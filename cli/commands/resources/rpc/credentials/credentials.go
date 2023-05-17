package credentials

import (
	"fmt"

	"github.com/edobtc/cloudkit/bitcoind/rpc"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	username string
	password string
)

func init() {
	Cmd.Flags().StringVarP(&username, "username", "u", "", "username")
	Cmd.Flags().StringVarP(&password, "password", "p", "", "password")
	Cmd.MarkFlagRequired("username")
}

// Cmd is the top level relay
// command to export
var Cmd = &cobra.Command{
	Use:   "credentials",
	Short: "credentials",
	Run: func(cmd *cobra.Command, args []string) {
		c := rpc.Credentials{
			Username: username,
			Password: password,
		}

		data, err := c.Generate()
		if err != nil {
			log.Error(err)
			return
		}

		fmt.Println("add this to your bitcoin.conf file:")
		fmt.Println(data)
		fmt.Println("")
		fmt.Println("your password is:", c.Password)
	},
}
