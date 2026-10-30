package models

type ComplexContent struct {
	Name string
	Annotated
	Mixed       string       `xml:"mixed,attr"`
	Extension   *Extension   `xml:"extension"`
	Restriction *Restriction `xml:"restriction"`
}
