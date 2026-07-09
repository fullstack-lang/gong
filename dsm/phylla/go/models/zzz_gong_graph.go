// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *AxesShape:
		ok = stage.IsStagedAxesShape(target)

	case *CircleGridShape:
		ok = stage.IsStagedCircleGridShape(target)

	case *ExplanationTextShape:
		ok = stage.IsStagedExplanationTextShape(target)

	case *GridPathShape:
		ok = stage.IsStagedGridPathShape(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *NextCircleShape:
		ok = stage.IsStagedNextCircleShape(target)

	case *Plant:
		ok = stage.IsStagedPlant(target)

	case *PlantCircumferenceShape:
		ok = stage.IsStagedPlantCircumferenceShape(target)

	case *PlantDiagram:
		ok = stage.IsStagedPlantDiagram(target)

	case *RhombusGridShape:
		ok = stage.IsStagedRhombusGridShape(target)

	case *RhombusShape:
		ok = stage.IsStagedRhombusShape(target)

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

	case *CircleGridShape:
		ok = stage.IsStagedCircleGridShape(target)

	case *ExplanationTextShape:
		ok = stage.IsStagedExplanationTextShape(target)

	case *GridPathShape:
		ok = stage.IsStagedGridPathShape(target)

	case *Library:
		ok = stage.IsStagedLibrary(target)

	case *NextCircleShape:
		ok = stage.IsStagedNextCircleShape(target)

	case *Plant:
		ok = stage.IsStagedPlant(target)

	case *PlantCircumferenceShape:
		ok = stage.IsStagedPlantCircumferenceShape(target)

	case *PlantDiagram:
		ok = stage.IsStagedPlantDiagram(target)

	case *RhombusGridShape:
		ok = stage.IsStagedRhombusGridShape(target)

	case *RhombusShape:
		ok = stage.IsStagedRhombusShape(target)

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

func (stage *Stage) IsStagedCircleGridShape(circlegridshape *CircleGridShape) (ok bool) {

	_, ok = stage.CircleGridShapes[circlegridshape]

	return
}

func (stage *Stage) IsStagedExplanationTextShape(explanationtextshape *ExplanationTextShape) (ok bool) {

	_, ok = stage.ExplanationTextShapes[explanationtextshape]

	return
}

func (stage *Stage) IsStagedGridPathShape(gridpathshape *GridPathShape) (ok bool) {

	_, ok = stage.GridPathShapes[gridpathshape]

	return
}

func (stage *Stage) IsStagedLibrary(library *Library) (ok bool) {

	_, ok = stage.Librarys[library]

	return
}

func (stage *Stage) IsStagedNextCircleShape(nextcircleshape *NextCircleShape) (ok bool) {

	_, ok = stage.NextCircleShapes[nextcircleshape]

	return
}

func (stage *Stage) IsStagedPlant(plant *Plant) (ok bool) {

	_, ok = stage.Plants[plant]

	return
}

func (stage *Stage) IsStagedPlantCircumferenceShape(plantcircumferenceshape *PlantCircumferenceShape) (ok bool) {

	_, ok = stage.PlantCircumferenceShapes[plantcircumferenceshape]

	return
}

func (stage *Stage) IsStagedPlantDiagram(plantdiagram *PlantDiagram) (ok bool) {

	_, ok = stage.PlantDiagrams[plantdiagram]

	return
}

func (stage *Stage) IsStagedRhombusGridShape(rhombusgridshape *RhombusGridShape) (ok bool) {

	_, ok = stage.RhombusGridShapes[rhombusgridshape]

	return
}

func (stage *Stage) IsStagedRhombusShape(rhombusshape *RhombusShape) (ok bool) {

	_, ok = stage.RhombusShapes[rhombusshape]

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

	case *CircleGridShape:
		stage.StageBranchCircleGridShape(target)

	case *ExplanationTextShape:
		stage.StageBranchExplanationTextShape(target)

	case *GridPathShape:
		stage.StageBranchGridPathShape(target)

	case *Library:
		stage.StageBranchLibrary(target)

	case *NextCircleShape:
		stage.StageBranchNextCircleShape(target)

	case *Plant:
		stage.StageBranchPlant(target)

	case *PlantCircumferenceShape:
		stage.StageBranchPlantCircumferenceShape(target)

	case *PlantDiagram:
		stage.StageBranchPlantDiagram(target)

	case *RhombusGridShape:
		stage.StageBranchRhombusGridShape(target)

	case *RhombusShape:
		stage.StageBranchRhombusShape(target)

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

func (stage *Stage) StageBranchCircleGridShape(circlegridshape *CircleGridShape) {

	// check if instance is already staged
	if IsStaged(stage, circlegridshape) {
		return
	}

	circlegridshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchExplanationTextShape(explanationtextshape *ExplanationTextShape) {

	// check if instance is already staged
	if IsStaged(stage, explanationtextshape) {
		return
	}

	explanationtextshape.Stage(stage)

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

func (stage *Stage) StageBranchNextCircleShape(nextcircleshape *NextCircleShape) {

	// check if instance is already staged
	if IsStaged(stage, nextcircleshape) {
		return
	}

	nextcircleshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPlant(plant *Plant) {

	// check if instance is already staged
	if IsStaged(stage, plant) {
		return
	}

	plant.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plant.AxesShape != nil {
		StageBranch(stage, plant.AxesShape)
	}
	if plant.ReferenceRhombus != nil {
		StageBranch(stage, plant.ReferenceRhombus)
	}
	if plant.PlantCircumferenceShape != nil {
		StageBranch(stage, plant.PlantCircumferenceShape)
	}
	if plant.GridPathShape != nil {
		StageBranch(stage, plant.GridPathShape)
	}
	if plant.RhombusGridShape != nil {
		StageBranch(stage, plant.RhombusGridShape)
	}
	if plant.ExplanationTextShape != nil {
		StageBranch(stage, plant.ExplanationTextShape)
	}
	if plant.RotatedReferenceRhombus != nil {
		StageBranch(stage, plant.RotatedReferenceRhombus)
	}
	if plant.RotatedPlantCircumferenceShape != nil {
		StageBranch(stage, plant.RotatedPlantCircumferenceShape)
	}
	if plant.RotatedGridPathShape != nil {
		StageBranch(stage, plant.RotatedGridPathShape)
	}
	if plant.RotatedRhombusGridShape != nil {
		StageBranch(stage, plant.RotatedRhombusGridShape)
	}
	if plant.GrowthPathRhombusGridShape != nil {
		StageBranch(stage, plant.GrowthPathRhombusGridShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plantdiagram := range plant.PlantDiagrams {
		StageBranch(stage, _plantdiagram)
	}

}

func (stage *Stage) StageBranchPlantCircumferenceShape(plantcircumferenceshape *PlantCircumferenceShape) {

	// check if instance is already staged
	if IsStaged(stage, plantcircumferenceshape) {
		return
	}

	plantcircumferenceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPlantDiagram(plantdiagram *PlantDiagram) {

	// check if instance is already staged
	if IsStaged(stage, plantdiagram) {
		return
	}

	plantdiagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchRhombusGridShape(rhombusgridshape *RhombusGridShape) {

	// check if instance is already staged
	if IsStaged(stage, rhombusgridshape) {
		return
	}

	rhombusgridshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rhombusshape := range rhombusgridshape.RhombusShapes {
		StageBranch(stage, _rhombusshape)
	}

}

func (stage *Stage) StageBranchRhombusShape(rhombusshape *RhombusShape) {

	// check if instance is already staged
	if IsStaged(stage, rhombusshape) {
		return
	}

	rhombusshape.Stage(stage)

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

	case *CircleGridShape:
		toT := CopyBranchCircleGridShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ExplanationTextShape:
		toT := CopyBranchExplanationTextShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GridPathShape:
		toT := CopyBranchGridPathShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Library:
		toT := CopyBranchLibrary(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *NextCircleShape:
		toT := CopyBranchNextCircleShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Plant:
		toT := CopyBranchPlant(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PlantCircumferenceShape:
		toT := CopyBranchPlantCircumferenceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PlantDiagram:
		toT := CopyBranchPlantDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RhombusGridShape:
		toT := CopyBranchRhombusGridShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RhombusShape:
		toT := CopyBranchRhombusShape(mapOrigCopy, fromT)
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

func CopyBranchCircleGridShape(mapOrigCopy map[any]any, circlegridshapeFrom *CircleGridShape) (circlegridshapeTo *CircleGridShape) {

	// circlegridshapeFrom has already been copied
	if _circlegridshapeTo, ok := mapOrigCopy[circlegridshapeFrom]; ok {
		circlegridshapeTo = _circlegridshapeTo.(*CircleGridShape)
		return
	}

	circlegridshapeTo = new(CircleGridShape)
	mapOrigCopy[circlegridshapeFrom] = circlegridshapeTo
	circlegridshapeFrom.CopyBasicFields(circlegridshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchExplanationTextShape(mapOrigCopy map[any]any, explanationtextshapeFrom *ExplanationTextShape) (explanationtextshapeTo *ExplanationTextShape) {

	// explanationtextshapeFrom has already been copied
	if _explanationtextshapeTo, ok := mapOrigCopy[explanationtextshapeFrom]; ok {
		explanationtextshapeTo = _explanationtextshapeTo.(*ExplanationTextShape)
		return
	}

	explanationtextshapeTo = new(ExplanationTextShape)
	mapOrigCopy[explanationtextshapeFrom] = explanationtextshapeTo
	explanationtextshapeFrom.CopyBasicFields(explanationtextshapeTo)

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

func CopyBranchNextCircleShape(mapOrigCopy map[any]any, nextcircleshapeFrom *NextCircleShape) (nextcircleshapeTo *NextCircleShape) {

	// nextcircleshapeFrom has already been copied
	if _nextcircleshapeTo, ok := mapOrigCopy[nextcircleshapeFrom]; ok {
		nextcircleshapeTo = _nextcircleshapeTo.(*NextCircleShape)
		return
	}

	nextcircleshapeTo = new(NextCircleShape)
	mapOrigCopy[nextcircleshapeFrom] = nextcircleshapeTo
	nextcircleshapeFrom.CopyBasicFields(nextcircleshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

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
	if plantFrom.AxesShape != nil {
		plantTo.AxesShape = CopyBranchAxesShape(mapOrigCopy, plantFrom.AxesShape)
	}
	if plantFrom.ReferenceRhombus != nil {
		plantTo.ReferenceRhombus = CopyBranchRhombusShape(mapOrigCopy, plantFrom.ReferenceRhombus)
	}
	if plantFrom.PlantCircumferenceShape != nil {
		plantTo.PlantCircumferenceShape = CopyBranchPlantCircumferenceShape(mapOrigCopy, plantFrom.PlantCircumferenceShape)
	}
	if plantFrom.GridPathShape != nil {
		plantTo.GridPathShape = CopyBranchGridPathShape(mapOrigCopy, plantFrom.GridPathShape)
	}
	if plantFrom.RhombusGridShape != nil {
		plantTo.RhombusGridShape = CopyBranchRhombusGridShape(mapOrigCopy, plantFrom.RhombusGridShape)
	}
	if plantFrom.ExplanationTextShape != nil {
		plantTo.ExplanationTextShape = CopyBranchExplanationTextShape(mapOrigCopy, plantFrom.ExplanationTextShape)
	}
	if plantFrom.RotatedReferenceRhombus != nil {
		plantTo.RotatedReferenceRhombus = CopyBranchRhombusShape(mapOrigCopy, plantFrom.RotatedReferenceRhombus)
	}
	if plantFrom.RotatedPlantCircumferenceShape != nil {
		plantTo.RotatedPlantCircumferenceShape = CopyBranchPlantCircumferenceShape(mapOrigCopy, plantFrom.RotatedPlantCircumferenceShape)
	}
	if plantFrom.RotatedGridPathShape != nil {
		plantTo.RotatedGridPathShape = CopyBranchGridPathShape(mapOrigCopy, plantFrom.RotatedGridPathShape)
	}
	if plantFrom.RotatedRhombusGridShape != nil {
		plantTo.RotatedRhombusGridShape = CopyBranchRhombusGridShape(mapOrigCopy, plantFrom.RotatedRhombusGridShape)
	}
	if plantFrom.GrowthPathRhombusGridShape != nil {
		plantTo.GrowthPathRhombusGridShape = CopyBranchRhombusGridShape(mapOrigCopy, plantFrom.GrowthPathRhombusGridShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plantdiagram := range plantFrom.PlantDiagrams {
		plantTo.PlantDiagrams = append(plantTo.PlantDiagrams, CopyBranchPlantDiagram(mapOrigCopy, _plantdiagram))
	}

	return
}

func CopyBranchPlantCircumferenceShape(mapOrigCopy map[any]any, plantcircumferenceshapeFrom *PlantCircumferenceShape) (plantcircumferenceshapeTo *PlantCircumferenceShape) {

	// plantcircumferenceshapeFrom has already been copied
	if _plantcircumferenceshapeTo, ok := mapOrigCopy[plantcircumferenceshapeFrom]; ok {
		plantcircumferenceshapeTo = _plantcircumferenceshapeTo.(*PlantCircumferenceShape)
		return
	}

	plantcircumferenceshapeTo = new(PlantCircumferenceShape)
	mapOrigCopy[plantcircumferenceshapeFrom] = plantcircumferenceshapeTo
	plantcircumferenceshapeFrom.CopyBasicFields(plantcircumferenceshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

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

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchRhombusGridShape(mapOrigCopy map[any]any, rhombusgridshapeFrom *RhombusGridShape) (rhombusgridshapeTo *RhombusGridShape) {

	// rhombusgridshapeFrom has already been copied
	if _rhombusgridshapeTo, ok := mapOrigCopy[rhombusgridshapeFrom]; ok {
		rhombusgridshapeTo = _rhombusgridshapeTo.(*RhombusGridShape)
		return
	}

	rhombusgridshapeTo = new(RhombusGridShape)
	mapOrigCopy[rhombusgridshapeFrom] = rhombusgridshapeTo
	rhombusgridshapeFrom.CopyBasicFields(rhombusgridshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rhombusshape := range rhombusgridshapeFrom.RhombusShapes {
		rhombusgridshapeTo.RhombusShapes = append(rhombusgridshapeTo.RhombusShapes, CopyBranchRhombusShape(mapOrigCopy, _rhombusshape))
	}

	return
}

func CopyBranchRhombusShape(mapOrigCopy map[any]any, rhombusshapeFrom *RhombusShape) (rhombusshapeTo *RhombusShape) {

	// rhombusshapeFrom has already been copied
	if _rhombusshapeTo, ok := mapOrigCopy[rhombusshapeFrom]; ok {
		rhombusshapeTo = _rhombusshapeTo.(*RhombusShape)
		return
	}

	rhombusshapeTo = new(RhombusShape)
	mapOrigCopy[rhombusshapeFrom] = rhombusshapeTo
	rhombusshapeFrom.CopyBasicFields(rhombusshapeTo)

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

	case *CircleGridShape:
		stage.UnstageBranchCircleGridShape(target)

	case *ExplanationTextShape:
		stage.UnstageBranchExplanationTextShape(target)

	case *GridPathShape:
		stage.UnstageBranchGridPathShape(target)

	case *Library:
		stage.UnstageBranchLibrary(target)

	case *NextCircleShape:
		stage.UnstageBranchNextCircleShape(target)

	case *Plant:
		stage.UnstageBranchPlant(target)

	case *PlantCircumferenceShape:
		stage.UnstageBranchPlantCircumferenceShape(target)

	case *PlantDiagram:
		stage.UnstageBranchPlantDiagram(target)

	case *RhombusGridShape:
		stage.UnstageBranchRhombusGridShape(target)

	case *RhombusShape:
		stage.UnstageBranchRhombusShape(target)

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

func (stage *Stage) UnstageBranchCircleGridShape(circlegridshape *CircleGridShape) {

	// check if instance is already staged
	if !IsStaged(stage, circlegridshape) {
		return
	}

	circlegridshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchExplanationTextShape(explanationtextshape *ExplanationTextShape) {

	// check if instance is already staged
	if !IsStaged(stage, explanationtextshape) {
		return
	}

	explanationtextshape.Unstage(stage)

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

func (stage *Stage) UnstageBranchNextCircleShape(nextcircleshape *NextCircleShape) {

	// check if instance is already staged
	if !IsStaged(stage, nextcircleshape) {
		return
	}

	nextcircleshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPlant(plant *Plant) {

	// check if instance is already staged
	if !IsStaged(stage, plant) {
		return
	}

	plant.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if plant.AxesShape != nil {
		UnstageBranch(stage, plant.AxesShape)
	}
	if plant.ReferenceRhombus != nil {
		UnstageBranch(stage, plant.ReferenceRhombus)
	}
	if plant.PlantCircumferenceShape != nil {
		UnstageBranch(stage, plant.PlantCircumferenceShape)
	}
	if plant.GridPathShape != nil {
		UnstageBranch(stage, plant.GridPathShape)
	}
	if plant.RhombusGridShape != nil {
		UnstageBranch(stage, plant.RhombusGridShape)
	}
	if plant.ExplanationTextShape != nil {
		UnstageBranch(stage, plant.ExplanationTextShape)
	}
	if plant.RotatedReferenceRhombus != nil {
		UnstageBranch(stage, plant.RotatedReferenceRhombus)
	}
	if plant.RotatedPlantCircumferenceShape != nil {
		UnstageBranch(stage, plant.RotatedPlantCircumferenceShape)
	}
	if plant.RotatedGridPathShape != nil {
		UnstageBranch(stage, plant.RotatedGridPathShape)
	}
	if plant.RotatedRhombusGridShape != nil {
		UnstageBranch(stage, plant.RotatedRhombusGridShape)
	}
	if plant.GrowthPathRhombusGridShape != nil {
		UnstageBranch(stage, plant.GrowthPathRhombusGridShape)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _plantdiagram := range plant.PlantDiagrams {
		UnstageBranch(stage, _plantdiagram)
	}

}

func (stage *Stage) UnstageBranchPlantCircumferenceShape(plantcircumferenceshape *PlantCircumferenceShape) {

	// check if instance is already staged
	if !IsStaged(stage, plantcircumferenceshape) {
		return
	}

	plantcircumferenceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPlantDiagram(plantdiagram *PlantDiagram) {

	// check if instance is already staged
	if !IsStaged(stage, plantdiagram) {
		return
	}

	plantdiagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchRhombusGridShape(rhombusgridshape *RhombusGridShape) {

	// check if instance is already staged
	if !IsStaged(stage, rhombusgridshape) {
		return
	}

	rhombusgridshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _rhombusshape := range rhombusgridshape.RhombusShapes {
		UnstageBranch(stage, _rhombusshape)
	}

}

func (stage *Stage) UnstageBranchRhombusShape(rhombusshape *RhombusShape) {

	// check if instance is already staged
	if !IsStaged(stage, rhombusshape) {
		return
	}

	rhombusshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *AxesShape) GongReconstructPointersFromReferences(stage *Stage, instance *AxesShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CircleGridShape) GongReconstructPointersFromReferences(stage *Stage, instance *CircleGridShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ExplanationTextShape) GongReconstructPointersFromReferences(stage *Stage, instance *ExplanationTextShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GridPathShape) GongReconstructPointersFromReferences(stage *Stage, instance *GridPathShape) {
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

func (reference *NextCircleShape) GongReconstructPointersFromReferences(stage *Stage, instance *NextCircleShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *Plant) GongReconstructPointersFromReferences(stage *Stage, instance *Plant) {
	// insertion point for pointers field
	if instance.AxesShape != nil {
		reference.AxesShape = stage.AxesShapes_reference[instance.AxesShape]
	}
	if instance.ReferenceRhombus != nil {
		reference.ReferenceRhombus = stage.RhombusShapes_reference[instance.ReferenceRhombus]
	}
	if instance.PlantCircumferenceShape != nil {
		reference.PlantCircumferenceShape = stage.PlantCircumferenceShapes_reference[instance.PlantCircumferenceShape]
	}
	if instance.GridPathShape != nil {
		reference.GridPathShape = stage.GridPathShapes_reference[instance.GridPathShape]
	}
	if instance.RhombusGridShape != nil {
		reference.RhombusGridShape = stage.RhombusGridShapes_reference[instance.RhombusGridShape]
	}
	if instance.ExplanationTextShape != nil {
		reference.ExplanationTextShape = stage.ExplanationTextShapes_reference[instance.ExplanationTextShape]
	}
	if instance.RotatedReferenceRhombus != nil {
		reference.RotatedReferenceRhombus = stage.RhombusShapes_reference[instance.RotatedReferenceRhombus]
	}
	if instance.RotatedPlantCircumferenceShape != nil {
		reference.RotatedPlantCircumferenceShape = stage.PlantCircumferenceShapes_reference[instance.RotatedPlantCircumferenceShape]
	}
	if instance.RotatedGridPathShape != nil {
		reference.RotatedGridPathShape = stage.GridPathShapes_reference[instance.RotatedGridPathShape]
	}
	if instance.RotatedRhombusGridShape != nil {
		reference.RotatedRhombusGridShape = stage.RhombusGridShapes_reference[instance.RotatedRhombusGridShape]
	}
	if instance.GrowthPathRhombusGridShape != nil {
		reference.GrowthPathRhombusGridShape = stage.RhombusGridShapes_reference[instance.GrowthPathRhombusGridShape]
	}
	// insertion point for slice of pointers field
	reference.PlantDiagrams = reference.PlantDiagrams[:0]
	for _, _b := range instance.PlantDiagrams {
		reference.PlantDiagrams = append(reference.PlantDiagrams, stage.PlantDiagrams_reference[_b])
	}
}

func (reference *PlantCircumferenceShape) GongReconstructPointersFromReferences(stage *Stage, instance *PlantCircumferenceShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PlantDiagram) GongReconstructPointersFromReferences(stage *Stage, instance *PlantDiagram) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *RhombusGridShape) GongReconstructPointersFromReferences(stage *Stage, instance *RhombusGridShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.RhombusShapes = reference.RhombusShapes[:0]
	for _, _b := range instance.RhombusShapes {
		reference.RhombusShapes = append(reference.RhombusShapes, stage.RhombusShapes_reference[_b])
	}
}

func (reference *RhombusShape) GongReconstructPointersFromReferences(stage *Stage, instance *RhombusShape) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AxesShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CircleGridShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ExplanationTextShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GridPathShape) GongReconstructPointersFromInstances(stage *Stage) {
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

func (reference *NextCircleShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *Plant) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.AxesShape; _reference != nil {
		reference.AxesShape = nil
		if _instance, ok := stage.AxesShapes_instance[_reference]; ok {
			reference.AxesShape = _instance
		}
	}
	if _reference := reference.ReferenceRhombus; _reference != nil {
		reference.ReferenceRhombus = nil
		if _instance, ok := stage.RhombusShapes_instance[_reference]; ok {
			reference.ReferenceRhombus = _instance
		}
	}
	if _reference := reference.PlantCircumferenceShape; _reference != nil {
		reference.PlantCircumferenceShape = nil
		if _instance, ok := stage.PlantCircumferenceShapes_instance[_reference]; ok {
			reference.PlantCircumferenceShape = _instance
		}
	}
	if _reference := reference.GridPathShape; _reference != nil {
		reference.GridPathShape = nil
		if _instance, ok := stage.GridPathShapes_instance[_reference]; ok {
			reference.GridPathShape = _instance
		}
	}
	if _reference := reference.RhombusGridShape; _reference != nil {
		reference.RhombusGridShape = nil
		if _instance, ok := stage.RhombusGridShapes_instance[_reference]; ok {
			reference.RhombusGridShape = _instance
		}
	}
	if _reference := reference.ExplanationTextShape; _reference != nil {
		reference.ExplanationTextShape = nil
		if _instance, ok := stage.ExplanationTextShapes_instance[_reference]; ok {
			reference.ExplanationTextShape = _instance
		}
	}
	if _reference := reference.RotatedReferenceRhombus; _reference != nil {
		reference.RotatedReferenceRhombus = nil
		if _instance, ok := stage.RhombusShapes_instance[_reference]; ok {
			reference.RotatedReferenceRhombus = _instance
		}
	}
	if _reference := reference.RotatedPlantCircumferenceShape; _reference != nil {
		reference.RotatedPlantCircumferenceShape = nil
		if _instance, ok := stage.PlantCircumferenceShapes_instance[_reference]; ok {
			reference.RotatedPlantCircumferenceShape = _instance
		}
	}
	if _reference := reference.RotatedGridPathShape; _reference != nil {
		reference.RotatedGridPathShape = nil
		if _instance, ok := stage.GridPathShapes_instance[_reference]; ok {
			reference.RotatedGridPathShape = _instance
		}
	}
	if _reference := reference.RotatedRhombusGridShape; _reference != nil {
		reference.RotatedRhombusGridShape = nil
		if _instance, ok := stage.RhombusGridShapes_instance[_reference]; ok {
			reference.RotatedRhombusGridShape = _instance
		}
	}
	if _reference := reference.GrowthPathRhombusGridShape; _reference != nil {
		reference.GrowthPathRhombusGridShape = nil
		if _instance, ok := stage.RhombusGridShapes_instance[_reference]; ok {
			reference.GrowthPathRhombusGridShape = _instance
		}
	}
	// insertion point for slice of pointers fields
	var _PlantDiagrams []*PlantDiagram
	for _, _reference := range reference.PlantDiagrams {
		if _instance, ok := stage.PlantDiagrams_instance[_reference]; ok {
			_PlantDiagrams = append(_PlantDiagrams, _instance)
		}
	}
	reference.PlantDiagrams = _PlantDiagrams
}

func (reference *PlantCircumferenceShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PlantDiagram) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *RhombusGridShape) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _RhombusShapes []*RhombusShape
	for _, _reference := range reference.RhombusShapes {
		if _instance, ok := stage.RhombusShapes_instance[_reference]; ok {
			_RhombusShapes = append(_RhombusShapes, _instance)
		}
	}
	reference.RhombusShapes = _RhombusShapes
}

func (reference *RhombusShape) GongReconstructPointersFromInstances(stage *Stage) {
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
	if axesshape.IsWithHiddenHandle != axesshapeOther.IsWithHiddenHandle {
		diffs = append(diffs, axesshape.GongMarshallField(stage, "IsWithHiddenHandle"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (circlegridshape *CircleGridShape) GongDiff(stage *Stage, circlegridshapeOther *CircleGridShape) (diffs []string) {
	// insertion point for field diffs
	if circlegridshape.Name != circlegridshapeOther.Name {
		diffs = append(diffs, circlegridshape.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (explanationtextshape *ExplanationTextShape) GongDiff(stage *Stage, explanationtextshapeOther *ExplanationTextShape) (diffs []string) {
	// insertion point for field diffs
	if explanationtextshape.Name != explanationtextshapeOther.Name {
		diffs = append(diffs, explanationtextshape.GongMarshallField(stage, "Name"))
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
func (nextcircleshape *NextCircleShape) GongDiff(stage *Stage, nextcircleshapeOther *NextCircleShape) (diffs []string) {
	// insertion point for field diffs
	if nextcircleshape.Name != nextcircleshapeOther.Name {
		diffs = append(diffs, nextcircleshape.GongMarshallField(stage, "Name"))
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
	if plant.IsSelected != plantOther.IsSelected {
		diffs = append(diffs, plant.GongMarshallField(stage, "IsSelected"))
	}
	if plant.IsPlantDiagramsNodeExpanded != plantOther.IsPlantDiagramsNodeExpanded {
		diffs = append(diffs, plant.GongMarshallField(stage, "IsPlantDiagramsNodeExpanded"))
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
	if (plant.AxesShape == nil) != (plantOther.AxesShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "AxesShape"))
	} else if plant.AxesShape != nil && plantOther.AxesShape != nil {
		if plant.AxesShape != plantOther.AxesShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "AxesShape"))
		}
	}
	if (plant.ReferenceRhombus == nil) != (plantOther.ReferenceRhombus == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "ReferenceRhombus"))
	} else if plant.ReferenceRhombus != nil && plantOther.ReferenceRhombus != nil {
		if plant.ReferenceRhombus != plantOther.ReferenceRhombus {
			diffs = append(diffs, plant.GongMarshallField(stage, "ReferenceRhombus"))
		}
	}
	if (plant.PlantCircumferenceShape == nil) != (plantOther.PlantCircumferenceShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "PlantCircumferenceShape"))
	} else if plant.PlantCircumferenceShape != nil && plantOther.PlantCircumferenceShape != nil {
		if plant.PlantCircumferenceShape != plantOther.PlantCircumferenceShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "PlantCircumferenceShape"))
		}
	}
	if (plant.GridPathShape == nil) != (plantOther.GridPathShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "GridPathShape"))
	} else if plant.GridPathShape != nil && plantOther.GridPathShape != nil {
		if plant.GridPathShape != plantOther.GridPathShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "GridPathShape"))
		}
	}
	if (plant.RhombusGridShape == nil) != (plantOther.RhombusGridShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "RhombusGridShape"))
	} else if plant.RhombusGridShape != nil && plantOther.RhombusGridShape != nil {
		if plant.RhombusGridShape != plantOther.RhombusGridShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "RhombusGridShape"))
		}
	}
	if (plant.ExplanationTextShape == nil) != (plantOther.ExplanationTextShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "ExplanationTextShape"))
	} else if plant.ExplanationTextShape != nil && plantOther.ExplanationTextShape != nil {
		if plant.ExplanationTextShape != plantOther.ExplanationTextShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "ExplanationTextShape"))
		}
	}
	if (plant.RotatedReferenceRhombus == nil) != (plantOther.RotatedReferenceRhombus == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "RotatedReferenceRhombus"))
	} else if plant.RotatedReferenceRhombus != nil && plantOther.RotatedReferenceRhombus != nil {
		if plant.RotatedReferenceRhombus != plantOther.RotatedReferenceRhombus {
			diffs = append(diffs, plant.GongMarshallField(stage, "RotatedReferenceRhombus"))
		}
	}
	if (plant.RotatedPlantCircumferenceShape == nil) != (plantOther.RotatedPlantCircumferenceShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "RotatedPlantCircumferenceShape"))
	} else if plant.RotatedPlantCircumferenceShape != nil && plantOther.RotatedPlantCircumferenceShape != nil {
		if plant.RotatedPlantCircumferenceShape != plantOther.RotatedPlantCircumferenceShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "RotatedPlantCircumferenceShape"))
		}
	}
	if (plant.RotatedGridPathShape == nil) != (plantOther.RotatedGridPathShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "RotatedGridPathShape"))
	} else if plant.RotatedGridPathShape != nil && plantOther.RotatedGridPathShape != nil {
		if plant.RotatedGridPathShape != plantOther.RotatedGridPathShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "RotatedGridPathShape"))
		}
	}
	if (plant.RotatedRhombusGridShape == nil) != (plantOther.RotatedRhombusGridShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "RotatedRhombusGridShape"))
	} else if plant.RotatedRhombusGridShape != nil && plantOther.RotatedRhombusGridShape != nil {
		if plant.RotatedRhombusGridShape != plantOther.RotatedRhombusGridShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "RotatedRhombusGridShape"))
		}
	}
	if (plant.GrowthPathRhombusGridShape == nil) != (plantOther.GrowthPathRhombusGridShape == nil) {
		diffs = append(diffs, plant.GongMarshallField(stage, "GrowthPathRhombusGridShape"))
	} else if plant.GrowthPathRhombusGridShape != nil && plantOther.GrowthPathRhombusGridShape != nil {
		if plant.GrowthPathRhombusGridShape != plantOther.GrowthPathRhombusGridShape {
			diffs = append(diffs, plant.GongMarshallField(stage, "GrowthPathRhombusGridShape"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (plantcircumferenceshape *PlantCircumferenceShape) GongDiff(stage *Stage, plantcircumferenceshapeOther *PlantCircumferenceShape) (diffs []string) {
	// insertion point for field diffs
	if plantcircumferenceshape.Name != plantcircumferenceshapeOther.Name {
		diffs = append(diffs, plantcircumferenceshape.GongMarshallField(stage, "Name"))
	}
	if plantcircumferenceshape.AngleDegree != plantcircumferenceshapeOther.AngleDegree {
		diffs = append(diffs, plantcircumferenceshape.GongMarshallField(stage, "AngleDegree"))
	}
	if plantcircumferenceshape.Length != plantcircumferenceshapeOther.Length {
		diffs = append(diffs, plantcircumferenceshape.GongMarshallField(stage, "Length"))
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
	if plantdiagram.IsHiddenAxesShape != plantdiagramOther.IsHiddenAxesShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenAxesShape"))
	}
	if plantdiagram.IsHiddenReferenceRhombus != plantdiagramOther.IsHiddenReferenceRhombus {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenReferenceRhombus"))
	}
	if plantdiagram.IsHiddenPlantCircumferenceShape != plantdiagramOther.IsHiddenPlantCircumferenceShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenPlantCircumferenceShape"))
	}
	if plantdiagram.IsHiddenGridPathShape != plantdiagramOther.IsHiddenGridPathShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenGridPathShape"))
	}
	if plantdiagram.IsHiddenRhombusGridShape != plantdiagramOther.IsHiddenRhombusGridShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenRhombusGridShape"))
	}
	if plantdiagram.IsHiddenExplanationTextShape != plantdiagramOther.IsHiddenExplanationTextShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenExplanationTextShape"))
	}
	if plantdiagram.IsHiddenRotatedReferenceRhombus != plantdiagramOther.IsHiddenRotatedReferenceRhombus {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenRotatedReferenceRhombus"))
	}
	if plantdiagram.IsHiddenRotatedPlantCircumferenceShape != plantdiagramOther.IsHiddenRotatedPlantCircumferenceShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenRotatedPlantCircumferenceShape"))
	}
	if plantdiagram.IsHiddenRotatedGridPathShape != plantdiagramOther.IsHiddenRotatedGridPathShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenRotatedGridPathShape"))
	}
	if plantdiagram.IsHiddenRotatedRhombusGridShape != plantdiagramOther.IsHiddenRotatedRhombusGridShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenRotatedRhombusGridShape"))
	}
	if plantdiagram.IsHiddenGrowthPathRhombusGridShape != plantdiagramOther.IsHiddenGrowthPathRhombusGridShape {
		diffs = append(diffs, plantdiagram.GongMarshallField(stage, "IsHiddenGrowthPathRhombusGridShape"))
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
func (rhombusgridshape *RhombusGridShape) GongDiff(stage *Stage, rhombusgridshapeOther *RhombusGridShape) (diffs []string) {
	// insertion point for field diffs
	if rhombusgridshape.Name != rhombusgridshapeOther.Name {
		diffs = append(diffs, rhombusgridshape.GongMarshallField(stage, "Name"))
	}
	RhombusShapesDifferent := false
	if len(rhombusgridshape.RhombusShapes) != len(rhombusgridshapeOther.RhombusShapes) {
		RhombusShapesDifferent = true
	} else {
		for i := range rhombusgridshape.RhombusShapes {
			if (rhombusgridshape.RhombusShapes[i] == nil) != (rhombusgridshapeOther.RhombusShapes[i] == nil) {
				RhombusShapesDifferent = true
				break
			} else if rhombusgridshape.RhombusShapes[i] != nil && rhombusgridshapeOther.RhombusShapes[i] != nil {
				// this is a pointer comparaison
				if rhombusgridshape.RhombusShapes[i] != rhombusgridshapeOther.RhombusShapes[i] {
					RhombusShapesDifferent = true
					break
				}
			}
		}
	}
	if RhombusShapesDifferent {
		ops := Diff(stage, rhombusgridshape, rhombusgridshapeOther, "RhombusShapes", rhombusgridshapeOther.RhombusShapes, rhombusgridshape.RhombusShapes)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (rhombusshape *RhombusShape) GongDiff(stage *Stage, rhombusshapeOther *RhombusShape) (diffs []string) {
	// insertion point for field diffs
	if rhombusshape.Name != rhombusshapeOther.Name {
		diffs = append(diffs, rhombusshape.GongMarshallField(stage, "Name"))
	}
	if rhombusshape.X != rhombusshapeOther.X {
		diffs = append(diffs, rhombusshape.GongMarshallField(stage, "X"))
	}
	if rhombusshape.Y != rhombusshapeOther.Y {
		diffs = append(diffs, rhombusshape.GongMarshallField(stage, "Y"))
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
