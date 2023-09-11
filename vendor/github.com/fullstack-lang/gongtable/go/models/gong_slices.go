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
func EvictInOtherSlices[T PointerToGongstruct, TF PointerToGongstruct](
	stage *StageStruct,
	owningInstance T,
	sliceField []TF,
	fieldName string) {

	// create a map of the field elements
	setOfFieldInstances := make(map[TF]any, 0)
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
		// tweaking, it might be streamlined
		if fieldName == "FormFields" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*FormDiv)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.FormFields).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.FormFields = make([]*FormField, 0)
				if any(_instance).(*FormDiv) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "CheckBoxs" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*FormDiv)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.CheckBoxs).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.CheckBoxs = make([]*CheckBox, 0)
				if any(_instance).(*FormDiv) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
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
		// tweaking, it might be streamlined
		if fieldName == "Options" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*FormFieldSelect)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Options).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Options = make([]*Option, 0)
				if any(_instance).(*FormFieldSelect) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
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
		// tweaking, it might be streamlined
		if fieldName == "FormDivs" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*FormGroup)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.FormDivs).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.FormDivs = make([]*FormDiv, 0)
				if any(_instance).(*FormGroup) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
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
		// tweaking, it might be streamlined
		if fieldName == "Cells" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Row)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Cells).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Cells = make([]*Cell, 0)
				if any(_instance).(*Row) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *Table:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "DisplayedColumns" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Table)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.DisplayedColumns).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.DisplayedColumns = make([]*DisplayedColumn, 0)
				if any(_instance).(*Table) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Rows" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Table)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Rows).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Rows = make([]*Row, 0)
				if any(_instance).(*Table) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}
