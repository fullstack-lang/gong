package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	test_fullstack "github.com/fullstack-lang/gong/test/go/fullstack"
	test_models "github.com/fullstack-lang/gong/test/go/models"

	test2_fullstack "github.com/fullstack-lang/gong/test2/go/fullstack"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")
	apiFlag    = flag.Bool("api", false, "it true, use api controllers instead of default controllers")

	unmarshallFromCodeStageTestA = flag.String("unmarshallFromCodeStageTestA", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommitStageTestA   = flag.String("marshallOnCommitStageTestA", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")

	unmarshallFromCodeStageTestB = flag.String("unmarshallFromCodeStageTestB", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommitStageTestB   = flag.String("marshallOnCommitStageTestB", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")
)

type BeforeCommitImplementation struct {
}

// commits on stacks on type "test"
func (impl *BeforeCommitImplementation) BeforeCommit(stage *test_models.StageStruct) {
	file, err := os.Create(fmt.Sprintf("./%s.go", *marshallOnCommitStageTestB))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	stage.Checkout()
	stage.Marshall(file, "github.com/fullstack-lang/gong/test/go/models", "main")
}

func main() {

	log.SetPrefix("test2: ")
	log.SetFlags(0)

	// parse program arguments
	flag.Parse()

	r := test2_fullstack.ServeStaticFiles(*logGINFlag)

	stageTest2 := test2_fullstack.NewStackInstance(r, "")
	_ = stageTest2

	stageTestA := test_fullstack.NewStackInstance(r, "A")
	stageTestB := test_fullstack.NewStackInstance(r, "B")

	// hook automatic marshall to go code at every commit
	if *marshallOnCommitStageTestA != "" {
		hook := new(BeforeCommitImplementation)
		stageTestA.OnInitCommitFromFrontCallback = hook
	}

	if *unmarshallFromCodeStageTestA != "" {
		stageTestA.Checkout()
		stageTestA.Reset()
		stageTestA.Commit()
		err := test_models.ParseAstFile(stageTestA, *unmarshallFromCodeStageTestA)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		stageTestA.Commit()
	}

	if *marshallOnCommitStageTestB != "" {
		hook := new(BeforeCommitImplementation)
		stageTestB.OnInitCommitFromFrontCallback = hook
	}

	if *unmarshallFromCodeStageTestB != "" {
		stageTestB.Checkout()
		stageTestB.Reset()
		stageTestB.Commit()
		err := test_models.ParseAstFile(stageTestB, *unmarshallFromCodeStageTestB)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		stageTestB.Commit()
	}

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}
