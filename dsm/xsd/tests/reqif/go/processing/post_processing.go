package processing

import "github.com/fullstack-lang/gong/dsm/xsd/tests/reqif/go/models"

func PostProcessing(stage *models.Stage) {
	for x := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_ENUMERATION](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_XHTML](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_ENUMERATION](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_DEFINITION_XHTML](stage) {
		x.Name = x.LONG_NAME
	}

	// Attribute Value

	// for x := range *models.GetGongstructInstancesSet[models.ATTRIBUTE_VALUE_ENUMERATION](stage) {
	// 	x.Name = x.DEFINITION
	// }

	// anonymous without ATTRIBUTE DEF

	for x := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](stage) {
		x.Name = x.ATTRIBUTE_DEFINITION_ENUMERATION_REF
	}

	for x := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_DEFINITION_XHTML_REF](stage) {
		x.Name = x.ATTRIBUTE_DEFINITION_XHTML_REF
	}

	// for x := range *models.GetGongstructInstancesSet[models.A_ATTRIBUTE_VALUE_XHTML_1](stage) {
	// 	if x.
	// 	x.Name = x.DATATYPE_DEFINITION_ENUMERATION_REF
	// }

	for x := range *models.GetGongstructInstancesSet[models.A_DATATYPE_DEFINITION_ENUMERATION_REF](stage) {
		x.Name = x.DATATYPE_DEFINITION_ENUMERATION_REF
	}

	for x := range *models.GetGongstructInstancesSet[models.A_DATATYPE_DEFINITION_XHTML_REF](stage) {
		x.Name = x.DATATYPE_DEFINITION_XHTML_REF
	}

	for x := range *models.GetGongstructInstancesSet[models.A_ENUM_VALUE_REF](stage) {
		x.Name = x.ENUM_VALUE_REF
	}

	for x := range *models.GetGongstructInstancesSet[models.A_OBJECT](stage) {
		x.Name = x.SPEC_OBJECT_REF
	}

	for x := range *models.GetGongstructInstancesSet[models.A_SPECIFICATION_TYPE_REF](stage) {
		x.Name = x.SPECIFICATION_TYPE_REF
	}

	for x := range *models.GetGongstructInstancesSet[models.A_SPEC_OBJECT_TYPE_REF](stage) {
		x.Name = x.SPEC_OBJECT_TYPE_REF
	}

	for x := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_ENUMERATION](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *models.GetGongstructInstancesSet[models.DATATYPE_DEFINITION_XHTML](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *models.GetGongstructInstancesSet[models.EMBEDDED_VALUE](stage) {
		x.Name = x.OTHER_CONTENT
	}

	for x := range *models.GetGongstructInstancesSet[models.ENUM_VALUE](stage) {
		x.Name = x.LONG_NAME
	}

	// DATATYPE_DEFINITION_XHTML EMBEDDED_VALUE ENUM_VALUE

	for x := range *models.GetGongstructInstancesSet[models.SPECIFICATION](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *models.GetGongstructInstancesSet[models.SPECIFICATION_TYPE](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *models.GetGongstructInstancesSet[models.SPEC_HIERARCHY](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *models.GetGongstructInstancesSet[models.SPEC_OBJECT_TYPE](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *models.GetGongstructInstancesSet[models.SPEC_RELATION](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *models.GetGongstructInstancesSet[models.SPEC_RELATION_TYPE](stage) {
		x.Name = x.LONG_NAME
	}
}
