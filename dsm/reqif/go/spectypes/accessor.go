package spectypes

import (
	"log"

	"github.com/fullstack-lang/gong/dsm/reqif/go/models"
)

func GetSpecObjectTypeRendering(
	stage *models.Stage,
	spectObjectType *models.SPEC_OBJECT_TYPE,
) (
	specObjectTypeRendering *models.SPEC_OBJECT_TYPE_Rendering) {

	// SPEC_OBJECT_TYPE_Rendering instances Names are the identifiers of the
	// SPEC_OBJECT_TYPE instance. Since, by ReqIF design, those identifiers are unique,
	// we can use the gong map of those instances
	map_ := models.GongGetMap[*models.SPEC_OBJECT_TYPE_Rendering](stage)

	var ok bool
	if specObjectTypeRendering, ok = map_[spectObjectType.GetIdentifier()]; !ok {
		log.Panic("GetSpecObjectTypeRendering: Unknown specObjectType", spectObjectType.GetName())
	}

	return
}

func GetSpecAttributeDefinitionRendering[
	AttrDef models.AttributeDefinition, AttrDefRendering models.AttributeDefinitionRendering](
	stage *models.Stage,
	specAttributeDefinition AttrDef,
) (
	specAttributeDefinitionRendering AttrDefRendering) {

	// ATTRIBUTE_DEFINITION_Rendering instances Names are the identifiers of the
	// ATTRIBUTE_DEFINITION instance. Since, by ReqIF design, those identifiers are unique,
	// we can use the gong map of those instances
	map_ := models.GongGetMap[AttrDefRendering](stage)

	var ok bool
	if specAttributeDefinitionRendering, ok = map_[specAttributeDefinition.GetIdentifier()]; !ok {
		log.Panic("GetSpecAttributeDefinitionRendering: Unknown specAttributeDefinition", specAttributeDefinition.GetName())
	}

	return
}

func GetATTRIBUTE_DEFINITION[
	Attr models.Attribute,
	AttrDef models.AttributeDefinition](stager *models.Stager, attribute Attr) (attributeDefinition AttrDef) {

	ref := attribute.GetAttributeDefinitionRef()
	var ok bool

	switch any(*new(AttrDef)).(type) {
	case *models.ATTRIBUTE_DEFINITION_XHTML:
		var ad *models.ATTRIBUTE_DEFINITION_XHTML
		if ad, ok = stager.Map_id_ATTRIBUTE_DEFINITION_XHTML[ref]; ok {
			attributeDefinition = any(ad).(AttrDef)
		}
	case *models.ATTRIBUTE_DEFINITION_STRING:
		var ad *models.ATTRIBUTE_DEFINITION_STRING
		if ad, ok = stager.Map_id_ATTRIBUTE_DEFINITION_STRING[ref]; ok {
			attributeDefinition = any(ad).(AttrDef)
		}
	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		var ad *models.ATTRIBUTE_DEFINITION_BOOLEAN
		if ad, ok = stager.Map_id_ATTRIBUTE_DEFINITION_BOOLEAN[ref]; ok {
			attributeDefinition = any(ad).(AttrDef)
		}
	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		var ad *models.ATTRIBUTE_DEFINITION_INTEGER
		if ad, ok = stager.Map_id_ATTRIBUTE_DEFINITION_INTEGER[ref]; ok {
			attributeDefinition = any(ad).(AttrDef)
		}
	case *models.ATTRIBUTE_DEFINITION_DATE:
		var ad *models.ATTRIBUTE_DEFINITION_DATE
		if ad, ok = stager.Map_id_ATTRIBUTE_DEFINITION_DATE[ref]; ok {
			attributeDefinition = any(ad).(AttrDef)
		}
	case *models.ATTRIBUTE_DEFINITION_REAL:
		var ad *models.ATTRIBUTE_DEFINITION_REAL
		if ad, ok = stager.Map_id_ATTRIBUTE_DEFINITION_REAL[ref]; ok {
			attributeDefinition = any(ad).(AttrDef)
		}
	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		var ad *models.ATTRIBUTE_DEFINITION_ENUMERATION
		if ad, ok = stager.Map_id_ATTRIBUTE_DEFINITION_ENUMERATION[ref]; ok {
			attributeDefinition = any(ad).(AttrDef)
		}
	}

	if !ok {
		log.Panic("unknown ref ", ref)
	}

	return
}
