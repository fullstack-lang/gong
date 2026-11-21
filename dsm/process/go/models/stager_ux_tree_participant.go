package models

import (
	"log"
	"slices"

	buttons "github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeParticipantWithinLibrary(
	library *Library,
	participant *Participant,
	parentNode *tree.Node,
) {
	participantNode := &tree.Node{
		Name:            participant.GetName(),
		IsExpanded:      slices.Contains(library.ParticipantsWhoseNodeIsExpanded, participant),
		IsNodeClickable: true,
		IsInEditMode:    participant.GetIsInRenameMode(),
	}
	parentNode.Children = append(parentNode.Children, participantNode)

	addRenameButton(participant, participantNode, stager)

	participantNode.OnNameChange = stager.onNameChange(participant)
	participantNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, participant, &library.ParticipantsWhoseNodeIsExpanded)
	participantNode.OnClick = onNodeClicked(stager, participant)
}

func (stager *Stager) treeParticipants(
	diagramProcess *DiagramProcess,
	participant *Participant,
	parentNode *tree.Node) {

	stage := stager.stage

	// find the shape (if any)
	shape, ok := diagramProcess.map_Participant_ParticipantShape[participant]
	node := &tree.Node{
		Name:                    participant.Name,
		IsExpanded:              slices.Contains(diagramProcess.ParticipantWhoseNodeIsExpanded, participant),
		IsNodeClickable:         true,
		IsInEditMode:            participant.isInRenameMode,
		HasCheckboxButton:       true,
		IsChecked:               ok,
		CheckboxHasToolTip:      true,
		CheckboxToolTipPosition: tree.Left,
		CheckboxToolTipText: func() string {
			if ok {
				return "Click to remove the participant shape"
			}
			return "Click to create a participant shape for this participant within this diagram"
		}(),
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

	node.OnNameChange = stager.onNameChange(participant)
	node.OnIsExpandedChange = onIsExpandedChangeSlice(stager, participant, &diagramProcess.ParticipantWhoseNodeIsExpanded)
	node.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			if shape != nil {
				log.Panic("adding a shape to an already product shape")
			}
			shape = newShapeToDiagram(participant, diagramProcess, &diagramProcess.Participant_Shapes, stager, node.ClientOnY)

			stage.Commit()
			return
		} else {
			if shape == nil {
				log.Panic("remove a non existing shape to product")
			}
			shape.UnstageVoid(stage)
			idx := slices.Index(diagramProcess.Participant_Shapes, shape)
			diagramProcess.Participant_Shapes = slices.Delete(diagramProcess.Participant_Shapes, idx, idx+1)
			stage.Commit()
			return
		}
	}
	node.OnClick = onNodeClicked(stager, participant)

	// Menu buttons on participant node: Add Task, Add Fork/Join & Add Decision Node
	stager.addCreateDecisionNodeButton(diagramProcess, participant, node, true)
	stager.addCreateForkJoinButton(diagramProcess, participant, node, true)
	confParticipantTaskMenu := ItemAndShapeButtonConfiguration[
		Task, *Task,
		Participant, *Participant,
		TaskShape, *TaskShape,
	]{
		ItemButtonConfiguration: ItemButtonConfiguration[
			Task, *Task,
			Participant, *Participant,
		]{
			parentNode:                         node,
			sliceForNewAddedItem:               &participant.Tasks,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeBySlice,
			parentNodeExpansionSliceEncoding:   &diagramProcess.ParticipantWhoseNodeIsExpanded,
			parentElement:                      participant,
			IsButtonInMenu:                     true,
		},
		receivingDiagram:      diagramProcess,
		sliceForNewAddedShape: &diagramProcess.Task_Shapes,
	}
	callbacksParticipantTaskMenu := addCreateItemAndShapeButton(stager, confParticipantTaskMenu)
	callbacksParticipantTaskMenu.button.Name = "Add Task"
	callbacksParticipantTaskMenu.button.ToolTipText = "Add a Task to \"" + participant.Name + "\""
	callbacksParticipantTaskMenu.OnBeforeCommit = func() {
		participant.IsTasksNodeExpanded = true
	}

	// Allocated Resources
	allocatedResourcesNode := &tree.Node{
		Name:            "Allocated Resources",
		FontStyle:       tree.ITALIC,
		IsExpanded:      participant.IsResourcesNodeExpanded,
		IsNodeClickable: true,
	}
	node.Children = append(node.Children, allocatedResourcesNode)
	allocatedResourcesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&participant.IsResourcesNodeExpanded)

	for _, resource := range participant.Resources {
		stager.treeAllocatedResourceWithinDiagramWithinParticipant(diagramProcess, resource, participant, allocatedResourcesNode)
	}

	// Tasks
	tasksNode := &tree.Node{
		Name:            "Tasks",
		FontStyle:       tree.ITALIC,
		IsExpanded:      participant.IsTasksNodeExpanded,
		IsNodeClickable: true,
	}
	node.Children = append(node.Children, tasksNode)
	tasksNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&participant.IsTasksNodeExpanded)

	for _, task := range participant.Tasks {
		stager.treeTask(diagramProcess, participant, task, tasksNode, &diagramProcess.TasksWhoseNodeIsExpanded)
	}

	// loook forward to https://github.com/golang/go/issues/61731
	// proposal: spec: support type inference on generic structs
	conf := ItemAndShapeButtonConfiguration[
		Task, *Task, // AT, PAT (Added Element)
		Participant, *Participant, // ParentAT, PParentAT (Parent Element)
		TaskShape, *TaskShape, // CT, PCT (Concrete Shape)

	]{
		ItemButtonConfiguration: ItemButtonConfiguration[
			Task, *Task, // AT, PAT (Added Element)
			Participant, *Participant, // ParentAT, PParentAT (Parent Element)
		]{
			parentNode:                         tasksNode,
			sliceForNewAddedItem:               &participant.Tasks,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &participant.IsTasksNodeExpanded,
			parentElement:                      participant,
		},
		receivingDiagram:      diagramProcess,
		sliceForNewAddedShape: &diagramProcess.Task_Shapes,
	}
	addCreateItemAndShapeButton(stager, conf)
	stager.addCreateDecisionNodeButton(diagramProcess, participant, tasksNode, true)
	stager.addCreateForkJoinButton(diagramProcess, participant, tasksNode, true)
	confTasksMenu := conf
	confTasksMenu.IsButtonInMenu = true
	callbacksTasksMenu := addCreateItemAndShapeButton(stager, confTasksMenu)
	callbacksTasksMenu.button.Name = "Add Task"
	callbacksTasksMenu.button.ToolTipText = "Add a Task to \"" + participant.Name + "\""

	controlflowsNode := &tree.Node{
		Name:            "ControlFlows",
		FontStyle:       tree.ITALIC,
		IsExpanded:      participant.IsControlFlowsNodeExpanded,
		IsNodeClickable: true,
	}
	node.Children = append(node.Children, controlflowsNode)
	controlflowsNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&participant.IsControlFlowsNodeExpanded)

	for _, controlflow := range participant.ControlFlows {
		stager.treeControlFlows(diagramProcess, controlflow, controlflowsNode, &diagramProcess.ControlFlowsWhoseNodeIsExpanded)
	}
}

func (stager *Stager) addCreateForkJoinButton(
	diagramProcess *DiagramProcess,
	participant *Participant,
	parentNode *tree.Node,
	isButtonInMenu bool,
) {
	confForkJoin := ItemAndShapeButtonConfiguration[
		Task, *Task,
		Participant, *Participant,
		TaskShape, *TaskShape,
	]{
		ItemButtonConfiguration: ItemButtonConfiguration[
			Task, *Task,
			Participant, *Participant,
		]{
			parentNode:                         parentNode,
			sliceForNewAddedItem:               &participant.Tasks,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &participant.IsTasksNodeExpanded,
			parentElement:                      participant,
			IsButtonInMenu:                     isButtonInMenu,
		},
		receivingDiagram:      diagramProcess,
		sliceForNewAddedShape: &diagramProcess.Task_Shapes,
	}
	callbacksForkJoin := addCreateItemAndShapeButton(stager, confForkJoin)
	callbacksForkJoin.button.Name = "Add Fork/Join"
	callbacksForkJoin.button.Icon = string(buttons.BUTTON_call_split)
	callbacksForkJoin.button.ToolTipText = "Add a Fork/Join to \"" + participant.Name + "\""
	callbacksForkJoin.OnBeforeCommit = func() {
		participant.IsTasksNodeExpanded = true
		if !slices.Contains(diagramProcess.ParticipantWhoseNodeIsExpanded, participant) {
			diagramProcess.ParticipantWhoseNodeIsExpanded = append(diagramProcess.ParticipantWhoseNodeIsExpanded, participant)
		}
		newForkJoin := callbacksForkJoin.createdItem
		newForkJoin.IsForkJoin = true
		if len(diagramProcess.Task_Shapes) > 0 {
			lastShape := diagramProcess.Task_Shapes[len(diagramProcess.Task_Shapes)-1]
			if lastShape.Task == newForkJoin {
				lastShape.SetWidth(ForkJoinDefaultLength)
				lastShape.SetHeight(ForkJoinBarThickness)
			}
		}
	}
}

func (stager *Stager) addCreateDecisionNodeButton(
	diagramProcess *DiagramProcess,
	participant *Participant,
	parentNode *tree.Node,
	isButtonInMenu bool,
) {
	confDecision := ItemAndShapeButtonConfiguration[
		Task, *Task,
		Participant, *Participant,
		TaskShape, *TaskShape,
	]{
		ItemButtonConfiguration: ItemButtonConfiguration[
			Task, *Task,
			Participant, *Participant,
		]{
			parentNode:                         parentNode,
			sliceForNewAddedItem:               &participant.Tasks,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &participant.IsTasksNodeExpanded,
			parentElement:                      participant,
			IsButtonInMenu:                     isButtonInMenu,
		},
		receivingDiagram:      diagramProcess,
		sliceForNewAddedShape: &diagramProcess.Task_Shapes,
	}
	callbacksDecision := addCreateItemAndShapeButton(stager, confDecision)
	callbacksDecision.button.Name = "Add Decision Node"
	callbacksDecision.button.Icon = string(buttons.BUTTON_diamond)
	callbacksDecision.button.ToolTipText = "Add a Decision Node to \"" + participant.Name + "\""
	callbacksDecision.OnBeforeCommit = func() {
		participant.IsTasksNodeExpanded = true
		if !slices.Contains(diagramProcess.ParticipantWhoseNodeIsExpanded, participant) {
			diagramProcess.ParticipantWhoseNodeIsExpanded = append(diagramProcess.ParticipantWhoseNodeIsExpanded, participant)
		}
		newDecision := callbacksDecision.createdItem
		newDecision.IsDecisionNode = true
		if len(diagramProcess.Task_Shapes) > 0 {
			lastShape := diagramProcess.Task_Shapes[len(diagramProcess.Task_Shapes)-1]
			if lastShape.Task == newDecision {
				lastShape.IsNameDisplayed = true
				lastShape.SetWidth(DecisionDiamondSize)
				lastShape.SetHeight(DecisionDiamondSize)
			}
		}
	}
}
