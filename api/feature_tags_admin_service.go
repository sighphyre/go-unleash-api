package api

type UpdateFeatureTagsBody struct {
	AddedTags   []FeatureTag `json:"addedTags"`
	RemovedTags []FeatureTag `json:"removedTags"`
}

type FeatureTagsResponse struct {
	Version int          `json:"version"`
	Tags    []FeatureTag `json:"tags"`
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

func (p *FeatureTagsService) CreateFeatureTags(featureName string, tag FeatureTag) (*FeatureTag, *Response, error) {
	req, err := p.client.newRequest("admin/features/"+featureName+"/tags", "POST", tag)
	if err != nil {
		return nil, nil, err
	}

	var tagsResponse FeatureTag

	resp, err := p.client.do(req, &tagsResponse)
	if err != nil {
		return nil, resp, err
	}

	return &tagsResponse, resp, err
}

func (p *FeatureTagsService) UpdateFeatureTags(featureName string, addedTags []FeatureTag, removedTags []FeatureTag) (*FeatureTagsResponse, *Response, error) {
	updateFeatureTagsBody := UpdateFeatureTagsBody{
		AddedTags:   addedTags,
		RemovedTags: removedTags,
	}

	req, err := p.client.newRequest("admin/features/"+featureName+"/tags", "PUT", updateFeatureTagsBody)
	if err != nil {
		return nil, nil, err
	}

	var updatedTagsResponse FeatureTagsResponse

	resp, err := p.client.do(req, &updatedTagsResponse)
	if err != nil {
		return nil, resp, err
	}

	return &updatedTagsResponse, resp, err
}

func (p *FeatureTagsService) DeleteFeatureTags(featureName string, tag FeatureTag) (*Response, error) {
	req, err := p.client.newRequest("admin/features/"+featureName+"/tags/"+tag.Type+"/"+tag.Value, "DELETE", nil)
	if err != nil {
		return nil, err
	}
	resp, err := p.client.do(req, nil)
	if err != nil {
		return resp, err
	}
	return nil, nil
}
