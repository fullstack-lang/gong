package models

func (*Stager) enforceShapesAbstractConsistency(stage *Stage, needCommit bool) bool {
	for _, instance := range stage.GetInstances() {
		if shape, ok := instance.(ConcreteType); ok {
			if shape.GetAbstractElement() == nil {
				shape.UnstageVoid(stage)
				needCommit = true
			}
			continue
		}
		if associationShape, ok := instance.(AssociationConcreteType); ok {
			if associationShape.GetAbstractStartElement() == nil {
				associationShape.UnstageVoid(stage)
				needCommit = true
				continue
			}
			if associationShape.GetAbstractEndElement() == nil {
				associationShape.UnstageVoid(stage)
				needCommit = true
				continue
			}
		}

	}
	return needCommit
}
