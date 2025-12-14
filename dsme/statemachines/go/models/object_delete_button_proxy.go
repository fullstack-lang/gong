package models

import (
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type ObjectDeleteButtonProxy struct {
	stager *Stager
	object *Object
}

// ButtonUpdated implements models.ButtonImplInterface.
func (proxy *ObjectDeleteButtonProxy) ButtonUpdated(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {

	proxy.object.Unstage(proxy.stager.stage)

	for _, message := range proxy.object.Messages {
		message.Unstage(proxy.stager.stage)
	}
	proxy.stager.stage.Commit()
}
