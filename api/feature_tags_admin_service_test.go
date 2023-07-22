package api

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/philips-labs/go-unleash-api/mocks"
)

type args struct {
	featureName string
	tags        []FeatureTag
}

var (
	featureTagsService *FeatureTagsService
	httpResponseMocks  map[string]*http.Response
)

func init() {
	featureTagsService = &FeatureTagsService{
		client: &ApiClient{
			client:    &mocks.MockClient{},
			apiUrl:    &url.URL{Path: "local"},
			authToken: "myToken",
		},
	}

	httpResponseMocks = make(map[string]*http.Response)
}

func TestFeatureTagsService_GetAllFeatureTags(t *testing.T) {
	httpResponseMocks["success"] = createHttpResponseMock(http.StatusOK, `{
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
	}`, http.MethodGet)
	httpResponseMocks["nocontent"] = createHttpResponseMock(http.StatusOK, `{
		"version": 1,
		"tags": []
	}`, http.MethodGet)
	httpResponseMocks["badrequest"] = createHttpResponseMock(http.StatusBadRequest, `{
		"id": "9c40958a-daac-400e-98fb-3bb438567008",
		"name": "ValidationError",
		"message": "The request payload you provided doesn't conform to the schema. The .parameters property should be object. You sent []."
	}`, http.MethodGet)

	scenarios := []struct {
		name            string
		p               *FeatureTagsService
		args            args
		mockedResponse  *http.Response
		wantFeatureTags *FeatureTagsResponse
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
			&FeatureTagsResponse{
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
			&FeatureTagsResponse{
				Version: 1,
				Types:   []FeatureTag{},
			},
			&Response{Response: httpResponseMocks["nocontent"]},
			false,
		},
		{
			"ReturnsBadRequest",
			featureTagsService,
			args{
				featureName: "",
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
			bodyResponse, response, err := tt.p.GetAllFeatureTags(tt.args.featureName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FeatureTagsService.GetAllFeatureTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(bodyResponse, tt.wantFeatureTags) {
				t.Errorf("FeatureTagsService.GetAllFeatureTags() bodyResponse = %v, want %v", bodyResponse, tt.wantFeatureTags)
			}
			if !reflect.DeepEqual(response, tt.wantResponse) {
				t.Errorf("FeatureTagsService.GetAllFeatureTags() response = %v, want %v", response, tt.wantResponse)
			}
		})
	}
}

func TestFeatureTagsService_CreateFeatureTags(t *testing.T) {
	httpResponseMocks["success"] = createHttpResponseMock(http.StatusOK, `{
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
	}`, http.MethodPost)
	httpResponseMocks["notfound"] = createHttpResponseMock(http.StatusNotFound, `{
		"id": "9c40958a-daac-400e-98fb-3bb438567008",
		"name": "ValidationError",
		"message": "The request payload you provided doesn't conform to the schema. The .parameters property should be object. You sent []."
	}`, http.MethodPost)
	httpResponseMocks["badrequest"] = createHttpResponseMock(http.StatusBadRequest, `{
		"id": "9c40958a-daac-400e-98fb-3bb438567008",
		"name": "ValidationError",
		"message": "The request payload you provided doesn't conform to the schema. The .parameters property should be object. You sent []."
	}`, http.MethodPost)
	scenarios := []struct {
		name            string
		p               *FeatureTagsService
		args            args
		mockedResponse  *http.Response
		wantFeatureTags *FeatureTagsResponse
		wantResponse    *Response
		wantErr         bool
	}{
		{
			"ReturnsFeatureTags",
			featureTagsService,
			args{
				featureName: "MyToggle",
				tags: []FeatureTag{
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
			httpResponseMocks["success"],
			&FeatureTagsResponse{
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
			"ReturnsNotFound",
			featureTagsService,
			args{
				featureName: "UnknownToggle",
				tags:        []FeatureTag{},
			},
			httpResponseMocks["notfound"],
			nil,
			&Response{Response: httpResponseMocks["notfound"]},
			true,
		},
		{
			"ReturnsBadRequest",
			featureTagsService,
			args{
				featureName: "UnknownToggle",
				tags:        []FeatureTag{},
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
			bodyResponse, response, err := tt.p.CreateFeatureTags(tt.args.featureName, tt.args.tags)
			if (err != nil) != tt.wantErr {
				t.Errorf("FeatureTagsService.CreateFeatureTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(bodyResponse, tt.wantFeatureTags) {

				t.Errorf("FeatureTagsService.CreateFeatureTags() bodyResponse = %v, want %v", bodyResponse, tt.wantFeatureTags)
			}
			if !reflect.DeepEqual(response, tt.wantResponse) {
				t.Errorf("FeatureTagsService.CreateFeatureTags() response = %v, want %v", response, tt.wantResponse)
			}
		})
	}
}

func TestFeatureTagsService_DeleteFeatureTags(t *testing.T) {
	httpResponseMocks["success"] = createHttpResponseMock(http.StatusOK, `nil`, http.MethodDelete)
	httpResponseMocks["notfound"] = createHttpResponseMock(http.StatusNotFound, `nil`, http.MethodDelete)

	scenarios := []struct {
		name           string
		p              *FeatureTagsService
		args           args
		mockedResponse *http.Response
		wantResponse   *Response
		wantErr        bool
	}{
		{
			"ReturnsOk",
			featureTagsService,
			args{
				featureName: "MyToggle",
				tags: []FeatureTag{
					{
						Type:  "simple",
						Value: "product1",
					},
				},
			},
			httpResponseMocks["success"],
			nil,
			false,
		},
		{
			"ReturnsNotFound",
			featureTagsService,
			args{
				featureName: "UnknownToggle",
				tags: []FeatureTag{
					{
						Type: "simple",
					},
				},
			},
			httpResponseMocks["notfound"],
			&Response{Response: httpResponseMocks["notfound"]},
			true,
		},
	}
	for _, tt := range scenarios {
		mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
			return tt.mockedResponse, nil
		}
		t.Run(tt.name, func(t *testing.T) {
			response, err := tt.p.DeleteFeatureTags(tt.args.featureName, tt.args.tags[0])
			if (err != nil) != tt.wantErr {
				t.Errorf("FeatureTagsService.DeleteFeatureTags() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(response, tt.wantResponse) {
				t.Errorf("FeatureTagsService.DeleteFeatureTags() response = %v, want %v", response, tt.wantResponse)
			}
		})
	}
}
