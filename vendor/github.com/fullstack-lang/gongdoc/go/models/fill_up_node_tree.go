package models

import (
	"log"
)

func FillUpNodeTree(diagramPackage *DiagramPackage) {

	// a node tree is agnostic of the node types it manages
	// therefore, a callback functiion is necessary
	onNodeCallbackStruct := new(NodeCallbacksSingloton)
	onNodeCallbackStruct.diagramPackage = diagramPackage

	FillUpDiagramNodeTree(diagramPackage, onNodeCallbackStruct)
	FillUpTreeOfIdentifiers(diagramPackage, onNodeCallbackStruct)

	updateNodesStates(&Stage, onNodeCallbackStruct)

	// set callbacks on node updates
	Stage.OnAfterNodeUpdateCallback = onNodeCallbackStruct
	Stage.OnAfterNodeCreateCallback = onNodeCallbackStruct
	Stage.OnAfterNodeDeleteCallback = onNodeCallbackStruct

	log.Printf("Parse found %d diagrams\n", len(diagramPackage.Classdiagrams))
}
