package api

type ProjectDetails struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Health       int      `json:"health"`
	UpdatedAt    string   `json:"updatedAt"`
	Environments []string `json:"environments"`
}

type ProjectsService struct {
	client *ApiClient
}

func (p *ProjectsService) GetProjectById(projectId string) (*ProjectDetails, *Response, error) {
	req, _ := p.client.newRequest("admin/projects/"+projectId, "GET", nil)

	var project ProjectDetails

	resp, err := p.client.do(req, &project)
	if err != nil {
		return nil, resp, err
	}

	return &project, resp, err
}
