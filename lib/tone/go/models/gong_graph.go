// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

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
func (stage *StageStruct) IsStagedFreqency(freqency *Freqency) (ok bool) {

	_, ok = stage.Freqencys[freqency]

	return
}

func (stage *StageStruct) IsStagedNote(note *Note) (ok bool) {

	_, ok = stage.Notes[note]

	return
}

func (stage *StageStruct) IsStagedPlayer(player *Player) (ok bool) {

	_, ok = stage.Players[player]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

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
func (stage *StageStruct) StageBranchFreqency(freqency *Freqency) {

	// check if instance is already staged
	if IsStaged(stage, freqency) {
		return
	}

	freqency.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchNote(note *Note) {

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

func (stage *StageStruct) StageBranchPlayer(player *Player) {

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
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

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
func (stage *StageStruct) UnstageBranchFreqency(freqency *Freqency) {

	// check if instance is already staged
	if !IsStaged(stage, freqency) {
		return
	}

	freqency.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchNote(note *Note) {

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

func (stage *StageStruct) UnstageBranchPlayer(player *Player) {

	// check if instance is already staged
	if !IsStaged(stage, player) {
		return
	}

	player.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}
