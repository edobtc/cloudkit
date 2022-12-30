package cloudflare

import (
	"context"

	"github.com/cloudflare/cloudflare-go"
	"github.com/sirupsen/logrus"
)

// Apply runs the CloudflareProvider end to end
// which in this instance is only associated with
// creating a dns record
func (p *CloudflareProvider) Apply() error {
	selection, err := p.Select()
	if err != nil {
		return err
	}

	logrus.Debug(selection)

	ctx := context.TODO()
	_, err = p.client.CreateDNSRecord(ctx, p.Config.ZoneID, cloudflare.DNSRecord{
		Type:    "A",
		Name:    p.Config.Name,
		Content: p.Config.Value,
		TTL:     p.Config.TTL,
	})
	return err
}
