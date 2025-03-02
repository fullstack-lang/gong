// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
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
	case *CheckBox:
		if stage.OnAfterCheckBoxCreateCallback != nil {
			stage.OnAfterCheckBoxCreateCallback.OnAfterCreate(stage, target)
		}
	case *DisplayedColumn:
		if stage.OnAfterDisplayedColumnCreateCallback != nil {
			stage.OnAfterDisplayedColumnCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormDiv:
		if stage.OnAfterFormDivCreateCallback != nil {
			stage.OnAfterFormDivCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormEditAssocButton:
		if stage.OnAfterFormEditAssocButtonCreateCallback != nil {
			stage.OnAfterFormEditAssocButtonCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormField:
		if stage.OnAfterFormFieldCreateCallback != nil {
			stage.OnAfterFormFieldCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormFieldDate:
		if stage.OnAfterFormFieldDateCreateCallback != nil {
			stage.OnAfterFormFieldDateCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormFieldDateTime:
		if stage.OnAfterFormFieldDateTimeCreateCallback != nil {
			stage.OnAfterFormFieldDateTimeCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormFieldFloat64:
		if stage.OnAfterFormFieldFloat64CreateCallback != nil {
			stage.OnAfterFormFieldFloat64CreateCallback.OnAfterCreate(stage, target)
		}
	case *FormFieldInt:
		if stage.OnAfterFormFieldIntCreateCallback != nil {
			stage.OnAfterFormFieldIntCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormFieldSelect:
		if stage.OnAfterFormFieldSelectCreateCallback != nil {
			stage.OnAfterFormFieldSelectCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormFieldString:
		if stage.OnAfterFormFieldStringCreateCallback != nil {
			stage.OnAfterFormFieldStringCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormFieldTime:
		if stage.OnAfterFormFieldTimeCreateCallback != nil {
			stage.OnAfterFormFieldTimeCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormGroup:
		if stage.OnAfterFormGroupCreateCallback != nil {
			stage.OnAfterFormGroupCreateCallback.OnAfterCreate(stage, target)
		}
	case *FormSortAssocButton:
		if stage.OnAfterFormSortAssocButtonCreateCallback != nil {
			stage.OnAfterFormSortAssocButtonCreateCallback.OnAfterCreate(stage, target)
		}
	case *Option:
		if stage.OnAfterOptionCreateCallback != nil {
			stage.OnAfterOptionCreateCallback.OnAfterCreate(stage, target)
		}
	case *Row:
		if stage.OnAfterRowCreateCallback != nil {
			stage.OnAfterRowCreateCallback.OnAfterCreate(stage, target)
		}
	case *Table:
		if stage.OnAfterTableCreateCallback != nil {
			stage.OnAfterTableCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

// AfterUpdateFromFront is called after a update from front
func AfterUpdateFromFront[Type Gongstruct](stage *StageStruct, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
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
	case *CheckBox:
		newTarget := any(new).(*CheckBox)
		if stage.OnAfterCheckBoxUpdateCallback != nil {
			stage.OnAfterCheckBoxUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DisplayedColumn:
		newTarget := any(new).(*DisplayedColumn)
		if stage.OnAfterDisplayedColumnUpdateCallback != nil {
			stage.OnAfterDisplayedColumnUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormDiv:
		newTarget := any(new).(*FormDiv)
		if stage.OnAfterFormDivUpdateCallback != nil {
			stage.OnAfterFormDivUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormEditAssocButton:
		newTarget := any(new).(*FormEditAssocButton)
		if stage.OnAfterFormEditAssocButtonUpdateCallback != nil {
			stage.OnAfterFormEditAssocButtonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormField:
		newTarget := any(new).(*FormField)
		if stage.OnAfterFormFieldUpdateCallback != nil {
			stage.OnAfterFormFieldUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormFieldDate:
		newTarget := any(new).(*FormFieldDate)
		if stage.OnAfterFormFieldDateUpdateCallback != nil {
			stage.OnAfterFormFieldDateUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormFieldDateTime:
		newTarget := any(new).(*FormFieldDateTime)
		if stage.OnAfterFormFieldDateTimeUpdateCallback != nil {
			stage.OnAfterFormFieldDateTimeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormFieldFloat64:
		newTarget := any(new).(*FormFieldFloat64)
		if stage.OnAfterFormFieldFloat64UpdateCallback != nil {
			stage.OnAfterFormFieldFloat64UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormFieldInt:
		newTarget := any(new).(*FormFieldInt)
		if stage.OnAfterFormFieldIntUpdateCallback != nil {
			stage.OnAfterFormFieldIntUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormFieldSelect:
		newTarget := any(new).(*FormFieldSelect)
		if stage.OnAfterFormFieldSelectUpdateCallback != nil {
			stage.OnAfterFormFieldSelectUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormFieldString:
		newTarget := any(new).(*FormFieldString)
		if stage.OnAfterFormFieldStringUpdateCallback != nil {
			stage.OnAfterFormFieldStringUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormFieldTime:
		newTarget := any(new).(*FormFieldTime)
		if stage.OnAfterFormFieldTimeUpdateCallback != nil {
			stage.OnAfterFormFieldTimeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormGroup:
		newTarget := any(new).(*FormGroup)
		if stage.OnAfterFormGroupUpdateCallback != nil {
			stage.OnAfterFormGroupUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *FormSortAssocButton:
		newTarget := any(new).(*FormSortAssocButton)
		if stage.OnAfterFormSortAssocButtonUpdateCallback != nil {
			stage.OnAfterFormSortAssocButtonUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Option:
		newTarget := any(new).(*Option)
		if stage.OnAfterOptionUpdateCallback != nil {
			stage.OnAfterOptionUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Row:
		newTarget := any(new).(*Row)
		if stage.OnAfterRowUpdateCallback != nil {
			stage.OnAfterRowUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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
func AfterDeleteFromFront[Type Gongstruct](stage *StageStruct, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
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
	case *CheckBox:
		if stage.OnAfterCheckBoxDeleteCallback != nil {
			staged := any(staged).(*CheckBox)
			stage.OnAfterCheckBoxDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DisplayedColumn:
		if stage.OnAfterDisplayedColumnDeleteCallback != nil {
			staged := any(staged).(*DisplayedColumn)
			stage.OnAfterDisplayedColumnDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormDiv:
		if stage.OnAfterFormDivDeleteCallback != nil {
			staged := any(staged).(*FormDiv)
			stage.OnAfterFormDivDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormEditAssocButton:
		if stage.OnAfterFormEditAssocButtonDeleteCallback != nil {
			staged := any(staged).(*FormEditAssocButton)
			stage.OnAfterFormEditAssocButtonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormField:
		if stage.OnAfterFormFieldDeleteCallback != nil {
			staged := any(staged).(*FormField)
			stage.OnAfterFormFieldDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormFieldDate:
		if stage.OnAfterFormFieldDateDeleteCallback != nil {
			staged := any(staged).(*FormFieldDate)
			stage.OnAfterFormFieldDateDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormFieldDateTime:
		if stage.OnAfterFormFieldDateTimeDeleteCallback != nil {
			staged := any(staged).(*FormFieldDateTime)
			stage.OnAfterFormFieldDateTimeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormFieldFloat64:
		if stage.OnAfterFormFieldFloat64DeleteCallback != nil {
			staged := any(staged).(*FormFieldFloat64)
			stage.OnAfterFormFieldFloat64DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormFieldInt:
		if stage.OnAfterFormFieldIntDeleteCallback != nil {
			staged := any(staged).(*FormFieldInt)
			stage.OnAfterFormFieldIntDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormFieldSelect:
		if stage.OnAfterFormFieldSelectDeleteCallback != nil {
			staged := any(staged).(*FormFieldSelect)
			stage.OnAfterFormFieldSelectDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormFieldString:
		if stage.OnAfterFormFieldStringDeleteCallback != nil {
			staged := any(staged).(*FormFieldString)
			stage.OnAfterFormFieldStringDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormFieldTime:
		if stage.OnAfterFormFieldTimeDeleteCallback != nil {
			staged := any(staged).(*FormFieldTime)
			stage.OnAfterFormFieldTimeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormGroup:
		if stage.OnAfterFormGroupDeleteCallback != nil {
			staged := any(staged).(*FormGroup)
			stage.OnAfterFormGroupDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *FormSortAssocButton:
		if stage.OnAfterFormSortAssocButtonDeleteCallback != nil {
			staged := any(staged).(*FormSortAssocButton)
			stage.OnAfterFormSortAssocButtonDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Option:
		if stage.OnAfterOptionDeleteCallback != nil {
			staged := any(staged).(*Option)
			stage.OnAfterOptionDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Row:
		if stage.OnAfterRowDeleteCallback != nil {
			staged := any(staged).(*Row)
			stage.OnAfterRowDeleteCallback.OnAfterDelete(stage, staged, front)
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
func AfterReadFromFront[Type Gongstruct](stage *StageStruct, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
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
	case *CheckBox:
		if stage.OnAfterCheckBoxReadCallback != nil {
			stage.OnAfterCheckBoxReadCallback.OnAfterRead(stage, target)
		}
	case *DisplayedColumn:
		if stage.OnAfterDisplayedColumnReadCallback != nil {
			stage.OnAfterDisplayedColumnReadCallback.OnAfterRead(stage, target)
		}
	case *FormDiv:
		if stage.OnAfterFormDivReadCallback != nil {
			stage.OnAfterFormDivReadCallback.OnAfterRead(stage, target)
		}
	case *FormEditAssocButton:
		if stage.OnAfterFormEditAssocButtonReadCallback != nil {
			stage.OnAfterFormEditAssocButtonReadCallback.OnAfterRead(stage, target)
		}
	case *FormField:
		if stage.OnAfterFormFieldReadCallback != nil {
			stage.OnAfterFormFieldReadCallback.OnAfterRead(stage, target)
		}
	case *FormFieldDate:
		if stage.OnAfterFormFieldDateReadCallback != nil {
			stage.OnAfterFormFieldDateReadCallback.OnAfterRead(stage, target)
		}
	case *FormFieldDateTime:
		if stage.OnAfterFormFieldDateTimeReadCallback != nil {
			stage.OnAfterFormFieldDateTimeReadCallback.OnAfterRead(stage, target)
		}
	case *FormFieldFloat64:
		if stage.OnAfterFormFieldFloat64ReadCallback != nil {
			stage.OnAfterFormFieldFloat64ReadCallback.OnAfterRead(stage, target)
		}
	case *FormFieldInt:
		if stage.OnAfterFormFieldIntReadCallback != nil {
			stage.OnAfterFormFieldIntReadCallback.OnAfterRead(stage, target)
		}
	case *FormFieldSelect:
		if stage.OnAfterFormFieldSelectReadCallback != nil {
			stage.OnAfterFormFieldSelectReadCallback.OnAfterRead(stage, target)
		}
	case *FormFieldString:
		if stage.OnAfterFormFieldStringReadCallback != nil {
			stage.OnAfterFormFieldStringReadCallback.OnAfterRead(stage, target)
		}
	case *FormFieldTime:
		if stage.OnAfterFormFieldTimeReadCallback != nil {
			stage.OnAfterFormFieldTimeReadCallback.OnAfterRead(stage, target)
		}
	case *FormGroup:
		if stage.OnAfterFormGroupReadCallback != nil {
			stage.OnAfterFormGroupReadCallback.OnAfterRead(stage, target)
		}
	case *FormSortAssocButton:
		if stage.OnAfterFormSortAssocButtonReadCallback != nil {
			stage.OnAfterFormSortAssocButtonReadCallback.OnAfterRead(stage, target)
		}
	case *Option:
		if stage.OnAfterOptionReadCallback != nil {
			stage.OnAfterOptionReadCallback.OnAfterRead(stage, target)
		}
	case *Row:
		if stage.OnAfterRowReadCallback != nil {
			stage.OnAfterRowReadCallback.OnAfterRead(stage, target)
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
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
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
	
	case *CheckBox:
		stage.OnAfterCheckBoxUpdateCallback = any(callback).(OnAfterUpdateInterface[CheckBox])
	
	case *DisplayedColumn:
		stage.OnAfterDisplayedColumnUpdateCallback = any(callback).(OnAfterUpdateInterface[DisplayedColumn])
	
	case *FormDiv:
		stage.OnAfterFormDivUpdateCallback = any(callback).(OnAfterUpdateInterface[FormDiv])
	
	case *FormEditAssocButton:
		stage.OnAfterFormEditAssocButtonUpdateCallback = any(callback).(OnAfterUpdateInterface[FormEditAssocButton])
	
	case *FormField:
		stage.OnAfterFormFieldUpdateCallback = any(callback).(OnAfterUpdateInterface[FormField])
	
	case *FormFieldDate:
		stage.OnAfterFormFieldDateUpdateCallback = any(callback).(OnAfterUpdateInterface[FormFieldDate])
	
	case *FormFieldDateTime:
		stage.OnAfterFormFieldDateTimeUpdateCallback = any(callback).(OnAfterUpdateInterface[FormFieldDateTime])
	
	case *FormFieldFloat64:
		stage.OnAfterFormFieldFloat64UpdateCallback = any(callback).(OnAfterUpdateInterface[FormFieldFloat64])
	
	case *FormFieldInt:
		stage.OnAfterFormFieldIntUpdateCallback = any(callback).(OnAfterUpdateInterface[FormFieldInt])
	
	case *FormFieldSelect:
		stage.OnAfterFormFieldSelectUpdateCallback = any(callback).(OnAfterUpdateInterface[FormFieldSelect])
	
	case *FormFieldString:
		stage.OnAfterFormFieldStringUpdateCallback = any(callback).(OnAfterUpdateInterface[FormFieldString])
	
	case *FormFieldTime:
		stage.OnAfterFormFieldTimeUpdateCallback = any(callback).(OnAfterUpdateInterface[FormFieldTime])
	
	case *FormGroup:
		stage.OnAfterFormGroupUpdateCallback = any(callback).(OnAfterUpdateInterface[FormGroup])
	
	case *FormSortAssocButton:
		stage.OnAfterFormSortAssocButtonUpdateCallback = any(callback).(OnAfterUpdateInterface[FormSortAssocButton])
	
	case *Option:
		stage.OnAfterOptionUpdateCallback = any(callback).(OnAfterUpdateInterface[Option])
	
	case *Row:
		stage.OnAfterRowUpdateCallback = any(callback).(OnAfterUpdateInterface[Row])
	
	case *Table:
		stage.OnAfterTableUpdateCallback = any(callback).(OnAfterUpdateInterface[Table])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
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
	
	case *CheckBox:
		stage.OnAfterCheckBoxCreateCallback = any(callback).(OnAfterCreateInterface[CheckBox])
	
	case *DisplayedColumn:
		stage.OnAfterDisplayedColumnCreateCallback = any(callback).(OnAfterCreateInterface[DisplayedColumn])
	
	case *FormDiv:
		stage.OnAfterFormDivCreateCallback = any(callback).(OnAfterCreateInterface[FormDiv])
	
	case *FormEditAssocButton:
		stage.OnAfterFormEditAssocButtonCreateCallback = any(callback).(OnAfterCreateInterface[FormEditAssocButton])
	
	case *FormField:
		stage.OnAfterFormFieldCreateCallback = any(callback).(OnAfterCreateInterface[FormField])
	
	case *FormFieldDate:
		stage.OnAfterFormFieldDateCreateCallback = any(callback).(OnAfterCreateInterface[FormFieldDate])
	
	case *FormFieldDateTime:
		stage.OnAfterFormFieldDateTimeCreateCallback = any(callback).(OnAfterCreateInterface[FormFieldDateTime])
	
	case *FormFieldFloat64:
		stage.OnAfterFormFieldFloat64CreateCallback = any(callback).(OnAfterCreateInterface[FormFieldFloat64])
	
	case *FormFieldInt:
		stage.OnAfterFormFieldIntCreateCallback = any(callback).(OnAfterCreateInterface[FormFieldInt])
	
	case *FormFieldSelect:
		stage.OnAfterFormFieldSelectCreateCallback = any(callback).(OnAfterCreateInterface[FormFieldSelect])
	
	case *FormFieldString:
		stage.OnAfterFormFieldStringCreateCallback = any(callback).(OnAfterCreateInterface[FormFieldString])
	
	case *FormFieldTime:
		stage.OnAfterFormFieldTimeCreateCallback = any(callback).(OnAfterCreateInterface[FormFieldTime])
	
	case *FormGroup:
		stage.OnAfterFormGroupCreateCallback = any(callback).(OnAfterCreateInterface[FormGroup])
	
	case *FormSortAssocButton:
		stage.OnAfterFormSortAssocButtonCreateCallback = any(callback).(OnAfterCreateInterface[FormSortAssocButton])
	
	case *Option:
		stage.OnAfterOptionCreateCallback = any(callback).(OnAfterCreateInterface[Option])
	
	case *Row:
		stage.OnAfterRowCreateCallback = any(callback).(OnAfterCreateInterface[Row])
	
	case *Table:
		stage.OnAfterTableCreateCallback = any(callback).(OnAfterCreateInterface[Table])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
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
	
	case *CheckBox:
		stage.OnAfterCheckBoxDeleteCallback = any(callback).(OnAfterDeleteInterface[CheckBox])
	
	case *DisplayedColumn:
		stage.OnAfterDisplayedColumnDeleteCallback = any(callback).(OnAfterDeleteInterface[DisplayedColumn])
	
	case *FormDiv:
		stage.OnAfterFormDivDeleteCallback = any(callback).(OnAfterDeleteInterface[FormDiv])
	
	case *FormEditAssocButton:
		stage.OnAfterFormEditAssocButtonDeleteCallback = any(callback).(OnAfterDeleteInterface[FormEditAssocButton])
	
	case *FormField:
		stage.OnAfterFormFieldDeleteCallback = any(callback).(OnAfterDeleteInterface[FormField])
	
	case *FormFieldDate:
		stage.OnAfterFormFieldDateDeleteCallback = any(callback).(OnAfterDeleteInterface[FormFieldDate])
	
	case *FormFieldDateTime:
		stage.OnAfterFormFieldDateTimeDeleteCallback = any(callback).(OnAfterDeleteInterface[FormFieldDateTime])
	
	case *FormFieldFloat64:
		stage.OnAfterFormFieldFloat64DeleteCallback = any(callback).(OnAfterDeleteInterface[FormFieldFloat64])
	
	case *FormFieldInt:
		stage.OnAfterFormFieldIntDeleteCallback = any(callback).(OnAfterDeleteInterface[FormFieldInt])
	
	case *FormFieldSelect:
		stage.OnAfterFormFieldSelectDeleteCallback = any(callback).(OnAfterDeleteInterface[FormFieldSelect])
	
	case *FormFieldString:
		stage.OnAfterFormFieldStringDeleteCallback = any(callback).(OnAfterDeleteInterface[FormFieldString])
	
	case *FormFieldTime:
		stage.OnAfterFormFieldTimeDeleteCallback = any(callback).(OnAfterDeleteInterface[FormFieldTime])
	
	case *FormGroup:
		stage.OnAfterFormGroupDeleteCallback = any(callback).(OnAfterDeleteInterface[FormGroup])
	
	case *FormSortAssocButton:
		stage.OnAfterFormSortAssocButtonDeleteCallback = any(callback).(OnAfterDeleteInterface[FormSortAssocButton])
	
	case *Option:
		stage.OnAfterOptionDeleteCallback = any(callback).(OnAfterDeleteInterface[Option])
	
	case *Row:
		stage.OnAfterRowDeleteCallback = any(callback).(OnAfterDeleteInterface[Row])
	
	case *Table:
		stage.OnAfterTableDeleteCallback = any(callback).(OnAfterDeleteInterface[Table])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *StageStruct, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
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
	
	case *CheckBox:
		stage.OnAfterCheckBoxReadCallback = any(callback).(OnAfterReadInterface[CheckBox])
	
	case *DisplayedColumn:
		stage.OnAfterDisplayedColumnReadCallback = any(callback).(OnAfterReadInterface[DisplayedColumn])
	
	case *FormDiv:
		stage.OnAfterFormDivReadCallback = any(callback).(OnAfterReadInterface[FormDiv])
	
	case *FormEditAssocButton:
		stage.OnAfterFormEditAssocButtonReadCallback = any(callback).(OnAfterReadInterface[FormEditAssocButton])
	
	case *FormField:
		stage.OnAfterFormFieldReadCallback = any(callback).(OnAfterReadInterface[FormField])
	
	case *FormFieldDate:
		stage.OnAfterFormFieldDateReadCallback = any(callback).(OnAfterReadInterface[FormFieldDate])
	
	case *FormFieldDateTime:
		stage.OnAfterFormFieldDateTimeReadCallback = any(callback).(OnAfterReadInterface[FormFieldDateTime])
	
	case *FormFieldFloat64:
		stage.OnAfterFormFieldFloat64ReadCallback = any(callback).(OnAfterReadInterface[FormFieldFloat64])
	
	case *FormFieldInt:
		stage.OnAfterFormFieldIntReadCallback = any(callback).(OnAfterReadInterface[FormFieldInt])
	
	case *FormFieldSelect:
		stage.OnAfterFormFieldSelectReadCallback = any(callback).(OnAfterReadInterface[FormFieldSelect])
	
	case *FormFieldString:
		stage.OnAfterFormFieldStringReadCallback = any(callback).(OnAfterReadInterface[FormFieldString])
	
	case *FormFieldTime:
		stage.OnAfterFormFieldTimeReadCallback = any(callback).(OnAfterReadInterface[FormFieldTime])
	
	case *FormGroup:
		stage.OnAfterFormGroupReadCallback = any(callback).(OnAfterReadInterface[FormGroup])
	
	case *FormSortAssocButton:
		stage.OnAfterFormSortAssocButtonReadCallback = any(callback).(OnAfterReadInterface[FormSortAssocButton])
	
	case *Option:
		stage.OnAfterOptionReadCallback = any(callback).(OnAfterReadInterface[Option])
	
	case *Row:
		stage.OnAfterRowReadCallback = any(callback).(OnAfterReadInterface[Row])
	
	case *Table:
		stage.OnAfterTableReadCallback = any(callback).(OnAfterReadInterface[Table])
	
	}
}
