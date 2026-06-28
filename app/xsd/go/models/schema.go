package models

import (
	"encoding/xml"
)

var SchemaSingloton Schema

type Schema struct {
	Name string
	Xs   string `xml:"xs,attr"`

	// Xmlns represents namespaces like xmlns:xsd="http://www.w3.org/2001/XMLSchema"
	// they are unmarshall via a custom unmarshaller
	Xmlns Xmlns `xml:"-"`

	Annotated
	Elements        []*Element        `xml:"element"`
	SimpleTypes     []*SimpleType     `xml:"simpleType"`
	ComplexTypes    []*ComplexType    `xml:"complexType"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`
	Groups          []*Group          `xml:"group"`

	ParticleAbstract
}

func (sch *Schema) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	sch.Xmlns = parseXmlns(start)

	type s Schema
	ss := (*s)(sch)
	return d.DecodeElement(ss, &start)
}

func (schema *Schema) SetParentAndChildren(parent Particle) {
	for _, p := range schema.ComplexTypes {
		p.SetParentAndChildren(schema)
	}
	for _, p := range schema.Groups {
		p.SetParentAndChildren(schema)
	}
	for _, p := range schema.Elements {
		p.SetParentAndChildren(schema)
	}
}
