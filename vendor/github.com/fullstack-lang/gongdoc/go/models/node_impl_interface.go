package models

type NodeImplInterface interface {
	OnAfterUpdate(stage *StageStruct, stagedNode, frontNode *Node)
	OnAfterDelete(stage *StageStruct, stagedNode, frontNode *Node)
	HasToBeChecked() bool
	SetHasToBeCheckedValue(value bool)
	HasToBeDisabled() bool
	SetHasToBeDisabledValue(value bool)
}
