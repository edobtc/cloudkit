package security_groups

import "github.com/edobtc/cloudkit/labels"

type Config struct {
	VpcId  string
	DryRun bool
	All    bool
	Kind   string
	Tags   []labels.Label // optionally define some extra tags
}

func MergeDefaultConfig(cfg *Config) *Config {
	if cfg.Tags == nil {
		cfg.Tags = []labels.Label{}
	}

	return cfg
}

func DefaultConfig() *Config {
	return &Config{
		Tags: []labels.Label{},
	}
}
