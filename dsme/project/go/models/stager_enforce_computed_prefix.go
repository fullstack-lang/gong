package models

import "fmt"

// enforceComputedPrefix checks that for every [models.Product] and [models.Task] present
// in a DAG attached to the root, the [models.Product.ComputedPrefix]
// and [models.Task.ComputedPrefix] are unique
// and follows a hierarchichal numbering form
//
// if a ComputedPrefix has to be changed, returns true
func (stager *Stager) enforceComputedPrefix() (needCommit bool) {
	root := stager.root

	for _, project := range root.Projects {
		if stager.numberProductNodes(project.RootProducts, "", []int{}) {
			needCommit = true
		}
		if stager.numberTaskNodes(project.RootTasks, "", []int{}) {
			needCommit = true
		}
	}

	return
}

func (stager *Stager) numberProductNodes(nodes []*Product, stringPrefix string, intPrefix []int) (needCommit bool) {
	for i, node := range nodes {
		var nodePrefix string
		if stringPrefix == "" {
			nodePrefix = fmt.Sprintf("%d", i+1)
		} else {
			nodePrefix = fmt.Sprintf("%s.%d", stringPrefix, i+1)
		}

		// compute intPrefix
		// we need a copy of the slice to avoid polluting the slice for other siblings
		nodeIntPrefix := make([]int, len(intPrefix))
		copy(nodeIntPrefix, intPrefix)
		nodeIntPrefix = append(nodeIntPrefix, i)

		if node.ComputedPrefix != nodePrefix {
			node.ComputedPrefix = nodePrefix
			needCommit = true
		}

		node.computedPrefix = nodeIntPrefix

		if stager.numberProductNodes(node.SubProducts, nodePrefix, nodeIntPrefix) {
			needCommit = true
		}

		// compute computedWidth
		// if the node is a leaf, it has a width of 1
		// otherwise, it is the sum of the width of its children
		if len(node.SubProducts) == 0 {
			node.computedWidth = 1
		} else {
			node.computedWidth = 0
			for _, subNode := range node.SubProducts {
				node.computedWidth += subNode.computedWidth
			}
		}
	}
	return
}

func (stager *Stager) numberTaskNodes(nodes []*Task, prefix string, intPrefix []int) (needCommit bool) {
	for i, node := range nodes {
		var nodePrefix string
		if prefix == "" {
			nodePrefix = fmt.Sprintf("%d", i+1)
		} else {
			nodePrefix = fmt.Sprintf("%s.%d", prefix, i+1)
		}

		// compute intPrefix
		// we need a copy of the slice to avoid polluting the slice for other siblings
		nodeIntPrefix := make([]int, len(intPrefix))
		copy(nodeIntPrefix, intPrefix)
		nodeIntPrefix = append(nodeIntPrefix, i)

		if node.ComputedPrefix != nodePrefix {
			node.ComputedPrefix = nodePrefix
			needCommit = true
		}

		node.computedPrefix = nodeIntPrefix

		if stager.numberTaskNodes(node.SubTasks, nodePrefix, nodeIntPrefix) {
			needCommit = true
		}

		// compute computedWidth
		// if the node is a leaf, it has a width of 1
		// otherwise, it is the sum of the width of its children
		if len(node.SubTasks) == 0 {
			node.computedWidth = 1
		} else {
			node.computedWidth = 0
			for _, subNode := range node.SubTasks {
				node.computedWidth += subNode.computedWidth
			}
		}
	}
	return
}
