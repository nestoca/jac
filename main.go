//go:generate antlr -Dlanguage=Go -o pkg/filtering/parser Pattern.g4
package main

import (
	"fmt"
	"github.com/nestoca/jac/cmd"
	"os"
)

var version = "v0.0.0"

func main() {
	if err := cmd.NewRootCmd(version).Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
