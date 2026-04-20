package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func onUpdateExpandableNode[
	AT interface {
		*AT_
		AbstractType
	},
	AT_ Gongstruct,
](
	stager *Stager,
	element AT,
	elementsWhoseNodeIsExpanded *[]AT,
) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			if frontNode.IsExpanded {
				if slices.Index(*elementsWhoseNodeIsExpanded, element) == -1 {
					*elementsWhoseNodeIsExpanded = append(*elementsWhoseNodeIsExpanded, element)
				}
			} else {
				if idx := slices.Index(*elementsWhoseNodeIsExpanded, element); idx != -1 {
					*elementsWhoseNodeIsExpanded = slices.Delete(*elementsWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			return
		}

		stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[AT]())
		stager.stage.Commit()
	}
}
