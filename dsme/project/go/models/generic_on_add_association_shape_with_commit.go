package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func onAddAssociationShapeWithCommit[
	ATstart AbstractType,
	ATend AbstractType,
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType
	},
	ACT_ Gongstruct,
](
	stager *Stager, start ATstart, end ATend, shapes *[]ACT) func(
	stage *tree.Stage, updatedButton *tree.Button) {
	return func(_ *tree.Stage, _ *tree.Button) {
		addAssociationShapeToDiagram(stager, start, end, shapes)
		stager.stage.Commit()
	}
}
