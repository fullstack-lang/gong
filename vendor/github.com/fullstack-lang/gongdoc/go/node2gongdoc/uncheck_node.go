package node2gongdoc

import gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

func UncheckAndDisable(node *gongdoc_models.Node, classdiagram *gongdoc_models.Classdiagram) {

	node.IsCheckboxDisabled = true
	node.IsChecked = false

	nodeImpl := node.Impl

	hasToBeDisabledValue := true
	if classdiagram != nil {
		hasToBeDisabledValue = !classdiagram.IsInDrawMode
	}

	if nodeImpl != nil {
		nodeImpl.SetHasToBeCheckedValue(false)
		nodeImpl.SetHasToBeDisabledValue(hasToBeDisabledValue)
	}

	for _, _node := range node.Children {
		UncheckAndDisable(_node, classdiagram)
	}
}
