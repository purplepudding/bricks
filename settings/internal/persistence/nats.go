package persistence

import (
	"context"
	"errors"
	"fmt"

	"github.com/nats-io/nats.go/jetstream"
	"github.com/purplepudding/bricks/lib/clients/natscli"
	"github.com/purplepudding/bricks/settings/internal/core/settings"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/structpb"
)

var _ settings.GlobalSettingsStore = (*NATSSettingsStore)(nil)
var _ settings.ServiceSettingsStore = (*NATSSettingsStore)(nil)

const (
	bucket = "settings"
)

type NATSSettingsStore struct {
	js jetstream.JetStream
}

func NewNATSSettingsStore(cfg natscli.Config) (*NATSSettingsStore, error) {
	js, err := natscli.NewJetStream(cfg)
	if err != nil {
		return nil, err
	}

	ctx := context.Background()
	o, err := js.CreateOrUpdateObjectStore(ctx, jetstream.ObjectStoreConfig{
		Bucket:  bucket,
		Storage: jetstream.FileStorage,
	})
	if err != nil {
		return nil, err
	}
	if o == nil {
		return nil, errors.New("object store not created")
	}

	return &NATSSettingsStore{js: js}, nil
}

func (s *NATSSettingsStore) Get(ctx context.Context, collection string, id string) (map[string]*structpb.Value, error) {
	kv, err := s.js.KeyValue(ctx, fmt.Sprintf("%s.%s", bucket, collection))
	if err != nil {
		return nil, err
	}

	kve, err := kv.Get(ctx, id)
	if err != nil {
		return nil, err
	}

	var pb structpb.Struct
	if err := proto.Unmarshal(kve.Value(), &pb); err != nil {
		return nil, err
	}

	return pb.Fields, err
}

func (s *NATSSettingsStore) Set(ctx context.Context, collection string, id string, entries map[string]*structpb.Value) error {
	kv, err := s.js.KeyValue(ctx, fmt.Sprintf("%s.%s", bucket, collection))
	if err != nil {
		return err
	}

	b, err := proto.Marshal(&structpb.Struct{Fields: entries})
	if err != nil {
		return err
	}

	_, err = kv.Put(ctx, id, b)
	if err != nil {
		return err
	}

	return nil
}
