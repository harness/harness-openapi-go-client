package nextgen

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"

	"github.com/antihax/optional"
)

func TestTemplateGitMetadataUpdate(t *testing.T) {
	requestBody := []TemplateUpdateGitMetadataRequest{{
		Version: "v1",
		GitDetails: &GitDetailsMetadata{
			FilePath:     ".harness/templates/example.yaml",
			RepoName:     "example-repo",
			ConnectorRef: "account.new_connector",
		},
	}}

	tests := []struct {
		name string
		path string
	}{
		{name: "account", path: "/v1/templates/git-metadata/example"},
		{name: "org", path: "/v1/orgs/example_org/templates/git-metadata/example"},
		{name: "project", path: "/v1/orgs/example_org/projects/example_project/templates/git-metadata/example"},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
				if r.Method != http.MethodPut {
					t.Errorf("method = %s, want PUT", r.Method)
				}
				if r.URL.Path != tt.path {
					t.Errorf("path = %s, want %s", r.URL.Path, tt.path)
				}
				if got := r.Header.Get("Harness-Account"); got != "example_account" {
					t.Errorf("Harness-Account = %q, want example_account", got)
				}
				if got := r.Header.Get("x-api-key"); got != "test-key" {
					t.Errorf("x-api-key = %q, want test-key", got)
				}
				if got := r.Header.Get("Content-Type"); got != "application/json" {
					t.Errorf("Content-Type = %q, want application/json", got)
				}

				var gotBody []TemplateUpdateGitMetadataRequest
				if err := json.NewDecoder(r.Body).Decode(&gotBody); err != nil {
					t.Errorf("decode request body: %v", err)
				}
				if !reflect.DeepEqual(gotBody, requestBody) {
					t.Errorf("request body = %#v, want %#v", gotBody, requestBody)
				}

				w.Header().Set("Content-Type", "application/json")
				_, _ = w.Write([]byte(`{"status":true}`))
			}))
			defer server.Close()

			cfg := NewConfiguration()
			cfg.BasePath = server.URL
			cfg.HTTPClient.RetryMax = 0
			client := NewAPIClient(cfg)
			ctx := context.WithValue(context.Background(), ContextAPIKey, APIKey{Key: "test-key"})

			var (
				response     TemplateUpdateGitMetadataResponse
				httpResponse *http.Response
				err          error
			)
			switch tt.name {
			case "account":
				response, httpResponse, err = client.AccountTemplateApi.UpdateGitMetadataDetails(ctx, "example", &AccountTemplateApiUpdateGitMetadataDetailsOpts{
					Body:           optional.NewInterface(requestBody),
					HarnessAccount: optional.NewString("example_account"),
				})
			case "org":
				response, httpResponse, err = client.OrgTemplateApi.UpdateGitMetadataDetailsOrg(ctx, "example", "example_org", &OrgTemplateApiUpdateGitMetadataDetailsOrgOpts{
					Body:           optional.NewInterface(requestBody),
					HarnessAccount: optional.NewString("example_account"),
				})
			case "project":
				response, httpResponse, err = client.ProjectTemplateApi.UpdateGitMetadataDetailsProject(ctx, "example", "example_org", "example_project", &ProjectTemplateApiUpdateGitMetadataDetailsProjectOpts{
					Body:           optional.NewInterface(requestBody),
					HarnessAccount: optional.NewString("example_account"),
				})
			}

			if err != nil {
				t.Fatalf("update Git metadata: %v", err)
			}
			if httpResponse.StatusCode != http.StatusOK {
				t.Fatalf("status = %d, want 200", httpResponse.StatusCode)
			}
			if !response.Status {
				t.Fatal("response status = false, want true")
			}
		})
	}
}

func TestTemplateGitMetadataUpdateReturnsAPIError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, _ *http.Request) {
		http.Error(w, "update failed", http.StatusInternalServerError)
	}))
	defer server.Close()

	cfg := NewConfiguration()
	cfg.BasePath = server.URL
	cfg.HTTPClient.RetryMax = 0
	client := NewAPIClient(cfg)

	_, _, err := client.AccountTemplateApi.UpdateGitMetadataDetails(context.Background(), "example", &AccountTemplateApiUpdateGitMetadataDetailsOpts{
		Body: optional.NewInterface([]TemplateUpdateGitMetadataRequest{{Version: "v1"}}),
	})
	if err == nil {
		t.Fatal("update Git metadata error = nil, want API error")
	}
}
