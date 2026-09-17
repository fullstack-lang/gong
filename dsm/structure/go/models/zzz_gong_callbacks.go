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
func (allocatedresourceshape *AllocatedResourceShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAllocatedResourceShapeCreateCallback != nil {
		stage.OnAfterAllocatedResourceShapeCreateCallback.OnAfterCreate(stage, allocatedresourceshape)
	}
}

func (allocatedresourceshape *AllocatedResourceShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAllocatedResourceShapeUpdateCallback != nil {
		var frontAllocatedResourceShape *AllocatedResourceShape
		if front != nil {
			frontAllocatedResourceShape, _ = front.(*AllocatedResourceShape)
		}
		stage.OnAfterAllocatedResourceShapeUpdateCallback.OnAfterUpdate(stage, allocatedresourceshape, frontAllocatedResourceShape)
	}
}

func (allocatedresourceshape *AllocatedResourceShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAllocatedResourceShapeDeleteCallback != nil {
		var frontAllocatedResourceShape *AllocatedResourceShape
		if front != nil {
			frontAllocatedResourceShape, _ = front.(*AllocatedResourceShape)
		}
		stage.OnAfterAllocatedResourceShapeDeleteCallback.OnAfterDelete(stage, allocatedresourceshape, frontAllocatedResourceShape)
	}
}

func (allocatedsystemshape *AllocatedSystemShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterAllocatedSystemShapeCreateCallback != nil {
		stage.OnAfterAllocatedSystemShapeCreateCallback.OnAfterCreate(stage, allocatedsystemshape)
	}
}

func (allocatedsystemshape *AllocatedSystemShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAllocatedSystemShapeUpdateCallback != nil {
		var frontAllocatedSystemShape *AllocatedSystemShape
		if front != nil {
			frontAllocatedSystemShape, _ = front.(*AllocatedSystemShape)
		}
		stage.OnAfterAllocatedSystemShapeUpdateCallback.OnAfterUpdate(stage, allocatedsystemshape, frontAllocatedSystemShape)
	}
}

func (allocatedsystemshape *AllocatedSystemShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterAllocatedSystemShapeDeleteCallback != nil {
		var frontAllocatedSystemShape *AllocatedSystemShape
		if front != nil {
			frontAllocatedSystemShape, _ = front.(*AllocatedSystemShape)
		}
		stage.OnAfterAllocatedSystemShapeDeleteCallback.OnAfterDelete(stage, allocatedsystemshape, frontAllocatedSystemShape)
	}
}

func (controlflow *ControlFlow) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterControlFlowCreateCallback != nil {
		stage.OnAfterControlFlowCreateCallback.OnAfterCreate(stage, controlflow)
	}
}

func (controlflow *ControlFlow) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterControlFlowUpdateCallback != nil {
		var frontControlFlow *ControlFlow
		if front != nil {
			frontControlFlow, _ = front.(*ControlFlow)
		}
		stage.OnAfterControlFlowUpdateCallback.OnAfterUpdate(stage, controlflow, frontControlFlow)
	}
}

func (controlflow *ControlFlow) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterControlFlowDeleteCallback != nil {
		var frontControlFlow *ControlFlow
		if front != nil {
			frontControlFlow, _ = front.(*ControlFlow)
		}
		stage.OnAfterControlFlowDeleteCallback.OnAfterDelete(stage, controlflow, frontControlFlow)
	}
}

func (controlflowshape *ControlFlowShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterControlFlowShapeCreateCallback != nil {
		stage.OnAfterControlFlowShapeCreateCallback.OnAfterCreate(stage, controlflowshape)
	}
}

func (controlflowshape *ControlFlowShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterControlFlowShapeUpdateCallback != nil {
		var frontControlFlowShape *ControlFlowShape
		if front != nil {
			frontControlFlowShape, _ = front.(*ControlFlowShape)
		}
		stage.OnAfterControlFlowShapeUpdateCallback.OnAfterUpdate(stage, controlflowshape, frontControlFlowShape)
	}
}

func (controlflowshape *ControlFlowShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterControlFlowShapeDeleteCallback != nil {
		var frontControlFlowShape *ControlFlowShape
		if front != nil {
			frontControlFlowShape, _ = front.(*ControlFlowShape)
		}
		stage.OnAfterControlFlowShapeDeleteCallback.OnAfterDelete(stage, controlflowshape, frontControlFlowShape)
	}
}

func (data *Data) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDataCreateCallback != nil {
		stage.OnAfterDataCreateCallback.OnAfterCreate(stage, data)
	}
}

func (data *Data) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataUpdateCallback != nil {
		var frontData *Data
		if front != nil {
			frontData, _ = front.(*Data)
		}
		stage.OnAfterDataUpdateCallback.OnAfterUpdate(stage, data, frontData)
	}
}

func (data *Data) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataDeleteCallback != nil {
		var frontData *Data
		if front != nil {
			frontData, _ = front.(*Data)
		}
		stage.OnAfterDataDeleteCallback.OnAfterDelete(stage, data, frontData)
	}
}

func (dataflow *DataFlow) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDataFlowCreateCallback != nil {
		stage.OnAfterDataFlowCreateCallback.OnAfterCreate(stage, dataflow)
	}
}

func (dataflow *DataFlow) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataFlowUpdateCallback != nil {
		var frontDataFlow *DataFlow
		if front != nil {
			frontDataFlow, _ = front.(*DataFlow)
		}
		stage.OnAfterDataFlowUpdateCallback.OnAfterUpdate(stage, dataflow, frontDataFlow)
	}
}

func (dataflow *DataFlow) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataFlowDeleteCallback != nil {
		var frontDataFlow *DataFlow
		if front != nil {
			frontDataFlow, _ = front.(*DataFlow)
		}
		stage.OnAfterDataFlowDeleteCallback.OnAfterDelete(stage, dataflow, frontDataFlow)
	}
}

func (dataflowshape *DataFlowShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDataFlowShapeCreateCallback != nil {
		stage.OnAfterDataFlowShapeCreateCallback.OnAfterCreate(stage, dataflowshape)
	}
}

func (dataflowshape *DataFlowShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataFlowShapeUpdateCallback != nil {
		var frontDataFlowShape *DataFlowShape
		if front != nil {
			frontDataFlowShape, _ = front.(*DataFlowShape)
		}
		stage.OnAfterDataFlowShapeUpdateCallback.OnAfterUpdate(stage, dataflowshape, frontDataFlowShape)
	}
}

func (dataflowshape *DataFlowShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataFlowShapeDeleteCallback != nil {
		var frontDataFlowShape *DataFlowShape
		if front != nil {
			frontDataFlowShape, _ = front.(*DataFlowShape)
		}
		stage.OnAfterDataFlowShapeDeleteCallback.OnAfterDelete(stage, dataflowshape, frontDataFlowShape)
	}
}

func (datashape *DataShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDataShapeCreateCallback != nil {
		stage.OnAfterDataShapeCreateCallback.OnAfterCreate(stage, datashape)
	}
}

func (datashape *DataShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataShapeUpdateCallback != nil {
		var frontDataShape *DataShape
		if front != nil {
			frontDataShape, _ = front.(*DataShape)
		}
		stage.OnAfterDataShapeUpdateCallback.OnAfterUpdate(stage, datashape, frontDataShape)
	}
}

func (datashape *DataShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDataShapeDeleteCallback != nil {
		var frontDataShape *DataShape
		if front != nil {
			frontDataShape, _ = front.(*DataShape)
		}
		stage.OnAfterDataShapeDeleteCallback.OnAfterDelete(stage, datashape, frontDataShape)
	}
}

func (diagramlayerstate *DiagramLayerState) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDiagramLayerStateCreateCallback != nil {
		stage.OnAfterDiagramLayerStateCreateCallback.OnAfterCreate(stage, diagramlayerstate)
	}
}

func (diagramlayerstate *DiagramLayerState) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramLayerStateUpdateCallback != nil {
		var frontDiagramLayerState *DiagramLayerState
		if front != nil {
			frontDiagramLayerState, _ = front.(*DiagramLayerState)
		}
		stage.OnAfterDiagramLayerStateUpdateCallback.OnAfterUpdate(stage, diagramlayerstate, frontDiagramLayerState)
	}
}

func (diagramlayerstate *DiagramLayerState) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramLayerStateDeleteCallback != nil {
		var frontDiagramLayerState *DiagramLayerState
		if front != nil {
			frontDiagramLayerState, _ = front.(*DiagramLayerState)
		}
		stage.OnAfterDiagramLayerStateDeleteCallback.OnAfterDelete(stage, diagramlayerstate, frontDiagramLayerState)
	}
}

func (diagramstructure *DiagramStructure) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterDiagramStructureCreateCallback != nil {
		stage.OnAfterDiagramStructureCreateCallback.OnAfterCreate(stage, diagramstructure)
	}
}

func (diagramstructure *DiagramStructure) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramStructureUpdateCallback != nil {
		var frontDiagramStructure *DiagramStructure
		if front != nil {
			frontDiagramStructure, _ = front.(*DiagramStructure)
		}
		stage.OnAfterDiagramStructureUpdateCallback.OnAfterUpdate(stage, diagramstructure, frontDiagramStructure)
	}
}

func (diagramstructure *DiagramStructure) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterDiagramStructureDeleteCallback != nil {
		var frontDiagramStructure *DiagramStructure
		if front != nil {
			frontDiagramStructure, _ = front.(*DiagramStructure)
		}
		stage.OnAfterDiagramStructureDeleteCallback.OnAfterDelete(stage, diagramstructure, frontDiagramStructure)
	}
}

func (externalpartshape *ExternalPartShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterExternalPartShapeCreateCallback != nil {
		stage.OnAfterExternalPartShapeCreateCallback.OnAfterCreate(stage, externalpartshape)
	}
}

func (externalpartshape *ExternalPartShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterExternalPartShapeUpdateCallback != nil {
		var frontExternalPartShape *ExternalPartShape
		if front != nil {
			frontExternalPartShape, _ = front.(*ExternalPartShape)
		}
		stage.OnAfterExternalPartShapeUpdateCallback.OnAfterUpdate(stage, externalpartshape, frontExternalPartShape)
	}
}

func (externalpartshape *ExternalPartShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterExternalPartShapeDeleteCallback != nil {
		var frontExternalPartShape *ExternalPartShape
		if front != nil {
			frontExternalPartShape, _ = front.(*ExternalPartShape)
		}
		stage.OnAfterExternalPartShapeDeleteCallback.OnAfterDelete(stage, externalpartshape, frontExternalPartShape)
	}
}

func (layerdefinition *LayerDefinition) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterLayerDefinitionCreateCallback != nil {
		stage.OnAfterLayerDefinitionCreateCallback.OnAfterCreate(stage, layerdefinition)
	}
}

func (layerdefinition *LayerDefinition) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLayerDefinitionUpdateCallback != nil {
		var frontLayerDefinition *LayerDefinition
		if front != nil {
			frontLayerDefinition, _ = front.(*LayerDefinition)
		}
		stage.OnAfterLayerDefinitionUpdateCallback.OnAfterUpdate(stage, layerdefinition, frontLayerDefinition)
	}
}

func (layerdefinition *LayerDefinition) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterLayerDefinitionDeleteCallback != nil {
		var frontLayerDefinition *LayerDefinition
		if front != nil {
			frontLayerDefinition, _ = front.(*LayerDefinition)
		}
		stage.OnAfterLayerDefinitionDeleteCallback.OnAfterDelete(stage, layerdefinition, frontLayerDefinition)
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

func (notepartshape *NotePartShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNotePartShapeCreateCallback != nil {
		stage.OnAfterNotePartShapeCreateCallback.OnAfterCreate(stage, notepartshape)
	}
}

func (notepartshape *NotePartShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNotePartShapeUpdateCallback != nil {
		var frontNotePartShape *NotePartShape
		if front != nil {
			frontNotePartShape, _ = front.(*NotePartShape)
		}
		stage.OnAfterNotePartShapeUpdateCallback.OnAfterUpdate(stage, notepartshape, frontNotePartShape)
	}
}

func (notepartshape *NotePartShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNotePartShapeDeleteCallback != nil {
		var frontNotePartShape *NotePartShape
		if front != nil {
			frontNotePartShape, _ = front.(*NotePartShape)
		}
		stage.OnAfterNotePartShapeDeleteCallback.OnAfterDelete(stage, notepartshape, frontNotePartShape)
	}
}

func (noteportshape *NotePortShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterNotePortShapeCreateCallback != nil {
		stage.OnAfterNotePortShapeCreateCallback.OnAfterCreate(stage, noteportshape)
	}
}

func (noteportshape *NotePortShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNotePortShapeUpdateCallback != nil {
		var frontNotePortShape *NotePortShape
		if front != nil {
			frontNotePortShape, _ = front.(*NotePortShape)
		}
		stage.OnAfterNotePortShapeUpdateCallback.OnAfterUpdate(stage, noteportshape, frontNotePortShape)
	}
}

func (noteportshape *NotePortShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterNotePortShapeDeleteCallback != nil {
		var frontNotePortShape *NotePortShape
		if front != nil {
			frontNotePortShape, _ = front.(*NotePortShape)
		}
		stage.OnAfterNotePortShapeDeleteCallback.OnAfterDelete(stage, noteportshape, frontNotePortShape)
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

func (part *Part) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartCreateCallback != nil {
		stage.OnAfterPartCreateCallback.OnAfterCreate(stage, part)
	}
}

func (part *Part) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartUpdateCallback != nil {
		var frontPart *Part
		if front != nil {
			frontPart, _ = front.(*Part)
		}
		stage.OnAfterPartUpdateCallback.OnAfterUpdate(stage, part, frontPart)
	}
}

func (part *Part) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartDeleteCallback != nil {
		var frontPart *Part
		if front != nil {
			frontPart, _ = front.(*Part)
		}
		stage.OnAfterPartDeleteCallback.OnAfterDelete(stage, part, frontPart)
	}
}

func (partanchoredpath *PartAnchoredPath) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartAnchoredPathCreateCallback != nil {
		stage.OnAfterPartAnchoredPathCreateCallback.OnAfterCreate(stage, partanchoredpath)
	}
}

func (partanchoredpath *PartAnchoredPath) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartAnchoredPathUpdateCallback != nil {
		var frontPartAnchoredPath *PartAnchoredPath
		if front != nil {
			frontPartAnchoredPath, _ = front.(*PartAnchoredPath)
		}
		stage.OnAfterPartAnchoredPathUpdateCallback.OnAfterUpdate(stage, partanchoredpath, frontPartAnchoredPath)
	}
}

func (partanchoredpath *PartAnchoredPath) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartAnchoredPathDeleteCallback != nil {
		var frontPartAnchoredPath *PartAnchoredPath
		if front != nil {
			frontPartAnchoredPath, _ = front.(*PartAnchoredPath)
		}
		stage.OnAfterPartAnchoredPathDeleteCallback.OnAfterDelete(stage, partanchoredpath, frontPartAnchoredPath)
	}
}

func (partshape *PartShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPartShapeCreateCallback != nil {
		stage.OnAfterPartShapeCreateCallback.OnAfterCreate(stage, partshape)
	}
}

func (partshape *PartShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartShapeUpdateCallback != nil {
		var frontPartShape *PartShape
		if front != nil {
			frontPartShape, _ = front.(*PartShape)
		}
		stage.OnAfterPartShapeUpdateCallback.OnAfterUpdate(stage, partshape, frontPartShape)
	}
}

func (partshape *PartShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPartShapeDeleteCallback != nil {
		var frontPartShape *PartShape
		if front != nil {
			frontPartShape, _ = front.(*PartShape)
		}
		stage.OnAfterPartShapeDeleteCallback.OnAfterDelete(stage, partshape, frontPartShape)
	}
}

func (port *Port) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPortCreateCallback != nil {
		stage.OnAfterPortCreateCallback.OnAfterCreate(stage, port)
	}
}

func (port *Port) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPortUpdateCallback != nil {
		var frontPort *Port
		if front != nil {
			frontPort, _ = front.(*Port)
		}
		stage.OnAfterPortUpdateCallback.OnAfterUpdate(stage, port, frontPort)
	}
}

func (port *Port) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPortDeleteCallback != nil {
		var frontPort *Port
		if front != nil {
			frontPort, _ = front.(*Port)
		}
		stage.OnAfterPortDeleteCallback.OnAfterDelete(stage, port, frontPort)
	}
}

func (portshape *PortShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterPortShapeCreateCallback != nil {
		stage.OnAfterPortShapeCreateCallback.OnAfterCreate(stage, portshape)
	}
}

func (portshape *PortShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPortShapeUpdateCallback != nil {
		var frontPortShape *PortShape
		if front != nil {
			frontPortShape, _ = front.(*PortShape)
		}
		stage.OnAfterPortShapeUpdateCallback.OnAfterUpdate(stage, portshape, frontPortShape)
	}
}

func (portshape *PortShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterPortShapeDeleteCallback != nil {
		var frontPortShape *PortShape
		if front != nil {
			frontPortShape, _ = front.(*PortShape)
		}
		stage.OnAfterPortShapeDeleteCallback.OnAfterDelete(stage, portshape, frontPortShape)
	}
}

func (resource *Resource) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterResourceCreateCallback != nil {
		stage.OnAfterResourceCreateCallback.OnAfterCreate(stage, resource)
	}
}

func (resource *Resource) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterResourceUpdateCallback != nil {
		var frontResource *Resource
		if front != nil {
			frontResource, _ = front.(*Resource)
		}
		stage.OnAfterResourceUpdateCallback.OnAfterUpdate(stage, resource, frontResource)
	}
}

func (resource *Resource) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterResourceDeleteCallback != nil {
		var frontResource *Resource
		if front != nil {
			frontResource, _ = front.(*Resource)
		}
		stage.OnAfterResourceDeleteCallback.OnAfterDelete(stage, resource, frontResource)
	}
}

func (semantictag *SemanticTag) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSemanticTagCreateCallback != nil {
		stage.OnAfterSemanticTagCreateCallback.OnAfterCreate(stage, semantictag)
	}
}

func (semantictag *SemanticTag) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSemanticTagUpdateCallback != nil {
		var frontSemanticTag *SemanticTag
		if front != nil {
			frontSemanticTag, _ = front.(*SemanticTag)
		}
		stage.OnAfterSemanticTagUpdateCallback.OnAfterUpdate(stage, semantictag, frontSemanticTag)
	}
}

func (semantictag *SemanticTag) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSemanticTagDeleteCallback != nil {
		var frontSemanticTag *SemanticTag
		if front != nil {
			frontSemanticTag, _ = front.(*SemanticTag)
		}
		stage.OnAfterSemanticTagDeleteCallback.OnAfterDelete(stage, semantictag, frontSemanticTag)
	}
}

func (system *System) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSystemCreateCallback != nil {
		stage.OnAfterSystemCreateCallback.OnAfterCreate(stage, system)
	}
}

func (system *System) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystemUpdateCallback != nil {
		var frontSystem *System
		if front != nil {
			frontSystem, _ = front.(*System)
		}
		stage.OnAfterSystemUpdateCallback.OnAfterUpdate(stage, system, frontSystem)
	}
}

func (system *System) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystemDeleteCallback != nil {
		var frontSystem *System
		if front != nil {
			frontSystem, _ = front.(*System)
		}
		stage.OnAfterSystemDeleteCallback.OnAfterDelete(stage, system, frontSystem)
	}
}

func (systemshape *SystemShape) GongAfterCreateFromFront(stage *Stage) {
	if stage.OnAfterSystemShapeCreateCallback != nil {
		stage.OnAfterSystemShapeCreateCallback.OnAfterCreate(stage, systemshape)
	}
}

func (systemshape *SystemShape) GongOnAfterUpdateFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystemShapeUpdateCallback != nil {
		var frontSystemShape *SystemShape
		if front != nil {
			frontSystemShape, _ = front.(*SystemShape)
		}
		stage.OnAfterSystemShapeUpdateCallback.OnAfterUpdate(stage, systemshape, frontSystemShape)
	}
}

func (systemshape *SystemShape) GongAfterDeleteFromFront(stage *Stage, front GongstructIF) {
	if stage.OnAfterSystemShapeDeleteCallback != nil {
		var frontSystemShape *SystemShape
		if front != nil {
			frontSystemShape, _ = front.(*SystemShape)
		}
		stage.OnAfterSystemShapeDeleteCallback.OnAfterDelete(stage, systemshape, frontSystemShape)
	}
}

