package security_groups

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	log "github.com/sirupsen/logrus"
)

func (b *SecurityGroupBuilder) CreateSubscriberSecurityGroup() (string, error) {
	subscriberSecurityGroupName := "zmq_subscriber_sg"
	subscriberSecurityGroupDescription := "Security group for zmq_subscriber resources"

	createSubscriberSecurityGroupInput := &ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(subscriberSecurityGroupName),
		Description: aws.String(subscriberSecurityGroupDescription),
		TagSpecifications: []*ec2.TagSpecification{
			{
				ResourceType: aws.String("security-group"),
				Tags: []*ec2.Tag{
					{
						Key:   aws.String("creator"),
						Value: aws.String("cloudkit"),
					},
				},
			},
		},
	}

	createSubscriberSecurityGroupOutput, err := b.svc.CreateSecurityGroup(createSubscriberSecurityGroupInput)
	if err != nil {
		log.Debug("Error creating security group for zmq_subscriber resources:", err)
		return "", err
	}

	return *createSubscriberSecurityGroupOutput.GroupId, nil
}

func (b *SecurityGroupBuilder) CreateZmqBroadcasterSecurityGroup() (string, error) {
	zmqBroadcasterSecurityGroupName := "zmq_broadcaster_sg"
	zmqBroadcasterSecurityGroupDescription := "Security group for zmq_broadcaster resources"

	createZmqBroadcasterSecurityGroupInput := &ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(zmqBroadcasterSecurityGroupName),
		Description: aws.String(zmqBroadcasterSecurityGroupDescription),
		VpcId:       aws.String(b.cfg.VpcId),
	}

	createZmqBroadcasterSecurityGroupOutput, err := b.svc.CreateSecurityGroup(createZmqBroadcasterSecurityGroupInput)
	if err != nil {
		log.Debug("Error creating security group for zmq_broadcaster resources:", err)
		return "", err
	}

	return *createZmqBroadcasterSecurityGroupOutput.GroupId, nil
}
