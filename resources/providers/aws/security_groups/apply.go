package security_groups

import (
	"github.com/edobtc/cloudkit/resources/providers/aws/auth"

	log "github.com/sirupsen/logrus"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

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
