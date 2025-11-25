package models

import (
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) updateSvgStage() {
	stager.svgStage.Reset()

	map_State_Rect := make(map[*State]*svg.Rect)

	svgStage := stager.svgStage
	svgObject := (&svg.SVG{Name: `SVG`})

	var diagram *Diagram
	{
		for diagram_ := range *GetGongstructInstancesSet[Diagram](stager.stage) {
			if diagram_.IsChecked {
				diagram = diagram_
			}
		}
	}

	if diagram == nil {
		stager.svgStage.Commit()
		return
	}

	svgImpl := &svgProxy{
		stager:                 stager,
		diagram:                diagram,
		map_SvgRect_StateShape: make(map[*svg.Rect]*StateShape),
	}
	svgObject.Impl = svgImpl

	svgObject.Name = diagram.Name
	svgObject.IsEditable = diagram.IsEditable()

	layer := (&svg.Layer{Name: "Layer 1"})
	svgObject.Layers = append(svgObject.Layers, layer)

	statesSet := make(map[*State]any)

	for _, stateShape_ := range diagram.State_Shapes {
		statesSet[stateShape_.State] = true
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

	for _, stateShape := range diagram.State_Shapes {

		if stateShape.State == nil {
			continue
		}
		state := stateShape.State

		var isSelected bool
		if selectedObject != nil {
			isSelected = selectedObject.State == state
		}

		rect := stager.svgGenerateRect(
			diagram,
			stateShape,
			isSelected,
			layer)
		map_State_Rect[state] = rect
		svgImpl.map_SvgRect_StateShape[rect] = stateShape
	}

	for _, transtionShape := range diagram.Transition_Shapes {
		transition := transtionShape.Transition

		if transition == nil || transition.Start == nil || transition.End == nil {
			continue
		}

		startRect := map_State_Rect[transition.Start]
		endRect := map_State_Rect[transition.End]

		linkName := transition.Name

		var rolesNames string
		{

			for _, role := range transition.RolesWithPermissions {

				rolesNames += role.Acronym

				for _, role_ := range role.RolesWithSamePermissions {
					rolesNames += role_.Acronym
				}
			}
		}
		if rolesNames != "" {
			linkName += " /" + rolesNames
		}

		var messageName string
		for idx, messageType := range transition.GeneratedMessages {
			messageName += messageType.Name

			l := len(transition.GeneratedMessages)
			if idx < l-1 {
				messageName += " + "
			}
		}
		stager.svgGenerateLink(
			startRect, endRect,
			linkName, messageName,
			&transtionShape.LinkShape, transition, layer, false)
	}

	svg.StageBranch(svgStage, svgObject)
	stager.svgStage.Commit()
}
