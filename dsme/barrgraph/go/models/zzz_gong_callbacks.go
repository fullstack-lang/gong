// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Category1:
		if stage.OnAfterCategory1CreateCallback != nil {
			stage.OnAfterCategory1CreateCallback.OnAfterCreate(stage, target)
		}
	case *Category1Shape:
		if stage.OnAfterCategory1ShapeCreateCallback != nil {
			stage.OnAfterCategory1ShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Category2:
		if stage.OnAfterCategory2CreateCallback != nil {
			stage.OnAfterCategory2CreateCallback.OnAfterCreate(stage, target)
		}
	case *Category2Shape:
		if stage.OnAfterCategory2ShapeCreateCallback != nil {
			stage.OnAfterCategory2ShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Category3:
		if stage.OnAfterCategory3CreateCallback != nil {
			stage.OnAfterCategory3CreateCallback.OnAfterCreate(stage, target)
		}
	case *Category3Shape:
		if stage.OnAfterCategory3ShapeCreateCallback != nil {
			stage.OnAfterCategory3ShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ControlPointShape:
		if stage.OnAfterControlPointShapeCreateCallback != nil {
			stage.OnAfterControlPointShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Desk:
		if stage.OnAfterDeskCreateCallback != nil {
			stage.OnAfterDeskCreateCallback.OnAfterCreate(stage, target)
		}
	case *Diagram:
		if stage.OnAfterDiagramCreateCallback != nil {
			stage.OnAfterDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *Influence:
		if stage.OnAfterInfluenceCreateCallback != nil {
			stage.OnAfterInfluenceCreateCallback.OnAfterCreate(stage, target)
		}
	case *InfluenceShape:
		if stage.OnAfterInfluenceShapeCreateCallback != nil {
			stage.OnAfterInfluenceShapeCreateCallback.OnAfterCreate(stage, target)
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
	case *Category1:
		newTarget := any(new).(*Category1)
		if stage.OnAfterCategory1UpdateCallback != nil {
			stage.OnAfterCategory1UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Category1Shape:
		newTarget := any(new).(*Category1Shape)
		if stage.OnAfterCategory1ShapeUpdateCallback != nil {
			stage.OnAfterCategory1ShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Category2:
		newTarget := any(new).(*Category2)
		if stage.OnAfterCategory2UpdateCallback != nil {
			stage.OnAfterCategory2UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Category2Shape:
		newTarget := any(new).(*Category2Shape)
		if stage.OnAfterCategory2ShapeUpdateCallback != nil {
			stage.OnAfterCategory2ShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Category3:
		newTarget := any(new).(*Category3)
		if stage.OnAfterCategory3UpdateCallback != nil {
			stage.OnAfterCategory3UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Category3Shape:
		newTarget := any(new).(*Category3Shape)
		if stage.OnAfterCategory3ShapeUpdateCallback != nil {
			stage.OnAfterCategory3ShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ControlPointShape:
		newTarget := any(new).(*ControlPointShape)
		if stage.OnAfterControlPointShapeUpdateCallback != nil {
			stage.OnAfterControlPointShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Desk:
		newTarget := any(new).(*Desk)
		if stage.OnAfterDeskUpdateCallback != nil {
			stage.OnAfterDeskUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Diagram:
		newTarget := any(new).(*Diagram)
		if stage.OnAfterDiagramUpdateCallback != nil {
			stage.OnAfterDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Influence:
		newTarget := any(new).(*Influence)
		if stage.OnAfterInfluenceUpdateCallback != nil {
			stage.OnAfterInfluenceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *InfluenceShape:
		newTarget := any(new).(*InfluenceShape)
		if stage.OnAfterInfluenceShapeUpdateCallback != nil {
			stage.OnAfterInfluenceShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Category1:
		if stage.OnAfterCategory1DeleteCallback != nil {
			staged := any(staged).(*Category1)
			stage.OnAfterCategory1DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Category1Shape:
		if stage.OnAfterCategory1ShapeDeleteCallback != nil {
			staged := any(staged).(*Category1Shape)
			stage.OnAfterCategory1ShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Category2:
		if stage.OnAfterCategory2DeleteCallback != nil {
			staged := any(staged).(*Category2)
			stage.OnAfterCategory2DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Category2Shape:
		if stage.OnAfterCategory2ShapeDeleteCallback != nil {
			staged := any(staged).(*Category2Shape)
			stage.OnAfterCategory2ShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Category3:
		if stage.OnAfterCategory3DeleteCallback != nil {
			staged := any(staged).(*Category3)
			stage.OnAfterCategory3DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Category3Shape:
		if stage.OnAfterCategory3ShapeDeleteCallback != nil {
			staged := any(staged).(*Category3Shape)
			stage.OnAfterCategory3ShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ControlPointShape:
		if stage.OnAfterControlPointShapeDeleteCallback != nil {
			staged := any(staged).(*ControlPointShape)
			stage.OnAfterControlPointShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Desk:
		if stage.OnAfterDeskDeleteCallback != nil {
			staged := any(staged).(*Desk)
			stage.OnAfterDeskDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Diagram:
		if stage.OnAfterDiagramDeleteCallback != nil {
			staged := any(staged).(*Diagram)
			stage.OnAfterDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Influence:
		if stage.OnAfterInfluenceDeleteCallback != nil {
			staged := any(staged).(*Influence)
			stage.OnAfterInfluenceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *InfluenceShape:
		if stage.OnAfterInfluenceShapeDeleteCallback != nil {
			staged := any(staged).(*InfluenceShape)
			stage.OnAfterInfluenceShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Category1:
		if stage.OnAfterCategory1ReadCallback != nil {
			stage.OnAfterCategory1ReadCallback.OnAfterRead(stage, target)
		}
	case *Category1Shape:
		if stage.OnAfterCategory1ShapeReadCallback != nil {
			stage.OnAfterCategory1ShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Category2:
		if stage.OnAfterCategory2ReadCallback != nil {
			stage.OnAfterCategory2ReadCallback.OnAfterRead(stage, target)
		}
	case *Category2Shape:
		if stage.OnAfterCategory2ShapeReadCallback != nil {
			stage.OnAfterCategory2ShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Category3:
		if stage.OnAfterCategory3ReadCallback != nil {
			stage.OnAfterCategory3ReadCallback.OnAfterRead(stage, target)
		}
	case *Category3Shape:
		if stage.OnAfterCategory3ShapeReadCallback != nil {
			stage.OnAfterCategory3ShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ControlPointShape:
		if stage.OnAfterControlPointShapeReadCallback != nil {
			stage.OnAfterControlPointShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Desk:
		if stage.OnAfterDeskReadCallback != nil {
			stage.OnAfterDeskReadCallback.OnAfterRead(stage, target)
		}
	case *Diagram:
		if stage.OnAfterDiagramReadCallback != nil {
			stage.OnAfterDiagramReadCallback.OnAfterRead(stage, target)
		}
	case *Influence:
		if stage.OnAfterInfluenceReadCallback != nil {
			stage.OnAfterInfluenceReadCallback.OnAfterRead(stage, target)
		}
	case *InfluenceShape:
		if stage.OnAfterInfluenceShapeReadCallback != nil {
			stage.OnAfterInfluenceShapeReadCallback.OnAfterRead(stage, target)
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
	case *Category1:
		stage.OnAfterCategory1UpdateCallback = any(callback).(OnAfterUpdateInterface[Category1])
	case *Category1Shape:
		stage.OnAfterCategory1ShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[Category1Shape])
	case *Category2:
		stage.OnAfterCategory2UpdateCallback = any(callback).(OnAfterUpdateInterface[Category2])
	case *Category2Shape:
		stage.OnAfterCategory2ShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[Category2Shape])
	case *Category3:
		stage.OnAfterCategory3UpdateCallback = any(callback).(OnAfterUpdateInterface[Category3])
	case *Category3Shape:
		stage.OnAfterCategory3ShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[Category3Shape])
	case *ControlPointShape:
		stage.OnAfterControlPointShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ControlPointShape])
	case *Desk:
		stage.OnAfterDeskUpdateCallback = any(callback).(OnAfterUpdateInterface[Desk])
	case *Diagram:
		stage.OnAfterDiagramUpdateCallback = any(callback).(OnAfterUpdateInterface[Diagram])
	case *Influence:
		stage.OnAfterInfluenceUpdateCallback = any(callback).(OnAfterUpdateInterface[Influence])
	case *InfluenceShape:
		stage.OnAfterInfluenceShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[InfluenceShape])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Category1:
		stage.OnAfterCategory1CreateCallback = any(callback).(OnAfterCreateInterface[Category1])
	case *Category1Shape:
		stage.OnAfterCategory1ShapeCreateCallback = any(callback).(OnAfterCreateInterface[Category1Shape])
	case *Category2:
		stage.OnAfterCategory2CreateCallback = any(callback).(OnAfterCreateInterface[Category2])
	case *Category2Shape:
		stage.OnAfterCategory2ShapeCreateCallback = any(callback).(OnAfterCreateInterface[Category2Shape])
	case *Category3:
		stage.OnAfterCategory3CreateCallback = any(callback).(OnAfterCreateInterface[Category3])
	case *Category3Shape:
		stage.OnAfterCategory3ShapeCreateCallback = any(callback).(OnAfterCreateInterface[Category3Shape])
	case *ControlPointShape:
		stage.OnAfterControlPointShapeCreateCallback = any(callback).(OnAfterCreateInterface[ControlPointShape])
	case *Desk:
		stage.OnAfterDeskCreateCallback = any(callback).(OnAfterCreateInterface[Desk])
	case *Diagram:
		stage.OnAfterDiagramCreateCallback = any(callback).(OnAfterCreateInterface[Diagram])
	case *Influence:
		stage.OnAfterInfluenceCreateCallback = any(callback).(OnAfterCreateInterface[Influence])
	case *InfluenceShape:
		stage.OnAfterInfluenceShapeCreateCallback = any(callback).(OnAfterCreateInterface[InfluenceShape])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Category1:
		stage.OnAfterCategory1DeleteCallback = any(callback).(OnAfterDeleteInterface[Category1])
	case *Category1Shape:
		stage.OnAfterCategory1ShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[Category1Shape])
	case *Category2:
		stage.OnAfterCategory2DeleteCallback = any(callback).(OnAfterDeleteInterface[Category2])
	case *Category2Shape:
		stage.OnAfterCategory2ShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[Category2Shape])
	case *Category3:
		stage.OnAfterCategory3DeleteCallback = any(callback).(OnAfterDeleteInterface[Category3])
	case *Category3Shape:
		stage.OnAfterCategory3ShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[Category3Shape])
	case *ControlPointShape:
		stage.OnAfterControlPointShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ControlPointShape])
	case *Desk:
		stage.OnAfterDeskDeleteCallback = any(callback).(OnAfterDeleteInterface[Desk])
	case *Diagram:
		stage.OnAfterDiagramDeleteCallback = any(callback).(OnAfterDeleteInterface[Diagram])
	case *Influence:
		stage.OnAfterInfluenceDeleteCallback = any(callback).(OnAfterDeleteInterface[Influence])
	case *InfluenceShape:
		stage.OnAfterInfluenceShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[InfluenceShape])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Category1:
		stage.OnAfterCategory1ReadCallback = any(callback).(OnAfterReadInterface[Category1])
	case *Category1Shape:
		stage.OnAfterCategory1ShapeReadCallback = any(callback).(OnAfterReadInterface[Category1Shape])
	case *Category2:
		stage.OnAfterCategory2ReadCallback = any(callback).(OnAfterReadInterface[Category2])
	case *Category2Shape:
		stage.OnAfterCategory2ShapeReadCallback = any(callback).(OnAfterReadInterface[Category2Shape])
	case *Category3:
		stage.OnAfterCategory3ReadCallback = any(callback).(OnAfterReadInterface[Category3])
	case *Category3Shape:
		stage.OnAfterCategory3ShapeReadCallback = any(callback).(OnAfterReadInterface[Category3Shape])
	case *ControlPointShape:
		stage.OnAfterControlPointShapeReadCallback = any(callback).(OnAfterReadInterface[ControlPointShape])
	case *Desk:
		stage.OnAfterDeskReadCallback = any(callback).(OnAfterReadInterface[Desk])
	case *Diagram:
		stage.OnAfterDiagramReadCallback = any(callback).(OnAfterReadInterface[Diagram])
	case *Influence:
		stage.OnAfterInfluenceReadCallback = any(callback).(OnAfterReadInterface[Influence])
	case *InfluenceShape:
		stage.OnAfterInfluenceShapeReadCallback = any(callback).(OnAfterReadInterface[InfluenceShape])
	}
}
