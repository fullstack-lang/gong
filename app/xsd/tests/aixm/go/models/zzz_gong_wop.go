// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var _ = time.Hour

// insertion point
type AIXMBasicMessage_WOP struct {
	// insertion point

	Name string

	Id string

	SequenceNumber int

	AggregationType string
}

func (from *AIXMBasicMessage) GongCopyBasicFields(to *AIXMBasicMessage) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SequenceNumber = from.SequenceNumber
	to.AggregationType = from.AggregationType
}

type AIXMBasicMessageType_WOP struct {
	// insertion point

	Name string

	Id string

	SequenceNumber int

	AggregationType string
}

func (from *AIXMBasicMessageType) GongCopyBasicFields(to *AIXMBasicMessageType) {
	// insertion point
	to.Name = from.Name
	to.Id = from.Id
	to.SequenceNumber = from.SequenceNumber
	to.AggregationType = from.AggregationType
}

type BasicMessageMemberAIXMPropertyType_WOP struct {
	// insertion point

	Name string

	Owns bool

	NilReason string

	RemoteSchema string

	Type_XSD string

	Href string

	Role string

	Arcrole string

	Title string

	Show string

	Actuate string
}

func (from *BasicMessageMemberAIXMPropertyType) GongCopyBasicFields(to *BasicMessageMemberAIXMPropertyType) {
	// insertion point
	*to = *from
}

// end of insertion point
