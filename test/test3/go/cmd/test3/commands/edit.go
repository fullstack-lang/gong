package commands

import (
	"log"
	"strconv"

	"github.com/fullstack-lang/gong/test/test3/go/level1stack"
	"github.com/fullstack-lang/gong/test/test3/go/models"
	"github.com/spf13/cobra"
)

var editPort int
var editEmbeddedDiagrams bool
var editOut string

var editCmd = &cobra.Command{
	Use:   "edit [file]",
	Short: "Edit a stage file",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		marshallOnCommit := file
		if editOut != "" {
			marshallOnCommit = editOut
		}

		stack := level1stack.NewLevel1StackDelta("test3", file, marshallOnCommit, true, editEmbeddedDiagrams, false)

		stack.Probe.Refresh()

		models.NewStager(
			stack.R,
			stack.Stage,
			stack.Probe,
		)

		log.Println("Server ready serve on localhost:" + strconv.Itoa(editPort))
		err := stack.R.Run(":" + strconv.Itoa(editPort))
		if err != nil {
			log.Fatalln(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(editCmd)
	editCmd.Flags().IntVarP(&editPort, "port", "p", 8080, "port server")
	editCmd.Flags().BoolVar(&editEmbeddedDiagrams, "embedded-diagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
	editCmd.Flags().StringVarP(&editOut, "out", "o", "", "specify a different file to save commits to")
}
