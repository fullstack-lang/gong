// generated code - do not edit
package models

// EvictInOtherSlices allows for adherance between
// the gong association model and go.
//
// Says you have a Astruct struct with a slice field "anarrayofb []*Bstruct"
//
// go allows many Astruct instance to have the anarrayofb field that have the
// same pointers. go slices are MANY-MANY association.
//
// With gong it is only ZERO-ONE-MANY associations, a Bstruct can be pointed only
// once by an Astruct instance through a given field. This follows the requirement
// that gong is suited for full stack programming and therefore the association
// is encoded as a reverse pointer (not as a joint table). In gong, a named struct
// is translated in a table and each table is a named struct.
//
// EvictInOtherSlices removes the fields instances from other
// fields of other instance
//
// Note : algo is in O(N)log(N) of nb of Astruct and Bstruct instances
func EvictInOtherSlices[OwningType PointerToGongstruct, FieldType PointerToGongstruct](
	stage *StageStruct,
	owningInstance OwningType,
	sliceField []FieldType,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[FieldType]any, 0)
	for _, fieldInstance := range sliceField {
		setOfFieldInstances[fieldInstance] = true
	}

	switch owningInstanceInfered := any(owningInstance).(type) {
	// insertion point
	case *Cell:
		// insertion point per field

	case *CellBoolean:
		// insertion point per field

	case *CellFloat64:
		// insertion point per field

	case *CellIcon:
		// insertion point per field

	case *CellInt:
		// insertion point per field

	case *CellString:
		// insertion point per field

	case *CheckBox:
		// insertion point per field

	case *DisplayedColumn:
		// insertion point per field

	case *FormDiv:
		// insertion point per field
		if fieldName == "FormFields" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*FormDiv) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*FormDiv)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.FormFields).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.FormFields = _inferedTypeInstance.FormFields[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.FormFields =
								append(_inferedTypeInstance.FormFields, any(fieldInstance).(*FormField))
						}
					}
				}
			}
		}
		if fieldName == "CheckBoxs" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*FormDiv) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*FormDiv)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.CheckBoxs).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.CheckBoxs = _inferedTypeInstance.CheckBoxs[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.CheckBoxs =
								append(_inferedTypeInstance.CheckBoxs, any(fieldInstance).(*CheckBox))
						}
					}
				}
			}
		}

	case *FormEditAssocButton:
		// insertion point per field

	case *FormField:
		// insertion point per field

	case *FormFieldDate:
		// insertion point per field

	case *FormFieldDateTime:
		// insertion point per field

	case *FormFieldFloat64:
		// insertion point per field

	case *FormFieldInt:
		// insertion point per field

	case *FormFieldSelect:
		// insertion point per field
		if fieldName == "Options" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*FormFieldSelect) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*FormFieldSelect)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Options).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Options = _inferedTypeInstance.Options[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Options =
								append(_inferedTypeInstance.Options, any(fieldInstance).(*Option))
						}
					}
				}
			}
		}

	case *FormFieldString:
		// insertion point per field

	case *FormFieldTime:
		// insertion point per field

	case *FormGroup:
		// insertion point per field
		if fieldName == "FormDivs" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*FormGroup) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*FormGroup)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.FormDivs).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.FormDivs = _inferedTypeInstance.FormDivs[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.FormDivs =
								append(_inferedTypeInstance.FormDivs, any(fieldInstance).(*FormDiv))
						}
					}
				}
			}
		}

	case *FormSortAssocButton:
		// insertion point per field

	case *Option:
		// insertion point per field

	case *Row:
		// insertion point per field
		if fieldName == "Cells" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Row) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Row)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Cells).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Cells = _inferedTypeInstance.Cells[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Cells =
								append(_inferedTypeInstance.Cells, any(fieldInstance).(*Cell))
						}
					}
				}
			}
		}

	case *Table:
		// insertion point per field
		if fieldName == "DisplayedColumns" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Table) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Table)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.DisplayedColumns).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.DisplayedColumns = _inferedTypeInstance.DisplayedColumns[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.DisplayedColumns =
								append(_inferedTypeInstance.DisplayedColumns, any(fieldInstance).(*DisplayedColumn))
						}
					}
				}
			}
		}
		if fieldName == "Rows" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Table) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Table)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Rows).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Rows = _inferedTypeInstance.Rows[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Rows =
								append(_inferedTypeInstance.Rows, any(fieldInstance).(*Row))
						}
					}
				}
			}
		}

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}
