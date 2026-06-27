package models

import (
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) svg() {
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
	stager.diagram = diagram
	stager.svgObject = svgObject

	stager.map_SvgRect_StateShape = make(map[*svg.Rect]*StateShape)
	stager.map_SvgRect_NoteShape = make(map[*svg.Rect]*NoteShape)
	svgObject.OnUpdate = stager.onUpdateSVG

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
		if stateShape.GetIsHidden() {
			continue
		}

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
		stager.map_SvgRect_StateShape[rect] = stateShape
	}

	for _, transtionShape := range diagram.Transition_Shapes {
		if transtionShape.GetIsHidden() {
			continue
		}

		transition := transtionShape.Transition

		if transition == nil || transition.Start == nil || transition.End == nil {
			continue
		}

		startRect := map_State_Rect[transition.Start]
		endRect := map_State_Rect[transition.End]

		if startRect == nil || endRect == nil {
			continue
		}

		stager.svgGenerateLink(
			startRect, endRect,
			&transtionShape.LinkShape, transition, layer, false)
	}

	map_Note_Rect := make(map[*Note]*svg.Rect)

	for _, noteShape := range diagram.Note_Shapes {
		if noteShape.GetIsHidden() {
			continue
		}

		if noteShape.Note == nil {
			continue
		}
		note := noteShape.Note

		var isSelected bool
		if selectedObject != nil {
			isSelected = selectedObject.Name == note.Name // Simplification since Object doesn't have Note
		}

		rect := stager.svgGenerateNoteRect(
			diagram,
			noteShape,
			isSelected,
			layer)
		map_Note_Rect[note] = rect
		stager.map_SvgRect_NoteShape[rect] = noteShape
	}

	for _, noteStateShape := range diagram.NoteState_Shapes {
		if noteStateShape.Note == nil || noteStateShape.State == nil {
			continue
		}

		startRect := map_Note_Rect[noteStateShape.Note]
		endRect := map_State_Rect[noteStateShape.State]

		if startRect == nil || endRect == nil {
			continue
		}

		stager.svgGenerateNoteLink(
			startRect, endRect,
			noteStateShape,
			layer,
		)
	}

	for _, noteTransitionShape := range diagram.NoteTransition_Shapes {
		if noteTransitionShape.Note == nil || noteTransitionShape.Transition == nil {
			continue
		}

		startRect := map_Note_Rect[noteTransitionShape.Note]
		// Link to the transition's Start State since we can't link to a link
		endRect := map_State_Rect[noteTransitionShape.Transition.Start]

		if startRect == nil || endRect == nil {
			continue
		}

		stager.svgGenerateNoteLink(
			startRect, endRect,
			noteTransitionShape,
			layer,
		)
	}

	svg.StageBranch(svgStage, svgObject)
	stager.svgStage.Commit()
}
