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
	case *CircleGridShape:
		if stage.OnAfterCircleGridShapeCreateCallback != nil {
			stage.OnAfterCircleGridShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GridPathShape:
		if stage.OnAfterGridPathShapeCreateCallback != nil {
			stage.OnAfterGridPathShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *GrowthVectorShape:
		if stage.OnAfterGrowthVectorShapeCreateCallback != nil {
			stage.OnAfterGrowthVectorShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
		}
	case *NextCircleShape:
		if stage.OnAfterNextCircleShapeCreateCallback != nil {
			stage.OnAfterNextCircleShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Plant:
		if stage.OnAfterPlantCreateCallback != nil {
			stage.OnAfterPlantCreateCallback.OnAfterCreate(stage, target)
		}
	case *PlantDiagram:
		if stage.OnAfterPlantDiagramCreateCallback != nil {
			stage.OnAfterPlantDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *ReferenceRhombus:
		if stage.OnAfterReferenceRhombusCreateCallback != nil {
			stage.OnAfterReferenceRhombusCreateCallback.OnAfterCreate(stage, target)
		}
	case *RhombusGridShape:
		if stage.OnAfterRhombusGridShapeCreateCallback != nil {
			stage.OnAfterRhombusGridShapeCreateCallback.OnAfterCreate(stage, target)
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
	case *CircleGridShape:
		newTarget := any(new).(*CircleGridShape)
		if stage.OnAfterCircleGridShapeUpdateCallback != nil {
			stage.OnAfterCircleGridShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GridPathShape:
		newTarget := any(new).(*GridPathShape)
		if stage.OnAfterGridPathShapeUpdateCallback != nil {
			stage.OnAfterGridPathShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *GrowthVectorShape:
		newTarget := any(new).(*GrowthVectorShape)
		if stage.OnAfterGrowthVectorShapeUpdateCallback != nil {
			stage.OnAfterGrowthVectorShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Library:
		newTarget := any(new).(*Library)
		if stage.OnAfterLibraryUpdateCallback != nil {
			stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *NextCircleShape:
		newTarget := any(new).(*NextCircleShape)
		if stage.OnAfterNextCircleShapeUpdateCallback != nil {
			stage.OnAfterNextCircleShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *ReferenceRhombus:
		newTarget := any(new).(*ReferenceRhombus)
		if stage.OnAfterReferenceRhombusUpdateCallback != nil {
			stage.OnAfterReferenceRhombusUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RhombusGridShape:
		newTarget := any(new).(*RhombusGridShape)
		if stage.OnAfterRhombusGridShapeUpdateCallback != nil {
			stage.OnAfterRhombusGridShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *CircleGridShape:
		if stage.OnAfterCircleGridShapeDeleteCallback != nil {
			staged := any(staged).(*CircleGridShape)
			stage.OnAfterCircleGridShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GridPathShape:
		if stage.OnAfterGridPathShapeDeleteCallback != nil {
			staged := any(staged).(*GridPathShape)
			stage.OnAfterGridPathShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *GrowthVectorShape:
		if stage.OnAfterGrowthVectorShapeDeleteCallback != nil {
			staged := any(staged).(*GrowthVectorShape)
			stage.OnAfterGrowthVectorShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Library:
		if stage.OnAfterLibraryDeleteCallback != nil {
			staged := any(staged).(*Library)
			stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *NextCircleShape:
		if stage.OnAfterNextCircleShapeDeleteCallback != nil {
			staged := any(staged).(*NextCircleShape)
			stage.OnAfterNextCircleShapeDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *ReferenceRhombus:
		if stage.OnAfterReferenceRhombusDeleteCallback != nil {
			staged := any(staged).(*ReferenceRhombus)
			stage.OnAfterReferenceRhombusDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RhombusGridShape:
		if stage.OnAfterRhombusGridShapeDeleteCallback != nil {
			staged := any(staged).(*RhombusGridShape)
			stage.OnAfterRhombusGridShapeDeleteCallback.OnAfterDelete(stage, staged, front)
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
	case *CircleGridShape:
		if stage.OnAfterCircleGridShapeReadCallback != nil {
			stage.OnAfterCircleGridShapeReadCallback.OnAfterRead(stage, target)
		}
	case *GridPathShape:
		if stage.OnAfterGridPathShapeReadCallback != nil {
			stage.OnAfterGridPathShapeReadCallback.OnAfterRead(stage, target)
		}
	case *GrowthVectorShape:
		if stage.OnAfterGrowthVectorShapeReadCallback != nil {
			stage.OnAfterGrowthVectorShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Library:
		if stage.OnAfterLibraryReadCallback != nil {
			stage.OnAfterLibraryReadCallback.OnAfterRead(stage, target)
		}
	case *NextCircleShape:
		if stage.OnAfterNextCircleShapeReadCallback != nil {
			stage.OnAfterNextCircleShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Plant:
		if stage.OnAfterPlantReadCallback != nil {
			stage.OnAfterPlantReadCallback.OnAfterRead(stage, target)
		}
	case *PlantDiagram:
		if stage.OnAfterPlantDiagramReadCallback != nil {
			stage.OnAfterPlantDiagramReadCallback.OnAfterRead(stage, target)
		}
	case *ReferenceRhombus:
		if stage.OnAfterReferenceRhombusReadCallback != nil {
			stage.OnAfterReferenceRhombusReadCallback.OnAfterRead(stage, target)
		}
	case *RhombusGridShape:
		if stage.OnAfterRhombusGridShapeReadCallback != nil {
			stage.OnAfterRhombusGridShapeReadCallback.OnAfterRead(stage, target)
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
	case *CircleGridShape:
		stage.OnAfterCircleGridShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[CircleGridShape])
	case *GridPathShape:
		stage.OnAfterGridPathShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GridPathShape])
	case *GrowthVectorShape:
		stage.OnAfterGrowthVectorShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[GrowthVectorShape])
	case *Library:
		stage.OnAfterLibraryUpdateCallback = any(callback).(OnAfterUpdateInterface[Library])
	case *NextCircleShape:
		stage.OnAfterNextCircleShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[NextCircleShape])
	case *Plant:
		stage.OnAfterPlantUpdateCallback = any(callback).(OnAfterUpdateInterface[Plant])
	case *PlantDiagram:
		stage.OnAfterPlantDiagramUpdateCallback = any(callback).(OnAfterUpdateInterface[PlantDiagram])
	case *ReferenceRhombus:
		stage.OnAfterReferenceRhombusUpdateCallback = any(callback).(OnAfterUpdateInterface[ReferenceRhombus])
	case *RhombusGridShape:
		stage.OnAfterRhombusGridShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[RhombusGridShape])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AxesShape:
		stage.OnAfterAxesShapeCreateCallback = any(callback).(OnAfterCreateInterface[AxesShape])
	case *CircleGridShape:
		stage.OnAfterCircleGridShapeCreateCallback = any(callback).(OnAfterCreateInterface[CircleGridShape])
	case *GridPathShape:
		stage.OnAfterGridPathShapeCreateCallback = any(callback).(OnAfterCreateInterface[GridPathShape])
	case *GrowthVectorShape:
		stage.OnAfterGrowthVectorShapeCreateCallback = any(callback).(OnAfterCreateInterface[GrowthVectorShape])
	case *Library:
		stage.OnAfterLibraryCreateCallback = any(callback).(OnAfterCreateInterface[Library])
	case *NextCircleShape:
		stage.OnAfterNextCircleShapeCreateCallback = any(callback).(OnAfterCreateInterface[NextCircleShape])
	case *Plant:
		stage.OnAfterPlantCreateCallback = any(callback).(OnAfterCreateInterface[Plant])
	case *PlantDiagram:
		stage.OnAfterPlantDiagramCreateCallback = any(callback).(OnAfterCreateInterface[PlantDiagram])
	case *ReferenceRhombus:
		stage.OnAfterReferenceRhombusCreateCallback = any(callback).(OnAfterCreateInterface[ReferenceRhombus])
	case *RhombusGridShape:
		stage.OnAfterRhombusGridShapeCreateCallback = any(callback).(OnAfterCreateInterface[RhombusGridShape])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AxesShape:
		stage.OnAfterAxesShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[AxesShape])
	case *CircleGridShape:
		stage.OnAfterCircleGridShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[CircleGridShape])
	case *GridPathShape:
		stage.OnAfterGridPathShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GridPathShape])
	case *GrowthVectorShape:
		stage.OnAfterGrowthVectorShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[GrowthVectorShape])
	case *Library:
		stage.OnAfterLibraryDeleteCallback = any(callback).(OnAfterDeleteInterface[Library])
	case *NextCircleShape:
		stage.OnAfterNextCircleShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[NextCircleShape])
	case *Plant:
		stage.OnAfterPlantDeleteCallback = any(callback).(OnAfterDeleteInterface[Plant])
	case *PlantDiagram:
		stage.OnAfterPlantDiagramDeleteCallback = any(callback).(OnAfterDeleteInterface[PlantDiagram])
	case *ReferenceRhombus:
		stage.OnAfterReferenceRhombusDeleteCallback = any(callback).(OnAfterDeleteInterface[ReferenceRhombus])
	case *RhombusGridShape:
		stage.OnAfterRhombusGridShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[RhombusGridShape])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *AxesShape:
		stage.OnAfterAxesShapeReadCallback = any(callback).(OnAfterReadInterface[AxesShape])
	case *CircleGridShape:
		stage.OnAfterCircleGridShapeReadCallback = any(callback).(OnAfterReadInterface[CircleGridShape])
	case *GridPathShape:
		stage.OnAfterGridPathShapeReadCallback = any(callback).(OnAfterReadInterface[GridPathShape])
	case *GrowthVectorShape:
		stage.OnAfterGrowthVectorShapeReadCallback = any(callback).(OnAfterReadInterface[GrowthVectorShape])
	case *Library:
		stage.OnAfterLibraryReadCallback = any(callback).(OnAfterReadInterface[Library])
	case *NextCircleShape:
		stage.OnAfterNextCircleShapeReadCallback = any(callback).(OnAfterReadInterface[NextCircleShape])
	case *Plant:
		stage.OnAfterPlantReadCallback = any(callback).(OnAfterReadInterface[Plant])
	case *PlantDiagram:
		stage.OnAfterPlantDiagramReadCallback = any(callback).(OnAfterReadInterface[PlantDiagram])
	case *ReferenceRhombus:
		stage.OnAfterReferenceRhombusReadCallback = any(callback).(OnAfterReadInterface[ReferenceRhombus])
	case *RhombusGridShape:
		stage.OnAfterRhombusGridShapeReadCallback = any(callback).(OnAfterReadInterface[RhombusGridShape])
	}
}
