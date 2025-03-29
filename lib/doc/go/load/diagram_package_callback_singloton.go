package load

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type DiagramPackageCallbacksSingloton struct {
	DiagramPackageCallback DiagramPackageCallback
	gongStage              *gong_models.Stage
	gongtreeStage          *gongtree_models.Stage
}

func (diagramPackageCallbacksSingloton *DiagramPackageCallbacksSingloton) OnAfterUpdate(
	gongdocStage *gongdoc_models.Stage,
	stagedDiagramPackage, frontDiagramPackage *gongdoc_models.DiagramPackage) {

	if stagedDiagramPackage.IsReloaded != frontDiagramPackage.IsReloaded {

		// reset the IsReloaded to false
		stagedDiagramPackage.Checkout(gongdocStage)
		stagedDiagramPackage.IsReloaded = false
		stagedDiagramPackage.Commit(gongdocStage)

		log.Println("Reload requested")
		if stagedDiagramPackage.IsEditable {
			Reload(
				diagramPackageCallbacksSingloton.gongStage,
				gongdocStage,
				diagramPackageCallbacksSingloton.gongtreeStage,
				stagedDiagramPackage)
		}
	}
}

type DiagramPackageCallback interface {
	HasSelected(gongstructName string)
}
