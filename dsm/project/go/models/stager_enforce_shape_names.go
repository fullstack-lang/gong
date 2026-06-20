package models

import (
	"fmt"
	"time"
)

func (stager *Stager) enforceShapeNames() (needCommit bool) {
	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		needCommit = enforceNodeShapeName(stager, diagram.Name, diagram.Product_Shapes) || needCommit
		needCommit = enforceNodeShapeName(stager, diagram.Name, diagram.Task_Shapes) || needCommit
		needCommit = enforceNodeShapeName(stager, diagram.Name, diagram.Note_Shapes) || needCommit
		needCommit = enforceNodeShapeName(stager, diagram.Name, diagram.Resource_Shapes) || needCommit
		needCommit = enforceNodeShapeName(stager, diagram.Name, diagram.TaskGroupShapes) || needCommit

		needCommit = enforceAssociationShapeName(stager, diagram.Name, diagram.ProductComposition_Shapes) || needCommit
		needCommit = enforceAssociationShapeName(stager, diagram.Name, diagram.TaskComposition_Shapes) || needCommit
		needCommit = enforceAssociationShapeName(stager, diagram.Name, diagram.TaskInputShapes) || needCommit
		needCommit = enforceAssociationShapeName(stager, diagram.Name, diagram.TaskOutputShapes) || needCommit
		needCommit = enforceAssociationShapeName(stager, diagram.Name, diagram.NoteProductShapes) || needCommit
		needCommit = enforceAssociationShapeName(stager, diagram.Name, diagram.NoteTaskShapes) || needCommit
		needCommit = enforceAssociationShapeName(stager, diagram.Name, diagram.NoteResourceShapes) || needCommit
		needCommit = enforceAssociationShapeName(stager, diagram.Name, diagram.ResourceComposition_Shapes) || needCommit
		needCommit = enforceAssociationShapeName(stager, diagram.Name, diagram.ResourceTaskShapes) || needCommit
	}
	return
}

func enforceNodeShapeName[T ConcreteType](stager *Stager, diagramName string, shapes []T) (needCommit bool) {
	for _, shape := range shapes {
		abstractElement := shape.GetAbstractElement()
		if abstractElement != nil {
			expectedName := fmt.Sprintf("%s-%s", diagramName, abstractElement.GetName())
			if shape.GetName() != expectedName {
				shape.SetName(expectedName)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Renamed shape to %s", expectedName))
				}
			}
		}
	}
	return
}

func enforceAssociationShapeName[T AssociationConcreteType](stager *Stager, diagramName string, shapes []T) (needCommit bool) {
	for _, shape := range shapes {
		startElement := shape.GetAbstractStartElement()
		endElement := shape.GetAbstractEndElement()
		if startElement != nil && endElement != nil {
			expectedName := fmt.Sprintf("%s-%s-%s", diagramName, startElement.GetName(), endElement.GetName())
			if shape.GetName() != expectedName {
				shape.SetName(expectedName)
				needCommit = true
				if stager.probeForm != nil {
					stager.probeForm.AddNotification(time.Now(), fmt.Sprintf("Renamed shape to %s", expectedName))
				}
			}
		}
	}
	return
}
