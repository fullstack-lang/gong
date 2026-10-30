package models

import (
	"encoding/xml"
)

var SchemaSingloton Schema

type Schema struct {
	Name string
	Xs   string `xml:"xs,attr"`

	TargetNamespace      string `xml:"targetNamespace,attr"`
	ElementFormDefault   string `xml:"elementFormDefault,attr"`
	AttributeFormDefault string `xml:"attributeFormDefault,attr"`
	Version              string `xml:"version,attr"`

	// Xmlns represents namespaces like xmlns:xsd="http://www.w3.org/2001/XMLSchema"
	// they are unmarshall via a custom unmarshaller
	Xmlns Xmlns `xml:"-"`

	// Multi-package resolution metadata
	FilePath      string `xml:"-"`
	PackageName   string `xml:"-"`
	PackageRelDir string `xml:"-"`
	IsRootSchema  bool   `xml:"-"`

	Annotated
	Imports         []*XSDImport      `xml:"import"`
	Includes        []*XSDInclude     `xml:"include"`
	Elements        []*Element        `xml:"element"`
	Attributes      []*Attribute      `xml:"attribute"`
	SimpleTypes     []*SimpleType     `xml:"simpleType"`
	ComplexTypes    []*ComplexType    `xml:"complexType"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`
	Groups          []*Group          `xml:"group"`

	Keys    []*Key    `xml:"key"`
	Uniques []*Unique `xml:"unique"`
	KeyRefs []*KeyRef `xml:"keyref"`

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
	for _, p := range schema.Attributes {
		p.SetParentAndChildren(schema)
	}
	for _, p := range schema.SimpleTypes {
		p.SetParentAndChildren(schema)
	}
	for _, p := range schema.AttributeGroups {
		p.SetParentAndChildren(schema)
	}
}
