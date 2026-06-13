package models

func (stager *Stager) enforceOrphansAbstractElement() (needCommit bool) {
	needCommit = reattachToLibraryRoots(
		stager,
		func() []*Product {
			roots := make([]*Product, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
				roots = append(roots, library.RootProducts...)
			}
			return roots
		},
		func(product *Product) {
			product.GetOwningLibrary().RootProducts = append(product.GetOwningLibrary().RootProducts, product)
		},
		func(product *Product) []*Product {
			return product.SubProducts
		},
	)

	needCommit = needCommit || reattachToLibraryRoots(
		stager,
		func() []*Task {
			roots := make([]*Task, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
				roots = append(roots, library.RootTasks...)
			}
			return roots
		},
		func(task *Task) {
			task.GetOwningLibrary().RootTasks = append(task.GetOwningLibrary().RootTasks, task)
		},
		func(task *Task) []*Task {
			return task.SubTasks
		},
	)

	needCommit = needCommit || reattachToLibraryRoots(
		stager,
		func() []*Note {
			roots := make([]*Note, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
				roots = append(roots, library.Notes...)
			}
			return roots
		},
		func(note *Note) {
			note.GetOwningLibrary().Notes = append(note.GetOwningLibrary().Notes, note)
		},
		func(note *Note) []*Note {
			return []*Note{}
		},
	)

	needCommit = needCommit || reattachToLibraryRoots(
		stager,
		func() []*Resource {
			roots := make([]*Resource, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
				roots = append(roots, library.RootResources...)
			}
			return roots
		},
		func(resource *Resource) {
			resource.GetOwningLibrary().RootResources = append(resource.GetOwningLibrary().RootResources, resource)
		},
		func(resource *Resource) []*Resource {
			return resource.SubResources
		},
	)

	needCommit = needCommit || reattachToLibraryRoots(
		stager,
		func() []*DiagramHierarchy {
			roots := make([]*DiagramHierarchy, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
				roots = append(roots, library.Diagrams...)
			}
			return roots
		},
		func(diagramHierarchy *DiagramHierarchy) {
			diagramHierarchy.GetOwningLibrary().Diagrams = append(diagramHierarchy.GetOwningLibrary().Diagrams, diagramHierarchy)
		},
		func(diagramHierarchy *DiagramHierarchy) []*DiagramHierarchy {
			return []*DiagramHierarchy{}
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
