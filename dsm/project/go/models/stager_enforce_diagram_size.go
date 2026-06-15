package models

func (stager *Stager) enforceDiagramSize() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {

		width := 0.0
		height := 0.0

		// parse all concrete shapes in the diagram that are not links
		//  to compute the size of the diagram
		// the size of the diagram is the max of the position of the shapes + their size
		for _, shape := range diagram.Product_Shapes {
			if shape.X+shape.Width > width {
				width = shape.X + shape.Width
			}
			if shape.Y+shape.Height > height {
				height = shape.Y + shape.Height
			}
		}

		for _, shape := range diagram.Task_Shapes {
			if shape.X+shape.Width > width {
				width = shape.X + shape.Width
			}
			if shape.Y+shape.Height > height {
				height = shape.Y + shape.Height
			}
		}

		for _, shape := range diagram.Note_Shapes {
			if shape.X+shape.Width > width {
				width = shape.X + shape.Width
			}
			if shape.Y+shape.Height > height {
				height = shape.Y + shape.Height
			}
		}
		for _, shape := range diagram.Resource_Shapes {
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

		if diagram.IsTimeDiagram {
			if width < diagram.XRightMargin+margin {
				width = diagram.XRightMargin + margin
			}

			var nbVisibleTaskGroups int
			for _, taskGroupShape := range diagram.TaskGroupShapes {
				if !taskGroupShape.IsHidden {
					nbVisibleTaskGroups++
				}
			}

			timeDiagramHeight := diagram.YTopMargin + diagram.LaneHeight*float64(nbVisibleTaskGroups)
			timeDiagramHeight += diagram.DateYOffset + margin // margin for date labels
			if height < timeDiagramHeight {
				height = timeDiagramHeight
			}
		}

		if width != diagram.Width {
			diagram.Width = width
			needCommit = true
		}

		if height != diagram.Height {
			diagram.Height = height
			needCommit = true
		}
	}

	return
}
