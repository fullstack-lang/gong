package models

import (
	"reflect"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	"github.com/fullstack-lang/gong/pkg/strutils"
)

const HeightBetween2AttributeShapes = 20

type diagramInterface interface {
	IsEditable() bool
}

func (stager *Stager) svgGenerateTaskRect(
	diagram diagramInterface, // might be ni
	TaskShape *TaskShape,
	layer *svg.Layer,
) *svg.Rect {
	root := stager.root

	Task := TaskShape.Task

	rect := new(svg.Rect)
	rect.Name = Task.Name
	rect.Stroke = svg.Black.ToString()
	rect.StrokeWidth = 2
	rect.StrokeOpacity = 1
	rect.X = TaskShape.X
	rect.Y = TaskShape.Y
	rect.Width = TaskShape.Width
	rect.Height = TaskShape.Height
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

			rect.Impl = NewRectShape_TaskProxy(
				&TaskShape.RectShape,
				Task,
				stager,
			)
			// for allowing later Stage() on the rect shape
			TaskShape.receiver = TaskShape
		}
	}

	TaskTitleText := new(svg.RectAnchoredText)
	{
		TaskTitleText.Name = Task.Name
		TaskTitleText.Content = Task.Name

		if rect.Width > 0 {
			TaskTitleText.Content = strutils.WrapString(TaskShape.Task.Name, int(rect.Width/root.NbPixPerCharacter))
		}
		TaskTitleText.Stroke = svg.Black.ToString()
		TaskTitleText.StrokeWidth = 1
		TaskTitleText.StrokeOpacity = 1
		TaskTitleText.Color = svg.Black.ToString()
		TaskTitleText.FillOpacity = 1

		TaskTitleText.FontSize = "16px"
		TaskTitleText.X_Offset = 0
		TaskTitleText.Y_Offset = 30
		TaskTitleText.RectAnchorType = svg.RECT_TOP
		TaskTitleText.TextAnchorType = svg.TEXT_ANCHOR_CENTER

		rect.RectAnchoredTexts = append(rect.RectAnchoredTexts, TaskTitleText)
	}

	layer.Rects = append(layer.Rects, rect)
	return rect
}

func NewRectShape_TaskProxy(
	TaskShapeInterface RectShapeInterface,
	Task *Task,
	stager *Stager,
) (proxy *ShapeRect_TaskProxy) {
	proxy = new(ShapeRect_TaskProxy)
	proxy.rectShapeInterface = TaskShapeInterface
	proxy.Task = Task
	proxy.stager = stager
	return
}

type ShapeRect_TaskProxy struct {
	rectShapeInterface RectShapeInterface
	Task               *Task
	stager             *Stager
}

// RectUpdated implements models.RectImplInterface.
func (p *ShapeRect_TaskProxy) RectUpdated(updatedRect *svg.Rect) {
	diffSize := p.rectShapeInterface.GetWidth() != updatedRect.Width ||
		p.rectShapeInterface.GetHeight() != updatedRect.Height

	diffPosition := p.rectShapeInterface.GetX() != updatedRect.X ||
		p.rectShapeInterface.GetY() != updatedRect.Y

	p.rectShapeInterface.SetX(updatedRect.X)
	p.rectShapeInterface.SetY(updatedRect.Y)
	p.rectShapeInterface.SetWidth(updatedRect.Width)
	p.rectShapeInterface.SetHeight(updatedRect.Height)

	if !diffSize && !diffPosition {
		p.stager.stage.CommitWithSuspendedCallbacks()
		p.stager.probeForm.FillUpFormFromGongstruct(p.Task, "Task")
	}

	if diffPosition {
		// Issue #7, this will allow multiple rect to be moved together
		p.stager.stage.CommitWithSuspendedCallbacks()
	}

	if diffSize {
		p.stager.stage.Commit()
	}
}
