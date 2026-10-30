// generated code - do not edit
package basictypes

import (
	"encoding/xml"
)

// to avoid compilation error if no xml element is generated
var _ xml.Attr

// insertion point for enums declaration

// From xsd simple type with enumerate restriction "NilReasonEnumeration_Member_1"
type Enum_NilReasonEnumeration_Member_1 string

const (
	Enum_NilReasonEnumeration_Member_1_Inapplicable Enum_NilReasonEnumeration_Member_1 = "inapplicable"
	Enum_NilReasonEnumeration_Member_1_Missing Enum_NilReasonEnumeration_Member_1 = "missing"
	Enum_NilReasonEnumeration_Member_1_Template Enum_NilReasonEnumeration_Member_1 = "template"
	Enum_NilReasonEnumeration_Member_1_Unknown Enum_NilReasonEnumeration_Member_1 = "unknown"
	Enum_NilReasonEnumeration_Member_1_Withheld Enum_NilReasonEnumeration_Member_1 = "withheld"
)

// From xsd simple type with enumerate restriction "SignType"
type Enum_SignType string

const (
	Enum_SignType_N__ Enum_SignType = "-"
	Enum_SignType_N___1 Enum_SignType = "+"
)

// insertion point for gongstructs declarations

// CodeType Named source named complex type "CodeType"
// gml:CodeType is a generalized type to be used for a term, keyword or name.
// It adds a XML attribute codeSpace to a term, where the value of the codeSpace attribute (if present) shall indicate a dictionary, thesaurus, classification scheme, authority, or pattern for the term.
type CodeType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "codeSpace"
	CodeSpace string `xml:"codeSpace,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// CodeWithAuthorityType Named source named complex type "CodeWithAuthorityType"
// gml:CodeWithAuthorityType requires that the codeSpace attribute is provided in an instance.
type CodeWithAuthorityType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "codeSpace"
	CodeSpace string `xml:"codeSpace,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// MeasureType Named source named complex type "MeasureType"
// gml:MeasureType supports recording an amount encoded as a value of XML Schema double, together with a units of measure indicated by an attribute uom, short for "units Of measure". The value of the uom attribute identifies a reference system for the amount, usually a ratio or interval scale.
type MeasureType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "uom"
	Uom string `xml:"uom,attr,omitempty"`

	// simpleContent character data
	EnclosedText float64 `xml:",chardata"`
}

// CoordinatesType Named source named complex type "CoordinatesType"
// This type is deprecated for tuples with ordinate values that are numbers.
// CoordinatesType is a text string, intended to be used to record an array of tuples or coordinates. 
// While it is not possible to enforce the internal structure of the string through schema validation, some optional attributes have been provided in previous versions of GML to support a description of the internal structure. These attributes are deprecated. The attributes were intended to be used as follows:
// Decimal	symbol used for a decimal point (default="." a stop or period)
// cs        	symbol used to separate components within a tuple or coordinate string (default="," a comma)
// ts        	symbol used to separate tuples or coordinate strings (default=" " a space)
// Since it is based on the XML Schema string type, CoordinatesType may be used in the construction of tables of tuples or arrays of tuples, including ones that contain mixed text and numeric values.
type CoordinatesType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "decimal"
	Decimal string `xml:"decimal,attr,omitempty"`

	// generated from attribute "cs"
	Cs string `xml:"cs,attr,omitempty"`

	// generated from attribute "ts"
	Ts string `xml:"ts,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// CodeListType Named source named complex type "CodeListType"
// gml:CodeListType provides for lists of terms. The values in an instance element shall all be valid according to the rules of the dictionary, classification scheme, or authority identified by the value of its codeSpace attribute.
type CodeListType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "codeSpace"
	CodeSpace string `xml:"codeSpace,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// CodeOrNilReasonListType Named source named complex type "CodeOrNilReasonListType"
// gml:CodeOrNilReasonListType provides for lists of terms. The values in an instance element shall all be valid according to the rules of the dictionary, classification scheme, or authority identified by the value of its codeSpace attribute. An instance element may also include embedded values from NilReasonType. It is intended to be used in situations where a term or classification is expected, but the value may be absent for some reason.
type CodeOrNilReasonListType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "codeSpace"
	CodeSpace string `xml:"codeSpace,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// MeasureListType Named source named complex type "MeasureListType"
// gml:MeasureListType provides for a list of quantities.
type MeasureListType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "uom"
	Uom string `xml:"uom,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}

// MeasureOrNilReasonListType Named source named complex type "MeasureOrNilReasonListType"
// gml:MeasureOrNilReasonListType provides for a list of quantities. An instance element may also include embedded values from NilReasonType. It is intended to be used in situations where a value is expected, but the value may be absent for some reason.
type MeasureOrNilReasonListType struct {
	Name string `xml:"-"`

	// insertion point for fields

	// generated from attribute "uom"
	Uom string `xml:"uom,attr,omitempty"`

	// simpleContent character data
	EnclosedText string `xml:",chardata"`
}
