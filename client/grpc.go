package client

import (
	"context"

	"github.com/edobtc/cloudkit/config"
	ctrl "github.com/edobtc/cloudkit/controlplane/grpc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func NewInsecureGrpcConn(ctx context.Context) (*grpc.ClientConn, error) {
	return grpc.DialContext(
		ctx,
		config.Read().GRPCListen,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	)
}

func NewSecureGrpcConnection(ctx context.Context) (*grpc.ClientConn, error) {
	creds, err := ctrl.LoadTLSClientCredentials()
	if err != nil {
		return nil, err
	}

	return grpc.DialContext(
		ctx,
		config.Read().GRPCListen,
		grpc.WithBlock(),
		grpc.WithTransportCredentials(creds),
	)
}
