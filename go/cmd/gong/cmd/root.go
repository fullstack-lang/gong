package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "gong",
	Short: "gong is a tool to generate code from a models package.",
	Long: `gong is a CLI tool to perform code generation for the gong framework.
It can generate backend code, frontend code, and documentation.`,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
