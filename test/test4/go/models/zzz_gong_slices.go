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

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.Astructs {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (astruct *Astruct) GongCopy() GongstructIF {
	newInstance := new(Astruct)
	astruct.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (astruct *Astruct) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(astruct).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(astruct), uint64(GetOrderPointerGongstruct(stage, astruct)))
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
	var astructs_newInstances []*Astruct
	var astructs_deletedInstances []*Astruct

	// parse all staged instances and check if they have a reference
	for astruct := range stage.Astructs {
		if ref, ok := stage.Astructs_reference[astruct]; !ok {
			astructs_newInstances = append(astructs_newInstances, astruct)
			newInstancesSlice = append(newInstancesSlice, astruct.GongMarshallIdentifier(stage))
			if stage.Astructs_referenceOrder == nil {
				stage.Astructs_referenceOrder = make(map[*Astruct]uint)
			}
			stage.Astructs_referenceOrder[astruct] = stage.Astruct_stagedOrder[astruct]
			newInstancesReverseSlice = append(newInstancesReverseSlice, astruct.GongMarshallUnstaging(stage))
			// delete(stage.Astructs_referenceOrder, astruct)
			fieldInitializers, pointersInitializations := astruct.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.Astruct_stagedOrder[ref] = stage.Astruct_stagedOrder[astruct]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := astruct.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, astruct)
			// delete(stage.Astruct_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if astruct.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", astruct.GetName())
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
	for _, ref := range stage.Astructs_reference {
		instance := stage.Astructs_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Astructs[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			astructs_deletedInstances = append(astructs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(astructs_newInstances)
	lenDeletedInstances += len(astructs_deletedInstances)

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

	// insertion point per named struct
	for instance := range stage.Astructs {
		reference := stage.Astructs_reference[instance]
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

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (astruct *Astruct) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", astruct.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Astruct")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(astruct.Name))
	return
}

// insertion point for unstaging
func (astruct *Astruct) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", astruct.GongGetReferenceIdentifier(stage))
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
