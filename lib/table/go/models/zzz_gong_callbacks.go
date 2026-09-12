// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Button:
		if stage.OnAfterButtonCreateCallback != nil {
			stage.OnAfterButtonCreateCallback.OnAfterCreate(stage, target)
		}
	case *Cell:
		if stage.OnAfterCellCreateCallback != nil {
			stage.OnAfterCellCreateCallback.OnAfterCreate(stage, target)
		}
	case *CellBoolean:
		if stage.OnAfterCellBooleanCreateCallback != nil {
			stage.OnAfterCellBooleanCreateCallback.OnAfterCreate(stage, target)
		}
	case *CellFloat64:
		if stage.OnAfterCellFloat64CreateCallback != nil {
			stage.OnAfterCellFloat64CreateCallback.OnAfterCreate(stage, target)
		}
	case *CellIcon:
		if stage.OnAfterCellIconCreateCallback != nil {
			stage.OnAfterCellIconCreateCallback.OnAfterCreate(stage, target)
		}
	case *CellInt:
		if stage.OnAfterCellIntCreateCallback != nil {
			stage.OnAfterCellIntCreateCallback.OnAfterCreate(stage, target)
		}
	case *CellString:
		if stage.OnAfterCellStringCreateCallback != nil {
			stage.OnAfterCellStringCreateCallback.OnAfterCreate(stage, target)
		}
	case *DisplayedColumn:
		if stage.OnAfterDisplayedColumnCreateCallback != nil {
			stage.OnAfterDisplayedColumnCreateCallback.OnAfterCreate(stage, target)
		}
	case *Row:
		if stage.OnAfterRowCreateCallback != nil {
			stage.OnAfterRowCreateCallback.OnAfterCreate(stage, target)
		}
	case *SVGIcon:
		if stage.OnAfterSVGIconCreateCallback != nil {
			stage.OnAfterSVGIconCreateCallback.OnAfterCreate(stage, target)
		}
	case *Table:
		if stage.OnAfterTableCreateCallback != nil {
			stage.OnAfterTableCreateCallback.OnAfterCreate(stage, target)
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
	case *Button:
		newTarget := any(new).(*Button)
		if stage.OnAfterButtonUpdateCallback != nil {
			stage.OnAfterButtonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Cell:
		newTarget := any(new).(*Cell)
		if stage.OnAfterCellUpdateCallback != nil {
			stage.OnAfterCellUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CellBoolean:
		newTarget := any(new).(*CellBoolean)
		if stage.OnAfterCellBooleanUpdateCallback != nil {
			stage.OnAfterCellBooleanUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CellFloat64:
		newTarget := any(new).(*CellFloat64)
		if stage.OnAfterCellFloat64UpdateCallback != nil {
			stage.OnAfterCellFloat64UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CellIcon:
		newTarget := any(new).(*CellIcon)
		if stage.OnAfterCellIconUpdateCallback != nil {
			stage.OnAfterCellIconUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CellInt:
		newTarget := any(new).(*CellInt)
		if stage.OnAfterCellIntUpdateCallback != nil {
			stage.OnAfterCellIntUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *CellString:
		newTarget := any(new).(*CellString)
		if stage.OnAfterCellStringUpdateCallback != nil {
			stage.OnAfterCellStringUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DisplayedColumn:
		newTarget := any(new).(*DisplayedColumn)
		if stage.OnAfterDisplayedColumnUpdateCallback != nil {
			stage.OnAfterDisplayedColumnUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Row:
		newTarget := any(new).(*Row)
		if stage.OnAfterRowUpdateCallback != nil {
			stage.OnAfterRowUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SVGIcon:
		newTarget := any(new).(*SVGIcon)
		if stage.OnAfterSVGIconUpdateCallback != nil {
			stage.OnAfterSVGIconUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Table:
		newTarget := any(new).(*Table)
		if stage.OnAfterTableUpdateCallback != nil {
			stage.OnAfterTableUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *Button:
		if stage.OnAfterButtonDeleteCallback != nil {
			staged := any(staged).(*Button)
			stage.OnAfterButtonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Cell:
		if stage.OnAfterCellDeleteCallback != nil {
			staged := any(staged).(*Cell)
			stage.OnAfterCellDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CellBoolean:
		if stage.OnAfterCellBooleanDeleteCallback != nil {
			staged := any(staged).(*CellBoolean)
			stage.OnAfterCellBooleanDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CellFloat64:
		if stage.OnAfterCellFloat64DeleteCallback != nil {
			staged := any(staged).(*CellFloat64)
			stage.OnAfterCellFloat64DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CellIcon:
		if stage.OnAfterCellIconDeleteCallback != nil {
			staged := any(staged).(*CellIcon)
			stage.OnAfterCellIconDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CellInt:
		if stage.OnAfterCellIntDeleteCallback != nil {
			staged := any(staged).(*CellInt)
			stage.OnAfterCellIntDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *CellString:
		if stage.OnAfterCellStringDeleteCallback != nil {
			staged := any(staged).(*CellString)
			stage.OnAfterCellStringDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DisplayedColumn:
		if stage.OnAfterDisplayedColumnDeleteCallback != nil {
			staged := any(staged).(*DisplayedColumn)
			stage.OnAfterDisplayedColumnDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Row:
		if stage.OnAfterRowDeleteCallback != nil {
			staged := any(staged).(*Row)
			stage.OnAfterRowDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SVGIcon:
		if stage.OnAfterSVGIconDeleteCallback != nil {
			staged := any(staged).(*SVGIcon)
			stage.OnAfterSVGIconDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Table:
		if stage.OnAfterTableDeleteCallback != nil {
			staged := any(staged).(*Table)
			stage.OnAfterTableDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
