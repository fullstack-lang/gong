package models

import "encoding/xml"

type All struct {
	Name string
	Annotated

	ModelGroup
}

func (a *All) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	a.Order = Order
	a.Depth = Depth
	Order = Order + 1

	type Alias All
	aux := (*Alias)(a)

	Depth = Depth + 1
	err := d.DecodeElement(aux, &start)
	Depth = Depth - 1

	return err
}
