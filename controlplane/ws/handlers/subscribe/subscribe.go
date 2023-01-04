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

	mt, _, err := c.ReadMessage()
	if err != nil {
		log.Error(err)
		return
	}

	pool := r.Context().Value(ws.PoolContextKey).(*ws.ConnectionPool)

	pool.Add(c)

	err = c.WriteMessage(mt, []byte(`
	  ____  |  |    ____   __ __   __| _/|  | __|__|_/  |_  /\    ______ __ __ \_ |__    ______  ____  _______ |__|\_ |__    ____    __| _/
	_/ ___\ |  |   /  _ \ |  |  \ / __ | |  |/ /|  |\   __\ \/   /  ___/|  |  \ | __ \  /  ___/_/ ___\ \_  __ \|  | | __ \ _/ __ \  / __ |
	\  \___ |  |__(  <_> )|  |  // /_/ | |    < |  | |  |   /\   \___ \ |  |  / | \_\ \ \___ \ \  \___  |  | \/|  | | \_\ \\  ___/ / /_/ |
	 \___  >|____/ \____/ |____/ \____ | |__|_ \|__| |__|   \/  /____  >|____/  |___  //____  > \___  > |__|   |__| |___  / \___  >\____ |
		 \/                           \/      \/                     \/             \/      \/      \/                  \/      \/      \/ `))

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
