package models

import "slices"

func (stager *Stager) enforceOrphansAbstractElement() (needCommit bool) {
	needCommit = reattachToLibraryRoots(
		stager,
		func() []*StateMachine {
			roots := make([]*StateMachine, 0)
			for _, library := range stager.stage.GetInstancesSorted[*Library]() {
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

	for _, note := range stager.stage.GetInstancesSorted[*Note]() {
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

	for _, state := range stager.stage.GetInstancesSorted[*State]() {
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

	if root := stager.GetRootLibrary(); root != nil {
		for _, role := range stager.stage.GetInstancesSorted[*Role]() {
			if !slices.Contains(root.Roles, role) {
				root.Roles = append(root.Roles, role)
				needCommit = true
			}
		}
		for _, messageType := range stager.stage.GetInstancesSorted[*MessageType]() {
			if !slices.Contains(root.MessageTypes, messageType) {
				root.MessageTypes = append(root.MessageTypes, messageType)
				needCommit = true
			}
		}
	}

	for _, object := range stager.stage.GetInstancesSorted[*Object]() {
		if object.State == nil {
			for _, sm := range stager.stage.GetInstancesSorted[*StateMachine]() {
				if sm.InitialState != nil {
					object.State = sm.InitialState
					needCommit = true
					break
				}
			}
		}
	}

	return
}
