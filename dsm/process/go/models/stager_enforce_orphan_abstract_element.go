package models

func (stager *Stager) enforceOrphansAbstractElement() (needCommit bool) {
	needCommit = reattachToLibraryRoots(
		stager,
		func() []*Process {
			roots := make([]*Process, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
				roots = append(roots, library.RootProcesses...)
			}
			return roots
		},
		func(process *Process) {
			process.GetOwningLibrary().RootProcesses = append(process.GetOwningLibrary().RootProcesses, process)
		},
		func(process *Process) []*Process {
			return process.SubProcesses
		},
	)

	if reattachToLibraryRoots(
		stager,
		func() []*Library {
			return stager.rootLibrary.SubLibraries
		},
		func(library *Library) {
			// attach to root, only if it is not the root library
			// (which is the only one without an owning library)
			if library != stager.rootLibrary {
				stager.rootLibrary.SubLibraries = append(stager.rootLibrary.SubLibraries, library)
				library.SetOwningLibrary(stager.rootLibrary)
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
