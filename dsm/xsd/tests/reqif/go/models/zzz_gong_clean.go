// generated code - do not edit
package models

// GongCleanSlice removes unstaged elements from a slice of pointers of type T.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanSlice[T PointerToGongstruct](stage *Stage, slice *[]T) (modified bool) {
	if *slice == nil {
		return false
	}

	var cleanedSlice []T
	for _, element := range *slice {
		if IsStagedPointerToGongstruct(stage, element) {
			cleanedSlice = append(cleanedSlice, element)
		}
	}
	modified = len(cleanedSlice) != len(*slice)
	*slice = cleanedSlice
	return
}

// GongCleanPointer sets the pointer to nil if the referenced element is not staged.
// T must be a pointer to a struct that implements PointerToGongstruct.
func GongCleanPointer[T PointerToGongstruct](stage *Stage, element *T) (modified bool) {
	if !IsStagedPointerToGongstruct(stage, *element) {
		var zero T
		*element = zero
		return true
	}
	return false
}

// insertion point per named struct
// Clean garbage collect unstaged instances that are referenced by ALTERNATIVE_ID
func (alternative_id *ALTERNATIVE_ID) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_DEFINITION_BOOLEAN
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &attribute_definition_boolean.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &attribute_definition_boolean.DEFAULT_VALUE) || modified
	modified = GongCleanPointer(stage, &attribute_definition_boolean.TYPE) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_DEFINITION_DATE
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &attribute_definition_date.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &attribute_definition_date.DEFAULT_VALUE) || modified
	modified = GongCleanPointer(stage, &attribute_definition_date.TYPE) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_DEFINITION_ENUMERATION
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &attribute_definition_enumeration.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &attribute_definition_enumeration.DEFAULT_VALUE) || modified
	modified = GongCleanPointer(stage, &attribute_definition_enumeration.TYPE) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_DEFINITION_INTEGER
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &attribute_definition_integer.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &attribute_definition_integer.DEFAULT_VALUE) || modified
	modified = GongCleanPointer(stage, &attribute_definition_integer.TYPE) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_DEFINITION_REAL
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &attribute_definition_real.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &attribute_definition_real.DEFAULT_VALUE) || modified
	modified = GongCleanPointer(stage, &attribute_definition_real.TYPE) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_DEFINITION_STRING
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &attribute_definition_string.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &attribute_definition_string.DEFAULT_VALUE) || modified
	modified = GongCleanPointer(stage, &attribute_definition_string.TYPE) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_DEFINITION_XHTML
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &attribute_definition_xhtml.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &attribute_definition_xhtml.DEFAULT_VALUE) || modified
	modified = GongCleanPointer(stage, &attribute_definition_xhtml.TYPE) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_VALUE_BOOLEAN
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &attribute_value_boolean.DEFINITION) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_VALUE_DATE
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &attribute_value_date.DEFINITION) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_VALUE_ENUMERATION
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &attribute_value_enumeration.DEFINITION) || modified
	modified = GongCleanPointer(stage, &attribute_value_enumeration.VALUES) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_VALUE_INTEGER
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &attribute_value_integer.DEFINITION) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_VALUE_REAL
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &attribute_value_real.DEFINITION) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_VALUE_STRING
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &attribute_value_string.DEFINITION) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by ATTRIBUTE_VALUE_XHTML
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &attribute_value_xhtml.THE_VALUE) || modified
	modified = GongCleanPointer(stage, &attribute_value_xhtml.THE_ORIGINAL_VALUE) || modified
	modified = GongCleanPointer(stage, &attribute_value_xhtml.DEFINITION) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ALTERNATIVE_ID
func (a_alternative_id *A_ALTERNATIVE_ID) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &a_alternative_id.ALTERNATIVE_ID) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_DEFINITION_BOOLEAN_REF
func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_DEFINITION_DATE_REF
func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_DEFINITION_ENUMERATION_REF
func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_DEFINITION_INTEGER_REF
func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_DEFINITION_REAL_REF
func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_DEFINITION_STRING_REF
func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_DEFINITION_XHTML_REF
func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_VALUE_BOOLEAN
func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_VALUE_DATE
func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_attribute_value_date.ATTRIBUTE_VALUE_DATE) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_VALUE_ENUMERATION
func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_VALUE_INTEGER
func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_VALUE_REAL
func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_attribute_value_real.ATTRIBUTE_VALUE_REAL) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_VALUE_STRING
func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_attribute_value_string.ATTRIBUTE_VALUE_STRING) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_VALUE_XHTML
func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ATTRIBUTE_VALUE_XHTML_1
func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN) || modified
	modified = GongCleanSlice(stage, &a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE) || modified
	modified = GongCleanSlice(stage, &a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION) || modified
	modified = GongCleanSlice(stage, &a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER) || modified
	modified = GongCleanSlice(stage, &a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL) || modified
	modified = GongCleanSlice(stage, &a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING) || modified
	modified = GongCleanSlice(stage, &a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_CHILDREN
func (a_children *A_CHILDREN) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_children.SPEC_HIERARCHY) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_CORE_CONTENT
func (a_core_content *A_CORE_CONTENT) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &a_core_content.REQ_IF_CONTENT) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by A_DATATYPES
func (a_datatypes *A_DATATYPES) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_datatypes.DATATYPE_DEFINITION_BOOLEAN) || modified
	modified = GongCleanSlice(stage, &a_datatypes.DATATYPE_DEFINITION_DATE) || modified
	modified = GongCleanSlice(stage, &a_datatypes.DATATYPE_DEFINITION_ENUMERATION) || modified
	modified = GongCleanSlice(stage, &a_datatypes.DATATYPE_DEFINITION_INTEGER) || modified
	modified = GongCleanSlice(stage, &a_datatypes.DATATYPE_DEFINITION_REAL) || modified
	modified = GongCleanSlice(stage, &a_datatypes.DATATYPE_DEFINITION_STRING) || modified
	modified = GongCleanSlice(stage, &a_datatypes.DATATYPE_DEFINITION_XHTML) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_DATATYPE_DEFINITION_BOOLEAN_REF
func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_DATATYPE_DEFINITION_DATE_REF
func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_DATATYPE_DEFINITION_ENUMERATION_REF
func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_DATATYPE_DEFINITION_INTEGER_REF
func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_DATATYPE_DEFINITION_REAL_REF
func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_DATATYPE_DEFINITION_STRING_REF
func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_DATATYPE_DEFINITION_XHTML_REF
func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_EDITABLE_ATTS
func (a_editable_atts *A_EDITABLE_ATTS) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_ENUM_VALUE_REF
func (a_enum_value_ref *A_ENUM_VALUE_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_OBJECT
func (a_object *A_OBJECT) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_PROPERTIES
func (a_properties *A_PROPERTIES) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &a_properties.EMBEDDED_VALUE) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by A_RELATION_GROUP_TYPE_REF
func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_SOURCE_1
func (a_source_1 *A_SOURCE_1) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_SOURCE_SPECIFICATION_1
func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_SPECIFICATIONS
func (a_specifications *A_SPECIFICATIONS) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_specifications.SPECIFICATION) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_SPECIFICATION_TYPE_REF
func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_SPECIFIED_VALUES
func (a_specified_values *A_SPECIFIED_VALUES) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_specified_values.ENUM_VALUE) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_SPEC_ATTRIBUTES
func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN) || modified
	modified = GongCleanSlice(stage, &a_spec_attributes.ATTRIBUTE_DEFINITION_DATE) || modified
	modified = GongCleanSlice(stage, &a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION) || modified
	modified = GongCleanSlice(stage, &a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER) || modified
	modified = GongCleanSlice(stage, &a_spec_attributes.ATTRIBUTE_DEFINITION_REAL) || modified
	modified = GongCleanSlice(stage, &a_spec_attributes.ATTRIBUTE_DEFINITION_STRING) || modified
	modified = GongCleanSlice(stage, &a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_SPEC_OBJECTS
func (a_spec_objects *A_SPEC_OBJECTS) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_spec_objects.SPEC_OBJECT) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_SPEC_OBJECT_TYPE_REF
func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_SPEC_RELATIONS
func (a_spec_relations *A_SPEC_RELATIONS) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_spec_relations.SPEC_RELATION) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_SPEC_RELATION_GROUPS
func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_spec_relation_groups.RELATION_GROUP) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_SPEC_RELATION_REF
func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_SPEC_RELATION_TYPE_REF
func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_SPEC_TYPES
func (a_spec_types *A_SPEC_TYPES) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_spec_types.RELATION_GROUP_TYPE) || modified
	modified = GongCleanSlice(stage, &a_spec_types.SPEC_OBJECT_TYPE) || modified
	modified = GongCleanSlice(stage, &a_spec_types.SPEC_RELATION_TYPE) || modified
	modified = GongCleanSlice(stage, &a_spec_types.SPECIFICATION_TYPE) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by A_THE_HEADER
func (a_the_header *A_THE_HEADER) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &a_the_header.REQ_IF_HEADER) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by A_TOOL_EXTENSIONS
func (a_tool_extensions *A_TOOL_EXTENSIONS) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	modified = GongCleanSlice(stage, &a_tool_extensions.REQ_IF_TOOL_EXTENSION) || modified
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by DATATYPE_DEFINITION_BOOLEAN
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &datatype_definition_boolean.ALTERNATIVE_ID) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DATATYPE_DEFINITION_DATE
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &datatype_definition_date.ALTERNATIVE_ID) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DATATYPE_DEFINITION_ENUMERATION
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &datatype_definition_enumeration.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &datatype_definition_enumeration.SPECIFIED_VALUES) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DATATYPE_DEFINITION_INTEGER
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &datatype_definition_integer.ALTERNATIVE_ID) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DATATYPE_DEFINITION_REAL
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &datatype_definition_real.ALTERNATIVE_ID) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DATATYPE_DEFINITION_STRING
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &datatype_definition_string.ALTERNATIVE_ID) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by DATATYPE_DEFINITION_XHTML
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &datatype_definition_xhtml.ALTERNATIVE_ID) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by EMBEDDED_VALUE
func (embedded_value *EMBEDDED_VALUE) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by ENUM_VALUE
func (enum_value *ENUM_VALUE) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &enum_value.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &enum_value.PROPERTIES) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RELATION_GROUP
func (relation_group *RELATION_GROUP) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &relation_group.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &relation_group.SOURCE_SPECIFICATION) || modified
	modified = GongCleanPointer(stage, &relation_group.SPEC_RELATIONS) || modified
	modified = GongCleanPointer(stage, &relation_group.TARGET_SPECIFICATION) || modified
	modified = GongCleanPointer(stage, &relation_group.TYPE) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by RELATION_GROUP_TYPE
func (relation_group_type *RELATION_GROUP_TYPE) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &relation_group_type.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &relation_group_type.SPEC_ATTRIBUTES) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by REQ_IF
func (req_if *REQ_IF) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &req_if.THE_HEADER) || modified
	modified = GongCleanPointer(stage, &req_if.CORE_CONTENT) || modified
	modified = GongCleanPointer(stage, &req_if.TOOL_EXTENSIONS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by REQ_IF_CONTENT
func (req_if_content *REQ_IF_CONTENT) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &req_if_content.DATATYPES) || modified
	modified = GongCleanPointer(stage, &req_if_content.SPEC_TYPES) || modified
	modified = GongCleanPointer(stage, &req_if_content.SPEC_OBJECTS) || modified
	modified = GongCleanPointer(stage, &req_if_content.SPEC_RELATIONS) || modified
	modified = GongCleanPointer(stage, &req_if_content.SPECIFICATIONS) || modified
	modified = GongCleanPointer(stage, &req_if_content.SPEC_RELATION_GROUPS) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by REQ_IF_HEADER
func (req_if_header *REQ_IF_HEADER) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by REQ_IF_TOOL_EXTENSION
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by SPECIFICATION
func (specification *SPECIFICATION) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &specification.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &specification.CHILDREN) || modified
	modified = GongCleanPointer(stage, &specification.VALUES) || modified
	modified = GongCleanPointer(stage, &specification.TYPE) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SPECIFICATION_TYPE
func (specification_type *SPECIFICATION_TYPE) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &specification_type.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &specification_type.SPEC_ATTRIBUTES) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SPEC_HIERARCHY
func (spec_hierarchy *SPEC_HIERARCHY) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &spec_hierarchy.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &spec_hierarchy.CHILDREN) || modified
	modified = GongCleanPointer(stage, &spec_hierarchy.EDITABLE_ATTS) || modified
	modified = GongCleanPointer(stage, &spec_hierarchy.OBJECT) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SPEC_OBJECT
func (spec_object *SPEC_OBJECT) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &spec_object.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &spec_object.VALUES) || modified
	modified = GongCleanPointer(stage, &spec_object.TYPE) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SPEC_OBJECT_TYPE
func (spec_object_type *SPEC_OBJECT_TYPE) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &spec_object_type.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &spec_object_type.SPEC_ATTRIBUTES) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SPEC_RELATION
func (spec_relation *SPEC_RELATION) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &spec_relation.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &spec_relation.VALUES) || modified
	modified = GongCleanPointer(stage, &spec_relation.SOURCE) || modified
	modified = GongCleanPointer(stage, &spec_relation.TARGET) || modified
	modified = GongCleanPointer(stage, &spec_relation.TYPE) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by SPEC_RELATION_TYPE
func (spec_relation_type *SPEC_RELATION_TYPE) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	modified = GongCleanPointer(stage, &spec_relation_type.ALTERNATIVE_ID) || modified
	modified = GongCleanPointer(stage, &spec_relation_type.SPEC_ATTRIBUTES) || modified
	return
}

// Clean garbage collect unstaged instances that are referenced by XHTML_CONTENT
func (xhtml_content *XHTML_CONTENT) GongClean(stage *Stage) (modified bool) {
	// insertion point per field
	// insertion point per field
	return
}

// Clean garbage collect unstaged instances that are referenced by staged elements
func (stage *Stage) Clean() (modified bool) {
	for _, instance := range stage.GetInstances() {
		modified = instance.GongClean(stage) || modified
	}
	return
}
