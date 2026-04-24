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
