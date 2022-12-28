package about

import (
	"fmt"

	"github.com/spf13/cobra"
)

var (
	data = `
 _    _              _              _  _    _  _
| |_ | |_  ___  ___ | | ___  _ _  _| || |_ |_|| |_
| . ||  _||  _||  _|| || . || | || . || '_|| ||  _|
|___||_|  |___||___||_||___||___||___||_,_||_||_|

Deploy and Manage The Timechain, by @edobtc
   `
)

var Cmd = &cobra.Command{
	Use: "about",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(data)
	},
}
