package models

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

		if diagram.MovementRectAnchorType == "" {
			diagram.MovementRectAnchorType = RECT_RIGHT
			needCommit = true
		}
		if diagram.MovementTextAnchorType == "" {
			diagram.MovementTextAnchorType = TEXT_ANCHOR_END
			needCommit = true
		}
		if diagram.MovementDateRectAnchorType == "" {
			diagram.MovementDateRectAnchorType = RECT_BOTTOM_LEFT
			needCommit = true
		}
		if diagram.MovementDateTextAnchorType == "" {
			diagram.MovementDateTextAnchorType = TEXT_ANCHOR_START
			needCommit = true
		}
		if diagram.MovementPlacesRectAnchorType == "" {
			diagram.MovementPlacesRectAnchorType = RECT_BOTTOM_RIGHT
			needCommit = true
		}
		if diagram.MovementPlacesTextAnchorType == "" {
			diagram.MovementPlacesTextAnchorType = TEXT_ANCHOR_END
			needCommit = true
		}
	}

	for _, movement := range GetGongstrucsSorted[*Movement](stager.stage) {
		//  movement cannot be minor AND major
		if movement.IsMajor && movement.IsMinor {
			movement.IsMinor = false
			needCommit = true
		}
	}

	// remove orphean movement shapes
	{
		rm := GetSliceOfPointersReverseMap[Diagram, MovementShape](GetAssociationName[Diagram]().MovementShapes[0].Name, stager.stage)
		for _, shape := range GetGongstrucsSorted[*MovementShape](stager.stage) {
			if shape.GetArtElement() == nil {
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
		rm := GetSliceOfPointersReverseMap[Diagram, ArtistShape](GetAssociationName[Diagram]().ArtistShapes[0].Name, stager.stage)
		for _, shape := range GetGongstrucsSorted[*ArtistShape](stager.stage) {
			if shape.GetArtElement() == nil {
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
		rm := GetSliceOfPointersReverseMap[Diagram, ArtefactTypeShape](GetAssociationName[Diagram]().ArtefactTypeShapes[0].Name, stager.stage)
		for _, shape := range GetGongstrucsSorted[*ArtefactTypeShape](stager.stage) {
			if shape.GetArtElement() == nil {
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
			if shape.GetArtElement() == nil {
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
		if influence.SourceMovement != nil {
			influence.source = influence.SourceMovement
		}
		if influence.SourceArtefactType != nil {
			influence.source = influence.SourceArtefactType
		}
		if influence.SourceArtist != nil {
			influence.source = influence.SourceArtist
		}

		if influence.TargetMovement != nil {
			influence.target = influence.TargetMovement
		}
		if influence.TargetArtefactType != nil {
			influence.target = influence.TargetArtefactType
		}
		if influence.TargetArtist != nil {
			influence.target = influence.TargetArtist
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
