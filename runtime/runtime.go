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
	execTick = 10 * time.Second
)

// Runtime is the standlone control plane orchestration
// runtime for receiving and coordinating submitted experiment
// plans
type Runtime struct {
	Orchestrator orchestrator.Orchestrator

	ctx        context.Context
	tick       <-chan time.Time
	Dispatcher *dispatcher.Dispatcher
}

// NewRuntime creates a new Runtime with defaults
func NewRuntime(ctx context.Context) *Runtime {
	return &Runtime{
		ctx:        ctx,
		tick:       time.Tick(execTick),
		Dispatcher: dispatcher.NewDispatcher(),
	}
}

// Run begins the runtime mechanism
func (r *Runtime) Run() {
	done := make(chan bool)

	// watch for runtime events that are logged
	// out from the dispatcher
	go func() {
		for msg := range <-r.Dispatcher.Log {
			log.Info(msg)
		}
	}()

	go func() {
		log.Info("run orchestrator")
		r.Orchestrator.Run(r.ctx)
	}()

	go func() {
		r.Start()
	}()

	for {
		select {
		case <-done:
		case <-r.ctx.Done():
			log.Info("orchestration done")
			r.Orchestrator.Teardown()
			close(done)
		}
	}
}

// Start starts the runtime reconciliation process
// to fix state on submitted/orchestrated experiments
func (r *Runtime) Start() {
	for range r.tick {
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
