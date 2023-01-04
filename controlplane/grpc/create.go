package grpc

import (
	"context"
	"encoding/json"
	"sync"

	"github.com/edobtc/cloudkit/config"
	"github.com/edobtc/cloudkit/controlplane/ws"
	"github.com/edobtc/cloudkit/events/publishers/webhook"
	"github.com/edobtc/cloudkit/resources/providers/cloudflare"
	"github.com/edobtc/cloudkit/resources/providers/digitalocean/droplet"
	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (s *ResourcesServiceServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	logrus.Debug("received create resource request")

	req, err := validateAndSetDefaults(req)
	if err != nil {
		return nil, err
	}

	logrus.Debug(req)

	response, err := droplet.Apply(req.Config.Name)
	if err != nil {
		return nil, err
	}

	wg := sync.WaitGroup{}

	for _, registration := range req.Registrations {
		wg.Add(1)
		go func(response *pb.ResourceResponse, registration *pb.Registration) {
			logrus.Infof("registering %s", registration.Name)

			p, _ := cloudflare.NewProviderFromConfig(&cloudflare.Config{
				Zone:  registration.Identifier,
				Type:  "A",
				Name:  response.Name,
				Value: response.Ip,
			})

			err = p.Apply()
			if err != nil {
				logrus.Error(err)
			}
		}(response, registration)
	}

	if data, err := json.Marshal(response); err == nil {
		if config.Read().Notifications.AllowWebsocketSubscribers {
			wg.Add(1)
			go func(d []byte) {
				ws.Pool.Publish(d)
			}(data)
		}

		if config.Read().Notifications.WebhookURL != "" {
			wg.Add(1)
			go func(d []byte) {
				wh, err := webhook.NewPublisher()
				if err != nil {
					logrus.Error(err)
					return
				}

				err = wh.Send(d)
				if err != nil {
					logrus.Error(err)
				}
			}(data)
		}
	} else {
		logrus.Error(err)
	}

	wg.Wait()

	return &pb.CreateResponse{
		Status: &pb.ResourceResponse{
			Success:    true,
			Identifier: uuid.New().String(),
		},
	}, nil
}

func validateAndSetDefaults(req *pb.CreateRequest) (*pb.CreateRequest, error) {
	if req.Config == nil {
		req.Config = &pb.ResourceConfiguration{}
	}

	if req.Config.Name == "" {
		req.Config.Name = uuid.New().String()
	}

	if req.Config.Region == "" {
		req.Config.Region = droplet.DefaultRegion
	}

	if req.Config.Size == "" {
		req.Config.Size = droplet.DefaultSize
	}

	if req.Config.Version == "" {
		req.Config.Version = "lnd15.03"
	}

	return req, nil
}
