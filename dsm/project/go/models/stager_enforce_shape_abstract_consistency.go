package models

// enforceShapesAbstractConsistency iterates over all staged instances to ensure
// that visual shapes maintain valid references to their underlying abstract elements.
//
// If a shape (ConcreteType) is missing its abstract element, or if a link
// (AssociationConcreteType) is missing either its start or end abstract element,
// the shape is considered orphaned and is removed from the stage (unstaged).
// It returns true if any modifications were made, signaling that a commit is needed.
func (stager *Stager) enforceShapesAbstractConsistency() bool {
	stage := stager.stage
	var needCommit bool
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
