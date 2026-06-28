package exporter

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"log"
	"path/filepath"
	"strings"
	"time"

	load "github.com/fullstack-lang/gong/lib/load/go/models"

	"github.com/fullstack-lang/gong/app/reqif/go/models"
)

func (exporter *Exporter) ExportAnonymousReqif(stager *models.Stager) {

	log.Println("Exporting the ReqIF file")

	var rootReqif *models.REQ_IF
	if rootReqif = stager.GetRootREQIF(); rootReqif == nil {
		log.Fatal("No REQ_IF element to marshall")
	}
	_ = rootReqif

	rootReqif.THE_HEADER.REQ_IF_HEADER.COMMENT = "anonymous COMMENT"
	rootReqif.THE_HEADER.REQ_IF_HEADER.TITLE = "anonymous TITLE"

	// datatypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.DATATYPES
	// for idx, datatype := range datatypes.DATATYPE_DEFINITION_XHTML {
	// 	datatype.LONG_NAME = "XHTML datatype " + fmt.Sprintf("%5d", idx)
	// }
	// for idx, datatype := range datatypes.DATATYPE_DEFINITION_STRING {
	// 	datatype.LONG_NAME = "STRING datatype " + fmt.Sprintf("%5d", idx)
	// }
	// for idx, datatype := range datatypes.DATATYPE_DEFINITION_BOOLEAN {
	// 	datatype.LONG_NAME = "BOODATATYPE_DEFINITION_BOOLEAN datatype " + fmt.Sprintf("%5d", idx)
	// }
	// for idx, datatype := range datatypes.DATATYPE_DEFINITION_INTEGER {
	// 	datatype.LONG_NAME = "INTEGER datatype " + fmt.Sprintf("%5d", idx)
	// }
	// for idx, datatype := range datatypes.DATATYPE_DEFINITION_REAL {
	// 	datatype.LONG_NAME = "REAL datatype " + fmt.Sprintf("%5d", idx)
	// }
	// for idx, datatype := range datatypes.DATATYPE_DEFINITION_DATE {
	// 	datatype.LONG_NAME = "DATE datatype " + fmt.Sprintf("%5d", idx)
	// }
	// for idx, datatype := range datatypes.DATATYPE_DEFINITION_ENUMERATION {
	// 	datatype.LONG_NAME = "ENUMERATION datatype " + fmt.Sprintf("%5d", idx)

	// 	for idx2, enum := range datatype.SPECIFIED_VALUES.ENUM_VALUE {
	// 		enum.LONG_NAME = "ENUMERATION datatype " + fmt.Sprintf("%5d", idx) + ", ENUM " + fmt.Sprintf("%5d", idx2)
	// 	}
	// }

	// spectypes := stager.GetRootREQIF().CORE_CONTENT.REQ_IF_CONTENT.SPEC_TYPES
	// for idx2, type_ := range spectypes.SPEC_OBJECT_TYPE {
	// 	type_.LONG_NAME = "spec object type " + fmt.Sprintf("%5d", idx2)

	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
	// 		attributeDefinition.LONG_NAME = "XHTML attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
	// 		attributeDefinition.LONG_NAME = "STRING attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
	// 		attributeDefinition.LONG_NAME = "BOOattributeDefinition_DEFINITION_BOOLEAN attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
	// 		attributeDefinition.LONG_NAME = "INTEGER attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
	// 		attributeDefinition.LONG_NAME = "REAL attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
	// 		attributeDefinition.LONG_NAME = "DATE attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
	// 		attributeDefinition.LONG_NAME = "ENUMERATION attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// }

	// for idx2, type_ := range spectypes.SPEC_RELATION_TYPE {
	// 	type_.LONG_NAME = "spec relation type " + fmt.Sprintf("%5d", idx2)

	// 	if type_.SPEC_ATTRIBUTES == nil {
	// 		continue
	// 	}

	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
	// 		attributeDefinition.LONG_NAME = "XHTML attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
	// 		attributeDefinition.LONG_NAME = "STRING attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
	// 		attributeDefinition.LONG_NAME = "BOOattributeDefinition_DEFINITION_BOOLEAN attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
	// 		attributeDefinition.LONG_NAME = "INTEGER attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
	// 		attributeDefinition.LONG_NAME = "REAL attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
	// 		attributeDefinition.LONG_NAME = "DATE attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
	// 		attributeDefinition.LONG_NAME = "ENUMERATION attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// }
	// for idx2, type_ := range spectypes.SPECIFICATION_TYPE {
	// 	type_.LONG_NAME = "specification type " + fmt.Sprintf("%5d", idx2)

	// 	if type_.SPEC_ATTRIBUTES == nil {
	// 		continue
	// 	}

	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
	// 		attributeDefinition.LONG_NAME = "XHTML attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_STRING {
	// 		attributeDefinition.LONG_NAME = "STRING attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_BOOLEAN {
	// 		attributeDefinition.LONG_NAME = "BOOattributeDefinition_DEFINITION_BOOLEAN attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_INTEGER {
	// 		attributeDefinition.LONG_NAME = "INTEGER attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_REAL {
	// 		attributeDefinition.LONG_NAME = "REAL attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_DATE {
	// 		attributeDefinition.LONG_NAME = "DATE attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// 	for idx, attributeDefinition := range type_.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_ENUMERATION {
	// 		attributeDefinition.LONG_NAME = "ENUMERATION attributeDefinition " + fmt.Sprintf("%5d", idx)
	// 	}
	// }

	for _, specObject := range rootReqif.CORE_CONTENT.REQ_IF_CONTENT.SPEC_OBJECTS.SPEC_OBJECT {
		_ = specObject

		specObject.LONG_NAME = "SpecObject-" + specObject.IDENTIFIER

		if specObject.VALUES == nil {
			continue
		}

		for idx, attrValue := range specObject.VALUES.ATTRIBUTE_VALUE_STRING {
			_ = idx
			attrValue.THE_VALUE = specObject.IDENTIFIER
		}
		for idx, attrValue := range specObject.VALUES.ATTRIBUTE_VALUE_INTEGER {
			_ = idx
			attrValue.THE_VALUE = HashStringToInt(specObject.IDENTIFIER)
		}
		for idx, attrValue := range specObject.VALUES.ATTRIBUTE_VALUE_REAL {
			_ = idx
			attrValue.THE_VALUE = HashStringToFloat64(specObject.IDENTIFIER)
		}
		for _, attrValue := range specObject.VALUES.ATTRIBUTE_VALUE_DATE {
			attrValue.THE_VALUE = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC).Local().UTC().Format(time.DateOnly)
		}
		for idx, attrValue := range specObject.VALUES.ATTRIBUTE_VALUE_XHTML {
			_ = idx
			attrValue.THE_VALUE.EnclosedText = "<xhtml:div>" + specObject.IDENTIFIER + "</xhtml:div>"
		}
	}

	for _, specification := range rootReqif.CORE_CONTENT.REQ_IF_CONTENT.SPECIFICATIONS.SPECIFICATION {
		_ = specification

		specification.LONG_NAME = "Specification-" + specification.IDENTIFIER

		if specification.VALUES == nil {
			continue
		}

		for idx, attrValue := range specification.VALUES.ATTRIBUTE_VALUE_STRING {
			_ = idx
			attrValue.THE_VALUE = "String Value : " + specification.IDENTIFIER
		}
		for idx, attrValue := range specification.VALUES.ATTRIBUTE_VALUE_INTEGER {
			_ = idx
			attrValue.THE_VALUE = HashStringToInt(specification.IDENTIFIER)
		}
		for idx, attrValue := range specification.VALUES.ATTRIBUTE_VALUE_REAL {
			_ = idx
			attrValue.THE_VALUE = HashStringToFloat64(specification.IDENTIFIER)
		}
		for _, attrValue := range specification.VALUES.ATTRIBUTE_VALUE_DATE {
			attrValue.THE_VALUE = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC).Local().UTC().Format(time.DateOnly)
		}
		for idx, attrValue := range specification.VALUES.ATTRIBUTE_VALUE_XHTML {
			_ = idx
			attrValue.THE_VALUE.EnclosedText = "<xhtml:div>" + specification.IDENTIFIER + "</xhtml:div>"
		}
	}

	if rootReqif.CORE_CONTENT.REQ_IF_CONTENT.SPEC_RELATIONS != nil {
		for _, specRelation := range rootReqif.CORE_CONTENT.REQ_IF_CONTENT.SPEC_RELATIONS.SPEC_RELATION {
			_ = specRelation

			specRelation.LONG_NAME = specRelation.IDENTIFIER

			if specRelation.VALUES == nil {
				continue
			}

			for idx, attrValue := range specRelation.VALUES.ATTRIBUTE_VALUE_STRING {
				_ = idx
				attrValue.THE_VALUE = specRelation.IDENTIFIER
			}
			for idx, attrValue := range specRelation.VALUES.ATTRIBUTE_VALUE_INTEGER {
				_ = idx
				attrValue.THE_VALUE = HashStringToInt(specRelation.IDENTIFIER)
			}
			for idx, attrValue := range specRelation.VALUES.ATTRIBUTE_VALUE_REAL {
				_ = idx
				attrValue.THE_VALUE = HashStringToFloat64(specRelation.IDENTIFIER)
			}
			for _, attrValue := range specRelation.VALUES.ATTRIBUTE_VALUE_DATE {
				attrValue.THE_VALUE = time.Date(1970, time.January, 1, 0, 0, 0, 0, time.UTC).Local().UTC().Format(time.DateOnly)
			}
			for idx, attrValue := range specRelation.VALUES.ATTRIBUTE_VALUE_XHTML {
				_ = idx
				attrValue.THE_VALUE.EnclosedText = "<xhtml:div>" + specRelation.IDENTIFIER + "</xhtml:div>"
			}
		}
	}

	// for idx, a_specified_values := range models.GetGongstrucsSorted[*models.A_SPECIFIED_VALUES](stager.GetStage()) {

	// 	for idx2, a_specified_value := range a_specified_values.ENUM_VALUE {
	// 		a_specified_value.PROPERTIES.EMBEDDED_VALUE.OTHER_CONTENT = fmt.Sprintf("%5d", idx) + " - " + fmt.Sprintf("%5d", idx2)
	// 	}
	// }

	// parse all spec objects and if the spec object has a chapter name field, then
	// change the type of the object

	// Marshal the struct into an indented XML byte slice.
	// Using MarshalIndent makes the output file human-readable.
	// The "" means no line prefix, and "  " means use two spaces for indentation.
	xmlData, err := xml.MarshalIndent(stager.GetRootREQIF(), "", "  ")
	if err != nil {
		fmt.Println("Error marshalling XML:", err)
		return
	}

	// Replace the REQ_IF tag with the desired one
	xmlString := string(xmlData)
	xmlString = strings.Replace(xmlString,
		"<REQ_IF>",
		`<REQ-IF xmlns="http://www.omg.org/spec/ReqIF/20110401/reqif.xsd" xmlns:xhtml="http://www.w3.org/1999/xhtml">`, 1)

	xmlString = strings.Replace(xmlString,
		"</REQ_IF>",
		`</REQ-IF>`, 1)

	// // Mangle UUIDs by replacing them with their first 8 characters
	// re := regexp.MustCompile(`[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}`)
	// xmlString = re.ReplaceAllStringFunc(xmlString, func(s string) string {
	// 	return s[:8]
	// })

	// Prepend the standard XML header to the marshalled data.
	// This makes it a valid XML file.
	outputData := []byte(xml.Header + xmlString)

	stager.GetLoadStage().Reset()

	fileToDownload := new(load.FileToDownload).Stage(stager.GetLoadStage())

	filename := stager.GetPathToReqifFile()
	filename = filepath.Base(filename)

	if strings.HasSuffix(filename, ".reqifz") {
		filename = strings.TrimSuffix(filename, "z")
	}

	filename = strings.TrimSuffix(filename, ".reqif")
	filename += "-scrambled.reqif"

	fileToDownload.Name = filename
	fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString(outputData)

	stager.GetLoadStage().Commit()

	time.Sleep(1 * time.Second)
	stager.UpdateAndCommitLoadReqifStage()

	log.Println("Finished exporting the anonymous ReqIF file", filename)

}
