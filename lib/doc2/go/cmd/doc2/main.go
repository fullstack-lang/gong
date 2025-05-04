package main

import (
	"flag"
	"log"
	"strconv"

	// insertion point for models import

	"github.com/fullstack-lang/gong/lib/doc2/go/prepare"

	doc2_go "github.com/fullstack-lang/gong/lib/doc2/go"
	doc2_static "github.com/fullstack-lang/gong/lib/doc2/go/static"

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

	log.SetPrefix("doc2: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := doc2_static.ServeStaticFiles(*logGINFlag)

	splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage
	receivingAsSplitArea := &split.AsSplitArea{
		Name: "Doc2 receiving area",
	}

	split.StageBranch(splitStage, &split.View{
		Name: "Receiving doc2",
		RootAsSplitAreas: []*split.AsSplitArea{
			receivingAsSplitArea,
		},
	})

	prepare.Prepare(r, *embeddedDiagrams, "./data/zzz_diagrams.go", "doc2test", doc2_go.GoModelsDir, doc2_go.GoDiagramsDir, receivingAsSplitArea)

	log.Println("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
