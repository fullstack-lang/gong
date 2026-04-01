package namer

import (
	m "github.com/fullstack-lang/gong/dsm/reqif/go/models"

	"fmt"
)

type ObjectNamer struct {
}

// SetNamesToElements fill up names of reqif objects
func (objectNamer *ObjectNamer) SetNamesToElements(stage *m.Stage, reqif *m.REQ_IF) {
	idx := 0

	for x := range *m.GetGongstructInstancesSet[m.ATTRIBUTE_DEFINITION_XHTML](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *m.GetGongstructInstancesSet[m.ATTRIBUTE_DEFINITION_ENUMERATION](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *m.GetGongstructInstancesSet[m.ATTRIBUTE_DEFINITION_BOOLEAN](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *m.GetGongstructInstancesSet[m.ATTRIBUTE_DEFINITION_STRING](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *m.GetGongstructInstancesSet[m.ATTRIBUTE_DEFINITION_REAL](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *m.GetGongstructInstancesSet[m.ATTRIBUTE_DEFINITION_INTEGER](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *m.GetGongstructInstancesSet[m.ATTRIBUTE_DEFINITION_DATE](stage) {
		x.Name = x.LONG_NAME
	}

	// Attribute Value

	// for x := range *m.GetGongstructInstancesSet[m.ATTRIBUTE_VALUE_ENUMERATION](stage) {
	// 	x.Name = x.DEFINITION
	// }

	// anonymous without ATTRIBUTE DEF

	for x := range *m.GetGongstructInstancesSet[m.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](stage) {
		x.Name = x.ATTRIBUTE_DEFINITION_ENUMERATION_REF
	}

	for x := range *m.GetGongstructInstancesSet[m.A_ATTRIBUTE_DEFINITION_XHTML_REF](stage) {
		x.Name = x.ATTRIBUTE_DEFINITION_XHTML_REF
	}

	// for x := range *m.GetGongstructInstancesSet[m.A_ATTRIBUTE_VALUE_XHTML_1](stage) {
	// 	if x.
	// 	x.Name = x.DATATYPE_DEFINITION_ENUMERATION_REF
	// }

	for x := range *m.GetGongstructInstancesSet[m.A_DATATYPE_DEFINITION_ENUMERATION_REF](stage) {
		x.Name = x.DATATYPE_DEFINITION_ENUMERATION_REF
	}

	for x := range *m.GetGongstructInstancesSet[m.A_DATATYPE_DEFINITION_XHTML_REF](stage) {
		x.Name = x.DATATYPE_DEFINITION_XHTML_REF
	}

	for x := range *m.GetGongstructInstancesSet[m.A_ENUM_VALUE_REF](stage) {
		x.Name = x.ENUM_VALUE_REF
	}

	for x := range *m.GetGongstructInstancesSet[m.A_OBJECT](stage) {
		x.Name = x.SPEC_OBJECT_REF
	}

	for x := range *m.GetGongstructInstancesSet[m.A_SPECIFICATION_TYPE_REF](stage) {
		x.Name = x.SPECIFICATION_TYPE_REF
	}

	for x := range *m.GetGongstructInstancesSet[m.A_SPEC_OBJECT_TYPE_REF](stage) {
		x.Name = x.SPEC_OBJECT_TYPE_REF
	}

	for x := range *m.GetGongstructInstancesSet[m.DATATYPE_DEFINITION_ENUMERATION](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *m.GetGongstructInstancesSet[m.DATATYPE_DEFINITION_BOOLEAN](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *m.GetGongstructInstancesSet[m.DATATYPE_DEFINITION_INTEGER](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *m.GetGongstructInstancesSet[m.DATATYPE_DEFINITION_REAL](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *m.GetGongstructInstancesSet[m.DATATYPE_DEFINITION_DATE](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *m.GetGongstructInstancesSet[m.DATATYPE_DEFINITION_STRING](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *m.GetGongstructInstancesSet[m.DATATYPE_DEFINITION_XHTML](stage) {
		x.Name = x.LONG_NAME
	}

	for x := range *m.GetGongstructInstancesSet[m.EMBEDDED_VALUE](stage) {
		x.Name = x.OTHER_CONTENT
	}

	for x := range *m.GetGongstructInstancesSet[m.ENUM_VALUE](stage) {
		x.Name = x.LONG_NAME
	}

	// DATATYPE_DEFINITION_XHTML EMBEDDED_VALUE ENUM_VALUE

	for x := range *m.GetGongstructInstancesSet[m.SPECIFICATION](stage) {
		x.Name = x.LONG_NAME
	}

	idx = 0
	for x := range *m.GetGongstructInstancesSet[m.SPECIFICATION_TYPE](stage) {
		x.Name = x.LONG_NAME
		if x.LONG_NAME == "" {
			x.Name = fmt.Sprintf("Specification_type_%.2d", idx)
			idx++
		}
	}

	for x := range *m.GetGongstructInstancesSet[m.SPEC_HIERARCHY](stage) {
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
	for x := range *m.GetGongstructInstancesSet[m.SPEC_OBJECT_TYPE](stage) {
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

	for x := range *m.GetGongstructInstancesSet[m.SPEC_RELATION_TYPE](stage) {
		x.Name = x.LONG_NAME
	}
}
