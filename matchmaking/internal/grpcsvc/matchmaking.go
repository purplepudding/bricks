package grpcsvc

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/grpc-ecosystem/go-grpc-middleware/v2/interceptors/realip"
	matchmakingv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/matchmaking"
	"github.com/purplepudding/bricks/matchmaking/internal/core"
	"go.opentelemetry.io/otel/attribute"
	"go.opentelemetry.io/otel/codes"
	"go.opentelemetry.io/otel/trace"
	"google.golang.org/grpc/metadata"
)

// Matchmaker interface for matchmaking logic
type Matchmaker interface {
	RequestMatch(ctx context.Context, playerID, playerAddr string) (<-chan []core.Player, error)
}

// MatchmakingService holds matchmaking service logic
type MatchmakingService struct {
	matchmakingv1.UnimplementedMatchmakingServiceServer
	matchmaker Matchmaker // Matchmaker interface implementation
}

// NewMatchmakingService creates a new instance of MatchmakingService with provided matchmaker
func NewMatchmakingService(m Matchmaker) *MatchmakingService {
	return &MatchmakingService{
		matchmaker: m,
	}
}

// RequestMatch implements the MatchmakingService server method, delegates to the internal matchmaker
func (s *MatchmakingService) RequestMatch(req *matchmakingv1.RequestMatchRequest, svr matchmakingv1.MatchmakingService_RequestMatchServer) error {
	ctx := svr.Context()

	//TODO pull this up into a general interceptor func
	span := trace.SpanFromContext(ctx)
	md, ok := metadata.FromIncomingContext(ctx)
	if !ok {
		return s.reportError(span, fmt.Errorf("metadata not found"))
	}
	for k, v := range md {
		span.SetAttributes(
			attribute.StringSlice("metadata."+k, v),
		)
	}

	//TODO Extract player ID from the request
	playerID := fmt.Sprintf("player-%s", time.Now())
	playerIP, found := realip.FromContext(ctx)
	if !found {
		return s.reportError(span, fmt.Errorf("player IP not found, %q", playerIP))
	}

	span.SetAttributes(
		attribute.String("player.id", playerID),
		attribute.String("player.ip", playerIP.StringExpanded()),
	)

	if req.Port == 0 {
		req.Port = 7777
	}
	playerAddr := fmt.Sprintf("%s:%d", playerIP, req.Port)

	span.SetAttributes(
		attribute.Int("player.port", int(req.Port)),
		attribute.String("player.addr", playerAddr),
	)

	// Call the matchmaker interface
	resChan, err := s.matchmaker.RequestMatch(ctx, playerID, playerAddr)
	if err != nil {
		return s.reportError(span, err)
	}

	// Send an awaiting match response
	err = svr.Send(&matchmakingv1.RequestMatchResponse{
		Update: &matchmakingv1.RequestMatchResponse_AwaitingMatch{},
	})
	if err != nil {
		return s.reportError(span, err)
	}

	ctx, cancel := context.WithTimeout(ctx, 25*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		return s.reportError(span, ctx.Err())
	case result := <-resChan:
		var players []*matchmakingv1.Player
		for _, p := range result {
			players = append(players, &matchmakingv1.Player{
				Id:   p.ID,
				Addr: p.Addr,
			})
		}

		return svr.Send(&matchmakingv1.RequestMatchResponse{
			Update: &matchmakingv1.RequestMatchResponse_MatchFound{
				MatchFound: &matchmakingv1.MatchFound{
					MatchId: uuid.New().String(),
					Players: players,
				},
			},
		})
	}
}

// TODO move this to a common lib and reuse
func (s *MatchmakingService) reportError(span trace.Span, err error) error {
	span.RecordError(err)
	span.SetStatus(codes.Error, err.Error())
	return err
}
