package autoload

import (
	"errors"

	"github.com/edobtc/cloudkit/resources/providers"
	"github.com/edobtc/cloudkit/resources/providers/aws/ec2"
	"github.com/edobtc/cloudkit/resources/providers/aws/ecs/fargate"
	"github.com/edobtc/cloudkit/resources/providers/aws/lambda"
	"github.com/edobtc/cloudkit/resources/providers/aws/route53"
	"github.com/edobtc/cloudkit/resources/providers/aws/security_groups"
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
	targetMap = map[pb.Target]string{
		pb.Target_TARGET_AWS_EC2:              "aws/ec2",
		pb.Target_TARGET_AWS_ECS:              "aws/fargate",
		pb.Target_TARGET_AWS_FARGATE:          "aws/fargate",
		pb.Target_TARGET_AWS_LAMBDA:           "aws/lambda",
		pb.Target_TARGET_AWS_ROUTE53:          "aws/route53",
		pb.Target_TARGET_AWS_SECURITY_GROUPS:  "aws/sg",
		pb.Target_TARGET_AWS_UNSPECIFIED:      "aws/ec2",
		pb.Target_TARGET_CLOUDFLARE:           "cloudflare",
		pb.Target_TARGET_DIGITALOCEAN_DROPLET: "digitalocean/droplet",
		pb.Target_TARGET_DIGITALOCEAN:         "digitalocean/droplet",
		pb.Target_TARGET_DOCKER:               "docker",
		pb.Target_TARGET_GCP:                  "gcp/gce",
		pb.Target_TARGET_K8S:                  "k8s/deployment",
		pb.Target_TARGET_KUBERNETES:           "k8s/deployment",
		pb.Target_TARGET_LINODE:               "linode",
		pb.Target_TARGET_MOCK_BLANK:           "test/blank",
		pb.Target_TARGET_MOCK_TIMED:           "test/timed",
	}

	registry = map[string]autoload{

		"cloudflare":           cloudflare.NewProvider,
		"digitalocean/droplet": droplet.NewProvider,
		"docker":               docker.NewProvider,
		"aws/route53":          route53.NewProvider,
		"aws/securitygroup":    security_groups.NewProvider,
		"aws/sg":               security_groups.NewProvider, // sg alias for security_groups
		"aws/lambda":           lambda.NewProvider,
		"aws/ec2":              ec2.NewProvisioner,
		"aws/fargate":          fargate.NewProvisioner,
		"aws/ecs":              fargate.NewProvisioner,

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

func ProtoTargetMap(t *pb.Target) string {
	if t == nil {
		return ""
	}

	if v, ok := targetMap[*t]; ok {
		return v
	}

	return ""
}
