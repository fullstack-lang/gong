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
func (gongbasicfield *GongBasicField) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.GongBasicFields[gongbasicfield]

	return
}

func (stage *Stage) IsStagedGongBasicField(gongbasicfield *GongBasicField) (ok bool) {

	return gongbasicfield.GongIsStaged(stage)
}

func (gongenum *GongEnum) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.GongEnums[gongenum]

	return
}

func (stage *Stage) IsStagedGongEnum(gongenum *GongEnum) (ok bool) {

	return gongenum.GongIsStaged(stage)
}

func (gongenumvalue *GongEnumValue) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.GongEnumValues[gongenumvalue]

	return
}

func (stage *Stage) IsStagedGongEnumValue(gongenumvalue *GongEnumValue) (ok bool) {

	return gongenumvalue.GongIsStaged(stage)
}

func (gonglink *GongLink) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.GongLinks[gonglink]

	return
}

func (stage *Stage) IsStagedGongLink(gonglink *GongLink) (ok bool) {

	return gonglink.GongIsStaged(stage)
}

func (gongnote *GongNote) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.GongNotes[gongnote]

	return
}

func (stage *Stage) IsStagedGongNote(gongnote *GongNote) (ok bool) {

	return gongnote.GongIsStaged(stage)
}

func (gongstruct *GongStruct) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.GongStructs[gongstruct]

	return
}

func (stage *Stage) IsStagedGongStruct(gongstruct *GongStruct) (ok bool) {

	return gongstruct.GongIsStaged(stage)
}

func (gongtimefield *GongTimeField) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.GongTimeFields[gongtimefield]

	return
}

func (stage *Stage) IsStagedGongTimeField(gongtimefield *GongTimeField) (ok bool) {

	return gongtimefield.GongIsStaged(stage)
}

func (metareference *MetaReference) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.MetaReferences[metareference]

	return
}

func (stage *Stage) IsStagedMetaReference(metareference *MetaReference) (ok bool) {

	return metareference.GongIsStaged(stage)
}

func (modelpkg *ModelPkg) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ModelPkgs[modelpkg]

	return
}

func (stage *Stage) IsStagedModelPkg(modelpkg *ModelPkg) (ok bool) {

	return modelpkg.GongIsStaged(stage)
}

func (pointertogongstructfield *PointerToGongStructField) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.PointerToGongStructFields[pointertogongstructfield]

	return
}

func (stage *Stage) IsStagedPointerToGongStructField(pointertogongstructfield *PointerToGongStructField) (ok bool) {

	return pointertogongstructfield.GongIsStaged(stage)
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield]

	return
}

func (stage *Stage) IsStagedSliceOfPointerToGongStructField(sliceofpointertogongstructfield *SliceOfPointerToGongStructField) (ok bool) {

	return sliceofpointertogongstructfield.GongIsStaged(stage)
}

func (stagesetfield *StageSetField) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StageSetFields[stagesetfield]

	return
}

func (stage *Stage) IsStagedStageSetField(stagesetfield *StageSetField) (ok bool) {

	return stagesetfield.GongIsStaged(stage)
}

func (stagesetmodel *StageSetModel) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.StageSetModels[stagesetmodel]

	return
}

func (stage *Stage) IsStagedStageSetModel(stagesetmodel *StageSetModel) (ok bool) {

	return stagesetmodel.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (gongbasicfield *GongBasicField) GongStageBranch(stage *Stage) {
	stage.StageBranchGongBasicField(gongbasicfield)
}

func (stage *Stage) StageBranchGongBasicField(gongbasicfield *GongBasicField) {

	// check if instance is already staged
	if stage.IsStaged(gongbasicfield) {
		return
	}

	gongbasicfield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if gongbasicfield.GongEnum != nil {
		stage.StageBranch(gongbasicfield.GongEnum)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gongenum *GongEnum) GongStageBranch(stage *Stage) {
	stage.StageBranchGongEnum(gongenum)
}

func (stage *Stage) StageBranchGongEnum(gongenum *GongEnum) {

	// check if instance is already staged
	if stage.IsStaged(gongenum) {
		return
	}

	gongenum.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalue := range gongenum.GongEnumValues {
		stage.StageBranch(_gongenumvalue)
	}

}

func (gongenumvalue *GongEnumValue) GongStageBranch(stage *Stage) {
	stage.StageBranchGongEnumValue(gongenumvalue)
}

func (stage *Stage) StageBranchGongEnumValue(gongenumvalue *GongEnumValue) {

	// check if instance is already staged
	if stage.IsStaged(gongenumvalue) {
		return
	}

	gongenumvalue.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gonglink *GongLink) GongStageBranch(stage *Stage) {
	stage.StageBranchGongLink(gonglink)
}

func (stage *Stage) StageBranchGongLink(gonglink *GongLink) {

	// check if instance is already staged
	if stage.IsStaged(gonglink) {
		return
	}

	gonglink.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gongnote *GongNote) GongStageBranch(stage *Stage) {
	stage.StageBranchGongNote(gongnote)
}

func (stage *Stage) StageBranchGongNote(gongnote *GongNote) {

	// check if instance is already staged
	if stage.IsStaged(gongnote) {
		return
	}

	gongnote.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gonglink := range gongnote.Links {
		stage.StageBranch(_gonglink)
	}

}

func (gongstruct *GongStruct) GongStageBranch(stage *Stage) {
	stage.StageBranchGongStruct(gongstruct)
}

func (stage *Stage) StageBranchGongStruct(gongstruct *GongStruct) {

	// check if instance is already staged
	if stage.IsStaged(gongstruct) {
		return
	}

	gongstruct.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongbasicfield := range gongstruct.GongBasicFields {
		stage.StageBranch(_gongbasicfield)
	}
	for _, _gongtimefield := range gongstruct.GongTimeFields {
		stage.StageBranch(_gongtimefield)
	}
	for _, _pointertogongstructfield := range gongstruct.PointerToGongStructFields {
		stage.StageBranch(_pointertogongstructfield)
	}
	for _, _sliceofpointertogongstructfield := range gongstruct.SliceOfPointerToGongStructFields {
		stage.StageBranch(_sliceofpointertogongstructfield)
	}

}

func (gongtimefield *GongTimeField) GongStageBranch(stage *Stage) {
	stage.StageBranchGongTimeField(gongtimefield)
}

func (stage *Stage) StageBranchGongTimeField(gongtimefield *GongTimeField) {

	// check if instance is already staged
	if stage.IsStaged(gongtimefield) {
		return
	}

	gongtimefield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metareference *MetaReference) GongStageBranch(stage *Stage) {
	stage.StageBranchMetaReference(metareference)
}

func (stage *Stage) StageBranchMetaReference(metareference *MetaReference) {

	// check if instance is already staged
	if stage.IsStaged(metareference) {
		return
	}

	metareference.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (modelpkg *ModelPkg) GongStageBranch(stage *Stage) {
	stage.StageBranchModelPkg(modelpkg)
}

func (stage *Stage) StageBranchModelPkg(modelpkg *ModelPkg) {

	// check if instance is already staged
	if stage.IsStaged(modelpkg) {
		return
	}

	modelpkg.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pointertogongstructfield *PointerToGongStructField) GongStageBranch(stage *Stage) {
	stage.StageBranchPointerToGongStructField(pointertogongstructfield)
}

func (stage *Stage) StageBranchPointerToGongStructField(pointertogongstructfield *PointerToGongStructField) {

	// check if instance is already staged
	if stage.IsStaged(pointertogongstructfield) {
		return
	}

	pointertogongstructfield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pointertogongstructfield.GongStruct != nil {
		stage.StageBranch(pointertogongstructfield.GongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongStageBranch(stage *Stage) {
	stage.StageBranchSliceOfPointerToGongStructField(sliceofpointertogongstructfield)
}

func (stage *Stage) StageBranchSliceOfPointerToGongStructField(sliceofpointertogongstructfield *SliceOfPointerToGongStructField) {

	// check if instance is already staged
	if stage.IsStaged(sliceofpointertogongstructfield) {
		return
	}

	sliceofpointertogongstructfield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sliceofpointertogongstructfield.GongStruct != nil {
		stage.StageBranch(sliceofpointertogongstructfield.GongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stagesetfield *StageSetField) GongStageBranch(stage *Stage) {
	stage.StageBranchStageSetField(stagesetfield)
}

func (stage *Stage) StageBranchStageSetField(stagesetfield *StageSetField) {

	// check if instance is already staged
	if stage.IsStaged(stagesetfield) {
		return
	}

	stagesetfield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stagesetmodel *StageSetModel) GongStageBranch(stage *Stage) {
	stage.StageBranchStageSetModel(stagesetmodel)
}

func (stage *Stage) StageBranchStageSetModel(stagesetmodel *StageSetModel) {

	// check if instance is already staged
	if stage.IsStaged(stagesetmodel) {
		return
	}

	stagesetmodel.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _stagesetfield := range stagesetmodel.Fields {
		stage.StageBranch(_stagesetfield)
	}

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
	case *GongBasicField:
		toT := GongCopyBranchGongBasicField(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongEnum:
		toT := GongCopyBranchGongEnum(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongEnumValue:
		toT := GongCopyBranchGongEnumValue(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongLink:
		toT := GongCopyBranchGongLink(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongNote:
		toT := GongCopyBranchGongNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongStruct:
		toT := GongCopyBranchGongStruct(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongTimeField:
		toT := GongCopyBranchGongTimeField(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MetaReference:
		toT := GongCopyBranchMetaReference(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ModelPkg:
		toT := GongCopyBranchModelPkg(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PointerToGongStructField:
		toT := GongCopyBranchPointerToGongStructField(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SliceOfPointerToGongStructField:
		toT := GongCopyBranchSliceOfPointerToGongStructField(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StageSetField:
		toT := GongCopyBranchStageSetField(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *StageSetModel:
		toT := GongCopyBranchStageSetModel(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchGongBasicField(mapOrigCopy map[any]any, gongbasicfieldFrom *GongBasicField) (gongbasicfieldTo *GongBasicField) {

	// gongbasicfieldFrom has already been copied
	if _gongbasicfieldTo, ok := mapOrigCopy[gongbasicfieldFrom]; ok {
		gongbasicfieldTo = _gongbasicfieldTo.(*GongBasicField)
		return
	}

	gongbasicfieldTo = new(GongBasicField)
	mapOrigCopy[gongbasicfieldFrom] = gongbasicfieldTo
	gongbasicfieldFrom.GongCopyBasicFields(gongbasicfieldTo)

	//insertion point for the staging of instances referenced by pointers
	if gongbasicfieldFrom.GongEnum != nil {
		gongbasicfieldTo.GongEnum = GongCopyBranchGongEnum(mapOrigCopy, gongbasicfieldFrom.GongEnum)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGongEnum(mapOrigCopy map[any]any, gongenumFrom *GongEnum) (gongenumTo *GongEnum) {

	// gongenumFrom has already been copied
	if _gongenumTo, ok := mapOrigCopy[gongenumFrom]; ok {
		gongenumTo = _gongenumTo.(*GongEnum)
		return
	}

	gongenumTo = new(GongEnum)
	mapOrigCopy[gongenumFrom] = gongenumTo
	gongenumFrom.GongCopyBasicFields(gongenumTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalue := range gongenumFrom.GongEnumValues {
		gongenumTo.GongEnumValues = append(gongenumTo.GongEnumValues, GongCopyBranchGongEnumValue(mapOrigCopy, _gongenumvalue))
	}

	return
}

func GongCopyBranchGongEnumValue(mapOrigCopy map[any]any, gongenumvalueFrom *GongEnumValue) (gongenumvalueTo *GongEnumValue) {

	// gongenumvalueFrom has already been copied
	if _gongenumvalueTo, ok := mapOrigCopy[gongenumvalueFrom]; ok {
		gongenumvalueTo = _gongenumvalueTo.(*GongEnumValue)
		return
	}

	gongenumvalueTo = new(GongEnumValue)
	mapOrigCopy[gongenumvalueFrom] = gongenumvalueTo
	gongenumvalueFrom.GongCopyBasicFields(gongenumvalueTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGongLink(mapOrigCopy map[any]any, gonglinkFrom *GongLink) (gonglinkTo *GongLink) {

	// gonglinkFrom has already been copied
	if _gonglinkTo, ok := mapOrigCopy[gonglinkFrom]; ok {
		gonglinkTo = _gonglinkTo.(*GongLink)
		return
	}

	gonglinkTo = new(GongLink)
	mapOrigCopy[gonglinkFrom] = gonglinkTo
	gonglinkFrom.GongCopyBasicFields(gonglinkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGongNote(mapOrigCopy map[any]any, gongnoteFrom *GongNote) (gongnoteTo *GongNote) {

	// gongnoteFrom has already been copied
	if _gongnoteTo, ok := mapOrigCopy[gongnoteFrom]; ok {
		gongnoteTo = _gongnoteTo.(*GongNote)
		return
	}

	gongnoteTo = new(GongNote)
	mapOrigCopy[gongnoteFrom] = gongnoteTo
	gongnoteFrom.GongCopyBasicFields(gongnoteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gonglink := range gongnoteFrom.Links {
		gongnoteTo.Links = append(gongnoteTo.Links, GongCopyBranchGongLink(mapOrigCopy, _gonglink))
	}

	return
}

func GongCopyBranchGongStruct(mapOrigCopy map[any]any, gongstructFrom *GongStruct) (gongstructTo *GongStruct) {

	// gongstructFrom has already been copied
	if _gongstructTo, ok := mapOrigCopy[gongstructFrom]; ok {
		gongstructTo = _gongstructTo.(*GongStruct)
		return
	}

	gongstructTo = new(GongStruct)
	mapOrigCopy[gongstructFrom] = gongstructTo
	gongstructFrom.GongCopyBasicFields(gongstructTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongbasicfield := range gongstructFrom.GongBasicFields {
		gongstructTo.GongBasicFields = append(gongstructTo.GongBasicFields, GongCopyBranchGongBasicField(mapOrigCopy, _gongbasicfield))
	}
	for _, _gongtimefield := range gongstructFrom.GongTimeFields {
		gongstructTo.GongTimeFields = append(gongstructTo.GongTimeFields, GongCopyBranchGongTimeField(mapOrigCopy, _gongtimefield))
	}
	for _, _pointertogongstructfield := range gongstructFrom.PointerToGongStructFields {
		gongstructTo.PointerToGongStructFields = append(gongstructTo.PointerToGongStructFields, GongCopyBranchPointerToGongStructField(mapOrigCopy, _pointertogongstructfield))
	}
	for _, _sliceofpointertogongstructfield := range gongstructFrom.SliceOfPointerToGongStructFields {
		gongstructTo.SliceOfPointerToGongStructFields = append(gongstructTo.SliceOfPointerToGongStructFields, GongCopyBranchSliceOfPointerToGongStructField(mapOrigCopy, _sliceofpointertogongstructfield))
	}

	return
}

func GongCopyBranchGongTimeField(mapOrigCopy map[any]any, gongtimefieldFrom *GongTimeField) (gongtimefieldTo *GongTimeField) {

	// gongtimefieldFrom has already been copied
	if _gongtimefieldTo, ok := mapOrigCopy[gongtimefieldFrom]; ok {
		gongtimefieldTo = _gongtimefieldTo.(*GongTimeField)
		return
	}

	gongtimefieldTo = new(GongTimeField)
	mapOrigCopy[gongtimefieldFrom] = gongtimefieldTo
	gongtimefieldFrom.GongCopyBasicFields(gongtimefieldTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMetaReference(mapOrigCopy map[any]any, metareferenceFrom *MetaReference) (metareferenceTo *MetaReference) {

	// metareferenceFrom has already been copied
	if _metareferenceTo, ok := mapOrigCopy[metareferenceFrom]; ok {
		metareferenceTo = _metareferenceTo.(*MetaReference)
		return
	}

	metareferenceTo = new(MetaReference)
	mapOrigCopy[metareferenceFrom] = metareferenceTo
	metareferenceFrom.GongCopyBasicFields(metareferenceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchModelPkg(mapOrigCopy map[any]any, modelpkgFrom *ModelPkg) (modelpkgTo *ModelPkg) {

	// modelpkgFrom has already been copied
	if _modelpkgTo, ok := mapOrigCopy[modelpkgFrom]; ok {
		modelpkgTo = _modelpkgTo.(*ModelPkg)
		return
	}

	modelpkgTo = new(ModelPkg)
	mapOrigCopy[modelpkgFrom] = modelpkgTo
	modelpkgFrom.GongCopyBasicFields(modelpkgTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPointerToGongStructField(mapOrigCopy map[any]any, pointertogongstructfieldFrom *PointerToGongStructField) (pointertogongstructfieldTo *PointerToGongStructField) {

	// pointertogongstructfieldFrom has already been copied
	if _pointertogongstructfieldTo, ok := mapOrigCopy[pointertogongstructfieldFrom]; ok {
		pointertogongstructfieldTo = _pointertogongstructfieldTo.(*PointerToGongStructField)
		return
	}

	pointertogongstructfieldTo = new(PointerToGongStructField)
	mapOrigCopy[pointertogongstructfieldFrom] = pointertogongstructfieldTo
	pointertogongstructfieldFrom.GongCopyBasicFields(pointertogongstructfieldTo)

	//insertion point for the staging of instances referenced by pointers
	if pointertogongstructfieldFrom.GongStruct != nil {
		pointertogongstructfieldTo.GongStruct = GongCopyBranchGongStruct(mapOrigCopy, pointertogongstructfieldFrom.GongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSliceOfPointerToGongStructField(mapOrigCopy map[any]any, sliceofpointertogongstructfieldFrom *SliceOfPointerToGongStructField) (sliceofpointertogongstructfieldTo *SliceOfPointerToGongStructField) {

	// sliceofpointertogongstructfieldFrom has already been copied
	if _sliceofpointertogongstructfieldTo, ok := mapOrigCopy[sliceofpointertogongstructfieldFrom]; ok {
		sliceofpointertogongstructfieldTo = _sliceofpointertogongstructfieldTo.(*SliceOfPointerToGongStructField)
		return
	}

	sliceofpointertogongstructfieldTo = new(SliceOfPointerToGongStructField)
	mapOrigCopy[sliceofpointertogongstructfieldFrom] = sliceofpointertogongstructfieldTo
	sliceofpointertogongstructfieldFrom.GongCopyBasicFields(sliceofpointertogongstructfieldTo)

	//insertion point for the staging of instances referenced by pointers
	if sliceofpointertogongstructfieldFrom.GongStruct != nil {
		sliceofpointertogongstructfieldTo.GongStruct = GongCopyBranchGongStruct(mapOrigCopy, sliceofpointertogongstructfieldFrom.GongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStageSetField(mapOrigCopy map[any]any, stagesetfieldFrom *StageSetField) (stagesetfieldTo *StageSetField) {

	// stagesetfieldFrom has already been copied
	if _stagesetfieldTo, ok := mapOrigCopy[stagesetfieldFrom]; ok {
		stagesetfieldTo = _stagesetfieldTo.(*StageSetField)
		return
	}

	stagesetfieldTo = new(StageSetField)
	mapOrigCopy[stagesetfieldFrom] = stagesetfieldTo
	stagesetfieldFrom.GongCopyBasicFields(stagesetfieldTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStageSetModel(mapOrigCopy map[any]any, stagesetmodelFrom *StageSetModel) (stagesetmodelTo *StageSetModel) {

	// stagesetmodelFrom has already been copied
	if _stagesetmodelTo, ok := mapOrigCopy[stagesetmodelFrom]; ok {
		stagesetmodelTo = _stagesetmodelTo.(*StageSetModel)
		return
	}

	stagesetmodelTo = new(StageSetModel)
	mapOrigCopy[stagesetmodelFrom] = stagesetmodelTo
	stagesetmodelFrom.GongCopyBasicFields(stagesetmodelTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _stagesetfield := range stagesetmodelFrom.Fields {
		stagesetmodelTo.Fields = append(stagesetmodelTo.Fields, GongCopyBranchStageSetField(mapOrigCopy, _stagesetfield))
	}

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
func (gongbasicfield *GongBasicField) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGongBasicField(gongbasicfield)
}

func (stage *Stage) UnstageBranchGongBasicField(gongbasicfield *GongBasicField) {

	// check if instance is already staged
	if !stage.IsStaged(gongbasicfield) {
		return
	}

	gongbasicfield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if gongbasicfield.GongEnum != nil {
		stage.UnstageBranch(gongbasicfield.GongEnum)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gongenum *GongEnum) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGongEnum(gongenum)
}

func (stage *Stage) UnstageBranchGongEnum(gongenum *GongEnum) {

	// check if instance is already staged
	if !stage.IsStaged(gongenum) {
		return
	}

	gongenum.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalue := range gongenum.GongEnumValues {
		stage.UnstageBranch(_gongenumvalue)
	}

}

func (gongenumvalue *GongEnumValue) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGongEnumValue(gongenumvalue)
}

func (stage *Stage) UnstageBranchGongEnumValue(gongenumvalue *GongEnumValue) {

	// check if instance is already staged
	if !stage.IsStaged(gongenumvalue) {
		return
	}

	gongenumvalue.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gonglink *GongLink) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGongLink(gonglink)
}

func (stage *Stage) UnstageBranchGongLink(gonglink *GongLink) {

	// check if instance is already staged
	if !stage.IsStaged(gonglink) {
		return
	}

	gonglink.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gongnote *GongNote) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGongNote(gongnote)
}

func (stage *Stage) UnstageBranchGongNote(gongnote *GongNote) {

	// check if instance is already staged
	if !stage.IsStaged(gongnote) {
		return
	}

	gongnote.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gonglink := range gongnote.Links {
		stage.UnstageBranch(_gonglink)
	}

}

func (gongstruct *GongStruct) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGongStruct(gongstruct)
}

func (stage *Stage) UnstageBranchGongStruct(gongstruct *GongStruct) {

	// check if instance is already staged
	if !stage.IsStaged(gongstruct) {
		return
	}

	gongstruct.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongbasicfield := range gongstruct.GongBasicFields {
		stage.UnstageBranch(_gongbasicfield)
	}
	for _, _gongtimefield := range gongstruct.GongTimeFields {
		stage.UnstageBranch(_gongtimefield)
	}
	for _, _pointertogongstructfield := range gongstruct.PointerToGongStructFields {
		stage.UnstageBranch(_pointertogongstructfield)
	}
	for _, _sliceofpointertogongstructfield := range gongstruct.SliceOfPointerToGongStructFields {
		stage.UnstageBranch(_sliceofpointertogongstructfield)
	}

}

func (gongtimefield *GongTimeField) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchGongTimeField(gongtimefield)
}

func (stage *Stage) UnstageBranchGongTimeField(gongtimefield *GongTimeField) {

	// check if instance is already staged
	if !stage.IsStaged(gongtimefield) {
		return
	}

	gongtimefield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metareference *MetaReference) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchMetaReference(metareference)
}

func (stage *Stage) UnstageBranchMetaReference(metareference *MetaReference) {

	// check if instance is already staged
	if !stage.IsStaged(metareference) {
		return
	}

	metareference.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (modelpkg *ModelPkg) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchModelPkg(modelpkg)
}

func (stage *Stage) UnstageBranchModelPkg(modelpkg *ModelPkg) {

	// check if instance is already staged
	if !stage.IsStaged(modelpkg) {
		return
	}

	modelpkg.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pointertogongstructfield *PointerToGongStructField) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchPointerToGongStructField(pointertogongstructfield)
}

func (stage *Stage) UnstageBranchPointerToGongStructField(pointertogongstructfield *PointerToGongStructField) {

	// check if instance is already staged
	if !stage.IsStaged(pointertogongstructfield) {
		return
	}

	pointertogongstructfield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pointertogongstructfield.GongStruct != nil {
		stage.UnstageBranch(pointertogongstructfield.GongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSliceOfPointerToGongStructField(sliceofpointertogongstructfield)
}

func (stage *Stage) UnstageBranchSliceOfPointerToGongStructField(sliceofpointertogongstructfield *SliceOfPointerToGongStructField) {

	// check if instance is already staged
	if !stage.IsStaged(sliceofpointertogongstructfield) {
		return
	}

	sliceofpointertogongstructfield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sliceofpointertogongstructfield.GongStruct != nil {
		stage.UnstageBranch(sliceofpointertogongstructfield.GongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stagesetfield *StageSetField) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStageSetField(stagesetfield)
}

func (stage *Stage) UnstageBranchStageSetField(stagesetfield *StageSetField) {

	// check if instance is already staged
	if !stage.IsStaged(stagesetfield) {
		return
	}

	stagesetfield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stagesetmodel *StageSetModel) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchStageSetModel(stagesetmodel)
}

func (stage *Stage) UnstageBranchStageSetModel(stagesetmodel *StageSetModel) {

	// check if instance is already staged
	if !stage.IsStaged(stagesetmodel) {
		return
	}

	stagesetmodel.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _stagesetfield := range stagesetmodel.Fields {
		stage.UnstageBranch(_stagesetfield)
	}

}

// insertion point for pointer reconstruction from references
func (reference *GongBasicField) GongReconstructPointersFromReferences(stage *Stage, instance *GongBasicField) {
	// insertion point for pointers field
	if instance.GongEnum != nil {
		reference.GongEnum = stage.GongEnums_reference[instance.GongEnum]
	}
	// insertion point for slice of pointers field
}

func (reference *GongEnum) GongReconstructPointersFromReferences(stage *Stage, instance *GongEnum) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.GongEnumValues = reference.GongEnumValues[:0]
	for _, _b := range instance.GongEnumValues {
		reference.GongEnumValues = append(reference.GongEnumValues, stage.GongEnumValues_reference[_b])
	}
}

func (reference *GongEnumValue) GongReconstructPointersFromReferences(stage *Stage, instance *GongEnumValue) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GongLink) GongReconstructPointersFromReferences(stage *Stage, instance *GongLink) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *GongNote) GongReconstructPointersFromReferences(stage *Stage, instance *GongNote) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Links = reference.Links[:0]
	for _, _b := range instance.Links {
		reference.Links = append(reference.Links, stage.GongLinks_reference[_b])
	}
}

func (reference *GongStruct) GongReconstructPointersFromReferences(stage *Stage, instance *GongStruct) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.GongBasicFields = reference.GongBasicFields[:0]
	for _, _b := range instance.GongBasicFields {
		reference.GongBasicFields = append(reference.GongBasicFields, stage.GongBasicFields_reference[_b])
	}
	reference.GongTimeFields = reference.GongTimeFields[:0]
	for _, _b := range instance.GongTimeFields {
		reference.GongTimeFields = append(reference.GongTimeFields, stage.GongTimeFields_reference[_b])
	}
	reference.PointerToGongStructFields = reference.PointerToGongStructFields[:0]
	for _, _b := range instance.PointerToGongStructFields {
		reference.PointerToGongStructFields = append(reference.PointerToGongStructFields, stage.PointerToGongStructFields_reference[_b])
	}
	reference.SliceOfPointerToGongStructFields = reference.SliceOfPointerToGongStructFields[:0]
	for _, _b := range instance.SliceOfPointerToGongStructFields {
		reference.SliceOfPointerToGongStructFields = append(reference.SliceOfPointerToGongStructFields, stage.SliceOfPointerToGongStructFields_reference[_b])
	}
}

func (reference *GongTimeField) GongReconstructPointersFromReferences(stage *Stage, instance *GongTimeField) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MetaReference) GongReconstructPointersFromReferences(stage *Stage, instance *MetaReference) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ModelPkg) GongReconstructPointersFromReferences(stage *Stage, instance *ModelPkg) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *PointerToGongStructField) GongReconstructPointersFromReferences(stage *Stage, instance *PointerToGongStructField) {
	// insertion point for pointers field
	if instance.GongStruct != nil {
		reference.GongStruct = stage.GongStructs_reference[instance.GongStruct]
	}
	// insertion point for slice of pointers field
}

func (reference *SliceOfPointerToGongStructField) GongReconstructPointersFromReferences(stage *Stage, instance *SliceOfPointerToGongStructField) {
	// insertion point for pointers field
	if instance.GongStruct != nil {
		reference.GongStruct = stage.GongStructs_reference[instance.GongStruct]
	}
	// insertion point for slice of pointers field
}

func (reference *StageSetField) GongReconstructPointersFromReferences(stage *Stage, instance *StageSetField) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StageSetModel) GongReconstructPointersFromReferences(stage *Stage, instance *StageSetModel) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.Fields = reference.Fields[:0]
	for _, _b := range instance.Fields {
		reference.Fields = append(reference.Fields, stage.StageSetFields_reference[_b])
	}
}

// insertion point for pointer reconstruction from instances
func (reference *GongBasicField) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.GongEnum; _reference != nil {
		reference.GongEnum = nil
		if _instance, ok := stage.GongEnums_instance[_reference]; ok {
			reference.GongEnum = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *GongEnum) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _GongEnumValues []*GongEnumValue
	for _, _reference := range reference.GongEnumValues {
		if _instance, ok := stage.GongEnumValues_instance[_reference]; ok {
			_GongEnumValues = append(_GongEnumValues, _instance)
		}
	}
	reference.GongEnumValues = _GongEnumValues
}

func (reference *GongEnumValue) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GongLink) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *GongNote) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Links []*GongLink
	for _, _reference := range reference.Links {
		if _instance, ok := stage.GongLinks_instance[_reference]; ok {
			_Links = append(_Links, _instance)
		}
	}
	reference.Links = _Links
}

func (reference *GongStruct) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _GongBasicFields []*GongBasicField
	for _, _reference := range reference.GongBasicFields {
		if _instance, ok := stage.GongBasicFields_instance[_reference]; ok {
			_GongBasicFields = append(_GongBasicFields, _instance)
		}
	}
	reference.GongBasicFields = _GongBasicFields
	var _GongTimeFields []*GongTimeField
	for _, _reference := range reference.GongTimeFields {
		if _instance, ok := stage.GongTimeFields_instance[_reference]; ok {
			_GongTimeFields = append(_GongTimeFields, _instance)
		}
	}
	reference.GongTimeFields = _GongTimeFields
	var _PointerToGongStructFields []*PointerToGongStructField
	for _, _reference := range reference.PointerToGongStructFields {
		if _instance, ok := stage.PointerToGongStructFields_instance[_reference]; ok {
			_PointerToGongStructFields = append(_PointerToGongStructFields, _instance)
		}
	}
	reference.PointerToGongStructFields = _PointerToGongStructFields
	var _SliceOfPointerToGongStructFields []*SliceOfPointerToGongStructField
	for _, _reference := range reference.SliceOfPointerToGongStructFields {
		if _instance, ok := stage.SliceOfPointerToGongStructFields_instance[_reference]; ok {
			_SliceOfPointerToGongStructFields = append(_SliceOfPointerToGongStructFields, _instance)
		}
	}
	reference.SliceOfPointerToGongStructFields = _SliceOfPointerToGongStructFields
}

func (reference *GongTimeField) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MetaReference) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ModelPkg) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *PointerToGongStructField) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.GongStruct; _reference != nil {
		reference.GongStruct = nil
		if _instance, ok := stage.GongStructs_instance[_reference]; ok {
			reference.GongStruct = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *SliceOfPointerToGongStructField) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.GongStruct; _reference != nil {
		reference.GongStruct = nil
		if _instance, ok := stage.GongStructs_instance[_reference]; ok {
			reference.GongStruct = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *StageSetField) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StageSetModel) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _Fields []*StageSetField
	for _, _reference := range reference.Fields {
		if _instance, ok := stage.StageSetFields_instance[_reference]; ok {
			_Fields = append(_Fields, _instance)
		}
	}
	reference.Fields = _Fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gongbasicfield *GongBasicField) GongDiff(stage *Stage, gongbasicfieldOther *GongBasicField) (diffs []string) {
	// insertion point for field diffs
	if gongbasicfield.Name != gongbasicfieldOther.Name {
		diffs = append(diffs, gongbasicfield.GongMarshallField(stage, "Name"))
	}
	if gongbasicfield.BasicKindName != gongbasicfieldOther.BasicKindName {
		diffs = append(diffs, gongbasicfield.GongMarshallField(stage, "BasicKindName"))
	}
	if (gongbasicfield.GongEnum == nil) != (gongbasicfieldOther.GongEnum == nil) {
		diffs = append(diffs, gongbasicfield.GongMarshallField(stage, "GongEnum"))
	} else if gongbasicfield.GongEnum != nil && gongbasicfieldOther.GongEnum != nil {
		if gongbasicfield.GongEnum != gongbasicfieldOther.GongEnum {
			diffs = append(diffs, gongbasicfield.GongMarshallField(stage, "GongEnum"))
		}
	}
	if gongbasicfield.DeclaredType != gongbasicfieldOther.DeclaredType {
		diffs = append(diffs, gongbasicfield.GongMarshallField(stage, "DeclaredType"))
	}
	if gongbasicfield.CompositeStructName != gongbasicfieldOther.CompositeStructName {
		diffs = append(diffs, gongbasicfield.GongMarshallField(stage, "CompositeStructName"))
	}
	if gongbasicfield.IsAccordionStart != gongbasicfieldOther.IsAccordionStart {
		diffs = append(diffs, gongbasicfield.GongMarshallField(stage, "IsAccordionStart"))
	}
	if gongbasicfield.AccordionName != gongbasicfieldOther.AccordionName {
		diffs = append(diffs, gongbasicfield.GongMarshallField(stage, "AccordionName"))
	}
	if gongbasicfield.IsAccordionEnd != gongbasicfieldOther.IsAccordionEnd {
		diffs = append(diffs, gongbasicfield.GongMarshallField(stage, "IsAccordionEnd"))
	}
	if gongbasicfield.Index != gongbasicfieldOther.Index {
		diffs = append(diffs, gongbasicfield.GongMarshallField(stage, "Index"))
	}
	if gongbasicfield.IsTextArea != gongbasicfieldOther.IsTextArea {
		diffs = append(diffs, gongbasicfield.GongMarshallField(stage, "IsTextArea"))
	}
	if gongbasicfield.IsBespokeWidth != gongbasicfieldOther.IsBespokeWidth {
		diffs = append(diffs, gongbasicfield.GongMarshallField(stage, "IsBespokeWidth"))
	}
	if gongbasicfield.BespokeWidth != gongbasicfieldOther.BespokeWidth {
		diffs = append(diffs, gongbasicfield.GongMarshallField(stage, "BespokeWidth"))
	}
	if gongbasicfield.IsBespokeHeight != gongbasicfieldOther.IsBespokeHeight {
		diffs = append(diffs, gongbasicfield.GongMarshallField(stage, "IsBespokeHeight"))
	}
	if gongbasicfield.BespokeHeight != gongbasicfieldOther.BespokeHeight {
		diffs = append(diffs, gongbasicfield.GongMarshallField(stage, "BespokeHeight"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gongenum *GongEnum) GongDiff(stage *Stage, gongenumOther *GongEnum) (diffs []string) {
	// insertion point for field diffs
	if gongenum.Name != gongenumOther.Name {
		diffs = append(diffs, gongenum.GongMarshallField(stage, "Name"))
	}
	if gongenum.Type != gongenumOther.Type {
		diffs = append(diffs, gongenum.GongMarshallField(stage, "Type"))
	}
	GongEnumValuesDifferent := false
	if len(gongenum.GongEnumValues) != len(gongenumOther.GongEnumValues) {
		GongEnumValuesDifferent = true
	} else {
		for i := range gongenum.GongEnumValues {
			if (gongenum.GongEnumValues[i] == nil) != (gongenumOther.GongEnumValues[i] == nil) {
				GongEnumValuesDifferent = true
				break
			} else if gongenum.GongEnumValues[i] != nil && gongenumOther.GongEnumValues[i] != nil {
				// this is a pointer comparaison
				if gongenum.GongEnumValues[i] != gongenumOther.GongEnumValues[i] {
					GongEnumValuesDifferent = true
					break
				}
			}
		}
	}
	if GongEnumValuesDifferent {
		ops := stage.Diff(
			gongenum,
			"GongEnumValues",
			len(gongenumOther.GongEnumValues),
			len(gongenum.GongEnumValues),
			func(i, j int) bool {
				return gongenumOther.GongEnumValues[i] == gongenum.GongEnumValues[j]
			},
			func(j int) string {
				return gongenum.GongEnumValues[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gongenumvalue *GongEnumValue) GongDiff(stage *Stage, gongenumvalueOther *GongEnumValue) (diffs []string) {
	// insertion point for field diffs
	if gongenumvalue.Name != gongenumvalueOther.Name {
		diffs = append(diffs, gongenumvalue.GongMarshallField(stage, "Name"))
	}
	if gongenumvalue.Value != gongenumvalueOther.Value {
		diffs = append(diffs, gongenumvalue.GongMarshallField(stage, "Value"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gonglink *GongLink) GongDiff(stage *Stage, gonglinkOther *GongLink) (diffs []string) {
	// insertion point for field diffs
	if gonglink.Name != gonglinkOther.Name {
		diffs = append(diffs, gonglink.GongMarshallField(stage, "Name"))
	}
	if gonglink.Recv != gonglinkOther.Recv {
		diffs = append(diffs, gonglink.GongMarshallField(stage, "Recv"))
	}
	if gonglink.ImportPath != gonglinkOther.ImportPath {
		diffs = append(diffs, gonglink.GongMarshallField(stage, "ImportPath"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gongnote *GongNote) GongDiff(stage *Stage, gongnoteOther *GongNote) (diffs []string) {
	// insertion point for field diffs
	if gongnote.Name != gongnoteOther.Name {
		diffs = append(diffs, gongnote.GongMarshallField(stage, "Name"))
	}
	if gongnote.Body != gongnoteOther.Body {
		diffs = append(diffs, gongnote.GongMarshallField(stage, "Body"))
	}
	if gongnote.BodyHTML != gongnoteOther.BodyHTML {
		diffs = append(diffs, gongnote.GongMarshallField(stage, "BodyHTML"))
	}
	LinksDifferent := false
	if len(gongnote.Links) != len(gongnoteOther.Links) {
		LinksDifferent = true
	} else {
		for i := range gongnote.Links {
			if (gongnote.Links[i] == nil) != (gongnoteOther.Links[i] == nil) {
				LinksDifferent = true
				break
			} else if gongnote.Links[i] != nil && gongnoteOther.Links[i] != nil {
				// this is a pointer comparaison
				if gongnote.Links[i] != gongnoteOther.Links[i] {
					LinksDifferent = true
					break
				}
			}
		}
	}
	if LinksDifferent {
		ops := stage.Diff(
			gongnote,
			"Links",
			len(gongnoteOther.Links),
			len(gongnote.Links),
			func(i, j int) bool {
				return gongnoteOther.Links[i] == gongnote.Links[j]
			},
			func(j int) string {
				return gongnote.Links[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gongstruct *GongStruct) GongDiff(stage *Stage, gongstructOther *GongStruct) (diffs []string) {
	// insertion point for field diffs
	if gongstruct.Name != gongstructOther.Name {
		diffs = append(diffs, gongstruct.GongMarshallField(stage, "Name"))
	}
	GongBasicFieldsDifferent := false
	if len(gongstruct.GongBasicFields) != len(gongstructOther.GongBasicFields) {
		GongBasicFieldsDifferent = true
	} else {
		for i := range gongstruct.GongBasicFields {
			if (gongstruct.GongBasicFields[i] == nil) != (gongstructOther.GongBasicFields[i] == nil) {
				GongBasicFieldsDifferent = true
				break
			} else if gongstruct.GongBasicFields[i] != nil && gongstructOther.GongBasicFields[i] != nil {
				// this is a pointer comparaison
				if gongstruct.GongBasicFields[i] != gongstructOther.GongBasicFields[i] {
					GongBasicFieldsDifferent = true
					break
				}
			}
		}
	}
	if GongBasicFieldsDifferent {
		ops := stage.Diff(
			gongstruct,
			"GongBasicFields",
			len(gongstructOther.GongBasicFields),
			len(gongstruct.GongBasicFields),
			func(i, j int) bool {
				return gongstructOther.GongBasicFields[i] == gongstruct.GongBasicFields[j]
			},
			func(j int) string {
				return gongstruct.GongBasicFields[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	GongTimeFieldsDifferent := false
	if len(gongstruct.GongTimeFields) != len(gongstructOther.GongTimeFields) {
		GongTimeFieldsDifferent = true
	} else {
		for i := range gongstruct.GongTimeFields {
			if (gongstruct.GongTimeFields[i] == nil) != (gongstructOther.GongTimeFields[i] == nil) {
				GongTimeFieldsDifferent = true
				break
			} else if gongstruct.GongTimeFields[i] != nil && gongstructOther.GongTimeFields[i] != nil {
				// this is a pointer comparaison
				if gongstruct.GongTimeFields[i] != gongstructOther.GongTimeFields[i] {
					GongTimeFieldsDifferent = true
					break
				}
			}
		}
	}
	if GongTimeFieldsDifferent {
		ops := stage.Diff(
			gongstruct,
			"GongTimeFields",
			len(gongstructOther.GongTimeFields),
			len(gongstruct.GongTimeFields),
			func(i, j int) bool {
				return gongstructOther.GongTimeFields[i] == gongstruct.GongTimeFields[j]
			},
			func(j int) string {
				return gongstruct.GongTimeFields[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	PointerToGongStructFieldsDifferent := false
	if len(gongstruct.PointerToGongStructFields) != len(gongstructOther.PointerToGongStructFields) {
		PointerToGongStructFieldsDifferent = true
	} else {
		for i := range gongstruct.PointerToGongStructFields {
			if (gongstruct.PointerToGongStructFields[i] == nil) != (gongstructOther.PointerToGongStructFields[i] == nil) {
				PointerToGongStructFieldsDifferent = true
				break
			} else if gongstruct.PointerToGongStructFields[i] != nil && gongstructOther.PointerToGongStructFields[i] != nil {
				// this is a pointer comparaison
				if gongstruct.PointerToGongStructFields[i] != gongstructOther.PointerToGongStructFields[i] {
					PointerToGongStructFieldsDifferent = true
					break
				}
			}
		}
	}
	if PointerToGongStructFieldsDifferent {
		ops := stage.Diff(
			gongstruct,
			"PointerToGongStructFields",
			len(gongstructOther.PointerToGongStructFields),
			len(gongstruct.PointerToGongStructFields),
			func(i, j int) bool {
				return gongstructOther.PointerToGongStructFields[i] == gongstruct.PointerToGongStructFields[j]
			},
			func(j int) string {
				return gongstruct.PointerToGongStructFields[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	SliceOfPointerToGongStructFieldsDifferent := false
	if len(gongstruct.SliceOfPointerToGongStructFields) != len(gongstructOther.SliceOfPointerToGongStructFields) {
		SliceOfPointerToGongStructFieldsDifferent = true
	} else {
		for i := range gongstruct.SliceOfPointerToGongStructFields {
			if (gongstruct.SliceOfPointerToGongStructFields[i] == nil) != (gongstructOther.SliceOfPointerToGongStructFields[i] == nil) {
				SliceOfPointerToGongStructFieldsDifferent = true
				break
			} else if gongstruct.SliceOfPointerToGongStructFields[i] != nil && gongstructOther.SliceOfPointerToGongStructFields[i] != nil {
				// this is a pointer comparaison
				if gongstruct.SliceOfPointerToGongStructFields[i] != gongstructOther.SliceOfPointerToGongStructFields[i] {
					SliceOfPointerToGongStructFieldsDifferent = true
					break
				}
			}
		}
	}
	if SliceOfPointerToGongStructFieldsDifferent {
		ops := stage.Diff(
			gongstruct,
			"SliceOfPointerToGongStructFields",
			len(gongstructOther.SliceOfPointerToGongStructFields),
			len(gongstruct.SliceOfPointerToGongStructFields),
			func(i, j int) bool {
				return gongstructOther.SliceOfPointerToGongStructFields[i] == gongstruct.SliceOfPointerToGongStructFields[j]
			},
			func(j int) string {
				return gongstruct.SliceOfPointerToGongStructFields[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if gongstruct.HasOnAfterUpdateSignature != gongstructOther.HasOnAfterUpdateSignature {
		diffs = append(diffs, gongstruct.GongMarshallField(stage, "HasOnAfterUpdateSignature"))
	}
	if gongstruct.IsIgnoredForFront != gongstructOther.IsIgnoredForFront {
		diffs = append(diffs, gongstruct.GongMarshallField(stage, "IsIgnoredForFront"))
	}
	if gongstruct.IsOmittedForMarshalling != gongstructOther.IsOmittedForMarshalling {
		diffs = append(diffs, gongstruct.GongMarshallField(stage, "IsOmittedForMarshalling"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gongtimefield *GongTimeField) GongDiff(stage *Stage, gongtimefieldOther *GongTimeField) (diffs []string) {
	// insertion point for field diffs
	if gongtimefield.Name != gongtimefieldOther.Name {
		diffs = append(diffs, gongtimefield.GongMarshallField(stage, "Name"))
	}
	if gongtimefield.Index != gongtimefieldOther.Index {
		diffs = append(diffs, gongtimefield.GongMarshallField(stage, "Index"))
	}
	if gongtimefield.CompositeStructName != gongtimefieldOther.CompositeStructName {
		diffs = append(diffs, gongtimefield.GongMarshallField(stage, "CompositeStructName"))
	}
	if gongtimefield.IsAccordionStart != gongtimefieldOther.IsAccordionStart {
		diffs = append(diffs, gongtimefield.GongMarshallField(stage, "IsAccordionStart"))
	}
	if gongtimefield.AccordionName != gongtimefieldOther.AccordionName {
		diffs = append(diffs, gongtimefield.GongMarshallField(stage, "AccordionName"))
	}
	if gongtimefield.IsAccordionEnd != gongtimefieldOther.IsAccordionEnd {
		diffs = append(diffs, gongtimefield.GongMarshallField(stage, "IsAccordionEnd"))
	}
	if gongtimefield.BespokeTimeFormat != gongtimefieldOther.BespokeTimeFormat {
		diffs = append(diffs, gongtimefield.GongMarshallField(stage, "BespokeTimeFormat"))
	}
	if gongtimefield.TimeFormOnly != gongtimefieldOther.TimeFormOnly {
		diffs = append(diffs, gongtimefield.GongMarshallField(stage, "TimeFormOnly"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (metareference *MetaReference) GongDiff(stage *Stage, metareferenceOther *MetaReference) (diffs []string) {
	// insertion point for field diffs
	if metareference.Name != metareferenceOther.Name {
		diffs = append(diffs, metareference.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (modelpkg *ModelPkg) GongDiff(stage *Stage, modelpkgOther *ModelPkg) (diffs []string) {
	// insertion point for field diffs
	if modelpkg.Name != modelpkgOther.Name {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "Name"))
	}
	if modelpkg.PkgGoName != modelpkgOther.PkgGoName {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "PkgGoName"))
	}
	if modelpkg.PkgPath != modelpkgOther.PkgPath {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "PkgPath"))
	}
	if modelpkg.PathToGoSubDirectory != modelpkgOther.PathToGoSubDirectory {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "PathToGoSubDirectory"))
	}
	if modelpkg.OrmPkgGenPath != modelpkgOther.OrmPkgGenPath {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "OrmPkgGenPath"))
	}
	if modelpkg.DbOrmPkgGenPath != modelpkgOther.DbOrmPkgGenPath {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "DbOrmPkgGenPath"))
	}
	if modelpkg.DbLiteOrmPkgGenPath != modelpkgOther.DbLiteOrmPkgGenPath {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "DbLiteOrmPkgGenPath"))
	}
	if modelpkg.DbPkgGenPath != modelpkgOther.DbPkgGenPath {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "DbPkgGenPath"))
	}
	if modelpkg.ControllersPkgGenPath != modelpkgOther.ControllersPkgGenPath {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "ControllersPkgGenPath"))
	}
	if modelpkg.FullstackPkgGenPath != modelpkgOther.FullstackPkgGenPath {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "FullstackPkgGenPath"))
	}
	if modelpkg.StackPkgGenPath != modelpkgOther.StackPkgGenPath {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "StackPkgGenPath"))
	}
	if modelpkg.Level1StackPkgGenPath != modelpkgOther.Level1StackPkgGenPath {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "Level1StackPkgGenPath"))
	}
	if modelpkg.StaticPkgGenPath != modelpkgOther.StaticPkgGenPath {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "StaticPkgGenPath"))
	}
	if modelpkg.ProbePkgGenPath != modelpkgOther.ProbePkgGenPath {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "ProbePkgGenPath"))
	}
	if modelpkg.NgWorkspacePath != modelpkgOther.NgWorkspacePath {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "NgWorkspacePath"))
	}
	if modelpkg.NgWorkspaceName != modelpkgOther.NgWorkspaceName {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "NgWorkspaceName"))
	}
	if modelpkg.NgDataLibrarySourceCodeDirectory != modelpkgOther.NgDataLibrarySourceCodeDirectory {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "NgDataLibrarySourceCodeDirectory"))
	}
	if modelpkg.NgSpecificLibrarySourceCodeDirectory != modelpkgOther.NgSpecificLibrarySourceCodeDirectory {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "NgSpecificLibrarySourceCodeDirectory"))
	}
	if modelpkg.MaterialLibDatamodelTargetPath != modelpkgOther.MaterialLibDatamodelTargetPath {
		diffs = append(diffs, modelpkg.GongMarshallField(stage, "MaterialLibDatamodelTargetPath"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pointertogongstructfield *PointerToGongStructField) GongDiff(stage *Stage, pointertogongstructfieldOther *PointerToGongStructField) (diffs []string) {
	// insertion point for field diffs
	if pointertogongstructfield.Name != pointertogongstructfieldOther.Name {
		diffs = append(diffs, pointertogongstructfield.GongMarshallField(stage, "Name"))
	}
	if (pointertogongstructfield.GongStruct == nil) != (pointertogongstructfieldOther.GongStruct == nil) {
		diffs = append(diffs, pointertogongstructfield.GongMarshallField(stage, "GongStruct"))
	} else if pointertogongstructfield.GongStruct != nil && pointertogongstructfieldOther.GongStruct != nil {
		if pointertogongstructfield.GongStruct != pointertogongstructfieldOther.GongStruct {
			diffs = append(diffs, pointertogongstructfield.GongMarshallField(stage, "GongStruct"))
		}
	}
	if pointertogongstructfield.Index != pointertogongstructfieldOther.Index {
		diffs = append(diffs, pointertogongstructfield.GongMarshallField(stage, "Index"))
	}
	if pointertogongstructfield.CompositeStructName != pointertogongstructfieldOther.CompositeStructName {
		diffs = append(diffs, pointertogongstructfield.GongMarshallField(stage, "CompositeStructName"))
	}
	if pointertogongstructfield.IsAccordionStart != pointertogongstructfieldOther.IsAccordionStart {
		diffs = append(diffs, pointertogongstructfield.GongMarshallField(stage, "IsAccordionStart"))
	}
	if pointertogongstructfield.AccordionName != pointertogongstructfieldOther.AccordionName {
		diffs = append(diffs, pointertogongstructfield.GongMarshallField(stage, "AccordionName"))
	}
	if pointertogongstructfield.IsAccordionEnd != pointertogongstructfieldOther.IsAccordionEnd {
		diffs = append(diffs, pointertogongstructfield.GongMarshallField(stage, "IsAccordionEnd"))
	}
	if pointertogongstructfield.IsType != pointertogongstructfieldOther.IsType {
		diffs = append(diffs, pointertogongstructfield.GongMarshallField(stage, "IsType"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongDiff(stage *Stage, sliceofpointertogongstructfieldOther *SliceOfPointerToGongStructField) (diffs []string) {
	// insertion point for field diffs
	if sliceofpointertogongstructfield.Name != sliceofpointertogongstructfieldOther.Name {
		diffs = append(diffs, sliceofpointertogongstructfield.GongMarshallField(stage, "Name"))
	}
	if (sliceofpointertogongstructfield.GongStruct == nil) != (sliceofpointertogongstructfieldOther.GongStruct == nil) {
		diffs = append(diffs, sliceofpointertogongstructfield.GongMarshallField(stage, "GongStruct"))
	} else if sliceofpointertogongstructfield.GongStruct != nil && sliceofpointertogongstructfieldOther.GongStruct != nil {
		if sliceofpointertogongstructfield.GongStruct != sliceofpointertogongstructfieldOther.GongStruct {
			diffs = append(diffs, sliceofpointertogongstructfield.GongMarshallField(stage, "GongStruct"))
		}
	}
	if sliceofpointertogongstructfield.Index != sliceofpointertogongstructfieldOther.Index {
		diffs = append(diffs, sliceofpointertogongstructfield.GongMarshallField(stage, "Index"))
	}
	if sliceofpointertogongstructfield.CompositeStructName != sliceofpointertogongstructfieldOther.CompositeStructName {
		diffs = append(diffs, sliceofpointertogongstructfield.GongMarshallField(stage, "CompositeStructName"))
	}
	if sliceofpointertogongstructfield.IsAccordionStart != sliceofpointertogongstructfieldOther.IsAccordionStart {
		diffs = append(diffs, sliceofpointertogongstructfield.GongMarshallField(stage, "IsAccordionStart"))
	}
	if sliceofpointertogongstructfield.AccordionName != sliceofpointertogongstructfieldOther.AccordionName {
		diffs = append(diffs, sliceofpointertogongstructfield.GongMarshallField(stage, "AccordionName"))
	}
	if sliceofpointertogongstructfield.IsAccordionEnd != sliceofpointertogongstructfieldOther.IsAccordionEnd {
		diffs = append(diffs, sliceofpointertogongstructfield.GongMarshallField(stage, "IsAccordionEnd"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stagesetfield *StageSetField) GongDiff(stage *Stage, stagesetfieldOther *StageSetField) (diffs []string) {
	// insertion point for field diffs
	if stagesetfield.Name != stagesetfieldOther.Name {
		diffs = append(diffs, stagesetfield.GongMarshallField(stage, "Name"))
	}
	if stagesetfield.PackageName != stagesetfieldOther.PackageName {
		diffs = append(diffs, stagesetfield.GongMarshallField(stage, "PackageName"))
	}
	if stagesetfield.PackagePath != stagesetfieldOther.PackagePath {
		diffs = append(diffs, stagesetfield.GongMarshallField(stage, "PackagePath"))
	}
	if stagesetfield.IsLocal != stagesetfieldOther.IsLocal {
		diffs = append(diffs, stagesetfield.GongMarshallField(stage, "IsLocal"))
	}
	if stagesetfield.ImportAlias != stagesetfieldOther.ImportAlias {
		diffs = append(diffs, stagesetfield.GongMarshallField(stage, "ImportAlias"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (stagesetmodel *StageSetModel) GongDiff(stage *Stage, stagesetmodelOther *StageSetModel) (diffs []string) {
	// insertion point for field diffs
	if stagesetmodel.Name != stagesetmodelOther.Name {
		diffs = append(diffs, stagesetmodel.GongMarshallField(stage, "Name"))
	}
	FieldsDifferent := false
	if len(stagesetmodel.Fields) != len(stagesetmodelOther.Fields) {
		FieldsDifferent = true
	} else {
		for i := range stagesetmodel.Fields {
			if (stagesetmodel.Fields[i] == nil) != (stagesetmodelOther.Fields[i] == nil) {
				FieldsDifferent = true
				break
			} else if stagesetmodel.Fields[i] != nil && stagesetmodelOther.Fields[i] != nil {
				// this is a pointer comparaison
				if stagesetmodel.Fields[i] != stagesetmodelOther.Fields[i] {
					FieldsDifferent = true
					break
				}
			}
		}
	}
	if FieldsDifferent {
		ops := stage.Diff(
			stagesetmodel,
			"Fields",
			len(stagesetmodelOther.Fields),
			len(stagesetmodel.Fields),
			func(i, j int) bool {
				return stagesetmodelOther.Fields[i] == stagesetmodel.Fields[j]
			},
			func(j int) string {
				return stagesetmodel.Fields[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	if stagesetmodel.IsManual != stagesetmodelOther.IsManual {
		diffs = append(diffs, stagesetmodel.GongMarshallField(stage, "IsManual"))
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
