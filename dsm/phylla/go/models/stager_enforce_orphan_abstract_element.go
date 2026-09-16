package models

import "fmt"

func (stager *Stager) enforceOrphansAbstractElement() (needCommit bool) {
	needCommit = reattachToLibraryRoots(
		stager,
		func() []*PlantAbstract {
			roots := make([]*PlantAbstract, 0)
			for _, library := range stager.stage.GetInstancesSorted[*Library]() {
				roots = append(roots, library.Plants...)
			}
			return roots
		},
		func(plant *PlantAbstract) {
			owningLib := plant.GetOwningLibrary()
			owningLib.Plants = append(owningLib.Plants, plant)
			stager.logAndNotify(fmt.Sprintf("Reattached orphan plant %s to library %s", plant.Name, owningLib.Name))
		},
		func(plant *PlantAbstract) []*PlantAbstract {
			return []*PlantAbstract{}
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
				stager.logAndNotify(fmt.Sprintf("Reattached orphan sub-library %s to root library", library.Name))
			}
		},
		func(library *Library) []*Library {
			return library.SubLibraries
		},
	)

	return
}
