package set

import (
	"fmt"
	"net/url"

	"github.com/edobtc/cloudkit/settings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	u string
)

func init() {
	Cmd.Flags().StringVarP(&u, "url", "u", "", "the URL to configure")
	Cmd.MarkFlagRequired("url")
}

// Cmd is the command for the cli tool
var Cmd = &cobra.Command{
	Use:   "set-url",
	Short: "set-url",
	Run: func(cmd *cobra.Command, args []string) {
		parsed, err := url.ParseRequestURI(u)
		if err != nil {
			log.Error("failed to parse submitted url as valid")
			log.Error(err)
			return
		}

		log.Info(fmt.Sprintf("setting url to %s", parsed.String()))

		s, err := settings.Read()

		s.Data.URL = parsed.String()

		if err != nil {
			log.Error(err)
			return
		}

		err = settings.Write()
		if err != nil {
			log.Error(err)
			return
		}
	},
}
