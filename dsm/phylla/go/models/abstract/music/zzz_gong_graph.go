// generated code - do not edit
package music

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (musicabstract *MusicAbstract) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.MusicAbstracts[musicabstract]

	return
}

func (stage *Stage) IsStagedMusicAbstract(musicabstract *MusicAbstract) (ok bool) {

	return musicabstract.GongIsStaged(stage)
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
func (musicabstract *MusicAbstract) GongStageBranch(stage *Stage) {
	stage.StageBranchMusicAbstract(musicabstract)
}

func (stage *Stage) StageBranchMusicAbstract(musicabstract *MusicAbstract) {

	// check if instance is already staged
	if stage.IsStaged(musicabstract) {
		return
	}

	musicabstract.Stage(stage)

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
	case *MusicAbstract:
		toT := GongCopyBranchMusicAbstract(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchMusicAbstract(mapOrigCopy map[any]any, musicabstractFrom *MusicAbstract) (musicabstractTo *MusicAbstract) {

	// musicabstractFrom has already been copied
	if _musicabstractTo, ok := mapOrigCopy[musicabstractFrom]; ok {
		musicabstractTo = _musicabstractTo.(*MusicAbstract)
		return
	}

	musicabstractTo = new(MusicAbstract)
	mapOrigCopy[musicabstractFrom] = musicabstractTo
	musicabstractFrom.GongCopyBasicFields(musicabstractTo)

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
func (musicabstract *MusicAbstract) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchMusicAbstract(musicabstract)
}

func (stage *Stage) UnstageBranchMusicAbstract(musicabstract *MusicAbstract) {

	// check if instance is already staged
	if !stage.IsStaged(musicabstract) {
		return
	}

	musicabstract.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *MusicAbstract) GongReconstructPointersFromReferences(stage *Stage, instance *MusicAbstract) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *MusicAbstract) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (musicabstract *MusicAbstract) GongDiff(stage *Stage, musicabstractOther *MusicAbstract) (diffs []string) {
	// insertion point for field diffs
	if musicabstract.Name != musicabstractOther.Name {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "Name"))
	}
	if musicabstract.IsChecked != musicabstractOther.IsChecked {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "IsChecked"))
	}
	if musicabstract.PitchHeight != musicabstractOther.PitchHeight {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "PitchHeight"))
	}
	if musicabstract.NbOfBeatsInTheme != musicabstractOther.NbOfBeatsInTheme {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "NbOfBeatsInTheme"))
	}
	if musicabstract.BeatsPerSecond != musicabstractOther.BeatsPerSecond {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "BeatsPerSecond"))
	}
	if musicabstract.FirstVoiceShiftX != musicabstractOther.FirstVoiceShiftX {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "FirstVoiceShiftX"))
	}
	if musicabstract.FirstVoiceShiftY != musicabstractOther.FirstVoiceShiftY {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "FirstVoiceShiftY"))
	}
	if musicabstract.PitchDifference != musicabstractOther.PitchDifference {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "PitchDifference"))
	}
	if musicabstract.Level != musicabstractOther.Level {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "Level"))
	}
	if musicabstract.ActualBeatsTemporalShift != musicabstractOther.ActualBeatsTemporalShift {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ActualBeatsTemporalShift"))
	}
	if musicabstract.IsMinor != musicabstractOther.IsMinor {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "IsMinor"))
	}
	if musicabstract.ThemeBinaryEncoding != musicabstractOther.ThemeBinaryEncoding {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ThemeBinaryEncoding"))
	}
	if musicabstract.BezierControlLengthRatio != musicabstractOther.BezierControlLengthRatio {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "BezierControlLengthRatio"))
	}
	if musicabstract.NbPitchLines != musicabstractOther.NbPitchLines {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "NbPitchLines"))
	}
	if musicabstract.NbBeatLines != musicabstractOther.NbBeatLines {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "NbBeatLines"))
	}
	if musicabstract.OriginX != musicabstractOther.OriginX {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "OriginX"))
	}
	if musicabstract.OriginY != musicabstractOther.OriginY {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "OriginY"))
	}
	if musicabstract.ScoreScale != musicabstractOther.ScoreScale {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ScoreScale"))
	}
	if musicabstract.ShowFirstVoice != musicabstractOther.ShowFirstVoice {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ShowFirstVoice"))
	}
	if musicabstract.ShowFirstVoiceShiftRight != musicabstractOther.ShowFirstVoiceShiftRight {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ShowFirstVoiceShiftRight"))
	}
	if musicabstract.ShowSecondVoice != musicabstractOther.ShowSecondVoice {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ShowSecondVoice"))
	}
	if musicabstract.ShowSecondVoiceShiftRight != musicabstractOther.ShowSecondVoiceShiftRight {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ShowSecondVoiceShiftRight"))
	}
	if musicabstract.ShowFirstVoiceNotes != musicabstractOther.ShowFirstVoiceNotes {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ShowFirstVoiceNotes"))
	}
	if musicabstract.ShowFirstVoiceNotesShiftRight != musicabstractOther.ShowFirstVoiceNotesShiftRight {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ShowFirstVoiceNotesShiftRight"))
	}
	if musicabstract.ShowSecondVoiceNotes != musicabstractOther.ShowSecondVoiceNotes {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ShowSecondVoiceNotes"))
	}
	if musicabstract.ShowSecondVoiceNotesShiftRight != musicabstractOther.ShowSecondVoiceNotesShiftRight {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "ShowSecondVoiceNotesShiftRight"))
	}
	if musicabstract.IsComposerNodeExpanded != musicabstractOther.IsComposerNodeExpanded {
		diffs = append(diffs, musicabstract.GongMarshallField(stage, "IsComposerNodeExpanded"))
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
