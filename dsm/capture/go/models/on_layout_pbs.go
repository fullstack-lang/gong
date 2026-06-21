package models

var (
	leftOffset = 50.0
	topOffset  = 50.0
)

// onLayoutPBS resets the diagram, creates and adds all possible deliverable-related concrete instances to the diagram,
// and organizes them in a breakdown structure
func onLayoutPBS(stager *Stager, diagram *Diagram) func() {
	return func() {
		// 1. Reset the diagram: remove all shapes from the stage to start fresh
		diagram.Deliverable_Shapes = []*DeliverableShape{}
		diagram.Concern_Shapes = []*ConcernShape{}
		diagram.Note_Shapes = []*NoteShape{}
		// Note: Task and Note shapes are not considered in PBS layout

		// 2. Identify Project
		project := stager.stage.Library_Diagrams_reverseMap[diagram]

		// layout is the function that
		// - layout the deliverable at startX and level
		// - returns the startX for the next deliverable at the same level
		// - call layout on all subDeliverables
		var layout func(p *Deliverable, startX float64, level int) float64
		layout = func(deliverable *Deliverable, startX float64, level int) float64 {
			// Create Shape
			shape := newShapeToDiagram(deliverable, diagram, &diagram.Deliverable_Shapes, stager.stage)
			x := (startX + float64(deliverable.GetComputedWidth()-1)/2.0) * diagram.DefaultBoxWidth * 1.5
			y := float64(level) * diagram.DefaultBoxHeigth * 2.0
			shape.X = x + leftOffset
			shape.Y = y + topOffset

			// Base Case: Leaf Node (no sub-deliverables)
			if len(deliverable.SubDeliverables) == 0 {
				return 1.0
			}

			// Recursive Step: Layout Children
			// We stack children horizontally starting from startX
			currentChildX := startX
			for _, subDeliverable := range deliverable.SubDeliverables {
				// Recursively layout the child and get its actual visual width
				childWidth := layout(subDeliverable, currentChildX, level+1)

				// Advance the X cursor by the child's width for the next sibling
				currentChildX += childWidth

				compositionShape := newConcreteAssociation(deliverable, subDeliverable, &diagram.DeliverableComposition_Shapes)
				compositionShape.CornerOffsetRatio = 1.5
				_ = compositionShape
			}
			return float64(deliverable.GetComputedWidth())
		}
		width := 0.0
		for _, subDeliverable := range project.RootDeliverables {
			width += layout(subDeliverable, width, 0)
		}

		// 5. Unstage the diagram and re-stage it with StageBranch to include all new shapes recursively
		diagram.Unstage(stager.stage)
		StageBranch(stager.stage, diagram)

		// Save changes
		stager.stage.Commit()
	}
}
