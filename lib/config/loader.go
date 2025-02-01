package config

import (
	"log/slog"
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/parsers/yaml"
	"github.com/knadh/koanf/providers/env"
	"github.com/knadh/koanf/providers/rawbytes"
	"github.com/purplepudding/foundation/lib/config/settings"
)

const (
	EnvPrefix = "FOUNDATION_"
)

func Load(service string, embeddedCfg []byte, cfg any) error {
	k := koanf.New(".")

	// Load from embedded YAML string (defaults)
	if err := k.Load(rawbytes.Provider(embeddedCfg), yaml.Parser()); err != nil {
		return err
	}

	//TODO Load from YAML at path (general overrides)

	//TODO consider how to externalise settings service url
	if err := k.Load(settings.NewProvider("settings:8080", service), nil); err != nil {
		slog.Error("error loading configuration from settings service - continuing", "err", err)
	}

	// Load from environment (specific overrides, last resort)
	if err := k.Load(env.Provider(EnvPrefix, ".", func(s string) string {
		return strings.ToLower(strings.Replace(strings.TrimPrefix(s, EnvPrefix), "_", ".", -1))
	}), nil); err != nil {
		slog.Error("error loading configuration from environment variable - continuing", "err", err)
	}

	// Unmarshal into supplied config struct
	if err := k.Unmarshal("", &cfg); err != nil {
		return err
	}

	return nil
}
