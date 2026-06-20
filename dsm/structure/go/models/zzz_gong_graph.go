// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *DiagramStructure:
		ok = stage.IsStagedDiagramStructure(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Link:
		ok = stage.IsStagedLink(target)

	case *LinkAssociationShape:
		ok = stage.IsStagedLinkAssociationShape(target)

	case *Part:
		ok = stage.IsStagedPart(target)

	case *PartShape:
		ok = stage.IsStagedPartShape(target)

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
	case *DiagramStructure:
		ok = stage.IsStagedDiagramStructure(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Link:
		ok = stage.IsStagedLink(target)

	case *LinkAssociationShape:
		ok = stage.IsStagedLinkAssociationShape(target)

	case *Part:
		ok = stage.IsStagedPart(target)

	case *PartShape:
		ok = stage.IsStagedPartShape(target)

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
func (stage *Stage) IsStagedDiagramStructure(diagramstructure *DiagramStructure) (ok bool) {

	_, ok = stage.DiagramStructures[diagramstructure]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedLink(link *Link) (ok bool) {

	_, ok = stage.Links[link]

	return
}

func (stage *Stage) IsStagedLinkAssociationShape(linkassociationshape *LinkAssociationShape) (ok bool) {

	_, ok = stage.LinkAssociationShapes[linkassociationshape]

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
	case *DiagramStructure:
		stage.StageBranchDiagramStructure(target)

	case *Library:
		stage.StageBranchLibrary(target)

	case *Link:
		stage.StageBranchLink(target)

	case *LinkAssociationShape:
		stage.StageBranchLinkAssociationShape(target)

	case *Part:
		stage.StageBranchPart(target)

	case *PartShape:
		stage.StageBranchPartShape(target)

	case *System:
		stage.StageBranchSystem(target)

	case *SystemShape:
		stage.StageBranchSystemShape(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
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
	for _, _partshape := range diagramstructure.Part_Shapes {
		StageBranch(stage, _partshape)
	}
	for _, _part := range diagramstructure.PartsWhoseNodeIsExpanded {
		StageBranch(stage, _part)
	}
	for _, _linkassociationshape := range diagramstructure.Link_Shapes {
		StageBranch(stage, _linkassociationshape)
	}
	for _, _link := range diagramstructure.LinksWhoseNodeIsExpanded {
		StageBranch(stage, _link)
	}

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
	for _, _system := range library.RootSystems {
		StageBranch(stage, _system)
	}
	for _, _system := range library.SystemsWhoseNodeIsExpanded {
		StageBranch(stage, _system)
	}

}

func (stage *Stage) StageBranchLink(link *Link) {

	// check if instance is already staged
	if IsStaged(stage, link) {
		return
	}

	link.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if link.Source != nil {
		StageBranch(stage, link.Source)
	}
	if link.Target != nil {
		StageBranch(stage, link.Target)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchLinkAssociationShape(linkassociationshape *LinkAssociationShape) {

	// check if instance is already staged
	if IsStaged(stage, linkassociationshape) {
		return
	}

	linkassociationshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if linkassociationshape.Link != nil {
		StageBranch(stage, linkassociationshape.Link)
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

func (stage *Stage) StageBranchSystem(system *System) {

	// check if instance is already staged
	if IsStaged(stage, system) {
		return
	}

	system.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _part := range system.Parts {
		StageBranch(stage, _part)
	}
	for _, _part := range system.PartsWhoseNodeIsExpanded {
		StageBranch(stage, _part)
	}
	for _, _system := range system.SubSystems {
		StageBranch(stage, _system)
	}
	for _, _system := range system.SubSystemsWhoseNodeIsExpanded {
		StageBranch(stage, _system)
	}
	for _, _link := range system.Links {
		StageBranch(stage, _link)
	}
	for _, _link := range system.LinksWhoseNodeIsExpanded {
		StageBranch(stage, _link)
	}
	for _, _diagramstructure := range system.DiagramStructures {
		StageBranch(stage, _diagramstructure)
	}
	for _, _diagramstructure := range system.DiagramStructuresWhoseNodeIsExpanded {
		StageBranch(stage, _diagramstructure)
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
	case *DiagramStructure:
		toT := CopyBranchDiagramStructure(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := CopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Link:
		toT := CopyBranchLink(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *LinkAssociationShape:
		toT := CopyBranchLinkAssociationShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Part:
		toT := CopyBranchPart(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PartShape:
		toT := CopyBranchPartShape(mapOrigCopy, fromT)
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
	for _, _partshape := range diagramstructureFrom.Part_Shapes {
		diagramstructureTo.Part_Shapes = append(diagramstructureTo.Part_Shapes, CopyBranchPartShape(mapOrigCopy, _partshape))
	}
	for _, _part := range diagramstructureFrom.PartsWhoseNodeIsExpanded {
		diagramstructureTo.PartsWhoseNodeIsExpanded = append(diagramstructureTo.PartsWhoseNodeIsExpanded, CopyBranchPart(mapOrigCopy, _part))
	}
	for _, _linkassociationshape := range diagramstructureFrom.Link_Shapes {
		diagramstructureTo.Link_Shapes = append(diagramstructureTo.Link_Shapes, CopyBranchLinkAssociationShape(mapOrigCopy, _linkassociationshape))
	}
	for _, _link := range diagramstructureFrom.LinksWhoseNodeIsExpanded {
		diagramstructureTo.LinksWhoseNodeIsExpanded = append(diagramstructureTo.LinksWhoseNodeIsExpanded, CopyBranchLink(mapOrigCopy, _link))
	}

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
	for _, _system := range libraryFrom.RootSystems {
		libraryTo.RootSystems = append(libraryTo.RootSystems, CopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _system := range libraryFrom.SystemsWhoseNodeIsExpanded {
		libraryTo.SystemsWhoseNodeIsExpanded = append(libraryTo.SystemsWhoseNodeIsExpanded, CopyBranchSystem(mapOrigCopy, _system))
	}

	return
}

func CopyBranchLink(mapOrigCopy map[any]any, linkFrom *Link) (linkTo *Link) {

	// linkFrom has already been copied
	if _linkTo, ok := mapOrigCopy[linkFrom]; ok {
		linkTo = _linkTo.(*Link)
		return
	}

	linkTo = new(Link)
	mapOrigCopy[linkFrom] = linkTo
	linkFrom.CopyBasicFields(linkTo)

	//insertion point for the staging of instances referenced by pointers
	if linkFrom.Source != nil {
		linkTo.Source = CopyBranchPart(mapOrigCopy, linkFrom.Source)
	}
	if linkFrom.Target != nil {
		linkTo.Target = CopyBranchPart(mapOrigCopy, linkFrom.Target)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchLinkAssociationShape(mapOrigCopy map[any]any, linkassociationshapeFrom *LinkAssociationShape) (linkassociationshapeTo *LinkAssociationShape) {

	// linkassociationshapeFrom has already been copied
	if _linkassociationshapeTo, ok := mapOrigCopy[linkassociationshapeFrom]; ok {
		linkassociationshapeTo = _linkassociationshapeTo.(*LinkAssociationShape)
		return
	}

	linkassociationshapeTo = new(LinkAssociationShape)
	mapOrigCopy[linkassociationshapeFrom] = linkassociationshapeTo
	linkassociationshapeFrom.CopyBasicFields(linkassociationshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if linkassociationshapeFrom.Link != nil {
		linkassociationshapeTo.Link = CopyBranchLink(mapOrigCopy, linkassociationshapeFrom.Link)
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
	for _, _part := range systemFrom.Parts {
		systemTo.Parts = append(systemTo.Parts, CopyBranchPart(mapOrigCopy, _part))
	}
	for _, _part := range systemFrom.PartsWhoseNodeIsExpanded {
		systemTo.PartsWhoseNodeIsExpanded = append(systemTo.PartsWhoseNodeIsExpanded, CopyBranchPart(mapOrigCopy, _part))
	}
	for _, _system := range systemFrom.SubSystems {
		systemTo.SubSystems = append(systemTo.SubSystems, CopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _system := range systemFrom.SubSystemsWhoseNodeIsExpanded {
		systemTo.SubSystemsWhoseNodeIsExpanded = append(systemTo.SubSystemsWhoseNodeIsExpanded, CopyBranchSystem(mapOrigCopy, _system))
	}
	for _, _link := range systemFrom.Links {
		systemTo.Links = append(systemTo.Links, CopyBranchLink(mapOrigCopy, _link))
	}
	for _, _link := range systemFrom.LinksWhoseNodeIsExpanded {
		systemTo.LinksWhoseNodeIsExpanded = append(systemTo.LinksWhoseNodeIsExpanded, CopyBranchLink(mapOrigCopy, _link))
	}
	for _, _diagramstructure := range systemFrom.DiagramStructures {
		systemTo.DiagramStructures = append(systemTo.DiagramStructures, CopyBranchDiagramStructure(mapOrigCopy, _diagramstructure))
	}
	for _, _diagramstructure := range systemFrom.DiagramStructuresWhoseNodeIsExpanded {
		systemTo.DiagramStructuresWhoseNodeIsExpanded = append(systemTo.DiagramStructuresWhoseNodeIsExpanded, CopyBranchDiagramStructure(mapOrigCopy, _diagramstructure))
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
	case *DiagramStructure:
		stage.UnstageBranchDiagramStructure(target)

	case *Library:
		stage.UnstageBranchLibrary(target)

	case *Link:
		stage.UnstageBranchLink(target)

	case *LinkAssociationShape:
		stage.UnstageBranchLinkAssociationShape(target)

	case *Part:
		stage.UnstageBranchPart(target)

	case *PartShape:
		stage.UnstageBranchPartShape(target)

	case *System:
		stage.UnstageBranchSystem(target)

	case *SystemShape:
		stage.UnstageBranchSystemShape(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
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
	for _, _partshape := range diagramstructure.Part_Shapes {
		UnstageBranch(stage, _partshape)
	}
	for _, _part := range diagramstructure.PartsWhoseNodeIsExpanded {
		UnstageBranch(stage, _part)
	}
	for _, _linkassociationshape := range diagramstructure.Link_Shapes {
		UnstageBranch(stage, _linkassociationshape)
	}
	for _, _link := range diagramstructure.LinksWhoseNodeIsExpanded {
		UnstageBranch(stage, _link)
	}

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
	for _, _system := range library.RootSystems {
		UnstageBranch(stage, _system)
	}
	for _, _system := range library.SystemsWhoseNodeIsExpanded {
		UnstageBranch(stage, _system)
	}

}

func (stage *Stage) UnstageBranchLink(link *Link) {

	// check if instance is already staged
	if !IsStaged(stage, link) {
		return
	}

	link.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if link.Source != nil {
		UnstageBranch(stage, link.Source)
	}
	if link.Target != nil {
		UnstageBranch(stage, link.Target)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchLinkAssociationShape(linkassociationshape *LinkAssociationShape) {

	// check if instance is already staged
	if !IsStaged(stage, linkassociationshape) {
		return
	}

	linkassociationshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if linkassociationshape.Link != nil {
		UnstageBranch(stage, linkassociationshape.Link)
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

func (stage *Stage) UnstageBranchSystem(system *System) {

	// check if instance is already staged
	if !IsStaged(stage, system) {
		return
	}

	system.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _part := range system.Parts {
		UnstageBranch(stage, _part)
	}
	for _, _part := range system.PartsWhoseNodeIsExpanded {
		UnstageBranch(stage, _part)
	}
	for _, _system := range system.SubSystems {
		UnstageBranch(stage, _system)
	}
	for _, _system := range system.SubSystemsWhoseNodeIsExpanded {
		UnstageBranch(stage, _system)
	}
	for _, _link := range system.Links {
		UnstageBranch(stage, _link)
	}
	for _, _link := range system.LinksWhoseNodeIsExpanded {
		UnstageBranch(stage, _link)
	}
	for _, _diagramstructure := range system.DiagramStructures {
		UnstageBranch(stage, _diagramstructure)
	}
	for _, _diagramstructure := range system.DiagramStructuresWhoseNodeIsExpanded {
		UnstageBranch(stage, _diagramstructure)
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
func (reference *DiagramStructure) GongReconstructPointersFromReferences(stage *Stage, instance *DiagramStructure) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.System_Shapes = reference.System_Shapes[:0]
	for _, _b := range instance.System_Shapes {
		reference.System_Shapes = append(reference.System_Shapes, stage.SystemShapes_reference[_b])
	}
	reference.Part_Shapes = reference.Part_Shapes[:0]
	for _, _b := range instance.Part_Shapes {
		reference.Part_Shapes = append(reference.Part_Shapes, stage.PartShapes_reference[_b])
	}
	reference.PartsWhoseNodeIsExpanded = reference.PartsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.PartsWhoseNodeIsExpanded {
		reference.PartsWhoseNodeIsExpanded = append(reference.PartsWhoseNodeIsExpanded, stage.Parts_reference[_b])
	}
	reference.Link_Shapes = reference.Link_Shapes[:0]
	for _, _b := range instance.Link_Shapes {
		reference.Link_Shapes = append(reference.Link_Shapes, stage.LinkAssociationShapes_reference[_b])
	}
	reference.LinksWhoseNodeIsExpanded = reference.LinksWhoseNodeIsExpanded[:0]
	for _, _b := range instance.LinksWhoseNodeIsExpanded {
		reference.LinksWhoseNodeIsExpanded = append(reference.LinksWhoseNodeIsExpanded, stage.Links_reference[_b])
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
	reference.RootSystems = reference.RootSystems[:0]
	for _, _b := range instance.RootSystems {
		reference.RootSystems = append(reference.RootSystems, stage.Systems_reference[_b])
	}
	reference.SystemsWhoseNodeIsExpanded = reference.SystemsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.SystemsWhoseNodeIsExpanded {
		reference.SystemsWhoseNodeIsExpanded = append(reference.SystemsWhoseNodeIsExpanded, stage.Systems_reference[_b])
	}
}

func (reference *Link) GongReconstructPointersFromReferences(stage *Stage, instance *Link) {
	// insertion point for pointers field
	if instance.Source != nil {
		reference.Source = stage.Parts_reference[instance.Source]
	}
	if instance.Target != nil {
		reference.Target = stage.Parts_reference[instance.Target]
	}
	// insertion point for slice of pointers field
}

func (reference *LinkAssociationShape) GongReconstructPointersFromReferences(stage *Stage, instance *LinkAssociationShape) {
	// insertion point for pointers field
	if instance.Link != nil {
		reference.Link = stage.Links_reference[instance.Link]
	}
	// insertion point for slice of pointers field
}

func (reference *Part) GongReconstructPointersFromReferences(stage *Stage, instance *Part) {
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

func (reference *System) GongReconstructPointersFromReferences(stage *Stage, instance *System) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Parts = reference.Parts[:0]
	for _, _b := range instance.Parts {
		reference.Parts = append(reference.Parts, stage.Parts_reference[_b])
	}
	reference.PartsWhoseNodeIsExpanded = reference.PartsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.PartsWhoseNodeIsExpanded {
		reference.PartsWhoseNodeIsExpanded = append(reference.PartsWhoseNodeIsExpanded, stage.Parts_reference[_b])
	}
	reference.SubSystems = reference.SubSystems[:0]
	for _, _b := range instance.SubSystems {
		reference.SubSystems = append(reference.SubSystems, stage.Systems_reference[_b])
	}
	reference.SubSystemsWhoseNodeIsExpanded = reference.SubSystemsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.SubSystemsWhoseNodeIsExpanded {
		reference.SubSystemsWhoseNodeIsExpanded = append(reference.SubSystemsWhoseNodeIsExpanded, stage.Systems_reference[_b])
	}
	reference.Links = reference.Links[:0]
	for _, _b := range instance.Links {
		reference.Links = append(reference.Links, stage.Links_reference[_b])
	}
	reference.LinksWhoseNodeIsExpanded = reference.LinksWhoseNodeIsExpanded[:0]
	for _, _b := range instance.LinksWhoseNodeIsExpanded {
		reference.LinksWhoseNodeIsExpanded = append(reference.LinksWhoseNodeIsExpanded, stage.Links_reference[_b])
	}
	reference.DiagramStructures = reference.DiagramStructures[:0]
	for _, _b := range instance.DiagramStructures {
		reference.DiagramStructures = append(reference.DiagramStructures, stage.DiagramStructures_reference[_b])
	}
	reference.DiagramStructuresWhoseNodeIsExpanded = reference.DiagramStructuresWhoseNodeIsExpanded[:0]
	for _, _b := range instance.DiagramStructuresWhoseNodeIsExpanded {
		reference.DiagramStructuresWhoseNodeIsExpanded = append(reference.DiagramStructuresWhoseNodeIsExpanded, stage.DiagramStructures_reference[_b])
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
	var _Part_Shapes []*PartShape
	for _, _reference := range reference.Part_Shapes {
		if _instance, ok := stage.PartShapes_instance[_reference]; ok {
			_Part_Shapes = append(_Part_Shapes, _instance)
		}
	}
	reference.Part_Shapes = _Part_Shapes
	var _PartsWhoseNodeIsExpanded []*Part
	for _, _reference := range reference.PartsWhoseNodeIsExpanded {
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			_PartsWhoseNodeIsExpanded = append(_PartsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.PartsWhoseNodeIsExpanded = _PartsWhoseNodeIsExpanded
	var _Link_Shapes []*LinkAssociationShape
	for _, _reference := range reference.Link_Shapes {
		if _instance, ok := stage.LinkAssociationShapes_instance[_reference]; ok {
			_Link_Shapes = append(_Link_Shapes, _instance)
		}
	}
	reference.Link_Shapes = _Link_Shapes
	var _LinksWhoseNodeIsExpanded []*Link
	for _, _reference := range reference.LinksWhoseNodeIsExpanded {
		if _instance, ok := stage.Links_instance[_reference]; ok {
			_LinksWhoseNodeIsExpanded = append(_LinksWhoseNodeIsExpanded, _instance)
		}
	}
	reference.LinksWhoseNodeIsExpanded = _LinksWhoseNodeIsExpanded
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
	var _RootSystems []*System
	for _, _reference := range reference.RootSystems {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_RootSystems = append(_RootSystems, _instance)
		}
	}
	reference.RootSystems = _RootSystems
	var _SystemsWhoseNodeIsExpanded []*System
	for _, _reference := range reference.SystemsWhoseNodeIsExpanded {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_SystemsWhoseNodeIsExpanded = append(_SystemsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.SystemsWhoseNodeIsExpanded = _SystemsWhoseNodeIsExpanded
}

func (reference *Link) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Source; _reference != nil {
		reference.Source = nil
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			reference.Source = _instance
		}
	}
	if _reference := reference.Target; _reference != nil {
		reference.Target = nil
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			reference.Target = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *LinkAssociationShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.Link; _reference != nil {
		reference.Link = nil
		if _instance, ok := stage.Links_instance[_reference]; ok {
			reference.Link = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *Part) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *System) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Parts []*Part
	for _, _reference := range reference.Parts {
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			_Parts = append(_Parts, _instance)
		}
	}
	reference.Parts = _Parts
	var _PartsWhoseNodeIsExpanded []*Part
	for _, _reference := range reference.PartsWhoseNodeIsExpanded {
		if _instance, ok := stage.Parts_instance[_reference]; ok {
			_PartsWhoseNodeIsExpanded = append(_PartsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.PartsWhoseNodeIsExpanded = _PartsWhoseNodeIsExpanded
	var _SubSystems []*System
	for _, _reference := range reference.SubSystems {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_SubSystems = append(_SubSystems, _instance)
		}
	}
	reference.SubSystems = _SubSystems
	var _SubSystemsWhoseNodeIsExpanded []*System
	for _, _reference := range reference.SubSystemsWhoseNodeIsExpanded {
		if _instance, ok := stage.Systems_instance[_reference]; ok {
			_SubSystemsWhoseNodeIsExpanded = append(_SubSystemsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.SubSystemsWhoseNodeIsExpanded = _SubSystemsWhoseNodeIsExpanded
	var _Links []*Link
	for _, _reference := range reference.Links {
		if _instance, ok := stage.Links_instance[_reference]; ok {
			_Links = append(_Links, _instance)
		}
	}
	reference.Links = _Links
	var _LinksWhoseNodeIsExpanded []*Link
	for _, _reference := range reference.LinksWhoseNodeIsExpanded {
		if _instance, ok := stage.Links_instance[_reference]; ok {
			_LinksWhoseNodeIsExpanded = append(_LinksWhoseNodeIsExpanded, _instance)
		}
	}
	reference.LinksWhoseNodeIsExpanded = _LinksWhoseNodeIsExpanded
	var _DiagramStructures []*DiagramStructure
	for _, _reference := range reference.DiagramStructures {
		if _instance, ok := stage.DiagramStructures_instance[_reference]; ok {
			_DiagramStructures = append(_DiagramStructures, _instance)
		}
	}
	reference.DiagramStructures = _DiagramStructures
	var _DiagramStructuresWhoseNodeIsExpanded []*DiagramStructure
	for _, _reference := range reference.DiagramStructuresWhoseNodeIsExpanded {
		if _instance, ok := stage.DiagramStructures_instance[_reference]; ok {
			_DiagramStructuresWhoseNodeIsExpanded = append(_DiagramStructuresWhoseNodeIsExpanded, _instance)
		}
	}
	reference.DiagramStructuresWhoseNodeIsExpanded = _DiagramStructuresWhoseNodeIsExpanded
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
func (diagramstructure *DiagramStructure) GongDiff(stage *Stage, diagramstructureOther *DiagramStructure) (diffs []string) {
	// insertion point for field diffs
	if diagramstructure.Name != diagramstructureOther.Name {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "Name"))
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
	if diagramstructure.Width != diagramstructureOther.Width {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "Width"))
	}
	if diagramstructure.Height != diagramstructureOther.Height {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "Height"))
	}
	if diagramstructure.DefaultBoxWidth != diagramstructureOther.DefaultBoxWidth {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "DefaultBoxWidth"))
	}
	if diagramstructure.DefaultBoxHeigth != diagramstructureOther.DefaultBoxHeigth {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "DefaultBoxHeigth"))
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
	PartsWhoseNodeIsExpandedDifferent := false
	if len(diagramstructure.PartsWhoseNodeIsExpanded) != len(diagramstructureOther.PartsWhoseNodeIsExpanded) {
		PartsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramstructure.PartsWhoseNodeIsExpanded {
			if (diagramstructure.PartsWhoseNodeIsExpanded[i] == nil) != (diagramstructureOther.PartsWhoseNodeIsExpanded[i] == nil) {
				PartsWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramstructure.PartsWhoseNodeIsExpanded[i] != nil && diagramstructureOther.PartsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.PartsWhoseNodeIsExpanded[i] != diagramstructureOther.PartsWhoseNodeIsExpanded[i] {
					PartsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if PartsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagramstructure, diagramstructureOther, "PartsWhoseNodeIsExpanded", diagramstructureOther.PartsWhoseNodeIsExpanded, diagramstructure.PartsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	Link_ShapesDifferent := false
	if len(diagramstructure.Link_Shapes) != len(diagramstructureOther.Link_Shapes) {
		Link_ShapesDifferent = true
	} else {
		for i := range diagramstructure.Link_Shapes {
			if (diagramstructure.Link_Shapes[i] == nil) != (diagramstructureOther.Link_Shapes[i] == nil) {
				Link_ShapesDifferent = true
				break
			} else if diagramstructure.Link_Shapes[i] != nil && diagramstructureOther.Link_Shapes[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.Link_Shapes[i] != diagramstructureOther.Link_Shapes[i] {
					Link_ShapesDifferent = true
					break
				}
			}
		}
	}
	if Link_ShapesDifferent {
		ops := Diff(stage, diagramstructure, diagramstructureOther, "Link_Shapes", diagramstructureOther.Link_Shapes, diagramstructure.Link_Shapes)
		diffs = append(diffs, ops)
	}
	if diagramstructure.IsLinksNodeExpanded != diagramstructureOther.IsLinksNodeExpanded {
		diffs = append(diffs, diagramstructure.GongMarshallField(stage, "IsLinksNodeExpanded"))
	}
	LinksWhoseNodeIsExpandedDifferent := false
	if len(diagramstructure.LinksWhoseNodeIsExpanded) != len(diagramstructureOther.LinksWhoseNodeIsExpanded) {
		LinksWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range diagramstructure.LinksWhoseNodeIsExpanded {
			if (diagramstructure.LinksWhoseNodeIsExpanded[i] == nil) != (diagramstructureOther.LinksWhoseNodeIsExpanded[i] == nil) {
				LinksWhoseNodeIsExpandedDifferent = true
				break
			} else if diagramstructure.LinksWhoseNodeIsExpanded[i] != nil && diagramstructureOther.LinksWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if diagramstructure.LinksWhoseNodeIsExpanded[i] != diagramstructureOther.LinksWhoseNodeIsExpanded[i] {
					LinksWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if LinksWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, diagramstructure, diagramstructureOther, "LinksWhoseNodeIsExpanded", diagramstructureOther.LinksWhoseNodeIsExpanded, diagramstructure.LinksWhoseNodeIsExpanded)
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
	RootSystemsDifferent := false
	if len(library.RootSystems) != len(libraryOther.RootSystems) {
		RootSystemsDifferent = true
	} else {
		for i := range library.RootSystems {
			if (library.RootSystems[i] == nil) != (libraryOther.RootSystems[i] == nil) {
				RootSystemsDifferent = true
				break
			} else if library.RootSystems[i] != nil && libraryOther.RootSystems[i] != nil {
				// this is a pointer comparaison
				if library.RootSystems[i] != libraryOther.RootSystems[i] {
					RootSystemsDifferent = true
					break
				}
			}
		}
	}
	if RootSystemsDifferent {
		ops := Diff(stage, library, libraryOther, "RootSystems", libraryOther.RootSystems, library.RootSystems)
		diffs = append(diffs, ops)
	}
	if library.IsSystemsNodeExpanded != libraryOther.IsSystemsNodeExpanded {
		diffs = append(diffs, library.GongMarshallField(stage, "IsSystemsNodeExpanded"))
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

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (link *Link) GongDiff(stage *Stage, linkOther *Link) (diffs []string) {
	// insertion point for field diffs
	if link.Name != linkOther.Name {
		diffs = append(diffs, link.GongMarshallField(stage, "Name"))
	}
	if link.ComputedPrefix != linkOther.ComputedPrefix {
		diffs = append(diffs, link.GongMarshallField(stage, "ComputedPrefix"))
	}
	if link.IsExpanded != linkOther.IsExpanded {
		diffs = append(diffs, link.GongMarshallField(stage, "IsExpanded"))
	}
	if link.LayoutDirection != linkOther.LayoutDirection {
		diffs = append(diffs, link.GongMarshallField(stage, "LayoutDirection"))
	}
	if (link.Source == nil) != (linkOther.Source == nil) {
		diffs = append(diffs, link.GongMarshallField(stage, "Source"))
	} else if link.Source != nil && linkOther.Source != nil {
		if link.Source != linkOther.Source {
			diffs = append(diffs, link.GongMarshallField(stage, "Source"))
		}
	}
	if (link.Target == nil) != (linkOther.Target == nil) {
		diffs = append(diffs, link.GongMarshallField(stage, "Target"))
	} else if link.Target != nil && linkOther.Target != nil {
		if link.Target != linkOther.Target {
			diffs = append(diffs, link.GongMarshallField(stage, "Target"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (linkassociationshape *LinkAssociationShape) GongDiff(stage *Stage, linkassociationshapeOther *LinkAssociationShape) (diffs []string) {
	// insertion point for field diffs
	if linkassociationshape.Name != linkassociationshapeOther.Name {
		diffs = append(diffs, linkassociationshape.GongMarshallField(stage, "Name"))
	}
	if (linkassociationshape.Link == nil) != (linkassociationshapeOther.Link == nil) {
		diffs = append(diffs, linkassociationshape.GongMarshallField(stage, "Link"))
	} else if linkassociationshape.Link != nil && linkassociationshapeOther.Link != nil {
		if linkassociationshape.Link != linkassociationshapeOther.Link {
			diffs = append(diffs, linkassociationshape.GongMarshallField(stage, "Link"))
		}
	}
	if linkassociationshape.StartRatio != linkassociationshapeOther.StartRatio {
		diffs = append(diffs, linkassociationshape.GongMarshallField(stage, "StartRatio"))
	}
	if linkassociationshape.EndRatio != linkassociationshapeOther.EndRatio {
		diffs = append(diffs, linkassociationshape.GongMarshallField(stage, "EndRatio"))
	}
	if linkassociationshape.StartOrientation != linkassociationshapeOther.StartOrientation {
		diffs = append(diffs, linkassociationshape.GongMarshallField(stage, "StartOrientation"))
	}
	if linkassociationshape.EndOrientation != linkassociationshapeOther.EndOrientation {
		diffs = append(diffs, linkassociationshape.GongMarshallField(stage, "EndOrientation"))
	}
	if linkassociationshape.CornerOffsetRatio != linkassociationshapeOther.CornerOffsetRatio {
		diffs = append(diffs, linkassociationshape.GongMarshallField(stage, "CornerOffsetRatio"))
	}
	if linkassociationshape.IsHidden != linkassociationshapeOther.IsHidden {
		diffs = append(diffs, linkassociationshape.GongMarshallField(stage, "IsHidden"))
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
	if part.ComputedPrefix != partOther.ComputedPrefix {
		diffs = append(diffs, part.GongMarshallField(stage, "ComputedPrefix"))
	}
	if part.IsExpanded != partOther.IsExpanded {
		diffs = append(diffs, part.GongMarshallField(stage, "IsExpanded"))
	}
	if part.LayoutDirection != partOther.LayoutDirection {
		diffs = append(diffs, part.GongMarshallField(stage, "LayoutDirection"))
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
	if partshape.OverideLayoutDirection != partshapeOther.OverideLayoutDirection {
		diffs = append(diffs, partshape.GongMarshallField(stage, "OverideLayoutDirection"))
	}
	if partshape.LayoutDirection != partshapeOther.LayoutDirection {
		diffs = append(diffs, partshape.GongMarshallField(stage, "LayoutDirection"))
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
	if system.ComputedPrefix != systemOther.ComputedPrefix {
		diffs = append(diffs, system.GongMarshallField(stage, "ComputedPrefix"))
	}
	if system.IsExpanded != systemOther.IsExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsExpanded"))
	}
	if system.LayoutDirection != systemOther.LayoutDirection {
		diffs = append(diffs, system.GongMarshallField(stage, "LayoutDirection"))
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
	if system.IsPartsNodeExpanded != systemOther.IsPartsNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsPartsNodeExpanded"))
	}
	PartsWhoseNodeIsExpandedDifferent := false
	if len(system.PartsWhoseNodeIsExpanded) != len(systemOther.PartsWhoseNodeIsExpanded) {
		PartsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range system.PartsWhoseNodeIsExpanded {
			if (system.PartsWhoseNodeIsExpanded[i] == nil) != (systemOther.PartsWhoseNodeIsExpanded[i] == nil) {
				PartsWhoseNodeIsExpandedDifferent = true
				break
			} else if system.PartsWhoseNodeIsExpanded[i] != nil && systemOther.PartsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if system.PartsWhoseNodeIsExpanded[i] != systemOther.PartsWhoseNodeIsExpanded[i] {
					PartsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if PartsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, system, systemOther, "PartsWhoseNodeIsExpanded", systemOther.PartsWhoseNodeIsExpanded, system.PartsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	SubSystemsDifferent := false
	if len(system.SubSystems) != len(systemOther.SubSystems) {
		SubSystemsDifferent = true
	} else {
		for i := range system.SubSystems {
			if (system.SubSystems[i] == nil) != (systemOther.SubSystems[i] == nil) {
				SubSystemsDifferent = true
				break
			} else if system.SubSystems[i] != nil && systemOther.SubSystems[i] != nil {
				// this is a pointer comparaison
				if system.SubSystems[i] != systemOther.SubSystems[i] {
					SubSystemsDifferent = true
					break
				}
			}
		}
	}
	if SubSystemsDifferent {
		ops := Diff(stage, system, systemOther, "SubSystems", systemOther.SubSystems, system.SubSystems)
		diffs = append(diffs, ops)
	}
	if system.IsSubSystemsNodeExpanded != systemOther.IsSubSystemsNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsSubSystemsNodeExpanded"))
	}
	SubSystemsWhoseNodeIsExpandedDifferent := false
	if len(system.SubSystemsWhoseNodeIsExpanded) != len(systemOther.SubSystemsWhoseNodeIsExpanded) {
		SubSystemsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range system.SubSystemsWhoseNodeIsExpanded {
			if (system.SubSystemsWhoseNodeIsExpanded[i] == nil) != (systemOther.SubSystemsWhoseNodeIsExpanded[i] == nil) {
				SubSystemsWhoseNodeIsExpandedDifferent = true
				break
			} else if system.SubSystemsWhoseNodeIsExpanded[i] != nil && systemOther.SubSystemsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if system.SubSystemsWhoseNodeIsExpanded[i] != systemOther.SubSystemsWhoseNodeIsExpanded[i] {
					SubSystemsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if SubSystemsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, system, systemOther, "SubSystemsWhoseNodeIsExpanded", systemOther.SubSystemsWhoseNodeIsExpanded, system.SubSystemsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	LinksDifferent := false
	if len(system.Links) != len(systemOther.Links) {
		LinksDifferent = true
	} else {
		for i := range system.Links {
			if (system.Links[i] == nil) != (systemOther.Links[i] == nil) {
				LinksDifferent = true
				break
			} else if system.Links[i] != nil && systemOther.Links[i] != nil {
				// this is a pointer comparaison
				if system.Links[i] != systemOther.Links[i] {
					LinksDifferent = true
					break
				}
			}
		}
	}
	if LinksDifferent {
		ops := Diff(stage, system, systemOther, "Links", systemOther.Links, system.Links)
		diffs = append(diffs, ops)
	}
	if system.IsLinksNodeExpanded != systemOther.IsLinksNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsLinksNodeExpanded"))
	}
	LinksWhoseNodeIsExpandedDifferent := false
	if len(system.LinksWhoseNodeIsExpanded) != len(systemOther.LinksWhoseNodeIsExpanded) {
		LinksWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range system.LinksWhoseNodeIsExpanded {
			if (system.LinksWhoseNodeIsExpanded[i] == nil) != (systemOther.LinksWhoseNodeIsExpanded[i] == nil) {
				LinksWhoseNodeIsExpandedDifferent = true
				break
			} else if system.LinksWhoseNodeIsExpanded[i] != nil && systemOther.LinksWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if system.LinksWhoseNodeIsExpanded[i] != systemOther.LinksWhoseNodeIsExpanded[i] {
					LinksWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if LinksWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, system, systemOther, "LinksWhoseNodeIsExpanded", systemOther.LinksWhoseNodeIsExpanded, system.LinksWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
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
	if system.IsDiagramStructuresNodeExpanded != systemOther.IsDiagramStructuresNodeExpanded {
		diffs = append(diffs, system.GongMarshallField(stage, "IsDiagramStructuresNodeExpanded"))
	}
	DiagramStructuresWhoseNodeIsExpandedDifferent := false
	if len(system.DiagramStructuresWhoseNodeIsExpanded) != len(systemOther.DiagramStructuresWhoseNodeIsExpanded) {
		DiagramStructuresWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range system.DiagramStructuresWhoseNodeIsExpanded {
			if (system.DiagramStructuresWhoseNodeIsExpanded[i] == nil) != (systemOther.DiagramStructuresWhoseNodeIsExpanded[i] == nil) {
				DiagramStructuresWhoseNodeIsExpandedDifferent = true
				break
			} else if system.DiagramStructuresWhoseNodeIsExpanded[i] != nil && systemOther.DiagramStructuresWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if system.DiagramStructuresWhoseNodeIsExpanded[i] != systemOther.DiagramStructuresWhoseNodeIsExpanded[i] {
					DiagramStructuresWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if DiagramStructuresWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, system, systemOther, "DiagramStructuresWhoseNodeIsExpanded", systemOther.DiagramStructuresWhoseNodeIsExpanded, system.DiagramStructuresWhoseNodeIsExpanded)
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
	if systemshape.OverideLayoutDirection != systemshapeOther.OverideLayoutDirection {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "OverideLayoutDirection"))
	}
	if systemshape.LayoutDirection != systemshapeOther.LayoutDirection {
		diffs = append(diffs, systemshape.GongMarshallField(stage, "LayoutDirection"))
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
