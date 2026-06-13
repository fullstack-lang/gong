package models

import (
	"fmt"
	"slices"
	"time"
)

// enforceTreesAndDAG ensures the structural integrity of the project model by enforcing tree and DAG properties.
//
// Specifically, it performs the following checks and corrections:
// 1. Hierarchy Trees: Ensures that Products, Resources, Tasks, and Libraries form valid forests (collections of trees).
//    - It prevents any node from having multiple parents.
//    - It prevents cyclic parent-child relationships.
// 2. Root Element Cleanup: Ensures that an element is either a root element or a sub-element, but not both.
//    - If a Product, Resource, Task, or Library is found to be a sub-element of another node, it is removed
//      from its library's Root list (e.g., RootProducts, RootTasks, etc.).
// 3. Dependency DAG: Ensures that the relationships between Tasks and Products (via Inputs and Outputs)
//    form a Directed Acyclic Graph.
//    - It prevents circular dependencies (e.g., Task A outputs Product 1, which is input to Task B, which outputs
//      Product 2, which is input to Task A).
//
// If any violations are found, they are automatically corrected (by breaking edges) and the stage is marked
// as needing a commit. Notifications are also added to the stager's probe form if available.
func (stager *Stager) enforceTreesAndDAG() (needCommit bool) {
	products := GetGongstrucsSorted[*Product](stager.stage)
	libraries := GetGongstrucsSorted[*Library](stager.stage)

	// 1. Hierarchy Tree for Product
	needCommit = EnforceTree(
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

	// remove products from library.RootProducts if they are a sub-product of another product
	isSubProduct := make(map[*Product]bool)
	for _, product := range products {
		for _, sub := range product.SubProducts {
			isSubProduct[sub] = true
		}
	}
	for _, lib := range libraries {
		var filteredRootProducts []*Product
		for _, prod := range lib.RootProducts {
			if !isSubProduct[prod] {
				filteredRootProducts = append(filteredRootProducts, prod)
			} else {
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Product %s is a sub-product, removing it from library %s RootProducts", prod.Name, lib.Name))
				}
				needCommit = true
			}
		}
		if len(filteredRootProducts) != len(lib.RootProducts) {
			lib.RootProducts = filteredRootProducts
		}
	}

	// 2. Hierarchy Tree for Resource
	resources := GetGongstrucsSorted[*Resource](stager.stage)
	needCommit = EnforceTree(
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

	// remove resources from library.RootResources if they are a sub-resource of another resource
	isSubResource := make(map[*Resource]bool)
	for _, resource := range resources {
		for _, sub := range resource.SubResources {
			isSubResource[sub] = true
		}
	}
	for _, lib := range libraries {
		var filteredRootResources []*Resource
		for _, res := range lib.RootResources {
			if !isSubResource[res] {
				filteredRootResources = append(filteredRootResources, res)
			} else {
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Resource %s is a sub-resource, removing it from library %s RootResources", res.Name, lib.Name))
				}
				needCommit = true
			}
		}
		if len(filteredRootResources) != len(lib.RootResources) {
			lib.RootResources = filteredRootResources
		}
	}

	// 3. Hierarchy Tree for Task
	tasks := GetGongstrucsSorted[*Task](stager.stage)
	needCommit = EnforceTree(
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

	// remove tasks from library.RootTasks if they are a sub-task of another task
	isSubTask := make(map[*Task]bool)
	for _, task := range tasks {
		for _, sub := range task.SubTasks {
			isSubTask[sub] = true
		}
	}
	for _, lib := range libraries {
		var filteredRootTasks []*Task
		for _, task := range lib.RootTasks {
			if !isSubTask[task] {
				filteredRootTasks = append(filteredRootTasks, task)
			} else {
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Task %s is a sub-task, removing it from library %s RootTasks", task.Name, lib.Name))
				}
				needCommit = true
			}
		}
		if len(filteredRootTasks) != len(lib.RootTasks) {
			lib.RootTasks = filteredRootTasks
		}
	}

	// 4. Hierarchy Tree for Library
	needCommit = EnforceTree(
		stager,
		libraries,
		func(l *Library) []*Library { return l.SubLibraries },
		func(parent, child *Library) {
			for j, sub := range parent.SubLibraries {
				if sub == child {
					parent.SubLibraries = slices.Delete(parent.SubLibraries, j, j+1)
					break
				}
			}
		},
		func(l *Library) string { return "Library " + l.Name },
	) || needCommit

	// remove libraries from root.Libraries if they are a sub-library of another library
	isSubLibrary := make(map[*Library]bool)
	for _, library := range libraries {
		if library == stager.getRootLibrary() {
			continue
		}
		for _, sub := range library.SubLibraries {
			isSubLibrary[sub] = true
		}
	}

	var filteredRootLibraries []*Library
	for _, lib := range stager.getRootLibrary().SubLibraries {
		if !isSubLibrary[lib] {
			filteredRootLibraries = append(filteredRootLibraries, lib)
		} else {
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Library %s is a sub-library, removing it from root.SubLibraries", lib.Name))
			}
			needCommit = true
		}
	}
	if len(filteredRootLibraries) != len(stager.getRootLibrary().SubLibraries) {
		stager.getRootLibrary().SubLibraries = filteredRootLibraries
	}

	// 5. Dependency DAG for Tasks and Products (Inputs/Outputs)
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

// EnforceTree checks that the graph is a forest (a collection of trees).
// It ensures that no node has more than one parent and that there are no cycles.
func EnforceTree[T comparable](
	stager *Stager,
	nodes []T,
	getChildren func(T) []T,
	removeChild func(parent T, child T),
	getName func(T) string,
) (needCommit bool) {
	// 1. Ensure no multiple parents
	parentMap := make(map[T]T)
	for _, node := range nodes {
		children := getChildren(node)
		childrenCopy := make([]T, len(children))
		copy(childrenCopy, children)

		for _, child := range childrenCopy {
			if existingParent, ok := parentMap[child]; ok {
				// Child already has a parent, remove from this node
				removeChild(node, child)
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Node \"%s\" has 2 parents ; \"%s\" and \"%s\". Link to the later is deleted",
							getName(child), getName(existingParent), getName(node)))
				}
				needCommit = true
			} else {
				parentMap[child] = node
			}
		}
	}

	// 2. Ensure no cycles
	needCommit = EnforceDAG(stager, nodes, getChildren, removeChild, getName) || needCommit

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
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(),
						fmt.Sprintf("Found loop involving %s, breaking edge with %s", getName(node), getName(child)))
				}
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
