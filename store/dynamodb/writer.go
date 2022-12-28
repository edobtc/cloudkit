package dynamodb

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// DynamoWriter is a single objective
// writer of objects to dynamo
type DynamoWriter struct {
	svc       *dynamodb.DynamoDB
	tableName string
}

// NewDynamoWriter returns a new, default value DynamoWriter
func NewDynamoWriter(tableName string) *DynamoWriter {
	return &DynamoWriter{
		svc:       dynamodb.New(newSession()),
		tableName: tableName,
	}
}

// newSession privately creates a new session object for passing
// to aws sdk go svc initializers
func newSession() *session.Session {
	return session.Must(session.NewSession())
}

func (d *DynamoWriter) WriteJson(entity interface{}) error {
	ety, err := dynamodbattribute.Marshal(entity)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      map[string]*dynamodb.AttributeValue{"object": ety},
		TableName: aws.String(d.tableName),
	}

	_, err = d.svc.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

func (d *DynamoWriter) Write(entity interface{}) error {
	ety, err := dynamodbattribute.MarshalMap(entity)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      ety,
		TableName: aws.String(d.tableName),
	}

	_, err = d.svc.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}

// Delete removes the item from the
// given table
func (d *DynamoWriter) Delete(entity interface{}) error {
	ety, err := dynamodbattribute.MarshalMap(d)
	if err != nil {
		return err
	}

	input := &dynamodb.DeleteItemInput{
		Key:       ety,
		TableName: aws.String(d.tableName),
	}

	_, err = d.svc.DeleteItem(input)
	if err != nil {
		return err
	}

	return nil
}
