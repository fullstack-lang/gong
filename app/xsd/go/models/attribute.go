package models

import "encoding/xml"

type Attribute struct {
	Name string
	ElementWithNameAttribute
	ElementWithTypeAttribute
	Annotated

	WithGoIdentifier

	Default         string      `xml:"default,attr"`
	Use             string      `xml:"use,attr"`
	Form            string      `xml:"form,attr"`
	Fixed           string      `xml:"fixed,attr"`
	Ref             string      `xml:"ref,attr"`
	TargetNamespace string      `xml:"targetNamespace,attr"`
	SimpleType      string      `xml:"simpleType,attr"`
	InnerSimpleType *SimpleType `xml:"simpleType"`
	IDXSD           string      `xml:"id,attr"`

	ParticleAbstract
}

func (a *Attribute) SetParentAndChildren(parent Particle) {
	a.Parent = parent
	if a.InnerSimpleType != nil {
		a.Children = append(a.Children, a.InnerSimpleType)
		a.InnerSimpleType.SetParentAndChildren(a)
	}
}

func (a *Attribute) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	a.Order = Order
	a.Depth = Depth
	Order = Order + 1

	type Alias Attribute
	aux := (*Alias)(a)

	Depth = Depth + 1
	err := d.DecodeElement(aux, &start)
	Depth = Depth - 1

	return err
}
