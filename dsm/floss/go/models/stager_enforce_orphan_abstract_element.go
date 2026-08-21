package models

func (stager *Stager) enforceOrphansAbstractElement() (needCommit bool) {
	needCommit = reattachToLibraryRoots(
		stager,
		func() []*System {
			roots := make([]*System, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
				roots = append(roots, library.RootSystemes...)
			}
			return roots
		},
		func(system *System) {
			system.GetOwningLibrary().RootSystemes = append(system.GetOwningLibrary().RootSystemes, system)
		},
		func(system *System) []*System {
			return system.SubSystemes
		},
	)

	if reattachToLibraryRoots(
		stager,
		func() []*Library {
			return stager.getRootLibrary().SubLibraries
		},
		func(library *Library) {
			// attach to root, only if it is not the root library
			// (which is the only one without an owning library)
			if library != stager.getRootLibrary() {
				stager.getRootLibrary().SubLibraries = append(stager.getRootLibrary().SubLibraries, library)
				library.SetOwningLibrary(stager.getRootLibrary())
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
