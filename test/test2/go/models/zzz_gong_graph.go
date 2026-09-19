// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (a *A) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.As[a]

	return
}

func (stage *Stage) IsStagedA(a *A) (ok bool) {

	return a.GongIsStaged(stage)
}

func (b *B) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Bs[b]

	return
}

func (stage *Stage) IsStagedB(b *B) (ok bool) {

	return b.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// StageBranch is a backward-compatible package-level forwarder.
func StageBranch(stage *Stage, instance GongstructIF) {
	stage.StageBranch(instance)
}

// insertion point for stage branch per struct
func (a *A) GongStageBranch(stage *Stage) {
	stage.StageBranchA(a)
}

func (stage *Stage) StageBranchA(a *A) {

	// check if instance is already staged
	if stage.IsStaged(a) {
		return
	}

	a.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a.B != nil {
		stage.StageBranch(a.B)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _b := range a.Bs {
		stage.StageBranch(_b)
	}

}

func (b *B) GongStageBranch(stage *Stage) {
	stage.StageBranchB(b)
}

func (stage *Stage) StageBranchB(b *B) {

	// check if instance is already staged
	if stage.IsStaged(b) {
		return
	}

	b.Stage(stage)

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
	case *A:
		toT := GongCopyBranchA(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *B:
		toT := GongCopyBranchB(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchA(mapOrigCopy map[any]any, aFrom *A) (aTo *A) {

	// aFrom has already been copied
	if _aTo, ok := mapOrigCopy[aFrom]; ok {
		aTo = _aTo.(*A)
		return
	}

	aTo = new(A)
	mapOrigCopy[aFrom] = aTo
	aFrom.GongCopyBasicFields(aTo)

	//insertion point for the staging of instances referenced by pointers
	if aFrom.B != nil {
		aTo.B = GongCopyBranchB(mapOrigCopy, aFrom.B)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _b := range aFrom.Bs {
		aTo.Bs = append(aTo.Bs, GongCopyBranchB(mapOrigCopy, _b))
	}

	return
}

func GongCopyBranchB(mapOrigCopy map[any]any, bFrom *B) (bTo *B) {

	// bFrom has already been copied
	if _bTo, ok := mapOrigCopy[bFrom]; ok {
		bTo = _bTo.(*B)
		return
	}

	bTo = new(B)
	mapOrigCopy[bFrom] = bTo
	bFrom.GongCopyBasicFields(bTo)

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

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (a *A) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA(a)
}

func (stage *Stage) UnstageBranchA(a *A) {

	// check if instance is already staged
	if !stage.IsStaged(a) {
		return
	}

	a.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a.B != nil {
		stage.UnstageBranch(a.B)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _b := range a.Bs {
		stage.UnstageBranch(_b)
	}

}

func (b *B) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchB(b)
}

func (stage *Stage) UnstageBranchB(b *B) {

	// check if instance is already staged
	if !stage.IsStaged(b) {
		return
	}

	b.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *A) GongReconstructPointersFromReferences(stage *Stage, instance *A) {
	// insertion point for pointers field
	if instance.B != nil {
		reference.B = stage.Bs_reference[instance.B]
	}
	// insertion point for slice of pointers field
	reference.Bs = reference.Bs[:0]
	for _, _b := range instance.Bs {
		reference.Bs = append(reference.Bs, stage.Bs_reference[_b])
	}
}

func (reference *B) GongReconstructPointersFromReferences(stage *Stage, instance *B) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *A) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.B; _reference != nil {
		reference.B = nil
		if _instance, ok := stage.Bs_instance[_reference]; ok {
			reference.B = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Bs []*B
	for _, _reference := range reference.Bs {
		if _instance, ok := stage.Bs_instance[_reference]; ok {
			_Bs = append(_Bs, _instance)
		}
	}
	reference.Bs = _Bs
}

func (reference *B) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
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
		diffs = append(diffs, a.GongMarshallField(stage, "B"))
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
		ops := stage.Diff(
			a,
			"Bs",
			len(aOther.Bs),
			len(a.Bs),
			func(i, j int) bool {
				return aOther.Bs[i] == a.Bs[j]
			},
			func(j int) string {
				return a.Bs[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if a.Foo != aOther.Foo {
		diffs = append(diffs, a.GongMarshallField(stage, "Foo"))
	}
	if a.Bar != aOther.Bar {
		diffs = append(diffs, a.GongMarshallField(stage, "Bar"))
	}
	if a.Zorgh != aOther.Zorgh {
		diffs = append(diffs, a.GongMarshallField(stage, "Zorgh"))
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
