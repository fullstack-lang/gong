package models

type staticType string

func (st staticType) GoName() string {
	return string(st)
}

func (ct staticType) GoTypeName() string {
	return ct.GoName()
}

func (st staticType) Attributes() []Attribute {
	return []Attribute{}
}

func (st staticType) Elements() []Element {
	return []Element{}
}

func (st staticType) Schema() *Schema {
	return nil
}

func (staticType) ContainsText() bool {
	return true
}

func (st staticType) compile(*Schema, *Element) {
}

var staticTypes = map[string]staticType{
	"string":             "string",
	"language":           "string",
	"dateTime":           "string",
	"date":               "string",
	"base64Binary":       "string",
	"normalizedString":   "string",
	"token":              "string",
	"NCName":             "string",
	"NMTOKENS":           "string",
	"anySimpleType":      "string",
	"anyType":            "string",
	"int":                "int",
	"integer":            "int64",
	"long":               "int64",
	"negativeInteger":    "int64",
	"nonNegativeInteger": "uint64",
	"anyURI":             "string",
	"double":             "float64",
	"decimal":            "float64", // no: http://books.xmlschemata.org/relaxng/ch19-77057.html
	"float":              "float64",
	"boolean":            "bool",
	"ID":                 "string",
	"IDREF":              "string",
	"positiveInteger":    "uint64",
	"unsignedInt":        "uint64",
	"gYear":              "string",
	"gYearMonth":         "string",
	"gMonthDay":          "string",
	"gDay":               "string",
	"gMonth":             "string",
	"time":               "string",
	"unsignedShort":      "uint16",
	"unsignedByte":       "uint8",
	"short":              "int16",
	"byte":               "int8",
}

func StaticType(name string) staticType {
	typ, found := staticTypes[name]
	if found {
		return typ
	}
	panic("Type xsd:" + name + " not implemented")
}

func IsStaticType(name string) bool {
	_, found := staticTypes[name]
	return found
}
