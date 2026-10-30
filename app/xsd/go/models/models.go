package models

type ElementWithNameAttribute struct {
	NameXSD string `xml:"name,attr"`
}

type ElementWithTypeAttribute struct {
	Type string `xml:"type,attr"`
}

type ElementWithValueAttribute struct {
	Value string `xml:"value,attr"`
}

type Annotated struct {
	Annotation *Annotation `xml:"annotation"`
}

type Annotation struct {
	Name           string
	Documentations []*Documentation `xml:"documentation"`
}

type Union struct {
	Name string
	Annotated
	MemberTypes string        `xml:"memberTypes,attr"`
	SimpleTypes []*SimpleType `xml:"simpleType"`
}

type List struct {
	Name string
	Annotated
	ItemType   string      `xml:"itemType,attr"`
	SimpleType *SimpleType `xml:"simpleType"`
}

type XSDImport struct {
	Name string
	Annotated
	Namespace      string `xml:"namespace,attr"`
	SchemaLocation string `xml:"schemaLocation,attr"`
}

type XSDInclude struct {
	Name string
	Annotated
	SchemaLocation string `xml:"schemaLocation,attr"`
}

type Selector struct {
	Name  string
	XPath string `xml:"xpath,attr"`
}

type FieldXPath struct {
	Name  string
	XPath string `xml:"xpath,attr"`
}

type Key struct {
	Name string
	Annotated
	ElementWithNameAttribute
	Selector *Selector     `xml:"selector"`
	Fields   []*FieldXPath `xml:"field"`
}

type Unique struct {
	Name string
	Annotated
	ElementWithNameAttribute
	Selector *Selector     `xml:"selector"`
	Fields   []*FieldXPath `xml:"field"`
}

type KeyRef struct {
	Name string
	Annotated
	ElementWithNameAttribute
	Refer    string        `xml:"refer,attr"`
	Selector *Selector     `xml:"selector"`
	Fields   []*FieldXPath `xml:"field"`
}
