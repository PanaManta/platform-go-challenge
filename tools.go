//go:build tools
// +build tools

package tools

import (
	_ "github.com/swaggo/swag/cmd/swag" // Swagger generator tool
	_ "github.com/urfave/cli/v2"        // Dependency for swag tool
	_ "sigs.k8s.io/yaml"                // Dependency for swag tool
)
