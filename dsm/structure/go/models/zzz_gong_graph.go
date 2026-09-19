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
func (allocatedresourceshape *AllocatedResourceShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.AllocatedResourceShapes[allocatedresourceshape]

	return
}

func (stage *Stage) IsStagedAllocatedResourceShape(allocatedresourceshape *AllocatedResourceShape) (ok bool) {

	return allocatedresourceshape.GongIsStaged(stage)
}

func (allocatedsystemshape *AllocatedSystemShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.AllocatedSystemShapes[allocatedsystemshape]

	return
}

func (stage *Stage) IsStagedAllocatedSystemShape(allocatedsystemshape *AllocatedSystemShape) (ok bool) {

	return allocatedsystemshape.GongIsStaged(stage)
}

func (controlflow *ControlFlow) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ControlFlows[controlflow]

	return
}

func (stage *Stage) IsStagedControlFlow(controlflow *ControlFlow) (ok bool) {

	return controlflow.GongIsStaged(stage)
}

func (controlflowshape *ControlFlowShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ControlFlowShapes[controlflowshape]

	return
}

func (stage *Stage) IsStagedControlFlowShape(controlflowshape *ControlFlowShape) (ok bool) {

	return controlflowshape.GongIsStaged(stage)
}

func (data *Data) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Datas[data]

	return
}

func (stage *Stage) IsStagedData(data *Data) (ok bool) {

	return data.GongIsStaged(stage)
}

func (dataflow *DataFlow) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DataFlows[dataflow]

	return
}

func (stage *Stage) IsStagedDataFlow(dataflow *DataFlow) (ok bool) {

	return dataflow.GongIsStaged(stage)
}

func (dataflowshape *DataFlowShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DataFlowShapes[dataflowshape]

	return
}

func (stage *Stage) IsStagedDataFlowShape(dataflowshape *DataFlowShape) (ok bool) {

	return dataflowshape.GongIsStaged(stage)
}

func (datashape *DataShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DataShapes[datashape]

	return
}

func (stage *Stage) IsStagedDataShape(datashape *DataShape) (ok bool) {

	return datashape.GongIsStaged(stage)
}

func (diagramlayerstate *DiagramLayerState) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DiagramLayerStates[diagramlayerstate]

	return
}

func (stage *Stage) IsStagedDiagramLayerState(diagramlayerstate *DiagramLayerState) (ok bool) {

	return diagramlayerstate.GongIsStaged(stage)
}

func (diagramstructure *DiagramStructure) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DiagramStructures[diagramstructure]

	return
}

func (stage *Stage) IsStagedDiagramStructure(diagramstructure *DiagramStructure) (ok bool) {

	return diagramstructure.GongIsStaged(stage)
}

func (externalpartshape *ExternalPartShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ExternalPartShapes[externalpartshape]

	return
}

func (stage *Stage) IsStagedExternalPartShape(externalpartshape *ExternalPartShape) (ok bool) {

	return externalpartshape.GongIsStaged(stage)
}

func (layerdefinition *LayerDefinition) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.LayerDefinitions[layerdefinition]

	return
}

func (stage *Stage) IsStagedLayerDefinition(layerdefinition *LayerDefinition) (ok bool) {

	return layerdefinition.GongIsStaged(stage)
}

func (library *Library) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	return library.GongIsStaged(stage)
}

func (note *Note) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Notes[note]

	return
}

func (stage *Stage) IsStagedNote(note *Note) (ok bool) {

	return note.GongIsStaged(stage)
}

func (notepartshape *NotePartShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.NotePartShapes[notepartshape]

	return
}

func (stage *Stage) IsStagedNotePartShape(notepartshape *NotePartShape) (ok bool) {

	return notepartshape.GongIsStaged(stage)
}

func (noteportshape *NotePortShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.NotePortShapes[noteportshape]

	return
}

func (stage *Stage) IsStagedNotePortShape(noteportshape *NotePortShape) (ok bool) {

	return noteportshape.GongIsStaged(stage)
}

func (noteshape *NoteShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.NoteShapes[noteshape]

	return
}

func (stage *Stage) IsStagedNoteShape(noteshape *NoteShape) (ok bool) {

	return noteshape.GongIsStaged(stage)
}

func (part *Part) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Parts[part]

	return
}

func (stage *Stage) IsStagedPart(part *Part) (ok bool) {

	return part.GongIsStaged(stage)
}

func (partanchoredpath *PartAnchoredPath) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PartAnchoredPaths[partanchoredpath]

	return
}

func (stage *Stage) IsStagedPartAnchoredPath(partanchoredpath *PartAnchoredPath) (ok bool) {

	return partanchoredpath.GongIsStaged(stage)
}

func (partshape *PartShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PartShapes[partshape]

	return
}

func (stage *Stage) IsStagedPartShape(partshape *PartShape) (ok bool) {

	return partshape.GongIsStaged(stage)
}

func (port *Port) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Ports[port]

	return
}

func (stage *Stage) IsStagedPort(port *Port) (ok bool) {

	return port.GongIsStaged(stage)
}

func (portshape *PortShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PortShapes[portshape]

	return
}

func (stage *Stage) IsStagedPortShape(portshape *PortShape) (ok bool) {

	return portshape.GongIsStaged(stage)
}

func (resource *Resource) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Resources[resource]

	return
}

func (stage *Stage) IsStagedResource(resource *Resource) (ok bool) {

	return resource.GongIsStaged(stage)
}

func (semantictag *SemanticTag) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SemanticTags[semantictag]

	return
}

func (stage *Stage) IsStagedSemanticTag(semantictag *SemanticTag) (ok bool) {

	return semantictag.GongIsStaged(stage)
}

func (system *System) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.Systems[system]

	return
}

func (stage *Stage) IsStagedSystem(system *System) (ok bool) {

	return system.GongIsStaged(stage)
}

func (systemshape *SystemShape) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SystemShapes[systemshape]

	return
}

func (stage *Stage) IsStagedSystemShape(systemshape *SystemShape) (ok bool) {

	return systemshape.GongIsStaged(stage)
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
func (allocatedresourceshape *AllocatedResourceShape) GongStageBranch(stage *Stage) {
	stage.StageBranchAllocatedResourceShape(allocatedresourceshape)
}

func (stage *Stage) StageBranchAllocatedResourceShape(allocatedresourceshape *AllocatedResourceShape) {

	// check if instance is already staged
	if stage.IsStaged(allocatedresourceshape) {
		return
	}

	allocatedresourceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if allocatedresourceshape.Part != nil {
		stage.StageBranch(allocatedresourceshape.Part)
	}
	if allocatedresourceshape.Resource != nil {
		stage.StageBranch(allocatedresourceshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (allocatedsystemshape *AllocatedSystemShape) GongStageBranch(stage *Stage) {
	stage.StageBranchAllocatedSystemShape(allocatedsystemshape)
}

func (stage *Stage) StageBranchAllocatedSystemShape(allocatedsystemshape *AllocatedSystemShape) {

	// check if instance is already staged
	if stage.IsStaged(allocatedsystemshape) {
		return
	}

	allocatedsystemshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if allocatedsystemshape.Part != nil {
		stage.StageBranch(allocatedsystemshape.Part)
	}
	if allocatedsystemshape.System != nil {
		stage.StageBranch(allocatedsystemshape.System)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (controlflow *ControlFlow) GongStageBranch(stage *Stage) {
	stage.StageBranchControlFlow(controlflow)
}

func (stage *Stage) StageBranchControlFlow(controlflow *ControlFlow) {

	// check if instance is already staged
	if stage.IsStaged(controlflow) {
		return
	}

	controlflow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlflow.Start != nil {
		stage.StageBranch(controlflow.Start)
	}
	if controlflow.End != nil {
		stage.StageBranch(controlflow.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (controlflowshape *ControlFlowShape) GongStageBranch(stage *Stage) {
	stage.StageBranchControlFlowShape(controlflowshape)
}

func (stage *Stage) StageBranchControlFlowShape(controlflowshape *ControlFlowShape) {

	// check if instance is already staged
	if stage.IsStaged(controlflowshape) {
		return
	}

	controlflowshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlflowshape.ControlFlow != nil {
		stage.StageBranch(controlflowshape.ControlFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (data *Data) GongStageBranch(stage *Stage) {
	stage.StageBranchData(data)
}

func (stage *Stage) StageBranchData(data *Data) {

	// check if instance is already staged
	if stage.IsStaged(data) {
		return
	}

	data.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dataflow *DataFlow) GongStageBranch(stage *Stage) {
	stage.StageBranchDataFlow(dataflow)
}

func (stage *Stage) StageBranchDataFlow(dataflow *DataFlow) {

	// check if instance is already staged
	if stage.IsStaged(dataflow) {
		return
	}

	dataflow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dataflow.StartPort != nil {
		stage.StageBranch(dataflow.StartPort)
	}
	if dataflow.EndPort != nil {
		stage.StageBranch(dataflow.EndPort)
	}
	if dataflow.StartExternalPart != nil {
		stage.StageBranch(dataflow.StartExternalPart)
	}
	if dataflow.EndExternalPart != nil {
		stage.StageBranch(dataflow.EndExternalPart)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _data := range dataflow.Datas {
		stage.StageBranch(_data)
	}

}

func (dataflowshape *DataFlowShape) GongStageBranch(stage *Stage) {
	stage.StageBranchDataFlowShape(dataflowshape)
}

func (stage *Stage) StageBranchDataFlowShape(dataflowshape *DataFlowShape) {

	// check if instance is already staged
	if stage.IsStaged(dataflowshape) {
		return
	}

	dataflowshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dataflowshape.DataFlow != nil {
		stage.StageBranch(dataflowshape.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datashape *DataShape) GongStageBranch(stage *Stage) {
	stage.StageBranchDataShape(datashape)
}

func (stage *Stage) StageBranchDataShape(datashape *DataShape) {

	// check if instance is already staged
	if stage.IsStaged(datashape) {
		return
	}

	datashape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datashape.Data != nil {
		stage.StageBranch(datashape.Data)
	}
	if datashape.DataFlow != nil {
		stage.StageBranch(datashape.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagramlayerstate *DiagramLayerState) GongStageBranch(stage *Stage) {
	stage.StageBranchDiagramLayerState(diagramlayerstate)
}

func (stage *Stage) StageBranchDiagramLayerState(diagramlayerstate *DiagramLayerState) {

	// check if instance is already staged
	if stage.IsStaged(diagramlayerstate) {
		return
	}

	diagramlayerstate.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if diagramlayerstate.DiagramStructure != nil {
		stage.StageBranch(diagramlayerstate.DiagramStructure)
	}
	if diagramlayerstate.LayerDefinition != nil {
		stage.StageBranch(diagramlayerstate.LayerDefinition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagramstructure *DiagramStructure) GongStageBranch(stage *Stage) {
	stage.StageBranchDiagramStructure(diagramstructure)
}

func (stage *Stage) StageBranchDiagramStructure(diagramstructure *DiagramStructure) {

	// check if instance is already staged
	if stage.IsStaged(diagramstructure) {
		return
	}

	diagramstructure.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _systemshape := range diagramstructure.System_Shapes {
		stage.StageBranch(_systemshape)
	}
	for _, _system := range diagramstructure.SystemsWhoseNodeIsExpanded {
		stage.StageBranch(_system)
	}
	for _, _partshape := range diagramstructure.Part_Shapes {
		stage.StageBranch(_partshape)
	}
	for _, _part := range diagramstructure.PartWhoseNodeIsExpanded {
		stage.StageBranch(_part)
	}
	for _, _externalpartshape := range diagramstructure.ExternalPart_Shapes {
		stage.StageBranch(_externalpartshape)
	}
	for _, _part := range diagramstructure.ExternalPartWhoseNodeIsExpanded {
		stage.StageBranch(_part)
	}
	for _, _part := range diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
		stage.StageBranch(_part)
	}
	for _, _part := range diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
		stage.StageBranch(_part)
	}
	for _, _port := range diagramstructure.PortsWhoseNodeIsExpanded {
		stage.StageBranch(_port)
	}
	for _, _portshape := range diagramstructure.Port_Shapes {
		stage.StageBranch(_portshape)
	}
	for _, _controlflow := range diagramstructure.ControlFlowsWhoseNodeIsExpanded {
		stage.StageBranch(_controlflow)
	}
	for _, _controlflowshape := range diagramstructure.ControlFlow_Shapes {
		stage.StageBranch(_controlflowshape)
	}
	for _, _dataflow := range diagramstructure.DataFlowsWhoseNodeIsExpanded {
		stage.StageBranch(_dataflow)
	}
	for _, _dataflowshape := range diagramstructure.DataFlow_Shapes {
		stage.StageBranch(_dataflowshape)
	}
	for _, _data := range diagramstructure.DatasWhoseNodeIsExpanded {
		stage.StageBranch(_data)
	}
	for _, _datashape := range diagramstructure.Data_Shapes {
		stage.StageBranch(_datashape)
	}
	for _, _dataflow := range diagramstructure.DataFlowsWhoseDataNodeIsExpanded {
		stage.StageBranch(_dataflow)
	}
	for _, _resource := range diagramstructure.AllocatedResourcesWhoseNodeIsExpanded {
		stage.StageBranch(_resource)
	}
	for _, _allocatedresourceshape := range diagramstructure.AllocatedResourceShapes {
		stage.StageBranch(_allocatedresourceshape)
	}
	for _, _system := range diagramstructure.AllocatedSystemesWhoseNodeIsExpanded {
		stage.StageBranch(_system)
	}
	for _, _allocatedsystemshape := range diagramstructure.AllocatedSystemShapes {
		stage.StageBranch(_allocatedsystemshape)
	}
	for _, _noteshape := range diagramstructure.Note_Shapes {
		stage.StageBranch(_noteshape)
	}
	for _, _note := range diagramstructure.NotesWhoseNodeIsExpanded {
		stage.StageBranch(_note)
	}
	for _, _noteportshape := range diagramstructure.NotePortShapes {
		stage.StageBranch(_noteportshape)
	}
	for _, _notepartshape := range diagramstructure.NotePartShapes {
		stage.StageBranch(_notepartshape)
	}

}

func (externalpartshape *ExternalPartShape) GongStageBranch(stage *Stage) {
	stage.StageBranchExternalPartShape(externalpartshape)
}

func (stage *Stage) StageBranchExternalPartShape(externalpartshape *ExternalPartShape) {

	// check if instance is already staged
	if stage.IsStaged(externalpartshape) {
		return
	}

	externalpartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if externalpartshape.Part != nil {
		stage.StageBranch(externalpartshape.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (layerdefinition *LayerDefinition) GongStageBranch(stage *Stage) {
	stage.StageBranchLayerDefinition(layerdefinition)
}

func (stage *Stage) StageBranchLayerDefinition(layerdefinition *LayerDefinition) {

	// check if instance is already staged
	if stage.IsStaged(layerdefinition) {
		return
	}

	layerdefinition.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _semantictag := range layerdefinition.Query {
		stage.StageBranch(_semantictag)
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
	for _, _system := range library.RootSystemes {
		stage.StageBranch(_system)
	}
	for _, _system := range library.SystemsWhoseNodeIsExpanded {
		stage.StageBranch(_system)
	}
	for _, _dataflow := range library.RootDataFlows {
		stage.StageBranch(_dataflow)
	}
	for _, _dataflow := range library.DataFlowsWhoseNodeIsExpanded {
		stage.StageBranch(_dataflow)
	}
	for _, _data := range library.RootDatas {
		stage.StageBranch(_data)
	}
	for _, _data := range library.DatasWhoseNodeIsExpanded {
		stage.StageBranch(_data)
	}
	for _, _resource := range library.RootResources {
		stage.StageBranch(_resource)
	}
	for _, _resource := range library.ResourcesWhoseNodeIsExpanded {
		stage.StageBranch(_resource)
	}
	for _, _part := range library.PartsWhoseNodeIsExpanded {
		stage.StageBranch(_part)
	}
	for _, _note := range library.RootNotes {
		stage.StageBranch(_note)
	}
	for _, _note := range library.NotesWhoseNodeIsExpanded {
		stage.StageBranch(_note)
	}

}

func (note *Note) GongStageBranch(stage *Stage) {
	stage.StageBranchNote(note)
}

func (stage *Stage) StageBranchNote(note *Note) {

	// check if instance is already staged
	if stage.IsStaged(note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _part := range note.Parts {
		stage.StageBranch(_part)
	}
	for _, _port := range note.Ports {
		stage.StageBranch(_port)
	}

}

func (notepartshape *NotePartShape) GongStageBranch(stage *Stage) {
	stage.StageBranchNotePartShape(notepartshape)
}

func (stage *Stage) StageBranchNotePartShape(notepartshape *NotePartShape) {

	// check if instance is already staged
	if stage.IsStaged(notepartshape) {
		return
	}

	notepartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notepartshape.Note != nil {
		stage.StageBranch(notepartshape.Note)
	}
	if notepartshape.Part != nil {
		stage.StageBranch(notepartshape.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteportshape *NotePortShape) GongStageBranch(stage *Stage) {
	stage.StageBranchNotePortShape(noteportshape)
}

func (stage *Stage) StageBranchNotePortShape(noteportshape *NotePortShape) {

	// check if instance is already staged
	if stage.IsStaged(noteportshape) {
		return
	}

	noteportshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteportshape.Note != nil {
		stage.StageBranch(noteportshape.Note)
	}
	if noteportshape.Port != nil {
		stage.StageBranch(noteportshape.Port)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteshape *NoteShape) GongStageBranch(stage *Stage) {
	stage.StageBranchNoteShape(noteshape)
}

func (stage *Stage) StageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if stage.IsStaged(noteshape) {
		return
	}

	noteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		stage.StageBranch(noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (part *Part) GongStageBranch(stage *Stage) {
	stage.StageBranchPart(part)
}

func (stage *Stage) StageBranchPart(part *Part) {

	// check if instance is already staged
	if stage.IsStaged(part) {
		return
	}

	part.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if part.TypeOfPart != nil {
		stage.StageBranch(part.TypeOfPart)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _port := range part.Ports {
		stage.StageBranch(_port)
	}
	for _, _controlflow := range part.ControlFlows {
		stage.StageBranch(_controlflow)
	}
	for _, _port := range part.PortWhoseOutControlFlowsNodeIsExpanded {
		stage.StageBranch(_port)
	}
	for _, _port := range part.PortWhoseInControlFlowsNodeIsExpanded {
		stage.StageBranch(_port)
	}
	for _, _port := range part.PortWhoseOutDataFlowsNodeIsExpanded {
		stage.StageBranch(_port)
	}
	for _, _port := range part.PortWhoseInDataFlowsNodeIsExpanded {
		stage.StageBranch(_port)
	}
	for _, _partanchoredpath := range part.PartAnchoredPath {
		stage.StageBranch(_partanchoredpath)
	}

}

func (partanchoredpath *PartAnchoredPath) GongStageBranch(stage *Stage) {
	stage.StageBranchPartAnchoredPath(partanchoredpath)
}

func (stage *Stage) StageBranchPartAnchoredPath(partanchoredpath *PartAnchoredPath) {

	// check if instance is already staged
	if stage.IsStaged(partanchoredpath) {
		return
	}

	partanchoredpath.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partshape *PartShape) GongStageBranch(stage *Stage) {
	stage.StageBranchPartShape(partshape)
}

func (stage *Stage) StageBranchPartShape(partshape *PartShape) {

	// check if instance is already staged
	if stage.IsStaged(partshape) {
		return
	}

	partshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if partshape.Part != nil {
		stage.StageBranch(partshape.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (port *Port) GongStageBranch(stage *Stage) {
	stage.StageBranchPort(port)
}

func (stage *Stage) StageBranchPort(port *Port) {

	// check if instance is already staged
	if stage.IsStaged(port) {
		return
	}

	port.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (portshape *PortShape) GongStageBranch(stage *Stage) {
	stage.StageBranchPortShape(portshape)
}

func (stage *Stage) StageBranchPortShape(portshape *PortShape) {

	// check if instance is already staged
	if stage.IsStaged(portshape) {
		return
	}

	portshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if portshape.Port != nil {
		stage.StageBranch(portshape.Port)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (resource *Resource) GongStageBranch(stage *Stage) {
	stage.StageBranchResource(resource)
}

func (stage *Stage) StageBranchResource(resource *Resource) {

	// check if instance is already staged
	if stage.IsStaged(resource) {
		return
	}

	resource.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (semantictag *SemanticTag) GongStageBranch(stage *Stage) {
	stage.StageBranchSemanticTag(semantictag)
}

func (stage *Stage) StageBranchSemanticTag(semantictag *SemanticTag) {

	// check if instance is already staged
	if stage.IsStaged(semantictag) {
		return
	}

	semantictag.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _part := range semantictag.Parts {
		stage.StageBranch(_part)
	}

}

func (system *System) GongStageBranch(stage *Stage) {
	stage.StageBranchSystem(system)
}

func (stage *Stage) StageBranchSystem(system *System) {

	// check if instance is already staged
	if stage.IsStaged(system) {
		return
	}

	system.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagramstructure := range system.DiagramStructures {
		stage.StageBranch(_diagramstructure)
	}
	for _, _diagramstructure := range system.DiagramStructureWhoseNodeIsExpanded {
		stage.StageBranch(_diagramstructure)
	}
	for _, _system := range system.SubSystemes {
		stage.StageBranch(_system)
	}
	for _, _part := range system.Parts {
		stage.StageBranch(_part)
	}
	for _, _part := range system.PartWhoseNodeIsExpanded {
		stage.StageBranch(_part)
	}
	for _, _dataflow := range system.DataFlows {
		stage.StageBranch(_dataflow)
	}
	for _, _part := range system.ExternalParts {
		stage.StageBranch(_part)
	}
	for _, _part := range system.ExternalPartWhoseNodeIsExpanded {
		stage.StageBranch(_part)
	}

}

func (systemshape *SystemShape) GongStageBranch(stage *Stage) {
	stage.StageBranchSystemShape(systemshape)
}

func (stage *Stage) StageBranchSystemShape(systemshape *SystemShape) {

	// check if instance is already staged
	if stage.IsStaged(systemshape) {
		return
	}

	systemshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if systemshape.System != nil {
		stage.StageBranch(systemshape.System)
	}

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
	case *AllocatedResourceShape:
		toT := GongCopyBranchAllocatedResourceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AllocatedSystemShape:
		toT := GongCopyBranchAllocatedSystemShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ControlFlow:
		toT := GongCopyBranchControlFlow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ControlFlowShape:
		toT := GongCopyBranchControlFlowShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Data:
		toT := GongCopyBranchData(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DataFlow:
		toT := GongCopyBranchDataFlow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DataFlowShape:
		toT := GongCopyBranchDataFlowShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DataShape:
		toT := GongCopyBranchDataShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DiagramLayerState:
		toT := GongCopyBranchDiagramLayerState(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DiagramStructure:
		toT := GongCopyBranchDiagramStructure(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ExternalPartShape:
		toT := GongCopyBranchExternalPartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LayerDefinition:
		toT := GongCopyBranchLayerDefinition(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := GongCopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := GongCopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NotePartShape:
		toT := GongCopyBranchNotePartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NotePortShape:
		toT := GongCopyBranchNotePortShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteShape:
		toT := GongCopyBranchNoteShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part:
		toT := GongCopyBranchPart(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartAnchoredPath:
		toT := GongCopyBranchPartAnchoredPath(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartShape:
		toT := GongCopyBranchPartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Port:
		toT := GongCopyBranchPort(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PortShape:
		toT := GongCopyBranchPortShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Resource:
		toT := GongCopyBranchResource(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SemanticTag:
		toT := GongCopyBranchSemanticTag(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *System:
		toT := GongCopyBranchSystem(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SystemShape:
		toT := GongCopyBranchSystemShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAllocatedResourceShape(mapOrigCopy map[any]any, allocatedresourceshapeFrom *AllocatedResourceShape) (allocatedresourceshapeTo *AllocatedResourceShape) {

	// allocatedresourceshapeFrom has already been copied
	if _allocatedresourceshapeTo, ok := mapOrigCopy[allocatedresourceshapeFrom]; ok {
		allocatedresourceshapeTo = _allocatedresourceshapeTo.(*AllocatedResourceShape)
		return
	}

	allocatedresourceshapeTo = new(AllocatedResourceShape)
	mapOrigCopy[allocatedresourceshapeFrom] = allocatedresourceshapeTo
	allocatedresourceshapeFrom.GongCopyBasicFields(allocatedresourceshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if allocatedresourceshapeFrom.Part != nil {
		allocatedresourceshapeTo.Part = GongCopyBranchPart(mapOrigCopy, allocatedresourceshapeFrom.Part)
	}
	if allocatedresourceshapeFrom.Resource != nil {
		allocatedresourceshapeTo.Resource = GongCopyBranchResource(mapOrigCopy, allocatedresourceshapeFrom.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAllocatedSystemShape(mapOrigCopy map[any]any, allocatedsystemshapeFrom *AllocatedSystemShape) (allocatedsystemshapeTo *AllocatedSystemShape) {

	// allocatedsystemshapeFrom has already been copied
	if _allocatedsystemshapeTo, ok := mapOrigCopy[allocatedsystemshapeFrom]; ok {
		allocatedsystemshapeTo = _allocatedsystemshapeTo.(*AllocatedSystemShape)
		return
	}

	allocatedsystemshapeTo = new(AllocatedSystemShape)
	mapOrigCopy[allocatedsystemshapeFrom] = allocatedsystemshapeTo
	allocatedsystemshapeFrom.GongCopyBasicFields(allocatedsystemshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if allocatedsystemshapeFrom.Part != nil {
		allocatedsystemshapeTo.Part = GongCopyBranchPart(mapOrigCopy, allocatedsystemshapeFrom.Part)
	}
	if allocatedsystemshapeFrom.System != nil {
		allocatedsystemshapeTo.System = GongCopyBranchSystem(mapOrigCopy, allocatedsystemshapeFrom.System)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchControlFlow(mapOrigCopy map[any]any, controlflowFrom *ControlFlow) (controlflowTo *ControlFlow) {

	// controlflowFrom has already been copied
	if _controlflowTo, ok := mapOrigCopy[controlflowFrom]; ok {
		controlflowTo = _controlflowTo.(*ControlFlow)
		return
	}

	controlflowTo = new(ControlFlow)
	mapOrigCopy[controlflowFrom] = controlflowTo
	controlflowFrom.GongCopyBasicFields(controlflowTo)

	//insertion point for the staging of instances referenced by pointers
	if controlflowFrom.Start != nil {
		controlflowTo.Start = GongCopyBranchPort(mapOrigCopy, controlflowFrom.Start)
	}
	if controlflowFrom.End != nil {
		controlflowTo.End = GongCopyBranchPort(mapOrigCopy, controlflowFrom.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchControlFlowShape(mapOrigCopy map[any]any, controlflowshapeFrom *ControlFlowShape) (controlflowshapeTo *ControlFlowShape) {

	// controlflowshapeFrom has already been copied
	if _controlflowshapeTo, ok := mapOrigCopy[controlflowshapeFrom]; ok {
		controlflowshapeTo = _controlflowshapeTo.(*ControlFlowShape)
		return
	}

	controlflowshapeTo = new(ControlFlowShape)
	mapOrigCopy[controlflowshapeFrom] = controlflowshapeTo
	controlflowshapeFrom.GongCopyBasicFields(controlflowshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if controlflowshapeFrom.ControlFlow != nil {
		controlflowshapeTo.ControlFlow = GongCopyBranchControlFlow(mapOrigCopy, controlflowshapeFrom.ControlFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchData(mapOrigCopy map[any]any, dataFrom *Data) (dataTo *Data) {

	// dataFrom has already been copied
	if _dataTo, ok := mapOrigCopy[dataFrom]; ok {
		dataTo = _dataTo.(*Data)
		return
	}

	dataTo = new(Data)
	mapOrigCopy[dataFrom] = dataTo
	dataFrom.GongCopyBasicFields(dataTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDataFlow(mapOrigCopy map[any]any, dataflowFrom *DataFlow) (dataflowTo *DataFlow) {

	// dataflowFrom has already been copied
	if _dataflowTo, ok := mapOrigCopy[dataflowFrom]; ok {
		dataflowTo = _dataflowTo.(*DataFlow)
		return
	}

	dataflowTo = new(DataFlow)
	mapOrigCopy[dataflowFrom] = dataflowTo
	dataflowFrom.GongCopyBasicFields(dataflowTo)

	//insertion point for the staging of instances referenced by pointers
	if dataflowFrom.StartPort != nil {
		dataflowTo.StartPort = GongCopyBranchPort(mapOrigCopy, dataflowFrom.StartPort)
	}
	if dataflowFrom.EndPort != nil {
		dataflowTo.EndPort = GongCopyBranchPort(mapOrigCopy, dataflowFrom.EndPort)
	}
	if dataflowFrom.StartExternalPart != nil {
		dataflowTo.StartExternalPart = GongCopyBranchPart(mapOrigCopy, dataflowFrom.StartExternalPart)
	}
	if dataflowFrom.EndExternalPart != nil {
		dataflowTo.EndExternalPart = GongCopyBranchPart(mapOrigCopy, dataflowFrom.EndExternalPart)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _data := range dataflowFrom.Datas {
		dataflowTo.Datas = append(dataflowTo.Datas, GongCopyBranchData(mapOrigCopy, _data))
	}

	return
}

func GongCopyBranchDataFlowShape(mapOrigCopy map[any]any, dataflowshapeFrom *DataFlowShape) (dataflowshapeTo *DataFlowShape) {

	// dataflowshapeFrom has already been copied
	if _dataflowshapeTo, ok := mapOrigCopy[dataflowshapeFrom]; ok {
		dataflowshapeTo = _dataflowshapeTo.(*DataFlowShape)
		return
	}

	dataflowshapeTo = new(DataFlowShape)
	mapOrigCopy[dataflowshapeFrom] = dataflowshapeTo
	dataflowshapeFrom.GongCopyBasicFields(dataflowshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if dataflowshapeFrom.DataFlow != nil {
		dataflowshapeTo.DataFlow = GongCopyBranchDataFlow(mapOrigCopy, dataflowshapeFrom.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDataShape(mapOrigCopy map[any]any, datashapeFrom *DataShape) (datashapeTo *DataShape) {

	// datashapeFrom has already been copied
	if _datashapeTo, ok := mapOrigCopy[datashapeFrom]; ok {
		datashapeTo = _datashapeTo.(*DataShape)
		return
	}

	datashapeTo = new(DataShape)
	mapOrigCopy[datashapeFrom] = datashapeTo
	datashapeFrom.GongCopyBasicFields(datashapeTo)

	//insertion point for the staging of instances referenced by pointers
	if datashapeFrom.Data != nil {
		datashapeTo.Data = GongCopyBranchData(mapOrigCopy, datashapeFrom.Data)
	}
	if datashapeFrom.DataFlow != nil {
		datashapeTo.DataFlow = GongCopyBranchDataFlow(mapOrigCopy, datashapeFrom.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDiagramLayerState(mapOrigCopy map[any]any, diagramlayerstateFrom *DiagramLayerState) (diagramlayerstateTo *DiagramLayerState) {

	// diagramlayerstateFrom has already been copied
	if _diagramlayerstateTo, ok := mapOrigCopy[diagramlayerstateFrom]; ok {
		diagramlayerstateTo = _diagramlayerstateTo.(*DiagramLayerState)
		return
	}

	diagramlayerstateTo = new(DiagramLayerState)
	mapOrigCopy[diagramlayerstateFrom] = diagramlayerstateTo
	diagramlayerstateFrom.GongCopyBasicFields(diagramlayerstateTo)

	//insertion point for the staging of instances referenced by pointers
	if diagramlayerstateFrom.DiagramStructure != nil {
		diagramlayerstateTo.DiagramStructure = GongCopyBranchDiagramStructure(mapOrigCopy, diagramlayerstateFrom.DiagramStructure)
	}
	if diagramlayerstateFrom.LayerDefinition != nil {
		diagramlayerstateTo.LayerDefinition = GongCopyBranchLayerDefinition(mapOrigCopy, diagramlayerstateFrom.LayerDefinition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDiagramStructure(mapOrigCopy map[any]any, diagramstructureFrom *DiagramStructure) (diagramstructureTo *DiagramStructure) {

	// diagramstructureFrom has already been copied
	if _diagramstructureTo, ok := mapOrigCopy[diagramstructureFrom]; ok {
		diagramstructureTo = _diagramstructureTo.(*DiagramStructure)
		return
	}

	diagramstructureTo = new(DiagramStructure)
	mapOrigCopy[diagramstructureFrom] = diagramstructureTo
	diagramstructureFrom.GongCopyBasicFields(diagramstructureTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _systemshape := range diagramstructureFrom.System_Shapes {
		diagramstructureTo.System_Shapes = append(diagramstructureTo.System_Shapes, GongCopyBranchSystemShape(mapOrigCopy, _systemshape))
	}
	for _, _system := range diagramstructureFrom.SystemsWhoseNodeIsExpanded {
		diagramstructureTo.SystemsWhoseNodeIsExpanded = append(diagramstructureTo.SystemsWhoseNodeIsExpanded, GongCopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _partshape := range diagramstructureFrom.Part_Shapes {
		diagramstructureTo.Part_Shapes = append(diagramstructureTo.Part_Shapes, GongCopyBranchPartShape(mapOrigCopy, _partshape))
	}
	for _, _part := range diagramstructureFrom.PartWhoseNodeIsExpanded {
		diagramstructureTo.PartWhoseNodeIsExpanded = append(diagramstructureTo.PartWhoseNodeIsExpanded, GongCopyBranchPart(mapOrigCopy, _part))
	}
	for _, _externalpartshape := range diagramstructureFrom.ExternalPart_Shapes {
		diagramstructureTo.ExternalPart_Shapes = append(diagramstructureTo.ExternalPart_Shapes, GongCopyBranchExternalPartShape(mapOrigCopy, _externalpartshape))
	}
	for _, _part := range diagramstructureFrom.ExternalPartWhoseNodeIsExpanded {
		diagramstructureTo.ExternalPartWhoseNodeIsExpanded = append(diagramstructureTo.ExternalPartWhoseNodeIsExpanded, GongCopyBranchPart(mapOrigCopy, _part))
	}
	for _, _part := range diagramstructureFrom.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
		diagramstructureTo.ExternalPartsWhoseOutDataFlowsNodeIsExpanded = append(diagramstructureTo.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, GongCopyBranchPart(mapOrigCopy, _part))
	}
	for _, _part := range diagramstructureFrom.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
		diagramstructureTo.ExternalPartsWhoseInDataFlowsNodeIsExpanded = append(diagramstructureTo.ExternalPartsWhoseInDataFlowsNodeIsExpanded, GongCopyBranchPart(mapOrigCopy, _part))
	}
	for _, _port := range diagramstructureFrom.PortsWhoseNodeIsExpanded {
		diagramstructureTo.PortsWhoseNodeIsExpanded = append(diagramstructureTo.PortsWhoseNodeIsExpanded, GongCopyBranchPort(mapOrigCopy, _port))
	}
	for _, _portshape := range diagramstructureFrom.Port_Shapes {
		diagramstructureTo.Port_Shapes = append(diagramstructureTo.Port_Shapes, GongCopyBranchPortShape(mapOrigCopy, _portshape))
	}
	for _, _controlflow := range diagramstructureFrom.ControlFlowsWhoseNodeIsExpanded {
		diagramstructureTo.ControlFlowsWhoseNodeIsExpanded = append(diagramstructureTo.ControlFlowsWhoseNodeIsExpanded, GongCopyBranchControlFlow(mapOrigCopy, _controlflow))
	}
	for _, _controlflowshape := range diagramstructureFrom.ControlFlow_Shapes {
		diagramstructureTo.ControlFlow_Shapes = append(diagramstructureTo.ControlFlow_Shapes, GongCopyBranchControlFlowShape(mapOrigCopy, _controlflowshape))
	}
	for _, _dataflow := range diagramstructureFrom.DataFlowsWhoseNodeIsExpanded {
		diagramstructureTo.DataFlowsWhoseNodeIsExpanded = append(diagramstructureTo.DataFlowsWhoseNodeIsExpanded, GongCopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _dataflowshape := range diagramstructureFrom.DataFlow_Shapes {
		diagramstructureTo.DataFlow_Shapes = append(diagramstructureTo.DataFlow_Shapes, GongCopyBranchDataFlowShape(mapOrigCopy, _dataflowshape))
	}
	for _, _data := range diagramstructureFrom.DatasWhoseNodeIsExpanded {
		diagramstructureTo.DatasWhoseNodeIsExpanded = append(diagramstructureTo.DatasWhoseNodeIsExpanded, GongCopyBranchData(mapOrigCopy, _data))
	}
	for _, _datashape := range diagramstructureFrom.Data_Shapes {
		diagramstructureTo.Data_Shapes = append(diagramstructureTo.Data_Shapes, GongCopyBranchDataShape(mapOrigCopy, _datashape))
	}
	for _, _dataflow := range diagramstructureFrom.DataFlowsWhoseDataNodeIsExpanded {
		diagramstructureTo.DataFlowsWhoseDataNodeIsExpanded = append(diagramstructureTo.DataFlowsWhoseDataNodeIsExpanded, GongCopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _resource := range diagramstructureFrom.AllocatedResourcesWhoseNodeIsExpanded {
		diagramstructureTo.AllocatedResourcesWhoseNodeIsExpanded = append(diagramstructureTo.AllocatedResourcesWhoseNodeIsExpanded, GongCopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _allocatedresourceshape := range diagramstructureFrom.AllocatedResourceShapes {
		diagramstructureTo.AllocatedResourceShapes = append(diagramstructureTo.AllocatedResourceShapes, GongCopyBranchAllocatedResourceShape(mapOrigCopy, _allocatedresourceshape))
	}
	for _, _system := range diagramstructureFrom.AllocatedSystemesWhoseNodeIsExpanded {
		diagramstructureTo.AllocatedSystemesWhoseNodeIsExpanded = append(diagramstructureTo.AllocatedSystemesWhoseNodeIsExpanded, GongCopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _allocatedsystemshape := range diagramstructureFrom.AllocatedSystemShapes {
		diagramstructureTo.AllocatedSystemShapes = append(diagramstructureTo.AllocatedSystemShapes, GongCopyBranchAllocatedSystemShape(mapOrigCopy, _allocatedsystemshape))
	}
	for _, _noteshape := range diagramstructureFrom.Note_Shapes {
		diagramstructureTo.Note_Shapes = append(diagramstructureTo.Note_Shapes, GongCopyBranchNoteShape(mapOrigCopy, _noteshape))
	}
	for _, _note := range diagramstructureFrom.NotesWhoseNodeIsExpanded {
		diagramstructureTo.NotesWhoseNodeIsExpanded = append(diagramstructureTo.NotesWhoseNodeIsExpanded, GongCopyBranchNote(mapOrigCopy, _note))
	}
	for _, _noteportshape := range diagramstructureFrom.NotePortShapes {
		diagramstructureTo.NotePortShapes = append(diagramstructureTo.NotePortShapes, GongCopyBranchNotePortShape(mapOrigCopy, _noteportshape))
	}
	for _, _notepartshape := range diagramstructureFrom.NotePartShapes {
		diagramstructureTo.NotePartShapes = append(diagramstructureTo.NotePartShapes, GongCopyBranchNotePartShape(mapOrigCopy, _notepartshape))
	}

	return
}

func GongCopyBranchExternalPartShape(mapOrigCopy map[any]any, externalpartshapeFrom *ExternalPartShape) (externalpartshapeTo *ExternalPartShape) {

	// externalpartshapeFrom has already been copied
	if _externalpartshapeTo, ok := mapOrigCopy[externalpartshapeFrom]; ok {
		externalpartshapeTo = _externalpartshapeTo.(*ExternalPartShape)
		return
	}

	externalpartshapeTo = new(ExternalPartShape)
	mapOrigCopy[externalpartshapeFrom] = externalpartshapeTo
	externalpartshapeFrom.GongCopyBasicFields(externalpartshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if externalpartshapeFrom.Part != nil {
		externalpartshapeTo.Part = GongCopyBranchPart(mapOrigCopy, externalpartshapeFrom.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLayerDefinition(mapOrigCopy map[any]any, layerdefinitionFrom *LayerDefinition) (layerdefinitionTo *LayerDefinition) {

	// layerdefinitionFrom has already been copied
	if _layerdefinitionTo, ok := mapOrigCopy[layerdefinitionFrom]; ok {
		layerdefinitionTo = _layerdefinitionTo.(*LayerDefinition)
		return
	}

	layerdefinitionTo = new(LayerDefinition)
	mapOrigCopy[layerdefinitionFrom] = layerdefinitionTo
	layerdefinitionFrom.GongCopyBasicFields(layerdefinitionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _semantictag := range layerdefinitionFrom.Query {
		layerdefinitionTo.Query = append(layerdefinitionTo.Query, GongCopyBranchSemanticTag(mapOrigCopy, _semantictag))
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
	for _, _system := range libraryFrom.RootSystemes {
		libraryTo.RootSystemes = append(libraryTo.RootSystemes, GongCopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _system := range libraryFrom.SystemsWhoseNodeIsExpanded {
		libraryTo.SystemsWhoseNodeIsExpanded = append(libraryTo.SystemsWhoseNodeIsExpanded, GongCopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _dataflow := range libraryFrom.RootDataFlows {
		libraryTo.RootDataFlows = append(libraryTo.RootDataFlows, GongCopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _dataflow := range libraryFrom.DataFlowsWhoseNodeIsExpanded {
		libraryTo.DataFlowsWhoseNodeIsExpanded = append(libraryTo.DataFlowsWhoseNodeIsExpanded, GongCopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _data := range libraryFrom.RootDatas {
		libraryTo.RootDatas = append(libraryTo.RootDatas, GongCopyBranchData(mapOrigCopy, _data))
	}
	for _, _data := range libraryFrom.DatasWhoseNodeIsExpanded {
		libraryTo.DatasWhoseNodeIsExpanded = append(libraryTo.DatasWhoseNodeIsExpanded, GongCopyBranchData(mapOrigCopy, _data))
	}
	for _, _resource := range libraryFrom.RootResources {
		libraryTo.RootResources = append(libraryTo.RootResources, GongCopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _resource := range libraryFrom.ResourcesWhoseNodeIsExpanded {
		libraryTo.ResourcesWhoseNodeIsExpanded = append(libraryTo.ResourcesWhoseNodeIsExpanded, GongCopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _part := range libraryFrom.PartsWhoseNodeIsExpanded {
		libraryTo.PartsWhoseNodeIsExpanded = append(libraryTo.PartsWhoseNodeIsExpanded, GongCopyBranchPart(mapOrigCopy, _part))
	}
	for _, _note := range libraryFrom.RootNotes {
		libraryTo.RootNotes = append(libraryTo.RootNotes, GongCopyBranchNote(mapOrigCopy, _note))
	}
	for _, _note := range libraryFrom.NotesWhoseNodeIsExpanded {
		libraryTo.NotesWhoseNodeIsExpanded = append(libraryTo.NotesWhoseNodeIsExpanded, GongCopyBranchNote(mapOrigCopy, _note))
	}

	return
}

func GongCopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {

	// noteFrom has already been copied
	if _noteTo, ok := mapOrigCopy[noteFrom]; ok {
		noteTo = _noteTo.(*Note)
		return
	}

	noteTo = new(Note)
	mapOrigCopy[noteFrom] = noteTo
	noteFrom.GongCopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _part := range noteFrom.Parts {
		noteTo.Parts = append(noteTo.Parts, GongCopyBranchPart(mapOrigCopy, _part))
	}
	for _, _port := range noteFrom.Ports {
		noteTo.Ports = append(noteTo.Ports, GongCopyBranchPort(mapOrigCopy, _port))
	}

	return
}

func GongCopyBranchNotePartShape(mapOrigCopy map[any]any, notepartshapeFrom *NotePartShape) (notepartshapeTo *NotePartShape) {

	// notepartshapeFrom has already been copied
	if _notepartshapeTo, ok := mapOrigCopy[notepartshapeFrom]; ok {
		notepartshapeTo = _notepartshapeTo.(*NotePartShape)
		return
	}

	notepartshapeTo = new(NotePartShape)
	mapOrigCopy[notepartshapeFrom] = notepartshapeTo
	notepartshapeFrom.GongCopyBasicFields(notepartshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if notepartshapeFrom.Note != nil {
		notepartshapeTo.Note = GongCopyBranchNote(mapOrigCopy, notepartshapeFrom.Note)
	}
	if notepartshapeFrom.Part != nil {
		notepartshapeTo.Part = GongCopyBranchPart(mapOrigCopy, notepartshapeFrom.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNotePortShape(mapOrigCopy map[any]any, noteportshapeFrom *NotePortShape) (noteportshapeTo *NotePortShape) {

	// noteportshapeFrom has already been copied
	if _noteportshapeTo, ok := mapOrigCopy[noteportshapeFrom]; ok {
		noteportshapeTo = _noteportshapeTo.(*NotePortShape)
		return
	}

	noteportshapeTo = new(NotePortShape)
	mapOrigCopy[noteportshapeFrom] = noteportshapeTo
	noteportshapeFrom.GongCopyBasicFields(noteportshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteportshapeFrom.Note != nil {
		noteportshapeTo.Note = GongCopyBranchNote(mapOrigCopy, noteportshapeFrom.Note)
	}
	if noteportshapeFrom.Port != nil {
		noteportshapeTo.Port = GongCopyBranchPort(mapOrigCopy, noteportshapeFrom.Port)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchNoteShape(mapOrigCopy map[any]any, noteshapeFrom *NoteShape) (noteshapeTo *NoteShape) {

	// noteshapeFrom has already been copied
	if _noteshapeTo, ok := mapOrigCopy[noteshapeFrom]; ok {
		noteshapeTo = _noteshapeTo.(*NoteShape)
		return
	}

	noteshapeTo = new(NoteShape)
	mapOrigCopy[noteshapeFrom] = noteshapeTo
	noteshapeFrom.GongCopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteshapeFrom.Note != nil {
		noteshapeTo.Note = GongCopyBranchNote(mapOrigCopy, noteshapeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPart(mapOrigCopy map[any]any, partFrom *Part) (partTo *Part) {

	// partFrom has already been copied
	if _partTo, ok := mapOrigCopy[partFrom]; ok {
		partTo = _partTo.(*Part)
		return
	}

	partTo = new(Part)
	mapOrigCopy[partFrom] = partTo
	partFrom.GongCopyBasicFields(partTo)

	//insertion point for the staging of instances referenced by pointers
	if partFrom.TypeOfPart != nil {
		partTo.TypeOfPart = GongCopyBranchSystem(mapOrigCopy, partFrom.TypeOfPart)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _port := range partFrom.Ports {
		partTo.Ports = append(partTo.Ports, GongCopyBranchPort(mapOrigCopy, _port))
	}
	for _, _controlflow := range partFrom.ControlFlows {
		partTo.ControlFlows = append(partTo.ControlFlows, GongCopyBranchControlFlow(mapOrigCopy, _controlflow))
	}
	for _, _port := range partFrom.PortWhoseOutControlFlowsNodeIsExpanded {
		partTo.PortWhoseOutControlFlowsNodeIsExpanded = append(partTo.PortWhoseOutControlFlowsNodeIsExpanded, GongCopyBranchPort(mapOrigCopy, _port))
	}
	for _, _port := range partFrom.PortWhoseInControlFlowsNodeIsExpanded {
		partTo.PortWhoseInControlFlowsNodeIsExpanded = append(partTo.PortWhoseInControlFlowsNodeIsExpanded, GongCopyBranchPort(mapOrigCopy, _port))
	}
	for _, _port := range partFrom.PortWhoseOutDataFlowsNodeIsExpanded {
		partTo.PortWhoseOutDataFlowsNodeIsExpanded = append(partTo.PortWhoseOutDataFlowsNodeIsExpanded, GongCopyBranchPort(mapOrigCopy, _port))
	}
	for _, _port := range partFrom.PortWhoseInDataFlowsNodeIsExpanded {
		partTo.PortWhoseInDataFlowsNodeIsExpanded = append(partTo.PortWhoseInDataFlowsNodeIsExpanded, GongCopyBranchPort(mapOrigCopy, _port))
	}
	for _, _partanchoredpath := range partFrom.PartAnchoredPath {
		partTo.PartAnchoredPath = append(partTo.PartAnchoredPath, GongCopyBranchPartAnchoredPath(mapOrigCopy, _partanchoredpath))
	}

	return
}

func GongCopyBranchPartAnchoredPath(mapOrigCopy map[any]any, partanchoredpathFrom *PartAnchoredPath) (partanchoredpathTo *PartAnchoredPath) {

	// partanchoredpathFrom has already been copied
	if _partanchoredpathTo, ok := mapOrigCopy[partanchoredpathFrom]; ok {
		partanchoredpathTo = _partanchoredpathTo.(*PartAnchoredPath)
		return
	}

	partanchoredpathTo = new(PartAnchoredPath)
	mapOrigCopy[partanchoredpathFrom] = partanchoredpathTo
	partanchoredpathFrom.GongCopyBasicFields(partanchoredpathTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartShape(mapOrigCopy map[any]any, partshapeFrom *PartShape) (partshapeTo *PartShape) {

	// partshapeFrom has already been copied
	if _partshapeTo, ok := mapOrigCopy[partshapeFrom]; ok {
		partshapeTo = _partshapeTo.(*PartShape)
		return
	}

	partshapeTo = new(PartShape)
	mapOrigCopy[partshapeFrom] = partshapeTo
	partshapeFrom.GongCopyBasicFields(partshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if partshapeFrom.Part != nil {
		partshapeTo.Part = GongCopyBranchPart(mapOrigCopy, partshapeFrom.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPort(mapOrigCopy map[any]any, portFrom *Port) (portTo *Port) {

	// portFrom has already been copied
	if _portTo, ok := mapOrigCopy[portFrom]; ok {
		portTo = _portTo.(*Port)
		return
	}

	portTo = new(Port)
	mapOrigCopy[portFrom] = portTo
	portFrom.GongCopyBasicFields(portTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPortShape(mapOrigCopy map[any]any, portshapeFrom *PortShape) (portshapeTo *PortShape) {

	// portshapeFrom has already been copied
	if _portshapeTo, ok := mapOrigCopy[portshapeFrom]; ok {
		portshapeTo = _portshapeTo.(*PortShape)
		return
	}

	portshapeTo = new(PortShape)
	mapOrigCopy[portshapeFrom] = portshapeTo
	portshapeFrom.GongCopyBasicFields(portshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if portshapeFrom.Port != nil {
		portshapeTo.Port = GongCopyBranchPort(mapOrigCopy, portshapeFrom.Port)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchResource(mapOrigCopy map[any]any, resourceFrom *Resource) (resourceTo *Resource) {

	// resourceFrom has already been copied
	if _resourceTo, ok := mapOrigCopy[resourceFrom]; ok {
		resourceTo = _resourceTo.(*Resource)
		return
	}

	resourceTo = new(Resource)
	mapOrigCopy[resourceFrom] = resourceTo
	resourceFrom.GongCopyBasicFields(resourceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSemanticTag(mapOrigCopy map[any]any, semantictagFrom *SemanticTag) (semantictagTo *SemanticTag) {

	// semantictagFrom has already been copied
	if _semantictagTo, ok := mapOrigCopy[semantictagFrom]; ok {
		semantictagTo = _semantictagTo.(*SemanticTag)
		return
	}

	semantictagTo = new(SemanticTag)
	mapOrigCopy[semantictagFrom] = semantictagTo
	semantictagFrom.GongCopyBasicFields(semantictagTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _part := range semantictagFrom.Parts {
		semantictagTo.Parts = append(semantictagTo.Parts, GongCopyBranchPart(mapOrigCopy, _part))
	}

	return
}

func GongCopyBranchSystem(mapOrigCopy map[any]any, systemFrom *System) (systemTo *System) {

	// systemFrom has already been copied
	if _systemTo, ok := mapOrigCopy[systemFrom]; ok {
		systemTo = _systemTo.(*System)
		return
	}

	systemTo = new(System)
	mapOrigCopy[systemFrom] = systemTo
	systemFrom.GongCopyBasicFields(systemTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagramstructure := range systemFrom.DiagramStructures {
		systemTo.DiagramStructures = append(systemTo.DiagramStructures, GongCopyBranchDiagramStructure(mapOrigCopy, _diagramstructure))
	}
	for _, _diagramstructure := range systemFrom.DiagramStructureWhoseNodeIsExpanded {
		systemTo.DiagramStructureWhoseNodeIsExpanded = append(systemTo.DiagramStructureWhoseNodeIsExpanded, GongCopyBranchDiagramStructure(mapOrigCopy, _diagramstructure))
	}
	for _, _system := range systemFrom.SubSystemes {
		systemTo.SubSystemes = append(systemTo.SubSystemes, GongCopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _part := range systemFrom.Parts {
		systemTo.Parts = append(systemTo.Parts, GongCopyBranchPart(mapOrigCopy, _part))
	}
	for _, _part := range systemFrom.PartWhoseNodeIsExpanded {
		systemTo.PartWhoseNodeIsExpanded = append(systemTo.PartWhoseNodeIsExpanded, GongCopyBranchPart(mapOrigCopy, _part))
	}
	for _, _dataflow := range systemFrom.DataFlows {
		systemTo.DataFlows = append(systemTo.DataFlows, GongCopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _part := range systemFrom.ExternalParts {
		systemTo.ExternalParts = append(systemTo.ExternalParts, GongCopyBranchPart(mapOrigCopy, _part))
	}
	for _, _part := range systemFrom.ExternalPartWhoseNodeIsExpanded {
		systemTo.ExternalPartWhoseNodeIsExpanded = append(systemTo.ExternalPartWhoseNodeIsExpanded, GongCopyBranchPart(mapOrigCopy, _part))
	}

	return
}

func GongCopyBranchSystemShape(mapOrigCopy map[any]any, systemshapeFrom *SystemShape) (systemshapeTo *SystemShape) {

	// systemshapeFrom has already been copied
	if _systemshapeTo, ok := mapOrigCopy[systemshapeFrom]; ok {
		systemshapeTo = _systemshapeTo.(*SystemShape)
		return
	}

	systemshapeTo = new(SystemShape)
	mapOrigCopy[systemshapeFrom] = systemshapeTo
	systemshapeFrom.GongCopyBasicFields(systemshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if systemshapeFrom.System != nil {
		systemshapeTo.System = GongCopyBranchSystem(mapOrigCopy, systemshapeFrom.System)
	}

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
func (allocatedresourceshape *AllocatedResourceShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchAllocatedResourceShape(allocatedresourceshape)
}

func (stage *Stage) UnstageBranchAllocatedResourceShape(allocatedresourceshape *AllocatedResourceShape) {

	// check if instance is already staged
	if !stage.IsStaged(allocatedresourceshape) {
		return
	}

	allocatedresourceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if allocatedresourceshape.Part != nil {
		stage.UnstageBranch(allocatedresourceshape.Part)
	}
	if allocatedresourceshape.Resource != nil {
		stage.UnstageBranch(allocatedresourceshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (allocatedsystemshape *AllocatedSystemShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchAllocatedSystemShape(allocatedsystemshape)
}

func (stage *Stage) UnstageBranchAllocatedSystemShape(allocatedsystemshape *AllocatedSystemShape) {

	// check if instance is already staged
	if !stage.IsStaged(allocatedsystemshape) {
		return
	}

	allocatedsystemshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if allocatedsystemshape.Part != nil {
		stage.UnstageBranch(allocatedsystemshape.Part)
	}
	if allocatedsystemshape.System != nil {
		stage.UnstageBranch(allocatedsystemshape.System)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (controlflow *ControlFlow) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchControlFlow(controlflow)
}

func (stage *Stage) UnstageBranchControlFlow(controlflow *ControlFlow) {

	// check if instance is already staged
	if !stage.IsStaged(controlflow) {
		return
	}

	controlflow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlflow.Start != nil {
		stage.UnstageBranch(controlflow.Start)
	}
	if controlflow.End != nil {
		stage.UnstageBranch(controlflow.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (controlflowshape *ControlFlowShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchControlFlowShape(controlflowshape)
}

func (stage *Stage) UnstageBranchControlFlowShape(controlflowshape *ControlFlowShape) {

	// check if instance is already staged
	if !stage.IsStaged(controlflowshape) {
		return
	}

	controlflowshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlflowshape.ControlFlow != nil {
		stage.UnstageBranch(controlflowshape.ControlFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (data *Data) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchData(data)
}

func (stage *Stage) UnstageBranchData(data *Data) {

	// check if instance is already staged
	if !stage.IsStaged(data) {
		return
	}

	data.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dataflow *DataFlow) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDataFlow(dataflow)
}

func (stage *Stage) UnstageBranchDataFlow(dataflow *DataFlow) {

	// check if instance is already staged
	if !stage.IsStaged(dataflow) {
		return
	}

	dataflow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dataflow.StartPort != nil {
		stage.UnstageBranch(dataflow.StartPort)
	}
	if dataflow.EndPort != nil {
		stage.UnstageBranch(dataflow.EndPort)
	}
	if dataflow.StartExternalPart != nil {
		stage.UnstageBranch(dataflow.StartExternalPart)
	}
	if dataflow.EndExternalPart != nil {
		stage.UnstageBranch(dataflow.EndExternalPart)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _data := range dataflow.Datas {
		stage.UnstageBranch(_data)
	}

}

func (dataflowshape *DataFlowShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDataFlowShape(dataflowshape)
}

func (stage *Stage) UnstageBranchDataFlowShape(dataflowshape *DataFlowShape) {

	// check if instance is already staged
	if !stage.IsStaged(dataflowshape) {
		return
	}

	dataflowshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dataflowshape.DataFlow != nil {
		stage.UnstageBranch(dataflowshape.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datashape *DataShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDataShape(datashape)
}

func (stage *Stage) UnstageBranchDataShape(datashape *DataShape) {

	// check if instance is already staged
	if !stage.IsStaged(datashape) {
		return
	}

	datashape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datashape.Data != nil {
		stage.UnstageBranch(datashape.Data)
	}
	if datashape.DataFlow != nil {
		stage.UnstageBranch(datashape.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagramlayerstate *DiagramLayerState) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDiagramLayerState(diagramlayerstate)
}

func (stage *Stage) UnstageBranchDiagramLayerState(diagramlayerstate *DiagramLayerState) {

	// check if instance is already staged
	if !stage.IsStaged(diagramlayerstate) {
		return
	}

	diagramlayerstate.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if diagramlayerstate.DiagramStructure != nil {
		stage.UnstageBranch(diagramlayerstate.DiagramStructure)
	}
	if diagramlayerstate.LayerDefinition != nil {
		stage.UnstageBranch(diagramlayerstate.LayerDefinition)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (diagramstructure *DiagramStructure) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDiagramStructure(diagramstructure)
}

func (stage *Stage) UnstageBranchDiagramStructure(diagramstructure *DiagramStructure) {

	// check if instance is already staged
	if !stage.IsStaged(diagramstructure) {
		return
	}

	diagramstructure.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _systemshape := range diagramstructure.System_Shapes {
		stage.UnstageBranch(_systemshape)
	}
	for _, _system := range diagramstructure.SystemsWhoseNodeIsExpanded {
		stage.UnstageBranch(_system)
	}
	for _, _partshape := range diagramstructure.Part_Shapes {
		stage.UnstageBranch(_partshape)
	}
	for _, _part := range diagramstructure.PartWhoseNodeIsExpanded {
		stage.UnstageBranch(_part)
	}
	for _, _externalpartshape := range diagramstructure.ExternalPart_Shapes {
		stage.UnstageBranch(_externalpartshape)
	}
	for _, _part := range diagramstructure.ExternalPartWhoseNodeIsExpanded {
		stage.UnstageBranch(_part)
	}
	for _, _part := range diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
		stage.UnstageBranch(_part)
	}
	for _, _part := range diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
		stage.UnstageBranch(_part)
	}
	for _, _port := range diagramstructure.PortsWhoseNodeIsExpanded {
		stage.UnstageBranch(_port)
	}
	for _, _portshape := range diagramstructure.Port_Shapes {
		stage.UnstageBranch(_portshape)
	}
	for _, _controlflow := range diagramstructure.ControlFlowsWhoseNodeIsExpanded {
		stage.UnstageBranch(_controlflow)
	}
	for _, _controlflowshape := range diagramstructure.ControlFlow_Shapes {
		stage.UnstageBranch(_controlflowshape)
	}
	for _, _dataflow := range diagramstructure.DataFlowsWhoseNodeIsExpanded {
		stage.UnstageBranch(_dataflow)
	}
	for _, _dataflowshape := range diagramstructure.DataFlow_Shapes {
		stage.UnstageBranch(_dataflowshape)
	}
	for _, _data := range diagramstructure.DatasWhoseNodeIsExpanded {
		stage.UnstageBranch(_data)
	}
	for _, _datashape := range diagramstructure.Data_Shapes {
		stage.UnstageBranch(_datashape)
	}
	for _, _dataflow := range diagramstructure.DataFlowsWhoseDataNodeIsExpanded {
		stage.UnstageBranch(_dataflow)
	}
	for _, _resource := range diagramstructure.AllocatedResourcesWhoseNodeIsExpanded {
		stage.UnstageBranch(_resource)
	}
	for _, _allocatedresourceshape := range diagramstructure.AllocatedResourceShapes {
		stage.UnstageBranch(_allocatedresourceshape)
	}
	for _, _system := range diagramstructure.AllocatedSystemesWhoseNodeIsExpanded {
		stage.UnstageBranch(_system)
	}
	for _, _allocatedsystemshape := range diagramstructure.AllocatedSystemShapes {
		stage.UnstageBranch(_allocatedsystemshape)
	}
	for _, _noteshape := range diagramstructure.Note_Shapes {
		stage.UnstageBranch(_noteshape)
	}
	for _, _note := range diagramstructure.NotesWhoseNodeIsExpanded {
		stage.UnstageBranch(_note)
	}
	for _, _noteportshape := range diagramstructure.NotePortShapes {
		stage.UnstageBranch(_noteportshape)
	}
	for _, _notepartshape := range diagramstructure.NotePartShapes {
		stage.UnstageBranch(_notepartshape)
	}

}

func (externalpartshape *ExternalPartShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchExternalPartShape(externalpartshape)
}

func (stage *Stage) UnstageBranchExternalPartShape(externalpartshape *ExternalPartShape) {

	// check if instance is already staged
	if !stage.IsStaged(externalpartshape) {
		return
	}

	externalpartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if externalpartshape.Part != nil {
		stage.UnstageBranch(externalpartshape.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (layerdefinition *LayerDefinition) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchLayerDefinition(layerdefinition)
}

func (stage *Stage) UnstageBranchLayerDefinition(layerdefinition *LayerDefinition) {

	// check if instance is already staged
	if !stage.IsStaged(layerdefinition) {
		return
	}

	layerdefinition.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _semantictag := range layerdefinition.Query {
		stage.UnstageBranch(_semantictag)
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
	for _, _system := range library.RootSystemes {
		stage.UnstageBranch(_system)
	}
	for _, _system := range library.SystemsWhoseNodeIsExpanded {
		stage.UnstageBranch(_system)
	}
	for _, _dataflow := range library.RootDataFlows {
		stage.UnstageBranch(_dataflow)
	}
	for _, _dataflow := range library.DataFlowsWhoseNodeIsExpanded {
		stage.UnstageBranch(_dataflow)
	}
	for _, _data := range library.RootDatas {
		stage.UnstageBranch(_data)
	}
	for _, _data := range library.DatasWhoseNodeIsExpanded {
		stage.UnstageBranch(_data)
	}
	for _, _resource := range library.RootResources {
		stage.UnstageBranch(_resource)
	}
	for _, _resource := range library.ResourcesWhoseNodeIsExpanded {
		stage.UnstageBranch(_resource)
	}
	for _, _part := range library.PartsWhoseNodeIsExpanded {
		stage.UnstageBranch(_part)
	}
	for _, _note := range library.RootNotes {
		stage.UnstageBranch(_note)
	}
	for _, _note := range library.NotesWhoseNodeIsExpanded {
		stage.UnstageBranch(_note)
	}

}

func (note *Note) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNote(note)
}

func (stage *Stage) UnstageBranchNote(note *Note) {

	// check if instance is already staged
	if !stage.IsStaged(note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _part := range note.Parts {
		stage.UnstageBranch(_part)
	}
	for _, _port := range note.Ports {
		stage.UnstageBranch(_port)
	}

}

func (notepartshape *NotePartShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNotePartShape(notepartshape)
}

func (stage *Stage) UnstageBranchNotePartShape(notepartshape *NotePartShape) {

	// check if instance is already staged
	if !stage.IsStaged(notepartshape) {
		return
	}

	notepartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if notepartshape.Note != nil {
		stage.UnstageBranch(notepartshape.Note)
	}
	if notepartshape.Part != nil {
		stage.UnstageBranch(notepartshape.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteportshape *NotePortShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNotePortShape(noteportshape)
}

func (stage *Stage) UnstageBranchNotePortShape(noteportshape *NotePortShape) {

	// check if instance is already staged
	if !stage.IsStaged(noteportshape) {
		return
	}

	noteportshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteportshape.Note != nil {
		stage.UnstageBranch(noteportshape.Note)
	}
	if noteportshape.Port != nil {
		stage.UnstageBranch(noteportshape.Port)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (noteshape *NoteShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchNoteShape(noteshape)
}

func (stage *Stage) UnstageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if !stage.IsStaged(noteshape) {
		return
	}

	noteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		stage.UnstageBranch(noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (part *Part) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPart(part)
}

func (stage *Stage) UnstageBranchPart(part *Part) {

	// check if instance is already staged
	if !stage.IsStaged(part) {
		return
	}

	part.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if part.TypeOfPart != nil {
		stage.UnstageBranch(part.TypeOfPart)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _port := range part.Ports {
		stage.UnstageBranch(_port)
	}
	for _, _controlflow := range part.ControlFlows {
		stage.UnstageBranch(_controlflow)
	}
	for _, _port := range part.PortWhoseOutControlFlowsNodeIsExpanded {
		stage.UnstageBranch(_port)
	}
	for _, _port := range part.PortWhoseInControlFlowsNodeIsExpanded {
		stage.UnstageBranch(_port)
	}
	for _, _port := range part.PortWhoseOutDataFlowsNodeIsExpanded {
		stage.UnstageBranch(_port)
	}
	for _, _port := range part.PortWhoseInDataFlowsNodeIsExpanded {
		stage.UnstageBranch(_port)
	}
	for _, _partanchoredpath := range part.PartAnchoredPath {
		stage.UnstageBranch(_partanchoredpath)
	}

}

func (partanchoredpath *PartAnchoredPath) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPartAnchoredPath(partanchoredpath)
}

func (stage *Stage) UnstageBranchPartAnchoredPath(partanchoredpath *PartAnchoredPath) {

	// check if instance is already staged
	if !stage.IsStaged(partanchoredpath) {
		return
	}

	partanchoredpath.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partshape *PartShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPartShape(partshape)
}

func (stage *Stage) UnstageBranchPartShape(partshape *PartShape) {

	// check if instance is already staged
	if !stage.IsStaged(partshape) {
		return
	}

	partshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if partshape.Part != nil {
		stage.UnstageBranch(partshape.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (port *Port) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPort(port)
}

func (stage *Stage) UnstageBranchPort(port *Port) {

	// check if instance is already staged
	if !stage.IsStaged(port) {
		return
	}

	port.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (portshape *PortShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPortShape(portshape)
}

func (stage *Stage) UnstageBranchPortShape(portshape *PortShape) {

	// check if instance is already staged
	if !stage.IsStaged(portshape) {
		return
	}

	portshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if portshape.Port != nil {
		stage.UnstageBranch(portshape.Port)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (resource *Resource) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchResource(resource)
}

func (stage *Stage) UnstageBranchResource(resource *Resource) {

	// check if instance is already staged
	if !stage.IsStaged(resource) {
		return
	}

	resource.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (semantictag *SemanticTag) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSemanticTag(semantictag)
}

func (stage *Stage) UnstageBranchSemanticTag(semantictag *SemanticTag) {

	// check if instance is already staged
	if !stage.IsStaged(semantictag) {
		return
	}

	semantictag.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _part := range semantictag.Parts {
		stage.UnstageBranch(_part)
	}

}

func (system *System) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSystem(system)
}

func (stage *Stage) UnstageBranchSystem(system *System) {

	// check if instance is already staged
	if !stage.IsStaged(system) {
		return
	}

	system.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagramstructure := range system.DiagramStructures {
		stage.UnstageBranch(_diagramstructure)
	}
	for _, _diagramstructure := range system.DiagramStructureWhoseNodeIsExpanded {
		stage.UnstageBranch(_diagramstructure)
	}
	for _, _system := range system.SubSystemes {
		stage.UnstageBranch(_system)
	}
	for _, _part := range system.Parts {
		stage.UnstageBranch(_part)
	}
	for _, _part := range system.PartWhoseNodeIsExpanded {
		stage.UnstageBranch(_part)
	}
	for _, _dataflow := range system.DataFlows {
		stage.UnstageBranch(_dataflow)
	}
	for _, _part := range system.ExternalParts {
		stage.UnstageBranch(_part)
	}
	for _, _part := range system.ExternalPartWhoseNodeIsExpanded {
		stage.UnstageBranch(_part)
	}

}

func (systemshape *SystemShape) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSystemShape(systemshape)
}

func (stage *Stage) UnstageBranchSystemShape(systemshape *SystemShape) {

	// check if instance is already staged
	if !stage.IsStaged(systemshape) {
		return
	}

	systemshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if systemshape.System != nil {
		stage.UnstageBranch(systemshape.System)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *AllocatedResourceShape) GongReconstructPointersFromReferences(stage *Stage, instance *AllocatedResourceShape) {
	// insertion point for pointers field
	if instance.Part != nil {
		reference.Part = stage.Parts_reference[instance.Part]
	}
	if instance.Resource != nil {
		reference.Resource = stage.Resources_reference[instance.Resource]
	}
	// insertion point for slice of pointers field
}

func (reference *AllocatedSystemShape) GongReconstructPointersFromReferences(stage *Stage, instance *AllocatedSystemShape) {
	// insertion point for pointers field
	if instance.Part != nil {
		reference.Part = stage.Parts_reference[instance.Part]
	}
	if instance.System != nil {
		reference.System = stage.Systems_reference[instance.System]
	}
	// insertion point for slice of pointers field
}

func (reference *ControlFlow) GongReconstructPointersFromReferences(stage *Stage, instance *ControlFlow) {
	// insertion point for pointers field
	if instance.Start != nil {
		reference.Start = stage.Ports_reference[instance.Start]
	}
	if instance.End != nil {
		reference.End = stage.Ports_reference[instance.End]
	}
	// insertion point for slice of pointers field
}

func (reference *ControlFlowShape) GongReconstructPointersFromReferences(stage *Stage, instance *ControlFlowShape) {
	// insertion point for pointers field
	if instance.ControlFlow != nil {
		reference.ControlFlow = stage.ControlFlows_reference[instance.ControlFlow]
	}
	// insertion point for slice of pointers field
}

func (reference *Data) GongReconstructPointersFromReferences(stage *Stage, instance *Data) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DataFlow) GongReconstructPointersFromReferences(stage *Stage, instance *DataFlow) {
	// insertion point for pointers field
	if instance.StartPort != nil {
		reference.StartPort = stage.Ports_reference[instance.StartPort]
	}
	if instance.EndPort != nil {
		reference.EndPort = stage.Ports_reference[instance.EndPort]
	}
	if instance.StartExternalPart != nil {
		reference.StartExternalPart = stage.Parts_reference[instance.StartExternalPart]
	}
	if instance.EndExternalPart != nil {
		reference.EndExternalPart = stage.Parts_reference[instance.EndExternalPart]
	}
	// insertion point for slice of pointers field
	reference.Datas = reference.Datas[:0]
	for _, _b := range instance.Datas {
		reference.Datas = append(reference.Datas, stage.Datas_reference[_b])
	}
}

func (reference *DataFlowShape) GongReconstructPointersFromReferences(stage *Stage, instance *DataFlowShape) {
	// insertion point for pointers field
	if instance.DataFlow != nil {
		reference.DataFlow = stage.DataFlows_reference[instance.DataFlow]
	}
	// insertion point for slice of pointers field
}

func (reference *DataShape) GongReconstructPointersFromReferences(stage *Stage, instance *DataShape) {
	// insertion point for pointers field
	if instance.Data != nil {
		reference.Data = stage.Datas_reference[instance.Data]
	}
	if instance.DataFlow != nil {
		reference.DataFlow = stage.DataFlows_reference[instance.DataFlow]
	}
	// insertion point for slice of pointers field
}

func (reference *DiagramLayerState) GongReconstructPointersFromReferences(stage *Stage, instance *DiagramLayerState) {
	// insertion point for pointers field
	if instance.DiagramStructure != nil {
		reference.DiagramStructure = stage.DiagramStructures_reference[instance.DiagramStructure]
	}
	if instance.LayerDefinition != nil {
		reference.LayerDefinition = stage.LayerDefinitions_reference[instance.LayerDefinition]
	}
	// insertion point for slice of pointers field
}

func (reference *DiagramStructure) GongReconstructPointersFromReferences(stage *Stage, instance *DiagramStructure) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.System_Shapes = reference.System_Shapes[:0]
	for _, _b := range instance.System_Shapes {
		reference.System_Shapes = append(reference.System_Shapes, stage.SystemShapes_reference[_b])
	}
	reference.SystemsWhoseNodeIsExpanded = reference.SystemsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.SystemsWhoseNodeIsExpanded {
		reference.SystemsWhoseNodeIsExpanded = append(reference.SystemsWhoseNodeIsExpanded, stage.Systems_reference[_b])
	}
	reference.Part_Shapes = reference.Part_Shapes[:0]
	for _, _b := range instance.Part_Shapes {
		reference.Part_Shapes = append(reference.Part_Shapes, stage.PartShapes_reference[_b])
	}
	reference.PartWhoseNodeIsExpanded = reference.PartWhoseNodeIsExpanded[:0]
	for _, _b := range instance.PartWhoseNodeIsExpanded {
		reference.PartWhoseNodeIsExpanded = append(reference.PartWhoseNodeIsExpanded, stage.Parts_reference[_b])
	}
	reference.ExternalPart_Shapes = reference.ExternalPart_Shapes[:0]
	for _, _b := range instance.ExternalPart_Shapes {
		reference.ExternalPart_Shapes = append(reference.ExternalPart_Shapes, stage.ExternalPartShapes_reference[_b])
	}
	reference.ExternalPartWhoseNodeIsExpanded = reference.ExternalPartWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ExternalPartWhoseNodeIsExpanded {
		reference.ExternalPartWhoseNodeIsExpanded = append(reference.ExternalPartWhoseNodeIsExpanded, stage.Parts_reference[_b])
	}
	reference.ExternalPartsWhoseOutDataFlowsNodeIsExpanded = reference.ExternalPartsWhoseOutDataFlowsNodeIsExpanded[:0]
	for _, _b := range instance.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
		reference.ExternalPartsWhoseOutDataFlowsNodeIsExpanded = append(reference.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, stage.Parts_reference[_b])
	}
	reference.ExternalPartsWhoseInDataFlowsNodeIsExpanded = reference.ExternalPartsWhoseInDataFlowsNodeIsExpanded[:0]
	for _, _b := range instance.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
		reference.ExternalPartsWhoseInDataFlowsNodeIsExpanded = append(reference.ExternalPartsWhoseInDataFlowsNodeIsExpanded, stage.Parts_reference[_b])
	}
	reference.PortsWhoseNodeIsExpanded = reference.PortsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.PortsWhoseNodeIsExpanded {
		reference.PortsWhoseNodeIsExpanded = append(reference.PortsWhoseNodeIsExpanded, stage.Ports_reference[_b])
	}
	reference.Port_Shapes = reference.Port_Shapes[:0]
	for _, _b := range instance.Port_Shapes {
		reference.Port_Shapes = append(reference.Port_Shapes, stage.PortShapes_reference[_b])
	}
	reference.ControlFlowsWhoseNodeIsExpanded = reference.ControlFlowsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ControlFlowsWhoseNodeIsExpanded {
		reference.ControlFlowsWhoseNodeIsExpanded = append(reference.ControlFlowsWhoseNodeIsExpanded, stage.ControlFlows_reference[_b])
	}
	reference.ControlFlow_Shapes = reference.ControlFlow_Shapes[:0]
	for _, _b := range instance.ControlFlow_Shapes {
		reference.ControlFlow_Shapes = append(reference.ControlFlow_Shapes, stage.ControlFlowShapes_reference[_b])
	}
	reference.DataFlowsWhoseNodeIsExpanded = reference.DataFlowsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.DataFlowsWhoseNodeIsExpanded {
		reference.DataFlowsWhoseNodeIsExpanded = append(reference.DataFlowsWhoseNodeIsExpanded, stage.DataFlows_reference[_b])
	}
	reference.DataFlow_Shapes = reference.DataFlow_Shapes[:0]
	for _, _b := range instance.DataFlow_Shapes {
		reference.DataFlow_Shapes = append(reference.DataFlow_Shapes, stage.DataFlowShapes_reference[_b])
	}
	reference.DatasWhoseNodeIsExpanded = reference.DatasWhoseNodeIsExpanded[:0]
	for _, _b := range instance.DatasWhoseNodeIsExpanded {
		reference.DatasWhoseNodeIsExpanded = append(reference.DatasWhoseNodeIsExpanded, stage.Datas_reference[_b])
	}
	reference.Data_Shapes = reference.Data_Shapes[:0]
	for _, _b := range instance.Data_Shapes {
		reference.Data_Shapes = append(reference.Data_Shapes, stage.DataShapes_reference[_b])
	}
	reference.DataFlowsWhoseDataNodeIsExpanded = reference.DataFlowsWhoseDataNodeIsExpanded[:0]
	for _, _b := range instance.DataFlowsWhoseDataNodeIsExpanded {
		reference.DataFlowsWhoseDataNodeIsExpanded = append(reference.DataFlowsWhoseDataNodeIsExpanded, stage.DataFlows_reference[_b])
	}
	reference.AllocatedResourcesWhoseNodeIsExpanded = reference.AllocatedResourcesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.AllocatedResourcesWhoseNodeIsExpanded {
		reference.AllocatedResourcesWhoseNodeIsExpanded = append(reference.AllocatedResourcesWhoseNodeIsExpanded, stage.Resources_reference[_b])
	}
	reference.AllocatedResourceShapes = reference.AllocatedResourceShapes[:0]
	for _, _b := range instance.AllocatedResourceShapes {
		reference.AllocatedResourceShapes = append(reference.AllocatedResourceShapes, stage.AllocatedResourceShapes_reference[_b])
	}
	reference.AllocatedSystemesWhoseNodeIsExpanded = reference.AllocatedSystemesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.AllocatedSystemesWhoseNodeIsExpanded {
		reference.AllocatedSystemesWhoseNodeIsExpanded = append(reference.AllocatedSystemesWhoseNodeIsExpanded, stage.Systems_reference[_b])
	}
	reference.AllocatedSystemShapes = reference.AllocatedSystemShapes[:0]
	for _, _b := range instance.AllocatedSystemShapes {
		reference.AllocatedSystemShapes = append(reference.AllocatedSystemShapes, stage.AllocatedSystemShapes_reference[_b])
	}
	reference.Note_Shapes = reference.Note_Shapes[:0]
	for _, _b := range instance.Note_Shapes {
		reference.Note_Shapes = append(reference.Note_Shapes, stage.NoteShapes_reference[_b])
	}
	reference.NotesWhoseNodeIsExpanded = reference.NotesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.NotesWhoseNodeIsExpanded {
		reference.NotesWhoseNodeIsExpanded = append(reference.NotesWhoseNodeIsExpanded, stage.Notes_reference[_b])
	}
	reference.NotePortShapes = reference.NotePortShapes[:0]
	for _, _b := range instance.NotePortShapes {
		reference.NotePortShapes = append(reference.NotePortShapes, stage.NotePortShapes_reference[_b])
	}
	reference.NotePartShapes = reference.NotePartShapes[:0]
	for _, _b := range instance.NotePartShapes {
		reference.NotePartShapes = append(reference.NotePartShapes, stage.NotePartShapes_reference[_b])
	}
}

func (reference *ExternalPartShape) GongReconstructPointersFromReferences(stage *Stage, instance *ExternalPartShape) {
	// insertion point for pointers field
	if instance.Part != nil {
		reference.Part = stage.Parts_reference[instance.Part]
	}
	// insertion point for slice of pointers field
}

func (reference *LayerDefinition) GongReconstructPointersFromReferences(stage *Stage, instance *LayerDefinition) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Query = reference.Query[:0]
	for _, _b := range instance.Query {
		reference.Query = append(reference.Query, stage.SemanticTags_reference[_b])
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
	reference.RootSystemes = reference.RootSystemes[:0]
	for _, _b := range instance.RootSystemes {
		reference.RootSystemes = append(reference.RootSystemes, stage.Systems_reference[_b])
	}
	reference.SystemsWhoseNodeIsExpanded = reference.SystemsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.SystemsWhoseNodeIsExpanded {
		reference.SystemsWhoseNodeIsExpanded = append(reference.SystemsWhoseNodeIsExpanded, stage.Systems_reference[_b])
	}
	reference.RootDataFlows = reference.RootDataFlows[:0]
	for _, _b := range instance.RootDataFlows {
		reference.RootDataFlows = append(reference.RootDataFlows, stage.DataFlows_reference[_b])
	}
	reference.DataFlowsWhoseNodeIsExpanded = reference.DataFlowsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.DataFlowsWhoseNodeIsExpanded {
		reference.DataFlowsWhoseNodeIsExpanded = append(reference.DataFlowsWhoseNodeIsExpanded, stage.DataFlows_reference[_b])
	}
	reference.RootDatas = reference.RootDatas[:0]
	for _, _b := range instance.RootDatas {
		reference.RootDatas = append(reference.RootDatas, stage.Datas_reference[_b])
	}
	reference.DatasWhoseNodeIsExpanded = reference.DatasWhoseNodeIsExpanded[:0]
	for _, _b := range instance.DatasWhoseNodeIsExpanded {
		reference.DatasWhoseNodeIsExpanded = append(reference.DatasWhoseNodeIsExpanded, stage.Datas_reference[_b])
	}
	reference.RootResources = reference.RootResources[:0]
	for _, _b := range instance.RootResources {
		reference.RootResources = append(reference.RootResources, stage.Resources_reference[_b])
	}
	reference.ResourcesWhoseNodeIsExpanded = reference.ResourcesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ResourcesWhoseNodeIsExpanded {
		reference.ResourcesWhoseNodeIsExpanded = append(reference.ResourcesWhoseNodeIsExpanded, stage.Resources_reference[_b])
	}
	reference.PartsWhoseNodeIsExpanded = reference.PartsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.PartsWhoseNodeIsExpanded {
		reference.PartsWhoseNodeIsExpanded = append(reference.PartsWhoseNodeIsExpanded, stage.Parts_reference[_b])
	}
	reference.RootNotes = reference.RootNotes[:0]
	for _, _b := range instance.RootNotes {
		reference.RootNotes = append(reference.RootNotes, stage.Notes_reference[_b])
	}
	reference.NotesWhoseNodeIsExpanded = reference.NotesWhoseNodeIsExpanded[:0]
	for _, _b := range instance.NotesWhoseNodeIsExpanded {
		reference.NotesWhoseNodeIsExpanded = append(reference.NotesWhoseNodeIsExpanded, stage.Notes_reference[_b])
	}
}

func (reference *Note) GongReconstructPointersFromReferences(stage *Stage, instance *Note) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Parts = reference.Parts[:0]
	for _, _b := range instance.Parts {
		reference.Parts = append(reference.Parts, stage.Parts_reference[_b])
	}
	reference.Ports = reference.Ports[:0]
	for _, _b := range instance.Ports {
		reference.Ports = append(reference.Ports, stage.Ports_reference[_b])
	}
}

func (reference *NotePartShape) GongReconstructPointersFromReferences(stage *Stage, instance *NotePartShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	if instance.Part != nil {
		reference.Part = stage.Parts_reference[instance.Part]
	}
	// insertion point for slice of pointers field
}

func (reference *NotePortShape) GongReconstructPointersFromReferences(stage *Stage, instance *NotePortShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	if instance.Port != nil {
		reference.Port = stage.Ports_reference[instance.Port]
	}
	// insertion point for slice of pointers field
}

func (reference *NoteShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteShape) {
	// insertion point for pointers field
	if instance.Note != nil {
		reference.Note = stage.Notes_reference[instance.Note]
	}
	// insertion point for slice of pointers field
}

func (reference *Part) GongReconstructPointersFromReferences(stage *Stage, instance *Part) {
	// insertion point for pointers field
	if instance.TypeOfPart != nil {
		reference.TypeOfPart = stage.Systems_reference[instance.TypeOfPart]
	}
	// insertion point for slice of pointers field
	reference.Ports = reference.Ports[:0]
	for _, _b := range instance.Ports {
		reference.Ports = append(reference.Ports, stage.Ports_reference[_b])
	}
	reference.ControlFlows = reference.ControlFlows[:0]
	for _, _b := range instance.ControlFlows {
		reference.ControlFlows = append(reference.ControlFlows, stage.ControlFlows_reference[_b])
	}
	reference.PortWhoseOutControlFlowsNodeIsExpanded = reference.PortWhoseOutControlFlowsNodeIsExpanded[:0]
	for _, _b := range instance.PortWhoseOutControlFlowsNodeIsExpanded {
		reference.PortWhoseOutControlFlowsNodeIsExpanded = append(reference.PortWhoseOutControlFlowsNodeIsExpanded, stage.Ports_reference[_b])
	}
	reference.PortWhoseInControlFlowsNodeIsExpanded = reference.PortWhoseInControlFlowsNodeIsExpanded[:0]
	for _, _b := range instance.PortWhoseInControlFlowsNodeIsExpanded {
		reference.PortWhoseInControlFlowsNodeIsExpanded = append(reference.PortWhoseInControlFlowsNodeIsExpanded, stage.Ports_reference[_b])
	}
	reference.PortWhoseOutDataFlowsNodeIsExpanded = reference.PortWhoseOutDataFlowsNodeIsExpanded[:0]
	for _, _b := range instance.PortWhoseOutDataFlowsNodeIsExpanded {
		reference.PortWhoseOutDataFlowsNodeIsExpanded = append(reference.PortWhoseOutDataFlowsNodeIsExpanded, stage.Ports_reference[_b])
	}
	reference.PortWhoseInDataFlowsNodeIsExpanded = reference.PortWhoseInDataFlowsNodeIsExpanded[:0]
	for _, _b := range instance.PortWhoseInDataFlowsNodeIsExpanded {
		reference.PortWhoseInDataFlowsNodeIsExpanded = append(reference.PortWhoseInDataFlowsNodeIsExpanded, stage.Ports_reference[_b])
	}
	reference.PartAnchoredPath = reference.PartAnchoredPath[:0]
	for _, _b := range instance.PartAnchoredPath {
		reference.PartAnchoredPath = append(reference.PartAnchoredPath, stage.PartAnchoredPaths_reference[_b])
	}
}

func (reference *PartAnchoredPath) GongReconstructPointersFromReferences(stage *Stage, instance *PartAnchoredPath) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartShape) GongReconstructPointersFromReferences(stage *Stage, instance *PartShape) {
	// insertion point for pointers field
	if instance.Part != nil {
		reference.Part = stage.Parts_reference[instance.Part]
	}
	// insertion point for slice of pointers field
}

func (reference *Port) GongReconstructPointersFromReferences(stage *Stage, instance *Port) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PortShape) GongReconstructPointersFromReferences(stage *Stage, instance *PortShape) {
	// insertion point for pointers field
	if instance.Port != nil {
		reference.Port = stage.Ports_reference[instance.Port]
	}
	// insertion point for slice of pointers field
}

func (reference *Resource) GongReconstructPointersFromReferences(stage *Stage, instance *Resource) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *SemanticTag) GongReconstructPointersFromReferences(stage *Stage, instance *SemanticTag) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Parts = reference.Parts[:0]
	for _, _b := range instance.Parts {
		reference.Parts = append(reference.Parts, stage.Parts_reference[_b])
	}
}

func (reference *System) GongReconstructPointersFromReferences(stage *Stage, instance *System) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.DiagramStructures = reference.DiagramStructures[:0]
	for _, _b := range instance.DiagramStructures {
		reference.DiagramStructures = append(reference.DiagramStructures, stage.DiagramStructures_reference[_b])
	}
	reference.DiagramStructureWhoseNodeIsExpanded = reference.DiagramStructureWhoseNodeIsExpanded[:0]
	for _, _b := range instance.DiagramStructureWhoseNodeIsExpanded {
		reference.DiagramStructureWhoseNodeIsExpanded = append(reference.DiagramStructureWhoseNodeIsExpanded, stage.DiagramStructures_reference[_b])
	}
	reference.SubSystemes = reference.SubSystemes[:0]
	for _, _b := range instance.SubSystemes {
		reference.SubSystemes = append(reference.SubSystemes, stage.Systems_reference[_b])
	}
	reference.Parts = reference.Parts[:0]
	for _, _b := range instance.Parts {
		reference.Parts = append(reference.Parts, stage.Parts_reference[_b])
	}
	reference.PartWhoseNodeIsExpanded = reference.PartWhoseNodeIsExpanded[:0]
	for _, _b := range instance.PartWhoseNodeIsExpanded {
		reference.PartWhoseNodeIsExpanded = append(reference.PartWhoseNodeIsExpanded, stage.Parts_reference[_b])
	}
	reference.DataFlows = reference.DataFlows[:0]
	for _, _b := range instance.DataFlows {
		reference.DataFlows = append(reference.DataFlows, stage.DataFlows_reference[_b])
	}
	reference.ExternalParts = reference.ExternalParts[:0]
	for _, _b := range instance.ExternalParts {
		reference.ExternalParts = append(reference.ExternalParts, stage.Parts_reference[_b])
	}
	reference.ExternalPartWhoseNodeIsExpanded = reference.ExternalPartWhoseNodeIsExpanded[:0]
	for _, _b := range instance.ExternalPartWhoseNodeIsExpanded {
		reference.ExternalPartWhoseNodeIsExpanded = append(reference.ExternalPartWhoseNodeIsExpanded, stage.Parts_reference[_b])
	}
}

func (reference *SystemShape) GongReconstructPointersFromReferences(stage *Stage, instance *SystemShape) {
	// insertion point for pointers field
	if instance.System != nil {
		reference.System = stage.Systems_reference[instance.System]
	}
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AllocatedResourceShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Part; _reference != nil {
		reference.Part = nil
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			reference.Part = _instance
		}
	}
	if _reference := reference.Resource; _reference != nil {
		reference.Resource = nil
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			reference.Resource = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *AllocatedSystemShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Part; _reference != nil {
		reference.Part = nil
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			reference.Part = _instance
		}
	}
	if _reference := reference.System; _reference != nil {
		reference.System = nil
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			reference.System = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ControlFlow) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Start; _reference != nil {
		reference.Start = nil
		if _instance, ok := stage.Ports_instance[_reference]; ok {
			reference.Start = _instance
		}
	}
	if _reference := reference.End; _reference != nil {
		reference.End = nil
		if _instance, ok := stage.Ports_instance[_reference]; ok {
			reference.End = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ControlFlowShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ControlFlow; _reference != nil {
		reference.ControlFlow = nil
		if _instance, ok := stage.ControlFlows_instance[_reference]; ok {
			reference.ControlFlow = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Data) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DataFlow) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.StartPort; _reference != nil {
		reference.StartPort = nil
		if _instance, ok := stage.Ports_instance[_reference]; ok {
			reference.StartPort = _instance
		}
	}
	if _reference := reference.EndPort; _reference != nil {
		reference.EndPort = nil
		if _instance, ok := stage.Ports_instance[_reference]; ok {
			reference.EndPort = _instance
		}
	}
	if _reference := reference.StartExternalPart; _reference != nil {
		reference.StartExternalPart = nil
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			reference.StartExternalPart = _instance
		}
	}
	if _reference := reference.EndExternalPart; _reference != nil {
		reference.EndExternalPart = nil
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			reference.EndExternalPart = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Datas []*Data
	for _, _reference := range reference.Datas {
		if _instance, ok := stage.Datas_instance[_reference]; ok {
			_Datas = append(_Datas, _instance)
		}
	}
	reference.Datas = _Datas
}

func (reference *DataFlowShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.DataFlow; _reference != nil {
		reference.DataFlow = nil
		if _instance, ok := stage.DataFlows_instance[_reference]; ok {
			reference.DataFlow = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *DataShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Data; _reference != nil {
		reference.Data = nil
		if _instance, ok := stage.Datas_instance[_reference]; ok {
			reference.Data = _instance
		}
	}
	if _reference := reference.DataFlow; _reference != nil {
		reference.DataFlow = nil
		if _instance, ok := stage.DataFlows_instance[_reference]; ok {
			reference.DataFlow = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *DiagramLayerState) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.DiagramStructure; _reference != nil {
		reference.DiagramStructure = nil
		if _instance, ok := stage.DiagramStructures_instance[_reference]; ok {
			reference.DiagramStructure = _instance
		}
	}
	if _reference := reference.LayerDefinition; _reference != nil {
		reference.LayerDefinition = nil
		if _instance, ok := stage.LayerDefinitions_instance[_reference]; ok {
			reference.LayerDefinition = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *DiagramStructure) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _System_Shapes []*SystemShape
	for _, _reference := range reference.System_Shapes {
		if _instance, ok := stage.SystemShapes_instance[_reference]; ok {
			_System_Shapes = append(_System_Shapes, _instance)
		}
	}
	reference.System_Shapes = _System_Shapes
	var _SystemsWhoseNodeIsExpanded []*System
	for _, _reference := range reference.SystemsWhoseNodeIsExpanded {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_SystemsWhoseNodeIsExpanded = append(_SystemsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.SystemsWhoseNodeIsExpanded = _SystemsWhoseNodeIsExpanded
	var _Part_Shapes []*PartShape
	for _, _reference := range reference.Part_Shapes {
		if _instance, ok := stage.PartShapes_instance[_reference]; ok {
			_Part_Shapes = append(_Part_Shapes, _instance)
		}
	}
	reference.Part_Shapes = _Part_Shapes
	var _PartWhoseNodeIsExpanded []*Part
	for _, _reference := range reference.PartWhoseNodeIsExpanded {
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			_PartWhoseNodeIsExpanded = append(_PartWhoseNodeIsExpanded, _instance)
		}
	}
	reference.PartWhoseNodeIsExpanded = _PartWhoseNodeIsExpanded
	var _ExternalPart_Shapes []*ExternalPartShape
	for _, _reference := range reference.ExternalPart_Shapes {
		if _instance, ok := stage.ExternalPartShapes_instance[_reference]; ok {
			_ExternalPart_Shapes = append(_ExternalPart_Shapes, _instance)
		}
	}
	reference.ExternalPart_Shapes = _ExternalPart_Shapes
	var _ExternalPartWhoseNodeIsExpanded []*Part
	for _, _reference := range reference.ExternalPartWhoseNodeIsExpanded {
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			_ExternalPartWhoseNodeIsExpanded = append(_ExternalPartWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ExternalPartWhoseNodeIsExpanded = _ExternalPartWhoseNodeIsExpanded
	var _ExternalPartsWhoseOutDataFlowsNodeIsExpanded []*Part
	for _, _reference := range reference.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			_ExternalPartsWhoseOutDataFlowsNodeIsExpanded = append(_ExternalPartsWhoseOutDataFlowsNodeIsExpanded, _instance)
		}
	}
	reference.ExternalPartsWhoseOutDataFlowsNodeIsExpanded = _ExternalPartsWhoseOutDataFlowsNodeIsExpanded
	var _ExternalPartsWhoseInDataFlowsNodeIsExpanded []*Part
	for _, _reference := range reference.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			_ExternalPartsWhoseInDataFlowsNodeIsExpanded = append(_ExternalPartsWhoseInDataFlowsNodeIsExpanded, _instance)
		}
	}
	reference.ExternalPartsWhoseInDataFlowsNodeIsExpanded = _ExternalPartsWhoseInDataFlowsNodeIsExpanded
	var _PortsWhoseNodeIsExpanded []*Port
	for _, _reference := range reference.PortsWhoseNodeIsExpanded {
		if _instance, ok := stage.Ports_instance[_reference]; ok {
			_PortsWhoseNodeIsExpanded = append(_PortsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.PortsWhoseNodeIsExpanded = _PortsWhoseNodeIsExpanded
	var _Port_Shapes []*PortShape
	for _, _reference := range reference.Port_Shapes {
		if _instance, ok := stage.PortShapes_instance[_reference]; ok {
			_Port_Shapes = append(_Port_Shapes, _instance)
		}
	}
	reference.Port_Shapes = _Port_Shapes
	var _ControlFlowsWhoseNodeIsExpanded []*ControlFlow
	for _, _reference := range reference.ControlFlowsWhoseNodeIsExpanded {
		if _instance, ok := stage.ControlFlows_instance[_reference]; ok {
			_ControlFlowsWhoseNodeIsExpanded = append(_ControlFlowsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ControlFlowsWhoseNodeIsExpanded = _ControlFlowsWhoseNodeIsExpanded
	var _ControlFlow_Shapes []*ControlFlowShape
	for _, _reference := range reference.ControlFlow_Shapes {
		if _instance, ok := stage.ControlFlowShapes_instance[_reference]; ok {
			_ControlFlow_Shapes = append(_ControlFlow_Shapes, _instance)
		}
	}
	reference.ControlFlow_Shapes = _ControlFlow_Shapes
	var _DataFlowsWhoseNodeIsExpanded []*DataFlow
	for _, _reference := range reference.DataFlowsWhoseNodeIsExpanded {
		if _instance, ok := stage.DataFlows_instance[_reference]; ok {
			_DataFlowsWhoseNodeIsExpanded = append(_DataFlowsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.DataFlowsWhoseNodeIsExpanded = _DataFlowsWhoseNodeIsExpanded
	var _DataFlow_Shapes []*DataFlowShape
	for _, _reference := range reference.DataFlow_Shapes {
		if _instance, ok := stage.DataFlowShapes_instance[_reference]; ok {
			_DataFlow_Shapes = append(_DataFlow_Shapes, _instance)
		}
	}
	reference.DataFlow_Shapes = _DataFlow_Shapes
	var _DatasWhoseNodeIsExpanded []*Data
	for _, _reference := range reference.DatasWhoseNodeIsExpanded {
		if _instance, ok := stage.Datas_instance[_reference]; ok {
			_DatasWhoseNodeIsExpanded = append(_DatasWhoseNodeIsExpanded, _instance)
		}
	}
	reference.DatasWhoseNodeIsExpanded = _DatasWhoseNodeIsExpanded
	var _Data_Shapes []*DataShape
	for _, _reference := range reference.Data_Shapes {
		if _instance, ok := stage.DataShapes_instance[_reference]; ok {
			_Data_Shapes = append(_Data_Shapes, _instance)
		}
	}
	reference.Data_Shapes = _Data_Shapes
	var _DataFlowsWhoseDataNodeIsExpanded []*DataFlow
	for _, _reference := range reference.DataFlowsWhoseDataNodeIsExpanded {
		if _instance, ok := stage.DataFlows_instance[_reference]; ok {
			_DataFlowsWhoseDataNodeIsExpanded = append(_DataFlowsWhoseDataNodeIsExpanded, _instance)
		}
	}
	reference.DataFlowsWhoseDataNodeIsExpanded = _DataFlowsWhoseDataNodeIsExpanded
	var _AllocatedResourcesWhoseNodeIsExpanded []*Resource
	for _, _reference := range reference.AllocatedResourcesWhoseNodeIsExpanded {
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			_AllocatedResourcesWhoseNodeIsExpanded = append(_AllocatedResourcesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.AllocatedResourcesWhoseNodeIsExpanded = _AllocatedResourcesWhoseNodeIsExpanded
	var _AllocatedResourceShapes []*AllocatedResourceShape
	for _, _reference := range reference.AllocatedResourceShapes {
		if _instance, ok := stage.AllocatedResourceShapes_instance[_reference]; ok {
			_AllocatedResourceShapes = append(_AllocatedResourceShapes, _instance)
		}
	}
	reference.AllocatedResourceShapes = _AllocatedResourceShapes
	var _AllocatedSystemesWhoseNodeIsExpanded []*System
	for _, _reference := range reference.AllocatedSystemesWhoseNodeIsExpanded {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_AllocatedSystemesWhoseNodeIsExpanded = append(_AllocatedSystemesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.AllocatedSystemesWhoseNodeIsExpanded = _AllocatedSystemesWhoseNodeIsExpanded
	var _AllocatedSystemShapes []*AllocatedSystemShape
	for _, _reference := range reference.AllocatedSystemShapes {
		if _instance, ok := stage.AllocatedSystemShapes_instance[_reference]; ok {
			_AllocatedSystemShapes = append(_AllocatedSystemShapes, _instance)
		}
	}
	reference.AllocatedSystemShapes = _AllocatedSystemShapes
	var _Note_Shapes []*NoteShape
	for _, _reference := range reference.Note_Shapes {
		if _instance, ok := stage.NoteShapes_instance[_reference]; ok {
			_Note_Shapes = append(_Note_Shapes, _instance)
		}
	}
	reference.Note_Shapes = _Note_Shapes
	var _NotesWhoseNodeIsExpanded []*Note
	for _, _reference := range reference.NotesWhoseNodeIsExpanded {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_NotesWhoseNodeIsExpanded = append(_NotesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.NotesWhoseNodeIsExpanded = _NotesWhoseNodeIsExpanded
	var _NotePortShapes []*NotePortShape
	for _, _reference := range reference.NotePortShapes {
		if _instance, ok := stage.NotePortShapes_instance[_reference]; ok {
			_NotePortShapes = append(_NotePortShapes, _instance)
		}
	}
	reference.NotePortShapes = _NotePortShapes
	var _NotePartShapes []*NotePartShape
	for _, _reference := range reference.NotePartShapes {
		if _instance, ok := stage.NotePartShapes_instance[_reference]; ok {
			_NotePartShapes = append(_NotePartShapes, _instance)
		}
	}
	reference.NotePartShapes = _NotePartShapes
}

func (reference *ExternalPartShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Part; _reference != nil {
		reference.Part = nil
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			reference.Part = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *LayerDefinition) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Query []*SemanticTag
	for _, _reference := range reference.Query {
		if _instance, ok := stage.SemanticTags_instance[_reference]; ok {
			_Query = append(_Query, _instance)
		}
	}
	reference.Query = _Query
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
	var _RootSystemes []*System
	for _, _reference := range reference.RootSystemes {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_RootSystemes = append(_RootSystemes, _instance)
		}
	}
	reference.RootSystemes = _RootSystemes
	var _SystemsWhoseNodeIsExpanded []*System
	for _, _reference := range reference.SystemsWhoseNodeIsExpanded {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_SystemsWhoseNodeIsExpanded = append(_SystemsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.SystemsWhoseNodeIsExpanded = _SystemsWhoseNodeIsExpanded
	var _RootDataFlows []*DataFlow
	for _, _reference := range reference.RootDataFlows {
		if _instance, ok := stage.DataFlows_instance[_reference]; ok {
			_RootDataFlows = append(_RootDataFlows, _instance)
		}
	}
	reference.RootDataFlows = _RootDataFlows
	var _DataFlowsWhoseNodeIsExpanded []*DataFlow
	for _, _reference := range reference.DataFlowsWhoseNodeIsExpanded {
		if _instance, ok := stage.DataFlows_instance[_reference]; ok {
			_DataFlowsWhoseNodeIsExpanded = append(_DataFlowsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.DataFlowsWhoseNodeIsExpanded = _DataFlowsWhoseNodeIsExpanded
	var _RootDatas []*Data
	for _, _reference := range reference.RootDatas {
		if _instance, ok := stage.Datas_instance[_reference]; ok {
			_RootDatas = append(_RootDatas, _instance)
		}
	}
	reference.RootDatas = _RootDatas
	var _DatasWhoseNodeIsExpanded []*Data
	for _, _reference := range reference.DatasWhoseNodeIsExpanded {
		if _instance, ok := stage.Datas_instance[_reference]; ok {
			_DatasWhoseNodeIsExpanded = append(_DatasWhoseNodeIsExpanded, _instance)
		}
	}
	reference.DatasWhoseNodeIsExpanded = _DatasWhoseNodeIsExpanded
	var _RootResources []*Resource
	for _, _reference := range reference.RootResources {
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			_RootResources = append(_RootResources, _instance)
		}
	}
	reference.RootResources = _RootResources
	var _ResourcesWhoseNodeIsExpanded []*Resource
	for _, _reference := range reference.ResourcesWhoseNodeIsExpanded {
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			_ResourcesWhoseNodeIsExpanded = append(_ResourcesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ResourcesWhoseNodeIsExpanded = _ResourcesWhoseNodeIsExpanded
	var _PartsWhoseNodeIsExpanded []*Part
	for _, _reference := range reference.PartsWhoseNodeIsExpanded {
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			_PartsWhoseNodeIsExpanded = append(_PartsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.PartsWhoseNodeIsExpanded = _PartsWhoseNodeIsExpanded
	var _RootNotes []*Note
	for _, _reference := range reference.RootNotes {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_RootNotes = append(_RootNotes, _instance)
		}
	}
	reference.RootNotes = _RootNotes
	var _NotesWhoseNodeIsExpanded []*Note
	for _, _reference := range reference.NotesWhoseNodeIsExpanded {
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			_NotesWhoseNodeIsExpanded = append(_NotesWhoseNodeIsExpanded, _instance)
		}
	}
	reference.NotesWhoseNodeIsExpanded = _NotesWhoseNodeIsExpanded
}

func (reference *Note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Parts []*Part
	for _, _reference := range reference.Parts {
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			_Parts = append(_Parts, _instance)
		}
	}
	reference.Parts = _Parts
	var _Ports []*Port
	for _, _reference := range reference.Ports {
		if _instance, ok := stage.Ports_instance[_reference]; ok {
			_Ports = append(_Ports, _instance)
		}
	}
	reference.Ports = _Ports
}

func (reference *NotePartShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	if _reference := reference.Part; _reference != nil {
		reference.Part = nil
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			reference.Part = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *NotePortShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	if _reference := reference.Port; _reference != nil {
		reference.Port = nil
		if _instance, ok := stage.Ports_instance[_reference]; ok {
			reference.Port = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *NoteShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Note; _reference != nil {
		reference.Note = nil
		if _instance, ok := stage.Notes_instance[_reference]; ok {
			reference.Note = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Part) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.TypeOfPart; _reference != nil {
		reference.TypeOfPart = nil
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			reference.TypeOfPart = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _Ports []*Port
	for _, _reference := range reference.Ports {
		if _instance, ok := stage.Ports_instance[_reference]; ok {
			_Ports = append(_Ports, _instance)
		}
	}
	reference.Ports = _Ports
	var _ControlFlows []*ControlFlow
	for _, _reference := range reference.ControlFlows {
		if _instance, ok := stage.ControlFlows_instance[_reference]; ok {
			_ControlFlows = append(_ControlFlows, _instance)
		}
	}
	reference.ControlFlows = _ControlFlows
	var _PortWhoseOutControlFlowsNodeIsExpanded []*Port
	for _, _reference := range reference.PortWhoseOutControlFlowsNodeIsExpanded {
		if _instance, ok := stage.Ports_instance[_reference]; ok {
			_PortWhoseOutControlFlowsNodeIsExpanded = append(_PortWhoseOutControlFlowsNodeIsExpanded, _instance)
		}
	}
	reference.PortWhoseOutControlFlowsNodeIsExpanded = _PortWhoseOutControlFlowsNodeIsExpanded
	var _PortWhoseInControlFlowsNodeIsExpanded []*Port
	for _, _reference := range reference.PortWhoseInControlFlowsNodeIsExpanded {
		if _instance, ok := stage.Ports_instance[_reference]; ok {
			_PortWhoseInControlFlowsNodeIsExpanded = append(_PortWhoseInControlFlowsNodeIsExpanded, _instance)
		}
	}
	reference.PortWhoseInControlFlowsNodeIsExpanded = _PortWhoseInControlFlowsNodeIsExpanded
	var _PortWhoseOutDataFlowsNodeIsExpanded []*Port
	for _, _reference := range reference.PortWhoseOutDataFlowsNodeIsExpanded {
		if _instance, ok := stage.Ports_instance[_reference]; ok {
			_PortWhoseOutDataFlowsNodeIsExpanded = append(_PortWhoseOutDataFlowsNodeIsExpanded, _instance)
		}
	}
	reference.PortWhoseOutDataFlowsNodeIsExpanded = _PortWhoseOutDataFlowsNodeIsExpanded
	var _PortWhoseInDataFlowsNodeIsExpanded []*Port
	for _, _reference := range reference.PortWhoseInDataFlowsNodeIsExpanded {
		if _instance, ok := stage.Ports_instance[_reference]; ok {
			_PortWhoseInDataFlowsNodeIsExpanded = append(_PortWhoseInDataFlowsNodeIsExpanded, _instance)
		}
	}
	reference.PortWhoseInDataFlowsNodeIsExpanded = _PortWhoseInDataFlowsNodeIsExpanded
	var _PartAnchoredPath []*PartAnchoredPath
	for _, _reference := range reference.PartAnchoredPath {
		if _instance, ok := stage.PartAnchoredPaths_instance[_reference]; ok {
			_PartAnchoredPath = append(_PartAnchoredPath, _instance)
		}
	}
	reference.PartAnchoredPath = _PartAnchoredPath
}

func (reference *PartAnchoredPath) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Part; _reference != nil {
		reference.Part = nil
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			reference.Part = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Port) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PortShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Port; _reference != nil {
		reference.Port = nil
		if _instance, ok := stage.Ports_instance[_reference]; ok {
			reference.Port = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Resource) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *SemanticTag) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Parts []*Part
	for _, _reference := range reference.Parts {
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			_Parts = append(_Parts, _instance)
		}
	}
	reference.Parts = _Parts
}

func (reference *System) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _DiagramStructures []*DiagramStructure
	for _, _reference := range reference.DiagramStructures {
		if _instance, ok := stage.DiagramStructures_instance[_reference]; ok {
			_DiagramStructures = append(_DiagramStructures, _instance)
		}
	}
	reference.DiagramStructures = _DiagramStructures
	var _DiagramStructureWhoseNodeIsExpanded []*DiagramStructure
	for _, _reference := range reference.DiagramStructureWhoseNodeIsExpanded {
		if _instance, ok := stage.DiagramStructures_instance[_reference]; ok {
			_DiagramStructureWhoseNodeIsExpanded = append(_DiagramStructureWhoseNodeIsExpanded, _instance)
		}
	}
	reference.DiagramStructureWhoseNodeIsExpanded = _DiagramStructureWhoseNodeIsExpanded
	var _SubSystemes []*System
	for _, _reference := range reference.SubSystemes {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_SubSystemes = append(_SubSystemes, _instance)
		}
	}
	reference.SubSystemes = _SubSystemes
	var _Parts []*Part
	for _, _reference := range reference.Parts {
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			_Parts = append(_Parts, _instance)
		}
	}
	reference.Parts = _Parts
	var _PartWhoseNodeIsExpanded []*Part
	for _, _reference := range reference.PartWhoseNodeIsExpanded {
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			_PartWhoseNodeIsExpanded = append(_PartWhoseNodeIsExpanded, _instance)
		}
	}
	reference.PartWhoseNodeIsExpanded = _PartWhoseNodeIsExpanded
	var _DataFlows []*DataFlow
	for _, _reference := range reference.DataFlows {
		if _instance, ok := stage.DataFlows_instance[_reference]; ok {
			_DataFlows = append(_DataFlows, _instance)
		}
	}
	reference.DataFlows = _DataFlows
	var _ExternalParts []*Part
	for _, _reference := range reference.ExternalParts {
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			_ExternalParts = append(_ExternalParts, _instance)
		}
	}
	reference.ExternalParts = _ExternalParts
	var _ExternalPartWhoseNodeIsExpanded []*Part
	for _, _reference := range reference.ExternalPartWhoseNodeIsExpanded {
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			_ExternalPartWhoseNodeIsExpanded = append(_ExternalPartWhoseNodeIsExpanded, _instance)
		}
	}
	reference.ExternalPartWhoseNodeIsExpanded = _ExternalPartWhoseNodeIsExpanded
}

func (reference *SystemShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.System; _reference != nil {
		reference.System = nil
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			reference.System = _instance
		}
	}
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (allocatedresourceshape *AllocatedResourceShape) GongDiff(stage *Stage, allocatedresourceshapeOther *AllocatedResourceShape) (diffs []string) {
	// insertion point for field diffs
	if allocatedresourceshape.Name != allocatedresourceshapeOther.Name {
		diffs = append(diffs, allocatedresourceshape.GongMarshallField(stage, "Name"))
	}
	if (allocatedresourceshape.Part == nil) != (allocatedresourceshapeOther.Part == nil) {
		diffs = append(diffs, allocatedresourceshape.GongMarshallField(stage, "Part"))
	} else if allocatedresourceshape.Part != nil && allocatedresourceshapeOther.Part != nil {
		if allocatedresourceshape.Part != allocatedresourceshapeOther.Part {
			diffs = append(diffs, allocatedresourceshape.GongMarshallField(stage, "Part"))
		}
	}
	if (allocatedresourceshape.Resource == nil) != (allocatedresourceshapeOther.Resource == nil) {
		diffs = append(diffs, allocatedresourceshape.GongMarshallField(stage, "Resource"))
	} else if allocatedresourceshape.Resource != nil && allocatedresourceshapeOther.Resource != nil {
		if allocatedresourceshape.Resource != allocatedresourceshapeOther.Resource {
			diffs = append(diffs, allocatedresourceshape.GongMarshallField(stage, "Resource"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (allocatedsystemshape *AllocatedSystemShape) GongDiff(stage *Stage, allocatedsystemshapeOther *AllocatedSystemShape) (diffs []string) {
	// insertion point for field diffs
	if allocatedsystemshape.Name != allocatedsystemshapeOther.Name {
		diffs = append(diffs, allocatedsystemshape.GongMarshallField(stage, "Name"))
	}
	if (allocatedsystemshape.Part == nil) != (allocatedsystemshapeOther.Part == nil) {
		diffs = append(diffs, allocatedsystemshape.GongMarshallField(stage, "Part"))
	} else if allocatedsystemshape.Part != nil && allocatedsystemshapeOther.Part != nil {
		if allocatedsystemshape.Part != allocatedsystemshapeOther.Part {
			diffs = append(diffs, allocatedsystemshape.GongMarshallField(stage, "Part"))
		}
	}
	if (allocatedsystemshape.System == nil) != (allocatedsystemshapeOther.System == nil) {
		diffs = append(diffs, allocatedsystemshape.GongMarshallField(stage, "System"))
	} else if allocatedsystemshape.System != nil && allocatedsystemshapeOther.System != nil {
		if allocatedsystemshape.System != allocatedsystemshapeOther.System {
			diffs = append(diffs, allocatedsystemshape.GongMarshallField(stage, "System"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (controlflow *ControlFlow) GongDiff(stage *Stage, controlflowOther *ControlFlow) (diffs []string) {
	// insertion point for field diffs
	if controlflow.Name != controlflowOther.Name {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "Name"))
	}
	if controlflow.Description != controlflowOther.Description {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "Description"))
	}
	if controlflow.ComputedPrefix != controlflowOther.ComputedPrefix {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "ComputedPrefix"))
	}
	if controlflow.IsExpanded != controlflowOther.IsExpanded {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "IsExpanded"))
	}
	if (controlflow.Start == nil) != (controlflowOther.Start == nil) {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "Start"))
	} else if controlflow.Start != nil && controlflowOther.Start != nil {
		if controlflow.Start != controlflowOther.Start {
			diffs = append(diffs, controlflow.GongMarshallField(stage, "Start"))
		}
	}
	if (controlflow.End == nil) != (controlflowOther.End == nil) {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "End"))
	} else if controlflow.End != nil && controlflowOther.End != nil {
		if controlflow.End != controlflowOther.End {
			diffs = append(diffs, controlflow.GongMarshallField(stage, "End"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (controlflowshape *ControlFlowShape) GongDiff(stage *Stage, controlflowshapeOther *ControlFlowShape) (diffs []string) {
	// insertion point for field diffs
	if controlflowshape.Name != controlflowshapeOther.Name {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "Name"))
	}
	if (controlflowshape.ControlFlow == nil) != (controlflowshapeOther.ControlFlow == nil) {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "ControlFlow"))
	} else if controlflowshape.ControlFlow != nil && controlflowshapeOther.ControlFlow != nil {
		if controlflowshape.ControlFlow != controlflowshapeOther.ControlFlow {
			diffs = append(diffs, controlflowshape.GongMarshallField(stage, "ControlFlow"))
		}
	}
	if controlflowshape.StartRatio != controlflowshapeOther.StartRatio {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "StartRatio"))
	}
	if controlflowshape.EndRatio != controlflowshapeOther.EndRatio {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "EndRatio"))
	}
	if controlflowshape.StartOrientation != controlflowshapeOther.StartOrientation {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "StartOrientation"))
	}
	if controlflowshape.EndOrientation != controlflowshapeOther.EndOrientation {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "EndOrientation"))
	}
	if controlflowshape.CornerOffsetRatio != controlflowshapeOther.CornerOffsetRatio {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if controlflowshape.IsHidden != controlflowshapeOther.IsHidden {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (data *Data) GongDiff(stage *Stage, dataOther *Data) (diffs []string) {
	// insertion point for field diffs
	if data.Name != dataOther.Name {
		diffs = append(diffs, data.GongMarshallField(stage, "Name"))
	}
	if data.Acronym != dataOther.Acronym {
		diffs = append(diffs, data.GongMarshallField(stage, "Acronym"))
	}
	if data.Description != dataOther.Description {
		diffs = append(diffs, data.GongMarshallField(stage, "Description"))
	}
	if data.ComputedPrefix != dataOther.ComputedPrefix {
		diffs = append(diffs, data.GongMarshallField(stage, "ComputedPrefix"))
	}
	if data.IsExpanded != dataOther.IsExpanded {
		diffs = append(diffs, data.GongMarshallField(stage, "IsExpanded"))
	}
	if data.SVG_Path != dataOther.SVG_Path {
		diffs = append(diffs, data.GongMarshallField(stage, "SVG_Path"))
	}
	if data.InverseAppliedScaling != dataOther.InverseAppliedScaling {
		diffs = append(diffs, data.GongMarshallField(stage, "InverseAppliedScaling"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (dataflow *DataFlow) GongDiff(stage *Stage, dataflowOther *DataFlow) (diffs []string) {
	// insertion point for field diffs
	if dataflow.Name != dataflowOther.Name {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "Name"))
	}
	if (dataflow.StartPort == nil) != (dataflowOther.StartPort == nil) {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "StartPort"))
	} else if dataflow.StartPort != nil && dataflowOther.StartPort != nil {
		if dataflow.StartPort != dataflowOther.StartPort {
			diffs = append(diffs, dataflow.GongMarshallField(stage, "StartPort"))
		}
	}
	if (dataflow.EndPort == nil) != (dataflowOther.EndPort == nil) {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "EndPort"))
	} else if dataflow.EndPort != nil && dataflowOther.EndPort != nil {
		if dataflow.EndPort != dataflowOther.EndPort {
			diffs = append(diffs, dataflow.GongMarshallField(stage, "EndPort"))
		}
	}
	if (dataflow.StartExternalPart == nil) != (dataflowOther.StartExternalPart == nil) {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "StartExternalPart"))
	} else if dataflow.StartExternalPart != nil && dataflowOther.StartExternalPart != nil {
		if dataflow.StartExternalPart != dataflowOther.StartExternalPart {
			diffs = append(diffs, dataflow.GongMarshallField(stage, "StartExternalPart"))
		}
	}
	if (dataflow.EndExternalPart == nil) != (dataflowOther.EndExternalPart == nil) {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "EndExternalPart"))
	} else if dataflow.EndExternalPart != nil && dataflowOther.EndExternalPart != nil {
		if dataflow.EndExternalPart != dataflowOther.EndExternalPart {
			diffs = append(diffs, dataflow.GongMarshallField(stage, "EndExternalPart"))
		}
	}
	DatasDifferent := false
	if len(dataflow.Datas) != len(dataflowOther.Datas) {
		DatasDifferent = true
	} else {
		for i := range dataflow.Datas {
			if (dataflow.Datas[i] == nil) != (dataflowOther.Datas[i] == nil) {
				DatasDifferent = true
				break
			} else if dataflow.Datas[i] != nil && dataflowOther.Datas[i] != nil {
				// this is a pointer comparaison
				if dataflow.Datas[i] != dataflowOther.Datas[i] {
					DatasDifferent = true
					break
				}
			}
		}
	}
	if DatasDifferent {
		ops := stage.Diff(
			dataflow,
			"Datas",
			len(dataflowOther.Datas),
			len(dataflow.Datas),
			func(i, j int) bool {
				return dataflowOther.Datas[i] == dataflow.Datas[j]
			},
			func(j int) string {
				return dataflow.Datas[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if dataflow.Description != dataflowOther.Description {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "Description"))
	}
	if dataflow.Direction != dataflowOther.Direction {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "Direction"))
	}
	if dataflow.IsDatasNodeExpanded != dataflowOther.IsDatasNodeExpanded {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "IsDatasNodeExpanded"))
	}
	if dataflow.ComputedPrefix != dataflowOther.ComputedPrefix {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "ComputedPrefix"))
	}
	if dataflow.IsExpanded != dataflowOther.IsExpanded {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "IsExpanded"))
	}
	if dataflow.Type != dataflowOther.Type {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "Type"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (dataflowshape *DataFlowShape) GongDiff(stage *Stage, dataflowshapeOther *DataFlowShape) (diffs []string) {
	// insertion point for field diffs
	if dataflowshape.Name != dataflowshapeOther.Name {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "Name"))
	}
	if (dataflowshape.DataFlow == nil) != (dataflowshapeOther.DataFlow == nil) {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "DataFlow"))
	} else if dataflowshape.DataFlow != nil && dataflowshapeOther.DataFlow != nil {
		if dataflowshape.DataFlow != dataflowshapeOther.DataFlow {
			diffs = append(diffs, dataflowshape.GongMarshallField(stage, "DataFlow"))
		}
	}
	if dataflowshape.StartRatio != dataflowshapeOther.StartRatio {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "StartRatio"))
	}
	if dataflowshape.EndRatio != dataflowshapeOther.EndRatio {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "EndRatio"))
	}
	if dataflowshape.StartOrientation != dataflowshapeOther.StartOrientation {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "StartOrientation"))
	}
	if dataflowshape.EndOrientation != dataflowshapeOther.EndOrientation {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "EndOrientation"))
	}
	if dataflowshape.CornerOffsetRatio != dataflowshapeOther.CornerOffsetRatio {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if dataflowshape.IsHidden != dataflowshapeOther.IsHidden {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (datashape *DataShape) GongDiff(stage *Stage, datashapeOther *DataShape) (diffs []string) {
	// insertion point for field diffs
	if datashape.Name != datashapeOther.Name {
		diffs = append(diffs, datashape.GongMarshallField(stage, "Name"))
	}
	if (datashape.Data == nil) != (datashapeOther.Data == nil) {
		diffs = append(diffs, datashape.GongMarshallField(stage, "Data"))
	} else if datashape.Data != nil && datashapeOther.Data != nil {
		if datashape.Data != datashapeOther.Data {
			diffs = append(diffs, datashape.GongMarshallField(stage, "Data"))
		}
	}
	if (datashape.DataFlow == nil) != (datashapeOther.DataFlow == nil) {
		diffs = append(diffs, datashape.GongMarshallField(stage, "DataFlow"))
	} else if datashape.DataFlow != nil && datashapeOther.DataFlow != nil {
		if datashape.DataFlow != datashapeOther.DataFlow {
			diffs = append(diffs, datashape.GongMarshallField(stage, "DataFlow"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (diagramlayerstate *DiagramLayerState) GongDiff(stage *Stage, diagramlayerstateOther *DiagramLayerState) (diffs []string) {
	// insertion point for field diffs
	if diagramlayerstate.Name != diagramlayerstateOther.Name {
		diffs = append(diffs, diagramlayerstate.GongMarshallField(stage, "Name"))
	}
	if (diagramlayerstate.DiagramStructure == nil) != (diagramlayerstateOther.DiagramStructure == nil) {
		diffs = append(diffs, diagramlayerstate.GongMarshallField(stage, "DiagramStructure"))
	} else if diagramlayerstate.DiagramStructure != nil && diagramlayerstateOther.DiagramStructure != nil {
		if diagramlayerstate.DiagramStructure != diagramlayerstateOther.DiagramStructure {
			diffs = append(diffs, diagramlayerstate.GongMarshallField(stage, "DiagramStructure"))
		}
	}
	if (diagramlayerstate.LayerDefinition == nil) != (diagramlayerstateOther.LayerDefinition == nil) {
		diffs = append(diffs, diagramlayerstate.GongMarshallField(stage, "LayerDefinition"))
	} else if diagramlayerstate.LayerDefinition != nil && diagramlayerstateOther.LayerDefinition != nil {
		if diagramlayerstate.LayerDefinition != diagramlayerstateOther.LayerDefinition {
			diffs = append(diffs, diagramlayerstate.GongMarshallField(stage, "LayerDefinition"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (diagramstructure *DiagramStructure) GongDiff(stage *Stage, diagramstructureOther *DiagramStructure) (diffs []string) {
	// insertion point for field diffs
	if diagramstructure.Name != diagramstructureOther.Name {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "Name"))
	}
	if diagramstructure.Description != diagramstructureOther.Description {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "Description"))
	}
	if diagramstructure.ComputedPrefix != diagramstructureOther.ComputedPrefix {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "ComputedPrefix"))
	}
	if diagramstructure.IsExpanded != diagramstructureOther.IsExpanded {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "IsExpanded"))
	}
	if diagramstructure.IsChecked != diagramstructureOther.IsChecked {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "IsChecked"))
	}
	if diagramstructure.IsEditable_ != diagramstructureOther.IsEditable_ {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "IsEditable_"))
	}
	if diagramstructure.IsShowPrefix != diagramstructureOther.IsShowPrefix {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "IsShowPrefix"))
	}
	if diagramstructure.DefaultBoxWidth != diagramstructureOther.DefaultBoxWidth {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "DefaultBoxWidth"))
	}
	if diagramstructure.DefaultBoxHeigth != diagramstructureOther.DefaultBoxHeigth {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "DefaultBoxHeigth"))
	}
	if diagramstructure.IsWithDiscretePorts != diagramstructureOther.IsWithDiscretePorts {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "IsWithDiscretePorts"))
	}
	if diagramstructure.Width != diagramstructureOther.Width {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "Width"))
	}
	if diagramstructure.Height != diagramstructureOther.Height {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "Height"))
	}
	System_ShapesDifferent := false
	if len(diagramstructure.System_Shapes) != len(diagramstructureOther.System_Shapes) {
		System_ShapesDifferent = true
	} else {
		for i := range diagramstructure.System_Shapes {
			if (diagramstructure.System_Shapes[i] == nil) != (diagramstructureOther.System_Shapes[i] == nil) {
				System_ShapesDifferent = true
				break
			} else if diagramstructure.System_Shapes[i] != nil && diagramstructureOther.System_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.System_Shapes[i] != diagramstructureOther.System_Shapes[i] {
					System_ShapesDifferent = true
					break
				}
			}
		}
	}
	if System_ShapesDifferent {
		ops := stage.Diff(
			diagramstructure,
			"System_Shapes",
			len(diagramstructureOther.System_Shapes),
			len(diagramstructure.System_Shapes),
			func(i, j int) bool {
				return diagramstructureOther.System_Shapes[i] == diagramstructure.System_Shapes[j]
			},
			func(j int) string {
				return diagramstructure.System_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if diagramstructure.IsSystemsNodeExpanded != diagramstructureOther.IsSystemsNodeExpanded {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "IsSystemsNodeExpanded"))
	}
	SystemsWhoseNodeIsExpandedDifferent := false
	if len(diagramstructure.SystemsWhoseNodeIsExpanded) != len(diagramstructureOther.SystemsWhoseNodeIsExpanded) {
		SystemsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramstructure.SystemsWhoseNodeIsExpanded {
			if (diagramstructure.SystemsWhoseNodeIsExpanded[i] == nil) != (diagramstructureOther.SystemsWhoseNodeIsExpanded[i] == nil) {
				SystemsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramstructure.SystemsWhoseNodeIsExpanded[i] != nil && diagramstructureOther.SystemsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.SystemsWhoseNodeIsExpanded[i] != diagramstructureOther.SystemsWhoseNodeIsExpanded[i] {
					SystemsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if SystemsWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramstructure,
			"SystemsWhoseNodeIsExpanded",
			len(diagramstructureOther.SystemsWhoseNodeIsExpanded),
			len(diagramstructure.SystemsWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramstructureOther.SystemsWhoseNodeIsExpanded[i] == diagramstructure.SystemsWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramstructure.SystemsWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	Part_ShapesDifferent := false
	if len(diagramstructure.Part_Shapes) != len(diagramstructureOther.Part_Shapes) {
		Part_ShapesDifferent = true
	} else {
		for i := range diagramstructure.Part_Shapes {
			if (diagramstructure.Part_Shapes[i] == nil) != (diagramstructureOther.Part_Shapes[i] == nil) {
				Part_ShapesDifferent = true
				break
			} else if diagramstructure.Part_Shapes[i] != nil && diagramstructureOther.Part_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.Part_Shapes[i] != diagramstructureOther.Part_Shapes[i] {
					Part_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Part_ShapesDifferent {
		ops := stage.Diff(
			diagramstructure,
			"Part_Shapes",
			len(diagramstructureOther.Part_Shapes),
			len(diagramstructure.Part_Shapes),
			func(i, j int) bool {
				return diagramstructureOther.Part_Shapes[i] == diagramstructure.Part_Shapes[j]
			},
			func(j int) string {
				return diagramstructure.Part_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if diagramstructure.IsPartsNodeExpanded != diagramstructureOther.IsPartsNodeExpanded {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "IsPartsNodeExpanded"))
	}
	PartWhoseNodeIsExpandedDifferent := false
	if len(diagramstructure.PartWhoseNodeIsExpanded) != len(diagramstructureOther.PartWhoseNodeIsExpanded) {
		PartWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramstructure.PartWhoseNodeIsExpanded {
			if (diagramstructure.PartWhoseNodeIsExpanded[i] == nil) != (diagramstructureOther.PartWhoseNodeIsExpanded[i] == nil) {
				PartWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramstructure.PartWhoseNodeIsExpanded[i] != nil && diagramstructureOther.PartWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.PartWhoseNodeIsExpanded[i] != diagramstructureOther.PartWhoseNodeIsExpanded[i] {
					PartWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if PartWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramstructure,
			"PartWhoseNodeIsExpanded",
			len(diagramstructureOther.PartWhoseNodeIsExpanded),
			len(diagramstructure.PartWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramstructureOther.PartWhoseNodeIsExpanded[i] == diagramstructure.PartWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramstructure.PartWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ExternalPart_ShapesDifferent := false
	if len(diagramstructure.ExternalPart_Shapes) != len(diagramstructureOther.ExternalPart_Shapes) {
		ExternalPart_ShapesDifferent = true
	} else {
		for i := range diagramstructure.ExternalPart_Shapes {
			if (diagramstructure.ExternalPart_Shapes[i] == nil) != (diagramstructureOther.ExternalPart_Shapes[i] == nil) {
				ExternalPart_ShapesDifferent = true
				break
			} else if diagramstructure.ExternalPart_Shapes[i] != nil && diagramstructureOther.ExternalPart_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.ExternalPart_Shapes[i] != diagramstructureOther.ExternalPart_Shapes[i] {
					ExternalPart_ShapesDifferent = true
					break
				}
			}
		}
	}
	if ExternalPart_ShapesDifferent {
		ops := stage.Diff(
			diagramstructure,
			"ExternalPart_Shapes",
			len(diagramstructureOther.ExternalPart_Shapes),
			len(diagramstructure.ExternalPart_Shapes),
			func(i, j int) bool {
				return diagramstructureOther.ExternalPart_Shapes[i] == diagramstructure.ExternalPart_Shapes[j]
			},
			func(j int) string {
				return diagramstructure.ExternalPart_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if diagramstructure.IsExternalPartsNodeExpanded != diagramstructureOther.IsExternalPartsNodeExpanded {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "IsExternalPartsNodeExpanded"))
	}
	ExternalPartWhoseNodeIsExpandedDifferent := false
	if len(diagramstructure.ExternalPartWhoseNodeIsExpanded) != len(diagramstructureOther.ExternalPartWhoseNodeIsExpanded) {
		ExternalPartWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramstructure.ExternalPartWhoseNodeIsExpanded {
			if (diagramstructure.ExternalPartWhoseNodeIsExpanded[i] == nil) != (diagramstructureOther.ExternalPartWhoseNodeIsExpanded[i] == nil) {
				ExternalPartWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramstructure.ExternalPartWhoseNodeIsExpanded[i] != nil && diagramstructureOther.ExternalPartWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.ExternalPartWhoseNodeIsExpanded[i] != diagramstructureOther.ExternalPartWhoseNodeIsExpanded[i] {
					ExternalPartWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ExternalPartWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramstructure,
			"ExternalPartWhoseNodeIsExpanded",
			len(diagramstructureOther.ExternalPartWhoseNodeIsExpanded),
			len(diagramstructure.ExternalPartWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramstructureOther.ExternalPartWhoseNodeIsExpanded[i] == diagramstructure.ExternalPartWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramstructure.ExternalPartWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ExternalPartsWhoseOutDataFlowsNodeIsExpandedDifferent := false
	if len(diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded) != len(diagramstructureOther.ExternalPartsWhoseOutDataFlowsNodeIsExpanded) {
		ExternalPartsWhoseOutDataFlowsNodeIsExpandedDifferent = true
	} else {
		for i := range diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
			if (diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded[i] == nil) != (diagramstructureOther.ExternalPartsWhoseOutDataFlowsNodeIsExpanded[i] == nil) {
				ExternalPartsWhoseOutDataFlowsNodeIsExpandedDifferent = true
				break
			} else if diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded[i] != nil && diagramstructureOther.ExternalPartsWhoseOutDataFlowsNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded[i] != diagramstructureOther.ExternalPartsWhoseOutDataFlowsNodeIsExpanded[i] {
					ExternalPartsWhoseOutDataFlowsNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ExternalPartsWhoseOutDataFlowsNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramstructure,
			"ExternalPartsWhoseOutDataFlowsNodeIsExpanded",
			len(diagramstructureOther.ExternalPartsWhoseOutDataFlowsNodeIsExpanded),
			len(diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded),
			func(i, j int) bool {
				return diagramstructureOther.ExternalPartsWhoseOutDataFlowsNodeIsExpanded[i] == diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ExternalPartsWhoseInDataFlowsNodeIsExpandedDifferent := false
	if len(diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded) != len(diagramstructureOther.ExternalPartsWhoseInDataFlowsNodeIsExpanded) {
		ExternalPartsWhoseInDataFlowsNodeIsExpandedDifferent = true
	} else {
		for i := range diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
			if (diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded[i] == nil) != (diagramstructureOther.ExternalPartsWhoseInDataFlowsNodeIsExpanded[i] == nil) {
				ExternalPartsWhoseInDataFlowsNodeIsExpandedDifferent = true
				break
			} else if diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded[i] != nil && diagramstructureOther.ExternalPartsWhoseInDataFlowsNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded[i] != diagramstructureOther.ExternalPartsWhoseInDataFlowsNodeIsExpanded[i] {
					ExternalPartsWhoseInDataFlowsNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ExternalPartsWhoseInDataFlowsNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramstructure,
			"ExternalPartsWhoseInDataFlowsNodeIsExpanded",
			len(diagramstructureOther.ExternalPartsWhoseInDataFlowsNodeIsExpanded),
			len(diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded),
			func(i, j int) bool {
				return diagramstructureOther.ExternalPartsWhoseInDataFlowsNodeIsExpanded[i] == diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	PortsWhoseNodeIsExpandedDifferent := false
	if len(diagramstructure.PortsWhoseNodeIsExpanded) != len(diagramstructureOther.PortsWhoseNodeIsExpanded) {
		PortsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramstructure.PortsWhoseNodeIsExpanded {
			if (diagramstructure.PortsWhoseNodeIsExpanded[i] == nil) != (diagramstructureOther.PortsWhoseNodeIsExpanded[i] == nil) {
				PortsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramstructure.PortsWhoseNodeIsExpanded[i] != nil && diagramstructureOther.PortsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.PortsWhoseNodeIsExpanded[i] != diagramstructureOther.PortsWhoseNodeIsExpanded[i] {
					PortsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if PortsWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramstructure,
			"PortsWhoseNodeIsExpanded",
			len(diagramstructureOther.PortsWhoseNodeIsExpanded),
			len(diagramstructure.PortsWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramstructureOther.PortsWhoseNodeIsExpanded[i] == diagramstructure.PortsWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramstructure.PortsWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	Port_ShapesDifferent := false
	if len(diagramstructure.Port_Shapes) != len(diagramstructureOther.Port_Shapes) {
		Port_ShapesDifferent = true
	} else {
		for i := range diagramstructure.Port_Shapes {
			if (diagramstructure.Port_Shapes[i] == nil) != (diagramstructureOther.Port_Shapes[i] == nil) {
				Port_ShapesDifferent = true
				break
			} else if diagramstructure.Port_Shapes[i] != nil && diagramstructureOther.Port_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.Port_Shapes[i] != diagramstructureOther.Port_Shapes[i] {
					Port_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Port_ShapesDifferent {
		ops := stage.Diff(
			diagramstructure,
			"Port_Shapes",
			len(diagramstructureOther.Port_Shapes),
			len(diagramstructure.Port_Shapes),
			func(i, j int) bool {
				return diagramstructureOther.Port_Shapes[i] == diagramstructure.Port_Shapes[j]
			},
			func(j int) string {
				return diagramstructure.Port_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ControlFlowsWhoseNodeIsExpandedDifferent := false
	if len(diagramstructure.ControlFlowsWhoseNodeIsExpanded) != len(diagramstructureOther.ControlFlowsWhoseNodeIsExpanded) {
		ControlFlowsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramstructure.ControlFlowsWhoseNodeIsExpanded {
			if (diagramstructure.ControlFlowsWhoseNodeIsExpanded[i] == nil) != (diagramstructureOther.ControlFlowsWhoseNodeIsExpanded[i] == nil) {
				ControlFlowsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramstructure.ControlFlowsWhoseNodeIsExpanded[i] != nil && diagramstructureOther.ControlFlowsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.ControlFlowsWhoseNodeIsExpanded[i] != diagramstructureOther.ControlFlowsWhoseNodeIsExpanded[i] {
					ControlFlowsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ControlFlowsWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramstructure,
			"ControlFlowsWhoseNodeIsExpanded",
			len(diagramstructureOther.ControlFlowsWhoseNodeIsExpanded),
			len(diagramstructure.ControlFlowsWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramstructureOther.ControlFlowsWhoseNodeIsExpanded[i] == diagramstructure.ControlFlowsWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramstructure.ControlFlowsWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ControlFlow_ShapesDifferent := false
	if len(diagramstructure.ControlFlow_Shapes) != len(diagramstructureOther.ControlFlow_Shapes) {
		ControlFlow_ShapesDifferent = true
	} else {
		for i := range diagramstructure.ControlFlow_Shapes {
			if (diagramstructure.ControlFlow_Shapes[i] == nil) != (diagramstructureOther.ControlFlow_Shapes[i] == nil) {
				ControlFlow_ShapesDifferent = true
				break
			} else if diagramstructure.ControlFlow_Shapes[i] != nil && diagramstructureOther.ControlFlow_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.ControlFlow_Shapes[i] != diagramstructureOther.ControlFlow_Shapes[i] {
					ControlFlow_ShapesDifferent = true
					break
				}
			}
		}
	}
	if ControlFlow_ShapesDifferent {
		ops := stage.Diff(
			diagramstructure,
			"ControlFlow_Shapes",
			len(diagramstructureOther.ControlFlow_Shapes),
			len(diagramstructure.ControlFlow_Shapes),
			func(i, j int) bool {
				return diagramstructureOther.ControlFlow_Shapes[i] == diagramstructure.ControlFlow_Shapes[j]
			},
			func(j int) string {
				return diagramstructure.ControlFlow_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DataFlowsWhoseNodeIsExpandedDifferent := false
	if len(diagramstructure.DataFlowsWhoseNodeIsExpanded) != len(diagramstructureOther.DataFlowsWhoseNodeIsExpanded) {
		DataFlowsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramstructure.DataFlowsWhoseNodeIsExpanded {
			if (diagramstructure.DataFlowsWhoseNodeIsExpanded[i] == nil) != (diagramstructureOther.DataFlowsWhoseNodeIsExpanded[i] == nil) {
				DataFlowsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramstructure.DataFlowsWhoseNodeIsExpanded[i] != nil && diagramstructureOther.DataFlowsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.DataFlowsWhoseNodeIsExpanded[i] != diagramstructureOther.DataFlowsWhoseNodeIsExpanded[i] {
					DataFlowsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if DataFlowsWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramstructure,
			"DataFlowsWhoseNodeIsExpanded",
			len(diagramstructureOther.DataFlowsWhoseNodeIsExpanded),
			len(diagramstructure.DataFlowsWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramstructureOther.DataFlowsWhoseNodeIsExpanded[i] == diagramstructure.DataFlowsWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramstructure.DataFlowsWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DataFlow_ShapesDifferent := false
	if len(diagramstructure.DataFlow_Shapes) != len(diagramstructureOther.DataFlow_Shapes) {
		DataFlow_ShapesDifferent = true
	} else {
		for i := range diagramstructure.DataFlow_Shapes {
			if (diagramstructure.DataFlow_Shapes[i] == nil) != (diagramstructureOther.DataFlow_Shapes[i] == nil) {
				DataFlow_ShapesDifferent = true
				break
			} else if diagramstructure.DataFlow_Shapes[i] != nil && diagramstructureOther.DataFlow_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.DataFlow_Shapes[i] != diagramstructureOther.DataFlow_Shapes[i] {
					DataFlow_ShapesDifferent = true
					break
				}
			}
		}
	}
	if DataFlow_ShapesDifferent {
		ops := stage.Diff(
			diagramstructure,
			"DataFlow_Shapes",
			len(diagramstructureOther.DataFlow_Shapes),
			len(diagramstructure.DataFlow_Shapes),
			func(i, j int) bool {
				return diagramstructureOther.DataFlow_Shapes[i] == diagramstructure.DataFlow_Shapes[j]
			},
			func(j int) string {
				return diagramstructure.DataFlow_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DatasWhoseNodeIsExpandedDifferent := false
	if len(diagramstructure.DatasWhoseNodeIsExpanded) != len(diagramstructureOther.DatasWhoseNodeIsExpanded) {
		DatasWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramstructure.DatasWhoseNodeIsExpanded {
			if (diagramstructure.DatasWhoseNodeIsExpanded[i] == nil) != (diagramstructureOther.DatasWhoseNodeIsExpanded[i] == nil) {
				DatasWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramstructure.DatasWhoseNodeIsExpanded[i] != nil && diagramstructureOther.DatasWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.DatasWhoseNodeIsExpanded[i] != diagramstructureOther.DatasWhoseNodeIsExpanded[i] {
					DatasWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if DatasWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramstructure,
			"DatasWhoseNodeIsExpanded",
			len(diagramstructureOther.DatasWhoseNodeIsExpanded),
			len(diagramstructure.DatasWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramstructureOther.DatasWhoseNodeIsExpanded[i] == diagramstructure.DatasWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramstructure.DatasWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	Data_ShapesDifferent := false
	if len(diagramstructure.Data_Shapes) != len(diagramstructureOther.Data_Shapes) {
		Data_ShapesDifferent = true
	} else {
		for i := range diagramstructure.Data_Shapes {
			if (diagramstructure.Data_Shapes[i] == nil) != (diagramstructureOther.Data_Shapes[i] == nil) {
				Data_ShapesDifferent = true
				break
			} else if diagramstructure.Data_Shapes[i] != nil && diagramstructureOther.Data_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.Data_Shapes[i] != diagramstructureOther.Data_Shapes[i] {
					Data_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Data_ShapesDifferent {
		ops := stage.Diff(
			diagramstructure,
			"Data_Shapes",
			len(diagramstructureOther.Data_Shapes),
			len(diagramstructure.Data_Shapes),
			func(i, j int) bool {
				return diagramstructureOther.Data_Shapes[i] == diagramstructure.Data_Shapes[j]
			},
			func(j int) string {
				return diagramstructure.Data_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DataFlowsWhoseDataNodeIsExpandedDifferent := false
	if len(diagramstructure.DataFlowsWhoseDataNodeIsExpanded) != len(diagramstructureOther.DataFlowsWhoseDataNodeIsExpanded) {
		DataFlowsWhoseDataNodeIsExpandedDifferent = true
	} else {
		for i := range diagramstructure.DataFlowsWhoseDataNodeIsExpanded {
			if (diagramstructure.DataFlowsWhoseDataNodeIsExpanded[i] == nil) != (diagramstructureOther.DataFlowsWhoseDataNodeIsExpanded[i] == nil) {
				DataFlowsWhoseDataNodeIsExpandedDifferent = true
				break
			} else if diagramstructure.DataFlowsWhoseDataNodeIsExpanded[i] != nil && diagramstructureOther.DataFlowsWhoseDataNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.DataFlowsWhoseDataNodeIsExpanded[i] != diagramstructureOther.DataFlowsWhoseDataNodeIsExpanded[i] {
					DataFlowsWhoseDataNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if DataFlowsWhoseDataNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramstructure,
			"DataFlowsWhoseDataNodeIsExpanded",
			len(diagramstructureOther.DataFlowsWhoseDataNodeIsExpanded),
			len(diagramstructure.DataFlowsWhoseDataNodeIsExpanded),
			func(i, j int) bool {
				return diagramstructureOther.DataFlowsWhoseDataNodeIsExpanded[i] == diagramstructure.DataFlowsWhoseDataNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramstructure.DataFlowsWhoseDataNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	AllocatedResourcesWhoseNodeIsExpandedDifferent := false
	if len(diagramstructure.AllocatedResourcesWhoseNodeIsExpanded) != len(diagramstructureOther.AllocatedResourcesWhoseNodeIsExpanded) {
		AllocatedResourcesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramstructure.AllocatedResourcesWhoseNodeIsExpanded {
			if (diagramstructure.AllocatedResourcesWhoseNodeIsExpanded[i] == nil) != (diagramstructureOther.AllocatedResourcesWhoseNodeIsExpanded[i] == nil) {
				AllocatedResourcesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramstructure.AllocatedResourcesWhoseNodeIsExpanded[i] != nil && diagramstructureOther.AllocatedResourcesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.AllocatedResourcesWhoseNodeIsExpanded[i] != diagramstructureOther.AllocatedResourcesWhoseNodeIsExpanded[i] {
					AllocatedResourcesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if AllocatedResourcesWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramstructure,
			"AllocatedResourcesWhoseNodeIsExpanded",
			len(diagramstructureOther.AllocatedResourcesWhoseNodeIsExpanded),
			len(diagramstructure.AllocatedResourcesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramstructureOther.AllocatedResourcesWhoseNodeIsExpanded[i] == diagramstructure.AllocatedResourcesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramstructure.AllocatedResourcesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	AllocatedResourceShapesDifferent := false
	if len(diagramstructure.AllocatedResourceShapes) != len(diagramstructureOther.AllocatedResourceShapes) {
		AllocatedResourceShapesDifferent = true
	} else {
		for i := range diagramstructure.AllocatedResourceShapes {
			if (diagramstructure.AllocatedResourceShapes[i] == nil) != (diagramstructureOther.AllocatedResourceShapes[i] == nil) {
				AllocatedResourceShapesDifferent = true
				break
			} else if diagramstructure.AllocatedResourceShapes[i] != nil && diagramstructureOther.AllocatedResourceShapes[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.AllocatedResourceShapes[i] != diagramstructureOther.AllocatedResourceShapes[i] {
					AllocatedResourceShapesDifferent = true
					break
				}
			}
		}
	}
	if AllocatedResourceShapesDifferent {
		ops := stage.Diff(
			diagramstructure,
			"AllocatedResourceShapes",
			len(diagramstructureOther.AllocatedResourceShapes),
			len(diagramstructure.AllocatedResourceShapes),
			func(i, j int) bool {
				return diagramstructureOther.AllocatedResourceShapes[i] == diagramstructure.AllocatedResourceShapes[j]
			},
			func(j int) string {
				return diagramstructure.AllocatedResourceShapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	AllocatedSystemesWhoseNodeIsExpandedDifferent := false
	if len(diagramstructure.AllocatedSystemesWhoseNodeIsExpanded) != len(diagramstructureOther.AllocatedSystemesWhoseNodeIsExpanded) {
		AllocatedSystemesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramstructure.AllocatedSystemesWhoseNodeIsExpanded {
			if (diagramstructure.AllocatedSystemesWhoseNodeIsExpanded[i] == nil) != (diagramstructureOther.AllocatedSystemesWhoseNodeIsExpanded[i] == nil) {
				AllocatedSystemesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramstructure.AllocatedSystemesWhoseNodeIsExpanded[i] != nil && diagramstructureOther.AllocatedSystemesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.AllocatedSystemesWhoseNodeIsExpanded[i] != diagramstructureOther.AllocatedSystemesWhoseNodeIsExpanded[i] {
					AllocatedSystemesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if AllocatedSystemesWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramstructure,
			"AllocatedSystemesWhoseNodeIsExpanded",
			len(diagramstructureOther.AllocatedSystemesWhoseNodeIsExpanded),
			len(diagramstructure.AllocatedSystemesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramstructureOther.AllocatedSystemesWhoseNodeIsExpanded[i] == diagramstructure.AllocatedSystemesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramstructure.AllocatedSystemesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	AllocatedSystemShapesDifferent := false
	if len(diagramstructure.AllocatedSystemShapes) != len(diagramstructureOther.AllocatedSystemShapes) {
		AllocatedSystemShapesDifferent = true
	} else {
		for i := range diagramstructure.AllocatedSystemShapes {
			if (diagramstructure.AllocatedSystemShapes[i] == nil) != (diagramstructureOther.AllocatedSystemShapes[i] == nil) {
				AllocatedSystemShapesDifferent = true
				break
			} else if diagramstructure.AllocatedSystemShapes[i] != nil && diagramstructureOther.AllocatedSystemShapes[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.AllocatedSystemShapes[i] != diagramstructureOther.AllocatedSystemShapes[i] {
					AllocatedSystemShapesDifferent = true
					break
				}
			}
		}
	}
	if AllocatedSystemShapesDifferent {
		ops := stage.Diff(
			diagramstructure,
			"AllocatedSystemShapes",
			len(diagramstructureOther.AllocatedSystemShapes),
			len(diagramstructure.AllocatedSystemShapes),
			func(i, j int) bool {
				return diagramstructureOther.AllocatedSystemShapes[i] == diagramstructure.AllocatedSystemShapes[j]
			},
			func(j int) string {
				return diagramstructure.AllocatedSystemShapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	Note_ShapesDifferent := false
	if len(diagramstructure.Note_Shapes) != len(diagramstructureOther.Note_Shapes) {
		Note_ShapesDifferent = true
	} else {
		for i := range diagramstructure.Note_Shapes {
			if (diagramstructure.Note_Shapes[i] == nil) != (diagramstructureOther.Note_Shapes[i] == nil) {
				Note_ShapesDifferent = true
				break
			} else if diagramstructure.Note_Shapes[i] != nil && diagramstructureOther.Note_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.Note_Shapes[i] != diagramstructureOther.Note_Shapes[i] {
					Note_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Note_ShapesDifferent {
		ops := stage.Diff(
			diagramstructure,
			"Note_Shapes",
			len(diagramstructureOther.Note_Shapes),
			len(diagramstructure.Note_Shapes),
			func(i, j int) bool {
				return diagramstructureOther.Note_Shapes[i] == diagramstructure.Note_Shapes[j]
			},
			func(j int) string {
				return diagramstructure.Note_Shapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	NotesWhoseNodeIsExpandedDifferent := false
	if len(diagramstructure.NotesWhoseNodeIsExpanded) != len(diagramstructureOther.NotesWhoseNodeIsExpanded) {
		NotesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramstructure.NotesWhoseNodeIsExpanded {
			if (diagramstructure.NotesWhoseNodeIsExpanded[i] == nil) != (diagramstructureOther.NotesWhoseNodeIsExpanded[i] == nil) {
				NotesWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramstructure.NotesWhoseNodeIsExpanded[i] != nil && diagramstructureOther.NotesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.NotesWhoseNodeIsExpanded[i] != diagramstructureOther.NotesWhoseNodeIsExpanded[i] {
					NotesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if NotesWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			diagramstructure,
			"NotesWhoseNodeIsExpanded",
			len(diagramstructureOther.NotesWhoseNodeIsExpanded),
			len(diagramstructure.NotesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return diagramstructureOther.NotesWhoseNodeIsExpanded[i] == diagramstructure.NotesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return diagramstructure.NotesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if diagramstructure.IsNotesNodeExpanded != diagramstructureOther.IsNotesNodeExpanded {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	NotePortShapesDifferent := false
	if len(diagramstructure.NotePortShapes) != len(diagramstructureOther.NotePortShapes) {
		NotePortShapesDifferent = true
	} else {
		for i := range diagramstructure.NotePortShapes {
			if (diagramstructure.NotePortShapes[i] == nil) != (diagramstructureOther.NotePortShapes[i] == nil) {
				NotePortShapesDifferent = true
				break
			} else if diagramstructure.NotePortShapes[i] != nil && diagramstructureOther.NotePortShapes[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.NotePortShapes[i] != diagramstructureOther.NotePortShapes[i] {
					NotePortShapesDifferent = true
					break
				}
			}
		}
	}
	if NotePortShapesDifferent {
		ops := stage.Diff(
			diagramstructure,
			"NotePortShapes",
			len(diagramstructureOther.NotePortShapes),
			len(diagramstructure.NotePortShapes),
			func(i, j int) bool {
				return diagramstructureOther.NotePortShapes[i] == diagramstructure.NotePortShapes[j]
			},
			func(j int) string {
				return diagramstructure.NotePortShapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	NotePartShapesDifferent := false
	if len(diagramstructure.NotePartShapes) != len(diagramstructureOther.NotePartShapes) {
		NotePartShapesDifferent = true
	} else {
		for i := range diagramstructure.NotePartShapes {
			if (diagramstructure.NotePartShapes[i] == nil) != (diagramstructureOther.NotePartShapes[i] == nil) {
				NotePartShapesDifferent = true
				break
			} else if diagramstructure.NotePartShapes[i] != nil && diagramstructureOther.NotePartShapes[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.NotePartShapes[i] != diagramstructureOther.NotePartShapes[i] {
					NotePartShapesDifferent = true
					break
				}
			}
		}
	}
	if NotePartShapesDifferent {
		ops := stage.Diff(
			diagramstructure,
			"NotePartShapes",
			len(diagramstructureOther.NotePartShapes),
			len(diagramstructure.NotePartShapes),
			func(i, j int) bool {
				return diagramstructureOther.NotePartShapes[i] == diagramstructure.NotePartShapes[j]
			},
			func(j int) string {
				return diagramstructure.NotePartShapes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (externalpartshape *ExternalPartShape) GongDiff(stage *Stage, externalpartshapeOther *ExternalPartShape) (diffs []string) {
	// insertion point for field diffs
	if externalpartshape.Name != externalpartshapeOther.Name {
		diffs = append(diffs, externalpartshape.GongMarshallField(stage, "Name"))
	}
	if (externalpartshape.Part == nil) != (externalpartshapeOther.Part == nil) {
		diffs = append(diffs, externalpartshape.GongMarshallField(stage, "Part"))
	} else if externalpartshape.Part != nil && externalpartshapeOther.Part != nil {
		if externalpartshape.Part != externalpartshapeOther.Part {
			diffs = append(diffs, externalpartshape.GongMarshallField(stage, "Part"))
		}
	}
	if externalpartshape.IsExpanded != externalpartshapeOther.IsExpanded {
		diffs = append(diffs, externalpartshape.GongMarshallField(stage, "IsExpanded"))
	}
	if externalpartshape.X != externalpartshapeOther.X {
		diffs = append(diffs, externalpartshape.GongMarshallField(stage, "X"))
	}
	if externalpartshape.Y != externalpartshapeOther.Y {
		diffs = append(diffs, externalpartshape.GongMarshallField(stage, "Y"))
	}
	if externalpartshape.Width != externalpartshapeOther.Width {
		diffs = append(diffs, externalpartshape.GongMarshallField(stage, "Width"))
	}
	if externalpartshape.Height != externalpartshapeOther.Height {
		diffs = append(diffs, externalpartshape.GongMarshallField(stage, "Height"))
	}
	if externalpartshape.IsHidden != externalpartshapeOther.IsHidden {
		diffs = append(diffs, externalpartshape.GongMarshallField(stage, "IsHidden"))
	}
	if externalpartshape.TailHeigth != externalpartshapeOther.TailHeigth {
		diffs = append(diffs, externalpartshape.GongMarshallField(stage, "TailHeigth"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (layerdefinition *LayerDefinition) GongDiff(stage *Stage, layerdefinitionOther *LayerDefinition) (diffs []string) {
	// insertion point for field diffs
	if layerdefinition.Name != layerdefinitionOther.Name {
		diffs = append(diffs, layerdefinition.GongMarshallField(stage, "Name"))
	}
	QueryDifferent := false
	if len(layerdefinition.Query) != len(layerdefinitionOther.Query) {
		QueryDifferent = true
	} else {
		for i := range layerdefinition.Query {
			if (layerdefinition.Query[i] == nil) != (layerdefinitionOther.Query[i] == nil) {
				QueryDifferent = true
				break
			} else if layerdefinition.Query[i] != nil && layerdefinitionOther.Query[i] != nil {
				// this is a pointer comparaison
				if layerdefinition.Query[i] != layerdefinitionOther.Query[i] {
					QueryDifferent = true
					break
				}
			}
		}
	}
	if QueryDifferent {
		ops := stage.Diff(
			layerdefinition,
			"Query",
			len(layerdefinitionOther.Query),
			len(layerdefinition.Query),
			func(i, j int) bool {
				return layerdefinitionOther.Query[i] == layerdefinition.Query[j]
			},
			func(j int) string {
				return layerdefinition.Query[j].GongGetIdentifier(stage)
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
	RootSystemesDifferent := false
	if len(library.RootSystemes) != len(libraryOther.RootSystemes) {
		RootSystemesDifferent = true
	} else {
		for i := range library.RootSystemes {
			if (library.RootSystemes[i] == nil) != (libraryOther.RootSystemes[i] == nil) {
				RootSystemesDifferent = true
				break
			} else if library.RootSystemes[i] != nil && libraryOther.RootSystemes[i] != nil {
				// this is a pointer comparaison
				if library.RootSystemes[i] != libraryOther.RootSystemes[i] {
					RootSystemesDifferent = true
					break
				}
			}
		}
	}
	if RootSystemesDifferent {
		ops := stage.Diff(
			library,
			"RootSystemes",
			len(libraryOther.RootSystemes),
			len(library.RootSystemes),
			func(i, j int) bool {
				return libraryOther.RootSystemes[i] == library.RootSystemes[j]
			},
			func(j int) string {
				return library.RootSystemes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if library.IsSystemesNodeExpanded != libraryOther.IsSystemesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsSystemesNodeExpanded"))
	}
	SystemsWhoseNodeIsExpandedDifferent := false
	if len(library.SystemsWhoseNodeIsExpanded) != len(libraryOther.SystemsWhoseNodeIsExpanded) {
		SystemsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.SystemsWhoseNodeIsExpanded {
			if (library.SystemsWhoseNodeIsExpanded[i] == nil) != (libraryOther.SystemsWhoseNodeIsExpanded[i] == nil) {
				SystemsWhoseNodeIsExpandedDifferent = true
				break
			} else if library.SystemsWhoseNodeIsExpanded[i] != nil && libraryOther.SystemsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.SystemsWhoseNodeIsExpanded[i] != libraryOther.SystemsWhoseNodeIsExpanded[i] {
					SystemsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if SystemsWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			library,
			"SystemsWhoseNodeIsExpanded",
			len(libraryOther.SystemsWhoseNodeIsExpanded),
			len(library.SystemsWhoseNodeIsExpanded),
			func(i, j int) bool {
				return libraryOther.SystemsWhoseNodeIsExpanded[i] == library.SystemsWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return library.SystemsWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	RootDataFlowsDifferent := false
	if len(library.RootDataFlows) != len(libraryOther.RootDataFlows) {
		RootDataFlowsDifferent = true
	} else {
		for i := range library.RootDataFlows {
			if (library.RootDataFlows[i] == nil) != (libraryOther.RootDataFlows[i] == nil) {
				RootDataFlowsDifferent = true
				break
			} else if library.RootDataFlows[i] != nil && libraryOther.RootDataFlows[i] != nil {
				// this is a pointer comparaison
				if library.RootDataFlows[i] != libraryOther.RootDataFlows[i] {
					RootDataFlowsDifferent = true
					break
				}
			}
		}
	}
	if RootDataFlowsDifferent {
		ops := stage.Diff(
			library,
			"RootDataFlows",
			len(libraryOther.RootDataFlows),
			len(library.RootDataFlows),
			func(i, j int) bool {
				return libraryOther.RootDataFlows[i] == library.RootDataFlows[j]
			},
			func(j int) string {
				return library.RootDataFlows[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if library.IsDataFlowsNodeExpanded != libraryOther.IsDataFlowsNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
	}
	DataFlowsWhoseNodeIsExpandedDifferent := false
	if len(library.DataFlowsWhoseNodeIsExpanded) != len(libraryOther.DataFlowsWhoseNodeIsExpanded) {
		DataFlowsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.DataFlowsWhoseNodeIsExpanded {
			if (library.DataFlowsWhoseNodeIsExpanded[i] == nil) != (libraryOther.DataFlowsWhoseNodeIsExpanded[i] == nil) {
				DataFlowsWhoseNodeIsExpandedDifferent = true
				break
			} else if library.DataFlowsWhoseNodeIsExpanded[i] != nil && libraryOther.DataFlowsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.DataFlowsWhoseNodeIsExpanded[i] != libraryOther.DataFlowsWhoseNodeIsExpanded[i] {
					DataFlowsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if DataFlowsWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			library,
			"DataFlowsWhoseNodeIsExpanded",
			len(libraryOther.DataFlowsWhoseNodeIsExpanded),
			len(library.DataFlowsWhoseNodeIsExpanded),
			func(i, j int) bool {
				return libraryOther.DataFlowsWhoseNodeIsExpanded[i] == library.DataFlowsWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return library.DataFlowsWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	RootDatasDifferent := false
	if len(library.RootDatas) != len(libraryOther.RootDatas) {
		RootDatasDifferent = true
	} else {
		for i := range library.RootDatas {
			if (library.RootDatas[i] == nil) != (libraryOther.RootDatas[i] == nil) {
				RootDatasDifferent = true
				break
			} else if library.RootDatas[i] != nil && libraryOther.RootDatas[i] != nil {
				// this is a pointer comparaison
				if library.RootDatas[i] != libraryOther.RootDatas[i] {
					RootDatasDifferent = true
					break
				}
			}
		}
	}
	if RootDatasDifferent {
		ops := stage.Diff(
			library,
			"RootDatas",
			len(libraryOther.RootDatas),
			len(library.RootDatas),
			func(i, j int) bool {
				return libraryOther.RootDatas[i] == library.RootDatas[j]
			},
			func(j int) string {
				return library.RootDatas[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if library.IsDatasNodeExpanded != libraryOther.IsDatasNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsDatasNodeExpanded"))
	}
	DatasWhoseNodeIsExpandedDifferent := false
	if len(library.DatasWhoseNodeIsExpanded) != len(libraryOther.DatasWhoseNodeIsExpanded) {
		DatasWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.DatasWhoseNodeIsExpanded {
			if (library.DatasWhoseNodeIsExpanded[i] == nil) != (libraryOther.DatasWhoseNodeIsExpanded[i] == nil) {
				DatasWhoseNodeIsExpandedDifferent = true
				break
			} else if library.DatasWhoseNodeIsExpanded[i] != nil && libraryOther.DatasWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.DatasWhoseNodeIsExpanded[i] != libraryOther.DatasWhoseNodeIsExpanded[i] {
					DatasWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if DatasWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			library,
			"DatasWhoseNodeIsExpanded",
			len(libraryOther.DatasWhoseNodeIsExpanded),
			len(library.DatasWhoseNodeIsExpanded),
			func(i, j int) bool {
				return libraryOther.DatasWhoseNodeIsExpanded[i] == library.DatasWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return library.DatasWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	RootResourcesDifferent := false
	if len(library.RootResources) != len(libraryOther.RootResources) {
		RootResourcesDifferent = true
	} else {
		for i := range library.RootResources {
			if (library.RootResources[i] == nil) != (libraryOther.RootResources[i] == nil) {
				RootResourcesDifferent = true
				break
			} else if library.RootResources[i] != nil && libraryOther.RootResources[i] != nil {
				// this is a pointer comparaison
				if library.RootResources[i] != libraryOther.RootResources[i] {
					RootResourcesDifferent = true
					break
				}
			}
		}
	}
	if RootResourcesDifferent {
		ops := stage.Diff(
			library,
			"RootResources",
			len(libraryOther.RootResources),
			len(library.RootResources),
			func(i, j int) bool {
				return libraryOther.RootResources[i] == library.RootResources[j]
			},
			func(j int) string {
				return library.RootResources[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if library.IsResourcesNodeExpanded != libraryOther.IsResourcesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsResourcesNodeExpanded"))
	}
	ResourcesWhoseNodeIsExpandedDifferent := false
	if len(library.ResourcesWhoseNodeIsExpanded) != len(libraryOther.ResourcesWhoseNodeIsExpanded) {
		ResourcesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.ResourcesWhoseNodeIsExpanded {
			if (library.ResourcesWhoseNodeIsExpanded[i] == nil) != (libraryOther.ResourcesWhoseNodeIsExpanded[i] == nil) {
				ResourcesWhoseNodeIsExpandedDifferent = true
				break
			} else if library.ResourcesWhoseNodeIsExpanded[i] != nil && libraryOther.ResourcesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.ResourcesWhoseNodeIsExpanded[i] != libraryOther.ResourcesWhoseNodeIsExpanded[i] {
					ResourcesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ResourcesWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			library,
			"ResourcesWhoseNodeIsExpanded",
			len(libraryOther.ResourcesWhoseNodeIsExpanded),
			len(library.ResourcesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return libraryOther.ResourcesWhoseNodeIsExpanded[i] == library.ResourcesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return library.ResourcesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	PartsWhoseNodeIsExpandedDifferent := false
	if len(library.PartsWhoseNodeIsExpanded) != len(libraryOther.PartsWhoseNodeIsExpanded) {
		PartsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.PartsWhoseNodeIsExpanded {
			if (library.PartsWhoseNodeIsExpanded[i] == nil) != (libraryOther.PartsWhoseNodeIsExpanded[i] == nil) {
				PartsWhoseNodeIsExpandedDifferent = true
				break
			} else if library.PartsWhoseNodeIsExpanded[i] != nil && libraryOther.PartsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.PartsWhoseNodeIsExpanded[i] != libraryOther.PartsWhoseNodeIsExpanded[i] {
					PartsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if PartsWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			library,
			"PartsWhoseNodeIsExpanded",
			len(libraryOther.PartsWhoseNodeIsExpanded),
			len(library.PartsWhoseNodeIsExpanded),
			func(i, j int) bool {
				return libraryOther.PartsWhoseNodeIsExpanded[i] == library.PartsWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return library.PartsWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	RootNotesDifferent := false
	if len(library.RootNotes) != len(libraryOther.RootNotes) {
		RootNotesDifferent = true
	} else {
		for i := range library.RootNotes {
			if (library.RootNotes[i] == nil) != (libraryOther.RootNotes[i] == nil) {
				RootNotesDifferent = true
				break
			} else if library.RootNotes[i] != nil && libraryOther.RootNotes[i] != nil {
				// this is a pointer comparaison
				if library.RootNotes[i] != libraryOther.RootNotes[i] {
					RootNotesDifferent = true
					break
				}
			}
		}
	}
	if RootNotesDifferent {
		ops := stage.Diff(
			library,
			"RootNotes",
			len(libraryOther.RootNotes),
			len(library.RootNotes),
			func(i, j int) bool {
				return libraryOther.RootNotes[i] == library.RootNotes[j]
			},
			func(j int) string {
				return library.RootNotes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if library.IsNotesNodeExpanded != libraryOther.IsNotesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	NotesWhoseNodeIsExpandedDifferent := false
	if len(library.NotesWhoseNodeIsExpanded) != len(libraryOther.NotesWhoseNodeIsExpanded) {
		NotesWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range library.NotesWhoseNodeIsExpanded {
			if (library.NotesWhoseNodeIsExpanded[i] == nil) != (libraryOther.NotesWhoseNodeIsExpanded[i] == nil) {
				NotesWhoseNodeIsExpandedDifferent = true
				break
			} else if library.NotesWhoseNodeIsExpanded[i] != nil && libraryOther.NotesWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if library.NotesWhoseNodeIsExpanded[i] != libraryOther.NotesWhoseNodeIsExpanded[i] {
					NotesWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if NotesWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			library,
			"NotesWhoseNodeIsExpanded",
			len(libraryOther.NotesWhoseNodeIsExpanded),
			len(library.NotesWhoseNodeIsExpanded),
			func(i, j int) bool {
				return libraryOther.NotesWhoseNodeIsExpanded[i] == library.NotesWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return library.NotesWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if library.IsExpandedTmp != libraryOther.IsExpandedTmp {
		diffs = append(diffs, library.GongMarshallField(stage, "IsExpandedTmp"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (note *Note) GongDiff(stage *Stage, noteOther *Note) (diffs []string) {
	// insertion point for field diffs
	if note.Name != noteOther.Name {
		diffs = append(diffs, note.GongMarshallField(stage, "Name"))
	}
	if note.Description != noteOther.Description {
		diffs = append(diffs, note.GongMarshallField(stage, "Description"))
	}
	if note.ComputedPrefix != noteOther.ComputedPrefix {
		diffs = append(diffs, note.GongMarshallField(stage, "ComputedPrefix"))
	}
	if note.IsExpanded != noteOther.IsExpanded {
		diffs = append(diffs, note.GongMarshallField(stage, "IsExpanded"))
	}
	if note.IsPartsNodeExpanded != noteOther.IsPartsNodeExpanded {
		diffs = append(diffs, note.GongMarshallField(stage, "IsPartsNodeExpanded"))
	}
	PartsDifferent := false
	if len(note.Parts) != len(noteOther.Parts) {
		PartsDifferent = true
	} else {
		for i := range note.Parts {
			if (note.Parts[i] == nil) != (noteOther.Parts[i] == nil) {
				PartsDifferent = true
				break
			} else if note.Parts[i] != nil && noteOther.Parts[i] != nil {
				// this is a pointer comparaison
				if note.Parts[i] != noteOther.Parts[i] {
					PartsDifferent = true
					break
				}
			}
		}
	}
	if PartsDifferent {
		ops := stage.Diff(
			note,
			"Parts",
			len(noteOther.Parts),
			len(note.Parts),
			func(i, j int) bool {
				return noteOther.Parts[i] == note.Parts[j]
			},
			func(j int) string {
				return note.Parts[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if note.IsPortsNodeExpanded != noteOther.IsPortsNodeExpanded {
		diffs = append(diffs, note.GongMarshallField(stage, "IsPortsNodeExpanded"))
	}
	PortsDifferent := false
	if len(note.Ports) != len(noteOther.Ports) {
		PortsDifferent = true
	} else {
		for i := range note.Ports {
			if (note.Ports[i] == nil) != (noteOther.Ports[i] == nil) {
				PortsDifferent = true
				break
			} else if note.Ports[i] != nil && noteOther.Ports[i] != nil {
				// this is a pointer comparaison
				if note.Ports[i] != noteOther.Ports[i] {
					PortsDifferent = true
					break
				}
			}
		}
	}
	if PortsDifferent {
		ops := stage.Diff(
			note,
			"Ports",
			len(noteOther.Ports),
			len(note.Ports),
			func(i, j int) bool {
				return noteOther.Ports[i] == note.Ports[j]
			},
			func(j int) string {
				return note.Ports[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (notepartshape *NotePartShape) GongDiff(stage *Stage, notepartshapeOther *NotePartShape) (diffs []string) {
	// insertion point for field diffs
	if notepartshape.Name != notepartshapeOther.Name {
		diffs = append(diffs, notepartshape.GongMarshallField(stage, "Name"))
	}
	if (notepartshape.Note == nil) != (notepartshapeOther.Note == nil) {
		diffs = append(diffs, notepartshape.GongMarshallField(stage, "Note"))
	} else if notepartshape.Note != nil && notepartshapeOther.Note != nil {
		if notepartshape.Note != notepartshapeOther.Note {
			diffs = append(diffs, notepartshape.GongMarshallField(stage, "Note"))
		}
	}
	if (notepartshape.Part == nil) != (notepartshapeOther.Part == nil) {
		diffs = append(diffs, notepartshape.GongMarshallField(stage, "Part"))
	} else if notepartshape.Part != nil && notepartshapeOther.Part != nil {
		if notepartshape.Part != notepartshapeOther.Part {
			diffs = append(diffs, notepartshape.GongMarshallField(stage, "Part"))
		}
	}
	if notepartshape.StartRatio != notepartshapeOther.StartRatio {
		diffs = append(diffs, notepartshape.GongMarshallField(stage, "StartRatio"))
	}
	if notepartshape.EndRatio != notepartshapeOther.EndRatio {
		diffs = append(diffs, notepartshape.GongMarshallField(stage, "EndRatio"))
	}
	if notepartshape.StartOrientation != notepartshapeOther.StartOrientation {
		diffs = append(diffs, notepartshape.GongMarshallField(stage, "StartOrientation"))
	}
	if notepartshape.EndOrientation != notepartshapeOther.EndOrientation {
		diffs = append(diffs, notepartshape.GongMarshallField(stage, "EndOrientation"))
	}
	if notepartshape.CornerOffsetRatio != notepartshapeOther.CornerOffsetRatio {
		diffs = append(diffs, notepartshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if notepartshape.IsHidden != notepartshapeOther.IsHidden {
		diffs = append(diffs, notepartshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (noteportshape *NotePortShape) GongDiff(stage *Stage, noteportshapeOther *NotePortShape) (diffs []string) {
	// insertion point for field diffs
	if noteportshape.Name != noteportshapeOther.Name {
		diffs = append(diffs, noteportshape.GongMarshallField(stage, "Name"))
	}
	if (noteportshape.Note == nil) != (noteportshapeOther.Note == nil) {
		diffs = append(diffs, noteportshape.GongMarshallField(stage, "Note"))
	} else if noteportshape.Note != nil && noteportshapeOther.Note != nil {
		if noteportshape.Note != noteportshapeOther.Note {
			diffs = append(diffs, noteportshape.GongMarshallField(stage, "Note"))
		}
	}
	if (noteportshape.Port == nil) != (noteportshapeOther.Port == nil) {
		diffs = append(diffs, noteportshape.GongMarshallField(stage, "Port"))
	} else if noteportshape.Port != nil && noteportshapeOther.Port != nil {
		if noteportshape.Port != noteportshapeOther.Port {
			diffs = append(diffs, noteportshape.GongMarshallField(stage, "Port"))
		}
	}
	if noteportshape.StartRatio != noteportshapeOther.StartRatio {
		diffs = append(diffs, noteportshape.GongMarshallField(stage, "StartRatio"))
	}
	if noteportshape.EndRatio != noteportshapeOther.EndRatio {
		diffs = append(diffs, noteportshape.GongMarshallField(stage, "EndRatio"))
	}
	if noteportshape.StartOrientation != noteportshapeOther.StartOrientation {
		diffs = append(diffs, noteportshape.GongMarshallField(stage, "StartOrientation"))
	}
	if noteportshape.EndOrientation != noteportshapeOther.EndOrientation {
		diffs = append(diffs, noteportshape.GongMarshallField(stage, "EndOrientation"))
	}
	if noteportshape.CornerOffsetRatio != noteportshapeOther.CornerOffsetRatio {
		diffs = append(diffs, noteportshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if noteportshape.IsHidden != noteportshapeOther.IsHidden {
		diffs = append(diffs, noteportshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (noteshape *NoteShape) GongDiff(stage *Stage, noteshapeOther *NoteShape) (diffs []string) {
	// insertion point for field diffs
	if noteshape.Name != noteshapeOther.Name {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Name"))
	}
	if (noteshape.Note == nil) != (noteshapeOther.Note == nil) {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
	} else if noteshape.Note != nil && noteshapeOther.Note != nil {
		if noteshape.Note != noteshapeOther.Note {
			diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
		}
	}
	if noteshape.X != noteshapeOther.X {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "X"))
	}
	if noteshape.Y != noteshapeOther.Y {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Y"))
	}
	if noteshape.Width != noteshapeOther.Width {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Width"))
	}
	if noteshape.Height != noteshapeOther.Height {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Height"))
	}
	if noteshape.IsHidden != noteshapeOther.IsHidden {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (part *Part) GongDiff(stage *Stage, partOther *Part) (diffs []string) {
	// insertion point for field diffs
	if part.Name != partOther.Name {
		diffs = append(diffs, part.GongMarshallField(stage, "Name"))
	}
	if part.Description != partOther.Description {
		diffs = append(diffs, part.GongMarshallField(stage, "Description"))
	}
	PortsDifferent := false
	if len(part.Ports) != len(partOther.Ports) {
		PortsDifferent = true
	} else {
		for i := range part.Ports {
			if (part.Ports[i] == nil) != (partOther.Ports[i] == nil) {
				PortsDifferent = true
				break
			} else if part.Ports[i] != nil && partOther.Ports[i] != nil {
				// this is a pointer comparaison
				if part.Ports[i] != partOther.Ports[i] {
					PortsDifferent = true
					break
				}
			}
		}
	}
	if PortsDifferent {
		ops := stage.Diff(
			part,
			"Ports",
			len(partOther.Ports),
			len(part.Ports),
			func(i, j int) bool {
				return partOther.Ports[i] == part.Ports[j]
			},
			func(j int) string {
				return part.Ports[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if (part.TypeOfPart == nil) != (partOther.TypeOfPart == nil) {
		diffs = append(diffs, part.GongMarshallField(stage, "TypeOfPart"))
	} else if part.TypeOfPart != nil && partOther.TypeOfPart != nil {
		if part.TypeOfPart != partOther.TypeOfPart {
			diffs = append(diffs, part.GongMarshallField(stage, "TypeOfPart"))
		}
	}
	if part.IsPartNameNotSystemName != partOther.IsPartNameNotSystemName {
		diffs = append(diffs, part.GongMarshallField(stage, "IsPartNameNotSystemName"))
	}
	if part.IsControlFlowsNodeExpanded != partOther.IsControlFlowsNodeExpanded {
		diffs = append(diffs, part.GongMarshallField(stage, "IsControlFlowsNodeExpanded"))
	}
	ControlFlowsDifferent := false
	if len(part.ControlFlows) != len(partOther.ControlFlows) {
		ControlFlowsDifferent = true
	} else {
		for i := range part.ControlFlows {
			if (part.ControlFlows[i] == nil) != (partOther.ControlFlows[i] == nil) {
				ControlFlowsDifferent = true
				break
			} else if part.ControlFlows[i] != nil && partOther.ControlFlows[i] != nil {
				// this is a pointer comparaison
				if part.ControlFlows[i] != partOther.ControlFlows[i] {
					ControlFlowsDifferent = true
					break
				}
			}
		}
	}
	if ControlFlowsDifferent {
		ops := stage.Diff(
			part,
			"ControlFlows",
			len(partOther.ControlFlows),
			len(part.ControlFlows),
			func(i, j int) bool {
				return partOther.ControlFlows[i] == part.ControlFlows[j]
			},
			func(j int) string {
				return part.ControlFlows[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	PortWhoseOutControlFlowsNodeIsExpandedDifferent := false
	if len(part.PortWhoseOutControlFlowsNodeIsExpanded) != len(partOther.PortWhoseOutControlFlowsNodeIsExpanded) {
		PortWhoseOutControlFlowsNodeIsExpandedDifferent = true
	} else {
		for i := range part.PortWhoseOutControlFlowsNodeIsExpanded {
			if (part.PortWhoseOutControlFlowsNodeIsExpanded[i] == nil) != (partOther.PortWhoseOutControlFlowsNodeIsExpanded[i] == nil) {
				PortWhoseOutControlFlowsNodeIsExpandedDifferent = true
				break
			} else if part.PortWhoseOutControlFlowsNodeIsExpanded[i] != nil && partOther.PortWhoseOutControlFlowsNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if part.PortWhoseOutControlFlowsNodeIsExpanded[i] != partOther.PortWhoseOutControlFlowsNodeIsExpanded[i] {
					PortWhoseOutControlFlowsNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if PortWhoseOutControlFlowsNodeIsExpandedDifferent {
		ops := stage.Diff(
			part,
			"PortWhoseOutControlFlowsNodeIsExpanded",
			len(partOther.PortWhoseOutControlFlowsNodeIsExpanded),
			len(part.PortWhoseOutControlFlowsNodeIsExpanded),
			func(i, j int) bool {
				return partOther.PortWhoseOutControlFlowsNodeIsExpanded[i] == part.PortWhoseOutControlFlowsNodeIsExpanded[j]
			},
			func(j int) string {
				return part.PortWhoseOutControlFlowsNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	PortWhoseInControlFlowsNodeIsExpandedDifferent := false
	if len(part.PortWhoseInControlFlowsNodeIsExpanded) != len(partOther.PortWhoseInControlFlowsNodeIsExpanded) {
		PortWhoseInControlFlowsNodeIsExpandedDifferent = true
	} else {
		for i := range part.PortWhoseInControlFlowsNodeIsExpanded {
			if (part.PortWhoseInControlFlowsNodeIsExpanded[i] == nil) != (partOther.PortWhoseInControlFlowsNodeIsExpanded[i] == nil) {
				PortWhoseInControlFlowsNodeIsExpandedDifferent = true
				break
			} else if part.PortWhoseInControlFlowsNodeIsExpanded[i] != nil && partOther.PortWhoseInControlFlowsNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if part.PortWhoseInControlFlowsNodeIsExpanded[i] != partOther.PortWhoseInControlFlowsNodeIsExpanded[i] {
					PortWhoseInControlFlowsNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if PortWhoseInControlFlowsNodeIsExpandedDifferent {
		ops := stage.Diff(
			part,
			"PortWhoseInControlFlowsNodeIsExpanded",
			len(partOther.PortWhoseInControlFlowsNodeIsExpanded),
			len(part.PortWhoseInControlFlowsNodeIsExpanded),
			func(i, j int) bool {
				return partOther.PortWhoseInControlFlowsNodeIsExpanded[i] == part.PortWhoseInControlFlowsNodeIsExpanded[j]
			},
			func(j int) string {
				return part.PortWhoseInControlFlowsNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if part.IsDataFlowsNodeExpanded != partOther.IsDataFlowsNodeExpanded {
		diffs = append(diffs, part.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
	}
	PortWhoseOutDataFlowsNodeIsExpandedDifferent := false
	if len(part.PortWhoseOutDataFlowsNodeIsExpanded) != len(partOther.PortWhoseOutDataFlowsNodeIsExpanded) {
		PortWhoseOutDataFlowsNodeIsExpandedDifferent = true
	} else {
		for i := range part.PortWhoseOutDataFlowsNodeIsExpanded {
			if (part.PortWhoseOutDataFlowsNodeIsExpanded[i] == nil) != (partOther.PortWhoseOutDataFlowsNodeIsExpanded[i] == nil) {
				PortWhoseOutDataFlowsNodeIsExpandedDifferent = true
				break
			} else if part.PortWhoseOutDataFlowsNodeIsExpanded[i] != nil && partOther.PortWhoseOutDataFlowsNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if part.PortWhoseOutDataFlowsNodeIsExpanded[i] != partOther.PortWhoseOutDataFlowsNodeIsExpanded[i] {
					PortWhoseOutDataFlowsNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if PortWhoseOutDataFlowsNodeIsExpandedDifferent {
		ops := stage.Diff(
			part,
			"PortWhoseOutDataFlowsNodeIsExpanded",
			len(partOther.PortWhoseOutDataFlowsNodeIsExpanded),
			len(part.PortWhoseOutDataFlowsNodeIsExpanded),
			func(i, j int) bool {
				return partOther.PortWhoseOutDataFlowsNodeIsExpanded[i] == part.PortWhoseOutDataFlowsNodeIsExpanded[j]
			},
			func(j int) string {
				return part.PortWhoseOutDataFlowsNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	PortWhoseInDataFlowsNodeIsExpandedDifferent := false
	if len(part.PortWhoseInDataFlowsNodeIsExpanded) != len(partOther.PortWhoseInDataFlowsNodeIsExpanded) {
		PortWhoseInDataFlowsNodeIsExpandedDifferent = true
	} else {
		for i := range part.PortWhoseInDataFlowsNodeIsExpanded {
			if (part.PortWhoseInDataFlowsNodeIsExpanded[i] == nil) != (partOther.PortWhoseInDataFlowsNodeIsExpanded[i] == nil) {
				PortWhoseInDataFlowsNodeIsExpandedDifferent = true
				break
			} else if part.PortWhoseInDataFlowsNodeIsExpanded[i] != nil && partOther.PortWhoseInDataFlowsNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if part.PortWhoseInDataFlowsNodeIsExpanded[i] != partOther.PortWhoseInDataFlowsNodeIsExpanded[i] {
					PortWhoseInDataFlowsNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if PortWhoseInDataFlowsNodeIsExpandedDifferent {
		ops := stage.Diff(
			part,
			"PortWhoseInDataFlowsNodeIsExpanded",
			len(partOther.PortWhoseInDataFlowsNodeIsExpanded),
			len(part.PortWhoseInDataFlowsNodeIsExpanded),
			func(i, j int) bool {
				return partOther.PortWhoseInDataFlowsNodeIsExpanded[i] == part.PortWhoseInDataFlowsNodeIsExpanded[j]
			},
			func(j int) string {
				return part.PortWhoseInDataFlowsNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	PartAnchoredPathDifferent := false
	if len(part.PartAnchoredPath) != len(partOther.PartAnchoredPath) {
		PartAnchoredPathDifferent = true
	} else {
		for i := range part.PartAnchoredPath {
			if (part.PartAnchoredPath[i] == nil) != (partOther.PartAnchoredPath[i] == nil) {
				PartAnchoredPathDifferent = true
				break
			} else if part.PartAnchoredPath[i] != nil && partOther.PartAnchoredPath[i] != nil {
				// this is a pointer comparaison
				if part.PartAnchoredPath[i] != partOther.PartAnchoredPath[i] {
					PartAnchoredPathDifferent = true
					break
				}
			}
		}
	}
	if PartAnchoredPathDifferent {
		ops := stage.Diff(
			part,
			"PartAnchoredPath",
			len(partOther.PartAnchoredPath),
			len(part.PartAnchoredPath),
			func(i, j int) bool {
				return partOther.PartAnchoredPath[i] == part.PartAnchoredPath[j]
			},
			func(j int) string {
				return part.PartAnchoredPath[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if part.ComputedPrefix != partOther.ComputedPrefix {
		diffs = append(diffs, part.GongMarshallField(stage, "ComputedPrefix"))
	}
	if part.IsExpanded != partOther.IsExpanded {
		diffs = append(diffs, part.GongMarshallField(stage, "IsExpanded"))
	}
	if part.IsPortsNodeExpanded != partOther.IsPortsNodeExpanded {
		diffs = append(diffs, part.GongMarshallField(stage, "IsPortsNodeExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partanchoredpath *PartAnchoredPath) GongDiff(stage *Stage, partanchoredpathOther *PartAnchoredPath) (diffs []string) {
	// insertion point for field diffs
	if partanchoredpath.Name != partanchoredpathOther.Name {
		diffs = append(diffs, partanchoredpath.GongMarshallField(stage, "Name"))
	}
	if partanchoredpath.Definition != partanchoredpathOther.Definition {
		diffs = append(diffs, partanchoredpath.GongMarshallField(stage, "Definition"))
	}
	if partanchoredpath.X_Offset != partanchoredpathOther.X_Offset {
		diffs = append(diffs, partanchoredpath.GongMarshallField(stage, "X_Offset"))
	}
	if partanchoredpath.Y_Offset != partanchoredpathOther.Y_Offset {
		diffs = append(diffs, partanchoredpath.GongMarshallField(stage, "Y_Offset"))
	}
	if partanchoredpath.RectAnchorType != partanchoredpathOther.RectAnchorType {
		diffs = append(diffs, partanchoredpath.GongMarshallField(stage, "RectAnchorType"))
	}
	if partanchoredpath.ScalePropotionnally != partanchoredpathOther.ScalePropotionnally {
		diffs = append(diffs, partanchoredpath.GongMarshallField(stage, "ScalePropotionnally"))
	}
	if partanchoredpath.AppliedScaling != partanchoredpathOther.AppliedScaling {
		diffs = append(diffs, partanchoredpath.GongMarshallField(stage, "AppliedScaling"))
	}
	if partanchoredpath.Color != partanchoredpathOther.Color {
		diffs = append(diffs, partanchoredpath.GongMarshallField(stage, "Color"))
	}
	if partanchoredpath.FillOpacity != partanchoredpathOther.FillOpacity {
		diffs = append(diffs, partanchoredpath.GongMarshallField(stage, "FillOpacity"))
	}
	if partanchoredpath.Stroke != partanchoredpathOther.Stroke {
		diffs = append(diffs, partanchoredpath.GongMarshallField(stage, "Stroke"))
	}
	if partanchoredpath.StrokeOpacity != partanchoredpathOther.StrokeOpacity {
		diffs = append(diffs, partanchoredpath.GongMarshallField(stage, "StrokeOpacity"))
	}
	if partanchoredpath.StrokeWidth != partanchoredpathOther.StrokeWidth {
		diffs = append(diffs, partanchoredpath.GongMarshallField(stage, "StrokeWidth"))
	}
	if partanchoredpath.StrokeDashArray != partanchoredpathOther.StrokeDashArray {
		diffs = append(diffs, partanchoredpath.GongMarshallField(stage, "StrokeDashArray"))
	}
	if partanchoredpath.StrokeDashArrayWhenSelected != partanchoredpathOther.StrokeDashArrayWhenSelected {
		diffs = append(diffs, partanchoredpath.GongMarshallField(stage, "StrokeDashArrayWhenSelected"))
	}
	if partanchoredpath.Transform != partanchoredpathOther.Transform {
		diffs = append(diffs, partanchoredpath.GongMarshallField(stage, "Transform"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (partshape *PartShape) GongDiff(stage *Stage, partshapeOther *PartShape) (diffs []string) {
	// insertion point for field diffs
	if partshape.Name != partshapeOther.Name {
		diffs = append(diffs, partshape.GongMarshallField(stage, "Name"))
	}
	if (partshape.Part == nil) != (partshapeOther.Part == nil) {
		diffs = append(diffs, partshape.GongMarshallField(stage, "Part"))
	} else if partshape.Part != nil && partshapeOther.Part != nil {
		if partshape.Part != partshapeOther.Part {
			diffs = append(diffs, partshape.GongMarshallField(stage, "Part"))
		}
	}
	if partshape.IsExpanded != partshapeOther.IsExpanded {
		diffs = append(diffs, partshape.GongMarshallField(stage, "IsExpanded"))
	}
	if partshape.X != partshapeOther.X {
		diffs = append(diffs, partshape.GongMarshallField(stage, "X"))
	}
	if partshape.Y != partshapeOther.Y {
		diffs = append(diffs, partshape.GongMarshallField(stage, "Y"))
	}
	if partshape.Width != partshapeOther.Width {
		diffs = append(diffs, partshape.GongMarshallField(stage, "Width"))
	}
	if partshape.Height != partshapeOther.Height {
		diffs = append(diffs, partshape.GongMarshallField(stage, "Height"))
	}
	if partshape.IsHidden != partshapeOther.IsHidden {
		diffs = append(diffs, partshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (port *Port) GongDiff(stage *Stage, portOther *Port) (diffs []string) {
	// insertion point for field diffs
	if port.Name != portOther.Name {
		diffs = append(diffs, port.GongMarshallField(stage, "Name"))
	}
	if port.Description != portOther.Description {
		diffs = append(diffs, port.GongMarshallField(stage, "Description"))
	}
	if port.ComputedPrefix != portOther.ComputedPrefix {
		diffs = append(diffs, port.GongMarshallField(stage, "ComputedPrefix"))
	}
	if port.IsExpanded != portOther.IsExpanded {
		diffs = append(diffs, port.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (portshape *PortShape) GongDiff(stage *Stage, portshapeOther *PortShape) (diffs []string) {
	// insertion point for field diffs
	if portshape.Name != portshapeOther.Name {
		diffs = append(diffs, portshape.GongMarshallField(stage, "Name"))
	}
	if (portshape.Port == nil) != (portshapeOther.Port == nil) {
		diffs = append(diffs, portshape.GongMarshallField(stage, "Port"))
	} else if portshape.Port != nil && portshapeOther.Port != nil {
		if portshape.Port != portshapeOther.Port {
			diffs = append(diffs, portshape.GongMarshallField(stage, "Port"))
		}
	}
	if portshape.IsExpanded != portshapeOther.IsExpanded {
		diffs = append(diffs, portshape.GongMarshallField(stage, "IsExpanded"))
	}
	if portshape.X != portshapeOther.X {
		diffs = append(diffs, portshape.GongMarshallField(stage, "X"))
	}
	if portshape.Y != portshapeOther.Y {
		diffs = append(diffs, portshape.GongMarshallField(stage, "Y"))
	}
	if portshape.Width != portshapeOther.Width {
		diffs = append(diffs, portshape.GongMarshallField(stage, "Width"))
	}
	if portshape.Height != portshapeOther.Height {
		diffs = append(diffs, portshape.GongMarshallField(stage, "Height"))
	}
	if portshape.IsHidden != portshapeOther.IsHidden {
		diffs = append(diffs, portshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (resource *Resource) GongDiff(stage *Stage, resourceOther *Resource) (diffs []string) {
	// insertion point for field diffs
	if resource.Name != resourceOther.Name {
		diffs = append(diffs, resource.GongMarshallField(stage, "Name"))
	}
	if resource.Acronym != resourceOther.Acronym {
		diffs = append(diffs, resource.GongMarshallField(stage, "Acronym"))
	}
	if resource.Description != resourceOther.Description {
		diffs = append(diffs, resource.GongMarshallField(stage, "Description"))
	}
	if resource.ComputedPrefix != resourceOther.ComputedPrefix {
		diffs = append(diffs, resource.GongMarshallField(stage, "ComputedPrefix"))
	}
	if resource.IsExpanded != resourceOther.IsExpanded {
		diffs = append(diffs, resource.GongMarshallField(stage, "IsExpanded"))
	}
	if resource.SVG_Path != resourceOther.SVG_Path {
		diffs = append(diffs, resource.GongMarshallField(stage, "SVG_Path"))
	}
	if resource.InverseAppliedScaling != resourceOther.InverseAppliedScaling {
		diffs = append(diffs, resource.GongMarshallField(stage, "InverseAppliedScaling"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (semantictag *SemanticTag) GongDiff(stage *Stage, semantictagOther *SemanticTag) (diffs []string) {
	// insertion point for field diffs
	if semantictag.Name != semantictagOther.Name {
		diffs = append(diffs, semantictag.GongMarshallField(stage, "Name"))
	}
	PartsDifferent := false
	if len(semantictag.Parts) != len(semantictagOther.Parts) {
		PartsDifferent = true
	} else {
		for i := range semantictag.Parts {
			if (semantictag.Parts[i] == nil) != (semantictagOther.Parts[i] == nil) {
				PartsDifferent = true
				break
			} else if semantictag.Parts[i] != nil && semantictagOther.Parts[i] != nil {
				// this is a pointer comparaison
				if semantictag.Parts[i] != semantictagOther.Parts[i] {
					PartsDifferent = true
					break
				}
			}
		}
	}
	if PartsDifferent {
		ops := stage.Diff(
			semantictag,
			"Parts",
			len(semantictagOther.Parts),
			len(semantictag.Parts),
			func(i, j int) bool {
				return semantictagOther.Parts[i] == semantictag.Parts[j]
			},
			func(j int) string {
				return semantictag.Parts[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (system *System) GongDiff(stage *Stage, systemOther *System) (diffs []string) {
	// insertion point for field diffs
	if system.Name != systemOther.Name {
		diffs = append(diffs, system.GongMarshallField(stage, "Name"))
	}
	if system.Description != systemOther.Description {
		diffs = append(diffs, system.GongMarshallField(stage, "Description"))
	}
	if system.ComputedPrefix != systemOther.ComputedPrefix {
		diffs = append(diffs, system.GongMarshallField(stage, "ComputedPrefix"))
	}
	if system.IsExpanded != systemOther.IsExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsExpanded"))
	}
	if system.SVG_Path != systemOther.SVG_Path {
		diffs = append(diffs, system.GongMarshallField(stage, "SVG_Path"))
	}
	if system.InverseAppliedScaling != systemOther.InverseAppliedScaling {
		diffs = append(diffs, system.GongMarshallField(stage, "InverseAppliedScaling"))
	}
	DiagramStructuresDifferent := false
	if len(system.DiagramStructures) != len(systemOther.DiagramStructures) {
		DiagramStructuresDifferent = true
	} else {
		for i := range system.DiagramStructures {
			if (system.DiagramStructures[i] == nil) != (systemOther.DiagramStructures[i] == nil) {
				DiagramStructuresDifferent = true
				break
			} else if system.DiagramStructures[i] != nil && systemOther.DiagramStructures[i] != nil {
				// this is a pointer comparaison
				if system.DiagramStructures[i] != systemOther.DiagramStructures[i] {
					DiagramStructuresDifferent = true
					break
				}
			}
		}
	}
	if DiagramStructuresDifferent {
		ops := stage.Diff(
			system,
			"DiagramStructures",
			len(systemOther.DiagramStructures),
			len(system.DiagramStructures),
			func(i, j int) bool {
				return systemOther.DiagramStructures[i] == system.DiagramStructures[j]
			},
			func(j int) string {
				return system.DiagramStructures[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DiagramStructureWhoseNodeIsExpandedDifferent := false
	if len(system.DiagramStructureWhoseNodeIsExpanded) != len(systemOther.DiagramStructureWhoseNodeIsExpanded) {
		DiagramStructureWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range system.DiagramStructureWhoseNodeIsExpanded {
			if (system.DiagramStructureWhoseNodeIsExpanded[i] == nil) != (systemOther.DiagramStructureWhoseNodeIsExpanded[i] == nil) {
				DiagramStructureWhoseNodeIsExpandedDifferent = true
				break
			} else if system.DiagramStructureWhoseNodeIsExpanded[i] != nil && systemOther.DiagramStructureWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if system.DiagramStructureWhoseNodeIsExpanded[i] != systemOther.DiagramStructureWhoseNodeIsExpanded[i] {
					DiagramStructureWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if DiagramStructureWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			system,
			"DiagramStructureWhoseNodeIsExpanded",
			len(systemOther.DiagramStructureWhoseNodeIsExpanded),
			len(system.DiagramStructureWhoseNodeIsExpanded),
			func(i, j int) bool {
				return systemOther.DiagramStructureWhoseNodeIsExpanded[i] == system.DiagramStructureWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return system.DiagramStructureWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if system.IsSubSystemNodeExpanded != systemOther.IsSubSystemNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsSubSystemNodeExpanded"))
	}
	SubSystemesDifferent := false
	if len(system.SubSystemes) != len(systemOther.SubSystemes) {
		SubSystemesDifferent = true
	} else {
		for i := range system.SubSystemes {
			if (system.SubSystemes[i] == nil) != (systemOther.SubSystemes[i] == nil) {
				SubSystemesDifferent = true
				break
			} else if system.SubSystemes[i] != nil && systemOther.SubSystemes[i] != nil {
				// this is a pointer comparaison
				if system.SubSystemes[i] != systemOther.SubSystemes[i] {
					SubSystemesDifferent = true
					break
				}
			}
		}
	}
	if SubSystemesDifferent {
		ops := stage.Diff(
			system,
			"SubSystemes",
			len(systemOther.SubSystemes),
			len(system.SubSystemes),
			func(i, j int) bool {
				return systemOther.SubSystemes[i] == system.SubSystemes[j]
			},
			func(j int) string {
				return system.SubSystemes[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	PartsDifferent := false
	if len(system.Parts) != len(systemOther.Parts) {
		PartsDifferent = true
	} else {
		for i := range system.Parts {
			if (system.Parts[i] == nil) != (systemOther.Parts[i] == nil) {
				PartsDifferent = true
				break
			} else if system.Parts[i] != nil && systemOther.Parts[i] != nil {
				// this is a pointer comparaison
				if system.Parts[i] != systemOther.Parts[i] {
					PartsDifferent = true
					break
				}
			}
		}
	}
	if PartsDifferent {
		ops := stage.Diff(
			system,
			"Parts",
			len(systemOther.Parts),
			len(system.Parts),
			func(i, j int) bool {
				return systemOther.Parts[i] == system.Parts[j]
			},
			func(j int) string {
				return system.Parts[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	PartWhoseNodeIsExpandedDifferent := false
	if len(system.PartWhoseNodeIsExpanded) != len(systemOther.PartWhoseNodeIsExpanded) {
		PartWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range system.PartWhoseNodeIsExpanded {
			if (system.PartWhoseNodeIsExpanded[i] == nil) != (systemOther.PartWhoseNodeIsExpanded[i] == nil) {
				PartWhoseNodeIsExpandedDifferent = true
				break
			} else if system.PartWhoseNodeIsExpanded[i] != nil && systemOther.PartWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if system.PartWhoseNodeIsExpanded[i] != systemOther.PartWhoseNodeIsExpanded[i] {
					PartWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if PartWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			system,
			"PartWhoseNodeIsExpanded",
			len(systemOther.PartWhoseNodeIsExpanded),
			len(system.PartWhoseNodeIsExpanded),
			func(i, j int) bool {
				return systemOther.PartWhoseNodeIsExpanded[i] == system.PartWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return system.PartWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DataFlowsDifferent := false
	if len(system.DataFlows) != len(systemOther.DataFlows) {
		DataFlowsDifferent = true
	} else {
		for i := range system.DataFlows {
			if (system.DataFlows[i] == nil) != (systemOther.DataFlows[i] == nil) {
				DataFlowsDifferent = true
				break
			} else if system.DataFlows[i] != nil && systemOther.DataFlows[i] != nil {
				// this is a pointer comparaison
				if system.DataFlows[i] != systemOther.DataFlows[i] {
					DataFlowsDifferent = true
					break
				}
			}
		}
	}
	if DataFlowsDifferent {
		ops := stage.Diff(
			system,
			"DataFlows",
			len(systemOther.DataFlows),
			len(system.DataFlows),
			func(i, j int) bool {
				return systemOther.DataFlows[i] == system.DataFlows[j]
			},
			func(j int) string {
				return system.DataFlows[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if system.IsDataFlowsNodeExpanded != systemOther.IsDataFlowsNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
	}
	ExternalPartsDifferent := false
	if len(system.ExternalParts) != len(systemOther.ExternalParts) {
		ExternalPartsDifferent = true
	} else {
		for i := range system.ExternalParts {
			if (system.ExternalParts[i] == nil) != (systemOther.ExternalParts[i] == nil) {
				ExternalPartsDifferent = true
				break
			} else if system.ExternalParts[i] != nil && systemOther.ExternalParts[i] != nil {
				// this is a pointer comparaison
				if system.ExternalParts[i] != systemOther.ExternalParts[i] {
					ExternalPartsDifferent = true
					break
				}
			}
		}
	}
	if ExternalPartsDifferent {
		ops := stage.Diff(
			system,
			"ExternalParts",
			len(systemOther.ExternalParts),
			len(system.ExternalParts),
			func(i, j int) bool {
				return systemOther.ExternalParts[i] == system.ExternalParts[j]
			},
			func(j int) string {
				return system.ExternalParts[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ExternalPartWhoseNodeIsExpandedDifferent := false
	if len(system.ExternalPartWhoseNodeIsExpanded) != len(systemOther.ExternalPartWhoseNodeIsExpanded) {
		ExternalPartWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range system.ExternalPartWhoseNodeIsExpanded {
			if (system.ExternalPartWhoseNodeIsExpanded[i] == nil) != (systemOther.ExternalPartWhoseNodeIsExpanded[i] == nil) {
				ExternalPartWhoseNodeIsExpandedDifferent = true
				break
			} else if system.ExternalPartWhoseNodeIsExpanded[i] != nil && systemOther.ExternalPartWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if system.ExternalPartWhoseNodeIsExpanded[i] != systemOther.ExternalPartWhoseNodeIsExpanded[i] {
					ExternalPartWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if ExternalPartWhoseNodeIsExpandedDifferent {
		ops := stage.Diff(
			system,
			"ExternalPartWhoseNodeIsExpanded",
			len(systemOther.ExternalPartWhoseNodeIsExpanded),
			len(system.ExternalPartWhoseNodeIsExpanded),
			func(i, j int) bool {
				return systemOther.ExternalPartWhoseNodeIsExpanded[i] == system.ExternalPartWhoseNodeIsExpanded[j]
			},
			func(j int) string {
				return system.ExternalPartWhoseNodeIsExpanded[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (systemshape *SystemShape) GongDiff(stage *Stage, systemshapeOther *SystemShape) (diffs []string) {
	// insertion point for field diffs
	if systemshape.Name != systemshapeOther.Name {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "Name"))
	}
	if (systemshape.System == nil) != (systemshapeOther.System == nil) {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "System"))
	} else if systemshape.System != nil && systemshapeOther.System != nil {
		if systemshape.System != systemshapeOther.System {
			diffs = append(diffs, systemshape.GongMarshallField(stage, "System"))
		}
	}
	if systemshape.IsExpanded != systemshapeOther.IsExpanded {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "IsExpanded"))
	}
	if systemshape.X != systemshapeOther.X {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "X"))
	}
	if systemshape.Y != systemshapeOther.Y {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "Y"))
	}
	if systemshape.Width != systemshapeOther.Width {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "Width"))
	}
	if systemshape.Height != systemshapeOther.Height {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "Height"))
	}
	if systemshape.IsHidden != systemshapeOther.IsHidden {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "IsHidden"))
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
