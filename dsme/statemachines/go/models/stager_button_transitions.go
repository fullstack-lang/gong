package models

import (
	button "github.com/fullstack-lang/gong/lib/button/go/models"
)

func (stager *Stager) updateButtonsStage() {

	stager.buttonTransitionsStage.Reset()

	layout := new(button.Layout).Stage(stager.buttonTransitionsStage)

	percentage := 100.0 / float64(len(stager.architecture.Roles))

	map_Role_buttonGroup := make(map[*Role]*button.Group)
	for _, role := range stager.architecture.Roles {
		group := new(button.Group).Stage(stager.buttonTransitionsStage)
		group.Name = role.Name
		group.NbColumns = 3
		group.Percentage = percentage
		layout.Groups = append(layout.Groups, group)

		map_Role_buttonGroup[role] = group
	}

	// get the selected object
	var selectedObject *Object
	{
		objects := *GetGongstructInstancesSet[Object](stager.stage)
		for object := range objects {
			if object.IsSelected {
				selectedObject = object
			}
		}
	}
	if selectedObject == nil {
		stager.buttonTransitionsStage.Commit()
		return
	}

	var selectedDiagram *Diagram
	{
		diagramSet := *GetGongstructInstancesSet[Diagram](stager.stage)
		for diagram_ := range diagramSet {
			if diagram_.IsChecked {
				selectedDiagram = diagram_
			}
		}
	}
	if selectedDiagram == nil {
		stager.buttonTransitionsStage.Commit()
		return
	}

	for _, transitionShape := range selectedDiagram.Transition_Shapes {
		transition := transitionShape.Transition

		if selectedObject.State != transition.Start {

			var objectIsInASubState bool
			if transition.Start.IsComposite() {
				for _, subState := range transition.Start.SubStates {
					if subState == selectedObject.State {
						objectIsInASubState = true
						continue
					}
				}
			} else {
				continue
			}

			if !objectIsInASubState {
				continue
			}
		}

		for _, role := range transition.RolesWithPermissions {
			button := NewButtonTranstion(stager,
				transition.Name,
				"",
				transition.Name,
				transition,
				stager.stage,
			)

			group, ok := map_Role_buttonGroup[role]
			if !ok {
				continue
			}
			group.Buttons = append(group.Buttons, button)

			for _, role_ := range role.RolesWithSamePermissions {
				group_, ok := map_Role_buttonGroup[role_]
				if !ok {
					continue
				}
				group_.Buttons = append(group_.Buttons, button)
			}
		}
	}

	stager.buttonTransitionsStage.Commit()
}
