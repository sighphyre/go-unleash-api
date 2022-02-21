package api

import (
	"bytes"
	"io/ioutil"
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/jdelucaa/go-unleash-api/mocks"
)

var (
	featureTogglesService *FeatureTogglesService
)

func init() {
	featureTogglesService = &FeatureTogglesService{
		client: &ApiClient{
			client:    &mocks.MockClient{},
			apiUrl:    &url.URL{Path: "local"},
			authToken: "myToken",
		},
	}

}

func TestFeatureTogglesService_GetFeatureByName(t *testing.T) {
	GetFeatureByNameMocks := make(map[string]*http.Response)
	GetFeatureByNameMocks["success"] = &http.Response{
		StatusCode: 200,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"name":"MyToggle","project":"default"}`))),
	}
	GetFeatureByNameMocks["notfound"] = &http.Response{
		StatusCode: 404,
		Body:       ioutil.NopCloser(bytes.NewReader([]byte(`{"name":"NotFoundError"`))),
		Request:    &http.Request{Method: "GET", RequestURI: "local"},
	}
	type args struct {
		projectId   string
		featureName string
	}
	tests := []struct {
		name           string
		p              *FeatureTogglesService
		args           args
		mockedResponse *http.Response
		wantFeature    *FeatureToggle
		wantResponse   *Response
		wantErr        bool
	}{
		{
			"ReturnsFeatureByName",
			featureTogglesService,
			args{
				projectId:   "default",
				featureName: "MyToggle",
			},
			GetFeatureByNameMocks["success"],
			&FeatureToggle{
				Project: "default",
				Name:    "MyToggle",
			},
			&Response{Response: GetFeatureByNameMocks["success"]},
			false,
		},
		{
			"ReturnsNotFoundError",
			featureTogglesService,
			args{
				projectId:   "default",
				featureName: "foo",
			},
			GetFeatureByNameMocks["notfound"],
			nil,
			&Response{Response: GetFeatureByNameMocks["notfound"]},
			true,
		},
	}
	for _, tt := range tests {
		mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
			return tt.mockedResponse, nil
		}
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.p.GetFeatureByName(tt.args.projectId, tt.args.featureName)
			if (err != nil) != tt.wantErr {
				t.Errorf("FeatureTogglesService.GetFeatureByName() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantFeature) {
				t.Errorf("FeatureTogglesService.GetFeatureByName() got = %v, want %v", got, tt.wantFeature)
			}
			if !reflect.DeepEqual(got1, tt.wantResponse) {
				t.Errorf("FeatureTogglesService.GetFeatureByName() got1 = %v, want %v", got1, tt.wantResponse)
			}
		})
	}
}
