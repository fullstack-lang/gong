// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (arrow *Arrow) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArrowCreateCallback != nil {
		stage.OnAfterArrowCreateCallback.OnAfterCreate(stage, arrow)
	}
}

func (arrow *Arrow) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrowUpdateCallback != nil {
		var frontArrow *Arrow
		if front != nil {
			frontArrow, _ = front.(*Arrow)
		}
		stage.OnAfterArrowUpdateCallback.OnAfterUpdate(stage, arrow, frontArrow)
	}
}

func (arrow *Arrow) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArrowDeleteCallback != nil {
		var frontArrow *Arrow
		if front != nil {
			frontArrow, _ = front.(*Arrow)
		}
		stage.OnAfterArrowDeleteCallback.OnAfterDelete(stage, arrow, frontArrow)
	}
}

func (bar *Bar) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterBarCreateCallback != nil {
		stage.OnAfterBarCreateCallback.OnAfterCreate(stage, bar)
	}
}

func (bar *Bar) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBarUpdateCallback != nil {
		var frontBar *Bar
		if front != nil {
			frontBar, _ = front.(*Bar)
		}
		stage.OnAfterBarUpdateCallback.OnAfterUpdate(stage, bar, frontBar)
	}
}

func (bar *Bar) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterBarDeleteCallback != nil {
		var frontBar *Bar
		if front != nil {
			frontBar, _ = front.(*Bar)
		}
		stage.OnAfterBarDeleteCallback.OnAfterDelete(stage, bar, frontBar)
	}
}

func (gantt *Gantt) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGanttCreateCallback != nil {
		stage.OnAfterGanttCreateCallback.OnAfterCreate(stage, gantt)
	}
}

func (gantt *Gantt) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGanttUpdateCallback != nil {
		var frontGantt *Gantt
		if front != nil {
			frontGantt, _ = front.(*Gantt)
		}
		stage.OnAfterGanttUpdateCallback.OnAfterUpdate(stage, gantt, frontGantt)
	}
}

func (gantt *Gantt) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGanttDeleteCallback != nil {
		var frontGantt *Gantt
		if front != nil {
			frontGantt, _ = front.(*Gantt)
		}
		stage.OnAfterGanttDeleteCallback.OnAfterDelete(stage, gantt, frontGantt)
	}
}

func (group *Group) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterGroupCreateCallback != nil {
		stage.OnAfterGroupCreateCallback.OnAfterCreate(stage, group)
	}
}

func (group *Group) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupUpdateCallback != nil {
		var frontGroup *Group
		if front != nil {
			frontGroup, _ = front.(*Group)
		}
		stage.OnAfterGroupUpdateCallback.OnAfterUpdate(stage, group, frontGroup)
	}
}

func (group *Group) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterGroupDeleteCallback != nil {
		var frontGroup *Group
		if front != nil {
			frontGroup, _ = front.(*Group)
		}
		stage.OnAfterGroupDeleteCallback.OnAfterDelete(stage, group, frontGroup)
	}
}

func (lane *Lane) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLaneCreateCallback != nil {
		stage.OnAfterLaneCreateCallback.OnAfterCreate(stage, lane)
	}
}

func (lane *Lane) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLaneUpdateCallback != nil {
		var frontLane *Lane
		if front != nil {
			frontLane, _ = front.(*Lane)
		}
		stage.OnAfterLaneUpdateCallback.OnAfterUpdate(stage, lane, frontLane)
	}
}

func (lane *Lane) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLaneDeleteCallback != nil {
		var frontLane *Lane
		if front != nil {
			frontLane, _ = front.(*Lane)
		}
		stage.OnAfterLaneDeleteCallback.OnAfterDelete(stage, lane, frontLane)
	}
}

func (laneuse *LaneUse) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLaneUseCreateCallback != nil {
		stage.OnAfterLaneUseCreateCallback.OnAfterCreate(stage, laneuse)
	}
}

func (laneuse *LaneUse) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLaneUseUpdateCallback != nil {
		var frontLaneUse *LaneUse
		if front != nil {
			frontLaneUse, _ = front.(*LaneUse)
		}
		stage.OnAfterLaneUseUpdateCallback.OnAfterUpdate(stage, laneuse, frontLaneUse)
	}
}

func (laneuse *LaneUse) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLaneUseDeleteCallback != nil {
		var frontLaneUse *LaneUse
		if front != nil {
			frontLaneUse, _ = front.(*LaneUse)
		}
		stage.OnAfterLaneUseDeleteCallback.OnAfterDelete(stage, laneuse, frontLaneUse)
	}
}

func (milestone *Milestone) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMilestoneCreateCallback != nil {
		stage.OnAfterMilestoneCreateCallback.OnAfterCreate(stage, milestone)
	}
}

func (milestone *Milestone) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMilestoneUpdateCallback != nil {
		var frontMilestone *Milestone
		if front != nil {
			frontMilestone, _ = front.(*Milestone)
		}
		stage.OnAfterMilestoneUpdateCallback.OnAfterUpdate(stage, milestone, frontMilestone)
	}
}

func (milestone *Milestone) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMilestoneDeleteCallback != nil {
		var frontMilestone *Milestone
		if front != nil {
			frontMilestone, _ = front.(*Milestone)
		}
		stage.OnAfterMilestoneDeleteCallback.OnAfterDelete(stage, milestone, frontMilestone)
	}
}

