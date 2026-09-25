// generated code - do not edit
package stool

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (stoolabstract *StoolAbstract) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StoolAbstracts[stoolabstract]

	return
}

func (stage *Stage) IsStagedStoolAbstract(stoolabstract *StoolAbstract) (ok bool) {

	return stoolabstract.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (stoolabstract *StoolAbstract) GongStageBranch(stage *Stage) {
	stage.StageBranchStoolAbstract(stoolabstract)
}

func (stage *Stage) StageBranchStoolAbstract(stoolabstract *StoolAbstract) {

	// check if instance is already staged
	if stage.IsStaged(stoolabstract) {
		return
	}

	stoolabstract.Stage(stage)

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
	case *StoolAbstract:
		toT := GongCopyBranchStoolAbstract(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchStoolAbstract(mapOrigCopy map[any]any, stoolabstractFrom *StoolAbstract) (stoolabstractTo *StoolAbstract) {

	// stoolabstractFrom has already been copied
	if _stoolabstractTo, ok := mapOrigCopy[stoolabstractFrom]; ok {
		stoolabstractTo = _stoolabstractTo.(*StoolAbstract)
		return
	}

	stoolabstractTo = new(StoolAbstract)
	mapOrigCopy[stoolabstractFrom] = stoolabstractTo
	stoolabstractFrom.GongCopyBasicFields(stoolabstractTo)

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
func (stoolabstract *StoolAbstract) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStoolAbstract(stoolabstract)
}

func (stage *Stage) UnstageBranchStoolAbstract(stoolabstract *StoolAbstract) {

	// check if instance is already staged
	if !stage.IsStaged(stoolabstract) {
		return
	}

	stoolabstract.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *StoolAbstract) GongReconstructPointersFromReferences(stage *Stage, instance *StoolAbstract) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *StoolAbstract) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stoolabstract *StoolAbstract) GongDiff(stage *Stage, stoolabstractOther *StoolAbstract) (diffs []string) {
	// insertion point for field diffs
	if stoolabstract.Name != stoolabstractOther.Name {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "Name"))
	}
	if stoolabstract.RadialRepetitions != stoolabstractOther.RadialRepetitions {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "RadialRepetitions"))
	}
	if stoolabstract.Transparency != stoolabstractOther.Transparency {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "Transparency"))
	}
	if stoolabstract.RelativeTubeDiameter != stoolabstractOther.RelativeTubeDiameter {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "RelativeTubeDiameter"))
	}
	if stoolabstract.RelativeHeight3DTorus != stoolabstractOther.RelativeHeight3DTorus {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "RelativeHeight3DTorus"))
	}
	if stoolabstract.StoolTorusVerticalScale != stoolabstractOther.StoolTorusVerticalScale {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "StoolTorusVerticalScale"))
	}
	if stoolabstract.RelativeHeight != stoolabstractOther.RelativeHeight {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "RelativeHeight"))
	}
	if stoolabstract.RelativeSeatThickness != stoolabstractOther.RelativeSeatThickness {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "RelativeSeatThickness"))
	}
	if stoolabstract.ProjectionAngle != stoolabstractOther.ProjectionAngle {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "ProjectionAngle"))
	}
	if stoolabstract.RelativeEyeSeparationCriteria != stoolabstractOther.RelativeEyeSeparationCriteria {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "RelativeEyeSeparationCriteria"))
	}
	if stoolabstract.RelativeEyeCornerControlVectorStrength != stoolabstractOther.RelativeEyeCornerControlVectorStrength {
		diffs = append(diffs, stoolabstract.GongMarshallField(stage, "RelativeEyeCornerControlVectorStrength"))
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
