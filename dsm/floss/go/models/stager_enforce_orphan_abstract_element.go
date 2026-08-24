package models

func (stager *Stager) enforceOrphansAbstractElement() (needCommit bool) {
	needCommit = reattachToLibraryRoots(
		stager,
		func() []*System {
			roots := make([]*System, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
				roots = append(roots, library.RootSystems...)
			}
			return roots
		},
		func(system *System) {
			system.GetOwningLibrary().RootSystems = append(system.GetOwningLibrary().RootSystems, system)
		},
		func(system *System) []*System {
			return system.SubSystems
		},
	)

	if reattachToLibraryRoots(
		stager,
		func() []*Complexity {
			roots := make([]*Complexity, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
				roots = append(roots, library.RootComplexitys...)
			}
			return roots
		},
		func(complexity *Complexity) {
			complexity.GetOwningLibrary().RootComplexitys = append(complexity.GetOwningLibrary().RootComplexitys, complexity)
		},
		func(complexity *Complexity) []*Complexity {
			return []*Complexity{}
		},
	) {
		needCommit = true
	}

	if reattachToLibraryRoots(
		stager,
		func() []*Performance {
			roots := make([]*Performance, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
				roots = append(roots, library.RootPerformances...)
			}
			return roots
		},
		func(performance *Performance) {
			performance.GetOwningLibrary().RootPerformances = append(performance.GetOwningLibrary().RootPerformances, performance)
		},
		func(performance *Performance) []*Performance {
			return []*Performance{}
		},
	) {
		needCommit = true
	}

	if reattachToLibraryRoots(
		stager,
		func() []*Effort {
			roots := make([]*Effort, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
				roots = append(roots, library.RootEfforts...)
			}
			return roots
		},
		func(effort *Effort) {
			effort.GetOwningLibrary().RootEfforts = append(effort.GetOwningLibrary().RootEfforts, effort)
		},
		func(effort *Effort) []*Effort {
			return []*Effort{}
		},
	) {
		needCommit = true
	}

	if reattachToLibraryRoots(
		stager,
		func() []*CompareAnalysis {
			roots := make([]*CompareAnalysis, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
				roots = append(roots, library.RootCompareAnalysis...)
			}
			return roots
		},
		func(compareAnalysis *CompareAnalysis) {
			compareAnalysis.GetOwningLibrary().RootCompareAnalysis = append(compareAnalysis.GetOwningLibrary().RootCompareAnalysis, compareAnalysis)
		},
		func(compareAnalysis *CompareAnalysis) []*CompareAnalysis {
			return []*CompareAnalysis{}
		},
	) {
		needCommit = true
	}

	if reattachToLibraryRoots(
		stager,
		func() []*Note {
			roots := make([]*Note, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
				roots = append(roots, library.RootNotes...)
			}
			return roots
		},
		func(note *Note) {
			note.GetOwningLibrary().RootNotes = append(note.GetOwningLibrary().RootNotes, note)
		},
		func(note *Note) []*Note {
			return []*Note{}
		},
	) {
		needCommit = true
	}

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
