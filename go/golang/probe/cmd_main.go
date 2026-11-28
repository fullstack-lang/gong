package probe

import "fmt"

func GetProbeCmdMainTemplate() string {
	splitImport := `split "github.com/fullstack-lang/gong/lib/split/go/models"`
	splitStackImport := `split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"`
	splitStaticImport := `split_static "github.com/fullstack-lang/gong/lib/split/go/static"`
	splitPkg := "split"

	return fmt.Sprintf(`//go:build !js

// generated code - do not edit
package main

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"{{PkgPathRoot}}/level1stack"
	"{{PkgPathRoot}}/models"

	%s
	%s
	%s
)

var (
	embeddedDiagrams   bool
	port               int
	unmarshallFromCode string
	marshallOnCommit   string
)

var (
	editOut       string
	editDir       bool
	editSkipGoMod bool
)

var editCmd = &cobra.Command{
	Use:     "edit [data/stage.go]",
	Aliases: []string{"edit-stageset", "stageset", "edit-multistage"},
	Short:   "Edit a stage file or directory",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
			marshallOnCommit = args[0]
		}
		if editOut != "" {
			marshallOnCommit = editOut
		}
		if editDir && marshallOnCommit != "" {
			_ = os.MkdirAll(marshallOnCommit, 0755)
		}
		executeServer()
	},
}

func executeServer() {
	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1Stack("{{PkgName}}", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams, editSkipGoMod)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// Create root split stage for the probe
	rootSplitStage := %s_stack.NewStack(stack.R, "", "", "", "", false, false).Stage

	rootSplitStage.StageBranch(&%s.View{
		Name: "Data Probe & Data Model",
		RootAsSplitAreas: []*%s.AsSplitArea{
			{
				Split: &%s.Split{
					StackName: stack.Stage.GetProbeSplitStageName(),
				},
			},
		},
	})
	rootSplitStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := %s_static.RunServer(stack.R, ":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

var (
	migrateOut       string
	migrateDir       bool
	migrateSkipGoMod bool
)

var migrateCmd = &cobra.Command{
	Use:   "migrate [data/stage.go]",
	Short: "Migrate a single-stage data file to multi-stage format or directory module",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := "data/stage.go"
		if len(args) > 0 {
			inputFile = args[0]
		}
		outputFile := migrateOut
		if outputFile == "" {
			if migrateDir {
				outputFile = "data_dir"
			} else {
				outputFile = "data/stageset.go"
			}
		}

		stage := models.NewStage("")
		if err := stage.ParseAstFile(inputFile, true); err != nil {
			log.Fatalf("failed to parse input stage file: %%v", err)
		}
		if migrateDir {
			if err := os.MkdirAll(outputFile, 0755); err != nil {
				log.Fatalf("failed to create output directory: %%v", err)
			}
			if err := stage.MarshallDirectory(outputFile, migrateSkipGoMod); err != nil {
				log.Fatalf("failed to marshall stage directory: %%v", err)
			}
		} else {
			stage.MarshallFile(outputFile, "main")
		}
		log.Printf("Successfully migrated %%s to %%s", inputFile, outputFile)
	},
}

var rootCmd = &cobra.Command{
	Use:   "{{ProbeCmdName}}",
	Short: "probe CLI for {{PkgName}}",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
			marshallOnCommit = args[0]
		}
		executeServer()
	},
}

func main() {
	editCmd.Flags().StringVarP(&editOut, "out", "o", "", "specify a different file or directory to save commits to")
	editCmd.Flags().BoolVarP(&editDir, "dir", "d", false, "persist stage data as a Go module directory")
	editCmd.Flags().BoolVar(&editSkipGoMod, "skip-go-mod", false, "do not generate go.mod in directory persistence mode")
	rootCmd.AddCommand(editCmd)
	migrateCmd.Flags().StringVarP(&migrateOut, "out", "o", "", "output file or directory path (default: data/stageset.go or data_dir)")
	migrateCmd.Flags().BoolVarP(&migrateDir, "dir", "d", false, "migrate into a Go module directory")
	migrateCmd.Flags().BoolVar(&migrateSkipGoMod, "skip-go-mod", false, "do not generate go.mod in directory persistence mode")
	rootCmd.AddCommand(migrateCmd)
	rootCmd.PersistentFlags().BoolVar(&embeddedDiagrams, "embedded-diagrams", true, "parse/analysis go/models and go/embeddedDiagrams")
	rootCmd.PersistentFlags().IntVar(&port, "port", 8080, "port server")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
`,
		splitImport,
		splitStackImport,
		splitStaticImport,
		splitPkg,
		splitPkg,
		splitPkg,
		splitPkg,
		splitPkg,
	)
}

var ProbeCmdMainTemplate = GetProbeCmdMainTemplate()

const ProbeCmdMainTemplateFullStack = `//go:build !js

// generated code - do not edit
package main

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"{{PkgPathRoot}}/stack"
	"{{PkgPathRoot}}/static"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

var (
	embeddedDiagrams   bool
	port               int
	unmarshallFromCode string
	marshallOnCommit   string
	editOut            string
)

var editCmd = &cobra.Command{
	Use:     "edit [data/stage.go]",
	Aliases: []string{"edit-stageset", "stageset", "edit-multistage"},
	Short:   "Edit a stage file",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
			marshallOnCommit = args[0]
		}
		if editOut != "" {
			marshallOnCommit = editOut
		}
		executeServer()
	},
}

func executeServer() {
	// setup the static file server and get the controller
	r := static.ServeStaticFiles(false)

	// setup model stack with its probe
	stack := stack.NewStack(r, "{{PkgName}}", unmarshallFromCode, marshallOnCommit, "", embeddedDiagrams, true)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// Create root split stage for the probe
	rootSplitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

	rootSplitStage.StageBranch(&split.View{
		Name: "Data Probe & Data Model",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stack.Stage.GetProbeSplitStageName(),
				},
			},
		},
	})
	rootSplitStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(port))
	err := static.RunServer(r, ":" + strconv.Itoa(port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}

var rootCmd = &cobra.Command{
	Use:   "{{ProbeCmdName}}",
	Short: "probe CLI for {{PkgName}}",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
			marshallOnCommit = args[0]
		}
		executeServer()
	},
}

func main() {
	editCmd.Flags().StringVar(&editOut, "out", "", "specify a different file to save commits to")
	rootCmd.AddCommand(editCmd)
	rootCmd.PersistentFlags().BoolVar(&embeddedDiagrams, "embedded-diagrams", true, "parse/analysis go/models and go/embeddedDiagrams")
	rootCmd.PersistentFlags().IntVar(&port, "port", 8080, "port server")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
`
