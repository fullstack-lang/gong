package models

func (stager *Stager) enforceDiagramSize() (needCommit bool) {
	for _, diagramHierarchy := range GetGongstrucsSorted[*DiagramHierarchy](stager.stage) {

		width := 0.0
		height := 0.0

		// parse all concrete shapes in the diagram that are not links
		//  to compute the size of the diagram
		// the size of the diagram is the max of the position of the shapes + their size
		for _, shape := range diagramHierarchy.Product_Shapes {
			if shape.X+shape.Width > width {
				width = shape.X + shape.Width
			}
			if shape.Y+shape.Height > height {
				height = shape.Y + shape.Height
			}
		}

		for _, shape := range diagramHierarchy.Task_Shapes {
			if shape.X+shape.Width > width {
				width = shape.X + shape.Width
			}
			if shape.Y+shape.Height > height {
				height = shape.Y + shape.Height
			}
		}

		for _, shape := range diagramHierarchy.Note_Shapes {
			if shape.X+shape.Width > width {
				width = shape.X + shape.Width
			}
			if shape.Y+shape.Height > height {
				height = shape.Y + shape.Height
			}
		}
		for _, shape := range diagramHierarchy.Resource_Shapes {
			if shape.X+shape.Width > width {
				width = shape.X + shape.Width
			}
			if shape.Y+shape.Height > height {
				height = shape.Y + shape.Height
			}
		}
		for _, shape := range diagramHierarchy.MilestoneShapes {
			if shape.X+shape.Width > width {
				width = shape.X + shape.Width
			}
			if shape.Y+shape.Height > height {
				height = shape.Y + shape.Height
			}
		}
		// add a margin to the diagram size
		margin := 100.0
		width += margin
		height += margin

		if diagramHierarchy.IsTimeDiagram {
			if width < diagramHierarchy.XRightMargin+margin {
				width = diagramHierarchy.XRightMargin + margin
			}

			var nbVisibleTaskGroups int
			for _, taskGroupShape := range diagramHierarchy.TaskGroupShapes {
				if !taskGroupShape.IsHidden {
					nbVisibleTaskGroups++
				}
			}

			timeDiagramHeight := diagramHierarchy.YTopMargin + diagramHierarchy.LaneHeight*float64(nbVisibleTaskGroups)
			timeDiagramHeight += diagramHierarchy.DateYOffset + margin // margin for date labels
			if height < timeDiagramHeight {
				height = timeDiagramHeight
			}
		}

		if width != diagramHierarchy.Width {
			diagramHierarchy.Width = width
			needCommit = true
		}

		if height != diagramHierarchy.Height {
			diagramHierarchy.Height = height
			needCommit = true
		}
	}

	return
}
