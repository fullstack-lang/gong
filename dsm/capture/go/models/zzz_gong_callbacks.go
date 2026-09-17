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
func (analysisneed *AnalysisNeed) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAnalysisNeedCreateCallback != nil {
		stage.OnAfterAnalysisNeedCreateCallback.OnAfterCreate(stage, analysisneed)
	}
}

func (analysisneed *AnalysisNeed) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAnalysisNeedUpdateCallback != nil {
		var frontAnalysisNeed *AnalysisNeed
		if front != nil {
			frontAnalysisNeed, _ = front.(*AnalysisNeed)
		}
		stage.OnAfterAnalysisNeedUpdateCallback.OnAfterUpdate(stage, analysisneed, frontAnalysisNeed)
	}
}

func (analysisneed *AnalysisNeed) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAnalysisNeedDeleteCallback != nil {
		var frontAnalysisNeed *AnalysisNeed
		if front != nil {
			frontAnalysisNeed, _ = front.(*AnalysisNeed)
		}
		stage.OnAfterAnalysisNeedDeleteCallback.OnAfterDelete(stage, analysisneed, frontAnalysisNeed)
	}
}

func (concept *Concept) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConceptCreateCallback != nil {
		stage.OnAfterConceptCreateCallback.OnAfterCreate(stage, concept)
	}
}

func (concept *Concept) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConceptUpdateCallback != nil {
		var frontConcept *Concept
		if front != nil {
			frontConcept, _ = front.(*Concept)
		}
		stage.OnAfterConceptUpdateCallback.OnAfterUpdate(stage, concept, frontConcept)
	}
}

func (concept *Concept) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConceptDeleteCallback != nil {
		var frontConcept *Concept
		if front != nil {
			frontConcept, _ = front.(*Concept)
		}
		stage.OnAfterConceptDeleteCallback.OnAfterDelete(stage, concept, frontConcept)
	}
}

func (conceptshape *ConceptShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConceptShapeCreateCallback != nil {
		stage.OnAfterConceptShapeCreateCallback.OnAfterCreate(stage, conceptshape)
	}
}

func (conceptshape *ConceptShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConceptShapeUpdateCallback != nil {
		var frontConceptShape *ConceptShape
		if front != nil {
			frontConceptShape, _ = front.(*ConceptShape)
		}
		stage.OnAfterConceptShapeUpdateCallback.OnAfterUpdate(stage, conceptshape, frontConceptShape)
	}
}

func (conceptshape *ConceptShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConceptShapeDeleteCallback != nil {
		var frontConceptShape *ConceptShape
		if front != nil {
			frontConceptShape, _ = front.(*ConceptShape)
		}
		stage.OnAfterConceptShapeDeleteCallback.OnAfterDelete(stage, conceptshape, frontConceptShape)
	}
}

func (concern *Concern) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConcernCreateCallback != nil {
		stage.OnAfterConcernCreateCallback.OnAfterCreate(stage, concern)
	}
}

func (concern *Concern) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConcernUpdateCallback != nil {
		var frontConcern *Concern
		if front != nil {
			frontConcern, _ = front.(*Concern)
		}
		stage.OnAfterConcernUpdateCallback.OnAfterUpdate(stage, concern, frontConcern)
	}
}

func (concern *Concern) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConcernDeleteCallback != nil {
		var frontConcern *Concern
		if front != nil {
			frontConcern, _ = front.(*Concern)
		}
		stage.OnAfterConcernDeleteCallback.OnAfterDelete(stage, concern, frontConcern)
	}
}

func (concerncompositionshape *ConcernCompositionShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConcernCompositionShapeCreateCallback != nil {
		stage.OnAfterConcernCompositionShapeCreateCallback.OnAfterCreate(stage, concerncompositionshape)
	}
}

func (concerncompositionshape *ConcernCompositionShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConcernCompositionShapeUpdateCallback != nil {
		var frontConcernCompositionShape *ConcernCompositionShape
		if front != nil {
			frontConcernCompositionShape, _ = front.(*ConcernCompositionShape)
		}
		stage.OnAfterConcernCompositionShapeUpdateCallback.OnAfterUpdate(stage, concerncompositionshape, frontConcernCompositionShape)
	}
}

func (concerncompositionshape *ConcernCompositionShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConcernCompositionShapeDeleteCallback != nil {
		var frontConcernCompositionShape *ConcernCompositionShape
		if front != nil {
			frontConcernCompositionShape, _ = front.(*ConcernCompositionShape)
		}
		stage.OnAfterConcernCompositionShapeDeleteCallback.OnAfterDelete(stage, concerncompositionshape, frontConcernCompositionShape)
	}
}

func (concerninputshape *ConcernInputShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConcernInputShapeCreateCallback != nil {
		stage.OnAfterConcernInputShapeCreateCallback.OnAfterCreate(stage, concerninputshape)
	}
}

func (concerninputshape *ConcernInputShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConcernInputShapeUpdateCallback != nil {
		var frontConcernInputShape *ConcernInputShape
		if front != nil {
			frontConcernInputShape, _ = front.(*ConcernInputShape)
		}
		stage.OnAfterConcernInputShapeUpdateCallback.OnAfterUpdate(stage, concerninputshape, frontConcernInputShape)
	}
}

func (concerninputshape *ConcernInputShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConcernInputShapeDeleteCallback != nil {
		var frontConcernInputShape *ConcernInputShape
		if front != nil {
			frontConcernInputShape, _ = front.(*ConcernInputShape)
		}
		stage.OnAfterConcernInputShapeDeleteCallback.OnAfterDelete(stage, concerninputshape, frontConcernInputShape)
	}
}

func (concernoutputshape *ConcernOutputShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConcernOutputShapeCreateCallback != nil {
		stage.OnAfterConcernOutputShapeCreateCallback.OnAfterCreate(stage, concernoutputshape)
	}
}

func (concernoutputshape *ConcernOutputShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConcernOutputShapeUpdateCallback != nil {
		var frontConcernOutputShape *ConcernOutputShape
		if front != nil {
			frontConcernOutputShape, _ = front.(*ConcernOutputShape)
		}
		stage.OnAfterConcernOutputShapeUpdateCallback.OnAfterUpdate(stage, concernoutputshape, frontConcernOutputShape)
	}
}

func (concernoutputshape *ConcernOutputShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConcernOutputShapeDeleteCallback != nil {
		var frontConcernOutputShape *ConcernOutputShape
		if front != nil {
			frontConcernOutputShape, _ = front.(*ConcernOutputShape)
		}
		stage.OnAfterConcernOutputShapeDeleteCallback.OnAfterDelete(stage, concernoutputshape, frontConcernOutputShape)
	}
}

func (concernshape *ConcernShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterConcernShapeCreateCallback != nil {
		stage.OnAfterConcernShapeCreateCallback.OnAfterCreate(stage, concernshape)
	}
}

func (concernshape *ConcernShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConcernShapeUpdateCallback != nil {
		var frontConcernShape *ConcernShape
		if front != nil {
			frontConcernShape, _ = front.(*ConcernShape)
		}
		stage.OnAfterConcernShapeUpdateCallback.OnAfterUpdate(stage, concernshape, frontConcernShape)
	}
}

func (concernshape *ConcernShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterConcernShapeDeleteCallback != nil {
		var frontConcernShape *ConcernShape
		if front != nil {
			frontConcernShape, _ = front.(*ConcernShape)
		}
		stage.OnAfterConcernShapeDeleteCallback.OnAfterDelete(stage, concernshape, frontConcernShape)
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

func (deliverable *Deliverable) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDeliverableCreateCallback != nil {
		stage.OnAfterDeliverableCreateCallback.OnAfterCreate(stage, deliverable)
	}
}

func (deliverable *Deliverable) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeliverableUpdateCallback != nil {
		var frontDeliverable *Deliverable
		if front != nil {
			frontDeliverable, _ = front.(*Deliverable)
		}
		stage.OnAfterDeliverableUpdateCallback.OnAfterUpdate(stage, deliverable, frontDeliverable)
	}
}

func (deliverable *Deliverable) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeliverableDeleteCallback != nil {
		var frontDeliverable *Deliverable
		if front != nil {
			frontDeliverable, _ = front.(*Deliverable)
		}
		stage.OnAfterDeliverableDeleteCallback.OnAfterDelete(stage, deliverable, frontDeliverable)
	}
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDeliverableCompositionShapeCreateCallback != nil {
		stage.OnAfterDeliverableCompositionShapeCreateCallback.OnAfterCreate(stage, deliverablecompositionshape)
	}
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeliverableCompositionShapeUpdateCallback != nil {
		var frontDeliverableCompositionShape *DeliverableCompositionShape
		if front != nil {
			frontDeliverableCompositionShape, _ = front.(*DeliverableCompositionShape)
		}
		stage.OnAfterDeliverableCompositionShapeUpdateCallback.OnAfterUpdate(stage, deliverablecompositionshape, frontDeliverableCompositionShape)
	}
}

func (deliverablecompositionshape *DeliverableCompositionShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeliverableCompositionShapeDeleteCallback != nil {
		var frontDeliverableCompositionShape *DeliverableCompositionShape
		if front != nil {
			frontDeliverableCompositionShape, _ = front.(*DeliverableCompositionShape)
		}
		stage.OnAfterDeliverableCompositionShapeDeleteCallback.OnAfterDelete(stage, deliverablecompositionshape, frontDeliverableCompositionShape)
	}
}

func (deliverableconceptshape *DeliverableConceptShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDeliverableConceptShapeCreateCallback != nil {
		stage.OnAfterDeliverableConceptShapeCreateCallback.OnAfterCreate(stage, deliverableconceptshape)
	}
}

func (deliverableconceptshape *DeliverableConceptShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeliverableConceptShapeUpdateCallback != nil {
		var frontDeliverableConceptShape *DeliverableConceptShape
		if front != nil {
			frontDeliverableConceptShape, _ = front.(*DeliverableConceptShape)
		}
		stage.OnAfterDeliverableConceptShapeUpdateCallback.OnAfterUpdate(stage, deliverableconceptshape, frontDeliverableConceptShape)
	}
}

func (deliverableconceptshape *DeliverableConceptShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeliverableConceptShapeDeleteCallback != nil {
		var frontDeliverableConceptShape *DeliverableConceptShape
		if front != nil {
			frontDeliverableConceptShape, _ = front.(*DeliverableConceptShape)
		}
		stage.OnAfterDeliverableConceptShapeDeleteCallback.OnAfterDelete(stage, deliverableconceptshape, frontDeliverableConceptShape)
	}
}

func (deliverableshape *DeliverableShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDeliverableShapeCreateCallback != nil {
		stage.OnAfterDeliverableShapeCreateCallback.OnAfterCreate(stage, deliverableshape)
	}
}

func (deliverableshape *DeliverableShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeliverableShapeUpdateCallback != nil {
		var frontDeliverableShape *DeliverableShape
		if front != nil {
			frontDeliverableShape, _ = front.(*DeliverableShape)
		}
		stage.OnAfterDeliverableShapeUpdateCallback.OnAfterUpdate(stage, deliverableshape, frontDeliverableShape)
	}
}

func (deliverableshape *DeliverableShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDeliverableShapeDeleteCallback != nil {
		var frontDeliverableShape *DeliverableShape
		if front != nil {
			frontDeliverableShape, _ = front.(*DeliverableShape)
		}
		stage.OnAfterDeliverableShapeDeleteCallback.OnAfterDelete(stage, deliverableshape, frontDeliverableShape)
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

func (diagramshape *DiagramShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDiagramShapeCreateCallback != nil {
		stage.OnAfterDiagramShapeCreateCallback.OnAfterCreate(stage, diagramshape)
	}
}

func (diagramshape *DiagramShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramShapeUpdateCallback != nil {
		var frontDiagramShape *DiagramShape
		if front != nil {
			frontDiagramShape, _ = front.(*DiagramShape)
		}
		stage.OnAfterDiagramShapeUpdateCallback.OnAfterUpdate(stage, diagramshape, frontDiagramShape)
	}
}

func (diagramshape *DiagramShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramShapeDeleteCallback != nil {
		var frontDiagramShape *DiagramShape
		if front != nil {
			frontDiagramShape, _ = front.(*DiagramShape)
		}
		stage.OnAfterDiagramShapeDeleteCallback.OnAfterDelete(stage, diagramshape, frontDiagramShape)
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

func (note *Note) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteCreateCallback != nil {
		stage.OnAfterNoteCreateCallback.OnAfterCreate(stage, note)
	}
}

func (note *Note) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteUpdateCallback != nil {
		var frontNote *Note
		if front != nil {
			frontNote, _ = front.(*Note)
		}
		stage.OnAfterNoteUpdateCallback.OnAfterUpdate(stage, note, frontNote)
	}
}

func (note *Note) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteDeleteCallback != nil {
		var frontNote *Note
		if front != nil {
			frontNote, _ = front.(*Note)
		}
		stage.OnAfterNoteDeleteCallback.OnAfterDelete(stage, note, frontNote)
	}
}

func (notedeliverableshape *NoteDeliverableShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteDeliverableShapeCreateCallback != nil {
		stage.OnAfterNoteDeliverableShapeCreateCallback.OnAfterCreate(stage, notedeliverableshape)
	}
}

func (notedeliverableshape *NoteDeliverableShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteDeliverableShapeUpdateCallback != nil {
		var frontNoteDeliverableShape *NoteDeliverableShape
		if front != nil {
			frontNoteDeliverableShape, _ = front.(*NoteDeliverableShape)
		}
		stage.OnAfterNoteDeliverableShapeUpdateCallback.OnAfterUpdate(stage, notedeliverableshape, frontNoteDeliverableShape)
	}
}

func (notedeliverableshape *NoteDeliverableShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteDeliverableShapeDeleteCallback != nil {
		var frontNoteDeliverableShape *NoteDeliverableShape
		if front != nil {
			frontNoteDeliverableShape, _ = front.(*NoteDeliverableShape)
		}
		stage.OnAfterNoteDeliverableShapeDeleteCallback.OnAfterDelete(stage, notedeliverableshape, frontNoteDeliverableShape)
	}
}

func (noteshape *NoteShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteShapeCreateCallback != nil {
		stage.OnAfterNoteShapeCreateCallback.OnAfterCreate(stage, noteshape)
	}
}

func (noteshape *NoteShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteShapeUpdateCallback != nil {
		var frontNoteShape *NoteShape
		if front != nil {
			frontNoteShape, _ = front.(*NoteShape)
		}
		stage.OnAfterNoteShapeUpdateCallback.OnAfterUpdate(stage, noteshape, frontNoteShape)
	}
}

func (noteshape *NoteShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteShapeDeleteCallback != nil {
		var frontNoteShape *NoteShape
		if front != nil {
			frontNoteShape, _ = front.(*NoteShape)
		}
		stage.OnAfterNoteShapeDeleteCallback.OnAfterDelete(stage, noteshape, frontNoteShape)
	}
}

func (notestakeholdershape *NoteStakeholderShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteStakeholderShapeCreateCallback != nil {
		stage.OnAfterNoteStakeholderShapeCreateCallback.OnAfterCreate(stage, notestakeholdershape)
	}
}

func (notestakeholdershape *NoteStakeholderShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteStakeholderShapeUpdateCallback != nil {
		var frontNoteStakeholderShape *NoteStakeholderShape
		if front != nil {
			frontNoteStakeholderShape, _ = front.(*NoteStakeholderShape)
		}
		stage.OnAfterNoteStakeholderShapeUpdateCallback.OnAfterUpdate(stage, notestakeholdershape, frontNoteStakeholderShape)
	}
}

func (notestakeholdershape *NoteStakeholderShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteStakeholderShapeDeleteCallback != nil {
		var frontNoteStakeholderShape *NoteStakeholderShape
		if front != nil {
			frontNoteStakeholderShape, _ = front.(*NoteStakeholderShape)
		}
		stage.OnAfterNoteStakeholderShapeDeleteCallback.OnAfterDelete(stage, notestakeholdershape, frontNoteStakeholderShape)
	}
}

func (notetaskshape *NoteTaskShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNoteTaskShapeCreateCallback != nil {
		stage.OnAfterNoteTaskShapeCreateCallback.OnAfterCreate(stage, notetaskshape)
	}
}

func (notetaskshape *NoteTaskShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteTaskShapeUpdateCallback != nil {
		var frontNoteTaskShape *NoteTaskShape
		if front != nil {
			frontNoteTaskShape, _ = front.(*NoteTaskShape)
		}
		stage.OnAfterNoteTaskShapeUpdateCallback.OnAfterUpdate(stage, notetaskshape, frontNoteTaskShape)
	}
}

func (notetaskshape *NoteTaskShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNoteTaskShapeDeleteCallback != nil {
		var frontNoteTaskShape *NoteTaskShape
		if front != nil {
			frontNoteTaskShape, _ = front.(*NoteTaskShape)
		}
		stage.OnAfterNoteTaskShapeDeleteCallback.OnAfterDelete(stage, notetaskshape, frontNoteTaskShape)
	}
}

func (requirement *Requirement) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRequirementCreateCallback != nil {
		stage.OnAfterRequirementCreateCallback.OnAfterCreate(stage, requirement)
	}
}

func (requirement *Requirement) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRequirementUpdateCallback != nil {
		var frontRequirement *Requirement
		if front != nil {
			frontRequirement, _ = front.(*Requirement)
		}
		stage.OnAfterRequirementUpdateCallback.OnAfterUpdate(stage, requirement, frontRequirement)
	}
}

func (requirement *Requirement) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRequirementDeleteCallback != nil {
		var frontRequirement *Requirement
		if front != nil {
			frontRequirement, _ = front.(*Requirement)
		}
		stage.OnAfterRequirementDeleteCallback.OnAfterDelete(stage, requirement, frontRequirement)
	}
}

func (requirementshape *RequirementShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterRequirementShapeCreateCallback != nil {
		stage.OnAfterRequirementShapeCreateCallback.OnAfterCreate(stage, requirementshape)
	}
}

func (requirementshape *RequirementShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRequirementShapeUpdateCallback != nil {
		var frontRequirementShape *RequirementShape
		if front != nil {
			frontRequirementShape, _ = front.(*RequirementShape)
		}
		stage.OnAfterRequirementShapeUpdateCallback.OnAfterUpdate(stage, requirementshape, frontRequirementShape)
	}
}

func (requirementshape *RequirementShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterRequirementShapeDeleteCallback != nil {
		var frontRequirementShape *RequirementShape
		if front != nil {
			frontRequirementShape, _ = front.(*RequirementShape)
		}
		stage.OnAfterRequirementShapeDeleteCallback.OnAfterDelete(stage, requirementshape, frontRequirementShape)
	}
}

func (stakeholder *Stakeholder) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStakeholderCreateCallback != nil {
		stage.OnAfterStakeholderCreateCallback.OnAfterCreate(stage, stakeholder)
	}
}

func (stakeholder *Stakeholder) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStakeholderUpdateCallback != nil {
		var frontStakeholder *Stakeholder
		if front != nil {
			frontStakeholder, _ = front.(*Stakeholder)
		}
		stage.OnAfterStakeholderUpdateCallback.OnAfterUpdate(stage, stakeholder, frontStakeholder)
	}
}

func (stakeholder *Stakeholder) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStakeholderDeleteCallback != nil {
		var frontStakeholder *Stakeholder
		if front != nil {
			frontStakeholder, _ = front.(*Stakeholder)
		}
		stage.OnAfterStakeholderDeleteCallback.OnAfterDelete(stage, stakeholder, frontStakeholder)
	}
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStakeholderCompositionShapeCreateCallback != nil {
		stage.OnAfterStakeholderCompositionShapeCreateCallback.OnAfterCreate(stage, stakeholdercompositionshape)
	}
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStakeholderCompositionShapeUpdateCallback != nil {
		var frontStakeholderCompositionShape *StakeholderCompositionShape
		if front != nil {
			frontStakeholderCompositionShape, _ = front.(*StakeholderCompositionShape)
		}
		stage.OnAfterStakeholderCompositionShapeUpdateCallback.OnAfterUpdate(stage, stakeholdercompositionshape, frontStakeholderCompositionShape)
	}
}

func (stakeholdercompositionshape *StakeholderCompositionShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStakeholderCompositionShapeDeleteCallback != nil {
		var frontStakeholderCompositionShape *StakeholderCompositionShape
		if front != nil {
			frontStakeholderCompositionShape, _ = front.(*StakeholderCompositionShape)
		}
		stage.OnAfterStakeholderCompositionShapeDeleteCallback.OnAfterDelete(stage, stakeholdercompositionshape, frontStakeholderCompositionShape)
	}
}

func (stakeholderconcernshape *StakeholderConcernShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStakeholderConcernShapeCreateCallback != nil {
		stage.OnAfterStakeholderConcernShapeCreateCallback.OnAfterCreate(stage, stakeholderconcernshape)
	}
}

func (stakeholderconcernshape *StakeholderConcernShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStakeholderConcernShapeUpdateCallback != nil {
		var frontStakeholderConcernShape *StakeholderConcernShape
		if front != nil {
			frontStakeholderConcernShape, _ = front.(*StakeholderConcernShape)
		}
		stage.OnAfterStakeholderConcernShapeUpdateCallback.OnAfterUpdate(stage, stakeholderconcernshape, frontStakeholderConcernShape)
	}
}

func (stakeholderconcernshape *StakeholderConcernShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStakeholderConcernShapeDeleteCallback != nil {
		var frontStakeholderConcernShape *StakeholderConcernShape
		if front != nil {
			frontStakeholderConcernShape, _ = front.(*StakeholderConcernShape)
		}
		stage.OnAfterStakeholderConcernShapeDeleteCallback.OnAfterDelete(stage, stakeholderconcernshape, frontStakeholderConcernShape)
	}
}

func (stakeholdershape *StakeholderShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterStakeholderShapeCreateCallback != nil {
		stage.OnAfterStakeholderShapeCreateCallback.OnAfterCreate(stage, stakeholdershape)
	}
}

func (stakeholdershape *StakeholderShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStakeholderShapeUpdateCallback != nil {
		var frontStakeholderShape *StakeholderShape
		if front != nil {
			frontStakeholderShape, _ = front.(*StakeholderShape)
		}
		stage.OnAfterStakeholderShapeUpdateCallback.OnAfterUpdate(stage, stakeholdershape, frontStakeholderShape)
	}
}

func (stakeholdershape *StakeholderShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterStakeholderShapeDeleteCallback != nil {
		var frontStakeholderShape *StakeholderShape
		if front != nil {
			frontStakeholderShape, _ = front.(*StakeholderShape)
		}
		stage.OnAfterStakeholderShapeDeleteCallback.OnAfterDelete(stage, stakeholdershape, frontStakeholderShape)
	}
}

func (supportlevel *SupportLevel) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSupportLevelCreateCallback != nil {
		stage.OnAfterSupportLevelCreateCallback.OnAfterCreate(stage, supportlevel)
	}
}

func (supportlevel *SupportLevel) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSupportLevelUpdateCallback != nil {
		var frontSupportLevel *SupportLevel
		if front != nil {
			frontSupportLevel, _ = front.(*SupportLevel)
		}
		stage.OnAfterSupportLevelUpdateCallback.OnAfterUpdate(stage, supportlevel, frontSupportLevel)
	}
}

func (supportlevel *SupportLevel) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSupportLevelDeleteCallback != nil {
		var frontSupportLevel *SupportLevel
		if front != nil {
			frontSupportLevel, _ = front.(*SupportLevel)
		}
		stage.OnAfterSupportLevelDeleteCallback.OnAfterDelete(stage, supportlevel, frontSupportLevel)
	}
}

func (tool *Tool) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterToolCreateCallback != nil {
		stage.OnAfterToolCreateCallback.OnAfterCreate(stage, tool)
	}
}

func (tool *Tool) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterToolUpdateCallback != nil {
		var frontTool *Tool
		if front != nil {
			frontTool, _ = front.(*Tool)
		}
		stage.OnAfterToolUpdateCallback.OnAfterUpdate(stage, tool, frontTool)
	}
}

func (tool *Tool) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterToolDeleteCallback != nil {
		var frontTool *Tool
		if front != nil {
			frontTool, _ = front.(*Tool)
		}
		stage.OnAfterToolDeleteCallback.OnAfterDelete(stage, tool, frontTool)
	}
}

