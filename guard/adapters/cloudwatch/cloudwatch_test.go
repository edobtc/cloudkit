package cloudwatch

import (
	"testing"

	log "github.com/sirupsen/logrus"
)

func TestEverything(t *testing.T) {
	c := Config{
		Metric:    "test",
		Namespace: "AWS/EC2",
		Dimensions: Dimensions{
			"huh": "what",
		},
		Tolerance: 1.5,
	}

	g := Guard{
		Config: c,
	}

	status, err := g.Check()

	if err != nil {
		log.Error(err)
	}

	log.Info(status)

}
