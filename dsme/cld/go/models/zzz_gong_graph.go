// generated code - do not edit
package models

import "fmt"

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Category1:
		ok = stage.IsStagedCategory1(target)

	case *Category1Shape:
		ok = stage.IsStagedCategory1Shape(target)

	case *Category2:
		ok = stage.IsStagedCategory2(target)

	case *Category2Shape:
		ok = stage.IsStagedCategory2Shape(target)

	case *Category3:
		ok = stage.IsStagedCategory3(target)

	case *Category3Shape:
		ok = stage.IsStagedCategory3Shape(target)

	case *ControlPointShape:
		ok = stage.IsStagedControlPointShape(target)

	case *Desk:
		ok = stage.IsStagedDesk(target)

	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Influence:
		ok = stage.IsStagedInfluence(target)

	case *InfluenceShape:
		ok = stage.IsStagedInfluenceShape(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *Category1:
		ok = stage.IsStagedCategory1(target)

	case *Category1Shape:
		ok = stage.IsStagedCategory1Shape(target)

	case *Category2:
		ok = stage.IsStagedCategory2(target)

	case *Category2Shape:
		ok = stage.IsStagedCategory2Shape(target)

	case *Category3:
		ok = stage.IsStagedCategory3(target)

	case *Category3Shape:
		ok = stage.IsStagedCategory3Shape(target)

	case *ControlPointShape:
		ok = stage.IsStagedControlPointShape(target)

	case *Desk:
		ok = stage.IsStagedDesk(target)

	case *Diagram:
		ok = stage.IsStagedDiagram(target)

	case *Influence:
		ok = stage.IsStagedInfluence(target)

	case *InfluenceShape:
		ok = stage.IsStagedInfluenceShape(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedCategory1(category1 *Category1) (ok bool) {

	_, ok = stage.Category1s[category1]

	return
}

func (stage *Stage) IsStagedCategory1Shape(category1shape *Category1Shape) (ok bool) {

	_, ok = stage.Category1Shapes[category1shape]

	return
}

func (stage *Stage) IsStagedCategory2(category2 *Category2) (ok bool) {

	_, ok = stage.Category2s[category2]

	return
}

func (stage *Stage) IsStagedCategory2Shape(category2shape *Category2Shape) (ok bool) {

	_, ok = stage.Category2Shapes[category2shape]

	return
}

func (stage *Stage) IsStagedCategory3(category3 *Category3) (ok bool) {

	_, ok = stage.Category3s[category3]

	return
}

func (stage *Stage) IsStagedCategory3Shape(category3shape *Category3Shape) (ok bool) {

	_, ok = stage.Category3Shapes[category3shape]

	return
}

func (stage *Stage) IsStagedControlPointShape(controlpointshape *ControlPointShape) (ok bool) {

	_, ok = stage.ControlPointShapes[controlpointshape]

	return
}

func (stage *Stage) IsStagedDesk(desk *Desk) (ok bool) {

	_, ok = stage.Desks[desk]

	return
}

func (stage *Stage) IsStagedDiagram(diagram *Diagram) (ok bool) {

	_, ok = stage.Diagrams[diagram]

	return
}

func (stage *Stage) IsStagedInfluence(influence *Influence) (ok bool) {

	_, ok = stage.Influences[influence]

	return
}

func (stage *Stage) IsStagedInfluenceShape(influenceshape *InfluenceShape) (ok bool) {

	_, ok = stage.InfluenceShapes[influenceshape]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *Category1:
		stage.StageBranchCategory1(target)

	case *Category1Shape:
		stage.StageBranchCategory1Shape(target)

	case *Category2:
		stage.StageBranchCategory2(target)

	case *Category2Shape:
		stage.StageBranchCategory2Shape(target)

	case *Category3:
		stage.StageBranchCategory3(target)

	case *Category3Shape:
		stage.StageBranchCategory3Shape(target)

	case *ControlPointShape:
		stage.StageBranchControlPointShape(target)

	case *Desk:
		stage.StageBranchDesk(target)

	case *Diagram:
		stage.StageBranchDiagram(target)

	case *Influence:
		stage.StageBranchInfluence(target)

	case *InfluenceShape:
		stage.StageBranchInfluenceShape(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchCategory1(category1 *Category1) {

	// check if instance is already staged
	if IsStaged(stage, category1) {
		return
	}

	category1.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCategory1Shape(category1shape *Category1Shape) {

	// check if instance is already staged
	if IsStaged(stage, category1shape) {
		return
	}

	category1shape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if category1shape.Category1 != nil {
		StageBranch(stage, category1shape.Category1)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCategory2(category2 *Category2) {

	// check if instance is already staged
	if IsStaged(stage, category2) {
		return
	}

	category2.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCategory2Shape(category2shape *Category2Shape) {

	// check if instance is already staged
	if IsStaged(stage, category2shape) {
		return
	}

	category2shape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if category2shape.Category2 != nil {
		StageBranch(stage, category2shape.Category2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCategory3(category3 *Category3) {

	// check if instance is already staged
	if IsStaged(stage, category3) {
		return
	}

	category3.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchCategory3Shape(category3shape *Category3Shape) {

	// check if instance is already staged
	if IsStaged(stage, category3shape) {
		return
	}

	category3shape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if category3shape.Category3 != nil {
		StageBranch(stage, category3shape.Category3)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchControlPointShape(controlpointshape *ControlPointShape) {

	// check if instance is already staged
	if IsStaged(stage, controlpointshape) {
		return
	}

	controlpointshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDesk(desk *Desk) {

	// check if instance is already staged
	if IsStaged(stage, desk) {
		return
	}

	desk.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if desk.SelectedDiagram != nil {
		StageBranch(stage, desk.SelectedDiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if IsStaged(stage, diagram) {
		return
	}

	diagram.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _category1shape := range diagram.Category1Shapes {
		StageBranch(stage, _category1shape)
	}
	for _, _category2shape := range diagram.Category2Shapes {
		StageBranch(stage, _category2shape)
	}
	for _, _category3shape := range diagram.Category3Shapes {
		StageBranch(stage, _category3shape)
	}
	for _, _influenceshape := range diagram.InfluenceShapes {
		StageBranch(stage, _influenceshape)
	}

}

func (stage *Stage) StageBranchInfluence(influence *Influence) {

	// check if instance is already staged
	if IsStaged(stage, influence) {
		return
	}

	influence.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if influence.SourceCategory1 != nil {
		StageBranch(stage, influence.SourceCategory1)
	}
	if influence.SourceCategory2 != nil {
		StageBranch(stage, influence.SourceCategory2)
	}
	if influence.SourceCategory3 != nil {
		StageBranch(stage, influence.SourceCategory3)
	}
	if influence.TargetCategory1 != nil {
		StageBranch(stage, influence.TargetCategory1)
	}
	if influence.TargetCategory2 != nil {
		StageBranch(stage, influence.TargetCategory2)
	}
	if influence.TargetCategory3 != nil {
		StageBranch(stage, influence.TargetCategory3)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchInfluenceShape(influenceshape *InfluenceShape) {

	// check if instance is already staged
	if IsStaged(stage, influenceshape) {
		return
	}

	influenceshape.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if influenceshape.Influence != nil {
		StageBranch(stage, influenceshape.Influence)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range influenceshape.ControlPointShapes {
		StageBranch(stage, _controlpointshape)
	}

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
	case *Category1:
		toT := CopyBranchCategory1(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Category1Shape:
		toT := CopyBranchCategory1Shape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Category2:
		toT := CopyBranchCategory2(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Category2Shape:
		toT := CopyBranchCategory2Shape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Category3:
		toT := CopyBranchCategory3(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Category3Shape:
		toT := CopyBranchCategory3Shape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ControlPointShape:
		toT := CopyBranchControlPointShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Desk:
		toT := CopyBranchDesk(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Diagram:
		toT := CopyBranchDiagram(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *Influence:
		toT := CopyBranchInfluence(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *InfluenceShape:
		toT := CopyBranchInfluenceShape(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchCategory1(mapOrigCopy map[any]any, category1From *Category1) (category1To *Category1) {

	// category1From has already been copied
	if _category1To, ok := mapOrigCopy[category1From]; ok {
		category1To = _category1To.(*Category1)
		return
	}

	category1To = new(Category1)
	mapOrigCopy[category1From] = category1To
	category1From.CopyBasicFields(category1To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCategory1Shape(mapOrigCopy map[any]any, category1shapeFrom *Category1Shape) (category1shapeTo *Category1Shape) {

	// category1shapeFrom has already been copied
	if _category1shapeTo, ok := mapOrigCopy[category1shapeFrom]; ok {
		category1shapeTo = _category1shapeTo.(*Category1Shape)
		return
	}

	category1shapeTo = new(Category1Shape)
	mapOrigCopy[category1shapeFrom] = category1shapeTo
	category1shapeFrom.CopyBasicFields(category1shapeTo)

	//insertion point for the staging of instances referenced by pointers
	if category1shapeFrom.Category1 != nil {
		category1shapeTo.Category1 = CopyBranchCategory1(mapOrigCopy, category1shapeFrom.Category1)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCategory2(mapOrigCopy map[any]any, category2From *Category2) (category2To *Category2) {

	// category2From has already been copied
	if _category2To, ok := mapOrigCopy[category2From]; ok {
		category2To = _category2To.(*Category2)
		return
	}

	category2To = new(Category2)
	mapOrigCopy[category2From] = category2To
	category2From.CopyBasicFields(category2To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCategory2Shape(mapOrigCopy map[any]any, category2shapeFrom *Category2Shape) (category2shapeTo *Category2Shape) {

	// category2shapeFrom has already been copied
	if _category2shapeTo, ok := mapOrigCopy[category2shapeFrom]; ok {
		category2shapeTo = _category2shapeTo.(*Category2Shape)
		return
	}

	category2shapeTo = new(Category2Shape)
	mapOrigCopy[category2shapeFrom] = category2shapeTo
	category2shapeFrom.CopyBasicFields(category2shapeTo)

	//insertion point for the staging of instances referenced by pointers
	if category2shapeFrom.Category2 != nil {
		category2shapeTo.Category2 = CopyBranchCategory2(mapOrigCopy, category2shapeFrom.Category2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCategory3(mapOrigCopy map[any]any, category3From *Category3) (category3To *Category3) {

	// category3From has already been copied
	if _category3To, ok := mapOrigCopy[category3From]; ok {
		category3To = _category3To.(*Category3)
		return
	}

	category3To = new(Category3)
	mapOrigCopy[category3From] = category3To
	category3From.CopyBasicFields(category3To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchCategory3Shape(mapOrigCopy map[any]any, category3shapeFrom *Category3Shape) (category3shapeTo *Category3Shape) {

	// category3shapeFrom has already been copied
	if _category3shapeTo, ok := mapOrigCopy[category3shapeFrom]; ok {
		category3shapeTo = _category3shapeTo.(*Category3Shape)
		return
	}

	category3shapeTo = new(Category3Shape)
	mapOrigCopy[category3shapeFrom] = category3shapeTo
	category3shapeFrom.CopyBasicFields(category3shapeTo)

	//insertion point for the staging of instances referenced by pointers
	if category3shapeFrom.Category3 != nil {
		category3shapeTo.Category3 = CopyBranchCategory3(mapOrigCopy, category3shapeFrom.Category3)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchControlPointShape(mapOrigCopy map[any]any, controlpointshapeFrom *ControlPointShape) (controlpointshapeTo *ControlPointShape) {

	// controlpointshapeFrom has already been copied
	if _controlpointshapeTo, ok := mapOrigCopy[controlpointshapeFrom]; ok {
		controlpointshapeTo = _controlpointshapeTo.(*ControlPointShape)
		return
	}

	controlpointshapeTo = new(ControlPointShape)
	mapOrigCopy[controlpointshapeFrom] = controlpointshapeTo
	controlpointshapeFrom.CopyBasicFields(controlpointshapeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDesk(mapOrigCopy map[any]any, deskFrom *Desk) (deskTo *Desk) {

	// deskFrom has already been copied
	if _deskTo, ok := mapOrigCopy[deskFrom]; ok {
		deskTo = _deskTo.(*Desk)
		return
	}

	deskTo = new(Desk)
	mapOrigCopy[deskFrom] = deskTo
	deskFrom.CopyBasicFields(deskTo)

	//insertion point for the staging of instances referenced by pointers
	if deskFrom.SelectedDiagram != nil {
		deskTo.SelectedDiagram = CopyBranchDiagram(mapOrigCopy, deskFrom.SelectedDiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchDiagram(mapOrigCopy map[any]any, diagramFrom *Diagram) (diagramTo *Diagram) {

	// diagramFrom has already been copied
	if _diagramTo, ok := mapOrigCopy[diagramFrom]; ok {
		diagramTo = _diagramTo.(*Diagram)
		return
	}

	diagramTo = new(Diagram)
	mapOrigCopy[diagramFrom] = diagramTo
	diagramFrom.CopyBasicFields(diagramTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _category1shape := range diagramFrom.Category1Shapes {
		diagramTo.Category1Shapes = append(diagramTo.Category1Shapes, CopyBranchCategory1Shape(mapOrigCopy, _category1shape))
	}
	for _, _category2shape := range diagramFrom.Category2Shapes {
		diagramTo.Category2Shapes = append(diagramTo.Category2Shapes, CopyBranchCategory2Shape(mapOrigCopy, _category2shape))
	}
	for _, _category3shape := range diagramFrom.Category3Shapes {
		diagramTo.Category3Shapes = append(diagramTo.Category3Shapes, CopyBranchCategory3Shape(mapOrigCopy, _category3shape))
	}
	for _, _influenceshape := range diagramFrom.InfluenceShapes {
		diagramTo.InfluenceShapes = append(diagramTo.InfluenceShapes, CopyBranchInfluenceShape(mapOrigCopy, _influenceshape))
	}

	return
}

func CopyBranchInfluence(mapOrigCopy map[any]any, influenceFrom *Influence) (influenceTo *Influence) {

	// influenceFrom has already been copied
	if _influenceTo, ok := mapOrigCopy[influenceFrom]; ok {
		influenceTo = _influenceTo.(*Influence)
		return
	}

	influenceTo = new(Influence)
	mapOrigCopy[influenceFrom] = influenceTo
	influenceFrom.CopyBasicFields(influenceTo)

	//insertion point for the staging of instances referenced by pointers
	if influenceFrom.SourceCategory1 != nil {
		influenceTo.SourceCategory1 = CopyBranchCategory1(mapOrigCopy, influenceFrom.SourceCategory1)
	}
	if influenceFrom.SourceCategory2 != nil {
		influenceTo.SourceCategory2 = CopyBranchCategory3(mapOrigCopy, influenceFrom.SourceCategory2)
	}
	if influenceFrom.SourceCategory3 != nil {
		influenceTo.SourceCategory3 = CopyBranchCategory2(mapOrigCopy, influenceFrom.SourceCategory3)
	}
	if influenceFrom.TargetCategory1 != nil {
		influenceTo.TargetCategory1 = CopyBranchCategory1(mapOrigCopy, influenceFrom.TargetCategory1)
	}
	if influenceFrom.TargetCategory2 != nil {
		influenceTo.TargetCategory2 = CopyBranchCategory3(mapOrigCopy, influenceFrom.TargetCategory2)
	}
	if influenceFrom.TargetCategory3 != nil {
		influenceTo.TargetCategory3 = CopyBranchCategory2(mapOrigCopy, influenceFrom.TargetCategory3)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchInfluenceShape(mapOrigCopy map[any]any, influenceshapeFrom *InfluenceShape) (influenceshapeTo *InfluenceShape) {

	// influenceshapeFrom has already been copied
	if _influenceshapeTo, ok := mapOrigCopy[influenceshapeFrom]; ok {
		influenceshapeTo = _influenceshapeTo.(*InfluenceShape)
		return
	}

	influenceshapeTo = new(InfluenceShape)
	mapOrigCopy[influenceshapeFrom] = influenceshapeTo
	influenceshapeFrom.CopyBasicFields(influenceshapeTo)

	//insertion point for the staging of instances referenced by pointers
	if influenceshapeFrom.Influence != nil {
		influenceshapeTo.Influence = CopyBranchInfluence(mapOrigCopy, influenceshapeFrom.Influence)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range influenceshapeFrom.ControlPointShapes {
		influenceshapeTo.ControlPointShapes = append(influenceshapeTo.ControlPointShapes, CopyBranchControlPointShape(mapOrigCopy, _controlpointshape))
	}

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
func UnstageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for unstage branch
	case *Category1:
		stage.UnstageBranchCategory1(target)

	case *Category1Shape:
		stage.UnstageBranchCategory1Shape(target)

	case *Category2:
		stage.UnstageBranchCategory2(target)

	case *Category2Shape:
		stage.UnstageBranchCategory2Shape(target)

	case *Category3:
		stage.UnstageBranchCategory3(target)

	case *Category3Shape:
		stage.UnstageBranchCategory3Shape(target)

	case *ControlPointShape:
		stage.UnstageBranchControlPointShape(target)

	case *Desk:
		stage.UnstageBranchDesk(target)

	case *Diagram:
		stage.UnstageBranchDiagram(target)

	case *Influence:
		stage.UnstageBranchInfluence(target)

	case *InfluenceShape:
		stage.UnstageBranchInfluenceShape(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchCategory1(category1 *Category1) {

	// check if instance is already staged
	if !IsStaged(stage, category1) {
		return
	}

	category1.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCategory1Shape(category1shape *Category1Shape) {

	// check if instance is already staged
	if !IsStaged(stage, category1shape) {
		return
	}

	category1shape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if category1shape.Category1 != nil {
		UnstageBranch(stage, category1shape.Category1)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCategory2(category2 *Category2) {

	// check if instance is already staged
	if !IsStaged(stage, category2) {
		return
	}

	category2.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCategory2Shape(category2shape *Category2Shape) {

	// check if instance is already staged
	if !IsStaged(stage, category2shape) {
		return
	}

	category2shape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if category2shape.Category2 != nil {
		UnstageBranch(stage, category2shape.Category2)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCategory3(category3 *Category3) {

	// check if instance is already staged
	if !IsStaged(stage, category3) {
		return
	}

	category3.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchCategory3Shape(category3shape *Category3Shape) {

	// check if instance is already staged
	if !IsStaged(stage, category3shape) {
		return
	}

	category3shape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if category3shape.Category3 != nil {
		UnstageBranch(stage, category3shape.Category3)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchControlPointShape(controlpointshape *ControlPointShape) {

	// check if instance is already staged
	if !IsStaged(stage, controlpointshape) {
		return
	}

	controlpointshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDesk(desk *Desk) {

	// check if instance is already staged
	if !IsStaged(stage, desk) {
		return
	}

	desk.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if desk.SelectedDiagram != nil {
		UnstageBranch(stage, desk.SelectedDiagram)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchDiagram(diagram *Diagram) {

	// check if instance is already staged
	if !IsStaged(stage, diagram) {
		return
	}

	diagram.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _category1shape := range diagram.Category1Shapes {
		UnstageBranch(stage, _category1shape)
	}
	for _, _category2shape := range diagram.Category2Shapes {
		UnstageBranch(stage, _category2shape)
	}
	for _, _category3shape := range diagram.Category3Shapes {
		UnstageBranch(stage, _category3shape)
	}
	for _, _influenceshape := range diagram.InfluenceShapes {
		UnstageBranch(stage, _influenceshape)
	}

}

func (stage *Stage) UnstageBranchInfluence(influence *Influence) {

	// check if instance is already staged
	if !IsStaged(stage, influence) {
		return
	}

	influence.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if influence.SourceCategory1 != nil {
		UnstageBranch(stage, influence.SourceCategory1)
	}
	if influence.SourceCategory2 != nil {
		UnstageBranch(stage, influence.SourceCategory2)
	}
	if influence.SourceCategory3 != nil {
		UnstageBranch(stage, influence.SourceCategory3)
	}
	if influence.TargetCategory1 != nil {
		UnstageBranch(stage, influence.TargetCategory1)
	}
	if influence.TargetCategory2 != nil {
		UnstageBranch(stage, influence.TargetCategory2)
	}
	if influence.TargetCategory3 != nil {
		UnstageBranch(stage, influence.TargetCategory3)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchInfluenceShape(influenceshape *InfluenceShape) {

	// check if instance is already staged
	if !IsStaged(stage, influenceshape) {
		return
	}

	influenceshape.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if influenceshape.Influence != nil {
		UnstageBranch(stage, influenceshape.Influence)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _controlpointshape := range influenceshape.ControlPointShapes {
		UnstageBranch(stage, _controlpointshape)
	}

}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (category1 *Category1) GongDiff(stage *Stage, category1Other *Category1) (diffs []string) {
	// insertion point for field diffs
	if category1.Name != category1Other.Name {
		diffs = append(diffs, category1.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (category1shape *Category1Shape) GongDiff(stage *Stage, category1shapeOther *Category1Shape) (diffs []string) {
	// insertion point for field diffs
	if category1shape.Name != category1shapeOther.Name {
		diffs = append(diffs, category1shape.GongMarshallField(stage, "Name"))
	}
	if (category1shape.Category1 == nil) != (category1shapeOther.Category1 == nil) {
		diffs = append(diffs, category1shape.GongMarshallField(stage, "Category1"))
	} else if category1shape.Category1 != nil && category1shapeOther.Category1 != nil {
		if category1shape.Category1 != category1shapeOther.Category1 {
			diffs = append(diffs, category1shape.GongMarshallField(stage, "Category1"))
		}
	}
	if category1shape.X != category1shapeOther.X {
		diffs = append(diffs, category1shape.GongMarshallField(stage, "X"))
	}
	if category1shape.Y != category1shapeOther.Y {
		diffs = append(diffs, category1shape.GongMarshallField(stage, "Y"))
	}
	if category1shape.Width != category1shapeOther.Width {
		diffs = append(diffs, category1shape.GongMarshallField(stage, "Width"))
	}
	if category1shape.Height != category1shapeOther.Height {
		diffs = append(diffs, category1shape.GongMarshallField(stage, "Height"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (category2 *Category2) GongDiff(stage *Stage, category2Other *Category2) (diffs []string) {
	// insertion point for field diffs
	if category2.Name != category2Other.Name {
		diffs = append(diffs, category2.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (category2shape *Category2Shape) GongDiff(stage *Stage, category2shapeOther *Category2Shape) (diffs []string) {
	// insertion point for field diffs
	if category2shape.Name != category2shapeOther.Name {
		diffs = append(diffs, category2shape.GongMarshallField(stage, "Name"))
	}
	if (category2shape.Category2 == nil) != (category2shapeOther.Category2 == nil) {
		diffs = append(diffs, category2shape.GongMarshallField(stage, "Category2"))
	} else if category2shape.Category2 != nil && category2shapeOther.Category2 != nil {
		if category2shape.Category2 != category2shapeOther.Category2 {
			diffs = append(diffs, category2shape.GongMarshallField(stage, "Category2"))
		}
	}
	if category2shape.X != category2shapeOther.X {
		diffs = append(diffs, category2shape.GongMarshallField(stage, "X"))
	}
	if category2shape.Y != category2shapeOther.Y {
		diffs = append(diffs, category2shape.GongMarshallField(stage, "Y"))
	}
	if category2shape.Width != category2shapeOther.Width {
		diffs = append(diffs, category2shape.GongMarshallField(stage, "Width"))
	}
	if category2shape.Height != category2shapeOther.Height {
		diffs = append(diffs, category2shape.GongMarshallField(stage, "Height"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (category3 *Category3) GongDiff(stage *Stage, category3Other *Category3) (diffs []string) {
	// insertion point for field diffs
	if category3.Name != category3Other.Name {
		diffs = append(diffs, category3.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (category3shape *Category3Shape) GongDiff(stage *Stage, category3shapeOther *Category3Shape) (diffs []string) {
	// insertion point for field diffs
	if category3shape.Name != category3shapeOther.Name {
		diffs = append(diffs, category3shape.GongMarshallField(stage, "Name"))
	}
	if (category3shape.Category3 == nil) != (category3shapeOther.Category3 == nil) {
		diffs = append(diffs, category3shape.GongMarshallField(stage, "Category3"))
	} else if category3shape.Category3 != nil && category3shapeOther.Category3 != nil {
		if category3shape.Category3 != category3shapeOther.Category3 {
			diffs = append(diffs, category3shape.GongMarshallField(stage, "Category3"))
		}
	}
	if category3shape.X != category3shapeOther.X {
		diffs = append(diffs, category3shape.GongMarshallField(stage, "X"))
	}
	if category3shape.Y != category3shapeOther.Y {
		diffs = append(diffs, category3shape.GongMarshallField(stage, "Y"))
	}
	if category3shape.Width != category3shapeOther.Width {
		diffs = append(diffs, category3shape.GongMarshallField(stage, "Width"))
	}
	if category3shape.Height != category3shapeOther.Height {
		diffs = append(diffs, category3shape.GongMarshallField(stage, "Height"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (controlpointshape *ControlPointShape) GongDiff(stage *Stage, controlpointshapeOther *ControlPointShape) (diffs []string) {
	// insertion point for field diffs
	if controlpointshape.Name != controlpointshapeOther.Name {
		diffs = append(diffs, controlpointshape.GongMarshallField(stage, "Name"))
	}
	if controlpointshape.X_Relative != controlpointshapeOther.X_Relative {
		diffs = append(diffs, controlpointshape.GongMarshallField(stage, "X_Relative"))
	}
	if controlpointshape.Y_Relative != controlpointshapeOther.Y_Relative {
		diffs = append(diffs, controlpointshape.GongMarshallField(stage, "Y_Relative"))
	}
	if controlpointshape.IsStartShapeTheClosestShape != controlpointshapeOther.IsStartShapeTheClosestShape {
		diffs = append(diffs, controlpointshape.GongMarshallField(stage, "IsStartShapeTheClosestShape"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (desk *Desk) GongDiff(stage *Stage, deskOther *Desk) (diffs []string) {
	// insertion point for field diffs
	if desk.Name != deskOther.Name {
		diffs = append(diffs, desk.GongMarshallField(stage, "Name"))
	}
	if (desk.SelectedDiagram == nil) != (deskOther.SelectedDiagram == nil) {
		diffs = append(diffs, desk.GongMarshallField(stage, "SelectedDiagram"))
	} else if desk.SelectedDiagram != nil && deskOther.SelectedDiagram != nil {
		if desk.SelectedDiagram != deskOther.SelectedDiagram {
			diffs = append(diffs, desk.GongMarshallField(stage, "SelectedDiagram"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (diagram *Diagram) GongDiff(stage *Stage, diagramOther *Diagram) (diffs []string) {
	// insertion point for field diffs
	if diagram.Name != diagramOther.Name {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Name"))
	}
	Category1ShapesDifferent := false
	if len(diagram.Category1Shapes) != len(diagramOther.Category1Shapes) {
		Category1ShapesDifferent = true
	} else {
		for i := range diagram.Category1Shapes {
			if (diagram.Category1Shapes[i] == nil) != (diagramOther.Category1Shapes[i] == nil) {
				Category1ShapesDifferent = true
				break
			} else if diagram.Category1Shapes[i] != nil && diagramOther.Category1Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.Category1Shapes[i] != diagramOther.Category1Shapes[i] {
					Category1ShapesDifferent = true
					break
				}
			}
		}
	}
	if Category1ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "Category1Shapes", diagramOther.Category1Shapes, diagram.Category1Shapes)
		diffs = append(diffs, ops)
	}
	Category2ShapesDifferent := false
	if len(diagram.Category2Shapes) != len(diagramOther.Category2Shapes) {
		Category2ShapesDifferent = true
	} else {
		for i := range diagram.Category2Shapes {
			if (diagram.Category2Shapes[i] == nil) != (diagramOther.Category2Shapes[i] == nil) {
				Category2ShapesDifferent = true
				break
			} else if diagram.Category2Shapes[i] != nil && diagramOther.Category2Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.Category2Shapes[i] != diagramOther.Category2Shapes[i] {
					Category2ShapesDifferent = true
					break
				}
			}
		}
	}
	if Category2ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "Category2Shapes", diagramOther.Category2Shapes, diagram.Category2Shapes)
		diffs = append(diffs, ops)
	}
	Category3ShapesDifferent := false
	if len(diagram.Category3Shapes) != len(diagramOther.Category3Shapes) {
		Category3ShapesDifferent = true
	} else {
		for i := range diagram.Category3Shapes {
			if (diagram.Category3Shapes[i] == nil) != (diagramOther.Category3Shapes[i] == nil) {
				Category3ShapesDifferent = true
				break
			} else if diagram.Category3Shapes[i] != nil && diagramOther.Category3Shapes[i] != nil {
				// this is a pointer comparaison
				if diagram.Category3Shapes[i] != diagramOther.Category3Shapes[i] {
					Category3ShapesDifferent = true
					break
				}
			}
		}
	}
	if Category3ShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "Category3Shapes", diagramOther.Category3Shapes, diagram.Category3Shapes)
		diffs = append(diffs, ops)
	}
	InfluenceShapesDifferent := false
	if len(diagram.InfluenceShapes) != len(diagramOther.InfluenceShapes) {
		InfluenceShapesDifferent = true
	} else {
		for i := range diagram.InfluenceShapes {
			if (diagram.InfluenceShapes[i] == nil) != (diagramOther.InfluenceShapes[i] == nil) {
				InfluenceShapesDifferent = true
				break
			} else if diagram.InfluenceShapes[i] != nil && diagramOther.InfluenceShapes[i] != nil {
				// this is a pointer comparaison
				if diagram.InfluenceShapes[i] != diagramOther.InfluenceShapes[i] {
					InfluenceShapesDifferent = true
					break
				}
			}
		}
	}
	if InfluenceShapesDifferent {
		ops := Diff(stage, diagram, diagramOther, "InfluenceShapes", diagramOther.InfluenceShapes, diagram.InfluenceShapes)
		diffs = append(diffs, ops)
	}
	if diagram.IsEditable != diagramOther.IsEditable {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsEditable"))
	}
	if diagram.IsNodeExpanded != diagramOther.IsNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsNodeExpanded"))
	}
	if diagram.IsCategory1NodeExpanded != diagramOther.IsCategory1NodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsCategory1NodeExpanded"))
	}
	if diagram.IsCategory2NodeExpanded != diagramOther.IsCategory2NodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsCategory2NodeExpanded"))
	}
	if diagram.IsCategory3NodeExpanded != diagramOther.IsCategory3NodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsCategory3NodeExpanded"))
	}
	if diagram.IsInfluenceCategoryNodeExpanded != diagramOther.IsInfluenceCategoryNodeExpanded {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsInfluenceCategoryNodeExpanded"))
	}
	if diagram.IsCategory1Shown != diagramOther.IsCategory1Shown {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsCategory1Shown"))
	}
	if diagram.IsCategory2Shown != diagramOther.IsCategory2Shown {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsCategory2Shown"))
	}
	if diagram.IsCategory3Shown != diagramOther.IsCategory3Shown {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsCategory3Shown"))
	}
	if diagram.IsInfluenceCategoryShown != diagramOther.IsInfluenceCategoryShown {
		diffs = append(diffs, diagram.GongMarshallField(stage, "IsInfluenceCategoryShown"))
	}
	if diagram.XMargin != diagramOther.XMargin {
		diffs = append(diffs, diagram.GongMarshallField(stage, "XMargin"))
	}
	if diagram.YMargin != diagramOther.YMargin {
		diffs = append(diffs, diagram.GongMarshallField(stage, "YMargin"))
	}
	if diagram.Height != diagramOther.Height {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Height"))
	}
	if diagram.Width != diagramOther.Width {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Width"))
	}
	if diagram.NbPixPerCharacter != diagramOther.NbPixPerCharacter {
		diffs = append(diffs, diagram.GongMarshallField(stage, "NbPixPerCharacter"))
	}
	if diagram.RedColorCode != diagramOther.RedColorCode {
		diffs = append(diffs, diagram.GongMarshallField(stage, "RedColorCode"))
	}
	if diagram.BackgroundGreyColorCode != diagramOther.BackgroundGreyColorCode {
		diffs = append(diffs, diagram.GongMarshallField(stage, "BackgroundGreyColorCode"))
	}
	if diagram.GrayColorCode != diagramOther.GrayColorCode {
		diffs = append(diffs, diagram.GongMarshallField(stage, "GrayColorCode"))
	}
	if diagram.Category1RectAnchorType != diagramOther.Category1RectAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category1RectAnchorType"))
	}
	if diagram.Category1TextAnchorType != diagramOther.Category1TextAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category1TextAnchorType"))
	}
	if diagram.Category1DominantBaselineType != diagramOther.Category1DominantBaselineType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category1DominantBaselineType"))
	}
	if diagram.Category1FontSize != diagramOther.Category1FontSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category1FontSize"))
	}
	if diagram.Category1FontWeigth != diagramOther.Category1FontWeigth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category1FontWeigth"))
	}
	if diagram.Category1FontFamily != diagramOther.Category1FontFamily {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category1FontFamily"))
	}
	if diagram.Category1LetterSpacing != diagramOther.Category1LetterSpacing {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category1LetterSpacing"))
	}
	if diagram.Category2TypeFontSize != diagramOther.Category2TypeFontSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category2TypeFontSize"))
	}
	if diagram.Category2TypeFontWeigth != diagramOther.Category2TypeFontWeigth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category2TypeFontWeigth"))
	}
	if diagram.Category2TypeFontFamily != diagramOther.Category2TypeFontFamily {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category2TypeFontFamily"))
	}
	if diagram.Category2TypeLetterSpacing != diagramOther.Category2TypeLetterSpacing {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category2TypeLetterSpacing"))
	}
	if diagram.Category2TypeRectAnchorType != diagramOther.Category2TypeRectAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category2TypeRectAnchorType"))
	}
	if diagram.Category2DominantBaselineType != diagramOther.Category2DominantBaselineType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category2DominantBaselineType"))
	}
	if diagram.Category2StrokeWidth != diagramOther.Category2StrokeWidth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category2StrokeWidth"))
	}
	if diagram.Category3RectAnchorType != diagramOther.Category3RectAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category3RectAnchorType"))
	}
	if diagram.Category3TextAnchorType != diagramOther.Category3TextAnchorType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category3TextAnchorType"))
	}
	if diagram.Category3DominantBaselineType != diagramOther.Category3DominantBaselineType {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category3DominantBaselineType"))
	}
	if diagram.Category3FontSize != diagramOther.Category3FontSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category3FontSize"))
	}
	if diagram.Category3FontWeigth != diagramOther.Category3FontWeigth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category3FontWeigth"))
	}
	if diagram.Category3FontFamily != diagramOther.Category3FontFamily {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category3FontFamily"))
	}
	if diagram.Category3LetterSpacing != diagramOther.Category3LetterSpacing {
		diffs = append(diffs, diagram.GongMarshallField(stage, "Category3LetterSpacing"))
	}
	if diagram.InfluenceStrokeWidth != diagramOther.InfluenceStrokeWidth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "InfluenceStrokeWidth"))
	}
	if diagram.InfluenceArrowSize != diagramOther.InfluenceArrowSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "InfluenceArrowSize"))
	}
	if diagram.InfluenceArrowStartOffset != diagramOther.InfluenceArrowStartOffset {
		diffs = append(diffs, diagram.GongMarshallField(stage, "InfluenceArrowStartOffset"))
	}
	if diagram.InfluenceArrowEndOffset != diagramOther.InfluenceArrowEndOffset {
		diffs = append(diffs, diagram.GongMarshallField(stage, "InfluenceArrowEndOffset"))
	}
	if diagram.InfluenceCornerRadius != diagramOther.InfluenceCornerRadius {
		diffs = append(diffs, diagram.GongMarshallField(stage, "InfluenceCornerRadius"))
	}
	if diagram.InfluenceFontSize != diagramOther.InfluenceFontSize {
		diffs = append(diffs, diagram.GongMarshallField(stage, "InfluenceFontSize"))
	}
	if diagram.InfluenceFontWeigth != diagramOther.InfluenceFontWeigth {
		diffs = append(diffs, diagram.GongMarshallField(stage, "InfluenceFontWeigth"))
	}
	if diagram.InfluenceFontFamily != diagramOther.InfluenceFontFamily {
		diffs = append(diffs, diagram.GongMarshallField(stage, "InfluenceFontFamily"))
	}
	if diagram.InfluenceLetterSpacing != diagramOther.InfluenceLetterSpacing {
		diffs = append(diffs, diagram.GongMarshallField(stage, "InfluenceLetterSpacing"))
	}
	if diagram.InfluenceDashedLinePattern != diagramOther.InfluenceDashedLinePattern {
		diffs = append(diffs, diagram.GongMarshallField(stage, "InfluenceDashedLinePattern"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (influence *Influence) GongDiff(stage *Stage, influenceOther *Influence) (diffs []string) {
	// insertion point for field diffs
	if influence.Name != influenceOther.Name {
		diffs = append(diffs, influence.GongMarshallField(stage, "Name"))
	}
	if (influence.SourceCategory1 == nil) != (influenceOther.SourceCategory1 == nil) {
		diffs = append(diffs, influence.GongMarshallField(stage, "SourceCategory1"))
	} else if influence.SourceCategory1 != nil && influenceOther.SourceCategory1 != nil {
		if influence.SourceCategory1 != influenceOther.SourceCategory1 {
			diffs = append(diffs, influence.GongMarshallField(stage, "SourceCategory1"))
		}
	}
	if (influence.SourceCategory2 == nil) != (influenceOther.SourceCategory2 == nil) {
		diffs = append(diffs, influence.GongMarshallField(stage, "SourceCategory2"))
	} else if influence.SourceCategory2 != nil && influenceOther.SourceCategory2 != nil {
		if influence.SourceCategory2 != influenceOther.SourceCategory2 {
			diffs = append(diffs, influence.GongMarshallField(stage, "SourceCategory2"))
		}
	}
	if (influence.SourceCategory3 == nil) != (influenceOther.SourceCategory3 == nil) {
		diffs = append(diffs, influence.GongMarshallField(stage, "SourceCategory3"))
	} else if influence.SourceCategory3 != nil && influenceOther.SourceCategory3 != nil {
		if influence.SourceCategory3 != influenceOther.SourceCategory3 {
			diffs = append(diffs, influence.GongMarshallField(stage, "SourceCategory3"))
		}
	}
	if (influence.TargetCategory1 == nil) != (influenceOther.TargetCategory1 == nil) {
		diffs = append(diffs, influence.GongMarshallField(stage, "TargetCategory1"))
	} else if influence.TargetCategory1 != nil && influenceOther.TargetCategory1 != nil {
		if influence.TargetCategory1 != influenceOther.TargetCategory1 {
			diffs = append(diffs, influence.GongMarshallField(stage, "TargetCategory1"))
		}
	}
	if (influence.TargetCategory2 == nil) != (influenceOther.TargetCategory2 == nil) {
		diffs = append(diffs, influence.GongMarshallField(stage, "TargetCategory2"))
	} else if influence.TargetCategory2 != nil && influenceOther.TargetCategory2 != nil {
		if influence.TargetCategory2 != influenceOther.TargetCategory2 {
			diffs = append(diffs, influence.GongMarshallField(stage, "TargetCategory2"))
		}
	}
	if (influence.TargetCategory3 == nil) != (influenceOther.TargetCategory3 == nil) {
		diffs = append(diffs, influence.GongMarshallField(stage, "TargetCategory3"))
	} else if influence.TargetCategory3 != nil && influenceOther.TargetCategory3 != nil {
		if influence.TargetCategory3 != influenceOther.TargetCategory3 {
			diffs = append(diffs, influence.GongMarshallField(stage, "TargetCategory3"))
		}
	}
	if influence.IsHypothtical != influenceOther.IsHypothtical {
		diffs = append(diffs, influence.GongMarshallField(stage, "IsHypothtical"))
	}
	if influence.TextAtEndOfArrow != influenceOther.TextAtEndOfArrow {
		diffs = append(diffs, influence.GongMarshallField(stage, "TextAtEndOfArrow"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (influenceshape *InfluenceShape) GongDiff(stage *Stage, influenceshapeOther *InfluenceShape) (diffs []string) {
	// insertion point for field diffs
	if influenceshape.Name != influenceshapeOther.Name {
		diffs = append(diffs, influenceshape.GongMarshallField(stage, "Name"))
	}
	if (influenceshape.Influence == nil) != (influenceshapeOther.Influence == nil) {
		diffs = append(diffs, influenceshape.GongMarshallField(stage, "Influence"))
	} else if influenceshape.Influence != nil && influenceshapeOther.Influence != nil {
		if influenceshape.Influence != influenceshapeOther.Influence {
			diffs = append(diffs, influenceshape.GongMarshallField(stage, "Influence"))
		}
	}
	ControlPointShapesDifferent := false
	if len(influenceshape.ControlPointShapes) != len(influenceshapeOther.ControlPointShapes) {
		ControlPointShapesDifferent = true
	} else {
		for i := range influenceshape.ControlPointShapes {
			if (influenceshape.ControlPointShapes[i] == nil) != (influenceshapeOther.ControlPointShapes[i] == nil) {
				ControlPointShapesDifferent = true
				break
			} else if influenceshape.ControlPointShapes[i] != nil && influenceshapeOther.ControlPointShapes[i] != nil {
				// this is a pointer comparaison
				if influenceshape.ControlPointShapes[i] != influenceshapeOther.ControlPointShapes[i] {
					ControlPointShapesDifferent = true
					break
				}
			}
		}
	}
	if ControlPointShapesDifferent {
		ops := Diff(stage, influenceshape, influenceshapeOther, "ControlPointShapes", influenceshapeOther.ControlPointShapes, influenceshape.ControlPointShapes)
		diffs = append(diffs, ops)
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
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, k+1)
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
