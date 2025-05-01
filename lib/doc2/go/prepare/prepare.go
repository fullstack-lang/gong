package prepare

import (
	"github.com/gin-gonic/gin"

	doc2_models "github.com/fullstack-lang/gong/lib/doc2/go/models"
	doc2_stack "github.com/fullstack-lang/gong/lib/doc2/go/stack"

	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"

	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"

	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
)

func Prepare(
	r *gin.Engine,
	embeddedDiagrams bool,
	pathToDocStageFile string,
	doc2StackName string,
) {

	stack := doc2_stack.NewStack(r, "", pathToDocStageFile+".go", pathToDocStageFile, "", embeddedDiagrams, true)
	// stack.SetMarshallPackageName("models")
	doc2Stage := stack.Stage

	splitStage := split_stack.NewStack(r, "", "", "", "", false, false).Stage
	treeStage := tree_stack.NewStack(r, doc2StackName+":doc2-sidebar", "", "", "", false, false).Stage
	svgStage := svg_stack.NewStack(r, "", "", "", "", false, false).Stage

	doc2_models.NewStager(
		r,
		doc2Stage,
		splitStage,
		treeStage,
		svgStage)
}
