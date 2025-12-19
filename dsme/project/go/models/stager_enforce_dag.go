package models

import "slices"

// EnforceDAG
// All [models.Product] are within at least one DAG
// This function checks that there is no cycles within the DAG.
// If a cycle is found the culprit association is deleted
func (stager *Stager) enforceDAG() (needCommit bool) {

	products := GetGongstrucsSorted[*Product](stager.stage)

	// 2. Call the generic EnforceDAG
	needCommit = needCommit || EnforceDAG(
		products,
		// getChildren: How to access sub-products
		func(p *Product) []*Product {
			return p.SubProducts
		},
		// removeChild: How to remove a sub-product link
		func(parent, child *Product) {
			for j, p := range parent.SubProducts {
				if p == child {
					parent.SubProducts = slices.Delete(parent.SubProducts, j, j+1)
					break
				}
			}
		},
	)

	tasks := GetGongstrucsSorted[*Task](stager.stage)

	needCommit = needCommit || EnforceDAG(
		tasks,
		// getChildren: How to access sub-tasks
		func(p *Task) []*Task {
			return p.SubTasks
		},
		// removeChild: How to remove a sub-task link
		func(parent, child *Task) {
			for j, p := range parent.SubTasks {
				if p == child {
					parent.SubTasks = slices.Delete(parent.SubTasks, j, j+1)
					break
				}
			}
		},
	)

	return
}

// EnforceDAG checks for cycles in a graph of nodes of type T.
// If a cycle is detected, it breaks the cycle by removing the edge to the visited node.
//
// generic parameters:
// nodes: a slice containing all nodes in the graph to ensure we visit disconnected components.
// getChildren: a function that returns the children/neighbors of a given node.
// removeChild: a function that removes a specific child from a parent node.
func EnforceDAG[T comparable](
	nodes []T,
	getChildren func(T) []T,
	removeChild func(parent T, child T),
) (needCommit bool) {

	// Sets for DFS cycle detection
	whiteSet := make(map[T]struct{}) // Not visited
	graySet := make(map[T]struct{})  // Visiting (current path)
	blackSet := make(map[T]struct{}) // Visited

	// Initialize white set with all nodes
	for _, node := range nodes {
		whiteSet[node] = struct{}{}
	}

	var dfs func(node T)

	dfs = func(node T) {
		// Move node from white to gray
		delete(whiteSet, node)
		graySet[node] = struct{}{}

		// Get children. We must iterate over a copy because removeChild might modify the slice
		children := getChildren(node)
		childrenCopy := make([]T, len(children))
		copy(childrenCopy, children)

		for _, child := range childrenCopy {
			// If child is in gray set, we found a back edge -> CYCLE
			if _, ok := graySet[child]; ok {
				removeChild(node, child)
				needCommit = true
				continue
			}

			// If child is in black set, it's already fully processed
			if _, ok := blackSet[child]; ok {
				continue
			}

			// If child is in white set, recurse
			if _, ok := whiteSet[child]; ok {
				dfs(child)
			}
		}

		// Move node from gray to black
		delete(graySet, node)
		blackSet[node] = struct{}{}
	}

	// Iterate over all nodes to handle disconnected subgraphs
	// We iterate over the input slice to ensure deterministic order if the slice is sorted
	for _, node := range nodes {
		if _, ok := whiteSet[node]; ok {
			dfs(node)
		}
	}

	return
}
