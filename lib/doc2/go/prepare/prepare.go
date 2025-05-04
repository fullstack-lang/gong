package prepare

import (
	"embed"

	"github.com/gin-gonic/gin"

	doc2_models "github.com/fullstack-lang/gong/lib/doc2/go/models"
	doc2_stack "github.com/fullstack-lang/gong/lib/doc2/go/stack"

	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"

	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"

	gong "github.com/fullstack-lang/gong/go/models"
	gong_stack "github.com/fullstack-lang/gong/go/stack"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
)

func Prepare(
	r *gin.Engine,
	embeddedDiagrams bool,
	pathToDocStageFile string,
	doc2StackName string,
	goModelsDir embed.FS,
	goDiagramsDir embed.FS,
	receivingAsSplitArea *split.AsSplitArea, // split area that will receive the doc2 areas
) {

	stack := doc2_stack.NewStack(r, doc2StackName+":doc2", pathToDocStageFile, pathToDocStageFile, "", embeddedDiagrams, false)
	// stack.SetMarshallPackageName("models")
	doc2Stage := stack.Stage

	doc2Stage.Checkout()

	treeStage := tree_stack.NewStack(r, doc2StackName+":doc2-sidebar", "", "", "", false, false).Stage
	svgStage := svg_stack.NewStack(r, doc2StackName+":doc2-svg", "", "", "", false, false).Stage
	gongStage := gong_stack.NewStack(r, doc2StackName+":doc2-gong", "", "", "", false, false).Stage

	// load the code of the model of interest into the gongStage
	gong.LoadEmbedded(gongStage, goModelsDir)

	doc2_models.NewStager(
		r,
		receivingAsSplitArea,
		doc2Stage,
		treeStage,
		svgStage,
		gongStage,
		embeddedDiagrams)
}
