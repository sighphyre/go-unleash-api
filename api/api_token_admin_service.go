package api

import (
	"bytes"
)

type AllApiTokensResponse struct {
	Tokens []ApiToken `json:"tokens"`
}

type ApiToken struct {
	Secret      string   `json:"secret,omitempty"`
	Username    string   `json:"username"`
	Type        string   `json:"type"`
	Environment string   `json:"environment,omitempty"`
	Projects    []string `json:"projects,omitempty"`
	ExpiresAt   string   `json:"expiresAt,omitempty"`
	CreatedAt   string   `json:"createdAt,omitempty"`
}

type ApiTokenService struct {
	client *ApiClient
}

func (p *ApiTokenService) GetAllApiTokens() (*AllApiTokensResponse, *Response, error) {
	req, _ := p.client.newRequest("admin/api-tokens", "GET", nil)

	var tokens AllApiTokensResponse

	resp, err := p.client.do(req, &tokens)
	if err != nil {
		return nil, resp, err
	}
	return &tokens, resp, err
}

func (p *ApiTokenService) CreateApiToken(token ApiToken) (*ApiToken, *Response, error) {
	req, _ := p.client.newRequest("admin/api-tokens", "POST", token)

	var tokenDetails ApiToken

	resp, err := p.client.do(req, &tokenDetails)
	if err != nil {
		return nil, resp, err
	}
	return &tokenDetails, resp, err
}

func (p *ApiTokenService) UpdateApiToken(secret string, token ApiToken) (bool, *Response, error) {
	if secret == "" {
		return false, nil, ErrRequiredParam("token")
	}
	req, err := p.client.newRequest("admin/api-tokens/"+secret, "PUT", token)
	if err != nil {
		return false, nil, err
	}

	var putResponse bytes.Buffer

	resp, err := p.client.do(req, &putResponse)
	if resp == nil {
		return false, resp, err
	}
	return true, resp, nil
}

func (p *ApiTokenService) DeleteApiToken(secret string) (bool, *Response, error) {
	if secret == "" {
		return false, nil, ErrRequiredParam("secret")
	}
	req, _ := p.client.newRequest("admin/api-tokens/"+secret, "DELETE", nil)

	var deleteResponse bytes.Buffer

	resp, err := p.client.do(req, &deleteResponse)
	if resp == nil {
		return false, resp, err
	}
	return true, resp, nil
}
