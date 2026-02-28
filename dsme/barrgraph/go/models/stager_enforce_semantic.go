package models

func (stager *Stager) enforceSemantic() {
	// VERY important because the probe only unstages objects
	// the stage might be dirty
	// with slices of pointer or pointer to unstaged instance
	stager.stage.Clean()
	needCommit := false

	// check that there is a Desk
	needCommit = stager.enforce_semantic_singlotons() || needCommit
	needCommit = stager.enforce_semantic_diagrams() || needCommit

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
