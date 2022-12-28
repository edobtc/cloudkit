package probe

import (
	"errors"
	"time"

	"github.com/edobtc/cloudkit/resources/providers"
	log "github.com/sirupsen/logrus"
)

type Kind int

const (
	// DefaultMaxFailures is the default limit when to trigger the probe
	DefaultMaxFailures = 5

	// DefaultPollInterval is the default interval at which to run the probe
	DefaultPollInterval = time.Second * 5

	// Http uses simple http checks to fulfil probe behavior
	HTTP Kind = iota

	// Listener implements a listener to some type of subscription channel
	// fulfil a passive probe
	Listener

	// Event awaits some trigger/event as a to fulfil a passive probe
	Event

	// Readiness will use the ProbeReadiness check in a provider
	Readiness

	// Liveness will use the ProbeLiveness check in a provider
	Liveness
)

var (
	// ErrMaxFailuresExceeded is returned when the probe is triggered
	ErrMaxFailuresExceeded = errors.New("max failures exceeded")

	// ErrNoProbeTypeSpecified if no probe type/kind is specific in the probe
	// config
	ErrNoProbeTypeSpecified = errors.New("no probe kind specified in probe.Kind")

	// ErrProbeTypeNotAvailable is returned if not probe of type is
	// available to load
	ErrProbeTypeNotAvailable = errors.New("probe kind not available")
)

// Config is a general config object for probes
type Config struct {
	Target string `yaml:"targer" json:"target"`
}

// Probe is a mechanism with which to check for some state, track said state, and
// trigger itself if a state is observed outside a tolerance
type Probe struct {
	Name        string        `yaml:"name" json:"name"`
	Kind        Kind          `yaml:"kind" json:"kind"`
	Interval    time.Duration `yaml:"duration" json:"duration"`
	MaxFailures int           `yaml:"maxFailures" json:"maxFailures"`
	Spec        Config        `yaml:"spec" json:"spec"`

	Performer adapter

	LastStatus bool `yaml:"lastStatus" json:"lastStatus"`
	Checks     int
	Failures   int `yaml:"failures" json:"failures"`
	provider   *providers.Provider

	wait chan bool
	err  chan error
	tick *time.Ticker
}

type adapter interface {
	Check() (bool, error)
}

// NewProbe returns an instantiated probe
func NewProbe() *Probe {
	return &Probe{
		MaxFailures: DefaultMaxFailures,
		Interval:    DefaultPollInterval,
		wait:        make(chan bool),
	}
}

// Await allows us to await a completion message from a probe
func (p *Probe) Await() chan bool {
	return p.wait
}

// Start begins continuously running a Probe
func (p *Probe) Start() chan bool {
	ticker := time.NewTicker(p.Interval)
	p.err = make(chan error)

	perf, err := load(p.Kind, p.Spec)
	if err != nil {
		p.err <- err
		return p.wait
	}

	p.Performer = perf

	go func() {
		for range ticker.C {
			success, err := p.Performer.Check()
			p.Checks++

			log.Infof("checked %d times", p.Checks)

			if err != nil {
				p.err <- err
			}

			if success == false {
				if p.LastStatus == false {
					p.Failures++
					if p.Failures >= p.MaxFailures {
						p.err <- ErrMaxFailuresExceeded
						p.wait <- true
						close(p.wait)
					}
				}
			} else {
				p.LastStatus = true
				p.Failures = 0
			}
		}
	}()

	go func() {
		for {
			select {
			case err := <-p.err:
				log.Error(err)
			}
		}
	}()

	return p.wait
}

// Stop stops a running Probe
func (p *Probe) Stop() error {
	p.wait <- true
	return nil
}

func load(k Kind, c Config) (adapter, error) {
	switch k {
	case HTTP:
		p := NewHTTPProbe()
		p.URL = c.Target
		return p, nil
	}

	return nil, ErrProbeTypeNotAvailable
}
