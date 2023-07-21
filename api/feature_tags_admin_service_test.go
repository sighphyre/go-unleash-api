package api

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/philips-labs/go-unleash-api/mocks"
)

var (
	featureTagsService *FeatureTagsService
)

func init() {
	featureTagsService = &FeatureTagsService{
		client: &ApiClient{
			client:    &mocks.MockClient{},
			apiUrl:    &url.URL{Path: "local"},
			authToken: "myToken",
		},
	}

}

func TestFeatureTagsService_GetAllFeatureTags(t *testing.T) {
	httpResponseMocks := make(map[string]*http.Response)
	httpResponseMocks["success"] = createHttpResponseMock(200, `{
		"version": 1,
		"tags": [
			{
				"value": "feature1",
				"type": "simple"
			},
			{
				"value": "product1",
				"type": "simple"
			}
		]
	}`, "GET")
	httpResponseMocks["nocontent"] = createHttpResponseMock(200, `{
		"version": 1,
		"tags": []
	}`, "GET")
	httpResponseMocks["badrequest"] = createHttpResponseMock(404, `{"name":"BadRequest"`, "GET")
	type args struct {
		featureName string
	}
	scenarios := []struct {
		name            string
		p               *FeatureTagsService
		args            args
		mockedResponse  *http.Response
		wantFeatureTags *AllFeatureTagsResponse
		wantResponse    *Response
		wantErr         bool
	}{
		{
			"ReturnsFeatureTags",
			featureTagsService,
			args{
				featureName: "MyToggle",
			},
			httpResponseMocks["success"],
			&AllFeatureTagsResponse{
				Version: 1,
				Types: []FeatureTag{
					{
						Type:  "simple",
						Value: "feature1",
					},
					{
						Type:  "simple",
						Value: "product1",
					},
				},
			},
			&Response{Response: httpResponseMocks["success"]},
			false,
		},
		{
			"ReturnsEmptyFeatureTags",
			featureTagsService,
			args{
				featureName: "MyToggleEmptyTags",
			},
			httpResponseMocks["nocontent"],
			&AllFeatureTagsResponse{
				Version: 1,
				Types:   []FeatureTag{},
			},
			&Response{Response: httpResponseMocks["nocontent"]},
			false,
		},
		{
			"ReturnsError",
			featureTagsService,
			args{
				featureName: "UnknownToggle",
			},
			httpResponseMocks["badrequest"],
			nil,
			&Response{Response: httpResponseMocks["badrequest"]},
			true,
		},
	}
	for _, tt := range scenarios {
		mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
			return tt.mockedResponse, nil
		}
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.p.GetAllFeatureTags(tt.args.featureName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FeatureTagsService.GetAllFeatureTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantFeatureTags) {
				t.Errorf("FeatureTagsService.GetAllFeatureTags() got = %v, want %v", got, tt.wantFeatureTags)
			}
			if !reflect.DeepEqual(got1, tt.wantResponse) {
				t.Errorf("FeatureTagsService.GetAllFeatureTags() got1 = %v, want %v", got1, tt.wantResponse)
			}
		})
	}
}
