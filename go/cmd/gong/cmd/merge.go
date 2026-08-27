package cmd

import (
	"log"

	"github.com/fullstack-lang/gong/go/golang"
	"github.com/spf13/cobra"
)

var (
	mergeOutputPath  string
	mergePackageName string
)

var mergeCmd = &cobra.Command{
	Use:   "merge [flags] <stage1.go> <stage2.go> ...",
	Short: "Merges multiple stage files into a single stage file",
	Long: `Merges multiple stage files by recomputing instance IDs to avoid identifier collisions across any data model.

Example:
  gong merge file1.go file2.go -o merged.go
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.SetPrefix("gong merge: ")
		log.SetFlags(0)

		if err := golang.MergeStageFiles(args, mergeOutputPath, mergePackageName); err != nil {
			log.Fatalf("Error merging stage files: %v", err)
		}

		log.Printf("Successfully merged %d files into %s", len(args), mergeOutputPath)
	},
}

func init() {
	mergeCmd.Flags().StringVarP(&mergeOutputPath, "output", "o", "merged.go", "output file for the merged stage")
	mergeCmd.Flags().StringVarP(&mergePackageName, "package", "p", "", "package name for the merged stage (default: package of the first input file)")

	rootCmd.AddCommand(mergeCmd)
}
