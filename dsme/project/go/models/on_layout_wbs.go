package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

// onLayoutWBS resets the diagram, creates and adds all possible task-related concrete instances to the diagram,
// and organizes them in a breakdown structure
func onLayoutWBS(stager *Stager, diagram *Diagram) func(
	stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
	return func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
		// 1. Reset the diagram: remove all shapes from the stage to start fresh
		diagram.Product_Shapes = []*ProductShape{}
		diagram.Task_Shapes = []*TaskShape{}
		diagram.Note_Shapes = []*NoteShape{}

		// 2. Identify Project
		project := stager.stage.Project_Diagrams_reverseMap[diagram]

		// layout is the function that
		// - layout the task at startX and level
		// - returns the startX for the next task at the same level
		// - call layout on all subTasks
		var layout func(p *Task, startX float64, level int) float64
		layout = func(task *Task, startX float64, level int) float64 {
			// Create Shape
			shape := newShapeToDiagram(task, diagram, &diagram.Task_Shapes, stager.stage)
			x := (startX + float64(task.GetComputedWidth()-1)/2.0) * diagram.DefaultBoxWidth * 1.5
			y := float64(level) * diagram.DefaultBoxHeigth * 2.0
			shape.X = x + leftOffset
			shape.Y = y + topOffset

			// Base Case: Leaf Node (no sub-tasks)
			if len(task.SubTasks) == 0 {
				return 1.0
			}

			// Recursive Step: Layout Children
			// We stack children horizontally starting from startX
			currentChildX := startX
			for _, subTask := range task.SubTasks {
				// Recursively layout the child and get its actual visual width
				childWidth := layout(subTask, currentChildX, level+1)

				// Advance the X cursor by the child's width for the next sibling
				currentChildX += childWidth

				compositionShape := newConcreteAssociation(task, subTask, &diagram.TaskComposition_Shapes)
				compositionShape.CornerOffsetRatio = 1.5
				_ = compositionShape
			}
			return float64(task.GetComputedWidth())
		}
		width := 0.0
		for _, subTask := range project.RootTasks {
			width += layout(subTask, width, 0)
		}

		// 5. Unstage the diagram and re-stage it with StageBranch to include all new shapes recursively
		diagram.Unstage(stager.stage)
		StageBranch(stager.stage, diagram)

		// Save changes
		stager.stage.Commit()
	}
}
