//go:build tools
// +build tools

package tools

import (
	_ "github.com/swaggo/swag/cmd/swag"
	_ "github.com/urfave/cli/v2"
	_ "sigs.k8s.io/yaml"
)
