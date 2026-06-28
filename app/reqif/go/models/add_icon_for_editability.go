package models

import (
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func AddIconForEditabilityOfAttribute(attributeIS_EDITABLE bool, attributeLONG_NAME string, nodeAttribute *tree.Node) {
	if attributeIS_EDITABLE {
		nodeAttribute.IsWithPreceedingIcon = true
		nodeAttribute.PreceedingIcon = string(buttons.BUTTON_edit)
	} else {
		nodeAttribute.IsWithPreceedingIcon = true
		nodeAttribute.PreceedingIcon = string(buttons.BUTTON_edit_off)
	}
}
