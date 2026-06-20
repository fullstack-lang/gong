// generated code (do not edit)
package models

func newConcreteAssociation[
	ATstart AbstractType,
	ATend AbstractType,
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType
	},
	ACT_ Gongstruct](start ATstart, end ATend, shapes *[]ACT) ACT {
	compositionShape := ACT(new(ACT_))
	compositionShape.SetName(start.GetName() + " to " + end.GetName())
	compositionShape.SetAbstractStartElement(start)
	compositionShape.SetAbstractEndElement(end)
	compositionShape.SetStartOrientation(ORIENTATION_VERTICAL)
	compositionShape.SetEndOrientation(ORIENTATION_VERTICAL)

	compositionShape.SetCornerOffsetRatio(1.68)
	compositionShape.SetStartRatio(0.5)
	compositionShape.SetEndRatio(0.5)
	*shapes = append(*shapes, compositionShape)

	return compositionShape
}
