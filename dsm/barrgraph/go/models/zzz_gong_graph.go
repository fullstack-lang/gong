// generated code - do not edit
package models

import (
	"fmt"
	"slices"
)

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (artefacttype *ArtefactType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ArtefactTypes[artefacttype]
	return ok
}

func (artefacttypeshape *ArtefactTypeShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ArtefactTypeShapes[artefacttypeshape]
	return ok
}

func (artist *Artist) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Artists[artist]
	return ok
}

func (artistshape *ArtistShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ArtistShapes[artistshape]
	return ok
}

func (controlpointshape *ControlPointShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ControlPointShapes[controlpointshape]
	return ok
}

func (desk *Desk) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Desks[desk]
	return ok
}

func (diagram *Diagram) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Diagrams[diagram]
	return ok
}

func (influence *Influence) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Influences[influence]
	return ok
}

func (influenceshape *InfluenceShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.InfluenceShapes[influenceshape]
	return ok
}

func (library *Library) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Librarys[library]
	return ok
}

func (movement *Movement) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Movements[movement]
	return ok
}

func (movementshape *MovementShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MovementShapes[movementshape]
	return ok
}

func (place *Place) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Places[place]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (artefacttype *ArtefactType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(artefacttype) {
		return
	}

	artefacttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (artefacttypeshape *ArtefactTypeShape) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(controlpointshape) {
		return
	}

	controlpointshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (desk *Desk) GongStageBranch(stage *Stage) {

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
	var alreadyCopied bool
	artefacttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, artefacttypeFrom)
	if alreadyCopied {
		return
	}
	artefacttypeFrom.GongCopyBasicFields(artefacttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArtefactTypeShape(mapOrigCopy map[any]any, artefacttypeshapeFrom *ArtefactTypeShape) (artefacttypeshapeTo *ArtefactTypeShape) {
	var alreadyCopied bool
	artefacttypeshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, artefacttypeshapeFrom)
	if alreadyCopied {
		return
	}
	artefacttypeshapeFrom.GongCopyBasicFields(artefacttypeshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if artefacttypeshapeFrom.ArtefactType != nil {
		artefacttypeshapeTo.ArtefactType = GongCopyBranchArtefactType(mapOrigCopy, artefacttypeshapeFrom.ArtefactType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArtist(mapOrigCopy map[any]any, artistFrom *Artist) (artistTo *Artist) {
	var alreadyCopied bool
	artistTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, artistFrom)
	if alreadyCopied {
		return
	}
	artistFrom.GongCopyBasicFields(artistTo)

	//insertion point for the staging of instances referenced by pointers
	if artistFrom.Place != nil {
		artistTo.Place = GongCopyBranchPlace(mapOrigCopy, artistFrom.Place)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchArtistShape(mapOrigCopy map[any]any, artistshapeFrom *ArtistShape) (artistshapeTo *ArtistShape) {
	var alreadyCopied bool
	artistshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, artistshapeFrom)
	if alreadyCopied {
		return
	}
	artistshapeFrom.GongCopyBasicFields(artistshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if artistshapeFrom.Artist != nil {
		artistshapeTo.Artist = GongCopyBranchArtist(mapOrigCopy, artistshapeFrom.Artist)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchControlPointShape(mapOrigCopy map[any]any, controlpointshapeFrom *ControlPointShape) (controlpointshapeTo *ControlPointShape) {
	var alreadyCopied bool
	controlpointshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, controlpointshapeFrom)
	if alreadyCopied {
		return
	}
	controlpointshapeFrom.GongCopyBasicFields(controlpointshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDesk(mapOrigCopy map[any]any, deskFrom *Desk) (deskTo *Desk) {
	var alreadyCopied bool
	deskTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, deskFrom)
	if alreadyCopied {
		return
	}
	deskFrom.GongCopyBasicFields(deskTo)

	//insertion point for the staging of instances referenced by pointers
	if deskFrom.SelectedDiagram != nil {
		deskTo.SelectedDiagram = GongCopyBranchDiagram(mapOrigCopy, deskFrom.SelectedDiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDiagram(mapOrigCopy map[any]any, diagramFrom *Diagram) (diagramTo *Diagram) {
	var alreadyCopied bool
	diagramTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, diagramFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	influenceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, influenceFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	influenceshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, influenceshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	libraryTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, libraryFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	movementTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, movementFrom)
	if alreadyCopied {
		return
	}
	movementFrom.GongCopyBasicFields(movementTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _place := range movementFrom.Places {
		movementTo.Places = append(movementTo.Places, GongCopyBranchPlace(mapOrigCopy, _place))
	}

	return
}

func GongCopyBranchMovementShape(mapOrigCopy map[any]any, movementshapeFrom *MovementShape) (movementshapeTo *MovementShape) {
	var alreadyCopied bool
	movementshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, movementshapeFrom)
	if alreadyCopied {
		return
	}
	movementshapeFrom.GongCopyBasicFields(movementshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if movementshapeFrom.Movement != nil {
		movementshapeTo.Movement = GongCopyBranchMovement(mapOrigCopy, movementshapeFrom.Movement)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPlace(mapOrigCopy map[any]any, placeFrom *Place) (placeTo *Place) {
	var alreadyCopied bool
	placeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, placeFrom)
	if alreadyCopied {
		return
	}
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

// insertion point for unstage branch per struct
func (artefacttype *ArtefactType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(artefacttype) {
		return
	}

	artefacttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (artefacttypeshape *ArtefactTypeShape) GongUnstageBranch(stage *Stage) {

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

	// check if instance is already staged
	if !stage.IsStaged(controlpointshape) {
		return
	}

	controlpointshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (desk *Desk) GongUnstageBranch(stage *Stage) {

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
	__gong__reconstructPointer(&reference.ArtefactType, stage.ArtefactTypes_reference, instance.ArtefactType)
	// insertion point for slice of pointers field
}

func (reference *Artist) GongReconstructPointersFromReferences(stage *Stage, instance *Artist) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Place, stage.Places_reference, instance.Place)
	// insertion point for slice of pointers field
}

func (reference *ArtistShape) GongReconstructPointersFromReferences(stage *Stage, instance *ArtistShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Artist, stage.Artists_reference, instance.Artist)
	// insertion point for slice of pointers field
}

func (reference *ControlPointShape) GongReconstructPointersFromReferences(stage *Stage, instance *ControlPointShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Desk) GongReconstructPointersFromReferences(stage *Stage, instance *Desk) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.SelectedDiagram, stage.Diagrams_reference, instance.SelectedDiagram)
	// insertion point for slice of pointers field
}

func (reference *Diagram) GongReconstructPointersFromReferences(stage *Stage, instance *Diagram) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.MovementShapes, stage.MovementShapes_reference, instance.MovementShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ArtefactTypeShapes, stage.ArtefactTypeShapes_reference, instance.ArtefactTypeShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ArtistShapes, stage.ArtistShapes_reference, instance.ArtistShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.InfluenceShapes, stage.InfluenceShapes_reference, instance.InfluenceShapes)
}

func (reference *Influence) GongReconstructPointersFromReferences(stage *Stage, instance *Influence) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.SourceMovement, stage.Movements_reference, instance.SourceMovement)
	__gong__reconstructPointer(&reference.SourceArtefactType, stage.ArtefactTypes_reference, instance.SourceArtefactType)
	__gong__reconstructPointer(&reference.SourceArtist, stage.Artists_reference, instance.SourceArtist)
	__gong__reconstructPointer(&reference.TargetMovement, stage.Movements_reference, instance.TargetMovement)
	__gong__reconstructPointer(&reference.TargetArtefactType, stage.ArtefactTypes_reference, instance.TargetArtefactType)
	__gong__reconstructPointer(&reference.TargetArtist, stage.Artists_reference, instance.TargetArtist)
	// insertion point for slice of pointers field
}

func (reference *InfluenceShape) GongReconstructPointersFromReferences(stage *Stage, instance *InfluenceShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Influence, stage.Influences_reference, instance.Influence)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlPointShapes, stage.ControlPointShapes_reference, instance.ControlPointShapes)
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubLibraries, stage.Librarys_reference, instance.SubLibraries)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_reference, instance.SubLibrariesWhoseNodeIsExpanded)
}

func (reference *Movement) GongReconstructPointersFromReferences(stage *Stage, instance *Movement) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Places, stage.Places_reference, instance.Places)
}

func (reference *MovementShape) GongReconstructPointersFromReferences(stage *Stage, instance *MovementShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Movement, stage.Movements_reference, instance.Movement)
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
	__gong__reconstructPointerFromInstance(&reference.ArtefactType, stage.ArtefactTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *Artist) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Place, stage.Places_instance)
	// insertion point for slice of pointers fields
}

func (reference *ArtistShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Artist, stage.Artists_instance)
	// insertion point for slice of pointers fields
}

func (reference *ControlPointShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Desk) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.SelectedDiagram, stage.Diagrams_instance)
	// insertion point for slice of pointers fields
}

func (reference *Diagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.MovementShapes, stage.MovementShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ArtefactTypeShapes, stage.ArtefactTypeShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ArtistShapes, stage.ArtistShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.InfluenceShapes, stage.InfluenceShapes_instance)
}

func (reference *Influence) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.SourceMovement, stage.Movements_instance)
	__gong__reconstructPointerFromInstance(&reference.SourceArtefactType, stage.ArtefactTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.SourceArtist, stage.Artists_instance)
	__gong__reconstructPointerFromInstance(&reference.TargetMovement, stage.Movements_instance)
	__gong__reconstructPointerFromInstance(&reference.TargetArtefactType, stage.ArtefactTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.TargetArtist, stage.Artists_instance)
	// insertion point for slice of pointers fields
}

func (reference *InfluenceShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Influence, stage.Influences_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlPointShapes, stage.ControlPointShapes_instance)
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubLibraries, stage.Librarys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_instance)
}

func (reference *Movement) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Places, stage.Places_instance)
}

func (reference *MovementShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Movement, stage.Movements_instance)
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
	if artefacttypeshape.ArtefactType != artefacttypeshapeOther.ArtefactType {
		diffs = append(diffs, artefacttypeshape.GongMarshallField(stage, "ArtefactType"))
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
	if artist.Place != artistOther.Place {
		diffs = append(diffs, artist.GongMarshallField(stage, "Place"))
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
	if artistshape.Artist != artistshapeOther.Artist {
		diffs = append(diffs, artistshape.GongMarshallField(stage, "Artist"))
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
	if desk.SelectedDiagram != deskOther.SelectedDiagram {
		diffs = append(diffs, desk.GongMarshallField(stage, "SelectedDiagram"))
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
	if ops := __gong__diffSliceOfPointers(stage, diagram, "MovementShapes", diagramOther.MovementShapes, diagram.MovementShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ArtefactTypeShapes", diagramOther.ArtefactTypeShapes, diagram.ArtefactTypeShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "ArtistShapes", diagramOther.ArtistShapes, diagram.ArtistShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagram, "InfluenceShapes", diagramOther.InfluenceShapes, diagram.InfluenceShapes); ops != "" {
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
	if influence.SourceMovement != influenceOther.SourceMovement {
		diffs = append(diffs, influence.GongMarshallField(stage, "SourceMovement"))
	}
	if influence.SourceArtefactType != influenceOther.SourceArtefactType {
		diffs = append(diffs, influence.GongMarshallField(stage, "SourceArtefactType"))
	}
	if influence.SourceArtist != influenceOther.SourceArtist {
		diffs = append(diffs, influence.GongMarshallField(stage, "SourceArtist"))
	}
	if influence.TargetMovement != influenceOther.TargetMovement {
		diffs = append(diffs, influence.GongMarshallField(stage, "TargetMovement"))
	}
	if influence.TargetArtefactType != influenceOther.TargetArtefactType {
		diffs = append(diffs, influence.GongMarshallField(stage, "TargetArtefactType"))
	}
	if influence.TargetArtist != influenceOther.TargetArtist {
		diffs = append(diffs, influence.GongMarshallField(stage, "TargetArtist"))
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
	if influenceshape.Influence != influenceshapeOther.Influence {
		diffs = append(diffs, influenceshape.GongMarshallField(stage, "Influence"))
	}
	if influenceshape.IsHidden != influenceshapeOther.IsHidden {
		diffs = append(diffs, influenceshape.GongMarshallField(stage, "IsHidden"))
	}
	if ops := __gong__diffSliceOfPointers(stage, influenceshape, "ControlPointShapes", influenceshapeOther.ControlPointShapes, influenceshape.ControlPointShapes); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, library, "SubLibraries", libraryOther.SubLibraries, library.SubLibraries); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsSubLibrariesNodeExpanded != libraryOther.IsSubLibrariesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsSubLibrariesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "SubLibrariesWhoseNodeIsExpanded", libraryOther.SubLibrariesWhoseNodeIsExpanded, library.SubLibrariesWhoseNodeIsExpanded); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, movement, "Places", movementOther.Places, movement.Places); ops != "" {
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
	if movementshape.Movement != movementshapeOther.Movement {
		diffs = append(diffs, movementshape.GongMarshallField(stage, "Movement"))
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

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
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
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}

func __gong__copyBranchCheck[T any](mapOrigCopy map[any]any, from *T) (*T, bool) {
	if to, ok := mapOrigCopy[from]; ok {
		return to.(*T), true
	}
	to := new(T)
	mapOrigCopy[from] = to
	return to, false
}

func __gong__reconstructPointer[T comparable](field *T, refMap map[T]T, instanceField T) {
	var zero T
	if instanceField != zero {
		*field = refMap[instanceField]
	}
}

func __gong__reconstructPointerFromInstance[T comparable](field *T, instMap map[T]T) {
	ref := *field
	var zero T
	if ref != zero {
		*field = zero
		if inst, ok := instMap[ref]; ok {
			*field = inst
		}
	}
}

func __gong__reconstructSliceOfPointersFromReferences[T comparable](field *[]T, refMap map[T]T, instanceSlice []T) {
	*field = (*field)[:0]
	for _, b := range instanceSlice {
		*field = append(*field, refMap[b])
	}
}

func __gong__reconstructSliceOfPointersFromInstances[T comparable](field *[]T, instMap map[T]T) {
	var res []T
	for _, ref := range *field {
		if inst, ok := instMap[ref]; ok {
			res = append(res, inst)
		}
	}
	*field = res
}

func __gong__diffSliceOfPointers[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	instance GongstructIF,
	fieldName string,
	oldSlice, newSlice []T,
) string {
	if slices.Equal(oldSlice, newSlice) {
		return ""
	}
	return stage.Diff(
		instance,
		fieldName,
		len(oldSlice),
		len(newSlice),
		func(i, j int) bool {
			return oldSlice[i] == newSlice[j]
		},
		func(j int) string {
			return newSlice[j].GongGetIdentifier(stage)
		},
	)
}
