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

		// 2. Compute the rank of each node (Product or Task)
		// The rank is determined by the longest path from a root node.
		// Rank 0 = Root. Rank N = Dependencies have Max(Rank) = N-1.
		map_Node_Rank := make(map[any]int)

		allProducts := GetGongstrucsSorted[*Product](stager.stage)
		project := stager.stage.Project_Diagrams_reverseMap[diagram]
		var products []*Product
		for _, product := range allProducts {
			if stager.productToProject[product] != project {
				continue
			}
			products = append(products, product)
			map_Node_Rank[product] = 0
		}

		// 3. Propagate ranks to children (Products and Tasks)
		// We use a fixed-point iteration to handle the depth propagation
		iterations := 0
		maxIterations := len(products) + 1 // Safe limit to prevent infinite loops in cyclic graphs
		for iterations < maxIterations {
			changed := false
			for _, product := range products {
				currentRank := map_Node_Rank[product]
				for _, subProduct := range product.SubProducts {
					if map_Node_Rank[subProduct] <= currentRank {
						map_Node_Rank[subProduct] = currentRank + 1
						changed = true
					}
				}
			}
			if !changed {
				break
			}
			iterations++
		}

		// 4. Create and position shapes

		// Use a tree layout algorithm for X positions
		map_Product_X := make(map[*Product]float64)
		var currentX float64 = 0
		visited := make(map[*Product]bool)

		// position recursively calculates the X coordinate
		var position func(p *Product)
		position = func(p *Product) {
			if visited[p] {
				return
			}
			visited[p] = true

			// Identify children that are part of the current project diagram
			var validChildren []*Product
			for _, sub := range p.SubProducts {
				if _, ok := map_Node_Rank[sub]; ok {
					validChildren = append(validChildren, sub)
				}
			}

			if len(validChildren) == 0 {
				// Leaf node: take the next available slot
				map_Product_X[p] = currentX
				currentX += 1.0
			} else {
				// Internal node: layout children first
				for _, child := range validChildren {
					position(child)
				}
				// Center parent over children
				minX := map_Product_X[validChildren[0]]
				maxX := map_Product_X[validChildren[len(validChildren)-1]]
				map_Product_X[p] = (minX + maxX) / 2.0
			}
		}

		// Apply layout starting from Roots (Rank 0)
		for _, product := range products {
			if map_Node_Rank[product] == 0 {
				position(product)
			}
		}

		// Handle any disconnected components or cycles not reached via roots
		for _, product := range products {
			if !visited[product] {
				position(product)
			}
		}

		// Create Product Shapes
		for _, product := range products {
			rank := map_Node_Rank[product]

			// Position
			x := map_Product_X[product] * defaultBoxWidth * 1.5
			y := float64(rank) * defaultBoxHeigth * 2.0

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
