package projects

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
