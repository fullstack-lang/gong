package models

import (
	"encoding/xml"
)

type SimpleType struct {
	Name string
	Annotated
	ElementWithNameAttribute
	Restriction *Restriction `xml:"restriction"`
	Union       *Union       `xml:"union"`

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

	// remove namespace from type
	base := restriction.Base
	if NsPrefix(base) != "" {
		base = Name(base)
	}
	if restriction.Base != "token" && restriction.Base != "string" {
		return false
	}

	for _, enumeration := range restriction.Enumerations {
		_ = enumeration
		// log.Println("Enum", enumeration.Value)
	}

	return true
}

func (simpleType *SimpleType) SetParentAndChildren(parent Particle) {
}

func (simpleType *SimpleType) generateGongEnum() string {

	enumValuesCode := ""
	for _, enumeration := range simpleType.Restriction.Enumerations {
		_ = enumeration
		enumValuesCode +=
			"\n\tEnum_" + xsdNameToGoIdentifier(simpleType.NameXSD) + "_" + xsdNameToGoIdentifier(enumeration.Value) +
				" " + "Enum_" + xsdNameToGoIdentifier(simpleType.NameXSD) +
				" = \"" + enumeration.Value + "\""
	}

	// log.Println("String Enumerate", st.Name)
	tmp := Replace3(
		Level1Code[Level1NamedEnumCode],

		"{{"+string(rune(Level2EnumComment))+"}}",
		"From xsd simple type with enumerate restriction \""+simpleType.NameXSD+"\"",

		"{{"+string(rune(Level2Enumname))+"}}",
		"Enum_"+xsdNameToGoIdentifier(simpleType.NameXSD),

		"{{"+string(rune(Level2EnumValues))+"}}",
		enumValuesCode,
	)
	return tmp
}
