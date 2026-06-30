package models

import (
	"fmt"
	"reflect"
	"strings"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"

	"github.com/fullstack-lang/gong/lib/strutils"
)

const HeightBetween2AttributeShapes = 20

type diagramInterface interface {
	IsEditable() bool
}

var smallRadius = 10.0
var bigRadius = 18.0

func (stager *Stager) svgGenerateRect(
	diagram diagramInterface, // might be ni
	stateShape *StateShape,
	isSelected bool,
	layer *svg.Layer,
) *svg.Rect {

	state := stateShape.State
	_, isStartState := stager.set_StartStates[state]

	rect := new(svg.Rect)
	rect.Name = state.Name
	rect.Stroke = svg.Black.ToString()
	rect.StrokeWidth = 2
	rect.StrokeOpacity = 1
	rect.X = stateShape.X
	rect.Y = stateShape.Y
	rect.Width = stateShape.Width
	rect.Height = stateShape.Height
	rect.RX = 15

	// rect is editable if diagram is not null
	if !reflect.ValueOf(diagram).IsNil() {
		if diagram.IsEditable() {
			rect.CanHaveBottomHandle = true
			rect.CanHaveLeftHandle = true
			rect.CanHaveRightHandle = true
			rect.CanHaveTopHandle = true

			rect.CanMoveHorizontaly = true
			rect.CanMoveVerticaly = true

			rect.OnSelect = func() {
				stager.stage.CommitWithSuspendedCallbacks()
				stager.probeForm.FillUpFormFromGongstruct(state, "State")
				stager.ux_tree()
			}
			rect.OnMove = func(x, y float64) {
				stateShape.X = x
				stateShape.Y = y
				// Issue #7, this will allow multiple rect to be moved together
				stager.stage.CommitWithSuspendedCallbacks()
				stager.ux_tree()
			}
			rect.OnResize = func(x, y, width, height float64) {
				stateShape.X = x
				stateShape.Y = y
				stateShape.Width = width
				stateShape.Height = height
				stager.stage.Commit()
			}
			// for allowing later Stage() on the rect shape
			stateShape.receiver = stateShape
		}
	}

	if isSelected {
		rect.Color = svg.Lightblue.ToString()
		rect.FillOpacity = 0.6
	} else {
		// rect.Color = svg.Lightgrey.ToString()
		// rect.FillOpacity = 0.3
	}

	stateTitleText := new(svg.RectAnchoredText)
	{
		stateTitleText.Name = state.Name
		stateTitleText.Content = state.Name

		if rect.Width > 0 {
			stateTitleText.Content = strutils.WrapString(stateShape.State.Name, int(rect.Width/stager.architecture.NbPixPerCharacter))
		}
		stateTitleText.Stroke = svg.Black.ToString()
		stateTitleText.StrokeWidth = 1
		stateTitleText.StrokeOpacity = 1
		stateTitleText.Color = svg.Black.ToString()
		stateTitleText.FillOpacity = 1

		stateTitleText.FontSize = "16px"
		stateTitleText.X_Offset = 0
		
		if state.Entry != nil || len(state.Activities) > 0 || state.Exit != nil {
			stateTitleText.Y_Offset = 15
			stateTitleText.RectAnchorType = svg.RECT_TOP
		} else {
			stateTitleText.Y_Offset = 0
			stateTitleText.RectAnchorType = svg.RECT_CENTER_MIDDLE
		}

		stateTitleText.TextAnchorType = svg.TEXT_ANCHOR_CENTER

		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, stateTitleText)
	}

	lineYOffset := stateTitleText.Y_Offset + float64(HeightBetween2AttributeShapes*(1+strings.Count(stateTitleText.Content, "\n")))
	
	if state.Entry != nil || len(state.Activities) > 0 || state.Exit != nil {
		line := new(svg.RectAnchoredPath)
		line.Name = "Separator"
		line.Definition = fmt.Sprintf("M 0 0 L %f 0", stateShape.Width)
		line.X_Offset = 0
		line.Y_Offset = lineYOffset - 5 // a little bit above the first element
		line.RectAnchorType = svg.RECT_TOP_LEFT
		line.Stroke = svg.Black.ToString()
		line.StrokeWidth = 1
		line.StrokeOpacity = 1
		rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, line)
	}

	currentY_Offset := lineYOffset + 10 // Start text 10px below the line

	x_offset := 10.0 // on the left of the state

	// Add the /Entry Action
	if action := state.Entry; action != nil {
		text := new(svg.RectAnchoredText)
		text.Name = action.Name
		content := "/entry " + action.Name

		if rect.Width > 0 {
			content = strutils.WrapString(content, int(rect.Width/stager.architecture.NbPixPerCharacter))
		}
		text.Content = content
		text.Stroke = svg.Black.ToString()
		text.StrokeWidth = 1
		text.StrokeOpacity = 1
		text.Color = svg.Black.ToString()
		text.FillOpacity = 1

		if action.Criticality == CriticalityCritical {
			text.Stroke = svg.Red.ToString()
			text.Color = svg.Red.ToString()
		}

		text.FontSize = "16px"
		text.X_Offset = x_offset
		text.Y_Offset = currentY_Offset
		text.RectAnchorType = svg.RECT_TOP_LEFT
		text.TextAnchorType = svg.TEXT_ANCHOR_START

		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, text)
		currentY_Offset += float64(HeightBetween2AttributeShapes * (1 + strings.Count(text.Content, "\n")))
	}

	// add the /Do activities

	for _, activity := range state.Activities {

		text := new(svg.RectAnchoredText)
		text.Name = activity.Name
		content := "/do " + activity.Name

		if rect.Width > 0 {
			content = strutils.WrapString(content, int(rect.Width/stager.architecture.NbPixPerCharacter))
		}
		text.Content = content
		text.Stroke = svg.Black.ToString()
		text.StrokeWidth = 1
		text.StrokeOpacity = 1
		text.Color = svg.Black.ToString()
		text.FillOpacity = 1

		if activity.Criticality == CriticalityCritical {
			text.Stroke = svg.Red.ToString()
			text.Color = svg.Red.ToString()
		}

		text.FontSize = "16px"
		text.X_Offset = x_offset
		text.Y_Offset = currentY_Offset
		text.RectAnchorType = svg.RECT_TOP_LEFT
		text.TextAnchorType = svg.TEXT_ANCHOR_START

		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, text)

		currentY_Offset += float64(HeightBetween2AttributeShapes * (1 + strings.Count(text.Content, "\n")))

	}

	// Add the /Entry Action
	if action := state.Exit; action != nil {
		text := new(svg.RectAnchoredText)
		text.Name = action.Name
		content := "/exit " + action.Name

		if rect.Width > 0 {
			content = strutils.WrapString(content, int(rect.Width/stager.architecture.NbPixPerCharacter))
		}
		text.Content = content
		text.Stroke = svg.Black.ToString()
		text.StrokeWidth = 1
		text.StrokeOpacity = 1
		text.Color = svg.Black.ToString()
		text.FillOpacity = 1

		if action.Criticality == CriticalityCritical {
			text.Stroke = svg.Red.ToString()
			text.Color = svg.Red.ToString()
		}

		text.FontSize = "16px"
		text.X_Offset = x_offset
		text.Y_Offset = currentY_Offset
		text.RectAnchorType = svg.RECT_TOP_LEFT
		text.TextAnchorType = svg.TEXT_ANCHOR_START

		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, text)
		currentY_Offset += float64(HeightBetween2AttributeShapes * (1 + strings.Count(text.Content, "\n")))
	}

	stager.addIconToState(rect, stateTitleText, state, isStartState, isSelected)

	if state.IsFictious {
		rect.StrokeDashArray = "5 5"
	}

	layer.Rects = append(layer.Rects, rect)
	return rect
}



func (stager *Stager) svgGenerateNoteRect(
	diagram diagramInterface,
	noteShape *NoteShape,
	isSelected bool,
	layer *svg.Layer,
) *svg.Rect {

	note := noteShape.Note

	rect := new(svg.Rect)
	rect.Name = note.Name
	rect.Stroke = "#FBC02D"
	rect.StrokeWidth = 1.0
	rect.StrokeOpacity = 1
	rect.Color = "#FFF9C4"
	rect.FillOpacity = 1
	rect.X = noteShape.X
	rect.Y = noteShape.Y
	rect.Width = noteShape.Width
	rect.Height = noteShape.Height
	rect.RX = 0

	if !reflect.ValueOf(diagram).IsNil() {
		if diagram.IsEditable() {
			rect.CanHaveBottomHandle = true
			rect.CanHaveLeftHandle = true
			rect.CanHaveRightHandle = true
			rect.CanHaveTopHandle = true

			rect.CanMoveHorizontaly = true
			rect.CanMoveVerticaly = true

			rect.OnSelect = func() {
				stager.stage.CommitWithSuspendedCallbacks()
				stager.probeForm.FillUpFormFromGongstruct(note, "Note")
				stager.ux_tree()
			}
			rect.OnMove = func(x, y float64) {
				noteShape.X = x
				noteShape.Y = y
				stager.stage.CommitWithSuspendedCallbacks()
				stager.ux_tree()
			}
			rect.OnResize = func(x, y, width, height float64) {
				noteShape.X = x
				noteShape.Y = y
				noteShape.Width = width
				noteShape.Height = height
				stager.stage.Commit()
			}
			noteShape.receiver = noteShape
		}
	}

	if isSelected {
		rect.Stroke = svg.Red.ToString()
		rect.StrokeWidth = 4
	}

	noteTitleText := new(svg.RectAnchoredText)
	noteTitleText.Name = note.Name
	content := "📝 " + note.Name

	margin := 20.0
	wrapWidth := rect.Width - margin
	if wrapWidth > 0 {
		content = strutils.WrapStringPreservingNewlines(content, int(wrapWidth/stager.architecture.NbPixPerCharacter))
	}

	noteTitleText.Content = content
	noteTitleText.Stroke = svg.Black.ToString()
	noteTitleText.StrokeWidth = 1
	noteTitleText.StrokeOpacity = 1
	noteTitleText.Color = svg.Black.ToString()
	noteTitleText.FillOpacity = 1
	noteTitleText.FontSize = "16px"
	noteTitleText.FontWeight = "normal"
	noteTitleText.FontStyle = "italic"
	noteTitleText.RectAnchorType = svg.RECT_TOP_LEFT
	noteTitleText.TextAnchorType = svg.TEXT_ANCHOR_START
	noteTitleText.X_Offset = 10
	noteTitleText.Y_Offset = 20
	rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, noteTitleText)

	layer.Rects = append(layer.Rects, rect)

	return rect
}



func (stager *Stager) addIconToState(
	rect *svg.Rect,
	stateTitleText *svg.RectAnchoredText,
	state *State,
	isStartState bool,
	isSelected bool,
) {
	if !state.IsDecisionNode && !isStartState && !state.IsEndState {
		return
	}

	stateTitleText.TextAnchorType = svg.TEXT_ANCHOR_START
	stateTitleText.RectAnchorType = svg.RECT_TOP_LEFT
	stateTitleText.DominantBaseline = svg.DominantBaselineCentral
	stateTitleText.WhiteSpace = svg.WhiteSpaceEnumPre
	stateTitleText.X_Offset = 0
	stateTitleText.Y_Offset = 0

	rect.CanHaveBottomHandle = false
	rect.CanHaveTopHandle = false

	if state.IsDecisionNode {
		diamond := new(svg.RectAnchoredPath)
		diamond.Stroke = svg.Black.ToString()
		diamond.StrokeWidth = 2
		diamond.StrokeOpacity = 1
		diamond.ScalePropotionnally = true
		diamond.AppliedScaling = 1

		diamond.Definition = "M 25 0 L 50 25 L 25 50 L 0 25 Z"
		diamond.X_Offset = -50
		diamond.Y_Offset = -25
		stateTitleText.X_Offset = -25
		diamond.RectAnchorType = svg.RECT_RIGHT
		rect.Height = 50
		rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, diamond)
	}

	if isStartState {
		if isSelected {
			circle := new(svg.RectAnchoredPath)

			circle.Color = svg.Lightblue.ToString()
			circle.FillOpacity = 0.6

			circle.Definition = fmt.Sprintf("M %f 0 A %f %f 0 0 1 %f %f A %f %f 0 0 1 %f 0 Z",
				bigRadius, bigRadius, bigRadius, bigRadius, 2*bigRadius, bigRadius, bigRadius, bigRadius)
			circle.X_Offset = -bigRadius
			circle.Y_Offset = -bigRadius
			circle.RectAnchorType = svg.RECT_RIGHT
			rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, circle)
		}

		circle := new(svg.RectAnchoredPath)
		circle.Stroke = svg.Black.ToString()
		circle.StrokeWidth = 2
		circle.StrokeOpacity = 1

		circle.Color = svg.Black.ToString()
		circle.FillOpacity = 1.0

		// we allow resizing for the sake of the text width
		if rect.Width < 2*smallRadius {
			rect.Width = 2 * smallRadius
		}
		rect.Height = 2 * smallRadius

		circle.Definition = fmt.Sprintf("M %f 0 A %f %f 0 0 1 %f %f A %f %f 0 0 1 %f 0 Z",
			smallRadius, smallRadius, smallRadius, smallRadius, 2*smallRadius, smallRadius, smallRadius, smallRadius)
		circle.X_Offset = -smallRadius
		circle.Y_Offset = -smallRadius
		circle.RectAnchorType = svg.RECT_RIGHT
		rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, circle)
	}

	if state.IsEndState {
		if rect.Width < 2*bigRadius {
			rect.Width = 2 * bigRadius
		}
		rect.Height = 2 * bigRadius

		{
			circle := new(svg.RectAnchoredPath)

			circle.Stroke = svg.Black.ToString()
			circle.StrokeWidth = 1
			circle.StrokeOpacity = 1.0

			circle.Definition = fmt.Sprintf("M %f 0 A %f %f 0 0 1 %f %f A %f %f 0 0 1 %f 0 Z",
				bigRadius, bigRadius, bigRadius, bigRadius, 2*bigRadius, bigRadius, bigRadius, bigRadius)
			circle.X_Offset = -2 * bigRadius
			circle.Y_Offset = -bigRadius
			circle.RectAnchorType = svg.RECT_RIGHT
			rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, circle)

			if isSelected {
				circle.Color = svg.Lightblue.ToString()
				circle.FillOpacity = 0.6
			}
		}

		{
			circle := new(svg.RectAnchoredPath)
			circle.Stroke = svg.Black.ToString()
			circle.StrokeWidth = 2
			circle.StrokeOpacity = 1

			circle.Color = svg.Black.ToString()
			circle.FillOpacity = 1.0

			circle.Definition = fmt.Sprintf("M %f 0 A %f %f 0 0 1 %f %f A %f %f 0 0 1 %f 0 Z",
				smallRadius, smallRadius, smallRadius, smallRadius, 2*smallRadius, smallRadius, smallRadius, smallRadius)
			circle.X_Offset = -smallRadius - bigRadius
			circle.Y_Offset = -smallRadius
			circle.RectAnchorType = svg.RECT_RIGHT
			rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, circle)
		}
	}
}
