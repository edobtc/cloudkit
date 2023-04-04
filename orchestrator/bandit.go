package orchestrator

import (
	"context"

	"github.com/edobtc/cloudkit/plan"
)

// TODO:
// NOT IMPLEMENTED
// Bandit s a standard Bandit test ie: a/b, a/b w/ control, a/b/c
type Bandit struct {
	plan *plan.Definition

	status   chan Status
	finished chan bool
}

// NewBanditExperiment returns an initialized Bandit struct
// with proper values configure
func NewBanditExperiment(p *plan.Definition) *Bandit {
	return &Bandit{
		plan:     p,
		status:   make(chan Status),
		finished: make(chan bool),
	}
}

// Run runs the entire process
func (b *Bandit) Run(ctx context.Context) { return }

// Bootstrap begins setting up the experiment
func (b *Bandit) Bootstrap() error { return nil }

// Start begins the experiment
func (b *Bandit) Start(ctx context.Context) error { return nil }

// Teardown will safely stop and de-provision the experiment abruptly
func (b *Bandit) Teardown() error { return nil }

// Abort will stop the experiment abruptly
func (b *Bandit) Abort() error { return nil }

// Finished will signal an event and close the channel when the experiment
// is finished
func (b *Bandit) Finished() chan bool { return make(chan bool) }

// Watch will allow you to await status updates of any orchestrated experiment
func (b *Bandit) Watch() chan Status { return make(chan Status) }
