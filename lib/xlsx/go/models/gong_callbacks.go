// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

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

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

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

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

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

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *DisplaySelection:
		if stage.OnAfterDisplaySelectionReadCallback != nil {
			stage.OnAfterDisplaySelectionReadCallback.OnAfterRead(stage, target)
		}
	case *XLCell:
		if stage.OnAfterXLCellReadCallback != nil {
			stage.OnAfterXLCellReadCallback.OnAfterRead(stage, target)
		}
	case *XLFile:
		if stage.OnAfterXLFileReadCallback != nil {
			stage.OnAfterXLFileReadCallback.OnAfterRead(stage, target)
		}
	case *XLRow:
		if stage.OnAfterXLRowReadCallback != nil {
			stage.OnAfterXLRowReadCallback.OnAfterRead(stage, target)
		}
	case *XLSheet:
		if stage.OnAfterXLSheetReadCallback != nil {
			stage.OnAfterXLSheetReadCallback.OnAfterRead(stage, target)
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
	case *DisplaySelection:
		stage.OnAfterDisplaySelectionUpdateCallback = any(callback).(OnAfterUpdateInterface[DisplaySelection])
	
	case *XLCell:
		stage.OnAfterXLCellUpdateCallback = any(callback).(OnAfterUpdateInterface[XLCell])
	
	case *XLFile:
		stage.OnAfterXLFileUpdateCallback = any(callback).(OnAfterUpdateInterface[XLFile])
	
	case *XLRow:
		stage.OnAfterXLRowUpdateCallback = any(callback).(OnAfterUpdateInterface[XLRow])
	
	case *XLSheet:
		stage.OnAfterXLSheetUpdateCallback = any(callback).(OnAfterUpdateInterface[XLSheet])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *DisplaySelection:
		stage.OnAfterDisplaySelectionCreateCallback = any(callback).(OnAfterCreateInterface[DisplaySelection])
	
	case *XLCell:
		stage.OnAfterXLCellCreateCallback = any(callback).(OnAfterCreateInterface[XLCell])
	
	case *XLFile:
		stage.OnAfterXLFileCreateCallback = any(callback).(OnAfterCreateInterface[XLFile])
	
	case *XLRow:
		stage.OnAfterXLRowCreateCallback = any(callback).(OnAfterCreateInterface[XLRow])
	
	case *XLSheet:
		stage.OnAfterXLSheetCreateCallback = any(callback).(OnAfterCreateInterface[XLSheet])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *DisplaySelection:
		stage.OnAfterDisplaySelectionDeleteCallback = any(callback).(OnAfterDeleteInterface[DisplaySelection])
	
	case *XLCell:
		stage.OnAfterXLCellDeleteCallback = any(callback).(OnAfterDeleteInterface[XLCell])
	
	case *XLFile:
		stage.OnAfterXLFileDeleteCallback = any(callback).(OnAfterDeleteInterface[XLFile])
	
	case *XLRow:
		stage.OnAfterXLRowDeleteCallback = any(callback).(OnAfterDeleteInterface[XLRow])
	
	case *XLSheet:
		stage.OnAfterXLSheetDeleteCallback = any(callback).(OnAfterDeleteInterface[XLSheet])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *DisplaySelection:
		stage.OnAfterDisplaySelectionReadCallback = any(callback).(OnAfterReadInterface[DisplaySelection])
	
	case *XLCell:
		stage.OnAfterXLCellReadCallback = any(callback).(OnAfterReadInterface[XLCell])
	
	case *XLFile:
		stage.OnAfterXLFileReadCallback = any(callback).(OnAfterReadInterface[XLFile])
	
	case *XLRow:
		stage.OnAfterXLRowReadCallback = any(callback).(OnAfterReadInterface[XLRow])
	
	case *XLSheet:
		stage.OnAfterXLSheetReadCallback = any(callback).(OnAfterReadInterface[XLSheet])
	
	}
}
