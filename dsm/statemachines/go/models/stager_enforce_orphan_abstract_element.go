package models

import "slices"

func (stager *Stager) enforceOrphansAbstractElement() (needCommit bool) {
	needCommit = reattachToLibraryRoots(
		stager,
		func() []*StateMachine {
			roots := make([]*StateMachine, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
				roots = append(roots, library.RootStateMachines...)
			}
			return roots
		},
		func(stateMachine *StateMachine) {
			stateMachine.GetOwningLibrary().RootStateMachines = append(stateMachine.GetOwningLibrary().RootStateMachines, stateMachine)
		},
		func(stateMachine *StateMachine) []*StateMachine {
			return []*StateMachine{}
		},
	)

	for _, note := range GetGongstrucsSorted[*Note](stager.stage) {
		if note.State == nil {
			note.Unstage(stager.stage)
			needCommit = true
			continue
		}
		if !slices.Contains(note.State.Notes, note) {
			note.State.Notes = append(note.State.Notes, note)
			needCommit = true
		}
	}

	for _, state := range GetGongstrucsSorted[*State](stager.stage) {
		var validNotes []*Note
		for _, note := range state.Notes {
			if note != nil && note.State == state {
				validNotes = append(validNotes, note)
			} else {
				needCommit = true
			}
		}
		state.Notes = validNotes
	}

	if reattachToLibraryRoots(
		stager,
		func() []*Library {
			return stager.GetRootLibrary().SubLibraries
		},
		func(library *Library) {
			if library != stager.GetRootLibrary() {
				stager.GetRootLibrary().SubLibraries = append(stager.GetRootLibrary().SubLibraries, library)
				library.SetOwningLibrary(stager.GetRootLibrary())
			}
		},
		func(library *Library) []*Library {
			return library.SubLibraries
		},
	) {
		needCommit = true
	}
	return
}
