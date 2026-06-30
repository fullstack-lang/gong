package models

import (
	"fmt"
	"time"

	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

var axisHandleBorderLength = 20.0

type Direction string

const (
	DIRECTION_TOP    Direction = "TOP"
	DIRECTION_BOTTOM Direction = "BOTTOM"
	DIRECTION_LEFT   Direction = "LEFT"
	DIRECTION_RIGHT  Direction = "RIGHT"
)



func (stager *Stager) drawVerticalAxis(
	diagram *Diagram,
	defaultLayer *gongsvg_models.Layer,
) {
	// creation of 2 transparant rects, one at each ends of the vertical
	verticalAxisTopRect := new(gongsvg_models.Rect)
	verticalAxisTopRect.Name = "Vertical axis top rect"
	defaultLayer.Rects = append(defaultLayer.Rects, verticalAxisTopRect)
	verticalAxisTopRect.X = diagram.AxisOrign_X - axisHandleBorderLength/2.0
	verticalAxisTopRect.Y = diagram.VerticalAxis_Top_Y - axisHandleBorderLength
	verticalAxisTopRect.Width = axisHandleBorderLength
	verticalAxisTopRect.Height = axisHandleBorderLength

	if diagram.IsInDrawMode {
		verticalAxisTopRect.Stroke = gongsvg_models.Bisque.ToString()
		verticalAxisTopRect.StrokeWidth = 4
		verticalAxisTopRect.StrokeOpacity = 1
		verticalAxisTopRect.IsSelectable = true
		verticalAxisTopRect.CanMoveVerticaly = true
		verticalAxisTopRect.OnMove = func(x, y float64) {
			diagram.VerticalAxis_Top_Y = y + axisHandleBorderLength
			stager.stage.Commit()
		}
	}

	verticalAxisBottomRect := new(gongsvg_models.Rect)
	verticalAxisBottomRect.Name = "Vertical axis bottom rect"
	defaultLayer.Rects = append(defaultLayer.Rects, verticalAxisBottomRect)
	verticalAxisBottomRect.X = diagram.AxisOrign_X - axisHandleBorderLength/2.0
	verticalAxisBottomRect.Y = diagram.VerticalAxis_Bottom_Y
	verticalAxisBottomRect.Width = axisHandleBorderLength
	verticalAxisBottomRect.Height = axisHandleBorderLength

	if diagram.IsInDrawMode {
		verticalAxisBottomRect.Stroke = gongsvg_models.Bisque.ToString()
		verticalAxisBottomRect.StrokeWidth = verticalAxisTopRect.StrokeWidth
		verticalAxisBottomRect.StrokeOpacity = 1
		verticalAxisBottomRect.CanMoveVerticaly = true
		verticalAxisBottomRect.OnMove = func(x, y float64) {
			diagram.VerticalAxis_Bottom_Y = y
			stager.stage.Commit()
		}
	}

	verticalAxis := new(gongsvg_models.Link)
	defaultLayer.Links = append(defaultLayer.Links, verticalAxis)

	verticalAxis.StrokeWidth = diagram.VerticalAxis_StrokeWidth
	verticalAxis.StrokeOpacity = 1
	verticalAxis.Name = "Vertical Axis"
	verticalAxis.Stroke = gongsvg_models.Black.ToString()
	verticalAxis.Start = verticalAxisTopRect
	verticalAxis.End = verticalAxisBottomRect

	verticalAxis.HasStartArrow = true
	verticalAxis.StartArrowSize = 8
	verticalAxis.HasEndArrow = true
	verticalAxis.CornerOffsetRatio = 2.0
	verticalAxis.EndArrowSize = 8
	verticalAxis.Type = gongsvg_models.LINK_TYPE_FLOATING_ORTHOGONAL

	verticalAxis.StartOrientation = gongsvg_models.ORIENTATION_VERTICAL
	verticalAxis.EndOrientation = gongsvg_models.ORIENTATION_VERTICAL
	verticalAxis.StartRatio = 0.5
	verticalAxis.EndRatio = 0.5
}

func (stager *Stager) drawHorizontalAxis(
	diagram *Diagram,
	axisLayer *gongsvg_models.Layer,
) {
	horizontalAxisLeftHandle := new(gongsvg_models.Rect)
	horizontalAxisLeftHandle.Name = "Horizontal axis left handle"
	axisLayer.Rects = append(axisLayer.Rects, horizontalAxisLeftHandle)
	horizontalAxisLeftHandle.X = diagram.AxisOrign_X - axisHandleBorderLength
	horizontalAxisLeftHandle.Y = diagram.AxisOrign_Y - axisHandleBorderLength/2.0
	horizontalAxisLeftHandle.Width = axisHandleBorderLength
	horizontalAxisLeftHandle.Height = axisHandleBorderLength

	if diagram.IsInDrawMode {
		horizontalAxisLeftHandle.Stroke = gongsvg_models.Bisque.ToString()
		horizontalAxisLeftHandle.StrokeWidth = 4
		horizontalAxisLeftHandle.StrokeOpacity = 1
		horizontalAxisLeftHandle.IsSelectable = true
		horizontalAxisLeftHandle.CanMoveHorizontaly = true
		horizontalAxisLeftHandle.CanMoveVerticaly = true
		horizontalAxisLeftHandle.OnMove = func(x, y float64) {
			diagram.AxisOrign_X = x + axisHandleBorderLength
			diagram.AxisOrign_Y = y + axisHandleBorderLength/2.0
			stager.stage.Commit()
		}
	}

	horizontalAxisRightHandle := new(gongsvg_models.Rect)
	horizontalAxisRightHandle.Name = "Horizontal axis rigth handle"
	axisLayer.Rects = append(axisLayer.Rects, horizontalAxisRightHandle)
	horizontalAxisRightHandle.X = diagram.HorizontalAxis_Right_X
	horizontalAxisRightHandle.Y = diagram.AxisOrign_Y - axisHandleBorderLength/2.0
	horizontalAxisRightHandle.Width = axisHandleBorderLength
	horizontalAxisRightHandle.Height = axisHandleBorderLength

	if diagram.IsInDrawMode {
		horizontalAxisRightHandle.Stroke = gongsvg_models.Bisque.ToString()
		horizontalAxisRightHandle.StrokeWidth = horizontalAxisLeftHandle.StrokeWidth
		horizontalAxisRightHandle.StrokeOpacity = 1
		horizontalAxisRightHandle.CanMoveHorizontaly = true
		horizontalAxisRightHandle.OnMove = func(x, y float64) {
			diagram.HorizontalAxis_Right_X = x
			stager.stage.Commit()
		}
	}

	horizontalAxis := new(gongsvg_models.Link)
	axisLayer.Links = append(axisLayer.Links, horizontalAxis)

	horizontalAxis.StrokeWidth = diagram.VerticalAxis_StrokeWidth
	horizontalAxis.StrokeOpacity = 1
	horizontalAxis.Name = "Horizontal Axis"
	horizontalAxis.Stroke = gongsvg_models.Black.ToString()
	horizontalAxis.Start = horizontalAxisLeftHandle
	horizontalAxis.End = horizontalAxisRightHandle

	horizontalAxis.HasStartArrow = false
	horizontalAxis.HasEndArrow = true
	horizontalAxis.CornerOffsetRatio = 2.0
	horizontalAxis.EndArrowSize = 8
	horizontalAxis.Type = gongsvg_models.LINK_TYPE_FLOATING_ORTHOGONAL

	horizontalAxis.StartOrientation = gongsvg_models.ORIENTATION_HORIZONTAL
	horizontalAxis.EndOrientation = gongsvg_models.ORIENTATION_HORIZONTAL
	horizontalAxis.StartRatio = 0.5
	horizontalAxis.EndRatio = 0.5

	if diagram.Start == diagram.End && diagram.Start.After(diagram.End) {
		return
	}
	firsts_of_January := diagram.GetYearsWithJanuaryFirstBetween()

	if diagram.NumberOfYearsBetweenTicks > 1 {
		sampledDownFirstsOfJanuary := make([]time.Time, 0)
		for idx, first_of_January := range firsts_of_January {
			if idx%diagram.NumberOfYearsBetweenTicks == 0 {
				sampledDownFirstsOfJanuary = append(sampledDownFirstsOfJanuary, first_of_January)
			}
		}
		firsts_of_January = sampledDownFirstsOfJanuary
	}

	for _, first_of_January := range firsts_of_January {
		nextFirst_if_january := first_of_January.AddDate(1, 0, 0)
		if diagram.NumberOfYearsBetweenTicks > 1 {
			nextFirst_if_january = nextFirst_if_january.AddDate(diagram.NumberOfYearsBetweenTicks-1, 0, 0)
		}

		x := TimeToX(diagram, first_of_January)
		xNext := TimeToX(diagram, nextFirst_if_january)
		y := horizontalAxis.Start.Y

		text := new(gongsvg_models.Text)
		text.Stroke = gongsvg_models.Black.ToString()
		text.StrokeWidth = 1.0
		text.StrokeOpacity = 0.5
		if diagram.NumberOfYearsBetweenTicks > 1 {
			text.Name = fmt.Sprintf("%d -> %d", first_of_January.Year(), nextFirst_if_january.Year())
			text.X = (x+xNext)/2 - 50
		} else {
			text.Name = fmt.Sprintf("%d", first_of_January.Year())
			text.X = (x+xNext)/2 - 20
		}
		text.Content = text.Name
		text.Color = gongsvg_models.Black.ToString()
		text.FillOpacity = 1.0
		text.Y = y + 35

		axisLayer.Texts = append(axisLayer.Texts, text)

		horizontalAxiTick := new(gongsvg_models.Line)
		axisLayer.Lines = append(axisLayer.Lines, horizontalAxiTick)

		horizontalAxiTick.Name = fmt.Sprintf("tick for year %d", first_of_January.Year())
		horizontalAxiTick.X1 = x
		horizontalAxiTick.X2 = x
		horizontalAxiTick.Y1 = y + axisHandleBorderLength/2
		horizontalAxiTick.Y2 = y + axisHandleBorderLength/2 + 5

		horizontalAxiTick.StrokeWidth = diagram.VerticalAxis_StrokeWidth
		horizontalAxiTick.StrokeOpacity = 1
		horizontalAxiTick.Stroke = gongsvg_models.Black.ToString()
	}
}
