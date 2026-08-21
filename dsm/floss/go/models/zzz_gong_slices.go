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
	// Compute reverse map for named struct DiagramFloss
	// insertion point per field
	stage.DiagramFloss_System_Shapes_reverseMap = make(map[*SystemShape]*DiagramFloss)
	for diagramfloss := range stage.DiagramFlosss {
		_ = diagramfloss
		for _, _systemshape := range diagramfloss.System_Shapes {
			stage.DiagramFloss_System_Shapes_reverseMap[_systemshape] = diagramfloss
		}
	}
	stage.DiagramFloss_SystemsWhoseNodeIsExpanded_reverseMap = make(map[*System]*DiagramFloss)
	for diagramfloss := range stage.DiagramFlosss {
		_ = diagramfloss
		for _, _system := range diagramfloss.SystemsWhoseNodeIsExpanded {
			stage.DiagramFloss_SystemsWhoseNodeIsExpanded_reverseMap[_system] = diagramfloss
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
	stage.Library_RootSystemes_reverseMap = make(map[*System]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _system := range library.RootSystemes {
			stage.Library_RootSystemes_reverseMap[_system] = library
		}
	}
	stage.Library_SystemsWhoseNodeIsExpanded_reverseMap = make(map[*System]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _system := range library.SystemsWhoseNodeIsExpanded {
			stage.Library_SystemsWhoseNodeIsExpanded_reverseMap[_system] = library
		}
	}

	// Compute reverse map for named struct System
	// insertion point per field
	stage.System_DiagramFlosses_reverseMap = make(map[*DiagramFloss]*System)
	for system := range stage.Systems {
		_ = system
		for _, _diagramfloss := range system.DiagramFlosses {
			stage.System_DiagramFlosses_reverseMap[_diagramfloss] = system
		}
	}
	stage.System_DiagramFlossWhoseNodeIsExpanded_reverseMap = make(map[*DiagramFloss]*System)
	for system := range stage.Systems {
		_ = system
		for _, _diagramfloss := range system.DiagramFlossWhoseNodeIsExpanded {
			stage.System_DiagramFlossWhoseNodeIsExpanded_reverseMap[_diagramfloss] = system
		}
	}
	stage.System_SubSystemes_reverseMap = make(map[*System]*System)
	for system := range stage.Systems {
		_ = system
		for _, _system := range system.SubSystemes {
			stage.System_SubSystemes_reverseMap[_system] = system
		}
	}

	// Compute reverse map for named struct SystemShape
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.DiagramFlosss {
		res = append(res, instance)
	}

	for instance := range stage.Librarys {
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
func (diagramfloss *DiagramFloss) GongCopy() GongstructIF {
	newInstance := new(DiagramFloss)
	diagramfloss.CopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.CopyBasicFields(newInstance)
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
func (diagramfloss *DiagramFloss) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagramfloss).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(diagramfloss), uint64(GetOrderPointerGongstruct(stage, diagramfloss)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(library), uint64(GetOrderPointerGongstruct(stage, library)))
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
	var diagramflosss_newInstances []*DiagramFloss
	var diagramflosss_deletedInstances []*DiagramFloss

	// parse all staged instances and check if they have a reference
	for diagramfloss := range stage.DiagramFlosss {
		if ref, ok := stage.DiagramFlosss_reference[diagramfloss]; !ok {
			diagramflosss_newInstances = append(diagramflosss_newInstances, diagramfloss)
			newInstancesSlice = append(newInstancesSlice, diagramfloss.GongMarshallIdentifier(stage))
			if stage.DiagramFlosss_referenceOrder == nil {
				stage.DiagramFlosss_referenceOrder = make(map[*DiagramFloss]uint)
			}
			stage.DiagramFlosss_referenceOrder[diagramfloss] = stage.DiagramFloss_stagedOrder[diagramfloss]
			newInstancesReverseSlice = append(newInstancesReverseSlice, diagramfloss.GongMarshallUnstaging(stage))
			// delete(stage.DiagramFlosss_referenceOrder, diagramfloss)
			fieldInitializers, pointersInitializations := diagramfloss.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DiagramFloss_stagedOrder[ref] = stage.DiagramFloss_stagedOrder[diagramfloss]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := diagramfloss.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, diagramfloss)
			// delete(stage.DiagramFloss_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if diagramfloss.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", diagramfloss.GetName())
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
	for _, ref := range stage.DiagramFlosss_reference {
		instance := stage.DiagramFlosss_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.DiagramFlosss[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			diagramflosss_deletedInstances = append(diagramflosss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(diagramflosss_newInstances)
	lenDeletedInstances += len(diagramflosss_deletedInstances)
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
	stage.DiagramFlosss_reference = make(map[*DiagramFloss]*DiagramFloss)
	stage.DiagramFlosss_referenceOrder = make(map[*DiagramFloss]uint) // diff Unstage needs the reference order
	stage.DiagramFlosss_instance = make(map[*DiagramFloss]*DiagramFloss)
	for instance := range stage.DiagramFlosss {
		_copy := instance.GongCopy().(*DiagramFloss)
		stage.DiagramFlosss_reference[instance] = _copy
		stage.DiagramFlosss_instance[_copy] = instance
		stage.DiagramFlosss_referenceOrder[_copy] = instance.GongGetOrder(stage)
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
	for instance := range stage.DiagramFlosss {
		reference := stage.DiagramFlosss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Librarys {
		reference := stage.Librarys_reference[instance]
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
func (diagramfloss *DiagramFloss) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DiagramFloss_stagedOrder[diagramfloss]; ok {
		return order
	}
	if order, ok := stage.DiagramFlosss_referenceOrder[diagramfloss]; ok {
		return order
	} else {
		log.Printf("instance %p of type DiagramFloss was not staged and does not have a reference order", diagramfloss)
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
func (diagramfloss *DiagramFloss) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramfloss.GongGetGongstructName(), diagramfloss.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagramfloss *DiagramFloss) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagramfloss.GongGetGongstructName(), diagramfloss.GongGetOrder(stage))
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
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
func (diagramfloss *DiagramFloss) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramfloss.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DiagramFloss")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(diagramfloss.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(library.Name))
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
func (diagramfloss *DiagramFloss) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagramfloss.GongGetReferenceIdentifier(stage))
	return
}

func (library *Library) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetReferenceIdentifier(stage))
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
