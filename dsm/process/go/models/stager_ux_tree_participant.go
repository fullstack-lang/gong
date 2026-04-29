package models

import (
	"log"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeParticipants(
	diagramProcess *DiagramProcess,
	participant *Participant,
	parentNode *tree.Node) {

	stage := stager.stage

	// find the shape (if any)
	shape, ok := diagramProcess.map_Participant_ParticipantShape[participant]
	node := &tree.Node{
		Name:              participant.Name,
		IsExpanded:        slices.Contains(diagramProcess.ParticipantWhoseNodeIsExpanded, participant),
		IsNodeClickable:   true,
		IsInEditMode:      participant.isInRenameMode,
		HasCheckboxButton: true,
		IsChecked:         ok,
	}
	parentNode.Children = append(parentNode.Children, node)
	addRenameButton(participant, node, stager)

	if shape, ok := diagramProcess.map_Participant_ParticipantShape[participant]; ok {
		node.IsChecked = true
		visibilityButton := &tree.Button{
			Name:            diagramProcess.GetName(),
			Icon:            string(buttons.BUTTON_visibility_off),
			ToolTipText:     "Hide from diagram",
			HasToolTip:      true,
			ToolTipPosition: tree.Right,
			OnClick: func() {
				shape.SetIsHidden(!shape.GetIsHidden())
				stage.Commit()
			},
		}
		if shape.GetIsHidden() {
			visibilityButton.Icon = string(buttons.BUTTON_visibility)
			visibilityButton.ToolTipText = "Show on diagram"
		}
		node.Buttons = append(node.Buttons, visibilityButton)
	}

	node.OnUpdate = func(_ *tree.Stage, stagedNode, frontNode *tree.Node) {
		if frontNode.Name != stagedNode.Name {
			participant.Name = frontNode.Name
			participant.isInRenameMode = false
			stage.Commit()
			return
		}
		if frontNode.IsChecked && !stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if shape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			shape = newShapeToDiagram(participant, diagramProcess, &diagramProcess.Participant_Shapes, stage)

			stage.Commit()
			return
		}
		if !frontNode.IsChecked && stagedNode.IsChecked {
			stagedNode.IsChecked = frontNode.IsChecked
			if shape == nil {
				log.Panic("remove a non existing shape to product")
			}
			shape.UnstageVoid(stage)
			idx := slices.Index(diagramProcess.Participant_Shapes, shape)
			diagramProcess.Participant_Shapes = slices.Delete(diagramProcess.Participant_Shapes, idx, idx+1)
			stage.Commit()
			return
		}
		if frontNode.IsExpanded != stagedNode.IsExpanded {
			if frontNode.IsExpanded {
				if !slices.Contains(diagramProcess.ParticipantWhoseNodeIsExpanded, participant) {
					diagramProcess.ParticipantWhoseNodeIsExpanded = append(diagramProcess.ParticipantWhoseNodeIsExpanded, participant)
				}
			} else {
				if idx := slices.Index(diagramProcess.ParticipantWhoseNodeIsExpanded, participant); idx != -1 {
					diagramProcess.ParticipantWhoseNodeIsExpanded = slices.Delete(diagramProcess.ParticipantWhoseNodeIsExpanded, idx, idx+1)
				}
			}
			stage.Commit()
			return
		}
		stager.probeForm.FillUpFormFromGongstruct(participant, GetPointerToGongstructName[*Participant]())
		stage.Commit()
	}

	tasksNode := &tree.Node{
		Name:            "Tasks",
		FontStyle:       tree.ITALIC,
		IsExpanded:      participant.IsTasksNodeExpanded,
		IsNodeClickable: true,
	}
	node.Children = append(node.Children, tasksNode)
	tasksNode.OnClick = func(frontNode *tree.Node) {
		if frontNode.IsExpanded != participant.IsTasksNodeExpanded {
			participant.IsTasksNodeExpanded = frontNode.IsExpanded
			stager.stage.Commit()
			return
		}
	}

	for _, task := range participant.Tasks {
		stager.treetasks(diagramProcess, task, tasksNode, &diagramProcess.TasksWhoseNodeIsExpanded)
	}

	// loook forward to https://github.com/golang/go/issues/61731
	// proposal: spec: support type inference on generic structs
	conf := addItemButtonConfiguration[
		Task, *Task, // AT, PAT (Added Element)
		Participant, *Participant, // ParentAT, PParentAT (Parent Element)
		TaskShape, *TaskShape, // CT, PCT (Concrete Shape)
		ControlFlowShape, *ControlFlowShape, // ACT, PACT (Association Shape)
	]{
		parentNode:                         tasksNode,
		sliceForNewAddedItem:               &participant.Tasks,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &participant.IsTasksNodeExpanded,
		parentElement:                      participant,
		isWithAdditionOfShape:              true,
		receivingDiagram:                   diagramProcess,
		sliceForNewAddedShape:              &diagramProcess.TaskShapes,
		isWithAdditionOfAssociationShape:   false,
		sliceForNewCompositionShapes:       &diagramProcess.ControlFlowShape,
	}

	addAddButton(stager, conf)

	addAddItemButtonSimple(
		stager,
		&diagramProcess.ParticipantWhoseNodeIsExpanded,
		participant,
		nil,
		node,
		&participant.Tasks)
}
