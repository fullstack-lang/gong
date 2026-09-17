// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (button *Button) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterButtonCreateCallback != nil {
		stage.OnAfterButtonCreateCallback.OnAfterCreate(stage, button)
	}
}

func (button *Button) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterButtonUpdateCallback != nil {
		var frontButton *Button
		if front != nil {
			frontButton, _ = front.(*Button)
		}
		stage.OnAfterButtonUpdateCallback.OnAfterUpdate(stage, button, frontButton)
	}
}

func (button *Button) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterButtonDeleteCallback != nil {
		var frontButton *Button
		if front != nil {
			frontButton, _ = front.(*Button)
		}
		stage.OnAfterButtonDeleteCallback.OnAfterDelete(stage, button, frontButton)
	}
}

func (menu *Menu) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMenuCreateCallback != nil {
		stage.OnAfterMenuCreateCallback.OnAfterCreate(stage, menu)
	}
}

func (menu *Menu) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMenuUpdateCallback != nil {
		var frontMenu *Menu
		if front != nil {
			frontMenu, _ = front.(*Menu)
		}
		stage.OnAfterMenuUpdateCallback.OnAfterUpdate(stage, menu, frontMenu)
	}
}

func (menu *Menu) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMenuDeleteCallback != nil {
		var frontMenu *Menu
		if front != nil {
			frontMenu, _ = front.(*Menu)
		}
		stage.OnAfterMenuDeleteCallback.OnAfterDelete(stage, menu, frontMenu)
	}
}

func (node *Node) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNodeCreateCallback != nil {
		stage.OnAfterNodeCreateCallback.OnAfterCreate(stage, node)
	}
}

func (node *Node) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNodeUpdateCallback != nil {
		var frontNode *Node
		if front != nil {
			frontNode, _ = front.(*Node)
		}
		stage.OnAfterNodeUpdateCallback.OnAfterUpdate(stage, node, frontNode)
	}
}

func (node *Node) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNodeDeleteCallback != nil {
		var frontNode *Node
		if front != nil {
			frontNode, _ = front.(*Node)
		}
		stage.OnAfterNodeDeleteCallback.OnAfterDelete(stage, node, frontNode)
	}
}

func (svgicon *SVGIcon) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSVGIconCreateCallback != nil {
		stage.OnAfterSVGIconCreateCallback.OnAfterCreate(stage, svgicon)
	}
}

func (svgicon *SVGIcon) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSVGIconUpdateCallback != nil {
		var frontSVGIcon *SVGIcon
		if front != nil {
			frontSVGIcon, _ = front.(*SVGIcon)
		}
		stage.OnAfterSVGIconUpdateCallback.OnAfterUpdate(stage, svgicon, frontSVGIcon)
	}
}

func (svgicon *SVGIcon) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSVGIconDeleteCallback != nil {
		var frontSVGIcon *SVGIcon
		if front != nil {
			frontSVGIcon, _ = front.(*SVGIcon)
		}
		stage.OnAfterSVGIconDeleteCallback.OnAfterDelete(stage, svgicon, frontSVGIcon)
	}
}

func (tree *Tree) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterTreeCreateCallback != nil {
		stage.OnAfterTreeCreateCallback.OnAfterCreate(stage, tree)
	}
}

func (tree *Tree) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTreeUpdateCallback != nil {
		var frontTree *Tree
		if front != nil {
			frontTree, _ = front.(*Tree)
		}
		stage.OnAfterTreeUpdateCallback.OnAfterUpdate(stage, tree, frontTree)
	}
}

func (tree *Tree) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterTreeDeleteCallback != nil {
		var frontTree *Tree
		if front != nil {
			frontTree, _ = front.(*Tree)
		}
		stage.OnAfterTreeDeleteCallback.OnAfterDelete(stage, tree, frontTree)
	}
}

