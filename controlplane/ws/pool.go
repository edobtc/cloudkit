package ws

import (
	"context"
	"net/http"

	"github.com/gorilla/websocket"
)

type contextKey string

const (
	PoolContextKey contextKey = "pool"
)

var (
	Pool = NewConnectionPool()
)

type ConnectionPool struct {
	connections map[*websocket.Conn]bool
}

func NewConnectionPool() *ConnectionPool {
	return &ConnectionPool{
		connections: map[*websocket.Conn]bool{},
	}
}

func (p *ConnectionPool) Add(c *websocket.Conn) bool {
	p.connections[c] = true
	return true
}

// Publish sends a message to all connections in the pool
//
// eg:
//
//	ws.Pool.Publish([]byte("wu tang killa-beez, we on a swarm"))
func (p *ConnectionPool) Publish(msg []byte) {
	for c := range p.connections {

		go func(client *websocket.Conn) {
			client.WriteMessage(websocket.TextMessage, msg)
		}(c)
	}
}

func (p *ConnectionPool) Remove(c *websocket.Conn) bool {
	delete(p.connections, c)
	return true
}

func InjectConnectionPool(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := context.WithValue(r.Context(), PoolContextKey, Pool)
		next.ServeHTTP(w, r.WithContext(ctx))
	})

}
