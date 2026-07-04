// generated code (do not edit)
package models

import (
	"fmt"
	"log"
	"maps"
	"slices"
	"time"
)

// Append is a generic helper that appends an item to a slice via a pointer
func Append[T any](slice *[]T, item T) {
	*slice = append(*slice, item)
}

func collectProjectElements[T PointerToGongstruct](
	rootElements []T,
	getChildren func(T) []T,
) []T {
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

	for _, root := range rootElements {
		collectReachable(root)
	}

	// 2. Convert map to slice
	return slices.Collect(maps.Keys(reachable))
}

type LibraryAbstractFields struct {
	owningLibrary *Library
}

type LibraryOwnedType interface {
	GetOwningLibrary() *Library
	SetOwningLibrary(library *Library)
}

func (r *LibraryAbstractFields) GetOwningLibrary() *Library {
	return r.owningLibrary
}

func (r *LibraryAbstractFields) SetOwningLibrary(library *Library) {
	r.owningLibrary = library
}

func (stager *Stager) enforceThereIsARootLibrary() (needCommit bool) {
	stage := stager.stage
	libraries := GetStructInstancesByOrderAuto[*Library](stage)
	if len(libraries) == 0 {
		rootLibrary := (&Library{Name: "", IsRootLibrary: true}).Stage(stage)
		if stager.probeForm != nil {
			stager.probeForm.AddNotification(time.Now(),
				"Created root library")
		}
		rootLibrary.Stage(stage)
		needCommit = true
		return
	}

	var rootLibrary *Library
	for _, library := range libraries {
		if library.IsRootLibrary {
			if rootLibrary == nil {
				rootLibrary = library
			} else {
				library.IsRootLibrary = false
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						"Detected an extra root library, unsetting it is a root library: "+library.GetName())
				}
				needCommit = true
			}
		}
	}

	if rootLibrary == nil {
		// we set the first library that has no parent as the root library
		hasParent := make(map[*Library]bool)
		for _, library := range libraries {
			for _, sub := range library.SubLibraries {
				hasParent[sub] = true
			}
		}

		var rootCandidates []*Library
		for _, library := range libraries {
			if !hasParent[library] {
				rootCandidates = append(rootCandidates, library)
			}
		}

		if len(rootCandidates) > 0 {
			slices.SortFunc(rootCandidates, CompareGongstructByName[*Library])
			rootLibrary = rootCandidates[0]
			rootLibrary.IsRootLibrary = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					"Set existing library as root: "+rootLibrary.GetName())
			}
			needCommit = true
		} else if len(libraries) > 0 {
			// Fallback in case of circular dependencies
			rootLibrary = libraries[0]
			rootLibrary.IsRootLibrary = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					"Set existing library as root (fallback): "+rootLibrary.GetName())
			}
			needCommit = true
		}
	}

	return
}

func (stager *Stager) GetRootLibrary() (rootLibrary *Library) {
	if stager.getRootLibrary() == nil {
		stager.enforceThereIsARootLibrary()
	}
	return stager.getRootLibrary()
}

func (stager *Stager) getRootLibrary() *Library {
	for library := range *GetGongstructInstancesSet[Library](stager.stage) {
		if library.IsRootLibrary {
			return library
		}
	}

	// should not happen
	log.Panic("No root library found")
	return nil
}

func newConcreteAssociation[
	ATstart AbstractType,
	ATend AbstractType,
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType
	},
	ACT_ Gongstruct](start ATstart, end ATend, shapes *[]ACT) ACT {
	compositionShape := ACT(new(ACT_))
	compositionShape.SetName(start.GetName() + " to " + end.GetName())
	compositionShape.SetAbstractStartElement(start)
	compositionShape.SetAbstractEndElement(end)
	compositionShape.SetStartOrientation(ORIENTATION_VERTICAL)
	compositionShape.SetEndOrientation(ORIENTATION_VERTICAL)

	compositionShape.SetCornerOffsetRatio(1.68)
	compositionShape.SetStartRatio(0.5)
	compositionShape.SetEndRatio(0.5)
	*shapes = append(*shapes, compositionShape)

	return compositionShape
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

	// 2. Find all nodes and, if not in a library, reattach to root library
	for _, object := range GetGongstrucsSorted[T](stager.stage) {
		if _, ok := reachable[object]; !ok {
			if object != any(stager.getRootLibrary()) {
				attachDirectlyToLibraryRoot(object)
				// object.UnstageVoid(stager.stage)
				needCommit = true
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Orphan \"%s\" of type \"%s\" was reattached to root library",
						object.GetName(), object.GongGetGongstructName()),
				)
			}
		}
	}

	return
}

// enforceShapesAbstractConsistency iterates over all staged instances to ensure
// that visual shapes maintain valid references to their underlying abstract elements.
//
// If a shape (ConcreteType) is missing its abstract element, or if a link
// (AssociationConcreteType) is missing either its start or end abstract element,
// the shape is considered orphaned and is removed from the stage (unstaged).
// It returns true if any modifications were made, signaling that a commit is needed.
func (stager *Stager) enforceShapesAbstractConsistency() bool {
	stage := stager.stage
	var needCommit bool
	for _, instance := range stage.GetInstances() {
		if shape, ok := instance.(ConcreteType); ok {
			if shape.GetAbstractElement() == nil {
				shape.UnstageVoid(stage)
				needCommit = true
			}
			continue
		}
		if associationShape, ok := instance.(AssociationConcreteType); ok {
			if associationShape.GetAbstractStartElement() == nil {
				associationShape.UnstageVoid(stage)
				needCommit = true
				continue
			}
			if associationShape.GetAbstractEndElement() == nil {
				associationShape.UnstageVoid(stage)
				needCommit = true
				continue
			}
		}

	}

	if needCommit {
		// some slice of pointers might be not clean
		stager.stage.Clean()
	}

	return needCommit
}
