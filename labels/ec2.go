package labels

import "github.com/aws/aws-sdk-go/service/ec2"

func FromEC2Tag(tag *ec2.Tag) Label {
	label := Label{
		Name:  *tag.Key,
		Value: *tag.Value,
	}

	return label
}

func FromEC2Tags(tags []*ec2.Tag) []Label {
	labels := []Label{}
	for _, tag := range tags {
		labels = append(labels, Label{
			Name:  *tag.Key,
			Value: *tag.Value,
		})
	}

	return labels
}
