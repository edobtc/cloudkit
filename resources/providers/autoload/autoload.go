package autoload

import (
	"errors"

	"github.com/edobtc/cloudkit/resources/providers"
	"github.com/edobtc/cloudkit/resources/providers/aws/ec2"
	"github.com/edobtc/cloudkit/resources/providers/aws/lambda"
	"github.com/edobtc/cloudkit/resources/providers/cloudflare"
	"github.com/edobtc/cloudkit/resources/providers/digitalocean/droplet"
	"github.com/edobtc/cloudkit/resources/providers/docker"
	"github.com/edobtc/cloudkit/resources/providers/mock/blank"
	"github.com/edobtc/cloudkit/resources/providers/mock/timed"

	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"
)

type autoload func(req *pb.CreateRequest) providers.Provider

type autoloadRegistration func(req *pb.Registration) providers.Provider

var (
	registry = map[string]autoload{

		"cloudflare":           cloudflare.NewProvider,
		"digitalocean/droplet": droplet.NewProvider,
		"docker":               docker.NewProvider,
		"aws/lambda":           lambda.NewProvider,
		"aws/ec2":              ec2.NewProvisioner,

		// mock/test providers for testing integrations
		// when running locally and wanting to make full requests
		// but NOT mess around with any resources
		// also useful in a test suite
		"test/blank": blank.NewProvisioner,
		"test/timed": timed.NewProvisioner,
	}

	registrations = map[string]autoloadRegistration{
		"cloudflare": cloudflare.NewRegistrationProvider,
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
func Load(p string, req *pb.CreateRequest) (providers.Provider, error) {

	if prv, ok := registry[p]; ok {
		return prv(req), nil
	}

	return nil, ErrProviderNotFound
}

// Load automatically loads a provider's provisioner from the
// registry
func LoadRegistration(p string, req *pb.Registration) (providers.Provider, error) {
	if prv, ok := registrations[p]; ok {
		return prv(req), nil
	}

	return nil, ErrProviderNotFound
}

// Exists allows you to validate and verify a given provider exists
// without having to also posses config we want to pass to the provider
func Exists(p string) bool {
	_, ok := registry[p]
	return ok
}
