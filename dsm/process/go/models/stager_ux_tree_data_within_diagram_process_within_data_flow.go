package models

import (
	"log"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDataWithinDiagramProcessWithinDataFlow(
	dataFlowNode *tree.Node,
	diagramProcess *DiagramProcess,
	dataFlow *DataFlow,
	dataFlowShape *DataFlowShape,
	isDataFlowShapePresent bool,
	data *Data,
) {
	stage := stager.stage

	// find the dataShape (if any)
	dataShape, isDataShapePresent := diagramProcess.map_Data_DataShape[data]

	dataNode := &tree.Node{
		Name:               data.GetName(),
		IsNodeClickable:    true,
		IsInEditMode:       data.GetIsInRenameMode(),
		HasCheckboxButton:  true,
		IsChecked:          isDataShapePresent,
		IsCheckboxDisabled: !isDataFlowShapePresent,
	}
	dataFlowNode.Children = append(dataFlowNode.Children, dataNode)

	dataNode.OnClick = func(frontNode *tree.Node) {
		if frontNode.IsChecked && !isDataShapePresent {
			isDataShapePresent = frontNode.IsChecked
			if dataShape != nil {
				log.Panic("adding a shape to an already product shape")
			}

			if dataFlowShape == nil {
				log.Panic("there should be a data flow shape")
			}

			dataShape := (&DataShape{
				Name: data.GetName() + "-" + diagramProcess.GetName(),
				Data: data,
			}).Stage(stager.stage)
			diagramProcess.Data_Shapes = append(diagramProcess.Data_Shapes, dataShape)
			stage.Commit()
			return
		}
		if !frontNode.IsChecked && isDataShapePresent {
			isDataShapePresent = frontNode.IsChecked
			if dataShape == nil {
				log.Panic("remove a non existing shape to product")
			}
			dataShape.UnstageVoid(stage)
			stage.Commit()
			return
		}
	}
}
