package models

import (
	"fmt"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treePBSRecursive(product *Product, parentNode *tree.Node) {
	productNode := &tree.Node{
		Name:            product.ComputedPrefix + " " + product.Name,
		IsExpanded:      product.IsExpanded,
		IsNodeClickable: true,
	}
	parentNode.Children = append(parentNode.Children, productNode)

	productNode.Impl = &tree.FunctionalNodeProxy{
		OnUpdate: OnUpdateAbstractElement(stager, product),
	}

	addAddItemButton(stager, productNode, &product.SubProducts)

	for _, product := range product.SubProducts {
		stager.treePBSRecursive(product, productNode)
	}

	if len(product.producers) > 0 {
		producersNode := &tree.Node{
			Name:                 fmt.Sprintf("(%d)", len(product.producers)),
			IsExpanded:           product.IsProducersNodeExpanded,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_input),
		}
		productNode.Children = append(productNode.Children, producersNode)
		producersNode.Impl = &tree.FunctionalNodeProxy{
			OnUpdate: stager.OnUpdateExpansion(&product.IsProducersNodeExpanded),
		}

		for _, task := range product.producers {
			inputProductNode := &tree.Node{
				Name:            task.GetName(),
				IsExpanded:      true,
				IsNodeClickable: true,
			}
			producersNode.Children = append(producersNode.Children, inputProductNode)
			inputProductNode.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: OnUpdateAbstractElement(stager, task),
			}
		}
	}

	if len(product.consumers) > 0 {
		consumersNode := &tree.Node{
			Name:                 fmt.Sprintf("(%d)", len(product.consumers)),
			IsExpanded:           product.IsConsumersNodeExpanded,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_output),
		}
		productNode.Children = append(productNode.Children, consumersNode)
		consumersNode.Impl = &tree.FunctionalNodeProxy{
			OnUpdate: stager.OnUpdateExpansion(&product.IsConsumersNodeExpanded),
		}

		for _, task := range product.consumers {
			outputTaskNode := &tree.Node{
				Name:            task.GetName(),
				IsExpanded:      true,
				IsNodeClickable: true,
			}
			consumersNode.Children = append(consumersNode.Children, outputTaskNode)
			outputTaskNode.Impl = &tree.FunctionalNodeProxy{
				OnUpdate: OnUpdateAbstractElement(stager, task),
			}
		}
	}
}

func (stager *Stager) treePBSRecusriveInDiagram(diagram *Diagram, product *Product, parentNode *tree.Node) {
	productNode := &tree.Node{
		Name:       product.ComputedPrefix + " " + product.Name,
		IsExpanded: slices.Index(diagram.ProductsWhoseNodeIsExpanded, product) != -1,

		HasCheckboxButton:  true,
		IsCheckboxDisabled: !diagram.IsChecked,

		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		ToolTipText:     "Add product to diagram",

		IsNodeClickable: true,
	}
	parentNode.Children = append(parentNode.Children, productNode)

	addAddItemButton(stager, productNode, &product.SubProducts)

	if _, ok := diagram.map_Product_ProductShape[product]; ok {
		productNode.IsChecked = true
	}

	// what to do when the product node is clicked
	productNode.Impl = &tree.FunctionalNodeProxy{
		OnUpdate: OnUpdateElementInDiagram(
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
					showHideCompositionButton.Icon = string(buttons.BUTTON_unfold_more)
					showHideCompositionButton.ToolTipText = "Show link from \"" + parentProduct.Name +
						"\" to \"" + product.Name + "\""

					// showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
					// 	OnUpdated: stager.OnAddProductCompositionShape(diagram, parentProduct, product),
					// }
					showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
						OnUpdated: OnAddAssociationShape(stager, diagram, parentProduct, product, &diagram.ProductComposition_Shapes),
					}
				} else {
					showHideCompositionButton.Icon = string(buttons.BUTTON_unfold_less)
					showHideCompositionButton.ToolTipText = "Hide link from \"" + parentProduct.Name +
						"\" to \"" + product.Name + "\""

					showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
						OnUpdated: OnRemoveAssociationShape(stager, diagram, compositionShape, &diagram.ProductComposition_Shapes),
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
