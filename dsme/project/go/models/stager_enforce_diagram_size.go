package models

func (stager *Stager) enforceDiagramSize() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {

		// parse all concrete shapes in the diagram that are not links
		//  to compute the size of the diagram
		// the size of the diagram is the max of the position of the shapes + their size
		for _, shape := range diagram.Product_Shapes {
			if shape.X+shape.Width > diagram.Width {
				diagram.Width = shape.X + shape.Width
				needCommit = true
			}
			if shape.Y+shape.Height > diagram.Height {
				diagram.Height = shape.Y + shape.Height
				needCommit = true
			}
		}

		for _, shape := range diagram.Task_Shapes {
			if shape.X+shape.Width > diagram.Width {
				diagram.Width = shape.X + shape.Width
				needCommit = true
			}
			if shape.Y+shape.Height > diagram.Height {
				diagram.Height = shape.Y + shape.Height
				needCommit = true
			}
		}

		for _, shape := range diagram.Note_Shapes {
			if shape.X+shape.Width > diagram.Width {
				diagram.Width = shape.X + shape.Width
				needCommit = true
			}
			if shape.Y+shape.Height > diagram.Height {
				diagram.Height = shape.Y + shape.Height
				needCommit = true
			}
		}
		for _, shape := range diagram.Resource_Shapes {
			if shape.X+shape.Width > diagram.Width {
				diagram.Width = shape.X + shape.Width
				needCommit = true
			}
			if shape.Y+shape.Height > diagram.Height {
				diagram.Height = shape.Y + shape.Height
				needCommit = true
			}
		}
		// add a margin to the diagram size
		margin := 100.0
		diagram.Width += margin
		diagram.Height += margin
	}

	return
}
