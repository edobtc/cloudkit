package droplet

import "context"

// DomainConfig is a configuration for a domain record.
type DomainConfig struct {
	Name      string
	Alias     string
	DropletID string
	IP        string
}

// RecordDomain will take a domainConfiguration and
// create a domain record for the given domain.
func RecordDomain(ctx context.Context, cfg DomainConfig) error {
	return nil
}
