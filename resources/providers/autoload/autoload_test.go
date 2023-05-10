package autoload

import (
	"fmt"
	"testing"

	pb "github.com/edobtc/cloudkit/rpc/controlplane/resources/v1"
	"github.com/stretchr/testify/assert"
)

func TestExists(t *testing.T) {
	if Exists("whatever/bloop") {
		t.Error("Exists should not find whatever/bloop")
	}

	if Exists("aws/lambda") == false {
		t.Error("Exists should find aws/lambda")
	}
}

func TestProtoTargetMap(t *testing.T) {
	type test struct {
		input    pb.Target
		expected string
	}

	tests := []test{
		{input: pb.Target_TARGET_AWS_UNSPECIFIED, expected: "aws/ec2"},
		{input: pb.Target_TARGET_GCP, expected: "gcp/gce"},
		{input: pb.Target_TARGET_CLOUDFLARE, expected: "cloudflare"},
		{input: pb.Target_TARGET_DIGITALOCEAN, expected: "digitalocean/droplet"},
		{input: pb.Target_TARGET_DIGITALOCEAN_DROPLET, expected: "digitalocean/droplet"},
		{input: pb.Target_TARGET_AWS_EC2, expected: "aws/ec2"},
		{input: pb.Target_TARGET_AWS_LAMBDA, expected: "aws/lambda"},
		{input: pb.Target_TARGET_AWS_FARGATE, expected: "aws/fargate"},
		{input: pb.Target_TARGET_LINODE, expected: "linode"},
		{input: pb.Target_TARGET_KUBERNETES, expected: "k8s/deployment"},
		{input: pb.Target_TARGET_K8S, expected: "k8s/deployment"},
		{input: pb.Target_TARGET_DOCKER, expected: "docker"},
		{input: pb.Target_TARGET_MOCK_BLANK, expected: "test/blank"},
		{input: pb.Target_TARGET_MOCK_TIMED, expected: "test/timed"},
		{input: pb.Target_TARGET_AWS_SECURITY_GROUPS, expected: "aws/sg"},
	}

	for _, tc := range tests {
		received := ProtoTargetMap(&tc.input)
		assert.Equal(t, tc.expected, received, fmt.Sprintf("should return %s from %s", tc.expected, tc.input.String()))
	}
}
