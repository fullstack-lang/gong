package models

import gongtree_models "github.com/fullstack-lang/gong/lib/tree/go/models"

type Object_Tree_Proxy struct {
	stager *Stager
	object *Object
}

// OnAfterUpdate implements models.NodeImplInterface.
func (proxy *Object_Tree_Proxy) OnAfterUpdate(stage *gongtree_models.Stage, stagedNode *gongtree_models.Node, frontNode *gongtree_models.Node) {
	if frontNode.IsChecked && !stagedNode.IsChecked {
		for object_ := range *GetGongstructInstancesSet[Object](proxy.stager.stage) {
			object_.IsSelected = false
		}
		proxy.object.IsSelected = true
		stagedNode.IsChecked = frontNode.IsChecked
		proxy.stager.stage.Commit()
	}
	if !frontNode.IsChecked && stagedNode.IsChecked {
		proxy.object.IsSelected = false
		stagedNode.IsChecked = frontNode.IsChecked

		for object_ := range *GetGongstructInstancesSet[Object](proxy.stager.stage) {
			object_.IsSelected = false
		}
		proxy.stager.stage.Commit()
	}
}
