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
	case *GongBasicField:
		// insertion point per field

	case *GongEnum:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "GongEnumValues" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*GongEnum)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.GongEnumValues).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.GongEnumValues = make([]*GongEnumValue, 0)
				if any(_instance).(*GongEnum) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *GongEnumValue:
		// insertion point per field

	case *GongLink:
		// insertion point per field

	case *GongNote:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Links" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*GongNote)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Links).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Links = make([]*GongLink, 0)
				if any(_instance).(*GongNote) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *GongStruct:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "GongBasicFields" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*GongStruct)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.GongBasicFields).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.GongBasicFields = make([]*GongBasicField, 0)
				if any(_instance).(*GongStruct) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "GongTimeFields" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*GongStruct)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.GongTimeFields).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.GongTimeFields = make([]*GongTimeField, 0)
				if any(_instance).(*GongStruct) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "PointerToGongStructFields" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*GongStruct)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.PointerToGongStructFields).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.PointerToGongStructFields = make([]*PointerToGongStructField, 0)
				if any(_instance).(*GongStruct) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "SliceOfPointerToGongStructFields" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*GongStruct)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.SliceOfPointerToGongStructFields).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.SliceOfPointerToGongStructFields = make([]*SliceOfPointerToGongStructField, 0)
				if any(_instance).(*GongStruct) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *GongTimeField:
		// insertion point per field

	case *Meta:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "MetaReferences" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Meta)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.MetaReferences).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.MetaReferences = make([]*MetaReference, 0)
				if any(_instance).(*Meta) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *MetaReference:
		// insertion point per field

	case *ModelPkg:
		// insertion point per field

	case *PointerToGongStructField:
		// insertion point per field

	case *SliceOfPointerToGongStructField:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}
