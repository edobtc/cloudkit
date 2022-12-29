package list

import (
	"github.com/edobtc/cloudkit/keys/ssh"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	name string
)

var Cmd = &cobra.Command{
	Use:   "list",
	Short: "list",
	Run: func(cmd *cobra.Command, args []string) {
		s := ssh.NewSSHKey(ssh.DefaultOptions())

		err := s.Generate()
		if err != nil {
			logrus.Error(err)
			return
		}

		err = s.SaveToFile()
		if err != nil {
			logrus.Error(err)
			return
		}
	},
}

func init() {
	Cmd.Flags().StringVarP(&name, "name", "n", "", "key name")
}
