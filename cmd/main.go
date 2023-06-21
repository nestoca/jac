//go:generate antlr -Dlanguage=Go -o parser -package parser Pattern.g4
package main

import (
	"fmt"
	"os"
)

var version = "v0.0.0"

func main() {
	rootCmd := createRootCmd()

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
