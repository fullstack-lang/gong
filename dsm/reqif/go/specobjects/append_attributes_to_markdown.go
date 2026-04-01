package specobjects

import (
	"fmt"
	"log"
	"regexp"
	"strings"

	"github.com/fullstack-lang/gong/dsm/reqif/go/models"
	m "github.com/fullstack-lang/gong/dsm/reqif/go/models"
	"github.com/fullstack-lang/gong/dsm/reqif/go/spectypes"
)

// Pre-compile regexes for efficiency.
var (
	// tagRegex matches any XML/XHTML-style tag.
	tagRegex = regexp.MustCompile(`<[^>]*>`)
	// spaceRegex matches one or more whitespace characters.
	spaceRegex = regexp.MustCompile(`\s+`)
)

// sanitizeForMarkdownTable escapes characters that can break a markdown table's structure.
func sanitizeForMarkdownTable(s string) string {
	// Replace pipe characters, which define columns
	s = strings.ReplaceAll(s, "|", "&#124;")
	// Replace newlines with a space to keep the row on a single line
	s = strings.ReplaceAll(s, "\n", " ")
	return s
}

// AppendAttributesToMarkdown populates a markdown string with a table of attributes from a SPEC_OBJECT.
func AppendAttributesToMarkdown(stager *m.Stager, specObject *m.SPEC_OBJECT, markdownContent *string) {
	var tableRows strings.Builder

	// Gather all attribute rows into the tableRows builder
	appendAttributeXHTMLRows(stager, specObject, &tableRows)
	appendAttributeStringRows(stager, specObject, &tableRows)
	appendAttributeBooleanRows(stager, specObject, &tableRows)
	appendAttributeIntegerRows(stager, specObject, &tableRows)
	appendAttributeRealRows(stager, specObject, &tableRows)
	appendAttributeDateRows(stager, specObject, &tableRows)
	appendAttributeEnumRows(stager, specObject, &tableRows)
	appendAttributeRelations(stager, specObject, &tableRows)

	// If any attributes were found, create the table.
	if tableRows.Len() > 0 {
		*markdownContent += "\n|  |  |\n"
		*markdownContent += "|---|---|\n"
		*markdownContent += tableRows.String()
	}
}

// appendAttributeXHTMLRows formats and appends XHTML attributes as markdown table rows.
func appendAttributeXHTMLRows(stager *m.Stager, specObject *m.SPEC_OBJECT, tableRows *strings.Builder) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_XHTML {
		attrDef := spectypes.GetATTRIBUTE_DEFINITION[*models.ATTRIBUTE_VALUE_XHTML, *models.ATTRIBUTE_DEFINITION_XHTML](stager, attribute)
		attributeDefintionRendering := GetAttributeDefinitionRendering(stager, attribute)

		if !*attributeDefintionRendering.GetShowInTablePtr() {
			continue
		}

		attributeDefinitionName := attrDef.LONG_NAME

		// Clean up XHTML content for markdown readability
		enclosedText := attribute.THE_VALUE.EnclosedText

		// 1. Remove all tags using a regular expression.
		textWithoutTags := tagRegex.ReplaceAllString(enclosedText, "")

		// 2. Normalize whitespace: replace newlines, tabs, and multiple spaces with a single space.
		normalizedText := spaceRegex.ReplaceAllString(textWithoutTags, " ")

		// 3. Trim leading/trailing space from the final string.
		finalText := strings.TrimSpace(normalizedText)

		tableRows.WriteString(fmt.Sprintf("| *%s*: | %s |\n",
			sanitizeForMarkdownTable(attributeDefinitionName),
			sanitizeForMarkdownTable(finalText)))
	}
}

// appendAttributeStringRows formats and appends String attributes as markdown table rows.
func appendAttributeStringRows(stager *m.Stager, specObject *m.SPEC_OBJECT, tableRows *strings.Builder) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_STRING {
		attrDef := spectypes.GetATTRIBUTE_DEFINITION[*models.ATTRIBUTE_VALUE_STRING, *models.ATTRIBUTE_DEFINITION_STRING](stager, attribute)
		attributeDefintionRendering := GetAttributeDefinitionRendering(stager, attribute)

		if !*attributeDefintionRendering.GetShowInTablePtr() {
			continue
		}

		attributeDefinitionName := attrDef.LONG_NAME
		tableRows.WriteString(fmt.Sprintf("| *%s*: | %s |\n",
			sanitizeForMarkdownTable(attributeDefinitionName),
			sanitizeForMarkdownTable(attribute.THE_VALUE)))
	}
}

// appendAttributeBooleanRows formats and appends Boolean attributes as markdown table rows.
func appendAttributeBooleanRows(stager *m.Stager, specObject *m.SPEC_OBJECT, tableRows *strings.Builder) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_BOOLEAN {
		attrDef := spectypes.GetATTRIBUTE_DEFINITION[*models.ATTRIBUTE_VALUE_BOOLEAN, *models.ATTRIBUTE_DEFINITION_BOOLEAN](stager, attribute)
		attributeDefintionRendering := GetAttributeDefinitionRendering(stager, attribute)

		if !*attributeDefintionRendering.GetShowInTablePtr() {
			continue
		}

		attributeDefinitionName := attrDef.LONG_NAME
		tableRows.WriteString(fmt.Sprintf("| *%s*: | %t |\n",
			sanitizeForMarkdownTable(attributeDefinitionName),
			attribute.THE_VALUE))
	}
}

// appendAttributeIntegerRows formats and appends Integer attributes as markdown table rows.
func appendAttributeIntegerRows(stager *m.Stager, specObject *m.SPEC_OBJECT, tableRows *strings.Builder) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_INTEGER {
		attrDef := spectypes.GetATTRIBUTE_DEFINITION[*models.ATTRIBUTE_VALUE_INTEGER, *models.ATTRIBUTE_DEFINITION_INTEGER](stager, attribute)
		attributeDefintionRendering := GetAttributeDefinitionRendering(stager, attribute)

		if !*attributeDefintionRendering.GetShowInTablePtr() {
			continue
		}

		attributeDefinitionName := attrDef.LONG_NAME
		tableRows.WriteString(fmt.Sprintf("| *%s*: | %d |\n",
			sanitizeForMarkdownTable(attributeDefinitionName),
			attribute.THE_VALUE))
	}
}

// appendAttributeDateRows formats and appends Date attributes as markdown table rows.
func appendAttributeDateRows(stager *m.Stager, specObject *m.SPEC_OBJECT, tableRows *strings.Builder) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_DATE {
		attrDef := spectypes.GetATTRIBUTE_DEFINITION[*models.ATTRIBUTE_VALUE_DATE, *models.ATTRIBUTE_DEFINITION_DATE](stager, attribute)
		attributeDefintionRendering := GetAttributeDefinitionRendering(stager, attribute)

		if !*attributeDefintionRendering.GetShowInTablePtr() {
			continue
		}

		attributeDefinitionName := attrDef.LONG_NAME
		tableRows.WriteString(fmt.Sprintf("| *%s*: | %s |\n",
			sanitizeForMarkdownTable(attributeDefinitionName),
			sanitizeForMarkdownTable(attribute.THE_VALUE)))
	}
}

// appendAttributeRealRows formats and appends Real attributes as markdown table rows.
func appendAttributeRealRows(stager *m.Stager, specObject *m.SPEC_OBJECT, tableRows *strings.Builder) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_REAL {
		attrDef := spectypes.GetATTRIBUTE_DEFINITION[*models.ATTRIBUTE_VALUE_REAL, *models.ATTRIBUTE_DEFINITION_REAL](stager, attribute)
		attributeDefintionRendering := GetAttributeDefinitionRendering(stager, attribute)

		if !*attributeDefintionRendering.GetShowInTablePtr() {
			continue
		}

		attributeDefinitionName := attrDef.LONG_NAME
		tableRows.WriteString(fmt.Sprintf("| *%s*: | %f |\n",
			sanitizeForMarkdownTable(attributeDefinitionName),
			attribute.THE_VALUE))
	}
}

// appendAttributeEnumRows formats and appends Enumeration attributes as markdown table rows.
func appendAttributeEnumRows(stager *m.Stager, specObject *m.SPEC_OBJECT, tableRows *strings.Builder) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
		attrDef := spectypes.GetATTRIBUTE_DEFINITION[*models.ATTRIBUTE_VALUE_ENUMERATION, *models.ATTRIBUTE_DEFINITION_ENUMERATION](stager, attribute)
		attributeDefintionRendering := GetAttributeDefinitionRendering(stager, attribute)

		if !*attributeDefintionRendering.GetShowInTablePtr() {
			continue
		}

		attributeDefinitionName := attrDef.LONG_NAME
		var enumValueString string
		if len(attribute.VALUES.ENUM_VALUE_REF) > 0 {
			valueIdentifier := attribute.VALUES.ENUM_VALUE_REF
			if enumValue, ok := stager.Map_id_ENUM_VALUE[valueIdentifier]; ok {
				enumValueString = enumValue.LONG_NAME
			} else {
				log.Panic("ENUM_VALUE_REF", valueIdentifier, "unknown ref in Map_id_ENUM_VALUE")
			}
		}

		tableRows.WriteString(fmt.Sprintf("| *%s*: | %s |\n",
			sanitizeForMarkdownTable(attributeDefinitionName),
			sanitizeForMarkdownTable(enumValueString)))
	}
}

// appendAttributeRelations appends rows to the markdown table for each of the spec object's relations,
// indicating the relation type and the related object.
func appendAttributeRelations(stager *m.Stager, specObject *m.SPEC_OBJECT, tableRows *strings.Builder) {

	specObjectType, ok := stager.Map_id_SPEC_OBJECT_TYPE[specObject.TYPE.SPEC_OBJECT_TYPE_REF]
	if !ok {
		log.Panic("specObject.TYPE.SPEC_OBJECT_TYPE_REF", specObject.TYPE.SPEC_OBJECT_TYPE_REF,
			"unknown ref")
	}

	specObjectTypeRendering := spectypes.GetSpecObjectTypeRendering(stager.GetStage(), specObjectType)
	if !specObjectTypeRendering.ShowRelations {
		return
	}

	for _, specRelation := range stager.Map_SPEC_OBJECT_relations_targets[specObject] {

		specRelationType, ok := stager.Map_id_SPEC_RELATION_TYPE[specRelation.TYPE.SPEC_RELATION_TYPE_REF]
		if !ok {
			log.Panic("specRelation.TYPE.SPEC_RELATION_TYPE_REF", specRelation.TYPE.SPEC_RELATION_TYPE_REF,
				"unknown relation type")
		}

		specObjectInRelation := stager.Map_id_SPEC_OBJECT[specRelation.SOURCE.SPEC_OBJECT_REF]

		var relationIdentification string
		AddIdentifierAndNameToTitle(stager, specObjectType, &relationIdentification, specObjectInRelation)
		relationIdentification += FillUpStringWithAttributes(stager, specObjectInRelation, Title)

		tableRows.WriteString(fmt.Sprintf("| *%s* >> : | %s |\n",
			sanitizeForMarkdownTable(specRelationType.GetName()),
			relationIdentification))
	}

	for _, specRelation := range stager.Map_SPEC_OBJECT_relations_sources[specObject] {

		specRelationType, ok := stager.Map_id_SPEC_RELATION_TYPE[specRelation.TYPE.SPEC_RELATION_TYPE_REF]
		if !ok {
			log.Panic("specRelation.TYPE.SPEC_RELATION_TYPE_REF", specRelation.TYPE.SPEC_RELATION_TYPE_REF,
				"unknown relation type")
		}

		specObjectInRelation := stager.Map_id_SPEC_OBJECT[specRelation.TARGET.SPEC_OBJECT_REF]
		var relationIdentification string
		AddIdentifierAndNameToTitle(stager, specObjectType, &relationIdentification, specObjectInRelation)
		relationIdentification += FillUpStringWithAttributes(stager, specObjectInRelation, Title)

		tableRows.WriteString(fmt.Sprintf("| *%s* << : | %s |\n",
			sanitizeForMarkdownTable(specRelationType.GetName()),
			relationIdentification))
	}
}
