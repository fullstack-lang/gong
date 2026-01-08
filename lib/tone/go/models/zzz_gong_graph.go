// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Freqency:
		ok = stage.IsStagedFreqency(target)

	case *Note:
		ok = stage.IsStagedNote(target)

	case *Player:
		ok = stage.IsStagedPlayer(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Freqency:
		ok = stage.IsStagedFreqency(target)

	case *Note:
		ok = stage.IsStagedNote(target)

	case *Player:
		ok = stage.IsStagedPlayer(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedFreqency(freqency *Freqency) (ok bool) {

	_, ok = stage.Freqencys[freqency]

	return
}

func (stage *Stage) IsStagedNote(note *Note) (ok bool) {

	_, ok = stage.Notes[note]

	return
}

func (stage *Stage) IsStagedPlayer(player *Player) (ok bool) {

	_, ok = stage.Players[player]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Freqency:
		stage.StageBranchFreqency(target)

	case *Note:
		stage.StageBranchNote(target)

	case *Player:
		stage.StageBranchPlayer(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchFreqency(freqency *Freqency) {

	// check if instance is already staged
	if IsStaged(stage, freqency) {
		return
	}

	freqency.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNote(note *Note) {

	// check if instance is already staged
	if IsStaged(stage, note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _freqency := range note.Frequencies {
		StageBranch(stage, _freqency)
	}

}

func (stage *Stage) StageBranchPlayer(player *Player) {

	// check if instance is already staged
	if IsStaged(stage, player) {
		return
	}

	player.Stage(stage)

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
	case *Freqency:
		toT := CopyBranchFreqency(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := CopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Player:
		toT := CopyBranchPlayer(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchFreqency(mapOrigCopy map[any]any, freqencyFrom *Freqency) (freqencyTo *Freqency) {

	// freqencyFrom has already been copied
	if _freqencyTo, ok := mapOrigCopy[freqencyFrom]; ok {
		freqencyTo = _freqencyTo.(*Freqency)
		return
	}

	freqencyTo = new(Freqency)
	mapOrigCopy[freqencyFrom] = freqencyTo
	freqencyFrom.CopyBasicFields(freqencyTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {

	// noteFrom has already been copied
	if _noteTo, ok := mapOrigCopy[noteFrom]; ok {
		noteTo = _noteTo.(*Note)
		return
	}

	noteTo = new(Note)
	mapOrigCopy[noteFrom] = noteTo
	noteFrom.CopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _freqency := range noteFrom.Frequencies {
		noteTo.Frequencies = append(noteTo.Frequencies, CopyBranchFreqency(mapOrigCopy, _freqency))
	}

	return
}

func CopyBranchPlayer(mapOrigCopy map[any]any, playerFrom *Player) (playerTo *Player) {

	// playerFrom has already been copied
	if _playerTo, ok := mapOrigCopy[playerFrom]; ok {
		playerTo = _playerTo.(*Player)
		return
	}

	playerTo = new(Player)
	mapOrigCopy[playerFrom] = playerTo
	playerFrom.CopyBasicFields(playerTo)

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
	case *Freqency:
		stage.UnstageBranchFreqency(target)

	case *Note:
		stage.UnstageBranchNote(target)

	case *Player:
		stage.UnstageBranchPlayer(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchFreqency(freqency *Freqency) {

	// check if instance is already staged
	if !IsStaged(stage, freqency) {
		return
	}

	freqency.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNote(note *Note) {

	// check if instance is already staged
	if !IsStaged(stage, note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _freqency := range note.Frequencies {
		UnstageBranch(stage, _freqency)
	}

}

func (stage *Stage) UnstageBranchPlayer(player *Player) {

	// check if instance is already staged
	if !IsStaged(stage, player) {
		return
	}

	player.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

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
	FrequenciesDifferent := false
	if len(note.Frequencies) != len(noteOther.Frequencies) {
		FrequenciesDifferent = true
	} else {
		for i := range note.Frequencies {
			if (note.Frequencies[i] == nil) != (noteOther.Frequencies[i] == nil) {
				FrequenciesDifferent = true
				break
			} else if note.Frequencies[i] != nil && noteOther.Frequencies[i] != nil {
				// this is a pointer comparaison
				if note.Frequencies[i] != noteOther.Frequencies[i] {
					FrequenciesDifferent = true
					break
				}
			}
		}
	}
	if FrequenciesDifferent {
		ops := Diff(stage, note, noteOther, "Frequencies", noteOther.Frequencies, note.Frequencies)
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
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, k+1)
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
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
