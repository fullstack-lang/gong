package node2gongdoc

import (
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"
)

// applyGongNodesConfRecursively
func applyGongNodesConfRecursively(node *gongtree_models.Node, isCheckboxDisabled, isChecked bool) {

	node.IsCheckboxDisabled = isCheckboxDisabled
	node.IsChecked = isChecked

	for _, _node := range node.Children {
		applyGongNodesConfRecursively(_node, isCheckboxDisabled, isChecked)
	}
}
