package models

func (stager *Stager) enforceOrphansAbstractElement() (needCommit bool) {
	needCommit = reattachToLibraryRoots(
		stager,
		func() []*Plant {
			roots := make([]*Plant, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
				roots = append(roots, library.RootPlants...)
			}
			return roots
		},
		func(plant *Plant) {
			plant.GetOwningLibrary().RootPlants = append(plant.GetOwningLibrary().RootPlants, plant)
		},
		func(plant *Plant) []*Plant {
			return []*Plant{}
		},
	)

	needCommit = needCommit || reattachToLibraryRoots(
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
	)

	return
}
