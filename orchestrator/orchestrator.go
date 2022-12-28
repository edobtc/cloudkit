package orchestrator

import (
	"context"
	"errors"

	"github.com/edobtc/cloudkit/plan"
)

// Status is the state of the experiment in the
// experiment lifecycle as observed by the orchestrator
type Status int

const (
	// Starting signals that the experiment is starting
	Starting Status = iota

	// Ready signals the experiment is ready to begin
	Ready

	// Waiting signals the experiment is ready, and waiting to begin
	Waiting

	// Active signals the experiment is running
	Active

	// Aborted signals the experiment has been abruptly aborted
	Aborted

	// Stopped signals the experiment has been prematurely stopped
	Stopped

	// Finishing signals the experiment is in a state of being torn down
	Finishing

	// Finished signals the experiment has finished
	Finished
)

var (
	// ErrOrchestratorDoesNotExist is returned is an orchestrator is loaded
	// by name but not recognized or implemented
	ErrOrchestratorDoesNotExist = errors.New("named orchestrator doesn't exist")
)

// Orchestrator is the interface which all implemented
// experiment orchestrators must conform to in order to be
// valid
type Orchestrator interface {
	Run(context.Context)
	Bootstrap() error
	Start(context.Context) error
	Abort() error
	Teardown() error

	Watch() chan Status
	Finished() chan bool
}

// Load will load an orchestrator by name
func Load(kind string, p *plan.Definition) (Orchestrator, error) {
	switch kind {
	case "split":
		return NewSplitExperiment(p), nil
	case "adaptive":
		return NewAdaptiveExperiment(p), nil
	case "bandit":
		return NewBanditExperiment(p), nil
	}

	return nil, ErrOrchestratorDoesNotExist
}
