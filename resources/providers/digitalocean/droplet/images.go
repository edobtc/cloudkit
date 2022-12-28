package droplet

import (
	"context"
	"strings"

	"github.com/digitalocean/godo"
)

type Image struct {
	ID         string
	Name       string
	ResourceID string
}

func FilterImages(ctx context.Context, pattern string) ([]Image, error) {
	var images []Image

	client := NewClient()
	opt := &godo.ListOptions{
		Page:    InitPageDefault,
		PerPage: PageSizeDefault,
	}

	snapshots, _, err := client.Snapshots.List(ctx, opt)
	if err != nil {
		return images, err
	}

	for _, snp := range snapshots {
		if strings.Contains(snp.Name, pattern) {
			images = append(images, Image{
				Name:       snp.Name,
				ID:         snp.ID,
				ResourceID: snp.ResourceID,
			})
		}
	}

	return images, nil
}
