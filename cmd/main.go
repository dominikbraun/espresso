package main

import (
	"github.com/dominikbraun/espresso/core"
	"github.com/dominikbraun/espresso/settings"
	"github.com/spf13/cobra"
	"log"
)

const (
	settingsFile string = "site"
)

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
			var s settings.Site

			if err := settings.FromFile(buildPath, settingsFile, &s); err != nil {
				return err
			}

			return core.RunBuild(buildPath, &s)
		},
	}

	espressoCmd.AddCommand(buildCmd)

	if err := espressoCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
