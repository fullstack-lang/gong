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

	dataNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.Name != stagedNode.Name {
			data.SetName(frontNode.Name)
			data.SetIsInRenameMode(false)
			stager.stage.Commit()
			return
		}
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			if frontNode.IsExpanded {
				if !slices.Contains(library.DatasWhoseNodeIsExpanded, data) {
					library.DatasWhoseNodeIsExpanded = append(library.DatasWhoseNodeIsExpanded, data)
				}
			} else {
				if idx := slices.Index(library.DatasWhoseNodeIsExpanded, data); idx != -1 {
					library.DatasWhoseNodeIsExpanded = slices.Delete(library.DatasWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			stager.stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(data, GetPointerToGongstructName[*Data]())
		stager.stage.Commit()
	}
}
