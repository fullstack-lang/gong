package models

import (
	"fmt"
	"log"
	"math"
	"time"

	"github.com/fullstack-lang/gong/lib/strutils"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) generateTimeDiagram(diagram *Diagram, svgObject *svg.SVG) {

	layer := svgObject.Layers[0]

	verticalLinesLayer := (&svg.Layer{Name: "Vertical Line Layers"})
	svgObject.Layers = append(svgObject.Layers, verticalLinesLayer)

	// If no duration, return early to prevent division by zero
	if diagram.HideWeekendsPeriod {
		if diagram.workDurationBetween(diagram.ComputedStart, diagram.ComputedEnd) <= 0 {
			return
		}
	} else if diagram.ComputedDuration == 0 {
		return
	}

	LaneHeight := diagram.LaneHeight
	RatioBarToLaneHeight := diagram.RatioBarToLaneHeight
	barHeigth := LaneHeight * RatioBarToLaneHeight

	var nbVisibleTaskGroups int
	for _, taskGroupShape := range diagram.TaskGroupShapes {
		if !taskGroupShape.IsHidden {
			nbVisibleTaskGroups++
		}
	}

	yTimeLine := LaneHeight*float64(nbVisibleTaskGroups) + diagram.YTopMargin

	// Time Line
	timeLine := new(svg.Line)
	timeLine.Name = "Time Line"
	timeLine.X1 = diagram.XLeftLanes
	timeLine.Y1 = yTimeLine
	timeLine.X2 = diagram.XRightMargin
	timeLine.Y2 = yTimeLine

	timeLine.Color = diagram.TimeLine_Color
	timeLine.FillOpacity = diagram.TimeLine_FillOpacity
	timeLine.Stroke = diagram.TimeLine_Stroke
	timeLine.StrokeWidth = diagram.TimeLine_StrokeWidth

	layer.Lines = append(layer.Lines, timeLine)

	// Dates
	DateYOffset := diagram.DateYOffset

	// put a date for every tick according to scale
	stager.drawTimeLine(diagram, diagram.XLeftLanes, diagram.XRightMargin, yTimeLine, DateYOffset, layer, verticalLinesLayer, diagram.YTopMargin)

	// Lanes
	currentY := diagram.YTopMargin
	laneIndex := 0

	// vertical separator for lane headers
	headerSeparator := new(svg.Line)
	headerSeparator.Name = "Header Separator"
	layer.Lines = append(layer.Lines, headerSeparator)
	headerSeparator.X1 = diagram.XLeftLanes
	headerSeparator.X2 = diagram.XLeftLanes
	headerSeparator.Y1 = diagram.YTopMargin
	headerSeparator.Y2 = yTimeLine
	headerSeparator.Stroke = "darkgrey"
	headerSeparator.StrokeWidth = 1.0

	mapTaskGroup_TextY := make(map[*TaskGroup]float64, 0)

	for _, taskGroupShape := range diagram.TaskGroupShapes {
		if taskGroupShape.IsHidden {
			continue
		}
		taskGroup := taskGroupShape.TaskGroup

		laneSVG := new(svg.Rect)
		layer.Rects = append(layer.Rects, laneSVG)
		laneSVG.Name = taskGroup.Name
		laneSVG.X = diagram.XLeftLanes
		laneSVG.Y = currentY

		laneSVG.Width = diagram.XRightMargin - diagram.XLeftLanes
		laneSVG.Height = LaneHeight

		laneSVG.Color = "lightgrey"
		laneSVG.FillOpacity = 0.2

		if laneIndex%2 == 0 {
			laneSVG.Color = "white"
			laneSVG.FillOpacity = 0.8
		}
		laneIndex = laneIndex + 1
		laneSVG.StrokeWidth = 1.5

		laneText := new(svg.Text)
		laneText.Name = taskGroup.Name
		laneText.Content = laneText.Name
		laneText.X = diagram.XLeftText
		laneText.Y = currentY + LaneHeight/2.0 + diagram.TextHeight/2.0
		mapTaskGroup_TextY[taskGroup] = laneText.Y
		laneText.Color = "black"
		laneText.FillOpacity = 1.0
		layer.Texts = append(layer.Texts, laneText)

		// Tasks
		for _, task := range taskGroup.Tasks {
			taskShape, ok := diagram.map_Task_TaskShape[task]
			if !ok {
				log.Println("task has no shape", task.Name)
				continue
			}
			if taskShape.IsHidden {
				log.Println("task", taskShape.Name, "is hidden")
				continue
			}

			rect4Bar := stager.displayTask(diagram, task, taskShape, layer, currentY, LaneHeight, barHeigth) // rounded corners

			if task.IsMilestone {
				// milestone rendering
				stager.displayMilestone(diagram, task, taskShape, verticalLinesLayer, yTimeLine, taskGroup, layer, mapTaskGroup_TextY)
			} else {
				stager.displayTaskCompletion(task, rect4Bar, barHeigth)
				// bar text using RectAnchoredText to ensure it renders on top of the bar
				stager.displayTaskTitle(task, diagram, rect4Bar)
			}
		}

		currentY = currentY + LaneHeight
	}

}

func (stager *Stager) displayTask(diagram *Diagram, task *Task, taskShape *TaskShape, layer *svg.Layer, currentY float64, LaneHeight float64, barHeigth float64) *svg.Rect {
	rect4Bar := new(svg.Rect)
	diagram.map_Task_Rect[task] = rect4Bar
	diagram.map_SvgRect_TaskShape[rect4Bar] = taskShape

	layer.Rects = append(layer.Rects, rect4Bar)
	rect4Bar.Name = task.Name
	rect4Bar.IsSelectable = true
	rect4Bar.CanHaveRightHandle = true
	rect4Bar.CanHaveLeftHandle = true
	rect4Bar.CanMoveHorizontaly = true

	rect4Bar.OnSelect = func() {
		stager.stage.CommitWithSuspendedCallbacks()
		stager.probeForm.FillUpFormFromGongstruct(task, "Task")
		stager.ux_tree()
	}
	rect4Bar.OnMove = func(x, y float64) {
		stager.stage.CommitWithSuspendedCallbacks() // just revert UI to backend state
	}
	rect4Bar.OnResize = func(x, y, width, height float64) {
		stager.stage.CommitWithSuspendedCallbacks() // just revert UI to backend state
	}

	var taskToDisplay = *task

	endToDisplay := taskToDisplay.End
	if taskToDisplay.IsAllDay && !taskToDisplay.IsMilestone {
		endToDisplay = endToDisplay.AddDate(0, 0, 1)
	}

	if diagram.UseManualStartAndEndDates {
		if task.Start.Before(diagram.ManualStart) {
			taskToDisplay.Start = diagram.ManualStart
		}
		if endToDisplay.After(diagram.ManualEnd) {
			endToDisplay = diagram.ManualEnd
		}
	}

	rect4Bar.X = diagram.dateToX(taskToDisplay.Start)
	rect4Bar.Y = currentY + (LaneHeight-barHeigth)/2.0
	rect4Bar.Height = barHeigth
	endX := diagram.dateToX(endToDisplay)
	rect4Bar.Width = endX - rect4Bar.X
	if rect4Bar.Width < 0 {
		rect4Bar.Width = 0
	}

	rect4Bar.Color = "steelblue"
	rect4Bar.FillOpacity = 0.6
	rect4Bar.Stroke = "darkblue"
	rect4Bar.StrokeWidth = 1.0
	rect4Bar.RX = 4.0
	return rect4Bar
}

func (stager *Stager) displayTaskCompletion(task *Task, rect *svg.Rect, barHeight float64) {
	if !task.IsWithCompletion {
		return
	}
	rect.IsScalingProportionally = false
	distanceFromBorder := 10.0
	iconHeight := 54.0
	yOffset := (barHeight - iconHeight) / 2.0
	if yOffset < 0 {
		yOffset = 0
	}
	if path := createCompletionRectAnchoredPath(task.Completion, distanceFromBorder, yOffset); path != nil {
		rect.RectAnchoredPaths = append(rect.RectAnchoredPaths, path)
	}
}

func (stager *Stager) displayTaskTitle(task *Task, diagram *Diagram, rect4Bar *svg.Rect) {
	barText := new(svg.RectAnchoredText)
	barText.Name = task.Name

	content := task.Name
	if diagram != nil && diagram.GetIsShowPrefix() {
		content = task.GetComputedPrefix() + " " + content
	}
	if task.IsImport && task.ReferencedTask != nil {
		content = "🔗 " + task.ReferencedTask.Name
		if diagram != nil && diagram.GetIsShowPrefix() {
			content = task.GetComputedPrefix() + " " + content
		}
	}

	root := stager.getRootLibrary()
	nbPixPerChar := 8.0
	if root != nil && root.NbPixPerCharacter > 0 {
		nbPixPerChar = root.NbPixPerCharacter
	}

	distanceFromBorder := 10.0
	iconWidth := 25.0
	leftPadding := diagram.XLeftText
	if task.IsWithCompletion {
		leftPadding += distanceFromBorder + iconWidth
	}

	switch task.TextPosition {
	case TEXT_POSITION_TOP:
		barText.RectAnchorType = svg.RECT_TOP
		barText.TextAnchorType = svg.TEXT_ANCHOR_CENTER
		barText.X_Offset = task.XOffset
	case TEXT_POSITION_BOTTOM:
		barText.RectAnchorType = svg.RECT_BOTTOM
		barText.TextAnchorType = svg.TEXT_ANCHOR_CENTER
		barText.X_Offset = task.XOffset
	case TEXT_POSITION_LEFT:
		barText.RectAnchorType = svg.RECT_LEFT
		barText.TextAnchorType = svg.TEXT_ANCHOR_END
		barText.X_Offset = -diagram.XLeftText + task.XOffset
	case TEXT_POSITION_RIGHT:
		barText.RectAnchorType = svg.RECT_RIGHT
		barText.TextAnchorType = svg.TEXT_ANCHOR_START
		barText.X_Offset = diagram.XLeftText + task.XOffset
	case TEXT_POSITION_CENTER:
		barText.RectAnchorType = svg.RECT_CENTER_MIDDLE
		barText.TextAnchorType = svg.TEXT_ANCHOR_CENTER
		barText.DominantBaseline = svg.DominantBaselineCentral
		if task.IsWithCompletion {
			barText.X_Offset = (distanceFromBorder + iconWidth)/2.0 + task.XOffset
		} else {
			barText.X_Offset = task.XOffset
		}
	default:
		barText.RectAnchorType = svg.RECT_LEFT_MIDDLE
		barText.TextAnchorType = svg.TEXT_ANCHOR_START
		barText.DominantBaseline = svg.DominantBaselineCentral
		barText.X_Offset = leftPadding + task.XOffset
	}
	barText.Y_Offset = task.YOffset

	// Wrap text for tasks where text is positioned inside or aligned with the bar
	if task.TextPosition != TEXT_POSITION_LEFT && task.TextPosition != TEXT_POSITION_RIGHT {
		margin := 2 * diagram.XLeftText
		if margin <= 0 {
			margin = 20.0
		}
		if task.IsWithCompletion {
			margin += distanceFromBorder + iconWidth
		}
		availableWidth := rect4Bar.Width - margin
		if availableWidth <= 0 && rect4Bar.Width > 0 {
			availableWidth = rect4Bar.Width
		}
		if availableWidth > 0 && nbPixPerChar > 0 {
			cutoff := int(availableWidth / nbPixPerChar)
			if cutoff > 0 {
				content = strutils.WrapStringPreservingNewlines(content, cutoff)
			}
		}
	}

	barText.Content = content
	barText.Color = "black"
	barText.FillOpacity = 1.0
	rect4Bar.RectAnchoredTexts = append(rect4Bar.RectAnchoredTexts, barText)
}

func (stager *Stager) displayMilestone(diagram *Diagram, task *Task, taskShape *TaskShape, verticalLinesLayer *svg.Layer, yTimeLine float64, taskGroup *TaskGroup, layer *svg.Layer, mapTaskGroup_TextY map[*TaskGroup]float64) {
	lineX := diagram.dateToX(task.Start)
	if task.DisplayVerticalBar {
		line := new(svg.Line)
		line.Name = task.Name
		verticalLinesLayer.Lines = append(verticalLinesLayer.Lines, line)
		line.X1 = lineX
		line.X2 = line.X1
		line.Y1 = diagram.YTopMargin
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
		diamond := new(svg.Rect)
		diagram.map_SvgRect_TaskShape[diamond] = taskShape

		layer.Rects = append(layer.Rects, diamond)
		diamond.Name = task.Name

		diamond.OnSelect = func() {
			stager.stage.CommitWithSuspendedCallbacks()
			stager.probeForm.FillUpFormFromGongstruct(task, "Task")
			stager.ux_tree()
		}
		diamond.OnMove = func(x, y float64) {
			stager.stage.CommitWithSuspendedCallbacks() // just revert UI to backend state
		}
		diamond.OnResize = func(x, y, width, height float64) {
			stager.stage.CommitWithSuspendedCallbacks() // just revert UI to backend state
		}

		diamond.X = lineX - diamondWidth/2.0
		diamond.Y = mapTaskGroup_TextY[taskGroupToDisplay] - diagram.TextHeight/2.0 - diamondWidth/2.0
		diamond.Width = diamondWidth
		diamond.Height = diamondWidth
		diamond.Color = "crimson"
		diamond.FillOpacity = 1.0
		diamond.Stroke = "darkred"
		diamond.StrokeWidth = 1.0
		diamond.Transform = fmt.Sprintf("rotate(%d %d %d)", 45, int64(diamond.X+diamondWidth/2.0), int64(diamond.Y+diamondWidth/2.0))

		// dummy rect to hold the text so it renders on top of everything
		dummyRect := new(svg.Rect)
		layer.Rects = append(layer.Rects, dummyRect)
		dummyRect.Name = task.Name + " text holder"
		dummyRect.X = diamond.X + diamondWidth
		dummyRect.Y = diamond.Y
		dummyRect.Width = 0
		dummyRect.Height = diamondWidth
		dummyRect.FillOpacity = 0.0
		dummyRect.StrokeOpacity = 0.0

		milestoneText := new(svg.RectAnchoredText)
		milestoneText.Name = task.Name
		milestoneText.Content = task.Name
		milestoneText.X_Offset = diagram.XLeftText + task.XOffset
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
}

func (stager *Stager) drawTimeLine(diagram *Diagram, XLeftLanes float64, XRightMargin float64, yTimeLine float64, DateYOffset float64, layer *svg.Layer, verticalLinesLayer *svg.Layer, YTopMargin float64) {
	timeStep := diagram.TimeStep
	if timeStep <= 0 {
		timeStep = 1
	}
	timeStepScale := diagram.TimeStepScale
	if timeStepScale == "" {
		timeStepScale = MONTHS
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
		if diagram.HideWeekendsPeriod && timeStepScale == DAYS && (currentTick.Weekday() == time.Saturday || currentTick.Weekday() == time.Sunday) {
			// skip weekend ticks when weekends are hidden in day view
		} else {
			ticks = append(ticks, currentTick)
		}

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

	for i := 0; i < len(ticks); i++ {
		tick := ticks[i]

		xOriginal := diagram.dateToX(tick)

		xVisible := xOriginal
		if xVisible < XLeftLanes {
			xVisible = XLeftLanes
		}
		if xVisible > XRightMargin {
			xVisible = XRightMargin
		}

		if i < len(ticks)-1 {
			nextTick := ticks[i+1]
			xNextOriginal := diagram.dateToX(nextTick)

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
		} else {
			if xOriginal < XLeftLanes {
				continue
			}
		}

		tickText := new(svg.Text)

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

	// Draw the vertical grid lines as thin Rects so they overlay the lanes perfectly
	if diagram.DrawVerticalTimeLines {
		for i := range ticksToDraw {
			tick := ticksToDraw[i]

			xOriginal := diagram.dateToX(tick)

			if xOriginal >= XLeftLanes && xOriginal <= XRightMargin {
				gridLine := new(svg.Line)
				gridLine.Name = fmt.Sprintf("grid line for %s", tick.Format("2006-01-02"))
				verticalLinesLayer.Lines = append(verticalLinesLayer.Lines, gridLine)
				gridLine.X1 = xOriginal
				gridLine.Y1 = YTopMargin
				gridLine.X2 = xOriginal
				gridLine.Y2 = yTimeLine

				gridLine.Stroke = "grey"
				gridLine.StrokeOpacity = 1.0
				gridLine.StrokeWidth = 1.0
				gridLine.StrokeDashArray = "5 5"
			}
		}
	}
}

// workTimeFromEpoch returns the cumulative working duration from a fixed epoch Monday (2000-01-03 00:00:00)
// in the timezone of t up to t. Weekends (Saturday 00:00:00 to Monday 00:00:00) contribute zero working time.
func workTimeFromEpoch(t time.Time) time.Duration {
	loc := t.Location()
	epoch := time.Date(2000, 1, 3, 0, 0, 0, 0, loc)

	dayStart := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, loc)
	timeOfDay := t.Sub(dayStart)

	weekday := dayStart.Weekday()
	daysSinceMonday := (int(weekday) + 6) % 7
	monday := dayStart.AddDate(0, 0, -daysSinceMonday)

	diffDays := int(math.Round(monday.Sub(epoch).Hours() / 24.0))
	fullWeeks := diffDays / 7

	workingDuration := time.Duration(fullWeeks) * (5 * 24 * time.Hour)

	var workDaysBeforeToday int
	switch weekday {
	case time.Monday:
		workDaysBeforeToday = 0
	case time.Tuesday:
		workDaysBeforeToday = 1
	case time.Wednesday:
		workDaysBeforeToday = 2
	case time.Thursday:
		workDaysBeforeToday = 3
	case time.Friday:
		workDaysBeforeToday = 4
	case time.Saturday, time.Sunday:
		workDaysBeforeToday = 5
	}
	workingDuration += time.Duration(workDaysBeforeToday) * 24 * time.Hour

	if weekday != time.Saturday && weekday != time.Sunday {
		workingDuration += timeOfDay
	}

	return workingDuration
}

func (diagram *Diagram) workTime(t time.Time) time.Duration {
	loc := diagram.ComputedStart.Location()
	tInLoc := t.In(loc)
	return workTimeFromEpoch(tInLoc)
}

func (diagram *Diagram) workDurationBetween(start, end time.Time) time.Duration {
	return diagram.workTime(end) - diagram.workTime(start)
}

// DateToX converts a date to the horizontal X coordinate within the diagram's lane area.
// If diagram.HideWeekendsPeriod is true, weekends are excluded from the timescale.
func (diagram *Diagram) DateToX(t time.Time) float64 {
	if diagram.ComputedDuration == 0 {
		return diagram.XLeftLanes
	}
	if diagram.HideWeekendsPeriod {
		totalWorkDuration := diagram.workDurationBetween(diagram.ComputedStart, diagram.ComputedEnd)
		if totalWorkDuration <= 0 {
			return diagram.XLeftLanes
		}
		tWorkDuration := diagram.workDurationBetween(diagram.ComputedStart, t)
		fraction := float64(tWorkDuration) / float64(totalWorkDuration)
		return diagram.XLeftLanes + (diagram.XRightMargin-diagram.XLeftLanes)*fraction
	}
	durationFromStart := t.Sub(diagram.ComputedStart)
	fraction := float64(durationFromStart) / float64(diagram.ComputedDuration)
	return diagram.XLeftLanes + (diagram.XRightMargin-diagram.XLeftLanes)*fraction
}

func (diagram *Diagram) dateToX(t time.Time) float64 {
	return diagram.DateToX(t)
}
