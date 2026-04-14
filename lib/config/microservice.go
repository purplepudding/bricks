package config

var _ Observable = (*Microservice)(nil)

type Microservice struct {
	ServingAddr          string `koanf:"servingAddr"`
	ObservabilityEnabled bool   `koanf:"observabilityEnabled"`
}

func (m Microservice) EnableObservability() bool {
	return m.ObservabilityEnabled
}

type Observable interface {
	EnableObservability() bool
}
