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
	case *Classdiagram:
		// insertion point per field
		if fieldName == "GongStructShapes" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Classdiagram) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Classdiagram)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.GongStructShapes).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.GongStructShapes = _inferedTypeInstance.GongStructShapes[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.GongStructShapes =
								append(_inferedTypeInstance.GongStructShapes, any(fieldInstance).(*GongStructShape))
						}
					}
				}
			}
		}
		if fieldName == "GongEnumShapes" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Classdiagram) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Classdiagram)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.GongEnumShapes).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.GongEnumShapes = _inferedTypeInstance.GongEnumShapes[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.GongEnumShapes =
								append(_inferedTypeInstance.GongEnumShapes, any(fieldInstance).(*GongEnumShape))
						}
					}
				}
			}
		}
		if fieldName == "NoteShapes" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Classdiagram) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Classdiagram)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.NoteShapes).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.NoteShapes = _inferedTypeInstance.NoteShapes[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.NoteShapes =
								append(_inferedTypeInstance.NoteShapes, any(fieldInstance).(*NoteShape))
						}
					}
				}
			}
		}

	case *DiagramPackage:
		// insertion point per field
		if fieldName == "Classdiagrams" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DiagramPackage) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DiagramPackage)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Classdiagrams).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Classdiagrams = _inferedTypeInstance.Classdiagrams[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Classdiagrams =
								append(_inferedTypeInstance.Classdiagrams, any(fieldInstance).(*Classdiagram))
						}
					}
				}
			}
		}
		if fieldName == "Umlscs" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*DiagramPackage) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*DiagramPackage)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Umlscs).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Umlscs = _inferedTypeInstance.Umlscs[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Umlscs =
								append(_inferedTypeInstance.Umlscs, any(fieldInstance).(*Umlsc))
						}
					}
				}
			}
		}

	case *Field:
		// insertion point per field

	case *GongEnumShape:
		// insertion point per field
		if fieldName == "GongEnumValueEntrys" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*GongEnumShape) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*GongEnumShape)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.GongEnumValueEntrys).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.GongEnumValueEntrys = _inferedTypeInstance.GongEnumValueEntrys[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.GongEnumValueEntrys =
								append(_inferedTypeInstance.GongEnumValueEntrys, any(fieldInstance).(*GongEnumValueEntry))
						}
					}
				}
			}
		}

	case *GongEnumValueEntry:
		// insertion point per field

	case *GongStructShape:
		// insertion point per field
		if fieldName == "Fields" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*GongStructShape) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*GongStructShape)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Fields).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Fields = _inferedTypeInstance.Fields[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Fields =
								append(_inferedTypeInstance.Fields, any(fieldInstance).(*Field))
						}
					}
				}
			}
		}
		if fieldName == "Links" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*GongStructShape) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*GongStructShape)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.Links).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.Links = _inferedTypeInstance.Links[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.Links =
								append(_inferedTypeInstance.Links, any(fieldInstance).(*Link))
						}
					}
				}
			}
		}

	case *Link:
		// insertion point per field

	case *NoteShape:
		// insertion point per field
		if fieldName == "NoteShapeLinks" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*NoteShape) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*NoteShape)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.NoteShapeLinks).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.NoteShapeLinks = _inferedTypeInstance.NoteShapeLinks[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.NoteShapeLinks =
								append(_inferedTypeInstance.NoteShapeLinks, any(fieldInstance).(*NoteShapeLink))
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
		if fieldName == "States" {

			// walk all instances of the owning type
			for _instance := range *GetGongstructInstancesSetFromPointerType[OwningType](stage) {
				if any(_instance).(*Umlsc) != owningInstanceInfered {
					_inferedTypeInstance := any(_instance).(*Umlsc)
					reference := make([]FieldType, 0)
					targetFieldSlice := any(_inferedTypeInstance.States).([]FieldType)
					copy(targetFieldSlice, reference)
					_inferedTypeInstance.States = _inferedTypeInstance.States[0:]
					for _, fieldInstance := range reference {
						if _, ok := setOfFieldInstances[any(fieldInstance).(FieldType)]; !ok {
							_inferedTypeInstance.States =
								append(_inferedTypeInstance.States, any(fieldInstance).(*UmlState))
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

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *StageStruct) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct Classdiagram
	// insertion point per field
	clear(stage.Classdiagram_GongStructShapes_reverseMap)
	stage.Classdiagram_GongStructShapes_reverseMap = make(map[*GongStructShape]*Classdiagram)
	for classdiagram := range stage.Classdiagrams {
		_ = classdiagram
		for _, _gongstructshape := range classdiagram.GongStructShapes {
			stage.Classdiagram_GongStructShapes_reverseMap[_gongstructshape] = classdiagram
		}
	}
	clear(stage.Classdiagram_GongEnumShapes_reverseMap)
	stage.Classdiagram_GongEnumShapes_reverseMap = make(map[*GongEnumShape]*Classdiagram)
	for classdiagram := range stage.Classdiagrams {
		_ = classdiagram
		for _, _gongenumshape := range classdiagram.GongEnumShapes {
			stage.Classdiagram_GongEnumShapes_reverseMap[_gongenumshape] = classdiagram
		}
	}
	clear(stage.Classdiagram_NoteShapes_reverseMap)
	stage.Classdiagram_NoteShapes_reverseMap = make(map[*NoteShape]*Classdiagram)
	for classdiagram := range stage.Classdiagrams {
		_ = classdiagram
		for _, _noteshape := range classdiagram.NoteShapes {
			stage.Classdiagram_NoteShapes_reverseMap[_noteshape] = classdiagram
		}
	}

	// Compute reverse map for named struct DiagramPackage
	// insertion point per field
	clear(stage.DiagramPackage_Classdiagrams_reverseMap)
	stage.DiagramPackage_Classdiagrams_reverseMap = make(map[*Classdiagram]*DiagramPackage)
	for diagrampackage := range stage.DiagramPackages {
		_ = diagrampackage
		for _, _classdiagram := range diagrampackage.Classdiagrams {
			stage.DiagramPackage_Classdiagrams_reverseMap[_classdiagram] = diagrampackage
		}
	}
	clear(stage.DiagramPackage_Umlscs_reverseMap)
	stage.DiagramPackage_Umlscs_reverseMap = make(map[*Umlsc]*DiagramPackage)
	for diagrampackage := range stage.DiagramPackages {
		_ = diagrampackage
		for _, _umlsc := range diagrampackage.Umlscs {
			stage.DiagramPackage_Umlscs_reverseMap[_umlsc] = diagrampackage
		}
	}

	// Compute reverse map for named struct Field
	// insertion point per field

	// Compute reverse map for named struct GongEnumShape
	// insertion point per field
	clear(stage.GongEnumShape_GongEnumValueEntrys_reverseMap)
	stage.GongEnumShape_GongEnumValueEntrys_reverseMap = make(map[*GongEnumValueEntry]*GongEnumShape)
	for gongenumshape := range stage.GongEnumShapes {
		_ = gongenumshape
		for _, _gongenumvalueentry := range gongenumshape.GongEnumValueEntrys {
			stage.GongEnumShape_GongEnumValueEntrys_reverseMap[_gongenumvalueentry] = gongenumshape
		}
	}

	// Compute reverse map for named struct GongEnumValueEntry
	// insertion point per field

	// Compute reverse map for named struct GongStructShape
	// insertion point per field
	clear(stage.GongStructShape_Fields_reverseMap)
	stage.GongStructShape_Fields_reverseMap = make(map[*Field]*GongStructShape)
	for gongstructshape := range stage.GongStructShapes {
		_ = gongstructshape
		for _, _field := range gongstructshape.Fields {
			stage.GongStructShape_Fields_reverseMap[_field] = gongstructshape
		}
	}
	clear(stage.GongStructShape_Links_reverseMap)
	stage.GongStructShape_Links_reverseMap = make(map[*Link]*GongStructShape)
	for gongstructshape := range stage.GongStructShapes {
		_ = gongstructshape
		for _, _link := range gongstructshape.Links {
			stage.GongStructShape_Links_reverseMap[_link] = gongstructshape
		}
	}

	// Compute reverse map for named struct Link
	// insertion point per field

	// Compute reverse map for named struct NoteShape
	// insertion point per field
	clear(stage.NoteShape_NoteShapeLinks_reverseMap)
	stage.NoteShape_NoteShapeLinks_reverseMap = make(map[*NoteShapeLink]*NoteShape)
	for noteshape := range stage.NoteShapes {
		_ = noteshape
		for _, _noteshapelink := range noteshape.NoteShapeLinks {
			stage.NoteShape_NoteShapeLinks_reverseMap[_noteshapelink] = noteshape
		}
	}

	// Compute reverse map for named struct NoteShapeLink
	// insertion point per field

	// Compute reverse map for named struct Position
	// insertion point per field

	// Compute reverse map for named struct UmlState
	// insertion point per field

	// Compute reverse map for named struct Umlsc
	// insertion point per field
	clear(stage.Umlsc_States_reverseMap)
	stage.Umlsc_States_reverseMap = make(map[*UmlState]*Umlsc)
	for umlsc := range stage.Umlscs {
		_ = umlsc
		for _, _umlstate := range umlsc.States {
			stage.Umlsc_States_reverseMap[_umlstate] = umlsc
		}
	}

	// Compute reverse map for named struct Vertice
	// insertion point per field

}
