// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AsSplit:
		if stage.OnAfterAsSplitCreateCallback != nil {
			stage.OnAfterAsSplitCreateCallback.OnAfterCreate(stage, target)
		}
	case *AsSplitArea:
		if stage.OnAfterAsSplitAreaCreateCallback != nil {
			stage.OnAfterAsSplitAreaCreateCallback.OnAfterCreate(stage, target)
		}
	case *Doc:
		if stage.OnAfterDocCreateCallback != nil {
			stage.OnAfterDocCreateCallback.OnAfterCreate(stage, target)
		}
	case *Form:
		if stage.OnAfterFormCreateCallback != nil {
			stage.OnAfterFormCreateCallback.OnAfterCreate(stage, target)
		}
	case *Svg:
		if stage.OnAfterSvgCreateCallback != nil {
			stage.OnAfterSvgCreateCallback.OnAfterCreate(stage, target)
		}
	case *Table:
		if stage.OnAfterTableCreateCallback != nil {
			stage.OnAfterTableCreateCallback.OnAfterCreate(stage, target)
		}
	case *Tree:
		if stage.OnAfterTreeCreateCallback != nil {
			stage.OnAfterTreeCreateCallback.OnAfterCreate(stage, target)
		}
	case *View:
		if stage.OnAfterViewCreateCallback != nil {
			stage.OnAfterViewCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *AsSplit:
		newTarget := any(new).(*AsSplit)
		if stage.OnAfterAsSplitUpdateCallback != nil {
			stage.OnAfterAsSplitUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *AsSplitArea:
		newTarget := any(new).(*AsSplitArea)
		if stage.OnAfterAsSplitAreaUpdateCallback != nil {
			stage.OnAfterAsSplitAreaUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Doc:
		newTarget := any(new).(*Doc)
		if stage.OnAfterDocUpdateCallback != nil {
			stage.OnAfterDocUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Form:
		newTarget := any(new).(*Form)
		if stage.OnAfterFormUpdateCallback != nil {
			stage.OnAfterFormUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Svg:
		newTarget := any(new).(*Svg)
		if stage.OnAfterSvgUpdateCallback != nil {
			stage.OnAfterSvgUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Table:
		newTarget := any(new).(*Table)
		if stage.OnAfterTableUpdateCallback != nil {
			stage.OnAfterTableUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Tree:
		newTarget := any(new).(*Tree)
		if stage.OnAfterTreeUpdateCallback != nil {
			stage.OnAfterTreeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *View:
		newTarget := any(new).(*View)
		if stage.OnAfterViewUpdateCallback != nil {
			stage.OnAfterViewUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *AsSplit:
		if stage.OnAfterAsSplitDeleteCallback != nil {
			staged := any(staged).(*AsSplit)
			stage.OnAfterAsSplitDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *AsSplitArea:
		if stage.OnAfterAsSplitAreaDeleteCallback != nil {
			staged := any(staged).(*AsSplitArea)
			stage.OnAfterAsSplitAreaDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Doc:
		if stage.OnAfterDocDeleteCallback != nil {
			staged := any(staged).(*Doc)
			stage.OnAfterDocDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Form:
		if stage.OnAfterFormDeleteCallback != nil {
			staged := any(staged).(*Form)
			stage.OnAfterFormDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Svg:
		if stage.OnAfterSvgDeleteCallback != nil {
			staged := any(staged).(*Svg)
			stage.OnAfterSvgDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Table:
		if stage.OnAfterTableDeleteCallback != nil {
			staged := any(staged).(*Table)
			stage.OnAfterTableDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Tree:
		if stage.OnAfterTreeDeleteCallback != nil {
			staged := any(staged).(*Tree)
			stage.OnAfterTreeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *View:
		if stage.OnAfterViewDeleteCallback != nil {
			staged := any(staged).(*View)
			stage.OnAfterViewDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *AsSplit:
		if stage.OnAfterAsSplitReadCallback != nil {
			stage.OnAfterAsSplitReadCallback.OnAfterRead(stage, target)
		}
	case *AsSplitArea:
		if stage.OnAfterAsSplitAreaReadCallback != nil {
			stage.OnAfterAsSplitAreaReadCallback.OnAfterRead(stage, target)
		}
	case *Doc:
		if stage.OnAfterDocReadCallback != nil {
			stage.OnAfterDocReadCallback.OnAfterRead(stage, target)
		}
	case *Form:
		if stage.OnAfterFormReadCallback != nil {
			stage.OnAfterFormReadCallback.OnAfterRead(stage, target)
		}
	case *Svg:
		if stage.OnAfterSvgReadCallback != nil {
			stage.OnAfterSvgReadCallback.OnAfterRead(stage, target)
		}
	case *Table:
		if stage.OnAfterTableReadCallback != nil {
			stage.OnAfterTableReadCallback.OnAfterRead(stage, target)
		}
	case *Tree:
		if stage.OnAfterTreeReadCallback != nil {
			stage.OnAfterTreeReadCallback.OnAfterRead(stage, target)
		}
	case *View:
		if stage.OnAfterViewReadCallback != nil {
			stage.OnAfterViewReadCallback.OnAfterRead(stage, target)
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
	case *AsSplit:
		stage.OnAfterAsSplitUpdateCallback = any(callback).(OnAfterUpdateInterface[AsSplit])
	
	case *AsSplitArea:
		stage.OnAfterAsSplitAreaUpdateCallback = any(callback).(OnAfterUpdateInterface[AsSplitArea])
	
	case *Doc:
		stage.OnAfterDocUpdateCallback = any(callback).(OnAfterUpdateInterface[Doc])
	
	case *Form:
		stage.OnAfterFormUpdateCallback = any(callback).(OnAfterUpdateInterface[Form])
	
	case *Svg:
		stage.OnAfterSvgUpdateCallback = any(callback).(OnAfterUpdateInterface[Svg])
	
	case *Table:
		stage.OnAfterTableUpdateCallback = any(callback).(OnAfterUpdateInterface[Table])
	
	case *Tree:
		stage.OnAfterTreeUpdateCallback = any(callback).(OnAfterUpdateInterface[Tree])
	
	case *View:
		stage.OnAfterViewUpdateCallback = any(callback).(OnAfterUpdateInterface[View])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AsSplit:
		stage.OnAfterAsSplitCreateCallback = any(callback).(OnAfterCreateInterface[AsSplit])
	
	case *AsSplitArea:
		stage.OnAfterAsSplitAreaCreateCallback = any(callback).(OnAfterCreateInterface[AsSplitArea])
	
	case *Doc:
		stage.OnAfterDocCreateCallback = any(callback).(OnAfterCreateInterface[Doc])
	
	case *Form:
		stage.OnAfterFormCreateCallback = any(callback).(OnAfterCreateInterface[Form])
	
	case *Svg:
		stage.OnAfterSvgCreateCallback = any(callback).(OnAfterCreateInterface[Svg])
	
	case *Table:
		stage.OnAfterTableCreateCallback = any(callback).(OnAfterCreateInterface[Table])
	
	case *Tree:
		stage.OnAfterTreeCreateCallback = any(callback).(OnAfterCreateInterface[Tree])
	
	case *View:
		stage.OnAfterViewCreateCallback = any(callback).(OnAfterCreateInterface[View])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AsSplit:
		stage.OnAfterAsSplitDeleteCallback = any(callback).(OnAfterDeleteInterface[AsSplit])
	
	case *AsSplitArea:
		stage.OnAfterAsSplitAreaDeleteCallback = any(callback).(OnAfterDeleteInterface[AsSplitArea])
	
	case *Doc:
		stage.OnAfterDocDeleteCallback = any(callback).(OnAfterDeleteInterface[Doc])
	
	case *Form:
		stage.OnAfterFormDeleteCallback = any(callback).(OnAfterDeleteInterface[Form])
	
	case *Svg:
		stage.OnAfterSvgDeleteCallback = any(callback).(OnAfterDeleteInterface[Svg])
	
	case *Table:
		stage.OnAfterTableDeleteCallback = any(callback).(OnAfterDeleteInterface[Table])
	
	case *Tree:
		stage.OnAfterTreeDeleteCallback = any(callback).(OnAfterDeleteInterface[Tree])
	
	case *View:
		stage.OnAfterViewDeleteCallback = any(callback).(OnAfterDeleteInterface[View])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *AsSplit:
		stage.OnAfterAsSplitReadCallback = any(callback).(OnAfterReadInterface[AsSplit])
	
	case *AsSplitArea:
		stage.OnAfterAsSplitAreaReadCallback = any(callback).(OnAfterReadInterface[AsSplitArea])
	
	case *Doc:
		stage.OnAfterDocReadCallback = any(callback).(OnAfterReadInterface[Doc])
	
	case *Form:
		stage.OnAfterFormReadCallback = any(callback).(OnAfterReadInterface[Form])
	
	case *Svg:
		stage.OnAfterSvgReadCallback = any(callback).(OnAfterReadInterface[Svg])
	
	case *Table:
		stage.OnAfterTableReadCallback = any(callback).(OnAfterReadInterface[Table])
	
	case *Tree:
		stage.OnAfterTreeReadCallback = any(callback).(OnAfterReadInterface[Tree])
	
	case *View:
		stage.OnAfterViewReadCallback = any(callback).(OnAfterReadInterface[View])
	
	}
}
