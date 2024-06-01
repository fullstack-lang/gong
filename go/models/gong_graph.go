// generated code - do not edit
package models

func IsStaged[Type Gongstruct](stage *StageStruct, instance *Type) (ok bool) {

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

	case *Meta:
		ok = stage.IsStagedMeta(target)

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
func (stage *StageStruct) IsStagedGongBasicField(gongbasicfield *GongBasicField) (ok bool) {

	_, ok = stage.GongBasicFields[gongbasicfield]

	return
}

func (stage *StageStruct) IsStagedGongEnum(gongenum *GongEnum) (ok bool) {

	_, ok = stage.GongEnums[gongenum]

	return
}

func (stage *StageStruct) IsStagedGongEnumValue(gongenumvalue *GongEnumValue) (ok bool) {

	_, ok = stage.GongEnumValues[gongenumvalue]

	return
}

func (stage *StageStruct) IsStagedGongLink(gonglink *GongLink) (ok bool) {

	_, ok = stage.GongLinks[gonglink]

	return
}

func (stage *StageStruct) IsStagedGongNote(gongnote *GongNote) (ok bool) {

	_, ok = stage.GongNotes[gongnote]

	return
}

func (stage *StageStruct) IsStagedGongStruct(gongstruct *GongStruct) (ok bool) {

	_, ok = stage.GongStructs[gongstruct]

	return
}

func (stage *StageStruct) IsStagedGongTimeField(gongtimefield *GongTimeField) (ok bool) {

	_, ok = stage.GongTimeFields[gongtimefield]

	return
}

func (stage *StageStruct) IsStagedMeta(meta *Meta) (ok bool) {

	_, ok = stage.Metas[meta]

	return
}

func (stage *StageStruct) IsStagedMetaReference(metareference *MetaReference) (ok bool) {

	_, ok = stage.MetaReferences[metareference]

	return
}

func (stage *StageStruct) IsStagedModelPkg(modelpkg *ModelPkg) (ok bool) {

	_, ok = stage.ModelPkgs[modelpkg]

	return
}

func (stage *StageStruct) IsStagedPointerToGongStructField(pointertogongstructfield *PointerToGongStructField) (ok bool) {

	_, ok = stage.PointerToGongStructFields[pointertogongstructfield]

	return
}

func (stage *StageStruct) IsStagedSliceOfPointerToGongStructField(sliceofpointertogongstructfield *SliceOfPointerToGongStructField) (ok bool) {

	_, ok = stage.SliceOfPointerToGongStructFields[sliceofpointertogongstructfield]

	return
}

// StageBranch stages instance and apply StageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func StageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

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

	case *Meta:
		stage.StageBranchMeta(target)

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
func (stage *StageStruct) StageBranchGongBasicField(gongbasicfield *GongBasicField) {

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

func (stage *StageStruct) StageBranchGongEnum(gongenum *GongEnum) {

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

func (stage *StageStruct) StageBranchGongEnumValue(gongenumvalue *GongEnumValue) {

	// check if instance is already staged
	if IsStaged(stage, gongenumvalue) {
		return
	}

	gongenumvalue.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGongLink(gonglink *GongLink) {

	// check if instance is already staged
	if IsStaged(stage, gonglink) {
		return
	}

	gonglink.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchGongNote(gongnote *GongNote) {

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

func (stage *StageStruct) StageBranchGongStruct(gongstruct *GongStruct) {

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

func (stage *StageStruct) StageBranchGongTimeField(gongtimefield *GongTimeField) {

	// check if instance is already staged
	if IsStaged(stage, gongtimefield) {
		return
	}

	gongtimefield.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchMeta(meta *Meta) {

	// check if instance is already staged
	if IsStaged(stage, meta) {
		return
	}

	meta.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _metareference := range meta.MetaReferences {
		StageBranch(stage, _metareference)
	}

}

func (stage *StageStruct) StageBranchMetaReference(metareference *MetaReference) {

	// check if instance is already staged
	if IsStaged(stage, metareference) {
		return
	}

	metareference.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchModelPkg(modelpkg *ModelPkg) {

	// check if instance is already staged
	if IsStaged(stage, modelpkg) {
		return
	}

	modelpkg.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) StageBranchPointerToGongStructField(pointertogongstructfield *PointerToGongStructField) {

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

func (stage *StageStruct) StageBranchSliceOfPointerToGongStructField(sliceofpointertogongstructfield *SliceOfPointerToGongStructField) {

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

	case *Meta:
		toT := CopyBranchMeta(mapOrigCopy, fromT)
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

func CopyBranchMeta(mapOrigCopy map[any]any, metaFrom *Meta) (metaTo *Meta) {

	// metaFrom has already been copied
	if _metaTo, ok := mapOrigCopy[metaFrom]; ok {
		metaTo = _metaTo.(*Meta)
		return
	}

	metaTo = new(Meta)
	mapOrigCopy[metaFrom] = metaTo
	metaFrom.CopyBasicFields(metaTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _metareference := range metaFrom.MetaReferences {
		metaTo.MetaReferences = append(metaTo.MetaReferences, CopyBranchMetaReference(mapOrigCopy, _metareference))
	}

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
func UnstageBranch[Type Gongstruct](stage *StageStruct, instance *Type) {

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

	case *Meta:
		stage.UnstageBranchMeta(target)

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
func (stage *StageStruct) UnstageBranchGongBasicField(gongbasicfield *GongBasicField) {

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

func (stage *StageStruct) UnstageBranchGongEnum(gongenum *GongEnum) {

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

func (stage *StageStruct) UnstageBranchGongEnumValue(gongenumvalue *GongEnumValue) {

	// check if instance is already staged
	if !IsStaged(stage, gongenumvalue) {
		return
	}

	gongenumvalue.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGongLink(gonglink *GongLink) {

	// check if instance is already staged
	if !IsStaged(stage, gonglink) {
		return
	}

	gonglink.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchGongNote(gongnote *GongNote) {

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

func (stage *StageStruct) UnstageBranchGongStruct(gongstruct *GongStruct) {

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

func (stage *StageStruct) UnstageBranchGongTimeField(gongtimefield *GongTimeField) {

	// check if instance is already staged
	if !IsStaged(stage, gongtimefield) {
		return
	}

	gongtimefield.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchMeta(meta *Meta) {

	// check if instance is already staged
	if !IsStaged(stage, meta) {
		return
	}

	meta.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _metareference := range meta.MetaReferences {
		UnstageBranch(stage, _metareference)
	}

}

func (stage *StageStruct) UnstageBranchMetaReference(metareference *MetaReference) {

	// check if instance is already staged
	if !IsStaged(stage, metareference) {
		return
	}

	metareference.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchModelPkg(modelpkg *ModelPkg) {

	// check if instance is already staged
	if !IsStaged(stage, modelpkg) {
		return
	}

	modelpkg.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (stage *StageStruct) UnstageBranchPointerToGongStructField(pointertogongstructfield *PointerToGongStructField) {

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

func (stage *StageStruct) UnstageBranchSliceOfPointerToGongStructField(sliceofpointertogongstructfield *SliceOfPointerToGongStructField) {

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
