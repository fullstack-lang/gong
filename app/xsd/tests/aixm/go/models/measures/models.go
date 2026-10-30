// generated code - do not edit
package measures

import (
	"encoding/xml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// insertion point for enums declaration

// insertion point for gongstructs declarations

// LengthType Named source named complex type "LengthType"
// This is a prototypical definition for a specific measure type defined as a vacuous extension (i.e. aliases) of gml:MeasureType. In this case, the content model supports the description of a length (or distance) quantity, with its units. The unit of measure referenced by uom shall be suitable for a length, such as metres or feet.
type LengthType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// ScaleType Named source named complex type "ScaleType"
type ScaleType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// TimeType Named source named complex type "TimeType"
type TimeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// GridLengthType Named source named complex type "GridLengthType"
type GridLengthType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// AreaType Named source named complex type "AreaType"
type AreaType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// VolumeType Named source named complex type "VolumeType"
type VolumeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// SpeedType Named source named complex type "SpeedType"
type SpeedType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// AngleType Named source named complex type "AngleType"
type AngleType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}
