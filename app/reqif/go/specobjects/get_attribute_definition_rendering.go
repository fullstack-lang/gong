package specobjects

import (
	"github.com/fullstack-lang/gong/app/reqif/go/models"
	"github.com/fullstack-lang/gong/app/reqif/go/spectypes"
)

func GetAttributeDefinitionRendering[Attr models.Attribute](stager *models.Stager, attribute Attr) (
	attrDefRendering models.AttributeDefinitionRendering) {

	switch any(attribute).(type) {
	case *models.ATTRIBUTE_VALUE_STRING:
		attrDef := spectypes.GetATTRIBUTE_DEFINITION[Attr, *models.ATTRIBUTE_DEFINITION_STRING](stager, attribute)
		attrDefRendering =
			spectypes.GetSpecAttributeDefinitionRendering[
				*models.ATTRIBUTE_DEFINITION_STRING,
				*models.ATTRIBUTE_DEFINITION_STRING_Rendering](stager.GetStage(), attrDef)
	case *models.ATTRIBUTE_VALUE_XHTML:
		attrDef := spectypes.GetATTRIBUTE_DEFINITION[Attr, *models.ATTRIBUTE_DEFINITION_XHTML](stager, attribute)
		attrDefRendering =
			spectypes.GetSpecAttributeDefinitionRendering[
				*models.ATTRIBUTE_DEFINITION_XHTML,
				*models.ATTRIBUTE_DEFINITION_XHTML_Rendering](stager.GetStage(), attrDef)
	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		attrDef := spectypes.GetATTRIBUTE_DEFINITION[Attr, *models.ATTRIBUTE_DEFINITION_BOOLEAN](stager, attribute)
		attrDefRendering =
			spectypes.GetSpecAttributeDefinitionRendering[
				*models.ATTRIBUTE_DEFINITION_BOOLEAN,
				*models.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering](stager.GetStage(), attrDef)
	case *models.ATTRIBUTE_VALUE_DATE:
		attrDef := spectypes.GetATTRIBUTE_DEFINITION[Attr, *models.ATTRIBUTE_DEFINITION_DATE](stager, attribute)
		attrDefRendering =
			spectypes.GetSpecAttributeDefinitionRendering[
				*models.ATTRIBUTE_DEFINITION_DATE,
				*models.ATTRIBUTE_DEFINITION_DATE_Rendering](stager.GetStage(), attrDef)
	case *models.ATTRIBUTE_VALUE_REAL:
		attrDef := spectypes.GetATTRIBUTE_DEFINITION[Attr, *models.ATTRIBUTE_DEFINITION_REAL](stager, attribute)
		attrDefRendering =
			spectypes.GetSpecAttributeDefinitionRendering[
				*models.ATTRIBUTE_DEFINITION_REAL,
				*models.ATTRIBUTE_DEFINITION_REAL_Rendering](stager.GetStage(), attrDef)
	case *models.ATTRIBUTE_VALUE_INTEGER:
		attrDef := spectypes.GetATTRIBUTE_DEFINITION[Attr, *models.ATTRIBUTE_DEFINITION_INTEGER](stager, attribute)
		attrDefRendering =
			spectypes.GetSpecAttributeDefinitionRendering[
				*models.ATTRIBUTE_DEFINITION_INTEGER,
				*models.ATTRIBUTE_DEFINITION_INTEGER_Rendering](stager.GetStage(), attrDef)
	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		attrDef := spectypes.GetATTRIBUTE_DEFINITION[Attr, *models.ATTRIBUTE_DEFINITION_ENUMERATION](stager, attribute)
		attrDefRendering =
			spectypes.GetSpecAttributeDefinitionRendering[
				*models.ATTRIBUTE_DEFINITION_ENUMERATION,
				*models.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering](stager.GetStage(), attrDef)

	}

	return
}
