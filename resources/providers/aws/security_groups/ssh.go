package security_groups

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
	log "github.com/sirupsen/logrus"
)

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
