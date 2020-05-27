//args: -Egoimports
//config: linters-settings.goimports.local-prefixes=github.com/golangci/golangci-lint
package goimports

import (
	"errors"
	"fmt"

	"github.com/golangci/golangci-lint/pkg/config"
)

func GoimportsLocalTest() {
	fmt.Print("x")
	_ = config.Config{}
	_ = errors.New("")
}
