package models

import (
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

func (stager *Stager) ComputeConsistency() {

	// VERY important because the probe only unstages objects
	// the stage might be dirty
	// with slices of pointer or pointer to unstaged instance
	stager.stage.Clean()
	needCommit := false

	// check that there is a Desk
	if len(GetGongstrucsSorted[*Desk](stager.stage)) == 0 {
		(&Desk{Name: "Desk"}).Stage(stager.stage)
		needCommit = true
	}

	stager.desk = GetGongstrucsSorted[*Desk](stager.stage)[0]

	// at least one diagram is welcome to ease the end user experience
	if len(GetGongstrucsSorted[*Diagram](stager.stage)) == 0 {
		diagram := (&Diagram{
			Name: "Default"}).Stage(stager.stage)
		stager.desk.SelectedDiagram = diagram
		needCommit = true
	}

	for _, diagram := range GetGongstrucsSorted[*Diagram](stager.stage) {

		if diagram.NbPixPerCharacter == 0 {
			diagram.NbPixPerCharacter = 8
			needCommit = true
		}

		if diagram.Width == 0 {
			diagram.Width = 800
			needCommit = true
		}
		if diagram.Height == 0 {
			diagram.Height = 2000
			needCommit = true
		}

		if diagram.GrayColorCode == "" {
			diagram.GrayColorCode = svg.Gray.ToString()
			needCommit = true
		}
		if diagram.RedColorCode == "" {
			diagram.RedColorCode = svg.Salmon.ToString()
			needCommit = true
		}
		if diagram.BackgroundGreyColorCode == "" {
			diagram.BackgroundGreyColorCode = svg.White.ToString()
			needCommit = true
		}
		if diagram.Category1RectAnchorType == "" {
			diagram.Category1RectAnchorType = RECT_RIGHT
			needCommit = true
		}
		if diagram.Category1TextAnchorType == "" {
			diagram.Category1TextAnchorType = TEXT_ANCHOR_END
			needCommit = true
		}

		if diagram.InfluenceStrokeWidth == 0 {
			diagram.InfluenceStrokeWidth = 1
			needCommit = true
		}
		if diagram.InfluenceArrowSize == 0 {
			diagram.InfluenceArrowSize = 10
			needCommit = true
		}
		if diagram.InfluenceCornerRadius == 0 {
			diagram.InfluenceCornerRadius = 8
			needCommit = true
		}
	}

	for _, movement := range GetGongstrucsSorted[*Category1](stager.stage) {
		_ = movement
	}

	// remove orphean movement shapes
	{
		rm := GetSliceOfPointersReverseMap[Diagram, Category1Shape](GetAssociationName[Diagram]().Category1Shapes[0].Name, stager.stage)
		for _, shape := range GetGongstrucsSorted[*Category1Shape](stager.stage) {
			if shape.GetCategory() == nil {
				shape.Unstage(stager.stage)
				needCommit = true
			}
			if _, ok := rm[shape]; !ok {
				shape.Unstage(stager.stage)
				needCommit = true
			}
		}
	}
	{
		rm := GetSliceOfPointersReverseMap[Diagram, Category2Shape](GetAssociationName[Diagram]().Category2Shapes[0].Name, stager.stage)
		for _, shape := range GetGongstrucsSorted[*Category2Shape](stager.stage) {
			if shape.GetCategory() == nil {
				shape.Unstage(stager.stage)
				needCommit = true
			}
			if _, ok := rm[shape]; !ok {
				shape.Unstage(stager.stage)
				needCommit = true
			}
		}
	}
	{
		rm := GetSliceOfPointersReverseMap[Diagram, Category3Shape](GetAssociationName[Diagram]().Category3Shapes[0].Name, stager.stage)
		for _, shape := range GetGongstrucsSorted[*Category3Shape](stager.stage) {
			if shape.GetCategory() == nil {
				shape.Unstage(stager.stage)
				needCommit = true
			}
			if _, ok := rm[shape]; !ok {
				shape.Unstage(stager.stage)
				needCommit = true
			}
		}
	}
	{
		rm := GetSliceOfPointersReverseMap[Diagram, InfluenceShape](GetAssociationName[Diagram]().InfluenceShapes[0].Name, stager.stage)
		for _, shape := range GetGongstrucsSorted[*InfluenceShape](stager.stage) {
			if shape.GetCategory() == nil {
				shape.Unstage(stager.stage)
				needCommit = true
			}
			if _, ok := rm[shape]; !ok {
				shape.Unstage(stager.stage)
				needCommit = true
			}
		}
	}

	{
		rm := GetSliceOfPointersReverseMap[InfluenceShape, ControlPointShape](GetAssociationName[InfluenceShape]().ControlPointShapes[0].Name, stager.stage)
		for _, shape := range GetGongstrucsSorted[*ControlPointShape](stager.stage) {
			if _, ok := rm[shape]; !ok {
				shape.Unstage(stager.stage)
				needCommit = true
			}
		}
	}

	for _, influence := range GetGongstrucsSorted[*Influence](stager.stage) {
		if influence.SourceCategory1 != nil {
			influence.source = influence.SourceCategory1
		}
		if influence.SourceCategory2 != nil {
			influence.source = influence.SourceCategory2
		}
		if influence.SourceCategory3 != nil {
			influence.source = influence.SourceCategory3
		}

		if influence.TargetCategory1 != nil {
			influence.target = influence.TargetCategory1
		}
		if influence.TargetCategory2 != nil {
			influence.target = influence.TargetCategory2
		}
		if influence.TargetCategory3 != nil {
			influence.target = influence.TargetCategory3
		}

		if influence.Name != influence.source.GetName()+" to "+influence.target.GetName() {
			influence.Name = influence.source.GetName() + " to " + influence.target.GetName()
			needCommit = true
		}
	}

	for _, influenceShape := range GetGongstrucsSorted[*InfluenceShape](stager.stage) {
		if influenceShape.Influence == nil {
			influenceShape.Unstage(stager.stage)
			needCommit = true
			continue
		}
		if influenceShape.Name != influenceShape.Influence.Name {
			influenceShape.Name = influenceShape.Influence.Name
			needCommit = true
		}
	}

	if needCommit {
		stager.stage.CommitWithSuspendedCallbacks()
		stager.stage.Checkout()
	}
}
