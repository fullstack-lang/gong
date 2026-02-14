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
		needCommit = numberNodes(project.RootProducts, "", []int{}, func(p *Product) []*Product { return p.SubProducts }) || needCommit
		needCommit = numberNodes(project.RootTasks, "", []int{}, func(t *Task) []*Task { return t.SubTasks }) || needCommit
		needCommit = numberNodes(project.Notes, "", []int{}, func(n *Note) []*Note { return nil }) || needCommit
		needCommit = numberNodes(project.RootResources, "", []int{}, func(r *Resource) []*Resource { return r.SubResources }) || needCommit
	}

	return
}

func numberNodes[T AbstractType](
	nodes []T,
	stringPrefix string,
	intPrefix []int,
	getChildren func(T) []T,
) (needCommit bool) {
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

		if node.GetComputedPrefix() != nodePrefix {
			node.SetComputedPrefix(nodePrefix)
			needCommit = true
		}

		node.SetComputedPrefixInt(nodeIntPrefix)

		if numberNodes(getChildren(node), nodePrefix, nodeIntPrefix, getChildren) {
			needCommit = true
		}

		// compute computedWidth
		// if the node is a leaf, it has a width of 1
		// otherwise, it is the sum of the width of its children
		children := getChildren(node)
		if len(children) == 0 {
			node.SetComputedWidth(1)
		} else {
			width := 0
			for _, subNode := range children {
				width += subNode.GetComputedWidth()
			}
			node.SetComputedWidth(width)
		}
	}
	return
}
