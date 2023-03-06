package api

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
)

type ProjectDetails struct {
	Name         string   `json:"name"`
	Description  string   `json:"description"`
	Health       int      `json:"health"`
	UpdatedAt    string   `json:"updatedAt"`
	Environments []string `json:"environments"`
}

type Project struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
}

type ProjectInfo struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}

type ErrorResponse struct {
	Error struct {
		Name    string `json:"name"`
		Message string `json:"message"`
	} `json:"error"`
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

	if resp == nil {
		return nil, nil, errors.New("response is nil")
	}

	return &project, resp, nil
}

func (p *ProjectsService) CreateProject(project Project) (*ProjectInfo, *Response, error) {
	req, _ := p.client.newRequest("admin/projects", "POST", project)

	var projectCreate ProjectInfo

	resp, err := p.client.do(req, &projectCreate)
	if err != nil {
		return nil, resp, err
	}

	return &projectCreate, resp, err
}

func (p *ProjectsService) UpdateProject(projectId string, project Project) (*ProjectInfo, *Response, error) {
	if projectId == "" {
		return nil, nil, ErrRequiredParam("projectId")
	}
	req, err := p.client.newRequest("admin/projects/"+projectId, "PUT", project)
	if err != nil {
		return nil, nil, err
	}

	var projectUpdate ProjectInfo

	resp, err := p.client.do(req, &projectUpdate)
	if err != nil {
		return nil, resp, err
	}

	return &projectUpdate, resp, err
}

func (p *ProjectsService) DeleteProject(projectId string) (*Response, error) {
	if projectId == "" {
		return nil, ErrRequiredParam("projectID")
	}
	req, err := p.client.newRequest("DELETE", "admin/projects/"+projectId, nil)
	if err != nil {
		return nil, err
	}
	var deleteResponse Response
	resp, err := p.client.do(req, &deleteResponse)
	if err != nil {
		return nil, err
	}
	if resp == nil {
		return nil, errors.New("response is nil")
	}
	if resp.StatusCode != http.StatusNoContent {
		var errorResponse ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s: %s", errorResponse.Error.Name, errorResponse.Error.Message)
	}
	return &deleteResponse, nil
}
