package grpc

import (
	"context"

	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

func (s *ResourcesServiceServer) List(ctx context.Context, req *pb.ListRequest) (*pb.ListResponse, error) {
	logrus.Debug("received list resource request", req)

	return &pb.ListResponse{
		Resources: []*pb.ResourceResponse{
			{Identifier: uuid.New().String()},
			{Identifier: uuid.New().String()},
			{Identifier: uuid.New().String()},
		},
	}, nil
}
