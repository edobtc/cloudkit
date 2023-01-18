package publish

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/edobtc/cloudkit/config"
	"github.com/edobtc/cloudkit/httpclient"
	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"
	"github.com/google/uuid"

	log "github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	payload string
)

func init() {
	// eg: -p "{'huh':'wut'}"
	Cmd.Flags().StringVarP(&payload, "payload", "p", "", "the payload to broadcast, leave empty to generate a default")
}

// Cmd is the command for the cli tool
var Cmd = &cobra.Command{
	Use:   "publish",
	Short: "publish a message to all connected websocket subscribers",
	Run: func(cmd *cobra.Command, args []string) {
		if payload == "" {
			log.Info("using a default event")
			payload = defaultEvent()
		}

		err := makeRequest(payload)
		if err != nil {
			log.Error(err)
			return
		}
		log.Info("published")
	},
}

func defaultEvent() string {
	resp := pb.ResourceResponse{
		Identifier: uuid.New().String(),
		Name:       "test event",
	}

	data, err := json.Marshal(&resp)
	if err != nil {
		log.Error(err)
		return ""
	}

	return string(data)
}

func makeRequest(data string) error {
	u := fmt.Sprintf("http://%s%s", config.Read().Listen, "/v1/events/ws/publish")
	log.Info(u)

	req, err := http.NewRequest("POST", u, bytes.NewBuffer([]byte(data)))
	if err != nil {
		return err
	}

	req.Header.Set("accept", "application/json")

	http := httpclient.New()

	resp, err := http.Do(req)

	if err != nil {
		return err
	}

	if resp.StatusCode != 200 {
		log.Error("failed to publish event")
		return errors.New("non 200 response code")
	}

	return nil
}
