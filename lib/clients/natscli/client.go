package natscli

import (
	"log/slog"

	"github.com/cenkalti/backoff/v4"
	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
)

func NewJetStream(cfg Config) (jetstream.JetStream, error) {
	if cfg.URL == "" {
		cfg.URL = nats.DefaultURL
	}

	var js jetstream.JetStream

	err := backoff.Retry(func() error {
		nc, err := nats.Connect(cfg.URL)
		if err != nil {
			slog.Error("error connecting to nats, backing off and retrying", "err", err, "url", cfg.URL)
			return err
		}

		js, err = jetstream.New(nc)
		if err != nil {
			slog.Error("error obtaining jetstream from nats connection, backing off and retrying", "err", err)
			return err
		}

		return nil
	}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5))

	if err != nil {
		return nil, err
	}

	return js, nil
}
