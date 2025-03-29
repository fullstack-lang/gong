package load

import (
	"path/filepath"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gong/lib/doc/go/models"
	gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func Reload(
	gongStage *gong_models.Stage,
	gongdocStage *gongdoc_models.Stage,
	gongtreeStage *gongtree_models.Stage,
	diagramPackage *gongdoc_models.DiagramPackage,
) {

	gongStage.Checkout()
	gongStage.Reset()
	modelPkg, _ := gong_models.LoadSource(gongStage,
		filepath.Join(diagramPackage.AbsolutePathToDiagramPackage, "../models"))
	gongStage.Commit()

	diagramPackage.Stage_.Checkout()
	diagramPackage.Stage_.Reset()
	diagramPackage.SelectedClassdiagram = nil
	diagramPackage.Stage_.Commit()

	diagramPackage.Classdiagrams = nil
	diagramPackage.Umlscs = nil
	diagramPackage.ModelPkg = modelPkg

	diagramPackage, _ = LoadDiagramPackage(
		gongdocStage,
		gongtreeStage,
		filepath.Join(diagramPackage.AbsolutePathToDiagramPackage, "../models"),
		modelPkg, true)

	// to be removed after fix of [issue](https://github.com/golang/go/issues/57559)
	gongdoc_models.SetupMapDocLinkRenaming(gongStage, diagramPackage.Stage_)
	// end of the be removed
	diagramPackage.Stage_.Commit()
}
