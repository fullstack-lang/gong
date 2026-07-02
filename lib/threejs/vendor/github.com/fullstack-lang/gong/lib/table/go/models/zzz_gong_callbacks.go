// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

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

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

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

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

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

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *Button:
		if stage.OnAfterButtonReadCallback != nil {
			stage.OnAfterButtonReadCallback.OnAfterRead(stage, target)
		}
	case *Cell:
		if stage.OnAfterCellReadCallback != nil {
			stage.OnAfterCellReadCallback.OnAfterRead(stage, target)
		}
	case *CellBoolean:
		if stage.OnAfterCellBooleanReadCallback != nil {
			stage.OnAfterCellBooleanReadCallback.OnAfterRead(stage, target)
		}
	case *CellFloat64:
		if stage.OnAfterCellFloat64ReadCallback != nil {
			stage.OnAfterCellFloat64ReadCallback.OnAfterRead(stage, target)
		}
	case *CellIcon:
		if stage.OnAfterCellIconReadCallback != nil {
			stage.OnAfterCellIconReadCallback.OnAfterRead(stage, target)
		}
	case *CellInt:
		if stage.OnAfterCellIntReadCallback != nil {
			stage.OnAfterCellIntReadCallback.OnAfterRead(stage, target)
		}
	case *CellString:
		if stage.OnAfterCellStringReadCallback != nil {
			stage.OnAfterCellStringReadCallback.OnAfterRead(stage, target)
		}
	case *DisplayedColumn:
		if stage.OnAfterDisplayedColumnReadCallback != nil {
			stage.OnAfterDisplayedColumnReadCallback.OnAfterRead(stage, target)
		}
	case *Row:
		if stage.OnAfterRowReadCallback != nil {
			stage.OnAfterRowReadCallback.OnAfterRead(stage, target)
		}
	case *SVGIcon:
		if stage.OnAfterSVGIconReadCallback != nil {
			stage.OnAfterSVGIconReadCallback.OnAfterRead(stage, target)
		}
	case *Table:
		if stage.OnAfterTableReadCallback != nil {
			stage.OnAfterTableReadCallback.OnAfterRead(stage, target)
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
	case *Button:
		stage.OnAfterButtonUpdateCallback = any(callback).(OnAfterUpdateInterface[Button])
	case *Cell:
		stage.OnAfterCellUpdateCallback = any(callback).(OnAfterUpdateInterface[Cell])
	case *CellBoolean:
		stage.OnAfterCellBooleanUpdateCallback = any(callback).(OnAfterUpdateInterface[CellBoolean])
	case *CellFloat64:
		stage.OnAfterCellFloat64UpdateCallback = any(callback).(OnAfterUpdateInterface[CellFloat64])
	case *CellIcon:
		stage.OnAfterCellIconUpdateCallback = any(callback).(OnAfterUpdateInterface[CellIcon])
	case *CellInt:
		stage.OnAfterCellIntUpdateCallback = any(callback).(OnAfterUpdateInterface[CellInt])
	case *CellString:
		stage.OnAfterCellStringUpdateCallback = any(callback).(OnAfterUpdateInterface[CellString])
	case *DisplayedColumn:
		stage.OnAfterDisplayedColumnUpdateCallback = any(callback).(OnAfterUpdateInterface[DisplayedColumn])
	case *Row:
		stage.OnAfterRowUpdateCallback = any(callback).(OnAfterUpdateInterface[Row])
	case *SVGIcon:
		stage.OnAfterSVGIconUpdateCallback = any(callback).(OnAfterUpdateInterface[SVGIcon])
	case *Table:
		stage.OnAfterTableUpdateCallback = any(callback).(OnAfterUpdateInterface[Table])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *Button:
		stage.OnAfterButtonCreateCallback = any(callback).(OnAfterCreateInterface[Button])
	case *Cell:
		stage.OnAfterCellCreateCallback = any(callback).(OnAfterCreateInterface[Cell])
	case *CellBoolean:
		stage.OnAfterCellBooleanCreateCallback = any(callback).(OnAfterCreateInterface[CellBoolean])
	case *CellFloat64:
		stage.OnAfterCellFloat64CreateCallback = any(callback).(OnAfterCreateInterface[CellFloat64])
	case *CellIcon:
		stage.OnAfterCellIconCreateCallback = any(callback).(OnAfterCreateInterface[CellIcon])
	case *CellInt:
		stage.OnAfterCellIntCreateCallback = any(callback).(OnAfterCreateInterface[CellInt])
	case *CellString:
		stage.OnAfterCellStringCreateCallback = any(callback).(OnAfterCreateInterface[CellString])
	case *DisplayedColumn:
		stage.OnAfterDisplayedColumnCreateCallback = any(callback).(OnAfterCreateInterface[DisplayedColumn])
	case *Row:
		stage.OnAfterRowCreateCallback = any(callback).(OnAfterCreateInterface[Row])
	case *SVGIcon:
		stage.OnAfterSVGIconCreateCallback = any(callback).(OnAfterCreateInterface[SVGIcon])
	case *Table:
		stage.OnAfterTableCreateCallback = any(callback).(OnAfterCreateInterface[Table])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *Button:
		stage.OnAfterButtonDeleteCallback = any(callback).(OnAfterDeleteInterface[Button])
	case *Cell:
		stage.OnAfterCellDeleteCallback = any(callback).(OnAfterDeleteInterface[Cell])
	case *CellBoolean:
		stage.OnAfterCellBooleanDeleteCallback = any(callback).(OnAfterDeleteInterface[CellBoolean])
	case *CellFloat64:
		stage.OnAfterCellFloat64DeleteCallback = any(callback).(OnAfterDeleteInterface[CellFloat64])
	case *CellIcon:
		stage.OnAfterCellIconDeleteCallback = any(callback).(OnAfterDeleteInterface[CellIcon])
	case *CellInt:
		stage.OnAfterCellIntDeleteCallback = any(callback).(OnAfterDeleteInterface[CellInt])
	case *CellString:
		stage.OnAfterCellStringDeleteCallback = any(callback).(OnAfterDeleteInterface[CellString])
	case *DisplayedColumn:
		stage.OnAfterDisplayedColumnDeleteCallback = any(callback).(OnAfterDeleteInterface[DisplayedColumn])
	case *Row:
		stage.OnAfterRowDeleteCallback = any(callback).(OnAfterDeleteInterface[Row])
	case *SVGIcon:
		stage.OnAfterSVGIconDeleteCallback = any(callback).(OnAfterDeleteInterface[SVGIcon])
	case *Table:
		stage.OnAfterTableDeleteCallback = any(callback).(OnAfterDeleteInterface[Table])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *Button:
		stage.OnAfterButtonReadCallback = any(callback).(OnAfterReadInterface[Button])
	case *Cell:
		stage.OnAfterCellReadCallback = any(callback).(OnAfterReadInterface[Cell])
	case *CellBoolean:
		stage.OnAfterCellBooleanReadCallback = any(callback).(OnAfterReadInterface[CellBoolean])
	case *CellFloat64:
		stage.OnAfterCellFloat64ReadCallback = any(callback).(OnAfterReadInterface[CellFloat64])
	case *CellIcon:
		stage.OnAfterCellIconReadCallback = any(callback).(OnAfterReadInterface[CellIcon])
	case *CellInt:
		stage.OnAfterCellIntReadCallback = any(callback).(OnAfterReadInterface[CellInt])
	case *CellString:
		stage.OnAfterCellStringReadCallback = any(callback).(OnAfterReadInterface[CellString])
	case *DisplayedColumn:
		stage.OnAfterDisplayedColumnReadCallback = any(callback).(OnAfterReadInterface[DisplayedColumn])
	case *Row:
		stage.OnAfterRowReadCallback = any(callback).(OnAfterReadInterface[Row])
	case *SVGIcon:
		stage.OnAfterSVGIconReadCallback = any(callback).(OnAfterReadInterface[SVGIcon])
	case *Table:
		stage.OnAfterTableReadCallback = any(callback).(OnAfterReadInterface[Table])
	}
}
