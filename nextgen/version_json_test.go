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
