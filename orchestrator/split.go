package orchestrator

import (
	"context"
	"errors"

	"github.com/edobtc/cloudkit/guard"
	"github.com/edobtc/cloudkit/plan"
	"github.com/edobtc/cloudkit/probe"
	"github.com/edobtc/cloudkit/runtime/agent"

	log "github.com/sirupsen/logrus"
)

// Split is a standard split test ie: a/b, a/b w/ control, a/b/c
type Split struct {
	plan *plan.Definition

	Running Running

	errors   chan error
	probes   chan error
	guards   chan error
	status   chan Status
	finished chan bool
}

// NewSplitExperiment returns an initialized Split struct
// with proper values configured
func NewSplitExperiment(p *plan.Definition) *Split {
	return &Split{
		Running: Running{},

		plan: p,

		errors:   make(chan error),
		guards:   make(chan error),
		probes:   make(chan error),
		status:   make(chan Status),
		finished: make(chan bool),
	}
}

// Running is a tracker for running elements of the
// experiment engine
type Running struct {
	Guards []*guard.Guard
	Probes []*probe.Probe
}

// Run will begin orchestration of the experiment
func (s *Split) Run(ctx context.Context) {
	if err := s.Bootstrap(); err != nil {
		log.Error(err)
		s.errors <- err
		return
	}

	if err := s.Start(ctx); err != nil {
		log.Error(err)
		s.errors <- err
		return
	}

	for {
		for status := range s.status {
			switch status {
			case Aborted:
				s.errors <- s.Abort()
			case Stopped:
				s.errors <- s.Teardown()
			case Finished:
				s.errors <- s.Teardown()
				s.Finished()
			}
		}
	}
}

// Bootstrap begins setting up the experiment
func (s *Split) Bootstrap() error {
	w, err := agent.Parse(*s.plan)
	if err != nil {
		return err
	}

	for _, w := range w {
		selection, err := w.Provider.Select()
		if err != nil {
			log.Error(err)
			return err
		}

		err = w.Provider.Apply()
		if err != nil {
			return err
		}

		log.Info(selection)

		// for _, fault := range s.plan.Spec.Faults {
		// 	if !fault.Spec.Targets.Empty() {
		// 		selection, err = w.Provider.Select()
		// 		if err != nil {
		// 			log.Error(err)
		// 			return err
		// 		}
		// 	}

		// 	resources := selection.Select(w.Parameters)
		// 	fault.Apply(resources)
		// }

		for _, g := range w.Guards {
			s.Running.Guards = append(s.Running.Guards, g)
		}

		for _, p := range w.Probes {
			s.Running.Probes = append(s.Running.Probes, p)
		}
	}

	return nil
}

// Start begins the experiment
func (s *Split) Start(ctx context.Context) error {
	//
	// Currently the only "start" operation for the split test
	// is to begin probes and guards.
	//
	// When configuration of routing is possible, we can enable the traffic
	// split in this start call also
	//
	for _, g := range s.Running.Guards {
		go func(gd *guard.Guard, errNotifier chan error) {
			status := gd.Start()

			select {
			case s := <-status:
				if s == guard.Failed {
					errNotifier <- errors.New("guard failed")
				}
			case <-ctx.Done():
				errNotifier <- gd.Stop()
			}
		}(g, s.guards)
	}

	for _, p := range s.Running.Probes {
		go func(pb *probe.Probe, errNotifier chan error) {
			status := pb.Start()

			select {
			case <-status:
				errNotifier <- errors.New("probe failed")
			case <-ctx.Done():
				errNotifier <- pb.Stop()
			}
		}(p, s.probes)
	}

	return nil
}

// Teardown will safely stop and de-provision the experiment abruptly
func (s *Split) Teardown() error { return nil }

// Abort will stop the experiment abruptly
func (s *Split) Abort() error { return nil }

// Finished will signal an event and close the channel when the experiment
// is finished
func (s *Split) Finished() chan bool { return make(chan bool) }

// Watch will allow you to await status updates of any orchestrated experiment
func (s *Split) Watch() chan Status { return make(chan Status) }
