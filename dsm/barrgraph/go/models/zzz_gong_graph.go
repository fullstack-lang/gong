// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (artefacttype *ArtefactType) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ArtefactTypes[artefacttype]

	return
}

func (stage *Stage) IsStagedArtefactType(artefacttype *ArtefactType) (ok bool) {

	return artefacttype.GongIsStaged(stage)
}

func (artefacttypeshape *ArtefactTypeShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ArtefactTypeShapes[artefacttypeshape]

	return
}

func (stage *Stage) IsStagedArtefactTypeShape(artefacttypeshape *ArtefactTypeShape) (ok bool) {

	return artefacttypeshape.GongIsStaged(stage)
}

func (artist *Artist) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Artists[artist]

	return
}

func (stage *Stage) IsStagedArtist(artist *Artist) (ok bool) {

	return artist.GongIsStaged(stage)
}

func (artistshape *ArtistShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ArtistShapes[artistshape]

	return
}

func (stage *Stage) IsStagedArtistShape(artistshape *ArtistShape) (ok bool) {

	return artistshape.GongIsStaged(stage)
}

func (controlpointshape *ControlPointShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ControlPointShapes[controlpointshape]

	return
}

func (stage *Stage) IsStagedControlPointShape(controlpointshape *ControlPointShape) (ok bool) {

	return controlpointshape.GongIsStaged(stage)
}

func (desk *Desk) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Desks[desk]

	return
}

func (stage *Stage) IsStagedDesk(desk *Desk) (ok bool) {

	return desk.GongIsStaged(stage)
}

func (diagram *Diagram) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Diagrams[diagram]

	return
}

func (stage *Stage) IsStagedDiagram(diagram *Diagram) (ok bool) {

	return diagram.GongIsStaged(stage)
}

func (influence *Influence) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Influences[influence]

	return
}

func (stage *Stage) IsStagedInfluence(influence *Influence) (ok bool) {

	return influence.GongIsStaged(stage)
}

func (influenceshape *InfluenceShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.InfluenceShapes[influenceshape]

	return
}

func (stage *Stage) IsStagedInfluenceShape(influenceshape *InfluenceShape) (ok bool) {

	return influenceshape.GongIsStaged(stage)
}

func (library *Library) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	return library.GongIsStaged(stage)
}

func (movement *Movement) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Movements[movement]

	return
}

func (stage *Stage) IsStagedMovement(movement *Movement) (ok bool) {

	return movement.GongIsStaged(stage)
}

func (movementshape *MovementShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.MovementShapes[movementshape]

	return
}

func (stage *Stage) IsStagedMovementShape(movementshape *MovementShape) (ok bool) {

	return movementshape.GongIsStaged(stage)
}

func (place *Place) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Places[place]

	return
}

func (stage *Stage) IsStagedPlace(place *Place) (ok bool) {

	return place.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// StageBranch is a backward-compatible package-level forwarder.
func StageBranch(stage *Stage, instance GongstructIF) {
	stage.StageBranch(instance)
}

// insertion point for stage branch per struct
func (artefacttype *ArtefactType) GongStageBranch(stage *Stage) {
	stage.StageBranchArtefactType(artefacttype)
}

func (stage *Stage) StageBranchArtefactType(artefacttype *ArtefactType) {

	// check if instance is already staged
	if stage.IsStaged(artefacttype) {
		return
	}

	artefacttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (artefacttypeshape *ArtefactTypeShape) GongStageBranch(stage *Stage) {
	stage.StageBranchArtefactTypeShape(artefacttypeshape)
}

func (stage *Stage) StageBranchArtefactTypeShape(artefacttypeshape *ArtefactTypeShape) {

	// check if instance is already staged
	if stage.IsStaged(artefacttypeshape) {
		return
	}

	artefacttypeshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if artefacttypeshape.ArtefactType != nil {
		stage.StageBranch(artefacttypeshape.ArtefactType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (artist *Artist) GongStageBranch(stage *Stage) {
	stage.StageBranchArtist(artist)
}

func (stage *Stage) StageBranchArtist(artist *Artist) {

	// check if instance is already staged
	if stage.IsStaged(artist) {
		return
	}

	artist.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if artist.Place != nil {
		stage.StageBranch(artist.Place)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (artistshape *ArtistShape) GongStageBranch(stage *Stage) {
	stage.StageBranchArtistShape(artistshape)
}

func (stage *Stage) StageBranchArtistShape(artistshape *ArtistShape) {

	// check if instance is already staged
	if stage.IsStaged(artistshape) {
		return
	}

	artistshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if artistshape.Artist != nil {
		stage.StageBranch(artistshape.Artist)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (controlpointshape *ControlPointShape) GongStageBranch(stage *Stage) {
	stage.StageBranchControlPointShape(controlpointshape)
}

func (stage *Stage) StageBranchControlPointShape(controlpointshape *ControlPointShape) {

	// check if instance is already staged
	if stage.IsStaged(controlpointshape) {
		return
	}

	controlpointshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (desk *Desk) GongStageBranch(stage *Stage) {
	stage.StageBranchDesk(desk)
}

func (stage *Stage) StageBranchDesk(desk *Desk) {

	// check if instance is already staged
	if stage.IsStaged(desk) {
		return
	}

	desk.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if desk.SelectedDiagram != nil {
		stage.StageBranch(desk.SelectedDiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagram *Diagram) GongStageBranch(stage *Stage) {
	stage.StageBranchDiagram(diagram)
}

func (stage *Stage) StageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if stage.IsStaged(diagram) {
		return
	}

	diagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _movementshape := range diagram.MovementShapes {
		stage.StageBranch(_movementshape)
	}
	for _, _artefacttypeshape := range diagram.ArtefactTypeShapes {
		stage.StageBranch(_artefacttypeshape)
	}
	for _, _artistshape := range diagram.ArtistShapes {
		stage.StageBranch(_artistshape)
	}
	for _, _influenceshape := range diagram.InfluenceShapes {
		stage.StageBranch(_influenceshape)
	}

}

func (influence *Influence) GongStageBranch(stage *Stage) {
	stage.StageBranchInfluence(influence)
}

func (stage *Stage) StageBranchInfluence(influence *Influence) {

	// check if instance is already staged
	if stage.IsStaged(influence) {
		return
	}

	influence.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if influence.SourceMovement != nil {
		stage.StageBranch(influence.SourceMovement)
	}
	if influence.SourceArtefactType != nil {
		stage.StageBranch(influence.SourceArtefactType)
	}
	if influence.SourceArtist != nil {
		stage.StageBranch(influence.SourceArtist)
	}
	if influence.TargetMovement != nil {
		stage.StageBranch(influence.TargetMovement)
	}
	if influence.TargetArtefactType != nil {
		stage.StageBranch(influence.TargetArtefactType)
	}
	if influence.TargetArtist != nil {
		stage.StageBranch(influence.TargetArtist)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (influenceshape *InfluenceShape) GongStageBranch(stage *Stage) {
	stage.StageBranchInfluenceShape(influenceshape)
}

func (stage *Stage) StageBranchInfluenceShape(influenceshape *InfluenceShape) {

	// check if instance is already staged
	if stage.IsStaged(influenceshape) {
		return
	}

	influenceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if influenceshape.Influence != nil {
		stage.StageBranch(influenceshape.Influence)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range influenceshape.ControlPointShapes {
		stage.StageBranch(_controlpointshape)
	}

}

func (library *Library) GongStageBranch(stage *Stage) {
	stage.StageBranchLibrary(library)
}

func (stage *Stage) StageBranchLibrary(library *Library) {

	// check if instance is already staged
	if stage.IsStaged(library) {
		return
	}

	library.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		stage.StageBranch(_library)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		stage.StageBranch(_library)
	}

}

func (movement *Movement) GongStageBranch(stage *Stage) {
	stage.StageBranchMovement(movement)
}

func (stage *Stage) StageBranchMovement(movement *Movement) {

	// check if instance is already staged
	if stage.IsStaged(movement) {
		return
	}

	movement.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _place := range movement.Places {
		stage.StageBranch(_place)
	}

}

func (movementshape *MovementShape) GongStageBranch(stage *Stage) {
	stage.StageBranchMovementShape(movementshape)
}

func (stage *Stage) StageBranchMovementShape(movementshape *MovementShape) {

	// check if instance is already staged
	if stage.IsStaged(movementshape) {
		return
	}

	movementshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if movementshape.Movement != nil {
		stage.StageBranch(movementshape.Movement)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (place *Place) GongStageBranch(stage *Stage) {
	stage.StageBranchPlace(place)
}

func (stage *Stage) StageBranchPlace(place *Place) {

	// check if instance is already staged
	if stage.IsStaged(place) {
		return
	}

	place.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *ArtefactType:
		toT := GongCopyBranchArtefactType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ArtefactTypeShape:
		toT := GongCopyBranchArtefactTypeShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Artist:
		toT := GongCopyBranchArtist(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ArtistShape:
		toT := GongCopyBranchArtistShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ControlPointShape:
		toT := GongCopyBranchControlPointShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Desk:
		toT := GongCopyBranchDesk(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Diagram:
		toT := GongCopyBranchDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Influence:
		toT := GongCopyBranchInfluence(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *InfluenceShape:
		toT := GongCopyBranchInfluenceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := GongCopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Movement:
		toT := GongCopyBranchMovement(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MovementShape:
		toT := GongCopyBranchMovementShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Place:
		toT := GongCopyBranchPlace(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchArtefactType(mapOrigCopy map[any]any, artefacttypeFrom *ArtefactType) (artefacttypeTo *ArtefactType) {

	// artefacttypeFrom has already been copied
	if _artefacttypeTo, ok := mapOrigCopy[artefacttypeFrom]; ok {
		artefacttypeTo = _artefacttypeTo.(*ArtefactType)
		return
	}

	artefacttypeTo = new(ArtefactType)
	mapOrigCopy[artefacttypeFrom] = artefacttypeTo
	artefacttypeFrom.GongCopyBasicFields(artefacttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArtefactTypeShape(mapOrigCopy map[any]any, artefacttypeshapeFrom *ArtefactTypeShape) (artefacttypeshapeTo *ArtefactTypeShape) {

	// artefacttypeshapeFrom has already been copied
	if _artefacttypeshapeTo, ok := mapOrigCopy[artefacttypeshapeFrom]; ok {
		artefacttypeshapeTo = _artefacttypeshapeTo.(*ArtefactTypeShape)
		return
	}

	artefacttypeshapeTo = new(ArtefactTypeShape)
	mapOrigCopy[artefacttypeshapeFrom] = artefacttypeshapeTo
	artefacttypeshapeFrom.GongCopyBasicFields(artefacttypeshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if artefacttypeshapeFrom.ArtefactType != nil {
		artefacttypeshapeTo.ArtefactType = GongCopyBranchArtefactType(mapOrigCopy, artefacttypeshapeFrom.ArtefactType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArtist(mapOrigCopy map[any]any, artistFrom *Artist) (artistTo *Artist) {

	// artistFrom has already been copied
	if _artistTo, ok := mapOrigCopy[artistFrom]; ok {
		artistTo = _artistTo.(*Artist)
		return
	}

	artistTo = new(Artist)
	mapOrigCopy[artistFrom] = artistTo
	artistFrom.GongCopyBasicFields(artistTo)

	//insertion point for the staging of instances referenced by pointers
	if artistFrom.Place != nil {
		artistTo.Place = GongCopyBranchPlace(mapOrigCopy, artistFrom.Place)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArtistShape(mapOrigCopy map[any]any, artistshapeFrom *ArtistShape) (artistshapeTo *ArtistShape) {

	// artistshapeFrom has already been copied
	if _artistshapeTo, ok := mapOrigCopy[artistshapeFrom]; ok {
		artistshapeTo = _artistshapeTo.(*ArtistShape)
		return
	}

	artistshapeTo = new(ArtistShape)
	mapOrigCopy[artistshapeFrom] = artistshapeTo
	artistshapeFrom.GongCopyBasicFields(artistshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if artistshapeFrom.Artist != nil {
		artistshapeTo.Artist = GongCopyBranchArtist(mapOrigCopy, artistshapeFrom.Artist)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchControlPointShape(mapOrigCopy map[any]any, controlpointshapeFrom *ControlPointShape) (controlpointshapeTo *ControlPointShape) {

	// controlpointshapeFrom has already been copied
	if _controlpointshapeTo, ok := mapOrigCopy[controlpointshapeFrom]; ok {
		controlpointshapeTo = _controlpointshapeTo.(*ControlPointShape)
		return
	}

	controlpointshapeTo = new(ControlPointShape)
	mapOrigCopy[controlpointshapeFrom] = controlpointshapeTo
	controlpointshapeFrom.GongCopyBasicFields(controlpointshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDesk(mapOrigCopy map[any]any, deskFrom *Desk) (deskTo *Desk) {

	// deskFrom has already been copied
	if _deskTo, ok := mapOrigCopy[deskFrom]; ok {
		deskTo = _deskTo.(*Desk)
		return
	}

	deskTo = new(Desk)
	mapOrigCopy[deskFrom] = deskTo
	deskFrom.GongCopyBasicFields(deskTo)

	//insertion point for the staging of instances referenced by pointers
	if deskFrom.SelectedDiagram != nil {
		deskTo.SelectedDiagram = GongCopyBranchDiagram(mapOrigCopy, deskFrom.SelectedDiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDiagram(mapOrigCopy map[any]any, diagramFrom *Diagram) (diagramTo *Diagram) {

	// diagramFrom has already been copied
	if _diagramTo, ok := mapOrigCopy[diagramFrom]; ok {
		diagramTo = _diagramTo.(*Diagram)
		return
	}

	diagramTo = new(Diagram)
	mapOrigCopy[diagramFrom] = diagramTo
	diagramFrom.GongCopyBasicFields(diagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _movementshape := range diagramFrom.MovementShapes {
		diagramTo.MovementShapes = append(diagramTo.MovementShapes, GongCopyBranchMovementShape(mapOrigCopy, _movementshape))
	}
	for _, _artefacttypeshape := range diagramFrom.ArtefactTypeShapes {
		diagramTo.ArtefactTypeShapes = append(diagramTo.ArtefactTypeShapes, GongCopyBranchArtefactTypeShape(mapOrigCopy, _artefacttypeshape))
	}
	for _, _artistshape := range diagramFrom.ArtistShapes {
		diagramTo.ArtistShapes = append(diagramTo.ArtistShapes, GongCopyBranchArtistShape(mapOrigCopy, _artistshape))
	}
	for _, _influenceshape := range diagramFrom.InfluenceShapes {
		diagramTo.InfluenceShapes = append(diagramTo.InfluenceShapes, GongCopyBranchInfluenceShape(mapOrigCopy, _influenceshape))
	}

	return
}

func GongCopyBranchInfluence(mapOrigCopy map[any]any, influenceFrom *Influence) (influenceTo *Influence) {

	// influenceFrom has already been copied
	if _influenceTo, ok := mapOrigCopy[influenceFrom]; ok {
		influenceTo = _influenceTo.(*Influence)
		return
	}

	influenceTo = new(Influence)
	mapOrigCopy[influenceFrom] = influenceTo
	influenceFrom.GongCopyBasicFields(influenceTo)

	//insertion point for the staging of instances referenced by pointers
	if influenceFrom.SourceMovement != nil {
		influenceTo.SourceMovement = GongCopyBranchMovement(mapOrigCopy, influenceFrom.SourceMovement)
	}
	if influenceFrom.SourceArtefactType != nil {
		influenceTo.SourceArtefactType = GongCopyBranchArtefactType(mapOrigCopy, influenceFrom.SourceArtefactType)
	}
	if influenceFrom.SourceArtist != nil {
		influenceTo.SourceArtist = GongCopyBranchArtist(mapOrigCopy, influenceFrom.SourceArtist)
	}
	if influenceFrom.TargetMovement != nil {
		influenceTo.TargetMovement = GongCopyBranchMovement(mapOrigCopy, influenceFrom.TargetMovement)
	}
	if influenceFrom.TargetArtefactType != nil {
		influenceTo.TargetArtefactType = GongCopyBranchArtefactType(mapOrigCopy, influenceFrom.TargetArtefactType)
	}
	if influenceFrom.TargetArtist != nil {
		influenceTo.TargetArtist = GongCopyBranchArtist(mapOrigCopy, influenceFrom.TargetArtist)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchInfluenceShape(mapOrigCopy map[any]any, influenceshapeFrom *InfluenceShape) (influenceshapeTo *InfluenceShape) {

	// influenceshapeFrom has already been copied
	if _influenceshapeTo, ok := mapOrigCopy[influenceshapeFrom]; ok {
		influenceshapeTo = _influenceshapeTo.(*InfluenceShape)
		return
	}

	influenceshapeTo = new(InfluenceShape)
	mapOrigCopy[influenceshapeFrom] = influenceshapeTo
	influenceshapeFrom.GongCopyBasicFields(influenceshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if influenceshapeFrom.Influence != nil {
		influenceshapeTo.Influence = GongCopyBranchInfluence(mapOrigCopy, influenceshapeFrom.Influence)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range influenceshapeFrom.ControlPointShapes {
		influenceshapeTo.ControlPointShapes = append(influenceshapeTo.ControlPointShapes, GongCopyBranchControlPointShape(mapOrigCopy, _controlpointshape))
	}

	return
}

func GongCopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {

	// libraryFrom has already been copied
	if _libraryTo, ok := mapOrigCopy[libraryFrom]; ok {
		libraryTo = _libraryTo.(*Library)
		return
	}

	libraryTo = new(Library)
	mapOrigCopy[libraryFrom] = libraryTo
	libraryFrom.GongCopyBasicFields(libraryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range libraryFrom.SubLibraries {
		libraryTo.SubLibraries = append(libraryTo.SubLibraries, GongCopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _library := range libraryFrom.SubLibrariesWhoseNodeIsExpanded {
		libraryTo.SubLibrariesWhoseNodeIsExpanded = append(libraryTo.SubLibrariesWhoseNodeIsExpanded, GongCopyBranchLibrary(mapOrigCopy, _library))
	}

	return
}

func GongCopyBranchMovement(mapOrigCopy map[any]any, movementFrom *Movement) (movementTo *Movement) {

	// movementFrom has already been copied
	if _movementTo, ok := mapOrigCopy[movementFrom]; ok {
		movementTo = _movementTo.(*Movement)
		return
	}

	movementTo = new(Movement)
	mapOrigCopy[movementFrom] = movementTo
	movementFrom.GongCopyBasicFields(movementTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _place := range movementFrom.Places {
		movementTo.Places = append(movementTo.Places, GongCopyBranchPlace(mapOrigCopy, _place))
	}

	return
}

func GongCopyBranchMovementShape(mapOrigCopy map[any]any, movementshapeFrom *MovementShape) (movementshapeTo *MovementShape) {

	// movementshapeFrom has already been copied
	if _movementshapeTo, ok := mapOrigCopy[movementshapeFrom]; ok {
		movementshapeTo = _movementshapeTo.(*MovementShape)
		return
	}

	movementshapeTo = new(MovementShape)
	mapOrigCopy[movementshapeFrom] = movementshapeTo
	movementshapeFrom.GongCopyBasicFields(movementshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if movementshapeFrom.Movement != nil {
		movementshapeTo.Movement = GongCopyBranchMovement(mapOrigCopy, movementshapeFrom.Movement)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPlace(mapOrigCopy map[any]any, placeFrom *Place) (placeTo *Place) {

	// placeFrom has already been copied
	if _placeTo, ok := mapOrigCopy[placeFrom]; ok {
		placeTo = _placeTo.(*Place)
		return
	}

	placeTo = new(Place)
	mapOrigCopy[placeFrom] = placeTo
	placeFrom.GongCopyBasicFields(placeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (artefacttype *ArtefactType) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchArtefactType(artefacttype)
}

func (stage *Stage) UnstageBranchArtefactType(artefacttype *ArtefactType) {

	// check if instance is already staged
	if !stage.IsStaged(artefacttype) {
		return
	}

	artefacttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (artefacttypeshape *ArtefactTypeShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchArtefactTypeShape(artefacttypeshape)
}

func (stage *Stage) UnstageBranchArtefactTypeShape(artefacttypeshape *ArtefactTypeShape) {

	// check if instance is already staged
	if !stage.IsStaged(artefacttypeshape) {
		return
	}

	artefacttypeshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if artefacttypeshape.ArtefactType != nil {
		stage.UnstageBranch(artefacttypeshape.ArtefactType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (artist *Artist) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchArtist(artist)
}

func (stage *Stage) UnstageBranchArtist(artist *Artist) {

	// check if instance is already staged
	if !stage.IsStaged(artist) {
		return
	}

	artist.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if artist.Place != nil {
		stage.UnstageBranch(artist.Place)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (artistshape *ArtistShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchArtistShape(artistshape)
}

func (stage *Stage) UnstageBranchArtistShape(artistshape *ArtistShape) {

	// check if instance is already staged
	if !stage.IsStaged(artistshape) {
		return
	}

	artistshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if artistshape.Artist != nil {
		stage.UnstageBranch(artistshape.Artist)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (controlpointshape *ControlPointShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchControlPointShape(controlpointshape)
}

func (stage *Stage) UnstageBranchControlPointShape(controlpointshape *ControlPointShape) {

	// check if instance is already staged
	if !stage.IsStaged(controlpointshape) {
		return
	}

	controlpointshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (desk *Desk) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDesk(desk)
}

func (stage *Stage) UnstageBranchDesk(desk *Desk) {

	// check if instance is already staged
	if !stage.IsStaged(desk) {
		return
	}

	desk.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if desk.SelectedDiagram != nil {
		stage.UnstageBranch(desk.SelectedDiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagram *Diagram) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDiagram(diagram)
}

func (stage *Stage) UnstageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if !stage.IsStaged(diagram) {
		return
	}

	diagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _movementshape := range diagram.MovementShapes {
		stage.UnstageBranch(_movementshape)
	}
	for _, _artefacttypeshape := range diagram.ArtefactTypeShapes {
		stage.UnstageBranch(_artefacttypeshape)
	}
	for _, _artistshape := range diagram.ArtistShapes {
		stage.UnstageBranch(_artistshape)
	}
	for _, _influenceshape := range diagram.InfluenceShapes {
		stage.UnstageBranch(_influenceshape)
	}

}

func (influence *Influence) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchInfluence(influence)
}

func (stage *Stage) UnstageBranchInfluence(influence *Influence) {

	// check if instance is already staged
	if !stage.IsStaged(influence) {
		return
	}

	influence.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if influence.SourceMovement != nil {
		stage.UnstageBranch(influence.SourceMovement)
	}
	if influence.SourceArtefactType != nil {
		stage.UnstageBranch(influence.SourceArtefactType)
	}
	if influence.SourceArtist != nil {
		stage.UnstageBranch(influence.SourceArtist)
	}
	if influence.TargetMovement != nil {
		stage.UnstageBranch(influence.TargetMovement)
	}
	if influence.TargetArtefactType != nil {
		stage.UnstageBranch(influence.TargetArtefactType)
	}
	if influence.TargetArtist != nil {
		stage.UnstageBranch(influence.TargetArtist)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (influenceshape *InfluenceShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchInfluenceShape(influenceshape)
}

func (stage *Stage) UnstageBranchInfluenceShape(influenceshape *InfluenceShape) {

	// check if instance is already staged
	if !stage.IsStaged(influenceshape) {
		return
	}

	influenceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if influenceshape.Influence != nil {
		stage.UnstageBranch(influenceshape.Influence)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range influenceshape.ControlPointShapes {
		stage.UnstageBranch(_controlpointshape)
	}

}

func (library *Library) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchLibrary(library)
}

func (stage *Stage) UnstageBranchLibrary(library *Library) {

	// check if instance is already staged
	if !stage.IsStaged(library) {
		return
	}

	library.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		stage.UnstageBranch(_library)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		stage.UnstageBranch(_library)
	}

}

func (movement *Movement) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchMovement(movement)
}

func (stage *Stage) UnstageBranchMovement(movement *Movement) {

	// check if instance is already staged
	if !stage.IsStaged(movement) {
		return
	}

	movement.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _place := range movement.Places {
		stage.UnstageBranch(_place)
	}

}

func (movementshape *MovementShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchMovementShape(movementshape)
}

func (stage *Stage) UnstageBranchMovementShape(movementshape *MovementShape) {

	// check if instance is already staged
	if !stage.IsStaged(movementshape) {
		return
	}

	movementshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if movementshape.Movement != nil {
		stage.UnstageBranch(movementshape.Movement)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (place *Place) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPlace(place)
}

func (stage *Stage) UnstageBranchPlace(place *Place) {

	// check if instance is already staged
	if !stage.IsStaged(place) {
		return
	}

	place.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *ArtefactType) GongReconstructPointersFromReferences(stage *Stage, instance *ArtefactType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ArtefactTypeShape) GongReconstructPointersFromReferences(stage *Stage, instance *ArtefactTypeShape) {
	// insertion point for pointers field
	if instance.ArtefactType != nil {
		reference.ArtefactType = stage.ArtefactTypes_reference[instance.ArtefactType]
	}
	// insertion point for slice of pointers field
}

func (reference *Artist) GongReconstructPointersFromReferences(stage *Stage, instance *Artist) {
	// insertion point for pointers field
	if instance.Place != nil {
		reference.Place = stage.Places_reference[instance.Place]
	}
	// insertion point for slice of pointers field
}

func (reference *ArtistShape) GongReconstructPointersFromReferences(stage *Stage, instance *ArtistShape) {
	// insertion point for pointers field
	if instance.Artist != nil {
		reference.Artist = stage.Artists_reference[instance.Artist]
	}
	// insertion point for slice of pointers field
}

func (reference *ControlPointShape) GongReconstructPointersFromReferences(stage *Stage, instance *ControlPointShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Desk) GongReconstructPointersFromReferences(stage *Stage, instance *Desk) {
	// insertion point for pointers field
	if instance.SelectedDiagram != nil {
		reference.SelectedDiagram = stage.Diagrams_reference[instance.SelectedDiagram]
	}
	// insertion point for slice of pointers field
}

func (reference *Diagram) GongReconstructPointersFromReferences(stage *Stage, instance *Diagram) {
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
}

func (reference *Influence) GongReconstructPointersFromReferences(stage *Stage, instance *Influence) {
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
}

func (reference *InfluenceShape) GongReconstructPointersFromReferences(stage *Stage, instance *InfluenceShape) {
	// insertion point for pointers field
	if instance.Influence != nil {
		reference.Influence = stage.Influences_reference[instance.Influence]
	}
	// insertion point for slice of pointers field
	reference.ControlPointShapes = reference.ControlPointShapes[:0]
	for _, _b := range instance.ControlPointShapes {
		reference.ControlPointShapes = append(reference.ControlPointShapes, stage.ControlPointShapes_reference[_b])
	}
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.SubLibraries = reference.SubLibraries[:0]
	for _, _b := range instance.SubLibraries {
		reference.SubLibraries = append(reference.SubLibraries, stage.Librarys_reference[_b])
	}
	reference.SubLibrariesWhoseNodeIsExpanded = reference.SubLibrariesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.SubLibrariesWhoseNodeIsExpanded {
		reference.SubLibrariesWhoseNodeIsExpanded = append(reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_reference[_b])
	}
}

func (reference *Movement) GongReconstructPointersFromReferences(stage *Stage, instance *Movement) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Places = reference.Places[:0]
	for _, _b := range instance.Places {
		reference.Places = append(reference.Places, stage.Places_reference[_b])
	}
}

func (reference *MovementShape) GongReconstructPointersFromReferences(stage *Stage, instance *MovementShape) {
	// insertion point for pointers field
	if instance.Movement != nil {
		reference.Movement = stage.Movements_reference[instance.Movement]
	}
	// insertion point for slice of pointers field
}

func (reference *Place) GongReconstructPointersFromReferences(stage *Stage, instance *Place) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *ArtefactType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ArtefactTypeShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ArtefactType; _reference != nil {
		reference.ArtefactType = nil
		if _instance, ok := stage.ArtefactTypes_instance[_reference]; ok {
			reference.ArtefactType = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Artist) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Place; _reference != nil {
		reference.Place = nil
		if _instance, ok := stage.Places_instance[_reference]; ok {
			reference.Place = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ArtistShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Artist; _reference != nil {
		reference.Artist = nil
		if _instance, ok := stage.Artists_instance[_reference]; ok {
			reference.Artist = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ControlPointShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Desk) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.SelectedDiagram; _reference != nil {
		reference.SelectedDiagram = nil
		if _instance, ok := stage.Diagrams_instance[_reference]; ok {
			reference.SelectedDiagram = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Diagram) GongReconstructPointersFromInstances(stage *Stage) {
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
}

func (reference *Influence) GongReconstructPointersFromInstances(stage *Stage) {
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
}

func (reference *InfluenceShape) GongReconstructPointersFromInstances(stage *Stage) {
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
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _SubLibraries []*Library
	for _, _reference := range reference.SubLibraries {
		if _instance, ok := stage.Librarys_instance[_reference]; ok {
			_SubLibraries = append(_SubLibraries, _instance)
		}
	}
	reference.SubLibraries = _SubLibraries
	var _SubLibrariesWhoseNodeIsExpanded []*Library
	for _, _reference := range reference.SubLibrariesWhoseNodeIsExpanded {
		if _instance, ok := stage.Librarys_instance[_reference]; ok {
			_SubLibrariesWhoseNodeIsExpanded = append(_SubLibrariesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.SubLibrariesWhoseNodeIsExpanded = _SubLibrariesWhoseNodeIsExpanded
}

func (reference *Movement) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Places []*Place
	for _, _reference := range reference.Places {
		if _instance, ok := stage.Places_instance[_reference]; ok {
			_Places = append(_Places, _instance)
		}
	}
	reference.Places = _Places
}

func (reference *MovementShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Movement; _reference != nil {
		reference.Movement = nil
		if _instance, ok := stage.Movements_instance[_reference]; ok {
			reference.Movement = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Place) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (artefacttype *ArtefactType) GongDiff(stage *Stage, artefacttypeOther *ArtefactType) (diffs []string) {
	// insertion point for field diffs
	if artefacttype.Name != artefacttypeOther.Name {
		diffs = append(diffs, artefacttype.GongMarshallField(stage, "Name"))
	}
	if artefacttype.ComputedPrefix != artefacttypeOther.ComputedPrefix {
		diffs = append(diffs, artefacttype.GongMarshallField(stage, "ComputedPrefix"))
	}
	if artefacttype.IsExpanded != artefacttypeOther.IsExpanded {
		diffs = append(diffs, artefacttype.GongMarshallField(stage, "IsExpanded"))
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
	if artefacttypeshape.IsHidden != artefacttypeshapeOther.IsHidden {
		diffs = append(diffs, artefacttypeshape.GongMarshallField(stage, "IsHidden"))
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
	if artist.ComputedPrefix != artistOther.ComputedPrefix {
		diffs = append(diffs, artist.GongMarshallField(stage, "ComputedPrefix"))
	}
	if artist.IsExpanded != artistOther.IsExpanded {
		diffs = append(diffs, artist.GongMarshallField(stage, "IsExpanded"))
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
	if artistshape.IsHidden != artistshapeOther.IsHidden {
		diffs = append(diffs, artistshape.GongMarshallField(stage, "IsHidden"))
	}
	if artistshape.ImagePng_X != artistshapeOther.ImagePng_X {
		diffs = append(diffs, artistshape.GongMarshallField(stage, "ImagePng_X"))
	}
	if artistshape.ImagePng_Y != artistshapeOther.ImagePng_Y {
		diffs = append(diffs, artistshape.GongMarshallField(stage, "ImagePng_Y"))
	}
	if artistshape.ImagePng_Width != artistshapeOther.ImagePng_Width {
		diffs = append(diffs, artistshape.GongMarshallField(stage, "ImagePng_Width"))
	}
	if artistshape.ImagePng_Height != artistshapeOther.ImagePng_Height {
		diffs = append(diffs, artistshape.GongMarshallField(stage, "ImagePng_Height"))
	}
	if artistshape.ImagePng_X_Offset != artistshapeOther.ImagePng_X_Offset {
		diffs = append(diffs, artistshape.GongMarshallField(stage, "ImagePng_X_Offset"))
	}
	if artistshape.ImagePng_Y_Offset != artistshapeOther.ImagePng_Y_Offset {
		diffs = append(diffs, artistshape.GongMarshallField(stage, "ImagePng_Y_Offset"))
	}
	if artistshape.ImagePng_RectAnchorType != artistshapeOther.ImagePng_RectAnchorType {
		diffs = append(diffs, artistshape.GongMarshallField(stage, "ImagePng_RectAnchorType"))
	}
	if artistshape.ImagePngBase64Content != artistshapeOther.ImagePngBase64Content {
		diffs = append(diffs, artistshape.GongMarshallField(stage, "ImagePngBase64Content"))
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
	if diagram.ComputedPrefix != diagramOther.ComputedPrefix {
		diffs = append(diffs, diagram.GongMarshallField(stage, "ComputedPrefix"))
	}
	if diagram.IsExpanded != diagramOther.IsExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsExpanded"))
	}
	if diagram.IsChecked != diagramOther.IsChecked {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsChecked"))
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
		ops := stage.Diff(
			diagram,
			"MovementShapes",
			len(diagramOther.MovementShapes),
			len(diagram.MovementShapes),
			func(i, j int) bool {
				return diagramOther.MovementShapes[i] == diagram.MovementShapes[j]
			},
			func(j int) string {
				return diagram.MovementShapes[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			diagram,
			"ArtefactTypeShapes",
			len(diagramOther.ArtefactTypeShapes),
			len(diagram.ArtefactTypeShapes),
			func(i, j int) bool {
				return diagramOther.ArtefactTypeShapes[i] == diagram.ArtefactTypeShapes[j]
			},
			func(j int) string {
				return diagram.ArtefactTypeShapes[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			diagram,
			"ArtistShapes",
			len(diagramOther.ArtistShapes),
			len(diagram.ArtistShapes),
			func(i, j int) bool {
				return diagramOther.ArtistShapes[i] == diagram.ArtistShapes[j]
			},
			func(j int) string {
				return diagram.ArtistShapes[j].GongGetIdentifier(stage)
			},
		)
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
		ops := stage.Diff(
			diagram,
			"InfluenceShapes",
			len(diagramOther.InfluenceShapes),
			len(diagram.InfluenceShapes),
			func(i, j int) bool {
				return diagramOther.InfluenceShapes[i] == diagram.InfluenceShapes[j]
			},
			func(j int) string {
				return diagram.InfluenceShapes[j].GongGetIdentifier(stage)
			},
		)
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
	if diagram.IsMovementCategoryHidden != diagramOther.IsMovementCategoryHidden {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsMovementCategoryHidden"))
	}
	if diagram.IsArtefactTypeCategoryHidden != diagramOther.IsArtefactTypeCategoryHidden {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsArtefactTypeCategoryHidden"))
	}
	if diagram.IsArtistCategoryHidden != diagramOther.IsArtistCategoryHidden {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsArtistCategoryHidden"))
	}
	if diagram.IsInfluenceCategoryHidden != diagramOther.IsInfluenceCategoryHidden {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsInfluenceCategoryHidden"))
	}
	if diagram.StartDate != diagramOther.StartDate {
		diffs = append(diffs, diagram.GongMarshallField(stage, "StartDate"))
	}
	if diagram.EndDate != diagramOther.EndDate {
		diffs = append(diffs, diagram.GongMarshallField(stage, "EndDate"))
	}
	if diagram.NbYearsForIntervals != diagramOther.NbYearsForIntervals {
		diffs = append(diffs, diagram.GongMarshallField(stage, "NbYearsForIntervals"))
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
	if influence.ComputedPrefix != influenceOther.ComputedPrefix {
		diffs = append(diffs, influence.GongMarshallField(stage, "ComputedPrefix"))
	}
	if influence.IsExpanded != influenceOther.IsExpanded {
		diffs = append(diffs, influence.GongMarshallField(stage, "IsExpanded"))
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
	if influenceshape.IsHidden != influenceshapeOther.IsHidden {
		diffs = append(diffs, influenceshape.GongMarshallField(stage, "IsHidden"))
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
		ops := stage.Diff(
			influenceshape,
			"ControlPointShapes",
			len(influenceshapeOther.ControlPointShapes),
			len(influenceshape.ControlPointShapes),
			func(i, j int) bool {
				return influenceshapeOther.ControlPointShapes[i] == influenceshape.ControlPointShapes[j]
			},
			func(j int) string {
				return influenceshape.ControlPointShapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (library *Library) GongDiff(stage *Stage, libraryOther *Library) (diffs []string) {
	// insertion point for field diffs
	if library.Name != libraryOther.Name {
		diffs = append(diffs, library.GongMarshallField(stage, "Name"))
	}
	if library.Description != libraryOther.Description {
		diffs = append(diffs, library.GongMarshallField(stage, "Description"))
	}
	if library.ComputedPrefix != libraryOther.ComputedPrefix {
		diffs = append(diffs, library.GongMarshallField(stage, "ComputedPrefix"))
	}
	if library.IsExpanded != libraryOther.IsExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpanded"))
	}
	if library.IsRootLibrary != libraryOther.IsRootLibrary {
		diffs = append(diffs, library.GongMarshallField(stage, "IsRootLibrary"))
	}
	SubLibrariesDifferent := false
	if len(library.SubLibraries) != len(libraryOther.SubLibraries) {
		SubLibrariesDifferent = true
	} else {
		for i := range library.SubLibraries {
			if (library.SubLibraries[i] == nil) != (libraryOther.SubLibraries[i] == nil) {
				SubLibrariesDifferent = true
				break
			} else if library.SubLibraries[i] != nil && libraryOther.SubLibraries[i] != nil {
				// this is a pointer comparaison
				if library.SubLibraries[i] != libraryOther.SubLibraries[i] {
					SubLibrariesDifferent = true
					break
				}
			}
		}
	}
	if SubLibrariesDifferent {
		ops := stage.Diff(
			library,
			"SubLibraries",
			len(libraryOther.SubLibraries),
			len(library.SubLibraries),
			func(i, j int) bool {
				return libraryOther.SubLibraries[i] == library.SubLibraries[j]
			},
			func(j int) string {
				return library.SubLibraries[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if library.IsSubLibrariesNodeExpanded != libraryOther.IsSubLibrariesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
	}
	SubLibrariesWhoseNodeIsExpandedDifferent := false
	if len(library.SubLibrariesWhoseNodeIsExpanded) != len(libraryOther.SubLibrariesWhoseNodeIsExpanded) {
		SubLibrariesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.SubLibrariesWhoseNodeIsExpanded {
			if (library.SubLibrariesWhoseNodeIsExpanded[i] == nil) != (libraryOther.SubLibrariesWhoseNodeIsExpanded[i] == nil) {
				SubLibrariesWhoseNodeIsExpandedDifferent = true
				break
			} else if library.SubLibrariesWhoseNodeIsExpanded[i] != nil && libraryOther.SubLibrariesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.SubLibrariesWhoseNodeIsExpanded[i] != libraryOther.SubLibrariesWhoseNodeIsExpanded[i] {
					SubLibrariesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if SubLibrariesWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			library,
			"SubLibrariesWhoseNodeIsExpanded",
			len(libraryOther.SubLibrariesWhoseNodeIsExpanded),
			len(library.SubLibrariesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return libraryOther.SubLibrariesWhoseNodeIsExpanded[i] == library.SubLibrariesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return library.SubLibrariesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if library.NbPixPerCharacter != libraryOther.NbPixPerCharacter {
		diffs = append(diffs, library.GongMarshallField(stage, "NbPixPerCharacter"))
	}
	if library.LogoSVGFile != libraryOther.LogoSVGFile {
		diffs = append(diffs, library.GongMarshallField(stage, "LogoSVGFile"))
	}
	if library.IsExpandedTmp != libraryOther.IsExpandedTmp {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpandedTmp"))
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
	if movement.ComputedPrefix != movementOther.ComputedPrefix {
		diffs = append(diffs, movement.GongMarshallField(stage, "ComputedPrefix"))
	}
	if movement.IsExpanded != movementOther.IsExpanded {
		diffs = append(diffs, movement.GongMarshallField(stage, "IsExpanded"))
	}
	if movement.Date != movementOther.Date {
		diffs = append(diffs, movement.GongMarshallField(stage, "Date"))
	}
	if movement.HideDate != movementOther.HideDate {
		diffs = append(diffs, movement.GongMarshallField(stage, "HideDate"))
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
		ops := stage.Diff(
			movement,
			"Places",
			len(movementOther.Places),
			len(movement.Places),
			func(i, j int) bool {
				return movementOther.Places[i] == movement.Places[j]
			},
			func(j int) string {
				return movement.Places[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if movement.HasTaxonomicFilter != movementOther.HasTaxonomicFilter {
		diffs = append(diffs, movement.GongMarshallField(stage, "HasTaxonomicFilter"))
	}
	if movement.TaxonomicFilter != movementOther.TaxonomicFilter {
		diffs = append(diffs, movement.GongMarshallField(stage, "TaxonomicFilter"))
	}
	if movement.IsFeatured != movementOther.IsFeatured {
		diffs = append(diffs, movement.GongMarshallField(stage, "IsFeatured"))
	}
	if movement.FeaturePrefix != movementOther.FeaturePrefix {
		diffs = append(diffs, movement.GongMarshallField(stage, "FeaturePrefix"))
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
	if movementshape.IsHidden != movementshapeOther.IsHidden {
		diffs = append(diffs, movementshape.GongMarshallField(stage, "IsHidden"))
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

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if equal(i, j) {
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
		if equal(i-1, j-1) {
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

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := 0; k < m; k++ {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := 0; k < n; k++ {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}
