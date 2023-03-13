package models

type NodeImplInterface interface {

	// OnAfterUpdate function is called each time a node is modified
	OnAfterUpdate(stage *StageStruct, stagedNode, frontNode *Node)

	// OnAfterUpdate function is called each time a node is deleted
	OnAfterDelete(stage *StageStruct, stagedNode, frontNode *Node)
}
