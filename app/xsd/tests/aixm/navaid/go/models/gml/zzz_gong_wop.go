// generated code - do not edit
package gml

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type AbstractFeatureMemberType_WOP struct {
	// insertion point

	Name string
}

func (from *AbstractFeatureMemberType) GongCopyBasicFields(to *AbstractFeatureMemberType) {
	// insertion point
	*to = *from
}

type AbstractFeatureType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string
}

func (from *AbstractFeatureType) GongCopyBasicFields(to *AbstractFeatureType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
}

type AbstractGMLType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string
}

func (from *AbstractGMLType) GongCopyBasicFields(to *AbstractGMLType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
}

type AbstractGeometricPrimitiveType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	SrsName string
}

func (from *AbstractGeometricPrimitiveType) GongCopyBasicFields(to *AbstractGeometricPrimitiveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.SrsName = from.SrsName
}

type AbstractGeometryType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	SrsName string
}

func (from *AbstractGeometryType) GongCopyBasicFields(to *AbstractGeometryType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.SrsName = from.SrsName
}

type AbstractMetadataPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *AbstractMetadataPropertyType) GongCopyBasicFields(to *AbstractMetadataPropertyType) {
	// insertion point
	*to = *from
}

type AbstractTimeGeometricPrimitiveType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	Frame string
}

func (from *AbstractTimeGeometricPrimitiveType) GongCopyBasicFields(to *AbstractTimeGeometricPrimitiveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.Frame = from.Frame
}

type AbstractTimeObjectType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string
}

func (from *AbstractTimeObjectType) GongCopyBasicFields(to *AbstractTimeObjectType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
}

type AbstractTimePrimitiveType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string
}

func (from *AbstractTimePrimitiveType) GongCopyBasicFields(to *AbstractTimePrimitiveType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
}

type AbstractTimeSliceType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string
}

func (from *AbstractTimeSliceType) GongCopyBasicFields(to *AbstractTimeSliceType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
}

type BoundingShapeType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *BoundingShapeType) GongCopyBasicFields(to *BoundingShapeType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
}

type CodeType_WOP struct {
	// insertion point

	Name string

	CodeSpace string

	EnclosedText string
}

func (from *CodeType) GongCopyBasicFields(to *CodeType) {
	// insertion point
	*to = *from
}

type CodeWithAuthorityType_WOP struct {
	// insertion point

	Name string

	CodeSpace string

	EnclosedText string
}

func (from *CodeWithAuthorityType) GongCopyBasicFields(to *CodeWithAuthorityType) {
	// insertion point
	*to = *from
}

type DirectPositionType_WOP struct {
	// insertion point

	Name string

	SrsName string

	EnclosedText string
}

func (from *DirectPositionType) GongCopyBasicFields(to *DirectPositionType) {
	// insertion point
	*to = *from
}

type DynamicFeatureType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string
}

func (from *DynamicFeatureType) GongCopyBasicFields(to *DynamicFeatureType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
}

type EnvelopeType_WOP struct {
	// insertion point

	Name string

	SrsName string
}

func (from *EnvelopeType) GongCopyBasicFields(to *EnvelopeType) {
	// insertion point
	to.Name = from.Name
	to.SrsName = from.SrsName
}

type PointType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	SrsName string
}

func (from *PointType) GongCopyBasicFields(to *PointType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.SrsName = from.SrsName
}

type TimeInstantType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	Frame string
}

func (from *TimeInstantType) GongCopyBasicFields(to *TimeInstantType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.Frame = from.Frame
}

type TimePeriodType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	Frame string
}

func (from *TimePeriodType) GongCopyBasicFields(to *TimePeriodType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.Frame = from.Frame
}

type TimePositionType_WOP struct {
	// insertion point

	Name string

	Frame string

	IndeterminatePosition Enum_TimeIndeterminateValueType

	EnclosedText string
}

func (from *TimePositionType) GongCopyBasicFields(to *TimePositionType) {
	// insertion point
	*to = *from
}

type TimePrimitivePropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *TimePrimitivePropertyType) GongCopyBasicFields(to *TimePrimitivePropertyType) {
	// insertion point
	to.Name = from.Name
	to.NilReason = from.NilReason
	to.Type_XSD = from.Type_XSD
	to.Href = from.Href
	to.Role = from.Role
	to.Arcrole = from.Arcrole
	to.Title = from.Title
	to.Show = from.Show
	to.Actuate = from.Actuate
}

// end of insertion point
