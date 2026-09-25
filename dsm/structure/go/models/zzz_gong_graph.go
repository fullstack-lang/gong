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
func (allocatedresourceshape *AllocatedResourceShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AllocatedResourceShapes[allocatedresourceshape]
	return ok
}

func (allocatedsystemshape *AllocatedSystemShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AllocatedSystemShapes[allocatedsystemshape]
	return ok
}

func (controlflow *ControlFlow) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ControlFlows[controlflow]
	return ok
}

func (controlflowshape *ControlFlowShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ControlFlowShapes[controlflowshape]
	return ok
}

func (data *Data) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Datas[data]
	return ok
}

func (dataflow *DataFlow) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DataFlows[dataflow]
	return ok
}

func (dataflowshape *DataFlowShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DataFlowShapes[dataflowshape]
	return ok
}

func (datashape *DataShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DataShapes[datashape]
	return ok
}

func (diagramlayerstate *DiagramLayerState) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DiagramLayerStates[diagramlayerstate]
	return ok
}

func (diagramstructure *DiagramStructure) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DiagramStructures[diagramstructure]
	return ok
}

func (externalpartshape *ExternalPartShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ExternalPartShapes[externalpartshape]
	return ok
}

func (layerdefinition *LayerDefinition) GongIsStaged(stage *Stage) bool {
	_, ok := stage.LayerDefinitions[layerdefinition]
	return ok
}

func (library *Library) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Librarys[library]
	return ok
}

func (note *Note) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Notes[note]
	return ok
}

func (notepartshape *NotePartShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NotePartShapes[notepartshape]
	return ok
}

func (noteportshape *NotePortShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NotePortShapes[noteportshape]
	return ok
}

func (noteshape *NoteShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.NoteShapes[noteshape]
	return ok
}

func (part *Part) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Parts[part]
	return ok
}

func (partanchoredpath *PartAnchoredPath) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartAnchoredPaths[partanchoredpath]
	return ok
}

func (partshape *PartShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PartShapes[partshape]
	return ok
}

func (port *Port) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Ports[port]
	return ok
}

func (portshape *PortShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PortShapes[portshape]
	return ok
}

func (resource *Resource) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Resources[resource]
	return ok
}

func (semantictag *SemanticTag) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SemanticTags[semantictag]
	return ok
}

func (system *System) GongIsStaged(stage *Stage) bool {
	_, ok := stage.Systems[system]
	return ok
}

func (systemshape *SystemShape) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SystemShapes[systemshape]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (allocatedresourceshape *AllocatedResourceShape) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(data) {
		return
	}

	data.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dataflow *DataFlow) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(partanchoredpath) {
		return
	}

	partanchoredpath.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partshape *PartShape) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(port) {
		return
	}

	port.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (portshape *PortShape) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(resource) {
		return
	}

	resource.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (semantictag *SemanticTag) GongStageBranch(stage *Stage) {

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
	var alreadyCopied bool
	allocatedresourceshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, allocatedresourceshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	allocatedsystemshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, allocatedsystemshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	controlflowTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, controlflowFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	controlflowshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, controlflowshapeFrom)
	if alreadyCopied {
		return
	}
	controlflowshapeFrom.GongCopyBasicFields(controlflowshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if controlflowshapeFrom.ControlFlow != nil {
		controlflowshapeTo.ControlFlow = GongCopyBranchControlFlow(mapOrigCopy, controlflowshapeFrom.ControlFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchData(mapOrigCopy map[any]any, dataFrom *Data) (dataTo *Data) {
	var alreadyCopied bool
	dataTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, dataFrom)
	if alreadyCopied {
		return
	}
	dataFrom.GongCopyBasicFields(dataTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDataFlow(mapOrigCopy map[any]any, dataflowFrom *DataFlow) (dataflowTo *DataFlow) {
	var alreadyCopied bool
	dataflowTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, dataflowFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	dataflowshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, dataflowshapeFrom)
	if alreadyCopied {
		return
	}
	dataflowshapeFrom.GongCopyBasicFields(dataflowshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if dataflowshapeFrom.DataFlow != nil {
		dataflowshapeTo.DataFlow = GongCopyBranchDataFlow(mapOrigCopy, dataflowshapeFrom.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDataShape(mapOrigCopy map[any]any, datashapeFrom *DataShape) (datashapeTo *DataShape) {
	var alreadyCopied bool
	datashapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, datashapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	diagramlayerstateTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, diagramlayerstateFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	diagramstructureTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, diagramstructureFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	externalpartshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, externalpartshapeFrom)
	if alreadyCopied {
		return
	}
	externalpartshapeFrom.GongCopyBasicFields(externalpartshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if externalpartshapeFrom.Part != nil {
		externalpartshapeTo.Part = GongCopyBranchPart(mapOrigCopy, externalpartshapeFrom.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchLayerDefinition(mapOrigCopy map[any]any, layerdefinitionFrom *LayerDefinition) (layerdefinitionTo *LayerDefinition) {
	var alreadyCopied bool
	layerdefinitionTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, layerdefinitionFrom)
	if alreadyCopied {
		return
	}
	layerdefinitionFrom.GongCopyBasicFields(layerdefinitionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _semantictag := range layerdefinitionFrom.Query {
		layerdefinitionTo.Query = append(layerdefinitionTo.Query, GongCopyBranchSemanticTag(mapOrigCopy, _semantictag))
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
	var alreadyCopied bool
	noteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	notepartshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, notepartshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	noteportshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteportshapeFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	noteshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, noteshapeFrom)
	if alreadyCopied {
		return
	}
	noteshapeFrom.GongCopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteshapeFrom.Note != nil {
		noteshapeTo.Note = GongCopyBranchNote(mapOrigCopy, noteshapeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPart(mapOrigCopy map[any]any, partFrom *Part) (partTo *Part) {
	var alreadyCopied bool
	partTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	partanchoredpathTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partanchoredpathFrom)
	if alreadyCopied {
		return
	}
	partanchoredpathFrom.GongCopyBasicFields(partanchoredpathTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPartShape(mapOrigCopy map[any]any, partshapeFrom *PartShape) (partshapeTo *PartShape) {
	var alreadyCopied bool
	partshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, partshapeFrom)
	if alreadyCopied {
		return
	}
	partshapeFrom.GongCopyBasicFields(partshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if partshapeFrom.Part != nil {
		partshapeTo.Part = GongCopyBranchPart(mapOrigCopy, partshapeFrom.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPort(mapOrigCopy map[any]any, portFrom *Port) (portTo *Port) {
	var alreadyCopied bool
	portTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, portFrom)
	if alreadyCopied {
		return
	}
	portFrom.GongCopyBasicFields(portTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPortShape(mapOrigCopy map[any]any, portshapeFrom *PortShape) (portshapeTo *PortShape) {
	var alreadyCopied bool
	portshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, portshapeFrom)
	if alreadyCopied {
		return
	}
	portshapeFrom.GongCopyBasicFields(portshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if portshapeFrom.Port != nil {
		portshapeTo.Port = GongCopyBranchPort(mapOrigCopy, portshapeFrom.Port)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchResource(mapOrigCopy map[any]any, resourceFrom *Resource) (resourceTo *Resource) {
	var alreadyCopied bool
	resourceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, resourceFrom)
	if alreadyCopied {
		return
	}
	resourceFrom.GongCopyBasicFields(resourceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSemanticTag(mapOrigCopy map[any]any, semantictagFrom *SemanticTag) (semantictagTo *SemanticTag) {
	var alreadyCopied bool
	semantictagTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, semantictagFrom)
	if alreadyCopied {
		return
	}
	semantictagFrom.GongCopyBasicFields(semantictagTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _part := range semantictagFrom.Parts {
		semantictagTo.Parts = append(semantictagTo.Parts, GongCopyBranchPart(mapOrigCopy, _part))
	}

	return
}

func GongCopyBranchSystem(mapOrigCopy map[any]any, systemFrom *System) (systemTo *System) {
	var alreadyCopied bool
	systemTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, systemFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	systemshapeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, systemshapeFrom)
	if alreadyCopied {
		return
	}
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

// insertion point for unstage branch per struct
func (allocatedresourceshape *AllocatedResourceShape) GongUnstageBranch(stage *Stage) {

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

	// check if instance is already staged
	if !stage.IsStaged(data) {
		return
	}

	data.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dataflow *DataFlow) GongUnstageBranch(stage *Stage) {

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

	// check if instance is already staged
	if !stage.IsStaged(partanchoredpath) {
		return
	}

	partanchoredpath.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (partshape *PartShape) GongUnstageBranch(stage *Stage) {

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

	// check if instance is already staged
	if !stage.IsStaged(port) {
		return
	}

	port.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (portshape *PortShape) GongUnstageBranch(stage *Stage) {

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

	// check if instance is already staged
	if !stage.IsStaged(resource) {
		return
	}

	resource.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (semantictag *SemanticTag) GongUnstageBranch(stage *Stage) {

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
	__gong__reconstructPointer(&reference.Part, stage.Parts_reference, instance.Part)
	__gong__reconstructPointer(&reference.Resource, stage.Resources_reference, instance.Resource)
	// insertion point for slice of pointers field
}

func (reference *AllocatedSystemShape) GongReconstructPointersFromReferences(stage *Stage, instance *AllocatedSystemShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Part, stage.Parts_reference, instance.Part)
	__gong__reconstructPointer(&reference.System, stage.Systems_reference, instance.System)
	// insertion point for slice of pointers field
}

func (reference *ControlFlow) GongReconstructPointersFromReferences(stage *Stage, instance *ControlFlow) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Start, stage.Ports_reference, instance.Start)
	__gong__reconstructPointer(&reference.End, stage.Ports_reference, instance.End)
	// insertion point for slice of pointers field
}

func (reference *ControlFlowShape) GongReconstructPointersFromReferences(stage *Stage, instance *ControlFlowShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.ControlFlow, stage.ControlFlows_reference, instance.ControlFlow)
	// insertion point for slice of pointers field
}

func (reference *Data) GongReconstructPointersFromReferences(stage *Stage, instance *Data) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DataFlow) GongReconstructPointersFromReferences(stage *Stage, instance *DataFlow) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.StartPort, stage.Ports_reference, instance.StartPort)
	__gong__reconstructPointer(&reference.EndPort, stage.Ports_reference, instance.EndPort)
	__gong__reconstructPointer(&reference.StartExternalPart, stage.Parts_reference, instance.StartExternalPart)
	__gong__reconstructPointer(&reference.EndExternalPart, stage.Parts_reference, instance.EndExternalPart)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Datas, stage.Datas_reference, instance.Datas)
}

func (reference *DataFlowShape) GongReconstructPointersFromReferences(stage *Stage, instance *DataFlowShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.DataFlow, stage.DataFlows_reference, instance.DataFlow)
	// insertion point for slice of pointers field
}

func (reference *DataShape) GongReconstructPointersFromReferences(stage *Stage, instance *DataShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Data, stage.Datas_reference, instance.Data)
	__gong__reconstructPointer(&reference.DataFlow, stage.DataFlows_reference, instance.DataFlow)
	// insertion point for slice of pointers field
}

func (reference *DiagramLayerState) GongReconstructPointersFromReferences(stage *Stage, instance *DiagramLayerState) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.DiagramStructure, stage.DiagramStructures_reference, instance.DiagramStructure)
	__gong__reconstructPointer(&reference.LayerDefinition, stage.LayerDefinitions_reference, instance.LayerDefinition)
	// insertion point for slice of pointers field
}

func (reference *DiagramStructure) GongReconstructPointersFromReferences(stage *Stage, instance *DiagramStructure) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.System_Shapes, stage.SystemShapes_reference, instance.System_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SystemsWhoseNodeIsExpanded, stage.Systems_reference, instance.SystemsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Part_Shapes, stage.PartShapes_reference, instance.Part_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PartWhoseNodeIsExpanded, stage.Parts_reference, instance.PartWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ExternalPart_Shapes, stage.ExternalPartShapes_reference, instance.ExternalPart_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ExternalPartWhoseNodeIsExpanded, stage.Parts_reference, instance.ExternalPartWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, stage.Parts_reference, instance.ExternalPartsWhoseOutDataFlowsNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ExternalPartsWhoseInDataFlowsNodeIsExpanded, stage.Parts_reference, instance.ExternalPartsWhoseInDataFlowsNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PortsWhoseNodeIsExpanded, stage.Ports_reference, instance.PortsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Port_Shapes, stage.PortShapes_reference, instance.Port_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlFlowsWhoseNodeIsExpanded, stage.ControlFlows_reference, instance.ControlFlowsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlFlow_Shapes, stage.ControlFlowShapes_reference, instance.ControlFlow_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DataFlowsWhoseNodeIsExpanded, stage.DataFlows_reference, instance.DataFlowsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DataFlow_Shapes, stage.DataFlowShapes_reference, instance.DataFlow_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DatasWhoseNodeIsExpanded, stage.Datas_reference, instance.DatasWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Data_Shapes, stage.DataShapes_reference, instance.Data_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DataFlowsWhoseDataNodeIsExpanded, stage.DataFlows_reference, instance.DataFlowsWhoseDataNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AllocatedResourcesWhoseNodeIsExpanded, stage.Resources_reference, instance.AllocatedResourcesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AllocatedResourceShapes, stage.AllocatedResourceShapes_reference, instance.AllocatedResourceShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AllocatedSystemesWhoseNodeIsExpanded, stage.Systems_reference, instance.AllocatedSystemesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.AllocatedSystemShapes, stage.AllocatedSystemShapes_reference, instance.AllocatedSystemShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Note_Shapes, stage.NoteShapes_reference, instance.Note_Shapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NotesWhoseNodeIsExpanded, stage.Notes_reference, instance.NotesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NotePortShapes, stage.NotePortShapes_reference, instance.NotePortShapes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NotePartShapes, stage.NotePartShapes_reference, instance.NotePartShapes)
}

func (reference *ExternalPartShape) GongReconstructPointersFromReferences(stage *Stage, instance *ExternalPartShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Part, stage.Parts_reference, instance.Part)
	// insertion point for slice of pointers field
}

func (reference *LayerDefinition) GongReconstructPointersFromReferences(stage *Stage, instance *LayerDefinition) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Query, stage.SemanticTags_reference, instance.Query)
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubLibraries, stage.Librarys_reference, instance.SubLibraries)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_reference, instance.SubLibrariesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootSystemes, stage.Systems_reference, instance.RootSystemes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SystemsWhoseNodeIsExpanded, stage.Systems_reference, instance.SystemsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootDataFlows, stage.DataFlows_reference, instance.RootDataFlows)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DataFlowsWhoseNodeIsExpanded, stage.DataFlows_reference, instance.DataFlowsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootDatas, stage.Datas_reference, instance.RootDatas)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DatasWhoseNodeIsExpanded, stage.Datas_reference, instance.DatasWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootResources, stage.Resources_reference, instance.RootResources)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ResourcesWhoseNodeIsExpanded, stage.Resources_reference, instance.ResourcesWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PartsWhoseNodeIsExpanded, stage.Parts_reference, instance.PartsWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.RootNotes, stage.Notes_reference, instance.RootNotes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.NotesWhoseNodeIsExpanded, stage.Notes_reference, instance.NotesWhoseNodeIsExpanded)
}

func (reference *Note) GongReconstructPointersFromReferences(stage *Stage, instance *Note) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Parts, stage.Parts_reference, instance.Parts)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Ports, stage.Ports_reference, instance.Ports)
}

func (reference *NotePartShape) GongReconstructPointersFromReferences(stage *Stage, instance *NotePartShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	__gong__reconstructPointer(&reference.Part, stage.Parts_reference, instance.Part)
	// insertion point for slice of pointers field
}

func (reference *NotePortShape) GongReconstructPointersFromReferences(stage *Stage, instance *NotePortShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	__gong__reconstructPointer(&reference.Port, stage.Ports_reference, instance.Port)
	// insertion point for slice of pointers field
}

func (reference *NoteShape) GongReconstructPointersFromReferences(stage *Stage, instance *NoteShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Note, stage.Notes_reference, instance.Note)
	// insertion point for slice of pointers field
}

func (reference *Part) GongReconstructPointersFromReferences(stage *Stage, instance *Part) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.TypeOfPart, stage.Systems_reference, instance.TypeOfPart)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Ports, stage.Ports_reference, instance.Ports)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ControlFlows, stage.ControlFlows_reference, instance.ControlFlows)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PortWhoseOutControlFlowsNodeIsExpanded, stage.Ports_reference, instance.PortWhoseOutControlFlowsNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PortWhoseInControlFlowsNodeIsExpanded, stage.Ports_reference, instance.PortWhoseInControlFlowsNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PortWhoseOutDataFlowsNodeIsExpanded, stage.Ports_reference, instance.PortWhoseOutDataFlowsNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PortWhoseInDataFlowsNodeIsExpanded, stage.Ports_reference, instance.PortWhoseInDataFlowsNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PartAnchoredPath, stage.PartAnchoredPaths_reference, instance.PartAnchoredPath)
}

func (reference *PartAnchoredPath) GongReconstructPointersFromReferences(stage *Stage, instance *PartAnchoredPath) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PartShape) GongReconstructPointersFromReferences(stage *Stage, instance *PartShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Part, stage.Parts_reference, instance.Part)
	// insertion point for slice of pointers field
}

func (reference *Port) GongReconstructPointersFromReferences(stage *Stage, instance *Port) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PortShape) GongReconstructPointersFromReferences(stage *Stage, instance *PortShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Port, stage.Ports_reference, instance.Port)
	// insertion point for slice of pointers field
}

func (reference *Resource) GongReconstructPointersFromReferences(stage *Stage, instance *Resource) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *SemanticTag) GongReconstructPointersFromReferences(stage *Stage, instance *SemanticTag) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Parts, stage.Parts_reference, instance.Parts)
}

func (reference *System) GongReconstructPointersFromReferences(stage *Stage, instance *System) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.DiagramStructures, stage.DiagramStructures_reference, instance.DiagramStructures)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DiagramStructureWhoseNodeIsExpanded, stage.DiagramStructures_reference, instance.DiagramStructureWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SubSystemes, stage.Systems_reference, instance.SubSystemes)
	__gong__reconstructSliceOfPointersFromReferences(&reference.Parts, stage.Parts_reference, instance.Parts)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PartWhoseNodeIsExpanded, stage.Parts_reference, instance.PartWhoseNodeIsExpanded)
	__gong__reconstructSliceOfPointersFromReferences(&reference.DataFlows, stage.DataFlows_reference, instance.DataFlows)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ExternalParts, stage.Parts_reference, instance.ExternalParts)
	__gong__reconstructSliceOfPointersFromReferences(&reference.ExternalPartWhoseNodeIsExpanded, stage.Parts_reference, instance.ExternalPartWhoseNodeIsExpanded)
}

func (reference *SystemShape) GongReconstructPointersFromReferences(stage *Stage, instance *SystemShape) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.System, stage.Systems_reference, instance.System)
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AllocatedResourceShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Part, stage.Parts_instance)
	__gong__reconstructPointerFromInstance(&reference.Resource, stage.Resources_instance)
	// insertion point for slice of pointers fields
}

func (reference *AllocatedSystemShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Part, stage.Parts_instance)
	__gong__reconstructPointerFromInstance(&reference.System, stage.Systems_instance)
	// insertion point for slice of pointers fields
}

func (reference *ControlFlow) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Start, stage.Ports_instance)
	__gong__reconstructPointerFromInstance(&reference.End, stage.Ports_instance)
	// insertion point for slice of pointers fields
}

func (reference *ControlFlowShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.ControlFlow, stage.ControlFlows_instance)
	// insertion point for slice of pointers fields
}

func (reference *Data) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DataFlow) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.StartPort, stage.Ports_instance)
	__gong__reconstructPointerFromInstance(&reference.EndPort, stage.Ports_instance)
	__gong__reconstructPointerFromInstance(&reference.StartExternalPart, stage.Parts_instance)
	__gong__reconstructPointerFromInstance(&reference.EndExternalPart, stage.Parts_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Datas, stage.Datas_instance)
}

func (reference *DataFlowShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.DataFlow, stage.DataFlows_instance)
	// insertion point for slice of pointers fields
}

func (reference *DataShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Data, stage.Datas_instance)
	__gong__reconstructPointerFromInstance(&reference.DataFlow, stage.DataFlows_instance)
	// insertion point for slice of pointers fields
}

func (reference *DiagramLayerState) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.DiagramStructure, stage.DiagramStructures_instance)
	__gong__reconstructPointerFromInstance(&reference.LayerDefinition, stage.LayerDefinitions_instance)
	// insertion point for slice of pointers fields
}

func (reference *DiagramStructure) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.System_Shapes, stage.SystemShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SystemsWhoseNodeIsExpanded, stage.Systems_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Part_Shapes, stage.PartShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PartWhoseNodeIsExpanded, stage.Parts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ExternalPart_Shapes, stage.ExternalPartShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ExternalPartWhoseNodeIsExpanded, stage.Parts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, stage.Parts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ExternalPartsWhoseInDataFlowsNodeIsExpanded, stage.Parts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PortsWhoseNodeIsExpanded, stage.Ports_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Port_Shapes, stage.PortShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlFlowsWhoseNodeIsExpanded, stage.ControlFlows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlFlow_Shapes, stage.ControlFlowShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DataFlowsWhoseNodeIsExpanded, stage.DataFlows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DataFlow_Shapes, stage.DataFlowShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DatasWhoseNodeIsExpanded, stage.Datas_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Data_Shapes, stage.DataShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DataFlowsWhoseDataNodeIsExpanded, stage.DataFlows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AllocatedResourcesWhoseNodeIsExpanded, stage.Resources_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AllocatedResourceShapes, stage.AllocatedResourceShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AllocatedSystemesWhoseNodeIsExpanded, stage.Systems_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.AllocatedSystemShapes, stage.AllocatedSystemShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Note_Shapes, stage.NoteShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NotesWhoseNodeIsExpanded, stage.Notes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NotePortShapes, stage.NotePortShapes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NotePartShapes, stage.NotePartShapes_instance)
}

func (reference *ExternalPartShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Part, stage.Parts_instance)
	// insertion point for slice of pointers fields
}

func (reference *LayerDefinition) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Query, stage.SemanticTags_instance)
}

func (reference *Library) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubLibraries, stage.Librarys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubLibrariesWhoseNodeIsExpanded, stage.Librarys_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootSystemes, stage.Systems_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SystemsWhoseNodeIsExpanded, stage.Systems_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootDataFlows, stage.DataFlows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DataFlowsWhoseNodeIsExpanded, stage.DataFlows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootDatas, stage.Datas_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DatasWhoseNodeIsExpanded, stage.Datas_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootResources, stage.Resources_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ResourcesWhoseNodeIsExpanded, stage.Resources_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PartsWhoseNodeIsExpanded, stage.Parts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.RootNotes, stage.Notes_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.NotesWhoseNodeIsExpanded, stage.Notes_instance)
}

func (reference *Note) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Parts, stage.Parts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Ports, stage.Ports_instance)
}

func (reference *NotePartShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	__gong__reconstructPointerFromInstance(&reference.Part, stage.Parts_instance)
	// insertion point for slice of pointers fields
}

func (reference *NotePortShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	__gong__reconstructPointerFromInstance(&reference.Port, stage.Ports_instance)
	// insertion point for slice of pointers fields
}

func (reference *NoteShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Note, stage.Notes_instance)
	// insertion point for slice of pointers fields
}

func (reference *Part) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.TypeOfPart, stage.Systems_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Ports, stage.Ports_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ControlFlows, stage.ControlFlows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PortWhoseOutControlFlowsNodeIsExpanded, stage.Ports_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PortWhoseInControlFlowsNodeIsExpanded, stage.Ports_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PortWhoseOutDataFlowsNodeIsExpanded, stage.Ports_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PortWhoseInDataFlowsNodeIsExpanded, stage.Ports_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PartAnchoredPath, stage.PartAnchoredPaths_instance)
}

func (reference *PartAnchoredPath) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PartShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Part, stage.Parts_instance)
	// insertion point for slice of pointers fields
}

func (reference *Port) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PortShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Port, stage.Ports_instance)
	// insertion point for slice of pointers fields
}

func (reference *Resource) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *SemanticTag) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Parts, stage.Parts_instance)
}

func (reference *System) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.DiagramStructures, stage.DiagramStructures_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DiagramStructureWhoseNodeIsExpanded, stage.DiagramStructures_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SubSystemes, stage.Systems_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.Parts, stage.Parts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PartWhoseNodeIsExpanded, stage.Parts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.DataFlows, stage.DataFlows_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ExternalParts, stage.Parts_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.ExternalPartWhoseNodeIsExpanded, stage.Parts_instance)
}

func (reference *SystemShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.System, stage.Systems_instance)
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
	if allocatedresourceshape.Part != allocatedresourceshapeOther.Part {
		diffs = append(diffs, allocatedresourceshape.GongMarshallField(stage, "Part"))
	}
	if allocatedresourceshape.Resource != allocatedresourceshapeOther.Resource {
		diffs = append(diffs, allocatedresourceshape.GongMarshallField(stage, "Resource"))
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
	if allocatedsystemshape.Part != allocatedsystemshapeOther.Part {
		diffs = append(diffs, allocatedsystemshape.GongMarshallField(stage, "Part"))
	}
	if allocatedsystemshape.System != allocatedsystemshapeOther.System {
		diffs = append(diffs, allocatedsystemshape.GongMarshallField(stage, "System"))
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
	if controlflow.Start != controlflowOther.Start {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "Start"))
	}
	if controlflow.End != controlflowOther.End {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "End"))
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
	if controlflowshape.ControlFlow != controlflowshapeOther.ControlFlow {
		diffs = append(diffs, controlflowshape.GongMarshallField(stage, "ControlFlow"))
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
	if dataflow.StartPort != dataflowOther.StartPort {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "StartPort"))
	}
	if dataflow.EndPort != dataflowOther.EndPort {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "EndPort"))
	}
	if dataflow.StartExternalPart != dataflowOther.StartExternalPart {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "StartExternalPart"))
	}
	if dataflow.EndExternalPart != dataflowOther.EndExternalPart {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "EndExternalPart"))
	}
	if ops := __gong__diffSliceOfPointers(stage, dataflow, "Datas", dataflowOther.Datas, dataflow.Datas); ops != "" {
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
	if dataflowshape.DataFlow != dataflowshapeOther.DataFlow {
		diffs = append(diffs, dataflowshape.GongMarshallField(stage, "DataFlow"))
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
	if datashape.Data != datashapeOther.Data {
		diffs = append(diffs, datashape.GongMarshallField(stage, "Data"))
	}
	if datashape.DataFlow != datashapeOther.DataFlow {
		diffs = append(diffs, datashape.GongMarshallField(stage, "DataFlow"))
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
	if diagramlayerstate.DiagramStructure != diagramlayerstateOther.DiagramStructure {
		diffs = append(diffs, diagramlayerstate.GongMarshallField(stage, "DiagramStructure"))
	}
	if diagramlayerstate.LayerDefinition != diagramlayerstateOther.LayerDefinition {
		diffs = append(diffs, diagramlayerstate.GongMarshallField(stage, "LayerDefinition"))
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
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "System_Shapes", diagramstructureOther.System_Shapes, diagramstructure.System_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagramstructure.IsSystemsNodeExpanded != diagramstructureOther.IsSystemsNodeExpanded {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "IsSystemsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "SystemsWhoseNodeIsExpanded", diagramstructureOther.SystemsWhoseNodeIsExpanded, diagramstructure.SystemsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "Part_Shapes", diagramstructureOther.Part_Shapes, diagramstructure.Part_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagramstructure.IsPartsNodeExpanded != diagramstructureOther.IsPartsNodeExpanded {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "IsPartsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "PartWhoseNodeIsExpanded", diagramstructureOther.PartWhoseNodeIsExpanded, diagramstructure.PartWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "ExternalPart_Shapes", diagramstructureOther.ExternalPart_Shapes, diagramstructure.ExternalPart_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagramstructure.IsExternalPartsNodeExpanded != diagramstructureOther.IsExternalPartsNodeExpanded {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "IsExternalPartsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "ExternalPartWhoseNodeIsExpanded", diagramstructureOther.ExternalPartWhoseNodeIsExpanded, diagramstructure.ExternalPartWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "ExternalPartsWhoseOutDataFlowsNodeIsExpanded", diagramstructureOther.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "ExternalPartsWhoseInDataFlowsNodeIsExpanded", diagramstructureOther.ExternalPartsWhoseInDataFlowsNodeIsExpanded, diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "PortsWhoseNodeIsExpanded", diagramstructureOther.PortsWhoseNodeIsExpanded, diagramstructure.PortsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "Port_Shapes", diagramstructureOther.Port_Shapes, diagramstructure.Port_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "ControlFlowsWhoseNodeIsExpanded", diagramstructureOther.ControlFlowsWhoseNodeIsExpanded, diagramstructure.ControlFlowsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "ControlFlow_Shapes", diagramstructureOther.ControlFlow_Shapes, diagramstructure.ControlFlow_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "DataFlowsWhoseNodeIsExpanded", diagramstructureOther.DataFlowsWhoseNodeIsExpanded, diagramstructure.DataFlowsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "DataFlow_Shapes", diagramstructureOther.DataFlow_Shapes, diagramstructure.DataFlow_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "DatasWhoseNodeIsExpanded", diagramstructureOther.DatasWhoseNodeIsExpanded, diagramstructure.DatasWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "Data_Shapes", diagramstructureOther.Data_Shapes, diagramstructure.Data_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "DataFlowsWhoseDataNodeIsExpanded", diagramstructureOther.DataFlowsWhoseDataNodeIsExpanded, diagramstructure.DataFlowsWhoseDataNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "AllocatedResourcesWhoseNodeIsExpanded", diagramstructureOther.AllocatedResourcesWhoseNodeIsExpanded, diagramstructure.AllocatedResourcesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "AllocatedResourceShapes", diagramstructureOther.AllocatedResourceShapes, diagramstructure.AllocatedResourceShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "AllocatedSystemesWhoseNodeIsExpanded", diagramstructureOther.AllocatedSystemesWhoseNodeIsExpanded, diagramstructure.AllocatedSystemesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "AllocatedSystemShapes", diagramstructureOther.AllocatedSystemShapes, diagramstructure.AllocatedSystemShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "Note_Shapes", diagramstructureOther.Note_Shapes, diagramstructure.Note_Shapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "NotesWhoseNodeIsExpanded", diagramstructureOther.NotesWhoseNodeIsExpanded, diagramstructure.NotesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if diagramstructure.IsNotesNodeExpanded != diagramstructureOther.IsNotesNodeExpanded {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "NotePortShapes", diagramstructureOther.NotePortShapes, diagramstructure.NotePortShapes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, diagramstructure, "NotePartShapes", diagramstructureOther.NotePartShapes, diagramstructure.NotePartShapes); ops != "" {
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
	if externalpartshape.Part != externalpartshapeOther.Part {
		diffs = append(diffs, externalpartshape.GongMarshallField(stage, "Part"))
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
	if ops := __gong__diffSliceOfPointers(stage, layerdefinition, "Query", layerdefinitionOther.Query, layerdefinition.Query); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, library, "RootSystemes", libraryOther.RootSystemes, library.RootSystemes); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsSystemesNodeExpanded != libraryOther.IsSystemesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsSystemesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "SystemsWhoseNodeIsExpanded", libraryOther.SystemsWhoseNodeIsExpanded, library.SystemsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootDataFlows", libraryOther.RootDataFlows, library.RootDataFlows); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsDataFlowsNodeExpanded != libraryOther.IsDataFlowsNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "DataFlowsWhoseNodeIsExpanded", libraryOther.DataFlowsWhoseNodeIsExpanded, library.DataFlowsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootDatas", libraryOther.RootDatas, library.RootDatas); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsDatasNodeExpanded != libraryOther.IsDatasNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsDatasNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "DatasWhoseNodeIsExpanded", libraryOther.DatasWhoseNodeIsExpanded, library.DatasWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootResources", libraryOther.RootResources, library.RootResources); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsResourcesNodeExpanded != libraryOther.IsResourcesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsResourcesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "ResourcesWhoseNodeIsExpanded", libraryOther.ResourcesWhoseNodeIsExpanded, library.ResourcesWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "PartsWhoseNodeIsExpanded", libraryOther.PartsWhoseNodeIsExpanded, library.PartsWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "RootNotes", libraryOther.RootNotes, library.RootNotes); ops != "" {
		diffs = append(diffs, ops)
	}
	if library.IsNotesNodeExpanded != libraryOther.IsNotesNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsNotesNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, library, "NotesWhoseNodeIsExpanded", libraryOther.NotesWhoseNodeIsExpanded, library.NotesWhoseNodeIsExpanded); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, note, "Parts", noteOther.Parts, note.Parts); ops != "" {
		diffs = append(diffs, ops)
	}
	if note.IsPortsNodeExpanded != noteOther.IsPortsNodeExpanded {
		diffs = append(diffs, note.GongMarshallField(stage, "IsPortsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, note, "Ports", noteOther.Ports, note.Ports); ops != "" {
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
	if notepartshape.Note != notepartshapeOther.Note {
		diffs = append(diffs, notepartshape.GongMarshallField(stage, "Note"))
	}
	if notepartshape.Part != notepartshapeOther.Part {
		diffs = append(diffs, notepartshape.GongMarshallField(stage, "Part"))
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
	if noteportshape.Note != noteportshapeOther.Note {
		diffs = append(diffs, noteportshape.GongMarshallField(stage, "Note"))
	}
	if noteportshape.Port != noteportshapeOther.Port {
		diffs = append(diffs, noteportshape.GongMarshallField(stage, "Port"))
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
	if noteshape.Note != noteshapeOther.Note {
		diffs = append(diffs, noteshape.GongMarshallField(stage, "Note"))
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
	if ops := __gong__diffSliceOfPointers(stage, part, "Ports", partOther.Ports, part.Ports); ops != "" {
		diffs = append(diffs, ops)
	}
	if part.TypeOfPart != partOther.TypeOfPart {
		diffs = append(diffs, part.GongMarshallField(stage, "TypeOfPart"))
	}
	if part.IsPartNameNotSystemName != partOther.IsPartNameNotSystemName {
		diffs = append(diffs, part.GongMarshallField(stage, "IsPartNameNotSystemName"))
	}
	if part.IsControlFlowsNodeExpanded != partOther.IsControlFlowsNodeExpanded {
		diffs = append(diffs, part.GongMarshallField(stage, "IsControlFlowsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, part, "ControlFlows", partOther.ControlFlows, part.ControlFlows); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, part, "PortWhoseOutControlFlowsNodeIsExpanded", partOther.PortWhoseOutControlFlowsNodeIsExpanded, part.PortWhoseOutControlFlowsNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, part, "PortWhoseInControlFlowsNodeIsExpanded", partOther.PortWhoseInControlFlowsNodeIsExpanded, part.PortWhoseInControlFlowsNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if part.IsDataFlowsNodeExpanded != partOther.IsDataFlowsNodeExpanded {
		diffs = append(diffs, part.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, part, "PortWhoseOutDataFlowsNodeIsExpanded", partOther.PortWhoseOutDataFlowsNodeIsExpanded, part.PortWhoseOutDataFlowsNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, part, "PortWhoseInDataFlowsNodeIsExpanded", partOther.PortWhoseInDataFlowsNodeIsExpanded, part.PortWhoseInDataFlowsNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, part, "PartAnchoredPath", partOther.PartAnchoredPath, part.PartAnchoredPath); ops != "" {
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
	if partshape.Part != partshapeOther.Part {
		diffs = append(diffs, partshape.GongMarshallField(stage, "Part"))
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
	if portshape.Port != portshapeOther.Port {
		diffs = append(diffs, portshape.GongMarshallField(stage, "Port"))
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
	if ops := __gong__diffSliceOfPointers(stage, semantictag, "Parts", semantictagOther.Parts, semantictag.Parts); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, system, "DiagramStructures", systemOther.DiagramStructures, system.DiagramStructures); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, system, "DiagramStructureWhoseNodeIsExpanded", systemOther.DiagramStructureWhoseNodeIsExpanded, system.DiagramStructureWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if system.IsSubSystemNodeExpanded != systemOther.IsSubSystemNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsSubSystemNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, system, "SubSystemes", systemOther.SubSystemes, system.SubSystemes); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, system, "Parts", systemOther.Parts, system.Parts); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, system, "PartWhoseNodeIsExpanded", systemOther.PartWhoseNodeIsExpanded, system.PartWhoseNodeIsExpanded); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, system, "DataFlows", systemOther.DataFlows, system.DataFlows); ops != "" {
		diffs = append(diffs, ops)
	}
	if system.IsDataFlowsNodeExpanded != systemOther.IsDataFlowsNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsDataFlowsNodeExpanded"))
	}
	if ops := __gong__diffSliceOfPointers(stage, system, "ExternalParts", systemOther.ExternalParts, system.ExternalParts); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, system, "ExternalPartWhoseNodeIsExpanded", systemOther.ExternalPartWhoseNodeIsExpanded, system.ExternalPartWhoseNodeIsExpanded); ops != "" {
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
	if systemshape.System != systemshapeOther.System {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "System"))
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
