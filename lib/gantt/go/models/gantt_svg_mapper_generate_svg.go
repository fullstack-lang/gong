package models

import (
	"fmt"
	"log"
	"math"
	"sort"
	"time"

	gongsvg_models "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (ganttSVGMapper *GanttSVGMapper) GenerateSvg(
	gongganttStage *Stage,
	gongsvgStage *gongsvg_models.Stage) {

	gongsvgStage.Reset()

	if len(*GetGongstructInstancesSet[Gantt](gongganttStage)) != 1 {
		log.Printf("It is supposed to have only one gantt chart")
		return
	}

	var gantt *Gantt
	for gantt_ := range *GetGongstructInstancesSet[Gantt](gongganttStage) {
		gantt = gantt_
	}
	gantt.ComputeStartAndEndDate()

	//
	// SVG
	//
	svg := new(gongsvg_models.SVG).Stage(gongsvgStage)
	svg.Name = "SVG"
	svg.IsEditable = true

	layerLanes := new(gongsvg_models.Layer).Stage(gongsvgStage)
	layerLanes.Name = "Lanes layer"
	svg.Layers = append(svg.Layers, layerLanes)

	layerBars := new(gongsvg_models.Layer).Stage(gongsvgStage)
	layerBars.Name = "Bars layer"
	svg.Layers = append(svg.Layers, layerBars)

	//
	// Time Line
	//
	// creates a tick lane
	LaneHeight := gantt.LaneHeight
	RatioBarToLaneHeight := gantt.RatioBarToLaneHeight
	barHeigth := LaneHeight * RatioBarToLaneHeight
	YTopMargin := gantt.YTopMargin
	yTimeLine := LaneHeight*float64(len(*GetGongstructInstancesSet[Lane](gongganttStage))) + YTopMargin

	XLeftText := gantt.XLeftText
	TextHeight := gantt.TextHeight

	XLeftLanes := gantt.XLeftLanes
	XRightMargin := gantt.XRightMargin

	timeLine := new(gongsvg_models.Line).Stage(gongsvgStage)
	timeLine.Name = "Time Line"
	timeLine.X1 = XLeftLanes
	timeLine.Y1 = yTimeLine
	timeLine.X2 = XRightMargin
	timeLine.Y2 = yTimeLine

	timeLine.Color = gantt.TimeLine_Color
	timeLine.FillOpacity = gantt.TimeLine_FillOpacity
	timeLine.Stroke = gantt.TimeLine_Stroke
	timeLine.StrokeWidth = gantt.TimeLine_StrokeWidth

	layerLanes.Lines = append(layerLanes.Lines, timeLine)

	//
	// tick lines
	//

	//
	// Dates
	//

	// Date offset. Where is the date below the TimeLine
	DateYOffset := 20.0

	// put a date for every year
	for year := gantt.ComputedStart.Year(); year <= gantt.ComputedEnd.Year(); year++ {
		yearTime := time.Date(year, time.January, 1, 0, 0, 0, 0, time.UTC)
		durationBetweenYearAndGanttStart := yearTime.Sub(gantt.ComputedStart)

		durationBetweenYearStartAndGanttStartRelativeToGanttDuration :=
			float64(durationBetweenYearAndGanttStart) / float64(gantt.ComputedEnd.Sub(gantt.ComputedStart))

		yearText := new(gongsvg_models.Text).Stage(gongsvgStage)
		yearText.Name = fmt.Sprintf("%d", year)
		yearText.Content = yearText.Name
		yearText.X = XLeftLanes + (XRightMargin-XLeftLanes)*durationBetweenYearStartAndGanttStartRelativeToGanttDuration
		yearText.Y = yTimeLine + DateYOffset
		yearText.Color = "black"
		yearText.FillOpacity = 1.0
		layerLanes.Texts = append(layerLanes.Texts, yearText)

		// log.Printf("year %d", year)
		//
		// draw the line
		//
		lineForYear := new(gongsvg_models.Line).Stage(gongsvgStage)
		lineForYear.Name = yearText.Name
		layerLanes.Lines = append(layerLanes.Lines, lineForYear)
		lineForYear.X1 = yearText.X
		lineForYear.X2 = lineForYear.X1
		lineForYear.Y1 = YTopMargin
		lineForYear.Y2 = yTimeLine
		lineForYear.Stroke = "black"
		lineForYear.StrokeWidth = 0.25
		lineForYear.StrokeDashArray = "4 4"
	}

	//
	// Lanes
	//
	currentY := YTopMargin

	mapLane_TextY := make(map[*Lane]float64, 0)
	mapLane_TopY := make(map[*Lane]float64, 0)

	//
	// Sort Lanes according to the Order
	//
	sort.Slice(gantt.Lanes, func(i, j int) bool {
		return gantt.Lanes[i].Order < gantt.Lanes[j].Order
	})

	laneIndex := 0

	// prepare a map of bar to barSVG
	mapBar_Rect := make(map[*Bar]*gongsvg_models.Rect)
	mapRect4Bar_Bar := make(map[*gongsvg_models.Rect]*Bar)

	for _, lane := range gantt.Lanes {

		laneSVG := new(gongsvg_models.Rect).Stage(gongsvgStage)
		layerLanes.Rects = append(layerLanes.Rects, laneSVG)
		laneSVG.Name = lane.Name
		laneSVG.X = XLeftLanes
		laneSVG.Y = currentY
		mapLane_TopY[lane] = currentY

		laneSVG.Width = XRightMargin - XLeftLanes
		laneSVG.Height = LaneHeight

		laneSVG.Color = "grey"
		laneSVG.FillOpacity = 0.1

		if laneIndex%2 == 0 {
			laneSVG.Color = "black"
		}
		laneIndex = laneIndex + 1
		laneSVG.StrokeWidth = 1.5

		laneText := new(gongsvg_models.Text).Stage(gongsvgStage)
		laneText.Name = lane.Name
		laneText.Content = laneText.Name
		laneText.X = XLeftText
		laneText.Y = currentY + LaneHeight/2.0 + TextHeight/2.0
		mapLane_TextY[lane] = laneText.Y
		laneText.Color = "black"
		laneText.FillOpacity = 1.0
		layerLanes.Texts = append(layerLanes.Texts, laneText)

		//
		// Bar
		//
		for _, bar := range lane.Bars {
			rect4Bar := new(gongsvg_models.Rect).Stage(gongsvgStage)
			mapBar_Rect[bar] = rect4Bar
			mapRect4Bar_Bar[rect4Bar] = bar
			layerBars.Rects = append(layerBars.Rects, rect4Bar)
			rect4Bar.Name = bar.Name
			rect4Bar.IsSelectable = true
			rect4Bar.CanHaveRightHandle = true
			rect4Bar.CanHaveLeftHandle = true
			rect4Bar.CanMoveHorizontaly = false

			// connects the call on rect to the bar
			barImpl := new(BarImpl)
			barImpl.Bar = bar
			barImpl.GanttToRender = gantt
			barImpl.GongganttStage = gongganttStage
			rect4Bar.Impl = barImpl

			var barToDisplay = *bar

			// if start and end dates of the gantt are set manualy, then
			// reset start and end dates of the bar to display
			if gantt.UseManualStartAndEndDates {
				if bar.Start.Before(gantt.ManualStart) {
					barToDisplay.Start = gantt.ManualStart
				}
				if bar.End.After(gantt.ManualEnd) {
					barToDisplay.End = gantt.ManualEnd
				}
			}

			durationFromGanttStartToBarStart := barToDisplay.Start.Sub(gantt.ComputedStart)
			durationBetweenBarStartAndGanttStartRelativeToGanttDuration :=
				float64(durationFromGanttStartToBarStart) / float64(gantt.ComputedEnd.Sub(gantt.ComputedStart))
			// log.Printf("Duration is %f", durationBetweenBarStartAndGanttStartRelativeToGanttDuration)

			durationFromBarEndAndBarStart := barToDisplay.End.Sub(barToDisplay.Start)
			durationBetweenBarEndAndBarStartRelativeToGanttDuration :=
				float64(durationFromBarEndAndBarStart) / float64(gantt.ComputedEnd.Sub(gantt.ComputedStart))
			// log.Printf("Relative Duration is %f", durationBetweenBarEndAndBarStartRelativeToGanttDuration)

			rect4Bar.X = XLeftLanes + (XRightMargin-XLeftLanes)*durationBetweenBarStartAndGanttStartRelativeToGanttDuration
			rect4Bar.Y = currentY + (LaneHeight-barHeigth)/2.0
			rect4Bar.Height = barHeigth
			rect4Bar.Width = (XRightMargin - XLeftLanes) * durationBetweenBarEndAndBarStartRelativeToGanttDuration

			rect4Bar.Color = "blue"
			if bar.OptionnalColor != "" {
				rect4Bar.Color = bar.OptionnalColor
			}
			rect4Bar.FillOpacity = 0.1
			if bar.FillOpacity != 0.0 {
				rect4Bar.FillOpacity = bar.FillOpacity
			}
			rect4Bar.Stroke = "blue"
			if bar.OptionnalStroke != "" {
				rect4Bar.Stroke = bar.OptionnalStroke
			}
			rect4Bar.StrokeWidth = 0.5
			if bar.StrokeWidth != 0.0 {
				rect4Bar.StrokeWidth = bar.StrokeWidth
			}
			if bar.StrokeDashArray != "" {
				rect4Bar.StrokeDashArray = bar.StrokeDashArray
			}

			// bar text
			barText := new(gongsvg_models.Text).Stage(gongsvgStage)
			barText.Name = bar.Name
			barText.Content = barText.Name
			barText.X = rect4Bar.X + XLeftText
			barText.Y = currentY + LaneHeight/2.0 + TextHeight/2.0
			barText.Color = "black"
			barText.FillOpacity = 1.0
			layerBars.Texts = append(layerBars.Texts, barText)
		}

		currentY = currentY + LaneHeight

	}

	//
	// Milestones
	//
	for _, milestone := range gantt.Milestones {

		if gantt.UseManualStartAndEndDates &&
			(milestone.Date.Before(gantt.ManualStart) ||
				milestone.Date.After(gantt.ManualEnd)) {
			continue
		}
		durationBetweenMilestoneAndGanttStart := milestone.Date.Sub(gantt.ComputedStart)

		durationBetweenMilestoneAndGanttStartString := durationBetweenMilestoneAndGanttStart.String()
		_ = durationBetweenMilestoneAndGanttStartString

		durationBetweenMilestoneAndGanttStartRelativeToGanttDuration :=
			float64(durationBetweenMilestoneAndGanttStart) / float64(gantt.ComputedEnd.Sub(gantt.ComputedStart))

		//
		// draw the line
		//
		lineX := XLeftLanes + (XRightMargin-XLeftLanes)*durationBetweenMilestoneAndGanttStartRelativeToGanttDuration
		if milestone.DisplayVerticalBar {
			line := new(gongsvg_models.Line).Stage(gongsvgStage)
			line.Name = milestone.Name
			layerLanes.Lines = append(layerLanes.Lines, line)
			line.X1 = lineX
			line.X2 = line.X1
			line.Y1 = YTopMargin
			line.Y2 = yTimeLine
			line.Stroke = "black"
			line.StrokeWidth = 0.5
			line.StrokeDashArray = "2 2"
		}

		//
		// draw diamond
		//
		diamondWidth := 18.0
		for _, laneToDisplay := range milestone.LanesToDisplay {

			diamond := new(gongsvg_models.Rect).Stage(gongsvgStage)
			layerLanes.Rects = append(layerLanes.Rects, diamond)
			diamond.Name = milestone.Name
			diamond.X = lineX - diamondWidth/2.0
			diamond.Y = mapLane_TextY[laneToDisplay] - diamondWidth + LaneHeight/2.0
			diamond.Width = diamondWidth
			diamond.Height = diamondWidth
			diamond.Color = "red"
			diamond.FillOpacity = 0.4
			diamond.Transform = fmt.Sprintf("rotate(%d %d %d)", 45, int64(diamond.X+diamondWidth/2.0), int64(diamond.Y+diamondWidth/2.0))

			// bar text
			milestoneText := new(gongsvg_models.Text).Stage(gongsvgStage)
			milestoneText.Name = milestone.Name
			milestoneText.Content = milestoneText.Name
			milestoneText.X = diamond.X + XLeftText + diamondWidth
			milestoneText.Y = diamond.Y + TextHeight/2.0 + 5 // manual fine tuning
			milestoneText.Color = "black"
			milestoneText.FillOpacity = 1.0
			layerLanes.Texts = append(layerLanes.Texts, milestoneText)
		}
	}

	//
	// Arrows, for each arrow, draw five lines,
	//
	// start : middle of the end of the "Start" bar on
	//
	for _, arrow := range gantt.Arrows {

		startBar := mapBar_Rect[arrow.From]
		endBar := mapBar_Rect[arrow.To]

		generate_arrow(
			gongsvgStage,
			layerLanes,
			startBar.X+startBar.Width,
			endBar.X,
			startBar.Y+barHeigth/2.0,
			endBar.Y+barHeigth/2.0,
			gantt.ArrowLengthToTheRightOfStartBar,
			gantt.ArrowTipLenght,
			arrow.Name,
			arrow.OptionnalColor,
			arrow.OptionnalStroke)
	}

	//
	// Groups of Lanes
	//
	for _, group := range gantt.Groups {

		if len(group.GroupLanes) == 0 {
			continue
		}

		groupSVG := new(gongsvg_models.Rect).Stage(gongsvgStage)
		layerLanes.Rects = append(layerLanes.Rects, groupSVG)
		groupSVG.Name = group.Name

		// compute X from list of lane
		groupTopY := math.MaxFloat64
		groupBottomY := 0.0
		for _, lane := range group.GroupLanes {
			if groupTopY > mapLane_TopY[lane] {
				groupTopY = mapLane_TopY[lane]
			}
			if groupBottomY < (mapLane_TopY[lane] + gantt.LaneHeight) {
				groupBottomY = mapLane_TopY[lane] + gantt.LaneHeight
			}
		}

		groupSVG.X = 0
		groupSVG.Y = groupTopY
		groupSVG.Width = gantt.XRightMargin - groupSVG.X
		groupSVG.Height = groupBottomY - groupTopY

		groupSVG.Stroke = gantt.Group_Stroke
		groupSVG.StrokeWidth = gantt.Group_StrokeWidth
		groupSVG.StrokeDashArray = gantt.Group_StrokeDashArray

		// text
		groupText := new(gongsvg_models.Text).Stage(gongsvgStage)
		groupText.Name = group.Name
		groupText.Content = groupText.Name
		groupText.X = XLeftText
		groupText.Y = groupTopY + gantt.DateYOffset
		groupText.Color = "blue"
		groupText.FillOpacity = 0.5
		layerLanes.Texts = append(layerLanes.Texts, groupText)
	}

	gongsvgStage.Commit()

	// log.Printf("Before Commit")
}
