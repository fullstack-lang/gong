// generated code - do not edit
package models

import "time"
var __GongSliceTemplate_time__dummyDeclaration time.Duration
var _ = __GongSliceTemplate_time__dummyDeclaration

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct AttributeShape
	// insertion point per field

	// Compute reverse map for named struct Classdiagram
	// insertion point per field
	stage.Classdiagram_GongStructShapes_reverseMap = make(map[*GongStructShape]*Classdiagram)
	for classdiagram := range stage.Classdiagrams {
		_ = classdiagram
		for _, _gongstructshape := range classdiagram.GongStructShapes {
			stage.Classdiagram_GongStructShapes_reverseMap[_gongstructshape] = classdiagram
		}
	}
	stage.Classdiagram_GongEnumShapes_reverseMap = make(map[*GongEnumShape]*Classdiagram)
	for classdiagram := range stage.Classdiagrams {
		_ = classdiagram
		for _, _gongenumshape := range classdiagram.GongEnumShapes {
			stage.Classdiagram_GongEnumShapes_reverseMap[_gongenumshape] = classdiagram
		}
	}
	stage.Classdiagram_GongNoteShapes_reverseMap = make(map[*GongNoteShape]*Classdiagram)
	for classdiagram := range stage.Classdiagrams {
		_ = classdiagram
		for _, _gongnoteshape := range classdiagram.GongNoteShapes {
			stage.Classdiagram_GongNoteShapes_reverseMap[_gongnoteshape] = classdiagram
		}
	}

	// Compute reverse map for named struct DiagramPackage
	// insertion point per field
	stage.DiagramPackage_Classdiagrams_reverseMap = make(map[*Classdiagram]*DiagramPackage)
	for diagrampackage := range stage.DiagramPackages {
		_ = diagrampackage
		for _, _classdiagram := range diagrampackage.Classdiagrams {
			stage.DiagramPackage_Classdiagrams_reverseMap[_classdiagram] = diagrampackage
		}
	}

	// Compute reverse map for named struct GongEnumShape
	// insertion point per field
	stage.GongEnumShape_GongEnumValueShapes_reverseMap = make(map[*GongEnumValueShape]*GongEnumShape)
	for gongenumshape := range stage.GongEnumShapes {
		_ = gongenumshape
		for _, _gongenumvalueshape := range gongenumshape.GongEnumValueShapes {
			stage.GongEnumShape_GongEnumValueShapes_reverseMap[_gongenumvalueshape] = gongenumshape
		}
	}

	// Compute reverse map for named struct GongEnumValueShape
	// insertion point per field

	// Compute reverse map for named struct GongNoteLinkShape
	// insertion point per field

	// Compute reverse map for named struct GongNoteShape
	// insertion point per field
	stage.GongNoteShape_GongNoteLinkShapes_reverseMap = make(map[*GongNoteLinkShape]*GongNoteShape)
	for gongnoteshape := range stage.GongNoteShapes {
		_ = gongnoteshape
		for _, _gongnotelinkshape := range gongnoteshape.GongNoteLinkShapes {
			stage.GongNoteShape_GongNoteLinkShapes_reverseMap[_gongnotelinkshape] = gongnoteshape
		}
	}

	// Compute reverse map for named struct GongStructShape
	// insertion point per field
	stage.GongStructShape_AttributeShapes_reverseMap = make(map[*AttributeShape]*GongStructShape)
	for gongstructshape := range stage.GongStructShapes {
		_ = gongstructshape
		for _, _attributeshape := range gongstructshape.AttributeShapes {
			stage.GongStructShape_AttributeShapes_reverseMap[_attributeshape] = gongstructshape
		}
	}
	stage.GongStructShape_LinkShapes_reverseMap = make(map[*LinkShape]*GongStructShape)
	for gongstructshape := range stage.GongStructShapes {
		_ = gongstructshape
		for _, _linkshape := range gongstructshape.LinkShapes {
			stage.GongStructShape_LinkShapes_reverseMap[_linkshape] = gongstructshape
		}
	}

	// Compute reverse map for named struct LinkShape
	// insertion point per field

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.AttributeShapes {
		res = append(res, instance)
	}

	for instance := range stage.Classdiagrams {
		res = append(res, instance)
	}

	for instance := range stage.DiagramPackages {
		res = append(res, instance)
	}

	for instance := range stage.GongEnumShapes {
		res = append(res, instance)
	}

	for instance := range stage.GongEnumValueShapes {
		res = append(res, instance)
	}

	for instance := range stage.GongNoteLinkShapes {
		res = append(res, instance)
	}

	for instance := range stage.GongNoteShapes {
		res = append(res, instance)
	}

	for instance := range stage.GongStructShapes {
		res = append(res, instance)
	}

	for instance := range stage.LinkShapes {
		res = append(res, instance)
	}

	return
}


// insertion point per named struct
func (attributeshape *AttributeShape) GongCopy() GongstructIF {
	newInstance := *attributeshape
	return &newInstance
}

func (classdiagram *Classdiagram) GongCopy() GongstructIF {
	newInstance := *classdiagram
	return &newInstance
}

func (diagrampackage *DiagramPackage) GongCopy() GongstructIF {
	newInstance := *diagrampackage
	return &newInstance
}

func (gongenumshape *GongEnumShape) GongCopy() GongstructIF {
	newInstance := *gongenumshape
	return &newInstance
}

func (gongenumvalueshape *GongEnumValueShape) GongCopy() GongstructIF {
	newInstance := *gongenumvalueshape
	return &newInstance
}

func (gongnotelinkshape *GongNoteLinkShape) GongCopy() GongstructIF {
	newInstance := *gongnotelinkshape
	return &newInstance
}

func (gongnoteshape *GongNoteShape) GongCopy() GongstructIF {
	newInstance := *gongnoteshape
	return &newInstance
}

func (gongstructshape *GongStructShape) GongCopy() GongstructIF {
	newInstance := *gongstructshape
	return &newInstance
}

func (linkshape *LinkShape) GongCopy() GongstructIF {
	newInstance := *linkshape
	return &newInstance
}


func (stage *Stage) ComputeDifference() {
	var lenNewInstances int
	var lenDeletedInstances int
	
	// insertion point per named struct
	var attributeshapes_newInstances []*AttributeShape
	var attributeshapes_deletedInstances []*AttributeShape

	// parse all staged instances and check if they have a reference
	for attributeshape := range stage.AttributeShapes {
		if _, ok := stage.AttributeShapes_reference[attributeshape]; !ok {
			attributeshapes_newInstances = append(attributeshapes_newInstances, attributeshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of AttributeShape "+attributeshape.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for attributeshape := range stage.AttributeShapes_reference {
		if _, ok := stage.AttributeShapes[attributeshape]; !ok {
			attributeshapes_deletedInstances = append(attributeshapes_deletedInstances, attributeshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of AttributeShape "+attributeshape.Name,
				)
			}
		}
	}

	lenNewInstances += len(attributeshapes_newInstances)
	lenDeletedInstances += len(attributeshapes_deletedInstances)
	var classdiagrams_newInstances []*Classdiagram
	var classdiagrams_deletedInstances []*Classdiagram

	// parse all staged instances and check if they have a reference
	for classdiagram := range stage.Classdiagrams {
		if _, ok := stage.Classdiagrams_reference[classdiagram]; !ok {
			classdiagrams_newInstances = append(classdiagrams_newInstances, classdiagram)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of Classdiagram "+classdiagram.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for classdiagram := range stage.Classdiagrams_reference {
		if _, ok := stage.Classdiagrams[classdiagram]; !ok {
			classdiagrams_deletedInstances = append(classdiagrams_deletedInstances, classdiagram)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of Classdiagram "+classdiagram.Name,
				)
			}
		}
	}

	lenNewInstances += len(classdiagrams_newInstances)
	lenDeletedInstances += len(classdiagrams_deletedInstances)
	var diagrampackages_newInstances []*DiagramPackage
	var diagrampackages_deletedInstances []*DiagramPackage

	// parse all staged instances and check if they have a reference
	for diagrampackage := range stage.DiagramPackages {
		if _, ok := stage.DiagramPackages_reference[diagrampackage]; !ok {
			diagrampackages_newInstances = append(diagrampackages_newInstances, diagrampackage)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of DiagramPackage "+diagrampackage.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for diagrampackage := range stage.DiagramPackages_reference {
		if _, ok := stage.DiagramPackages[diagrampackage]; !ok {
			diagrampackages_deletedInstances = append(diagrampackages_deletedInstances, diagrampackage)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of DiagramPackage "+diagrampackage.Name,
				)
			}
		}
	}

	lenNewInstances += len(diagrampackages_newInstances)
	lenDeletedInstances += len(diagrampackages_deletedInstances)
	var gongenumshapes_newInstances []*GongEnumShape
	var gongenumshapes_deletedInstances []*GongEnumShape

	// parse all staged instances and check if they have a reference
	for gongenumshape := range stage.GongEnumShapes {
		if _, ok := stage.GongEnumShapes_reference[gongenumshape]; !ok {
			gongenumshapes_newInstances = append(gongenumshapes_newInstances, gongenumshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of GongEnumShape "+gongenumshape.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongenumshape := range stage.GongEnumShapes_reference {
		if _, ok := stage.GongEnumShapes[gongenumshape]; !ok {
			gongenumshapes_deletedInstances = append(gongenumshapes_deletedInstances, gongenumshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of GongEnumShape "+gongenumshape.Name,
				)
			}
		}
	}

	lenNewInstances += len(gongenumshapes_newInstances)
	lenDeletedInstances += len(gongenumshapes_deletedInstances)
	var gongenumvalueshapes_newInstances []*GongEnumValueShape
	var gongenumvalueshapes_deletedInstances []*GongEnumValueShape

	// parse all staged instances and check if they have a reference
	for gongenumvalueshape := range stage.GongEnumValueShapes {
		if _, ok := stage.GongEnumValueShapes_reference[gongenumvalueshape]; !ok {
			gongenumvalueshapes_newInstances = append(gongenumvalueshapes_newInstances, gongenumvalueshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of GongEnumValueShape "+gongenumvalueshape.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongenumvalueshape := range stage.GongEnumValueShapes_reference {
		if _, ok := stage.GongEnumValueShapes[gongenumvalueshape]; !ok {
			gongenumvalueshapes_deletedInstances = append(gongenumvalueshapes_deletedInstances, gongenumvalueshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of GongEnumValueShape "+gongenumvalueshape.Name,
				)
			}
		}
	}

	lenNewInstances += len(gongenumvalueshapes_newInstances)
	lenDeletedInstances += len(gongenumvalueshapes_deletedInstances)
	var gongnotelinkshapes_newInstances []*GongNoteLinkShape
	var gongnotelinkshapes_deletedInstances []*GongNoteLinkShape

	// parse all staged instances and check if they have a reference
	for gongnotelinkshape := range stage.GongNoteLinkShapes {
		if _, ok := stage.GongNoteLinkShapes_reference[gongnotelinkshape]; !ok {
			gongnotelinkshapes_newInstances = append(gongnotelinkshapes_newInstances, gongnotelinkshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of GongNoteLinkShape "+gongnotelinkshape.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongnotelinkshape := range stage.GongNoteLinkShapes_reference {
		if _, ok := stage.GongNoteLinkShapes[gongnotelinkshape]; !ok {
			gongnotelinkshapes_deletedInstances = append(gongnotelinkshapes_deletedInstances, gongnotelinkshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of GongNoteLinkShape "+gongnotelinkshape.Name,
				)
			}
		}
	}

	lenNewInstances += len(gongnotelinkshapes_newInstances)
	lenDeletedInstances += len(gongnotelinkshapes_deletedInstances)
	var gongnoteshapes_newInstances []*GongNoteShape
	var gongnoteshapes_deletedInstances []*GongNoteShape

	// parse all staged instances and check if they have a reference
	for gongnoteshape := range stage.GongNoteShapes {
		if _, ok := stage.GongNoteShapes_reference[gongnoteshape]; !ok {
			gongnoteshapes_newInstances = append(gongnoteshapes_newInstances, gongnoteshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of GongNoteShape "+gongnoteshape.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongnoteshape := range stage.GongNoteShapes_reference {
		if _, ok := stage.GongNoteShapes[gongnoteshape]; !ok {
			gongnoteshapes_deletedInstances = append(gongnoteshapes_deletedInstances, gongnoteshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of GongNoteShape "+gongnoteshape.Name,
				)
			}
		}
	}

	lenNewInstances += len(gongnoteshapes_newInstances)
	lenDeletedInstances += len(gongnoteshapes_deletedInstances)
	var gongstructshapes_newInstances []*GongStructShape
	var gongstructshapes_deletedInstances []*GongStructShape

	// parse all staged instances and check if they have a reference
	for gongstructshape := range stage.GongStructShapes {
		if _, ok := stage.GongStructShapes_reference[gongstructshape]; !ok {
			gongstructshapes_newInstances = append(gongstructshapes_newInstances, gongstructshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of GongStructShape "+gongstructshape.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for gongstructshape := range stage.GongStructShapes_reference {
		if _, ok := stage.GongStructShapes[gongstructshape]; !ok {
			gongstructshapes_deletedInstances = append(gongstructshapes_deletedInstances, gongstructshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of GongStructShape "+gongstructshape.Name,
				)
			}
		}
	}

	lenNewInstances += len(gongstructshapes_newInstances)
	lenDeletedInstances += len(gongstructshapes_deletedInstances)
	var linkshapes_newInstances []*LinkShape
	var linkshapes_deletedInstances []*LinkShape

	// parse all staged instances and check if they have a reference
	for linkshape := range stage.LinkShapes {
		if _, ok := stage.LinkShapes_reference[linkshape]; !ok {
			linkshapes_newInstances = append(linkshapes_newInstances, linkshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"New instance of LinkShape "+linkshape.Name,
				)
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for linkshape := range stage.LinkShapes_reference {
		if _, ok := stage.LinkShapes[linkshape]; !ok {
			linkshapes_deletedInstances = append(linkshapes_deletedInstances, linkshape)
			if stage.GetProbeIF() != nil {
				stage.GetProbeIF().AddNotification(
					time.Now(),
					"Deleted instance of LinkShape "+linkshape.Name,
				)
			}
		}
	}

	lenNewInstances += len(linkshapes_newInstances)
	lenDeletedInstances += len(linkshapes_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 {
		if stage.GetProbeIF() != nil {
			stage.GetProbeIF().CommitNotificationTable()
		}
	}
}

// ComputeReference will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReference() {

	// insertion point per named struct
	stage.AttributeShapes_reference = make(map[*AttributeShape]*AttributeShape)
	for instance := range stage.AttributeShapes {
		stage.AttributeShapes_reference[instance] = instance
	}

	stage.Classdiagrams_reference = make(map[*Classdiagram]*Classdiagram)
	for instance := range stage.Classdiagrams {
		stage.Classdiagrams_reference[instance] = instance
	}

	stage.DiagramPackages_reference = make(map[*DiagramPackage]*DiagramPackage)
	for instance := range stage.DiagramPackages {
		stage.DiagramPackages_reference[instance] = instance
	}

	stage.GongEnumShapes_reference = make(map[*GongEnumShape]*GongEnumShape)
	for instance := range stage.GongEnumShapes {
		stage.GongEnumShapes_reference[instance] = instance
	}

	stage.GongEnumValueShapes_reference = make(map[*GongEnumValueShape]*GongEnumValueShape)
	for instance := range stage.GongEnumValueShapes {
		stage.GongEnumValueShapes_reference[instance] = instance
	}

	stage.GongNoteLinkShapes_reference = make(map[*GongNoteLinkShape]*GongNoteLinkShape)
	for instance := range stage.GongNoteLinkShapes {
		stage.GongNoteLinkShapes_reference[instance] = instance
	}

	stage.GongNoteShapes_reference = make(map[*GongNoteShape]*GongNoteShape)
	for instance := range stage.GongNoteShapes {
		stage.GongNoteShapes_reference[instance] = instance
	}

	stage.GongStructShapes_reference = make(map[*GongStructShape]*GongStructShape)
	for instance := range stage.GongStructShapes {
		stage.GongStructShapes_reference[instance] = instance
	}

	stage.LinkShapes_reference = make(map[*LinkShape]*LinkShape)
	for instance := range stage.LinkShapes {
		stage.LinkShapes_reference[instance] = instance
	}

}
