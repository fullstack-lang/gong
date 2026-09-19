package models

func (stager *Stager) enforceOrphansAbstractElement() (needCommit bool) {
	needCommit = reattachToLibraryRoots(
		stager,
		func() []*Analysis {
			roots := make([]*Analysis, 0)
			for _, library := range stager.stage.GetInstancesSorted[*Library]() {
				roots = append(roots, library.Analyses...)
			}
			return roots
		},
		func(analysis *Analysis) {
			stager.getRootLibrary().Analyses = append(stager.getRootLibrary().Analyses, analysis)
			analysis.SetOwningLibrary(stager.getRootLibrary())
		},
		func(analysis *Analysis) []*Analysis {
			return nil
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
