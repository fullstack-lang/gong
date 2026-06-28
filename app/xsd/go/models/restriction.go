package models

type Restriction struct {
	Name string
	Annotated

	Base string `xml:"base,attr"`

	Enumerations []*Enumeration `xml:"enumeration"`
	MinInclusive *MinInclusive  `xml:"minInclusive"`
	MaxInclusive *MaxInclusive  `xml:"maxInclusive"`
	Pattern      *Pattern       `xml:"pattern"`
	WhiteSpace   *WhiteSpace    `xml:"whiteSpace"`
	MinLength    *MinLength     `xml:"minLength"`
	MaxLength    *MaxLength     `xml:"maxLength"`
	Length       *Length        `xml:"length"`
	TotalDigit   *TotalDigit    `xml:"totalDigits"`
}
