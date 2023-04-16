package main

import (
	"context"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
	"github.com/aws/aws-lambda-go/lambda"
	"github.com/google/go-github/v51/github"
	"github.com/shivam-909/portfolio-backend/internal/dynamo"
	dynamo_projects "github.com/shivam-909/portfolio-backend/internal/dynamo/projects"
	"github.com/shivam-909/portfolio-backend/pkg/projects"
	"github.com/shivam-909/portfolio-backend/pkg/response"
	"golang.org/x/oauth2"
)

var (
	GITHUB_TOKEN string
)

func main() {
	lambda.Start(handler)
}

// Create a flag called GITHUB_TOKEN and set it to the github token

// Makes a request call to the github api to get the latest projects
// whichever projects are no longer present in the response are deleted,
// and the new projects are added to the database.
func handler(req events.APIGatewayProxyRequest) (events.APIGatewayProxyResponse, error) {

	ts := oauth2.StaticTokenSource(&oauth2.Token{
		AccessToken: GITHUB_TOKEN,
	})
	tc := oauth2.NewClient(context.Background(), ts)
	gc := github.NewClient(tc)

	sess := dynamo.CreateDynamoSession()

	githubProjects, _, err := gc.Repositories.ListByOrg(context.Background(), "shivam-909", &github.RepositoryListByOrgOptions{
		Type: "public",
	})

	if err != nil {
		return response.New(err.Error(), http.StatusInternalServerError), nil
	}

	parsedProjects := []projects.Project{}

	for _, gh_project := range githubProjects {
		parsedProjects = append(parsedProjects, projects.FromGithubProject(*gh_project))
	}

	err = dynamo_projects.SyncProjects(sess, parsedProjects)
	if err != nil {
		return response.New(err.Error(), http.StatusInternalServerError), nil
	}

	return response.New("", http.StatusAccepted), nil
}
