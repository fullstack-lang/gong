package models

import (
	"log"
	"strings"
)

func generateAttributes(
	attrs []*Attribute,
	stMap map[string]*SimpleType,
	setOfGoIdentifiers map[string]any,
	fields *string) {
	for _, attr := range attrs {

		// remove namespace from type
		if NsPrefix(attr.Type) != "" {
			attr.Type = Name(attr.Type)
		}

		goType := generateGoTypeFromType(attr.Type, stMap)

		var name string

		if goType == "" {

			prefix := "xlink:"
			if strings.HasPrefix(attr.Ref, prefix) {
				goType = "string"
				nameDec := strings.TrimPrefix(attr.Ref, prefix)
				name = xsdNameToGoIdentifier(nameDec)

				// overide the name of the attr
				attr.NameXSD = "http://www.w3.org/1999/xlink " + nameDec
			}

			prefix = "xml:"
			if strings.HasPrefix(attr.Ref, prefix) {
				goType = "string"
				nameDec := strings.TrimPrefix(attr.Ref, prefix)
				name = xsdNameToGoIdentifier(nameDec)

				// overide the name of the attr
				attr.NameXSD = "http://www.w3.org/XML/1998/namespace " + nameDec
			}
		} else {
			name = xsdNameToGoIdentifier(attr.Name)
		}

		if goType == "" {
			log.Fatalln("No resolution of attribute type", attr.Type)
		}

		switch name {
		case "Name":
			name = "NameXSD"
		}

		computeGoIdentifier(name, &attr.WithGoIdentifier, setOfGoIdentifiers)

		*fields += "\n\n\t// generated from attribute \"" + attr.NameXSD +
			"\n\t" + attr.GoIdentifier + " " + goType + " " + "`" + `xml:"` + attr.NameXSD + `,attr,omitempty"` + "`"
	}
}
