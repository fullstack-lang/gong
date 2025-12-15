// generated code - do not edit
package models

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
