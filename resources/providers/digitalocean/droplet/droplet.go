package droplet

import (
	"context"
	"errors"
	"strconv"

	"github.com/digitalocean/godo"
	"github.com/sirupsen/logrus"
)

func CreateDroplet(ctx context.Context, cfg *Config) (*godo.Droplet, error) {
	client := NewClient()

	imageID, err := strconv.Atoi(cfg.ImageID)
	if err != nil {
		return nil, errors.New("invalid image id")
	}

	createRequest := &godo.DropletCreateRequest{
		Name:   cfg.Name,
		Region: "nyc3",
		Size:   "s-2vcpu-2gb",
		Image: godo.DropletCreateImage{
			ID: imageID,
		},
		UserData:   "",
		Backups:    true,
		IPv6:       true,
		Monitoring: true,
		SSHKeys: []godo.DropletCreateSSHKey{
			{Fingerprint: cfg.SSHKey},
		},
	}

	newDroplet, _, err := client.Droplets.Create(ctx, createRequest)

	if err != nil {
		logrus.Error(err)
		return nil, err
	}

	return newDroplet, nil
}
