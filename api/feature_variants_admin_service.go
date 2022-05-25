package api

type VariantsResponse struct {
	Version  string    `json:"version"`
	Variants []Variant `json:"variants"`
}

type VariantsService struct {
	client *ApiClient
}

func (p *VariantsService) AddVariantsForFeatureToggle(projectId string, featureName string, variants []Variant) (*VariantsResponse, *Response, error) {
	req, err := p.client.newRequest("admin/projects/"+projectId+"/features/"+featureName+"/variants", "PUT", variants)
	if err != nil {
		return nil, nil, err
	}

	var variantsResponse VariantsResponse

	resp, err := p.client.do(req, &variantsResponse)
	if err != nil {
		return nil, resp, err
	}

	return &variantsResponse, resp, err
}
