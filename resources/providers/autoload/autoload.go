package autoload

import (
	"errors"

	"github.com/edobtc/cloudkit/resources/providers"
	"github.com/edobtc/cloudkit/resources/providers/aws/lambda"
	"github.com/edobtc/cloudkit/resources/providers/cloudflare"
	"github.com/edobtc/cloudkit/resources/providers/digitalocean/droplet"
	"github.com/edobtc/cloudkit/resources/providers/docker"

	"gopkg.in/yaml.v2"
)

type autoload func(cfg []byte) providers.Provider

var (
	registry = map[string]autoload{
		"aws/lambda":           lambda.NewProvider,
		"cloudflare":           cloudflare.NewProvider,
		"digitalocean/droplet": droplet.NewProvider,
		"docker":               docker.NewProvider,
		// "aws/ec2":         ec2.NewProvisioner,
		// "aws/elasticache": elasticache.NewProvisioner,
		// "k8s/pods":        pods.NewProvisioner,
		// "k8s/deployment":  deployment.NewProvisioner,
	}

	// ErrProviderNotFound is returned when we have failed to
	// lookup a provider
	ErrProviderNotFound = errors.New("provider not implemented or registered")

	// ErrConfigSerializationFailed is returned when we have failed to
	// serialize the generic parsed config back to yaml
	ErrConfigSerializationFailed = errors.New("failed to serialize GenericConfig to yaml")
)

// Load automatically loads a provider's provisioner from the
// registry
func Load(p string, cfg providers.GenericConfig) (providers.Provider, error) {
	config, err := yaml.Marshal(cfg)

	if err != nil {
		return nil, ErrConfigSerializationFailed
	}

	if prv, ok := registry[p]; ok {
		return prv(config), nil
	}

	return nil, ErrProviderNotFound
}

// Exists allows you to validate and verify a given provider exists
// without having to also posses config we want to pass to the provider
func Exists(p string) bool {
	_, ok := registry[p]
	return ok
}
