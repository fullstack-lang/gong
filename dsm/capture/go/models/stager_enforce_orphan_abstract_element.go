package models

import (
	"fmt"
	"time"
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
	)

	return
}

func reattachToLibraryRoots[T interface {
	AbstractType
	LibraryOwnedType
	comparable
}](
	stager *Stager,
	getRoots func() []T,
	attachDirectlyToLibraryRoot func(T),
	getChildren func(T) []T,
) (needCommit bool) {
	// 1. Find all reachable nodes
	reachable := make(map[T]struct{})
	var collectReachable func(node T)
	collectReachable = func(node T) {
		var zero T
		if node == zero {
			return
		}
		if _, ok := reachable[node]; ok {
			return // already visited
		}
		reachable[node] = struct{}{}

		for _, child := range getChildren(node) {
			collectReachable(child)
		}
	}

	for _, root := range getRoots() {
		collectReachable(root)
	}

	// 2. Find all nodes and delete them
	for _, object := range GetGongstrucsSorted[T](stager.stage) {
		if _, ok := reachable[object]; !ok {
			if object != any(stager.rootLibrary) {
				object.UnstageVoid(stager.stage)
				needCommit = true
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Orphan \"%s\" of type \"%s\" was deleted",
						object.GetName(), object.GongGetGongstructName()),
				)
			}
		}
	}

	return
}
