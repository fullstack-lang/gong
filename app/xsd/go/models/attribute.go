package models

type Attribute struct {
	Name string
	ElementWithNameAttribute
	ElementWithTypeAttribute
	Annotated

	WithGoIdentifier

	Default         string `xml:"default,attr"`
	Use             string `xml:"use,attr"`
	Form            string `xml:"form,attr"`
	Fixed           string `xml:"fixed,attr"`
	Ref             string `xml:"ref,attr"`
	TargetNamespace string `xml:"targetNamespace,attr"`
	SimpleType      string `xml:"simpleType,attr"`
	IDXSD           string `xml:"id,attr"`
}
