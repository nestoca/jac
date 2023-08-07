package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

func main() {
	if err := NewRootCmd().Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func NewRootCmd() *cobra.Command {
	cmd := &cobra.Command{
		Use:   "render",
		Short: "Render teams page",
		Args:  cobra.ExactArgs(3),
		RunE: func(cmd *cobra.Command, args []string) error {
			catalogDir := args[0]
			templateFile := args[1]
			outputFile := args[2]
			return Render(catalogDir, templateFile, outputFile)
		},
	}

	return cmd
}
