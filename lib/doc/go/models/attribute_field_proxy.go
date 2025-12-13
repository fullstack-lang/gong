package models

import (
	gong "github.com/fullstack-lang/gong/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type AttributeFieldNodeProxy struct {
	node            *tree.Node
	stager          *Stager
	classDiagram    *Classdiagram
	gongstruct      *gong.GongStruct
	gongStructShape *GongStructShape
	field           gong.FieldInterface
	attributeShape  *AttributeShape
}

func (proxy *AttributeFieldNodeProxy) OnAfterUpdate(
	stage *tree.Stage,
	staged, front *tree.Node) {

	// intercept update to the node that are when the node is checked
	if front.IsChecked && !staged.IsChecked {
		// uncheck all other diagram
		proxy.classDiagram.AddAttributeFieldShape(
			proxy.stager.stage,
			proxy.stager.gongStage,
			proxy.gongstruct,
			proxy.field,
			proxy.gongStructShape)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitFormStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}

	// the checked node is unchecked
	if !front.IsChecked && staged.IsChecked {
		proxy.classDiagram.RemoveAttributeFieldShape(proxy.stager.stage, proxy.attributeShape, proxy.gongStructShape)

		proxy.stager.UpdateAndCommitTreeStage()
		proxy.stager.UpdateAndCommitFormStage()
		proxy.stager.UpdateAndCommitSVGStage()

		proxy.stager.stage.Commit()
	}
}
