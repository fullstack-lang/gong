package models

import (
	"fmt"
	"slices"
	"time"
)

func (stager *Stager) enforceDAG() (needCommit bool) {

	products := GetGongstrucsSorted[*Product](stager.stage)
	resources := GetGongstrucsSorted[*Resource](stager.stage)
	tasks := GetGongstrucsSorted[*Task](stager.stage)

	// 1. Hierarchy DAG for Product
	needCommit = EnforceDAG(
		stager,
		products,
		func(p *Product) []*Product { return p.SubProducts },
		func(parent, child *Product) {
			for j, sub := range parent.SubProducts {
				if sub == child {
					parent.SubProducts = slices.Delete(parent.SubProducts, j, j+1)
					break
				}
			}
		},
		func(p *Product) string { return "Product " + p.Name },
	) || needCommit

	// 2. Hierarchy DAG for Resource
	needCommit = EnforceDAG(
		stager,
		resources,
		func(r *Resource) []*Resource { return r.SubResources },
		func(parent, child *Resource) {
			for j, sub := range parent.SubResources {
				if sub == child {
					parent.SubResources = slices.Delete(parent.SubResources, j, j+1)
					break
				}
			}
		},
		func(r *Resource) string { return "Resource " + r.Name },
	) || needCommit

	// 3. Hierarchy DAG for Task
	needCommit = EnforceDAG(
		stager,
		tasks,
		func(t *Task) []*Task { return t.SubTasks },
		func(parent, child *Task) {
			for j, sub := range parent.SubTasks {
				if sub == child {
					parent.SubTasks = slices.Delete(parent.SubTasks, j, j+1)
					break
				}
			}
		},
		func(t *Task) string { return "Task " + t.Name },
	) || needCommit

	// 4. Dependency DAG for Tasks and Products (Inputs/Outputs)
	// Build a map of Product -> Tasks that consume it (InputProducts)
	// This allows us to traverse the "Product -> Task" edge efficiently
	productConsumers := make(map[*Product][]*Task)
	for _, task := range tasks {
		for _, inputProd := range task.Inputs {
			productConsumers[inputProd] = append(productConsumers[inputProd], task)
		}
	}

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

	// Call the generic EnforceDAG on the dependency graph
	needCommit = EnforceDAG(
		stager,
		allNodes,
		// getChildren: Returns all nodes that 'n' depends on / points to
		func(n DAGNode) []DAGNode {
			children := []DAGNode{}

			if n.Task != nil {
				// Edge: Task -> OutputProducts
				for _, outProd := range n.Task.Outputs {
					children = append(children, DAGNode{Product: outProd})
				}
			} else if n.Product != nil {
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
			if parent.Task != nil && child.Product != nil {
				// Break Task -> OutputProduct
				p := parent.Task
				c := child.Product
				for j, out := range p.Outputs {
					if out == c {
						p.Outputs = slices.Delete(p.Outputs, j, j+1)
						break
					}
				}
			} else if parent.Product != nil && child.Task != nil {
				// Break Product -> Task (InputProduct)
				// Note: The pointer is stored in the Task (Child), so we remove the Product (Parent) from there
				p := parent.Product
				c := child.Task
				for j, inp := range c.Inputs {
					if inp == p {
						c.Inputs = slices.Delete(c.Inputs, j, j+1)
						break
					}
				}
			}
		},
		func(n DAGNode) string {
			if n.Task != nil {
				return "Task " + n.Task.Name
			} else if n.Product != nil {
				return "Product " + n.Product.Name
			}
			return ""
		},
	) || needCommit

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
	stager *Stager,
	nodes []T,
	getChildren func(T) []T,
	removeChild func(parent T, child T),
	getName func(T) string,
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
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Found loop involving %s, breaking edge with %s", getName(node), getName(child)))
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
