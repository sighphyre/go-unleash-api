package api

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/philips-labs/go-unleash-api/mocks"
)

var (
	featureTypesService *FeatureTypesService
)

func init() {
	featureTypesService = &FeatureTypesService{
		client: &ApiClient{
			client:    &mocks.MockClient{},
			apiUrl:    &url.URL{Path: "local"},
			authToken: "myToken",
		},
	}

}

func TestFeatureTypesService_GetAllFeatureTypes(t *testing.T) {
	httpResponseMocks := make(map[string]*http.Response)
	httpResponseMocks["success"] = createHttpResponseMock(200, `{"version":1,"types":[{"id":"release","name":"Release"}]}`, "GET")
	httpResponseMocks["badrequest"] = createHttpResponseMock(404, `{"name":"BadRequest"`, "GET")
	tests := []struct {
		name             string
		p                *FeatureTypesService
		mockedResponse   *http.Response
		wantFeatureTypes *AllFeatureTypesResponse
		wantResponse     *Response
		wantErr          bool
	}{
		{
			"ReturnsFeatureTypes",
			featureTypesService,
			httpResponseMocks["success"],
			&AllFeatureTypesResponse{
				Version: 1,
				Types: []FeatureType{
					{
						ID:   "release",
						Name: "Release",
					},
				},
			},
			&Response{Response: httpResponseMocks["success"]},
			false,
		},
		{
			"ReturnsError",
			featureTypesService,
			httpResponseMocks["badrequest"],
			nil,
			&Response{Response: httpResponseMocks["badrequest"]},
			true,
		},
	}
	for _, tt := range tests {
		mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
			return tt.mockedResponse, nil
		}
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.p.GetAllFeatureTypes()
			if (err != nil) != tt.wantErr {
				t.Errorf("FeatureTypesService.GetAllFeatureTypes() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantFeatureTypes) {
				t.Errorf("FeatureTypesService.GetAllFeatureTypes() got = %v, want %v", got, tt.wantFeatureTypes)
			}
			if !reflect.DeepEqual(got1, tt.wantResponse) {
				t.Errorf("FeatureTypesService.GetAllFeatureTypes() got1 = %v, want %v", got1, tt.wantResponse)
			}
		})
	}
}
