package security_groups

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/edobtc/cloudkit/resources/providers/aws/auth"
	"github.com/edobtc/cloudkit/target"
)

func (p *SecurityGroupsProvider) Select() (target.Selection, error) {
	tagKey := "creator"
	tagValue := "cloudkit"
	t := target.Selection{
		Selected: target.Resources{},
	}

	sess, err := auth.Session()
	if err != nil {
		return t, err
	}

	svc := ec2.New(sess)

	filter := &ec2.Filter{
		Name: aws.String("tag:" + tagKey),
		Values: []*string{
			aws.String(tagValue),
		},
	}

	describeSecurityGroupsInput := &ec2.DescribeSecurityGroupsInput{
		Filters: []*ec2.Filter{filter},
	}

	describeSecurityGroupsOutput, err := svc.DescribeSecurityGroups(describeSecurityGroupsInput)
	if err != nil {
		return t, err
	}

	for _, sg := range describeSecurityGroupsOutput.SecurityGroups {
		r := target.Resource{
			ID:   *sg.GroupId,
			Name: string(*sg.GroupName),
			Meta: sg.Tags,
		}
		t.Selected = append(t.Selected, r)
	}

	return t, nil

}
