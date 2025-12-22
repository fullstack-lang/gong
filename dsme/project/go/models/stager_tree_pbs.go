package models

import (
	"fmt"

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