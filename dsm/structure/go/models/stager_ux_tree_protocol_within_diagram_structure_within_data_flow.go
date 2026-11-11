package models

import (
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeProtocolWithinDiagramStructureWithinDataFlow(
	dataFlowNode *tree.Node,
	diagramStructure *DiagramStructure,
	dataFlow *DataFlow,
	dataFlowShape *DataFlowShape,
	isDataFlowShapePresent bool,
	protocol *Protocol,
) {
	stage := stager.stage

	protoKey := protocolShapeKey{
		dataFlow: dataFlow,
		protocol: protocol,
	}
	protocolShape, isProtocolShapePresent := diagramStructure.map_ProtocolShapeKey_ProtocolShape[protoKey]

	protocolNode := &tree.Node{
		Name:                    protocol.GetName(),
		IsExpanded:              slices.Contains(diagramStructure.DataFlowsWhoseProtocolNodeIsExpanded, dataFlow),
		IsNodeClickable:         true,
		IsInEditMode:            protocol.GetIsInRenameMode(),
		HasCheckboxButton:       true,
		IsChecked:               isProtocolShapePresent,
		IsCheckboxDisabled:      !isDataFlowShapePresent,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if !isDataFlowShapePresent {
				return "A protocol shape cannot be created if the data flow shape is not present"
			}
			if isProtocolShapePresent {
				return "Click to remove the protocol shape"
			}
			return "Click to create a protocol shape for this protocol within this data flow"
		}(),
	}
	dataFlowNode.Children = append(dataFlowNode.Children, protocolNode)
	addRenameButton(protocol, protocolNode, stager)

	if isProtocolShapePresent && protocolShape != nil {
		visibilityButton := &tree.Button{
			Name:            diagramStructure.GetName(),
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				protocolShape.SetIsHidden(!protocolShape.GetIsHidden())
				stage.Commit()
			},
		}
		if protocolShape.GetIsHidden() {
			visibilityButton.Icon = string(buttons.BUTTON_visibility)
			visibilityButton.ToolTipText = "Show on diagram"
		}
		protocolNode.Buttons = append(protocolNode.Buttons, visibilityButton)
	}

	protocolNode.OnNameChange = stager.onNameChange(protocol)
	protocolNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, dataFlow, &diagramStructure.DataFlowsWhoseProtocolNodeIsExpanded)

	protocolNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked && !isProtocolShapePresent {
			shapeName := protocol.GetName() + "-" + dataFlow.GetName() + "-" + diagramStructure.GetName()
			if dataFlowShape != nil {
				shapeName = protocol.GetName() + "-" + dataFlowShape.GetName() + "-" + diagramStructure.GetName()
			}
			newProtocolShape := (&ProtocolShape{
				Name:     shapeName,
				Protocol: protocol,
				DataFlow: dataFlow,
			}).Stage(stage)
			newProtocolShape.SetWidth(140.0)
			diagramStructure.Protocol_Shapes = append(diagramStructure.Protocol_Shapes, newProtocolShape)
			stage.Commit()
			return
		}
		if !isChecked && isProtocolShapePresent {
			protocolShape.UnstageVoid(stage)
			stage.Commit()
			return
		}
	}
	protocolNode.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(protocol, "Protocol")
		if isProtocolShapePresent && protocolShape != nil {
			fillUpShapeAppearanceForm(stager, protocolShape)
		} else {
			clearShapeAppearanceForm(stager)
		}
		stager.stage.CommitWithSuspendedCallbacks()
	}

	for _, protocolLayer := range protocol.ProtocolLayers {
		if protocolLayer == nil {
			continue
		}
		layer := protocolLayer
		layerKey := protocolLayerShapeKey{
			dataFlow:      dataFlow,
			protocolLayer: layer,
		}
		layerShape, isLayerShapePresent := diagramStructure.map_ProtocolLayerShapeKey_ProtocolLayerShape[layerKey]

		layerNode := &tree.Node{
			Name:                    layer.GetName(),
			IsNodeClickable:         true,
			IsInEditMode:            layer.GetIsInRenameMode(),
			HasCheckboxButton:       true,
			IsChecked:               isLayerShapePresent,
			IsCheckboxDisabled:      !isDataFlowShapePresent,
			CheckboxHasToolTip:      true,
			CheckboxToolTipPosition: tree.Left,
			CheckboxToolTipText: func() string {
				if !isDataFlowShapePresent {
					return "A protocol layer shape cannot be created if the data flow shape is not present"
				}
				if isLayerShapePresent {
					return "Click to remove the protocol layer shape"
				}
				return "Click to create a protocol layer shape within this data flow"
			}(),
		}
		protocolNode.Children = append(protocolNode.Children, layerNode)
		addRenameButton(layer, layerNode, stager)

		if isLayerShapePresent && layerShape != nil {
			visibilityButton := &tree.Button{
				Name:            diagramStructure.GetName(),
				Icon:            string(buttons.BUTTON_visibility_off),
				ToolTipText:     "Hide from diagram",
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
				OnClick: func() {
					layerShape.SetIsHidden(!layerShape.GetIsHidden())
					stage.Commit()
				},
			}
			if layerShape.GetIsHidden() {
				visibilityButton.Icon = string(buttons.BUTTON_visibility)
				visibilityButton.ToolTipText = "Show on diagram"
			}
			layerNode.Buttons = append(layerNode.Buttons, visibilityButton)
		}

		layerNode.OnNameChange = stager.onNameChange(layer)
		layerNode.OnIsCheckedChanged = func(isChecked bool) {
			if isChecked && !isLayerShapePresent {
				shapeName := layer.GetName() + "-" + dataFlow.GetName() + "-" + diagramStructure.GetName()
				if dataFlowShape != nil {
					shapeName = layer.GetName() + "-" + dataFlowShape.GetName() + "-" + diagramStructure.GetName()
				}
				newLayerShape := (&ProtocolLayerShape{
					Name:          shapeName,
					ProtocolLayer: layer,
					DataFlow:      dataFlow,
				}).Stage(stage)
				newLayerShape.SetWidth(140.0)
				diagramStructure.ProtocolLayer_Shapes = append(diagramStructure.ProtocolLayer_Shapes, newLayerShape)
				stage.Commit()
				return
			}
			if !isChecked && isLayerShapePresent {
				layerShape.UnstageVoid(stage)
				stage.Commit()
				return
			}
		}
		layerNode.OnClick = func(frontNode *tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(layer, "ProtocolLayer")
			if isLayerShapePresent && layerShape != nil {
				fillUpShapeAppearanceForm(stager, layerShape)
			} else {
				clearShapeAppearanceForm(stager)
			}
			stager.stage.CommitWithSuspendedCallbacks()
		}
	}
}
