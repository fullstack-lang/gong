package generator

import (
	"fmt"
	"log"
	"os"
	"strings"

	m "github.com/fullstack-lang/gong/dsm/reqif/go/models"
)

const fileProlog = `// Generated code, do not edit
package models

import (
	"log"
	"strings"
)

`

const specObjectProlog = `
		case "%s":
			specObjectInstance := (&%s{
				Name: specObject.Name,
			}).Stage(stage)
			_ = specObjectInstance
			
			specObjectInstance.DESC = specObject.DESC
			specObjectInstance.IDENTIFIER = specObject.IDENTIFIER
			specObjectInstance.LAST_CHANGE = specObject.LAST_CHANGE
			specObjectInstance.LONG_NAME = specObject.LONG_NAME
			`

const stagerFunctionStart = `

func stageReqifObjects(stager *Stager, stage *Stage) {
	objects := stager.reqifStager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_OBJECTS

	for _, specObject := range objects.SPEC_OBJECT {

		specObjectType, ok := stager.reqifStager.Map_id_specobjectTypes[specObject.TYPE.SPEC_OBJECT_TYPE_REF]
		if !ok {
			log.Panic("specObject.TYPE.SPEC_OBJECT_TYPE_REF", specObject.TYPE.SPEC_OBJECT_TYPE_REF,
				"unknown object type")
		}

		switch specObjectType.Name {
`

const attributeXHTMLStart = `
			for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_XHTML {
				// provide the type
				var attributeDefinition string
				if datatype, ok := stager.reqifStager.Map_id_attributeDefinitionXHTML[attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF]; ok {
					attributeDefinition = datatype.LONG_NAME
				} else {
					log.Panic("ATTRIBUTE_DEFINITION_XHTML_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF,
						"unknown ref")
				}
				_ = attributeDefinition

				enclosedText := attribute.THE_VALUE.EnclosedText

				enclosedText = strings.ReplaceAll(enclosedText, "<reqif-xhtml:div>", " ")
				enclosedText = strings.ReplaceAll(enclosedText, "</reqif-xhtml:div>", "\n")
				enclosedText = strings.ReplaceAll(enclosedText, "<reqif-xhtml:br >", "-")
				enclosedText = strings.ReplaceAll(enclosedText, "</reqif-xhtml:br >", "\n")
				enclosedText = strings.ReplaceAll(enclosedText, "<reqif-xhtml:br />", "\n")

				switch attributeDefinition {
`

const attributeStringStart = `
			for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_STRING {
				// provide the type
				var attributeDefinition string
				if datatype, ok := stager.reqifStager.Map_id_attributeDefinitionString[attribute.DEFINITION.ATTRIBUTE_DEFINITION_STRING_REF]; ok {
					attributeDefinition = datatype.LONG_NAME
				} else {
					log.Panic("ATTRIBUTE_DEFINITION_STRING_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_STRING_REF,
						"unknown ref")
				}
				_ = attributeDefinition

				switch attributeDefinition {
`

const attributeEnumStart = `
			for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
				// provide the type
				var attributeDefinition string
				if datatype, ok := stager.reqifStager.Map_id_attributeDefinitionEnum[attribute.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF]; ok {
					attributeDefinition = datatype.LONG_NAME
				} else {
					log.Panic("ATTRIBUTE_DEFINITION_STRING_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF,
						"unknown ref")
				}
				_ = attributeDefinition

				switch attributeDefinition {
`

const attributeBooleanStart = `
			for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
				// provide the type
				var attributeDefinition string
				if datatype, ok := stager.reqifStager.Map_id_attributeDefinitionBoolean[attribute.DEFINITION.ATTRIBUTE_DEFINITION_BOOLEAN_REF]; ok {
					attributeDefinition = datatype.LONG_NAME
				} else {
					log.Panic("ATTRIBUTE_DEFINITION_BOOLEAN_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_BOOLEAN_REF,
						"unknown ref")
				}
				_ = attributeDefinition

				switch attributeDefinition {
`

const enumTemplate = `
				case "%s":
					var __val__ %s
					value, err1 := stager.reqifStager.Map_id_enumValues[attribute.VALUES.ENUM_VALUE_REF]
					if err1 != ok {
						log.Println(attribute.GetName(), " is not a good enum ref")
					}
					err := __val__.FromString(value.Name)
					if err != nil {
						log.Println(attribute.GetName(), " is not a good enum")
					}
					specObjectInstance.%s = __val__
`

const attributeDateStart = `
			for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_DATE {
				// provide the type
				var attributeDefinition string
				if datatype, ok := stager.reqifStager.Map_id_attributeDefinitionDate[attribute.DEFINITION.ATTRIBUTE_DEFINITION_DATE_REF]; ok {
					attributeDefinition = datatype.LONG_NAME
				} else {
					log.Panic("ATTRIBUTE_DEFINITION_DATE_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_DATE_REF,
						"unknown ref")
				}
				_ = attributeDefinition

				switch attributeDefinition {
`

type ModelGenerator struct {
	PathToGoModelFile string
}

// GenerateModels implements m.ModelGeneratorInterface.
func (modelGenerator *ModelGenerator) GenerateModels(stager *m.Stager) {
	var sb strings.Builder

	sb.WriteString(fileProlog)

	// generates the enums
	datatypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.DATATYPES
	for _, datatype := range datatypes.DATATYPE_DEFINITION_ENUMERATION {

		sb.WriteString(fmt.Sprintf(`type %s string`, GenerateGoIdentifier(datatype.Name)))
		sb.WriteString("\n")
		sb.WriteString("\n")

		sb.WriteString("const (\n")

		for _, enum := range datatype.SPECIFIED_VALUES.ENUM_VALUE {

			sb.WriteString(fmt.Sprintf("\t%s_%s %s = \"%s\"\n",
				GenerateGoIdentifier(datatype.Name),
				GenerateGoIdentifier(enum.LONG_NAME),
				GenerateGoIdentifier(datatype.Name),
				enum.LONG_NAME))

		}

		sb.WriteString("\n)\n")
	}

	// generates the struct
	spectypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_TYPES

	for _, specObjectType := range spectypes.SPEC_OBJECT_TYPE {

		sb.WriteString(fmt.Sprintf("type %s struct {", GenerateGoIdentifier(specObjectType.Name)))
		sb.WriteString("\tName string\n")
		sb.WriteString("\n")
		sb.WriteString("\nDESC string\n")
		sb.WriteString("\n")
		sb.WriteString("\tIDENTIFIER string\n")
		sb.WriteString("\n")
		sb.WriteString("\nLAST_CHANGE string\n")
		sb.WriteString("\n")
		sb.WriteString("\nLONG_NAME string\n")
		sb.WriteString("\n")

		for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
			sb.WriteString(fmt.Sprintf("\t%s string\n", GenerateGoIdentifier(attribute.Name)))
		}

		for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {

			var attributeType string
			if datatype, ok := stager.Map_id_DATATYPE_DEFINITION_ENUMERATION[attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF]; ok {
				attributeType = datatype.LONG_NAME
			} else {
				log.Panic("attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF", attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF,
					"unknown ref")
			}

			sb.WriteString(fmt.Sprintf("\t%s %s\n",
				GenerateGoIdentifier(attribute.Name),
				GenerateGoIdentifier(attributeType),
			))
		}

		for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
			sb.WriteString(fmt.Sprintf("\t%s bool\n", GenerateGoIdentifier(attribute.Name)))
		}
		for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
			sb.WriteString(fmt.Sprintf("\t%s string\n", GenerateGoIdentifier(attribute.Name)))
		}
		for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
			sb.WriteString(fmt.Sprintf("\t%s int\n", GenerateGoIdentifier(attribute.Name)))
		}
		for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
			sb.WriteString(fmt.Sprintf("\t%s float64\n", GenerateGoIdentifier(attribute.Name)))
		}
		for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
			sb.WriteString(fmt.Sprintf("\t%s string\n", GenerateGoIdentifier(attribute.Name)))
		}

		sb.WriteString("\n}\n")
	}

	// generates the stager function
	sb.WriteString(stagerFunctionStart)

	for _, specObjectType := range spectypes.SPEC_OBJECT_TYPE {
		sb.WriteString(fmt.Sprintf(specObjectProlog,
			specObjectType.GetName(),
			GenerateGoIdentifier(specObjectType.GetName()),
		))

		{
			// attributes XHTML
			sb.WriteString(attributeXHTMLStart)

			for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
				sb.WriteString(fmt.Sprintf("\t\t\t\tcase \"%s\":\n", attribute.Name))
				sb.WriteString(fmt.Sprintf("\t\t\t\t\tspecObjectInstance.%s = enclosedText\n", GenerateGoIdentifier(attribute.Name)))
			}

			sb.WriteString("\t\t\t\tdefault:\n")
			sb.WriteString("\t\t\t\t\tlog.Println(\"Unkown field name (attribute definition)\", attributeDefinition)\n")
			sb.WriteString("\t\t\t\t} // end of the switch on the attribute id\n")

			sb.WriteString("\t\t\t} // end of the loop on the xhtml attribute parsing\n")
		}
		{
			// attributes STRING
			sb.WriteString(attributeStringStart)

			for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
				sb.WriteString(fmt.Sprintf("\t\t\t\tcase \"%s\":\n", attribute.Name))
				sb.WriteString(fmt.Sprintf("\t\t\t\t\tspecObjectInstance.%s = attribute.THE_VALUE\n", GenerateGoIdentifier(attribute.Name)))
			}

			sb.WriteString("\t\t\t\tdefault:\n")
			sb.WriteString("\t\t\t\t\tlog.Println(\"Unkown field name (attribute definition)\", attributeDefinition)\n")
			sb.WriteString("\t\t\t\t} // end of the switch on the attribute id\n")

			sb.WriteString("\t\t\t} // end of the loop on the string attribute parsing\n")
		}
		{
			// attributes Enumeration
			sb.WriteString(attributeEnumStart)

			for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {

				var attributeType string
				if datatype, ok := stager.Map_id_DATATYPE_DEFINITION_ENUMERATION[attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF]; ok {
					attributeType = datatype.LONG_NAME
				} else {
					log.Panic("attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF", attribute.TYPE.DATATYPE_DEFINITION_ENUMERATION_REF,
						"unknown ref")
				}

				sb.WriteString(fmt.Sprintf(enumTemplate,
					attribute.Name,
					GenerateGoIdentifier(attributeType),
					GenerateGoIdentifier(attribute.Name)))
			}

			sb.WriteString("\t\t\t\tdefault:\n")
			sb.WriteString("\t\t\t\t\tlog.Println(\"Unkown field name (attribute definition)\", attributeDefinition)\n")
			sb.WriteString("\t\t\t\t} // end of the switch on the attribute id\n")

			sb.WriteString("\t\t\t} // end of the loop on the enumeration attribute parsing\n")
		}
		{
			// attributes BOOLEAN
			sb.WriteString(attributeBooleanStart)

			for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
				sb.WriteString(fmt.Sprintf("\t\t\t\tcase \"%s\":\n", attribute.Name))
				sb.WriteString(fmt.Sprintf("\t\t\t\t\tspecObjectInstance.%s = attribute.THE_VALUE\n", GenerateGoIdentifier(attribute.Name)))
			}

			sb.WriteString("\t\t\t\tdefault:\n")
			sb.WriteString("\t\t\t\t\tlog.Println(\"Unkown field name (attribute definition)\", attributeDefinition)\n")
			sb.WriteString("\t\t\t\t} // end of the switch on the attribute id\n")

			sb.WriteString("\t\t\t} // end of the loop on the boolean attribute parsing\n")
		}
		{
			// attributes DATE
			sb.WriteString(attributeDateStart)

			for _, attribute := range specObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
				sb.WriteString(fmt.Sprintf("\t\t\t\tcase \"%s\":\n", attribute.Name))
				sb.WriteString(fmt.Sprintf("\t\t\t\t\tspecObjectInstance.%s = attribute.THE_VALUE\n", GenerateGoIdentifier(attribute.Name)))
			}

			sb.WriteString("\t\t\t\tdefault:\n")
			sb.WriteString("\t\t\t\t\tlog.Println(\"Unkown field name (attribute definition)\", attributeDefinition)\n")
			sb.WriteString("\t\t\t\t} // end of the switch on the attribute id\n")

			sb.WriteString("\t\t\t} // end of the loop on the date attribute parsing\n")
		}
	}

	sb.WriteString("\t\t} // end of the switch on spec object types attributes xhtml\n")
	sb.WriteString("\t}// end of the loop on spec object types\n")
	sb.WriteString("}\n")

	result := sb.String()
	err := os.WriteFile(modelGenerator.PathToGoModelFile, []byte(result), 0644)
	if err != nil {
		log.Panicln("Generate go models", err.Error())
	}
}
