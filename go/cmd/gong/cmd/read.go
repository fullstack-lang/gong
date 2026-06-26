package cmd

import (
	"log"
	"strconv"

	"github.com/spf13/cobra"

	embeddedgo "github.com/fullstack-lang/gong/go"
	gong_models "github.com/fullstack-lang/gong/go/models"
	"github.com/fullstack-lang/gong/go/probe"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
	split_static "github.com/fullstack-lang/gong/lib/split/go/static"
)

var readPort int
var readEmbeddedDiagrams bool
var logGINFlag bool

var readCmd = &cobra.Command{
	Use:   "read [path to models package]",
	Short: "read the models and run the default server with the probe",
	Args:  cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		var pkgPath string
		if len(args) > 0 {
			pkgPath = args[0]
		} else {
			pkgPath = "."
		}

		// initiate model package
		modelStage := gong_models.NewStage("gong")
		_, err := gong_models.LoadSource(modelStage, pkgPath)
		if err != nil {
			log.Fatal(err)
		}

		modelStage.Commit()

		// setup the static file server and get the controller
		r := split_static.ServeStaticFiles(logGINFlag)

		// setup model probe
		prb := probe.NewProbe(
			r,
			embeddedgo.GoModelsDir,
			embeddedgo.GoDiagramsDir,
			readEmbeddedDiagrams,
			modelStage,
		)
		modelStage.SetProbeIF(prb)
		prb.Refresh()

		// the root split name is "" by convention. Is is the same for all gong applications
		// that do not develop their specific angular component
		splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

		// set the title of the application (name of the tab)
		(&split.Title{Name: "Gong"}).Stage(splitStage)
		(&split.FavIcon{Name: "Gong",
			SVG: `<svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 24 24" fill="none">
				<circle cx="12" cy="12" r="10" fill="#ff6b6b"/>
				<path d="M8 14s1.5 2 4 2 4-2 4-2" stroke="white" stroke-width="2" stroke-linecap="round"/>
				<circle cx="9" cy="9" r="1" fill="white"/>
				<circle cx="15" cy="9" r="1" fill="white"/>
			</svg>`,
		}).Stage(splitStage)

		stager := gong_models.NewStager(r, modelStage, splitStage)

		// one for the probe of the
		split.StageBranch(splitStage, &split.View{
			Name: modelStage.GetName() + " with Probe",
			RootAsSplitAreas: []*split.AsSplitArea{
				(&split.AsSplitArea{
					Size: 50,
					AsSplit: (&split.AsSplit{
						Direction: split.Horizontal,
						AsSplitAreas: []*split.AsSplitArea{
							stager.GetAsSplitArea(),
						},
					}),
				}),
				(&split.AsSplitArea{
					Size:    50,
					AsSplit: prb.GetDataEditor(),
				}),
			},
		})

		split.StageBranch(splitStage, &split.View{
			Name: "Diagram Editor",
			RootAsSplitAreas: []*split.AsSplitArea{
				prb.GetDiagramEditor(),
			},
		})

		// commit the split stage (this will initiate the front components)
		splitStage.Commit()

		log.Println("Server ready serve on localhost:" + strconv.Itoa(readPort))
		err = r.Run(":" + strconv.Itoa(readPort))
		if err != nil {
			log.Fatalln(err.Error())
		}
	},
}

func init() {
	rootCmd.AddCommand(readCmd)
	readCmd.Flags().IntVarP(&readPort, "port", "p", 8080, "port server")
	readCmd.Flags().BoolVar(&readEmbeddedDiagrams, "embedded-diagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
	readCmd.Flags().BoolVar(&logGINFlag, "logGINFlag", false, "log gin operations")
}
