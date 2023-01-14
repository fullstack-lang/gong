package models

import (
	"log"
)

func FillUpNodeTree(pkgelt *DiagramPackage) {

	// a node tree is agnostic of the node types it manages
	// therefore, a callback functiion is necessary
	onNodeCallbackStruct := new(NodeCallbacksSingloton)

	FillUpDiagramNodeTree(pkgelt, onNodeCallbackStruct)
	FillUpTreeOfIdentifiers(pkgelt, onNodeCallbackStruct)

	updateNodesStates(&Stage, onNodeCallbackStruct)

	// set callbacks on node updates
	Stage.OnAfterNodeUpdateCallback = onNodeCallbackStruct
	Stage.OnAfterNodeCreateCallback = onNodeCallbackStruct
	Stage.OnAfterNodeDeleteCallback = onNodeCallbackStruct

	log.Printf("Parse found %d diagrams\n", len(pkgelt.Classdiagrams))
}
