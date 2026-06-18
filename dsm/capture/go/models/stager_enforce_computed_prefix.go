package models

import (
	"fmt"
	"time"
)

// enforceComputedPrefix checks that for every [models.Deliverable] and [models.Task] present
// in a DAG attached to the root, the [models.Deliverable.ComputedPrefix]
// and [models.Task.ComputedPrefix] are unique
// and follows a hierarchichal numbering form
//
// if a ComputedPrefix has to be changed, returns true
func (stager *Stager) enforceComputedPrefix() (needCommit bool) {

	for library := range *GetGongstructInstancesSetFromPointerType[*Library](stager.stage) {
		needCommit = numberNodes(stager, library.RootDeliverables, "", []int{}, func(p *Deliverable) []*Deliverable { return p.SubDeliverables }, make(map[*Deliverable]bool)) || needCommit
		needCommit = numberNodes(stager, library.RootConcerns, "", []int{}, func(t *Concern) []*Concern { return t.SubConcerns }, make(map[*Concern]bool)) || needCommit
		needCommit = numberNodes(stager, library.Notes, "", []int{}, func(n *Note) []*Note { return nil }, make(map[*Note]bool)) || needCommit
		needCommit = numberNodes(stager, library.RootStakeholders, "", []int{}, func(r *Stakeholder) []*Stakeholder { return r.SubStakeholders }, make(map[*Stakeholder]bool)) || needCommit
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
