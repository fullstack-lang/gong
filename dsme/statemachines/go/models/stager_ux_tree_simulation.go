package models

import (
	"sort"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
)

func (stager *Stager) treeObjectSimulation() {

	stager.treeObjectsSimulationStage.Reset()

	treeInstance := new(tree.Tree).Stage(stager.treeObjectsSimulationStage)
	treeInstance.Name = string(ObjectTreeName)

	for _, stateMachine := range stager.architecture.StateMachines {

		nodeForAddButton := new(tree.Node).Stage(stager.treeObjectsSimulationStage)
		nodeForAddButton.Name = stateMachine.Name
		treeInstance.RootNodes = append(treeInstance.RootNodes, nodeForAddButton)

		addButton := (&tree.Button{
			Name: "Object" + " " + string(buttons.BUTTON_add),
			Icon: string(buttons.BUTTON_add)}).Stage(stager.treeObjectsSimulationStage)
		nodeForAddButton.Buttons = append(nodeForAddButton.Buttons, addButton)
		addButton.Impl = &ObjectAddButtonProxy{
			stager:       stager,
			stateMachine: stateMachine,
		}

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

			p := new(Object_Tree_Proxy)
			p.stager = stager
			p.object = object
			object.Proxy = p
			nodeObject.Impl = p

			treeInstance.RootNodes = append(treeInstance.RootNodes, nodeObject)

			nodeState := new(tree.Node).Stage(stager.treeObjectsSimulationStage)
			nodeState.Name = "State: " + object.State.Name
			nodeObject.Children = append(nodeObject.Children, nodeState)

			// ajout d'un bouton delete
			deleteButton := (&tree.Button{
				Name: "State" + " " + string(buttons.BUTTON_delete),
				Icon: string(buttons.BUTTON_delete)}).Stage(stager.treeObjectsSimulationStage)
			nodeObject.Buttons = append(nodeObject.Buttons, deleteButton)
			deleteButton.Impl = &ObjectDeleteButtonProxy{
				stager: stager,
				object: object,
			}

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
