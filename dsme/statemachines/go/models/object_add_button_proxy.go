package models

import (
	"fmt"
	"log"
	"time"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

type ObjectAddButtonProxy struct {
	stager       *Stager
	stateMachine *StateMachine
}

// ButtonUpdated implements models.ButtonImplInterface.
func (p *ObjectAddButtonProxy) ButtonUpdated(stage *tree.Stage, button *tree.Button, updatedButton *tree.Button) {

	for object_ := range *GetGongstructInstancesSet[Object](p.stager.stage) {
		object_.IsSelected = false
	}

	objectNumber++
	if objectNumber == 13 {
		objectNumber++
	}
	nbObject := fmt.Sprintf("%02d/MI", objectNumber)

	if p.stateMachine.InitialState == nil {
		log.Println("State Machine", p.stateMachine.Name, "has no Start State")
		return
	}

	(&Object{
		Name: nbObject +
			" DOF/ " + time.Now().Add(time.Hour*24).Format("2006-01-02") +
			" DEP/ " + time.Now().Add(time.Hour*24).Format("150405"),
		DOF:        time.Now().Add(time.Hour * 24),
		State:      p.stateMachine.InitialState,
		IsSelected: true,
	}).Stage(p.stager.stage)

	p.stager.stage.Commit()
}

var objectNumber int
