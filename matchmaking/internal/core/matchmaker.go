package core

import (
	"context"
)

type Player struct {
	ID   string
	Addr string
}

// MatchmakingClient defines an interface for matchmaking service operations
// that can query and respond to matchmaking requests based on a specific game ID.
type MatchmakingClient interface {
	RequestMatch(ctx context.Context, player Player, gameID string) (<-chan []Player, error)
}

// Matchmaker implements the matchmaking logic using an internal MemoryMatchClient
// to handle player pair matching.
type Matchmaker struct {
	matchmakingClient MatchmakingClient // Thread-safe client for matchmaking
}

// NewMatchmaker creates a new instance of Matchmaker with a newly initialized MemoryMatchClient.
func NewMatchmaker(matchmakingClient MatchmakingClient) *Matchmaker {
	return &Matchmaker{
		matchmakingClient: matchmakingClient,
	}
}

// RequestMatch delegates to the MemoryMatchClient to handle the actual request.
// Returns appropriate response based on whether a player was found to match with.
func (m *Matchmaker) RequestMatch(ctx context.Context, playerID, playerAddr string) (<-chan []Player, error) {

	return m.matchmakingClient.RequestMatch(ctx, Player{
		ID:   playerID,
		Addr: playerAddr,
	}, "gameID")
}
