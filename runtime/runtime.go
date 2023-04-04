package runtime

import (
	"context"
	"time"

	"github.com/edobtc/cloudkit/orchestrator"
	"github.com/edobtc/cloudkit/runtime/agent"
	"github.com/edobtc/cloudkit/runtime/dispatcher"

	log "github.com/sirupsen/logrus"
)

const (
	DefaultExec = 10 * time.Second
)

// Runtime is the standlone control plane orchestration
// runtime for receiving and coordinating submitted experiment
// plans
type Runtime struct {
	Orchestrator orchestrator.Orchestrator

	ctx        context.Context
	tick       *time.Ticker
	err        <-chan error
	done       chan bool
	Dispatcher *dispatcher.Dispatcher
}

type Config struct {
	Interval time.Duration
	ErrChan  <-chan error
}

// New creates a new Runtime with defaults
func New(ctx context.Context) *Runtime {
	return &Runtime{
		ctx:        ctx,
		tick:       time.NewTicker(DefaultExec),
		err:        make(<-chan error),
		done:       make(chan bool),
		Dispatcher: dispatcher.NewDispatcher(),
	}
}

// NewWithConfiguration creates a new Runtime with optional
// configuration settings
func NewWithConfiguration(ctx context.Context, cfg Config) *Runtime {
	if cfg.Interval == 0 {
		cfg.Interval = DefaultExec
	}

	if cfg.ErrChan == nil {
		cfg.ErrChan = make(<-chan error)
	}

	return &Runtime{
		ctx:        ctx,
		tick:       time.NewTicker(cfg.Interval),
		err:        cfg.ErrChan,
		done:       make(chan bool),
		Dispatcher: dispatcher.NewDispatcher(),
	}
}

// Run begins the runtime mechanism
func (r *Runtime) Run() {

	// watch for runtime events that are logged
	// out from the dispatcher
	go func() {
		for msg := range <-r.Dispatcher.Log {
			log.Info(msg)
		}
	}()

	go func() {
		log.Info("run orchestrator")
		// r.Orchestrator.Run(r.ctx)
	}()

	go func() {
		r.Start()
	}()

	for {
		select {
		case <-r.done:
		case <-r.ctx.Done():
			log.Info("orchestration done")
			// r.Orchestrator.Teardown()
			close(r.done)
		}
	}
}

func (r *Runtime) Await() chan bool {
	return r.done
}

// Start starts the runtime reconciliation process
// to fix state on submitted provision or management
// events
//
// though it can be set up and called directly, it is managed
// by the runtime.Run command
func (r *Runtime) Start() {
	for range r.tick.C {
		log.Debug("running reconciliation")
		for _, a := range r.Dispatcher.Queue {
			go func(ap *agent.Agent) {
				status, err := ap.Reconcile()
				log.Infof("reconcile status %v", status)
				if err != nil {
					log.Error(err)
				}
			}(a)
		}
	}
}
