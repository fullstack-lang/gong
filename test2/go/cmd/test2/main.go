package main

import (
	"embed"
	"flag"
	"fmt"
	"io/fs"
	"log"
	"net/http"
	"os"

	"github.com/gin-contrib/static"

	"github.com/fullstack-lang/gong/test/go/fullstack"
	test_fullstack "github.com/fullstack-lang/gong/test/go/fullstack"
	"github.com/fullstack-lang/gong/test/go/models"
	test_models "github.com/fullstack-lang/gong/test/go/models"

	test2_fullstack "github.com/fullstack-lang/gong/test2/go/fullstack"
)

var (
	logDBFlag  = flag.Bool("logDB", false, "log mode for db")
	logGINFlag = flag.Bool("logGIN", false, "log mode for gin")
	apiFlag    = flag.Bool("api", false, "it true, use api controllers instead of default controllers")

	unmarshallFromCodeStageTestA = flag.String("unmarshallFromCodeStageTestA", "", "unmarshall data from go file and '.go' (must be lowercased without spaces), If unmarshallFromCode arg is '', no unmarshalling")
	marshallOnCommitStageTestA   = flag.String("marshallOnCommitStageTestA", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")
	marshallOnCommitStageTestB   = flag.String("marshallOnCommitStageTestB", "", "on all commits, marshall staged data to a go file with the marshall name and '.go' (must be lowercased without spaces). If marshall arg is '', no marshalling")
)

type BeforeCommitImplementation struct {
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *models.StageStruct) {
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

	r := fullstack.ServeStaticFiles(*logGINFlag)

	test2_fullstack.NewStackInstance(r, "")
	stageTestA := test_fullstack.NewStackInstance(r, "A")
	stageTestB := test_fullstack.NewStackInstance(r, "B")

	// hook automatic marshall to go code at every commit
	if *marshallOnCommitStageTestA != "" {
		hook := new(BeforeCommitImplementation)
		stageTestA.OnInitCommitFromFrontCallback = hook
	}

	if *marshallOnCommitStageTestB != "" {
		hook := new(BeforeCommitImplementation)
		stageTestB.OnInitCommitFromFrontCallback = hook
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

	log.Printf("Server ready serve on localhost:8080")
	r.Run()
}

type embedFileSystem struct {
	http.FileSystem
}

func (e embedFileSystem) Exists(prefix string, path string) bool {
	_, err := e.Open(path)
	return err == nil
}

func EmbedFolder(fsEmbed embed.FS, targetPath string) static.ServeFileSystem {
	fsys, err := fs.Sub(fsEmbed, targetPath)
	if err != nil {
		panic(err)
	}
	return embedFileSystem{
		FileSystem: http.FS(fsys),
	}
}
