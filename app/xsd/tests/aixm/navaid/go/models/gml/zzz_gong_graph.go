// generated code - do not edit
package gml

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
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractFeatureMemberTypes[abstractfeaturemembertype]
	return ok
}

func (abstractfeaturetype *AbstractFeatureType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractFeatureTypes[abstractfeaturetype]
	return ok
}

func (abstractgmltype *AbstractGMLType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractGMLTypes[abstractgmltype]
	return ok
}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractGeometricPrimitiveTypes[abstractgeometricprimitivetype]
	return ok
}

func (abstractgeometrytype *AbstractGeometryType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractGeometryTypes[abstractgeometrytype]
	return ok
}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractMetadataPropertyTypes[abstractmetadatapropertytype]
	return ok
}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractTimeGeometricPrimitiveTypes[abstracttimegeometricprimitivetype]
	return ok
}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractTimeObjectTypes[abstracttimeobjecttype]
	return ok
}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractTimePrimitiveTypes[abstracttimeprimitivetype]
	return ok
}

func (abstracttimeslicetype *AbstractTimeSliceType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.AbstractTimeSliceTypes[abstracttimeslicetype]
	return ok
}

func (boundingshapetype *BoundingShapeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.BoundingShapeTypes[boundingshapetype]
	return ok
}

func (codetype *CodeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeTypes[codetype]
	return ok
}

func (codewithauthoritytype *CodeWithAuthorityType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.CodeWithAuthorityTypes[codewithauthoritytype]
	return ok
}

func (directpositiontype *DirectPositionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DirectPositionTypes[directpositiontype]
	return ok
}

func (dynamicfeaturetype *DynamicFeatureType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.DynamicFeatureTypes[dynamicfeaturetype]
	return ok
}

func (envelopetype *EnvelopeType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.EnvelopeTypes[envelopetype]
	return ok
}

func (pointtype *PointType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.PointTypes[pointtype]
	return ok
}

func (timeinstanttype *TimeInstantType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TimeInstantTypes[timeinstanttype]
	return ok
}

func (timeperiodtype *TimePeriodType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TimePeriodTypes[timeperiodtype]
	return ok
}

func (timepositiontype *TimePositionType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TimePositionTypes[timepositiontype]
	return ok
}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongIsStaged(stage *Stage) bool {
	_, ok := stage.TimePrimitivePropertyTypes[timeprimitivepropertytype]
	return ok
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// insertion point for stage branch per struct
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractfeaturemembertype) {
		return
	}

	abstractfeaturemembertype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractfeaturetype *AbstractFeatureType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractfeaturetype) {
		return
	}

	abstractfeaturetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractfeaturetype.Identifier != nil {
		stage.StageBranch(abstractfeaturetype.Identifier)
	}
	if abstractfeaturetype.BoundedBy != nil {
		stage.StageBranch(abstractfeaturetype.BoundedBy)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractfeaturetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (abstractgmltype *AbstractGMLType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractgmltype) {
		return
	}

	abstractgmltype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractgmltype.Identifier != nil {
		stage.StageBranch(abstractgmltype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgmltype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractgeometricprimitivetype) {
		return
	}

	abstractgeometricprimitivetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractgeometricprimitivetype.Identifier != nil {
		stage.StageBranch(abstractgeometricprimitivetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgeometricprimitivetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (abstractgeometrytype *AbstractGeometryType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractgeometrytype) {
		return
	}

	abstractgeometrytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractgeometrytype.Identifier != nil {
		stage.StageBranch(abstractgeometrytype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgeometrytype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstractmetadatapropertytype) {
		return
	}

	abstractmetadatapropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstracttimegeometricprimitivetype) {
		return
	}

	abstracttimegeometricprimitivetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimegeometricprimitivetype.Identifier != nil {
		stage.StageBranch(abstracttimegeometricprimitivetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimegeometricprimitivetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstracttimeobjecttype) {
		return
	}

	abstracttimeobjecttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeobjecttype.Identifier != nil {
		stage.StageBranch(abstracttimeobjecttype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeobjecttype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstracttimeprimitivetype) {
		return
	}

	abstracttimeprimitivetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeprimitivetype.Identifier != nil {
		stage.StageBranch(abstracttimeprimitivetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeprimitivetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (abstracttimeslicetype *AbstractTimeSliceType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(abstracttimeslicetype) {
		return
	}

	abstracttimeslicetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeslicetype.Identifier != nil {
		stage.StageBranch(abstracttimeslicetype.Identifier)
	}
	if abstracttimeslicetype.ValidTime != nil {
		stage.StageBranch(abstracttimeslicetype.ValidTime)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeslicetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (boundingshapetype *BoundingShapeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(boundingshapetype) {
		return
	}

	boundingshapetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if boundingshapetype.Envelope != nil {
		stage.StageBranch(boundingshapetype.Envelope)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetype *CodeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codetype) {
		return
	}

	codetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codewithauthoritytype *CodeWithAuthorityType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(codewithauthoritytype) {
		return
	}

	codewithauthoritytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (directpositiontype *DirectPositionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(directpositiontype) {
		return
	}

	directpositiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dynamicfeaturetype *DynamicFeatureType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(dynamicfeaturetype) {
		return
	}

	dynamicfeaturetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dynamicfeaturetype.Identifier != nil {
		stage.StageBranch(dynamicfeaturetype.Identifier)
	}
	if dynamicfeaturetype.BoundedBy != nil {
		stage.StageBranch(dynamicfeaturetype.BoundedBy)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range dynamicfeaturetype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (envelopetype *EnvelopeType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(envelopetype) {
		return
	}

	envelopetype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if envelopetype.LowerCorner != nil {
		stage.StageBranch(envelopetype.LowerCorner)
	}
	if envelopetype.UpperCorner != nil {
		stage.StageBranch(envelopetype.UpperCorner)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pointtype *PointType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(pointtype) {
		return
	}

	pointtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pointtype.Identifier != nil {
		stage.StageBranch(pointtype.Identifier)
	}
	if pointtype.Pos != nil {
		stage.StageBranch(pointtype.Pos)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range pointtype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (timeinstanttype *TimeInstantType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timeinstanttype) {
		return
	}

	timeinstanttype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if timeinstanttype.Identifier != nil {
		stage.StageBranch(timeinstanttype.Identifier)
	}
	if timeinstanttype.TimePosition != nil {
		stage.StageBranch(timeinstanttype.TimePosition)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range timeinstanttype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (timeperiodtype *TimePeriodType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timeperiodtype) {
		return
	}

	timeperiodtype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if timeperiodtype.Identifier != nil {
		stage.StageBranch(timeperiodtype.Identifier)
	}
	if timeperiodtype.BeginPosition != nil {
		stage.StageBranch(timeperiodtype.BeginPosition)
	}
	if timeperiodtype.EndPosition != nil {
		stage.StageBranch(timeperiodtype.EndPosition)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range timeperiodtype.NameXSD {
		stage.StageBranch(_codetype)
	}

}

func (timepositiontype *TimePositionType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timepositiontype) {
		return
	}

	timepositiontype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongStageBranch(stage *Stage) {

	// check if instance is already staged
	if stage.IsStaged(timeprimitivepropertytype) {
		return
	}

	timeprimitivepropertytype.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if timeprimitivepropertytype.TimeInstant != nil {
		stage.StageBranch(timeprimitivepropertytype.TimeInstant)
	}
	if timeprimitivepropertytype.TimePeriod != nil {
		stage.StageBranch(timeprimitivepropertytype.TimePeriod)
	}

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
	case *AbstractFeatureMemberType:
		toT := GongCopyBranchAbstractFeatureMemberType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractFeatureType:
		toT := GongCopyBranchAbstractFeatureType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractGMLType:
		toT := GongCopyBranchAbstractGMLType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractGeometricPrimitiveType:
		toT := GongCopyBranchAbstractGeometricPrimitiveType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractGeometryType:
		toT := GongCopyBranchAbstractGeometryType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractMetadataPropertyType:
		toT := GongCopyBranchAbstractMetadataPropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractTimeGeometricPrimitiveType:
		toT := GongCopyBranchAbstractTimeGeometricPrimitiveType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractTimeObjectType:
		toT := GongCopyBranchAbstractTimeObjectType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractTimePrimitiveType:
		toT := GongCopyBranchAbstractTimePrimitiveType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *AbstractTimeSliceType:
		toT := GongCopyBranchAbstractTimeSliceType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *BoundingShapeType:
		toT := GongCopyBranchBoundingShapeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeType:
		toT := GongCopyBranchCodeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *CodeWithAuthorityType:
		toT := GongCopyBranchCodeWithAuthorityType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DirectPositionType:
		toT := GongCopyBranchDirectPositionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DynamicFeatureType:
		toT := GongCopyBranchDynamicFeatureType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EnvelopeType:
		toT := GongCopyBranchEnvelopeType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *PointType:
		toT := GongCopyBranchPointType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TimeInstantType:
		toT := GongCopyBranchTimeInstantType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TimePeriodType:
		toT := GongCopyBranchTimePeriodType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TimePositionType:
		toT := GongCopyBranchTimePositionType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *TimePrimitivePropertyType:
		toT := GongCopyBranchTimePrimitivePropertyType(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchAbstractFeatureMemberType(mapOrigCopy map[any]any, abstractfeaturemembertypeFrom *AbstractFeatureMemberType) (abstractfeaturemembertypeTo *AbstractFeatureMemberType) {
	var alreadyCopied bool
	abstractfeaturemembertypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractfeaturemembertypeFrom)
	if alreadyCopied {
		return
	}
	abstractfeaturemembertypeFrom.GongCopyBasicFields(abstractfeaturemembertypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractFeatureType(mapOrigCopy map[any]any, abstractfeaturetypeFrom *AbstractFeatureType) (abstractfeaturetypeTo *AbstractFeatureType) {
	var alreadyCopied bool
	abstractfeaturetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractfeaturetypeFrom)
	if alreadyCopied {
		return
	}
	abstractfeaturetypeFrom.GongCopyBasicFields(abstractfeaturetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstractfeaturetypeFrom.Identifier != nil {
		abstractfeaturetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstractfeaturetypeFrom.Identifier)
	}
	if abstractfeaturetypeFrom.BoundedBy != nil {
		abstractfeaturetypeTo.BoundedBy = GongCopyBranchBoundingShapeType(mapOrigCopy, abstractfeaturetypeFrom.BoundedBy)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractfeaturetypeFrom.NameXSD {
		abstractfeaturetypeTo.NameXSD = append(abstractfeaturetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAbstractGMLType(mapOrigCopy map[any]any, abstractgmltypeFrom *AbstractGMLType) (abstractgmltypeTo *AbstractGMLType) {
	var alreadyCopied bool
	abstractgmltypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractgmltypeFrom)
	if alreadyCopied {
		return
	}
	abstractgmltypeFrom.GongCopyBasicFields(abstractgmltypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstractgmltypeFrom.Identifier != nil {
		abstractgmltypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstractgmltypeFrom.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgmltypeFrom.NameXSD {
		abstractgmltypeTo.NameXSD = append(abstractgmltypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAbstractGeometricPrimitiveType(mapOrigCopy map[any]any, abstractgeometricprimitivetypeFrom *AbstractGeometricPrimitiveType) (abstractgeometricprimitivetypeTo *AbstractGeometricPrimitiveType) {
	var alreadyCopied bool
	abstractgeometricprimitivetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractgeometricprimitivetypeFrom)
	if alreadyCopied {
		return
	}
	abstractgeometricprimitivetypeFrom.GongCopyBasicFields(abstractgeometricprimitivetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstractgeometricprimitivetypeFrom.Identifier != nil {
		abstractgeometricprimitivetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstractgeometricprimitivetypeFrom.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgeometricprimitivetypeFrom.NameXSD {
		abstractgeometricprimitivetypeTo.NameXSD = append(abstractgeometricprimitivetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAbstractGeometryType(mapOrigCopy map[any]any, abstractgeometrytypeFrom *AbstractGeometryType) (abstractgeometrytypeTo *AbstractGeometryType) {
	var alreadyCopied bool
	abstractgeometrytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractgeometrytypeFrom)
	if alreadyCopied {
		return
	}
	abstractgeometrytypeFrom.GongCopyBasicFields(abstractgeometrytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstractgeometrytypeFrom.Identifier != nil {
		abstractgeometrytypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstractgeometrytypeFrom.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgeometrytypeFrom.NameXSD {
		abstractgeometrytypeTo.NameXSD = append(abstractgeometrytypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAbstractMetadataPropertyType(mapOrigCopy map[any]any, abstractmetadatapropertytypeFrom *AbstractMetadataPropertyType) (abstractmetadatapropertytypeTo *AbstractMetadataPropertyType) {
	var alreadyCopied bool
	abstractmetadatapropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstractmetadatapropertytypeFrom)
	if alreadyCopied {
		return
	}
	abstractmetadatapropertytypeFrom.GongCopyBasicFields(abstractmetadatapropertytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchAbstractTimeGeometricPrimitiveType(mapOrigCopy map[any]any, abstracttimegeometricprimitivetypeFrom *AbstractTimeGeometricPrimitiveType) (abstracttimegeometricprimitivetypeTo *AbstractTimeGeometricPrimitiveType) {
	var alreadyCopied bool
	abstracttimegeometricprimitivetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstracttimegeometricprimitivetypeFrom)
	if alreadyCopied {
		return
	}
	abstracttimegeometricprimitivetypeFrom.GongCopyBasicFields(abstracttimegeometricprimitivetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimegeometricprimitivetypeFrom.Identifier != nil {
		abstracttimegeometricprimitivetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstracttimegeometricprimitivetypeFrom.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimegeometricprimitivetypeFrom.NameXSD {
		abstracttimegeometricprimitivetypeTo.NameXSD = append(abstracttimegeometricprimitivetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAbstractTimeObjectType(mapOrigCopy map[any]any, abstracttimeobjecttypeFrom *AbstractTimeObjectType) (abstracttimeobjecttypeTo *AbstractTimeObjectType) {
	var alreadyCopied bool
	abstracttimeobjecttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstracttimeobjecttypeFrom)
	if alreadyCopied {
		return
	}
	abstracttimeobjecttypeFrom.GongCopyBasicFields(abstracttimeobjecttypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeobjecttypeFrom.Identifier != nil {
		abstracttimeobjecttypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstracttimeobjecttypeFrom.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeobjecttypeFrom.NameXSD {
		abstracttimeobjecttypeTo.NameXSD = append(abstracttimeobjecttypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAbstractTimePrimitiveType(mapOrigCopy map[any]any, abstracttimeprimitivetypeFrom *AbstractTimePrimitiveType) (abstracttimeprimitivetypeTo *AbstractTimePrimitiveType) {
	var alreadyCopied bool
	abstracttimeprimitivetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstracttimeprimitivetypeFrom)
	if alreadyCopied {
		return
	}
	abstracttimeprimitivetypeFrom.GongCopyBasicFields(abstracttimeprimitivetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeprimitivetypeFrom.Identifier != nil {
		abstracttimeprimitivetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstracttimeprimitivetypeFrom.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeprimitivetypeFrom.NameXSD {
		abstracttimeprimitivetypeTo.NameXSD = append(abstracttimeprimitivetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchAbstractTimeSliceType(mapOrigCopy map[any]any, abstracttimeslicetypeFrom *AbstractTimeSliceType) (abstracttimeslicetypeTo *AbstractTimeSliceType) {
	var alreadyCopied bool
	abstracttimeslicetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, abstracttimeslicetypeFrom)
	if alreadyCopied {
		return
	}
	abstracttimeslicetypeFrom.GongCopyBasicFields(abstracttimeslicetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeslicetypeFrom.Identifier != nil {
		abstracttimeslicetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, abstracttimeslicetypeFrom.Identifier)
	}
	if abstracttimeslicetypeFrom.ValidTime != nil {
		abstracttimeslicetypeTo.ValidTime = GongCopyBranchTimePrimitivePropertyType(mapOrigCopy, abstracttimeslicetypeFrom.ValidTime)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeslicetypeFrom.NameXSD {
		abstracttimeslicetypeTo.NameXSD = append(abstracttimeslicetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchBoundingShapeType(mapOrigCopy map[any]any, boundingshapetypeFrom *BoundingShapeType) (boundingshapetypeTo *BoundingShapeType) {
	var alreadyCopied bool
	boundingshapetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, boundingshapetypeFrom)
	if alreadyCopied {
		return
	}
	boundingshapetypeFrom.GongCopyBasicFields(boundingshapetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if boundingshapetypeFrom.Envelope != nil {
		boundingshapetypeTo.Envelope = GongCopyBranchEnvelopeType(mapOrigCopy, boundingshapetypeFrom.Envelope)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeType(mapOrigCopy map[any]any, codetypeFrom *CodeType) (codetypeTo *CodeType) {
	var alreadyCopied bool
	codetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codetypeFrom)
	if alreadyCopied {
		return
	}
	codetypeFrom.GongCopyBasicFields(codetypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchCodeWithAuthorityType(mapOrigCopy map[any]any, codewithauthoritytypeFrom *CodeWithAuthorityType) (codewithauthoritytypeTo *CodeWithAuthorityType) {
	var alreadyCopied bool
	codewithauthoritytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, codewithauthoritytypeFrom)
	if alreadyCopied {
		return
	}
	codewithauthoritytypeFrom.GongCopyBasicFields(codewithauthoritytypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDirectPositionType(mapOrigCopy map[any]any, directpositiontypeFrom *DirectPositionType) (directpositiontypeTo *DirectPositionType) {
	var alreadyCopied bool
	directpositiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, directpositiontypeFrom)
	if alreadyCopied {
		return
	}
	directpositiontypeFrom.GongCopyBasicFields(directpositiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDynamicFeatureType(mapOrigCopy map[any]any, dynamicfeaturetypeFrom *DynamicFeatureType) (dynamicfeaturetypeTo *DynamicFeatureType) {
	var alreadyCopied bool
	dynamicfeaturetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, dynamicfeaturetypeFrom)
	if alreadyCopied {
		return
	}
	dynamicfeaturetypeFrom.GongCopyBasicFields(dynamicfeaturetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if dynamicfeaturetypeFrom.Identifier != nil {
		dynamicfeaturetypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, dynamicfeaturetypeFrom.Identifier)
	}
	if dynamicfeaturetypeFrom.BoundedBy != nil {
		dynamicfeaturetypeTo.BoundedBy = GongCopyBranchBoundingShapeType(mapOrigCopy, dynamicfeaturetypeFrom.BoundedBy)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range dynamicfeaturetypeFrom.NameXSD {
		dynamicfeaturetypeTo.NameXSD = append(dynamicfeaturetypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchEnvelopeType(mapOrigCopy map[any]any, envelopetypeFrom *EnvelopeType) (envelopetypeTo *EnvelopeType) {
	var alreadyCopied bool
	envelopetypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, envelopetypeFrom)
	if alreadyCopied {
		return
	}
	envelopetypeFrom.GongCopyBasicFields(envelopetypeTo)

	//insertion point for the staging of instances referenced by pointers
	if envelopetypeFrom.LowerCorner != nil {
		envelopetypeTo.LowerCorner = GongCopyBranchDirectPositionType(mapOrigCopy, envelopetypeFrom.LowerCorner)
	}
	if envelopetypeFrom.UpperCorner != nil {
		envelopetypeTo.UpperCorner = GongCopyBranchDirectPositionType(mapOrigCopy, envelopetypeFrom.UpperCorner)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchPointType(mapOrigCopy map[any]any, pointtypeFrom *PointType) (pointtypeTo *PointType) {
	var alreadyCopied bool
	pointtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, pointtypeFrom)
	if alreadyCopied {
		return
	}
	pointtypeFrom.GongCopyBasicFields(pointtypeTo)

	//insertion point for the staging of instances referenced by pointers
	if pointtypeFrom.Identifier != nil {
		pointtypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, pointtypeFrom.Identifier)
	}
	if pointtypeFrom.Pos != nil {
		pointtypeTo.Pos = GongCopyBranchDirectPositionType(mapOrigCopy, pointtypeFrom.Pos)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range pointtypeFrom.NameXSD {
		pointtypeTo.NameXSD = append(pointtypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchTimeInstantType(mapOrigCopy map[any]any, timeinstanttypeFrom *TimeInstantType) (timeinstanttypeTo *TimeInstantType) {
	var alreadyCopied bool
	timeinstanttypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timeinstanttypeFrom)
	if alreadyCopied {
		return
	}
	timeinstanttypeFrom.GongCopyBasicFields(timeinstanttypeTo)

	//insertion point for the staging of instances referenced by pointers
	if timeinstanttypeFrom.Identifier != nil {
		timeinstanttypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, timeinstanttypeFrom.Identifier)
	}
	if timeinstanttypeFrom.TimePosition != nil {
		timeinstanttypeTo.TimePosition = GongCopyBranchTimePositionType(mapOrigCopy, timeinstanttypeFrom.TimePosition)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range timeinstanttypeFrom.NameXSD {
		timeinstanttypeTo.NameXSD = append(timeinstanttypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchTimePeriodType(mapOrigCopy map[any]any, timeperiodtypeFrom *TimePeriodType) (timeperiodtypeTo *TimePeriodType) {
	var alreadyCopied bool
	timeperiodtypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timeperiodtypeFrom)
	if alreadyCopied {
		return
	}
	timeperiodtypeFrom.GongCopyBasicFields(timeperiodtypeTo)

	//insertion point for the staging of instances referenced by pointers
	if timeperiodtypeFrom.Identifier != nil {
		timeperiodtypeTo.Identifier = GongCopyBranchCodeWithAuthorityType(mapOrigCopy, timeperiodtypeFrom.Identifier)
	}
	if timeperiodtypeFrom.BeginPosition != nil {
		timeperiodtypeTo.BeginPosition = GongCopyBranchTimePositionType(mapOrigCopy, timeperiodtypeFrom.BeginPosition)
	}
	if timeperiodtypeFrom.EndPosition != nil {
		timeperiodtypeTo.EndPosition = GongCopyBranchTimePositionType(mapOrigCopy, timeperiodtypeFrom.EndPosition)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range timeperiodtypeFrom.NameXSD {
		timeperiodtypeTo.NameXSD = append(timeperiodtypeTo.NameXSD, GongCopyBranchCodeType(mapOrigCopy, _codetype))
	}

	return
}

func GongCopyBranchTimePositionType(mapOrigCopy map[any]any, timepositiontypeFrom *TimePositionType) (timepositiontypeTo *TimePositionType) {
	var alreadyCopied bool
	timepositiontypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timepositiontypeFrom)
	if alreadyCopied {
		return
	}
	timepositiontypeFrom.GongCopyBasicFields(timepositiontypeTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchTimePrimitivePropertyType(mapOrigCopy map[any]any, timeprimitivepropertytypeFrom *TimePrimitivePropertyType) (timeprimitivepropertytypeTo *TimePrimitivePropertyType) {
	var alreadyCopied bool
	timeprimitivepropertytypeTo, alreadyCopied = __gong__copyBranchCheck(mapOrigCopy, timeprimitivepropertytypeFrom)
	if alreadyCopied {
		return
	}
	timeprimitivepropertytypeFrom.GongCopyBasicFields(timeprimitivepropertytypeTo)

	//insertion point for the staging of instances referenced by pointers
	if timeprimitivepropertytypeFrom.TimeInstant != nil {
		timeprimitivepropertytypeTo.TimeInstant = GongCopyBranchTimeInstantType(mapOrigCopy, timeprimitivepropertytypeFrom.TimeInstant)
	}
	if timeprimitivepropertytypeFrom.TimePeriod != nil {
		timeprimitivepropertytypeTo.TimePeriod = GongCopyBranchTimePeriodType(mapOrigCopy, timeprimitivepropertytypeFrom.TimePeriod)
	}

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
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractfeaturemembertype) {
		return
	}

	abstractfeaturemembertype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstractfeaturetype *AbstractFeatureType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractfeaturetype) {
		return
	}

	abstractfeaturetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractfeaturetype.Identifier != nil {
		stage.UnstageBranch(abstractfeaturetype.Identifier)
	}
	if abstractfeaturetype.BoundedBy != nil {
		stage.UnstageBranch(abstractfeaturetype.BoundedBy)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractfeaturetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (abstractgmltype *AbstractGMLType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractgmltype) {
		return
	}

	abstractgmltype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractgmltype.Identifier != nil {
		stage.UnstageBranch(abstractgmltype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgmltype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractgeometricprimitivetype) {
		return
	}

	abstractgeometricprimitivetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractgeometricprimitivetype.Identifier != nil {
		stage.UnstageBranch(abstractgeometricprimitivetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgeometricprimitivetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (abstractgeometrytype *AbstractGeometryType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractgeometrytype) {
		return
	}

	abstractgeometrytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstractgeometrytype.Identifier != nil {
		stage.UnstageBranch(abstractgeometrytype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstractgeometrytype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstractmetadatapropertytype) {
		return
	}

	abstractmetadatapropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstracttimegeometricprimitivetype) {
		return
	}

	abstracttimegeometricprimitivetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimegeometricprimitivetype.Identifier != nil {
		stage.UnstageBranch(abstracttimegeometricprimitivetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimegeometricprimitivetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (abstracttimeobjecttype *AbstractTimeObjectType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstracttimeobjecttype) {
		return
	}

	abstracttimeobjecttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeobjecttype.Identifier != nil {
		stage.UnstageBranch(abstracttimeobjecttype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeobjecttype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstracttimeprimitivetype) {
		return
	}

	abstracttimeprimitivetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeprimitivetype.Identifier != nil {
		stage.UnstageBranch(abstracttimeprimitivetype.Identifier)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeprimitivetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (abstracttimeslicetype *AbstractTimeSliceType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(abstracttimeslicetype) {
		return
	}

	abstracttimeslicetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if abstracttimeslicetype.Identifier != nil {
		stage.UnstageBranch(abstracttimeslicetype.Identifier)
	}
	if abstracttimeslicetype.ValidTime != nil {
		stage.UnstageBranch(abstracttimeslicetype.ValidTime)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range abstracttimeslicetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (boundingshapetype *BoundingShapeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(boundingshapetype) {
		return
	}

	boundingshapetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if boundingshapetype.Envelope != nil {
		stage.UnstageBranch(boundingshapetype.Envelope)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codetype *CodeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codetype) {
		return
	}

	codetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (codewithauthoritytype *CodeWithAuthorityType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(codewithauthoritytype) {
		return
	}

	codewithauthoritytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (directpositiontype *DirectPositionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(directpositiontype) {
		return
	}

	directpositiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (dynamicfeaturetype *DynamicFeatureType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(dynamicfeaturetype) {
		return
	}

	dynamicfeaturetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if dynamicfeaturetype.Identifier != nil {
		stage.UnstageBranch(dynamicfeaturetype.Identifier)
	}
	if dynamicfeaturetype.BoundedBy != nil {
		stage.UnstageBranch(dynamicfeaturetype.BoundedBy)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range dynamicfeaturetype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (envelopetype *EnvelopeType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(envelopetype) {
		return
	}

	envelopetype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if envelopetype.LowerCorner != nil {
		stage.UnstageBranch(envelopetype.LowerCorner)
	}
	if envelopetype.UpperCorner != nil {
		stage.UnstageBranch(envelopetype.UpperCorner)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (pointtype *PointType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(pointtype) {
		return
	}

	pointtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if pointtype.Identifier != nil {
		stage.UnstageBranch(pointtype.Identifier)
	}
	if pointtype.Pos != nil {
		stage.UnstageBranch(pointtype.Pos)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range pointtype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (timeinstanttype *TimeInstantType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timeinstanttype) {
		return
	}

	timeinstanttype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if timeinstanttype.Identifier != nil {
		stage.UnstageBranch(timeinstanttype.Identifier)
	}
	if timeinstanttype.TimePosition != nil {
		stage.UnstageBranch(timeinstanttype.TimePosition)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range timeinstanttype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (timeperiodtype *TimePeriodType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timeperiodtype) {
		return
	}

	timeperiodtype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if timeperiodtype.Identifier != nil {
		stage.UnstageBranch(timeperiodtype.Identifier)
	}
	if timeperiodtype.BeginPosition != nil {
		stage.UnstageBranch(timeperiodtype.BeginPosition)
	}
	if timeperiodtype.EndPosition != nil {
		stage.UnstageBranch(timeperiodtype.EndPosition)
	}

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _codetype := range timeperiodtype.NameXSD {
		stage.UnstageBranch(_codetype)
	}

}

func (timepositiontype *TimePositionType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timepositiontype) {
		return
	}

	timepositiontype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (timeprimitivepropertytype *TimePrimitivePropertyType) GongUnstageBranch(stage *Stage) {

	// check if instance is already staged
	if !stage.IsStaged(timeprimitivepropertytype) {
		return
	}

	timeprimitivepropertytype.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if timeprimitivepropertytype.TimeInstant != nil {
		stage.UnstageBranch(timeprimitivepropertytype.TimeInstant)
	}
	if timeprimitivepropertytype.TimePeriod != nil {
		stage.UnstageBranch(timeprimitivepropertytype.TimePeriod)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *AbstractFeatureMemberType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractFeatureMemberType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AbstractFeatureType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractFeatureType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	__gong__reconstructPointer(&reference.BoundedBy, stage.BoundingShapeTypes_reference, instance.BoundedBy)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AbstractGMLType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractGMLType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AbstractGeometricPrimitiveType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractGeometricPrimitiveType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AbstractGeometryType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractGeometryType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AbstractMetadataPropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractMetadataPropertyType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *AbstractTimeGeometricPrimitiveType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractTimeGeometricPrimitiveType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AbstractTimeObjectType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractTimeObjectType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AbstractTimePrimitiveType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractTimePrimitiveType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *AbstractTimeSliceType) GongReconstructPointersFromReferences(stage *Stage, instance *AbstractTimeSliceType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	__gong__reconstructPointer(&reference.ValidTime, stage.TimePrimitivePropertyTypes_reference, instance.ValidTime)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *BoundingShapeType) GongReconstructPointersFromReferences(stage *Stage, instance *BoundingShapeType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Envelope, stage.EnvelopeTypes_reference, instance.Envelope)
	// insertion point for slice of pointers field
}

func (reference *CodeType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *CodeWithAuthorityType) GongReconstructPointersFromReferences(stage *Stage, instance *CodeWithAuthorityType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DirectPositionType) GongReconstructPointersFromReferences(stage *Stage, instance *DirectPositionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *DynamicFeatureType) GongReconstructPointersFromReferences(stage *Stage, instance *DynamicFeatureType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	__gong__reconstructPointer(&reference.BoundedBy, stage.BoundingShapeTypes_reference, instance.BoundedBy)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *EnvelopeType) GongReconstructPointersFromReferences(stage *Stage, instance *EnvelopeType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.LowerCorner, stage.DirectPositionTypes_reference, instance.LowerCorner)
	__gong__reconstructPointer(&reference.UpperCorner, stage.DirectPositionTypes_reference, instance.UpperCorner)
	// insertion point for slice of pointers field
}

func (reference *PointType) GongReconstructPointersFromReferences(stage *Stage, instance *PointType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	__gong__reconstructPointer(&reference.Pos, stage.DirectPositionTypes_reference, instance.Pos)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *TimeInstantType) GongReconstructPointersFromReferences(stage *Stage, instance *TimeInstantType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	__gong__reconstructPointer(&reference.TimePosition, stage.TimePositionTypes_reference, instance.TimePosition)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *TimePeriodType) GongReconstructPointersFromReferences(stage *Stage, instance *TimePeriodType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.Identifier, stage.CodeWithAuthorityTypes_reference, instance.Identifier)
	__gong__reconstructPointer(&reference.BeginPosition, stage.TimePositionTypes_reference, instance.BeginPosition)
	__gong__reconstructPointer(&reference.EndPosition, stage.TimePositionTypes_reference, instance.EndPosition)
	// insertion point for slice of pointers field
	__gong__reconstructSliceOfPointersFromReferences(&reference.NameXSD, stage.CodeTypes_reference, instance.NameXSD)
}

func (reference *TimePositionType) GongReconstructPointersFromReferences(stage *Stage, instance *TimePositionType) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *TimePrimitivePropertyType) GongReconstructPointersFromReferences(stage *Stage, instance *TimePrimitivePropertyType) {
	// insertion point for pointers field
	__gong__reconstructPointer(&reference.TimeInstant, stage.TimeInstantTypes_reference, instance.TimeInstant)
	__gong__reconstructPointer(&reference.TimePeriod, stage.TimePeriodTypes_reference, instance.TimePeriod)
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *AbstractFeatureMemberType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AbstractFeatureType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.BoundedBy, stage.BoundingShapeTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AbstractGMLType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AbstractGeometricPrimitiveType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AbstractGeometryType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AbstractMetadataPropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *AbstractTimeGeometricPrimitiveType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AbstractTimeObjectType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AbstractTimePrimitiveType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *AbstractTimeSliceType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.ValidTime, stage.TimePrimitivePropertyTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *BoundingShapeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Envelope, stage.EnvelopeTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *CodeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *CodeWithAuthorityType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DirectPositionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *DynamicFeatureType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.BoundedBy, stage.BoundingShapeTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *EnvelopeType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.LowerCorner, stage.DirectPositionTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.UpperCorner, stage.DirectPositionTypes_instance)
	// insertion point for slice of pointers fields
}

func (reference *PointType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.Pos, stage.DirectPositionTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *TimeInstantType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.TimePosition, stage.TimePositionTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *TimePeriodType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.Identifier, stage.CodeWithAuthorityTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.BeginPosition, stage.TimePositionTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.EndPosition, stage.TimePositionTypes_instance)
	// insertion point for slice of pointers fields
	__gong__reconstructSliceOfPointersFromInstances(&reference.NameXSD, stage.CodeTypes_instance)
}

func (reference *TimePositionType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *TimePrimitivePropertyType) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	__gong__reconstructPointerFromInstance(&reference.TimeInstant, stage.TimeInstantTypes_instance)
	__gong__reconstructPointerFromInstance(&reference.TimePeriod, stage.TimePeriodTypes_instance)
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractfeaturemembertype *AbstractFeatureMemberType) GongDiff(stage *Stage, abstractfeaturemembertypeOther *AbstractFeatureMemberType) (diffs []string) {
	// insertion point for field diffs
	if abstractfeaturemembertype.Name != abstractfeaturemembertypeOther.Name {
		diffs = append(diffs, abstractfeaturemembertype.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractfeaturetype *AbstractFeatureType) GongDiff(stage *Stage, abstractfeaturetypeOther *AbstractFeatureType) (diffs []string) {
	// insertion point for field diffs
	if abstractfeaturetype.Name != abstractfeaturetypeOther.Name {
		diffs = append(diffs, abstractfeaturetype.GongMarshallField(stage, "Name"))
	}
	if abstractfeaturetype.Id != abstractfeaturetypeOther.Id {
		diffs = append(diffs, abstractfeaturetype.GongMarshallField(stage, "Id"))
	}
	if abstractfeaturetype.Description != abstractfeaturetypeOther.Description {
		diffs = append(diffs, abstractfeaturetype.GongMarshallField(stage, "Description"))
	}
	if abstractfeaturetype.Identifier != abstractfeaturetypeOther.Identifier {
		diffs = append(diffs, abstractfeaturetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstractfeaturetype, "NameXSD", abstractfeaturetypeOther.NameXSD, abstractfeaturetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if abstractfeaturetype.BoundedBy != abstractfeaturetypeOther.BoundedBy {
		diffs = append(diffs, abstractfeaturetype.GongMarshallField(stage, "BoundedBy"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractgmltype *AbstractGMLType) GongDiff(stage *Stage, abstractgmltypeOther *AbstractGMLType) (diffs []string) {
	// insertion point for field diffs
	if abstractgmltype.Name != abstractgmltypeOther.Name {
		diffs = append(diffs, abstractgmltype.GongMarshallField(stage, "Name"))
	}
	if abstractgmltype.Id != abstractgmltypeOther.Id {
		diffs = append(diffs, abstractgmltype.GongMarshallField(stage, "Id"))
	}
	if abstractgmltype.Description != abstractgmltypeOther.Description {
		diffs = append(diffs, abstractgmltype.GongMarshallField(stage, "Description"))
	}
	if abstractgmltype.Identifier != abstractgmltypeOther.Identifier {
		diffs = append(diffs, abstractgmltype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstractgmltype, "NameXSD", abstractgmltypeOther.NameXSD, abstractgmltype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractgeometricprimitivetype *AbstractGeometricPrimitiveType) GongDiff(stage *Stage, abstractgeometricprimitivetypeOther *AbstractGeometricPrimitiveType) (diffs []string) {
	// insertion point for field diffs
	if abstractgeometricprimitivetype.Name != abstractgeometricprimitivetypeOther.Name {
		diffs = append(diffs, abstractgeometricprimitivetype.GongMarshallField(stage, "Name"))
	}
	if abstractgeometricprimitivetype.Id != abstractgeometricprimitivetypeOther.Id {
		diffs = append(diffs, abstractgeometricprimitivetype.GongMarshallField(stage, "Id"))
	}
	if abstractgeometricprimitivetype.Description != abstractgeometricprimitivetypeOther.Description {
		diffs = append(diffs, abstractgeometricprimitivetype.GongMarshallField(stage, "Description"))
	}
	if abstractgeometricprimitivetype.Identifier != abstractgeometricprimitivetypeOther.Identifier {
		diffs = append(diffs, abstractgeometricprimitivetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstractgeometricprimitivetype, "NameXSD", abstractgeometricprimitivetypeOther.NameXSD, abstractgeometricprimitivetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if abstractgeometricprimitivetype.SrsName != abstractgeometricprimitivetypeOther.SrsName {
		diffs = append(diffs, abstractgeometricprimitivetype.GongMarshallField(stage, "SrsName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractgeometrytype *AbstractGeometryType) GongDiff(stage *Stage, abstractgeometrytypeOther *AbstractGeometryType) (diffs []string) {
	// insertion point for field diffs
	if abstractgeometrytype.Name != abstractgeometrytypeOther.Name {
		diffs = append(diffs, abstractgeometrytype.GongMarshallField(stage, "Name"))
	}
	if abstractgeometrytype.Id != abstractgeometrytypeOther.Id {
		diffs = append(diffs, abstractgeometrytype.GongMarshallField(stage, "Id"))
	}
	if abstractgeometrytype.Description != abstractgeometrytypeOther.Description {
		diffs = append(diffs, abstractgeometrytype.GongMarshallField(stage, "Description"))
	}
	if abstractgeometrytype.Identifier != abstractgeometrytypeOther.Identifier {
		diffs = append(diffs, abstractgeometrytype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstractgeometrytype, "NameXSD", abstractgeometrytypeOther.NameXSD, abstractgeometrytype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if abstractgeometrytype.SrsName != abstractgeometrytypeOther.SrsName {
		diffs = append(diffs, abstractgeometrytype.GongMarshallField(stage, "SrsName"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstractmetadatapropertytype *AbstractMetadataPropertyType) GongDiff(stage *Stage, abstractmetadatapropertytypeOther *AbstractMetadataPropertyType) (diffs []string) {
	// insertion point for field diffs
	if abstractmetadatapropertytype.Name != abstractmetadatapropertytypeOther.Name {
		diffs = append(diffs, abstractmetadatapropertytype.GongMarshallField(stage, "Name"))
	}
	if abstractmetadatapropertytype.Owns != abstractmetadatapropertytypeOther.Owns {
		diffs = append(diffs, abstractmetadatapropertytype.GongMarshallField(stage, "Owns"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstracttimegeometricprimitivetype *AbstractTimeGeometricPrimitiveType) GongDiff(stage *Stage, abstracttimegeometricprimitivetypeOther *AbstractTimeGeometricPrimitiveType) (diffs []string) {
	// insertion point for field diffs
	if abstracttimegeometricprimitivetype.Name != abstracttimegeometricprimitivetypeOther.Name {
		diffs = append(diffs, abstracttimegeometricprimitivetype.GongMarshallField(stage, "Name"))
	}
	if abstracttimegeometricprimitivetype.Id != abstracttimegeometricprimitivetypeOther.Id {
		diffs = append(diffs, abstracttimegeometricprimitivetype.GongMarshallField(stage, "Id"))
	}
	if abstracttimegeometricprimitivetype.Description != abstracttimegeometricprimitivetypeOther.Description {
		diffs = append(diffs, abstracttimegeometricprimitivetype.GongMarshallField(stage, "Description"))
	}
	if abstracttimegeometricprimitivetype.Identifier != abstracttimegeometricprimitivetypeOther.Identifier {
		diffs = append(diffs, abstracttimegeometricprimitivetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstracttimegeometricprimitivetype, "NameXSD", abstracttimegeometricprimitivetypeOther.NameXSD, abstracttimegeometricprimitivetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if abstracttimegeometricprimitivetype.Frame != abstracttimegeometricprimitivetypeOther.Frame {
		diffs = append(diffs, abstracttimegeometricprimitivetype.GongMarshallField(stage, "Frame"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstracttimeobjecttype *AbstractTimeObjectType) GongDiff(stage *Stage, abstracttimeobjecttypeOther *AbstractTimeObjectType) (diffs []string) {
	// insertion point for field diffs
	if abstracttimeobjecttype.Name != abstracttimeobjecttypeOther.Name {
		diffs = append(diffs, abstracttimeobjecttype.GongMarshallField(stage, "Name"))
	}
	if abstracttimeobjecttype.Id != abstracttimeobjecttypeOther.Id {
		diffs = append(diffs, abstracttimeobjecttype.GongMarshallField(stage, "Id"))
	}
	if abstracttimeobjecttype.Description != abstracttimeobjecttypeOther.Description {
		diffs = append(diffs, abstracttimeobjecttype.GongMarshallField(stage, "Description"))
	}
	if abstracttimeobjecttype.Identifier != abstracttimeobjecttypeOther.Identifier {
		diffs = append(diffs, abstracttimeobjecttype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstracttimeobjecttype, "NameXSD", abstracttimeobjecttypeOther.NameXSD, abstracttimeobjecttype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstracttimeprimitivetype *AbstractTimePrimitiveType) GongDiff(stage *Stage, abstracttimeprimitivetypeOther *AbstractTimePrimitiveType) (diffs []string) {
	// insertion point for field diffs
	if abstracttimeprimitivetype.Name != abstracttimeprimitivetypeOther.Name {
		diffs = append(diffs, abstracttimeprimitivetype.GongMarshallField(stage, "Name"))
	}
	if abstracttimeprimitivetype.Id != abstracttimeprimitivetypeOther.Id {
		diffs = append(diffs, abstracttimeprimitivetype.GongMarshallField(stage, "Id"))
	}
	if abstracttimeprimitivetype.Description != abstracttimeprimitivetypeOther.Description {
		diffs = append(diffs, abstracttimeprimitivetype.GongMarshallField(stage, "Description"))
	}
	if abstracttimeprimitivetype.Identifier != abstracttimeprimitivetypeOther.Identifier {
		diffs = append(diffs, abstracttimeprimitivetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstracttimeprimitivetype, "NameXSD", abstracttimeprimitivetypeOther.NameXSD, abstracttimeprimitivetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (abstracttimeslicetype *AbstractTimeSliceType) GongDiff(stage *Stage, abstracttimeslicetypeOther *AbstractTimeSliceType) (diffs []string) {
	// insertion point for field diffs
	if abstracttimeslicetype.Name != abstracttimeslicetypeOther.Name {
		diffs = append(diffs, abstracttimeslicetype.GongMarshallField(stage, "Name"))
	}
	if abstracttimeslicetype.Id != abstracttimeslicetypeOther.Id {
		diffs = append(diffs, abstracttimeslicetype.GongMarshallField(stage, "Id"))
	}
	if abstracttimeslicetype.Description != abstracttimeslicetypeOther.Description {
		diffs = append(diffs, abstracttimeslicetype.GongMarshallField(stage, "Description"))
	}
	if abstracttimeslicetype.Identifier != abstracttimeslicetypeOther.Identifier {
		diffs = append(diffs, abstracttimeslicetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, abstracttimeslicetype, "NameXSD", abstracttimeslicetypeOther.NameXSD, abstracttimeslicetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if abstracttimeslicetype.ValidTime != abstracttimeslicetypeOther.ValidTime {
		diffs = append(diffs, abstracttimeslicetype.GongMarshallField(stage, "ValidTime"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (boundingshapetype *BoundingShapeType) GongDiff(stage *Stage, boundingshapetypeOther *BoundingShapeType) (diffs []string) {
	// insertion point for field diffs
	if boundingshapetype.Name != boundingshapetypeOther.Name {
		diffs = append(diffs, boundingshapetype.GongMarshallField(stage, "Name"))
	}
	if boundingshapetype.NilReason != boundingshapetypeOther.NilReason {
		diffs = append(diffs, boundingshapetype.GongMarshallField(stage, "NilReason"))
	}
	if boundingshapetype.Envelope != boundingshapetypeOther.Envelope {
		diffs = append(diffs, boundingshapetype.GongMarshallField(stage, "Envelope"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codetype *CodeType) GongDiff(stage *Stage, codetypeOther *CodeType) (diffs []string) {
	// insertion point for field diffs
	if codetype.Name != codetypeOther.Name {
		diffs = append(diffs, codetype.GongMarshallField(stage, "Name"))
	}
	if codetype.CodeSpace != codetypeOther.CodeSpace {
		diffs = append(diffs, codetype.GongMarshallField(stage, "CodeSpace"))
	}
	if codetype.EnclosedText != codetypeOther.EnclosedText {
		diffs = append(diffs, codetype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (codewithauthoritytype *CodeWithAuthorityType) GongDiff(stage *Stage, codewithauthoritytypeOther *CodeWithAuthorityType) (diffs []string) {
	// insertion point for field diffs
	if codewithauthoritytype.Name != codewithauthoritytypeOther.Name {
		diffs = append(diffs, codewithauthoritytype.GongMarshallField(stage, "Name"))
	}
	if codewithauthoritytype.CodeSpace != codewithauthoritytypeOther.CodeSpace {
		diffs = append(diffs, codewithauthoritytype.GongMarshallField(stage, "CodeSpace"))
	}
	if codewithauthoritytype.EnclosedText != codewithauthoritytypeOther.EnclosedText {
		diffs = append(diffs, codewithauthoritytype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (directpositiontype *DirectPositionType) GongDiff(stage *Stage, directpositiontypeOther *DirectPositionType) (diffs []string) {
	// insertion point for field diffs
	if directpositiontype.Name != directpositiontypeOther.Name {
		diffs = append(diffs, directpositiontype.GongMarshallField(stage, "Name"))
	}
	if directpositiontype.SrsName != directpositiontypeOther.SrsName {
		diffs = append(diffs, directpositiontype.GongMarshallField(stage, "SrsName"))
	}
	if directpositiontype.EnclosedText != directpositiontypeOther.EnclosedText {
		diffs = append(diffs, directpositiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (dynamicfeaturetype *DynamicFeatureType) GongDiff(stage *Stage, dynamicfeaturetypeOther *DynamicFeatureType) (diffs []string) {
	// insertion point for field diffs
	if dynamicfeaturetype.Name != dynamicfeaturetypeOther.Name {
		diffs = append(diffs, dynamicfeaturetype.GongMarshallField(stage, "Name"))
	}
	if dynamicfeaturetype.Id != dynamicfeaturetypeOther.Id {
		diffs = append(diffs, dynamicfeaturetype.GongMarshallField(stage, "Id"))
	}
	if dynamicfeaturetype.Description != dynamicfeaturetypeOther.Description {
		diffs = append(diffs, dynamicfeaturetype.GongMarshallField(stage, "Description"))
	}
	if dynamicfeaturetype.Identifier != dynamicfeaturetypeOther.Identifier {
		diffs = append(diffs, dynamicfeaturetype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, dynamicfeaturetype, "NameXSD", dynamicfeaturetypeOther.NameXSD, dynamicfeaturetype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if dynamicfeaturetype.BoundedBy != dynamicfeaturetypeOther.BoundedBy {
		diffs = append(diffs, dynamicfeaturetype.GongMarshallField(stage, "BoundedBy"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (envelopetype *EnvelopeType) GongDiff(stage *Stage, envelopetypeOther *EnvelopeType) (diffs []string) {
	// insertion point for field diffs
	if envelopetype.Name != envelopetypeOther.Name {
		diffs = append(diffs, envelopetype.GongMarshallField(stage, "Name"))
	}
	if envelopetype.SrsName != envelopetypeOther.SrsName {
		diffs = append(diffs, envelopetype.GongMarshallField(stage, "SrsName"))
	}
	if envelopetype.LowerCorner != envelopetypeOther.LowerCorner {
		diffs = append(diffs, envelopetype.GongMarshallField(stage, "LowerCorner"))
	}
	if envelopetype.UpperCorner != envelopetypeOther.UpperCorner {
		diffs = append(diffs, envelopetype.GongMarshallField(stage, "UpperCorner"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (pointtype *PointType) GongDiff(stage *Stage, pointtypeOther *PointType) (diffs []string) {
	// insertion point for field diffs
	if pointtype.Name != pointtypeOther.Name {
		diffs = append(diffs, pointtype.GongMarshallField(stage, "Name"))
	}
	if pointtype.Id != pointtypeOther.Id {
		diffs = append(diffs, pointtype.GongMarshallField(stage, "Id"))
	}
	if pointtype.Description != pointtypeOther.Description {
		diffs = append(diffs, pointtype.GongMarshallField(stage, "Description"))
	}
	if pointtype.Identifier != pointtypeOther.Identifier {
		diffs = append(diffs, pointtype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, pointtype, "NameXSD", pointtypeOther.NameXSD, pointtype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if pointtype.SrsName != pointtypeOther.SrsName {
		diffs = append(diffs, pointtype.GongMarshallField(stage, "SrsName"))
	}
	if pointtype.Pos != pointtypeOther.Pos {
		diffs = append(diffs, pointtype.GongMarshallField(stage, "Pos"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (timeinstanttype *TimeInstantType) GongDiff(stage *Stage, timeinstanttypeOther *TimeInstantType) (diffs []string) {
	// insertion point for field diffs
	if timeinstanttype.Name != timeinstanttypeOther.Name {
		diffs = append(diffs, timeinstanttype.GongMarshallField(stage, "Name"))
	}
	if timeinstanttype.Id != timeinstanttypeOther.Id {
		diffs = append(diffs, timeinstanttype.GongMarshallField(stage, "Id"))
	}
	if timeinstanttype.Description != timeinstanttypeOther.Description {
		diffs = append(diffs, timeinstanttype.GongMarshallField(stage, "Description"))
	}
	if timeinstanttype.Identifier != timeinstanttypeOther.Identifier {
		diffs = append(diffs, timeinstanttype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, timeinstanttype, "NameXSD", timeinstanttypeOther.NameXSD, timeinstanttype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if timeinstanttype.Frame != timeinstanttypeOther.Frame {
		diffs = append(diffs, timeinstanttype.GongMarshallField(stage, "Frame"))
	}
	if timeinstanttype.TimePosition != timeinstanttypeOther.TimePosition {
		diffs = append(diffs, timeinstanttype.GongMarshallField(stage, "TimePosition"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (timeperiodtype *TimePeriodType) GongDiff(stage *Stage, timeperiodtypeOther *TimePeriodType) (diffs []string) {
	// insertion point for field diffs
	if timeperiodtype.Name != timeperiodtypeOther.Name {
		diffs = append(diffs, timeperiodtype.GongMarshallField(stage, "Name"))
	}
	if timeperiodtype.Id != timeperiodtypeOther.Id {
		diffs = append(diffs, timeperiodtype.GongMarshallField(stage, "Id"))
	}
	if timeperiodtype.Description != timeperiodtypeOther.Description {
		diffs = append(diffs, timeperiodtype.GongMarshallField(stage, "Description"))
	}
	if timeperiodtype.Identifier != timeperiodtypeOther.Identifier {
		diffs = append(diffs, timeperiodtype.GongMarshallField(stage, "Identifier"))
	}
	if ops := __gong__diffSliceOfPointers(stage, timeperiodtype, "NameXSD", timeperiodtypeOther.NameXSD, timeperiodtype.NameXSD); ops != "" {
		diffs = append(diffs, ops)
	}
	if timeperiodtype.Frame != timeperiodtypeOther.Frame {
		diffs = append(diffs, timeperiodtype.GongMarshallField(stage, "Frame"))
	}
	if timeperiodtype.BeginPosition != timeperiodtypeOther.BeginPosition {
		diffs = append(diffs, timeperiodtype.GongMarshallField(stage, "BeginPosition"))
	}
	if timeperiodtype.EndPosition != timeperiodtypeOther.EndPosition {
		diffs = append(diffs, timeperiodtype.GongMarshallField(stage, "EndPosition"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (timepositiontype *TimePositionType) GongDiff(stage *Stage, timepositiontypeOther *TimePositionType) (diffs []string) {
	// insertion point for field diffs
	if timepositiontype.Name != timepositiontypeOther.Name {
		diffs = append(diffs, timepositiontype.GongMarshallField(stage, "Name"))
	}
	if timepositiontype.Frame != timepositiontypeOther.Frame {
		diffs = append(diffs, timepositiontype.GongMarshallField(stage, "Frame"))
	}
	if timepositiontype.IndeterminatePosition != timepositiontypeOther.IndeterminatePosition {
		diffs = append(diffs, timepositiontype.GongMarshallField(stage, "IndeterminatePosition"))
	}
	if timepositiontype.EnclosedText != timepositiontypeOther.EnclosedText {
		diffs = append(diffs, timepositiontype.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (timeprimitivepropertytype *TimePrimitivePropertyType) GongDiff(stage *Stage, timeprimitivepropertytypeOther *TimePrimitivePropertyType) (diffs []string) {
	// insertion point for field diffs
	if timeprimitivepropertytype.Name != timeprimitivepropertytypeOther.Name {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "Name"))
	}
	if timeprimitivepropertytype.NilReason != timeprimitivepropertytypeOther.NilReason {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "NilReason"))
	}
	if timeprimitivepropertytype.Type_XSD != timeprimitivepropertytypeOther.Type_XSD {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "Type_XSD"))
	}
	if timeprimitivepropertytype.Href != timeprimitivepropertytypeOther.Href {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "Href"))
	}
	if timeprimitivepropertytype.Role != timeprimitivepropertytypeOther.Role {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "Role"))
	}
	if timeprimitivepropertytype.Arcrole != timeprimitivepropertytypeOther.Arcrole {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "Arcrole"))
	}
	if timeprimitivepropertytype.Title != timeprimitivepropertytypeOther.Title {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "Title"))
	}
	if timeprimitivepropertytype.Show != timeprimitivepropertytypeOther.Show {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "Show"))
	}
	if timeprimitivepropertytype.Actuate != timeprimitivepropertytypeOther.Actuate {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "Actuate"))
	}
	if timeprimitivepropertytype.TimeInstant != timeprimitivepropertytypeOther.TimeInstant {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "TimeInstant"))
	}
	if timeprimitivepropertytype.TimePeriod != timeprimitivepropertytypeOther.TimePeriod {
		diffs = append(diffs, timeprimitivepropertytype.GongMarshallField(stage, "TimePeriod"))
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
