package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

// reader of objects to dynamo
type DynamoReader struct {
	svc       *dynamodb.DynamoDB
	tableName string
}

// NewDynamoReader returns a new, default value DynamoReader
func NewDynamoReader(tableName string) *DynamoReader {
	return &DynamoReader{
		svc:       dynamodb.New(newSession()),
		tableName: tableName,
	}
}

func (dr *DynamoReader) GetByID(id string, sort string) (*dynamodb.GetItemOutput, error) {
	return dr.svc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String(dr.tableName),
		Key: map[string]*dynamodb.AttributeValue{
			"PK": {
				S: aws.String(id),
			},
			"SK": {
				S: aws.String(sort),
			},
		},
	})
}
