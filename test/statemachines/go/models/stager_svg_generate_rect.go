package models

import (
	"fmt"
	"reflect"
	"strings"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
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

			rect.Impl = NewRectShape_EtatProxy(
				&stateShape.RectShape,
				state,
				stager,
			)
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
			stateTitleText.Content = WrapString(stateShape.State.Name, int(rect.Width/stager.architecture.NbPixPerCharacter))
		}
		stateTitleText.Stroke = svg.Black.ToString()
		stateTitleText.StrokeWidth = 1
		stateTitleText.StrokeOpacity = 1
		stateTitleText.Color = svg.Black.ToString()
		stateTitleText.FillOpacity = 1

		stateTitleText.FontSize = "16px"
		stateTitleText.X_Offset = 0
		stateTitleText.Y_Offset = 30
		stateTitleText.RectAnchorType = svg.RECT_TOP
		stateTitleText.TextAnchorType = svg.TEXT_ANCHOR_CENTER

		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, stateTitleText)
	}

	// add the /Do actions
	currentY_Offset := stateTitleText.Y_Offset + float64(HeightBetween2AttributeShapes*(1+strings.Count(stateTitleText.Content, "\n")))
	for _, doAction := range state.DoActions {
		{
			doActionText := new(svg.RectAnchoredText)
			doActionText.Name = doAction.Name
			content := "/do " + doAction.Name

			if rect.Width > 0 {
				content = WrapString(content, int(rect.Width/stager.architecture.NbPixPerCharacter))
			}
			doActionText.Content = content
			doActionText.Stroke = svg.Black.ToString()
			doActionText.StrokeWidth = 1
			doActionText.StrokeOpacity = 1
			doActionText.Color = svg.Black.ToString()
			doActionText.FillOpacity = 1

			if doAction.Criticality == DoActionCritical {
				doActionText.Stroke = svg.Red.ToString()
				doActionText.Color = svg.Red.ToString()
			}

			doActionText.FontSize = "16px"
			doActionText.X_Offset = 0
			doActionText.Y_Offset = currentY_Offset
			doActionText.RectAnchorType = svg.RECT_TOP
			doActionText.TextAnchorType = svg.TEXT_ANCHOR_CENTER

			rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, doActionText)

			currentY_Offset += float64(HeightBetween2AttributeShapes * (1 + strings.Count(doActionText.Content, "\n")))
		}
	}

	if state.IsDecisionNode {
		stateTitleText.TextAnchorType = svg.TEXT_ANCHOR_START
		stateTitleText.RectAnchorType = svg.RECT_TOP_LEFT

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

		rect.StrokeOpacity = 0.0
		rect.FillOpacity = 0.0
		// force size
		rect.CanHaveBottomHandle = false
		rect.CanHaveTopHandle = false

		if isSelected {
			rect.Color = svg.Lightblue.ToString()
			rect.FillOpacity = 0.6
		}
	}

	if isStartState {
		stateTitleText.TextAnchorType = svg.TEXT_ANCHOR_START
		stateTitleText.RectAnchorType = svg.RECT_TOP_LEFT
		stateTitleText.DominantBaseline = svg.DominantBaselineCentral
		stateTitleText.WhiteSpace = svg.WhiteSpaceEnumPre
		stateTitleText.X_Offset = 0
		stateTitleText.Y_Offset = 0

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

		// force size
		rect.CanHaveBottomHandle = false
		rect.CanHaveTopHandle = false

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

		rect.StrokeOpacity = 0.0
		rect.FillOpacity = 0.0
	}

	if state.IsEndState {
		stateTitleText.TextAnchorType = svg.TEXT_ANCHOR_START
		stateTitleText.RectAnchorType = svg.RECT_TOP_LEFT
		stateTitleText.DominantBaseline = svg.DominantBaselineCentral
		stateTitleText.WhiteSpace = svg.WhiteSpaceEnumPre
		stateTitleText.X_Offset = 0
		stateTitleText.Y_Offset = 0

		rect.CanHaveBottomHandle = false
		rect.CanHaveTopHandle = false
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

		rect.StrokeOpacity = 0.0
		rect.FillOpacity = 0.0

	}

	if state.IsFictif {
		rect.StrokeDashArray = "5 5"
	}

	layer.Rects = append(layer.Rects, rect)
	return rect
}

func NewRectShape_EtatProxy(
	stateShapeInterface RectShapeInterface,
	state *State,
	stager *Stager,
) (proxy *ShapeRect_EtatProxy) {
	proxy = new(ShapeRect_EtatProxy)
	proxy.rectShapeInterface = stateShapeInterface
	proxy.state = state
	proxy.stager = stager
	return
}

type ShapeRect_EtatProxy struct {
	rectShapeInterface RectShapeInterface
	state              *State
	stager             *Stager
}

// RectUpdated implements models.RectImplInterface.
func (p *ShapeRect_EtatProxy) RectUpdated(updatedRect *svg.Rect) {

	diffSize :=
		p.rectShapeInterface.GetWidth() != updatedRect.Width ||
			p.rectShapeInterface.GetHeight() != updatedRect.Height

	diffPosition :=
		p.rectShapeInterface.GetX() != updatedRect.X ||
			p.rectShapeInterface.GetY() != updatedRect.Y

	p.rectShapeInterface.SetX(updatedRect.X)
	p.rectShapeInterface.SetY(updatedRect.Y)
	p.rectShapeInterface.SetWidth(updatedRect.Width)
	p.rectShapeInterface.SetHeight(updatedRect.Height)

	if !diffSize && !diffPosition {
		p.stager.stage.CommitWithSuspendedCallbacks()
		p.stager.probeForm.FillUpFormFromGongstruct(p.state, "State")
	}

	if diffPosition {
		// Issue #7, this will allow multiple rect to be moved together
		p.stager.stage.CommitWithSuspendedCallbacks()
	}

	if diffSize {
		p.stager.stage.Commit()
	}

}
