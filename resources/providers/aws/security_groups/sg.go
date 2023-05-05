package main

import (
	"github.com/edobtc/cloudkit/resources/providers/aws/auth"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func CreateSecurityGroups() error {
	sess, err := auth.Session()
	if err != nil {
		return err
	}

	svc := ec2.New(sess)

	subscriberSecurityGroupId, err := createSubscriberSecurityGroup(svc)
	if err != nil {
		return err
	}
	log.Debug("Created security group for zmq_subscriber resources:", subscriberSecurityGroupId)

	selfSecurityGroupInput := createSelfSecurityGroupInput(subscriberSecurityGroupId)
	authorizeSelfSecurityGroupIngress(svc, selfSecurityGroupInput)
	log.Debug("Authorized ingress from the same security group for security group:", subscriberSecurityGroupId)

	sshSecurityGroupId := createSshSecurityGroup(svc)
	log.Debug("Created security group for SSH access:", sshSecurityGroupId)

	sshSecurityGroupInput := createSshSecurityGroupInput(sshSecurityGroupId)
	authorizeSshSecurityGroupIngress(svc, sshSecurityGroupInput)
	log.Debug("Authorized ingress for SSH access for security group:", sshSecurityGroupId)

	bitcoinP2PSecurityGroupId, err := createBitcoinP2PSecurityGroup(svc)
	if err != nil {
		return err
	}
	log.Debug("Created security group for bitcoin_p2p resources:", bitcoinP2PSecurityGroupId)

	bitcoinP2PSecurityGroupInput := createBitcoinP2PSecurityGroupInput(bitcoinP2PSecurityGroupId)
	authorizeBitcoinP2PSecurityGroupIngress(svc, bitcoinP2PSecurityGroupInput)
	log.Debug("Authorized ingress for bitcoin_p2p access for security group:", bitcoinP2PSecurityGroupId)

	zmqBroadcasterSecurityGroupId, err := createZmqBroadcasterSecurityGroup(svc)
	if err != nil {
		return err
	}
	log.Debug("Created security group for zmq_broadcaster resources:", zmqBroadcasterSecurityGroupId)

	return nil
}

func createSubscriberSecurityGroup(svc *ec2.EC2) (string, error) {
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

	createSubscriberSecurityGroupOutput, err := svc.CreateSecurityGroup(createSubscriberSecurityGroupInput)
	if err != nil {
		log.Debug("Error creating security group for zmq_subscriber resources:", err)
		return "", err
	}

	return *createSubscriberSecurityGroupOutput.GroupId, nil
}

func createSshSecurityGroupInput(sshSecurityGroupId string) *ec2.AuthorizeSecurityGroupIngressInput {
	return &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: aws.String(sshSecurityGroupId),
		IpPermissions: []*ec2.IpPermission{
			{
				IpProtocol: aws.String("tcp"),
				FromPort:   aws.Int64(22),
				ToPort:     aws.Int64(22),
				IpRanges: []*ec2.IpRange{
					{
						CidrIp: aws.String("0.0.0.0/0"),
					},
				},
			},
		},
	}
}

func createBitcoinP2PSecurityGroup(svc *ec2.EC2) (string, error) {
	bitcoinP2PSecurityGroupName := "bitcoin_p2p_sg"
	bitcoinP2PSecurityGroupDescription := "Security group for bitcoin_p2p resources"

	createBitcoinP2PSecurityGroupInput := &ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(bitcoinP2PSecurityGroupName),
		Description: aws.String(bitcoinP2PSecurityGroupDescription),
	}

	createBitcoinP2PSecurityGroupOutput, err := svc.CreateSecurityGroup(createBitcoinP2PSecurityGroupInput)
	if err != nil {
		log.Debug("Error creating security group for bitcoin_p2p resources:", err)
		return "", err
	}

	return *createBitcoinP2PSecurityGroupOutput.GroupId, nil
}

func createBitcoinP2PSecurityGroupInput(bitcoinP2PSecurityGroupId string) *ec2.AuthorizeSecurityGroupIngressInput {
	return &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: aws.String(bitcoinP2PSecurityGroupId),
		IpPermissions: []*ec2.IpPermission{
			{
				IpProtocol: aws.String("tcp"),
				FromPort:   aws.Int64(8333),
				ToPort:     aws.Int64(8333),
				IpRanges: []*ec2.IpRange{
					{
						CidrIp: aws.String("0.0.0.0/0"),
					},
				},
			},
		},
	}
}

func authorizeBitcoinP2PSecurityGroupIngress(svc *ec2.EC2, input *ec2.AuthorizeSecurityGroupIngressInput) error {
	_, err := svc.AuthorizeSecurityGroupIngress(input)
	if err != nil {
		log.Debug("Error authorizing ingress for SSH access:", err)
		return err
	}
	return nil
}

func authorizeSshSecurityGroupIngress(svc *ec2.EC2, input *ec2.AuthorizeSecurityGroupIngressInput) error {
	_, err := svc.AuthorizeSecurityGroupIngress(input)
	if err != nil {
		log.Debug("Error authorizing ingress for SSH access:", err)
		return err
	}
	return nil
}

func createSelfSecurityGroupInput(subscriberSecurityGroupId string) *ec2.AuthorizeSecurityGroupIngressInput {
	return &ec2.AuthorizeSecurityGroupIngressInput{
		GroupId: aws.String(subscriberSecurityGroupId),
		IpPermissions: []*ec2.IpPermission{
			{
				IpProtocol: aws.String("-1"),
				UserIdGroupPairs: []*ec2.UserIdGroupPair{
					{
						GroupId: aws.String(subscriberSecurityGroupId),
					},
				},
			},
		},
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
}

func authorizeSelfSecurityGroupIngress(svc *ec2.EC2, input *ec2.AuthorizeSecurityGroupIngressInput) error {
	_, err := svc.AuthorizeSecurityGroupIngress(input)
	if err != nil {
		log.Debug("Error authorizing ingress from the same security group:", err)
		return err
	}

	return nil
}

func createSshSecurityGroup(svc *ec2.EC2) string {
	sshSecurityGroupName := "ssh_sg"
	sshSecurityGroupDescription := "Security group for SSH access"

	createSshSecurityGroupInput := &ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(sshSecurityGroupName),
		Description: aws.String(sshSecurityGroupDescription),
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

	createSshSecurityGroupOutput, err := svc.CreateSecurityGroup(createSshSecurityGroupInput)
	if err != nil {
		log.Debug("Error creating security group for SSH access:", err)
		return ""
	}

	return *createSshSecurityGroupOutput.GroupId
}

func createZmqBroadcasterSecurityGroup(svc *ec2.EC2) (string, error) {
	zmqBroadcasterSecurityGroupName := "zmq_broadcaster_sg"
	zmqBroadcasterSecurityGroupDescription := "Security group for zmq_broadcaster resources"

	createZmqBroadcasterSecurityGroupInput := &ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(zmqBroadcasterSecurityGroupName),
		Description: aws.String(zmqBroadcasterSecurityGroupDescription),
	}

	createZmqBroadcasterSecurityGroupOutput, err := svc.CreateSecurityGroup(createZmqBroadcasterSecurityGroupInput)
	if err != nil {
		log.Debug("Error creating security group for zmq_broadcaster resources:", err)
		return "", err
	}

	return *createZmqBroadcasterSecurityGroupOutput.GroupId, nil
}
