package models

import (
	"maps"
	"slices"
)

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
