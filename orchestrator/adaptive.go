package orchestrator

import (
	"context"

	"github.com/edobtc/cloudkit/plan"
)

// TODO:
// NOT IMPLEMENTED
//

// Adaptive s a standard Adaptive test ie: a/b, a/b w/ control, a/b/c
type Adaptive struct {
	plan *plan.Definition

	status   chan Status
	finished chan bool
}

// NewAdaptiveExperiment returns an initialized Adaptive struct
// with proper values configure
func NewAdaptiveExperiment(p *plan.Definition) *Adaptive {
	return &Adaptive{
		plan:     p,
		status:   make(chan Status),
		finished: make(chan bool),
	}
}

// Run runs the entire process
func (a *Adaptive) Run(ctx context.Context) { return }

// Bootstrap begins setting up the experiment
func (a *Adaptive) Bootstrap() error { return nil }

// Start begins setting up the experiment
func (a *Adaptive) Start(ctx context.Context) error { return nil }

// Teardown will safely stop and de-provision the experiment abruptly
func (a *Adaptive) Teardown() error { return nil }

// Abort will stop the experiment abruptly
func (a *Adaptive) Abort() error { return nil }

// Finished will signal an event and close the channel when the experiment
// is finished
func (a *Adaptive) Finished() chan bool { return make(chan bool) }

// Watch will allow you to await status updates of any orchestrated experiment
func (a *Adaptive) Watch() chan Status { return make(chan Status) }
