package models

import (
	"slices"
)

func onRemoveAssociationShape[
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType
	},
	ACT_ Gongstruct](stager *Stager, compositionShape ACT, shapes *[]ACT) func() {
	return func() {
		compositionShape.UnstageVoid(stager.stage)
		idx := slices.Index(*shapes, compositionShape)
		*shapes = slices.Delete(*shapes, idx, idx+1)
	}
}
