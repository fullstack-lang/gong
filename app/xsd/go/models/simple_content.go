package models

type SimpleContent struct {
	Name string

	Extension   *Extension   `xml:"extension"`
	Restriction *Restriction `xml:"restriction"`
}
