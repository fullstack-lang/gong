package main

import (
	"flag"
	"fmt"
	"log"
	"os"
	"strings"

	gong_go "github.com/fullstack-lang/gong/go"
	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong_models "github.com/fullstack-lang/gong/go/models"
	gong_static "github.com/fullstack-lang/gong/go/static"

	gongdoc_load "github.com/fullstack-lang/gongdoc/go/load"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")

	marshallOnStartup  = flag.String("marshallOnStartup", "", "at startup, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")
	unmarshallFromCode = flag.String("unmarshallFromCode", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	unmarshall         = flag.String("unmarshall", "", "unmarshall data from marshall name and '.go' (must be lowercased without spaces), If unmarshall arg is '', no unmarshalling")
	marshallOnCommit   = flag.String("marshallOnCommit", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	diagrams         = flag.Bool("diagrams", true, "parse/analysis go/models and go/diagrams")
	embeddedDiagrams = flag.Bool("embeddedDiagrams", false, "parse/analysis go/models and go/embeddedDiagrams")
)

// InjectionGateway is the singloton that stores all functions
// that can set the objects the stage
// InjectionGateway stores function as a map of names
var InjectionGateway = make(map[string](func()))

// hook marhalling to stage
type BeforeCommitImplementation struct {
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *gong_models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommit))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stage.Checkout()
	stage.Marshall(file, "github.com/fullstack-lang/gong/go/models", "main")
}

func main() {

	log.SetPrefix("gong: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	// setup the static file server and get the controller
	r := gong_static.ServeStaticFiles(*logGINFlag)

	// setup stack
	var stage *gong_models.StageStruct
	if *marshallOnCommit != "" {
		// persistence in a SQLite file on disk in memory
		stage = gong_fullstack.NewStackInstance(r, "github.com/fullstack-lang/gong/go/models")
	} else {
		// persistence in a SQLite file on disk
		stage = gong_fullstack.NewStackInstance(r, "github.com/fullstack-lang/gong/go/models", "./test.db")
	}

	// generate injection code from the stage
	if *marshallOnStartup != "" {

		if strings.Contains(*marshallOnStartup, " ") {
			log.Fatalln(*marshallOnStartup + " must not contains blank spaces")
		}
		if strings.ToLower(*marshallOnStartup) != *marshallOnStartup {
			log.Fatalln(*marshallOnStartup + " must be lowercases")
		}

		file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnStartup))
		if err != nil {
			log.Fatal(err.Error())
		}
		defer file.Close()

		stage.Checkout()
		stage.Marshall(file, "github.com/fullstack-lang/gong/go/models", "main")
		os.Exit(0)
	}

	// setup the stage by injecting the code from code database
	if *unmarshall != "" {
		stage.Checkout()
		stage.Reset()
		stage.Commit()
		if InjectionGateway[*unmarshall] != nil {
			InjectionGateway[*unmarshall]()
		}
		stage.Commit()
	} else {
		// in case the database is used, checkout the content to the stage
		stage.Checkout()
	}

	if *unmarshallFromCode != "" {
		stage.Checkout()
		stage.Reset()
		stage.Commit()
		err := gong_models.ParseAstFile(stage, *unmarshallFromCode)

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
		stage.OnInitCommitFromFrontCallback = hook
	}

	gongdoc_load.Load(
		"gong",
		"github.com/fullstack-lang/gong/go/models",
		gong_go.GoModelsDir,
		gong_go.GoDiagramsDir,
		r,
		*embeddedDiagrams,
		&stage.Map_GongStructName_InstancesNb)

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}
