package agent

import (
	"github.com/edobtc/cloudkit/guard"
	"github.com/edobtc/cloudkit/parameters"
	"github.com/edobtc/cloudkit/plan"
	"github.com/edobtc/cloudkit/probe"
	"github.com/edobtc/cloudkit/resources/providers"
	"github.com/edobtc/cloudkit/resources/providers/autoload"
	"github.com/edobtc/cloudkit/target"

	"github.com/google/uuid"

	"github.com/pkg/errors"
	log "github.com/sirupsen/logrus"
)

// Status is the status of a agent that is completing work
type Status int

var (
	// ErrAborted means any work was abandoned
	ErrAborted = errors.New("aborted")

	// ErrStatusUnknown we are unsure of any status of a provisioner
	ErrStatusUnknown = errors.New("no known status known")

	// ErrFailed means the agent has failed to do it's work
	ErrFailed = errors.New("failed")
)

const (
	// Submitted signals that work has been received
	Submitted Status = iota

	// Provisioning signals that things are being set up
	Provisioning

	// Running signals the work is being undertaken
	Running

	// Aborted signals the work was abandoned
	Aborted

	// Failed signals the attempted work failed to complete with an error
	Failed

	// Complete signals any work has successfully reached the end of it's lifecycle
	Complete
)

// Agent is a worker that accepts a declaration
// and provisioner and completes the work to
// manage state and resolve it.
// for example, provisioning a new node
type Agent struct {
	ID uuid.UUID

	Provider     providers.Provider
	ProviderName string
	Target       target.Target
	Parameters   parameters.Parameters
	Guards       []*guard.Guard
	Probes       []*probe.Probe

	err    error
	status Status

	Next *Agent
}

// NewAgent instantiates an Agent with defaults
func NewAgent(prov providers.Provider) *Agent {
	return &Agent{
		ID:       uuid.New(),
		Provider: prov,
		status:   Submitted,
	}
}

// Parse takes an plan and constructs a workflow
// of agents as a sequence to be executed
func Parse(p plan.Definition) ([]*Agent, error) {
	workflow := []*Agent{}

	c := providers.GenericConfig{}
	c["target"] = p.Spec.Target

	prov, err := autoload.Load(p.Spec.Provider, c)
	if err != nil {
		return workflow, err
	}

	a := Agent{
		Target:       p.Spec.Target,
		ProviderName: p.Spec.Provider,
		Provider:     prov,
		Parameters:   p.Spec.Parameters,
		Guards:       p.Spec.Guards,
		Probes:       p.Spec.Probes,
	}

	workflow = append(workflow, &a)

	// OLD Variant application
	//
	// for _, v := range p.Spec.Variants {
	// 	if v.Target.Empty() {
	// 		v.Config["target"] = p.Spec.Target
	// 	} else {
	// 		v.Config["target"] = v.Target
	// 	}

	// 	prov, err := autoload.Load(p.Spec.Provider, v.Config)

	// 	if err != nil {
	// 		return workflow, err
	// 	}

	// 	av := Agent{
	// 		ProviderName: p.Spec.Provider,
	// 		Provider:     prov,
	// 		Parameters:   p.Spec.Parameters,
	// 		Guards:       v.Guards,
	// 		Probes:       v.Probes,
	// 	}

	// 	if v.Target.Empty() {
	// 		av.Target = p.Spec.Target
	// 	} else {
	// 		av.Target = v.Target
	// 	}

	// 	workflow = append(workflow, &av)
	// }

	return workflow, nil
}

// Reconcile checks the current status of a process being managed
// and ensures the state is resolved
func (a *Agent) Reconcile() (Status, error) {
	switch a.status {
	case Submitted:
		log.Info("State is submitted, about to apply")
		a.status = Provisioning
		err := a.Provider.Apply()
		if err != nil {
			a.err = err
			a.status = Failed
			return a.status, a.err
		}
		a.status = Running
	case Provisioning:
		_, err := a.Provider.ProbeReadiness()
		if err != nil {
			a.err = err
			a.status = Failed
			return a.status, a.err
		}
	case Running:
	case Aborted:
		return Aborted, ErrAborted
	case Failed:
		return Failed, errors.Wrap(a.err, "failed")
	case Complete:
		return Complete, nil
	}

	return a.status, nil
}

// State returns the state of a process managed
// by an Agent
func (a *Agent) State() Status {
	return a.status
}
