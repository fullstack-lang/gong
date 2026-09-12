// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *DisplaySelection:
		if stage.OnAfterDisplaySelectionCreateCallback != nil {
			stage.OnAfterDisplaySelectionCreateCallback.OnAfterCreate(stage, target)
		}
	case *XLCell:
		if stage.OnAfterXLCellCreateCallback != nil {
			stage.OnAfterXLCellCreateCallback.OnAfterCreate(stage, target)
		}
	case *XLFile:
		if stage.OnAfterXLFileCreateCallback != nil {
			stage.OnAfterXLFileCreateCallback.OnAfterCreate(stage, target)
		}
	case *XLRow:
		if stage.OnAfterXLRowCreateCallback != nil {
			stage.OnAfterXLRowCreateCallback.OnAfterCreate(stage, target)
		}
	case *XLSheet:
		if stage.OnAfterXLSheetCreateCallback != nil {
			stage.OnAfterXLSheetCreateCallback.OnAfterCreate(stage, target)
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
	case *DisplaySelection:
		newTarget := any(new).(*DisplaySelection)
		if stage.OnAfterDisplaySelectionUpdateCallback != nil {
			stage.OnAfterDisplaySelectionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *XLCell:
		newTarget := any(new).(*XLCell)
		if stage.OnAfterXLCellUpdateCallback != nil {
			stage.OnAfterXLCellUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *XLFile:
		newTarget := any(new).(*XLFile)
		if stage.OnAfterXLFileUpdateCallback != nil {
			stage.OnAfterXLFileUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *XLRow:
		newTarget := any(new).(*XLRow)
		if stage.OnAfterXLRowUpdateCallback != nil {
			stage.OnAfterXLRowUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *XLSheet:
		newTarget := any(new).(*XLSheet)
		if stage.OnAfterXLSheetUpdateCallback != nil {
			stage.OnAfterXLSheetUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
	case *DisplaySelection:
		if stage.OnAfterDisplaySelectionDeleteCallback != nil {
			staged := any(staged).(*DisplaySelection)
			stage.OnAfterDisplaySelectionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *XLCell:
		if stage.OnAfterXLCellDeleteCallback != nil {
			staged := any(staged).(*XLCell)
			stage.OnAfterXLCellDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *XLFile:
		if stage.OnAfterXLFileDeleteCallback != nil {
			staged := any(staged).(*XLFile)
			stage.OnAfterXLFileDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *XLRow:
		if stage.OnAfterXLRowDeleteCallback != nil {
			staged := any(staged).(*XLRow)
			stage.OnAfterXLRowDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *XLSheet:
		if stage.OnAfterXLSheetDeleteCallback != nil {
			staged := any(staged).(*XLSheet)
			stage.OnAfterXLSheetDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
