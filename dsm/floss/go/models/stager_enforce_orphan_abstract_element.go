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
