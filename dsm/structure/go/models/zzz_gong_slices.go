// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct DiagramStructure
	// insertion point per field
	stage.DiagramStructure_System_Shapes_reverseMap = make(map[*SystemShape]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _systemshape := range diagramstructure.System_Shapes {
			stage.DiagramStructure_System_Shapes_reverseMap[_systemshape] = diagramstructure
		}
	}
	stage.DiagramStructure_Part_Shapes_reverseMap = make(map[*PartShape]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _partshape := range diagramstructure.Part_Shapes {
			stage.DiagramStructure_Part_Shapes_reverseMap[_partshape] = diagramstructure
		}
	}
	stage.DiagramStructure_PartsWhoseNodeIsExpanded_reverseMap = make(map[*Part]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _part := range diagramstructure.PartsWhoseNodeIsExpanded {
			stage.DiagramStructure_PartsWhoseNodeIsExpanded_reverseMap[_part] = diagramstructure
		}
	}
	stage.DiagramStructure_Link_Shapes_reverseMap = make(map[*LinkAssociationShape]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _linkassociationshape := range diagramstructure.Link_Shapes {
			stage.DiagramStructure_Link_Shapes_reverseMap[_linkassociationshape] = diagramstructure
		}
	}
	stage.DiagramStructure_LinksWhoseNodeIsExpanded_reverseMap = make(map[*Link]*DiagramStructure)
	for diagramstructure := range stage.DiagramStructures {
		_ = diagramstructure
		for _, _link := range diagramstructure.LinksWhoseNodeIsExpanded {
			stage.DiagramStructure_LinksWhoseNodeIsExpanded_reverseMap[_link] = diagramstructure
		}
	}

	// Compute reverse map for named struct Library
	// insertion point per field
	stage.Library_SubLibraries_reverseMap = make(map[*Library]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _library := range library.SubLibraries {
			stage.Library_SubLibraries_reverseMap[_library] = library
		}
	}
	stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap = make(map[*Library]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
			stage.Library_SubLibrariesWhoseNodeIsExpanded_reverseMap[_library] = library
		}
	}
	stage.Library_RootSystems_reverseMap = make(map[*System]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _system := range library.RootSystems {
			stage.Library_RootSystems_reverseMap[_system] = library
		}
	}
	stage.Library_SystemsWhoseNodeIsExpanded_reverseMap = make(map[*System]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _system := range library.SystemsWhoseNodeIsExpanded {
			stage.Library_SystemsWhoseNodeIsExpanded_reverseMap[_system] = library
		}
	}

	// Compute reverse map for named struct Link
	// insertion point per field

	// Compute reverse map for named struct LinkAssociationShape
	// insertion point per field

	// Compute reverse map for named struct Part
	// insertion point per field

	// Compute reverse map for named struct PartShape
	// insertion point per field

	// Compute reverse map for named struct System
	// insertion point per field
	stage.System_Parts_reverseMap = make(map[*Part]*System)
	for system := range stage.Systems {
		_ = system
		for _, _part := range system.Parts {
			stage.System_Parts_reverseMap[_part] = system
		}
	}
	stage.System_PartsWhoseNodeIsExpanded_reverseMap = make(map[*Part]*System)
	for system := range stage.Systems {
		_ = system
		for _, _part := range system.PartsWhoseNodeIsExpanded {
			stage.System_PartsWhoseNodeIsExpanded_reverseMap[_part] = system
		}
	}
	stage.System_SubSystems_reverseMap = make(map[*System]*System)
	for system := range stage.Systems {
		_ = system
		for _, _system := range system.SubSystems {
			stage.System_SubSystems_reverseMap[_system] = system
		}
	}
	stage.System_SubSystemsWhoseNodeIsExpanded_reverseMap = make(map[*System]*System)
	for system := range stage.Systems {
		_ = system
		for _, _system := range system.SubSystemsWhoseNodeIsExpanded {
			stage.System_SubSystemsWhoseNodeIsExpanded_reverseMap[_system] = system
		}
	}
	stage.System_Links_reverseMap = make(map[*Link]*System)
	for system := range stage.Systems {
		_ = system
		for _, _link := range system.Links {
			stage.System_Links_reverseMap[_link] = system
		}
	}
	stage.System_LinksWhoseNodeIsExpanded_reverseMap = make(map[*Link]*System)
	for system := range stage.Systems {
		_ = system
		for _, _link := range system.LinksWhoseNodeIsExpanded {
			stage.System_LinksWhoseNodeIsExpanded_reverseMap[_link] = system
		}
	}
	stage.System_DiagramStructures_reverseMap = make(map[*DiagramStructure]*System)
	for system := range stage.Systems {
		_ = system
		for _, _diagramstructure := range system.DiagramStructures {
			stage.System_DiagramStructures_reverseMap[_diagramstructure] = system
		}
	}
	stage.System_DiagramStructuresWhoseNodeIsExpanded_reverseMap = make(map[*DiagramStructure]*System)
	for system := range stage.Systems {
		_ = system
		for _, _diagramstructure := range system.DiagramStructuresWhoseNodeIsExpanded {
			stage.System_DiagramStructuresWhoseNodeIsExpanded_reverseMap[_diagramstructure] = system
		}
	}

	// Compute reverse map for named struct SystemShape
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.DiagramStructures {
		res = append(res, instance)
	}

	for instance := range stage.Librarys {
		res = append(res, instance)
	}

	for instance := range stage.Links {
		res = append(res, instance)
	}

	for instance := range stage.LinkAssociationShapes {
		res = append(res, instance)
	}

	for instance := range stage.Parts {
		res = append(res, instance)
	}

	for instance := range stage.PartShapes {
		res = append(res, instance)
	}

	for instance := range stage.Systems {
		res = append(res, instance)
	}

	for instance := range stage.SystemShapes {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (diagramstructure *DiagramStructure) GongCopy() GongstructIF {
	newInstance := new(DiagramStructure)
	diagramstructure.CopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.CopyBasicFields(newInstance)
	return newInstance
}

func (link *Link) GongCopy() GongstructIF {
	newInstance := new(Link)
	link.CopyBasicFields(newInstance)
	return newInstance
}

func (linkassociationshape *LinkAssociationShape) GongCopy() GongstructIF {
	newInstance := new(LinkAssociationShape)
	linkassociationshape.CopyBasicFields(newInstance)
	return newInstance
}

func (part *Part) GongCopy() GongstructIF {
	newInstance := new(Part)
	part.CopyBasicFields(newInstance)
	return newInstance
}

func (partshape *PartShape) GongCopy() GongstructIF {
	newInstance := new(PartShape)
	partshape.CopyBasicFields(newInstance)
	return newInstance
}

func (system *System) GongCopy() GongstructIF {
	newInstance := new(System)
	system.CopyBasicFields(newInstance)
	return newInstance
}

func (systemshape *SystemShape) GongCopy() GongstructIF {
	newInstance := new(SystemShape)
	systemshape.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (diagramstructure *DiagramStructure) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagramstructure).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(diagramstructure), uint64(GetOrderPointerGongstruct(stage, diagramstructure)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(library), uint64(GetOrderPointerGongstruct(stage, library)))
	return
}

func (link *Link) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(link).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(link), uint64(GetOrderPointerGongstruct(stage, link)))
	return
}

func (linkassociationshape *LinkAssociationShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(linkassociationshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(linkassociationshape), uint64(GetOrderPointerGongstruct(stage, linkassociationshape)))
	return
}

func (part *Part) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(part).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(part), uint64(GetOrderPointerGongstruct(stage, part)))
	return
}

func (partshape *PartShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(partshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(partshape), uint64(GetOrderPointerGongstruct(stage, partshape)))
	return
}

func (system *System) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(system).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(system), uint64(GetOrderPointerGongstruct(stage, system)))
	return
}

func (systemshape *SystemShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(systemshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(systemshape), uint64(GetOrderPointerGongstruct(stage, systemshape)))
	return
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	var diagramstructures_newInstances []*DiagramStructure
	var diagramstructures_deletedInstances []*DiagramStructure

	// parse all staged instances and check if they have a reference
	for diagramstructure := range stage.DiagramStructures {
		if ref, ok := stage.DiagramStructures_reference[diagramstructure]; !ok {
			diagramstructures_newInstances = append(diagramstructures_newInstances, diagramstructure)
			newInstancesSlice = append(newInstancesSlice, diagramstructure.GongMarshallIdentifier(stage))
			if stage.DiagramStructures_referenceOrder == nil {
				stage.DiagramStructures_referenceOrder = make(map[*DiagramStructure]uint)
			}
			stage.DiagramStructures_referenceOrder[diagramstructure] = stage.DiagramStructure_stagedOrder[diagramstructure]
			newInstancesReverseSlice = append(newInstancesReverseSlice, diagramstructure.GongMarshallUnstaging(stage))
			// delete(stage.DiagramStructures_referenceOrder, diagramstructure)
			fieldInitializers, pointersInitializations := diagramstructure.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DiagramStructure_stagedOrder[ref] = stage.DiagramStructure_stagedOrder[diagramstructure]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := diagramstructure.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, diagramstructure)
			// delete(stage.DiagramStructure_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if diagramstructure.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", diagramstructure.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.DiagramStructures_reference {
		instance := stage.DiagramStructures_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DiagramStructures[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			diagramstructures_deletedInstances = append(diagramstructures_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(diagramstructures_newInstances)
	lenDeletedInstances += len(diagramstructures_deletedInstances)
	var librarys_newInstances []*Library
	var librarys_deletedInstances []*Library

	// parse all staged instances and check if they have a reference
	for library := range stage.Librarys {
		if ref, ok := stage.Librarys_reference[library]; !ok {
			librarys_newInstances = append(librarys_newInstances, library)
			newInstancesSlice = append(newInstancesSlice, library.GongMarshallIdentifier(stage))
			if stage.Librarys_referenceOrder == nil {
				stage.Librarys_referenceOrder = make(map[*Library]uint)
			}
			stage.Librarys_referenceOrder[library] = stage.Library_stagedOrder[library]
			newInstancesReverseSlice = append(newInstancesReverseSlice, library.GongMarshallUnstaging(stage))
			// delete(stage.Librarys_referenceOrder, library)
			fieldInitializers, pointersInitializations := library.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Library_stagedOrder[ref] = stage.Library_stagedOrder[library]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := library.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, library)
			// delete(stage.Library_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if library.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", library.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Librarys_reference {
		instance := stage.Librarys_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Librarys[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			librarys_deletedInstances = append(librarys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(librarys_newInstances)
	lenDeletedInstances += len(librarys_deletedInstances)
	var links_newInstances []*Link
	var links_deletedInstances []*Link

	// parse all staged instances and check if they have a reference
	for link := range stage.Links {
		if ref, ok := stage.Links_reference[link]; !ok {
			links_newInstances = append(links_newInstances, link)
			newInstancesSlice = append(newInstancesSlice, link.GongMarshallIdentifier(stage))
			if stage.Links_referenceOrder == nil {
				stage.Links_referenceOrder = make(map[*Link]uint)
			}
			stage.Links_referenceOrder[link] = stage.Link_stagedOrder[link]
			newInstancesReverseSlice = append(newInstancesReverseSlice, link.GongMarshallUnstaging(stage))
			// delete(stage.Links_referenceOrder, link)
			fieldInitializers, pointersInitializations := link.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Link_stagedOrder[ref] = stage.Link_stagedOrder[link]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := link.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, link)
			// delete(stage.Link_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if link.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", link.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Links_reference {
		instance := stage.Links_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Links[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			links_deletedInstances = append(links_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(links_newInstances)
	lenDeletedInstances += len(links_deletedInstances)
	var linkassociationshapes_newInstances []*LinkAssociationShape
	var linkassociationshapes_deletedInstances []*LinkAssociationShape

	// parse all staged instances and check if they have a reference
	for linkassociationshape := range stage.LinkAssociationShapes {
		if ref, ok := stage.LinkAssociationShapes_reference[linkassociationshape]; !ok {
			linkassociationshapes_newInstances = append(linkassociationshapes_newInstances, linkassociationshape)
			newInstancesSlice = append(newInstancesSlice, linkassociationshape.GongMarshallIdentifier(stage))
			if stage.LinkAssociationShapes_referenceOrder == nil {
				stage.LinkAssociationShapes_referenceOrder = make(map[*LinkAssociationShape]uint)
			}
			stage.LinkAssociationShapes_referenceOrder[linkassociationshape] = stage.LinkAssociationShape_stagedOrder[linkassociationshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, linkassociationshape.GongMarshallUnstaging(stage))
			// delete(stage.LinkAssociationShapes_referenceOrder, linkassociationshape)
			fieldInitializers, pointersInitializations := linkassociationshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.LinkAssociationShape_stagedOrder[ref] = stage.LinkAssociationShape_stagedOrder[linkassociationshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := linkassociationshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, linkassociationshape)
			// delete(stage.LinkAssociationShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if linkassociationshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", linkassociationshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.LinkAssociationShapes_reference {
		instance := stage.LinkAssociationShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.LinkAssociationShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			linkassociationshapes_deletedInstances = append(linkassociationshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(linkassociationshapes_newInstances)
	lenDeletedInstances += len(linkassociationshapes_deletedInstances)
	var parts_newInstances []*Part
	var parts_deletedInstances []*Part

	// parse all staged instances and check if they have a reference
	for part := range stage.Parts {
		if ref, ok := stage.Parts_reference[part]; !ok {
			parts_newInstances = append(parts_newInstances, part)
			newInstancesSlice = append(newInstancesSlice, part.GongMarshallIdentifier(stage))
			if stage.Parts_referenceOrder == nil {
				stage.Parts_referenceOrder = make(map[*Part]uint)
			}
			stage.Parts_referenceOrder[part] = stage.Part_stagedOrder[part]
			newInstancesReverseSlice = append(newInstancesReverseSlice, part.GongMarshallUnstaging(stage))
			// delete(stage.Parts_referenceOrder, part)
			fieldInitializers, pointersInitializations := part.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Part_stagedOrder[ref] = stage.Part_stagedOrder[part]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := part.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, part)
			// delete(stage.Part_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if part.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", part.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Parts_reference {
		instance := stage.Parts_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Parts[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			parts_deletedInstances = append(parts_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(parts_newInstances)
	lenDeletedInstances += len(parts_deletedInstances)
	var partshapes_newInstances []*PartShape
	var partshapes_deletedInstances []*PartShape

	// parse all staged instances and check if they have a reference
	for partshape := range stage.PartShapes {
		if ref, ok := stage.PartShapes_reference[partshape]; !ok {
			partshapes_newInstances = append(partshapes_newInstances, partshape)
			newInstancesSlice = append(newInstancesSlice, partshape.GongMarshallIdentifier(stage))
			if stage.PartShapes_referenceOrder == nil {
				stage.PartShapes_referenceOrder = make(map[*PartShape]uint)
			}
			stage.PartShapes_referenceOrder[partshape] = stage.PartShape_stagedOrder[partshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, partshape.GongMarshallUnstaging(stage))
			// delete(stage.PartShapes_referenceOrder, partshape)
			fieldInitializers, pointersInitializations := partshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PartShape_stagedOrder[ref] = stage.PartShape_stagedOrder[partshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := partshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, partshape)
			// delete(stage.PartShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if partshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", partshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.PartShapes_reference {
		instance := stage.PartShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PartShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			partshapes_deletedInstances = append(partshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(partshapes_newInstances)
	lenDeletedInstances += len(partshapes_deletedInstances)
	var systems_newInstances []*System
	var systems_deletedInstances []*System

	// parse all staged instances and check if they have a reference
	for system := range stage.Systems {
		if ref, ok := stage.Systems_reference[system]; !ok {
			systems_newInstances = append(systems_newInstances, system)
			newInstancesSlice = append(newInstancesSlice, system.GongMarshallIdentifier(stage))
			if stage.Systems_referenceOrder == nil {
				stage.Systems_referenceOrder = make(map[*System]uint)
			}
			stage.Systems_referenceOrder[system] = stage.System_stagedOrder[system]
			newInstancesReverseSlice = append(newInstancesReverseSlice, system.GongMarshallUnstaging(stage))
			// delete(stage.Systems_referenceOrder, system)
			fieldInitializers, pointersInitializations := system.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.System_stagedOrder[ref] = stage.System_stagedOrder[system]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := system.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, system)
			// delete(stage.System_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if system.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", system.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.Systems_reference {
		instance := stage.Systems_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Systems[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			systems_deletedInstances = append(systems_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(systems_newInstances)
	lenDeletedInstances += len(systems_deletedInstances)
	var systemshapes_newInstances []*SystemShape
	var systemshapes_deletedInstances []*SystemShape

	// parse all staged instances and check if they have a reference
	for systemshape := range stage.SystemShapes {
		if ref, ok := stage.SystemShapes_reference[systemshape]; !ok {
			systemshapes_newInstances = append(systemshapes_newInstances, systemshape)
			newInstancesSlice = append(newInstancesSlice, systemshape.GongMarshallIdentifier(stage))
			if stage.SystemShapes_referenceOrder == nil {
				stage.SystemShapes_referenceOrder = make(map[*SystemShape]uint)
			}
			stage.SystemShapes_referenceOrder[systemshape] = stage.SystemShape_stagedOrder[systemshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, systemshape.GongMarshallUnstaging(stage))
			// delete(stage.SystemShapes_referenceOrder, systemshape)
			fieldInitializers, pointersInitializations := systemshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SystemShape_stagedOrder[ref] = stage.SystemShape_stagedOrder[systemshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := systemshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, systemshape)
			// delete(stage.SystemShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if systemshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", systemshape.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range stage.SystemShapes_reference {
		instance := stage.SystemShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.SystemShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			systemshapes_deletedInstances = append(systemshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(systemshapes_newInstances)
	lenDeletedInstances += len(systemshapes_deletedInstances)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	stage.DiagramStructures_reference = make(map[*DiagramStructure]*DiagramStructure)
	stage.DiagramStructures_referenceOrder = make(map[*DiagramStructure]uint) // diff Unstage needs the reference order
	stage.DiagramStructures_instance = make(map[*DiagramStructure]*DiagramStructure)
	for instance := range stage.DiagramStructures {
		_copy := instance.GongCopy().(*DiagramStructure)
		stage.DiagramStructures_reference[instance] = _copy
		stage.DiagramStructures_instance[_copy] = instance
		stage.DiagramStructures_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Librarys_reference = make(map[*Library]*Library)
	stage.Librarys_referenceOrder = make(map[*Library]uint) // diff Unstage needs the reference order
	stage.Librarys_instance = make(map[*Library]*Library)
	for instance := range stage.Librarys {
		_copy := instance.GongCopy().(*Library)
		stage.Librarys_reference[instance] = _copy
		stage.Librarys_instance[_copy] = instance
		stage.Librarys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Links_reference = make(map[*Link]*Link)
	stage.Links_referenceOrder = make(map[*Link]uint) // diff Unstage needs the reference order
	stage.Links_instance = make(map[*Link]*Link)
	for instance := range stage.Links {
		_copy := instance.GongCopy().(*Link)
		stage.Links_reference[instance] = _copy
		stage.Links_instance[_copy] = instance
		stage.Links_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.LinkAssociationShapes_reference = make(map[*LinkAssociationShape]*LinkAssociationShape)
	stage.LinkAssociationShapes_referenceOrder = make(map[*LinkAssociationShape]uint) // diff Unstage needs the reference order
	stage.LinkAssociationShapes_instance = make(map[*LinkAssociationShape]*LinkAssociationShape)
	for instance := range stage.LinkAssociationShapes {
		_copy := instance.GongCopy().(*LinkAssociationShape)
		stage.LinkAssociationShapes_reference[instance] = _copy
		stage.LinkAssociationShapes_instance[_copy] = instance
		stage.LinkAssociationShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Parts_reference = make(map[*Part]*Part)
	stage.Parts_referenceOrder = make(map[*Part]uint) // diff Unstage needs the reference order
	stage.Parts_instance = make(map[*Part]*Part)
	for instance := range stage.Parts {
		_copy := instance.GongCopy().(*Part)
		stage.Parts_reference[instance] = _copy
		stage.Parts_instance[_copy] = instance
		stage.Parts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PartShapes_reference = make(map[*PartShape]*PartShape)
	stage.PartShapes_referenceOrder = make(map[*PartShape]uint) // diff Unstage needs the reference order
	stage.PartShapes_instance = make(map[*PartShape]*PartShape)
	for instance := range stage.PartShapes {
		_copy := instance.GongCopy().(*PartShape)
		stage.PartShapes_reference[instance] = _copy
		stage.PartShapes_instance[_copy] = instance
		stage.PartShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Systems_reference = make(map[*System]*System)
	stage.Systems_referenceOrder = make(map[*System]uint) // diff Unstage needs the reference order
	stage.Systems_instance = make(map[*System]*System)
	for instance := range stage.Systems {
		_copy := instance.GongCopy().(*System)
		stage.Systems_reference[instance] = _copy
		stage.Systems_instance[_copy] = instance
		stage.Systems_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SystemShapes_reference = make(map[*SystemShape]*SystemShape)
	stage.SystemShapes_referenceOrder = make(map[*SystemShape]uint) // diff Unstage needs the reference order
	stage.SystemShapes_instance = make(map[*SystemShape]*SystemShape)
	for instance := range stage.SystemShapes {
		_copy := instance.GongCopy().(*SystemShape)
		stage.SystemShapes_reference[instance] = _copy
		stage.SystemShapes_instance[_copy] = instance
		stage.SystemShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.DiagramStructures {
		reference := stage.DiagramStructures_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Librarys {
		reference := stage.Librarys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Links {
		reference := stage.Links_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.LinkAssociationShapes {
		reference := stage.LinkAssociationShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Parts {
		reference := stage.Parts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PartShapes {
		reference := stage.PartShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Systems {
		reference := stage.Systems_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SystemShapes {
		reference := stage.SystemShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (diagramstructure *DiagramStructure) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DiagramStructure_stagedOrder[diagramstructure]; ok {
		return order
	}
	if order, ok := stage.DiagramStructures_referenceOrder[diagramstructure]; ok {
		return order
	} else {
		log.Printf("instance %p of type DiagramStructure was not staged and does not have a reference order", diagramstructure)
		return 0
	}
}

func (library *Library) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Library_stagedOrder[library]; ok {
		return order
	}
	if order, ok := stage.Librarys_referenceOrder[library]; ok {
		return order
	} else {
		log.Printf("instance %p of type Library was not staged and does not have a reference order", library)
		return 0
	}
}

func (link *Link) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Link_stagedOrder[link]; ok {
		return order
	}
	if order, ok := stage.Links_referenceOrder[link]; ok {
		return order
	} else {
		log.Printf("instance %p of type Link was not staged and does not have a reference order", link)
		return 0
	}
}

func (linkassociationshape *LinkAssociationShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LinkAssociationShape_stagedOrder[linkassociationshape]; ok {
		return order
	}
	if order, ok := stage.LinkAssociationShapes_referenceOrder[linkassociationshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type LinkAssociationShape was not staged and does not have a reference order", linkassociationshape)
		return 0
	}
}

func (part *Part) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Part_stagedOrder[part]; ok {
		return order
	}
	if order, ok := stage.Parts_referenceOrder[part]; ok {
		return order
	} else {
		log.Printf("instance %p of type Part was not staged and does not have a reference order", part)
		return 0
	}
}

func (partshape *PartShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PartShape_stagedOrder[partshape]; ok {
		return order
	}
	if order, ok := stage.PartShapes_referenceOrder[partshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type PartShape was not staged and does not have a reference order", partshape)
		return 0
	}
}

func (system *System) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.System_stagedOrder[system]; ok {
		return order
	}
	if order, ok := stage.Systems_referenceOrder[system]; ok {
		return order
	} else {
		log.Printf("instance %p of type System was not staged and does not have a reference order", system)
		return 0
	}
}

func (systemshape *SystemShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SystemShape_stagedOrder[systemshape]; ok {
		return order
	}
	if order, ok := stage.SystemShapes_referenceOrder[systemshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type SystemShape was not staged and does not have a reference order", systemshape)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (diagramstructure *DiagramStructure) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramstructure.GongGetGongstructName(), diagramstructure.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagramstructure *DiagramStructure) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramstructure.GongGetGongstructName(), diagramstructure.GongGetOrder(stage))
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

func (link *Link) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", link.GongGetGongstructName(), link.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (link *Link) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", link.GongGetGongstructName(), link.GongGetOrder(stage))
}

func (linkassociationshape *LinkAssociationShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", linkassociationshape.GongGetGongstructName(), linkassociationshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (linkassociationshape *LinkAssociationShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", linkassociationshape.GongGetGongstructName(), linkassociationshape.GongGetOrder(stage))
}

func (part *Part) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part.GongGetGongstructName(), part.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part *Part) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part.GongGetGongstructName(), part.GongGetOrder(stage))
}

func (partshape *PartShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partshape.GongGetGongstructName(), partshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (partshape *PartShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", partshape.GongGetGongstructName(), partshape.GongGetOrder(stage))
}

func (system *System) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", system.GongGetGongstructName(), system.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (system *System) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", system.GongGetGongstructName(), system.GongGetOrder(stage))
}

func (systemshape *SystemShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", systemshape.GongGetGongstructName(), systemshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (systemshape *SystemShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", systemshape.GongGetGongstructName(), systemshape.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (diagramstructure *DiagramStructure) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramstructure.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DiagramStructure")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagramstructure.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(library.Name))
	return
}

func (link *Link) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", link.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Link")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(link.Name))
	return
}

func (linkassociationshape *LinkAssociationShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", linkassociationshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LinkAssociationShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(linkassociationshape.Name))
	return
}

func (part *Part) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Part")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(part.Name))
	return
}

func (partshape *PartShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PartShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(partshape.Name))
	return
}

func (system *System) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", system.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "System")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(system.Name))
	return
}

func (systemshape *SystemShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", systemshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SystemShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(systemshape.Name))
	return
}

// insertion point for unstaging
func (diagramstructure *DiagramStructure) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramstructure.GongGetReferenceIdentifier(stage))
	return
}

func (library *Library) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetReferenceIdentifier(stage))
	return
}

func (link *Link) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", link.GongGetReferenceIdentifier(stage))
	return
}

func (linkassociationshape *LinkAssociationShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", linkassociationshape.GongGetReferenceIdentifier(stage))
	return
}

func (part *Part) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part.GongGetReferenceIdentifier(stage))
	return
}

func (partshape *PartShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", partshape.GongGetReferenceIdentifier(stage))
	return
}

func (system *System) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", system.GongGetReferenceIdentifier(stage))
	return
}

func (systemshape *SystemShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", systemshape.GongGetReferenceIdentifier(stage))
	return
}

func IntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += IntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// end of template
