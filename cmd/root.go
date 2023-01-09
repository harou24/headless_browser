package cmd

import (
	"github.com/spf13/cobra"
)

var (
	rootCmd = &cobra.Command{
		Use:   "headless_browser",
		Short: "Headless browser that makes easier to test UI components in websites.",
		Long:  "",
	}
)

func Execute() error {
	return rootCmd.Execute()
}
