package models

import (
	"fmt"
	"time"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) generateTimeDiagram(diagram *Diagram, svgObject *svg.SVG) {

	layer := svgObject.Layers[0]

	verticalLinesLayer := (&svg.Layer{Name: "Vertical Line Layers"}).Stage(stager.svgStage)
	svgObject.Layers = append(svgObject.Layers, verticalLinesLayer)

	// If no duration, return early to prevent division by zero
	if diagram.ComputedDuration == 0 {
		return
	}

	LaneHeight := diagram.LaneHeight
	RatioBarToLaneHeight := diagram.RatioBarToLaneHeight
	barHeigth := LaneHeight * RatioBarToLaneHeight
	YTopMargin := diagram.YTopMargin

	var nbVisibleTaskGroups int
	for _, taskGroupShape := range diagram.TaskGroupShapes {
		if !taskGroupShape.IsHidden {
			nbVisibleTaskGroups++
		}
	}

	yTimeLine := LaneHeight*float64(nbVisibleTaskGroups) + YTopMargin

	XLeftText := diagram.XLeftText
	TextHeight := diagram.TextHeight

	XLeftLanes := diagram.XLeftLanes
	XRightMargin := diagram.XRightMargin

	// Time Line
	timeLine := new(svg.Line).Stage(stager.svgStage)
	timeLine.Name = "Time Line"
	timeLine.X1 = XLeftLanes
	timeLine.Y1 = yTimeLine
	timeLine.X2 = XRightMargin
	timeLine.Y2 = yTimeLine

	timeLine.Color = diagram.TimeLine_Color
	timeLine.FillOpacity = diagram.TimeLine_FillOpacity
	timeLine.Stroke = diagram.TimeLine_Stroke
	timeLine.StrokeWidth = diagram.TimeLine_StrokeWidth

	layer.Lines = append(layer.Lines, timeLine)

	// Dates
	DateYOffset := diagram.DateYOffset

	// put a date for every tick according to scale
	timeStep := diagram.TimeStep
	if timeStep <= 0 {
		timeStep = 1
	}
	timeStepScale := diagram.TimeStepScale
	if timeStepScale == "" {
		timeStepScale = YEARS
	}

	var ticks []time.Time
	start := diagram.ComputedStart
	var currentTick time.Time

	switch timeStepScale {
	case YEARS:
		currentTick = time.Date(start.Year(), time.January, 1, 0, 0, 0, 0, start.Location())
	case MONTHS:
		currentTick = time.Date(start.Year(), start.Month(), 1, 0, 0, 0, 0, start.Location())
	case WEEKS:
		offset := int(time.Monday - start.Weekday())
		if offset > 0 {
			offset -= 7
		}
		currentTick = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location()).AddDate(0, 0, offset)
	case DAYS:
		currentTick = time.Date(start.Year(), start.Month(), start.Day(), 0, 0, 0, 0, start.Location())
	default:
		currentTick = time.Date(start.Year(), time.January, 1, 0, 0, 0, 0, start.Location())
	}

	for currentTick.Before(diagram.ComputedEnd) || currentTick.Equal(diagram.ComputedEnd) {
		ticks = append(ticks, currentTick)

		switch timeStepScale {
		case YEARS:
			currentTick = currentTick.AddDate(timeStep, 0, 0)
		case MONTHS:
			currentTick = currentTick.AddDate(0, timeStep, 0)
		case WEEKS:
			currentTick = currentTick.AddDate(0, 0, 7*timeStep)
		case DAYS:
			currentTick = currentTick.AddDate(0, 0, timeStep)
		default:
			currentTick = currentTick.AddDate(timeStep, 0, 0)
		}
	}
	// Move tick drawing to the end so they are drawn over the lanes
	ticksToDraw := ticks

	// Ticks text only at this point

	for i := 0; i < len(ticks)-1; i++ {
		tick := ticks[i]
		nextTick := ticks[i+1]

		durationBetweenTickAndGanttStart := tick.Sub(diagram.ComputedStart)
		durationBetweenTickAndGanttStartRelativeToGanttDuration :=
			float64(durationBetweenTickAndGanttStart) / float64(diagram.ComputedDuration)

		xOriginal := XLeftLanes + (XRightMargin-XLeftLanes)*durationBetweenTickAndGanttStartRelativeToGanttDuration

		durationBetweenNextTickAndGanttStart := nextTick.Sub(diagram.ComputedStart)
		durationBetweenNextTickAndGanttStartRelativeToGanttDuration :=
			float64(durationBetweenNextTickAndGanttStart) / float64(diagram.ComputedDuration)

		xNextOriginal := XLeftLanes + (XRightMargin-XLeftLanes)*durationBetweenNextTickAndGanttStartRelativeToGanttDuration

		xVisible := xOriginal
		if xVisible < XLeftLanes {
			xVisible = XLeftLanes
		}
		if xVisible > XRightMargin {
			xVisible = XRightMargin
		}

		xNextVisible := xNextOriginal
		if xNextVisible < XLeftLanes {
			xNextVisible = XLeftLanes
		}
		if xNextVisible > XRightMargin {
			xNextVisible = XRightMargin
		}

		if xVisible == xNextVisible {
			continue
		}

		tickText := new(svg.Text).Stage(stager.svgStage)

		var tickLabel string
		switch timeStepScale {
		case YEARS:
			tickLabel = fmt.Sprintf("%d", tick.Year())
		case MONTHS:
			tickLabel = tick.Format("Jan '06")
		case WEEKS:
			_, week := tick.ISOWeek()
			tickLabel = fmt.Sprintf("W%02d", week)
		case DAYS:
			tickLabel = tick.Format("02 Jan")
		default:
			tickLabel = fmt.Sprintf("%d", tick.Year())
		}

		tickText.Name = tickLabel
		tickText.X = xVisible - float64(len(tickLabel))*4.0
		tickText.Content = tickText.Name
		tickText.Y = yTimeLine + DateYOffset
		tickText.Color = "black"
		tickText.FillOpacity = 1.0
		layer.Texts = append(layer.Texts, tickText)
	}

	// Lanes
	currentY := YTopMargin
	laneIndex := 0

	// vertical separator for lane headers
	headerSeparator := new(svg.Line).Stage(stager.svgStage)
	headerSeparator.Name = "Header Separator"
	layer.Lines = append(layer.Lines, headerSeparator)
	headerSeparator.X1 = XLeftLanes
	headerSeparator.X2 = XLeftLanes
	headerSeparator.Y1 = YTopMargin
	headerSeparator.Y2 = yTimeLine
	headerSeparator.Stroke = "darkgrey"
	headerSeparator.StrokeWidth = 1.0

	mapTaskGroup_TextY := make(map[*TaskGroup]float64, 0)

	for _, taskGroupShape := range diagram.TaskGroupShapes {
		if taskGroupShape.IsHidden {
			continue
		}
		taskGroup := taskGroupShape.TaskGroup

		laneSVG := new(svg.Rect).Stage(stager.svgStage)
		layer.Rects = append(layer.Rects, laneSVG)
		laneSVG.Name = taskGroup.Name
		laneSVG.X = XLeftLanes
		laneSVG.Y = currentY

		laneSVG.Width = XRightMargin - XLeftLanes
		laneSVG.Height = LaneHeight

		laneSVG.Color = "lightgrey"
		laneSVG.FillOpacity = 0.2

		if laneIndex%2 == 0 {
			laneSVG.Color = "white"
			laneSVG.FillOpacity = 0.8
		}
		laneIndex = laneIndex + 1
		laneSVG.StrokeWidth = 1.5

		laneText := new(svg.Text).Stage(stager.svgStage)
		laneText.Name = taskGroup.Name
		laneText.Content = laneText.Name
		laneText.X = XLeftText
		laneText.Y = currentY + LaneHeight/2.0 + TextHeight/2.0
		mapTaskGroup_TextY[taskGroup] = laneText.Y
		laneText.Color = "black"
		laneText.FillOpacity = 1.0
		layer.Texts = append(layer.Texts, laneText)

		// Tasks
		for _, task := range taskGroup.Tasks {
			taskShape, ok := diagram.map_Task_TaskShape[task]
			if !ok || taskShape.IsHidden {
				continue
			}

			rect4Bar := new(svg.Rect).Stage(stager.svgStage)
			diagram.map_Task_Rect[task] = rect4Bar
			diagram.map_SvgRect_TaskShape[rect4Bar] = taskShape

			layer.Rects = append(layer.Rects, rect4Bar)
			rect4Bar.Name = task.Name
			rect4Bar.IsSelectable = true
			rect4Bar.CanHaveRightHandle = true
			rect4Bar.CanHaveLeftHandle = true
			rect4Bar.CanMoveHorizontaly = true

			rect4Bar.OnUpdate = func(updatedRect *svg.Rect) {
				// We don't save size or position changes because time diagrams are driven by dates
				// but we want to allow clicking to open the probe form
				diffSize := rect4Bar.Width != updatedRect.Width || rect4Bar.Height != updatedRect.Height
				diffPosition := rect4Bar.X != updatedRect.X || rect4Bar.Y != updatedRect.Y

				if !diffSize && !diffPosition {
					stager.stage.CommitWithSuspendedCallbacks()
					stager.probeForm.FillUpFormFromGongstruct(task, "Task")
					stager.ux_tree()
				} else {
					stager.stage.CommitWithSuspendedCallbacks() // just revert UI to backend state
				}
			}

			var taskToDisplay = *task

			if diagram.UseManualStartAndEndDates {
				if task.Start.Before(diagram.ManualStart) {
					taskToDisplay.Start = diagram.ManualStart
				}
				if task.End.After(diagram.ManualEnd) {
					taskToDisplay.End = diagram.ManualEnd
				}
			}

			durationFromGanttStartToBarStart := taskToDisplay.Start.Sub(diagram.ComputedStart)
			durationBetweenBarStartAndGanttStartRelativeToGanttDuration :=
				float64(durationFromGanttStartToBarStart) / float64(diagram.ComputedDuration)

			durationFromBarEndAndBarStart := taskToDisplay.End.Sub(taskToDisplay.Start)
			durationBetweenBarEndAndBarStartRelativeToGanttDuration :=
				float64(durationFromBarEndAndBarStart) / float64(diagram.ComputedDuration)

			rect4Bar.X = XLeftLanes + (XRightMargin-XLeftLanes)*durationBetweenBarStartAndGanttStartRelativeToGanttDuration
			rect4Bar.Y = currentY + (LaneHeight-barHeigth)/2.0
			rect4Bar.Height = barHeigth
			rect4Bar.Width = (XRightMargin - XLeftLanes) * durationBetweenBarEndAndBarStartRelativeToGanttDuration

			rect4Bar.Color = "steelblue"
			rect4Bar.FillOpacity = 0.6
			rect4Bar.Stroke = "darkblue"
			rect4Bar.StrokeWidth = 1.0
			rect4Bar.RX = 4.0 // rounded corners

			if task.IsMilestone {
				// milestone rendering
				durationBetweenMilestoneAndGanttStart := taskToDisplay.Start.Sub(diagram.ComputedStart)
				durationBetweenMilestoneAndGanttStartRelativeToGanttDuration :=
					float64(durationBetweenMilestoneAndGanttStart) / float64(diagram.ComputedDuration)

				lineX := XLeftLanes + (XRightMargin-XLeftLanes)*durationBetweenMilestoneAndGanttStartRelativeToGanttDuration
				if task.DisplayVerticalBar {
					line := new(svg.Line).Stage(stager.svgStage)
					line.Name = task.Name
					verticalLinesLayer.Lines = append(verticalLinesLayer.Lines, line)
					line.X1 = lineX
					line.X2 = line.X1
					line.Y1 = YTopMargin
					line.Y2 = yTimeLine
					line.Stroke = "black"
					line.StrokeOpacity = 1
					line.StrokeWidth = 0.5
					line.StrokeDashArray = "2 2"
				}

				diamondWidth := 18.0

				// if no specific task group is assigned to display the diamond, we just put it on its own lane
				taskGroupsToDisplay := task.TaskGroupsToDisplay
				if len(taskGroupsToDisplay) == 0 {
					taskGroupsToDisplay = append(taskGroupsToDisplay, taskGroup)
				}

				for _, taskGroupToDisplay := range taskGroupsToDisplay {
					diamond := new(svg.Rect).Stage(stager.svgStage)
					layer.Rects = append(layer.Rects, diamond)
					diamond.Name = task.Name
					diamond.X = lineX - diamondWidth/2.0
					diamond.Y = mapTaskGroup_TextY[taskGroupToDisplay] - TextHeight/2.0 - diamondWidth/2.0
					diamond.Width = diamondWidth
					diamond.Height = diamondWidth
					diamond.Color = "crimson"
					diamond.FillOpacity = 1.0
					diamond.Stroke = "darkred"
					diamond.StrokeWidth = 1.0
					diamond.Transform = fmt.Sprintf("rotate(%d %d %d)", 45, int64(diamond.X+diamondWidth/2.0), int64(diamond.Y+diamondWidth/2.0))

					// dummy rect to hold the text so it renders on top of everything
					dummyRect := new(svg.Rect).Stage(stager.svgStage)
					layer.Rects = append(layer.Rects, dummyRect)
					dummyRect.Name = task.Name + " text holder"
					dummyRect.X = diamond.X + diamondWidth
					dummyRect.Y = diamond.Y
					dummyRect.Width = 0
					dummyRect.Height = diamondWidth
					dummyRect.FillOpacity = 0.0
					dummyRect.StrokeOpacity = 0.0

					milestoneText := new(svg.RectAnchoredText).Stage(stager.svgStage)
					milestoneText.Name = task.Name
					milestoneText.Content = task.Name
					milestoneText.X_Offset = XLeftText + task.XOffset
					milestoneText.Y_Offset = task.YOffset

					switch task.TextPosition {
					case TEXT_POSITION_TOP:
						milestoneText.RectAnchorType = svg.RECT_TOP
						milestoneText.TextAnchorType = svg.TEXT_ANCHOR_CENTER
					case TEXT_POSITION_BOTTOM:
						milestoneText.RectAnchorType = svg.RECT_BOTTOM
						milestoneText.TextAnchorType = svg.TEXT_ANCHOR_CENTER
					case TEXT_POSITION_LEFT:
						milestoneText.RectAnchorType = svg.RECT_LEFT
						milestoneText.TextAnchorType = svg.TEXT_ANCHOR_END
					case TEXT_POSITION_RIGHT:
						milestoneText.RectAnchorType = svg.RECT_RIGHT
						milestoneText.TextAnchorType = svg.TEXT_ANCHOR_START
					case TEXT_POSITION_CENTER:
						milestoneText.RectAnchorType = svg.RECT_CENTER_MIDDLE
						milestoneText.TextAnchorType = svg.TEXT_ANCHOR_CENTER
					default:
						milestoneText.RectAnchorType = svg.RECT_LEFT
						milestoneText.TextAnchorType = svg.TEXT_ANCHOR_START
					}

					milestoneText.Color = "black"
					milestoneText.FillOpacity = 1.0

					dummyRect.RectAnchoredTexts = append(dummyRect.RectAnchoredTexts, milestoneText)
				}
			} else {
				// bar text using RectAnchoredText to ensure it renders on top of the bar
				barText := new(svg.RectAnchoredText).Stage(stager.svgStage)
				barText.Name = task.Name
				barText.Content = task.Name
				barText.X_Offset = XLeftText + task.XOffset
				barText.Y_Offset = task.YOffset

				switch task.TextPosition {
				case TEXT_POSITION_TOP:
					barText.RectAnchorType = svg.RECT_TOP
					barText.TextAnchorType = svg.TEXT_ANCHOR_CENTER
				case TEXT_POSITION_BOTTOM:
					barText.RectAnchorType = svg.RECT_BOTTOM
					barText.TextAnchorType = svg.TEXT_ANCHOR_CENTER
				case TEXT_POSITION_LEFT:
					barText.RectAnchorType = svg.RECT_LEFT
					barText.TextAnchorType = svg.TEXT_ANCHOR_END
				case TEXT_POSITION_RIGHT:
					barText.RectAnchorType = svg.RECT_RIGHT
					barText.TextAnchorType = svg.TEXT_ANCHOR_START
				case TEXT_POSITION_CENTER:
					barText.RectAnchorType = svg.RECT_CENTER_MIDDLE
					barText.TextAnchorType = svg.TEXT_ANCHOR_CENTER
				default:
					barText.RectAnchorType = svg.RECT_LEFT
					barText.TextAnchorType = svg.TEXT_ANCHOR_START
				}

				barText.Color = "black"
				barText.FillOpacity = 1.0
				rect4Bar.RectAnchoredTexts = append(rect4Bar.RectAnchoredTexts, barText)
			}
		}

		currentY = currentY + LaneHeight
	}

	// Draw the vertical grid lines as thin Rects so they overlay the lanes perfectly
	if diagram.DrawVerticalTimeLines {
		for i := 0; i < len(ticksToDraw)-1; i++ {
			tick := ticksToDraw[i]

			durationBetweenTickAndGanttStart := tick.Sub(diagram.ComputedStart)
			durationBetweenTickAndGanttStartRelativeToGanttDuration :=
				float64(durationBetweenTickAndGanttStart) / float64(diagram.ComputedDuration)

			xOriginal := XLeftLanes + (XRightMargin-XLeftLanes)*durationBetweenTickAndGanttStartRelativeToGanttDuration

			if xOriginal >= XLeftLanes && xOriginal <= XRightMargin {
				gridLine := new(svg.Rect).Stage(stager.svgStage)
				gridLine.Name = fmt.Sprintf("grid line for %s", tick.Format("2006-01-02"))
				layer.Rects = append(layer.Rects, gridLine)
				gridLine.X = xOriginal
				gridLine.Y = YTopMargin
				gridLine.Width = 1.0 // thin solid line
				gridLine.Height = yTimeLine - YTopMargin
				gridLine.Color = "lightgrey"
				gridLine.FillOpacity = 0.5 // semi-transparent to blend nicely
				gridLine.StrokeOpacity = 0.0
			}
		}
	}
}
