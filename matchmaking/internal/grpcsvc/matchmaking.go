package grpcsvc

import (
	"context"
	"fmt"
	"time"

	matchmakingv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/matchmaking"
)

// Matchmaker interface for matchmaking logic
type Matchmaker interface {
	RequestMatch(ctx context.Context, playerID string) (<-chan []string, error)
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
	//TODO Extract player ID from the request
	playerID := fmt.Sprintf("player-%s", time.Now())

	// Call the matchmaker interface
	resChan, err := s.matchmaker.RequestMatch(svr.Context(), playerID)
	if err != nil {
		return err
	}

	err = svr.Send(&matchmakingv1.RequestMatchResponse{
		Update: &matchmakingv1.RequestMatchResponse_AwaitingMatch{},
	})
	if err != nil {
		return err
	}

	ctx, cancel := context.WithTimeout(svr.Context(), 25*time.Second)
	defer cancel()

	select {
	case <-ctx.Done():
		return ctx.Err()
	case result := <-resChan:

		// Print the debug info
		fmt.Printf("Matchmaker RequestMatch Results: %s\n", result)
		fmt.Printf("Matchmaker RequestMatch Error: %v\n", err)

		return svr.Send(&matchmakingv1.RequestMatchResponse{
			Update: &matchmakingv1.RequestMatchResponse_MatchFound{},
		})
	}
}
