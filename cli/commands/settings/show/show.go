package show

import (
	"github.com/edobtc/cloudkit/settings"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

// Cmd is the command for the cli tool
var Cmd = &cobra.Command{
	Use:   "show",
	Short: "show",
	Run: func(cmd *cobra.Command, args []string) {
		s, err := settings.Read()
		if err != nil {
			log.Error(err)
			return
		}

		data, err := s.ToYAML()
		if err != nil {
			log.Error(err)
			return
		}

		log.Info(string(data))
	},
}
