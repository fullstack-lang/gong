package models

import (
	"fmt"
	"log"
	"sort"
	"time"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

var objectNumber int

func (stager *Stager) treeSimulation() {

	stager.treeObjectsSimulationStage.Reset()

	treeInstance := new(tree.Tree).Stage(stager.treeObjectsSimulationStage)
	treeInstance.Name = string(ObjectTreeName)

	for _, stateMachine := range stager.architecture.StateMachines {

		nodeForAddButton := new(tree.Node).Stage(stager.treeObjectsSimulationStage)
		nodeForAddButton.Name = stateMachine.Name
		treeInstance.RootNodes = append(treeInstance.RootNodes, nodeForAddButton)

		addButton := (&tree.Button{
			Name: "Object" + " " + string(buttons.BUTTON_add),
			Icon: string(buttons.BUTTON_add),
			OnClick: func() {
				for object_ := range *GetGongstructInstancesSet[Object](stager.stage) {
					object_.IsSelected = false
				}

				objectNumber++
				if objectNumber == 13 {
					objectNumber++
				}
				nbObject := fmt.Sprintf("%02d/MI", objectNumber)

				if stateMachine.InitialState == nil {
					log.Println("State Machine", stateMachine.Name, "has no Start State")
					return
				}

				(&Object{
					Name: nbObject +
						" DOF/ " + time.Now().Add(time.Hour*24).Format("2006-01-02") +
						" DEP/ " + time.Now().Add(time.Hour*24).Format("150405"),
					DOF:        time.Now().Add(time.Hour * 24),
					State:      stateMachine.InitialState,
					IsSelected: true,
				}).Stage(stager.stage)

				stager.stage.Commit()
			},
		}).Stage(stager.treeObjectsSimulationStage)
		nodeForAddButton.Buttons = append(nodeForAddButton.Buttons, addButton)

		objects := stager.map_stateMachine_objects[stateMachine]

		sort.Slice(objects,
			func(i, j int) bool {
				return objects[i].DOF.After(objects[j].DOF)
			})

		for _, object := range objects {
			nodeObject := new(tree.Node).Stage(stager.treeObjectsSimulationStage)
			nodeObject.Name = object.Name
			nodeObject.HasCheckboxButton = true
			nodeObject.IsChecked = object.IsSelected
			nodeObject.IsExpanded = true
			nodeObject.OnUpdate = func(_ *tree.Stage, stagedNode, frontNode *tree.Node) {
				if frontNode.IsChecked && !stagedNode.IsChecked {
					for object_ := range *GetGongstructInstancesSet[Object](stager.stage) {
						object_.IsSelected = false
					}
					object.IsSelected = true
					stagedNode.IsChecked = frontNode.IsChecked
					stager.stage.Commit()
				}
				if !frontNode.IsChecked && stagedNode.IsChecked {
					object.IsSelected = false
					stagedNode.IsChecked = frontNode.IsChecked

					for object_ := range *GetGongstructInstancesSet[Object](stager.stage) {
						object_.IsSelected = false
					}
					stager.stage.Commit()
				}
			}

			treeInstance.RootNodes = append(treeInstance.RootNodes, nodeObject)

			nodeState := new(tree.Node).Stage(stager.treeObjectsSimulationStage)
			nodeState.Name = "State: " + object.State.Name
			nodeObject.Children = append(nodeObject.Children, nodeState)

			// ajout d'un bouton delete
			deleteButton := (&tree.Button{
				Name: "State" + " " + string(buttons.BUTTON_delete),
				Icon: string(buttons.BUTTON_delete),
				OnClick: func() {
					object.Unstage(stager.stage)

					for _, message := range object.Messages {
						message.Unstage(stager.stage)
					}
					stager.stage.Commit()
				},
			}).Stage(stager.treeObjectsSimulationStage)
			nodeObject.Buttons = append(nodeObject.Buttons, deleteButton)

			for _, message := range object.Messages {
				nodeMessage := new(tree.Node).Stage(stager.treeObjectsSimulationStage)

				nodeMessage.Name = message.Name
				nodeMessage.HasCheckboxButton = false
				nodeMessage.IsChecked = message.IsSelected

				nodeObject.Children = append(nodeObject.Children, nodeMessage)
			}
		}
	}

	stager.treeObjectsSimulationStage.Commit()
}
