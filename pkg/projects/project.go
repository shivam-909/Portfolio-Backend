package projects

import (
	"encoding/json"

	"github.com/aws/aws-lambda-go/events"
)

type Demo struct {
	Exists   bool              `json:"exists"`
	Metadata map[string]string `json:"metadata"`
}

type Project struct {
	ID          int32    `json:"id"`
	Title       string   `json:"title"`
	Description string   `json:"description"`
	Created     int32    `json:"created"`
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
