package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

var (
	leftOffset = 10.0
	topOffset  = 10.0
)

// onLayoutPBS resets the diagram, creates and add all all possible product related concrete instances to the diagram,
// and organize them in a breakdown structure
func onLayoutPBS(stager *Stager, diagram *Diagram) func(
	stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
	return func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
		// 1. Reset the diagram: remove all shapes from the stage to start fresh
		stager.removeAllShapes(diagram)

		// 2. Identify Project
		project := stager.stage.Project_Diagrams_reverseMap[diagram]

		// 3. Algorithm: Recursive Tree Layout using Computed Width
		map_Product_X := make(map[*Product]float64)
		map_Product_Level := make(map[*Product]int)

		// recursive layout function
		var layout func(p *Product, startX float64, level int)
		layout = func(p *Product, startX float64, level int) {
			map_Product_Level[p] = level

			// If leaf (no sub-products), center in the 1-unit slot
			if len(p.SubProducts) == 0 {
				map_Product_X[p] = startX + 0.5
				return
			}

			// If parent, layout children sequentially
			currentChildX := startX
			for _, sub := range p.SubProducts {
				layout(sub, currentChildX, level+1)
				// Increment offset by the pre-computed width of the child's subtree
				currentChildX += float64(sub.GetComputedWidth())
			}

			// Center parent over its children
			// We align the parent with the visual center of its immediate children
			firstChildX := map_Product_X[p.SubProducts[0]]
			lastChildX := map_Product_X[p.SubProducts[len(p.SubProducts)-1]]
			map_Product_X[p] = (firstChildX + lastChildX) / 2.0
		}

		// Apply layout iterating over the defined RootProducts of the project
		// This respects the user-defined order of trees
		var currentRootStartSlot float64 = 0.0
		for _, root := range project.RootProducts {
			layout(root, currentRootStartSlot, 0)
			// Advance by the root's total width plus a gap
			currentRootStartSlot += float64(root.GetComputedWidth()) + 0.5
		}

		// 4. Create Shapes
		// We iterate over the mapped products, which ensures we only create shapes for
		// reachable nodes in the project's hierarchy.
		for product, x_slot := range map_Product_X {
			level := map_Product_Level[product]

			// Position mapping
			// X: slot * box_width * spacing_factor
			// Y: level * box_height * spacing_factor
			x := x_slot * diagram.DefaultBoxWidth * 1.5
			y := float64(level) * diagram.DefaultBoxHeigth * 2.0

			// Create Shape
			shape := newShapeToDiagram(product, diagram, &diagram.Product_Shapes, stager.stage)
			shape.X = x + leftOffset
			shape.Y = y + topOffset

			// Create Composition Shapes (Links to children)
			for _, subProduct := range product.SubProducts {
				compositionShape := newConcreteAssociation(product, subProduct, &diagram.ProductComposition_Shapes)
				_ = compositionShape
			}
		}

		// 5. Unstage the diagram and re-stage it with StageBranch to include all new shapes recursively
		diagram.Unstage(stager.stage)
		StageBranch(stager.stage, diagram)

		// Save changes
		stager.stage.Commit()
	}
}
