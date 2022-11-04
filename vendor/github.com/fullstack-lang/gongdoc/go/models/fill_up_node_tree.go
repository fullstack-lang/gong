package models

import (
	"log"
)

func FillUpNodeTree(pkgelt *DiagramPackage) {

	onNodeCallbackStruct := new(NodeCallbacksSingloton)
	FillUpDiagramNodeTree(pkgelt, onNodeCallbackStruct)
	FillUpTreeOfIdentifiers(pkgelt, onNodeCallbackStruct)
	updateNodesStates(&Stage, onNodeCallbackStruct)

	// set callbacks on node updates
	Stage.OnAfterNodeUpdateCallback = onNodeCallbackStruct
	Stage.OnAfterNodeCreateCallback = onNodeCallbackStruct
	Stage.OnAfterNodeDeleteCallback = onNodeCallbackStruct

	log.Printf("Parse found %d diagrams\n", len(pkgelt.Classdiagrams))
	log.Printf("Server ready to serve on http://localhost:8080/")
}
