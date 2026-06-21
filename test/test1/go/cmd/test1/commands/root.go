package commands

import (
	"fmt"
	"log"
	"os"

	"github.com/spf13/cobra"
)

var logFile string
var logGINFlag bool

var rootCmd = &cobra.Command{
	Use:   "test",
	Short: "test is a pseudo CLI for test project",
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVar(&logFile, "logFile", "", "path to the log file")
	rootCmd.PersistentFlags().BoolVar(&logGINFlag, "logGIN", false, "log mode for gin")

	cobra.OnInitialize(initConfig)
}

func initConfig() {
	log.SetPrefix("test: ")
	log.SetFlags(log.Lmicroseconds)

	if logFile != "" {
		file, err := os.OpenFile(logFile, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)
		if err != nil {
			log.Fatalf("Failed to open log file: %v", err)
		}
		log.SetOutput(file)
	}
}
