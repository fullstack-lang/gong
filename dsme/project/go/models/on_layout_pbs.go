package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

var (
	leftOffset = 10.0
	topOffset  = 10.0
)

// onLayoutPBS resets the diagram, creates and adds all possible product-related concrete instances to the diagram,
// and organizes them in a breakdown structure
func onLayoutPBS(stager *Stager, diagram *Diagram) func(
	stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
	return func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
		// 1. Reset the diagram: remove all shapes from the stage to start fresh
		stager.removeAllShapes(diagram)

		// 2. Identify Project
		project := stager.stage.Project_Diagrams_reverseMap[diagram]

		// layout is the function that
		// - layout the product at startX and level
		// - returns the startX for the next product at the same level
		// - call layout on all subProducts
		var layout func(p *Product, startX float64, level int) float64
		layout = func(product *Product, startX float64, level int) float64 {
			// Create Shape
			shape := newShapeToDiagram(product, diagram, &diagram.Product_Shapes, stager.stage)
			x := (startX + float64(product.GetComputedWidth()-1)/2.0) * diagram.DefaultBoxWidth * 1.5
			y := float64(level) * diagram.DefaultBoxHeigth * 2.0
			shape.X = x + leftOffset
			shape.Y = y + topOffset

			// Base Case: Leaf Node (no sub-products)
			if len(product.SubProducts) == 0 {
				return 1.0
			}

			// Recursive Step: Layout Children
			// We stack children horizontally starting from startX
			currentChildX := startX
			for _, subProduct := range product.SubProducts {
				// Recursively layout the child and get its actual visual width
				childWidth := layout(subProduct, currentChildX, level+1)

				// Advance the X cursor by the child's width for the next sibling
				currentChildX += childWidth

				compositionShape := newConcreteAssociation(product, subProduct, &diagram.ProductComposition_Shapes)
				compositionShape.CornerOffsetRatio = 1.5
				_ = compositionShape
			}
			return float64(product.GetComputedWidth())
		}
		width := 0.0
		for _, subProduct := range project.RootProducts {
			width += layout(subProduct, width, 0)
		}

		// 5. Unstage the diagram and re-stage it with StageBranch to include all new shapes recursively
		diagram.Unstage(stager.stage)
		StageBranch(stager.stage, diagram)

		// Save changes
		stager.stage.Commit()
	}
}
