# harness-openapi-go-client

Go client for the Harness NextGen APIs used by `terraform-provider-harness`. The module path is `github.com/harness/harness-openapi-go-client`. Import the NextGen package as `github.com/harness/harness-openapi-go-client/nextgen`.

The tree was first produced by [swagger-codegen](https://github.com/swagger-api/swagger-codegen) (`GoClientCodegen`). A full regenerate rewrites every model and API file, including hand fixes, so do not run codegen for a single field.

## Add or change a response field

Match the published pipeline-service OpenAPI, then edit only the files for that model. `store_type` on get pipeline is the example.

1. `nextgen/api/swagger.yaml` — add the property on the schema, including `description` and `enum` when the API has one.
2. `nextgen/model_<name>.go` — add the Go field and `json` tag. Use `omitempty` for optional response fields. Put `store_type` next to `git_details`, not inside it. An inline pipeline has a store type and no git details.
3. `nextgen/docs/<Model>.md` — add the property row. These docs are what reviewers check.
4. `nextgen/version_json_test.go` or a test next to the model — unmarshal a sample payload and assert the new field. Also assert that an omitted field stays empty.

`git_details.connector_ref` is already on `GitDetails`. Do not add a second connector field on the get response.

## Check the change

Go 1.19 or newer.

```bash
go test ./nextgen/ -count=1
```

## Release

Versions are tags `v0.0.N` on `master`. `terraform-provider-harness` depends on that tag in `go.mod`. After this repo merges, tag the next patch and bump the provider.

Commit messages:

```text
feat: [PIPE-12345]: Short description
```

GitHub blocks new branches on `harness/harness-openapi-go-client`. Push the branch to your fork and open the pull request from there.

`nextgen/git_push.sh` is leftover codegen. It force-pushes `master`. Do not run it.

## Layout

| Path | What it is |
| --- | --- |
| `nextgen/` | Generated client, models, and per-model docs |
| `nextgen/api/swagger.yaml` | Spec checked into this repo |
| `openapi.yaml` | Older spec snapshot. New pipeline fields go in `nextgen/api/swagger.yaml` |
| `utils/` | Small HTTP helpers |

`nextgen/README.md` is the original codegen text. Use this file.
