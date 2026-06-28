package models

import "log"

// specific for RE-QIF
//
// many anonymous complex types
//
//	A_TYPE_1 Named source within outer element "TYPE"
//	type A_TYPE_1 struct {
func (schema *Schema) RenameTypeAnonymousComplexType() {
	for _, namedComplexType := range schema.ComplexTypes {
		for _, all := range namedComplexType.Alls {
			for _, element := range all.Elements {
				switch element.NameXSD {
				case "TYPE",
					"DEFINITION",
					"DEFAULT-VALUE",
					"SPEC-RELATIONS",
					"VALUES",
					"ATTRIBUTE-VALUE-XHTML":
					log.Println("Found an instance of anonymous complex type", element.NameXSD)
					if anonymousComplexType := element.ComplexType; anonymousComplexType != nil {
						for _, choice := range anonymousComplexType.Choices {
							for _, _e := range choice.Elements {
								_ = _e
								// We force here the name of the anonymous complex type
								// anonymousComplexType.Name = "Renamed_" + _e.NameXSD
								anonymousComplexType.Name = "A_" + _e.NameXSD
								log.Println("Renamed anonymous complex type to", anonymousComplexType.NameXSD)
							}
						}
					}
				default:

				}
			}

		}
	}

}
