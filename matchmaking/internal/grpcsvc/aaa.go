package grpcsvc

import (
	matchmakingv1 "github.com/purplepudding/bricks/api/pkg/pb/bricks/v1/matchmaking"
)

var _ matchmakingv1.MatchmakingServiceServer = (*MatchmakingService)(nil)

type MatchmakingService struct {
	matchmakingv1.UnimplementedMatchmakingServiceServer
}
