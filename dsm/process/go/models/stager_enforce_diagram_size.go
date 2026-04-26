package models

func (stager *Stager) enforceDiagramSize() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*DiagramProcess](stager.stage) {

		diagram.Width = 0
		diagram.Height = 0

		// parse all concrete shapes in the diagram that are not links
		//  to compute the size of the diagram
		// the size of the diagram is the max of the position of the shapes + their size
		needCommit = updateDiagramSize(diagram.Process_Shapes, &diagram.Width, &diagram.Height) || needCommit
		needCommit = updateDiagramSize(diagram.Participant_Shapes, &diagram.Width, &diagram.Height) || needCommit

		// add a margin to the diagram size
		margin := 300.0
		diagram.Width += margin
		diagram.Height += margin
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
