package exporter

import (
	"encoding/base64"
	"encoding/xml"
	"fmt"
	"log"
	"time"

	load "github.com/fullstack-lang/gong/lib/load/go/models"

	"github.com/fullstack-lang/gong/dsm/reqif/go/models"
)

func (exporter *Exporter) ExportReqif(stager *models.Stager) {

	log.Println("Exporting the ReqIF file")

	if stager.GetRootREQIF() == nil {
		log.Fatal("No REQ_IF element to marshall")
	}

	// // Get the path to the system's temporary directory.
	// // os.TempDir() provides a cross-platform way to get this path.
	// // For example:
	// // - Linux/macOS: /tmp
	// // - Windows: C:\Users\YourUser\AppData\Local\Temp
	// tempDir := os.TempDir()
	// filePath := filepath.Join(tempDir, stager.pathToOutputReqifFile)

	// // Write the final byte slice to the specified file.
	// // os.WriteFile handles creating/truncating the file and writing the data.
	// // 0644 is a standard file permission.
	// err = os.WriteFile(filePath, outputData, 0644)
	// if err != nil {
	// 	fmt.Printf("Error writing to file %s: %v\n", filePath, err)
	// 	return
	// }

	// fmt.Printf("Successfully wrote REQ_IF data to %s\n", filePath)

	// get the root of the SPEC_TYPE slice
	var specTypes *models.A_SPEC_TYPES

	specTypesInstances := models.GetGongstrucsSorted[*models.A_SPEC_TYPES](stager.GetStage())

	if len(specTypesInstances) != 1 {
		log.Println("Problem, there should be one models.A_SPEC_TYPES instance")
		return
	}

	suffixNewObjectType := "-Text"
	specTypes = specTypesInstances[0]
	_ = specTypes

	if len(specTypes.SPEC_OBJECT_TYPE) != 1 {
		log.Println("Problem, there should be one SPEC_OBJECT_TYPE in DOORS legacy")
		return
	}
	legacyDoorsObjectType := specTypes.SPEC_OBJECT_TYPE[0]

	// add a spec object type
	newSpecObjectType := new(models.SPEC_OBJECT_TYPE)
	specTypes.SPEC_OBJECT_TYPE = append(specTypes.SPEC_OBJECT_TYPE, newSpecObjectType)
	newSpecObjectType.Name = legacyDoorsObjectType.Name + suffixNewObjectType
	newSpecObjectType.IDENTIFIER = legacyDoorsObjectType.IDENTIFIER + suffixNewObjectType

	// one needs to duplicate all the attribute definition for the type
	newSpecObjectType.SPEC_ATTRIBUTES = new(models.A_SPEC_ATTRIBUTES)
	for _, _attributeDefinitionXHTML := range legacyDoorsObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
		attributeDefinitionXHTML := new(models.ATTRIBUTE_DEFINITION_XHTML)
		attributeDefinitionXHTML.Name = _attributeDefinitionXHTML.Name + suffixNewObjectType
		attributeDefinitionXHTML.IDENTIFIER = _attributeDefinitionXHTML.IDENTIFIER + suffixNewObjectType
		attributeDefinitionXHTML.IS_EDITABLE = _attributeDefinitionXHTML.IS_EDITABLE
		attributeDefinitionXHTML.LONG_NAME = _attributeDefinitionXHTML.LONG_NAME + suffixNewObjectType
		attributeDefinitionXHTML.TYPE = _attributeDefinitionXHTML.TYPE

		newSpecObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML =
			append(newSpecObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML, attributeDefinitionXHTML)
	}

	// get the attribute definition that match "chapter name"
	var chapterNameAttributeDefinitionXHTML *models.ATTRIBUTE_DEFINITION_XHTML
	chapterNameAttributeName := "ReqIF.ChapterName"
	for _, _attributeDefinitionXHTML := range legacyDoorsObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
		if _attributeDefinitionXHTML.Name == chapterNameAttributeName {
			chapterNameAttributeDefinitionXHTML = _attributeDefinitionXHTML
		}
	}
	if chapterNameAttributeDefinitionXHTML == nil {
		log.Println("Found no XHTML attribute definition with name ", chapterNameAttributeName, ". Returning")
		return
	}

	var textAttributeDefinitionXHTML *models.ATTRIBUTE_DEFINITION_XHTML
	textAttributeName := "ReqIF.Text"
	for _, _attributeDefinitionXHTML := range legacyDoorsObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
		if _attributeDefinitionXHTML.Name == textAttributeName {
			textAttributeDefinitionXHTML = _attributeDefinitionXHTML
		}
	}
	if textAttributeDefinitionXHTML == nil {
		log.Println("Found no XHTML attribute definition with name ", textAttributeName, ". Returning")
		return
	}

	var PUIDAttributeDefinitionXHTML *models.ATTRIBUTE_DEFINITION_XHTML
	pUIDAttributeName := "IE PUID"
	for _, _attributeDefinitionXHTML := range legacyDoorsObjectType.SPEC_ATTRIBUTES.ATTRIBUTE_DEFINITION_XHTML {
		if _attributeDefinitionXHTML.Name == pUIDAttributeName {
			PUIDAttributeDefinitionXHTML = _attributeDefinitionXHTML
		}
	}
	if PUIDAttributeDefinitionXHTML == nil {
		log.Println("Found no XHTML attribute definition with name ", pUIDAttributeName, ". Returning")
		return
	}

	specObjectsInstances := models.GetGongstrucsSorted[*models.A_SPEC_OBJECTS](stager.GetStage())

	if len(specObjectsInstances) != 1 {
		log.Println("Problem, there should be one models.A_SPEC_OBJECTS instance")
		return
	}

	nbTextObjects := 0
	for _, specObject := range (specObjectsInstances[0]).SPEC_OBJECT {

		var isHeading bool
		var isRequirement bool
		var hasText bool

		// get the value to gather wether it has a chapter name attribute
		for _, attributeValueXHTML := range specObject.VALUES.ATTRIBUTE_VALUE_XHTML {
			if attributeValueXHTML.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF == chapterNameAttributeDefinitionXHTML.IDENTIFIER {
				isHeading = true
			}
			if attributeValueXHTML.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF == PUIDAttributeDefinitionXHTML.IDENTIFIER {
				isRequirement = true
			}
			if attributeValueXHTML.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF == textAttributeDefinitionXHTML.IDENTIFIER {
				hasText = true
			}

		}
		if isHeading || isRequirement {
			continue
		}

		nbTextObjects++
		specObject.TYPE.SPEC_OBJECT_TYPE_REF = newSpecObjectType.IDENTIFIER

		if !hasText {
			log.Println("Problem, item is not heading / requirement but has no text field")
		}

		// one needs to weed all its attributes but the XHTML attributes that have to be typed
		// with the attribute definition of the new Text object type
		clear(specObject.VALUES.ATTRIBUTE_VALUE_DATE)
		clear(specObject.VALUES.ATTRIBUTE_VALUE_INTEGER)
		clear(specObject.VALUES.ATTRIBUTE_VALUE_ENUMERATION)
		clear(specObject.VALUES.ATTRIBUTE_VALUE_REAL)
		clear(specObject.VALUES.ATTRIBUTE_VALUE_STRING)
		for _, attributeValueXHTML := range specObject.VALUES.ATTRIBUTE_VALUE_XHTML {
			attributeValueXHTML.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF =
				attributeValueXHTML.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF + suffixNewObjectType
		}

	}
	log.Println("Found ", nbTextObjects, "Text Objects")

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

	// Prepend the standard XML header to the marshalled data.
	// This makes it a valid XML file.
	outputData := []byte(xml.Header + string(xmlData))

	stager.GetLoadStage().Reset()

	fileToDownload := new(load.FileToDownload).Stage(stager.GetLoadStage())

	fileToDownload.Name = stager.GetPathToOutputReqifFile()
	fileToDownload.Base64EncodedContent = base64.StdEncoding.EncodeToString(outputData)

	stager.GetLoadStage().Commit()

	time.Sleep(1 * time.Second)
	stager.UpdateAndCommitLoadReqifStage()

	log.Println("Finished exporting the ReqIF file", stager.GetPathToOutputReqifFile())

}
