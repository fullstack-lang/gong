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
	// Compute reverse map for named struct AxesShape
	// insertion point per field

	// Compute reverse map for named struct Library
	// insertion point per field
	stage.Library_SubLibraries_reverseMap = make(map[*Library]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _library := range library.SubLibraries {
			stage.Library_SubLibraries_reverseMap[_library] = library
		}
	}
	stage.Library_Plants_reverseMap = make(map[*Plant]*Library)
	for library := range stage.Librarys {
		_ = library
		for _, _plant := range library.Plants {
			stage.Library_Plants_reverseMap[_plant] = library
		}
	}

	// Compute reverse map for named struct Plant
	// insertion point per field
	stage.Plant_PlantDiagramsWhoseNodeIsExpanded_reverseMap = make(map[*PlantDiagram]*Plant)
	for plant := range stage.Plants {
		_ = plant
		for _, _plantdiagram := range plant.PlantDiagramsWhoseNodeIsExpanded {
			stage.Plant_PlantDiagramsWhoseNodeIsExpanded_reverseMap[_plantdiagram] = plant
		}
	}
	stage.Plant_PlantDiagrams_reverseMap = make(map[*PlantDiagram]*Plant)
	for plant := range stage.Plants {
		_ = plant
		for _, _plantdiagram := range plant.PlantDiagrams {
			stage.Plant_PlantDiagrams_reverseMap[_plantdiagram] = plant
		}
	}

	// Compute reverse map for named struct PlantDiagram
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.AxesShapes {
		res = append(res, instance)
	}

	for instance := range stage.Librarys {
		res = append(res, instance)
	}

	for instance := range stage.Plants {
		res = append(res, instance)
	}

	for instance := range stage.PlantDiagrams {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (axesshape *AxesShape) GongCopy() GongstructIF {
	newInstance := new(AxesShape)
	axesshape.CopyBasicFields(newInstance)
	return newInstance
}

func (library *Library) GongCopy() GongstructIF {
	newInstance := new(Library)
	library.CopyBasicFields(newInstance)
	return newInstance
}

func (plant *Plant) GongCopy() GongstructIF {
	newInstance := new(Plant)
	plant.CopyBasicFields(newInstance)
	return newInstance
}

func (plantdiagram *PlantDiagram) GongCopy() GongstructIF {
	newInstance := new(PlantDiagram)
	plantdiagram.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (axesshape *AxesShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(axesshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(axesshape), uint64(GetOrderPointerGongstruct(stage, axesshape)))
	return
}

func (library *Library) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(library).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(library), uint64(GetOrderPointerGongstruct(stage, library)))
	return
}

func (plant *Plant) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(plant).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(plant), uint64(GetOrderPointerGongstruct(stage, plant)))
	return
}

func (plantdiagram *PlantDiagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(plantdiagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(plantdiagram), uint64(GetOrderPointerGongstruct(stage, plantdiagram)))
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
	var axesshapes_newInstances []*AxesShape
	var axesshapes_deletedInstances []*AxesShape

	// parse all staged instances and check if they have a reference
	for axesshape := range stage.AxesShapes {
		if ref, ok := stage.AxesShapes_reference[axesshape]; !ok {
			axesshapes_newInstances = append(axesshapes_newInstances, axesshape)
			newInstancesSlice = append(newInstancesSlice, axesshape.GongMarshallIdentifier(stage))
			if stage.AxesShapes_referenceOrder == nil {
				stage.AxesShapes_referenceOrder = make(map[*AxesShape]uint)
			}
			stage.AxesShapes_referenceOrder[axesshape] = stage.AxesShape_stagedOrder[axesshape]
			newInstancesReverseSlice = append(newInstancesReverseSlice, axesshape.GongMarshallUnstaging(stage))
			// delete(stage.AxesShapes_referenceOrder, axesshape)
			fieldInitializers, pointersInitializations := axesshape.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AxesShape_stagedOrder[ref] = stage.AxesShape_stagedOrder[axesshape]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := axesshape.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, axesshape)
			// delete(stage.AxesShape_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if axesshape.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", axesshape.GetName())
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
	for _, ref := range stage.AxesShapes_reference {
		instance := stage.AxesShapes_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.AxesShapes[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			axesshapes_deletedInstances = append(axesshapes_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(axesshapes_newInstances)
	lenDeletedInstances += len(axesshapes_deletedInstances)
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
	var plants_newInstances []*Plant
	var plants_deletedInstances []*Plant

	// parse all staged instances and check if they have a reference
	for plant := range stage.Plants {
		if ref, ok := stage.Plants_reference[plant]; !ok {
			plants_newInstances = append(plants_newInstances, plant)
			newInstancesSlice = append(newInstancesSlice, plant.GongMarshallIdentifier(stage))
			if stage.Plants_referenceOrder == nil {
				stage.Plants_referenceOrder = make(map[*Plant]uint)
			}
			stage.Plants_referenceOrder[plant] = stage.Plant_stagedOrder[plant]
			newInstancesReverseSlice = append(newInstancesReverseSlice, plant.GongMarshallUnstaging(stage))
			// delete(stage.Plants_referenceOrder, plant)
			fieldInitializers, pointersInitializations := plant.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Plant_stagedOrder[ref] = stage.Plant_stagedOrder[plant]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := plant.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, plant)
			// delete(stage.Plant_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if plant.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", plant.GetName())
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
	for _, ref := range stage.Plants_reference {
		instance := stage.Plants_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Plants[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			plants_deletedInstances = append(plants_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(plants_newInstances)
	lenDeletedInstances += len(plants_deletedInstances)
	var plantdiagrams_newInstances []*PlantDiagram
	var plantdiagrams_deletedInstances []*PlantDiagram

	// parse all staged instances and check if they have a reference
	for plantdiagram := range stage.PlantDiagrams {
		if ref, ok := stage.PlantDiagrams_reference[plantdiagram]; !ok {
			plantdiagrams_newInstances = append(plantdiagrams_newInstances, plantdiagram)
			newInstancesSlice = append(newInstancesSlice, plantdiagram.GongMarshallIdentifier(stage))
			if stage.PlantDiagrams_referenceOrder == nil {
				stage.PlantDiagrams_referenceOrder = make(map[*PlantDiagram]uint)
			}
			stage.PlantDiagrams_referenceOrder[plantdiagram] = stage.PlantDiagram_stagedOrder[plantdiagram]
			newInstancesReverseSlice = append(newInstancesReverseSlice, plantdiagram.GongMarshallUnstaging(stage))
			// delete(stage.PlantDiagrams_referenceOrder, plantdiagram)
			fieldInitializers, pointersInitializations := plantdiagram.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.PlantDiagram_stagedOrder[ref] = stage.PlantDiagram_stagedOrder[plantdiagram]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := plantdiagram.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, plantdiagram)
			// delete(stage.PlantDiagram_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if plantdiagram.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", plantdiagram.GetName())
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
	for _, ref := range stage.PlantDiagrams_reference {
		instance := stage.PlantDiagrams_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.PlantDiagrams[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			plantdiagrams_deletedInstances = append(plantdiagrams_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(plantdiagrams_newInstances)
	lenDeletedInstances += len(plantdiagrams_deletedInstances)

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
	stage.AxesShapes_reference = make(map[*AxesShape]*AxesShape)
	stage.AxesShapes_referenceOrder = make(map[*AxesShape]uint) // diff Unstage needs the reference order
	stage.AxesShapes_instance = make(map[*AxesShape]*AxesShape)
	for instance := range stage.AxesShapes {
		_copy := instance.GongCopy().(*AxesShape)
		stage.AxesShapes_reference[instance] = _copy
		stage.AxesShapes_instance[_copy] = instance
		stage.AxesShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
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

	stage.Plants_reference = make(map[*Plant]*Plant)
	stage.Plants_referenceOrder = make(map[*Plant]uint) // diff Unstage needs the reference order
	stage.Plants_instance = make(map[*Plant]*Plant)
	for instance := range stage.Plants {
		_copy := instance.GongCopy().(*Plant)
		stage.Plants_reference[instance] = _copy
		stage.Plants_instance[_copy] = instance
		stage.Plants_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.PlantDiagrams_reference = make(map[*PlantDiagram]*PlantDiagram)
	stage.PlantDiagrams_referenceOrder = make(map[*PlantDiagram]uint) // diff Unstage needs the reference order
	stage.PlantDiagrams_instance = make(map[*PlantDiagram]*PlantDiagram)
	for instance := range stage.PlantDiagrams {
		_copy := instance.GongCopy().(*PlantDiagram)
		stage.PlantDiagrams_reference[instance] = _copy
		stage.PlantDiagrams_instance[_copy] = instance
		stage.PlantDiagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.AxesShapes {
		reference := stage.AxesShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Librarys {
		reference := stage.Librarys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Plants {
		reference := stage.Plants_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.PlantDiagrams {
		reference := stage.PlantDiagrams_reference[instance]
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
func (axesshape *AxesShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AxesShape_stagedOrder[axesshape]; ok {
		return order
	}
	if order, ok := stage.AxesShapes_referenceOrder[axesshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type AxesShape was not staged and does not have a reference order", axesshape)
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

func (plant *Plant) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Plant_stagedOrder[plant]; ok {
		return order
	}
	if order, ok := stage.Plants_referenceOrder[plant]; ok {
		return order
	} else {
		log.Printf("instance %p of type Plant was not staged and does not have a reference order", plant)
		return 0
	}
}

func (plantdiagram *PlantDiagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.PlantDiagram_stagedOrder[plantdiagram]; ok {
		return order
	}
	if order, ok := stage.PlantDiagrams_referenceOrder[plantdiagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type PlantDiagram was not staged and does not have a reference order", plantdiagram)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (axesshape *AxesShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", axesshape.GongGetGongstructName(), axesshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (axesshape *AxesShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", axesshape.GongGetGongstructName(), axesshape.GongGetOrder(stage))
}

func (library *Library) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (library *Library) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", library.GongGetGongstructName(), library.GongGetOrder(stage))
}

func (plant *Plant) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plant.GongGetGongstructName(), plant.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (plant *Plant) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plant.GongGetGongstructName(), plant.GongGetOrder(stage))
}

func (plantdiagram *PlantDiagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plantdiagram.GongGetGongstructName(), plantdiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (plantdiagram *PlantDiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", plantdiagram.GongGetGongstructName(), plantdiagram.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (axesshape *AxesShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", axesshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AxesShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(axesshape.Name))
	return
}

func (library *Library) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Library")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(library.Name))
	return
}

func (plant *Plant) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plant.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Plant")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(plant.Name))
	return
}

func (plantdiagram *PlantDiagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plantdiagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "PlantDiagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(plantdiagram.Name))
	return
}

// insertion point for unstaging
func (axesshape *AxesShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", axesshape.GongGetReferenceIdentifier(stage))
	return
}

func (library *Library) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", library.GongGetReferenceIdentifier(stage))
	return
}

func (plant *Plant) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plant.GongGetReferenceIdentifier(stage))
	return
}

func (plantdiagram *PlantDiagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", plantdiagram.GongGetReferenceIdentifier(stage))
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
