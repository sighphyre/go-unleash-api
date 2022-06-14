package api

import (
	"bytes"
)

type UserDetails struct {
	Id         int    `json:"id"`
	Name       string `json:"name,omitempty"`
	Username   string `json:"username,omitempty"`
	Email      string `json:"email,omitempty"`
	ImageUrl   string `json:"imageUrl,omitempty"`
	CreatedAt  string `json:"createdAt,omitempty"`
	InviteLink string `json:"inviteLink,omitempty"`
	EmailSent  bool   `json:"emailSent,omitempty"`
	RootRole   int    `json:"rootRole,omitempty"`
}

type User struct {
	Name      string `json:"name"`
	Username  string `json:"username"`
	Email     string `json:"email"`
	RootRole  int    `json:"rootRole"`
	SendEmail bool   `json:"sendEmail"`
}

type UsersService struct {
	client *ApiClient
}

func (p *UsersService) GetUserById(userId string) (*UserDetails, *Response, error) {
	if userId == "" {
		return nil, nil, ErrRequiredParam("userId")
	}
	req, _ := p.client.newRequest("admin/user-admin/"+userId, "GET", nil)

	var user UserDetails

	resp, err := p.client.do(req, &user)
	if err != nil {
		return nil, resp, err
	}
	return &user, resp, err
}

func (p *UsersService) CreateUser(user User) (*UserDetails, *Response, error) {
	req, _ := p.client.newRequest("admin/user-admin", "POST", user)

	var userDetails UserDetails

	resp, err := p.client.do(req, &userDetails)
	if err != nil {
		return nil, resp, err
	}
	return &userDetails, resp, err
}

func (p *UsersService) UpdateUser(userId string, user User) (*UserDetails, *Response, error) {
	if userId == "" {
		return nil, nil, ErrRequiredParam("userId")
	}
	req, err := p.client.newRequest("admin/user-admin/"+userId, "PUT", user)
	if err != nil {
		return nil, nil, err
	}

	var updatedUser UserDetails

	resp, err := p.client.do(req, &updatedUser)
	if err != nil {
		return nil, resp, err
	}
	return &updatedUser, resp, err
}

func (p *UsersService) DeleteUser(userId string) (bool, *Response, error) {
	if userId == "" {
		return false, nil, ErrRequiredParam("userId")
	}
	req, _ := p.client.newRequest("admin/user-admin/"+userId, "DELETE", nil)

	var deleteResponse bytes.Buffer

	resp, err := p.client.do(req, &deleteResponse)
	if resp == nil {
		return false, resp, err
	}
	return true, resp, nil
}

func (p *UsersService) SearchUser(query string) (*[]UserDetails, *Response, error) {
	if query == "" {
		return nil, nil, ErrRequiredParam("query")
	}
	req, _ := p.client.newRequest("admin/user-admin/search?q="+query, "GET", nil)

	var users []UserDetails

	resp, err := p.client.do(req, &users)
	if err != nil {
		return nil, resp, err
	}
	return &users, resp, err
}
