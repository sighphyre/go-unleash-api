package api

import (
	"bytes"
)

type AllStrategiesResponse struct {
	Version    int        `json:"version"`
	Strategies []Strategy `json:"strategies"`
}

// FeatureToggle represents a Feature Toggle resource
type Strategy struct {
	ID          string              `json:"id,omitempty"`
	Name        string              `json:"name"`
	DisplayName string              `json:"displayName,omitempty"`
	Description string              `json:"description"`
	Editable    bool                `json:"editable"`
	Deprecated  bool                `json:"deprecated"`
	Parameters  []StrategyParameter `json:"parameters"`
}

type StrategyParameter struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Required    bool   `json:"required"`
	Type        string `json:"type"`
}

type StrategiesService struct {
	client *ApiClient
}

func (p *StrategiesService) CreateStrategy(strategy Strategy) (*Strategy, *Response, error) {
	req, _ := p.client.newRequest("admin/strategies", "POST", strategy)

	var createdStrategy Strategy

	resp, err := p.client.do(req, &createdStrategy)
	if err != nil {
		return nil, resp, err
	}
	return &createdStrategy, resp, err
}

func (p *StrategiesService) UpdateStrategy(strategy Strategy) (*Strategy, *Response, error) {
	req, _ := p.client.newRequest("admin/strategies/"+strategy.Name, "PUT", strategy)

	var updatedStrategy Strategy

	resp, err := p.client.do(req, &updatedStrategy)
	if err != nil {
		return nil, resp, err
	}
	return &updatedStrategy, resp, err
}

func (p *FeatureTogglesService) DeprecateStrategy(strategyName string) (bool, *Response, error) {
	req, err := p.client.newRequest("admin/strategies/"+strategyName+"/deprecate", "POST", nil)
	if err != nil {
		return false, nil, err
	}

	var deprecateResponse bytes.Buffer

	resp, err := p.client.do(req, &deprecateResponse)
	if err != nil {
		return false, resp, err
	}
	return true, resp, nil
}

func (p *FeatureTogglesService) ReactivateStrategy(strategyName string) (bool, *Response, error) {
	req, err := p.client.newRequest("admin/strategies/"+strategyName+"/reactivate", "POST", nil)
	if err != nil {
		return false, nil, err
	}

	var reactivateResponse bytes.Buffer

	resp, err := p.client.do(req, &reactivateResponse)
	if err != nil {
		return false, resp, err
	}
	return true, resp, nil
}

func (p *StrategiesService) GetAllStrategies() (*AllStrategiesResponse, *Response, error) {
	req, err := p.client.newRequest("admin/strategies", "GET", nil)
	if err != nil {
		return nil, nil, err
	}

	var strategies AllStrategiesResponse

	resp, err := p.client.do(req, &strategies)
	if err != nil {
		return nil, resp, err
	}
	return &strategies, resp, err
}

func (p *StrategiesService) GetStrategyByName(strategyName string) (*Strategy, *Response, error) {
	req, err := p.client.newRequest("admin/strategies/"+strategyName, "GET", nil)
	if err != nil {
		return nil, nil, err
	}

	var strategy Strategy

	resp, err := p.client.do(req, &strategy)
	if err != nil {
		return nil, resp, err
	}

	return &strategy, resp, err
}
