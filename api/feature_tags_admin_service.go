package api

type AllFeatureTagsResponse struct {
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

func (p *FeatureTagsService) GetAllFeatureTags(featureName string) (*AllFeatureTagsResponse, *Response, error) {
	req, _ := p.client.newRequest("admin/features/"+featureName+"/tags", "GET", nil)

	var featureTags AllFeatureTagsResponse

	resp, err := p.client.do(req, &featureTags)
	if err != nil {
		return nil, resp, err
	}
	return &featureTags, resp, err
}
