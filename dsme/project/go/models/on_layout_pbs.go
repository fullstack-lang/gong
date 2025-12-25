package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

// OnLayoutPBS resets the diagram, creates and add all all possible product related concrete instances to the diagram,
// and organize them in a breakdown structure
func OnLayoutPBS(stager *Stager, diagram *Diagram) func(
	stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
	return func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
		// 1. Reset the diagram: remove all shapes from the stage to start fresh
		for _, shape := range diagram.Product_Shapes {
			shape.Unstage(stager.stage)
		}
		diagram.Product_Shapes = nil

		for _, shape := range diagram.Task_Shapes {
			shape.Unstage(stager.stage)
		}
		diagram.Task_Shapes = nil

		for _, shape := range diagram.ProductComposition_Shapes {
			shape.Unstage(stager.stage)
		}
		diagram.ProductComposition_Shapes = nil

		for _, shape := range diagram.TaskComposition_Shapes {
			shape.Unstage(stager.stage)
		}
		diagram.TaskComposition_Shapes = nil

		for _, shape := range diagram.TaskInputShapes {
			shape.Unstage(stager.stage)
		}
		diagram.TaskInputShapes = nil

		for _, shape := range diagram.TaskOutputShapes {
			shape.Unstage(stager.stage)
		}
		diagram.TaskOutputShapes = nil

		// 2. Compute the rank of each node (Product or Task)
		// The rank is determined by the longest path from a root node.
		// Rank 0 = Root. Rank N = Dependencies have Max(Rank) = N-1.
		map_Node_Rank := make(map[any]int)

		products := GetGongstrucsSorted[*Product](stager.stage)
		for _, product := range products {
			map_Node_Rank[product] = 0
		}
		tasks := GetGongstrucsSorted[*Task](stager.stage)
		for _, task := range tasks {
			map_Node_Rank[task] = 0
		}

		// to be coded

		// 5. Unstage the diagram and re-stage it with StageBranch to include all new shapes recursively
		diagram.Unstage(stager.stage)
		StageBranch(stager.stage, diagram)

		// Save changes
		stager.stage.Commit()
	}
}
