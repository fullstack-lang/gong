package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func onAddAssociationShape[
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
	return func(stage *tree.Stage, updatedButton *tree.Button) {
		addAssociationShapeToDiagram(stager, start, end, shapes)
	}
}
