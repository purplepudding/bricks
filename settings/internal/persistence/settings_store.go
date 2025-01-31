package persistence

import (
	"context"
	"fmt"
	"log/slog"
	"maps"

	"github.com/purplepudding/foundation/settings/internal/core/settings"
	"github.com/valkey-io/valkey-go"
	"github.com/vmihailenco/msgpack/v5"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ settings.GlobalSettingsStore = (*ValkeySettingsStore)(nil)
var _ settings.ServiceSettingsStore = (*ValkeySettingsStore)(nil)

type ValkeySettingsStore struct {
	valkeyCli valkey.Client
}

func NewValkeySettingsStore(valkeyCli valkey.Client) *ValkeySettingsStore {
	return &ValkeySettingsStore{valkeyCli: valkeyCli}
}

func unpackEntries(ctx context.Context, entries map[string]*structpb.Value, prefix string) (map[string]string, error) {
	result := make(map[string]string, len(entries))

	for k, v := range entries {
		if sv := v.GetStructValue(); sv != nil {
			var newPrefix string
			if prefix == "" {
				newPrefix = k + ":"
			} else {
				newPrefix = fmt.Sprintf("%s:%s:", prefix, k)
			}

			unpacked, err := unpackEntries(ctx, sv.Fields, newPrefix)
			if err != nil {
				return nil, err
			}

			for k, v := range unpacked {
				result[k] = v
			}
			continue
		}

		// Not a struct, so we embed the value at this exact key
		b, err := msgpack.Marshal(v.AsInterface())
		if err != nil {
			slog.Error("error marshalling value with msgpack", "err", err)
			return nil, err
		}

		result[k] = string(b)
	}

	return result, nil
}

func (s *ValkeySettingsStore) Get(ctx context.Context, collection string) (map[string]*structpb.Value, error) {
	res := s.valkeyCli.Do(ctx, s.valkeyCli.B().Hgetall().Key(collection).Build())

	dbEntries, err := res.AsStrMap()
	if err != nil {
		//TODO sentinel and wrapping
		return nil, err
	}

	results := make(map[string]*structpb.Value, len(dbEntries))
	for k, v := range dbEntries {
		var val any
		err := msgpack.Unmarshal([]byte(v), &val)
		if err != nil {
			//TODO sentinel and wrapping
			return nil, err
		}

		sv, err := structpb.NewValue(val)
		if err != nil {
			//TODO sentinel and wrapping
			return nil, err
		}

		results[k] = sv
	}

	return results, nil
}

func (s *ValkeySettingsStore) Set(ctx context.Context, collection string, entries map[string]*structpb.Value) error {
	// Build marshalled entries for the hash
	marshalledEntries, err := unpackEntries(ctx, entries, "")
	if err != nil {
		return err
	}

	// Write values into Valkey hash
	res := s.valkeyCli.Do(ctx, s.valkeyCli.B().Hmset().Key(collection).FieldValue().FieldValueIter(maps.All(marshalledEntries)).Build())

	return res.Error()
}
