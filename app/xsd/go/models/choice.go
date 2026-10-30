package models

import "encoding/xml"

type Choice struct {
	Name string
	Annotated

	ModelGroup

	IsDuplicatedInXSD bool

	OuterParticle Particle
}

func (c *Choice) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	c.Order = Order
	c.Depth = Depth
	Order = Order + 1

	type Alias Choice
	aux := (*Alias)(c)

	Depth = Depth + 1
	err := d.DecodeElement(aux, &start)
	Depth = Depth - 1

	return err
}
