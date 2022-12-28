package dispatcher

import (
	"sync"

	"github.com/edobtc/cloudkit/runtime/agent"
)

// Dispatcher handles the processing of submitted
// experiment executor agents
type Dispatcher struct {
	sync  *sync.Mutex
	Log   chan []byte
	Queue []*agent.Agent
}

// NewDispatcher returns a new Dispatcher with defaults
func NewDispatcher() *Dispatcher {
	return &Dispatcher{
		sync:  &sync.Mutex{},
		Log:   make(chan []byte),
		Queue: []*agent.Agent{},
	}
}

// Add adds an agent to the dispatcher queue
func (d *Dispatcher) Add(a *agent.Agent) int {
	d.sync.Lock()
	d.Queue = append(d.Queue, a)
	index := len(d.Queue)
	d.sync.Unlock()

	d.Log <- []byte([]byte(a.ID.String()))
	return index
}

// Remove removes an agent by id from the dispatcher queue
func (d *Dispatcher) Remove(id string) {
	d.sync.Lock()
	for i, a := range d.Queue {
		if a.ID.String() == id {
			if len(d.Queue) == 1 {
				d.Queue = []*agent.Agent{}
			} else {
				d.Queue = append(d.Queue[:i], d.Queue[i+1:]...)
			}
		}
	}

	d.sync.Unlock()
}
