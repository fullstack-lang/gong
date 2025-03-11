// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Arrow:
		if stage.OnAfterArrowCreateCallback != nil {
			stage.OnAfterArrowCreateCallback.OnAfterCreate(stage, target)
		}
	case *Bar:
		if stage.OnAfterBarCreateCallback != nil {
			stage.OnAfterBarCreateCallback.OnAfterCreate(stage, target)
		}
	case *Gantt:
		if stage.OnAfterGanttCreateCallback != nil {
			stage.OnAfterGanttCreateCallback.OnAfterCreate(stage, target)
		}
	case *Group:
		if stage.OnAfterGroupCreateCallback != nil {
			stage.OnAfterGroupCreateCallback.OnAfterCreate(stage, target)
		}
	case *Lane:
		if stage.OnAfterLaneCreateCallback != nil {
			stage.OnAfterLaneCreateCallback.OnAfterCreate(stage, target)
		}
	case *LaneUse:
		if stage.OnAfterLaneUseCreateCallback != nil {
			stage.OnAfterLaneUseCreateCallback.OnAfterCreate(stage, target)
		}
	case *Milestone:
		if stage.OnAfterMilestoneCreateCallback != nil {
			stage.OnAfterMilestoneCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *Arrow:
		newTarget := any(new).(*Arrow)
		if stage.OnAfterArrowUpdateCallback != nil {
			stage.OnAfterArrowUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Bar:
		newTarget := any(new).(*Bar)
		if stage.OnAfterBarUpdateCallback != nil {
			stage.OnAfterBarUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Gantt:
		newTarget := any(new).(*Gantt)
		if stage.OnAfterGanttUpdateCallback != nil {
			stage.OnAfterGanttUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Group:
		newTarget := any(new).(*Group)
		if stage.OnAfterGroupUpdateCallback != nil {
			stage.OnAfterGroupUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Lane:
		newTarget := any(new).(*Lane)
		if stage.OnAfterLaneUpdateCallback != nil {
			stage.OnAfterLaneUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *LaneUse:
		newTarget := any(new).(*LaneUse)
		if stage.OnAfterLaneUseUpdateCallback != nil {
			stage.OnAfterLaneUseUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Milestone:
		newTarget := any(new).(*Milestone)
		if stage.OnAfterMilestoneUpdateCallback != nil {
			stage.OnAfterMilestoneUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *Arrow:
		if stage.OnAfterArrowDeleteCallback != nil {
			staged := any(staged).(*Arrow)
			stage.OnAfterArrowDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Bar:
		if stage.OnAfterBarDeleteCallback != nil {
			staged := any(staged).(*Bar)
			stage.OnAfterBarDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Gantt:
		if stage.OnAfterGanttDeleteCallback != nil {
			staged := any(staged).(*Gantt)
			stage.OnAfterGanttDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Group:
		if stage.OnAfterGroupDeleteCallback != nil {
			staged := any(staged).(*Group)
			stage.OnAfterGroupDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Lane:
		if stage.OnAfterLaneDeleteCallback != nil {
			staged := any(staged).(*Lane)
			stage.OnAfterLaneDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *LaneUse:
		if stage.OnAfterLaneUseDeleteCallback != nil {
			staged := any(staged).(*LaneUse)
			stage.OnAfterLaneUseDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Milestone:
		if stage.OnAfterMilestoneDeleteCallback != nil {
			staged := any(staged).(*Milestone)
			stage.OnAfterMilestoneDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Arrow:
		if stage.OnAfterArrowReadCallback != nil {
			stage.OnAfterArrowReadCallback.OnAfterRead(stage, target)
		}
	case *Bar:
		if stage.OnAfterBarReadCallback != nil {
			stage.OnAfterBarReadCallback.OnAfterRead(stage, target)
		}
	case *Gantt:
		if stage.OnAfterGanttReadCallback != nil {
			stage.OnAfterGanttReadCallback.OnAfterRead(stage, target)
		}
	case *Group:
		if stage.OnAfterGroupReadCallback != nil {
			stage.OnAfterGroupReadCallback.OnAfterRead(stage, target)
		}
	case *Lane:
		if stage.OnAfterLaneReadCallback != nil {
			stage.OnAfterLaneReadCallback.OnAfterRead(stage, target)
		}
	case *LaneUse:
		if stage.OnAfterLaneUseReadCallback != nil {
			stage.OnAfterLaneUseReadCallback.OnAfterRead(stage, target)
		}
	case *Milestone:
		if stage.OnAfterMilestoneReadCallback != nil {
			stage.OnAfterMilestoneReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Arrow:
		stage.OnAfterArrowUpdateCallback = any(callback).(OnAfterUpdateInterface[Arrow])
	
	case *Bar:
		stage.OnAfterBarUpdateCallback = any(callback).(OnAfterUpdateInterface[Bar])
	
	case *Gantt:
		stage.OnAfterGanttUpdateCallback = any(callback).(OnAfterUpdateInterface[Gantt])
	
	case *Group:
		stage.OnAfterGroupUpdateCallback = any(callback).(OnAfterUpdateInterface[Group])
	
	case *Lane:
		stage.OnAfterLaneUpdateCallback = any(callback).(OnAfterUpdateInterface[Lane])
	
	case *LaneUse:
		stage.OnAfterLaneUseUpdateCallback = any(callback).(OnAfterUpdateInterface[LaneUse])
	
	case *Milestone:
		stage.OnAfterMilestoneUpdateCallback = any(callback).(OnAfterUpdateInterface[Milestone])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Arrow:
		stage.OnAfterArrowCreateCallback = any(callback).(OnAfterCreateInterface[Arrow])
	
	case *Bar:
		stage.OnAfterBarCreateCallback = any(callback).(OnAfterCreateInterface[Bar])
	
	case *Gantt:
		stage.OnAfterGanttCreateCallback = any(callback).(OnAfterCreateInterface[Gantt])
	
	case *Group:
		stage.OnAfterGroupCreateCallback = any(callback).(OnAfterCreateInterface[Group])
	
	case *Lane:
		stage.OnAfterLaneCreateCallback = any(callback).(OnAfterCreateInterface[Lane])
	
	case *LaneUse:
		stage.OnAfterLaneUseCreateCallback = any(callback).(OnAfterCreateInterface[LaneUse])
	
	case *Milestone:
		stage.OnAfterMilestoneCreateCallback = any(callback).(OnAfterCreateInterface[Milestone])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Arrow:
		stage.OnAfterArrowDeleteCallback = any(callback).(OnAfterDeleteInterface[Arrow])
	
	case *Bar:
		stage.OnAfterBarDeleteCallback = any(callback).(OnAfterDeleteInterface[Bar])
	
	case *Gantt:
		stage.OnAfterGanttDeleteCallback = any(callback).(OnAfterDeleteInterface[Gantt])
	
	case *Group:
		stage.OnAfterGroupDeleteCallback = any(callback).(OnAfterDeleteInterface[Group])
	
	case *Lane:
		stage.OnAfterLaneDeleteCallback = any(callback).(OnAfterDeleteInterface[Lane])
	
	case *LaneUse:
		stage.OnAfterLaneUseDeleteCallback = any(callback).(OnAfterDeleteInterface[LaneUse])
	
	case *Milestone:
		stage.OnAfterMilestoneDeleteCallback = any(callback).(OnAfterDeleteInterface[Milestone])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *Arrow:
		stage.OnAfterArrowReadCallback = any(callback).(OnAfterReadInterface[Arrow])
	
	case *Bar:
		stage.OnAfterBarReadCallback = any(callback).(OnAfterReadInterface[Bar])
	
	case *Gantt:
		stage.OnAfterGanttReadCallback = any(callback).(OnAfterReadInterface[Gantt])
	
	case *Group:
		stage.OnAfterGroupReadCallback = any(callback).(OnAfterReadInterface[Group])
	
	case *Lane:
		stage.OnAfterLaneReadCallback = any(callback).(OnAfterReadInterface[Lane])
	
	case *LaneUse:
		stage.OnAfterLaneUseReadCallback = any(callback).(OnAfterReadInterface[LaneUse])
	
	case *Milestone:
		stage.OnAfterMilestoneReadCallback = any(callback).(OnAfterReadInterface[Milestone])
	
	}
}
