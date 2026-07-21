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
	// Compute reverse map for named struct A
	// insertion point per field
	stage.A_Bs_reverseMap = make(map[*B]*A)
	for a := range stage.As {
		_ = a
		for _, _b := range a.Bs {
			stage.A_Bs_reverseMap[_b] = a
		}
	}

	// Compute reverse map for named struct B
	// insertion point per field

	// Compute reverse map for named struct C
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.As {
		res = append(res, instance)
	}

	for instance := range stage.Bs {
		res = append(res, instance)
	}

	for instance := range stage.Cs {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (a *A) GongCopy() GongstructIF {
	newInstance := new(A)
	a.CopyBasicFields(newInstance)
	return newInstance
}

func (b *B) GongCopy() GongstructIF {
	newInstance := new(B)
	b.CopyBasicFields(newInstance)
	return newInstance
}

func (c *C) GongCopy() GongstructIF {
	newInstance := new(C)
	c.CopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (a *A) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(a), uint64(GetOrderPointerGongstruct(stage, a)))
	return
}

func (b *B) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(b).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(b), uint64(GetOrderPointerGongstruct(stage, b)))
	return
}

func (c *C) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(c).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GenerateReproducibleUUIDv4(GetGongstructNameFromPointer(c), uint64(GetOrderPointerGongstruct(stage, c)))
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
	var as_newInstances []*A
	var as_deletedInstances []*A

	// parse all staged instances and check if they have a reference
	for a := range stage.As {
		if ref, ok := stage.As_reference[a]; !ok {
			as_newInstances = append(as_newInstances, a)
			newInstancesSlice = append(newInstancesSlice, a.GongMarshallIdentifier(stage))
			if stage.As_referenceOrder == nil {
				stage.As_referenceOrder = make(map[*A]uint)
			}
			stage.As_referenceOrder[a] = stage.A_stagedOrder[a]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a.GongMarshallUnstaging(stage))
			// delete(stage.As_referenceOrder, a)
			fieldInitializers, pointersInitializations := a.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_stagedOrder[ref] = stage.A_stagedOrder[a]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := a.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a)
			// delete(stage.A_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if a.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", a.GetName())
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
	for _, ref := range stage.As_reference {
		instance := stage.As_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.As[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			as_deletedInstances = append(as_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(as_newInstances)
	lenDeletedInstances += len(as_deletedInstances)
	var bs_newInstances []*B
	var bs_deletedInstances []*B

	// parse all staged instances and check if they have a reference
	for b := range stage.Bs {
		if ref, ok := stage.Bs_reference[b]; !ok {
			bs_newInstances = append(bs_newInstances, b)
			newInstancesSlice = append(newInstancesSlice, b.GongMarshallIdentifier(stage))
			if stage.Bs_referenceOrder == nil {
				stage.Bs_referenceOrder = make(map[*B]uint)
			}
			stage.Bs_referenceOrder[b] = stage.B_stagedOrder[b]
			newInstancesReverseSlice = append(newInstancesReverseSlice, b.GongMarshallUnstaging(stage))
			// delete(stage.Bs_referenceOrder, b)
			fieldInitializers, pointersInitializations := b.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.B_stagedOrder[ref] = stage.B_stagedOrder[b]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := b.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, b)
			// delete(stage.B_stagedOrder, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				if b.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", b.GetName())
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
	for _, ref := range stage.Bs_reference {
		instance := stage.Bs_instance[ref]    // get the instance corresponding to the reference
		if _, ok := stage.Bs[instance]; !ok { // if the instance is not staged anymore,  it means it has been unstaged
			bs_deletedInstances = append(bs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(bs_newInstances)
	lenDeletedInstances += len(bs_deletedInstances)

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
	stage.As_reference = make(map[*A]*A)
	stage.As_referenceOrder = make(map[*A]uint) // diff Unstage needs the reference order
	stage.As_instance = make(map[*A]*A)
	for instance := range stage.As {
		_copy := instance.GongCopy().(*A)
		stage.As_reference[instance] = _copy
		stage.As_instance[_copy] = instance
		stage.As_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Bs_reference = make(map[*B]*B)
	stage.Bs_referenceOrder = make(map[*B]uint) // diff Unstage needs the reference order
	stage.Bs_instance = make(map[*B]*B)
	for instance := range stage.Bs {
		_copy := instance.GongCopy().(*B)
		stage.Bs_reference[instance] = _copy
		stage.Bs_instance[_copy] = instance
		stage.Bs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Cs_reference = make(map[*C]*C)
	stage.Cs_referenceOrder = make(map[*C]uint) // diff Unstage needs the reference order
	stage.Cs_instance = make(map[*C]*C)
	for instance := range stage.Cs {
		_copy := instance.GongCopy().(*C)
		stage.Cs_reference[instance] = _copy
		stage.Cs_instance[_copy] = instance
		stage.Cs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.As {
		reference := stage.As_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Bs {
		reference := stage.Bs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Cs {
		reference := stage.Cs_reference[instance]
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
func (a *A) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_stagedOrder[a]; ok {
		return order
	}
	if order, ok := stage.As_referenceOrder[a]; ok {
		return order
	} else {
		log.Printf("instance %p of type A was not staged and does not have a reference order", a)
		return 0
	}
}

func (b *B) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.B_stagedOrder[b]; ok {
		return order
	}
	if order, ok := stage.Bs_referenceOrder[b]; ok {
		return order
	} else {
		log.Printf("instance %p of type B was not staged and does not have a reference order", b)
		return 0
	}
}

func (c *C) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.C_stagedOrder[c]; ok {
		return order
	}
	if order, ok := stage.Cs_referenceOrder[c]; ok {
		return order
	} else {
		log.Printf("instance %p of type C was not staged and does not have a reference order", c)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (a *A) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a.GongGetGongstructName(), a.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a *A) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a.GongGetGongstructName(), a.GongGetOrder(stage))
}

func (b *B) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", b.GongGetGongstructName(), b.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (b *B) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", b.GongGetGongstructName(), b.GongGetOrder(stage))
}

func (c *C) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", c.GongGetGongstructName(), c.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (c *C) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", c.GongGetGongstructName(), c.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (a *A) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(a.Name))
	return
}

func (b *B) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", b.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "B")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(b.Name))
	return
}

func (c *C) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", c.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "C")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", ToRawStringLiteral(c.Name))
	return
}

// insertion point for unstaging
func (a *A) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a.GongGetReferenceIdentifier(stage))
	return
}

func (b *B) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", b.GongGetReferenceIdentifier(stage))
	return
}

func (c *C) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", c.GongGetReferenceIdentifier(stage))
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
