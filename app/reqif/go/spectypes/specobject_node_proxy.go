package spectypes

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	m "github.com/fullstack-lang/gong/app/reqif/go/models"
)

type SpecObjectTypeNodeProxy struct {
	stager         *m.Stager
	specObjectType *m.SPEC_OBJECT_TYPE
}

func (proxy *SpecObjectTypeNodeProxy) OnAfterUpdate(
	treeStage *tree.Stage,
	staged, front *tree.Node) {

	if staged.IsExpanded != front.IsExpanded {
		specObjectTypeRendering := GetSpecObjectTypeRendering(proxy.stager.GetStage(), proxy.specObjectType)
		specObjectTypeRendering.IsNodeExpanded = !specObjectTypeRendering.IsNodeExpanded
		staged.IsExpanded = front.IsExpanded
	}
}
