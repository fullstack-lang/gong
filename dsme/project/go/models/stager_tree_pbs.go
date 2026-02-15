package models

import (
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treePBSRecusriveInDiagram(diagram *Diagram, product *Product, parentNode *tree.Node) {
	productNode := &tree.Node{
		Name:         product.Name,
		IsWithPrefix: true,
		Prefix:       product.ComputedPrefix,

		IsExpanded: slices.Index(diagram.ProductsWhoseNodeIsExpanded, product) != -1,

		HasCheckboxButton:  true,
		IsCheckboxDisabled: !diagram.IsChecked,

		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		ToolTipText:     "Add product to diagram",

		IsNodeClickable: true,

		IsInEditMode: product.IsInRenameMode,
	}
	parentNode.Children = append(parentNode.Children, productNode)

	if !product.IsInRenameMode {
		productNode.Buttons = append(productNode.Buttons,
			&tree.Button{
				Name: product.GetName() + " " + string(buttons.BUTTON_edit_note),
				Icon: string(buttons.BUTTON_edit_note),
				Impl: &tree.FunctionalButtonProxy{
					OnUpdated: func(stage *tree.Stage, button, updatedButton *tree.Button) {
						product.IsInRenameMode = true
						stager.stage.Commit()
					},
				},
				HasToolTip:      true,
				ToolTipText:     "Rename the product",
				ToolTipPosition: tree.Above,
			})
	} else {
		productNode.Buttons = append(productNode.Buttons,
			&tree.Button{
				Name: product.GetName() + " " + string(buttons.BUTTON_edit_off),
				Icon: string(buttons.BUTTON_edit_off),
				Impl: &tree.FunctionalButtonProxy{
					OnUpdated: func(stage *tree.Stage, button, updatedButton *tree.Button) {
						product.IsInRenameMode = false
						stager.stage.Commit()
					},
				},
				HasToolTip:      true,
				ToolTipText:     "Cancel renaming",
				ToolTipPosition: tree.Above,
			})
	}

	addAddItemButton(stager, &diagram.ProductsWhoseNodeIsExpanded, product, nil, productNode, &product.SubProducts, diagram, &diagram.Product_Shapes, &diagram.ProductComposition_Shapes)

	if _, ok := diagram.map_Product_ProductShape[product]; ok {
		productNode.IsChecked = true
	}

	// what to do when the product node is clicked
	productNode.Impl = &tree.FunctionalNodeProxy{
		OnUpdate: onUpdateElementInDiagram(
			stager,
			diagram,
			product,
			&diagram.ProductsWhoseNodeIsExpanded,
			&diagram.Product_Shapes,
			&diagram.map_Product_ProductShape),
	}

	// if product has a parent product, add a button to show/hide the link to the parent
	if parentProduct := product.parentProduct; parentProduct != nil {
		if _, ok := diagram.map_Product_ProductShape[parentProduct]; ok {
			if _, ok := diagram.map_Product_ProductShape[product]; ok {

				showHideCompositionButton := &tree.Button{
					Name: GetGongstructNameFromPointer(product) + " " + string(buttons.BUTTON_add),

					HasToolTip:      true,
					ToolTipPosition: tree.Right,
				}

				if compositionShape, ok := diagram.map_Product_ProductCompositionShape[product]; !ok {
					showHideCompositionButton.Icon = string(buttons.BUTTON_visibility)
					showHideCompositionButton.ToolTipText = "Show link from \"" + parentProduct.Name +
						"\" to \"" + product.Name + "\""

					// showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
					// 	OnUpdated: stager.OnAddProductCompositionShape(diagram, parentProduct, product),
					// }
					showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
						OnUpdated: onAddAssociationShape(stager, parentProduct, product, &diagram.ProductComposition_Shapes),
					}
				} else {
					showHideCompositionButton.Icon = string(buttons.BUTTON_visibility_off)
					showHideCompositionButton.ToolTipText = "Hide link from \"" + parentProduct.Name +
						"\" to \"" + product.Name + "\""

					showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
						OnUpdated: onRemoveAssociationShape(stager, compositionShape, &diagram.ProductComposition_Shapes),
					}
				}
				productNode.Buttons = append(productNode.Buttons, showHideCompositionButton)
			}
		}
	}

	for _, product := range product.SubProducts {
		stager.treePBSRecusriveInDiagram(diagram, product, productNode)
	}
}
