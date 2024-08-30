//go:build tools
// +build tools

package langfuse

import (
	_ "github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen"
)

//go:generate go run github.com/oapi-codegen/oapi-codegen/v2/cmd/oapi-codegen --config=codegen_config.yaml https://cloud.langfuse.com/generated/api/openapi.yml
