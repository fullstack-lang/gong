package models

// deletePBSShapes
func (stager *Stager) deletePBSShapes(diagram *Diagram) {
	for _, shape := range diagram.Product_Shapes {
		shape.Unstage(stager.stage)
	}
	diagram.Product_Shapes = nil

	for _, shape := range diagram.ProductComposition_Shapes {
		shape.Unstage(stager.stage)
	}
	diagram.ProductComposition_Shapes = nil
}

func (stager *Stager) deleteWBSShapes(diagram *Diagram) {
	for _, shape := range diagram.Task_Shapes {
		shape.Unstage(stager.stage)
	}
	diagram.Task_Shapes = nil

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
}
