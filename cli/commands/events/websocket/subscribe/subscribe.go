package subscribe

import (
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"time"

	"github.com/edobtc/cloudkit/cli/commands/events/subscribe/debug"

	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	host string
)

func init() {
	Cmd.AddCommand(debug.Cmd)

	Cmd.Flags().StringVarP(&host, "host", "n", "127.0.0.1:8081", "host")
}

// Cmd is the command for the cli tool
var Cmd = &cobra.Command{
	Use:   "subscribe",
	Short: "subscribe",
	Run: func(cmd *cobra.Command, args []string) {
		interrupt := make(chan os.Signal, 1)
		signal.Notify(interrupt, os.Interrupt)

		u := url.URL{Scheme: "ws", Host: host, Path: "/ws/subscribe"}

		c, _, err := websocket.DefaultDialer.Dial(u.String(), nil)
		if err != nil {
			logrus.Error("dial:", err)
			return
		}
		defer c.Close()

		done := make(chan struct{})

		go func() {
			defer close(done)
			for {
				_, message, err := c.ReadMessage()
				if err != nil {
					logrus.Info("read:", err)
					return
				}
				fmt.Println(string(message))
			}
		}()

		ticker := time.NewTicker(time.Second)
		defer ticker.Stop()

		for {
			select {
			case <-done:
				return
			case t := <-ticker.C:
				err := c.WriteMessage(websocket.TextMessage, []byte(t.String()))
				if err != nil {
					logrus.Info("write:", err)
					return
				}
			case <-interrupt:
				logrus.Debug("interrupt")

				// Cleanly close the connection by sending a close message and then
				// waiting (with timeout) for the server to close the connection.
				err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
				if err != nil {
					logrus.Println("write close:", err)
					return
				}
				select {
				case <-done:
				case <-time.After(time.Second):
				}
				return
			}
		}
	},
}
