// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AxesShape:
		if stage.OnAfterAxesShapeCreateCallback != nil {
			stage.OnAfterAxesShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
		}
	case *Plant:
		if stage.OnAfterPlantCreateCallback != nil {
			stage.OnAfterPlantCreateCallback.OnAfterCreate(stage, target)
		}
	case *PlantDiagram:
		if stage.OnAfterPlantDiagramCreateCallback != nil {
			stage.OnAfterPlantDiagramCreateCallback.OnAfterCreate(stage, target)
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
	case *AxesShape:
		newTarget := any(new).(*AxesShape)
		if stage.OnAfterAxesShapeUpdateCallback != nil {
			stage.OnAfterAxesShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Library:
		newTarget := any(new).(*Library)
		if stage.OnAfterLibraryUpdateCallback != nil {
			stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Plant:
		newTarget := any(new).(*Plant)
		if stage.OnAfterPlantUpdateCallback != nil {
			stage.OnAfterPlantUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *PlantDiagram:
		newTarget := any(new).(*PlantDiagram)
		if stage.OnAfterPlantDiagramUpdateCallback != nil {
			stage.OnAfterPlantDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *AxesShape:
		if stage.OnAfterAxesShapeDeleteCallback != nil {
			staged := any(staged).(*AxesShape)
			stage.OnAfterAxesShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Library:
		if stage.OnAfterLibraryDeleteCallback != nil {
			staged := any(staged).(*Library)
			stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Plant:
		if stage.OnAfterPlantDeleteCallback != nil {
			staged := any(staged).(*Plant)
			stage.OnAfterPlantDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *PlantDiagram:
		if stage.OnAfterPlantDiagramDeleteCallback != nil {
			staged := any(staged).(*PlantDiagram)
			stage.OnAfterPlantDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AxesShape:
		if stage.OnAfterAxesShapeReadCallback != nil {
			stage.OnAfterAxesShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryReadCallback != nil {
			stage.OnAfterLibraryReadCallback.OnAfterRead(stage, target)
		}
	case *Plant:
		if stage.OnAfterPlantReadCallback != nil {
			stage.OnAfterPlantReadCallback.OnAfterRead(stage, target)
		}
	case *PlantDiagram:
		if stage.OnAfterPlantDiagramReadCallback != nil {
			stage.OnAfterPlantDiagramReadCallback.OnAfterRead(stage, target)
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
	case *AxesShape:
		stage.OnAfterAxesShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[AxesShape])
	case *Library:
		stage.OnAfterLibraryUpdateCallback = any(callback).(OnAfterUpdateInterface[Library])
	case *Plant:
		stage.OnAfterPlantUpdateCallback = any(callback).(OnAfterUpdateInterface[Plant])
	case *PlantDiagram:
		stage.OnAfterPlantDiagramUpdateCallback = any(callback).(OnAfterUpdateInterface[PlantDiagram])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AxesShape:
		stage.OnAfterAxesShapeCreateCallback = any(callback).(OnAfterCreateInterface[AxesShape])
	case *Library:
		stage.OnAfterLibraryCreateCallback = any(callback).(OnAfterCreateInterface[Library])
	case *Plant:
		stage.OnAfterPlantCreateCallback = any(callback).(OnAfterCreateInterface[Plant])
	case *PlantDiagram:
		stage.OnAfterPlantDiagramCreateCallback = any(callback).(OnAfterCreateInterface[PlantDiagram])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AxesShape:
		stage.OnAfterAxesShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[AxesShape])
	case *Library:
		stage.OnAfterLibraryDeleteCallback = any(callback).(OnAfterDeleteInterface[Library])
	case *Plant:
		stage.OnAfterPlantDeleteCallback = any(callback).(OnAfterDeleteInterface[Plant])
	case *PlantDiagram:
		stage.OnAfterPlantDiagramDeleteCallback = any(callback).(OnAfterDeleteInterface[PlantDiagram])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AxesShape:
		stage.OnAfterAxesShapeReadCallback = any(callback).(OnAfterReadInterface[AxesShape])
	case *Library:
		stage.OnAfterLibraryReadCallback = any(callback).(OnAfterReadInterface[Library])
	case *Plant:
		stage.OnAfterPlantReadCallback = any(callback).(OnAfterReadInterface[Plant])
	case *PlantDiagram:
		stage.OnAfterPlantDiagramReadCallback = any(callback).(OnAfterReadInterface[PlantDiagram])
	}
}
