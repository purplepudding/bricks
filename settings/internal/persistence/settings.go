package persistence

import (
	"context"
	"log/slog"

	"github.com/purplepudding/foundation/settings/internal/core/settings"
	"github.com/valkey-io/valkey-go"
	"github.com/vmihailenco/msgpack/v5"
)

var _ settings.SettingsStore = (*ValkeySettingsStore)(nil)

type ValkeySettingsStore struct {
	valkeyCli valkey.Client
}

func NewValkeySettingsStore(valkeyCli valkey.Client) *ValkeySettingsStore {
	return &ValkeySettingsStore{valkeyCli: valkeyCli}
}

func (s *ValkeySettingsStore) Set(ctx context.Context, collection string, entries map[string]any) error {
	// Build marshalled entries for the hash
	iter := func(yield func(string, string) bool) {
		for k, v := range entries {
			b, err := msgpack.Marshal(v)
			if err != nil {
				slog.Error("error marshalling value with msgpack", "err", err)
				return
			}

			if !yield(k, string(b)) {
				return
			}
		}
	}

	// Write values into Valkey hash
	res := s.valkeyCli.Do(ctx, s.valkeyCli.B().Hmset().Key(collection).FieldValue().FieldValueIter(iter).Build())

	return res.Error()
}
