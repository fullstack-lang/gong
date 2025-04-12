package models

import gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

func NewTreeNodeImplNamedStruct(namedStruct NamedStructInterface, probe *Probe2) (res *TreeNodeImplNamedStruct) {

	res = new(TreeNodeImplNamedStruct)

	res.namedStruct = namedStruct
	res.probe = probe

	return

}

type TreeNodeImplNamedStruct struct {
	namedStruct NamedStructInterface
	probe       *Probe2
}

func (i *TreeNodeImplNamedStruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.Stage,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}
}
