// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

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

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

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

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

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

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
