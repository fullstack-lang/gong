package prepare

import (
	"embed"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/fullstack-lang/gong/lib/doc2/go/fullstack"
	"github.com/fullstack-lang/gong/lib/doc2/go/models"

	tree_fullstack "github.com/fullstack-lang/gong/lib/tree/go/fullstack"

	svg_fullstack "github.com/fullstack-lang/gong/lib/svg/go/fullstack"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong "github.com/fullstack-lang/gong/go/models"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
)

// hook marhalling to stage
type beforeCommitImplementation struct {
	marshallOnCommit string

	packageName string
}

func (beforeCommitImplementation *beforeCommitImplementation) BeforeCommit(stage *models.Stage) {

	// the ".go" is not provided
	filename := beforeCommitImplementation.marshallOnCommit
	if !strings.HasSuffix(filename, ".go") {
		filename = filename + ".go"
	}

	file, err := os.Create(fmt.Sprintf("./%s", filename))
	if err != nil {
		log.Fatal(err.Error())
	}
	defer file.Close()

	packageName := beforeCommitImplementation.packageName
	if packageName == "" {
		packageName = "main"
	}

	stage.Checkout()
	stage.Marshall(file, "github.com/fullstack-lang/gong/lib/doc2/go/models", packageName)
}

func Prepare(
	r *gin.Engine,
	embeddedDiagrams bool,
	doc2StackName string,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
	receivingAsSplitArea *split.AsSplitArea, // split area that will receive the doc2 areas
) {
	var stage *models.Stage

	stage, _ = fullstack.NewStackInstance(r, doc2StackName)

	stage.Checkout()
	stage.Reset()
	stage.Commit()

	if !embeddedDiagrams {
		err := models.ParseAstFile(stage, "../../diagrams/diagrams.go")

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		BeforeCommitImplementation := &beforeCommitImplementation{
			marshallOnCommit: "../../diagrams/diagrams.go",
			packageName:      "diagrams", // necessity because the diagram file is in a diagrams package
		}
		stage.OnInitCommitCallback = BeforeCommitImplementation
	} else {
		err := models.ParseAstEmbeddedFile(stage, goDiagramsDir, "diagrams/diagrams.go")

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}
	}

	stage.Commit()

	stage.Checkout()

	treeStage, _ := tree_fullstack.NewStackInstance(r, doc2StackName+":doc2-sidebar", "", "")
	svgStage, _ := svg_fullstack.NewStackInstance(r, doc2StackName+":doc2-svg", "", "", "")
	gongStage, _ := gong_fullstack.NewStackInstance(r, doc2StackName+":doc2-gong", "", "")

	// load the code of the model of interest into the gongStage
	gong.LoadEmbedded(gongStage, goModelsDir)

	models.NewStager(
		r,
		receivingAsSplitArea,
		stage,
		treeStage,
		svgStage,
		gongStage,
		embeddedDiagrams)
}
