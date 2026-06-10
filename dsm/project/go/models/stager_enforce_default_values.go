package models

import (
	"fmt"
	"time"
)

// enforceDefaultValues enforce defaut values when there are not suitable
func (stager *Stager) enforceDefaultValues() (needCommit bool) {
	const (
		defaultBoxWidth  = 250.0
		defaultBoxHeigth = 70.0
	)

	root := stager.getRootLibrary()
	if root.NbPixPerCharacter == 0 {
		root.NbPixPerCharacter = 8
		needCommit = true

		stager.probeForm.AddNotification(time.Now(),
			fmt.Sprintf("Root: setting nbPixPerCharacter to %f", root.NbPixPerCharacter))
	}

	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {
		if diagram.DefaultBoxHeigth == 0 {
			diagram.DefaultBoxHeigth = defaultBoxHeigth
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Diagram %s: setting default box height to %f", diagram.Name, defaultBoxHeigth))
			}
		}
		if diagram.DefaultBoxWidth == 0 {
			diagram.DefaultBoxWidth = defaultBoxWidth
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Diagram %s: setting default box width to %f", diagram.Name, defaultBoxWidth))
			}
		}

		if diagram.IsTimeDiagram {
			if diagram.LaneHeight == 0 {
				diagram.LaneHeight = 30.0
				needCommit = true
			}
			if diagram.RatioBarToLaneHeight == 0 {
				diagram.RatioBarToLaneHeight = 0.8
				needCommit = true
			}
			if diagram.YTopMargin == 0 {
				diagram.YTopMargin = 40.0
				needCommit = true
			}
			if diagram.XLeftText == 0 {
				diagram.XLeftText = 10.0
				needCommit = true
			}
			if diagram.TextHeight == 0 {
				diagram.TextHeight = 15.0
				needCommit = true
			}
			if diagram.XLeftLanes == 0 {
				diagram.XLeftLanes = 200.0
				needCommit = true
			}
			if diagram.XRightMargin == 0 {
				diagram.XRightMargin = 50.0
				needCommit = true
			}
			if diagram.ArrowLengthToTheRightOfStartBar == 0 {
				diagram.ArrowLengthToTheRightOfStartBar = 15.0
				needCommit = true
			}
			if diagram.ArrowTipLenght == 0 {
				diagram.ArrowTipLenght = 5.0
				needCommit = true
			}
			if diagram.TimeLine_Color == "" {
				diagram.TimeLine_Color = "grey"
				needCommit = true
			}
			if diagram.TimeLine_FillOpacity == 0 {
				diagram.TimeLine_FillOpacity = 0.1
				needCommit = true
			}
			if diagram.TimeLine_Stroke == "" {
				diagram.TimeLine_Stroke = "grey"
				needCommit = true
			}
			if diagram.TimeLine_StrokeWidth == 0 {
				diagram.TimeLine_StrokeWidth = 1.0
				needCommit = true
			}
			if diagram.Group_Stroke == "" {
				diagram.Group_Stroke = "black"
				needCommit = true
			}
			if diagram.Group_StrokeWidth == 0 {
				diagram.Group_StrokeWidth = 1.0
				needCommit = true
			}
			if diagram.Group_StrokeDashArray == "" {
				diagram.Group_StrokeDashArray = "2 2"
				needCommit = true
			}
			if diagram.DateYOffset == 0 {
				diagram.DateYOffset = 15.0
				needCommit = true
			}
		}
	}
	return
}
