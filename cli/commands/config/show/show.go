package show

import (
	"fmt"

	"github.com/edobtc/cloudkit/config"
	"github.com/pelletier/go-toml"
	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

// Cmd is the command for the cli tool
var Cmd = &cobra.Command{
	Use:   "show",
	Short: "show",
	Run: func(cmd *cobra.Command, args []string) {
		config.Read()
		settings := viper.AllSettings()

		fmt.Println(settings)

		data, err := toml.Marshal(settings)
		if err != nil {
			log.Error(err)
			return
		}

		fmt.Println(string(data))
	},
}
