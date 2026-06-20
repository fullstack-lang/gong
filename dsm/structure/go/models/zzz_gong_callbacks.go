// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *DiagramStructure:
		if stage.OnAfterDiagramStructureCreateCallback != nil {
			stage.OnAfterDiagramStructureCreateCallback.OnAfterCreate(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Link:
		if stage.OnAfterLinkCreateCallback != nil {
			stage.OnAfterLinkCreateCallback.OnAfterCreate(stage, target)
		}
	case *LinkAssociationShape:
		if stage.OnAfterLinkAssociationShapeCreateCallback != nil {
			stage.OnAfterLinkAssociationShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Part:
		if stage.OnAfterPartCreateCallback != nil {
			stage.OnAfterPartCreateCallback.OnAfterCreate(stage, target)
		}
	case *PartShape:
		if stage.OnAfterPartShapeCreateCallback != nil {
			stage.OnAfterPartShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *System:
		if stage.OnAfterSystemCreateCallback != nil {
			stage.OnAfterSystemCreateCallback.OnAfterCreate(stage, target)
		}
	case *SystemShape:
		if stage.OnAfterSystemShapeCreateCallback != nil {
			stage.OnAfterSystemShapeCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *DiagramStructure:
		newTarget := any(new).(*DiagramStructure)
		if stage.OnAfterDiagramStructureUpdateCallback != nil {
			stage.OnAfterDiagramStructureUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Library:
		newTarget := any(new).(*Library)
		if stage.OnAfterLibraryUpdateCallback != nil {
			stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Link:
		newTarget := any(new).(*Link)
		if stage.OnAfterLinkUpdateCallback != nil {
			stage.OnAfterLinkUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *LinkAssociationShape:
		newTarget := any(new).(*LinkAssociationShape)
		if stage.OnAfterLinkAssociationShapeUpdateCallback != nil {
			stage.OnAfterLinkAssociationShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Part:
		newTarget := any(new).(*Part)
		if stage.OnAfterPartUpdateCallback != nil {
			stage.OnAfterPartUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PartShape:
		newTarget := any(new).(*PartShape)
		if stage.OnAfterPartShapeUpdateCallback != nil {
			stage.OnAfterPartShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *System:
		newTarget := any(new).(*System)
		if stage.OnAfterSystemUpdateCallback != nil {
			stage.OnAfterSystemUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SystemShape:
		newTarget := any(new).(*SystemShape)
		if stage.OnAfterSystemShapeUpdateCallback != nil {
			stage.OnAfterSystemShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *DiagramStructure:
		if stage.OnAfterDiagramStructureDeleteCallback != nil {
			staged := any(staged).(*DiagramStructure)
			stage.OnAfterDiagramStructureDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Library:
		if stage.OnAfterLibraryDeleteCallback != nil {
			staged := any(staged).(*Library)
			stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Link:
		if stage.OnAfterLinkDeleteCallback != nil {
			staged := any(staged).(*Link)
			stage.OnAfterLinkDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *LinkAssociationShape:
		if stage.OnAfterLinkAssociationShapeDeleteCallback != nil {
			staged := any(staged).(*LinkAssociationShape)
			stage.OnAfterLinkAssociationShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Part:
		if stage.OnAfterPartDeleteCallback != nil {
			staged := any(staged).(*Part)
			stage.OnAfterPartDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PartShape:
		if stage.OnAfterPartShapeDeleteCallback != nil {
			staged := any(staged).(*PartShape)
			stage.OnAfterPartShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *System:
		if stage.OnAfterSystemDeleteCallback != nil {
			staged := any(staged).(*System)
			stage.OnAfterSystemDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SystemShape:
		if stage.OnAfterSystemShapeDeleteCallback != nil {
			staged := any(staged).(*SystemShape)
			stage.OnAfterSystemShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *DiagramStructure:
		if stage.OnAfterDiagramStructureReadCallback != nil {
			stage.OnAfterDiagramStructureReadCallback.OnAfterRead(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryReadCallback != nil {
			stage.OnAfterLibraryReadCallback.OnAfterRead(stage, target)
		}
	case *Link:
		if stage.OnAfterLinkReadCallback != nil {
			stage.OnAfterLinkReadCallback.OnAfterRead(stage, target)
		}
	case *LinkAssociationShape:
		if stage.OnAfterLinkAssociationShapeReadCallback != nil {
			stage.OnAfterLinkAssociationShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Part:
		if stage.OnAfterPartReadCallback != nil {
			stage.OnAfterPartReadCallback.OnAfterRead(stage, target)
		}
	case *PartShape:
		if stage.OnAfterPartShapeReadCallback != nil {
			stage.OnAfterPartShapeReadCallback.OnAfterRead(stage, target)
		}
	case *System:
		if stage.OnAfterSystemReadCallback != nil {
			stage.OnAfterSystemReadCallback.OnAfterRead(stage, target)
		}
	case *SystemShape:
		if stage.OnAfterSystemShapeReadCallback != nil {
			stage.OnAfterSystemShapeReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *DiagramStructure:
		stage.OnAfterDiagramStructureUpdateCallback = any(callback).(OnAfterUpdateInterface[DiagramStructure])
	case *Library:
		stage.OnAfterLibraryUpdateCallback = any(callback).(OnAfterUpdateInterface[Library])
	case *Link:
		stage.OnAfterLinkUpdateCallback = any(callback).(OnAfterUpdateInterface[Link])
	case *LinkAssociationShape:
		stage.OnAfterLinkAssociationShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[LinkAssociationShape])
	case *Part:
		stage.OnAfterPartUpdateCallback = any(callback).(OnAfterUpdateInterface[Part])
	case *PartShape:
		stage.OnAfterPartShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[PartShape])
	case *System:
		stage.OnAfterSystemUpdateCallback = any(callback).(OnAfterUpdateInterface[System])
	case *SystemShape:
		stage.OnAfterSystemShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[SystemShape])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *DiagramStructure:
		stage.OnAfterDiagramStructureCreateCallback = any(callback).(OnAfterCreateInterface[DiagramStructure])
	case *Library:
		stage.OnAfterLibraryCreateCallback = any(callback).(OnAfterCreateInterface[Library])
	case *Link:
		stage.OnAfterLinkCreateCallback = any(callback).(OnAfterCreateInterface[Link])
	case *LinkAssociationShape:
		stage.OnAfterLinkAssociationShapeCreateCallback = any(callback).(OnAfterCreateInterface[LinkAssociationShape])
	case *Part:
		stage.OnAfterPartCreateCallback = any(callback).(OnAfterCreateInterface[Part])
	case *PartShape:
		stage.OnAfterPartShapeCreateCallback = any(callback).(OnAfterCreateInterface[PartShape])
	case *System:
		stage.OnAfterSystemCreateCallback = any(callback).(OnAfterCreateInterface[System])
	case *SystemShape:
		stage.OnAfterSystemShapeCreateCallback = any(callback).(OnAfterCreateInterface[SystemShape])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *DiagramStructure:
		stage.OnAfterDiagramStructureDeleteCallback = any(callback).(OnAfterDeleteInterface[DiagramStructure])
	case *Library:
		stage.OnAfterLibraryDeleteCallback = any(callback).(OnAfterDeleteInterface[Library])
	case *Link:
		stage.OnAfterLinkDeleteCallback = any(callback).(OnAfterDeleteInterface[Link])
	case *LinkAssociationShape:
		stage.OnAfterLinkAssociationShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[LinkAssociationShape])
	case *Part:
		stage.OnAfterPartDeleteCallback = any(callback).(OnAfterDeleteInterface[Part])
	case *PartShape:
		stage.OnAfterPartShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[PartShape])
	case *System:
		stage.OnAfterSystemDeleteCallback = any(callback).(OnAfterDeleteInterface[System])
	case *SystemShape:
		stage.OnAfterSystemShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[SystemShape])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *DiagramStructure:
		stage.OnAfterDiagramStructureReadCallback = any(callback).(OnAfterReadInterface[DiagramStructure])
	case *Library:
		stage.OnAfterLibraryReadCallback = any(callback).(OnAfterReadInterface[Library])
	case *Link:
		stage.OnAfterLinkReadCallback = any(callback).(OnAfterReadInterface[Link])
	case *LinkAssociationShape:
		stage.OnAfterLinkAssociationShapeReadCallback = any(callback).(OnAfterReadInterface[LinkAssociationShape])
	case *Part:
		stage.OnAfterPartReadCallback = any(callback).(OnAfterReadInterface[Part])
	case *PartShape:
		stage.OnAfterPartShapeReadCallback = any(callback).(OnAfterReadInterface[PartShape])
	case *System:
		stage.OnAfterSystemReadCallback = any(callback).(OnAfterReadInterface[System])
	case *SystemShape:
		stage.OnAfterSystemShapeReadCallback = any(callback).(OnAfterReadInterface[SystemShape])
	}
}
