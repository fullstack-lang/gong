// generated code - do not edit
package models

func IsStagedPointerToGongstruct[Type PointerToGongstruct](stage *Stage, instance Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *GongBasicField:
		ok = stage.IsStagedGongBasicField(target)

	case *GongEnum:
		ok = stage.IsStagedGongEnum(target)

	case *GongEnumValue:
		ok = stage.IsStagedGongEnumValue(target)

	case *GongLink:
		ok = stage.IsStagedGongLink(target)

	case *GongNote:
		ok = stage.IsStagedGongNote(target)

	case *GongStruct:
		ok = stage.IsStagedGongStruct(target)

	case *GongTimeField:
		ok = stage.IsStagedGongTimeField(target)

	case *MetaReference:
		ok = stage.IsStagedMetaReference(target)

	case *ModelPkg:
		ok = stage.IsStagedModelPkg(target)

	case *PointerToGongStructField:
		ok = stage.IsStagedPointerToGongStructField(target)

	case *SliceOfPointerToGongStructField:
		ok = stage.IsStagedSliceOfPointerToGongStructField(target)

	default:
		_ = target
	}
	return
}

func IsStaged[Type Gongstruct](stage *Stage, instance *Type) (ok bool) {

	switch target := any(instance).(type) {
	// insertion point for stage
	case *GongBasicField:
		ok = stage.IsStagedGongBasicField(target)

	case *GongEnum:
		ok = stage.IsStagedGongEnum(target)

	case *GongEnumValue:
		ok = stage.IsStagedGongEnumValue(target)

	case *GongLink:
		ok = stage.IsStagedGongLink(target)

	case *GongNote:
		ok = stage.IsStagedGongNote(target)

	case *GongStruct:
		ok = stage.IsStagedGongStruct(target)

	case *GongTimeField:
		ok = stage.IsStagedGongTimeField(target)

	case *MetaReference:
		ok = stage.IsStagedMetaReference(target)

	case *ModelPkg:
		ok = stage.IsStagedModelPkg(target)

	case *PointerToGongStructField:
		ok = stage.IsStagedPointerToGongStructField(target)

	case *SliceOfPointerToGongStructField:
		ok = stage.IsStagedSliceOfPointerToGongStructField(target)

	default:
		_ = target
	}
	return
}

// insertion point for stage per struct
func (stage *Stage) IsStagedGongBasicField(gongbasicfield *GongBasicField) (ok bool) {

	_, ok = stage.GongBasicFields[gongbasicfield]

	return
}

func (stage *Stage) IsStagedGongEnum(gongenum *GongEnum) (ok bool) {

	_, ok = stage.GongEnums[gongenum]

	return
}

func (stage *Stage) IsStagedGongEnumValue(gongenumvalue *GongEnumValue) (ok bool) {

	_, ok = stage.GongEnumValues[gongenumvalue]

	return
}

func (stage *Stage) IsStagedGongLink(gonglink *GongLink) (ok bool) {

	_, ok = stage.GongLinks[gonglink]

	return
}

func (stage *Stage) IsStagedGongNote(gongnote *GongNote) (ok bool) {

	_, ok = stage.GongNotes[gongnote]

	return
}

func (stage *Stage) IsStagedGongStruct(gongstruct *GongStruct) (ok bool) {

	_, ok = stage.GongStructs[gongstruct]

	return
}

func (stage *Stage) IsStagedGongTimeField(gongtimefield *GongTimeField) (ok bool) {

	_, ok = stage.GongTimeFields[gongtimefield]

	return
}

func (stage *Stage) IsStagedMetaReference(metareference *MetaReference) (ok bool) {

	_, ok = stage.MetaReferences[metareference]

	return
}

func (stage *Stage) IsStagedModelPkg(modelpkg *ModelPkg) (ok bool) {

	_, ok = stage.ModelPkgs[modelpkg]

	return
}

func (stage *Stage) IsStagedPointerToGongStructField(pointertogongstructfield *PointerToGongStructField) (ok bool) {

	_, ok = stage.PointerToGongStructFields[pointertogongstructfield]

	return
}

func (stage *Stage) IsStagedSliceOfPointerToGongStructField(sliceofpointertogongstructfield *SliceOfPointerToGongStructField) (ok bool) {

	_, ok = stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point for stage branch
	case *GongBasicField:
		stage.StageBranchGongBasicField(target)

	case *GongEnum:
		stage.StageBranchGongEnum(target)

	case *GongEnumValue:
		stage.StageBranchGongEnumValue(target)

	case *GongLink:
		stage.StageBranchGongLink(target)

	case *GongNote:
		stage.StageBranchGongNote(target)

	case *GongStruct:
		stage.StageBranchGongStruct(target)

	case *GongTimeField:
		stage.StageBranchGongTimeField(target)

	case *MetaReference:
		stage.StageBranchMetaReference(target)

	case *ModelPkg:
		stage.StageBranchModelPkg(target)

	case *PointerToGongStructField:
		stage.StageBranchPointerToGongStructField(target)

	case *SliceOfPointerToGongStructField:
		stage.StageBranchSliceOfPointerToGongStructField(target)

	default:
		_ = target
	}
}

// insertion point for stage branch per struct
func (stage *Stage) StageBranchGongBasicField(gongbasicfield *GongBasicField) {

	// check if instance is already staged
	if IsStaged(stage, gongbasicfield) {
		return
	}

	gongbasicfield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if gongbasicfield.GongEnum != nil {
		StageBranch(stage, gongbasicfield.GongEnum)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGongEnum(gongenum *GongEnum) {

	// check if instance is already staged
	if IsStaged(stage, gongenum) {
		return
	}

	gongenum.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalue := range gongenum.GongEnumValues {
		StageBranch(stage, _gongenumvalue)
	}

}

func (stage *Stage) StageBranchGongEnumValue(gongenumvalue *GongEnumValue) {

	// check if instance is already staged
	if IsStaged(stage, gongenumvalue) {
		return
	}

	gongenumvalue.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGongLink(gonglink *GongLink) {

	// check if instance is already staged
	if IsStaged(stage, gonglink) {
		return
	}

	gonglink.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchGongNote(gongnote *GongNote) {

	// check if instance is already staged
	if IsStaged(stage, gongnote) {
		return
	}

	gongnote.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gonglink := range gongnote.Links {
		StageBranch(stage, _gonglink)
	}

}

func (stage *Stage) StageBranchGongStruct(gongstruct *GongStruct) {

	// check if instance is already staged
	if IsStaged(stage, gongstruct) {
		return
	}

	gongstruct.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongbasicfield := range gongstruct.GongBasicFields {
		StageBranch(stage, _gongbasicfield)
	}
	for _, _gongtimefield := range gongstruct.GongTimeFields {
		StageBranch(stage, _gongtimefield)
	}
	for _, _pointertogongstructfield := range gongstruct.PointerToGongStructFields {
		StageBranch(stage, _pointertogongstructfield)
	}
	for _, _sliceofpointertogongstructfield := range gongstruct.SliceOfPointerToGongStructFields {
		StageBranch(stage, _sliceofpointertogongstructfield)
	}

}

func (stage *Stage) StageBranchGongTimeField(gongtimefield *GongTimeField) {

	// check if instance is already staged
	if IsStaged(stage, gongtimefield) {
		return
	}

	gongtimefield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchMetaReference(metareference *MetaReference) {

	// check if instance is already staged
	if IsStaged(stage, metareference) {
		return
	}

	metareference.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchModelPkg(modelpkg *ModelPkg) {

	// check if instance is already staged
	if IsStaged(stage, modelpkg) {
		return
	}

	modelpkg.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchPointerToGongStructField(pointertogongstructfield *PointerToGongStructField) {

	// check if instance is already staged
	if IsStaged(stage, pointertogongstructfield) {
		return
	}

	pointertogongstructfield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pointertogongstructfield.GongStruct != nil {
		StageBranch(stage, pointertogongstructfield.GongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) StageBranchSliceOfPointerToGongStructField(sliceofpointertogongstructfield *SliceOfPointerToGongStructField) {

	// check if instance is already staged
	if IsStaged(stage, sliceofpointertogongstructfield) {
		return
	}

	sliceofpointertogongstructfield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sliceofpointertogongstructfield.GongStruct != nil {
		StageBranch(stage, sliceofpointertogongstructfield.GongStruct)
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
	case *GongBasicField:
		toT := CopyBranchGongBasicField(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongEnum:
		toT := CopyBranchGongEnum(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongEnumValue:
		toT := CopyBranchGongEnumValue(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongLink:
		toT := CopyBranchGongLink(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongNote:
		toT := CopyBranchGongNote(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongStruct:
		toT := CopyBranchGongStruct(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *GongTimeField:
		toT := CopyBranchGongTimeField(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MetaReference:
		toT := CopyBranchMetaReference(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ModelPkg:
		toT := CopyBranchModelPkg(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PointerToGongStructField:
		toT := CopyBranchPointerToGongStructField(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SliceOfPointerToGongStructField:
		toT := CopyBranchSliceOfPointerToGongStructField(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func CopyBranchGongBasicField(mapOrigCopy map[any]any, gongbasicfieldFrom *GongBasicField) (gongbasicfieldTo *GongBasicField) {

	// gongbasicfieldFrom has already been copied
	if _gongbasicfieldTo, ok := mapOrigCopy[gongbasicfieldFrom]; ok {
		gongbasicfieldTo = _gongbasicfieldTo.(*GongBasicField)
		return
	}

	gongbasicfieldTo = new(GongBasicField)
	mapOrigCopy[gongbasicfieldFrom] = gongbasicfieldTo
	gongbasicfieldFrom.CopyBasicFields(gongbasicfieldTo)

	//insertion point for the staging of instances referenced by pointers
	if gongbasicfieldFrom.GongEnum != nil {
		gongbasicfieldTo.GongEnum = CopyBranchGongEnum(mapOrigCopy, gongbasicfieldFrom.GongEnum)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGongEnum(mapOrigCopy map[any]any, gongenumFrom *GongEnum) (gongenumTo *GongEnum) {

	// gongenumFrom has already been copied
	if _gongenumTo, ok := mapOrigCopy[gongenumFrom]; ok {
		gongenumTo = _gongenumTo.(*GongEnum)
		return
	}

	gongenumTo = new(GongEnum)
	mapOrigCopy[gongenumFrom] = gongenumTo
	gongenumFrom.CopyBasicFields(gongenumTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalue := range gongenumFrom.GongEnumValues {
		gongenumTo.GongEnumValues = append(gongenumTo.GongEnumValues, CopyBranchGongEnumValue(mapOrigCopy, _gongenumvalue))
	}

	return
}

func CopyBranchGongEnumValue(mapOrigCopy map[any]any, gongenumvalueFrom *GongEnumValue) (gongenumvalueTo *GongEnumValue) {

	// gongenumvalueFrom has already been copied
	if _gongenumvalueTo, ok := mapOrigCopy[gongenumvalueFrom]; ok {
		gongenumvalueTo = _gongenumvalueTo.(*GongEnumValue)
		return
	}

	gongenumvalueTo = new(GongEnumValue)
	mapOrigCopy[gongenumvalueFrom] = gongenumvalueTo
	gongenumvalueFrom.CopyBasicFields(gongenumvalueTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGongLink(mapOrigCopy map[any]any, gonglinkFrom *GongLink) (gonglinkTo *GongLink) {

	// gonglinkFrom has already been copied
	if _gonglinkTo, ok := mapOrigCopy[gonglinkFrom]; ok {
		gonglinkTo = _gonglinkTo.(*GongLink)
		return
	}

	gonglinkTo = new(GongLink)
	mapOrigCopy[gonglinkFrom] = gonglinkTo
	gonglinkFrom.CopyBasicFields(gonglinkTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchGongNote(mapOrigCopy map[any]any, gongnoteFrom *GongNote) (gongnoteTo *GongNote) {

	// gongnoteFrom has already been copied
	if _gongnoteTo, ok := mapOrigCopy[gongnoteFrom]; ok {
		gongnoteTo = _gongnoteTo.(*GongNote)
		return
	}

	gongnoteTo = new(GongNote)
	mapOrigCopy[gongnoteFrom] = gongnoteTo
	gongnoteFrom.CopyBasicFields(gongnoteTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gonglink := range gongnoteFrom.Links {
		gongnoteTo.Links = append(gongnoteTo.Links, CopyBranchGongLink(mapOrigCopy, _gonglink))
	}

	return
}

func CopyBranchGongStruct(mapOrigCopy map[any]any, gongstructFrom *GongStruct) (gongstructTo *GongStruct) {

	// gongstructFrom has already been copied
	if _gongstructTo, ok := mapOrigCopy[gongstructFrom]; ok {
		gongstructTo = _gongstructTo.(*GongStruct)
		return
	}

	gongstructTo = new(GongStruct)
	mapOrigCopy[gongstructFrom] = gongstructTo
	gongstructFrom.CopyBasicFields(gongstructTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongbasicfield := range gongstructFrom.GongBasicFields {
		gongstructTo.GongBasicFields = append(gongstructTo.GongBasicFields, CopyBranchGongBasicField(mapOrigCopy, _gongbasicfield))
	}
	for _, _gongtimefield := range gongstructFrom.GongTimeFields {
		gongstructTo.GongTimeFields = append(gongstructTo.GongTimeFields, CopyBranchGongTimeField(mapOrigCopy, _gongtimefield))
	}
	for _, _pointertogongstructfield := range gongstructFrom.PointerToGongStructFields {
		gongstructTo.PointerToGongStructFields = append(gongstructTo.PointerToGongStructFields, CopyBranchPointerToGongStructField(mapOrigCopy, _pointertogongstructfield))
	}
	for _, _sliceofpointertogongstructfield := range gongstructFrom.SliceOfPointerToGongStructFields {
		gongstructTo.SliceOfPointerToGongStructFields = append(gongstructTo.SliceOfPointerToGongStructFields, CopyBranchSliceOfPointerToGongStructField(mapOrigCopy, _sliceofpointertogongstructfield))
	}

	return
}

func CopyBranchGongTimeField(mapOrigCopy map[any]any, gongtimefieldFrom *GongTimeField) (gongtimefieldTo *GongTimeField) {

	// gongtimefieldFrom has already been copied
	if _gongtimefieldTo, ok := mapOrigCopy[gongtimefieldFrom]; ok {
		gongtimefieldTo = _gongtimefieldTo.(*GongTimeField)
		return
	}

	gongtimefieldTo = new(GongTimeField)
	mapOrigCopy[gongtimefieldFrom] = gongtimefieldTo
	gongtimefieldFrom.CopyBasicFields(gongtimefieldTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchMetaReference(mapOrigCopy map[any]any, metareferenceFrom *MetaReference) (metareferenceTo *MetaReference) {

	// metareferenceFrom has already been copied
	if _metareferenceTo, ok := mapOrigCopy[metareferenceFrom]; ok {
		metareferenceTo = _metareferenceTo.(*MetaReference)
		return
	}

	metareferenceTo = new(MetaReference)
	mapOrigCopy[metareferenceFrom] = metareferenceTo
	metareferenceFrom.CopyBasicFields(metareferenceTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchModelPkg(mapOrigCopy map[any]any, modelpkgFrom *ModelPkg) (modelpkgTo *ModelPkg) {

	// modelpkgFrom has already been copied
	if _modelpkgTo, ok := mapOrigCopy[modelpkgFrom]; ok {
		modelpkgTo = _modelpkgTo.(*ModelPkg)
		return
	}

	modelpkgTo = new(ModelPkg)
	mapOrigCopy[modelpkgFrom] = modelpkgTo
	modelpkgFrom.CopyBasicFields(modelpkgTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchPointerToGongStructField(mapOrigCopy map[any]any, pointertogongstructfieldFrom *PointerToGongStructField) (pointertogongstructfieldTo *PointerToGongStructField) {

	// pointertogongstructfieldFrom has already been copied
	if _pointertogongstructfieldTo, ok := mapOrigCopy[pointertogongstructfieldFrom]; ok {
		pointertogongstructfieldTo = _pointertogongstructfieldTo.(*PointerToGongStructField)
		return
	}

	pointertogongstructfieldTo = new(PointerToGongStructField)
	mapOrigCopy[pointertogongstructfieldFrom] = pointertogongstructfieldTo
	pointertogongstructfieldFrom.CopyBasicFields(pointertogongstructfieldTo)

	//insertion point for the staging of instances referenced by pointers
	if pointertogongstructfieldFrom.GongStruct != nil {
		pointertogongstructfieldTo.GongStruct = CopyBranchGongStruct(mapOrigCopy, pointertogongstructfieldFrom.GongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func CopyBranchSliceOfPointerToGongStructField(mapOrigCopy map[any]any, sliceofpointertogongstructfieldFrom *SliceOfPointerToGongStructField) (sliceofpointertogongstructfieldTo *SliceOfPointerToGongStructField) {

	// sliceofpointertogongstructfieldFrom has already been copied
	if _sliceofpointertogongstructfieldTo, ok := mapOrigCopy[sliceofpointertogongstructfieldFrom]; ok {
		sliceofpointertogongstructfieldTo = _sliceofpointertogongstructfieldTo.(*SliceOfPointerToGongStructField)
		return
	}

	sliceofpointertogongstructfieldTo = new(SliceOfPointerToGongStructField)
	mapOrigCopy[sliceofpointertogongstructfieldFrom] = sliceofpointertogongstructfieldTo
	sliceofpointertogongstructfieldFrom.CopyBasicFields(sliceofpointertogongstructfieldTo)

	//insertion point for the staging of instances referenced by pointers
	if sliceofpointertogongstructfieldFrom.GongStruct != nil {
		sliceofpointertogongstructfieldTo.GongStruct = CopyBranchGongStruct(mapOrigCopy, sliceofpointertogongstructfieldFrom.GongStruct)
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
	case *GongBasicField:
		stage.UnstageBranchGongBasicField(target)

	case *GongEnum:
		stage.UnstageBranchGongEnum(target)

	case *GongEnumValue:
		stage.UnstageBranchGongEnumValue(target)

	case *GongLink:
		stage.UnstageBranchGongLink(target)

	case *GongNote:
		stage.UnstageBranchGongNote(target)

	case *GongStruct:
		stage.UnstageBranchGongStruct(target)

	case *GongTimeField:
		stage.UnstageBranchGongTimeField(target)

	case *MetaReference:
		stage.UnstageBranchMetaReference(target)

	case *ModelPkg:
		stage.UnstageBranchModelPkg(target)

	case *PointerToGongStructField:
		stage.UnstageBranchPointerToGongStructField(target)

	case *SliceOfPointerToGongStructField:
		stage.UnstageBranchSliceOfPointerToGongStructField(target)

	default:
		_ = target
	}
}

// insertion point for unstage branch per struct
func (stage *Stage) UnstageBranchGongBasicField(gongbasicfield *GongBasicField) {

	// check if instance is already staged
	if !IsStaged(stage, gongbasicfield) {
		return
	}

	gongbasicfield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if gongbasicfield.GongEnum != nil {
		UnstageBranch(stage, gongbasicfield.GongEnum)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGongEnum(gongenum *GongEnum) {

	// check if instance is already staged
	if !IsStaged(stage, gongenum) {
		return
	}

	gongenum.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongenumvalue := range gongenum.GongEnumValues {
		UnstageBranch(stage, _gongenumvalue)
	}

}

func (stage *Stage) UnstageBranchGongEnumValue(gongenumvalue *GongEnumValue) {

	// check if instance is already staged
	if !IsStaged(stage, gongenumvalue) {
		return
	}

	gongenumvalue.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGongLink(gonglink *GongLink) {

	// check if instance is already staged
	if !IsStaged(stage, gonglink) {
		return
	}

	gonglink.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchGongNote(gongnote *GongNote) {

	// check if instance is already staged
	if !IsStaged(stage, gongnote) {
		return
	}

	gongnote.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gonglink := range gongnote.Links {
		UnstageBranch(stage, _gonglink)
	}

}

func (stage *Stage) UnstageBranchGongStruct(gongstruct *GongStruct) {

	// check if instance is already staged
	if !IsStaged(stage, gongstruct) {
		return
	}

	gongstruct.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _gongbasicfield := range gongstruct.GongBasicFields {
		UnstageBranch(stage, _gongbasicfield)
	}
	for _, _gongtimefield := range gongstruct.GongTimeFields {
		UnstageBranch(stage, _gongtimefield)
	}
	for _, _pointertogongstructfield := range gongstruct.PointerToGongStructFields {
		UnstageBranch(stage, _pointertogongstructfield)
	}
	for _, _sliceofpointertogongstructfield := range gongstruct.SliceOfPointerToGongStructFields {
		UnstageBranch(stage, _sliceofpointertogongstructfield)
	}

}

func (stage *Stage) UnstageBranchGongTimeField(gongtimefield *GongTimeField) {

	// check if instance is already staged
	if !IsStaged(stage, gongtimefield) {
		return
	}

	gongtimefield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchMetaReference(metareference *MetaReference) {

	// check if instance is already staged
	if !IsStaged(stage, metareference) {
		return
	}

	metareference.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchModelPkg(modelpkg *ModelPkg) {

	// check if instance is already staged
	if !IsStaged(stage, modelpkg) {
		return
	}

	modelpkg.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchPointerToGongStructField(pointertogongstructfield *PointerToGongStructField) {

	// check if instance is already staged
	if !IsStaged(stage, pointertogongstructfield) {
		return
	}

	pointertogongstructfield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pointertogongstructfield.GongStruct != nil {
		UnstageBranch(stage, pointertogongstructfield.GongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *Stage) UnstageBranchSliceOfPointerToGongStructField(sliceofpointertogongstructfield *SliceOfPointerToGongStructField) {

	// check if instance is already staged
	if !IsStaged(stage, sliceofpointertogongstructfield) {
		return
	}

	sliceofpointertogongstructfield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if sliceofpointertogongstructfield.GongStruct != nil {
		UnstageBranch(stage, sliceofpointertogongstructfield.GongStruct)
	}

	//insertion point for the staging of instances referenced by slice of pointers

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
		diffs = append(diffs, "GongEnum")
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
		diffs = append(diffs, gongenum.GongMarshallField(stage, "GongEnumValues"))
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
		diffs = append(diffs, gongnote.GongMarshallField(stage, "Links"))
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
		diffs = append(diffs, gongstruct.GongMarshallField(stage, "GongBasicFields"))
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
		diffs = append(diffs, gongstruct.GongMarshallField(stage, "GongTimeFields"))
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
		diffs = append(diffs, gongstruct.GongMarshallField(stage, "PointerToGongStructFields"))
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
		diffs = append(diffs, gongstruct.GongMarshallField(stage, "SliceOfPointerToGongStructFields"))
	}
	if gongstruct.HasOnAfterUpdateSignature != gongstructOther.HasOnAfterUpdateSignature {
		diffs = append(diffs, gongstruct.GongMarshallField(stage, "HasOnAfterUpdateSignature"))
	}
	if gongstruct.IsIgnoredForFront != gongstructOther.IsIgnoredForFront {
		diffs = append(diffs, gongstruct.GongMarshallField(stage, "IsIgnoredForFront"))
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
	if gongtimefield.BespokeTimeFormat != gongtimefieldOther.BespokeTimeFormat {
		diffs = append(diffs, gongtimefield.GongMarshallField(stage, "BespokeTimeFormat"))
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
		diffs = append(diffs, "GongStruct")
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
		diffs = append(diffs, "GongStruct")
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

	return
}
