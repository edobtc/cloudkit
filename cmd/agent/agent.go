package agent

import (
	"fmt"
	"os"

	"github.com/edobtc/cloudkit/logging"

	// relay route setup

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

func init() {
	logging.Setup()
}

var rootCmd = &cobra.Command{
	Use:   "cloudkit-agent",
	Short: "cloudkit-agent is a small utility that should run on new nodes during provisioning",

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("running the agent")
	},
}

func main() {
	if err := rootCmd.Execute(); err != nil {
		log.Error(err)
		os.Exit(1)
	}
}
