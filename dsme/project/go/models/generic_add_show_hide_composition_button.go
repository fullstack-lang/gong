package models

import (
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func addShowHideCompositionButton[
	T Gongstruct,
	PT interface {
		*T
		AbstractType
	},
	CT interface {
		*S
		RectShapeInterface
		ConcreteType
	},
	S Gongstruct,
	ACT interface {
		*ACT_
		LinkShapeInterface
		AssociationConcreteType
	},
	ACT_ Gongstruct,
](
	stager *Stager,
	element PT,
	parentElement PT,
	node *tree.Node,
	map_Element_Shape map[PT]CT,
	map_Element_CompositionShape map[PT]ACT,
	compositionShapes *[]ACT,
) {
	if parentElement == nil {
		return
	}

	if _, ok := map_Element_Shape[parentElement]; ok {
		if _, ok := map_Element_Shape[element]; ok {

			showHideCompositionButton := &tree.Button{
				Name:            GetGongstructNameFromPointer(element) + " " + string(buttons.BUTTON_add),
				HasToolTip:      true,
				ToolTipPosition: tree.Right,
			}

			if compositionShape, ok := map_Element_CompositionShape[element]; !ok {
				showHideCompositionButton.Icon = string(buttons.BUTTON_visibility)
				showHideCompositionButton.ToolTipText = "Show link from \"" + parentElement.GetName() +
					"\" to \"" + element.GetName() + "\""

				showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
					OnUpdated: onAddAssociationShapeWithCommit(stager, parentElement, element, compositionShapes),
				}
			} else {
				showHideCompositionButton.Icon = string(buttons.BUTTON_visibility_off)
				showHideCompositionButton.ToolTipText = "Hide link from \"" + parentElement.GetName() +
					"\" to \"" + element.GetName() + "\""

				showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
					OnUpdated: onRemoveAssociationShapeWithCommit(stager, compositionShape, compositionShapes),
				}
			}
			node.Buttons = append(node.Buttons, showHideCompositionButton)
		}
	}
}
