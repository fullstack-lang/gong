package models

import (
	"encoding/xml"
	"strconv"
)

type SimpleType struct {
	Name string
	Annotated
	ElementWithNameAttribute
	Restriction *Restriction `xml:"restriction"`
	Union       *Union       `xml:"union"`
	List        *List        `xml:"list"`

	WithGoIdentifier
	ParticleAbstract
}

func (e *SimpleType) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	e.Order = Order
	e.Depth = Depth
	Order = Order + 1

	type Alias SimpleType
	aux := (*Alias)(e)

	Depth = Depth + 1
	err := d.DecodeElement(aux, &start)
	Depth = Depth - 1

	return err
}

func (simpleType *SimpleType) IsStringEnumerate() bool {

	restriction := simpleType.Restriction
	if restriction == nil {
		return false
	}
	if len(restriction.Enumerations) == 0 {
		return false
	}

	// remove namespace from type
	base := restriction.Base
	if NsPrefix(base) != "" {
		base = Name(base)
	}
	switch base {
	case "token", "string", "NCName", "NMTOKEN", "Name", "normalizedString", "language", "anyURI", "ID":
		return true
	default:
		return false
	}
}

func (simpleType *SimpleType) SetParentAndChildren(parent Particle) {
}

func (simpleType *SimpleType) generateGongEnum() string {
	enumName := simpleType.GoIdentifier
	if enumName == "" {
		enumName = "Enum_" + xsdNameToGoIdentifier(simpleType.NameXSD)
	}

	setOfEnumValueIdents := make(map[string]any)
	enumValuesCode := ""
	for _, enumeration := range simpleType.Restriction.Enumerations {
		var wid WithGoIdentifier
		computeGoIdentifier(enumName+"_"+xsdNameToGoIdentifier(enumeration.Value), &wid, setOfEnumValueIdents)
		enumValuesCode +=
			"\n\t" + wid.GoIdentifier +
				" " + enumName +
				" = " + strconv.Quote(enumeration.Value)
	}

	tmp := Replace3(
		Level1Code[Level1NamedEnumCode],

		"{{"+string(rune(Level2EnumComment))+"}}",
		"From xsd simple type with enumerate restriction \""+simpleType.NameXSD+"\"",

		"{{"+string(rune(Level2Enumname))+"}}",
		enumName,

		"{{"+string(rune(Level2EnumValues))+"}}",
		enumValuesCode,
	)
	return tmp
}
