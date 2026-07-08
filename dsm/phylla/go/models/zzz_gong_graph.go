// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *AxesShape:
		ok = stage.IsStagedAxesShape(target)

	case *GridPathShape:
		ok = stage.IsStagedGridPathShape(target)

	case *GrowthVectorShape:
		ok = stage.IsStagedGrowthVectorShape(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Plant:
		ok = stage.IsStagedPlant(target)

	case *PlantDiagram:
		ok = stage.IsStagedPlantDiagram(target)

	case *ReferenceRhombus:
		ok = stage.IsStagedReferenceRhombus(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *AxesShape:
		ok = stage.IsStagedAxesShape(target)

	case *GridPathShape:
		ok = stage.IsStagedGridPathShape(target)

	case *GrowthVectorShape:
		ok = stage.IsStagedGrowthVectorShape(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *Plant:
		ok = stage.IsStagedPlant(target)

	case *PlantDiagram:
		ok = stage.IsStagedPlantDiagram(target)

	case *ReferenceRhombus:
		ok = stage.IsStagedReferenceRhombus(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedAxesShape(axesshape *AxesShape) (ok bool) {

	_, ok = stage.AxesShapes[axesshape]

	return
}

func (stage *Stage) IsStagedGridPathShape(gridpathshape *GridPathShape) (ok bool) {

	_, ok = stage.GridPathShapes[gridpathshape]

	return
}

func (stage *Stage) IsStagedGrowthVectorShape(growthvectorshape *GrowthVectorShape) (ok bool) {

	_, ok = stage.GrowthVectorShapes[growthvectorshape]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedPlant(plant *Plant) (ok bool) {

	_, ok = stage.Plants[plant]

	return
}

func (stage *Stage) IsStagedPlantDiagram(plantdiagram *PlantDiagram) (ok bool) {

	_, ok = stage.PlantDiagrams[plantdiagram]

	return
}

func (stage *Stage) IsStagedReferenceRhombus(referencerhombus *ReferenceRhombus) (ok bool) {

	_, ok = stage.ReferenceRhombuss[referencerhombus]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *AxesShape:
		stage.StageBranchAxesShape(target)

	case *GridPathShape:
		stage.StageBranchGridPathShape(target)

	case *GrowthVectorShape:
		stage.StageBranchGrowthVectorShape(target)

	case *Library:
		stage.StageBranchLibrary(target)

	case *Plant:
		stage.StageBranchPlant(target)

	case *PlantDiagram:
		stage.StageBranchPlantDiagram(target)

	case *ReferenceRhombus:
		stage.StageBranchReferenceRhombus(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchAxesShape(axesshape *AxesShape) {

	// check if instance is already staged
	if IsStaged(stage, axesshape) {
		return
	}

	axesshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGridPathShape(gridpathshape *GridPathShape) {

	// check if instance is already staged
	if IsStaged(stage, gridpathshape) {
		return
	}

	gridpathshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGrowthVectorShape(growthvectorshape *GrowthVectorShape) {

	// check if instance is already staged
	if IsStaged(stage, growthvectorshape) {
		return
	}

	growthvectorshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

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
	for _, _plant := range library.Plants {
		StageBranch(stage, _plant)
	}

}

func (stage *Stage) StageBranchPlant(plant *Plant) {

	// check if instance is already staged
	if IsStaged(stage, plant) {
		return
	}

	plant.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plantdiagram := range plant.PlantDiagramsWhoseNodeIsExpanded {
		StageBranch(stage, _plantdiagram)
	}
	for _, _plantdiagram := range plant.PlantDiagrams {
		StageBranch(stage, _plantdiagram)
	}

}

func (stage *Stage) StageBranchPlantDiagram(plantdiagram *PlantDiagram) {

	// check if instance is already staged
	if IsStaged(stage, plantdiagram) {
		return
	}

	plantdiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plantdiagram.AxesShape != nil {
		StageBranch(stage, plantdiagram.AxesShape)
	}
	if plantdiagram.ReferenceRhombus != nil {
		StageBranch(stage, plantdiagram.ReferenceRhombus)
	}
	if plantdiagram.GrowthVectorShape != nil {
		StageBranch(stage, plantdiagram.GrowthVectorShape)
	}
	if plantdiagram.GridPathShape != nil {
		StageBranch(stage, plantdiagram.GridPathShape)
	}
	if plantdiagram.RotatedReferenceRhombus != nil {
		StageBranch(stage, plantdiagram.RotatedReferenceRhombus)
	}
	if plantdiagram.RotatedGrowthVectorShape != nil {
		StageBranch(stage, plantdiagram.RotatedGrowthVectorShape)
	}
	if plantdiagram.RotatedGridPathShape != nil {
		StageBranch(stage, plantdiagram.RotatedGridPathShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchReferenceRhombus(referencerhombus *ReferenceRhombus) {

	// check if instance is already staged
	if IsStaged(stage, referencerhombus) {
		return
	}

	referencerhombus.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

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
	case *AxesShape:
		toT := CopyBranchAxesShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GridPathShape:
		toT := CopyBranchGridPathShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GrowthVectorShape:
		toT := CopyBranchGrowthVectorShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := CopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Plant:
		toT := CopyBranchPlant(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PlantDiagram:
		toT := CopyBranchPlantDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ReferenceRhombus:
		toT := CopyBranchReferenceRhombus(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchAxesShape(mapOrigCopy map[any]any, axesshapeFrom *AxesShape) (axesshapeTo *AxesShape) {

	// axesshapeFrom has already been copied
	if _axesshapeTo, ok := mapOrigCopy[axesshapeFrom]; ok {
		axesshapeTo = _axesshapeTo.(*AxesShape)
		return
	}

	axesshapeTo = new(AxesShape)
	mapOrigCopy[axesshapeFrom] = axesshapeTo
	axesshapeFrom.CopyBasicFields(axesshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGridPathShape(mapOrigCopy map[any]any, gridpathshapeFrom *GridPathShape) (gridpathshapeTo *GridPathShape) {

	// gridpathshapeFrom has already been copied
	if _gridpathshapeTo, ok := mapOrigCopy[gridpathshapeFrom]; ok {
		gridpathshapeTo = _gridpathshapeTo.(*GridPathShape)
		return
	}

	gridpathshapeTo = new(GridPathShape)
	mapOrigCopy[gridpathshapeFrom] = gridpathshapeTo
	gridpathshapeFrom.CopyBasicFields(gridpathshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGrowthVectorShape(mapOrigCopy map[any]any, growthvectorshapeFrom *GrowthVectorShape) (growthvectorshapeTo *GrowthVectorShape) {

	// growthvectorshapeFrom has already been copied
	if _growthvectorshapeTo, ok := mapOrigCopy[growthvectorshapeFrom]; ok {
		growthvectorshapeTo = _growthvectorshapeTo.(*GrowthVectorShape)
		return
	}

	growthvectorshapeTo = new(GrowthVectorShape)
	mapOrigCopy[growthvectorshapeFrom] = growthvectorshapeTo
	growthvectorshapeFrom.CopyBasicFields(growthvectorshapeTo)

	//insertion point for the staging of instances referenced by pointers

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
	for _, _plant := range libraryFrom.Plants {
		libraryTo.Plants = append(libraryTo.Plants, CopyBranchPlant(mapOrigCopy, _plant))
	}

	return
}

func CopyBranchPlant(mapOrigCopy map[any]any, plantFrom *Plant) (plantTo *Plant) {

	// plantFrom has already been copied
	if _plantTo, ok := mapOrigCopy[plantFrom]; ok {
		plantTo = _plantTo.(*Plant)
		return
	}

	plantTo = new(Plant)
	mapOrigCopy[plantFrom] = plantTo
	plantFrom.CopyBasicFields(plantTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plantdiagram := range plantFrom.PlantDiagramsWhoseNodeIsExpanded {
		plantTo.PlantDiagramsWhoseNodeIsExpanded = append(plantTo.PlantDiagramsWhoseNodeIsExpanded, CopyBranchPlantDiagram(mapOrigCopy, _plantdiagram))
	}
	for _, _plantdiagram := range plantFrom.PlantDiagrams {
		plantTo.PlantDiagrams = append(plantTo.PlantDiagrams, CopyBranchPlantDiagram(mapOrigCopy, _plantdiagram))
	}

	return
}

func CopyBranchPlantDiagram(mapOrigCopy map[any]any, plantdiagramFrom *PlantDiagram) (plantdiagramTo *PlantDiagram) {

	// plantdiagramFrom has already been copied
	if _plantdiagramTo, ok := mapOrigCopy[plantdiagramFrom]; ok {
		plantdiagramTo = _plantdiagramTo.(*PlantDiagram)
		return
	}

	plantdiagramTo = new(PlantDiagram)
	mapOrigCopy[plantdiagramFrom] = plantdiagramTo
	plantdiagramFrom.CopyBasicFields(plantdiagramTo)

	//insertion point for the staging of instances referenced by pointers
	if plantdiagramFrom.AxesShape != nil {
		plantdiagramTo.AxesShape = CopyBranchAxesShape(mapOrigCopy, plantdiagramFrom.AxesShape)
	}
	if plantdiagramFrom.ReferenceRhombus != nil {
		plantdiagramTo.ReferenceRhombus = CopyBranchReferenceRhombus(mapOrigCopy, plantdiagramFrom.ReferenceRhombus)
	}
	if plantdiagramFrom.GrowthVectorShape != nil {
		plantdiagramTo.GrowthVectorShape = CopyBranchGrowthVectorShape(mapOrigCopy, plantdiagramFrom.GrowthVectorShape)
	}
	if plantdiagramFrom.GridPathShape != nil {
		plantdiagramTo.GridPathShape = CopyBranchGridPathShape(mapOrigCopy, plantdiagramFrom.GridPathShape)
	}
	if plantdiagramFrom.RotatedReferenceRhombus != nil {
		plantdiagramTo.RotatedReferenceRhombus = CopyBranchReferenceRhombus(mapOrigCopy, plantdiagramFrom.RotatedReferenceRhombus)
	}
	if plantdiagramFrom.RotatedGrowthVectorShape != nil {
		plantdiagramTo.RotatedGrowthVectorShape = CopyBranchGrowthVectorShape(mapOrigCopy, plantdiagramFrom.RotatedGrowthVectorShape)
	}
	if plantdiagramFrom.RotatedGridPathShape != nil {
		plantdiagramTo.RotatedGridPathShape = CopyBranchGridPathShape(mapOrigCopy, plantdiagramFrom.RotatedGridPathShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchReferenceRhombus(mapOrigCopy map[any]any, referencerhombusFrom *ReferenceRhombus) (referencerhombusTo *ReferenceRhombus) {

	// referencerhombusFrom has already been copied
	if _referencerhombusTo, ok := mapOrigCopy[referencerhombusFrom]; ok {
		referencerhombusTo = _referencerhombusTo.(*ReferenceRhombus)
		return
	}

	referencerhombusTo = new(ReferenceRhombus)
	mapOrigCopy[referencerhombusFrom] = referencerhombusTo
	referencerhombusFrom.CopyBasicFields(referencerhombusTo)

	//insertion point for the staging of instances referenced by pointers

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
	case *AxesShape:
		stage.UnstageBranchAxesShape(target)

	case *GridPathShape:
		stage.UnstageBranchGridPathShape(target)

	case *GrowthVectorShape:
		stage.UnstageBranchGrowthVectorShape(target)

	case *Library:
		stage.UnstageBranchLibrary(target)

	case *Plant:
		stage.UnstageBranchPlant(target)

	case *PlantDiagram:
		stage.UnstageBranchPlantDiagram(target)

	case *ReferenceRhombus:
		stage.UnstageBranchReferenceRhombus(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchAxesShape(axesshape *AxesShape) {

	// check if instance is already staged
	if !IsStaged(stage, axesshape) {
		return
	}

	axesshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGridPathShape(gridpathshape *GridPathShape) {

	// check if instance is already staged
	if !IsStaged(stage, gridpathshape) {
		return
	}

	gridpathshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGrowthVectorShape(growthvectorshape *GrowthVectorShape) {

	// check if instance is already staged
	if !IsStaged(stage, growthvectorshape) {
		return
	}

	growthvectorshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

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
	for _, _plant := range library.Plants {
		UnstageBranch(stage, _plant)
	}

}

func (stage *Stage) UnstageBranchPlant(plant *Plant) {

	// check if instance is already staged
	if !IsStaged(stage, plant) {
		return
	}

	plant.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plantdiagram := range plant.PlantDiagramsWhoseNodeIsExpanded {
		UnstageBranch(stage, _plantdiagram)
	}
	for _, _plantdiagram := range plant.PlantDiagrams {
		UnstageBranch(stage, _plantdiagram)
	}

}

func (stage *Stage) UnstageBranchPlantDiagram(plantdiagram *PlantDiagram) {

	// check if instance is already staged
	if !IsStaged(stage, plantdiagram) {
		return
	}

	plantdiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plantdiagram.AxesShape != nil {
		UnstageBranch(stage, plantdiagram.AxesShape)
	}
	if plantdiagram.ReferenceRhombus != nil {
		UnstageBranch(stage, plantdiagram.ReferenceRhombus)
	}
	if plantdiagram.GrowthVectorShape != nil {
		UnstageBranch(stage, plantdiagram.GrowthVectorShape)
	}
	if plantdiagram.GridPathShape != nil {
		UnstageBranch(stage, plantdiagram.GridPathShape)
	}
	if plantdiagram.RotatedReferenceRhombus != nil {
		UnstageBranch(stage, plantdiagram.RotatedReferenceRhombus)
	}
	if plantdiagram.RotatedGrowthVectorShape != nil {
		UnstageBranch(stage, plantdiagram.RotatedGrowthVectorShape)
	}
	if plantdiagram.RotatedGridPathShape != nil {
		UnstageBranch(stage, plantdiagram.RotatedGridPathShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchReferenceRhombus(referencerhombus *ReferenceRhombus) {

	// check if instance is already staged
	if !IsStaged(stage, referencerhombus) {
		return
	}

	referencerhombus.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *AxesShape) GongReconstructPointersFromReferences(stage *Stage, instance *AxesShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GridPathShape) GongReconstructPointersFromReferences(stage *Stage, instance *GridPathShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GrowthVectorShape) GongReconstructPointersFromReferences(stage *Stage, instance *GrowthVectorShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Library) GongReconstructPointersFromReferences(stage *Stage, instance *Library) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.SubLibraries = reference.SubLibraries[:0]
	for _, _b := range instance.SubLibraries {
		reference.SubLibraries = append(reference.SubLibraries, stage.Librarys_reference[_b])
	}
	reference.Plants = reference.Plants[:0]
	for _, _b := range instance.Plants {
		reference.Plants = append(reference.Plants, stage.Plants_reference[_b])
	}
}

func (reference *Plant) GongReconstructPointersFromReferences(stage *Stage, instance *Plant) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.PlantDiagramsWhoseNodeIsExpanded = reference.PlantDiagramsWhoseNodeIsExpanded[:0]
	for _, _b := range instance.PlantDiagramsWhoseNodeIsExpanded {
		reference.PlantDiagramsWhoseNodeIsExpanded = append(reference.PlantDiagramsWhoseNodeIsExpanded, stage.PlantDiagrams_reference[_b])
	}
	reference.PlantDiagrams = reference.PlantDiagrams[:0]
	for _, _b := range instance.PlantDiagrams {
		reference.PlantDiagrams = append(reference.PlantDiagrams, stage.PlantDiagrams_reference[_b])
	}
}

func (reference *PlantDiagram) GongReconstructPointersFromReferences(stage *Stage, instance *PlantDiagram) {
	// insertion point for pointers field
	if instance.AxesShape != nil {
		reference.AxesShape = stage.AxesShapes_reference[instance.AxesShape]
	}
	if instance.ReferenceRhombus != nil {
		reference.ReferenceRhombus = stage.ReferenceRhombuss_reference[instance.ReferenceRhombus]
	}
	if instance.GrowthVectorShape != nil {
		reference.GrowthVectorShape = stage.GrowthVectorShapes_reference[instance.GrowthVectorShape]
	}
	if instance.GridPathShape != nil {
		reference.GridPathShape = stage.GridPathShapes_reference[instance.GridPathShape]
	}
	if instance.RotatedReferenceRhombus != nil {
		reference.RotatedReferenceRhombus = stage.ReferenceRhombuss_reference[instance.RotatedReferenceRhombus]
	}
	if instance.RotatedGrowthVectorShape != nil {
		reference.RotatedGrowthVectorShape = stage.GrowthVectorShapes_reference[instance.RotatedGrowthVectorShape]
	}
	if instance.RotatedGridPathShape != nil {
		reference.RotatedGridPathShape = stage.GridPathShapes_reference[instance.RotatedGridPathShape]
	}
	// insertion point for slice of pointers field
}

func (reference *ReferenceRhombus) GongReconstructPointersFromReferences(stage *Stage, instance *ReferenceRhombus) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AxesShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GridPathShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GrowthVectorShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
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
	var _Plants []*Plant
	for _, _reference := range reference.Plants {
		if _instance, ok := stage.Plants_instance[_reference]; ok {
			_Plants = append(_Plants, _instance)
		}
	}
	reference.Plants = _Plants
}

func (reference *Plant) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _PlantDiagramsWhoseNodeIsExpanded []*PlantDiagram
	for _, _reference := range reference.PlantDiagramsWhoseNodeIsExpanded {
		if _instance, ok := stage.PlantDiagrams_instance[_reference]; ok {
			_PlantDiagramsWhoseNodeIsExpanded = append(_PlantDiagramsWhoseNodeIsExpanded, _instance)
		}
	}
	reference.PlantDiagramsWhoseNodeIsExpanded = _PlantDiagramsWhoseNodeIsExpanded
	var _PlantDiagrams []*PlantDiagram
	for _, _reference := range reference.PlantDiagrams {
		if _instance, ok := stage.PlantDiagrams_instance[_reference]; ok {
			_PlantDiagrams = append(_PlantDiagrams, _instance)
		}
	}
	reference.PlantDiagrams = _PlantDiagrams
}

func (reference *PlantDiagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.AxesShape; _reference != nil {
		reference.AxesShape = nil
		if _instance, ok := stage.AxesShapes_instance[_reference]; ok {
			reference.AxesShape = _instance
		}
	}
	if _reference := reference.ReferenceRhombus; _reference != nil {
		reference.ReferenceRhombus = nil
		if _instance, ok := stage.ReferenceRhombuss_instance[_reference]; ok {
			reference.ReferenceRhombus = _instance
		}
	}
	if _reference := reference.GrowthVectorShape; _reference != nil {
		reference.GrowthVectorShape = nil
		if _instance, ok := stage.GrowthVectorShapes_instance[_reference]; ok {
			reference.GrowthVectorShape = _instance
		}
	}
	if _reference := reference.GridPathShape; _reference != nil {
		reference.GridPathShape = nil
		if _instance, ok := stage.GridPathShapes_instance[_reference]; ok {
			reference.GridPathShape = _instance
		}
	}
	if _reference := reference.RotatedReferenceRhombus; _reference != nil {
		reference.RotatedReferenceRhombus = nil
		if _instance, ok := stage.ReferenceRhombuss_instance[_reference]; ok {
			reference.RotatedReferenceRhombus = _instance
		}
	}
	if _reference := reference.RotatedGrowthVectorShape; _reference != nil {
		reference.RotatedGrowthVectorShape = nil
		if _instance, ok := stage.GrowthVectorShapes_instance[_reference]; ok {
			reference.RotatedGrowthVectorShape = _instance
		}
	}
	if _reference := reference.RotatedGridPathShape; _reference != nil {
		reference.RotatedGridPathShape = nil
		if _instance, ok := stage.GridPathShapes_instance[_reference]; ok {
			reference.RotatedGridPathShape = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ReferenceRhombus) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (axesshape *AxesShape) GongDiff(stage *Stage, axesshapeOther *AxesShape) (diffs []string) {
	// insertion point for field diffs
	if axesshape.Name != axesshapeOther.Name {
		diffs = append(diffs, axesshape.GongMarshallField(stage, "Name"))
	}
	if axesshape.LengthX != axesshapeOther.LengthX {
		diffs = append(diffs, axesshape.GongMarshallField(stage, "LengthX"))
	}
	if axesshape.LengthY != axesshapeOther.LengthY {
		diffs = append(diffs, axesshape.GongMarshallField(stage, "LengthY"))
	}
	if axesshape.IsHidden != axesshapeOther.IsHidden {
		diffs = append(diffs, axesshape.GongMarshallField(stage, "IsHidden"))
	}
	if axesshape.IsWithHiddenHandle != axesshapeOther.IsWithHiddenHandle {
		diffs = append(diffs, axesshape.GongMarshallField(stage, "IsWithHiddenHandle"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gridpathshape *GridPathShape) GongDiff(stage *Stage, gridpathshapeOther *GridPathShape) (diffs []string) {
	// insertion point for field diffs
	if gridpathshape.Name != gridpathshapeOther.Name {
		diffs = append(diffs, gridpathshape.GongMarshallField(stage, "Name"))
	}
	if gridpathshape.IsHidden != gridpathshapeOther.IsHidden {
		diffs = append(diffs, gridpathshape.GongMarshallField(stage, "IsHidden"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (growthvectorshape *GrowthVectorShape) GongDiff(stage *Stage, growthvectorshapeOther *GrowthVectorShape) (diffs []string) {
	// insertion point for field diffs
	if growthvectorshape.Name != growthvectorshapeOther.Name {
		diffs = append(diffs, growthvectorshape.GongMarshallField(stage, "Name"))
	}
	if growthvectorshape.AngleDegree != growthvectorshapeOther.AngleDegree {
		diffs = append(diffs, growthvectorshape.GongMarshallField(stage, "AngleDegree"))
	}
	if growthvectorshape.Length != growthvectorshapeOther.Length {
		diffs = append(diffs, growthvectorshape.GongMarshallField(stage, "Length"))
	}
	if growthvectorshape.IsHidden != growthvectorshapeOther.IsHidden {
		diffs = append(diffs, growthvectorshape.GongMarshallField(stage, "IsHidden"))
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
	if library.IsRootLibrary != libraryOther.IsRootLibrary {
		diffs = append(diffs, library.GongMarshallField(stage, "IsRootLibrary"))
	}
	PlantsDifferent := false
	if len(library.Plants) != len(libraryOther.Plants) {
		PlantsDifferent = true
	} else {
		for i := range library.Plants {
			if (library.Plants[i] == nil) != (libraryOther.Plants[i] == nil) {
				PlantsDifferent = true
				break
			} else if library.Plants[i] != nil && libraryOther.Plants[i] != nil {
				// this is a pointer comparaison
				if library.Plants[i] != libraryOther.Plants[i] {
					PlantsDifferent = true
					break
				}
			}
		}
	}
	if PlantsDifferent {
		ops := Diff(stage, library, libraryOther, "Plants", libraryOther.Plants, library.Plants)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (plant *Plant) GongDiff(stage *Stage, plantOther *Plant) (diffs []string) {
	// insertion point for field diffs
	if plant.Name != plantOther.Name {
		diffs = append(diffs, plant.GongMarshallField(stage, "Name"))
	}
	if plant.N != plantOther.N {
		diffs = append(diffs, plant.GongMarshallField(stage, "N"))
	}
	if plant.M != plantOther.M {
		diffs = append(diffs, plant.GongMarshallField(stage, "M"))
	}
	if plant.Z != plantOther.Z {
		diffs = append(diffs, plant.GongMarshallField(stage, "Z"))
	}
	if plant.RhombusInsideAngle != plantOther.RhombusInsideAngle {
		diffs = append(diffs, plant.GongMarshallField(stage, "RhombusInsideAngle"))
	}
	if plant.RhombusSideLength != plantOther.RhombusSideLength {
		diffs = append(diffs, plant.GongMarshallField(stage, "RhombusSideLength"))
	}
	if plant.ShiftToNearestCircle != plantOther.ShiftToNearestCircle {
		diffs = append(diffs, plant.GongMarshallField(stage, "ShiftToNearestCircle"))
	}
	if plant.ComputedPrefix != plantOther.ComputedPrefix {
		diffs = append(diffs, plant.GongMarshallField(stage, "ComputedPrefix"))
	}
	if plant.IsExpanded != plantOther.IsExpanded {
		diffs = append(diffs, plant.GongMarshallField(stage, "IsExpanded"))
	}
	if plant.IsPlantDiagramsNodeExpanded != plantOther.IsPlantDiagramsNodeExpanded {
		diffs = append(diffs, plant.GongMarshallField(stage, "IsPlantDiagramsNodeExpanded"))
	}
	PlantDiagramsWhoseNodeIsExpandedDifferent := false
	if len(plant.PlantDiagramsWhoseNodeIsExpanded) != len(plantOther.PlantDiagramsWhoseNodeIsExpanded) {
		PlantDiagramsWhoseNodeIsExpandedDifferent = true
	} else {
		for i := range plant.PlantDiagramsWhoseNodeIsExpanded {
			if (plant.PlantDiagramsWhoseNodeIsExpanded[i] == nil) != (plantOther.PlantDiagramsWhoseNodeIsExpanded[i] == nil) {
				PlantDiagramsWhoseNodeIsExpandedDifferent = true
				break
			} else if plant.PlantDiagramsWhoseNodeIsExpanded[i] != nil && plantOther.PlantDiagramsWhoseNodeIsExpanded[i] != nil {
				// this is a pointer comparaison
				if plant.PlantDiagramsWhoseNodeIsExpanded[i] != plantOther.PlantDiagramsWhoseNodeIsExpanded[i] {
					PlantDiagramsWhoseNodeIsExpandedDifferent = true
					break
				}
			}
		}
	}
	if PlantDiagramsWhoseNodeIsExpandedDifferent {
		ops := Diff(stage, plant, plantOther, "PlantDiagramsWhoseNodeIsExpanded", plantOther.PlantDiagramsWhoseNodeIsExpanded, plant.PlantDiagramsWhoseNodeIsExpanded)
		diffs = append(diffs, ops)
	}
	PlantDiagramsDifferent := false
	if len(plant.PlantDiagrams) != len(plantOther.PlantDiagrams) {
		PlantDiagramsDifferent = true
	} else {
		for i := range plant.PlantDiagrams {
			if (plant.PlantDiagrams[i] == nil) != (plantOther.PlantDiagrams[i] == nil) {
				PlantDiagramsDifferent = true
				break
			} else if plant.PlantDiagrams[i] != nil && plantOther.PlantDiagrams[i] != nil {
				// this is a pointer comparaison
				if plant.PlantDiagrams[i] != plantOther.PlantDiagrams[i] {
					PlantDiagramsDifferent = true
					break
				}
			}
		}
	}
	if PlantDiagramsDifferent {
		ops := Diff(stage, plant, plantOther, "PlantDiagrams", plantOther.PlantDiagrams, plant.PlantDiagrams)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (plantdiagram *PlantDiagram) GongDiff(stage *Stage, plantdiagramOther *PlantDiagram) (diffs []string) {
	// insertion point for field diffs
	if plantdiagram.Name != plantdiagramOther.Name {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "Name"))
	}
	if plantdiagram.OriginX != plantdiagramOther.OriginX {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "OriginX"))
	}
	if plantdiagram.OriginY != plantdiagramOther.OriginY {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "OriginY"))
	}
	if (plantdiagram.AxesShape == nil) != (plantdiagramOther.AxesShape == nil) {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "AxesShape"))
	} else if plantdiagram.AxesShape != nil && plantdiagramOther.AxesShape != nil {
		if plantdiagram.AxesShape != plantdiagramOther.AxesShape {
			diffs = append(diffs, plantdiagram.GongMarshallField(stage, "AxesShape"))
		}
	}
	if (plantdiagram.ReferenceRhombus == nil) != (plantdiagramOther.ReferenceRhombus == nil) {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "ReferenceRhombus"))
	} else if plantdiagram.ReferenceRhombus != nil && plantdiagramOther.ReferenceRhombus != nil {
		if plantdiagram.ReferenceRhombus != plantdiagramOther.ReferenceRhombus {
			diffs = append(diffs, plantdiagram.GongMarshallField(stage, "ReferenceRhombus"))
		}
	}
	if (plantdiagram.GrowthVectorShape == nil) != (plantdiagramOther.GrowthVectorShape == nil) {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "GrowthVectorShape"))
	} else if plantdiagram.GrowthVectorShape != nil && plantdiagramOther.GrowthVectorShape != nil {
		if plantdiagram.GrowthVectorShape != plantdiagramOther.GrowthVectorShape {
			diffs = append(diffs, plantdiagram.GongMarshallField(stage, "GrowthVectorShape"))
		}
	}
	if (plantdiagram.GridPathShape == nil) != (plantdiagramOther.GridPathShape == nil) {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "GridPathShape"))
	} else if plantdiagram.GridPathShape != nil && plantdiagramOther.GridPathShape != nil {
		if plantdiagram.GridPathShape != plantdiagramOther.GridPathShape {
			diffs = append(diffs, plantdiagram.GongMarshallField(stage, "GridPathShape"))
		}
	}
	if (plantdiagram.RotatedReferenceRhombus == nil) != (plantdiagramOther.RotatedReferenceRhombus == nil) {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "RotatedReferenceRhombus"))
	} else if plantdiagram.RotatedReferenceRhombus != nil && plantdiagramOther.RotatedReferenceRhombus != nil {
		if plantdiagram.RotatedReferenceRhombus != plantdiagramOther.RotatedReferenceRhombus {
			diffs = append(diffs, plantdiagram.GongMarshallField(stage, "RotatedReferenceRhombus"))
		}
	}
	if (plantdiagram.RotatedGrowthVectorShape == nil) != (plantdiagramOther.RotatedGrowthVectorShape == nil) {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "RotatedGrowthVectorShape"))
	} else if plantdiagram.RotatedGrowthVectorShape != nil && plantdiagramOther.RotatedGrowthVectorShape != nil {
		if plantdiagram.RotatedGrowthVectorShape != plantdiagramOther.RotatedGrowthVectorShape {
			diffs = append(diffs, plantdiagram.GongMarshallField(stage, "RotatedGrowthVectorShape"))
		}
	}
	if (plantdiagram.RotatedGridPathShape == nil) != (plantdiagramOther.RotatedGridPathShape == nil) {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "RotatedGridPathShape"))
	} else if plantdiagram.RotatedGridPathShape != nil && plantdiagramOther.RotatedGridPathShape != nil {
		if plantdiagram.RotatedGridPathShape != plantdiagramOther.RotatedGridPathShape {
			diffs = append(diffs, plantdiagram.GongMarshallField(stage, "RotatedGridPathShape"))
		}
	}
	if plantdiagram.IsChecked != plantdiagramOther.IsChecked {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsChecked"))
	}
	if plantdiagram.ComputedPrefix != plantdiagramOther.ComputedPrefix {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "ComputedPrefix"))
	}
	if plantdiagram.IsExpanded != plantdiagramOther.IsExpanded {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsExpanded"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (referencerhombus *ReferenceRhombus) GongDiff(stage *Stage, referencerhombusOther *ReferenceRhombus) (diffs []string) {
	// insertion point for field diffs
	if referencerhombus.Name != referencerhombusOther.Name {
		diffs = append(diffs, referencerhombus.GongMarshallField(stage, "Name"))
	}
	if referencerhombus.IsHidden != referencerhombusOther.IsHidden {
		diffs = append(diffs, referencerhombus.GongMarshallField(stage, "IsHidden"))
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
