package models

import (
	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) treeDiagramCapture(library *Library, diagram *Diagram, libraryNode *tree.Node) {
	diagramNode := &tree.Node{
		Name:              diagram.Name,
		IsExpanded:        diagram.IsExpanded,
		IsNodeClickable:   true,
		HasCheckboxButton: true,
		IsChecked:         diagram.IsChecked,

		IsInEditMode: diagram.GetIsInRenameMode(),
	}
	libraryNode.Children = append(libraryNode.Children, diagramNode)

	diagramNode.Menu = &tree.Menu{
		Name: "Menu",
	}

	element := diagram
	node := diagramNode

	if !element.GetIsInRenameMode() {
		node.Menu.Buttons = append(node.Menu.Buttons,
			&tree.Button{
				Name: element.GetName() + " " + string(buttons.BUTTON_edit_note),
				Icon: string(buttons.BUTTON_edit_note),
				OnClick: func() {
					element.SetIsInRenameMode(true)
					stager.stage.Commit()
				},
				HasToolTip:      true,
				ToolTipText:     "Rename the " + GetGongstructNameFromPointer(element),
				ToolTipPosition: tree.Above,
			})
	} else {
		node.Menu.Buttons = append(node.Menu.Buttons,
			&tree.Button{
				Name: element.GetName() + " " + string(buttons.BUTTON_edit_off),
				Icon: string(buttons.BUTTON_edit_off),
				OnClick: func() {
					element.SetIsInRenameMode(false)
					stager.stage.Commit()
				},
				HasToolTip:      true,
				ToolTipText:     "Cancel renaming",
				ToolTipPosition: tree.Above,
			})
	}

	diagramNode.OnIsCheckedChanged = func(isChecked bool) {
		if isChecked {
			// reset all ddiagram selection
			for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
				diagram_.IsChecked = false
			}
			diagram.IsChecked = true
		} else {
			diagram.IsChecked = false
			// reset all ddiagram selection
			for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
				diagram_.IsChecked = false
			}
		}
		stager.stage.Commit()
	}
	diagramNode.OnIsExpandedChange = func(isExpanded bool) {
		diagram.IsExpanded = isExpanded
		stager.stage.Commit()
	}
	diagramNode.OnNameChange = func(newName string) {
		diagram.Name = newName
		diagram.SetIsInRenameMode(false)
		stager.stage.Commit()
	}
	diagramNode.OnClick = func(frontNode *tree.Node) {
		stager.probeForm.FillUpFormFromGongstruct(diagram, "Diagram")
		stager.stage.Commit()
	}

	{
		editButton := &tree.Button{
			Name:            "Diagram Editability",
			Icon:            string(buttons.BUTTON_edit),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,
			ToolTipText:     "Edit diagram",
			OnClick: func() {
				diagram.IsEditable_ = !diagram.IsEditable_
				stager.stage.Commit()
			},
		}
		if diagram.IsEditable_ {
			editButton.Icon = string(buttons.BUTTON_edit_off)
			editButton.ToolTipText = "Stop editing diagram"
		}
		diagramNode.Buttons = append(diagramNode.Buttons, editButton)
	}
	{
		copyButton := &tree.Button{
			Name:            "Diagram Copy",
			Icon:            string(buttons.BUTTON_copy_all),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,
			ToolTipText:     "Copy Diagram",
			OnClick:         onCopyDiagram(stager, diagram),
		}
		diagramNode.Menu.Buttons = append(diagramNode.Menu.Buttons, copyButton)
	}
	{
		showAllButton := &tree.Button{
			Name:            "Diagram Show All",
			Icon:            string(buttons.BUTTON_all_out),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,
			ToolTipText:     "Show All Elements in the Diagram",
			OnClick:         onShowAllInDiagram(stager, diagram),
		}
		diagramNode.Menu.Buttons = append(diagramNode.Menu.Buttons, showAllButton)
	}
	{
		showAllButton := &tree.Button{
			Name:            "Diagram Prefix",
			Icon:            string(buttons.BUTTON_show_chart),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,

			OnClick: func() {
				diagram.ShowPrefix = !diagram.ShowPrefix
				stager.stage.Commit()
			},
		}
		if !diagram.ShowPrefix {
			showAllButton.Icon = string(buttons.BUTTON_label)
			showAllButton.ToolTipText = "Show Prefix"
		} else {
			showAllButton.Icon = string(buttons.BUTTON_label_off)
			showAllButton.ToolTipText = "Hide Prefix"
		}
		diagramNode.Menu.Buttons = append(diagramNode.Menu.Buttons, showAllButton)
	}

	diagramsCategoryNode := &tree.Node{
		Name:            "Diagrams",
		FontStyle:       tree.ITALIC,
		IsExpanded:      diagram.IsDiagramsNodeExpanded,
		IsNodeClickable: true,
	}
	diagramNode.Children = append(diagramNode.Children, diagramsCategoryNode)
	diagramsCategoryNode.OnIsExpandedChange = stager.OnUpdateExpansion(&diagram.IsDiagramsNodeExpanded)

	if diagram.IsDiagramsNodeExpanded {
		stager.treeDiagramBSinDiagram(diagram, stager.GetRootLibrary(), diagramsCategoryNode)
	}

	wbsNode := &tree.Node{
		Name:            "Concerns",
		FontStyle:       tree.ITALIC,
		IsExpanded:      diagram.IsConcernsNodeExpanded,
		IsNodeClickable: true,
	}
	diagramNode.Children = append(diagramNode.Children, wbsNode)
	wbsNode.OnIsExpandedChange = stager.OnUpdateExpansion(&diagram.IsConcernsNodeExpanded)

	confRootConcerns := ItemAndShapeButtonConfiguration[
		Concern, *Concern,
		Library, *Library,
		ConcernShape, *ConcernShape,
	]{
		ItemButtonConfiguration: ItemButtonConfiguration[
			Concern, *Concern,
			Library, *Library,
		]{
			parentNode:                         wbsNode,
			sliceForNewAddedItem:               &library.RootConcerns,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagram.IsConcernsNodeExpanded,
		},
		receivingDiagram:      diagram,
		sliceForNewAddedShape: &diagram.Concern_Shapes,
	}
	addCreateItemAndShapeButton(stager, confRootConcerns)

	diagram.map_Concern_ConcernCompositionShape = make(map[*Concern]*ConcernCompositionShape)
	for _, shape := range diagram.ConcernComposition_Shapes {
		if shape.Concern != nil {
			diagram.map_Concern_ConcernCompositionShape[shape.Concern] = shape
		}
	}

	stakeholderNode := &tree.Node{
		Name:            "Stakeholders",
		FontStyle:       tree.ITALIC,
		IsExpanded:      diagram.IsStakeholdersNodeExpanded,
		IsNodeClickable: true,
	}
	diagramNode.Children = append(diagramNode.Children, stakeholderNode)
	stakeholderNode.OnIsExpandedChange = stager.OnUpdateExpansion(&diagram.IsStakeholdersNodeExpanded)

	confRootStakeholders := ItemAndShapeButtonConfiguration[
		Stakeholder, *Stakeholder,
		Library, *Library,
		StakeholderShape, *StakeholderShape,
	]{
		ItemButtonConfiguration: ItemButtonConfiguration[
			Stakeholder, *Stakeholder,
			Library, *Library,
		]{
			parentNode:                         stakeholderNode,
			sliceForNewAddedItem:               &library.RootStakeholders,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagram.IsStakeholdersNodeExpanded,
		},
		receivingDiagram:      diagram,
		sliceForNewAddedShape: &diagram.Stakeholder_Shapes,
	}
	addCreateItemAndShapeButton(stager, confRootStakeholders)

	for _, x := range library.RootStakeholders {
		stager.treeStakeholderBSinDiagram(diagram, x, stakeholderNode)
	}

	for _, x := range library.RootConcerns {
		stager.treeConcernBSinDiagram(diagram, x, wbsNode)
	}

	delivarablesNode := &tree.Node{
		Name:            "Deliverables",
		FontStyle:       tree.ITALIC,
		IsExpanded:      diagram.IsPBSNodeExpanded,
		IsNodeClickable: true,
	}
	diagramNode.Children = append(diagramNode.Children, delivarablesNode)
	delivarablesNode.OnIsExpandedChange = stager.OnUpdateExpansion(&diagram.IsPBSNodeExpanded)

	confRootDeliverables := ItemAndShapeButtonConfiguration[
		Deliverable, *Deliverable,
		Library, *Library,
		DeliverableShape, *DeliverableShape,
	]{
		ItemButtonConfiguration: ItemButtonConfiguration[
			Deliverable, *Deliverable,
			Library, *Library,
		]{
			parentNode:                         delivarablesNode,
			sliceForNewAddedItem:               &library.RootDeliverables,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagram.IsPBSNodeExpanded,
		},
		receivingDiagram:      diagram,
		sliceForNewAddedShape: &diagram.Deliverable_Shapes,
	}
	addCreateItemAndShapeButton(stager, confRootDeliverables)

	for _, x := range library.RootDeliverables {
		stager.treeDeliverableRecusriveInDiagram(diagram, x, delivarablesNode)
	}

	requirementsNode := &tree.Node{
		Name:            "Requirements",
		FontStyle:       tree.ITALIC,
		IsExpanded:      diagram.IsRequirementsNodeExpanded,
		IsNodeClickable: true,
	}
	diagramNode.Children = append(diagramNode.Children, requirementsNode)
	requirementsNode.OnIsExpandedChange = stager.OnUpdateExpansion(&diagram.IsRequirementsNodeExpanded)

	confRootRequirements := ItemAndShapeButtonConfiguration[
		Requirement, *Requirement,
		Library, *Library,
		RequirementShape, *RequirementShape,
	]{
		ItemButtonConfiguration: ItemButtonConfiguration[
			Requirement, *Requirement,
			Library, *Library,
		]{
			parentNode:                         requirementsNode,
			sliceForNewAddedItem:               &library.RootRequirements,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagram.IsRequirementsNodeExpanded,
		},
		receivingDiagram:      diagram,
		sliceForNewAddedShape: &diagram.Requirement_Shapes,
	}
	addCreateItemAndShapeButton(stager, confRootRequirements)

	for _, req := range library.RootRequirements {
		stager.treeRequirementBSinDiagram(diagram, req, requirementsNode)
	}

	conceptsNode := &tree.Node{
		Name:            "Concepts",
		FontStyle:       tree.ITALIC,
		IsExpanded:      diagram.IsConceptsNodeExpanded,
		IsNodeClickable: true,
	}
	diagramNode.Children = append(diagramNode.Children, conceptsNode)
	conceptsNode.OnIsExpandedChange = stager.OnUpdateExpansion(&diagram.IsConceptsNodeExpanded)

	confRootConcepts := ItemAndShapeButtonConfiguration[
		Concept, *Concept,
		Library, *Library,
		ConceptShape, *ConceptShape,
	]{
		ItemButtonConfiguration: ItemButtonConfiguration[
			Concept, *Concept,
			Library, *Library,
		]{
			parentNode:                         conceptsNode,
			sliceForNewAddedItem:               &library.RootConcepts,
			isParentNodeExpandedByAddOperation: true,
			parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
			parentNodeExpansionBooleanValue:    &diagram.IsConceptsNodeExpanded,
		},
		receivingDiagram:      diagram,
		sliceForNewAddedShape: &diagram.Concept_Shapes,
	}
	addCreateItemAndShapeButton(stager, confRootConcepts)

	for _, concept := range library.RootConcepts {
		stager.treeConceptBSinDiagram(diagram, concept, conceptsNode)
	}

	{
		notesNode := &tree.Node{
			Name:            "Notes",
			FontStyle:       tree.ITALIC,
			IsExpanded:      diagram.IsNotesNodeExpanded,
			IsNodeClickable: true,
		}
		diagramNode.Children = append(diagramNode.Children, notesNode)
		notesNode.OnIsExpandedChange = stager.OnUpdateExpansion(&diagram.IsNotesNodeExpanded)

		confNotes := ItemAndShapeButtonConfiguration[
			Note, *Note,
			Library, *Library,
			NoteShape, *NoteShape,
		]{
			ItemButtonConfiguration: ItemButtonConfiguration[
				Note, *Note,
				Library, *Library,
			]{
				parentNode:                         notesNode,
				sliceForNewAddedItem:               &library.Notes,
				isParentNodeExpandedByAddOperation: true,
				parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
				parentNodeExpansionBooleanValue:    &diagram.IsNotesNodeExpanded,
			},
			receivingDiagram:      diagram,
			sliceForNewAddedShape: &diagram.Note_Shapes,
		}
		addCreateItemAndShapeButton(stager, confNotes)

		for _, note := range library.Notes {
			confNode := TreeNodeAndShapeConfigurationWithoutLink[
				*Note, Note,
				*Note, Note, // Parent is not used
				*NoteShape, NoteShape,
				*Diagram,
			]{
				diagram:                     diagram,
				parentNode:                  notesNode,
				element:                     note,
				parentElement:               nil,
				elementsWhoseNodeIsExpanded: &diagram.NotesWhoseNodeIsExpanded,
				shapes:                      &diagram.Note_Shapes,
				shapesMap:                   diagram.map_Note_NoteShape,
			}
			noteNode := addNodeToTreeWithoutLink(stager, confNode)

			// allow display of associations note to deliverables
			for _, deliverable := range note.Deliverables {
				nodeDeliverable := &tree.Node{
					Name:            deliverable.Name,
					IsNodeClickable: true,
				}
				noteNode.Children = append(noteNode.Children, nodeDeliverable)

				showHideRelationButton := &tree.Button{
					Name: GetGongstructNameFromPointer(deliverable) + "- showHideRelationButton" + note.Name + " - " + deliverable.Name,

					HasToolTip:      true,
					ToolTipPosition: tree.Right,
				}
				nodeDeliverable.Buttons = append(nodeDeliverable.Buttons, showHideRelationButton)

				if _, ok := diagram.map_Deliverable_DeliverableShape[deliverable]; ok {
					if _, ok := diagram.map_Note_NoteShape[note]; ok {

						noteDeliverableShape, ok := diagram.map_Note_NoteDeliverableShape[noteDeliverableKey{Note: note, Deliverable: deliverable}]
						nodeDeliverable.IsChecked = ok

						if ok {
							showHideRelationButton.Icon = string(buttons.BUTTON_visibility_off)
							showHideRelationButton.ToolTipText = "Hide link from note \"" + note.Name +
								"\" to deliverable \"" + deliverable.Name + "\""
							// what to do when the deliverable node is clicked
							showHideRelationButton.OnClick = func() { onRemoveAssociationShape(stager, noteDeliverableShape, &diagram.NoteDeliverableShapes)() }
						} else {
							showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
							showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
								"\" to deliverable \"" + deliverable.Name + "\""
							showHideRelationButton.OnClick = func() { onAddAssociationShape(stager, note, deliverable, &diagram.NoteDeliverableShapes)() }
						}
					}
				}
			}

			for _, task := range note.Tasks {
				nodeTask := &tree.Node{
					Name:            task.Name,
					IsNodeClickable: true,
				}
				noteNode.Children = append(noteNode.Children, nodeTask)
				showHideRelationButton := &tree.Button{
					Name:            GetGongstructNameFromPointer(task),
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
				}
				nodeTask.Buttons = append(nodeTask.Buttons, showHideRelationButton)
				if _, ok := diagram.map_Concern_ConcernShape[task]; ok {
					if _, ok := diagram.map_Note_NoteShape[note]; ok {
						noteTaskShape, ok := diagram.map_Note_NoteTaskShape[noteTaskKey{Note: note, Task: task}]
						nodeTask.IsChecked = ok

						if ok {
							showHideRelationButton.Icon = string(buttons.BUTTON_visibility_off)
							showHideRelationButton.ToolTipText = "Hide link from note \"" + note.Name +
								"\" to task \"" + task.Name + "\""
							// what to do when the deliverable node is clicked
							showHideRelationButton.OnClick = func() { onRemoveAssociationShape(stager, noteTaskShape, &diagram.NoteTaskShapes)() }
						} else {
							showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
							showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
								"\" to task \"" + task.Name + "\""
							showHideRelationButton.OnClick = func() { onAddAssociationShape(stager, note, task, &diagram.NoteTaskShapes)() }
						}
					}
				}
			}

			for _, resource := range note.Resources {
				nodeResource := &tree.Node{
					Name:            resource.Name,
					IsNodeClickable: true,
				}
				noteNode.Children = append(noteNode.Children, nodeResource)
				showHideRelationButton := &tree.Button{
					Name:            GetGongstructNameFromPointer(resource),
					HasToolTip:      true,
					ToolTipPosition: tree.Right,
				}
				nodeResource.Buttons = append(nodeResource.Buttons, showHideRelationButton)
				if _, ok := diagram.map_Stakeholder_StakeholderShape[resource]; ok {
					if _, ok := diagram.map_Note_NoteShape[note]; ok {
						noteResourceShape, ok := diagram.map_Note_NoteResourceShape[noteResourceKey{Note: note, Stakeholder: resource}]
						nodeResource.IsChecked = ok
						if ok {
							showHideRelationButton.Icon = string(buttons.BUTTON_visibility_off)
							showHideRelationButton.ToolTipText = "Hide link from note \"" + note.Name +
								"\" to resource \"" + resource.Name + "\""
							// what to do when the deliverable node is clicked
							showHideRelationButton.OnClick = func() { onRemoveAssociationShape(stager, noteResourceShape, &diagram.NoteResourceShapes)() }
						} else {
							showHideRelationButton.Icon = string(buttons.BUTTON_visibility)
							showHideRelationButton.ToolTipText = "Show link from note \"" + note.Name +
								"\" to resource \"" + resource.Name + "\""
							showHideRelationButton.OnClick = func() { onAddAssociationShape(stager, note, resource, &diagram.NoteResourceShapes)() }
						}
					}
				}
			}
		}
	}
}
