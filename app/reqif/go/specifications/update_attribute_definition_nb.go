package specifications

import (
	"log"

	m "github.com/fullstack-lang/gong/app/reqif/go/models"
)

// UpdateAttributeDefinitionNb parses the selected specification and updates the number of attributes definitions
func (o *SpecificationsTreeStageUpdater) UpdateAttributeDefinitionNb(
	stager *m.Stager,
) {
	stage := stager.GetStage()

	specification := GetSelectedSpecification(stage)

	if specification == nil {
		return
	}

	stager.Map_SPEC_OBJECT_TYPE_Spec_nbInstance = initializeNbInstanceMap[m.SPEC_OBJECT_TYPE]()
	stager.Map_SPECIFICATION_TYPE_Spec_nbInstance = initializeNbInstanceMap[m.SPECIFICATION_TYPE]()
	stager.Map_SPEC_RELATION_TYPE_Spec_nbInstance = initializeNbInstanceMap[m.SPEC_RELATION_TYPE]()

	specificationType, ok := stager.Map_id_SPECIFICATION_TYPE[specification.TYPE.SPECIFICATION_TYPE_REF]
	if !ok {
		log.Panic("specRelation.TYPE.SPECIFICATION_TYPE_REF", specification.TYPE.SPECIFICATION_TYPE_REF,
			"unknown relation type")
	}
	stager.Map_SPECIFICATION_TYPE_Spec_nbInstance[specificationType]++

	stager.Map_ATTRIBUTE_DEFINITION_XHTML_Spec_nbInstance = initializeNbInstanceMap[m.ATTRIBUTE_DEFINITION_XHTML]()
	stager.Map_ATTRIBUTE_DEFINITION_STRING_Spec_nbInstance = initializeNbInstanceMap[m.ATTRIBUTE_DEFINITION_STRING]()
	stager.Map_ATTRIBUTE_DEFINITION_BOOLEAN_Spec_nbInstance = initializeNbInstanceMap[m.ATTRIBUTE_DEFINITION_BOOLEAN]()
	stager.Map_ATTRIBUTE_DEFINITION_INTEGER_Spec_nbInstance = initializeNbInstanceMap[m.ATTRIBUTE_DEFINITION_INTEGER]()
	stager.Map_ATTRIBUTE_DEFINITION_DATE_Spec_nbInstance = initializeNbInstanceMap[m.ATTRIBUTE_DEFINITION_DATE]()
	stager.Map_ATTRIBUTE_DEFINITION_REAL_Spec_nbInstance = initializeNbInstanceMap[m.ATTRIBUTE_DEFINITION_REAL]()
	stager.Map_ATTRIBUTE_DEFINITION_ENUMERATION_Spec_nbInstance = initializeNbInstanceMap[m.ATTRIBUTE_DEFINITION_ENUMERATION]()
	stager.Map_ENUM_VALUE_Spec_nbInstance = initializeNbInstanceMap[m.ENUM_VALUE]()

	for _, specHierarchy := range specification.CHILDREN.SPEC_HIERARCHY {
		UpdateAttributeDefinitionNbRecursice(
			stager,
			specHierarchy,
		)
	}

}
func UpdateAttributeDefinitionNbRecursice(
	stager *m.Stager,
	specHierarchy *m.SPEC_HIERARCHY,
) {

	specObject, ok := stager.Map_id_SPEC_OBJECT[specHierarchy.OBJECT.SPEC_OBJECT_REF]
	if !ok {
		log.Panic("specHierarchy.OBJECT.SPEC_OBJECT_REF", specHierarchy.OBJECT.SPEC_OBJECT_REF,
			"unknown ref")
	}

	specObjectType, ok := stager.Map_id_SPEC_OBJECT_TYPE[specObject.TYPE.SPEC_OBJECT_TYPE_REF]
	if !ok {
		log.Panic("specObject.TYPE.SPEC_OBJECT_TYPE_REF", specObject.TYPE.SPEC_OBJECT_TYPE_REF,
			"unknown ref")
	}

	stager.Map_SPEC_OBJECT_TYPE_Spec_nbInstance[specObjectType]++

	appendAttributeXHTML(stager, specObject)
	appendAttributeString(stager, specObject)
	appendAttributeBoolean(stager, specObject)
	appendAttributeInteger(stager, specObject)
	appendAttributeReal(stager, specObject)
	appendAttributeDate(stager, specObject)
	appendAttributeEnum(stager, specObject)

	if specHierarchy.CHILDREN != nil {
		for _, specHierarchyChildren := range specHierarchy.CHILDREN.SPEC_HIERARCHY {
			UpdateAttributeDefinitionNbRecursice(stager, specHierarchyChildren)
		}
	}

}

// appendAttributeXHTMLRows formats and appends XHTML attributes as markdown table rows.
func appendAttributeXHTML(stager *m.Stager, specObject *m.SPEC_OBJECT) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_XHTML {

		if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_XHTML[attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF]; ok {
			stager.Map_ATTRIBUTE_DEFINITION_XHTML_Spec_nbInstance[datatype]++
		} else {
			log.Panic("ATTRIBUTE_DEFINITION_XHTML_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_XHTML_REF, "unknown ref")
		}

	}
}

// appendAttributeStringRows formats and appends String attributes as markdown table rows.
func appendAttributeString(stager *m.Stager, specObject *m.SPEC_OBJECT) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_STRING {

		if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_STRING[attribute.DEFINITION.ATTRIBUTE_DEFINITION_STRING_REF]; ok {
			stager.Map_ATTRIBUTE_DEFINITION_STRING_Spec_nbInstance[datatype]++
		} else {
			log.Panic("ATTRIBUTE_DEFINITION_STRING_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_STRING_REF, "unknown ref")
		}

	}
}

// appendAttributeBooleanRows formats and appends Boolean attributes as markdown table rows.
func appendAttributeBoolean(stager *m.Stager, specObject *m.SPEC_OBJECT) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_BOOLEAN {

		if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_BOOLEAN[attribute.DEFINITION.ATTRIBUTE_DEFINITION_BOOLEAN_REF]; ok {
			stager.Map_ATTRIBUTE_DEFINITION_BOOLEAN_Spec_nbInstance[datatype]++
		} else {
			log.Panic("ATTRIBUTE_DEFINITION_BOOLEAN_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_BOOLEAN_REF, "unknown ref")
		}

	}
}

// appendAttributeIntegerRows formats and appends Integer attributes as markdown table rows.
func appendAttributeInteger(stager *m.Stager, specObject *m.SPEC_OBJECT) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_INTEGER {

		if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_INTEGER[attribute.DEFINITION.ATTRIBUTE_DEFINITION_INTEGER_REF]; ok {
			stager.Map_ATTRIBUTE_DEFINITION_INTEGER_Spec_nbInstance[datatype]++
		} else {
			log.Panic("ATTRIBUTE_DEFINITION_INTEGER_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_INTEGER_REF, "unknown ref")
		}

	}
}

// appendAttributeDateRows formats and appends Date attributes as markdown table rows.
func appendAttributeDate(stager *m.Stager, specObject *m.SPEC_OBJECT) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_DATE {

		if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_DATE[attribute.DEFINITION.ATTRIBUTE_DEFINITION_DATE_REF]; ok {
			stager.Map_ATTRIBUTE_DEFINITION_DATE_Spec_nbInstance[datatype]++
		} else {
			log.Panic("ATTRIBUTE_DEFINITION_DATE_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_DATE_REF, "unknown ref")
		}

	}
}

// appendAttributeRealRows formats and appends Real attributes as markdown table rows.
func appendAttributeReal(stager *m.Stager, specObject *m.SPEC_OBJECT) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_REAL {

		if datatype, ok := stager.Map_id_ATTRIBUTE_DEFINITION_REAL[attribute.DEFINITION.ATTRIBUTE_DEFINITION_REAL_REF]; ok {
			stager.Map_ATTRIBUTE_DEFINITION_REAL_Spec_nbInstance[datatype]++
		} else {
			log.Panic("ATTRIBUTE_DEFINITION_REAL_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_REAL_REF, "unknown ref")
		}

	}
}

// appendAttributeEnumRows formats and appends Enumeration attributes as markdown table rows.
func appendAttributeEnum(stager *m.Stager, specObject *m.SPEC_OBJECT) {
	for _, attribute := range specObject.VALUES.ATTRIBUTE_VALUE_ENUMERATION {
		if enumType, ok := stager.Map_id_ATTRIBUTE_DEFINITION_ENUMERATION[attribute.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF]; ok {
			stager.Map_ATTRIBUTE_DEFINITION_ENUMERATION_Spec_nbInstance[enumType]++
		} else {
			log.Panic("ATTRIBUTE_DEFINITION_ENUMERATION_REF", attribute.DEFINITION.ATTRIBUTE_DEFINITION_ENUMERATION_REF, "unknown ref")
		}

		if len(attribute.VALUES.ENUM_VALUE_REF) > 0 {
			valueIdentifier := attribute.VALUES.ENUM_VALUE_REF
			if enumValue, ok := stager.Map_id_ENUM_VALUE[valueIdentifier]; ok {
				stager.Map_ENUM_VALUE_Spec_nbInstance[enumValue]++
			} else {
				log.Panic("ENUM_VALUE_REF", valueIdentifier, "unknown ref in Map_id_ENUM_VALUE")
			}
		}
	}
}

// Generic function to initialize a map with *T as key and int as value
func initializeNbInstanceMap[T any]() map[*T]int {
	return make(map[*T]int)
}
