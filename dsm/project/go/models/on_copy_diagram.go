package models

func onCopyDiagram(stager *Stager, diagram *Diagram) func() {
	return func() {
		newDiagram := new(Diagram)
		newDiagram.Name = diagram.Name + " copy"
		newDiagram.IsEditable_ = true

		library := stager.stage.Library_Diagrams_reverseMap[diagram]
		Append(&library.Diagrams, newDiagram)

		// Copy nodes
		for _, node := range diagram.GetNodes() {
			newNode := node.GongCopy().(ConcreteType)
			newNode.SetAbstractElement(node.GetAbstractElement())
			newDiagram.AddNode(newNode)
		}

		// Copy links
		for _, link := range diagram.GetLinks() {
			newLink := link.GongCopy().(AssociationConcreteType)
			newLink.SetAbstractStartElement(link.GetAbstractStartElement())
			newLink.SetAbstractEndElement(link.GetAbstractEndElement())
			newDiagram.AddLink(newLink)
		}

		StageBranch(stager.stage, newDiagram)
		stager.stage.Commit()
	}
}
