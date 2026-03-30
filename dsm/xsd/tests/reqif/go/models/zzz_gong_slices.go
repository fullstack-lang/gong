// generated code - do not edit
package models

import (
	"fmt"
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
	// Compute reverse map for named struct ALTERNATIVE_ID
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_BOOLEAN
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_DATE
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_ENUMERATION
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_INTEGER
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_REAL
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_STRING
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_XHTML
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_BOOLEAN
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_DATE
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_ENUMERATION
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_INTEGER
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_REAL
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_STRING
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_VALUE_XHTML
	// insertion point per field

	// Compute reverse map for named struct A_ALTERNATIVE_ID
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_BOOLEAN_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_DATE_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_ENUMERATION_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_INTEGER_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_REAL_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_STRING_REF
	// insertion point per field

	// Compute reverse map for named struct A_ATTRIBUTE_DEFINITION_XHTML_REF
	// insertion point per field

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

	// Compute reverse map for named struct A_CORE_CONTENT
	// insertion point per field

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

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_BOOLEAN_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_DATE_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_ENUMERATION_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_INTEGER_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_REAL_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_STRING_REF
	// insertion point per field

	// Compute reverse map for named struct A_DATATYPE_DEFINITION_XHTML_REF
	// insertion point per field

	// Compute reverse map for named struct A_EDITABLE_ATTS
	// insertion point per field

	// Compute reverse map for named struct A_ENUM_VALUE_REF
	// insertion point per field

	// Compute reverse map for named struct A_OBJECT
	// insertion point per field

	// Compute reverse map for named struct A_PROPERTIES
	// insertion point per field

	// Compute reverse map for named struct A_RELATION_GROUP_TYPE_REF
	// insertion point per field

	// Compute reverse map for named struct A_SOURCE_1
	// insertion point per field

	// Compute reverse map for named struct A_SOURCE_SPECIFICATION_1
	// insertion point per field

	// Compute reverse map for named struct A_SPECIFICATIONS
	// insertion point per field
	stage.A_SPECIFICATIONS_SPECIFICATION_reverseMap = make(map[*SPECIFICATION]*A_SPECIFICATIONS)
	for a_specifications := range stage.A_SPECIFICATIONSs {
		_ = a_specifications
		for _, _specification := range a_specifications.SPECIFICATION {
			stage.A_SPECIFICATIONS_SPECIFICATION_reverseMap[_specification] = a_specifications
		}
	}

	// Compute reverse map for named struct A_SPECIFICATION_TYPE_REF
	// insertion point per field

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

	// Compute reverse map for named struct A_SPEC_OBJECT_TYPE_REF
	// insertion point per field

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

	// Compute reverse map for named struct A_SPEC_RELATION_REF
	// insertion point per field

	// Compute reverse map for named struct A_SPEC_RELATION_TYPE_REF
	// insertion point per field

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

	// Compute reverse map for named struct A_THE_HEADER
	// insertion point per field

	// Compute reverse map for named struct A_TOOL_EXTENSIONS
	// insertion point per field
	stage.A_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap = make(map[*REQ_IF_TOOL_EXTENSION]*A_TOOL_EXTENSIONS)
	for a_tool_extensions := range stage.A_TOOL_EXTENSIONSs {
		_ = a_tool_extensions
		for _, _req_if_tool_extension := range a_tool_extensions.REQ_IF_TOOL_EXTENSION {
			stage.A_TOOL_EXTENSIONS_REQ_IF_TOOL_EXTENSION_reverseMap[_req_if_tool_extension] = a_tool_extensions
		}
	}

	// Compute reverse map for named struct DATATYPE_DEFINITION_BOOLEAN
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_DATE
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_ENUMERATION
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_INTEGER
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_REAL
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_STRING
	// insertion point per field

	// Compute reverse map for named struct DATATYPE_DEFINITION_XHTML
	// insertion point per field

	// Compute reverse map for named struct EMBEDDED_VALUE
	// insertion point per field

	// Compute reverse map for named struct ENUM_VALUE
	// insertion point per field

	// Compute reverse map for named struct RELATION_GROUP
	// insertion point per field

	// Compute reverse map for named struct RELATION_GROUP_TYPE
	// insertion point per field

	// Compute reverse map for named struct REQ_IF
	// insertion point per field

	// Compute reverse map for named struct REQ_IF_CONTENT
	// insertion point per field

	// Compute reverse map for named struct REQ_IF_HEADER
	// insertion point per field

	// Compute reverse map for named struct REQ_IF_TOOL_EXTENSION
	// insertion point per field

	// Compute reverse map for named struct SPECIFICATION
	// insertion point per field

	// Compute reverse map for named struct SPECIFICATION_TYPE
	// insertion point per field

	// Compute reverse map for named struct SPEC_HIERARCHY
	// insertion point per field

	// Compute reverse map for named struct SPEC_OBJECT
	// insertion point per field

	// Compute reverse map for named struct SPEC_OBJECT_TYPE
	// insertion point per field

	// Compute reverse map for named struct SPEC_RELATION
	// insertion point per field

	// Compute reverse map for named struct SPEC_RELATION_TYPE
	// insertion point per field

	// Compute reverse map for named struct XHTML_CONTENT
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.ALTERNATIVE_IDs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_DATEs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_REALs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_BOOLEANs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_DATEs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_ENUMERATIONs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_INTEGERs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_REALs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_STRINGs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_XHTMLs {
		res = append(res, instance)
	}

	for instance := range stage.A_ALTERNATIVE_IDs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_DATE_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_REAL_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_STRING_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_BOOLEANs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_DATEs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_INTEGERs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_REALs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_STRINGs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
		res = append(res, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		res = append(res, instance)
	}

	for instance := range stage.A_CHILDRENs {
		res = append(res, instance)
	}

	for instance := range stage.A_CORE_CONTENTs {
		res = append(res, instance)
	}

	for instance := range stage.A_DATATYPESs {
		res = append(res, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_DATE_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_INTEGER_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_REAL_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_STRING_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_XHTML_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_EDITABLE_ATTSs {
		res = append(res, instance)
	}

	for instance := range stage.A_ENUM_VALUE_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_OBJECTs {
		res = append(res, instance)
	}

	for instance := range stage.A_PROPERTIESs {
		res = append(res, instance)
	}

	for instance := range stage.A_RELATION_GROUP_TYPE_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_SOURCE_1s {
		res = append(res, instance)
	}

	for instance := range stage.A_SOURCE_SPECIFICATION_1s {
		res = append(res, instance)
	}

	for instance := range stage.A_SPECIFICATIONSs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPECIFICATION_TYPE_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPECIFIED_VALUESs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPEC_ATTRIBUTESs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPEC_OBJECTSs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPEC_OBJECT_TYPE_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPEC_RELATIONSs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPEC_RELATION_GROUPSs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPEC_RELATION_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPEC_RELATION_TYPE_REFs {
		res = append(res, instance)
	}

	for instance := range stage.A_SPEC_TYPESs {
		res = append(res, instance)
	}

	for instance := range stage.A_THE_HEADERs {
		res = append(res, instance)
	}

	for instance := range stage.A_TOOL_EXTENSIONSs {
		res = append(res, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_BOOLEANs {
		res = append(res, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_DATEs {
		res = append(res, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		res = append(res, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_INTEGERs {
		res = append(res, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_REALs {
		res = append(res, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_STRINGs {
		res = append(res, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_XHTMLs {
		res = append(res, instance)
	}

	for instance := range stage.EMBEDDED_VALUEs {
		res = append(res, instance)
	}

	for instance := range stage.ENUM_VALUEs {
		res = append(res, instance)
	}

	for instance := range stage.RELATION_GROUPs {
		res = append(res, instance)
	}

	for instance := range stage.RELATION_GROUP_TYPEs {
		res = append(res, instance)
	}

	for instance := range stage.REQ_IFs {
		res = append(res, instance)
	}

	for instance := range stage.REQ_IF_CONTENTs {
		res = append(res, instance)
	}

	for instance := range stage.REQ_IF_HEADERs {
		res = append(res, instance)
	}

	for instance := range stage.REQ_IF_TOOL_EXTENSIONs {
		res = append(res, instance)
	}

	for instance := range stage.SPECIFICATIONs {
		res = append(res, instance)
	}

	for instance := range stage.SPECIFICATION_TYPEs {
		res = append(res, instance)
	}

	for instance := range stage.SPEC_HIERARCHYs {
		res = append(res, instance)
	}

	for instance := range stage.SPEC_OBJECTs {
		res = append(res, instance)
	}

	for instance := range stage.SPEC_OBJECT_TYPEs {
		res = append(res, instance)
	}

	for instance := range stage.SPEC_RELATIONs {
		res = append(res, instance)
	}

	for instance := range stage.SPEC_RELATION_TYPEs {
		res = append(res, instance)
	}

	for instance := range stage.XHTML_CONTENTs {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (alternative_id *ALTERNATIVE_ID) GongCopy() GongstructIF {
	newInstance := *alternative_id
	return &newInstance
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongCopy() GongstructIF {
	newInstance := *attribute_definition_boolean
	return &newInstance
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongCopy() GongstructIF {
	newInstance := *attribute_definition_date
	return &newInstance
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongCopy() GongstructIF {
	newInstance := *attribute_definition_enumeration
	return &newInstance
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongCopy() GongstructIF {
	newInstance := *attribute_definition_integer
	return &newInstance
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongCopy() GongstructIF {
	newInstance := *attribute_definition_real
	return &newInstance
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongCopy() GongstructIF {
	newInstance := *attribute_definition_string
	return &newInstance
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongCopy() GongstructIF {
	newInstance := *attribute_definition_xhtml
	return &newInstance
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongCopy() GongstructIF {
	newInstance := *attribute_value_boolean
	return &newInstance
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongCopy() GongstructIF {
	newInstance := *attribute_value_date
	return &newInstance
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongCopy() GongstructIF {
	newInstance := *attribute_value_enumeration
	return &newInstance
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongCopy() GongstructIF {
	newInstance := *attribute_value_integer
	return &newInstance
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongCopy() GongstructIF {
	newInstance := *attribute_value_real
	return &newInstance
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongCopy() GongstructIF {
	newInstance := *attribute_value_string
	return &newInstance
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongCopy() GongstructIF {
	newInstance := *attribute_value_xhtml
	return &newInstance
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongCopy() GongstructIF {
	newInstance := *a_alternative_id
	return &newInstance
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongCopy() GongstructIF {
	newInstance := *a_attribute_definition_boolean_ref
	return &newInstance
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongCopy() GongstructIF {
	newInstance := *a_attribute_definition_date_ref
	return &newInstance
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongCopy() GongstructIF {
	newInstance := *a_attribute_definition_enumeration_ref
	return &newInstance
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongCopy() GongstructIF {
	newInstance := *a_attribute_definition_integer_ref
	return &newInstance
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongCopy() GongstructIF {
	newInstance := *a_attribute_definition_real_ref
	return &newInstance
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongCopy() GongstructIF {
	newInstance := *a_attribute_definition_string_ref
	return &newInstance
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongCopy() GongstructIF {
	newInstance := *a_attribute_definition_xhtml_ref
	return &newInstance
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongCopy() GongstructIF {
	newInstance := *a_attribute_value_boolean
	return &newInstance
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongCopy() GongstructIF {
	newInstance := *a_attribute_value_date
	return &newInstance
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongCopy() GongstructIF {
	newInstance := *a_attribute_value_enumeration
	return &newInstance
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongCopy() GongstructIF {
	newInstance := *a_attribute_value_integer
	return &newInstance
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongCopy() GongstructIF {
	newInstance := *a_attribute_value_real
	return &newInstance
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongCopy() GongstructIF {
	newInstance := *a_attribute_value_string
	return &newInstance
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongCopy() GongstructIF {
	newInstance := *a_attribute_value_xhtml
	return &newInstance
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongCopy() GongstructIF {
	newInstance := *a_attribute_value_xhtml_1
	return &newInstance
}

func (a_children *A_CHILDREN) GongCopy() GongstructIF {
	newInstance := *a_children
	return &newInstance
}

func (a_core_content *A_CORE_CONTENT) GongCopy() GongstructIF {
	newInstance := *a_core_content
	return &newInstance
}

func (a_datatypes *A_DATATYPES) GongCopy() GongstructIF {
	newInstance := *a_datatypes
	return &newInstance
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongCopy() GongstructIF {
	newInstance := *a_datatype_definition_boolean_ref
	return &newInstance
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongCopy() GongstructIF {
	newInstance := *a_datatype_definition_date_ref
	return &newInstance
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongCopy() GongstructIF {
	newInstance := *a_datatype_definition_enumeration_ref
	return &newInstance
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongCopy() GongstructIF {
	newInstance := *a_datatype_definition_integer_ref
	return &newInstance
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongCopy() GongstructIF {
	newInstance := *a_datatype_definition_real_ref
	return &newInstance
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongCopy() GongstructIF {
	newInstance := *a_datatype_definition_string_ref
	return &newInstance
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongCopy() GongstructIF {
	newInstance := *a_datatype_definition_xhtml_ref
	return &newInstance
}

func (a_editable_atts *A_EDITABLE_ATTS) GongCopy() GongstructIF {
	newInstance := *a_editable_atts
	return &newInstance
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongCopy() GongstructIF {
	newInstance := *a_enum_value_ref
	return &newInstance
}

func (a_object *A_OBJECT) GongCopy() GongstructIF {
	newInstance := *a_object
	return &newInstance
}

func (a_properties *A_PROPERTIES) GongCopy() GongstructIF {
	newInstance := *a_properties
	return &newInstance
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongCopy() GongstructIF {
	newInstance := *a_relation_group_type_ref
	return &newInstance
}

func (a_source_1 *A_SOURCE_1) GongCopy() GongstructIF {
	newInstance := *a_source_1
	return &newInstance
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongCopy() GongstructIF {
	newInstance := *a_source_specification_1
	return &newInstance
}

func (a_specifications *A_SPECIFICATIONS) GongCopy() GongstructIF {
	newInstance := *a_specifications
	return &newInstance
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongCopy() GongstructIF {
	newInstance := *a_specification_type_ref
	return &newInstance
}

func (a_specified_values *A_SPECIFIED_VALUES) GongCopy() GongstructIF {
	newInstance := *a_specified_values
	return &newInstance
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongCopy() GongstructIF {
	newInstance := *a_spec_attributes
	return &newInstance
}

func (a_spec_objects *A_SPEC_OBJECTS) GongCopy() GongstructIF {
	newInstance := *a_spec_objects
	return &newInstance
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongCopy() GongstructIF {
	newInstance := *a_spec_object_type_ref
	return &newInstance
}

func (a_spec_relations *A_SPEC_RELATIONS) GongCopy() GongstructIF {
	newInstance := *a_spec_relations
	return &newInstance
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongCopy() GongstructIF {
	newInstance := *a_spec_relation_groups
	return &newInstance
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongCopy() GongstructIF {
	newInstance := *a_spec_relation_ref
	return &newInstance
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongCopy() GongstructIF {
	newInstance := *a_spec_relation_type_ref
	return &newInstance
}

func (a_spec_types *A_SPEC_TYPES) GongCopy() GongstructIF {
	newInstance := *a_spec_types
	return &newInstance
}

func (a_the_header *A_THE_HEADER) GongCopy() GongstructIF {
	newInstance := *a_the_header
	return &newInstance
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongCopy() GongstructIF {
	newInstance := *a_tool_extensions
	return &newInstance
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongCopy() GongstructIF {
	newInstance := *datatype_definition_boolean
	return &newInstance
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongCopy() GongstructIF {
	newInstance := *datatype_definition_date
	return &newInstance
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongCopy() GongstructIF {
	newInstance := *datatype_definition_enumeration
	return &newInstance
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongCopy() GongstructIF {
	newInstance := *datatype_definition_integer
	return &newInstance
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongCopy() GongstructIF {
	newInstance := *datatype_definition_real
	return &newInstance
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongCopy() GongstructIF {
	newInstance := *datatype_definition_string
	return &newInstance
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongCopy() GongstructIF {
	newInstance := *datatype_definition_xhtml
	return &newInstance
}

func (embedded_value *EMBEDDED_VALUE) GongCopy() GongstructIF {
	newInstance := *embedded_value
	return &newInstance
}

func (enum_value *ENUM_VALUE) GongCopy() GongstructIF {
	newInstance := *enum_value
	return &newInstance
}

func (relation_group *RELATION_GROUP) GongCopy() GongstructIF {
	newInstance := *relation_group
	return &newInstance
}

func (relation_group_type *RELATION_GROUP_TYPE) GongCopy() GongstructIF {
	newInstance := *relation_group_type
	return &newInstance
}

func (req_if *REQ_IF) GongCopy() GongstructIF {
	newInstance := *req_if
	return &newInstance
}

func (req_if_content *REQ_IF_CONTENT) GongCopy() GongstructIF {
	newInstance := *req_if_content
	return &newInstance
}

func (req_if_header *REQ_IF_HEADER) GongCopy() GongstructIF {
	newInstance := *req_if_header
	return &newInstance
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongCopy() GongstructIF {
	newInstance := *req_if_tool_extension
	return &newInstance
}

func (specification *SPECIFICATION) GongCopy() GongstructIF {
	newInstance := *specification
	return &newInstance
}

func (specification_type *SPECIFICATION_TYPE) GongCopy() GongstructIF {
	newInstance := *specification_type
	return &newInstance
}

func (spec_hierarchy *SPEC_HIERARCHY) GongCopy() GongstructIF {
	newInstance := *spec_hierarchy
	return &newInstance
}

func (spec_object *SPEC_OBJECT) GongCopy() GongstructIF {
	newInstance := *spec_object
	return &newInstance
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongCopy() GongstructIF {
	newInstance := *spec_object_type
	return &newInstance
}

func (spec_relation *SPEC_RELATION) GongCopy() GongstructIF {
	newInstance := *spec_relation
	return &newInstance
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongCopy() GongstructIF {
	newInstance := *spec_relation_type
	return &newInstance
}

func (xhtml_content *XHTML_CONTENT) GongCopy() GongstructIF {
	newInstance := *xhtml_content
	return &newInstance
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
	var alternative_ids_newInstances []*ALTERNATIVE_ID
	var alternative_ids_deletedInstances []*ALTERNATIVE_ID

	// parse all staged instances and check if they have a reference
	for alternative_id := range stage.ALTERNATIVE_IDs {
		if ref, ok := stage.ALTERNATIVE_IDs_reference[alternative_id]; !ok {
			alternative_ids_newInstances = append(alternative_ids_newInstances, alternative_id)
			newInstancesSlice = append(newInstancesSlice, alternative_id.GongMarshallIdentifier(stage))
			if stage.ALTERNATIVE_IDs_referenceOrder == nil {
				stage.ALTERNATIVE_IDs_referenceOrder = make(map[*ALTERNATIVE_ID]uint)
			}
			stage.ALTERNATIVE_IDs_referenceOrder[alternative_id] = stage.ALTERNATIVE_IDMap_Staged_Order[alternative_id]
			newInstancesReverseSlice = append(newInstancesReverseSlice, alternative_id.GongMarshallUnstaging(stage))
			delete(stage.ALTERNATIVE_IDs_referenceOrder, alternative_id)
			fieldInitializers, pointersInitializations := alternative_id.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ALTERNATIVE_IDMap_Staged_Order[ref] = stage.ALTERNATIVE_IDMap_Staged_Order[alternative_id]
			diffs := alternative_id.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, alternative_id)
			delete(stage.ALTERNATIVE_IDMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", alternative_id.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ALTERNATIVE_IDs_reference {
		if _, ok := stage.ALTERNATIVE_IDs[ref]; !ok {
			alternative_ids_deletedInstances = append(alternative_ids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(alternative_ids_newInstances)
	lenDeletedInstances += len(alternative_ids_deletedInstances)
	var attribute_definition_booleans_newInstances []*ATTRIBUTE_DEFINITION_BOOLEAN
	var attribute_definition_booleans_deletedInstances []*ATTRIBUTE_DEFINITION_BOOLEAN

	// parse all staged instances and check if they have a reference
	for attribute_definition_boolean := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		if ref, ok := stage.ATTRIBUTE_DEFINITION_BOOLEANs_reference[attribute_definition_boolean]; !ok {
			attribute_definition_booleans_newInstances = append(attribute_definition_booleans_newInstances, attribute_definition_boolean)
			newInstancesSlice = append(newInstancesSlice, attribute_definition_boolean.GongMarshallIdentifier(stage))
			if stage.ATTRIBUTE_DEFINITION_BOOLEANs_referenceOrder == nil {
				stage.ATTRIBUTE_DEFINITION_BOOLEANs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]uint)
			}
			stage.ATTRIBUTE_DEFINITION_BOOLEANs_referenceOrder[attribute_definition_boolean] = stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order[attribute_definition_boolean]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attribute_definition_boolean.GongMarshallUnstaging(stage))
			delete(stage.ATTRIBUTE_DEFINITION_BOOLEANs_referenceOrder, attribute_definition_boolean)
			fieldInitializers, pointersInitializations := attribute_definition_boolean.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order[ref] = stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order[attribute_definition_boolean]
			diffs := attribute_definition_boolean.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attribute_definition_boolean)
			delete(stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attribute_definition_boolean.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ATTRIBUTE_DEFINITION_BOOLEANs_reference {
		if _, ok := stage.ATTRIBUTE_DEFINITION_BOOLEANs[ref]; !ok {
			attribute_definition_booleans_deletedInstances = append(attribute_definition_booleans_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attribute_definition_booleans_newInstances)
	lenDeletedInstances += len(attribute_definition_booleans_deletedInstances)
	var attribute_definition_dates_newInstances []*ATTRIBUTE_DEFINITION_DATE
	var attribute_definition_dates_deletedInstances []*ATTRIBUTE_DEFINITION_DATE

	// parse all staged instances and check if they have a reference
	for attribute_definition_date := range stage.ATTRIBUTE_DEFINITION_DATEs {
		if ref, ok := stage.ATTRIBUTE_DEFINITION_DATEs_reference[attribute_definition_date]; !ok {
			attribute_definition_dates_newInstances = append(attribute_definition_dates_newInstances, attribute_definition_date)
			newInstancesSlice = append(newInstancesSlice, attribute_definition_date.GongMarshallIdentifier(stage))
			if stage.ATTRIBUTE_DEFINITION_DATEs_referenceOrder == nil {
				stage.ATTRIBUTE_DEFINITION_DATEs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_DATE]uint)
			}
			stage.ATTRIBUTE_DEFINITION_DATEs_referenceOrder[attribute_definition_date] = stage.ATTRIBUTE_DEFINITION_DATEMap_Staged_Order[attribute_definition_date]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attribute_definition_date.GongMarshallUnstaging(stage))
			delete(stage.ATTRIBUTE_DEFINITION_DATEs_referenceOrder, attribute_definition_date)
			fieldInitializers, pointersInitializations := attribute_definition_date.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ATTRIBUTE_DEFINITION_DATEMap_Staged_Order[ref] = stage.ATTRIBUTE_DEFINITION_DATEMap_Staged_Order[attribute_definition_date]
			diffs := attribute_definition_date.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attribute_definition_date)
			delete(stage.ATTRIBUTE_DEFINITION_DATEMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attribute_definition_date.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ATTRIBUTE_DEFINITION_DATEs_reference {
		if _, ok := stage.ATTRIBUTE_DEFINITION_DATEs[ref]; !ok {
			attribute_definition_dates_deletedInstances = append(attribute_definition_dates_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attribute_definition_dates_newInstances)
	lenDeletedInstances += len(attribute_definition_dates_deletedInstances)
	var attribute_definition_enumerations_newInstances []*ATTRIBUTE_DEFINITION_ENUMERATION
	var attribute_definition_enumerations_deletedInstances []*ATTRIBUTE_DEFINITION_ENUMERATION

	// parse all staged instances and check if they have a reference
	for attribute_definition_enumeration := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		if ref, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_reference[attribute_definition_enumeration]; !ok {
			attribute_definition_enumerations_newInstances = append(attribute_definition_enumerations_newInstances, attribute_definition_enumeration)
			newInstancesSlice = append(newInstancesSlice, attribute_definition_enumeration.GongMarshallIdentifier(stage))
			if stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_referenceOrder == nil {
				stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]uint)
			}
			stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_referenceOrder[attribute_definition_enumeration] = stage.ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order[attribute_definition_enumeration]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attribute_definition_enumeration.GongMarshallUnstaging(stage))
			delete(stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_referenceOrder, attribute_definition_enumeration)
			fieldInitializers, pointersInitializations := attribute_definition_enumeration.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order[ref] = stage.ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order[attribute_definition_enumeration]
			diffs := attribute_definition_enumeration.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attribute_definition_enumeration)
			delete(stage.ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attribute_definition_enumeration.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_reference {
		if _, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATIONs[ref]; !ok {
			attribute_definition_enumerations_deletedInstances = append(attribute_definition_enumerations_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attribute_definition_enumerations_newInstances)
	lenDeletedInstances += len(attribute_definition_enumerations_deletedInstances)
	var attribute_definition_integers_newInstances []*ATTRIBUTE_DEFINITION_INTEGER
	var attribute_definition_integers_deletedInstances []*ATTRIBUTE_DEFINITION_INTEGER

	// parse all staged instances and check if they have a reference
	for attribute_definition_integer := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		if ref, ok := stage.ATTRIBUTE_DEFINITION_INTEGERs_reference[attribute_definition_integer]; !ok {
			attribute_definition_integers_newInstances = append(attribute_definition_integers_newInstances, attribute_definition_integer)
			newInstancesSlice = append(newInstancesSlice, attribute_definition_integer.GongMarshallIdentifier(stage))
			if stage.ATTRIBUTE_DEFINITION_INTEGERs_referenceOrder == nil {
				stage.ATTRIBUTE_DEFINITION_INTEGERs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_INTEGER]uint)
			}
			stage.ATTRIBUTE_DEFINITION_INTEGERs_referenceOrder[attribute_definition_integer] = stage.ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order[attribute_definition_integer]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attribute_definition_integer.GongMarshallUnstaging(stage))
			delete(stage.ATTRIBUTE_DEFINITION_INTEGERs_referenceOrder, attribute_definition_integer)
			fieldInitializers, pointersInitializations := attribute_definition_integer.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order[ref] = stage.ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order[attribute_definition_integer]
			diffs := attribute_definition_integer.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attribute_definition_integer)
			delete(stage.ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attribute_definition_integer.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ATTRIBUTE_DEFINITION_INTEGERs_reference {
		if _, ok := stage.ATTRIBUTE_DEFINITION_INTEGERs[ref]; !ok {
			attribute_definition_integers_deletedInstances = append(attribute_definition_integers_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attribute_definition_integers_newInstances)
	lenDeletedInstances += len(attribute_definition_integers_deletedInstances)
	var attribute_definition_reals_newInstances []*ATTRIBUTE_DEFINITION_REAL
	var attribute_definition_reals_deletedInstances []*ATTRIBUTE_DEFINITION_REAL

	// parse all staged instances and check if they have a reference
	for attribute_definition_real := range stage.ATTRIBUTE_DEFINITION_REALs {
		if ref, ok := stage.ATTRIBUTE_DEFINITION_REALs_reference[attribute_definition_real]; !ok {
			attribute_definition_reals_newInstances = append(attribute_definition_reals_newInstances, attribute_definition_real)
			newInstancesSlice = append(newInstancesSlice, attribute_definition_real.GongMarshallIdentifier(stage))
			if stage.ATTRIBUTE_DEFINITION_REALs_referenceOrder == nil {
				stage.ATTRIBUTE_DEFINITION_REALs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_REAL]uint)
			}
			stage.ATTRIBUTE_DEFINITION_REALs_referenceOrder[attribute_definition_real] = stage.ATTRIBUTE_DEFINITION_REALMap_Staged_Order[attribute_definition_real]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attribute_definition_real.GongMarshallUnstaging(stage))
			delete(stage.ATTRIBUTE_DEFINITION_REALs_referenceOrder, attribute_definition_real)
			fieldInitializers, pointersInitializations := attribute_definition_real.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ATTRIBUTE_DEFINITION_REALMap_Staged_Order[ref] = stage.ATTRIBUTE_DEFINITION_REALMap_Staged_Order[attribute_definition_real]
			diffs := attribute_definition_real.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attribute_definition_real)
			delete(stage.ATTRIBUTE_DEFINITION_REALMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attribute_definition_real.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ATTRIBUTE_DEFINITION_REALs_reference {
		if _, ok := stage.ATTRIBUTE_DEFINITION_REALs[ref]; !ok {
			attribute_definition_reals_deletedInstances = append(attribute_definition_reals_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attribute_definition_reals_newInstances)
	lenDeletedInstances += len(attribute_definition_reals_deletedInstances)
	var attribute_definition_strings_newInstances []*ATTRIBUTE_DEFINITION_STRING
	var attribute_definition_strings_deletedInstances []*ATTRIBUTE_DEFINITION_STRING

	// parse all staged instances and check if they have a reference
	for attribute_definition_string := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		if ref, ok := stage.ATTRIBUTE_DEFINITION_STRINGs_reference[attribute_definition_string]; !ok {
			attribute_definition_strings_newInstances = append(attribute_definition_strings_newInstances, attribute_definition_string)
			newInstancesSlice = append(newInstancesSlice, attribute_definition_string.GongMarshallIdentifier(stage))
			if stage.ATTRIBUTE_DEFINITION_STRINGs_referenceOrder == nil {
				stage.ATTRIBUTE_DEFINITION_STRINGs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_STRING]uint)
			}
			stage.ATTRIBUTE_DEFINITION_STRINGs_referenceOrder[attribute_definition_string] = stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order[attribute_definition_string]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attribute_definition_string.GongMarshallUnstaging(stage))
			delete(stage.ATTRIBUTE_DEFINITION_STRINGs_referenceOrder, attribute_definition_string)
			fieldInitializers, pointersInitializations := attribute_definition_string.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order[ref] = stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order[attribute_definition_string]
			diffs := attribute_definition_string.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attribute_definition_string)
			delete(stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attribute_definition_string.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ATTRIBUTE_DEFINITION_STRINGs_reference {
		if _, ok := stage.ATTRIBUTE_DEFINITION_STRINGs[ref]; !ok {
			attribute_definition_strings_deletedInstances = append(attribute_definition_strings_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attribute_definition_strings_newInstances)
	lenDeletedInstances += len(attribute_definition_strings_deletedInstances)
	var attribute_definition_xhtmls_newInstances []*ATTRIBUTE_DEFINITION_XHTML
	var attribute_definition_xhtmls_deletedInstances []*ATTRIBUTE_DEFINITION_XHTML

	// parse all staged instances and check if they have a reference
	for attribute_definition_xhtml := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		if ref, ok := stage.ATTRIBUTE_DEFINITION_XHTMLs_reference[attribute_definition_xhtml]; !ok {
			attribute_definition_xhtmls_newInstances = append(attribute_definition_xhtmls_newInstances, attribute_definition_xhtml)
			newInstancesSlice = append(newInstancesSlice, attribute_definition_xhtml.GongMarshallIdentifier(stage))
			if stage.ATTRIBUTE_DEFINITION_XHTMLs_referenceOrder == nil {
				stage.ATTRIBUTE_DEFINITION_XHTMLs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_XHTML]uint)
			}
			stage.ATTRIBUTE_DEFINITION_XHTMLs_referenceOrder[attribute_definition_xhtml] = stage.ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order[attribute_definition_xhtml]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attribute_definition_xhtml.GongMarshallUnstaging(stage))
			delete(stage.ATTRIBUTE_DEFINITION_XHTMLs_referenceOrder, attribute_definition_xhtml)
			fieldInitializers, pointersInitializations := attribute_definition_xhtml.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order[ref] = stage.ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order[attribute_definition_xhtml]
			diffs := attribute_definition_xhtml.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attribute_definition_xhtml)
			delete(stage.ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attribute_definition_xhtml.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ATTRIBUTE_DEFINITION_XHTMLs_reference {
		if _, ok := stage.ATTRIBUTE_DEFINITION_XHTMLs[ref]; !ok {
			attribute_definition_xhtmls_deletedInstances = append(attribute_definition_xhtmls_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attribute_definition_xhtmls_newInstances)
	lenDeletedInstances += len(attribute_definition_xhtmls_deletedInstances)
	var attribute_value_booleans_newInstances []*ATTRIBUTE_VALUE_BOOLEAN
	var attribute_value_booleans_deletedInstances []*ATTRIBUTE_VALUE_BOOLEAN

	// parse all staged instances and check if they have a reference
	for attribute_value_boolean := range stage.ATTRIBUTE_VALUE_BOOLEANs {
		if ref, ok := stage.ATTRIBUTE_VALUE_BOOLEANs_reference[attribute_value_boolean]; !ok {
			attribute_value_booleans_newInstances = append(attribute_value_booleans_newInstances, attribute_value_boolean)
			newInstancesSlice = append(newInstancesSlice, attribute_value_boolean.GongMarshallIdentifier(stage))
			if stage.ATTRIBUTE_VALUE_BOOLEANs_referenceOrder == nil {
				stage.ATTRIBUTE_VALUE_BOOLEANs_referenceOrder = make(map[*ATTRIBUTE_VALUE_BOOLEAN]uint)
			}
			stage.ATTRIBUTE_VALUE_BOOLEANs_referenceOrder[attribute_value_boolean] = stage.ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[attribute_value_boolean]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attribute_value_boolean.GongMarshallUnstaging(stage))
			delete(stage.ATTRIBUTE_VALUE_BOOLEANs_referenceOrder, attribute_value_boolean)
			fieldInitializers, pointersInitializations := attribute_value_boolean.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[ref] = stage.ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[attribute_value_boolean]
			diffs := attribute_value_boolean.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attribute_value_boolean)
			delete(stage.ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attribute_value_boolean.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ATTRIBUTE_VALUE_BOOLEANs_reference {
		if _, ok := stage.ATTRIBUTE_VALUE_BOOLEANs[ref]; !ok {
			attribute_value_booleans_deletedInstances = append(attribute_value_booleans_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attribute_value_booleans_newInstances)
	lenDeletedInstances += len(attribute_value_booleans_deletedInstances)
	var attribute_value_dates_newInstances []*ATTRIBUTE_VALUE_DATE
	var attribute_value_dates_deletedInstances []*ATTRIBUTE_VALUE_DATE

	// parse all staged instances and check if they have a reference
	for attribute_value_date := range stage.ATTRIBUTE_VALUE_DATEs {
		if ref, ok := stage.ATTRIBUTE_VALUE_DATEs_reference[attribute_value_date]; !ok {
			attribute_value_dates_newInstances = append(attribute_value_dates_newInstances, attribute_value_date)
			newInstancesSlice = append(newInstancesSlice, attribute_value_date.GongMarshallIdentifier(stage))
			if stage.ATTRIBUTE_VALUE_DATEs_referenceOrder == nil {
				stage.ATTRIBUTE_VALUE_DATEs_referenceOrder = make(map[*ATTRIBUTE_VALUE_DATE]uint)
			}
			stage.ATTRIBUTE_VALUE_DATEs_referenceOrder[attribute_value_date] = stage.ATTRIBUTE_VALUE_DATEMap_Staged_Order[attribute_value_date]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attribute_value_date.GongMarshallUnstaging(stage))
			delete(stage.ATTRIBUTE_VALUE_DATEs_referenceOrder, attribute_value_date)
			fieldInitializers, pointersInitializations := attribute_value_date.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ATTRIBUTE_VALUE_DATEMap_Staged_Order[ref] = stage.ATTRIBUTE_VALUE_DATEMap_Staged_Order[attribute_value_date]
			diffs := attribute_value_date.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attribute_value_date)
			delete(stage.ATTRIBUTE_VALUE_DATEMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attribute_value_date.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ATTRIBUTE_VALUE_DATEs_reference {
		if _, ok := stage.ATTRIBUTE_VALUE_DATEs[ref]; !ok {
			attribute_value_dates_deletedInstances = append(attribute_value_dates_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attribute_value_dates_newInstances)
	lenDeletedInstances += len(attribute_value_dates_deletedInstances)
	var attribute_value_enumerations_newInstances []*ATTRIBUTE_VALUE_ENUMERATION
	var attribute_value_enumerations_deletedInstances []*ATTRIBUTE_VALUE_ENUMERATION

	// parse all staged instances and check if they have a reference
	for attribute_value_enumeration := range stage.ATTRIBUTE_VALUE_ENUMERATIONs {
		if ref, ok := stage.ATTRIBUTE_VALUE_ENUMERATIONs_reference[attribute_value_enumeration]; !ok {
			attribute_value_enumerations_newInstances = append(attribute_value_enumerations_newInstances, attribute_value_enumeration)
			newInstancesSlice = append(newInstancesSlice, attribute_value_enumeration.GongMarshallIdentifier(stage))
			if stage.ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder == nil {
				stage.ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder = make(map[*ATTRIBUTE_VALUE_ENUMERATION]uint)
			}
			stage.ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder[attribute_value_enumeration] = stage.ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[attribute_value_enumeration]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attribute_value_enumeration.GongMarshallUnstaging(stage))
			delete(stage.ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder, attribute_value_enumeration)
			fieldInitializers, pointersInitializations := attribute_value_enumeration.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[ref] = stage.ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[attribute_value_enumeration]
			diffs := attribute_value_enumeration.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attribute_value_enumeration)
			delete(stage.ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attribute_value_enumeration.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ATTRIBUTE_VALUE_ENUMERATIONs_reference {
		if _, ok := stage.ATTRIBUTE_VALUE_ENUMERATIONs[ref]; !ok {
			attribute_value_enumerations_deletedInstances = append(attribute_value_enumerations_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attribute_value_enumerations_newInstances)
	lenDeletedInstances += len(attribute_value_enumerations_deletedInstances)
	var attribute_value_integers_newInstances []*ATTRIBUTE_VALUE_INTEGER
	var attribute_value_integers_deletedInstances []*ATTRIBUTE_VALUE_INTEGER

	// parse all staged instances and check if they have a reference
	for attribute_value_integer := range stage.ATTRIBUTE_VALUE_INTEGERs {
		if ref, ok := stage.ATTRIBUTE_VALUE_INTEGERs_reference[attribute_value_integer]; !ok {
			attribute_value_integers_newInstances = append(attribute_value_integers_newInstances, attribute_value_integer)
			newInstancesSlice = append(newInstancesSlice, attribute_value_integer.GongMarshallIdentifier(stage))
			if stage.ATTRIBUTE_VALUE_INTEGERs_referenceOrder == nil {
				stage.ATTRIBUTE_VALUE_INTEGERs_referenceOrder = make(map[*ATTRIBUTE_VALUE_INTEGER]uint)
			}
			stage.ATTRIBUTE_VALUE_INTEGERs_referenceOrder[attribute_value_integer] = stage.ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[attribute_value_integer]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attribute_value_integer.GongMarshallUnstaging(stage))
			delete(stage.ATTRIBUTE_VALUE_INTEGERs_referenceOrder, attribute_value_integer)
			fieldInitializers, pointersInitializations := attribute_value_integer.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[ref] = stage.ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[attribute_value_integer]
			diffs := attribute_value_integer.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attribute_value_integer)
			delete(stage.ATTRIBUTE_VALUE_INTEGERMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attribute_value_integer.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ATTRIBUTE_VALUE_INTEGERs_reference {
		if _, ok := stage.ATTRIBUTE_VALUE_INTEGERs[ref]; !ok {
			attribute_value_integers_deletedInstances = append(attribute_value_integers_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attribute_value_integers_newInstances)
	lenDeletedInstances += len(attribute_value_integers_deletedInstances)
	var attribute_value_reals_newInstances []*ATTRIBUTE_VALUE_REAL
	var attribute_value_reals_deletedInstances []*ATTRIBUTE_VALUE_REAL

	// parse all staged instances and check if they have a reference
	for attribute_value_real := range stage.ATTRIBUTE_VALUE_REALs {
		if ref, ok := stage.ATTRIBUTE_VALUE_REALs_reference[attribute_value_real]; !ok {
			attribute_value_reals_newInstances = append(attribute_value_reals_newInstances, attribute_value_real)
			newInstancesSlice = append(newInstancesSlice, attribute_value_real.GongMarshallIdentifier(stage))
			if stage.ATTRIBUTE_VALUE_REALs_referenceOrder == nil {
				stage.ATTRIBUTE_VALUE_REALs_referenceOrder = make(map[*ATTRIBUTE_VALUE_REAL]uint)
			}
			stage.ATTRIBUTE_VALUE_REALs_referenceOrder[attribute_value_real] = stage.ATTRIBUTE_VALUE_REALMap_Staged_Order[attribute_value_real]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attribute_value_real.GongMarshallUnstaging(stage))
			delete(stage.ATTRIBUTE_VALUE_REALs_referenceOrder, attribute_value_real)
			fieldInitializers, pointersInitializations := attribute_value_real.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ATTRIBUTE_VALUE_REALMap_Staged_Order[ref] = stage.ATTRIBUTE_VALUE_REALMap_Staged_Order[attribute_value_real]
			diffs := attribute_value_real.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attribute_value_real)
			delete(stage.ATTRIBUTE_VALUE_REALMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attribute_value_real.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ATTRIBUTE_VALUE_REALs_reference {
		if _, ok := stage.ATTRIBUTE_VALUE_REALs[ref]; !ok {
			attribute_value_reals_deletedInstances = append(attribute_value_reals_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attribute_value_reals_newInstances)
	lenDeletedInstances += len(attribute_value_reals_deletedInstances)
	var attribute_value_strings_newInstances []*ATTRIBUTE_VALUE_STRING
	var attribute_value_strings_deletedInstances []*ATTRIBUTE_VALUE_STRING

	// parse all staged instances and check if they have a reference
	for attribute_value_string := range stage.ATTRIBUTE_VALUE_STRINGs {
		if ref, ok := stage.ATTRIBUTE_VALUE_STRINGs_reference[attribute_value_string]; !ok {
			attribute_value_strings_newInstances = append(attribute_value_strings_newInstances, attribute_value_string)
			newInstancesSlice = append(newInstancesSlice, attribute_value_string.GongMarshallIdentifier(stage))
			if stage.ATTRIBUTE_VALUE_STRINGs_referenceOrder == nil {
				stage.ATTRIBUTE_VALUE_STRINGs_referenceOrder = make(map[*ATTRIBUTE_VALUE_STRING]uint)
			}
			stage.ATTRIBUTE_VALUE_STRINGs_referenceOrder[attribute_value_string] = stage.ATTRIBUTE_VALUE_STRINGMap_Staged_Order[attribute_value_string]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attribute_value_string.GongMarshallUnstaging(stage))
			delete(stage.ATTRIBUTE_VALUE_STRINGs_referenceOrder, attribute_value_string)
			fieldInitializers, pointersInitializations := attribute_value_string.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ATTRIBUTE_VALUE_STRINGMap_Staged_Order[ref] = stage.ATTRIBUTE_VALUE_STRINGMap_Staged_Order[attribute_value_string]
			diffs := attribute_value_string.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attribute_value_string)
			delete(stage.ATTRIBUTE_VALUE_STRINGMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attribute_value_string.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ATTRIBUTE_VALUE_STRINGs_reference {
		if _, ok := stage.ATTRIBUTE_VALUE_STRINGs[ref]; !ok {
			attribute_value_strings_deletedInstances = append(attribute_value_strings_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attribute_value_strings_newInstances)
	lenDeletedInstances += len(attribute_value_strings_deletedInstances)
	var attribute_value_xhtmls_newInstances []*ATTRIBUTE_VALUE_XHTML
	var attribute_value_xhtmls_deletedInstances []*ATTRIBUTE_VALUE_XHTML

	// parse all staged instances and check if they have a reference
	for attribute_value_xhtml := range stage.ATTRIBUTE_VALUE_XHTMLs {
		if ref, ok := stage.ATTRIBUTE_VALUE_XHTMLs_reference[attribute_value_xhtml]; !ok {
			attribute_value_xhtmls_newInstances = append(attribute_value_xhtmls_newInstances, attribute_value_xhtml)
			newInstancesSlice = append(newInstancesSlice, attribute_value_xhtml.GongMarshallIdentifier(stage))
			if stage.ATTRIBUTE_VALUE_XHTMLs_referenceOrder == nil {
				stage.ATTRIBUTE_VALUE_XHTMLs_referenceOrder = make(map[*ATTRIBUTE_VALUE_XHTML]uint)
			}
			stage.ATTRIBUTE_VALUE_XHTMLs_referenceOrder[attribute_value_xhtml] = stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[attribute_value_xhtml]
			newInstancesReverseSlice = append(newInstancesReverseSlice, attribute_value_xhtml.GongMarshallUnstaging(stage))
			delete(stage.ATTRIBUTE_VALUE_XHTMLs_referenceOrder, attribute_value_xhtml)
			fieldInitializers, pointersInitializations := attribute_value_xhtml.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[ref] = stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[attribute_value_xhtml]
			diffs := attribute_value_xhtml.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, attribute_value_xhtml)
			delete(stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", attribute_value_xhtml.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ATTRIBUTE_VALUE_XHTMLs_reference {
		if _, ok := stage.ATTRIBUTE_VALUE_XHTMLs[ref]; !ok {
			attribute_value_xhtmls_deletedInstances = append(attribute_value_xhtmls_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(attribute_value_xhtmls_newInstances)
	lenDeletedInstances += len(attribute_value_xhtmls_deletedInstances)
	var a_alternative_ids_newInstances []*A_ALTERNATIVE_ID
	var a_alternative_ids_deletedInstances []*A_ALTERNATIVE_ID

	// parse all staged instances and check if they have a reference
	for a_alternative_id := range stage.A_ALTERNATIVE_IDs {
		if ref, ok := stage.A_ALTERNATIVE_IDs_reference[a_alternative_id]; !ok {
			a_alternative_ids_newInstances = append(a_alternative_ids_newInstances, a_alternative_id)
			newInstancesSlice = append(newInstancesSlice, a_alternative_id.GongMarshallIdentifier(stage))
			if stage.A_ALTERNATIVE_IDs_referenceOrder == nil {
				stage.A_ALTERNATIVE_IDs_referenceOrder = make(map[*A_ALTERNATIVE_ID]uint)
			}
			stage.A_ALTERNATIVE_IDs_referenceOrder[a_alternative_id] = stage.A_ALTERNATIVE_IDMap_Staged_Order[a_alternative_id]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_alternative_id.GongMarshallUnstaging(stage))
			delete(stage.A_ALTERNATIVE_IDs_referenceOrder, a_alternative_id)
			fieldInitializers, pointersInitializations := a_alternative_id.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ALTERNATIVE_IDMap_Staged_Order[ref] = stage.A_ALTERNATIVE_IDMap_Staged_Order[a_alternative_id]
			diffs := a_alternative_id.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_alternative_id)
			delete(stage.A_ALTERNATIVE_IDMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_alternative_id.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ALTERNATIVE_IDs_reference {
		if _, ok := stage.A_ALTERNATIVE_IDs[ref]; !ok {
			a_alternative_ids_deletedInstances = append(a_alternative_ids_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_alternative_ids_newInstances)
	lenDeletedInstances += len(a_alternative_ids_deletedInstances)
	var a_attribute_definition_boolean_refs_newInstances []*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF
	var a_attribute_definition_boolean_refs_deletedInstances []*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF

	// parse all staged instances and check if they have a reference
	for a_attribute_definition_boolean_ref := range stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs {
		if ref, ok := stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_reference[a_attribute_definition_boolean_ref]; !ok {
			a_attribute_definition_boolean_refs_newInstances = append(a_attribute_definition_boolean_refs_newInstances, a_attribute_definition_boolean_ref)
			newInstancesSlice = append(newInstancesSlice, a_attribute_definition_boolean_ref.GongMarshallIdentifier(stage))
			if stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_referenceOrder == nil {
				stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]uint)
			}
			stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_referenceOrder[a_attribute_definition_boolean_ref] = stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_attribute_definition_boolean_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_attribute_definition_boolean_ref.GongMarshallUnstaging(stage))
			delete(stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_referenceOrder, a_attribute_definition_boolean_ref)
			fieldInitializers, pointersInitializations := a_attribute_definition_boolean_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFMap_Staged_Order[ref] = stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_attribute_definition_boolean_ref]
			diffs := a_attribute_definition_boolean_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_attribute_definition_boolean_ref)
			delete(stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_attribute_definition_boolean_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_reference {
		if _, ok := stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs[ref]; !ok {
			a_attribute_definition_boolean_refs_deletedInstances = append(a_attribute_definition_boolean_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_attribute_definition_boolean_refs_newInstances)
	lenDeletedInstances += len(a_attribute_definition_boolean_refs_deletedInstances)
	var a_attribute_definition_date_refs_newInstances []*A_ATTRIBUTE_DEFINITION_DATE_REF
	var a_attribute_definition_date_refs_deletedInstances []*A_ATTRIBUTE_DEFINITION_DATE_REF

	// parse all staged instances and check if they have a reference
	for a_attribute_definition_date_ref := range stage.A_ATTRIBUTE_DEFINITION_DATE_REFs {
		if ref, ok := stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_reference[a_attribute_definition_date_ref]; !ok {
			a_attribute_definition_date_refs_newInstances = append(a_attribute_definition_date_refs_newInstances, a_attribute_definition_date_ref)
			newInstancesSlice = append(newInstancesSlice, a_attribute_definition_date_ref.GongMarshallIdentifier(stage))
			if stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_referenceOrder == nil {
				stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_DATE_REF]uint)
			}
			stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_referenceOrder[a_attribute_definition_date_ref] = stage.A_ATTRIBUTE_DEFINITION_DATE_REFMap_Staged_Order[a_attribute_definition_date_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_attribute_definition_date_ref.GongMarshallUnstaging(stage))
			delete(stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_referenceOrder, a_attribute_definition_date_ref)
			fieldInitializers, pointersInitializations := a_attribute_definition_date_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ATTRIBUTE_DEFINITION_DATE_REFMap_Staged_Order[ref] = stage.A_ATTRIBUTE_DEFINITION_DATE_REFMap_Staged_Order[a_attribute_definition_date_ref]
			diffs := a_attribute_definition_date_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_attribute_definition_date_ref)
			delete(stage.A_ATTRIBUTE_DEFINITION_DATE_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_attribute_definition_date_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_reference {
		if _, ok := stage.A_ATTRIBUTE_DEFINITION_DATE_REFs[ref]; !ok {
			a_attribute_definition_date_refs_deletedInstances = append(a_attribute_definition_date_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_attribute_definition_date_refs_newInstances)
	lenDeletedInstances += len(a_attribute_definition_date_refs_deletedInstances)
	var a_attribute_definition_enumeration_refs_newInstances []*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF
	var a_attribute_definition_enumeration_refs_deletedInstances []*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF

	// parse all staged instances and check if they have a reference
	for a_attribute_definition_enumeration_ref := range stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs {
		if ref, ok := stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_reference[a_attribute_definition_enumeration_ref]; !ok {
			a_attribute_definition_enumeration_refs_newInstances = append(a_attribute_definition_enumeration_refs_newInstances, a_attribute_definition_enumeration_ref)
			newInstancesSlice = append(newInstancesSlice, a_attribute_definition_enumeration_ref.GongMarshallIdentifier(stage))
			if stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_referenceOrder == nil {
				stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]uint)
			}
			stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_referenceOrder[a_attribute_definition_enumeration_ref] = stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_attribute_definition_enumeration_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_attribute_definition_enumeration_ref.GongMarshallUnstaging(stage))
			delete(stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_referenceOrder, a_attribute_definition_enumeration_ref)
			fieldInitializers, pointersInitializations := a_attribute_definition_enumeration_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFMap_Staged_Order[ref] = stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_attribute_definition_enumeration_ref]
			diffs := a_attribute_definition_enumeration_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_attribute_definition_enumeration_ref)
			delete(stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_attribute_definition_enumeration_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_reference {
		if _, ok := stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs[ref]; !ok {
			a_attribute_definition_enumeration_refs_deletedInstances = append(a_attribute_definition_enumeration_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_attribute_definition_enumeration_refs_newInstances)
	lenDeletedInstances += len(a_attribute_definition_enumeration_refs_deletedInstances)
	var a_attribute_definition_integer_refs_newInstances []*A_ATTRIBUTE_DEFINITION_INTEGER_REF
	var a_attribute_definition_integer_refs_deletedInstances []*A_ATTRIBUTE_DEFINITION_INTEGER_REF

	// parse all staged instances and check if they have a reference
	for a_attribute_definition_integer_ref := range stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs {
		if ref, ok := stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_reference[a_attribute_definition_integer_ref]; !ok {
			a_attribute_definition_integer_refs_newInstances = append(a_attribute_definition_integer_refs_newInstances, a_attribute_definition_integer_ref)
			newInstancesSlice = append(newInstancesSlice, a_attribute_definition_integer_ref.GongMarshallIdentifier(stage))
			if stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_referenceOrder == nil {
				stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_INTEGER_REF]uint)
			}
			stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_referenceOrder[a_attribute_definition_integer_ref] = stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFMap_Staged_Order[a_attribute_definition_integer_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_attribute_definition_integer_ref.GongMarshallUnstaging(stage))
			delete(stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_referenceOrder, a_attribute_definition_integer_ref)
			fieldInitializers, pointersInitializations := a_attribute_definition_integer_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFMap_Staged_Order[ref] = stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFMap_Staged_Order[a_attribute_definition_integer_ref]
			diffs := a_attribute_definition_integer_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_attribute_definition_integer_ref)
			delete(stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_attribute_definition_integer_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_reference {
		if _, ok := stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs[ref]; !ok {
			a_attribute_definition_integer_refs_deletedInstances = append(a_attribute_definition_integer_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_attribute_definition_integer_refs_newInstances)
	lenDeletedInstances += len(a_attribute_definition_integer_refs_deletedInstances)
	var a_attribute_definition_real_refs_newInstances []*A_ATTRIBUTE_DEFINITION_REAL_REF
	var a_attribute_definition_real_refs_deletedInstances []*A_ATTRIBUTE_DEFINITION_REAL_REF

	// parse all staged instances and check if they have a reference
	for a_attribute_definition_real_ref := range stage.A_ATTRIBUTE_DEFINITION_REAL_REFs {
		if ref, ok := stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_reference[a_attribute_definition_real_ref]; !ok {
			a_attribute_definition_real_refs_newInstances = append(a_attribute_definition_real_refs_newInstances, a_attribute_definition_real_ref)
			newInstancesSlice = append(newInstancesSlice, a_attribute_definition_real_ref.GongMarshallIdentifier(stage))
			if stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_referenceOrder == nil {
				stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_REAL_REF]uint)
			}
			stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_referenceOrder[a_attribute_definition_real_ref] = stage.A_ATTRIBUTE_DEFINITION_REAL_REFMap_Staged_Order[a_attribute_definition_real_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_attribute_definition_real_ref.GongMarshallUnstaging(stage))
			delete(stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_referenceOrder, a_attribute_definition_real_ref)
			fieldInitializers, pointersInitializations := a_attribute_definition_real_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ATTRIBUTE_DEFINITION_REAL_REFMap_Staged_Order[ref] = stage.A_ATTRIBUTE_DEFINITION_REAL_REFMap_Staged_Order[a_attribute_definition_real_ref]
			diffs := a_attribute_definition_real_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_attribute_definition_real_ref)
			delete(stage.A_ATTRIBUTE_DEFINITION_REAL_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_attribute_definition_real_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_reference {
		if _, ok := stage.A_ATTRIBUTE_DEFINITION_REAL_REFs[ref]; !ok {
			a_attribute_definition_real_refs_deletedInstances = append(a_attribute_definition_real_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_attribute_definition_real_refs_newInstances)
	lenDeletedInstances += len(a_attribute_definition_real_refs_deletedInstances)
	var a_attribute_definition_string_refs_newInstances []*A_ATTRIBUTE_DEFINITION_STRING_REF
	var a_attribute_definition_string_refs_deletedInstances []*A_ATTRIBUTE_DEFINITION_STRING_REF

	// parse all staged instances and check if they have a reference
	for a_attribute_definition_string_ref := range stage.A_ATTRIBUTE_DEFINITION_STRING_REFs {
		if ref, ok := stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_reference[a_attribute_definition_string_ref]; !ok {
			a_attribute_definition_string_refs_newInstances = append(a_attribute_definition_string_refs_newInstances, a_attribute_definition_string_ref)
			newInstancesSlice = append(newInstancesSlice, a_attribute_definition_string_ref.GongMarshallIdentifier(stage))
			if stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_referenceOrder == nil {
				stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_STRING_REF]uint)
			}
			stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_referenceOrder[a_attribute_definition_string_ref] = stage.A_ATTRIBUTE_DEFINITION_STRING_REFMap_Staged_Order[a_attribute_definition_string_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_attribute_definition_string_ref.GongMarshallUnstaging(stage))
			delete(stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_referenceOrder, a_attribute_definition_string_ref)
			fieldInitializers, pointersInitializations := a_attribute_definition_string_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ATTRIBUTE_DEFINITION_STRING_REFMap_Staged_Order[ref] = stage.A_ATTRIBUTE_DEFINITION_STRING_REFMap_Staged_Order[a_attribute_definition_string_ref]
			diffs := a_attribute_definition_string_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_attribute_definition_string_ref)
			delete(stage.A_ATTRIBUTE_DEFINITION_STRING_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_attribute_definition_string_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_reference {
		if _, ok := stage.A_ATTRIBUTE_DEFINITION_STRING_REFs[ref]; !ok {
			a_attribute_definition_string_refs_deletedInstances = append(a_attribute_definition_string_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_attribute_definition_string_refs_newInstances)
	lenDeletedInstances += len(a_attribute_definition_string_refs_deletedInstances)
	var a_attribute_definition_xhtml_refs_newInstances []*A_ATTRIBUTE_DEFINITION_XHTML_REF
	var a_attribute_definition_xhtml_refs_deletedInstances []*A_ATTRIBUTE_DEFINITION_XHTML_REF

	// parse all staged instances and check if they have a reference
	for a_attribute_definition_xhtml_ref := range stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs {
		if ref, ok := stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_reference[a_attribute_definition_xhtml_ref]; !ok {
			a_attribute_definition_xhtml_refs_newInstances = append(a_attribute_definition_xhtml_refs_newInstances, a_attribute_definition_xhtml_ref)
			newInstancesSlice = append(newInstancesSlice, a_attribute_definition_xhtml_ref.GongMarshallIdentifier(stage))
			if stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_referenceOrder == nil {
				stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_XHTML_REF]uint)
			}
			stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_referenceOrder[a_attribute_definition_xhtml_ref] = stage.A_ATTRIBUTE_DEFINITION_XHTML_REFMap_Staged_Order[a_attribute_definition_xhtml_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_attribute_definition_xhtml_ref.GongMarshallUnstaging(stage))
			delete(stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_referenceOrder, a_attribute_definition_xhtml_ref)
			fieldInitializers, pointersInitializations := a_attribute_definition_xhtml_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ATTRIBUTE_DEFINITION_XHTML_REFMap_Staged_Order[ref] = stage.A_ATTRIBUTE_DEFINITION_XHTML_REFMap_Staged_Order[a_attribute_definition_xhtml_ref]
			diffs := a_attribute_definition_xhtml_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_attribute_definition_xhtml_ref)
			delete(stage.A_ATTRIBUTE_DEFINITION_XHTML_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_attribute_definition_xhtml_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_reference {
		if _, ok := stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs[ref]; !ok {
			a_attribute_definition_xhtml_refs_deletedInstances = append(a_attribute_definition_xhtml_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_attribute_definition_xhtml_refs_newInstances)
	lenDeletedInstances += len(a_attribute_definition_xhtml_refs_deletedInstances)
	var a_attribute_value_booleans_newInstances []*A_ATTRIBUTE_VALUE_BOOLEAN
	var a_attribute_value_booleans_deletedInstances []*A_ATTRIBUTE_VALUE_BOOLEAN

	// parse all staged instances and check if they have a reference
	for a_attribute_value_boolean := range stage.A_ATTRIBUTE_VALUE_BOOLEANs {
		if ref, ok := stage.A_ATTRIBUTE_VALUE_BOOLEANs_reference[a_attribute_value_boolean]; !ok {
			a_attribute_value_booleans_newInstances = append(a_attribute_value_booleans_newInstances, a_attribute_value_boolean)
			newInstancesSlice = append(newInstancesSlice, a_attribute_value_boolean.GongMarshallIdentifier(stage))
			if stage.A_ATTRIBUTE_VALUE_BOOLEANs_referenceOrder == nil {
				stage.A_ATTRIBUTE_VALUE_BOOLEANs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_BOOLEAN]uint)
			}
			stage.A_ATTRIBUTE_VALUE_BOOLEANs_referenceOrder[a_attribute_value_boolean] = stage.A_ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[a_attribute_value_boolean]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_attribute_value_boolean.GongMarshallUnstaging(stage))
			delete(stage.A_ATTRIBUTE_VALUE_BOOLEANs_referenceOrder, a_attribute_value_boolean)
			fieldInitializers, pointersInitializations := a_attribute_value_boolean.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[ref] = stage.A_ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[a_attribute_value_boolean]
			diffs := a_attribute_value_boolean.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_attribute_value_boolean)
			delete(stage.A_ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_attribute_value_boolean.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ATTRIBUTE_VALUE_BOOLEANs_reference {
		if _, ok := stage.A_ATTRIBUTE_VALUE_BOOLEANs[ref]; !ok {
			a_attribute_value_booleans_deletedInstances = append(a_attribute_value_booleans_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_attribute_value_booleans_newInstances)
	lenDeletedInstances += len(a_attribute_value_booleans_deletedInstances)
	var a_attribute_value_dates_newInstances []*A_ATTRIBUTE_VALUE_DATE
	var a_attribute_value_dates_deletedInstances []*A_ATTRIBUTE_VALUE_DATE

	// parse all staged instances and check if they have a reference
	for a_attribute_value_date := range stage.A_ATTRIBUTE_VALUE_DATEs {
		if ref, ok := stage.A_ATTRIBUTE_VALUE_DATEs_reference[a_attribute_value_date]; !ok {
			a_attribute_value_dates_newInstances = append(a_attribute_value_dates_newInstances, a_attribute_value_date)
			newInstancesSlice = append(newInstancesSlice, a_attribute_value_date.GongMarshallIdentifier(stage))
			if stage.A_ATTRIBUTE_VALUE_DATEs_referenceOrder == nil {
				stage.A_ATTRIBUTE_VALUE_DATEs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_DATE]uint)
			}
			stage.A_ATTRIBUTE_VALUE_DATEs_referenceOrder[a_attribute_value_date] = stage.A_ATTRIBUTE_VALUE_DATEMap_Staged_Order[a_attribute_value_date]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_attribute_value_date.GongMarshallUnstaging(stage))
			delete(stage.A_ATTRIBUTE_VALUE_DATEs_referenceOrder, a_attribute_value_date)
			fieldInitializers, pointersInitializations := a_attribute_value_date.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ATTRIBUTE_VALUE_DATEMap_Staged_Order[ref] = stage.A_ATTRIBUTE_VALUE_DATEMap_Staged_Order[a_attribute_value_date]
			diffs := a_attribute_value_date.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_attribute_value_date)
			delete(stage.A_ATTRIBUTE_VALUE_DATEMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_attribute_value_date.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ATTRIBUTE_VALUE_DATEs_reference {
		if _, ok := stage.A_ATTRIBUTE_VALUE_DATEs[ref]; !ok {
			a_attribute_value_dates_deletedInstances = append(a_attribute_value_dates_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_attribute_value_dates_newInstances)
	lenDeletedInstances += len(a_attribute_value_dates_deletedInstances)
	var a_attribute_value_enumerations_newInstances []*A_ATTRIBUTE_VALUE_ENUMERATION
	var a_attribute_value_enumerations_deletedInstances []*A_ATTRIBUTE_VALUE_ENUMERATION

	// parse all staged instances and check if they have a reference
	for a_attribute_value_enumeration := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
		if ref, ok := stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_reference[a_attribute_value_enumeration]; !ok {
			a_attribute_value_enumerations_newInstances = append(a_attribute_value_enumerations_newInstances, a_attribute_value_enumeration)
			newInstancesSlice = append(newInstancesSlice, a_attribute_value_enumeration.GongMarshallIdentifier(stage))
			if stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder == nil {
				stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_ENUMERATION]uint)
			}
			stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder[a_attribute_value_enumeration] = stage.A_ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[a_attribute_value_enumeration]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_attribute_value_enumeration.GongMarshallUnstaging(stage))
			delete(stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder, a_attribute_value_enumeration)
			fieldInitializers, pointersInitializations := a_attribute_value_enumeration.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[ref] = stage.A_ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[a_attribute_value_enumeration]
			diffs := a_attribute_value_enumeration.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_attribute_value_enumeration)
			delete(stage.A_ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_attribute_value_enumeration.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_reference {
		if _, ok := stage.A_ATTRIBUTE_VALUE_ENUMERATIONs[ref]; !ok {
			a_attribute_value_enumerations_deletedInstances = append(a_attribute_value_enumerations_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_attribute_value_enumerations_newInstances)
	lenDeletedInstances += len(a_attribute_value_enumerations_deletedInstances)
	var a_attribute_value_integers_newInstances []*A_ATTRIBUTE_VALUE_INTEGER
	var a_attribute_value_integers_deletedInstances []*A_ATTRIBUTE_VALUE_INTEGER

	// parse all staged instances and check if they have a reference
	for a_attribute_value_integer := range stage.A_ATTRIBUTE_VALUE_INTEGERs {
		if ref, ok := stage.A_ATTRIBUTE_VALUE_INTEGERs_reference[a_attribute_value_integer]; !ok {
			a_attribute_value_integers_newInstances = append(a_attribute_value_integers_newInstances, a_attribute_value_integer)
			newInstancesSlice = append(newInstancesSlice, a_attribute_value_integer.GongMarshallIdentifier(stage))
			if stage.A_ATTRIBUTE_VALUE_INTEGERs_referenceOrder == nil {
				stage.A_ATTRIBUTE_VALUE_INTEGERs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_INTEGER]uint)
			}
			stage.A_ATTRIBUTE_VALUE_INTEGERs_referenceOrder[a_attribute_value_integer] = stage.A_ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[a_attribute_value_integer]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_attribute_value_integer.GongMarshallUnstaging(stage))
			delete(stage.A_ATTRIBUTE_VALUE_INTEGERs_referenceOrder, a_attribute_value_integer)
			fieldInitializers, pointersInitializations := a_attribute_value_integer.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[ref] = stage.A_ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[a_attribute_value_integer]
			diffs := a_attribute_value_integer.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_attribute_value_integer)
			delete(stage.A_ATTRIBUTE_VALUE_INTEGERMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_attribute_value_integer.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ATTRIBUTE_VALUE_INTEGERs_reference {
		if _, ok := stage.A_ATTRIBUTE_VALUE_INTEGERs[ref]; !ok {
			a_attribute_value_integers_deletedInstances = append(a_attribute_value_integers_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_attribute_value_integers_newInstances)
	lenDeletedInstances += len(a_attribute_value_integers_deletedInstances)
	var a_attribute_value_reals_newInstances []*A_ATTRIBUTE_VALUE_REAL
	var a_attribute_value_reals_deletedInstances []*A_ATTRIBUTE_VALUE_REAL

	// parse all staged instances and check if they have a reference
	for a_attribute_value_real := range stage.A_ATTRIBUTE_VALUE_REALs {
		if ref, ok := stage.A_ATTRIBUTE_VALUE_REALs_reference[a_attribute_value_real]; !ok {
			a_attribute_value_reals_newInstances = append(a_attribute_value_reals_newInstances, a_attribute_value_real)
			newInstancesSlice = append(newInstancesSlice, a_attribute_value_real.GongMarshallIdentifier(stage))
			if stage.A_ATTRIBUTE_VALUE_REALs_referenceOrder == nil {
				stage.A_ATTRIBUTE_VALUE_REALs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_REAL]uint)
			}
			stage.A_ATTRIBUTE_VALUE_REALs_referenceOrder[a_attribute_value_real] = stage.A_ATTRIBUTE_VALUE_REALMap_Staged_Order[a_attribute_value_real]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_attribute_value_real.GongMarshallUnstaging(stage))
			delete(stage.A_ATTRIBUTE_VALUE_REALs_referenceOrder, a_attribute_value_real)
			fieldInitializers, pointersInitializations := a_attribute_value_real.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ATTRIBUTE_VALUE_REALMap_Staged_Order[ref] = stage.A_ATTRIBUTE_VALUE_REALMap_Staged_Order[a_attribute_value_real]
			diffs := a_attribute_value_real.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_attribute_value_real)
			delete(stage.A_ATTRIBUTE_VALUE_REALMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_attribute_value_real.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ATTRIBUTE_VALUE_REALs_reference {
		if _, ok := stage.A_ATTRIBUTE_VALUE_REALs[ref]; !ok {
			a_attribute_value_reals_deletedInstances = append(a_attribute_value_reals_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_attribute_value_reals_newInstances)
	lenDeletedInstances += len(a_attribute_value_reals_deletedInstances)
	var a_attribute_value_strings_newInstances []*A_ATTRIBUTE_VALUE_STRING
	var a_attribute_value_strings_deletedInstances []*A_ATTRIBUTE_VALUE_STRING

	// parse all staged instances and check if they have a reference
	for a_attribute_value_string := range stage.A_ATTRIBUTE_VALUE_STRINGs {
		if ref, ok := stage.A_ATTRIBUTE_VALUE_STRINGs_reference[a_attribute_value_string]; !ok {
			a_attribute_value_strings_newInstances = append(a_attribute_value_strings_newInstances, a_attribute_value_string)
			newInstancesSlice = append(newInstancesSlice, a_attribute_value_string.GongMarshallIdentifier(stage))
			if stage.A_ATTRIBUTE_VALUE_STRINGs_referenceOrder == nil {
				stage.A_ATTRIBUTE_VALUE_STRINGs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_STRING]uint)
			}
			stage.A_ATTRIBUTE_VALUE_STRINGs_referenceOrder[a_attribute_value_string] = stage.A_ATTRIBUTE_VALUE_STRINGMap_Staged_Order[a_attribute_value_string]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_attribute_value_string.GongMarshallUnstaging(stage))
			delete(stage.A_ATTRIBUTE_VALUE_STRINGs_referenceOrder, a_attribute_value_string)
			fieldInitializers, pointersInitializations := a_attribute_value_string.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ATTRIBUTE_VALUE_STRINGMap_Staged_Order[ref] = stage.A_ATTRIBUTE_VALUE_STRINGMap_Staged_Order[a_attribute_value_string]
			diffs := a_attribute_value_string.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_attribute_value_string)
			delete(stage.A_ATTRIBUTE_VALUE_STRINGMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_attribute_value_string.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ATTRIBUTE_VALUE_STRINGs_reference {
		if _, ok := stage.A_ATTRIBUTE_VALUE_STRINGs[ref]; !ok {
			a_attribute_value_strings_deletedInstances = append(a_attribute_value_strings_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_attribute_value_strings_newInstances)
	lenDeletedInstances += len(a_attribute_value_strings_deletedInstances)
	var a_attribute_value_xhtmls_newInstances []*A_ATTRIBUTE_VALUE_XHTML
	var a_attribute_value_xhtmls_deletedInstances []*A_ATTRIBUTE_VALUE_XHTML

	// parse all staged instances and check if they have a reference
	for a_attribute_value_xhtml := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
		if ref, ok := stage.A_ATTRIBUTE_VALUE_XHTMLs_reference[a_attribute_value_xhtml]; !ok {
			a_attribute_value_xhtmls_newInstances = append(a_attribute_value_xhtmls_newInstances, a_attribute_value_xhtml)
			newInstancesSlice = append(newInstancesSlice, a_attribute_value_xhtml.GongMarshallIdentifier(stage))
			if stage.A_ATTRIBUTE_VALUE_XHTMLs_referenceOrder == nil {
				stage.A_ATTRIBUTE_VALUE_XHTMLs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_XHTML]uint)
			}
			stage.A_ATTRIBUTE_VALUE_XHTMLs_referenceOrder[a_attribute_value_xhtml] = stage.A_ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[a_attribute_value_xhtml]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_attribute_value_xhtml.GongMarshallUnstaging(stage))
			delete(stage.A_ATTRIBUTE_VALUE_XHTMLs_referenceOrder, a_attribute_value_xhtml)
			fieldInitializers, pointersInitializations := a_attribute_value_xhtml.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[ref] = stage.A_ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[a_attribute_value_xhtml]
			diffs := a_attribute_value_xhtml.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_attribute_value_xhtml)
			delete(stage.A_ATTRIBUTE_VALUE_XHTMLMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_attribute_value_xhtml.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ATTRIBUTE_VALUE_XHTMLs_reference {
		if _, ok := stage.A_ATTRIBUTE_VALUE_XHTMLs[ref]; !ok {
			a_attribute_value_xhtmls_deletedInstances = append(a_attribute_value_xhtmls_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_attribute_value_xhtmls_newInstances)
	lenDeletedInstances += len(a_attribute_value_xhtmls_deletedInstances)
	var a_attribute_value_xhtml_1s_newInstances []*A_ATTRIBUTE_VALUE_XHTML_1
	var a_attribute_value_xhtml_1s_deletedInstances []*A_ATTRIBUTE_VALUE_XHTML_1

	// parse all staged instances and check if they have a reference
	for a_attribute_value_xhtml_1 := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		if ref, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1s_reference[a_attribute_value_xhtml_1]; !ok {
			a_attribute_value_xhtml_1s_newInstances = append(a_attribute_value_xhtml_1s_newInstances, a_attribute_value_xhtml_1)
			newInstancesSlice = append(newInstancesSlice, a_attribute_value_xhtml_1.GongMarshallIdentifier(stage))
			if stage.A_ATTRIBUTE_VALUE_XHTML_1s_referenceOrder == nil {
				stage.A_ATTRIBUTE_VALUE_XHTML_1s_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_XHTML_1]uint)
			}
			stage.A_ATTRIBUTE_VALUE_XHTML_1s_referenceOrder[a_attribute_value_xhtml_1] = stage.A_ATTRIBUTE_VALUE_XHTML_1Map_Staged_Order[a_attribute_value_xhtml_1]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_attribute_value_xhtml_1.GongMarshallUnstaging(stage))
			delete(stage.A_ATTRIBUTE_VALUE_XHTML_1s_referenceOrder, a_attribute_value_xhtml_1)
			fieldInitializers, pointersInitializations := a_attribute_value_xhtml_1.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ATTRIBUTE_VALUE_XHTML_1Map_Staged_Order[ref] = stage.A_ATTRIBUTE_VALUE_XHTML_1Map_Staged_Order[a_attribute_value_xhtml_1]
			diffs := a_attribute_value_xhtml_1.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_attribute_value_xhtml_1)
			delete(stage.A_ATTRIBUTE_VALUE_XHTML_1Map_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_attribute_value_xhtml_1.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ATTRIBUTE_VALUE_XHTML_1s_reference {
		if _, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1s[ref]; !ok {
			a_attribute_value_xhtml_1s_deletedInstances = append(a_attribute_value_xhtml_1s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_attribute_value_xhtml_1s_newInstances)
	lenDeletedInstances += len(a_attribute_value_xhtml_1s_deletedInstances)
	var a_childrens_newInstances []*A_CHILDREN
	var a_childrens_deletedInstances []*A_CHILDREN

	// parse all staged instances and check if they have a reference
	for a_children := range stage.A_CHILDRENs {
		if ref, ok := stage.A_CHILDRENs_reference[a_children]; !ok {
			a_childrens_newInstances = append(a_childrens_newInstances, a_children)
			newInstancesSlice = append(newInstancesSlice, a_children.GongMarshallIdentifier(stage))
			if stage.A_CHILDRENs_referenceOrder == nil {
				stage.A_CHILDRENs_referenceOrder = make(map[*A_CHILDREN]uint)
			}
			stage.A_CHILDRENs_referenceOrder[a_children] = stage.A_CHILDRENMap_Staged_Order[a_children]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_children.GongMarshallUnstaging(stage))
			delete(stage.A_CHILDRENs_referenceOrder, a_children)
			fieldInitializers, pointersInitializations := a_children.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_CHILDRENMap_Staged_Order[ref] = stage.A_CHILDRENMap_Staged_Order[a_children]
			diffs := a_children.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_children)
			delete(stage.A_CHILDRENMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_children.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_CHILDRENs_reference {
		if _, ok := stage.A_CHILDRENs[ref]; !ok {
			a_childrens_deletedInstances = append(a_childrens_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_childrens_newInstances)
	lenDeletedInstances += len(a_childrens_deletedInstances)
	var a_core_contents_newInstances []*A_CORE_CONTENT
	var a_core_contents_deletedInstances []*A_CORE_CONTENT

	// parse all staged instances and check if they have a reference
	for a_core_content := range stage.A_CORE_CONTENTs {
		if ref, ok := stage.A_CORE_CONTENTs_reference[a_core_content]; !ok {
			a_core_contents_newInstances = append(a_core_contents_newInstances, a_core_content)
			newInstancesSlice = append(newInstancesSlice, a_core_content.GongMarshallIdentifier(stage))
			if stage.A_CORE_CONTENTs_referenceOrder == nil {
				stage.A_CORE_CONTENTs_referenceOrder = make(map[*A_CORE_CONTENT]uint)
			}
			stage.A_CORE_CONTENTs_referenceOrder[a_core_content] = stage.A_CORE_CONTENTMap_Staged_Order[a_core_content]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_core_content.GongMarshallUnstaging(stage))
			delete(stage.A_CORE_CONTENTs_referenceOrder, a_core_content)
			fieldInitializers, pointersInitializations := a_core_content.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_CORE_CONTENTMap_Staged_Order[ref] = stage.A_CORE_CONTENTMap_Staged_Order[a_core_content]
			diffs := a_core_content.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_core_content)
			delete(stage.A_CORE_CONTENTMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_core_content.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_CORE_CONTENTs_reference {
		if _, ok := stage.A_CORE_CONTENTs[ref]; !ok {
			a_core_contents_deletedInstances = append(a_core_contents_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_core_contents_newInstances)
	lenDeletedInstances += len(a_core_contents_deletedInstances)
	var a_datatypess_newInstances []*A_DATATYPES
	var a_datatypess_deletedInstances []*A_DATATYPES

	// parse all staged instances and check if they have a reference
	for a_datatypes := range stage.A_DATATYPESs {
		if ref, ok := stage.A_DATATYPESs_reference[a_datatypes]; !ok {
			a_datatypess_newInstances = append(a_datatypess_newInstances, a_datatypes)
			newInstancesSlice = append(newInstancesSlice, a_datatypes.GongMarshallIdentifier(stage))
			if stage.A_DATATYPESs_referenceOrder == nil {
				stage.A_DATATYPESs_referenceOrder = make(map[*A_DATATYPES]uint)
			}
			stage.A_DATATYPESs_referenceOrder[a_datatypes] = stage.A_DATATYPESMap_Staged_Order[a_datatypes]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_datatypes.GongMarshallUnstaging(stage))
			delete(stage.A_DATATYPESs_referenceOrder, a_datatypes)
			fieldInitializers, pointersInitializations := a_datatypes.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_DATATYPESMap_Staged_Order[ref] = stage.A_DATATYPESMap_Staged_Order[a_datatypes]
			diffs := a_datatypes.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_datatypes)
			delete(stage.A_DATATYPESMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_datatypes.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_DATATYPESs_reference {
		if _, ok := stage.A_DATATYPESs[ref]; !ok {
			a_datatypess_deletedInstances = append(a_datatypess_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_datatypess_newInstances)
	lenDeletedInstances += len(a_datatypess_deletedInstances)
	var a_datatype_definition_boolean_refs_newInstances []*A_DATATYPE_DEFINITION_BOOLEAN_REF
	var a_datatype_definition_boolean_refs_deletedInstances []*A_DATATYPE_DEFINITION_BOOLEAN_REF

	// parse all staged instances and check if they have a reference
	for a_datatype_definition_boolean_ref := range stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs {
		if ref, ok := stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_reference[a_datatype_definition_boolean_ref]; !ok {
			a_datatype_definition_boolean_refs_newInstances = append(a_datatype_definition_boolean_refs_newInstances, a_datatype_definition_boolean_ref)
			newInstancesSlice = append(newInstancesSlice, a_datatype_definition_boolean_ref.GongMarshallIdentifier(stage))
			if stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_referenceOrder == nil {
				stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_BOOLEAN_REF]uint)
			}
			stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_referenceOrder[a_datatype_definition_boolean_ref] = stage.A_DATATYPE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_datatype_definition_boolean_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_datatype_definition_boolean_ref.GongMarshallUnstaging(stage))
			delete(stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_referenceOrder, a_datatype_definition_boolean_ref)
			fieldInitializers, pointersInitializations := a_datatype_definition_boolean_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_DATATYPE_DEFINITION_BOOLEAN_REFMap_Staged_Order[ref] = stage.A_DATATYPE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_datatype_definition_boolean_ref]
			diffs := a_datatype_definition_boolean_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_datatype_definition_boolean_ref)
			delete(stage.A_DATATYPE_DEFINITION_BOOLEAN_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_datatype_definition_boolean_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_reference {
		if _, ok := stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs[ref]; !ok {
			a_datatype_definition_boolean_refs_deletedInstances = append(a_datatype_definition_boolean_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_datatype_definition_boolean_refs_newInstances)
	lenDeletedInstances += len(a_datatype_definition_boolean_refs_deletedInstances)
	var a_datatype_definition_date_refs_newInstances []*A_DATATYPE_DEFINITION_DATE_REF
	var a_datatype_definition_date_refs_deletedInstances []*A_DATATYPE_DEFINITION_DATE_REF

	// parse all staged instances and check if they have a reference
	for a_datatype_definition_date_ref := range stage.A_DATATYPE_DEFINITION_DATE_REFs {
		if ref, ok := stage.A_DATATYPE_DEFINITION_DATE_REFs_reference[a_datatype_definition_date_ref]; !ok {
			a_datatype_definition_date_refs_newInstances = append(a_datatype_definition_date_refs_newInstances, a_datatype_definition_date_ref)
			newInstancesSlice = append(newInstancesSlice, a_datatype_definition_date_ref.GongMarshallIdentifier(stage))
			if stage.A_DATATYPE_DEFINITION_DATE_REFs_referenceOrder == nil {
				stage.A_DATATYPE_DEFINITION_DATE_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_DATE_REF]uint)
			}
			stage.A_DATATYPE_DEFINITION_DATE_REFs_referenceOrder[a_datatype_definition_date_ref] = stage.A_DATATYPE_DEFINITION_DATE_REFMap_Staged_Order[a_datatype_definition_date_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_datatype_definition_date_ref.GongMarshallUnstaging(stage))
			delete(stage.A_DATATYPE_DEFINITION_DATE_REFs_referenceOrder, a_datatype_definition_date_ref)
			fieldInitializers, pointersInitializations := a_datatype_definition_date_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_DATATYPE_DEFINITION_DATE_REFMap_Staged_Order[ref] = stage.A_DATATYPE_DEFINITION_DATE_REFMap_Staged_Order[a_datatype_definition_date_ref]
			diffs := a_datatype_definition_date_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_datatype_definition_date_ref)
			delete(stage.A_DATATYPE_DEFINITION_DATE_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_datatype_definition_date_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_DATATYPE_DEFINITION_DATE_REFs_reference {
		if _, ok := stage.A_DATATYPE_DEFINITION_DATE_REFs[ref]; !ok {
			a_datatype_definition_date_refs_deletedInstances = append(a_datatype_definition_date_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_datatype_definition_date_refs_newInstances)
	lenDeletedInstances += len(a_datatype_definition_date_refs_deletedInstances)
	var a_datatype_definition_enumeration_refs_newInstances []*A_DATATYPE_DEFINITION_ENUMERATION_REF
	var a_datatype_definition_enumeration_refs_deletedInstances []*A_DATATYPE_DEFINITION_ENUMERATION_REF

	// parse all staged instances and check if they have a reference
	for a_datatype_definition_enumeration_ref := range stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs {
		if ref, ok := stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_reference[a_datatype_definition_enumeration_ref]; !ok {
			a_datatype_definition_enumeration_refs_newInstances = append(a_datatype_definition_enumeration_refs_newInstances, a_datatype_definition_enumeration_ref)
			newInstancesSlice = append(newInstancesSlice, a_datatype_definition_enumeration_ref.GongMarshallIdentifier(stage))
			if stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_referenceOrder == nil {
				stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_ENUMERATION_REF]uint)
			}
			stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_referenceOrder[a_datatype_definition_enumeration_ref] = stage.A_DATATYPE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_datatype_definition_enumeration_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_datatype_definition_enumeration_ref.GongMarshallUnstaging(stage))
			delete(stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_referenceOrder, a_datatype_definition_enumeration_ref)
			fieldInitializers, pointersInitializations := a_datatype_definition_enumeration_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_DATATYPE_DEFINITION_ENUMERATION_REFMap_Staged_Order[ref] = stage.A_DATATYPE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_datatype_definition_enumeration_ref]
			diffs := a_datatype_definition_enumeration_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_datatype_definition_enumeration_ref)
			delete(stage.A_DATATYPE_DEFINITION_ENUMERATION_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_datatype_definition_enumeration_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_reference {
		if _, ok := stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs[ref]; !ok {
			a_datatype_definition_enumeration_refs_deletedInstances = append(a_datatype_definition_enumeration_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_datatype_definition_enumeration_refs_newInstances)
	lenDeletedInstances += len(a_datatype_definition_enumeration_refs_deletedInstances)
	var a_datatype_definition_integer_refs_newInstances []*A_DATATYPE_DEFINITION_INTEGER_REF
	var a_datatype_definition_integer_refs_deletedInstances []*A_DATATYPE_DEFINITION_INTEGER_REF

	// parse all staged instances and check if they have a reference
	for a_datatype_definition_integer_ref := range stage.A_DATATYPE_DEFINITION_INTEGER_REFs {
		if ref, ok := stage.A_DATATYPE_DEFINITION_INTEGER_REFs_reference[a_datatype_definition_integer_ref]; !ok {
			a_datatype_definition_integer_refs_newInstances = append(a_datatype_definition_integer_refs_newInstances, a_datatype_definition_integer_ref)
			newInstancesSlice = append(newInstancesSlice, a_datatype_definition_integer_ref.GongMarshallIdentifier(stage))
			if stage.A_DATATYPE_DEFINITION_INTEGER_REFs_referenceOrder == nil {
				stage.A_DATATYPE_DEFINITION_INTEGER_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_INTEGER_REF]uint)
			}
			stage.A_DATATYPE_DEFINITION_INTEGER_REFs_referenceOrder[a_datatype_definition_integer_ref] = stage.A_DATATYPE_DEFINITION_INTEGER_REFMap_Staged_Order[a_datatype_definition_integer_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_datatype_definition_integer_ref.GongMarshallUnstaging(stage))
			delete(stage.A_DATATYPE_DEFINITION_INTEGER_REFs_referenceOrder, a_datatype_definition_integer_ref)
			fieldInitializers, pointersInitializations := a_datatype_definition_integer_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_DATATYPE_DEFINITION_INTEGER_REFMap_Staged_Order[ref] = stage.A_DATATYPE_DEFINITION_INTEGER_REFMap_Staged_Order[a_datatype_definition_integer_ref]
			diffs := a_datatype_definition_integer_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_datatype_definition_integer_ref)
			delete(stage.A_DATATYPE_DEFINITION_INTEGER_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_datatype_definition_integer_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_DATATYPE_DEFINITION_INTEGER_REFs_reference {
		if _, ok := stage.A_DATATYPE_DEFINITION_INTEGER_REFs[ref]; !ok {
			a_datatype_definition_integer_refs_deletedInstances = append(a_datatype_definition_integer_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_datatype_definition_integer_refs_newInstances)
	lenDeletedInstances += len(a_datatype_definition_integer_refs_deletedInstances)
	var a_datatype_definition_real_refs_newInstances []*A_DATATYPE_DEFINITION_REAL_REF
	var a_datatype_definition_real_refs_deletedInstances []*A_DATATYPE_DEFINITION_REAL_REF

	// parse all staged instances and check if they have a reference
	for a_datatype_definition_real_ref := range stage.A_DATATYPE_DEFINITION_REAL_REFs {
		if ref, ok := stage.A_DATATYPE_DEFINITION_REAL_REFs_reference[a_datatype_definition_real_ref]; !ok {
			a_datatype_definition_real_refs_newInstances = append(a_datatype_definition_real_refs_newInstances, a_datatype_definition_real_ref)
			newInstancesSlice = append(newInstancesSlice, a_datatype_definition_real_ref.GongMarshallIdentifier(stage))
			if stage.A_DATATYPE_DEFINITION_REAL_REFs_referenceOrder == nil {
				stage.A_DATATYPE_DEFINITION_REAL_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_REAL_REF]uint)
			}
			stage.A_DATATYPE_DEFINITION_REAL_REFs_referenceOrder[a_datatype_definition_real_ref] = stage.A_DATATYPE_DEFINITION_REAL_REFMap_Staged_Order[a_datatype_definition_real_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_datatype_definition_real_ref.GongMarshallUnstaging(stage))
			delete(stage.A_DATATYPE_DEFINITION_REAL_REFs_referenceOrder, a_datatype_definition_real_ref)
			fieldInitializers, pointersInitializations := a_datatype_definition_real_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_DATATYPE_DEFINITION_REAL_REFMap_Staged_Order[ref] = stage.A_DATATYPE_DEFINITION_REAL_REFMap_Staged_Order[a_datatype_definition_real_ref]
			diffs := a_datatype_definition_real_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_datatype_definition_real_ref)
			delete(stage.A_DATATYPE_DEFINITION_REAL_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_datatype_definition_real_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_DATATYPE_DEFINITION_REAL_REFs_reference {
		if _, ok := stage.A_DATATYPE_DEFINITION_REAL_REFs[ref]; !ok {
			a_datatype_definition_real_refs_deletedInstances = append(a_datatype_definition_real_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_datatype_definition_real_refs_newInstances)
	lenDeletedInstances += len(a_datatype_definition_real_refs_deletedInstances)
	var a_datatype_definition_string_refs_newInstances []*A_DATATYPE_DEFINITION_STRING_REF
	var a_datatype_definition_string_refs_deletedInstances []*A_DATATYPE_DEFINITION_STRING_REF

	// parse all staged instances and check if they have a reference
	for a_datatype_definition_string_ref := range stage.A_DATATYPE_DEFINITION_STRING_REFs {
		if ref, ok := stage.A_DATATYPE_DEFINITION_STRING_REFs_reference[a_datatype_definition_string_ref]; !ok {
			a_datatype_definition_string_refs_newInstances = append(a_datatype_definition_string_refs_newInstances, a_datatype_definition_string_ref)
			newInstancesSlice = append(newInstancesSlice, a_datatype_definition_string_ref.GongMarshallIdentifier(stage))
			if stage.A_DATATYPE_DEFINITION_STRING_REFs_referenceOrder == nil {
				stage.A_DATATYPE_DEFINITION_STRING_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_STRING_REF]uint)
			}
			stage.A_DATATYPE_DEFINITION_STRING_REFs_referenceOrder[a_datatype_definition_string_ref] = stage.A_DATATYPE_DEFINITION_STRING_REFMap_Staged_Order[a_datatype_definition_string_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_datatype_definition_string_ref.GongMarshallUnstaging(stage))
			delete(stage.A_DATATYPE_DEFINITION_STRING_REFs_referenceOrder, a_datatype_definition_string_ref)
			fieldInitializers, pointersInitializations := a_datatype_definition_string_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_DATATYPE_DEFINITION_STRING_REFMap_Staged_Order[ref] = stage.A_DATATYPE_DEFINITION_STRING_REFMap_Staged_Order[a_datatype_definition_string_ref]
			diffs := a_datatype_definition_string_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_datatype_definition_string_ref)
			delete(stage.A_DATATYPE_DEFINITION_STRING_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_datatype_definition_string_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_DATATYPE_DEFINITION_STRING_REFs_reference {
		if _, ok := stage.A_DATATYPE_DEFINITION_STRING_REFs[ref]; !ok {
			a_datatype_definition_string_refs_deletedInstances = append(a_datatype_definition_string_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_datatype_definition_string_refs_newInstances)
	lenDeletedInstances += len(a_datatype_definition_string_refs_deletedInstances)
	var a_datatype_definition_xhtml_refs_newInstances []*A_DATATYPE_DEFINITION_XHTML_REF
	var a_datatype_definition_xhtml_refs_deletedInstances []*A_DATATYPE_DEFINITION_XHTML_REF

	// parse all staged instances and check if they have a reference
	for a_datatype_definition_xhtml_ref := range stage.A_DATATYPE_DEFINITION_XHTML_REFs {
		if ref, ok := stage.A_DATATYPE_DEFINITION_XHTML_REFs_reference[a_datatype_definition_xhtml_ref]; !ok {
			a_datatype_definition_xhtml_refs_newInstances = append(a_datatype_definition_xhtml_refs_newInstances, a_datatype_definition_xhtml_ref)
			newInstancesSlice = append(newInstancesSlice, a_datatype_definition_xhtml_ref.GongMarshallIdentifier(stage))
			if stage.A_DATATYPE_DEFINITION_XHTML_REFs_referenceOrder == nil {
				stage.A_DATATYPE_DEFINITION_XHTML_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_XHTML_REF]uint)
			}
			stage.A_DATATYPE_DEFINITION_XHTML_REFs_referenceOrder[a_datatype_definition_xhtml_ref] = stage.A_DATATYPE_DEFINITION_XHTML_REFMap_Staged_Order[a_datatype_definition_xhtml_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_datatype_definition_xhtml_ref.GongMarshallUnstaging(stage))
			delete(stage.A_DATATYPE_DEFINITION_XHTML_REFs_referenceOrder, a_datatype_definition_xhtml_ref)
			fieldInitializers, pointersInitializations := a_datatype_definition_xhtml_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_DATATYPE_DEFINITION_XHTML_REFMap_Staged_Order[ref] = stage.A_DATATYPE_DEFINITION_XHTML_REFMap_Staged_Order[a_datatype_definition_xhtml_ref]
			diffs := a_datatype_definition_xhtml_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_datatype_definition_xhtml_ref)
			delete(stage.A_DATATYPE_DEFINITION_XHTML_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_datatype_definition_xhtml_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_DATATYPE_DEFINITION_XHTML_REFs_reference {
		if _, ok := stage.A_DATATYPE_DEFINITION_XHTML_REFs[ref]; !ok {
			a_datatype_definition_xhtml_refs_deletedInstances = append(a_datatype_definition_xhtml_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_datatype_definition_xhtml_refs_newInstances)
	lenDeletedInstances += len(a_datatype_definition_xhtml_refs_deletedInstances)
	var a_editable_attss_newInstances []*A_EDITABLE_ATTS
	var a_editable_attss_deletedInstances []*A_EDITABLE_ATTS

	// parse all staged instances and check if they have a reference
	for a_editable_atts := range stage.A_EDITABLE_ATTSs {
		if ref, ok := stage.A_EDITABLE_ATTSs_reference[a_editable_atts]; !ok {
			a_editable_attss_newInstances = append(a_editable_attss_newInstances, a_editable_atts)
			newInstancesSlice = append(newInstancesSlice, a_editable_atts.GongMarshallIdentifier(stage))
			if stage.A_EDITABLE_ATTSs_referenceOrder == nil {
				stage.A_EDITABLE_ATTSs_referenceOrder = make(map[*A_EDITABLE_ATTS]uint)
			}
			stage.A_EDITABLE_ATTSs_referenceOrder[a_editable_atts] = stage.A_EDITABLE_ATTSMap_Staged_Order[a_editable_atts]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_editable_atts.GongMarshallUnstaging(stage))
			delete(stage.A_EDITABLE_ATTSs_referenceOrder, a_editable_atts)
			fieldInitializers, pointersInitializations := a_editable_atts.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_EDITABLE_ATTSMap_Staged_Order[ref] = stage.A_EDITABLE_ATTSMap_Staged_Order[a_editable_atts]
			diffs := a_editable_atts.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_editable_atts)
			delete(stage.A_EDITABLE_ATTSMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_editable_atts.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_EDITABLE_ATTSs_reference {
		if _, ok := stage.A_EDITABLE_ATTSs[ref]; !ok {
			a_editable_attss_deletedInstances = append(a_editable_attss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_editable_attss_newInstances)
	lenDeletedInstances += len(a_editable_attss_deletedInstances)
	var a_enum_value_refs_newInstances []*A_ENUM_VALUE_REF
	var a_enum_value_refs_deletedInstances []*A_ENUM_VALUE_REF

	// parse all staged instances and check if they have a reference
	for a_enum_value_ref := range stage.A_ENUM_VALUE_REFs {
		if ref, ok := stage.A_ENUM_VALUE_REFs_reference[a_enum_value_ref]; !ok {
			a_enum_value_refs_newInstances = append(a_enum_value_refs_newInstances, a_enum_value_ref)
			newInstancesSlice = append(newInstancesSlice, a_enum_value_ref.GongMarshallIdentifier(stage))
			if stage.A_ENUM_VALUE_REFs_referenceOrder == nil {
				stage.A_ENUM_VALUE_REFs_referenceOrder = make(map[*A_ENUM_VALUE_REF]uint)
			}
			stage.A_ENUM_VALUE_REFs_referenceOrder[a_enum_value_ref] = stage.A_ENUM_VALUE_REFMap_Staged_Order[a_enum_value_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_enum_value_ref.GongMarshallUnstaging(stage))
			delete(stage.A_ENUM_VALUE_REFs_referenceOrder, a_enum_value_ref)
			fieldInitializers, pointersInitializations := a_enum_value_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_ENUM_VALUE_REFMap_Staged_Order[ref] = stage.A_ENUM_VALUE_REFMap_Staged_Order[a_enum_value_ref]
			diffs := a_enum_value_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_enum_value_ref)
			delete(stage.A_ENUM_VALUE_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_enum_value_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_ENUM_VALUE_REFs_reference {
		if _, ok := stage.A_ENUM_VALUE_REFs[ref]; !ok {
			a_enum_value_refs_deletedInstances = append(a_enum_value_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_enum_value_refs_newInstances)
	lenDeletedInstances += len(a_enum_value_refs_deletedInstances)
	var a_objects_newInstances []*A_OBJECT
	var a_objects_deletedInstances []*A_OBJECT

	// parse all staged instances and check if they have a reference
	for a_object := range stage.A_OBJECTs {
		if ref, ok := stage.A_OBJECTs_reference[a_object]; !ok {
			a_objects_newInstances = append(a_objects_newInstances, a_object)
			newInstancesSlice = append(newInstancesSlice, a_object.GongMarshallIdentifier(stage))
			if stage.A_OBJECTs_referenceOrder == nil {
				stage.A_OBJECTs_referenceOrder = make(map[*A_OBJECT]uint)
			}
			stage.A_OBJECTs_referenceOrder[a_object] = stage.A_OBJECTMap_Staged_Order[a_object]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_object.GongMarshallUnstaging(stage))
			delete(stage.A_OBJECTs_referenceOrder, a_object)
			fieldInitializers, pointersInitializations := a_object.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_OBJECTMap_Staged_Order[ref] = stage.A_OBJECTMap_Staged_Order[a_object]
			diffs := a_object.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_object)
			delete(stage.A_OBJECTMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_object.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_OBJECTs_reference {
		if _, ok := stage.A_OBJECTs[ref]; !ok {
			a_objects_deletedInstances = append(a_objects_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_objects_newInstances)
	lenDeletedInstances += len(a_objects_deletedInstances)
	var a_propertiess_newInstances []*A_PROPERTIES
	var a_propertiess_deletedInstances []*A_PROPERTIES

	// parse all staged instances and check if they have a reference
	for a_properties := range stage.A_PROPERTIESs {
		if ref, ok := stage.A_PROPERTIESs_reference[a_properties]; !ok {
			a_propertiess_newInstances = append(a_propertiess_newInstances, a_properties)
			newInstancesSlice = append(newInstancesSlice, a_properties.GongMarshallIdentifier(stage))
			if stage.A_PROPERTIESs_referenceOrder == nil {
				stage.A_PROPERTIESs_referenceOrder = make(map[*A_PROPERTIES]uint)
			}
			stage.A_PROPERTIESs_referenceOrder[a_properties] = stage.A_PROPERTIESMap_Staged_Order[a_properties]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_properties.GongMarshallUnstaging(stage))
			delete(stage.A_PROPERTIESs_referenceOrder, a_properties)
			fieldInitializers, pointersInitializations := a_properties.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_PROPERTIESMap_Staged_Order[ref] = stage.A_PROPERTIESMap_Staged_Order[a_properties]
			diffs := a_properties.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_properties)
			delete(stage.A_PROPERTIESMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_properties.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_PROPERTIESs_reference {
		if _, ok := stage.A_PROPERTIESs[ref]; !ok {
			a_propertiess_deletedInstances = append(a_propertiess_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_propertiess_newInstances)
	lenDeletedInstances += len(a_propertiess_deletedInstances)
	var a_relation_group_type_refs_newInstances []*A_RELATION_GROUP_TYPE_REF
	var a_relation_group_type_refs_deletedInstances []*A_RELATION_GROUP_TYPE_REF

	// parse all staged instances and check if they have a reference
	for a_relation_group_type_ref := range stage.A_RELATION_GROUP_TYPE_REFs {
		if ref, ok := stage.A_RELATION_GROUP_TYPE_REFs_reference[a_relation_group_type_ref]; !ok {
			a_relation_group_type_refs_newInstances = append(a_relation_group_type_refs_newInstances, a_relation_group_type_ref)
			newInstancesSlice = append(newInstancesSlice, a_relation_group_type_ref.GongMarshallIdentifier(stage))
			if stage.A_RELATION_GROUP_TYPE_REFs_referenceOrder == nil {
				stage.A_RELATION_GROUP_TYPE_REFs_referenceOrder = make(map[*A_RELATION_GROUP_TYPE_REF]uint)
			}
			stage.A_RELATION_GROUP_TYPE_REFs_referenceOrder[a_relation_group_type_ref] = stage.A_RELATION_GROUP_TYPE_REFMap_Staged_Order[a_relation_group_type_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_relation_group_type_ref.GongMarshallUnstaging(stage))
			delete(stage.A_RELATION_GROUP_TYPE_REFs_referenceOrder, a_relation_group_type_ref)
			fieldInitializers, pointersInitializations := a_relation_group_type_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_RELATION_GROUP_TYPE_REFMap_Staged_Order[ref] = stage.A_RELATION_GROUP_TYPE_REFMap_Staged_Order[a_relation_group_type_ref]
			diffs := a_relation_group_type_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_relation_group_type_ref)
			delete(stage.A_RELATION_GROUP_TYPE_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_relation_group_type_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_RELATION_GROUP_TYPE_REFs_reference {
		if _, ok := stage.A_RELATION_GROUP_TYPE_REFs[ref]; !ok {
			a_relation_group_type_refs_deletedInstances = append(a_relation_group_type_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_relation_group_type_refs_newInstances)
	lenDeletedInstances += len(a_relation_group_type_refs_deletedInstances)
	var a_source_1s_newInstances []*A_SOURCE_1
	var a_source_1s_deletedInstances []*A_SOURCE_1

	// parse all staged instances and check if they have a reference
	for a_source_1 := range stage.A_SOURCE_1s {
		if ref, ok := stage.A_SOURCE_1s_reference[a_source_1]; !ok {
			a_source_1s_newInstances = append(a_source_1s_newInstances, a_source_1)
			newInstancesSlice = append(newInstancesSlice, a_source_1.GongMarshallIdentifier(stage))
			if stage.A_SOURCE_1s_referenceOrder == nil {
				stage.A_SOURCE_1s_referenceOrder = make(map[*A_SOURCE_1]uint)
			}
			stage.A_SOURCE_1s_referenceOrder[a_source_1] = stage.A_SOURCE_1Map_Staged_Order[a_source_1]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_source_1.GongMarshallUnstaging(stage))
			delete(stage.A_SOURCE_1s_referenceOrder, a_source_1)
			fieldInitializers, pointersInitializations := a_source_1.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_SOURCE_1Map_Staged_Order[ref] = stage.A_SOURCE_1Map_Staged_Order[a_source_1]
			diffs := a_source_1.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_source_1)
			delete(stage.A_SOURCE_1Map_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_source_1.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_SOURCE_1s_reference {
		if _, ok := stage.A_SOURCE_1s[ref]; !ok {
			a_source_1s_deletedInstances = append(a_source_1s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_source_1s_newInstances)
	lenDeletedInstances += len(a_source_1s_deletedInstances)
	var a_source_specification_1s_newInstances []*A_SOURCE_SPECIFICATION_1
	var a_source_specification_1s_deletedInstances []*A_SOURCE_SPECIFICATION_1

	// parse all staged instances and check if they have a reference
	for a_source_specification_1 := range stage.A_SOURCE_SPECIFICATION_1s {
		if ref, ok := stage.A_SOURCE_SPECIFICATION_1s_reference[a_source_specification_1]; !ok {
			a_source_specification_1s_newInstances = append(a_source_specification_1s_newInstances, a_source_specification_1)
			newInstancesSlice = append(newInstancesSlice, a_source_specification_1.GongMarshallIdentifier(stage))
			if stage.A_SOURCE_SPECIFICATION_1s_referenceOrder == nil {
				stage.A_SOURCE_SPECIFICATION_1s_referenceOrder = make(map[*A_SOURCE_SPECIFICATION_1]uint)
			}
			stage.A_SOURCE_SPECIFICATION_1s_referenceOrder[a_source_specification_1] = stage.A_SOURCE_SPECIFICATION_1Map_Staged_Order[a_source_specification_1]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_source_specification_1.GongMarshallUnstaging(stage))
			delete(stage.A_SOURCE_SPECIFICATION_1s_referenceOrder, a_source_specification_1)
			fieldInitializers, pointersInitializations := a_source_specification_1.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_SOURCE_SPECIFICATION_1Map_Staged_Order[ref] = stage.A_SOURCE_SPECIFICATION_1Map_Staged_Order[a_source_specification_1]
			diffs := a_source_specification_1.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_source_specification_1)
			delete(stage.A_SOURCE_SPECIFICATION_1Map_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_source_specification_1.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_SOURCE_SPECIFICATION_1s_reference {
		if _, ok := stage.A_SOURCE_SPECIFICATION_1s[ref]; !ok {
			a_source_specification_1s_deletedInstances = append(a_source_specification_1s_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_source_specification_1s_newInstances)
	lenDeletedInstances += len(a_source_specification_1s_deletedInstances)
	var a_specificationss_newInstances []*A_SPECIFICATIONS
	var a_specificationss_deletedInstances []*A_SPECIFICATIONS

	// parse all staged instances and check if they have a reference
	for a_specifications := range stage.A_SPECIFICATIONSs {
		if ref, ok := stage.A_SPECIFICATIONSs_reference[a_specifications]; !ok {
			a_specificationss_newInstances = append(a_specificationss_newInstances, a_specifications)
			newInstancesSlice = append(newInstancesSlice, a_specifications.GongMarshallIdentifier(stage))
			if stage.A_SPECIFICATIONSs_referenceOrder == nil {
				stage.A_SPECIFICATIONSs_referenceOrder = make(map[*A_SPECIFICATIONS]uint)
			}
			stage.A_SPECIFICATIONSs_referenceOrder[a_specifications] = stage.A_SPECIFICATIONSMap_Staged_Order[a_specifications]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_specifications.GongMarshallUnstaging(stage))
			delete(stage.A_SPECIFICATIONSs_referenceOrder, a_specifications)
			fieldInitializers, pointersInitializations := a_specifications.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_SPECIFICATIONSMap_Staged_Order[ref] = stage.A_SPECIFICATIONSMap_Staged_Order[a_specifications]
			diffs := a_specifications.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_specifications)
			delete(stage.A_SPECIFICATIONSMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_specifications.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_SPECIFICATIONSs_reference {
		if _, ok := stage.A_SPECIFICATIONSs[ref]; !ok {
			a_specificationss_deletedInstances = append(a_specificationss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_specificationss_newInstances)
	lenDeletedInstances += len(a_specificationss_deletedInstances)
	var a_specification_type_refs_newInstances []*A_SPECIFICATION_TYPE_REF
	var a_specification_type_refs_deletedInstances []*A_SPECIFICATION_TYPE_REF

	// parse all staged instances and check if they have a reference
	for a_specification_type_ref := range stage.A_SPECIFICATION_TYPE_REFs {
		if ref, ok := stage.A_SPECIFICATION_TYPE_REFs_reference[a_specification_type_ref]; !ok {
			a_specification_type_refs_newInstances = append(a_specification_type_refs_newInstances, a_specification_type_ref)
			newInstancesSlice = append(newInstancesSlice, a_specification_type_ref.GongMarshallIdentifier(stage))
			if stage.A_SPECIFICATION_TYPE_REFs_referenceOrder == nil {
				stage.A_SPECIFICATION_TYPE_REFs_referenceOrder = make(map[*A_SPECIFICATION_TYPE_REF]uint)
			}
			stage.A_SPECIFICATION_TYPE_REFs_referenceOrder[a_specification_type_ref] = stage.A_SPECIFICATION_TYPE_REFMap_Staged_Order[a_specification_type_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_specification_type_ref.GongMarshallUnstaging(stage))
			delete(stage.A_SPECIFICATION_TYPE_REFs_referenceOrder, a_specification_type_ref)
			fieldInitializers, pointersInitializations := a_specification_type_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_SPECIFICATION_TYPE_REFMap_Staged_Order[ref] = stage.A_SPECIFICATION_TYPE_REFMap_Staged_Order[a_specification_type_ref]
			diffs := a_specification_type_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_specification_type_ref)
			delete(stage.A_SPECIFICATION_TYPE_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_specification_type_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_SPECIFICATION_TYPE_REFs_reference {
		if _, ok := stage.A_SPECIFICATION_TYPE_REFs[ref]; !ok {
			a_specification_type_refs_deletedInstances = append(a_specification_type_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_specification_type_refs_newInstances)
	lenDeletedInstances += len(a_specification_type_refs_deletedInstances)
	var a_specified_valuess_newInstances []*A_SPECIFIED_VALUES
	var a_specified_valuess_deletedInstances []*A_SPECIFIED_VALUES

	// parse all staged instances and check if they have a reference
	for a_specified_values := range stage.A_SPECIFIED_VALUESs {
		if ref, ok := stage.A_SPECIFIED_VALUESs_reference[a_specified_values]; !ok {
			a_specified_valuess_newInstances = append(a_specified_valuess_newInstances, a_specified_values)
			newInstancesSlice = append(newInstancesSlice, a_specified_values.GongMarshallIdentifier(stage))
			if stage.A_SPECIFIED_VALUESs_referenceOrder == nil {
				stage.A_SPECIFIED_VALUESs_referenceOrder = make(map[*A_SPECIFIED_VALUES]uint)
			}
			stage.A_SPECIFIED_VALUESs_referenceOrder[a_specified_values] = stage.A_SPECIFIED_VALUESMap_Staged_Order[a_specified_values]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_specified_values.GongMarshallUnstaging(stage))
			delete(stage.A_SPECIFIED_VALUESs_referenceOrder, a_specified_values)
			fieldInitializers, pointersInitializations := a_specified_values.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_SPECIFIED_VALUESMap_Staged_Order[ref] = stage.A_SPECIFIED_VALUESMap_Staged_Order[a_specified_values]
			diffs := a_specified_values.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_specified_values)
			delete(stage.A_SPECIFIED_VALUESMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_specified_values.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_SPECIFIED_VALUESs_reference {
		if _, ok := stage.A_SPECIFIED_VALUESs[ref]; !ok {
			a_specified_valuess_deletedInstances = append(a_specified_valuess_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_specified_valuess_newInstances)
	lenDeletedInstances += len(a_specified_valuess_deletedInstances)
	var a_spec_attributess_newInstances []*A_SPEC_ATTRIBUTES
	var a_spec_attributess_deletedInstances []*A_SPEC_ATTRIBUTES

	// parse all staged instances and check if they have a reference
	for a_spec_attributes := range stage.A_SPEC_ATTRIBUTESs {
		if ref, ok := stage.A_SPEC_ATTRIBUTESs_reference[a_spec_attributes]; !ok {
			a_spec_attributess_newInstances = append(a_spec_attributess_newInstances, a_spec_attributes)
			newInstancesSlice = append(newInstancesSlice, a_spec_attributes.GongMarshallIdentifier(stage))
			if stage.A_SPEC_ATTRIBUTESs_referenceOrder == nil {
				stage.A_SPEC_ATTRIBUTESs_referenceOrder = make(map[*A_SPEC_ATTRIBUTES]uint)
			}
			stage.A_SPEC_ATTRIBUTESs_referenceOrder[a_spec_attributes] = stage.A_SPEC_ATTRIBUTESMap_Staged_Order[a_spec_attributes]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_spec_attributes.GongMarshallUnstaging(stage))
			delete(stage.A_SPEC_ATTRIBUTESs_referenceOrder, a_spec_attributes)
			fieldInitializers, pointersInitializations := a_spec_attributes.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_SPEC_ATTRIBUTESMap_Staged_Order[ref] = stage.A_SPEC_ATTRIBUTESMap_Staged_Order[a_spec_attributes]
			diffs := a_spec_attributes.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_spec_attributes)
			delete(stage.A_SPEC_ATTRIBUTESMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_spec_attributes.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_SPEC_ATTRIBUTESs_reference {
		if _, ok := stage.A_SPEC_ATTRIBUTESs[ref]; !ok {
			a_spec_attributess_deletedInstances = append(a_spec_attributess_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_spec_attributess_newInstances)
	lenDeletedInstances += len(a_spec_attributess_deletedInstances)
	var a_spec_objectss_newInstances []*A_SPEC_OBJECTS
	var a_spec_objectss_deletedInstances []*A_SPEC_OBJECTS

	// parse all staged instances and check if they have a reference
	for a_spec_objects := range stage.A_SPEC_OBJECTSs {
		if ref, ok := stage.A_SPEC_OBJECTSs_reference[a_spec_objects]; !ok {
			a_spec_objectss_newInstances = append(a_spec_objectss_newInstances, a_spec_objects)
			newInstancesSlice = append(newInstancesSlice, a_spec_objects.GongMarshallIdentifier(stage))
			if stage.A_SPEC_OBJECTSs_referenceOrder == nil {
				stage.A_SPEC_OBJECTSs_referenceOrder = make(map[*A_SPEC_OBJECTS]uint)
			}
			stage.A_SPEC_OBJECTSs_referenceOrder[a_spec_objects] = stage.A_SPEC_OBJECTSMap_Staged_Order[a_spec_objects]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_spec_objects.GongMarshallUnstaging(stage))
			delete(stage.A_SPEC_OBJECTSs_referenceOrder, a_spec_objects)
			fieldInitializers, pointersInitializations := a_spec_objects.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_SPEC_OBJECTSMap_Staged_Order[ref] = stage.A_SPEC_OBJECTSMap_Staged_Order[a_spec_objects]
			diffs := a_spec_objects.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_spec_objects)
			delete(stage.A_SPEC_OBJECTSMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_spec_objects.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_SPEC_OBJECTSs_reference {
		if _, ok := stage.A_SPEC_OBJECTSs[ref]; !ok {
			a_spec_objectss_deletedInstances = append(a_spec_objectss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_spec_objectss_newInstances)
	lenDeletedInstances += len(a_spec_objectss_deletedInstances)
	var a_spec_object_type_refs_newInstances []*A_SPEC_OBJECT_TYPE_REF
	var a_spec_object_type_refs_deletedInstances []*A_SPEC_OBJECT_TYPE_REF

	// parse all staged instances and check if they have a reference
	for a_spec_object_type_ref := range stage.A_SPEC_OBJECT_TYPE_REFs {
		if ref, ok := stage.A_SPEC_OBJECT_TYPE_REFs_reference[a_spec_object_type_ref]; !ok {
			a_spec_object_type_refs_newInstances = append(a_spec_object_type_refs_newInstances, a_spec_object_type_ref)
			newInstancesSlice = append(newInstancesSlice, a_spec_object_type_ref.GongMarshallIdentifier(stage))
			if stage.A_SPEC_OBJECT_TYPE_REFs_referenceOrder == nil {
				stage.A_SPEC_OBJECT_TYPE_REFs_referenceOrder = make(map[*A_SPEC_OBJECT_TYPE_REF]uint)
			}
			stage.A_SPEC_OBJECT_TYPE_REFs_referenceOrder[a_spec_object_type_ref] = stage.A_SPEC_OBJECT_TYPE_REFMap_Staged_Order[a_spec_object_type_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_spec_object_type_ref.GongMarshallUnstaging(stage))
			delete(stage.A_SPEC_OBJECT_TYPE_REFs_referenceOrder, a_spec_object_type_ref)
			fieldInitializers, pointersInitializations := a_spec_object_type_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_SPEC_OBJECT_TYPE_REFMap_Staged_Order[ref] = stage.A_SPEC_OBJECT_TYPE_REFMap_Staged_Order[a_spec_object_type_ref]
			diffs := a_spec_object_type_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_spec_object_type_ref)
			delete(stage.A_SPEC_OBJECT_TYPE_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_spec_object_type_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_SPEC_OBJECT_TYPE_REFs_reference {
		if _, ok := stage.A_SPEC_OBJECT_TYPE_REFs[ref]; !ok {
			a_spec_object_type_refs_deletedInstances = append(a_spec_object_type_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_spec_object_type_refs_newInstances)
	lenDeletedInstances += len(a_spec_object_type_refs_deletedInstances)
	var a_spec_relationss_newInstances []*A_SPEC_RELATIONS
	var a_spec_relationss_deletedInstances []*A_SPEC_RELATIONS

	// parse all staged instances and check if they have a reference
	for a_spec_relations := range stage.A_SPEC_RELATIONSs {
		if ref, ok := stage.A_SPEC_RELATIONSs_reference[a_spec_relations]; !ok {
			a_spec_relationss_newInstances = append(a_spec_relationss_newInstances, a_spec_relations)
			newInstancesSlice = append(newInstancesSlice, a_spec_relations.GongMarshallIdentifier(stage))
			if stage.A_SPEC_RELATIONSs_referenceOrder == nil {
				stage.A_SPEC_RELATIONSs_referenceOrder = make(map[*A_SPEC_RELATIONS]uint)
			}
			stage.A_SPEC_RELATIONSs_referenceOrder[a_spec_relations] = stage.A_SPEC_RELATIONSMap_Staged_Order[a_spec_relations]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_spec_relations.GongMarshallUnstaging(stage))
			delete(stage.A_SPEC_RELATIONSs_referenceOrder, a_spec_relations)
			fieldInitializers, pointersInitializations := a_spec_relations.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_SPEC_RELATIONSMap_Staged_Order[ref] = stage.A_SPEC_RELATIONSMap_Staged_Order[a_spec_relations]
			diffs := a_spec_relations.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_spec_relations)
			delete(stage.A_SPEC_RELATIONSMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_spec_relations.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_SPEC_RELATIONSs_reference {
		if _, ok := stage.A_SPEC_RELATIONSs[ref]; !ok {
			a_spec_relationss_deletedInstances = append(a_spec_relationss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_spec_relationss_newInstances)
	lenDeletedInstances += len(a_spec_relationss_deletedInstances)
	var a_spec_relation_groupss_newInstances []*A_SPEC_RELATION_GROUPS
	var a_spec_relation_groupss_deletedInstances []*A_SPEC_RELATION_GROUPS

	// parse all staged instances and check if they have a reference
	for a_spec_relation_groups := range stage.A_SPEC_RELATION_GROUPSs {
		if ref, ok := stage.A_SPEC_RELATION_GROUPSs_reference[a_spec_relation_groups]; !ok {
			a_spec_relation_groupss_newInstances = append(a_spec_relation_groupss_newInstances, a_spec_relation_groups)
			newInstancesSlice = append(newInstancesSlice, a_spec_relation_groups.GongMarshallIdentifier(stage))
			if stage.A_SPEC_RELATION_GROUPSs_referenceOrder == nil {
				stage.A_SPEC_RELATION_GROUPSs_referenceOrder = make(map[*A_SPEC_RELATION_GROUPS]uint)
			}
			stage.A_SPEC_RELATION_GROUPSs_referenceOrder[a_spec_relation_groups] = stage.A_SPEC_RELATION_GROUPSMap_Staged_Order[a_spec_relation_groups]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_spec_relation_groups.GongMarshallUnstaging(stage))
			delete(stage.A_SPEC_RELATION_GROUPSs_referenceOrder, a_spec_relation_groups)
			fieldInitializers, pointersInitializations := a_spec_relation_groups.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_SPEC_RELATION_GROUPSMap_Staged_Order[ref] = stage.A_SPEC_RELATION_GROUPSMap_Staged_Order[a_spec_relation_groups]
			diffs := a_spec_relation_groups.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_spec_relation_groups)
			delete(stage.A_SPEC_RELATION_GROUPSMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_spec_relation_groups.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_SPEC_RELATION_GROUPSs_reference {
		if _, ok := stage.A_SPEC_RELATION_GROUPSs[ref]; !ok {
			a_spec_relation_groupss_deletedInstances = append(a_spec_relation_groupss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_spec_relation_groupss_newInstances)
	lenDeletedInstances += len(a_spec_relation_groupss_deletedInstances)
	var a_spec_relation_refs_newInstances []*A_SPEC_RELATION_REF
	var a_spec_relation_refs_deletedInstances []*A_SPEC_RELATION_REF

	// parse all staged instances and check if they have a reference
	for a_spec_relation_ref := range stage.A_SPEC_RELATION_REFs {
		if ref, ok := stage.A_SPEC_RELATION_REFs_reference[a_spec_relation_ref]; !ok {
			a_spec_relation_refs_newInstances = append(a_spec_relation_refs_newInstances, a_spec_relation_ref)
			newInstancesSlice = append(newInstancesSlice, a_spec_relation_ref.GongMarshallIdentifier(stage))
			if stage.A_SPEC_RELATION_REFs_referenceOrder == nil {
				stage.A_SPEC_RELATION_REFs_referenceOrder = make(map[*A_SPEC_RELATION_REF]uint)
			}
			stage.A_SPEC_RELATION_REFs_referenceOrder[a_spec_relation_ref] = stage.A_SPEC_RELATION_REFMap_Staged_Order[a_spec_relation_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_spec_relation_ref.GongMarshallUnstaging(stage))
			delete(stage.A_SPEC_RELATION_REFs_referenceOrder, a_spec_relation_ref)
			fieldInitializers, pointersInitializations := a_spec_relation_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_SPEC_RELATION_REFMap_Staged_Order[ref] = stage.A_SPEC_RELATION_REFMap_Staged_Order[a_spec_relation_ref]
			diffs := a_spec_relation_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_spec_relation_ref)
			delete(stage.A_SPEC_RELATION_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_spec_relation_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_SPEC_RELATION_REFs_reference {
		if _, ok := stage.A_SPEC_RELATION_REFs[ref]; !ok {
			a_spec_relation_refs_deletedInstances = append(a_spec_relation_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_spec_relation_refs_newInstances)
	lenDeletedInstances += len(a_spec_relation_refs_deletedInstances)
	var a_spec_relation_type_refs_newInstances []*A_SPEC_RELATION_TYPE_REF
	var a_spec_relation_type_refs_deletedInstances []*A_SPEC_RELATION_TYPE_REF

	// parse all staged instances and check if they have a reference
	for a_spec_relation_type_ref := range stage.A_SPEC_RELATION_TYPE_REFs {
		if ref, ok := stage.A_SPEC_RELATION_TYPE_REFs_reference[a_spec_relation_type_ref]; !ok {
			a_spec_relation_type_refs_newInstances = append(a_spec_relation_type_refs_newInstances, a_spec_relation_type_ref)
			newInstancesSlice = append(newInstancesSlice, a_spec_relation_type_ref.GongMarshallIdentifier(stage))
			if stage.A_SPEC_RELATION_TYPE_REFs_referenceOrder == nil {
				stage.A_SPEC_RELATION_TYPE_REFs_referenceOrder = make(map[*A_SPEC_RELATION_TYPE_REF]uint)
			}
			stage.A_SPEC_RELATION_TYPE_REFs_referenceOrder[a_spec_relation_type_ref] = stage.A_SPEC_RELATION_TYPE_REFMap_Staged_Order[a_spec_relation_type_ref]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_spec_relation_type_ref.GongMarshallUnstaging(stage))
			delete(stage.A_SPEC_RELATION_TYPE_REFs_referenceOrder, a_spec_relation_type_ref)
			fieldInitializers, pointersInitializations := a_spec_relation_type_ref.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_SPEC_RELATION_TYPE_REFMap_Staged_Order[ref] = stage.A_SPEC_RELATION_TYPE_REFMap_Staged_Order[a_spec_relation_type_ref]
			diffs := a_spec_relation_type_ref.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_spec_relation_type_ref)
			delete(stage.A_SPEC_RELATION_TYPE_REFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_spec_relation_type_ref.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_SPEC_RELATION_TYPE_REFs_reference {
		if _, ok := stage.A_SPEC_RELATION_TYPE_REFs[ref]; !ok {
			a_spec_relation_type_refs_deletedInstances = append(a_spec_relation_type_refs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_spec_relation_type_refs_newInstances)
	lenDeletedInstances += len(a_spec_relation_type_refs_deletedInstances)
	var a_spec_typess_newInstances []*A_SPEC_TYPES
	var a_spec_typess_deletedInstances []*A_SPEC_TYPES

	// parse all staged instances and check if they have a reference
	for a_spec_types := range stage.A_SPEC_TYPESs {
		if ref, ok := stage.A_SPEC_TYPESs_reference[a_spec_types]; !ok {
			a_spec_typess_newInstances = append(a_spec_typess_newInstances, a_spec_types)
			newInstancesSlice = append(newInstancesSlice, a_spec_types.GongMarshallIdentifier(stage))
			if stage.A_SPEC_TYPESs_referenceOrder == nil {
				stage.A_SPEC_TYPESs_referenceOrder = make(map[*A_SPEC_TYPES]uint)
			}
			stage.A_SPEC_TYPESs_referenceOrder[a_spec_types] = stage.A_SPEC_TYPESMap_Staged_Order[a_spec_types]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_spec_types.GongMarshallUnstaging(stage))
			delete(stage.A_SPEC_TYPESs_referenceOrder, a_spec_types)
			fieldInitializers, pointersInitializations := a_spec_types.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_SPEC_TYPESMap_Staged_Order[ref] = stage.A_SPEC_TYPESMap_Staged_Order[a_spec_types]
			diffs := a_spec_types.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_spec_types)
			delete(stage.A_SPEC_TYPESMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_spec_types.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_SPEC_TYPESs_reference {
		if _, ok := stage.A_SPEC_TYPESs[ref]; !ok {
			a_spec_typess_deletedInstances = append(a_spec_typess_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_spec_typess_newInstances)
	lenDeletedInstances += len(a_spec_typess_deletedInstances)
	var a_the_headers_newInstances []*A_THE_HEADER
	var a_the_headers_deletedInstances []*A_THE_HEADER

	// parse all staged instances and check if they have a reference
	for a_the_header := range stage.A_THE_HEADERs {
		if ref, ok := stage.A_THE_HEADERs_reference[a_the_header]; !ok {
			a_the_headers_newInstances = append(a_the_headers_newInstances, a_the_header)
			newInstancesSlice = append(newInstancesSlice, a_the_header.GongMarshallIdentifier(stage))
			if stage.A_THE_HEADERs_referenceOrder == nil {
				stage.A_THE_HEADERs_referenceOrder = make(map[*A_THE_HEADER]uint)
			}
			stage.A_THE_HEADERs_referenceOrder[a_the_header] = stage.A_THE_HEADERMap_Staged_Order[a_the_header]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_the_header.GongMarshallUnstaging(stage))
			delete(stage.A_THE_HEADERs_referenceOrder, a_the_header)
			fieldInitializers, pointersInitializations := a_the_header.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_THE_HEADERMap_Staged_Order[ref] = stage.A_THE_HEADERMap_Staged_Order[a_the_header]
			diffs := a_the_header.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_the_header)
			delete(stage.A_THE_HEADERMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_the_header.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_THE_HEADERs_reference {
		if _, ok := stage.A_THE_HEADERs[ref]; !ok {
			a_the_headers_deletedInstances = append(a_the_headers_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_the_headers_newInstances)
	lenDeletedInstances += len(a_the_headers_deletedInstances)
	var a_tool_extensionss_newInstances []*A_TOOL_EXTENSIONS
	var a_tool_extensionss_deletedInstances []*A_TOOL_EXTENSIONS

	// parse all staged instances and check if they have a reference
	for a_tool_extensions := range stage.A_TOOL_EXTENSIONSs {
		if ref, ok := stage.A_TOOL_EXTENSIONSs_reference[a_tool_extensions]; !ok {
			a_tool_extensionss_newInstances = append(a_tool_extensionss_newInstances, a_tool_extensions)
			newInstancesSlice = append(newInstancesSlice, a_tool_extensions.GongMarshallIdentifier(stage))
			if stage.A_TOOL_EXTENSIONSs_referenceOrder == nil {
				stage.A_TOOL_EXTENSIONSs_referenceOrder = make(map[*A_TOOL_EXTENSIONS]uint)
			}
			stage.A_TOOL_EXTENSIONSs_referenceOrder[a_tool_extensions] = stage.A_TOOL_EXTENSIONSMap_Staged_Order[a_tool_extensions]
			newInstancesReverseSlice = append(newInstancesReverseSlice, a_tool_extensions.GongMarshallUnstaging(stage))
			delete(stage.A_TOOL_EXTENSIONSs_referenceOrder, a_tool_extensions)
			fieldInitializers, pointersInitializations := a_tool_extensions.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.A_TOOL_EXTENSIONSMap_Staged_Order[ref] = stage.A_TOOL_EXTENSIONSMap_Staged_Order[a_tool_extensions]
			diffs := a_tool_extensions.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, a_tool_extensions)
			delete(stage.A_TOOL_EXTENSIONSMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", a_tool_extensions.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.A_TOOL_EXTENSIONSs_reference {
		if _, ok := stage.A_TOOL_EXTENSIONSs[ref]; !ok {
			a_tool_extensionss_deletedInstances = append(a_tool_extensionss_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(a_tool_extensionss_newInstances)
	lenDeletedInstances += len(a_tool_extensionss_deletedInstances)
	var datatype_definition_booleans_newInstances []*DATATYPE_DEFINITION_BOOLEAN
	var datatype_definition_booleans_deletedInstances []*DATATYPE_DEFINITION_BOOLEAN

	// parse all staged instances and check if they have a reference
	for datatype_definition_boolean := range stage.DATATYPE_DEFINITION_BOOLEANs {
		if ref, ok := stage.DATATYPE_DEFINITION_BOOLEANs_reference[datatype_definition_boolean]; !ok {
			datatype_definition_booleans_newInstances = append(datatype_definition_booleans_newInstances, datatype_definition_boolean)
			newInstancesSlice = append(newInstancesSlice, datatype_definition_boolean.GongMarshallIdentifier(stage))
			if stage.DATATYPE_DEFINITION_BOOLEANs_referenceOrder == nil {
				stage.DATATYPE_DEFINITION_BOOLEANs_referenceOrder = make(map[*DATATYPE_DEFINITION_BOOLEAN]uint)
			}
			stage.DATATYPE_DEFINITION_BOOLEANs_referenceOrder[datatype_definition_boolean] = stage.DATATYPE_DEFINITION_BOOLEANMap_Staged_Order[datatype_definition_boolean]
			newInstancesReverseSlice = append(newInstancesReverseSlice, datatype_definition_boolean.GongMarshallUnstaging(stage))
			delete(stage.DATATYPE_DEFINITION_BOOLEANs_referenceOrder, datatype_definition_boolean)
			fieldInitializers, pointersInitializations := datatype_definition_boolean.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DATATYPE_DEFINITION_BOOLEANMap_Staged_Order[ref] = stage.DATATYPE_DEFINITION_BOOLEANMap_Staged_Order[datatype_definition_boolean]
			diffs := datatype_definition_boolean.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, datatype_definition_boolean)
			delete(stage.DATATYPE_DEFINITION_BOOLEANMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", datatype_definition_boolean.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.DATATYPE_DEFINITION_BOOLEANs_reference {
		if _, ok := stage.DATATYPE_DEFINITION_BOOLEANs[ref]; !ok {
			datatype_definition_booleans_deletedInstances = append(datatype_definition_booleans_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(datatype_definition_booleans_newInstances)
	lenDeletedInstances += len(datatype_definition_booleans_deletedInstances)
	var datatype_definition_dates_newInstances []*DATATYPE_DEFINITION_DATE
	var datatype_definition_dates_deletedInstances []*DATATYPE_DEFINITION_DATE

	// parse all staged instances and check if they have a reference
	for datatype_definition_date := range stage.DATATYPE_DEFINITION_DATEs {
		if ref, ok := stage.DATATYPE_DEFINITION_DATEs_reference[datatype_definition_date]; !ok {
			datatype_definition_dates_newInstances = append(datatype_definition_dates_newInstances, datatype_definition_date)
			newInstancesSlice = append(newInstancesSlice, datatype_definition_date.GongMarshallIdentifier(stage))
			if stage.DATATYPE_DEFINITION_DATEs_referenceOrder == nil {
				stage.DATATYPE_DEFINITION_DATEs_referenceOrder = make(map[*DATATYPE_DEFINITION_DATE]uint)
			}
			stage.DATATYPE_DEFINITION_DATEs_referenceOrder[datatype_definition_date] = stage.DATATYPE_DEFINITION_DATEMap_Staged_Order[datatype_definition_date]
			newInstancesReverseSlice = append(newInstancesReverseSlice, datatype_definition_date.GongMarshallUnstaging(stage))
			delete(stage.DATATYPE_DEFINITION_DATEs_referenceOrder, datatype_definition_date)
			fieldInitializers, pointersInitializations := datatype_definition_date.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DATATYPE_DEFINITION_DATEMap_Staged_Order[ref] = stage.DATATYPE_DEFINITION_DATEMap_Staged_Order[datatype_definition_date]
			diffs := datatype_definition_date.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, datatype_definition_date)
			delete(stage.DATATYPE_DEFINITION_DATEMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", datatype_definition_date.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.DATATYPE_DEFINITION_DATEs_reference {
		if _, ok := stage.DATATYPE_DEFINITION_DATEs[ref]; !ok {
			datatype_definition_dates_deletedInstances = append(datatype_definition_dates_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(datatype_definition_dates_newInstances)
	lenDeletedInstances += len(datatype_definition_dates_deletedInstances)
	var datatype_definition_enumerations_newInstances []*DATATYPE_DEFINITION_ENUMERATION
	var datatype_definition_enumerations_deletedInstances []*DATATYPE_DEFINITION_ENUMERATION

	// parse all staged instances and check if they have a reference
	for datatype_definition_enumeration := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		if ref, ok := stage.DATATYPE_DEFINITION_ENUMERATIONs_reference[datatype_definition_enumeration]; !ok {
			datatype_definition_enumerations_newInstances = append(datatype_definition_enumerations_newInstances, datatype_definition_enumeration)
			newInstancesSlice = append(newInstancesSlice, datatype_definition_enumeration.GongMarshallIdentifier(stage))
			if stage.DATATYPE_DEFINITION_ENUMERATIONs_referenceOrder == nil {
				stage.DATATYPE_DEFINITION_ENUMERATIONs_referenceOrder = make(map[*DATATYPE_DEFINITION_ENUMERATION]uint)
			}
			stage.DATATYPE_DEFINITION_ENUMERATIONs_referenceOrder[datatype_definition_enumeration] = stage.DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order[datatype_definition_enumeration]
			newInstancesReverseSlice = append(newInstancesReverseSlice, datatype_definition_enumeration.GongMarshallUnstaging(stage))
			delete(stage.DATATYPE_DEFINITION_ENUMERATIONs_referenceOrder, datatype_definition_enumeration)
			fieldInitializers, pointersInitializations := datatype_definition_enumeration.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order[ref] = stage.DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order[datatype_definition_enumeration]
			diffs := datatype_definition_enumeration.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, datatype_definition_enumeration)
			delete(stage.DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", datatype_definition_enumeration.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.DATATYPE_DEFINITION_ENUMERATIONs_reference {
		if _, ok := stage.DATATYPE_DEFINITION_ENUMERATIONs[ref]; !ok {
			datatype_definition_enumerations_deletedInstances = append(datatype_definition_enumerations_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(datatype_definition_enumerations_newInstances)
	lenDeletedInstances += len(datatype_definition_enumerations_deletedInstances)
	var datatype_definition_integers_newInstances []*DATATYPE_DEFINITION_INTEGER
	var datatype_definition_integers_deletedInstances []*DATATYPE_DEFINITION_INTEGER

	// parse all staged instances and check if they have a reference
	for datatype_definition_integer := range stage.DATATYPE_DEFINITION_INTEGERs {
		if ref, ok := stage.DATATYPE_DEFINITION_INTEGERs_reference[datatype_definition_integer]; !ok {
			datatype_definition_integers_newInstances = append(datatype_definition_integers_newInstances, datatype_definition_integer)
			newInstancesSlice = append(newInstancesSlice, datatype_definition_integer.GongMarshallIdentifier(stage))
			if stage.DATATYPE_DEFINITION_INTEGERs_referenceOrder == nil {
				stage.DATATYPE_DEFINITION_INTEGERs_referenceOrder = make(map[*DATATYPE_DEFINITION_INTEGER]uint)
			}
			stage.DATATYPE_DEFINITION_INTEGERs_referenceOrder[datatype_definition_integer] = stage.DATATYPE_DEFINITION_INTEGERMap_Staged_Order[datatype_definition_integer]
			newInstancesReverseSlice = append(newInstancesReverseSlice, datatype_definition_integer.GongMarshallUnstaging(stage))
			delete(stage.DATATYPE_DEFINITION_INTEGERs_referenceOrder, datatype_definition_integer)
			fieldInitializers, pointersInitializations := datatype_definition_integer.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DATATYPE_DEFINITION_INTEGERMap_Staged_Order[ref] = stage.DATATYPE_DEFINITION_INTEGERMap_Staged_Order[datatype_definition_integer]
			diffs := datatype_definition_integer.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, datatype_definition_integer)
			delete(stage.DATATYPE_DEFINITION_INTEGERMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", datatype_definition_integer.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.DATATYPE_DEFINITION_INTEGERs_reference {
		if _, ok := stage.DATATYPE_DEFINITION_INTEGERs[ref]; !ok {
			datatype_definition_integers_deletedInstances = append(datatype_definition_integers_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(datatype_definition_integers_newInstances)
	lenDeletedInstances += len(datatype_definition_integers_deletedInstances)
	var datatype_definition_reals_newInstances []*DATATYPE_DEFINITION_REAL
	var datatype_definition_reals_deletedInstances []*DATATYPE_DEFINITION_REAL

	// parse all staged instances and check if they have a reference
	for datatype_definition_real := range stage.DATATYPE_DEFINITION_REALs {
		if ref, ok := stage.DATATYPE_DEFINITION_REALs_reference[datatype_definition_real]; !ok {
			datatype_definition_reals_newInstances = append(datatype_definition_reals_newInstances, datatype_definition_real)
			newInstancesSlice = append(newInstancesSlice, datatype_definition_real.GongMarshallIdentifier(stage))
			if stage.DATATYPE_DEFINITION_REALs_referenceOrder == nil {
				stage.DATATYPE_DEFINITION_REALs_referenceOrder = make(map[*DATATYPE_DEFINITION_REAL]uint)
			}
			stage.DATATYPE_DEFINITION_REALs_referenceOrder[datatype_definition_real] = stage.DATATYPE_DEFINITION_REALMap_Staged_Order[datatype_definition_real]
			newInstancesReverseSlice = append(newInstancesReverseSlice, datatype_definition_real.GongMarshallUnstaging(stage))
			delete(stage.DATATYPE_DEFINITION_REALs_referenceOrder, datatype_definition_real)
			fieldInitializers, pointersInitializations := datatype_definition_real.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DATATYPE_DEFINITION_REALMap_Staged_Order[ref] = stage.DATATYPE_DEFINITION_REALMap_Staged_Order[datatype_definition_real]
			diffs := datatype_definition_real.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, datatype_definition_real)
			delete(stage.DATATYPE_DEFINITION_REALMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", datatype_definition_real.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.DATATYPE_DEFINITION_REALs_reference {
		if _, ok := stage.DATATYPE_DEFINITION_REALs[ref]; !ok {
			datatype_definition_reals_deletedInstances = append(datatype_definition_reals_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(datatype_definition_reals_newInstances)
	lenDeletedInstances += len(datatype_definition_reals_deletedInstances)
	var datatype_definition_strings_newInstances []*DATATYPE_DEFINITION_STRING
	var datatype_definition_strings_deletedInstances []*DATATYPE_DEFINITION_STRING

	// parse all staged instances and check if they have a reference
	for datatype_definition_string := range stage.DATATYPE_DEFINITION_STRINGs {
		if ref, ok := stage.DATATYPE_DEFINITION_STRINGs_reference[datatype_definition_string]; !ok {
			datatype_definition_strings_newInstances = append(datatype_definition_strings_newInstances, datatype_definition_string)
			newInstancesSlice = append(newInstancesSlice, datatype_definition_string.GongMarshallIdentifier(stage))
			if stage.DATATYPE_DEFINITION_STRINGs_referenceOrder == nil {
				stage.DATATYPE_DEFINITION_STRINGs_referenceOrder = make(map[*DATATYPE_DEFINITION_STRING]uint)
			}
			stage.DATATYPE_DEFINITION_STRINGs_referenceOrder[datatype_definition_string] = stage.DATATYPE_DEFINITION_STRINGMap_Staged_Order[datatype_definition_string]
			newInstancesReverseSlice = append(newInstancesReverseSlice, datatype_definition_string.GongMarshallUnstaging(stage))
			delete(stage.DATATYPE_DEFINITION_STRINGs_referenceOrder, datatype_definition_string)
			fieldInitializers, pointersInitializations := datatype_definition_string.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DATATYPE_DEFINITION_STRINGMap_Staged_Order[ref] = stage.DATATYPE_DEFINITION_STRINGMap_Staged_Order[datatype_definition_string]
			diffs := datatype_definition_string.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, datatype_definition_string)
			delete(stage.DATATYPE_DEFINITION_STRINGMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", datatype_definition_string.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.DATATYPE_DEFINITION_STRINGs_reference {
		if _, ok := stage.DATATYPE_DEFINITION_STRINGs[ref]; !ok {
			datatype_definition_strings_deletedInstances = append(datatype_definition_strings_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(datatype_definition_strings_newInstances)
	lenDeletedInstances += len(datatype_definition_strings_deletedInstances)
	var datatype_definition_xhtmls_newInstances []*DATATYPE_DEFINITION_XHTML
	var datatype_definition_xhtmls_deletedInstances []*DATATYPE_DEFINITION_XHTML

	// parse all staged instances and check if they have a reference
	for datatype_definition_xhtml := range stage.DATATYPE_DEFINITION_XHTMLs {
		if ref, ok := stage.DATATYPE_DEFINITION_XHTMLs_reference[datatype_definition_xhtml]; !ok {
			datatype_definition_xhtmls_newInstances = append(datatype_definition_xhtmls_newInstances, datatype_definition_xhtml)
			newInstancesSlice = append(newInstancesSlice, datatype_definition_xhtml.GongMarshallIdentifier(stage))
			if stage.DATATYPE_DEFINITION_XHTMLs_referenceOrder == nil {
				stage.DATATYPE_DEFINITION_XHTMLs_referenceOrder = make(map[*DATATYPE_DEFINITION_XHTML]uint)
			}
			stage.DATATYPE_DEFINITION_XHTMLs_referenceOrder[datatype_definition_xhtml] = stage.DATATYPE_DEFINITION_XHTMLMap_Staged_Order[datatype_definition_xhtml]
			newInstancesReverseSlice = append(newInstancesReverseSlice, datatype_definition_xhtml.GongMarshallUnstaging(stage))
			delete(stage.DATATYPE_DEFINITION_XHTMLs_referenceOrder, datatype_definition_xhtml)
			fieldInitializers, pointersInitializations := datatype_definition_xhtml.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.DATATYPE_DEFINITION_XHTMLMap_Staged_Order[ref] = stage.DATATYPE_DEFINITION_XHTMLMap_Staged_Order[datatype_definition_xhtml]
			diffs := datatype_definition_xhtml.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, datatype_definition_xhtml)
			delete(stage.DATATYPE_DEFINITION_XHTMLMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", datatype_definition_xhtml.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.DATATYPE_DEFINITION_XHTMLs_reference {
		if _, ok := stage.DATATYPE_DEFINITION_XHTMLs[ref]; !ok {
			datatype_definition_xhtmls_deletedInstances = append(datatype_definition_xhtmls_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(datatype_definition_xhtmls_newInstances)
	lenDeletedInstances += len(datatype_definition_xhtmls_deletedInstances)
	var embedded_values_newInstances []*EMBEDDED_VALUE
	var embedded_values_deletedInstances []*EMBEDDED_VALUE

	// parse all staged instances and check if they have a reference
	for embedded_value := range stage.EMBEDDED_VALUEs {
		if ref, ok := stage.EMBEDDED_VALUEs_reference[embedded_value]; !ok {
			embedded_values_newInstances = append(embedded_values_newInstances, embedded_value)
			newInstancesSlice = append(newInstancesSlice, embedded_value.GongMarshallIdentifier(stage))
			if stage.EMBEDDED_VALUEs_referenceOrder == nil {
				stage.EMBEDDED_VALUEs_referenceOrder = make(map[*EMBEDDED_VALUE]uint)
			}
			stage.EMBEDDED_VALUEs_referenceOrder[embedded_value] = stage.EMBEDDED_VALUEMap_Staged_Order[embedded_value]
			newInstancesReverseSlice = append(newInstancesReverseSlice, embedded_value.GongMarshallUnstaging(stage))
			delete(stage.EMBEDDED_VALUEs_referenceOrder, embedded_value)
			fieldInitializers, pointersInitializations := embedded_value.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.EMBEDDED_VALUEMap_Staged_Order[ref] = stage.EMBEDDED_VALUEMap_Staged_Order[embedded_value]
			diffs := embedded_value.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, embedded_value)
			delete(stage.EMBEDDED_VALUEMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", embedded_value.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.EMBEDDED_VALUEs_reference {
		if _, ok := stage.EMBEDDED_VALUEs[ref]; !ok {
			embedded_values_deletedInstances = append(embedded_values_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(embedded_values_newInstances)
	lenDeletedInstances += len(embedded_values_deletedInstances)
	var enum_values_newInstances []*ENUM_VALUE
	var enum_values_deletedInstances []*ENUM_VALUE

	// parse all staged instances and check if they have a reference
	for enum_value := range stage.ENUM_VALUEs {
		if ref, ok := stage.ENUM_VALUEs_reference[enum_value]; !ok {
			enum_values_newInstances = append(enum_values_newInstances, enum_value)
			newInstancesSlice = append(newInstancesSlice, enum_value.GongMarshallIdentifier(stage))
			if stage.ENUM_VALUEs_referenceOrder == nil {
				stage.ENUM_VALUEs_referenceOrder = make(map[*ENUM_VALUE]uint)
			}
			stage.ENUM_VALUEs_referenceOrder[enum_value] = stage.ENUM_VALUEMap_Staged_Order[enum_value]
			newInstancesReverseSlice = append(newInstancesReverseSlice, enum_value.GongMarshallUnstaging(stage))
			delete(stage.ENUM_VALUEs_referenceOrder, enum_value)
			fieldInitializers, pointersInitializations := enum_value.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.ENUM_VALUEMap_Staged_Order[ref] = stage.ENUM_VALUEMap_Staged_Order[enum_value]
			diffs := enum_value.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, enum_value)
			delete(stage.ENUM_VALUEMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", enum_value.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.ENUM_VALUEs_reference {
		if _, ok := stage.ENUM_VALUEs[ref]; !ok {
			enum_values_deletedInstances = append(enum_values_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(enum_values_newInstances)
	lenDeletedInstances += len(enum_values_deletedInstances)
	var relation_groups_newInstances []*RELATION_GROUP
	var relation_groups_deletedInstances []*RELATION_GROUP

	// parse all staged instances and check if they have a reference
	for relation_group := range stage.RELATION_GROUPs {
		if ref, ok := stage.RELATION_GROUPs_reference[relation_group]; !ok {
			relation_groups_newInstances = append(relation_groups_newInstances, relation_group)
			newInstancesSlice = append(newInstancesSlice, relation_group.GongMarshallIdentifier(stage))
			if stage.RELATION_GROUPs_referenceOrder == nil {
				stage.RELATION_GROUPs_referenceOrder = make(map[*RELATION_GROUP]uint)
			}
			stage.RELATION_GROUPs_referenceOrder[relation_group] = stage.RELATION_GROUPMap_Staged_Order[relation_group]
			newInstancesReverseSlice = append(newInstancesReverseSlice, relation_group.GongMarshallUnstaging(stage))
			delete(stage.RELATION_GROUPs_referenceOrder, relation_group)
			fieldInitializers, pointersInitializations := relation_group.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RELATION_GROUPMap_Staged_Order[ref] = stage.RELATION_GROUPMap_Staged_Order[relation_group]
			diffs := relation_group.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, relation_group)
			delete(stage.RELATION_GROUPMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", relation_group.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.RELATION_GROUPs_reference {
		if _, ok := stage.RELATION_GROUPs[ref]; !ok {
			relation_groups_deletedInstances = append(relation_groups_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(relation_groups_newInstances)
	lenDeletedInstances += len(relation_groups_deletedInstances)
	var relation_group_types_newInstances []*RELATION_GROUP_TYPE
	var relation_group_types_deletedInstances []*RELATION_GROUP_TYPE

	// parse all staged instances and check if they have a reference
	for relation_group_type := range stage.RELATION_GROUP_TYPEs {
		if ref, ok := stage.RELATION_GROUP_TYPEs_reference[relation_group_type]; !ok {
			relation_group_types_newInstances = append(relation_group_types_newInstances, relation_group_type)
			newInstancesSlice = append(newInstancesSlice, relation_group_type.GongMarshallIdentifier(stage))
			if stage.RELATION_GROUP_TYPEs_referenceOrder == nil {
				stage.RELATION_GROUP_TYPEs_referenceOrder = make(map[*RELATION_GROUP_TYPE]uint)
			}
			stage.RELATION_GROUP_TYPEs_referenceOrder[relation_group_type] = stage.RELATION_GROUP_TYPEMap_Staged_Order[relation_group_type]
			newInstancesReverseSlice = append(newInstancesReverseSlice, relation_group_type.GongMarshallUnstaging(stage))
			delete(stage.RELATION_GROUP_TYPEs_referenceOrder, relation_group_type)
			fieldInitializers, pointersInitializations := relation_group_type.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.RELATION_GROUP_TYPEMap_Staged_Order[ref] = stage.RELATION_GROUP_TYPEMap_Staged_Order[relation_group_type]
			diffs := relation_group_type.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, relation_group_type)
			delete(stage.RELATION_GROUP_TYPEMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", relation_group_type.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.RELATION_GROUP_TYPEs_reference {
		if _, ok := stage.RELATION_GROUP_TYPEs[ref]; !ok {
			relation_group_types_deletedInstances = append(relation_group_types_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(relation_group_types_newInstances)
	lenDeletedInstances += len(relation_group_types_deletedInstances)
	var req_ifs_newInstances []*REQ_IF
	var req_ifs_deletedInstances []*REQ_IF

	// parse all staged instances and check if they have a reference
	for req_if := range stage.REQ_IFs {
		if ref, ok := stage.REQ_IFs_reference[req_if]; !ok {
			req_ifs_newInstances = append(req_ifs_newInstances, req_if)
			newInstancesSlice = append(newInstancesSlice, req_if.GongMarshallIdentifier(stage))
			if stage.REQ_IFs_referenceOrder == nil {
				stage.REQ_IFs_referenceOrder = make(map[*REQ_IF]uint)
			}
			stage.REQ_IFs_referenceOrder[req_if] = stage.REQ_IFMap_Staged_Order[req_if]
			newInstancesReverseSlice = append(newInstancesReverseSlice, req_if.GongMarshallUnstaging(stage))
			delete(stage.REQ_IFs_referenceOrder, req_if)
			fieldInitializers, pointersInitializations := req_if.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.REQ_IFMap_Staged_Order[ref] = stage.REQ_IFMap_Staged_Order[req_if]
			diffs := req_if.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, req_if)
			delete(stage.REQ_IFMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", req_if.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.REQ_IFs_reference {
		if _, ok := stage.REQ_IFs[ref]; !ok {
			req_ifs_deletedInstances = append(req_ifs_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(req_ifs_newInstances)
	lenDeletedInstances += len(req_ifs_deletedInstances)
	var req_if_contents_newInstances []*REQ_IF_CONTENT
	var req_if_contents_deletedInstances []*REQ_IF_CONTENT

	// parse all staged instances and check if they have a reference
	for req_if_content := range stage.REQ_IF_CONTENTs {
		if ref, ok := stage.REQ_IF_CONTENTs_reference[req_if_content]; !ok {
			req_if_contents_newInstances = append(req_if_contents_newInstances, req_if_content)
			newInstancesSlice = append(newInstancesSlice, req_if_content.GongMarshallIdentifier(stage))
			if stage.REQ_IF_CONTENTs_referenceOrder == nil {
				stage.REQ_IF_CONTENTs_referenceOrder = make(map[*REQ_IF_CONTENT]uint)
			}
			stage.REQ_IF_CONTENTs_referenceOrder[req_if_content] = stage.REQ_IF_CONTENTMap_Staged_Order[req_if_content]
			newInstancesReverseSlice = append(newInstancesReverseSlice, req_if_content.GongMarshallUnstaging(stage))
			delete(stage.REQ_IF_CONTENTs_referenceOrder, req_if_content)
			fieldInitializers, pointersInitializations := req_if_content.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.REQ_IF_CONTENTMap_Staged_Order[ref] = stage.REQ_IF_CONTENTMap_Staged_Order[req_if_content]
			diffs := req_if_content.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, req_if_content)
			delete(stage.REQ_IF_CONTENTMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", req_if_content.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.REQ_IF_CONTENTs_reference {
		if _, ok := stage.REQ_IF_CONTENTs[ref]; !ok {
			req_if_contents_deletedInstances = append(req_if_contents_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(req_if_contents_newInstances)
	lenDeletedInstances += len(req_if_contents_deletedInstances)
	var req_if_headers_newInstances []*REQ_IF_HEADER
	var req_if_headers_deletedInstances []*REQ_IF_HEADER

	// parse all staged instances and check if they have a reference
	for req_if_header := range stage.REQ_IF_HEADERs {
		if ref, ok := stage.REQ_IF_HEADERs_reference[req_if_header]; !ok {
			req_if_headers_newInstances = append(req_if_headers_newInstances, req_if_header)
			newInstancesSlice = append(newInstancesSlice, req_if_header.GongMarshallIdentifier(stage))
			if stage.REQ_IF_HEADERs_referenceOrder == nil {
				stage.REQ_IF_HEADERs_referenceOrder = make(map[*REQ_IF_HEADER]uint)
			}
			stage.REQ_IF_HEADERs_referenceOrder[req_if_header] = stage.REQ_IF_HEADERMap_Staged_Order[req_if_header]
			newInstancesReverseSlice = append(newInstancesReverseSlice, req_if_header.GongMarshallUnstaging(stage))
			delete(stage.REQ_IF_HEADERs_referenceOrder, req_if_header)
			fieldInitializers, pointersInitializations := req_if_header.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.REQ_IF_HEADERMap_Staged_Order[ref] = stage.REQ_IF_HEADERMap_Staged_Order[req_if_header]
			diffs := req_if_header.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, req_if_header)
			delete(stage.REQ_IF_HEADERMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", req_if_header.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.REQ_IF_HEADERs_reference {
		if _, ok := stage.REQ_IF_HEADERs[ref]; !ok {
			req_if_headers_deletedInstances = append(req_if_headers_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(req_if_headers_newInstances)
	lenDeletedInstances += len(req_if_headers_deletedInstances)
	var req_if_tool_extensions_newInstances []*REQ_IF_TOOL_EXTENSION
	var req_if_tool_extensions_deletedInstances []*REQ_IF_TOOL_EXTENSION

	// parse all staged instances and check if they have a reference
	for req_if_tool_extension := range stage.REQ_IF_TOOL_EXTENSIONs {
		if ref, ok := stage.REQ_IF_TOOL_EXTENSIONs_reference[req_if_tool_extension]; !ok {
			req_if_tool_extensions_newInstances = append(req_if_tool_extensions_newInstances, req_if_tool_extension)
			newInstancesSlice = append(newInstancesSlice, req_if_tool_extension.GongMarshallIdentifier(stage))
			if stage.REQ_IF_TOOL_EXTENSIONs_referenceOrder == nil {
				stage.REQ_IF_TOOL_EXTENSIONs_referenceOrder = make(map[*REQ_IF_TOOL_EXTENSION]uint)
			}
			stage.REQ_IF_TOOL_EXTENSIONs_referenceOrder[req_if_tool_extension] = stage.REQ_IF_TOOL_EXTENSIONMap_Staged_Order[req_if_tool_extension]
			newInstancesReverseSlice = append(newInstancesReverseSlice, req_if_tool_extension.GongMarshallUnstaging(stage))
			delete(stage.REQ_IF_TOOL_EXTENSIONs_referenceOrder, req_if_tool_extension)
			fieldInitializers, pointersInitializations := req_if_tool_extension.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.REQ_IF_TOOL_EXTENSIONMap_Staged_Order[ref] = stage.REQ_IF_TOOL_EXTENSIONMap_Staged_Order[req_if_tool_extension]
			diffs := req_if_tool_extension.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, req_if_tool_extension)
			delete(stage.REQ_IF_TOOL_EXTENSIONMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", req_if_tool_extension.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.REQ_IF_TOOL_EXTENSIONs_reference {
		if _, ok := stage.REQ_IF_TOOL_EXTENSIONs[ref]; !ok {
			req_if_tool_extensions_deletedInstances = append(req_if_tool_extensions_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(req_if_tool_extensions_newInstances)
	lenDeletedInstances += len(req_if_tool_extensions_deletedInstances)
	var specifications_newInstances []*SPECIFICATION
	var specifications_deletedInstances []*SPECIFICATION

	// parse all staged instances and check if they have a reference
	for specification := range stage.SPECIFICATIONs {
		if ref, ok := stage.SPECIFICATIONs_reference[specification]; !ok {
			specifications_newInstances = append(specifications_newInstances, specification)
			newInstancesSlice = append(newInstancesSlice, specification.GongMarshallIdentifier(stage))
			if stage.SPECIFICATIONs_referenceOrder == nil {
				stage.SPECIFICATIONs_referenceOrder = make(map[*SPECIFICATION]uint)
			}
			stage.SPECIFICATIONs_referenceOrder[specification] = stage.SPECIFICATIONMap_Staged_Order[specification]
			newInstancesReverseSlice = append(newInstancesReverseSlice, specification.GongMarshallUnstaging(stage))
			delete(stage.SPECIFICATIONs_referenceOrder, specification)
			fieldInitializers, pointersInitializations := specification.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SPECIFICATIONMap_Staged_Order[ref] = stage.SPECIFICATIONMap_Staged_Order[specification]
			diffs := specification.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, specification)
			delete(stage.SPECIFICATIONMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", specification.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.SPECIFICATIONs_reference {
		if _, ok := stage.SPECIFICATIONs[ref]; !ok {
			specifications_deletedInstances = append(specifications_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(specifications_newInstances)
	lenDeletedInstances += len(specifications_deletedInstances)
	var specification_types_newInstances []*SPECIFICATION_TYPE
	var specification_types_deletedInstances []*SPECIFICATION_TYPE

	// parse all staged instances and check if they have a reference
	for specification_type := range stage.SPECIFICATION_TYPEs {
		if ref, ok := stage.SPECIFICATION_TYPEs_reference[specification_type]; !ok {
			specification_types_newInstances = append(specification_types_newInstances, specification_type)
			newInstancesSlice = append(newInstancesSlice, specification_type.GongMarshallIdentifier(stage))
			if stage.SPECIFICATION_TYPEs_referenceOrder == nil {
				stage.SPECIFICATION_TYPEs_referenceOrder = make(map[*SPECIFICATION_TYPE]uint)
			}
			stage.SPECIFICATION_TYPEs_referenceOrder[specification_type] = stage.SPECIFICATION_TYPEMap_Staged_Order[specification_type]
			newInstancesReverseSlice = append(newInstancesReverseSlice, specification_type.GongMarshallUnstaging(stage))
			delete(stage.SPECIFICATION_TYPEs_referenceOrder, specification_type)
			fieldInitializers, pointersInitializations := specification_type.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SPECIFICATION_TYPEMap_Staged_Order[ref] = stage.SPECIFICATION_TYPEMap_Staged_Order[specification_type]
			diffs := specification_type.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, specification_type)
			delete(stage.SPECIFICATION_TYPEMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", specification_type.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.SPECIFICATION_TYPEs_reference {
		if _, ok := stage.SPECIFICATION_TYPEs[ref]; !ok {
			specification_types_deletedInstances = append(specification_types_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(specification_types_newInstances)
	lenDeletedInstances += len(specification_types_deletedInstances)
	var spec_hierarchys_newInstances []*SPEC_HIERARCHY
	var spec_hierarchys_deletedInstances []*SPEC_HIERARCHY

	// parse all staged instances and check if they have a reference
	for spec_hierarchy := range stage.SPEC_HIERARCHYs {
		if ref, ok := stage.SPEC_HIERARCHYs_reference[spec_hierarchy]; !ok {
			spec_hierarchys_newInstances = append(spec_hierarchys_newInstances, spec_hierarchy)
			newInstancesSlice = append(newInstancesSlice, spec_hierarchy.GongMarshallIdentifier(stage))
			if stage.SPEC_HIERARCHYs_referenceOrder == nil {
				stage.SPEC_HIERARCHYs_referenceOrder = make(map[*SPEC_HIERARCHY]uint)
			}
			stage.SPEC_HIERARCHYs_referenceOrder[spec_hierarchy] = stage.SPEC_HIERARCHYMap_Staged_Order[spec_hierarchy]
			newInstancesReverseSlice = append(newInstancesReverseSlice, spec_hierarchy.GongMarshallUnstaging(stage))
			delete(stage.SPEC_HIERARCHYs_referenceOrder, spec_hierarchy)
			fieldInitializers, pointersInitializations := spec_hierarchy.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SPEC_HIERARCHYMap_Staged_Order[ref] = stage.SPEC_HIERARCHYMap_Staged_Order[spec_hierarchy]
			diffs := spec_hierarchy.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, spec_hierarchy)
			delete(stage.SPEC_HIERARCHYMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", spec_hierarchy.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.SPEC_HIERARCHYs_reference {
		if _, ok := stage.SPEC_HIERARCHYs[ref]; !ok {
			spec_hierarchys_deletedInstances = append(spec_hierarchys_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(spec_hierarchys_newInstances)
	lenDeletedInstances += len(spec_hierarchys_deletedInstances)
	var spec_objects_newInstances []*SPEC_OBJECT
	var spec_objects_deletedInstances []*SPEC_OBJECT

	// parse all staged instances and check if they have a reference
	for spec_object := range stage.SPEC_OBJECTs {
		if ref, ok := stage.SPEC_OBJECTs_reference[spec_object]; !ok {
			spec_objects_newInstances = append(spec_objects_newInstances, spec_object)
			newInstancesSlice = append(newInstancesSlice, spec_object.GongMarshallIdentifier(stage))
			if stage.SPEC_OBJECTs_referenceOrder == nil {
				stage.SPEC_OBJECTs_referenceOrder = make(map[*SPEC_OBJECT]uint)
			}
			stage.SPEC_OBJECTs_referenceOrder[spec_object] = stage.SPEC_OBJECTMap_Staged_Order[spec_object]
			newInstancesReverseSlice = append(newInstancesReverseSlice, spec_object.GongMarshallUnstaging(stage))
			delete(stage.SPEC_OBJECTs_referenceOrder, spec_object)
			fieldInitializers, pointersInitializations := spec_object.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SPEC_OBJECTMap_Staged_Order[ref] = stage.SPEC_OBJECTMap_Staged_Order[spec_object]
			diffs := spec_object.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, spec_object)
			delete(stage.SPEC_OBJECTMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", spec_object.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.SPEC_OBJECTs_reference {
		if _, ok := stage.SPEC_OBJECTs[ref]; !ok {
			spec_objects_deletedInstances = append(spec_objects_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(spec_objects_newInstances)
	lenDeletedInstances += len(spec_objects_deletedInstances)
	var spec_object_types_newInstances []*SPEC_OBJECT_TYPE
	var spec_object_types_deletedInstances []*SPEC_OBJECT_TYPE

	// parse all staged instances and check if they have a reference
	for spec_object_type := range stage.SPEC_OBJECT_TYPEs {
		if ref, ok := stage.SPEC_OBJECT_TYPEs_reference[spec_object_type]; !ok {
			spec_object_types_newInstances = append(spec_object_types_newInstances, spec_object_type)
			newInstancesSlice = append(newInstancesSlice, spec_object_type.GongMarshallIdentifier(stage))
			if stage.SPEC_OBJECT_TYPEs_referenceOrder == nil {
				stage.SPEC_OBJECT_TYPEs_referenceOrder = make(map[*SPEC_OBJECT_TYPE]uint)
			}
			stage.SPEC_OBJECT_TYPEs_referenceOrder[spec_object_type] = stage.SPEC_OBJECT_TYPEMap_Staged_Order[spec_object_type]
			newInstancesReverseSlice = append(newInstancesReverseSlice, spec_object_type.GongMarshallUnstaging(stage))
			delete(stage.SPEC_OBJECT_TYPEs_referenceOrder, spec_object_type)
			fieldInitializers, pointersInitializations := spec_object_type.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SPEC_OBJECT_TYPEMap_Staged_Order[ref] = stage.SPEC_OBJECT_TYPEMap_Staged_Order[spec_object_type]
			diffs := spec_object_type.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, spec_object_type)
			delete(stage.SPEC_OBJECT_TYPEMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", spec_object_type.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.SPEC_OBJECT_TYPEs_reference {
		if _, ok := stage.SPEC_OBJECT_TYPEs[ref]; !ok {
			spec_object_types_deletedInstances = append(spec_object_types_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(spec_object_types_newInstances)
	lenDeletedInstances += len(spec_object_types_deletedInstances)
	var spec_relations_newInstances []*SPEC_RELATION
	var spec_relations_deletedInstances []*SPEC_RELATION

	// parse all staged instances and check if they have a reference
	for spec_relation := range stage.SPEC_RELATIONs {
		if ref, ok := stage.SPEC_RELATIONs_reference[spec_relation]; !ok {
			spec_relations_newInstances = append(spec_relations_newInstances, spec_relation)
			newInstancesSlice = append(newInstancesSlice, spec_relation.GongMarshallIdentifier(stage))
			if stage.SPEC_RELATIONs_referenceOrder == nil {
				stage.SPEC_RELATIONs_referenceOrder = make(map[*SPEC_RELATION]uint)
			}
			stage.SPEC_RELATIONs_referenceOrder[spec_relation] = stage.SPEC_RELATIONMap_Staged_Order[spec_relation]
			newInstancesReverseSlice = append(newInstancesReverseSlice, spec_relation.GongMarshallUnstaging(stage))
			delete(stage.SPEC_RELATIONs_referenceOrder, spec_relation)
			fieldInitializers, pointersInitializations := spec_relation.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SPEC_RELATIONMap_Staged_Order[ref] = stage.SPEC_RELATIONMap_Staged_Order[spec_relation]
			diffs := spec_relation.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, spec_relation)
			delete(stage.SPEC_RELATIONMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", spec_relation.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.SPEC_RELATIONs_reference {
		if _, ok := stage.SPEC_RELATIONs[ref]; !ok {
			spec_relations_deletedInstances = append(spec_relations_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(spec_relations_newInstances)
	lenDeletedInstances += len(spec_relations_deletedInstances)
	var spec_relation_types_newInstances []*SPEC_RELATION_TYPE
	var spec_relation_types_deletedInstances []*SPEC_RELATION_TYPE

	// parse all staged instances and check if they have a reference
	for spec_relation_type := range stage.SPEC_RELATION_TYPEs {
		if ref, ok := stage.SPEC_RELATION_TYPEs_reference[spec_relation_type]; !ok {
			spec_relation_types_newInstances = append(spec_relation_types_newInstances, spec_relation_type)
			newInstancesSlice = append(newInstancesSlice, spec_relation_type.GongMarshallIdentifier(stage))
			if stage.SPEC_RELATION_TYPEs_referenceOrder == nil {
				stage.SPEC_RELATION_TYPEs_referenceOrder = make(map[*SPEC_RELATION_TYPE]uint)
			}
			stage.SPEC_RELATION_TYPEs_referenceOrder[spec_relation_type] = stage.SPEC_RELATION_TYPEMap_Staged_Order[spec_relation_type]
			newInstancesReverseSlice = append(newInstancesReverseSlice, spec_relation_type.GongMarshallUnstaging(stage))
			delete(stage.SPEC_RELATION_TYPEs_referenceOrder, spec_relation_type)
			fieldInitializers, pointersInitializations := spec_relation_type.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.SPEC_RELATION_TYPEMap_Staged_Order[ref] = stage.SPEC_RELATION_TYPEMap_Staged_Order[spec_relation_type]
			diffs := spec_relation_type.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, spec_relation_type)
			delete(stage.SPEC_RELATION_TYPEMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", spec_relation_type.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.SPEC_RELATION_TYPEs_reference {
		if _, ok := stage.SPEC_RELATION_TYPEs[ref]; !ok {
			spec_relation_types_deletedInstances = append(spec_relation_types_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(spec_relation_types_newInstances)
	lenDeletedInstances += len(spec_relation_types_deletedInstances)
	var xhtml_contents_newInstances []*XHTML_CONTENT
	var xhtml_contents_deletedInstances []*XHTML_CONTENT

	// parse all staged instances and check if they have a reference
	for xhtml_content := range stage.XHTML_CONTENTs {
		if ref, ok := stage.XHTML_CONTENTs_reference[xhtml_content]; !ok {
			xhtml_contents_newInstances = append(xhtml_contents_newInstances, xhtml_content)
			newInstancesSlice = append(newInstancesSlice, xhtml_content.GongMarshallIdentifier(stage))
			if stage.XHTML_CONTENTs_referenceOrder == nil {
				stage.XHTML_CONTENTs_referenceOrder = make(map[*XHTML_CONTENT]uint)
			}
			stage.XHTML_CONTENTs_referenceOrder[xhtml_content] = stage.XHTML_CONTENTMap_Staged_Order[xhtml_content]
			newInstancesReverseSlice = append(newInstancesReverseSlice, xhtml_content.GongMarshallUnstaging(stage))
			delete(stage.XHTML_CONTENTs_referenceOrder, xhtml_content)
			fieldInitializers, pointersInitializations := xhtml_content.GongMarshallAllFields(stage)
			fieldsEditSlice = append(fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stage.XHTML_CONTENTMap_Staged_Order[ref] = stage.XHTML_CONTENTMap_Staged_Order[xhtml_content]
			diffs := xhtml_content.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, xhtml_content)
			delete(stage.XHTML_CONTENTMap_Staged_Order, ref)
			if len(diffs) > 0 {
				var fieldsEdit string
				fieldsEdit += fmt.Sprintf("\n\t// %s", xhtml_content.GetName())
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				fieldsEditSlice = append(fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					fieldsEditReverseSlice = append(fieldsEditReverseSlice, reverseDiff)
				}
				lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for ref := range stage.XHTML_CONTENTs_reference {
		if _, ok := stage.XHTML_CONTENTs[ref]; !ok {
			xhtml_contents_deletedInstances = append(xhtml_contents_deletedInstances, ref)
			deletedInstancesSlice = append(deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			deletedInstancesReverseSlice = append(deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			fieldsEditReverseSlice = append(fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	lenNewInstances += len(xhtml_contents_newInstances)
	lenDeletedInstances += len(xhtml_contents_deletedInstances)

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
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	stage.ALTERNATIVE_IDs_reference = make(map[*ALTERNATIVE_ID]*ALTERNATIVE_ID)
	stage.ALTERNATIVE_IDs_referenceOrder = make(map[*ALTERNATIVE_ID]uint) // diff Unstage needs the reference order
	for instance := range stage.ALTERNATIVE_IDs {
		stage.ALTERNATIVE_IDs_reference[instance] = instance.GongCopy().(*ALTERNATIVE_ID)
		stage.ALTERNATIVE_IDs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_BOOLEANs_reference = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*ATTRIBUTE_DEFINITION_BOOLEAN)
	stage.ATTRIBUTE_DEFINITION_BOOLEANs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]uint) // diff Unstage needs the reference order
	for instance := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		stage.ATTRIBUTE_DEFINITION_BOOLEANs_reference[instance] = instance.GongCopy().(*ATTRIBUTE_DEFINITION_BOOLEAN)
		stage.ATTRIBUTE_DEFINITION_BOOLEANs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_DATEs_reference = make(map[*ATTRIBUTE_DEFINITION_DATE]*ATTRIBUTE_DEFINITION_DATE)
	stage.ATTRIBUTE_DEFINITION_DATEs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_DATE]uint) // diff Unstage needs the reference order
	for instance := range stage.ATTRIBUTE_DEFINITION_DATEs {
		stage.ATTRIBUTE_DEFINITION_DATEs_reference[instance] = instance.GongCopy().(*ATTRIBUTE_DEFINITION_DATE)
		stage.ATTRIBUTE_DEFINITION_DATEs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_reference = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*ATTRIBUTE_DEFINITION_ENUMERATION)
	stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]uint) // diff Unstage needs the reference order
	for instance := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_reference[instance] = instance.GongCopy().(*ATTRIBUTE_DEFINITION_ENUMERATION)
		stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_INTEGERs_reference = make(map[*ATTRIBUTE_DEFINITION_INTEGER]*ATTRIBUTE_DEFINITION_INTEGER)
	stage.ATTRIBUTE_DEFINITION_INTEGERs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_INTEGER]uint) // diff Unstage needs the reference order
	for instance := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		stage.ATTRIBUTE_DEFINITION_INTEGERs_reference[instance] = instance.GongCopy().(*ATTRIBUTE_DEFINITION_INTEGER)
		stage.ATTRIBUTE_DEFINITION_INTEGERs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_REALs_reference = make(map[*ATTRIBUTE_DEFINITION_REAL]*ATTRIBUTE_DEFINITION_REAL)
	stage.ATTRIBUTE_DEFINITION_REALs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_REAL]uint) // diff Unstage needs the reference order
	for instance := range stage.ATTRIBUTE_DEFINITION_REALs {
		stage.ATTRIBUTE_DEFINITION_REALs_reference[instance] = instance.GongCopy().(*ATTRIBUTE_DEFINITION_REAL)
		stage.ATTRIBUTE_DEFINITION_REALs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_STRINGs_reference = make(map[*ATTRIBUTE_DEFINITION_STRING]*ATTRIBUTE_DEFINITION_STRING)
	stage.ATTRIBUTE_DEFINITION_STRINGs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_STRING]uint) // diff Unstage needs the reference order
	for instance := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		stage.ATTRIBUTE_DEFINITION_STRINGs_reference[instance] = instance.GongCopy().(*ATTRIBUTE_DEFINITION_STRING)
		stage.ATTRIBUTE_DEFINITION_STRINGs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_XHTMLs_reference = make(map[*ATTRIBUTE_DEFINITION_XHTML]*ATTRIBUTE_DEFINITION_XHTML)
	stage.ATTRIBUTE_DEFINITION_XHTMLs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_XHTML]uint) // diff Unstage needs the reference order
	for instance := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		stage.ATTRIBUTE_DEFINITION_XHTMLs_reference[instance] = instance.GongCopy().(*ATTRIBUTE_DEFINITION_XHTML)
		stage.ATTRIBUTE_DEFINITION_XHTMLs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_VALUE_BOOLEANs_reference = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*ATTRIBUTE_VALUE_BOOLEAN)
	stage.ATTRIBUTE_VALUE_BOOLEANs_referenceOrder = make(map[*ATTRIBUTE_VALUE_BOOLEAN]uint) // diff Unstage needs the reference order
	for instance := range stage.ATTRIBUTE_VALUE_BOOLEANs {
		stage.ATTRIBUTE_VALUE_BOOLEANs_reference[instance] = instance.GongCopy().(*ATTRIBUTE_VALUE_BOOLEAN)
		stage.ATTRIBUTE_VALUE_BOOLEANs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_VALUE_DATEs_reference = make(map[*ATTRIBUTE_VALUE_DATE]*ATTRIBUTE_VALUE_DATE)
	stage.ATTRIBUTE_VALUE_DATEs_referenceOrder = make(map[*ATTRIBUTE_VALUE_DATE]uint) // diff Unstage needs the reference order
	for instance := range stage.ATTRIBUTE_VALUE_DATEs {
		stage.ATTRIBUTE_VALUE_DATEs_reference[instance] = instance.GongCopy().(*ATTRIBUTE_VALUE_DATE)
		stage.ATTRIBUTE_VALUE_DATEs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_VALUE_ENUMERATIONs_reference = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*ATTRIBUTE_VALUE_ENUMERATION)
	stage.ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder = make(map[*ATTRIBUTE_VALUE_ENUMERATION]uint) // diff Unstage needs the reference order
	for instance := range stage.ATTRIBUTE_VALUE_ENUMERATIONs {
		stage.ATTRIBUTE_VALUE_ENUMERATIONs_reference[instance] = instance.GongCopy().(*ATTRIBUTE_VALUE_ENUMERATION)
		stage.ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_VALUE_INTEGERs_reference = make(map[*ATTRIBUTE_VALUE_INTEGER]*ATTRIBUTE_VALUE_INTEGER)
	stage.ATTRIBUTE_VALUE_INTEGERs_referenceOrder = make(map[*ATTRIBUTE_VALUE_INTEGER]uint) // diff Unstage needs the reference order
	for instance := range stage.ATTRIBUTE_VALUE_INTEGERs {
		stage.ATTRIBUTE_VALUE_INTEGERs_reference[instance] = instance.GongCopy().(*ATTRIBUTE_VALUE_INTEGER)
		stage.ATTRIBUTE_VALUE_INTEGERs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_VALUE_REALs_reference = make(map[*ATTRIBUTE_VALUE_REAL]*ATTRIBUTE_VALUE_REAL)
	stage.ATTRIBUTE_VALUE_REALs_referenceOrder = make(map[*ATTRIBUTE_VALUE_REAL]uint) // diff Unstage needs the reference order
	for instance := range stage.ATTRIBUTE_VALUE_REALs {
		stage.ATTRIBUTE_VALUE_REALs_reference[instance] = instance.GongCopy().(*ATTRIBUTE_VALUE_REAL)
		stage.ATTRIBUTE_VALUE_REALs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_VALUE_STRINGs_reference = make(map[*ATTRIBUTE_VALUE_STRING]*ATTRIBUTE_VALUE_STRING)
	stage.ATTRIBUTE_VALUE_STRINGs_referenceOrder = make(map[*ATTRIBUTE_VALUE_STRING]uint) // diff Unstage needs the reference order
	for instance := range stage.ATTRIBUTE_VALUE_STRINGs {
		stage.ATTRIBUTE_VALUE_STRINGs_reference[instance] = instance.GongCopy().(*ATTRIBUTE_VALUE_STRING)
		stage.ATTRIBUTE_VALUE_STRINGs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_VALUE_XHTMLs_reference = make(map[*ATTRIBUTE_VALUE_XHTML]*ATTRIBUTE_VALUE_XHTML)
	stage.ATTRIBUTE_VALUE_XHTMLs_referenceOrder = make(map[*ATTRIBUTE_VALUE_XHTML]uint) // diff Unstage needs the reference order
	for instance := range stage.ATTRIBUTE_VALUE_XHTMLs {
		stage.ATTRIBUTE_VALUE_XHTMLs_reference[instance] = instance.GongCopy().(*ATTRIBUTE_VALUE_XHTML)
		stage.ATTRIBUTE_VALUE_XHTMLs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ALTERNATIVE_IDs_reference = make(map[*A_ALTERNATIVE_ID]*A_ALTERNATIVE_ID)
	stage.A_ALTERNATIVE_IDs_referenceOrder = make(map[*A_ALTERNATIVE_ID]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ALTERNATIVE_IDs {
		stage.A_ALTERNATIVE_IDs_reference[instance] = instance.GongCopy().(*A_ALTERNATIVE_ID)
		stage.A_ALTERNATIVE_IDs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_reference = make(map[*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs {
		stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_reference[instance] = instance.GongCopy().(*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
		stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_reference = make(map[*A_ATTRIBUTE_DEFINITION_DATE_REF]*A_ATTRIBUTE_DEFINITION_DATE_REF)
	stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_DATE_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ATTRIBUTE_DEFINITION_DATE_REFs {
		stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_reference[instance] = instance.GongCopy().(*A_ATTRIBUTE_DEFINITION_DATE_REF)
		stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_reference = make(map[*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs {
		stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_reference[instance] = instance.GongCopy().(*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
		stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_reference = make(map[*A_ATTRIBUTE_DEFINITION_INTEGER_REF]*A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_INTEGER_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs {
		stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_reference[instance] = instance.GongCopy().(*A_ATTRIBUTE_DEFINITION_INTEGER_REF)
		stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_reference = make(map[*A_ATTRIBUTE_DEFINITION_REAL_REF]*A_ATTRIBUTE_DEFINITION_REAL_REF)
	stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_REAL_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ATTRIBUTE_DEFINITION_REAL_REFs {
		stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_reference[instance] = instance.GongCopy().(*A_ATTRIBUTE_DEFINITION_REAL_REF)
		stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_reference = make(map[*A_ATTRIBUTE_DEFINITION_STRING_REF]*A_ATTRIBUTE_DEFINITION_STRING_REF)
	stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_STRING_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ATTRIBUTE_DEFINITION_STRING_REFs {
		stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_reference[instance] = instance.GongCopy().(*A_ATTRIBUTE_DEFINITION_STRING_REF)
		stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_reference = make(map[*A_ATTRIBUTE_DEFINITION_XHTML_REF]*A_ATTRIBUTE_DEFINITION_XHTML_REF)
	stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_XHTML_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs {
		stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_reference[instance] = instance.GongCopy().(*A_ATTRIBUTE_DEFINITION_XHTML_REF)
		stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_VALUE_BOOLEANs_reference = make(map[*A_ATTRIBUTE_VALUE_BOOLEAN]*A_ATTRIBUTE_VALUE_BOOLEAN)
	stage.A_ATTRIBUTE_VALUE_BOOLEANs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_BOOLEAN]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ATTRIBUTE_VALUE_BOOLEANs {
		stage.A_ATTRIBUTE_VALUE_BOOLEANs_reference[instance] = instance.GongCopy().(*A_ATTRIBUTE_VALUE_BOOLEAN)
		stage.A_ATTRIBUTE_VALUE_BOOLEANs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_VALUE_DATEs_reference = make(map[*A_ATTRIBUTE_VALUE_DATE]*A_ATTRIBUTE_VALUE_DATE)
	stage.A_ATTRIBUTE_VALUE_DATEs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_DATE]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ATTRIBUTE_VALUE_DATEs {
		stage.A_ATTRIBUTE_VALUE_DATEs_reference[instance] = instance.GongCopy().(*A_ATTRIBUTE_VALUE_DATE)
		stage.A_ATTRIBUTE_VALUE_DATEs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_reference = make(map[*A_ATTRIBUTE_VALUE_ENUMERATION]*A_ATTRIBUTE_VALUE_ENUMERATION)
	stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_ENUMERATION]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
		stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_reference[instance] = instance.GongCopy().(*A_ATTRIBUTE_VALUE_ENUMERATION)
		stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_VALUE_INTEGERs_reference = make(map[*A_ATTRIBUTE_VALUE_INTEGER]*A_ATTRIBUTE_VALUE_INTEGER)
	stage.A_ATTRIBUTE_VALUE_INTEGERs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_INTEGER]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ATTRIBUTE_VALUE_INTEGERs {
		stage.A_ATTRIBUTE_VALUE_INTEGERs_reference[instance] = instance.GongCopy().(*A_ATTRIBUTE_VALUE_INTEGER)
		stage.A_ATTRIBUTE_VALUE_INTEGERs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_VALUE_REALs_reference = make(map[*A_ATTRIBUTE_VALUE_REAL]*A_ATTRIBUTE_VALUE_REAL)
	stage.A_ATTRIBUTE_VALUE_REALs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_REAL]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ATTRIBUTE_VALUE_REALs {
		stage.A_ATTRIBUTE_VALUE_REALs_reference[instance] = instance.GongCopy().(*A_ATTRIBUTE_VALUE_REAL)
		stage.A_ATTRIBUTE_VALUE_REALs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_VALUE_STRINGs_reference = make(map[*A_ATTRIBUTE_VALUE_STRING]*A_ATTRIBUTE_VALUE_STRING)
	stage.A_ATTRIBUTE_VALUE_STRINGs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_STRING]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ATTRIBUTE_VALUE_STRINGs {
		stage.A_ATTRIBUTE_VALUE_STRINGs_reference[instance] = instance.GongCopy().(*A_ATTRIBUTE_VALUE_STRING)
		stage.A_ATTRIBUTE_VALUE_STRINGs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_VALUE_XHTMLs_reference = make(map[*A_ATTRIBUTE_VALUE_XHTML]*A_ATTRIBUTE_VALUE_XHTML)
	stage.A_ATTRIBUTE_VALUE_XHTMLs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_XHTML]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
		stage.A_ATTRIBUTE_VALUE_XHTMLs_reference[instance] = instance.GongCopy().(*A_ATTRIBUTE_VALUE_XHTML)
		stage.A_ATTRIBUTE_VALUE_XHTMLs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_VALUE_XHTML_1s_reference = make(map[*A_ATTRIBUTE_VALUE_XHTML_1]*A_ATTRIBUTE_VALUE_XHTML_1)
	stage.A_ATTRIBUTE_VALUE_XHTML_1s_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_XHTML_1]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		stage.A_ATTRIBUTE_VALUE_XHTML_1s_reference[instance] = instance.GongCopy().(*A_ATTRIBUTE_VALUE_XHTML_1)
		stage.A_ATTRIBUTE_VALUE_XHTML_1s_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_CHILDRENs_reference = make(map[*A_CHILDREN]*A_CHILDREN)
	stage.A_CHILDRENs_referenceOrder = make(map[*A_CHILDREN]uint) // diff Unstage needs the reference order
	for instance := range stage.A_CHILDRENs {
		stage.A_CHILDRENs_reference[instance] = instance.GongCopy().(*A_CHILDREN)
		stage.A_CHILDRENs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_CORE_CONTENTs_reference = make(map[*A_CORE_CONTENT]*A_CORE_CONTENT)
	stage.A_CORE_CONTENTs_referenceOrder = make(map[*A_CORE_CONTENT]uint) // diff Unstage needs the reference order
	for instance := range stage.A_CORE_CONTENTs {
		stage.A_CORE_CONTENTs_reference[instance] = instance.GongCopy().(*A_CORE_CONTENT)
		stage.A_CORE_CONTENTs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_DATATYPESs_reference = make(map[*A_DATATYPES]*A_DATATYPES)
	stage.A_DATATYPESs_referenceOrder = make(map[*A_DATATYPES]uint) // diff Unstage needs the reference order
	for instance := range stage.A_DATATYPESs {
		stage.A_DATATYPESs_reference[instance] = instance.GongCopy().(*A_DATATYPES)
		stage.A_DATATYPESs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_reference = make(map[*A_DATATYPE_DEFINITION_BOOLEAN_REF]*A_DATATYPE_DEFINITION_BOOLEAN_REF)
	stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_BOOLEAN_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs {
		stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_reference[instance] = instance.GongCopy().(*A_DATATYPE_DEFINITION_BOOLEAN_REF)
		stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_DATATYPE_DEFINITION_DATE_REFs_reference = make(map[*A_DATATYPE_DEFINITION_DATE_REF]*A_DATATYPE_DEFINITION_DATE_REF)
	stage.A_DATATYPE_DEFINITION_DATE_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_DATE_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_DATATYPE_DEFINITION_DATE_REFs {
		stage.A_DATATYPE_DEFINITION_DATE_REFs_reference[instance] = instance.GongCopy().(*A_DATATYPE_DEFINITION_DATE_REF)
		stage.A_DATATYPE_DEFINITION_DATE_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_reference = make(map[*A_DATATYPE_DEFINITION_ENUMERATION_REF]*A_DATATYPE_DEFINITION_ENUMERATION_REF)
	stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_ENUMERATION_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs {
		stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_reference[instance] = instance.GongCopy().(*A_DATATYPE_DEFINITION_ENUMERATION_REF)
		stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_DATATYPE_DEFINITION_INTEGER_REFs_reference = make(map[*A_DATATYPE_DEFINITION_INTEGER_REF]*A_DATATYPE_DEFINITION_INTEGER_REF)
	stage.A_DATATYPE_DEFINITION_INTEGER_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_INTEGER_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_DATATYPE_DEFINITION_INTEGER_REFs {
		stage.A_DATATYPE_DEFINITION_INTEGER_REFs_reference[instance] = instance.GongCopy().(*A_DATATYPE_DEFINITION_INTEGER_REF)
		stage.A_DATATYPE_DEFINITION_INTEGER_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_DATATYPE_DEFINITION_REAL_REFs_reference = make(map[*A_DATATYPE_DEFINITION_REAL_REF]*A_DATATYPE_DEFINITION_REAL_REF)
	stage.A_DATATYPE_DEFINITION_REAL_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_REAL_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_DATATYPE_DEFINITION_REAL_REFs {
		stage.A_DATATYPE_DEFINITION_REAL_REFs_reference[instance] = instance.GongCopy().(*A_DATATYPE_DEFINITION_REAL_REF)
		stage.A_DATATYPE_DEFINITION_REAL_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_DATATYPE_DEFINITION_STRING_REFs_reference = make(map[*A_DATATYPE_DEFINITION_STRING_REF]*A_DATATYPE_DEFINITION_STRING_REF)
	stage.A_DATATYPE_DEFINITION_STRING_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_STRING_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_DATATYPE_DEFINITION_STRING_REFs {
		stage.A_DATATYPE_DEFINITION_STRING_REFs_reference[instance] = instance.GongCopy().(*A_DATATYPE_DEFINITION_STRING_REF)
		stage.A_DATATYPE_DEFINITION_STRING_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_DATATYPE_DEFINITION_XHTML_REFs_reference = make(map[*A_DATATYPE_DEFINITION_XHTML_REF]*A_DATATYPE_DEFINITION_XHTML_REF)
	stage.A_DATATYPE_DEFINITION_XHTML_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_XHTML_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_DATATYPE_DEFINITION_XHTML_REFs {
		stage.A_DATATYPE_DEFINITION_XHTML_REFs_reference[instance] = instance.GongCopy().(*A_DATATYPE_DEFINITION_XHTML_REF)
		stage.A_DATATYPE_DEFINITION_XHTML_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_EDITABLE_ATTSs_reference = make(map[*A_EDITABLE_ATTS]*A_EDITABLE_ATTS)
	stage.A_EDITABLE_ATTSs_referenceOrder = make(map[*A_EDITABLE_ATTS]uint) // diff Unstage needs the reference order
	for instance := range stage.A_EDITABLE_ATTSs {
		stage.A_EDITABLE_ATTSs_reference[instance] = instance.GongCopy().(*A_EDITABLE_ATTS)
		stage.A_EDITABLE_ATTSs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_ENUM_VALUE_REFs_reference = make(map[*A_ENUM_VALUE_REF]*A_ENUM_VALUE_REF)
	stage.A_ENUM_VALUE_REFs_referenceOrder = make(map[*A_ENUM_VALUE_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_ENUM_VALUE_REFs {
		stage.A_ENUM_VALUE_REFs_reference[instance] = instance.GongCopy().(*A_ENUM_VALUE_REF)
		stage.A_ENUM_VALUE_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_OBJECTs_reference = make(map[*A_OBJECT]*A_OBJECT)
	stage.A_OBJECTs_referenceOrder = make(map[*A_OBJECT]uint) // diff Unstage needs the reference order
	for instance := range stage.A_OBJECTs {
		stage.A_OBJECTs_reference[instance] = instance.GongCopy().(*A_OBJECT)
		stage.A_OBJECTs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_PROPERTIESs_reference = make(map[*A_PROPERTIES]*A_PROPERTIES)
	stage.A_PROPERTIESs_referenceOrder = make(map[*A_PROPERTIES]uint) // diff Unstage needs the reference order
	for instance := range stage.A_PROPERTIESs {
		stage.A_PROPERTIESs_reference[instance] = instance.GongCopy().(*A_PROPERTIES)
		stage.A_PROPERTIESs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_RELATION_GROUP_TYPE_REFs_reference = make(map[*A_RELATION_GROUP_TYPE_REF]*A_RELATION_GROUP_TYPE_REF)
	stage.A_RELATION_GROUP_TYPE_REFs_referenceOrder = make(map[*A_RELATION_GROUP_TYPE_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_RELATION_GROUP_TYPE_REFs {
		stage.A_RELATION_GROUP_TYPE_REFs_reference[instance] = instance.GongCopy().(*A_RELATION_GROUP_TYPE_REF)
		stage.A_RELATION_GROUP_TYPE_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_SOURCE_1s_reference = make(map[*A_SOURCE_1]*A_SOURCE_1)
	stage.A_SOURCE_1s_referenceOrder = make(map[*A_SOURCE_1]uint) // diff Unstage needs the reference order
	for instance := range stage.A_SOURCE_1s {
		stage.A_SOURCE_1s_reference[instance] = instance.GongCopy().(*A_SOURCE_1)
		stage.A_SOURCE_1s_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_SOURCE_SPECIFICATION_1s_reference = make(map[*A_SOURCE_SPECIFICATION_1]*A_SOURCE_SPECIFICATION_1)
	stage.A_SOURCE_SPECIFICATION_1s_referenceOrder = make(map[*A_SOURCE_SPECIFICATION_1]uint) // diff Unstage needs the reference order
	for instance := range stage.A_SOURCE_SPECIFICATION_1s {
		stage.A_SOURCE_SPECIFICATION_1s_reference[instance] = instance.GongCopy().(*A_SOURCE_SPECIFICATION_1)
		stage.A_SOURCE_SPECIFICATION_1s_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_SPECIFICATIONSs_reference = make(map[*A_SPECIFICATIONS]*A_SPECIFICATIONS)
	stage.A_SPECIFICATIONSs_referenceOrder = make(map[*A_SPECIFICATIONS]uint) // diff Unstage needs the reference order
	for instance := range stage.A_SPECIFICATIONSs {
		stage.A_SPECIFICATIONSs_reference[instance] = instance.GongCopy().(*A_SPECIFICATIONS)
		stage.A_SPECIFICATIONSs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_SPECIFICATION_TYPE_REFs_reference = make(map[*A_SPECIFICATION_TYPE_REF]*A_SPECIFICATION_TYPE_REF)
	stage.A_SPECIFICATION_TYPE_REFs_referenceOrder = make(map[*A_SPECIFICATION_TYPE_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_SPECIFICATION_TYPE_REFs {
		stage.A_SPECIFICATION_TYPE_REFs_reference[instance] = instance.GongCopy().(*A_SPECIFICATION_TYPE_REF)
		stage.A_SPECIFICATION_TYPE_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_SPECIFIED_VALUESs_reference = make(map[*A_SPECIFIED_VALUES]*A_SPECIFIED_VALUES)
	stage.A_SPECIFIED_VALUESs_referenceOrder = make(map[*A_SPECIFIED_VALUES]uint) // diff Unstage needs the reference order
	for instance := range stage.A_SPECIFIED_VALUESs {
		stage.A_SPECIFIED_VALUESs_reference[instance] = instance.GongCopy().(*A_SPECIFIED_VALUES)
		stage.A_SPECIFIED_VALUESs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_SPEC_ATTRIBUTESs_reference = make(map[*A_SPEC_ATTRIBUTES]*A_SPEC_ATTRIBUTES)
	stage.A_SPEC_ATTRIBUTESs_referenceOrder = make(map[*A_SPEC_ATTRIBUTES]uint) // diff Unstage needs the reference order
	for instance := range stage.A_SPEC_ATTRIBUTESs {
		stage.A_SPEC_ATTRIBUTESs_reference[instance] = instance.GongCopy().(*A_SPEC_ATTRIBUTES)
		stage.A_SPEC_ATTRIBUTESs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_SPEC_OBJECTSs_reference = make(map[*A_SPEC_OBJECTS]*A_SPEC_OBJECTS)
	stage.A_SPEC_OBJECTSs_referenceOrder = make(map[*A_SPEC_OBJECTS]uint) // diff Unstage needs the reference order
	for instance := range stage.A_SPEC_OBJECTSs {
		stage.A_SPEC_OBJECTSs_reference[instance] = instance.GongCopy().(*A_SPEC_OBJECTS)
		stage.A_SPEC_OBJECTSs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_SPEC_OBJECT_TYPE_REFs_reference = make(map[*A_SPEC_OBJECT_TYPE_REF]*A_SPEC_OBJECT_TYPE_REF)
	stage.A_SPEC_OBJECT_TYPE_REFs_referenceOrder = make(map[*A_SPEC_OBJECT_TYPE_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_SPEC_OBJECT_TYPE_REFs {
		stage.A_SPEC_OBJECT_TYPE_REFs_reference[instance] = instance.GongCopy().(*A_SPEC_OBJECT_TYPE_REF)
		stage.A_SPEC_OBJECT_TYPE_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_SPEC_RELATIONSs_reference = make(map[*A_SPEC_RELATIONS]*A_SPEC_RELATIONS)
	stage.A_SPEC_RELATIONSs_referenceOrder = make(map[*A_SPEC_RELATIONS]uint) // diff Unstage needs the reference order
	for instance := range stage.A_SPEC_RELATIONSs {
		stage.A_SPEC_RELATIONSs_reference[instance] = instance.GongCopy().(*A_SPEC_RELATIONS)
		stage.A_SPEC_RELATIONSs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_SPEC_RELATION_GROUPSs_reference = make(map[*A_SPEC_RELATION_GROUPS]*A_SPEC_RELATION_GROUPS)
	stage.A_SPEC_RELATION_GROUPSs_referenceOrder = make(map[*A_SPEC_RELATION_GROUPS]uint) // diff Unstage needs the reference order
	for instance := range stage.A_SPEC_RELATION_GROUPSs {
		stage.A_SPEC_RELATION_GROUPSs_reference[instance] = instance.GongCopy().(*A_SPEC_RELATION_GROUPS)
		stage.A_SPEC_RELATION_GROUPSs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_SPEC_RELATION_REFs_reference = make(map[*A_SPEC_RELATION_REF]*A_SPEC_RELATION_REF)
	stage.A_SPEC_RELATION_REFs_referenceOrder = make(map[*A_SPEC_RELATION_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_SPEC_RELATION_REFs {
		stage.A_SPEC_RELATION_REFs_reference[instance] = instance.GongCopy().(*A_SPEC_RELATION_REF)
		stage.A_SPEC_RELATION_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_SPEC_RELATION_TYPE_REFs_reference = make(map[*A_SPEC_RELATION_TYPE_REF]*A_SPEC_RELATION_TYPE_REF)
	stage.A_SPEC_RELATION_TYPE_REFs_referenceOrder = make(map[*A_SPEC_RELATION_TYPE_REF]uint) // diff Unstage needs the reference order
	for instance := range stage.A_SPEC_RELATION_TYPE_REFs {
		stage.A_SPEC_RELATION_TYPE_REFs_reference[instance] = instance.GongCopy().(*A_SPEC_RELATION_TYPE_REF)
		stage.A_SPEC_RELATION_TYPE_REFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_SPEC_TYPESs_reference = make(map[*A_SPEC_TYPES]*A_SPEC_TYPES)
	stage.A_SPEC_TYPESs_referenceOrder = make(map[*A_SPEC_TYPES]uint) // diff Unstage needs the reference order
	for instance := range stage.A_SPEC_TYPESs {
		stage.A_SPEC_TYPESs_reference[instance] = instance.GongCopy().(*A_SPEC_TYPES)
		stage.A_SPEC_TYPESs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_THE_HEADERs_reference = make(map[*A_THE_HEADER]*A_THE_HEADER)
	stage.A_THE_HEADERs_referenceOrder = make(map[*A_THE_HEADER]uint) // diff Unstage needs the reference order
	for instance := range stage.A_THE_HEADERs {
		stage.A_THE_HEADERs_reference[instance] = instance.GongCopy().(*A_THE_HEADER)
		stage.A_THE_HEADERs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.A_TOOL_EXTENSIONSs_reference = make(map[*A_TOOL_EXTENSIONS]*A_TOOL_EXTENSIONS)
	stage.A_TOOL_EXTENSIONSs_referenceOrder = make(map[*A_TOOL_EXTENSIONS]uint) // diff Unstage needs the reference order
	for instance := range stage.A_TOOL_EXTENSIONSs {
		stage.A_TOOL_EXTENSIONSs_reference[instance] = instance.GongCopy().(*A_TOOL_EXTENSIONS)
		stage.A_TOOL_EXTENSIONSs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.DATATYPE_DEFINITION_BOOLEANs_reference = make(map[*DATATYPE_DEFINITION_BOOLEAN]*DATATYPE_DEFINITION_BOOLEAN)
	stage.DATATYPE_DEFINITION_BOOLEANs_referenceOrder = make(map[*DATATYPE_DEFINITION_BOOLEAN]uint) // diff Unstage needs the reference order
	for instance := range stage.DATATYPE_DEFINITION_BOOLEANs {
		stage.DATATYPE_DEFINITION_BOOLEANs_reference[instance] = instance.GongCopy().(*DATATYPE_DEFINITION_BOOLEAN)
		stage.DATATYPE_DEFINITION_BOOLEANs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.DATATYPE_DEFINITION_DATEs_reference = make(map[*DATATYPE_DEFINITION_DATE]*DATATYPE_DEFINITION_DATE)
	stage.DATATYPE_DEFINITION_DATEs_referenceOrder = make(map[*DATATYPE_DEFINITION_DATE]uint) // diff Unstage needs the reference order
	for instance := range stage.DATATYPE_DEFINITION_DATEs {
		stage.DATATYPE_DEFINITION_DATEs_reference[instance] = instance.GongCopy().(*DATATYPE_DEFINITION_DATE)
		stage.DATATYPE_DEFINITION_DATEs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.DATATYPE_DEFINITION_ENUMERATIONs_reference = make(map[*DATATYPE_DEFINITION_ENUMERATION]*DATATYPE_DEFINITION_ENUMERATION)
	stage.DATATYPE_DEFINITION_ENUMERATIONs_referenceOrder = make(map[*DATATYPE_DEFINITION_ENUMERATION]uint) // diff Unstage needs the reference order
	for instance := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		stage.DATATYPE_DEFINITION_ENUMERATIONs_reference[instance] = instance.GongCopy().(*DATATYPE_DEFINITION_ENUMERATION)
		stage.DATATYPE_DEFINITION_ENUMERATIONs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.DATATYPE_DEFINITION_INTEGERs_reference = make(map[*DATATYPE_DEFINITION_INTEGER]*DATATYPE_DEFINITION_INTEGER)
	stage.DATATYPE_DEFINITION_INTEGERs_referenceOrder = make(map[*DATATYPE_DEFINITION_INTEGER]uint) // diff Unstage needs the reference order
	for instance := range stage.DATATYPE_DEFINITION_INTEGERs {
		stage.DATATYPE_DEFINITION_INTEGERs_reference[instance] = instance.GongCopy().(*DATATYPE_DEFINITION_INTEGER)
		stage.DATATYPE_DEFINITION_INTEGERs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.DATATYPE_DEFINITION_REALs_reference = make(map[*DATATYPE_DEFINITION_REAL]*DATATYPE_DEFINITION_REAL)
	stage.DATATYPE_DEFINITION_REALs_referenceOrder = make(map[*DATATYPE_DEFINITION_REAL]uint) // diff Unstage needs the reference order
	for instance := range stage.DATATYPE_DEFINITION_REALs {
		stage.DATATYPE_DEFINITION_REALs_reference[instance] = instance.GongCopy().(*DATATYPE_DEFINITION_REAL)
		stage.DATATYPE_DEFINITION_REALs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.DATATYPE_DEFINITION_STRINGs_reference = make(map[*DATATYPE_DEFINITION_STRING]*DATATYPE_DEFINITION_STRING)
	stage.DATATYPE_DEFINITION_STRINGs_referenceOrder = make(map[*DATATYPE_DEFINITION_STRING]uint) // diff Unstage needs the reference order
	for instance := range stage.DATATYPE_DEFINITION_STRINGs {
		stage.DATATYPE_DEFINITION_STRINGs_reference[instance] = instance.GongCopy().(*DATATYPE_DEFINITION_STRING)
		stage.DATATYPE_DEFINITION_STRINGs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.DATATYPE_DEFINITION_XHTMLs_reference = make(map[*DATATYPE_DEFINITION_XHTML]*DATATYPE_DEFINITION_XHTML)
	stage.DATATYPE_DEFINITION_XHTMLs_referenceOrder = make(map[*DATATYPE_DEFINITION_XHTML]uint) // diff Unstage needs the reference order
	for instance := range stage.DATATYPE_DEFINITION_XHTMLs {
		stage.DATATYPE_DEFINITION_XHTMLs_reference[instance] = instance.GongCopy().(*DATATYPE_DEFINITION_XHTML)
		stage.DATATYPE_DEFINITION_XHTMLs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.EMBEDDED_VALUEs_reference = make(map[*EMBEDDED_VALUE]*EMBEDDED_VALUE)
	stage.EMBEDDED_VALUEs_referenceOrder = make(map[*EMBEDDED_VALUE]uint) // diff Unstage needs the reference order
	for instance := range stage.EMBEDDED_VALUEs {
		stage.EMBEDDED_VALUEs_reference[instance] = instance.GongCopy().(*EMBEDDED_VALUE)
		stage.EMBEDDED_VALUEs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.ENUM_VALUEs_reference = make(map[*ENUM_VALUE]*ENUM_VALUE)
	stage.ENUM_VALUEs_referenceOrder = make(map[*ENUM_VALUE]uint) // diff Unstage needs the reference order
	for instance := range stage.ENUM_VALUEs {
		stage.ENUM_VALUEs_reference[instance] = instance.GongCopy().(*ENUM_VALUE)
		stage.ENUM_VALUEs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.RELATION_GROUPs_reference = make(map[*RELATION_GROUP]*RELATION_GROUP)
	stage.RELATION_GROUPs_referenceOrder = make(map[*RELATION_GROUP]uint) // diff Unstage needs the reference order
	for instance := range stage.RELATION_GROUPs {
		stage.RELATION_GROUPs_reference[instance] = instance.GongCopy().(*RELATION_GROUP)
		stage.RELATION_GROUPs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.RELATION_GROUP_TYPEs_reference = make(map[*RELATION_GROUP_TYPE]*RELATION_GROUP_TYPE)
	stage.RELATION_GROUP_TYPEs_referenceOrder = make(map[*RELATION_GROUP_TYPE]uint) // diff Unstage needs the reference order
	for instance := range stage.RELATION_GROUP_TYPEs {
		stage.RELATION_GROUP_TYPEs_reference[instance] = instance.GongCopy().(*RELATION_GROUP_TYPE)
		stage.RELATION_GROUP_TYPEs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.REQ_IFs_reference = make(map[*REQ_IF]*REQ_IF)
	stage.REQ_IFs_referenceOrder = make(map[*REQ_IF]uint) // diff Unstage needs the reference order
	for instance := range stage.REQ_IFs {
		stage.REQ_IFs_reference[instance] = instance.GongCopy().(*REQ_IF)
		stage.REQ_IFs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.REQ_IF_CONTENTs_reference = make(map[*REQ_IF_CONTENT]*REQ_IF_CONTENT)
	stage.REQ_IF_CONTENTs_referenceOrder = make(map[*REQ_IF_CONTENT]uint) // diff Unstage needs the reference order
	for instance := range stage.REQ_IF_CONTENTs {
		stage.REQ_IF_CONTENTs_reference[instance] = instance.GongCopy().(*REQ_IF_CONTENT)
		stage.REQ_IF_CONTENTs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.REQ_IF_HEADERs_reference = make(map[*REQ_IF_HEADER]*REQ_IF_HEADER)
	stage.REQ_IF_HEADERs_referenceOrder = make(map[*REQ_IF_HEADER]uint) // diff Unstage needs the reference order
	for instance := range stage.REQ_IF_HEADERs {
		stage.REQ_IF_HEADERs_reference[instance] = instance.GongCopy().(*REQ_IF_HEADER)
		stage.REQ_IF_HEADERs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.REQ_IF_TOOL_EXTENSIONs_reference = make(map[*REQ_IF_TOOL_EXTENSION]*REQ_IF_TOOL_EXTENSION)
	stage.REQ_IF_TOOL_EXTENSIONs_referenceOrder = make(map[*REQ_IF_TOOL_EXTENSION]uint) // diff Unstage needs the reference order
	for instance := range stage.REQ_IF_TOOL_EXTENSIONs {
		stage.REQ_IF_TOOL_EXTENSIONs_reference[instance] = instance.GongCopy().(*REQ_IF_TOOL_EXTENSION)
		stage.REQ_IF_TOOL_EXTENSIONs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SPECIFICATIONs_reference = make(map[*SPECIFICATION]*SPECIFICATION)
	stage.SPECIFICATIONs_referenceOrder = make(map[*SPECIFICATION]uint) // diff Unstage needs the reference order
	for instance := range stage.SPECIFICATIONs {
		stage.SPECIFICATIONs_reference[instance] = instance.GongCopy().(*SPECIFICATION)
		stage.SPECIFICATIONs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SPECIFICATION_TYPEs_reference = make(map[*SPECIFICATION_TYPE]*SPECIFICATION_TYPE)
	stage.SPECIFICATION_TYPEs_referenceOrder = make(map[*SPECIFICATION_TYPE]uint) // diff Unstage needs the reference order
	for instance := range stage.SPECIFICATION_TYPEs {
		stage.SPECIFICATION_TYPEs_reference[instance] = instance.GongCopy().(*SPECIFICATION_TYPE)
		stage.SPECIFICATION_TYPEs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SPEC_HIERARCHYs_reference = make(map[*SPEC_HIERARCHY]*SPEC_HIERARCHY)
	stage.SPEC_HIERARCHYs_referenceOrder = make(map[*SPEC_HIERARCHY]uint) // diff Unstage needs the reference order
	for instance := range stage.SPEC_HIERARCHYs {
		stage.SPEC_HIERARCHYs_reference[instance] = instance.GongCopy().(*SPEC_HIERARCHY)
		stage.SPEC_HIERARCHYs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SPEC_OBJECTs_reference = make(map[*SPEC_OBJECT]*SPEC_OBJECT)
	stage.SPEC_OBJECTs_referenceOrder = make(map[*SPEC_OBJECT]uint) // diff Unstage needs the reference order
	for instance := range stage.SPEC_OBJECTs {
		stage.SPEC_OBJECTs_reference[instance] = instance.GongCopy().(*SPEC_OBJECT)
		stage.SPEC_OBJECTs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SPEC_OBJECT_TYPEs_reference = make(map[*SPEC_OBJECT_TYPE]*SPEC_OBJECT_TYPE)
	stage.SPEC_OBJECT_TYPEs_referenceOrder = make(map[*SPEC_OBJECT_TYPE]uint) // diff Unstage needs the reference order
	for instance := range stage.SPEC_OBJECT_TYPEs {
		stage.SPEC_OBJECT_TYPEs_reference[instance] = instance.GongCopy().(*SPEC_OBJECT_TYPE)
		stage.SPEC_OBJECT_TYPEs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SPEC_RELATIONs_reference = make(map[*SPEC_RELATION]*SPEC_RELATION)
	stage.SPEC_RELATIONs_referenceOrder = make(map[*SPEC_RELATION]uint) // diff Unstage needs the reference order
	for instance := range stage.SPEC_RELATIONs {
		stage.SPEC_RELATIONs_reference[instance] = instance.GongCopy().(*SPEC_RELATION)
		stage.SPEC_RELATIONs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.SPEC_RELATION_TYPEs_reference = make(map[*SPEC_RELATION_TYPE]*SPEC_RELATION_TYPE)
	stage.SPEC_RELATION_TYPEs_referenceOrder = make(map[*SPEC_RELATION_TYPE]uint) // diff Unstage needs the reference order
	for instance := range stage.SPEC_RELATION_TYPEs {
		stage.SPEC_RELATION_TYPEs_reference[instance] = instance.GongCopy().(*SPEC_RELATION_TYPE)
		stage.SPEC_RELATION_TYPEs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

	stage.XHTML_CONTENTs_reference = make(map[*XHTML_CONTENT]*XHTML_CONTENT)
	stage.XHTML_CONTENTs_referenceOrder = make(map[*XHTML_CONTENT]uint) // diff Unstage needs the reference order
	for instance := range stage.XHTML_CONTENTs {
		stage.XHTML_CONTENTs_reference[instance] = instance.GongCopy().(*XHTML_CONTENT)
		stage.XHTML_CONTENTs_referenceOrder[instance] = instance.GongGetOrder(stage)
	}

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
	if order, ok := stage.ALTERNATIVE_IDMap_Staged_Order[alternative_id]; ok {
		return order
	}
	return stage.ALTERNATIVE_IDs_referenceOrder[alternative_id]
}

func (alternative_id *ALTERNATIVE_ID) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ALTERNATIVE_IDs_referenceOrder[alternative_id]
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_BOOLEANMap_Staged_Order[attribute_definition_boolean]; ok {
		return order
	}
	return stage.ATTRIBUTE_DEFINITION_BOOLEANs_referenceOrder[attribute_definition_boolean]
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ATTRIBUTE_DEFINITION_BOOLEANs_referenceOrder[attribute_definition_boolean]
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_DATEMap_Staged_Order[attribute_definition_date]; ok {
		return order
	}
	return stage.ATTRIBUTE_DEFINITION_DATEs_referenceOrder[attribute_definition_date]
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ATTRIBUTE_DEFINITION_DATEs_referenceOrder[attribute_definition_date]
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATIONMap_Staged_Order[attribute_definition_enumeration]; ok {
		return order
	}
	return stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_referenceOrder[attribute_definition_enumeration]
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_referenceOrder[attribute_definition_enumeration]
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_INTEGERMap_Staged_Order[attribute_definition_integer]; ok {
		return order
	}
	return stage.ATTRIBUTE_DEFINITION_INTEGERs_referenceOrder[attribute_definition_integer]
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ATTRIBUTE_DEFINITION_INTEGERs_referenceOrder[attribute_definition_integer]
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_REALMap_Staged_Order[attribute_definition_real]; ok {
		return order
	}
	return stage.ATTRIBUTE_DEFINITION_REALs_referenceOrder[attribute_definition_real]
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ATTRIBUTE_DEFINITION_REALs_referenceOrder[attribute_definition_real]
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_STRINGMap_Staged_Order[attribute_definition_string]; ok {
		return order
	}
	return stage.ATTRIBUTE_DEFINITION_STRINGs_referenceOrder[attribute_definition_string]
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ATTRIBUTE_DEFINITION_STRINGs_referenceOrder[attribute_definition_string]
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_XHTMLMap_Staged_Order[attribute_definition_xhtml]; ok {
		return order
	}
	return stage.ATTRIBUTE_DEFINITION_XHTMLs_referenceOrder[attribute_definition_xhtml]
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ATTRIBUTE_DEFINITION_XHTMLs_referenceOrder[attribute_definition_xhtml]
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[attribute_value_boolean]; ok {
		return order
	}
	return stage.ATTRIBUTE_VALUE_BOOLEANs_referenceOrder[attribute_value_boolean]
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ATTRIBUTE_VALUE_BOOLEANs_referenceOrder[attribute_value_boolean]
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_VALUE_DATEMap_Staged_Order[attribute_value_date]; ok {
		return order
	}
	return stage.ATTRIBUTE_VALUE_DATEs_referenceOrder[attribute_value_date]
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ATTRIBUTE_VALUE_DATEs_referenceOrder[attribute_value_date]
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[attribute_value_enumeration]; ok {
		return order
	}
	return stage.ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder[attribute_value_enumeration]
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder[attribute_value_enumeration]
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[attribute_value_integer]; ok {
		return order
	}
	return stage.ATTRIBUTE_VALUE_INTEGERs_referenceOrder[attribute_value_integer]
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ATTRIBUTE_VALUE_INTEGERs_referenceOrder[attribute_value_integer]
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_VALUE_REALMap_Staged_Order[attribute_value_real]; ok {
		return order
	}
	return stage.ATTRIBUTE_VALUE_REALs_referenceOrder[attribute_value_real]
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ATTRIBUTE_VALUE_REALs_referenceOrder[attribute_value_real]
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_VALUE_STRINGMap_Staged_Order[attribute_value_string]; ok {
		return order
	}
	return stage.ATTRIBUTE_VALUE_STRINGs_referenceOrder[attribute_value_string]
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ATTRIBUTE_VALUE_STRINGs_referenceOrder[attribute_value_string]
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[attribute_value_xhtml]; ok {
		return order
	}
	return stage.ATTRIBUTE_VALUE_XHTMLs_referenceOrder[attribute_value_xhtml]
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ATTRIBUTE_VALUE_XHTMLs_referenceOrder[attribute_value_xhtml]
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ALTERNATIVE_IDMap_Staged_Order[a_alternative_id]; ok {
		return order
	}
	return stage.A_ALTERNATIVE_IDs_referenceOrder[a_alternative_id]
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ALTERNATIVE_IDs_referenceOrder[a_alternative_id]
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_attribute_definition_boolean_ref]; ok {
		return order
	}
	return stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_referenceOrder[a_attribute_definition_boolean_ref]
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_referenceOrder[a_attribute_definition_boolean_ref]
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_DATE_REFMap_Staged_Order[a_attribute_definition_date_ref]; ok {
		return order
	}
	return stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_referenceOrder[a_attribute_definition_date_ref]
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_referenceOrder[a_attribute_definition_date_ref]
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_attribute_definition_enumeration_ref]; ok {
		return order
	}
	return stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_referenceOrder[a_attribute_definition_enumeration_ref]
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_referenceOrder[a_attribute_definition_enumeration_ref]
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFMap_Staged_Order[a_attribute_definition_integer_ref]; ok {
		return order
	}
	return stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_referenceOrder[a_attribute_definition_integer_ref]
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_referenceOrder[a_attribute_definition_integer_ref]
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_REAL_REFMap_Staged_Order[a_attribute_definition_real_ref]; ok {
		return order
	}
	return stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_referenceOrder[a_attribute_definition_real_ref]
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_referenceOrder[a_attribute_definition_real_ref]
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_STRING_REFMap_Staged_Order[a_attribute_definition_string_ref]; ok {
		return order
	}
	return stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_referenceOrder[a_attribute_definition_string_ref]
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_referenceOrder[a_attribute_definition_string_ref]
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_XHTML_REFMap_Staged_Order[a_attribute_definition_xhtml_ref]; ok {
		return order
	}
	return stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_referenceOrder[a_attribute_definition_xhtml_ref]
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_referenceOrder[a_attribute_definition_xhtml_ref]
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_VALUE_BOOLEANMap_Staged_Order[a_attribute_value_boolean]; ok {
		return order
	}
	return stage.A_ATTRIBUTE_VALUE_BOOLEANs_referenceOrder[a_attribute_value_boolean]
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ATTRIBUTE_VALUE_BOOLEANs_referenceOrder[a_attribute_value_boolean]
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_VALUE_DATEMap_Staged_Order[a_attribute_value_date]; ok {
		return order
	}
	return stage.A_ATTRIBUTE_VALUE_DATEs_referenceOrder[a_attribute_value_date]
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ATTRIBUTE_VALUE_DATEs_referenceOrder[a_attribute_value_date]
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_VALUE_ENUMERATIONMap_Staged_Order[a_attribute_value_enumeration]; ok {
		return order
	}
	return stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder[a_attribute_value_enumeration]
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder[a_attribute_value_enumeration]
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_VALUE_INTEGERMap_Staged_Order[a_attribute_value_integer]; ok {
		return order
	}
	return stage.A_ATTRIBUTE_VALUE_INTEGERs_referenceOrder[a_attribute_value_integer]
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ATTRIBUTE_VALUE_INTEGERs_referenceOrder[a_attribute_value_integer]
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_VALUE_REALMap_Staged_Order[a_attribute_value_real]; ok {
		return order
	}
	return stage.A_ATTRIBUTE_VALUE_REALs_referenceOrder[a_attribute_value_real]
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ATTRIBUTE_VALUE_REALs_referenceOrder[a_attribute_value_real]
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_VALUE_STRINGMap_Staged_Order[a_attribute_value_string]; ok {
		return order
	}
	return stage.A_ATTRIBUTE_VALUE_STRINGs_referenceOrder[a_attribute_value_string]
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ATTRIBUTE_VALUE_STRINGs_referenceOrder[a_attribute_value_string]
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_VALUE_XHTMLMap_Staged_Order[a_attribute_value_xhtml]; ok {
		return order
	}
	return stage.A_ATTRIBUTE_VALUE_XHTMLs_referenceOrder[a_attribute_value_xhtml]
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ATTRIBUTE_VALUE_XHTMLs_referenceOrder[a_attribute_value_xhtml]
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1Map_Staged_Order[a_attribute_value_xhtml_1]; ok {
		return order
	}
	return stage.A_ATTRIBUTE_VALUE_XHTML_1s_referenceOrder[a_attribute_value_xhtml_1]
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ATTRIBUTE_VALUE_XHTML_1s_referenceOrder[a_attribute_value_xhtml_1]
}

func (a_children *A_CHILDREN) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_CHILDRENMap_Staged_Order[a_children]; ok {
		return order
	}
	return stage.A_CHILDRENs_referenceOrder[a_children]
}

func (a_children *A_CHILDREN) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_CHILDRENs_referenceOrder[a_children]
}

func (a_core_content *A_CORE_CONTENT) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_CORE_CONTENTMap_Staged_Order[a_core_content]; ok {
		return order
	}
	return stage.A_CORE_CONTENTs_referenceOrder[a_core_content]
}

func (a_core_content *A_CORE_CONTENT) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_CORE_CONTENTs_referenceOrder[a_core_content]
}

func (a_datatypes *A_DATATYPES) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_DATATYPESMap_Staged_Order[a_datatypes]; ok {
		return order
	}
	return stage.A_DATATYPESs_referenceOrder[a_datatypes]
}

func (a_datatypes *A_DATATYPES) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_DATATYPESs_referenceOrder[a_datatypes]
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_DATATYPE_DEFINITION_BOOLEAN_REFMap_Staged_Order[a_datatype_definition_boolean_ref]; ok {
		return order
	}
	return stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_referenceOrder[a_datatype_definition_boolean_ref]
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_referenceOrder[a_datatype_definition_boolean_ref]
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_DATATYPE_DEFINITION_DATE_REFMap_Staged_Order[a_datatype_definition_date_ref]; ok {
		return order
	}
	return stage.A_DATATYPE_DEFINITION_DATE_REFs_referenceOrder[a_datatype_definition_date_ref]
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_DATATYPE_DEFINITION_DATE_REFs_referenceOrder[a_datatype_definition_date_ref]
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_DATATYPE_DEFINITION_ENUMERATION_REFMap_Staged_Order[a_datatype_definition_enumeration_ref]; ok {
		return order
	}
	return stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_referenceOrder[a_datatype_definition_enumeration_ref]
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_referenceOrder[a_datatype_definition_enumeration_ref]
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_DATATYPE_DEFINITION_INTEGER_REFMap_Staged_Order[a_datatype_definition_integer_ref]; ok {
		return order
	}
	return stage.A_DATATYPE_DEFINITION_INTEGER_REFs_referenceOrder[a_datatype_definition_integer_ref]
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_DATATYPE_DEFINITION_INTEGER_REFs_referenceOrder[a_datatype_definition_integer_ref]
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_DATATYPE_DEFINITION_REAL_REFMap_Staged_Order[a_datatype_definition_real_ref]; ok {
		return order
	}
	return stage.A_DATATYPE_DEFINITION_REAL_REFs_referenceOrder[a_datatype_definition_real_ref]
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_DATATYPE_DEFINITION_REAL_REFs_referenceOrder[a_datatype_definition_real_ref]
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_DATATYPE_DEFINITION_STRING_REFMap_Staged_Order[a_datatype_definition_string_ref]; ok {
		return order
	}
	return stage.A_DATATYPE_DEFINITION_STRING_REFs_referenceOrder[a_datatype_definition_string_ref]
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_DATATYPE_DEFINITION_STRING_REFs_referenceOrder[a_datatype_definition_string_ref]
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_DATATYPE_DEFINITION_XHTML_REFMap_Staged_Order[a_datatype_definition_xhtml_ref]; ok {
		return order
	}
	return stage.A_DATATYPE_DEFINITION_XHTML_REFs_referenceOrder[a_datatype_definition_xhtml_ref]
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_DATATYPE_DEFINITION_XHTML_REFs_referenceOrder[a_datatype_definition_xhtml_ref]
}

func (a_editable_atts *A_EDITABLE_ATTS) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_EDITABLE_ATTSMap_Staged_Order[a_editable_atts]; ok {
		return order
	}
	return stage.A_EDITABLE_ATTSs_referenceOrder[a_editable_atts]
}

func (a_editable_atts *A_EDITABLE_ATTS) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_EDITABLE_ATTSs_referenceOrder[a_editable_atts]
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ENUM_VALUE_REFMap_Staged_Order[a_enum_value_ref]; ok {
		return order
	}
	return stage.A_ENUM_VALUE_REFs_referenceOrder[a_enum_value_ref]
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_ENUM_VALUE_REFs_referenceOrder[a_enum_value_ref]
}

func (a_object *A_OBJECT) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_OBJECTMap_Staged_Order[a_object]; ok {
		return order
	}
	return stage.A_OBJECTs_referenceOrder[a_object]
}

func (a_object *A_OBJECT) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_OBJECTs_referenceOrder[a_object]
}

func (a_properties *A_PROPERTIES) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_PROPERTIESMap_Staged_Order[a_properties]; ok {
		return order
	}
	return stage.A_PROPERTIESs_referenceOrder[a_properties]
}

func (a_properties *A_PROPERTIES) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_PROPERTIESs_referenceOrder[a_properties]
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_RELATION_GROUP_TYPE_REFMap_Staged_Order[a_relation_group_type_ref]; ok {
		return order
	}
	return stage.A_RELATION_GROUP_TYPE_REFs_referenceOrder[a_relation_group_type_ref]
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_RELATION_GROUP_TYPE_REFs_referenceOrder[a_relation_group_type_ref]
}

func (a_source_1 *A_SOURCE_1) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SOURCE_1Map_Staged_Order[a_source_1]; ok {
		return order
	}
	return stage.A_SOURCE_1s_referenceOrder[a_source_1]
}

func (a_source_1 *A_SOURCE_1) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_SOURCE_1s_referenceOrder[a_source_1]
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SOURCE_SPECIFICATION_1Map_Staged_Order[a_source_specification_1]; ok {
		return order
	}
	return stage.A_SOURCE_SPECIFICATION_1s_referenceOrder[a_source_specification_1]
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_SOURCE_SPECIFICATION_1s_referenceOrder[a_source_specification_1]
}

func (a_specifications *A_SPECIFICATIONS) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPECIFICATIONSMap_Staged_Order[a_specifications]; ok {
		return order
	}
	return stage.A_SPECIFICATIONSs_referenceOrder[a_specifications]
}

func (a_specifications *A_SPECIFICATIONS) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_SPECIFICATIONSs_referenceOrder[a_specifications]
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPECIFICATION_TYPE_REFMap_Staged_Order[a_specification_type_ref]; ok {
		return order
	}
	return stage.A_SPECIFICATION_TYPE_REFs_referenceOrder[a_specification_type_ref]
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_SPECIFICATION_TYPE_REFs_referenceOrder[a_specification_type_ref]
}

func (a_specified_values *A_SPECIFIED_VALUES) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPECIFIED_VALUESMap_Staged_Order[a_specified_values]; ok {
		return order
	}
	return stage.A_SPECIFIED_VALUESs_referenceOrder[a_specified_values]
}

func (a_specified_values *A_SPECIFIED_VALUES) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_SPECIFIED_VALUESs_referenceOrder[a_specified_values]
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPEC_ATTRIBUTESMap_Staged_Order[a_spec_attributes]; ok {
		return order
	}
	return stage.A_SPEC_ATTRIBUTESs_referenceOrder[a_spec_attributes]
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_SPEC_ATTRIBUTESs_referenceOrder[a_spec_attributes]
}

func (a_spec_objects *A_SPEC_OBJECTS) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPEC_OBJECTSMap_Staged_Order[a_spec_objects]; ok {
		return order
	}
	return stage.A_SPEC_OBJECTSs_referenceOrder[a_spec_objects]
}

func (a_spec_objects *A_SPEC_OBJECTS) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_SPEC_OBJECTSs_referenceOrder[a_spec_objects]
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPEC_OBJECT_TYPE_REFMap_Staged_Order[a_spec_object_type_ref]; ok {
		return order
	}
	return stage.A_SPEC_OBJECT_TYPE_REFs_referenceOrder[a_spec_object_type_ref]
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_SPEC_OBJECT_TYPE_REFs_referenceOrder[a_spec_object_type_ref]
}

func (a_spec_relations *A_SPEC_RELATIONS) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPEC_RELATIONSMap_Staged_Order[a_spec_relations]; ok {
		return order
	}
	return stage.A_SPEC_RELATIONSs_referenceOrder[a_spec_relations]
}

func (a_spec_relations *A_SPEC_RELATIONS) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_SPEC_RELATIONSs_referenceOrder[a_spec_relations]
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPEC_RELATION_GROUPSMap_Staged_Order[a_spec_relation_groups]; ok {
		return order
	}
	return stage.A_SPEC_RELATION_GROUPSs_referenceOrder[a_spec_relation_groups]
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_SPEC_RELATION_GROUPSs_referenceOrder[a_spec_relation_groups]
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPEC_RELATION_REFMap_Staged_Order[a_spec_relation_ref]; ok {
		return order
	}
	return stage.A_SPEC_RELATION_REFs_referenceOrder[a_spec_relation_ref]
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_SPEC_RELATION_REFs_referenceOrder[a_spec_relation_ref]
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPEC_RELATION_TYPE_REFMap_Staged_Order[a_spec_relation_type_ref]; ok {
		return order
	}
	return stage.A_SPEC_RELATION_TYPE_REFs_referenceOrder[a_spec_relation_type_ref]
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_SPEC_RELATION_TYPE_REFs_referenceOrder[a_spec_relation_type_ref]
}

func (a_spec_types *A_SPEC_TYPES) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPEC_TYPESMap_Staged_Order[a_spec_types]; ok {
		return order
	}
	return stage.A_SPEC_TYPESs_referenceOrder[a_spec_types]
}

func (a_spec_types *A_SPEC_TYPES) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_SPEC_TYPESs_referenceOrder[a_spec_types]
}

func (a_the_header *A_THE_HEADER) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_THE_HEADERMap_Staged_Order[a_the_header]; ok {
		return order
	}
	return stage.A_THE_HEADERs_referenceOrder[a_the_header]
}

func (a_the_header *A_THE_HEADER) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_THE_HEADERs_referenceOrder[a_the_header]
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_TOOL_EXTENSIONSMap_Staged_Order[a_tool_extensions]; ok {
		return order
	}
	return stage.A_TOOL_EXTENSIONSs_referenceOrder[a_tool_extensions]
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongGetReferenceOrder(stage *Stage) uint {
	return stage.A_TOOL_EXTENSIONSs_referenceOrder[a_tool_extensions]
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DATATYPE_DEFINITION_BOOLEANMap_Staged_Order[datatype_definition_boolean]; ok {
		return order
	}
	return stage.DATATYPE_DEFINITION_BOOLEANs_referenceOrder[datatype_definition_boolean]
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongGetReferenceOrder(stage *Stage) uint {
	return stage.DATATYPE_DEFINITION_BOOLEANs_referenceOrder[datatype_definition_boolean]
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DATATYPE_DEFINITION_DATEMap_Staged_Order[datatype_definition_date]; ok {
		return order
	}
	return stage.DATATYPE_DEFINITION_DATEs_referenceOrder[datatype_definition_date]
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongGetReferenceOrder(stage *Stage) uint {
	return stage.DATATYPE_DEFINITION_DATEs_referenceOrder[datatype_definition_date]
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DATATYPE_DEFINITION_ENUMERATIONMap_Staged_Order[datatype_definition_enumeration]; ok {
		return order
	}
	return stage.DATATYPE_DEFINITION_ENUMERATIONs_referenceOrder[datatype_definition_enumeration]
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongGetReferenceOrder(stage *Stage) uint {
	return stage.DATATYPE_DEFINITION_ENUMERATIONs_referenceOrder[datatype_definition_enumeration]
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DATATYPE_DEFINITION_INTEGERMap_Staged_Order[datatype_definition_integer]; ok {
		return order
	}
	return stage.DATATYPE_DEFINITION_INTEGERs_referenceOrder[datatype_definition_integer]
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongGetReferenceOrder(stage *Stage) uint {
	return stage.DATATYPE_DEFINITION_INTEGERs_referenceOrder[datatype_definition_integer]
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DATATYPE_DEFINITION_REALMap_Staged_Order[datatype_definition_real]; ok {
		return order
	}
	return stage.DATATYPE_DEFINITION_REALs_referenceOrder[datatype_definition_real]
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongGetReferenceOrder(stage *Stage) uint {
	return stage.DATATYPE_DEFINITION_REALs_referenceOrder[datatype_definition_real]
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DATATYPE_DEFINITION_STRINGMap_Staged_Order[datatype_definition_string]; ok {
		return order
	}
	return stage.DATATYPE_DEFINITION_STRINGs_referenceOrder[datatype_definition_string]
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongGetReferenceOrder(stage *Stage) uint {
	return stage.DATATYPE_DEFINITION_STRINGs_referenceOrder[datatype_definition_string]
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DATATYPE_DEFINITION_XHTMLMap_Staged_Order[datatype_definition_xhtml]; ok {
		return order
	}
	return stage.DATATYPE_DEFINITION_XHTMLs_referenceOrder[datatype_definition_xhtml]
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongGetReferenceOrder(stage *Stage) uint {
	return stage.DATATYPE_DEFINITION_XHTMLs_referenceOrder[datatype_definition_xhtml]
}

func (embedded_value *EMBEDDED_VALUE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EMBEDDED_VALUEMap_Staged_Order[embedded_value]; ok {
		return order
	}
	return stage.EMBEDDED_VALUEs_referenceOrder[embedded_value]
}

func (embedded_value *EMBEDDED_VALUE) GongGetReferenceOrder(stage *Stage) uint {
	return stage.EMBEDDED_VALUEs_referenceOrder[embedded_value]
}

func (enum_value *ENUM_VALUE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ENUM_VALUEMap_Staged_Order[enum_value]; ok {
		return order
	}
	return stage.ENUM_VALUEs_referenceOrder[enum_value]
}

func (enum_value *ENUM_VALUE) GongGetReferenceOrder(stage *Stage) uint {
	return stage.ENUM_VALUEs_referenceOrder[enum_value]
}

func (relation_group *RELATION_GROUP) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RELATION_GROUPMap_Staged_Order[relation_group]; ok {
		return order
	}
	return stage.RELATION_GROUPs_referenceOrder[relation_group]
}

func (relation_group *RELATION_GROUP) GongGetReferenceOrder(stage *Stage) uint {
	return stage.RELATION_GROUPs_referenceOrder[relation_group]
}

func (relation_group_type *RELATION_GROUP_TYPE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RELATION_GROUP_TYPEMap_Staged_Order[relation_group_type]; ok {
		return order
	}
	return stage.RELATION_GROUP_TYPEs_referenceOrder[relation_group_type]
}

func (relation_group_type *RELATION_GROUP_TYPE) GongGetReferenceOrder(stage *Stage) uint {
	return stage.RELATION_GROUP_TYPEs_referenceOrder[relation_group_type]
}

func (req_if *REQ_IF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.REQ_IFMap_Staged_Order[req_if]; ok {
		return order
	}
	return stage.REQ_IFs_referenceOrder[req_if]
}

func (req_if *REQ_IF) GongGetReferenceOrder(stage *Stage) uint {
	return stage.REQ_IFs_referenceOrder[req_if]
}

func (req_if_content *REQ_IF_CONTENT) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.REQ_IF_CONTENTMap_Staged_Order[req_if_content]; ok {
		return order
	}
	return stage.REQ_IF_CONTENTs_referenceOrder[req_if_content]
}

func (req_if_content *REQ_IF_CONTENT) GongGetReferenceOrder(stage *Stage) uint {
	return stage.REQ_IF_CONTENTs_referenceOrder[req_if_content]
}

func (req_if_header *REQ_IF_HEADER) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.REQ_IF_HEADERMap_Staged_Order[req_if_header]; ok {
		return order
	}
	return stage.REQ_IF_HEADERs_referenceOrder[req_if_header]
}

func (req_if_header *REQ_IF_HEADER) GongGetReferenceOrder(stage *Stage) uint {
	return stage.REQ_IF_HEADERs_referenceOrder[req_if_header]
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.REQ_IF_TOOL_EXTENSIONMap_Staged_Order[req_if_tool_extension]; ok {
		return order
	}
	return stage.REQ_IF_TOOL_EXTENSIONs_referenceOrder[req_if_tool_extension]
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongGetReferenceOrder(stage *Stage) uint {
	return stage.REQ_IF_TOOL_EXTENSIONs_referenceOrder[req_if_tool_extension]
}

func (specification *SPECIFICATION) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SPECIFICATIONMap_Staged_Order[specification]; ok {
		return order
	}
	return stage.SPECIFICATIONs_referenceOrder[specification]
}

func (specification *SPECIFICATION) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SPECIFICATIONs_referenceOrder[specification]
}

func (specification_type *SPECIFICATION_TYPE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SPECIFICATION_TYPEMap_Staged_Order[specification_type]; ok {
		return order
	}
	return stage.SPECIFICATION_TYPEs_referenceOrder[specification_type]
}

func (specification_type *SPECIFICATION_TYPE) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SPECIFICATION_TYPEs_referenceOrder[specification_type]
}

func (spec_hierarchy *SPEC_HIERARCHY) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SPEC_HIERARCHYMap_Staged_Order[spec_hierarchy]; ok {
		return order
	}
	return stage.SPEC_HIERARCHYs_referenceOrder[spec_hierarchy]
}

func (spec_hierarchy *SPEC_HIERARCHY) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SPEC_HIERARCHYs_referenceOrder[spec_hierarchy]
}

func (spec_object *SPEC_OBJECT) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SPEC_OBJECTMap_Staged_Order[spec_object]; ok {
		return order
	}
	return stage.SPEC_OBJECTs_referenceOrder[spec_object]
}

func (spec_object *SPEC_OBJECT) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SPEC_OBJECTs_referenceOrder[spec_object]
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SPEC_OBJECT_TYPEMap_Staged_Order[spec_object_type]; ok {
		return order
	}
	return stage.SPEC_OBJECT_TYPEs_referenceOrder[spec_object_type]
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SPEC_OBJECT_TYPEs_referenceOrder[spec_object_type]
}

func (spec_relation *SPEC_RELATION) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SPEC_RELATIONMap_Staged_Order[spec_relation]; ok {
		return order
	}
	return stage.SPEC_RELATIONs_referenceOrder[spec_relation]
}

func (spec_relation *SPEC_RELATION) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SPEC_RELATIONs_referenceOrder[spec_relation]
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SPEC_RELATION_TYPEMap_Staged_Order[spec_relation_type]; ok {
		return order
	}
	return stage.SPEC_RELATION_TYPEs_referenceOrder[spec_relation_type]
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongGetReferenceOrder(stage *Stage) uint {
	return stage.SPEC_RELATION_TYPEs_referenceOrder[spec_relation_type]
}

func (xhtml_content *XHTML_CONTENT) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.XHTML_CONTENTMap_Staged_Order[xhtml_content]; ok {
		return order
	}
	return stage.XHTML_CONTENTs_referenceOrder[xhtml_content]
}

func (xhtml_content *XHTML_CONTENT) GongGetReferenceOrder(stage *Stage) uint {
	return stage.XHTML_CONTENTs_referenceOrder[xhtml_content]
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (alternative_id *ALTERNATIVE_ID) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", alternative_id.GongGetGongstructName(), alternative_id.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (alternative_id *ALTERNATIVE_ID) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", alternative_id.GongGetGongstructName(), alternative_id.GongGetReferenceOrder(stage))
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_boolean.GongGetGongstructName(), attribute_definition_boolean.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_boolean.GongGetGongstructName(), attribute_definition_boolean.GongGetReferenceOrder(stage))
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_date.GongGetGongstructName(), attribute_definition_date.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_date.GongGetGongstructName(), attribute_definition_date.GongGetReferenceOrder(stage))
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_enumeration.GongGetGongstructName(), attribute_definition_enumeration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_enumeration.GongGetGongstructName(), attribute_definition_enumeration.GongGetReferenceOrder(stage))
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_integer.GongGetGongstructName(), attribute_definition_integer.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_integer.GongGetGongstructName(), attribute_definition_integer.GongGetReferenceOrder(stage))
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_real.GongGetGongstructName(), attribute_definition_real.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_real.GongGetGongstructName(), attribute_definition_real.GongGetReferenceOrder(stage))
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_string.GongGetGongstructName(), attribute_definition_string.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_string.GongGetGongstructName(), attribute_definition_string.GongGetReferenceOrder(stage))
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_xhtml.GongGetGongstructName(), attribute_definition_xhtml.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_xhtml.GongGetGongstructName(), attribute_definition_xhtml.GongGetReferenceOrder(stage))
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_boolean.GongGetGongstructName(), attribute_value_boolean.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_boolean.GongGetGongstructName(), attribute_value_boolean.GongGetReferenceOrder(stage))
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_date.GongGetGongstructName(), attribute_value_date.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_date.GongGetGongstructName(), attribute_value_date.GongGetReferenceOrder(stage))
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_enumeration.GongGetGongstructName(), attribute_value_enumeration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_enumeration.GongGetGongstructName(), attribute_value_enumeration.GongGetReferenceOrder(stage))
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_integer.GongGetGongstructName(), attribute_value_integer.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_integer.GongGetGongstructName(), attribute_value_integer.GongGetReferenceOrder(stage))
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_real.GongGetGongstructName(), attribute_value_real.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_real.GongGetGongstructName(), attribute_value_real.GongGetReferenceOrder(stage))
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_string.GongGetGongstructName(), attribute_value_string.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_string.GongGetGongstructName(), attribute_value_string.GongGetReferenceOrder(stage))
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_xhtml.GongGetGongstructName(), attribute_value_xhtml.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_xhtml.GongGetGongstructName(), attribute_value_xhtml.GongGetReferenceOrder(stage))
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_alternative_id.GongGetGongstructName(), a_alternative_id.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_alternative_id *A_ALTERNATIVE_ID) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_alternative_id.GongGetGongstructName(), a_alternative_id.GongGetReferenceOrder(stage))
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_boolean_ref.GongGetGongstructName(), a_attribute_definition_boolean_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_boolean_ref.GongGetGongstructName(), a_attribute_definition_boolean_ref.GongGetReferenceOrder(stage))
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_date_ref.GongGetGongstructName(), a_attribute_definition_date_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_date_ref.GongGetGongstructName(), a_attribute_definition_date_ref.GongGetReferenceOrder(stage))
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_enumeration_ref.GongGetGongstructName(), a_attribute_definition_enumeration_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_enumeration_ref.GongGetGongstructName(), a_attribute_definition_enumeration_ref.GongGetReferenceOrder(stage))
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_integer_ref.GongGetGongstructName(), a_attribute_definition_integer_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_integer_ref.GongGetGongstructName(), a_attribute_definition_integer_ref.GongGetReferenceOrder(stage))
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_real_ref.GongGetGongstructName(), a_attribute_definition_real_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_real_ref.GongGetGongstructName(), a_attribute_definition_real_ref.GongGetReferenceOrder(stage))
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_string_ref.GongGetGongstructName(), a_attribute_definition_string_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_string_ref.GongGetGongstructName(), a_attribute_definition_string_ref.GongGetReferenceOrder(stage))
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_xhtml_ref.GongGetGongstructName(), a_attribute_definition_xhtml_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_xhtml_ref.GongGetGongstructName(), a_attribute_definition_xhtml_ref.GongGetReferenceOrder(stage))
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_boolean.GongGetGongstructName(), a_attribute_value_boolean.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_boolean.GongGetGongstructName(), a_attribute_value_boolean.GongGetReferenceOrder(stage))
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_date.GongGetGongstructName(), a_attribute_value_date.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_date.GongGetGongstructName(), a_attribute_value_date.GongGetReferenceOrder(stage))
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_enumeration.GongGetGongstructName(), a_attribute_value_enumeration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_enumeration.GongGetGongstructName(), a_attribute_value_enumeration.GongGetReferenceOrder(stage))
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_integer.GongGetGongstructName(), a_attribute_value_integer.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_integer.GongGetGongstructName(), a_attribute_value_integer.GongGetReferenceOrder(stage))
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_real.GongGetGongstructName(), a_attribute_value_real.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_real.GongGetGongstructName(), a_attribute_value_real.GongGetReferenceOrder(stage))
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_string.GongGetGongstructName(), a_attribute_value_string.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_string.GongGetGongstructName(), a_attribute_value_string.GongGetReferenceOrder(stage))
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_xhtml.GongGetGongstructName(), a_attribute_value_xhtml.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_xhtml.GongGetGongstructName(), a_attribute_value_xhtml.GongGetReferenceOrder(stage))
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_xhtml_1.GongGetGongstructName(), a_attribute_value_xhtml_1.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_xhtml_1.GongGetGongstructName(), a_attribute_value_xhtml_1.GongGetReferenceOrder(stage))
}

func (a_children *A_CHILDREN) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_children.GongGetGongstructName(), a_children.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_children *A_CHILDREN) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_children.GongGetGongstructName(), a_children.GongGetReferenceOrder(stage))
}

func (a_core_content *A_CORE_CONTENT) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_core_content.GongGetGongstructName(), a_core_content.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_core_content *A_CORE_CONTENT) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_core_content.GongGetGongstructName(), a_core_content.GongGetReferenceOrder(stage))
}

func (a_datatypes *A_DATATYPES) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatypes.GongGetGongstructName(), a_datatypes.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatypes *A_DATATYPES) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatypes.GongGetGongstructName(), a_datatypes.GongGetReferenceOrder(stage))
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_boolean_ref.GongGetGongstructName(), a_datatype_definition_boolean_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_boolean_ref.GongGetGongstructName(), a_datatype_definition_boolean_ref.GongGetReferenceOrder(stage))
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_date_ref.GongGetGongstructName(), a_datatype_definition_date_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_date_ref.GongGetGongstructName(), a_datatype_definition_date_ref.GongGetReferenceOrder(stage))
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_enumeration_ref.GongGetGongstructName(), a_datatype_definition_enumeration_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_enumeration_ref.GongGetGongstructName(), a_datatype_definition_enumeration_ref.GongGetReferenceOrder(stage))
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_integer_ref.GongGetGongstructName(), a_datatype_definition_integer_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_integer_ref.GongGetGongstructName(), a_datatype_definition_integer_ref.GongGetReferenceOrder(stage))
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_real_ref.GongGetGongstructName(), a_datatype_definition_real_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_real_ref.GongGetGongstructName(), a_datatype_definition_real_ref.GongGetReferenceOrder(stage))
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_string_ref.GongGetGongstructName(), a_datatype_definition_string_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_string_ref.GongGetGongstructName(), a_datatype_definition_string_ref.GongGetReferenceOrder(stage))
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_xhtml_ref.GongGetGongstructName(), a_datatype_definition_xhtml_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_xhtml_ref.GongGetGongstructName(), a_datatype_definition_xhtml_ref.GongGetReferenceOrder(stage))
}

func (a_editable_atts *A_EDITABLE_ATTS) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_editable_atts.GongGetGongstructName(), a_editable_atts.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_editable_atts *A_EDITABLE_ATTS) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_editable_atts.GongGetGongstructName(), a_editable_atts.GongGetReferenceOrder(stage))
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_enum_value_ref.GongGetGongstructName(), a_enum_value_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_enum_value_ref *A_ENUM_VALUE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_enum_value_ref.GongGetGongstructName(), a_enum_value_ref.GongGetReferenceOrder(stage))
}

func (a_object *A_OBJECT) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_object.GongGetGongstructName(), a_object.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_object *A_OBJECT) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_object.GongGetGongstructName(), a_object.GongGetReferenceOrder(stage))
}

func (a_properties *A_PROPERTIES) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_properties.GongGetGongstructName(), a_properties.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_properties *A_PROPERTIES) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_properties.GongGetGongstructName(), a_properties.GongGetReferenceOrder(stage))
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_relation_group_type_ref.GongGetGongstructName(), a_relation_group_type_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_relation_group_type_ref.GongGetGongstructName(), a_relation_group_type_ref.GongGetReferenceOrder(stage))
}

func (a_source_1 *A_SOURCE_1) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_source_1.GongGetGongstructName(), a_source_1.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_source_1 *A_SOURCE_1) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_source_1.GongGetGongstructName(), a_source_1.GongGetReferenceOrder(stage))
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_source_specification_1.GongGetGongstructName(), a_source_specification_1.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_source_specification_1.GongGetGongstructName(), a_source_specification_1.GongGetReferenceOrder(stage))
}

func (a_specifications *A_SPECIFICATIONS) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_specifications.GongGetGongstructName(), a_specifications.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_specifications *A_SPECIFICATIONS) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_specifications.GongGetGongstructName(), a_specifications.GongGetReferenceOrder(stage))
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_specification_type_ref.GongGetGongstructName(), a_specification_type_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_specification_type_ref.GongGetGongstructName(), a_specification_type_ref.GongGetReferenceOrder(stage))
}

func (a_specified_values *A_SPECIFIED_VALUES) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_specified_values.GongGetGongstructName(), a_specified_values.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_specified_values *A_SPECIFIED_VALUES) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_specified_values.GongGetGongstructName(), a_specified_values.GongGetReferenceOrder(stage))
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_attributes.GongGetGongstructName(), a_spec_attributes.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_attributes.GongGetGongstructName(), a_spec_attributes.GongGetReferenceOrder(stage))
}

func (a_spec_objects *A_SPEC_OBJECTS) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_objects.GongGetGongstructName(), a_spec_objects.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_objects *A_SPEC_OBJECTS) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_objects.GongGetGongstructName(), a_spec_objects.GongGetReferenceOrder(stage))
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_object_type_ref.GongGetGongstructName(), a_spec_object_type_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_object_type_ref.GongGetGongstructName(), a_spec_object_type_ref.GongGetReferenceOrder(stage))
}

func (a_spec_relations *A_SPEC_RELATIONS) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_relations.GongGetGongstructName(), a_spec_relations.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_relations *A_SPEC_RELATIONS) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_relations.GongGetGongstructName(), a_spec_relations.GongGetReferenceOrder(stage))
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_relation_groups.GongGetGongstructName(), a_spec_relation_groups.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_relation_groups.GongGetGongstructName(), a_spec_relation_groups.GongGetReferenceOrder(stage))
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_relation_ref.GongGetGongstructName(), a_spec_relation_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_relation_ref.GongGetGongstructName(), a_spec_relation_ref.GongGetReferenceOrder(stage))
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_relation_type_ref.GongGetGongstructName(), a_spec_relation_type_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_relation_type_ref.GongGetGongstructName(), a_spec_relation_type_ref.GongGetReferenceOrder(stage))
}

func (a_spec_types *A_SPEC_TYPES) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_types.GongGetGongstructName(), a_spec_types.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_types *A_SPEC_TYPES) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_types.GongGetGongstructName(), a_spec_types.GongGetReferenceOrder(stage))
}

func (a_the_header *A_THE_HEADER) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_the_header.GongGetGongstructName(), a_the_header.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_the_header *A_THE_HEADER) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_the_header.GongGetGongstructName(), a_the_header.GongGetReferenceOrder(stage))
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_tool_extensions.GongGetGongstructName(), a_tool_extensions.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_tool_extensions *A_TOOL_EXTENSIONS) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_tool_extensions.GongGetGongstructName(), a_tool_extensions.GongGetReferenceOrder(stage))
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_boolean.GongGetGongstructName(), datatype_definition_boolean.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_boolean.GongGetGongstructName(), datatype_definition_boolean.GongGetReferenceOrder(stage))
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_date.GongGetGongstructName(), datatype_definition_date.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_date.GongGetGongstructName(), datatype_definition_date.GongGetReferenceOrder(stage))
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_enumeration.GongGetGongstructName(), datatype_definition_enumeration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_enumeration.GongGetGongstructName(), datatype_definition_enumeration.GongGetReferenceOrder(stage))
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_integer.GongGetGongstructName(), datatype_definition_integer.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_integer.GongGetGongstructName(), datatype_definition_integer.GongGetReferenceOrder(stage))
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_real.GongGetGongstructName(), datatype_definition_real.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_real.GongGetGongstructName(), datatype_definition_real.GongGetReferenceOrder(stage))
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_string.GongGetGongstructName(), datatype_definition_string.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_string.GongGetGongstructName(), datatype_definition_string.GongGetReferenceOrder(stage))
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_xhtml.GongGetGongstructName(), datatype_definition_xhtml.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_xhtml.GongGetGongstructName(), datatype_definition_xhtml.GongGetReferenceOrder(stage))
}

func (embedded_value *EMBEDDED_VALUE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", embedded_value.GongGetGongstructName(), embedded_value.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (embedded_value *EMBEDDED_VALUE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", embedded_value.GongGetGongstructName(), embedded_value.GongGetReferenceOrder(stage))
}

func (enum_value *ENUM_VALUE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", enum_value.GongGetGongstructName(), enum_value.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (enum_value *ENUM_VALUE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", enum_value.GongGetGongstructName(), enum_value.GongGetReferenceOrder(stage))
}

func (relation_group *RELATION_GROUP) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", relation_group.GongGetGongstructName(), relation_group.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (relation_group *RELATION_GROUP) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", relation_group.GongGetGongstructName(), relation_group.GongGetReferenceOrder(stage))
}

func (relation_group_type *RELATION_GROUP_TYPE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", relation_group_type.GongGetGongstructName(), relation_group_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (relation_group_type *RELATION_GROUP_TYPE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", relation_group_type.GongGetGongstructName(), relation_group_type.GongGetReferenceOrder(stage))
}

func (req_if *REQ_IF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", req_if.GongGetGongstructName(), req_if.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (req_if *REQ_IF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", req_if.GongGetGongstructName(), req_if.GongGetReferenceOrder(stage))
}

func (req_if_content *REQ_IF_CONTENT) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", req_if_content.GongGetGongstructName(), req_if_content.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (req_if_content *REQ_IF_CONTENT) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", req_if_content.GongGetGongstructName(), req_if_content.GongGetReferenceOrder(stage))
}

func (req_if_header *REQ_IF_HEADER) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", req_if_header.GongGetGongstructName(), req_if_header.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (req_if_header *REQ_IF_HEADER) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", req_if_header.GongGetGongstructName(), req_if_header.GongGetReferenceOrder(stage))
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", req_if_tool_extension.GongGetGongstructName(), req_if_tool_extension.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", req_if_tool_extension.GongGetGongstructName(), req_if_tool_extension.GongGetReferenceOrder(stage))
}

func (specification *SPECIFICATION) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", specification.GongGetGongstructName(), specification.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specification *SPECIFICATION) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", specification.GongGetGongstructName(), specification.GongGetReferenceOrder(stage))
}

func (specification_type *SPECIFICATION_TYPE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", specification_type.GongGetGongstructName(), specification_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specification_type *SPECIFICATION_TYPE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", specification_type.GongGetGongstructName(), specification_type.GongGetReferenceOrder(stage))
}

func (spec_hierarchy *SPEC_HIERARCHY) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_hierarchy.GongGetGongstructName(), spec_hierarchy.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spec_hierarchy *SPEC_HIERARCHY) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_hierarchy.GongGetGongstructName(), spec_hierarchy.GongGetReferenceOrder(stage))
}

func (spec_object *SPEC_OBJECT) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_object.GongGetGongstructName(), spec_object.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spec_object *SPEC_OBJECT) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_object.GongGetGongstructName(), spec_object.GongGetReferenceOrder(stage))
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_object_type.GongGetGongstructName(), spec_object_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spec_object_type *SPEC_OBJECT_TYPE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_object_type.GongGetGongstructName(), spec_object_type.GongGetReferenceOrder(stage))
}

func (spec_relation *SPEC_RELATION) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_relation.GongGetGongstructName(), spec_relation.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spec_relation *SPEC_RELATION) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_relation.GongGetGongstructName(), spec_relation.GongGetReferenceOrder(stage))
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_relation_type.GongGetGongstructName(), spec_relation_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spec_relation_type *SPEC_RELATION_TYPE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_relation_type.GongGetGongstructName(), spec_relation_type.GongGetReferenceOrder(stage))
}

func (xhtml_content *XHTML_CONTENT) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xhtml_content.GongGetGongstructName(), xhtml_content.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (xhtml_content *XHTML_CONTENT) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xhtml_content.GongGetGongstructName(), xhtml_content.GongGetReferenceOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (alternative_id *ALTERNATIVE_ID) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", alternative_id.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ALTERNATIVE_ID")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", alternative_id.Name)
	return
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_boolean.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_BOOLEAN")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_boolean.Name)
	return
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_date.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_DATE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_date.Name)
	return
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_enumeration.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_ENUMERATION")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_enumeration.Name)
	return
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_integer.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_INTEGER")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_integer.Name)
	return
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_real.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_REAL")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_real.Name)
	return
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_string.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_STRING")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_string.Name)
	return
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_xhtml.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_XHTML")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_definition_xhtml.Name)
	return
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_boolean.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_BOOLEAN")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_boolean.Name)
	return
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_date.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_DATE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_date.Name)
	return
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_enumeration.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_ENUMERATION")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_enumeration.Name)
	return
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_integer.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_INTEGER")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_integer.Name)
	return
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_real.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_REAL")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_real.Name)
	return
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_string.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_STRING")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_string.Name)
	return
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_xhtml.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_XHTML")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", attribute_value_xhtml.Name)
	return
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_alternative_id.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ALTERNATIVE_ID")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_alternative_id.Name)
	return
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_boolean_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_definition_boolean_ref.Name)
	return
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_date_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_DATE_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_definition_date_ref.Name)
	return
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_enumeration_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_definition_enumeration_ref.Name)
	return
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_integer_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_INTEGER_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_definition_integer_ref.Name)
	return
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_real_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_REAL_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_definition_real_ref.Name)
	return
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_string_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_STRING_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_definition_string_ref.Name)
	return
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_xhtml_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_XHTML_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_definition_xhtml_ref.Name)
	return
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_boolean.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_BOOLEAN")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_value_boolean.Name)
	return
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_date.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_DATE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_value_date.Name)
	return
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_enumeration.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_ENUMERATION")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_value_enumeration.Name)
	return
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_integer.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_INTEGER")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_value_integer.Name)
	return
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_real.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_REAL")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_value_real.Name)
	return
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_string.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_STRING")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_value_string.Name)
	return
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_xhtml.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_XHTML")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_value_xhtml.Name)
	return
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_xhtml_1.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_XHTML_1")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_attribute_value_xhtml_1.Name)
	return
}

func (a_children *A_CHILDREN) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_children.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_CHILDREN")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_children.Name)
	return
}

func (a_core_content *A_CORE_CONTENT) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_core_content.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_CORE_CONTENT")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_core_content.Name)
	return
}

func (a_datatypes *A_DATATYPES) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatypes.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPES")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatypes.Name)
	return
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_boolean_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_BOOLEAN_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatype_definition_boolean_ref.Name)
	return
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_date_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_DATE_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatype_definition_date_ref.Name)
	return
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_enumeration_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_ENUMERATION_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatype_definition_enumeration_ref.Name)
	return
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_integer_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_INTEGER_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatype_definition_integer_ref.Name)
	return
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_real_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_REAL_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatype_definition_real_ref.Name)
	return
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_string_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_STRING_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatype_definition_string_ref.Name)
	return
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_xhtml_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_XHTML_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_datatype_definition_xhtml_ref.Name)
	return
}

func (a_editable_atts *A_EDITABLE_ATTS) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_editable_atts.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_EDITABLE_ATTS")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_editable_atts.Name)
	return
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_enum_value_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ENUM_VALUE_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_enum_value_ref.Name)
	return
}

func (a_object *A_OBJECT) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_object.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_OBJECT")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_object.Name)
	return
}

func (a_properties *A_PROPERTIES) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_properties.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_PROPERTIES")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_properties.Name)
	return
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_relation_group_type_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_RELATION_GROUP_TYPE_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_relation_group_type_ref.Name)
	return
}

func (a_source_1 *A_SOURCE_1) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_source_1.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SOURCE_1")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_source_1.Name)
	return
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_source_specification_1.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SOURCE_SPECIFICATION_1")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_source_specification_1.Name)
	return
}

func (a_specifications *A_SPECIFICATIONS) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_specifications.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPECIFICATIONS")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_specifications.Name)
	return
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_specification_type_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPECIFICATION_TYPE_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_specification_type_ref.Name)
	return
}

func (a_specified_values *A_SPECIFIED_VALUES) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_specified_values.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPECIFIED_VALUES")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_specified_values.Name)
	return
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_attributes.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_ATTRIBUTES")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_attributes.Name)
	return
}

func (a_spec_objects *A_SPEC_OBJECTS) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_objects.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_OBJECTS")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_objects.Name)
	return
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_object_type_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_OBJECT_TYPE_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_object_type_ref.Name)
	return
}

func (a_spec_relations *A_SPEC_RELATIONS) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_relations.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_RELATIONS")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_relations.Name)
	return
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_relation_groups.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_RELATION_GROUPS")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_relation_groups.Name)
	return
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_relation_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_RELATION_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_relation_ref.Name)
	return
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_relation_type_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_RELATION_TYPE_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_relation_type_ref.Name)
	return
}

func (a_spec_types *A_SPEC_TYPES) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_types.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_TYPES")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_spec_types.Name)
	return
}

func (a_the_header *A_THE_HEADER) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_the_header.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_THE_HEADER")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_the_header.Name)
	return
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_tool_extensions.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_TOOL_EXTENSIONS")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", a_tool_extensions.Name)
	return
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_boolean.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_BOOLEAN")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_boolean.Name)
	return
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_date.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_DATE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_date.Name)
	return
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_enumeration.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_ENUMERATION")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_enumeration.Name)
	return
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_integer.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_INTEGER")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_integer.Name)
	return
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_real.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_REAL")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_real.Name)
	return
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_string.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_STRING")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_string.Name)
	return
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_xhtml.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_XHTML")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", datatype_definition_xhtml.Name)
	return
}

func (embedded_value *EMBEDDED_VALUE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", embedded_value.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EMBEDDED_VALUE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", embedded_value.Name)
	return
}

func (enum_value *ENUM_VALUE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", enum_value.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ENUM_VALUE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", enum_value.Name)
	return
}

func (relation_group *RELATION_GROUP) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RELATION_GROUP")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", relation_group.Name)
	return
}

func (relation_group_type *RELATION_GROUP_TYPE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", relation_group_type.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RELATION_GROUP_TYPE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", relation_group_type.Name)
	return
}

func (req_if *REQ_IF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", req_if.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQ_IF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", req_if.Name)
	return
}

func (req_if_content *REQ_IF_CONTENT) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", req_if_content.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQ_IF_CONTENT")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", req_if_content.Name)
	return
}

func (req_if_header *REQ_IF_HEADER) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", req_if_header.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQ_IF_HEADER")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", req_if_header.Name)
	return
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", req_if_tool_extension.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQ_IF_TOOL_EXTENSION")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", req_if_tool_extension.Name)
	return
}

func (specification *SPECIFICATION) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", specification.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECIFICATION")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specification.Name)
	return
}

func (specification_type *SPECIFICATION_TYPE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", specification_type.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECIFICATION_TYPE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", specification_type.Name)
	return
}

func (spec_hierarchy *SPEC_HIERARCHY) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_HIERARCHY")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spec_hierarchy.Name)
	return
}

func (spec_object *SPEC_OBJECT) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_object.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_OBJECT")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spec_object.Name)
	return
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_object_type.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_OBJECT_TYPE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spec_object_type.Name)
	return
}

func (spec_relation *SPEC_RELATION) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_RELATION")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spec_relation.Name)
	return
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_relation_type.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_RELATION_TYPE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", spec_relation_type.Name)
	return
}

func (xhtml_content *XHTML_CONTENT) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xhtml_content.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XHTML_CONTENT")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", xhtml_content.Name)
	return
}

// insertion point for unstaging
func (alternative_id *ALTERNATIVE_ID) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", alternative_id.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_boolean.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_date.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_enumeration.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_integer.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_real.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_string.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_xhtml.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_boolean.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_date.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_enumeration.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_integer.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_real.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_string.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_xhtml.GongGetReferenceIdentifier(stage))
	return
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_alternative_id.GongGetReferenceIdentifier(stage))
	return
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_boolean_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_date_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_enumeration_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_integer_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_real_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_string_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_xhtml_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_boolean.GongGetReferenceIdentifier(stage))
	return
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_date.GongGetReferenceIdentifier(stage))
	return
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_enumeration.GongGetReferenceIdentifier(stage))
	return
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_integer.GongGetReferenceIdentifier(stage))
	return
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_real.GongGetReferenceIdentifier(stage))
	return
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_string.GongGetReferenceIdentifier(stage))
	return
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_xhtml.GongGetReferenceIdentifier(stage))
	return
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_xhtml_1.GongGetReferenceIdentifier(stage))
	return
}

func (a_children *A_CHILDREN) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_children.GongGetReferenceIdentifier(stage))
	return
}

func (a_core_content *A_CORE_CONTENT) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_core_content.GongGetReferenceIdentifier(stage))
	return
}

func (a_datatypes *A_DATATYPES) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatypes.GongGetReferenceIdentifier(stage))
	return
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_boolean_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_date_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_enumeration_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_integer_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_real_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_string_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_xhtml_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_editable_atts *A_EDITABLE_ATTS) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_editable_atts.GongGetReferenceIdentifier(stage))
	return
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_enum_value_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_object *A_OBJECT) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_object.GongGetReferenceIdentifier(stage))
	return
}

func (a_properties *A_PROPERTIES) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_properties.GongGetReferenceIdentifier(stage))
	return
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_relation_group_type_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_source_1 *A_SOURCE_1) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_source_1.GongGetReferenceIdentifier(stage))
	return
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_source_specification_1.GongGetReferenceIdentifier(stage))
	return
}

func (a_specifications *A_SPECIFICATIONS) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_specifications.GongGetReferenceIdentifier(stage))
	return
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_specification_type_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_specified_values *A_SPECIFIED_VALUES) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_specified_values.GongGetReferenceIdentifier(stage))
	return
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_attributes.GongGetReferenceIdentifier(stage))
	return
}

func (a_spec_objects *A_SPEC_OBJECTS) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_objects.GongGetReferenceIdentifier(stage))
	return
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_object_type_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_spec_relations *A_SPEC_RELATIONS) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_relations.GongGetReferenceIdentifier(stage))
	return
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_relation_groups.GongGetReferenceIdentifier(stage))
	return
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_relation_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_relation_type_ref.GongGetReferenceIdentifier(stage))
	return
}

func (a_spec_types *A_SPEC_TYPES) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_types.GongGetReferenceIdentifier(stage))
	return
}

func (a_the_header *A_THE_HEADER) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_the_header.GongGetReferenceIdentifier(stage))
	return
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_tool_extensions.GongGetReferenceIdentifier(stage))
	return
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_boolean.GongGetReferenceIdentifier(stage))
	return
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_date.GongGetReferenceIdentifier(stage))
	return
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_enumeration.GongGetReferenceIdentifier(stage))
	return
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_integer.GongGetReferenceIdentifier(stage))
	return
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_real.GongGetReferenceIdentifier(stage))
	return
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_string.GongGetReferenceIdentifier(stage))
	return
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_xhtml.GongGetReferenceIdentifier(stage))
	return
}

func (embedded_value *EMBEDDED_VALUE) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", embedded_value.GongGetReferenceIdentifier(stage))
	return
}

func (enum_value *ENUM_VALUE) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", enum_value.GongGetReferenceIdentifier(stage))
	return
}

func (relation_group *RELATION_GROUP) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", relation_group.GongGetReferenceIdentifier(stage))
	return
}

func (relation_group_type *RELATION_GROUP_TYPE) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", relation_group_type.GongGetReferenceIdentifier(stage))
	return
}

func (req_if *REQ_IF) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", req_if.GongGetReferenceIdentifier(stage))
	return
}

func (req_if_content *REQ_IF_CONTENT) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", req_if_content.GongGetReferenceIdentifier(stage))
	return
}

func (req_if_header *REQ_IF_HEADER) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", req_if_header.GongGetReferenceIdentifier(stage))
	return
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", req_if_tool_extension.GongGetReferenceIdentifier(stage))
	return
}

func (specification *SPECIFICATION) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", specification.GongGetReferenceIdentifier(stage))
	return
}

func (specification_type *SPECIFICATION_TYPE) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", specification_type.GongGetReferenceIdentifier(stage))
	return
}

func (spec_hierarchy *SPEC_HIERARCHY) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_hierarchy.GongGetReferenceIdentifier(stage))
	return
}

func (spec_object *SPEC_OBJECT) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_object.GongGetReferenceIdentifier(stage))
	return
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_object_type.GongGetReferenceIdentifier(stage))
	return
}

func (spec_relation *SPEC_RELATION) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_relation.GongGetReferenceIdentifier(stage))
	return
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_relation_type.GongGetReferenceIdentifier(stage))
	return
}

func (xhtml_content *XHTML_CONTENT) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xhtml_content.GongGetReferenceIdentifier(stage))
	return
}

// end of template
