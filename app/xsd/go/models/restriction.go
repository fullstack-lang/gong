package models

type Restriction struct {
	Name string
	Annotated

	Base string `xml:"base,attr"`

	SimpleType *SimpleType `xml:"simpleType"`

	Enumerations  []*Enumeration `xml:"enumeration"`
	MinInclusive  *MinInclusive  `xml:"minInclusive"`
	MaxInclusive  *MaxInclusive  `xml:"maxInclusive"`
	MinExclusive  *MinExclusive  `xml:"minExclusive"`
	MaxExclusive  *MaxExclusive  `xml:"maxExclusive"`
	Pattern       *Pattern       `xml:"pattern"`
	WhiteSpace    *WhiteSpace    `xml:"whiteSpace"`
	MinLength     *MinLength     `xml:"minLength"`
	MaxLength     *MaxLength     `xml:"maxLength"`
	Length        *Length        `xml:"length"`
	TotalDigit    *TotalDigit    `xml:"totalDigits"`
	FractionDigit *FractionDigit `xml:"fractionDigits"`

	// For complexContent restriction
	ModelGroup
	Attributes      []*Attribute      `xml:"attribute"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`
}
