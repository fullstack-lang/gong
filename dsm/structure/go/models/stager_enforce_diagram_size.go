package models

func (stager *Stager) enforceDiagramSize() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*DiagramStructure](stager.stage) {

		width := 0.0
		height := 0.0

		// parse all concrete shapes in the diagram that are not links
		//  to compute the size of the diagram
		// the size of the diagram is the max of the position of the shapes + their size
		updateDiagramSize(diagram.System_Shapes, &width, &height)
		updateDiagramSize(diagram.Part_Shapes, &width, &height)
		updateDiagramSize(diagram.ExternalPart_Shapes, &width, &height)
		updateDiagramSize(diagram.Port_Shapes, &width, &height)
		updateDiagramSize(diagram.Note_Shapes, &width, &height)

		// add a margin to the diagram size
		margin := 300.0
		width += margin
		height += margin

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

func updateDiagramSize[T RectShapeInterface](shapes []T, width, height *float64) (needCommit bool) {
	for _, shape := range shapes {
		if shape.GetX()+shape.GetWidth() > *width {
			*width = shape.GetX() + shape.GetWidth()
			needCommit = true
		}
		if shape.GetY()+shape.GetHeight() > *height {
			*height = shape.GetY() + shape.GetHeight()
			needCommit = true
		}
	}
	return
}
