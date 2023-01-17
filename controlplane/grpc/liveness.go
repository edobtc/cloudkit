package grpc

import (
	"context"

	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"
	"github.com/edobtc/cloudkit/version"
	"github.com/sirupsen/logrus"
)

func (s *ResourcesServiceServer) Liveness(ctx context.Context, req *pb.LivenessRequest) (*pb.LivenessResponse, error) {
	logrus.Debug("received list resource request", req)

	return &pb.LivenessResponse{
		Message: version.String(),
	}, nil
}
