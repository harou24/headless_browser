package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var port string

func init() {
	rootCmd.AddCommand(screenshot)
	screenshot.PersistentFlags().StringVarP(&port, "url", "u", "", "take a screenshot on a giving URL")
}

var screenshot = &cobra.Command{
	Use:   "screenshot",
	Short: "take a screenshot on a giving URL",
	Long:  "",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Taking a screenshot")
	},
}
