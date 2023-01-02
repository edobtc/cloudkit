package autoload

import (
	"testing"

	"github.com/edobtc/cloudkit/events/publishers/aws/firehose"
	"github.com/edobtc/cloudkit/events/publishers/aws/kinesis"
	"github.com/edobtc/cloudkit/events/publishers/aws/lambda"
	"github.com/edobtc/cloudkit/events/publishers/aws/sns"
	"github.com/edobtc/cloudkit/events/publishers/filesystem"
	"github.com/edobtc/cloudkit/events/publishers/webhook"
	"github.com/edobtc/cloudkit/events/publishers/websocket"
	"github.com/stretchr/testify/assert"
)

func TestNewPublisherFirehose(t *testing.T) {
	fp, err := NewPublisher("firehose")
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, fp, &firehose.Publisher{})
}

func TestNewPublisherSns(t *testing.T) {
	fp, err := NewPublisher("sns")
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, fp, &sns.Publisher{})
}

func TestNewPublisherLambda(t *testing.T) {
	fp, err := NewPublisher("lambda")
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, fp, &lambda.Publisher{})
}

func TestNewPublisherKinesis(t *testing.T) {
	fp, err := NewPublisher("kinesis")
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, fp, &kinesis.Publisher{})
}

func TestNewPublisherFilesystem(t *testing.T) {
	fp, err := NewPublisher("filesystem")
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, fp, &filesystem.Publisher{})
}

func TestNewPublisherFilesystemShort(t *testing.T) {
	fp, err := NewPublisher("fs")
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, fp, &filesystem.Publisher{})
}

func TestNewPublisherWebhook(t *testing.T) {
	fp, err := NewPublisher("webhook")
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, fp, &webhook.Publisher{})
}

func TestNewPublisherWebhookShort(t *testing.T) {
	fp, err := NewPublisher("wh")
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, fp, &webhook.Publisher{})
}

func TestNewPublisherWebsocket(t *testing.T) {
	fp, err := NewPublisher("websocket")
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, fp, &websocket.Publisher{})
}

func TestNewPublisherWebsocketShort(t *testing.T) {
	fp, err := NewPublisher("ws")
	if err != nil {
		t.Fatal(err)
	}

	assert.IsType(t, fp, &websocket.Publisher{})
}
