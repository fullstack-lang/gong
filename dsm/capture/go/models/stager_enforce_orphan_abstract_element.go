package models

func (stager *Stager) enforceOrphansAbstractElement() (needCommit bool) {
	needCommit = reattachToLibraryRoots(
		stager,
		func() []*Deliverable {
			roots := make([]*Deliverable, 0)
			for _, library := range stager.stage.GetInstancesSorted[*Library]() {
				roots = append(roots, library.RootDeliverables...)
			}
			return roots
		},
		func(deliverable *Deliverable) {
			deliverable.GetOwningLibrary().RootDeliverables = append(deliverable.GetOwningLibrary().RootDeliverables, deliverable)
		},
		func(deliverable *Deliverable) []*Deliverable {
			return deliverable.SubDeliverables
		},
	)

	needCommit = needCommit || reattachToLibraryRoots(
		stager,
		func() []*Concern {
			roots := make([]*Concern, 0)
			for _, library := range stager.stage.GetInstancesSorted[*Library]() {
				roots = append(roots, library.RootConcerns...)
			}
			return roots
		},
		func(task *Concern) {
			task.GetOwningLibrary().RootConcerns = append(task.GetOwningLibrary().RootConcerns, task)
		},
		func(task *Concern) []*Concern {
			return task.SubConcerns
		},
	)

	needCommit = needCommit || reattachToLibraryRoots(
		stager,
		func() []*Note {
			roots := make([]*Note, 0)
			for _, library := range stager.stage.GetInstancesSorted[*Library]() {
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
		func() []*Stakeholder {
			roots := make([]*Stakeholder, 0)
			for _, library := range stager.stage.GetInstancesSorted[*Library]() {
				roots = append(roots, library.RootStakeholders...)
			}
			return roots
		},
		func(resource *Stakeholder) {
			resource.GetOwningLibrary().RootStakeholders = append(resource.GetOwningLibrary().RootStakeholders, resource)
		},
		func(resource *Stakeholder) []*Stakeholder {
			return resource.SubStakeholders
		},
	)

	needCommit = needCommit || reattachToLibraryRoots(
		stager,
		func() []*Requirement {
			roots := make([]*Requirement, 0)
			for _, library := range stager.stage.GetInstancesSorted[*Library]() {
				roots = append(roots, library.RootRequirements...)
			}
			return roots
		},
		func(req *Requirement) {
			req.GetOwningLibrary().RootRequirements = append(req.GetOwningLibrary().RootRequirements, req)
		},
		func(req *Requirement) []*Requirement {
			return []*Requirement{}
		},
	)

	needCommit = needCommit || reattachToLibraryRoots(
		stager,
		func() []*View {
			roots := make([]*View, 0)
			for _, library := range stager.stage.GetInstancesSorted[*Library]() {
				roots = append(roots, library.RootViews...)
			}
			return roots
		},
		func(view *View) {
			view.GetOwningLibrary().RootViews = append(view.GetOwningLibrary().RootViews, view)
		},
		func(view *View) []*View {
			return []*View{}
		},
	)

	needCommit = needCommit || reattachToLibraryRoots(
		stager,
		func() []*Concept {
			roots := make([]*Concept, 0)
			for _, library := range stager.stage.GetInstancesSorted[*Library]() {
				roots = append(roots, library.RootConcepts...)
			}
			return roots
		},
		func(concept *Concept) {
			concept.GetOwningLibrary().RootConcepts = append(concept.GetOwningLibrary().RootConcepts, concept)
		},
		func(concept *Concept) []*Concept {
			return []*Concept{}
		},
	)

	needCommit = needCommit || reattachToLibraryRoots(
		stager,
		func() []*Diagram {
			roots := make([]*Diagram, 0)
			for _, library := range stager.stage.GetInstancesSorted[*Library]() {
				roots = append(roots, library.Diagrams...)
			}
			return roots
		},
		func(diagram *Diagram) {
			diagram.GetOwningLibrary().Diagrams = append(diagram.GetOwningLibrary().Diagrams, diagram)
		},
		func(diagram *Diagram) []*Diagram {
			return []*Diagram{}
		},
	)

	needCommit = needCommit || reattachToLibraryRoots(
		stager,
		func() []*Library {
			return stager.GetRootLibrary().SubLibraries
		},
		func(library *Library) {
			// attach to root, only if it is not the root library
			// (which is the only one without an owning library)
			if library != stager.GetRootLibrary() {
				stager.GetRootLibrary().SubLibraries = append(stager.GetRootLibrary().SubLibraries, library)
				library.SetOwningLibrary(stager.GetRootLibrary())
			}
		},
		func(library *Library) []*Library {
			return library.SubLibraries
		},
	)

	return
}
