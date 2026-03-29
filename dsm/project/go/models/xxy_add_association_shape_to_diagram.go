package models

func addAssociationShapeToDiagram[
	ATstart AbstractType,
	ATend AbstractType,
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType // the association concrete type shape
	},
	ACT_ Gongstruct](stager *Stager, start ATstart, end ATend, shapes *[]ACT) {
	stage := stager.stage
	_ = stage

	compositionShape := newConcreteAssociation(start, end, shapes)
	compositionShape.StageVoid(stager.stage)
}
