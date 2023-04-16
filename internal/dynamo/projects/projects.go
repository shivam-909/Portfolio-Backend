package projects

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"github.com/shivam-909/portfolio-backend/pkg/projects"
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

func RetrieveProject(db *dynamodb.DynamoDB, id int32) (projects.Project, error) {

	var project projects.Project

	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(string(id)),
			},
		},
		TableName: aws.String(projectTableName),
	}

	result, err := db.GetItem(input)
	if err != nil {
		return project, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &project)
	if err != nil {
		return project, err
	}

	return project, nil
}
