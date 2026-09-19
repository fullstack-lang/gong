// generated code - do not edit
package models

import "fmt"

// IsStaged is the Stage method checking if a gongstruct instance is staged.
func (stage *Stage) IsStaged(instance GongstructIF) (ok bool) {
	if instance != nil {
		return instance.GongIsStaged(stage)
	}
	return false
}

// insertion point for stage per struct
func (alternative_id *ALTERNATIVE_ID) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ALTERNATIVE_IDs[alternative_id]

	return
}

func (stage *Stage) IsStagedALTERNATIVE_ID(alternative_id *ALTERNATIVE_ID) (ok bool) {

	return alternative_id.GongIsStaged(stage)
}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_BOOLEANs[attribute_definition_boolean]

	return
}

func (stage *Stage) IsStagedATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) (ok bool) {

	return attribute_definition_boolean.GongIsStaged(stage)
}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_DATEs[attribute_definition_date]

	return
}

func (stage *Stage) IsStagedATTRIBUTE_DEFINITION_DATE(attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) (ok bool) {

	return attribute_definition_date.GongIsStaged(stage)
}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_ENUMERATIONs[attribute_definition_enumeration]

	return
}

func (stage *Stage) IsStagedATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) (ok bool) {

	return attribute_definition_enumeration.GongIsStaged(stage)
}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_INTEGERs[attribute_definition_integer]

	return
}

func (stage *Stage) IsStagedATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) (ok bool) {

	return attribute_definition_integer.GongIsStaged(stage)
}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_REALs[attribute_definition_real]

	return
}

func (stage *Stage) IsStagedATTRIBUTE_DEFINITION_REAL(attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) (ok bool) {

	return attribute_definition_real.GongIsStaged(stage)
}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_STRINGs[attribute_definition_string]

	return
}

func (stage *Stage) IsStagedATTRIBUTE_DEFINITION_STRING(attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) (ok bool) {

	return attribute_definition_string.GongIsStaged(stage)
}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ATTRIBUTE_DEFINITION_XHTMLs[attribute_definition_xhtml]

	return
}

func (stage *Stage) IsStagedATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) (ok bool) {

	return attribute_definition_xhtml.GongIsStaged(stage)
}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_BOOLEANs[attribute_value_boolean]

	return
}

func (stage *Stage) IsStagedATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) (ok bool) {

	return attribute_value_boolean.GongIsStaged(stage)
}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_DATEs[attribute_value_date]

	return
}

func (stage *Stage) IsStagedATTRIBUTE_VALUE_DATE(attribute_value_date *ATTRIBUTE_VALUE_DATE) (ok bool) {

	return attribute_value_date.GongIsStaged(stage)
}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_ENUMERATIONs[attribute_value_enumeration]

	return
}

func (stage *Stage) IsStagedATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) (ok bool) {

	return attribute_value_enumeration.GongIsStaged(stage)
}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_INTEGERs[attribute_value_integer]

	return
}

func (stage *Stage) IsStagedATTRIBUTE_VALUE_INTEGER(attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) (ok bool) {

	return attribute_value_integer.GongIsStaged(stage)
}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_REALs[attribute_value_real]

	return
}

func (stage *Stage) IsStagedATTRIBUTE_VALUE_REAL(attribute_value_real *ATTRIBUTE_VALUE_REAL) (ok bool) {

	return attribute_value_real.GongIsStaged(stage)
}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_STRINGs[attribute_value_string]

	return
}

func (stage *Stage) IsStagedATTRIBUTE_VALUE_STRING(attribute_value_string *ATTRIBUTE_VALUE_STRING) (ok bool) {

	return attribute_value_string.GongIsStaged(stage)
}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ATTRIBUTE_VALUE_XHTMLs[attribute_value_xhtml]

	return
}

func (stage *Stage) IsStagedATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) (ok bool) {

	return attribute_value_xhtml.GongIsStaged(stage)
}

func (a_alternative_id *A_ALTERNATIVE_ID) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ALTERNATIVE_IDs[a_alternative_id]

	return
}

func (stage *Stage) IsStagedA_ALTERNATIVE_ID(a_alternative_id *A_ALTERNATIVE_ID) (ok bool) {

	return a_alternative_id.GongIsStaged(stage)
}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs[a_attribute_definition_boolean_ref]

	return
}

func (stage *Stage) IsStagedA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) (ok bool) {

	return a_attribute_definition_boolean_ref.GongIsStaged(stage)
}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_DEFINITION_DATE_REFs[a_attribute_definition_date_ref]

	return
}

func (stage *Stage) IsStagedA_ATTRIBUTE_DEFINITION_DATE_REF(a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) (ok bool) {

	return a_attribute_definition_date_ref.GongIsStaged(stage)
}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs[a_attribute_definition_enumeration_ref]

	return
}

func (stage *Stage) IsStagedA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) (ok bool) {

	return a_attribute_definition_enumeration_ref.GongIsStaged(stage)
}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs[a_attribute_definition_integer_ref]

	return
}

func (stage *Stage) IsStagedA_ATTRIBUTE_DEFINITION_INTEGER_REF(a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) (ok bool) {

	return a_attribute_definition_integer_ref.GongIsStaged(stage)
}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_DEFINITION_REAL_REFs[a_attribute_definition_real_ref]

	return
}

func (stage *Stage) IsStagedA_ATTRIBUTE_DEFINITION_REAL_REF(a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) (ok bool) {

	return a_attribute_definition_real_ref.GongIsStaged(stage)
}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_DEFINITION_STRING_REFs[a_attribute_definition_string_ref]

	return
}

func (stage *Stage) IsStagedA_ATTRIBUTE_DEFINITION_STRING_REF(a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) (ok bool) {

	return a_attribute_definition_string_ref.GongIsStaged(stage)
}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs[a_attribute_definition_xhtml_ref]

	return
}

func (stage *Stage) IsStagedA_ATTRIBUTE_DEFINITION_XHTML_REF(a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) (ok bool) {

	return a_attribute_definition_xhtml_ref.GongIsStaged(stage)
}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_VALUE_BOOLEANs[a_attribute_value_boolean]

	return
}

func (stage *Stage) IsStagedA_ATTRIBUTE_VALUE_BOOLEAN(a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) (ok bool) {

	return a_attribute_value_boolean.GongIsStaged(stage)
}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_VALUE_DATEs[a_attribute_value_date]

	return
}

func (stage *Stage) IsStagedA_ATTRIBUTE_VALUE_DATE(a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) (ok bool) {

	return a_attribute_value_date.GongIsStaged(stage)
}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_VALUE_ENUMERATIONs[a_attribute_value_enumeration]

	return
}

func (stage *Stage) IsStagedA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) (ok bool) {

	return a_attribute_value_enumeration.GongIsStaged(stage)
}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_VALUE_INTEGERs[a_attribute_value_integer]

	return
}

func (stage *Stage) IsStagedA_ATTRIBUTE_VALUE_INTEGER(a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) (ok bool) {

	return a_attribute_value_integer.GongIsStaged(stage)
}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_VALUE_REALs[a_attribute_value_real]

	return
}

func (stage *Stage) IsStagedA_ATTRIBUTE_VALUE_REAL(a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) (ok bool) {

	return a_attribute_value_real.GongIsStaged(stage)
}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_VALUE_STRINGs[a_attribute_value_string]

	return
}

func (stage *Stage) IsStagedA_ATTRIBUTE_VALUE_STRING(a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) (ok bool) {

	return a_attribute_value_string.GongIsStaged(stage)
}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_VALUE_XHTMLs[a_attribute_value_xhtml]

	return
}

func (stage *Stage) IsStagedA_ATTRIBUTE_VALUE_XHTML(a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) (ok bool) {

	return a_attribute_value_xhtml.GongIsStaged(stage)
}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ATTRIBUTE_VALUE_XHTML_1s[a_attribute_value_xhtml_1]

	return
}

func (stage *Stage) IsStagedA_ATTRIBUTE_VALUE_XHTML_1(a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) (ok bool) {

	return a_attribute_value_xhtml_1.GongIsStaged(stage)
}

func (a_children *A_CHILDREN) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_CHILDRENs[a_children]

	return
}

func (stage *Stage) IsStagedA_CHILDREN(a_children *A_CHILDREN) (ok bool) {

	return a_children.GongIsStaged(stage)
}

func (a_core_content *A_CORE_CONTENT) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_CORE_CONTENTs[a_core_content]

	return
}

func (stage *Stage) IsStagedA_CORE_CONTENT(a_core_content *A_CORE_CONTENT) (ok bool) {

	return a_core_content.GongIsStaged(stage)
}

func (a_datatypes *A_DATATYPES) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_DATATYPESs[a_datatypes]

	return
}

func (stage *Stage) IsStagedA_DATATYPES(a_datatypes *A_DATATYPES) (ok bool) {

	return a_datatypes.GongIsStaged(stage)
}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs[a_datatype_definition_boolean_ref]

	return
}

func (stage *Stage) IsStagedA_DATATYPE_DEFINITION_BOOLEAN_REF(a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) (ok bool) {

	return a_datatype_definition_boolean_ref.GongIsStaged(stage)
}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_DATATYPE_DEFINITION_DATE_REFs[a_datatype_definition_date_ref]

	return
}

func (stage *Stage) IsStagedA_DATATYPE_DEFINITION_DATE_REF(a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) (ok bool) {

	return a_datatype_definition_date_ref.GongIsStaged(stage)
}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs[a_datatype_definition_enumeration_ref]

	return
}

func (stage *Stage) IsStagedA_DATATYPE_DEFINITION_ENUMERATION_REF(a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) (ok bool) {

	return a_datatype_definition_enumeration_ref.GongIsStaged(stage)
}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_DATATYPE_DEFINITION_INTEGER_REFs[a_datatype_definition_integer_ref]

	return
}

func (stage *Stage) IsStagedA_DATATYPE_DEFINITION_INTEGER_REF(a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) (ok bool) {

	return a_datatype_definition_integer_ref.GongIsStaged(stage)
}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_DATATYPE_DEFINITION_REAL_REFs[a_datatype_definition_real_ref]

	return
}

func (stage *Stage) IsStagedA_DATATYPE_DEFINITION_REAL_REF(a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) (ok bool) {

	return a_datatype_definition_real_ref.GongIsStaged(stage)
}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_DATATYPE_DEFINITION_STRING_REFs[a_datatype_definition_string_ref]

	return
}

func (stage *Stage) IsStagedA_DATATYPE_DEFINITION_STRING_REF(a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) (ok bool) {

	return a_datatype_definition_string_ref.GongIsStaged(stage)
}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_DATATYPE_DEFINITION_XHTML_REFs[a_datatype_definition_xhtml_ref]

	return
}

func (stage *Stage) IsStagedA_DATATYPE_DEFINITION_XHTML_REF(a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) (ok bool) {

	return a_datatype_definition_xhtml_ref.GongIsStaged(stage)
}

func (a_editable_atts *A_EDITABLE_ATTS) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_EDITABLE_ATTSs[a_editable_atts]

	return
}

func (stage *Stage) IsStagedA_EDITABLE_ATTS(a_editable_atts *A_EDITABLE_ATTS) (ok bool) {

	return a_editable_atts.GongIsStaged(stage)
}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_ENUM_VALUE_REFs[a_enum_value_ref]

	return
}

func (stage *Stage) IsStagedA_ENUM_VALUE_REF(a_enum_value_ref *A_ENUM_VALUE_REF) (ok bool) {

	return a_enum_value_ref.GongIsStaged(stage)
}

func (a_object *A_OBJECT) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_OBJECTs[a_object]

	return
}

func (stage *Stage) IsStagedA_OBJECT(a_object *A_OBJECT) (ok bool) {

	return a_object.GongIsStaged(stage)
}

func (a_properties *A_PROPERTIES) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_PROPERTIESs[a_properties]

	return
}

func (stage *Stage) IsStagedA_PROPERTIES(a_properties *A_PROPERTIES) (ok bool) {

	return a_properties.GongIsStaged(stage)
}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_RELATION_GROUP_TYPE_REFs[a_relation_group_type_ref]

	return
}

func (stage *Stage) IsStagedA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) (ok bool) {

	return a_relation_group_type_ref.GongIsStaged(stage)
}

func (a_source_1 *A_SOURCE_1) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_SOURCE_1s[a_source_1]

	return
}

func (stage *Stage) IsStagedA_SOURCE_1(a_source_1 *A_SOURCE_1) (ok bool) {

	return a_source_1.GongIsStaged(stage)
}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_SOURCE_SPECIFICATION_1s[a_source_specification_1]

	return
}

func (stage *Stage) IsStagedA_SOURCE_SPECIFICATION_1(a_source_specification_1 *A_SOURCE_SPECIFICATION_1) (ok bool) {

	return a_source_specification_1.GongIsStaged(stage)
}

func (a_specifications *A_SPECIFICATIONS) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_SPECIFICATIONSs[a_specifications]

	return
}

func (stage *Stage) IsStagedA_SPECIFICATIONS(a_specifications *A_SPECIFICATIONS) (ok bool) {

	return a_specifications.GongIsStaged(stage)
}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_SPECIFICATION_TYPE_REFs[a_specification_type_ref]

	return
}

func (stage *Stage) IsStagedA_SPECIFICATION_TYPE_REF(a_specification_type_ref *A_SPECIFICATION_TYPE_REF) (ok bool) {

	return a_specification_type_ref.GongIsStaged(stage)
}

func (a_specified_values *A_SPECIFIED_VALUES) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_SPECIFIED_VALUESs[a_specified_values]

	return
}

func (stage *Stage) IsStagedA_SPECIFIED_VALUES(a_specified_values *A_SPECIFIED_VALUES) (ok bool) {

	return a_specified_values.GongIsStaged(stage)
}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_SPEC_ATTRIBUTESs[a_spec_attributes]

	return
}

func (stage *Stage) IsStagedA_SPEC_ATTRIBUTES(a_spec_attributes *A_SPEC_ATTRIBUTES) (ok bool) {

	return a_spec_attributes.GongIsStaged(stage)
}

func (a_spec_objects *A_SPEC_OBJECTS) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_SPEC_OBJECTSs[a_spec_objects]

	return
}

func (stage *Stage) IsStagedA_SPEC_OBJECTS(a_spec_objects *A_SPEC_OBJECTS) (ok bool) {

	return a_spec_objects.GongIsStaged(stage)
}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_SPEC_OBJECT_TYPE_REFs[a_spec_object_type_ref]

	return
}

func (stage *Stage) IsStagedA_SPEC_OBJECT_TYPE_REF(a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) (ok bool) {

	return a_spec_object_type_ref.GongIsStaged(stage)
}

func (a_spec_relations *A_SPEC_RELATIONS) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_SPEC_RELATIONSs[a_spec_relations]

	return
}

func (stage *Stage) IsStagedA_SPEC_RELATIONS(a_spec_relations *A_SPEC_RELATIONS) (ok bool) {

	return a_spec_relations.GongIsStaged(stage)
}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_SPEC_RELATION_GROUPSs[a_spec_relation_groups]

	return
}

func (stage *Stage) IsStagedA_SPEC_RELATION_GROUPS(a_spec_relation_groups *A_SPEC_RELATION_GROUPS) (ok bool) {

	return a_spec_relation_groups.GongIsStaged(stage)
}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_SPEC_RELATION_REFs[a_spec_relation_ref]

	return
}

func (stage *Stage) IsStagedA_SPEC_RELATION_REF(a_spec_relation_ref *A_SPEC_RELATION_REF) (ok bool) {

	return a_spec_relation_ref.GongIsStaged(stage)
}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_SPEC_RELATION_TYPE_REFs[a_spec_relation_type_ref]

	return
}

func (stage *Stage) IsStagedA_SPEC_RELATION_TYPE_REF(a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) (ok bool) {

	return a_spec_relation_type_ref.GongIsStaged(stage)
}

func (a_spec_types *A_SPEC_TYPES) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_SPEC_TYPESs[a_spec_types]

	return
}

func (stage *Stage) IsStagedA_SPEC_TYPES(a_spec_types *A_SPEC_TYPES) (ok bool) {

	return a_spec_types.GongIsStaged(stage)
}

func (a_the_header *A_THE_HEADER) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_THE_HEADERs[a_the_header]

	return
}

func (stage *Stage) IsStagedA_THE_HEADER(a_the_header *A_THE_HEADER) (ok bool) {

	return a_the_header.GongIsStaged(stage)
}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.A_TOOL_EXTENSIONSs[a_tool_extensions]

	return
}

func (stage *Stage) IsStagedA_TOOL_EXTENSIONS(a_tool_extensions *A_TOOL_EXTENSIONS) (ok bool) {

	return a_tool_extensions.GongIsStaged(stage)
}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_BOOLEANs[datatype_definition_boolean]

	return
}

func (stage *Stage) IsStagedDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) (ok bool) {

	return datatype_definition_boolean.GongIsStaged(stage)
}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_DATEs[datatype_definition_date]

	return
}

func (stage *Stage) IsStagedDATATYPE_DEFINITION_DATE(datatype_definition_date *DATATYPE_DEFINITION_DATE) (ok bool) {

	return datatype_definition_date.GongIsStaged(stage)
}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_ENUMERATIONs[datatype_definition_enumeration]

	return
}

func (stage *Stage) IsStagedDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) (ok bool) {

	return datatype_definition_enumeration.GongIsStaged(stage)
}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_INTEGERs[datatype_definition_integer]

	return
}

func (stage *Stage) IsStagedDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) (ok bool) {

	return datatype_definition_integer.GongIsStaged(stage)
}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_REALs[datatype_definition_real]

	return
}

func (stage *Stage) IsStagedDATATYPE_DEFINITION_REAL(datatype_definition_real *DATATYPE_DEFINITION_REAL) (ok bool) {

	return datatype_definition_real.GongIsStaged(stage)
}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_STRINGs[datatype_definition_string]

	return
}

func (stage *Stage) IsStagedDATATYPE_DEFINITION_STRING(datatype_definition_string *DATATYPE_DEFINITION_STRING) (ok bool) {

	return datatype_definition_string.GongIsStaged(stage)
}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.DATATYPE_DEFINITION_XHTMLs[datatype_definition_xhtml]

	return
}

func (stage *Stage) IsStagedDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) (ok bool) {

	return datatype_definition_xhtml.GongIsStaged(stage)
}

func (embedded_value *EMBEDDED_VALUE) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.EMBEDDED_VALUEs[embedded_value]

	return
}

func (stage *Stage) IsStagedEMBEDDED_VALUE(embedded_value *EMBEDDED_VALUE) (ok bool) {

	return embedded_value.GongIsStaged(stage)
}

func (enum_value *ENUM_VALUE) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.ENUM_VALUEs[enum_value]

	return
}

func (stage *Stage) IsStagedENUM_VALUE(enum_value *ENUM_VALUE) (ok bool) {

	return enum_value.GongIsStaged(stage)
}

func (relation_group *RELATION_GROUP) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.RELATION_GROUPs[relation_group]

	return
}

func (stage *Stage) IsStagedRELATION_GROUP(relation_group *RELATION_GROUP) (ok bool) {

	return relation_group.GongIsStaged(stage)
}

func (relation_group_type *RELATION_GROUP_TYPE) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.RELATION_GROUP_TYPEs[relation_group_type]

	return
}

func (stage *Stage) IsStagedRELATION_GROUP_TYPE(relation_group_type *RELATION_GROUP_TYPE) (ok bool) {

	return relation_group_type.GongIsStaged(stage)
}

func (req_if *REQ_IF) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.REQ_IFs[req_if]

	return
}

func (stage *Stage) IsStagedREQ_IF(req_if *REQ_IF) (ok bool) {

	return req_if.GongIsStaged(stage)
}

func (req_if_content *REQ_IF_CONTENT) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.REQ_IF_CONTENTs[req_if_content]

	return
}

func (stage *Stage) IsStagedREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT) (ok bool) {

	return req_if_content.GongIsStaged(stage)
}

func (req_if_header *REQ_IF_HEADER) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.REQ_IF_HEADERs[req_if_header]

	return
}

func (stage *Stage) IsStagedREQ_IF_HEADER(req_if_header *REQ_IF_HEADER) (ok bool) {

	return req_if_header.GongIsStaged(stage)
}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.REQ_IF_TOOL_EXTENSIONs[req_if_tool_extension]

	return
}

func (stage *Stage) IsStagedREQ_IF_TOOL_EXTENSION(req_if_tool_extension *REQ_IF_TOOL_EXTENSION) (ok bool) {

	return req_if_tool_extension.GongIsStaged(stage)
}

func (specification *SPECIFICATION) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SPECIFICATIONs[specification]

	return
}

func (stage *Stage) IsStagedSPECIFICATION(specification *SPECIFICATION) (ok bool) {

	return specification.GongIsStaged(stage)
}

func (specification_type *SPECIFICATION_TYPE) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SPECIFICATION_TYPEs[specification_type]

	return
}

func (stage *Stage) IsStagedSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE) (ok bool) {

	return specification_type.GongIsStaged(stage)
}

func (spec_hierarchy *SPEC_HIERARCHY) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SPEC_HIERARCHYs[spec_hierarchy]

	return
}

func (stage *Stage) IsStagedSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY) (ok bool) {

	return spec_hierarchy.GongIsStaged(stage)
}

func (spec_object *SPEC_OBJECT) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SPEC_OBJECTs[spec_object]

	return
}

func (stage *Stage) IsStagedSPEC_OBJECT(spec_object *SPEC_OBJECT) (ok bool) {

	return spec_object.GongIsStaged(stage)
}

func (spec_object_type *SPEC_OBJECT_TYPE) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SPEC_OBJECT_TYPEs[spec_object_type]

	return
}

func (stage *Stage) IsStagedSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE) (ok bool) {

	return spec_object_type.GongIsStaged(stage)
}

func (spec_relation *SPEC_RELATION) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SPEC_RELATIONs[spec_relation]

	return
}

func (stage *Stage) IsStagedSPEC_RELATION(spec_relation *SPEC_RELATION) (ok bool) {

	return spec_relation.GongIsStaged(stage)
}

func (spec_relation_type *SPEC_RELATION_TYPE) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.SPEC_RELATION_TYPEs[spec_relation_type]

	return
}

func (stage *Stage) IsStagedSPEC_RELATION_TYPE(spec_relation_type *SPEC_RELATION_TYPE) (ok bool) {

	return spec_relation_type.GongIsStaged(stage)
}

func (xhtml_content *XHTML_CONTENT) GongIsStaged(stage *Stage) (ok bool) {

	_, ok = stage.XHTML_CONTENTs[xhtml_content]

	return
}

func (stage *Stage) IsStagedXHTML_CONTENT(xhtml_content *XHTML_CONTENT) (ok bool) {

	return xhtml_content.GongIsStaged(stage)
}

// StageBranch is the Stage method that stages instance and applies StageBranch recursively.
func (stage *Stage) StageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongStageBranch(stage)
	}
}

// StageBranch is a backward-compatible package-level forwarder.
func StageBranch(stage *Stage, instance GongstructIF) {
	stage.StageBranch(instance)
}

// insertion point for stage branch per struct
func (alternative_id *ALTERNATIVE_ID) GongStageBranch(stage *Stage) {
	stage.StageBranchALTERNATIVE_ID(alternative_id)
}

func (stage *Stage) StageBranchALTERNATIVE_ID(alternative_id *ALTERNATIVE_ID) {

	// check if instance is already staged
	if stage.IsStaged(alternative_id) {
		return
	}

	alternative_id.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongStageBranch(stage *Stage) {
	stage.StageBranchATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean)
}

func (stage *Stage) StageBranchATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) {

	// check if instance is already staged
	if stage.IsStaged(attribute_definition_boolean) {
		return
	}

	attribute_definition_boolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_boolean.ALTERNATIVE_ID != nil {
		stage.StageBranch(attribute_definition_boolean.ALTERNATIVE_ID)
	}
	if attribute_definition_boolean.DEFAULT_VALUE != nil {
		stage.StageBranch(attribute_definition_boolean.DEFAULT_VALUE)
	}
	if attribute_definition_boolean.TYPE != nil {
		stage.StageBranch(attribute_definition_boolean.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongStageBranch(stage *Stage) {
	stage.StageBranchATTRIBUTE_DEFINITION_DATE(attribute_definition_date)
}

func (stage *Stage) StageBranchATTRIBUTE_DEFINITION_DATE(attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) {

	// check if instance is already staged
	if stage.IsStaged(attribute_definition_date) {
		return
	}

	attribute_definition_date.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_date.ALTERNATIVE_ID != nil {
		stage.StageBranch(attribute_definition_date.ALTERNATIVE_ID)
	}
	if attribute_definition_date.DEFAULT_VALUE != nil {
		stage.StageBranch(attribute_definition_date.DEFAULT_VALUE)
	}
	if attribute_definition_date.TYPE != nil {
		stage.StageBranch(attribute_definition_date.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongStageBranch(stage *Stage) {
	stage.StageBranchATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration)
}

func (stage *Stage) StageBranchATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) {

	// check if instance is already staged
	if stage.IsStaged(attribute_definition_enumeration) {
		return
	}

	attribute_definition_enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_enumeration.ALTERNATIVE_ID != nil {
		stage.StageBranch(attribute_definition_enumeration.ALTERNATIVE_ID)
	}
	if attribute_definition_enumeration.DEFAULT_VALUE != nil {
		stage.StageBranch(attribute_definition_enumeration.DEFAULT_VALUE)
	}
	if attribute_definition_enumeration.TYPE != nil {
		stage.StageBranch(attribute_definition_enumeration.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongStageBranch(stage *Stage) {
	stage.StageBranchATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer)
}

func (stage *Stage) StageBranchATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) {

	// check if instance is already staged
	if stage.IsStaged(attribute_definition_integer) {
		return
	}

	attribute_definition_integer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_integer.ALTERNATIVE_ID != nil {
		stage.StageBranch(attribute_definition_integer.ALTERNATIVE_ID)
	}
	if attribute_definition_integer.DEFAULT_VALUE != nil {
		stage.StageBranch(attribute_definition_integer.DEFAULT_VALUE)
	}
	if attribute_definition_integer.TYPE != nil {
		stage.StageBranch(attribute_definition_integer.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongStageBranch(stage *Stage) {
	stage.StageBranchATTRIBUTE_DEFINITION_REAL(attribute_definition_real)
}

func (stage *Stage) StageBranchATTRIBUTE_DEFINITION_REAL(attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) {

	// check if instance is already staged
	if stage.IsStaged(attribute_definition_real) {
		return
	}

	attribute_definition_real.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_real.ALTERNATIVE_ID != nil {
		stage.StageBranch(attribute_definition_real.ALTERNATIVE_ID)
	}
	if attribute_definition_real.DEFAULT_VALUE != nil {
		stage.StageBranch(attribute_definition_real.DEFAULT_VALUE)
	}
	if attribute_definition_real.TYPE != nil {
		stage.StageBranch(attribute_definition_real.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongStageBranch(stage *Stage) {
	stage.StageBranchATTRIBUTE_DEFINITION_STRING(attribute_definition_string)
}

func (stage *Stage) StageBranchATTRIBUTE_DEFINITION_STRING(attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) {

	// check if instance is already staged
	if stage.IsStaged(attribute_definition_string) {
		return
	}

	attribute_definition_string.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_string.ALTERNATIVE_ID != nil {
		stage.StageBranch(attribute_definition_string.ALTERNATIVE_ID)
	}
	if attribute_definition_string.DEFAULT_VALUE != nil {
		stage.StageBranch(attribute_definition_string.DEFAULT_VALUE)
	}
	if attribute_definition_string.TYPE != nil {
		stage.StageBranch(attribute_definition_string.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongStageBranch(stage *Stage) {
	stage.StageBranchATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml)
}

func (stage *Stage) StageBranchATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) {

	// check if instance is already staged
	if stage.IsStaged(attribute_definition_xhtml) {
		return
	}

	attribute_definition_xhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_xhtml.ALTERNATIVE_ID != nil {
		stage.StageBranch(attribute_definition_xhtml.ALTERNATIVE_ID)
	}
	if attribute_definition_xhtml.DEFAULT_VALUE != nil {
		stage.StageBranch(attribute_definition_xhtml.DEFAULT_VALUE)
	}
	if attribute_definition_xhtml.TYPE != nil {
		stage.StageBranch(attribute_definition_xhtml.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongStageBranch(stage *Stage) {
	stage.StageBranchATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean)
}

func (stage *Stage) StageBranchATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) {

	// check if instance is already staged
	if stage.IsStaged(attribute_value_boolean) {
		return
	}

	attribute_value_boolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_boolean.DEFINITION != nil {
		stage.StageBranch(attribute_value_boolean.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongStageBranch(stage *Stage) {
	stage.StageBranchATTRIBUTE_VALUE_DATE(attribute_value_date)
}

func (stage *Stage) StageBranchATTRIBUTE_VALUE_DATE(attribute_value_date *ATTRIBUTE_VALUE_DATE) {

	// check if instance is already staged
	if stage.IsStaged(attribute_value_date) {
		return
	}

	attribute_value_date.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_date.DEFINITION != nil {
		stage.StageBranch(attribute_value_date.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongStageBranch(stage *Stage) {
	stage.StageBranchATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration)
}

func (stage *Stage) StageBranchATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) {

	// check if instance is already staged
	if stage.IsStaged(attribute_value_enumeration) {
		return
	}

	attribute_value_enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_enumeration.DEFINITION != nil {
		stage.StageBranch(attribute_value_enumeration.DEFINITION)
	}
	if attribute_value_enumeration.VALUES != nil {
		stage.StageBranch(attribute_value_enumeration.VALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongStageBranch(stage *Stage) {
	stage.StageBranchATTRIBUTE_VALUE_INTEGER(attribute_value_integer)
}

func (stage *Stage) StageBranchATTRIBUTE_VALUE_INTEGER(attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) {

	// check if instance is already staged
	if stage.IsStaged(attribute_value_integer) {
		return
	}

	attribute_value_integer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_integer.DEFINITION != nil {
		stage.StageBranch(attribute_value_integer.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongStageBranch(stage *Stage) {
	stage.StageBranchATTRIBUTE_VALUE_REAL(attribute_value_real)
}

func (stage *Stage) StageBranchATTRIBUTE_VALUE_REAL(attribute_value_real *ATTRIBUTE_VALUE_REAL) {

	// check if instance is already staged
	if stage.IsStaged(attribute_value_real) {
		return
	}

	attribute_value_real.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_real.DEFINITION != nil {
		stage.StageBranch(attribute_value_real.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongStageBranch(stage *Stage) {
	stage.StageBranchATTRIBUTE_VALUE_STRING(attribute_value_string)
}

func (stage *Stage) StageBranchATTRIBUTE_VALUE_STRING(attribute_value_string *ATTRIBUTE_VALUE_STRING) {

	// check if instance is already staged
	if stage.IsStaged(attribute_value_string) {
		return
	}

	attribute_value_string.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_string.DEFINITION != nil {
		stage.StageBranch(attribute_value_string.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongStageBranch(stage *Stage) {
	stage.StageBranchATTRIBUTE_VALUE_XHTML(attribute_value_xhtml)
}

func (stage *Stage) StageBranchATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) {

	// check if instance is already staged
	if stage.IsStaged(attribute_value_xhtml) {
		return
	}

	attribute_value_xhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_xhtml.THE_VALUE != nil {
		stage.StageBranch(attribute_value_xhtml.THE_VALUE)
	}
	if attribute_value_xhtml.THE_ORIGINAL_VALUE != nil {
		stage.StageBranch(attribute_value_xhtml.THE_ORIGINAL_VALUE)
	}
	if attribute_value_xhtml.DEFINITION != nil {
		stage.StageBranch(attribute_value_xhtml.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_alternative_id *A_ALTERNATIVE_ID) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ALTERNATIVE_ID(a_alternative_id)
}

func (stage *Stage) StageBranchA_ALTERNATIVE_ID(a_alternative_id *A_ALTERNATIVE_ID) {

	// check if instance is already staged
	if stage.IsStaged(a_alternative_id) {
		return
	}

	a_alternative_id.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_alternative_id.ALTERNATIVE_ID != nil {
		stage.StageBranch(a_alternative_id.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(a_attribute_definition_boolean_ref)
}

func (stage *Stage) StageBranchA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_attribute_definition_boolean_ref) {
		return
	}

	a_attribute_definition_boolean_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ATTRIBUTE_DEFINITION_DATE_REF(a_attribute_definition_date_ref)
}

func (stage *Stage) StageBranchA_ATTRIBUTE_DEFINITION_DATE_REF(a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_attribute_definition_date_ref) {
		return
	}

	a_attribute_definition_date_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(a_attribute_definition_enumeration_ref)
}

func (stage *Stage) StageBranchA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_attribute_definition_enumeration_ref) {
		return
	}

	a_attribute_definition_enumeration_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ATTRIBUTE_DEFINITION_INTEGER_REF(a_attribute_definition_integer_ref)
}

func (stage *Stage) StageBranchA_ATTRIBUTE_DEFINITION_INTEGER_REF(a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_attribute_definition_integer_ref) {
		return
	}

	a_attribute_definition_integer_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ATTRIBUTE_DEFINITION_REAL_REF(a_attribute_definition_real_ref)
}

func (stage *Stage) StageBranchA_ATTRIBUTE_DEFINITION_REAL_REF(a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_attribute_definition_real_ref) {
		return
	}

	a_attribute_definition_real_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ATTRIBUTE_DEFINITION_STRING_REF(a_attribute_definition_string_ref)
}

func (stage *Stage) StageBranchA_ATTRIBUTE_DEFINITION_STRING_REF(a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_attribute_definition_string_ref) {
		return
	}

	a_attribute_definition_string_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ATTRIBUTE_DEFINITION_XHTML_REF(a_attribute_definition_xhtml_ref)
}

func (stage *Stage) StageBranchA_ATTRIBUTE_DEFINITION_XHTML_REF(a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_attribute_definition_xhtml_ref) {
		return
	}

	a_attribute_definition_xhtml_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ATTRIBUTE_VALUE_BOOLEAN(a_attribute_value_boolean)
}

func (stage *Stage) StageBranchA_ATTRIBUTE_VALUE_BOOLEAN(a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) {

	// check if instance is already staged
	if stage.IsStaged(a_attribute_value_boolean) {
		return
	}

	a_attribute_value_boolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_boolean := range a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN {
		stage.StageBranch(_attribute_value_boolean)
	}

}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ATTRIBUTE_VALUE_DATE(a_attribute_value_date)
}

func (stage *Stage) StageBranchA_ATTRIBUTE_VALUE_DATE(a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) {

	// check if instance is already staged
	if stage.IsStaged(a_attribute_value_date) {
		return
	}

	a_attribute_value_date.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_date := range a_attribute_value_date.ATTRIBUTE_VALUE_DATE {
		stage.StageBranch(_attribute_value_date)
	}

}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration)
}

func (stage *Stage) StageBranchA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) {

	// check if instance is already staged
	if stage.IsStaged(a_attribute_value_enumeration) {
		return
	}

	a_attribute_value_enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_enumeration := range a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION {
		stage.StageBranch(_attribute_value_enumeration)
	}

}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ATTRIBUTE_VALUE_INTEGER(a_attribute_value_integer)
}

func (stage *Stage) StageBranchA_ATTRIBUTE_VALUE_INTEGER(a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) {

	// check if instance is already staged
	if stage.IsStaged(a_attribute_value_integer) {
		return
	}

	a_attribute_value_integer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_integer := range a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER {
		stage.StageBranch(_attribute_value_integer)
	}

}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ATTRIBUTE_VALUE_REAL(a_attribute_value_real)
}

func (stage *Stage) StageBranchA_ATTRIBUTE_VALUE_REAL(a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) {

	// check if instance is already staged
	if stage.IsStaged(a_attribute_value_real) {
		return
	}

	a_attribute_value_real.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_real := range a_attribute_value_real.ATTRIBUTE_VALUE_REAL {
		stage.StageBranch(_attribute_value_real)
	}

}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ATTRIBUTE_VALUE_STRING(a_attribute_value_string)
}

func (stage *Stage) StageBranchA_ATTRIBUTE_VALUE_STRING(a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) {

	// check if instance is already staged
	if stage.IsStaged(a_attribute_value_string) {
		return
	}

	a_attribute_value_string.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_string := range a_attribute_value_string.ATTRIBUTE_VALUE_STRING {
		stage.StageBranch(_attribute_value_string)
	}

}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ATTRIBUTE_VALUE_XHTML(a_attribute_value_xhtml)
}

func (stage *Stage) StageBranchA_ATTRIBUTE_VALUE_XHTML(a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) {

	// check if instance is already staged
	if stage.IsStaged(a_attribute_value_xhtml) {
		return
	}

	a_attribute_value_xhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_xhtml := range a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML {
		stage.StageBranch(_attribute_value_xhtml)
	}

}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ATTRIBUTE_VALUE_XHTML_1(a_attribute_value_xhtml_1)
}

func (stage *Stage) StageBranchA_ATTRIBUTE_VALUE_XHTML_1(a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) {

	// check if instance is already staged
	if stage.IsStaged(a_attribute_value_xhtml_1) {
		return
	}

	a_attribute_value_xhtml_1.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_boolean := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN {
		stage.StageBranch(_attribute_value_boolean)
	}
	for _, _attribute_value_date := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE {
		stage.StageBranch(_attribute_value_date)
	}
	for _, _attribute_value_enumeration := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION {
		stage.StageBranch(_attribute_value_enumeration)
	}
	for _, _attribute_value_integer := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER {
		stage.StageBranch(_attribute_value_integer)
	}
	for _, _attribute_value_real := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL {
		stage.StageBranch(_attribute_value_real)
	}
	for _, _attribute_value_string := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING {
		stage.StageBranch(_attribute_value_string)
	}
	for _, _attribute_value_xhtml := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML {
		stage.StageBranch(_attribute_value_xhtml)
	}

}

func (a_children *A_CHILDREN) GongStageBranch(stage *Stage) {
	stage.StageBranchA_CHILDREN(a_children)
}

func (stage *Stage) StageBranchA_CHILDREN(a_children *A_CHILDREN) {

	// check if instance is already staged
	if stage.IsStaged(a_children) {
		return
	}

	a_children.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_hierarchy := range a_children.SPEC_HIERARCHY {
		stage.StageBranch(_spec_hierarchy)
	}

}

func (a_core_content *A_CORE_CONTENT) GongStageBranch(stage *Stage) {
	stage.StageBranchA_CORE_CONTENT(a_core_content)
}

func (stage *Stage) StageBranchA_CORE_CONTENT(a_core_content *A_CORE_CONTENT) {

	// check if instance is already staged
	if stage.IsStaged(a_core_content) {
		return
	}

	a_core_content.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_core_content.REQ_IF_CONTENT != nil {
		stage.StageBranch(a_core_content.REQ_IF_CONTENT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_datatypes *A_DATATYPES) GongStageBranch(stage *Stage) {
	stage.StageBranchA_DATATYPES(a_datatypes)
}

func (stage *Stage) StageBranchA_DATATYPES(a_datatypes *A_DATATYPES) {

	// check if instance is already staged
	if stage.IsStaged(a_datatypes) {
		return
	}

	a_datatypes.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _datatype_definition_boolean := range a_datatypes.DATATYPE_DEFINITION_BOOLEAN {
		stage.StageBranch(_datatype_definition_boolean)
	}
	for _, _datatype_definition_date := range a_datatypes.DATATYPE_DEFINITION_DATE {
		stage.StageBranch(_datatype_definition_date)
	}
	for _, _datatype_definition_enumeration := range a_datatypes.DATATYPE_DEFINITION_ENUMERATION {
		stage.StageBranch(_datatype_definition_enumeration)
	}
	for _, _datatype_definition_integer := range a_datatypes.DATATYPE_DEFINITION_INTEGER {
		stage.StageBranch(_datatype_definition_integer)
	}
	for _, _datatype_definition_real := range a_datatypes.DATATYPE_DEFINITION_REAL {
		stage.StageBranch(_datatype_definition_real)
	}
	for _, _datatype_definition_string := range a_datatypes.DATATYPE_DEFINITION_STRING {
		stage.StageBranch(_datatype_definition_string)
	}
	for _, _datatype_definition_xhtml := range a_datatypes.DATATYPE_DEFINITION_XHTML {
		stage.StageBranch(_datatype_definition_xhtml)
	}

}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_DATATYPE_DEFINITION_BOOLEAN_REF(a_datatype_definition_boolean_ref)
}

func (stage *Stage) StageBranchA_DATATYPE_DEFINITION_BOOLEAN_REF(a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_datatype_definition_boolean_ref) {
		return
	}

	a_datatype_definition_boolean_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_DATATYPE_DEFINITION_DATE_REF(a_datatype_definition_date_ref)
}

func (stage *Stage) StageBranchA_DATATYPE_DEFINITION_DATE_REF(a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_datatype_definition_date_ref) {
		return
	}

	a_datatype_definition_date_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_DATATYPE_DEFINITION_ENUMERATION_REF(a_datatype_definition_enumeration_ref)
}

func (stage *Stage) StageBranchA_DATATYPE_DEFINITION_ENUMERATION_REF(a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_datatype_definition_enumeration_ref) {
		return
	}

	a_datatype_definition_enumeration_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_DATATYPE_DEFINITION_INTEGER_REF(a_datatype_definition_integer_ref)
}

func (stage *Stage) StageBranchA_DATATYPE_DEFINITION_INTEGER_REF(a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_datatype_definition_integer_ref) {
		return
	}

	a_datatype_definition_integer_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_DATATYPE_DEFINITION_REAL_REF(a_datatype_definition_real_ref)
}

func (stage *Stage) StageBranchA_DATATYPE_DEFINITION_REAL_REF(a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_datatype_definition_real_ref) {
		return
	}

	a_datatype_definition_real_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_DATATYPE_DEFINITION_STRING_REF(a_datatype_definition_string_ref)
}

func (stage *Stage) StageBranchA_DATATYPE_DEFINITION_STRING_REF(a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_datatype_definition_string_ref) {
		return
	}

	a_datatype_definition_string_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_DATATYPE_DEFINITION_XHTML_REF(a_datatype_definition_xhtml_ref)
}

func (stage *Stage) StageBranchA_DATATYPE_DEFINITION_XHTML_REF(a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_datatype_definition_xhtml_ref) {
		return
	}

	a_datatype_definition_xhtml_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_editable_atts *A_EDITABLE_ATTS) GongStageBranch(stage *Stage) {
	stage.StageBranchA_EDITABLE_ATTS(a_editable_atts)
}

func (stage *Stage) StageBranchA_EDITABLE_ATTS(a_editable_atts *A_EDITABLE_ATTS) {

	// check if instance is already staged
	if stage.IsStaged(a_editable_atts) {
		return
	}

	a_editable_atts.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_ENUM_VALUE_REF(a_enum_value_ref)
}

func (stage *Stage) StageBranchA_ENUM_VALUE_REF(a_enum_value_ref *A_ENUM_VALUE_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_enum_value_ref) {
		return
	}

	a_enum_value_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_object *A_OBJECT) GongStageBranch(stage *Stage) {
	stage.StageBranchA_OBJECT(a_object)
}

func (stage *Stage) StageBranchA_OBJECT(a_object *A_OBJECT) {

	// check if instance is already staged
	if stage.IsStaged(a_object) {
		return
	}

	a_object.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_properties *A_PROPERTIES) GongStageBranch(stage *Stage) {
	stage.StageBranchA_PROPERTIES(a_properties)
}

func (stage *Stage) StageBranchA_PROPERTIES(a_properties *A_PROPERTIES) {

	// check if instance is already staged
	if stage.IsStaged(a_properties) {
		return
	}

	a_properties.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_properties.EMBEDDED_VALUE != nil {
		stage.StageBranch(a_properties.EMBEDDED_VALUE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref)
}

func (stage *Stage) StageBranchA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_relation_group_type_ref) {
		return
	}

	a_relation_group_type_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_source_1 *A_SOURCE_1) GongStageBranch(stage *Stage) {
	stage.StageBranchA_SOURCE_1(a_source_1)
}

func (stage *Stage) StageBranchA_SOURCE_1(a_source_1 *A_SOURCE_1) {

	// check if instance is already staged
	if stage.IsStaged(a_source_1) {
		return
	}

	a_source_1.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongStageBranch(stage *Stage) {
	stage.StageBranchA_SOURCE_SPECIFICATION_1(a_source_specification_1)
}

func (stage *Stage) StageBranchA_SOURCE_SPECIFICATION_1(a_source_specification_1 *A_SOURCE_SPECIFICATION_1) {

	// check if instance is already staged
	if stage.IsStaged(a_source_specification_1) {
		return
	}

	a_source_specification_1.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_specifications *A_SPECIFICATIONS) GongStageBranch(stage *Stage) {
	stage.StageBranchA_SPECIFICATIONS(a_specifications)
}

func (stage *Stage) StageBranchA_SPECIFICATIONS(a_specifications *A_SPECIFICATIONS) {

	// check if instance is already staged
	if stage.IsStaged(a_specifications) {
		return
	}

	a_specifications.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _specification := range a_specifications.SPECIFICATION {
		stage.StageBranch(_specification)
	}

}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_SPECIFICATION_TYPE_REF(a_specification_type_ref)
}

func (stage *Stage) StageBranchA_SPECIFICATION_TYPE_REF(a_specification_type_ref *A_SPECIFICATION_TYPE_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_specification_type_ref) {
		return
	}

	a_specification_type_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_specified_values *A_SPECIFIED_VALUES) GongStageBranch(stage *Stage) {
	stage.StageBranchA_SPECIFIED_VALUES(a_specified_values)
}

func (stage *Stage) StageBranchA_SPECIFIED_VALUES(a_specified_values *A_SPECIFIED_VALUES) {

	// check if instance is already staged
	if stage.IsStaged(a_specified_values) {
		return
	}

	a_specified_values.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enum_value := range a_specified_values.ENUM_VALUE {
		stage.StageBranch(_enum_value)
	}

}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongStageBranch(stage *Stage) {
	stage.StageBranchA_SPEC_ATTRIBUTES(a_spec_attributes)
}

func (stage *Stage) StageBranchA_SPEC_ATTRIBUTES(a_spec_attributes *A_SPEC_ATTRIBUTES) {

	// check if instance is already staged
	if stage.IsStaged(a_spec_attributes) {
		return
	}

	a_spec_attributes.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_definition_boolean := range a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN {
		stage.StageBranch(_attribute_definition_boolean)
	}
	for _, _attribute_definition_date := range a_spec_attributes.ATTRIBUTE_DEFINITION_DATE {
		stage.StageBranch(_attribute_definition_date)
	}
	for _, _attribute_definition_enumeration := range a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION {
		stage.StageBranch(_attribute_definition_enumeration)
	}
	for _, _attribute_definition_integer := range a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER {
		stage.StageBranch(_attribute_definition_integer)
	}
	for _, _attribute_definition_real := range a_spec_attributes.ATTRIBUTE_DEFINITION_REAL {
		stage.StageBranch(_attribute_definition_real)
	}
	for _, _attribute_definition_string := range a_spec_attributes.ATTRIBUTE_DEFINITION_STRING {
		stage.StageBranch(_attribute_definition_string)
	}
	for _, _attribute_definition_xhtml := range a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML {
		stage.StageBranch(_attribute_definition_xhtml)
	}

}

func (a_spec_objects *A_SPEC_OBJECTS) GongStageBranch(stage *Stage) {
	stage.StageBranchA_SPEC_OBJECTS(a_spec_objects)
}

func (stage *Stage) StageBranchA_SPEC_OBJECTS(a_spec_objects *A_SPEC_OBJECTS) {

	// check if instance is already staged
	if stage.IsStaged(a_spec_objects) {
		return
	}

	a_spec_objects.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_object := range a_spec_objects.SPEC_OBJECT {
		stage.StageBranch(_spec_object)
	}

}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_SPEC_OBJECT_TYPE_REF(a_spec_object_type_ref)
}

func (stage *Stage) StageBranchA_SPEC_OBJECT_TYPE_REF(a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_spec_object_type_ref) {
		return
	}

	a_spec_object_type_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_spec_relations *A_SPEC_RELATIONS) GongStageBranch(stage *Stage) {
	stage.StageBranchA_SPEC_RELATIONS(a_spec_relations)
}

func (stage *Stage) StageBranchA_SPEC_RELATIONS(a_spec_relations *A_SPEC_RELATIONS) {

	// check if instance is already staged
	if stage.IsStaged(a_spec_relations) {
		return
	}

	a_spec_relations.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_relation := range a_spec_relations.SPEC_RELATION {
		stage.StageBranch(_spec_relation)
	}

}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongStageBranch(stage *Stage) {
	stage.StageBranchA_SPEC_RELATION_GROUPS(a_spec_relation_groups)
}

func (stage *Stage) StageBranchA_SPEC_RELATION_GROUPS(a_spec_relation_groups *A_SPEC_RELATION_GROUPS) {

	// check if instance is already staged
	if stage.IsStaged(a_spec_relation_groups) {
		return
	}

	a_spec_relation_groups.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relation_group := range a_spec_relation_groups.RELATION_GROUP {
		stage.StageBranch(_relation_group)
	}

}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_SPEC_RELATION_REF(a_spec_relation_ref)
}

func (stage *Stage) StageBranchA_SPEC_RELATION_REF(a_spec_relation_ref *A_SPEC_RELATION_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_spec_relation_ref) {
		return
	}

	a_spec_relation_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongStageBranch(stage *Stage) {
	stage.StageBranchA_SPEC_RELATION_TYPE_REF(a_spec_relation_type_ref)
}

func (stage *Stage) StageBranchA_SPEC_RELATION_TYPE_REF(a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) {

	// check if instance is already staged
	if stage.IsStaged(a_spec_relation_type_ref) {
		return
	}

	a_spec_relation_type_ref.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_spec_types *A_SPEC_TYPES) GongStageBranch(stage *Stage) {
	stage.StageBranchA_SPEC_TYPES(a_spec_types)
}

func (stage *Stage) StageBranchA_SPEC_TYPES(a_spec_types *A_SPEC_TYPES) {

	// check if instance is already staged
	if stage.IsStaged(a_spec_types) {
		return
	}

	a_spec_types.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relation_group_type := range a_spec_types.RELATION_GROUP_TYPE {
		stage.StageBranch(_relation_group_type)
	}
	for _, _spec_object_type := range a_spec_types.SPEC_OBJECT_TYPE {
		stage.StageBranch(_spec_object_type)
	}
	for _, _spec_relation_type := range a_spec_types.SPEC_RELATION_TYPE {
		stage.StageBranch(_spec_relation_type)
	}
	for _, _specification_type := range a_spec_types.SPECIFICATION_TYPE {
		stage.StageBranch(_specification_type)
	}

}

func (a_the_header *A_THE_HEADER) GongStageBranch(stage *Stage) {
	stage.StageBranchA_THE_HEADER(a_the_header)
}

func (stage *Stage) StageBranchA_THE_HEADER(a_the_header *A_THE_HEADER) {

	// check if instance is already staged
	if stage.IsStaged(a_the_header) {
		return
	}

	a_the_header.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_the_header.REQ_IF_HEADER != nil {
		stage.StageBranch(a_the_header.REQ_IF_HEADER)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongStageBranch(stage *Stage) {
	stage.StageBranchA_TOOL_EXTENSIONS(a_tool_extensions)
}

func (stage *Stage) StageBranchA_TOOL_EXTENSIONS(a_tool_extensions *A_TOOL_EXTENSIONS) {

	// check if instance is already staged
	if stage.IsStaged(a_tool_extensions) {
		return
	}

	a_tool_extensions.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _req_if_tool_extension := range a_tool_extensions.REQ_IF_TOOL_EXTENSION {
		stage.StageBranch(_req_if_tool_extension)
	}

}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongStageBranch(stage *Stage) {
	stage.StageBranchDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean)
}

func (stage *Stage) StageBranchDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) {

	// check if instance is already staged
	if stage.IsStaged(datatype_definition_boolean) {
		return
	}

	datatype_definition_boolean.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_boolean.ALTERNATIVE_ID != nil {
		stage.StageBranch(datatype_definition_boolean.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongStageBranch(stage *Stage) {
	stage.StageBranchDATATYPE_DEFINITION_DATE(datatype_definition_date)
}

func (stage *Stage) StageBranchDATATYPE_DEFINITION_DATE(datatype_definition_date *DATATYPE_DEFINITION_DATE) {

	// check if instance is already staged
	if stage.IsStaged(datatype_definition_date) {
		return
	}

	datatype_definition_date.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_date.ALTERNATIVE_ID != nil {
		stage.StageBranch(datatype_definition_date.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongStageBranch(stage *Stage) {
	stage.StageBranchDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration)
}

func (stage *Stage) StageBranchDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) {

	// check if instance is already staged
	if stage.IsStaged(datatype_definition_enumeration) {
		return
	}

	datatype_definition_enumeration.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_enumeration.ALTERNATIVE_ID != nil {
		stage.StageBranch(datatype_definition_enumeration.ALTERNATIVE_ID)
	}
	if datatype_definition_enumeration.SPECIFIED_VALUES != nil {
		stage.StageBranch(datatype_definition_enumeration.SPECIFIED_VALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongStageBranch(stage *Stage) {
	stage.StageBranchDATATYPE_DEFINITION_INTEGER(datatype_definition_integer)
}

func (stage *Stage) StageBranchDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) {

	// check if instance is already staged
	if stage.IsStaged(datatype_definition_integer) {
		return
	}

	datatype_definition_integer.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_integer.ALTERNATIVE_ID != nil {
		stage.StageBranch(datatype_definition_integer.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongStageBranch(stage *Stage) {
	stage.StageBranchDATATYPE_DEFINITION_REAL(datatype_definition_real)
}

func (stage *Stage) StageBranchDATATYPE_DEFINITION_REAL(datatype_definition_real *DATATYPE_DEFINITION_REAL) {

	// check if instance is already staged
	if stage.IsStaged(datatype_definition_real) {
		return
	}

	datatype_definition_real.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_real.ALTERNATIVE_ID != nil {
		stage.StageBranch(datatype_definition_real.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongStageBranch(stage *Stage) {
	stage.StageBranchDATATYPE_DEFINITION_STRING(datatype_definition_string)
}

func (stage *Stage) StageBranchDATATYPE_DEFINITION_STRING(datatype_definition_string *DATATYPE_DEFINITION_STRING) {

	// check if instance is already staged
	if stage.IsStaged(datatype_definition_string) {
		return
	}

	datatype_definition_string.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_string.ALTERNATIVE_ID != nil {
		stage.StageBranch(datatype_definition_string.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongStageBranch(stage *Stage) {
	stage.StageBranchDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml)
}

func (stage *Stage) StageBranchDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) {

	// check if instance is already staged
	if stage.IsStaged(datatype_definition_xhtml) {
		return
	}

	datatype_definition_xhtml.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_xhtml.ALTERNATIVE_ID != nil {
		stage.StageBranch(datatype_definition_xhtml.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (embedded_value *EMBEDDED_VALUE) GongStageBranch(stage *Stage) {
	stage.StageBranchEMBEDDED_VALUE(embedded_value)
}

func (stage *Stage) StageBranchEMBEDDED_VALUE(embedded_value *EMBEDDED_VALUE) {

	// check if instance is already staged
	if stage.IsStaged(embedded_value) {
		return
	}

	embedded_value.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (enum_value *ENUM_VALUE) GongStageBranch(stage *Stage) {
	stage.StageBranchENUM_VALUE(enum_value)
}

func (stage *Stage) StageBranchENUM_VALUE(enum_value *ENUM_VALUE) {

	// check if instance is already staged
	if stage.IsStaged(enum_value) {
		return
	}

	enum_value.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if enum_value.ALTERNATIVE_ID != nil {
		stage.StageBranch(enum_value.ALTERNATIVE_ID)
	}
	if enum_value.PROPERTIES != nil {
		stage.StageBranch(enum_value.PROPERTIES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (relation_group *RELATION_GROUP) GongStageBranch(stage *Stage) {
	stage.StageBranchRELATION_GROUP(relation_group)
}

func (stage *Stage) StageBranchRELATION_GROUP(relation_group *RELATION_GROUP) {

	// check if instance is already staged
	if stage.IsStaged(relation_group) {
		return
	}

	relation_group.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if relation_group.ALTERNATIVE_ID != nil {
		stage.StageBranch(relation_group.ALTERNATIVE_ID)
	}
	if relation_group.SOURCE_SPECIFICATION != nil {
		stage.StageBranch(relation_group.SOURCE_SPECIFICATION)
	}
	if relation_group.SPEC_RELATIONS != nil {
		stage.StageBranch(relation_group.SPEC_RELATIONS)
	}
	if relation_group.TARGET_SPECIFICATION != nil {
		stage.StageBranch(relation_group.TARGET_SPECIFICATION)
	}
	if relation_group.TYPE != nil {
		stage.StageBranch(relation_group.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (relation_group_type *RELATION_GROUP_TYPE) GongStageBranch(stage *Stage) {
	stage.StageBranchRELATION_GROUP_TYPE(relation_group_type)
}

func (stage *Stage) StageBranchRELATION_GROUP_TYPE(relation_group_type *RELATION_GROUP_TYPE) {

	// check if instance is already staged
	if stage.IsStaged(relation_group_type) {
		return
	}

	relation_group_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if relation_group_type.ALTERNATIVE_ID != nil {
		stage.StageBranch(relation_group_type.ALTERNATIVE_ID)
	}
	if relation_group_type.SPEC_ATTRIBUTES != nil {
		stage.StageBranch(relation_group_type.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (req_if *REQ_IF) GongStageBranch(stage *Stage) {
	stage.StageBranchREQ_IF(req_if)
}

func (stage *Stage) StageBranchREQ_IF(req_if *REQ_IF) {

	// check if instance is already staged
	if stage.IsStaged(req_if) {
		return
	}

	req_if.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if req_if.THE_HEADER != nil {
		stage.StageBranch(req_if.THE_HEADER)
	}
	if req_if.CORE_CONTENT != nil {
		stage.StageBranch(req_if.CORE_CONTENT)
	}
	if req_if.TOOL_EXTENSIONS != nil {
		stage.StageBranch(req_if.TOOL_EXTENSIONS)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (req_if_content *REQ_IF_CONTENT) GongStageBranch(stage *Stage) {
	stage.StageBranchREQ_IF_CONTENT(req_if_content)
}

func (stage *Stage) StageBranchREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT) {

	// check if instance is already staged
	if stage.IsStaged(req_if_content) {
		return
	}

	req_if_content.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if req_if_content.DATATYPES != nil {
		stage.StageBranch(req_if_content.DATATYPES)
	}
	if req_if_content.SPEC_TYPES != nil {
		stage.StageBranch(req_if_content.SPEC_TYPES)
	}
	if req_if_content.SPEC_OBJECTS != nil {
		stage.StageBranch(req_if_content.SPEC_OBJECTS)
	}
	if req_if_content.SPEC_RELATIONS != nil {
		stage.StageBranch(req_if_content.SPEC_RELATIONS)
	}
	if req_if_content.SPECIFICATIONS != nil {
		stage.StageBranch(req_if_content.SPECIFICATIONS)
	}
	if req_if_content.SPEC_RELATION_GROUPS != nil {
		stage.StageBranch(req_if_content.SPEC_RELATION_GROUPS)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (req_if_header *REQ_IF_HEADER) GongStageBranch(stage *Stage) {
	stage.StageBranchREQ_IF_HEADER(req_if_header)
}

func (stage *Stage) StageBranchREQ_IF_HEADER(req_if_header *REQ_IF_HEADER) {

	// check if instance is already staged
	if stage.IsStaged(req_if_header) {
		return
	}

	req_if_header.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongStageBranch(stage *Stage) {
	stage.StageBranchREQ_IF_TOOL_EXTENSION(req_if_tool_extension)
}

func (stage *Stage) StageBranchREQ_IF_TOOL_EXTENSION(req_if_tool_extension *REQ_IF_TOOL_EXTENSION) {

	// check if instance is already staged
	if stage.IsStaged(req_if_tool_extension) {
		return
	}

	req_if_tool_extension.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (specification *SPECIFICATION) GongStageBranch(stage *Stage) {
	stage.StageBranchSPECIFICATION(specification)
}

func (stage *Stage) StageBranchSPECIFICATION(specification *SPECIFICATION) {

	// check if instance is already staged
	if stage.IsStaged(specification) {
		return
	}

	specification.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specification.ALTERNATIVE_ID != nil {
		stage.StageBranch(specification.ALTERNATIVE_ID)
	}
	if specification.CHILDREN != nil {
		stage.StageBranch(specification.CHILDREN)
	}
	if specification.VALUES != nil {
		stage.StageBranch(specification.VALUES)
	}
	if specification.TYPE != nil {
		stage.StageBranch(specification.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (specification_type *SPECIFICATION_TYPE) GongStageBranch(stage *Stage) {
	stage.StageBranchSPECIFICATION_TYPE(specification_type)
}

func (stage *Stage) StageBranchSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE) {

	// check if instance is already staged
	if stage.IsStaged(specification_type) {
		return
	}

	specification_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specification_type.ALTERNATIVE_ID != nil {
		stage.StageBranch(specification_type.ALTERNATIVE_ID)
	}
	if specification_type.SPEC_ATTRIBUTES != nil {
		stage.StageBranch(specification_type.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (spec_hierarchy *SPEC_HIERARCHY) GongStageBranch(stage *Stage) {
	stage.StageBranchSPEC_HIERARCHY(spec_hierarchy)
}

func (stage *Stage) StageBranchSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY) {

	// check if instance is already staged
	if stage.IsStaged(spec_hierarchy) {
		return
	}

	spec_hierarchy.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_hierarchy.ALTERNATIVE_ID != nil {
		stage.StageBranch(spec_hierarchy.ALTERNATIVE_ID)
	}
	if spec_hierarchy.CHILDREN != nil {
		stage.StageBranch(spec_hierarchy.CHILDREN)
	}
	if spec_hierarchy.EDITABLE_ATTS != nil {
		stage.StageBranch(spec_hierarchy.EDITABLE_ATTS)
	}
	if spec_hierarchy.OBJECT != nil {
		stage.StageBranch(spec_hierarchy.OBJECT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (spec_object *SPEC_OBJECT) GongStageBranch(stage *Stage) {
	stage.StageBranchSPEC_OBJECT(spec_object)
}

func (stage *Stage) StageBranchSPEC_OBJECT(spec_object *SPEC_OBJECT) {

	// check if instance is already staged
	if stage.IsStaged(spec_object) {
		return
	}

	spec_object.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_object.ALTERNATIVE_ID != nil {
		stage.StageBranch(spec_object.ALTERNATIVE_ID)
	}
	if spec_object.VALUES != nil {
		stage.StageBranch(spec_object.VALUES)
	}
	if spec_object.TYPE != nil {
		stage.StageBranch(spec_object.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (spec_object_type *SPEC_OBJECT_TYPE) GongStageBranch(stage *Stage) {
	stage.StageBranchSPEC_OBJECT_TYPE(spec_object_type)
}

func (stage *Stage) StageBranchSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE) {

	// check if instance is already staged
	if stage.IsStaged(spec_object_type) {
		return
	}

	spec_object_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_object_type.ALTERNATIVE_ID != nil {
		stage.StageBranch(spec_object_type.ALTERNATIVE_ID)
	}
	if spec_object_type.SPEC_ATTRIBUTES != nil {
		stage.StageBranch(spec_object_type.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (spec_relation *SPEC_RELATION) GongStageBranch(stage *Stage) {
	stage.StageBranchSPEC_RELATION(spec_relation)
}

func (stage *Stage) StageBranchSPEC_RELATION(spec_relation *SPEC_RELATION) {

	// check if instance is already staged
	if stage.IsStaged(spec_relation) {
		return
	}

	spec_relation.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_relation.ALTERNATIVE_ID != nil {
		stage.StageBranch(spec_relation.ALTERNATIVE_ID)
	}
	if spec_relation.VALUES != nil {
		stage.StageBranch(spec_relation.VALUES)
	}
	if spec_relation.SOURCE != nil {
		stage.StageBranch(spec_relation.SOURCE)
	}
	if spec_relation.TARGET != nil {
		stage.StageBranch(spec_relation.TARGET)
	}
	if spec_relation.TYPE != nil {
		stage.StageBranch(spec_relation.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (spec_relation_type *SPEC_RELATION_TYPE) GongStageBranch(stage *Stage) {
	stage.StageBranchSPEC_RELATION_TYPE(spec_relation_type)
}

func (stage *Stage) StageBranchSPEC_RELATION_TYPE(spec_relation_type *SPEC_RELATION_TYPE) {

	// check if instance is already staged
	if stage.IsStaged(spec_relation_type) {
		return
	}

	spec_relation_type.Stage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_relation_type.ALTERNATIVE_ID != nil {
		stage.StageBranch(spec_relation_type.ALTERNATIVE_ID)
	}
	if spec_relation_type.SPEC_ATTRIBUTES != nil {
		stage.StageBranch(spec_relation_type.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (xhtml_content *XHTML_CONTENT) GongStageBranch(stage *Stage) {
	stage.StageBranchXHTML_CONTENT(xhtml_content)
}

func (stage *Stage) StageBranchXHTML_CONTENT(xhtml_content *XHTML_CONTENT) {

	// check if instance is already staged
	if stage.IsStaged(xhtml_content) {
		return
	}

	xhtml_content.Stage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// GongCopyBranch stages instance and apply GongCopyBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the instance
//
// the algorithm stops along the course of graph if a vertex is already staged
func GongCopyBranch[Type Gongstruct](from *Type) (to *Type) {

	mapOrigCopy := make(map[any]any)
	_ = mapOrigCopy

	switch fromT := any(from).(type) {
	// insertion point for stage branch
	case *ALTERNATIVE_ID:
		toT := GongCopyBranchALTERNATIVE_ID(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		toT := GongCopyBranchATTRIBUTE_DEFINITION_BOOLEAN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_DATE:
		toT := GongCopyBranchATTRIBUTE_DEFINITION_DATE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		toT := GongCopyBranchATTRIBUTE_DEFINITION_ENUMERATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_INTEGER:
		toT := GongCopyBranchATTRIBUTE_DEFINITION_INTEGER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_REAL:
		toT := GongCopyBranchATTRIBUTE_DEFINITION_REAL(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_STRING:
		toT := GongCopyBranchATTRIBUTE_DEFINITION_STRING(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_DEFINITION_XHTML:
		toT := GongCopyBranchATTRIBUTE_DEFINITION_XHTML(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_BOOLEAN:
		toT := GongCopyBranchATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_DATE:
		toT := GongCopyBranchATTRIBUTE_VALUE_DATE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_ENUMERATION:
		toT := GongCopyBranchATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_INTEGER:
		toT := GongCopyBranchATTRIBUTE_VALUE_INTEGER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_REAL:
		toT := GongCopyBranchATTRIBUTE_VALUE_REAL(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_STRING:
		toT := GongCopyBranchATTRIBUTE_VALUE_STRING(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ATTRIBUTE_VALUE_XHTML:
		toT := GongCopyBranchATTRIBUTE_VALUE_XHTML(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ALTERNATIVE_ID:
		toT := GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		toT := GongCopyBranchA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		toT := GongCopyBranchA_ATTRIBUTE_DEFINITION_DATE_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		toT := GongCopyBranchA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		toT := GongCopyBranchA_ATTRIBUTE_DEFINITION_INTEGER_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		toT := GongCopyBranchA_ATTRIBUTE_DEFINITION_REAL_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		toT := GongCopyBranchA_ATTRIBUTE_DEFINITION_STRING_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		toT := GongCopyBranchA_ATTRIBUTE_DEFINITION_XHTML_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		toT := GongCopyBranchA_ATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_VALUE_DATE:
		toT := GongCopyBranchA_ATTRIBUTE_VALUE_DATE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		toT := GongCopyBranchA_ATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_VALUE_INTEGER:
		toT := GongCopyBranchA_ATTRIBUTE_VALUE_INTEGER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_VALUE_REAL:
		toT := GongCopyBranchA_ATTRIBUTE_VALUE_REAL(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_VALUE_STRING:
		toT := GongCopyBranchA_ATTRIBUTE_VALUE_STRING(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_VALUE_XHTML:
		toT := GongCopyBranchA_ATTRIBUTE_VALUE_XHTML(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ATTRIBUTE_VALUE_XHTML_1:
		toT := GongCopyBranchA_ATTRIBUTE_VALUE_XHTML_1(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_CHILDREN:
		toT := GongCopyBranchA_CHILDREN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_CORE_CONTENT:
		toT := GongCopyBranchA_CORE_CONTENT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPES:
		toT := GongCopyBranchA_DATATYPES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		toT := GongCopyBranchA_DATATYPE_DEFINITION_BOOLEAN_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPE_DEFINITION_DATE_REF:
		toT := GongCopyBranchA_DATATYPE_DEFINITION_DATE_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		toT := GongCopyBranchA_DATATYPE_DEFINITION_ENUMERATION_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		toT := GongCopyBranchA_DATATYPE_DEFINITION_INTEGER_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPE_DEFINITION_REAL_REF:
		toT := GongCopyBranchA_DATATYPE_DEFINITION_REAL_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPE_DEFINITION_STRING_REF:
		toT := GongCopyBranchA_DATATYPE_DEFINITION_STRING_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_DATATYPE_DEFINITION_XHTML_REF:
		toT := GongCopyBranchA_DATATYPE_DEFINITION_XHTML_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_EDITABLE_ATTS:
		toT := GongCopyBranchA_EDITABLE_ATTS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_ENUM_VALUE_REF:
		toT := GongCopyBranchA_ENUM_VALUE_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_OBJECT:
		toT := GongCopyBranchA_OBJECT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_PROPERTIES:
		toT := GongCopyBranchA_PROPERTIES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_RELATION_GROUP_TYPE_REF:
		toT := GongCopyBranchA_RELATION_GROUP_TYPE_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SOURCE_1:
		toT := GongCopyBranchA_SOURCE_1(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SOURCE_SPECIFICATION_1:
		toT := GongCopyBranchA_SOURCE_SPECIFICATION_1(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPECIFICATIONS:
		toT := GongCopyBranchA_SPECIFICATIONS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPECIFICATION_TYPE_REF:
		toT := GongCopyBranchA_SPECIFICATION_TYPE_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPECIFIED_VALUES:
		toT := GongCopyBranchA_SPECIFIED_VALUES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_ATTRIBUTES:
		toT := GongCopyBranchA_SPEC_ATTRIBUTES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_OBJECTS:
		toT := GongCopyBranchA_SPEC_OBJECTS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_OBJECT_TYPE_REF:
		toT := GongCopyBranchA_SPEC_OBJECT_TYPE_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_RELATIONS:
		toT := GongCopyBranchA_SPEC_RELATIONS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_RELATION_GROUPS:
		toT := GongCopyBranchA_SPEC_RELATION_GROUPS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_RELATION_REF:
		toT := GongCopyBranchA_SPEC_RELATION_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_RELATION_TYPE_REF:
		toT := GongCopyBranchA_SPEC_RELATION_TYPE_REF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_SPEC_TYPES:
		toT := GongCopyBranchA_SPEC_TYPES(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_THE_HEADER:
		toT := GongCopyBranchA_THE_HEADER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *A_TOOL_EXTENSIONS:
		toT := GongCopyBranchA_TOOL_EXTENSIONS(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_BOOLEAN:
		toT := GongCopyBranchDATATYPE_DEFINITION_BOOLEAN(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_DATE:
		toT := GongCopyBranchDATATYPE_DEFINITION_DATE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_ENUMERATION:
		toT := GongCopyBranchDATATYPE_DEFINITION_ENUMERATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_INTEGER:
		toT := GongCopyBranchDATATYPE_DEFINITION_INTEGER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_REAL:
		toT := GongCopyBranchDATATYPE_DEFINITION_REAL(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_STRING:
		toT := GongCopyBranchDATATYPE_DEFINITION_STRING(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *DATATYPE_DEFINITION_XHTML:
		toT := GongCopyBranchDATATYPE_DEFINITION_XHTML(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *EMBEDDED_VALUE:
		toT := GongCopyBranchEMBEDDED_VALUE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *ENUM_VALUE:
		toT := GongCopyBranchENUM_VALUE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RELATION_GROUP:
		toT := GongCopyBranchRELATION_GROUP(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *RELATION_GROUP_TYPE:
		toT := GongCopyBranchRELATION_GROUP_TYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQ_IF:
		toT := GongCopyBranchREQ_IF(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQ_IF_CONTENT:
		toT := GongCopyBranchREQ_IF_CONTENT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQ_IF_HEADER:
		toT := GongCopyBranchREQ_IF_HEADER(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *REQ_IF_TOOL_EXTENSION:
		toT := GongCopyBranchREQ_IF_TOOL_EXTENSION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECIFICATION:
		toT := GongCopyBranchSPECIFICATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPECIFICATION_TYPE:
		toT := GongCopyBranchSPECIFICATION_TYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPEC_HIERARCHY:
		toT := GongCopyBranchSPEC_HIERARCHY(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPEC_OBJECT:
		toT := GongCopyBranchSPEC_OBJECT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPEC_OBJECT_TYPE:
		toT := GongCopyBranchSPEC_OBJECT_TYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPEC_RELATION:
		toT := GongCopyBranchSPEC_RELATION(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *SPEC_RELATION_TYPE:
		toT := GongCopyBranchSPEC_RELATION_TYPE(mapOrigCopy, fromT)
		return any(toT).(*Type)

	case *XHTML_CONTENT:
		toT := GongCopyBranchXHTML_CONTENT(mapOrigCopy, fromT)
		return any(toT).(*Type)

	default:
		_ = fromT // to espace compilation issue when model is empty
	}
	return
}

// insertion point for stage branch per struct
func GongCopyBranchALTERNATIVE_ID(mapOrigCopy map[any]any, alternative_idFrom *ALTERNATIVE_ID) (alternative_idTo *ALTERNATIVE_ID) {

	// alternative_idFrom has already been copied
	if _alternative_idTo, ok := mapOrigCopy[alternative_idFrom]; ok {
		alternative_idTo = _alternative_idTo.(*ALTERNATIVE_ID)
		return
	}

	alternative_idTo = new(ALTERNATIVE_ID)
	mapOrigCopy[alternative_idFrom] = alternative_idTo
	alternative_idFrom.GongCopyBasicFields(alternative_idTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchATTRIBUTE_DEFINITION_BOOLEAN(mapOrigCopy map[any]any, attribute_definition_booleanFrom *ATTRIBUTE_DEFINITION_BOOLEAN) (attribute_definition_booleanTo *ATTRIBUTE_DEFINITION_BOOLEAN) {

	// attribute_definition_booleanFrom has already been copied
	if _attribute_definition_booleanTo, ok := mapOrigCopy[attribute_definition_booleanFrom]; ok {
		attribute_definition_booleanTo = _attribute_definition_booleanTo.(*ATTRIBUTE_DEFINITION_BOOLEAN)
		return
	}

	attribute_definition_booleanTo = new(ATTRIBUTE_DEFINITION_BOOLEAN)
	mapOrigCopy[attribute_definition_booleanFrom] = attribute_definition_booleanTo
	attribute_definition_booleanFrom.GongCopyBasicFields(attribute_definition_booleanTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_booleanFrom.ALTERNATIVE_ID != nil {
		attribute_definition_booleanTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, attribute_definition_booleanFrom.ALTERNATIVE_ID)
	}
	if attribute_definition_booleanFrom.DEFAULT_VALUE != nil {
		attribute_definition_booleanTo.DEFAULT_VALUE = GongCopyBranchA_ATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, attribute_definition_booleanFrom.DEFAULT_VALUE)
	}
	if attribute_definition_booleanFrom.TYPE != nil {
		attribute_definition_booleanTo.TYPE = GongCopyBranchA_DATATYPE_DEFINITION_BOOLEAN_REF(mapOrigCopy, attribute_definition_booleanFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchATTRIBUTE_DEFINITION_DATE(mapOrigCopy map[any]any, attribute_definition_dateFrom *ATTRIBUTE_DEFINITION_DATE) (attribute_definition_dateTo *ATTRIBUTE_DEFINITION_DATE) {

	// attribute_definition_dateFrom has already been copied
	if _attribute_definition_dateTo, ok := mapOrigCopy[attribute_definition_dateFrom]; ok {
		attribute_definition_dateTo = _attribute_definition_dateTo.(*ATTRIBUTE_DEFINITION_DATE)
		return
	}

	attribute_definition_dateTo = new(ATTRIBUTE_DEFINITION_DATE)
	mapOrigCopy[attribute_definition_dateFrom] = attribute_definition_dateTo
	attribute_definition_dateFrom.GongCopyBasicFields(attribute_definition_dateTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_dateFrom.ALTERNATIVE_ID != nil {
		attribute_definition_dateTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, attribute_definition_dateFrom.ALTERNATIVE_ID)
	}
	if attribute_definition_dateFrom.DEFAULT_VALUE != nil {
		attribute_definition_dateTo.DEFAULT_VALUE = GongCopyBranchA_ATTRIBUTE_VALUE_DATE(mapOrigCopy, attribute_definition_dateFrom.DEFAULT_VALUE)
	}
	if attribute_definition_dateFrom.TYPE != nil {
		attribute_definition_dateTo.TYPE = GongCopyBranchA_DATATYPE_DEFINITION_DATE_REF(mapOrigCopy, attribute_definition_dateFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchATTRIBUTE_DEFINITION_ENUMERATION(mapOrigCopy map[any]any, attribute_definition_enumerationFrom *ATTRIBUTE_DEFINITION_ENUMERATION) (attribute_definition_enumerationTo *ATTRIBUTE_DEFINITION_ENUMERATION) {

	// attribute_definition_enumerationFrom has already been copied
	if _attribute_definition_enumerationTo, ok := mapOrigCopy[attribute_definition_enumerationFrom]; ok {
		attribute_definition_enumerationTo = _attribute_definition_enumerationTo.(*ATTRIBUTE_DEFINITION_ENUMERATION)
		return
	}

	attribute_definition_enumerationTo = new(ATTRIBUTE_DEFINITION_ENUMERATION)
	mapOrigCopy[attribute_definition_enumerationFrom] = attribute_definition_enumerationTo
	attribute_definition_enumerationFrom.GongCopyBasicFields(attribute_definition_enumerationTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_enumerationFrom.ALTERNATIVE_ID != nil {
		attribute_definition_enumerationTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, attribute_definition_enumerationFrom.ALTERNATIVE_ID)
	}
	if attribute_definition_enumerationFrom.DEFAULT_VALUE != nil {
		attribute_definition_enumerationTo.DEFAULT_VALUE = GongCopyBranchA_ATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, attribute_definition_enumerationFrom.DEFAULT_VALUE)
	}
	if attribute_definition_enumerationFrom.TYPE != nil {
		attribute_definition_enumerationTo.TYPE = GongCopyBranchA_DATATYPE_DEFINITION_ENUMERATION_REF(mapOrigCopy, attribute_definition_enumerationFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchATTRIBUTE_DEFINITION_INTEGER(mapOrigCopy map[any]any, attribute_definition_integerFrom *ATTRIBUTE_DEFINITION_INTEGER) (attribute_definition_integerTo *ATTRIBUTE_DEFINITION_INTEGER) {

	// attribute_definition_integerFrom has already been copied
	if _attribute_definition_integerTo, ok := mapOrigCopy[attribute_definition_integerFrom]; ok {
		attribute_definition_integerTo = _attribute_definition_integerTo.(*ATTRIBUTE_DEFINITION_INTEGER)
		return
	}

	attribute_definition_integerTo = new(ATTRIBUTE_DEFINITION_INTEGER)
	mapOrigCopy[attribute_definition_integerFrom] = attribute_definition_integerTo
	attribute_definition_integerFrom.GongCopyBasicFields(attribute_definition_integerTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_integerFrom.ALTERNATIVE_ID != nil {
		attribute_definition_integerTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, attribute_definition_integerFrom.ALTERNATIVE_ID)
	}
	if attribute_definition_integerFrom.DEFAULT_VALUE != nil {
		attribute_definition_integerTo.DEFAULT_VALUE = GongCopyBranchA_ATTRIBUTE_VALUE_INTEGER(mapOrigCopy, attribute_definition_integerFrom.DEFAULT_VALUE)
	}
	if attribute_definition_integerFrom.TYPE != nil {
		attribute_definition_integerTo.TYPE = GongCopyBranchA_DATATYPE_DEFINITION_INTEGER_REF(mapOrigCopy, attribute_definition_integerFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchATTRIBUTE_DEFINITION_REAL(mapOrigCopy map[any]any, attribute_definition_realFrom *ATTRIBUTE_DEFINITION_REAL) (attribute_definition_realTo *ATTRIBUTE_DEFINITION_REAL) {

	// attribute_definition_realFrom has already been copied
	if _attribute_definition_realTo, ok := mapOrigCopy[attribute_definition_realFrom]; ok {
		attribute_definition_realTo = _attribute_definition_realTo.(*ATTRIBUTE_DEFINITION_REAL)
		return
	}

	attribute_definition_realTo = new(ATTRIBUTE_DEFINITION_REAL)
	mapOrigCopy[attribute_definition_realFrom] = attribute_definition_realTo
	attribute_definition_realFrom.GongCopyBasicFields(attribute_definition_realTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_realFrom.ALTERNATIVE_ID != nil {
		attribute_definition_realTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, attribute_definition_realFrom.ALTERNATIVE_ID)
	}
	if attribute_definition_realFrom.DEFAULT_VALUE != nil {
		attribute_definition_realTo.DEFAULT_VALUE = GongCopyBranchA_ATTRIBUTE_VALUE_REAL(mapOrigCopy, attribute_definition_realFrom.DEFAULT_VALUE)
	}
	if attribute_definition_realFrom.TYPE != nil {
		attribute_definition_realTo.TYPE = GongCopyBranchA_DATATYPE_DEFINITION_REAL_REF(mapOrigCopy, attribute_definition_realFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchATTRIBUTE_DEFINITION_STRING(mapOrigCopy map[any]any, attribute_definition_stringFrom *ATTRIBUTE_DEFINITION_STRING) (attribute_definition_stringTo *ATTRIBUTE_DEFINITION_STRING) {

	// attribute_definition_stringFrom has already been copied
	if _attribute_definition_stringTo, ok := mapOrigCopy[attribute_definition_stringFrom]; ok {
		attribute_definition_stringTo = _attribute_definition_stringTo.(*ATTRIBUTE_DEFINITION_STRING)
		return
	}

	attribute_definition_stringTo = new(ATTRIBUTE_DEFINITION_STRING)
	mapOrigCopy[attribute_definition_stringFrom] = attribute_definition_stringTo
	attribute_definition_stringFrom.GongCopyBasicFields(attribute_definition_stringTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_stringFrom.ALTERNATIVE_ID != nil {
		attribute_definition_stringTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, attribute_definition_stringFrom.ALTERNATIVE_ID)
	}
	if attribute_definition_stringFrom.DEFAULT_VALUE != nil {
		attribute_definition_stringTo.DEFAULT_VALUE = GongCopyBranchA_ATTRIBUTE_VALUE_STRING(mapOrigCopy, attribute_definition_stringFrom.DEFAULT_VALUE)
	}
	if attribute_definition_stringFrom.TYPE != nil {
		attribute_definition_stringTo.TYPE = GongCopyBranchA_DATATYPE_DEFINITION_STRING_REF(mapOrigCopy, attribute_definition_stringFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchATTRIBUTE_DEFINITION_XHTML(mapOrigCopy map[any]any, attribute_definition_xhtmlFrom *ATTRIBUTE_DEFINITION_XHTML) (attribute_definition_xhtmlTo *ATTRIBUTE_DEFINITION_XHTML) {

	// attribute_definition_xhtmlFrom has already been copied
	if _attribute_definition_xhtmlTo, ok := mapOrigCopy[attribute_definition_xhtmlFrom]; ok {
		attribute_definition_xhtmlTo = _attribute_definition_xhtmlTo.(*ATTRIBUTE_DEFINITION_XHTML)
		return
	}

	attribute_definition_xhtmlTo = new(ATTRIBUTE_DEFINITION_XHTML)
	mapOrigCopy[attribute_definition_xhtmlFrom] = attribute_definition_xhtmlTo
	attribute_definition_xhtmlFrom.GongCopyBasicFields(attribute_definition_xhtmlTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_xhtmlFrom.ALTERNATIVE_ID != nil {
		attribute_definition_xhtmlTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, attribute_definition_xhtmlFrom.ALTERNATIVE_ID)
	}
	if attribute_definition_xhtmlFrom.DEFAULT_VALUE != nil {
		attribute_definition_xhtmlTo.DEFAULT_VALUE = GongCopyBranchA_ATTRIBUTE_VALUE_XHTML(mapOrigCopy, attribute_definition_xhtmlFrom.DEFAULT_VALUE)
	}
	if attribute_definition_xhtmlFrom.TYPE != nil {
		attribute_definition_xhtmlTo.TYPE = GongCopyBranchA_DATATYPE_DEFINITION_XHTML_REF(mapOrigCopy, attribute_definition_xhtmlFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy map[any]any, attribute_value_booleanFrom *ATTRIBUTE_VALUE_BOOLEAN) (attribute_value_booleanTo *ATTRIBUTE_VALUE_BOOLEAN) {

	// attribute_value_booleanFrom has already been copied
	if _attribute_value_booleanTo, ok := mapOrigCopy[attribute_value_booleanFrom]; ok {
		attribute_value_booleanTo = _attribute_value_booleanTo.(*ATTRIBUTE_VALUE_BOOLEAN)
		return
	}

	attribute_value_booleanTo = new(ATTRIBUTE_VALUE_BOOLEAN)
	mapOrigCopy[attribute_value_booleanFrom] = attribute_value_booleanTo
	attribute_value_booleanFrom.GongCopyBasicFields(attribute_value_booleanTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_booleanFrom.DEFINITION != nil {
		attribute_value_booleanTo.DEFINITION = GongCopyBranchA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(mapOrigCopy, attribute_value_booleanFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchATTRIBUTE_VALUE_DATE(mapOrigCopy map[any]any, attribute_value_dateFrom *ATTRIBUTE_VALUE_DATE) (attribute_value_dateTo *ATTRIBUTE_VALUE_DATE) {

	// attribute_value_dateFrom has already been copied
	if _attribute_value_dateTo, ok := mapOrigCopy[attribute_value_dateFrom]; ok {
		attribute_value_dateTo = _attribute_value_dateTo.(*ATTRIBUTE_VALUE_DATE)
		return
	}

	attribute_value_dateTo = new(ATTRIBUTE_VALUE_DATE)
	mapOrigCopy[attribute_value_dateFrom] = attribute_value_dateTo
	attribute_value_dateFrom.GongCopyBasicFields(attribute_value_dateTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_dateFrom.DEFINITION != nil {
		attribute_value_dateTo.DEFINITION = GongCopyBranchA_ATTRIBUTE_DEFINITION_DATE_REF(mapOrigCopy, attribute_value_dateFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy map[any]any, attribute_value_enumerationFrom *ATTRIBUTE_VALUE_ENUMERATION) (attribute_value_enumerationTo *ATTRIBUTE_VALUE_ENUMERATION) {

	// attribute_value_enumerationFrom has already been copied
	if _attribute_value_enumerationTo, ok := mapOrigCopy[attribute_value_enumerationFrom]; ok {
		attribute_value_enumerationTo = _attribute_value_enumerationTo.(*ATTRIBUTE_VALUE_ENUMERATION)
		return
	}

	attribute_value_enumerationTo = new(ATTRIBUTE_VALUE_ENUMERATION)
	mapOrigCopy[attribute_value_enumerationFrom] = attribute_value_enumerationTo
	attribute_value_enumerationFrom.GongCopyBasicFields(attribute_value_enumerationTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_enumerationFrom.DEFINITION != nil {
		attribute_value_enumerationTo.DEFINITION = GongCopyBranchA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(mapOrigCopy, attribute_value_enumerationFrom.DEFINITION)
	}
	if attribute_value_enumerationFrom.VALUES != nil {
		attribute_value_enumerationTo.VALUES = GongCopyBranchA_ENUM_VALUE_REF(mapOrigCopy, attribute_value_enumerationFrom.VALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchATTRIBUTE_VALUE_INTEGER(mapOrigCopy map[any]any, attribute_value_integerFrom *ATTRIBUTE_VALUE_INTEGER) (attribute_value_integerTo *ATTRIBUTE_VALUE_INTEGER) {

	// attribute_value_integerFrom has already been copied
	if _attribute_value_integerTo, ok := mapOrigCopy[attribute_value_integerFrom]; ok {
		attribute_value_integerTo = _attribute_value_integerTo.(*ATTRIBUTE_VALUE_INTEGER)
		return
	}

	attribute_value_integerTo = new(ATTRIBUTE_VALUE_INTEGER)
	mapOrigCopy[attribute_value_integerFrom] = attribute_value_integerTo
	attribute_value_integerFrom.GongCopyBasicFields(attribute_value_integerTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_integerFrom.DEFINITION != nil {
		attribute_value_integerTo.DEFINITION = GongCopyBranchA_ATTRIBUTE_DEFINITION_INTEGER_REF(mapOrigCopy, attribute_value_integerFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchATTRIBUTE_VALUE_REAL(mapOrigCopy map[any]any, attribute_value_realFrom *ATTRIBUTE_VALUE_REAL) (attribute_value_realTo *ATTRIBUTE_VALUE_REAL) {

	// attribute_value_realFrom has already been copied
	if _attribute_value_realTo, ok := mapOrigCopy[attribute_value_realFrom]; ok {
		attribute_value_realTo = _attribute_value_realTo.(*ATTRIBUTE_VALUE_REAL)
		return
	}

	attribute_value_realTo = new(ATTRIBUTE_VALUE_REAL)
	mapOrigCopy[attribute_value_realFrom] = attribute_value_realTo
	attribute_value_realFrom.GongCopyBasicFields(attribute_value_realTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_realFrom.DEFINITION != nil {
		attribute_value_realTo.DEFINITION = GongCopyBranchA_ATTRIBUTE_DEFINITION_REAL_REF(mapOrigCopy, attribute_value_realFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchATTRIBUTE_VALUE_STRING(mapOrigCopy map[any]any, attribute_value_stringFrom *ATTRIBUTE_VALUE_STRING) (attribute_value_stringTo *ATTRIBUTE_VALUE_STRING) {

	// attribute_value_stringFrom has already been copied
	if _attribute_value_stringTo, ok := mapOrigCopy[attribute_value_stringFrom]; ok {
		attribute_value_stringTo = _attribute_value_stringTo.(*ATTRIBUTE_VALUE_STRING)
		return
	}

	attribute_value_stringTo = new(ATTRIBUTE_VALUE_STRING)
	mapOrigCopy[attribute_value_stringFrom] = attribute_value_stringTo
	attribute_value_stringFrom.GongCopyBasicFields(attribute_value_stringTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_stringFrom.DEFINITION != nil {
		attribute_value_stringTo.DEFINITION = GongCopyBranchA_ATTRIBUTE_DEFINITION_STRING_REF(mapOrigCopy, attribute_value_stringFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchATTRIBUTE_VALUE_XHTML(mapOrigCopy map[any]any, attribute_value_xhtmlFrom *ATTRIBUTE_VALUE_XHTML) (attribute_value_xhtmlTo *ATTRIBUTE_VALUE_XHTML) {

	// attribute_value_xhtmlFrom has already been copied
	if _attribute_value_xhtmlTo, ok := mapOrigCopy[attribute_value_xhtmlFrom]; ok {
		attribute_value_xhtmlTo = _attribute_value_xhtmlTo.(*ATTRIBUTE_VALUE_XHTML)
		return
	}

	attribute_value_xhtmlTo = new(ATTRIBUTE_VALUE_XHTML)
	mapOrigCopy[attribute_value_xhtmlFrom] = attribute_value_xhtmlTo
	attribute_value_xhtmlFrom.GongCopyBasicFields(attribute_value_xhtmlTo)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_xhtmlFrom.THE_VALUE != nil {
		attribute_value_xhtmlTo.THE_VALUE = GongCopyBranchXHTML_CONTENT(mapOrigCopy, attribute_value_xhtmlFrom.THE_VALUE)
	}
	if attribute_value_xhtmlFrom.THE_ORIGINAL_VALUE != nil {
		attribute_value_xhtmlTo.THE_ORIGINAL_VALUE = GongCopyBranchXHTML_CONTENT(mapOrigCopy, attribute_value_xhtmlFrom.THE_ORIGINAL_VALUE)
	}
	if attribute_value_xhtmlFrom.DEFINITION != nil {
		attribute_value_xhtmlTo.DEFINITION = GongCopyBranchA_ATTRIBUTE_DEFINITION_XHTML_REF(mapOrigCopy, attribute_value_xhtmlFrom.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy map[any]any, a_alternative_idFrom *A_ALTERNATIVE_ID) (a_alternative_idTo *A_ALTERNATIVE_ID) {

	// a_alternative_idFrom has already been copied
	if _a_alternative_idTo, ok := mapOrigCopy[a_alternative_idFrom]; ok {
		a_alternative_idTo = _a_alternative_idTo.(*A_ALTERNATIVE_ID)
		return
	}

	a_alternative_idTo = new(A_ALTERNATIVE_ID)
	mapOrigCopy[a_alternative_idFrom] = a_alternative_idTo
	a_alternative_idFrom.GongCopyBasicFields(a_alternative_idTo)

	//insertion point for the staging of instances referenced by pointers
	if a_alternative_idFrom.ALTERNATIVE_ID != nil {
		a_alternative_idTo.ALTERNATIVE_ID = GongCopyBranchALTERNATIVE_ID(mapOrigCopy, a_alternative_idFrom.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(mapOrigCopy map[any]any, a_attribute_definition_boolean_refFrom *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) (a_attribute_definition_boolean_refTo *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) {

	// a_attribute_definition_boolean_refFrom has already been copied
	if _a_attribute_definition_boolean_refTo, ok := mapOrigCopy[a_attribute_definition_boolean_refFrom]; ok {
		a_attribute_definition_boolean_refTo = _a_attribute_definition_boolean_refTo.(*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
		return
	}

	a_attribute_definition_boolean_refTo = new(A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
	mapOrigCopy[a_attribute_definition_boolean_refFrom] = a_attribute_definition_boolean_refTo
	a_attribute_definition_boolean_refFrom.GongCopyBasicFields(a_attribute_definition_boolean_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_ATTRIBUTE_DEFINITION_DATE_REF(mapOrigCopy map[any]any, a_attribute_definition_date_refFrom *A_ATTRIBUTE_DEFINITION_DATE_REF) (a_attribute_definition_date_refTo *A_ATTRIBUTE_DEFINITION_DATE_REF) {

	// a_attribute_definition_date_refFrom has already been copied
	if _a_attribute_definition_date_refTo, ok := mapOrigCopy[a_attribute_definition_date_refFrom]; ok {
		a_attribute_definition_date_refTo = _a_attribute_definition_date_refTo.(*A_ATTRIBUTE_DEFINITION_DATE_REF)
		return
	}

	a_attribute_definition_date_refTo = new(A_ATTRIBUTE_DEFINITION_DATE_REF)
	mapOrigCopy[a_attribute_definition_date_refFrom] = a_attribute_definition_date_refTo
	a_attribute_definition_date_refFrom.GongCopyBasicFields(a_attribute_definition_date_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(mapOrigCopy map[any]any, a_attribute_definition_enumeration_refFrom *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) (a_attribute_definition_enumeration_refTo *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) {

	// a_attribute_definition_enumeration_refFrom has already been copied
	if _a_attribute_definition_enumeration_refTo, ok := mapOrigCopy[a_attribute_definition_enumeration_refFrom]; ok {
		a_attribute_definition_enumeration_refTo = _a_attribute_definition_enumeration_refTo.(*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
		return
	}

	a_attribute_definition_enumeration_refTo = new(A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
	mapOrigCopy[a_attribute_definition_enumeration_refFrom] = a_attribute_definition_enumeration_refTo
	a_attribute_definition_enumeration_refFrom.GongCopyBasicFields(a_attribute_definition_enumeration_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_ATTRIBUTE_DEFINITION_INTEGER_REF(mapOrigCopy map[any]any, a_attribute_definition_integer_refFrom *A_ATTRIBUTE_DEFINITION_INTEGER_REF) (a_attribute_definition_integer_refTo *A_ATTRIBUTE_DEFINITION_INTEGER_REF) {

	// a_attribute_definition_integer_refFrom has already been copied
	if _a_attribute_definition_integer_refTo, ok := mapOrigCopy[a_attribute_definition_integer_refFrom]; ok {
		a_attribute_definition_integer_refTo = _a_attribute_definition_integer_refTo.(*A_ATTRIBUTE_DEFINITION_INTEGER_REF)
		return
	}

	a_attribute_definition_integer_refTo = new(A_ATTRIBUTE_DEFINITION_INTEGER_REF)
	mapOrigCopy[a_attribute_definition_integer_refFrom] = a_attribute_definition_integer_refTo
	a_attribute_definition_integer_refFrom.GongCopyBasicFields(a_attribute_definition_integer_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_ATTRIBUTE_DEFINITION_REAL_REF(mapOrigCopy map[any]any, a_attribute_definition_real_refFrom *A_ATTRIBUTE_DEFINITION_REAL_REF) (a_attribute_definition_real_refTo *A_ATTRIBUTE_DEFINITION_REAL_REF) {

	// a_attribute_definition_real_refFrom has already been copied
	if _a_attribute_definition_real_refTo, ok := mapOrigCopy[a_attribute_definition_real_refFrom]; ok {
		a_attribute_definition_real_refTo = _a_attribute_definition_real_refTo.(*A_ATTRIBUTE_DEFINITION_REAL_REF)
		return
	}

	a_attribute_definition_real_refTo = new(A_ATTRIBUTE_DEFINITION_REAL_REF)
	mapOrigCopy[a_attribute_definition_real_refFrom] = a_attribute_definition_real_refTo
	a_attribute_definition_real_refFrom.GongCopyBasicFields(a_attribute_definition_real_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_ATTRIBUTE_DEFINITION_STRING_REF(mapOrigCopy map[any]any, a_attribute_definition_string_refFrom *A_ATTRIBUTE_DEFINITION_STRING_REF) (a_attribute_definition_string_refTo *A_ATTRIBUTE_DEFINITION_STRING_REF) {

	// a_attribute_definition_string_refFrom has already been copied
	if _a_attribute_definition_string_refTo, ok := mapOrigCopy[a_attribute_definition_string_refFrom]; ok {
		a_attribute_definition_string_refTo = _a_attribute_definition_string_refTo.(*A_ATTRIBUTE_DEFINITION_STRING_REF)
		return
	}

	a_attribute_definition_string_refTo = new(A_ATTRIBUTE_DEFINITION_STRING_REF)
	mapOrigCopy[a_attribute_definition_string_refFrom] = a_attribute_definition_string_refTo
	a_attribute_definition_string_refFrom.GongCopyBasicFields(a_attribute_definition_string_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_ATTRIBUTE_DEFINITION_XHTML_REF(mapOrigCopy map[any]any, a_attribute_definition_xhtml_refFrom *A_ATTRIBUTE_DEFINITION_XHTML_REF) (a_attribute_definition_xhtml_refTo *A_ATTRIBUTE_DEFINITION_XHTML_REF) {

	// a_attribute_definition_xhtml_refFrom has already been copied
	if _a_attribute_definition_xhtml_refTo, ok := mapOrigCopy[a_attribute_definition_xhtml_refFrom]; ok {
		a_attribute_definition_xhtml_refTo = _a_attribute_definition_xhtml_refTo.(*A_ATTRIBUTE_DEFINITION_XHTML_REF)
		return
	}

	a_attribute_definition_xhtml_refTo = new(A_ATTRIBUTE_DEFINITION_XHTML_REF)
	mapOrigCopy[a_attribute_definition_xhtml_refFrom] = a_attribute_definition_xhtml_refTo
	a_attribute_definition_xhtml_refFrom.GongCopyBasicFields(a_attribute_definition_xhtml_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_ATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy map[any]any, a_attribute_value_booleanFrom *A_ATTRIBUTE_VALUE_BOOLEAN) (a_attribute_value_booleanTo *A_ATTRIBUTE_VALUE_BOOLEAN) {

	// a_attribute_value_booleanFrom has already been copied
	if _a_attribute_value_booleanTo, ok := mapOrigCopy[a_attribute_value_booleanFrom]; ok {
		a_attribute_value_booleanTo = _a_attribute_value_booleanTo.(*A_ATTRIBUTE_VALUE_BOOLEAN)
		return
	}

	a_attribute_value_booleanTo = new(A_ATTRIBUTE_VALUE_BOOLEAN)
	mapOrigCopy[a_attribute_value_booleanFrom] = a_attribute_value_booleanTo
	a_attribute_value_booleanFrom.GongCopyBasicFields(a_attribute_value_booleanTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_boolean := range a_attribute_value_booleanFrom.ATTRIBUTE_VALUE_BOOLEAN {
		a_attribute_value_booleanTo.ATTRIBUTE_VALUE_BOOLEAN = append(a_attribute_value_booleanTo.ATTRIBUTE_VALUE_BOOLEAN, GongCopyBranchATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, _attribute_value_boolean))
	}

	return
}

func GongCopyBranchA_ATTRIBUTE_VALUE_DATE(mapOrigCopy map[any]any, a_attribute_value_dateFrom *A_ATTRIBUTE_VALUE_DATE) (a_attribute_value_dateTo *A_ATTRIBUTE_VALUE_DATE) {

	// a_attribute_value_dateFrom has already been copied
	if _a_attribute_value_dateTo, ok := mapOrigCopy[a_attribute_value_dateFrom]; ok {
		a_attribute_value_dateTo = _a_attribute_value_dateTo.(*A_ATTRIBUTE_VALUE_DATE)
		return
	}

	a_attribute_value_dateTo = new(A_ATTRIBUTE_VALUE_DATE)
	mapOrigCopy[a_attribute_value_dateFrom] = a_attribute_value_dateTo
	a_attribute_value_dateFrom.GongCopyBasicFields(a_attribute_value_dateTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_date := range a_attribute_value_dateFrom.ATTRIBUTE_VALUE_DATE {
		a_attribute_value_dateTo.ATTRIBUTE_VALUE_DATE = append(a_attribute_value_dateTo.ATTRIBUTE_VALUE_DATE, GongCopyBranchATTRIBUTE_VALUE_DATE(mapOrigCopy, _attribute_value_date))
	}

	return
}

func GongCopyBranchA_ATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy map[any]any, a_attribute_value_enumerationFrom *A_ATTRIBUTE_VALUE_ENUMERATION) (a_attribute_value_enumerationTo *A_ATTRIBUTE_VALUE_ENUMERATION) {

	// a_attribute_value_enumerationFrom has already been copied
	if _a_attribute_value_enumerationTo, ok := mapOrigCopy[a_attribute_value_enumerationFrom]; ok {
		a_attribute_value_enumerationTo = _a_attribute_value_enumerationTo.(*A_ATTRIBUTE_VALUE_ENUMERATION)
		return
	}

	a_attribute_value_enumerationTo = new(A_ATTRIBUTE_VALUE_ENUMERATION)
	mapOrigCopy[a_attribute_value_enumerationFrom] = a_attribute_value_enumerationTo
	a_attribute_value_enumerationFrom.GongCopyBasicFields(a_attribute_value_enumerationTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_enumeration := range a_attribute_value_enumerationFrom.ATTRIBUTE_VALUE_ENUMERATION {
		a_attribute_value_enumerationTo.ATTRIBUTE_VALUE_ENUMERATION = append(a_attribute_value_enumerationTo.ATTRIBUTE_VALUE_ENUMERATION, GongCopyBranchATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, _attribute_value_enumeration))
	}

	return
}

func GongCopyBranchA_ATTRIBUTE_VALUE_INTEGER(mapOrigCopy map[any]any, a_attribute_value_integerFrom *A_ATTRIBUTE_VALUE_INTEGER) (a_attribute_value_integerTo *A_ATTRIBUTE_VALUE_INTEGER) {

	// a_attribute_value_integerFrom has already been copied
	if _a_attribute_value_integerTo, ok := mapOrigCopy[a_attribute_value_integerFrom]; ok {
		a_attribute_value_integerTo = _a_attribute_value_integerTo.(*A_ATTRIBUTE_VALUE_INTEGER)
		return
	}

	a_attribute_value_integerTo = new(A_ATTRIBUTE_VALUE_INTEGER)
	mapOrigCopy[a_attribute_value_integerFrom] = a_attribute_value_integerTo
	a_attribute_value_integerFrom.GongCopyBasicFields(a_attribute_value_integerTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_integer := range a_attribute_value_integerFrom.ATTRIBUTE_VALUE_INTEGER {
		a_attribute_value_integerTo.ATTRIBUTE_VALUE_INTEGER = append(a_attribute_value_integerTo.ATTRIBUTE_VALUE_INTEGER, GongCopyBranchATTRIBUTE_VALUE_INTEGER(mapOrigCopy, _attribute_value_integer))
	}

	return
}

func GongCopyBranchA_ATTRIBUTE_VALUE_REAL(mapOrigCopy map[any]any, a_attribute_value_realFrom *A_ATTRIBUTE_VALUE_REAL) (a_attribute_value_realTo *A_ATTRIBUTE_VALUE_REAL) {

	// a_attribute_value_realFrom has already been copied
	if _a_attribute_value_realTo, ok := mapOrigCopy[a_attribute_value_realFrom]; ok {
		a_attribute_value_realTo = _a_attribute_value_realTo.(*A_ATTRIBUTE_VALUE_REAL)
		return
	}

	a_attribute_value_realTo = new(A_ATTRIBUTE_VALUE_REAL)
	mapOrigCopy[a_attribute_value_realFrom] = a_attribute_value_realTo
	a_attribute_value_realFrom.GongCopyBasicFields(a_attribute_value_realTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_real := range a_attribute_value_realFrom.ATTRIBUTE_VALUE_REAL {
		a_attribute_value_realTo.ATTRIBUTE_VALUE_REAL = append(a_attribute_value_realTo.ATTRIBUTE_VALUE_REAL, GongCopyBranchATTRIBUTE_VALUE_REAL(mapOrigCopy, _attribute_value_real))
	}

	return
}

func GongCopyBranchA_ATTRIBUTE_VALUE_STRING(mapOrigCopy map[any]any, a_attribute_value_stringFrom *A_ATTRIBUTE_VALUE_STRING) (a_attribute_value_stringTo *A_ATTRIBUTE_VALUE_STRING) {

	// a_attribute_value_stringFrom has already been copied
	if _a_attribute_value_stringTo, ok := mapOrigCopy[a_attribute_value_stringFrom]; ok {
		a_attribute_value_stringTo = _a_attribute_value_stringTo.(*A_ATTRIBUTE_VALUE_STRING)
		return
	}

	a_attribute_value_stringTo = new(A_ATTRIBUTE_VALUE_STRING)
	mapOrigCopy[a_attribute_value_stringFrom] = a_attribute_value_stringTo
	a_attribute_value_stringFrom.GongCopyBasicFields(a_attribute_value_stringTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_string := range a_attribute_value_stringFrom.ATTRIBUTE_VALUE_STRING {
		a_attribute_value_stringTo.ATTRIBUTE_VALUE_STRING = append(a_attribute_value_stringTo.ATTRIBUTE_VALUE_STRING, GongCopyBranchATTRIBUTE_VALUE_STRING(mapOrigCopy, _attribute_value_string))
	}

	return
}

func GongCopyBranchA_ATTRIBUTE_VALUE_XHTML(mapOrigCopy map[any]any, a_attribute_value_xhtmlFrom *A_ATTRIBUTE_VALUE_XHTML) (a_attribute_value_xhtmlTo *A_ATTRIBUTE_VALUE_XHTML) {

	// a_attribute_value_xhtmlFrom has already been copied
	if _a_attribute_value_xhtmlTo, ok := mapOrigCopy[a_attribute_value_xhtmlFrom]; ok {
		a_attribute_value_xhtmlTo = _a_attribute_value_xhtmlTo.(*A_ATTRIBUTE_VALUE_XHTML)
		return
	}

	a_attribute_value_xhtmlTo = new(A_ATTRIBUTE_VALUE_XHTML)
	mapOrigCopy[a_attribute_value_xhtmlFrom] = a_attribute_value_xhtmlTo
	a_attribute_value_xhtmlFrom.GongCopyBasicFields(a_attribute_value_xhtmlTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_xhtml := range a_attribute_value_xhtmlFrom.ATTRIBUTE_VALUE_XHTML {
		a_attribute_value_xhtmlTo.ATTRIBUTE_VALUE_XHTML = append(a_attribute_value_xhtmlTo.ATTRIBUTE_VALUE_XHTML, GongCopyBranchATTRIBUTE_VALUE_XHTML(mapOrigCopy, _attribute_value_xhtml))
	}

	return
}

func GongCopyBranchA_ATTRIBUTE_VALUE_XHTML_1(mapOrigCopy map[any]any, a_attribute_value_xhtml_1From *A_ATTRIBUTE_VALUE_XHTML_1) (a_attribute_value_xhtml_1To *A_ATTRIBUTE_VALUE_XHTML_1) {

	// a_attribute_value_xhtml_1From has already been copied
	if _a_attribute_value_xhtml_1To, ok := mapOrigCopy[a_attribute_value_xhtml_1From]; ok {
		a_attribute_value_xhtml_1To = _a_attribute_value_xhtml_1To.(*A_ATTRIBUTE_VALUE_XHTML_1)
		return
	}

	a_attribute_value_xhtml_1To = new(A_ATTRIBUTE_VALUE_XHTML_1)
	mapOrigCopy[a_attribute_value_xhtml_1From] = a_attribute_value_xhtml_1To
	a_attribute_value_xhtml_1From.GongCopyBasicFields(a_attribute_value_xhtml_1To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_boolean := range a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_BOOLEAN {
		a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_BOOLEAN = append(a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_BOOLEAN, GongCopyBranchATTRIBUTE_VALUE_BOOLEAN(mapOrigCopy, _attribute_value_boolean))
	}
	for _, _attribute_value_date := range a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_DATE {
		a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_DATE = append(a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_DATE, GongCopyBranchATTRIBUTE_VALUE_DATE(mapOrigCopy, _attribute_value_date))
	}
	for _, _attribute_value_enumeration := range a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_ENUMERATION {
		a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_ENUMERATION = append(a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_ENUMERATION, GongCopyBranchATTRIBUTE_VALUE_ENUMERATION(mapOrigCopy, _attribute_value_enumeration))
	}
	for _, _attribute_value_integer := range a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_INTEGER {
		a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_INTEGER = append(a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_INTEGER, GongCopyBranchATTRIBUTE_VALUE_INTEGER(mapOrigCopy, _attribute_value_integer))
	}
	for _, _attribute_value_real := range a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_REAL {
		a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_REAL = append(a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_REAL, GongCopyBranchATTRIBUTE_VALUE_REAL(mapOrigCopy, _attribute_value_real))
	}
	for _, _attribute_value_string := range a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_STRING {
		a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_STRING = append(a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_STRING, GongCopyBranchATTRIBUTE_VALUE_STRING(mapOrigCopy, _attribute_value_string))
	}
	for _, _attribute_value_xhtml := range a_attribute_value_xhtml_1From.ATTRIBUTE_VALUE_XHTML {
		a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_XHTML = append(a_attribute_value_xhtml_1To.ATTRIBUTE_VALUE_XHTML, GongCopyBranchATTRIBUTE_VALUE_XHTML(mapOrigCopy, _attribute_value_xhtml))
	}

	return
}

func GongCopyBranchA_CHILDREN(mapOrigCopy map[any]any, a_childrenFrom *A_CHILDREN) (a_childrenTo *A_CHILDREN) {

	// a_childrenFrom has already been copied
	if _a_childrenTo, ok := mapOrigCopy[a_childrenFrom]; ok {
		a_childrenTo = _a_childrenTo.(*A_CHILDREN)
		return
	}

	a_childrenTo = new(A_CHILDREN)
	mapOrigCopy[a_childrenFrom] = a_childrenTo
	a_childrenFrom.GongCopyBasicFields(a_childrenTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_hierarchy := range a_childrenFrom.SPEC_HIERARCHY {
		a_childrenTo.SPEC_HIERARCHY = append(a_childrenTo.SPEC_HIERARCHY, GongCopyBranchSPEC_HIERARCHY(mapOrigCopy, _spec_hierarchy))
	}

	return
}

func GongCopyBranchA_CORE_CONTENT(mapOrigCopy map[any]any, a_core_contentFrom *A_CORE_CONTENT) (a_core_contentTo *A_CORE_CONTENT) {

	// a_core_contentFrom has already been copied
	if _a_core_contentTo, ok := mapOrigCopy[a_core_contentFrom]; ok {
		a_core_contentTo = _a_core_contentTo.(*A_CORE_CONTENT)
		return
	}

	a_core_contentTo = new(A_CORE_CONTENT)
	mapOrigCopy[a_core_contentFrom] = a_core_contentTo
	a_core_contentFrom.GongCopyBasicFields(a_core_contentTo)

	//insertion point for the staging of instances referenced by pointers
	if a_core_contentFrom.REQ_IF_CONTENT != nil {
		a_core_contentTo.REQ_IF_CONTENT = GongCopyBranchREQ_IF_CONTENT(mapOrigCopy, a_core_contentFrom.REQ_IF_CONTENT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_DATATYPES(mapOrigCopy map[any]any, a_datatypesFrom *A_DATATYPES) (a_datatypesTo *A_DATATYPES) {

	// a_datatypesFrom has already been copied
	if _a_datatypesTo, ok := mapOrigCopy[a_datatypesFrom]; ok {
		a_datatypesTo = _a_datatypesTo.(*A_DATATYPES)
		return
	}

	a_datatypesTo = new(A_DATATYPES)
	mapOrigCopy[a_datatypesFrom] = a_datatypesTo
	a_datatypesFrom.GongCopyBasicFields(a_datatypesTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _datatype_definition_boolean := range a_datatypesFrom.DATATYPE_DEFINITION_BOOLEAN {
		a_datatypesTo.DATATYPE_DEFINITION_BOOLEAN = append(a_datatypesTo.DATATYPE_DEFINITION_BOOLEAN, GongCopyBranchDATATYPE_DEFINITION_BOOLEAN(mapOrigCopy, _datatype_definition_boolean))
	}
	for _, _datatype_definition_date := range a_datatypesFrom.DATATYPE_DEFINITION_DATE {
		a_datatypesTo.DATATYPE_DEFINITION_DATE = append(a_datatypesTo.DATATYPE_DEFINITION_DATE, GongCopyBranchDATATYPE_DEFINITION_DATE(mapOrigCopy, _datatype_definition_date))
	}
	for _, _datatype_definition_enumeration := range a_datatypesFrom.DATATYPE_DEFINITION_ENUMERATION {
		a_datatypesTo.DATATYPE_DEFINITION_ENUMERATION = append(a_datatypesTo.DATATYPE_DEFINITION_ENUMERATION, GongCopyBranchDATATYPE_DEFINITION_ENUMERATION(mapOrigCopy, _datatype_definition_enumeration))
	}
	for _, _datatype_definition_integer := range a_datatypesFrom.DATATYPE_DEFINITION_INTEGER {
		a_datatypesTo.DATATYPE_DEFINITION_INTEGER = append(a_datatypesTo.DATATYPE_DEFINITION_INTEGER, GongCopyBranchDATATYPE_DEFINITION_INTEGER(mapOrigCopy, _datatype_definition_integer))
	}
	for _, _datatype_definition_real := range a_datatypesFrom.DATATYPE_DEFINITION_REAL {
		a_datatypesTo.DATATYPE_DEFINITION_REAL = append(a_datatypesTo.DATATYPE_DEFINITION_REAL, GongCopyBranchDATATYPE_DEFINITION_REAL(mapOrigCopy, _datatype_definition_real))
	}
	for _, _datatype_definition_string := range a_datatypesFrom.DATATYPE_DEFINITION_STRING {
		a_datatypesTo.DATATYPE_DEFINITION_STRING = append(a_datatypesTo.DATATYPE_DEFINITION_STRING, GongCopyBranchDATATYPE_DEFINITION_STRING(mapOrigCopy, _datatype_definition_string))
	}
	for _, _datatype_definition_xhtml := range a_datatypesFrom.DATATYPE_DEFINITION_XHTML {
		a_datatypesTo.DATATYPE_DEFINITION_XHTML = append(a_datatypesTo.DATATYPE_DEFINITION_XHTML, GongCopyBranchDATATYPE_DEFINITION_XHTML(mapOrigCopy, _datatype_definition_xhtml))
	}

	return
}

func GongCopyBranchA_DATATYPE_DEFINITION_BOOLEAN_REF(mapOrigCopy map[any]any, a_datatype_definition_boolean_refFrom *A_DATATYPE_DEFINITION_BOOLEAN_REF) (a_datatype_definition_boolean_refTo *A_DATATYPE_DEFINITION_BOOLEAN_REF) {

	// a_datatype_definition_boolean_refFrom has already been copied
	if _a_datatype_definition_boolean_refTo, ok := mapOrigCopy[a_datatype_definition_boolean_refFrom]; ok {
		a_datatype_definition_boolean_refTo = _a_datatype_definition_boolean_refTo.(*A_DATATYPE_DEFINITION_BOOLEAN_REF)
		return
	}

	a_datatype_definition_boolean_refTo = new(A_DATATYPE_DEFINITION_BOOLEAN_REF)
	mapOrigCopy[a_datatype_definition_boolean_refFrom] = a_datatype_definition_boolean_refTo
	a_datatype_definition_boolean_refFrom.GongCopyBasicFields(a_datatype_definition_boolean_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_DATATYPE_DEFINITION_DATE_REF(mapOrigCopy map[any]any, a_datatype_definition_date_refFrom *A_DATATYPE_DEFINITION_DATE_REF) (a_datatype_definition_date_refTo *A_DATATYPE_DEFINITION_DATE_REF) {

	// a_datatype_definition_date_refFrom has already been copied
	if _a_datatype_definition_date_refTo, ok := mapOrigCopy[a_datatype_definition_date_refFrom]; ok {
		a_datatype_definition_date_refTo = _a_datatype_definition_date_refTo.(*A_DATATYPE_DEFINITION_DATE_REF)
		return
	}

	a_datatype_definition_date_refTo = new(A_DATATYPE_DEFINITION_DATE_REF)
	mapOrigCopy[a_datatype_definition_date_refFrom] = a_datatype_definition_date_refTo
	a_datatype_definition_date_refFrom.GongCopyBasicFields(a_datatype_definition_date_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_DATATYPE_DEFINITION_ENUMERATION_REF(mapOrigCopy map[any]any, a_datatype_definition_enumeration_refFrom *A_DATATYPE_DEFINITION_ENUMERATION_REF) (a_datatype_definition_enumeration_refTo *A_DATATYPE_DEFINITION_ENUMERATION_REF) {

	// a_datatype_definition_enumeration_refFrom has already been copied
	if _a_datatype_definition_enumeration_refTo, ok := mapOrigCopy[a_datatype_definition_enumeration_refFrom]; ok {
		a_datatype_definition_enumeration_refTo = _a_datatype_definition_enumeration_refTo.(*A_DATATYPE_DEFINITION_ENUMERATION_REF)
		return
	}

	a_datatype_definition_enumeration_refTo = new(A_DATATYPE_DEFINITION_ENUMERATION_REF)
	mapOrigCopy[a_datatype_definition_enumeration_refFrom] = a_datatype_definition_enumeration_refTo
	a_datatype_definition_enumeration_refFrom.GongCopyBasicFields(a_datatype_definition_enumeration_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_DATATYPE_DEFINITION_INTEGER_REF(mapOrigCopy map[any]any, a_datatype_definition_integer_refFrom *A_DATATYPE_DEFINITION_INTEGER_REF) (a_datatype_definition_integer_refTo *A_DATATYPE_DEFINITION_INTEGER_REF) {

	// a_datatype_definition_integer_refFrom has already been copied
	if _a_datatype_definition_integer_refTo, ok := mapOrigCopy[a_datatype_definition_integer_refFrom]; ok {
		a_datatype_definition_integer_refTo = _a_datatype_definition_integer_refTo.(*A_DATATYPE_DEFINITION_INTEGER_REF)
		return
	}

	a_datatype_definition_integer_refTo = new(A_DATATYPE_DEFINITION_INTEGER_REF)
	mapOrigCopy[a_datatype_definition_integer_refFrom] = a_datatype_definition_integer_refTo
	a_datatype_definition_integer_refFrom.GongCopyBasicFields(a_datatype_definition_integer_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_DATATYPE_DEFINITION_REAL_REF(mapOrigCopy map[any]any, a_datatype_definition_real_refFrom *A_DATATYPE_DEFINITION_REAL_REF) (a_datatype_definition_real_refTo *A_DATATYPE_DEFINITION_REAL_REF) {

	// a_datatype_definition_real_refFrom has already been copied
	if _a_datatype_definition_real_refTo, ok := mapOrigCopy[a_datatype_definition_real_refFrom]; ok {
		a_datatype_definition_real_refTo = _a_datatype_definition_real_refTo.(*A_DATATYPE_DEFINITION_REAL_REF)
		return
	}

	a_datatype_definition_real_refTo = new(A_DATATYPE_DEFINITION_REAL_REF)
	mapOrigCopy[a_datatype_definition_real_refFrom] = a_datatype_definition_real_refTo
	a_datatype_definition_real_refFrom.GongCopyBasicFields(a_datatype_definition_real_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_DATATYPE_DEFINITION_STRING_REF(mapOrigCopy map[any]any, a_datatype_definition_string_refFrom *A_DATATYPE_DEFINITION_STRING_REF) (a_datatype_definition_string_refTo *A_DATATYPE_DEFINITION_STRING_REF) {

	// a_datatype_definition_string_refFrom has already been copied
	if _a_datatype_definition_string_refTo, ok := mapOrigCopy[a_datatype_definition_string_refFrom]; ok {
		a_datatype_definition_string_refTo = _a_datatype_definition_string_refTo.(*A_DATATYPE_DEFINITION_STRING_REF)
		return
	}

	a_datatype_definition_string_refTo = new(A_DATATYPE_DEFINITION_STRING_REF)
	mapOrigCopy[a_datatype_definition_string_refFrom] = a_datatype_definition_string_refTo
	a_datatype_definition_string_refFrom.GongCopyBasicFields(a_datatype_definition_string_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_DATATYPE_DEFINITION_XHTML_REF(mapOrigCopy map[any]any, a_datatype_definition_xhtml_refFrom *A_DATATYPE_DEFINITION_XHTML_REF) (a_datatype_definition_xhtml_refTo *A_DATATYPE_DEFINITION_XHTML_REF) {

	// a_datatype_definition_xhtml_refFrom has already been copied
	if _a_datatype_definition_xhtml_refTo, ok := mapOrigCopy[a_datatype_definition_xhtml_refFrom]; ok {
		a_datatype_definition_xhtml_refTo = _a_datatype_definition_xhtml_refTo.(*A_DATATYPE_DEFINITION_XHTML_REF)
		return
	}

	a_datatype_definition_xhtml_refTo = new(A_DATATYPE_DEFINITION_XHTML_REF)
	mapOrigCopy[a_datatype_definition_xhtml_refFrom] = a_datatype_definition_xhtml_refTo
	a_datatype_definition_xhtml_refFrom.GongCopyBasicFields(a_datatype_definition_xhtml_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_EDITABLE_ATTS(mapOrigCopy map[any]any, a_editable_attsFrom *A_EDITABLE_ATTS) (a_editable_attsTo *A_EDITABLE_ATTS) {

	// a_editable_attsFrom has already been copied
	if _a_editable_attsTo, ok := mapOrigCopy[a_editable_attsFrom]; ok {
		a_editable_attsTo = _a_editable_attsTo.(*A_EDITABLE_ATTS)
		return
	}

	a_editable_attsTo = new(A_EDITABLE_ATTS)
	mapOrigCopy[a_editable_attsFrom] = a_editable_attsTo
	a_editable_attsFrom.GongCopyBasicFields(a_editable_attsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_ENUM_VALUE_REF(mapOrigCopy map[any]any, a_enum_value_refFrom *A_ENUM_VALUE_REF) (a_enum_value_refTo *A_ENUM_VALUE_REF) {

	// a_enum_value_refFrom has already been copied
	if _a_enum_value_refTo, ok := mapOrigCopy[a_enum_value_refFrom]; ok {
		a_enum_value_refTo = _a_enum_value_refTo.(*A_ENUM_VALUE_REF)
		return
	}

	a_enum_value_refTo = new(A_ENUM_VALUE_REF)
	mapOrigCopy[a_enum_value_refFrom] = a_enum_value_refTo
	a_enum_value_refFrom.GongCopyBasicFields(a_enum_value_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_OBJECT(mapOrigCopy map[any]any, a_objectFrom *A_OBJECT) (a_objectTo *A_OBJECT) {

	// a_objectFrom has already been copied
	if _a_objectTo, ok := mapOrigCopy[a_objectFrom]; ok {
		a_objectTo = _a_objectTo.(*A_OBJECT)
		return
	}

	a_objectTo = new(A_OBJECT)
	mapOrigCopy[a_objectFrom] = a_objectTo
	a_objectFrom.GongCopyBasicFields(a_objectTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_PROPERTIES(mapOrigCopy map[any]any, a_propertiesFrom *A_PROPERTIES) (a_propertiesTo *A_PROPERTIES) {

	// a_propertiesFrom has already been copied
	if _a_propertiesTo, ok := mapOrigCopy[a_propertiesFrom]; ok {
		a_propertiesTo = _a_propertiesTo.(*A_PROPERTIES)
		return
	}

	a_propertiesTo = new(A_PROPERTIES)
	mapOrigCopy[a_propertiesFrom] = a_propertiesTo
	a_propertiesFrom.GongCopyBasicFields(a_propertiesTo)

	//insertion point for the staging of instances referenced by pointers
	if a_propertiesFrom.EMBEDDED_VALUE != nil {
		a_propertiesTo.EMBEDDED_VALUE = GongCopyBranchEMBEDDED_VALUE(mapOrigCopy, a_propertiesFrom.EMBEDDED_VALUE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_RELATION_GROUP_TYPE_REF(mapOrigCopy map[any]any, a_relation_group_type_refFrom *A_RELATION_GROUP_TYPE_REF) (a_relation_group_type_refTo *A_RELATION_GROUP_TYPE_REF) {

	// a_relation_group_type_refFrom has already been copied
	if _a_relation_group_type_refTo, ok := mapOrigCopy[a_relation_group_type_refFrom]; ok {
		a_relation_group_type_refTo = _a_relation_group_type_refTo.(*A_RELATION_GROUP_TYPE_REF)
		return
	}

	a_relation_group_type_refTo = new(A_RELATION_GROUP_TYPE_REF)
	mapOrigCopy[a_relation_group_type_refFrom] = a_relation_group_type_refTo
	a_relation_group_type_refFrom.GongCopyBasicFields(a_relation_group_type_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_SOURCE_1(mapOrigCopy map[any]any, a_source_1From *A_SOURCE_1) (a_source_1To *A_SOURCE_1) {

	// a_source_1From has already been copied
	if _a_source_1To, ok := mapOrigCopy[a_source_1From]; ok {
		a_source_1To = _a_source_1To.(*A_SOURCE_1)
		return
	}

	a_source_1To = new(A_SOURCE_1)
	mapOrigCopy[a_source_1From] = a_source_1To
	a_source_1From.GongCopyBasicFields(a_source_1To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_SOURCE_SPECIFICATION_1(mapOrigCopy map[any]any, a_source_specification_1From *A_SOURCE_SPECIFICATION_1) (a_source_specification_1To *A_SOURCE_SPECIFICATION_1) {

	// a_source_specification_1From has already been copied
	if _a_source_specification_1To, ok := mapOrigCopy[a_source_specification_1From]; ok {
		a_source_specification_1To = _a_source_specification_1To.(*A_SOURCE_SPECIFICATION_1)
		return
	}

	a_source_specification_1To = new(A_SOURCE_SPECIFICATION_1)
	mapOrigCopy[a_source_specification_1From] = a_source_specification_1To
	a_source_specification_1From.GongCopyBasicFields(a_source_specification_1To)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_SPECIFICATIONS(mapOrigCopy map[any]any, a_specificationsFrom *A_SPECIFICATIONS) (a_specificationsTo *A_SPECIFICATIONS) {

	// a_specificationsFrom has already been copied
	if _a_specificationsTo, ok := mapOrigCopy[a_specificationsFrom]; ok {
		a_specificationsTo = _a_specificationsTo.(*A_SPECIFICATIONS)
		return
	}

	a_specificationsTo = new(A_SPECIFICATIONS)
	mapOrigCopy[a_specificationsFrom] = a_specificationsTo
	a_specificationsFrom.GongCopyBasicFields(a_specificationsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _specification := range a_specificationsFrom.SPECIFICATION {
		a_specificationsTo.SPECIFICATION = append(a_specificationsTo.SPECIFICATION, GongCopyBranchSPECIFICATION(mapOrigCopy, _specification))
	}

	return
}

func GongCopyBranchA_SPECIFICATION_TYPE_REF(mapOrigCopy map[any]any, a_specification_type_refFrom *A_SPECIFICATION_TYPE_REF) (a_specification_type_refTo *A_SPECIFICATION_TYPE_REF) {

	// a_specification_type_refFrom has already been copied
	if _a_specification_type_refTo, ok := mapOrigCopy[a_specification_type_refFrom]; ok {
		a_specification_type_refTo = _a_specification_type_refTo.(*A_SPECIFICATION_TYPE_REF)
		return
	}

	a_specification_type_refTo = new(A_SPECIFICATION_TYPE_REF)
	mapOrigCopy[a_specification_type_refFrom] = a_specification_type_refTo
	a_specification_type_refFrom.GongCopyBasicFields(a_specification_type_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_SPECIFIED_VALUES(mapOrigCopy map[any]any, a_specified_valuesFrom *A_SPECIFIED_VALUES) (a_specified_valuesTo *A_SPECIFIED_VALUES) {

	// a_specified_valuesFrom has already been copied
	if _a_specified_valuesTo, ok := mapOrigCopy[a_specified_valuesFrom]; ok {
		a_specified_valuesTo = _a_specified_valuesTo.(*A_SPECIFIED_VALUES)
		return
	}

	a_specified_valuesTo = new(A_SPECIFIED_VALUES)
	mapOrigCopy[a_specified_valuesFrom] = a_specified_valuesTo
	a_specified_valuesFrom.GongCopyBasicFields(a_specified_valuesTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enum_value := range a_specified_valuesFrom.ENUM_VALUE {
		a_specified_valuesTo.ENUM_VALUE = append(a_specified_valuesTo.ENUM_VALUE, GongCopyBranchENUM_VALUE(mapOrigCopy, _enum_value))
	}

	return
}

func GongCopyBranchA_SPEC_ATTRIBUTES(mapOrigCopy map[any]any, a_spec_attributesFrom *A_SPEC_ATTRIBUTES) (a_spec_attributesTo *A_SPEC_ATTRIBUTES) {

	// a_spec_attributesFrom has already been copied
	if _a_spec_attributesTo, ok := mapOrigCopy[a_spec_attributesFrom]; ok {
		a_spec_attributesTo = _a_spec_attributesTo.(*A_SPEC_ATTRIBUTES)
		return
	}

	a_spec_attributesTo = new(A_SPEC_ATTRIBUTES)
	mapOrigCopy[a_spec_attributesFrom] = a_spec_attributesTo
	a_spec_attributesFrom.GongCopyBasicFields(a_spec_attributesTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_definition_boolean := range a_spec_attributesFrom.ATTRIBUTE_DEFINITION_BOOLEAN {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_BOOLEAN = append(a_spec_attributesTo.ATTRIBUTE_DEFINITION_BOOLEAN, GongCopyBranchATTRIBUTE_DEFINITION_BOOLEAN(mapOrigCopy, _attribute_definition_boolean))
	}
	for _, _attribute_definition_date := range a_spec_attributesFrom.ATTRIBUTE_DEFINITION_DATE {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_DATE = append(a_spec_attributesTo.ATTRIBUTE_DEFINITION_DATE, GongCopyBranchATTRIBUTE_DEFINITION_DATE(mapOrigCopy, _attribute_definition_date))
	}
	for _, _attribute_definition_enumeration := range a_spec_attributesFrom.ATTRIBUTE_DEFINITION_ENUMERATION {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_ENUMERATION = append(a_spec_attributesTo.ATTRIBUTE_DEFINITION_ENUMERATION, GongCopyBranchATTRIBUTE_DEFINITION_ENUMERATION(mapOrigCopy, _attribute_definition_enumeration))
	}
	for _, _attribute_definition_integer := range a_spec_attributesFrom.ATTRIBUTE_DEFINITION_INTEGER {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_INTEGER = append(a_spec_attributesTo.ATTRIBUTE_DEFINITION_INTEGER, GongCopyBranchATTRIBUTE_DEFINITION_INTEGER(mapOrigCopy, _attribute_definition_integer))
	}
	for _, _attribute_definition_real := range a_spec_attributesFrom.ATTRIBUTE_DEFINITION_REAL {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_REAL = append(a_spec_attributesTo.ATTRIBUTE_DEFINITION_REAL, GongCopyBranchATTRIBUTE_DEFINITION_REAL(mapOrigCopy, _attribute_definition_real))
	}
	for _, _attribute_definition_string := range a_spec_attributesFrom.ATTRIBUTE_DEFINITION_STRING {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_STRING = append(a_spec_attributesTo.ATTRIBUTE_DEFINITION_STRING, GongCopyBranchATTRIBUTE_DEFINITION_STRING(mapOrigCopy, _attribute_definition_string))
	}
	for _, _attribute_definition_xhtml := range a_spec_attributesFrom.ATTRIBUTE_DEFINITION_XHTML {
		a_spec_attributesTo.ATTRIBUTE_DEFINITION_XHTML = append(a_spec_attributesTo.ATTRIBUTE_DEFINITION_XHTML, GongCopyBranchATTRIBUTE_DEFINITION_XHTML(mapOrigCopy, _attribute_definition_xhtml))
	}

	return
}

func GongCopyBranchA_SPEC_OBJECTS(mapOrigCopy map[any]any, a_spec_objectsFrom *A_SPEC_OBJECTS) (a_spec_objectsTo *A_SPEC_OBJECTS) {

	// a_spec_objectsFrom has already been copied
	if _a_spec_objectsTo, ok := mapOrigCopy[a_spec_objectsFrom]; ok {
		a_spec_objectsTo = _a_spec_objectsTo.(*A_SPEC_OBJECTS)
		return
	}

	a_spec_objectsTo = new(A_SPEC_OBJECTS)
	mapOrigCopy[a_spec_objectsFrom] = a_spec_objectsTo
	a_spec_objectsFrom.GongCopyBasicFields(a_spec_objectsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_object := range a_spec_objectsFrom.SPEC_OBJECT {
		a_spec_objectsTo.SPEC_OBJECT = append(a_spec_objectsTo.SPEC_OBJECT, GongCopyBranchSPEC_OBJECT(mapOrigCopy, _spec_object))
	}

	return
}

func GongCopyBranchA_SPEC_OBJECT_TYPE_REF(mapOrigCopy map[any]any, a_spec_object_type_refFrom *A_SPEC_OBJECT_TYPE_REF) (a_spec_object_type_refTo *A_SPEC_OBJECT_TYPE_REF) {

	// a_spec_object_type_refFrom has already been copied
	if _a_spec_object_type_refTo, ok := mapOrigCopy[a_spec_object_type_refFrom]; ok {
		a_spec_object_type_refTo = _a_spec_object_type_refTo.(*A_SPEC_OBJECT_TYPE_REF)
		return
	}

	a_spec_object_type_refTo = new(A_SPEC_OBJECT_TYPE_REF)
	mapOrigCopy[a_spec_object_type_refFrom] = a_spec_object_type_refTo
	a_spec_object_type_refFrom.GongCopyBasicFields(a_spec_object_type_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_SPEC_RELATIONS(mapOrigCopy map[any]any, a_spec_relationsFrom *A_SPEC_RELATIONS) (a_spec_relationsTo *A_SPEC_RELATIONS) {

	// a_spec_relationsFrom has already been copied
	if _a_spec_relationsTo, ok := mapOrigCopy[a_spec_relationsFrom]; ok {
		a_spec_relationsTo = _a_spec_relationsTo.(*A_SPEC_RELATIONS)
		return
	}

	a_spec_relationsTo = new(A_SPEC_RELATIONS)
	mapOrigCopy[a_spec_relationsFrom] = a_spec_relationsTo
	a_spec_relationsFrom.GongCopyBasicFields(a_spec_relationsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_relation := range a_spec_relationsFrom.SPEC_RELATION {
		a_spec_relationsTo.SPEC_RELATION = append(a_spec_relationsTo.SPEC_RELATION, GongCopyBranchSPEC_RELATION(mapOrigCopy, _spec_relation))
	}

	return
}

func GongCopyBranchA_SPEC_RELATION_GROUPS(mapOrigCopy map[any]any, a_spec_relation_groupsFrom *A_SPEC_RELATION_GROUPS) (a_spec_relation_groupsTo *A_SPEC_RELATION_GROUPS) {

	// a_spec_relation_groupsFrom has already been copied
	if _a_spec_relation_groupsTo, ok := mapOrigCopy[a_spec_relation_groupsFrom]; ok {
		a_spec_relation_groupsTo = _a_spec_relation_groupsTo.(*A_SPEC_RELATION_GROUPS)
		return
	}

	a_spec_relation_groupsTo = new(A_SPEC_RELATION_GROUPS)
	mapOrigCopy[a_spec_relation_groupsFrom] = a_spec_relation_groupsTo
	a_spec_relation_groupsFrom.GongCopyBasicFields(a_spec_relation_groupsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relation_group := range a_spec_relation_groupsFrom.RELATION_GROUP {
		a_spec_relation_groupsTo.RELATION_GROUP = append(a_spec_relation_groupsTo.RELATION_GROUP, GongCopyBranchRELATION_GROUP(mapOrigCopy, _relation_group))
	}

	return
}

func GongCopyBranchA_SPEC_RELATION_REF(mapOrigCopy map[any]any, a_spec_relation_refFrom *A_SPEC_RELATION_REF) (a_spec_relation_refTo *A_SPEC_RELATION_REF) {

	// a_spec_relation_refFrom has already been copied
	if _a_spec_relation_refTo, ok := mapOrigCopy[a_spec_relation_refFrom]; ok {
		a_spec_relation_refTo = _a_spec_relation_refTo.(*A_SPEC_RELATION_REF)
		return
	}

	a_spec_relation_refTo = new(A_SPEC_RELATION_REF)
	mapOrigCopy[a_spec_relation_refFrom] = a_spec_relation_refTo
	a_spec_relation_refFrom.GongCopyBasicFields(a_spec_relation_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_SPEC_RELATION_TYPE_REF(mapOrigCopy map[any]any, a_spec_relation_type_refFrom *A_SPEC_RELATION_TYPE_REF) (a_spec_relation_type_refTo *A_SPEC_RELATION_TYPE_REF) {

	// a_spec_relation_type_refFrom has already been copied
	if _a_spec_relation_type_refTo, ok := mapOrigCopy[a_spec_relation_type_refFrom]; ok {
		a_spec_relation_type_refTo = _a_spec_relation_type_refTo.(*A_SPEC_RELATION_TYPE_REF)
		return
	}

	a_spec_relation_type_refTo = new(A_SPEC_RELATION_TYPE_REF)
	mapOrigCopy[a_spec_relation_type_refFrom] = a_spec_relation_type_refTo
	a_spec_relation_type_refFrom.GongCopyBasicFields(a_spec_relation_type_refTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_SPEC_TYPES(mapOrigCopy map[any]any, a_spec_typesFrom *A_SPEC_TYPES) (a_spec_typesTo *A_SPEC_TYPES) {

	// a_spec_typesFrom has already been copied
	if _a_spec_typesTo, ok := mapOrigCopy[a_spec_typesFrom]; ok {
		a_spec_typesTo = _a_spec_typesTo.(*A_SPEC_TYPES)
		return
	}

	a_spec_typesTo = new(A_SPEC_TYPES)
	mapOrigCopy[a_spec_typesFrom] = a_spec_typesTo
	a_spec_typesFrom.GongCopyBasicFields(a_spec_typesTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relation_group_type := range a_spec_typesFrom.RELATION_GROUP_TYPE {
		a_spec_typesTo.RELATION_GROUP_TYPE = append(a_spec_typesTo.RELATION_GROUP_TYPE, GongCopyBranchRELATION_GROUP_TYPE(mapOrigCopy, _relation_group_type))
	}
	for _, _spec_object_type := range a_spec_typesFrom.SPEC_OBJECT_TYPE {
		a_spec_typesTo.SPEC_OBJECT_TYPE = append(a_spec_typesTo.SPEC_OBJECT_TYPE, GongCopyBranchSPEC_OBJECT_TYPE(mapOrigCopy, _spec_object_type))
	}
	for _, _spec_relation_type := range a_spec_typesFrom.SPEC_RELATION_TYPE {
		a_spec_typesTo.SPEC_RELATION_TYPE = append(a_spec_typesTo.SPEC_RELATION_TYPE, GongCopyBranchSPEC_RELATION_TYPE(mapOrigCopy, _spec_relation_type))
	}
	for _, _specification_type := range a_spec_typesFrom.SPECIFICATION_TYPE {
		a_spec_typesTo.SPECIFICATION_TYPE = append(a_spec_typesTo.SPECIFICATION_TYPE, GongCopyBranchSPECIFICATION_TYPE(mapOrigCopy, _specification_type))
	}

	return
}

func GongCopyBranchA_THE_HEADER(mapOrigCopy map[any]any, a_the_headerFrom *A_THE_HEADER) (a_the_headerTo *A_THE_HEADER) {

	// a_the_headerFrom has already been copied
	if _a_the_headerTo, ok := mapOrigCopy[a_the_headerFrom]; ok {
		a_the_headerTo = _a_the_headerTo.(*A_THE_HEADER)
		return
	}

	a_the_headerTo = new(A_THE_HEADER)
	mapOrigCopy[a_the_headerFrom] = a_the_headerTo
	a_the_headerFrom.GongCopyBasicFields(a_the_headerTo)

	//insertion point for the staging of instances referenced by pointers
	if a_the_headerFrom.REQ_IF_HEADER != nil {
		a_the_headerTo.REQ_IF_HEADER = GongCopyBranchREQ_IF_HEADER(mapOrigCopy, a_the_headerFrom.REQ_IF_HEADER)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchA_TOOL_EXTENSIONS(mapOrigCopy map[any]any, a_tool_extensionsFrom *A_TOOL_EXTENSIONS) (a_tool_extensionsTo *A_TOOL_EXTENSIONS) {

	// a_tool_extensionsFrom has already been copied
	if _a_tool_extensionsTo, ok := mapOrigCopy[a_tool_extensionsFrom]; ok {
		a_tool_extensionsTo = _a_tool_extensionsTo.(*A_TOOL_EXTENSIONS)
		return
	}

	a_tool_extensionsTo = new(A_TOOL_EXTENSIONS)
	mapOrigCopy[a_tool_extensionsFrom] = a_tool_extensionsTo
	a_tool_extensionsFrom.GongCopyBasicFields(a_tool_extensionsTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _req_if_tool_extension := range a_tool_extensionsFrom.REQ_IF_TOOL_EXTENSION {
		a_tool_extensionsTo.REQ_IF_TOOL_EXTENSION = append(a_tool_extensionsTo.REQ_IF_TOOL_EXTENSION, GongCopyBranchREQ_IF_TOOL_EXTENSION(mapOrigCopy, _req_if_tool_extension))
	}

	return
}

func GongCopyBranchDATATYPE_DEFINITION_BOOLEAN(mapOrigCopy map[any]any, datatype_definition_booleanFrom *DATATYPE_DEFINITION_BOOLEAN) (datatype_definition_booleanTo *DATATYPE_DEFINITION_BOOLEAN) {

	// datatype_definition_booleanFrom has already been copied
	if _datatype_definition_booleanTo, ok := mapOrigCopy[datatype_definition_booleanFrom]; ok {
		datatype_definition_booleanTo = _datatype_definition_booleanTo.(*DATATYPE_DEFINITION_BOOLEAN)
		return
	}

	datatype_definition_booleanTo = new(DATATYPE_DEFINITION_BOOLEAN)
	mapOrigCopy[datatype_definition_booleanFrom] = datatype_definition_booleanTo
	datatype_definition_booleanFrom.GongCopyBasicFields(datatype_definition_booleanTo)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_booleanFrom.ALTERNATIVE_ID != nil {
		datatype_definition_booleanTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, datatype_definition_booleanFrom.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDATATYPE_DEFINITION_DATE(mapOrigCopy map[any]any, datatype_definition_dateFrom *DATATYPE_DEFINITION_DATE) (datatype_definition_dateTo *DATATYPE_DEFINITION_DATE) {

	// datatype_definition_dateFrom has already been copied
	if _datatype_definition_dateTo, ok := mapOrigCopy[datatype_definition_dateFrom]; ok {
		datatype_definition_dateTo = _datatype_definition_dateTo.(*DATATYPE_DEFINITION_DATE)
		return
	}

	datatype_definition_dateTo = new(DATATYPE_DEFINITION_DATE)
	mapOrigCopy[datatype_definition_dateFrom] = datatype_definition_dateTo
	datatype_definition_dateFrom.GongCopyBasicFields(datatype_definition_dateTo)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_dateFrom.ALTERNATIVE_ID != nil {
		datatype_definition_dateTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, datatype_definition_dateFrom.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDATATYPE_DEFINITION_ENUMERATION(mapOrigCopy map[any]any, datatype_definition_enumerationFrom *DATATYPE_DEFINITION_ENUMERATION) (datatype_definition_enumerationTo *DATATYPE_DEFINITION_ENUMERATION) {

	// datatype_definition_enumerationFrom has already been copied
	if _datatype_definition_enumerationTo, ok := mapOrigCopy[datatype_definition_enumerationFrom]; ok {
		datatype_definition_enumerationTo = _datatype_definition_enumerationTo.(*DATATYPE_DEFINITION_ENUMERATION)
		return
	}

	datatype_definition_enumerationTo = new(DATATYPE_DEFINITION_ENUMERATION)
	mapOrigCopy[datatype_definition_enumerationFrom] = datatype_definition_enumerationTo
	datatype_definition_enumerationFrom.GongCopyBasicFields(datatype_definition_enumerationTo)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_enumerationFrom.ALTERNATIVE_ID != nil {
		datatype_definition_enumerationTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, datatype_definition_enumerationFrom.ALTERNATIVE_ID)
	}
	if datatype_definition_enumerationFrom.SPECIFIED_VALUES != nil {
		datatype_definition_enumerationTo.SPECIFIED_VALUES = GongCopyBranchA_SPECIFIED_VALUES(mapOrigCopy, datatype_definition_enumerationFrom.SPECIFIED_VALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDATATYPE_DEFINITION_INTEGER(mapOrigCopy map[any]any, datatype_definition_integerFrom *DATATYPE_DEFINITION_INTEGER) (datatype_definition_integerTo *DATATYPE_DEFINITION_INTEGER) {

	// datatype_definition_integerFrom has already been copied
	if _datatype_definition_integerTo, ok := mapOrigCopy[datatype_definition_integerFrom]; ok {
		datatype_definition_integerTo = _datatype_definition_integerTo.(*DATATYPE_DEFINITION_INTEGER)
		return
	}

	datatype_definition_integerTo = new(DATATYPE_DEFINITION_INTEGER)
	mapOrigCopy[datatype_definition_integerFrom] = datatype_definition_integerTo
	datatype_definition_integerFrom.GongCopyBasicFields(datatype_definition_integerTo)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_integerFrom.ALTERNATIVE_ID != nil {
		datatype_definition_integerTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, datatype_definition_integerFrom.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDATATYPE_DEFINITION_REAL(mapOrigCopy map[any]any, datatype_definition_realFrom *DATATYPE_DEFINITION_REAL) (datatype_definition_realTo *DATATYPE_DEFINITION_REAL) {

	// datatype_definition_realFrom has already been copied
	if _datatype_definition_realTo, ok := mapOrigCopy[datatype_definition_realFrom]; ok {
		datatype_definition_realTo = _datatype_definition_realTo.(*DATATYPE_DEFINITION_REAL)
		return
	}

	datatype_definition_realTo = new(DATATYPE_DEFINITION_REAL)
	mapOrigCopy[datatype_definition_realFrom] = datatype_definition_realTo
	datatype_definition_realFrom.GongCopyBasicFields(datatype_definition_realTo)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_realFrom.ALTERNATIVE_ID != nil {
		datatype_definition_realTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, datatype_definition_realFrom.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDATATYPE_DEFINITION_STRING(mapOrigCopy map[any]any, datatype_definition_stringFrom *DATATYPE_DEFINITION_STRING) (datatype_definition_stringTo *DATATYPE_DEFINITION_STRING) {

	// datatype_definition_stringFrom has already been copied
	if _datatype_definition_stringTo, ok := mapOrigCopy[datatype_definition_stringFrom]; ok {
		datatype_definition_stringTo = _datatype_definition_stringTo.(*DATATYPE_DEFINITION_STRING)
		return
	}

	datatype_definition_stringTo = new(DATATYPE_DEFINITION_STRING)
	mapOrigCopy[datatype_definition_stringFrom] = datatype_definition_stringTo
	datatype_definition_stringFrom.GongCopyBasicFields(datatype_definition_stringTo)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_stringFrom.ALTERNATIVE_ID != nil {
		datatype_definition_stringTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, datatype_definition_stringFrom.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchDATATYPE_DEFINITION_XHTML(mapOrigCopy map[any]any, datatype_definition_xhtmlFrom *DATATYPE_DEFINITION_XHTML) (datatype_definition_xhtmlTo *DATATYPE_DEFINITION_XHTML) {

	// datatype_definition_xhtmlFrom has already been copied
	if _datatype_definition_xhtmlTo, ok := mapOrigCopy[datatype_definition_xhtmlFrom]; ok {
		datatype_definition_xhtmlTo = _datatype_definition_xhtmlTo.(*DATATYPE_DEFINITION_XHTML)
		return
	}

	datatype_definition_xhtmlTo = new(DATATYPE_DEFINITION_XHTML)
	mapOrigCopy[datatype_definition_xhtmlFrom] = datatype_definition_xhtmlTo
	datatype_definition_xhtmlFrom.GongCopyBasicFields(datatype_definition_xhtmlTo)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_xhtmlFrom.ALTERNATIVE_ID != nil {
		datatype_definition_xhtmlTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, datatype_definition_xhtmlFrom.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchEMBEDDED_VALUE(mapOrigCopy map[any]any, embedded_valueFrom *EMBEDDED_VALUE) (embedded_valueTo *EMBEDDED_VALUE) {

	// embedded_valueFrom has already been copied
	if _embedded_valueTo, ok := mapOrigCopy[embedded_valueFrom]; ok {
		embedded_valueTo = _embedded_valueTo.(*EMBEDDED_VALUE)
		return
	}

	embedded_valueTo = new(EMBEDDED_VALUE)
	mapOrigCopy[embedded_valueFrom] = embedded_valueTo
	embedded_valueFrom.GongCopyBasicFields(embedded_valueTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchENUM_VALUE(mapOrigCopy map[any]any, enum_valueFrom *ENUM_VALUE) (enum_valueTo *ENUM_VALUE) {

	// enum_valueFrom has already been copied
	if _enum_valueTo, ok := mapOrigCopy[enum_valueFrom]; ok {
		enum_valueTo = _enum_valueTo.(*ENUM_VALUE)
		return
	}

	enum_valueTo = new(ENUM_VALUE)
	mapOrigCopy[enum_valueFrom] = enum_valueTo
	enum_valueFrom.GongCopyBasicFields(enum_valueTo)

	//insertion point for the staging of instances referenced by pointers
	if enum_valueFrom.ALTERNATIVE_ID != nil {
		enum_valueTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, enum_valueFrom.ALTERNATIVE_ID)
	}
	if enum_valueFrom.PROPERTIES != nil {
		enum_valueTo.PROPERTIES = GongCopyBranchA_PROPERTIES(mapOrigCopy, enum_valueFrom.PROPERTIES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRELATION_GROUP(mapOrigCopy map[any]any, relation_groupFrom *RELATION_GROUP) (relation_groupTo *RELATION_GROUP) {

	// relation_groupFrom has already been copied
	if _relation_groupTo, ok := mapOrigCopy[relation_groupFrom]; ok {
		relation_groupTo = _relation_groupTo.(*RELATION_GROUP)
		return
	}

	relation_groupTo = new(RELATION_GROUP)
	mapOrigCopy[relation_groupFrom] = relation_groupTo
	relation_groupFrom.GongCopyBasicFields(relation_groupTo)

	//insertion point for the staging of instances referenced by pointers
	if relation_groupFrom.ALTERNATIVE_ID != nil {
		relation_groupTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, relation_groupFrom.ALTERNATIVE_ID)
	}
	if relation_groupFrom.SOURCE_SPECIFICATION != nil {
		relation_groupTo.SOURCE_SPECIFICATION = GongCopyBranchA_SOURCE_SPECIFICATION_1(mapOrigCopy, relation_groupFrom.SOURCE_SPECIFICATION)
	}
	if relation_groupFrom.SPEC_RELATIONS != nil {
		relation_groupTo.SPEC_RELATIONS = GongCopyBranchA_SPEC_RELATION_REF(mapOrigCopy, relation_groupFrom.SPEC_RELATIONS)
	}
	if relation_groupFrom.TARGET_SPECIFICATION != nil {
		relation_groupTo.TARGET_SPECIFICATION = GongCopyBranchA_SOURCE_SPECIFICATION_1(mapOrigCopy, relation_groupFrom.TARGET_SPECIFICATION)
	}
	if relation_groupFrom.TYPE != nil {
		relation_groupTo.TYPE = GongCopyBranchA_RELATION_GROUP_TYPE_REF(mapOrigCopy, relation_groupFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchRELATION_GROUP_TYPE(mapOrigCopy map[any]any, relation_group_typeFrom *RELATION_GROUP_TYPE) (relation_group_typeTo *RELATION_GROUP_TYPE) {

	// relation_group_typeFrom has already been copied
	if _relation_group_typeTo, ok := mapOrigCopy[relation_group_typeFrom]; ok {
		relation_group_typeTo = _relation_group_typeTo.(*RELATION_GROUP_TYPE)
		return
	}

	relation_group_typeTo = new(RELATION_GROUP_TYPE)
	mapOrigCopy[relation_group_typeFrom] = relation_group_typeTo
	relation_group_typeFrom.GongCopyBasicFields(relation_group_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if relation_group_typeFrom.ALTERNATIVE_ID != nil {
		relation_group_typeTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, relation_group_typeFrom.ALTERNATIVE_ID)
	}
	if relation_group_typeFrom.SPEC_ATTRIBUTES != nil {
		relation_group_typeTo.SPEC_ATTRIBUTES = GongCopyBranchA_SPEC_ATTRIBUTES(mapOrigCopy, relation_group_typeFrom.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchREQ_IF(mapOrigCopy map[any]any, req_ifFrom *REQ_IF) (req_ifTo *REQ_IF) {

	// req_ifFrom has already been copied
	if _req_ifTo, ok := mapOrigCopy[req_ifFrom]; ok {
		req_ifTo = _req_ifTo.(*REQ_IF)
		return
	}

	req_ifTo = new(REQ_IF)
	mapOrigCopy[req_ifFrom] = req_ifTo
	req_ifFrom.GongCopyBasicFields(req_ifTo)

	//insertion point for the staging of instances referenced by pointers
	if req_ifFrom.THE_HEADER != nil {
		req_ifTo.THE_HEADER = GongCopyBranchA_THE_HEADER(mapOrigCopy, req_ifFrom.THE_HEADER)
	}
	if req_ifFrom.CORE_CONTENT != nil {
		req_ifTo.CORE_CONTENT = GongCopyBranchA_CORE_CONTENT(mapOrigCopy, req_ifFrom.CORE_CONTENT)
	}
	if req_ifFrom.TOOL_EXTENSIONS != nil {
		req_ifTo.TOOL_EXTENSIONS = GongCopyBranchA_TOOL_EXTENSIONS(mapOrigCopy, req_ifFrom.TOOL_EXTENSIONS)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchREQ_IF_CONTENT(mapOrigCopy map[any]any, req_if_contentFrom *REQ_IF_CONTENT) (req_if_contentTo *REQ_IF_CONTENT) {

	// req_if_contentFrom has already been copied
	if _req_if_contentTo, ok := mapOrigCopy[req_if_contentFrom]; ok {
		req_if_contentTo = _req_if_contentTo.(*REQ_IF_CONTENT)
		return
	}

	req_if_contentTo = new(REQ_IF_CONTENT)
	mapOrigCopy[req_if_contentFrom] = req_if_contentTo
	req_if_contentFrom.GongCopyBasicFields(req_if_contentTo)

	//insertion point for the staging of instances referenced by pointers
	if req_if_contentFrom.DATATYPES != nil {
		req_if_contentTo.DATATYPES = GongCopyBranchA_DATATYPES(mapOrigCopy, req_if_contentFrom.DATATYPES)
	}
	if req_if_contentFrom.SPEC_TYPES != nil {
		req_if_contentTo.SPEC_TYPES = GongCopyBranchA_SPEC_TYPES(mapOrigCopy, req_if_contentFrom.SPEC_TYPES)
	}
	if req_if_contentFrom.SPEC_OBJECTS != nil {
		req_if_contentTo.SPEC_OBJECTS = GongCopyBranchA_SPEC_OBJECTS(mapOrigCopy, req_if_contentFrom.SPEC_OBJECTS)
	}
	if req_if_contentFrom.SPEC_RELATIONS != nil {
		req_if_contentTo.SPEC_RELATIONS = GongCopyBranchA_SPEC_RELATIONS(mapOrigCopy, req_if_contentFrom.SPEC_RELATIONS)
	}
	if req_if_contentFrom.SPECIFICATIONS != nil {
		req_if_contentTo.SPECIFICATIONS = GongCopyBranchA_SPECIFICATIONS(mapOrigCopy, req_if_contentFrom.SPECIFICATIONS)
	}
	if req_if_contentFrom.SPEC_RELATION_GROUPS != nil {
		req_if_contentTo.SPEC_RELATION_GROUPS = GongCopyBranchA_SPEC_RELATION_GROUPS(mapOrigCopy, req_if_contentFrom.SPEC_RELATION_GROUPS)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchREQ_IF_HEADER(mapOrigCopy map[any]any, req_if_headerFrom *REQ_IF_HEADER) (req_if_headerTo *REQ_IF_HEADER) {

	// req_if_headerFrom has already been copied
	if _req_if_headerTo, ok := mapOrigCopy[req_if_headerFrom]; ok {
		req_if_headerTo = _req_if_headerTo.(*REQ_IF_HEADER)
		return
	}

	req_if_headerTo = new(REQ_IF_HEADER)
	mapOrigCopy[req_if_headerFrom] = req_if_headerTo
	req_if_headerFrom.GongCopyBasicFields(req_if_headerTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchREQ_IF_TOOL_EXTENSION(mapOrigCopy map[any]any, req_if_tool_extensionFrom *REQ_IF_TOOL_EXTENSION) (req_if_tool_extensionTo *REQ_IF_TOOL_EXTENSION) {

	// req_if_tool_extensionFrom has already been copied
	if _req_if_tool_extensionTo, ok := mapOrigCopy[req_if_tool_extensionFrom]; ok {
		req_if_tool_extensionTo = _req_if_tool_extensionTo.(*REQ_IF_TOOL_EXTENSION)
		return
	}

	req_if_tool_extensionTo = new(REQ_IF_TOOL_EXTENSION)
	mapOrigCopy[req_if_tool_extensionFrom] = req_if_tool_extensionTo
	req_if_tool_extensionFrom.GongCopyBasicFields(req_if_tool_extensionTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSPECIFICATION(mapOrigCopy map[any]any, specificationFrom *SPECIFICATION) (specificationTo *SPECIFICATION) {

	// specificationFrom has already been copied
	if _specificationTo, ok := mapOrigCopy[specificationFrom]; ok {
		specificationTo = _specificationTo.(*SPECIFICATION)
		return
	}

	specificationTo = new(SPECIFICATION)
	mapOrigCopy[specificationFrom] = specificationTo
	specificationFrom.GongCopyBasicFields(specificationTo)

	//insertion point for the staging of instances referenced by pointers
	if specificationFrom.ALTERNATIVE_ID != nil {
		specificationTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, specificationFrom.ALTERNATIVE_ID)
	}
	if specificationFrom.CHILDREN != nil {
		specificationTo.CHILDREN = GongCopyBranchA_CHILDREN(mapOrigCopy, specificationFrom.CHILDREN)
	}
	if specificationFrom.VALUES != nil {
		specificationTo.VALUES = GongCopyBranchA_ATTRIBUTE_VALUE_XHTML_1(mapOrigCopy, specificationFrom.VALUES)
	}
	if specificationFrom.TYPE != nil {
		specificationTo.TYPE = GongCopyBranchA_SPECIFICATION_TYPE_REF(mapOrigCopy, specificationFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSPECIFICATION_TYPE(mapOrigCopy map[any]any, specification_typeFrom *SPECIFICATION_TYPE) (specification_typeTo *SPECIFICATION_TYPE) {

	// specification_typeFrom has already been copied
	if _specification_typeTo, ok := mapOrigCopy[specification_typeFrom]; ok {
		specification_typeTo = _specification_typeTo.(*SPECIFICATION_TYPE)
		return
	}

	specification_typeTo = new(SPECIFICATION_TYPE)
	mapOrigCopy[specification_typeFrom] = specification_typeTo
	specification_typeFrom.GongCopyBasicFields(specification_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if specification_typeFrom.ALTERNATIVE_ID != nil {
		specification_typeTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, specification_typeFrom.ALTERNATIVE_ID)
	}
	if specification_typeFrom.SPEC_ATTRIBUTES != nil {
		specification_typeTo.SPEC_ATTRIBUTES = GongCopyBranchA_SPEC_ATTRIBUTES(mapOrigCopy, specification_typeFrom.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSPEC_HIERARCHY(mapOrigCopy map[any]any, spec_hierarchyFrom *SPEC_HIERARCHY) (spec_hierarchyTo *SPEC_HIERARCHY) {

	// spec_hierarchyFrom has already been copied
	if _spec_hierarchyTo, ok := mapOrigCopy[spec_hierarchyFrom]; ok {
		spec_hierarchyTo = _spec_hierarchyTo.(*SPEC_HIERARCHY)
		return
	}

	spec_hierarchyTo = new(SPEC_HIERARCHY)
	mapOrigCopy[spec_hierarchyFrom] = spec_hierarchyTo
	spec_hierarchyFrom.GongCopyBasicFields(spec_hierarchyTo)

	//insertion point for the staging of instances referenced by pointers
	if spec_hierarchyFrom.ALTERNATIVE_ID != nil {
		spec_hierarchyTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, spec_hierarchyFrom.ALTERNATIVE_ID)
	}
	if spec_hierarchyFrom.CHILDREN != nil {
		spec_hierarchyTo.CHILDREN = GongCopyBranchA_CHILDREN(mapOrigCopy, spec_hierarchyFrom.CHILDREN)
	}
	if spec_hierarchyFrom.EDITABLE_ATTS != nil {
		spec_hierarchyTo.EDITABLE_ATTS = GongCopyBranchA_EDITABLE_ATTS(mapOrigCopy, spec_hierarchyFrom.EDITABLE_ATTS)
	}
	if spec_hierarchyFrom.OBJECT != nil {
		spec_hierarchyTo.OBJECT = GongCopyBranchA_OBJECT(mapOrigCopy, spec_hierarchyFrom.OBJECT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSPEC_OBJECT(mapOrigCopy map[any]any, spec_objectFrom *SPEC_OBJECT) (spec_objectTo *SPEC_OBJECT) {

	// spec_objectFrom has already been copied
	if _spec_objectTo, ok := mapOrigCopy[spec_objectFrom]; ok {
		spec_objectTo = _spec_objectTo.(*SPEC_OBJECT)
		return
	}

	spec_objectTo = new(SPEC_OBJECT)
	mapOrigCopy[spec_objectFrom] = spec_objectTo
	spec_objectFrom.GongCopyBasicFields(spec_objectTo)

	//insertion point for the staging of instances referenced by pointers
	if spec_objectFrom.ALTERNATIVE_ID != nil {
		spec_objectTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, spec_objectFrom.ALTERNATIVE_ID)
	}
	if spec_objectFrom.VALUES != nil {
		spec_objectTo.VALUES = GongCopyBranchA_ATTRIBUTE_VALUE_XHTML_1(mapOrigCopy, spec_objectFrom.VALUES)
	}
	if spec_objectFrom.TYPE != nil {
		spec_objectTo.TYPE = GongCopyBranchA_SPEC_OBJECT_TYPE_REF(mapOrigCopy, spec_objectFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSPEC_OBJECT_TYPE(mapOrigCopy map[any]any, spec_object_typeFrom *SPEC_OBJECT_TYPE) (spec_object_typeTo *SPEC_OBJECT_TYPE) {

	// spec_object_typeFrom has already been copied
	if _spec_object_typeTo, ok := mapOrigCopy[spec_object_typeFrom]; ok {
		spec_object_typeTo = _spec_object_typeTo.(*SPEC_OBJECT_TYPE)
		return
	}

	spec_object_typeTo = new(SPEC_OBJECT_TYPE)
	mapOrigCopy[spec_object_typeFrom] = spec_object_typeTo
	spec_object_typeFrom.GongCopyBasicFields(spec_object_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if spec_object_typeFrom.ALTERNATIVE_ID != nil {
		spec_object_typeTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, spec_object_typeFrom.ALTERNATIVE_ID)
	}
	if spec_object_typeFrom.SPEC_ATTRIBUTES != nil {
		spec_object_typeTo.SPEC_ATTRIBUTES = GongCopyBranchA_SPEC_ATTRIBUTES(mapOrigCopy, spec_object_typeFrom.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSPEC_RELATION(mapOrigCopy map[any]any, spec_relationFrom *SPEC_RELATION) (spec_relationTo *SPEC_RELATION) {

	// spec_relationFrom has already been copied
	if _spec_relationTo, ok := mapOrigCopy[spec_relationFrom]; ok {
		spec_relationTo = _spec_relationTo.(*SPEC_RELATION)
		return
	}

	spec_relationTo = new(SPEC_RELATION)
	mapOrigCopy[spec_relationFrom] = spec_relationTo
	spec_relationFrom.GongCopyBasicFields(spec_relationTo)

	//insertion point for the staging of instances referenced by pointers
	if spec_relationFrom.ALTERNATIVE_ID != nil {
		spec_relationTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, spec_relationFrom.ALTERNATIVE_ID)
	}
	if spec_relationFrom.VALUES != nil {
		spec_relationTo.VALUES = GongCopyBranchA_ATTRIBUTE_VALUE_XHTML_1(mapOrigCopy, spec_relationFrom.VALUES)
	}
	if spec_relationFrom.SOURCE != nil {
		spec_relationTo.SOURCE = GongCopyBranchA_SOURCE_1(mapOrigCopy, spec_relationFrom.SOURCE)
	}
	if spec_relationFrom.TARGET != nil {
		spec_relationTo.TARGET = GongCopyBranchA_SOURCE_1(mapOrigCopy, spec_relationFrom.TARGET)
	}
	if spec_relationFrom.TYPE != nil {
		spec_relationTo.TYPE = GongCopyBranchA_SPEC_RELATION_TYPE_REF(mapOrigCopy, spec_relationFrom.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchSPEC_RELATION_TYPE(mapOrigCopy map[any]any, spec_relation_typeFrom *SPEC_RELATION_TYPE) (spec_relation_typeTo *SPEC_RELATION_TYPE) {

	// spec_relation_typeFrom has already been copied
	if _spec_relation_typeTo, ok := mapOrigCopy[spec_relation_typeFrom]; ok {
		spec_relation_typeTo = _spec_relation_typeTo.(*SPEC_RELATION_TYPE)
		return
	}

	spec_relation_typeTo = new(SPEC_RELATION_TYPE)
	mapOrigCopy[spec_relation_typeFrom] = spec_relation_typeTo
	spec_relation_typeFrom.GongCopyBasicFields(spec_relation_typeTo)

	//insertion point for the staging of instances referenced by pointers
	if spec_relation_typeFrom.ALTERNATIVE_ID != nil {
		spec_relation_typeTo.ALTERNATIVE_ID = GongCopyBranchA_ALTERNATIVE_ID(mapOrigCopy, spec_relation_typeFrom.ALTERNATIVE_ID)
	}
	if spec_relation_typeFrom.SPEC_ATTRIBUTES != nil {
		spec_relation_typeTo.SPEC_ATTRIBUTES = GongCopyBranchA_SPEC_ATTRIBUTES(mapOrigCopy, spec_relation_typeFrom.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

func GongCopyBranchXHTML_CONTENT(mapOrigCopy map[any]any, xhtml_contentFrom *XHTML_CONTENT) (xhtml_contentTo *XHTML_CONTENT) {

	// xhtml_contentFrom has already been copied
	if _xhtml_contentTo, ok := mapOrigCopy[xhtml_contentFrom]; ok {
		xhtml_contentTo = _xhtml_contentTo.(*XHTML_CONTENT)
		return
	}

	xhtml_contentTo = new(XHTML_CONTENT)
	mapOrigCopy[xhtml_contentFrom] = xhtml_contentTo
	xhtml_contentFrom.GongCopyBasicFields(xhtml_contentTo)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

	return
}

// UnstageBranch stages instance and apply UnstageBranch on all gongstruct instances that are
// referenced by pointers or slices of pointers of the insance
//
// the algorithm stops along the course of graph if a vertex is already staged
// UnstageBranch is the Stage method that unstages instance and applies UnstageBranch recursively.
func (stage *Stage) UnstageBranch(instance GongstructIF) {
	if instance != nil {
		instance.GongUnstageBranch(stage)
	}
}

// UnstageBranch is a backward-compatible package-level forwarder.
func UnstageBranch(stage *Stage, instance GongstructIF) {
	stage.UnstageBranch(instance)
}

// insertion point for unstage branch per struct
func (alternative_id *ALTERNATIVE_ID) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchALTERNATIVE_ID(alternative_id)
}

func (stage *Stage) UnstageBranchALTERNATIVE_ID(alternative_id *ALTERNATIVE_ID) {

	// check if instance is already staged
	if !stage.IsStaged(alternative_id) {
		return
	}

	alternative_id.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean)
}

func (stage *Stage) UnstageBranchATTRIBUTE_DEFINITION_BOOLEAN(attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) {

	// check if instance is already staged
	if !stage.IsStaged(attribute_definition_boolean) {
		return
	}

	attribute_definition_boolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_boolean.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(attribute_definition_boolean.ALTERNATIVE_ID)
	}
	if attribute_definition_boolean.DEFAULT_VALUE != nil {
		stage.UnstageBranch(attribute_definition_boolean.DEFAULT_VALUE)
	}
	if attribute_definition_boolean.TYPE != nil {
		stage.UnstageBranch(attribute_definition_boolean.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchATTRIBUTE_DEFINITION_DATE(attribute_definition_date)
}

func (stage *Stage) UnstageBranchATTRIBUTE_DEFINITION_DATE(attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) {

	// check if instance is already staged
	if !stage.IsStaged(attribute_definition_date) {
		return
	}

	attribute_definition_date.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_date.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(attribute_definition_date.ALTERNATIVE_ID)
	}
	if attribute_definition_date.DEFAULT_VALUE != nil {
		stage.UnstageBranch(attribute_definition_date.DEFAULT_VALUE)
	}
	if attribute_definition_date.TYPE != nil {
		stage.UnstageBranch(attribute_definition_date.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration)
}

func (stage *Stage) UnstageBranchATTRIBUTE_DEFINITION_ENUMERATION(attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) {

	// check if instance is already staged
	if !stage.IsStaged(attribute_definition_enumeration) {
		return
	}

	attribute_definition_enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_enumeration.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(attribute_definition_enumeration.ALTERNATIVE_ID)
	}
	if attribute_definition_enumeration.DEFAULT_VALUE != nil {
		stage.UnstageBranch(attribute_definition_enumeration.DEFAULT_VALUE)
	}
	if attribute_definition_enumeration.TYPE != nil {
		stage.UnstageBranch(attribute_definition_enumeration.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer)
}

func (stage *Stage) UnstageBranchATTRIBUTE_DEFINITION_INTEGER(attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) {

	// check if instance is already staged
	if !stage.IsStaged(attribute_definition_integer) {
		return
	}

	attribute_definition_integer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_integer.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(attribute_definition_integer.ALTERNATIVE_ID)
	}
	if attribute_definition_integer.DEFAULT_VALUE != nil {
		stage.UnstageBranch(attribute_definition_integer.DEFAULT_VALUE)
	}
	if attribute_definition_integer.TYPE != nil {
		stage.UnstageBranch(attribute_definition_integer.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchATTRIBUTE_DEFINITION_REAL(attribute_definition_real)
}

func (stage *Stage) UnstageBranchATTRIBUTE_DEFINITION_REAL(attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) {

	// check if instance is already staged
	if !stage.IsStaged(attribute_definition_real) {
		return
	}

	attribute_definition_real.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_real.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(attribute_definition_real.ALTERNATIVE_ID)
	}
	if attribute_definition_real.DEFAULT_VALUE != nil {
		stage.UnstageBranch(attribute_definition_real.DEFAULT_VALUE)
	}
	if attribute_definition_real.TYPE != nil {
		stage.UnstageBranch(attribute_definition_real.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchATTRIBUTE_DEFINITION_STRING(attribute_definition_string)
}

func (stage *Stage) UnstageBranchATTRIBUTE_DEFINITION_STRING(attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) {

	// check if instance is already staged
	if !stage.IsStaged(attribute_definition_string) {
		return
	}

	attribute_definition_string.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_string.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(attribute_definition_string.ALTERNATIVE_ID)
	}
	if attribute_definition_string.DEFAULT_VALUE != nil {
		stage.UnstageBranch(attribute_definition_string.DEFAULT_VALUE)
	}
	if attribute_definition_string.TYPE != nil {
		stage.UnstageBranch(attribute_definition_string.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml)
}

func (stage *Stage) UnstageBranchATTRIBUTE_DEFINITION_XHTML(attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) {

	// check if instance is already staged
	if !stage.IsStaged(attribute_definition_xhtml) {
		return
	}

	attribute_definition_xhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_definition_xhtml.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(attribute_definition_xhtml.ALTERNATIVE_ID)
	}
	if attribute_definition_xhtml.DEFAULT_VALUE != nil {
		stage.UnstageBranch(attribute_definition_xhtml.DEFAULT_VALUE)
	}
	if attribute_definition_xhtml.TYPE != nil {
		stage.UnstageBranch(attribute_definition_xhtml.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean)
}

func (stage *Stage) UnstageBranchATTRIBUTE_VALUE_BOOLEAN(attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) {

	// check if instance is already staged
	if !stage.IsStaged(attribute_value_boolean) {
		return
	}

	attribute_value_boolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_boolean.DEFINITION != nil {
		stage.UnstageBranch(attribute_value_boolean.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchATTRIBUTE_VALUE_DATE(attribute_value_date)
}

func (stage *Stage) UnstageBranchATTRIBUTE_VALUE_DATE(attribute_value_date *ATTRIBUTE_VALUE_DATE) {

	// check if instance is already staged
	if !stage.IsStaged(attribute_value_date) {
		return
	}

	attribute_value_date.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_date.DEFINITION != nil {
		stage.UnstageBranch(attribute_value_date.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration)
}

func (stage *Stage) UnstageBranchATTRIBUTE_VALUE_ENUMERATION(attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) {

	// check if instance is already staged
	if !stage.IsStaged(attribute_value_enumeration) {
		return
	}

	attribute_value_enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_enumeration.DEFINITION != nil {
		stage.UnstageBranch(attribute_value_enumeration.DEFINITION)
	}
	if attribute_value_enumeration.VALUES != nil {
		stage.UnstageBranch(attribute_value_enumeration.VALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchATTRIBUTE_VALUE_INTEGER(attribute_value_integer)
}

func (stage *Stage) UnstageBranchATTRIBUTE_VALUE_INTEGER(attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) {

	// check if instance is already staged
	if !stage.IsStaged(attribute_value_integer) {
		return
	}

	attribute_value_integer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_integer.DEFINITION != nil {
		stage.UnstageBranch(attribute_value_integer.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchATTRIBUTE_VALUE_REAL(attribute_value_real)
}

func (stage *Stage) UnstageBranchATTRIBUTE_VALUE_REAL(attribute_value_real *ATTRIBUTE_VALUE_REAL) {

	// check if instance is already staged
	if !stage.IsStaged(attribute_value_real) {
		return
	}

	attribute_value_real.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_real.DEFINITION != nil {
		stage.UnstageBranch(attribute_value_real.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchATTRIBUTE_VALUE_STRING(attribute_value_string)
}

func (stage *Stage) UnstageBranchATTRIBUTE_VALUE_STRING(attribute_value_string *ATTRIBUTE_VALUE_STRING) {

	// check if instance is already staged
	if !stage.IsStaged(attribute_value_string) {
		return
	}

	attribute_value_string.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_string.DEFINITION != nil {
		stage.UnstageBranch(attribute_value_string.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchATTRIBUTE_VALUE_XHTML(attribute_value_xhtml)
}

func (stage *Stage) UnstageBranchATTRIBUTE_VALUE_XHTML(attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) {

	// check if instance is already staged
	if !stage.IsStaged(attribute_value_xhtml) {
		return
	}

	attribute_value_xhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if attribute_value_xhtml.THE_VALUE != nil {
		stage.UnstageBranch(attribute_value_xhtml.THE_VALUE)
	}
	if attribute_value_xhtml.THE_ORIGINAL_VALUE != nil {
		stage.UnstageBranch(attribute_value_xhtml.THE_ORIGINAL_VALUE)
	}
	if attribute_value_xhtml.DEFINITION != nil {
		stage.UnstageBranch(attribute_value_xhtml.DEFINITION)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_alternative_id *A_ALTERNATIVE_ID) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ALTERNATIVE_ID(a_alternative_id)
}

func (stage *Stage) UnstageBranchA_ALTERNATIVE_ID(a_alternative_id *A_ALTERNATIVE_ID) {

	// check if instance is already staged
	if !stage.IsStaged(a_alternative_id) {
		return
	}

	a_alternative_id.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_alternative_id.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(a_alternative_id.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(a_attribute_definition_boolean_ref)
}

func (stage *Stage) UnstageBranchA_ATTRIBUTE_DEFINITION_BOOLEAN_REF(a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_attribute_definition_boolean_ref) {
		return
	}

	a_attribute_definition_boolean_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ATTRIBUTE_DEFINITION_DATE_REF(a_attribute_definition_date_ref)
}

func (stage *Stage) UnstageBranchA_ATTRIBUTE_DEFINITION_DATE_REF(a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_attribute_definition_date_ref) {
		return
	}

	a_attribute_definition_date_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(a_attribute_definition_enumeration_ref)
}

func (stage *Stage) UnstageBranchA_ATTRIBUTE_DEFINITION_ENUMERATION_REF(a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_attribute_definition_enumeration_ref) {
		return
	}

	a_attribute_definition_enumeration_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ATTRIBUTE_DEFINITION_INTEGER_REF(a_attribute_definition_integer_ref)
}

func (stage *Stage) UnstageBranchA_ATTRIBUTE_DEFINITION_INTEGER_REF(a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_attribute_definition_integer_ref) {
		return
	}

	a_attribute_definition_integer_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ATTRIBUTE_DEFINITION_REAL_REF(a_attribute_definition_real_ref)
}

func (stage *Stage) UnstageBranchA_ATTRIBUTE_DEFINITION_REAL_REF(a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_attribute_definition_real_ref) {
		return
	}

	a_attribute_definition_real_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ATTRIBUTE_DEFINITION_STRING_REF(a_attribute_definition_string_ref)
}

func (stage *Stage) UnstageBranchA_ATTRIBUTE_DEFINITION_STRING_REF(a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_attribute_definition_string_ref) {
		return
	}

	a_attribute_definition_string_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ATTRIBUTE_DEFINITION_XHTML_REF(a_attribute_definition_xhtml_ref)
}

func (stage *Stage) UnstageBranchA_ATTRIBUTE_DEFINITION_XHTML_REF(a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_attribute_definition_xhtml_ref) {
		return
	}

	a_attribute_definition_xhtml_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ATTRIBUTE_VALUE_BOOLEAN(a_attribute_value_boolean)
}

func (stage *Stage) UnstageBranchA_ATTRIBUTE_VALUE_BOOLEAN(a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) {

	// check if instance is already staged
	if !stage.IsStaged(a_attribute_value_boolean) {
		return
	}

	a_attribute_value_boolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_boolean := range a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN {
		stage.UnstageBranch(_attribute_value_boolean)
	}

}

func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ATTRIBUTE_VALUE_DATE(a_attribute_value_date)
}

func (stage *Stage) UnstageBranchA_ATTRIBUTE_VALUE_DATE(a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) {

	// check if instance is already staged
	if !stage.IsStaged(a_attribute_value_date) {
		return
	}

	a_attribute_value_date.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_date := range a_attribute_value_date.ATTRIBUTE_VALUE_DATE {
		stage.UnstageBranch(_attribute_value_date)
	}

}

func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration)
}

func (stage *Stage) UnstageBranchA_ATTRIBUTE_VALUE_ENUMERATION(a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) {

	// check if instance is already staged
	if !stage.IsStaged(a_attribute_value_enumeration) {
		return
	}

	a_attribute_value_enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_enumeration := range a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION {
		stage.UnstageBranch(_attribute_value_enumeration)
	}

}

func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ATTRIBUTE_VALUE_INTEGER(a_attribute_value_integer)
}

func (stage *Stage) UnstageBranchA_ATTRIBUTE_VALUE_INTEGER(a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) {

	// check if instance is already staged
	if !stage.IsStaged(a_attribute_value_integer) {
		return
	}

	a_attribute_value_integer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_integer := range a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER {
		stage.UnstageBranch(_attribute_value_integer)
	}

}

func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ATTRIBUTE_VALUE_REAL(a_attribute_value_real)
}

func (stage *Stage) UnstageBranchA_ATTRIBUTE_VALUE_REAL(a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) {

	// check if instance is already staged
	if !stage.IsStaged(a_attribute_value_real) {
		return
	}

	a_attribute_value_real.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_real := range a_attribute_value_real.ATTRIBUTE_VALUE_REAL {
		stage.UnstageBranch(_attribute_value_real)
	}

}

func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ATTRIBUTE_VALUE_STRING(a_attribute_value_string)
}

func (stage *Stage) UnstageBranchA_ATTRIBUTE_VALUE_STRING(a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) {

	// check if instance is already staged
	if !stage.IsStaged(a_attribute_value_string) {
		return
	}

	a_attribute_value_string.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_string := range a_attribute_value_string.ATTRIBUTE_VALUE_STRING {
		stage.UnstageBranch(_attribute_value_string)
	}

}

func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ATTRIBUTE_VALUE_XHTML(a_attribute_value_xhtml)
}

func (stage *Stage) UnstageBranchA_ATTRIBUTE_VALUE_XHTML(a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) {

	// check if instance is already staged
	if !stage.IsStaged(a_attribute_value_xhtml) {
		return
	}

	a_attribute_value_xhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_xhtml := range a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML {
		stage.UnstageBranch(_attribute_value_xhtml)
	}

}

func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ATTRIBUTE_VALUE_XHTML_1(a_attribute_value_xhtml_1)
}

func (stage *Stage) UnstageBranchA_ATTRIBUTE_VALUE_XHTML_1(a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) {

	// check if instance is already staged
	if !stage.IsStaged(a_attribute_value_xhtml_1) {
		return
	}

	a_attribute_value_xhtml_1.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_value_boolean := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN {
		stage.UnstageBranch(_attribute_value_boolean)
	}
	for _, _attribute_value_date := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE {
		stage.UnstageBranch(_attribute_value_date)
	}
	for _, _attribute_value_enumeration := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION {
		stage.UnstageBranch(_attribute_value_enumeration)
	}
	for _, _attribute_value_integer := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER {
		stage.UnstageBranch(_attribute_value_integer)
	}
	for _, _attribute_value_real := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL {
		stage.UnstageBranch(_attribute_value_real)
	}
	for _, _attribute_value_string := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING {
		stage.UnstageBranch(_attribute_value_string)
	}
	for _, _attribute_value_xhtml := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML {
		stage.UnstageBranch(_attribute_value_xhtml)
	}

}

func (a_children *A_CHILDREN) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_CHILDREN(a_children)
}

func (stage *Stage) UnstageBranchA_CHILDREN(a_children *A_CHILDREN) {

	// check if instance is already staged
	if !stage.IsStaged(a_children) {
		return
	}

	a_children.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_hierarchy := range a_children.SPEC_HIERARCHY {
		stage.UnstageBranch(_spec_hierarchy)
	}

}

func (a_core_content *A_CORE_CONTENT) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_CORE_CONTENT(a_core_content)
}

func (stage *Stage) UnstageBranchA_CORE_CONTENT(a_core_content *A_CORE_CONTENT) {

	// check if instance is already staged
	if !stage.IsStaged(a_core_content) {
		return
	}

	a_core_content.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_core_content.REQ_IF_CONTENT != nil {
		stage.UnstageBranch(a_core_content.REQ_IF_CONTENT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_datatypes *A_DATATYPES) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_DATATYPES(a_datatypes)
}

func (stage *Stage) UnstageBranchA_DATATYPES(a_datatypes *A_DATATYPES) {

	// check if instance is already staged
	if !stage.IsStaged(a_datatypes) {
		return
	}

	a_datatypes.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _datatype_definition_boolean := range a_datatypes.DATATYPE_DEFINITION_BOOLEAN {
		stage.UnstageBranch(_datatype_definition_boolean)
	}
	for _, _datatype_definition_date := range a_datatypes.DATATYPE_DEFINITION_DATE {
		stage.UnstageBranch(_datatype_definition_date)
	}
	for _, _datatype_definition_enumeration := range a_datatypes.DATATYPE_DEFINITION_ENUMERATION {
		stage.UnstageBranch(_datatype_definition_enumeration)
	}
	for _, _datatype_definition_integer := range a_datatypes.DATATYPE_DEFINITION_INTEGER {
		stage.UnstageBranch(_datatype_definition_integer)
	}
	for _, _datatype_definition_real := range a_datatypes.DATATYPE_DEFINITION_REAL {
		stage.UnstageBranch(_datatype_definition_real)
	}
	for _, _datatype_definition_string := range a_datatypes.DATATYPE_DEFINITION_STRING {
		stage.UnstageBranch(_datatype_definition_string)
	}
	for _, _datatype_definition_xhtml := range a_datatypes.DATATYPE_DEFINITION_XHTML {
		stage.UnstageBranch(_datatype_definition_xhtml)
	}

}

func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_DATATYPE_DEFINITION_BOOLEAN_REF(a_datatype_definition_boolean_ref)
}

func (stage *Stage) UnstageBranchA_DATATYPE_DEFINITION_BOOLEAN_REF(a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_datatype_definition_boolean_ref) {
		return
	}

	a_datatype_definition_boolean_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_DATATYPE_DEFINITION_DATE_REF(a_datatype_definition_date_ref)
}

func (stage *Stage) UnstageBranchA_DATATYPE_DEFINITION_DATE_REF(a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_datatype_definition_date_ref) {
		return
	}

	a_datatype_definition_date_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_DATATYPE_DEFINITION_ENUMERATION_REF(a_datatype_definition_enumeration_ref)
}

func (stage *Stage) UnstageBranchA_DATATYPE_DEFINITION_ENUMERATION_REF(a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_datatype_definition_enumeration_ref) {
		return
	}

	a_datatype_definition_enumeration_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_DATATYPE_DEFINITION_INTEGER_REF(a_datatype_definition_integer_ref)
}

func (stage *Stage) UnstageBranchA_DATATYPE_DEFINITION_INTEGER_REF(a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_datatype_definition_integer_ref) {
		return
	}

	a_datatype_definition_integer_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_DATATYPE_DEFINITION_REAL_REF(a_datatype_definition_real_ref)
}

func (stage *Stage) UnstageBranchA_DATATYPE_DEFINITION_REAL_REF(a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_datatype_definition_real_ref) {
		return
	}

	a_datatype_definition_real_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_DATATYPE_DEFINITION_STRING_REF(a_datatype_definition_string_ref)
}

func (stage *Stage) UnstageBranchA_DATATYPE_DEFINITION_STRING_REF(a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_datatype_definition_string_ref) {
		return
	}

	a_datatype_definition_string_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_DATATYPE_DEFINITION_XHTML_REF(a_datatype_definition_xhtml_ref)
}

func (stage *Stage) UnstageBranchA_DATATYPE_DEFINITION_XHTML_REF(a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_datatype_definition_xhtml_ref) {
		return
	}

	a_datatype_definition_xhtml_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_editable_atts *A_EDITABLE_ATTS) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_EDITABLE_ATTS(a_editable_atts)
}

func (stage *Stage) UnstageBranchA_EDITABLE_ATTS(a_editable_atts *A_EDITABLE_ATTS) {

	// check if instance is already staged
	if !stage.IsStaged(a_editable_atts) {
		return
	}

	a_editable_atts.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_enum_value_ref *A_ENUM_VALUE_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_ENUM_VALUE_REF(a_enum_value_ref)
}

func (stage *Stage) UnstageBranchA_ENUM_VALUE_REF(a_enum_value_ref *A_ENUM_VALUE_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_enum_value_ref) {
		return
	}

	a_enum_value_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_object *A_OBJECT) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_OBJECT(a_object)
}

func (stage *Stage) UnstageBranchA_OBJECT(a_object *A_OBJECT) {

	// check if instance is already staged
	if !stage.IsStaged(a_object) {
		return
	}

	a_object.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_properties *A_PROPERTIES) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_PROPERTIES(a_properties)
}

func (stage *Stage) UnstageBranchA_PROPERTIES(a_properties *A_PROPERTIES) {

	// check if instance is already staged
	if !stage.IsStaged(a_properties) {
		return
	}

	a_properties.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_properties.EMBEDDED_VALUE != nil {
		stage.UnstageBranch(a_properties.EMBEDDED_VALUE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref)
}

func (stage *Stage) UnstageBranchA_RELATION_GROUP_TYPE_REF(a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_relation_group_type_ref) {
		return
	}

	a_relation_group_type_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_source_1 *A_SOURCE_1) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_SOURCE_1(a_source_1)
}

func (stage *Stage) UnstageBranchA_SOURCE_1(a_source_1 *A_SOURCE_1) {

	// check if instance is already staged
	if !stage.IsStaged(a_source_1) {
		return
	}

	a_source_1.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_SOURCE_SPECIFICATION_1(a_source_specification_1)
}

func (stage *Stage) UnstageBranchA_SOURCE_SPECIFICATION_1(a_source_specification_1 *A_SOURCE_SPECIFICATION_1) {

	// check if instance is already staged
	if !stage.IsStaged(a_source_specification_1) {
		return
	}

	a_source_specification_1.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_specifications *A_SPECIFICATIONS) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_SPECIFICATIONS(a_specifications)
}

func (stage *Stage) UnstageBranchA_SPECIFICATIONS(a_specifications *A_SPECIFICATIONS) {

	// check if instance is already staged
	if !stage.IsStaged(a_specifications) {
		return
	}

	a_specifications.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _specification := range a_specifications.SPECIFICATION {
		stage.UnstageBranch(_specification)
	}

}

func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_SPECIFICATION_TYPE_REF(a_specification_type_ref)
}

func (stage *Stage) UnstageBranchA_SPECIFICATION_TYPE_REF(a_specification_type_ref *A_SPECIFICATION_TYPE_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_specification_type_ref) {
		return
	}

	a_specification_type_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_specified_values *A_SPECIFIED_VALUES) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_SPECIFIED_VALUES(a_specified_values)
}

func (stage *Stage) UnstageBranchA_SPECIFIED_VALUES(a_specified_values *A_SPECIFIED_VALUES) {

	// check if instance is already staged
	if !stage.IsStaged(a_specified_values) {
		return
	}

	a_specified_values.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _enum_value := range a_specified_values.ENUM_VALUE {
		stage.UnstageBranch(_enum_value)
	}

}

func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_SPEC_ATTRIBUTES(a_spec_attributes)
}

func (stage *Stage) UnstageBranchA_SPEC_ATTRIBUTES(a_spec_attributes *A_SPEC_ATTRIBUTES) {

	// check if instance is already staged
	if !stage.IsStaged(a_spec_attributes) {
		return
	}

	a_spec_attributes.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _attribute_definition_boolean := range a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN {
		stage.UnstageBranch(_attribute_definition_boolean)
	}
	for _, _attribute_definition_date := range a_spec_attributes.ATTRIBUTE_DEFINITION_DATE {
		stage.UnstageBranch(_attribute_definition_date)
	}
	for _, _attribute_definition_enumeration := range a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION {
		stage.UnstageBranch(_attribute_definition_enumeration)
	}
	for _, _attribute_definition_integer := range a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER {
		stage.UnstageBranch(_attribute_definition_integer)
	}
	for _, _attribute_definition_real := range a_spec_attributes.ATTRIBUTE_DEFINITION_REAL {
		stage.UnstageBranch(_attribute_definition_real)
	}
	for _, _attribute_definition_string := range a_spec_attributes.ATTRIBUTE_DEFINITION_STRING {
		stage.UnstageBranch(_attribute_definition_string)
	}
	for _, _attribute_definition_xhtml := range a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML {
		stage.UnstageBranch(_attribute_definition_xhtml)
	}

}

func (a_spec_objects *A_SPEC_OBJECTS) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_SPEC_OBJECTS(a_spec_objects)
}

func (stage *Stage) UnstageBranchA_SPEC_OBJECTS(a_spec_objects *A_SPEC_OBJECTS) {

	// check if instance is already staged
	if !stage.IsStaged(a_spec_objects) {
		return
	}

	a_spec_objects.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_object := range a_spec_objects.SPEC_OBJECT {
		stage.UnstageBranch(_spec_object)
	}

}

func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_SPEC_OBJECT_TYPE_REF(a_spec_object_type_ref)
}

func (stage *Stage) UnstageBranchA_SPEC_OBJECT_TYPE_REF(a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_spec_object_type_ref) {
		return
	}

	a_spec_object_type_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_spec_relations *A_SPEC_RELATIONS) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_SPEC_RELATIONS(a_spec_relations)
}

func (stage *Stage) UnstageBranchA_SPEC_RELATIONS(a_spec_relations *A_SPEC_RELATIONS) {

	// check if instance is already staged
	if !stage.IsStaged(a_spec_relations) {
		return
	}

	a_spec_relations.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _spec_relation := range a_spec_relations.SPEC_RELATION {
		stage.UnstageBranch(_spec_relation)
	}

}

func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_SPEC_RELATION_GROUPS(a_spec_relation_groups)
}

func (stage *Stage) UnstageBranchA_SPEC_RELATION_GROUPS(a_spec_relation_groups *A_SPEC_RELATION_GROUPS) {

	// check if instance is already staged
	if !stage.IsStaged(a_spec_relation_groups) {
		return
	}

	a_spec_relation_groups.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relation_group := range a_spec_relation_groups.RELATION_GROUP {
		stage.UnstageBranch(_relation_group)
	}

}

func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_SPEC_RELATION_REF(a_spec_relation_ref)
}

func (stage *Stage) UnstageBranchA_SPEC_RELATION_REF(a_spec_relation_ref *A_SPEC_RELATION_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_spec_relation_ref) {
		return
	}

	a_spec_relation_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_SPEC_RELATION_TYPE_REF(a_spec_relation_type_ref)
}

func (stage *Stage) UnstageBranchA_SPEC_RELATION_TYPE_REF(a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) {

	// check if instance is already staged
	if !stage.IsStaged(a_spec_relation_type_ref) {
		return
	}

	a_spec_relation_type_ref.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_spec_types *A_SPEC_TYPES) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_SPEC_TYPES(a_spec_types)
}

func (stage *Stage) UnstageBranchA_SPEC_TYPES(a_spec_types *A_SPEC_TYPES) {

	// check if instance is already staged
	if !stage.IsStaged(a_spec_types) {
		return
	}

	a_spec_types.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _relation_group_type := range a_spec_types.RELATION_GROUP_TYPE {
		stage.UnstageBranch(_relation_group_type)
	}
	for _, _spec_object_type := range a_spec_types.SPEC_OBJECT_TYPE {
		stage.UnstageBranch(_spec_object_type)
	}
	for _, _spec_relation_type := range a_spec_types.SPEC_RELATION_TYPE {
		stage.UnstageBranch(_spec_relation_type)
	}
	for _, _specification_type := range a_spec_types.SPECIFICATION_TYPE {
		stage.UnstageBranch(_specification_type)
	}

}

func (a_the_header *A_THE_HEADER) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_THE_HEADER(a_the_header)
}

func (stage *Stage) UnstageBranchA_THE_HEADER(a_the_header *A_THE_HEADER) {

	// check if instance is already staged
	if !stage.IsStaged(a_the_header) {
		return
	}

	a_the_header.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if a_the_header.REQ_IF_HEADER != nil {
		stage.UnstageBranch(a_the_header.REQ_IF_HEADER)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (a_tool_extensions *A_TOOL_EXTENSIONS) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchA_TOOL_EXTENSIONS(a_tool_extensions)
}

func (stage *Stage) UnstageBranchA_TOOL_EXTENSIONS(a_tool_extensions *A_TOOL_EXTENSIONS) {

	// check if instance is already staged
	if !stage.IsStaged(a_tool_extensions) {
		return
	}

	a_tool_extensions.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers
	for _, _req_if_tool_extension := range a_tool_extensions.REQ_IF_TOOL_EXTENSION {
		stage.UnstageBranch(_req_if_tool_extension)
	}

}

func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean)
}

func (stage *Stage) UnstageBranchDATATYPE_DEFINITION_BOOLEAN(datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) {

	// check if instance is already staged
	if !stage.IsStaged(datatype_definition_boolean) {
		return
	}

	datatype_definition_boolean.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_boolean.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(datatype_definition_boolean.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDATATYPE_DEFINITION_DATE(datatype_definition_date)
}

func (stage *Stage) UnstageBranchDATATYPE_DEFINITION_DATE(datatype_definition_date *DATATYPE_DEFINITION_DATE) {

	// check if instance is already staged
	if !stage.IsStaged(datatype_definition_date) {
		return
	}

	datatype_definition_date.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_date.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(datatype_definition_date.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration)
}

func (stage *Stage) UnstageBranchDATATYPE_DEFINITION_ENUMERATION(datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) {

	// check if instance is already staged
	if !stage.IsStaged(datatype_definition_enumeration) {
		return
	}

	datatype_definition_enumeration.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_enumeration.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(datatype_definition_enumeration.ALTERNATIVE_ID)
	}
	if datatype_definition_enumeration.SPECIFIED_VALUES != nil {
		stage.UnstageBranch(datatype_definition_enumeration.SPECIFIED_VALUES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDATATYPE_DEFINITION_INTEGER(datatype_definition_integer)
}

func (stage *Stage) UnstageBranchDATATYPE_DEFINITION_INTEGER(datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) {

	// check if instance is already staged
	if !stage.IsStaged(datatype_definition_integer) {
		return
	}

	datatype_definition_integer.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_integer.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(datatype_definition_integer.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDATATYPE_DEFINITION_REAL(datatype_definition_real)
}

func (stage *Stage) UnstageBranchDATATYPE_DEFINITION_REAL(datatype_definition_real *DATATYPE_DEFINITION_REAL) {

	// check if instance is already staged
	if !stage.IsStaged(datatype_definition_real) {
		return
	}

	datatype_definition_real.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_real.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(datatype_definition_real.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDATATYPE_DEFINITION_STRING(datatype_definition_string)
}

func (stage *Stage) UnstageBranchDATATYPE_DEFINITION_STRING(datatype_definition_string *DATATYPE_DEFINITION_STRING) {

	// check if instance is already staged
	if !stage.IsStaged(datatype_definition_string) {
		return
	}

	datatype_definition_string.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_string.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(datatype_definition_string.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml)
}

func (stage *Stage) UnstageBranchDATATYPE_DEFINITION_XHTML(datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) {

	// check if instance is already staged
	if !stage.IsStaged(datatype_definition_xhtml) {
		return
	}

	datatype_definition_xhtml.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if datatype_definition_xhtml.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(datatype_definition_xhtml.ALTERNATIVE_ID)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (embedded_value *EMBEDDED_VALUE) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchEMBEDDED_VALUE(embedded_value)
}

func (stage *Stage) UnstageBranchEMBEDDED_VALUE(embedded_value *EMBEDDED_VALUE) {

	// check if instance is already staged
	if !stage.IsStaged(embedded_value) {
		return
	}

	embedded_value.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (enum_value *ENUM_VALUE) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchENUM_VALUE(enum_value)
}

func (stage *Stage) UnstageBranchENUM_VALUE(enum_value *ENUM_VALUE) {

	// check if instance is already staged
	if !stage.IsStaged(enum_value) {
		return
	}

	enum_value.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if enum_value.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(enum_value.ALTERNATIVE_ID)
	}
	if enum_value.PROPERTIES != nil {
		stage.UnstageBranch(enum_value.PROPERTIES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (relation_group *RELATION_GROUP) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchRELATION_GROUP(relation_group)
}

func (stage *Stage) UnstageBranchRELATION_GROUP(relation_group *RELATION_GROUP) {

	// check if instance is already staged
	if !stage.IsStaged(relation_group) {
		return
	}

	relation_group.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if relation_group.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(relation_group.ALTERNATIVE_ID)
	}
	if relation_group.SOURCE_SPECIFICATION != nil {
		stage.UnstageBranch(relation_group.SOURCE_SPECIFICATION)
	}
	if relation_group.SPEC_RELATIONS != nil {
		stage.UnstageBranch(relation_group.SPEC_RELATIONS)
	}
	if relation_group.TARGET_SPECIFICATION != nil {
		stage.UnstageBranch(relation_group.TARGET_SPECIFICATION)
	}
	if relation_group.TYPE != nil {
		stage.UnstageBranch(relation_group.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (relation_group_type *RELATION_GROUP_TYPE) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchRELATION_GROUP_TYPE(relation_group_type)
}

func (stage *Stage) UnstageBranchRELATION_GROUP_TYPE(relation_group_type *RELATION_GROUP_TYPE) {

	// check if instance is already staged
	if !stage.IsStaged(relation_group_type) {
		return
	}

	relation_group_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if relation_group_type.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(relation_group_type.ALTERNATIVE_ID)
	}
	if relation_group_type.SPEC_ATTRIBUTES != nil {
		stage.UnstageBranch(relation_group_type.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (req_if *REQ_IF) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchREQ_IF(req_if)
}

func (stage *Stage) UnstageBranchREQ_IF(req_if *REQ_IF) {

	// check if instance is already staged
	if !stage.IsStaged(req_if) {
		return
	}

	req_if.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if req_if.THE_HEADER != nil {
		stage.UnstageBranch(req_if.THE_HEADER)
	}
	if req_if.CORE_CONTENT != nil {
		stage.UnstageBranch(req_if.CORE_CONTENT)
	}
	if req_if.TOOL_EXTENSIONS != nil {
		stage.UnstageBranch(req_if.TOOL_EXTENSIONS)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (req_if_content *REQ_IF_CONTENT) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchREQ_IF_CONTENT(req_if_content)
}

func (stage *Stage) UnstageBranchREQ_IF_CONTENT(req_if_content *REQ_IF_CONTENT) {

	// check if instance is already staged
	if !stage.IsStaged(req_if_content) {
		return
	}

	req_if_content.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if req_if_content.DATATYPES != nil {
		stage.UnstageBranch(req_if_content.DATATYPES)
	}
	if req_if_content.SPEC_TYPES != nil {
		stage.UnstageBranch(req_if_content.SPEC_TYPES)
	}
	if req_if_content.SPEC_OBJECTS != nil {
		stage.UnstageBranch(req_if_content.SPEC_OBJECTS)
	}
	if req_if_content.SPEC_RELATIONS != nil {
		stage.UnstageBranch(req_if_content.SPEC_RELATIONS)
	}
	if req_if_content.SPECIFICATIONS != nil {
		stage.UnstageBranch(req_if_content.SPECIFICATIONS)
	}
	if req_if_content.SPEC_RELATION_GROUPS != nil {
		stage.UnstageBranch(req_if_content.SPEC_RELATION_GROUPS)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (req_if_header *REQ_IF_HEADER) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchREQ_IF_HEADER(req_if_header)
}

func (stage *Stage) UnstageBranchREQ_IF_HEADER(req_if_header *REQ_IF_HEADER) {

	// check if instance is already staged
	if !stage.IsStaged(req_if_header) {
		return
	}

	req_if_header.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchREQ_IF_TOOL_EXTENSION(req_if_tool_extension)
}

func (stage *Stage) UnstageBranchREQ_IF_TOOL_EXTENSION(req_if_tool_extension *REQ_IF_TOOL_EXTENSION) {

	// check if instance is already staged
	if !stage.IsStaged(req_if_tool_extension) {
		return
	}

	req_if_tool_extension.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

func (specification *SPECIFICATION) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSPECIFICATION(specification)
}

func (stage *Stage) UnstageBranchSPECIFICATION(specification *SPECIFICATION) {

	// check if instance is already staged
	if !stage.IsStaged(specification) {
		return
	}

	specification.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specification.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(specification.ALTERNATIVE_ID)
	}
	if specification.CHILDREN != nil {
		stage.UnstageBranch(specification.CHILDREN)
	}
	if specification.VALUES != nil {
		stage.UnstageBranch(specification.VALUES)
	}
	if specification.TYPE != nil {
		stage.UnstageBranch(specification.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (specification_type *SPECIFICATION_TYPE) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSPECIFICATION_TYPE(specification_type)
}

func (stage *Stage) UnstageBranchSPECIFICATION_TYPE(specification_type *SPECIFICATION_TYPE) {

	// check if instance is already staged
	if !stage.IsStaged(specification_type) {
		return
	}

	specification_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if specification_type.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(specification_type.ALTERNATIVE_ID)
	}
	if specification_type.SPEC_ATTRIBUTES != nil {
		stage.UnstageBranch(specification_type.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (spec_hierarchy *SPEC_HIERARCHY) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSPEC_HIERARCHY(spec_hierarchy)
}

func (stage *Stage) UnstageBranchSPEC_HIERARCHY(spec_hierarchy *SPEC_HIERARCHY) {

	// check if instance is already staged
	if !stage.IsStaged(spec_hierarchy) {
		return
	}

	spec_hierarchy.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_hierarchy.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(spec_hierarchy.ALTERNATIVE_ID)
	}
	if spec_hierarchy.CHILDREN != nil {
		stage.UnstageBranch(spec_hierarchy.CHILDREN)
	}
	if spec_hierarchy.EDITABLE_ATTS != nil {
		stage.UnstageBranch(spec_hierarchy.EDITABLE_ATTS)
	}
	if spec_hierarchy.OBJECT != nil {
		stage.UnstageBranch(spec_hierarchy.OBJECT)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (spec_object *SPEC_OBJECT) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSPEC_OBJECT(spec_object)
}

func (stage *Stage) UnstageBranchSPEC_OBJECT(spec_object *SPEC_OBJECT) {

	// check if instance is already staged
	if !stage.IsStaged(spec_object) {
		return
	}

	spec_object.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_object.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(spec_object.ALTERNATIVE_ID)
	}
	if spec_object.VALUES != nil {
		stage.UnstageBranch(spec_object.VALUES)
	}
	if spec_object.TYPE != nil {
		stage.UnstageBranch(spec_object.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (spec_object_type *SPEC_OBJECT_TYPE) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSPEC_OBJECT_TYPE(spec_object_type)
}

func (stage *Stage) UnstageBranchSPEC_OBJECT_TYPE(spec_object_type *SPEC_OBJECT_TYPE) {

	// check if instance is already staged
	if !stage.IsStaged(spec_object_type) {
		return
	}

	spec_object_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_object_type.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(spec_object_type.ALTERNATIVE_ID)
	}
	if spec_object_type.SPEC_ATTRIBUTES != nil {
		stage.UnstageBranch(spec_object_type.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (spec_relation *SPEC_RELATION) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSPEC_RELATION(spec_relation)
}

func (stage *Stage) UnstageBranchSPEC_RELATION(spec_relation *SPEC_RELATION) {

	// check if instance is already staged
	if !stage.IsStaged(spec_relation) {
		return
	}

	spec_relation.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_relation.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(spec_relation.ALTERNATIVE_ID)
	}
	if spec_relation.VALUES != nil {
		stage.UnstageBranch(spec_relation.VALUES)
	}
	if spec_relation.SOURCE != nil {
		stage.UnstageBranch(spec_relation.SOURCE)
	}
	if spec_relation.TARGET != nil {
		stage.UnstageBranch(spec_relation.TARGET)
	}
	if spec_relation.TYPE != nil {
		stage.UnstageBranch(spec_relation.TYPE)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (spec_relation_type *SPEC_RELATION_TYPE) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchSPEC_RELATION_TYPE(spec_relation_type)
}

func (stage *Stage) UnstageBranchSPEC_RELATION_TYPE(spec_relation_type *SPEC_RELATION_TYPE) {

	// check if instance is already staged
	if !stage.IsStaged(spec_relation_type) {
		return
	}

	spec_relation_type.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers
	if spec_relation_type.ALTERNATIVE_ID != nil {
		stage.UnstageBranch(spec_relation_type.ALTERNATIVE_ID)
	}
	if spec_relation_type.SPEC_ATTRIBUTES != nil {
		stage.UnstageBranch(spec_relation_type.SPEC_ATTRIBUTES)
	}

	//insertion point for the staging of instances referenced by slice of pointers

}

func (xhtml_content *XHTML_CONTENT) GongUnstageBranch(stage *Stage) {
	stage.UnstageBranchXHTML_CONTENT(xhtml_content)
}

func (stage *Stage) UnstageBranchXHTML_CONTENT(xhtml_content *XHTML_CONTENT) {

	// check if instance is already staged
	if !stage.IsStaged(xhtml_content) {
		return
	}

	xhtml_content.Unstage(stage)

	//insertion point for the staging of instances referenced by pointers

	//insertion point for the staging of instances referenced by slice of pointers

}

// insertion point for pointer reconstruction from references
func (reference *ALTERNATIVE_ID) GongReconstructPointersFromReferences(stage *Stage, instance *ALTERNATIVE_ID) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ATTRIBUTE_DEFINITION_BOOLEAN) GongReconstructPointersFromReferences(stage *Stage, instance *ATTRIBUTE_DEFINITION_BOOLEAN) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.DEFAULT_VALUE != nil {
		reference.DEFAULT_VALUE = stage.A_ATTRIBUTE_VALUE_BOOLEANs_reference[instance.DEFAULT_VALUE]
	}
	if instance.TYPE != nil {
		reference.TYPE = stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_reference[instance.TYPE]
	}
	// insertion point for slice of pointers field
}

func (reference *ATTRIBUTE_DEFINITION_DATE) GongReconstructPointersFromReferences(stage *Stage, instance *ATTRIBUTE_DEFINITION_DATE) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.DEFAULT_VALUE != nil {
		reference.DEFAULT_VALUE = stage.A_ATTRIBUTE_VALUE_DATEs_reference[instance.DEFAULT_VALUE]
	}
	if instance.TYPE != nil {
		reference.TYPE = stage.A_DATATYPE_DEFINITION_DATE_REFs_reference[instance.TYPE]
	}
	// insertion point for slice of pointers field
}

func (reference *ATTRIBUTE_DEFINITION_ENUMERATION) GongReconstructPointersFromReferences(stage *Stage, instance *ATTRIBUTE_DEFINITION_ENUMERATION) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.DEFAULT_VALUE != nil {
		reference.DEFAULT_VALUE = stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_reference[instance.DEFAULT_VALUE]
	}
	if instance.TYPE != nil {
		reference.TYPE = stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_reference[instance.TYPE]
	}
	// insertion point for slice of pointers field
}

func (reference *ATTRIBUTE_DEFINITION_INTEGER) GongReconstructPointersFromReferences(stage *Stage, instance *ATTRIBUTE_DEFINITION_INTEGER) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.DEFAULT_VALUE != nil {
		reference.DEFAULT_VALUE = stage.A_ATTRIBUTE_VALUE_INTEGERs_reference[instance.DEFAULT_VALUE]
	}
	if instance.TYPE != nil {
		reference.TYPE = stage.A_DATATYPE_DEFINITION_INTEGER_REFs_reference[instance.TYPE]
	}
	// insertion point for slice of pointers field
}

func (reference *ATTRIBUTE_DEFINITION_REAL) GongReconstructPointersFromReferences(stage *Stage, instance *ATTRIBUTE_DEFINITION_REAL) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.DEFAULT_VALUE != nil {
		reference.DEFAULT_VALUE = stage.A_ATTRIBUTE_VALUE_REALs_reference[instance.DEFAULT_VALUE]
	}
	if instance.TYPE != nil {
		reference.TYPE = stage.A_DATATYPE_DEFINITION_REAL_REFs_reference[instance.TYPE]
	}
	// insertion point for slice of pointers field
}

func (reference *ATTRIBUTE_DEFINITION_STRING) GongReconstructPointersFromReferences(stage *Stage, instance *ATTRIBUTE_DEFINITION_STRING) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.DEFAULT_VALUE != nil {
		reference.DEFAULT_VALUE = stage.A_ATTRIBUTE_VALUE_STRINGs_reference[instance.DEFAULT_VALUE]
	}
	if instance.TYPE != nil {
		reference.TYPE = stage.A_DATATYPE_DEFINITION_STRING_REFs_reference[instance.TYPE]
	}
	// insertion point for slice of pointers field
}

func (reference *ATTRIBUTE_DEFINITION_XHTML) GongReconstructPointersFromReferences(stage *Stage, instance *ATTRIBUTE_DEFINITION_XHTML) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.DEFAULT_VALUE != nil {
		reference.DEFAULT_VALUE = stage.A_ATTRIBUTE_VALUE_XHTMLs_reference[instance.DEFAULT_VALUE]
	}
	if instance.TYPE != nil {
		reference.TYPE = stage.A_DATATYPE_DEFINITION_XHTML_REFs_reference[instance.TYPE]
	}
	// insertion point for slice of pointers field
}

func (reference *ATTRIBUTE_VALUE_BOOLEAN) GongReconstructPointersFromReferences(stage *Stage, instance *ATTRIBUTE_VALUE_BOOLEAN) {
	// insertion point for pointers field
	if instance.DEFINITION != nil {
		reference.DEFINITION = stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_reference[instance.DEFINITION]
	}
	// insertion point for slice of pointers field
}

func (reference *ATTRIBUTE_VALUE_DATE) GongReconstructPointersFromReferences(stage *Stage, instance *ATTRIBUTE_VALUE_DATE) {
	// insertion point for pointers field
	if instance.DEFINITION != nil {
		reference.DEFINITION = stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_reference[instance.DEFINITION]
	}
	// insertion point for slice of pointers field
}

func (reference *ATTRIBUTE_VALUE_ENUMERATION) GongReconstructPointersFromReferences(stage *Stage, instance *ATTRIBUTE_VALUE_ENUMERATION) {
	// insertion point for pointers field
	if instance.DEFINITION != nil {
		reference.DEFINITION = stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_reference[instance.DEFINITION]
	}
	if instance.VALUES != nil {
		reference.VALUES = stage.A_ENUM_VALUE_REFs_reference[instance.VALUES]
	}
	// insertion point for slice of pointers field
}

func (reference *ATTRIBUTE_VALUE_INTEGER) GongReconstructPointersFromReferences(stage *Stage, instance *ATTRIBUTE_VALUE_INTEGER) {
	// insertion point for pointers field
	if instance.DEFINITION != nil {
		reference.DEFINITION = stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_reference[instance.DEFINITION]
	}
	// insertion point for slice of pointers field
}

func (reference *ATTRIBUTE_VALUE_REAL) GongReconstructPointersFromReferences(stage *Stage, instance *ATTRIBUTE_VALUE_REAL) {
	// insertion point for pointers field
	if instance.DEFINITION != nil {
		reference.DEFINITION = stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_reference[instance.DEFINITION]
	}
	// insertion point for slice of pointers field
}

func (reference *ATTRIBUTE_VALUE_STRING) GongReconstructPointersFromReferences(stage *Stage, instance *ATTRIBUTE_VALUE_STRING) {
	// insertion point for pointers field
	if instance.DEFINITION != nil {
		reference.DEFINITION = stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_reference[instance.DEFINITION]
	}
	// insertion point for slice of pointers field
}

func (reference *ATTRIBUTE_VALUE_XHTML) GongReconstructPointersFromReferences(stage *Stage, instance *ATTRIBUTE_VALUE_XHTML) {
	// insertion point for pointers field
	if instance.THE_VALUE != nil {
		reference.THE_VALUE = stage.XHTML_CONTENTs_reference[instance.THE_VALUE]
	}
	if instance.THE_ORIGINAL_VALUE != nil {
		reference.THE_ORIGINAL_VALUE = stage.XHTML_CONTENTs_reference[instance.THE_ORIGINAL_VALUE]
	}
	if instance.DEFINITION != nil {
		reference.DEFINITION = stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_reference[instance.DEFINITION]
	}
	// insertion point for slice of pointers field
}

func (reference *A_ALTERNATIVE_ID) GongReconstructPointersFromReferences(stage *Stage, instance *A_ALTERNATIVE_ID) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	// insertion point for slice of pointers field
}

func (reference *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_ATTRIBUTE_DEFINITION_DATE_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_ATTRIBUTE_DEFINITION_DATE_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_ATTRIBUTE_DEFINITION_INTEGER_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_ATTRIBUTE_DEFINITION_REAL_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_ATTRIBUTE_DEFINITION_REAL_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_ATTRIBUTE_DEFINITION_STRING_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_ATTRIBUTE_DEFINITION_STRING_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_ATTRIBUTE_DEFINITION_XHTML_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_ATTRIBUTE_VALUE_BOOLEAN) GongReconstructPointersFromReferences(stage *Stage, instance *A_ATTRIBUTE_VALUE_BOOLEAN) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.ATTRIBUTE_VALUE_BOOLEAN = reference.ATTRIBUTE_VALUE_BOOLEAN[:0]
	for _, _b := range instance.ATTRIBUTE_VALUE_BOOLEAN {
		reference.ATTRIBUTE_VALUE_BOOLEAN = append(reference.ATTRIBUTE_VALUE_BOOLEAN, stage.ATTRIBUTE_VALUE_BOOLEANs_reference[_b])
	}
}

func (reference *A_ATTRIBUTE_VALUE_DATE) GongReconstructPointersFromReferences(stage *Stage, instance *A_ATTRIBUTE_VALUE_DATE) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.ATTRIBUTE_VALUE_DATE = reference.ATTRIBUTE_VALUE_DATE[:0]
	for _, _b := range instance.ATTRIBUTE_VALUE_DATE {
		reference.ATTRIBUTE_VALUE_DATE = append(reference.ATTRIBUTE_VALUE_DATE, stage.ATTRIBUTE_VALUE_DATEs_reference[_b])
	}
}

func (reference *A_ATTRIBUTE_VALUE_ENUMERATION) GongReconstructPointersFromReferences(stage *Stage, instance *A_ATTRIBUTE_VALUE_ENUMERATION) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.ATTRIBUTE_VALUE_ENUMERATION = reference.ATTRIBUTE_VALUE_ENUMERATION[:0]
	for _, _b := range instance.ATTRIBUTE_VALUE_ENUMERATION {
		reference.ATTRIBUTE_VALUE_ENUMERATION = append(reference.ATTRIBUTE_VALUE_ENUMERATION, stage.ATTRIBUTE_VALUE_ENUMERATIONs_reference[_b])
	}
}

func (reference *A_ATTRIBUTE_VALUE_INTEGER) GongReconstructPointersFromReferences(stage *Stage, instance *A_ATTRIBUTE_VALUE_INTEGER) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.ATTRIBUTE_VALUE_INTEGER = reference.ATTRIBUTE_VALUE_INTEGER[:0]
	for _, _b := range instance.ATTRIBUTE_VALUE_INTEGER {
		reference.ATTRIBUTE_VALUE_INTEGER = append(reference.ATTRIBUTE_VALUE_INTEGER, stage.ATTRIBUTE_VALUE_INTEGERs_reference[_b])
	}
}

func (reference *A_ATTRIBUTE_VALUE_REAL) GongReconstructPointersFromReferences(stage *Stage, instance *A_ATTRIBUTE_VALUE_REAL) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.ATTRIBUTE_VALUE_REAL = reference.ATTRIBUTE_VALUE_REAL[:0]
	for _, _b := range instance.ATTRIBUTE_VALUE_REAL {
		reference.ATTRIBUTE_VALUE_REAL = append(reference.ATTRIBUTE_VALUE_REAL, stage.ATTRIBUTE_VALUE_REALs_reference[_b])
	}
}

func (reference *A_ATTRIBUTE_VALUE_STRING) GongReconstructPointersFromReferences(stage *Stage, instance *A_ATTRIBUTE_VALUE_STRING) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.ATTRIBUTE_VALUE_STRING = reference.ATTRIBUTE_VALUE_STRING[:0]
	for _, _b := range instance.ATTRIBUTE_VALUE_STRING {
		reference.ATTRIBUTE_VALUE_STRING = append(reference.ATTRIBUTE_VALUE_STRING, stage.ATTRIBUTE_VALUE_STRINGs_reference[_b])
	}
}

func (reference *A_ATTRIBUTE_VALUE_XHTML) GongReconstructPointersFromReferences(stage *Stage, instance *A_ATTRIBUTE_VALUE_XHTML) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.ATTRIBUTE_VALUE_XHTML = reference.ATTRIBUTE_VALUE_XHTML[:0]
	for _, _b := range instance.ATTRIBUTE_VALUE_XHTML {
		reference.ATTRIBUTE_VALUE_XHTML = append(reference.ATTRIBUTE_VALUE_XHTML, stage.ATTRIBUTE_VALUE_XHTMLs_reference[_b])
	}
}

func (reference *A_ATTRIBUTE_VALUE_XHTML_1) GongReconstructPointersFromReferences(stage *Stage, instance *A_ATTRIBUTE_VALUE_XHTML_1) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.ATTRIBUTE_VALUE_BOOLEAN = reference.ATTRIBUTE_VALUE_BOOLEAN[:0]
	for _, _b := range instance.ATTRIBUTE_VALUE_BOOLEAN {
		reference.ATTRIBUTE_VALUE_BOOLEAN = append(reference.ATTRIBUTE_VALUE_BOOLEAN, stage.ATTRIBUTE_VALUE_BOOLEANs_reference[_b])
	}
	reference.ATTRIBUTE_VALUE_DATE = reference.ATTRIBUTE_VALUE_DATE[:0]
	for _, _b := range instance.ATTRIBUTE_VALUE_DATE {
		reference.ATTRIBUTE_VALUE_DATE = append(reference.ATTRIBUTE_VALUE_DATE, stage.ATTRIBUTE_VALUE_DATEs_reference[_b])
	}
	reference.ATTRIBUTE_VALUE_ENUMERATION = reference.ATTRIBUTE_VALUE_ENUMERATION[:0]
	for _, _b := range instance.ATTRIBUTE_VALUE_ENUMERATION {
		reference.ATTRIBUTE_VALUE_ENUMERATION = append(reference.ATTRIBUTE_VALUE_ENUMERATION, stage.ATTRIBUTE_VALUE_ENUMERATIONs_reference[_b])
	}
	reference.ATTRIBUTE_VALUE_INTEGER = reference.ATTRIBUTE_VALUE_INTEGER[:0]
	for _, _b := range instance.ATTRIBUTE_VALUE_INTEGER {
		reference.ATTRIBUTE_VALUE_INTEGER = append(reference.ATTRIBUTE_VALUE_INTEGER, stage.ATTRIBUTE_VALUE_INTEGERs_reference[_b])
	}
	reference.ATTRIBUTE_VALUE_REAL = reference.ATTRIBUTE_VALUE_REAL[:0]
	for _, _b := range instance.ATTRIBUTE_VALUE_REAL {
		reference.ATTRIBUTE_VALUE_REAL = append(reference.ATTRIBUTE_VALUE_REAL, stage.ATTRIBUTE_VALUE_REALs_reference[_b])
	}
	reference.ATTRIBUTE_VALUE_STRING = reference.ATTRIBUTE_VALUE_STRING[:0]
	for _, _b := range instance.ATTRIBUTE_VALUE_STRING {
		reference.ATTRIBUTE_VALUE_STRING = append(reference.ATTRIBUTE_VALUE_STRING, stage.ATTRIBUTE_VALUE_STRINGs_reference[_b])
	}
	reference.ATTRIBUTE_VALUE_XHTML = reference.ATTRIBUTE_VALUE_XHTML[:0]
	for _, _b := range instance.ATTRIBUTE_VALUE_XHTML {
		reference.ATTRIBUTE_VALUE_XHTML = append(reference.ATTRIBUTE_VALUE_XHTML, stage.ATTRIBUTE_VALUE_XHTMLs_reference[_b])
	}
}

func (reference *A_CHILDREN) GongReconstructPointersFromReferences(stage *Stage, instance *A_CHILDREN) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.SPEC_HIERARCHY = reference.SPEC_HIERARCHY[:0]
	for _, _b := range instance.SPEC_HIERARCHY {
		reference.SPEC_HIERARCHY = append(reference.SPEC_HIERARCHY, stage.SPEC_HIERARCHYs_reference[_b])
	}
}

func (reference *A_CORE_CONTENT) GongReconstructPointersFromReferences(stage *Stage, instance *A_CORE_CONTENT) {
	// insertion point for pointers field
	if instance.REQ_IF_CONTENT != nil {
		reference.REQ_IF_CONTENT = stage.REQ_IF_CONTENTs_reference[instance.REQ_IF_CONTENT]
	}
	// insertion point for slice of pointers field
}

func (reference *A_DATATYPES) GongReconstructPointersFromReferences(stage *Stage, instance *A_DATATYPES) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.DATATYPE_DEFINITION_BOOLEAN = reference.DATATYPE_DEFINITION_BOOLEAN[:0]
	for _, _b := range instance.DATATYPE_DEFINITION_BOOLEAN {
		reference.DATATYPE_DEFINITION_BOOLEAN = append(reference.DATATYPE_DEFINITION_BOOLEAN, stage.DATATYPE_DEFINITION_BOOLEANs_reference[_b])
	}
	reference.DATATYPE_DEFINITION_DATE = reference.DATATYPE_DEFINITION_DATE[:0]
	for _, _b := range instance.DATATYPE_DEFINITION_DATE {
		reference.DATATYPE_DEFINITION_DATE = append(reference.DATATYPE_DEFINITION_DATE, stage.DATATYPE_DEFINITION_DATEs_reference[_b])
	}
	reference.DATATYPE_DEFINITION_ENUMERATION = reference.DATATYPE_DEFINITION_ENUMERATION[:0]
	for _, _b := range instance.DATATYPE_DEFINITION_ENUMERATION {
		reference.DATATYPE_DEFINITION_ENUMERATION = append(reference.DATATYPE_DEFINITION_ENUMERATION, stage.DATATYPE_DEFINITION_ENUMERATIONs_reference[_b])
	}
	reference.DATATYPE_DEFINITION_INTEGER = reference.DATATYPE_DEFINITION_INTEGER[:0]
	for _, _b := range instance.DATATYPE_DEFINITION_INTEGER {
		reference.DATATYPE_DEFINITION_INTEGER = append(reference.DATATYPE_DEFINITION_INTEGER, stage.DATATYPE_DEFINITION_INTEGERs_reference[_b])
	}
	reference.DATATYPE_DEFINITION_REAL = reference.DATATYPE_DEFINITION_REAL[:0]
	for _, _b := range instance.DATATYPE_DEFINITION_REAL {
		reference.DATATYPE_DEFINITION_REAL = append(reference.DATATYPE_DEFINITION_REAL, stage.DATATYPE_DEFINITION_REALs_reference[_b])
	}
	reference.DATATYPE_DEFINITION_STRING = reference.DATATYPE_DEFINITION_STRING[:0]
	for _, _b := range instance.DATATYPE_DEFINITION_STRING {
		reference.DATATYPE_DEFINITION_STRING = append(reference.DATATYPE_DEFINITION_STRING, stage.DATATYPE_DEFINITION_STRINGs_reference[_b])
	}
	reference.DATATYPE_DEFINITION_XHTML = reference.DATATYPE_DEFINITION_XHTML[:0]
	for _, _b := range instance.DATATYPE_DEFINITION_XHTML {
		reference.DATATYPE_DEFINITION_XHTML = append(reference.DATATYPE_DEFINITION_XHTML, stage.DATATYPE_DEFINITION_XHTMLs_reference[_b])
	}
}

func (reference *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_DATATYPE_DEFINITION_BOOLEAN_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_DATATYPE_DEFINITION_DATE_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_DATATYPE_DEFINITION_DATE_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_DATATYPE_DEFINITION_ENUMERATION_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_DATATYPE_DEFINITION_INTEGER_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_DATATYPE_DEFINITION_INTEGER_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_DATATYPE_DEFINITION_REAL_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_DATATYPE_DEFINITION_REAL_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_DATATYPE_DEFINITION_STRING_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_DATATYPE_DEFINITION_STRING_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_DATATYPE_DEFINITION_XHTML_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_DATATYPE_DEFINITION_XHTML_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_EDITABLE_ATTS) GongReconstructPointersFromReferences(stage *Stage, instance *A_EDITABLE_ATTS) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_ENUM_VALUE_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_ENUM_VALUE_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_OBJECT) GongReconstructPointersFromReferences(stage *Stage, instance *A_OBJECT) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_PROPERTIES) GongReconstructPointersFromReferences(stage *Stage, instance *A_PROPERTIES) {
	// insertion point for pointers field
	if instance.EMBEDDED_VALUE != nil {
		reference.EMBEDDED_VALUE = stage.EMBEDDED_VALUEs_reference[instance.EMBEDDED_VALUE]
	}
	// insertion point for slice of pointers field
}

func (reference *A_RELATION_GROUP_TYPE_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_RELATION_GROUP_TYPE_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_SOURCE_1) GongReconstructPointersFromReferences(stage *Stage, instance *A_SOURCE_1) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_SOURCE_SPECIFICATION_1) GongReconstructPointersFromReferences(stage *Stage, instance *A_SOURCE_SPECIFICATION_1) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_SPECIFICATIONS) GongReconstructPointersFromReferences(stage *Stage, instance *A_SPECIFICATIONS) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.SPECIFICATION = reference.SPECIFICATION[:0]
	for _, _b := range instance.SPECIFICATION {
		reference.SPECIFICATION = append(reference.SPECIFICATION, stage.SPECIFICATIONs_reference[_b])
	}
}

func (reference *A_SPECIFICATION_TYPE_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_SPECIFICATION_TYPE_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_SPECIFIED_VALUES) GongReconstructPointersFromReferences(stage *Stage, instance *A_SPECIFIED_VALUES) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.ENUM_VALUE = reference.ENUM_VALUE[:0]
	for _, _b := range instance.ENUM_VALUE {
		reference.ENUM_VALUE = append(reference.ENUM_VALUE, stage.ENUM_VALUEs_reference[_b])
	}
}

func (reference *A_SPEC_ATTRIBUTES) GongReconstructPointersFromReferences(stage *Stage, instance *A_SPEC_ATTRIBUTES) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.ATTRIBUTE_DEFINITION_BOOLEAN = reference.ATTRIBUTE_DEFINITION_BOOLEAN[:0]
	for _, _b := range instance.ATTRIBUTE_DEFINITION_BOOLEAN {
		reference.ATTRIBUTE_DEFINITION_BOOLEAN = append(reference.ATTRIBUTE_DEFINITION_BOOLEAN, stage.ATTRIBUTE_DEFINITION_BOOLEANs_reference[_b])
	}
	reference.ATTRIBUTE_DEFINITION_DATE = reference.ATTRIBUTE_DEFINITION_DATE[:0]
	for _, _b := range instance.ATTRIBUTE_DEFINITION_DATE {
		reference.ATTRIBUTE_DEFINITION_DATE = append(reference.ATTRIBUTE_DEFINITION_DATE, stage.ATTRIBUTE_DEFINITION_DATEs_reference[_b])
	}
	reference.ATTRIBUTE_DEFINITION_ENUMERATION = reference.ATTRIBUTE_DEFINITION_ENUMERATION[:0]
	for _, _b := range instance.ATTRIBUTE_DEFINITION_ENUMERATION {
		reference.ATTRIBUTE_DEFINITION_ENUMERATION = append(reference.ATTRIBUTE_DEFINITION_ENUMERATION, stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_reference[_b])
	}
	reference.ATTRIBUTE_DEFINITION_INTEGER = reference.ATTRIBUTE_DEFINITION_INTEGER[:0]
	for _, _b := range instance.ATTRIBUTE_DEFINITION_INTEGER {
		reference.ATTRIBUTE_DEFINITION_INTEGER = append(reference.ATTRIBUTE_DEFINITION_INTEGER, stage.ATTRIBUTE_DEFINITION_INTEGERs_reference[_b])
	}
	reference.ATTRIBUTE_DEFINITION_REAL = reference.ATTRIBUTE_DEFINITION_REAL[:0]
	for _, _b := range instance.ATTRIBUTE_DEFINITION_REAL {
		reference.ATTRIBUTE_DEFINITION_REAL = append(reference.ATTRIBUTE_DEFINITION_REAL, stage.ATTRIBUTE_DEFINITION_REALs_reference[_b])
	}
	reference.ATTRIBUTE_DEFINITION_STRING = reference.ATTRIBUTE_DEFINITION_STRING[:0]
	for _, _b := range instance.ATTRIBUTE_DEFINITION_STRING {
		reference.ATTRIBUTE_DEFINITION_STRING = append(reference.ATTRIBUTE_DEFINITION_STRING, stage.ATTRIBUTE_DEFINITION_STRINGs_reference[_b])
	}
	reference.ATTRIBUTE_DEFINITION_XHTML = reference.ATTRIBUTE_DEFINITION_XHTML[:0]
	for _, _b := range instance.ATTRIBUTE_DEFINITION_XHTML {
		reference.ATTRIBUTE_DEFINITION_XHTML = append(reference.ATTRIBUTE_DEFINITION_XHTML, stage.ATTRIBUTE_DEFINITION_XHTMLs_reference[_b])
	}
}

func (reference *A_SPEC_OBJECTS) GongReconstructPointersFromReferences(stage *Stage, instance *A_SPEC_OBJECTS) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.SPEC_OBJECT = reference.SPEC_OBJECT[:0]
	for _, _b := range instance.SPEC_OBJECT {
		reference.SPEC_OBJECT = append(reference.SPEC_OBJECT, stage.SPEC_OBJECTs_reference[_b])
	}
}

func (reference *A_SPEC_OBJECT_TYPE_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_SPEC_OBJECT_TYPE_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_SPEC_RELATIONS) GongReconstructPointersFromReferences(stage *Stage, instance *A_SPEC_RELATIONS) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.SPEC_RELATION = reference.SPEC_RELATION[:0]
	for _, _b := range instance.SPEC_RELATION {
		reference.SPEC_RELATION = append(reference.SPEC_RELATION, stage.SPEC_RELATIONs_reference[_b])
	}
}

func (reference *A_SPEC_RELATION_GROUPS) GongReconstructPointersFromReferences(stage *Stage, instance *A_SPEC_RELATION_GROUPS) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.RELATION_GROUP = reference.RELATION_GROUP[:0]
	for _, _b := range instance.RELATION_GROUP {
		reference.RELATION_GROUP = append(reference.RELATION_GROUP, stage.RELATION_GROUPs_reference[_b])
	}
}

func (reference *A_SPEC_RELATION_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_SPEC_RELATION_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_SPEC_RELATION_TYPE_REF) GongReconstructPointersFromReferences(stage *Stage, instance *A_SPEC_RELATION_TYPE_REF) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *A_SPEC_TYPES) GongReconstructPointersFromReferences(stage *Stage, instance *A_SPEC_TYPES) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.RELATION_GROUP_TYPE = reference.RELATION_GROUP_TYPE[:0]
	for _, _b := range instance.RELATION_GROUP_TYPE {
		reference.RELATION_GROUP_TYPE = append(reference.RELATION_GROUP_TYPE, stage.RELATION_GROUP_TYPEs_reference[_b])
	}
	reference.SPEC_OBJECT_TYPE = reference.SPEC_OBJECT_TYPE[:0]
	for _, _b := range instance.SPEC_OBJECT_TYPE {
		reference.SPEC_OBJECT_TYPE = append(reference.SPEC_OBJECT_TYPE, stage.SPEC_OBJECT_TYPEs_reference[_b])
	}
	reference.SPEC_RELATION_TYPE = reference.SPEC_RELATION_TYPE[:0]
	for _, _b := range instance.SPEC_RELATION_TYPE {
		reference.SPEC_RELATION_TYPE = append(reference.SPEC_RELATION_TYPE, stage.SPEC_RELATION_TYPEs_reference[_b])
	}
	reference.SPECIFICATION_TYPE = reference.SPECIFICATION_TYPE[:0]
	for _, _b := range instance.SPECIFICATION_TYPE {
		reference.SPECIFICATION_TYPE = append(reference.SPECIFICATION_TYPE, stage.SPECIFICATION_TYPEs_reference[_b])
	}
}

func (reference *A_THE_HEADER) GongReconstructPointersFromReferences(stage *Stage, instance *A_THE_HEADER) {
	// insertion point for pointers field
	if instance.REQ_IF_HEADER != nil {
		reference.REQ_IF_HEADER = stage.REQ_IF_HEADERs_reference[instance.REQ_IF_HEADER]
	}
	// insertion point for slice of pointers field
}

func (reference *A_TOOL_EXTENSIONS) GongReconstructPointersFromReferences(stage *Stage, instance *A_TOOL_EXTENSIONS) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
	reference.REQ_IF_TOOL_EXTENSION = reference.REQ_IF_TOOL_EXTENSION[:0]
	for _, _b := range instance.REQ_IF_TOOL_EXTENSION {
		reference.REQ_IF_TOOL_EXTENSION = append(reference.REQ_IF_TOOL_EXTENSION, stage.REQ_IF_TOOL_EXTENSIONs_reference[_b])
	}
}

func (reference *DATATYPE_DEFINITION_BOOLEAN) GongReconstructPointersFromReferences(stage *Stage, instance *DATATYPE_DEFINITION_BOOLEAN) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	// insertion point for slice of pointers field
}

func (reference *DATATYPE_DEFINITION_DATE) GongReconstructPointersFromReferences(stage *Stage, instance *DATATYPE_DEFINITION_DATE) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	// insertion point for slice of pointers field
}

func (reference *DATATYPE_DEFINITION_ENUMERATION) GongReconstructPointersFromReferences(stage *Stage, instance *DATATYPE_DEFINITION_ENUMERATION) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.SPECIFIED_VALUES != nil {
		reference.SPECIFIED_VALUES = stage.A_SPECIFIED_VALUESs_reference[instance.SPECIFIED_VALUES]
	}
	// insertion point for slice of pointers field
}

func (reference *DATATYPE_DEFINITION_INTEGER) GongReconstructPointersFromReferences(stage *Stage, instance *DATATYPE_DEFINITION_INTEGER) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	// insertion point for slice of pointers field
}

func (reference *DATATYPE_DEFINITION_REAL) GongReconstructPointersFromReferences(stage *Stage, instance *DATATYPE_DEFINITION_REAL) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	// insertion point for slice of pointers field
}

func (reference *DATATYPE_DEFINITION_STRING) GongReconstructPointersFromReferences(stage *Stage, instance *DATATYPE_DEFINITION_STRING) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	// insertion point for slice of pointers field
}

func (reference *DATATYPE_DEFINITION_XHTML) GongReconstructPointersFromReferences(stage *Stage, instance *DATATYPE_DEFINITION_XHTML) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	// insertion point for slice of pointers field
}

func (reference *EMBEDDED_VALUE) GongReconstructPointersFromReferences(stage *Stage, instance *EMBEDDED_VALUE) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *ENUM_VALUE) GongReconstructPointersFromReferences(stage *Stage, instance *ENUM_VALUE) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.PROPERTIES != nil {
		reference.PROPERTIES = stage.A_PROPERTIESs_reference[instance.PROPERTIES]
	}
	// insertion point for slice of pointers field
}

func (reference *RELATION_GROUP) GongReconstructPointersFromReferences(stage *Stage, instance *RELATION_GROUP) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.SOURCE_SPECIFICATION != nil {
		reference.SOURCE_SPECIFICATION = stage.A_SOURCE_SPECIFICATION_1s_reference[instance.SOURCE_SPECIFICATION]
	}
	if instance.SPEC_RELATIONS != nil {
		reference.SPEC_RELATIONS = stage.A_SPEC_RELATION_REFs_reference[instance.SPEC_RELATIONS]
	}
	if instance.TARGET_SPECIFICATION != nil {
		reference.TARGET_SPECIFICATION = stage.A_SOURCE_SPECIFICATION_1s_reference[instance.TARGET_SPECIFICATION]
	}
	if instance.TYPE != nil {
		reference.TYPE = stage.A_RELATION_GROUP_TYPE_REFs_reference[instance.TYPE]
	}
	// insertion point for slice of pointers field
}

func (reference *RELATION_GROUP_TYPE) GongReconstructPointersFromReferences(stage *Stage, instance *RELATION_GROUP_TYPE) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.SPEC_ATTRIBUTES != nil {
		reference.SPEC_ATTRIBUTES = stage.A_SPEC_ATTRIBUTESs_reference[instance.SPEC_ATTRIBUTES]
	}
	// insertion point for slice of pointers field
}

func (reference *REQ_IF) GongReconstructPointersFromReferences(stage *Stage, instance *REQ_IF) {
	// insertion point for pointers field
	if instance.THE_HEADER != nil {
		reference.THE_HEADER = stage.A_THE_HEADERs_reference[instance.THE_HEADER]
	}
	if instance.CORE_CONTENT != nil {
		reference.CORE_CONTENT = stage.A_CORE_CONTENTs_reference[instance.CORE_CONTENT]
	}
	if instance.TOOL_EXTENSIONS != nil {
		reference.TOOL_EXTENSIONS = stage.A_TOOL_EXTENSIONSs_reference[instance.TOOL_EXTENSIONS]
	}
	// insertion point for slice of pointers field
}

func (reference *REQ_IF_CONTENT) GongReconstructPointersFromReferences(stage *Stage, instance *REQ_IF_CONTENT) {
	// insertion point for pointers field
	if instance.DATATYPES != nil {
		reference.DATATYPES = stage.A_DATATYPESs_reference[instance.DATATYPES]
	}
	if instance.SPEC_TYPES != nil {
		reference.SPEC_TYPES = stage.A_SPEC_TYPESs_reference[instance.SPEC_TYPES]
	}
	if instance.SPEC_OBJECTS != nil {
		reference.SPEC_OBJECTS = stage.A_SPEC_OBJECTSs_reference[instance.SPEC_OBJECTS]
	}
	if instance.SPEC_RELATIONS != nil {
		reference.SPEC_RELATIONS = stage.A_SPEC_RELATIONSs_reference[instance.SPEC_RELATIONS]
	}
	if instance.SPECIFICATIONS != nil {
		reference.SPECIFICATIONS = stage.A_SPECIFICATIONSs_reference[instance.SPECIFICATIONS]
	}
	if instance.SPEC_RELATION_GROUPS != nil {
		reference.SPEC_RELATION_GROUPS = stage.A_SPEC_RELATION_GROUPSs_reference[instance.SPEC_RELATION_GROUPS]
	}
	// insertion point for slice of pointers field
}

func (reference *REQ_IF_HEADER) GongReconstructPointersFromReferences(stage *Stage, instance *REQ_IF_HEADER) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *REQ_IF_TOOL_EXTENSION) GongReconstructPointersFromReferences(stage *Stage, instance *REQ_IF_TOOL_EXTENSION) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

func (reference *SPECIFICATION) GongReconstructPointersFromReferences(stage *Stage, instance *SPECIFICATION) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.CHILDREN != nil {
		reference.CHILDREN = stage.A_CHILDRENs_reference[instance.CHILDREN]
	}
	if instance.VALUES != nil {
		reference.VALUES = stage.A_ATTRIBUTE_VALUE_XHTML_1s_reference[instance.VALUES]
	}
	if instance.TYPE != nil {
		reference.TYPE = stage.A_SPECIFICATION_TYPE_REFs_reference[instance.TYPE]
	}
	// insertion point for slice of pointers field
}

func (reference *SPECIFICATION_TYPE) GongReconstructPointersFromReferences(stage *Stage, instance *SPECIFICATION_TYPE) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.SPEC_ATTRIBUTES != nil {
		reference.SPEC_ATTRIBUTES = stage.A_SPEC_ATTRIBUTESs_reference[instance.SPEC_ATTRIBUTES]
	}
	// insertion point for slice of pointers field
}

func (reference *SPEC_HIERARCHY) GongReconstructPointersFromReferences(stage *Stage, instance *SPEC_HIERARCHY) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.CHILDREN != nil {
		reference.CHILDREN = stage.A_CHILDRENs_reference[instance.CHILDREN]
	}
	if instance.EDITABLE_ATTS != nil {
		reference.EDITABLE_ATTS = stage.A_EDITABLE_ATTSs_reference[instance.EDITABLE_ATTS]
	}
	if instance.OBJECT != nil {
		reference.OBJECT = stage.A_OBJECTs_reference[instance.OBJECT]
	}
	// insertion point for slice of pointers field
}

func (reference *SPEC_OBJECT) GongReconstructPointersFromReferences(stage *Stage, instance *SPEC_OBJECT) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.VALUES != nil {
		reference.VALUES = stage.A_ATTRIBUTE_VALUE_XHTML_1s_reference[instance.VALUES]
	}
	if instance.TYPE != nil {
		reference.TYPE = stage.A_SPEC_OBJECT_TYPE_REFs_reference[instance.TYPE]
	}
	// insertion point for slice of pointers field
}

func (reference *SPEC_OBJECT_TYPE) GongReconstructPointersFromReferences(stage *Stage, instance *SPEC_OBJECT_TYPE) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.SPEC_ATTRIBUTES != nil {
		reference.SPEC_ATTRIBUTES = stage.A_SPEC_ATTRIBUTESs_reference[instance.SPEC_ATTRIBUTES]
	}
	// insertion point for slice of pointers field
}

func (reference *SPEC_RELATION) GongReconstructPointersFromReferences(stage *Stage, instance *SPEC_RELATION) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.VALUES != nil {
		reference.VALUES = stage.A_ATTRIBUTE_VALUE_XHTML_1s_reference[instance.VALUES]
	}
	if instance.SOURCE != nil {
		reference.SOURCE = stage.A_SOURCE_1s_reference[instance.SOURCE]
	}
	if instance.TARGET != nil {
		reference.TARGET = stage.A_SOURCE_1s_reference[instance.TARGET]
	}
	if instance.TYPE != nil {
		reference.TYPE = stage.A_SPEC_RELATION_TYPE_REFs_reference[instance.TYPE]
	}
	// insertion point for slice of pointers field
}

func (reference *SPEC_RELATION_TYPE) GongReconstructPointersFromReferences(stage *Stage, instance *SPEC_RELATION_TYPE) {
	// insertion point for pointers field
	if instance.ALTERNATIVE_ID != nil {
		reference.ALTERNATIVE_ID = stage.A_ALTERNATIVE_IDs_reference[instance.ALTERNATIVE_ID]
	}
	if instance.SPEC_ATTRIBUTES != nil {
		reference.SPEC_ATTRIBUTES = stage.A_SPEC_ATTRIBUTESs_reference[instance.SPEC_ATTRIBUTES]
	}
	// insertion point for slice of pointers field
}

func (reference *XHTML_CONTENT) GongReconstructPointersFromReferences(stage *Stage, instance *XHTML_CONTENT) {
	// insertion point for pointers field
	// insertion point for slice of pointers field
}

// insertion point for pointer reconstruction from instances
func (reference *ALTERNATIVE_ID) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ATTRIBUTE_DEFINITION_BOOLEAN) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.DEFAULT_VALUE; _reference != nil {
		reference.DEFAULT_VALUE = nil
		if _instance, ok := stage.A_ATTRIBUTE_VALUE_BOOLEANs_instance[_reference]; ok {
			reference.DEFAULT_VALUE = _instance
		}
	}
	if _reference := reference.TYPE; _reference != nil {
		reference.TYPE = nil
		if _instance, ok := stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs_instance[_reference]; ok {
			reference.TYPE = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ATTRIBUTE_DEFINITION_DATE) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.DEFAULT_VALUE; _reference != nil {
		reference.DEFAULT_VALUE = nil
		if _instance, ok := stage.A_ATTRIBUTE_VALUE_DATEs_instance[_reference]; ok {
			reference.DEFAULT_VALUE = _instance
		}
	}
	if _reference := reference.TYPE; _reference != nil {
		reference.TYPE = nil
		if _instance, ok := stage.A_DATATYPE_DEFINITION_DATE_REFs_instance[_reference]; ok {
			reference.TYPE = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ATTRIBUTE_DEFINITION_ENUMERATION) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.DEFAULT_VALUE; _reference != nil {
		reference.DEFAULT_VALUE = nil
		if _instance, ok := stage.A_ATTRIBUTE_VALUE_ENUMERATIONs_instance[_reference]; ok {
			reference.DEFAULT_VALUE = _instance
		}
	}
	if _reference := reference.TYPE; _reference != nil {
		reference.TYPE = nil
		if _instance, ok := stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs_instance[_reference]; ok {
			reference.TYPE = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ATTRIBUTE_DEFINITION_INTEGER) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.DEFAULT_VALUE; _reference != nil {
		reference.DEFAULT_VALUE = nil
		if _instance, ok := stage.A_ATTRIBUTE_VALUE_INTEGERs_instance[_reference]; ok {
			reference.DEFAULT_VALUE = _instance
		}
	}
	if _reference := reference.TYPE; _reference != nil {
		reference.TYPE = nil
		if _instance, ok := stage.A_DATATYPE_DEFINITION_INTEGER_REFs_instance[_reference]; ok {
			reference.TYPE = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ATTRIBUTE_DEFINITION_REAL) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.DEFAULT_VALUE; _reference != nil {
		reference.DEFAULT_VALUE = nil
		if _instance, ok := stage.A_ATTRIBUTE_VALUE_REALs_instance[_reference]; ok {
			reference.DEFAULT_VALUE = _instance
		}
	}
	if _reference := reference.TYPE; _reference != nil {
		reference.TYPE = nil
		if _instance, ok := stage.A_DATATYPE_DEFINITION_REAL_REFs_instance[_reference]; ok {
			reference.TYPE = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ATTRIBUTE_DEFINITION_STRING) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.DEFAULT_VALUE; _reference != nil {
		reference.DEFAULT_VALUE = nil
		if _instance, ok := stage.A_ATTRIBUTE_VALUE_STRINGs_instance[_reference]; ok {
			reference.DEFAULT_VALUE = _instance
		}
	}
	if _reference := reference.TYPE; _reference != nil {
		reference.TYPE = nil
		if _instance, ok := stage.A_DATATYPE_DEFINITION_STRING_REFs_instance[_reference]; ok {
			reference.TYPE = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ATTRIBUTE_DEFINITION_XHTML) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.DEFAULT_VALUE; _reference != nil {
		reference.DEFAULT_VALUE = nil
		if _instance, ok := stage.A_ATTRIBUTE_VALUE_XHTMLs_instance[_reference]; ok {
			reference.DEFAULT_VALUE = _instance
		}
	}
	if _reference := reference.TYPE; _reference != nil {
		reference.TYPE = nil
		if _instance, ok := stage.A_DATATYPE_DEFINITION_XHTML_REFs_instance[_reference]; ok {
			reference.TYPE = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ATTRIBUTE_VALUE_BOOLEAN) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.DEFINITION; _reference != nil {
		reference.DEFINITION = nil
		if _instance, ok := stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs_instance[_reference]; ok {
			reference.DEFINITION = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ATTRIBUTE_VALUE_DATE) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.DEFINITION; _reference != nil {
		reference.DEFINITION = nil
		if _instance, ok := stage.A_ATTRIBUTE_DEFINITION_DATE_REFs_instance[_reference]; ok {
			reference.DEFINITION = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ATTRIBUTE_VALUE_ENUMERATION) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.DEFINITION; _reference != nil {
		reference.DEFINITION = nil
		if _instance, ok := stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs_instance[_reference]; ok {
			reference.DEFINITION = _instance
		}
	}
	if _reference := reference.VALUES; _reference != nil {
		reference.VALUES = nil
		if _instance, ok := stage.A_ENUM_VALUE_REFs_instance[_reference]; ok {
			reference.VALUES = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ATTRIBUTE_VALUE_INTEGER) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.DEFINITION; _reference != nil {
		reference.DEFINITION = nil
		if _instance, ok := stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs_instance[_reference]; ok {
			reference.DEFINITION = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ATTRIBUTE_VALUE_REAL) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.DEFINITION; _reference != nil {
		reference.DEFINITION = nil
		if _instance, ok := stage.A_ATTRIBUTE_DEFINITION_REAL_REFs_instance[_reference]; ok {
			reference.DEFINITION = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ATTRIBUTE_VALUE_STRING) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.DEFINITION; _reference != nil {
		reference.DEFINITION = nil
		if _instance, ok := stage.A_ATTRIBUTE_DEFINITION_STRING_REFs_instance[_reference]; ok {
			reference.DEFINITION = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *ATTRIBUTE_VALUE_XHTML) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.THE_VALUE; _reference != nil {
		reference.THE_VALUE = nil
		if _instance, ok := stage.XHTML_CONTENTs_instance[_reference]; ok {
			reference.THE_VALUE = _instance
		}
	}
	if _reference := reference.THE_ORIGINAL_VALUE; _reference != nil {
		reference.THE_ORIGINAL_VALUE = nil
		if _instance, ok := stage.XHTML_CONTENTs_instance[_reference]; ok {
			reference.THE_ORIGINAL_VALUE = _instance
		}
	}
	if _reference := reference.DEFINITION; _reference != nil {
		reference.DEFINITION = nil
		if _instance, ok := stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs_instance[_reference]; ok {
			reference.DEFINITION = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *A_ALTERNATIVE_ID) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_ATTRIBUTE_DEFINITION_DATE_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_ATTRIBUTE_DEFINITION_REAL_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_ATTRIBUTE_DEFINITION_STRING_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_ATTRIBUTE_VALUE_BOOLEAN) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _ATTRIBUTE_VALUE_BOOLEAN []*ATTRIBUTE_VALUE_BOOLEAN
	for _, _reference := range reference.ATTRIBUTE_VALUE_BOOLEAN {
		if _instance, ok := stage.ATTRIBUTE_VALUE_BOOLEANs_instance[_reference]; ok {
			_ATTRIBUTE_VALUE_BOOLEAN = append(_ATTRIBUTE_VALUE_BOOLEAN, _instance)
		}
	}
	reference.ATTRIBUTE_VALUE_BOOLEAN = _ATTRIBUTE_VALUE_BOOLEAN
}

func (reference *A_ATTRIBUTE_VALUE_DATE) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _ATTRIBUTE_VALUE_DATE []*ATTRIBUTE_VALUE_DATE
	for _, _reference := range reference.ATTRIBUTE_VALUE_DATE {
		if _instance, ok := stage.ATTRIBUTE_VALUE_DATEs_instance[_reference]; ok {
			_ATTRIBUTE_VALUE_DATE = append(_ATTRIBUTE_VALUE_DATE, _instance)
		}
	}
	reference.ATTRIBUTE_VALUE_DATE = _ATTRIBUTE_VALUE_DATE
}

func (reference *A_ATTRIBUTE_VALUE_ENUMERATION) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _ATTRIBUTE_VALUE_ENUMERATION []*ATTRIBUTE_VALUE_ENUMERATION
	for _, _reference := range reference.ATTRIBUTE_VALUE_ENUMERATION {
		if _instance, ok := stage.ATTRIBUTE_VALUE_ENUMERATIONs_instance[_reference]; ok {
			_ATTRIBUTE_VALUE_ENUMERATION = append(_ATTRIBUTE_VALUE_ENUMERATION, _instance)
		}
	}
	reference.ATTRIBUTE_VALUE_ENUMERATION = _ATTRIBUTE_VALUE_ENUMERATION
}

func (reference *A_ATTRIBUTE_VALUE_INTEGER) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _ATTRIBUTE_VALUE_INTEGER []*ATTRIBUTE_VALUE_INTEGER
	for _, _reference := range reference.ATTRIBUTE_VALUE_INTEGER {
		if _instance, ok := stage.ATTRIBUTE_VALUE_INTEGERs_instance[_reference]; ok {
			_ATTRIBUTE_VALUE_INTEGER = append(_ATTRIBUTE_VALUE_INTEGER, _instance)
		}
	}
	reference.ATTRIBUTE_VALUE_INTEGER = _ATTRIBUTE_VALUE_INTEGER
}

func (reference *A_ATTRIBUTE_VALUE_REAL) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _ATTRIBUTE_VALUE_REAL []*ATTRIBUTE_VALUE_REAL
	for _, _reference := range reference.ATTRIBUTE_VALUE_REAL {
		if _instance, ok := stage.ATTRIBUTE_VALUE_REALs_instance[_reference]; ok {
			_ATTRIBUTE_VALUE_REAL = append(_ATTRIBUTE_VALUE_REAL, _instance)
		}
	}
	reference.ATTRIBUTE_VALUE_REAL = _ATTRIBUTE_VALUE_REAL
}

func (reference *A_ATTRIBUTE_VALUE_STRING) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _ATTRIBUTE_VALUE_STRING []*ATTRIBUTE_VALUE_STRING
	for _, _reference := range reference.ATTRIBUTE_VALUE_STRING {
		if _instance, ok := stage.ATTRIBUTE_VALUE_STRINGs_instance[_reference]; ok {
			_ATTRIBUTE_VALUE_STRING = append(_ATTRIBUTE_VALUE_STRING, _instance)
		}
	}
	reference.ATTRIBUTE_VALUE_STRING = _ATTRIBUTE_VALUE_STRING
}

func (reference *A_ATTRIBUTE_VALUE_XHTML) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _ATTRIBUTE_VALUE_XHTML []*ATTRIBUTE_VALUE_XHTML
	for _, _reference := range reference.ATTRIBUTE_VALUE_XHTML {
		if _instance, ok := stage.ATTRIBUTE_VALUE_XHTMLs_instance[_reference]; ok {
			_ATTRIBUTE_VALUE_XHTML = append(_ATTRIBUTE_VALUE_XHTML, _instance)
		}
	}
	reference.ATTRIBUTE_VALUE_XHTML = _ATTRIBUTE_VALUE_XHTML
}

func (reference *A_ATTRIBUTE_VALUE_XHTML_1) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _ATTRIBUTE_VALUE_BOOLEAN []*ATTRIBUTE_VALUE_BOOLEAN
	for _, _reference := range reference.ATTRIBUTE_VALUE_BOOLEAN {
		if _instance, ok := stage.ATTRIBUTE_VALUE_BOOLEANs_instance[_reference]; ok {
			_ATTRIBUTE_VALUE_BOOLEAN = append(_ATTRIBUTE_VALUE_BOOLEAN, _instance)
		}
	}
	reference.ATTRIBUTE_VALUE_BOOLEAN = _ATTRIBUTE_VALUE_BOOLEAN
	var _ATTRIBUTE_VALUE_DATE []*ATTRIBUTE_VALUE_DATE
	for _, _reference := range reference.ATTRIBUTE_VALUE_DATE {
		if _instance, ok := stage.ATTRIBUTE_VALUE_DATEs_instance[_reference]; ok {
			_ATTRIBUTE_VALUE_DATE = append(_ATTRIBUTE_VALUE_DATE, _instance)
		}
	}
	reference.ATTRIBUTE_VALUE_DATE = _ATTRIBUTE_VALUE_DATE
	var _ATTRIBUTE_VALUE_ENUMERATION []*ATTRIBUTE_VALUE_ENUMERATION
	for _, _reference := range reference.ATTRIBUTE_VALUE_ENUMERATION {
		if _instance, ok := stage.ATTRIBUTE_VALUE_ENUMERATIONs_instance[_reference]; ok {
			_ATTRIBUTE_VALUE_ENUMERATION = append(_ATTRIBUTE_VALUE_ENUMERATION, _instance)
		}
	}
	reference.ATTRIBUTE_VALUE_ENUMERATION = _ATTRIBUTE_VALUE_ENUMERATION
	var _ATTRIBUTE_VALUE_INTEGER []*ATTRIBUTE_VALUE_INTEGER
	for _, _reference := range reference.ATTRIBUTE_VALUE_INTEGER {
		if _instance, ok := stage.ATTRIBUTE_VALUE_INTEGERs_instance[_reference]; ok {
			_ATTRIBUTE_VALUE_INTEGER = append(_ATTRIBUTE_VALUE_INTEGER, _instance)
		}
	}
	reference.ATTRIBUTE_VALUE_INTEGER = _ATTRIBUTE_VALUE_INTEGER
	var _ATTRIBUTE_VALUE_REAL []*ATTRIBUTE_VALUE_REAL
	for _, _reference := range reference.ATTRIBUTE_VALUE_REAL {
		if _instance, ok := stage.ATTRIBUTE_VALUE_REALs_instance[_reference]; ok {
			_ATTRIBUTE_VALUE_REAL = append(_ATTRIBUTE_VALUE_REAL, _instance)
		}
	}
	reference.ATTRIBUTE_VALUE_REAL = _ATTRIBUTE_VALUE_REAL
	var _ATTRIBUTE_VALUE_STRING []*ATTRIBUTE_VALUE_STRING
	for _, _reference := range reference.ATTRIBUTE_VALUE_STRING {
		if _instance, ok := stage.ATTRIBUTE_VALUE_STRINGs_instance[_reference]; ok {
			_ATTRIBUTE_VALUE_STRING = append(_ATTRIBUTE_VALUE_STRING, _instance)
		}
	}
	reference.ATTRIBUTE_VALUE_STRING = _ATTRIBUTE_VALUE_STRING
	var _ATTRIBUTE_VALUE_XHTML []*ATTRIBUTE_VALUE_XHTML
	for _, _reference := range reference.ATTRIBUTE_VALUE_XHTML {
		if _instance, ok := stage.ATTRIBUTE_VALUE_XHTMLs_instance[_reference]; ok {
			_ATTRIBUTE_VALUE_XHTML = append(_ATTRIBUTE_VALUE_XHTML, _instance)
		}
	}
	reference.ATTRIBUTE_VALUE_XHTML = _ATTRIBUTE_VALUE_XHTML
}

func (reference *A_CHILDREN) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _SPEC_HIERARCHY []*SPEC_HIERARCHY
	for _, _reference := range reference.SPEC_HIERARCHY {
		if _instance, ok := stage.SPEC_HIERARCHYs_instance[_reference]; ok {
			_SPEC_HIERARCHY = append(_SPEC_HIERARCHY, _instance)
		}
	}
	reference.SPEC_HIERARCHY = _SPEC_HIERARCHY
}

func (reference *A_CORE_CONTENT) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.REQ_IF_CONTENT; _reference != nil {
		reference.REQ_IF_CONTENT = nil
		if _instance, ok := stage.REQ_IF_CONTENTs_instance[_reference]; ok {
			reference.REQ_IF_CONTENT = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *A_DATATYPES) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _DATATYPE_DEFINITION_BOOLEAN []*DATATYPE_DEFINITION_BOOLEAN
	for _, _reference := range reference.DATATYPE_DEFINITION_BOOLEAN {
		if _instance, ok := stage.DATATYPE_DEFINITION_BOOLEANs_instance[_reference]; ok {
			_DATATYPE_DEFINITION_BOOLEAN = append(_DATATYPE_DEFINITION_BOOLEAN, _instance)
		}
	}
	reference.DATATYPE_DEFINITION_BOOLEAN = _DATATYPE_DEFINITION_BOOLEAN
	var _DATATYPE_DEFINITION_DATE []*DATATYPE_DEFINITION_DATE
	for _, _reference := range reference.DATATYPE_DEFINITION_DATE {
		if _instance, ok := stage.DATATYPE_DEFINITION_DATEs_instance[_reference]; ok {
			_DATATYPE_DEFINITION_DATE = append(_DATATYPE_DEFINITION_DATE, _instance)
		}
	}
	reference.DATATYPE_DEFINITION_DATE = _DATATYPE_DEFINITION_DATE
	var _DATATYPE_DEFINITION_ENUMERATION []*DATATYPE_DEFINITION_ENUMERATION
	for _, _reference := range reference.DATATYPE_DEFINITION_ENUMERATION {
		if _instance, ok := stage.DATATYPE_DEFINITION_ENUMERATIONs_instance[_reference]; ok {
			_DATATYPE_DEFINITION_ENUMERATION = append(_DATATYPE_DEFINITION_ENUMERATION, _instance)
		}
	}
	reference.DATATYPE_DEFINITION_ENUMERATION = _DATATYPE_DEFINITION_ENUMERATION
	var _DATATYPE_DEFINITION_INTEGER []*DATATYPE_DEFINITION_INTEGER
	for _, _reference := range reference.DATATYPE_DEFINITION_INTEGER {
		if _instance, ok := stage.DATATYPE_DEFINITION_INTEGERs_instance[_reference]; ok {
			_DATATYPE_DEFINITION_INTEGER = append(_DATATYPE_DEFINITION_INTEGER, _instance)
		}
	}
	reference.DATATYPE_DEFINITION_INTEGER = _DATATYPE_DEFINITION_INTEGER
	var _DATATYPE_DEFINITION_REAL []*DATATYPE_DEFINITION_REAL
	for _, _reference := range reference.DATATYPE_DEFINITION_REAL {
		if _instance, ok := stage.DATATYPE_DEFINITION_REALs_instance[_reference]; ok {
			_DATATYPE_DEFINITION_REAL = append(_DATATYPE_DEFINITION_REAL, _instance)
		}
	}
	reference.DATATYPE_DEFINITION_REAL = _DATATYPE_DEFINITION_REAL
	var _DATATYPE_DEFINITION_STRING []*DATATYPE_DEFINITION_STRING
	for _, _reference := range reference.DATATYPE_DEFINITION_STRING {
		if _instance, ok := stage.DATATYPE_DEFINITION_STRINGs_instance[_reference]; ok {
			_DATATYPE_DEFINITION_STRING = append(_DATATYPE_DEFINITION_STRING, _instance)
		}
	}
	reference.DATATYPE_DEFINITION_STRING = _DATATYPE_DEFINITION_STRING
	var _DATATYPE_DEFINITION_XHTML []*DATATYPE_DEFINITION_XHTML
	for _, _reference := range reference.DATATYPE_DEFINITION_XHTML {
		if _instance, ok := stage.DATATYPE_DEFINITION_XHTMLs_instance[_reference]; ok {
			_DATATYPE_DEFINITION_XHTML = append(_DATATYPE_DEFINITION_XHTML, _instance)
		}
	}
	reference.DATATYPE_DEFINITION_XHTML = _DATATYPE_DEFINITION_XHTML
}

func (reference *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_DATATYPE_DEFINITION_DATE_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_DATATYPE_DEFINITION_INTEGER_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_DATATYPE_DEFINITION_REAL_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_DATATYPE_DEFINITION_STRING_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_DATATYPE_DEFINITION_XHTML_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_EDITABLE_ATTS) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_ENUM_VALUE_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_OBJECT) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_PROPERTIES) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.EMBEDDED_VALUE; _reference != nil {
		reference.EMBEDDED_VALUE = nil
		if _instance, ok := stage.EMBEDDED_VALUEs_instance[_reference]; ok {
			reference.EMBEDDED_VALUE = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *A_RELATION_GROUP_TYPE_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_SOURCE_1) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_SOURCE_SPECIFICATION_1) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_SPECIFICATIONS) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _SPECIFICATION []*SPECIFICATION
	for _, _reference := range reference.SPECIFICATION {
		if _instance, ok := stage.SPECIFICATIONs_instance[_reference]; ok {
			_SPECIFICATION = append(_SPECIFICATION, _instance)
		}
	}
	reference.SPECIFICATION = _SPECIFICATION
}

func (reference *A_SPECIFICATION_TYPE_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_SPECIFIED_VALUES) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _ENUM_VALUE []*ENUM_VALUE
	for _, _reference := range reference.ENUM_VALUE {
		if _instance, ok := stage.ENUM_VALUEs_instance[_reference]; ok {
			_ENUM_VALUE = append(_ENUM_VALUE, _instance)
		}
	}
	reference.ENUM_VALUE = _ENUM_VALUE
}

func (reference *A_SPEC_ATTRIBUTES) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _ATTRIBUTE_DEFINITION_BOOLEAN []*ATTRIBUTE_DEFINITION_BOOLEAN
	for _, _reference := range reference.ATTRIBUTE_DEFINITION_BOOLEAN {
		if _instance, ok := stage.ATTRIBUTE_DEFINITION_BOOLEANs_instance[_reference]; ok {
			_ATTRIBUTE_DEFINITION_BOOLEAN = append(_ATTRIBUTE_DEFINITION_BOOLEAN, _instance)
		}
	}
	reference.ATTRIBUTE_DEFINITION_BOOLEAN = _ATTRIBUTE_DEFINITION_BOOLEAN
	var _ATTRIBUTE_DEFINITION_DATE []*ATTRIBUTE_DEFINITION_DATE
	for _, _reference := range reference.ATTRIBUTE_DEFINITION_DATE {
		if _instance, ok := stage.ATTRIBUTE_DEFINITION_DATEs_instance[_reference]; ok {
			_ATTRIBUTE_DEFINITION_DATE = append(_ATTRIBUTE_DEFINITION_DATE, _instance)
		}
	}
	reference.ATTRIBUTE_DEFINITION_DATE = _ATTRIBUTE_DEFINITION_DATE
	var _ATTRIBUTE_DEFINITION_ENUMERATION []*ATTRIBUTE_DEFINITION_ENUMERATION
	for _, _reference := range reference.ATTRIBUTE_DEFINITION_ENUMERATION {
		if _instance, ok := stage.ATTRIBUTE_DEFINITION_ENUMERATIONs_instance[_reference]; ok {
			_ATTRIBUTE_DEFINITION_ENUMERATION = append(_ATTRIBUTE_DEFINITION_ENUMERATION, _instance)
		}
	}
	reference.ATTRIBUTE_DEFINITION_ENUMERATION = _ATTRIBUTE_DEFINITION_ENUMERATION
	var _ATTRIBUTE_DEFINITION_INTEGER []*ATTRIBUTE_DEFINITION_INTEGER
	for _, _reference := range reference.ATTRIBUTE_DEFINITION_INTEGER {
		if _instance, ok := stage.ATTRIBUTE_DEFINITION_INTEGERs_instance[_reference]; ok {
			_ATTRIBUTE_DEFINITION_INTEGER = append(_ATTRIBUTE_DEFINITION_INTEGER, _instance)
		}
	}
	reference.ATTRIBUTE_DEFINITION_INTEGER = _ATTRIBUTE_DEFINITION_INTEGER
	var _ATTRIBUTE_DEFINITION_REAL []*ATTRIBUTE_DEFINITION_REAL
	for _, _reference := range reference.ATTRIBUTE_DEFINITION_REAL {
		if _instance, ok := stage.ATTRIBUTE_DEFINITION_REALs_instance[_reference]; ok {
			_ATTRIBUTE_DEFINITION_REAL = append(_ATTRIBUTE_DEFINITION_REAL, _instance)
		}
	}
	reference.ATTRIBUTE_DEFINITION_REAL = _ATTRIBUTE_DEFINITION_REAL
	var _ATTRIBUTE_DEFINITION_STRING []*ATTRIBUTE_DEFINITION_STRING
	for _, _reference := range reference.ATTRIBUTE_DEFINITION_STRING {
		if _instance, ok := stage.ATTRIBUTE_DEFINITION_STRINGs_instance[_reference]; ok {
			_ATTRIBUTE_DEFINITION_STRING = append(_ATTRIBUTE_DEFINITION_STRING, _instance)
		}
	}
	reference.ATTRIBUTE_DEFINITION_STRING = _ATTRIBUTE_DEFINITION_STRING
	var _ATTRIBUTE_DEFINITION_XHTML []*ATTRIBUTE_DEFINITION_XHTML
	for _, _reference := range reference.ATTRIBUTE_DEFINITION_XHTML {
		if _instance, ok := stage.ATTRIBUTE_DEFINITION_XHTMLs_instance[_reference]; ok {
			_ATTRIBUTE_DEFINITION_XHTML = append(_ATTRIBUTE_DEFINITION_XHTML, _instance)
		}
	}
	reference.ATTRIBUTE_DEFINITION_XHTML = _ATTRIBUTE_DEFINITION_XHTML
}

func (reference *A_SPEC_OBJECTS) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _SPEC_OBJECT []*SPEC_OBJECT
	for _, _reference := range reference.SPEC_OBJECT {
		if _instance, ok := stage.SPEC_OBJECTs_instance[_reference]; ok {
			_SPEC_OBJECT = append(_SPEC_OBJECT, _instance)
		}
	}
	reference.SPEC_OBJECT = _SPEC_OBJECT
}

func (reference *A_SPEC_OBJECT_TYPE_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_SPEC_RELATIONS) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _SPEC_RELATION []*SPEC_RELATION
	for _, _reference := range reference.SPEC_RELATION {
		if _instance, ok := stage.SPEC_RELATIONs_instance[_reference]; ok {
			_SPEC_RELATION = append(_SPEC_RELATION, _instance)
		}
	}
	reference.SPEC_RELATION = _SPEC_RELATION
}

func (reference *A_SPEC_RELATION_GROUPS) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _RELATION_GROUP []*RELATION_GROUP
	for _, _reference := range reference.RELATION_GROUP {
		if _instance, ok := stage.RELATION_GROUPs_instance[_reference]; ok {
			_RELATION_GROUP = append(_RELATION_GROUP, _instance)
		}
	}
	reference.RELATION_GROUP = _RELATION_GROUP
}

func (reference *A_SPEC_RELATION_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_SPEC_RELATION_TYPE_REF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *A_SPEC_TYPES) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _RELATION_GROUP_TYPE []*RELATION_GROUP_TYPE
	for _, _reference := range reference.RELATION_GROUP_TYPE {
		if _instance, ok := stage.RELATION_GROUP_TYPEs_instance[_reference]; ok {
			_RELATION_GROUP_TYPE = append(_RELATION_GROUP_TYPE, _instance)
		}
	}
	reference.RELATION_GROUP_TYPE = _RELATION_GROUP_TYPE
	var _SPEC_OBJECT_TYPE []*SPEC_OBJECT_TYPE
	for _, _reference := range reference.SPEC_OBJECT_TYPE {
		if _instance, ok := stage.SPEC_OBJECT_TYPEs_instance[_reference]; ok {
			_SPEC_OBJECT_TYPE = append(_SPEC_OBJECT_TYPE, _instance)
		}
	}
	reference.SPEC_OBJECT_TYPE = _SPEC_OBJECT_TYPE
	var _SPEC_RELATION_TYPE []*SPEC_RELATION_TYPE
	for _, _reference := range reference.SPEC_RELATION_TYPE {
		if _instance, ok := stage.SPEC_RELATION_TYPEs_instance[_reference]; ok {
			_SPEC_RELATION_TYPE = append(_SPEC_RELATION_TYPE, _instance)
		}
	}
	reference.SPEC_RELATION_TYPE = _SPEC_RELATION_TYPE
	var _SPECIFICATION_TYPE []*SPECIFICATION_TYPE
	for _, _reference := range reference.SPECIFICATION_TYPE {
		if _instance, ok := stage.SPECIFICATION_TYPEs_instance[_reference]; ok {
			_SPECIFICATION_TYPE = append(_SPECIFICATION_TYPE, _instance)
		}
	}
	reference.SPECIFICATION_TYPE = _SPECIFICATION_TYPE
}

func (reference *A_THE_HEADER) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.REQ_IF_HEADER; _reference != nil {
		reference.REQ_IF_HEADER = nil
		if _instance, ok := stage.REQ_IF_HEADERs_instance[_reference]; ok {
			reference.REQ_IF_HEADER = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *A_TOOL_EXTENSIONS) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
	var _REQ_IF_TOOL_EXTENSION []*REQ_IF_TOOL_EXTENSION
	for _, _reference := range reference.REQ_IF_TOOL_EXTENSION {
		if _instance, ok := stage.REQ_IF_TOOL_EXTENSIONs_instance[_reference]; ok {
			_REQ_IF_TOOL_EXTENSION = append(_REQ_IF_TOOL_EXTENSION, _instance)
		}
	}
	reference.REQ_IF_TOOL_EXTENSION = _REQ_IF_TOOL_EXTENSION
}

func (reference *DATATYPE_DEFINITION_BOOLEAN) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *DATATYPE_DEFINITION_DATE) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *DATATYPE_DEFINITION_ENUMERATION) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.SPECIFIED_VALUES; _reference != nil {
		reference.SPECIFIED_VALUES = nil
		if _instance, ok := stage.A_SPECIFIED_VALUESs_instance[_reference]; ok {
			reference.SPECIFIED_VALUES = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *DATATYPE_DEFINITION_INTEGER) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *DATATYPE_DEFINITION_REAL) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *DATATYPE_DEFINITION_STRING) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *DATATYPE_DEFINITION_XHTML) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *EMBEDDED_VALUE) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *ENUM_VALUE) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.PROPERTIES; _reference != nil {
		reference.PROPERTIES = nil
		if _instance, ok := stage.A_PROPERTIESs_instance[_reference]; ok {
			reference.PROPERTIES = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *RELATION_GROUP) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.SOURCE_SPECIFICATION; _reference != nil {
		reference.SOURCE_SPECIFICATION = nil
		if _instance, ok := stage.A_SOURCE_SPECIFICATION_1s_instance[_reference]; ok {
			reference.SOURCE_SPECIFICATION = _instance
		}
	}
	if _reference := reference.SPEC_RELATIONS; _reference != nil {
		reference.SPEC_RELATIONS = nil
		if _instance, ok := stage.A_SPEC_RELATION_REFs_instance[_reference]; ok {
			reference.SPEC_RELATIONS = _instance
		}
	}
	if _reference := reference.TARGET_SPECIFICATION; _reference != nil {
		reference.TARGET_SPECIFICATION = nil
		if _instance, ok := stage.A_SOURCE_SPECIFICATION_1s_instance[_reference]; ok {
			reference.TARGET_SPECIFICATION = _instance
		}
	}
	if _reference := reference.TYPE; _reference != nil {
		reference.TYPE = nil
		if _instance, ok := stage.A_RELATION_GROUP_TYPE_REFs_instance[_reference]; ok {
			reference.TYPE = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *RELATION_GROUP_TYPE) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.SPEC_ATTRIBUTES; _reference != nil {
		reference.SPEC_ATTRIBUTES = nil
		if _instance, ok := stage.A_SPEC_ATTRIBUTESs_instance[_reference]; ok {
			reference.SPEC_ATTRIBUTES = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *REQ_IF) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.THE_HEADER; _reference != nil {
		reference.THE_HEADER = nil
		if _instance, ok := stage.A_THE_HEADERs_instance[_reference]; ok {
			reference.THE_HEADER = _instance
		}
	}
	if _reference := reference.CORE_CONTENT; _reference != nil {
		reference.CORE_CONTENT = nil
		if _instance, ok := stage.A_CORE_CONTENTs_instance[_reference]; ok {
			reference.CORE_CONTENT = _instance
		}
	}
	if _reference := reference.TOOL_EXTENSIONS; _reference != nil {
		reference.TOOL_EXTENSIONS = nil
		if _instance, ok := stage.A_TOOL_EXTENSIONSs_instance[_reference]; ok {
			reference.TOOL_EXTENSIONS = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *REQ_IF_CONTENT) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.DATATYPES; _reference != nil {
		reference.DATATYPES = nil
		if _instance, ok := stage.A_DATATYPESs_instance[_reference]; ok {
			reference.DATATYPES = _instance
		}
	}
	if _reference := reference.SPEC_TYPES; _reference != nil {
		reference.SPEC_TYPES = nil
		if _instance, ok := stage.A_SPEC_TYPESs_instance[_reference]; ok {
			reference.SPEC_TYPES = _instance
		}
	}
	if _reference := reference.SPEC_OBJECTS; _reference != nil {
		reference.SPEC_OBJECTS = nil
		if _instance, ok := stage.A_SPEC_OBJECTSs_instance[_reference]; ok {
			reference.SPEC_OBJECTS = _instance
		}
	}
	if _reference := reference.SPEC_RELATIONS; _reference != nil {
		reference.SPEC_RELATIONS = nil
		if _instance, ok := stage.A_SPEC_RELATIONSs_instance[_reference]; ok {
			reference.SPEC_RELATIONS = _instance
		}
	}
	if _reference := reference.SPECIFICATIONS; _reference != nil {
		reference.SPECIFICATIONS = nil
		if _instance, ok := stage.A_SPECIFICATIONSs_instance[_reference]; ok {
			reference.SPECIFICATIONS = _instance
		}
	}
	if _reference := reference.SPEC_RELATION_GROUPS; _reference != nil {
		reference.SPEC_RELATION_GROUPS = nil
		if _instance, ok := stage.A_SPEC_RELATION_GROUPSs_instance[_reference]; ok {
			reference.SPEC_RELATION_GROUPS = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *REQ_IF_HEADER) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *REQ_IF_TOOL_EXTENSION) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

func (reference *SPECIFICATION) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.CHILDREN; _reference != nil {
		reference.CHILDREN = nil
		if _instance, ok := stage.A_CHILDRENs_instance[_reference]; ok {
			reference.CHILDREN = _instance
		}
	}
	if _reference := reference.VALUES; _reference != nil {
		reference.VALUES = nil
		if _instance, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1s_instance[_reference]; ok {
			reference.VALUES = _instance
		}
	}
	if _reference := reference.TYPE; _reference != nil {
		reference.TYPE = nil
		if _instance, ok := stage.A_SPECIFICATION_TYPE_REFs_instance[_reference]; ok {
			reference.TYPE = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *SPECIFICATION_TYPE) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.SPEC_ATTRIBUTES; _reference != nil {
		reference.SPEC_ATTRIBUTES = nil
		if _instance, ok := stage.A_SPEC_ATTRIBUTESs_instance[_reference]; ok {
			reference.SPEC_ATTRIBUTES = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *SPEC_HIERARCHY) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.CHILDREN; _reference != nil {
		reference.CHILDREN = nil
		if _instance, ok := stage.A_CHILDRENs_instance[_reference]; ok {
			reference.CHILDREN = _instance
		}
	}
	if _reference := reference.EDITABLE_ATTS; _reference != nil {
		reference.EDITABLE_ATTS = nil
		if _instance, ok := stage.A_EDITABLE_ATTSs_instance[_reference]; ok {
			reference.EDITABLE_ATTS = _instance
		}
	}
	if _reference := reference.OBJECT; _reference != nil {
		reference.OBJECT = nil
		if _instance, ok := stage.A_OBJECTs_instance[_reference]; ok {
			reference.OBJECT = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *SPEC_OBJECT) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.VALUES; _reference != nil {
		reference.VALUES = nil
		if _instance, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1s_instance[_reference]; ok {
			reference.VALUES = _instance
		}
	}
	if _reference := reference.TYPE; _reference != nil {
		reference.TYPE = nil
		if _instance, ok := stage.A_SPEC_OBJECT_TYPE_REFs_instance[_reference]; ok {
			reference.TYPE = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *SPEC_OBJECT_TYPE) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.SPEC_ATTRIBUTES; _reference != nil {
		reference.SPEC_ATTRIBUTES = nil
		if _instance, ok := stage.A_SPEC_ATTRIBUTESs_instance[_reference]; ok {
			reference.SPEC_ATTRIBUTES = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *SPEC_RELATION) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.VALUES; _reference != nil {
		reference.VALUES = nil
		if _instance, ok := stage.A_ATTRIBUTE_VALUE_XHTML_1s_instance[_reference]; ok {
			reference.VALUES = _instance
		}
	}
	if _reference := reference.SOURCE; _reference != nil {
		reference.SOURCE = nil
		if _instance, ok := stage.A_SOURCE_1s_instance[_reference]; ok {
			reference.SOURCE = _instance
		}
	}
	if _reference := reference.TARGET; _reference != nil {
		reference.TARGET = nil
		if _instance, ok := stage.A_SOURCE_1s_instance[_reference]; ok {
			reference.TARGET = _instance
		}
	}
	if _reference := reference.TYPE; _reference != nil {
		reference.TYPE = nil
		if _instance, ok := stage.A_SPEC_RELATION_TYPE_REFs_instance[_reference]; ok {
			reference.TYPE = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *SPEC_RELATION_TYPE) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	if _reference := reference.ALTERNATIVE_ID; _reference != nil {
		reference.ALTERNATIVE_ID = nil
		if _instance, ok := stage.A_ALTERNATIVE_IDs_instance[_reference]; ok {
			reference.ALTERNATIVE_ID = _instance
		}
	}
	if _reference := reference.SPEC_ATTRIBUTES; _reference != nil {
		reference.SPEC_ATTRIBUTES = nil
		if _instance, ok := stage.A_SPEC_ATTRIBUTESs_instance[_reference]; ok {
			reference.SPEC_ATTRIBUTES = _instance
		}
	}
	// insertion point for slice of pointers fields
}

func (reference *XHTML_CONTENT) GongReconstructPointersFromInstances(stage *Stage) {
	// insertion point for pointers field
	// insertion point for slice of pointers fields
}

// insertion point for diff per struct
// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (alternative_id *ALTERNATIVE_ID) GongDiff(stage *Stage, alternative_idOther *ALTERNATIVE_ID) (diffs []string) {
	// insertion point for field diffs
	if alternative_id.Name != alternative_idOther.Name {
		diffs = append(diffs, alternative_id.GongMarshallField(stage, "Name"))
	}
	if alternative_id.IDENTIFIER != alternative_idOther.IDENTIFIER {
		diffs = append(diffs, alternative_id.GongMarshallField(stage, "IDENTIFIER"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attribute_definition_boolean *ATTRIBUTE_DEFINITION_BOOLEAN) GongDiff(stage *Stage, attribute_definition_booleanOther *ATTRIBUTE_DEFINITION_BOOLEAN) (diffs []string) {
	// insertion point for field diffs
	if attribute_definition_boolean.Name != attribute_definition_booleanOther.Name {
		diffs = append(diffs, attribute_definition_boolean.GongMarshallField(stage, "Name"))
	}
	if attribute_definition_boolean.DESC != attribute_definition_booleanOther.DESC {
		diffs = append(diffs, attribute_definition_boolean.GongMarshallField(stage, "DESC"))
	}
	if attribute_definition_boolean.IDENTIFIER != attribute_definition_booleanOther.IDENTIFIER {
		diffs = append(diffs, attribute_definition_boolean.GongMarshallField(stage, "IDENTIFIER"))
	}
	if attribute_definition_boolean.IS_EDITABLE != attribute_definition_booleanOther.IS_EDITABLE {
		diffs = append(diffs, attribute_definition_boolean.GongMarshallField(stage, "IS_EDITABLE"))
	}
	if attribute_definition_boolean.LAST_CHANGE != attribute_definition_booleanOther.LAST_CHANGE {
		diffs = append(diffs, attribute_definition_boolean.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if attribute_definition_boolean.LONG_NAME != attribute_definition_booleanOther.LONG_NAME {
		diffs = append(diffs, attribute_definition_boolean.GongMarshallField(stage, "LONG_NAME"))
	}
	if (attribute_definition_boolean.ALTERNATIVE_ID == nil) != (attribute_definition_booleanOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, attribute_definition_boolean.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if attribute_definition_boolean.ALTERNATIVE_ID != nil && attribute_definition_booleanOther.ALTERNATIVE_ID != nil {
		if attribute_definition_boolean.ALTERNATIVE_ID != attribute_definition_booleanOther.ALTERNATIVE_ID {
			diffs = append(diffs, attribute_definition_boolean.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (attribute_definition_boolean.DEFAULT_VALUE == nil) != (attribute_definition_booleanOther.DEFAULT_VALUE == nil) {
		diffs = append(diffs, attribute_definition_boolean.GongMarshallField(stage, "DEFAULT_VALUE"))
	} else if attribute_definition_boolean.DEFAULT_VALUE != nil && attribute_definition_booleanOther.DEFAULT_VALUE != nil {
		if attribute_definition_boolean.DEFAULT_VALUE != attribute_definition_booleanOther.DEFAULT_VALUE {
			diffs = append(diffs, attribute_definition_boolean.GongMarshallField(stage, "DEFAULT_VALUE"))
		}
	}
	if (attribute_definition_boolean.TYPE == nil) != (attribute_definition_booleanOther.TYPE == nil) {
		diffs = append(diffs, attribute_definition_boolean.GongMarshallField(stage, "TYPE"))
	} else if attribute_definition_boolean.TYPE != nil && attribute_definition_booleanOther.TYPE != nil {
		if attribute_definition_boolean.TYPE != attribute_definition_booleanOther.TYPE {
			diffs = append(diffs, attribute_definition_boolean.GongMarshallField(stage, "TYPE"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attribute_definition_date *ATTRIBUTE_DEFINITION_DATE) GongDiff(stage *Stage, attribute_definition_dateOther *ATTRIBUTE_DEFINITION_DATE) (diffs []string) {
	// insertion point for field diffs
	if attribute_definition_date.Name != attribute_definition_dateOther.Name {
		diffs = append(diffs, attribute_definition_date.GongMarshallField(stage, "Name"))
	}
	if attribute_definition_date.DESC != attribute_definition_dateOther.DESC {
		diffs = append(diffs, attribute_definition_date.GongMarshallField(stage, "DESC"))
	}
	if attribute_definition_date.IDENTIFIER != attribute_definition_dateOther.IDENTIFIER {
		diffs = append(diffs, attribute_definition_date.GongMarshallField(stage, "IDENTIFIER"))
	}
	if attribute_definition_date.IS_EDITABLE != attribute_definition_dateOther.IS_EDITABLE {
		diffs = append(diffs, attribute_definition_date.GongMarshallField(stage, "IS_EDITABLE"))
	}
	if attribute_definition_date.LAST_CHANGE != attribute_definition_dateOther.LAST_CHANGE {
		diffs = append(diffs, attribute_definition_date.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if attribute_definition_date.LONG_NAME != attribute_definition_dateOther.LONG_NAME {
		diffs = append(diffs, attribute_definition_date.GongMarshallField(stage, "LONG_NAME"))
	}
	if (attribute_definition_date.ALTERNATIVE_ID == nil) != (attribute_definition_dateOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, attribute_definition_date.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if attribute_definition_date.ALTERNATIVE_ID != nil && attribute_definition_dateOther.ALTERNATIVE_ID != nil {
		if attribute_definition_date.ALTERNATIVE_ID != attribute_definition_dateOther.ALTERNATIVE_ID {
			diffs = append(diffs, attribute_definition_date.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (attribute_definition_date.DEFAULT_VALUE == nil) != (attribute_definition_dateOther.DEFAULT_VALUE == nil) {
		diffs = append(diffs, attribute_definition_date.GongMarshallField(stage, "DEFAULT_VALUE"))
	} else if attribute_definition_date.DEFAULT_VALUE != nil && attribute_definition_dateOther.DEFAULT_VALUE != nil {
		if attribute_definition_date.DEFAULT_VALUE != attribute_definition_dateOther.DEFAULT_VALUE {
			diffs = append(diffs, attribute_definition_date.GongMarshallField(stage, "DEFAULT_VALUE"))
		}
	}
	if (attribute_definition_date.TYPE == nil) != (attribute_definition_dateOther.TYPE == nil) {
		diffs = append(diffs, attribute_definition_date.GongMarshallField(stage, "TYPE"))
	} else if attribute_definition_date.TYPE != nil && attribute_definition_dateOther.TYPE != nil {
		if attribute_definition_date.TYPE != attribute_definition_dateOther.TYPE {
			diffs = append(diffs, attribute_definition_date.GongMarshallField(stage, "TYPE"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attribute_definition_enumeration *ATTRIBUTE_DEFINITION_ENUMERATION) GongDiff(stage *Stage, attribute_definition_enumerationOther *ATTRIBUTE_DEFINITION_ENUMERATION) (diffs []string) {
	// insertion point for field diffs
	if attribute_definition_enumeration.Name != attribute_definition_enumerationOther.Name {
		diffs = append(diffs, attribute_definition_enumeration.GongMarshallField(stage, "Name"))
	}
	if attribute_definition_enumeration.DESC != attribute_definition_enumerationOther.DESC {
		diffs = append(diffs, attribute_definition_enumeration.GongMarshallField(stage, "DESC"))
	}
	if attribute_definition_enumeration.IDENTIFIER != attribute_definition_enumerationOther.IDENTIFIER {
		diffs = append(diffs, attribute_definition_enumeration.GongMarshallField(stage, "IDENTIFIER"))
	}
	if attribute_definition_enumeration.IS_EDITABLE != attribute_definition_enumerationOther.IS_EDITABLE {
		diffs = append(diffs, attribute_definition_enumeration.GongMarshallField(stage, "IS_EDITABLE"))
	}
	if attribute_definition_enumeration.LAST_CHANGE != attribute_definition_enumerationOther.LAST_CHANGE {
		diffs = append(diffs, attribute_definition_enumeration.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if attribute_definition_enumeration.LONG_NAME != attribute_definition_enumerationOther.LONG_NAME {
		diffs = append(diffs, attribute_definition_enumeration.GongMarshallField(stage, "LONG_NAME"))
	}
	if attribute_definition_enumeration.MULTI_VALUED != attribute_definition_enumerationOther.MULTI_VALUED {
		diffs = append(diffs, attribute_definition_enumeration.GongMarshallField(stage, "MULTI_VALUED"))
	}
	if (attribute_definition_enumeration.ALTERNATIVE_ID == nil) != (attribute_definition_enumerationOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, attribute_definition_enumeration.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if attribute_definition_enumeration.ALTERNATIVE_ID != nil && attribute_definition_enumerationOther.ALTERNATIVE_ID != nil {
		if attribute_definition_enumeration.ALTERNATIVE_ID != attribute_definition_enumerationOther.ALTERNATIVE_ID {
			diffs = append(diffs, attribute_definition_enumeration.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (attribute_definition_enumeration.DEFAULT_VALUE == nil) != (attribute_definition_enumerationOther.DEFAULT_VALUE == nil) {
		diffs = append(diffs, attribute_definition_enumeration.GongMarshallField(stage, "DEFAULT_VALUE"))
	} else if attribute_definition_enumeration.DEFAULT_VALUE != nil && attribute_definition_enumerationOther.DEFAULT_VALUE != nil {
		if attribute_definition_enumeration.DEFAULT_VALUE != attribute_definition_enumerationOther.DEFAULT_VALUE {
			diffs = append(diffs, attribute_definition_enumeration.GongMarshallField(stage, "DEFAULT_VALUE"))
		}
	}
	if (attribute_definition_enumeration.TYPE == nil) != (attribute_definition_enumerationOther.TYPE == nil) {
		diffs = append(diffs, attribute_definition_enumeration.GongMarshallField(stage, "TYPE"))
	} else if attribute_definition_enumeration.TYPE != nil && attribute_definition_enumerationOther.TYPE != nil {
		if attribute_definition_enumeration.TYPE != attribute_definition_enumerationOther.TYPE {
			diffs = append(diffs, attribute_definition_enumeration.GongMarshallField(stage, "TYPE"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attribute_definition_integer *ATTRIBUTE_DEFINITION_INTEGER) GongDiff(stage *Stage, attribute_definition_integerOther *ATTRIBUTE_DEFINITION_INTEGER) (diffs []string) {
	// insertion point for field diffs
	if attribute_definition_integer.Name != attribute_definition_integerOther.Name {
		diffs = append(diffs, attribute_definition_integer.GongMarshallField(stage, "Name"))
	}
	if attribute_definition_integer.DESC != attribute_definition_integerOther.DESC {
		diffs = append(diffs, attribute_definition_integer.GongMarshallField(stage, "DESC"))
	}
	if attribute_definition_integer.IDENTIFIER != attribute_definition_integerOther.IDENTIFIER {
		diffs = append(diffs, attribute_definition_integer.GongMarshallField(stage, "IDENTIFIER"))
	}
	if attribute_definition_integer.IS_EDITABLE != attribute_definition_integerOther.IS_EDITABLE {
		diffs = append(diffs, attribute_definition_integer.GongMarshallField(stage, "IS_EDITABLE"))
	}
	if attribute_definition_integer.LAST_CHANGE != attribute_definition_integerOther.LAST_CHANGE {
		diffs = append(diffs, attribute_definition_integer.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if attribute_definition_integer.LONG_NAME != attribute_definition_integerOther.LONG_NAME {
		diffs = append(diffs, attribute_definition_integer.GongMarshallField(stage, "LONG_NAME"))
	}
	if (attribute_definition_integer.ALTERNATIVE_ID == nil) != (attribute_definition_integerOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, attribute_definition_integer.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if attribute_definition_integer.ALTERNATIVE_ID != nil && attribute_definition_integerOther.ALTERNATIVE_ID != nil {
		if attribute_definition_integer.ALTERNATIVE_ID != attribute_definition_integerOther.ALTERNATIVE_ID {
			diffs = append(diffs, attribute_definition_integer.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (attribute_definition_integer.DEFAULT_VALUE == nil) != (attribute_definition_integerOther.DEFAULT_VALUE == nil) {
		diffs = append(diffs, attribute_definition_integer.GongMarshallField(stage, "DEFAULT_VALUE"))
	} else if attribute_definition_integer.DEFAULT_VALUE != nil && attribute_definition_integerOther.DEFAULT_VALUE != nil {
		if attribute_definition_integer.DEFAULT_VALUE != attribute_definition_integerOther.DEFAULT_VALUE {
			diffs = append(diffs, attribute_definition_integer.GongMarshallField(stage, "DEFAULT_VALUE"))
		}
	}
	if (attribute_definition_integer.TYPE == nil) != (attribute_definition_integerOther.TYPE == nil) {
		diffs = append(diffs, attribute_definition_integer.GongMarshallField(stage, "TYPE"))
	} else if attribute_definition_integer.TYPE != nil && attribute_definition_integerOther.TYPE != nil {
		if attribute_definition_integer.TYPE != attribute_definition_integerOther.TYPE {
			diffs = append(diffs, attribute_definition_integer.GongMarshallField(stage, "TYPE"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attribute_definition_real *ATTRIBUTE_DEFINITION_REAL) GongDiff(stage *Stage, attribute_definition_realOther *ATTRIBUTE_DEFINITION_REAL) (diffs []string) {
	// insertion point for field diffs
	if attribute_definition_real.Name != attribute_definition_realOther.Name {
		diffs = append(diffs, attribute_definition_real.GongMarshallField(stage, "Name"))
	}
	if attribute_definition_real.DESC != attribute_definition_realOther.DESC {
		diffs = append(diffs, attribute_definition_real.GongMarshallField(stage, "DESC"))
	}
	if attribute_definition_real.IDENTIFIER != attribute_definition_realOther.IDENTIFIER {
		diffs = append(diffs, attribute_definition_real.GongMarshallField(stage, "IDENTIFIER"))
	}
	if attribute_definition_real.IS_EDITABLE != attribute_definition_realOther.IS_EDITABLE {
		diffs = append(diffs, attribute_definition_real.GongMarshallField(stage, "IS_EDITABLE"))
	}
	if attribute_definition_real.LAST_CHANGE != attribute_definition_realOther.LAST_CHANGE {
		diffs = append(diffs, attribute_definition_real.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if attribute_definition_real.LONG_NAME != attribute_definition_realOther.LONG_NAME {
		diffs = append(diffs, attribute_definition_real.GongMarshallField(stage, "LONG_NAME"))
	}
	if (attribute_definition_real.ALTERNATIVE_ID == nil) != (attribute_definition_realOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, attribute_definition_real.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if attribute_definition_real.ALTERNATIVE_ID != nil && attribute_definition_realOther.ALTERNATIVE_ID != nil {
		if attribute_definition_real.ALTERNATIVE_ID != attribute_definition_realOther.ALTERNATIVE_ID {
			diffs = append(diffs, attribute_definition_real.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (attribute_definition_real.DEFAULT_VALUE == nil) != (attribute_definition_realOther.DEFAULT_VALUE == nil) {
		diffs = append(diffs, attribute_definition_real.GongMarshallField(stage, "DEFAULT_VALUE"))
	} else if attribute_definition_real.DEFAULT_VALUE != nil && attribute_definition_realOther.DEFAULT_VALUE != nil {
		if attribute_definition_real.DEFAULT_VALUE != attribute_definition_realOther.DEFAULT_VALUE {
			diffs = append(diffs, attribute_definition_real.GongMarshallField(stage, "DEFAULT_VALUE"))
		}
	}
	if (attribute_definition_real.TYPE == nil) != (attribute_definition_realOther.TYPE == nil) {
		diffs = append(diffs, attribute_definition_real.GongMarshallField(stage, "TYPE"))
	} else if attribute_definition_real.TYPE != nil && attribute_definition_realOther.TYPE != nil {
		if attribute_definition_real.TYPE != attribute_definition_realOther.TYPE {
			diffs = append(diffs, attribute_definition_real.GongMarshallField(stage, "TYPE"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attribute_definition_string *ATTRIBUTE_DEFINITION_STRING) GongDiff(stage *Stage, attribute_definition_stringOther *ATTRIBUTE_DEFINITION_STRING) (diffs []string) {
	// insertion point for field diffs
	if attribute_definition_string.Name != attribute_definition_stringOther.Name {
		diffs = append(diffs, attribute_definition_string.GongMarshallField(stage, "Name"))
	}
	if attribute_definition_string.DESC != attribute_definition_stringOther.DESC {
		diffs = append(diffs, attribute_definition_string.GongMarshallField(stage, "DESC"))
	}
	if attribute_definition_string.IDENTIFIER != attribute_definition_stringOther.IDENTIFIER {
		diffs = append(diffs, attribute_definition_string.GongMarshallField(stage, "IDENTIFIER"))
	}
	if attribute_definition_string.IS_EDITABLE != attribute_definition_stringOther.IS_EDITABLE {
		diffs = append(diffs, attribute_definition_string.GongMarshallField(stage, "IS_EDITABLE"))
	}
	if attribute_definition_string.LAST_CHANGE != attribute_definition_stringOther.LAST_CHANGE {
		diffs = append(diffs, attribute_definition_string.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if attribute_definition_string.LONG_NAME != attribute_definition_stringOther.LONG_NAME {
		diffs = append(diffs, attribute_definition_string.GongMarshallField(stage, "LONG_NAME"))
	}
	if (attribute_definition_string.ALTERNATIVE_ID == nil) != (attribute_definition_stringOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, attribute_definition_string.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if attribute_definition_string.ALTERNATIVE_ID != nil && attribute_definition_stringOther.ALTERNATIVE_ID != nil {
		if attribute_definition_string.ALTERNATIVE_ID != attribute_definition_stringOther.ALTERNATIVE_ID {
			diffs = append(diffs, attribute_definition_string.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (attribute_definition_string.DEFAULT_VALUE == nil) != (attribute_definition_stringOther.DEFAULT_VALUE == nil) {
		diffs = append(diffs, attribute_definition_string.GongMarshallField(stage, "DEFAULT_VALUE"))
	} else if attribute_definition_string.DEFAULT_VALUE != nil && attribute_definition_stringOther.DEFAULT_VALUE != nil {
		if attribute_definition_string.DEFAULT_VALUE != attribute_definition_stringOther.DEFAULT_VALUE {
			diffs = append(diffs, attribute_definition_string.GongMarshallField(stage, "DEFAULT_VALUE"))
		}
	}
	if (attribute_definition_string.TYPE == nil) != (attribute_definition_stringOther.TYPE == nil) {
		diffs = append(diffs, attribute_definition_string.GongMarshallField(stage, "TYPE"))
	} else if attribute_definition_string.TYPE != nil && attribute_definition_stringOther.TYPE != nil {
		if attribute_definition_string.TYPE != attribute_definition_stringOther.TYPE {
			diffs = append(diffs, attribute_definition_string.GongMarshallField(stage, "TYPE"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attribute_definition_xhtml *ATTRIBUTE_DEFINITION_XHTML) GongDiff(stage *Stage, attribute_definition_xhtmlOther *ATTRIBUTE_DEFINITION_XHTML) (diffs []string) {
	// insertion point for field diffs
	if attribute_definition_xhtml.Name != attribute_definition_xhtmlOther.Name {
		diffs = append(diffs, attribute_definition_xhtml.GongMarshallField(stage, "Name"))
	}
	if attribute_definition_xhtml.DESC != attribute_definition_xhtmlOther.DESC {
		diffs = append(diffs, attribute_definition_xhtml.GongMarshallField(stage, "DESC"))
	}
	if attribute_definition_xhtml.IDENTIFIER != attribute_definition_xhtmlOther.IDENTIFIER {
		diffs = append(diffs, attribute_definition_xhtml.GongMarshallField(stage, "IDENTIFIER"))
	}
	if attribute_definition_xhtml.IS_EDITABLE != attribute_definition_xhtmlOther.IS_EDITABLE {
		diffs = append(diffs, attribute_definition_xhtml.GongMarshallField(stage, "IS_EDITABLE"))
	}
	if attribute_definition_xhtml.LAST_CHANGE != attribute_definition_xhtmlOther.LAST_CHANGE {
		diffs = append(diffs, attribute_definition_xhtml.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if attribute_definition_xhtml.LONG_NAME != attribute_definition_xhtmlOther.LONG_NAME {
		diffs = append(diffs, attribute_definition_xhtml.GongMarshallField(stage, "LONG_NAME"))
	}
	if (attribute_definition_xhtml.ALTERNATIVE_ID == nil) != (attribute_definition_xhtmlOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, attribute_definition_xhtml.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if attribute_definition_xhtml.ALTERNATIVE_ID != nil && attribute_definition_xhtmlOther.ALTERNATIVE_ID != nil {
		if attribute_definition_xhtml.ALTERNATIVE_ID != attribute_definition_xhtmlOther.ALTERNATIVE_ID {
			diffs = append(diffs, attribute_definition_xhtml.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (attribute_definition_xhtml.DEFAULT_VALUE == nil) != (attribute_definition_xhtmlOther.DEFAULT_VALUE == nil) {
		diffs = append(diffs, attribute_definition_xhtml.GongMarshallField(stage, "DEFAULT_VALUE"))
	} else if attribute_definition_xhtml.DEFAULT_VALUE != nil && attribute_definition_xhtmlOther.DEFAULT_VALUE != nil {
		if attribute_definition_xhtml.DEFAULT_VALUE != attribute_definition_xhtmlOther.DEFAULT_VALUE {
			diffs = append(diffs, attribute_definition_xhtml.GongMarshallField(stage, "DEFAULT_VALUE"))
		}
	}
	if (attribute_definition_xhtml.TYPE == nil) != (attribute_definition_xhtmlOther.TYPE == nil) {
		diffs = append(diffs, attribute_definition_xhtml.GongMarshallField(stage, "TYPE"))
	} else if attribute_definition_xhtml.TYPE != nil && attribute_definition_xhtmlOther.TYPE != nil {
		if attribute_definition_xhtml.TYPE != attribute_definition_xhtmlOther.TYPE {
			diffs = append(diffs, attribute_definition_xhtml.GongMarshallField(stage, "TYPE"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attribute_value_boolean *ATTRIBUTE_VALUE_BOOLEAN) GongDiff(stage *Stage, attribute_value_booleanOther *ATTRIBUTE_VALUE_BOOLEAN) (diffs []string) {
	// insertion point for field diffs
	if attribute_value_boolean.Name != attribute_value_booleanOther.Name {
		diffs = append(diffs, attribute_value_boolean.GongMarshallField(stage, "Name"))
	}
	if attribute_value_boolean.THE_VALUE != attribute_value_booleanOther.THE_VALUE {
		diffs = append(diffs, attribute_value_boolean.GongMarshallField(stage, "THE_VALUE"))
	}
	if (attribute_value_boolean.DEFINITION == nil) != (attribute_value_booleanOther.DEFINITION == nil) {
		diffs = append(diffs, attribute_value_boolean.GongMarshallField(stage, "DEFINITION"))
	} else if attribute_value_boolean.DEFINITION != nil && attribute_value_booleanOther.DEFINITION != nil {
		if attribute_value_boolean.DEFINITION != attribute_value_booleanOther.DEFINITION {
			diffs = append(diffs, attribute_value_boolean.GongMarshallField(stage, "DEFINITION"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attribute_value_date *ATTRIBUTE_VALUE_DATE) GongDiff(stage *Stage, attribute_value_dateOther *ATTRIBUTE_VALUE_DATE) (diffs []string) {
	// insertion point for field diffs
	if attribute_value_date.Name != attribute_value_dateOther.Name {
		diffs = append(diffs, attribute_value_date.GongMarshallField(stage, "Name"))
	}
	if attribute_value_date.THE_VALUE != attribute_value_dateOther.THE_VALUE {
		diffs = append(diffs, attribute_value_date.GongMarshallField(stage, "THE_VALUE"))
	}
	if (attribute_value_date.DEFINITION == nil) != (attribute_value_dateOther.DEFINITION == nil) {
		diffs = append(diffs, attribute_value_date.GongMarshallField(stage, "DEFINITION"))
	} else if attribute_value_date.DEFINITION != nil && attribute_value_dateOther.DEFINITION != nil {
		if attribute_value_date.DEFINITION != attribute_value_dateOther.DEFINITION {
			diffs = append(diffs, attribute_value_date.GongMarshallField(stage, "DEFINITION"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attribute_value_enumeration *ATTRIBUTE_VALUE_ENUMERATION) GongDiff(stage *Stage, attribute_value_enumerationOther *ATTRIBUTE_VALUE_ENUMERATION) (diffs []string) {
	// insertion point for field diffs
	if attribute_value_enumeration.Name != attribute_value_enumerationOther.Name {
		diffs = append(diffs, attribute_value_enumeration.GongMarshallField(stage, "Name"))
	}
	if (attribute_value_enumeration.DEFINITION == nil) != (attribute_value_enumerationOther.DEFINITION == nil) {
		diffs = append(diffs, attribute_value_enumeration.GongMarshallField(stage, "DEFINITION"))
	} else if attribute_value_enumeration.DEFINITION != nil && attribute_value_enumerationOther.DEFINITION != nil {
		if attribute_value_enumeration.DEFINITION != attribute_value_enumerationOther.DEFINITION {
			diffs = append(diffs, attribute_value_enumeration.GongMarshallField(stage, "DEFINITION"))
		}
	}
	if (attribute_value_enumeration.VALUES == nil) != (attribute_value_enumerationOther.VALUES == nil) {
		diffs = append(diffs, attribute_value_enumeration.GongMarshallField(stage, "VALUES"))
	} else if attribute_value_enumeration.VALUES != nil && attribute_value_enumerationOther.VALUES != nil {
		if attribute_value_enumeration.VALUES != attribute_value_enumerationOther.VALUES {
			diffs = append(diffs, attribute_value_enumeration.GongMarshallField(stage, "VALUES"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attribute_value_integer *ATTRIBUTE_VALUE_INTEGER) GongDiff(stage *Stage, attribute_value_integerOther *ATTRIBUTE_VALUE_INTEGER) (diffs []string) {
	// insertion point for field diffs
	if attribute_value_integer.Name != attribute_value_integerOther.Name {
		diffs = append(diffs, attribute_value_integer.GongMarshallField(stage, "Name"))
	}
	if attribute_value_integer.THE_VALUE != attribute_value_integerOther.THE_VALUE {
		diffs = append(diffs, attribute_value_integer.GongMarshallField(stage, "THE_VALUE"))
	}
	if (attribute_value_integer.DEFINITION == nil) != (attribute_value_integerOther.DEFINITION == nil) {
		diffs = append(diffs, attribute_value_integer.GongMarshallField(stage, "DEFINITION"))
	} else if attribute_value_integer.DEFINITION != nil && attribute_value_integerOther.DEFINITION != nil {
		if attribute_value_integer.DEFINITION != attribute_value_integerOther.DEFINITION {
			diffs = append(diffs, attribute_value_integer.GongMarshallField(stage, "DEFINITION"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attribute_value_real *ATTRIBUTE_VALUE_REAL) GongDiff(stage *Stage, attribute_value_realOther *ATTRIBUTE_VALUE_REAL) (diffs []string) {
	// insertion point for field diffs
	if attribute_value_real.Name != attribute_value_realOther.Name {
		diffs = append(diffs, attribute_value_real.GongMarshallField(stage, "Name"))
	}
	if attribute_value_real.THE_VALUE != attribute_value_realOther.THE_VALUE {
		diffs = append(diffs, attribute_value_real.GongMarshallField(stage, "THE_VALUE"))
	}
	if (attribute_value_real.DEFINITION == nil) != (attribute_value_realOther.DEFINITION == nil) {
		diffs = append(diffs, attribute_value_real.GongMarshallField(stage, "DEFINITION"))
	} else if attribute_value_real.DEFINITION != nil && attribute_value_realOther.DEFINITION != nil {
		if attribute_value_real.DEFINITION != attribute_value_realOther.DEFINITION {
			diffs = append(diffs, attribute_value_real.GongMarshallField(stage, "DEFINITION"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attribute_value_string *ATTRIBUTE_VALUE_STRING) GongDiff(stage *Stage, attribute_value_stringOther *ATTRIBUTE_VALUE_STRING) (diffs []string) {
	// insertion point for field diffs
	if attribute_value_string.Name != attribute_value_stringOther.Name {
		diffs = append(diffs, attribute_value_string.GongMarshallField(stage, "Name"))
	}
	if attribute_value_string.THE_VALUE != attribute_value_stringOther.THE_VALUE {
		diffs = append(diffs, attribute_value_string.GongMarshallField(stage, "THE_VALUE"))
	}
	if (attribute_value_string.DEFINITION == nil) != (attribute_value_stringOther.DEFINITION == nil) {
		diffs = append(diffs, attribute_value_string.GongMarshallField(stage, "DEFINITION"))
	} else if attribute_value_string.DEFINITION != nil && attribute_value_stringOther.DEFINITION != nil {
		if attribute_value_string.DEFINITION != attribute_value_stringOther.DEFINITION {
			diffs = append(diffs, attribute_value_string.GongMarshallField(stage, "DEFINITION"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (attribute_value_xhtml *ATTRIBUTE_VALUE_XHTML) GongDiff(stage *Stage, attribute_value_xhtmlOther *ATTRIBUTE_VALUE_XHTML) (diffs []string) {
	// insertion point for field diffs
	if attribute_value_xhtml.Name != attribute_value_xhtmlOther.Name {
		diffs = append(diffs, attribute_value_xhtml.GongMarshallField(stage, "Name"))
	}
	if attribute_value_xhtml.IS_SIMPLIFIED != attribute_value_xhtmlOther.IS_SIMPLIFIED {
		diffs = append(diffs, attribute_value_xhtml.GongMarshallField(stage, "IS_SIMPLIFIED"))
	}
	if (attribute_value_xhtml.THE_VALUE == nil) != (attribute_value_xhtmlOther.THE_VALUE == nil) {
		diffs = append(diffs, attribute_value_xhtml.GongMarshallField(stage, "THE_VALUE"))
	} else if attribute_value_xhtml.THE_VALUE != nil && attribute_value_xhtmlOther.THE_VALUE != nil {
		if attribute_value_xhtml.THE_VALUE != attribute_value_xhtmlOther.THE_VALUE {
			diffs = append(diffs, attribute_value_xhtml.GongMarshallField(stage, "THE_VALUE"))
		}
	}
	if (attribute_value_xhtml.THE_ORIGINAL_VALUE == nil) != (attribute_value_xhtmlOther.THE_ORIGINAL_VALUE == nil) {
		diffs = append(diffs, attribute_value_xhtml.GongMarshallField(stage, "THE_ORIGINAL_VALUE"))
	} else if attribute_value_xhtml.THE_ORIGINAL_VALUE != nil && attribute_value_xhtmlOther.THE_ORIGINAL_VALUE != nil {
		if attribute_value_xhtml.THE_ORIGINAL_VALUE != attribute_value_xhtmlOther.THE_ORIGINAL_VALUE {
			diffs = append(diffs, attribute_value_xhtml.GongMarshallField(stage, "THE_ORIGINAL_VALUE"))
		}
	}
	if (attribute_value_xhtml.DEFINITION == nil) != (attribute_value_xhtmlOther.DEFINITION == nil) {
		diffs = append(diffs, attribute_value_xhtml.GongMarshallField(stage, "DEFINITION"))
	} else if attribute_value_xhtml.DEFINITION != nil && attribute_value_xhtmlOther.DEFINITION != nil {
		if attribute_value_xhtml.DEFINITION != attribute_value_xhtmlOther.DEFINITION {
			diffs = append(diffs, attribute_value_xhtml.GongMarshallField(stage, "DEFINITION"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_alternative_id *A_ALTERNATIVE_ID) GongDiff(stage *Stage, a_alternative_idOther *A_ALTERNATIVE_ID) (diffs []string) {
	// insertion point for field diffs
	if a_alternative_id.Name != a_alternative_idOther.Name {
		diffs = append(diffs, a_alternative_id.GongMarshallField(stage, "Name"))
	}
	if (a_alternative_id.ALTERNATIVE_ID == nil) != (a_alternative_idOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, a_alternative_id.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if a_alternative_id.ALTERNATIVE_ID != nil && a_alternative_idOther.ALTERNATIVE_ID != nil {
		if a_alternative_id.ALTERNATIVE_ID != a_alternative_idOther.ALTERNATIVE_ID {
			diffs = append(diffs, a_alternative_id.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_attribute_definition_boolean_ref *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) GongDiff(stage *Stage, a_attribute_definition_boolean_refOther *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF) (diffs []string) {
	// insertion point for field diffs
	if a_attribute_definition_boolean_ref.Name != a_attribute_definition_boolean_refOther.Name {
		diffs = append(diffs, a_attribute_definition_boolean_ref.GongMarshallField(stage, "Name"))
	}
	if a_attribute_definition_boolean_ref.ATTRIBUTE_DEFINITION_BOOLEAN_REF != a_attribute_definition_boolean_refOther.ATTRIBUTE_DEFINITION_BOOLEAN_REF {
		diffs = append(diffs, a_attribute_definition_boolean_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_BOOLEAN_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_attribute_definition_date_ref *A_ATTRIBUTE_DEFINITION_DATE_REF) GongDiff(stage *Stage, a_attribute_definition_date_refOther *A_ATTRIBUTE_DEFINITION_DATE_REF) (diffs []string) {
	// insertion point for field diffs
	if a_attribute_definition_date_ref.Name != a_attribute_definition_date_refOther.Name {
		diffs = append(diffs, a_attribute_definition_date_ref.GongMarshallField(stage, "Name"))
	}
	if a_attribute_definition_date_ref.ATTRIBUTE_DEFINITION_DATE_REF != a_attribute_definition_date_refOther.ATTRIBUTE_DEFINITION_DATE_REF {
		diffs = append(diffs, a_attribute_definition_date_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_DATE_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_attribute_definition_enumeration_ref *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) GongDiff(stage *Stage, a_attribute_definition_enumeration_refOther *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF) (diffs []string) {
	// insertion point for field diffs
	if a_attribute_definition_enumeration_ref.Name != a_attribute_definition_enumeration_refOther.Name {
		diffs = append(diffs, a_attribute_definition_enumeration_ref.GongMarshallField(stage, "Name"))
	}
	if a_attribute_definition_enumeration_ref.ATTRIBUTE_DEFINITION_ENUMERATION_REF != a_attribute_definition_enumeration_refOther.ATTRIBUTE_DEFINITION_ENUMERATION_REF {
		diffs = append(diffs, a_attribute_definition_enumeration_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_ENUMERATION_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_attribute_definition_integer_ref *A_ATTRIBUTE_DEFINITION_INTEGER_REF) GongDiff(stage *Stage, a_attribute_definition_integer_refOther *A_ATTRIBUTE_DEFINITION_INTEGER_REF) (diffs []string) {
	// insertion point for field diffs
	if a_attribute_definition_integer_ref.Name != a_attribute_definition_integer_refOther.Name {
		diffs = append(diffs, a_attribute_definition_integer_ref.GongMarshallField(stage, "Name"))
	}
	if a_attribute_definition_integer_ref.ATTRIBUTE_DEFINITION_INTEGER_REF != a_attribute_definition_integer_refOther.ATTRIBUTE_DEFINITION_INTEGER_REF {
		diffs = append(diffs, a_attribute_definition_integer_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_INTEGER_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_attribute_definition_real_ref *A_ATTRIBUTE_DEFINITION_REAL_REF) GongDiff(stage *Stage, a_attribute_definition_real_refOther *A_ATTRIBUTE_DEFINITION_REAL_REF) (diffs []string) {
	// insertion point for field diffs
	if a_attribute_definition_real_ref.Name != a_attribute_definition_real_refOther.Name {
		diffs = append(diffs, a_attribute_definition_real_ref.GongMarshallField(stage, "Name"))
	}
	if a_attribute_definition_real_ref.ATTRIBUTE_DEFINITION_REAL_REF != a_attribute_definition_real_refOther.ATTRIBUTE_DEFINITION_REAL_REF {
		diffs = append(diffs, a_attribute_definition_real_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_REAL_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_attribute_definition_string_ref *A_ATTRIBUTE_DEFINITION_STRING_REF) GongDiff(stage *Stage, a_attribute_definition_string_refOther *A_ATTRIBUTE_DEFINITION_STRING_REF) (diffs []string) {
	// insertion point for field diffs
	if a_attribute_definition_string_ref.Name != a_attribute_definition_string_refOther.Name {
		diffs = append(diffs, a_attribute_definition_string_ref.GongMarshallField(stage, "Name"))
	}
	if a_attribute_definition_string_ref.ATTRIBUTE_DEFINITION_STRING_REF != a_attribute_definition_string_refOther.ATTRIBUTE_DEFINITION_STRING_REF {
		diffs = append(diffs, a_attribute_definition_string_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_STRING_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_attribute_definition_xhtml_ref *A_ATTRIBUTE_DEFINITION_XHTML_REF) GongDiff(stage *Stage, a_attribute_definition_xhtml_refOther *A_ATTRIBUTE_DEFINITION_XHTML_REF) (diffs []string) {
	// insertion point for field diffs
	if a_attribute_definition_xhtml_ref.Name != a_attribute_definition_xhtml_refOther.Name {
		diffs = append(diffs, a_attribute_definition_xhtml_ref.GongMarshallField(stage, "Name"))
	}
	if a_attribute_definition_xhtml_ref.ATTRIBUTE_DEFINITION_XHTML_REF != a_attribute_definition_xhtml_refOther.ATTRIBUTE_DEFINITION_XHTML_REF {
		diffs = append(diffs, a_attribute_definition_xhtml_ref.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_XHTML_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_attribute_value_boolean *A_ATTRIBUTE_VALUE_BOOLEAN) GongDiff(stage *Stage, a_attribute_value_booleanOther *A_ATTRIBUTE_VALUE_BOOLEAN) (diffs []string) {
	// insertion point for field diffs
	if a_attribute_value_boolean.Name != a_attribute_value_booleanOther.Name {
		diffs = append(diffs, a_attribute_value_boolean.GongMarshallField(stage, "Name"))
	}
	ATTRIBUTE_VALUE_BOOLEANDifferent := false
	if len(a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN) != len(a_attribute_value_booleanOther.ATTRIBUTE_VALUE_BOOLEAN) {
		ATTRIBUTE_VALUE_BOOLEANDifferent = true
	} else {
		for i := range a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN {
			if (a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN[i] == nil) != (a_attribute_value_booleanOther.ATTRIBUTE_VALUE_BOOLEAN[i] == nil) {
				ATTRIBUTE_VALUE_BOOLEANDifferent = true
				break
			} else if a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN[i] != nil && a_attribute_value_booleanOther.ATTRIBUTE_VALUE_BOOLEAN[i] != nil {
				// this is a pointer comparaison
				if a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN[i] != a_attribute_value_booleanOther.ATTRIBUTE_VALUE_BOOLEAN[i] {
					ATTRIBUTE_VALUE_BOOLEANDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_VALUE_BOOLEANDifferent {
		ops := stage.Diff(
			a_attribute_value_boolean,
			"ATTRIBUTE_VALUE_BOOLEAN",
			len(a_attribute_value_booleanOther.ATTRIBUTE_VALUE_BOOLEAN),
			len(a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN),
			func(i, j int) bool {
				return a_attribute_value_booleanOther.ATTRIBUTE_VALUE_BOOLEAN[i] == a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN[j]
			},
			func(j int) string {
				return a_attribute_value_boolean.ATTRIBUTE_VALUE_BOOLEAN[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_attribute_value_date *A_ATTRIBUTE_VALUE_DATE) GongDiff(stage *Stage, a_attribute_value_dateOther *A_ATTRIBUTE_VALUE_DATE) (diffs []string) {
	// insertion point for field diffs
	if a_attribute_value_date.Name != a_attribute_value_dateOther.Name {
		diffs = append(diffs, a_attribute_value_date.GongMarshallField(stage, "Name"))
	}
	ATTRIBUTE_VALUE_DATEDifferent := false
	if len(a_attribute_value_date.ATTRIBUTE_VALUE_DATE) != len(a_attribute_value_dateOther.ATTRIBUTE_VALUE_DATE) {
		ATTRIBUTE_VALUE_DATEDifferent = true
	} else {
		for i := range a_attribute_value_date.ATTRIBUTE_VALUE_DATE {
			if (a_attribute_value_date.ATTRIBUTE_VALUE_DATE[i] == nil) != (a_attribute_value_dateOther.ATTRIBUTE_VALUE_DATE[i] == nil) {
				ATTRIBUTE_VALUE_DATEDifferent = true
				break
			} else if a_attribute_value_date.ATTRIBUTE_VALUE_DATE[i] != nil && a_attribute_value_dateOther.ATTRIBUTE_VALUE_DATE[i] != nil {
				// this is a pointer comparaison
				if a_attribute_value_date.ATTRIBUTE_VALUE_DATE[i] != a_attribute_value_dateOther.ATTRIBUTE_VALUE_DATE[i] {
					ATTRIBUTE_VALUE_DATEDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_VALUE_DATEDifferent {
		ops := stage.Diff(
			a_attribute_value_date,
			"ATTRIBUTE_VALUE_DATE",
			len(a_attribute_value_dateOther.ATTRIBUTE_VALUE_DATE),
			len(a_attribute_value_date.ATTRIBUTE_VALUE_DATE),
			func(i, j int) bool {
				return a_attribute_value_dateOther.ATTRIBUTE_VALUE_DATE[i] == a_attribute_value_date.ATTRIBUTE_VALUE_DATE[j]
			},
			func(j int) string {
				return a_attribute_value_date.ATTRIBUTE_VALUE_DATE[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_attribute_value_enumeration *A_ATTRIBUTE_VALUE_ENUMERATION) GongDiff(stage *Stage, a_attribute_value_enumerationOther *A_ATTRIBUTE_VALUE_ENUMERATION) (diffs []string) {
	// insertion point for field diffs
	if a_attribute_value_enumeration.Name != a_attribute_value_enumerationOther.Name {
		diffs = append(diffs, a_attribute_value_enumeration.GongMarshallField(stage, "Name"))
	}
	ATTRIBUTE_VALUE_ENUMERATIONDifferent := false
	if len(a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION) != len(a_attribute_value_enumerationOther.ATTRIBUTE_VALUE_ENUMERATION) {
		ATTRIBUTE_VALUE_ENUMERATIONDifferent = true
	} else {
		for i := range a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION {
			if (a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION[i] == nil) != (a_attribute_value_enumerationOther.ATTRIBUTE_VALUE_ENUMERATION[i] == nil) {
				ATTRIBUTE_VALUE_ENUMERATIONDifferent = true
				break
			} else if a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION[i] != nil && a_attribute_value_enumerationOther.ATTRIBUTE_VALUE_ENUMERATION[i] != nil {
				// this is a pointer comparaison
				if a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION[i] != a_attribute_value_enumerationOther.ATTRIBUTE_VALUE_ENUMERATION[i] {
					ATTRIBUTE_VALUE_ENUMERATIONDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_VALUE_ENUMERATIONDifferent {
		ops := stage.Diff(
			a_attribute_value_enumeration,
			"ATTRIBUTE_VALUE_ENUMERATION",
			len(a_attribute_value_enumerationOther.ATTRIBUTE_VALUE_ENUMERATION),
			len(a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION),
			func(i, j int) bool {
				return a_attribute_value_enumerationOther.ATTRIBUTE_VALUE_ENUMERATION[i] == a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION[j]
			},
			func(j int) string {
				return a_attribute_value_enumeration.ATTRIBUTE_VALUE_ENUMERATION[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_attribute_value_integer *A_ATTRIBUTE_VALUE_INTEGER) GongDiff(stage *Stage, a_attribute_value_integerOther *A_ATTRIBUTE_VALUE_INTEGER) (diffs []string) {
	// insertion point for field diffs
	if a_attribute_value_integer.Name != a_attribute_value_integerOther.Name {
		diffs = append(diffs, a_attribute_value_integer.GongMarshallField(stage, "Name"))
	}
	ATTRIBUTE_VALUE_INTEGERDifferent := false
	if len(a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER) != len(a_attribute_value_integerOther.ATTRIBUTE_VALUE_INTEGER) {
		ATTRIBUTE_VALUE_INTEGERDifferent = true
	} else {
		for i := range a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER {
			if (a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER[i] == nil) != (a_attribute_value_integerOther.ATTRIBUTE_VALUE_INTEGER[i] == nil) {
				ATTRIBUTE_VALUE_INTEGERDifferent = true
				break
			} else if a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER[i] != nil && a_attribute_value_integerOther.ATTRIBUTE_VALUE_INTEGER[i] != nil {
				// this is a pointer comparaison
				if a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER[i] != a_attribute_value_integerOther.ATTRIBUTE_VALUE_INTEGER[i] {
					ATTRIBUTE_VALUE_INTEGERDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_VALUE_INTEGERDifferent {
		ops := stage.Diff(
			a_attribute_value_integer,
			"ATTRIBUTE_VALUE_INTEGER",
			len(a_attribute_value_integerOther.ATTRIBUTE_VALUE_INTEGER),
			len(a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER),
			func(i, j int) bool {
				return a_attribute_value_integerOther.ATTRIBUTE_VALUE_INTEGER[i] == a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER[j]
			},
			func(j int) string {
				return a_attribute_value_integer.ATTRIBUTE_VALUE_INTEGER[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_attribute_value_real *A_ATTRIBUTE_VALUE_REAL) GongDiff(stage *Stage, a_attribute_value_realOther *A_ATTRIBUTE_VALUE_REAL) (diffs []string) {
	// insertion point for field diffs
	if a_attribute_value_real.Name != a_attribute_value_realOther.Name {
		diffs = append(diffs, a_attribute_value_real.GongMarshallField(stage, "Name"))
	}
	ATTRIBUTE_VALUE_REALDifferent := false
	if len(a_attribute_value_real.ATTRIBUTE_VALUE_REAL) != len(a_attribute_value_realOther.ATTRIBUTE_VALUE_REAL) {
		ATTRIBUTE_VALUE_REALDifferent = true
	} else {
		for i := range a_attribute_value_real.ATTRIBUTE_VALUE_REAL {
			if (a_attribute_value_real.ATTRIBUTE_VALUE_REAL[i] == nil) != (a_attribute_value_realOther.ATTRIBUTE_VALUE_REAL[i] == nil) {
				ATTRIBUTE_VALUE_REALDifferent = true
				break
			} else if a_attribute_value_real.ATTRIBUTE_VALUE_REAL[i] != nil && a_attribute_value_realOther.ATTRIBUTE_VALUE_REAL[i] != nil {
				// this is a pointer comparaison
				if a_attribute_value_real.ATTRIBUTE_VALUE_REAL[i] != a_attribute_value_realOther.ATTRIBUTE_VALUE_REAL[i] {
					ATTRIBUTE_VALUE_REALDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_VALUE_REALDifferent {
		ops := stage.Diff(
			a_attribute_value_real,
			"ATTRIBUTE_VALUE_REAL",
			len(a_attribute_value_realOther.ATTRIBUTE_VALUE_REAL),
			len(a_attribute_value_real.ATTRIBUTE_VALUE_REAL),
			func(i, j int) bool {
				return a_attribute_value_realOther.ATTRIBUTE_VALUE_REAL[i] == a_attribute_value_real.ATTRIBUTE_VALUE_REAL[j]
			},
			func(j int) string {
				return a_attribute_value_real.ATTRIBUTE_VALUE_REAL[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_attribute_value_string *A_ATTRIBUTE_VALUE_STRING) GongDiff(stage *Stage, a_attribute_value_stringOther *A_ATTRIBUTE_VALUE_STRING) (diffs []string) {
	// insertion point for field diffs
	if a_attribute_value_string.Name != a_attribute_value_stringOther.Name {
		diffs = append(diffs, a_attribute_value_string.GongMarshallField(stage, "Name"))
	}
	ATTRIBUTE_VALUE_STRINGDifferent := false
	if len(a_attribute_value_string.ATTRIBUTE_VALUE_STRING) != len(a_attribute_value_stringOther.ATTRIBUTE_VALUE_STRING) {
		ATTRIBUTE_VALUE_STRINGDifferent = true
	} else {
		for i := range a_attribute_value_string.ATTRIBUTE_VALUE_STRING {
			if (a_attribute_value_string.ATTRIBUTE_VALUE_STRING[i] == nil) != (a_attribute_value_stringOther.ATTRIBUTE_VALUE_STRING[i] == nil) {
				ATTRIBUTE_VALUE_STRINGDifferent = true
				break
			} else if a_attribute_value_string.ATTRIBUTE_VALUE_STRING[i] != nil && a_attribute_value_stringOther.ATTRIBUTE_VALUE_STRING[i] != nil {
				// this is a pointer comparaison
				if a_attribute_value_string.ATTRIBUTE_VALUE_STRING[i] != a_attribute_value_stringOther.ATTRIBUTE_VALUE_STRING[i] {
					ATTRIBUTE_VALUE_STRINGDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_VALUE_STRINGDifferent {
		ops := stage.Diff(
			a_attribute_value_string,
			"ATTRIBUTE_VALUE_STRING",
			len(a_attribute_value_stringOther.ATTRIBUTE_VALUE_STRING),
			len(a_attribute_value_string.ATTRIBUTE_VALUE_STRING),
			func(i, j int) bool {
				return a_attribute_value_stringOther.ATTRIBUTE_VALUE_STRING[i] == a_attribute_value_string.ATTRIBUTE_VALUE_STRING[j]
			},
			func(j int) string {
				return a_attribute_value_string.ATTRIBUTE_VALUE_STRING[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_attribute_value_xhtml *A_ATTRIBUTE_VALUE_XHTML) GongDiff(stage *Stage, a_attribute_value_xhtmlOther *A_ATTRIBUTE_VALUE_XHTML) (diffs []string) {
	// insertion point for field diffs
	if a_attribute_value_xhtml.Name != a_attribute_value_xhtmlOther.Name {
		diffs = append(diffs, a_attribute_value_xhtml.GongMarshallField(stage, "Name"))
	}
	ATTRIBUTE_VALUE_XHTMLDifferent := false
	if len(a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML) != len(a_attribute_value_xhtmlOther.ATTRIBUTE_VALUE_XHTML) {
		ATTRIBUTE_VALUE_XHTMLDifferent = true
	} else {
		for i := range a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML {
			if (a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML[i] == nil) != (a_attribute_value_xhtmlOther.ATTRIBUTE_VALUE_XHTML[i] == nil) {
				ATTRIBUTE_VALUE_XHTMLDifferent = true
				break
			} else if a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML[i] != nil && a_attribute_value_xhtmlOther.ATTRIBUTE_VALUE_XHTML[i] != nil {
				// this is a pointer comparaison
				if a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML[i] != a_attribute_value_xhtmlOther.ATTRIBUTE_VALUE_XHTML[i] {
					ATTRIBUTE_VALUE_XHTMLDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_VALUE_XHTMLDifferent {
		ops := stage.Diff(
			a_attribute_value_xhtml,
			"ATTRIBUTE_VALUE_XHTML",
			len(a_attribute_value_xhtmlOther.ATTRIBUTE_VALUE_XHTML),
			len(a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML),
			func(i, j int) bool {
				return a_attribute_value_xhtmlOther.ATTRIBUTE_VALUE_XHTML[i] == a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML[j]
			},
			func(j int) string {
				return a_attribute_value_xhtml.ATTRIBUTE_VALUE_XHTML[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_attribute_value_xhtml_1 *A_ATTRIBUTE_VALUE_XHTML_1) GongDiff(stage *Stage, a_attribute_value_xhtml_1Other *A_ATTRIBUTE_VALUE_XHTML_1) (diffs []string) {
	// insertion point for field diffs
	if a_attribute_value_xhtml_1.Name != a_attribute_value_xhtml_1Other.Name {
		diffs = append(diffs, a_attribute_value_xhtml_1.GongMarshallField(stage, "Name"))
	}
	ATTRIBUTE_VALUE_BOOLEANDifferent := false
	if len(a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN) != len(a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_BOOLEAN) {
		ATTRIBUTE_VALUE_BOOLEANDifferent = true
	} else {
		for i := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN {
			if (a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN[i] == nil) != (a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_BOOLEAN[i] == nil) {
				ATTRIBUTE_VALUE_BOOLEANDifferent = true
				break
			} else if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN[i] != nil && a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_BOOLEAN[i] != nil {
				// this is a pointer comparaison
				if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN[i] != a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_BOOLEAN[i] {
					ATTRIBUTE_VALUE_BOOLEANDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_VALUE_BOOLEANDifferent {
		ops := stage.Diff(
			a_attribute_value_xhtml_1,
			"ATTRIBUTE_VALUE_BOOLEAN",
			len(a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_BOOLEAN),
			len(a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN),
			func(i, j int) bool {
				return a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_BOOLEAN[i] == a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN[j]
			},
			func(j int) string {
				return a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_BOOLEAN[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ATTRIBUTE_VALUE_DATEDifferent := false
	if len(a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE) != len(a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_DATE) {
		ATTRIBUTE_VALUE_DATEDifferent = true
	} else {
		for i := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE {
			if (a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE[i] == nil) != (a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_DATE[i] == nil) {
				ATTRIBUTE_VALUE_DATEDifferent = true
				break
			} else if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE[i] != nil && a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_DATE[i] != nil {
				// this is a pointer comparaison
				if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE[i] != a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_DATE[i] {
					ATTRIBUTE_VALUE_DATEDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_VALUE_DATEDifferent {
		ops := stage.Diff(
			a_attribute_value_xhtml_1,
			"ATTRIBUTE_VALUE_DATE",
			len(a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_DATE),
			len(a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE),
			func(i, j int) bool {
				return a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_DATE[i] == a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE[j]
			},
			func(j int) string {
				return a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_DATE[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ATTRIBUTE_VALUE_ENUMERATIONDifferent := false
	if len(a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION) != len(a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_ENUMERATION) {
		ATTRIBUTE_VALUE_ENUMERATIONDifferent = true
	} else {
		for i := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION {
			if (a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION[i] == nil) != (a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_ENUMERATION[i] == nil) {
				ATTRIBUTE_VALUE_ENUMERATIONDifferent = true
				break
			} else if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION[i] != nil && a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_ENUMERATION[i] != nil {
				// this is a pointer comparaison
				if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION[i] != a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_ENUMERATION[i] {
					ATTRIBUTE_VALUE_ENUMERATIONDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_VALUE_ENUMERATIONDifferent {
		ops := stage.Diff(
			a_attribute_value_xhtml_1,
			"ATTRIBUTE_VALUE_ENUMERATION",
			len(a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_ENUMERATION),
			len(a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION),
			func(i, j int) bool {
				return a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_ENUMERATION[i] == a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION[j]
			},
			func(j int) string {
				return a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_ENUMERATION[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ATTRIBUTE_VALUE_INTEGERDifferent := false
	if len(a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER) != len(a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_INTEGER) {
		ATTRIBUTE_VALUE_INTEGERDifferent = true
	} else {
		for i := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER {
			if (a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER[i] == nil) != (a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_INTEGER[i] == nil) {
				ATTRIBUTE_VALUE_INTEGERDifferent = true
				break
			} else if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER[i] != nil && a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_INTEGER[i] != nil {
				// this is a pointer comparaison
				if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER[i] != a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_INTEGER[i] {
					ATTRIBUTE_VALUE_INTEGERDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_VALUE_INTEGERDifferent {
		ops := stage.Diff(
			a_attribute_value_xhtml_1,
			"ATTRIBUTE_VALUE_INTEGER",
			len(a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_INTEGER),
			len(a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER),
			func(i, j int) bool {
				return a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_INTEGER[i] == a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER[j]
			},
			func(j int) string {
				return a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_INTEGER[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ATTRIBUTE_VALUE_REALDifferent := false
	if len(a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL) != len(a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_REAL) {
		ATTRIBUTE_VALUE_REALDifferent = true
	} else {
		for i := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL {
			if (a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL[i] == nil) != (a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_REAL[i] == nil) {
				ATTRIBUTE_VALUE_REALDifferent = true
				break
			} else if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL[i] != nil && a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_REAL[i] != nil {
				// this is a pointer comparaison
				if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL[i] != a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_REAL[i] {
					ATTRIBUTE_VALUE_REALDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_VALUE_REALDifferent {
		ops := stage.Diff(
			a_attribute_value_xhtml_1,
			"ATTRIBUTE_VALUE_REAL",
			len(a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_REAL),
			len(a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL),
			func(i, j int) bool {
				return a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_REAL[i] == a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL[j]
			},
			func(j int) string {
				return a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_REAL[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ATTRIBUTE_VALUE_STRINGDifferent := false
	if len(a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING) != len(a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_STRING) {
		ATTRIBUTE_VALUE_STRINGDifferent = true
	} else {
		for i := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING {
			if (a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING[i] == nil) != (a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_STRING[i] == nil) {
				ATTRIBUTE_VALUE_STRINGDifferent = true
				break
			} else if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING[i] != nil && a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_STRING[i] != nil {
				// this is a pointer comparaison
				if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING[i] != a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_STRING[i] {
					ATTRIBUTE_VALUE_STRINGDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_VALUE_STRINGDifferent {
		ops := stage.Diff(
			a_attribute_value_xhtml_1,
			"ATTRIBUTE_VALUE_STRING",
			len(a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_STRING),
			len(a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING),
			func(i, j int) bool {
				return a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_STRING[i] == a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING[j]
			},
			func(j int) string {
				return a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_STRING[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ATTRIBUTE_VALUE_XHTMLDifferent := false
	if len(a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML) != len(a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_XHTML) {
		ATTRIBUTE_VALUE_XHTMLDifferent = true
	} else {
		for i := range a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML {
			if (a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML[i] == nil) != (a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_XHTML[i] == nil) {
				ATTRIBUTE_VALUE_XHTMLDifferent = true
				break
			} else if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML[i] != nil && a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_XHTML[i] != nil {
				// this is a pointer comparaison
				if a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML[i] != a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_XHTML[i] {
					ATTRIBUTE_VALUE_XHTMLDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_VALUE_XHTMLDifferent {
		ops := stage.Diff(
			a_attribute_value_xhtml_1,
			"ATTRIBUTE_VALUE_XHTML",
			len(a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_XHTML),
			len(a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML),
			func(i, j int) bool {
				return a_attribute_value_xhtml_1Other.ATTRIBUTE_VALUE_XHTML[i] == a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML[j]
			},
			func(j int) string {
				return a_attribute_value_xhtml_1.ATTRIBUTE_VALUE_XHTML[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_children *A_CHILDREN) GongDiff(stage *Stage, a_childrenOther *A_CHILDREN) (diffs []string) {
	// insertion point for field diffs
	if a_children.Name != a_childrenOther.Name {
		diffs = append(diffs, a_children.GongMarshallField(stage, "Name"))
	}
	SPEC_HIERARCHYDifferent := false
	if len(a_children.SPEC_HIERARCHY) != len(a_childrenOther.SPEC_HIERARCHY) {
		SPEC_HIERARCHYDifferent = true
	} else {
		for i := range a_children.SPEC_HIERARCHY {
			if (a_children.SPEC_HIERARCHY[i] == nil) != (a_childrenOther.SPEC_HIERARCHY[i] == nil) {
				SPEC_HIERARCHYDifferent = true
				break
			} else if a_children.SPEC_HIERARCHY[i] != nil && a_childrenOther.SPEC_HIERARCHY[i] != nil {
				// this is a pointer comparaison
				if a_children.SPEC_HIERARCHY[i] != a_childrenOther.SPEC_HIERARCHY[i] {
					SPEC_HIERARCHYDifferent = true
					break
				}
			}
		}
	}
	if SPEC_HIERARCHYDifferent {
		ops := stage.Diff(
			a_children,
			"SPEC_HIERARCHY",
			len(a_childrenOther.SPEC_HIERARCHY),
			len(a_children.SPEC_HIERARCHY),
			func(i, j int) bool {
				return a_childrenOther.SPEC_HIERARCHY[i] == a_children.SPEC_HIERARCHY[j]
			},
			func(j int) string {
				return a_children.SPEC_HIERARCHY[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_core_content *A_CORE_CONTENT) GongDiff(stage *Stage, a_core_contentOther *A_CORE_CONTENT) (diffs []string) {
	// insertion point for field diffs
	if a_core_content.Name != a_core_contentOther.Name {
		diffs = append(diffs, a_core_content.GongMarshallField(stage, "Name"))
	}
	if (a_core_content.REQ_IF_CONTENT == nil) != (a_core_contentOther.REQ_IF_CONTENT == nil) {
		diffs = append(diffs, a_core_content.GongMarshallField(stage, "REQ_IF_CONTENT"))
	} else if a_core_content.REQ_IF_CONTENT != nil && a_core_contentOther.REQ_IF_CONTENT != nil {
		if a_core_content.REQ_IF_CONTENT != a_core_contentOther.REQ_IF_CONTENT {
			diffs = append(diffs, a_core_content.GongMarshallField(stage, "REQ_IF_CONTENT"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_datatypes *A_DATATYPES) GongDiff(stage *Stage, a_datatypesOther *A_DATATYPES) (diffs []string) {
	// insertion point for field diffs
	if a_datatypes.Name != a_datatypesOther.Name {
		diffs = append(diffs, a_datatypes.GongMarshallField(stage, "Name"))
	}
	DATATYPE_DEFINITION_BOOLEANDifferent := false
	if len(a_datatypes.DATATYPE_DEFINITION_BOOLEAN) != len(a_datatypesOther.DATATYPE_DEFINITION_BOOLEAN) {
		DATATYPE_DEFINITION_BOOLEANDifferent = true
	} else {
		for i := range a_datatypes.DATATYPE_DEFINITION_BOOLEAN {
			if (a_datatypes.DATATYPE_DEFINITION_BOOLEAN[i] == nil) != (a_datatypesOther.DATATYPE_DEFINITION_BOOLEAN[i] == nil) {
				DATATYPE_DEFINITION_BOOLEANDifferent = true
				break
			} else if a_datatypes.DATATYPE_DEFINITION_BOOLEAN[i] != nil && a_datatypesOther.DATATYPE_DEFINITION_BOOLEAN[i] != nil {
				// this is a pointer comparaison
				if a_datatypes.DATATYPE_DEFINITION_BOOLEAN[i] != a_datatypesOther.DATATYPE_DEFINITION_BOOLEAN[i] {
					DATATYPE_DEFINITION_BOOLEANDifferent = true
					break
				}
			}
		}
	}
	if DATATYPE_DEFINITION_BOOLEANDifferent {
		ops := stage.Diff(
			a_datatypes,
			"DATATYPE_DEFINITION_BOOLEAN",
			len(a_datatypesOther.DATATYPE_DEFINITION_BOOLEAN),
			len(a_datatypes.DATATYPE_DEFINITION_BOOLEAN),
			func(i, j int) bool {
				return a_datatypesOther.DATATYPE_DEFINITION_BOOLEAN[i] == a_datatypes.DATATYPE_DEFINITION_BOOLEAN[j]
			},
			func(j int) string {
				return a_datatypes.DATATYPE_DEFINITION_BOOLEAN[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DATATYPE_DEFINITION_DATEDifferent := false
	if len(a_datatypes.DATATYPE_DEFINITION_DATE) != len(a_datatypesOther.DATATYPE_DEFINITION_DATE) {
		DATATYPE_DEFINITION_DATEDifferent = true
	} else {
		for i := range a_datatypes.DATATYPE_DEFINITION_DATE {
			if (a_datatypes.DATATYPE_DEFINITION_DATE[i] == nil) != (a_datatypesOther.DATATYPE_DEFINITION_DATE[i] == nil) {
				DATATYPE_DEFINITION_DATEDifferent = true
				break
			} else if a_datatypes.DATATYPE_DEFINITION_DATE[i] != nil && a_datatypesOther.DATATYPE_DEFINITION_DATE[i] != nil {
				// this is a pointer comparaison
				if a_datatypes.DATATYPE_DEFINITION_DATE[i] != a_datatypesOther.DATATYPE_DEFINITION_DATE[i] {
					DATATYPE_DEFINITION_DATEDifferent = true
					break
				}
			}
		}
	}
	if DATATYPE_DEFINITION_DATEDifferent {
		ops := stage.Diff(
			a_datatypes,
			"DATATYPE_DEFINITION_DATE",
			len(a_datatypesOther.DATATYPE_DEFINITION_DATE),
			len(a_datatypes.DATATYPE_DEFINITION_DATE),
			func(i, j int) bool {
				return a_datatypesOther.DATATYPE_DEFINITION_DATE[i] == a_datatypes.DATATYPE_DEFINITION_DATE[j]
			},
			func(j int) string {
				return a_datatypes.DATATYPE_DEFINITION_DATE[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DATATYPE_DEFINITION_ENUMERATIONDifferent := false
	if len(a_datatypes.DATATYPE_DEFINITION_ENUMERATION) != len(a_datatypesOther.DATATYPE_DEFINITION_ENUMERATION) {
		DATATYPE_DEFINITION_ENUMERATIONDifferent = true
	} else {
		for i := range a_datatypes.DATATYPE_DEFINITION_ENUMERATION {
			if (a_datatypes.DATATYPE_DEFINITION_ENUMERATION[i] == nil) != (a_datatypesOther.DATATYPE_DEFINITION_ENUMERATION[i] == nil) {
				DATATYPE_DEFINITION_ENUMERATIONDifferent = true
				break
			} else if a_datatypes.DATATYPE_DEFINITION_ENUMERATION[i] != nil && a_datatypesOther.DATATYPE_DEFINITION_ENUMERATION[i] != nil {
				// this is a pointer comparaison
				if a_datatypes.DATATYPE_DEFINITION_ENUMERATION[i] != a_datatypesOther.DATATYPE_DEFINITION_ENUMERATION[i] {
					DATATYPE_DEFINITION_ENUMERATIONDifferent = true
					break
				}
			}
		}
	}
	if DATATYPE_DEFINITION_ENUMERATIONDifferent {
		ops := stage.Diff(
			a_datatypes,
			"DATATYPE_DEFINITION_ENUMERATION",
			len(a_datatypesOther.DATATYPE_DEFINITION_ENUMERATION),
			len(a_datatypes.DATATYPE_DEFINITION_ENUMERATION),
			func(i, j int) bool {
				return a_datatypesOther.DATATYPE_DEFINITION_ENUMERATION[i] == a_datatypes.DATATYPE_DEFINITION_ENUMERATION[j]
			},
			func(j int) string {
				return a_datatypes.DATATYPE_DEFINITION_ENUMERATION[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DATATYPE_DEFINITION_INTEGERDifferent := false
	if len(a_datatypes.DATATYPE_DEFINITION_INTEGER) != len(a_datatypesOther.DATATYPE_DEFINITION_INTEGER) {
		DATATYPE_DEFINITION_INTEGERDifferent = true
	} else {
		for i := range a_datatypes.DATATYPE_DEFINITION_INTEGER {
			if (a_datatypes.DATATYPE_DEFINITION_INTEGER[i] == nil) != (a_datatypesOther.DATATYPE_DEFINITION_INTEGER[i] == nil) {
				DATATYPE_DEFINITION_INTEGERDifferent = true
				break
			} else if a_datatypes.DATATYPE_DEFINITION_INTEGER[i] != nil && a_datatypesOther.DATATYPE_DEFINITION_INTEGER[i] != nil {
				// this is a pointer comparaison
				if a_datatypes.DATATYPE_DEFINITION_INTEGER[i] != a_datatypesOther.DATATYPE_DEFINITION_INTEGER[i] {
					DATATYPE_DEFINITION_INTEGERDifferent = true
					break
				}
			}
		}
	}
	if DATATYPE_DEFINITION_INTEGERDifferent {
		ops := stage.Diff(
			a_datatypes,
			"DATATYPE_DEFINITION_INTEGER",
			len(a_datatypesOther.DATATYPE_DEFINITION_INTEGER),
			len(a_datatypes.DATATYPE_DEFINITION_INTEGER),
			func(i, j int) bool {
				return a_datatypesOther.DATATYPE_DEFINITION_INTEGER[i] == a_datatypes.DATATYPE_DEFINITION_INTEGER[j]
			},
			func(j int) string {
				return a_datatypes.DATATYPE_DEFINITION_INTEGER[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DATATYPE_DEFINITION_REALDifferent := false
	if len(a_datatypes.DATATYPE_DEFINITION_REAL) != len(a_datatypesOther.DATATYPE_DEFINITION_REAL) {
		DATATYPE_DEFINITION_REALDifferent = true
	} else {
		for i := range a_datatypes.DATATYPE_DEFINITION_REAL {
			if (a_datatypes.DATATYPE_DEFINITION_REAL[i] == nil) != (a_datatypesOther.DATATYPE_DEFINITION_REAL[i] == nil) {
				DATATYPE_DEFINITION_REALDifferent = true
				break
			} else if a_datatypes.DATATYPE_DEFINITION_REAL[i] != nil && a_datatypesOther.DATATYPE_DEFINITION_REAL[i] != nil {
				// this is a pointer comparaison
				if a_datatypes.DATATYPE_DEFINITION_REAL[i] != a_datatypesOther.DATATYPE_DEFINITION_REAL[i] {
					DATATYPE_DEFINITION_REALDifferent = true
					break
				}
			}
		}
	}
	if DATATYPE_DEFINITION_REALDifferent {
		ops := stage.Diff(
			a_datatypes,
			"DATATYPE_DEFINITION_REAL",
			len(a_datatypesOther.DATATYPE_DEFINITION_REAL),
			len(a_datatypes.DATATYPE_DEFINITION_REAL),
			func(i, j int) bool {
				return a_datatypesOther.DATATYPE_DEFINITION_REAL[i] == a_datatypes.DATATYPE_DEFINITION_REAL[j]
			},
			func(j int) string {
				return a_datatypes.DATATYPE_DEFINITION_REAL[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DATATYPE_DEFINITION_STRINGDifferent := false
	if len(a_datatypes.DATATYPE_DEFINITION_STRING) != len(a_datatypesOther.DATATYPE_DEFINITION_STRING) {
		DATATYPE_DEFINITION_STRINGDifferent = true
	} else {
		for i := range a_datatypes.DATATYPE_DEFINITION_STRING {
			if (a_datatypes.DATATYPE_DEFINITION_STRING[i] == nil) != (a_datatypesOther.DATATYPE_DEFINITION_STRING[i] == nil) {
				DATATYPE_DEFINITION_STRINGDifferent = true
				break
			} else if a_datatypes.DATATYPE_DEFINITION_STRING[i] != nil && a_datatypesOther.DATATYPE_DEFINITION_STRING[i] != nil {
				// this is a pointer comparaison
				if a_datatypes.DATATYPE_DEFINITION_STRING[i] != a_datatypesOther.DATATYPE_DEFINITION_STRING[i] {
					DATATYPE_DEFINITION_STRINGDifferent = true
					break
				}
			}
		}
	}
	if DATATYPE_DEFINITION_STRINGDifferent {
		ops := stage.Diff(
			a_datatypes,
			"DATATYPE_DEFINITION_STRING",
			len(a_datatypesOther.DATATYPE_DEFINITION_STRING),
			len(a_datatypes.DATATYPE_DEFINITION_STRING),
			func(i, j int) bool {
				return a_datatypesOther.DATATYPE_DEFINITION_STRING[i] == a_datatypes.DATATYPE_DEFINITION_STRING[j]
			},
			func(j int) string {
				return a_datatypes.DATATYPE_DEFINITION_STRING[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	DATATYPE_DEFINITION_XHTMLDifferent := false
	if len(a_datatypes.DATATYPE_DEFINITION_XHTML) != len(a_datatypesOther.DATATYPE_DEFINITION_XHTML) {
		DATATYPE_DEFINITION_XHTMLDifferent = true
	} else {
		for i := range a_datatypes.DATATYPE_DEFINITION_XHTML {
			if (a_datatypes.DATATYPE_DEFINITION_XHTML[i] == nil) != (a_datatypesOther.DATATYPE_DEFINITION_XHTML[i] == nil) {
				DATATYPE_DEFINITION_XHTMLDifferent = true
				break
			} else if a_datatypes.DATATYPE_DEFINITION_XHTML[i] != nil && a_datatypesOther.DATATYPE_DEFINITION_XHTML[i] != nil {
				// this is a pointer comparaison
				if a_datatypes.DATATYPE_DEFINITION_XHTML[i] != a_datatypesOther.DATATYPE_DEFINITION_XHTML[i] {
					DATATYPE_DEFINITION_XHTMLDifferent = true
					break
				}
			}
		}
	}
	if DATATYPE_DEFINITION_XHTMLDifferent {
		ops := stage.Diff(
			a_datatypes,
			"DATATYPE_DEFINITION_XHTML",
			len(a_datatypesOther.DATATYPE_DEFINITION_XHTML),
			len(a_datatypes.DATATYPE_DEFINITION_XHTML),
			func(i, j int) bool {
				return a_datatypesOther.DATATYPE_DEFINITION_XHTML[i] == a_datatypes.DATATYPE_DEFINITION_XHTML[j]
			},
			func(j int) string {
				return a_datatypes.DATATYPE_DEFINITION_XHTML[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_datatype_definition_boolean_ref *A_DATATYPE_DEFINITION_BOOLEAN_REF) GongDiff(stage *Stage, a_datatype_definition_boolean_refOther *A_DATATYPE_DEFINITION_BOOLEAN_REF) (diffs []string) {
	// insertion point for field diffs
	if a_datatype_definition_boolean_ref.Name != a_datatype_definition_boolean_refOther.Name {
		diffs = append(diffs, a_datatype_definition_boolean_ref.GongMarshallField(stage, "Name"))
	}
	if a_datatype_definition_boolean_ref.DATATYPE_DEFINITION_BOOLEAN_REF != a_datatype_definition_boolean_refOther.DATATYPE_DEFINITION_BOOLEAN_REF {
		diffs = append(diffs, a_datatype_definition_boolean_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_BOOLEAN_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_datatype_definition_date_ref *A_DATATYPE_DEFINITION_DATE_REF) GongDiff(stage *Stage, a_datatype_definition_date_refOther *A_DATATYPE_DEFINITION_DATE_REF) (diffs []string) {
	// insertion point for field diffs
	if a_datatype_definition_date_ref.Name != a_datatype_definition_date_refOther.Name {
		diffs = append(diffs, a_datatype_definition_date_ref.GongMarshallField(stage, "Name"))
	}
	if a_datatype_definition_date_ref.DATATYPE_DEFINITION_DATE_REF != a_datatype_definition_date_refOther.DATATYPE_DEFINITION_DATE_REF {
		diffs = append(diffs, a_datatype_definition_date_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_DATE_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_datatype_definition_enumeration_ref *A_DATATYPE_DEFINITION_ENUMERATION_REF) GongDiff(stage *Stage, a_datatype_definition_enumeration_refOther *A_DATATYPE_DEFINITION_ENUMERATION_REF) (diffs []string) {
	// insertion point for field diffs
	if a_datatype_definition_enumeration_ref.Name != a_datatype_definition_enumeration_refOther.Name {
		diffs = append(diffs, a_datatype_definition_enumeration_ref.GongMarshallField(stage, "Name"))
	}
	if a_datatype_definition_enumeration_ref.DATATYPE_DEFINITION_ENUMERATION_REF != a_datatype_definition_enumeration_refOther.DATATYPE_DEFINITION_ENUMERATION_REF {
		diffs = append(diffs, a_datatype_definition_enumeration_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_ENUMERATION_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_datatype_definition_integer_ref *A_DATATYPE_DEFINITION_INTEGER_REF) GongDiff(stage *Stage, a_datatype_definition_integer_refOther *A_DATATYPE_DEFINITION_INTEGER_REF) (diffs []string) {
	// insertion point for field diffs
	if a_datatype_definition_integer_ref.Name != a_datatype_definition_integer_refOther.Name {
		diffs = append(diffs, a_datatype_definition_integer_ref.GongMarshallField(stage, "Name"))
	}
	if a_datatype_definition_integer_ref.DATATYPE_DEFINITION_INTEGER_REF != a_datatype_definition_integer_refOther.DATATYPE_DEFINITION_INTEGER_REF {
		diffs = append(diffs, a_datatype_definition_integer_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_INTEGER_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_datatype_definition_real_ref *A_DATATYPE_DEFINITION_REAL_REF) GongDiff(stage *Stage, a_datatype_definition_real_refOther *A_DATATYPE_DEFINITION_REAL_REF) (diffs []string) {
	// insertion point for field diffs
	if a_datatype_definition_real_ref.Name != a_datatype_definition_real_refOther.Name {
		diffs = append(diffs, a_datatype_definition_real_ref.GongMarshallField(stage, "Name"))
	}
	if a_datatype_definition_real_ref.DATATYPE_DEFINITION_REAL_REF != a_datatype_definition_real_refOther.DATATYPE_DEFINITION_REAL_REF {
		diffs = append(diffs, a_datatype_definition_real_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_REAL_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_datatype_definition_string_ref *A_DATATYPE_DEFINITION_STRING_REF) GongDiff(stage *Stage, a_datatype_definition_string_refOther *A_DATATYPE_DEFINITION_STRING_REF) (diffs []string) {
	// insertion point for field diffs
	if a_datatype_definition_string_ref.Name != a_datatype_definition_string_refOther.Name {
		diffs = append(diffs, a_datatype_definition_string_ref.GongMarshallField(stage, "Name"))
	}
	if a_datatype_definition_string_ref.DATATYPE_DEFINITION_STRING_REF != a_datatype_definition_string_refOther.DATATYPE_DEFINITION_STRING_REF {
		diffs = append(diffs, a_datatype_definition_string_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_STRING_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_datatype_definition_xhtml_ref *A_DATATYPE_DEFINITION_XHTML_REF) GongDiff(stage *Stage, a_datatype_definition_xhtml_refOther *A_DATATYPE_DEFINITION_XHTML_REF) (diffs []string) {
	// insertion point for field diffs
	if a_datatype_definition_xhtml_ref.Name != a_datatype_definition_xhtml_refOther.Name {
		diffs = append(diffs, a_datatype_definition_xhtml_ref.GongMarshallField(stage, "Name"))
	}
	if a_datatype_definition_xhtml_ref.DATATYPE_DEFINITION_XHTML_REF != a_datatype_definition_xhtml_refOther.DATATYPE_DEFINITION_XHTML_REF {
		diffs = append(diffs, a_datatype_definition_xhtml_ref.GongMarshallField(stage, "DATATYPE_DEFINITION_XHTML_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_editable_atts *A_EDITABLE_ATTS) GongDiff(stage *Stage, a_editable_attsOther *A_EDITABLE_ATTS) (diffs []string) {
	// insertion point for field diffs
	if a_editable_atts.Name != a_editable_attsOther.Name {
		diffs = append(diffs, a_editable_atts.GongMarshallField(stage, "Name"))
	}
	if a_editable_atts.ATTRIBUTE_DEFINITION_BOOLEAN_REF != a_editable_attsOther.ATTRIBUTE_DEFINITION_BOOLEAN_REF {
		diffs = append(diffs, a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_BOOLEAN_REF"))
	}
	if a_editable_atts.ATTRIBUTE_DEFINITION_DATE_REF != a_editable_attsOther.ATTRIBUTE_DEFINITION_DATE_REF {
		diffs = append(diffs, a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_DATE_REF"))
	}
	if a_editable_atts.ATTRIBUTE_DEFINITION_ENUMERATION_REF != a_editable_attsOther.ATTRIBUTE_DEFINITION_ENUMERATION_REF {
		diffs = append(diffs, a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_ENUMERATION_REF"))
	}
	if a_editable_atts.ATTRIBUTE_DEFINITION_INTEGER_REF != a_editable_attsOther.ATTRIBUTE_DEFINITION_INTEGER_REF {
		diffs = append(diffs, a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_INTEGER_REF"))
	}
	if a_editable_atts.ATTRIBUTE_DEFINITION_REAL_REF != a_editable_attsOther.ATTRIBUTE_DEFINITION_REAL_REF {
		diffs = append(diffs, a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_REAL_REF"))
	}
	if a_editable_atts.ATTRIBUTE_DEFINITION_STRING_REF != a_editable_attsOther.ATTRIBUTE_DEFINITION_STRING_REF {
		diffs = append(diffs, a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_STRING_REF"))
	}
	if a_editable_atts.ATTRIBUTE_DEFINITION_XHTML_REF != a_editable_attsOther.ATTRIBUTE_DEFINITION_XHTML_REF {
		diffs = append(diffs, a_editable_atts.GongMarshallField(stage, "ATTRIBUTE_DEFINITION_XHTML_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_enum_value_ref *A_ENUM_VALUE_REF) GongDiff(stage *Stage, a_enum_value_refOther *A_ENUM_VALUE_REF) (diffs []string) {
	// insertion point for field diffs
	if a_enum_value_ref.Name != a_enum_value_refOther.Name {
		diffs = append(diffs, a_enum_value_ref.GongMarshallField(stage, "Name"))
	}
	if a_enum_value_ref.ENUM_VALUE_REF != a_enum_value_refOther.ENUM_VALUE_REF {
		diffs = append(diffs, a_enum_value_ref.GongMarshallField(stage, "ENUM_VALUE_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_object *A_OBJECT) GongDiff(stage *Stage, a_objectOther *A_OBJECT) (diffs []string) {
	// insertion point for field diffs
	if a_object.Name != a_objectOther.Name {
		diffs = append(diffs, a_object.GongMarshallField(stage, "Name"))
	}
	if a_object.SPEC_OBJECT_REF != a_objectOther.SPEC_OBJECT_REF {
		diffs = append(diffs, a_object.GongMarshallField(stage, "SPEC_OBJECT_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_properties *A_PROPERTIES) GongDiff(stage *Stage, a_propertiesOther *A_PROPERTIES) (diffs []string) {
	// insertion point for field diffs
	if a_properties.Name != a_propertiesOther.Name {
		diffs = append(diffs, a_properties.GongMarshallField(stage, "Name"))
	}
	if (a_properties.EMBEDDED_VALUE == nil) != (a_propertiesOther.EMBEDDED_VALUE == nil) {
		diffs = append(diffs, a_properties.GongMarshallField(stage, "EMBEDDED_VALUE"))
	} else if a_properties.EMBEDDED_VALUE != nil && a_propertiesOther.EMBEDDED_VALUE != nil {
		if a_properties.EMBEDDED_VALUE != a_propertiesOther.EMBEDDED_VALUE {
			diffs = append(diffs, a_properties.GongMarshallField(stage, "EMBEDDED_VALUE"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_relation_group_type_ref *A_RELATION_GROUP_TYPE_REF) GongDiff(stage *Stage, a_relation_group_type_refOther *A_RELATION_GROUP_TYPE_REF) (diffs []string) {
	// insertion point for field diffs
	if a_relation_group_type_ref.Name != a_relation_group_type_refOther.Name {
		diffs = append(diffs, a_relation_group_type_ref.GongMarshallField(stage, "Name"))
	}
	if a_relation_group_type_ref.RELATION_GROUP_TYPE_REF != a_relation_group_type_refOther.RELATION_GROUP_TYPE_REF {
		diffs = append(diffs, a_relation_group_type_ref.GongMarshallField(stage, "RELATION_GROUP_TYPE_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_source_1 *A_SOURCE_1) GongDiff(stage *Stage, a_source_1Other *A_SOURCE_1) (diffs []string) {
	// insertion point for field diffs
	if a_source_1.Name != a_source_1Other.Name {
		diffs = append(diffs, a_source_1.GongMarshallField(stage, "Name"))
	}
	if a_source_1.SPEC_OBJECT_REF != a_source_1Other.SPEC_OBJECT_REF {
		diffs = append(diffs, a_source_1.GongMarshallField(stage, "SPEC_OBJECT_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_source_specification_1 *A_SOURCE_SPECIFICATION_1) GongDiff(stage *Stage, a_source_specification_1Other *A_SOURCE_SPECIFICATION_1) (diffs []string) {
	// insertion point for field diffs
	if a_source_specification_1.Name != a_source_specification_1Other.Name {
		diffs = append(diffs, a_source_specification_1.GongMarshallField(stage, "Name"))
	}
	if a_source_specification_1.SPECIFICATION_REF != a_source_specification_1Other.SPECIFICATION_REF {
		diffs = append(diffs, a_source_specification_1.GongMarshallField(stage, "SPECIFICATION_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_specifications *A_SPECIFICATIONS) GongDiff(stage *Stage, a_specificationsOther *A_SPECIFICATIONS) (diffs []string) {
	// insertion point for field diffs
	if a_specifications.Name != a_specificationsOther.Name {
		diffs = append(diffs, a_specifications.GongMarshallField(stage, "Name"))
	}
	SPECIFICATIONDifferent := false
	if len(a_specifications.SPECIFICATION) != len(a_specificationsOther.SPECIFICATION) {
		SPECIFICATIONDifferent = true
	} else {
		for i := range a_specifications.SPECIFICATION {
			if (a_specifications.SPECIFICATION[i] == nil) != (a_specificationsOther.SPECIFICATION[i] == nil) {
				SPECIFICATIONDifferent = true
				break
			} else if a_specifications.SPECIFICATION[i] != nil && a_specificationsOther.SPECIFICATION[i] != nil {
				// this is a pointer comparaison
				if a_specifications.SPECIFICATION[i] != a_specificationsOther.SPECIFICATION[i] {
					SPECIFICATIONDifferent = true
					break
				}
			}
		}
	}
	if SPECIFICATIONDifferent {
		ops := stage.Diff(
			a_specifications,
			"SPECIFICATION",
			len(a_specificationsOther.SPECIFICATION),
			len(a_specifications.SPECIFICATION),
			func(i, j int) bool {
				return a_specificationsOther.SPECIFICATION[i] == a_specifications.SPECIFICATION[j]
			},
			func(j int) string {
				return a_specifications.SPECIFICATION[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_specification_type_ref *A_SPECIFICATION_TYPE_REF) GongDiff(stage *Stage, a_specification_type_refOther *A_SPECIFICATION_TYPE_REF) (diffs []string) {
	// insertion point for field diffs
	if a_specification_type_ref.Name != a_specification_type_refOther.Name {
		diffs = append(diffs, a_specification_type_ref.GongMarshallField(stage, "Name"))
	}
	if a_specification_type_ref.SPECIFICATION_TYPE_REF != a_specification_type_refOther.SPECIFICATION_TYPE_REF {
		diffs = append(diffs, a_specification_type_ref.GongMarshallField(stage, "SPECIFICATION_TYPE_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_specified_values *A_SPECIFIED_VALUES) GongDiff(stage *Stage, a_specified_valuesOther *A_SPECIFIED_VALUES) (diffs []string) {
	// insertion point for field diffs
	if a_specified_values.Name != a_specified_valuesOther.Name {
		diffs = append(diffs, a_specified_values.GongMarshallField(stage, "Name"))
	}
	ENUM_VALUEDifferent := false
	if len(a_specified_values.ENUM_VALUE) != len(a_specified_valuesOther.ENUM_VALUE) {
		ENUM_VALUEDifferent = true
	} else {
		for i := range a_specified_values.ENUM_VALUE {
			if (a_specified_values.ENUM_VALUE[i] == nil) != (a_specified_valuesOther.ENUM_VALUE[i] == nil) {
				ENUM_VALUEDifferent = true
				break
			} else if a_specified_values.ENUM_VALUE[i] != nil && a_specified_valuesOther.ENUM_VALUE[i] != nil {
				// this is a pointer comparaison
				if a_specified_values.ENUM_VALUE[i] != a_specified_valuesOther.ENUM_VALUE[i] {
					ENUM_VALUEDifferent = true
					break
				}
			}
		}
	}
	if ENUM_VALUEDifferent {
		ops := stage.Diff(
			a_specified_values,
			"ENUM_VALUE",
			len(a_specified_valuesOther.ENUM_VALUE),
			len(a_specified_values.ENUM_VALUE),
			func(i, j int) bool {
				return a_specified_valuesOther.ENUM_VALUE[i] == a_specified_values.ENUM_VALUE[j]
			},
			func(j int) string {
				return a_specified_values.ENUM_VALUE[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_spec_attributes *A_SPEC_ATTRIBUTES) GongDiff(stage *Stage, a_spec_attributesOther *A_SPEC_ATTRIBUTES) (diffs []string) {
	// insertion point for field diffs
	if a_spec_attributes.Name != a_spec_attributesOther.Name {
		diffs = append(diffs, a_spec_attributes.GongMarshallField(stage, "Name"))
	}
	ATTRIBUTE_DEFINITION_BOOLEANDifferent := false
	if len(a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN) != len(a_spec_attributesOther.ATTRIBUTE_DEFINITION_BOOLEAN) {
		ATTRIBUTE_DEFINITION_BOOLEANDifferent = true
	} else {
		for i := range a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN {
			if (a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN[i] == nil) != (a_spec_attributesOther.ATTRIBUTE_DEFINITION_BOOLEAN[i] == nil) {
				ATTRIBUTE_DEFINITION_BOOLEANDifferent = true
				break
			} else if a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN[i] != nil && a_spec_attributesOther.ATTRIBUTE_DEFINITION_BOOLEAN[i] != nil {
				// this is a pointer comparaison
				if a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN[i] != a_spec_attributesOther.ATTRIBUTE_DEFINITION_BOOLEAN[i] {
					ATTRIBUTE_DEFINITION_BOOLEANDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_DEFINITION_BOOLEANDifferent {
		ops := stage.Diff(
			a_spec_attributes,
			"ATTRIBUTE_DEFINITION_BOOLEAN",
			len(a_spec_attributesOther.ATTRIBUTE_DEFINITION_BOOLEAN),
			len(a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN),
			func(i, j int) bool {
				return a_spec_attributesOther.ATTRIBUTE_DEFINITION_BOOLEAN[i] == a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN[j]
			},
			func(j int) string {
				return a_spec_attributes.ATTRIBUTE_DEFINITION_BOOLEAN[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ATTRIBUTE_DEFINITION_DATEDifferent := false
	if len(a_spec_attributes.ATTRIBUTE_DEFINITION_DATE) != len(a_spec_attributesOther.ATTRIBUTE_DEFINITION_DATE) {
		ATTRIBUTE_DEFINITION_DATEDifferent = true
	} else {
		for i := range a_spec_attributes.ATTRIBUTE_DEFINITION_DATE {
			if (a_spec_attributes.ATTRIBUTE_DEFINITION_DATE[i] == nil) != (a_spec_attributesOther.ATTRIBUTE_DEFINITION_DATE[i] == nil) {
				ATTRIBUTE_DEFINITION_DATEDifferent = true
				break
			} else if a_spec_attributes.ATTRIBUTE_DEFINITION_DATE[i] != nil && a_spec_attributesOther.ATTRIBUTE_DEFINITION_DATE[i] != nil {
				// this is a pointer comparaison
				if a_spec_attributes.ATTRIBUTE_DEFINITION_DATE[i] != a_spec_attributesOther.ATTRIBUTE_DEFINITION_DATE[i] {
					ATTRIBUTE_DEFINITION_DATEDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_DEFINITION_DATEDifferent {
		ops := stage.Diff(
			a_spec_attributes,
			"ATTRIBUTE_DEFINITION_DATE",
			len(a_spec_attributesOther.ATTRIBUTE_DEFINITION_DATE),
			len(a_spec_attributes.ATTRIBUTE_DEFINITION_DATE),
			func(i, j int) bool {
				return a_spec_attributesOther.ATTRIBUTE_DEFINITION_DATE[i] == a_spec_attributes.ATTRIBUTE_DEFINITION_DATE[j]
			},
			func(j int) string {
				return a_spec_attributes.ATTRIBUTE_DEFINITION_DATE[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ATTRIBUTE_DEFINITION_ENUMERATIONDifferent := false
	if len(a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION) != len(a_spec_attributesOther.ATTRIBUTE_DEFINITION_ENUMERATION) {
		ATTRIBUTE_DEFINITION_ENUMERATIONDifferent = true
	} else {
		for i := range a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION {
			if (a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION[i] == nil) != (a_spec_attributesOther.ATTRIBUTE_DEFINITION_ENUMERATION[i] == nil) {
				ATTRIBUTE_DEFINITION_ENUMERATIONDifferent = true
				break
			} else if a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION[i] != nil && a_spec_attributesOther.ATTRIBUTE_DEFINITION_ENUMERATION[i] != nil {
				// this is a pointer comparaison
				if a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION[i] != a_spec_attributesOther.ATTRIBUTE_DEFINITION_ENUMERATION[i] {
					ATTRIBUTE_DEFINITION_ENUMERATIONDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_DEFINITION_ENUMERATIONDifferent {
		ops := stage.Diff(
			a_spec_attributes,
			"ATTRIBUTE_DEFINITION_ENUMERATION",
			len(a_spec_attributesOther.ATTRIBUTE_DEFINITION_ENUMERATION),
			len(a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION),
			func(i, j int) bool {
				return a_spec_attributesOther.ATTRIBUTE_DEFINITION_ENUMERATION[i] == a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION[j]
			},
			func(j int) string {
				return a_spec_attributes.ATTRIBUTE_DEFINITION_ENUMERATION[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ATTRIBUTE_DEFINITION_INTEGERDifferent := false
	if len(a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER) != len(a_spec_attributesOther.ATTRIBUTE_DEFINITION_INTEGER) {
		ATTRIBUTE_DEFINITION_INTEGERDifferent = true
	} else {
		for i := range a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER {
			if (a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER[i] == nil) != (a_spec_attributesOther.ATTRIBUTE_DEFINITION_INTEGER[i] == nil) {
				ATTRIBUTE_DEFINITION_INTEGERDifferent = true
				break
			} else if a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER[i] != nil && a_spec_attributesOther.ATTRIBUTE_DEFINITION_INTEGER[i] != nil {
				// this is a pointer comparaison
				if a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER[i] != a_spec_attributesOther.ATTRIBUTE_DEFINITION_INTEGER[i] {
					ATTRIBUTE_DEFINITION_INTEGERDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_DEFINITION_INTEGERDifferent {
		ops := stage.Diff(
			a_spec_attributes,
			"ATTRIBUTE_DEFINITION_INTEGER",
			len(a_spec_attributesOther.ATTRIBUTE_DEFINITION_INTEGER),
			len(a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER),
			func(i, j int) bool {
				return a_spec_attributesOther.ATTRIBUTE_DEFINITION_INTEGER[i] == a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER[j]
			},
			func(j int) string {
				return a_spec_attributes.ATTRIBUTE_DEFINITION_INTEGER[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ATTRIBUTE_DEFINITION_REALDifferent := false
	if len(a_spec_attributes.ATTRIBUTE_DEFINITION_REAL) != len(a_spec_attributesOther.ATTRIBUTE_DEFINITION_REAL) {
		ATTRIBUTE_DEFINITION_REALDifferent = true
	} else {
		for i := range a_spec_attributes.ATTRIBUTE_DEFINITION_REAL {
			if (a_spec_attributes.ATTRIBUTE_DEFINITION_REAL[i] == nil) != (a_spec_attributesOther.ATTRIBUTE_DEFINITION_REAL[i] == nil) {
				ATTRIBUTE_DEFINITION_REALDifferent = true
				break
			} else if a_spec_attributes.ATTRIBUTE_DEFINITION_REAL[i] != nil && a_spec_attributesOther.ATTRIBUTE_DEFINITION_REAL[i] != nil {
				// this is a pointer comparaison
				if a_spec_attributes.ATTRIBUTE_DEFINITION_REAL[i] != a_spec_attributesOther.ATTRIBUTE_DEFINITION_REAL[i] {
					ATTRIBUTE_DEFINITION_REALDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_DEFINITION_REALDifferent {
		ops := stage.Diff(
			a_spec_attributes,
			"ATTRIBUTE_DEFINITION_REAL",
			len(a_spec_attributesOther.ATTRIBUTE_DEFINITION_REAL),
			len(a_spec_attributes.ATTRIBUTE_DEFINITION_REAL),
			func(i, j int) bool {
				return a_spec_attributesOther.ATTRIBUTE_DEFINITION_REAL[i] == a_spec_attributes.ATTRIBUTE_DEFINITION_REAL[j]
			},
			func(j int) string {
				return a_spec_attributes.ATTRIBUTE_DEFINITION_REAL[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ATTRIBUTE_DEFINITION_STRINGDifferent := false
	if len(a_spec_attributes.ATTRIBUTE_DEFINITION_STRING) != len(a_spec_attributesOther.ATTRIBUTE_DEFINITION_STRING) {
		ATTRIBUTE_DEFINITION_STRINGDifferent = true
	} else {
		for i := range a_spec_attributes.ATTRIBUTE_DEFINITION_STRING {
			if (a_spec_attributes.ATTRIBUTE_DEFINITION_STRING[i] == nil) != (a_spec_attributesOther.ATTRIBUTE_DEFINITION_STRING[i] == nil) {
				ATTRIBUTE_DEFINITION_STRINGDifferent = true
				break
			} else if a_spec_attributes.ATTRIBUTE_DEFINITION_STRING[i] != nil && a_spec_attributesOther.ATTRIBUTE_DEFINITION_STRING[i] != nil {
				// this is a pointer comparaison
				if a_spec_attributes.ATTRIBUTE_DEFINITION_STRING[i] != a_spec_attributesOther.ATTRIBUTE_DEFINITION_STRING[i] {
					ATTRIBUTE_DEFINITION_STRINGDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_DEFINITION_STRINGDifferent {
		ops := stage.Diff(
			a_spec_attributes,
			"ATTRIBUTE_DEFINITION_STRING",
			len(a_spec_attributesOther.ATTRIBUTE_DEFINITION_STRING),
			len(a_spec_attributes.ATTRIBUTE_DEFINITION_STRING),
			func(i, j int) bool {
				return a_spec_attributesOther.ATTRIBUTE_DEFINITION_STRING[i] == a_spec_attributes.ATTRIBUTE_DEFINITION_STRING[j]
			},
			func(j int) string {
				return a_spec_attributes.ATTRIBUTE_DEFINITION_STRING[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	ATTRIBUTE_DEFINITION_XHTMLDifferent := false
	if len(a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML) != len(a_spec_attributesOther.ATTRIBUTE_DEFINITION_XHTML) {
		ATTRIBUTE_DEFINITION_XHTMLDifferent = true
	} else {
		for i := range a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML {
			if (a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML[i] == nil) != (a_spec_attributesOther.ATTRIBUTE_DEFINITION_XHTML[i] == nil) {
				ATTRIBUTE_DEFINITION_XHTMLDifferent = true
				break
			} else if a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML[i] != nil && a_spec_attributesOther.ATTRIBUTE_DEFINITION_XHTML[i] != nil {
				// this is a pointer comparaison
				if a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML[i] != a_spec_attributesOther.ATTRIBUTE_DEFINITION_XHTML[i] {
					ATTRIBUTE_DEFINITION_XHTMLDifferent = true
					break
				}
			}
		}
	}
	if ATTRIBUTE_DEFINITION_XHTMLDifferent {
		ops := stage.Diff(
			a_spec_attributes,
			"ATTRIBUTE_DEFINITION_XHTML",
			len(a_spec_attributesOther.ATTRIBUTE_DEFINITION_XHTML),
			len(a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML),
			func(i, j int) bool {
				return a_spec_attributesOther.ATTRIBUTE_DEFINITION_XHTML[i] == a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML[j]
			},
			func(j int) string {
				return a_spec_attributes.ATTRIBUTE_DEFINITION_XHTML[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_spec_objects *A_SPEC_OBJECTS) GongDiff(stage *Stage, a_spec_objectsOther *A_SPEC_OBJECTS) (diffs []string) {
	// insertion point for field diffs
	if a_spec_objects.Name != a_spec_objectsOther.Name {
		diffs = append(diffs, a_spec_objects.GongMarshallField(stage, "Name"))
	}
	SPEC_OBJECTDifferent := false
	if len(a_spec_objects.SPEC_OBJECT) != len(a_spec_objectsOther.SPEC_OBJECT) {
		SPEC_OBJECTDifferent = true
	} else {
		for i := range a_spec_objects.SPEC_OBJECT {
			if (a_spec_objects.SPEC_OBJECT[i] == nil) != (a_spec_objectsOther.SPEC_OBJECT[i] == nil) {
				SPEC_OBJECTDifferent = true
				break
			} else if a_spec_objects.SPEC_OBJECT[i] != nil && a_spec_objectsOther.SPEC_OBJECT[i] != nil {
				// this is a pointer comparaison
				if a_spec_objects.SPEC_OBJECT[i] != a_spec_objectsOther.SPEC_OBJECT[i] {
					SPEC_OBJECTDifferent = true
					break
				}
			}
		}
	}
	if SPEC_OBJECTDifferent {
		ops := stage.Diff(
			a_spec_objects,
			"SPEC_OBJECT",
			len(a_spec_objectsOther.SPEC_OBJECT),
			len(a_spec_objects.SPEC_OBJECT),
			func(i, j int) bool {
				return a_spec_objectsOther.SPEC_OBJECT[i] == a_spec_objects.SPEC_OBJECT[j]
			},
			func(j int) string {
				return a_spec_objects.SPEC_OBJECT[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_spec_object_type_ref *A_SPEC_OBJECT_TYPE_REF) GongDiff(stage *Stage, a_spec_object_type_refOther *A_SPEC_OBJECT_TYPE_REF) (diffs []string) {
	// insertion point for field diffs
	if a_spec_object_type_ref.Name != a_spec_object_type_refOther.Name {
		diffs = append(diffs, a_spec_object_type_ref.GongMarshallField(stage, "Name"))
	}
	if a_spec_object_type_ref.SPEC_OBJECT_TYPE_REF != a_spec_object_type_refOther.SPEC_OBJECT_TYPE_REF {
		diffs = append(diffs, a_spec_object_type_ref.GongMarshallField(stage, "SPEC_OBJECT_TYPE_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_spec_relations *A_SPEC_RELATIONS) GongDiff(stage *Stage, a_spec_relationsOther *A_SPEC_RELATIONS) (diffs []string) {
	// insertion point for field diffs
	if a_spec_relations.Name != a_spec_relationsOther.Name {
		diffs = append(diffs, a_spec_relations.GongMarshallField(stage, "Name"))
	}
	SPEC_RELATIONDifferent := false
	if len(a_spec_relations.SPEC_RELATION) != len(a_spec_relationsOther.SPEC_RELATION) {
		SPEC_RELATIONDifferent = true
	} else {
		for i := range a_spec_relations.SPEC_RELATION {
			if (a_spec_relations.SPEC_RELATION[i] == nil) != (a_spec_relationsOther.SPEC_RELATION[i] == nil) {
				SPEC_RELATIONDifferent = true
				break
			} else if a_spec_relations.SPEC_RELATION[i] != nil && a_spec_relationsOther.SPEC_RELATION[i] != nil {
				// this is a pointer comparaison
				if a_spec_relations.SPEC_RELATION[i] != a_spec_relationsOther.SPEC_RELATION[i] {
					SPEC_RELATIONDifferent = true
					break
				}
			}
		}
	}
	if SPEC_RELATIONDifferent {
		ops := stage.Diff(
			a_spec_relations,
			"SPEC_RELATION",
			len(a_spec_relationsOther.SPEC_RELATION),
			len(a_spec_relations.SPEC_RELATION),
			func(i, j int) bool {
				return a_spec_relationsOther.SPEC_RELATION[i] == a_spec_relations.SPEC_RELATION[j]
			},
			func(j int) string {
				return a_spec_relations.SPEC_RELATION[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_spec_relation_groups *A_SPEC_RELATION_GROUPS) GongDiff(stage *Stage, a_spec_relation_groupsOther *A_SPEC_RELATION_GROUPS) (diffs []string) {
	// insertion point for field diffs
	if a_spec_relation_groups.Name != a_spec_relation_groupsOther.Name {
		diffs = append(diffs, a_spec_relation_groups.GongMarshallField(stage, "Name"))
	}
	RELATION_GROUPDifferent := false
	if len(a_spec_relation_groups.RELATION_GROUP) != len(a_spec_relation_groupsOther.RELATION_GROUP) {
		RELATION_GROUPDifferent = true
	} else {
		for i := range a_spec_relation_groups.RELATION_GROUP {
			if (a_spec_relation_groups.RELATION_GROUP[i] == nil) != (a_spec_relation_groupsOther.RELATION_GROUP[i] == nil) {
				RELATION_GROUPDifferent = true
				break
			} else if a_spec_relation_groups.RELATION_GROUP[i] != nil && a_spec_relation_groupsOther.RELATION_GROUP[i] != nil {
				// this is a pointer comparaison
				if a_spec_relation_groups.RELATION_GROUP[i] != a_spec_relation_groupsOther.RELATION_GROUP[i] {
					RELATION_GROUPDifferent = true
					break
				}
			}
		}
	}
	if RELATION_GROUPDifferent {
		ops := stage.Diff(
			a_spec_relation_groups,
			"RELATION_GROUP",
			len(a_spec_relation_groupsOther.RELATION_GROUP),
			len(a_spec_relation_groups.RELATION_GROUP),
			func(i, j int) bool {
				return a_spec_relation_groupsOther.RELATION_GROUP[i] == a_spec_relation_groups.RELATION_GROUP[j]
			},
			func(j int) string {
				return a_spec_relation_groups.RELATION_GROUP[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_spec_relation_ref *A_SPEC_RELATION_REF) GongDiff(stage *Stage, a_spec_relation_refOther *A_SPEC_RELATION_REF) (diffs []string) {
	// insertion point for field diffs
	if a_spec_relation_ref.Name != a_spec_relation_refOther.Name {
		diffs = append(diffs, a_spec_relation_ref.GongMarshallField(stage, "Name"))
	}
	if a_spec_relation_ref.SPEC_RELATION_REF != a_spec_relation_refOther.SPEC_RELATION_REF {
		diffs = append(diffs, a_spec_relation_ref.GongMarshallField(stage, "SPEC_RELATION_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_spec_relation_type_ref *A_SPEC_RELATION_TYPE_REF) GongDiff(stage *Stage, a_spec_relation_type_refOther *A_SPEC_RELATION_TYPE_REF) (diffs []string) {
	// insertion point for field diffs
	if a_spec_relation_type_ref.Name != a_spec_relation_type_refOther.Name {
		diffs = append(diffs, a_spec_relation_type_ref.GongMarshallField(stage, "Name"))
	}
	if a_spec_relation_type_ref.SPEC_RELATION_TYPE_REF != a_spec_relation_type_refOther.SPEC_RELATION_TYPE_REF {
		diffs = append(diffs, a_spec_relation_type_ref.GongMarshallField(stage, "SPEC_RELATION_TYPE_REF"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_spec_types *A_SPEC_TYPES) GongDiff(stage *Stage, a_spec_typesOther *A_SPEC_TYPES) (diffs []string) {
	// insertion point for field diffs
	if a_spec_types.Name != a_spec_typesOther.Name {
		diffs = append(diffs, a_spec_types.GongMarshallField(stage, "Name"))
	}
	RELATION_GROUP_TYPEDifferent := false
	if len(a_spec_types.RELATION_GROUP_TYPE) != len(a_spec_typesOther.RELATION_GROUP_TYPE) {
		RELATION_GROUP_TYPEDifferent = true
	} else {
		for i := range a_spec_types.RELATION_GROUP_TYPE {
			if (a_spec_types.RELATION_GROUP_TYPE[i] == nil) != (a_spec_typesOther.RELATION_GROUP_TYPE[i] == nil) {
				RELATION_GROUP_TYPEDifferent = true
				break
			} else if a_spec_types.RELATION_GROUP_TYPE[i] != nil && a_spec_typesOther.RELATION_GROUP_TYPE[i] != nil {
				// this is a pointer comparaison
				if a_spec_types.RELATION_GROUP_TYPE[i] != a_spec_typesOther.RELATION_GROUP_TYPE[i] {
					RELATION_GROUP_TYPEDifferent = true
					break
				}
			}
		}
	}
	if RELATION_GROUP_TYPEDifferent {
		ops := stage.Diff(
			a_spec_types,
			"RELATION_GROUP_TYPE",
			len(a_spec_typesOther.RELATION_GROUP_TYPE),
			len(a_spec_types.RELATION_GROUP_TYPE),
			func(i, j int) bool {
				return a_spec_typesOther.RELATION_GROUP_TYPE[i] == a_spec_types.RELATION_GROUP_TYPE[j]
			},
			func(j int) string {
				return a_spec_types.RELATION_GROUP_TYPE[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	SPEC_OBJECT_TYPEDifferent := false
	if len(a_spec_types.SPEC_OBJECT_TYPE) != len(a_spec_typesOther.SPEC_OBJECT_TYPE) {
		SPEC_OBJECT_TYPEDifferent = true
	} else {
		for i := range a_spec_types.SPEC_OBJECT_TYPE {
			if (a_spec_types.SPEC_OBJECT_TYPE[i] == nil) != (a_spec_typesOther.SPEC_OBJECT_TYPE[i] == nil) {
				SPEC_OBJECT_TYPEDifferent = true
				break
			} else if a_spec_types.SPEC_OBJECT_TYPE[i] != nil && a_spec_typesOther.SPEC_OBJECT_TYPE[i] != nil {
				// this is a pointer comparaison
				if a_spec_types.SPEC_OBJECT_TYPE[i] != a_spec_typesOther.SPEC_OBJECT_TYPE[i] {
					SPEC_OBJECT_TYPEDifferent = true
					break
				}
			}
		}
	}
	if SPEC_OBJECT_TYPEDifferent {
		ops := stage.Diff(
			a_spec_types,
			"SPEC_OBJECT_TYPE",
			len(a_spec_typesOther.SPEC_OBJECT_TYPE),
			len(a_spec_types.SPEC_OBJECT_TYPE),
			func(i, j int) bool {
				return a_spec_typesOther.SPEC_OBJECT_TYPE[i] == a_spec_types.SPEC_OBJECT_TYPE[j]
			},
			func(j int) string {
				return a_spec_types.SPEC_OBJECT_TYPE[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	SPEC_RELATION_TYPEDifferent := false
	if len(a_spec_types.SPEC_RELATION_TYPE) != len(a_spec_typesOther.SPEC_RELATION_TYPE) {
		SPEC_RELATION_TYPEDifferent = true
	} else {
		for i := range a_spec_types.SPEC_RELATION_TYPE {
			if (a_spec_types.SPEC_RELATION_TYPE[i] == nil) != (a_spec_typesOther.SPEC_RELATION_TYPE[i] == nil) {
				SPEC_RELATION_TYPEDifferent = true
				break
			} else if a_spec_types.SPEC_RELATION_TYPE[i] != nil && a_spec_typesOther.SPEC_RELATION_TYPE[i] != nil {
				// this is a pointer comparaison
				if a_spec_types.SPEC_RELATION_TYPE[i] != a_spec_typesOther.SPEC_RELATION_TYPE[i] {
					SPEC_RELATION_TYPEDifferent = true
					break
				}
			}
		}
	}
	if SPEC_RELATION_TYPEDifferent {
		ops := stage.Diff(
			a_spec_types,
			"SPEC_RELATION_TYPE",
			len(a_spec_typesOther.SPEC_RELATION_TYPE),
			len(a_spec_types.SPEC_RELATION_TYPE),
			func(i, j int) bool {
				return a_spec_typesOther.SPEC_RELATION_TYPE[i] == a_spec_types.SPEC_RELATION_TYPE[j]
			},
			func(j int) string {
				return a_spec_types.SPEC_RELATION_TYPE[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}
	SPECIFICATION_TYPEDifferent := false
	if len(a_spec_types.SPECIFICATION_TYPE) != len(a_spec_typesOther.SPECIFICATION_TYPE) {
		SPECIFICATION_TYPEDifferent = true
	} else {
		for i := range a_spec_types.SPECIFICATION_TYPE {
			if (a_spec_types.SPECIFICATION_TYPE[i] == nil) != (a_spec_typesOther.SPECIFICATION_TYPE[i] == nil) {
				SPECIFICATION_TYPEDifferent = true
				break
			} else if a_spec_types.SPECIFICATION_TYPE[i] != nil && a_spec_typesOther.SPECIFICATION_TYPE[i] != nil {
				// this is a pointer comparaison
				if a_spec_types.SPECIFICATION_TYPE[i] != a_spec_typesOther.SPECIFICATION_TYPE[i] {
					SPECIFICATION_TYPEDifferent = true
					break
				}
			}
		}
	}
	if SPECIFICATION_TYPEDifferent {
		ops := stage.Diff(
			a_spec_types,
			"SPECIFICATION_TYPE",
			len(a_spec_typesOther.SPECIFICATION_TYPE),
			len(a_spec_types.SPECIFICATION_TYPE),
			func(i, j int) bool {
				return a_spec_typesOther.SPECIFICATION_TYPE[i] == a_spec_types.SPECIFICATION_TYPE[j]
			},
			func(j int) string {
				return a_spec_types.SPECIFICATION_TYPE[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_the_header *A_THE_HEADER) GongDiff(stage *Stage, a_the_headerOther *A_THE_HEADER) (diffs []string) {
	// insertion point for field diffs
	if a_the_header.Name != a_the_headerOther.Name {
		diffs = append(diffs, a_the_header.GongMarshallField(stage, "Name"))
	}
	if (a_the_header.REQ_IF_HEADER == nil) != (a_the_headerOther.REQ_IF_HEADER == nil) {
		diffs = append(diffs, a_the_header.GongMarshallField(stage, "REQ_IF_HEADER"))
	} else if a_the_header.REQ_IF_HEADER != nil && a_the_headerOther.REQ_IF_HEADER != nil {
		if a_the_header.REQ_IF_HEADER != a_the_headerOther.REQ_IF_HEADER {
			diffs = append(diffs, a_the_header.GongMarshallField(stage, "REQ_IF_HEADER"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (a_tool_extensions *A_TOOL_EXTENSIONS) GongDiff(stage *Stage, a_tool_extensionsOther *A_TOOL_EXTENSIONS) (diffs []string) {
	// insertion point for field diffs
	if a_tool_extensions.Name != a_tool_extensionsOther.Name {
		diffs = append(diffs, a_tool_extensions.GongMarshallField(stage, "Name"))
	}
	REQ_IF_TOOL_EXTENSIONDifferent := false
	if len(a_tool_extensions.REQ_IF_TOOL_EXTENSION) != len(a_tool_extensionsOther.REQ_IF_TOOL_EXTENSION) {
		REQ_IF_TOOL_EXTENSIONDifferent = true
	} else {
		for i := range a_tool_extensions.REQ_IF_TOOL_EXTENSION {
			if (a_tool_extensions.REQ_IF_TOOL_EXTENSION[i] == nil) != (a_tool_extensionsOther.REQ_IF_TOOL_EXTENSION[i] == nil) {
				REQ_IF_TOOL_EXTENSIONDifferent = true
				break
			} else if a_tool_extensions.REQ_IF_TOOL_EXTENSION[i] != nil && a_tool_extensionsOther.REQ_IF_TOOL_EXTENSION[i] != nil {
				// this is a pointer comparaison
				if a_tool_extensions.REQ_IF_TOOL_EXTENSION[i] != a_tool_extensionsOther.REQ_IF_TOOL_EXTENSION[i] {
					REQ_IF_TOOL_EXTENSIONDifferent = true
					break
				}
			}
		}
	}
	if REQ_IF_TOOL_EXTENSIONDifferent {
		ops := stage.Diff(
			a_tool_extensions,
			"REQ_IF_TOOL_EXTENSION",
			len(a_tool_extensionsOther.REQ_IF_TOOL_EXTENSION),
			len(a_tool_extensions.REQ_IF_TOOL_EXTENSION),
			func(i, j int) bool {
				return a_tool_extensionsOther.REQ_IF_TOOL_EXTENSION[i] == a_tool_extensions.REQ_IF_TOOL_EXTENSION[j]
			},
			func(j int) string {
				return a_tool_extensions.REQ_IF_TOOL_EXTENSION[j].GongGetIdentifier(stage)
			},
		)
		diffs = append(diffs, ops)
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (datatype_definition_boolean *DATATYPE_DEFINITION_BOOLEAN) GongDiff(stage *Stage, datatype_definition_booleanOther *DATATYPE_DEFINITION_BOOLEAN) (diffs []string) {
	// insertion point for field diffs
	if datatype_definition_boolean.Name != datatype_definition_booleanOther.Name {
		diffs = append(diffs, datatype_definition_boolean.GongMarshallField(stage, "Name"))
	}
	if datatype_definition_boolean.DESC != datatype_definition_booleanOther.DESC {
		diffs = append(diffs, datatype_definition_boolean.GongMarshallField(stage, "DESC"))
	}
	if datatype_definition_boolean.IDENTIFIER != datatype_definition_booleanOther.IDENTIFIER {
		diffs = append(diffs, datatype_definition_boolean.GongMarshallField(stage, "IDENTIFIER"))
	}
	if datatype_definition_boolean.LAST_CHANGE != datatype_definition_booleanOther.LAST_CHANGE {
		diffs = append(diffs, datatype_definition_boolean.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if datatype_definition_boolean.LONG_NAME != datatype_definition_booleanOther.LONG_NAME {
		diffs = append(diffs, datatype_definition_boolean.GongMarshallField(stage, "LONG_NAME"))
	}
	if (datatype_definition_boolean.ALTERNATIVE_ID == nil) != (datatype_definition_booleanOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, datatype_definition_boolean.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if datatype_definition_boolean.ALTERNATIVE_ID != nil && datatype_definition_booleanOther.ALTERNATIVE_ID != nil {
		if datatype_definition_boolean.ALTERNATIVE_ID != datatype_definition_booleanOther.ALTERNATIVE_ID {
			diffs = append(diffs, datatype_definition_boolean.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (datatype_definition_date *DATATYPE_DEFINITION_DATE) GongDiff(stage *Stage, datatype_definition_dateOther *DATATYPE_DEFINITION_DATE) (diffs []string) {
	// insertion point for field diffs
	if datatype_definition_date.Name != datatype_definition_dateOther.Name {
		diffs = append(diffs, datatype_definition_date.GongMarshallField(stage, "Name"))
	}
	if datatype_definition_date.DESC != datatype_definition_dateOther.DESC {
		diffs = append(diffs, datatype_definition_date.GongMarshallField(stage, "DESC"))
	}
	if datatype_definition_date.IDENTIFIER != datatype_definition_dateOther.IDENTIFIER {
		diffs = append(diffs, datatype_definition_date.GongMarshallField(stage, "IDENTIFIER"))
	}
	if datatype_definition_date.LAST_CHANGE != datatype_definition_dateOther.LAST_CHANGE {
		diffs = append(diffs, datatype_definition_date.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if datatype_definition_date.LONG_NAME != datatype_definition_dateOther.LONG_NAME {
		diffs = append(diffs, datatype_definition_date.GongMarshallField(stage, "LONG_NAME"))
	}
	if (datatype_definition_date.ALTERNATIVE_ID == nil) != (datatype_definition_dateOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, datatype_definition_date.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if datatype_definition_date.ALTERNATIVE_ID != nil && datatype_definition_dateOther.ALTERNATIVE_ID != nil {
		if datatype_definition_date.ALTERNATIVE_ID != datatype_definition_dateOther.ALTERNATIVE_ID {
			diffs = append(diffs, datatype_definition_date.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (datatype_definition_enumeration *DATATYPE_DEFINITION_ENUMERATION) GongDiff(stage *Stage, datatype_definition_enumerationOther *DATATYPE_DEFINITION_ENUMERATION) (diffs []string) {
	// insertion point for field diffs
	if datatype_definition_enumeration.Name != datatype_definition_enumerationOther.Name {
		diffs = append(diffs, datatype_definition_enumeration.GongMarshallField(stage, "Name"))
	}
	if datatype_definition_enumeration.DESC != datatype_definition_enumerationOther.DESC {
		diffs = append(diffs, datatype_definition_enumeration.GongMarshallField(stage, "DESC"))
	}
	if datatype_definition_enumeration.IDENTIFIER != datatype_definition_enumerationOther.IDENTIFIER {
		diffs = append(diffs, datatype_definition_enumeration.GongMarshallField(stage, "IDENTIFIER"))
	}
	if datatype_definition_enumeration.LAST_CHANGE != datatype_definition_enumerationOther.LAST_CHANGE {
		diffs = append(diffs, datatype_definition_enumeration.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if datatype_definition_enumeration.LONG_NAME != datatype_definition_enumerationOther.LONG_NAME {
		diffs = append(diffs, datatype_definition_enumeration.GongMarshallField(stage, "LONG_NAME"))
	}
	if (datatype_definition_enumeration.ALTERNATIVE_ID == nil) != (datatype_definition_enumerationOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, datatype_definition_enumeration.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if datatype_definition_enumeration.ALTERNATIVE_ID != nil && datatype_definition_enumerationOther.ALTERNATIVE_ID != nil {
		if datatype_definition_enumeration.ALTERNATIVE_ID != datatype_definition_enumerationOther.ALTERNATIVE_ID {
			diffs = append(diffs, datatype_definition_enumeration.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (datatype_definition_enumeration.SPECIFIED_VALUES == nil) != (datatype_definition_enumerationOther.SPECIFIED_VALUES == nil) {
		diffs = append(diffs, datatype_definition_enumeration.GongMarshallField(stage, "SPECIFIED_VALUES"))
	} else if datatype_definition_enumeration.SPECIFIED_VALUES != nil && datatype_definition_enumerationOther.SPECIFIED_VALUES != nil {
		if datatype_definition_enumeration.SPECIFIED_VALUES != datatype_definition_enumerationOther.SPECIFIED_VALUES {
			diffs = append(diffs, datatype_definition_enumeration.GongMarshallField(stage, "SPECIFIED_VALUES"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (datatype_definition_integer *DATATYPE_DEFINITION_INTEGER) GongDiff(stage *Stage, datatype_definition_integerOther *DATATYPE_DEFINITION_INTEGER) (diffs []string) {
	// insertion point for field diffs
	if datatype_definition_integer.Name != datatype_definition_integerOther.Name {
		diffs = append(diffs, datatype_definition_integer.GongMarshallField(stage, "Name"))
	}
	if datatype_definition_integer.DESC != datatype_definition_integerOther.DESC {
		diffs = append(diffs, datatype_definition_integer.GongMarshallField(stage, "DESC"))
	}
	if datatype_definition_integer.IDENTIFIER != datatype_definition_integerOther.IDENTIFIER {
		diffs = append(diffs, datatype_definition_integer.GongMarshallField(stage, "IDENTIFIER"))
	}
	if datatype_definition_integer.LAST_CHANGE != datatype_definition_integerOther.LAST_CHANGE {
		diffs = append(diffs, datatype_definition_integer.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if datatype_definition_integer.LONG_NAME != datatype_definition_integerOther.LONG_NAME {
		diffs = append(diffs, datatype_definition_integer.GongMarshallField(stage, "LONG_NAME"))
	}
	if datatype_definition_integer.MAX != datatype_definition_integerOther.MAX {
		diffs = append(diffs, datatype_definition_integer.GongMarshallField(stage, "MAX"))
	}
	if datatype_definition_integer.MIN != datatype_definition_integerOther.MIN {
		diffs = append(diffs, datatype_definition_integer.GongMarshallField(stage, "MIN"))
	}
	if (datatype_definition_integer.ALTERNATIVE_ID == nil) != (datatype_definition_integerOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, datatype_definition_integer.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if datatype_definition_integer.ALTERNATIVE_ID != nil && datatype_definition_integerOther.ALTERNATIVE_ID != nil {
		if datatype_definition_integer.ALTERNATIVE_ID != datatype_definition_integerOther.ALTERNATIVE_ID {
			diffs = append(diffs, datatype_definition_integer.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (datatype_definition_real *DATATYPE_DEFINITION_REAL) GongDiff(stage *Stage, datatype_definition_realOther *DATATYPE_DEFINITION_REAL) (diffs []string) {
	// insertion point for field diffs
	if datatype_definition_real.Name != datatype_definition_realOther.Name {
		diffs = append(diffs, datatype_definition_real.GongMarshallField(stage, "Name"))
	}
	if datatype_definition_real.ACCURACY != datatype_definition_realOther.ACCURACY {
		diffs = append(diffs, datatype_definition_real.GongMarshallField(stage, "ACCURACY"))
	}
	if datatype_definition_real.DESC != datatype_definition_realOther.DESC {
		diffs = append(diffs, datatype_definition_real.GongMarshallField(stage, "DESC"))
	}
	if datatype_definition_real.IDENTIFIER != datatype_definition_realOther.IDENTIFIER {
		diffs = append(diffs, datatype_definition_real.GongMarshallField(stage, "IDENTIFIER"))
	}
	if datatype_definition_real.LAST_CHANGE != datatype_definition_realOther.LAST_CHANGE {
		diffs = append(diffs, datatype_definition_real.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if datatype_definition_real.LONG_NAME != datatype_definition_realOther.LONG_NAME {
		diffs = append(diffs, datatype_definition_real.GongMarshallField(stage, "LONG_NAME"))
	}
	if datatype_definition_real.MAX != datatype_definition_realOther.MAX {
		diffs = append(diffs, datatype_definition_real.GongMarshallField(stage, "MAX"))
	}
	if datatype_definition_real.MIN != datatype_definition_realOther.MIN {
		diffs = append(diffs, datatype_definition_real.GongMarshallField(stage, "MIN"))
	}
	if (datatype_definition_real.ALTERNATIVE_ID == nil) != (datatype_definition_realOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, datatype_definition_real.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if datatype_definition_real.ALTERNATIVE_ID != nil && datatype_definition_realOther.ALTERNATIVE_ID != nil {
		if datatype_definition_real.ALTERNATIVE_ID != datatype_definition_realOther.ALTERNATIVE_ID {
			diffs = append(diffs, datatype_definition_real.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (datatype_definition_string *DATATYPE_DEFINITION_STRING) GongDiff(stage *Stage, datatype_definition_stringOther *DATATYPE_DEFINITION_STRING) (diffs []string) {
	// insertion point for field diffs
	if datatype_definition_string.Name != datatype_definition_stringOther.Name {
		diffs = append(diffs, datatype_definition_string.GongMarshallField(stage, "Name"))
	}
	if datatype_definition_string.DESC != datatype_definition_stringOther.DESC {
		diffs = append(diffs, datatype_definition_string.GongMarshallField(stage, "DESC"))
	}
	if datatype_definition_string.IDENTIFIER != datatype_definition_stringOther.IDENTIFIER {
		diffs = append(diffs, datatype_definition_string.GongMarshallField(stage, "IDENTIFIER"))
	}
	if datatype_definition_string.LAST_CHANGE != datatype_definition_stringOther.LAST_CHANGE {
		diffs = append(diffs, datatype_definition_string.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if datatype_definition_string.LONG_NAME != datatype_definition_stringOther.LONG_NAME {
		diffs = append(diffs, datatype_definition_string.GongMarshallField(stage, "LONG_NAME"))
	}
	if datatype_definition_string.MAX_LENGTH != datatype_definition_stringOther.MAX_LENGTH {
		diffs = append(diffs, datatype_definition_string.GongMarshallField(stage, "MAX_LENGTH"))
	}
	if (datatype_definition_string.ALTERNATIVE_ID == nil) != (datatype_definition_stringOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, datatype_definition_string.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if datatype_definition_string.ALTERNATIVE_ID != nil && datatype_definition_stringOther.ALTERNATIVE_ID != nil {
		if datatype_definition_string.ALTERNATIVE_ID != datatype_definition_stringOther.ALTERNATIVE_ID {
			diffs = append(diffs, datatype_definition_string.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (datatype_definition_xhtml *DATATYPE_DEFINITION_XHTML) GongDiff(stage *Stage, datatype_definition_xhtmlOther *DATATYPE_DEFINITION_XHTML) (diffs []string) {
	// insertion point for field diffs
	if datatype_definition_xhtml.Name != datatype_definition_xhtmlOther.Name {
		diffs = append(diffs, datatype_definition_xhtml.GongMarshallField(stage, "Name"))
	}
	if datatype_definition_xhtml.DESC != datatype_definition_xhtmlOther.DESC {
		diffs = append(diffs, datatype_definition_xhtml.GongMarshallField(stage, "DESC"))
	}
	if datatype_definition_xhtml.IDENTIFIER != datatype_definition_xhtmlOther.IDENTIFIER {
		diffs = append(diffs, datatype_definition_xhtml.GongMarshallField(stage, "IDENTIFIER"))
	}
	if datatype_definition_xhtml.LAST_CHANGE != datatype_definition_xhtmlOther.LAST_CHANGE {
		diffs = append(diffs, datatype_definition_xhtml.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if datatype_definition_xhtml.LONG_NAME != datatype_definition_xhtmlOther.LONG_NAME {
		diffs = append(diffs, datatype_definition_xhtml.GongMarshallField(stage, "LONG_NAME"))
	}
	if (datatype_definition_xhtml.ALTERNATIVE_ID == nil) != (datatype_definition_xhtmlOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, datatype_definition_xhtml.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if datatype_definition_xhtml.ALTERNATIVE_ID != nil && datatype_definition_xhtmlOther.ALTERNATIVE_ID != nil {
		if datatype_definition_xhtml.ALTERNATIVE_ID != datatype_definition_xhtmlOther.ALTERNATIVE_ID {
			diffs = append(diffs, datatype_definition_xhtml.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (embedded_value *EMBEDDED_VALUE) GongDiff(stage *Stage, embedded_valueOther *EMBEDDED_VALUE) (diffs []string) {
	// insertion point for field diffs
	if embedded_value.Name != embedded_valueOther.Name {
		diffs = append(diffs, embedded_value.GongMarshallField(stage, "Name"))
	}
	if embedded_value.KEY != embedded_valueOther.KEY {
		diffs = append(diffs, embedded_value.GongMarshallField(stage, "KEY"))
	}
	if embedded_value.OTHER_CONTENT != embedded_valueOther.OTHER_CONTENT {
		diffs = append(diffs, embedded_value.GongMarshallField(stage, "OTHER_CONTENT"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (enum_value *ENUM_VALUE) GongDiff(stage *Stage, enum_valueOther *ENUM_VALUE) (diffs []string) {
	// insertion point for field diffs
	if enum_value.Name != enum_valueOther.Name {
		diffs = append(diffs, enum_value.GongMarshallField(stage, "Name"))
	}
	if enum_value.DESC != enum_valueOther.DESC {
		diffs = append(diffs, enum_value.GongMarshallField(stage, "DESC"))
	}
	if enum_value.IDENTIFIER != enum_valueOther.IDENTIFIER {
		diffs = append(diffs, enum_value.GongMarshallField(stage, "IDENTIFIER"))
	}
	if enum_value.LAST_CHANGE != enum_valueOther.LAST_CHANGE {
		diffs = append(diffs, enum_value.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if enum_value.LONG_NAME != enum_valueOther.LONG_NAME {
		diffs = append(diffs, enum_value.GongMarshallField(stage, "LONG_NAME"))
	}
	if (enum_value.ALTERNATIVE_ID == nil) != (enum_valueOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, enum_value.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if enum_value.ALTERNATIVE_ID != nil && enum_valueOther.ALTERNATIVE_ID != nil {
		if enum_value.ALTERNATIVE_ID != enum_valueOther.ALTERNATIVE_ID {
			diffs = append(diffs, enum_value.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (enum_value.PROPERTIES == nil) != (enum_valueOther.PROPERTIES == nil) {
		diffs = append(diffs, enum_value.GongMarshallField(stage, "PROPERTIES"))
	} else if enum_value.PROPERTIES != nil && enum_valueOther.PROPERTIES != nil {
		if enum_value.PROPERTIES != enum_valueOther.PROPERTIES {
			diffs = append(diffs, enum_value.GongMarshallField(stage, "PROPERTIES"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (relation_group *RELATION_GROUP) GongDiff(stage *Stage, relation_groupOther *RELATION_GROUP) (diffs []string) {
	// insertion point for field diffs
	if relation_group.Name != relation_groupOther.Name {
		diffs = append(diffs, relation_group.GongMarshallField(stage, "Name"))
	}
	if relation_group.DESC != relation_groupOther.DESC {
		diffs = append(diffs, relation_group.GongMarshallField(stage, "DESC"))
	}
	if relation_group.IDENTIFIER != relation_groupOther.IDENTIFIER {
		diffs = append(diffs, relation_group.GongMarshallField(stage, "IDENTIFIER"))
	}
	if relation_group.LAST_CHANGE != relation_groupOther.LAST_CHANGE {
		diffs = append(diffs, relation_group.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if relation_group.LONG_NAME != relation_groupOther.LONG_NAME {
		diffs = append(diffs, relation_group.GongMarshallField(stage, "LONG_NAME"))
	}
	if (relation_group.ALTERNATIVE_ID == nil) != (relation_groupOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, relation_group.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if relation_group.ALTERNATIVE_ID != nil && relation_groupOther.ALTERNATIVE_ID != nil {
		if relation_group.ALTERNATIVE_ID != relation_groupOther.ALTERNATIVE_ID {
			diffs = append(diffs, relation_group.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (relation_group.SOURCE_SPECIFICATION == nil) != (relation_groupOther.SOURCE_SPECIFICATION == nil) {
		diffs = append(diffs, relation_group.GongMarshallField(stage, "SOURCE_SPECIFICATION"))
	} else if relation_group.SOURCE_SPECIFICATION != nil && relation_groupOther.SOURCE_SPECIFICATION != nil {
		if relation_group.SOURCE_SPECIFICATION != relation_groupOther.SOURCE_SPECIFICATION {
			diffs = append(diffs, relation_group.GongMarshallField(stage, "SOURCE_SPECIFICATION"))
		}
	}
	if (relation_group.SPEC_RELATIONS == nil) != (relation_groupOther.SPEC_RELATIONS == nil) {
		diffs = append(diffs, relation_group.GongMarshallField(stage, "SPEC_RELATIONS"))
	} else if relation_group.SPEC_RELATIONS != nil && relation_groupOther.SPEC_RELATIONS != nil {
		if relation_group.SPEC_RELATIONS != relation_groupOther.SPEC_RELATIONS {
			diffs = append(diffs, relation_group.GongMarshallField(stage, "SPEC_RELATIONS"))
		}
	}
	if (relation_group.TARGET_SPECIFICATION == nil) != (relation_groupOther.TARGET_SPECIFICATION == nil) {
		diffs = append(diffs, relation_group.GongMarshallField(stage, "TARGET_SPECIFICATION"))
	} else if relation_group.TARGET_SPECIFICATION != nil && relation_groupOther.TARGET_SPECIFICATION != nil {
		if relation_group.TARGET_SPECIFICATION != relation_groupOther.TARGET_SPECIFICATION {
			diffs = append(diffs, relation_group.GongMarshallField(stage, "TARGET_SPECIFICATION"))
		}
	}
	if (relation_group.TYPE == nil) != (relation_groupOther.TYPE == nil) {
		diffs = append(diffs, relation_group.GongMarshallField(stage, "TYPE"))
	} else if relation_group.TYPE != nil && relation_groupOther.TYPE != nil {
		if relation_group.TYPE != relation_groupOther.TYPE {
			diffs = append(diffs, relation_group.GongMarshallField(stage, "TYPE"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (relation_group_type *RELATION_GROUP_TYPE) GongDiff(stage *Stage, relation_group_typeOther *RELATION_GROUP_TYPE) (diffs []string) {
	// insertion point for field diffs
	if relation_group_type.Name != relation_group_typeOther.Name {
		diffs = append(diffs, relation_group_type.GongMarshallField(stage, "Name"))
	}
	if relation_group_type.DESC != relation_group_typeOther.DESC {
		diffs = append(diffs, relation_group_type.GongMarshallField(stage, "DESC"))
	}
	if relation_group_type.IDENTIFIER != relation_group_typeOther.IDENTIFIER {
		diffs = append(diffs, relation_group_type.GongMarshallField(stage, "IDENTIFIER"))
	}
	if relation_group_type.LAST_CHANGE != relation_group_typeOther.LAST_CHANGE {
		diffs = append(diffs, relation_group_type.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if relation_group_type.LONG_NAME != relation_group_typeOther.LONG_NAME {
		diffs = append(diffs, relation_group_type.GongMarshallField(stage, "LONG_NAME"))
	}
	if (relation_group_type.ALTERNATIVE_ID == nil) != (relation_group_typeOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, relation_group_type.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if relation_group_type.ALTERNATIVE_ID != nil && relation_group_typeOther.ALTERNATIVE_ID != nil {
		if relation_group_type.ALTERNATIVE_ID != relation_group_typeOther.ALTERNATIVE_ID {
			diffs = append(diffs, relation_group_type.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (relation_group_type.SPEC_ATTRIBUTES == nil) != (relation_group_typeOther.SPEC_ATTRIBUTES == nil) {
		diffs = append(diffs, relation_group_type.GongMarshallField(stage, "SPEC_ATTRIBUTES"))
	} else if relation_group_type.SPEC_ATTRIBUTES != nil && relation_group_typeOther.SPEC_ATTRIBUTES != nil {
		if relation_group_type.SPEC_ATTRIBUTES != relation_group_typeOther.SPEC_ATTRIBUTES {
			diffs = append(diffs, relation_group_type.GongMarshallField(stage, "SPEC_ATTRIBUTES"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (req_if *REQ_IF) GongDiff(stage *Stage, req_ifOther *REQ_IF) (diffs []string) {
	// insertion point for field diffs
	if req_if.Name != req_ifOther.Name {
		diffs = append(diffs, req_if.GongMarshallField(stage, "Name"))
	}
	if req_if.Lang != req_ifOther.Lang {
		diffs = append(diffs, req_if.GongMarshallField(stage, "Lang"))
	}
	if (req_if.THE_HEADER == nil) != (req_ifOther.THE_HEADER == nil) {
		diffs = append(diffs, req_if.GongMarshallField(stage, "THE_HEADER"))
	} else if req_if.THE_HEADER != nil && req_ifOther.THE_HEADER != nil {
		if req_if.THE_HEADER != req_ifOther.THE_HEADER {
			diffs = append(diffs, req_if.GongMarshallField(stage, "THE_HEADER"))
		}
	}
	if (req_if.CORE_CONTENT == nil) != (req_ifOther.CORE_CONTENT == nil) {
		diffs = append(diffs, req_if.GongMarshallField(stage, "CORE_CONTENT"))
	} else if req_if.CORE_CONTENT != nil && req_ifOther.CORE_CONTENT != nil {
		if req_if.CORE_CONTENT != req_ifOther.CORE_CONTENT {
			diffs = append(diffs, req_if.GongMarshallField(stage, "CORE_CONTENT"))
		}
	}
	if (req_if.TOOL_EXTENSIONS == nil) != (req_ifOther.TOOL_EXTENSIONS == nil) {
		diffs = append(diffs, req_if.GongMarshallField(stage, "TOOL_EXTENSIONS"))
	} else if req_if.TOOL_EXTENSIONS != nil && req_ifOther.TOOL_EXTENSIONS != nil {
		if req_if.TOOL_EXTENSIONS != req_ifOther.TOOL_EXTENSIONS {
			diffs = append(diffs, req_if.GongMarshallField(stage, "TOOL_EXTENSIONS"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (req_if_content *REQ_IF_CONTENT) GongDiff(stage *Stage, req_if_contentOther *REQ_IF_CONTENT) (diffs []string) {
	// insertion point for field diffs
	if req_if_content.Name != req_if_contentOther.Name {
		diffs = append(diffs, req_if_content.GongMarshallField(stage, "Name"))
	}
	if (req_if_content.DATATYPES == nil) != (req_if_contentOther.DATATYPES == nil) {
		diffs = append(diffs, req_if_content.GongMarshallField(stage, "DATATYPES"))
	} else if req_if_content.DATATYPES != nil && req_if_contentOther.DATATYPES != nil {
		if req_if_content.DATATYPES != req_if_contentOther.DATATYPES {
			diffs = append(diffs, req_if_content.GongMarshallField(stage, "DATATYPES"))
		}
	}
	if (req_if_content.SPEC_TYPES == nil) != (req_if_contentOther.SPEC_TYPES == nil) {
		diffs = append(diffs, req_if_content.GongMarshallField(stage, "SPEC_TYPES"))
	} else if req_if_content.SPEC_TYPES != nil && req_if_contentOther.SPEC_TYPES != nil {
		if req_if_content.SPEC_TYPES != req_if_contentOther.SPEC_TYPES {
			diffs = append(diffs, req_if_content.GongMarshallField(stage, "SPEC_TYPES"))
		}
	}
	if (req_if_content.SPEC_OBJECTS == nil) != (req_if_contentOther.SPEC_OBJECTS == nil) {
		diffs = append(diffs, req_if_content.GongMarshallField(stage, "SPEC_OBJECTS"))
	} else if req_if_content.SPEC_OBJECTS != nil && req_if_contentOther.SPEC_OBJECTS != nil {
		if req_if_content.SPEC_OBJECTS != req_if_contentOther.SPEC_OBJECTS {
			diffs = append(diffs, req_if_content.GongMarshallField(stage, "SPEC_OBJECTS"))
		}
	}
	if (req_if_content.SPEC_RELATIONS == nil) != (req_if_contentOther.SPEC_RELATIONS == nil) {
		diffs = append(diffs, req_if_content.GongMarshallField(stage, "SPEC_RELATIONS"))
	} else if req_if_content.SPEC_RELATIONS != nil && req_if_contentOther.SPEC_RELATIONS != nil {
		if req_if_content.SPEC_RELATIONS != req_if_contentOther.SPEC_RELATIONS {
			diffs = append(diffs, req_if_content.GongMarshallField(stage, "SPEC_RELATIONS"))
		}
	}
	if (req_if_content.SPECIFICATIONS == nil) != (req_if_contentOther.SPECIFICATIONS == nil) {
		diffs = append(diffs, req_if_content.GongMarshallField(stage, "SPECIFICATIONS"))
	} else if req_if_content.SPECIFICATIONS != nil && req_if_contentOther.SPECIFICATIONS != nil {
		if req_if_content.SPECIFICATIONS != req_if_contentOther.SPECIFICATIONS {
			diffs = append(diffs, req_if_content.GongMarshallField(stage, "SPECIFICATIONS"))
		}
	}
	if (req_if_content.SPEC_RELATION_GROUPS == nil) != (req_if_contentOther.SPEC_RELATION_GROUPS == nil) {
		diffs = append(diffs, req_if_content.GongMarshallField(stage, "SPEC_RELATION_GROUPS"))
	} else if req_if_content.SPEC_RELATION_GROUPS != nil && req_if_contentOther.SPEC_RELATION_GROUPS != nil {
		if req_if_content.SPEC_RELATION_GROUPS != req_if_contentOther.SPEC_RELATION_GROUPS {
			diffs = append(diffs, req_if_content.GongMarshallField(stage, "SPEC_RELATION_GROUPS"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (req_if_header *REQ_IF_HEADER) GongDiff(stage *Stage, req_if_headerOther *REQ_IF_HEADER) (diffs []string) {
	// insertion point for field diffs
	if req_if_header.Name != req_if_headerOther.Name {
		diffs = append(diffs, req_if_header.GongMarshallField(stage, "Name"))
	}
	if req_if_header.IDENTIFIER != req_if_headerOther.IDENTIFIER {
		diffs = append(diffs, req_if_header.GongMarshallField(stage, "IDENTIFIER"))
	}
	if req_if_header.COMMENT != req_if_headerOther.COMMENT {
		diffs = append(diffs, req_if_header.GongMarshallField(stage, "COMMENT"))
	}
	if req_if_header.CREATION_TIME != req_if_headerOther.CREATION_TIME {
		diffs = append(diffs, req_if_header.GongMarshallField(stage, "CREATION_TIME"))
	}
	if req_if_header.REPOSITORY_ID != req_if_headerOther.REPOSITORY_ID {
		diffs = append(diffs, req_if_header.GongMarshallField(stage, "REPOSITORY_ID"))
	}
	if req_if_header.REQ_IF_TOOL_ID != req_if_headerOther.REQ_IF_TOOL_ID {
		diffs = append(diffs, req_if_header.GongMarshallField(stage, "REQ_IF_TOOL_ID"))
	}
	if req_if_header.REQ_IF_VERSION != req_if_headerOther.REQ_IF_VERSION {
		diffs = append(diffs, req_if_header.GongMarshallField(stage, "REQ_IF_VERSION"))
	}
	if req_if_header.SOURCE_TOOL_ID != req_if_headerOther.SOURCE_TOOL_ID {
		diffs = append(diffs, req_if_header.GongMarshallField(stage, "SOURCE_TOOL_ID"))
	}
	if req_if_header.TITLE != req_if_headerOther.TITLE {
		diffs = append(diffs, req_if_header.GongMarshallField(stage, "TITLE"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (req_if_tool_extension *REQ_IF_TOOL_EXTENSION) GongDiff(stage *Stage, req_if_tool_extensionOther *REQ_IF_TOOL_EXTENSION) (diffs []string) {
	// insertion point for field diffs
	if req_if_tool_extension.Name != req_if_tool_extensionOther.Name {
		diffs = append(diffs, req_if_tool_extension.GongMarshallField(stage, "Name"))
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (specification *SPECIFICATION) GongDiff(stage *Stage, specificationOther *SPECIFICATION) (diffs []string) {
	// insertion point for field diffs
	if specification.Name != specificationOther.Name {
		diffs = append(diffs, specification.GongMarshallField(stage, "Name"))
	}
	if specification.DESC != specificationOther.DESC {
		diffs = append(diffs, specification.GongMarshallField(stage, "DESC"))
	}
	if specification.IDENTIFIER != specificationOther.IDENTIFIER {
		diffs = append(diffs, specification.GongMarshallField(stage, "IDENTIFIER"))
	}
	if specification.LAST_CHANGE != specificationOther.LAST_CHANGE {
		diffs = append(diffs, specification.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if specification.LONG_NAME != specificationOther.LONG_NAME {
		diffs = append(diffs, specification.GongMarshallField(stage, "LONG_NAME"))
	}
	if (specification.ALTERNATIVE_ID == nil) != (specificationOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, specification.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if specification.ALTERNATIVE_ID != nil && specificationOther.ALTERNATIVE_ID != nil {
		if specification.ALTERNATIVE_ID != specificationOther.ALTERNATIVE_ID {
			diffs = append(diffs, specification.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (specification.CHILDREN == nil) != (specificationOther.CHILDREN == nil) {
		diffs = append(diffs, specification.GongMarshallField(stage, "CHILDREN"))
	} else if specification.CHILDREN != nil && specificationOther.CHILDREN != nil {
		if specification.CHILDREN != specificationOther.CHILDREN {
			diffs = append(diffs, specification.GongMarshallField(stage, "CHILDREN"))
		}
	}
	if (specification.VALUES == nil) != (specificationOther.VALUES == nil) {
		diffs = append(diffs, specification.GongMarshallField(stage, "VALUES"))
	} else if specification.VALUES != nil && specificationOther.VALUES != nil {
		if specification.VALUES != specificationOther.VALUES {
			diffs = append(diffs, specification.GongMarshallField(stage, "VALUES"))
		}
	}
	if (specification.TYPE == nil) != (specificationOther.TYPE == nil) {
		diffs = append(diffs, specification.GongMarshallField(stage, "TYPE"))
	} else if specification.TYPE != nil && specificationOther.TYPE != nil {
		if specification.TYPE != specificationOther.TYPE {
			diffs = append(diffs, specification.GongMarshallField(stage, "TYPE"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (specification_type *SPECIFICATION_TYPE) GongDiff(stage *Stage, specification_typeOther *SPECIFICATION_TYPE) (diffs []string) {
	// insertion point for field diffs
	if specification_type.Name != specification_typeOther.Name {
		diffs = append(diffs, specification_type.GongMarshallField(stage, "Name"))
	}
	if specification_type.DESC != specification_typeOther.DESC {
		diffs = append(diffs, specification_type.GongMarshallField(stage, "DESC"))
	}
	if specification_type.IDENTIFIER != specification_typeOther.IDENTIFIER {
		diffs = append(diffs, specification_type.GongMarshallField(stage, "IDENTIFIER"))
	}
	if specification_type.LAST_CHANGE != specification_typeOther.LAST_CHANGE {
		diffs = append(diffs, specification_type.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if specification_type.LONG_NAME != specification_typeOther.LONG_NAME {
		diffs = append(diffs, specification_type.GongMarshallField(stage, "LONG_NAME"))
	}
	if (specification_type.ALTERNATIVE_ID == nil) != (specification_typeOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, specification_type.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if specification_type.ALTERNATIVE_ID != nil && specification_typeOther.ALTERNATIVE_ID != nil {
		if specification_type.ALTERNATIVE_ID != specification_typeOther.ALTERNATIVE_ID {
			diffs = append(diffs, specification_type.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (specification_type.SPEC_ATTRIBUTES == nil) != (specification_typeOther.SPEC_ATTRIBUTES == nil) {
		diffs = append(diffs, specification_type.GongMarshallField(stage, "SPEC_ATTRIBUTES"))
	} else if specification_type.SPEC_ATTRIBUTES != nil && specification_typeOther.SPEC_ATTRIBUTES != nil {
		if specification_type.SPEC_ATTRIBUTES != specification_typeOther.SPEC_ATTRIBUTES {
			diffs = append(diffs, specification_type.GongMarshallField(stage, "SPEC_ATTRIBUTES"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (spec_hierarchy *SPEC_HIERARCHY) GongDiff(stage *Stage, spec_hierarchyOther *SPEC_HIERARCHY) (diffs []string) {
	// insertion point for field diffs
	if spec_hierarchy.Name != spec_hierarchyOther.Name {
		diffs = append(diffs, spec_hierarchy.GongMarshallField(stage, "Name"))
	}
	if spec_hierarchy.DESC != spec_hierarchyOther.DESC {
		diffs = append(diffs, spec_hierarchy.GongMarshallField(stage, "DESC"))
	}
	if spec_hierarchy.IDENTIFIER != spec_hierarchyOther.IDENTIFIER {
		diffs = append(diffs, spec_hierarchy.GongMarshallField(stage, "IDENTIFIER"))
	}
	if spec_hierarchy.IS_EDITABLE != spec_hierarchyOther.IS_EDITABLE {
		diffs = append(diffs, spec_hierarchy.GongMarshallField(stage, "IS_EDITABLE"))
	}
	if spec_hierarchy.IS_TABLE_INTERNAL != spec_hierarchyOther.IS_TABLE_INTERNAL {
		diffs = append(diffs, spec_hierarchy.GongMarshallField(stage, "IS_TABLE_INTERNAL"))
	}
	if spec_hierarchy.LAST_CHANGE != spec_hierarchyOther.LAST_CHANGE {
		diffs = append(diffs, spec_hierarchy.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if spec_hierarchy.LONG_NAME != spec_hierarchyOther.LONG_NAME {
		diffs = append(diffs, spec_hierarchy.GongMarshallField(stage, "LONG_NAME"))
	}
	if (spec_hierarchy.ALTERNATIVE_ID == nil) != (spec_hierarchyOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, spec_hierarchy.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if spec_hierarchy.ALTERNATIVE_ID != nil && spec_hierarchyOther.ALTERNATIVE_ID != nil {
		if spec_hierarchy.ALTERNATIVE_ID != spec_hierarchyOther.ALTERNATIVE_ID {
			diffs = append(diffs, spec_hierarchy.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (spec_hierarchy.CHILDREN == nil) != (spec_hierarchyOther.CHILDREN == nil) {
		diffs = append(diffs, spec_hierarchy.GongMarshallField(stage, "CHILDREN"))
	} else if spec_hierarchy.CHILDREN != nil && spec_hierarchyOther.CHILDREN != nil {
		if spec_hierarchy.CHILDREN != spec_hierarchyOther.CHILDREN {
			diffs = append(diffs, spec_hierarchy.GongMarshallField(stage, "CHILDREN"))
		}
	}
	if (spec_hierarchy.EDITABLE_ATTS == nil) != (spec_hierarchyOther.EDITABLE_ATTS == nil) {
		diffs = append(diffs, spec_hierarchy.GongMarshallField(stage, "EDITABLE_ATTS"))
	} else if spec_hierarchy.EDITABLE_ATTS != nil && spec_hierarchyOther.EDITABLE_ATTS != nil {
		if spec_hierarchy.EDITABLE_ATTS != spec_hierarchyOther.EDITABLE_ATTS {
			diffs = append(diffs, spec_hierarchy.GongMarshallField(stage, "EDITABLE_ATTS"))
		}
	}
	if (spec_hierarchy.OBJECT == nil) != (spec_hierarchyOther.OBJECT == nil) {
		diffs = append(diffs, spec_hierarchy.GongMarshallField(stage, "OBJECT"))
	} else if spec_hierarchy.OBJECT != nil && spec_hierarchyOther.OBJECT != nil {
		if spec_hierarchy.OBJECT != spec_hierarchyOther.OBJECT {
			diffs = append(diffs, spec_hierarchy.GongMarshallField(stage, "OBJECT"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (spec_object *SPEC_OBJECT) GongDiff(stage *Stage, spec_objectOther *SPEC_OBJECT) (diffs []string) {
	// insertion point for field diffs
	if spec_object.Name != spec_objectOther.Name {
		diffs = append(diffs, spec_object.GongMarshallField(stage, "Name"))
	}
	if spec_object.DESC != spec_objectOther.DESC {
		diffs = append(diffs, spec_object.GongMarshallField(stage, "DESC"))
	}
	if spec_object.IDENTIFIER != spec_objectOther.IDENTIFIER {
		diffs = append(diffs, spec_object.GongMarshallField(stage, "IDENTIFIER"))
	}
	if spec_object.LAST_CHANGE != spec_objectOther.LAST_CHANGE {
		diffs = append(diffs, spec_object.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if spec_object.LONG_NAME != spec_objectOther.LONG_NAME {
		diffs = append(diffs, spec_object.GongMarshallField(stage, "LONG_NAME"))
	}
	if (spec_object.ALTERNATIVE_ID == nil) != (spec_objectOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, spec_object.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if spec_object.ALTERNATIVE_ID != nil && spec_objectOther.ALTERNATIVE_ID != nil {
		if spec_object.ALTERNATIVE_ID != spec_objectOther.ALTERNATIVE_ID {
			diffs = append(diffs, spec_object.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (spec_object.VALUES == nil) != (spec_objectOther.VALUES == nil) {
		diffs = append(diffs, spec_object.GongMarshallField(stage, "VALUES"))
	} else if spec_object.VALUES != nil && spec_objectOther.VALUES != nil {
		if spec_object.VALUES != spec_objectOther.VALUES {
			diffs = append(diffs, spec_object.GongMarshallField(stage, "VALUES"))
		}
	}
	if (spec_object.TYPE == nil) != (spec_objectOther.TYPE == nil) {
		diffs = append(diffs, spec_object.GongMarshallField(stage, "TYPE"))
	} else if spec_object.TYPE != nil && spec_objectOther.TYPE != nil {
		if spec_object.TYPE != spec_objectOther.TYPE {
			diffs = append(diffs, spec_object.GongMarshallField(stage, "TYPE"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (spec_object_type *SPEC_OBJECT_TYPE) GongDiff(stage *Stage, spec_object_typeOther *SPEC_OBJECT_TYPE) (diffs []string) {
	// insertion point for field diffs
	if spec_object_type.Name != spec_object_typeOther.Name {
		diffs = append(diffs, spec_object_type.GongMarshallField(stage, "Name"))
	}
	if spec_object_type.DESC != spec_object_typeOther.DESC {
		diffs = append(diffs, spec_object_type.GongMarshallField(stage, "DESC"))
	}
	if spec_object_type.IDENTIFIER != spec_object_typeOther.IDENTIFIER {
		diffs = append(diffs, spec_object_type.GongMarshallField(stage, "IDENTIFIER"))
	}
	if spec_object_type.LAST_CHANGE != spec_object_typeOther.LAST_CHANGE {
		diffs = append(diffs, spec_object_type.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if spec_object_type.LONG_NAME != spec_object_typeOther.LONG_NAME {
		diffs = append(diffs, spec_object_type.GongMarshallField(stage, "LONG_NAME"))
	}
	if (spec_object_type.ALTERNATIVE_ID == nil) != (spec_object_typeOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, spec_object_type.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if spec_object_type.ALTERNATIVE_ID != nil && spec_object_typeOther.ALTERNATIVE_ID != nil {
		if spec_object_type.ALTERNATIVE_ID != spec_object_typeOther.ALTERNATIVE_ID {
			diffs = append(diffs, spec_object_type.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (spec_object_type.SPEC_ATTRIBUTES == nil) != (spec_object_typeOther.SPEC_ATTRIBUTES == nil) {
		diffs = append(diffs, spec_object_type.GongMarshallField(stage, "SPEC_ATTRIBUTES"))
	} else if spec_object_type.SPEC_ATTRIBUTES != nil && spec_object_typeOther.SPEC_ATTRIBUTES != nil {
		if spec_object_type.SPEC_ATTRIBUTES != spec_object_typeOther.SPEC_ATTRIBUTES {
			diffs = append(diffs, spec_object_type.GongMarshallField(stage, "SPEC_ATTRIBUTES"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (spec_relation *SPEC_RELATION) GongDiff(stage *Stage, spec_relationOther *SPEC_RELATION) (diffs []string) {
	// insertion point for field diffs
	if spec_relation.Name != spec_relationOther.Name {
		diffs = append(diffs, spec_relation.GongMarshallField(stage, "Name"))
	}
	if spec_relation.DESC != spec_relationOther.DESC {
		diffs = append(diffs, spec_relation.GongMarshallField(stage, "DESC"))
	}
	if spec_relation.IDENTIFIER != spec_relationOther.IDENTIFIER {
		diffs = append(diffs, spec_relation.GongMarshallField(stage, "IDENTIFIER"))
	}
	if spec_relation.LAST_CHANGE != spec_relationOther.LAST_CHANGE {
		diffs = append(diffs, spec_relation.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if spec_relation.LONG_NAME != spec_relationOther.LONG_NAME {
		diffs = append(diffs, spec_relation.GongMarshallField(stage, "LONG_NAME"))
	}
	if (spec_relation.ALTERNATIVE_ID == nil) != (spec_relationOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, spec_relation.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if spec_relation.ALTERNATIVE_ID != nil && spec_relationOther.ALTERNATIVE_ID != nil {
		if spec_relation.ALTERNATIVE_ID != spec_relationOther.ALTERNATIVE_ID {
			diffs = append(diffs, spec_relation.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (spec_relation.VALUES == nil) != (spec_relationOther.VALUES == nil) {
		diffs = append(diffs, spec_relation.GongMarshallField(stage, "VALUES"))
	} else if spec_relation.VALUES != nil && spec_relationOther.VALUES != nil {
		if spec_relation.VALUES != spec_relationOther.VALUES {
			diffs = append(diffs, spec_relation.GongMarshallField(stage, "VALUES"))
		}
	}
	if (spec_relation.SOURCE == nil) != (spec_relationOther.SOURCE == nil) {
		diffs = append(diffs, spec_relation.GongMarshallField(stage, "SOURCE"))
	} else if spec_relation.SOURCE != nil && spec_relationOther.SOURCE != nil {
		if spec_relation.SOURCE != spec_relationOther.SOURCE {
			diffs = append(diffs, spec_relation.GongMarshallField(stage, "SOURCE"))
		}
	}
	if (spec_relation.TARGET == nil) != (spec_relationOther.TARGET == nil) {
		diffs = append(diffs, spec_relation.GongMarshallField(stage, "TARGET"))
	} else if spec_relation.TARGET != nil && spec_relationOther.TARGET != nil {
		if spec_relation.TARGET != spec_relationOther.TARGET {
			diffs = append(diffs, spec_relation.GongMarshallField(stage, "TARGET"))
		}
	}
	if (spec_relation.TYPE == nil) != (spec_relationOther.TYPE == nil) {
		diffs = append(diffs, spec_relation.GongMarshallField(stage, "TYPE"))
	} else if spec_relation.TYPE != nil && spec_relationOther.TYPE != nil {
		if spec_relation.TYPE != spec_relationOther.TYPE {
			diffs = append(diffs, spec_relation.GongMarshallField(stage, "TYPE"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (spec_relation_type *SPEC_RELATION_TYPE) GongDiff(stage *Stage, spec_relation_typeOther *SPEC_RELATION_TYPE) (diffs []string) {
	// insertion point for field diffs
	if spec_relation_type.Name != spec_relation_typeOther.Name {
		diffs = append(diffs, spec_relation_type.GongMarshallField(stage, "Name"))
	}
	if spec_relation_type.DESC != spec_relation_typeOther.DESC {
		diffs = append(diffs, spec_relation_type.GongMarshallField(stage, "DESC"))
	}
	if spec_relation_type.IDENTIFIER != spec_relation_typeOther.IDENTIFIER {
		diffs = append(diffs, spec_relation_type.GongMarshallField(stage, "IDENTIFIER"))
	}
	if spec_relation_type.LAST_CHANGE != spec_relation_typeOther.LAST_CHANGE {
		diffs = append(diffs, spec_relation_type.GongMarshallField(stage, "LAST_CHANGE"))
	}
	if spec_relation_type.LONG_NAME != spec_relation_typeOther.LONG_NAME {
		diffs = append(diffs, spec_relation_type.GongMarshallField(stage, "LONG_NAME"))
	}
	if (spec_relation_type.ALTERNATIVE_ID == nil) != (spec_relation_typeOther.ALTERNATIVE_ID == nil) {
		diffs = append(diffs, spec_relation_type.GongMarshallField(stage, "ALTERNATIVE_ID"))
	} else if spec_relation_type.ALTERNATIVE_ID != nil && spec_relation_typeOther.ALTERNATIVE_ID != nil {
		if spec_relation_type.ALTERNATIVE_ID != spec_relation_typeOther.ALTERNATIVE_ID {
			diffs = append(diffs, spec_relation_type.GongMarshallField(stage, "ALTERNATIVE_ID"))
		}
	}
	if (spec_relation_type.SPEC_ATTRIBUTES == nil) != (spec_relation_typeOther.SPEC_ATTRIBUTES == nil) {
		diffs = append(diffs, spec_relation_type.GongMarshallField(stage, "SPEC_ATTRIBUTES"))
	} else if spec_relation_type.SPEC_ATTRIBUTES != nil && spec_relation_typeOther.SPEC_ATTRIBUTES != nil {
		if spec_relation_type.SPEC_ATTRIBUTES != spec_relation_typeOther.SPEC_ATTRIBUTES {
			diffs = append(diffs, spec_relation_type.GongMarshallField(stage, "SPEC_ATTRIBUTES"))
		}
	}

	return
}

// GongDiff computes the diff between the instance and another instance of same gong struct type
// and returns the list of differences as strings
func (xhtml_content *XHTML_CONTENT) GongDiff(stage *Stage, xhtml_contentOther *XHTML_CONTENT) (diffs []string) {
	// insertion point for field diffs
	if xhtml_content.Name != xhtml_contentOther.Name {
		diffs = append(diffs, xhtml_content.GongMarshallField(stage, "Name"))
	}
	if xhtml_content.EnclosedText != xhtml_contentOther.EnclosedText {
		diffs = append(diffs, xhtml_content.GongMarshallField(stage, "EnclosedText"))
	}

	return
}

// Diff is the Stage method that returns the sequence of operations to transform oldSlice into newSlice.
func (stage *Stage) Diff(
	a GongstructIF,
	fieldName string,
	lenOld, lenNew int,
	equal func(i, j int) bool,
	getNewIdentifier func(j int) string,
) (ops string) {
	m, n := lenOld, lenNew

	// 1. Build the LCS (Longest Common Subsequence) Matrix
	// This helps us find the "anchor" elements that shouldn't move.
	dp := make([][]int, m+1)
	for i := range dp {
		dp[i] = make([]int, n+1)
	}

	for i := range m {
		for j := range n {
			if equal(i, j) {
				dp[i+1][j+1] = dp[i][j] + 1
			} else {
				// Take the maximum of previous options
				dp[i+1][j+1] = max(dp[i][j+1], dp[i+1][j])
			}
		}
	}

	// 2. Backtrack to find which indices in oldSlice are part of the LCS
	// We use a map for O(1) lookups.
	keptIndices := make(map[int]bool)
	i, j := m, n
	for i > 0 && j > 0 {
		if equal(i-1, j-1) {
			keptIndices[i-1] = true
			i--
			j--
		} else if dp[i-1][j] > dp[i][j-1] {
			i--
		} else {
			j--
		}
	}

	// 3. PHASE 1: Generate Deletions
	// MUST go from High Index -> Low Index to preserve validity of lower indices.
	for k := m - 1; k >= 0; k-- {
		if !keptIndices[k] {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Delete( %s.%s, %d, %d)", a.GongGetReferenceIdentifier(stage), fieldName, a.GongGetReferenceIdentifier(stage), fieldName, k, k+1)
		}
	}

	// 4. PHASE 2: Generate Insertions
	// We simulate the state of the slice after deletions to determine insertion points.
	// The 'current' slice essentially consists of only the kept LCS items.

	// Track kept indices in old slice
	keptOldIndices := make([]int, 0, len(keptIndices))
	for k := range m {
		if keptIndices[k] {
			keptOldIndices = append(keptOldIndices, k)
		}
	}

	lcsIdx := 0
	// Iterate through the NEW slice. If it matches the current LCS head, we keep it.
	// If it doesn't match, it must be inserted here.
	for k := range n {
		if lcsIdx < len(keptOldIndices) && equal(keptOldIndices[lcsIdx], k) {
			lcsIdx++
		} else {
			ops += fmt.Sprintf("\n\t%s.%s = slices.Insert( %s.%s, %d, %s)", a.GongGetIdentifier(stage), fieldName, a.GongGetIdentifier(stage), fieldName, k, getNewIdentifier(k))
		}
	}

	return ops
}
