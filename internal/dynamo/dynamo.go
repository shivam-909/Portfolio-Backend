package dynamo

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/shivam-909/portfolio-backend/internal/projects"
)

const projectTableName = "projects"

func CreateProject(db *dynamodb.DynamoDB, project projects.Project) error {

	projectMap, err := dynamodbattribute.MarshalMap(project)
	if err != nil {
		return err
	}

	input := &dynamodb.PutItemInput{
		Item:      projectMap,
		TableName: aws.String(projectTableName),
	}

	_, err = db.PutItem(input)
	if err != nil {
		return err
	}

	return nil
}
