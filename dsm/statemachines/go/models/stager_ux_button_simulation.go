package models

import (
	button "github.com/fullstack-lang/gong/lib/button/go/models"
	buttons "github.com/fullstack-lang/gong/lib/button/go/models"
)

func (stager *Stager) buttonSimulation() {

	stager.buttonTransitionsStage.Reset()

	layout := new(button.Layout).Stage(stager.buttonTransitionsStage)

	var roles []*Role
	if root := stager.getRootLibrary(); root != nil {
		roles = root.Roles
	}
	if len(roles) == 0 {
		roles = GetGongstrucsSorted[*Role](stager.stage)
	}

	if len(roles) == 0 {
		stager.buttonTransitionsStage.Commit()
		return
	}

	percentage := 100.0 / float64(len(roles))

	map_Role_buttonGroup := make(map[*Role]*button.Group)
	for _, role := range roles {
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
			if transition.Start != nil && transition.Start.IsComposite() {
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
			label := transition.Name
			if label == "" {
				label = "> " + transition.End.GetName()
			}

			button := (&buttons.Button{
				Name:  transition.Name,
				Icon:  "",
				Label: label,
				OnClick: func() {
					transition.performTransition(stager.stage)
				},
			}).Stage(stager.buttonTransitionsStage)

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
