//go:build !js

package main

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var (
	logGINFlag         bool
	unmarshallFromCode string
	marshallOnCommit   string
	embeddedDiagrams   bool
	port               int
)

var rootCmd = &cobra.Command{
	Use:   "aixm",
	Short: "aixm is a pseudo CLI for aixm project",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVar(&logGINFlag, "logGIN", false, "log mode for gin")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	log.SetPrefix("aixm: ")
	log.SetFlags(log.Ltime | log.Lmicroseconds | log.Lmsgprefix)
}
