// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *ArtefactType:
		if stage.OnAfterArtefactTypeCreateCallback != nil {
			stage.OnAfterArtefactTypeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ArtefactTypeShape:
		if stage.OnAfterArtefactTypeShapeCreateCallback != nil {
			stage.OnAfterArtefactTypeShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Artist:
		if stage.OnAfterArtistCreateCallback != nil {
			stage.OnAfterArtistCreateCallback.OnAfterCreate(stage, target)
		}
	case *ArtistShape:
		if stage.OnAfterArtistShapeCreateCallback != nil {
			stage.OnAfterArtistShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *ControlPointShape:
		if stage.OnAfterControlPointShapeCreateCallback != nil {
			stage.OnAfterControlPointShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Desk:
		if stage.OnAfterDeskCreateCallback != nil {
			stage.OnAfterDeskCreateCallback.OnAfterCreate(stage, target)
		}
	case *Diagram:
		if stage.OnAfterDiagramCreateCallback != nil {
			stage.OnAfterDiagramCreateCallback.OnAfterCreate(stage, target)
		}
	case *Influence:
		if stage.OnAfterInfluenceCreateCallback != nil {
			stage.OnAfterInfluenceCreateCallback.OnAfterCreate(stage, target)
		}
	case *InfluenceShape:
		if stage.OnAfterInfluenceShapeCreateCallback != nil {
			stage.OnAfterInfluenceShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Movement:
		if stage.OnAfterMovementCreateCallback != nil {
			stage.OnAfterMovementCreateCallback.OnAfterCreate(stage, target)
		}
	case *MovementShape:
		if stage.OnAfterMovementShapeCreateCallback != nil {
			stage.OnAfterMovementShapeCreateCallback.OnAfterCreate(stage, target)
		}
	case *Place:
		if stage.OnAfterPlaceCreateCallback != nil {
			stage.OnAfterPlaceCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *ArtefactType:
		newTarget := any(new).(*ArtefactType)
		if stage.OnAfterArtefactTypeUpdateCallback != nil {
			stage.OnAfterArtefactTypeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ArtefactTypeShape:
		newTarget := any(new).(*ArtefactTypeShape)
		if stage.OnAfterArtefactTypeShapeUpdateCallback != nil {
			stage.OnAfterArtefactTypeShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Artist:
		newTarget := any(new).(*Artist)
		if stage.OnAfterArtistUpdateCallback != nil {
			stage.OnAfterArtistUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ArtistShape:
		newTarget := any(new).(*ArtistShape)
		if stage.OnAfterArtistShapeUpdateCallback != nil {
			stage.OnAfterArtistShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ControlPointShape:
		newTarget := any(new).(*ControlPointShape)
		if stage.OnAfterControlPointShapeUpdateCallback != nil {
			stage.OnAfterControlPointShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Desk:
		newTarget := any(new).(*Desk)
		if stage.OnAfterDeskUpdateCallback != nil {
			stage.OnAfterDeskUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Diagram:
		newTarget := any(new).(*Diagram)
		if stage.OnAfterDiagramUpdateCallback != nil {
			stage.OnAfterDiagramUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Influence:
		newTarget := any(new).(*Influence)
		if stage.OnAfterInfluenceUpdateCallback != nil {
			stage.OnAfterInfluenceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *InfluenceShape:
		newTarget := any(new).(*InfluenceShape)
		if stage.OnAfterInfluenceShapeUpdateCallback != nil {
			stage.OnAfterInfluenceShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Movement:
		newTarget := any(new).(*Movement)
		if stage.OnAfterMovementUpdateCallback != nil {
			stage.OnAfterMovementUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *MovementShape:
		newTarget := any(new).(*MovementShape)
		if stage.OnAfterMovementShapeUpdateCallback != nil {
			stage.OnAfterMovementShapeUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *Place:
		newTarget := any(new).(*Place)
		if stage.OnAfterPlaceUpdateCallback != nil {
			stage.OnAfterPlaceUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *ArtefactType:
		if stage.OnAfterArtefactTypeDeleteCallback != nil {
			staged := any(staged).(*ArtefactType)
			stage.OnAfterArtefactTypeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ArtefactTypeShape:
		if stage.OnAfterArtefactTypeShapeDeleteCallback != nil {
			staged := any(staged).(*ArtefactTypeShape)
			stage.OnAfterArtefactTypeShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Artist:
		if stage.OnAfterArtistDeleteCallback != nil {
			staged := any(staged).(*Artist)
			stage.OnAfterArtistDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ArtistShape:
		if stage.OnAfterArtistShapeDeleteCallback != nil {
			staged := any(staged).(*ArtistShape)
			stage.OnAfterArtistShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ControlPointShape:
		if stage.OnAfterControlPointShapeDeleteCallback != nil {
			staged := any(staged).(*ControlPointShape)
			stage.OnAfterControlPointShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Desk:
		if stage.OnAfterDeskDeleteCallback != nil {
			staged := any(staged).(*Desk)
			stage.OnAfterDeskDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Diagram:
		if stage.OnAfterDiagramDeleteCallback != nil {
			staged := any(staged).(*Diagram)
			stage.OnAfterDiagramDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Influence:
		if stage.OnAfterInfluenceDeleteCallback != nil {
			staged := any(staged).(*Influence)
			stage.OnAfterInfluenceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *InfluenceShape:
		if stage.OnAfterInfluenceShapeDeleteCallback != nil {
			staged := any(staged).(*InfluenceShape)
			stage.OnAfterInfluenceShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Movement:
		if stage.OnAfterMovementDeleteCallback != nil {
			staged := any(staged).(*Movement)
			stage.OnAfterMovementDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *MovementShape:
		if stage.OnAfterMovementShapeDeleteCallback != nil {
			staged := any(staged).(*MovementShape)
			stage.OnAfterMovementShapeDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *Place:
		if stage.OnAfterPlaceDeleteCallback != nil {
			staged := any(staged).(*Place)
			stage.OnAfterPlaceDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *ArtefactType:
		if stage.OnAfterArtefactTypeReadCallback != nil {
			stage.OnAfterArtefactTypeReadCallback.OnAfterRead(stage, target)
		}
	case *ArtefactTypeShape:
		if stage.OnAfterArtefactTypeShapeReadCallback != nil {
			stage.OnAfterArtefactTypeShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Artist:
		if stage.OnAfterArtistReadCallback != nil {
			stage.OnAfterArtistReadCallback.OnAfterRead(stage, target)
		}
	case *ArtistShape:
		if stage.OnAfterArtistShapeReadCallback != nil {
			stage.OnAfterArtistShapeReadCallback.OnAfterRead(stage, target)
		}
	case *ControlPointShape:
		if stage.OnAfterControlPointShapeReadCallback != nil {
			stage.OnAfterControlPointShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Desk:
		if stage.OnAfterDeskReadCallback != nil {
			stage.OnAfterDeskReadCallback.OnAfterRead(stage, target)
		}
	case *Diagram:
		if stage.OnAfterDiagramReadCallback != nil {
			stage.OnAfterDiagramReadCallback.OnAfterRead(stage, target)
		}
	case *Influence:
		if stage.OnAfterInfluenceReadCallback != nil {
			stage.OnAfterInfluenceReadCallback.OnAfterRead(stage, target)
		}
	case *InfluenceShape:
		if stage.OnAfterInfluenceShapeReadCallback != nil {
			stage.OnAfterInfluenceShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Movement:
		if stage.OnAfterMovementReadCallback != nil {
			stage.OnAfterMovementReadCallback.OnAfterRead(stage, target)
		}
	case *MovementShape:
		if stage.OnAfterMovementShapeReadCallback != nil {
			stage.OnAfterMovementShapeReadCallback.OnAfterRead(stage, target)
		}
	case *Place:
		if stage.OnAfterPlaceReadCallback != nil {
			stage.OnAfterPlaceReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ArtefactType:
		stage.OnAfterArtefactTypeUpdateCallback = any(callback).(OnAfterUpdateInterface[ArtefactType])
	
	case *ArtefactTypeShape:
		stage.OnAfterArtefactTypeShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ArtefactTypeShape])
	
	case *Artist:
		stage.OnAfterArtistUpdateCallback = any(callback).(OnAfterUpdateInterface[Artist])
	
	case *ArtistShape:
		stage.OnAfterArtistShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ArtistShape])
	
	case *ControlPointShape:
		stage.OnAfterControlPointShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[ControlPointShape])
	
	case *Desk:
		stage.OnAfterDeskUpdateCallback = any(callback).(OnAfterUpdateInterface[Desk])
	
	case *Diagram:
		stage.OnAfterDiagramUpdateCallback = any(callback).(OnAfterUpdateInterface[Diagram])
	
	case *Influence:
		stage.OnAfterInfluenceUpdateCallback = any(callback).(OnAfterUpdateInterface[Influence])
	
	case *InfluenceShape:
		stage.OnAfterInfluenceShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[InfluenceShape])
	
	case *Movement:
		stage.OnAfterMovementUpdateCallback = any(callback).(OnAfterUpdateInterface[Movement])
	
	case *MovementShape:
		stage.OnAfterMovementShapeUpdateCallback = any(callback).(OnAfterUpdateInterface[MovementShape])
	
	case *Place:
		stage.OnAfterPlaceUpdateCallback = any(callback).(OnAfterUpdateInterface[Place])
	
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ArtefactType:
		stage.OnAfterArtefactTypeCreateCallback = any(callback).(OnAfterCreateInterface[ArtefactType])
	
	case *ArtefactTypeShape:
		stage.OnAfterArtefactTypeShapeCreateCallback = any(callback).(OnAfterCreateInterface[ArtefactTypeShape])
	
	case *Artist:
		stage.OnAfterArtistCreateCallback = any(callback).(OnAfterCreateInterface[Artist])
	
	case *ArtistShape:
		stage.OnAfterArtistShapeCreateCallback = any(callback).(OnAfterCreateInterface[ArtistShape])
	
	case *ControlPointShape:
		stage.OnAfterControlPointShapeCreateCallback = any(callback).(OnAfterCreateInterface[ControlPointShape])
	
	case *Desk:
		stage.OnAfterDeskCreateCallback = any(callback).(OnAfterCreateInterface[Desk])
	
	case *Diagram:
		stage.OnAfterDiagramCreateCallback = any(callback).(OnAfterCreateInterface[Diagram])
	
	case *Influence:
		stage.OnAfterInfluenceCreateCallback = any(callback).(OnAfterCreateInterface[Influence])
	
	case *InfluenceShape:
		stage.OnAfterInfluenceShapeCreateCallback = any(callback).(OnAfterCreateInterface[InfluenceShape])
	
	case *Movement:
		stage.OnAfterMovementCreateCallback = any(callback).(OnAfterCreateInterface[Movement])
	
	case *MovementShape:
		stage.OnAfterMovementShapeCreateCallback = any(callback).(OnAfterCreateInterface[MovementShape])
	
	case *Place:
		stage.OnAfterPlaceCreateCallback = any(callback).(OnAfterCreateInterface[Place])
	
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ArtefactType:
		stage.OnAfterArtefactTypeDeleteCallback = any(callback).(OnAfterDeleteInterface[ArtefactType])
	
	case *ArtefactTypeShape:
		stage.OnAfterArtefactTypeShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ArtefactTypeShape])
	
	case *Artist:
		stage.OnAfterArtistDeleteCallback = any(callback).(OnAfterDeleteInterface[Artist])
	
	case *ArtistShape:
		stage.OnAfterArtistShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ArtistShape])
	
	case *ControlPointShape:
		stage.OnAfterControlPointShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[ControlPointShape])
	
	case *Desk:
		stage.OnAfterDeskDeleteCallback = any(callback).(OnAfterDeleteInterface[Desk])
	
	case *Diagram:
		stage.OnAfterDiagramDeleteCallback = any(callback).(OnAfterDeleteInterface[Diagram])
	
	case *Influence:
		stage.OnAfterInfluenceDeleteCallback = any(callback).(OnAfterDeleteInterface[Influence])
	
	case *InfluenceShape:
		stage.OnAfterInfluenceShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[InfluenceShape])
	
	case *Movement:
		stage.OnAfterMovementDeleteCallback = any(callback).(OnAfterDeleteInterface[Movement])
	
	case *MovementShape:
		stage.OnAfterMovementShapeDeleteCallback = any(callback).(OnAfterDeleteInterface[MovementShape])
	
	case *Place:
		stage.OnAfterPlaceDeleteCallback = any(callback).(OnAfterDeleteInterface[Place])
	
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ArtefactType:
		stage.OnAfterArtefactTypeReadCallback = any(callback).(OnAfterReadInterface[ArtefactType])
	
	case *ArtefactTypeShape:
		stage.OnAfterArtefactTypeShapeReadCallback = any(callback).(OnAfterReadInterface[ArtefactTypeShape])
	
	case *Artist:
		stage.OnAfterArtistReadCallback = any(callback).(OnAfterReadInterface[Artist])
	
	case *ArtistShape:
		stage.OnAfterArtistShapeReadCallback = any(callback).(OnAfterReadInterface[ArtistShape])
	
	case *ControlPointShape:
		stage.OnAfterControlPointShapeReadCallback = any(callback).(OnAfterReadInterface[ControlPointShape])
	
	case *Desk:
		stage.OnAfterDeskReadCallback = any(callback).(OnAfterReadInterface[Desk])
	
	case *Diagram:
		stage.OnAfterDiagramReadCallback = any(callback).(OnAfterReadInterface[Diagram])
	
	case *Influence:
		stage.OnAfterInfluenceReadCallback = any(callback).(OnAfterReadInterface[Influence])
	
	case *InfluenceShape:
		stage.OnAfterInfluenceShapeReadCallback = any(callback).(OnAfterReadInterface[InfluenceShape])
	
	case *Movement:
		stage.OnAfterMovementReadCallback = any(callback).(OnAfterReadInterface[Movement])
	
	case *MovementShape:
		stage.OnAfterMovementShapeReadCallback = any(callback).(OnAfterReadInterface[MovementShape])
	
	case *Place:
		stage.OnAfterPlaceReadCallback = any(callback).(OnAfterReadInterface[Place])
	
	}
}
