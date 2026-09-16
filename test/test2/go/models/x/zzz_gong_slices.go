// generated code - do not edit
package x

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
	// Compute reverse map for named struct X
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.Xs {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (x *X) GongCopy() GongstructIF {
	newInstance := new(X)
	x.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (x *X) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(x).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(x), uint64(stage.GetOrder(x)))
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
	var xs_newInstances []*X
	var xs_deletedInstances []*X

	// parse all staged instances and check if they have a reference
	for x := range stage.Xs {
		if ref, ok := stage.Xs_reference[x]; !ok {
			xs_newInstances = append(xs_newInstances, x)
			newInstancesSlice = append(newInstancesSlice, x.GongMarshallIdentifier(stage))
			if stage.Xs_referenceOrder == nil {
				stage.Xs_referenceOrder = make(map[*X]uint)
			}
			stage.Xs_referenceOrder[x] = stage.X_stagedOrder[x]
			newInstancesReverseSlice = append(newInstancesReverseSlice, x.GongMarshallUnstaging(stage))
			// delete(stage.Xs_referenceOrder, x)
			fieldInitializers, pointersInitializations := x.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.X_stagedOrder[ref] = stage.X_stagedOrder[x]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := x.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, x)
			// delete(stage.X_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if x.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", x.GetName())
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
	for _, ref := range stage.Xs_reference {
		instance := stage.Xs_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Xs[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			xs_deletedInstances = append(xs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(xs_newInstances)
	lenDeletedInstances += len(xs_deletedInstances)

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
	stage.Xs_reference = make(map[*X]*X)
	stage.Xs_referenceOrder = make(map[*X]uint) // diff Unstage needs the reference order
	stage.Xs_instance = make(map[*X]*X)
	for instance := range stage.Xs {
		_copy := instance.GongCopy().(*X)
		stage.Xs_reference[instance] = _copy
		stage.Xs_instance[_copy] = instance
		stage.Xs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.Xs {
		reference := stage.Xs_reference[instance]
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
func (x *X) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.X_stagedOrder[x]; ok {
		return order
	}
	if order, ok := stage.Xs_referenceOrder[x]; ok {
		return order
	} else {
		log.Printf("instance %p of type X was not staged and does not have a reference order", x)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (x *X) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", x.GongGetGongstructName(), x.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (x *X) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", x.GongGetGongstructName(), x.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (x *X) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", x.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "X")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(x.Name))
	return
}

// insertion point for unstaging
func (x *X) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", x.GongGetReferenceIdentifier(stage))
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
