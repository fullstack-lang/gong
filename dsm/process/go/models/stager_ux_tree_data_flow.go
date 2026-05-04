package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDataFlows(
	library *Library,
	dataFlow *DataFlow,
	parentNode *tree.Node,
) {
	dataFlowNode := &tree.Node{
		Name:            dataFlow.GetName(),
		IsExpanded:      slices.Contains(library.DataFlowsWhoseNodeIsExpanded, dataFlow),
		IsNodeClickable: true,
		IsInEditMode:    dataFlow.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, dataFlowNode)

	addRenameButton(dataFlow, dataFlowNode, stager)

	dataFlowNode.OnUpdate = func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.Name != stagedNode.Name {
			dataFlow.SetName(frontNode.Name)
			dataFlow.SetIsInRenameMode(false)
			stager.stage.Commit()
			return
		}
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			if frontNode.IsExpanded {
				if !slices.Contains(library.DataFlowsWhoseNodeIsExpanded, dataFlow) {
					library.DataFlowsWhoseNodeIsExpanded = append(library.DataFlowsWhoseNodeIsExpanded, dataFlow)
				}
			} else {
				if idx := slices.Index(library.DataFlowsWhoseNodeIsExpanded, dataFlow); idx != -1 {
					library.DataFlowsWhoseNodeIsExpanded = slices.Delete(library.DataFlowsWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			stager.stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(dataFlow, GetPointerToGongstructName[*DataFlow]())
		stager.stage.Commit()
	}

	//
	// Data
	//
	datasNode := &tree.Node{
		Name:            "Data",
		FontStyle:       tree.ITALIC,
		IsExpanded:      dataFlow.IsDatasNodeExpanded,
		IsNodeClickable: true,
	}
	dataFlowNode.Children = append(dataFlowNode.Children, datasNode)
	datasNode.OnClick = func(frontNode *tree.Node) {
		if frontNode.IsExpanded != dataFlow.IsDatasNodeExpanded {
			dataFlow.IsDatasNodeExpanded = frontNode.IsExpanded
			stager.stage.Commit()
			return
		}
	}

	for _, data := range dataFlow.Datas {
		stager.treeData(library, data, datasNode)
	}
}
