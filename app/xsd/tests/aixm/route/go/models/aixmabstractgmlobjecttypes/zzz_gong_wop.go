// generated code - do not edit
package aixmabstractgmlobjecttypes

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type AbstractAIXMFeatureBaseType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string
}

func (from *AbstractAIXMFeatureBaseType) GongCopyBasicFields(to *AbstractAIXMFeatureBaseType) {
	// insertion point
	*to = *from
}

type AbstractAIXMFeatureType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string
}

func (from *AbstractAIXMFeatureType) GongCopyBasicFields(to *AbstractAIXMFeatureType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
}

type AbstractAIXMMessageBaseType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	SequenceNumber int
}

func (from *AbstractAIXMMessageBaseType) GongCopyBasicFields(to *AbstractAIXMMessageBaseType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.SequenceNumber = from.SequenceNumber
}

type AbstractAIXMMessageType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	SequenceNumber int

	AggregationType string
}

func (from *AbstractAIXMMessageType) GongCopyBasicFields(to *AbstractAIXMMessageType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.SequenceNumber = from.SequenceNumber
	to.AggregationType = from.AggregationType
}

type AbstractAIXMObjectType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *AbstractAIXMObjectType) GongCopyBasicFields(to *AbstractAIXMObjectType) {
	// insertion point
	*to = *from
}

type AbstractAIXMPropertyType_WOP struct {
	// insertion point

	Name string

	NilReason string
}

func (from *AbstractAIXMPropertyType) GongCopyBasicFields(to *AbstractAIXMPropertyType) {
	// insertion point
	*to = *from
}

type AbstractAIXMTimeSliceBaseType_WOP struct {
	// insertion point

	Name string

	Id string
}

func (from *AbstractAIXMTimeSliceBaseType) GongCopyBasicFields(to *AbstractAIXMTimeSliceBaseType) {
	// insertion point
	*to = *from
}

type AbstractAIXMTimeSliceType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string

	Interpretation Enum_Interpretation_SimpleType

	SequenceNumber int

	CorrectionNumber int
}

func (from *AbstractAIXMTimeSliceType) GongCopyBasicFields(to *AbstractAIXMTimeSliceType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.Description = from.Description
	to.Interpretation = from.Interpretation
	to.SequenceNumber = from.SequenceNumber
	to.CorrectionNumber = from.CorrectionNumber
}

type AbstractExtensionType_WOP struct {
	// insertion point

	Name string

	Id string

	Description string
}

func (from *AbstractExtensionType) GongCopyBasicFields(to *AbstractExtensionType) {
	// insertion point
	*to = *from
}

type FeatureMetadataPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *FeatureMetadataPropertyType) GongCopyBasicFields(to *FeatureMetadataPropertyType) {
	// insertion point
	*to = *from
}

type FeatureTimeSliceMetadataPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *FeatureTimeSliceMetadataPropertyType) GongCopyBasicFields(to *FeatureTimeSliceMetadataPropertyType) {
	// insertion point
	*to = *from
}

type MessageMetadataPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool
}

func (from *MessageMetadataPropertyType) GongCopyBasicFields(to *MessageMetadataPropertyType) {
	// insertion point
	*to = *from
}

// end of insertion point
