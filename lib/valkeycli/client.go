package valkeycli

import (
	"log/slog"

	"github.com/cenkalti/backoff/v4"
	"github.com/valkey-io/valkey-go"
)

func New(cfg Config) (valkey.Client, error) {
	var valkeyCli valkey.Client

	err := backoff.Retry(func() error {
		var err error
		valkeyCli, err = valkey.NewClient(valkey.ClientOption{InitAddress: []string{cfg.Addr}})
		if err != nil {
			slog.Error("error connecting to valkey, backing off and retrying", "err", err)
			return err
		}
		return nil
	}, backoff.WithMaxRetries(backoff.NewExponentialBackOff(), 5))

	if err != nil {
		return nil, err
	}
	return valkeyCli, nil
}
