// Package main provides the Espresso commands, parses the CLI input and
// calls the appropriate Espresso core functions for further processing.
package main

import (
	"fmt"
	"github.com/dominikbraun/espresso/config"
	"github.com/dominikbraun/espresso/core"
	"github.com/spf13/cobra"
	"log"
)

const (
	settingsFile string = "site"
)

var version = "UNSPECIFIED"

// func main builds all CLI commands and processes the CLI input.
func main() {
	espressoCmd := &cobra.Command{
		Use: "espresso",
		RunE: func(cmd *cobra.Command, args []string) error {
			return cmd.Help()
		},
	}

	buildCmd := &cobra.Command{
		Use:  "build <PATH>",
		Args: cobra.ExactArgs(1),
		RunE: func(cmd *cobra.Command, args []string) error {
			buildPath := args[0]
			var s config.Site

			if err := config.FromFile(buildPath, settingsFile, &s); err != nil {
				return err
			}

			return core.RunBuild(buildPath, &s)
		},
	}

	versionCmd := &cobra.Command{
		Use: "version",
		RunE: func(cmd *cobra.Command, args []string) error {
			fmt.Printf("Espresso %s\n", version)
			return nil
		},
	}

	espressoCmd.AddCommand(buildCmd)
	espressoCmd.AddCommand(versionCmd)

	if err := espressoCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
