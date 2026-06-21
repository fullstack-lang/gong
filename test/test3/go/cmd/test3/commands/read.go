package commands

import (
	"log"
	"strconv"

	"github.com/fullstack-lang/gong/test/test3/go/level1stack"
	"github.com/fullstack-lang/gong/test/test3/go/models"
	"github.com/spf13/cobra"
)

var readPort int
var readEmbeddedDiagrams bool

var readCmd = &cobra.Command{
	Use:   "read [file]",
	Short: "Read a stage file without marshalling on commit",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		file := args[0]

		// Empty string for marshallOnCommit prevents writing
		stack := level1stack.NewLevel1StackDelta("test3", file, "", true, readEmbeddedDiagrams, false)

		stack.Probe.Refresh()

		models.NewStager(
			stack.R,
			stack.Stage,
			stack.Probe,
		)

		log.Println("Server ready serve on localhost:" + strconv.Itoa(readPort))
		err := stack.R.Run(":" + strconv.Itoa(readPort))
		if err != nil {
			log.Fatalln(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().IntVarP(&readPort, "port", "p", 8080, "port server")
	readCmd.Flags().BoolVar(&readEmbeddedDiagrams, "embedded-diagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
}
