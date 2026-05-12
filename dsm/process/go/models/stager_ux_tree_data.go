package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeData(
	library *Library,
	data *Data,
	parentNode *tree.Node,
) {
	dataNode := &tree.Node{
		Name:            data.GetName(),
		IsExpanded:      slices.Contains(library.DatasWhoseNodeIsExpanded, data),
		IsNodeClickable: true,
		IsInEditMode:    data.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, dataNode)

	addRenameButton(data, dataNode, stager)

	dataNode.OnNameChange = stager.onNameChange(data)
	dataNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, data, &library.DatasWhoseNodeIsExpanded)
	dataNode.OnClick = onNodeClicked(stager, data)
}
