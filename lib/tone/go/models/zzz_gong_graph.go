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
func (freqency *Freqency) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Freqencys[freqency]
	return ok
}

func (note *Note) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Notes[note]
	return ok
}

func (player *Player) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Players[player]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (freqency *Freqency) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(freqency) {
		return
	}

	freqency.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (note *Note) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _freqency := range note.Frequencies {
		stage.StageBranch(_freqency)
	}

}

func (player *Player) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(player) {
		return
	}

	player.Stage(stage)

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
	case *Freqency:
		toT := GongCopyBranchFreqency(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := GongCopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Player:
		toT := GongCopyBranchPlayer(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchFreqency(mapOrigCopy map[any]any, freqencyFrom *Freqency) (freqencyTo *Freqency) {
	var alreadyCopied bool
	freqencyTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, freqencyFrom)
	if alreadyCopied {
		return
	}
	freqencyFrom.GongCopyBasicFields(freqencyTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {
	var alreadyCopied bool
	noteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteFrom)
	if alreadyCopied {
		return
	}
	noteFrom.GongCopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _freqency := range noteFrom.Frequencies {
		noteTo.Frequencies = append(noteTo.Frequencies, GongCopyBranchFreqency(mapOrigCopy, _freqency))
	}

	return
}

func GongCopyBranchPlayer(mapOrigCopy map[any]any, playerFrom *Player) (playerTo *Player) {
	var alreadyCopied bool
	playerTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, playerFrom)
	if alreadyCopied {
		return
	}
	playerFrom.GongCopyBasicFields(playerTo)

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
func (freqency *Freqency) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(freqency) {
		return
	}

	freqency.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (note *Note) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _freqency := range note.Frequencies {
		stage.UnstageBranch(_freqency)
	}

}

func (player *Player) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(player) {
		return
	}

	player.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *Freqency) GongReconstructPointersFromReferences(stage *Stage, instance *Freqency) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Note) GongReconstructPointersFromReferences(stage *Stage, instance *Note) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Frequencies, stage.Freqencys_reference, instance.Frequencies)
}

func (reference *Player) GongReconstructPointersFromReferences(stage *Stage, instance *Player) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *Freqency) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Frequencies, stage.Freqencys_instance)
}

func (reference *Player) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (freqency *Freqency) GongDiff(stage *Stage, freqencyOther *Freqency) (diffs []string) {
	// insertion point for field diffs
	if freqency.Name != freqencyOther.Name {
		diffs = append(diffs, freqency.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (note *Note) GongDiff(stage *Stage, noteOther *Note) (diffs []string) {
	// insertion point for field diffs
	if note.Name != noteOther.Name {
		diffs = append(diffs, note.GongMarshallField(stage, "Name"))
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Frequencies", noteOther.Frequencies, note.Frequencies); ops != "" {
		diffs = append(diffs, ops)
	}
	if note.Start != noteOther.Start {
		diffs = append(diffs, note.GongMarshallField(stage, "Start"))
	}
	if note.Duration != noteOther.Duration {
		diffs = append(diffs, note.GongMarshallField(stage, "Duration"))
	}
	if note.Velocity != noteOther.Velocity {
		diffs = append(diffs, note.GongMarshallField(stage, "Velocity"))
	}
	if note.Info != noteOther.Info {
		diffs = append(diffs, note.GongMarshallField(stage, "Info"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (player *Player) GongDiff(stage *Stage, playerOther *Player) (diffs []string) {
	// insertion point for field diffs
	if player.Name != playerOther.Name {
		diffs = append(diffs, player.GongMarshallField(stage, "Name"))
	}
	if player.Status != playerOther.Status {
		diffs = append(diffs, player.GongMarshallField(stage, "Status"))
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
