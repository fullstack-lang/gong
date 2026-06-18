package models

import (
)

func (stager *Stager) enforceOrphansAbstractElement() (needCommit bool) {
	needCommit = reattachToLibraryRoots(
		stager,
		func() []*Deliverable {
			roots := make([]*Deliverable, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
				roots = append(roots, library.RootDeliverables...)
			}
			return roots
		},
		func(product *Deliverable) {
			product.GetOwningLibrary().RootDeliverables = append(product.GetOwningLibrary().RootDeliverables, product)
		},
		func(product *Deliverable) []*Deliverable {
			return product.SubProducts
		},
	)

	needCommit = needCommit || reattachToLibraryRoots(
		stager,
		func() []*Concern {
			roots := make([]*Concern, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
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
		func() []*Stakeholder {
			roots := make([]*Stakeholder, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
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
		func() []*Diagram {
			roots := make([]*Diagram, 0)
			for _, library := range GetGongstrucsSorted[*Library](stager.stage) {
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


