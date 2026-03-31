package xsd

type SchemaVersion1 struct {
	Redefinables []*ComplexType `xml:"complexType"`
}

type SchemaVersion2 struct {
	Redefinables []*Redefinable `xml:",inline"`
}

type Redefinable struct {
	ComplexType *ComplexType `xml:"complexType"`
}
