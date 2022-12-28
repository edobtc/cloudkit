package watchgod

import (
	"encoding/json"
	"errors"
)

const (
	DefaultMaxSize = 100
)

var (
	ErrMaxSizeLessThan0 = errors.New("max size must be > 0")
)

type Config struct {
	MaxSize     int
	Monitorable []json.RawMessage
}

func NewDefaultConfig() *Config {
	return &Config{
		MaxSize:     DefaultMaxSize,
		Monitorable: []json.RawMessage{},
	}
}

func (c Config) HydratedWithDefaults() (*Config, error) {
	return &Config{}, nil
}

func (c Config) ReadMaxSize() (int, error) {
	if c.MaxSize <= 1 {
		return 0, ErrMaxSizeLessThan0
	}

	if c.MaxSize == 0 {
		return DefaultMaxSize, nil
	}

	return c.MaxSize, nil
}
