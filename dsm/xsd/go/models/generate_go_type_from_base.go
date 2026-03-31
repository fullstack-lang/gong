package models

func generateGoTypeFromType(type_ string, stMap map[string]*SimpleType) (goType string) {

	switch type_ {
	// String types
	case
		"dateTime",
		"string", "normalizedString", "token", "language", "Name", "NCName",
		"NMTOKEN", "NMTOKENS", "ID", "IDREF", "IDREFS", "ENTITY", "ENTITIES",
		"QName", "anyURI", "NOTATION", "decimal":
		goType = "string"

	// Numeric types
	case "integer", "nonPositiveInteger", "negativeInteger", "long",
		"int", "short", "byte", "nonNegativeInteger", "unsignedLong", "unsignedInt",
		"unsignedShort", "unsignedByte", "positiveInteger":
		goType = "int"

		// Numeric types
	case "double", "float":
		goType = "float64"

	// Boolean type
	case "boolean":
		goType = "bool"

	}

	// a base can refer a simple type
	if goType == "" {

		if st, ok := stMap[type_]; ok {

			if st.IsStringEnumerate() {
				// log.Println("string enum")
				return "Enum_" + xsdNameToGoIdentifier(st.NameXSD)
			}
			if st.Restriction != nil {

				// remove namespace from type
				if NsPrefix(st.Restriction.Base) != "" {
					st.Restriction.Base = Name(st.Restriction.Base)
				}

				return generateGoTypeFromType(st.Restriction.Base, stMap)
			} else {
				//
				// the union of type can be a go type int, float64, string or boolean
				// on simplify to string
				return "string"
				// types := strings.Split(st.Union.MemberTypes, " ")
				// return generateGoTypeFromSimpleType(types[0], stMap)
			}
		}
		//
		// type="empty": This indicates that the element is expected to be empty
		// (i.e., it does not have any child elements or text content).
		if type_ == "empty" {
			return "string"
		}
	}
	return
}
