package models

import (
	"fmt"
	"time"
)

// enforceComputedPrefix checks that for every [models.Product] and [models.Task] present
// in a DAG attached to the root, the [models.Product.ComputedPrefix]
// and [models.Task.ComputedPrefix] are unique
// and follows a hierarchichal numbering form
//
// if a ComputedPrefix has to be changed, returns true
func (stager *Stager) enforceComputedPrefix() (needCommit bool) {

	for library := range *GetGongstructInstancesSetFromPointerType[*Library](stager.stage) {
		needCommit = numberNodes(stager, library.RootProducts, "", []int{}, func(p *Product) []*Product { return p.SubProducts }, make(map[*Product]bool)) || needCommit
		needCommit = numberNodes(stager, library.RootTasks, "", []int{}, func(t *Task) []*Task { return t.SubTasks }, make(map[*Task]bool)) || needCommit
		needCommit = numberNodes(stager, library.Notes, "", []int{}, func(n *Note) []*Note { return nil }, make(map[*Note]bool)) || needCommit
		needCommit = numberNodes(stager, library.RootResources, "", []int{}, func(r *Resource) []*Resource { return r.SubResources }, make(map[*Resource]bool)) || needCommit
		needCommit = numberNodes(stager, library.Diagrams, "", []int{}, func(d *Diagram) []*Diagram { return nil }, make(map[*Diagram]bool)) || needCommit
		// needCommit = numberNodes(stager, library.SubLibraries, "", []int{}, func(l *Library) []*Library { return l.SubLibraries }, make(map[*Library]bool))
	}

	return
}

func numberNodes[T interface {
	AbstractType
	comparable
}](
	stager *Stager,
	nodes []T,
	stringPrefix string,
	intPrefix []int,
	getChildren func(T) []T,
	visited map[T]bool,
) (needCommit bool) {
	index := 0
	for _, node := range nodes {
		if visited[node] {
			continue
		}
		visited[node] = true

		var nodePrefix string
		if stringPrefix == "" {
			nodePrefix = fmt.Sprintf("%d", index+1)
		} else {
			nodePrefix = fmt.Sprintf("%s.%d", stringPrefix, index+1)
		}

		// compute intPrefix
		// we need a copy of the slice to avoid polluting the slice for other siblings
		nodeIntPrefix := make([]int, len(intPrefix))
		copy(nodeIntPrefix, intPrefix)
		nodeIntPrefix = append(nodeIntPrefix, index)

		if oldComputedPrefix := node.GetComputedPrefix(); oldComputedPrefix != nodePrefix {
			node.SetComputedPrefix(nodePrefix)
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("number node renamed prefix of object \"%s\" to \"%s\"",
					oldComputedPrefix, node.GetComputedPrefix()))
			}
		}

		node.SetComputedPrefixInt(nodeIntPrefix)

		if numberNodes(stager, getChildren(node), nodePrefix, nodeIntPrefix, getChildren, visited) {
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

		index++
	}
	return
}
