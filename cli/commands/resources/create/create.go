package create

import (
	"context"
	"os"

	"github.com/edobtc/cloudkit/client"
	"github.com/edobtc/cloudkit/plan"

	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	name      string
	version   string
	config    string
	transport string
)

// Cmd is the top level relay
// command to export
var Cmd = &cobra.Command{
	Use:   "create",
	Short: "create",
	Run: func(cmd *cobra.Command, args []string) {
		if name == "" {
			log.Info("you must provide a name")
			return
		}

		if transport == "grpc" {
			createGRPC()
		} else {
			createHTTP()
		}

	},
}

func createGRPC() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	conn, err := client.NewInsecureGrpcConn(ctx)

	if err != nil {
		log.Infof("could not connect to server: %v", err)
		return
	}
	defer conn.Close()

	client := pb.NewResourcesClient(conn)
	resp, err := client.Create(context.TODO(), &pb.CreateRequest{
		Config: &pb.ResourceConfiguration{
			Name: name,
		},
	})

	if err != nil {
		log.Errorf("could not create streaming client: %v", err)
		return
	}

	log.Info(resp)
}

func createHTTP() {
	c, err := client.NewClientWithToken()
	if err != nil {
		log.Error(err)
		return
	}

	p := plan.Definition{}

	_, err = c.CreateResource(p)

	if err != nil {
		log.Error(err)
		os.Exit(1)
	}

	log.Info("resource created")
}

func init() {
	Cmd.Flags().StringVarP(&name, "name", "n", "", "resource name")
	Cmd.Flags().StringVarP(&version, "version", "v", "lnd-15.5", "version")
	Cmd.Flags().StringVarP(&config, "config", "c", "", "config file")
	Cmd.Flags().StringVarP(&transport, "transports", "t", "grpc", "(grpc|http)")
}
