package projects

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
	"github.com/google/go-github/v51/github"
)

type Demo struct {
	Exists   bool              `json:"exists"`
	Metadata map[string]string `json:"metadata"`
}

type Project struct {
	ID          int64    `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Created     int64    `json:"created"`
	LinkToRepo  string   `json:"link_to_repo"`
	ImageURL    string   `json:"image_url"`
	Tags        []string `json:"tags"`
	Demo        Demo     `json:"demo"`
}

func Parse(req events.APIGatewayProxyRequest) (Project, error) {
	var project Project
	err := json.Unmarshal([]byte(req.Body), &project)
	if err != nil {
		return Project{}, err
	}
	return project, nil
}

// Create a method on Project called String that turns it into a JSON string and escapes it for HTML.
func (p Project) String() string {
	json, err := json.Marshal(p)
	if err != nil {
		return ""
	}
	return string(json)
}

func FromGithubProject(ghp github.Repository) Project {
	return Project{
		ID:          ghp.GetID(),
		Title:       ghp.GetName(),
		Description: ghp.GetDescription(),
		Created:     ghp.GetCreatedAt().Unix(),
		LinkToRepo:  ghp.GetHTMLURL(),
		ImageURL:    "",
		Tags:        []string{},
		Demo: Demo{
			Exists:   false,
			Metadata: map[string]string{},
		},
	}
}
