package nextgen

import (
	"encoding/json"
	"strings"
	"testing"
)

func TestPipelineRequestBodyIncludesVersionWhenSet(t *testing.T) {
	createWith, err := json.Marshal(PipelineCreateRequestBody{
		Identifier: "id", Name: "name", PipelineYaml: "yaml", Version: "1",
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(createWith), `"version":"1"`) {
		t.Fatalf("create body missing version: %s", createWith)
	}

	createWithout, err := json.Marshal(PipelineCreateRequestBody{
		Identifier: "id", Name: "name", PipelineYaml: "yaml",
	})
	if err != nil {
		t.Fatal(err)
	}
	if strings.Contains(string(createWithout), `"version"`) {
		t.Fatalf("empty version should be omitted: %s", createWithout)
	}

	updateWith, err := json.Marshal(PipelineUpdateRequestBody{
		Identifier: "id", Name: "name", PipelineYaml: "yaml", Version: "1",
	})
	if err != nil {
		t.Fatal(err)
	}
	if !strings.Contains(string(updateWith), `"version":"1"`) {
		t.Fatalf("update body missing version: %s", updateWith)
	}
}

func TestPipelineGetResponseReadsStoreType(t *testing.T) {
	var remote PipelineGetResponseBody
	if err := json.Unmarshal([]byte(`{"store_type":"REMOTE","git_details":{"connector_ref":"account.conn","branch_name":"main"}}`), &remote); err != nil {
		t.Fatal(err)
	}
	if remote.StoreType != "REMOTE" {
		t.Fatalf("store_type = %q", remote.StoreType)
	}
	if remote.GitDetails == nil || remote.GitDetails.ConnectorRef != "account.conn" {
		t.Fatalf("connector_ref = %#v", remote.GitDetails)
	}

	var inline PipelineGetResponseBody
	if err := json.Unmarshal([]byte(`{"store_type":"INLINE"}`), &inline); err != nil {
		t.Fatal(err)
	}
	if inline.StoreType != "INLINE" {
		t.Fatalf("store_type = %q", inline.StoreType)
	}
	if inline.GitDetails != nil {
		t.Fatalf("inline git_details = %#v", inline.GitDetails)
	}
}
