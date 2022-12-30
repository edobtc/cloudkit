package cloudflare

const (
	DefaultTTL  = 1
	DefaultType = "A"
)

// Config holds allowed values
// for an implemented resource provider. Any value outside of this config
// is unable to be modified during resource creation.
type Config struct {
	Zone   string
	ZoneID string
	Name   string
	TTL    int
	Value  string
	Type   string
}

func MergeDefaultConfig(cfg *Config) *Config {
	if cfg.TTL == 0 {
		cfg.TTL = DefaultTTL
	}

	if cfg.Type == "" {
		cfg.Type = DefaultType
	}

	return cfg
}

func DefaultConfig() *Config {
	return &Config{
		Type: DefaultType,
		TTL:  DefaultTTL,
	}
}
