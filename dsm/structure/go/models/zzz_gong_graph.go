// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *AllocatedResourceShape:
		ok = stage.IsStagedAllocatedResourceShape(target)

	case *AllocatedSystemShape:
		ok = stage.IsStagedAllocatedSystemShape(target)

	case *ControlFlow:
		ok = stage.IsStagedControlFlow(target)

	case *ControlFlowShape:
		ok = stage.IsStagedControlFlowShape(target)

	case *Data:
		ok = stage.IsStagedData(target)

	case *DataFlow:
		ok = stage.IsStagedDataFlow(target)

	case *DataFlowShape:
		ok = stage.IsStagedDataFlowShape(target)

	case *DataShape:
		ok = stage.IsStagedDataShape(target)

	case *DiagramStructure:
		ok = stage.IsStagedDiagramStructure(target)

	case *ExternalPartShape:
		ok = stage.IsStagedExternalPartShape(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Note:
		ok = stage.IsStagedNote(target)

	case *NotePortShape:
		ok = stage.IsStagedNotePortShape(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *Part:
		ok = stage.IsStagedPart(target)

	case *PartShape:
		ok = stage.IsStagedPartShape(target)

	case *Port:
		ok = stage.IsStagedPort(target)

	case *PortShape:
		ok = stage.IsStagedPortShape(target)

	case *Resource:
		ok = stage.IsStagedResource(target)

	case *System:
		ok = stage.IsStagedSystem(target)

	case *SystemShape:
		ok = stage.IsStagedSystemShape(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *AllocatedResourceShape:
		ok = stage.IsStagedAllocatedResourceShape(target)

	case *AllocatedSystemShape:
		ok = stage.IsStagedAllocatedSystemShape(target)

	case *ControlFlow:
		ok = stage.IsStagedControlFlow(target)

	case *ControlFlowShape:
		ok = stage.IsStagedControlFlowShape(target)

	case *Data:
		ok = stage.IsStagedData(target)

	case *DataFlow:
		ok = stage.IsStagedDataFlow(target)

	case *DataFlowShape:
		ok = stage.IsStagedDataFlowShape(target)

	case *DataShape:
		ok = stage.IsStagedDataShape(target)

	case *DiagramStructure:
		ok = stage.IsStagedDiagramStructure(target)

	case *ExternalPartShape:
		ok = stage.IsStagedExternalPartShape(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Note:
		ok = stage.IsStagedNote(target)

	case *NotePortShape:
		ok = stage.IsStagedNotePortShape(target)

	case *NoteShape:
		ok = stage.IsStagedNoteShape(target)

	case *Part:
		ok = stage.IsStagedPart(target)

	case *PartShape:
		ok = stage.IsStagedPartShape(target)

	case *Port:
		ok = stage.IsStagedPort(target)

	case *PortShape:
		ok = stage.IsStagedPortShape(target)

	case *Resource:
		ok = stage.IsStagedResource(target)

	case *System:
		ok = stage.IsStagedSystem(target)

	case *SystemShape:
		ok = stage.IsStagedSystemShape(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedAllocatedResourceShape(allocatedresourceshape *AllocatedResourceShape) (ok bool) {

	_, ok = stage.AllocatedResourceShapes[allocatedresourceshape]

	return
}

func (stage *Stage) IsStagedAllocatedSystemShape(allocatedsystemshape *AllocatedSystemShape) (ok bool) {

	_, ok = stage.AllocatedSystemShapes[allocatedsystemshape]

	return
}

func (stage *Stage) IsStagedControlFlow(controlflow *ControlFlow) (ok bool) {

	_, ok = stage.ControlFlows[controlflow]

	return
}

func (stage *Stage) IsStagedControlFlowShape(controlflowshape *ControlFlowShape) (ok bool) {

	_, ok = stage.ControlFlowShapes[controlflowshape]

	return
}

func (stage *Stage) IsStagedData(data *Data) (ok bool) {

	_, ok = stage.Datas[data]

	return
}

func (stage *Stage) IsStagedDataFlow(dataflow *DataFlow) (ok bool) {

	_, ok = stage.DataFlows[dataflow]

	return
}

func (stage *Stage) IsStagedDataFlowShape(dataflowshape *DataFlowShape) (ok bool) {

	_, ok = stage.DataFlowShapes[dataflowshape]

	return
}

func (stage *Stage) IsStagedDataShape(datashape *DataShape) (ok bool) {

	_, ok = stage.DataShapes[datashape]

	return
}

func (stage *Stage) IsStagedDiagramStructure(diagramstructure *DiagramStructure) (ok bool) {

	_, ok = stage.DiagramStructures[diagramstructure]

	return
}

func (stage *Stage) IsStagedExternalPartShape(externalpartshape *ExternalPartShape) (ok bool) {

	_, ok = stage.ExternalPartShapes[externalpartshape]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedNote(note *Note) (ok bool) {

	_, ok = stage.Notes[note]

	return
}

func (stage *Stage) IsStagedNotePortShape(noteportshape *NotePortShape) (ok bool) {

	_, ok = stage.NotePortShapes[noteportshape]

	return
}

func (stage *Stage) IsStagedNoteShape(noteshape *NoteShape) (ok bool) {

	_, ok = stage.NoteShapes[noteshape]

	return
}

func (stage *Stage) IsStagedPart(part *Part) (ok bool) {

	_, ok = stage.Parts[part]

	return
}

func (stage *Stage) IsStagedPartShape(partshape *PartShape) (ok bool) {

	_, ok = stage.PartShapes[partshape]

	return
}

func (stage *Stage) IsStagedPort(port *Port) (ok bool) {

	_, ok = stage.Ports[port]

	return
}

func (stage *Stage) IsStagedPortShape(portshape *PortShape) (ok bool) {

	_, ok = stage.PortShapes[portshape]

	return
}

func (stage *Stage) IsStagedResource(resource *Resource) (ok bool) {

	_, ok = stage.Resources[resource]

	return
}

func (stage *Stage) IsStagedSystem(system *System) (ok bool) {

	_, ok = stage.Systems[system]

	return
}

func (stage *Stage) IsStagedSystemShape(systemshape *SystemShape) (ok bool) {

	_, ok = stage.SystemShapes[systemshape]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *AllocatedResourceShape:
		stage.StageBranchAllocatedResourceShape(target)

	case *AllocatedSystemShape:
		stage.StageBranchAllocatedSystemShape(target)

	case *ControlFlow:
		stage.StageBranchControlFlow(target)

	case *ControlFlowShape:
		stage.StageBranchControlFlowShape(target)

	case *Data:
		stage.StageBranchData(target)

	case *DataFlow:
		stage.StageBranchDataFlow(target)

	case *DataFlowShape:
		stage.StageBranchDataFlowShape(target)

	case *DataShape:
		stage.StageBranchDataShape(target)

	case *DiagramStructure:
		stage.StageBranchDiagramStructure(target)

	case *ExternalPartShape:
		stage.StageBranchExternalPartShape(target)

	case *Library:
		stage.StageBranchLibrary(target)

	case *Note:
		stage.StageBranchNote(target)

	case *NotePortShape:
		stage.StageBranchNotePortShape(target)

	case *NoteShape:
		stage.StageBranchNoteShape(target)

	case *Part:
		stage.StageBranchPart(target)

	case *PartShape:
		stage.StageBranchPartShape(target)

	case *Port:
		stage.StageBranchPort(target)

	case *PortShape:
		stage.StageBranchPortShape(target)

	case *Resource:
		stage.StageBranchResource(target)

	case *System:
		stage.StageBranchSystem(target)

	case *SystemShape:
		stage.StageBranchSystemShape(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchAllocatedResourceShape(allocatedresourceshape *AllocatedResourceShape) {

	// check if instance is already staged
	if IsStaged(stage, allocatedresourceshape) {
		return
	}

	allocatedresourceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if allocatedresourceshape.Part != nil {
		StageBranch(stage, allocatedresourceshape.Part)
	}
	if allocatedresourceshape.Resource != nil {
		StageBranch(stage, allocatedresourceshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchAllocatedSystemShape(allocatedsystemshape *AllocatedSystemShape) {

	// check if instance is already staged
	if IsStaged(stage, allocatedsystemshape) {
		return
	}

	allocatedsystemshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if allocatedsystemshape.Part != nil {
		StageBranch(stage, allocatedsystemshape.Part)
	}
	if allocatedsystemshape.System != nil {
		StageBranch(stage, allocatedsystemshape.System)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchControlFlow(controlflow *ControlFlow) {

	// check if instance is already staged
	if IsStaged(stage, controlflow) {
		return
	}

	controlflow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlflow.Start != nil {
		StageBranch(stage, controlflow.Start)
	}
	if controlflow.End != nil {
		StageBranch(stage, controlflow.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchControlFlowShape(controlflowshape *ControlFlowShape) {

	// check if instance is already staged
	if IsStaged(stage, controlflowshape) {
		return
	}

	controlflowshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlflowshape.ControlFlow != nil {
		StageBranch(stage, controlflowshape.ControlFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchData(data *Data) {

	// check if instance is already staged
	if IsStaged(stage, data) {
		return
	}

	data.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDataFlow(dataflow *DataFlow) {

	// check if instance is already staged
	if IsStaged(stage, dataflow) {
		return
	}

	dataflow.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dataflow.StartPort != nil {
		StageBranch(stage, dataflow.StartPort)
	}
	if dataflow.EndPort != nil {
		StageBranch(stage, dataflow.EndPort)
	}
	if dataflow.StartExternalPart != nil {
		StageBranch(stage, dataflow.StartExternalPart)
	}
	if dataflow.EndExternalPart != nil {
		StageBranch(stage, dataflow.EndExternalPart)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _data := range dataflow.Datas {
		StageBranch(stage, _data)
	}

}

func (stage *Stage) StageBranchDataFlowShape(dataflowshape *DataFlowShape) {

	// check if instance is already staged
	if IsStaged(stage, dataflowshape) {
		return
	}

	dataflowshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dataflowshape.DataFlow != nil {
		StageBranch(stage, dataflowshape.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDataShape(datashape *DataShape) {

	// check if instance is already staged
	if IsStaged(stage, datashape) {
		return
	}

	datashape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datashape.Data != nil {
		StageBranch(stage, datashape.Data)
	}
	if datashape.DataFlow != nil {
		StageBranch(stage, datashape.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDiagramStructure(diagramstructure *DiagramStructure) {

	// check if instance is already staged
	if IsStaged(stage, diagramstructure) {
		return
	}

	diagramstructure.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _systemshape := range diagramstructure.System_Shapes {
		StageBranch(stage, _systemshape)
	}
	for _, _system := range diagramstructure.SystemsWhoseNodeIsExpanded {
		StageBranch(stage, _system)
	}
	for _, _partshape := range diagramstructure.Part_Shapes {
		StageBranch(stage, _partshape)
	}
	for _, _part := range diagramstructure.PartWhoseNodeIsExpanded {
		StageBranch(stage, _part)
	}
	for _, _externalpartshape := range diagramstructure.ExternalPart_Shapes {
		StageBranch(stage, _externalpartshape)
	}
	for _, _part := range diagramstructure.ExternalPartWhoseNodeIsExpanded {
		StageBranch(stage, _part)
	}
	for _, _part := range diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
		StageBranch(stage, _part)
	}
	for _, _part := range diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
		StageBranch(stage, _part)
	}
	for _, _port := range diagramstructure.PortsWhoseNodeIsExpanded {
		StageBranch(stage, _port)
	}
	for _, _portshape := range diagramstructure.Port_Shapes {
		StageBranch(stage, _portshape)
	}
	for _, _controlflow := range diagramstructure.ControlFlowsWhoseNodeIsExpanded {
		StageBranch(stage, _controlflow)
	}
	for _, _controlflowshape := range diagramstructure.ControlFlow_Shapes {
		StageBranch(stage, _controlflowshape)
	}
	for _, _dataflow := range diagramstructure.DataFlowsWhoseNodeIsExpanded {
		StageBranch(stage, _dataflow)
	}
	for _, _dataflowshape := range diagramstructure.DataFlow_Shapes {
		StageBranch(stage, _dataflowshape)
	}
	for _, _data := range diagramstructure.DatasWhoseNodeIsExpanded {
		StageBranch(stage, _data)
	}
	for _, _datashape := range diagramstructure.Data_Shapes {
		StageBranch(stage, _datashape)
	}
	for _, _dataflow := range diagramstructure.DataFlowsWhoseDataNodeIsExpanded {
		StageBranch(stage, _dataflow)
	}
	for _, _resource := range diagramstructure.AllocatedResourcesWhoseNodeIsExpanded {
		StageBranch(stage, _resource)
	}
	for _, _allocatedresourceshape := range diagramstructure.AllocatedResourceShapes {
		StageBranch(stage, _allocatedresourceshape)
	}
	for _, _system := range diagramstructure.AllocatedSystemesWhoseNodeIsExpanded {
		StageBranch(stage, _system)
	}
	for _, _allocatedsystemshape := range diagramstructure.AllocatedSystemShapes {
		StageBranch(stage, _allocatedsystemshape)
	}
	for _, _noteshape := range diagramstructure.Note_Shapes {
		StageBranch(stage, _noteshape)
	}
	for _, _note := range diagramstructure.NotesWhoseNodeIsExpanded {
		StageBranch(stage, _note)
	}
	for _, _noteportshape := range diagramstructure.NotePortShapes {
		StageBranch(stage, _noteportshape)
	}

}

func (stage *Stage) StageBranchExternalPartShape(externalpartshape *ExternalPartShape) {

	// check if instance is already staged
	if IsStaged(stage, externalpartshape) {
		return
	}

	externalpartshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if externalpartshape.Part != nil {
		StageBranch(stage, externalpartshape.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchLibrary(library *Library) {

	// check if instance is already staged
	if IsStaged(stage, library) {
		return
	}

	library.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		StageBranch(stage, _library)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		StageBranch(stage, _library)
	}
	for _, _system := range library.RootSystemes {
		StageBranch(stage, _system)
	}
	for _, _system := range library.SystemsWhoseNodeIsExpanded {
		StageBranch(stage, _system)
	}
	for _, _dataflow := range library.RootDataFlows {
		StageBranch(stage, _dataflow)
	}
	for _, _dataflow := range library.DataFlowsWhoseNodeIsExpanded {
		StageBranch(stage, _dataflow)
	}
	for _, _data := range library.RootDatas {
		StageBranch(stage, _data)
	}
	for _, _data := range library.DatasWhoseNodeIsExpanded {
		StageBranch(stage, _data)
	}
	for _, _resource := range library.RootResources {
		StageBranch(stage, _resource)
	}
	for _, _resource := range library.ResourcesWhoseNodeIsExpanded {
		StageBranch(stage, _resource)
	}
	for _, _part := range library.PartsWhoseNodeIsExpanded {
		StageBranch(stage, _part)
	}
	for _, _note := range library.RootNotes {
		StageBranch(stage, _note)
	}
	for _, _note := range library.NotesWhoseNodeIsExpanded {
		StageBranch(stage, _note)
	}

}

func (stage *Stage) StageBranchNote(note *Note) {

	// check if instance is already staged
	if IsStaged(stage, note) {
		return
	}

	note.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _port := range note.Ports {
		StageBranch(stage, _port)
	}

}

func (stage *Stage) StageBranchNotePortShape(noteportshape *NotePortShape) {

	// check if instance is already staged
	if IsStaged(stage, noteportshape) {
		return
	}

	noteportshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteportshape.Note != nil {
		StageBranch(stage, noteportshape.Note)
	}
	if noteportshape.Port != nil {
		StageBranch(stage, noteportshape.Port)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if IsStaged(stage, noteshape) {
		return
	}

	noteshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		StageBranch(stage, noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPart(part *Part) {

	// check if instance is already staged
	if IsStaged(stage, part) {
		return
	}

	part.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _resource := range part.Resources {
		StageBranch(stage, _resource)
	}
	for _, _system := range part.Systemes {
		StageBranch(stage, _system)
	}
	for _, _port := range part.Ports {
		StageBranch(stage, _port)
	}
	for _, _controlflow := range part.ControlFlows {
		StageBranch(stage, _controlflow)
	}
	for _, _port := range part.PortWhoseOutControlFlowsNodeIsExpanded {
		StageBranch(stage, _port)
	}
	for _, _port := range part.PortWhoseInControlFlowsNodeIsExpanded {
		StageBranch(stage, _port)
	}
	for _, _port := range part.PortWhoseOutDataFlowsNodeIsExpanded {
		StageBranch(stage, _port)
	}
	for _, _port := range part.PortWhoseInDataFlowsNodeIsExpanded {
		StageBranch(stage, _port)
	}

}

func (stage *Stage) StageBranchPartShape(partshape *PartShape) {

	// check if instance is already staged
	if IsStaged(stage, partshape) {
		return
	}

	partshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if partshape.Part != nil {
		StageBranch(stage, partshape.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPort(port *Port) {

	// check if instance is already staged
	if IsStaged(stage, port) {
		return
	}

	port.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if port.Type != nil {
		StageBranch(stage, port.Type)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPortShape(portshape *PortShape) {

	// check if instance is already staged
	if IsStaged(stage, portshape) {
		return
	}

	portshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if portshape.Port != nil {
		StageBranch(stage, portshape.Port)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchResource(resource *Resource) {

	// check if instance is already staged
	if IsStaged(stage, resource) {
		return
	}

	resource.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSystem(system *System) {

	// check if instance is already staged
	if IsStaged(stage, system) {
		return
	}

	system.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagramstructure := range system.DiagramStructures {
		StageBranch(stage, _diagramstructure)
	}
	for _, _diagramstructure := range system.DiagramStructureWhoseNodeIsExpanded {
		StageBranch(stage, _diagramstructure)
	}
	for _, _system := range system.SubSystemes {
		StageBranch(stage, _system)
	}
	for _, _part := range system.Parts {
		StageBranch(stage, _part)
	}
	for _, _part := range system.PartWhoseNodeIsExpanded {
		StageBranch(stage, _part)
	}
	for _, _dataflow := range system.DataFlows {
		StageBranch(stage, _dataflow)
	}
	for _, _part := range system.ExternalParts {
		StageBranch(stage, _part)
	}
	for _, _part := range system.ExternalPartWhoseNodeIsExpanded {
		StageBranch(stage, _part)
	}

}

func (stage *Stage) StageBranchSystemShape(systemshape *SystemShape) {

	// check if instance is already staged
	if IsStaged(stage, systemshape) {
		return
	}

	systemshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if systemshape.System != nil {
		StageBranch(stage, systemshape.System)
	}

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
	case *AllocatedResourceShape:
		toT := CopyBranchAllocatedResourceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AllocatedSystemShape:
		toT := CopyBranchAllocatedSystemShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ControlFlow:
		toT := CopyBranchControlFlow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ControlFlowShape:
		toT := CopyBranchControlFlowShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Data:
		toT := CopyBranchData(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DataFlow:
		toT := CopyBranchDataFlow(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DataFlowShape:
		toT := CopyBranchDataFlowShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DataShape:
		toT := CopyBranchDataShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DiagramStructure:
		toT := CopyBranchDiagramStructure(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ExternalPartShape:
		toT := CopyBranchExternalPartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := CopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Note:
		toT := CopyBranchNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NotePortShape:
		toT := CopyBranchNotePortShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NoteShape:
		toT := CopyBranchNoteShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part:
		toT := CopyBranchPart(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartShape:
		toT := CopyBranchPartShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Port:
		toT := CopyBranchPort(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PortShape:
		toT := CopyBranchPortShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Resource:
		toT := CopyBranchResource(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *System:
		toT := CopyBranchSystem(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SystemShape:
		toT := CopyBranchSystemShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchAllocatedResourceShape(mapOrigCopy map[any]any, allocatedresourceshapeFrom *AllocatedResourceShape) (allocatedresourceshapeTo *AllocatedResourceShape) {

	// allocatedresourceshapeFrom has already been copied
	if _allocatedresourceshapeTo, ok := mapOrigCopy[allocatedresourceshapeFrom]; ok {
		allocatedresourceshapeTo = _allocatedresourceshapeTo.(*AllocatedResourceShape)
		return
	}

	allocatedresourceshapeTo = new(AllocatedResourceShape)
	mapOrigCopy[allocatedresourceshapeFrom] = allocatedresourceshapeTo
	allocatedresourceshapeFrom.CopyBasicFields(allocatedresourceshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if allocatedresourceshapeFrom.Part != nil {
		allocatedresourceshapeTo.Part = CopyBranchPart(mapOrigCopy, allocatedresourceshapeFrom.Part)
	}
	if allocatedresourceshapeFrom.Resource != nil {
		allocatedresourceshapeTo.Resource = CopyBranchResource(mapOrigCopy, allocatedresourceshapeFrom.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchAllocatedSystemShape(mapOrigCopy map[any]any, allocatedsystemshapeFrom *AllocatedSystemShape) (allocatedsystemshapeTo *AllocatedSystemShape) {

	// allocatedsystemshapeFrom has already been copied
	if _allocatedsystemshapeTo, ok := mapOrigCopy[allocatedsystemshapeFrom]; ok {
		allocatedsystemshapeTo = _allocatedsystemshapeTo.(*AllocatedSystemShape)
		return
	}

	allocatedsystemshapeTo = new(AllocatedSystemShape)
	mapOrigCopy[allocatedsystemshapeFrom] = allocatedsystemshapeTo
	allocatedsystemshapeFrom.CopyBasicFields(allocatedsystemshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if allocatedsystemshapeFrom.Part != nil {
		allocatedsystemshapeTo.Part = CopyBranchPart(mapOrigCopy, allocatedsystemshapeFrom.Part)
	}
	if allocatedsystemshapeFrom.System != nil {
		allocatedsystemshapeTo.System = CopyBranchSystem(mapOrigCopy, allocatedsystemshapeFrom.System)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchControlFlow(mapOrigCopy map[any]any, controlflowFrom *ControlFlow) (controlflowTo *ControlFlow) {

	// controlflowFrom has already been copied
	if _controlflowTo, ok := mapOrigCopy[controlflowFrom]; ok {
		controlflowTo = _controlflowTo.(*ControlFlow)
		return
	}

	controlflowTo = new(ControlFlow)
	mapOrigCopy[controlflowFrom] = controlflowTo
	controlflowFrom.CopyBasicFields(controlflowTo)

	//insertion point for the staging of instances referenced by pointers
	if controlflowFrom.Start != nil {
		controlflowTo.Start = CopyBranchPort(mapOrigCopy, controlflowFrom.Start)
	}
	if controlflowFrom.End != nil {
		controlflowTo.End = CopyBranchPort(mapOrigCopy, controlflowFrom.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchControlFlowShape(mapOrigCopy map[any]any, controlflowshapeFrom *ControlFlowShape) (controlflowshapeTo *ControlFlowShape) {

	// controlflowshapeFrom has already been copied
	if _controlflowshapeTo, ok := mapOrigCopy[controlflowshapeFrom]; ok {
		controlflowshapeTo = _controlflowshapeTo.(*ControlFlowShape)
		return
	}

	controlflowshapeTo = new(ControlFlowShape)
	mapOrigCopy[controlflowshapeFrom] = controlflowshapeTo
	controlflowshapeFrom.CopyBasicFields(controlflowshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if controlflowshapeFrom.ControlFlow != nil {
		controlflowshapeTo.ControlFlow = CopyBranchControlFlow(mapOrigCopy, controlflowshapeFrom.ControlFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchData(mapOrigCopy map[any]any, dataFrom *Data) (dataTo *Data) {

	// dataFrom has already been copied
	if _dataTo, ok := mapOrigCopy[dataFrom]; ok {
		dataTo = _dataTo.(*Data)
		return
	}

	dataTo = new(Data)
	mapOrigCopy[dataFrom] = dataTo
	dataFrom.CopyBasicFields(dataTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDataFlow(mapOrigCopy map[any]any, dataflowFrom *DataFlow) (dataflowTo *DataFlow) {

	// dataflowFrom has already been copied
	if _dataflowTo, ok := mapOrigCopy[dataflowFrom]; ok {
		dataflowTo = _dataflowTo.(*DataFlow)
		return
	}

	dataflowTo = new(DataFlow)
	mapOrigCopy[dataflowFrom] = dataflowTo
	dataflowFrom.CopyBasicFields(dataflowTo)

	//insertion point for the staging of instances referenced by pointers
	if dataflowFrom.StartPort != nil {
		dataflowTo.StartPort = CopyBranchPort(mapOrigCopy, dataflowFrom.StartPort)
	}
	if dataflowFrom.EndPort != nil {
		dataflowTo.EndPort = CopyBranchPort(mapOrigCopy, dataflowFrom.EndPort)
	}
	if dataflowFrom.StartExternalPart != nil {
		dataflowTo.StartExternalPart = CopyBranchPart(mapOrigCopy, dataflowFrom.StartExternalPart)
	}
	if dataflowFrom.EndExternalPart != nil {
		dataflowTo.EndExternalPart = CopyBranchPart(mapOrigCopy, dataflowFrom.EndExternalPart)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _data := range dataflowFrom.Datas {
		dataflowTo.Datas = append(dataflowTo.Datas, CopyBranchData(mapOrigCopy, _data))
	}

	return
}

func CopyBranchDataFlowShape(mapOrigCopy map[any]any, dataflowshapeFrom *DataFlowShape) (dataflowshapeTo *DataFlowShape) {

	// dataflowshapeFrom has already been copied
	if _dataflowshapeTo, ok := mapOrigCopy[dataflowshapeFrom]; ok {
		dataflowshapeTo = _dataflowshapeTo.(*DataFlowShape)
		return
	}

	dataflowshapeTo = new(DataFlowShape)
	mapOrigCopy[dataflowshapeFrom] = dataflowshapeTo
	dataflowshapeFrom.CopyBasicFields(dataflowshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if dataflowshapeFrom.DataFlow != nil {
		dataflowshapeTo.DataFlow = CopyBranchDataFlow(mapOrigCopy, dataflowshapeFrom.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDataShape(mapOrigCopy map[any]any, datashapeFrom *DataShape) (datashapeTo *DataShape) {

	// datashapeFrom has already been copied
	if _datashapeTo, ok := mapOrigCopy[datashapeFrom]; ok {
		datashapeTo = _datashapeTo.(*DataShape)
		return
	}

	datashapeTo = new(DataShape)
	mapOrigCopy[datashapeFrom] = datashapeTo
	datashapeFrom.CopyBasicFields(datashapeTo)

	//insertion point for the staging of instances referenced by pointers
	if datashapeFrom.Data != nil {
		datashapeTo.Data = CopyBranchData(mapOrigCopy, datashapeFrom.Data)
	}
	if datashapeFrom.DataFlow != nil {
		datashapeTo.DataFlow = CopyBranchDataFlow(mapOrigCopy, datashapeFrom.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDiagramStructure(mapOrigCopy map[any]any, diagramstructureFrom *DiagramStructure) (diagramstructureTo *DiagramStructure) {

	// diagramstructureFrom has already been copied
	if _diagramstructureTo, ok := mapOrigCopy[diagramstructureFrom]; ok {
		diagramstructureTo = _diagramstructureTo.(*DiagramStructure)
		return
	}

	diagramstructureTo = new(DiagramStructure)
	mapOrigCopy[diagramstructureFrom] = diagramstructureTo
	diagramstructureFrom.CopyBasicFields(diagramstructureTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _systemshape := range diagramstructureFrom.System_Shapes {
		diagramstructureTo.System_Shapes = append(diagramstructureTo.System_Shapes, CopyBranchSystemShape(mapOrigCopy, _systemshape))
	}
	for _, _system := range diagramstructureFrom.SystemsWhoseNodeIsExpanded {
		diagramstructureTo.SystemsWhoseNodeIsExpanded = append(diagramstructureTo.SystemsWhoseNodeIsExpanded, CopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _partshape := range diagramstructureFrom.Part_Shapes {
		diagramstructureTo.Part_Shapes = append(diagramstructureTo.Part_Shapes, CopyBranchPartShape(mapOrigCopy, _partshape))
	}
	for _, _part := range diagramstructureFrom.PartWhoseNodeIsExpanded {
		diagramstructureTo.PartWhoseNodeIsExpanded = append(diagramstructureTo.PartWhoseNodeIsExpanded, CopyBranchPart(mapOrigCopy, _part))
	}
	for _, _externalpartshape := range diagramstructureFrom.ExternalPart_Shapes {
		diagramstructureTo.ExternalPart_Shapes = append(diagramstructureTo.ExternalPart_Shapes, CopyBranchExternalPartShape(mapOrigCopy, _externalpartshape))
	}
	for _, _part := range diagramstructureFrom.ExternalPartWhoseNodeIsExpanded {
		diagramstructureTo.ExternalPartWhoseNodeIsExpanded = append(diagramstructureTo.ExternalPartWhoseNodeIsExpanded, CopyBranchPart(mapOrigCopy, _part))
	}
	for _, _part := range diagramstructureFrom.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
		diagramstructureTo.ExternalPartsWhoseOutDataFlowsNodeIsExpanded = append(diagramstructureTo.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, CopyBranchPart(mapOrigCopy, _part))
	}
	for _, _part := range diagramstructureFrom.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
		diagramstructureTo.ExternalPartsWhoseInDataFlowsNodeIsExpanded = append(diagramstructureTo.ExternalPartsWhoseInDataFlowsNodeIsExpanded, CopyBranchPart(mapOrigCopy, _part))
	}
	for _, _port := range diagramstructureFrom.PortsWhoseNodeIsExpanded {
		diagramstructureTo.PortsWhoseNodeIsExpanded = append(diagramstructureTo.PortsWhoseNodeIsExpanded, CopyBranchPort(mapOrigCopy, _port))
	}
	for _, _portshape := range diagramstructureFrom.Port_Shapes {
		diagramstructureTo.Port_Shapes = append(diagramstructureTo.Port_Shapes, CopyBranchPortShape(mapOrigCopy, _portshape))
	}
	for _, _controlflow := range diagramstructureFrom.ControlFlowsWhoseNodeIsExpanded {
		diagramstructureTo.ControlFlowsWhoseNodeIsExpanded = append(diagramstructureTo.ControlFlowsWhoseNodeIsExpanded, CopyBranchControlFlow(mapOrigCopy, _controlflow))
	}
	for _, _controlflowshape := range diagramstructureFrom.ControlFlow_Shapes {
		diagramstructureTo.ControlFlow_Shapes = append(diagramstructureTo.ControlFlow_Shapes, CopyBranchControlFlowShape(mapOrigCopy, _controlflowshape))
	}
	for _, _dataflow := range diagramstructureFrom.DataFlowsWhoseNodeIsExpanded {
		diagramstructureTo.DataFlowsWhoseNodeIsExpanded = append(diagramstructureTo.DataFlowsWhoseNodeIsExpanded, CopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _dataflowshape := range diagramstructureFrom.DataFlow_Shapes {
		diagramstructureTo.DataFlow_Shapes = append(diagramstructureTo.DataFlow_Shapes, CopyBranchDataFlowShape(mapOrigCopy, _dataflowshape))
	}
	for _, _data := range diagramstructureFrom.DatasWhoseNodeIsExpanded {
		diagramstructureTo.DatasWhoseNodeIsExpanded = append(diagramstructureTo.DatasWhoseNodeIsExpanded, CopyBranchData(mapOrigCopy, _data))
	}
	for _, _datashape := range diagramstructureFrom.Data_Shapes {
		diagramstructureTo.Data_Shapes = append(diagramstructureTo.Data_Shapes, CopyBranchDataShape(mapOrigCopy, _datashape))
	}
	for _, _dataflow := range diagramstructureFrom.DataFlowsWhoseDataNodeIsExpanded {
		diagramstructureTo.DataFlowsWhoseDataNodeIsExpanded = append(diagramstructureTo.DataFlowsWhoseDataNodeIsExpanded, CopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _resource := range diagramstructureFrom.AllocatedResourcesWhoseNodeIsExpanded {
		diagramstructureTo.AllocatedResourcesWhoseNodeIsExpanded = append(diagramstructureTo.AllocatedResourcesWhoseNodeIsExpanded, CopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _allocatedresourceshape := range diagramstructureFrom.AllocatedResourceShapes {
		diagramstructureTo.AllocatedResourceShapes = append(diagramstructureTo.AllocatedResourceShapes, CopyBranchAllocatedResourceShape(mapOrigCopy, _allocatedresourceshape))
	}
	for _, _system := range diagramstructureFrom.AllocatedSystemesWhoseNodeIsExpanded {
		diagramstructureTo.AllocatedSystemesWhoseNodeIsExpanded = append(diagramstructureTo.AllocatedSystemesWhoseNodeIsExpanded, CopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _allocatedsystemshape := range diagramstructureFrom.AllocatedSystemShapes {
		diagramstructureTo.AllocatedSystemShapes = append(diagramstructureTo.AllocatedSystemShapes, CopyBranchAllocatedSystemShape(mapOrigCopy, _allocatedsystemshape))
	}
	for _, _noteshape := range diagramstructureFrom.Note_Shapes {
		diagramstructureTo.Note_Shapes = append(diagramstructureTo.Note_Shapes, CopyBranchNoteShape(mapOrigCopy, _noteshape))
	}
	for _, _note := range diagramstructureFrom.NotesWhoseNodeIsExpanded {
		diagramstructureTo.NotesWhoseNodeIsExpanded = append(diagramstructureTo.NotesWhoseNodeIsExpanded, CopyBranchNote(mapOrigCopy, _note))
	}
	for _, _noteportshape := range diagramstructureFrom.NotePortShapes {
		diagramstructureTo.NotePortShapes = append(diagramstructureTo.NotePortShapes, CopyBranchNotePortShape(mapOrigCopy, _noteportshape))
	}

	return
}

func CopyBranchExternalPartShape(mapOrigCopy map[any]any, externalpartshapeFrom *ExternalPartShape) (externalpartshapeTo *ExternalPartShape) {

	// externalpartshapeFrom has already been copied
	if _externalpartshapeTo, ok := mapOrigCopy[externalpartshapeFrom]; ok {
		externalpartshapeTo = _externalpartshapeTo.(*ExternalPartShape)
		return
	}

	externalpartshapeTo = new(ExternalPartShape)
	mapOrigCopy[externalpartshapeFrom] = externalpartshapeTo
	externalpartshapeFrom.CopyBasicFields(externalpartshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if externalpartshapeFrom.Part != nil {
		externalpartshapeTo.Part = CopyBranchPart(mapOrigCopy, externalpartshapeFrom.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLibrary(mapOrigCopy map[any]any, libraryFrom *Library) (libraryTo *Library) {

	// libraryFrom has already been copied
	if _libraryTo, ok := mapOrigCopy[libraryFrom]; ok {
		libraryTo = _libraryTo.(*Library)
		return
	}

	libraryTo = new(Library)
	mapOrigCopy[libraryFrom] = libraryTo
	libraryFrom.CopyBasicFields(libraryTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range libraryFrom.SubLibraries {
		libraryTo.SubLibraries = append(libraryTo.SubLibraries, CopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _library := range libraryFrom.SubLibrariesWhoseNodeIsExpanded {
		libraryTo.SubLibrariesWhoseNodeIsExpanded = append(libraryTo.SubLibrariesWhoseNodeIsExpanded, CopyBranchLibrary(mapOrigCopy, _library))
	}
	for _, _system := range libraryFrom.RootSystemes {
		libraryTo.RootSystemes = append(libraryTo.RootSystemes, CopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _system := range libraryFrom.SystemsWhoseNodeIsExpanded {
		libraryTo.SystemsWhoseNodeIsExpanded = append(libraryTo.SystemsWhoseNodeIsExpanded, CopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _dataflow := range libraryFrom.RootDataFlows {
		libraryTo.RootDataFlows = append(libraryTo.RootDataFlows, CopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _dataflow := range libraryFrom.DataFlowsWhoseNodeIsExpanded {
		libraryTo.DataFlowsWhoseNodeIsExpanded = append(libraryTo.DataFlowsWhoseNodeIsExpanded, CopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _data := range libraryFrom.RootDatas {
		libraryTo.RootDatas = append(libraryTo.RootDatas, CopyBranchData(mapOrigCopy, _data))
	}
	for _, _data := range libraryFrom.DatasWhoseNodeIsExpanded {
		libraryTo.DatasWhoseNodeIsExpanded = append(libraryTo.DatasWhoseNodeIsExpanded, CopyBranchData(mapOrigCopy, _data))
	}
	for _, _resource := range libraryFrom.RootResources {
		libraryTo.RootResources = append(libraryTo.RootResources, CopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _resource := range libraryFrom.ResourcesWhoseNodeIsExpanded {
		libraryTo.ResourcesWhoseNodeIsExpanded = append(libraryTo.ResourcesWhoseNodeIsExpanded, CopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _part := range libraryFrom.PartsWhoseNodeIsExpanded {
		libraryTo.PartsWhoseNodeIsExpanded = append(libraryTo.PartsWhoseNodeIsExpanded, CopyBranchPart(mapOrigCopy, _part))
	}
	for _, _note := range libraryFrom.RootNotes {
		libraryTo.RootNotes = append(libraryTo.RootNotes, CopyBranchNote(mapOrigCopy, _note))
	}
	for _, _note := range libraryFrom.NotesWhoseNodeIsExpanded {
		libraryTo.NotesWhoseNodeIsExpanded = append(libraryTo.NotesWhoseNodeIsExpanded, CopyBranchNote(mapOrigCopy, _note))
	}

	return
}

func CopyBranchNote(mapOrigCopy map[any]any, noteFrom *Note) (noteTo *Note) {

	// noteFrom has already been copied
	if _noteTo, ok := mapOrigCopy[noteFrom]; ok {
		noteTo = _noteTo.(*Note)
		return
	}

	noteTo = new(Note)
	mapOrigCopy[noteFrom] = noteTo
	noteFrom.CopyBasicFields(noteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _port := range noteFrom.Ports {
		noteTo.Ports = append(noteTo.Ports, CopyBranchPort(mapOrigCopy, _port))
	}

	return
}

func CopyBranchNotePortShape(mapOrigCopy map[any]any, noteportshapeFrom *NotePortShape) (noteportshapeTo *NotePortShape) {

	// noteportshapeFrom has already been copied
	if _noteportshapeTo, ok := mapOrigCopy[noteportshapeFrom]; ok {
		noteportshapeTo = _noteportshapeTo.(*NotePortShape)
		return
	}

	noteportshapeTo = new(NotePortShape)
	mapOrigCopy[noteportshapeFrom] = noteportshapeTo
	noteportshapeFrom.CopyBasicFields(noteportshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteportshapeFrom.Note != nil {
		noteportshapeTo.Note = CopyBranchNote(mapOrigCopy, noteportshapeFrom.Note)
	}
	if noteportshapeFrom.Port != nil {
		noteportshapeTo.Port = CopyBranchPort(mapOrigCopy, noteportshapeFrom.Port)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchNoteShape(mapOrigCopy map[any]any, noteshapeFrom *NoteShape) (noteshapeTo *NoteShape) {

	// noteshapeFrom has already been copied
	if _noteshapeTo, ok := mapOrigCopy[noteshapeFrom]; ok {
		noteshapeTo = _noteshapeTo.(*NoteShape)
		return
	}

	noteshapeTo = new(NoteShape)
	mapOrigCopy[noteshapeFrom] = noteshapeTo
	noteshapeFrom.CopyBasicFields(noteshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if noteshapeFrom.Note != nil {
		noteshapeTo.Note = CopyBranchNote(mapOrigCopy, noteshapeFrom.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPart(mapOrigCopy map[any]any, partFrom *Part) (partTo *Part) {

	// partFrom has already been copied
	if _partTo, ok := mapOrigCopy[partFrom]; ok {
		partTo = _partTo.(*Part)
		return
	}

	partTo = new(Part)
	mapOrigCopy[partFrom] = partTo
	partFrom.CopyBasicFields(partTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _resource := range partFrom.Resources {
		partTo.Resources = append(partTo.Resources, CopyBranchResource(mapOrigCopy, _resource))
	}
	for _, _system := range partFrom.Systemes {
		partTo.Systemes = append(partTo.Systemes, CopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _port := range partFrom.Ports {
		partTo.Ports = append(partTo.Ports, CopyBranchPort(mapOrigCopy, _port))
	}
	for _, _controlflow := range partFrom.ControlFlows {
		partTo.ControlFlows = append(partTo.ControlFlows, CopyBranchControlFlow(mapOrigCopy, _controlflow))
	}
	for _, _port := range partFrom.PortWhoseOutControlFlowsNodeIsExpanded {
		partTo.PortWhoseOutControlFlowsNodeIsExpanded = append(partTo.PortWhoseOutControlFlowsNodeIsExpanded, CopyBranchPort(mapOrigCopy, _port))
	}
	for _, _port := range partFrom.PortWhoseInControlFlowsNodeIsExpanded {
		partTo.PortWhoseInControlFlowsNodeIsExpanded = append(partTo.PortWhoseInControlFlowsNodeIsExpanded, CopyBranchPort(mapOrigCopy, _port))
	}
	for _, _port := range partFrom.PortWhoseOutDataFlowsNodeIsExpanded {
		partTo.PortWhoseOutDataFlowsNodeIsExpanded = append(partTo.PortWhoseOutDataFlowsNodeIsExpanded, CopyBranchPort(mapOrigCopy, _port))
	}
	for _, _port := range partFrom.PortWhoseInDataFlowsNodeIsExpanded {
		partTo.PortWhoseInDataFlowsNodeIsExpanded = append(partTo.PortWhoseInDataFlowsNodeIsExpanded, CopyBranchPort(mapOrigCopy, _port))
	}

	return
}

func CopyBranchPartShape(mapOrigCopy map[any]any, partshapeFrom *PartShape) (partshapeTo *PartShape) {

	// partshapeFrom has already been copied
	if _partshapeTo, ok := mapOrigCopy[partshapeFrom]; ok {
		partshapeTo = _partshapeTo.(*PartShape)
		return
	}

	partshapeTo = new(PartShape)
	mapOrigCopy[partshapeFrom] = partshapeTo
	partshapeFrom.CopyBasicFields(partshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if partshapeFrom.Part != nil {
		partshapeTo.Part = CopyBranchPart(mapOrigCopy, partshapeFrom.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPort(mapOrigCopy map[any]any, portFrom *Port) (portTo *Port) {

	// portFrom has already been copied
	if _portTo, ok := mapOrigCopy[portFrom]; ok {
		portTo = _portTo.(*Port)
		return
	}

	portTo = new(Port)
	mapOrigCopy[portFrom] = portTo
	portFrom.CopyBasicFields(portTo)

	//insertion point for the staging of instances referenced by pointers
	if portFrom.Type != nil {
		portTo.Type = CopyBranchSystem(mapOrigCopy, portFrom.Type)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPortShape(mapOrigCopy map[any]any, portshapeFrom *PortShape) (portshapeTo *PortShape) {

	// portshapeFrom has already been copied
	if _portshapeTo, ok := mapOrigCopy[portshapeFrom]; ok {
		portshapeTo = _portshapeTo.(*PortShape)
		return
	}

	portshapeTo = new(PortShape)
	mapOrigCopy[portshapeFrom] = portshapeTo
	portshapeFrom.CopyBasicFields(portshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if portshapeFrom.Port != nil {
		portshapeTo.Port = CopyBranchPort(mapOrigCopy, portshapeFrom.Port)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchResource(mapOrigCopy map[any]any, resourceFrom *Resource) (resourceTo *Resource) {

	// resourceFrom has already been copied
	if _resourceTo, ok := mapOrigCopy[resourceFrom]; ok {
		resourceTo = _resourceTo.(*Resource)
		return
	}

	resourceTo = new(Resource)
	mapOrigCopy[resourceFrom] = resourceTo
	resourceFrom.CopyBasicFields(resourceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSystem(mapOrigCopy map[any]any, systemFrom *System) (systemTo *System) {

	// systemFrom has already been copied
	if _systemTo, ok := mapOrigCopy[systemFrom]; ok {
		systemTo = _systemTo.(*System)
		return
	}

	systemTo = new(System)
	mapOrigCopy[systemFrom] = systemTo
	systemFrom.CopyBasicFields(systemTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagramstructure := range systemFrom.DiagramStructures {
		systemTo.DiagramStructures = append(systemTo.DiagramStructures, CopyBranchDiagramStructure(mapOrigCopy, _diagramstructure))
	}
	for _, _diagramstructure := range systemFrom.DiagramStructureWhoseNodeIsExpanded {
		systemTo.DiagramStructureWhoseNodeIsExpanded = append(systemTo.DiagramStructureWhoseNodeIsExpanded, CopyBranchDiagramStructure(mapOrigCopy, _diagramstructure))
	}
	for _, _system := range systemFrom.SubSystemes {
		systemTo.SubSystemes = append(systemTo.SubSystemes, CopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _part := range systemFrom.Parts {
		systemTo.Parts = append(systemTo.Parts, CopyBranchPart(mapOrigCopy, _part))
	}
	for _, _part := range systemFrom.PartWhoseNodeIsExpanded {
		systemTo.PartWhoseNodeIsExpanded = append(systemTo.PartWhoseNodeIsExpanded, CopyBranchPart(mapOrigCopy, _part))
	}
	for _, _dataflow := range systemFrom.DataFlows {
		systemTo.DataFlows = append(systemTo.DataFlows, CopyBranchDataFlow(mapOrigCopy, _dataflow))
	}
	for _, _part := range systemFrom.ExternalParts {
		systemTo.ExternalParts = append(systemTo.ExternalParts, CopyBranchPart(mapOrigCopy, _part))
	}
	for _, _part := range systemFrom.ExternalPartWhoseNodeIsExpanded {
		systemTo.ExternalPartWhoseNodeIsExpanded = append(systemTo.ExternalPartWhoseNodeIsExpanded, CopyBranchPart(mapOrigCopy, _part))
	}

	return
}

func CopyBranchSystemShape(mapOrigCopy map[any]any, systemshapeFrom *SystemShape) (systemshapeTo *SystemShape) {

	// systemshapeFrom has already been copied
	if _systemshapeTo, ok := mapOrigCopy[systemshapeFrom]; ok {
		systemshapeTo = _systemshapeTo.(*SystemShape)
		return
	}

	systemshapeTo = new(SystemShape)
	mapOrigCopy[systemshapeFrom] = systemshapeTo
	systemshapeFrom.CopyBasicFields(systemshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if systemshapeFrom.System != nil {
		systemshapeTo.System = CopyBranchSystem(mapOrigCopy, systemshapeFrom.System)
	}

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
	case *AllocatedResourceShape:
		stage.UnstageBranchAllocatedResourceShape(target)

	case *AllocatedSystemShape:
		stage.UnstageBranchAllocatedSystemShape(target)

	case *ControlFlow:
		stage.UnstageBranchControlFlow(target)

	case *ControlFlowShape:
		stage.UnstageBranchControlFlowShape(target)

	case *Data:
		stage.UnstageBranchData(target)

	case *DataFlow:
		stage.UnstageBranchDataFlow(target)

	case *DataFlowShape:
		stage.UnstageBranchDataFlowShape(target)

	case *DataShape:
		stage.UnstageBranchDataShape(target)

	case *DiagramStructure:
		stage.UnstageBranchDiagramStructure(target)

	case *ExternalPartShape:
		stage.UnstageBranchExternalPartShape(target)

	case *Library:
		stage.UnstageBranchLibrary(target)

	case *Note:
		stage.UnstageBranchNote(target)

	case *NotePortShape:
		stage.UnstageBranchNotePortShape(target)

	case *NoteShape:
		stage.UnstageBranchNoteShape(target)

	case *Part:
		stage.UnstageBranchPart(target)

	case *PartShape:
		stage.UnstageBranchPartShape(target)

	case *Port:
		stage.UnstageBranchPort(target)

	case *PortShape:
		stage.UnstageBranchPortShape(target)

	case *Resource:
		stage.UnstageBranchResource(target)

	case *System:
		stage.UnstageBranchSystem(target)

	case *SystemShape:
		stage.UnstageBranchSystemShape(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchAllocatedResourceShape(allocatedresourceshape *AllocatedResourceShape) {

	// check if instance is already staged
	if !IsStaged(stage, allocatedresourceshape) {
		return
	}

	allocatedresourceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if allocatedresourceshape.Part != nil {
		UnstageBranch(stage, allocatedresourceshape.Part)
	}
	if allocatedresourceshape.Resource != nil {
		UnstageBranch(stage, allocatedresourceshape.Resource)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchAllocatedSystemShape(allocatedsystemshape *AllocatedSystemShape) {

	// check if instance is already staged
	if !IsStaged(stage, allocatedsystemshape) {
		return
	}

	allocatedsystemshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if allocatedsystemshape.Part != nil {
		UnstageBranch(stage, allocatedsystemshape.Part)
	}
	if allocatedsystemshape.System != nil {
		UnstageBranch(stage, allocatedsystemshape.System)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchControlFlow(controlflow *ControlFlow) {

	// check if instance is already staged
	if !IsStaged(stage, controlflow) {
		return
	}

	controlflow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlflow.Start != nil {
		UnstageBranch(stage, controlflow.Start)
	}
	if controlflow.End != nil {
		UnstageBranch(stage, controlflow.End)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchControlFlowShape(controlflowshape *ControlFlowShape) {

	// check if instance is already staged
	if !IsStaged(stage, controlflowshape) {
		return
	}

	controlflowshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if controlflowshape.ControlFlow != nil {
		UnstageBranch(stage, controlflowshape.ControlFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchData(data *Data) {

	// check if instance is already staged
	if !IsStaged(stage, data) {
		return
	}

	data.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDataFlow(dataflow *DataFlow) {

	// check if instance is already staged
	if !IsStaged(stage, dataflow) {
		return
	}

	dataflow.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dataflow.StartPort != nil {
		UnstageBranch(stage, dataflow.StartPort)
	}
	if dataflow.EndPort != nil {
		UnstageBranch(stage, dataflow.EndPort)
	}
	if dataflow.StartExternalPart != nil {
		UnstageBranch(stage, dataflow.StartExternalPart)
	}
	if dataflow.EndExternalPart != nil {
		UnstageBranch(stage, dataflow.EndExternalPart)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _data := range dataflow.Datas {
		UnstageBranch(stage, _data)
	}

}

func (stage *Stage) UnstageBranchDataFlowShape(dataflowshape *DataFlowShape) {

	// check if instance is already staged
	if !IsStaged(stage, dataflowshape) {
		return
	}

	dataflowshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dataflowshape.DataFlow != nil {
		UnstageBranch(stage, dataflowshape.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDataShape(datashape *DataShape) {

	// check if instance is already staged
	if !IsStaged(stage, datashape) {
		return
	}

	datashape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datashape.Data != nil {
		UnstageBranch(stage, datashape.Data)
	}
	if datashape.DataFlow != nil {
		UnstageBranch(stage, datashape.DataFlow)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDiagramStructure(diagramstructure *DiagramStructure) {

	// check if instance is already staged
	if !IsStaged(stage, diagramstructure) {
		return
	}

	diagramstructure.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _systemshape := range diagramstructure.System_Shapes {
		UnstageBranch(stage, _systemshape)
	}
	for _, _system := range diagramstructure.SystemsWhoseNodeIsExpanded {
		UnstageBranch(stage, _system)
	}
	for _, _partshape := range diagramstructure.Part_Shapes {
		UnstageBranch(stage, _partshape)
	}
	for _, _part := range diagramstructure.PartWhoseNodeIsExpanded {
		UnstageBranch(stage, _part)
	}
	for _, _externalpartshape := range diagramstructure.ExternalPart_Shapes {
		UnstageBranch(stage, _externalpartshape)
	}
	for _, _part := range diagramstructure.ExternalPartWhoseNodeIsExpanded {
		UnstageBranch(stage, _part)
	}
	for _, _part := range diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded {
		UnstageBranch(stage, _part)
	}
	for _, _part := range diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded {
		UnstageBranch(stage, _part)
	}
	for _, _port := range diagramstructure.PortsWhoseNodeIsExpanded {
		UnstageBranch(stage, _port)
	}
	for _, _portshape := range diagramstructure.Port_Shapes {
		UnstageBranch(stage, _portshape)
	}
	for _, _controlflow := range diagramstructure.ControlFlowsWhoseNodeIsExpanded {
		UnstageBranch(stage, _controlflow)
	}
	for _, _controlflowshape := range diagramstructure.ControlFlow_Shapes {
		UnstageBranch(stage, _controlflowshape)
	}
	for _, _dataflow := range diagramstructure.DataFlowsWhoseNodeIsExpanded {
		UnstageBranch(stage, _dataflow)
	}
	for _, _dataflowshape := range diagramstructure.DataFlow_Shapes {
		UnstageBranch(stage, _dataflowshape)
	}
	for _, _data := range diagramstructure.DatasWhoseNodeIsExpanded {
		UnstageBranch(stage, _data)
	}
	for _, _datashape := range diagramstructure.Data_Shapes {
		UnstageBranch(stage, _datashape)
	}
	for _, _dataflow := range diagramstructure.DataFlowsWhoseDataNodeIsExpanded {
		UnstageBranch(stage, _dataflow)
	}
	for _, _resource := range diagramstructure.AllocatedResourcesWhoseNodeIsExpanded {
		UnstageBranch(stage, _resource)
	}
	for _, _allocatedresourceshape := range diagramstructure.AllocatedResourceShapes {
		UnstageBranch(stage, _allocatedresourceshape)
	}
	for _, _system := range diagramstructure.AllocatedSystemesWhoseNodeIsExpanded {
		UnstageBranch(stage, _system)
	}
	for _, _allocatedsystemshape := range diagramstructure.AllocatedSystemShapes {
		UnstageBranch(stage, _allocatedsystemshape)
	}
	for _, _noteshape := range diagramstructure.Note_Shapes {
		UnstageBranch(stage, _noteshape)
	}
	for _, _note := range diagramstructure.NotesWhoseNodeIsExpanded {
		UnstageBranch(stage, _note)
	}
	for _, _noteportshape := range diagramstructure.NotePortShapes {
		UnstageBranch(stage, _noteportshape)
	}

}

func (stage *Stage) UnstageBranchExternalPartShape(externalpartshape *ExternalPartShape) {

	// check if instance is already staged
	if !IsStaged(stage, externalpartshape) {
		return
	}

	externalpartshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if externalpartshape.Part != nil {
		UnstageBranch(stage, externalpartshape.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchLibrary(library *Library) {

	// check if instance is already staged
	if !IsStaged(stage, library) {
		return
	}

	library.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _library := range library.SubLibraries {
		UnstageBranch(stage, _library)
	}
	for _, _library := range library.SubLibrariesWhoseNodeIsExpanded {
		UnstageBranch(stage, _library)
	}
	for _, _system := range library.RootSystemes {
		UnstageBranch(stage, _system)
	}
	for _, _system := range library.SystemsWhoseNodeIsExpanded {
		UnstageBranch(stage, _system)
	}
	for _, _dataflow := range library.RootDataFlows {
		UnstageBranch(stage, _dataflow)
	}
	for _, _dataflow := range library.DataFlowsWhoseNodeIsExpanded {
		UnstageBranch(stage, _dataflow)
	}
	for _, _data := range library.RootDatas {
		UnstageBranch(stage, _data)
	}
	for _, _data := range library.DatasWhoseNodeIsExpanded {
		UnstageBranch(stage, _data)
	}
	for _, _resource := range library.RootResources {
		UnstageBranch(stage, _resource)
	}
	for _, _resource := range library.ResourcesWhoseNodeIsExpanded {
		UnstageBranch(stage, _resource)
	}
	for _, _part := range library.PartsWhoseNodeIsExpanded {
		UnstageBranch(stage, _part)
	}
	for _, _note := range library.RootNotes {
		UnstageBranch(stage, _note)
	}
	for _, _note := range library.NotesWhoseNodeIsExpanded {
		UnstageBranch(stage, _note)
	}

}

func (stage *Stage) UnstageBranchNote(note *Note) {

	// check if instance is already staged
	if !IsStaged(stage, note) {
		return
	}

	note.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _port := range note.Ports {
		UnstageBranch(stage, _port)
	}

}

func (stage *Stage) UnstageBranchNotePortShape(noteportshape *NotePortShape) {

	// check if instance is already staged
	if !IsStaged(stage, noteportshape) {
		return
	}

	noteportshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteportshape.Note != nil {
		UnstageBranch(stage, noteportshape.Note)
	}
	if noteportshape.Port != nil {
		UnstageBranch(stage, noteportshape.Port)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchNoteShape(noteshape *NoteShape) {

	// check if instance is already staged
	if !IsStaged(stage, noteshape) {
		return
	}

	noteshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if noteshape.Note != nil {
		UnstageBranch(stage, noteshape.Note)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPart(part *Part) {

	// check if instance is already staged
	if !IsStaged(stage, part) {
		return
	}

	part.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _resource := range part.Resources {
		UnstageBranch(stage, _resource)
	}
	for _, _system := range part.Systemes {
		UnstageBranch(stage, _system)
	}
	for _, _port := range part.Ports {
		UnstageBranch(stage, _port)
	}
	for _, _controlflow := range part.ControlFlows {
		UnstageBranch(stage, _controlflow)
	}
	for _, _port := range part.PortWhoseOutControlFlowsNodeIsExpanded {
		UnstageBranch(stage, _port)
	}
	for _, _port := range part.PortWhoseInControlFlowsNodeIsExpanded {
		UnstageBranch(stage, _port)
	}
	for _, _port := range part.PortWhoseOutDataFlowsNodeIsExpanded {
		UnstageBranch(stage, _port)
	}
	for _, _port := range part.PortWhoseInDataFlowsNodeIsExpanded {
		UnstageBranch(stage, _port)
	}

}

func (stage *Stage) UnstageBranchPartShape(partshape *PartShape) {

	// check if instance is already staged
	if !IsStaged(stage, partshape) {
		return
	}

	partshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if partshape.Part != nil {
		UnstageBranch(stage, partshape.Part)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPort(port *Port) {

	// check if instance is already staged
	if !IsStaged(stage, port) {
		return
	}

	port.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if port.Type != nil {
		UnstageBranch(stage, port.Type)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPortShape(portshape *PortShape) {

	// check if instance is already staged
	if !IsStaged(stage, portshape) {
		return
	}

	portshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if portshape.Port != nil {
		UnstageBranch(stage, portshape.Port)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchResource(resource *Resource) {

	// check if instance is already staged
	if !IsStaged(stage, resource) {
		return
	}

	resource.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSystem(system *System) {

	// check if instance is already staged
	if !IsStaged(stage, system) {
		return
	}

	system.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _diagramstructure := range system.DiagramStructures {
		UnstageBranch(stage, _diagramstructure)
	}
	for _, _diagramstructure := range system.DiagramStructureWhoseNodeIsExpanded {
		UnstageBranch(stage, _diagramstructure)
	}
	for _, _system := range system.SubSystemes {
		UnstageBranch(stage, _system)
	}
	for _, _part := range system.Parts {
		UnstageBranch(stage, _part)
	}
	for _, _part := range system.PartWhoseNodeIsExpanded {
		UnstageBranch(stage, _part)
	}
	for _, _dataflow := range system.DataFlows {
		UnstageBranch(stage, _dataflow)
	}
	for _, _part := range system.ExternalParts {
		UnstageBranch(stage, _part)
	}
	for _, _part := range system.ExternalPartWhoseNodeIsExpanded {
		UnstageBranch(stage, _part)
	}

}

func (stage *Stage) UnstageBranchSystemShape(systemshape *SystemShape) {

	// check if instance is already staged
	if !IsStaged(stage, systemshape) {
		return
	}

	systemshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if systemshape.System != nil {
		UnstageBranch(stage, systemshape.System)
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
}

func (reference *ExternalPartShape) GongReconstructPointersFromReferences(stage *Stage, instance *ExternalPartShape) {
	// insertion point for pointers field
	if instance.Part != nil {
		reference.Part = stage.Parts_reference[instance.Part]
	}
	// insertion point for slice of pointers field
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
	reference.Ports = reference.Ports[:0]
	for _, _b := range instance.Ports {
		reference.Ports = append(reference.Ports, stage.Ports_reference[_b])
	}
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
	// insertion point for slice of pointers field
	reference.Resources = reference.Resources[:0]
	for _, _b := range instance.Resources {
		reference.Resources = append(reference.Resources, stage.Resources_reference[_b])
	}
	reference.Systemes = reference.Systemes[:0]
	for _, _b := range instance.Systemes {
		reference.Systemes = append(reference.Systemes, stage.Systems_reference[_b])
	}
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
	if instance.Type != nil {
		reference.Type = stage.Systems_reference[instance.Type]
	}
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
	var _Ports []*Port
	for _, _reference := range reference.Ports {
		if _instance, ok := stage.Ports_instance[_reference]; ok {
			_Ports = append(_Ports, _instance)
		}
	}
	reference.Ports = _Ports
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
	// insertion point for slice of pointers fields
	var _Resources []*Resource
	for _, _reference := range reference.Resources {
		if _instance, ok := stage.Resources_instance[_reference]; ok {
			_Resources = append(_Resources, _instance)
		}
	}
	reference.Resources = _Resources
	var _Systemes []*System
	for _, _reference := range reference.Systemes {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_Systemes = append(_Systemes, _instance)
		}
	}
	reference.Systemes = _Systemes
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
	if _reference := reference.Type; _reference != nil {
		reference.Type = nil
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			reference.Type = _instance
		}
	}
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
	if controlflow.LayoutDirection != controlflowOther.LayoutDirection {
		diffs = append(diffs, controlflow.GongMarshallField(stage, "LayoutDirection"))
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
	if data.LayoutDirection != dataOther.LayoutDirection {
		diffs = append(diffs, data.GongMarshallField(stage, "LayoutDirection"))
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
		ops := Diff(stage, dataflow, dataflowOther, "Datas", dataflowOther.Datas, dataflow.Datas)
		diffs = append(diffs, ops)
	}
	if dataflow.Description != dataflowOther.Description {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "Description"))
	}
	if dataflow.ComputedPrefix != dataflowOther.ComputedPrefix {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "ComputedPrefix"))
	}
	if dataflow.IsExpanded != dataflowOther.IsExpanded {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "IsExpanded"))
	}
	if dataflow.LayoutDirection != dataflowOther.LayoutDirection {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "LayoutDirection"))
	}
	if dataflow.Type != dataflowOther.Type {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "Type"))
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
	if dataflow.IsDatasNodeExpanded != dataflowOther.IsDatasNodeExpanded {
		diffs = append(diffs, dataflow.GongMarshallField(stage, "IsDatasNodeExpanded"))
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
	if diagramstructure.LayoutDirection != diagramstructureOther.LayoutDirection {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "LayoutDirection"))
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "System_Shapes", diagramstructureOther.System_Shapes, diagramstructure.System_Shapes)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "SystemsWhoseNodeIsExpanded", diagramstructureOther.SystemsWhoseNodeIsExpanded, diagramstructure.SystemsWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "Part_Shapes", diagramstructureOther.Part_Shapes, diagramstructure.Part_Shapes)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "PartWhoseNodeIsExpanded", diagramstructureOther.PartWhoseNodeIsExpanded, diagramstructure.PartWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "ExternalPart_Shapes", diagramstructureOther.ExternalPart_Shapes, diagramstructure.ExternalPart_Shapes)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "ExternalPartWhoseNodeIsExpanded", diagramstructureOther.ExternalPartWhoseNodeIsExpanded, diagramstructure.ExternalPartWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "ExternalPartsWhoseOutDataFlowsNodeIsExpanded", diagramstructureOther.ExternalPartsWhoseOutDataFlowsNodeIsExpanded, diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "ExternalPartsWhoseInDataFlowsNodeIsExpanded", diagramstructureOther.ExternalPartsWhoseInDataFlowsNodeIsExpanded, diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "PortsWhoseNodeIsExpanded", diagramstructureOther.PortsWhoseNodeIsExpanded, diagramstructure.PortsWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "Port_Shapes", diagramstructureOther.Port_Shapes, diagramstructure.Port_Shapes)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "ControlFlowsWhoseNodeIsExpanded", diagramstructureOther.ControlFlowsWhoseNodeIsExpanded, diagramstructure.ControlFlowsWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "ControlFlow_Shapes", diagramstructureOther.ControlFlow_Shapes, diagramstructure.ControlFlow_Shapes)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "DataFlowsWhoseNodeIsExpanded", diagramstructureOther.DataFlowsWhoseNodeIsExpanded, diagramstructure.DataFlowsWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "DataFlow_Shapes", diagramstructureOther.DataFlow_Shapes, diagramstructure.DataFlow_Shapes)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "DatasWhoseNodeIsExpanded", diagramstructureOther.DatasWhoseNodeIsExpanded, diagramstructure.DatasWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "Data_Shapes", diagramstructureOther.Data_Shapes, diagramstructure.Data_Shapes)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "DataFlowsWhoseDataNodeIsExpanded", diagramstructureOther.DataFlowsWhoseDataNodeIsExpanded, diagramstructure.DataFlowsWhoseDataNodeIsExpanded)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "AllocatedResourcesWhoseNodeIsExpanded", diagramstructureOther.AllocatedResourcesWhoseNodeIsExpanded, diagramstructure.AllocatedResourcesWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "AllocatedResourceShapes", diagramstructureOther.AllocatedResourceShapes, diagramstructure.AllocatedResourceShapes)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "AllocatedSystemesWhoseNodeIsExpanded", diagramstructureOther.AllocatedSystemesWhoseNodeIsExpanded, diagramstructure.AllocatedSystemesWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "AllocatedSystemShapes", diagramstructureOther.AllocatedSystemShapes, diagramstructure.AllocatedSystemShapes)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "Note_Shapes", diagramstructureOther.Note_Shapes, diagramstructure.Note_Shapes)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "NotesWhoseNodeIsExpanded", diagramstructureOther.NotesWhoseNodeIsExpanded, diagramstructure.NotesWhoseNodeIsExpanded)
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
		ops := Diff(stage, diagramstructure, diagramstructureOther, "NotePortShapes", diagramstructureOther.NotePortShapes, diagramstructure.NotePortShapes)
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
	if library.LayoutDirection != libraryOther.LayoutDirection {
		diffs = append(diffs, library.GongMarshallField(stage, "LayoutDirection"))
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
		ops := Diff(stage, library, libraryOther, "SubLibraries", libraryOther.SubLibraries, library.SubLibraries)
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
		ops := Diff(stage, library, libraryOther, "SubLibrariesWhoseNodeIsExpanded", libraryOther.SubLibrariesWhoseNodeIsExpanded, library.SubLibrariesWhoseNodeIsExpanded)
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
		ops := Diff(stage, library, libraryOther, "RootSystemes", libraryOther.RootSystemes, library.RootSystemes)
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
		ops := Diff(stage, library, libraryOther, "SystemsWhoseNodeIsExpanded", libraryOther.SystemsWhoseNodeIsExpanded, library.SystemsWhoseNodeIsExpanded)
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
		ops := Diff(stage, library, libraryOther, "RootDataFlows", libraryOther.RootDataFlows, library.RootDataFlows)
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
		ops := Diff(stage, library, libraryOther, "DataFlowsWhoseNodeIsExpanded", libraryOther.DataFlowsWhoseNodeIsExpanded, library.DataFlowsWhoseNodeIsExpanded)
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
		ops := Diff(stage, library, libraryOther, "RootDatas", libraryOther.RootDatas, library.RootDatas)
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
		ops := Diff(stage, library, libraryOther, "DatasWhoseNodeIsExpanded", libraryOther.DatasWhoseNodeIsExpanded, library.DatasWhoseNodeIsExpanded)
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
		ops := Diff(stage, library, libraryOther, "RootResources", libraryOther.RootResources, library.RootResources)
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
		ops := Diff(stage, library, libraryOther, "ResourcesWhoseNodeIsExpanded", libraryOther.ResourcesWhoseNodeIsExpanded, library.ResourcesWhoseNodeIsExpanded)
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
		ops := Diff(stage, library, libraryOther, "PartsWhoseNodeIsExpanded", libraryOther.PartsWhoseNodeIsExpanded, library.PartsWhoseNodeIsExpanded)
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
		ops := Diff(stage, library, libraryOther, "RootNotes", libraryOther.RootNotes, library.RootNotes)
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
		ops := Diff(stage, library, libraryOther, "NotesWhoseNodeIsExpanded", libraryOther.NotesWhoseNodeIsExpanded, library.NotesWhoseNodeIsExpanded)
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
	if note.LayoutDirection != noteOther.LayoutDirection {
		diffs = append(diffs, note.GongMarshallField(stage, "LayoutDirection"))
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
		ops := Diff(stage, note, noteOther, "Ports", noteOther.Ports, note.Ports)
		diffs = append(diffs, ops)
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
	if part.IsSystemResource != partOther.IsSystemResource {
		diffs = append(diffs, part.GongMarshallField(stage, "IsSystemResource"))
	}
	if part.Description != partOther.Description {
		diffs = append(diffs, part.GongMarshallField(stage, "Description"))
	}
	ResourcesDifferent := false
	if len(part.Resources) != len(partOther.Resources) {
		ResourcesDifferent = true
	} else {
		for i := range part.Resources {
			if (part.Resources[i] == nil) != (partOther.Resources[i] == nil) {
				ResourcesDifferent = true
				break
			} else if part.Resources[i] != nil && partOther.Resources[i] != nil {
				// this is a pointer comparaison
				if part.Resources[i] != partOther.Resources[i] {
					ResourcesDifferent = true
					break
				}
			}
		}
	}
	if ResourcesDifferent {
		ops := Diff(stage, part, partOther, "Resources", partOther.Resources, part.Resources)
		diffs = append(diffs, ops)
	}
	if part.IsResourcesNodeExpanded != partOther.IsResourcesNodeExpanded {
		diffs = append(diffs, part.GongMarshallField(stage, "IsResourcesNodeExpanded"))
	}
	SystemesDifferent := false
	if len(part.Systemes) != len(partOther.Systemes) {
		SystemesDifferent = true
	} else {
		for i := range part.Systemes {
			if (part.Systemes[i] == nil) != (partOther.Systemes[i] == nil) {
				SystemesDifferent = true
				break
			} else if part.Systemes[i] != nil && partOther.Systemes[i] != nil {
				// this is a pointer comparaison
				if part.Systemes[i] != partOther.Systemes[i] {
					SystemesDifferent = true
					break
				}
			}
		}
	}
	if SystemesDifferent {
		ops := Diff(stage, part, partOther, "Systemes", partOther.Systemes, part.Systemes)
		diffs = append(diffs, ops)
	}
	if part.IsSystemesNodeExpanded != partOther.IsSystemesNodeExpanded {
		diffs = append(diffs, part.GongMarshallField(stage, "IsSystemesNodeExpanded"))
	}
	if part.ComputedPrefix != partOther.ComputedPrefix {
		diffs = append(diffs, part.GongMarshallField(stage, "ComputedPrefix"))
	}
	if part.IsExpanded != partOther.IsExpanded {
		diffs = append(diffs, part.GongMarshallField(stage, "IsExpanded"))
	}
	if part.LayoutDirection != partOther.LayoutDirection {
		diffs = append(diffs, part.GongMarshallField(stage, "LayoutDirection"))
	}
	if part.IsPortsNodeExpanded != partOther.IsPortsNodeExpanded {
		diffs = append(diffs, part.GongMarshallField(stage, "IsPortsNodeExpanded"))
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
		ops := Diff(stage, part, partOther, "Ports", partOther.Ports, part.Ports)
		diffs = append(diffs, ops)
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
		ops := Diff(stage, part, partOther, "ControlFlows", partOther.ControlFlows, part.ControlFlows)
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
		ops := Diff(stage, part, partOther, "PortWhoseOutControlFlowsNodeIsExpanded", partOther.PortWhoseOutControlFlowsNodeIsExpanded, part.PortWhoseOutControlFlowsNodeIsExpanded)
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
		ops := Diff(stage, part, partOther, "PortWhoseInControlFlowsNodeIsExpanded", partOther.PortWhoseInControlFlowsNodeIsExpanded, part.PortWhoseInControlFlowsNodeIsExpanded)
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
		ops := Diff(stage, part, partOther, "PortWhoseOutDataFlowsNodeIsExpanded", partOther.PortWhoseOutDataFlowsNodeIsExpanded, part.PortWhoseOutDataFlowsNodeIsExpanded)
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
		ops := Diff(stage, part, partOther, "PortWhoseInDataFlowsNodeIsExpanded", partOther.PortWhoseInDataFlowsNodeIsExpanded, part.PortWhoseInDataFlowsNodeIsExpanded)
		diffs = append(diffs, ops)
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
	if partshape.WidthWeight != partshapeOther.WidthWeight {
		diffs = append(diffs, partshape.GongMarshallField(stage, "WidthWeight"))
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
	if port.LayoutDirection != portOther.LayoutDirection {
		diffs = append(diffs, port.GongMarshallField(stage, "LayoutDirection"))
	}
	if port.IsStartPort != portOther.IsStartPort {
		diffs = append(diffs, port.GongMarshallField(stage, "IsStartPort"))
	}
	if port.IsEndPort != portOther.IsEndPort {
		diffs = append(diffs, port.GongMarshallField(stage, "IsEndPort"))
	}
	if (port.Type == nil) != (portOther.Type == nil) {
		diffs = append(diffs, port.GongMarshallField(stage, "Type"))
	} else if port.Type != nil && portOther.Type != nil {
		if port.Type != portOther.Type {
			diffs = append(diffs, port.GongMarshallField(stage, "Type"))
		}
	}
	if port.IsPortNameNotSystemName != portOther.IsPortNameNotSystemName {
		diffs = append(diffs, port.GongMarshallField(stage, "IsPortNameNotSystemName"))
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
	if resource.LayoutDirection != resourceOther.LayoutDirection {
		diffs = append(diffs, resource.GongMarshallField(stage, "LayoutDirection"))
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
	if system.LayoutDirection != systemOther.LayoutDirection {
		diffs = append(diffs, system.GongMarshallField(stage, "LayoutDirection"))
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
		ops := Diff(stage, system, systemOther, "DiagramStructures", systemOther.DiagramStructures, system.DiagramStructures)
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
		ops := Diff(stage, system, systemOther, "DiagramStructureWhoseNodeIsExpanded", systemOther.DiagramStructureWhoseNodeIsExpanded, system.DiagramStructureWhoseNodeIsExpanded)
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
		ops := Diff(stage, system, systemOther, "SubSystemes", systemOther.SubSystemes, system.SubSystemes)
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
		ops := Diff(stage, system, systemOther, "Parts", systemOther.Parts, system.Parts)
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
		ops := Diff(stage, system, systemOther, "PartWhoseNodeIsExpanded", systemOther.PartWhoseNodeIsExpanded, system.PartWhoseNodeIsExpanded)
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
		ops := Diff(stage, system, systemOther, "DataFlows", systemOther.DataFlows, system.DataFlows)
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
		ops := Diff(stage, system, systemOther, "ExternalParts", systemOther.ExternalParts, system.ExternalParts)
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
		ops := Diff(stage, system, systemOther, "ExternalPartWhoseNodeIsExpanded", systemOther.ExternalPartWhoseNodeIsExpanded, system.ExternalPartWhoseNodeIsExpanded)
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
