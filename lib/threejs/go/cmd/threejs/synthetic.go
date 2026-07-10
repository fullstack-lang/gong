//go:build !js

package main

import (
	"github.com/spf13/cobra"
)

var syntheticCmd = &cobra.Command{
	Use:   "synthetic",
	Short: "Generate synthetic strange forms for threejs",
	Run: func(cmd *cobra.Command, args []string) {
		generateStrangeFormsFlag = true
		executeServer()
	},
}

func init() {
	rootCmd.AddCommand(syntheticCmd)
	syntheticCmd.Flags().IntVar(&port, "port", 8080, "port server")
}
