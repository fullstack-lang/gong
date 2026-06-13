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

	for _, diagramHierarchy := range GetGongstrucsSorted[*DiagramHierarchy](stager.stage) {
		if diagramHierarchy.DefaultBoxHeigth == 0 {
			diagramHierarchy.DefaultBoxHeigth = defaultBoxHeigth
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Diagram %s: setting default box height to %f", diagramHierarchy.Name, defaultBoxHeigth))
			}
		}
		if diagramHierarchy.DefaultBoxWidth == 0 {
			diagramHierarchy.DefaultBoxWidth = defaultBoxWidth
			needCommit = true
			if stager.probeForm != nil {
				stager.probeForm.AddNotification(time.Now(),
					fmt.Sprintf("Diagram %s: setting default box width to %f", diagramHierarchy.Name, defaultBoxWidth))
			}
		}

		if diagramHierarchy.IsTimeDiagram {
			if diagramHierarchy.LaneHeight == 0 {
				diagramHierarchy.LaneHeight = 100.0
				needCommit = true
			}
			if diagramHierarchy.RatioBarToLaneHeight == 0 {
				diagramHierarchy.RatioBarToLaneHeight = 0.8
				needCommit = true
			}
			if diagramHierarchy.YTopMargin == 0 {
				diagramHierarchy.YTopMargin = 40.0
				needCommit = true
			}
			if diagramHierarchy.XLeftText == 0 {
				diagramHierarchy.XLeftText = 10.0
				needCommit = true
			}
			if diagramHierarchy.TextHeight == 0 {
				diagramHierarchy.TextHeight = 15.0
				needCommit = true
			}
			if diagramHierarchy.XLeftLanes == 0 {
				diagramHierarchy.XLeftLanes = 200.0
				needCommit = true
			}
			if diagramHierarchy.XRightMargin == 0 {
				diagramHierarchy.XRightMargin = 1000.0
				needCommit = true
			}
			if diagramHierarchy.ArrowLengthToTheRightOfStartBar == 0 {
				diagramHierarchy.ArrowLengthToTheRightOfStartBar = 15.0
				needCommit = true
			}
			if diagramHierarchy.ArrowTipLenght == 0 {
				diagramHierarchy.ArrowTipLenght = 5.0
				needCommit = true
			}
			if diagramHierarchy.TimeLine_Color == "" {
				diagramHierarchy.TimeLine_Color = "grey"
				needCommit = true
			}
			if diagramHierarchy.TimeLine_FillOpacity == 0 {
				diagramHierarchy.TimeLine_FillOpacity = 0.1
				needCommit = true
			}
			if diagramHierarchy.TimeLine_Stroke == "" {
				diagramHierarchy.TimeLine_Stroke = "grey"
				needCommit = true
			}
			if diagramHierarchy.TimeLine_StrokeWidth == 0 {
				diagramHierarchy.TimeLine_StrokeWidth = 1.0
				needCommit = true
			}
			if diagramHierarchy.Group_Stroke == "" {
				diagramHierarchy.Group_Stroke = "black"
				needCommit = true
			}
			if diagramHierarchy.Group_StrokeWidth == 0 {
				diagramHierarchy.Group_StrokeWidth = 1.0
				needCommit = true
			}
			if diagramHierarchy.Group_StrokeDashArray == "" {
				diagramHierarchy.Group_StrokeDashArray = "2 2"
				needCommit = true
			}
			if diagramHierarchy.DateYOffset == 0 {
				diagramHierarchy.DateYOffset = 15.0
				needCommit = true
			}
		}
	}
	return
}
