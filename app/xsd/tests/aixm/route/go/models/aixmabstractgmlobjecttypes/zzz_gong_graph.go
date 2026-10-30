// generated code - do not edit
package aixmabstractgmlobjecttypes

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
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractAIXMFeatureBaseTypes[abstractaixmfeaturebasetype]
	return ok
}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractAIXMFeatureTypes[abstractaixmfeaturetype]
	return ok
}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractAIXMMessageBaseTypes[abstractaixmmessagebasetype]
	return ok
}

func (abstractaixmmessagetype *AbstractAIXMMessageType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractAIXMMessageTypes[abstractaixmmessagetype]
	return ok
}

func (abstractaixmobjecttype *AbstractAIXMObjectType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractAIXMObjectTypes[abstractaixmobjecttype]
	return ok
}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractAIXMPropertyTypes[abstractaixmpropertytype]
	return ok
}

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractAIXMTimeSliceBaseTypes[abstractaixmtimeslicebasetype]
	return ok
}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractAIXMTimeSliceTypes[abstractaixmtimeslicetype]
	return ok
}

func (abstractextensiontype *AbstractExtensionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractExtensionTypes[abstractextensiontype]
	return ok
}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FeatureMetadataPropertyTypes[featuremetadatapropertytype]
	return ok
}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.FeatureTimeSliceMetadataPropertyTypes[featuretimeslicemetadatapropertytype]
	return ok
}

func (messagemetadatapropertytype *MessageMetadataPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.MessageMetadataPropertyTypes[messagemetadatapropertytype]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractaixmfeaturebasetype) {
		return
	}

	abstractaixmfeaturebasetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractaixmfeaturetype) {
		return
	}

	abstractaixmfeaturetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractaixmfeaturetype.FeatureMetadata != nil {
		stage.StageBranch(abstractaixmfeaturetype.FeatureMetadata)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractaixmmessagebasetype) {
		return
	}

	abstractaixmmessagebasetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractaixmmessagebasetype.MessageMetadata != nil {
		stage.StageBranch(abstractaixmmessagebasetype.MessageMetadata)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractaixmmessagetype *AbstractAIXMMessageType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractaixmmessagetype) {
		return
	}

	abstractaixmmessagetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractaixmmessagetype.MessageMetadata != nil {
		stage.StageBranch(abstractaixmmessagetype.MessageMetadata)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractaixmobjecttype *AbstractAIXMObjectType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractaixmobjecttype) {
		return
	}

	abstractaixmobjecttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractaixmpropertytype) {
		return
	}

	abstractaixmpropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractaixmtimeslicebasetype) {
		return
	}

	abstractaixmtimeslicebasetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractaixmtimeslicetype) {
		return
	}

	abstractaixmtimeslicetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractaixmtimeslicetype.TimeSliceMetadata != nil {
		stage.StageBranch(abstractaixmtimeslicetype.TimeSliceMetadata)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractextensiontype *AbstractExtensionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractextensiontype) {
		return
	}

	abstractextensiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(featuremetadatapropertytype) {
		return
	}

	featuremetadatapropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(featuretimeslicemetadatapropertytype) {
		return
	}

	featuretimeslicemetadatapropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (messagemetadatapropertytype *MessageMetadataPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(messagemetadatapropertytype) {
		return
	}

	messagemetadatapropertytype.Stage(stage)

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
	case *AbstractAIXMFeatureBaseType:
		toT := GongCopyBranchAbstractAIXMFeatureBaseType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractAIXMFeatureType:
		toT := GongCopyBranchAbstractAIXMFeatureType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractAIXMMessageBaseType:
		toT := GongCopyBranchAbstractAIXMMessageBaseType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractAIXMMessageType:
		toT := GongCopyBranchAbstractAIXMMessageType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractAIXMObjectType:
		toT := GongCopyBranchAbstractAIXMObjectType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractAIXMPropertyType:
		toT := GongCopyBranchAbstractAIXMPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractAIXMTimeSliceBaseType:
		toT := GongCopyBranchAbstractAIXMTimeSliceBaseType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractAIXMTimeSliceType:
		toT := GongCopyBranchAbstractAIXMTimeSliceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractExtensionType:
		toT := GongCopyBranchAbstractExtensionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FeatureMetadataPropertyType:
		toT := GongCopyBranchFeatureMetadataPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *FeatureTimeSliceMetadataPropertyType:
		toT := GongCopyBranchFeatureTimeSliceMetadataPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *MessageMetadataPropertyType:
		toT := GongCopyBranchMessageMetadataPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAbstractAIXMFeatureBaseType(mapOrigCopy map[any]any, abstractaixmfeaturebasetypeFrom *AbstractAIXMFeatureBaseType) (abstractaixmfeaturebasetypeTo *AbstractAIXMFeatureBaseType) {
	var alreadyCopied bool
	abstractaixmfeaturebasetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractaixmfeaturebasetypeFrom)
	if alreadyCopied {
		return
	}
	abstractaixmfeaturebasetypeFrom.GongCopyBasicFields(abstractaixmfeaturebasetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractAIXMFeatureType(mapOrigCopy map[any]any, abstractaixmfeaturetypeFrom *AbstractAIXMFeatureType) (abstractaixmfeaturetypeTo *AbstractAIXMFeatureType) {
	var alreadyCopied bool
	abstractaixmfeaturetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractaixmfeaturetypeFrom)
	if alreadyCopied {
		return
	}
	abstractaixmfeaturetypeFrom.GongCopyBasicFields(abstractaixmfeaturetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstractaixmfeaturetypeFrom.FeatureMetadata != nil {
		abstractaixmfeaturetypeTo.FeatureMetadata = GongCopyBranchFeatureMetadataPropertyType(mapOrigCopy, abstractaixmfeaturetypeFrom.FeatureMetadata)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractAIXMMessageBaseType(mapOrigCopy map[any]any, abstractaixmmessagebasetypeFrom *AbstractAIXMMessageBaseType) (abstractaixmmessagebasetypeTo *AbstractAIXMMessageBaseType) {
	var alreadyCopied bool
	abstractaixmmessagebasetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractaixmmessagebasetypeFrom)
	if alreadyCopied {
		return
	}
	abstractaixmmessagebasetypeFrom.GongCopyBasicFields(abstractaixmmessagebasetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstractaixmmessagebasetypeFrom.MessageMetadata != nil {
		abstractaixmmessagebasetypeTo.MessageMetadata = GongCopyBranchMessageMetadataPropertyType(mapOrigCopy, abstractaixmmessagebasetypeFrom.MessageMetadata)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractAIXMMessageType(mapOrigCopy map[any]any, abstractaixmmessagetypeFrom *AbstractAIXMMessageType) (abstractaixmmessagetypeTo *AbstractAIXMMessageType) {
	var alreadyCopied bool
	abstractaixmmessagetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractaixmmessagetypeFrom)
	if alreadyCopied {
		return
	}
	abstractaixmmessagetypeFrom.GongCopyBasicFields(abstractaixmmessagetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstractaixmmessagetypeFrom.MessageMetadata != nil {
		abstractaixmmessagetypeTo.MessageMetadata = GongCopyBranchMessageMetadataPropertyType(mapOrigCopy, abstractaixmmessagetypeFrom.MessageMetadata)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractAIXMObjectType(mapOrigCopy map[any]any, abstractaixmobjecttypeFrom *AbstractAIXMObjectType) (abstractaixmobjecttypeTo *AbstractAIXMObjectType) {
	var alreadyCopied bool
	abstractaixmobjecttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractaixmobjecttypeFrom)
	if alreadyCopied {
		return
	}
	abstractaixmobjecttypeFrom.GongCopyBasicFields(abstractaixmobjecttypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractAIXMPropertyType(mapOrigCopy map[any]any, abstractaixmpropertytypeFrom *AbstractAIXMPropertyType) (abstractaixmpropertytypeTo *AbstractAIXMPropertyType) {
	var alreadyCopied bool
	abstractaixmpropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractaixmpropertytypeFrom)
	if alreadyCopied {
		return
	}
	abstractaixmpropertytypeFrom.GongCopyBasicFields(abstractaixmpropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractAIXMTimeSliceBaseType(mapOrigCopy map[any]any, abstractaixmtimeslicebasetypeFrom *AbstractAIXMTimeSliceBaseType) (abstractaixmtimeslicebasetypeTo *AbstractAIXMTimeSliceBaseType) {
	var alreadyCopied bool
	abstractaixmtimeslicebasetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractaixmtimeslicebasetypeFrom)
	if alreadyCopied {
		return
	}
	abstractaixmtimeslicebasetypeFrom.GongCopyBasicFields(abstractaixmtimeslicebasetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractAIXMTimeSliceType(mapOrigCopy map[any]any, abstractaixmtimeslicetypeFrom *AbstractAIXMTimeSliceType) (abstractaixmtimeslicetypeTo *AbstractAIXMTimeSliceType) {
	var alreadyCopied bool
	abstractaixmtimeslicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractaixmtimeslicetypeFrom)
	if alreadyCopied {
		return
	}
	abstractaixmtimeslicetypeFrom.GongCopyBasicFields(abstractaixmtimeslicetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstractaixmtimeslicetypeFrom.TimeSliceMetadata != nil {
		abstractaixmtimeslicetypeTo.TimeSliceMetadata = GongCopyBranchFeatureTimeSliceMetadataPropertyType(mapOrigCopy, abstractaixmtimeslicetypeFrom.TimeSliceMetadata)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractExtensionType(mapOrigCopy map[any]any, abstractextensiontypeFrom *AbstractExtensionType) (abstractextensiontypeTo *AbstractExtensionType) {
	var alreadyCopied bool
	abstractextensiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractextensiontypeFrom)
	if alreadyCopied {
		return
	}
	abstractextensiontypeFrom.GongCopyBasicFields(abstractextensiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFeatureMetadataPropertyType(mapOrigCopy map[any]any, featuremetadatapropertytypeFrom *FeatureMetadataPropertyType) (featuremetadatapropertytypeTo *FeatureMetadataPropertyType) {
	var alreadyCopied bool
	featuremetadatapropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, featuremetadatapropertytypeFrom)
	if alreadyCopied {
		return
	}
	featuremetadatapropertytypeFrom.GongCopyBasicFields(featuremetadatapropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchFeatureTimeSliceMetadataPropertyType(mapOrigCopy map[any]any, featuretimeslicemetadatapropertytypeFrom *FeatureTimeSliceMetadataPropertyType) (featuretimeslicemetadatapropertytypeTo *FeatureTimeSliceMetadataPropertyType) {
	var alreadyCopied bool
	featuretimeslicemetadatapropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, featuretimeslicemetadatapropertytypeFrom)
	if alreadyCopied {
		return
	}
	featuretimeslicemetadatapropertytypeFrom.GongCopyBasicFields(featuretimeslicemetadatapropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchMessageMetadataPropertyType(mapOrigCopy map[any]any, messagemetadatapropertytypeFrom *MessageMetadataPropertyType) (messagemetadatapropertytypeTo *MessageMetadataPropertyType) {
	var alreadyCopied bool
	messagemetadatapropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, messagemetadatapropertytypeFrom)
	if alreadyCopied {
		return
	}
	messagemetadatapropertytypeFrom.GongCopyBasicFields(messagemetadatapropertytypeTo)

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
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractaixmfeaturebasetype) {
		return
	}

	abstractaixmfeaturebasetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractaixmfeaturetype) {
		return
	}

	abstractaixmfeaturetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractaixmfeaturetype.FeatureMetadata != nil {
		stage.UnstageBranch(abstractaixmfeaturetype.FeatureMetadata)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractaixmmessagebasetype) {
		return
	}

	abstractaixmmessagebasetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractaixmmessagebasetype.MessageMetadata != nil {
		stage.UnstageBranch(abstractaixmmessagebasetype.MessageMetadata)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractaixmmessagetype *AbstractAIXMMessageType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractaixmmessagetype) {
		return
	}

	abstractaixmmessagetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractaixmmessagetype.MessageMetadata != nil {
		stage.UnstageBranch(abstractaixmmessagetype.MessageMetadata)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractaixmobjecttype *AbstractAIXMObjectType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractaixmobjecttype) {
		return
	}

	abstractaixmobjecttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractaixmpropertytype) {
		return
	}

	abstractaixmpropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractaixmtimeslicebasetype) {
		return
	}

	abstractaixmtimeslicebasetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractaixmtimeslicetype) {
		return
	}

	abstractaixmtimeslicetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractaixmtimeslicetype.TimeSliceMetadata != nil {
		stage.UnstageBranch(abstractaixmtimeslicetype.TimeSliceMetadata)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractextensiontype *AbstractExtensionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractextensiontype) {
		return
	}

	abstractextensiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(featuremetadatapropertytype) {
		return
	}

	featuremetadatapropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(featuretimeslicemetadatapropertytype) {
		return
	}

	featuretimeslicemetadatapropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (messagemetadatapropertytype *MessageMetadataPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(messagemetadatapropertytype) {
		return
	}

	messagemetadatapropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *AbstractAIXMFeatureBaseType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractAIXMFeatureBaseType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AbstractAIXMFeatureType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractAIXMFeatureType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.FeatureMetadata, stage.FeatureMetadataPropertyTypes_reference, instance.FeatureMetadata)
	// insertion point for slice of pointers field
}

func (reference *AbstractAIXMMessageBaseType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractAIXMMessageBaseType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MessageMetadata, stage.MessageMetadataPropertyTypes_reference, instance.MessageMetadata)
	// insertion point for slice of pointers field
}

func (reference *AbstractAIXMMessageType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractAIXMMessageType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.MessageMetadata, stage.MessageMetadataPropertyTypes_reference, instance.MessageMetadata)
	// insertion point for slice of pointers field
}

func (reference *AbstractAIXMObjectType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractAIXMObjectType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AbstractAIXMPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractAIXMPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AbstractAIXMTimeSliceBaseType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractAIXMTimeSliceBaseType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AbstractAIXMTimeSliceType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractAIXMTimeSliceType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.TimeSliceMetadata, stage.FeatureTimeSliceMetadataPropertyTypes_reference, instance.TimeSliceMetadata)
	// insertion point for slice of pointers field
}

func (reference *AbstractExtensionType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractExtensionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FeatureMetadataPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *FeatureMetadataPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *FeatureTimeSliceMetadataPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *FeatureTimeSliceMetadataPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *MessageMetadataPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *MessageMetadataPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AbstractAIXMFeatureBaseType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AbstractAIXMFeatureType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.FeatureMetadata, stage.FeatureMetadataPropertyTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *AbstractAIXMMessageBaseType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MessageMetadata, stage.MessageMetadataPropertyTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *AbstractAIXMMessageType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.MessageMetadata, stage.MessageMetadataPropertyTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *AbstractAIXMObjectType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AbstractAIXMPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AbstractAIXMTimeSliceBaseType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AbstractAIXMTimeSliceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.TimeSliceMetadata, stage.FeatureTimeSliceMetadataPropertyTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *AbstractExtensionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FeatureMetadataPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *FeatureTimeSliceMetadataPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *MessageMetadataPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractaixmfeaturebasetype *AbstractAIXMFeatureBaseType) GongDiff(stage *Stage, abstractaixmfeaturebasetypeOther *AbstractAIXMFeatureBaseType) (diffs []string) {
	// insertion point for field diffs
	if abstractaixmfeaturebasetype.Name != abstractaixmfeaturebasetypeOther.Name {
		diffs = append(diffs, abstractaixmfeaturebasetype.GongMarshallField(stage, "Name"))
	}
	if abstractaixmfeaturebasetype.Id != abstractaixmfeaturebasetypeOther.Id {
		diffs = append(diffs, abstractaixmfeaturebasetype.GongMarshallField(stage, "Id"))
	}
	if abstractaixmfeaturebasetype.Description != abstractaixmfeaturebasetypeOther.Description {
		diffs = append(diffs, abstractaixmfeaturebasetype.GongMarshallField(stage, "Description"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractaixmfeaturetype *AbstractAIXMFeatureType) GongDiff(stage *Stage, abstractaixmfeaturetypeOther *AbstractAIXMFeatureType) (diffs []string) {
	// insertion point for field diffs
	if abstractaixmfeaturetype.Name != abstractaixmfeaturetypeOther.Name {
		diffs = append(diffs, abstractaixmfeaturetype.GongMarshallField(stage, "Name"))
	}
	if abstractaixmfeaturetype.Id != abstractaixmfeaturetypeOther.Id {
		diffs = append(diffs, abstractaixmfeaturetype.GongMarshallField(stage, "Id"))
	}
	if abstractaixmfeaturetype.Description != abstractaixmfeaturetypeOther.Description {
		diffs = append(diffs, abstractaixmfeaturetype.GongMarshallField(stage, "Description"))
	}
	if abstractaixmfeaturetype.FeatureMetadata != abstractaixmfeaturetypeOther.FeatureMetadata {
		diffs = append(diffs, abstractaixmfeaturetype.GongMarshallField(stage, "FeatureMetadata"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractaixmmessagebasetype *AbstractAIXMMessageBaseType) GongDiff(stage *Stage, abstractaixmmessagebasetypeOther *AbstractAIXMMessageBaseType) (diffs []string) {
	// insertion point for field diffs
	if abstractaixmmessagebasetype.Name != abstractaixmmessagebasetypeOther.Name {
		diffs = append(diffs, abstractaixmmessagebasetype.GongMarshallField(stage, "Name"))
	}
	if abstractaixmmessagebasetype.Id != abstractaixmmessagebasetypeOther.Id {
		diffs = append(diffs, abstractaixmmessagebasetype.GongMarshallField(stage, "Id"))
	}
	if abstractaixmmessagebasetype.Description != abstractaixmmessagebasetypeOther.Description {
		diffs = append(diffs, abstractaixmmessagebasetype.GongMarshallField(stage, "Description"))
	}
	if abstractaixmmessagebasetype.SequenceNumber != abstractaixmmessagebasetypeOther.SequenceNumber {
		diffs = append(diffs, abstractaixmmessagebasetype.GongMarshallField(stage, "SequenceNumber"))
	}
	if abstractaixmmessagebasetype.MessageMetadata != abstractaixmmessagebasetypeOther.MessageMetadata {
		diffs = append(diffs, abstractaixmmessagebasetype.GongMarshallField(stage, "MessageMetadata"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractaixmmessagetype *AbstractAIXMMessageType) GongDiff(stage *Stage, abstractaixmmessagetypeOther *AbstractAIXMMessageType) (diffs []string) {
	// insertion point for field diffs
	if abstractaixmmessagetype.Name != abstractaixmmessagetypeOther.Name {
		diffs = append(diffs, abstractaixmmessagetype.GongMarshallField(stage, "Name"))
	}
	if abstractaixmmessagetype.Id != abstractaixmmessagetypeOther.Id {
		diffs = append(diffs, abstractaixmmessagetype.GongMarshallField(stage, "Id"))
	}
	if abstractaixmmessagetype.Description != abstractaixmmessagetypeOther.Description {
		diffs = append(diffs, abstractaixmmessagetype.GongMarshallField(stage, "Description"))
	}
	if abstractaixmmessagetype.SequenceNumber != abstractaixmmessagetypeOther.SequenceNumber {
		diffs = append(diffs, abstractaixmmessagetype.GongMarshallField(stage, "SequenceNumber"))
	}
	if abstractaixmmessagetype.MessageMetadata != abstractaixmmessagetypeOther.MessageMetadata {
		diffs = append(diffs, abstractaixmmessagetype.GongMarshallField(stage, "MessageMetadata"))
	}
	if abstractaixmmessagetype.AggregationType != abstractaixmmessagetypeOther.AggregationType {
		diffs = append(diffs, abstractaixmmessagetype.GongMarshallField(stage, "AggregationType"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractaixmobjecttype *AbstractAIXMObjectType) GongDiff(stage *Stage, abstractaixmobjecttypeOther *AbstractAIXMObjectType) (diffs []string) {
	// insertion point for field diffs
	if abstractaixmobjecttype.Name != abstractaixmobjecttypeOther.Name {
		diffs = append(diffs, abstractaixmobjecttype.GongMarshallField(stage, "Name"))
	}
	if abstractaixmobjecttype.Id != abstractaixmobjecttypeOther.Id {
		diffs = append(diffs, abstractaixmobjecttype.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractaixmpropertytype *AbstractAIXMPropertyType) GongDiff(stage *Stage, abstractaixmpropertytypeOther *AbstractAIXMPropertyType) (diffs []string) {
	// insertion point for field diffs
	if abstractaixmpropertytype.Name != abstractaixmpropertytypeOther.Name {
		diffs = append(diffs, abstractaixmpropertytype.GongMarshallField(stage, "Name"))
	}
	if abstractaixmpropertytype.NilReason != abstractaixmpropertytypeOther.NilReason {
		diffs = append(diffs, abstractaixmpropertytype.GongMarshallField(stage, "NilReason"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractaixmtimeslicebasetype *AbstractAIXMTimeSliceBaseType) GongDiff(stage *Stage, abstractaixmtimeslicebasetypeOther *AbstractAIXMTimeSliceBaseType) (diffs []string) {
	// insertion point for field diffs
	if abstractaixmtimeslicebasetype.Name != abstractaixmtimeslicebasetypeOther.Name {
		diffs = append(diffs, abstractaixmtimeslicebasetype.GongMarshallField(stage, "Name"))
	}
	if abstractaixmtimeslicebasetype.Id != abstractaixmtimeslicebasetypeOther.Id {
		diffs = append(diffs, abstractaixmtimeslicebasetype.GongMarshallField(stage, "Id"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractaixmtimeslicetype *AbstractAIXMTimeSliceType) GongDiff(stage *Stage, abstractaixmtimeslicetypeOther *AbstractAIXMTimeSliceType) (diffs []string) {
	// insertion point for field diffs
	if abstractaixmtimeslicetype.Name != abstractaixmtimeslicetypeOther.Name {
		diffs = append(diffs, abstractaixmtimeslicetype.GongMarshallField(stage, "Name"))
	}
	if abstractaixmtimeslicetype.Id != abstractaixmtimeslicetypeOther.Id {
		diffs = append(diffs, abstractaixmtimeslicetype.GongMarshallField(stage, "Id"))
	}
	if abstractaixmtimeslicetype.Description != abstractaixmtimeslicetypeOther.Description {
		diffs = append(diffs, abstractaixmtimeslicetype.GongMarshallField(stage, "Description"))
	}
	if abstractaixmtimeslicetype.Interpretation != abstractaixmtimeslicetypeOther.Interpretation {
		diffs = append(diffs, abstractaixmtimeslicetype.GongMarshallField(stage, "Interpretation"))
	}
	if abstractaixmtimeslicetype.SequenceNumber != abstractaixmtimeslicetypeOther.SequenceNumber {
		diffs = append(diffs, abstractaixmtimeslicetype.GongMarshallField(stage, "SequenceNumber"))
	}
	if abstractaixmtimeslicetype.CorrectionNumber != abstractaixmtimeslicetypeOther.CorrectionNumber {
		diffs = append(diffs, abstractaixmtimeslicetype.GongMarshallField(stage, "CorrectionNumber"))
	}
	if abstractaixmtimeslicetype.TimeSliceMetadata != abstractaixmtimeslicetypeOther.TimeSliceMetadata {
		diffs = append(diffs, abstractaixmtimeslicetype.GongMarshallField(stage, "TimeSliceMetadata"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractextensiontype *AbstractExtensionType) GongDiff(stage *Stage, abstractextensiontypeOther *AbstractExtensionType) (diffs []string) {
	// insertion point for field diffs
	if abstractextensiontype.Name != abstractextensiontypeOther.Name {
		diffs = append(diffs, abstractextensiontype.GongMarshallField(stage, "Name"))
	}
	if abstractextensiontype.Id != abstractextensiontypeOther.Id {
		diffs = append(diffs, abstractextensiontype.GongMarshallField(stage, "Id"))
	}
	if abstractextensiontype.Description != abstractextensiontypeOther.Description {
		diffs = append(diffs, abstractextensiontype.GongMarshallField(stage, "Description"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (featuremetadatapropertytype *FeatureMetadataPropertyType) GongDiff(stage *Stage, featuremetadatapropertytypeOther *FeatureMetadataPropertyType) (diffs []string) {
	// insertion point for field diffs
	if featuremetadatapropertytype.Name != featuremetadatapropertytypeOther.Name {
		diffs = append(diffs, featuremetadatapropertytype.GongMarshallField(stage, "Name"))
	}
	if featuremetadatapropertytype.Owns != featuremetadatapropertytypeOther.Owns {
		diffs = append(diffs, featuremetadatapropertytype.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (featuretimeslicemetadatapropertytype *FeatureTimeSliceMetadataPropertyType) GongDiff(stage *Stage, featuretimeslicemetadatapropertytypeOther *FeatureTimeSliceMetadataPropertyType) (diffs []string) {
	// insertion point for field diffs
	if featuretimeslicemetadatapropertytype.Name != featuretimeslicemetadatapropertytypeOther.Name {
		diffs = append(diffs, featuretimeslicemetadatapropertytype.GongMarshallField(stage, "Name"))
	}
	if featuretimeslicemetadatapropertytype.Owns != featuretimeslicemetadatapropertytypeOther.Owns {
		diffs = append(diffs, featuretimeslicemetadatapropertytype.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (messagemetadatapropertytype *MessageMetadataPropertyType) GongDiff(stage *Stage, messagemetadatapropertytypeOther *MessageMetadataPropertyType) (diffs []string) {
	// insertion point for field diffs
	if messagemetadatapropertytype.Name != messagemetadatapropertytypeOther.Name {
		diffs = append(diffs, messagemetadatapropertytype.GongMarshallField(stage, "Name"))
	}
	if messagemetadatapropertytype.Owns != messagemetadatapropertytypeOther.Owns {
		diffs = append(diffs, messagemetadatapropertytype.GongMarshallField(stage, "Owns"))
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
