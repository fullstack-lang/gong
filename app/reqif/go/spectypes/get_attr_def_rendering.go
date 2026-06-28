package spectypes

import (
	m "github.com/fullstack-lang/gong/app/reqif/go/models"
)

func GetAttrDefRendering[AttrDef m.AttributeDefinition](
	stager *m.Stager,
	attributeDefinition AttrDef,
) (
	attrDefRendering m.AttributeDefinitionRendering) {

	switch any(attributeDefinition).(type) {
	case *m.ATTRIBUTE_DEFINITION_XHTML:
		attrDefRendering = GetSpecAttributeDefinitionRendering[
			AttrDef, *m.ATTRIBUTE_DEFINITION_XHTML_Rendering](stager.GetStage(),
			attributeDefinition)
	case *m.ATTRIBUTE_DEFINITION_STRING:
		attrDefRendering = GetSpecAttributeDefinitionRendering[
			AttrDef, *m.ATTRIBUTE_DEFINITION_STRING_Rendering](stager.GetStage(),
			attributeDefinition)
	case *m.ATTRIBUTE_DEFINITION_BOOLEAN:
		attrDefRendering = GetSpecAttributeDefinitionRendering[
			AttrDef, *m.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering](stager.GetStage(),
			attributeDefinition)
	case *m.ATTRIBUTE_DEFINITION_INTEGER:
		attrDefRendering = GetSpecAttributeDefinitionRendering[
			AttrDef, *m.ATTRIBUTE_DEFINITION_INTEGER_Rendering](stager.GetStage(),
			attributeDefinition)
	case *m.ATTRIBUTE_DEFINITION_DATE:
		attrDefRendering = GetSpecAttributeDefinitionRendering[
			AttrDef, *m.ATTRIBUTE_DEFINITION_DATE_Rendering](stager.GetStage(),
			attributeDefinition)
	case *m.ATTRIBUTE_DEFINITION_REAL:
		attrDefRendering = GetSpecAttributeDefinitionRendering[
			AttrDef, *m.ATTRIBUTE_DEFINITION_REAL_Rendering](stager.GetStage(),
			attributeDefinition)
	case *m.ATTRIBUTE_DEFINITION_ENUMERATION:
		attrDefRendering = GetSpecAttributeDefinitionRendering[
			AttrDef, *m.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering](stager.GetStage(),
			attributeDefinition)
	}

	return
}
