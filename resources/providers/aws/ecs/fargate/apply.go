package fargate

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ecs"
	"github.com/edobtc/cloudkit/resources/providers/aws/auth"
	log "github.com/sirupsen/logrus"
)

// Apply runs the provisioner end to end, so calls
// read and clone

func (p *Provisioner) Apply() error {
	sess, err := auth.Session()
	if err != nil {
		return err
	}
	svc := ecs.New(sess)

	// Create a new task definition
	taskDefinitionInput := &ecs.RegisterTaskDefinitionInput{
		ContainerDefinitions: []*ecs.ContainerDefinition{
			{
				Name:  aws.String(p.Config.Name),
				Image: aws.String(p.Config.Version),
			},
		},
		Family: aws.String(fmt.Sprintf("%s-task", p.Config.Name)),
	}

	taskDefinitionOutput, err := svc.RegisterTaskDefinition(taskDefinitionInput)
	if err != nil {
		return err
	}

	log.Info("Created task definition:", aws.StringValue(taskDefinitionOutput.TaskDefinition.TaskDefinitionArn))

	// Create a new service
	serviceInput := &ecs.CreateServiceInput{
		Cluster:        aws.String(p.Config.Target), // Replace with your desired cluster
		ServiceName:    aws.String(fmt.Sprintf("%s-service", p.Config.Name)),
		TaskDefinition: aws.String(aws.StringValue(taskDefinitionOutput.TaskDefinition.TaskDefinitionArn)),
		LaunchType:     aws.String("FARGATE"),
		NetworkConfiguration: &ecs.NetworkConfiguration{
			AwsvpcConfiguration: &ecs.AwsVpcConfiguration{
				Subnets:        p.getSubnets(),
				SecurityGroups: p.getSGs(),
			},
		},
		PlatformVersion: aws.String("LATEST"),
		DesiredCount:    aws.Int64(1),
		Tags: []*ecs.Tag{
			{
				Key:   aws.String("cloudkit"),
				Value: aws.String(p.Config.Name),
			},
			{
				Key:   aws.String("cloudkit:provisioner"),
				Value: aws.String("cloudkit"),
			},
		},
	}

	serviceOutput, err := svc.CreateService(serviceInput)
	if err != nil {
		return err
	}

	log.Info("Created service:", aws.StringValue(serviceOutput.Service.ServiceArn))

	return nil
}

func (p *Provisioner) getSubnets() []*string {
	// scan and read from target environment if missing
	nets := []*string{}
	for _, net := range p.Config.SubnetIDs {
		nets = append(nets, aws.String(net))
	}
	return nets
}

func (p *Provisioner) getSGs() []*string {
	// scan and read from existing configuration if missing
	sgs := []*string{}
	for _, sg := range p.Config.SecurityGroups {
		sgs = append(sgs, aws.String(sg))
	}
	return sgs
}
