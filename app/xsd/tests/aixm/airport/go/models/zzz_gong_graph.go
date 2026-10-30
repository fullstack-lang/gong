// generated code - do not edit
package models

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
func (aixmbasicmessage *AIXMBasicMessage) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AIXMBasicMessages[aixmbasicmessage]
	return ok
}

func (aixmbasicmessagetype *AIXMBasicMessageType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AIXMBasicMessageTypes[aixmbasicmessagetype]
	return ok
}

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.BasicMessageMemberAIXMPropertyTypes[basicmessagememberaixmpropertytype]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (aixmbasicmessage *AIXMBasicMessage) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(aixmbasicmessage) {
		return
	}

	aixmbasicmessage.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _basicmessagememberaixmpropertytype := range aixmbasicmessage.HasMember {
		stage.StageBranch(_basicmessagememberaixmpropertytype)
	}

}

func (aixmbasicmessagetype *AIXMBasicMessageType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(aixmbasicmessagetype) {
		return
	}

	aixmbasicmessagetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _basicmessagememberaixmpropertytype := range aixmbasicmessagetype.HasMember {
		stage.StageBranch(_basicmessagememberaixmpropertytype)
	}

}

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(basicmessagememberaixmpropertytype) {
		return
	}

	basicmessagememberaixmpropertytype.Stage(stage)

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
	case *AIXMBasicMessage:
		toT := GongCopyBranchAIXMBasicMessage(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AIXMBasicMessageType:
		toT := GongCopyBranchAIXMBasicMessageType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BasicMessageMemberAIXMPropertyType:
		toT := GongCopyBranchBasicMessageMemberAIXMPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAIXMBasicMessage(mapOrigCopy map[any]any, aixmbasicmessageFrom *AIXMBasicMessage) (aixmbasicmessageTo *AIXMBasicMessage) {
	var alreadyCopied bool
	aixmbasicmessageTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, aixmbasicmessageFrom)
	if alreadyCopied {
		return
	}
	aixmbasicmessageFrom.GongCopyBasicFields(aixmbasicmessageTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _basicmessagememberaixmpropertytype := range aixmbasicmessageFrom.HasMember {
		aixmbasicmessageTo.HasMember = append(aixmbasicmessageTo.HasMember, GongCopyBranchBasicMessageMemberAIXMPropertyType(mapOrigCopy, _basicmessagememberaixmpropertytype))
	}

	return
}

func GongCopyBranchAIXMBasicMessageType(mapOrigCopy map[any]any, aixmbasicmessagetypeFrom *AIXMBasicMessageType) (aixmbasicmessagetypeTo *AIXMBasicMessageType) {
	var alreadyCopied bool
	aixmbasicmessagetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, aixmbasicmessagetypeFrom)
	if alreadyCopied {
		return
	}
	aixmbasicmessagetypeFrom.GongCopyBasicFields(aixmbasicmessagetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _basicmessagememberaixmpropertytype := range aixmbasicmessagetypeFrom.HasMember {
		aixmbasicmessagetypeTo.HasMember = append(aixmbasicmessagetypeTo.HasMember, GongCopyBranchBasicMessageMemberAIXMPropertyType(mapOrigCopy, _basicmessagememberaixmpropertytype))
	}

	return
}

func GongCopyBranchBasicMessageMemberAIXMPropertyType(mapOrigCopy map[any]any, basicmessagememberaixmpropertytypeFrom *BasicMessageMemberAIXMPropertyType) (basicmessagememberaixmpropertytypeTo *BasicMessageMemberAIXMPropertyType) {
	var alreadyCopied bool
	basicmessagememberaixmpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, basicmessagememberaixmpropertytypeFrom)
	if alreadyCopied {
		return
	}
	basicmessagememberaixmpropertytypeFrom.GongCopyBasicFields(basicmessagememberaixmpropertytypeTo)

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
func (aixmbasicmessage *AIXMBasicMessage) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(aixmbasicmessage) {
		return
	}

	aixmbasicmessage.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _basicmessagememberaixmpropertytype := range aixmbasicmessage.HasMember {
		stage.UnstageBranch(_basicmessagememberaixmpropertytype)
	}

}

func (aixmbasicmessagetype *AIXMBasicMessageType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(aixmbasicmessagetype) {
		return
	}

	aixmbasicmessagetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _basicmessagememberaixmpropertytype := range aixmbasicmessagetype.HasMember {
		stage.UnstageBranch(_basicmessagememberaixmpropertytype)
	}

}

func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(basicmessagememberaixmpropertytype) {
		return
	}

	basicmessagememberaixmpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *AIXMBasicMessage) GongReconstructPointersFromReferences(stage *Stage, instance *AIXMBasicMessage) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.HasMember, stage.BasicMessageMemberAIXMPropertyTypes_reference, instance.HasMember)
}

func (reference *AIXMBasicMessageType) GongReconstructPointersFromReferences(stage *Stage, instance *AIXMBasicMessageType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.HasMember, stage.BasicMessageMemberAIXMPropertyTypes_reference, instance.HasMember)
}

func (reference *BasicMessageMemberAIXMPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *BasicMessageMemberAIXMPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AIXMBasicMessage) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.HasMember, stage.BasicMessageMemberAIXMPropertyTypes_instance)
}

func (reference *AIXMBasicMessageType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.HasMember, stage.BasicMessageMemberAIXMPropertyTypes_instance)
}

func (reference *BasicMessageMemberAIXMPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (aixmbasicmessage *AIXMBasicMessage) GongDiff(stage *Stage, aixmbasicmessageOther *AIXMBasicMessage) (diffs []string) {
	// insertion point for field diffs
	if aixmbasicmessage.Name != aixmbasicmessageOther.Name {
		diffs = append(diffs, aixmbasicmessage.GongMarshallField(stage, "Name"))
	}
	if aixmbasicmessage.Id != aixmbasicmessageOther.Id {
		diffs = append(diffs, aixmbasicmessage.GongMarshallField(stage, "Id"))
	}
	if aixmbasicmessage.Description != aixmbasicmessageOther.Description {
		diffs = append(diffs, aixmbasicmessage.GongMarshallField(stage, "Description"))
	}
	if aixmbasicmessage.SequenceNumber != aixmbasicmessageOther.SequenceNumber {
		diffs = append(diffs, aixmbasicmessage.GongMarshallField(stage, "SequenceNumber"))
	}
	if aixmbasicmessage.AggregationType != aixmbasicmessageOther.AggregationType {
		diffs = append(diffs, aixmbasicmessage.GongMarshallField(stage, "AggregationType"))
	}
	if ops := __gong__diffSliceOfPointers(stage, aixmbasicmessage, "HasMember", aixmbasicmessageOther.HasMember, aixmbasicmessage.HasMember); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (aixmbasicmessagetype *AIXMBasicMessageType) GongDiff(stage *Stage, aixmbasicmessagetypeOther *AIXMBasicMessageType) (diffs []string) {
	// insertion point for field diffs
	if aixmbasicmessagetype.Name != aixmbasicmessagetypeOther.Name {
		diffs = append(diffs, aixmbasicmessagetype.GongMarshallField(stage, "Name"))
	}
	if aixmbasicmessagetype.Id != aixmbasicmessagetypeOther.Id {
		diffs = append(diffs, aixmbasicmessagetype.GongMarshallField(stage, "Id"))
	}
	if aixmbasicmessagetype.Description != aixmbasicmessagetypeOther.Description {
		diffs = append(diffs, aixmbasicmessagetype.GongMarshallField(stage, "Description"))
	}
	if aixmbasicmessagetype.SequenceNumber != aixmbasicmessagetypeOther.SequenceNumber {
		diffs = append(diffs, aixmbasicmessagetype.GongMarshallField(stage, "SequenceNumber"))
	}
	if aixmbasicmessagetype.AggregationType != aixmbasicmessagetypeOther.AggregationType {
		diffs = append(diffs, aixmbasicmessagetype.GongMarshallField(stage, "AggregationType"))
	}
	if ops := __gong__diffSliceOfPointers(stage, aixmbasicmessagetype, "HasMember", aixmbasicmessagetypeOther.HasMember, aixmbasicmessagetype.HasMember); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (basicmessagememberaixmpropertytype *BasicMessageMemberAIXMPropertyType) GongDiff(stage *Stage, basicmessagememberaixmpropertytypeOther *BasicMessageMemberAIXMPropertyType) (diffs []string) {
	// insertion point for field diffs
	if basicmessagememberaixmpropertytype.Name != basicmessagememberaixmpropertytypeOther.Name {
		diffs = append(diffs, basicmessagememberaixmpropertytype.GongMarshallField(stage, "Name"))
	}
	if basicmessagememberaixmpropertytype.NilReason != basicmessagememberaixmpropertytypeOther.NilReason {
		diffs = append(diffs, basicmessagememberaixmpropertytype.GongMarshallField(stage, "NilReason"))
	}
	if basicmessagememberaixmpropertytype.Type_XSD != basicmessagememberaixmpropertytypeOther.Type_XSD {
		diffs = append(diffs, basicmessagememberaixmpropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if basicmessagememberaixmpropertytype.Href != basicmessagememberaixmpropertytypeOther.Href {
		diffs = append(diffs, basicmessagememberaixmpropertytype.GongMarshallField(stage, "Href"))
	}
	if basicmessagememberaixmpropertytype.Role != basicmessagememberaixmpropertytypeOther.Role {
		diffs = append(diffs, basicmessagememberaixmpropertytype.GongMarshallField(stage, "Role"))
	}
	if basicmessagememberaixmpropertytype.Arcrole != basicmessagememberaixmpropertytypeOther.Arcrole {
		diffs = append(diffs, basicmessagememberaixmpropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if basicmessagememberaixmpropertytype.Title != basicmessagememberaixmpropertytypeOther.Title {
		diffs = append(diffs, basicmessagememberaixmpropertytype.GongMarshallField(stage, "Title"))
	}
	if basicmessagememberaixmpropertytype.Show != basicmessagememberaixmpropertytypeOther.Show {
		diffs = append(diffs, basicmessagememberaixmpropertytype.GongMarshallField(stage, "Show"))
	}
	if basicmessagememberaixmpropertytype.Actuate != basicmessagememberaixmpropertytypeOther.Actuate {
		diffs = append(diffs, basicmessagememberaixmpropertytype.GongMarshallField(stage, "Actuate"))
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
