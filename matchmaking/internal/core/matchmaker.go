package core

import (
	"context"

	memorymatch "github.com/purplepudding/bricks/matchmaking/internal/memorymatch"
)

// MatchmakingClient defines an interface for matchmaking service operations
// that can query and respond to matchmaking requests based on a specific game ID.
type MatchmakingClient interface {
	RequestMatch(ctx context.Context, playerID string, gameID string) ([]string, error)
}

// Matchmaker implements the matchmaking logic using an internal MemoryMatchClient
// to handle player pair matching.
type Matchmaker struct {
	matchmakingClient *memorymatch.MemoryMatchClient // Thread-safe client for matchmaking
}

// NewMatchmaker creates a new instance of Matchmaker with a newly initialized MemoryMatchClient.
func NewMatchmaker() *Matchmaker {
	return &Matchmaker{
		matchmakingClient: memorymatch.NewMemoryMatchClient(),
	}
}

// RequestMatch delegates to the MemoryMatchClient to handle the actual request.
// Returns appropriate response based on whether a player was found to match with.
func (m *Matchmaker) RequestMatch(ctx context.Context, playerID, playerAddr string) (<-chan []string, error) {

	return m.matchmakingClient.RequestMatch(ctx, playerID, "gameID") //TODO more fleshed out impl here
}
