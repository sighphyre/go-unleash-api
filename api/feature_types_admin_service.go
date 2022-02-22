package api

type AllFeatureTypesResponse struct {
	Version int           `json:"version"`
	Types   []FeatureType `json:"types"`
}

type FeatureType struct {
	ID           string `json:"id,omitempty"`
	Name         string `json:"name"`
	Description  string `json:"description"`
	LifetimeDays int    `json:"lifetimeDays"`
}

type FeatureTypesService struct {
	client *ApiClient
}

func (p *FeatureTypesService) GetAllFeatureTypes() (*AllFeatureTypesResponse, *Response, error) {
	req, _ := p.client.newRequest("admin/feature-types", "GET", nil)

	var featureTypes AllFeatureTypesResponse

	resp, err := p.client.do(req, &featureTypes)
	if err != nil {
		return nil, resp, err
	}
	return &featureTypes, resp, err
}
