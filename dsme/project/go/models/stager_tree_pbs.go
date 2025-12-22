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

	productNode.Impl = &NodeProxy[*Product]{
		stager:   stager,
		node:     productNode,
		instance: product,
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
		producersNode.Impl = &expandableNodeProxy{
			node:           producersNode,
			stager:         stager,
			isNodeExpanded: &product.IsProducersNodeExpanded,
		}

		for _, task := range product.producers {
			inputProductNode := &tree.Node{
				Name:            task.GetName(),
				IsExpanded:      true,
				IsNodeClickable: true,
			}
			producersNode.Children = append(producersNode.Children, inputProductNode)
			inputProductNode.Impl = &NodeProxy[*Task]{
				stager:   stager,
				node:     inputProductNode,
				instance: task,
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
		consumersNode.Impl = &expandableNodeProxy{
			node:           consumersNode,
			stager:         stager,
			isNodeExpanded: &product.IsConsumersNodeExpanded,
		}

		for _, task := range product.consumers {
			outputTaskNode := &tree.Node{
				Name:            task.GetName(),
				IsExpanded:      true,
				IsNodeClickable: true,
			}
			consumersNode.Children = append(consumersNode.Children, outputTaskNode)
			outputTaskNode.Impl = &NodeProxy[*Task]{
				stager:   stager,
				node:     outputTaskNode,
				instance: task,
			}
		}
	}
}

func (stager *Stager) treePBSinDiagram(diagram *Diagram, product *Product, parentNode *tree.Node) {
	productNode := &tree.Node{
		Name:               product.ComputedPrefix + " " + product.Name,
		IsExpanded:         slices.Index(diagram.ProductsWhoseNodeIsExpanded, product) != -1,
		IsNodeClickable:    true,
		HasCheckboxButton:  true,
		IsCheckboxDisabled: !diagram.IsChecked,
		HasToolTip:         true,
		ToolTipPosition:    tree.Above,
		ToolTipText:        "Add product to diagram",
	}
	parentNode.Children = append(parentNode.Children, productNode)
	var productShape *ProductShape
	var ok bool
	if productShape, ok = diagram.map_Product_ProductShape[product]; ok {
		productNode.IsChecked = true
	}

	productNode.Impl = &ProductNodeProxyInDiagram[*Product]{
		stager:       stager,
		node:         productNode,
		product:      product,
		diagram:      diagram,
		productShape: productShape,
	}

	for _, product := range product.SubProducts {
		stager.treePBSinDiagram(diagram, product, productNode)
	}
}

type ProductNodeProxyInDiagram[T ProjectElementType] struct {
	stager       *Stager
	node         *tree.Node
	product      *Product
	productShape *ProductShape
	diagram      *Diagram
}

// OnAfterUpdate implements models.NodeImplInterface.
func (p *ProductNodeProxyInDiagram[T]) OnAfterUpdate(stage *tree.Stage, stagedNode *tree.Node, frontNode *tree.Node) {
	if frontNode.IsChecked && !stagedNode.IsChecked {
		stagedNode.IsChecked = frontNode.IsChecked

		// add the Product_Shape
		if p.productShape != nil {
			log.Panic("adding a shape to an already product shape")
		}

		productShape := newProductShapeToDiagram(p.product, p.diagram).Stage(p.stager.stage)
		p.productShape = productShape

		p.stager.stage.Commit()
		return
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {
		stagedNode.IsChecked = frontNode.IsChecked

		// one need to remove the Product_Shape
		if p.productShape == nil {
			log.Panic("remove a non existing shape to product")
		}

		productShape := p.productShape
		p.productShape = nil
		productShape.Unstage(p.stager.stage)

		idx := slices.Index(p.diagram.Product_Shapes, productShape)
		p.diagram.Product_Shapes = slices.Delete(p.diagram.Product_Shapes, idx, idx+1)

		p.stager.stage.Commit()
		return
	}

	if frontNode.IsExpanded != stagedNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded

		if frontNode.IsExpanded {
			if slices.Index(p.diagram.ProductsWhoseNodeIsExpanded, p.product) == -1 {
				p.diagram.ProductsWhoseNodeIsExpanded = append(p.diagram.ProductsWhoseNodeIsExpanded, p.product)
			}
		} else {
			if idx := slices.Index(p.diagram.ProductsWhoseNodeIsExpanded, p.product); idx != -1 {
				p.diagram.ProductsWhoseNodeIsExpanded = slices.Delete(p.diagram.ProductsWhoseNodeIsExpanded, idx, idx+1)
			}
		}
		return
	}

	p.stager.probeForm.FillUpFormFromGongstruct(p.product, GetPointerToGongstructName[*Product]())

	p.stager.stage.Commit()
}
