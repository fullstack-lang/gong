// generated code - do not edit
package gcobase

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (abstractobject_type *AbstractObject_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractObject_Types[abstractobject_type]
	return ok
}

func (codelistvalue_type *CodeListValue_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeListValue_Types[codelistvalue_type]
	return ok
}

func (objectreference_propertytype *ObjectReference_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ObjectReference_PropertyTypes[objectreference_propertytype]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (abstractobject_type *AbstractObject_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractobject_type) {
		return
	}

	abstractobject_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelistvalue_type *CodeListValue_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codelistvalue_type) {
		return
	}

	codelistvalue_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (objectreference_propertytype *ObjectReference_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(objectreference_propertytype) {
		return
	}

	objectreference_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *AbstractObject_Type:
		toT := GongCopyBranchAbstractObject_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeListValue_Type:
		toT := GongCopyBranchCodeListValue_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ObjectReference_PropertyType:
		toT := GongCopyBranchObjectReference_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAbstractObject_Type(mapOrigCopy map[any]any, abstractobject_typeFrom *AbstractObject_Type) (abstractobject_typeTo *AbstractObject_Type) {
	var alreadyCopied bool
	abstractobject_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractobject_typeFrom)
	if alreadyCopied {
		return
	}
	abstractobject_typeFrom.GongCopyBasicFields(abstractobject_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeListValue_Type(mapOrigCopy map[any]any, codelistvalue_typeFrom *CodeListValue_Type) (codelistvalue_typeTo *CodeListValue_Type) {
	var alreadyCopied bool
	codelistvalue_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codelistvalue_typeFrom)
	if alreadyCopied {
		return
	}
	codelistvalue_typeFrom.GongCopyBasicFields(codelistvalue_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchObjectReference_PropertyType(mapOrigCopy map[any]any, objectreference_propertytypeFrom *ObjectReference_PropertyType) (objectreference_propertytypeTo *ObjectReference_PropertyType) {
	var alreadyCopied bool
	objectreference_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, objectreference_propertytypeFrom)
	if alreadyCopied {
		return
	}
	objectreference_propertytypeFrom.GongCopyBasicFields(objectreference_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// insertion point for unstage branch per struct
func (abstractobject_type *AbstractObject_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractobject_type) {
		return
	}

	abstractobject_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codelistvalue_type *CodeListValue_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codelistvalue_type) {
		return
	}

	codelistvalue_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (objectreference_propertytype *ObjectReference_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(objectreference_propertytype) {
		return
	}

	objectreference_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *AbstractObject_Type) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractObject_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeListValue_Type) GongReconstructPointersFromReferences(stage *Stage, instance *CodeListValue_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ObjectReference_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *ObjectReference_PropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AbstractObject_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeListValue_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ObjectReference_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractobject_type *AbstractObject_Type) GongDiff(stage *Stage, abstractobject_typeOther *AbstractObject_Type) (diffs []string) {
	// insertion point for field diffs
	if abstractobject_type.Name != abstractobject_typeOther.Name {
		diffs = append(diffs, abstractobject_type.GongMarshallField(stage, "Name"))
	}
	if abstractobject_type.Id != abstractobject_typeOther.Id {
		diffs = append(diffs, abstractobject_type.GongMarshallField(stage, "Id"))
	}
	if abstractobject_type.Uuid != abstractobject_typeOther.Uuid {
		diffs = append(diffs, abstractobject_type.GongMarshallField(stage, "Uuid"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codelistvalue_type *CodeListValue_Type) GongDiff(stage *Stage, codelistvalue_typeOther *CodeListValue_Type) (diffs []string) {
	// insertion point for field diffs
	if codelistvalue_type.Name != codelistvalue_typeOther.Name {
		diffs = append(diffs, codelistvalue_type.GongMarshallField(stage, "Name"))
	}
	if codelistvalue_type.CodeList != codelistvalue_typeOther.CodeList {
		diffs = append(diffs, codelistvalue_type.GongMarshallField(stage, "CodeList"))
	}
	if codelistvalue_type.CodeListValue != codelistvalue_typeOther.CodeListValue {
		diffs = append(diffs, codelistvalue_type.GongMarshallField(stage, "CodeListValue"))
	}
	if codelistvalue_type.CodeSpace != codelistvalue_typeOther.CodeSpace {
		diffs = append(diffs, codelistvalue_type.GongMarshallField(stage, "CodeSpace"))
	}
	if codelistvalue_type.EnclosedText != codelistvalue_typeOther.EnclosedText {
		diffs = append(diffs, codelistvalue_type.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (objectreference_propertytype *ObjectReference_PropertyType) GongDiff(stage *Stage, objectreference_propertytypeOther *ObjectReference_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if objectreference_propertytype.Name != objectreference_propertytypeOther.Name {
		diffs = append(diffs, objectreference_propertytype.GongMarshallField(stage, "Name"))
	}
	if objectreference_propertytype.NilReason != objectreference_propertytypeOther.NilReason {
		diffs = append(diffs, objectreference_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if objectreference_propertytype.Uuidref != objectreference_propertytypeOther.Uuidref {
		diffs = append(diffs, objectreference_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if objectreference_propertytype.Type_XSD != objectreference_propertytypeOther.Type_XSD {
		diffs = append(diffs, objectreference_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if objectreference_propertytype.Href != objectreference_propertytypeOther.Href {
		diffs = append(diffs, objectreference_propertytype.GongMarshallField(stage, "Href"))
	}
	if objectreference_propertytype.Role != objectreference_propertytypeOther.Role {
		diffs = append(diffs, objectreference_propertytype.GongMarshallField(stage, "Role"))
	}
	if objectreference_propertytype.Arcrole != objectreference_propertytypeOther.Arcrole {
		diffs = append(diffs, objectreference_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if objectreference_propertytype.Title != objectreference_propertytypeOther.Title {
		diffs = append(diffs, objectreference_propertytype.GongMarshallField(stage, "Title"))
	}
	if objectreference_propertytype.Show != objectreference_propertytypeOther.Show {
		diffs = append(diffs, objectreference_propertytype.GongMarshallField(stage, "Show"))
	}
	if objectreference_propertytype.Actuate != objectreference_propertytypeOther.Actuate {
		diffs = append(diffs, objectreference_propertytype.GongMarshallField(stage, "Actuate"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
