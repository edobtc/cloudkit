package status

import (
	"github.com/edobtc/cloudkit/clients/bitcoind/http"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Cmd is the top level config
// command to export
var Cmd = &cobra.Command{
	Use:   "status",
	Short: "status",
	Run: func(cmd *cobra.Command, args []string) {
		err := blockCount()
		if err != nil {
			log.Error(err)
		}
		cmd.Help()
	},
}

func init() {
	// Cmd.AddCommand(show.Cmd)
}

func blockCount() error {
	client, err := http.NewBasicRpcClient()
	if err != nil {
		panic(err)
	}
	defer client.Shutdown()

	// Query the RPC server for the current block count and display it.
	blockCount, err := client.GetBlockCount()
	if err != nil {
		return err
	}

	log.Printf("Block count: %d", blockCount)
	return nil
}
