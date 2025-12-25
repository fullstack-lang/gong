package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

// OnShowAllInDiagram resets the diagram, creates and add all all possible concrete instances to the diagram,
// and organize them in a pseudo pert diagram
func OnShowAllInDiagram(stager *Stager, diagram *Diagram) func(
	stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
	return func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
		// 1. Reset the diagram: remove all shapes to start fresh
		// We unstage existing shapes to ensure they are removed from the stage
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

		// Relax edges to propagate ranks.
		// We loop enough times to cover the maximum possible path length (number of nodes).
		// Edges considered:
		// - Product -> SubProduct
		// - Task -> SubTask
		// - Task -> Output (Product)
		// - Product -> Input (Task)
		nbNodes := len(products) + len(tasks)
		for i := 0; i < nbNodes; i++ {

			// Product Composition: Parent Rank < Child Rank
			for _, product := range products {
				for _, subProduct := range product.SubProducts {
					if map_Node_Rank[subProduct] <= map_Node_Rank[product] {
						map_Node_Rank[subProduct] = map_Node_Rank[product] + 1
					}
				}
			}

			// Task Composition: Parent Rank < Child Rank
			for _, task := range tasks {
				for _, subTask := range task.SubTasks {
					if map_Node_Rank[subTask] <= map_Node_Rank[task] {
						map_Node_Rank[subTask] = map_Node_Rank[task] + 1
					}
				}
			}

			// Task Output: Task Rank < Output Product Rank
			for _, task := range tasks {
				for _, output := range task.Outputs {
					if map_Node_Rank[output] <= map_Node_Rank[task] {
						map_Node_Rank[output] = map_Node_Rank[task] + 1
					}
				}
			}

			// Task Input: Input Product Rank < Task Rank
			for _, task := range tasks {
				for _, input := range task.Inputs {
					if map_Node_Rank[task] <= map_Node_Rank[input] {
						map_Node_Rank[task] = map_Node_Rank[input] + 1
					}
				}
			}
		}

		// 3. Create and Position Shapes
		const (
			X_STEP   = 300.0
			Y_STEP   = 100.0
			OFFSET_X = 50.0
			OFFSET_Y = 50.0
		)

		// Helper to keep track of vertical stacking within a rank column
		rankCounters := make(map[int]int)

		// Create Product Shapes
		for _, product := range products {
			rank := map_Node_Rank[product]
			index := rankCounters[rank]
			rankCounters[rank]++

			x := OFFSET_X + float64(rank)*X_STEP
			y := OFFSET_Y + float64(index)*Y_STEP

			// Note: We do not Stage the shape here, StageBranch will handle it later
			shape := &ProductShape{Name: product.Name}
			shape.Product = product
			shape.X = x
			shape.Y = y
			shape.Width = 200
			shape.Height = 60
			diagram.Product_Shapes = append(diagram.Product_Shapes, shape)
		}

		// Create Task Shapes
		for _, task := range tasks {
			rank := map_Node_Rank[task]
			index := rankCounters[rank]
			rankCounters[rank]++

			x := OFFSET_X + float64(rank)*X_STEP
			y := OFFSET_Y + float64(index)*Y_STEP

			// Note: We do not Stage the shape here, StageBranch will handle it later
			shape := &TaskShape{Name: task.Name}
			shape.Task = task
			shape.X = x
			shape.Y = y
			shape.Width = 200
			shape.Height = 60
			diagram.Task_Shapes = append(diagram.Task_Shapes, shape)
		}

		// 4. Create Link Shapes using newConcreteAssociation
		// Product Composition
		for _, product := range products {
			for _, subProduct := range product.SubProducts {
				link := newConcreteAssociation[*Product, *Product, *ProductCompositionShape, ProductCompositionShape](product, subProduct)
				diagram.ProductComposition_Shapes = append(diagram.ProductComposition_Shapes, link)
			}
		}

		// Task Composition
		for _, task := range tasks {
			for _, subTask := range task.SubTasks {
				link := newConcreteAssociation[*Task, *Task, *TaskCompositionShape, TaskCompositionShape](task, subTask)
				diagram.TaskComposition_Shapes = append(diagram.TaskComposition_Shapes, link)
			}
		}

		// Task Inputs (Product -> Task)
		for _, task := range tasks {
			for _, input := range task.Inputs {
				// Start: Task, End: Product (as per Gong struct definition for TaskInputShape)
				link := newConcreteAssociation[*Task, *Product, *TaskInputShape, TaskInputShape](task, input)
				diagram.TaskInputShapes = append(diagram.TaskInputShapes, link)
			}
			for _, output := range task.Outputs {
				// Start: Task, End: Product
				link := newConcreteAssociation[*Task, *Product, *TaskOutputShape, TaskOutputShape](task, output)
				diagram.TaskOutputShapes = append(diagram.TaskOutputShapes, link)
			}
		}

		// 5. Unstage the diagram and re-stage it with StageBranch to include all new shapes
		diagram.Unstage(stager.stage)
		StageBranch(stager.stage, diagram)

		// Save changes
		stager.stage.Commit()
	}
}
