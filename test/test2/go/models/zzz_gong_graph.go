// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *A:
		ok = stage.IsStagedA(target)

	case *B:
		ok = stage.IsStagedB(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *A:
		ok = stage.IsStagedA(target)

	case *B:
		ok = stage.IsStagedB(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedA(a *A) (ok bool) {

	_, ok = stage.As[a]

	return
}

func (stage *Stage) IsStagedB(b *B) (ok bool) {

	_, ok = stage.Bs[b]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *A:
		stage.StageBranchA(target)

	case *B:
		stage.StageBranchB(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchA(a *A) {

	// check if instance is already staged
	if IsStaged(stage, a) {
		return
	}

	a.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a.B != nil {
		StageBranch(stage, a.B)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _b := range a.Bs {
		StageBranch(stage, _b)
	}

}

func (stage *Stage) StageBranchB(b *B) {

	// check if instance is already staged
	if IsStaged(stage, b) {
		return
	}

	b.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *A:
		toT := CopyBranchA(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *B:
		toT := CopyBranchB(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchA(mapOrigCopy map[any]any, aFrom *A) (aTo *A) {

	// aFrom has already been copied
	if _aTo, ok := mapOrigCopy[aFrom]; ok {
		aTo = _aTo.(*A)
		return
	}

	aTo = new(A)
	mapOrigCopy[aFrom] = aTo
	aFrom.CopyBasicFields(aTo)

	//insertion point for the staging of instances referenced by pointers
	if aFrom.B != nil {
		aTo.B = CopyBranchB(mapOrigCopy, aFrom.B)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _b := range aFrom.Bs {
		aTo.Bs = append(aTo.Bs, CopyBranchB(mapOrigCopy, _b))
	}

	return
}

func CopyBranchB(mapOrigCopy map[any]any, bFrom *B) (bTo *B) {

	// bFrom has already been copied
	if _bTo, ok := mapOrigCopy[bFrom]; ok {
		bTo = _bTo.(*B)
		return
	}

	bTo = new(B)
	mapOrigCopy[bFrom] = bTo
	bFrom.CopyBasicFields(bTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *A:
		stage.UnstageBranchA(target)

	case *B:
		stage.UnstageBranchB(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchA(a *A) {

	// check if instance is already staged
	if !IsStaged(stage, a) {
		return
	}

	a.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a.B != nil {
		UnstageBranch(stage, a.B)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _b := range a.Bs {
		UnstageBranch(stage, _b)
	}

}

func (stage *Stage) UnstageBranchB(b *B) {

	// check if instance is already staged
	if !IsStaged(stage, b) {
		return
	}

	b.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a *A) GongDiff(stage *Stage, aOther *A) (diffs []string) {
	// insertion point for field diffs
	if a.Name != aOther.Name {
		diffs = append(diffs, a.GongMarshallField(stage, "Name"))
	}
	if a.NumberField != aOther.NumberField {
		diffs = append(diffs, a.GongMarshallField(stage, "NumberField"))
	}
	if (a.B == nil) != (aOther.B == nil) {
		diffs = append(diffs, "B")
	} else if a.B != nil && aOther.B != nil {
		if a.B != aOther.B {
			diffs = append(diffs, a.GongMarshallField(stage, "B"))
		}
	}
	BsDifferent := false
	if len(a.Bs) != len(aOther.Bs) {
		BsDifferent = true
	} else {
		for i := range a.Bs {
			if (a.Bs[i] == nil) != (aOther.Bs[i] == nil) {
				BsDifferent = true
				break
			} else if a.Bs[i] != nil && aOther.Bs[i] != nil {
				// this is a pointer comparaison
				if a.Bs[i] != aOther.Bs[i] {
					BsDifferent = true
					break
				}
			}
		}
	}
	if BsDifferent {
		ops := Diff(stage, a, aOther, "Bs", aOther.Bs, a.Bs)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (b *B) GongDiff(stage *Stage, bOther *B) (diffs []string) {
	// insertion point for field diffs
	if b.Name != bOther.Name {
		diffs = append(diffs, b.GongMarshallField(stage, "Name"))
	}

	return
}

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
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
			ops += fmt.Sprintf("\t%s.%s = slices.Delete( %s.%s, %d, %d)\n", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\t%s.%s = slices.Insert( %s.%s, %d, %s)\n",  a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
