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

func RetrieveAllProjects(db *dynamodb.DynamoDB) ([]projects.Project, error) {

	var projects []projects.Project

	var queryInput = &dynamodb.ScanInput{
		TableName: aws.String(projectTableName),
	}

	var result, err = db.Scan(queryInput)
	if err != nil {
		return projects, err
	}

	err = dynamodbattribute.UnmarshalListOfMaps(result.Items, &projects)
	if err != nil {
		return projects, err
	}

	return projects, nil
}

func DeleteProject(db *dynamodb.DynamoDB, id int64) error {

	var queryInput = &dynamodb.DeleteItemInput{
		TableName: aws.String(projectTableName),
		Key: map[string]*dynamodb.AttributeValue{
			"id": {
				N: aws.String(fmt.Sprintf("%d", id)),
			},
		},
	}

	_, err := db.DeleteItem(queryInput)
	if err != nil {
		return err
	}

	return nil
}

func SyncProjects(db *dynamodb.DynamoDB, projects []projects.Project) error {

	existingProjects, err := RetrieveAllProjects(db)
	if err != nil {
		return err
	}

	// Get all the ones that exist
	existingIDs := make(map[int64]bool)
	for _, project := range existingProjects {
		existingIDs[project.ID] = true
	}

	// For each new project, if it exists, mark it as resolved or create it.
	for _, project := range projects {
		if !existingIDs[project.ID] {
			err = CreateProject(db, project)
			if err != nil {
				return err
			}
		} else {
			existingIDs[project.ID] = false
		}
	}

	// Anything we already have but isn't in the new list, delete.
	for id, exists := range existingIDs {
		if exists {
			err = DeleteProject(db, id)
			if err != nil {
				return err
			}
		}
	}

	return nil
}
