package security_groups

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	"github.com/edobtc/cloudkit/resources/providers/aws/auth"
)

func Select() ([]*ec2.SecurityGroup, error) {
	tagKey := "creator"
	tagValue := "cloudkit"

	sess, err := auth.Session()
	if err != nil {
		return nil, err
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
		return nil, err
	}

	return describeSecurityGroupsOutput.SecurityGroups, nil

}
