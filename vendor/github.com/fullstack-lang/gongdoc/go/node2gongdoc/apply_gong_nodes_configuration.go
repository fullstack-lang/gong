package node2gongdoc

import gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

func applyGongNodesConfiguration(node *gongdoc_models.Node, isCheckboxDisabled, isChecked bool) {

	node.IsCheckboxDisabled = isCheckboxDisabled
	node.IsChecked = isChecked

	for _, _node := range node.Children {
		applyGongNodesConfiguration(_node, isCheckboxDisabled, isChecked)
	}
}
