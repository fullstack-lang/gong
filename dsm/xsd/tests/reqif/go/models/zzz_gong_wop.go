// generated code - do not edit
package models

import "time"

// to avoid compile error if no time field is present
var __GONG_time_The_fool_doth_think_he_is_wise__ = time.Hour

var _ = __GONG_time_The_fool_doth_think_he_is_wise__

// insertion point
type ALTERNATIVE_ID_WOP struct {
	// insertion point

	Name string

	IDENTIFIER string
}

func (from *ALTERNATIVE_ID) CopyBasicFields(to *ALTERNATIVE_ID) {
	// insertion point
	to.Name = from.Name
	to.IDENTIFIER = from.IDENTIFIER
}

type ATTRIBUTE_DEFINITION_BOOLEAN_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	IS_EDITABLE bool

	LAST_CHANGE string

	LONG_NAME string
}

func (from *ATTRIBUTE_DEFINITION_BOOLEAN) CopyBasicFields(to *ATTRIBUTE_DEFINITION_BOOLEAN) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.IS_EDITABLE = from.IS_EDITABLE
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type ATTRIBUTE_DEFINITION_DATE_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	IS_EDITABLE bool

	LAST_CHANGE string

	LONG_NAME string
}

func (from *ATTRIBUTE_DEFINITION_DATE) CopyBasicFields(to *ATTRIBUTE_DEFINITION_DATE) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.IS_EDITABLE = from.IS_EDITABLE
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type ATTRIBUTE_DEFINITION_ENUMERATION_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	IS_EDITABLE bool

	LAST_CHANGE string

	LONG_NAME string

	MULTI_VALUED bool
}

func (from *ATTRIBUTE_DEFINITION_ENUMERATION) CopyBasicFields(to *ATTRIBUTE_DEFINITION_ENUMERATION) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.IS_EDITABLE = from.IS_EDITABLE
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
	to.MULTI_VALUED = from.MULTI_VALUED
}

type ATTRIBUTE_DEFINITION_INTEGER_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	IS_EDITABLE bool

	LAST_CHANGE string

	LONG_NAME string
}

func (from *ATTRIBUTE_DEFINITION_INTEGER) CopyBasicFields(to *ATTRIBUTE_DEFINITION_INTEGER) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.IS_EDITABLE = from.IS_EDITABLE
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type ATTRIBUTE_DEFINITION_REAL_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	IS_EDITABLE bool

	LAST_CHANGE string

	LONG_NAME string
}

func (from *ATTRIBUTE_DEFINITION_REAL) CopyBasicFields(to *ATTRIBUTE_DEFINITION_REAL) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.IS_EDITABLE = from.IS_EDITABLE
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type ATTRIBUTE_DEFINITION_STRING_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	IS_EDITABLE bool

	LAST_CHANGE string

	LONG_NAME string
}

func (from *ATTRIBUTE_DEFINITION_STRING) CopyBasicFields(to *ATTRIBUTE_DEFINITION_STRING) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.IS_EDITABLE = from.IS_EDITABLE
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type ATTRIBUTE_DEFINITION_XHTML_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	IS_EDITABLE bool

	LAST_CHANGE string

	LONG_NAME string
}

func (from *ATTRIBUTE_DEFINITION_XHTML) CopyBasicFields(to *ATTRIBUTE_DEFINITION_XHTML) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.IS_EDITABLE = from.IS_EDITABLE
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type ATTRIBUTE_VALUE_BOOLEAN_WOP struct {
	// insertion point

	Name string

	THE_VALUE bool
}

func (from *ATTRIBUTE_VALUE_BOOLEAN) CopyBasicFields(to *ATTRIBUTE_VALUE_BOOLEAN) {
	// insertion point
	to.Name = from.Name
	to.THE_VALUE = from.THE_VALUE
}

type ATTRIBUTE_VALUE_DATE_WOP struct {
	// insertion point

	Name string

	THE_VALUE string
}

func (from *ATTRIBUTE_VALUE_DATE) CopyBasicFields(to *ATTRIBUTE_VALUE_DATE) {
	// insertion point
	to.Name = from.Name
	to.THE_VALUE = from.THE_VALUE
}

type ATTRIBUTE_VALUE_ENUMERATION_WOP struct {
	// insertion point

	Name string
}

func (from *ATTRIBUTE_VALUE_ENUMERATION) CopyBasicFields(to *ATTRIBUTE_VALUE_ENUMERATION) {
	// insertion point
	to.Name = from.Name
}

type ATTRIBUTE_VALUE_INTEGER_WOP struct {
	// insertion point

	Name string

	THE_VALUE int
}

func (from *ATTRIBUTE_VALUE_INTEGER) CopyBasicFields(to *ATTRIBUTE_VALUE_INTEGER) {
	// insertion point
	to.Name = from.Name
	to.THE_VALUE = from.THE_VALUE
}

type ATTRIBUTE_VALUE_REAL_WOP struct {
	// insertion point

	Name string

	THE_VALUE float64
}

func (from *ATTRIBUTE_VALUE_REAL) CopyBasicFields(to *ATTRIBUTE_VALUE_REAL) {
	// insertion point
	to.Name = from.Name
	to.THE_VALUE = from.THE_VALUE
}

type ATTRIBUTE_VALUE_STRING_WOP struct {
	// insertion point

	Name string

	THE_VALUE string
}

func (from *ATTRIBUTE_VALUE_STRING) CopyBasicFields(to *ATTRIBUTE_VALUE_STRING) {
	// insertion point
	to.Name = from.Name
	to.THE_VALUE = from.THE_VALUE
}

type ATTRIBUTE_VALUE_XHTML_WOP struct {
	// insertion point

	Name string

	IS_SIMPLIFIED bool
}

func (from *ATTRIBUTE_VALUE_XHTML) CopyBasicFields(to *ATTRIBUTE_VALUE_XHTML) {
	// insertion point
	to.Name = from.Name
	to.IS_SIMPLIFIED = from.IS_SIMPLIFIED
}

type A_ALTERNATIVE_ID_WOP struct {
	// insertion point

	Name string
}

func (from *A_ALTERNATIVE_ID) CopyBasicFields(to *A_ALTERNATIVE_ID) {
	// insertion point
	to.Name = from.Name
}

type A_ATTRIBUTE_DEFINITION_BOOLEAN_REF_WOP struct {
	// insertion point

	Name string

	ATTRIBUTE_DEFINITION_BOOLEAN_REF string
}

func (from *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) CopyBasicFields(to *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) {
	// insertion point
	to.Name = from.Name
	to.ATTRIBUTE_DEFINITION_BOOLEAN_REF = from.ATTRIBUTE_DEFINITION_BOOLEAN_REF
}

type A_ATTRIBUTE_DEFINITION_DATE_REF_WOP struct {
	// insertion point

	Name string

	ATTRIBUTE_DEFINITION_DATE_REF string
}

func (from *A_ATTRIBUTE_DEFINITION_DATE_REF) CopyBasicFields(to *A_ATTRIBUTE_DEFINITION_DATE_REF) {
	// insertion point
	to.Name = from.Name
	to.ATTRIBUTE_DEFINITION_DATE_REF = from.ATTRIBUTE_DEFINITION_DATE_REF
}

type A_ATTRIBUTE_DEFINITION_ENUMERATION_REF_WOP struct {
	// insertion point

	Name string

	ATTRIBUTE_DEFINITION_ENUMERATION_REF string
}

func (from *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) CopyBasicFields(to *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) {
	// insertion point
	to.Name = from.Name
	to.ATTRIBUTE_DEFINITION_ENUMERATION_REF = from.ATTRIBUTE_DEFINITION_ENUMERATION_REF
}

type A_ATTRIBUTE_DEFINITION_INTEGER_REF_WOP struct {
	// insertion point

	Name string

	ATTRIBUTE_DEFINITION_INTEGER_REF string
}

func (from *A_ATTRIBUTE_DEFINITION_INTEGER_REF) CopyBasicFields(to *A_ATTRIBUTE_DEFINITION_INTEGER_REF) {
	// insertion point
	to.Name = from.Name
	to.ATTRIBUTE_DEFINITION_INTEGER_REF = from.ATTRIBUTE_DEFINITION_INTEGER_REF
}

type A_ATTRIBUTE_DEFINITION_REAL_REF_WOP struct {
	// insertion point

	Name string

	ATTRIBUTE_DEFINITION_REAL_REF string
}

func (from *A_ATTRIBUTE_DEFINITION_REAL_REF) CopyBasicFields(to *A_ATTRIBUTE_DEFINITION_REAL_REF) {
	// insertion point
	to.Name = from.Name
	to.ATTRIBUTE_DEFINITION_REAL_REF = from.ATTRIBUTE_DEFINITION_REAL_REF
}

type A_ATTRIBUTE_DEFINITION_STRING_REF_WOP struct {
	// insertion point

	Name string

	ATTRIBUTE_DEFINITION_STRING_REF string
}

func (from *A_ATTRIBUTE_DEFINITION_STRING_REF) CopyBasicFields(to *A_ATTRIBUTE_DEFINITION_STRING_REF) {
	// insertion point
	to.Name = from.Name
	to.ATTRIBUTE_DEFINITION_STRING_REF = from.ATTRIBUTE_DEFINITION_STRING_REF
}

type A_ATTRIBUTE_DEFINITION_XHTML_REF_WOP struct {
	// insertion point

	Name string

	ATTRIBUTE_DEFINITION_XHTML_REF string
}

func (from *A_ATTRIBUTE_DEFINITION_XHTML_REF) CopyBasicFields(to *A_ATTRIBUTE_DEFINITION_XHTML_REF) {
	// insertion point
	to.Name = from.Name
	to.ATTRIBUTE_DEFINITION_XHTML_REF = from.ATTRIBUTE_DEFINITION_XHTML_REF
}

type A_ATTRIBUTE_VALUE_BOOLEAN_WOP struct {
	// insertion point

	Name string
}

func (from *A_ATTRIBUTE_VALUE_BOOLEAN) CopyBasicFields(to *A_ATTRIBUTE_VALUE_BOOLEAN) {
	// insertion point
	to.Name = from.Name
}

type A_ATTRIBUTE_VALUE_DATE_WOP struct {
	// insertion point

	Name string
}

func (from *A_ATTRIBUTE_VALUE_DATE) CopyBasicFields(to *A_ATTRIBUTE_VALUE_DATE) {
	// insertion point
	to.Name = from.Name
}

type A_ATTRIBUTE_VALUE_ENUMERATION_WOP struct {
	// insertion point

	Name string
}

func (from *A_ATTRIBUTE_VALUE_ENUMERATION) CopyBasicFields(to *A_ATTRIBUTE_VALUE_ENUMERATION) {
	// insertion point
	to.Name = from.Name
}

type A_ATTRIBUTE_VALUE_INTEGER_WOP struct {
	// insertion point

	Name string
}

func (from *A_ATTRIBUTE_VALUE_INTEGER) CopyBasicFields(to *A_ATTRIBUTE_VALUE_INTEGER) {
	// insertion point
	to.Name = from.Name
}

type A_ATTRIBUTE_VALUE_REAL_WOP struct {
	// insertion point

	Name string
}

func (from *A_ATTRIBUTE_VALUE_REAL) CopyBasicFields(to *A_ATTRIBUTE_VALUE_REAL) {
	// insertion point
	to.Name = from.Name
}

type A_ATTRIBUTE_VALUE_STRING_WOP struct {
	// insertion point

	Name string
}

func (from *A_ATTRIBUTE_VALUE_STRING) CopyBasicFields(to *A_ATTRIBUTE_VALUE_STRING) {
	// insertion point
	to.Name = from.Name
}

type A_ATTRIBUTE_VALUE_XHTML_WOP struct {
	// insertion point

	Name string
}

func (from *A_ATTRIBUTE_VALUE_XHTML) CopyBasicFields(to *A_ATTRIBUTE_VALUE_XHTML) {
	// insertion point
	to.Name = from.Name
}

type A_ATTRIBUTE_VALUE_XHTML_1_WOP struct {
	// insertion point

	Name string
}

func (from *A_ATTRIBUTE_VALUE_XHTML_1) CopyBasicFields(to *A_ATTRIBUTE_VALUE_XHTML_1) {
	// insertion point
	to.Name = from.Name
}

type A_CHILDREN_WOP struct {
	// insertion point

	Name string
}

func (from *A_CHILDREN) CopyBasicFields(to *A_CHILDREN) {
	// insertion point
	to.Name = from.Name
}

type A_CORE_CONTENT_WOP struct {
	// insertion point

	Name string
}

func (from *A_CORE_CONTENT) CopyBasicFields(to *A_CORE_CONTENT) {
	// insertion point
	to.Name = from.Name
}

type A_DATATYPES_WOP struct {
	// insertion point

	Name string
}

func (from *A_DATATYPES) CopyBasicFields(to *A_DATATYPES) {
	// insertion point
	to.Name = from.Name
}

type A_DATATYPE_DEFINITION_BOOLEAN_REF_WOP struct {
	// insertion point

	Name string

	DATATYPE_DEFINITION_BOOLEAN_REF string
}

func (from *A_DATATYPE_DEFINITION_BOOLEAN_REF) CopyBasicFields(to *A_DATATYPE_DEFINITION_BOOLEAN_REF) {
	// insertion point
	to.Name = from.Name
	to.DATATYPE_DEFINITION_BOOLEAN_REF = from.DATATYPE_DEFINITION_BOOLEAN_REF
}

type A_DATATYPE_DEFINITION_DATE_REF_WOP struct {
	// insertion point

	Name string

	DATATYPE_DEFINITION_DATE_REF string
}

func (from *A_DATATYPE_DEFINITION_DATE_REF) CopyBasicFields(to *A_DATATYPE_DEFINITION_DATE_REF) {
	// insertion point
	to.Name = from.Name
	to.DATATYPE_DEFINITION_DATE_REF = from.DATATYPE_DEFINITION_DATE_REF
}

type A_DATATYPE_DEFINITION_ENUMERATION_REF_WOP struct {
	// insertion point

	Name string

	DATATYPE_DEFINITION_ENUMERATION_REF string
}

func (from *A_DATATYPE_DEFINITION_ENUMERATION_REF) CopyBasicFields(to *A_DATATYPE_DEFINITION_ENUMERATION_REF) {
	// insertion point
	to.Name = from.Name
	to.DATATYPE_DEFINITION_ENUMERATION_REF = from.DATATYPE_DEFINITION_ENUMERATION_REF
}

type A_DATATYPE_DEFINITION_INTEGER_REF_WOP struct {
	// insertion point

	Name string

	DATATYPE_DEFINITION_INTEGER_REF string
}

func (from *A_DATATYPE_DEFINITION_INTEGER_REF) CopyBasicFields(to *A_DATATYPE_DEFINITION_INTEGER_REF) {
	// insertion point
	to.Name = from.Name
	to.DATATYPE_DEFINITION_INTEGER_REF = from.DATATYPE_DEFINITION_INTEGER_REF
}

type A_DATATYPE_DEFINITION_REAL_REF_WOP struct {
	// insertion point

	Name string

	DATATYPE_DEFINITION_REAL_REF string
}

func (from *A_DATATYPE_DEFINITION_REAL_REF) CopyBasicFields(to *A_DATATYPE_DEFINITION_REAL_REF) {
	// insertion point
	to.Name = from.Name
	to.DATATYPE_DEFINITION_REAL_REF = from.DATATYPE_DEFINITION_REAL_REF
}

type A_DATATYPE_DEFINITION_STRING_REF_WOP struct {
	// insertion point

	Name string

	DATATYPE_DEFINITION_STRING_REF string
}

func (from *A_DATATYPE_DEFINITION_STRING_REF) CopyBasicFields(to *A_DATATYPE_DEFINITION_STRING_REF) {
	// insertion point
	to.Name = from.Name
	to.DATATYPE_DEFINITION_STRING_REF = from.DATATYPE_DEFINITION_STRING_REF
}

type A_DATATYPE_DEFINITION_XHTML_REF_WOP struct {
	// insertion point

	Name string

	DATATYPE_DEFINITION_XHTML_REF string
}

func (from *A_DATATYPE_DEFINITION_XHTML_REF) CopyBasicFields(to *A_DATATYPE_DEFINITION_XHTML_REF) {
	// insertion point
	to.Name = from.Name
	to.DATATYPE_DEFINITION_XHTML_REF = from.DATATYPE_DEFINITION_XHTML_REF
}

type A_EDITABLE_ATTS_WOP struct {
	// insertion point

	Name string

	ATTRIBUTE_DEFINITION_BOOLEAN_REF string

	ATTRIBUTE_DEFINITION_DATE_REF string

	ATTRIBUTE_DEFINITION_ENUMERATION_REF string

	ATTRIBUTE_DEFINITION_INTEGER_REF string

	ATTRIBUTE_DEFINITION_REAL_REF string

	ATTRIBUTE_DEFINITION_STRING_REF string

	ATTRIBUTE_DEFINITION_XHTML_REF string
}

func (from *A_EDITABLE_ATTS) CopyBasicFields(to *A_EDITABLE_ATTS) {
	// insertion point
	to.Name = from.Name
	to.ATTRIBUTE_DEFINITION_BOOLEAN_REF = from.ATTRIBUTE_DEFINITION_BOOLEAN_REF
	to.ATTRIBUTE_DEFINITION_DATE_REF = from.ATTRIBUTE_DEFINITION_DATE_REF
	to.ATTRIBUTE_DEFINITION_ENUMERATION_REF = from.ATTRIBUTE_DEFINITION_ENUMERATION_REF
	to.ATTRIBUTE_DEFINITION_INTEGER_REF = from.ATTRIBUTE_DEFINITION_INTEGER_REF
	to.ATTRIBUTE_DEFINITION_REAL_REF = from.ATTRIBUTE_DEFINITION_REAL_REF
	to.ATTRIBUTE_DEFINITION_STRING_REF = from.ATTRIBUTE_DEFINITION_STRING_REF
	to.ATTRIBUTE_DEFINITION_XHTML_REF = from.ATTRIBUTE_DEFINITION_XHTML_REF
}

type A_ENUM_VALUE_REF_WOP struct {
	// insertion point

	Name string

	ENUM_VALUE_REF string
}

func (from *A_ENUM_VALUE_REF) CopyBasicFields(to *A_ENUM_VALUE_REF) {
	// insertion point
	to.Name = from.Name
	to.ENUM_VALUE_REF = from.ENUM_VALUE_REF
}

type A_OBJECT_WOP struct {
	// insertion point

	Name string

	SPEC_OBJECT_REF string
}

func (from *A_OBJECT) CopyBasicFields(to *A_OBJECT) {
	// insertion point
	to.Name = from.Name
	to.SPEC_OBJECT_REF = from.SPEC_OBJECT_REF
}

type A_PROPERTIES_WOP struct {
	// insertion point

	Name string
}

func (from *A_PROPERTIES) CopyBasicFields(to *A_PROPERTIES) {
	// insertion point
	to.Name = from.Name
}

type A_RELATION_GROUP_TYPE_REF_WOP struct {
	// insertion point

	Name string

	RELATION_GROUP_TYPE_REF string
}

func (from *A_RELATION_GROUP_TYPE_REF) CopyBasicFields(to *A_RELATION_GROUP_TYPE_REF) {
	// insertion point
	to.Name = from.Name
	to.RELATION_GROUP_TYPE_REF = from.RELATION_GROUP_TYPE_REF
}

type A_SOURCE_1_WOP struct {
	// insertion point

	Name string

	SPEC_OBJECT_REF string
}

func (from *A_SOURCE_1) CopyBasicFields(to *A_SOURCE_1) {
	// insertion point
	to.Name = from.Name
	to.SPEC_OBJECT_REF = from.SPEC_OBJECT_REF
}

type A_SOURCE_SPECIFICATION_1_WOP struct {
	// insertion point

	Name string

	SPECIFICATION_REF Enum_GLOBAL_REF
}

func (from *A_SOURCE_SPECIFICATION_1) CopyBasicFields(to *A_SOURCE_SPECIFICATION_1) {
	// insertion point
	to.Name = from.Name
	to.SPECIFICATION_REF = from.SPECIFICATION_REF
}

type A_SPECIFICATIONS_WOP struct {
	// insertion point

	Name string
}

func (from *A_SPECIFICATIONS) CopyBasicFields(to *A_SPECIFICATIONS) {
	// insertion point
	to.Name = from.Name
}

type A_SPECIFICATION_TYPE_REF_WOP struct {
	// insertion point

	Name string

	SPECIFICATION_TYPE_REF string
}

func (from *A_SPECIFICATION_TYPE_REF) CopyBasicFields(to *A_SPECIFICATION_TYPE_REF) {
	// insertion point
	to.Name = from.Name
	to.SPECIFICATION_TYPE_REF = from.SPECIFICATION_TYPE_REF
}

type A_SPECIFIED_VALUES_WOP struct {
	// insertion point

	Name string
}

func (from *A_SPECIFIED_VALUES) CopyBasicFields(to *A_SPECIFIED_VALUES) {
	// insertion point
	to.Name = from.Name
}

type A_SPEC_ATTRIBUTES_WOP struct {
	// insertion point

	Name string
}

func (from *A_SPEC_ATTRIBUTES) CopyBasicFields(to *A_SPEC_ATTRIBUTES) {
	// insertion point
	to.Name = from.Name
}

type A_SPEC_OBJECTS_WOP struct {
	// insertion point

	Name string
}

func (from *A_SPEC_OBJECTS) CopyBasicFields(to *A_SPEC_OBJECTS) {
	// insertion point
	to.Name = from.Name
}

type A_SPEC_OBJECT_TYPE_REF_WOP struct {
	// insertion point

	Name string

	SPEC_OBJECT_TYPE_REF string
}

func (from *A_SPEC_OBJECT_TYPE_REF) CopyBasicFields(to *A_SPEC_OBJECT_TYPE_REF) {
	// insertion point
	to.Name = from.Name
	to.SPEC_OBJECT_TYPE_REF = from.SPEC_OBJECT_TYPE_REF
}

type A_SPEC_RELATIONS_WOP struct {
	// insertion point

	Name string
}

func (from *A_SPEC_RELATIONS) CopyBasicFields(to *A_SPEC_RELATIONS) {
	// insertion point
	to.Name = from.Name
}

type A_SPEC_RELATION_GROUPS_WOP struct {
	// insertion point

	Name string
}

func (from *A_SPEC_RELATION_GROUPS) CopyBasicFields(to *A_SPEC_RELATION_GROUPS) {
	// insertion point
	to.Name = from.Name
}

type A_SPEC_RELATION_REF_WOP struct {
	// insertion point

	Name string

	SPEC_RELATION_REF string
}

func (from *A_SPEC_RELATION_REF) CopyBasicFields(to *A_SPEC_RELATION_REF) {
	// insertion point
	to.Name = from.Name
	to.SPEC_RELATION_REF = from.SPEC_RELATION_REF
}

type A_SPEC_RELATION_TYPE_REF_WOP struct {
	// insertion point

	Name string

	SPEC_RELATION_TYPE_REF string
}

func (from *A_SPEC_RELATION_TYPE_REF) CopyBasicFields(to *A_SPEC_RELATION_TYPE_REF) {
	// insertion point
	to.Name = from.Name
	to.SPEC_RELATION_TYPE_REF = from.SPEC_RELATION_TYPE_REF
}

type A_SPEC_TYPES_WOP struct {
	// insertion point

	Name string
}

func (from *A_SPEC_TYPES) CopyBasicFields(to *A_SPEC_TYPES) {
	// insertion point
	to.Name = from.Name
}

type A_THE_HEADER_WOP struct {
	// insertion point

	Name string
}

func (from *A_THE_HEADER) CopyBasicFields(to *A_THE_HEADER) {
	// insertion point
	to.Name = from.Name
}

type A_TOOL_EXTENSIONS_WOP struct {
	// insertion point

	Name string
}

func (from *A_TOOL_EXTENSIONS) CopyBasicFields(to *A_TOOL_EXTENSIONS) {
	// insertion point
	to.Name = from.Name
}

type DATATYPE_DEFINITION_BOOLEAN_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	LAST_CHANGE string

	LONG_NAME string
}

func (from *DATATYPE_DEFINITION_BOOLEAN) CopyBasicFields(to *DATATYPE_DEFINITION_BOOLEAN) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type DATATYPE_DEFINITION_DATE_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	LAST_CHANGE string

	LONG_NAME string
}

func (from *DATATYPE_DEFINITION_DATE) CopyBasicFields(to *DATATYPE_DEFINITION_DATE) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type DATATYPE_DEFINITION_ENUMERATION_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	LAST_CHANGE string

	LONG_NAME string
}

func (from *DATATYPE_DEFINITION_ENUMERATION) CopyBasicFields(to *DATATYPE_DEFINITION_ENUMERATION) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type DATATYPE_DEFINITION_INTEGER_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	LAST_CHANGE string

	LONG_NAME string

	MAX int

	MIN int
}

func (from *DATATYPE_DEFINITION_INTEGER) CopyBasicFields(to *DATATYPE_DEFINITION_INTEGER) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
	to.MAX = from.MAX
	to.MIN = from.MIN
}

type DATATYPE_DEFINITION_REAL_WOP struct {
	// insertion point

	Name string

	ACCURACY int

	DESC string

	IDENTIFIER string

	LAST_CHANGE string

	LONG_NAME string

	MAX float64

	MIN float64
}

func (from *DATATYPE_DEFINITION_REAL) CopyBasicFields(to *DATATYPE_DEFINITION_REAL) {
	// insertion point
	to.Name = from.Name
	to.ACCURACY = from.ACCURACY
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
	to.MAX = from.MAX
	to.MIN = from.MIN
}

type DATATYPE_DEFINITION_STRING_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	LAST_CHANGE string

	LONG_NAME string

	MAX_LENGTH int
}

func (from *DATATYPE_DEFINITION_STRING) CopyBasicFields(to *DATATYPE_DEFINITION_STRING) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
	to.MAX_LENGTH = from.MAX_LENGTH
}

type DATATYPE_DEFINITION_XHTML_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	LAST_CHANGE string

	LONG_NAME string
}

func (from *DATATYPE_DEFINITION_XHTML) CopyBasicFields(to *DATATYPE_DEFINITION_XHTML) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type EMBEDDED_VALUE_WOP struct {
	// insertion point

	Name string

	KEY int

	OTHER_CONTENT string
}

func (from *EMBEDDED_VALUE) CopyBasicFields(to *EMBEDDED_VALUE) {
	// insertion point
	to.Name = from.Name
	to.KEY = from.KEY
	to.OTHER_CONTENT = from.OTHER_CONTENT
}

type ENUM_VALUE_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	LAST_CHANGE string

	LONG_NAME string
}

func (from *ENUM_VALUE) CopyBasicFields(to *ENUM_VALUE) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type RELATION_GROUP_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	LAST_CHANGE string

	LONG_NAME string
}

func (from *RELATION_GROUP) CopyBasicFields(to *RELATION_GROUP) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type RELATION_GROUP_TYPE_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	LAST_CHANGE string

	LONG_NAME string
}

func (from *RELATION_GROUP_TYPE) CopyBasicFields(to *RELATION_GROUP_TYPE) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type REQ_IF_WOP struct {
	// insertion point

	Name string

	Lang string
}

func (from *REQ_IF) CopyBasicFields(to *REQ_IF) {
	// insertion point
	to.Name = from.Name
	to.Lang = from.Lang
}

type REQ_IF_CONTENT_WOP struct {
	// insertion point

	Name string
}

func (from *REQ_IF_CONTENT) CopyBasicFields(to *REQ_IF_CONTENT) {
	// insertion point
	to.Name = from.Name
}

type REQ_IF_HEADER_WOP struct {
	// insertion point

	Name string

	IDENTIFIER string

	COMMENT string

	CREATION_TIME string

	REPOSITORY_ID string

	REQ_IF_TOOL_ID string

	REQ_IF_VERSION string

	SOURCE_TOOL_ID string

	TITLE string
}

func (from *REQ_IF_HEADER) CopyBasicFields(to *REQ_IF_HEADER) {
	// insertion point
	to.Name = from.Name
	to.IDENTIFIER = from.IDENTIFIER
	to.COMMENT = from.COMMENT
	to.CREATION_TIME = from.CREATION_TIME
	to.REPOSITORY_ID = from.REPOSITORY_ID
	to.REQ_IF_TOOL_ID = from.REQ_IF_TOOL_ID
	to.REQ_IF_VERSION = from.REQ_IF_VERSION
	to.SOURCE_TOOL_ID = from.SOURCE_TOOL_ID
	to.TITLE = from.TITLE
}

type REQ_IF_TOOL_EXTENSION_WOP struct {
	// insertion point

	Name string
}

func (from *REQ_IF_TOOL_EXTENSION) CopyBasicFields(to *REQ_IF_TOOL_EXTENSION) {
	// insertion point
	to.Name = from.Name
}

type SPECIFICATION_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	LAST_CHANGE string

	LONG_NAME string
}

func (from *SPECIFICATION) CopyBasicFields(to *SPECIFICATION) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type SPECIFICATION_TYPE_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	LAST_CHANGE string

	LONG_NAME string
}

func (from *SPECIFICATION_TYPE) CopyBasicFields(to *SPECIFICATION_TYPE) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type SPEC_HIERARCHY_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	IS_EDITABLE bool

	IS_TABLE_INTERNAL bool

	LAST_CHANGE string

	LONG_NAME string
}

func (from *SPEC_HIERARCHY) CopyBasicFields(to *SPEC_HIERARCHY) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.IS_EDITABLE = from.IS_EDITABLE
	to.IS_TABLE_INTERNAL = from.IS_TABLE_INTERNAL
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type SPEC_OBJECT_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	LAST_CHANGE string

	LONG_NAME string
}

func (from *SPEC_OBJECT) CopyBasicFields(to *SPEC_OBJECT) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type SPEC_OBJECT_TYPE_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	LAST_CHANGE string

	LONG_NAME string
}

func (from *SPEC_OBJECT_TYPE) CopyBasicFields(to *SPEC_OBJECT_TYPE) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type SPEC_RELATION_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	LAST_CHANGE string

	LONG_NAME string
}

func (from *SPEC_RELATION) CopyBasicFields(to *SPEC_RELATION) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type SPEC_RELATION_TYPE_WOP struct {
	// insertion point

	Name string

	DESC string

	IDENTIFIER string

	LAST_CHANGE string

	LONG_NAME string
}

func (from *SPEC_RELATION_TYPE) CopyBasicFields(to *SPEC_RELATION_TYPE) {
	// insertion point
	to.Name = from.Name
	to.DESC = from.DESC
	to.IDENTIFIER = from.IDENTIFIER
	to.LAST_CHANGE = from.LAST_CHANGE
	to.LONG_NAME = from.LONG_NAME
}

type XHTML_CONTENT_WOP struct {
	// insertion point

	Name string

	EnclosedText string

	PureText string
}

func (from *XHTML_CONTENT) CopyBasicFields(to *XHTML_CONTENT) {
	// insertion point
	to.Name = from.Name
	to.EnclosedText = from.EnclosedText
	to.PureText = from.PureText
}

