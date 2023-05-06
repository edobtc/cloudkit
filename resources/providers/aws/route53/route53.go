package route53

import (
	"errors"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
	"github.com/cloudflare/cloudflare-go"
	"github.com/edobtc/cloudkit/labels"
	"github.com/edobtc/cloudkit/resources/providers"
	"github.com/edobtc/cloudkit/resources/providers/aws/auth"
	"github.com/edobtc/cloudkit/target"

	log "github.com/sirupsen/logrus"

	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"
)

var (
	ErrZoneNotFound = errors.New("no zone found")
)

// Route53Provider implements a Route53Provider
type Route53Provider struct {
	// Config holds our internal configuration options
	// for the instance of the Route53Provider
	Config *Config
	svc    *route53.Route53

	Zone    string
	Records []cloudflare.DNSRecord
}

// NewProvider initializes a Route53Provider
// with defaults
func NewProvider(req *pb.CreateRequest) providers.Provider {
	sess, err := auth.Session()
	if err != nil {
		return nil
	}

	svc := route53.New(sess)

	return &Route53Provider{
		Config: &Config{
			Zone: req.Config.Name,
		},
		svc: svc,
	}
}

// NewRegistrationProvider initializes a Route53Provider
// for use with a registration callback vs a full blown
// provider implementation, extracting and mapping the correct
// part of the request configuration
func NewRegistrationProvider(req *pb.Registration) providers.Provider {
	sess, err := auth.Session()
	if err != nil {
		return nil
	}

	svc := route53.New(sess)

	return &Route53Provider{
		Config: &Config{
			Zone: req.Identifier,
		},
		svc: svc,
	}
}

func NewProviderFromConfig(cfg *Config) (providers.Provider, error) {
	sess, err := auth.Session()
	if err != nil {
		return nil, err
	}

	svc := route53.New(sess)

	return &Route53Provider{
		Config: MergeDefaultConfig(cfg),
		svc:    svc,
	}, nil
}

// Select is similar to Read yet copies a selection of
// resources based on the Target configuration
func (p *Route53Provider) Select() (target.Selection, error) {
	listZonesInput := &route53.ListHostedZonesByNameInput{
		DNSName: aws.String(p.Config.Zone),
	}
	listZonesOutput, err := p.svc.ListHostedZonesByName(listZonesInput)
	if err != nil {
		return target.Selection{}, err
	}

	if len(listZonesOutput.HostedZones) == 0 {
		return target.Selection{}, ErrZoneNotFound
	}

	log.Debug(listZonesOutput.HostedZones[0])

	// The hosted zone name includes a trailing dot,
	// #which must be removed before using the zone ID.
	zoneId := aws.StringValue(listZonesOutput.HostedZones[0].Id)
	zoneId = zoneId[:len(zoneId)-1]

	return target.Selection{
		Data: target.Resources{
			target.Resource{
				Name: aws.StringValue(listZonesOutput.HostedZones[0].Name),
				ID:   zoneId,
				Meta: listZonesOutput.HostedZones[0],
			},
		},
	}, nil
}

// Read fetches and stores the configuration for an existing
// cloudflare zone. What is read of the existing resource acts
// as the template/configuration to implement a clone via creating a
// new resource with the existing output as input for a variant
func (p *Route53Provider) Read() error {
	return nil
}

// Clone creates a modified variant
func (p *Route53Provider) Clone() error {
	return nil
}

// ProbeReadiness checks that the provisioned resource is available and
// ready to be included in a request
func (p *Route53Provider) ProbeReadiness() (bool, error) {
	return false, nil
}

// Teardown eradicates any resource that has been
// provisioned as part of a variant
func (p *Route53Provider) Teardown() error {
	// Needs to look up variants based on
	// labels / tags which identify a variant name, experiment,
	// and ideally a namespace
	return nil
}

// Cancel will abort and running or submitted Route53Provider
func (p *Route53Provider) Cancel() error { return nil }

// Stop will stop any running Route53Provider
func (p *Route53Provider) Stop() error { return nil }

// AwaitReadiness should be implemented to detect
// when a Route53Provider has finished setting up a variant
// and can begin using it in an experiment
func (p *Route53Provider) AwaitReadiness() chan error { return make(chan error) }

// Annotate should implement applying labels or tags for a given resource type
func (p *Route53Provider) Annotate(id string, l labels.Labels) error { return nil }
