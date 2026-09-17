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
	// Compute reverse map for named struct ALTERNATIVE_ID
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_BOOLEAN
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_BOOLEAN_Rendering
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_DATE
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_DATE_Rendering
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_ENUMERATION
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_ENUMERATION_Rendering
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_INTEGER
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_INTEGER_Rendering
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_REAL
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_REAL_Rendering
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_Rendering
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_STRING
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_STRING_Rendering
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_XHTML
	// insertion point per field

	// Compute reverse map for named struct ATTRIBUTE_DEFINITION_XHTML_Rendering
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

	// Compute reverse map for named struct EmbeddedJpgImage
	// insertion point per field

	// Compute reverse map for named struct EmbeddedPngImage
	// insertion point per field

	// Compute reverse map for named struct EmbeddedSvgImage
	// insertion point per field

	// Compute reverse map for named struct Kill
	// insertion point per field

	// Compute reverse map for named struct Map_identifier_bool
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

	// Compute reverse map for named struct SPECIFICATION_Rendering
	// insertion point per field

	// Compute reverse map for named struct SPECIFICATION_TYPE
	// insertion point per field

	// Compute reverse map for named struct SPEC_HIERARCHY
	// insertion point per field

	// Compute reverse map for named struct SPEC_OBJECT
	// insertion point per field

	// Compute reverse map for named struct SPEC_OBJECT_TYPE
	// insertion point per field

	// Compute reverse map for named struct SPEC_OBJECT_TYPE_Rendering
	// insertion point per field

	// Compute reverse map for named struct SPEC_RELATION
	// insertion point per field

	// Compute reverse map for named struct SPEC_RELATION_TYPE
	// insertion point per field

	// Compute reverse map for named struct StaticWebSite
	// insertion point per field
	stage.StaticWebSite_Chapters_reverseMap = make(map[*StaticWebSiteChapter]*StaticWebSite)
	for staticwebsite := range stage.StaticWebSites {
		_ = staticwebsite
		for _, _staticwebsitechapter := range staticwebsite.Chapters {
			stage.StaticWebSite_Chapters_reverseMap[_staticwebsitechapter] = staticwebsite
		}
	}

	// Compute reverse map for named struct StaticWebSiteChapter
	// insertion point per field
	stage.StaticWebSiteChapter_Paragraphs_reverseMap = make(map[*StaticWebSiteParagraph]*StaticWebSiteChapter)
	for staticwebsitechapter := range stage.StaticWebSiteChapters {
		_ = staticwebsitechapter
		for _, _staticwebsiteparagraph := range staticwebsitechapter.Paragraphs {
			stage.StaticWebSiteChapter_Paragraphs_reverseMap[_staticwebsiteparagraph] = staticwebsitechapter
		}
	}

	// Compute reverse map for named struct StaticWebSiteGeneratedImage
	// insertion point per field

	// Compute reverse map for named struct StaticWebSiteImage
	// insertion point per field

	// Compute reverse map for named struct StaticWebSiteParagraph
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

	for instance := range stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_DATEs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_DATE_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_REALs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_REAL_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_STRING_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		res = append(res, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_XHTML_Renderings {
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

	for instance := range stage.EmbeddedJpgImages {
		res = append(res, instance)
	}

	for instance := range stage.EmbeddedPngImages {
		res = append(res, instance)
	}

	for instance := range stage.EmbeddedSvgImages {
		res = append(res, instance)
	}

	for instance := range stage.Kills {
		res = append(res, instance)
	}

	for instance := range stage.Map_identifier_bools {
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

	for instance := range stage.SPECIFICATION_Renderings {
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

	for instance := range stage.SPEC_OBJECT_TYPE_Renderings {
		res = append(res, instance)
	}

	for instance := range stage.SPEC_RELATIONs {
		res = append(res, instance)
	}

	for instance := range stage.SPEC_RELATION_TYPEs {
		res = append(res, instance)
	}

	for instance := range stage.StaticWebSites {
		res = append(res, instance)
	}

	for instance := range stage.StaticWebSiteChapters {
		res = append(res, instance)
	}

	for instance := range stage.StaticWebSiteGeneratedImages {
		res = append(res, instance)
	}

	for instance := range stage.StaticWebSiteImages {
		res = append(res, instance)
	}

	for instance := range stage.StaticWebSiteParagraphs {
		res = append(res, instance)
	}

	for instance := range stage.XHTML_CONTENTs {
		res = append(res, instance)
	}

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

func (attribute_definition_boolean_rendering *ATTRIBUTE_DEFINITION_BOOLEAN_Rendering) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_BOOLEAN_Rendering)
	attribute_definition_boolean_rendering.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_DATE)
	attribute_definition_date.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_date_rendering *ATTRIBUTE_DEFINITION_DATE_Rendering) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_DATE_Rendering)
	attribute_definition_date_rendering.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_ENUMERATION)
	attribute_definition_enumeration.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_enumeration_rendering *ATTRIBUTE_DEFINITION_ENUMERATION_Rendering) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_ENUMERATION_Rendering)
	attribute_definition_enumeration_rendering.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_INTEGER)
	attribute_definition_integer.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_integer_rendering *ATTRIBUTE_DEFINITION_INTEGER_Rendering) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_INTEGER_Rendering)
	attribute_definition_integer_rendering.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_REAL)
	attribute_definition_real.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_real_rendering *ATTRIBUTE_DEFINITION_REAL_Rendering) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_REAL_Rendering)
	attribute_definition_real_rendering.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_rendering *ATTRIBUTE_DEFINITION_Rendering) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_Rendering)
	attribute_definition_rendering.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_STRING)
	attribute_definition_string.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_string_rendering *ATTRIBUTE_DEFINITION_STRING_Rendering) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_STRING_Rendering)
	attribute_definition_string_rendering.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_XHTML)
	attribute_definition_xhtml.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attribute_definition_xhtml_rendering *ATTRIBUTE_DEFINITION_XHTML_Rendering) GongCopy() GongstructIF {
	newInstance := new(ATTRIBUTE_DEFINITION_XHTML_Rendering)
	attribute_definition_xhtml_rendering.GongCopyBasicFields(newInstance)
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

func (embeddedjpgimage *EmbeddedJpgImage) GongCopy() GongstructIF {
	newInstance := new(EmbeddedJpgImage)
	embeddedjpgimage.GongCopyBasicFields(newInstance)
	return newInstance
}

func (embeddedpngimage *EmbeddedPngImage) GongCopy() GongstructIF {
	newInstance := new(EmbeddedPngImage)
	embeddedpngimage.GongCopyBasicFields(newInstance)
	return newInstance
}

func (embeddedsvgimage *EmbeddedSvgImage) GongCopy() GongstructIF {
	newInstance := new(EmbeddedSvgImage)
	embeddedsvgimage.GongCopyBasicFields(newInstance)
	return newInstance
}

func (kill *Kill) GongCopy() GongstructIF {
	newInstance := new(Kill)
	kill.GongCopyBasicFields(newInstance)
	return newInstance
}

func (map_identifier_bool *Map_identifier_bool) GongCopy() GongstructIF {
	newInstance := new(Map_identifier_bool)
	map_identifier_bool.GongCopyBasicFields(newInstance)
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

func (specification_rendering *SPECIFICATION_Rendering) GongCopy() GongstructIF {
	newInstance := new(SPECIFICATION_Rendering)
	specification_rendering.GongCopyBasicFields(newInstance)
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

func (spec_object_type_rendering *SPEC_OBJECT_TYPE_Rendering) GongCopy() GongstructIF {
	newInstance := new(SPEC_OBJECT_TYPE_Rendering)
	spec_object_type_rendering.GongCopyBasicFields(newInstance)
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

func (staticwebsite *StaticWebSite) GongCopy() GongstructIF {
	newInstance := new(StaticWebSite)
	staticwebsite.GongCopyBasicFields(newInstance)
	return newInstance
}

func (staticwebsitechapter *StaticWebSiteChapter) GongCopy() GongstructIF {
	newInstance := new(StaticWebSiteChapter)
	staticwebsitechapter.GongCopyBasicFields(newInstance)
	return newInstance
}

func (staticwebsitegeneratedimage *StaticWebSiteGeneratedImage) GongCopy() GongstructIF {
	newInstance := new(StaticWebSiteGeneratedImage)
	staticwebsitegeneratedimage.GongCopyBasicFields(newInstance)
	return newInstance
}

func (staticwebsiteimage *StaticWebSiteImage) GongCopy() GongstructIF {
	newInstance := new(StaticWebSiteImage)
	staticwebsiteimage.GongCopyBasicFields(newInstance)
	return newInstance
}

func (staticwebsiteparagraph *StaticWebSiteParagraph) GongCopy() GongstructIF {
	newInstance := new(StaticWebSiteParagraph)
	staticwebsiteparagraph.GongCopyBasicFields(newInstance)
	return newInstance
}

func (xhtml_content *XHTML_CONTENT) GongCopy() GongstructIF {
	newInstance := new(XHTML_CONTENT)
	xhtml_content.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (alternative_id *ALTERNATIVE_ID) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(alternative_id).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(alternative_id), uint64(stage.GetOrder(alternative_id)))
	return
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_definition_boolean).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_definition_boolean), uint64(stage.GetOrder(attribute_definition_boolean)))
	return
}

func (attribute_definition_boolean_rendering *ATTRIBUTE_DEFINITION_BOOLEAN_Rendering) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_definition_boolean_rendering).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_definition_boolean_rendering), uint64(stage.GetOrder(attribute_definition_boolean_rendering)))
	return
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_definition_date).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_definition_date), uint64(stage.GetOrder(attribute_definition_date)))
	return
}

func (attribute_definition_date_rendering *ATTRIBUTE_DEFINITION_DATE_Rendering) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_definition_date_rendering).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_definition_date_rendering), uint64(stage.GetOrder(attribute_definition_date_rendering)))
	return
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_definition_enumeration).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_definition_enumeration), uint64(stage.GetOrder(attribute_definition_enumeration)))
	return
}

func (attribute_definition_enumeration_rendering *ATTRIBUTE_DEFINITION_ENUMERATION_Rendering) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_definition_enumeration_rendering).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_definition_enumeration_rendering), uint64(stage.GetOrder(attribute_definition_enumeration_rendering)))
	return
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_definition_integer).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_definition_integer), uint64(stage.GetOrder(attribute_definition_integer)))
	return
}

func (attribute_definition_integer_rendering *ATTRIBUTE_DEFINITION_INTEGER_Rendering) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_definition_integer_rendering).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_definition_integer_rendering), uint64(stage.GetOrder(attribute_definition_integer_rendering)))
	return
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_definition_real).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_definition_real), uint64(stage.GetOrder(attribute_definition_real)))
	return
}

func (attribute_definition_real_rendering *ATTRIBUTE_DEFINITION_REAL_Rendering) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_definition_real_rendering).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_definition_real_rendering), uint64(stage.GetOrder(attribute_definition_real_rendering)))
	return
}

func (attribute_definition_rendering *ATTRIBUTE_DEFINITION_Rendering) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_definition_rendering).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_definition_rendering), uint64(stage.GetOrder(attribute_definition_rendering)))
	return
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_definition_string).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_definition_string), uint64(stage.GetOrder(attribute_definition_string)))
	return
}

func (attribute_definition_string_rendering *ATTRIBUTE_DEFINITION_STRING_Rendering) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_definition_string_rendering).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_definition_string_rendering), uint64(stage.GetOrder(attribute_definition_string_rendering)))
	return
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_definition_xhtml).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_definition_xhtml), uint64(stage.GetOrder(attribute_definition_xhtml)))
	return
}

func (attribute_definition_xhtml_rendering *ATTRIBUTE_DEFINITION_XHTML_Rendering) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_definition_xhtml_rendering).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_definition_xhtml_rendering), uint64(stage.GetOrder(attribute_definition_xhtml_rendering)))
	return
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_value_boolean).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_value_boolean), uint64(stage.GetOrder(attribute_value_boolean)))
	return
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_value_date).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_value_date), uint64(stage.GetOrder(attribute_value_date)))
	return
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_value_enumeration).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_value_enumeration), uint64(stage.GetOrder(attribute_value_enumeration)))
	return
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_value_integer).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_value_integer), uint64(stage.GetOrder(attribute_value_integer)))
	return
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_value_real).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_value_real), uint64(stage.GetOrder(attribute_value_real)))
	return
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_value_string).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_value_string), uint64(stage.GetOrder(attribute_value_string)))
	return
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attribute_value_xhtml).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attribute_value_xhtml), uint64(stage.GetOrder(attribute_value_xhtml)))
	return
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_alternative_id).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_alternative_id), uint64(stage.GetOrder(a_alternative_id)))
	return
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_attribute_definition_boolean_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_attribute_definition_boolean_ref), uint64(stage.GetOrder(a_attribute_definition_boolean_ref)))
	return
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_attribute_definition_date_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_attribute_definition_date_ref), uint64(stage.GetOrder(a_attribute_definition_date_ref)))
	return
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_attribute_definition_enumeration_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_attribute_definition_enumeration_ref), uint64(stage.GetOrder(a_attribute_definition_enumeration_ref)))
	return
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_attribute_definition_integer_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_attribute_definition_integer_ref), uint64(stage.GetOrder(a_attribute_definition_integer_ref)))
	return
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_attribute_definition_real_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_attribute_definition_real_ref), uint64(stage.GetOrder(a_attribute_definition_real_ref)))
	return
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_attribute_definition_string_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_attribute_definition_string_ref), uint64(stage.GetOrder(a_attribute_definition_string_ref)))
	return
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_attribute_definition_xhtml_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_attribute_definition_xhtml_ref), uint64(stage.GetOrder(a_attribute_definition_xhtml_ref)))
	return
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_attribute_value_boolean).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_attribute_value_boolean), uint64(stage.GetOrder(a_attribute_value_boolean)))
	return
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_attribute_value_date).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_attribute_value_date), uint64(stage.GetOrder(a_attribute_value_date)))
	return
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_attribute_value_enumeration).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_attribute_value_enumeration), uint64(stage.GetOrder(a_attribute_value_enumeration)))
	return
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_attribute_value_integer).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_attribute_value_integer), uint64(stage.GetOrder(a_attribute_value_integer)))
	return
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_attribute_value_real).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_attribute_value_real), uint64(stage.GetOrder(a_attribute_value_real)))
	return
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_attribute_value_string).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_attribute_value_string), uint64(stage.GetOrder(a_attribute_value_string)))
	return
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_attribute_value_xhtml).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_attribute_value_xhtml), uint64(stage.GetOrder(a_attribute_value_xhtml)))
	return
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_attribute_value_xhtml_1).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_attribute_value_xhtml_1), uint64(stage.GetOrder(a_attribute_value_xhtml_1)))
	return
}

func (a_children *A_CHILDREN) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_children).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_children), uint64(stage.GetOrder(a_children)))
	return
}

func (a_core_content *A_CORE_CONTENT) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_core_content).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_core_content), uint64(stage.GetOrder(a_core_content)))
	return
}

func (a_datatypes *A_DATATYPES) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_datatypes).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_datatypes), uint64(stage.GetOrder(a_datatypes)))
	return
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_datatype_definition_boolean_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_datatype_definition_boolean_ref), uint64(stage.GetOrder(a_datatype_definition_boolean_ref)))
	return
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_datatype_definition_date_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_datatype_definition_date_ref), uint64(stage.GetOrder(a_datatype_definition_date_ref)))
	return
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_datatype_definition_enumeration_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_datatype_definition_enumeration_ref), uint64(stage.GetOrder(a_datatype_definition_enumeration_ref)))
	return
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_datatype_definition_integer_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_datatype_definition_integer_ref), uint64(stage.GetOrder(a_datatype_definition_integer_ref)))
	return
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_datatype_definition_real_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_datatype_definition_real_ref), uint64(stage.GetOrder(a_datatype_definition_real_ref)))
	return
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_datatype_definition_string_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_datatype_definition_string_ref), uint64(stage.GetOrder(a_datatype_definition_string_ref)))
	return
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_datatype_definition_xhtml_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_datatype_definition_xhtml_ref), uint64(stage.GetOrder(a_datatype_definition_xhtml_ref)))
	return
}

func (a_editable_atts *A_EDITABLE_ATTS) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_editable_atts).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_editable_atts), uint64(stage.GetOrder(a_editable_atts)))
	return
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_enum_value_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_enum_value_ref), uint64(stage.GetOrder(a_enum_value_ref)))
	return
}

func (a_object *A_OBJECT) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_object).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_object), uint64(stage.GetOrder(a_object)))
	return
}

func (a_properties *A_PROPERTIES) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_properties).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_properties), uint64(stage.GetOrder(a_properties)))
	return
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_relation_group_type_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_relation_group_type_ref), uint64(stage.GetOrder(a_relation_group_type_ref)))
	return
}

func (a_source_1 *A_SOURCE_1) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_source_1).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_source_1), uint64(stage.GetOrder(a_source_1)))
	return
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_source_specification_1).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_source_specification_1), uint64(stage.GetOrder(a_source_specification_1)))
	return
}

func (a_specifications *A_SPECIFICATIONS) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_specifications).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_specifications), uint64(stage.GetOrder(a_specifications)))
	return
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_specification_type_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_specification_type_ref), uint64(stage.GetOrder(a_specification_type_ref)))
	return
}

func (a_specified_values *A_SPECIFIED_VALUES) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_specified_values).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_specified_values), uint64(stage.GetOrder(a_specified_values)))
	return
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_spec_attributes).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_spec_attributes), uint64(stage.GetOrder(a_spec_attributes)))
	return
}

func (a_spec_objects *A_SPEC_OBJECTS) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_spec_objects).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_spec_objects), uint64(stage.GetOrder(a_spec_objects)))
	return
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_spec_object_type_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_spec_object_type_ref), uint64(stage.GetOrder(a_spec_object_type_ref)))
	return
}

func (a_spec_relations *A_SPEC_RELATIONS) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_spec_relations).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_spec_relations), uint64(stage.GetOrder(a_spec_relations)))
	return
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_spec_relation_groups).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_spec_relation_groups), uint64(stage.GetOrder(a_spec_relation_groups)))
	return
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_spec_relation_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_spec_relation_ref), uint64(stage.GetOrder(a_spec_relation_ref)))
	return
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_spec_relation_type_ref).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_spec_relation_type_ref), uint64(stage.GetOrder(a_spec_relation_type_ref)))
	return
}

func (a_spec_types *A_SPEC_TYPES) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_spec_types).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_spec_types), uint64(stage.GetOrder(a_spec_types)))
	return
}

func (a_the_header *A_THE_HEADER) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_the_header).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_the_header), uint64(stage.GetOrder(a_the_header)))
	return
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_tool_extensions).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_tool_extensions), uint64(stage.GetOrder(a_tool_extensions)))
	return
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(datatype_definition_boolean).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(datatype_definition_boolean), uint64(stage.GetOrder(datatype_definition_boolean)))
	return
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(datatype_definition_date).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(datatype_definition_date), uint64(stage.GetOrder(datatype_definition_date)))
	return
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(datatype_definition_enumeration).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(datatype_definition_enumeration), uint64(stage.GetOrder(datatype_definition_enumeration)))
	return
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(datatype_definition_integer).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(datatype_definition_integer), uint64(stage.GetOrder(datatype_definition_integer)))
	return
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(datatype_definition_real).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(datatype_definition_real), uint64(stage.GetOrder(datatype_definition_real)))
	return
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(datatype_definition_string).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(datatype_definition_string), uint64(stage.GetOrder(datatype_definition_string)))
	return
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(datatype_definition_xhtml).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(datatype_definition_xhtml), uint64(stage.GetOrder(datatype_definition_xhtml)))
	return
}

func (embedded_value *EMBEDDED_VALUE) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(embedded_value).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(embedded_value), uint64(stage.GetOrder(embedded_value)))
	return
}

func (enum_value *ENUM_VALUE) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(enum_value).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(enum_value), uint64(stage.GetOrder(enum_value)))
	return
}

func (embeddedjpgimage *EmbeddedJpgImage) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(embeddedjpgimage).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(embeddedjpgimage), uint64(stage.GetOrder(embeddedjpgimage)))
	return
}

func (embeddedpngimage *EmbeddedPngImage) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(embeddedpngimage).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(embeddedpngimage), uint64(stage.GetOrder(embeddedpngimage)))
	return
}

func (embeddedsvgimage *EmbeddedSvgImage) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(embeddedsvgimage).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(embeddedsvgimage), uint64(stage.GetOrder(embeddedsvgimage)))
	return
}

func (kill *Kill) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(kill).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(kill), uint64(stage.GetOrder(kill)))
	return
}

func (map_identifier_bool *Map_identifier_bool) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(map_identifier_bool).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(map_identifier_bool), uint64(stage.GetOrder(map_identifier_bool)))
	return
}

func (relation_group *RELATION_GROUP) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(relation_group).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(relation_group), uint64(stage.GetOrder(relation_group)))
	return
}

func (relation_group_type *RELATION_GROUP_TYPE) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(relation_group_type).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(relation_group_type), uint64(stage.GetOrder(relation_group_type)))
	return
}

func (req_if *REQ_IF) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(req_if).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(req_if), uint64(stage.GetOrder(req_if)))
	return
}

func (req_if_content *REQ_IF_CONTENT) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(req_if_content).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(req_if_content), uint64(stage.GetOrder(req_if_content)))
	return
}

func (req_if_header *REQ_IF_HEADER) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(req_if_header).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(req_if_header), uint64(stage.GetOrder(req_if_header)))
	return
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(req_if_tool_extension).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(req_if_tool_extension), uint64(stage.GetOrder(req_if_tool_extension)))
	return
}

func (specification *SPECIFICATION) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(specification).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(specification), uint64(stage.GetOrder(specification)))
	return
}

func (specification_rendering *SPECIFICATION_Rendering) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(specification_rendering).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(specification_rendering), uint64(stage.GetOrder(specification_rendering)))
	return
}

func (specification_type *SPECIFICATION_TYPE) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(specification_type).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(specification_type), uint64(stage.GetOrder(specification_type)))
	return
}

func (spec_hierarchy *SPEC_HIERARCHY) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spec_hierarchy).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(spec_hierarchy), uint64(stage.GetOrder(spec_hierarchy)))
	return
}

func (spec_object *SPEC_OBJECT) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spec_object).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(spec_object), uint64(stage.GetOrder(spec_object)))
	return
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spec_object_type).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(spec_object_type), uint64(stage.GetOrder(spec_object_type)))
	return
}

func (spec_object_type_rendering *SPEC_OBJECT_TYPE_Rendering) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spec_object_type_rendering).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(spec_object_type_rendering), uint64(stage.GetOrder(spec_object_type_rendering)))
	return
}

func (spec_relation *SPEC_RELATION) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spec_relation).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(spec_relation), uint64(stage.GetOrder(spec_relation)))
	return
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(spec_relation_type).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(spec_relation_type), uint64(stage.GetOrder(spec_relation_type)))
	return
}

func (staticwebsite *StaticWebSite) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(staticwebsite).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(staticwebsite), uint64(stage.GetOrder(staticwebsite)))
	return
}

func (staticwebsitechapter *StaticWebSiteChapter) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(staticwebsitechapter).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(staticwebsitechapter), uint64(stage.GetOrder(staticwebsitechapter)))
	return
}

func (staticwebsitegeneratedimage *StaticWebSiteGeneratedImage) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(staticwebsitegeneratedimage).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(staticwebsitegeneratedimage), uint64(stage.GetOrder(staticwebsitegeneratedimage)))
	return
}

func (staticwebsiteimage *StaticWebSiteImage) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(staticwebsiteimage).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(staticwebsiteimage), uint64(stage.GetOrder(staticwebsiteimage)))
	return
}

func (staticwebsiteparagraph *StaticWebSiteParagraph) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(staticwebsiteparagraph).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(staticwebsiteparagraph), uint64(stage.GetOrder(staticwebsiteparagraph)))
	return
}

func (xhtml_content *XHTML_CONTENT) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(xhtml_content).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(xhtml_content), uint64(stage.GetOrder(xhtml_content)))
	return
}


type GongstructDiffable[T any] interface {
	PointerToGongstruct
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
		stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings,
		stage.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering_stagedOrder,
		stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings_reference,
		&stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings_referenceOrder,
		stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings_instance,
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
		stage.ATTRIBUTE_DEFINITION_DATE_Renderings,
		stage.ATTRIBUTE_DEFINITION_DATE_Rendering_stagedOrder,
		stage.ATTRIBUTE_DEFINITION_DATE_Renderings_reference,
		&stage.ATTRIBUTE_DEFINITION_DATE_Renderings_referenceOrder,
		stage.ATTRIBUTE_DEFINITION_DATE_Renderings_instance,
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
		stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings,
		stage.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering_stagedOrder,
		stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings_reference,
		&stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings_referenceOrder,
		stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings_instance,
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
		stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings,
		stage.ATTRIBUTE_DEFINITION_INTEGER_Rendering_stagedOrder,
		stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings_reference,
		&stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings_referenceOrder,
		stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings_instance,
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
		stage.ATTRIBUTE_DEFINITION_REAL_Renderings,
		stage.ATTRIBUTE_DEFINITION_REAL_Rendering_stagedOrder,
		stage.ATTRIBUTE_DEFINITION_REAL_Renderings_reference,
		&stage.ATTRIBUTE_DEFINITION_REAL_Renderings_referenceOrder,
		stage.ATTRIBUTE_DEFINITION_REAL_Renderings_instance,
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
		stage.ATTRIBUTE_DEFINITION_Renderings,
		stage.ATTRIBUTE_DEFINITION_Rendering_stagedOrder,
		stage.ATTRIBUTE_DEFINITION_Renderings_reference,
		&stage.ATTRIBUTE_DEFINITION_Renderings_referenceOrder,
		stage.ATTRIBUTE_DEFINITION_Renderings_instance,
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
		stage.ATTRIBUTE_DEFINITION_STRING_Renderings,
		stage.ATTRIBUTE_DEFINITION_STRING_Rendering_stagedOrder,
		stage.ATTRIBUTE_DEFINITION_STRING_Renderings_reference,
		&stage.ATTRIBUTE_DEFINITION_STRING_Renderings_referenceOrder,
		stage.ATTRIBUTE_DEFINITION_STRING_Renderings_instance,
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
		stage.ATTRIBUTE_DEFINITION_XHTML_Renderings,
		stage.ATTRIBUTE_DEFINITION_XHTML_Rendering_stagedOrder,
		stage.ATTRIBUTE_DEFINITION_XHTML_Renderings_reference,
		&stage.ATTRIBUTE_DEFINITION_XHTML_Renderings_referenceOrder,
		stage.ATTRIBUTE_DEFINITION_XHTML_Renderings_instance,
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
		stage.EmbeddedJpgImages,
		stage.EmbeddedJpgImage_stagedOrder,
		stage.EmbeddedJpgImages_reference,
		&stage.EmbeddedJpgImages_referenceOrder,
		stage.EmbeddedJpgImages_instance,
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
		stage.EmbeddedPngImages,
		stage.EmbeddedPngImage_stagedOrder,
		stage.EmbeddedPngImages_reference,
		&stage.EmbeddedPngImages_referenceOrder,
		stage.EmbeddedPngImages_instance,
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
		stage.EmbeddedSvgImages,
		stage.EmbeddedSvgImage_stagedOrder,
		stage.EmbeddedSvgImages_reference,
		&stage.EmbeddedSvgImages_referenceOrder,
		stage.EmbeddedSvgImages_instance,
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
		stage.Kills,
		stage.Kill_stagedOrder,
		stage.Kills_reference,
		&stage.Kills_referenceOrder,
		stage.Kills_instance,
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
		stage.Map_identifier_bools,
		stage.Map_identifier_bool_stagedOrder,
		stage.Map_identifier_bools_reference,
		&stage.Map_identifier_bools_referenceOrder,
		stage.Map_identifier_bools_instance,
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
		stage.SPECIFICATION_Renderings,
		stage.SPECIFICATION_Rendering_stagedOrder,
		stage.SPECIFICATION_Renderings_reference,
		&stage.SPECIFICATION_Renderings_referenceOrder,
		stage.SPECIFICATION_Renderings_instance,
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
		stage.SPEC_OBJECT_TYPE_Renderings,
		stage.SPEC_OBJECT_TYPE_Rendering_stagedOrder,
		stage.SPEC_OBJECT_TYPE_Renderings_reference,
		&stage.SPEC_OBJECT_TYPE_Renderings_referenceOrder,
		stage.SPEC_OBJECT_TYPE_Renderings_instance,
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
		stage.StaticWebSites,
		stage.StaticWebSite_stagedOrder,
		stage.StaticWebSites_reference,
		&stage.StaticWebSites_referenceOrder,
		stage.StaticWebSites_instance,
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
		stage.StaticWebSiteChapters,
		stage.StaticWebSiteChapter_stagedOrder,
		stage.StaticWebSiteChapters_reference,
		&stage.StaticWebSiteChapters_referenceOrder,
		stage.StaticWebSiteChapters_instance,
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
		stage.StaticWebSiteGeneratedImages,
		stage.StaticWebSiteGeneratedImage_stagedOrder,
		stage.StaticWebSiteGeneratedImages_reference,
		&stage.StaticWebSiteGeneratedImages_referenceOrder,
		stage.StaticWebSiteGeneratedImages_instance,
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
		stage.StaticWebSiteImages,
		stage.StaticWebSiteImage_stagedOrder,
		stage.StaticWebSiteImages_reference,
		&stage.StaticWebSiteImages_referenceOrder,
		stage.StaticWebSiteImages_instance,
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
		stage.StaticWebSiteParagraphs,
		stage.StaticWebSiteParagraph_stagedOrder,
		stage.StaticWebSiteParagraphs_reference,
		&stage.StaticWebSiteParagraphs_referenceOrder,
		stage.StaticWebSiteParagraphs_instance,
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
	stage.ALTERNATIVE_IDs_reference = make(map[*ALTERNATIVE_ID]*ALTERNATIVE_ID)
	stage.ALTERNATIVE_IDs_referenceOrder = make(map[*ALTERNATIVE_ID]uint) // diff Unstage needs the reference order
	stage.ALTERNATIVE_IDs_instance = make(map[*ALTERNATIVE_ID]*ALTERNATIVE_ID)
	for instance := range stage.ALTERNATIVE_IDs {
		_copy := instance.GongCopy().(*ALTERNATIVE_ID)
		stage.ALTERNATIVE_IDs_reference[instance] = _copy
		stage.ALTERNATIVE_IDs_instance[_copy] = instance
		stage.ALTERNATIVE_IDs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_BOOLEANs_reference = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*ATTRIBUTE_DEFINITION_BOOLEAN)
	stage.ATTRIBUTE_DEFINITION_BOOLEANs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_DEFINITION_BOOLEANs_instance = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN]*ATTRIBUTE_DEFINITION_BOOLEAN)
	for instance := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		_copy := instance.GongCopy().(*ATTRIBUTE_DEFINITION_BOOLEAN)
		stage.ATTRIBUTE_DEFINITION_BOOLEANs_reference[instance] = _copy
		stage.ATTRIBUTE_DEFINITION_BOOLEANs_instance[_copy] = instance
		stage.ATTRIBUTE_DEFINITION_BOOLEANs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings_reference = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN_Rendering]*ATTRIBUTE_DEFINITION_BOOLEAN_Rendering)
	stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN_Rendering]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings_instance = make(map[*ATTRIBUTE_DEFINITION_BOOLEAN_Rendering]*ATTRIBUTE_DEFINITION_BOOLEAN_Rendering)
	for instance := range stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings {
		_copy := instance.GongCopy().(*ATTRIBUTE_DEFINITION_BOOLEAN_Rendering)
		stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings_reference[instance] = _copy
		stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings_instance[_copy] = instance
		stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_DATEs_reference = make(map[*ATTRIBUTE_DEFINITION_DATE]*ATTRIBUTE_DEFINITION_DATE)
	stage.ATTRIBUTE_DEFINITION_DATEs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_DATE]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_DEFINITION_DATEs_instance = make(map[*ATTRIBUTE_DEFINITION_DATE]*ATTRIBUTE_DEFINITION_DATE)
	for instance := range stage.ATTRIBUTE_DEFINITION_DATEs {
		_copy := instance.GongCopy().(*ATTRIBUTE_DEFINITION_DATE)
		stage.ATTRIBUTE_DEFINITION_DATEs_reference[instance] = _copy
		stage.ATTRIBUTE_DEFINITION_DATEs_instance[_copy] = instance
		stage.ATTRIBUTE_DEFINITION_DATEs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_DATE_Renderings_reference = make(map[*ATTRIBUTE_DEFINITION_DATE_Rendering]*ATTRIBUTE_DEFINITION_DATE_Rendering)
	stage.ATTRIBUTE_DEFINITION_DATE_Renderings_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_DATE_Rendering]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_DEFINITION_DATE_Renderings_instance = make(map[*ATTRIBUTE_DEFINITION_DATE_Rendering]*ATTRIBUTE_DEFINITION_DATE_Rendering)
	for instance := range stage.ATTRIBUTE_DEFINITION_DATE_Renderings {
		_copy := instance.GongCopy().(*ATTRIBUTE_DEFINITION_DATE_Rendering)
		stage.ATTRIBUTE_DEFINITION_DATE_Renderings_reference[instance] = _copy
		stage.ATTRIBUTE_DEFINITION_DATE_Renderings_instance[_copy] = instance
		stage.ATTRIBUTE_DEFINITION_DATE_Renderings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_reference = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*ATTRIBUTE_DEFINITION_ENUMERATION)
	stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_instance = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION]*ATTRIBUTE_DEFINITION_ENUMERATION)
	for instance := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		_copy := instance.GongCopy().(*ATTRIBUTE_DEFINITION_ENUMERATION)
		stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_reference[instance] = _copy
		stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_instance[_copy] = instance
		stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings_reference = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION_Rendering]*ATTRIBUTE_DEFINITION_ENUMERATION_Rendering)
	stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION_Rendering]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings_instance = make(map[*ATTRIBUTE_DEFINITION_ENUMERATION_Rendering]*ATTRIBUTE_DEFINITION_ENUMERATION_Rendering)
	for instance := range stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings {
		_copy := instance.GongCopy().(*ATTRIBUTE_DEFINITION_ENUMERATION_Rendering)
		stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings_reference[instance] = _copy
		stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings_instance[_copy] = instance
		stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_INTEGERs_reference = make(map[*ATTRIBUTE_DEFINITION_INTEGER]*ATTRIBUTE_DEFINITION_INTEGER)
	stage.ATTRIBUTE_DEFINITION_INTEGERs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_INTEGER]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_DEFINITION_INTEGERs_instance = make(map[*ATTRIBUTE_DEFINITION_INTEGER]*ATTRIBUTE_DEFINITION_INTEGER)
	for instance := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		_copy := instance.GongCopy().(*ATTRIBUTE_DEFINITION_INTEGER)
		stage.ATTRIBUTE_DEFINITION_INTEGERs_reference[instance] = _copy
		stage.ATTRIBUTE_DEFINITION_INTEGERs_instance[_copy] = instance
		stage.ATTRIBUTE_DEFINITION_INTEGERs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings_reference = make(map[*ATTRIBUTE_DEFINITION_INTEGER_Rendering]*ATTRIBUTE_DEFINITION_INTEGER_Rendering)
	stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_INTEGER_Rendering]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings_instance = make(map[*ATTRIBUTE_DEFINITION_INTEGER_Rendering]*ATTRIBUTE_DEFINITION_INTEGER_Rendering)
	for instance := range stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings {
		_copy := instance.GongCopy().(*ATTRIBUTE_DEFINITION_INTEGER_Rendering)
		stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings_reference[instance] = _copy
		stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings_instance[_copy] = instance
		stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_REALs_reference = make(map[*ATTRIBUTE_DEFINITION_REAL]*ATTRIBUTE_DEFINITION_REAL)
	stage.ATTRIBUTE_DEFINITION_REALs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_REAL]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_DEFINITION_REALs_instance = make(map[*ATTRIBUTE_DEFINITION_REAL]*ATTRIBUTE_DEFINITION_REAL)
	for instance := range stage.ATTRIBUTE_DEFINITION_REALs {
		_copy := instance.GongCopy().(*ATTRIBUTE_DEFINITION_REAL)
		stage.ATTRIBUTE_DEFINITION_REALs_reference[instance] = _copy
		stage.ATTRIBUTE_DEFINITION_REALs_instance[_copy] = instance
		stage.ATTRIBUTE_DEFINITION_REALs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_REAL_Renderings_reference = make(map[*ATTRIBUTE_DEFINITION_REAL_Rendering]*ATTRIBUTE_DEFINITION_REAL_Rendering)
	stage.ATTRIBUTE_DEFINITION_REAL_Renderings_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_REAL_Rendering]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_DEFINITION_REAL_Renderings_instance = make(map[*ATTRIBUTE_DEFINITION_REAL_Rendering]*ATTRIBUTE_DEFINITION_REAL_Rendering)
	for instance := range stage.ATTRIBUTE_DEFINITION_REAL_Renderings {
		_copy := instance.GongCopy().(*ATTRIBUTE_DEFINITION_REAL_Rendering)
		stage.ATTRIBUTE_DEFINITION_REAL_Renderings_reference[instance] = _copy
		stage.ATTRIBUTE_DEFINITION_REAL_Renderings_instance[_copy] = instance
		stage.ATTRIBUTE_DEFINITION_REAL_Renderings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_Renderings_reference = make(map[*ATTRIBUTE_DEFINITION_Rendering]*ATTRIBUTE_DEFINITION_Rendering)
	stage.ATTRIBUTE_DEFINITION_Renderings_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_Rendering]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_DEFINITION_Renderings_instance = make(map[*ATTRIBUTE_DEFINITION_Rendering]*ATTRIBUTE_DEFINITION_Rendering)
	for instance := range stage.ATTRIBUTE_DEFINITION_Renderings {
		_copy := instance.GongCopy().(*ATTRIBUTE_DEFINITION_Rendering)
		stage.ATTRIBUTE_DEFINITION_Renderings_reference[instance] = _copy
		stage.ATTRIBUTE_DEFINITION_Renderings_instance[_copy] = instance
		stage.ATTRIBUTE_DEFINITION_Renderings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_STRINGs_reference = make(map[*ATTRIBUTE_DEFINITION_STRING]*ATTRIBUTE_DEFINITION_STRING)
	stage.ATTRIBUTE_DEFINITION_STRINGs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_STRING]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_DEFINITION_STRINGs_instance = make(map[*ATTRIBUTE_DEFINITION_STRING]*ATTRIBUTE_DEFINITION_STRING)
	for instance := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		_copy := instance.GongCopy().(*ATTRIBUTE_DEFINITION_STRING)
		stage.ATTRIBUTE_DEFINITION_STRINGs_reference[instance] = _copy
		stage.ATTRIBUTE_DEFINITION_STRINGs_instance[_copy] = instance
		stage.ATTRIBUTE_DEFINITION_STRINGs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_STRING_Renderings_reference = make(map[*ATTRIBUTE_DEFINITION_STRING_Rendering]*ATTRIBUTE_DEFINITION_STRING_Rendering)
	stage.ATTRIBUTE_DEFINITION_STRING_Renderings_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_STRING_Rendering]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_DEFINITION_STRING_Renderings_instance = make(map[*ATTRIBUTE_DEFINITION_STRING_Rendering]*ATTRIBUTE_DEFINITION_STRING_Rendering)
	for instance := range stage.ATTRIBUTE_DEFINITION_STRING_Renderings {
		_copy := instance.GongCopy().(*ATTRIBUTE_DEFINITION_STRING_Rendering)
		stage.ATTRIBUTE_DEFINITION_STRING_Renderings_reference[instance] = _copy
		stage.ATTRIBUTE_DEFINITION_STRING_Renderings_instance[_copy] = instance
		stage.ATTRIBUTE_DEFINITION_STRING_Renderings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_XHTMLs_reference = make(map[*ATTRIBUTE_DEFINITION_XHTML]*ATTRIBUTE_DEFINITION_XHTML)
	stage.ATTRIBUTE_DEFINITION_XHTMLs_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_XHTML]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_DEFINITION_XHTMLs_instance = make(map[*ATTRIBUTE_DEFINITION_XHTML]*ATTRIBUTE_DEFINITION_XHTML)
	for instance := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		_copy := instance.GongCopy().(*ATTRIBUTE_DEFINITION_XHTML)
		stage.ATTRIBUTE_DEFINITION_XHTMLs_reference[instance] = _copy
		stage.ATTRIBUTE_DEFINITION_XHTMLs_instance[_copy] = instance
		stage.ATTRIBUTE_DEFINITION_XHTMLs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_DEFINITION_XHTML_Renderings_reference = make(map[*ATTRIBUTE_DEFINITION_XHTML_Rendering]*ATTRIBUTE_DEFINITION_XHTML_Rendering)
	stage.ATTRIBUTE_DEFINITION_XHTML_Renderings_referenceOrder = make(map[*ATTRIBUTE_DEFINITION_XHTML_Rendering]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_DEFINITION_XHTML_Renderings_instance = make(map[*ATTRIBUTE_DEFINITION_XHTML_Rendering]*ATTRIBUTE_DEFINITION_XHTML_Rendering)
	for instance := range stage.ATTRIBUTE_DEFINITION_XHTML_Renderings {
		_copy := instance.GongCopy().(*ATTRIBUTE_DEFINITION_XHTML_Rendering)
		stage.ATTRIBUTE_DEFINITION_XHTML_Renderings_reference[instance] = _copy
		stage.ATTRIBUTE_DEFINITION_XHTML_Renderings_instance[_copy] = instance
		stage.ATTRIBUTE_DEFINITION_XHTML_Renderings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_VALUE_BOOLEANs_reference = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*ATTRIBUTE_VALUE_BOOLEAN)
	stage.ATTRIBUTE_VALUE_BOOLEANs_referenceOrder = make(map[*ATTRIBUTE_VALUE_BOOLEAN]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_VALUE_BOOLEANs_instance = make(map[*ATTRIBUTE_VALUE_BOOLEAN]*ATTRIBUTE_VALUE_BOOLEAN)
	for instance := range stage.ATTRIBUTE_VALUE_BOOLEANs {
		_copy := instance.GongCopy().(*ATTRIBUTE_VALUE_BOOLEAN)
		stage.ATTRIBUTE_VALUE_BOOLEANs_reference[instance] = _copy
		stage.ATTRIBUTE_VALUE_BOOLEANs_instance[_copy] = instance
		stage.ATTRIBUTE_VALUE_BOOLEANs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_VALUE_DATEs_reference = make(map[*ATTRIBUTE_VALUE_DATE]*ATTRIBUTE_VALUE_DATE)
	stage.ATTRIBUTE_VALUE_DATEs_referenceOrder = make(map[*ATTRIBUTE_VALUE_DATE]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_VALUE_DATEs_instance = make(map[*ATTRIBUTE_VALUE_DATE]*ATTRIBUTE_VALUE_DATE)
	for instance := range stage.ATTRIBUTE_VALUE_DATEs {
		_copy := instance.GongCopy().(*ATTRIBUTE_VALUE_DATE)
		stage.ATTRIBUTE_VALUE_DATEs_reference[instance] = _copy
		stage.ATTRIBUTE_VALUE_DATEs_instance[_copy] = instance
		stage.ATTRIBUTE_VALUE_DATEs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_VALUE_ENUMERATIONs_reference = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*ATTRIBUTE_VALUE_ENUMERATION)
	stage.ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder = make(map[*ATTRIBUTE_VALUE_ENUMERATION]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_VALUE_ENUMERATIONs_instance = make(map[*ATTRIBUTE_VALUE_ENUMERATION]*ATTRIBUTE_VALUE_ENUMERATION)
	for instance := range stage.ATTRIBUTE_VALUE_ENUMERATIONs {
		_copy := instance.GongCopy().(*ATTRIBUTE_VALUE_ENUMERATION)
		stage.ATTRIBUTE_VALUE_ENUMERATIONs_reference[instance] = _copy
		stage.ATTRIBUTE_VALUE_ENUMERATIONs_instance[_copy] = instance
		stage.ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_VALUE_INTEGERs_reference = make(map[*ATTRIBUTE_VALUE_INTEGER]*ATTRIBUTE_VALUE_INTEGER)
	stage.ATTRIBUTE_VALUE_INTEGERs_referenceOrder = make(map[*ATTRIBUTE_VALUE_INTEGER]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_VALUE_INTEGERs_instance = make(map[*ATTRIBUTE_VALUE_INTEGER]*ATTRIBUTE_VALUE_INTEGER)
	for instance := range stage.ATTRIBUTE_VALUE_INTEGERs {
		_copy := instance.GongCopy().(*ATTRIBUTE_VALUE_INTEGER)
		stage.ATTRIBUTE_VALUE_INTEGERs_reference[instance] = _copy
		stage.ATTRIBUTE_VALUE_INTEGERs_instance[_copy] = instance
		stage.ATTRIBUTE_VALUE_INTEGERs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_VALUE_REALs_reference = make(map[*ATTRIBUTE_VALUE_REAL]*ATTRIBUTE_VALUE_REAL)
	stage.ATTRIBUTE_VALUE_REALs_referenceOrder = make(map[*ATTRIBUTE_VALUE_REAL]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_VALUE_REALs_instance = make(map[*ATTRIBUTE_VALUE_REAL]*ATTRIBUTE_VALUE_REAL)
	for instance := range stage.ATTRIBUTE_VALUE_REALs {
		_copy := instance.GongCopy().(*ATTRIBUTE_VALUE_REAL)
		stage.ATTRIBUTE_VALUE_REALs_reference[instance] = _copy
		stage.ATTRIBUTE_VALUE_REALs_instance[_copy] = instance
		stage.ATTRIBUTE_VALUE_REALs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_VALUE_STRINGs_reference = make(map[*ATTRIBUTE_VALUE_STRING]*ATTRIBUTE_VALUE_STRING)
	stage.ATTRIBUTE_VALUE_STRINGs_referenceOrder = make(map[*ATTRIBUTE_VALUE_STRING]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_VALUE_STRINGs_instance = make(map[*ATTRIBUTE_VALUE_STRING]*ATTRIBUTE_VALUE_STRING)
	for instance := range stage.ATTRIBUTE_VALUE_STRINGs {
		_copy := instance.GongCopy().(*ATTRIBUTE_VALUE_STRING)
		stage.ATTRIBUTE_VALUE_STRINGs_reference[instance] = _copy
		stage.ATTRIBUTE_VALUE_STRINGs_instance[_copy] = instance
		stage.ATTRIBUTE_VALUE_STRINGs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ATTRIBUTE_VALUE_XHTMLs_reference = make(map[*ATTRIBUTE_VALUE_XHTML]*ATTRIBUTE_VALUE_XHTML)
	stage.ATTRIBUTE_VALUE_XHTMLs_referenceOrder = make(map[*ATTRIBUTE_VALUE_XHTML]uint) // diff Unstage needs the reference order
	stage.ATTRIBUTE_VALUE_XHTMLs_instance = make(map[*ATTRIBUTE_VALUE_XHTML]*ATTRIBUTE_VALUE_XHTML)
	for instance := range stage.ATTRIBUTE_VALUE_XHTMLs {
		_copy := instance.GongCopy().(*ATTRIBUTE_VALUE_XHTML)
		stage.ATTRIBUTE_VALUE_XHTMLs_reference[instance] = _copy
		stage.ATTRIBUTE_VALUE_XHTMLs_instance[_copy] = instance
		stage.ATTRIBUTE_VALUE_XHTMLs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ALTERNATIVE_IDs_reference = make(map[*A_ALTERNATIVE_ID]*A_ALTERNATIVE_ID)
	stage.A_ALTERNATIVE_IDs_referenceOrder = make(map[*A_ALTERNATIVE_ID]uint) // diff Unstage needs the reference order
	stage.A_ALTERNATIVE_IDs_instance = make(map[*A_ALTERNATIVE_ID]*A_ALTERNATIVE_ID)
	for instance := range stage.A_ALTERNATIVE_IDs {
		_copy := instance.GongCopy().(*A_ALTERNATIVE_ID)
		stage.A_ALTERNATIVE_IDs_reference[instance] = _copy
		stage.A_ALTERNATIVE_IDs_instance[_copy] = instance
		stage.A_ALTERNATIVE_IDs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_reference = make(map[*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]uint) // diff Unstage needs the reference order
	stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_instance = make(map[*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF]*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	for instance := range stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs {
		_copy := instance.GongCopy().(*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
		stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_reference[instance] = _copy
		stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_instance[_copy] = instance
		stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_reference = make(map[*A_ATTRIBUTE_DEFINITION_DATE_REF]*A_ATTRIBUTE_DEFINITION_DATE_REF)
	stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_DATE_REF]uint) // diff Unstage needs the reference order
	stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_instance = make(map[*A_ATTRIBUTE_DEFINITION_DATE_REF]*A_ATTRIBUTE_DEFINITION_DATE_REF)
	for instance := range stage.A_ATTRIBUTE_DEFINITION_DATE_REFs {
		_copy := instance.GongCopy().(*A_ATTRIBUTE_DEFINITION_DATE_REF)
		stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_reference[instance] = _copy
		stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_instance[_copy] = instance
		stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_reference = make(map[*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]uint) // diff Unstage needs the reference order
	stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_instance = make(map[*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF]*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	for instance := range stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs {
		_copy := instance.GongCopy().(*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
		stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_reference[instance] = _copy
		stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_instance[_copy] = instance
		stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_reference = make(map[*A_ATTRIBUTE_DEFINITION_INTEGER_REF]*A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_INTEGER_REF]uint) // diff Unstage needs the reference order
	stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_instance = make(map[*A_ATTRIBUTE_DEFINITION_INTEGER_REF]*A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	for instance := range stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs {
		_copy := instance.GongCopy().(*A_ATTRIBUTE_DEFINITION_INTEGER_REF)
		stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_reference[instance] = _copy
		stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_instance[_copy] = instance
		stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_reference = make(map[*A_ATTRIBUTE_DEFINITION_REAL_REF]*A_ATTRIBUTE_DEFINITION_REAL_REF)
	stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_REAL_REF]uint) // diff Unstage needs the reference order
	stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_instance = make(map[*A_ATTRIBUTE_DEFINITION_REAL_REF]*A_ATTRIBUTE_DEFINITION_REAL_REF)
	for instance := range stage.A_ATTRIBUTE_DEFINITION_REAL_REFs {
		_copy := instance.GongCopy().(*A_ATTRIBUTE_DEFINITION_REAL_REF)
		stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_reference[instance] = _copy
		stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_instance[_copy] = instance
		stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_reference = make(map[*A_ATTRIBUTE_DEFINITION_STRING_REF]*A_ATTRIBUTE_DEFINITION_STRING_REF)
	stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_STRING_REF]uint) // diff Unstage needs the reference order
	stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_instance = make(map[*A_ATTRIBUTE_DEFINITION_STRING_REF]*A_ATTRIBUTE_DEFINITION_STRING_REF)
	for instance := range stage.A_ATTRIBUTE_DEFINITION_STRING_REFs {
		_copy := instance.GongCopy().(*A_ATTRIBUTE_DEFINITION_STRING_REF)
		stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_reference[instance] = _copy
		stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_instance[_copy] = instance
		stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_reference = make(map[*A_ATTRIBUTE_DEFINITION_XHTML_REF]*A_ATTRIBUTE_DEFINITION_XHTML_REF)
	stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_referenceOrder = make(map[*A_ATTRIBUTE_DEFINITION_XHTML_REF]uint) // diff Unstage needs the reference order
	stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_instance = make(map[*A_ATTRIBUTE_DEFINITION_XHTML_REF]*A_ATTRIBUTE_DEFINITION_XHTML_REF)
	for instance := range stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs {
		_copy := instance.GongCopy().(*A_ATTRIBUTE_DEFINITION_XHTML_REF)
		stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_reference[instance] = _copy
		stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_instance[_copy] = instance
		stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_VALUE_BOOLEANs_reference = make(map[*A_ATTRIBUTE_VALUE_BOOLEAN]*A_ATTRIBUTE_VALUE_BOOLEAN)
	stage.A_ATTRIBUTE_VALUE_BOOLEANs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_BOOLEAN]uint) // diff Unstage needs the reference order
	stage.A_ATTRIBUTE_VALUE_BOOLEANs_instance = make(map[*A_ATTRIBUTE_VALUE_BOOLEAN]*A_ATTRIBUTE_VALUE_BOOLEAN)
	for instance := range stage.A_ATTRIBUTE_VALUE_BOOLEANs {
		_copy := instance.GongCopy().(*A_ATTRIBUTE_VALUE_BOOLEAN)
		stage.A_ATTRIBUTE_VALUE_BOOLEANs_reference[instance] = _copy
		stage.A_ATTRIBUTE_VALUE_BOOLEANs_instance[_copy] = instance
		stage.A_ATTRIBUTE_VALUE_BOOLEANs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_VALUE_DATEs_reference = make(map[*A_ATTRIBUTE_VALUE_DATE]*A_ATTRIBUTE_VALUE_DATE)
	stage.A_ATTRIBUTE_VALUE_DATEs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_DATE]uint) // diff Unstage needs the reference order
	stage.A_ATTRIBUTE_VALUE_DATEs_instance = make(map[*A_ATTRIBUTE_VALUE_DATE]*A_ATTRIBUTE_VALUE_DATE)
	for instance := range stage.A_ATTRIBUTE_VALUE_DATEs {
		_copy := instance.GongCopy().(*A_ATTRIBUTE_VALUE_DATE)
		stage.A_ATTRIBUTE_VALUE_DATEs_reference[instance] = _copy
		stage.A_ATTRIBUTE_VALUE_DATEs_instance[_copy] = instance
		stage.A_ATTRIBUTE_VALUE_DATEs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_reference = make(map[*A_ATTRIBUTE_VALUE_ENUMERATION]*A_ATTRIBUTE_VALUE_ENUMERATION)
	stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_ENUMERATION]uint) // diff Unstage needs the reference order
	stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_instance = make(map[*A_ATTRIBUTE_VALUE_ENUMERATION]*A_ATTRIBUTE_VALUE_ENUMERATION)
	for instance := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
		_copy := instance.GongCopy().(*A_ATTRIBUTE_VALUE_ENUMERATION)
		stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_reference[instance] = _copy
		stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_instance[_copy] = instance
		stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_VALUE_INTEGERs_reference = make(map[*A_ATTRIBUTE_VALUE_INTEGER]*A_ATTRIBUTE_VALUE_INTEGER)
	stage.A_ATTRIBUTE_VALUE_INTEGERs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_INTEGER]uint) // diff Unstage needs the reference order
	stage.A_ATTRIBUTE_VALUE_INTEGERs_instance = make(map[*A_ATTRIBUTE_VALUE_INTEGER]*A_ATTRIBUTE_VALUE_INTEGER)
	for instance := range stage.A_ATTRIBUTE_VALUE_INTEGERs {
		_copy := instance.GongCopy().(*A_ATTRIBUTE_VALUE_INTEGER)
		stage.A_ATTRIBUTE_VALUE_INTEGERs_reference[instance] = _copy
		stage.A_ATTRIBUTE_VALUE_INTEGERs_instance[_copy] = instance
		stage.A_ATTRIBUTE_VALUE_INTEGERs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_VALUE_REALs_reference = make(map[*A_ATTRIBUTE_VALUE_REAL]*A_ATTRIBUTE_VALUE_REAL)
	stage.A_ATTRIBUTE_VALUE_REALs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_REAL]uint) // diff Unstage needs the reference order
	stage.A_ATTRIBUTE_VALUE_REALs_instance = make(map[*A_ATTRIBUTE_VALUE_REAL]*A_ATTRIBUTE_VALUE_REAL)
	for instance := range stage.A_ATTRIBUTE_VALUE_REALs {
		_copy := instance.GongCopy().(*A_ATTRIBUTE_VALUE_REAL)
		stage.A_ATTRIBUTE_VALUE_REALs_reference[instance] = _copy
		stage.A_ATTRIBUTE_VALUE_REALs_instance[_copy] = instance
		stage.A_ATTRIBUTE_VALUE_REALs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_VALUE_STRINGs_reference = make(map[*A_ATTRIBUTE_VALUE_STRING]*A_ATTRIBUTE_VALUE_STRING)
	stage.A_ATTRIBUTE_VALUE_STRINGs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_STRING]uint) // diff Unstage needs the reference order
	stage.A_ATTRIBUTE_VALUE_STRINGs_instance = make(map[*A_ATTRIBUTE_VALUE_STRING]*A_ATTRIBUTE_VALUE_STRING)
	for instance := range stage.A_ATTRIBUTE_VALUE_STRINGs {
		_copy := instance.GongCopy().(*A_ATTRIBUTE_VALUE_STRING)
		stage.A_ATTRIBUTE_VALUE_STRINGs_reference[instance] = _copy
		stage.A_ATTRIBUTE_VALUE_STRINGs_instance[_copy] = instance
		stage.A_ATTRIBUTE_VALUE_STRINGs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_VALUE_XHTMLs_reference = make(map[*A_ATTRIBUTE_VALUE_XHTML]*A_ATTRIBUTE_VALUE_XHTML)
	stage.A_ATTRIBUTE_VALUE_XHTMLs_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_XHTML]uint) // diff Unstage needs the reference order
	stage.A_ATTRIBUTE_VALUE_XHTMLs_instance = make(map[*A_ATTRIBUTE_VALUE_XHTML]*A_ATTRIBUTE_VALUE_XHTML)
	for instance := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
		_copy := instance.GongCopy().(*A_ATTRIBUTE_VALUE_XHTML)
		stage.A_ATTRIBUTE_VALUE_XHTMLs_reference[instance] = _copy
		stage.A_ATTRIBUTE_VALUE_XHTMLs_instance[_copy] = instance
		stage.A_ATTRIBUTE_VALUE_XHTMLs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ATTRIBUTE_VALUE_XHTML_1s_reference = make(map[*A_ATTRIBUTE_VALUE_XHTML_1]*A_ATTRIBUTE_VALUE_XHTML_1)
	stage.A_ATTRIBUTE_VALUE_XHTML_1s_referenceOrder = make(map[*A_ATTRIBUTE_VALUE_XHTML_1]uint) // diff Unstage needs the reference order
	stage.A_ATTRIBUTE_VALUE_XHTML_1s_instance = make(map[*A_ATTRIBUTE_VALUE_XHTML_1]*A_ATTRIBUTE_VALUE_XHTML_1)
	for instance := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		_copy := instance.GongCopy().(*A_ATTRIBUTE_VALUE_XHTML_1)
		stage.A_ATTRIBUTE_VALUE_XHTML_1s_reference[instance] = _copy
		stage.A_ATTRIBUTE_VALUE_XHTML_1s_instance[_copy] = instance
		stage.A_ATTRIBUTE_VALUE_XHTML_1s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_CHILDRENs_reference = make(map[*A_CHILDREN]*A_CHILDREN)
	stage.A_CHILDRENs_referenceOrder = make(map[*A_CHILDREN]uint) // diff Unstage needs the reference order
	stage.A_CHILDRENs_instance = make(map[*A_CHILDREN]*A_CHILDREN)
	for instance := range stage.A_CHILDRENs {
		_copy := instance.GongCopy().(*A_CHILDREN)
		stage.A_CHILDRENs_reference[instance] = _copy
		stage.A_CHILDRENs_instance[_copy] = instance
		stage.A_CHILDRENs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_CORE_CONTENTs_reference = make(map[*A_CORE_CONTENT]*A_CORE_CONTENT)
	stage.A_CORE_CONTENTs_referenceOrder = make(map[*A_CORE_CONTENT]uint) // diff Unstage needs the reference order
	stage.A_CORE_CONTENTs_instance = make(map[*A_CORE_CONTENT]*A_CORE_CONTENT)
	for instance := range stage.A_CORE_CONTENTs {
		_copy := instance.GongCopy().(*A_CORE_CONTENT)
		stage.A_CORE_CONTENTs_reference[instance] = _copy
		stage.A_CORE_CONTENTs_instance[_copy] = instance
		stage.A_CORE_CONTENTs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_DATATYPESs_reference = make(map[*A_DATATYPES]*A_DATATYPES)
	stage.A_DATATYPESs_referenceOrder = make(map[*A_DATATYPES]uint) // diff Unstage needs the reference order
	stage.A_DATATYPESs_instance = make(map[*A_DATATYPES]*A_DATATYPES)
	for instance := range stage.A_DATATYPESs {
		_copy := instance.GongCopy().(*A_DATATYPES)
		stage.A_DATATYPESs_reference[instance] = _copy
		stage.A_DATATYPESs_instance[_copy] = instance
		stage.A_DATATYPESs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_reference = make(map[*A_DATATYPE_DEFINITION_BOOLEAN_REF]*A_DATATYPE_DEFINITION_BOOLEAN_REF)
	stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_BOOLEAN_REF]uint) // diff Unstage needs the reference order
	stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_instance = make(map[*A_DATATYPE_DEFINITION_BOOLEAN_REF]*A_DATATYPE_DEFINITION_BOOLEAN_REF)
	for instance := range stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs {
		_copy := instance.GongCopy().(*A_DATATYPE_DEFINITION_BOOLEAN_REF)
		stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_reference[instance] = _copy
		stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_instance[_copy] = instance
		stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_DATATYPE_DEFINITION_DATE_REFs_reference = make(map[*A_DATATYPE_DEFINITION_DATE_REF]*A_DATATYPE_DEFINITION_DATE_REF)
	stage.A_DATATYPE_DEFINITION_DATE_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_DATE_REF]uint) // diff Unstage needs the reference order
	stage.A_DATATYPE_DEFINITION_DATE_REFs_instance = make(map[*A_DATATYPE_DEFINITION_DATE_REF]*A_DATATYPE_DEFINITION_DATE_REF)
	for instance := range stage.A_DATATYPE_DEFINITION_DATE_REFs {
		_copy := instance.GongCopy().(*A_DATATYPE_DEFINITION_DATE_REF)
		stage.A_DATATYPE_DEFINITION_DATE_REFs_reference[instance] = _copy
		stage.A_DATATYPE_DEFINITION_DATE_REFs_instance[_copy] = instance
		stage.A_DATATYPE_DEFINITION_DATE_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_reference = make(map[*A_DATATYPE_DEFINITION_ENUMERATION_REF]*A_DATATYPE_DEFINITION_ENUMERATION_REF)
	stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_ENUMERATION_REF]uint) // diff Unstage needs the reference order
	stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_instance = make(map[*A_DATATYPE_DEFINITION_ENUMERATION_REF]*A_DATATYPE_DEFINITION_ENUMERATION_REF)
	for instance := range stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs {
		_copy := instance.GongCopy().(*A_DATATYPE_DEFINITION_ENUMERATION_REF)
		stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_reference[instance] = _copy
		stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_instance[_copy] = instance
		stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_DATATYPE_DEFINITION_INTEGER_REFs_reference = make(map[*A_DATATYPE_DEFINITION_INTEGER_REF]*A_DATATYPE_DEFINITION_INTEGER_REF)
	stage.A_DATATYPE_DEFINITION_INTEGER_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_INTEGER_REF]uint) // diff Unstage needs the reference order
	stage.A_DATATYPE_DEFINITION_INTEGER_REFs_instance = make(map[*A_DATATYPE_DEFINITION_INTEGER_REF]*A_DATATYPE_DEFINITION_INTEGER_REF)
	for instance := range stage.A_DATATYPE_DEFINITION_INTEGER_REFs {
		_copy := instance.GongCopy().(*A_DATATYPE_DEFINITION_INTEGER_REF)
		stage.A_DATATYPE_DEFINITION_INTEGER_REFs_reference[instance] = _copy
		stage.A_DATATYPE_DEFINITION_INTEGER_REFs_instance[_copy] = instance
		stage.A_DATATYPE_DEFINITION_INTEGER_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_DATATYPE_DEFINITION_REAL_REFs_reference = make(map[*A_DATATYPE_DEFINITION_REAL_REF]*A_DATATYPE_DEFINITION_REAL_REF)
	stage.A_DATATYPE_DEFINITION_REAL_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_REAL_REF]uint) // diff Unstage needs the reference order
	stage.A_DATATYPE_DEFINITION_REAL_REFs_instance = make(map[*A_DATATYPE_DEFINITION_REAL_REF]*A_DATATYPE_DEFINITION_REAL_REF)
	for instance := range stage.A_DATATYPE_DEFINITION_REAL_REFs {
		_copy := instance.GongCopy().(*A_DATATYPE_DEFINITION_REAL_REF)
		stage.A_DATATYPE_DEFINITION_REAL_REFs_reference[instance] = _copy
		stage.A_DATATYPE_DEFINITION_REAL_REFs_instance[_copy] = instance
		stage.A_DATATYPE_DEFINITION_REAL_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_DATATYPE_DEFINITION_STRING_REFs_reference = make(map[*A_DATATYPE_DEFINITION_STRING_REF]*A_DATATYPE_DEFINITION_STRING_REF)
	stage.A_DATATYPE_DEFINITION_STRING_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_STRING_REF]uint) // diff Unstage needs the reference order
	stage.A_DATATYPE_DEFINITION_STRING_REFs_instance = make(map[*A_DATATYPE_DEFINITION_STRING_REF]*A_DATATYPE_DEFINITION_STRING_REF)
	for instance := range stage.A_DATATYPE_DEFINITION_STRING_REFs {
		_copy := instance.GongCopy().(*A_DATATYPE_DEFINITION_STRING_REF)
		stage.A_DATATYPE_DEFINITION_STRING_REFs_reference[instance] = _copy
		stage.A_DATATYPE_DEFINITION_STRING_REFs_instance[_copy] = instance
		stage.A_DATATYPE_DEFINITION_STRING_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_DATATYPE_DEFINITION_XHTML_REFs_reference = make(map[*A_DATATYPE_DEFINITION_XHTML_REF]*A_DATATYPE_DEFINITION_XHTML_REF)
	stage.A_DATATYPE_DEFINITION_XHTML_REFs_referenceOrder = make(map[*A_DATATYPE_DEFINITION_XHTML_REF]uint) // diff Unstage needs the reference order
	stage.A_DATATYPE_DEFINITION_XHTML_REFs_instance = make(map[*A_DATATYPE_DEFINITION_XHTML_REF]*A_DATATYPE_DEFINITION_XHTML_REF)
	for instance := range stage.A_DATATYPE_DEFINITION_XHTML_REFs {
		_copy := instance.GongCopy().(*A_DATATYPE_DEFINITION_XHTML_REF)
		stage.A_DATATYPE_DEFINITION_XHTML_REFs_reference[instance] = _copy
		stage.A_DATATYPE_DEFINITION_XHTML_REFs_instance[_copy] = instance
		stage.A_DATATYPE_DEFINITION_XHTML_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_EDITABLE_ATTSs_reference = make(map[*A_EDITABLE_ATTS]*A_EDITABLE_ATTS)
	stage.A_EDITABLE_ATTSs_referenceOrder = make(map[*A_EDITABLE_ATTS]uint) // diff Unstage needs the reference order
	stage.A_EDITABLE_ATTSs_instance = make(map[*A_EDITABLE_ATTS]*A_EDITABLE_ATTS)
	for instance := range stage.A_EDITABLE_ATTSs {
		_copy := instance.GongCopy().(*A_EDITABLE_ATTS)
		stage.A_EDITABLE_ATTSs_reference[instance] = _copy
		stage.A_EDITABLE_ATTSs_instance[_copy] = instance
		stage.A_EDITABLE_ATTSs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_ENUM_VALUE_REFs_reference = make(map[*A_ENUM_VALUE_REF]*A_ENUM_VALUE_REF)
	stage.A_ENUM_VALUE_REFs_referenceOrder = make(map[*A_ENUM_VALUE_REF]uint) // diff Unstage needs the reference order
	stage.A_ENUM_VALUE_REFs_instance = make(map[*A_ENUM_VALUE_REF]*A_ENUM_VALUE_REF)
	for instance := range stage.A_ENUM_VALUE_REFs {
		_copy := instance.GongCopy().(*A_ENUM_VALUE_REF)
		stage.A_ENUM_VALUE_REFs_reference[instance] = _copy
		stage.A_ENUM_VALUE_REFs_instance[_copy] = instance
		stage.A_ENUM_VALUE_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_OBJECTs_reference = make(map[*A_OBJECT]*A_OBJECT)
	stage.A_OBJECTs_referenceOrder = make(map[*A_OBJECT]uint) // diff Unstage needs the reference order
	stage.A_OBJECTs_instance = make(map[*A_OBJECT]*A_OBJECT)
	for instance := range stage.A_OBJECTs {
		_copy := instance.GongCopy().(*A_OBJECT)
		stage.A_OBJECTs_reference[instance] = _copy
		stage.A_OBJECTs_instance[_copy] = instance
		stage.A_OBJECTs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_PROPERTIESs_reference = make(map[*A_PROPERTIES]*A_PROPERTIES)
	stage.A_PROPERTIESs_referenceOrder = make(map[*A_PROPERTIES]uint) // diff Unstage needs the reference order
	stage.A_PROPERTIESs_instance = make(map[*A_PROPERTIES]*A_PROPERTIES)
	for instance := range stage.A_PROPERTIESs {
		_copy := instance.GongCopy().(*A_PROPERTIES)
		stage.A_PROPERTIESs_reference[instance] = _copy
		stage.A_PROPERTIESs_instance[_copy] = instance
		stage.A_PROPERTIESs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_RELATION_GROUP_TYPE_REFs_reference = make(map[*A_RELATION_GROUP_TYPE_REF]*A_RELATION_GROUP_TYPE_REF)
	stage.A_RELATION_GROUP_TYPE_REFs_referenceOrder = make(map[*A_RELATION_GROUP_TYPE_REF]uint) // diff Unstage needs the reference order
	stage.A_RELATION_GROUP_TYPE_REFs_instance = make(map[*A_RELATION_GROUP_TYPE_REF]*A_RELATION_GROUP_TYPE_REF)
	for instance := range stage.A_RELATION_GROUP_TYPE_REFs {
		_copy := instance.GongCopy().(*A_RELATION_GROUP_TYPE_REF)
		stage.A_RELATION_GROUP_TYPE_REFs_reference[instance] = _copy
		stage.A_RELATION_GROUP_TYPE_REFs_instance[_copy] = instance
		stage.A_RELATION_GROUP_TYPE_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_SOURCE_1s_reference = make(map[*A_SOURCE_1]*A_SOURCE_1)
	stage.A_SOURCE_1s_referenceOrder = make(map[*A_SOURCE_1]uint) // diff Unstage needs the reference order
	stage.A_SOURCE_1s_instance = make(map[*A_SOURCE_1]*A_SOURCE_1)
	for instance := range stage.A_SOURCE_1s {
		_copy := instance.GongCopy().(*A_SOURCE_1)
		stage.A_SOURCE_1s_reference[instance] = _copy
		stage.A_SOURCE_1s_instance[_copy] = instance
		stage.A_SOURCE_1s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_SOURCE_SPECIFICATION_1s_reference = make(map[*A_SOURCE_SPECIFICATION_1]*A_SOURCE_SPECIFICATION_1)
	stage.A_SOURCE_SPECIFICATION_1s_referenceOrder = make(map[*A_SOURCE_SPECIFICATION_1]uint) // diff Unstage needs the reference order
	stage.A_SOURCE_SPECIFICATION_1s_instance = make(map[*A_SOURCE_SPECIFICATION_1]*A_SOURCE_SPECIFICATION_1)
	for instance := range stage.A_SOURCE_SPECIFICATION_1s {
		_copy := instance.GongCopy().(*A_SOURCE_SPECIFICATION_1)
		stage.A_SOURCE_SPECIFICATION_1s_reference[instance] = _copy
		stage.A_SOURCE_SPECIFICATION_1s_instance[_copy] = instance
		stage.A_SOURCE_SPECIFICATION_1s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_SPECIFICATIONSs_reference = make(map[*A_SPECIFICATIONS]*A_SPECIFICATIONS)
	stage.A_SPECIFICATIONSs_referenceOrder = make(map[*A_SPECIFICATIONS]uint) // diff Unstage needs the reference order
	stage.A_SPECIFICATIONSs_instance = make(map[*A_SPECIFICATIONS]*A_SPECIFICATIONS)
	for instance := range stage.A_SPECIFICATIONSs {
		_copy := instance.GongCopy().(*A_SPECIFICATIONS)
		stage.A_SPECIFICATIONSs_reference[instance] = _copy
		stage.A_SPECIFICATIONSs_instance[_copy] = instance
		stage.A_SPECIFICATIONSs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_SPECIFICATION_TYPE_REFs_reference = make(map[*A_SPECIFICATION_TYPE_REF]*A_SPECIFICATION_TYPE_REF)
	stage.A_SPECIFICATION_TYPE_REFs_referenceOrder = make(map[*A_SPECIFICATION_TYPE_REF]uint) // diff Unstage needs the reference order
	stage.A_SPECIFICATION_TYPE_REFs_instance = make(map[*A_SPECIFICATION_TYPE_REF]*A_SPECIFICATION_TYPE_REF)
	for instance := range stage.A_SPECIFICATION_TYPE_REFs {
		_copy := instance.GongCopy().(*A_SPECIFICATION_TYPE_REF)
		stage.A_SPECIFICATION_TYPE_REFs_reference[instance] = _copy
		stage.A_SPECIFICATION_TYPE_REFs_instance[_copy] = instance
		stage.A_SPECIFICATION_TYPE_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_SPECIFIED_VALUESs_reference = make(map[*A_SPECIFIED_VALUES]*A_SPECIFIED_VALUES)
	stage.A_SPECIFIED_VALUESs_referenceOrder = make(map[*A_SPECIFIED_VALUES]uint) // diff Unstage needs the reference order
	stage.A_SPECIFIED_VALUESs_instance = make(map[*A_SPECIFIED_VALUES]*A_SPECIFIED_VALUES)
	for instance := range stage.A_SPECIFIED_VALUESs {
		_copy := instance.GongCopy().(*A_SPECIFIED_VALUES)
		stage.A_SPECIFIED_VALUESs_reference[instance] = _copy
		stage.A_SPECIFIED_VALUESs_instance[_copy] = instance
		stage.A_SPECIFIED_VALUESs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_SPEC_ATTRIBUTESs_reference = make(map[*A_SPEC_ATTRIBUTES]*A_SPEC_ATTRIBUTES)
	stage.A_SPEC_ATTRIBUTESs_referenceOrder = make(map[*A_SPEC_ATTRIBUTES]uint) // diff Unstage needs the reference order
	stage.A_SPEC_ATTRIBUTESs_instance = make(map[*A_SPEC_ATTRIBUTES]*A_SPEC_ATTRIBUTES)
	for instance := range stage.A_SPEC_ATTRIBUTESs {
		_copy := instance.GongCopy().(*A_SPEC_ATTRIBUTES)
		stage.A_SPEC_ATTRIBUTESs_reference[instance] = _copy
		stage.A_SPEC_ATTRIBUTESs_instance[_copy] = instance
		stage.A_SPEC_ATTRIBUTESs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_SPEC_OBJECTSs_reference = make(map[*A_SPEC_OBJECTS]*A_SPEC_OBJECTS)
	stage.A_SPEC_OBJECTSs_referenceOrder = make(map[*A_SPEC_OBJECTS]uint) // diff Unstage needs the reference order
	stage.A_SPEC_OBJECTSs_instance = make(map[*A_SPEC_OBJECTS]*A_SPEC_OBJECTS)
	for instance := range stage.A_SPEC_OBJECTSs {
		_copy := instance.GongCopy().(*A_SPEC_OBJECTS)
		stage.A_SPEC_OBJECTSs_reference[instance] = _copy
		stage.A_SPEC_OBJECTSs_instance[_copy] = instance
		stage.A_SPEC_OBJECTSs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_SPEC_OBJECT_TYPE_REFs_reference = make(map[*A_SPEC_OBJECT_TYPE_REF]*A_SPEC_OBJECT_TYPE_REF)
	stage.A_SPEC_OBJECT_TYPE_REFs_referenceOrder = make(map[*A_SPEC_OBJECT_TYPE_REF]uint) // diff Unstage needs the reference order
	stage.A_SPEC_OBJECT_TYPE_REFs_instance = make(map[*A_SPEC_OBJECT_TYPE_REF]*A_SPEC_OBJECT_TYPE_REF)
	for instance := range stage.A_SPEC_OBJECT_TYPE_REFs {
		_copy := instance.GongCopy().(*A_SPEC_OBJECT_TYPE_REF)
		stage.A_SPEC_OBJECT_TYPE_REFs_reference[instance] = _copy
		stage.A_SPEC_OBJECT_TYPE_REFs_instance[_copy] = instance
		stage.A_SPEC_OBJECT_TYPE_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_SPEC_RELATIONSs_reference = make(map[*A_SPEC_RELATIONS]*A_SPEC_RELATIONS)
	stage.A_SPEC_RELATIONSs_referenceOrder = make(map[*A_SPEC_RELATIONS]uint) // diff Unstage needs the reference order
	stage.A_SPEC_RELATIONSs_instance = make(map[*A_SPEC_RELATIONS]*A_SPEC_RELATIONS)
	for instance := range stage.A_SPEC_RELATIONSs {
		_copy := instance.GongCopy().(*A_SPEC_RELATIONS)
		stage.A_SPEC_RELATIONSs_reference[instance] = _copy
		stage.A_SPEC_RELATIONSs_instance[_copy] = instance
		stage.A_SPEC_RELATIONSs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_SPEC_RELATION_GROUPSs_reference = make(map[*A_SPEC_RELATION_GROUPS]*A_SPEC_RELATION_GROUPS)
	stage.A_SPEC_RELATION_GROUPSs_referenceOrder = make(map[*A_SPEC_RELATION_GROUPS]uint) // diff Unstage needs the reference order
	stage.A_SPEC_RELATION_GROUPSs_instance = make(map[*A_SPEC_RELATION_GROUPS]*A_SPEC_RELATION_GROUPS)
	for instance := range stage.A_SPEC_RELATION_GROUPSs {
		_copy := instance.GongCopy().(*A_SPEC_RELATION_GROUPS)
		stage.A_SPEC_RELATION_GROUPSs_reference[instance] = _copy
		stage.A_SPEC_RELATION_GROUPSs_instance[_copy] = instance
		stage.A_SPEC_RELATION_GROUPSs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_SPEC_RELATION_REFs_reference = make(map[*A_SPEC_RELATION_REF]*A_SPEC_RELATION_REF)
	stage.A_SPEC_RELATION_REFs_referenceOrder = make(map[*A_SPEC_RELATION_REF]uint) // diff Unstage needs the reference order
	stage.A_SPEC_RELATION_REFs_instance = make(map[*A_SPEC_RELATION_REF]*A_SPEC_RELATION_REF)
	for instance := range stage.A_SPEC_RELATION_REFs {
		_copy := instance.GongCopy().(*A_SPEC_RELATION_REF)
		stage.A_SPEC_RELATION_REFs_reference[instance] = _copy
		stage.A_SPEC_RELATION_REFs_instance[_copy] = instance
		stage.A_SPEC_RELATION_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_SPEC_RELATION_TYPE_REFs_reference = make(map[*A_SPEC_RELATION_TYPE_REF]*A_SPEC_RELATION_TYPE_REF)
	stage.A_SPEC_RELATION_TYPE_REFs_referenceOrder = make(map[*A_SPEC_RELATION_TYPE_REF]uint) // diff Unstage needs the reference order
	stage.A_SPEC_RELATION_TYPE_REFs_instance = make(map[*A_SPEC_RELATION_TYPE_REF]*A_SPEC_RELATION_TYPE_REF)
	for instance := range stage.A_SPEC_RELATION_TYPE_REFs {
		_copy := instance.GongCopy().(*A_SPEC_RELATION_TYPE_REF)
		stage.A_SPEC_RELATION_TYPE_REFs_reference[instance] = _copy
		stage.A_SPEC_RELATION_TYPE_REFs_instance[_copy] = instance
		stage.A_SPEC_RELATION_TYPE_REFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_SPEC_TYPESs_reference = make(map[*A_SPEC_TYPES]*A_SPEC_TYPES)
	stage.A_SPEC_TYPESs_referenceOrder = make(map[*A_SPEC_TYPES]uint) // diff Unstage needs the reference order
	stage.A_SPEC_TYPESs_instance = make(map[*A_SPEC_TYPES]*A_SPEC_TYPES)
	for instance := range stage.A_SPEC_TYPESs {
		_copy := instance.GongCopy().(*A_SPEC_TYPES)
		stage.A_SPEC_TYPESs_reference[instance] = _copy
		stage.A_SPEC_TYPESs_instance[_copy] = instance
		stage.A_SPEC_TYPESs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_THE_HEADERs_reference = make(map[*A_THE_HEADER]*A_THE_HEADER)
	stage.A_THE_HEADERs_referenceOrder = make(map[*A_THE_HEADER]uint) // diff Unstage needs the reference order
	stage.A_THE_HEADERs_instance = make(map[*A_THE_HEADER]*A_THE_HEADER)
	for instance := range stage.A_THE_HEADERs {
		_copy := instance.GongCopy().(*A_THE_HEADER)
		stage.A_THE_HEADERs_reference[instance] = _copy
		stage.A_THE_HEADERs_instance[_copy] = instance
		stage.A_THE_HEADERs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_TOOL_EXTENSIONSs_reference = make(map[*A_TOOL_EXTENSIONS]*A_TOOL_EXTENSIONS)
	stage.A_TOOL_EXTENSIONSs_referenceOrder = make(map[*A_TOOL_EXTENSIONS]uint) // diff Unstage needs the reference order
	stage.A_TOOL_EXTENSIONSs_instance = make(map[*A_TOOL_EXTENSIONS]*A_TOOL_EXTENSIONS)
	for instance := range stage.A_TOOL_EXTENSIONSs {
		_copy := instance.GongCopy().(*A_TOOL_EXTENSIONS)
		stage.A_TOOL_EXTENSIONSs_reference[instance] = _copy
		stage.A_TOOL_EXTENSIONSs_instance[_copy] = instance
		stage.A_TOOL_EXTENSIONSs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DATATYPE_DEFINITION_BOOLEANs_reference = make(map[*DATATYPE_DEFINITION_BOOLEAN]*DATATYPE_DEFINITION_BOOLEAN)
	stage.DATATYPE_DEFINITION_BOOLEANs_referenceOrder = make(map[*DATATYPE_DEFINITION_BOOLEAN]uint) // diff Unstage needs the reference order
	stage.DATATYPE_DEFINITION_BOOLEANs_instance = make(map[*DATATYPE_DEFINITION_BOOLEAN]*DATATYPE_DEFINITION_BOOLEAN)
	for instance := range stage.DATATYPE_DEFINITION_BOOLEANs {
		_copy := instance.GongCopy().(*DATATYPE_DEFINITION_BOOLEAN)
		stage.DATATYPE_DEFINITION_BOOLEANs_reference[instance] = _copy
		stage.DATATYPE_DEFINITION_BOOLEANs_instance[_copy] = instance
		stage.DATATYPE_DEFINITION_BOOLEANs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DATATYPE_DEFINITION_DATEs_reference = make(map[*DATATYPE_DEFINITION_DATE]*DATATYPE_DEFINITION_DATE)
	stage.DATATYPE_DEFINITION_DATEs_referenceOrder = make(map[*DATATYPE_DEFINITION_DATE]uint) // diff Unstage needs the reference order
	stage.DATATYPE_DEFINITION_DATEs_instance = make(map[*DATATYPE_DEFINITION_DATE]*DATATYPE_DEFINITION_DATE)
	for instance := range stage.DATATYPE_DEFINITION_DATEs {
		_copy := instance.GongCopy().(*DATATYPE_DEFINITION_DATE)
		stage.DATATYPE_DEFINITION_DATEs_reference[instance] = _copy
		stage.DATATYPE_DEFINITION_DATEs_instance[_copy] = instance
		stage.DATATYPE_DEFINITION_DATEs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DATATYPE_DEFINITION_ENUMERATIONs_reference = make(map[*DATATYPE_DEFINITION_ENUMERATION]*DATATYPE_DEFINITION_ENUMERATION)
	stage.DATATYPE_DEFINITION_ENUMERATIONs_referenceOrder = make(map[*DATATYPE_DEFINITION_ENUMERATION]uint) // diff Unstage needs the reference order
	stage.DATATYPE_DEFINITION_ENUMERATIONs_instance = make(map[*DATATYPE_DEFINITION_ENUMERATION]*DATATYPE_DEFINITION_ENUMERATION)
	for instance := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		_copy := instance.GongCopy().(*DATATYPE_DEFINITION_ENUMERATION)
		stage.DATATYPE_DEFINITION_ENUMERATIONs_reference[instance] = _copy
		stage.DATATYPE_DEFINITION_ENUMERATIONs_instance[_copy] = instance
		stage.DATATYPE_DEFINITION_ENUMERATIONs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DATATYPE_DEFINITION_INTEGERs_reference = make(map[*DATATYPE_DEFINITION_INTEGER]*DATATYPE_DEFINITION_INTEGER)
	stage.DATATYPE_DEFINITION_INTEGERs_referenceOrder = make(map[*DATATYPE_DEFINITION_INTEGER]uint) // diff Unstage needs the reference order
	stage.DATATYPE_DEFINITION_INTEGERs_instance = make(map[*DATATYPE_DEFINITION_INTEGER]*DATATYPE_DEFINITION_INTEGER)
	for instance := range stage.DATATYPE_DEFINITION_INTEGERs {
		_copy := instance.GongCopy().(*DATATYPE_DEFINITION_INTEGER)
		stage.DATATYPE_DEFINITION_INTEGERs_reference[instance] = _copy
		stage.DATATYPE_DEFINITION_INTEGERs_instance[_copy] = instance
		stage.DATATYPE_DEFINITION_INTEGERs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DATATYPE_DEFINITION_REALs_reference = make(map[*DATATYPE_DEFINITION_REAL]*DATATYPE_DEFINITION_REAL)
	stage.DATATYPE_DEFINITION_REALs_referenceOrder = make(map[*DATATYPE_DEFINITION_REAL]uint) // diff Unstage needs the reference order
	stage.DATATYPE_DEFINITION_REALs_instance = make(map[*DATATYPE_DEFINITION_REAL]*DATATYPE_DEFINITION_REAL)
	for instance := range stage.DATATYPE_DEFINITION_REALs {
		_copy := instance.GongCopy().(*DATATYPE_DEFINITION_REAL)
		stage.DATATYPE_DEFINITION_REALs_reference[instance] = _copy
		stage.DATATYPE_DEFINITION_REALs_instance[_copy] = instance
		stage.DATATYPE_DEFINITION_REALs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DATATYPE_DEFINITION_STRINGs_reference = make(map[*DATATYPE_DEFINITION_STRING]*DATATYPE_DEFINITION_STRING)
	stage.DATATYPE_DEFINITION_STRINGs_referenceOrder = make(map[*DATATYPE_DEFINITION_STRING]uint) // diff Unstage needs the reference order
	stage.DATATYPE_DEFINITION_STRINGs_instance = make(map[*DATATYPE_DEFINITION_STRING]*DATATYPE_DEFINITION_STRING)
	for instance := range stage.DATATYPE_DEFINITION_STRINGs {
		_copy := instance.GongCopy().(*DATATYPE_DEFINITION_STRING)
		stage.DATATYPE_DEFINITION_STRINGs_reference[instance] = _copy
		stage.DATATYPE_DEFINITION_STRINGs_instance[_copy] = instance
		stage.DATATYPE_DEFINITION_STRINGs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.DATATYPE_DEFINITION_XHTMLs_reference = make(map[*DATATYPE_DEFINITION_XHTML]*DATATYPE_DEFINITION_XHTML)
	stage.DATATYPE_DEFINITION_XHTMLs_referenceOrder = make(map[*DATATYPE_DEFINITION_XHTML]uint) // diff Unstage needs the reference order
	stage.DATATYPE_DEFINITION_XHTMLs_instance = make(map[*DATATYPE_DEFINITION_XHTML]*DATATYPE_DEFINITION_XHTML)
	for instance := range stage.DATATYPE_DEFINITION_XHTMLs {
		_copy := instance.GongCopy().(*DATATYPE_DEFINITION_XHTML)
		stage.DATATYPE_DEFINITION_XHTMLs_reference[instance] = _copy
		stage.DATATYPE_DEFINITION_XHTMLs_instance[_copy] = instance
		stage.DATATYPE_DEFINITION_XHTMLs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.EMBEDDED_VALUEs_reference = make(map[*EMBEDDED_VALUE]*EMBEDDED_VALUE)
	stage.EMBEDDED_VALUEs_referenceOrder = make(map[*EMBEDDED_VALUE]uint) // diff Unstage needs the reference order
	stage.EMBEDDED_VALUEs_instance = make(map[*EMBEDDED_VALUE]*EMBEDDED_VALUE)
	for instance := range stage.EMBEDDED_VALUEs {
		_copy := instance.GongCopy().(*EMBEDDED_VALUE)
		stage.EMBEDDED_VALUEs_reference[instance] = _copy
		stage.EMBEDDED_VALUEs_instance[_copy] = instance
		stage.EMBEDDED_VALUEs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.ENUM_VALUEs_reference = make(map[*ENUM_VALUE]*ENUM_VALUE)
	stage.ENUM_VALUEs_referenceOrder = make(map[*ENUM_VALUE]uint) // diff Unstage needs the reference order
	stage.ENUM_VALUEs_instance = make(map[*ENUM_VALUE]*ENUM_VALUE)
	for instance := range stage.ENUM_VALUEs {
		_copy := instance.GongCopy().(*ENUM_VALUE)
		stage.ENUM_VALUEs_reference[instance] = _copy
		stage.ENUM_VALUEs_instance[_copy] = instance
		stage.ENUM_VALUEs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.EmbeddedJpgImages_reference = make(map[*EmbeddedJpgImage]*EmbeddedJpgImage)
	stage.EmbeddedJpgImages_referenceOrder = make(map[*EmbeddedJpgImage]uint) // diff Unstage needs the reference order
	stage.EmbeddedJpgImages_instance = make(map[*EmbeddedJpgImage]*EmbeddedJpgImage)
	for instance := range stage.EmbeddedJpgImages {
		_copy := instance.GongCopy().(*EmbeddedJpgImage)
		stage.EmbeddedJpgImages_reference[instance] = _copy
		stage.EmbeddedJpgImages_instance[_copy] = instance
		stage.EmbeddedJpgImages_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.EmbeddedPngImages_reference = make(map[*EmbeddedPngImage]*EmbeddedPngImage)
	stage.EmbeddedPngImages_referenceOrder = make(map[*EmbeddedPngImage]uint) // diff Unstage needs the reference order
	stage.EmbeddedPngImages_instance = make(map[*EmbeddedPngImage]*EmbeddedPngImage)
	for instance := range stage.EmbeddedPngImages {
		_copy := instance.GongCopy().(*EmbeddedPngImage)
		stage.EmbeddedPngImages_reference[instance] = _copy
		stage.EmbeddedPngImages_instance[_copy] = instance
		stage.EmbeddedPngImages_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.EmbeddedSvgImages_reference = make(map[*EmbeddedSvgImage]*EmbeddedSvgImage)
	stage.EmbeddedSvgImages_referenceOrder = make(map[*EmbeddedSvgImage]uint) // diff Unstage needs the reference order
	stage.EmbeddedSvgImages_instance = make(map[*EmbeddedSvgImage]*EmbeddedSvgImage)
	for instance := range stage.EmbeddedSvgImages {
		_copy := instance.GongCopy().(*EmbeddedSvgImage)
		stage.EmbeddedSvgImages_reference[instance] = _copy
		stage.EmbeddedSvgImages_instance[_copy] = instance
		stage.EmbeddedSvgImages_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Kills_reference = make(map[*Kill]*Kill)
	stage.Kills_referenceOrder = make(map[*Kill]uint) // diff Unstage needs the reference order
	stage.Kills_instance = make(map[*Kill]*Kill)
	for instance := range stage.Kills {
		_copy := instance.GongCopy().(*Kill)
		stage.Kills_reference[instance] = _copy
		stage.Kills_instance[_copy] = instance
		stage.Kills_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Map_identifier_bools_reference = make(map[*Map_identifier_bool]*Map_identifier_bool)
	stage.Map_identifier_bools_referenceOrder = make(map[*Map_identifier_bool]uint) // diff Unstage needs the reference order
	stage.Map_identifier_bools_instance = make(map[*Map_identifier_bool]*Map_identifier_bool)
	for instance := range stage.Map_identifier_bools {
		_copy := instance.GongCopy().(*Map_identifier_bool)
		stage.Map_identifier_bools_reference[instance] = _copy
		stage.Map_identifier_bools_instance[_copy] = instance
		stage.Map_identifier_bools_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.RELATION_GROUPs_reference = make(map[*RELATION_GROUP]*RELATION_GROUP)
	stage.RELATION_GROUPs_referenceOrder = make(map[*RELATION_GROUP]uint) // diff Unstage needs the reference order
	stage.RELATION_GROUPs_instance = make(map[*RELATION_GROUP]*RELATION_GROUP)
	for instance := range stage.RELATION_GROUPs {
		_copy := instance.GongCopy().(*RELATION_GROUP)
		stage.RELATION_GROUPs_reference[instance] = _copy
		stage.RELATION_GROUPs_instance[_copy] = instance
		stage.RELATION_GROUPs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.RELATION_GROUP_TYPEs_reference = make(map[*RELATION_GROUP_TYPE]*RELATION_GROUP_TYPE)
	stage.RELATION_GROUP_TYPEs_referenceOrder = make(map[*RELATION_GROUP_TYPE]uint) // diff Unstage needs the reference order
	stage.RELATION_GROUP_TYPEs_instance = make(map[*RELATION_GROUP_TYPE]*RELATION_GROUP_TYPE)
	for instance := range stage.RELATION_GROUP_TYPEs {
		_copy := instance.GongCopy().(*RELATION_GROUP_TYPE)
		stage.RELATION_GROUP_TYPEs_reference[instance] = _copy
		stage.RELATION_GROUP_TYPEs_instance[_copy] = instance
		stage.RELATION_GROUP_TYPEs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.REQ_IFs_reference = make(map[*REQ_IF]*REQ_IF)
	stage.REQ_IFs_referenceOrder = make(map[*REQ_IF]uint) // diff Unstage needs the reference order
	stage.REQ_IFs_instance = make(map[*REQ_IF]*REQ_IF)
	for instance := range stage.REQ_IFs {
		_copy := instance.GongCopy().(*REQ_IF)
		stage.REQ_IFs_reference[instance] = _copy
		stage.REQ_IFs_instance[_copy] = instance
		stage.REQ_IFs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.REQ_IF_CONTENTs_reference = make(map[*REQ_IF_CONTENT]*REQ_IF_CONTENT)
	stage.REQ_IF_CONTENTs_referenceOrder = make(map[*REQ_IF_CONTENT]uint) // diff Unstage needs the reference order
	stage.REQ_IF_CONTENTs_instance = make(map[*REQ_IF_CONTENT]*REQ_IF_CONTENT)
	for instance := range stage.REQ_IF_CONTENTs {
		_copy := instance.GongCopy().(*REQ_IF_CONTENT)
		stage.REQ_IF_CONTENTs_reference[instance] = _copy
		stage.REQ_IF_CONTENTs_instance[_copy] = instance
		stage.REQ_IF_CONTENTs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.REQ_IF_HEADERs_reference = make(map[*REQ_IF_HEADER]*REQ_IF_HEADER)
	stage.REQ_IF_HEADERs_referenceOrder = make(map[*REQ_IF_HEADER]uint) // diff Unstage needs the reference order
	stage.REQ_IF_HEADERs_instance = make(map[*REQ_IF_HEADER]*REQ_IF_HEADER)
	for instance := range stage.REQ_IF_HEADERs {
		_copy := instance.GongCopy().(*REQ_IF_HEADER)
		stage.REQ_IF_HEADERs_reference[instance] = _copy
		stage.REQ_IF_HEADERs_instance[_copy] = instance
		stage.REQ_IF_HEADERs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.REQ_IF_TOOL_EXTENSIONs_reference = make(map[*REQ_IF_TOOL_EXTENSION]*REQ_IF_TOOL_EXTENSION)
	stage.REQ_IF_TOOL_EXTENSIONs_referenceOrder = make(map[*REQ_IF_TOOL_EXTENSION]uint) // diff Unstage needs the reference order
	stage.REQ_IF_TOOL_EXTENSIONs_instance = make(map[*REQ_IF_TOOL_EXTENSION]*REQ_IF_TOOL_EXTENSION)
	for instance := range stage.REQ_IF_TOOL_EXTENSIONs {
		_copy := instance.GongCopy().(*REQ_IF_TOOL_EXTENSION)
		stage.REQ_IF_TOOL_EXTENSIONs_reference[instance] = _copy
		stage.REQ_IF_TOOL_EXTENSIONs_instance[_copy] = instance
		stage.REQ_IF_TOOL_EXTENSIONs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SPECIFICATIONs_reference = make(map[*SPECIFICATION]*SPECIFICATION)
	stage.SPECIFICATIONs_referenceOrder = make(map[*SPECIFICATION]uint) // diff Unstage needs the reference order
	stage.SPECIFICATIONs_instance = make(map[*SPECIFICATION]*SPECIFICATION)
	for instance := range stage.SPECIFICATIONs {
		_copy := instance.GongCopy().(*SPECIFICATION)
		stage.SPECIFICATIONs_reference[instance] = _copy
		stage.SPECIFICATIONs_instance[_copy] = instance
		stage.SPECIFICATIONs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SPECIFICATION_Renderings_reference = make(map[*SPECIFICATION_Rendering]*SPECIFICATION_Rendering)
	stage.SPECIFICATION_Renderings_referenceOrder = make(map[*SPECIFICATION_Rendering]uint) // diff Unstage needs the reference order
	stage.SPECIFICATION_Renderings_instance = make(map[*SPECIFICATION_Rendering]*SPECIFICATION_Rendering)
	for instance := range stage.SPECIFICATION_Renderings {
		_copy := instance.GongCopy().(*SPECIFICATION_Rendering)
		stage.SPECIFICATION_Renderings_reference[instance] = _copy
		stage.SPECIFICATION_Renderings_instance[_copy] = instance
		stage.SPECIFICATION_Renderings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SPECIFICATION_TYPEs_reference = make(map[*SPECIFICATION_TYPE]*SPECIFICATION_TYPE)
	stage.SPECIFICATION_TYPEs_referenceOrder = make(map[*SPECIFICATION_TYPE]uint) // diff Unstage needs the reference order
	stage.SPECIFICATION_TYPEs_instance = make(map[*SPECIFICATION_TYPE]*SPECIFICATION_TYPE)
	for instance := range stage.SPECIFICATION_TYPEs {
		_copy := instance.GongCopy().(*SPECIFICATION_TYPE)
		stage.SPECIFICATION_TYPEs_reference[instance] = _copy
		stage.SPECIFICATION_TYPEs_instance[_copy] = instance
		stage.SPECIFICATION_TYPEs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SPEC_HIERARCHYs_reference = make(map[*SPEC_HIERARCHY]*SPEC_HIERARCHY)
	stage.SPEC_HIERARCHYs_referenceOrder = make(map[*SPEC_HIERARCHY]uint) // diff Unstage needs the reference order
	stage.SPEC_HIERARCHYs_instance = make(map[*SPEC_HIERARCHY]*SPEC_HIERARCHY)
	for instance := range stage.SPEC_HIERARCHYs {
		_copy := instance.GongCopy().(*SPEC_HIERARCHY)
		stage.SPEC_HIERARCHYs_reference[instance] = _copy
		stage.SPEC_HIERARCHYs_instance[_copy] = instance
		stage.SPEC_HIERARCHYs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SPEC_OBJECTs_reference = make(map[*SPEC_OBJECT]*SPEC_OBJECT)
	stage.SPEC_OBJECTs_referenceOrder = make(map[*SPEC_OBJECT]uint) // diff Unstage needs the reference order
	stage.SPEC_OBJECTs_instance = make(map[*SPEC_OBJECT]*SPEC_OBJECT)
	for instance := range stage.SPEC_OBJECTs {
		_copy := instance.GongCopy().(*SPEC_OBJECT)
		stage.SPEC_OBJECTs_reference[instance] = _copy
		stage.SPEC_OBJECTs_instance[_copy] = instance
		stage.SPEC_OBJECTs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SPEC_OBJECT_TYPEs_reference = make(map[*SPEC_OBJECT_TYPE]*SPEC_OBJECT_TYPE)
	stage.SPEC_OBJECT_TYPEs_referenceOrder = make(map[*SPEC_OBJECT_TYPE]uint) // diff Unstage needs the reference order
	stage.SPEC_OBJECT_TYPEs_instance = make(map[*SPEC_OBJECT_TYPE]*SPEC_OBJECT_TYPE)
	for instance := range stage.SPEC_OBJECT_TYPEs {
		_copy := instance.GongCopy().(*SPEC_OBJECT_TYPE)
		stage.SPEC_OBJECT_TYPEs_reference[instance] = _copy
		stage.SPEC_OBJECT_TYPEs_instance[_copy] = instance
		stage.SPEC_OBJECT_TYPEs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SPEC_OBJECT_TYPE_Renderings_reference = make(map[*SPEC_OBJECT_TYPE_Rendering]*SPEC_OBJECT_TYPE_Rendering)
	stage.SPEC_OBJECT_TYPE_Renderings_referenceOrder = make(map[*SPEC_OBJECT_TYPE_Rendering]uint) // diff Unstage needs the reference order
	stage.SPEC_OBJECT_TYPE_Renderings_instance = make(map[*SPEC_OBJECT_TYPE_Rendering]*SPEC_OBJECT_TYPE_Rendering)
	for instance := range stage.SPEC_OBJECT_TYPE_Renderings {
		_copy := instance.GongCopy().(*SPEC_OBJECT_TYPE_Rendering)
		stage.SPEC_OBJECT_TYPE_Renderings_reference[instance] = _copy
		stage.SPEC_OBJECT_TYPE_Renderings_instance[_copy] = instance
		stage.SPEC_OBJECT_TYPE_Renderings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SPEC_RELATIONs_reference = make(map[*SPEC_RELATION]*SPEC_RELATION)
	stage.SPEC_RELATIONs_referenceOrder = make(map[*SPEC_RELATION]uint) // diff Unstage needs the reference order
	stage.SPEC_RELATIONs_instance = make(map[*SPEC_RELATION]*SPEC_RELATION)
	for instance := range stage.SPEC_RELATIONs {
		_copy := instance.GongCopy().(*SPEC_RELATION)
		stage.SPEC_RELATIONs_reference[instance] = _copy
		stage.SPEC_RELATIONs_instance[_copy] = instance
		stage.SPEC_RELATIONs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.SPEC_RELATION_TYPEs_reference = make(map[*SPEC_RELATION_TYPE]*SPEC_RELATION_TYPE)
	stage.SPEC_RELATION_TYPEs_referenceOrder = make(map[*SPEC_RELATION_TYPE]uint) // diff Unstage needs the reference order
	stage.SPEC_RELATION_TYPEs_instance = make(map[*SPEC_RELATION_TYPE]*SPEC_RELATION_TYPE)
	for instance := range stage.SPEC_RELATION_TYPEs {
		_copy := instance.GongCopy().(*SPEC_RELATION_TYPE)
		stage.SPEC_RELATION_TYPEs_reference[instance] = _copy
		stage.SPEC_RELATION_TYPEs_instance[_copy] = instance
		stage.SPEC_RELATION_TYPEs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StaticWebSites_reference = make(map[*StaticWebSite]*StaticWebSite)
	stage.StaticWebSites_referenceOrder = make(map[*StaticWebSite]uint) // diff Unstage needs the reference order
	stage.StaticWebSites_instance = make(map[*StaticWebSite]*StaticWebSite)
	for instance := range stage.StaticWebSites {
		_copy := instance.GongCopy().(*StaticWebSite)
		stage.StaticWebSites_reference[instance] = _copy
		stage.StaticWebSites_instance[_copy] = instance
		stage.StaticWebSites_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StaticWebSiteChapters_reference = make(map[*StaticWebSiteChapter]*StaticWebSiteChapter)
	stage.StaticWebSiteChapters_referenceOrder = make(map[*StaticWebSiteChapter]uint) // diff Unstage needs the reference order
	stage.StaticWebSiteChapters_instance = make(map[*StaticWebSiteChapter]*StaticWebSiteChapter)
	for instance := range stage.StaticWebSiteChapters {
		_copy := instance.GongCopy().(*StaticWebSiteChapter)
		stage.StaticWebSiteChapters_reference[instance] = _copy
		stage.StaticWebSiteChapters_instance[_copy] = instance
		stage.StaticWebSiteChapters_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StaticWebSiteGeneratedImages_reference = make(map[*StaticWebSiteGeneratedImage]*StaticWebSiteGeneratedImage)
	stage.StaticWebSiteGeneratedImages_referenceOrder = make(map[*StaticWebSiteGeneratedImage]uint) // diff Unstage needs the reference order
	stage.StaticWebSiteGeneratedImages_instance = make(map[*StaticWebSiteGeneratedImage]*StaticWebSiteGeneratedImage)
	for instance := range stage.StaticWebSiteGeneratedImages {
		_copy := instance.GongCopy().(*StaticWebSiteGeneratedImage)
		stage.StaticWebSiteGeneratedImages_reference[instance] = _copy
		stage.StaticWebSiteGeneratedImages_instance[_copy] = instance
		stage.StaticWebSiteGeneratedImages_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StaticWebSiteImages_reference = make(map[*StaticWebSiteImage]*StaticWebSiteImage)
	stage.StaticWebSiteImages_referenceOrder = make(map[*StaticWebSiteImage]uint) // diff Unstage needs the reference order
	stage.StaticWebSiteImages_instance = make(map[*StaticWebSiteImage]*StaticWebSiteImage)
	for instance := range stage.StaticWebSiteImages {
		_copy := instance.GongCopy().(*StaticWebSiteImage)
		stage.StaticWebSiteImages_reference[instance] = _copy
		stage.StaticWebSiteImages_instance[_copy] = instance
		stage.StaticWebSiteImages_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.StaticWebSiteParagraphs_reference = make(map[*StaticWebSiteParagraph]*StaticWebSiteParagraph)
	stage.StaticWebSiteParagraphs_referenceOrder = make(map[*StaticWebSiteParagraph]uint) // diff Unstage needs the reference order
	stage.StaticWebSiteParagraphs_instance = make(map[*StaticWebSiteParagraph]*StaticWebSiteParagraph)
	for instance := range stage.StaticWebSiteParagraphs {
		_copy := instance.GongCopy().(*StaticWebSiteParagraph)
		stage.StaticWebSiteParagraphs_reference[instance] = _copy
		stage.StaticWebSiteParagraphs_instance[_copy] = instance
		stage.StaticWebSiteParagraphs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.XHTML_CONTENTs_reference = make(map[*XHTML_CONTENT]*XHTML_CONTENT)
	stage.XHTML_CONTENTs_referenceOrder = make(map[*XHTML_CONTENT]uint) // diff Unstage needs the reference order
	stage.XHTML_CONTENTs_instance = make(map[*XHTML_CONTENT]*XHTML_CONTENT)
	for instance := range stage.XHTML_CONTENTs {
		_copy := instance.GongCopy().(*XHTML_CONTENT)
		stage.XHTML_CONTENTs_reference[instance] = _copy
		stage.XHTML_CONTENTs_instance[_copy] = instance
		stage.XHTML_CONTENTs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.ALTERNATIVE_IDs {
		reference := stage.ALTERNATIVE_IDs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
		reference := stage.ATTRIBUTE_DEFINITION_BOOLEANs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings {
		reference := stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_DATEs {
		reference := stage.ATTRIBUTE_DEFINITION_DATEs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_DATE_Renderings {
		reference := stage.ATTRIBUTE_DEFINITION_DATE_Renderings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
		reference := stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings {
		reference := stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
		reference := stage.ATTRIBUTE_DEFINITION_INTEGERs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings {
		reference := stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_REALs {
		reference := stage.ATTRIBUTE_DEFINITION_REALs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_REAL_Renderings {
		reference := stage.ATTRIBUTE_DEFINITION_REAL_Renderings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_Renderings {
		reference := stage.ATTRIBUTE_DEFINITION_Renderings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_STRINGs {
		reference := stage.ATTRIBUTE_DEFINITION_STRINGs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_STRING_Renderings {
		reference := stage.ATTRIBUTE_DEFINITION_STRING_Renderings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
		reference := stage.ATTRIBUTE_DEFINITION_XHTMLs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_DEFINITION_XHTML_Renderings {
		reference := stage.ATTRIBUTE_DEFINITION_XHTML_Renderings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_BOOLEANs {
		reference := stage.ATTRIBUTE_VALUE_BOOLEANs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_DATEs {
		reference := stage.ATTRIBUTE_VALUE_DATEs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_ENUMERATIONs {
		reference := stage.ATTRIBUTE_VALUE_ENUMERATIONs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_INTEGERs {
		reference := stage.ATTRIBUTE_VALUE_INTEGERs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_REALs {
		reference := stage.ATTRIBUTE_VALUE_REALs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_STRINGs {
		reference := stage.ATTRIBUTE_VALUE_STRINGs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ATTRIBUTE_VALUE_XHTMLs {
		reference := stage.ATTRIBUTE_VALUE_XHTMLs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ALTERNATIVE_IDs {
		reference := stage.A_ALTERNATIVE_IDs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs {
		reference := stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_DATE_REFs {
		reference := stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs {
		reference := stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs {
		reference := stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_REAL_REFs {
		reference := stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_STRING_REFs {
		reference := stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs {
		reference := stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_BOOLEANs {
		reference := stage.A_ATTRIBUTE_VALUE_BOOLEANs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_DATEs {
		reference := stage.A_ATTRIBUTE_VALUE_DATEs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
		reference := stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_INTEGERs {
		reference := stage.A_ATTRIBUTE_VALUE_INTEGERs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_REALs {
		reference := stage.A_ATTRIBUTE_VALUE_REALs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_STRINGs {
		reference := stage.A_ATTRIBUTE_VALUE_STRINGs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
		reference := stage.A_ATTRIBUTE_VALUE_XHTMLs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
		reference := stage.A_ATTRIBUTE_VALUE_XHTML_1s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_CHILDRENs {
		reference := stage.A_CHILDRENs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_CORE_CONTENTs {
		reference := stage.A_CORE_CONTENTs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_DATATYPESs {
		reference := stage.A_DATATYPESs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs {
		reference := stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_DATE_REFs {
		reference := stage.A_DATATYPE_DEFINITION_DATE_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs {
		reference := stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_INTEGER_REFs {
		reference := stage.A_DATATYPE_DEFINITION_INTEGER_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_REAL_REFs {
		reference := stage.A_DATATYPE_DEFINITION_REAL_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_STRING_REFs {
		reference := stage.A_DATATYPE_DEFINITION_STRING_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_DATATYPE_DEFINITION_XHTML_REFs {
		reference := stage.A_DATATYPE_DEFINITION_XHTML_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_EDITABLE_ATTSs {
		reference := stage.A_EDITABLE_ATTSs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_ENUM_VALUE_REFs {
		reference := stage.A_ENUM_VALUE_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_OBJECTs {
		reference := stage.A_OBJECTs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_PROPERTIESs {
		reference := stage.A_PROPERTIESs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_RELATION_GROUP_TYPE_REFs {
		reference := stage.A_RELATION_GROUP_TYPE_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_SOURCE_1s {
		reference := stage.A_SOURCE_1s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_SOURCE_SPECIFICATION_1s {
		reference := stage.A_SOURCE_SPECIFICATION_1s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_SPECIFICATIONSs {
		reference := stage.A_SPECIFICATIONSs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_SPECIFICATION_TYPE_REFs {
		reference := stage.A_SPECIFICATION_TYPE_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_SPECIFIED_VALUESs {
		reference := stage.A_SPECIFIED_VALUESs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_SPEC_ATTRIBUTESs {
		reference := stage.A_SPEC_ATTRIBUTESs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_SPEC_OBJECTSs {
		reference := stage.A_SPEC_OBJECTSs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_SPEC_OBJECT_TYPE_REFs {
		reference := stage.A_SPEC_OBJECT_TYPE_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_SPEC_RELATIONSs {
		reference := stage.A_SPEC_RELATIONSs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_SPEC_RELATION_GROUPSs {
		reference := stage.A_SPEC_RELATION_GROUPSs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_SPEC_RELATION_REFs {
		reference := stage.A_SPEC_RELATION_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_SPEC_RELATION_TYPE_REFs {
		reference := stage.A_SPEC_RELATION_TYPE_REFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_SPEC_TYPESs {
		reference := stage.A_SPEC_TYPESs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_THE_HEADERs {
		reference := stage.A_THE_HEADERs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_TOOL_EXTENSIONSs {
		reference := stage.A_TOOL_EXTENSIONSs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_BOOLEANs {
		reference := stage.DATATYPE_DEFINITION_BOOLEANs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_DATEs {
		reference := stage.DATATYPE_DEFINITION_DATEs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
		reference := stage.DATATYPE_DEFINITION_ENUMERATIONs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_INTEGERs {
		reference := stage.DATATYPE_DEFINITION_INTEGERs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_REALs {
		reference := stage.DATATYPE_DEFINITION_REALs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_STRINGs {
		reference := stage.DATATYPE_DEFINITION_STRINGs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.DATATYPE_DEFINITION_XHTMLs {
		reference := stage.DATATYPE_DEFINITION_XHTMLs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EMBEDDED_VALUEs {
		reference := stage.EMBEDDED_VALUEs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.ENUM_VALUEs {
		reference := stage.ENUM_VALUEs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EmbeddedJpgImages {
		reference := stage.EmbeddedJpgImages_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EmbeddedPngImages {
		reference := stage.EmbeddedPngImages_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.EmbeddedSvgImages {
		reference := stage.EmbeddedSvgImages_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Kills {
		reference := stage.Kills_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Map_identifier_bools {
		reference := stage.Map_identifier_bools_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.RELATION_GROUPs {
		reference := stage.RELATION_GROUPs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.RELATION_GROUP_TYPEs {
		reference := stage.RELATION_GROUP_TYPEs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.REQ_IFs {
		reference := stage.REQ_IFs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.REQ_IF_CONTENTs {
		reference := stage.REQ_IF_CONTENTs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.REQ_IF_HEADERs {
		reference := stage.REQ_IF_HEADERs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.REQ_IF_TOOL_EXTENSIONs {
		reference := stage.REQ_IF_TOOL_EXTENSIONs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SPECIFICATIONs {
		reference := stage.SPECIFICATIONs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SPECIFICATION_Renderings {
		reference := stage.SPECIFICATION_Renderings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SPECIFICATION_TYPEs {
		reference := stage.SPECIFICATION_TYPEs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SPEC_HIERARCHYs {
		reference := stage.SPEC_HIERARCHYs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SPEC_OBJECTs {
		reference := stage.SPEC_OBJECTs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SPEC_OBJECT_TYPEs {
		reference := stage.SPEC_OBJECT_TYPEs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SPEC_OBJECT_TYPE_Renderings {
		reference := stage.SPEC_OBJECT_TYPE_Renderings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SPEC_RELATIONs {
		reference := stage.SPEC_RELATIONs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.SPEC_RELATION_TYPEs {
		reference := stage.SPEC_RELATION_TYPEs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StaticWebSites {
		reference := stage.StaticWebSites_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StaticWebSiteChapters {
		reference := stage.StaticWebSiteChapters_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StaticWebSiteGeneratedImages {
		reference := stage.StaticWebSiteGeneratedImages_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StaticWebSiteImages {
		reference := stage.StaticWebSiteImages_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.StaticWebSiteParagraphs {
		reference := stage.StaticWebSiteParagraphs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.XHTML_CONTENTs {
		reference := stage.XHTML_CONTENTs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
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
	if order, ok := stage.ALTERNATIVE_ID_stagedOrder[alternative_id]; ok {
		return order
	}
	if order, ok := stage.ALTERNATIVE_IDs_referenceOrder[alternative_id]; ok {
		return order
	} else {
		log.Printf("instance %p of type ALTERNATIVE_ID was not staged and does not have a reference order", alternative_id)
		return 0
	}
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_BOOLEAN_stagedOrder[attribute_definition_boolean]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_DEFINITION_BOOLEANs_referenceOrder[attribute_definition_boolean]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_DEFINITION_BOOLEAN was not staged and does not have a reference order", attribute_definition_boolean)
		return 0
	}
}

func (attribute_definition_boolean_rendering *ATTRIBUTE_DEFINITION_BOOLEAN_Rendering) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_BOOLEAN_Rendering_stagedOrder[attribute_definition_boolean_rendering]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings_referenceOrder[attribute_definition_boolean_rendering]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_DEFINITION_BOOLEAN_Rendering was not staged and does not have a reference order", attribute_definition_boolean_rendering)
		return 0
	}
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_DATE_stagedOrder[attribute_definition_date]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_DEFINITION_DATEs_referenceOrder[attribute_definition_date]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_DEFINITION_DATE was not staged and does not have a reference order", attribute_definition_date)
		return 0
	}
}

func (attribute_definition_date_rendering *ATTRIBUTE_DEFINITION_DATE_Rendering) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_DATE_Rendering_stagedOrder[attribute_definition_date_rendering]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_DEFINITION_DATE_Renderings_referenceOrder[attribute_definition_date_rendering]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_DEFINITION_DATE_Rendering was not staged and does not have a reference order", attribute_definition_date_rendering)
		return 0
	}
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATION_stagedOrder[attribute_definition_enumeration]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_referenceOrder[attribute_definition_enumeration]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_DEFINITION_ENUMERATION was not staged and does not have a reference order", attribute_definition_enumeration)
		return 0
	}
}

func (attribute_definition_enumeration_rendering *ATTRIBUTE_DEFINITION_ENUMERATION_Rendering) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATION_Rendering_stagedOrder[attribute_definition_enumeration_rendering]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings_referenceOrder[attribute_definition_enumeration_rendering]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_DEFINITION_ENUMERATION_Rendering was not staged and does not have a reference order", attribute_definition_enumeration_rendering)
		return 0
	}
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_INTEGER_stagedOrder[attribute_definition_integer]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_DEFINITION_INTEGERs_referenceOrder[attribute_definition_integer]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_DEFINITION_INTEGER was not staged and does not have a reference order", attribute_definition_integer)
		return 0
	}
}

func (attribute_definition_integer_rendering *ATTRIBUTE_DEFINITION_INTEGER_Rendering) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_INTEGER_Rendering_stagedOrder[attribute_definition_integer_rendering]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings_referenceOrder[attribute_definition_integer_rendering]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_DEFINITION_INTEGER_Rendering was not staged and does not have a reference order", attribute_definition_integer_rendering)
		return 0
	}
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_REAL_stagedOrder[attribute_definition_real]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_DEFINITION_REALs_referenceOrder[attribute_definition_real]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_DEFINITION_REAL was not staged and does not have a reference order", attribute_definition_real)
		return 0
	}
}

func (attribute_definition_real_rendering *ATTRIBUTE_DEFINITION_REAL_Rendering) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_REAL_Rendering_stagedOrder[attribute_definition_real_rendering]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_DEFINITION_REAL_Renderings_referenceOrder[attribute_definition_real_rendering]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_DEFINITION_REAL_Rendering was not staged and does not have a reference order", attribute_definition_real_rendering)
		return 0
	}
}

func (attribute_definition_rendering *ATTRIBUTE_DEFINITION_Rendering) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_Rendering_stagedOrder[attribute_definition_rendering]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_DEFINITION_Renderings_referenceOrder[attribute_definition_rendering]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_DEFINITION_Rendering was not staged and does not have a reference order", attribute_definition_rendering)
		return 0
	}
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_STRING_stagedOrder[attribute_definition_string]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_DEFINITION_STRINGs_referenceOrder[attribute_definition_string]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_DEFINITION_STRING was not staged and does not have a reference order", attribute_definition_string)
		return 0
	}
}

func (attribute_definition_string_rendering *ATTRIBUTE_DEFINITION_STRING_Rendering) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_STRING_Rendering_stagedOrder[attribute_definition_string_rendering]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_DEFINITION_STRING_Renderings_referenceOrder[attribute_definition_string_rendering]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_DEFINITION_STRING_Rendering was not staged and does not have a reference order", attribute_definition_string_rendering)
		return 0
	}
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_XHTML_stagedOrder[attribute_definition_xhtml]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_DEFINITION_XHTMLs_referenceOrder[attribute_definition_xhtml]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_DEFINITION_XHTML was not staged and does not have a reference order", attribute_definition_xhtml)
		return 0
	}
}

func (attribute_definition_xhtml_rendering *ATTRIBUTE_DEFINITION_XHTML_Rendering) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_DEFINITION_XHTML_Rendering_stagedOrder[attribute_definition_xhtml_rendering]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_DEFINITION_XHTML_Renderings_referenceOrder[attribute_definition_xhtml_rendering]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_DEFINITION_XHTML_Rendering was not staged and does not have a reference order", attribute_definition_xhtml_rendering)
		return 0
	}
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_VALUE_BOOLEAN_stagedOrder[attribute_value_boolean]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_VALUE_BOOLEANs_referenceOrder[attribute_value_boolean]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_VALUE_BOOLEAN was not staged and does not have a reference order", attribute_value_boolean)
		return 0
	}
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_VALUE_DATE_stagedOrder[attribute_value_date]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_VALUE_DATEs_referenceOrder[attribute_value_date]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_VALUE_DATE was not staged and does not have a reference order", attribute_value_date)
		return 0
	}
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_VALUE_ENUMERATION_stagedOrder[attribute_value_enumeration]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder[attribute_value_enumeration]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_VALUE_ENUMERATION was not staged and does not have a reference order", attribute_value_enumeration)
		return 0
	}
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_VALUE_INTEGER_stagedOrder[attribute_value_integer]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_VALUE_INTEGERs_referenceOrder[attribute_value_integer]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_VALUE_INTEGER was not staged and does not have a reference order", attribute_value_integer)
		return 0
	}
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_VALUE_REAL_stagedOrder[attribute_value_real]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_VALUE_REALs_referenceOrder[attribute_value_real]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_VALUE_REAL was not staged and does not have a reference order", attribute_value_real)
		return 0
	}
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_VALUE_STRING_stagedOrder[attribute_value_string]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_VALUE_STRINGs_referenceOrder[attribute_value_string]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_VALUE_STRING was not staged and does not have a reference order", attribute_value_string)
		return 0
	}
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ATTRIBUTE_VALUE_XHTML_stagedOrder[attribute_value_xhtml]; ok {
		return order
	}
	if order, ok := stage.ATTRIBUTE_VALUE_XHTMLs_referenceOrder[attribute_value_xhtml]; ok {
		return order
	} else {
		log.Printf("instance %p of type ATTRIBUTE_VALUE_XHTML was not staged and does not have a reference order", attribute_value_xhtml)
		return 0
	}
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ALTERNATIVE_ID_stagedOrder[a_alternative_id]; ok {
		return order
	}
	if order, ok := stage.A_ALTERNATIVE_IDs_referenceOrder[a_alternative_id]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ALTERNATIVE_ID was not staged and does not have a reference order", a_alternative_id)
		return 0
	}
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_stagedOrder[a_attribute_definition_boolean_ref]; ok {
		return order
	}
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_referenceOrder[a_attribute_definition_boolean_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ATTRIBUTE_DEFINITION_BOOLEAN_REF was not staged and does not have a reference order", a_attribute_definition_boolean_ref)
		return 0
	}
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_DATE_REF_stagedOrder[a_attribute_definition_date_ref]; ok {
		return order
	}
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_referenceOrder[a_attribute_definition_date_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ATTRIBUTE_DEFINITION_DATE_REF was not staged and does not have a reference order", a_attribute_definition_date_ref)
		return 0
	}
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_stagedOrder[a_attribute_definition_enumeration_ref]; ok {
		return order
	}
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_referenceOrder[a_attribute_definition_enumeration_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ATTRIBUTE_DEFINITION_ENUMERATION_REF was not staged and does not have a reference order", a_attribute_definition_enumeration_ref)
		return 0
	}
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_INTEGER_REF_stagedOrder[a_attribute_definition_integer_ref]; ok {
		return order
	}
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_referenceOrder[a_attribute_definition_integer_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ATTRIBUTE_DEFINITION_INTEGER_REF was not staged and does not have a reference order", a_attribute_definition_integer_ref)
		return 0
	}
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_REAL_REF_stagedOrder[a_attribute_definition_real_ref]; ok {
		return order
	}
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_referenceOrder[a_attribute_definition_real_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ATTRIBUTE_DEFINITION_REAL_REF was not staged and does not have a reference order", a_attribute_definition_real_ref)
		return 0
	}
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_STRING_REF_stagedOrder[a_attribute_definition_string_ref]; ok {
		return order
	}
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_referenceOrder[a_attribute_definition_string_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ATTRIBUTE_DEFINITION_STRING_REF was not staged and does not have a reference order", a_attribute_definition_string_ref)
		return 0
	}
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_XHTML_REF_stagedOrder[a_attribute_definition_xhtml_ref]; ok {
		return order
	}
	if order, ok := stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_referenceOrder[a_attribute_definition_xhtml_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ATTRIBUTE_DEFINITION_XHTML_REF was not staged and does not have a reference order", a_attribute_definition_xhtml_ref)
		return 0
	}
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_VALUE_BOOLEAN_stagedOrder[a_attribute_value_boolean]; ok {
		return order
	}
	if order, ok := stage.A_ATTRIBUTE_VALUE_BOOLEANs_referenceOrder[a_attribute_value_boolean]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ATTRIBUTE_VALUE_BOOLEAN was not staged and does not have a reference order", a_attribute_value_boolean)
		return 0
	}
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_VALUE_DATE_stagedOrder[a_attribute_value_date]; ok {
		return order
	}
	if order, ok := stage.A_ATTRIBUTE_VALUE_DATEs_referenceOrder[a_attribute_value_date]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ATTRIBUTE_VALUE_DATE was not staged and does not have a reference order", a_attribute_value_date)
		return 0
	}
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_VALUE_ENUMERATION_stagedOrder[a_attribute_value_enumeration]; ok {
		return order
	}
	if order, ok := stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_referenceOrder[a_attribute_value_enumeration]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ATTRIBUTE_VALUE_ENUMERATION was not staged and does not have a reference order", a_attribute_value_enumeration)
		return 0
	}
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_VALUE_INTEGER_stagedOrder[a_attribute_value_integer]; ok {
		return order
	}
	if order, ok := stage.A_ATTRIBUTE_VALUE_INTEGERs_referenceOrder[a_attribute_value_integer]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ATTRIBUTE_VALUE_INTEGER was not staged and does not have a reference order", a_attribute_value_integer)
		return 0
	}
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_VALUE_REAL_stagedOrder[a_attribute_value_real]; ok {
		return order
	}
	if order, ok := stage.A_ATTRIBUTE_VALUE_REALs_referenceOrder[a_attribute_value_real]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ATTRIBUTE_VALUE_REAL was not staged and does not have a reference order", a_attribute_value_real)
		return 0
	}
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_VALUE_STRING_stagedOrder[a_attribute_value_string]; ok {
		return order
	}
	if order, ok := stage.A_ATTRIBUTE_VALUE_STRINGs_referenceOrder[a_attribute_value_string]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ATTRIBUTE_VALUE_STRING was not staged and does not have a reference order", a_attribute_value_string)
		return 0
	}
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_VALUE_XHTML_stagedOrder[a_attribute_value_xhtml]; ok {
		return order
	}
	if order, ok := stage.A_ATTRIBUTE_VALUE_XHTMLs_referenceOrder[a_attribute_value_xhtml]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ATTRIBUTE_VALUE_XHTML was not staged and does not have a reference order", a_attribute_value_xhtml)
		return 0
	}
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1_stagedOrder[a_attribute_value_xhtml_1]; ok {
		return order
	}
	if order, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1s_referenceOrder[a_attribute_value_xhtml_1]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ATTRIBUTE_VALUE_XHTML_1 was not staged and does not have a reference order", a_attribute_value_xhtml_1)
		return 0
	}
}

func (a_children *A_CHILDREN) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_CHILDREN_stagedOrder[a_children]; ok {
		return order
	}
	if order, ok := stage.A_CHILDRENs_referenceOrder[a_children]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_CHILDREN was not staged and does not have a reference order", a_children)
		return 0
	}
}

func (a_core_content *A_CORE_CONTENT) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_CORE_CONTENT_stagedOrder[a_core_content]; ok {
		return order
	}
	if order, ok := stage.A_CORE_CONTENTs_referenceOrder[a_core_content]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_CORE_CONTENT was not staged and does not have a reference order", a_core_content)
		return 0
	}
}

func (a_datatypes *A_DATATYPES) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_DATATYPES_stagedOrder[a_datatypes]; ok {
		return order
	}
	if order, ok := stage.A_DATATYPESs_referenceOrder[a_datatypes]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_DATATYPES was not staged and does not have a reference order", a_datatypes)
		return 0
	}
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_DATATYPE_DEFINITION_BOOLEAN_REF_stagedOrder[a_datatype_definition_boolean_ref]; ok {
		return order
	}
	if order, ok := stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_referenceOrder[a_datatype_definition_boolean_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_DATATYPE_DEFINITION_BOOLEAN_REF was not staged and does not have a reference order", a_datatype_definition_boolean_ref)
		return 0
	}
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_DATATYPE_DEFINITION_DATE_REF_stagedOrder[a_datatype_definition_date_ref]; ok {
		return order
	}
	if order, ok := stage.A_DATATYPE_DEFINITION_DATE_REFs_referenceOrder[a_datatype_definition_date_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_DATATYPE_DEFINITION_DATE_REF was not staged and does not have a reference order", a_datatype_definition_date_ref)
		return 0
	}
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_DATATYPE_DEFINITION_ENUMERATION_REF_stagedOrder[a_datatype_definition_enumeration_ref]; ok {
		return order
	}
	if order, ok := stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_referenceOrder[a_datatype_definition_enumeration_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_DATATYPE_DEFINITION_ENUMERATION_REF was not staged and does not have a reference order", a_datatype_definition_enumeration_ref)
		return 0
	}
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_DATATYPE_DEFINITION_INTEGER_REF_stagedOrder[a_datatype_definition_integer_ref]; ok {
		return order
	}
	if order, ok := stage.A_DATATYPE_DEFINITION_INTEGER_REFs_referenceOrder[a_datatype_definition_integer_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_DATATYPE_DEFINITION_INTEGER_REF was not staged and does not have a reference order", a_datatype_definition_integer_ref)
		return 0
	}
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_DATATYPE_DEFINITION_REAL_REF_stagedOrder[a_datatype_definition_real_ref]; ok {
		return order
	}
	if order, ok := stage.A_DATATYPE_DEFINITION_REAL_REFs_referenceOrder[a_datatype_definition_real_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_DATATYPE_DEFINITION_REAL_REF was not staged and does not have a reference order", a_datatype_definition_real_ref)
		return 0
	}
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_DATATYPE_DEFINITION_STRING_REF_stagedOrder[a_datatype_definition_string_ref]; ok {
		return order
	}
	if order, ok := stage.A_DATATYPE_DEFINITION_STRING_REFs_referenceOrder[a_datatype_definition_string_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_DATATYPE_DEFINITION_STRING_REF was not staged and does not have a reference order", a_datatype_definition_string_ref)
		return 0
	}
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_DATATYPE_DEFINITION_XHTML_REF_stagedOrder[a_datatype_definition_xhtml_ref]; ok {
		return order
	}
	if order, ok := stage.A_DATATYPE_DEFINITION_XHTML_REFs_referenceOrder[a_datatype_definition_xhtml_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_DATATYPE_DEFINITION_XHTML_REF was not staged and does not have a reference order", a_datatype_definition_xhtml_ref)
		return 0
	}
}

func (a_editable_atts *A_EDITABLE_ATTS) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_EDITABLE_ATTS_stagedOrder[a_editable_atts]; ok {
		return order
	}
	if order, ok := stage.A_EDITABLE_ATTSs_referenceOrder[a_editable_atts]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_EDITABLE_ATTS was not staged and does not have a reference order", a_editable_atts)
		return 0
	}
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_ENUM_VALUE_REF_stagedOrder[a_enum_value_ref]; ok {
		return order
	}
	if order, ok := stage.A_ENUM_VALUE_REFs_referenceOrder[a_enum_value_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_ENUM_VALUE_REF was not staged and does not have a reference order", a_enum_value_ref)
		return 0
	}
}

func (a_object *A_OBJECT) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_OBJECT_stagedOrder[a_object]; ok {
		return order
	}
	if order, ok := stage.A_OBJECTs_referenceOrder[a_object]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_OBJECT was not staged and does not have a reference order", a_object)
		return 0
	}
}

func (a_properties *A_PROPERTIES) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_PROPERTIES_stagedOrder[a_properties]; ok {
		return order
	}
	if order, ok := stage.A_PROPERTIESs_referenceOrder[a_properties]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_PROPERTIES was not staged and does not have a reference order", a_properties)
		return 0
	}
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_RELATION_GROUP_TYPE_REF_stagedOrder[a_relation_group_type_ref]; ok {
		return order
	}
	if order, ok := stage.A_RELATION_GROUP_TYPE_REFs_referenceOrder[a_relation_group_type_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_RELATION_GROUP_TYPE_REF was not staged and does not have a reference order", a_relation_group_type_ref)
		return 0
	}
}

func (a_source_1 *A_SOURCE_1) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SOURCE_1_stagedOrder[a_source_1]; ok {
		return order
	}
	if order, ok := stage.A_SOURCE_1s_referenceOrder[a_source_1]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_SOURCE_1 was not staged and does not have a reference order", a_source_1)
		return 0
	}
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SOURCE_SPECIFICATION_1_stagedOrder[a_source_specification_1]; ok {
		return order
	}
	if order, ok := stage.A_SOURCE_SPECIFICATION_1s_referenceOrder[a_source_specification_1]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_SOURCE_SPECIFICATION_1 was not staged and does not have a reference order", a_source_specification_1)
		return 0
	}
}

func (a_specifications *A_SPECIFICATIONS) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPECIFICATIONS_stagedOrder[a_specifications]; ok {
		return order
	}
	if order, ok := stage.A_SPECIFICATIONSs_referenceOrder[a_specifications]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_SPECIFICATIONS was not staged and does not have a reference order", a_specifications)
		return 0
	}
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPECIFICATION_TYPE_REF_stagedOrder[a_specification_type_ref]; ok {
		return order
	}
	if order, ok := stage.A_SPECIFICATION_TYPE_REFs_referenceOrder[a_specification_type_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_SPECIFICATION_TYPE_REF was not staged and does not have a reference order", a_specification_type_ref)
		return 0
	}
}

func (a_specified_values *A_SPECIFIED_VALUES) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPECIFIED_VALUES_stagedOrder[a_specified_values]; ok {
		return order
	}
	if order, ok := stage.A_SPECIFIED_VALUESs_referenceOrder[a_specified_values]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_SPECIFIED_VALUES was not staged and does not have a reference order", a_specified_values)
		return 0
	}
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPEC_ATTRIBUTES_stagedOrder[a_spec_attributes]; ok {
		return order
	}
	if order, ok := stage.A_SPEC_ATTRIBUTESs_referenceOrder[a_spec_attributes]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_SPEC_ATTRIBUTES was not staged and does not have a reference order", a_spec_attributes)
		return 0
	}
}

func (a_spec_objects *A_SPEC_OBJECTS) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPEC_OBJECTS_stagedOrder[a_spec_objects]; ok {
		return order
	}
	if order, ok := stage.A_SPEC_OBJECTSs_referenceOrder[a_spec_objects]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_SPEC_OBJECTS was not staged and does not have a reference order", a_spec_objects)
		return 0
	}
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPEC_OBJECT_TYPE_REF_stagedOrder[a_spec_object_type_ref]; ok {
		return order
	}
	if order, ok := stage.A_SPEC_OBJECT_TYPE_REFs_referenceOrder[a_spec_object_type_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_SPEC_OBJECT_TYPE_REF was not staged and does not have a reference order", a_spec_object_type_ref)
		return 0
	}
}

func (a_spec_relations *A_SPEC_RELATIONS) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPEC_RELATIONS_stagedOrder[a_spec_relations]; ok {
		return order
	}
	if order, ok := stage.A_SPEC_RELATIONSs_referenceOrder[a_spec_relations]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_SPEC_RELATIONS was not staged and does not have a reference order", a_spec_relations)
		return 0
	}
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPEC_RELATION_GROUPS_stagedOrder[a_spec_relation_groups]; ok {
		return order
	}
	if order, ok := stage.A_SPEC_RELATION_GROUPSs_referenceOrder[a_spec_relation_groups]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_SPEC_RELATION_GROUPS was not staged and does not have a reference order", a_spec_relation_groups)
		return 0
	}
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPEC_RELATION_REF_stagedOrder[a_spec_relation_ref]; ok {
		return order
	}
	if order, ok := stage.A_SPEC_RELATION_REFs_referenceOrder[a_spec_relation_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_SPEC_RELATION_REF was not staged and does not have a reference order", a_spec_relation_ref)
		return 0
	}
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPEC_RELATION_TYPE_REF_stagedOrder[a_spec_relation_type_ref]; ok {
		return order
	}
	if order, ok := stage.A_SPEC_RELATION_TYPE_REFs_referenceOrder[a_spec_relation_type_ref]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_SPEC_RELATION_TYPE_REF was not staged and does not have a reference order", a_spec_relation_type_ref)
		return 0
	}
}

func (a_spec_types *A_SPEC_TYPES) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_SPEC_TYPES_stagedOrder[a_spec_types]; ok {
		return order
	}
	if order, ok := stage.A_SPEC_TYPESs_referenceOrder[a_spec_types]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_SPEC_TYPES was not staged and does not have a reference order", a_spec_types)
		return 0
	}
}

func (a_the_header *A_THE_HEADER) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_THE_HEADER_stagedOrder[a_the_header]; ok {
		return order
	}
	if order, ok := stage.A_THE_HEADERs_referenceOrder[a_the_header]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_THE_HEADER was not staged and does not have a reference order", a_the_header)
		return 0
	}
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_TOOL_EXTENSIONS_stagedOrder[a_tool_extensions]; ok {
		return order
	}
	if order, ok := stage.A_TOOL_EXTENSIONSs_referenceOrder[a_tool_extensions]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_TOOL_EXTENSIONS was not staged and does not have a reference order", a_tool_extensions)
		return 0
	}
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DATATYPE_DEFINITION_BOOLEAN_stagedOrder[datatype_definition_boolean]; ok {
		return order
	}
	if order, ok := stage.DATATYPE_DEFINITION_BOOLEANs_referenceOrder[datatype_definition_boolean]; ok {
		return order
	} else {
		log.Printf("instance %p of type DATATYPE_DEFINITION_BOOLEAN was not staged and does not have a reference order", datatype_definition_boolean)
		return 0
	}
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DATATYPE_DEFINITION_DATE_stagedOrder[datatype_definition_date]; ok {
		return order
	}
	if order, ok := stage.DATATYPE_DEFINITION_DATEs_referenceOrder[datatype_definition_date]; ok {
		return order
	} else {
		log.Printf("instance %p of type DATATYPE_DEFINITION_DATE was not staged and does not have a reference order", datatype_definition_date)
		return 0
	}
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DATATYPE_DEFINITION_ENUMERATION_stagedOrder[datatype_definition_enumeration]; ok {
		return order
	}
	if order, ok := stage.DATATYPE_DEFINITION_ENUMERATIONs_referenceOrder[datatype_definition_enumeration]; ok {
		return order
	} else {
		log.Printf("instance %p of type DATATYPE_DEFINITION_ENUMERATION was not staged and does not have a reference order", datatype_definition_enumeration)
		return 0
	}
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DATATYPE_DEFINITION_INTEGER_stagedOrder[datatype_definition_integer]; ok {
		return order
	}
	if order, ok := stage.DATATYPE_DEFINITION_INTEGERs_referenceOrder[datatype_definition_integer]; ok {
		return order
	} else {
		log.Printf("instance %p of type DATATYPE_DEFINITION_INTEGER was not staged and does not have a reference order", datatype_definition_integer)
		return 0
	}
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DATATYPE_DEFINITION_REAL_stagedOrder[datatype_definition_real]; ok {
		return order
	}
	if order, ok := stage.DATATYPE_DEFINITION_REALs_referenceOrder[datatype_definition_real]; ok {
		return order
	} else {
		log.Printf("instance %p of type DATATYPE_DEFINITION_REAL was not staged and does not have a reference order", datatype_definition_real)
		return 0
	}
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DATATYPE_DEFINITION_STRING_stagedOrder[datatype_definition_string]; ok {
		return order
	}
	if order, ok := stage.DATATYPE_DEFINITION_STRINGs_referenceOrder[datatype_definition_string]; ok {
		return order
	} else {
		log.Printf("instance %p of type DATATYPE_DEFINITION_STRING was not staged and does not have a reference order", datatype_definition_string)
		return 0
	}
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.DATATYPE_DEFINITION_XHTML_stagedOrder[datatype_definition_xhtml]; ok {
		return order
	}
	if order, ok := stage.DATATYPE_DEFINITION_XHTMLs_referenceOrder[datatype_definition_xhtml]; ok {
		return order
	} else {
		log.Printf("instance %p of type DATATYPE_DEFINITION_XHTML was not staged and does not have a reference order", datatype_definition_xhtml)
		return 0
	}
}

func (embedded_value *EMBEDDED_VALUE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EMBEDDED_VALUE_stagedOrder[embedded_value]; ok {
		return order
	}
	if order, ok := stage.EMBEDDED_VALUEs_referenceOrder[embedded_value]; ok {
		return order
	} else {
		log.Printf("instance %p of type EMBEDDED_VALUE was not staged and does not have a reference order", embedded_value)
		return 0
	}
}

func (enum_value *ENUM_VALUE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.ENUM_VALUE_stagedOrder[enum_value]; ok {
		return order
	}
	if order, ok := stage.ENUM_VALUEs_referenceOrder[enum_value]; ok {
		return order
	} else {
		log.Printf("instance %p of type ENUM_VALUE was not staged and does not have a reference order", enum_value)
		return 0
	}
}

func (embeddedjpgimage *EmbeddedJpgImage) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EmbeddedJpgImage_stagedOrder[embeddedjpgimage]; ok {
		return order
	}
	if order, ok := stage.EmbeddedJpgImages_referenceOrder[embeddedjpgimage]; ok {
		return order
	} else {
		log.Printf("instance %p of type EmbeddedJpgImage was not staged and does not have a reference order", embeddedjpgimage)
		return 0
	}
}

func (embeddedpngimage *EmbeddedPngImage) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EmbeddedPngImage_stagedOrder[embeddedpngimage]; ok {
		return order
	}
	if order, ok := stage.EmbeddedPngImages_referenceOrder[embeddedpngimage]; ok {
		return order
	} else {
		log.Printf("instance %p of type EmbeddedPngImage was not staged and does not have a reference order", embeddedpngimage)
		return 0
	}
}

func (embeddedsvgimage *EmbeddedSvgImage) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.EmbeddedSvgImage_stagedOrder[embeddedsvgimage]; ok {
		return order
	}
	if order, ok := stage.EmbeddedSvgImages_referenceOrder[embeddedsvgimage]; ok {
		return order
	} else {
		log.Printf("instance %p of type EmbeddedSvgImage was not staged and does not have a reference order", embeddedsvgimage)
		return 0
	}
}

func (kill *Kill) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Kill_stagedOrder[kill]; ok {
		return order
	}
	if order, ok := stage.Kills_referenceOrder[kill]; ok {
		return order
	} else {
		log.Printf("instance %p of type Kill was not staged and does not have a reference order", kill)
		return 0
	}
}

func (map_identifier_bool *Map_identifier_bool) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Map_identifier_bool_stagedOrder[map_identifier_bool]; ok {
		return order
	}
	if order, ok := stage.Map_identifier_bools_referenceOrder[map_identifier_bool]; ok {
		return order
	} else {
		log.Printf("instance %p of type Map_identifier_bool was not staged and does not have a reference order", map_identifier_bool)
		return 0
	}
}

func (relation_group *RELATION_GROUP) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RELATION_GROUP_stagedOrder[relation_group]; ok {
		return order
	}
	if order, ok := stage.RELATION_GROUPs_referenceOrder[relation_group]; ok {
		return order
	} else {
		log.Printf("instance %p of type RELATION_GROUP was not staged and does not have a reference order", relation_group)
		return 0
	}
}

func (relation_group_type *RELATION_GROUP_TYPE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.RELATION_GROUP_TYPE_stagedOrder[relation_group_type]; ok {
		return order
	}
	if order, ok := stage.RELATION_GROUP_TYPEs_referenceOrder[relation_group_type]; ok {
		return order
	} else {
		log.Printf("instance %p of type RELATION_GROUP_TYPE was not staged and does not have a reference order", relation_group_type)
		return 0
	}
}

func (req_if *REQ_IF) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.REQ_IF_stagedOrder[req_if]; ok {
		return order
	}
	if order, ok := stage.REQ_IFs_referenceOrder[req_if]; ok {
		return order
	} else {
		log.Printf("instance %p of type REQ_IF was not staged and does not have a reference order", req_if)
		return 0
	}
}

func (req_if_content *REQ_IF_CONTENT) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.REQ_IF_CONTENT_stagedOrder[req_if_content]; ok {
		return order
	}
	if order, ok := stage.REQ_IF_CONTENTs_referenceOrder[req_if_content]; ok {
		return order
	} else {
		log.Printf("instance %p of type REQ_IF_CONTENT was not staged and does not have a reference order", req_if_content)
		return 0
	}
}

func (req_if_header *REQ_IF_HEADER) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.REQ_IF_HEADER_stagedOrder[req_if_header]; ok {
		return order
	}
	if order, ok := stage.REQ_IF_HEADERs_referenceOrder[req_if_header]; ok {
		return order
	} else {
		log.Printf("instance %p of type REQ_IF_HEADER was not staged and does not have a reference order", req_if_header)
		return 0
	}
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.REQ_IF_TOOL_EXTENSION_stagedOrder[req_if_tool_extension]; ok {
		return order
	}
	if order, ok := stage.REQ_IF_TOOL_EXTENSIONs_referenceOrder[req_if_tool_extension]; ok {
		return order
	} else {
		log.Printf("instance %p of type REQ_IF_TOOL_EXTENSION was not staged and does not have a reference order", req_if_tool_extension)
		return 0
	}
}

func (specification *SPECIFICATION) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SPECIFICATION_stagedOrder[specification]; ok {
		return order
	}
	if order, ok := stage.SPECIFICATIONs_referenceOrder[specification]; ok {
		return order
	} else {
		log.Printf("instance %p of type SPECIFICATION was not staged and does not have a reference order", specification)
		return 0
	}
}

func (specification_rendering *SPECIFICATION_Rendering) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SPECIFICATION_Rendering_stagedOrder[specification_rendering]; ok {
		return order
	}
	if order, ok := stage.SPECIFICATION_Renderings_referenceOrder[specification_rendering]; ok {
		return order
	} else {
		log.Printf("instance %p of type SPECIFICATION_Rendering was not staged and does not have a reference order", specification_rendering)
		return 0
	}
}

func (specification_type *SPECIFICATION_TYPE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SPECIFICATION_TYPE_stagedOrder[specification_type]; ok {
		return order
	}
	if order, ok := stage.SPECIFICATION_TYPEs_referenceOrder[specification_type]; ok {
		return order
	} else {
		log.Printf("instance %p of type SPECIFICATION_TYPE was not staged and does not have a reference order", specification_type)
		return 0
	}
}

func (spec_hierarchy *SPEC_HIERARCHY) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SPEC_HIERARCHY_stagedOrder[spec_hierarchy]; ok {
		return order
	}
	if order, ok := stage.SPEC_HIERARCHYs_referenceOrder[spec_hierarchy]; ok {
		return order
	} else {
		log.Printf("instance %p of type SPEC_HIERARCHY was not staged and does not have a reference order", spec_hierarchy)
		return 0
	}
}

func (spec_object *SPEC_OBJECT) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SPEC_OBJECT_stagedOrder[spec_object]; ok {
		return order
	}
	if order, ok := stage.SPEC_OBJECTs_referenceOrder[spec_object]; ok {
		return order
	} else {
		log.Printf("instance %p of type SPEC_OBJECT was not staged and does not have a reference order", spec_object)
		return 0
	}
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SPEC_OBJECT_TYPE_stagedOrder[spec_object_type]; ok {
		return order
	}
	if order, ok := stage.SPEC_OBJECT_TYPEs_referenceOrder[spec_object_type]; ok {
		return order
	} else {
		log.Printf("instance %p of type SPEC_OBJECT_TYPE was not staged and does not have a reference order", spec_object_type)
		return 0
	}
}

func (spec_object_type_rendering *SPEC_OBJECT_TYPE_Rendering) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SPEC_OBJECT_TYPE_Rendering_stagedOrder[spec_object_type_rendering]; ok {
		return order
	}
	if order, ok := stage.SPEC_OBJECT_TYPE_Renderings_referenceOrder[spec_object_type_rendering]; ok {
		return order
	} else {
		log.Printf("instance %p of type SPEC_OBJECT_TYPE_Rendering was not staged and does not have a reference order", spec_object_type_rendering)
		return 0
	}
}

func (spec_relation *SPEC_RELATION) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SPEC_RELATION_stagedOrder[spec_relation]; ok {
		return order
	}
	if order, ok := stage.SPEC_RELATIONs_referenceOrder[spec_relation]; ok {
		return order
	} else {
		log.Printf("instance %p of type SPEC_RELATION was not staged and does not have a reference order", spec_relation)
		return 0
	}
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.SPEC_RELATION_TYPE_stagedOrder[spec_relation_type]; ok {
		return order
	}
	if order, ok := stage.SPEC_RELATION_TYPEs_referenceOrder[spec_relation_type]; ok {
		return order
	} else {
		log.Printf("instance %p of type SPEC_RELATION_TYPE was not staged and does not have a reference order", spec_relation_type)
		return 0
	}
}

func (staticwebsite *StaticWebSite) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StaticWebSite_stagedOrder[staticwebsite]; ok {
		return order
	}
	if order, ok := stage.StaticWebSites_referenceOrder[staticwebsite]; ok {
		return order
	} else {
		log.Printf("instance %p of type StaticWebSite was not staged and does not have a reference order", staticwebsite)
		return 0
	}
}

func (staticwebsitechapter *StaticWebSiteChapter) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StaticWebSiteChapter_stagedOrder[staticwebsitechapter]; ok {
		return order
	}
	if order, ok := stage.StaticWebSiteChapters_referenceOrder[staticwebsitechapter]; ok {
		return order
	} else {
		log.Printf("instance %p of type StaticWebSiteChapter was not staged and does not have a reference order", staticwebsitechapter)
		return 0
	}
}

func (staticwebsitegeneratedimage *StaticWebSiteGeneratedImage) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StaticWebSiteGeneratedImage_stagedOrder[staticwebsitegeneratedimage]; ok {
		return order
	}
	if order, ok := stage.StaticWebSiteGeneratedImages_referenceOrder[staticwebsitegeneratedimage]; ok {
		return order
	} else {
		log.Printf("instance %p of type StaticWebSiteGeneratedImage was not staged and does not have a reference order", staticwebsitegeneratedimage)
		return 0
	}
}

func (staticwebsiteimage *StaticWebSiteImage) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StaticWebSiteImage_stagedOrder[staticwebsiteimage]; ok {
		return order
	}
	if order, ok := stage.StaticWebSiteImages_referenceOrder[staticwebsiteimage]; ok {
		return order
	} else {
		log.Printf("instance %p of type StaticWebSiteImage was not staged and does not have a reference order", staticwebsiteimage)
		return 0
	}
}

func (staticwebsiteparagraph *StaticWebSiteParagraph) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.StaticWebSiteParagraph_stagedOrder[staticwebsiteparagraph]; ok {
		return order
	}
	if order, ok := stage.StaticWebSiteParagraphs_referenceOrder[staticwebsiteparagraph]; ok {
		return order
	} else {
		log.Printf("instance %p of type StaticWebSiteParagraph was not staged and does not have a reference order", staticwebsiteparagraph)
		return 0
	}
}

func (xhtml_content *XHTML_CONTENT) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.XHTML_CONTENT_stagedOrder[xhtml_content]; ok {
		return order
	}
	if order, ok := stage.XHTML_CONTENTs_referenceOrder[xhtml_content]; ok {
		return order
	} else {
		log.Printf("instance %p of type XHTML_CONTENT was not staged and does not have a reference order", xhtml_content)
		return 0
	}
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
	return fmt.Sprintf("__%s__%08d_", alternative_id.GongGetGongstructName(), alternative_id.GongGetOrder(stage))
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_boolean.GongGetGongstructName(), attribute_definition_boolean.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_boolean.GongGetGongstructName(), attribute_definition_boolean.GongGetOrder(stage))
}

func (attribute_definition_boolean_rendering *ATTRIBUTE_DEFINITION_BOOLEAN_Rendering) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_boolean_rendering.GongGetGongstructName(), attribute_definition_boolean_rendering.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_boolean_rendering *ATTRIBUTE_DEFINITION_BOOLEAN_Rendering) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_boolean_rendering.GongGetGongstructName(), attribute_definition_boolean_rendering.GongGetOrder(stage))
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_date.GongGetGongstructName(), attribute_definition_date.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_date.GongGetGongstructName(), attribute_definition_date.GongGetOrder(stage))
}

func (attribute_definition_date_rendering *ATTRIBUTE_DEFINITION_DATE_Rendering) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_date_rendering.GongGetGongstructName(), attribute_definition_date_rendering.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_date_rendering *ATTRIBUTE_DEFINITION_DATE_Rendering) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_date_rendering.GongGetGongstructName(), attribute_definition_date_rendering.GongGetOrder(stage))
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_enumeration.GongGetGongstructName(), attribute_definition_enumeration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_enumeration.GongGetGongstructName(), attribute_definition_enumeration.GongGetOrder(stage))
}

func (attribute_definition_enumeration_rendering *ATTRIBUTE_DEFINITION_ENUMERATION_Rendering) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_enumeration_rendering.GongGetGongstructName(), attribute_definition_enumeration_rendering.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_enumeration_rendering *ATTRIBUTE_DEFINITION_ENUMERATION_Rendering) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_enumeration_rendering.GongGetGongstructName(), attribute_definition_enumeration_rendering.GongGetOrder(stage))
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_integer.GongGetGongstructName(), attribute_definition_integer.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_integer.GongGetGongstructName(), attribute_definition_integer.GongGetOrder(stage))
}

func (attribute_definition_integer_rendering *ATTRIBUTE_DEFINITION_INTEGER_Rendering) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_integer_rendering.GongGetGongstructName(), attribute_definition_integer_rendering.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_integer_rendering *ATTRIBUTE_DEFINITION_INTEGER_Rendering) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_integer_rendering.GongGetGongstructName(), attribute_definition_integer_rendering.GongGetOrder(stage))
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_real.GongGetGongstructName(), attribute_definition_real.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_real.GongGetGongstructName(), attribute_definition_real.GongGetOrder(stage))
}

func (attribute_definition_real_rendering *ATTRIBUTE_DEFINITION_REAL_Rendering) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_real_rendering.GongGetGongstructName(), attribute_definition_real_rendering.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_real_rendering *ATTRIBUTE_DEFINITION_REAL_Rendering) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_real_rendering.GongGetGongstructName(), attribute_definition_real_rendering.GongGetOrder(stage))
}

func (attribute_definition_rendering *ATTRIBUTE_DEFINITION_Rendering) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_rendering.GongGetGongstructName(), attribute_definition_rendering.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_rendering *ATTRIBUTE_DEFINITION_Rendering) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_rendering.GongGetGongstructName(), attribute_definition_rendering.GongGetOrder(stage))
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_string.GongGetGongstructName(), attribute_definition_string.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_string.GongGetGongstructName(), attribute_definition_string.GongGetOrder(stage))
}

func (attribute_definition_string_rendering *ATTRIBUTE_DEFINITION_STRING_Rendering) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_string_rendering.GongGetGongstructName(), attribute_definition_string_rendering.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_string_rendering *ATTRIBUTE_DEFINITION_STRING_Rendering) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_string_rendering.GongGetGongstructName(), attribute_definition_string_rendering.GongGetOrder(stage))
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_xhtml.GongGetGongstructName(), attribute_definition_xhtml.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_xhtml.GongGetGongstructName(), attribute_definition_xhtml.GongGetOrder(stage))
}

func (attribute_definition_xhtml_rendering *ATTRIBUTE_DEFINITION_XHTML_Rendering) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_xhtml_rendering.GongGetGongstructName(), attribute_definition_xhtml_rendering.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_definition_xhtml_rendering *ATTRIBUTE_DEFINITION_XHTML_Rendering) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_definition_xhtml_rendering.GongGetGongstructName(), attribute_definition_xhtml_rendering.GongGetOrder(stage))
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_boolean.GongGetGongstructName(), attribute_value_boolean.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_boolean.GongGetGongstructName(), attribute_value_boolean.GongGetOrder(stage))
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_date.GongGetGongstructName(), attribute_value_date.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_date.GongGetGongstructName(), attribute_value_date.GongGetOrder(stage))
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_enumeration.GongGetGongstructName(), attribute_value_enumeration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_enumeration.GongGetGongstructName(), attribute_value_enumeration.GongGetOrder(stage))
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_integer.GongGetGongstructName(), attribute_value_integer.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_integer.GongGetGongstructName(), attribute_value_integer.GongGetOrder(stage))
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_real.GongGetGongstructName(), attribute_value_real.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_real.GongGetGongstructName(), attribute_value_real.GongGetOrder(stage))
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_string.GongGetGongstructName(), attribute_value_string.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_string.GongGetGongstructName(), attribute_value_string.GongGetOrder(stage))
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_xhtml.GongGetGongstructName(), attribute_value_xhtml.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attribute_value_xhtml.GongGetGongstructName(), attribute_value_xhtml.GongGetOrder(stage))
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_alternative_id.GongGetGongstructName(), a_alternative_id.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_alternative_id *A_ALTERNATIVE_ID) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_alternative_id.GongGetGongstructName(), a_alternative_id.GongGetOrder(stage))
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_boolean_ref.GongGetGongstructName(), a_attribute_definition_boolean_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_boolean_ref.GongGetGongstructName(), a_attribute_definition_boolean_ref.GongGetOrder(stage))
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_date_ref.GongGetGongstructName(), a_attribute_definition_date_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_date_ref.GongGetGongstructName(), a_attribute_definition_date_ref.GongGetOrder(stage))
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_enumeration_ref.GongGetGongstructName(), a_attribute_definition_enumeration_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_enumeration_ref.GongGetGongstructName(), a_attribute_definition_enumeration_ref.GongGetOrder(stage))
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_integer_ref.GongGetGongstructName(), a_attribute_definition_integer_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_integer_ref.GongGetGongstructName(), a_attribute_definition_integer_ref.GongGetOrder(stage))
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_real_ref.GongGetGongstructName(), a_attribute_definition_real_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_real_ref.GongGetGongstructName(), a_attribute_definition_real_ref.GongGetOrder(stage))
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_string_ref.GongGetGongstructName(), a_attribute_definition_string_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_string_ref.GongGetGongstructName(), a_attribute_definition_string_ref.GongGetOrder(stage))
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_xhtml_ref.GongGetGongstructName(), a_attribute_definition_xhtml_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_definition_xhtml_ref.GongGetGongstructName(), a_attribute_definition_xhtml_ref.GongGetOrder(stage))
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_boolean.GongGetGongstructName(), a_attribute_value_boolean.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_boolean.GongGetGongstructName(), a_attribute_value_boolean.GongGetOrder(stage))
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_date.GongGetGongstructName(), a_attribute_value_date.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_date.GongGetGongstructName(), a_attribute_value_date.GongGetOrder(stage))
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_enumeration.GongGetGongstructName(), a_attribute_value_enumeration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_enumeration.GongGetGongstructName(), a_attribute_value_enumeration.GongGetOrder(stage))
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_integer.GongGetGongstructName(), a_attribute_value_integer.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_integer.GongGetGongstructName(), a_attribute_value_integer.GongGetOrder(stage))
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_real.GongGetGongstructName(), a_attribute_value_real.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_real.GongGetGongstructName(), a_attribute_value_real.GongGetOrder(stage))
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_string.GongGetGongstructName(), a_attribute_value_string.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_string.GongGetGongstructName(), a_attribute_value_string.GongGetOrder(stage))
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_xhtml.GongGetGongstructName(), a_attribute_value_xhtml.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_xhtml.GongGetGongstructName(), a_attribute_value_xhtml.GongGetOrder(stage))
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_xhtml_1.GongGetGongstructName(), a_attribute_value_xhtml_1.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_attribute_value_xhtml_1.GongGetGongstructName(), a_attribute_value_xhtml_1.GongGetOrder(stage))
}

func (a_children *A_CHILDREN) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_children.GongGetGongstructName(), a_children.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_children *A_CHILDREN) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_children.GongGetGongstructName(), a_children.GongGetOrder(stage))
}

func (a_core_content *A_CORE_CONTENT) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_core_content.GongGetGongstructName(), a_core_content.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_core_content *A_CORE_CONTENT) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_core_content.GongGetGongstructName(), a_core_content.GongGetOrder(stage))
}

func (a_datatypes *A_DATATYPES) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatypes.GongGetGongstructName(), a_datatypes.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatypes *A_DATATYPES) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatypes.GongGetGongstructName(), a_datatypes.GongGetOrder(stage))
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_boolean_ref.GongGetGongstructName(), a_datatype_definition_boolean_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_boolean_ref.GongGetGongstructName(), a_datatype_definition_boolean_ref.GongGetOrder(stage))
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_date_ref.GongGetGongstructName(), a_datatype_definition_date_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_date_ref.GongGetGongstructName(), a_datatype_definition_date_ref.GongGetOrder(stage))
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_enumeration_ref.GongGetGongstructName(), a_datatype_definition_enumeration_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_enumeration_ref.GongGetGongstructName(), a_datatype_definition_enumeration_ref.GongGetOrder(stage))
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_integer_ref.GongGetGongstructName(), a_datatype_definition_integer_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_integer_ref.GongGetGongstructName(), a_datatype_definition_integer_ref.GongGetOrder(stage))
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_real_ref.GongGetGongstructName(), a_datatype_definition_real_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_real_ref.GongGetGongstructName(), a_datatype_definition_real_ref.GongGetOrder(stage))
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_string_ref.GongGetGongstructName(), a_datatype_definition_string_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_string_ref.GongGetGongstructName(), a_datatype_definition_string_ref.GongGetOrder(stage))
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_xhtml_ref.GongGetGongstructName(), a_datatype_definition_xhtml_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_datatype_definition_xhtml_ref.GongGetGongstructName(), a_datatype_definition_xhtml_ref.GongGetOrder(stage))
}

func (a_editable_atts *A_EDITABLE_ATTS) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_editable_atts.GongGetGongstructName(), a_editable_atts.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_editable_atts *A_EDITABLE_ATTS) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_editable_atts.GongGetGongstructName(), a_editable_atts.GongGetOrder(stage))
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_enum_value_ref.GongGetGongstructName(), a_enum_value_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_enum_value_ref *A_ENUM_VALUE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_enum_value_ref.GongGetGongstructName(), a_enum_value_ref.GongGetOrder(stage))
}

func (a_object *A_OBJECT) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_object.GongGetGongstructName(), a_object.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_object *A_OBJECT) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_object.GongGetGongstructName(), a_object.GongGetOrder(stage))
}

func (a_properties *A_PROPERTIES) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_properties.GongGetGongstructName(), a_properties.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_properties *A_PROPERTIES) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_properties.GongGetGongstructName(), a_properties.GongGetOrder(stage))
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_relation_group_type_ref.GongGetGongstructName(), a_relation_group_type_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_relation_group_type_ref.GongGetGongstructName(), a_relation_group_type_ref.GongGetOrder(stage))
}

func (a_source_1 *A_SOURCE_1) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_source_1.GongGetGongstructName(), a_source_1.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_source_1 *A_SOURCE_1) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_source_1.GongGetGongstructName(), a_source_1.GongGetOrder(stage))
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_source_specification_1.GongGetGongstructName(), a_source_specification_1.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_source_specification_1.GongGetGongstructName(), a_source_specification_1.GongGetOrder(stage))
}

func (a_specifications *A_SPECIFICATIONS) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_specifications.GongGetGongstructName(), a_specifications.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_specifications *A_SPECIFICATIONS) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_specifications.GongGetGongstructName(), a_specifications.GongGetOrder(stage))
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_specification_type_ref.GongGetGongstructName(), a_specification_type_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_specification_type_ref.GongGetGongstructName(), a_specification_type_ref.GongGetOrder(stage))
}

func (a_specified_values *A_SPECIFIED_VALUES) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_specified_values.GongGetGongstructName(), a_specified_values.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_specified_values *A_SPECIFIED_VALUES) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_specified_values.GongGetGongstructName(), a_specified_values.GongGetOrder(stage))
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_attributes.GongGetGongstructName(), a_spec_attributes.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_attributes.GongGetGongstructName(), a_spec_attributes.GongGetOrder(stage))
}

func (a_spec_objects *A_SPEC_OBJECTS) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_objects.GongGetGongstructName(), a_spec_objects.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_objects *A_SPEC_OBJECTS) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_objects.GongGetGongstructName(), a_spec_objects.GongGetOrder(stage))
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_object_type_ref.GongGetGongstructName(), a_spec_object_type_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_object_type_ref.GongGetGongstructName(), a_spec_object_type_ref.GongGetOrder(stage))
}

func (a_spec_relations *A_SPEC_RELATIONS) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_relations.GongGetGongstructName(), a_spec_relations.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_relations *A_SPEC_RELATIONS) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_relations.GongGetGongstructName(), a_spec_relations.GongGetOrder(stage))
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_relation_groups.GongGetGongstructName(), a_spec_relation_groups.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_relation_groups.GongGetGongstructName(), a_spec_relation_groups.GongGetOrder(stage))
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_relation_ref.GongGetGongstructName(), a_spec_relation_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_relation_ref.GongGetGongstructName(), a_spec_relation_ref.GongGetOrder(stage))
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_relation_type_ref.GongGetGongstructName(), a_spec_relation_type_ref.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_relation_type_ref.GongGetGongstructName(), a_spec_relation_type_ref.GongGetOrder(stage))
}

func (a_spec_types *A_SPEC_TYPES) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_types.GongGetGongstructName(), a_spec_types.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_spec_types *A_SPEC_TYPES) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_spec_types.GongGetGongstructName(), a_spec_types.GongGetOrder(stage))
}

func (a_the_header *A_THE_HEADER) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_the_header.GongGetGongstructName(), a_the_header.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_the_header *A_THE_HEADER) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_the_header.GongGetGongstructName(), a_the_header.GongGetOrder(stage))
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_tool_extensions.GongGetGongstructName(), a_tool_extensions.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_tool_extensions *A_TOOL_EXTENSIONS) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_tool_extensions.GongGetGongstructName(), a_tool_extensions.GongGetOrder(stage))
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_boolean.GongGetGongstructName(), datatype_definition_boolean.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_boolean.GongGetGongstructName(), datatype_definition_boolean.GongGetOrder(stage))
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_date.GongGetGongstructName(), datatype_definition_date.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_date.GongGetGongstructName(), datatype_definition_date.GongGetOrder(stage))
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_enumeration.GongGetGongstructName(), datatype_definition_enumeration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_enumeration.GongGetGongstructName(), datatype_definition_enumeration.GongGetOrder(stage))
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_integer.GongGetGongstructName(), datatype_definition_integer.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_integer.GongGetGongstructName(), datatype_definition_integer.GongGetOrder(stage))
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_real.GongGetGongstructName(), datatype_definition_real.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_real.GongGetGongstructName(), datatype_definition_real.GongGetOrder(stage))
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_string.GongGetGongstructName(), datatype_definition_string.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_string.GongGetGongstructName(), datatype_definition_string.GongGetOrder(stage))
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_xhtml.GongGetGongstructName(), datatype_definition_xhtml.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", datatype_definition_xhtml.GongGetGongstructName(), datatype_definition_xhtml.GongGetOrder(stage))
}

func (embedded_value *EMBEDDED_VALUE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", embedded_value.GongGetGongstructName(), embedded_value.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (embedded_value *EMBEDDED_VALUE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", embedded_value.GongGetGongstructName(), embedded_value.GongGetOrder(stage))
}

func (enum_value *ENUM_VALUE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", enum_value.GongGetGongstructName(), enum_value.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (enum_value *ENUM_VALUE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", enum_value.GongGetGongstructName(), enum_value.GongGetOrder(stage))
}

func (embeddedjpgimage *EmbeddedJpgImage) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", embeddedjpgimage.GongGetGongstructName(), embeddedjpgimage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (embeddedjpgimage *EmbeddedJpgImage) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", embeddedjpgimage.GongGetGongstructName(), embeddedjpgimage.GongGetOrder(stage))
}

func (embeddedpngimage *EmbeddedPngImage) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", embeddedpngimage.GongGetGongstructName(), embeddedpngimage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (embeddedpngimage *EmbeddedPngImage) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", embeddedpngimage.GongGetGongstructName(), embeddedpngimage.GongGetOrder(stage))
}

func (embeddedsvgimage *EmbeddedSvgImage) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", embeddedsvgimage.GongGetGongstructName(), embeddedsvgimage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (embeddedsvgimage *EmbeddedSvgImage) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", embeddedsvgimage.GongGetGongstructName(), embeddedsvgimage.GongGetOrder(stage))
}

func (kill *Kill) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", kill.GongGetGongstructName(), kill.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (kill *Kill) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", kill.GongGetGongstructName(), kill.GongGetOrder(stage))
}

func (map_identifier_bool *Map_identifier_bool) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", map_identifier_bool.GongGetGongstructName(), map_identifier_bool.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (map_identifier_bool *Map_identifier_bool) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", map_identifier_bool.GongGetGongstructName(), map_identifier_bool.GongGetOrder(stage))
}

func (relation_group *RELATION_GROUP) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", relation_group.GongGetGongstructName(), relation_group.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (relation_group *RELATION_GROUP) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", relation_group.GongGetGongstructName(), relation_group.GongGetOrder(stage))
}

func (relation_group_type *RELATION_GROUP_TYPE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", relation_group_type.GongGetGongstructName(), relation_group_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (relation_group_type *RELATION_GROUP_TYPE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", relation_group_type.GongGetGongstructName(), relation_group_type.GongGetOrder(stage))
}

func (req_if *REQ_IF) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", req_if.GongGetGongstructName(), req_if.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (req_if *REQ_IF) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", req_if.GongGetGongstructName(), req_if.GongGetOrder(stage))
}

func (req_if_content *REQ_IF_CONTENT) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", req_if_content.GongGetGongstructName(), req_if_content.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (req_if_content *REQ_IF_CONTENT) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", req_if_content.GongGetGongstructName(), req_if_content.GongGetOrder(stage))
}

func (req_if_header *REQ_IF_HEADER) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", req_if_header.GongGetGongstructName(), req_if_header.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (req_if_header *REQ_IF_HEADER) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", req_if_header.GongGetGongstructName(), req_if_header.GongGetOrder(stage))
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", req_if_tool_extension.GongGetGongstructName(), req_if_tool_extension.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", req_if_tool_extension.GongGetGongstructName(), req_if_tool_extension.GongGetOrder(stage))
}

func (specification *SPECIFICATION) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", specification.GongGetGongstructName(), specification.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specification *SPECIFICATION) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", specification.GongGetGongstructName(), specification.GongGetOrder(stage))
}

func (specification_rendering *SPECIFICATION_Rendering) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", specification_rendering.GongGetGongstructName(), specification_rendering.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specification_rendering *SPECIFICATION_Rendering) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", specification_rendering.GongGetGongstructName(), specification_rendering.GongGetOrder(stage))
}

func (specification_type *SPECIFICATION_TYPE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", specification_type.GongGetGongstructName(), specification_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (specification_type *SPECIFICATION_TYPE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", specification_type.GongGetGongstructName(), specification_type.GongGetOrder(stage))
}

func (spec_hierarchy *SPEC_HIERARCHY) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_hierarchy.GongGetGongstructName(), spec_hierarchy.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spec_hierarchy *SPEC_HIERARCHY) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_hierarchy.GongGetGongstructName(), spec_hierarchy.GongGetOrder(stage))
}

func (spec_object *SPEC_OBJECT) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_object.GongGetGongstructName(), spec_object.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spec_object *SPEC_OBJECT) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_object.GongGetGongstructName(), spec_object.GongGetOrder(stage))
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_object_type.GongGetGongstructName(), spec_object_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spec_object_type *SPEC_OBJECT_TYPE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_object_type.GongGetGongstructName(), spec_object_type.GongGetOrder(stage))
}

func (spec_object_type_rendering *SPEC_OBJECT_TYPE_Rendering) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_object_type_rendering.GongGetGongstructName(), spec_object_type_rendering.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spec_object_type_rendering *SPEC_OBJECT_TYPE_Rendering) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_object_type_rendering.GongGetGongstructName(), spec_object_type_rendering.GongGetOrder(stage))
}

func (spec_relation *SPEC_RELATION) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_relation.GongGetGongstructName(), spec_relation.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spec_relation *SPEC_RELATION) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_relation.GongGetGongstructName(), spec_relation.GongGetOrder(stage))
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_relation_type.GongGetGongstructName(), spec_relation_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (spec_relation_type *SPEC_RELATION_TYPE) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", spec_relation_type.GongGetGongstructName(), spec_relation_type.GongGetOrder(stage))
}

func (staticwebsite *StaticWebSite) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staticwebsite.GongGetGongstructName(), staticwebsite.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staticwebsite *StaticWebSite) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staticwebsite.GongGetGongstructName(), staticwebsite.GongGetOrder(stage))
}

func (staticwebsitechapter *StaticWebSiteChapter) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staticwebsitechapter.GongGetGongstructName(), staticwebsitechapter.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staticwebsitechapter *StaticWebSiteChapter) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staticwebsitechapter.GongGetGongstructName(), staticwebsitechapter.GongGetOrder(stage))
}

func (staticwebsitegeneratedimage *StaticWebSiteGeneratedImage) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staticwebsitegeneratedimage.GongGetGongstructName(), staticwebsitegeneratedimage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staticwebsitegeneratedimage *StaticWebSiteGeneratedImage) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staticwebsitegeneratedimage.GongGetGongstructName(), staticwebsitegeneratedimage.GongGetOrder(stage))
}

func (staticwebsiteimage *StaticWebSiteImage) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staticwebsiteimage.GongGetGongstructName(), staticwebsiteimage.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staticwebsiteimage *StaticWebSiteImage) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staticwebsiteimage.GongGetGongstructName(), staticwebsiteimage.GongGetOrder(stage))
}

func (staticwebsiteparagraph *StaticWebSiteParagraph) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staticwebsiteparagraph.GongGetGongstructName(), staticwebsiteparagraph.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staticwebsiteparagraph *StaticWebSiteParagraph) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staticwebsiteparagraph.GongGetGongstructName(), staticwebsiteparagraph.GongGetOrder(stage))
}

func (xhtml_content *XHTML_CONTENT) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xhtml_content.GongGetGongstructName(), xhtml_content.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (xhtml_content *XHTML_CONTENT) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", xhtml_content.GongGetGongstructName(), xhtml_content.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (alternative_id *ALTERNATIVE_ID) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", alternative_id.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ALTERNATIVE_ID")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(alternative_id.Name))
	return
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_boolean.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_BOOLEAN")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_definition_boolean.Name))
	return
}

func (attribute_definition_boolean_rendering *ATTRIBUTE_DEFINITION_BOOLEAN_Rendering) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_boolean_rendering.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_BOOLEAN_Rendering")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_definition_boolean_rendering.Name))
	return
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_date.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_DATE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_definition_date.Name))
	return
}

func (attribute_definition_date_rendering *ATTRIBUTE_DEFINITION_DATE_Rendering) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_date_rendering.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_DATE_Rendering")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_definition_date_rendering.Name))
	return
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_enumeration.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_ENUMERATION")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_definition_enumeration.Name))
	return
}

func (attribute_definition_enumeration_rendering *ATTRIBUTE_DEFINITION_ENUMERATION_Rendering) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_enumeration_rendering.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_ENUMERATION_Rendering")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_definition_enumeration_rendering.Name))
	return
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_integer.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_INTEGER")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_definition_integer.Name))
	return
}

func (attribute_definition_integer_rendering *ATTRIBUTE_DEFINITION_INTEGER_Rendering) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_integer_rendering.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_INTEGER_Rendering")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_definition_integer_rendering.Name))
	return
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_real.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_REAL")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_definition_real.Name))
	return
}

func (attribute_definition_real_rendering *ATTRIBUTE_DEFINITION_REAL_Rendering) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_real_rendering.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_REAL_Rendering")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_definition_real_rendering.Name))
	return
}

func (attribute_definition_rendering *ATTRIBUTE_DEFINITION_Rendering) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_rendering.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_Rendering")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_definition_rendering.Name))
	return
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_string.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_STRING")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_definition_string.Name))
	return
}

func (attribute_definition_string_rendering *ATTRIBUTE_DEFINITION_STRING_Rendering) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_string_rendering.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_STRING_Rendering")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_definition_string_rendering.Name))
	return
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_xhtml.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_XHTML")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_definition_xhtml.Name))
	return
}

func (attribute_definition_xhtml_rendering *ATTRIBUTE_DEFINITION_XHTML_Rendering) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_xhtml_rendering.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_DEFINITION_XHTML_Rendering")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_definition_xhtml_rendering.Name))
	return
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_boolean.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_BOOLEAN")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_value_boolean.Name))
	return
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_date.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_DATE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_value_date.Name))
	return
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_enumeration.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_ENUMERATION")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_value_enumeration.Name))
	return
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_integer.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_INTEGER")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_value_integer.Name))
	return
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_real.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_REAL")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_value_real.Name))
	return
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_string.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_STRING")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_value_string.Name))
	return
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_value_xhtml.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ATTRIBUTE_VALUE_XHTML")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attribute_value_xhtml.Name))
	return
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_alternative_id.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ALTERNATIVE_ID")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_alternative_id.Name))
	return
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_boolean_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_attribute_definition_boolean_ref.Name))
	return
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_date_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_DATE_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_attribute_definition_date_ref.Name))
	return
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_enumeration_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_attribute_definition_enumeration_ref.Name))
	return
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_integer_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_INTEGER_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_attribute_definition_integer_ref.Name))
	return
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_real_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_REAL_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_attribute_definition_real_ref.Name))
	return
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_string_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_STRING_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_attribute_definition_string_ref.Name))
	return
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_definition_xhtml_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_DEFINITION_XHTML_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_attribute_definition_xhtml_ref.Name))
	return
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_boolean.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_BOOLEAN")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_attribute_value_boolean.Name))
	return
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_date.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_DATE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_attribute_value_date.Name))
	return
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_enumeration.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_ENUMERATION")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_attribute_value_enumeration.Name))
	return
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_integer.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_INTEGER")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_attribute_value_integer.Name))
	return
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_real.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_REAL")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_attribute_value_real.Name))
	return
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_string.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_STRING")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_attribute_value_string.Name))
	return
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_xhtml.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_XHTML")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_attribute_value_xhtml.Name))
	return
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_attribute_value_xhtml_1.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ATTRIBUTE_VALUE_XHTML_1")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_attribute_value_xhtml_1.Name))
	return
}

func (a_children *A_CHILDREN) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_children.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_CHILDREN")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_children.Name))
	return
}

func (a_core_content *A_CORE_CONTENT) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_core_content.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_CORE_CONTENT")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_core_content.Name))
	return
}

func (a_datatypes *A_DATATYPES) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatypes.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPES")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_datatypes.Name))
	return
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_boolean_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_BOOLEAN_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_datatype_definition_boolean_ref.Name))
	return
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_date_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_DATE_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_datatype_definition_date_ref.Name))
	return
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_enumeration_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_ENUMERATION_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_datatype_definition_enumeration_ref.Name))
	return
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_integer_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_INTEGER_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_datatype_definition_integer_ref.Name))
	return
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_real_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_REAL_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_datatype_definition_real_ref.Name))
	return
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_string_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_STRING_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_datatype_definition_string_ref.Name))
	return
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_datatype_definition_xhtml_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_DATATYPE_DEFINITION_XHTML_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_datatype_definition_xhtml_ref.Name))
	return
}

func (a_editable_atts *A_EDITABLE_ATTS) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_editable_atts.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_EDITABLE_ATTS")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_editable_atts.Name))
	return
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_enum_value_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_ENUM_VALUE_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_enum_value_ref.Name))
	return
}

func (a_object *A_OBJECT) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_object.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_OBJECT")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_object.Name))
	return
}

func (a_properties *A_PROPERTIES) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_properties.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_PROPERTIES")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_properties.Name))
	return
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_relation_group_type_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_RELATION_GROUP_TYPE_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_relation_group_type_ref.Name))
	return
}

func (a_source_1 *A_SOURCE_1) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_source_1.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SOURCE_1")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_source_1.Name))
	return
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_source_specification_1.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SOURCE_SPECIFICATION_1")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_source_specification_1.Name))
	return
}

func (a_specifications *A_SPECIFICATIONS) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_specifications.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPECIFICATIONS")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_specifications.Name))
	return
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_specification_type_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPECIFICATION_TYPE_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_specification_type_ref.Name))
	return
}

func (a_specified_values *A_SPECIFIED_VALUES) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_specified_values.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPECIFIED_VALUES")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_specified_values.Name))
	return
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_attributes.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_ATTRIBUTES")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_spec_attributes.Name))
	return
}

func (a_spec_objects *A_SPEC_OBJECTS) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_objects.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_OBJECTS")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_spec_objects.Name))
	return
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_object_type_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_OBJECT_TYPE_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_spec_object_type_ref.Name))
	return
}

func (a_spec_relations *A_SPEC_RELATIONS) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_relations.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_RELATIONS")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_spec_relations.Name))
	return
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_relation_groups.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_RELATION_GROUPS")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_spec_relation_groups.Name))
	return
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_relation_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_RELATION_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_spec_relation_ref.Name))
	return
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_relation_type_ref.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_RELATION_TYPE_REF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_spec_relation_type_ref.Name))
	return
}

func (a_spec_types *A_SPEC_TYPES) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_spec_types.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_SPEC_TYPES")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_spec_types.Name))
	return
}

func (a_the_header *A_THE_HEADER) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_the_header.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_THE_HEADER")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_the_header.Name))
	return
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_tool_extensions.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_TOOL_EXTENSIONS")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_tool_extensions.Name))
	return
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_boolean.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_BOOLEAN")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(datatype_definition_boolean.Name))
	return
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_date.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_DATE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(datatype_definition_date.Name))
	return
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_enumeration.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_ENUMERATION")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(datatype_definition_enumeration.Name))
	return
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_integer.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_INTEGER")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(datatype_definition_integer.Name))
	return
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_real.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_REAL")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(datatype_definition_real.Name))
	return
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_string.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_STRING")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(datatype_definition_string.Name))
	return
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", datatype_definition_xhtml.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "DATATYPE_DEFINITION_XHTML")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(datatype_definition_xhtml.Name))
	return
}

func (embedded_value *EMBEDDED_VALUE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", embedded_value.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EMBEDDED_VALUE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(embedded_value.Name))
	return
}

func (enum_value *ENUM_VALUE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", enum_value.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "ENUM_VALUE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(enum_value.Name))
	return
}

func (embeddedjpgimage *EmbeddedJpgImage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", embeddedjpgimage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EmbeddedJpgImage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(embeddedjpgimage.Name))
	return
}

func (embeddedpngimage *EmbeddedPngImage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", embeddedpngimage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EmbeddedPngImage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(embeddedpngimage.Name))
	return
}

func (embeddedsvgimage *EmbeddedSvgImage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", embeddedsvgimage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "EmbeddedSvgImage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(embeddedsvgimage.Name))
	return
}

func (kill *Kill) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", kill.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Kill")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(kill.Name))
	return
}

func (map_identifier_bool *Map_identifier_bool) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", map_identifier_bool.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Map_identifier_bool")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(map_identifier_bool.Name))
	return
}

func (relation_group *RELATION_GROUP) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", relation_group.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RELATION_GROUP")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(relation_group.Name))
	return
}

func (relation_group_type *RELATION_GROUP_TYPE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", relation_group_type.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "RELATION_GROUP_TYPE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(relation_group_type.Name))
	return
}

func (req_if *REQ_IF) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", req_if.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQ_IF")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(req_if.Name))
	return
}

func (req_if_content *REQ_IF_CONTENT) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", req_if_content.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQ_IF_CONTENT")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(req_if_content.Name))
	return
}

func (req_if_header *REQ_IF_HEADER) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", req_if_header.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQ_IF_HEADER")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(req_if_header.Name))
	return
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", req_if_tool_extension.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "REQ_IF_TOOL_EXTENSION")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(req_if_tool_extension.Name))
	return
}

func (specification *SPECIFICATION) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", specification.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECIFICATION")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(specification.Name))
	return
}

func (specification_rendering *SPECIFICATION_Rendering) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", specification_rendering.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECIFICATION_Rendering")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(specification_rendering.Name))
	return
}

func (specification_type *SPECIFICATION_TYPE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", specification_type.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPECIFICATION_TYPE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(specification_type.Name))
	return
}

func (spec_hierarchy *SPEC_HIERARCHY) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_hierarchy.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_HIERARCHY")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(spec_hierarchy.Name))
	return
}

func (spec_object *SPEC_OBJECT) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_object.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_OBJECT")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(spec_object.Name))
	return
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_object_type.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_OBJECT_TYPE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(spec_object_type.Name))
	return
}

func (spec_object_type_rendering *SPEC_OBJECT_TYPE_Rendering) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_object_type_rendering.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_OBJECT_TYPE_Rendering")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(spec_object_type_rendering.Name))
	return
}

func (spec_relation *SPEC_RELATION) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_relation.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_RELATION")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(spec_relation.Name))
	return
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_relation_type.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "SPEC_RELATION_TYPE")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(spec_relation_type.Name))
	return
}

func (staticwebsite *StaticWebSite) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staticwebsite.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StaticWebSite")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(staticwebsite.Name))
	return
}

func (staticwebsitechapter *StaticWebSiteChapter) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staticwebsitechapter.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StaticWebSiteChapter")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(staticwebsitechapter.Name))
	return
}

func (staticwebsitegeneratedimage *StaticWebSiteGeneratedImage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staticwebsitegeneratedimage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StaticWebSiteGeneratedImage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(staticwebsitegeneratedimage.Name))
	return
}

func (staticwebsiteimage *StaticWebSiteImage) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staticwebsiteimage.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StaticWebSiteImage")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(staticwebsiteimage.Name))
	return
}

func (staticwebsiteparagraph *StaticWebSiteParagraph) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staticwebsiteparagraph.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "StaticWebSiteParagraph")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(staticwebsiteparagraph.Name))
	return
}

func (xhtml_content *XHTML_CONTENT) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xhtml_content.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "XHTML_CONTENT")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(xhtml_content.Name))
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

func (attribute_definition_boolean_rendering *ATTRIBUTE_DEFINITION_BOOLEAN_Rendering) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_boolean_rendering.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_date.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_date_rendering *ATTRIBUTE_DEFINITION_DATE_Rendering) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_date_rendering.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_enumeration.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_enumeration_rendering *ATTRIBUTE_DEFINITION_ENUMERATION_Rendering) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_enumeration_rendering.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_integer.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_integer_rendering *ATTRIBUTE_DEFINITION_INTEGER_Rendering) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_integer_rendering.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_real.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_real_rendering *ATTRIBUTE_DEFINITION_REAL_Rendering) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_real_rendering.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_rendering *ATTRIBUTE_DEFINITION_Rendering) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_rendering.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_string.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_string_rendering *ATTRIBUTE_DEFINITION_STRING_Rendering) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_string_rendering.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_xhtml.GongGetReferenceIdentifier(stage))
	return
}

func (attribute_definition_xhtml_rendering *ATTRIBUTE_DEFINITION_XHTML_Rendering) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attribute_definition_xhtml_rendering.GongGetReferenceIdentifier(stage))
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

func (embeddedjpgimage *EmbeddedJpgImage) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", embeddedjpgimage.GongGetReferenceIdentifier(stage))
	return
}

func (embeddedpngimage *EmbeddedPngImage) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", embeddedpngimage.GongGetReferenceIdentifier(stage))
	return
}

func (embeddedsvgimage *EmbeddedSvgImage) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", embeddedsvgimage.GongGetReferenceIdentifier(stage))
	return
}

func (kill *Kill) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", kill.GongGetReferenceIdentifier(stage))
	return
}

func (map_identifier_bool *Map_identifier_bool) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", map_identifier_bool.GongGetReferenceIdentifier(stage))
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

func (specification_rendering *SPECIFICATION_Rendering) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", specification_rendering.GongGetReferenceIdentifier(stage))
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

func (spec_object_type_rendering *SPEC_OBJECT_TYPE_Rendering) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", spec_object_type_rendering.GongGetReferenceIdentifier(stage))
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

func (staticwebsite *StaticWebSite) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staticwebsite.GongGetReferenceIdentifier(stage))
	return
}

func (staticwebsitechapter *StaticWebSiteChapter) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staticwebsitechapter.GongGetReferenceIdentifier(stage))
	return
}

func (staticwebsitegeneratedimage *StaticWebSiteGeneratedImage) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staticwebsitegeneratedimage.GongGetReferenceIdentifier(stage))
	return
}

func (staticwebsiteimage *StaticWebSiteImage) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staticwebsiteimage.GongGetReferenceIdentifier(stage))
	return
}

func (staticwebsiteparagraph *StaticWebSiteParagraph) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staticwebsiteparagraph.GongGetReferenceIdentifier(stage))
	return
}

func (xhtml_content *XHTML_CONTENT) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", xhtml_content.GongGetReferenceIdentifier(stage))
	return
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

// end of template
