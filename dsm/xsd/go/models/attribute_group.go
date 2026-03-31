package models

import "encoding/xml"

type AttributeGroup struct {
	Name string
	ElementWithNameAttribute
	Annotated
	WithGoIdentifier

	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`

	Ref string `xml:"ref,attr"`

	Attributes []*Attribute `xml:"attribute"`

	ParticleAbstract
}

func (ag *AttributeGroup) generateAttributes(
	agMap map[string]*AttributeGroup,
	stMap map[string]*SimpleType,
	setOfGoIdentifiers map[string]any,
	fields *string) {

	for _, referencedAg := range ag.AttributeGroups {

		if namedAg, ok := agMap[referencedAg.Ref]; ok {
			generateAttributes(namedAg.Attributes, stMap, setOfGoIdentifiers, fields)
			namedAg.generateAttributes(agMap, stMap, setOfGoIdentifiers, fields)
		}
	}

}

func (ag *AttributeGroup) SetParentAndChildren(parent Particle) {
	ag.Parent = parent
}

func (e *AttributeGroup) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {

	e.Order = Order
	e.Depth = Depth
	Order = Order + 1

	type Alias AttributeGroup
	aux := (*Alias)(e)

	Depth = Depth + 1
	err := d.DecodeElement(aux, &start)
	Depth = Depth - 1

	return err
}
