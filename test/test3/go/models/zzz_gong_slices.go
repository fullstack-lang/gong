// generated code - do not edit
package models

import (
	"fmt"
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

}

func (stage *Stage) GetInstances() (res []GongstructIF) {

	// insertion point per named struct
	for instance := range stage.As {
		res = append(res, instance)
	}

	for instance := range stage.Bs {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (a *A) GongCopy() GongstructIF {
	newInstance := *a
	return &newInstance
}

func (b *B) GongCopy() GongstructIF {
	newInstance := *b
	return &newInstance
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
			stage.As_referenceOrder[a] = stage.AMap_Staged_Order[a]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a.GongMarshallUnstaging(stage))
			delete(stage.As_referenceOrder, a)
			fieldInitializers, pointersInitializations := a.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.AMap_Staged_Order[ref] = stage.AMap_Staged_Order[a]
			diffs := a.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a)
			delete(stage.AMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a.GetName())
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
	for ref := range stage.As_reference {
		if _, ok := stage.As[ref]; !ok {
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
			stage.Bs_referenceOrder[b] = stage.BMap_Staged_Order[b]
			newInstancesReverseSlice = append(newInstancesReverseSlice, b.GongMarshallUnstaging(stage))
			delete(stage.Bs_referenceOrder, b)
			fieldInitializers, pointersInitializations := b.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.BMap_Staged_Order[ref] = stage.BMap_Staged_Order[b]
			diffs := b.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, b)
			delete(stage.BMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", b.GetName())
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
	for ref := range stage.Bs_reference {
		if _, ok := stage.Bs[ref]; !ok {
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
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {

	// insertion point per named struct
	stage.As_reference = make(map[*A]*A)
	stage.As_referenceOrder = make(map[*A]uint) // diff Unstage needs the reference order
	for instance := range stage.As {
		stage.As_reference[instance] = instance.GongCopy().(*A)
		stage.As_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.Bs_reference = make(map[*B]*B)
	stage.Bs_referenceOrder = make(map[*B]uint) // diff Unstage needs the reference order
	for instance := range stage.Bs {
		stage.Bs_reference[instance] = instance.GongCopy().(*B)
		stage.Bs_referenceOrder[instance] = instance.GongGetOrder(stage)
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
	if order, ok := stage.AMap_Staged_Order[a]; ok {
		return order
	}
	return stage.As_referenceOrder[a]
}

func (a *A) GongGetReferenceOrder(stage *Stage) uint {
	return stage.As_referenceOrder[a]
}

func (b *B) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.BMap_Staged_Order[b]; ok {
		return order
	}
	return stage.Bs_referenceOrder[b]
}

func (b *B) GongGetReferenceOrder(stage *Stage) uint {
	return stage.Bs_referenceOrder[b]
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
	return fmt.Sprintf("__%s__%08d_", a.GongGetGongstructName(), a.GongGetReferenceOrder(stage))
}

func (b *B) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", b.GongGetGongstructName(), b.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (b *B) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", b.GongGetGongstructName(), b.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (a *A) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a.Name)
	return
}
func (b *B) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", b.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "B")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", b.Name)
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
