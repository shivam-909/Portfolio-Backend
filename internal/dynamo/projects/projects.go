package projects

import (
	"fmt"

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

	var queryInput = &dynamodb.QueryInput{
		TableName: aws.String(projectTableName),
		KeyConditions: map[string]*dynamodb.Condition{
			"id": {
				ComparisonOperator: aws.String("EQ"),
				AttributeValueList: []*dynamodb.AttributeValue{
					{
						N: aws.String(fmt.Sprintf("%d", id)),
					},
				},
			},
		},
	}

	var result, err = db.Query(queryInput)
	if err != nil {
		return project, err
	}

	if len(result.Items) == 0 {
		return project, fmt.Errorf("no project found with id %d", id)
	}

	err = dynamodbattribute.UnmarshalMap(result.Items[0], &project)
	if err != nil {
		return project, err
	}

	return project, nil
}
