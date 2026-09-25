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
func (gongbasicfield *GongBasicField) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GongBasicFields[gongbasicfield]
	return ok
}

func (gongenum *GongEnum) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GongEnums[gongenum]
	return ok
}

func (gongenumvalue *GongEnumValue) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GongEnumValues[gongenumvalue]
	return ok
}

func (gonglink *GongLink) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GongLinks[gonglink]
	return ok
}

func (gongnote *GongNote) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GongNotes[gongnote]
	return ok
}

func (gongstruct *GongStruct) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GongStructs[gongstruct]
	return ok
}

func (gongtimefield *GongTimeField) GongIsStaged(stage *Stage) bool {
	_, ok := stage.GongTimeFields[gongtimefield]
	return ok
}

func (metareference *MetaReference) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MetaReferences[metareference]
	return ok
}

func (modelpkg *ModelPkg) GongIsStaged(stage *Stage) bool {
	_, ok := stage.ModelPkgs[modelpkg]
	return ok
}

func (pointertogongstructfield *PointerToGongStructField) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PointerToGongStructFields[pointertogongstructfield]
	return ok
}

func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongIsStaged(stage *Stage) bool {
	_, ok := stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield]
	return ok
}

func (stagesetfield *StageSetField) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StageSetFields[stagesetfield]
	return ok
}

func (stagesetmodel *StageSetModel) GongIsStaged(stage *Stage) bool {
	_, ok := stage.StageSetModels[stagesetmodel]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (gongbasicfield *GongBasicField) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(gongenumvalue) {
		return
	}

	gongenumvalue.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gonglink *GongLink) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(gonglink) {
		return
	}

	gonglink.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gongnote *GongNote) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(gongtimefield) {
		return
	}

	gongtimefield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metareference *MetaReference) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(metareference) {
		return
	}

	metareference.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (modelpkg *ModelPkg) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(modelpkg) {
		return
	}

	modelpkg.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pointertogongstructfield *PointerToGongStructField) GongStageBranch(stage *Stage) {

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

	// check if instance is already staged
	if stage.IsStaged(stagesetfield) {
		return
	}

	stagesetfield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stagesetmodel *StageSetModel) GongStageBranch(stage *Stage) {

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
	var alreadyCopied bool
	gongbasicfieldTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, gongbasicfieldFrom)
	if alreadyCopied {
		return
	}
	gongbasicfieldFrom.GongCopyBasicFields(gongbasicfieldTo)

	//insertion point for the staging of instances referenced by pointers
	if gongbasicfieldFrom.GongEnum != nil {
		gongbasicfieldTo.GongEnum = GongCopyBranchGongEnum(mapOrigCopy, gongbasicfieldFrom.GongEnum)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGongEnum(mapOrigCopy map[any]any, gongenumFrom *GongEnum) (gongenumTo *GongEnum) {
	var alreadyCopied bool
	gongenumTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, gongenumFrom)
	if alreadyCopied {
		return
	}
	gongenumFrom.GongCopyBasicFields(gongenumTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalue := range gongenumFrom.GongEnumValues {
		gongenumTo.GongEnumValues = append(gongenumTo.GongEnumValues, GongCopyBranchGongEnumValue(mapOrigCopy, _gongenumvalue))
	}

	return
}

func GongCopyBranchGongEnumValue(mapOrigCopy map[any]any, gongenumvalueFrom *GongEnumValue) (gongenumvalueTo *GongEnumValue) {
	var alreadyCopied bool
	gongenumvalueTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, gongenumvalueFrom)
	if alreadyCopied {
		return
	}
	gongenumvalueFrom.GongCopyBasicFields(gongenumvalueTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGongLink(mapOrigCopy map[any]any, gonglinkFrom *GongLink) (gonglinkTo *GongLink) {
	var alreadyCopied bool
	gonglinkTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, gonglinkFrom)
	if alreadyCopied {
		return
	}
	gonglinkFrom.GongCopyBasicFields(gonglinkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchGongNote(mapOrigCopy map[any]any, gongnoteFrom *GongNote) (gongnoteTo *GongNote) {
	var alreadyCopied bool
	gongnoteTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, gongnoteFrom)
	if alreadyCopied {
		return
	}
	gongnoteFrom.GongCopyBasicFields(gongnoteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gonglink := range gongnoteFrom.Links {
		gongnoteTo.Links = append(gongnoteTo.Links, GongCopyBranchGongLink(mapOrigCopy, _gonglink))
	}

	return
}

func GongCopyBranchGongStruct(mapOrigCopy map[any]any, gongstructFrom *GongStruct) (gongstructTo *GongStruct) {
	var alreadyCopied bool
	gongstructTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, gongstructFrom)
	if alreadyCopied {
		return
	}
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
	var alreadyCopied bool
	gongtimefieldTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, gongtimefieldFrom)
	if alreadyCopied {
		return
	}
	gongtimefieldFrom.GongCopyBasicFields(gongtimefieldTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMetaReference(mapOrigCopy map[any]any, metareferenceFrom *MetaReference) (metareferenceTo *MetaReference) {
	var alreadyCopied bool
	metareferenceTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, metareferenceFrom)
	if alreadyCopied {
		return
	}
	metareferenceFrom.GongCopyBasicFields(metareferenceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchModelPkg(mapOrigCopy map[any]any, modelpkgFrom *ModelPkg) (modelpkgTo *ModelPkg) {
	var alreadyCopied bool
	modelpkgTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, modelpkgFrom)
	if alreadyCopied {
		return
	}
	modelpkgFrom.GongCopyBasicFields(modelpkgTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPointerToGongStructField(mapOrigCopy map[any]any, pointertogongstructfieldFrom *PointerToGongStructField) (pointertogongstructfieldTo *PointerToGongStructField) {
	var alreadyCopied bool
	pointertogongstructfieldTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pointertogongstructfieldFrom)
	if alreadyCopied {
		return
	}
	pointertogongstructfieldFrom.GongCopyBasicFields(pointertogongstructfieldTo)

	//insertion point for the staging of instances referenced by pointers
	if pointertogongstructfieldFrom.GongStruct != nil {
		pointertogongstructfieldTo.GongStruct = GongCopyBranchGongStruct(mapOrigCopy, pointertogongstructfieldFrom.GongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSliceOfPointerToGongStructField(mapOrigCopy map[any]any, sliceofpointertogongstructfieldFrom *SliceOfPointerToGongStructField) (sliceofpointertogongstructfieldTo *SliceOfPointerToGongStructField) {
	var alreadyCopied bool
	sliceofpointertogongstructfieldTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, sliceofpointertogongstructfieldFrom)
	if alreadyCopied {
		return
	}
	sliceofpointertogongstructfieldFrom.GongCopyBasicFields(sliceofpointertogongstructfieldTo)

	//insertion point for the staging of instances referenced by pointers
	if sliceofpointertogongstructfieldFrom.GongStruct != nil {
		sliceofpointertogongstructfieldTo.GongStruct = GongCopyBranchGongStruct(mapOrigCopy, sliceofpointertogongstructfieldFrom.GongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStageSetField(mapOrigCopy map[any]any, stagesetfieldFrom *StageSetField) (stagesetfieldTo *StageSetField) {
	var alreadyCopied bool
	stagesetfieldTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stagesetfieldFrom)
	if alreadyCopied {
		return
	}
	stagesetfieldFrom.GongCopyBasicFields(stagesetfieldTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchStageSetModel(mapOrigCopy map[any]any, stagesetmodelFrom *StageSetModel) (stagesetmodelTo *StageSetModel) {
	var alreadyCopied bool
	stagesetmodelTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, stagesetmodelFrom)
	if alreadyCopied {
		return
	}
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

	// check if instance is already staged
	if !stage.IsStaged(gongenumvalue) {
		return
	}

	gongenumvalue.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gonglink *GongLink) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(gonglink) {
		return
	}

	gonglink.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (gongnote *GongNote) GongUnstageBranch(stage *Stage) {

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

	// check if instance is already staged
	if !stage.IsStaged(gongtimefield) {
		return
	}

	gongtimefield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (metareference *MetaReference) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(metareference) {
		return
	}

	metareference.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (modelpkg *ModelPkg) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(modelpkg) {
		return
	}

	modelpkg.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pointertogongstructfield *PointerToGongStructField) GongUnstageBranch(stage *Stage) {

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

	// check if instance is already staged
	if !stage.IsStaged(stagesetfield) {
		return
	}

	stagesetfield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stagesetmodel *StageSetModel) GongUnstageBranch(stage *Stage) {

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
	__gong__reconstructPointer(&reference.GongEnum, stage.GongEnums_reference, instance.GongEnum)
	// insertion point for slice of pointers field
}

func (reference *GongEnum) GongReconstructPointersFromReferences(stage *Stage, instance *GongEnum) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.GongEnumValues, stage.GongEnumValues_reference, instance.GongEnumValues)
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
	__gong__reconstructSliceOfPointersFromReferences(&reference.Links, stage.GongLinks_reference, instance.Links)
}

func (reference *GongStruct) GongReconstructPointersFromReferences(stage *Stage, instance *GongStruct) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.GongBasicFields, stage.GongBasicFields_reference, instance.GongBasicFields)
	__gong__reconstructSliceOfPointersFromReferences(&reference.GongTimeFields, stage.GongTimeFields_reference, instance.GongTimeFields)
	__gong__reconstructSliceOfPointersFromReferences(&reference.PointerToGongStructFields, stage.PointerToGongStructFields_reference, instance.PointerToGongStructFields)
	__gong__reconstructSliceOfPointersFromReferences(&reference.SliceOfPointerToGongStructFields, stage.SliceOfPointerToGongStructFields_reference, instance.SliceOfPointerToGongStructFields)
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
	__gong__reconstructPointer(&reference.GongStruct, stage.GongStructs_reference, instance.GongStruct)
	// insertion point for slice of pointers field
}

func (reference *SliceOfPointerToGongStructField) GongReconstructPointersFromReferences(stage *Stage, instance *SliceOfPointerToGongStructField) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.GongStruct, stage.GongStructs_reference, instance.GongStruct)
	// insertion point for slice of pointers field
}

func (reference *StageSetField) GongReconstructPointersFromReferences(stage *Stage, instance *StageSetField) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *StageSetModel) GongReconstructPointersFromReferences(stage *Stage, instance *StageSetModel) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.Fields, stage.StageSetFields_reference, instance.Fields)
}

// insertion point for pointer reconstruction from instances
func (reference *GongBasicField) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.GongEnum, stage.GongEnums_instance)
	// insertion point for slice of pointers fields
}

func (reference *GongEnum) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.GongEnumValues, stage.GongEnumValues_instance)
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
	__gong__reconstructSliceOfPointersFromInstances(&reference.Links, stage.GongLinks_instance)
}

func (reference *GongStruct) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.GongBasicFields, stage.GongBasicFields_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.GongTimeFields, stage.GongTimeFields_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.PointerToGongStructFields, stage.PointerToGongStructFields_instance)
	__gong__reconstructSliceOfPointersFromInstances(&reference.SliceOfPointerToGongStructFields, stage.SliceOfPointerToGongStructFields_instance)
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
	__gong__reconstructPointerFromInstance(&reference.GongStruct, stage.GongStructs_instance)
	// insertion point for slice of pointers fields
}

func (reference *SliceOfPointerToGongStructField) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.GongStruct, stage.GongStructs_instance)
	// insertion point for slice of pointers fields
}

func (reference *StageSetField) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *StageSetModel) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.Fields, stage.StageSetFields_instance)
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
	if gongbasicfield.GongEnum != gongbasicfieldOther.GongEnum {
		diffs = append(diffs, gongbasicfield.GongMarshallField(stage, "GongEnum"))
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
	if ops := __gong__diffSliceOfPointers(stage, gongenum, "GongEnumValues", gongenumOther.GongEnumValues, gongenum.GongEnumValues); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, gongnote, "Links", gongnoteOther.Links, gongnote.Links); ops != "" {
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
	if ops := __gong__diffSliceOfPointers(stage, gongstruct, "GongBasicFields", gongstructOther.GongBasicFields, gongstruct.GongBasicFields); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, gongstruct, "GongTimeFields", gongstructOther.GongTimeFields, gongstruct.GongTimeFields); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, gongstruct, "PointerToGongStructFields", gongstructOther.PointerToGongStructFields, gongstruct.PointerToGongStructFields); ops != "" {
		diffs = append(diffs, ops)
	}
	if ops := __gong__diffSliceOfPointers(stage, gongstruct, "SliceOfPointerToGongStructFields", gongstructOther.SliceOfPointerToGongStructFields, gongstruct.SliceOfPointerToGongStructFields); ops != "" {
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
	if pointertogongstructfield.GongStruct != pointertogongstructfieldOther.GongStruct {
		diffs = append(diffs, pointertogongstructfield.GongMarshallField(stage, "GongStruct"))
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
	if sliceofpointertogongstructfield.GongStruct != sliceofpointertogongstructfieldOther.GongStruct {
		diffs = append(diffs, sliceofpointertogongstructfield.GongMarshallField(stage, "GongStruct"))
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
	if ops := __gong__diffSliceOfPointers(stage, stagesetmodel, "Fields", stagesetmodelOther.Fields, stagesetmodel.Fields); ops != "" {
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
