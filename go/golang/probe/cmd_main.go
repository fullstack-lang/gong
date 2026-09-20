package probe

import "fmt"

func GetProbeCmdMainTemplate(useSplitlite bool, hasStageSet bool) string {
	splitImport := `split "github.com/fullstack-lang/gong/lib/split/go/models"`
	splitStackImport := `split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"`
	splitStaticImport := `split_static "github.com/fullstack-lang/gong/lib/split/go/static"`
	splitPkg := "split"
	if useSplitlite {
		splitImport = `splitlite "github.com/fullstack-lang/gong/lib/splitlite/go/models"`
		splitStackImport = `splitlite_stack "github.com/fullstack-lang/gong/lib/splitlite/go/stack"`
		splitStaticImport = `splitlite_static "github.com/fullstack-lang/gong/lib/splitlite/go/static"`
		splitPkg = "splitlite"
	}

	stageSetCmd := ""
	addStageSetCmd := ""
	modelsImport := ""
	if hasStageSet {
		modelsImport = "\n\t\"{{PkgPathRoot}}/models\""
		stageSetCmd = fmt.Sprintf(`
var editStageSetCmd = &cobra.Command{
	Use:     "edit-stageset [data/stage.go]",
	Aliases: []string{"stageset", "edit-multistage"},
	Short:   "Edit a multi-stage StageSet file (temporary command)",
	Args:    cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
			marshallOnCommit = args[0]
		}
		executeServerStageSet()
	},
}

func executeServerStageSet() {
	stack := level1stack.NewLevel1StackStageSet("{{PkgName}}", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)

	stack.Probe.Refresh()
	if stack.StageSetProbe != nil {
		stack.StageSetProbe.Refresh()
	}

	rootSplitStage := %s_stack.NewStack(stack.R, "", "", "", "", false, false).Stage

	if stack.StageSet != nil {
		%s.StageBranch(rootSplitStage, &%s.View{
			Name: "StageSet Probe",
			RootAsSplitAreas: []*%s.AsSplitArea{
				{
					Split: &%s.Split{
						StackName: stack.StageSet.GetProbeSplitStageName(),
					},
				},
			},
		})
	}

	%s.StageBranch(rootSplitStage, &%s.View{
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

var migrateOut string

var migrateCmd = &cobra.Command{
	Use:   "migrate [data/stage.go]",
	Short: "Migrate a single-stage data file to multiple-stage StageSet format",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		inputFile := "data/stage.go"
		if len(args) > 0 {
			inputFile = args[0]
		}
		outputFile := migrateOut
		if outputFile == "" {
			outputFile = "data/stageset.go"
		}

		stage := models.NewStage("")
		if err := stage.ParseAstFile(inputFile, true); err != nil {
			log.Fatalf("failed to parse input stage file: %%v", err)
		}
		stageSet := models.NewStageSetFromStage(stage)
		stageSet.MarshallFile(outputFile, "main")
		log.Printf("Successfully migrated %%s to %%s", inputFile, outputFile)
	},
}
`, splitPkg, splitPkg, splitPkg, splitPkg, splitPkg, splitPkg, splitPkg, splitPkg, splitPkg, splitPkg)
		addStageSetCmd = "\n\trootCmd.AddCommand(editStageSetCmd)\n\tmigrateCmd.Flags().StringVar(&migrateOut, \"out\", \"\", \"output file path (default: data/stageset.go)\")\n\trootCmd.AddCommand(migrateCmd)"
	}

	return fmt.Sprintf(`//go:build !js

// generated code - do not edit
package main

import (
	"log"
	"os"
	"strconv"

	"github.com/spf13/cobra"

	"{{PkgPathRoot}}/level1stack"%s

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

var editCmd = &cobra.Command{
	Use:   "edit [data/stage.go]",
	Short: "Edit a stage file",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
			marshallOnCommit = args[0]
		}
		executeServer()
	},
}

func executeServer() {
	// setup
	// - model level1 stack with its probe
	// - unmarshall/marshall go file with stage data
	stack := level1stack.NewLevel1Stack("{{PkgName}}", unmarshallFromCode, marshallOnCommit, true, embeddedDiagrams)

	// refresh the probe, therefore we can see what has been unmarshalled
	stack.Probe.Refresh()

	// Create root split stage for the probe
	rootSplitStage := %s_stack.NewStack(stack.R, "", "", "", "", false, false).Stage

	%s.StageBranch(rootSplitStage, &%s.View{
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
%s
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
	rootCmd.AddCommand(editCmd)%s
	rootCmd.PersistentFlags().BoolVar(&embeddedDiagrams, "embedded-diagrams", true, "parse/analysis go/models and go/embeddedDiagrams")
	rootCmd.PersistentFlags().IntVar(&port, "port", 8080, "port server")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
`,
		modelsImport,
		splitImport,
		splitStackImport,
		splitStaticImport,
		splitPkg,
		splitPkg,
		splitPkg,
		splitPkg,
		splitPkg,
		splitPkg,
		stageSetCmd,
		addStageSetCmd,
	)
}

var ProbeCmdMainTemplate = GetProbeCmdMainTemplate(false, false)
var ProbeCmdMainSplitliteTemplate = GetProbeCmdMainTemplate(true, false)

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
)

var editCmd = &cobra.Command{
	Use:   "edit [data/stage.go]",
	Short: "Edit a stage file",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) > 0 {
			unmarshallFromCode = args[0]
			marshallOnCommit = args[0]
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

	split.StageBranch(rootSplitStage, &split.View{
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
	rootCmd.AddCommand(editCmd)
	rootCmd.PersistentFlags().BoolVar(&embeddedDiagrams, "embedded-diagrams", true, "parse/analysis go/models and go/embeddedDiagrams")
	rootCmd.PersistentFlags().IntVar(&port, "port", 8080, "port server")

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
		os.Exit(1)
	}
}
`
