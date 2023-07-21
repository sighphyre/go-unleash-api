package api

type FeatureTagsResponse struct {
	Version int          `json:"version"`
	Types   []FeatureTag `json:"tags"`
}

type FeatureTag struct {
	Type  string `json:"type"`
	Value string `json:"value"`
}

type FeatureTagsService struct {
	client *ApiClient
}

func (p *FeatureTagsService) GetAllFeatureTags(featureName string) (*FeatureTagsResponse, *Response, error) {
	req, _ := p.client.newRequest("admin/features/"+featureName+"/tags", "GET", nil)

	var featureTags FeatureTagsResponse

	resp, err := p.client.do(req, &featureTags)
	if err != nil {
		return nil, resp, err
	}
	return &featureTags, resp, err
}

func (p *FeatureTagsService) PutFeatureTags(featureName string, tags []FeatureTag) (*FeatureTagsResponse, *Response, error) {
	req, err := p.client.newRequest("admin/features/"+featureName+"/tags", "PUT", tags)
	if err != nil {
		return nil, nil, err
	}

	var tagsResponse FeatureTagsResponse

	resp, err := p.client.do(req, &tagsResponse)
	if err != nil {
		return nil, resp, err
	}

	return &tagsResponse, resp, err
}
