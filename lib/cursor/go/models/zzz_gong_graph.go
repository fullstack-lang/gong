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
func (cursor *Cursor) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Cursors[cursor]

	return
}

func (stage *Stage) IsStagedCursor(cursor *Cursor) (ok bool) {

	return cursor.GongIsStaged(stage)
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
func (cursor *Cursor) GongStageBranch(stage *Stage) {
	stage.StageBranchCursor(cursor)
}

func (stage *Stage) StageBranchCursor(cursor *Cursor) {

	// check if instance is already staged
	if stage.IsStaged(cursor) {
		return
	}

	cursor.Stage(stage)

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
	case *Cursor:
		toT := GongCopyBranchCursor(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchCursor(mapOrigCopy map[any]any, cursorFrom *Cursor) (cursorTo *Cursor) {

	// cursorFrom has already been copied
	if _cursorTo, ok := mapOrigCopy[cursorFrom]; ok {
		cursorTo = _cursorTo.(*Cursor)
		return
	}

	cursorTo = new(Cursor)
	mapOrigCopy[cursorFrom] = cursorTo
	cursorFrom.GongCopyBasicFields(cursorTo)

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
func (cursor *Cursor) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchCursor(cursor)
}

func (stage *Stage) UnstageBranchCursor(cursor *Cursor) {

	// check if instance is already staged
	if !stage.IsStaged(cursor) {
		return
	}

	cursor.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Cursor) GongReconstructPointersFromReferences(stage *Stage, instance *Cursor) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Cursor) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (cursor *Cursor) GongDiff(stage *Stage, cursorOther *Cursor) (diffs []string) {
	// insertion point for field diffs
	if cursor.Name != cursorOther.Name {
		diffs = append(diffs, cursor.GongMarshallField(stage, "Name"))
	}
	if cursor.StartX != cursorOther.StartX {
		diffs = append(diffs, cursor.GongMarshallField(stage, "StartX"))
	}
	if cursor.EndX != cursorOther.EndX {
		diffs = append(diffs, cursor.GongMarshallField(stage, "EndX"))
	}
	if cursor.Y1 != cursorOther.Y1 {
		diffs = append(diffs, cursor.GongMarshallField(stage, "Y1"))
	}
	if cursor.Y2 != cursorOther.Y2 {
		diffs = append(diffs, cursor.GongMarshallField(stage, "Y2"))
	}
	if cursor.DurationSeconds != cursorOther.DurationSeconds {
		diffs = append(diffs, cursor.GongMarshallField(stage, "DurationSeconds"))
	}
	if cursor.Color != cursorOther.Color {
		diffs = append(diffs, cursor.GongMarshallField(stage, "Color"))
	}
	if cursor.FillOpacity != cursorOther.FillOpacity {
		diffs = append(diffs, cursor.GongMarshallField(stage, "FillOpacity"))
	}
	if cursor.Stroke != cursorOther.Stroke {
		diffs = append(diffs, cursor.GongMarshallField(stage, "Stroke"))
	}
	if cursor.StrokeOpacity != cursorOther.StrokeOpacity {
		diffs = append(diffs, cursor.GongMarshallField(stage, "StrokeOpacity"))
	}
	if cursor.StrokeWidth != cursorOther.StrokeWidth {
		diffs = append(diffs, cursor.GongMarshallField(stage, "StrokeWidth"))
	}
	if cursor.StrokeDashArray != cursorOther.StrokeDashArray {
		diffs = append(diffs, cursor.GongMarshallField(stage, "StrokeDashArray"))
	}
	if cursor.StrokeDashArrayWhenSelected != cursorOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, cursor.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if cursor.Transform != cursorOther.Transform {
		diffs = append(diffs, cursor.GongMarshallField(stage, "Transform"))
	}
	if cursor.IsPlaying != cursorOther.IsPlaying {
		diffs = append(diffs, cursor.GongMarshallField(stage, "IsPlaying"))
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

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if equal(i, j) {
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
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := 0; k < n; k++ {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}
