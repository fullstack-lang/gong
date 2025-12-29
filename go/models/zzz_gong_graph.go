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
func (gongbasicfield *GongBasicField) GongDiff(gongbasicfieldOther *GongBasicField) (diffs []string) {
	// insertion point for field diffs
	if gongbasicfield.Name != gongbasicfieldOther.Name {
		diffs = append(diffs, "Name")
	}
	if gongbasicfield.BasicKindName != gongbasicfieldOther.BasicKindName {
		diffs = append(diffs, "BasicKindName")
	}
	if (gongbasicfield.GongEnum == nil) != (gongbasicfieldOther.GongEnum == nil) {
		diffs = append(diffs, "GongEnum")
	} else if gongbasicfield.GongEnum != nil && gongbasicfieldOther.GongEnum != nil {
		if gongbasicfield.GongEnum != gongbasicfieldOther.GongEnum {
			diffs = append(diffs, "GongEnum")
		}
	}
	if gongbasicfield.DeclaredType != gongbasicfieldOther.DeclaredType {
		diffs = append(diffs, "DeclaredType")
	}
	if gongbasicfield.CompositeStructName != gongbasicfieldOther.CompositeStructName {
		diffs = append(diffs, "CompositeStructName")
	}
	if gongbasicfield.Index != gongbasicfieldOther.Index {
		diffs = append(diffs, "Index")
	}
	if gongbasicfield.IsTextArea != gongbasicfieldOther.IsTextArea {
		diffs = append(diffs, "IsTextArea")
	}
	if gongbasicfield.IsBespokeWidth != gongbasicfieldOther.IsBespokeWidth {
		diffs = append(diffs, "IsBespokeWidth")
	}
	if gongbasicfield.BespokeWidth != gongbasicfieldOther.BespokeWidth {
		diffs = append(diffs, "BespokeWidth")
	}
	if gongbasicfield.IsBespokeHeight != gongbasicfieldOther.IsBespokeHeight {
		diffs = append(diffs, "IsBespokeHeight")
	}
	if gongbasicfield.BespokeHeight != gongbasicfieldOther.BespokeHeight {
		diffs = append(diffs, "BespokeHeight")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gongenum *GongEnum) GongDiff(gongenumOther *GongEnum) (diffs []string) {
	// insertion point for field diffs
	if gongenum.Name != gongenumOther.Name {
		diffs = append(diffs, "Name")
	}
	if gongenum.Type != gongenumOther.Type {
		diffs = append(diffs, "Type")
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
				if len(gongenum.GongEnumValues[i].GongDiff(gongenumOther.GongEnumValues[i])) > 0 {
					GongEnumValuesDifferent = true
					break
				}
			}
		}
	}
	if GongEnumValuesDifferent {
		diffs = append(diffs, "GongEnumValues")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gongenumvalue *GongEnumValue) GongDiff(gongenumvalueOther *GongEnumValue) (diffs []string) {
	// insertion point for field diffs
	if gongenumvalue.Name != gongenumvalueOther.Name {
		diffs = append(diffs, "Name")
	}
	if gongenumvalue.Value != gongenumvalueOther.Value {
		diffs = append(diffs, "Value")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gonglink *GongLink) GongDiff(gonglinkOther *GongLink) (diffs []string) {
	// insertion point for field diffs
	if gonglink.Name != gonglinkOther.Name {
		diffs = append(diffs, "Name")
	}
	if gonglink.Recv != gonglinkOther.Recv {
		diffs = append(diffs, "Recv")
	}
	if gonglink.ImportPath != gonglinkOther.ImportPath {
		diffs = append(diffs, "ImportPath")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gongnote *GongNote) GongDiff(gongnoteOther *GongNote) (diffs []string) {
	// insertion point for field diffs
	if gongnote.Name != gongnoteOther.Name {
		diffs = append(diffs, "Name")
	}
	if gongnote.Body != gongnoteOther.Body {
		diffs = append(diffs, "Body")
	}
	if gongnote.BodyHTML != gongnoteOther.BodyHTML {
		diffs = append(diffs, "BodyHTML")
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
				if len(gongnote.Links[i].GongDiff(gongnoteOther.Links[i])) > 0 {
					LinksDifferent = true
					break
				}
			}
		}
	}
	if LinksDifferent {
		diffs = append(diffs, "Links")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gongstruct *GongStruct) GongDiff(gongstructOther *GongStruct) (diffs []string) {
	// insertion point for field diffs
	if gongstruct.Name != gongstructOther.Name {
		diffs = append(diffs, "Name")
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
				if len(gongstruct.GongBasicFields[i].GongDiff(gongstructOther.GongBasicFields[i])) > 0 {
					GongBasicFieldsDifferent = true
					break
				}
			}
		}
	}
	if GongBasicFieldsDifferent {
		diffs = append(diffs, "GongBasicFields")
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
				if len(gongstruct.GongTimeFields[i].GongDiff(gongstructOther.GongTimeFields[i])) > 0 {
					GongTimeFieldsDifferent = true
					break
				}
			}
		}
	}
	if GongTimeFieldsDifferent {
		diffs = append(diffs, "GongTimeFields")
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
				if len(gongstruct.PointerToGongStructFields[i].GongDiff(gongstructOther.PointerToGongStructFields[i])) > 0 {
					PointerToGongStructFieldsDifferent = true
					break
				}
			}
		}
	}
	if PointerToGongStructFieldsDifferent {
		diffs = append(diffs, "PointerToGongStructFields")
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
				if len(gongstruct.SliceOfPointerToGongStructFields[i].GongDiff(gongstructOther.SliceOfPointerToGongStructFields[i])) > 0 {
					SliceOfPointerToGongStructFieldsDifferent = true
					break
				}
			}
		}
	}
	if SliceOfPointerToGongStructFieldsDifferent {
		diffs = append(diffs, "SliceOfPointerToGongStructFields")
	}
	if gongstruct.HasOnAfterUpdateSignature != gongstructOther.HasOnAfterUpdateSignature {
		diffs = append(diffs, "HasOnAfterUpdateSignature")
	}
	if gongstruct.IsIgnoredForFront != gongstructOther.IsIgnoredForFront {
		diffs = append(diffs, "IsIgnoredForFront")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (gongtimefield *GongTimeField) GongDiff(gongtimefieldOther *GongTimeField) (diffs []string) {
	// insertion point for field diffs
	if gongtimefield.Name != gongtimefieldOther.Name {
		diffs = append(diffs, "Name")
	}
	if gongtimefield.Index != gongtimefieldOther.Index {
		diffs = append(diffs, "Index")
	}
	if gongtimefield.CompositeStructName != gongtimefieldOther.CompositeStructName {
		diffs = append(diffs, "CompositeStructName")
	}
	if gongtimefield.BespokeTimeFormat != gongtimefieldOther.BespokeTimeFormat {
		diffs = append(diffs, "BespokeTimeFormat")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (metareference *MetaReference) GongDiff(metareferenceOther *MetaReference) (diffs []string) {
	// insertion point for field diffs
	if metareference.Name != metareferenceOther.Name {
		diffs = append(diffs, "Name")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (modelpkg *ModelPkg) GongDiff(modelpkgOther *ModelPkg) (diffs []string) {
	// insertion point for field diffs
	if modelpkg.Name != modelpkgOther.Name {
		diffs = append(diffs, "Name")
	}
	if modelpkg.PkgPath != modelpkgOther.PkgPath {
		diffs = append(diffs, "PkgPath")
	}
	if modelpkg.PathToGoSubDirectory != modelpkgOther.PathToGoSubDirectory {
		diffs = append(diffs, "PathToGoSubDirectory")
	}
	if modelpkg.OrmPkgGenPath != modelpkgOther.OrmPkgGenPath {
		diffs = append(diffs, "OrmPkgGenPath")
	}
	if modelpkg.DbOrmPkgGenPath != modelpkgOther.DbOrmPkgGenPath {
		diffs = append(diffs, "DbOrmPkgGenPath")
	}
	if modelpkg.DbLiteOrmPkgGenPath != modelpkgOther.DbLiteOrmPkgGenPath {
		diffs = append(diffs, "DbLiteOrmPkgGenPath")
	}
	if modelpkg.DbPkgGenPath != modelpkgOther.DbPkgGenPath {
		diffs = append(diffs, "DbPkgGenPath")
	}
	if modelpkg.ControllersPkgGenPath != modelpkgOther.ControllersPkgGenPath {
		diffs = append(diffs, "ControllersPkgGenPath")
	}
	if modelpkg.FullstackPkgGenPath != modelpkgOther.FullstackPkgGenPath {
		diffs = append(diffs, "FullstackPkgGenPath")
	}
	if modelpkg.StackPkgGenPath != modelpkgOther.StackPkgGenPath {
		diffs = append(diffs, "StackPkgGenPath")
	}
	if modelpkg.Level1StackPkgGenPath != modelpkgOther.Level1StackPkgGenPath {
		diffs = append(diffs, "Level1StackPkgGenPath")
	}
	if modelpkg.StaticPkgGenPath != modelpkgOther.StaticPkgGenPath {
		diffs = append(diffs, "StaticPkgGenPath")
	}
	if modelpkg.ProbePkgGenPath != modelpkgOther.ProbePkgGenPath {
		diffs = append(diffs, "ProbePkgGenPath")
	}
	if modelpkg.NgWorkspacePath != modelpkgOther.NgWorkspacePath {
		diffs = append(diffs, "NgWorkspacePath")
	}
	if modelpkg.NgWorkspaceName != modelpkgOther.NgWorkspaceName {
		diffs = append(diffs, "NgWorkspaceName")
	}
	if modelpkg.NgDataLibrarySourceCodeDirectory != modelpkgOther.NgDataLibrarySourceCodeDirectory {
		diffs = append(diffs, "NgDataLibrarySourceCodeDirectory")
	}
	if modelpkg.NgSpecificLibrarySourceCodeDirectory != modelpkgOther.NgSpecificLibrarySourceCodeDirectory {
		diffs = append(diffs, "NgSpecificLibrarySourceCodeDirectory")
	}
	if modelpkg.MaterialLibDatamodelTargetPath != modelpkgOther.MaterialLibDatamodelTargetPath {
		diffs = append(diffs, "MaterialLibDatamodelTargetPath")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pointertogongstructfield *PointerToGongStructField) GongDiff(pointertogongstructfieldOther *PointerToGongStructField) (diffs []string) {
	// insertion point for field diffs
	if pointertogongstructfield.Name != pointertogongstructfieldOther.Name {
		diffs = append(diffs, "Name")
	}
	if (pointertogongstructfield.GongStruct == nil) != (pointertogongstructfieldOther.GongStruct == nil) {
		diffs = append(diffs, "GongStruct")
	} else if pointertogongstructfield.GongStruct != nil && pointertogongstructfieldOther.GongStruct != nil {
		if pointertogongstructfield.GongStruct != pointertogongstructfieldOther.GongStruct {
			diffs = append(diffs, "GongStruct")
		}
	}
	if pointertogongstructfield.Index != pointertogongstructfieldOther.Index {
		diffs = append(diffs, "Index")
	}
	if pointertogongstructfield.CompositeStructName != pointertogongstructfieldOther.CompositeStructName {
		diffs = append(diffs, "CompositeStructName")
	}
	if pointertogongstructfield.IsType != pointertogongstructfieldOther.IsType {
		diffs = append(diffs, "IsType")
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (sliceofpointertogongstructfield *SliceOfPointerToGongStructField) GongDiff(sliceofpointertogongstructfieldOther *SliceOfPointerToGongStructField) (diffs []string) {
	// insertion point for field diffs
	if sliceofpointertogongstructfield.Name != sliceofpointertogongstructfieldOther.Name {
		diffs = append(diffs, "Name")
	}
	if (sliceofpointertogongstructfield.GongStruct == nil) != (sliceofpointertogongstructfieldOther.GongStruct == nil) {
		diffs = append(diffs, "GongStruct")
	} else if sliceofpointertogongstructfield.GongStruct != nil && sliceofpointertogongstructfieldOther.GongStruct != nil {
		if sliceofpointertogongstructfield.GongStruct != sliceofpointertogongstructfieldOther.GongStruct {
			diffs = append(diffs, "GongStruct")
		}
	}
	if sliceofpointertogongstructfield.Index != sliceofpointertogongstructfieldOther.Index {
		diffs = append(diffs, "Index")
	}
	if sliceofpointertogongstructfield.CompositeStructName != sliceofpointertogongstructfieldOther.CompositeStructName {
		diffs = append(diffs, "CompositeStructName")
	}

	return
}
