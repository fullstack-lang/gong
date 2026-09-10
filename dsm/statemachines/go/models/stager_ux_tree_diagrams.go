package models

import (
	"log"
	"slices"

	"github.com/fullstack-lang/gong/lib/tree/go/buttons"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
)

func (stager *Stager) ux_tree() {
	stager.treeStage.Reset()

	treeInstance := &tree.Tree{Name: string(DiagramTreeName)}

	stager.probeForm.AddCommitNavigationNode(func(gni GongNodeIF) {
		treeInstance.RootNodes = append(treeInstance.RootNodes, gni.(*tree.Node))
	})

	stager.treeLibrary(stager.GetRootLibrary(), &treeInstance.RootNodes)

	tree.StageBranch(stager.treeStage, treeInstance)

	stager.treeStage.Commit()
}

func (stager *Stager) treeLibrary(library *Library, parentNodes *[]*tree.Node) {
	libraryNode := &tree.Node{
		Name:                 library.Name,
		IsExpanded:           library.IsExpandedTmp,
		IsNodeClickable:      true,
		IsInEditMode:         library.isInRenameMode,
		IsWithPreceedingIcon: true,
		PreceedingIcon:       string(buttons.BUTTON_local_library),
	}
	*parentNodes = append(*parentNodes, libraryNode)

	if libraryNode.Menu == nil {
		libraryNode.Menu = &tree.Menu{Name: "Menu"}
	}
	stager.addNodeRenameButton(libraryNode, "library", library.GetIsInRenameMode(), library.SetIsInRenameMode)
	libraryNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsExpandedTmp)
	libraryNode.OnNameChange = stager.onNameChange(library)
	libraryNode.OnClick = onNodeClicked(stager, library)

	exportSysMLButton := &tree.Button{
		Name:            "Export to SysML V2",
		Icon:            string(buttons.BUTTON_file_download),
		ToolTipText:     "Export to SysML V2 format",
		HasToolTip:      true,
		ToolTipPosition: tree.Above,
		OnClick: func() {
			stager.exportSysML(library)
		},
	}
	libraryNode.Buttons = append(libraryNode.Buttons, exportSysMLButton)
	if libraryNode.Menu != nil {
		libraryNode.Menu.Buttons = append(libraryNode.Menu.Buttons, exportSysMLButton)
	}

	//
	// SubLibraries
	//
	subLibrariesNode := &tree.Node{
		Name:                 "Sub Libraries",
		FontStyle:            tree.ITALIC,
		IsExpanded:           library.IsSubLibrariesNodeExpanded,
		IsNodeClickable:      true,
		IsWithPreceedingIcon: true,
		PreceedingIcon:       string(buttons.BUTTON_folder),
	}
	libraryNode.Children = append(libraryNode.Children, subLibrariesNode)
	subLibrariesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&library.IsSubLibrariesNodeExpanded)
	subLibrariesNode.OnClick = onNodeClicked(stager, library)

	for _, subLibrary := range library.SubLibraries {
		stager.treeLibrary(subLibrary, &subLibrariesNode.Children)
	}

	// add sub library button
	confSubLibraries := ItemButtonConfiguration[
		Library, *Library,
		Library, *Library,
	]{
		parentNode:                         subLibrariesNode,
		sliceForNewAddedItem:               &library.SubLibraries,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsSubLibrariesNodeExpanded,
	}
	addCreateItemButton(stager, confSubLibraries)

	// add a statemachine to the library button
	confRootStateMachines := ItemButtonConfiguration[
		StateMachine, *StateMachine,
		Library, *Library,
	]{
		parentNode:                         libraryNode,
		sliceForNewAddedItem:               &library.RootStateMachines,
		isParentNodeExpandedByAddOperation: true,
		parentNodeExpansionType:            parentNodeExpansionTypeByBooleanValue,
		parentNodeExpansionBooleanValue:    &library.IsExpandedTmp,
	}
	addCreateItemButton(stager, confRootStateMachines)
	if libraryNode.Menu != nil && len(libraryNode.Buttons) > 0 {
		libraryNode.Menu.Buttons = append(libraryNode.Menu.Buttons, libraryNode.Buttons[0])
	}

	for _, stateMachine := range library.RootStateMachines {
		stager.treeStateMachines(stateMachine, libraryNode, &library.StateMachinesWhoseNodeIsExpanded)
	}

}

func (stager *Stager) treeStateMachines(
	stateMachine *StateMachine,
	parentNode *tree.Node,
	stateMachinesWhoseNodeIsExpanded *[]*StateMachine,
) {
	stateMachineNode := &tree.Node{
		Name:                 stateMachine.GetName(),
		IsExpanded:           slices.Contains(*stateMachinesWhoseNodeIsExpanded, stateMachine),
		IsNodeClickable:      true,
		IsInEditMode:         stateMachine.GetIsInRenameMode(),
		IsWithPreceedingIcon: true,
		PreceedingIcon:       string(buttons.BUTTON_account_tree),
	}
	parentNode.Children = append(parentNode.Children, stateMachineNode)

	if stateMachineNode.Menu == nil {
		stateMachineNode.Menu = &tree.Menu{Name: "Menu"}
	}
	stager.addNodeRenameButton(stateMachineNode, "state machine", stateMachine.GetIsInRenameMode(), stateMachine.SetIsInRenameMode)
	stateMachineNode.OnNameChange = stager.onNameChange(stateMachine)
	stateMachineNode.OnIsExpandedChange = onIsExpandedChangeSlice(stager, stateMachine, stateMachinesWhoseNodeIsExpanded)
	stateMachineNode.OnClick = onNodeClicked(stager, stateMachine)

	{
		addButton := &tree.Button{
			Name:            "Diagram" + " " + string(buttons.BUTTON_add),
			Icon:            string(buttons.BUTTON_add),
			HasToolTip:      true,
			ToolTipPosition: tree.Above,
			ToolTipText:     "Add a Diagram to the state machine",
			OnClick: func() {
				s := stager.stage
				newDiagram := (&Diagram{
					Name:        "New Diagram",
					IsEditable_: true,
				}).Stage(s)

				stateMachine.Diagrams = append(stateMachine.Diagrams, newDiagram)
				stager.stage.Commit()
			},
		}
		if stateMachineNode.Menu == nil {
			stateMachineNode.Menu = &tree.Menu{Name: "Menu"}
		}
		stateMachineNode.Menu.Buttons = append(stateMachineNode.Menu.Buttons, addButton)
	}

	transitionsSet := *GetGongstructInstancesSet[Transition](stager.stage)
	transitionSlice := SortGongstructSetByName(transitionsSet)

	for _, diagram := range stateMachine.Diagrams {
		diagramNode := new(tree.Node)
		stateMachineNode.Children = append(stateMachineNode.Children, diagramNode)
		diagramNode.Name = diagram.Name
		diagramNode.IsChecked = diagram.IsChecked
		diagramNode.IsExpanded = diagram.IsExpanded
		diagramNode.IsWithPreceedingIcon = true
		diagramNode.PreceedingIcon = string(buttons.BUTTON_schema)

		diagramNode.IsInEditMode = diagram.isInRenameMode

		diagramNode.IsNodeClickable = true
		diagramProxy := new(Diagram_Tree_DiagramProxy)
		diagramProxy.stager = stager
		diagramProxy.diagram = diagram

		diagramNode.HasCheckboxButton = true
		diagramNode.Impl = diagramProxy

		if diagramNode.Menu == nil {
			diagramNode.Menu = &tree.Menu{Name: "Menu"}
		}
		stager.addNodeRenameButton(diagramNode, "diagram", diagram.isInRenameMode, func(v bool) { diagram.isInRenameMode = v })
		diagramNode.OnNameChange = func(newName string) {
			diagram.Name = newName
			diagram.isInRenameMode = false
			stager.stage.Commit()
		}

		{
			copyButton := &tree.Button{
				Name:            "Diagram Copy" + " " + string(buttons.BUTTON_copy_all),
				Icon:            string(buttons.BUTTON_copy_all),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,
				ToolTipText:     "Copy Diagram",
				OnClick: func() {
					proxy := &DiagramCopyButtonProxy{
						stager:  stager,
						diagram: diagram,
					}
					proxy.ButtonUpdated(nil, nil, nil)
				},
			}
			diagramNode.Menu.Buttons = append(diagramNode.Menu.Buttons, copyButton)
		}
		// for displaying wether the State node is checked
		map_State__StateShape := make(map[*State]*StateShape)
		for _, stateShape := range diagram.State_Shapes {
			map_State__StateShape[stateShape.State] = stateShape
		}

		// for displaying wether the State_Transition node is checked
		map_Transition__TransitionShape := make(map[*Transition]*Transition_Shape)
		for _, transitionShape := range diagram.Transition_Shapes {
			map_Transition__TransitionShape[transitionShape.Transition] = transitionShape
		}

		map_Note__NoteShape := make(map[*Note]*NoteShape)
		for _, noteShape := range diagram.Note_Shapes {
			if noteShape.Note != nil {
				map_Note__NoteShape[noteShape.Note] = noteShape
			}
		}

		statesNode := &tree.Node{
			Name:                 "States",
			FontStyle:            tree.ITALIC,
			IsExpanded:           diagram.IsStatesNodeExpanded,
			IsNodeClickable:      true,
			IsWithPreceedingIcon: true,
			PreceedingIcon:       string(buttons.BUTTON_folder),
		}
		diagramNode.Children = append(diagramNode.Children, statesNode)
		statesNode.OnIsExpandedChange = stager.onIsExpandedChangeBool(&diagram.IsStatesNodeExpanded)
		statesNode.OnClick = func(*tree.Node) {
			stager.probeForm.FillUpFormFromGongstruct(diagram, "Diagram")
		}

		if diagram.IsChecked {
			var addButton *tree.Button
			addButton = &tree.Button{
				Name:            "Diagram" + " " + string(buttons.BUTTON_add),
				Icon:            string(buttons.BUTTON_add),
				HasToolTip:      true,
				ToolTipPosition: tree.Above,
				ToolTipText:     "Add a State to the State Machine and add it to the diagram",
			}
			addButton.OnClick = func() {
				s := stager.stage
				newState := new(State).Stage(s)

				newState.Name = "New State"
				stateMachine.States = append(stateMachine.States, newState)

				newStateShapeToDiagram(newState, diagram, addButton.ClientOnY).Stage(stager.stage)

				stager.stage.ComputeReverseMaps()
				if stager.probeForm != nil {
					stager.probeForm.FillUpFormFromGongstruct(newState, "State")
				}

				stager.stage.Commit()
			}
			statesNode.Buttons = append(statesNode.Buttons, addButton)
		}

		for _, state := range stateMachine.States {
			diagramStateNode := new(tree.Node)
			diagramStateNode.Name = state.Name
			diagramStateNode.HasCheckboxButton = true
			diagramStateNode.IsNodeClickable = true
			diagramStateNode.IsInEditMode = state.isInRenameMode
			diagramStateNode.IsExpanded = slices.Contains(diagram.StatesWhoseNodeIsExpanded, state)
			diagramStateNode.IsWithPreceedingIcon = true
			diagramStateNode.PreceedingIcon = string(buttons.BUTTON_crop_square)

			stager.addNodeRenameButton(diagramStateNode, "state", state.isInRenameMode, func(v bool) { state.isInRenameMode = v })

			var stateShape *StateShape
			var ok bool
			if stateShape, ok = map_State__StateShape[state]; ok {
				diagramStateNode.IsChecked = true

				howHideButton := &tree.Button{
					Name:            "Show State Shape",
					Icon:            string(buttons.BUTTON_visibility),
					HasToolTip:      true,
					ToolTipPosition: tree.Above,
					ToolTipText:     "Show State Shape",
					OnClick: func() {
						stateShape.SetIsHidden(!stateShape.GetIsHidden())
						stager.stage.Commit()
					},
				}
				diagramStateNode.Buttons = append(diagramStateNode.Buttons, howHideButton)

				if !stateShape.GetIsHidden() {
					howHideButton.Icon = string(buttons.BUTTON_visibility_off)
					howHideButton.ToolTipText = "Hide State Shape"
				}

				// range over transitions that have the state as a source or target
				// if the target is present, enable the check button
				for _, transition_ := range transitionSlice {
					if transition_.Start == state && transition_.End != nil {
						transitionNode := new(tree.Node)
						transitionNode.Name = transition_.Name + " --> " + transition_.End.Name

						if transition_.isInRenameMode {
							transitionNode.Name = transition_.Name
						}

						transitionNode.HasCheckboxButton = true
						transitionNode.IsNodeClickable = true
						transitionNode.IsInEditMode = transition_.isInRenameMode
						transitionNode.IsWithPreceedingIcon = true
						transitionNode.PreceedingIcon = string(buttons.BUTTON_arrow_forward)

						stager.addNodeRenameButton(transitionNode, "transition", transition_.isInRenameMode, func(v bool) { transition_.isInRenameMode = v })
						transitionNode.OnNameChange = func(newName string) {
							transition_.Name = newName
							transition_.isInRenameMode = false
							stager.stage.Commit()
						}

						proxy := new(DiagramTree_Transition_Proxy)
						proxy.diagram = diagram
						proxy.stager = stager
						proxy.transition = transition_
						transitionNode.Impl = proxy

						// set the transition of the proxy (can be nil, )
						var transitionShape *Transition_Shape
						if transitionShape_, ok := map_Transition__TransitionShape[transition_]; ok {
							transitionShape = transitionShape_
							transitionNode.IsChecked = true
						}

						proxy.transitionShape = transitionShape

						// disable if end is not present
						if _, ok := map_State__StateShape[transition_.End]; !ok {
							transitionNode.IsCheckboxDisabled = true
						}

						showHideButton := &tree.Button{
							Name:            "Show Transition Shape",
							Icon:            string(buttons.BUTTON_visibility),
							HasToolTip:      true,
							ToolTipPosition: tree.Above,
							ToolTipText:     "Show Transition Shape",
							OnClick: func() {
								if transitionShape != nil {
									transitionShape.SetIsHidden(!transitionShape.GetIsHidden())
									stager.stage.Commit()
								}
							},
						}
						transitionNode.Buttons = append(transitionNode.Buttons, showHideButton)
						if transitionShape != nil && !transitionShape.GetIsHidden() {
							showHideButton.Icon = string(buttons.BUTTON_visibility_off)
							showHideButton.ToolTipText = "Hide Transition Shape"
						}

						diagramStateNode.Children = append(diagramStateNode.Children, transitionNode)
					}
				}

				addNoteButton := &tree.Button{
					Name:            state.GetName() + " " + string(buttons.BUTTON_note_add),
					Icon:            string(buttons.BUTTON_note_add),
					HasToolTip:      true,
					ToolTipPosition: tree.Above,
					ToolTipText:     "Add a Note to this State and to the diagram",
					OnClick: func() {
						s := stager.stage
						newNote := new(Note).Stage(s)
						newNote.Name = "New Note"
						newNote.State = state
						state.Notes = append(state.Notes, newNote)

						newShape := new(NoteShape)
						newShape.Note = newNote
						newShape.Name = newNote.GetName() + "-" + diagram.GetName()
						newShape.Height = 80
						newShape.Width = 200
						newShape.X = stateShape.X + 250
						newShape.Y = stateShape.Y
						diagram.Note_Shapes = append(diagram.Note_Shapes, newShape)
						newShape.Stage(s)

						noteStateShape := new(NoteStateShape).Stage(s)
						noteStateShape.Name = newNote.GetName() + "-" + state.GetName()
						noteStateShape.Note = newNote
						noteStateShape.State = state
						noteStateShape.StartOrientation = ORIENTATION_HORIZONTAL
						noteStateShape.EndOrientation = ORIENTATION_HORIZONTAL
						noteStateShape.CornerOffsetRatio = 1.2
						noteStateShape.StartRatio = 0.5
						noteStateShape.EndRatio = 0.5
						diagram.NoteState_Shapes = append(diagram.NoteState_Shapes, noteStateShape)

						stager.stage.ComputeReverseMaps()
						if stager.probeForm != nil {
							stager.probeForm.FillUpFormFromGongstruct(newNote, "Note")
						}

						stager.stage.Commit()
					},
				}
				diagramStateNode.Buttons = append(diagramStateNode.Buttons, addNoteButton)

				for _, note := range state.Notes {
					noteNode := new(tree.Node)
					noteNode.Name = note.Name
					noteNode.HasCheckboxButton = true
					noteNode.IsNodeClickable = true
					noteNode.IsInEditMode = note.isInRenameMode
					noteNode.IsWithPreceedingIcon = true
					noteNode.PreceedingIcon = string(buttons.BUTTON_description)

					var noteShape *NoteShape
					var isNoteChecked bool
					if shape, ok := map_Note__NoteShape[note]; ok {
						noteShape = shape
						isNoteChecked = true
						noteNode.IsChecked = true
					}

					stager.addNodeRenameButton(noteNode, "note", note.isInRenameMode, func(v bool) { note.isInRenameMode = v })

					if isNoteChecked && noteShape != nil {
						howHideButton := &tree.Button{
							Name:            "Show Note Shape",
							Icon:            string(buttons.BUTTON_visibility),
							HasToolTip:      true,
							ToolTipPosition: tree.Above,
							ToolTipText:     "Show Note Shape",
							OnClick: func() {
								noteShape.SetIsHidden(!noteShape.GetIsHidden())
								stager.stage.Commit()
							},
						}
						noteNode.Buttons = append(noteNode.Buttons, howHideButton)

						if !noteShape.GetIsHidden() {
							howHideButton.Icon = string(buttons.BUTTON_visibility_off)
							howHideButton.ToolTipText = "Hide Note Shape"
						}
					}

					noteNode.OnIsCheckedChanged = func(isChecked bool) {
						if isChecked {
							if noteShape != nil {
								log.Fatalln("adding a shape to an already note shape")
							}
							newShape := new(NoteShape)
							newShape.Note = note
							newShape.Name = note.GetName() + "-" + diagram.GetName()
							newShape.Height = 80
							newShape.Width = 200
							newShape.X = stateShape.X + 250
							newShape.Y = stateShape.Y
							diagram.Note_Shapes = append(diagram.Note_Shapes, newShape)
							newShape.Stage(stager.stage)

							hasNoteStateShape := false
							for _, nss := range diagram.NoteState_Shapes {
								if nss.Note == note && nss.State == state {
									hasNoteStateShape = true
									break
								}
							}
							if !hasNoteStateShape {
								nss := new(NoteStateShape).Stage(stager.stage)
								nss.Name = note.GetName() + "-" + state.GetName()
								nss.Note = note
								nss.State = state
								nss.StartOrientation = ORIENTATION_HORIZONTAL
								nss.EndOrientation = ORIENTATION_HORIZONTAL
								nss.CornerOffsetRatio = 1.2
								nss.StartRatio = 0.5
								nss.EndRatio = 0.5
								diagram.NoteState_Shapes = append(diagram.NoteState_Shapes, nss)
							}

							stager.stage.Commit()
						} else {
							if noteShape != nil {
								noteShape.Unstage(stager.stage)
								idx := slices.Index(diagram.Note_Shapes, noteShape)
								if idx != -1 {
									diagram.Note_Shapes = slices.Delete(diagram.Note_Shapes, idx, idx+1)
								}
							}
							var remainingNSS []*NoteStateShape
							for _, nss := range diagram.NoteState_Shapes {
								if nss.Note == note && nss.State == state {
									nss.Unstage(stager.stage)
								} else {
									remainingNSS = append(remainingNSS, nss)
								}
							}
							diagram.NoteState_Shapes = remainingNSS

							stager.stage.Commit()
						}
					}

					noteNode.OnNameChange = func(newName string) {
						note.Name = newName
						note.isInRenameMode = false
						stager.stage.Commit()
					}

					noteNode.OnClick = func(frontNode *tree.Node) {
						stager.probeForm.FillUpFormFromGongstruct(note, "Note")
					}

					diagramStateNode.Children = append(diagramStateNode.Children, noteNode)
				}
			}

			diagramStateNode.OnIsCheckedChanged = func(isChecked bool) {
				if isChecked {
					// add the State_Shape
					if stateShape != nil {
						log.Fatalln("adding a shape to an already state shape")
					}

					newStateShapeToDiagram(state, diagram, diagramStateNode.ClientOnY).Stage(stager.stage)
					stager.stage.Commit()
				} else {
					// one need to remove the State_Shape
					if stateShape == nil {
						log.Fatalln("remove a non existing shape to state")
					}

					stateShape.Unstage(stager.stage)
					idx := slices.Index(diagram.State_Shapes, stateShape)
					diagram.State_Shapes = slices.Delete(diagram.State_Shapes, idx, idx+1)

					stager.stage.Commit()
				}
			}

			diagramStateNode.OnNameChange = func(newName string) {
				state.Name = newName
				state.isInRenameMode = false
				stager.stage.Commit()
			}

			diagramStateNode.OnIsExpandedChange = func(isExpanded bool) {
				if isExpanded {
					if slices.Index(diagram.StatesWhoseNodeIsExpanded, state) == -1 {
						diagram.StatesWhoseNodeIsExpanded = append(diagram.StatesWhoseNodeIsExpanded, state)
					}
				} else {
					if idx := slices.Index(diagram.StatesWhoseNodeIsExpanded, state); idx != -1 {
						diagram.StatesWhoseNodeIsExpanded = slices.Delete(diagram.StatesWhoseNodeIsExpanded, idx, idx+1)
					}
				}
				stager.stage.Commit()
			}

			diagramStateNode.OnClick = func(frontNode *tree.Node) {
				stager.probeForm.FillUpFormFromGongstruct(state, "State")
			}

			statesNode.Children = append(statesNode.Children, diagramStateNode)
		}
	}
}

func (stager *Stager) addNodeRenameButton(
	node *tree.Node,
	typeName string,
	isInRenameMode bool,
	setIsInRenameMode func(bool),
) {
	if !isInRenameMode {
		renameButton := &tree.Button{
			Name:            "Rename " + typeName,
			Icon:            string(buttons.BUTTON_edit_note),
			HasToolTip:      true,
			ToolTipText:     "Rename the " + typeName,
			ToolTipPosition: tree.Above,
			OnClick: func() {
				setIsInRenameMode(true)
				stager.stage.Commit()
			},
		}
		node.Buttons = append(node.Buttons, renameButton)
		if node.Menu != nil {
			node.Menu.Buttons = append(node.Menu.Buttons, renameButton)
		}
	} else {
		cancelButton := &tree.Button{
			Name:            "Cancel rename " + typeName,
			Icon:            string(buttons.BUTTON_edit_off),
			HasToolTip:      true,
			ToolTipText:     "Cancel renaming",
			ToolTipPosition: tree.Above,
			OnClick: func() {
				setIsInRenameMode(false)
				stager.stage.Commit()
			},
		}
		node.Buttons = append(node.Buttons, cancelButton)
		if node.Menu != nil {
			node.Menu.Buttons = append(node.Menu.Buttons, cancelButton)
		}
	}
}

