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

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.AttributeShapes)

	res = __gong__appendInstances(res, stage.Classdiagrams)

	res = __gong__appendInstances(res, stage.DiagramPackages)

	res = __gong__appendInstances(res, stage.GongEnumShapes)

	res = __gong__appendInstances(res, stage.GongEnumValueShapes)

	res = __gong__appendInstances(res, stage.GongNoteLinkShapes)

	res = __gong__appendInstances(res, stage.GongNoteShapes)

	res = __gong__appendInstances(res, stage.GongStructShapes)

	res = __gong__appendInstances(res, stage.LinkShapes)

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
func (attributeshape *AttributeShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attributeshape)
}

func (classdiagram *Classdiagram) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, classdiagram)
}

func (diagrampackage *DiagramPackage) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, diagrampackage)
}

func (gongenumshape *GongEnumShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, gongenumshape)
}

func (gongenumvalueshape *GongEnumValueShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, gongenumvalueshape)
}

func (gongnotelinkshape *GongNoteLinkShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, gongnotelinkshape)
}

func (gongnoteshape *GongNoteShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, gongnoteshape)
}

func (gongstructshape *GongStructShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, gongstructshape)
}

func (linkshape *LinkShape) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, linkshape)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
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
	__gong__computeReferencePass1(stage, stage.AttributeShapes, &stage.AttributeShapes_reference, &stage.AttributeShapes_referenceOrder, &stage.AttributeShapes_instance)

	__gong__computeReferencePass1(stage, stage.Classdiagrams, &stage.Classdiagrams_reference, &stage.Classdiagrams_referenceOrder, &stage.Classdiagrams_instance)

	__gong__computeReferencePass1(stage, stage.DiagramPackages, &stage.DiagramPackages_reference, &stage.DiagramPackages_referenceOrder, &stage.DiagramPackages_instance)

	__gong__computeReferencePass1(stage, stage.GongEnumShapes, &stage.GongEnumShapes_reference, &stage.GongEnumShapes_referenceOrder, &stage.GongEnumShapes_instance)

	__gong__computeReferencePass1(stage, stage.GongEnumValueShapes, &stage.GongEnumValueShapes_reference, &stage.GongEnumValueShapes_referenceOrder, &stage.GongEnumValueShapes_instance)

	__gong__computeReferencePass1(stage, stage.GongNoteLinkShapes, &stage.GongNoteLinkShapes_reference, &stage.GongNoteLinkShapes_referenceOrder, &stage.GongNoteLinkShapes_instance)

	__gong__computeReferencePass1(stage, stage.GongNoteShapes, &stage.GongNoteShapes_reference, &stage.GongNoteShapes_referenceOrder, &stage.GongNoteShapes_instance)

	__gong__computeReferencePass1(stage, stage.GongStructShapes, &stage.GongStructShapes_reference, &stage.GongStructShapes_referenceOrder, &stage.GongStructShapes_instance)

	__gong__computeReferencePass1(stage, stage.LinkShapes, &stage.LinkShapes_reference, &stage.LinkShapes_referenceOrder, &stage.LinkShapes_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.AttributeShapes, stage.AttributeShapes_reference, stage)

	__gong__computeReferencePass2(stage.Classdiagrams, stage.Classdiagrams_reference, stage)

	__gong__computeReferencePass2(stage.DiagramPackages, stage.DiagramPackages_reference, stage)

	__gong__computeReferencePass2(stage.GongEnumShapes, stage.GongEnumShapes_reference, stage)

	__gong__computeReferencePass2(stage.GongEnumValueShapes, stage.GongEnumValueShapes_reference, stage)

	__gong__computeReferencePass2(stage.GongNoteLinkShapes, stage.GongNoteLinkShapes_reference, stage)

	__gong__computeReferencePass2(stage.GongNoteShapes, stage.GongNoteShapes_reference, stage)

	__gong__computeReferencePass2(stage.GongStructShapes, stage.GongStructShapes_reference, stage)

	__gong__computeReferencePass2(stage.LinkShapes, stage.LinkShapes_reference, stage)

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
	return __gong__getOrder(stage.AttributeShape_stagedOrder, stage.AttributeShapes_referenceOrder, attributeshape, "AttributeShape")
}

func (classdiagram *Classdiagram) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Classdiagram_stagedOrder, stage.Classdiagrams_referenceOrder, classdiagram, "Classdiagram")
}

func (diagrampackage *DiagramPackage) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DiagramPackage_stagedOrder, stage.DiagramPackages_referenceOrder, diagrampackage, "DiagramPackage")
}

func (gongenumshape *GongEnumShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GongEnumShape_stagedOrder, stage.GongEnumShapes_referenceOrder, gongenumshape, "GongEnumShape")
}

func (gongenumvalueshape *GongEnumValueShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GongEnumValueShape_stagedOrder, stage.GongEnumValueShapes_referenceOrder, gongenumvalueshape, "GongEnumValueShape")
}

func (gongnotelinkshape *GongNoteLinkShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GongNoteLinkShape_stagedOrder, stage.GongNoteLinkShapes_referenceOrder, gongnotelinkshape, "GongNoteLinkShape")
}

func (gongnoteshape *GongNoteShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GongNoteShape_stagedOrder, stage.GongNoteShapes_referenceOrder, gongnoteshape, "GongNoteShape")
}

func (gongstructshape *GongStructShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.GongStructShape_stagedOrder, stage.GongStructShapes_referenceOrder, gongstructshape, "GongStructShape")
}

func (linkshape *LinkShape) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.LinkShape_stagedOrder, stage.LinkShapes_referenceOrder, linkshape, "LinkShape")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (attributeshape *AttributeShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attributeshape, attributeshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attributeshape *AttributeShape) GongGetReferenceIdentifier(stage *Stage) string {
	return attributeshape.GongGetIdentifier(stage)
}

func (classdiagram *Classdiagram) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(classdiagram, classdiagram.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (classdiagram *Classdiagram) GongGetReferenceIdentifier(stage *Stage) string {
	return classdiagram.GongGetIdentifier(stage)
}

func (diagrampackage *DiagramPackage) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(diagrampackage, diagrampackage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (diagrampackage *DiagramPackage) GongGetReferenceIdentifier(stage *Stage) string {
	return diagrampackage.GongGetIdentifier(stage)
}

func (gongenumshape *GongEnumShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(gongenumshape, gongenumshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongenumshape *GongEnumShape) GongGetReferenceIdentifier(stage *Stage) string {
	return gongenumshape.GongGetIdentifier(stage)
}

func (gongenumvalueshape *GongEnumValueShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(gongenumvalueshape, gongenumvalueshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongenumvalueshape *GongEnumValueShape) GongGetReferenceIdentifier(stage *Stage) string {
	return gongenumvalueshape.GongGetIdentifier(stage)
}

func (gongnotelinkshape *GongNoteLinkShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(gongnotelinkshape, gongnotelinkshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongnotelinkshape *GongNoteLinkShape) GongGetReferenceIdentifier(stage *Stage) string {
	return gongnotelinkshape.GongGetIdentifier(stage)
}

func (gongnoteshape *GongNoteShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(gongnoteshape, gongnoteshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongnoteshape *GongNoteShape) GongGetReferenceIdentifier(stage *Stage) string {
	return gongnoteshape.GongGetIdentifier(stage)
}

func (gongstructshape *GongStructShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(gongstructshape, gongstructshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gongstructshape *GongStructShape) GongGetReferenceIdentifier(stage *Stage) string {
	return gongstructshape.GongGetIdentifier(stage)
}

func (linkshape *LinkShape) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(linkshape, linkshape.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (linkshape *LinkShape) GongGetReferenceIdentifier(stage *Stage) string {
	return linkshape.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (attributeshape *AttributeShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attributeshape.GongGetIdentifier(stage), "AttributeShape", attributeshape.Name)
}

func (classdiagram *Classdiagram) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(classdiagram.GongGetIdentifier(stage), "Classdiagram", classdiagram.Name)
}

func (diagrampackage *DiagramPackage) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(diagrampackage.GongGetIdentifier(stage), "DiagramPackage", diagrampackage.Name)
}

func (gongenumshape *GongEnumShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(gongenumshape.GongGetIdentifier(stage), "GongEnumShape", gongenumshape.Name)
}

func (gongenumvalueshape *GongEnumValueShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(gongenumvalueshape.GongGetIdentifier(stage), "GongEnumValueShape", gongenumvalueshape.Name)
}

func (gongnotelinkshape *GongNoteLinkShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(gongnotelinkshape.GongGetIdentifier(stage), "GongNoteLinkShape", gongnotelinkshape.Name)
}

func (gongnoteshape *GongNoteShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(gongnoteshape.GongGetIdentifier(stage), "GongNoteShape", gongnoteshape.Name)
}

func (gongstructshape *GongStructShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(gongstructshape.GongGetIdentifier(stage), "GongStructShape", gongstructshape.Name)
}

func (linkshape *LinkShape) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(linkshape.GongGetIdentifier(stage), "LinkShape", linkshape.Name)
}

// insertion point for unstaging
func (attributeshape *AttributeShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attributeshape.GongGetReferenceIdentifier(stage))
}

func (classdiagram *Classdiagram) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(classdiagram.GongGetReferenceIdentifier(stage))
}

func (diagrampackage *DiagramPackage) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(diagrampackage.GongGetReferenceIdentifier(stage))
}

func (gongenumshape *GongEnumShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(gongenumshape.GongGetReferenceIdentifier(stage))
}

func (gongenumvalueshape *GongEnumValueShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(gongenumvalueshape.GongGetReferenceIdentifier(stage))
}

func (gongnotelinkshape *GongNoteLinkShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(gongnotelinkshape.GongGetReferenceIdentifier(stage))
}

func (gongnoteshape *GongNoteShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(gongnoteshape.GongGetReferenceIdentifier(stage))
}

func (gongstructshape *GongStructShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(gongstructshape.GongGetReferenceIdentifier(stage))
}

func (linkshape *LinkShape) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(linkshape.GongGetReferenceIdentifier(stage))
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

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
