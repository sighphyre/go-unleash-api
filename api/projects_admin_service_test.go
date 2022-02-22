package api

import (
	"net/http"
	"net/url"
	"reflect"
	"testing"

	"github.com/philips-labs/go-unleash-api/mocks"
)

var (
	projectsService *ProjectsService
)

func init() {
	projectsService = &ProjectsService{
		client: &ApiClient{
			client:    &mocks.MockClient{},
			apiUrl:    &url.URL{Path: "local"},
			authToken: "myToken",
		},
	}

}

func TestProjectsService_GetProjectById(t *testing.T) {
	httpResponseMocks := make(map[string]*http.Response)
	httpResponseMocks["success"] = createHttpResponseMock(200, `{"name":"Default","description":"Default project","environments":["development","production"],"features":[]}`, "GET")
	httpResponseMocks["notfound"] = createHttpResponseMock(404, `{"name":"NotFoundError"`, "GET")
	type args struct {
		projectId string
	}
	tests := []struct {
		name           string
		p              *ProjectsService
		args           args
		mockedResponse *http.Response
		wantProject    *ProjectDetails
		wantResponse   *Response
		wantErr        bool
	}{
		{
			"ReturnsProjectDetails",
			projectsService,
			args{
				projectId: "default",
			},
			httpResponseMocks["success"],
			&ProjectDetails{
				Name:         "Default",
				Description:  "Default project",
				Environments: []string{"development", "production"},
			},
			&Response{Response: httpResponseMocks["success"]},
			false,
		},
		{
			"ReturnsError",
			projectsService,
			args{
				projectId: "foo",
			},
			httpResponseMocks["notfound"],
			nil,
			&Response{Response: httpResponseMocks["notfound"]},
			true,
		},
	}
	for _, tt := range tests {
		mocks.GetDoFunc = func(*http.Request) (*http.Response, error) {
			return tt.mockedResponse, nil
		}
		t.Run(tt.name, func(t *testing.T) {
			got, got1, err := tt.p.GetProjectById(tt.args.projectId)
			if (err != nil) != tt.wantErr {
				t.Errorf("ProjectsService.GetProjectById() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.wantProject) {
				t.Errorf("ProjectsService.GetProjectById() got = %v, want %v", got, tt.wantProject)
			}
			if !reflect.DeepEqual(got1, tt.wantResponse) {
				t.Errorf("ProjectsService.GetProjectById() got1 = %v, want %v", got1, tt.wantResponse)
			}
		})
	}
}
