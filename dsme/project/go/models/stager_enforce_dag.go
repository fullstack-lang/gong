package models

import "slices"

// EnforceDAG
// Check for cycles in the combined graph of Tasks and Products
//
// The edges are:
// - Task -> SubTask
// - Product -> SubProduct
// - Task -> Product (OutputProducts)
// - Product -> Task (InputProducts)
func (stager *Stager) enforceDAG() (needCommit bool) {

	// 1. Build a map of Product -> Tasks that consume it (InputProducts)
	// This allows us to traverse the "Product -> Task" edge efficiently
	productConsumers := make(map[*Product][]*Task)
	tasks := GetGongstrucsSorted[*Task](stager.stage)
	for _, task := range tasks {
		for _, inputProd := range task.InputProducts {
			productConsumers[inputProd] = append(productConsumers[inputProd], task)
		}
	}

	products := GetGongstrucsSorted[*Product](stager.stage)

	// Define a wrapper node for the mixed graph because EnforceDAG expects a single type T
	type DAGNode struct {
		Product *Product
		Task    *Task
	}

	// Create the list of all nodes in the graph
	allNodes := make([]DAGNode, 0, len(tasks)+len(products))
	for _, t := range tasks {
		allNodes = append(allNodes, DAGNode{Task: t})
	}
	for _, p := range products {
		allNodes = append(allNodes, DAGNode{Product: p})
	}

	// 2. Call the generic EnforceDAG on the unified graph
	needCommit = EnforceDAG(
		allNodes,
		// getChildren: Returns all nodes that 'n' depends on / points to
		func(n DAGNode) []DAGNode {
			children := []DAGNode{}

			if n.Task != nil {
				// Edge: Task -> SubTasks
				for _, subTask := range n.Task.SubTasks {
					children = append(children, DAGNode{Task: subTask})
				}
				// Edge: Task -> OutputProducts
				for _, outProd := range n.Task.OutputProducts {
					children = append(children, DAGNode{Product: outProd})
				}
			} else if n.Product != nil {
				// Edge: Product -> SubProducts
				for _, subProd := range n.Product.SubProducts {
					children = append(children, DAGNode{Product: subProd})
				}
				// Edge: Product -> Consuming Tasks (InputProducts)
				if consumers, ok := productConsumers[n.Product]; ok {
					for _, consumerTask := range consumers {
						children = append(children, DAGNode{Task: consumerTask})
					}
				}
			}
			return children
		},
		// removeChild: Breaks the cycle by removing the specific edge
		func(parent, child DAGNode) {
			if parent.Task != nil && child.Task != nil {
				// Break Task -> SubTask
				p := parent.Task
				c := child.Task
				for j, sub := range p.SubTasks {
					if sub == c {
						p.SubTasks = slices.Delete(p.SubTasks, j, j+1)
						break
					}
				}
			} else if parent.Task != nil && child.Product != nil {
				// Break Task -> OutputProduct
				p := parent.Task
				c := child.Product
				for j, out := range p.OutputProducts {
					if out == c {
						p.OutputProducts = slices.Delete(p.OutputProducts, j, j+1)
						break
					}
				}
			} else if parent.Product != nil && child.Product != nil {
				// Break Product -> SubProduct
				p := parent.Product
				c := child.Product
				for j, sub := range p.SubProducts {
					if sub == c {
						p.SubProducts = slices.Delete(p.SubProducts, j, j+1)
						break
					}
				}
			} else if parent.Product != nil && child.Task != nil {
				// Break Product -> Task (InputProduct)
				// Note: The pointer is stored in the Task (Child), so we remove the Product (Parent) from there
				p := parent.Product
				c := child.Task
				for j, inp := range c.InputProducts {
					if inp == p {
						c.InputProducts = slices.Delete(c.InputProducts, j, j+1)
						break
					}
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
