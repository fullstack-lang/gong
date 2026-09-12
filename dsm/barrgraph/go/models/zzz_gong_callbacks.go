// generated code - do not edit
package models

// AfterCreateFromFront is the Stage method called after a create from front.
func (stage *Stage) AfterCreateFromFront[Type Gongstruct](instance *Type) {

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
	case *Library:
		if stage.OnAfterLibraryCreateCallback != nil {
			stage.OnAfterLibraryCreateCallback.OnAfterCreate(stage, target)
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

// AfterCreateFromFront is a backward-compatible package-level forwarder.
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {
	stage.AfterCreateFromFront(instance)
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is the Stage method called after an update from front.
func (stage *Stage) OnAfterUpdateFromFront[Type Gongstruct](old, new *Type) {

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
	case *Library:
		newTarget := any(new).(*Library)
		if stage.OnAfterLibraryUpdateCallback != nil {
			stage.OnAfterLibraryUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
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

// OnAfterUpdateFromFront is a backward-compatible package-level forwarder.
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {
	stage.OnAfterUpdateFromFront(old, new)
}

// AfterDeleteFromFront is the Stage method called after a delete from front.
func (stage *Stage) AfterDeleteFromFront[Type Gongstruct](staged, front *Type) {

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
	case *Library:
		if stage.OnAfterLibraryDeleteCallback != nil {
			staged := any(staged).(*Library)
			stage.OnAfterLibraryDeleteCallback.OnAfterDelete(stage, staged, front)
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

// AfterDeleteFromFront is a backward-compatible package-level forwarder.
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {
	stage.AfterDeleteFromFront(staged, front)
}
