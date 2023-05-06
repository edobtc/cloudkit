package route53

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/route53"
)

const (
	RecordType = "A"
)

// Apply runs the Route53Provider end to end
// which in this instance is only associated with
// creating a dns record
func (p *Route53Provider) Apply() error {
	target, err := p.Select()
	if err != nil {
		return err
	}

	createRecordInput := &route53.ChangeResourceRecordSetsInput{
		HostedZoneId: aws.String(target.Data[0].ID),
		ChangeBatch: &route53.ChangeBatch{
			Changes: []*route53.Change{
				{
					Action: aws.String("CREATE"),
					ResourceRecordSet: &route53.ResourceRecordSet{
						Name: aws.String(p.Config.Name),
						Type: aws.String(RecordType),
						TTL:  aws.Int64(300),
						ResourceRecords: []*route53.ResourceRecord{
							{
								Value: aws.String(p.Config.Value),
							},
						},
					},
				},
			},
		},
	}
	_, err = p.svc.ChangeResourceRecordSets(createRecordInput)
	if err != nil {
		return err
	}

	return nil
}
