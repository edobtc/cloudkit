package droplet

import (
	"github.com/edobtc/cloudkit/config"

	"github.com/digitalocean/godo"
)

func NewClient() *godo.Client {
	return godo.NewFromToken(config.Read().DigitalOceanToken)
}
