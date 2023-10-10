package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strconv"

	test2_go "github.com/fullstack-lang/gong/test2/go"
	test2_fullstack "github.com/fullstack-lang/gong/test2/go/fullstack"
	test2_models "github.com/fullstack-lang/gong/test2/go/models"
	test2_orm "github.com/fullstack-lang/gong/test2/go/orm"
	test2_probe "github.com/fullstack-lang/gong/test2/go/probe"

	test2_static "github.com/fullstack-lang/gong/test2/go/static"
)

var (
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")

	testPackageHierarchy = flag.Bool("testPackageHierarchy", false, "performs some tests related to package hierarchy")

	port = flag.Int("port", 8080, "port server")
)

// InjectionGateway is the singloton that stores all functions
// that can set the objects the stage
// InjectionGateway stores function as a map of names
var InjectionGateway = make(map[string](func()))

// hook marhalling to stage
type BeforeCommitImplementation struct {
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *test2_models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stage.Checkout()
	stage.Marshall(file, "github.com/fullstack-lang/gong/test2/go/models", "main")
}

func main() {

	log.SetPrefix("test2: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := test2_static.ServeStaticFiles(*logGINFlag)

	if *testPackageHierarchy {
		stage, backRepo := test2_fullstack.NewStackInstance(r, "test2")
		_ = stage
		_ = backRepo
		// stage_x, backRepo_x := test2_fullstack_x.NewStackInstance(r, "test2")
		// _ = stage_x
		// _ = backRepo_x
	}

	if !*testPackageHierarchy {
		// setup stack
		var stage *test2_models.StageStruct
		var backRepo *test2_orm.BackRepoStruct

		if *marshallOnCommit != "" {
			// persistence in a SQLite file on disk in memory
			stage, backRepo = test2_fullstack.NewStackInstance(r, "test2")
		} else {
			// persistence in a SQLite file on disk
			stage, backRepo = test2_fullstack.NewStackInstance(r, "test2", "./test2.db")
		}

		if *unmarshallFromCode != "" {
			stage.Checkout()
			stage.Reset()
			stage.Commit()
			err := test2_models.ParseAstFile(stage, *unmarshallFromCode)

			// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
			// xxx.go might be absent the first time. However, this shall not be a show stopper.
			if err != nil {
				log.Println("no file to read " + err.Error())
			}

			stage.Commit()
		} else {
			// in case the database is used, checkout the content to the stage
			stage.Checkout()
		}

		// hook automatic marshall to go code at every commit
		if *marshallOnCommit != "" {
			hook := new(BeforeCommitImplementation)
			stage.OnInitCommitCallback = hook
		}
		test2_probe.NewProbe(r, test2_go.GoModelsDir, test2_go.GoDiagramsDir,
			*embeddedDiagrams, "test2", stage, backRepo)
	}

	log.Printf("Server ready serve on localhost:" + strconv.Itoa(*port))
	err := r.Run(":" + strconv.Itoa(*port))
	if err != nil {
		log.Fatalln(err.Error())
	}
}
