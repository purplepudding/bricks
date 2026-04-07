// Package memorymatch provides a channel-based matchmaking system for player pairing.
package memorymatch

import (
	"context"
	"sync"

	"github.com/purplepudding/bricks/matchmaking/internal/core"
)

// MemoryMatchClient implements thread-safe channel-based player matching.
type MemoryMatchClient struct {
	players map[string]*PlayerConnection // Map of player IDs to connection information
	mu      sync.Mutex                   // Mutex for thread-safe operations
}

// PlayerConnection encapsulates player data and channel for communication.
type PlayerConnection struct {
	Player    core.Player
	MessageCh chan []core.Player
}

// NewMemoryMatchClient returns an initialized MemoryMatchClient.
func NewMemoryMatchClient() *MemoryMatchClient {
	return &MemoryMatchClient{
		players: make(map[string]*PlayerConnection),
	}
}

// RequestMatch implements the channel-based matchmaking system:
//   - If queue is empty, creates a new connection and returns nil (player has a channel)
//   - If queue has players, finds a matching player and returns MatchFound with a channel
//   - If matching fails, player is requeued
func (m *MemoryMatchClient) RequestMatch(ctx context.Context, player core.Player, gameID string) (<-chan []core.Player, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	println("locked for playerID: %s, gameID: %s", player.ID, gameID)

	// If player already has a connection, return it
	if pc, exists := m.players[player.ID]; exists {
		return pc.MessageCh, nil
	}

	newChan := make(chan []core.Player, 1)

	// Else find the first available player to match with
	for key, pc := range m.players {
		if pc.Player.ID != player.ID {
			println("found matching player: %s for playerID: %s", pc.Player.ID, player.ID)

			newChan <- []core.Player{pc.Player}
			pc.MessageCh <- []core.Player{player}

			delete(m.players, key)

			return newChan, nil
		}
	}

	// Else no match, create a new channel for this player and queue it
	println("no match found for playerID: %s, queuing", player)
	m.players[player.ID] = &PlayerConnection{
		Player:    player,
		MessageCh: newChan,
	}
	return newChan, nil
}
