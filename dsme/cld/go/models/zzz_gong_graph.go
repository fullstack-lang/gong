// generated code - do not edit
package models

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *ArtefactType:
		ok = stage.IsStagedArtefactType(target)

	case *ArtefactTypeShape:
		ok = stage.IsStagedArtefactTypeShape(target)

	case *Artist:
		ok = stage.IsStagedArtist(target)

	case *ArtistShape:
		ok = stage.IsStagedArtistShape(target)

	case *ControlPointShape:
		ok = stage.IsStagedControlPointShape(target)

	case *Desk:
		ok = stage.IsStagedDesk(target)

	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Influence:
		ok = stage.IsStagedInfluence(target)

	case *InfluenceShape:
		ok = stage.IsStagedInfluenceShape(target)

	case *Movement:
		ok = stage.IsStagedMovement(target)

	case *MovementShape:
		ok = stage.IsStagedMovementShape(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *ArtefactType:
		ok = stage.IsStagedArtefactType(target)

	case *ArtefactTypeShape:
		ok = stage.IsStagedArtefactTypeShape(target)

	case *Artist:
		ok = stage.IsStagedArtist(target)

	case *ArtistShape:
		ok = stage.IsStagedArtistShape(target)

	case *ControlPointShape:
		ok = stage.IsStagedControlPointShape(target)

	case *Desk:
		ok = stage.IsStagedDesk(target)

	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Influence:
		ok = stage.IsStagedInfluence(target)

	case *InfluenceShape:
		ok = stage.IsStagedInfluenceShape(target)

	case *Movement:
		ok = stage.IsStagedMovement(target)

	case *MovementShape:
		ok = stage.IsStagedMovementShape(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedArtefactType(artefacttype *ArtefactType) (ok bool) {

	_, ok = stage.ArtefactTypes[artefacttype]

	return
}

func (stage *Stage) IsStagedArtefactTypeShape(artefacttypeshape *ArtefactTypeShape) (ok bool) {

	_, ok = stage.ArtefactTypeShapes[artefacttypeshape]

	return
}

func (stage *Stage) IsStagedArtist(artist *Artist) (ok bool) {

	_, ok = stage.Artists[artist]

	return
}

func (stage *Stage) IsStagedArtistShape(artistshape *ArtistShape) (ok bool) {

	_, ok = stage.ArtistShapes[artistshape]

	return
}

func (stage *Stage) IsStagedControlPointShape(controlpointshape *ControlPointShape) (ok bool) {

	_, ok = stage.ControlPointShapes[controlpointshape]

	return
}

func (stage *Stage) IsStagedDesk(desk *Desk) (ok bool) {

	_, ok = stage.Desks[desk]

	return
}

func (stage *Stage) IsStagedDiagram(diagram *Diagram) (ok bool) {

	_, ok = stage.Diagrams[diagram]

	return
}

func (stage *Stage) IsStagedInfluence(influence *Influence) (ok bool) {

	_, ok = stage.Influences[influence]

	return
}

func (stage *Stage) IsStagedInfluenceShape(influenceshape *InfluenceShape) (ok bool) {

	_, ok = stage.InfluenceShapes[influenceshape]

	return
}

func (stage *Stage) IsStagedMovement(movement *Movement) (ok bool) {

	_, ok = stage.Movements[movement]

	return
}

func (stage *Stage) IsStagedMovementShape(movementshape *MovementShape) (ok bool) {

	_, ok = stage.MovementShapes[movementshape]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *ArtefactType:
		stage.StageBranchArtefactType(target)

	case *ArtefactTypeShape:
		stage.StageBranchArtefactTypeShape(target)

	case *Artist:
		stage.StageBranchArtist(target)

	case *ArtistShape:
		stage.StageBranchArtistShape(target)

	case *ControlPointShape:
		stage.StageBranchControlPointShape(target)

	case *Desk:
		stage.StageBranchDesk(target)

	case *Diagram:
		stage.StageBranchDiagram(target)

	case *Influence:
		stage.StageBranchInfluence(target)

	case *InfluenceShape:
		stage.StageBranchInfluenceShape(target)

	case *Movement:
		stage.StageBranchMovement(target)

	case *MovementShape:
		stage.StageBranchMovementShape(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchArtefactType(artefacttype *ArtefactType) {

	// check if instance is already staged
	if IsStaged(stage, artefacttype) {
		return
	}

	artefacttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchArtefactTypeShape(artefacttypeshape *ArtefactTypeShape) {

	// check if instance is already staged
	if IsStaged(stage, artefacttypeshape) {
		return
	}

	artefacttypeshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if artefacttypeshape.ArtefactType != nil {
		StageBranch(stage, artefacttypeshape.ArtefactType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchArtist(artist *Artist) {

	// check if instance is already staged
	if IsStaged(stage, artist) {
		return
	}

	artist.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchArtistShape(artistshape *ArtistShape) {

	// check if instance is already staged
	if IsStaged(stage, artistshape) {
		return
	}

	artistshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if artistshape.Artist != nil {
		StageBranch(stage, artistshape.Artist)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchControlPointShape(controlpointshape *ControlPointShape) {

	// check if instance is already staged
	if IsStaged(stage, controlpointshape) {
		return
	}

	controlpointshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDesk(desk *Desk) {

	// check if instance is already staged
	if IsStaged(stage, desk) {
		return
	}

	desk.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if desk.SelectedDiagram != nil {
		StageBranch(stage, desk.SelectedDiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if IsStaged(stage, diagram) {
		return
	}

	diagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _movementshape := range diagram.MovementShapes {
		StageBranch(stage, _movementshape)
	}
	for _, _artefacttypeshape := range diagram.ArtefactTypeShapes {
		StageBranch(stage, _artefacttypeshape)
	}
	for _, _artistshape := range diagram.ArtistShapes {
		StageBranch(stage, _artistshape)
	}
	for _, _influenceshape := range diagram.InfluenceShapes {
		StageBranch(stage, _influenceshape)
	}

}

func (stage *Stage) StageBranchInfluence(influence *Influence) {

	// check if instance is already staged
	if IsStaged(stage, influence) {
		return
	}

	influence.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if influence.SourceMovement != nil {
		StageBranch(stage, influence.SourceMovement)
	}
	if influence.SourceArtefactType != nil {
		StageBranch(stage, influence.SourceArtefactType)
	}
	if influence.SourceArtist != nil {
		StageBranch(stage, influence.SourceArtist)
	}
	if influence.TargetMovement != nil {
		StageBranch(stage, influence.TargetMovement)
	}
	if influence.TargetArtefactType != nil {
		StageBranch(stage, influence.TargetArtefactType)
	}
	if influence.TargetArtist != nil {
		StageBranch(stage, influence.TargetArtist)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchInfluenceShape(influenceshape *InfluenceShape) {

	// check if instance is already staged
	if IsStaged(stage, influenceshape) {
		return
	}

	influenceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if influenceshape.Influence != nil {
		StageBranch(stage, influenceshape.Influence)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range influenceshape.ControlPointShapes {
		StageBranch(stage, _controlpointshape)
	}

}

func (stage *Stage) StageBranchMovement(movement *Movement) {

	// check if instance is already staged
	if IsStaged(stage, movement) {
		return
	}

	movement.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchMovementShape(movementshape *MovementShape) {

	// check if instance is already staged
	if IsStaged(stage, movementshape) {
		return
	}

	movementshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if movementshape.Movement != nil {
		StageBranch(stage, movementshape.Movement)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// CopyBranch stages instance and apply CopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func CopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *ArtefactType:
		toT := CopyBranchArtefactType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ArtefactTypeShape:
		toT := CopyBranchArtefactTypeShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Artist:
		toT := CopyBranchArtist(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ArtistShape:
		toT := CopyBranchArtistShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ControlPointShape:
		toT := CopyBranchControlPointShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Desk:
		toT := CopyBranchDesk(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Diagram:
		toT := CopyBranchDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Influence:
		toT := CopyBranchInfluence(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *InfluenceShape:
		toT := CopyBranchInfluenceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Movement:
		toT := CopyBranchMovement(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MovementShape:
		toT := CopyBranchMovementShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchArtefactType(mapOrigCopy map[any]any, artefacttypeFrom *ArtefactType) (artefacttypeTo *ArtefactType) {

	// artefacttypeFrom has already been copied
	if _artefacttypeTo, ok := mapOrigCopy[artefacttypeFrom]; ok {
		artefacttypeTo = _artefacttypeTo.(*ArtefactType)
		return
	}

	artefacttypeTo = new(ArtefactType)
	mapOrigCopy[artefacttypeFrom] = artefacttypeTo
	artefacttypeFrom.CopyBasicFields(artefacttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchArtefactTypeShape(mapOrigCopy map[any]any, artefacttypeshapeFrom *ArtefactTypeShape) (artefacttypeshapeTo *ArtefactTypeShape) {

	// artefacttypeshapeFrom has already been copied
	if _artefacttypeshapeTo, ok := mapOrigCopy[artefacttypeshapeFrom]; ok {
		artefacttypeshapeTo = _artefacttypeshapeTo.(*ArtefactTypeShape)
		return
	}

	artefacttypeshapeTo = new(ArtefactTypeShape)
	mapOrigCopy[artefacttypeshapeFrom] = artefacttypeshapeTo
	artefacttypeshapeFrom.CopyBasicFields(artefacttypeshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if artefacttypeshapeFrom.ArtefactType != nil {
		artefacttypeshapeTo.ArtefactType = CopyBranchArtefactType(mapOrigCopy, artefacttypeshapeFrom.ArtefactType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchArtist(mapOrigCopy map[any]any, artistFrom *Artist) (artistTo *Artist) {

	// artistFrom has already been copied
	if _artistTo, ok := mapOrigCopy[artistFrom]; ok {
		artistTo = _artistTo.(*Artist)
		return
	}

	artistTo = new(Artist)
	mapOrigCopy[artistFrom] = artistTo
	artistFrom.CopyBasicFields(artistTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchArtistShape(mapOrigCopy map[any]any, artistshapeFrom *ArtistShape) (artistshapeTo *ArtistShape) {

	// artistshapeFrom has already been copied
	if _artistshapeTo, ok := mapOrigCopy[artistshapeFrom]; ok {
		artistshapeTo = _artistshapeTo.(*ArtistShape)
		return
	}

	artistshapeTo = new(ArtistShape)
	mapOrigCopy[artistshapeFrom] = artistshapeTo
	artistshapeFrom.CopyBasicFields(artistshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if artistshapeFrom.Artist != nil {
		artistshapeTo.Artist = CopyBranchArtist(mapOrigCopy, artistshapeFrom.Artist)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchControlPointShape(mapOrigCopy map[any]any, controlpointshapeFrom *ControlPointShape) (controlpointshapeTo *ControlPointShape) {

	// controlpointshapeFrom has already been copied
	if _controlpointshapeTo, ok := mapOrigCopy[controlpointshapeFrom]; ok {
		controlpointshapeTo = _controlpointshapeTo.(*ControlPointShape)
		return
	}

	controlpointshapeTo = new(ControlPointShape)
	mapOrigCopy[controlpointshapeFrom] = controlpointshapeTo
	controlpointshapeFrom.CopyBasicFields(controlpointshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDesk(mapOrigCopy map[any]any, deskFrom *Desk) (deskTo *Desk) {

	// deskFrom has already been copied
	if _deskTo, ok := mapOrigCopy[deskFrom]; ok {
		deskTo = _deskTo.(*Desk)
		return
	}

	deskTo = new(Desk)
	mapOrigCopy[deskFrom] = deskTo
	deskFrom.CopyBasicFields(deskTo)

	//insertion point for the staging of instances referenced by pointers
	if deskFrom.SelectedDiagram != nil {
		deskTo.SelectedDiagram = CopyBranchDiagram(mapOrigCopy, deskFrom.SelectedDiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDiagram(mapOrigCopy map[any]any, diagramFrom *Diagram) (diagramTo *Diagram) {

	// diagramFrom has already been copied
	if _diagramTo, ok := mapOrigCopy[diagramFrom]; ok {
		diagramTo = _diagramTo.(*Diagram)
		return
	}

	diagramTo = new(Diagram)
	mapOrigCopy[diagramFrom] = diagramTo
	diagramFrom.CopyBasicFields(diagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _movementshape := range diagramFrom.MovementShapes {
		diagramTo.MovementShapes = append(diagramTo.MovementShapes, CopyBranchMovementShape(mapOrigCopy, _movementshape))
	}
	for _, _artefacttypeshape := range diagramFrom.ArtefactTypeShapes {
		diagramTo.ArtefactTypeShapes = append(diagramTo.ArtefactTypeShapes, CopyBranchArtefactTypeShape(mapOrigCopy, _artefacttypeshape))
	}
	for _, _artistshape := range diagramFrom.ArtistShapes {
		diagramTo.ArtistShapes = append(diagramTo.ArtistShapes, CopyBranchArtistShape(mapOrigCopy, _artistshape))
	}
	for _, _influenceshape := range diagramFrom.InfluenceShapes {
		diagramTo.InfluenceShapes = append(diagramTo.InfluenceShapes, CopyBranchInfluenceShape(mapOrigCopy, _influenceshape))
	}

	return
}

func CopyBranchInfluence(mapOrigCopy map[any]any, influenceFrom *Influence) (influenceTo *Influence) {

	// influenceFrom has already been copied
	if _influenceTo, ok := mapOrigCopy[influenceFrom]; ok {
		influenceTo = _influenceTo.(*Influence)
		return
	}

	influenceTo = new(Influence)
	mapOrigCopy[influenceFrom] = influenceTo
	influenceFrom.CopyBasicFields(influenceTo)

	//insertion point for the staging of instances referenced by pointers
	if influenceFrom.SourceMovement != nil {
		influenceTo.SourceMovement = CopyBranchMovement(mapOrigCopy, influenceFrom.SourceMovement)
	}
	if influenceFrom.SourceArtefactType != nil {
		influenceTo.SourceArtefactType = CopyBranchArtefactType(mapOrigCopy, influenceFrom.SourceArtefactType)
	}
	if influenceFrom.SourceArtist != nil {
		influenceTo.SourceArtist = CopyBranchArtist(mapOrigCopy, influenceFrom.SourceArtist)
	}
	if influenceFrom.TargetMovement != nil {
		influenceTo.TargetMovement = CopyBranchMovement(mapOrigCopy, influenceFrom.TargetMovement)
	}
	if influenceFrom.TargetArtefactType != nil {
		influenceTo.TargetArtefactType = CopyBranchArtefactType(mapOrigCopy, influenceFrom.TargetArtefactType)
	}
	if influenceFrom.TargetArtist != nil {
		influenceTo.TargetArtist = CopyBranchArtist(mapOrigCopy, influenceFrom.TargetArtist)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchInfluenceShape(mapOrigCopy map[any]any, influenceshapeFrom *InfluenceShape) (influenceshapeTo *InfluenceShape) {

	// influenceshapeFrom has already been copied
	if _influenceshapeTo, ok := mapOrigCopy[influenceshapeFrom]; ok {
		influenceshapeTo = _influenceshapeTo.(*InfluenceShape)
		return
	}

	influenceshapeTo = new(InfluenceShape)
	mapOrigCopy[influenceshapeFrom] = influenceshapeTo
	influenceshapeFrom.CopyBasicFields(influenceshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if influenceshapeFrom.Influence != nil {
		influenceshapeTo.Influence = CopyBranchInfluence(mapOrigCopy, influenceshapeFrom.Influence)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range influenceshapeFrom.ControlPointShapes {
		influenceshapeTo.ControlPointShapes = append(influenceshapeTo.ControlPointShapes, CopyBranchControlPointShape(mapOrigCopy, _controlpointshape))
	}

	return
}

func CopyBranchMovement(mapOrigCopy map[any]any, movementFrom *Movement) (movementTo *Movement) {

	// movementFrom has already been copied
	if _movementTo, ok := mapOrigCopy[movementFrom]; ok {
		movementTo = _movementTo.(*Movement)
		return
	}

	movementTo = new(Movement)
	mapOrigCopy[movementFrom] = movementTo
	movementFrom.CopyBasicFields(movementTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMovementShape(mapOrigCopy map[any]any, movementshapeFrom *MovementShape) (movementshapeTo *MovementShape) {

	// movementshapeFrom has already been copied
	if _movementshapeTo, ok := mapOrigCopy[movementshapeFrom]; ok {
		movementshapeTo = _movementshapeTo.(*MovementShape)
		return
	}

	movementshapeTo = new(MovementShape)
	mapOrigCopy[movementshapeFrom] = movementshapeTo
	movementshapeFrom.CopyBasicFields(movementshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if movementshapeFrom.Movement != nil {
		movementshapeTo.Movement = CopyBranchMovement(mapOrigCopy, movementshapeFrom.Movement)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *ArtefactType:
		stage.UnstageBranchArtefactType(target)

	case *ArtefactTypeShape:
		stage.UnstageBranchArtefactTypeShape(target)

	case *Artist:
		stage.UnstageBranchArtist(target)

	case *ArtistShape:
		stage.UnstageBranchArtistShape(target)

	case *ControlPointShape:
		stage.UnstageBranchControlPointShape(target)

	case *Desk:
		stage.UnstageBranchDesk(target)

	case *Diagram:
		stage.UnstageBranchDiagram(target)

	case *Influence:
		stage.UnstageBranchInfluence(target)

	case *InfluenceShape:
		stage.UnstageBranchInfluenceShape(target)

	case *Movement:
		stage.UnstageBranchMovement(target)

	case *MovementShape:
		stage.UnstageBranchMovementShape(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchArtefactType(artefacttype *ArtefactType) {

	// check if instance is already staged
	if !IsStaged(stage, artefacttype) {
		return
	}

	artefacttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchArtefactTypeShape(artefacttypeshape *ArtefactTypeShape) {

	// check if instance is already staged
	if !IsStaged(stage, artefacttypeshape) {
		return
	}

	artefacttypeshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if artefacttypeshape.ArtefactType != nil {
		UnstageBranch(stage, artefacttypeshape.ArtefactType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchArtist(artist *Artist) {

	// check if instance is already staged
	if !IsStaged(stage, artist) {
		return
	}

	artist.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchArtistShape(artistshape *ArtistShape) {

	// check if instance is already staged
	if !IsStaged(stage, artistshape) {
		return
	}

	artistshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if artistshape.Artist != nil {
		UnstageBranch(stage, artistshape.Artist)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchControlPointShape(controlpointshape *ControlPointShape) {

	// check if instance is already staged
	if !IsStaged(stage, controlpointshape) {
		return
	}

	controlpointshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDesk(desk *Desk) {

	// check if instance is already staged
	if !IsStaged(stage, desk) {
		return
	}

	desk.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if desk.SelectedDiagram != nil {
		UnstageBranch(stage, desk.SelectedDiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if !IsStaged(stage, diagram) {
		return
	}

	diagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _movementshape := range diagram.MovementShapes {
		UnstageBranch(stage, _movementshape)
	}
	for _, _artefacttypeshape := range diagram.ArtefactTypeShapes {
		UnstageBranch(stage, _artefacttypeshape)
	}
	for _, _artistshape := range diagram.ArtistShapes {
		UnstageBranch(stage, _artistshape)
	}
	for _, _influenceshape := range diagram.InfluenceShapes {
		UnstageBranch(stage, _influenceshape)
	}

}

func (stage *Stage) UnstageBranchInfluence(influence *Influence) {

	// check if instance is already staged
	if !IsStaged(stage, influence) {
		return
	}

	influence.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if influence.SourceMovement != nil {
		UnstageBranch(stage, influence.SourceMovement)
	}
	if influence.SourceArtefactType != nil {
		UnstageBranch(stage, influence.SourceArtefactType)
	}
	if influence.SourceArtist != nil {
		UnstageBranch(stage, influence.SourceArtist)
	}
	if influence.TargetMovement != nil {
		UnstageBranch(stage, influence.TargetMovement)
	}
	if influence.TargetArtefactType != nil {
		UnstageBranch(stage, influence.TargetArtefactType)
	}
	if influence.TargetArtist != nil {
		UnstageBranch(stage, influence.TargetArtist)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchInfluenceShape(influenceshape *InfluenceShape) {

	// check if instance is already staged
	if !IsStaged(stage, influenceshape) {
		return
	}

	influenceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if influenceshape.Influence != nil {
		UnstageBranch(stage, influenceshape.Influence)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range influenceshape.ControlPointShapes {
		UnstageBranch(stage, _controlpointshape)
	}

}

func (stage *Stage) UnstageBranchMovement(movement *Movement) {

	// check if instance is already staged
	if !IsStaged(stage, movement) {
		return
	}

	movement.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchMovementShape(movementshape *MovementShape) {

	// check if instance is already staged
	if !IsStaged(stage, movementshape) {
		return
	}

	movementshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if movementshape.Movement != nil {
		UnstageBranch(stage, movementshape.Movement)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}
