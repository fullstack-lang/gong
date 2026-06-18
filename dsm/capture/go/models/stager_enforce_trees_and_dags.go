package models

import (
	"fmt"
	"slices"
	"time"
)

func (stager *Stager) enforceTreesAndDAG() (needCommit bool) {
	deliverables := GetGongstrucsSorted[*Deliverable](stager.stage)

	// 1. Hierarchy Tree for Deliverable
	needCommit = EnforceTree(
		stager,
		deliverables,
		func(p *Deliverable) []*Deliverable { return p.SubDeliverables },
		func(parent, child *Deliverable) {
			for j, sub := range parent.SubDeliverables {
				if sub == child {
					parent.SubDeliverables = slices.Delete(parent.SubDeliverables, j, j+1)
					break
				}
			}
		},
		func(p *Deliverable) string { return "Deliverable " + p.Name },
	) || needCommit

	// 2. Hierarchy Tree for Resource
	resources := GetGongstrucsSorted[*Stakeholder](stager.stage)
	needCommit = EnforceTree(
		stager,
		resources,
		func(r *Stakeholder) []*Stakeholder { return r.SubStakeholders },
		func(parent, child *Stakeholder) {
			for j, sub := range parent.SubStakeholders {
				if sub == child {
					parent.SubStakeholders = slices.Delete(parent.SubStakeholders, j, j+1)
					break
				}
			}
		},
		func(r *Stakeholder) string { return "Resource " + r.Name },
	) || needCommit

	// 3. Hierarchy Tree for Task
	tasks := GetGongstrucsSorted[*Concern](stager.stage)
	needCommit = EnforceTree(
		stager,
		tasks,
		func(t *Concern) []*Concern { return t.SubConcerns },
		func(parent, child *Concern) {
			for j, sub := range parent.SubConcerns {
				if sub == child {
					parent.SubConcerns = slices.Delete(parent.SubConcerns, j, j+1)
					break
				}
			}
		},
		func(t *Concern) string { return "Task " + t.Name },
	) || needCommit

	// 4. Hierarchy Tree for Library
	libraries := GetGongstrucsSorted[*Library](stager.stage)
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
		if library == stager.GetRootLibrary() {
			continue
		}
		for _, sub := range library.SubLibraries {
			isSubLibrary[sub] = true
		}
	}

	var filteredRootLibraries []*Library
	for _, lib := range stager.GetRootLibrary().SubLibraries {
		if !isSubLibrary[lib] {
			filteredRootLibraries = append(filteredRootLibraries, lib)
		} else {
			stager.probeForm.AddNotification(time.Now(),
				fmt.Sprintf("Library %s is a sub-library, removing it from root.SubLibraries", lib.Name))
			needCommit = true
		}
	}
	if len(filteredRootLibraries) != len(stager.GetRootLibrary().SubLibraries) {
		stager.GetRootLibrary().SubLibraries = filteredRootLibraries
	}

	// 5. Dependency DAG for Tasks and Deliverables (Inputs/Outputs)
	// Build a map of Deliverable -> Tasks that consume it (InputDeliverables)
	// This allows us to traverse the "Deliverable -> Task" edge efficiently
	deliverableConsumers := make(map[*Deliverable][]*Concern)
	for _, task := range tasks {
		for _, inputProd := range task.Inputs {
			deliverableConsumers[inputProd] = append(deliverableConsumers[inputProd], task)
		}
	}

	// Define a wrapper node for the mixed graph because EnforceDAG expects a single type T
	type DAGNode struct {
		Deliverable *Deliverable
		Task    *Concern
	}

	// Create the list of all nodes in the graph
	allNodes := make([]DAGNode, 0, len(tasks)+len(deliverables))
	for _, t := range tasks {
		allNodes = append(allNodes, DAGNode{Task: t})
	}
	for _, p := range deliverables {
		allNodes = append(allNodes, DAGNode{Deliverable: p})
	}

	// Call the generic EnforceDAG on the dependency graph
	needCommit = EnforceDAG(
		stager,
		allNodes,
		// getChildren: Returns all nodes that 'n' depends on / points to
		func(n DAGNode) []DAGNode {
			children := []DAGNode{}

			if n.Task != nil {
				// Edge: Task -> OutputDeliverables
				for _, outProd := range n.Task.Outputs {
					children = append(children, DAGNode{Deliverable: outProd})
				}
			} else if n.Deliverable != nil {
				// Edge: Deliverable -> Consuming Tasks (InputDeliverables)
				if consumers, ok := deliverableConsumers[n.Deliverable]; ok {
					for _, consumerTask := range consumers {
						children = append(children, DAGNode{Task: consumerTask})
					}
				}
			}
			return children
		},
		// removeChild: Breaks the cycle by removing the specific edge
		func(parent, child DAGNode) {
			if parent.Task != nil && child.Deliverable != nil {
				// Break Task -> OutputDeliverable
				p := parent.Task
				c := child.Deliverable
				for j, out := range p.Outputs {
					if out == c {
						p.Outputs = slices.Delete(p.Outputs, j, j+1)
						break
					}
				}
			} else if parent.Deliverable != nil && child.Task != nil {
				// Break Deliverable -> Task (InputDeliverable)
				// Note: The pointer is stored in the Task (Child), so we remove the Deliverable (Parent) from there
				p := parent.Deliverable
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
			} else if n.Deliverable != nil {
				return "Deliverable " + n.Deliverable.Name
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
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Node \"%s\" has 2 parents ; \"%s\" and \"%s\". Link to the later is deleted",
						getName(child), getName(existingParent), getName(node)))
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
