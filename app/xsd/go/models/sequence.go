package models

import "encoding/xml"

type Sequence struct {
	Name string
	Annotated

	ModelGroup

	OuterParticleOwnerAbstract
}

func (s *Sequence) UnmarshalXML(d *xml.Decoder, start xml.StartElement) error {
	s.Order = Order
	s.Depth = Depth
	Order = Order + 1

	type Alias Sequence
	aux := (*Alias)(s)

	Depth = Depth + 1
	err := d.DecodeElement(aux, &start)
	Depth = Depth - 1

	return err
}
