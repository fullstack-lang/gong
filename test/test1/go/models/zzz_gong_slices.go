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
	// Compute reverse map for named struct Astruct
	// insertion point per field
	stage.Astruct_Anarrayofb_reverseMap = make(map[*Bstruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _bstruct := range astruct.Anarrayofb {
			stage.Astruct_Anarrayofb_reverseMap[_bstruct] = astruct
		}
	}
	stage.Astruct_Dstruct4s_reverseMap = make(map[*Dstruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _dstruct := range astruct.Dstruct4s {
			stage.Astruct_Dstruct4s_reverseMap[_dstruct] = astruct
		}
	}
	stage.Astruct_Anarrayofa_reverseMap = make(map[*Astruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _astruct := range astruct.Anarrayofa {
			stage.Astruct_Anarrayofa_reverseMap[_astruct] = astruct
		}
	}
	stage.Astruct_Anotherarrayofb_reverseMap = make(map[*Bstruct]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _bstruct := range astruct.Anotherarrayofb {
			stage.Astruct_Anotherarrayofb_reverseMap[_bstruct] = astruct
		}
	}
	stage.Astruct_AnarrayofbUse_reverseMap = make(map[*AstructBstructUse]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _astructbstructuse := range astruct.AnarrayofbUse {
			stage.Astruct_AnarrayofbUse_reverseMap[_astructbstructuse] = astruct
		}
	}
	stage.Astruct_Anarrayofb2Use_reverseMap = make(map[*AstructBstruct2Use]*Astruct)
	for astruct := range stage.Astructs {
		_ = astruct
		for _, _astructbstruct2use := range astruct.Anarrayofb2Use {
			stage.Astruct_Anarrayofb2Use_reverseMap[_astructbstruct2use] = astruct
		}
	}

	// Compute reverse map for named struct Dstruct
	// insertion point per field
	stage.Dstruct_Anarrayofb_reverseMap = make(map[*Bstruct]*Dstruct)
	for dstruct := range stage.Dstructs {
		_ = dstruct
		for _, _bstruct := range dstruct.Anarrayofb {
			stage.Dstruct_Anarrayofb_reverseMap[_bstruct] = dstruct
		}
	}
	stage.Dstruct_Gstructs_reverseMap = make(map[*Gstruct]*Dstruct)
	for dstruct := range stage.Dstructs {
		_ = dstruct
		for _, _gstruct := range dstruct.Gstructs {
			stage.Dstruct_Gstructs_reverseMap[_gstruct] = dstruct
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.Astructs)

	res = __gong__appendInstances(res, stage.AstructBstruct2Uses)

	res = __gong__appendInstances(res, stage.AstructBstructUses)

	res = __gong__appendInstances(res, stage.Bstructs)

	res = __gong__appendInstances(res, stage.Dstructs)

	res = __gong__appendInstances(res, stage.F0123456789012345678901234567890s)

	res = __gong__appendInstances(res, stage.Gstructs)

	return
}

// insertion point per named struct
func (astruct *Astruct) GongCopy() GongstructIF {
	newInstance := new(Astruct)
	astruct.GongCopyBasicFields(newInstance)
	return newInstance
}

func (astructbstruct2use *AstructBstruct2Use) GongCopy() GongstructIF {
	newInstance := new(AstructBstruct2Use)
	astructbstruct2use.GongCopyBasicFields(newInstance)
	return newInstance
}

func (astructbstructuse *AstructBstructUse) GongCopy() GongstructIF {
	newInstance := new(AstructBstructUse)
	astructbstructuse.GongCopyBasicFields(newInstance)
	return newInstance
}

func (bstruct *Bstruct) GongCopy() GongstructIF {
	newInstance := new(Bstruct)
	bstruct.GongCopyBasicFields(newInstance)
	return newInstance
}

func (dstruct *Dstruct) GongCopy() GongstructIF {
	newInstance := new(Dstruct)
	dstruct.GongCopyBasicFields(newInstance)
	return newInstance
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongCopy() GongstructIF {
	newInstance := new(F0123456789012345678901234567890)
	f0123456789012345678901234567890.GongCopyBasicFields(newInstance)
	return newInstance
}

func (gstruct *Gstruct) GongCopy() GongstructIF {
	newInstance := new(Gstruct)
	gstruct.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (astruct *Astruct) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, astruct)
}

func (astructbstruct2use *AstructBstruct2Use) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, astructbstruct2use)
}

func (astructbstructuse *AstructBstructUse) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, astructbstructuse)
}

func (bstruct *Bstruct) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, bstruct)
}

func (dstruct *Dstruct) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, dstruct)
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, f0123456789012345678901234567890)
}

func (gstruct *Gstruct) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, gstruct)
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
		stage.Astructs,
		stage.Astruct_stagedOrder,
		stage.Astructs_reference,
		&stage.Astructs_referenceOrder,
		stage.Astructs_instance,
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
		stage.AstructBstruct2Uses,
		stage.AstructBstruct2Use_stagedOrder,
		stage.AstructBstruct2Uses_reference,
		&stage.AstructBstruct2Uses_referenceOrder,
		stage.AstructBstruct2Uses_instance,
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
		stage.AstructBstructUses,
		stage.AstructBstructUse_stagedOrder,
		stage.AstructBstructUses_reference,
		&stage.AstructBstructUses_referenceOrder,
		stage.AstructBstructUses_instance,
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
		stage.Bstructs,
		stage.Bstruct_stagedOrder,
		stage.Bstructs_reference,
		&stage.Bstructs_referenceOrder,
		stage.Bstructs_instance,
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
		stage.Dstructs,
		stage.Dstruct_stagedOrder,
		stage.Dstructs_reference,
		&stage.Dstructs_referenceOrder,
		stage.Dstructs_instance,
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
		stage.F0123456789012345678901234567890s,
		stage.F0123456789012345678901234567890_stagedOrder,
		stage.F0123456789012345678901234567890s_reference,
		&stage.F0123456789012345678901234567890s_referenceOrder,
		stage.F0123456789012345678901234567890s_instance,
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
		stage.Gstructs,
		stage.Gstruct_stagedOrder,
		stage.Gstructs_reference,
		&stage.Gstructs_referenceOrder,
		stage.Gstructs_instance,
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
	__gong__computeReferencePass1(stage, stage.Astructs, &stage.Astructs_reference, &stage.Astructs_referenceOrder, &stage.Astructs_instance)

	__gong__computeReferencePass1(stage, stage.AstructBstruct2Uses, &stage.AstructBstruct2Uses_reference, &stage.AstructBstruct2Uses_referenceOrder, &stage.AstructBstruct2Uses_instance)

	__gong__computeReferencePass1(stage, stage.AstructBstructUses, &stage.AstructBstructUses_reference, &stage.AstructBstructUses_referenceOrder, &stage.AstructBstructUses_instance)

	__gong__computeReferencePass1(stage, stage.Bstructs, &stage.Bstructs_reference, &stage.Bstructs_referenceOrder, &stage.Bstructs_instance)

	__gong__computeReferencePass1(stage, stage.Dstructs, &stage.Dstructs_reference, &stage.Dstructs_referenceOrder, &stage.Dstructs_instance)

	__gong__computeReferencePass1(stage, stage.F0123456789012345678901234567890s, &stage.F0123456789012345678901234567890s_reference, &stage.F0123456789012345678901234567890s_referenceOrder, &stage.F0123456789012345678901234567890s_instance)

	__gong__computeReferencePass1(stage, stage.Gstructs, &stage.Gstructs_reference, &stage.Gstructs_referenceOrder, &stage.Gstructs_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.Astructs, stage.Astructs_reference, stage)

	__gong__computeReferencePass2(stage.AstructBstruct2Uses, stage.AstructBstruct2Uses_reference, stage)

	__gong__computeReferencePass2(stage.AstructBstructUses, stage.AstructBstructUses_reference, stage)

	__gong__computeReferencePass2(stage.Bstructs, stage.Bstructs_reference, stage)

	__gong__computeReferencePass2(stage.Dstructs, stage.Dstructs_reference, stage)

	__gong__computeReferencePass2(stage.F0123456789012345678901234567890s, stage.F0123456789012345678901234567890s_reference, stage)

	__gong__computeReferencePass2(stage.Gstructs, stage.Gstructs_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (astruct *Astruct) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Astruct_stagedOrder, stage.Astructs_referenceOrder, astruct, "Astruct")
}

func (astructbstruct2use *AstructBstruct2Use) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AstructBstruct2Use_stagedOrder, stage.AstructBstruct2Uses_referenceOrder, astructbstruct2use, "AstructBstruct2Use")
}

func (astructbstructuse *AstructBstructUse) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.AstructBstructUse_stagedOrder, stage.AstructBstructUses_referenceOrder, astructbstructuse, "AstructBstructUse")
}

func (bstruct *Bstruct) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Bstruct_stagedOrder, stage.Bstructs_referenceOrder, bstruct, "Bstruct")
}

func (dstruct *Dstruct) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Dstruct_stagedOrder, stage.Dstructs_referenceOrder, dstruct, "Dstruct")
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.F0123456789012345678901234567890_stagedOrder, stage.F0123456789012345678901234567890s_referenceOrder, f0123456789012345678901234567890, "F0123456789012345678901234567890")
}

func (gstruct *Gstruct) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Gstruct_stagedOrder, stage.Gstructs_referenceOrder, gstruct, "Gstruct")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (astruct *Astruct) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(astruct, astruct.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (astruct *Astruct) GongGetReferenceIdentifier(stage *Stage) string {
	return astruct.GongGetIdentifier(stage)
}

func (astructbstruct2use *AstructBstruct2Use) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(astructbstruct2use, astructbstruct2use.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (astructbstruct2use *AstructBstruct2Use) GongGetReferenceIdentifier(stage *Stage) string {
	return astructbstruct2use.GongGetIdentifier(stage)
}

func (astructbstructuse *AstructBstructUse) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(astructbstructuse, astructbstructuse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (astructbstructuse *AstructBstructUse) GongGetReferenceIdentifier(stage *Stage) string {
	return astructbstructuse.GongGetIdentifier(stage)
}

func (bstruct *Bstruct) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(bstruct, bstruct.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bstruct *Bstruct) GongGetReferenceIdentifier(stage *Stage) string {
	return bstruct.GongGetIdentifier(stage)
}

func (dstruct *Dstruct) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(dstruct, dstruct.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dstruct *Dstruct) GongGetReferenceIdentifier(stage *Stage) string {
	return dstruct.GongGetIdentifier(stage)
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(f0123456789012345678901234567890, f0123456789012345678901234567890.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongGetReferenceIdentifier(stage *Stage) string {
	return f0123456789012345678901234567890.GongGetIdentifier(stage)
}

func (gstruct *Gstruct) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(gstruct, gstruct.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gstruct *Gstruct) GongGetReferenceIdentifier(stage *Stage) string {
	return gstruct.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (astruct *Astruct) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(astruct.GongGetIdentifier(stage), "Astruct", astruct.Name)
}

func (astructbstruct2use *AstructBstruct2Use) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(astructbstruct2use.GongGetIdentifier(stage), "AstructBstruct2Use", astructbstruct2use.Name)
}

func (astructbstructuse *AstructBstructUse) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(astructbstructuse.GongGetIdentifier(stage), "AstructBstructUse", astructbstructuse.Name)
}

func (bstruct *Bstruct) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(bstruct.GongGetIdentifier(stage), "Bstruct", bstruct.Name)
}

func (dstruct *Dstruct) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(dstruct.GongGetIdentifier(stage), "Dstruct", dstruct.Name)
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(f0123456789012345678901234567890.GongGetIdentifier(stage), "F0123456789012345678901234567890", f0123456789012345678901234567890.Name)
}

func (gstruct *Gstruct) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(gstruct.GongGetIdentifier(stage), "Gstruct", gstruct.Name)
}

// insertion point for unstaging
func (astruct *Astruct) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(astruct.GongGetReferenceIdentifier(stage))
}

func (astructbstruct2use *AstructBstruct2Use) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(astructbstruct2use.GongGetReferenceIdentifier(stage))
}

func (astructbstructuse *AstructBstructUse) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(astructbstructuse.GongGetReferenceIdentifier(stage))
}

func (bstruct *Bstruct) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(bstruct.GongGetReferenceIdentifier(stage))
}

func (dstruct *Dstruct) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(dstruct.GongGetReferenceIdentifier(stage))
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(f0123456789012345678901234567890.GongGetReferenceIdentifier(stage))
}

func (gstruct *Gstruct) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(gstruct.GongGetReferenceIdentifier(stage))
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
