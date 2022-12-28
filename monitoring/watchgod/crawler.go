package watchgod

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	log "github.com/sirupsen/logrus"
)

func NewCrawler(ctx context.Context, cfg Config) *Crawler {
	s := make(chan os.Signal, 1)

	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	return &Crawler{
		ctx:            ctx,
		tickInterval:   defaultTick,
		s:              s,
		InfoCheckPoint: DefaultInfoCheckPoint,
		EventChan:      make(chan Monitorable),
		Finished:       make(chan bool),
		ErrChan:        make(chan error),
	}
}

func NewCrawlerFromConfig() *Crawler {
	s := make(chan os.Signal, 1)

	signal.Notify(s, syscall.SIGINT, syscall.SIGTERM)

	return &Crawler{
		tickInterval:    defaultTick,
		InfoCheckPoint:  DefaultInfoCheckPoint,
		s:               s,
		MaxErrThreshold: DefaultCrawlerMaxErrThreshold,
		EventChan:       make(chan Monitorable),
		Finished:        make(chan bool),
		ErrChan:         make(chan error),
	}
}

type Crawler struct {
	sync.Mutex

	ctx             context.Context
	Monitorable     []*Monitorable
	tickInterval    time.Duration
	tickCount       int
	InfoCheckPoint  int
	ErrCount        int
	MaxErrThreshold int
	s               chan (os.Signal)
	EventChan       chan (Monitorable)
	ErrChan         chan (error)
	Finished        chan (bool)
}

func (c *Crawler) Start() {
	ticker := time.NewTicker(c.tickInterval)
	defer ticker.Stop()

	go func() {
		log.Info("starting watch")
		<-c.s
		log.Info("received signal")
		c.Stop()
	}()

	for {
		select {
		case <-c.Finished:
			fmt.Println("Done!")
		case <-ticker.C:
			c.progress()

			for _, mtr := range c.Monitorable {
				if !mtr.ShouldCheck() {
					continue
				}
				go func(m *Monitorable) {
					check, err := m.Check(context.TODO())
					if check {
						// Success
						c.EventChan <- *m
					} else {
						// Fail case handling
						log.Info("failed")

						if m.ErrThresholdExceeded() {
							c.Remove(m)
							c.ErrChan <- fmt.Errorf("%s err threshold exceeded. removing", m.URL)
						}

						if err != nil {
							c.ErrChan <- err
						}
					}
				}(mtr)
			}
		}
	}
}

func (c *Crawler) Add(m *Monitorable) int {
	c.Lock()
	defer c.Unlock()
	c.Monitorable = append(c.Monitorable, m)

	return len(c.Monitorable)
}

func (c *Crawler) Remove(m *Monitorable) bool {
	for idx, mtr := range c.Monitorable {
		if mtr.URL == m.URL {
			return c.RemovePosition(idx)
		}
	}
	return false
}

func (c *Crawler) progress() {
	c.tickCount++

	if (c.tickCount % c.InfoCheckPoint) == 0 {
		log.Info("Size:", len(c.Monitorable))
	}
}

func (c *Crawler) RemovePosition(idx int) bool {
	c.Lock()
	defer c.Unlock()
	replacement := make([]*Monitorable, 0)
	replacement = append(replacement, c.Monitorable[:idx]...)
	c.Monitorable = append(replacement, c.Monitorable[idx+1:]...)

	return true
}

func (c *Crawler) Wait() chan (bool) {
	return c.Finished
}

func (c *Crawler) Status() (chan (string), error) {
	return nil, nil
}

func (c *Crawler) Stop() chan (string) {
	log.Info("stopping")
	c.Finished <- true
	return nil
}
