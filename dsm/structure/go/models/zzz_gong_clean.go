// generated code - do not edit
package models

import "time"

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !IsStagedPointerToGongstruct(stage, *element) {
		*element = zero
		modified = true
		return
	}
	return
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by AllocatedResourceShape
func (allocatedresourceshape *AllocatedResourceShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &allocatedresourceshape.Part) || modified
	modified = GongCleanPointer(stage, &allocatedresourceshape.Resource) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AllocatedSystemShape
func (allocatedsystemshape *AllocatedSystemShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &allocatedsystemshape.Part) || modified
	modified = GongCleanPointer(stage, &allocatedsystemshape.System) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ControlFlow
func (controlflow *ControlFlow) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &controlflow.Start) || modified
	modified = GongCleanPointer(stage, &controlflow.End) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ControlFlowShape
func (controlflowshape *ControlFlowShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &controlflowshape.ControlFlow) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Data
func (data *Data) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DataFlow
func (dataflow *DataFlow) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &dataflow.Datas) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &dataflow.StartPort) || modified
	modified = GongCleanPointer(stage, &dataflow.EndPort) || modified
	modified = GongCleanPointer(stage, &dataflow.StartExternalPart) || modified
	modified = GongCleanPointer(stage, &dataflow.EndExternalPart) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DataFlowShape
func (dataflowshape *DataFlowShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &dataflowshape.DataFlow) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DataShape
func (datashape *DataShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &datashape.Data) || modified
	modified = GongCleanPointer(stage, &datashape.DataFlow) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DiagramStructure
func (diagramstructure *DiagramStructure) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &diagramstructure.System_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramstructure.SystemsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramstructure.Part_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramstructure.PartWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramstructure.ExternalPart_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramstructure.ExternalPartWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramstructure.PortsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramstructure.Port_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramstructure.ControlFlowsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramstructure.ControlFlow_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramstructure.DataFlowsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramstructure.DataFlow_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramstructure.DatasWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramstructure.Data_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramstructure.DataFlowsWhoseDataNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramstructure.AllocatedResourcesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramstructure.AllocatedResourceShapes) || modified
	modified = GongCleanSlice(stage, &diagramstructure.AllocatedSystemesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramstructure.AllocatedSystemShapes) || modified
	modified = GongCleanSlice(stage, &diagramstructure.Note_Shapes) || modified
	modified = GongCleanSlice(stage, &diagramstructure.NotesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &diagramstructure.NotePortShapes) || modified
	modified = GongCleanSlice(stage, &diagramstructure.NotePartShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ExternalPartShape
func (externalpartshape *ExternalPartShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &externalpartshape.Part) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &library.SubLibraries) || modified
	modified = GongCleanSlice(stage, &library.SubLibrariesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.RootSystemes) || modified
	modified = GongCleanSlice(stage, &library.SystemsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.RootDataFlows) || modified
	modified = GongCleanSlice(stage, &library.DataFlowsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.RootDatas) || modified
	modified = GongCleanSlice(stage, &library.DatasWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.RootResources) || modified
	modified = GongCleanSlice(stage, &library.ResourcesWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.PartsWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &library.RootNotes) || modified
	modified = GongCleanSlice(stage, &library.NotesWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Note
func (note *Note) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &note.Parts) || modified
	modified = GongCleanSlice(stage, &note.Ports) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NotePartShape
func (notepartshape *NotePartShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &notepartshape.Note) || modified
	modified = GongCleanPointer(stage, &notepartshape.Part) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NotePortShape
func (noteportshape *NotePortShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &noteportshape.Note) || modified
	modified = GongCleanPointer(stage, &noteportshape.Port) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteShape
func (noteshape *NoteShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &noteshape.Note) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Part
func (part *Part) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &part.Ports) || modified
	modified = GongCleanSlice(stage, &part.ControlFlows) || modified
	modified = GongCleanSlice(stage, &part.PortWhoseOutControlFlowsNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &part.PortWhoseInControlFlowsNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &part.PortWhoseOutDataFlowsNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &part.PortWhoseInDataFlowsNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &part.PartAnchoredPath) || modified
	// insertion point per field
	modified = GongCleanPointer(stage, &part.TypeOfPart) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by PartAnchoredPath
func (partanchoredpath *PartAnchoredPath) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PartShape
func (partshape *PartShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &partshape.Part) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Port
func (port *Port) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by PortShape
func (portshape *PortShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &portshape.Port) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Resource
func (resource *Resource) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by System
func (system *System) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &system.DiagramStructures) || modified
	modified = GongCleanSlice(stage, &system.DiagramStructureWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &system.SubSystemes) || modified
	modified = GongCleanSlice(stage, &system.Parts) || modified
	modified = GongCleanSlice(stage, &system.PartWhoseNodeIsExpanded) || modified
	modified = GongCleanSlice(stage, &system.DataFlows) || modified
	modified = GongCleanSlice(stage, &system.ExternalParts) || modified
	modified = GongCleanSlice(stage, &system.ExternalPartWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SystemShape
func (systemshape *SystemShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &systemshape.System) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	if modified {
		if stage.probeIF != nil {
			stage.probeIF.AddNotification(time.Now(), "Stage clean generated a modification")
		}
	}
	return
}
