package security_groups

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	log "github.com/sirupsen/logrus"
)

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
