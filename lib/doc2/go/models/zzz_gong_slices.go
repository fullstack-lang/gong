// generated code - do not edit
package models

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct AttributeShape
	// insertion point per field

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
	clear(stage.GongStructShape_AttributeShapes_reverseMap)
	stage.GongStructShape_AttributeShapes_reverseMap = make(map[*AttributeShape]*GongStructShape)
	for gongstructshape := range stage.GongStructShapes {
		_ = gongstructshape
		for _, _attributeshape := range gongstructshape.AttributeShapes {
			stage.GongStructShape_AttributeShapes_reverseMap[_attributeshape] = gongstructshape
		}
	}
	clear(stage.GongStructShape_LinkShapes_reverseMap)
	stage.GongStructShape_LinkShapes_reverseMap = make(map[*LinkShape]*GongStructShape)
	for gongstructshape := range stage.GongStructShapes {
		_ = gongstructshape
		for _, _linkshape := range gongstructshape.LinkShapes {
			stage.GongStructShape_LinkShapes_reverseMap[_linkshape] = gongstructshape
		}
	}

	// Compute reverse map for named struct LinkShape
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

}
