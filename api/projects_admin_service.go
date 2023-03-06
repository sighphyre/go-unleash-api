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

type CreateProjectResponse struct {
	Id          string `json:"id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"createdAt"`
}

type AddUserRoleResponse struct {
	UserId    int    `json:"userId"`
	ProjectId string `json:"projectId"`
	RoleId    int    `json:"roleId"`
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

func (p *ProjectsService) CreateProject(project Project) (*CreateProjectResponse, *Response, error) {
	req, _ := p.client.newRequest("admin/projects", "POST", project)

	var projectCreate CreateProjectResponse

	resp, err := p.client.do(req, &projectCreate)
	if err != nil {
		return nil, resp, err
	}

	return &projectCreate, resp, err
}

func (p *ProjectsService) UpdateProject(projectId string, project Project) (*CreateProjectResponse, *Response, error) {
	if projectId == "" {
		return nil, nil, ErrRequiredParam("projectId")
	}
	req, err := p.client.newRequest("admin/projects/"+projectId, "PUT", project)
	if err != nil {
		return nil, nil, err
	}

	var projectUpdate CreateProjectResponse

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

func (p *ProjectsService) AddUserRole(userId int, projectId string, roleId int) (*AddUserRoleResponse, *Response, error) {
	if projectId == "" {
		return nil, nil, ErrRequiredParam("projectID")
	}
	if userId == 0 {
		return nil, nil, ErrRequiredParam("userId")
	}
	if roleId == 0 {
		return nil, nil, ErrRequiredParam("roleId")
	}

	path := fmt.Sprintf("admin/projects/%s/users/%d/roles/%d", projectId, userId, roleId)
	req, err := p.client.newRequest("PUT", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var addRoleResponse AddUserRoleResponse
	resp, err := p.client.do(req, &addRoleResponse)
	if err != nil {
		return nil, nil, err
	}
	if resp == nil {
		return nil, nil, errors.New("response is nil")
	}
	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, nil, err
		}
		return nil, nil, fmt.Errorf("%s: %s", errorResponse.Error.Name, errorResponse.Error.Message)
	}
	return &addRoleResponse, resp, err
}

func (p *ProjectsService) UpdateUserRole(projectId string, userId int, roleId int) (*AddUserRoleResponse, *Response, error) {
	if projectId == "" {
		return nil, nil, ErrRequiredParam("projectID")
	}
	if userId == 0 {
		return nil, nil, ErrRequiredParam("userID")
	}
	if roleId == 0 {
		return nil, nil, ErrRequiredParam("roleID")
	}

	path := fmt.Sprintf("admin/projects/%s/users/%d/roles/%d", projectId, userId, roleId)
	req, err := p.client.newRequest("POST", path, nil)
	if err != nil {
		return nil, nil, err
	}

	var updateRoleResponse AddUserRoleResponse
	resp, err := p.client.do(req, &updateRoleResponse)
	if err != nil {
		return nil, nil, err
	}
	if resp == nil {
		return nil, nil, errors.New("response is nil")
	}
	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, nil, err
		}
		return nil, nil, fmt.Errorf("%s: %s", errorResponse.Error.Name, errorResponse.Error.Message)
	}
	return &updateRoleResponse, resp, nil
}

func (p *ProjectsService) DeleteUserRole(projectId string, userId int, roleId int) (*Response, error) {
	if projectId == "" {
		return nil, ErrRequiredParam("projectID")
	}
	if userId == 0 {
		return nil, ErrRequiredParam("userID")
	}
	if roleId == 0 {
		return nil, ErrRequiredParam("roleID")
	}

	path := fmt.Sprintf("admin/projects/%s/users/%d/roles/%d", projectId, userId, roleId)
	req, err := p.client.newRequest("DELETE", path, nil)
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

	if resp.StatusCode != http.StatusOK {
		var errorResponse ErrorResponse
		err := json.NewDecoder(resp.Body).Decode(&errorResponse)
		if err != nil {
			return nil, err
		}
		return nil, fmt.Errorf("%s: %s", errorResponse.Error.Name, errorResponse.Error.Message)
	}

	return &deleteResponse, nil
}
