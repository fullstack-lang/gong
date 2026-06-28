package models

type Extension struct {
	Name string
	ModelGroup

	Base string `xml:"base,attr"`
	Ref  string `xml:"ref,attr"`

	Attributes      []*Attribute      `xml:"attribute"`
	AttributeGroups []*AttributeGroup `xml:"attributeGroup"`
}
