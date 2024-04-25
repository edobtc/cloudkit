package autoload

import (
	"testing"

	eclair "github.com/edobtc/cloudkit/events/subscribers/lightning/eclair/ws"
	zmq "github.com/edobtc/cloudkit/events/subscribers/zmq"

	"github.com/stretchr/testify/assert"
)

func TestNewSubscriberZMQ(t *testing.T) {
	fp, err := NewSubscriber("zmq")
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, fp)
	assert.IsType(t, fp, &zmq.Subscriber{})
}

func TestNewSubscriberEclair(t *testing.T) {
	fp, err := NewSubscriber("eclair")
	if err != nil {
		t.Fatal(err)
	}

	assert.NotNil(t, fp)
	assert.IsType(t, fp, &eclair.EclairSubscriber{})
}
