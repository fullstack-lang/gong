package models

import (
	"fmt"
	"time"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) generateTimeDiagram(diagram *Diagram, layer *svg.Layer) {
	diagram.ComputeStartAndEndDate()

	// If no duration, return early to prevent division by zero
	if diagram.ComputedDuration == 0 {
		return
	}

	LaneHeight := diagram.LaneHeight
	RatioBarToLaneHeight := diagram.RatioBarToLaneHeight
	barHeigth := LaneHeight * RatioBarToLaneHeight
	YTopMargin := diagram.YTopMargin

	yTimeLine := LaneHeight*float64(len(diagram.TaskGroups)) + YTopMargin

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

	// put a date for every year
	for year := diagram.ComputedStart.Year(); year <= diagram.ComputedEnd.Year(); year++ {
		yearTime := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
		durationBetweenYearAndGanttStart := yearTime.Sub(diagram.ComputedStart)

		durationBetweenYearStartAndGanttStartRelativeToGanttDuration :=
			float64(durationBetweenYearAndGanttStart) / float64(diagram.ComputedDuration)

		yearText := new(svg.Text).Stage(stager.svgStage)
		yearText.Name = fmt.Sprintf("%d", year)
		yearText.Content = yearText.Name
		yearText.X = XLeftLanes + (XRightMargin-XLeftLanes)*durationBetweenYearStartAndGanttStartRelativeToGanttDuration
		yearText.Y = yTimeLine + DateYOffset
		yearText.Color = "black"
		yearText.FillOpacity = 1.0
		layer.Texts = append(layer.Texts, yearText)

		// draw the line
		lineForYear := new(svg.Line).Stage(stager.svgStage)
		lineForYear.Name = yearText.Name
		layer.Lines = append(layer.Lines, lineForYear)
		lineForYear.X1 = yearText.X
		lineForYear.X2 = lineForYear.X1
		lineForYear.Y1 = YTopMargin
		lineForYear.Y2 = yTimeLine
		lineForYear.Stroke = "black"
		lineForYear.StrokeWidth = 0.25
		lineForYear.StrokeDashArray = "4 4"
	}

	// Lanes
	currentY := YTopMargin
	laneIndex := 0

	for _, taskGroup := range diagram.TaskGroups {

		laneSVG := new(svg.Rect).Stage(stager.svgStage)
		layer.Rects = append(layer.Rects, laneSVG)
		laneSVG.Name = taskGroup.Name
		laneSVG.X = XLeftLanes
		laneSVG.Y = currentY

		laneSVG.Width = XRightMargin - XLeftLanes
		laneSVG.Height = LaneHeight

		laneSVG.Color = "grey"
		laneSVG.FillOpacity = 0.1

		if laneIndex%2 == 0 {
			laneSVG.Color = "black"
		}
		laneIndex = laneIndex + 1
		laneSVG.StrokeWidth = 1.5

		laneText := new(svg.Text).Stage(stager.svgStage)
		laneText.Name = taskGroup.Name
		laneText.Content = laneText.Name
		laneText.X = XLeftText
		laneText.Y = currentY + LaneHeight/2.0 + TextHeight/2.0
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

			rect4Bar.Color = "blue"
			rect4Bar.FillOpacity = 0.1
			rect4Bar.Stroke = "blue"
			rect4Bar.StrokeWidth = 0.5

			// bar text
			barText := new(svg.Text).Stage(stager.svgStage)
			barText.Name = task.Name
			barText.Content = barText.Name
			barText.X = rect4Bar.X + XLeftText
			barText.Y = currentY + LaneHeight/2.0 + TextHeight/2.0
			barText.Color = "black"
			barText.FillOpacity = 1.0
			layer.Texts = append(layer.Texts, barText)
		}

		currentY = currentY + LaneHeight
	}
}
