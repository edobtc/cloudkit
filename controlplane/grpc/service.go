package grpc

import (
	"crypto/tls"
	"crypto/x509"
	"fmt"
	"os"

	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
)

type ResourcesServiceServer struct {
	Close chan os.Signal
	Error chan error

	pb.UnimplementedResourcesServer
}

func NewResourcesServiceServer() *ResourcesServiceServer {
	return &ResourcesServiceServer{
		Close: make(chan os.Signal),
		Error: make(chan error),
	}
}

func NewTlSServer() (*grpc.Server, error) {
	tlsCredentials, err := LoadTLSCredentials()
	if err != nil {
		return nil, err
	}

	grpcServer := grpc.NewServer(
		grpc.Creds(tlsCredentials),
	)

	return grpcServer, nil
}

func LoadTLSCredentials() (credentials.TransportCredentials, error) {
	// Load server's certificate and private key
	serverCert, err := tls.LoadX509KeyPair("cert/server-cert.pem", "cert/server-key.pem")
	if err != nil {
		return nil, err
	}

	// Create the credentials and return it
	config := &tls.Config{
		Certificates: []tls.Certificate{serverCert},
		ClientAuth:   tls.NoClientCert,
	}

	return credentials.NewTLS(config), nil
}

func LoadTLSClientCredentials() (credentials.TransportCredentials, error) {
	// Load certificate of the CA who signed server's certificate
	pemServerCA, err := os.ReadFile("cert/ca-cert.pem")
	if err != nil {
		return nil, err
	}

	certPool := x509.NewCertPool()
	if !certPool.AppendCertsFromPEM(pemServerCA) {
		return nil, fmt.Errorf("failed to add server CA's certificate")
	}

	// Create the credentials and return it
	config := &tls.Config{
		RootCAs: certPool,
	}

	return credentials.NewTLS(config), nil
}
