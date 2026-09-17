// generated code - do not edit
package models

import "time"

// CleanSlice is the Stage method that removes unstaged elements from a slice of pointers.
func (stage *Stage) CleanSlice[T GongstructPtr](slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if stage.IsStaged(element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	if modified {
		*slice = cleanedSlice
	}
	return
}

// CleanPointer is the Stage method that sets the pointer to nil if the referenced element is not staged.
func (stage *Stage) CleanPointer[T GongstructPtr](element *T) (modified bool) {
	var zero T
	if *element == zero {
		return
	}

	if !stage.IsStaged(*element) {
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
	modified = stage.CleanPointer(&allocatedresourceshape.Part) || modified
	modified = stage.CleanPointer(&allocatedresourceshape.Resource) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by AllocatedSystemShape
func (allocatedsystemshape *AllocatedSystemShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&allocatedsystemshape.Part) || modified
	modified = stage.CleanPointer(&allocatedsystemshape.System) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ControlFlow
func (controlflow *ControlFlow) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&controlflow.Start) || modified
	modified = stage.CleanPointer(&controlflow.End) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ControlFlowShape
func (controlflowshape *ControlFlowShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&controlflowshape.ControlFlow) || modified
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
	modified = stage.CleanSlice(&dataflow.Datas) || modified
	// insertion point per field
	modified = stage.CleanPointer(&dataflow.StartPort) || modified
	modified = stage.CleanPointer(&dataflow.EndPort) || modified
	modified = stage.CleanPointer(&dataflow.StartExternalPart) || modified
	modified = stage.CleanPointer(&dataflow.EndExternalPart) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DataFlowShape
func (dataflowshape *DataFlowShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&dataflowshape.DataFlow) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DataShape
func (datashape *DataShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&datashape.Data) || modified
	modified = stage.CleanPointer(&datashape.DataFlow) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DiagramLayerState
func (diagramlayerstate *DiagramLayerState) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&diagramlayerstate.DiagramStructure) || modified
	modified = stage.CleanPointer(&diagramlayerstate.LayerDefinition) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DiagramStructure
func (diagramstructure *DiagramStructure) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&diagramstructure.System_Shapes) || modified
	modified = stage.CleanSlice(&diagramstructure.SystemsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramstructure.Part_Shapes) || modified
	modified = stage.CleanSlice(&diagramstructure.PartWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramstructure.ExternalPart_Shapes) || modified
	modified = stage.CleanSlice(&diagramstructure.ExternalPartWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramstructure.ExternalPartsWhoseOutDataFlowsNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramstructure.ExternalPartsWhoseInDataFlowsNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramstructure.PortsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramstructure.Port_Shapes) || modified
	modified = stage.CleanSlice(&diagramstructure.ControlFlowsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramstructure.ControlFlow_Shapes) || modified
	modified = stage.CleanSlice(&diagramstructure.DataFlowsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramstructure.DataFlow_Shapes) || modified
	modified = stage.CleanSlice(&diagramstructure.DatasWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramstructure.Data_Shapes) || modified
	modified = stage.CleanSlice(&diagramstructure.DataFlowsWhoseDataNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramstructure.AllocatedResourcesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramstructure.AllocatedResourceShapes) || modified
	modified = stage.CleanSlice(&diagramstructure.AllocatedSystemesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramstructure.AllocatedSystemShapes) || modified
	modified = stage.CleanSlice(&diagramstructure.Note_Shapes) || modified
	modified = stage.CleanSlice(&diagramstructure.NotesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&diagramstructure.NotePortShapes) || modified
	modified = stage.CleanSlice(&diagramstructure.NotePartShapes) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ExternalPartShape
func (externalpartshape *ExternalPartShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&externalpartshape.Part) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by LayerDefinition
func (layerdefinition *LayerDefinition) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&layerdefinition.Query) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Library
func (library *Library) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&library.SubLibraries) || modified
	modified = stage.CleanSlice(&library.SubLibrariesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.RootSystemes) || modified
	modified = stage.CleanSlice(&library.SystemsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.RootDataFlows) || modified
	modified = stage.CleanSlice(&library.DataFlowsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.RootDatas) || modified
	modified = stage.CleanSlice(&library.DatasWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.RootResources) || modified
	modified = stage.CleanSlice(&library.ResourcesWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.PartsWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&library.RootNotes) || modified
	modified = stage.CleanSlice(&library.NotesWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by Note
func (note *Note) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&note.Parts) || modified
	modified = stage.CleanSlice(&note.Ports) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by NotePartShape
func (notepartshape *NotePartShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&notepartshape.Note) || modified
	modified = stage.CleanPointer(&notepartshape.Part) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NotePortShape
func (noteportshape *NotePortShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&noteportshape.Note) || modified
	modified = stage.CleanPointer(&noteportshape.Port) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by NoteShape
func (noteshape *NoteShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&noteshape.Note) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Part
func (part *Part) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&part.Ports) || modified
	modified = stage.CleanSlice(&part.ControlFlows) || modified
	modified = stage.CleanSlice(&part.PortWhoseOutControlFlowsNodeIsExpanded) || modified
	modified = stage.CleanSlice(&part.PortWhoseInControlFlowsNodeIsExpanded) || modified
	modified = stage.CleanSlice(&part.PortWhoseOutDataFlowsNodeIsExpanded) || modified
	modified = stage.CleanSlice(&part.PortWhoseInDataFlowsNodeIsExpanded) || modified
	modified = stage.CleanSlice(&part.PartAnchoredPath) || modified
	// insertion point per field
	modified = stage.CleanPointer(&part.TypeOfPart) || modified
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
	modified = stage.CleanPointer(&partshape.Part) || modified
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
	modified = stage.CleanPointer(&portshape.Port) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by Resource
func (resource *Resource) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SemanticTag
func (semantictag *SemanticTag) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&semantictag.Parts) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by System
func (system *System) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = stage.CleanSlice(&system.DiagramStructures) || modified
	modified = stage.CleanSlice(&system.DiagramStructureWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&system.SubSystemes) || modified
	modified = stage.CleanSlice(&system.Parts) || modified
	modified = stage.CleanSlice(&system.PartWhoseNodeIsExpanded) || modified
	modified = stage.CleanSlice(&system.DataFlows) || modified
	modified = stage.CleanSlice(&system.ExternalParts) || modified
	modified = stage.CleanSlice(&system.ExternalPartWhoseNodeIsExpanded) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SystemShape
func (systemshape *SystemShape) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = stage.CleanPointer(&systemshape.System) || modified
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
