// generated code - do not edit
package models

import "log"

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
	case *Astruct:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Anarrayofb" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Astruct)
				log.Println("Instance", _inferedTypeInstance.GetName(), "Anarrayofb", len(_inferedTypeInstance.Anarrayofb))
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Anarrayofb).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Anarrayofb = make([]*Bstruct, 0)
				if any(_instance).(*Astruct) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
				log.Println("Instance", _inferedTypeInstance.GetName(), "Anarrayofb", len(_inferedTypeInstance.Anarrayofb))
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Anotherarrayofb" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Astruct)
				log.Println("Instance", _inferedTypeInstance.GetName(), "Anotherarrayofb", len(_inferedTypeInstance.Anotherarrayofb))
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Anotherarrayofb).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Anotherarrayofb = make([]*Bstruct, 0)
				if any(_instance).(*Astruct) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
				log.Println("Instance", _inferedTypeInstance.GetName(), "Anotherarrayofb", len(_inferedTypeInstance.Anotherarrayofb))
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Anarrayofa" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Astruct)
				log.Println("Instance", _inferedTypeInstance.GetName(), "Anarrayofa", len(_inferedTypeInstance.Anarrayofa))
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Anarrayofa).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Anarrayofa = make([]*Astruct, 0)
				if any(_instance).(*Astruct) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
				log.Println("Instance", _inferedTypeInstance.GetName(), "Anarrayofa", len(_inferedTypeInstance.Anarrayofa))
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "AnarrayofbUse" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Astruct)
				log.Println("Instance", _inferedTypeInstance.GetName(), "AnarrayofbUse", len(_inferedTypeInstance.AnarrayofbUse))
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.AnarrayofbUse).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.AnarrayofbUse = make([]*AstructBstructUse, 0)
				if any(_instance).(*Astruct) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
				log.Println("Instance", _inferedTypeInstance.GetName(), "AnarrayofbUse", len(_inferedTypeInstance.AnarrayofbUse))
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Anarrayofb2Use" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Astruct)
				log.Println("Instance", _inferedTypeInstance.GetName(), "Anarrayofb2Use", len(_inferedTypeInstance.Anarrayofb2Use))
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Anarrayofb2Use).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Anarrayofb2Use = make([]*AstructBstruct2Use, 0)
				if any(_instance).(*Astruct) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
				log.Println("Instance", _inferedTypeInstance.GetName(), "Anarrayofb2Use", len(_inferedTypeInstance.Anarrayofb2Use))
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

	}
}
