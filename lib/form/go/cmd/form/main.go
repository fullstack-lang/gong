package main

import (
	"flag"
	"log"
	"strconv"

	// insertion point for models import{{modelsImportDirective}}
	form_stack "github.com/fullstack-lang/gong/lib/form/go/stack"
	form_static "github.com/fullstack-lang/gong/lib/form/go/static"

	// form_models "github.com/fullstack-lang/gong/lib/form/go/models"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	port = flag.Int("port", 8080, "port server")
)

func main() {

	log.SetPrefix("form: ")
	log.SetFlags(log.Lmicroseconds)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := form_static.ServeStaticFiles(*logGINFlag)

	// setup model stack with its probe
	stack := form_stack.NewStack(r, "form", *unmarshallFromCode, *marshallOnCommit, "", *embeddedDiagrams, true)
	stack.Probe.Refresh()
	stack.Stage.Commit()

	splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage

	// form_models.NewStager(r, stack.Stage, stack.Probe)
	split.StageBranch(splitStage, &split.View{
		Name: "Split",
		RootAsSplitAreas: []*split.AsSplitArea{
			(&split.AsSplitArea{
				Size: 50,
				AsSplit: (&split.AsSplit{
					Direction: split.Horizontal,
					AsSplitAreas: []*split.AsSplitArea{
						(&split.AsSplitArea{
							Name: "Form",
							Size: 50,
							Form2: (&split.Form2{
								Name:      "Form2",
								StackName: stack.Stage.GetName(),
							}),
						}),
					},
				}),
			}),
			(&split.AsSplitArea{
				Size: 50,
				Split: (&split.Split{
					StackName: stack.Stage.GetProbeSplitStageName(),
				}),
			}),
		},
	})

	splitStage.Commit()

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
