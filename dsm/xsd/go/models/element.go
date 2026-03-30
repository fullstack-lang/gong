package models

import (
	"encoding/xml"
	"log"
)

var Order = 1
var Depth int

type Element struct {
	Name string

	ParticleAbstract

	// analysis
	WithGoIdentifier

	Annotated
	ElementWithNameAttribute
	ElementWithTypeAttribute

	OccurrenceDefinitionAbstract

	Default  string `xml:"default,attr"`
	Fixed    string `xml:"fixed,attr"`
	Nillable string `xml:"nillable,attr"`
	Ref      string `xml:"ref,attr"`
	Abstract string `xml:"abstract,attr"`
	Form     string `xml:"form,attr"`
	Block    string `xml:"block,attr"`
	Final    string `xml:"final,attr"`

	SimpleType  *SimpleType  `xml:"simpleType"`
	ComplexType *ComplexType `xml:"complexType"`
	Groups      []*Group     `xml:"group"`

	OuterParticleOwnerAbstract

	IsDuplicatedInXSD bool
}

func (e *Element) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	e.Order = Order
	e.Depth = Depth
	Order = Order + 1

	type Alias Element
	aux := (*Alias)(e)

	Depth = Depth + 1
	err := d.DecodeElement(aux, &start)
	Depth = Depth - 1

	return err
}

func (e *Element) SetParentAndChildren(parent Particle) {

	if parent == e {
		log.Fatalln("setting parent as itself")
	}

	e.Parent = parent
	if e.ComplexType != nil {
		e.Children = append(e.Children, e.ComplexType)
		e.ComplexType.SetParentAndChildren(e)
	}
	if e.SimpleType != nil {
		e.Children = append(e.Children, e.ComplexType)
		e.SimpleType.SetParentAndChildren(e)
	}
}
