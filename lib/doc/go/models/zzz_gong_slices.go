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

	// end of insertion point per named struct
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
	newInstance := new(AttributeShape)
	attributeshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (classdiagram *Classdiagram) GongCopy() GongstructIF {
	newInstance := new(Classdiagram)
	classdiagram.GongCopyBasicFields(newInstance)
	return newInstance
}

func (diagrampackage *DiagramPackage) GongCopy() GongstructIF {
	newInstance := new(DiagramPackage)
	diagrampackage.GongCopyBasicFields(newInstance)
	return newInstance
}

func (gongenumshape *GongEnumShape) GongCopy() GongstructIF {
	newInstance := new(GongEnumShape)
	gongenumshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (gongenumvalueshape *GongEnumValueShape) GongCopy() GongstructIF {
	newInstance := new(GongEnumValueShape)
	gongenumvalueshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (gongnotelinkshape *GongNoteLinkShape) GongCopy() GongstructIF {
	newInstance := new(GongNoteLinkShape)
	gongnotelinkshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (gongnoteshape *GongNoteShape) GongCopy() GongstructIF {
	newInstance := new(GongNoteShape)
	gongnoteshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (gongstructshape *GongStructShape) GongCopy() GongstructIF {
	newInstance := new(GongStructShape)
	gongstructshape.GongCopyBasicFields(newInstance)
	return newInstance
}

func (linkshape *LinkShape) GongCopy() GongstructIF {
	newInstance := new(LinkShape)
	linkshape.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (attributeshape *AttributeShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attributeshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attributeshape), uint64(stage.GetOrder(attributeshape)))
	return
}

func (classdiagram *Classdiagram) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(classdiagram).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(classdiagram), uint64(stage.GetOrder(classdiagram)))
	return
}

func (diagrampackage *DiagramPackage) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(diagrampackage).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(diagrampackage), uint64(stage.GetOrder(diagrampackage)))
	return
}

func (gongenumshape *GongEnumShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(gongenumshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(gongenumshape), uint64(stage.GetOrder(gongenumshape)))
	return
}

func (gongenumvalueshape *GongEnumValueShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(gongenumvalueshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(gongenumvalueshape), uint64(stage.GetOrder(gongenumvalueshape)))
	return
}

func (gongnotelinkshape *GongNoteLinkShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(gongnotelinkshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(gongnotelinkshape), uint64(stage.GetOrder(gongnotelinkshape)))
	return
}

func (gongnoteshape *GongNoteShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(gongnoteshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(gongnoteshape), uint64(stage.GetOrder(gongnoteshape)))
	return
}

func (gongstructshape *GongStructShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(gongstructshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(gongstructshape), uint64(stage.GetOrder(gongstructshape)))
	return
}

func (linkshape *LinkShape) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(linkshape).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(linkshape), uint64(stage.GetOrder(linkshape)))
	return
}


type GongstructDiffable[T any] interface {
	PointerToGongstruct
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
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
	computeCommitsForType(
		stage,
		stage.AttributeShapes,
		stage.AttributeShape_stagedOrder,
		stage.AttributeShapes_reference,
		&stage.AttributeShapes_referenceOrder,
		stage.AttributeShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Classdiagrams,
		stage.Classdiagram_stagedOrder,
		stage.Classdiagrams_reference,
		&stage.Classdiagrams_referenceOrder,
		stage.Classdiagrams_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DiagramPackages,
		stage.DiagramPackage_stagedOrder,
		stage.DiagramPackages_reference,
		&stage.DiagramPackages_referenceOrder,
		stage.DiagramPackages_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GongEnumShapes,
		stage.GongEnumShape_stagedOrder,
		stage.GongEnumShapes_reference,
		&stage.GongEnumShapes_referenceOrder,
		stage.GongEnumShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GongEnumValueShapes,
		stage.GongEnumValueShape_stagedOrder,
		stage.GongEnumValueShapes_reference,
		&stage.GongEnumValueShapes_referenceOrder,
		stage.GongEnumValueShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GongNoteLinkShapes,
		stage.GongNoteLinkShape_stagedOrder,
		stage.GongNoteLinkShapes_reference,
		&stage.GongNoteLinkShapes_referenceOrder,
		stage.GongNoteLinkShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GongNoteShapes,
		stage.GongNoteShape_stagedOrder,
		stage.GongNoteShapes_reference,
		&stage.GongNoteShapes_referenceOrder,
		stage.GongNoteShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.GongStructShapes,
		stage.GongStructShape_stagedOrder,
		stage.GongStructShapes_reference,
		&stage.GongStructShapes_referenceOrder,
		stage.GongStructShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.LinkShapes,
		stage.LinkShape_stagedOrder,
		stage.LinkShapes_reference,
		&stage.LinkShapes_referenceOrder,
		stage.LinkShapes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

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
	stage.AttributeShapes_reference = make(map[*AttributeShape]*AttributeShape)
	stage.AttributeShapes_referenceOrder = make(map[*AttributeShape]uint) // diff Unstage needs the reference order
	stage.AttributeShapes_instance = make(map[*AttributeShape]*AttributeShape)
	for instance := range stage.AttributeShapes {
		_copy := instance.GongCopy().(*AttributeShape)
		stage.AttributeShapes_reference[instance] = _copy
		stage.AttributeShapes_instance[_copy] = instance
		stage.AttributeShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Classdiagrams_reference = make(map[*Classdiagram]*Classdiagram)
	stage.Classdiagrams_referenceOrder = make(map[*Classdiagram]uint) // diff Unstage needs the reference order
	stage.Classdiagrams_instance = make(map[*Classdiagram]*Classdiagram)
	for instance := range stage.Classdiagrams {
		_copy := instance.GongCopy().(*Classdiagram)
		stage.Classdiagrams_reference[instance] = _copy
		stage.Classdiagrams_instance[_copy] = instance
		stage.Classdiagrams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DiagramPackages_reference = make(map[*DiagramPackage]*DiagramPackage)
	stage.DiagramPackages_referenceOrder = make(map[*DiagramPackage]uint) // diff Unstage needs the reference order
	stage.DiagramPackages_instance = make(map[*DiagramPackage]*DiagramPackage)
	for instance := range stage.DiagramPackages {
		_copy := instance.GongCopy().(*DiagramPackage)
		stage.DiagramPackages_reference[instance] = _copy
		stage.DiagramPackages_instance[_copy] = instance
		stage.DiagramPackages_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GongEnumShapes_reference = make(map[*GongEnumShape]*GongEnumShape)
	stage.GongEnumShapes_referenceOrder = make(map[*GongEnumShape]uint) // diff Unstage needs the reference order
	stage.GongEnumShapes_instance = make(map[*GongEnumShape]*GongEnumShape)
	for instance := range stage.GongEnumShapes {
		_copy := instance.GongCopy().(*GongEnumShape)
		stage.GongEnumShapes_reference[instance] = _copy
		stage.GongEnumShapes_instance[_copy] = instance
		stage.GongEnumShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GongEnumValueShapes_reference = make(map[*GongEnumValueShape]*GongEnumValueShape)
	stage.GongEnumValueShapes_referenceOrder = make(map[*GongEnumValueShape]uint) // diff Unstage needs the reference order
	stage.GongEnumValueShapes_instance = make(map[*GongEnumValueShape]*GongEnumValueShape)
	for instance := range stage.GongEnumValueShapes {
		_copy := instance.GongCopy().(*GongEnumValueShape)
		stage.GongEnumValueShapes_reference[instance] = _copy
		stage.GongEnumValueShapes_instance[_copy] = instance
		stage.GongEnumValueShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GongNoteLinkShapes_reference = make(map[*GongNoteLinkShape]*GongNoteLinkShape)
	stage.GongNoteLinkShapes_referenceOrder = make(map[*GongNoteLinkShape]uint) // diff Unstage needs the reference order
	stage.GongNoteLinkShapes_instance = make(map[*GongNoteLinkShape]*GongNoteLinkShape)
	for instance := range stage.GongNoteLinkShapes {
		_copy := instance.GongCopy().(*GongNoteLinkShape)
		stage.GongNoteLinkShapes_reference[instance] = _copy
		stage.GongNoteLinkShapes_instance[_copy] = instance
		stage.GongNoteLinkShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GongNoteShapes_reference = make(map[*GongNoteShape]*GongNoteShape)
	stage.GongNoteShapes_referenceOrder = make(map[*GongNoteShape]uint) // diff Unstage needs the reference order
	stage.GongNoteShapes_instance = make(map[*GongNoteShape]*GongNoteShape)
	for instance := range stage.GongNoteShapes {
		_copy := instance.GongCopy().(*GongNoteShape)
		stage.GongNoteShapes_reference[instance] = _copy
		stage.GongNoteShapes_instance[_copy] = instance
		stage.GongNoteShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.GongStructShapes_reference = make(map[*GongStructShape]*GongStructShape)
	stage.GongStructShapes_referenceOrder = make(map[*GongStructShape]uint) // diff Unstage needs the reference order
	stage.GongStructShapes_instance = make(map[*GongStructShape]*GongStructShape)
	for instance := range stage.GongStructShapes {
		_copy := instance.GongCopy().(*GongStructShape)
		stage.GongStructShapes_reference[instance] = _copy
		stage.GongStructShapes_instance[_copy] = instance
		stage.GongStructShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.LinkShapes_reference = make(map[*LinkShape]*LinkShape)
	stage.LinkShapes_referenceOrder = make(map[*LinkShape]uint) // diff Unstage needs the reference order
	stage.LinkShapes_instance = make(map[*LinkShape]*LinkShape)
	for instance := range stage.LinkShapes {
		_copy := instance.GongCopy().(*LinkShape)
		stage.LinkShapes_reference[instance] = _copy
		stage.LinkShapes_instance[_copy] = instance
		stage.LinkShapes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.AttributeShapes {
		reference := stage.AttributeShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Classdiagrams {
		reference := stage.Classdiagrams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DiagramPackages {
		reference := stage.DiagramPackages_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GongEnumShapes {
		reference := stage.GongEnumShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GongEnumValueShapes {
		reference := stage.GongEnumValueShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GongNoteLinkShapes {
		reference := stage.GongNoteLinkShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GongNoteShapes {
		reference := stage.GongNoteShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.GongStructShapes {
		reference := stage.GongStructShapes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.LinkShapes {
		reference := stage.LinkShapes_reference[instance]
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
func (attributeshape *AttributeShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AttributeShape_stagedOrder[attributeshape]; ok {
		return order
	}
	if order, ok := stage.AttributeShapes_referenceOrder[attributeshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type AttributeShape was not staged and does not have a reference order", attributeshape)
		return 0
	}
}

func (classdiagram *Classdiagram) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Classdiagram_stagedOrder[classdiagram]; ok {
		return order
	}
	if order, ok := stage.Classdiagrams_referenceOrder[classdiagram]; ok {
		return order
	} else {
		log.Printf("instance %p of type Classdiagram was not staged and does not have a reference order", classdiagram)
		return 0
	}
}

func (diagrampackage *DiagramPackage) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DiagramPackage_stagedOrder[diagrampackage]; ok {
		return order
	}
	if order, ok := stage.DiagramPackages_referenceOrder[diagrampackage]; ok {
		return order
	} else {
		log.Printf("instance %p of type DiagramPackage was not staged and does not have a reference order", diagrampackage)
		return 0
	}
}

func (gongenumshape *GongEnumShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongEnumShape_stagedOrder[gongenumshape]; ok {
		return order
	}
	if order, ok := stage.GongEnumShapes_referenceOrder[gongenumshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type GongEnumShape was not staged and does not have a reference order", gongenumshape)
		return 0
	}
}

func (gongenumvalueshape *GongEnumValueShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongEnumValueShape_stagedOrder[gongenumvalueshape]; ok {
		return order
	}
	if order, ok := stage.GongEnumValueShapes_referenceOrder[gongenumvalueshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type GongEnumValueShape was not staged and does not have a reference order", gongenumvalueshape)
		return 0
	}
}

func (gongnotelinkshape *GongNoteLinkShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongNoteLinkShape_stagedOrder[gongnotelinkshape]; ok {
		return order
	}
	if order, ok := stage.GongNoteLinkShapes_referenceOrder[gongnotelinkshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type GongNoteLinkShape was not staged and does not have a reference order", gongnotelinkshape)
		return 0
	}
}

func (gongnoteshape *GongNoteShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongNoteShape_stagedOrder[gongnoteshape]; ok {
		return order
	}
	if order, ok := stage.GongNoteShapes_referenceOrder[gongnoteshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type GongNoteShape was not staged and does not have a reference order", gongnoteshape)
		return 0
	}
}

func (gongstructshape *GongStructShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.GongStructShape_stagedOrder[gongstructshape]; ok {
		return order
	}
	if order, ok := stage.GongStructShapes_referenceOrder[gongstructshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type GongStructShape was not staged and does not have a reference order", gongstructshape)
		return 0
	}
}

func (linkshape *LinkShape) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.LinkShape_stagedOrder[linkshape]; ok {
		return order
	}
	if order, ok := stage.LinkShapes_referenceOrder[linkshape]; ok {
		return order
	} else {
		log.Printf("instance %p of type LinkShape was not staged and does not have a reference order", linkshape)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (attributeshape *AttributeShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attributeshape.GongGetGongstructName(), attributeshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attributeshape *AttributeShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attributeshape.GongGetGongstructName(), attributeshape.GongGetOrder(stage))
}

func (classdiagram *Classdiagram) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", classdiagram.GongGetGongstructName(), classdiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (classdiagram *Classdiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", classdiagram.GongGetGongstructName(), classdiagram.GongGetOrder(stage))
}

func (diagrampackage *DiagramPackage) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagrampackage.GongGetGongstructName(), diagrampackage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagrampackage *DiagramPackage) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", diagrampackage.GongGetGongstructName(), diagrampackage.GongGetOrder(stage))
}

func (gongenumshape *GongEnumShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongenumshape.GongGetGongstructName(), gongenumshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongenumshape *GongEnumShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongenumshape.GongGetGongstructName(), gongenumshape.GongGetOrder(stage))
}

func (gongenumvalueshape *GongEnumValueShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongenumvalueshape.GongGetGongstructName(), gongenumvalueshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongenumvalueshape *GongEnumValueShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongenumvalueshape.GongGetGongstructName(), gongenumvalueshape.GongGetOrder(stage))
}

func (gongnotelinkshape *GongNoteLinkShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongnotelinkshape.GongGetGongstructName(), gongnotelinkshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongnotelinkshape *GongNoteLinkShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongnotelinkshape.GongGetGongstructName(), gongnotelinkshape.GongGetOrder(stage))
}

func (gongnoteshape *GongNoteShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongnoteshape.GongGetGongstructName(), gongnoteshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongnoteshape *GongNoteShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongnoteshape.GongGetGongstructName(), gongnoteshape.GongGetOrder(stage))
}

func (gongstructshape *GongStructShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongstructshape.GongGetGongstructName(), gongstructshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongstructshape *GongStructShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gongstructshape.GongGetGongstructName(), gongstructshape.GongGetOrder(stage))
}

func (linkshape *LinkShape) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", linkshape.GongGetGongstructName(), linkshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (linkshape *LinkShape) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", linkshape.GongGetGongstructName(), linkshape.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (attributeshape *AttributeShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attributeshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AttributeShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attributeshape.Name))
	return
}

func (classdiagram *Classdiagram) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", classdiagram.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Classdiagram")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(classdiagram.Name))
	return
}

func (diagrampackage *DiagramPackage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagrampackage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DiagramPackage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(diagrampackage.Name))
	return
}

func (gongenumshape *GongEnumShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongenumshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongEnumShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(gongenumshape.Name))
	return
}

func (gongenumvalueshape *GongEnumValueShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongenumvalueshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongEnumValueShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(gongenumvalueshape.Name))
	return
}

func (gongnotelinkshape *GongNoteLinkShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongnotelinkshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongNoteLinkShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(gongnotelinkshape.Name))
	return
}

func (gongnoteshape *GongNoteShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongnoteshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongNoteShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(gongnoteshape.Name))
	return
}

func (gongstructshape *GongStructShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongstructshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "GongStructShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(gongstructshape.Name))
	return
}

func (linkshape *LinkShape) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", linkshape.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "LinkShape")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(linkshape.Name))
	return
}

// insertion point for unstaging
func (attributeshape *AttributeShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attributeshape.GongGetReferenceIdentifier(stage))
	return
}

func (classdiagram *Classdiagram) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", classdiagram.GongGetReferenceIdentifier(stage))
	return
}

func (diagrampackage *DiagramPackage) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", diagrampackage.GongGetReferenceIdentifier(stage))
	return
}

func (gongenumshape *GongEnumShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongenumshape.GongGetReferenceIdentifier(stage))
	return
}

func (gongenumvalueshape *GongEnumValueShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongenumvalueshape.GongGetReferenceIdentifier(stage))
	return
}

func (gongnotelinkshape *GongNoteLinkShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongnotelinkshape.GongGetReferenceIdentifier(stage))
	return
}

func (gongnoteshape *GongNoteShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongnoteshape.GongGetReferenceIdentifier(stage))
	return
}

func (gongstructshape *GongStructShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gongstructshape.GongGetReferenceIdentifier(stage))
	return
}

func (linkshape *LinkShape) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", linkshape.GongGetReferenceIdentifier(stage))
	return
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
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
