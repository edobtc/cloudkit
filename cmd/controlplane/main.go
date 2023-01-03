package main

import (
	"context"
	"net"
	"net/http"
	"os"
	"os/signal"
	"syscall"

	"github.com/sirupsen/logrus"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"

	"github.com/grpc-ecosystem/grpc-gateway/v2/runtime"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	// Services
	service "github.com/edobtc/cloudkit/controlplane/grpc"
	"github.com/edobtc/cloudkit/controlplane/ws"
	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"
	api "github.com/edobtc/cloudkit/rpc/gateway/controlplane/resources/v1"

	"github.com/edobtc/cloudkit/config"
	"github.com/edobtc/cloudkit/controlplane"

	"github.com/edobtc/cloudkit/http/middleware/auth/key"
	"github.com/edobtc/cloudkit/http/middleware/namespace"

	"github.com/edobtc/cloudkit/logging"
	"github.com/edobtc/cloudkit/server"
)

func init() {
	logging.Setup()
}

var rootCmd = &cobra.Command{
	Use:   "control-plane",
	Short: "bck is the main command to launch the notification server",
	Run: func(cmd *cobra.Command, args []string) {

		go func() {
			serveGRPC()
		}()

		<-serveHTTP()
	},
}

func serveHTTP() chan bool {
	srv := server.NewServer()

	r := srv.Router()

	// middleware attachments
	r.Use(namespace.DecorateNamespace) // namespace middleware
	r.Use(ws.InjectConnectionPool)     // maintenance of websocket connections

	// route bindings
	controlplane.BindRoutes(r)

	// optional API key enabled
	if config.Read().EnableApiKey {
		log.Info("enabling api key")
		r.Use(key.ApiKey)
	}

	log.Info("starting control plane")

	signal.Notify(srv.Close, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	go func() {
		s := <-srv.Close
		log.Infof("received signal %v", s)
		srv.GracefulShutdown()
	}()

	return srv.Start()
}

func serveGRPCGateway() chan error {
	await := make(chan error)
	listen := config.Read().GRPCGatewayListen

	mux := runtime.NewServeMux()
	proxyOpts := []grpc.DialOption{
		grpc.WithTransportCredentials(insecure.NewCredentials()),
	}

	err := api.RegisterResourcesHandlerFromEndpoint(
		context.Background(),
		mux,
		config.Read().GRPCListen,
		proxyOpts,
	)

	if err != nil {
		panic(err)
	}

	go func() {
		log.Infof("grpc gateway starting and listening on %s", listen)
		await <- http.ListenAndServe(listen, mux)
	}()

	return await
}

func serveGRPC() chan bool {
	listen := config.Read().GRPCListen
	lis, err := net.Listen("tcp", listen)
	if err != nil {
		panic("failed to listen")
	}

	srv := service.NewResourcesServiceServer()

	signal.Notify(srv.Close, os.Interrupt, syscall.SIGTERM, syscall.SIGINT)

	grpcServer := grpc.NewServer()

	pb.RegisterResourcesServer(grpcServer, srv)

	if config.Read().GRPCGatewayEnabled {
		go func() {
			log.Infof("grpc starting and listening on %s", listen)
			<-serveGRPCGateway()
		}()
	}

	go func() {
		s := <-srv.Close
		logrus.Infof("got signal %v, attempting graceful shutdown", s)

		grpcServer.GracefulStop()

	}()

	go func() {
		err = grpcServer.Serve(lis)
		if err != nil {
			panic(err)
		}
	}()

	return make(chan bool)
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
