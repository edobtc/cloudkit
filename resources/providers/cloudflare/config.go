package cloudflare

// Config holds allowed values
// for an implemented resource provider. Any value outside of this config
// is unable to be modified during resource creation.
type Config struct {
	Zone   string
	ZoneID string
	Name   string
	TTL    int
	Value  string
}

func DefaultConfig() *Config {
	return &Config{
		TTL: 1,
	}
}
