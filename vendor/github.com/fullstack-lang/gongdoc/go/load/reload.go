package load

import (
	"path/filepath"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"
	gongdoc_node2gongdoc "github.com/fullstack-lang/gongdoc/go/node2gongdoc"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

func Reload(
	gongdocStage *gongdoc_models.StageStruct,
	gongtreeStage *gongtree_models.StageStruct,
	diagramPackage *gongdoc_models.DiagramPackage,
) {

	gong_models.GetDefaultStage().Checkout()
	gong_models.GetDefaultStage().Reset()
	modelPkg, _ := gong_models.LoadSource(gong_models.GetDefaultStage(),
		filepath.Join(diagramPackage.AbsolutePathToDiagramPackage, "../models"))
	gong_models.GetDefaultStage().Commit()

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
	gongdoc_models.SetupMapDocLinkRenaming(gong_models.GetDefaultStage(), diagramPackage.Stage_)
	// end of the be removed
	gongdoc_node2gongdoc.FillUpNodeTree(gongdocStage, gongtreeStage, diagramPackage)
	diagramPackage.Stage_.Commit()
}
