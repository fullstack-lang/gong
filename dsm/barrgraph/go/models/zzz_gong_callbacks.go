// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront(instance GongstructIF) {
	if instance != nil {
		instance.GongAfterCreateFromFront(stage)
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront(old, new GongstructIF) {
	if old != nil {
		old.GongOnAfterUpdateFromFront(stage, new)
	}
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront(staged, front GongstructIF) {
	if staged != nil {
		staged.GongAfterDeleteFromFront(stage, front)
	}
}

// insertion point
func (artefacttype *ArtefactType) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArtefactTypeCreateCallback != nil {
		stage.OnAfterArtefactTypeCreateCallback.OnAfterCreate(stage, artefacttype)
	}
}

func (artefacttype *ArtefactType) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArtefactTypeUpdateCallback != nil {
		var frontArtefactType *ArtefactType
		if front != nil {
			frontArtefactType, _ = front.(*ArtefactType)
		}
		stage.OnAfterArtefactTypeUpdateCallback.OnAfterUpdate(stage, artefacttype, frontArtefactType)
	}
}

func (artefacttype *ArtefactType) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArtefactTypeDeleteCallback != nil {
		var frontArtefactType *ArtefactType
		if front != nil {
			frontArtefactType, _ = front.(*ArtefactType)
		}
		stage.OnAfterArtefactTypeDeleteCallback.OnAfterDelete(stage, artefacttype, frontArtefactType)
	}
}

func (artefacttypeshape *ArtefactTypeShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArtefactTypeShapeCreateCallback != nil {
		stage.OnAfterArtefactTypeShapeCreateCallback.OnAfterCreate(stage, artefacttypeshape)
	}
}

func (artefacttypeshape *ArtefactTypeShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArtefactTypeShapeUpdateCallback != nil {
		var frontArtefactTypeShape *ArtefactTypeShape
		if front != nil {
			frontArtefactTypeShape, _ = front.(*ArtefactTypeShape)
		}
		stage.OnAfterArtefactTypeShapeUpdateCallback.OnAfterUpdate(stage, artefacttypeshape, frontArtefactTypeShape)
	}
}

func (artefacttypeshape *ArtefactTypeShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArtefactTypeShapeDeleteCallback != nil {
		var frontArtefactTypeShape *ArtefactTypeShape
		if front != nil {
			frontArtefactTypeShape, _ = front.(*ArtefactTypeShape)
		}
		stage.OnAfterArtefactTypeShapeDeleteCallback.OnAfterDelete(stage, artefacttypeshape, frontArtefactTypeShape)
	}
}

func (artist *Artist) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArtistCreateCallback != nil {
		stage.OnAfterArtistCreateCallback.OnAfterCreate(stage, artist)
	}
}

func (artist *Artist) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArtistUpdateCallback != nil {
		var frontArtist *Artist
		if front != nil {
			frontArtist, _ = front.(*Artist)
		}
		stage.OnAfterArtistUpdateCallback.OnAfterUpdate(stage, artist, frontArtist)
	}
}

func (artist *Artist) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArtistDeleteCallback != nil {
		var frontArtist *Artist
		if front != nil {
			frontArtist, _ = front.(*Artist)
		}
		stage.OnAfterArtistDeleteCallback.OnAfterDelete(stage, artist, frontArtist)
	}
}

func (artistshape *ArtistShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterArtistShapeCreateCallback != nil {
		stage.OnAfterArtistShapeCreateCallback.OnAfterCreate(stage, artistshape)
	}
}

func (artistshape *ArtistShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArtistShapeUpdateCallback != nil {
		var frontArtistShape *ArtistShape
		if front != nil {
			frontArtistShape, _ = front.(*ArtistShape)
		}
		stage.OnAfterArtistShapeUpdateCallback.OnAfterUpdate(stage, artistshape, frontArtistShape)
	}
}

func (artistshape *ArtistShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterArtistShapeDeleteCallback != nil {
		var frontArtistShape *ArtistShape
		if front != nil {
			frontArtistShape, _ = front.(*ArtistShape)
		}
		stage.OnAfterArtistShapeDeleteCallback.OnAfterDelete(stage, artistshape, frontArtistShape)
	}
}

func (controlpointshape *ControlPointShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterControlPointShapeCreateCallback != nil {
		stage.OnAfterControlPointShapeCreateCallback.OnAfterCreate(stage, controlpointshape)
	}
}

func (controlpointshape *ControlPointShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterControlPointShapeUpdateCallback != nil {
		var frontControlPointShape *ControlPointShape
		if front != nil {
			frontControlPointShape, _ = front.(*ControlPointShape)
		}
		stage.OnAfterControlPointShapeUpdateCallback.OnAfterUpdate(stage, controlpointshape, frontControlPointShape)
	}
}

func (controlpointshape *ControlPointShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterControlPointShapeDeleteCallback != nil {
		var frontControlPointShape *ControlPointShape
		if front != nil {
			frontControlPointShape, _ = front.(*ControlPointShape)
		}
		stage.OnAfterControlPointShapeDeleteCallback.OnAfterDelete(stage, controlpointshape, frontControlPointShape)
	}
}

func (desk *Desk) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDeskCreateCallback != nil {
		stage.OnAfterDeskCreateCallback.OnAfterCreate(stage, desk)
	}
}

func (desk *Desk) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeskUpdateCallback != nil {
		var frontDesk *Desk
		if front != nil {
			frontDesk, _ = front.(*Desk)
		}
		stage.OnAfterDeskUpdateCallback.OnAfterUpdate(stage, desk, frontDesk)
	}
}

func (desk *Desk) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeskDeleteCallback != nil {
		var frontDesk *Desk
		if front != nil {
			frontDesk, _ = front.(*Desk)
		}
		stage.OnAfterDeskDeleteCallback.OnAfterDelete(stage, desk, frontDesk)
	}
}

func (diagram *Diagram) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDiagramCreateCallback != nil {
		stage.OnAfterDiagramCreateCallback.OnAfterCreate(stage, diagram)
	}
}

func (diagram *Diagram) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramUpdateCallback != nil {
		var frontDiagram *Diagram
		if front != nil {
			frontDiagram, _ = front.(*Diagram)
		}
		stage.OnAfterDiagramUpdateCallback.OnAfterUpdate(stage, diagram, frontDiagram)
	}
}

func (diagram *Diagram) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramDeleteCallback != nil {
		var frontDiagram *Diagram
		if front != nil {
			frontDiagram, _ = front.(*Diagram)
		}
		stage.OnAfterDiagramDeleteCallback.OnAfterDelete(stage, diagram, frontDiagram)
	}
}

func (influence *Influence) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInfluenceCreateCallback != nil {
		stage.OnAfterInfluenceCreateCallback.OnAfterCreate(stage, influence)
	}
}

func (influence *Influence) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInfluenceUpdateCallback != nil {
		var frontInfluence *Influence
		if front != nil {
			frontInfluence, _ = front.(*Influence)
		}
		stage.OnAfterInfluenceUpdateCallback.OnAfterUpdate(stage, influence, frontInfluence)
	}
}

func (influence *Influence) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInfluenceDeleteCallback != nil {
		var frontInfluence *Influence
		if front != nil {
			frontInfluence, _ = front.(*Influence)
		}
		stage.OnAfterInfluenceDeleteCallback.OnAfterDelete(stage, influence, frontInfluence)
	}
}

func (influenceshape *InfluenceShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterInfluenceShapeCreateCallback != nil {
		stage.OnAfterInfluenceShapeCreateCallback.OnAfterCreate(stage, influenceshape)
	}
}

func (influenceshape *InfluenceShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInfluenceShapeUpdateCallback != nil {
		var frontInfluenceShape *InfluenceShape
		if front != nil {
			frontInfluenceShape, _ = front.(*InfluenceShape)
		}
		stage.OnAfterInfluenceShapeUpdateCallback.OnAfterUpdate(stage, influenceshape, frontInfluenceShape)
	}
}

func (influenceshape *InfluenceShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterInfluenceShapeDeleteCallback != nil {
		var frontInfluenceShape *InfluenceShape
		if front != nil {
			frontInfluenceShape, _ = front.(*InfluenceShape)
		}
		stage.OnAfterInfluenceShapeDeleteCallback.OnAfterDelete(stage, influenceshape, frontInfluenceShape)
	}
}

func (library *Library) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLibraryCreateCallback != nil {
		stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, library)
	}
}

func (library *Library) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLibraryUpdateCallback != nil {
		var frontLibrary *Library
		if front != nil {
			frontLibrary, _ = front.(*Library)
		}
		stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, library, frontLibrary)
	}
}

func (library *Library) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLibraryDeleteCallback != nil {
		var frontLibrary *Library
		if front != nil {
			frontLibrary, _ = front.(*Library)
		}
		stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, library, frontLibrary)
	}
}

func (movement *Movement) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMovementCreateCallback != nil {
		stage.OnAfterMovementCreateCallback.OnAfterCreate(stage, movement)
	}
}

func (movement *Movement) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMovementUpdateCallback != nil {
		var frontMovement *Movement
		if front != nil {
			frontMovement, _ = front.(*Movement)
		}
		stage.OnAfterMovementUpdateCallback.OnAfterUpdate(stage, movement, frontMovement)
	}
}

func (movement *Movement) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMovementDeleteCallback != nil {
		var frontMovement *Movement
		if front != nil {
			frontMovement, _ = front.(*Movement)
		}
		stage.OnAfterMovementDeleteCallback.OnAfterDelete(stage, movement, frontMovement)
	}
}

func (movementshape *MovementShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterMovementShapeCreateCallback != nil {
		stage.OnAfterMovementShapeCreateCallback.OnAfterCreate(stage, movementshape)
	}
}

func (movementshape *MovementShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMovementShapeUpdateCallback != nil {
		var frontMovementShape *MovementShape
		if front != nil {
			frontMovementShape, _ = front.(*MovementShape)
		}
		stage.OnAfterMovementShapeUpdateCallback.OnAfterUpdate(stage, movementshape, frontMovementShape)
	}
}

func (movementshape *MovementShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterMovementShapeDeleteCallback != nil {
		var frontMovementShape *MovementShape
		if front != nil {
			frontMovementShape, _ = front.(*MovementShape)
		}
		stage.OnAfterMovementShapeDeleteCallback.OnAfterDelete(stage, movementshape, frontMovementShape)
	}
}

func (place *Place) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPlaceCreateCallback != nil {
		stage.OnAfterPlaceCreateCallback.OnAfterCreate(stage, place)
	}
}

func (place *Place) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlaceUpdateCallback != nil {
		var frontPlace *Place
		if front != nil {
			frontPlace, _ = front.(*Place)
		}
		stage.OnAfterPlaceUpdateCallback.OnAfterUpdate(stage, place, frontPlace)
	}
}

func (place *Place) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPlaceDeleteCallback != nil {
		var frontPlace *Place
		if front != nil {
			frontPlace, _ = front.(*Place)
		}
		stage.OnAfterPlaceDeleteCallback.OnAfterDelete(stage, place, frontPlace)
	}
}

