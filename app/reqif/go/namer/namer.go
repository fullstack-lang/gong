package namer

import (
	m "github.com/fullstack-lang/gong/app/reqif/go/models"

	"fmt"
)

type ObjectNamer struct {
}

// SetNamesToElements fill up names of reqif objects
func (objectNamer *ObjectNamer) SetNamesToElements(stage *m.Stage, reqif *m.REQ_IF) {
	idx := 0

	for x := range *stage.GetInstancesSet[*m.ATTRIBUTE_DEFINITION_XHTML]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*m.ATTRIBUTE_DEFINITION_ENUMERATION]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*m.ATTRIBUTE_DEFINITION_BOOLEAN]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*m.ATTRIBUTE_DEFINITION_STRING]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*m.ATTRIBUTE_DEFINITION_REAL]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*m.ATTRIBUTE_DEFINITION_INTEGER]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*m.ATTRIBUTE_DEFINITION_DATE]() {
		x.Name = x.LONG_NAME
	}

	// Attribute Value

	// for x := range *stage.GetInstancesSet[*m.ATTRIBUTE_VALUE_ENUMERATION]() {
	// 	x.Name = x.DEFINITION
	// }

	// anonymous without ATTRIBUTE DEF

	for x := range *stage.GetInstancesSet[*m.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]() {
		x.Name = x.ATTRIBUTE_DEFINITION_ENUMERATION_REF
	}

	for x := range *stage.GetInstancesSet[*m.A_ATTRIBUTE_DEFINITION_XHTML_REF]() {
		x.Name = x.ATTRIBUTE_DEFINITION_XHTML_REF
	}

	// for x := range *stage.GetInstancesSet[*m.A_ATTRIBUTE_VALUE_XHTML_1]() {
	// 	if x.
	// 	x.Name = x.DATATYPE_DEFINITION_ENUMERATION_REF
	// }

	for x := range *stage.GetInstancesSet[*m.A_DATATYPE_DEFINITION_ENUMERATION_REF]() {
		x.Name = x.DATATYPE_DEFINITION_ENUMERATION_REF
	}

	for x := range *stage.GetInstancesSet[*m.A_DATATYPE_DEFINITION_XHTML_REF]() {
		x.Name = x.DATATYPE_DEFINITION_XHTML_REF
	}

	for x := range *stage.GetInstancesSet[*m.A_ENUM_VALUE_REF]() {
		x.Name = x.ENUM_VALUE_REF
	}

	for x := range *stage.GetInstancesSet[*m.A_OBJECT]() {
		x.Name = x.SPEC_OBJECT_REF
	}

	for x := range *stage.GetInstancesSet[*m.A_SPECIFICATION_TYPE_REF]() {
		x.Name = x.SPECIFICATION_TYPE_REF
	}

	for x := range *stage.GetInstancesSet[*m.A_SPEC_OBJECT_TYPE_REF]() {
		x.Name = x.SPEC_OBJECT_TYPE_REF
	}

	for x := range *stage.GetInstancesSet[*m.DATATYPE_DEFINITION_ENUMERATION]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*m.DATATYPE_DEFINITION_BOOLEAN]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*m.DATATYPE_DEFINITION_INTEGER]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*m.DATATYPE_DEFINITION_REAL]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*m.DATATYPE_DEFINITION_DATE]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*m.DATATYPE_DEFINITION_STRING]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*m.DATATYPE_DEFINITION_XHTML]() {
		x.Name = x.LONG_NAME
	}

	for x := range *stage.GetInstancesSet[*m.EMBEDDED_VALUE]() {
		x.Name = x.OTHER_CONTENT
	}

	for x := range *stage.GetInstancesSet[*m.ENUM_VALUE]() {
		x.Name = x.LONG_NAME
	}

	// DATATYPE_DEFINITION_XHTML EMBEDDED_VALUE ENUM_VALUE

	for x := range *stage.GetInstancesSet[*m.SPECIFICATION]() {
		x.Name = x.LONG_NAME
	}

	idx = 0
	for x := range *stage.GetInstancesSet[*m.SPECIFICATION_TYPE]() {
		x.Name = x.LONG_NAME
		if x.LONG_NAME == "" {
			x.Name = fmt.Sprintf("Specification_type_%.2d", idx)
			idx++
		}
	}

	for x := range *stage.GetInstancesSet[*m.SPEC_HIERARCHY]() {
		x.Name = x.LONG_NAME
	}

	objects := reqif.CORE_CONTENT.REQ_IF_CONTENT.SPEC_OBJECTS

	if objects != nil {
		for idx, x := range objects.SPEC_OBJECT {
			x.Name = x.LONG_NAME
			if x.LONG_NAME == "" {
				x.Name = fmt.Sprintf("Spec_object_%.5d", idx)
				idx++
			}
		}
	}

	idx = 0
	for x := range *stage.GetInstancesSet[*m.SPEC_OBJECT_TYPE]() {
		x.Name = x.LONG_NAME
		if x.LONG_NAME == "" {
			x.Name = fmt.Sprintf("Spec_object_type_%.2d", idx)
			idx++
		}
	}

	relations := reqif.CORE_CONTENT.REQ_IF_CONTENT.SPEC_RELATIONS
	if relations != nil {
		for idx, x := range relations.SPEC_RELATION {
			x.Name = x.LONG_NAME
			if x.LONG_NAME == "" {
				x.Name = fmt.Sprintf("Spec_relation_%.4d", idx)
				idx++
			}
		}
	}

	for x := range *stage.GetInstancesSet[*m.SPEC_RELATION_TYPE]() {
		x.Name = x.LONG_NAME
	}
}
