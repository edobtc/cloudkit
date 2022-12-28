package watchgod

import "time"

const (
	// CrawlerDefaults
	defaultTick                   = 100 * time.Millisecond
	DefaultInfoCheckPoint         = 100
	DefaultCrawlerMaxErrThreshold = 20

	// Monitorable Defaults
	DefaultWindowLength        = 5
	DefaultMaxErrThreshold     = 4
	DefaultMonitorableInterval = 45 * time.Second
)
