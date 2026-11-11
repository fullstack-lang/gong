package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeProtocolWithinLibrary(
	library *Library,
	protocol *Protocol,
	parentNode *tree.Node,
) {
	protocolNode := &tree.Node{
		Name:            protocol.GetName(),
		IsExpanded:      slices.Contains(library.ProtocolsWhoseNodeIsExpanded, protocol),
		IsNodeClickable: true,
		IsInEditMode:    protocol.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, protocolNode)

	addRenameButton(protocol, protocolNode, stager)

	protocolNode.OnNameChange = stager.onNameChange(protocol)
	protocolNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, protocol, &library.ProtocolsWhoseNodeIsExpanded)
	protocolNode.OnClick = onNodeClicked(stager, protocol)

	layersNode := &tree.Node{
		Name:            "Protocol Layers",
		FontStyle:       tree.ITALIC,
		IsExpanded:      protocol.IsProtocolLayersNodeExpanded,
		IsNodeClickable: true,
	}
	protocolNode.Children = append(protocolNode.Children, layersNode)
	layersNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&protocol.IsProtocolLayersNodeExpanded)
	layersNode.OnClick = onNodeClicked(stager, protocol)

	for _, layer := range protocol.ProtocolLayers {
		stager.treeProtocolLayer(library, layer, layersNode)
	}

	confLayer := ItemButtonConfiguration[
		ProtocolLayer, *ProtocolLayer,
		Protocol, *Protocol,
	]{
		parentNode:                         layersNode,
		sliceForNewAddedItem:               &protocol.ProtocolLayers,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &protocol.IsProtocolLayersNodeExpanded,
	}
	addCreateItemButton(stager, confLayer)

	for _, dataFlow := range stager.rm_Protocol_DataFlows[protocol] {
		stager.treeDataFlowWithinLibrary(library, dataFlow, protocolNode)
	}
}

func (stager *Stager) treeProtocolLayerWithinLibrary(
	library *Library,
	protocolLayer *ProtocolLayer,
	parentNode *tree.Node,
) {
	layerNode := &tree.Node{
		Name:            protocolLayer.GetName(),
		IsExpanded:      slices.Contains(library.ProtocolLayersWhoseNodeIsExpanded, protocolLayer),
		IsNodeClickable: true,
		IsInEditMode:    protocolLayer.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, layerNode)

	addRenameButton(protocolLayer, layerNode, stager)

	layerNode.OnNameChange = stager.onNameChange(protocolLayer)
	layerNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, protocolLayer, &library.ProtocolLayersWhoseNodeIsExpanded)
	layerNode.OnClick = onNodeClicked(stager, protocolLayer)

	for _, protocol := range stager.rm_ProtocolLayer_Protocols[protocolLayer] {
		stager.treeProtocol(library, protocol, layerNode)
	}
}

func (stager *Stager) treeProtocolLayer(
	library *Library,
	protocolLayer *ProtocolLayer,
	parentNode *tree.Node,
) {
	layerNode := &tree.Node{
		Name:            protocolLayer.GetName(),
		IsExpanded:      slices.Contains(library.ProtocolLayersWhoseNodeIsExpanded, protocolLayer),
		IsNodeClickable: true,
		IsInEditMode:    protocolLayer.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, layerNode)

	addRenameButton(protocolLayer, layerNode, stager)

	layerNode.OnNameChange = stager.onNameChange(protocolLayer)
	layerNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, protocolLayer, &library.ProtocolLayersWhoseNodeIsExpanded)
	layerNode.OnClick = onNodeClicked(stager, protocolLayer)
}

func (stager *Stager) treeProtocol(
	library *Library,
	protocol *Protocol,
	parentNode *tree.Node,
) {
	protocolNode := &tree.Node{
		Name:            protocol.GetName(),
		IsExpanded:      slices.Contains(library.ProtocolsWhoseNodeIsExpanded, protocol),
		IsNodeClickable: true,
		IsInEditMode:    protocol.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, protocolNode)

	addRenameButton(protocol, protocolNode, stager)

	protocolNode.OnNameChange = stager.onNameChange(protocol)
	protocolNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, protocol, &library.ProtocolsWhoseNodeIsExpanded)
	protocolNode.OnClick = onNodeClicked(stager, protocol)

	for _, layer := range protocol.ProtocolLayers {
		stager.treeProtocolLayer(library, layer, protocolNode)
	}
}
