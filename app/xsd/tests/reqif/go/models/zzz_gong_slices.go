// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_BOOLEAN
	// insertion point per field
	stage.A_ATTRIBUTE_VALUE_BOOLEAN_ATTRIBUTE_VALUE_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*A_ATTRIBUTE_VALUE_BOOLEAN)
	for a_attribute_value_boolean := range stage.A_ATTRIBUTE_VALUE_BOOLEANs {
		_ = a_attribute_value_boolean
		for _, _attribute_value_boolean := range a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN {
			stage.A_ATTRIBUTE_VALUE_BOOLEAN_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[_attribute_value_boolean] = a_attribute_value_boolean
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_DATE
	// insertion point per field
	stage.A_ATTRIBUTE_VALUE_DATE_ATTRIBUTE_VALUE_DATE_reverseMap = make(map[*ATTRIBUTE_VALUE_DATE]*A_ATTRIBUTE_VALUE_DATE)
	for a_attribute_value_date := range stage.A_ATTRIBUTE_VALUE_DATEs {
		_ = a_attribute_value_date
		for _, _attribute_value_date := range a_attribute_value_date.ATTRIBUTE_VALUE_DATE {
			stage.A_ATTRIBUTE_VALUE_DATE_ATTRIBUTE_VALUE_DATE_reverseMap[_attribute_value_date] = a_attribute_value_date
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_ENUMERATION
	// insertion point per field
	stage.A_ATTRIBUTE_VALUE_ENUMERATION_ATTRIBUTE_VALUE_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*A_ATTRIBUTE_VALUE_ENUMERATION)
	for a_attribute_value_enumeration := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
		_ = a_attribute_value_enumeration
		for _, _attribute_value_enumeration := range a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION {
			stage.A_ATTRIBUTE_VALUE_ENUMERATION_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[_attribute_value_enumeration] = a_attribute_value_enumeration
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_INTEGER
	// insertion point per field
	stage.A_ATTRIBUTE_VALUE_INTEGER_ATTRIBUTE_VALUE_INTEGER_reverseMap = make(map[*ATTRIBUTE_VALUE_INTEGER]*A_ATTRIBUTE_VALUE_INTEGER)
	for a_attribute_value_integer := range stage.A_ATTRIBUTE_VALUE_INTEGERs {
		_ = a_attribute_value_integer
		for _, _attribute_value_integer := range a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER {
			stage.A_ATTRIBUTE_VALUE_INTEGER_ATTRIBUTE_VALUE_INTEGER_reverseMap[_attribute_value_integer] = a_attribute_value_integer
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_REAL
	// insertion point per field
	stage.A_ATTRIBUTE_VALUE_REAL_ATTRIBUTE_VALUE_REAL_reverseMap = make(map[*ATTRIBUTE_VALUE_REAL]*A_ATTRIBUTE_VALUE_REAL)
	for a_attribute_value_real := range stage.A_ATTRIBUTE_VALUE_REALs {
		_ = a_attribute_value_real
		for _, _attribute_value_real := range a_attribute_value_real.ATTRIBUTE_VALUE_REAL {
			stage.A_ATTRIBUTE_VALUE_REAL_ATTRIBUTE_VALUE_REAL_reverseMap[_attribute_value_real] = a_attribute_value_real
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_STRING
	// insertion point per field
	stage.A_ATTRIBUTE_VALUE_STRING_ATTRIBUTE_VALUE_STRING_reverseMap = make(map[*ATTRIBUTE_VALUE_STRING]*A_ATTRIBUTE_VALUE_STRING)
	for a_attribute_value_string := range stage.A_ATTRIBUTE_VALUE_STRINGs {
		_ = a_attribute_value_string
		for _, _attribute_value_string := range a_attribute_value_string.ATTRIBUTE_VALUE_STRING {
			stage.A_ATTRIBUTE_VALUE_STRING_ATTRIBUTE_VALUE_STRING_reverseMap[_attribute_value_string] = a_attribute_value_string
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_XHTML
	// insertion point per field
	stage.A_ATTRIBUTE_VALUE_XHTML_ATTRIBUTE_VALUE_XHTML_reverseMap = make(map[*ATTRIBUTE_VALUE_XHTML]*A_ATTRIBUTE_VALUE_XHTML)
	for a_attribute_value_xhtml := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
		_ = a_attribute_value_xhtml
		for _, _attribute_value_xhtml := range a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML {
			stage.A_ATTRIBUTE_VALUE_XHTML_ATTRIBUTE_VALUE_XHTML_reverseMap[_attribute_value_xhtml] = a_attribute_value_xhtml
		}
	}

	// Compute reverse map for named struct A_ATTRIBUTE_VALUE_XHTML_1
	// insertion point per field
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_boolean := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_BOOLEAN_reverseMap[_attribute_value_boolean] = a_attribute_value_xhtml_1
		}
	}
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_DATE_reverseMap = make(map[*ATTRIBUTE_VALUE_DATE]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_date := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_DATE_reverseMap[_attribute_value_date] = a_attribute_value_xhtml_1
		}
	}
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_enumeration := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_ENUMERATION_reverseMap[_attribute_value_enumeration] = a_attribute_value_xhtml_1
		}
	}
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_INTEGER_reverseMap = make(map[*ATTRIBUTE_VALUE_INTEGER]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_integer := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_INTEGER_reverseMap[_attribute_value_integer] = a_attribute_value_xhtml_1
		}
	}
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_REAL_reverseMap = make(map[*ATTRIBUTE_VALUE_REAL]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_real := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_REAL_reverseMap[_attribute_value_real] = a_attribute_value_xhtml_1
		}
	}
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_STRING_reverseMap = make(map[*ATTRIBUTE_VALUE_STRING]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_string := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_STRING_reverseMap[_attribute_value_string] = a_attribute_value_xhtml_1
		}
	}
	stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_XHTML_reverseMap = make(map[*ATTRIBUTE_VALUE_XHTML]*A_ATTRIBUTE_VALUE_XHTML_1)
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_ = a_attribute_value_xhtml_1
		for _, _attribute_value_xhtml := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML {
			stage.A_ATTRIBUTE_VALUE_XHTML_1_ATTRIBUTE_VALUE_XHTML_reverseMap[_attribute_value_xhtml] = a_attribute_value_xhtml_1
		}
	}

	// Compute reverse map for named struct A_CHILDREN
	// insertion point per field
	stage.A_CHILDREN_SPEC_HIERARCHY_reverseMap = make(map[*SPEC_HIERARCHY]*A_CHILDREN)
	for a_children := range stage.A_CHILDRENs {
		_ = a_children
		for _, _spec_hierarchy := range a_children.SPEC_HIERARCHY {
			stage.A_CHILDREN_SPEC_HIERARCHY_reverseMap[_spec_hierarchy] = a_children
		}
	}

	// Compute reverse map for named struct A_DATATYPES
	// insertion point per field
	stage.A_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap = make(map[*DATATYPE_DEFINITION_BOOLEAN]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_boolean := range a_datatypes.DATATYPE_DEFINITION_BOOLEAN {
			stage.A_DATATYPES_DATATYPE_DEFINITION_BOOLEAN_reverseMap[_datatype_definition_boolean] = a_datatypes
		}
	}
	stage.A_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap = make(map[*DATATYPE_DEFINITION_DATE]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_date := range a_datatypes.DATATYPE_DEFINITION_DATE {
			stage.A_DATATYPES_DATATYPE_DEFINITION_DATE_reverseMap[_datatype_definition_date] = a_datatypes
		}
	}
	stage.A_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap = make(map[*DATATYPE_DEFINITION_ENUMERATION]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_enumeration := range a_datatypes.DATATYPE_DEFINITION_ENUMERATION {
			stage.A_DATATYPES_DATATYPE_DEFINITION_ENUMERATION_reverseMap[_datatype_definition_enumeration] = a_datatypes
		}
	}
	stage.A_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap = make(map[*DATATYPE_DEFINITION_INTEGER]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_integer := range a_datatypes.DATATYPE_DEFINITION_INTEGER {
			stage.A_DATATYPES_DATATYPE_DEFINITION_INTEGER_reverseMap[_datatype_definition_integer] = a_datatypes
		}
	}
	stage.A_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap = make(map[*DATATYPE_DEFINITION_REAL]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_real := range a_datatypes.DATATYPE_DEFINITION_REAL {
			stage.A_DATATYPES_DATATYPE_DEFINITION_REAL_reverseMap[_datatype_definition_real] = a_datatypes
		}
	}
	stage.A_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap = make(map[*DATATYPE_DEFINITION_STRING]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_string := range a_datatypes.DATATYPE_DEFINITION_STRING {
			stage.A_DATATYPES_DATATYPE_DEFINITION_STRING_reverseMap[_datatype_definition_string] = a_datatypes
		}
	}
	stage.A_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap = make(map[*DATATYPE_DEFINITION_XHTML]*A_DATATYPES)
	for a_datatypes := range stage.A_DATATYPESs {
		_ = a_datatypes
		for _, _datatype_definition_xhtml := range a_datatypes.DATATYPE_DEFINITION_XHTML {
			stage.A_DATATYPES_DATATYPE_DEFINITION_XHTML_reverseMap[_datatype_definition_xhtml] = a_datatypes
		}
	}

	// Compute reverse map for named struct A_SPECIFICATIONS
	// insertion point per field
	stage.A_SPECIFICATIONS_SPECIFICATION_reverseMap = make(map[*SPECIFICATION]*A_SPECIFICATIONS)
	for a_specifications := range stage.A_SPECIFICATIONSs {
		_ = a_specifications
		for _, _specification := range a_specifications.SPECIFICATION {
			stage.A_SPECIFICATIONS_SPECIFICATION_reverseMap[_specification] = a_specifications
		}
	}

	// Compute reverse map for named struct A_SPECIFIED_VALUES
	// insertion point per field
	stage.A_SPECIFIED_VALUES_ENUM_VALUE_reverseMap = make(map[*ENUM_VALUE]*A_SPECIFIED_VALUES)
	for a_specified_values := range stage.A_SPECIFIED_VALUESs {
		_ = a_specified_values
		for _, _enum_value := range a_specified_values.ENUM_VALUE {
			stage.A_SPECIFIED_VALUES_ENUM_VALUE_reverseMap[_enum_value] = a_specified_values
		}
	}

	// Compute reverse map for named struct A_SPEC_ATTRIBUTES
	// insertion point per field
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_boolean := range a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_BOOLEAN_reverseMap[_attribute_definition_boolean] = a_spec_attributes
		}
	}
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap = make(map[*ATTRIBUTE_DEFINITION_DATE]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_date := range a_spec_attributes.ATTRIBUTE_DEFINITION_DATE {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_DATE_reverseMap[_attribute_definition_date] = a_spec_attributes
		}
	}
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_enumeration := range a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_ENUMERATION_reverseMap[_attribute_definition_enumeration] = a_spec_attributes
		}
	}
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap = make(map[*ATTRIBUTE_DEFINITION_INTEGER]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_integer := range a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_INTEGER_reverseMap[_attribute_definition_integer] = a_spec_attributes
		}
	}
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap = make(map[*ATTRIBUTE_DEFINITION_REAL]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_real := range a_spec_attributes.ATTRIBUTE_DEFINITION_REAL {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_REAL_reverseMap[_attribute_definition_real] = a_spec_attributes
		}
	}
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap = make(map[*ATTRIBUTE_DEFINITION_STRING]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_string := range a_spec_attributes.ATTRIBUTE_DEFINITION_STRING {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_STRING_reverseMap[_attribute_definition_string] = a_spec_attributes
		}
	}
	stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap = make(map[*ATTRIBUTE_DEFINITION_XHTML]*A_SPEC_ATTRIBUTES)
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		_ = a_spec_attributes
		for _, _attribute_definition_xhtml := range a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML {
			stage.A_SPEC_ATTRIBUTES_ATTRIBUTE_DEFINITION_XHTML_reverseMap[_attribute_definition_xhtml] = a_spec_attributes
		}
	}

	// Compute reverse map for named struct A_SPEC_OBJECTS
	// insertion point per field
	stage.A_SPEC_OBJECTS_SPEC_OBJECT_reverseMap = make(map[*SPEC_OBJECT]*A_SPEC_OBJECTS)
	for a_spec_objects := range stage.A_SPEC_OBJECTSs {
		_ = a_spec_objects
		for _, _spec_object := range a_spec_objects.SPEC_OBJECT {
			stage.A_SPEC_OBJECTS_SPEC_OBJECT_reverseMap[_spec_object] = a_spec_objects
		}
	}

	// Compute reverse map for named struct A_SPEC_RELATIONS
	// insertion point per field
	stage.A_SPEC_RELATIONS_SPEC_RELATION_reverseMap = make(map[*SPEC_RELATION]*A_SPEC_RELATIONS)
	for a_spec_relations := range stage.A_SPEC_RELATIONSs {
		_ = a_spec_relations
		for _, _spec_relation := range a_spec_relations.SPEC_RELATION {
			stage.A_SPEC_RELATIONS_SPEC_RELATION_reverseMap[_spec_relation] = a_spec_relations
		}
	}

	// Compute reverse map for named struct A_SPEC_RELATION_GROUPS
	// insertion point per field
	stage.A_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap = make(map[*RELATION_GROUP]*A_SPEC_RELATION_GROUPS)
	for a_spec_relation_groups := range stage.A_SPEC_RELATION_GROUPSs {
		_ = a_spec_relation_groups
		for _, _relation_group := range a_spec_relation_groups.RELATION_GROUP {
			stage.A_SPEC_RELATION_GROUPS_RELATION_GROUP_reverseMap[_relation_group] = a_spec_relation_groups
		}
	}

	// Compute reverse map for named struct A_SPEC_TYPES
	// insertion point per field
	stage.A_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap = make(map[*RELATION_GROUP_TYPE]*A_SPEC_TYPES)
	for a_spec_types := range stage.A_SPEC_TYPESs {
		_ = a_spec_types
		for _, _relation_group_type := range a_spec_types.RELATION_GROUP_TYPE {
			stage.A_SPEC_TYPES_RELATION_GROUP_TYPE_reverseMap[_relation_group_type] = a_spec_types
		}
	}
	stage.A_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap = make(map[*SPEC_OBJECT_TYPE]*A_SPEC_TYPES)
	for a_spec_types := range stage.A_SPEC_TYPESs {
		_ = a_spec_types
		for _, _spec_object_type := range a_spec_types.SPEC_OBJECT_TYPE {
			stage.A_SPEC_TYPES_SPEC_OBJECT_TYPE_reverseMap[_spec_object_type] = a_spec_types
		}
	}
	stage.A_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap = make(map[*SPEC_RELATION_TYPE]*A_SPEC_TYPES)
	for a_spec_types := range stage.A_SPEC_TYPESs {
		_ = a_spec_types
		for _, _spec_relation_type := range a_spec_types.SPEC_RELATION_TYPE {
			stage.A_SPEC_TYPES_SPEC_RELATION_TYPE_reverseMap[_spec_relation_type] = a_spec_types
		}
	}
	stage.A_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap = make(map[*SPECIFICATION_TYPE]*A_SPEC_TYPES)
	for a_spec_types := range stage.A_SPEC_TYPESs {
		_ = a_spec_types
		for _, _specification_type := range a_spec_types.SPECIFICATION_TYPE {
			stage.A_SPEC_TYPES_SPECIFICATION_TYPE_reverseMap[_specification_type] = a_spec_types
		}
	}

	// Compute reverse map for named struct A_TOOL_EXTENSIONS
	// insertion point per field
	stage.A_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap = make(map[*REQ_IF_TOOL_EXTENSION]*A_TOOL_EXTENSIONS)
	for a_tool_extensions := range stage.A_TOOL_EXTENSIONSs {
		_ = a_tool_extensions
		for _, _req_if_tool_extension := range a_tool_extensions.REQ_IF_TOOL_EXTENSION {
			stage.A_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap[_req_if_tool_extension] = a_tool_extensions
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.ALTERNATIVE_IDs)

	res = __gong__appendInstances(res, stage.ATTRIBUTE_DEFINITION_BOOLEANs)

	res = __gong__appendInstances(res, stage.ATTRIBUTE_DEFINITION_DATEs)

	res = __gong__appendInstances(res, stage.ATTRIBUTE_DEFINITION_ENUMERATIONs)

	res = __gong__appendInstances(res, stage.ATTRIBUTE_DEFINITION_INTEGERs)

	res = __gong__appendInstances(res, stage.ATTRIBUTE_DEFINITION_REALs)

	res = __gong__appendInstances(res, stage.ATTRIBUTE_DEFINITION_STRINGs)

	res = __gong__appendInstances(res, stage.ATTRIBUTE_DEFINITION_XHTMLs)

	res = __gong__appendInstances(res, stage.ATTRIBUTE_VALUE_BOOLEANs)

	res = __gong__appendInstances(res, stage.ATTRIBUTE_VALUE_DATEs)

	res = __gong__appendInstances(res, stage.ATTRIBUTE_VALUE_ENUMERATIONs)

	res = __gong__appendInstances(res, stage.ATTRIBUTE_VALUE_INTEGERs)

	res = __gong__appendInstances(res, stage.ATTRIBUTE_VALUE_REALs)

	res = __gong__appendInstances(res, stage.ATTRIBUTE_VALUE_STRINGs)

	res = __gong__appendInstances(res, stage.ATTRIBUTE_VALUE_XHTMLs)

	res = __gong__appendInstances(res, stage.A_ALTERNATIVE_IDs)

	res = __gong__appendInstances(res, stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs)

	res = __gong__appendInstances(res, stage.A_ATTRIBUTE_DEFINITION_DATE_REFs)

	res = __gong__appendInstances(res, stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs)

	res = __gong__appendInstances(res, stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs)

	res = __gong__appendInstances(res, stage.A_ATTRIBUTE_DEFINITION_REAL_REFs)

	res = __gong__appendInstances(res, stage.A_ATTRIBUTE_DEFINITION_STRING_REFs)

	res = __gong__appendInstances(res, stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs)

	res = __gong__appendInstances(res, stage.A_ATTRIBUTE_VALUE_BOOLEANs)

	res = __gong__appendInstances(res, stage.A_ATTRIBUTE_VALUE_DATEs)

	res = __gong__appendInstances(res, stage.A_ATTRIBUTE_VALUE_ENUMERATIONs)

	res = __gong__appendInstances(res, stage.A_ATTRIBUTE_VALUE_INTEGERs)

	res = __gong__appendInstances(res, stage.A_ATTRIBUTE_VALUE_REALs)

	res = __gong__appendInstances(res, stage.A_ATTRIBUTE_VALUE_STRINGs)

	res = __gong__appendInstances(res, stage.A_ATTRIBUTE_VALUE_XHTMLs)

	res = __gong__appendInstances(res, stage.A_ATTRIBUTE_VALUE_XHTML_1s)

	res = __gong__appendInstances(res, stage.A_CHILDRENs)

	res = __gong__appendInstances(res, stage.A_CORE_CONTENTs)

	res = __gong__appendInstances(res, stage.A_DATATYPESs)

	res = __gong__appendInstances(res, stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs)

	res = __gong__appendInstances(res, stage.A_DATATYPE_DEFINITION_DATE_REFs)

	res = __gong__appendInstances(res, stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs)

	res = __gong__appendInstances(res, stage.A_DATATYPE_DEFINITION_INTEGER_REFs)

	res = __gong__appendInstances(res, stage.A_DATATYPE_DEFINITION_REAL_REFs)

	res = __gong__appendInstances(res, stage.A_DATATYPE_DEFINITION_STRING_REFs)

	res = __gong__appendInstances(res, stage.A_DATATYPE_DEFINITION_XHTML_REFs)

	res = __gong__appendInstances(res, stage.A_EDITABLE_ATTSs)

	res = __gong__appendInstances(res, stage.A_ENUM_VALUE_REFs)

	res = __gong__appendInstances(res, stage.A_OBJECTs)

	res = __gong__appendInstances(res, stage.A_PROPERTIESs)

	res = __gong__appendInstances(res, stage.A_RELATION_GROUP_TYPE_REFs)

	res = __gong__appendInstances(res, stage.A_SOURCE_1s)

	res = __gong__appendInstances(res, stage.A_SOURCE_SPECIFICATION_1s)

	res = __gong__appendInstances(res, stage.A_SPECIFICATIONSs)

	res = __gong__appendInstances(res, stage.A_SPECIFICATION_TYPE_REFs)

	res = __gong__appendInstances(res, stage.A_SPECIFIED_VALUESs)

	res = __gong__appendInstances(res, stage.A_SPEC_ATTRIBUTESs)

	res = __gong__appendInstances(res, stage.A_SPEC_OBJECTSs)

	res = __gong__appendInstances(res, stage.A_SPEC_OBJECT_TYPE_REFs)

	res = __gong__appendInstances(res, stage.A_SPEC_RELATIONSs)

	res = __gong__appendInstances(res, stage.A_SPEC_RELATION_GROUPSs)

	res = __gong__appendInstances(res, stage.A_SPEC_RELATION_REFs)

	res = __gong__appendInstances(res, stage.A_SPEC_RELATION_TYPE_REFs)

	res = __gong__appendInstances(res, stage.A_SPEC_TYPESs)

	res = __gong__appendInstances(res, stage.A_THE_HEADERs)

	res = __gong__appendInstances(res, stage.A_TOOL_EXTENSIONSs)

	res = __gong__appendInstances(res, stage.DATATYPE_DEFINITION_BOOLEANs)

	res = __gong__appendInstances(res, stage.DATATYPE_DEFINITION_DATEs)

	res = __gong__appendInstances(res, stage.DATATYPE_DEFINITION_ENUMERATIONs)

	res = __gong__appendInstances(res, stage.DATATYPE_DEFINITION_INTEGERs)

	res = __gong__appendInstances(res, stage.DATATYPE_DEFINITION_REALs)

	res = __gong__appendInstances(res, stage.DATATYPE_DEFINITION_STRINGs)

	res = __gong__appendInstances(res, stage.DATATYPE_DEFINITION_XHTMLs)

	res = __gong__appendInstances(res, stage.EMBEDDED_VALUEs)

	res = __gong__appendInstances(res, stage.ENUM_VALUEs)

	res = __gong__appendInstances(res, stage.RELATION_GROUPs)

	res = __gong__appendInstances(res, stage.RELATION_GROUP_TYPEs)

	res = __gong__appendInstances(res, stage.REQ_IFs)

	res = __gong__appendInstances(res, stage.REQ_IF_CONTENTs)

	res = __gong__appendInstances(res, stage.REQ_IF_HEADERs)

	res = __gong__appendInstances(res, stage.REQ_IF_TOOL_EXTENSIONs)

	res = __gong__appendInstances(res, stage.SPECIFICATIONs)

	res = __gong__appendInstances(res, stage.SPECIFICATION_TYPEs)

	res = __gong__appendInstances(res, stage.SPEC_HIERARCHYs)

	res = __gong__appendInstances(res, stage.SPEC_OBJECTs)

	res = __gong__appendInstances(res, stage.SPEC_OBJECT_TYPEs)

	res = __gong__appendInstances(res, stage.SPEC_RELATIONs)

	res = __gong__appendInstances(res, stage.SPEC_RELATION_TYPEs)

	res = __gong__appendInstances(res, stage.XHTML_CONTENTs)

	return
}

// insertion point per named struct
func (alternative_id *ALTERNATIVE_ID) GongCopy() GongstructIF {
	newInstance := new(ALTERNATIVE_ID)
	alternative_id.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_BOOLEAN)
	attribute_definition_boolean.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_DATE)
	attribute_definition_date.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_ENUMERATION)
	attribute_definition_enumeration.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_INTEGER)
	attribute_definition_integer.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_REAL)
	attribute_definition_real.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_STRING)
	attribute_definition_string.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_XHTML)
	attribute_definition_xhtml.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_VALUE_BOOLEAN)
	attribute_value_boolean.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_VALUE_DATE)
	attribute_value_date.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_VALUE_ENUMERATION)
	attribute_value_enumeration.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_VALUE_INTEGER)
	attribute_value_integer.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_VALUE_REAL)
	attribute_value_real.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_VALUE_STRING)
	attribute_value_string.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_VALUE_XHTML)
	attribute_value_xhtml.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongCopy() GongstructIF {
	newInstance := new(A_ALTERNATIVE_ID)
	a_alternative_id.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongCopy() GongstructIF {
	newInstance := new(A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	a_attribute_definition_boolean_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongCopy() GongstructIF {
	newInstance := new(A_ATTRIBUTE_DEFINITION_DATE_REF)
	a_attribute_definition_date_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongCopy() GongstructIF {
	newInstance := new(A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	a_attribute_definition_enumeration_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongCopy() GongstructIF {
	newInstance := new(A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	a_attribute_definition_integer_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongCopy() GongstructIF {
	newInstance := new(A_ATTRIBUTE_DEFINITION_REAL_REF)
	a_attribute_definition_real_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongCopy() GongstructIF {
	newInstance := new(A_ATTRIBUTE_DEFINITION_STRING_REF)
	a_attribute_definition_string_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongCopy() GongstructIF {
	newInstance := new(A_ATTRIBUTE_DEFINITION_XHTML_REF)
	a_attribute_definition_xhtml_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongCopy() GongstructIF {
	newInstance := new(A_ATTRIBUTE_VALUE_BOOLEAN)
	a_attribute_value_boolean.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongCopy() GongstructIF {
	newInstance := new(A_ATTRIBUTE_VALUE_DATE)
	a_attribute_value_date.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongCopy() GongstructIF {
	newInstance := new(A_ATTRIBUTE_VALUE_ENUMERATION)
	a_attribute_value_enumeration.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongCopy() GongstructIF {
	newInstance := new(A_ATTRIBUTE_VALUE_INTEGER)
	a_attribute_value_integer.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongCopy() GongstructIF {
	newInstance := new(A_ATTRIBUTE_VALUE_REAL)
	a_attribute_value_real.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongCopy() GongstructIF {
	newInstance := new(A_ATTRIBUTE_VALUE_STRING)
	a_attribute_value_string.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongCopy() GongstructIF {
	newInstance := new(A_ATTRIBUTE_VALUE_XHTML)
	a_attribute_value_xhtml.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongCopy() GongstructIF {
	newInstance := new(A_ATTRIBUTE_VALUE_XHTML_1)
	a_attribute_value_xhtml_1.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_children *A_CHILDREN) GongCopy() GongstructIF {
	newInstance := new(A_CHILDREN)
	a_children.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_core_content *A_CORE_CONTENT) GongCopy() GongstructIF {
	newInstance := new(A_CORE_CONTENT)
	a_core_content.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_datatypes *A_DATATYPES) GongCopy() GongstructIF {
	newInstance := new(A_DATATYPES)
	a_datatypes.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongCopy() GongstructIF {
	newInstance := new(A_DATATYPE_DEFINITION_BOOLEAN_REF)
	a_datatype_definition_boolean_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongCopy() GongstructIF {
	newInstance := new(A_DATATYPE_DEFINITION_DATE_REF)
	a_datatype_definition_date_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongCopy() GongstructIF {
	newInstance := new(A_DATATYPE_DEFINITION_ENUMERATION_REF)
	a_datatype_definition_enumeration_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongCopy() GongstructIF {
	newInstance := new(A_DATATYPE_DEFINITION_INTEGER_REF)
	a_datatype_definition_integer_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongCopy() GongstructIF {
	newInstance := new(A_DATATYPE_DEFINITION_REAL_REF)
	a_datatype_definition_real_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongCopy() GongstructIF {
	newInstance := new(A_DATATYPE_DEFINITION_STRING_REF)
	a_datatype_definition_string_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongCopy() GongstructIF {
	newInstance := new(A_DATATYPE_DEFINITION_XHTML_REF)
	a_datatype_definition_xhtml_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_editable_atts *A_EDITABLE_ATTS) GongCopy() GongstructIF {
	newInstance := new(A_EDITABLE_ATTS)
	a_editable_atts.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongCopy() GongstructIF {
	newInstance := new(A_ENUM_VALUE_REF)
	a_enum_value_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_object *A_OBJECT) GongCopy() GongstructIF {
	newInstance := new(A_OBJECT)
	a_object.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_properties *A_PROPERTIES) GongCopy() GongstructIF {
	newInstance := new(A_PROPERTIES)
	a_properties.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongCopy() GongstructIF {
	newInstance := new(A_RELATION_GROUP_TYPE_REF)
	a_relation_group_type_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_source_1 *A_SOURCE_1) GongCopy() GongstructIF {
	newInstance := new(A_SOURCE_1)
	a_source_1.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongCopy() GongstructIF {
	newInstance := new(A_SOURCE_SPECIFICATION_1)
	a_source_specification_1.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_specifications *A_SPECIFICATIONS) GongCopy() GongstructIF {
	newInstance := new(A_SPECIFICATIONS)
	a_specifications.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongCopy() GongstructIF {
	newInstance := new(A_SPECIFICATION_TYPE_REF)
	a_specification_type_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_specified_values *A_SPECIFIED_VALUES) GongCopy() GongstructIF {
	newInstance := new(A_SPECIFIED_VALUES)
	a_specified_values.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongCopy() GongstructIF {
	newInstance := new(A_SPEC_ATTRIBUTES)
	a_spec_attributes.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_spec_objects *A_SPEC_OBJECTS) GongCopy() GongstructIF {
	newInstance := new(A_SPEC_OBJECTS)
	a_spec_objects.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongCopy() GongstructIF {
	newInstance := new(A_SPEC_OBJECT_TYPE_REF)
	a_spec_object_type_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_spec_relations *A_SPEC_RELATIONS) GongCopy() GongstructIF {
	newInstance := new(A_SPEC_RELATIONS)
	a_spec_relations.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongCopy() GongstructIF {
	newInstance := new(A_SPEC_RELATION_GROUPS)
	a_spec_relation_groups.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongCopy() GongstructIF {
	newInstance := new(A_SPEC_RELATION_REF)
	a_spec_relation_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongCopy() GongstructIF {
	newInstance := new(A_SPEC_RELATION_TYPE_REF)
	a_spec_relation_type_ref.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_spec_types *A_SPEC_TYPES) GongCopy() GongstructIF {
	newInstance := new(A_SPEC_TYPES)
	a_spec_types.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_the_header *A_THE_HEADER) GongCopy() GongstructIF {
	newInstance := new(A_THE_HEADER)
	a_the_header.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongCopy() GongstructIF {
	newInstance := new(A_TOOL_EXTENSIONS)
	a_tool_extensions.GongCopyBasicFields(newInstance)
	return newInstance
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongCopy() GongstructIF {
	newInstance := new(DATATYPE_DEFINITION_BOOLEAN)
	datatype_definition_boolean.GongCopyBasicFields(newInstance)
	return newInstance
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongCopy() GongstructIF {
	newInstance := new(DATATYPE_DEFINITION_DATE)
	datatype_definition_date.GongCopyBasicFields(newInstance)
	return newInstance
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongCopy() GongstructIF {
	newInstance := new(DATATYPE_DEFINITION_ENUMERATION)
	datatype_definition_enumeration.GongCopyBasicFields(newInstance)
	return newInstance
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongCopy() GongstructIF {
	newInstance := new(DATATYPE_DEFINITION_INTEGER)
	datatype_definition_integer.GongCopyBasicFields(newInstance)
	return newInstance
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongCopy() GongstructIF {
	newInstance := new(DATATYPE_DEFINITION_REAL)
	datatype_definition_real.GongCopyBasicFields(newInstance)
	return newInstance
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongCopy() GongstructIF {
	newInstance := new(DATATYPE_DEFINITION_STRING)
	datatype_definition_string.GongCopyBasicFields(newInstance)
	return newInstance
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongCopy() GongstructIF {
	newInstance := new(DATATYPE_DEFINITION_XHTML)
	datatype_definition_xhtml.GongCopyBasicFields(newInstance)
	return newInstance
}

func (embedded_value *EMBEDDED_VALUE) GongCopy() GongstructIF {
	newInstance := new(EMBEDDED_VALUE)
	embedded_value.GongCopyBasicFields(newInstance)
	return newInstance
}

func (enum_value *ENUM_VALUE) GongCopy() GongstructIF {
	newInstance := new(ENUM_VALUE)
	enum_value.GongCopyBasicFields(newInstance)
	return newInstance
}

func (relation_group *RELATION_GROUP) GongCopy() GongstructIF {
	newInstance := new(RELATION_GROUP)
	relation_group.GongCopyBasicFields(newInstance)
	return newInstance
}

func (relation_group_type *RELATION_GROUP_TYPE) GongCopy() GongstructIF {
	newInstance := new(RELATION_GROUP_TYPE)
	relation_group_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (req_if *REQ_IF) GongCopy() GongstructIF {
	newInstance := new(REQ_IF)
	req_if.GongCopyBasicFields(newInstance)
	return newInstance
}

func (req_if_content *REQ_IF_CONTENT) GongCopy() GongstructIF {
	newInstance := new(REQ_IF_CONTENT)
	req_if_content.GongCopyBasicFields(newInstance)
	return newInstance
}

func (req_if_header *REQ_IF_HEADER) GongCopy() GongstructIF {
	newInstance := new(REQ_IF_HEADER)
	req_if_header.GongCopyBasicFields(newInstance)
	return newInstance
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongCopy() GongstructIF {
	newInstance := new(REQ_IF_TOOL_EXTENSION)
	req_if_tool_extension.GongCopyBasicFields(newInstance)
	return newInstance
}

func (specification *SPECIFICATION) GongCopy() GongstructIF {
	newInstance := new(SPECIFICATION)
	specification.GongCopyBasicFields(newInstance)
	return newInstance
}

func (specification_type *SPECIFICATION_TYPE) GongCopy() GongstructIF {
	newInstance := new(SPECIFICATION_TYPE)
	specification_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (spec_hierarchy *SPEC_HIERARCHY) GongCopy() GongstructIF {
	newInstance := new(SPEC_HIERARCHY)
	spec_hierarchy.GongCopyBasicFields(newInstance)
	return newInstance
}

func (spec_object *SPEC_OBJECT) GongCopy() GongstructIF {
	newInstance := new(SPEC_OBJECT)
	spec_object.GongCopyBasicFields(newInstance)
	return newInstance
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongCopy() GongstructIF {
	newInstance := new(SPEC_OBJECT_TYPE)
	spec_object_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (spec_relation *SPEC_RELATION) GongCopy() GongstructIF {
	newInstance := new(SPEC_RELATION)
	spec_relation.GongCopyBasicFields(newInstance)
	return newInstance
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongCopy() GongstructIF {
	newInstance := new(SPEC_RELATION_TYPE)
	spec_relation_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (xhtml_content *XHTML_CONTENT) GongCopy() GongstructIF {
	newInstance := new(XHTML_CONTENT)
	xhtml_content.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (alternative_id *ALTERNATIVE_ID) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, alternative_id)
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attribute_definition_boolean)
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attribute_definition_date)
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attribute_definition_enumeration)
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attribute_definition_integer)
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attribute_definition_real)
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attribute_definition_string)
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attribute_definition_xhtml)
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attribute_value_boolean)
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attribute_value_date)
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attribute_value_enumeration)
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attribute_value_integer)
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attribute_value_real)
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attribute_value_string)
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attribute_value_xhtml)
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_alternative_id)
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_attribute_definition_boolean_ref)
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_attribute_definition_date_ref)
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_attribute_definition_enumeration_ref)
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_attribute_definition_integer_ref)
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_attribute_definition_real_ref)
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_attribute_definition_string_ref)
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_attribute_definition_xhtml_ref)
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_attribute_value_boolean)
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_attribute_value_date)
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_attribute_value_enumeration)
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_attribute_value_integer)
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_attribute_value_real)
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_attribute_value_string)
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_attribute_value_xhtml)
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_attribute_value_xhtml_1)
}

func (a_children *A_CHILDREN) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_children)
}

func (a_core_content *A_CORE_CONTENT) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_core_content)
}

func (a_datatypes *A_DATATYPES) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_datatypes)
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_datatype_definition_boolean_ref)
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_datatype_definition_date_ref)
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_datatype_definition_enumeration_ref)
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_datatype_definition_integer_ref)
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_datatype_definition_real_ref)
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_datatype_definition_string_ref)
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_datatype_definition_xhtml_ref)
}

func (a_editable_atts *A_EDITABLE_ATTS) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_editable_atts)
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_enum_value_ref)
}

func (a_object *A_OBJECT) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_object)
}

func (a_properties *A_PROPERTIES) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_properties)
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_relation_group_type_ref)
}

func (a_source_1 *A_SOURCE_1) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_source_1)
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_source_specification_1)
}

func (a_specifications *A_SPECIFICATIONS) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_specifications)
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_specification_type_ref)
}

func (a_specified_values *A_SPECIFIED_VALUES) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_specified_values)
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_spec_attributes)
}

func (a_spec_objects *A_SPEC_OBJECTS) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_spec_objects)
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_spec_object_type_ref)
}

func (a_spec_relations *A_SPEC_RELATIONS) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_spec_relations)
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_spec_relation_groups)
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_spec_relation_ref)
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_spec_relation_type_ref)
}

func (a_spec_types *A_SPEC_TYPES) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_spec_types)
}

func (a_the_header *A_THE_HEADER) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_the_header)
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_tool_extensions)
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, datatype_definition_boolean)
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, datatype_definition_date)
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, datatype_definition_enumeration)
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, datatype_definition_integer)
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, datatype_definition_real)
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, datatype_definition_string)
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, datatype_definition_xhtml)
}

func (embedded_value *EMBEDDED_VALUE) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, embedded_value)
}

func (enum_value *ENUM_VALUE) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, enum_value)
}

func (relation_group *RELATION_GROUP) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, relation_group)
}

func (relation_group_type *RELATION_GROUP_TYPE) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, relation_group_type)
}

func (req_if *REQ_IF) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, req_if)
}

func (req_if_content *REQ_IF_CONTENT) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, req_if_content)
}

func (req_if_header *REQ_IF_HEADER) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, req_if_header)
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, req_if_tool_extension)
}

func (specification *SPECIFICATION) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, specification)
}

func (specification_type *SPECIFICATION_TYPE) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, specification_type)
}

func (spec_hierarchy *SPEC_HIERARCHY) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, spec_hierarchy)
}

func (spec_object *SPEC_OBJECT) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, spec_object)
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, spec_object_type)
}

func (spec_relation *SPEC_RELATION) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, spec_relation)
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, spec_relation_type)
}

func (xhtml_content *XHTML_CONTENT) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, xhtml_content)
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	computeCommitsForType(
		stage,
		stage.ALTERNATIVE_IDs,
		stage.ALTERNATIVE_ID_stagedOrder,
		stage.ALTERNATIVE_IDs_reference,
		&stage.ALTERNATIVE_IDs_referenceOrder,
		stage.ALTERNATIVE_IDs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ATTRIBUTE_DEFINITION_BOOLEANs,
		stage.ATTRIBUTE_DEFINITION_BOOLEAN_stagedOrder,
		stage.ATTRIBUTE_DEFINITION_BOOLEANs_reference,
		&stage.ATTRIBUTE_DEFINITION_BOOLEANs_referenceOrder,
		stage.ATTRIBUTE_DEFINITION_BOOLEANs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ATTRIBUTE_DEFINITION_DATEs,
		stage.ATTRIBUTE_DEFINITION_DATE_stagedOrder,
		stage.ATTRIBUTE_DEFINITION_DATEs_reference,
		&stage.ATTRIBUTE_DEFINITION_DATEs_referenceOrder,
		stage.ATTRIBUTE_DEFINITION_DATEs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ATTRIBUTE_DEFINITION_ENUMERATIONs,
		stage.ATTRIBUTE_DEFINITION_ENUMERATION_stagedOrder,
		stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_reference,
		&stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_referenceOrder,
		stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ATTRIBUTE_DEFINITION_INTEGERs,
		stage.ATTRIBUTE_DEFINITION_INTEGER_stagedOrder,
		stage.ATTRIBUTE_DEFINITION_INTEGERs_reference,
		&stage.ATTRIBUTE_DEFINITION_INTEGERs_referenceOrder,
		stage.ATTRIBUTE_DEFINITION_INTEGERs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ATTRIBUTE_DEFINITION_REALs,
		stage.ATTRIBUTE_DEFINITION_REAL_stagedOrder,
		stage.ATTRIBUTE_DEFINITION_REALs_reference,
		&stage.ATTRIBUTE_DEFINITION_REALs_referenceOrder,
		stage.ATTRIBUTE_DEFINITION_REALs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ATTRIBUTE_DEFINITION_STRINGs,
		stage.ATTRIBUTE_DEFINITION_STRING_stagedOrder,
		stage.ATTRIBUTE_DEFINITION_STRINGs_reference,
		&stage.ATTRIBUTE_DEFINITION_STRINGs_referenceOrder,
		stage.ATTRIBUTE_DEFINITION_STRINGs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ATTRIBUTE_DEFINITION_XHTMLs,
		stage.ATTRIBUTE_DEFINITION_XHTML_stagedOrder,
		stage.ATTRIBUTE_DEFINITION_XHTMLs_reference,
		&stage.ATTRIBUTE_DEFINITION_XHTMLs_referenceOrder,
		stage.ATTRIBUTE_DEFINITION_XHTMLs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ATTRIBUTE_VALUE_BOOLEANs,
		stage.ATTRIBUTE_VALUE_BOOLEAN_stagedOrder,
		stage.ATTRIBUTE_VALUE_BOOLEANs_reference,
		&stage.ATTRIBUTE_VALUE_BOOLEANs_referenceOrder,
		stage.ATTRIBUTE_VALUE_BOOLEANs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ATTRIBUTE_VALUE_DATEs,
		stage.ATTRIBUTE_VALUE_DATE_stagedOrder,
		stage.ATTRIBUTE_VALUE_DATEs_reference,
		&stage.ATTRIBUTE_VALUE_DATEs_referenceOrder,
		stage.ATTRIBUTE_VALUE_DATEs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ATTRIBUTE_VALUE_ENUMERATIONs,
		stage.ATTRIBUTE_VALUE_ENUMERATION_stagedOrder,
		stage.ATTRIBUTE_VALUE_ENUMERATIONs_reference,
		&stage.ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder,
		stage.ATTRIBUTE_VALUE_ENUMERATIONs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ATTRIBUTE_VALUE_INTEGERs,
		stage.ATTRIBUTE_VALUE_INTEGER_stagedOrder,
		stage.ATTRIBUTE_VALUE_INTEGERs_reference,
		&stage.ATTRIBUTE_VALUE_INTEGERs_referenceOrder,
		stage.ATTRIBUTE_VALUE_INTEGERs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ATTRIBUTE_VALUE_REALs,
		stage.ATTRIBUTE_VALUE_REAL_stagedOrder,
		stage.ATTRIBUTE_VALUE_REALs_reference,
		&stage.ATTRIBUTE_VALUE_REALs_referenceOrder,
		stage.ATTRIBUTE_VALUE_REALs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ATTRIBUTE_VALUE_STRINGs,
		stage.ATTRIBUTE_VALUE_STRING_stagedOrder,
		stage.ATTRIBUTE_VALUE_STRINGs_reference,
		&stage.ATTRIBUTE_VALUE_STRINGs_referenceOrder,
		stage.ATTRIBUTE_VALUE_STRINGs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ATTRIBUTE_VALUE_XHTMLs,
		stage.ATTRIBUTE_VALUE_XHTML_stagedOrder,
		stage.ATTRIBUTE_VALUE_XHTMLs_reference,
		&stage.ATTRIBUTE_VALUE_XHTMLs_referenceOrder,
		stage.ATTRIBUTE_VALUE_XHTMLs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ALTERNATIVE_IDs,
		stage.A_ALTERNATIVE_ID_stagedOrder,
		stage.A_ALTERNATIVE_IDs_reference,
		&stage.A_ALTERNATIVE_IDs_referenceOrder,
		stage.A_ALTERNATIVE_IDs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs,
		stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_stagedOrder,
		stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_reference,
		&stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_referenceOrder,
		stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ATTRIBUTE_DEFINITION_DATE_REFs,
		stage.A_ATTRIBUTE_DEFINITION_DATE_REF_stagedOrder,
		stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_reference,
		&stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_referenceOrder,
		stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs,
		stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_stagedOrder,
		stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_reference,
		&stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_referenceOrder,
		stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs,
		stage.A_ATTRIBUTE_DEFINITION_INTEGER_REF_stagedOrder,
		stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_reference,
		&stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_referenceOrder,
		stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ATTRIBUTE_DEFINITION_REAL_REFs,
		stage.A_ATTRIBUTE_DEFINITION_REAL_REF_stagedOrder,
		stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_reference,
		&stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_referenceOrder,
		stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ATTRIBUTE_DEFINITION_STRING_REFs,
		stage.A_ATTRIBUTE_DEFINITION_STRING_REF_stagedOrder,
		stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_reference,
		&stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_referenceOrder,
		stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs,
		stage.A_ATTRIBUTE_DEFINITION_XHTML_REF_stagedOrder,
		stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_reference,
		&stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_referenceOrder,
		stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ATTRIBUTE_VALUE_BOOLEANs,
		stage.A_ATTRIBUTE_VALUE_BOOLEAN_stagedOrder,
		stage.A_ATTRIBUTE_VALUE_BOOLEANs_reference,
		&stage.A_ATTRIBUTE_VALUE_BOOLEANs_referenceOrder,
		stage.A_ATTRIBUTE_VALUE_BOOLEANs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ATTRIBUTE_VALUE_DATEs,
		stage.A_ATTRIBUTE_VALUE_DATE_stagedOrder,
		stage.A_ATTRIBUTE_VALUE_DATEs_reference,
		&stage.A_ATTRIBUTE_VALUE_DATEs_referenceOrder,
		stage.A_ATTRIBUTE_VALUE_DATEs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ATTRIBUTE_VALUE_ENUMERATIONs,
		stage.A_ATTRIBUTE_VALUE_ENUMERATION_stagedOrder,
		stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_reference,
		&stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder,
		stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ATTRIBUTE_VALUE_INTEGERs,
		stage.A_ATTRIBUTE_VALUE_INTEGER_stagedOrder,
		stage.A_ATTRIBUTE_VALUE_INTEGERs_reference,
		&stage.A_ATTRIBUTE_VALUE_INTEGERs_referenceOrder,
		stage.A_ATTRIBUTE_VALUE_INTEGERs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ATTRIBUTE_VALUE_REALs,
		stage.A_ATTRIBUTE_VALUE_REAL_stagedOrder,
		stage.A_ATTRIBUTE_VALUE_REALs_reference,
		&stage.A_ATTRIBUTE_VALUE_REALs_referenceOrder,
		stage.A_ATTRIBUTE_VALUE_REALs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ATTRIBUTE_VALUE_STRINGs,
		stage.A_ATTRIBUTE_VALUE_STRING_stagedOrder,
		stage.A_ATTRIBUTE_VALUE_STRINGs_reference,
		&stage.A_ATTRIBUTE_VALUE_STRINGs_referenceOrder,
		stage.A_ATTRIBUTE_VALUE_STRINGs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ATTRIBUTE_VALUE_XHTMLs,
		stage.A_ATTRIBUTE_VALUE_XHTML_stagedOrder,
		stage.A_ATTRIBUTE_VALUE_XHTMLs_reference,
		&stage.A_ATTRIBUTE_VALUE_XHTMLs_referenceOrder,
		stage.A_ATTRIBUTE_VALUE_XHTMLs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ATTRIBUTE_VALUE_XHTML_1s,
		stage.A_ATTRIBUTE_VALUE_XHTML_1_stagedOrder,
		stage.A_ATTRIBUTE_VALUE_XHTML_1s_reference,
		&stage.A_ATTRIBUTE_VALUE_XHTML_1s_referenceOrder,
		stage.A_ATTRIBUTE_VALUE_XHTML_1s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_CHILDRENs,
		stage.A_CHILDREN_stagedOrder,
		stage.A_CHILDRENs_reference,
		&stage.A_CHILDRENs_referenceOrder,
		stage.A_CHILDRENs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_CORE_CONTENTs,
		stage.A_CORE_CONTENT_stagedOrder,
		stage.A_CORE_CONTENTs_reference,
		&stage.A_CORE_CONTENTs_referenceOrder,
		stage.A_CORE_CONTENTs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_DATATYPESs,
		stage.A_DATATYPES_stagedOrder,
		stage.A_DATATYPESs_reference,
		&stage.A_DATATYPESs_referenceOrder,
		stage.A_DATATYPESs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs,
		stage.A_DATATYPE_DEFINITION_BOOLEAN_REF_stagedOrder,
		stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_reference,
		&stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_referenceOrder,
		stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_DATATYPE_DEFINITION_DATE_REFs,
		stage.A_DATATYPE_DEFINITION_DATE_REF_stagedOrder,
		stage.A_DATATYPE_DEFINITION_DATE_REFs_reference,
		&stage.A_DATATYPE_DEFINITION_DATE_REFs_referenceOrder,
		stage.A_DATATYPE_DEFINITION_DATE_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs,
		stage.A_DATATYPE_DEFINITION_ENUMERATION_REF_stagedOrder,
		stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_reference,
		&stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_referenceOrder,
		stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_DATATYPE_DEFINITION_INTEGER_REFs,
		stage.A_DATATYPE_DEFINITION_INTEGER_REF_stagedOrder,
		stage.A_DATATYPE_DEFINITION_INTEGER_REFs_reference,
		&stage.A_DATATYPE_DEFINITION_INTEGER_REFs_referenceOrder,
		stage.A_DATATYPE_DEFINITION_INTEGER_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_DATATYPE_DEFINITION_REAL_REFs,
		stage.A_DATATYPE_DEFINITION_REAL_REF_stagedOrder,
		stage.A_DATATYPE_DEFINITION_REAL_REFs_reference,
		&stage.A_DATATYPE_DEFINITION_REAL_REFs_referenceOrder,
		stage.A_DATATYPE_DEFINITION_REAL_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_DATATYPE_DEFINITION_STRING_REFs,
		stage.A_DATATYPE_DEFINITION_STRING_REF_stagedOrder,
		stage.A_DATATYPE_DEFINITION_STRING_REFs_reference,
		&stage.A_DATATYPE_DEFINITION_STRING_REFs_referenceOrder,
		stage.A_DATATYPE_DEFINITION_STRING_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_DATATYPE_DEFINITION_XHTML_REFs,
		stage.A_DATATYPE_DEFINITION_XHTML_REF_stagedOrder,
		stage.A_DATATYPE_DEFINITION_XHTML_REFs_reference,
		&stage.A_DATATYPE_DEFINITION_XHTML_REFs_referenceOrder,
		stage.A_DATATYPE_DEFINITION_XHTML_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_EDITABLE_ATTSs,
		stage.A_EDITABLE_ATTS_stagedOrder,
		stage.A_EDITABLE_ATTSs_reference,
		&stage.A_EDITABLE_ATTSs_referenceOrder,
		stage.A_EDITABLE_ATTSs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_ENUM_VALUE_REFs,
		stage.A_ENUM_VALUE_REF_stagedOrder,
		stage.A_ENUM_VALUE_REFs_reference,
		&stage.A_ENUM_VALUE_REFs_referenceOrder,
		stage.A_ENUM_VALUE_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_OBJECTs,
		stage.A_OBJECT_stagedOrder,
		stage.A_OBJECTs_reference,
		&stage.A_OBJECTs_referenceOrder,
		stage.A_OBJECTs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_PROPERTIESs,
		stage.A_PROPERTIES_stagedOrder,
		stage.A_PROPERTIESs_reference,
		&stage.A_PROPERTIESs_referenceOrder,
		stage.A_PROPERTIESs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_RELATION_GROUP_TYPE_REFs,
		stage.A_RELATION_GROUP_TYPE_REF_stagedOrder,
		stage.A_RELATION_GROUP_TYPE_REFs_reference,
		&stage.A_RELATION_GROUP_TYPE_REFs_referenceOrder,
		stage.A_RELATION_GROUP_TYPE_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_SOURCE_1s,
		stage.A_SOURCE_1_stagedOrder,
		stage.A_SOURCE_1s_reference,
		&stage.A_SOURCE_1s_referenceOrder,
		stage.A_SOURCE_1s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_SOURCE_SPECIFICATION_1s,
		stage.A_SOURCE_SPECIFICATION_1_stagedOrder,
		stage.A_SOURCE_SPECIFICATION_1s_reference,
		&stage.A_SOURCE_SPECIFICATION_1s_referenceOrder,
		stage.A_SOURCE_SPECIFICATION_1s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_SPECIFICATIONSs,
		stage.A_SPECIFICATIONS_stagedOrder,
		stage.A_SPECIFICATIONSs_reference,
		&stage.A_SPECIFICATIONSs_referenceOrder,
		stage.A_SPECIFICATIONSs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_SPECIFICATION_TYPE_REFs,
		stage.A_SPECIFICATION_TYPE_REF_stagedOrder,
		stage.A_SPECIFICATION_TYPE_REFs_reference,
		&stage.A_SPECIFICATION_TYPE_REFs_referenceOrder,
		stage.A_SPECIFICATION_TYPE_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_SPECIFIED_VALUESs,
		stage.A_SPECIFIED_VALUES_stagedOrder,
		stage.A_SPECIFIED_VALUESs_reference,
		&stage.A_SPECIFIED_VALUESs_referenceOrder,
		stage.A_SPECIFIED_VALUESs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_SPEC_ATTRIBUTESs,
		stage.A_SPEC_ATTRIBUTES_stagedOrder,
		stage.A_SPEC_ATTRIBUTESs_reference,
		&stage.A_SPEC_ATTRIBUTESs_referenceOrder,
		stage.A_SPEC_ATTRIBUTESs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_SPEC_OBJECTSs,
		stage.A_SPEC_OBJECTS_stagedOrder,
		stage.A_SPEC_OBJECTSs_reference,
		&stage.A_SPEC_OBJECTSs_referenceOrder,
		stage.A_SPEC_OBJECTSs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_SPEC_OBJECT_TYPE_REFs,
		stage.A_SPEC_OBJECT_TYPE_REF_stagedOrder,
		stage.A_SPEC_OBJECT_TYPE_REFs_reference,
		&stage.A_SPEC_OBJECT_TYPE_REFs_referenceOrder,
		stage.A_SPEC_OBJECT_TYPE_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_SPEC_RELATIONSs,
		stage.A_SPEC_RELATIONS_stagedOrder,
		stage.A_SPEC_RELATIONSs_reference,
		&stage.A_SPEC_RELATIONSs_referenceOrder,
		stage.A_SPEC_RELATIONSs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_SPEC_RELATION_GROUPSs,
		stage.A_SPEC_RELATION_GROUPS_stagedOrder,
		stage.A_SPEC_RELATION_GROUPSs_reference,
		&stage.A_SPEC_RELATION_GROUPSs_referenceOrder,
		stage.A_SPEC_RELATION_GROUPSs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_SPEC_RELATION_REFs,
		stage.A_SPEC_RELATION_REF_stagedOrder,
		stage.A_SPEC_RELATION_REFs_reference,
		&stage.A_SPEC_RELATION_REFs_referenceOrder,
		stage.A_SPEC_RELATION_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_SPEC_RELATION_TYPE_REFs,
		stage.A_SPEC_RELATION_TYPE_REF_stagedOrder,
		stage.A_SPEC_RELATION_TYPE_REFs_reference,
		&stage.A_SPEC_RELATION_TYPE_REFs_referenceOrder,
		stage.A_SPEC_RELATION_TYPE_REFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_SPEC_TYPESs,
		stage.A_SPEC_TYPES_stagedOrder,
		stage.A_SPEC_TYPESs_reference,
		&stage.A_SPEC_TYPESs_referenceOrder,
		stage.A_SPEC_TYPESs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_THE_HEADERs,
		stage.A_THE_HEADER_stagedOrder,
		stage.A_THE_HEADERs_reference,
		&stage.A_THE_HEADERs_referenceOrder,
		stage.A_THE_HEADERs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_TOOL_EXTENSIONSs,
		stage.A_TOOL_EXTENSIONS_stagedOrder,
		stage.A_TOOL_EXTENSIONSs_reference,
		&stage.A_TOOL_EXTENSIONSs_referenceOrder,
		stage.A_TOOL_EXTENSIONSs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DATATYPE_DEFINITION_BOOLEANs,
		stage.DATATYPE_DEFINITION_BOOLEAN_stagedOrder,
		stage.DATATYPE_DEFINITION_BOOLEANs_reference,
		&stage.DATATYPE_DEFINITION_BOOLEANs_referenceOrder,
		stage.DATATYPE_DEFINITION_BOOLEANs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DATATYPE_DEFINITION_DATEs,
		stage.DATATYPE_DEFINITION_DATE_stagedOrder,
		stage.DATATYPE_DEFINITION_DATEs_reference,
		&stage.DATATYPE_DEFINITION_DATEs_referenceOrder,
		stage.DATATYPE_DEFINITION_DATEs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DATATYPE_DEFINITION_ENUMERATIONs,
		stage.DATATYPE_DEFINITION_ENUMERATION_stagedOrder,
		stage.DATATYPE_DEFINITION_ENUMERATIONs_reference,
		&stage.DATATYPE_DEFINITION_ENUMERATIONs_referenceOrder,
		stage.DATATYPE_DEFINITION_ENUMERATIONs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DATATYPE_DEFINITION_INTEGERs,
		stage.DATATYPE_DEFINITION_INTEGER_stagedOrder,
		stage.DATATYPE_DEFINITION_INTEGERs_reference,
		&stage.DATATYPE_DEFINITION_INTEGERs_referenceOrder,
		stage.DATATYPE_DEFINITION_INTEGERs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DATATYPE_DEFINITION_REALs,
		stage.DATATYPE_DEFINITION_REAL_stagedOrder,
		stage.DATATYPE_DEFINITION_REALs_reference,
		&stage.DATATYPE_DEFINITION_REALs_referenceOrder,
		stage.DATATYPE_DEFINITION_REALs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DATATYPE_DEFINITION_STRINGs,
		stage.DATATYPE_DEFINITION_STRING_stagedOrder,
		stage.DATATYPE_DEFINITION_STRINGs_reference,
		&stage.DATATYPE_DEFINITION_STRINGs_referenceOrder,
		stage.DATATYPE_DEFINITION_STRINGs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.DATATYPE_DEFINITION_XHTMLs,
		stage.DATATYPE_DEFINITION_XHTML_stagedOrder,
		stage.DATATYPE_DEFINITION_XHTMLs_reference,
		&stage.DATATYPE_DEFINITION_XHTMLs_referenceOrder,
		stage.DATATYPE_DEFINITION_XHTMLs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.EMBEDDED_VALUEs,
		stage.EMBEDDED_VALUE_stagedOrder,
		stage.EMBEDDED_VALUEs_reference,
		&stage.EMBEDDED_VALUEs_referenceOrder,
		stage.EMBEDDED_VALUEs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.ENUM_VALUEs,
		stage.ENUM_VALUE_stagedOrder,
		stage.ENUM_VALUEs_reference,
		&stage.ENUM_VALUEs_referenceOrder,
		stage.ENUM_VALUEs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RELATION_GROUPs,
		stage.RELATION_GROUP_stagedOrder,
		stage.RELATION_GROUPs_reference,
		&stage.RELATION_GROUPs_referenceOrder,
		stage.RELATION_GROUPs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.RELATION_GROUP_TYPEs,
		stage.RELATION_GROUP_TYPE_stagedOrder,
		stage.RELATION_GROUP_TYPEs_reference,
		&stage.RELATION_GROUP_TYPEs_referenceOrder,
		stage.RELATION_GROUP_TYPEs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.REQ_IFs,
		stage.REQ_IF_stagedOrder,
		stage.REQ_IFs_reference,
		&stage.REQ_IFs_referenceOrder,
		stage.REQ_IFs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.REQ_IF_CONTENTs,
		stage.REQ_IF_CONTENT_stagedOrder,
		stage.REQ_IF_CONTENTs_reference,
		&stage.REQ_IF_CONTENTs_referenceOrder,
		stage.REQ_IF_CONTENTs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.REQ_IF_HEADERs,
		stage.REQ_IF_HEADER_stagedOrder,
		stage.REQ_IF_HEADERs_reference,
		&stage.REQ_IF_HEADERs_referenceOrder,
		stage.REQ_IF_HEADERs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.REQ_IF_TOOL_EXTENSIONs,
		stage.REQ_IF_TOOL_EXTENSION_stagedOrder,
		stage.REQ_IF_TOOL_EXTENSIONs_reference,
		&stage.REQ_IF_TOOL_EXTENSIONs_referenceOrder,
		stage.REQ_IF_TOOL_EXTENSIONs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SPECIFICATIONs,
		stage.SPECIFICATION_stagedOrder,
		stage.SPECIFICATIONs_reference,
		&stage.SPECIFICATIONs_referenceOrder,
		stage.SPECIFICATIONs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SPECIFICATION_TYPEs,
		stage.SPECIFICATION_TYPE_stagedOrder,
		stage.SPECIFICATION_TYPEs_reference,
		&stage.SPECIFICATION_TYPEs_referenceOrder,
		stage.SPECIFICATION_TYPEs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SPEC_HIERARCHYs,
		stage.SPEC_HIERARCHY_stagedOrder,
		stage.SPEC_HIERARCHYs_reference,
		&stage.SPEC_HIERARCHYs_referenceOrder,
		stage.SPEC_HIERARCHYs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SPEC_OBJECTs,
		stage.SPEC_OBJECT_stagedOrder,
		stage.SPEC_OBJECTs_reference,
		&stage.SPEC_OBJECTs_referenceOrder,
		stage.SPEC_OBJECTs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SPEC_OBJECT_TYPEs,
		stage.SPEC_OBJECT_TYPE_stagedOrder,
		stage.SPEC_OBJECT_TYPEs_reference,
		&stage.SPEC_OBJECT_TYPEs_referenceOrder,
		stage.SPEC_OBJECT_TYPEs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SPEC_RELATIONs,
		stage.SPEC_RELATION_stagedOrder,
		stage.SPEC_RELATIONs_reference,
		&stage.SPEC_RELATIONs_referenceOrder,
		stage.SPEC_RELATIONs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.SPEC_RELATION_TYPEs,
		stage.SPEC_RELATION_TYPE_stagedOrder,
		stage.SPEC_RELATION_TYPEs_reference,
		&stage.SPEC_RELATION_TYPEs_referenceOrder,
		stage.SPEC_RELATION_TYPEs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.XHTML_CONTENTs,
		stage.XHTML_CONTENT_stagedOrder,
		stage.XHTML_CONTENTs_reference,
		&stage.XHTML_CONTENTs_referenceOrder,
		stage.XHTML_CONTENTs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	__gong__computeReferencePass1(stage, stage.ALTERNATIVE_IDs, &stage.ALTERNATIVE_IDs_reference, &stage.ALTERNATIVE_IDs_referenceOrder, &stage.ALTERNATIVE_IDs_instance)

	__gong__computeReferencePass1(stage, stage.ATTRIBUTE_DEFINITION_BOOLEANs, &stage.ATTRIBUTE_DEFINITION_BOOLEANs_reference, &stage.ATTRIBUTE_DEFINITION_BOOLEANs_referenceOrder, &stage.ATTRIBUTE_DEFINITION_BOOLEANs_instance)

	__gong__computeReferencePass1(stage, stage.ATTRIBUTE_DEFINITION_DATEs, &stage.ATTRIBUTE_DEFINITION_DATEs_reference, &stage.ATTRIBUTE_DEFINITION_DATEs_referenceOrder, &stage.ATTRIBUTE_DEFINITION_DATEs_instance)

	__gong__computeReferencePass1(stage, stage.ATTRIBUTE_DEFINITION_ENUMERATIONs, &stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_reference, &stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_referenceOrder, &stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_instance)

	__gong__computeReferencePass1(stage, stage.ATTRIBUTE_DEFINITION_INTEGERs, &stage.ATTRIBUTE_DEFINITION_INTEGERs_reference, &stage.ATTRIBUTE_DEFINITION_INTEGERs_referenceOrder, &stage.ATTRIBUTE_DEFINITION_INTEGERs_instance)

	__gong__computeReferencePass1(stage, stage.ATTRIBUTE_DEFINITION_REALs, &stage.ATTRIBUTE_DEFINITION_REALs_reference, &stage.ATTRIBUTE_DEFINITION_REALs_referenceOrder, &stage.ATTRIBUTE_DEFINITION_REALs_instance)

	__gong__computeReferencePass1(stage, stage.ATTRIBUTE_DEFINITION_STRINGs, &stage.ATTRIBUTE_DEFINITION_STRINGs_reference, &stage.ATTRIBUTE_DEFINITION_STRINGs_referenceOrder, &stage.ATTRIBUTE_DEFINITION_STRINGs_instance)

	__gong__computeReferencePass1(stage, stage.ATTRIBUTE_DEFINITION_XHTMLs, &stage.ATTRIBUTE_DEFINITION_XHTMLs_reference, &stage.ATTRIBUTE_DEFINITION_XHTMLs_referenceOrder, &stage.ATTRIBUTE_DEFINITION_XHTMLs_instance)

	__gong__computeReferencePass1(stage, stage.ATTRIBUTE_VALUE_BOOLEANs, &stage.ATTRIBUTE_VALUE_BOOLEANs_reference, &stage.ATTRIBUTE_VALUE_BOOLEANs_referenceOrder, &stage.ATTRIBUTE_VALUE_BOOLEANs_instance)

	__gong__computeReferencePass1(stage, stage.ATTRIBUTE_VALUE_DATEs, &stage.ATTRIBUTE_VALUE_DATEs_reference, &stage.ATTRIBUTE_VALUE_DATEs_referenceOrder, &stage.ATTRIBUTE_VALUE_DATEs_instance)

	__gong__computeReferencePass1(stage, stage.ATTRIBUTE_VALUE_ENUMERATIONs, &stage.ATTRIBUTE_VALUE_ENUMERATIONs_reference, &stage.ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder, &stage.ATTRIBUTE_VALUE_ENUMERATIONs_instance)

	__gong__computeReferencePass1(stage, stage.ATTRIBUTE_VALUE_INTEGERs, &stage.ATTRIBUTE_VALUE_INTEGERs_reference, &stage.ATTRIBUTE_VALUE_INTEGERs_referenceOrder, &stage.ATTRIBUTE_VALUE_INTEGERs_instance)

	__gong__computeReferencePass1(stage, stage.ATTRIBUTE_VALUE_REALs, &stage.ATTRIBUTE_VALUE_REALs_reference, &stage.ATTRIBUTE_VALUE_REALs_referenceOrder, &stage.ATTRIBUTE_VALUE_REALs_instance)

	__gong__computeReferencePass1(stage, stage.ATTRIBUTE_VALUE_STRINGs, &stage.ATTRIBUTE_VALUE_STRINGs_reference, &stage.ATTRIBUTE_VALUE_STRINGs_referenceOrder, &stage.ATTRIBUTE_VALUE_STRINGs_instance)

	__gong__computeReferencePass1(stage, stage.ATTRIBUTE_VALUE_XHTMLs, &stage.ATTRIBUTE_VALUE_XHTMLs_reference, &stage.ATTRIBUTE_VALUE_XHTMLs_referenceOrder, &stage.ATTRIBUTE_VALUE_XHTMLs_instance)

	__gong__computeReferencePass1(stage, stage.A_ALTERNATIVE_IDs, &stage.A_ALTERNATIVE_IDs_reference, &stage.A_ALTERNATIVE_IDs_referenceOrder, &stage.A_ALTERNATIVE_IDs_instance)

	__gong__computeReferencePass1(stage, stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs, &stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_reference, &stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_referenceOrder, &stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_ATTRIBUTE_DEFINITION_DATE_REFs, &stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_reference, &stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_referenceOrder, &stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs, &stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_reference, &stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_referenceOrder, &stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs, &stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_reference, &stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_referenceOrder, &stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_ATTRIBUTE_DEFINITION_REAL_REFs, &stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_reference, &stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_referenceOrder, &stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_ATTRIBUTE_DEFINITION_STRING_REFs, &stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_reference, &stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_referenceOrder, &stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs, &stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_reference, &stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_referenceOrder, &stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_ATTRIBUTE_VALUE_BOOLEANs, &stage.A_ATTRIBUTE_VALUE_BOOLEANs_reference, &stage.A_ATTRIBUTE_VALUE_BOOLEANs_referenceOrder, &stage.A_ATTRIBUTE_VALUE_BOOLEANs_instance)

	__gong__computeReferencePass1(stage, stage.A_ATTRIBUTE_VALUE_DATEs, &stage.A_ATTRIBUTE_VALUE_DATEs_reference, &stage.A_ATTRIBUTE_VALUE_DATEs_referenceOrder, &stage.A_ATTRIBUTE_VALUE_DATEs_instance)

	__gong__computeReferencePass1(stage, stage.A_ATTRIBUTE_VALUE_ENUMERATIONs, &stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_reference, &stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder, &stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_instance)

	__gong__computeReferencePass1(stage, stage.A_ATTRIBUTE_VALUE_INTEGERs, &stage.A_ATTRIBUTE_VALUE_INTEGERs_reference, &stage.A_ATTRIBUTE_VALUE_INTEGERs_referenceOrder, &stage.A_ATTRIBUTE_VALUE_INTEGERs_instance)

	__gong__computeReferencePass1(stage, stage.A_ATTRIBUTE_VALUE_REALs, &stage.A_ATTRIBUTE_VALUE_REALs_reference, &stage.A_ATTRIBUTE_VALUE_REALs_referenceOrder, &stage.A_ATTRIBUTE_VALUE_REALs_instance)

	__gong__computeReferencePass1(stage, stage.A_ATTRIBUTE_VALUE_STRINGs, &stage.A_ATTRIBUTE_VALUE_STRINGs_reference, &stage.A_ATTRIBUTE_VALUE_STRINGs_referenceOrder, &stage.A_ATTRIBUTE_VALUE_STRINGs_instance)

	__gong__computeReferencePass1(stage, stage.A_ATTRIBUTE_VALUE_XHTMLs, &stage.A_ATTRIBUTE_VALUE_XHTMLs_reference, &stage.A_ATTRIBUTE_VALUE_XHTMLs_referenceOrder, &stage.A_ATTRIBUTE_VALUE_XHTMLs_instance)

	__gong__computeReferencePass1(stage, stage.A_ATTRIBUTE_VALUE_XHTML_1s, &stage.A_ATTRIBUTE_VALUE_XHTML_1s_reference, &stage.A_ATTRIBUTE_VALUE_XHTML_1s_referenceOrder, &stage.A_ATTRIBUTE_VALUE_XHTML_1s_instance)

	__gong__computeReferencePass1(stage, stage.A_CHILDRENs, &stage.A_CHILDRENs_reference, &stage.A_CHILDRENs_referenceOrder, &stage.A_CHILDRENs_instance)

	__gong__computeReferencePass1(stage, stage.A_CORE_CONTENTs, &stage.A_CORE_CONTENTs_reference, &stage.A_CORE_CONTENTs_referenceOrder, &stage.A_CORE_CONTENTs_instance)

	__gong__computeReferencePass1(stage, stage.A_DATATYPESs, &stage.A_DATATYPESs_reference, &stage.A_DATATYPESs_referenceOrder, &stage.A_DATATYPESs_instance)

	__gong__computeReferencePass1(stage, stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs, &stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_reference, &stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_referenceOrder, &stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_DATATYPE_DEFINITION_DATE_REFs, &stage.A_DATATYPE_DEFINITION_DATE_REFs_reference, &stage.A_DATATYPE_DEFINITION_DATE_REFs_referenceOrder, &stage.A_DATATYPE_DEFINITION_DATE_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs, &stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_reference, &stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_referenceOrder, &stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_DATATYPE_DEFINITION_INTEGER_REFs, &stage.A_DATATYPE_DEFINITION_INTEGER_REFs_reference, &stage.A_DATATYPE_DEFINITION_INTEGER_REFs_referenceOrder, &stage.A_DATATYPE_DEFINITION_INTEGER_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_DATATYPE_DEFINITION_REAL_REFs, &stage.A_DATATYPE_DEFINITION_REAL_REFs_reference, &stage.A_DATATYPE_DEFINITION_REAL_REFs_referenceOrder, &stage.A_DATATYPE_DEFINITION_REAL_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_DATATYPE_DEFINITION_STRING_REFs, &stage.A_DATATYPE_DEFINITION_STRING_REFs_reference, &stage.A_DATATYPE_DEFINITION_STRING_REFs_referenceOrder, &stage.A_DATATYPE_DEFINITION_STRING_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_DATATYPE_DEFINITION_XHTML_REFs, &stage.A_DATATYPE_DEFINITION_XHTML_REFs_reference, &stage.A_DATATYPE_DEFINITION_XHTML_REFs_referenceOrder, &stage.A_DATATYPE_DEFINITION_XHTML_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_EDITABLE_ATTSs, &stage.A_EDITABLE_ATTSs_reference, &stage.A_EDITABLE_ATTSs_referenceOrder, &stage.A_EDITABLE_ATTSs_instance)

	__gong__computeReferencePass1(stage, stage.A_ENUM_VALUE_REFs, &stage.A_ENUM_VALUE_REFs_reference, &stage.A_ENUM_VALUE_REFs_referenceOrder, &stage.A_ENUM_VALUE_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_OBJECTs, &stage.A_OBJECTs_reference, &stage.A_OBJECTs_referenceOrder, &stage.A_OBJECTs_instance)

	__gong__computeReferencePass1(stage, stage.A_PROPERTIESs, &stage.A_PROPERTIESs_reference, &stage.A_PROPERTIESs_referenceOrder, &stage.A_PROPERTIESs_instance)

	__gong__computeReferencePass1(stage, stage.A_RELATION_GROUP_TYPE_REFs, &stage.A_RELATION_GROUP_TYPE_REFs_reference, &stage.A_RELATION_GROUP_TYPE_REFs_referenceOrder, &stage.A_RELATION_GROUP_TYPE_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_SOURCE_1s, &stage.A_SOURCE_1s_reference, &stage.A_SOURCE_1s_referenceOrder, &stage.A_SOURCE_1s_instance)

	__gong__computeReferencePass1(stage, stage.A_SOURCE_SPECIFICATION_1s, &stage.A_SOURCE_SPECIFICATION_1s_reference, &stage.A_SOURCE_SPECIFICATION_1s_referenceOrder, &stage.A_SOURCE_SPECIFICATION_1s_instance)

	__gong__computeReferencePass1(stage, stage.A_SPECIFICATIONSs, &stage.A_SPECIFICATIONSs_reference, &stage.A_SPECIFICATIONSs_referenceOrder, &stage.A_SPECIFICATIONSs_instance)

	__gong__computeReferencePass1(stage, stage.A_SPECIFICATION_TYPE_REFs, &stage.A_SPECIFICATION_TYPE_REFs_reference, &stage.A_SPECIFICATION_TYPE_REFs_referenceOrder, &stage.A_SPECIFICATION_TYPE_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_SPECIFIED_VALUESs, &stage.A_SPECIFIED_VALUESs_reference, &stage.A_SPECIFIED_VALUESs_referenceOrder, &stage.A_SPECIFIED_VALUESs_instance)

	__gong__computeReferencePass1(stage, stage.A_SPEC_ATTRIBUTESs, &stage.A_SPEC_ATTRIBUTESs_reference, &stage.A_SPEC_ATTRIBUTESs_referenceOrder, &stage.A_SPEC_ATTRIBUTESs_instance)

	__gong__computeReferencePass1(stage, stage.A_SPEC_OBJECTSs, &stage.A_SPEC_OBJECTSs_reference, &stage.A_SPEC_OBJECTSs_referenceOrder, &stage.A_SPEC_OBJECTSs_instance)

	__gong__computeReferencePass1(stage, stage.A_SPEC_OBJECT_TYPE_REFs, &stage.A_SPEC_OBJECT_TYPE_REFs_reference, &stage.A_SPEC_OBJECT_TYPE_REFs_referenceOrder, &stage.A_SPEC_OBJECT_TYPE_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_SPEC_RELATIONSs, &stage.A_SPEC_RELATIONSs_reference, &stage.A_SPEC_RELATIONSs_referenceOrder, &stage.A_SPEC_RELATIONSs_instance)

	__gong__computeReferencePass1(stage, stage.A_SPEC_RELATION_GROUPSs, &stage.A_SPEC_RELATION_GROUPSs_reference, &stage.A_SPEC_RELATION_GROUPSs_referenceOrder, &stage.A_SPEC_RELATION_GROUPSs_instance)

	__gong__computeReferencePass1(stage, stage.A_SPEC_RELATION_REFs, &stage.A_SPEC_RELATION_REFs_reference, &stage.A_SPEC_RELATION_REFs_referenceOrder, &stage.A_SPEC_RELATION_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_SPEC_RELATION_TYPE_REFs, &stage.A_SPEC_RELATION_TYPE_REFs_reference, &stage.A_SPEC_RELATION_TYPE_REFs_referenceOrder, &stage.A_SPEC_RELATION_TYPE_REFs_instance)

	__gong__computeReferencePass1(stage, stage.A_SPEC_TYPESs, &stage.A_SPEC_TYPESs_reference, &stage.A_SPEC_TYPESs_referenceOrder, &stage.A_SPEC_TYPESs_instance)

	__gong__computeReferencePass1(stage, stage.A_THE_HEADERs, &stage.A_THE_HEADERs_reference, &stage.A_THE_HEADERs_referenceOrder, &stage.A_THE_HEADERs_instance)

	__gong__computeReferencePass1(stage, stage.A_TOOL_EXTENSIONSs, &stage.A_TOOL_EXTENSIONSs_reference, &stage.A_TOOL_EXTENSIONSs_referenceOrder, &stage.A_TOOL_EXTENSIONSs_instance)

	__gong__computeReferencePass1(stage, stage.DATATYPE_DEFINITION_BOOLEANs, &stage.DATATYPE_DEFINITION_BOOLEANs_reference, &stage.DATATYPE_DEFINITION_BOOLEANs_referenceOrder, &stage.DATATYPE_DEFINITION_BOOLEANs_instance)

	__gong__computeReferencePass1(stage, stage.DATATYPE_DEFINITION_DATEs, &stage.DATATYPE_DEFINITION_DATEs_reference, &stage.DATATYPE_DEFINITION_DATEs_referenceOrder, &stage.DATATYPE_DEFINITION_DATEs_instance)

	__gong__computeReferencePass1(stage, stage.DATATYPE_DEFINITION_ENUMERATIONs, &stage.DATATYPE_DEFINITION_ENUMERATIONs_reference, &stage.DATATYPE_DEFINITION_ENUMERATIONs_referenceOrder, &stage.DATATYPE_DEFINITION_ENUMERATIONs_instance)

	__gong__computeReferencePass1(stage, stage.DATATYPE_DEFINITION_INTEGERs, &stage.DATATYPE_DEFINITION_INTEGERs_reference, &stage.DATATYPE_DEFINITION_INTEGERs_referenceOrder, &stage.DATATYPE_DEFINITION_INTEGERs_instance)

	__gong__computeReferencePass1(stage, stage.DATATYPE_DEFINITION_REALs, &stage.DATATYPE_DEFINITION_REALs_reference, &stage.DATATYPE_DEFINITION_REALs_referenceOrder, &stage.DATATYPE_DEFINITION_REALs_instance)

	__gong__computeReferencePass1(stage, stage.DATATYPE_DEFINITION_STRINGs, &stage.DATATYPE_DEFINITION_STRINGs_reference, &stage.DATATYPE_DEFINITION_STRINGs_referenceOrder, &stage.DATATYPE_DEFINITION_STRINGs_instance)

	__gong__computeReferencePass1(stage, stage.DATATYPE_DEFINITION_XHTMLs, &stage.DATATYPE_DEFINITION_XHTMLs_reference, &stage.DATATYPE_DEFINITION_XHTMLs_referenceOrder, &stage.DATATYPE_DEFINITION_XHTMLs_instance)

	__gong__computeReferencePass1(stage, stage.EMBEDDED_VALUEs, &stage.EMBEDDED_VALUEs_reference, &stage.EMBEDDED_VALUEs_referenceOrder, &stage.EMBEDDED_VALUEs_instance)

	__gong__computeReferencePass1(stage, stage.ENUM_VALUEs, &stage.ENUM_VALUEs_reference, &stage.ENUM_VALUEs_referenceOrder, &stage.ENUM_VALUEs_instance)

	__gong__computeReferencePass1(stage, stage.RELATION_GROUPs, &stage.RELATION_GROUPs_reference, &stage.RELATION_GROUPs_referenceOrder, &stage.RELATION_GROUPs_instance)

	__gong__computeReferencePass1(stage, stage.RELATION_GROUP_TYPEs, &stage.RELATION_GROUP_TYPEs_reference, &stage.RELATION_GROUP_TYPEs_referenceOrder, &stage.RELATION_GROUP_TYPEs_instance)

	__gong__computeReferencePass1(stage, stage.REQ_IFs, &stage.REQ_IFs_reference, &stage.REQ_IFs_referenceOrder, &stage.REQ_IFs_instance)

	__gong__computeReferencePass1(stage, stage.REQ_IF_CONTENTs, &stage.REQ_IF_CONTENTs_reference, &stage.REQ_IF_CONTENTs_referenceOrder, &stage.REQ_IF_CONTENTs_instance)

	__gong__computeReferencePass1(stage, stage.REQ_IF_HEADERs, &stage.REQ_IF_HEADERs_reference, &stage.REQ_IF_HEADERs_referenceOrder, &stage.REQ_IF_HEADERs_instance)

	__gong__computeReferencePass1(stage, stage.REQ_IF_TOOL_EXTENSIONs, &stage.REQ_IF_TOOL_EXTENSIONs_reference, &stage.REQ_IF_TOOL_EXTENSIONs_referenceOrder, &stage.REQ_IF_TOOL_EXTENSIONs_instance)

	__gong__computeReferencePass1(stage, stage.SPECIFICATIONs, &stage.SPECIFICATIONs_reference, &stage.SPECIFICATIONs_referenceOrder, &stage.SPECIFICATIONs_instance)

	__gong__computeReferencePass1(stage, stage.SPECIFICATION_TYPEs, &stage.SPECIFICATION_TYPEs_reference, &stage.SPECIFICATION_TYPEs_referenceOrder, &stage.SPECIFICATION_TYPEs_instance)

	__gong__computeReferencePass1(stage, stage.SPEC_HIERARCHYs, &stage.SPEC_HIERARCHYs_reference, &stage.SPEC_HIERARCHYs_referenceOrder, &stage.SPEC_HIERARCHYs_instance)

	__gong__computeReferencePass1(stage, stage.SPEC_OBJECTs, &stage.SPEC_OBJECTs_reference, &stage.SPEC_OBJECTs_referenceOrder, &stage.SPEC_OBJECTs_instance)

	__gong__computeReferencePass1(stage, stage.SPEC_OBJECT_TYPEs, &stage.SPEC_OBJECT_TYPEs_reference, &stage.SPEC_OBJECT_TYPEs_referenceOrder, &stage.SPEC_OBJECT_TYPEs_instance)

	__gong__computeReferencePass1(stage, stage.SPEC_RELATIONs, &stage.SPEC_RELATIONs_reference, &stage.SPEC_RELATIONs_referenceOrder, &stage.SPEC_RELATIONs_instance)

	__gong__computeReferencePass1(stage, stage.SPEC_RELATION_TYPEs, &stage.SPEC_RELATION_TYPEs_reference, &stage.SPEC_RELATION_TYPEs_referenceOrder, &stage.SPEC_RELATION_TYPEs_instance)

	__gong__computeReferencePass1(stage, stage.XHTML_CONTENTs, &stage.XHTML_CONTENTs_reference, &stage.XHTML_CONTENTs_referenceOrder, &stage.XHTML_CONTENTs_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.ALTERNATIVE_IDs, stage.ALTERNATIVE_IDs_reference, stage)

	__gong__computeReferencePass2(stage.ATTRIBUTE_DEFINITION_BOOLEANs, stage.ATTRIBUTE_DEFINITION_BOOLEANs_reference, stage)

	__gong__computeReferencePass2(stage.ATTRIBUTE_DEFINITION_DATEs, stage.ATTRIBUTE_DEFINITION_DATEs_reference, stage)

	__gong__computeReferencePass2(stage.ATTRIBUTE_DEFINITION_ENUMERATIONs, stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_reference, stage)

	__gong__computeReferencePass2(stage.ATTRIBUTE_DEFINITION_INTEGERs, stage.ATTRIBUTE_DEFINITION_INTEGERs_reference, stage)

	__gong__computeReferencePass2(stage.ATTRIBUTE_DEFINITION_REALs, stage.ATTRIBUTE_DEFINITION_REALs_reference, stage)

	__gong__computeReferencePass2(stage.ATTRIBUTE_DEFINITION_STRINGs, stage.ATTRIBUTE_DEFINITION_STRINGs_reference, stage)

	__gong__computeReferencePass2(stage.ATTRIBUTE_DEFINITION_XHTMLs, stage.ATTRIBUTE_DEFINITION_XHTMLs_reference, stage)

	__gong__computeReferencePass2(stage.ATTRIBUTE_VALUE_BOOLEANs, stage.ATTRIBUTE_VALUE_BOOLEANs_reference, stage)

	__gong__computeReferencePass2(stage.ATTRIBUTE_VALUE_DATEs, stage.ATTRIBUTE_VALUE_DATEs_reference, stage)

	__gong__computeReferencePass2(stage.ATTRIBUTE_VALUE_ENUMERATIONs, stage.ATTRIBUTE_VALUE_ENUMERATIONs_reference, stage)

	__gong__computeReferencePass2(stage.ATTRIBUTE_VALUE_INTEGERs, stage.ATTRIBUTE_VALUE_INTEGERs_reference, stage)

	__gong__computeReferencePass2(stage.ATTRIBUTE_VALUE_REALs, stage.ATTRIBUTE_VALUE_REALs_reference, stage)

	__gong__computeReferencePass2(stage.ATTRIBUTE_VALUE_STRINGs, stage.ATTRIBUTE_VALUE_STRINGs_reference, stage)

	__gong__computeReferencePass2(stage.ATTRIBUTE_VALUE_XHTMLs, stage.ATTRIBUTE_VALUE_XHTMLs_reference, stage)

	__gong__computeReferencePass2(stage.A_ALTERNATIVE_IDs, stage.A_ALTERNATIVE_IDs_reference, stage)

	__gong__computeReferencePass2(stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs, stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_ATTRIBUTE_DEFINITION_DATE_REFs, stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs, stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs, stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_ATTRIBUTE_DEFINITION_REAL_REFs, stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_ATTRIBUTE_DEFINITION_STRING_REFs, stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs, stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_ATTRIBUTE_VALUE_BOOLEANs, stage.A_ATTRIBUTE_VALUE_BOOLEANs_reference, stage)

	__gong__computeReferencePass2(stage.A_ATTRIBUTE_VALUE_DATEs, stage.A_ATTRIBUTE_VALUE_DATEs_reference, stage)

	__gong__computeReferencePass2(stage.A_ATTRIBUTE_VALUE_ENUMERATIONs, stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_reference, stage)

	__gong__computeReferencePass2(stage.A_ATTRIBUTE_VALUE_INTEGERs, stage.A_ATTRIBUTE_VALUE_INTEGERs_reference, stage)

	__gong__computeReferencePass2(stage.A_ATTRIBUTE_VALUE_REALs, stage.A_ATTRIBUTE_VALUE_REALs_reference, stage)

	__gong__computeReferencePass2(stage.A_ATTRIBUTE_VALUE_STRINGs, stage.A_ATTRIBUTE_VALUE_STRINGs_reference, stage)

	__gong__computeReferencePass2(stage.A_ATTRIBUTE_VALUE_XHTMLs, stage.A_ATTRIBUTE_VALUE_XHTMLs_reference, stage)

	__gong__computeReferencePass2(stage.A_ATTRIBUTE_VALUE_XHTML_1s, stage.A_ATTRIBUTE_VALUE_XHTML_1s_reference, stage)

	__gong__computeReferencePass2(stage.A_CHILDRENs, stage.A_CHILDRENs_reference, stage)

	__gong__computeReferencePass2(stage.A_CORE_CONTENTs, stage.A_CORE_CONTENTs_reference, stage)

	__gong__computeReferencePass2(stage.A_DATATYPESs, stage.A_DATATYPESs_reference, stage)

	__gong__computeReferencePass2(stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs, stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_DATATYPE_DEFINITION_DATE_REFs, stage.A_DATATYPE_DEFINITION_DATE_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs, stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_DATATYPE_DEFINITION_INTEGER_REFs, stage.A_DATATYPE_DEFINITION_INTEGER_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_DATATYPE_DEFINITION_REAL_REFs, stage.A_DATATYPE_DEFINITION_REAL_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_DATATYPE_DEFINITION_STRING_REFs, stage.A_DATATYPE_DEFINITION_STRING_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_DATATYPE_DEFINITION_XHTML_REFs, stage.A_DATATYPE_DEFINITION_XHTML_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_EDITABLE_ATTSs, stage.A_EDITABLE_ATTSs_reference, stage)

	__gong__computeReferencePass2(stage.A_ENUM_VALUE_REFs, stage.A_ENUM_VALUE_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_OBJECTs, stage.A_OBJECTs_reference, stage)

	__gong__computeReferencePass2(stage.A_PROPERTIESs, stage.A_PROPERTIESs_reference, stage)

	__gong__computeReferencePass2(stage.A_RELATION_GROUP_TYPE_REFs, stage.A_RELATION_GROUP_TYPE_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_SOURCE_1s, stage.A_SOURCE_1s_reference, stage)

	__gong__computeReferencePass2(stage.A_SOURCE_SPECIFICATION_1s, stage.A_SOURCE_SPECIFICATION_1s_reference, stage)

	__gong__computeReferencePass2(stage.A_SPECIFICATIONSs, stage.A_SPECIFICATIONSs_reference, stage)

	__gong__computeReferencePass2(stage.A_SPECIFICATION_TYPE_REFs, stage.A_SPECIFICATION_TYPE_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_SPECIFIED_VALUESs, stage.A_SPECIFIED_VALUESs_reference, stage)

	__gong__computeReferencePass2(stage.A_SPEC_ATTRIBUTESs, stage.A_SPEC_ATTRIBUTESs_reference, stage)

	__gong__computeReferencePass2(stage.A_SPEC_OBJECTSs, stage.A_SPEC_OBJECTSs_reference, stage)

	__gong__computeReferencePass2(stage.A_SPEC_OBJECT_TYPE_REFs, stage.A_SPEC_OBJECT_TYPE_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_SPEC_RELATIONSs, stage.A_SPEC_RELATIONSs_reference, stage)

	__gong__computeReferencePass2(stage.A_SPEC_RELATION_GROUPSs, stage.A_SPEC_RELATION_GROUPSs_reference, stage)

	__gong__computeReferencePass2(stage.A_SPEC_RELATION_REFs, stage.A_SPEC_RELATION_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_SPEC_RELATION_TYPE_REFs, stage.A_SPEC_RELATION_TYPE_REFs_reference, stage)

	__gong__computeReferencePass2(stage.A_SPEC_TYPESs, stage.A_SPEC_TYPESs_reference, stage)

	__gong__computeReferencePass2(stage.A_THE_HEADERs, stage.A_THE_HEADERs_reference, stage)

	__gong__computeReferencePass2(stage.A_TOOL_EXTENSIONSs, stage.A_TOOL_EXTENSIONSs_reference, stage)

	__gong__computeReferencePass2(stage.DATATYPE_DEFINITION_BOOLEANs, stage.DATATYPE_DEFINITION_BOOLEANs_reference, stage)

	__gong__computeReferencePass2(stage.DATATYPE_DEFINITION_DATEs, stage.DATATYPE_DEFINITION_DATEs_reference, stage)

	__gong__computeReferencePass2(stage.DATATYPE_DEFINITION_ENUMERATIONs, stage.DATATYPE_DEFINITION_ENUMERATIONs_reference, stage)

	__gong__computeReferencePass2(stage.DATATYPE_DEFINITION_INTEGERs, stage.DATATYPE_DEFINITION_INTEGERs_reference, stage)

	__gong__computeReferencePass2(stage.DATATYPE_DEFINITION_REALs, stage.DATATYPE_DEFINITION_REALs_reference, stage)

	__gong__computeReferencePass2(stage.DATATYPE_DEFINITION_STRINGs, stage.DATATYPE_DEFINITION_STRINGs_reference, stage)

	__gong__computeReferencePass2(stage.DATATYPE_DEFINITION_XHTMLs, stage.DATATYPE_DEFINITION_XHTMLs_reference, stage)

	__gong__computeReferencePass2(stage.EMBEDDED_VALUEs, stage.EMBEDDED_VALUEs_reference, stage)

	__gong__computeReferencePass2(stage.ENUM_VALUEs, stage.ENUM_VALUEs_reference, stage)

	__gong__computeReferencePass2(stage.RELATION_GROUPs, stage.RELATION_GROUPs_reference, stage)

	__gong__computeReferencePass2(stage.RELATION_GROUP_TYPEs, stage.RELATION_GROUP_TYPEs_reference, stage)

	__gong__computeReferencePass2(stage.REQ_IFs, stage.REQ_IFs_reference, stage)

	__gong__computeReferencePass2(stage.REQ_IF_CONTENTs, stage.REQ_IF_CONTENTs_reference, stage)

	__gong__computeReferencePass2(stage.REQ_IF_HEADERs, stage.REQ_IF_HEADERs_reference, stage)

	__gong__computeReferencePass2(stage.REQ_IF_TOOL_EXTENSIONs, stage.REQ_IF_TOOL_EXTENSIONs_reference, stage)

	__gong__computeReferencePass2(stage.SPECIFICATIONs, stage.SPECIFICATIONs_reference, stage)

	__gong__computeReferencePass2(stage.SPECIFICATION_TYPEs, stage.SPECIFICATION_TYPEs_reference, stage)

	__gong__computeReferencePass2(stage.SPEC_HIERARCHYs, stage.SPEC_HIERARCHYs_reference, stage)

	__gong__computeReferencePass2(stage.SPEC_OBJECTs, stage.SPEC_OBJECTs_reference, stage)

	__gong__computeReferencePass2(stage.SPEC_OBJECT_TYPEs, stage.SPEC_OBJECT_TYPEs_reference, stage)

	__gong__computeReferencePass2(stage.SPEC_RELATIONs, stage.SPEC_RELATIONs_reference, stage)

	__gong__computeReferencePass2(stage.SPEC_RELATION_TYPEs, stage.SPEC_RELATION_TYPEs_reference, stage)

	__gong__computeReferencePass2(stage.XHTML_CONTENTs, stage.XHTML_CONTENTs_reference, stage)

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (alternative_id *ALTERNATIVE_ID) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ALTERNATIVE_ID_stagedOrder, stage.ALTERNATIVE_IDs_referenceOrder, alternative_id, "ALTERNATIVE_ID")
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ATTRIBUTE_DEFINITION_BOOLEAN_stagedOrder, stage.ATTRIBUTE_DEFINITION_BOOLEANs_referenceOrder, attribute_definition_boolean, "ATTRIBUTE_DEFINITION_BOOLEAN")
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ATTRIBUTE_DEFINITION_DATE_stagedOrder, stage.ATTRIBUTE_DEFINITION_DATEs_referenceOrder, attribute_definition_date, "ATTRIBUTE_DEFINITION_DATE")
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ATTRIBUTE_DEFINITION_ENUMERATION_stagedOrder, stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_referenceOrder, attribute_definition_enumeration, "ATTRIBUTE_DEFINITION_ENUMERATION")
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ATTRIBUTE_DEFINITION_INTEGER_stagedOrder, stage.ATTRIBUTE_DEFINITION_INTEGERs_referenceOrder, attribute_definition_integer, "ATTRIBUTE_DEFINITION_INTEGER")
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ATTRIBUTE_DEFINITION_REAL_stagedOrder, stage.ATTRIBUTE_DEFINITION_REALs_referenceOrder, attribute_definition_real, "ATTRIBUTE_DEFINITION_REAL")
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ATTRIBUTE_DEFINITION_STRING_stagedOrder, stage.ATTRIBUTE_DEFINITION_STRINGs_referenceOrder, attribute_definition_string, "ATTRIBUTE_DEFINITION_STRING")
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ATTRIBUTE_DEFINITION_XHTML_stagedOrder, stage.ATTRIBUTE_DEFINITION_XHTMLs_referenceOrder, attribute_definition_xhtml, "ATTRIBUTE_DEFINITION_XHTML")
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ATTRIBUTE_VALUE_BOOLEAN_stagedOrder, stage.ATTRIBUTE_VALUE_BOOLEANs_referenceOrder, attribute_value_boolean, "ATTRIBUTE_VALUE_BOOLEAN")
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ATTRIBUTE_VALUE_DATE_stagedOrder, stage.ATTRIBUTE_VALUE_DATEs_referenceOrder, attribute_value_date, "ATTRIBUTE_VALUE_DATE")
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ATTRIBUTE_VALUE_ENUMERATION_stagedOrder, stage.ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder, attribute_value_enumeration, "ATTRIBUTE_VALUE_ENUMERATION")
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ATTRIBUTE_VALUE_INTEGER_stagedOrder, stage.ATTRIBUTE_VALUE_INTEGERs_referenceOrder, attribute_value_integer, "ATTRIBUTE_VALUE_INTEGER")
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ATTRIBUTE_VALUE_REAL_stagedOrder, stage.ATTRIBUTE_VALUE_REALs_referenceOrder, attribute_value_real, "ATTRIBUTE_VALUE_REAL")
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ATTRIBUTE_VALUE_STRING_stagedOrder, stage.ATTRIBUTE_VALUE_STRINGs_referenceOrder, attribute_value_string, "ATTRIBUTE_VALUE_STRING")
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ATTRIBUTE_VALUE_XHTML_stagedOrder, stage.ATTRIBUTE_VALUE_XHTMLs_referenceOrder, attribute_value_xhtml, "ATTRIBUTE_VALUE_XHTML")
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ALTERNATIVE_ID_stagedOrder, stage.A_ALTERNATIVE_IDs_referenceOrder, a_alternative_id, "A_ALTERNATIVE_ID")
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_stagedOrder, stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_referenceOrder, a_attribute_definition_boolean_ref, "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF")
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ATTRIBUTE_DEFINITION_DATE_REF_stagedOrder, stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_referenceOrder, a_attribute_definition_date_ref, "A_ATTRIBUTE_DEFINITION_DATE_REF")
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_stagedOrder, stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_referenceOrder, a_attribute_definition_enumeration_ref, "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF")
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ATTRIBUTE_DEFINITION_INTEGER_REF_stagedOrder, stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_referenceOrder, a_attribute_definition_integer_ref, "A_ATTRIBUTE_DEFINITION_INTEGER_REF")
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ATTRIBUTE_DEFINITION_REAL_REF_stagedOrder, stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_referenceOrder, a_attribute_definition_real_ref, "A_ATTRIBUTE_DEFINITION_REAL_REF")
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ATTRIBUTE_DEFINITION_STRING_REF_stagedOrder, stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_referenceOrder, a_attribute_definition_string_ref, "A_ATTRIBUTE_DEFINITION_STRING_REF")
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ATTRIBUTE_DEFINITION_XHTML_REF_stagedOrder, stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_referenceOrder, a_attribute_definition_xhtml_ref, "A_ATTRIBUTE_DEFINITION_XHTML_REF")
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ATTRIBUTE_VALUE_BOOLEAN_stagedOrder, stage.A_ATTRIBUTE_VALUE_BOOLEANs_referenceOrder, a_attribute_value_boolean, "A_ATTRIBUTE_VALUE_BOOLEAN")
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ATTRIBUTE_VALUE_DATE_stagedOrder, stage.A_ATTRIBUTE_VALUE_DATEs_referenceOrder, a_attribute_value_date, "A_ATTRIBUTE_VALUE_DATE")
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ATTRIBUTE_VALUE_ENUMERATION_stagedOrder, stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder, a_attribute_value_enumeration, "A_ATTRIBUTE_VALUE_ENUMERATION")
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ATTRIBUTE_VALUE_INTEGER_stagedOrder, stage.A_ATTRIBUTE_VALUE_INTEGERs_referenceOrder, a_attribute_value_integer, "A_ATTRIBUTE_VALUE_INTEGER")
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ATTRIBUTE_VALUE_REAL_stagedOrder, stage.A_ATTRIBUTE_VALUE_REALs_referenceOrder, a_attribute_value_real, "A_ATTRIBUTE_VALUE_REAL")
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ATTRIBUTE_VALUE_STRING_stagedOrder, stage.A_ATTRIBUTE_VALUE_STRINGs_referenceOrder, a_attribute_value_string, "A_ATTRIBUTE_VALUE_STRING")
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ATTRIBUTE_VALUE_XHTML_stagedOrder, stage.A_ATTRIBUTE_VALUE_XHTMLs_referenceOrder, a_attribute_value_xhtml, "A_ATTRIBUTE_VALUE_XHTML")
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ATTRIBUTE_VALUE_XHTML_1_stagedOrder, stage.A_ATTRIBUTE_VALUE_XHTML_1s_referenceOrder, a_attribute_value_xhtml_1, "A_ATTRIBUTE_VALUE_XHTML_1")
}

func (a_children *A_CHILDREN) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_CHILDREN_stagedOrder, stage.A_CHILDRENs_referenceOrder, a_children, "A_CHILDREN")
}

func (a_core_content *A_CORE_CONTENT) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_CORE_CONTENT_stagedOrder, stage.A_CORE_CONTENTs_referenceOrder, a_core_content, "A_CORE_CONTENT")
}

func (a_datatypes *A_DATATYPES) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_DATATYPES_stagedOrder, stage.A_DATATYPESs_referenceOrder, a_datatypes, "A_DATATYPES")
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_DATATYPE_DEFINITION_BOOLEAN_REF_stagedOrder, stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_referenceOrder, a_datatype_definition_boolean_ref, "A_DATATYPE_DEFINITION_BOOLEAN_REF")
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_DATATYPE_DEFINITION_DATE_REF_stagedOrder, stage.A_DATATYPE_DEFINITION_DATE_REFs_referenceOrder, a_datatype_definition_date_ref, "A_DATATYPE_DEFINITION_DATE_REF")
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_DATATYPE_DEFINITION_ENUMERATION_REF_stagedOrder, stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_referenceOrder, a_datatype_definition_enumeration_ref, "A_DATATYPE_DEFINITION_ENUMERATION_REF")
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_DATATYPE_DEFINITION_INTEGER_REF_stagedOrder, stage.A_DATATYPE_DEFINITION_INTEGER_REFs_referenceOrder, a_datatype_definition_integer_ref, "A_DATATYPE_DEFINITION_INTEGER_REF")
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_DATATYPE_DEFINITION_REAL_REF_stagedOrder, stage.A_DATATYPE_DEFINITION_REAL_REFs_referenceOrder, a_datatype_definition_real_ref, "A_DATATYPE_DEFINITION_REAL_REF")
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_DATATYPE_DEFINITION_STRING_REF_stagedOrder, stage.A_DATATYPE_DEFINITION_STRING_REFs_referenceOrder, a_datatype_definition_string_ref, "A_DATATYPE_DEFINITION_STRING_REF")
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_DATATYPE_DEFINITION_XHTML_REF_stagedOrder, stage.A_DATATYPE_DEFINITION_XHTML_REFs_referenceOrder, a_datatype_definition_xhtml_ref, "A_DATATYPE_DEFINITION_XHTML_REF")
}

func (a_editable_atts *A_EDITABLE_ATTS) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_EDITABLE_ATTS_stagedOrder, stage.A_EDITABLE_ATTSs_referenceOrder, a_editable_atts, "A_EDITABLE_ATTS")
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_ENUM_VALUE_REF_stagedOrder, stage.A_ENUM_VALUE_REFs_referenceOrder, a_enum_value_ref, "A_ENUM_VALUE_REF")
}

func (a_object *A_OBJECT) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_OBJECT_stagedOrder, stage.A_OBJECTs_referenceOrder, a_object, "A_OBJECT")
}

func (a_properties *A_PROPERTIES) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_PROPERTIES_stagedOrder, stage.A_PROPERTIESs_referenceOrder, a_properties, "A_PROPERTIES")
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_RELATION_GROUP_TYPE_REF_stagedOrder, stage.A_RELATION_GROUP_TYPE_REFs_referenceOrder, a_relation_group_type_ref, "A_RELATION_GROUP_TYPE_REF")
}

func (a_source_1 *A_SOURCE_1) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_SOURCE_1_stagedOrder, stage.A_SOURCE_1s_referenceOrder, a_source_1, "A_SOURCE_1")
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_SOURCE_SPECIFICATION_1_stagedOrder, stage.A_SOURCE_SPECIFICATION_1s_referenceOrder, a_source_specification_1, "A_SOURCE_SPECIFICATION_1")
}

func (a_specifications *A_SPECIFICATIONS) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_SPECIFICATIONS_stagedOrder, stage.A_SPECIFICATIONSs_referenceOrder, a_specifications, "A_SPECIFICATIONS")
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_SPECIFICATION_TYPE_REF_stagedOrder, stage.A_SPECIFICATION_TYPE_REFs_referenceOrder, a_specification_type_ref, "A_SPECIFICATION_TYPE_REF")
}

func (a_specified_values *A_SPECIFIED_VALUES) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_SPECIFIED_VALUES_stagedOrder, stage.A_SPECIFIED_VALUESs_referenceOrder, a_specified_values, "A_SPECIFIED_VALUES")
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_SPEC_ATTRIBUTES_stagedOrder, stage.A_SPEC_ATTRIBUTESs_referenceOrder, a_spec_attributes, "A_SPEC_ATTRIBUTES")
}

func (a_spec_objects *A_SPEC_OBJECTS) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_SPEC_OBJECTS_stagedOrder, stage.A_SPEC_OBJECTSs_referenceOrder, a_spec_objects, "A_SPEC_OBJECTS")
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_SPEC_OBJECT_TYPE_REF_stagedOrder, stage.A_SPEC_OBJECT_TYPE_REFs_referenceOrder, a_spec_object_type_ref, "A_SPEC_OBJECT_TYPE_REF")
}

func (a_spec_relations *A_SPEC_RELATIONS) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_SPEC_RELATIONS_stagedOrder, stage.A_SPEC_RELATIONSs_referenceOrder, a_spec_relations, "A_SPEC_RELATIONS")
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_SPEC_RELATION_GROUPS_stagedOrder, stage.A_SPEC_RELATION_GROUPSs_referenceOrder, a_spec_relation_groups, "A_SPEC_RELATION_GROUPS")
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_SPEC_RELATION_REF_stagedOrder, stage.A_SPEC_RELATION_REFs_referenceOrder, a_spec_relation_ref, "A_SPEC_RELATION_REF")
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_SPEC_RELATION_TYPE_REF_stagedOrder, stage.A_SPEC_RELATION_TYPE_REFs_referenceOrder, a_spec_relation_type_ref, "A_SPEC_RELATION_TYPE_REF")
}

func (a_spec_types *A_SPEC_TYPES) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_SPEC_TYPES_stagedOrder, stage.A_SPEC_TYPESs_referenceOrder, a_spec_types, "A_SPEC_TYPES")
}

func (a_the_header *A_THE_HEADER) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_THE_HEADER_stagedOrder, stage.A_THE_HEADERs_referenceOrder, a_the_header, "A_THE_HEADER")
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_TOOL_EXTENSIONS_stagedOrder, stage.A_TOOL_EXTENSIONSs_referenceOrder, a_tool_extensions, "A_TOOL_EXTENSIONS")
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DATATYPE_DEFINITION_BOOLEAN_stagedOrder, stage.DATATYPE_DEFINITION_BOOLEANs_referenceOrder, datatype_definition_boolean, "DATATYPE_DEFINITION_BOOLEAN")
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DATATYPE_DEFINITION_DATE_stagedOrder, stage.DATATYPE_DEFINITION_DATEs_referenceOrder, datatype_definition_date, "DATATYPE_DEFINITION_DATE")
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DATATYPE_DEFINITION_ENUMERATION_stagedOrder, stage.DATATYPE_DEFINITION_ENUMERATIONs_referenceOrder, datatype_definition_enumeration, "DATATYPE_DEFINITION_ENUMERATION")
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DATATYPE_DEFINITION_INTEGER_stagedOrder, stage.DATATYPE_DEFINITION_INTEGERs_referenceOrder, datatype_definition_integer, "DATATYPE_DEFINITION_INTEGER")
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DATATYPE_DEFINITION_REAL_stagedOrder, stage.DATATYPE_DEFINITION_REALs_referenceOrder, datatype_definition_real, "DATATYPE_DEFINITION_REAL")
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DATATYPE_DEFINITION_STRING_stagedOrder, stage.DATATYPE_DEFINITION_STRINGs_referenceOrder, datatype_definition_string, "DATATYPE_DEFINITION_STRING")
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.DATATYPE_DEFINITION_XHTML_stagedOrder, stage.DATATYPE_DEFINITION_XHTMLs_referenceOrder, datatype_definition_xhtml, "DATATYPE_DEFINITION_XHTML")
}

func (embedded_value *EMBEDDED_VALUE) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.EMBEDDED_VALUE_stagedOrder, stage.EMBEDDED_VALUEs_referenceOrder, embedded_value, "EMBEDDED_VALUE")
}

func (enum_value *ENUM_VALUE) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.ENUM_VALUE_stagedOrder, stage.ENUM_VALUEs_referenceOrder, enum_value, "ENUM_VALUE")
}

func (relation_group *RELATION_GROUP) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RELATION_GROUP_stagedOrder, stage.RELATION_GROUPs_referenceOrder, relation_group, "RELATION_GROUP")
}

func (relation_group_type *RELATION_GROUP_TYPE) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.RELATION_GROUP_TYPE_stagedOrder, stage.RELATION_GROUP_TYPEs_referenceOrder, relation_group_type, "RELATION_GROUP_TYPE")
}

func (req_if *REQ_IF) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.REQ_IF_stagedOrder, stage.REQ_IFs_referenceOrder, req_if, "REQ_IF")
}

func (req_if_content *REQ_IF_CONTENT) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.REQ_IF_CONTENT_stagedOrder, stage.REQ_IF_CONTENTs_referenceOrder, req_if_content, "REQ_IF_CONTENT")
}

func (req_if_header *REQ_IF_HEADER) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.REQ_IF_HEADER_stagedOrder, stage.REQ_IF_HEADERs_referenceOrder, req_if_header, "REQ_IF_HEADER")
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.REQ_IF_TOOL_EXTENSION_stagedOrder, stage.REQ_IF_TOOL_EXTENSIONs_referenceOrder, req_if_tool_extension, "REQ_IF_TOOL_EXTENSION")
}

func (specification *SPECIFICATION) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SPECIFICATION_stagedOrder, stage.SPECIFICATIONs_referenceOrder, specification, "SPECIFICATION")
}

func (specification_type *SPECIFICATION_TYPE) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SPECIFICATION_TYPE_stagedOrder, stage.SPECIFICATION_TYPEs_referenceOrder, specification_type, "SPECIFICATION_TYPE")
}

func (spec_hierarchy *SPEC_HIERARCHY) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SPEC_HIERARCHY_stagedOrder, stage.SPEC_HIERARCHYs_referenceOrder, spec_hierarchy, "SPEC_HIERARCHY")
}

func (spec_object *SPEC_OBJECT) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SPEC_OBJECT_stagedOrder, stage.SPEC_OBJECTs_referenceOrder, spec_object, "SPEC_OBJECT")
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SPEC_OBJECT_TYPE_stagedOrder, stage.SPEC_OBJECT_TYPEs_referenceOrder, spec_object_type, "SPEC_OBJECT_TYPE")
}

func (spec_relation *SPEC_RELATION) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SPEC_RELATION_stagedOrder, stage.SPEC_RELATIONs_referenceOrder, spec_relation, "SPEC_RELATION")
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.SPEC_RELATION_TYPE_stagedOrder, stage.SPEC_RELATION_TYPEs_referenceOrder, spec_relation_type, "SPEC_RELATION_TYPE")
}

func (xhtml_content *XHTML_CONTENT) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.XHTML_CONTENT_stagedOrder, stage.XHTML_CONTENTs_referenceOrder, xhtml_content, "XHTML_CONTENT")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (alternative_id *ALTERNATIVE_ID) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(alternative_id, alternative_id.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (alternative_id *ALTERNATIVE_ID) GongGetReferenceIdentifier(stage *Stage) string {
	return alternative_id.GongGetIdentifier(stage)
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attribute_definition_boolean, attribute_definition_boolean.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongGetReferenceIdentifier(stage *Stage) string {
	return attribute_definition_boolean.GongGetIdentifier(stage)
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attribute_definition_date, attribute_definition_date.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongGetReferenceIdentifier(stage *Stage) string {
	return attribute_definition_date.GongGetIdentifier(stage)
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attribute_definition_enumeration, attribute_definition_enumeration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongGetReferenceIdentifier(stage *Stage) string {
	return attribute_definition_enumeration.GongGetIdentifier(stage)
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attribute_definition_integer, attribute_definition_integer.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongGetReferenceIdentifier(stage *Stage) string {
	return attribute_definition_integer.GongGetIdentifier(stage)
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attribute_definition_real, attribute_definition_real.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongGetReferenceIdentifier(stage *Stage) string {
	return attribute_definition_real.GongGetIdentifier(stage)
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attribute_definition_string, attribute_definition_string.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongGetReferenceIdentifier(stage *Stage) string {
	return attribute_definition_string.GongGetIdentifier(stage)
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attribute_definition_xhtml, attribute_definition_xhtml.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongGetReferenceIdentifier(stage *Stage) string {
	return attribute_definition_xhtml.GongGetIdentifier(stage)
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attribute_value_boolean, attribute_value_boolean.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongGetReferenceIdentifier(stage *Stage) string {
	return attribute_value_boolean.GongGetIdentifier(stage)
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attribute_value_date, attribute_value_date.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongGetReferenceIdentifier(stage *Stage) string {
	return attribute_value_date.GongGetIdentifier(stage)
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attribute_value_enumeration, attribute_value_enumeration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongGetReferenceIdentifier(stage *Stage) string {
	return attribute_value_enumeration.GongGetIdentifier(stage)
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attribute_value_integer, attribute_value_integer.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongGetReferenceIdentifier(stage *Stage) string {
	return attribute_value_integer.GongGetIdentifier(stage)
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attribute_value_real, attribute_value_real.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongGetReferenceIdentifier(stage *Stage) string {
	return attribute_value_real.GongGetIdentifier(stage)
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attribute_value_string, attribute_value_string.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongGetReferenceIdentifier(stage *Stage) string {
	return attribute_value_string.GongGetIdentifier(stage)
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attribute_value_xhtml, attribute_value_xhtml.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongGetReferenceIdentifier(stage *Stage) string {
	return attribute_value_xhtml.GongGetIdentifier(stage)
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_alternative_id, a_alternative_id.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_alternative_id *A_ALTERNATIVE_ID) GongGetReferenceIdentifier(stage *Stage) string {
	return a_alternative_id.GongGetIdentifier(stage)
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_attribute_definition_boolean_ref, a_attribute_definition_boolean_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_attribute_definition_boolean_ref.GongGetIdentifier(stage)
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_attribute_definition_date_ref, a_attribute_definition_date_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_attribute_definition_date_ref.GongGetIdentifier(stage)
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_attribute_definition_enumeration_ref, a_attribute_definition_enumeration_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_attribute_definition_enumeration_ref.GongGetIdentifier(stage)
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_attribute_definition_integer_ref, a_attribute_definition_integer_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_attribute_definition_integer_ref.GongGetIdentifier(stage)
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_attribute_definition_real_ref, a_attribute_definition_real_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_attribute_definition_real_ref.GongGetIdentifier(stage)
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_attribute_definition_string_ref, a_attribute_definition_string_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_attribute_definition_string_ref.GongGetIdentifier(stage)
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_attribute_definition_xhtml_ref, a_attribute_definition_xhtml_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_attribute_definition_xhtml_ref.GongGetIdentifier(stage)
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_attribute_value_boolean, a_attribute_value_boolean.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongGetReferenceIdentifier(stage *Stage) string {
	return a_attribute_value_boolean.GongGetIdentifier(stage)
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_attribute_value_date, a_attribute_value_date.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongGetReferenceIdentifier(stage *Stage) string {
	return a_attribute_value_date.GongGetIdentifier(stage)
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_attribute_value_enumeration, a_attribute_value_enumeration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongGetReferenceIdentifier(stage *Stage) string {
	return a_attribute_value_enumeration.GongGetIdentifier(stage)
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_attribute_value_integer, a_attribute_value_integer.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongGetReferenceIdentifier(stage *Stage) string {
	return a_attribute_value_integer.GongGetIdentifier(stage)
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_attribute_value_real, a_attribute_value_real.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongGetReferenceIdentifier(stage *Stage) string {
	return a_attribute_value_real.GongGetIdentifier(stage)
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_attribute_value_string, a_attribute_value_string.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongGetReferenceIdentifier(stage *Stage) string {
	return a_attribute_value_string.GongGetIdentifier(stage)
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_attribute_value_xhtml, a_attribute_value_xhtml.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongGetReferenceIdentifier(stage *Stage) string {
	return a_attribute_value_xhtml.GongGetIdentifier(stage)
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_attribute_value_xhtml_1, a_attribute_value_xhtml_1.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongGetReferenceIdentifier(stage *Stage) string {
	return a_attribute_value_xhtml_1.GongGetIdentifier(stage)
}

func (a_children *A_CHILDREN) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_children, a_children.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_children *A_CHILDREN) GongGetReferenceIdentifier(stage *Stage) string {
	return a_children.GongGetIdentifier(stage)
}

func (a_core_content *A_CORE_CONTENT) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_core_content, a_core_content.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_core_content *A_CORE_CONTENT) GongGetReferenceIdentifier(stage *Stage) string {
	return a_core_content.GongGetIdentifier(stage)
}

func (a_datatypes *A_DATATYPES) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_datatypes, a_datatypes.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatypes *A_DATATYPES) GongGetReferenceIdentifier(stage *Stage) string {
	return a_datatypes.GongGetIdentifier(stage)
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_datatype_definition_boolean_ref, a_datatype_definition_boolean_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_datatype_definition_boolean_ref.GongGetIdentifier(stage)
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_datatype_definition_date_ref, a_datatype_definition_date_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_datatype_definition_date_ref.GongGetIdentifier(stage)
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_datatype_definition_enumeration_ref, a_datatype_definition_enumeration_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_datatype_definition_enumeration_ref.GongGetIdentifier(stage)
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_datatype_definition_integer_ref, a_datatype_definition_integer_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_datatype_definition_integer_ref.GongGetIdentifier(stage)
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_datatype_definition_real_ref, a_datatype_definition_real_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_datatype_definition_real_ref.GongGetIdentifier(stage)
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_datatype_definition_string_ref, a_datatype_definition_string_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_datatype_definition_string_ref.GongGetIdentifier(stage)
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_datatype_definition_xhtml_ref, a_datatype_definition_xhtml_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_datatype_definition_xhtml_ref.GongGetIdentifier(stage)
}

func (a_editable_atts *A_EDITABLE_ATTS) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_editable_atts, a_editable_atts.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_editable_atts *A_EDITABLE_ATTS) GongGetReferenceIdentifier(stage *Stage) string {
	return a_editable_atts.GongGetIdentifier(stage)
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_enum_value_ref, a_enum_value_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_enum_value_ref *A_ENUM_VALUE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_enum_value_ref.GongGetIdentifier(stage)
}

func (a_object *A_OBJECT) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_object, a_object.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_object *A_OBJECT) GongGetReferenceIdentifier(stage *Stage) string {
	return a_object.GongGetIdentifier(stage)
}

func (a_properties *A_PROPERTIES) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_properties, a_properties.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_properties *A_PROPERTIES) GongGetReferenceIdentifier(stage *Stage) string {
	return a_properties.GongGetIdentifier(stage)
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_relation_group_type_ref, a_relation_group_type_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_relation_group_type_ref.GongGetIdentifier(stage)
}

func (a_source_1 *A_SOURCE_1) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_source_1, a_source_1.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_source_1 *A_SOURCE_1) GongGetReferenceIdentifier(stage *Stage) string {
	return a_source_1.GongGetIdentifier(stage)
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_source_specification_1, a_source_specification_1.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongGetReferenceIdentifier(stage *Stage) string {
	return a_source_specification_1.GongGetIdentifier(stage)
}

func (a_specifications *A_SPECIFICATIONS) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_specifications, a_specifications.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_specifications *A_SPECIFICATIONS) GongGetReferenceIdentifier(stage *Stage) string {
	return a_specifications.GongGetIdentifier(stage)
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_specification_type_ref, a_specification_type_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_specification_type_ref.GongGetIdentifier(stage)
}

func (a_specified_values *A_SPECIFIED_VALUES) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_specified_values, a_specified_values.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_specified_values *A_SPECIFIED_VALUES) GongGetReferenceIdentifier(stage *Stage) string {
	return a_specified_values.GongGetIdentifier(stage)
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_spec_attributes, a_spec_attributes.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongGetReferenceIdentifier(stage *Stage) string {
	return a_spec_attributes.GongGetIdentifier(stage)
}

func (a_spec_objects *A_SPEC_OBJECTS) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_spec_objects, a_spec_objects.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_objects *A_SPEC_OBJECTS) GongGetReferenceIdentifier(stage *Stage) string {
	return a_spec_objects.GongGetIdentifier(stage)
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_spec_object_type_ref, a_spec_object_type_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_spec_object_type_ref.GongGetIdentifier(stage)
}

func (a_spec_relations *A_SPEC_RELATIONS) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_spec_relations, a_spec_relations.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_relations *A_SPEC_RELATIONS) GongGetReferenceIdentifier(stage *Stage) string {
	return a_spec_relations.GongGetIdentifier(stage)
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_spec_relation_groups, a_spec_relation_groups.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongGetReferenceIdentifier(stage *Stage) string {
	return a_spec_relation_groups.GongGetIdentifier(stage)
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_spec_relation_ref, a_spec_relation_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_spec_relation_ref.GongGetIdentifier(stage)
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_spec_relation_type_ref, a_spec_relation_type_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return a_spec_relation_type_ref.GongGetIdentifier(stage)
}

func (a_spec_types *A_SPEC_TYPES) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_spec_types, a_spec_types.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_types *A_SPEC_TYPES) GongGetReferenceIdentifier(stage *Stage) string {
	return a_spec_types.GongGetIdentifier(stage)
}

func (a_the_header *A_THE_HEADER) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_the_header, a_the_header.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_the_header *A_THE_HEADER) GongGetReferenceIdentifier(stage *Stage) string {
	return a_the_header.GongGetIdentifier(stage)
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_tool_extensions, a_tool_extensions.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_tool_extensions *A_TOOL_EXTENSIONS) GongGetReferenceIdentifier(stage *Stage) string {
	return a_tool_extensions.GongGetIdentifier(stage)
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(datatype_definition_boolean, datatype_definition_boolean.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongGetReferenceIdentifier(stage *Stage) string {
	return datatype_definition_boolean.GongGetIdentifier(stage)
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(datatype_definition_date, datatype_definition_date.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongGetReferenceIdentifier(stage *Stage) string {
	return datatype_definition_date.GongGetIdentifier(stage)
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(datatype_definition_enumeration, datatype_definition_enumeration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongGetReferenceIdentifier(stage *Stage) string {
	return datatype_definition_enumeration.GongGetIdentifier(stage)
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(datatype_definition_integer, datatype_definition_integer.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongGetReferenceIdentifier(stage *Stage) string {
	return datatype_definition_integer.GongGetIdentifier(stage)
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(datatype_definition_real, datatype_definition_real.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongGetReferenceIdentifier(stage *Stage) string {
	return datatype_definition_real.GongGetIdentifier(stage)
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(datatype_definition_string, datatype_definition_string.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongGetReferenceIdentifier(stage *Stage) string {
	return datatype_definition_string.GongGetIdentifier(stage)
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(datatype_definition_xhtml, datatype_definition_xhtml.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongGetReferenceIdentifier(stage *Stage) string {
	return datatype_definition_xhtml.GongGetIdentifier(stage)
}

func (embedded_value *EMBEDDED_VALUE) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(embedded_value, embedded_value.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (embedded_value *EMBEDDED_VALUE) GongGetReferenceIdentifier(stage *Stage) string {
	return embedded_value.GongGetIdentifier(stage)
}

func (enum_value *ENUM_VALUE) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(enum_value, enum_value.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (enum_value *ENUM_VALUE) GongGetReferenceIdentifier(stage *Stage) string {
	return enum_value.GongGetIdentifier(stage)
}

func (relation_group *RELATION_GROUP) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(relation_group, relation_group.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (relation_group *RELATION_GROUP) GongGetReferenceIdentifier(stage *Stage) string {
	return relation_group.GongGetIdentifier(stage)
}

func (relation_group_type *RELATION_GROUP_TYPE) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(relation_group_type, relation_group_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (relation_group_type *RELATION_GROUP_TYPE) GongGetReferenceIdentifier(stage *Stage) string {
	return relation_group_type.GongGetIdentifier(stage)
}

func (req_if *REQ_IF) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(req_if, req_if.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (req_if *REQ_IF) GongGetReferenceIdentifier(stage *Stage) string {
	return req_if.GongGetIdentifier(stage)
}

func (req_if_content *REQ_IF_CONTENT) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(req_if_content, req_if_content.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (req_if_content *REQ_IF_CONTENT) GongGetReferenceIdentifier(stage *Stage) string {
	return req_if_content.GongGetIdentifier(stage)
}

func (req_if_header *REQ_IF_HEADER) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(req_if_header, req_if_header.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (req_if_header *REQ_IF_HEADER) GongGetReferenceIdentifier(stage *Stage) string {
	return req_if_header.GongGetIdentifier(stage)
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(req_if_tool_extension, req_if_tool_extension.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongGetReferenceIdentifier(stage *Stage) string {
	return req_if_tool_extension.GongGetIdentifier(stage)
}

func (specification *SPECIFICATION) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(specification, specification.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specification *SPECIFICATION) GongGetReferenceIdentifier(stage *Stage) string {
	return specification.GongGetIdentifier(stage)
}

func (specification_type *SPECIFICATION_TYPE) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(specification_type, specification_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specification_type *SPECIFICATION_TYPE) GongGetReferenceIdentifier(stage *Stage) string {
	return specification_type.GongGetIdentifier(stage)
}

func (spec_hierarchy *SPEC_HIERARCHY) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(spec_hierarchy, spec_hierarchy.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spec_hierarchy *SPEC_HIERARCHY) GongGetReferenceIdentifier(stage *Stage) string {
	return spec_hierarchy.GongGetIdentifier(stage)
}

func (spec_object *SPEC_OBJECT) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(spec_object, spec_object.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spec_object *SPEC_OBJECT) GongGetReferenceIdentifier(stage *Stage) string {
	return spec_object.GongGetIdentifier(stage)
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(spec_object_type, spec_object_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spec_object_type *SPEC_OBJECT_TYPE) GongGetReferenceIdentifier(stage *Stage) string {
	return spec_object_type.GongGetIdentifier(stage)
}

func (spec_relation *SPEC_RELATION) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(spec_relation, spec_relation.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spec_relation *SPEC_RELATION) GongGetReferenceIdentifier(stage *Stage) string {
	return spec_relation.GongGetIdentifier(stage)
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(spec_relation_type, spec_relation_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spec_relation_type *SPEC_RELATION_TYPE) GongGetReferenceIdentifier(stage *Stage) string {
	return spec_relation_type.GongGetIdentifier(stage)
}

func (xhtml_content *XHTML_CONTENT) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(xhtml_content, xhtml_content.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (xhtml_content *XHTML_CONTENT) GongGetReferenceIdentifier(stage *Stage) string {
	return xhtml_content.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (alternative_id *ALTERNATIVE_ID) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(alternative_id.GongGetIdentifier(stage), "ALTERNATIVE_ID", alternative_id.Name)
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attribute_definition_boolean.GongGetIdentifier(stage), "ATTRIBUTE_DEFINITION_BOOLEAN", attribute_definition_boolean.Name)
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attribute_definition_date.GongGetIdentifier(stage), "ATTRIBUTE_DEFINITION_DATE", attribute_definition_date.Name)
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attribute_definition_enumeration.GongGetIdentifier(stage), "ATTRIBUTE_DEFINITION_ENUMERATION", attribute_definition_enumeration.Name)
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attribute_definition_integer.GongGetIdentifier(stage), "ATTRIBUTE_DEFINITION_INTEGER", attribute_definition_integer.Name)
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attribute_definition_real.GongGetIdentifier(stage), "ATTRIBUTE_DEFINITION_REAL", attribute_definition_real.Name)
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attribute_definition_string.GongGetIdentifier(stage), "ATTRIBUTE_DEFINITION_STRING", attribute_definition_string.Name)
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attribute_definition_xhtml.GongGetIdentifier(stage), "ATTRIBUTE_DEFINITION_XHTML", attribute_definition_xhtml.Name)
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attribute_value_boolean.GongGetIdentifier(stage), "ATTRIBUTE_VALUE_BOOLEAN", attribute_value_boolean.Name)
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attribute_value_date.GongGetIdentifier(stage), "ATTRIBUTE_VALUE_DATE", attribute_value_date.Name)
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attribute_value_enumeration.GongGetIdentifier(stage), "ATTRIBUTE_VALUE_ENUMERATION", attribute_value_enumeration.Name)
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attribute_value_integer.GongGetIdentifier(stage), "ATTRIBUTE_VALUE_INTEGER", attribute_value_integer.Name)
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attribute_value_real.GongGetIdentifier(stage), "ATTRIBUTE_VALUE_REAL", attribute_value_real.Name)
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attribute_value_string.GongGetIdentifier(stage), "ATTRIBUTE_VALUE_STRING", attribute_value_string.Name)
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attribute_value_xhtml.GongGetIdentifier(stage), "ATTRIBUTE_VALUE_XHTML", attribute_value_xhtml.Name)
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_alternative_id.GongGetIdentifier(stage), "A_ALTERNATIVE_ID", a_alternative_id.Name)
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_attribute_definition_boolean_ref.GongGetIdentifier(stage), "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF", a_attribute_definition_boolean_ref.Name)
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_attribute_definition_date_ref.GongGetIdentifier(stage), "A_ATTRIBUTE_DEFINITION_DATE_REF", a_attribute_definition_date_ref.Name)
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_attribute_definition_enumeration_ref.GongGetIdentifier(stage), "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF", a_attribute_definition_enumeration_ref.Name)
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_attribute_definition_integer_ref.GongGetIdentifier(stage), "A_ATTRIBUTE_DEFINITION_INTEGER_REF", a_attribute_definition_integer_ref.Name)
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_attribute_definition_real_ref.GongGetIdentifier(stage), "A_ATTRIBUTE_DEFINITION_REAL_REF", a_attribute_definition_real_ref.Name)
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_attribute_definition_string_ref.GongGetIdentifier(stage), "A_ATTRIBUTE_DEFINITION_STRING_REF", a_attribute_definition_string_ref.Name)
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_attribute_definition_xhtml_ref.GongGetIdentifier(stage), "A_ATTRIBUTE_DEFINITION_XHTML_REF", a_attribute_definition_xhtml_ref.Name)
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_attribute_value_boolean.GongGetIdentifier(stage), "A_ATTRIBUTE_VALUE_BOOLEAN", a_attribute_value_boolean.Name)
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_attribute_value_date.GongGetIdentifier(stage), "A_ATTRIBUTE_VALUE_DATE", a_attribute_value_date.Name)
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_attribute_value_enumeration.GongGetIdentifier(stage), "A_ATTRIBUTE_VALUE_ENUMERATION", a_attribute_value_enumeration.Name)
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_attribute_value_integer.GongGetIdentifier(stage), "A_ATTRIBUTE_VALUE_INTEGER", a_attribute_value_integer.Name)
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_attribute_value_real.GongGetIdentifier(stage), "A_ATTRIBUTE_VALUE_REAL", a_attribute_value_real.Name)
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_attribute_value_string.GongGetIdentifier(stage), "A_ATTRIBUTE_VALUE_STRING", a_attribute_value_string.Name)
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_attribute_value_xhtml.GongGetIdentifier(stage), "A_ATTRIBUTE_VALUE_XHTML", a_attribute_value_xhtml.Name)
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_attribute_value_xhtml_1.GongGetIdentifier(stage), "A_ATTRIBUTE_VALUE_XHTML_1", a_attribute_value_xhtml_1.Name)
}

func (a_children *A_CHILDREN) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_children.GongGetIdentifier(stage), "A_CHILDREN", a_children.Name)
}

func (a_core_content *A_CORE_CONTENT) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_core_content.GongGetIdentifier(stage), "A_CORE_CONTENT", a_core_content.Name)
}

func (a_datatypes *A_DATATYPES) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_datatypes.GongGetIdentifier(stage), "A_DATATYPES", a_datatypes.Name)
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_datatype_definition_boolean_ref.GongGetIdentifier(stage), "A_DATATYPE_DEFINITION_BOOLEAN_REF", a_datatype_definition_boolean_ref.Name)
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_datatype_definition_date_ref.GongGetIdentifier(stage), "A_DATATYPE_DEFINITION_DATE_REF", a_datatype_definition_date_ref.Name)
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_datatype_definition_enumeration_ref.GongGetIdentifier(stage), "A_DATATYPE_DEFINITION_ENUMERATION_REF", a_datatype_definition_enumeration_ref.Name)
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_datatype_definition_integer_ref.GongGetIdentifier(stage), "A_DATATYPE_DEFINITION_INTEGER_REF", a_datatype_definition_integer_ref.Name)
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_datatype_definition_real_ref.GongGetIdentifier(stage), "A_DATATYPE_DEFINITION_REAL_REF", a_datatype_definition_real_ref.Name)
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_datatype_definition_string_ref.GongGetIdentifier(stage), "A_DATATYPE_DEFINITION_STRING_REF", a_datatype_definition_string_ref.Name)
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_datatype_definition_xhtml_ref.GongGetIdentifier(stage), "A_DATATYPE_DEFINITION_XHTML_REF", a_datatype_definition_xhtml_ref.Name)
}

func (a_editable_atts *A_EDITABLE_ATTS) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_editable_atts.GongGetIdentifier(stage), "A_EDITABLE_ATTS", a_editable_atts.Name)
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_enum_value_ref.GongGetIdentifier(stage), "A_ENUM_VALUE_REF", a_enum_value_ref.Name)
}

func (a_object *A_OBJECT) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_object.GongGetIdentifier(stage), "A_OBJECT", a_object.Name)
}

func (a_properties *A_PROPERTIES) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_properties.GongGetIdentifier(stage), "A_PROPERTIES", a_properties.Name)
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_relation_group_type_ref.GongGetIdentifier(stage), "A_RELATION_GROUP_TYPE_REF", a_relation_group_type_ref.Name)
}

func (a_source_1 *A_SOURCE_1) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_source_1.GongGetIdentifier(stage), "A_SOURCE_1", a_source_1.Name)
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_source_specification_1.GongGetIdentifier(stage), "A_SOURCE_SPECIFICATION_1", a_source_specification_1.Name)
}

func (a_specifications *A_SPECIFICATIONS) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_specifications.GongGetIdentifier(stage), "A_SPECIFICATIONS", a_specifications.Name)
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_specification_type_ref.GongGetIdentifier(stage), "A_SPECIFICATION_TYPE_REF", a_specification_type_ref.Name)
}

func (a_specified_values *A_SPECIFIED_VALUES) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_specified_values.GongGetIdentifier(stage), "A_SPECIFIED_VALUES", a_specified_values.Name)
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_spec_attributes.GongGetIdentifier(stage), "A_SPEC_ATTRIBUTES", a_spec_attributes.Name)
}

func (a_spec_objects *A_SPEC_OBJECTS) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_spec_objects.GongGetIdentifier(stage), "A_SPEC_OBJECTS", a_spec_objects.Name)
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_spec_object_type_ref.GongGetIdentifier(stage), "A_SPEC_OBJECT_TYPE_REF", a_spec_object_type_ref.Name)
}

func (a_spec_relations *A_SPEC_RELATIONS) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_spec_relations.GongGetIdentifier(stage), "A_SPEC_RELATIONS", a_spec_relations.Name)
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_spec_relation_groups.GongGetIdentifier(stage), "A_SPEC_RELATION_GROUPS", a_spec_relation_groups.Name)
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_spec_relation_ref.GongGetIdentifier(stage), "A_SPEC_RELATION_REF", a_spec_relation_ref.Name)
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_spec_relation_type_ref.GongGetIdentifier(stage), "A_SPEC_RELATION_TYPE_REF", a_spec_relation_type_ref.Name)
}

func (a_spec_types *A_SPEC_TYPES) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_spec_types.GongGetIdentifier(stage), "A_SPEC_TYPES", a_spec_types.Name)
}

func (a_the_header *A_THE_HEADER) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_the_header.GongGetIdentifier(stage), "A_THE_HEADER", a_the_header.Name)
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_tool_extensions.GongGetIdentifier(stage), "A_TOOL_EXTENSIONS", a_tool_extensions.Name)
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(datatype_definition_boolean.GongGetIdentifier(stage), "DATATYPE_DEFINITION_BOOLEAN", datatype_definition_boolean.Name)
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(datatype_definition_date.GongGetIdentifier(stage), "DATATYPE_DEFINITION_DATE", datatype_definition_date.Name)
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(datatype_definition_enumeration.GongGetIdentifier(stage), "DATATYPE_DEFINITION_ENUMERATION", datatype_definition_enumeration.Name)
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(datatype_definition_integer.GongGetIdentifier(stage), "DATATYPE_DEFINITION_INTEGER", datatype_definition_integer.Name)
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(datatype_definition_real.GongGetIdentifier(stage), "DATATYPE_DEFINITION_REAL", datatype_definition_real.Name)
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(datatype_definition_string.GongGetIdentifier(stage), "DATATYPE_DEFINITION_STRING", datatype_definition_string.Name)
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(datatype_definition_xhtml.GongGetIdentifier(stage), "DATATYPE_DEFINITION_XHTML", datatype_definition_xhtml.Name)
}

func (embedded_value *EMBEDDED_VALUE) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(embedded_value.GongGetIdentifier(stage), "EMBEDDED_VALUE", embedded_value.Name)
}

func (enum_value *ENUM_VALUE) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(enum_value.GongGetIdentifier(stage), "ENUM_VALUE", enum_value.Name)
}

func (relation_group *RELATION_GROUP) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(relation_group.GongGetIdentifier(stage), "RELATION_GROUP", relation_group.Name)
}

func (relation_group_type *RELATION_GROUP_TYPE) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(relation_group_type.GongGetIdentifier(stage), "RELATION_GROUP_TYPE", relation_group_type.Name)
}

func (req_if *REQ_IF) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(req_if.GongGetIdentifier(stage), "REQ_IF", req_if.Name)
}

func (req_if_content *REQ_IF_CONTENT) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(req_if_content.GongGetIdentifier(stage), "REQ_IF_CONTENT", req_if_content.Name)
}

func (req_if_header *REQ_IF_HEADER) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(req_if_header.GongGetIdentifier(stage), "REQ_IF_HEADER", req_if_header.Name)
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(req_if_tool_extension.GongGetIdentifier(stage), "REQ_IF_TOOL_EXTENSION", req_if_tool_extension.Name)
}

func (specification *SPECIFICATION) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(specification.GongGetIdentifier(stage), "SPECIFICATION", specification.Name)
}

func (specification_type *SPECIFICATION_TYPE) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(specification_type.GongGetIdentifier(stage), "SPECIFICATION_TYPE", specification_type.Name)
}

func (spec_hierarchy *SPEC_HIERARCHY) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(spec_hierarchy.GongGetIdentifier(stage), "SPEC_HIERARCHY", spec_hierarchy.Name)
}

func (spec_object *SPEC_OBJECT) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(spec_object.GongGetIdentifier(stage), "SPEC_OBJECT", spec_object.Name)
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(spec_object_type.GongGetIdentifier(stage), "SPEC_OBJECT_TYPE", spec_object_type.Name)
}

func (spec_relation *SPEC_RELATION) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(spec_relation.GongGetIdentifier(stage), "SPEC_RELATION", spec_relation.Name)
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(spec_relation_type.GongGetIdentifier(stage), "SPEC_RELATION_TYPE", spec_relation_type.Name)
}

func (xhtml_content *XHTML_CONTENT) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(xhtml_content.GongGetIdentifier(stage), "XHTML_CONTENT", xhtml_content.Name)
}

// insertion point for unstaging
func (alternative_id *ALTERNATIVE_ID) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(alternative_id.GongGetReferenceIdentifier(stage))
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attribute_definition_boolean.GongGetReferenceIdentifier(stage))
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attribute_definition_date.GongGetReferenceIdentifier(stage))
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attribute_definition_enumeration.GongGetReferenceIdentifier(stage))
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attribute_definition_integer.GongGetReferenceIdentifier(stage))
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attribute_definition_real.GongGetReferenceIdentifier(stage))
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attribute_definition_string.GongGetReferenceIdentifier(stage))
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attribute_definition_xhtml.GongGetReferenceIdentifier(stage))
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attribute_value_boolean.GongGetReferenceIdentifier(stage))
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attribute_value_date.GongGetReferenceIdentifier(stage))
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attribute_value_enumeration.GongGetReferenceIdentifier(stage))
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attribute_value_integer.GongGetReferenceIdentifier(stage))
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attribute_value_real.GongGetReferenceIdentifier(stage))
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attribute_value_string.GongGetReferenceIdentifier(stage))
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attribute_value_xhtml.GongGetReferenceIdentifier(stage))
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_alternative_id.GongGetReferenceIdentifier(stage))
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_attribute_definition_boolean_ref.GongGetReferenceIdentifier(stage))
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_attribute_definition_date_ref.GongGetReferenceIdentifier(stage))
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_attribute_definition_enumeration_ref.GongGetReferenceIdentifier(stage))
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_attribute_definition_integer_ref.GongGetReferenceIdentifier(stage))
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_attribute_definition_real_ref.GongGetReferenceIdentifier(stage))
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_attribute_definition_string_ref.GongGetReferenceIdentifier(stage))
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_attribute_definition_xhtml_ref.GongGetReferenceIdentifier(stage))
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_attribute_value_boolean.GongGetReferenceIdentifier(stage))
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_attribute_value_date.GongGetReferenceIdentifier(stage))
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_attribute_value_enumeration.GongGetReferenceIdentifier(stage))
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_attribute_value_integer.GongGetReferenceIdentifier(stage))
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_attribute_value_real.GongGetReferenceIdentifier(stage))
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_attribute_value_string.GongGetReferenceIdentifier(stage))
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_attribute_value_xhtml.GongGetReferenceIdentifier(stage))
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_attribute_value_xhtml_1.GongGetReferenceIdentifier(stage))
}

func (a_children *A_CHILDREN) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_children.GongGetReferenceIdentifier(stage))
}

func (a_core_content *A_CORE_CONTENT) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_core_content.GongGetReferenceIdentifier(stage))
}

func (a_datatypes *A_DATATYPES) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_datatypes.GongGetReferenceIdentifier(stage))
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_datatype_definition_boolean_ref.GongGetReferenceIdentifier(stage))
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_datatype_definition_date_ref.GongGetReferenceIdentifier(stage))
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_datatype_definition_enumeration_ref.GongGetReferenceIdentifier(stage))
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_datatype_definition_integer_ref.GongGetReferenceIdentifier(stage))
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_datatype_definition_real_ref.GongGetReferenceIdentifier(stage))
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_datatype_definition_string_ref.GongGetReferenceIdentifier(stage))
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_datatype_definition_xhtml_ref.GongGetReferenceIdentifier(stage))
}

func (a_editable_atts *A_EDITABLE_ATTS) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_editable_atts.GongGetReferenceIdentifier(stage))
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_enum_value_ref.GongGetReferenceIdentifier(stage))
}

func (a_object *A_OBJECT) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_object.GongGetReferenceIdentifier(stage))
}

func (a_properties *A_PROPERTIES) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_properties.GongGetReferenceIdentifier(stage))
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_relation_group_type_ref.GongGetReferenceIdentifier(stage))
}

func (a_source_1 *A_SOURCE_1) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_source_1.GongGetReferenceIdentifier(stage))
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_source_specification_1.GongGetReferenceIdentifier(stage))
}

func (a_specifications *A_SPECIFICATIONS) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_specifications.GongGetReferenceIdentifier(stage))
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_specification_type_ref.GongGetReferenceIdentifier(stage))
}

func (a_specified_values *A_SPECIFIED_VALUES) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_specified_values.GongGetReferenceIdentifier(stage))
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_spec_attributes.GongGetReferenceIdentifier(stage))
}

func (a_spec_objects *A_SPEC_OBJECTS) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_spec_objects.GongGetReferenceIdentifier(stage))
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_spec_object_type_ref.GongGetReferenceIdentifier(stage))
}

func (a_spec_relations *A_SPEC_RELATIONS) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_spec_relations.GongGetReferenceIdentifier(stage))
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_spec_relation_groups.GongGetReferenceIdentifier(stage))
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_spec_relation_ref.GongGetReferenceIdentifier(stage))
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_spec_relation_type_ref.GongGetReferenceIdentifier(stage))
}

func (a_spec_types *A_SPEC_TYPES) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_spec_types.GongGetReferenceIdentifier(stage))
}

func (a_the_header *A_THE_HEADER) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_the_header.GongGetReferenceIdentifier(stage))
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_tool_extensions.GongGetReferenceIdentifier(stage))
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(datatype_definition_boolean.GongGetReferenceIdentifier(stage))
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(datatype_definition_date.GongGetReferenceIdentifier(stage))
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(datatype_definition_enumeration.GongGetReferenceIdentifier(stage))
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(datatype_definition_integer.GongGetReferenceIdentifier(stage))
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(datatype_definition_real.GongGetReferenceIdentifier(stage))
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(datatype_definition_string.GongGetReferenceIdentifier(stage))
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(datatype_definition_xhtml.GongGetReferenceIdentifier(stage))
}

func (embedded_value *EMBEDDED_VALUE) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(embedded_value.GongGetReferenceIdentifier(stage))
}

func (enum_value *ENUM_VALUE) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(enum_value.GongGetReferenceIdentifier(stage))
}

func (relation_group *RELATION_GROUP) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(relation_group.GongGetReferenceIdentifier(stage))
}

func (relation_group_type *RELATION_GROUP_TYPE) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(relation_group_type.GongGetReferenceIdentifier(stage))
}

func (req_if *REQ_IF) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(req_if.GongGetReferenceIdentifier(stage))
}

func (req_if_content *REQ_IF_CONTENT) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(req_if_content.GongGetReferenceIdentifier(stage))
}

func (req_if_header *REQ_IF_HEADER) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(req_if_header.GongGetReferenceIdentifier(stage))
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(req_if_tool_extension.GongGetReferenceIdentifier(stage))
}

func (specification *SPECIFICATION) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(specification.GongGetReferenceIdentifier(stage))
}

func (specification_type *SPECIFICATION_TYPE) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(specification_type.GongGetReferenceIdentifier(stage))
}

func (spec_hierarchy *SPEC_HIERARCHY) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(spec_hierarchy.GongGetReferenceIdentifier(stage))
}

func (spec_object *SPEC_OBJECT) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(spec_object.GongGetReferenceIdentifier(stage))
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(spec_object_type.GongGetReferenceIdentifier(stage))
}

func (spec_relation *SPEC_RELATION) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(spec_relation.GongGetReferenceIdentifier(stage))
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(spec_relation_type.GongGetReferenceIdentifier(stage))
}

func (xhtml_content *XHTML_CONTENT) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(xhtml_content.GongGetReferenceIdentifier(stage))
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
