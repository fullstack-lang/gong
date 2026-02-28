// generated code - do not edit
package models

import "fmt"

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

	case *Place:
		ok = stage.IsStagedPlace(target)

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

	case *Place:
		ok = stage.IsStagedPlace(target)

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

func (stage *Stage) IsStagedPlace(place *Place) (ok bool) {

	_, ok = stage.Places[place]

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

	case *Place:
		stage.StageBranchPlace(target)

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
	if artist.Place != nil {
		StageBranch(stage, artist.Place)
	}

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
	for _, _place := range movement.Places {
		StageBranch(stage, _place)
	}

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

func (stage *Stage) StageBranchPlace(place *Place) {

	// check if instance is already staged
	if IsStaged(stage, place) {
		return
	}

	place.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

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

	case *Place:
		toT := CopyBranchPlace(mapOrigCopy, fromT)
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
	if artistFrom.Place != nil {
		artistTo.Place = CopyBranchPlace(mapOrigCopy, artistFrom.Place)
	}

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
	for _, _place := range movementFrom.Places {
		movementTo.Places = append(movementTo.Places, CopyBranchPlace(mapOrigCopy, _place))
	}

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

func CopyBranchPlace(mapOrigCopy map[any]any, placeFrom *Place) (placeTo *Place) {

	// placeFrom has already been copied
	if _placeTo, ok := mapOrigCopy[placeFrom]; ok {
		placeTo = _placeTo.(*Place)
		return
	}

	placeTo = new(Place)
	mapOrigCopy[placeFrom] = placeTo
	placeFrom.CopyBasicFields(placeTo)

	//insertion point for the staging of instances referenced by pointers

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

	case *Place:
		stage.UnstageBranchPlace(target)

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
	if artist.Place != nil {
		UnstageBranch(stage, artist.Place)
	}

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
	for _, _place := range movement.Places {
		UnstageBranch(stage, _place)
	}

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

func (stage *Stage) UnstageBranchPlace(place *Place) {

	// check if instance is already staged
	if !IsStaged(stage, place) {
		return
	}

	place.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *ArtefactType) GongReconstructPointersFromReferences(stage *Stage, instance *ArtefactType) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field

	return
}

func (reference *ArtefactTypeShape) GongReconstructPointersFromReferences(stage *Stage, instance *ArtefactTypeShape) () {
	// insertion point for pointers field
	if instance.ArtefactType != nil {
		reference.ArtefactType = stage.ArtefactTypes_reference[instance.ArtefactType]
	}
	// insertion point for slice of pointers field

	return
}

func (reference *Artist) GongReconstructPointersFromReferences(stage *Stage, instance *Artist) () {
	// insertion point for pointers field
	if instance.Place != nil {
		reference.Place = stage.Places_reference[instance.Place]
	}
	// insertion point for slice of pointers field

	return
}

func (reference *ArtistShape) GongReconstructPointersFromReferences(stage *Stage, instance *ArtistShape) () {
	// insertion point for pointers field
	if instance.Artist != nil {
		reference.Artist = stage.Artists_reference[instance.Artist]
	}
	// insertion point for slice of pointers field

	return
}

func (reference *ControlPointShape) GongReconstructPointersFromReferences(stage *Stage, instance *ControlPointShape) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field

	return
}

func (reference *Desk) GongReconstructPointersFromReferences(stage *Stage, instance *Desk) () {
	// insertion point for pointers field
	if instance.SelectedDiagram != nil {
		reference.SelectedDiagram = stage.Diagrams_reference[instance.SelectedDiagram]
	}
	// insertion point for slice of pointers field

	return
}

func (reference *Diagram) GongReconstructPointersFromReferences(stage *Stage, instance *Diagram) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.MovementShapes = reference.MovementShapes[:0]
	for _, _b := range instance.MovementShapes {
		reference.MovementShapes = append(reference.MovementShapes, stage.MovementShapes_reference[_b])
	}
	reference.ArtefactTypeShapes = reference.ArtefactTypeShapes[:0]
	for _, _b := range instance.ArtefactTypeShapes {
		reference.ArtefactTypeShapes = append(reference.ArtefactTypeShapes, stage.ArtefactTypeShapes_reference[_b])
	}
	reference.ArtistShapes = reference.ArtistShapes[:0]
	for _, _b := range instance.ArtistShapes {
		reference.ArtistShapes = append(reference.ArtistShapes, stage.ArtistShapes_reference[_b])
	}
	reference.InfluenceShapes = reference.InfluenceShapes[:0]
	for _, _b := range instance.InfluenceShapes {
		reference.InfluenceShapes = append(reference.InfluenceShapes, stage.InfluenceShapes_reference[_b])
	}

	return
}

func (reference *Influence) GongReconstructPointersFromReferences(stage *Stage, instance *Influence) () {
	// insertion point for pointers field
	if instance.SourceMovement != nil {
		reference.SourceMovement = stage.Movements_reference[instance.SourceMovement]
	}
	if instance.SourceArtefactType != nil {
		reference.SourceArtefactType = stage.ArtefactTypes_reference[instance.SourceArtefactType]
	}
	if instance.SourceArtist != nil {
		reference.SourceArtist = stage.Artists_reference[instance.SourceArtist]
	}
	if instance.TargetMovement != nil {
		reference.TargetMovement = stage.Movements_reference[instance.TargetMovement]
	}
	if instance.TargetArtefactType != nil {
		reference.TargetArtefactType = stage.ArtefactTypes_reference[instance.TargetArtefactType]
	}
	if instance.TargetArtist != nil {
		reference.TargetArtist = stage.Artists_reference[instance.TargetArtist]
	}
	// insertion point for slice of pointers field

	return
}

func (reference *InfluenceShape) GongReconstructPointersFromReferences(stage *Stage, instance *InfluenceShape) () {
	// insertion point for pointers field
	if instance.Influence != nil {
		reference.Influence = stage.Influences_reference[instance.Influence]
	}
	// insertion point for slice of pointers field
	reference.ControlPointShapes = reference.ControlPointShapes[:0]
	for _, _b := range instance.ControlPointShapes {
		reference.ControlPointShapes = append(reference.ControlPointShapes, stage.ControlPointShapes_reference[_b])
	}

	return
}

func (reference *Movement) GongReconstructPointersFromReferences(stage *Stage, instance *Movement) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Places = reference.Places[:0]
	for _, _b := range instance.Places {
		reference.Places = append(reference.Places, stage.Places_reference[_b])
	}

	return
}

func (reference *MovementShape) GongReconstructPointersFromReferences(stage *Stage, instance *MovementShape) () {
	// insertion point for pointers field
	if instance.Movement != nil {
		reference.Movement = stage.Movements_reference[instance.Movement]
	}
	// insertion point for slice of pointers field

	return
}

func (reference *Place) GongReconstructPointersFromReferences(stage *Stage, instance *Place) () {
	// insertion point for pointers field
	// insertion point for slice of pointers field

	return
}

// insertion point for pointer reconstruction from instances
func (reference *ArtefactType) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields

	return
}

func (reference *ArtefactTypeShape) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	if _reference := reference.ArtefactType; _reference != nil {
		reference.ArtefactType = nil
		if _instance, ok := stage.ArtefactTypes_instance[_reference]; ok {
			reference.ArtefactType = _instance
		}
	}
	// insertion point for slice of pointers fields

	return
}

func (reference *Artist) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	if _reference := reference.Place; _reference != nil {
		reference.Place = nil
		if _instance, ok := stage.Places_instance[_reference]; ok {
			reference.Place = _instance
		}
	}
	// insertion point for slice of pointers fields

	return
}

func (reference *ArtistShape) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	if _reference := reference.Artist; _reference != nil {
		reference.Artist = nil
		if _instance, ok := stage.Artists_instance[_reference]; ok {
			reference.Artist = _instance
		}
	}
	// insertion point for slice of pointers fields

	return
}

func (reference *ControlPointShape) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields

	return
}

func (reference *Desk) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	if _reference := reference.SelectedDiagram; _reference != nil {
		reference.SelectedDiagram = nil
		if _instance, ok := stage.Diagrams_instance[_reference]; ok {
			reference.SelectedDiagram = _instance
		}
	}
	// insertion point for slice of pointers fields

	return
}

func (reference *Diagram) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _MovementShapes []*MovementShape
	for _, _reference := range reference.MovementShapes {
		if _instance, ok := stage.MovementShapes_instance[_reference]; ok {
			_MovementShapes = append(_MovementShapes, _instance)
		}
	}
	reference.MovementShapes = _MovementShapes
	var _ArtefactTypeShapes []*ArtefactTypeShape
	for _, _reference := range reference.ArtefactTypeShapes {
		if _instance, ok := stage.ArtefactTypeShapes_instance[_reference]; ok {
			_ArtefactTypeShapes = append(_ArtefactTypeShapes, _instance)
		}
	}
	reference.ArtefactTypeShapes = _ArtefactTypeShapes
	var _ArtistShapes []*ArtistShape
	for _, _reference := range reference.ArtistShapes {
		if _instance, ok := stage.ArtistShapes_instance[_reference]; ok {
			_ArtistShapes = append(_ArtistShapes, _instance)
		}
	}
	reference.ArtistShapes = _ArtistShapes
	var _InfluenceShapes []*InfluenceShape
	for _, _reference := range reference.InfluenceShapes {
		if _instance, ok := stage.InfluenceShapes_instance[_reference]; ok {
			_InfluenceShapes = append(_InfluenceShapes, _instance)
		}
	}
	reference.InfluenceShapes = _InfluenceShapes

	return
}

func (reference *Influence) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	if _reference := reference.SourceMovement; _reference != nil {
		reference.SourceMovement = nil
		if _instance, ok := stage.Movements_instance[_reference]; ok {
			reference.SourceMovement = _instance
		}
	}
	if _reference := reference.SourceArtefactType; _reference != nil {
		reference.SourceArtefactType = nil
		if _instance, ok := stage.ArtefactTypes_instance[_reference]; ok {
			reference.SourceArtefactType = _instance
		}
	}
	if _reference := reference.SourceArtist; _reference != nil {
		reference.SourceArtist = nil
		if _instance, ok := stage.Artists_instance[_reference]; ok {
			reference.SourceArtist = _instance
		}
	}
	if _reference := reference.TargetMovement; _reference != nil {
		reference.TargetMovement = nil
		if _instance, ok := stage.Movements_instance[_reference]; ok {
			reference.TargetMovement = _instance
		}
	}
	if _reference := reference.TargetArtefactType; _reference != nil {
		reference.TargetArtefactType = nil
		if _instance, ok := stage.ArtefactTypes_instance[_reference]; ok {
			reference.TargetArtefactType = _instance
		}
	}
	if _reference := reference.TargetArtist; _reference != nil {
		reference.TargetArtist = nil
		if _instance, ok := stage.Artists_instance[_reference]; ok {
			reference.TargetArtist = _instance
		}
	}
	// insertion point for slice of pointers fields

	return
}

func (reference *InfluenceShape) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	if _reference := reference.Influence; _reference != nil {
		reference.Influence = nil
		if _instance, ok := stage.Influences_instance[_reference]; ok {
			reference.Influence = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _ControlPointShapes []*ControlPointShape
	for _, _reference := range reference.ControlPointShapes {
		if _instance, ok := stage.ControlPointShapes_instance[_reference]; ok {
			_ControlPointShapes = append(_ControlPointShapes, _instance)
		}
	}
	reference.ControlPointShapes = _ControlPointShapes

	return
}

func (reference *Movement) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Places []*Place
	for _, _reference := range reference.Places {
		if _instance, ok := stage.Places_instance[_reference]; ok {
			_Places = append(_Places, _instance)
		}
	}
	reference.Places = _Places

	return
}

func (reference *MovementShape) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	if _reference := reference.Movement; _reference != nil {
		reference.Movement = nil
		if _instance, ok := stage.Movements_instance[_reference]; ok {
			reference.Movement = _instance
		}
	}
	// insertion point for slice of pointers fields

	return
}

func (reference *Place) GongReconstructPointersFromInstances(stage *Stage) () {
	// insertion point for pointers field
	// insertion point for slice of pointers fields

	return
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (artefacttype *ArtefactType) GongDiff(stage *Stage, artefacttypeOther *ArtefactType) (diffs []string) {
	// insertion point for field diffs
	if artefacttype.Name != artefacttypeOther.Name {
		diffs = append(diffs, artefacttype.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (artefacttypeshape *ArtefactTypeShape) GongDiff(stage *Stage, artefacttypeshapeOther *ArtefactTypeShape) (diffs []string) {
	// insertion point for field diffs
	if artefacttypeshape.Name != artefacttypeshapeOther.Name {
		diffs = append(diffs, artefacttypeshape.GongMarshallField(stage, "Name"))
	}
	if (artefacttypeshape.ArtefactType == nil) != (artefacttypeshapeOther.ArtefactType == nil) {
		diffs = append(diffs, artefacttypeshape.GongMarshallField(stage, "ArtefactType"))
	} else if artefacttypeshape.ArtefactType != nil && artefacttypeshapeOther.ArtefactType != nil {
		if artefacttypeshape.ArtefactType != artefacttypeshapeOther.ArtefactType {
			diffs = append(diffs, artefacttypeshape.GongMarshallField(stage, "ArtefactType"))
		}
	}
	if artefacttypeshape.X != artefacttypeshapeOther.X {
		diffs = append(diffs, artefacttypeshape.GongMarshallField(stage, "X"))
	}
	if artefacttypeshape.Y != artefacttypeshapeOther.Y {
		diffs = append(diffs, artefacttypeshape.GongMarshallField(stage, "Y"))
	}
	if artefacttypeshape.Width != artefacttypeshapeOther.Width {
		diffs = append(diffs, artefacttypeshape.GongMarshallField(stage, "Width"))
	}
	if artefacttypeshape.Height != artefacttypeshapeOther.Height {
		diffs = append(diffs, artefacttypeshape.GongMarshallField(stage, "Height"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (artist *Artist) GongDiff(stage *Stage, artistOther *Artist) (diffs []string) {
	// insertion point for field diffs
	if artist.Name != artistOther.Name {
		diffs = append(diffs, artist.GongMarshallField(stage, "Name"))
	}
	if artist.IsDead != artistOther.IsDead {
		diffs = append(diffs, artist.GongMarshallField(stage, "IsDead"))
	}
	if artist.DateOfDeath != artistOther.DateOfDeath {
		diffs = append(diffs, artist.GongMarshallField(stage, "DateOfDeath"))
	}
	if (artist.Place == nil) != (artistOther.Place == nil) {
		diffs = append(diffs, artist.GongMarshallField(stage, "Place"))
	} else if artist.Place != nil && artistOther.Place != nil {
		if artist.Place != artistOther.Place {
			diffs = append(diffs, artist.GongMarshallField(stage, "Place"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (artistshape *ArtistShape) GongDiff(stage *Stage, artistshapeOther *ArtistShape) (diffs []string) {
	// insertion point for field diffs
	if artistshape.Name != artistshapeOther.Name {
		diffs = append(diffs, artistshape.GongMarshallField(stage, "Name"))
	}
	if (artistshape.Artist == nil) != (artistshapeOther.Artist == nil) {
		diffs = append(diffs, artistshape.GongMarshallField(stage, "Artist"))
	} else if artistshape.Artist != nil && artistshapeOther.Artist != nil {
		if artistshape.Artist != artistshapeOther.Artist {
			diffs = append(diffs, artistshape.GongMarshallField(stage, "Artist"))
		}
	}
	if artistshape.X != artistshapeOther.X {
		diffs = append(diffs, artistshape.GongMarshallField(stage, "X"))
	}
	if artistshape.Y != artistshapeOther.Y {
		diffs = append(diffs, artistshape.GongMarshallField(stage, "Y"))
	}
	if artistshape.Width != artistshapeOther.Width {
		diffs = append(diffs, artistshape.GongMarshallField(stage, "Width"))
	}
	if artistshape.Height != artistshapeOther.Height {
		diffs = append(diffs, artistshape.GongMarshallField(stage, "Height"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (controlpointshape *ControlPointShape) GongDiff(stage *Stage, controlpointshapeOther *ControlPointShape) (diffs []string) {
	// insertion point for field diffs
	if controlpointshape.Name != controlpointshapeOther.Name {
		diffs = append(diffs, controlpointshape.GongMarshallField(stage, "Name"))
	}
	if controlpointshape.X_Relative != controlpointshapeOther.X_Relative {
		diffs = append(diffs, controlpointshape.GongMarshallField(stage, "X_Relative"))
	}
	if controlpointshape.Y_Relative != controlpointshapeOther.Y_Relative {
		diffs = append(diffs, controlpointshape.GongMarshallField(stage, "Y_Relative"))
	}
	if controlpointshape.IsStartShapeTheClosestShape != controlpointshapeOther.IsStartShapeTheClosestShape {
		diffs = append(diffs, controlpointshape.GongMarshallField(stage, "IsStartShapeTheClosestShape"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (desk *Desk) GongDiff(stage *Stage, deskOther *Desk) (diffs []string) {
	// insertion point for field diffs
	if desk.Name != deskOther.Name {
		diffs = append(diffs, desk.GongMarshallField(stage, "Name"))
	}
	if (desk.SelectedDiagram == nil) != (deskOther.SelectedDiagram == nil) {
		diffs = append(diffs, desk.GongMarshallField(stage, "SelectedDiagram"))
	} else if desk.SelectedDiagram != nil && deskOther.SelectedDiagram != nil {
		if desk.SelectedDiagram != deskOther.SelectedDiagram {
			diffs = append(diffs, desk.GongMarshallField(stage, "SelectedDiagram"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (diagram *Diagram) GongDiff(stage *Stage, diagramOther *Diagram) (diffs []string) {
	// insertion point for field diffs
	if diagram.Name != diagramOther.Name {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Name"))
	}
	MovementShapesDifferent := false
	if len(diagram.MovementShapes) != len(diagramOther.MovementShapes) {
		MovementShapesDifferent = true
	} else {
		for i := range diagram.MovementShapes {
			if (diagram.MovementShapes[i] == nil) != (diagramOther.MovementShapes[i] == nil) {
				MovementShapesDifferent = true
				break
			} else if diagram.MovementShapes[i] != nil && diagramOther.MovementShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.MovementShapes[i] != diagramOther.MovementShapes[i] {
					MovementShapesDifferent = true
					break
				}
			}
		}
	}
	if MovementShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "MovementShapes", diagramOther.MovementShapes, diagram.MovementShapes)
		diffs = append(diffs, ops)
	}
	ArtefactTypeShapesDifferent := false
	if len(diagram.ArtefactTypeShapes) != len(diagramOther.ArtefactTypeShapes) {
		ArtefactTypeShapesDifferent = true
	} else {
		for i := range diagram.ArtefactTypeShapes {
			if (diagram.ArtefactTypeShapes[i] == nil) != (diagramOther.ArtefactTypeShapes[i] == nil) {
				ArtefactTypeShapesDifferent = true
				break
			} else if diagram.ArtefactTypeShapes[i] != nil && diagramOther.ArtefactTypeShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.ArtefactTypeShapes[i] != diagramOther.ArtefactTypeShapes[i] {
					ArtefactTypeShapesDifferent = true
					break
				}
			}
		}
	}
	if ArtefactTypeShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "ArtefactTypeShapes", diagramOther.ArtefactTypeShapes, diagram.ArtefactTypeShapes)
		diffs = append(diffs, ops)
	}
	ArtistShapesDifferent := false
	if len(diagram.ArtistShapes) != len(diagramOther.ArtistShapes) {
		ArtistShapesDifferent = true
	} else {
		for i := range diagram.ArtistShapes {
			if (diagram.ArtistShapes[i] == nil) != (diagramOther.ArtistShapes[i] == nil) {
				ArtistShapesDifferent = true
				break
			} else if diagram.ArtistShapes[i] != nil && diagramOther.ArtistShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.ArtistShapes[i] != diagramOther.ArtistShapes[i] {
					ArtistShapesDifferent = true
					break
				}
			}
		}
	}
	if ArtistShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "ArtistShapes", diagramOther.ArtistShapes, diagram.ArtistShapes)
		diffs = append(diffs, ops)
	}
	InfluenceShapesDifferent := false
	if len(diagram.InfluenceShapes) != len(diagramOther.InfluenceShapes) {
		InfluenceShapesDifferent = true
	} else {
		for i := range diagram.InfluenceShapes {
			if (diagram.InfluenceShapes[i] == nil) != (diagramOther.InfluenceShapes[i] == nil) {
				InfluenceShapesDifferent = true
				break
			} else if diagram.InfluenceShapes[i] != nil && diagramOther.InfluenceShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.InfluenceShapes[i] != diagramOther.InfluenceShapes[i] {
					InfluenceShapesDifferent = true
					break
				}
			}
		}
	}
	if InfluenceShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "InfluenceShapes", diagramOther.InfluenceShapes, diagram.InfluenceShapes)
		diffs = append(diffs, ops)
	}
	if diagram.IsEditable != diagramOther.IsEditable {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsEditable"))
	}
	if diagram.IsNodeExpanded != diagramOther.IsNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsNodeExpanded"))
	}
	if diagram.IsMovementCategoryNodeExpanded != diagramOther.IsMovementCategoryNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsMovementCategoryNodeExpanded"))
	}
	if diagram.IsArtefactTypeCategoryNodeExpanded != diagramOther.IsArtefactTypeCategoryNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsArtefactTypeCategoryNodeExpanded"))
	}
	if diagram.IsArtistCategoryNodeExpanded != diagramOther.IsArtistCategoryNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsArtistCategoryNodeExpanded"))
	}
	if diagram.IsInfluenceCategoryNodeExpanded != diagramOther.IsInfluenceCategoryNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsInfluenceCategoryNodeExpanded"))
	}
	if diagram.IsMovementCategoryShown != diagramOther.IsMovementCategoryShown {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsMovementCategoryShown"))
	}
	if diagram.IsArtefactTypeCategoryShown != diagramOther.IsArtefactTypeCategoryShown {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsArtefactTypeCategoryShown"))
	}
	if diagram.IsArtistCategoryShown != diagramOther.IsArtistCategoryShown {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsArtistCategoryShown"))
	}
	if diagram.IsInfluenceCategoryShown != diagramOther.IsInfluenceCategoryShown {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsInfluenceCategoryShown"))
	}
	if diagram.StartDate != diagramOther.StartDate {
		diffs = append(diffs, diagram.GongMarshallField(stage, "StartDate"))
	}
	if diagram.EndDate != diagramOther.EndDate {
		diffs = append(diffs, diagram.GongMarshallField(stage, "EndDate"))
	}
	if diagram.XMargin != diagramOther.XMargin {
		diffs = append(diffs, diagram.GongMarshallField(stage, "XMargin"))
	}
	if diagram.YMargin != diagramOther.YMargin {
		diffs = append(diffs, diagram.GongMarshallField(stage, "YMargin"))
	}
	if diagram.Height != diagramOther.Height {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Height"))
	}
	if diagram.NextVerticalDateXMargin != diagramOther.NextVerticalDateXMargin {
		diffs = append(diffs, diagram.GongMarshallField(stage, "NextVerticalDateXMargin"))
	}
	if diagram.RedColorCode != diagramOther.RedColorCode {
		diffs = append(diffs, diagram.GongMarshallField(stage, "RedColorCode"))
	}
	if diagram.BackgroundGreyColorCode != diagramOther.BackgroundGreyColorCode {
		diffs = append(diffs, diagram.GongMarshallField(stage, "BackgroundGreyColorCode"))
	}
	if diagram.GrayColorCode != diagramOther.GrayColorCode {
		diffs = append(diffs, diagram.GongMarshallField(stage, "GrayColorCode"))
	}
	if diagram.BottomBoxYOffset != diagramOther.BottomBoxYOffset {
		diffs = append(diffs, diagram.GongMarshallField(stage, "BottomBoxYOffset"))
	}
	if diagram.BottomBoxWidth != diagramOther.BottomBoxWidth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "BottomBoxWidth"))
	}
	if diagram.BottomBoxHeigth != diagramOther.BottomBoxHeigth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "BottomBoxHeigth"))
	}
	if diagram.BottomBoxFontSize != diagramOther.BottomBoxFontSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "BottomBoxFontSize"))
	}
	if diagram.BottomBoxFontWeigth != diagramOther.BottomBoxFontWeigth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "BottomBoxFontWeigth"))
	}
	if diagram.BottomBoxFontFamily != diagramOther.BottomBoxFontFamily {
		diffs = append(diffs, diagram.GongMarshallField(stage, "BottomBoxFontFamily"))
	}
	if diagram.BottomBoxLetterSpacing != diagramOther.BottomBoxLetterSpacing {
		diffs = append(diffs, diagram.GongMarshallField(stage, "BottomBoxLetterSpacing"))
	}
	if diagram.BottomBoxLetterColorCode != diagramOther.BottomBoxLetterColorCode {
		diffs = append(diffs, diagram.GongMarshallField(stage, "BottomBoxLetterColorCode"))
	}
	if diagram.MovementRectAnchorType != diagramOther.MovementRectAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementRectAnchorType"))
	}
	if diagram.MovementTextAnchorType != diagramOther.MovementTextAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementTextAnchorType"))
	}
	if diagram.MovementDominantBaselineType != diagramOther.MovementDominantBaselineType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementDominantBaselineType"))
	}
	if diagram.MovementFontSize != diagramOther.MovementFontSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementFontSize"))
	}
	if diagram.MajorMovementFontSize != diagramOther.MajorMovementFontSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MajorMovementFontSize"))
	}
	if diagram.MinorMovementFontSize != diagramOther.MinorMovementFontSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MinorMovementFontSize"))
	}
	if diagram.MovementFontWeigth != diagramOther.MovementFontWeigth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementFontWeigth"))
	}
	if diagram.MovementFontFamily != diagramOther.MovementFontFamily {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementFontFamily"))
	}
	if diagram.MovementLetterSpacing != diagramOther.MovementLetterSpacing {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementLetterSpacing"))
	}
	if diagram.AbstractMovementFontSize != diagramOther.AbstractMovementFontSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "AbstractMovementFontSize"))
	}
	if diagram.AbstractMovementRectAnchorType != diagramOther.AbstractMovementRectAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "AbstractMovementRectAnchorType"))
	}
	if diagram.AbstractMovementTextAnchorType != diagramOther.AbstractMovementTextAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "AbstractMovementTextAnchorType"))
	}
	if diagram.AbstractDominantBaselineType != diagramOther.AbstractDominantBaselineType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "AbstractDominantBaselineType"))
	}
	if diagram.MovementDateRectAnchorType != diagramOther.MovementDateRectAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementDateRectAnchorType"))
	}
	if diagram.MovementDateTextAnchorType != diagramOther.MovementDateTextAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementDateTextAnchorType"))
	}
	if diagram.MovementDateTextDominantBaselineType != diagramOther.MovementDateTextDominantBaselineType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementDateTextDominantBaselineType"))
	}
	if diagram.MovementDateAndPlacesFontSize != diagramOther.MovementDateAndPlacesFontSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementDateAndPlacesFontSize"))
	}
	if diagram.MovementDateAndPlacesFontWeigth != diagramOther.MovementDateAndPlacesFontWeigth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementDateAndPlacesFontWeigth"))
	}
	if diagram.MovementDateAndPlacesFontFamily != diagramOther.MovementDateAndPlacesFontFamily {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementDateAndPlacesFontFamily"))
	}
	if diagram.MovementDateAndPlacesLetterSpacing != diagramOther.MovementDateAndPlacesLetterSpacing {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementDateAndPlacesLetterSpacing"))
	}
	if diagram.MovementBelowArcY_Offset != diagramOther.MovementBelowArcY_Offset {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementBelowArcY_Offset"))
	}
	if diagram.MovementBelowArcY_OffsetPerPlace != diagramOther.MovementBelowArcY_OffsetPerPlace {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementBelowArcY_OffsetPerPlace"))
	}
	if diagram.MovementPlacesRectAnchorType != diagramOther.MovementPlacesRectAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementPlacesRectAnchorType"))
	}
	if diagram.MovementPlacesTextAnchorType != diagramOther.MovementPlacesTextAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementPlacesTextAnchorType"))
	}
	if diagram.MovementPlacesDominantBaselineType != diagramOther.MovementPlacesDominantBaselineType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MovementPlacesDominantBaselineType"))
	}
	if diagram.ArtefactTypeFontSize != diagramOther.ArtefactTypeFontSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtefactTypeFontSize"))
	}
	if diagram.ArtefactTypeFontWeigth != diagramOther.ArtefactTypeFontWeigth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtefactTypeFontWeigth"))
	}
	if diagram.ArtefactTypeFontFamily != diagramOther.ArtefactTypeFontFamily {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtefactTypeFontFamily"))
	}
	if diagram.ArtefactTypeLetterSpacing != diagramOther.ArtefactTypeLetterSpacing {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtefactTypeLetterSpacing"))
	}
	if diagram.ArtefactTypeRectAnchorType != diagramOther.ArtefactTypeRectAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtefactTypeRectAnchorType"))
	}
	if diagram.ArtefactDominantBaselineType != diagramOther.ArtefactDominantBaselineType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtefactDominantBaselineType"))
	}
	if diagram.ArtefactTypeStrokeWidth != diagramOther.ArtefactTypeStrokeWidth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtefactTypeStrokeWidth"))
	}
	if diagram.ArtistRectAnchorType != diagramOther.ArtistRectAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistRectAnchorType"))
	}
	if diagram.ArtistTextAnchorType != diagramOther.ArtistTextAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistTextAnchorType"))
	}
	if diagram.ArtistDominantBaselineType != diagramOther.ArtistDominantBaselineType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistDominantBaselineType"))
	}
	if diagram.ArtistFontSize != diagramOther.ArtistFontSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistFontSize"))
	}
	if diagram.MajorArtistFontSize != diagramOther.MajorArtistFontSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MajorArtistFontSize"))
	}
	if diagram.MinorArtistFontSize != diagramOther.MinorArtistFontSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "MinorArtistFontSize"))
	}
	if diagram.ArtistFontWeigth != diagramOther.ArtistFontWeigth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistFontWeigth"))
	}
	if diagram.ArtistFontFamily != diagramOther.ArtistFontFamily {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistFontFamily"))
	}
	if diagram.ArtistLetterSpacing != diagramOther.ArtistLetterSpacing {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistLetterSpacing"))
	}
	if diagram.ArtistDateRectAnchorType != diagramOther.ArtistDateRectAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistDateRectAnchorType"))
	}
	if diagram.ArtistDateTextAnchorType != diagramOther.ArtistDateTextAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistDateTextAnchorType"))
	}
	if diagram.ArtistDateDominantBaselineType != diagramOther.ArtistDateDominantBaselineType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistDateDominantBaselineType"))
	}
	if diagram.ArtistDateAndPlacesFontSize != diagramOther.ArtistDateAndPlacesFontSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistDateAndPlacesFontSize"))
	}
	if diagram.ArtistDateAndPlacesFontWeigth != diagramOther.ArtistDateAndPlacesFontWeigth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistDateAndPlacesFontWeigth"))
	}
	if diagram.ArtistDateAndPlacesFontFamily != diagramOther.ArtistDateAndPlacesFontFamily {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistDateAndPlacesFontFamily"))
	}
	if diagram.ArtistDateAndPlacesLetterSpacing != diagramOther.ArtistDateAndPlacesLetterSpacing {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistDateAndPlacesLetterSpacing"))
	}
	if diagram.ArtistPlacesRectAnchorType != diagramOther.ArtistPlacesRectAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistPlacesRectAnchorType"))
	}
	if diagram.ArtistPlacesTextAnchorType != diagramOther.ArtistPlacesTextAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistPlacesTextAnchorType"))
	}
	if diagram.ArtistPlacesDominantBaselineType != diagramOther.ArtistPlacesDominantBaselineType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ArtistPlacesDominantBaselineType"))
	}
	if diagram.InfluenceArrowSize != diagramOther.InfluenceArrowSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "InfluenceArrowSize"))
	}
	if diagram.InfluenceArrowStartOffset != diagramOther.InfluenceArrowStartOffset {
		diffs = append(diffs, diagram.GongMarshallField(stage, "InfluenceArrowStartOffset"))
	}
	if diagram.InfluenceArrowEndOffset != diagramOther.InfluenceArrowEndOffset {
		diffs = append(diffs, diagram.GongMarshallField(stage, "InfluenceArrowEndOffset"))
	}
	if diagram.InfluenceCornerRadius != diagramOther.InfluenceCornerRadius {
		diffs = append(diffs, diagram.GongMarshallField(stage, "InfluenceCornerRadius"))
	}
	if diagram.InfluenceDashedLinePattern != diagramOther.InfluenceDashedLinePattern {
		diffs = append(diffs, diagram.GongMarshallField(stage, "InfluenceDashedLinePattern"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (influence *Influence) GongDiff(stage *Stage, influenceOther *Influence) (diffs []string) {
	// insertion point for field diffs
	if influence.Name != influenceOther.Name {
		diffs = append(diffs, influence.GongMarshallField(stage, "Name"))
	}
	if (influence.SourceMovement == nil) != (influenceOther.SourceMovement == nil) {
		diffs = append(diffs, influence.GongMarshallField(stage, "SourceMovement"))
	} else if influence.SourceMovement != nil && influenceOther.SourceMovement != nil {
		if influence.SourceMovement != influenceOther.SourceMovement {
			diffs = append(diffs, influence.GongMarshallField(stage, "SourceMovement"))
		}
	}
	if (influence.SourceArtefactType == nil) != (influenceOther.SourceArtefactType == nil) {
		diffs = append(diffs, influence.GongMarshallField(stage, "SourceArtefactType"))
	} else if influence.SourceArtefactType != nil && influenceOther.SourceArtefactType != nil {
		if influence.SourceArtefactType != influenceOther.SourceArtefactType {
			diffs = append(diffs, influence.GongMarshallField(stage, "SourceArtefactType"))
		}
	}
	if (influence.SourceArtist == nil) != (influenceOther.SourceArtist == nil) {
		diffs = append(diffs, influence.GongMarshallField(stage, "SourceArtist"))
	} else if influence.SourceArtist != nil && influenceOther.SourceArtist != nil {
		if influence.SourceArtist != influenceOther.SourceArtist {
			diffs = append(diffs, influence.GongMarshallField(stage, "SourceArtist"))
		}
	}
	if (influence.TargetMovement == nil) != (influenceOther.TargetMovement == nil) {
		diffs = append(diffs, influence.GongMarshallField(stage, "TargetMovement"))
	} else if influence.TargetMovement != nil && influenceOther.TargetMovement != nil {
		if influence.TargetMovement != influenceOther.TargetMovement {
			diffs = append(diffs, influence.GongMarshallField(stage, "TargetMovement"))
		}
	}
	if (influence.TargetArtefactType == nil) != (influenceOther.TargetArtefactType == nil) {
		diffs = append(diffs, influence.GongMarshallField(stage, "TargetArtefactType"))
	} else if influence.TargetArtefactType != nil && influenceOther.TargetArtefactType != nil {
		if influence.TargetArtefactType != influenceOther.TargetArtefactType {
			diffs = append(diffs, influence.GongMarshallField(stage, "TargetArtefactType"))
		}
	}
	if (influence.TargetArtist == nil) != (influenceOther.TargetArtist == nil) {
		diffs = append(diffs, influence.GongMarshallField(stage, "TargetArtist"))
	} else if influence.TargetArtist != nil && influenceOther.TargetArtist != nil {
		if influence.TargetArtist != influenceOther.TargetArtist {
			diffs = append(diffs, influence.GongMarshallField(stage, "TargetArtist"))
		}
	}
	if influence.IsHypothtical != influenceOther.IsHypothtical {
		diffs = append(diffs, influence.GongMarshallField(stage, "IsHypothtical"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (influenceshape *InfluenceShape) GongDiff(stage *Stage, influenceshapeOther *InfluenceShape) (diffs []string) {
	// insertion point for field diffs
	if influenceshape.Name != influenceshapeOther.Name {
		diffs = append(diffs, influenceshape.GongMarshallField(stage, "Name"))
	}
	if (influenceshape.Influence == nil) != (influenceshapeOther.Influence == nil) {
		diffs = append(diffs, influenceshape.GongMarshallField(stage, "Influence"))
	} else if influenceshape.Influence != nil && influenceshapeOther.Influence != nil {
		if influenceshape.Influence != influenceshapeOther.Influence {
			diffs = append(diffs, influenceshape.GongMarshallField(stage, "Influence"))
		}
	}
	ControlPointShapesDifferent := false
	if len(influenceshape.ControlPointShapes) != len(influenceshapeOther.ControlPointShapes) {
		ControlPointShapesDifferent = true
	} else {
		for i := range influenceshape.ControlPointShapes {
			if (influenceshape.ControlPointShapes[i] == nil) != (influenceshapeOther.ControlPointShapes[i] == nil) {
				ControlPointShapesDifferent = true
				break
			} else if influenceshape.ControlPointShapes[i] != nil && influenceshapeOther.ControlPointShapes[i] != nil {
				// this is a pointer comparaison
				if influenceshape.ControlPointShapes[i] != influenceshapeOther.ControlPointShapes[i] {
					ControlPointShapesDifferent = true
					break
				}
			}
		}
	}
	if ControlPointShapesDifferent {
		ops := Diff(stage, influenceshape, influenceshapeOther, "ControlPointShapes", influenceshapeOther.ControlPointShapes, influenceshape.ControlPointShapes)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (movement *Movement) GongDiff(stage *Stage, movementOther *Movement) (diffs []string) {
	// insertion point for field diffs
	if movement.Name != movementOther.Name {
		diffs = append(diffs, movement.GongMarshallField(stage, "Name"))
	}
	if movement.Date != movementOther.Date {
		diffs = append(diffs, movement.GongMarshallField(stage, "Date"))
	}
	PlacesDifferent := false
	if len(movement.Places) != len(movementOther.Places) {
		PlacesDifferent = true
	} else {
		for i := range movement.Places {
			if (movement.Places[i] == nil) != (movementOther.Places[i] == nil) {
				PlacesDifferent = true
				break
			} else if movement.Places[i] != nil && movementOther.Places[i] != nil {
				// this is a pointer comparaison
				if movement.Places[i] != movementOther.Places[i] {
					PlacesDifferent = true
					break
				}
			}
		}
	}
	if PlacesDifferent {
		ops := Diff(stage, movement, movementOther, "Places", movementOther.Places, movement.Places)
		diffs = append(diffs, ops)
	}
	if movement.IsAbstract != movementOther.IsAbstract {
		diffs = append(diffs, movement.GongMarshallField(stage, "IsAbstract"))
	}
	if movement.IsModern != movementOther.IsModern {
		diffs = append(diffs, movement.GongMarshallField(stage, "IsModern"))
	}
	if movement.IsMajor != movementOther.IsMajor {
		diffs = append(diffs, movement.GongMarshallField(stage, "IsMajor"))
	}
	if movement.IsMinor != movementOther.IsMinor {
		diffs = append(diffs, movement.GongMarshallField(stage, "IsMinor"))
	}
	if movement.AdditionnalName != movementOther.AdditionnalName {
		diffs = append(diffs, movement.GongMarshallField(stage, "AdditionnalName"))
	}
	if movement.HideDate != movementOther.HideDate {
		diffs = append(diffs, movement.GongMarshallField(stage, "HideDate"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (movementshape *MovementShape) GongDiff(stage *Stage, movementshapeOther *MovementShape) (diffs []string) {
	// insertion point for field diffs
	if movementshape.Name != movementshapeOther.Name {
		diffs = append(diffs, movementshape.GongMarshallField(stage, "Name"))
	}
	if (movementshape.Movement == nil) != (movementshapeOther.Movement == nil) {
		diffs = append(diffs, movementshape.GongMarshallField(stage, "Movement"))
	} else if movementshape.Movement != nil && movementshapeOther.Movement != nil {
		if movementshape.Movement != movementshapeOther.Movement {
			diffs = append(diffs, movementshape.GongMarshallField(stage, "Movement"))
		}
	}
	if movementshape.X != movementshapeOther.X {
		diffs = append(diffs, movementshape.GongMarshallField(stage, "X"))
	}
	if movementshape.Y != movementshapeOther.Y {
		diffs = append(diffs, movementshape.GongMarshallField(stage, "Y"))
	}
	if movementshape.Width != movementshapeOther.Width {
		diffs = append(diffs, movementshape.GongMarshallField(stage, "Width"))
	}
	if movementshape.Height != movementshapeOther.Height {
		diffs = append(diffs, movementshape.GongMarshallField(stage, "Height"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (place *Place) GongDiff(stage *Stage, placeOther *Place) (diffs []string) {
	// insertion point for field diffs
	if place.Name != placeOther.Name {
		diffs = append(diffs, place.GongMarshallField(stage, "Name"))
	}

	return
}

// Diff returns the sequence of operations to transform oldSlice into newSlice.
// It requires type T to be comparable (e.g., pointers, ints, strings).
func Diff[T1, T2 PointerToGongstruct](stage *Stage, a, b T1, fieldName string, oldSlice, newSlice []T2) (ops string) {
	m, n := len(oldSlice), len(newSlice)

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if oldSlice[i] == newSlice[j] {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				if dp[i][j+1] > dp[i+1][j] {
					dp[i+1][j+1] = dp[i][j+1]
				} else {
					dp[i+1][j+1] = dp[i+1][j]
				}
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if oldSlice[i-1] == newSlice[j-1] {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Create a temporary view of what's left after deletions for tracking matches
	var currentLCS []T2
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			currentLCS = append(currentLCS, oldSlice[k])
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k, targetVal := range newSlice {
		if lcsIdx < len(currentLCS) && currentLCS[lcsIdx] == targetVal {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, targetVal.GongGetIdentifier(stage))
		}
	}

	return ops
}
