package security_groups

import (
	"github.com/edobtc/cloudkit/resources/providers/aws/auth"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

type Config struct {
	VpcId  string
	DryRun bool
}

type SecurityGroupBuilder struct {
	cfg Config
	svc *ec2.EC2
}

func NewSecurityGroupBuilder(cfg Config) (*SecurityGroupBuilder, error) {
	sess, err := auth.Session()
	if err != nil {
		return nil, err
	}

	return &SecurityGroupBuilder{
		cfg: cfg,
		svc: ec2.New(sess),
	}, nil
}

func (b *SecurityGroupBuilder) CreateSecurityGroups() error {
	subscriberSecurityGroupId, err := b.CreateSubscriberSecurityGroup()
	if err != nil {
		return err
	}
	log.Debug("Created security group for zmq_subscriber resources:", subscriberSecurityGroupId)

	selfSecurityGroupInput := b.ListenerSelfSecurityGroupInput(subscriberSecurityGroupId)
	b.AuthorizeSecurityGroupIngress(selfSecurityGroupInput)
	log.Debug("Authorized ingress from the same security group for security group:", subscriberSecurityGroupId)

	sshSecurityGroupId := b.CreateSshSecurityGroup()
	log.Debug("Created security group for SSH access:", sshSecurityGroupId)

	sshSecurityGroupInput := b.CreateSshSecurityGroupInput(sshSecurityGroupId)
	b.AuthorizeSecurityGroupIngress(sshSecurityGroupInput)
	log.Debug("Authorized ingress for SSH access for security group:", sshSecurityGroupId)

	bitcoinP2PSecurityGroupId, err := b.CreateBitcoinP2PSecurityGroup()
	if err != nil {
		return err
	}
	log.Debug("Created security group for bitcoin_p2p resources:", bitcoinP2PSecurityGroupId)

	bitcoinP2PSecurityGroupInput := b.CreateBitcoinP2PSecurityGroupInput(bitcoinP2PSecurityGroupId)
	b.AuthorizeSecurityGroupIngress(bitcoinP2PSecurityGroupInput)
	log.Debug("Authorized ingress for bitcoin_p2p access for security group:", bitcoinP2PSecurityGroupId)

	zmqBroadcasterSecurityGroupId, err := b.CreateZmqBroadcasterSecurityGroup()
	if err != nil {
		return err
	}
	log.Debug("Created security group for zmq_broadcaster resources:", zmqBroadcasterSecurityGroupId)

	return nil
}

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

func (b *SecurityGroupBuilder) CreateSshSecurityGroupInput(sshSecurityGroupId string) *ec2.AuthorizeSecurityGroupIngressInput {
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

func (b *SecurityGroupBuilder) CreateBitcoinP2PSecurityGroup() (string, error) {
	bitcoinP2PSecurityGroupName := "bitcoin_p2p_sg"
	bitcoinP2PSecurityGroupDescription := "Security group for bitcoin_p2p resources"

	createBitcoinP2PSecurityGroupInput := &ec2.CreateSecurityGroupInput{
		GroupName:   aws.String(bitcoinP2PSecurityGroupName),
		Description: aws.String(bitcoinP2PSecurityGroupDescription),
	}

	createBitcoinP2PSecurityGroupOutput, err := b.svc.CreateSecurityGroup(createBitcoinP2PSecurityGroupInput)
	if err != nil {
		log.Debug("Error creating security group for bitcoin_p2p resources:", err)
		return "", err
	}

	return *createBitcoinP2PSecurityGroupOutput.GroupId, nil
}

func (b *SecurityGroupBuilder) CreateBitcoinP2PSecurityGroupInput(bitcoinP2PSecurityGroupId string) *ec2.AuthorizeSecurityGroupIngressInput {
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

func (b *SecurityGroupBuilder) AuthorizeSecurityGroupIngress(input *ec2.AuthorizeSecurityGroupIngressInput) error {
	_, err := b.svc.AuthorizeSecurityGroupIngress(input)
	if err != nil {
		return err
	}
	return nil
}

func (b *SecurityGroupBuilder) ListenerSelfSecurityGroupInput(subscriberSecurityGroupId string) *ec2.AuthorizeSecurityGroupIngressInput {
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

func (b *SecurityGroupBuilder) CreateSshSecurityGroup() string {
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

	createSshSecurityGroupOutput, err := b.svc.CreateSecurityGroup(createSshSecurityGroupInput)
	if err != nil {
		log.Debug("Error creating security group for SSH access:", err)
		return ""
	}

	return *createSshSecurityGroupOutput.GroupId
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
