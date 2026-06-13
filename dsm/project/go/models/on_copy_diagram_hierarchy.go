package models

func onCopyDiagram(stager *Stager, diagramHierarchy *DiagramHierarchy) func() {
	return func() {
		newDiagramHierarchy := new(DiagramHierarchy)
		newDiagramHierarchy.Name = diagramHierarchy.Name + " copy"
		newDiagramHierarchy.IsEditable_ = true

		library := stager.stage.Library_Diagrams_reverseMap[diagramHierarchy]
		Append(&library.Diagrams, newDiagramHierarchy)

		// Copy nodes
		for _, node := range diagramHierarchy.GetNodes() {
			newNode := node.GongCopy().(ConcreteType)
			newNode.SetAbstractElement(node.GetAbstractElement())
			newDiagramHierarchy.AddNode(newNode)
		}

		// Copy links
		for _, link := range diagramHierarchy.GetLinks() {
			newLink := link.GongCopy().(AssociationConcreteType)
			newLink.SetAbstractStartElement(link.GetAbstractStartElement())
			newLink.SetAbstractEndElement(link.GetAbstractEndElement())
			newDiagramHierarchy.AddLink(newLink)
		}

		StageBranch(stager.stage, newDiagramHierarchy)
		stager.stage.Commit()
	}
}
