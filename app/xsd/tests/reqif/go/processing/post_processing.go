package processing

import "github.com/fullstack-lang/gong/app/xsd/tests/reqif/go/models"

func PostProcessing(stage *models.Stage) {
	for x := range *stage.GetInstancesSet[*models.ATTRIBUTE_DEFINITION_ENUMERATION]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*models.ATTRIBUTE_DEFINITION_XHTML]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*models.ATTRIBUTE_DEFINITION_ENUMERATION]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*models.ATTRIBUTE_DEFINITION_XHTML]() {
		x.Name = x.LONG_NAME
	}

	// Attribute Value

	// for x := range *stage.GetInstancesSet[*models.ATTRIBUTE_VALUE_ENUMERATION]() {
	// 	x.Name = x.DEFINITION
	// }

	// anonymous without ATTRIBUTE DEF

	for x := range *stage.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]() {
		x.Name = x.ATTRIBUTE_DEFINITION_ENUMERATION_REF
	}

	for x := range *stage.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_XHTML_REF]() {
		x.Name = x.ATTRIBUTE_DEFINITION_XHTML_REF
	}

	// for x := range *stage.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_XHTML_1]() {
	// 	if x.
	// 	x.Name = x.DATATYPE_DEFINITION_ENUMERATION_REF
	// }

	for x := range *stage.GetInstancesSet[*models.A_DATATYPE_DEFINITION_ENUMERATION_REF]() {
		x.Name = x.DATATYPE_DEFINITION_ENUMERATION_REF
	}

	for x := range *stage.GetInstancesSet[*models.A_DATATYPE_DEFINITION_XHTML_REF]() {
		x.Name = x.DATATYPE_DEFINITION_XHTML_REF
	}

	for x := range *stage.GetInstancesSet[*models.A_ENUM_VALUE_REF]() {
		x.Name = x.ENUM_VALUE_REF
	}

	for x := range *stage.GetInstancesSet[*models.A_OBJECT]() {
		x.Name = x.SPEC_OBJECT_REF
	}

	for x := range *stage.GetInstancesSet[*models.A_SPECIFICATION_TYPE_REF]() {
		x.Name = x.SPECIFICATION_TYPE_REF
	}

	for x := range *stage.GetInstancesSet[*models.A_SPEC_OBJECT_TYPE_REF]() {
		x.Name = x.SPEC_OBJECT_TYPE_REF
	}

	for x := range *stage.GetInstancesSet[*models.DATATYPE_DEFINITION_ENUMERATION]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*models.DATATYPE_DEFINITION_XHTML]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*models.EMBEDDED_VALUE]() {
		x.Name = x.OTHER_CONTENT
	}

	for x := range *stage.GetInstancesSet[*models.ENUM_VALUE]() {
		x.Name = x.LONG_NAME
	}

	// DATATYPE_DEFINITION_XHTML EMBEDDED_VALUE ENUM_VALUE

	for x := range *stage.GetInstancesSet[*models.SPECIFICATION]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*models.SPECIFICATION_TYPE]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*models.SPEC_HIERARCHY]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*models.SPEC_OBJECT]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*models.SPEC_OBJECT_TYPE]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*models.SPEC_RELATION]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*models.SPEC_RELATION_TYPE]() {
		x.Name = x.LONG_NAME
	}
}
