package cloudflare

import (
	"context"
	"log"

	"github.com/cloudflare/cloudflare-go"
	"github.com/edobtc/cloudkit/config"
	"github.com/edobtc/cloudkit/labels"
	"github.com/edobtc/cloudkit/resources/providers"
	"github.com/edobtc/cloudkit/target"
	"gopkg.in/yaml.v2"
)

// CloudflareProvider implements a CloudflareProvider
type CloudflareProvider struct {
	// Config holds our internal configuration options
	// for the instance of the CloudflareProvider
	Config *Config
	client *cloudflare.API

	Zone    cloudflare.Zone
	Records []cloudflare.DNSRecord
}

// NewProvider initializes a CloudflareProvider
// with defaults
func NewProvider(yml []byte) providers.Provider {
	cfg := Config{}
	err := yaml.Unmarshal(yml, &cfg)

	if err != nil {
		return nil
	}

	return &CloudflareProvider{
		Config: &cfg,
	}
}

func NewProviderFromConfig(cfg *Config) (providers.Provider, error) {
	client, err := cloudflare.NewWithAPIToken(config.Read().CloudflareAPIToken)

	if err != nil {
		return nil, err
	}
	return &CloudflareProvider{
		Config: cfg,
		client: client,
	}, nil
}

// Select is similar to Read yet copies a selection of
// resources based on the Target configuration
func (p *CloudflareProvider) Select() (target.Selection, error) {
	ctx := context.Background()

	if p.Config.Zone != "" && p.Config.ZoneID == "" {
		id, err := p.client.ZoneIDByName(p.Config.Zone)
		if err != nil {
			return target.Selection{}, err
		}
		p.Config.ZoneID = id
	}

	zone, err := p.client.ZoneDetails(ctx, p.Config.ZoneID)
	if err != nil {
		log.Fatal(err)
	}

	p.Zone = zone

	records, err := p.client.DNSRecords(ctx, p.Config.ZoneID, cloudflare.DNSRecord{})
	if err != nil {
		return target.Selection{}, err
	}

	p.Records = records

	return target.Selection{
		Data: target.Resources{
			target.Resource{
				Name: p.Zone.Name,
				ID:   p.Zone.ID,
				Meta: p.Records,
			},
		},
	}, nil
}

// Read fetches and stores the configuration for an existing
// cloudflare zone. What is read of the existing resource acts
// as the template/configuration to implement a clone via creating a
// new resource with the existing output as input for a variant
func (p *CloudflareProvider) Read() error {
	return nil
}

// Clone creates a modified variant
func (p *CloudflareProvider) Clone() error {
	return nil
}

// ProbeReadiness checks that the provisioned resource is available and
// ready to be included in a request
func (p *CloudflareProvider) ProbeReadiness() (bool, error) {
	return false, nil
}

// Teardown eradicates any resource that has been
// provisioned as part of a variant
func (p *CloudflareProvider) Teardown() error {
	// Needs to look up variants based on
	// labels / tags which identify a variant name, experiment,
	// and ideally a namespace
	return nil
}

// Cancel will abort and running or submitted CloudflareProvider
func (p *CloudflareProvider) Cancel() error { return nil }

// Stop will stop any running CloudflareProvider
func (p *CloudflareProvider) Stop() error { return nil }

// AwaitReadiness should be implemented to detect
// when a CloudflareProvider has finished setting up a variant
// and can begin using it in an experiment
func (p *CloudflareProvider) AwaitReadiness() chan error { return make(chan error) }

// Annotate should implement applying labels or tags for a given resource type
func (p *CloudflareProvider) Annotate(id string, l labels.Labels) error { return nil }
