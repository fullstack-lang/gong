// generated code - do not edit
package metadataextension

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
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_DatatypeCode_PropertyTypes[md_datatypecode_propertytype]
	return ok
}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_ExtendedElementInformation_PropertyTypes[md_extendedelementinformation_propertytype]
	return ok
}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_ExtendedElementInformation_Types[md_extendedelementinformation_type]
	return ok
}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_MetadataExtensionInformation_PropertyTypes[md_metadataextensioninformation_propertytype]
	return ok
}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_MetadataExtensionInformation_Types[md_metadataextensioninformation_type]
	return ok
}

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MD_ObligationCode_PropertyTypes[md_obligationcode_propertytype]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_datatypecode_propertytype) {
		return
	}

	md_datatypecode_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_extendedelementinformation_propertytype) {
		return
	}

	md_extendedelementinformation_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_extendedelementinformation_propertytype.MD_ExtendedElementInformation != nil {
		stage.StageBranch(md_extendedelementinformation_propertytype.MD_ExtendedElementInformation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_extendedelementinformation_type) {
		return
	}

	md_extendedelementinformation_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_extendedelementinformation_type.Obligation != nil {
		stage.StageBranch(md_extendedelementinformation_type.Obligation)
	}
	if md_extendedelementinformation_type.DataType != nil {
		stage.StageBranch(md_extendedelementinformation_type.DataType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_metadataextensioninformation_propertytype) {
		return
	}

	md_metadataextensioninformation_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_metadataextensioninformation_propertytype.MD_MetadataExtensionInformation != nil {
		stage.StageBranch(md_metadataextensioninformation_propertytype.MD_MetadataExtensionInformation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_metadataextensioninformation_type) {
		return
	}

	md_metadataextensioninformation_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_extendedelementinformation_propertytype := range md_metadataextensioninformation_type.ExtendedElementInformation {
		stage.StageBranch(_md_extendedelementinformation_propertytype)
	}

}

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(md_obligationcode_propertytype) {
		return
	}

	md_obligationcode_propertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

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
	case *MD_DatatypeCode_PropertyType:
		toT := GongCopyBranchMD_DatatypeCode_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_ExtendedElementInformation_PropertyType:
		toT := GongCopyBranchMD_ExtendedElementInformation_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_ExtendedElementInformation_Type:
		toT := GongCopyBranchMD_ExtendedElementInformation_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_MetadataExtensionInformation_PropertyType:
		toT := GongCopyBranchMD_MetadataExtensionInformation_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_MetadataExtensionInformation_Type:
		toT := GongCopyBranchMD_MetadataExtensionInformation_Type(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MD_ObligationCode_PropertyType:
		toT := GongCopyBranchMD_ObligationCode_PropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchMD_DatatypeCode_PropertyType(mapOrigCopy map[any]any, md_datatypecode_propertytypeFrom *MD_DatatypeCode_PropertyType) (md_datatypecode_propertytypeTo *MD_DatatypeCode_PropertyType) {
	var alreadyCopied bool
	md_datatypecode_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_datatypecode_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_datatypecode_propertytypeFrom.GongCopyBasicFields(md_datatypecode_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_ExtendedElementInformation_PropertyType(mapOrigCopy map[any]any, md_extendedelementinformation_propertytypeFrom *MD_ExtendedElementInformation_PropertyType) (md_extendedelementinformation_propertytypeTo *MD_ExtendedElementInformation_PropertyType) {
	var alreadyCopied bool
	md_extendedelementinformation_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_extendedelementinformation_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_extendedelementinformation_propertytypeFrom.GongCopyBasicFields(md_extendedelementinformation_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_extendedelementinformation_propertytypeFrom.MD_ExtendedElementInformation != nil {
		md_extendedelementinformation_propertytypeTo.MD_ExtendedElementInformation = GongCopyBranchMD_ExtendedElementInformation_Type(mapOrigCopy, md_extendedelementinformation_propertytypeFrom.MD_ExtendedElementInformation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_ExtendedElementInformation_Type(mapOrigCopy map[any]any, md_extendedelementinformation_typeFrom *MD_ExtendedElementInformation_Type) (md_extendedelementinformation_typeTo *MD_ExtendedElementInformation_Type) {
	var alreadyCopied bool
	md_extendedelementinformation_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_extendedelementinformation_typeFrom)
	if alreadyCopied {
		return
	}
	md_extendedelementinformation_typeFrom.GongCopyBasicFields(md_extendedelementinformation_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_extendedelementinformation_typeFrom.Obligation != nil {
		md_extendedelementinformation_typeTo.Obligation = GongCopyBranchMD_ObligationCode_PropertyType(mapOrigCopy, md_extendedelementinformation_typeFrom.Obligation)
	}
	if md_extendedelementinformation_typeFrom.DataType != nil {
		md_extendedelementinformation_typeTo.DataType = GongCopyBranchMD_DatatypeCode_PropertyType(mapOrigCopy, md_extendedelementinformation_typeFrom.DataType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_MetadataExtensionInformation_PropertyType(mapOrigCopy map[any]any, md_metadataextensioninformation_propertytypeFrom *MD_MetadataExtensionInformation_PropertyType) (md_metadataextensioninformation_propertytypeTo *MD_MetadataExtensionInformation_PropertyType) {
	var alreadyCopied bool
	md_metadataextensioninformation_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_metadataextensioninformation_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_metadataextensioninformation_propertytypeFrom.GongCopyBasicFields(md_metadataextensioninformation_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if md_metadataextensioninformation_propertytypeFrom.MD_MetadataExtensionInformation != nil {
		md_metadataextensioninformation_propertytypeTo.MD_MetadataExtensionInformation = GongCopyBranchMD_MetadataExtensionInformation_Type(mapOrigCopy, md_metadataextensioninformation_propertytypeFrom.MD_MetadataExtensionInformation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMD_MetadataExtensionInformation_Type(mapOrigCopy map[any]any, md_metadataextensioninformation_typeFrom *MD_MetadataExtensionInformation_Type) (md_metadataextensioninformation_typeTo *MD_MetadataExtensionInformation_Type) {
	var alreadyCopied bool
	md_metadataextensioninformation_typeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_metadataextensioninformation_typeFrom)
	if alreadyCopied {
		return
	}
	md_metadataextensioninformation_typeFrom.GongCopyBasicFields(md_metadataextensioninformation_typeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_extendedelementinformation_propertytype := range md_metadataextensioninformation_typeFrom.ExtendedElementInformation {
		md_metadataextensioninformation_typeTo.ExtendedElementInformation = append(md_metadataextensioninformation_typeTo.ExtendedElementInformation, GongCopyBranchMD_ExtendedElementInformation_PropertyType(mapOrigCopy, _md_extendedelementinformation_propertytype))
	}

	return
}

func GongCopyBranchMD_ObligationCode_PropertyType(mapOrigCopy map[any]any, md_obligationcode_propertytypeFrom *MD_ObligationCode_PropertyType) (md_obligationcode_propertytypeTo *MD_ObligationCode_PropertyType) {
	var alreadyCopied bool
	md_obligationcode_propertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, md_obligationcode_propertytypeFrom)
	if alreadyCopied {
		return
	}
	md_obligationcode_propertytypeFrom.GongCopyBasicFields(md_obligationcode_propertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

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
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_datatypecode_propertytype) {
		return
	}

	md_datatypecode_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_extendedelementinformation_propertytype) {
		return
	}

	md_extendedelementinformation_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_extendedelementinformation_propertytype.MD_ExtendedElementInformation != nil {
		stage.UnstageBranch(md_extendedelementinformation_propertytype.MD_ExtendedElementInformation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_extendedelementinformation_type) {
		return
	}

	md_extendedelementinformation_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_extendedelementinformation_type.Obligation != nil {
		stage.UnstageBranch(md_extendedelementinformation_type.Obligation)
	}
	if md_extendedelementinformation_type.DataType != nil {
		stage.UnstageBranch(md_extendedelementinformation_type.DataType)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_metadataextensioninformation_propertytype) {
		return
	}

	md_metadataextensioninformation_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if md_metadataextensioninformation_propertytype.MD_MetadataExtensionInformation != nil {
		stage.UnstageBranch(md_metadataextensioninformation_propertytype.MD_MetadataExtensionInformation)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_metadataextensioninformation_type) {
		return
	}

	md_metadataextensioninformation_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _md_extendedelementinformation_propertytype := range md_metadataextensioninformation_type.ExtendedElementInformation {
		stage.UnstageBranch(_md_extendedelementinformation_propertytype)
	}

}

func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(md_obligationcode_propertytype) {
		return
	}

	md_obligationcode_propertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *MD_DatatypeCode_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_DatatypeCode_PropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MD_ExtendedElementInformation_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_ExtendedElementInformation_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_ExtendedElementInformation, stage.MD_ExtendedElementInformation_Types_reference, instance.MD_ExtendedElementInformation)
	// insertion point for slice of pointers field
}

func (reference *MD_ExtendedElementInformation_Type) GongReconstructPointersFromReferences(stage *Stage, instance *MD_ExtendedElementInformation_Type) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Obligation, stage.MD_ObligationCode_PropertyTypes_reference, instance.Obligation)
	__gong__reconstructPointer(&reference.DataType, stage.MD_DatatypeCode_PropertyTypes_reference, instance.DataType)
	// insertion point for slice of pointers field
}

func (reference *MD_MetadataExtensionInformation_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_MetadataExtensionInformation_PropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MD_MetadataExtensionInformation, stage.MD_MetadataExtensionInformation_Types_reference, instance.MD_MetadataExtensionInformation)
	// insertion point for slice of pointers field
}

func (reference *MD_MetadataExtensionInformation_Type) GongReconstructPointersFromReferences(stage *Stage, instance *MD_MetadataExtensionInformation_Type) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.ExtendedElementInformation, stage.MD_ExtendedElementInformation_PropertyTypes_reference, instance.ExtendedElementInformation)
}

func (reference *MD_ObligationCode_PropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MD_ObligationCode_PropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *MD_DatatypeCode_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MD_ExtendedElementInformation_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_ExtendedElementInformation, stage.MD_ExtendedElementInformation_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_ExtendedElementInformation_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Obligation, stage.MD_ObligationCode_PropertyTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.DataType, stage.MD_DatatypeCode_PropertyTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_MetadataExtensionInformation_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MD_MetadataExtensionInformation, stage.MD_MetadataExtensionInformation_Types_instance)
	// insertion point for slice of pointers fields
}

func (reference *MD_MetadataExtensionInformation_Type) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.ExtendedElementInformation, stage.MD_ExtendedElementInformation_PropertyTypes_instance)
}

func (reference *MD_ObligationCode_PropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_datatypecode_propertytype *MD_DatatypeCode_PropertyType) GongDiff(stage *Stage, md_datatypecode_propertytypeOther *MD_DatatypeCode_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_datatypecode_propertytype.Name != md_datatypecode_propertytypeOther.Name {
		diffs = append(diffs, md_datatypecode_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_datatypecode_propertytype.NilReason != md_datatypecode_propertytypeOther.NilReason {
		diffs = append(diffs, md_datatypecode_propertytype.GongMarshallField(stage, "NilReason"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_extendedelementinformation_propertytype *MD_ExtendedElementInformation_PropertyType) GongDiff(stage *Stage, md_extendedelementinformation_propertytypeOther *MD_ExtendedElementInformation_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_extendedelementinformation_propertytype.Name != md_extendedelementinformation_propertytypeOther.Name {
		diffs = append(diffs, md_extendedelementinformation_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_extendedelementinformation_propertytype.NilReason != md_extendedelementinformation_propertytypeOther.NilReason {
		diffs = append(diffs, md_extendedelementinformation_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_extendedelementinformation_propertytype.Uuidref != md_extendedelementinformation_propertytypeOther.Uuidref {
		diffs = append(diffs, md_extendedelementinformation_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_extendedelementinformation_propertytype.Type_XSD != md_extendedelementinformation_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_extendedelementinformation_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_extendedelementinformation_propertytype.Href != md_extendedelementinformation_propertytypeOther.Href {
		diffs = append(diffs, md_extendedelementinformation_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_extendedelementinformation_propertytype.Role != md_extendedelementinformation_propertytypeOther.Role {
		diffs = append(diffs, md_extendedelementinformation_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_extendedelementinformation_propertytype.Arcrole != md_extendedelementinformation_propertytypeOther.Arcrole {
		diffs = append(diffs, md_extendedelementinformation_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_extendedelementinformation_propertytype.Title != md_extendedelementinformation_propertytypeOther.Title {
		diffs = append(diffs, md_extendedelementinformation_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_extendedelementinformation_propertytype.Show != md_extendedelementinformation_propertytypeOther.Show {
		diffs = append(diffs, md_extendedelementinformation_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_extendedelementinformation_propertytype.Actuate != md_extendedelementinformation_propertytypeOther.Actuate {
		diffs = append(diffs, md_extendedelementinformation_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_extendedelementinformation_propertytype.MD_ExtendedElementInformation != md_extendedelementinformation_propertytypeOther.MD_ExtendedElementInformation {
		diffs = append(diffs, md_extendedelementinformation_propertytype.GongMarshallField(stage, "MD_ExtendedElementInformation"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_extendedelementinformation_type *MD_ExtendedElementInformation_Type) GongDiff(stage *Stage, md_extendedelementinformation_typeOther *MD_ExtendedElementInformation_Type) (diffs []string) {
	// insertion point for field diffs
	if md_extendedelementinformation_type.Name != md_extendedelementinformation_typeOther.Name {
		diffs = append(diffs, md_extendedelementinformation_type.GongMarshallField(stage, "Name"))
	}
	if md_extendedelementinformation_type.Id != md_extendedelementinformation_typeOther.Id {
		diffs = append(diffs, md_extendedelementinformation_type.GongMarshallField(stage, "Id"))
	}
	if md_extendedelementinformation_type.Uuid != md_extendedelementinformation_typeOther.Uuid {
		diffs = append(diffs, md_extendedelementinformation_type.GongMarshallField(stage, "Uuid"))
	}
	if md_extendedelementinformation_type.Obligation != md_extendedelementinformation_typeOther.Obligation {
		diffs = append(diffs, md_extendedelementinformation_type.GongMarshallField(stage, "Obligation"))
	}
	if md_extendedelementinformation_type.DataType != md_extendedelementinformation_typeOther.DataType {
		diffs = append(diffs, md_extendedelementinformation_type.GongMarshallField(stage, "DataType"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_metadataextensioninformation_propertytype *MD_MetadataExtensionInformation_PropertyType) GongDiff(stage *Stage, md_metadataextensioninformation_propertytypeOther *MD_MetadataExtensionInformation_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_metadataextensioninformation_propertytype.Name != md_metadataextensioninformation_propertytypeOther.Name {
		diffs = append(diffs, md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_metadataextensioninformation_propertytype.NilReason != md_metadataextensioninformation_propertytypeOther.NilReason {
		diffs = append(diffs, md_metadataextensioninformation_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_metadataextensioninformation_propertytype.Uuidref != md_metadataextensioninformation_propertytypeOther.Uuidref {
		diffs = append(diffs, md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Uuidref"))
	}
	if md_metadataextensioninformation_propertytype.Type_XSD != md_metadataextensioninformation_propertytypeOther.Type_XSD {
		diffs = append(diffs, md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if md_metadataextensioninformation_propertytype.Href != md_metadataextensioninformation_propertytypeOther.Href {
		diffs = append(diffs, md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Href"))
	}
	if md_metadataextensioninformation_propertytype.Role != md_metadataextensioninformation_propertytypeOther.Role {
		diffs = append(diffs, md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Role"))
	}
	if md_metadataextensioninformation_propertytype.Arcrole != md_metadataextensioninformation_propertytypeOther.Arcrole {
		diffs = append(diffs, md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Arcrole"))
	}
	if md_metadataextensioninformation_propertytype.Title != md_metadataextensioninformation_propertytypeOther.Title {
		diffs = append(diffs, md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Title"))
	}
	if md_metadataextensioninformation_propertytype.Show != md_metadataextensioninformation_propertytypeOther.Show {
		diffs = append(diffs, md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Show"))
	}
	if md_metadataextensioninformation_propertytype.Actuate != md_metadataextensioninformation_propertytypeOther.Actuate {
		diffs = append(diffs, md_metadataextensioninformation_propertytype.GongMarshallField(stage, "Actuate"))
	}
	if md_metadataextensioninformation_propertytype.MD_MetadataExtensionInformation != md_metadataextensioninformation_propertytypeOther.MD_MetadataExtensionInformation {
		diffs = append(diffs, md_metadataextensioninformation_propertytype.GongMarshallField(stage, "MD_MetadataExtensionInformation"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_metadataextensioninformation_type *MD_MetadataExtensionInformation_Type) GongDiff(stage *Stage, md_metadataextensioninformation_typeOther *MD_MetadataExtensionInformation_Type) (diffs []string) {
	// insertion point for field diffs
	if md_metadataextensioninformation_type.Name != md_metadataextensioninformation_typeOther.Name {
		diffs = append(diffs, md_metadataextensioninformation_type.GongMarshallField(stage, "Name"))
	}
	if md_metadataextensioninformation_type.Id != md_metadataextensioninformation_typeOther.Id {
		diffs = append(diffs, md_metadataextensioninformation_type.GongMarshallField(stage, "Id"))
	}
	if md_metadataextensioninformation_type.Uuid != md_metadataextensioninformation_typeOther.Uuid {
		diffs = append(diffs, md_metadataextensioninformation_type.GongMarshallField(stage, "Uuid"))
	}
	if ops := __gong__diffSliceOfPointers(stage, md_metadataextensioninformation_type, "ExtendedElementInformation", md_metadataextensioninformation_typeOther.ExtendedElementInformation, md_metadataextensioninformation_type.ExtendedElementInformation); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (md_obligationcode_propertytype *MD_ObligationCode_PropertyType) GongDiff(stage *Stage, md_obligationcode_propertytypeOther *MD_ObligationCode_PropertyType) (diffs []string) {
	// insertion point for field diffs
	if md_obligationcode_propertytype.Name != md_obligationcode_propertytypeOther.Name {
		diffs = append(diffs, md_obligationcode_propertytype.GongMarshallField(stage, "Name"))
	}
	if md_obligationcode_propertytype.NilReason != md_obligationcode_propertytypeOther.NilReason {
		diffs = append(diffs, md_obligationcode_propertytype.GongMarshallField(stage, "NilReason"))
	}
	if md_obligationcode_propertytype.MD_ObligationCode != md_obligationcode_propertytypeOther.MD_ObligationCode {
		diffs = append(diffs, md_obligationcode_propertytype.GongMarshallField(stage, "MD_ObligationCode"))
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
