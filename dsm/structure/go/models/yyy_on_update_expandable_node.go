// generated code (do not edit)
package models

import (
	"slices"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func setCallbacksExpandableNode[
	AT interface {
		*AT_
		AbstractType
	},
	AT_ Gongstruct,
](
	stager *Stager,
	node *tree.Node,
	element AT,
	elementsWhoseNodeIsExpanded *[]AT,
) {
	node.OnIsExpandedChange = func(isExpanded bool) {
		if isExpanded {
			if slices.Index(*elementsWhoseNodeIsExpanded, element) == -1 {
				*elementsWhoseNodeIsExpanded = append(*elementsWhoseNodeIsExpanded, element)
			}
		} else {
			if idx := slices.Index(*elementsWhoseNodeIsExpanded, element); idx != -1 {
				*elementsWhoseNodeIsExpanded = slices.Delete(*elementsWhoseNodeIsExpanded, idx, idx+1)
			}
		}
		stager.stage.Commit()
	}

	node.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(element, GetPointerToGongstructName[AT]())
		stager.stage.Commit()
	}
}
