package subscribe

import (
	"net/http"

	"github.com/edobtc/cloudkit/controlplane/ws"
	"github.com/gorilla/websocket"
	log "github.com/sirupsen/logrus"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func SubscribeCommand(w http.ResponseWriter, r *http.Request) {
	c, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Error(err)
		return
	}

	defer c.Close()

	mt, message, err := c.ReadMessage()
	if err != nil {
		log.Error(err)
		return
	}

	log.Info(string(message))

	pool := r.Context().Value(ws.PoolContextKey).(*ws.ConnectionPool)

	pool.Add(c)

	err = c.WriteMessage(mt, []byte("subscribed"))

	if err != nil {
		log.Error(err)
		return
	}

	for {
		if _, _, err := c.NextReader(); err != nil {
			// handle error form a disconnection
			c.Close()
			break
		}
	}
}
