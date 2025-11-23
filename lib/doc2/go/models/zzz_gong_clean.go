// generated code - do not edit
package models

// Clean computes the reverse map, for all intances, for all clean to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) Clean() {
	// insertion point per named struct
	// Compute reverse map for named struct AttributeShape
	for attributeshape := range stage.AttributeShapes {
		_ = attributeshape
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct Classdiagram
	for classdiagram := range stage.Classdiagrams {
		_ = classdiagram
		// insertion point per field
		var _GongStructShapes []*GongStructShape
		for _, _gongstructshape := range classdiagram.GongStructShapes {
			if IsStaged(stage, _gongstructshape) {
			 	_GongStructShapes = append(_GongStructShapes, _gongstructshape)
			}
		}
		classdiagram.GongStructShapes = _GongStructShapes
		var _GongEnumShapes []*GongEnumShape
		for _, _gongenumshape := range classdiagram.GongEnumShapes {
			if IsStaged(stage, _gongenumshape) {
			 	_GongEnumShapes = append(_GongEnumShapes, _gongenumshape)
			}
		}
		classdiagram.GongEnumShapes = _GongEnumShapes
		var _GongNoteShapes []*GongNoteShape
		for _, _gongnoteshape := range classdiagram.GongNoteShapes {
			if IsStaged(stage, _gongnoteshape) {
			 	_GongNoteShapes = append(_GongNoteShapes, _gongnoteshape)
			}
		}
		classdiagram.GongNoteShapes = _GongNoteShapes
		// insertion point per field
	}

	// Compute reverse map for named struct DiagramPackage
	for diagrampackage := range stage.DiagramPackages {
		_ = diagrampackage
		// insertion point per field
		var _Classdiagrams []*Classdiagram
		for _, _classdiagram := range diagrampackage.Classdiagrams {
			if IsStaged(stage, _classdiagram) {
			 	_Classdiagrams = append(_Classdiagrams, _classdiagram)
			}
		}
		diagrampackage.Classdiagrams = _Classdiagrams
		// insertion point per field
		if !IsStaged(stage, diagrampackage.SelectedClassdiagram) {
			diagrampackage.SelectedClassdiagram = nil
		}
	}

	// Compute reverse map for named struct GongEnumShape
	for gongenumshape := range stage.GongEnumShapes {
		_ = gongenumshape
		// insertion point per field
		var _GongEnumValueShapes []*GongEnumValueShape
		for _, _gongenumvalueshape := range gongenumshape.GongEnumValueShapes {
			if IsStaged(stage, _gongenumvalueshape) {
			 	_GongEnumValueShapes = append(_GongEnumValueShapes, _gongenumvalueshape)
			}
		}
		gongenumshape.GongEnumValueShapes = _GongEnumValueShapes
		// insertion point per field
	}

	// Compute reverse map for named struct GongEnumValueShape
	for gongenumvalueshape := range stage.GongEnumValueShapes {
		_ = gongenumvalueshape
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct GongNoteLinkShape
	for gongnotelinkshape := range stage.GongNoteLinkShapes {
		_ = gongnotelinkshape
		// insertion point per field
		// insertion point per field
	}

	// Compute reverse map for named struct GongNoteShape
	for gongnoteshape := range stage.GongNoteShapes {
		_ = gongnoteshape
		// insertion point per field
		var _GongNoteLinkShapes []*GongNoteLinkShape
		for _, _gongnotelinkshape := range gongnoteshape.GongNoteLinkShapes {
			if IsStaged(stage, _gongnotelinkshape) {
			 	_GongNoteLinkShapes = append(_GongNoteLinkShapes, _gongnotelinkshape)
			}
		}
		gongnoteshape.GongNoteLinkShapes = _GongNoteLinkShapes
		// insertion point per field
	}

	// Compute reverse map for named struct GongStructShape
	for gongstructshape := range stage.GongStructShapes {
		_ = gongstructshape
		// insertion point per field
		var _AttributeShapes []*AttributeShape
		for _, _attributeshape := range gongstructshape.AttributeShapes {
			if IsStaged(stage, _attributeshape) {
			 	_AttributeShapes = append(_AttributeShapes, _attributeshape)
			}
		}
		gongstructshape.AttributeShapes = _AttributeShapes
		var _LinkShapes []*LinkShape
		for _, _linkshape := range gongstructshape.LinkShapes {
			if IsStaged(stage, _linkshape) {
			 	_LinkShapes = append(_LinkShapes, _linkshape)
			}
		}
		gongstructshape.LinkShapes = _LinkShapes
		// insertion point per field
	}

	// Compute reverse map for named struct LinkShape
	for linkshape := range stage.LinkShapes {
		_ = linkshape
		// insertion point per field
		// insertion point per field
	}

}
