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
	case *Astruct:
		// insertion point per field
		if fieldName == "Anarrayofb" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Astruct) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Astruct)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Anarrayofb).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Anarrayofb = _inferedTypeInstance.Anarrayofb[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Anarrayofb =
								append(_inferedTypeInstance.Anarrayofb, any(fieldInstance).(*Bstruct))
						}
					}
				}
			}
		}
		if fieldName == "Dstruct4s" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Astruct) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Astruct)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Dstruct4s).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Dstruct4s = _inferedTypeInstance.Dstruct4s[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Dstruct4s =
								append(_inferedTypeInstance.Dstruct4s, any(fieldInstance).(*Dstruct))
						}
					}
				}
			}
		}
		if fieldName == "Anarrayofa" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Astruct) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Astruct)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Anarrayofa).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Anarrayofa = _inferedTypeInstance.Anarrayofa[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Anarrayofa =
								append(_inferedTypeInstance.Anarrayofa, any(fieldInstance).(*Astruct))
						}
					}
				}
			}
		}
		if fieldName == "Anotherarrayofb" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Astruct) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Astruct)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Anotherarrayofb).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Anotherarrayofb = _inferedTypeInstance.Anotherarrayofb[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Anotherarrayofb =
								append(_inferedTypeInstance.Anotherarrayofb, any(fieldInstance).(*Bstruct))
						}
					}
				}
			}
		}
		if fieldName == "AnarrayofbUse" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Astruct) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Astruct)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.AnarrayofbUse).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.AnarrayofbUse = _inferedTypeInstance.AnarrayofbUse[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.AnarrayofbUse =
								append(_inferedTypeInstance.AnarrayofbUse, any(fieldInstance).(*AstructBstructUse))
						}
					}
				}
			}
		}
		if fieldName == "Anarrayofb2Use" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Astruct) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Astruct)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Anarrayofb2Use).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Anarrayofb2Use = _inferedTypeInstance.Anarrayofb2Use[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Anarrayofb2Use =
								append(_inferedTypeInstance.Anarrayofb2Use, any(fieldInstance).(*AstructBstruct2Use))
						}
					}
				}
			}
		}

	case *AstructBstruct2Use:
		// insertion point per field

	case *AstructBstructUse:
		// insertion point per field

	case *Bstruct:
		// insertion point per field

	case *Dstruct:
		// insertion point per field
		if fieldName == "Anarrayofb" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Dstruct) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Dstruct)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Anarrayofb).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Anarrayofb = _inferedTypeInstance.Anarrayofb[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Anarrayofb =
								append(_inferedTypeInstance.Anarrayofb, any(fieldInstance).(*Bstruct))
						}
					}
				}
			}
		}

	case *Fstruct:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Astruct
	// insertion point per field
	clear(stage.Astruct_Anarrayofb_reverseMap)
	stage.Astruct_Anarrayofb_reverseMap = make(map[*Bstruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _bstruct := range astruct.Anarrayofb {
			stage.Astruct_Anarrayofb_reverseMap[_bstruct] = astruct
		}
	}
	clear(stage.Astruct_Dstruct4s_reverseMap)
	stage.Astruct_Dstruct4s_reverseMap = make(map[*Dstruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _dstruct := range astruct.Dstruct4s {
			stage.Astruct_Dstruct4s_reverseMap[_dstruct] = astruct
		}
	}
	clear(stage.Astruct_Anarrayofa_reverseMap)
	stage.Astruct_Anarrayofa_reverseMap = make(map[*Astruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _astruct := range astruct.Anarrayofa {
			stage.Astruct_Anarrayofa_reverseMap[_astruct] = astruct
		}
	}
	clear(stage.Astruct_Anotherarrayofb_reverseMap)
	stage.Astruct_Anotherarrayofb_reverseMap = make(map[*Bstruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _bstruct := range astruct.Anotherarrayofb {
			stage.Astruct_Anotherarrayofb_reverseMap[_bstruct] = astruct
		}
	}
	clear(stage.Astruct_AnarrayofbUse_reverseMap)
	stage.Astruct_AnarrayofbUse_reverseMap = make(map[*AstructBstructUse]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _astructbstructuse := range astruct.AnarrayofbUse {
			stage.Astruct_AnarrayofbUse_reverseMap[_astructbstructuse] = astruct
		}
	}
	clear(stage.Astruct_Anarrayofb2Use_reverseMap)
	stage.Astruct_Anarrayofb2Use_reverseMap = make(map[*AstructBstruct2Use]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _astructbstruct2use := range astruct.Anarrayofb2Use {
			stage.Astruct_Anarrayofb2Use_reverseMap[_astructbstruct2use] = astruct
		}
	}

	// Compute reverse map for named struct AstructBstruct2Use
	// insertion point per field

	// Compute reverse map for named struct AstructBstructUse
	// insertion point per field

	// Compute reverse map for named struct Bstruct
	// insertion point per field

	// Compute reverse map for named struct Dstruct
	// insertion point per field
	clear(stage.Dstruct_Anarrayofb_reverseMap)
	stage.Dstruct_Anarrayofb_reverseMap = make(map[*Bstruct]*Dstruct)
	for dstruct := range stage.Dstructs {
		_ = dstruct
		for _, _bstruct := range dstruct.Anarrayofb {
			stage.Dstruct_Anarrayofb_reverseMap[_bstruct] = dstruct
		}
	}

	// Compute reverse map for named struct Fstruct
	// insertion point per field

}
