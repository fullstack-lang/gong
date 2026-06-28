/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

// readCmd represents the read command
var readCmd = &cobra.Command{
	Use:   "read",
	Short: "read an xsd file",
	Args:  cobra.ExactArgs(1),
	Long:  `...`,
	Run: func(cmd *cobra.Command, args []string) {

		r, stack := process(args)

		stack.Stage.Commit()
		stack.Probe.Refresh()

		// xlFilePath := filepath.Join(
		// 	filepath.Dir(*outputModelFilePath),
		// 	"schema.xlsx")

		// models.SerializeStage(stack.Stage, xlFilePath)

		log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
		err := r.Run(":" + strconv.Itoa(*port))
		if err != nil {
			log.Fatalln(err.Error())
		}
	},
}
