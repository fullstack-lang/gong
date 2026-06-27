package models

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
