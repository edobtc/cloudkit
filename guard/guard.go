package guard

import (
	"time"

	"github.com/edobtc/cloudkit/guard/adapters"

	log "github.com/sirupsen/logrus"
)

type (
	// Kind represents the type of guard
	// to use and how the data as input to the
	// guard is sourced
	Kind int
	// Status is the guard's suggestion of a
	// state of the input data
	Status int

	// Action is the response to take to a guard being
	// triggered
	Action int
)

const (
	// HTTP will make an HTTP request
	// to some data source to test a guard
	HTTP Kind = iota

	// Query Means query a data source
	Query

	// OK means the guard was not triggered
	OK Status = iota

	// Failing status can be used when any status in the window
	// has failed
	Failing

	// Failed means the guard was triggered
	Failed

	// Warn will send a warning
	Warn Action = iota

	// Alert will trigger an alert
	Alert

	// Abort will unsafely and suddenly end the guarded operation
	Abort

	// Teardown will safely begin teardown for any guarded operation
	Teardown

	// DefaultPollInterval is the default checking interval
	// for a guide
	DefaultPollInterval = time.Second * 1

	// DefaultIntegration is the default integration to use
	DefaultIntegration = "cloudwatch"

	// DefaultKind is the default guard type to use
	DefaultKind = Query

	// DefaultWindowSize is the default windowSize to keep of
	// historical statuses
	DefaultWindowSize = 10
)

// Guard is a mechanism for watching some data and using that
// as signal to see if
type Guard struct {
	Kind               Kind          `json:"kind"`
	Name               string        `json:"name"`
	Status             Status        `json:"status"`
	StatusDistribution []Status      `json:"statusDist"`
	Count              int           `json:"count"`
	Threshold          int           `json:"threshold"`
	Interval           time.Duration `json:"interval"`
	WindowSize         int           `json:"windowSize"`
	Tolerance          float32       `json:"tolerance"`
	Spec               interface{}   `json:"spec"`
	Integration        string        `json:"integration"`
	done               chan bool
	notifier           chan Status
	tick               *time.Ticker

	Err error

	adapter adapters.Adapter
}

// NewGuard returns an initialized Guard
func NewGuard() *Guard {
	return &Guard{
		Interval:    DefaultPollInterval,
		Kind:        Query,
		Integration: "cloudwatch",
		notifier:    make(chan Status),
		WindowSize:  DefaultWindowSize,
	}
}

// Start begins continuously running a guard
func (g *Guard) Start() chan Status {
	ticker := time.NewTicker(DefaultPollInterval)
	g.tick = ticker
	g.done = make(chan bool)

	a, err := adapters.Load(g.Integration, g.Spec)

	if err != nil {
		log.Error(err)
		g.Stop()
	}

	g.adapter = a

	go func() {
		for range ticker.C {
			success, err := g.adapter.Check()

			g.Count++

			if err != nil {
				log.Error(err)
				g.Err = err
				g.Stop()
				continue
			}

			if success {
				g.addStatus(OK)
			}

			if !success {
				g.addStatus(Failed)
			}

			g.Status = g.Evaluate()

			if g.Failed() {
				g.notifier <- g.Status
				g.Stop()
			}
		}
	}()

	go func() {
		select {
		case <-g.done:
			g.notifier <- g.Status
			close(g.done)
			close(g.notifier)
		}
	}()

	return g.notifier
}

// Stop stops a running guard
func (g *Guard) Stop() error {
	g.tick.Stop()
	g.done <- true
	return nil
}

func (g *Guard) addStatus(s Status) {
	g.StatusDistribution = append(
		g.StatusDistribution,
		s,
	)

	if len(g.StatusDistribution) > g.WindowSize {
		g.StatusDistribution = append(
			g.StatusDistribution[:1],
			g.StatusDistribution[1+1:]...,
		)
	}
}

// Evaluate uses the distribution of previous checks to see
// if we cross a defined threshold for the guard. Needs improving
func (g *Guard) Evaluate() Status {
	var failed int
	for _, status := range g.StatusDistribution {
		if status != OK {
			failed++
		}
	}

	if failed >= 0 {
		if failed >= g.Threshold {
			g.Status = Failed
		} else {
			g.Status = Failing
		}
	}

	g.Status = OK

	return g.Status
}

// Failed returns true if we've determined the guard to be failing
func (g *Guard) Failed() bool {
	return g.Status == Failed
}
