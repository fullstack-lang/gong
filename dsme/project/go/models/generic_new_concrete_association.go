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

	if taskInputShape, ok := any(compositionShape).(*TaskInputShape); ok {
		taskInputShape.SetStartOrientation(ORIENTATION_HORIZONTAL)
		taskInputShape.SetEndOrientation(ORIENTATION_HORIZONTAL)
		taskInputShape.CornerOffsetRatio = 0.2
	}
	if taskOutputShape, ok := any(compositionShape).(*TaskOutputShape); ok {
		taskOutputShape.SetStartOrientation(ORIENTATION_HORIZONTAL)
		taskOutputShape.SetEndOrientation(ORIENTATION_HORIZONTAL)
		taskOutputShape.CornerOffsetRatio = 0.2
	}

	compositionShape.SetCornerOffsetRatio(1.68)
	compositionShape.SetStartRatio(0.5)
	compositionShape.SetEndRatio(0.5)
	*shapes = append(*shapes, compositionShape)

	return compositionShape
}
