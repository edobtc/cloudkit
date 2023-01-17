package grpc

import (
	"context"

	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"

	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (s *ResourcesServiceServer) Status(ctx context.Context, req *pb.StatusRequest) (*pb.StatusResponse, error) {
	logrus.Debug("received list resource request", req)

	return &pb.StatusResponse{
		Resource: &pb.ResourceResponse{
			Identifier: uuid.New().String(),
		},
	}, nil
}
