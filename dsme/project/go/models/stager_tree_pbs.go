package models

import (
	"fmt"
	"log"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treePBS(product *Product, parentNode *tree.Node) {
	productNode := &tree.Node{
		Name:            product.ComputedPrefix + " " + product.Name,
		IsExpanded:      product.IsExpanded,
		IsNodeClickable: true,
	}
	parentNode.Children = append(parentNode.Children, productNode)

	productNode.Impl = &tree.FunctionalNodeProxy{
		OnUpdate: stager.OnUpdateProduct(product),
	}

	addAddItemButton(stager, productNode, &product.SubProducts)

	for _, product := range product.SubProducts {
		stager.treePBS(product, productNode)
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
				OnUpdate: stager.OnUpdateTask(task),
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
				OnUpdate: stager.OnUpdateTask(task),
			}
		}
	}
}

func (stager *Stager) treePBSinDiagram(diagram *Diagram, product *Product, parentNode *tree.Node) {
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

	if _, ok := diagram.map_Product_ProductShape[product]; ok {
		productNode.IsChecked = true
	}

	// what to do when the product node is clicked
	productNode.Impl = &tree.FunctionalNodeProxy{
		OnUpdate: stager.OnUpdateProductInDiagram(diagram, product),
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
						OnUpdated: OnAddCompositionShape(stager, diagram, parentProduct, product, &diagram.ProductComposition_Shapes),
					}
				} else {
					showHideCompositionButton.Icon = string(buttons.BUTTON_unfold_less)
					showHideCompositionButton.ToolTipText = "Hide link from \"" + parentProduct.Name +
						"\" to \"" + product.Name + "\""

					showHideCompositionButton.Impl = &tree.FunctionalButtonProxy{
						OnUpdated: stager.OnRemoveProductCompositionShape(diagram, compositionShape),
					}
				}
				productNode.Buttons = append(productNode.Buttons, showHideCompositionButton)
			}
		}
	}

	for _, product := range product.SubProducts {
		stager.treePBSinDiagram(diagram, product, productNode)
	}
}

// Helper callbacks

func (stager *Stager) OnUpdateProduct(product *Product) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			product.IsExpanded = frontNode.IsExpanded
		} else {
			stager.probeForm.FillUpFormFromGongstruct(product, GetPointerToGongstructName[*Product]())
		}
		stager.stage.Commit()
	}
}

func (stager *Stager) OnUpdateTask(task *Task) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			task.IsExpanded = frontNode.IsExpanded
		} else {
			stager.probeForm.FillUpFormFromGongstruct(task, GetPointerToGongstructName[*Task]())
		}
		stager.stage.Commit()
	}
}

func (stager *Stager) OnUpdateProductInDiagram(diagram *Diagram, product *Product) func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
	return func(stage *tree.Stage, stagedNode, frontNode *tree.Node) {
		// find the shape (if any)
		productShape := diagram.map_Product_ProductShape[product]

		if frontNode.IsChecked && !stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if productShape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			productShape = newProductShapeToDiagram(product, diagram).Stage(stager.stage)
			stager.stage.Commit()
			return
		}
		if !frontNode.IsChecked && stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if productShape == nil {
				log.Panic("remove a non existing shape to product")
			}
			productShape.Unstage(stager.stage)
			idx := slices.Index(diagram.Product_Shapes, productShape)
			diagram.Product_Shapes = slices.Delete(diagram.Product_Shapes, idx, idx+1)
			stager.stage.Commit()
			return
		}

		if frontNode.IsExpanded != stagedNode.IsExpanded {
			stagedNode.IsExpanded = frontNode.IsExpanded
			if frontNode.IsExpanded {
				if slices.Index(diagram.ProductsWhoseNodeIsExpanded, product) == -1 {
					diagram.ProductsWhoseNodeIsExpanded = append(diagram.ProductsWhoseNodeIsExpanded, product)
				}
			} else {
				if idx := slices.Index(diagram.ProductsWhoseNodeIsExpanded, product); idx != -1 {
					diagram.ProductsWhoseNodeIsExpanded = slices.Delete(diagram.ProductsWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			return
		}

		stager.probeForm.FillUpFormFromGongstruct(product, GetPointerToGongstructName[*Product]())
		stager.stage.Commit()
	}
}

func (stager *Stager) OnAddProductCompositionShape(diagram *Diagram, parentProduct, product *Product) func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
	return func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
		compositionShape := new(ProductCompositionShape).Stage(stager.stage)
		compositionShape.Name = parentProduct.Name + " to " + product.Name
		compositionShape.Product = product
		compositionShape.StartOrientation = ORIENTATION_VERTICAL
		compositionShape.EndOrientation = ORIENTATION_VERTICAL
		compositionShape.CornerOffsetRatio = 1.68 // near the golden ratio
		compositionShape.StartRatio = 0.5
		compositionShape.EndRatio = 0.5

		diagram.ProductComposition_Shapes = append(diagram.ProductComposition_Shapes, compositionShape)
		stager.stage.Commit()
	}
}

func (stager *Stager) OnRemoveProductCompositionShape(diagram *Diagram, compositionShape *ProductCompositionShape) func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
	return func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
		compositionShape.Unstage(stager.stage)
		idx := slices.Index(diagram.ProductComposition_Shapes, compositionShape)
		diagram.ProductComposition_Shapes = slices.Delete(diagram.ProductComposition_Shapes, idx, idx+1)
		stager.stage.Commit()
	}
}

func OnAddCompositionShape[
	AT AbstractType,
	CCT interface {
		*CCT_
		LinkShapeInterface
		CompositionConcreteType
	},
	CCT_ Gongstruct](
	stager *Stager, diagram *Diagram, parent, child AT, shapes *[]CCT) func(
	stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
	return func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
		compositionShape := CCT(new(CCT_))
		compositionShape.StageVoid(stager.stage)
		compositionShape.SetName(parent.GetName() + " to " + child.GetName())
		compositionShape.SetAbstractChildElement(child)
		compositionShape.SetStartOrientation(ORIENTATION_VERTICAL)
		compositionShape.SetEndOrientation(ORIENTATION_VERTICAL)
		compositionShape.SetCornerOffsetRatio(1.68)
		compositionShape.SetStartRatio(0.5)
		compositionShape.SetEndRatio(0.5)

		*shapes = append(*shapes, compositionShape)
		stager.stage.Commit()
	}
}
