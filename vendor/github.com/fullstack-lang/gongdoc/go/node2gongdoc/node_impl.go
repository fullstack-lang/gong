package node2gongdoc

import gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

type NodeImpl struct {
	node   *gongdoc_models.Node
	nodeCb *NodeCB

	hasRelatedDiagramElt     bool
	canHaveRelatedDiagramElt bool
}

// HasToBeChecked implements the interface needed by the tree/node package
func (nodeImpl *NodeImpl) HasToBeChecked() bool {
	return nodeImpl.hasRelatedDiagramElt
}

func (nodeImpl *NodeImpl) SetHasToBeCheckedValue(value bool) {
	nodeImpl.hasRelatedDiagramElt = value
}

// HasToBeDisabled implements the interface needed by the tree/node package
func (nodeImpl *NodeImpl) HasToBeDisabled() bool {
	return nodeImpl.canHaveRelatedDiagramElt
}

func (nodeImpl *NodeImpl) SetHasToBeDisabledValue(value bool) {
	nodeImpl.canHaveRelatedDiagramElt = value
}
