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
	case *Classdiagram:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "GongStructShapes" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Classdiagram)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.GongStructShapes).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.GongStructShapes = make([]*GongStructShape, 0)
				if any(_instance).(*Classdiagram) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "GongEnumShapes" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Classdiagram)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.GongEnumShapes).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.GongEnumShapes = make([]*GongEnumShape, 0)
				if any(_instance).(*Classdiagram) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "NoteShapes" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Classdiagram)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.NoteShapes).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.NoteShapes = make([]*NoteShape, 0)
				if any(_instance).(*Classdiagram) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *DiagramPackage:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Classdiagrams" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*DiagramPackage)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Classdiagrams).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Classdiagrams = make([]*Classdiagram, 0)
				if any(_instance).(*DiagramPackage) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Umlscs" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*DiagramPackage)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Umlscs).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Umlscs = make([]*Umlsc, 0)
				if any(_instance).(*DiagramPackage) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *Field:
		// insertion point per field

	case *GongEnumShape:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "GongEnumValueEntrys" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*GongEnumShape)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.GongEnumValueEntrys).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.GongEnumValueEntrys = make([]*GongEnumValueEntry, 0)
				if any(_instance).(*GongEnumShape) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *GongEnumValueEntry:
		// insertion point per field

	case *GongStructShape:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "Fields" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*GongStructShape)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Fields).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Fields = make([]*Field, 0)
				if any(_instance).(*GongStructShape) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}
		// tweaking, it might be streamlined
		if fieldName == "Links" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*GongStructShape)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.Links).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.Links = make([]*Link, 0)
				if any(_instance).(*GongStructShape) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *Link:
		// insertion point per field

	case *NoteShape:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "NoteShapeLinks" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*NoteShape)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.NoteShapeLinks).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.NoteShapeLinks = make([]*NoteShapeLink, 0)
				if any(_instance).(*NoteShape) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *NoteShapeLink:
		// insertion point per field

	case *Position:
		// insertion point per field

	case *UmlState:
		// insertion point per field

	case *Umlsc:
		// insertion point per field
		// tweaking, it might be streamlined
		if fieldName == "States" {
			for _instance := range *GetGongstructInstancesSetFromPointerType[T](stage) {
				_inferedTypeInstance := any(_instance).(*Umlsc)
				reference := make([]TF, 0)
				targetFieldSlice := any(_inferedTypeInstance.States).([]TF)
				copy(targetFieldSlice, reference)
				_inferedTypeInstance.States = make([]*UmlState, 0)
				if any(_instance).(*Umlsc) != owningInstanceInfered {
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(TF)]; !ok {
							targetFieldSlice = append(targetFieldSlice, fieldInstance)
						}
					}
				}
			}
		}

	case *Vertice:
		// insertion point per field

	default:
		_ = owningInstanceInfered // to avoid "declared and not used" error if no named struct has slices
	}
}
