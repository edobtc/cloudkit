package grpc

import (
	"context"

	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"
	"github.com/google/uuid"
	log "github.com/sirupsen/logrus"
)

func (s *ResourcesServiceServer) Submit(ctx context.Context, _req *pb.SubmitRequest) (*pb.SubmitResponse, error) {
	log.Debug("received submit resource request")

	log.Debug(_req)

	req := _req.Data

	log.Info(s.runtime)

	log.Debug(req)

	return &pb.SubmitResponse{
		Data: &pb.CreateResponse{
			Status: &pb.ResourceResponse{
				Success:    true,
				Identifier: uuid.New().String(),
			},
		},
	}, nil
}
