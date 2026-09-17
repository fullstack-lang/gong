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

	// Compute reverse map for named struct AstructBstruct2Use
	// insertion point per field

	// Compute reverse map for named struct AstructBstructUse
	// insertion point per field

	// Compute reverse map for named struct Bstruct
	// insertion point per field

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

	// Compute reverse map for named struct F0123456789012345678901234567890
	// insertion point per field

	// Compute reverse map for named struct Gstruct
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.Astructs {
		res = append(res, instance)
	}

	for instance := range stage.AstructBstruct2Uses {
		res = append(res, instance)
	}

	for instance := range stage.AstructBstructUses {
		res = append(res, instance)
	}

	for instance := range stage.Bstructs {
		res = append(res, instance)
	}

	for instance := range stage.Dstructs {
		res = append(res, instance)
	}

	for instance := range stage.F0123456789012345678901234567890s {
		res = append(res, instance)
	}

	for instance := range stage.Gstructs {
		res = append(res, instance)
	}

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
func (astruct *Astruct) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(astruct).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(astruct), uint64(stage.GetOrder(astruct)))
	return
}

func (astructbstruct2use *AstructBstruct2Use) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(astructbstruct2use).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(astructbstruct2use), uint64(stage.GetOrder(astructbstruct2use)))
	return
}

func (astructbstructuse *AstructBstructUse) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(astructbstructuse).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(astructbstructuse), uint64(stage.GetOrder(astructbstructuse)))
	return
}

func (bstruct *Bstruct) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bstruct).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(bstruct), uint64(stage.GetOrder(bstruct)))
	return
}

func (dstruct *Dstruct) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(dstruct).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(dstruct), uint64(stage.GetOrder(dstruct)))
	return
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(f0123456789012345678901234567890).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(f0123456789012345678901234567890), uint64(stage.GetOrder(f0123456789012345678901234567890)))
	return
}

func (gstruct *Gstruct) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(gstruct).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(gstruct), uint64(stage.GetOrder(gstruct)))
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
	stage.Astructs_reference = make(map[*Astruct]*Astruct)
	stage.Astructs_referenceOrder = make(map[*Astruct]uint) // diff Unstage needs the reference order
	stage.Astructs_instance = make(map[*Astruct]*Astruct)
	for instance := range stage.Astructs {
		_copy := instance.GongCopy().(*Astruct)
		stage.Astructs_reference[instance] = _copy
		stage.Astructs_instance[_copy] = instance
		stage.Astructs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.AstructBstruct2Uses_reference = make(map[*AstructBstruct2Use]*AstructBstruct2Use)
	stage.AstructBstruct2Uses_referenceOrder = make(map[*AstructBstruct2Use]uint) // diff Unstage needs the reference order
	stage.AstructBstruct2Uses_instance = make(map[*AstructBstruct2Use]*AstructBstruct2Use)
	for instance := range stage.AstructBstruct2Uses {
		_copy := instance.GongCopy().(*AstructBstruct2Use)
		stage.AstructBstruct2Uses_reference[instance] = _copy
		stage.AstructBstruct2Uses_instance[_copy] = instance
		stage.AstructBstruct2Uses_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.AstructBstructUses_reference = make(map[*AstructBstructUse]*AstructBstructUse)
	stage.AstructBstructUses_referenceOrder = make(map[*AstructBstructUse]uint) // diff Unstage needs the reference order
	stage.AstructBstructUses_instance = make(map[*AstructBstructUse]*AstructBstructUse)
	for instance := range stage.AstructBstructUses {
		_copy := instance.GongCopy().(*AstructBstructUse)
		stage.AstructBstructUses_reference[instance] = _copy
		stage.AstructBstructUses_instance[_copy] = instance
		stage.AstructBstructUses_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Bstructs_reference = make(map[*Bstruct]*Bstruct)
	stage.Bstructs_referenceOrder = make(map[*Bstruct]uint) // diff Unstage needs the reference order
	stage.Bstructs_instance = make(map[*Bstruct]*Bstruct)
	for instance := range stage.Bstructs {
		_copy := instance.GongCopy().(*Bstruct)
		stage.Bstructs_reference[instance] = _copy
		stage.Bstructs_instance[_copy] = instance
		stage.Bstructs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Dstructs_reference = make(map[*Dstruct]*Dstruct)
	stage.Dstructs_referenceOrder = make(map[*Dstruct]uint) // diff Unstage needs the reference order
	stage.Dstructs_instance = make(map[*Dstruct]*Dstruct)
	for instance := range stage.Dstructs {
		_copy := instance.GongCopy().(*Dstruct)
		stage.Dstructs_reference[instance] = _copy
		stage.Dstructs_instance[_copy] = instance
		stage.Dstructs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.F0123456789012345678901234567890s_reference = make(map[*F0123456789012345678901234567890]*F0123456789012345678901234567890)
	stage.F0123456789012345678901234567890s_referenceOrder = make(map[*F0123456789012345678901234567890]uint) // diff Unstage needs the reference order
	stage.F0123456789012345678901234567890s_instance = make(map[*F0123456789012345678901234567890]*F0123456789012345678901234567890)
	for instance := range stage.F0123456789012345678901234567890s {
		_copy := instance.GongCopy().(*F0123456789012345678901234567890)
		stage.F0123456789012345678901234567890s_reference[instance] = _copy
		stage.F0123456789012345678901234567890s_instance[_copy] = instance
		stage.F0123456789012345678901234567890s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Gstructs_reference = make(map[*Gstruct]*Gstruct)
	stage.Gstructs_referenceOrder = make(map[*Gstruct]uint) // diff Unstage needs the reference order
	stage.Gstructs_instance = make(map[*Gstruct]*Gstruct)
	for instance := range stage.Gstructs {
		_copy := instance.GongCopy().(*Gstruct)
		stage.Gstructs_reference[instance] = _copy
		stage.Gstructs_instance[_copy] = instance
		stage.Gstructs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.Astructs {
		reference := stage.Astructs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.AstructBstruct2Uses {
		reference := stage.AstructBstruct2Uses_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.AstructBstructUses {
		reference := stage.AstructBstructUses_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Bstructs {
		reference := stage.Bstructs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Dstructs {
		reference := stage.Dstructs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.F0123456789012345678901234567890s {
		reference := stage.F0123456789012345678901234567890s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Gstructs {
		reference := stage.Gstructs_reference[instance]
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
func (astruct *Astruct) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Astruct_stagedOrder[astruct]; ok {
		return order
	}
	if order, ok := stage.Astructs_referenceOrder[astruct]; ok {
		return order
	} else {
		log.Printf("instance %p of type Astruct was not staged and does not have a reference order", astruct)
		return 0
	}
}

func (astructbstruct2use *AstructBstruct2Use) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AstructBstruct2Use_stagedOrder[astructbstruct2use]; ok {
		return order
	}
	if order, ok := stage.AstructBstruct2Uses_referenceOrder[astructbstruct2use]; ok {
		return order
	} else {
		log.Printf("instance %p of type AstructBstruct2Use was not staged and does not have a reference order", astructbstruct2use)
		return 0
	}
}

func (astructbstructuse *AstructBstructUse) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.AstructBstructUse_stagedOrder[astructbstructuse]; ok {
		return order
	}
	if order, ok := stage.AstructBstructUses_referenceOrder[astructbstructuse]; ok {
		return order
	} else {
		log.Printf("instance %p of type AstructBstructUse was not staged and does not have a reference order", astructbstructuse)
		return 0
	}
}

func (bstruct *Bstruct) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Bstruct_stagedOrder[bstruct]; ok {
		return order
	}
	if order, ok := stage.Bstructs_referenceOrder[bstruct]; ok {
		return order
	} else {
		log.Printf("instance %p of type Bstruct was not staged and does not have a reference order", bstruct)
		return 0
	}
}

func (dstruct *Dstruct) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Dstruct_stagedOrder[dstruct]; ok {
		return order
	}
	if order, ok := stage.Dstructs_referenceOrder[dstruct]; ok {
		return order
	} else {
		log.Printf("instance %p of type Dstruct was not staged and does not have a reference order", dstruct)
		return 0
	}
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.F0123456789012345678901234567890_stagedOrder[f0123456789012345678901234567890]; ok {
		return order
	}
	if order, ok := stage.F0123456789012345678901234567890s_referenceOrder[f0123456789012345678901234567890]; ok {
		return order
	} else {
		log.Printf("instance %p of type F0123456789012345678901234567890 was not staged and does not have a reference order", f0123456789012345678901234567890)
		return 0
	}
}

func (gstruct *Gstruct) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Gstruct_stagedOrder[gstruct]; ok {
		return order
	}
	if order, ok := stage.Gstructs_referenceOrder[gstruct]; ok {
		return order
	} else {
		log.Printf("instance %p of type Gstruct was not staged and does not have a reference order", gstruct)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (astruct *Astruct) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", astruct.GongGetGongstructName(), astruct.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (astruct *Astruct) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", astruct.GongGetGongstructName(), astruct.GongGetOrder(stage))
}

func (astructbstruct2use *AstructBstruct2Use) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", astructbstruct2use.GongGetGongstructName(), astructbstruct2use.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (astructbstruct2use *AstructBstruct2Use) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", astructbstruct2use.GongGetGongstructName(), astructbstruct2use.GongGetOrder(stage))
}

func (astructbstructuse *AstructBstructUse) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", astructbstructuse.GongGetGongstructName(), astructbstructuse.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (astructbstructuse *AstructBstructUse) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", astructbstructuse.GongGetGongstructName(), astructbstructuse.GongGetOrder(stage))
}

func (bstruct *Bstruct) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bstruct.GongGetGongstructName(), bstruct.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bstruct *Bstruct) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bstruct.GongGetGongstructName(), bstruct.GongGetOrder(stage))
}

func (dstruct *Dstruct) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dstruct.GongGetGongstructName(), dstruct.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dstruct *Dstruct) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dstruct.GongGetGongstructName(), dstruct.GongGetOrder(stage))
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", f0123456789012345678901234567890.GongGetGongstructName(), f0123456789012345678901234567890.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", f0123456789012345678901234567890.GongGetGongstructName(), f0123456789012345678901234567890.GongGetOrder(stage))
}

func (gstruct *Gstruct) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gstruct.GongGetGongstructName(), gstruct.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (gstruct *Gstruct) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", gstruct.GongGetGongstructName(), gstruct.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (astruct *Astruct) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", astruct.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Astruct")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(astruct.Name))
	return
}

func (astructbstruct2use *AstructBstruct2Use) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", astructbstruct2use.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AstructBstruct2Use")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(astructbstruct2use.Name))
	return
}

func (astructbstructuse *AstructBstructUse) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", astructbstructuse.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "AstructBstructUse")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(astructbstructuse.Name))
	return
}

func (bstruct *Bstruct) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bstruct.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bstruct")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(bstruct.Name))
	return
}

func (dstruct *Dstruct) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dstruct.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Dstruct")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(dstruct.Name))
	return
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", f0123456789012345678901234567890.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "F0123456789012345678901234567890")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(f0123456789012345678901234567890.Name))
	return
}

func (gstruct *Gstruct) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gstruct.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Gstruct")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(gstruct.Name))
	return
}

// insertion point for unstaging
func (astruct *Astruct) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", astruct.GongGetReferenceIdentifier(stage))
	return
}

func (astructbstruct2use *AstructBstruct2Use) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", astructbstruct2use.GongGetReferenceIdentifier(stage))
	return
}

func (astructbstructuse *AstructBstructUse) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", astructbstructuse.GongGetReferenceIdentifier(stage))
	return
}

func (bstruct *Bstruct) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bstruct.GongGetReferenceIdentifier(stage))
	return
}

func (dstruct *Dstruct) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dstruct.GongGetReferenceIdentifier(stage))
	return
}

func (f0123456789012345678901234567890 *F0123456789012345678901234567890) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", f0123456789012345678901234567890.GongGetReferenceIdentifier(stage))
	return
}

func (gstruct *Gstruct) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", gstruct.GongGetReferenceIdentifier(stage))
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
