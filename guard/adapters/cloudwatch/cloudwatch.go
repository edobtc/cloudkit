package cloudwatch

import (
	"time"

	"github.com/edobtc/cloudkit/aws/session"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/cloudwatch"
	log "github.com/sirupsen/logrus"
)

type Dimensions map[string]string

// Config is the collection of configuration options specific
// to the cloudwatch guard implementation
type Config struct {
	Metric     string
	Namespace  string
	Dimensions Dimensions
	Tolerance  float32
}

// Guard is the cloudwatch guard
type Guard struct {
	Config Config
}

// Check is the method for performing the guard check
// that will be called by any experiment orchestrator
// to maintain ongoing checks of remote state/status, drift,
// change, etc...
func (g *Guard) Check() (bool, error) {
	svc := cloudwatch.New(session.NewDynamicSession())

	input := &cloudwatch.GetMetricStatisticsInput{
		MetricName: aws.String(g.Config.Metric),
		Namespace:  aws.String(g.Config.Namespace),
		StartTime:  aws.Time(time.Now().Add(-1 * time.Minute)),
		EndTime:    aws.Time(time.Now()),
		Period:     aws.Int64(60),
		ExtendedStatistics: []*string{
			aws.String("p0.0"),
			aws.String("p0.5"),
			aws.String("p99"),
			aws.String("p100"),
		},
	}

	if len(g.Config.Dimensions) > 0 {
		filter := []*cloudwatch.Dimension{}
		for name, value := range g.Config.Dimensions {
			filter = append(
				filter,
				&cloudwatch.Dimension{
					Name:  aws.String(name),
					Value: aws.String(value),
				},
			)
		}
		input.Dimensions = filter
	}

	result, err := svc.GetMetricStatistics(input)

	if err != nil {
		return false, err
	}

	log.Info(result.Datapoints)

	// TODO: Check that we're within tolerance here
	// return true or not otherwise

	return true, nil
}
