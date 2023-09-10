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
// Note : algo is in O(N) of nb of Astruct and Bstruct instances
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
	case *Astruct:

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
	}
}
