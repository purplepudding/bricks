// Package memorymatch provides a channel-based matchmaking system for player pairing.
package memorymatch

import (
	"context"
	"sync"
)

// MemoryMatchClient implements thread-safe channel-based player matching.
type MemoryMatchClient struct {
	players map[string]*PlayerConnection // Map of player IDs to connection information
	mu      sync.Mutex                   // Mutex for thread-safe operations
}

// PlayerConnection encapsulates player data and channel for communication.
type PlayerConnection struct {
	PlayerID  string
	MessageCh chan []string
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
func (m *MemoryMatchClient) RequestMatch(ctx context.Context, playerID string, gameID string) (<-chan []string, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	println("locked for playerID: %s, gameID: %s", playerID, gameID)

	// If player already has a connection, return it
	if pc, exists := m.players[playerID]; exists {
		return pc.MessageCh, nil
	}

	// Else find the first available player to match with
	for key, pc := range m.players {
		if pc.PlayerID != playerID {
			println("found matching player: %s for playerID: %s", pc.PlayerID, playerID)

			newChan := make(chan []string, 1)

			newChan <- []string{pc.PlayerID}
			pc.MessageCh <- []string{playerID}

			delete(m.players, key)

			return newChan, nil
		}
	}

	// Else no match, create a new channel for this player and queue it
	println("no match found for playerID: %s, queuing", playerID)
	newChan := make(chan []string, 1)
	m.players[playerID] = &PlayerConnection{
		PlayerID:  playerID,
		MessageCh: newChan,
	}
	return newChan, nil
}
