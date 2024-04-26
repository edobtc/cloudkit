package publishers

import (
	"github.com/edobtc/cloudkit/events/publishers/aws/firehose"
	"github.com/edobtc/cloudkit/events/publishers/aws/kinesis"
	"github.com/edobtc/cloudkit/events/publishers/aws/sns"
	"github.com/edobtc/cloudkit/events/publishers/filesystem"
	"github.com/edobtc/cloudkit/events/publishers/rmq"
)

// Publisher defines the interface for event
// publishers which are currently defined as:
//
// Kinesis (publish directly to kinesis)
// Firehose (send to a kinesis firehose)
// Filesystem (for local events)
type Publisher interface {
	Send([]byte) error
	Close() error
}

var (
	_ Publisher = (*firehose.Publisher)(nil)
	_ Publisher = (*kinesis.Publisher)(nil)
	_ Publisher = (*filesystem.Publisher)(nil)
	_ Publisher = (*sns.Publisher)(nil)
	_ Publisher = (*rmq.Publisher)(nil)
)
