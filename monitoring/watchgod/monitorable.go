package watchgod

import (
	"context"
	"net/http"
	"time"

	"github.com/edobtc/cloudkit/httpclient"
)

type Monitorable struct {
	URL             string
	Status          bool
	StatusWindow    []bool
	WindowLength    int
	MaxErrThreshold int
	Interval        time.Duration
	NextCheck       *time.Time
	LastChecked     *time.Time
	Protocol        Protocol
	clientHTTP      *http.Client
}

func NewMonitorable() *Monitorable {
	return &Monitorable{
		StatusWindow:    []bool{},
		Protocol:        HTTP,
		WindowLength:    DefaultWindowLength,
		MaxErrThreshold: DefaultMaxErrThreshold,
		clientHTTP:      httpclient.New(),
		Interval:        DefaultMonitorableInterval,
	}
}

func (m *Monitorable) ShouldCheck() bool {
	if m.LastChecked == nil {
		return true
	}

	return time.Since(*m.LastChecked) >= m.Interval
}

func (m *Monitorable) AddStatus(status bool) int {
	if m.WindowLength > 0 {
		if len(m.StatusWindow) >= m.WindowLength {
			m.StatusWindow = m.StatusWindow[1:]
		}
	}

	m.StatusWindow = append(m.StatusWindow, status)
	return len(m.StatusWindow)
}

func (m *Monitorable) Check(ctx context.Context) (bool, error) {
	status, err := m.dispatch()
	if err != nil {
		m.AddStatus(status)
		return false, err
	}

	m.AddStatus(status)

	return status, nil
}

func (m *Monitorable) dispatch() (bool, error) {
	switch m.Protocol {
	case HTTP:
		return m.fetchHTTP()
	default:
		return false, nil
	}
}

func (m *Monitorable) ErrThresholdExceeded() bool {
	var errcount = 0
	for _, status := range m.StatusWindow {
		if !status {
			errcount++
		}
	}

	return errcount >= m.MaxErrThreshold
}

func (m *Monitorable) fetchHTTP() (bool, error) {
	if m.clientHTTP == nil {
		m.clientHTTP = httpclient.New()
	}

	checkTime := time.Now()
	m.LastChecked = &checkTime

	req, err := http.NewRequest("GET", m.URL, nil)
	if err != nil {
		return false, err
	}

	client := httpclient.New()

	resp, err := client.Do(req)

	if err != nil {
		return false, err
	}

	defer resp.Body.Close()

	status := (resp.StatusCode == http.StatusOK)

	return status, nil
}
