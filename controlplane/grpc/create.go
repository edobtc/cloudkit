package grpc

import (
	"context"

	"github.com/edobtc/cloudkit/resources/providers/digitalocean/droplet"
	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (s *ResourcesServiceServer) Create(ctx context.Context, req *pb.CreateRequest) (*pb.CreateResponse, error) {
	logrus.Info("received create resource request")

	logrus.Info(ctx)
	logrus.Info(req)

	droplet.Apply(req.Config.Name)

	return &pb.CreateResponse{
		Status: &pb.ResourceResponse{
			Success:    true,
			Identifier: uuid.New().String(),
		},
	}, nil

}
