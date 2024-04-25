package zmq

import (
	"encoding/hex"
	"fmt"
	"os"
	"sync"

	"github.com/edobtc/cloudkit/config"
	delivery "github.com/edobtc/cloudkit/events/publishers"

	"github.com/btcsuite/btcutil"
	log "github.com/sirupsen/logrus"
	"github.com/zeromq/goczmq"
)

type Subscriber struct {
	sync.Mutex

	sock *goczmq.Channeler

	destinations []delivery.Publisher

	listener chan interface{}
	Close    chan os.Signal
	Wait     chan bool
	Error    chan error
}

func NewSubscriber() *Subscriber {
	host := config.Read().Streams.ZeroMqListenAddr

	log.Debug(host)

	return &Subscriber{
		sock: goczmq.NewSubChanneler(host),

		listener: make(chan interface{}),
		Close:    make(chan os.Signal),
		Wait:     make(chan bool),
		Error:    make(chan error),
	}
}

func (s *Subscriber) Sock() *goczmq.Channeler {
	return s.sock
}

func (s *Subscriber) Listen() <-chan interface{} {
	return s.listener
}

func (s *Subscriber) AddDestination(d delivery.Publisher) int {
	s.destinations = append(s.destinations, d)
	return len(s.destinations)
}

func (s *Subscriber) RemoveDestination(d delivery.Publisher) int {
	s.Lock()

	for idx, destination := range s.destinations {
		if destination == d {
			s.destinations = append(s.destinations[:idx], s.destinations[idx+1:]...)
		}
	}
	s.Unlock()
	return len(s.destinations)
}

func (s *Subscriber) Detach() error {
	s.sock.Destroy()

	// close channels
	close(s.Error)
	close(s.Wait)
	close(s.Close)

	return nil
}

func (s *Subscriber) Subscribe(channel string) {
	s.sock.Subscribe(channel)
}

func (s *Subscriber) BatchSubscribe(channels []string) {
	for _, channel := range channels {
		s.sock.Subscribe(channel)
	}
}

func (s *Subscriber) Start() chan bool {
	go func() {
		for {
			request := <-s.sock.RecvChan
			switch string(request[0]) {
			case "hashtx":
				fmt.Println(string(request[0]))
				fmt.Println(hex.EncodeToString(request[1]))
				tx := hex.EncodeToString(request[1])
				for _, d := range s.destinations {

					// :-/
					go func(p delivery.Publisher) {
						p.Send([]byte(tx))
					}(d)
				}

			case "hashblock":
				fmt.Println(string(request[0]))
				fmt.Println(hex.EncodeToString(request[1]))
			case "rawtx":
				fmt.Println(string(request[0]))
				txString := hex.EncodeToString(request[1])
				tx, err := btcutil.NewTxFromBytes(request[1])
				if err != nil {
					fmt.Println(err)
				}
				fmt.Println(txString)
				fmt.Println(tx.Index())
			}
		}
	}()

	return s.Wait
}
