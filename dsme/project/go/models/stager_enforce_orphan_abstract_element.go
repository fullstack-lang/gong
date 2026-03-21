package models

import (
	"fmt"
	"time"
)

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
			product.GetOwnlingLibrary().RootProducts = append(product.GetOwnlingLibrary().RootProducts, product)
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
			task.GetOwnlingLibrary().RootTasks = append(task.GetOwnlingLibrary().RootTasks, task)
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
			note.GetOwnlingLibrary().Notes = append(note.GetOwnlingLibrary().Notes, note)
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
			resource.GetOwnlingLibrary().RootResources = append(resource.GetOwnlingLibrary().RootResources, resource)
		},
		func(resource *Resource) []*Resource {
			return resource.SubResources
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
			diagram.GetOwnlingLibrary().Diagrams = append(diagram.GetOwnlingLibrary().Diagrams, diagram)
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
			stager.rootLibrary.SubLibraries = append(stager.rootLibrary.SubLibraries, library)
			library.SetOwningLibrary(stager.rootLibrary)
		},
		func(library *Library) []*Library {
			return []*Library{}
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
	appendToRoot func(T),
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
			// attach to root, only if it is not the root library (which is the only one without an owning library)

			if any(object) == any(stager.rootLibrary) {
				continue
			}
			appendToRoot(object)
			needCommit = true
			stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Orphan %s %s, was reattached to the root of library %s",
				object.GongGetGongstructName(), object.GetName(), object.GetOwnlingLibrary().GetName()))
		}
	}

	return
}
