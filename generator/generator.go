package generator

import (
	"time"

	log "github.com/sirupsen/logrus"
)

const (
	DefaultInterval = 5 * time.Second
	DefaultMaxCount = 200
)

// Default implementation
type Config struct {
	Handler  Handler
	Interval time.Duration
	MaxCount int
}

func DefaultConfig() Config {
	return Config{
		Interval: DefaultInterval,
		MaxCount: DefaultMaxCount,
	}
}

type Handler func() error

// Generator interface if there is a need
type Generator interface {
	Start() chan bool
	Stop() bool
	Errors() chan error
}

// Default implementation
type DefaultGenerator struct {
	Handler        Handler
	Interval       time.Duration
	MaxCount       int
	ProcessedCount int
	errChan        chan (error)
	finished       chan (bool)
}

// defaultHandler is the NOOP default handler
// that should over ridden and injected
var defaultHandler = func() error {
	log.Info("generated event")
	return nil
}

// Return a new Generator returning interface
// with an injected handler
func NewWitConfig(cfg Config) Generator {
	return &DefaultGenerator{
		Handler:  cfg.Handler,
		Interval: cfg.Interval,
		MaxCount: cfg.MaxCount,
		errChan:  make(chan error),
		finished: make(chan bool),
	}
}

// Return a new Generator returning interface
// with an injected handler
func NewWithHandler(h Handler) Generator {
	return &DefaultGenerator{
		Handler:  h,
		Interval: DefaultInterval,
		MaxCount: DefaultMaxCount,
		errChan:  make(chan error),
		finished: make(chan bool),
	}
}

// Return a new Generator returning interface
// with the default handler
func New() Generator {
	return &DefaultGenerator{
		Handler:  defaultHandler,
		Interval: DefaultInterval,
		MaxCount: DefaultMaxCount,
		errChan:  make(chan error),
		finished: make(chan bool),
	}
}

// Start
func (d *DefaultGenerator) Start() chan bool {
	ticker := time.NewTicker(d.Interval)

	go func() {
		for range ticker.C {
			d.Generate()

			d.ProcessedCount++

			if d.ProcessedCount >= d.MaxCount {
				ticker.Stop()
				d.Stop()
			}
		}
	}()

	return d.finished
}

func (d *DefaultGenerator) Generate() {
	err := d.Handler()
	if err != nil {
		d.errChan <- err
	}
}

func (d *DefaultGenerator) Errors() chan error {
	return d.errChan
}

func (d *DefaultGenerator) Stop() bool {
	close(d.finished)
	return true
}
