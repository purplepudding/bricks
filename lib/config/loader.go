package config

import (
	"strings"

	"github.com/knadh/koanf"
	"github.com/knadh/koanf/providers/env"
)

const (
	EnvPrefix = "FOUNDATION_"
)

func Load(cfg any) error {
	k := koanf.New(".")

	//TODO Load from embedded YAML string (defaults)

	//TODO Load from YAML at path (general overrides)

	// Load from environment (specific overrides, last resort)
	if err := k.Load(env.Provider(EnvPrefix, ".", func(s string) string {
		return strings.ToLower(strings.Replace(strings.TrimPrefix(s, EnvPrefix), "_", ".", -1))
	}), nil); err != nil {
		return err
	}

	// Unmarshal into supplied config struct
	if err := k.Unmarshal("", &cfg); err != nil {
		return err
	}

	return nil
}
