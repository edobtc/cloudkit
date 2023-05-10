package labels

import (
	"testing"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/ec2"
)

func TestConversionFromEc2Tag(t *testing.T) {
	tag := &ec2.Tag{
		Key:   aws.String("foo"),
		Value: aws.String("bar"),
	}

	label := FromEC2Tag(tag)

	if label.Name != "foo" {
		t.Errorf("Expected label name to be foo, got %s", label.Name)
	}

	if label.Value != "bar" {
		t.Errorf("Expected label value to be bar, got %s", label.Value)
	}
}

func TestConversionFromEc2Tags(t *testing.T) {
	tags := []*ec2.Tag{
		{
			Key:   aws.String("foo"),
			Value: aws.String("bar"),
		},
		{
			Key:   aws.String("baz"),
			Value: aws.String("qux"),
		},
	}

	labels := FromEC2Tags(tags)

	if len(labels) != 2 {
		t.Errorf("Expected 2 labels, got %d", len(labels))
	}

	if labels[0].Name != "foo" {
		t.Errorf("Expected label name to be foo, got %s", labels[0].Name)
	}

	if labels[0].Value != "bar" {
		t.Errorf("Expected label value to be bar, got %s", labels[0].Value)
	}

	if labels[1].Name != "baz" {
		t.Errorf("Expected label name to be baz, got %s", labels[1].Name)
	}

	if labels[1].Value != "qux" {
		t.Errorf("Expected label value to be qux, got %s", labels[1].Value)
	}
}
