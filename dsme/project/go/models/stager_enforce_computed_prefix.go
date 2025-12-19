package models

import "fmt"

// enforceComputedPrefix checks that for every [models.Product] present
// in a DAG attached to the root, the [models.Product.ComputedPrefix] is unique
// and follows a hierarchichal numbering form
//
// if a ComputedPrefix has to be changed, returns true
func (stager *Stager) enforceComputedPrefix() (needCommit bool) {

	root := stager.root

	for _, project := range root.Projects {
		if stager.numberNodes(project.RootProducts, "") {
			needCommit = true
		}
	}

	return
}

func (stager *Stager) numberNodes(nodes []*Product, prefix string) (needCommit bool) {

	for i, node := range nodes {
		var nodePrefix string
		if prefix == "" {
			nodePrefix = fmt.Sprintf("%d", i+1)
		} else {
			nodePrefix = fmt.Sprintf("%s.%d", prefix, i+1)
		}

		if node.ComputedPrefix != nodePrefix {
			node.ComputedPrefix = nodePrefix
			needCommit = true
		}

		if stager.numberNodes(node.SubProducts, nodePrefix) {
			needCommit = true
		}
	}
	return
}
