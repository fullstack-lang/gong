package models

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
	stager *Stager, start ATstart, end ATend, shapes *[]ACT) func() {
	return func() {
		addAssociationShapeToDiagram(stager, start, end, shapes)
	}
}
