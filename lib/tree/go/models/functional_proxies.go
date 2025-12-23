package models

// If you want to avoid cluttering your package with many tiny struct definitions (like AddProductButtonNodeProxy),
// you can define one generic adapter struct that delegates logic to a function.
//
// This allows you to write the logic "anonymously" (inline) using a closure.
/*
	addButton := &tree.Button{
		Name: "Product" + " " + string(buttons.BUTTON_add),
		Icon: string(buttons.BUTTON_add),
	}
	node.Buttons = append(node.Buttons, addButton)

	// Use the generic proxy and define logic inline
	addButton.Impl = &FunctionalButtonProxy{
		OnUpdated: func(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {
			...
		},
	}
*/
// Define a function signature that matches the interface method logic
type ButtonUpdatedFunc func(stage *Stage, button *Button, updatedButton *Button)

// Generic Proxy that implements models.ButtonImplInterface
type FunctionalButtonProxy struct {
	OnUpdated ButtonUpdatedFunc
}

// Implement the interface method
func (p *FunctionalButtonProxy) ButtonUpdated(stage *Stage, button *Button, updatedButton *Button) {
	if p.OnUpdated != nil {
		p.OnUpdated(stage, button, updatedButton)
	}
}

// Define a function signature that matches the interface method logic
type NodeOnAfterUpdateFunc func(stage *Stage, stagedNode *Node, frontNode *Node)

type FunctionalNodeProxy struct {
	OnUpdate NodeOnAfterUpdateFunc
}

// Implement the interface method
func (p *FunctionalNodeProxy) OnAfterUpdate(stage *Stage, stagedNode *Node, frontNode *Node) {
	if p.OnUpdate != nil {
		p.OnUpdate(stage, stagedNode, frontNode)
	}
}
