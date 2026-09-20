package cmd

import (
	"log"

	"github.com/fullstack-lang/gong/go/golang"
	"github.com/spf13/cobra"
)

var (
	migrateOutputPath  string
	migratePackageName string
	migrateInPlace     bool
)

var migrateCmd = &cobra.Command{
	Use:   "migrate [flags] <stage.go> [stage2.go ...]",
	Short: "Migrates single-stage data files to multi-stage StageSet format",
	Long: `Migrates one or more legacy single-stage data files (stage.go) to the new multi-stage StageSet format (stageset.go).

Examples:
  # Migrate a single stage file to stageset.go
  gong migrate data/stage.go -o data/stageset.go

  # Migrate in-place
  gong migrate data/stage.go -i

  # Migrate multiple files
  gong migrate test1/data/stage.go test2/data/stage.go
`,
	Args: cobra.MinimumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		log.SetPrefix("gong migrate: ")
		log.SetFlags(0)

		if migrateInPlace && migrateOutputPath != "" {
			log.Fatalf("cannot use both --in-place (-i) and --output (-o)")
		}

		if len(args) > 1 && migrateOutputPath != "" {
			log.Fatalf("cannot use --output (-o) when migrating multiple files")
		}

		for _, filePath := range args {
			outPath := migrateOutputPath
			resPath, err := golang.MigrateStageFile(filePath, outPath, migratePackageName, migrateInPlace)
			if err != nil {
				log.Fatalf("Error migrating %s: %v", filePath, err)
			}
			log.Printf("Successfully migrated %s -> %s", filePath, resPath)
		}
	},
}

func init() {
	migrateCmd.Flags().StringVarP(&migrateOutputPath, "output", "o", "", "output file for the migrated stage (default: stageset.go)")
	migrateCmd.Flags().BoolVarP(&migrateInPlace, "in-place", "i", false, "overwrite input file in place")
	migrateCmd.Flags().StringVarP(&migratePackageName, "package", "p", "", "package name override for the migrated stage")

	rootCmd.AddCommand(migrateCmd)
}
