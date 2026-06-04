// generated code (do not edit)
package models

import (
	"fmt"
	"time"
)

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
