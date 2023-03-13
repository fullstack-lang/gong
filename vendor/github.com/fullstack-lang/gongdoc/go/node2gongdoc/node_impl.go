package node2gongdoc

import gongdoc_models "github.com/fullstack-lang/gongdoc/go/models"

type NodeImpl struct {
	node   *gongdoc_models.Node
	nodeCb *NodeCB
}

func (nodeImpl *NodeImpl) DisableNodeCheckbox() {
	nodeImpl.node.IsCheckboxDisabled = true
}

func (nodeImpl *NodeImpl) CheckNode() {
	nodeImpl.node.IsChecked = true
}
