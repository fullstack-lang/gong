package node2gongdoc

import gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

// applyGongNodesConfRecursively
func applyGongNodesConfRecursively(node *gongdoc_models.Node, isCheckboxDisabled, isChecked bool) {

	node.IsCheckboxDisabled = isCheckboxDisabled
	node.IsChecked = isChecked

	for _, _node := range node.Children {
		applyGongNodesConfRecursively(_node, isCheckboxDisabled, isChecked)
	}
}
