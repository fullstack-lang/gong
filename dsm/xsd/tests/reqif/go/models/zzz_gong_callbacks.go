// generated code - do not edit
package models

// AfterCreateFromFront is called after a create from front
func AfterCreateFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *ALTERNATIVE_ID:
		if stage.OnAfterALTERNATIVE_IDCreateCallback != nil {
			stage.OnAfterALTERNATIVE_IDCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		if stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANCreateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_DATE:
		if stage.OnAfterATTRIBUTE_DEFINITION_DATECreateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_DATECreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		if stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONCreateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_INTEGER:
		if stage.OnAfterATTRIBUTE_DEFINITION_INTEGERCreateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_INTEGERCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_REAL:
		if stage.OnAfterATTRIBUTE_DEFINITION_REALCreateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_REALCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_STRING:
		if stage.OnAfterATTRIBUTE_DEFINITION_STRINGCreateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_STRINGCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_XHTML:
		if stage.OnAfterATTRIBUTE_DEFINITION_XHTMLCreateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_XHTMLCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_VALUE_BOOLEAN:
		if stage.OnAfterATTRIBUTE_VALUE_BOOLEANCreateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_BOOLEANCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_VALUE_DATE:
		if stage.OnAfterATTRIBUTE_VALUE_DATECreateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_DATECreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_VALUE_ENUMERATION:
		if stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONCreateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_VALUE_INTEGER:
		if stage.OnAfterATTRIBUTE_VALUE_INTEGERCreateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_INTEGERCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_VALUE_REAL:
		if stage.OnAfterATTRIBUTE_VALUE_REALCreateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_REALCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_VALUE_STRING:
		if stage.OnAfterATTRIBUTE_VALUE_STRINGCreateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_STRINGCreateCallback.OnAfterCreate(stage, target)
		}
	case *ATTRIBUTE_VALUE_XHTML:
		if stage.OnAfterATTRIBUTE_VALUE_XHTMLCreateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_XHTMLCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ALTERNATIVE_ID:
		if stage.OnAfterA_ALTERNATIVE_IDCreateCallback != nil {
			stage.OnAfterA_ALTERNATIVE_IDCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		if stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_DATE:
		if stage.OnAfterA_ATTRIBUTE_VALUE_DATECreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_DATECreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		if stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_INTEGER:
		if stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_REAL:
		if stage.OnAfterA_ATTRIBUTE_VALUE_REALCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_REALCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_STRING:
		if stage.OnAfterA_ATTRIBUTE_VALUE_STRINGCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_STRINGCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_XHTML:
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLCreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1CreateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1CreateCallback.OnAfterCreate(stage, target)
		}
	case *A_CHILDREN:
		if stage.OnAfterA_CHILDRENCreateCallback != nil {
			stage.OnAfterA_CHILDRENCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_CORE_CONTENT:
		if stage.OnAfterA_CORE_CONTENTCreateCallback != nil {
			stage.OnAfterA_CORE_CONTENTCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_DATATYPES:
		if stage.OnAfterA_DATATYPESCreateCallback != nil {
			stage.OnAfterA_DATATYPESCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFCreateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_DATATYPE_DEFINITION_DATE_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFCreateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFCreateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFCreateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_DATATYPE_DEFINITION_REAL_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFCreateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_DATATYPE_DEFINITION_STRING_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFCreateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFCreateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_EDITABLE_ATTS:
		if stage.OnAfterA_EDITABLE_ATTSCreateCallback != nil {
			stage.OnAfterA_EDITABLE_ATTSCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_ENUM_VALUE_REF:
		if stage.OnAfterA_ENUM_VALUE_REFCreateCallback != nil {
			stage.OnAfterA_ENUM_VALUE_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_OBJECT:
		if stage.OnAfterA_OBJECTCreateCallback != nil {
			stage.OnAfterA_OBJECTCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_PROPERTIES:
		if stage.OnAfterA_PROPERTIESCreateCallback != nil {
			stage.OnAfterA_PROPERTIESCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_RELATION_GROUP_TYPE_REF:
		if stage.OnAfterA_RELATION_GROUP_TYPE_REFCreateCallback != nil {
			stage.OnAfterA_RELATION_GROUP_TYPE_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SOURCE_1:
		if stage.OnAfterA_SOURCE_1CreateCallback != nil {
			stage.OnAfterA_SOURCE_1CreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SOURCE_SPECIFICATION_1:
		if stage.OnAfterA_SOURCE_SPECIFICATION_1CreateCallback != nil {
			stage.OnAfterA_SOURCE_SPECIFICATION_1CreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPECIFICATIONS:
		if stage.OnAfterA_SPECIFICATIONSCreateCallback != nil {
			stage.OnAfterA_SPECIFICATIONSCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPECIFICATION_TYPE_REF:
		if stage.OnAfterA_SPECIFICATION_TYPE_REFCreateCallback != nil {
			stage.OnAfterA_SPECIFICATION_TYPE_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPECIFIED_VALUES:
		if stage.OnAfterA_SPECIFIED_VALUESCreateCallback != nil {
			stage.OnAfterA_SPECIFIED_VALUESCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPEC_ATTRIBUTES:
		if stage.OnAfterA_SPEC_ATTRIBUTESCreateCallback != nil {
			stage.OnAfterA_SPEC_ATTRIBUTESCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPEC_OBJECTS:
		if stage.OnAfterA_SPEC_OBJECTSCreateCallback != nil {
			stage.OnAfterA_SPEC_OBJECTSCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPEC_OBJECT_TYPE_REF:
		if stage.OnAfterA_SPEC_OBJECT_TYPE_REFCreateCallback != nil {
			stage.OnAfterA_SPEC_OBJECT_TYPE_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPEC_RELATIONS:
		if stage.OnAfterA_SPEC_RELATIONSCreateCallback != nil {
			stage.OnAfterA_SPEC_RELATIONSCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPEC_RELATION_GROUPS:
		if stage.OnAfterA_SPEC_RELATION_GROUPSCreateCallback != nil {
			stage.OnAfterA_SPEC_RELATION_GROUPSCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPEC_RELATION_REF:
		if stage.OnAfterA_SPEC_RELATION_REFCreateCallback != nil {
			stage.OnAfterA_SPEC_RELATION_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPEC_RELATION_TYPE_REF:
		if stage.OnAfterA_SPEC_RELATION_TYPE_REFCreateCallback != nil {
			stage.OnAfterA_SPEC_RELATION_TYPE_REFCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_SPEC_TYPES:
		if stage.OnAfterA_SPEC_TYPESCreateCallback != nil {
			stage.OnAfterA_SPEC_TYPESCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_THE_HEADER:
		if stage.OnAfterA_THE_HEADERCreateCallback != nil {
			stage.OnAfterA_THE_HEADERCreateCallback.OnAfterCreate(stage, target)
		}
	case *A_TOOL_EXTENSIONS:
		if stage.OnAfterA_TOOL_EXTENSIONSCreateCallback != nil {
			stage.OnAfterA_TOOL_EXTENSIONSCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPE_DEFINITION_BOOLEAN:
		if stage.OnAfterDATATYPE_DEFINITION_BOOLEANCreateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_BOOLEANCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPE_DEFINITION_DATE:
		if stage.OnAfterDATATYPE_DEFINITION_DATECreateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_DATECreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPE_DEFINITION_ENUMERATION:
		if stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONCreateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPE_DEFINITION_INTEGER:
		if stage.OnAfterDATATYPE_DEFINITION_INTEGERCreateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_INTEGERCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPE_DEFINITION_REAL:
		if stage.OnAfterDATATYPE_DEFINITION_REALCreateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_REALCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPE_DEFINITION_STRING:
		if stage.OnAfterDATATYPE_DEFINITION_STRINGCreateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_STRINGCreateCallback.OnAfterCreate(stage, target)
		}
	case *DATATYPE_DEFINITION_XHTML:
		if stage.OnAfterDATATYPE_DEFINITION_XHTMLCreateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_XHTMLCreateCallback.OnAfterCreate(stage, target)
		}
	case *EMBEDDED_VALUE:
		if stage.OnAfterEMBEDDED_VALUECreateCallback != nil {
			stage.OnAfterEMBEDDED_VALUECreateCallback.OnAfterCreate(stage, target)
		}
	case *ENUM_VALUE:
		if stage.OnAfterENUM_VALUECreateCallback != nil {
			stage.OnAfterENUM_VALUECreateCallback.OnAfterCreate(stage, target)
		}
	case *RELATION_GROUP:
		if stage.OnAfterRELATION_GROUPCreateCallback != nil {
			stage.OnAfterRELATION_GROUPCreateCallback.OnAfterCreate(stage, target)
		}
	case *RELATION_GROUP_TYPE:
		if stage.OnAfterRELATION_GROUP_TYPECreateCallback != nil {
			stage.OnAfterRELATION_GROUP_TYPECreateCallback.OnAfterCreate(stage, target)
		}
	case *REQ_IF:
		if stage.OnAfterREQ_IFCreateCallback != nil {
			stage.OnAfterREQ_IFCreateCallback.OnAfterCreate(stage, target)
		}
	case *REQ_IF_CONTENT:
		if stage.OnAfterREQ_IF_CONTENTCreateCallback != nil {
			stage.OnAfterREQ_IF_CONTENTCreateCallback.OnAfterCreate(stage, target)
		}
	case *REQ_IF_HEADER:
		if stage.OnAfterREQ_IF_HEADERCreateCallback != nil {
			stage.OnAfterREQ_IF_HEADERCreateCallback.OnAfterCreate(stage, target)
		}
	case *REQ_IF_TOOL_EXTENSION:
		if stage.OnAfterREQ_IF_TOOL_EXTENSIONCreateCallback != nil {
			stage.OnAfterREQ_IF_TOOL_EXTENSIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECIFICATION:
		if stage.OnAfterSPECIFICATIONCreateCallback != nil {
			stage.OnAfterSPECIFICATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPECIFICATION_TYPE:
		if stage.OnAfterSPECIFICATION_TYPECreateCallback != nil {
			stage.OnAfterSPECIFICATION_TYPECreateCallback.OnAfterCreate(stage, target)
		}
	case *SPEC_HIERARCHY:
		if stage.OnAfterSPEC_HIERARCHYCreateCallback != nil {
			stage.OnAfterSPEC_HIERARCHYCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPEC_OBJECT:
		if stage.OnAfterSPEC_OBJECTCreateCallback != nil {
			stage.OnAfterSPEC_OBJECTCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPEC_OBJECT_TYPE:
		if stage.OnAfterSPEC_OBJECT_TYPECreateCallback != nil {
			stage.OnAfterSPEC_OBJECT_TYPECreateCallback.OnAfterCreate(stage, target)
		}
	case *SPEC_RELATION:
		if stage.OnAfterSPEC_RELATIONCreateCallback != nil {
			stage.OnAfterSPEC_RELATIONCreateCallback.OnAfterCreate(stage, target)
		}
	case *SPEC_RELATION_TYPE:
		if stage.OnAfterSPEC_RELATION_TYPECreateCallback != nil {
			stage.OnAfterSPEC_RELATION_TYPECreateCallback.OnAfterCreate(stage, target)
		}
	case *XHTML_CONTENT:
		if stage.OnAfterXHTML_CONTENTCreateCallback != nil {
			stage.OnAfterXHTML_CONTENTCreateCallback.OnAfterCreate(stage, target)
		}
	default:
		_ = target
	}
}

type Gong__MouseEvent struct {
	ShiftKey bool
}

// OnAfterUpdateFromFront is called after a update from front
func OnAfterUpdateFromFront[Type Gongstruct](stage *Stage, old, new *Type) {

	switch oldTarget := any(old).(type) {
	// insertion point
	case *ALTERNATIVE_ID:
		newTarget := any(new).(*ALTERNATIVE_ID)
		if stage.OnAfterALTERNATIVE_IDUpdateCallback != nil {
			stage.OnAfterALTERNATIVE_IDUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_BOOLEAN)
		if stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_DEFINITION_DATE:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_DATE)
		if stage.OnAfterATTRIBUTE_DEFINITION_DATEUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_DATEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_ENUMERATION)
		if stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_DEFINITION_INTEGER:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_INTEGER)
		if stage.OnAfterATTRIBUTE_DEFINITION_INTEGERUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_INTEGERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_DEFINITION_REAL:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_REAL)
		if stage.OnAfterATTRIBUTE_DEFINITION_REALUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_REALUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_DEFINITION_STRING:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_STRING)
		if stage.OnAfterATTRIBUTE_DEFINITION_STRINGUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_STRINGUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_DEFINITION_XHTML:
		newTarget := any(new).(*ATTRIBUTE_DEFINITION_XHTML)
		if stage.OnAfterATTRIBUTE_DEFINITION_XHTMLUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_XHTMLUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_VALUE_BOOLEAN:
		newTarget := any(new).(*ATTRIBUTE_VALUE_BOOLEAN)
		if stage.OnAfterATTRIBUTE_VALUE_BOOLEANUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_BOOLEANUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_VALUE_DATE:
		newTarget := any(new).(*ATTRIBUTE_VALUE_DATE)
		if stage.OnAfterATTRIBUTE_VALUE_DATEUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_DATEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_VALUE_ENUMERATION:
		newTarget := any(new).(*ATTRIBUTE_VALUE_ENUMERATION)
		if stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_VALUE_INTEGER:
		newTarget := any(new).(*ATTRIBUTE_VALUE_INTEGER)
		if stage.OnAfterATTRIBUTE_VALUE_INTEGERUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_INTEGERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_VALUE_REAL:
		newTarget := any(new).(*ATTRIBUTE_VALUE_REAL)
		if stage.OnAfterATTRIBUTE_VALUE_REALUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_REALUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_VALUE_STRING:
		newTarget := any(new).(*ATTRIBUTE_VALUE_STRING)
		if stage.OnAfterATTRIBUTE_VALUE_STRINGUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_STRINGUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ATTRIBUTE_VALUE_XHTML:
		newTarget := any(new).(*ATTRIBUTE_VALUE_XHTML)
		if stage.OnAfterATTRIBUTE_VALUE_XHTMLUpdateCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_XHTMLUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ALTERNATIVE_ID:
		newTarget := any(new).(*A_ALTERNATIVE_ID)
		if stage.OnAfterA_ALTERNATIVE_IDUpdateCallback != nil {
			stage.OnAfterA_ALTERNATIVE_IDUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		newTarget := any(new).(*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFUpdateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		newTarget := any(new).(*A_ATTRIBUTE_DEFINITION_DATE_REF)
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFUpdateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		newTarget := any(new).(*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFUpdateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		newTarget := any(new).(*A_ATTRIBUTE_DEFINITION_INTEGER_REF)
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFUpdateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		newTarget := any(new).(*A_ATTRIBUTE_DEFINITION_REAL_REF)
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFUpdateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		newTarget := any(new).(*A_ATTRIBUTE_DEFINITION_STRING_REF)
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFUpdateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		newTarget := any(new).(*A_ATTRIBUTE_DEFINITION_XHTML_REF)
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFUpdateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		newTarget := any(new).(*A_ATTRIBUTE_VALUE_BOOLEAN)
		if stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANUpdateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ATTRIBUTE_VALUE_DATE:
		newTarget := any(new).(*A_ATTRIBUTE_VALUE_DATE)
		if stage.OnAfterA_ATTRIBUTE_VALUE_DATEUpdateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_DATEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		newTarget := any(new).(*A_ATTRIBUTE_VALUE_ENUMERATION)
		if stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONUpdateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ATTRIBUTE_VALUE_INTEGER:
		newTarget := any(new).(*A_ATTRIBUTE_VALUE_INTEGER)
		if stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERUpdateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ATTRIBUTE_VALUE_REAL:
		newTarget := any(new).(*A_ATTRIBUTE_VALUE_REAL)
		if stage.OnAfterA_ATTRIBUTE_VALUE_REALUpdateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_REALUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ATTRIBUTE_VALUE_STRING:
		newTarget := any(new).(*A_ATTRIBUTE_VALUE_STRING)
		if stage.OnAfterA_ATTRIBUTE_VALUE_STRINGUpdateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_STRINGUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ATTRIBUTE_VALUE_XHTML:
		newTarget := any(new).(*A_ATTRIBUTE_VALUE_XHTML)
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLUpdateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		newTarget := any(new).(*A_ATTRIBUTE_VALUE_XHTML_1)
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1UpdateCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_CHILDREN:
		newTarget := any(new).(*A_CHILDREN)
		if stage.OnAfterA_CHILDRENUpdateCallback != nil {
			stage.OnAfterA_CHILDRENUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_CORE_CONTENT:
		newTarget := any(new).(*A_CORE_CONTENT)
		if stage.OnAfterA_CORE_CONTENTUpdateCallback != nil {
			stage.OnAfterA_CORE_CONTENTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_DATATYPES:
		newTarget := any(new).(*A_DATATYPES)
		if stage.OnAfterA_DATATYPESUpdateCallback != nil {
			stage.OnAfterA_DATATYPESUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		newTarget := any(new).(*A_DATATYPE_DEFINITION_BOOLEAN_REF)
		if stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFUpdateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_DATATYPE_DEFINITION_DATE_REF:
		newTarget := any(new).(*A_DATATYPE_DEFINITION_DATE_REF)
		if stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFUpdateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		newTarget := any(new).(*A_DATATYPE_DEFINITION_ENUMERATION_REF)
		if stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFUpdateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		newTarget := any(new).(*A_DATATYPE_DEFINITION_INTEGER_REF)
		if stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFUpdateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_DATATYPE_DEFINITION_REAL_REF:
		newTarget := any(new).(*A_DATATYPE_DEFINITION_REAL_REF)
		if stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFUpdateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_DATATYPE_DEFINITION_STRING_REF:
		newTarget := any(new).(*A_DATATYPE_DEFINITION_STRING_REF)
		if stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFUpdateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		newTarget := any(new).(*A_DATATYPE_DEFINITION_XHTML_REF)
		if stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFUpdateCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_EDITABLE_ATTS:
		newTarget := any(new).(*A_EDITABLE_ATTS)
		if stage.OnAfterA_EDITABLE_ATTSUpdateCallback != nil {
			stage.OnAfterA_EDITABLE_ATTSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_ENUM_VALUE_REF:
		newTarget := any(new).(*A_ENUM_VALUE_REF)
		if stage.OnAfterA_ENUM_VALUE_REFUpdateCallback != nil {
			stage.OnAfterA_ENUM_VALUE_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_OBJECT:
		newTarget := any(new).(*A_OBJECT)
		if stage.OnAfterA_OBJECTUpdateCallback != nil {
			stage.OnAfterA_OBJECTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_PROPERTIES:
		newTarget := any(new).(*A_PROPERTIES)
		if stage.OnAfterA_PROPERTIESUpdateCallback != nil {
			stage.OnAfterA_PROPERTIESUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_RELATION_GROUP_TYPE_REF:
		newTarget := any(new).(*A_RELATION_GROUP_TYPE_REF)
		if stage.OnAfterA_RELATION_GROUP_TYPE_REFUpdateCallback != nil {
			stage.OnAfterA_RELATION_GROUP_TYPE_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_SOURCE_1:
		newTarget := any(new).(*A_SOURCE_1)
		if stage.OnAfterA_SOURCE_1UpdateCallback != nil {
			stage.OnAfterA_SOURCE_1UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_SOURCE_SPECIFICATION_1:
		newTarget := any(new).(*A_SOURCE_SPECIFICATION_1)
		if stage.OnAfterA_SOURCE_SPECIFICATION_1UpdateCallback != nil {
			stage.OnAfterA_SOURCE_SPECIFICATION_1UpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_SPECIFICATIONS:
		newTarget := any(new).(*A_SPECIFICATIONS)
		if stage.OnAfterA_SPECIFICATIONSUpdateCallback != nil {
			stage.OnAfterA_SPECIFICATIONSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_SPECIFICATION_TYPE_REF:
		newTarget := any(new).(*A_SPECIFICATION_TYPE_REF)
		if stage.OnAfterA_SPECIFICATION_TYPE_REFUpdateCallback != nil {
			stage.OnAfterA_SPECIFICATION_TYPE_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_SPECIFIED_VALUES:
		newTarget := any(new).(*A_SPECIFIED_VALUES)
		if stage.OnAfterA_SPECIFIED_VALUESUpdateCallback != nil {
			stage.OnAfterA_SPECIFIED_VALUESUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_SPEC_ATTRIBUTES:
		newTarget := any(new).(*A_SPEC_ATTRIBUTES)
		if stage.OnAfterA_SPEC_ATTRIBUTESUpdateCallback != nil {
			stage.OnAfterA_SPEC_ATTRIBUTESUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_SPEC_OBJECTS:
		newTarget := any(new).(*A_SPEC_OBJECTS)
		if stage.OnAfterA_SPEC_OBJECTSUpdateCallback != nil {
			stage.OnAfterA_SPEC_OBJECTSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_SPEC_OBJECT_TYPE_REF:
		newTarget := any(new).(*A_SPEC_OBJECT_TYPE_REF)
		if stage.OnAfterA_SPEC_OBJECT_TYPE_REFUpdateCallback != nil {
			stage.OnAfterA_SPEC_OBJECT_TYPE_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_SPEC_RELATIONS:
		newTarget := any(new).(*A_SPEC_RELATIONS)
		if stage.OnAfterA_SPEC_RELATIONSUpdateCallback != nil {
			stage.OnAfterA_SPEC_RELATIONSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_SPEC_RELATION_GROUPS:
		newTarget := any(new).(*A_SPEC_RELATION_GROUPS)
		if stage.OnAfterA_SPEC_RELATION_GROUPSUpdateCallback != nil {
			stage.OnAfterA_SPEC_RELATION_GROUPSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_SPEC_RELATION_REF:
		newTarget := any(new).(*A_SPEC_RELATION_REF)
		if stage.OnAfterA_SPEC_RELATION_REFUpdateCallback != nil {
			stage.OnAfterA_SPEC_RELATION_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_SPEC_RELATION_TYPE_REF:
		newTarget := any(new).(*A_SPEC_RELATION_TYPE_REF)
		if stage.OnAfterA_SPEC_RELATION_TYPE_REFUpdateCallback != nil {
			stage.OnAfterA_SPEC_RELATION_TYPE_REFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_SPEC_TYPES:
		newTarget := any(new).(*A_SPEC_TYPES)
		if stage.OnAfterA_SPEC_TYPESUpdateCallback != nil {
			stage.OnAfterA_SPEC_TYPESUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_THE_HEADER:
		newTarget := any(new).(*A_THE_HEADER)
		if stage.OnAfterA_THE_HEADERUpdateCallback != nil {
			stage.OnAfterA_THE_HEADERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *A_TOOL_EXTENSIONS:
		newTarget := any(new).(*A_TOOL_EXTENSIONS)
		if stage.OnAfterA_TOOL_EXTENSIONSUpdateCallback != nil {
			stage.OnAfterA_TOOL_EXTENSIONSUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPE_DEFINITION_BOOLEAN:
		newTarget := any(new).(*DATATYPE_DEFINITION_BOOLEAN)
		if stage.OnAfterDATATYPE_DEFINITION_BOOLEANUpdateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_BOOLEANUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPE_DEFINITION_DATE:
		newTarget := any(new).(*DATATYPE_DEFINITION_DATE)
		if stage.OnAfterDATATYPE_DEFINITION_DATEUpdateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_DATEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPE_DEFINITION_ENUMERATION:
		newTarget := any(new).(*DATATYPE_DEFINITION_ENUMERATION)
		if stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONUpdateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPE_DEFINITION_INTEGER:
		newTarget := any(new).(*DATATYPE_DEFINITION_INTEGER)
		if stage.OnAfterDATATYPE_DEFINITION_INTEGERUpdateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_INTEGERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPE_DEFINITION_REAL:
		newTarget := any(new).(*DATATYPE_DEFINITION_REAL)
		if stage.OnAfterDATATYPE_DEFINITION_REALUpdateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_REALUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPE_DEFINITION_STRING:
		newTarget := any(new).(*DATATYPE_DEFINITION_STRING)
		if stage.OnAfterDATATYPE_DEFINITION_STRINGUpdateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_STRINGUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *DATATYPE_DEFINITION_XHTML:
		newTarget := any(new).(*DATATYPE_DEFINITION_XHTML)
		if stage.OnAfterDATATYPE_DEFINITION_XHTMLUpdateCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_XHTMLUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *EMBEDDED_VALUE:
		newTarget := any(new).(*EMBEDDED_VALUE)
		if stage.OnAfterEMBEDDED_VALUEUpdateCallback != nil {
			stage.OnAfterEMBEDDED_VALUEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *ENUM_VALUE:
		newTarget := any(new).(*ENUM_VALUE)
		if stage.OnAfterENUM_VALUEUpdateCallback != nil {
			stage.OnAfterENUM_VALUEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RELATION_GROUP:
		newTarget := any(new).(*RELATION_GROUP)
		if stage.OnAfterRELATION_GROUPUpdateCallback != nil {
			stage.OnAfterRELATION_GROUPUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *RELATION_GROUP_TYPE:
		newTarget := any(new).(*RELATION_GROUP_TYPE)
		if stage.OnAfterRELATION_GROUP_TYPEUpdateCallback != nil {
			stage.OnAfterRELATION_GROUP_TYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *REQ_IF:
		newTarget := any(new).(*REQ_IF)
		if stage.OnAfterREQ_IFUpdateCallback != nil {
			stage.OnAfterREQ_IFUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *REQ_IF_CONTENT:
		newTarget := any(new).(*REQ_IF_CONTENT)
		if stage.OnAfterREQ_IF_CONTENTUpdateCallback != nil {
			stage.OnAfterREQ_IF_CONTENTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *REQ_IF_HEADER:
		newTarget := any(new).(*REQ_IF_HEADER)
		if stage.OnAfterREQ_IF_HEADERUpdateCallback != nil {
			stage.OnAfterREQ_IF_HEADERUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *REQ_IF_TOOL_EXTENSION:
		newTarget := any(new).(*REQ_IF_TOOL_EXTENSION)
		if stage.OnAfterREQ_IF_TOOL_EXTENSIONUpdateCallback != nil {
			stage.OnAfterREQ_IF_TOOL_EXTENSIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECIFICATION:
		newTarget := any(new).(*SPECIFICATION)
		if stage.OnAfterSPECIFICATIONUpdateCallback != nil {
			stage.OnAfterSPECIFICATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPECIFICATION_TYPE:
		newTarget := any(new).(*SPECIFICATION_TYPE)
		if stage.OnAfterSPECIFICATION_TYPEUpdateCallback != nil {
			stage.OnAfterSPECIFICATION_TYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPEC_HIERARCHY:
		newTarget := any(new).(*SPEC_HIERARCHY)
		if stage.OnAfterSPEC_HIERARCHYUpdateCallback != nil {
			stage.OnAfterSPEC_HIERARCHYUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPEC_OBJECT:
		newTarget := any(new).(*SPEC_OBJECT)
		if stage.OnAfterSPEC_OBJECTUpdateCallback != nil {
			stage.OnAfterSPEC_OBJECTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPEC_OBJECT_TYPE:
		newTarget := any(new).(*SPEC_OBJECT_TYPE)
		if stage.OnAfterSPEC_OBJECT_TYPEUpdateCallback != nil {
			stage.OnAfterSPEC_OBJECT_TYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPEC_RELATION:
		newTarget := any(new).(*SPEC_RELATION)
		if stage.OnAfterSPEC_RELATIONUpdateCallback != nil {
			stage.OnAfterSPEC_RELATIONUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *SPEC_RELATION_TYPE:
		newTarget := any(new).(*SPEC_RELATION_TYPE)
		if stage.OnAfterSPEC_RELATION_TYPEUpdateCallback != nil {
			stage.OnAfterSPEC_RELATION_TYPEUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	case *XHTML_CONTENT:
		newTarget := any(new).(*XHTML_CONTENT)
		if stage.OnAfterXHTML_CONTENTUpdateCallback != nil {
			stage.OnAfterXHTML_CONTENTUpdateCallback.OnAfterUpdate(stage, oldTarget, newTarget)
		}
	default:
		_ = oldTarget
	}
}

// AfterDeleteFromFront is called after a delete from front
func AfterDeleteFromFront[Type Gongstruct](stage *Stage, staged, front *Type) {

	switch front := any(front).(type) {
	// insertion point
	case *ALTERNATIVE_ID:
		if stage.OnAfterALTERNATIVE_IDDeleteCallback != nil {
			staged := any(staged).(*ALTERNATIVE_ID)
			stage.OnAfterALTERNATIVE_IDDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		if stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_DEFINITION_BOOLEAN)
			stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_DEFINITION_DATE:
		if stage.OnAfterATTRIBUTE_DEFINITION_DATEDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_DEFINITION_DATE)
			stage.OnAfterATTRIBUTE_DEFINITION_DATEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		if stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_DEFINITION_ENUMERATION)
			stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_DEFINITION_INTEGER:
		if stage.OnAfterATTRIBUTE_DEFINITION_INTEGERDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_DEFINITION_INTEGER)
			stage.OnAfterATTRIBUTE_DEFINITION_INTEGERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_DEFINITION_REAL:
		if stage.OnAfterATTRIBUTE_DEFINITION_REALDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_DEFINITION_REAL)
			stage.OnAfterATTRIBUTE_DEFINITION_REALDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_DEFINITION_STRING:
		if stage.OnAfterATTRIBUTE_DEFINITION_STRINGDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_DEFINITION_STRING)
			stage.OnAfterATTRIBUTE_DEFINITION_STRINGDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_DEFINITION_XHTML:
		if stage.OnAfterATTRIBUTE_DEFINITION_XHTMLDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_DEFINITION_XHTML)
			stage.OnAfterATTRIBUTE_DEFINITION_XHTMLDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_VALUE_BOOLEAN:
		if stage.OnAfterATTRIBUTE_VALUE_BOOLEANDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_VALUE_BOOLEAN)
			stage.OnAfterATTRIBUTE_VALUE_BOOLEANDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_VALUE_DATE:
		if stage.OnAfterATTRIBUTE_VALUE_DATEDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_VALUE_DATE)
			stage.OnAfterATTRIBUTE_VALUE_DATEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_VALUE_ENUMERATION:
		if stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_VALUE_ENUMERATION)
			stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_VALUE_INTEGER:
		if stage.OnAfterATTRIBUTE_VALUE_INTEGERDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_VALUE_INTEGER)
			stage.OnAfterATTRIBUTE_VALUE_INTEGERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_VALUE_REAL:
		if stage.OnAfterATTRIBUTE_VALUE_REALDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_VALUE_REAL)
			stage.OnAfterATTRIBUTE_VALUE_REALDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_VALUE_STRING:
		if stage.OnAfterATTRIBUTE_VALUE_STRINGDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_VALUE_STRING)
			stage.OnAfterATTRIBUTE_VALUE_STRINGDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ATTRIBUTE_VALUE_XHTML:
		if stage.OnAfterATTRIBUTE_VALUE_XHTMLDeleteCallback != nil {
			staged := any(staged).(*ATTRIBUTE_VALUE_XHTML)
			stage.OnAfterATTRIBUTE_VALUE_XHTMLDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ALTERNATIVE_ID:
		if stage.OnAfterA_ALTERNATIVE_IDDeleteCallback != nil {
			staged := any(staged).(*A_ALTERNATIVE_ID)
			stage.OnAfterA_ALTERNATIVE_IDDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)
			stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_DEFINITION_DATE_REF)
			stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)
			stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_DEFINITION_INTEGER_REF)
			stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_DEFINITION_REAL_REF)
			stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_DEFINITION_STRING_REF)
			stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_DEFINITION_XHTML_REF)
			stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		if stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_VALUE_BOOLEAN)
			stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_VALUE_DATE:
		if stage.OnAfterA_ATTRIBUTE_VALUE_DATEDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_VALUE_DATE)
			stage.OnAfterA_ATTRIBUTE_VALUE_DATEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		if stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_VALUE_ENUMERATION)
			stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_VALUE_INTEGER:
		if stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_VALUE_INTEGER)
			stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_VALUE_REAL:
		if stage.OnAfterA_ATTRIBUTE_VALUE_REALDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_VALUE_REAL)
			stage.OnAfterA_ATTRIBUTE_VALUE_REALDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_VALUE_STRING:
		if stage.OnAfterA_ATTRIBUTE_VALUE_STRINGDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_VALUE_STRING)
			stage.OnAfterA_ATTRIBUTE_VALUE_STRINGDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_VALUE_XHTML:
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLDeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_VALUE_XHTML)
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1DeleteCallback != nil {
			staged := any(staged).(*A_ATTRIBUTE_VALUE_XHTML_1)
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_CHILDREN:
		if stage.OnAfterA_CHILDRENDeleteCallback != nil {
			staged := any(staged).(*A_CHILDREN)
			stage.OnAfterA_CHILDRENDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_CORE_CONTENT:
		if stage.OnAfterA_CORE_CONTENTDeleteCallback != nil {
			staged := any(staged).(*A_CORE_CONTENT)
			stage.OnAfterA_CORE_CONTENTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_DATATYPES:
		if stage.OnAfterA_DATATYPESDeleteCallback != nil {
			staged := any(staged).(*A_DATATYPES)
			stage.OnAfterA_DATATYPESDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFDeleteCallback != nil {
			staged := any(staged).(*A_DATATYPE_DEFINITION_BOOLEAN_REF)
			stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_DATATYPE_DEFINITION_DATE_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFDeleteCallback != nil {
			staged := any(staged).(*A_DATATYPE_DEFINITION_DATE_REF)
			stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFDeleteCallback != nil {
			staged := any(staged).(*A_DATATYPE_DEFINITION_ENUMERATION_REF)
			stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFDeleteCallback != nil {
			staged := any(staged).(*A_DATATYPE_DEFINITION_INTEGER_REF)
			stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_DATATYPE_DEFINITION_REAL_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFDeleteCallback != nil {
			staged := any(staged).(*A_DATATYPE_DEFINITION_REAL_REF)
			stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_DATATYPE_DEFINITION_STRING_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFDeleteCallback != nil {
			staged := any(staged).(*A_DATATYPE_DEFINITION_STRING_REF)
			stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFDeleteCallback != nil {
			staged := any(staged).(*A_DATATYPE_DEFINITION_XHTML_REF)
			stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_EDITABLE_ATTS:
		if stage.OnAfterA_EDITABLE_ATTSDeleteCallback != nil {
			staged := any(staged).(*A_EDITABLE_ATTS)
			stage.OnAfterA_EDITABLE_ATTSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_ENUM_VALUE_REF:
		if stage.OnAfterA_ENUM_VALUE_REFDeleteCallback != nil {
			staged := any(staged).(*A_ENUM_VALUE_REF)
			stage.OnAfterA_ENUM_VALUE_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_OBJECT:
		if stage.OnAfterA_OBJECTDeleteCallback != nil {
			staged := any(staged).(*A_OBJECT)
			stage.OnAfterA_OBJECTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_PROPERTIES:
		if stage.OnAfterA_PROPERTIESDeleteCallback != nil {
			staged := any(staged).(*A_PROPERTIES)
			stage.OnAfterA_PROPERTIESDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_RELATION_GROUP_TYPE_REF:
		if stage.OnAfterA_RELATION_GROUP_TYPE_REFDeleteCallback != nil {
			staged := any(staged).(*A_RELATION_GROUP_TYPE_REF)
			stage.OnAfterA_RELATION_GROUP_TYPE_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SOURCE_1:
		if stage.OnAfterA_SOURCE_1DeleteCallback != nil {
			staged := any(staged).(*A_SOURCE_1)
			stage.OnAfterA_SOURCE_1DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SOURCE_SPECIFICATION_1:
		if stage.OnAfterA_SOURCE_SPECIFICATION_1DeleteCallback != nil {
			staged := any(staged).(*A_SOURCE_SPECIFICATION_1)
			stage.OnAfterA_SOURCE_SPECIFICATION_1DeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPECIFICATIONS:
		if stage.OnAfterA_SPECIFICATIONSDeleteCallback != nil {
			staged := any(staged).(*A_SPECIFICATIONS)
			stage.OnAfterA_SPECIFICATIONSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPECIFICATION_TYPE_REF:
		if stage.OnAfterA_SPECIFICATION_TYPE_REFDeleteCallback != nil {
			staged := any(staged).(*A_SPECIFICATION_TYPE_REF)
			stage.OnAfterA_SPECIFICATION_TYPE_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPECIFIED_VALUES:
		if stage.OnAfterA_SPECIFIED_VALUESDeleteCallback != nil {
			staged := any(staged).(*A_SPECIFIED_VALUES)
			stage.OnAfterA_SPECIFIED_VALUESDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPEC_ATTRIBUTES:
		if stage.OnAfterA_SPEC_ATTRIBUTESDeleteCallback != nil {
			staged := any(staged).(*A_SPEC_ATTRIBUTES)
			stage.OnAfterA_SPEC_ATTRIBUTESDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPEC_OBJECTS:
		if stage.OnAfterA_SPEC_OBJECTSDeleteCallback != nil {
			staged := any(staged).(*A_SPEC_OBJECTS)
			stage.OnAfterA_SPEC_OBJECTSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPEC_OBJECT_TYPE_REF:
		if stage.OnAfterA_SPEC_OBJECT_TYPE_REFDeleteCallback != nil {
			staged := any(staged).(*A_SPEC_OBJECT_TYPE_REF)
			stage.OnAfterA_SPEC_OBJECT_TYPE_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPEC_RELATIONS:
		if stage.OnAfterA_SPEC_RELATIONSDeleteCallback != nil {
			staged := any(staged).(*A_SPEC_RELATIONS)
			stage.OnAfterA_SPEC_RELATIONSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPEC_RELATION_GROUPS:
		if stage.OnAfterA_SPEC_RELATION_GROUPSDeleteCallback != nil {
			staged := any(staged).(*A_SPEC_RELATION_GROUPS)
			stage.OnAfterA_SPEC_RELATION_GROUPSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPEC_RELATION_REF:
		if stage.OnAfterA_SPEC_RELATION_REFDeleteCallback != nil {
			staged := any(staged).(*A_SPEC_RELATION_REF)
			stage.OnAfterA_SPEC_RELATION_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPEC_RELATION_TYPE_REF:
		if stage.OnAfterA_SPEC_RELATION_TYPE_REFDeleteCallback != nil {
			staged := any(staged).(*A_SPEC_RELATION_TYPE_REF)
			stage.OnAfterA_SPEC_RELATION_TYPE_REFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_SPEC_TYPES:
		if stage.OnAfterA_SPEC_TYPESDeleteCallback != nil {
			staged := any(staged).(*A_SPEC_TYPES)
			stage.OnAfterA_SPEC_TYPESDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_THE_HEADER:
		if stage.OnAfterA_THE_HEADERDeleteCallback != nil {
			staged := any(staged).(*A_THE_HEADER)
			stage.OnAfterA_THE_HEADERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *A_TOOL_EXTENSIONS:
		if stage.OnAfterA_TOOL_EXTENSIONSDeleteCallback != nil {
			staged := any(staged).(*A_TOOL_EXTENSIONS)
			stage.OnAfterA_TOOL_EXTENSIONSDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPE_DEFINITION_BOOLEAN:
		if stage.OnAfterDATATYPE_DEFINITION_BOOLEANDeleteCallback != nil {
			staged := any(staged).(*DATATYPE_DEFINITION_BOOLEAN)
			stage.OnAfterDATATYPE_DEFINITION_BOOLEANDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPE_DEFINITION_DATE:
		if stage.OnAfterDATATYPE_DEFINITION_DATEDeleteCallback != nil {
			staged := any(staged).(*DATATYPE_DEFINITION_DATE)
			stage.OnAfterDATATYPE_DEFINITION_DATEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPE_DEFINITION_ENUMERATION:
		if stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONDeleteCallback != nil {
			staged := any(staged).(*DATATYPE_DEFINITION_ENUMERATION)
			stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPE_DEFINITION_INTEGER:
		if stage.OnAfterDATATYPE_DEFINITION_INTEGERDeleteCallback != nil {
			staged := any(staged).(*DATATYPE_DEFINITION_INTEGER)
			stage.OnAfterDATATYPE_DEFINITION_INTEGERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPE_DEFINITION_REAL:
		if stage.OnAfterDATATYPE_DEFINITION_REALDeleteCallback != nil {
			staged := any(staged).(*DATATYPE_DEFINITION_REAL)
			stage.OnAfterDATATYPE_DEFINITION_REALDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPE_DEFINITION_STRING:
		if stage.OnAfterDATATYPE_DEFINITION_STRINGDeleteCallback != nil {
			staged := any(staged).(*DATATYPE_DEFINITION_STRING)
			stage.OnAfterDATATYPE_DEFINITION_STRINGDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *DATATYPE_DEFINITION_XHTML:
		if stage.OnAfterDATATYPE_DEFINITION_XHTMLDeleteCallback != nil {
			staged := any(staged).(*DATATYPE_DEFINITION_XHTML)
			stage.OnAfterDATATYPE_DEFINITION_XHTMLDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *EMBEDDED_VALUE:
		if stage.OnAfterEMBEDDED_VALUEDeleteCallback != nil {
			staged := any(staged).(*EMBEDDED_VALUE)
			stage.OnAfterEMBEDDED_VALUEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *ENUM_VALUE:
		if stage.OnAfterENUM_VALUEDeleteCallback != nil {
			staged := any(staged).(*ENUM_VALUE)
			stage.OnAfterENUM_VALUEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RELATION_GROUP:
		if stage.OnAfterRELATION_GROUPDeleteCallback != nil {
			staged := any(staged).(*RELATION_GROUP)
			stage.OnAfterRELATION_GROUPDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *RELATION_GROUP_TYPE:
		if stage.OnAfterRELATION_GROUP_TYPEDeleteCallback != nil {
			staged := any(staged).(*RELATION_GROUP_TYPE)
			stage.OnAfterRELATION_GROUP_TYPEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *REQ_IF:
		if stage.OnAfterREQ_IFDeleteCallback != nil {
			staged := any(staged).(*REQ_IF)
			stage.OnAfterREQ_IFDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *REQ_IF_CONTENT:
		if stage.OnAfterREQ_IF_CONTENTDeleteCallback != nil {
			staged := any(staged).(*REQ_IF_CONTENT)
			stage.OnAfterREQ_IF_CONTENTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *REQ_IF_HEADER:
		if stage.OnAfterREQ_IF_HEADERDeleteCallback != nil {
			staged := any(staged).(*REQ_IF_HEADER)
			stage.OnAfterREQ_IF_HEADERDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *REQ_IF_TOOL_EXTENSION:
		if stage.OnAfterREQ_IF_TOOL_EXTENSIONDeleteCallback != nil {
			staged := any(staged).(*REQ_IF_TOOL_EXTENSION)
			stage.OnAfterREQ_IF_TOOL_EXTENSIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECIFICATION:
		if stage.OnAfterSPECIFICATIONDeleteCallback != nil {
			staged := any(staged).(*SPECIFICATION)
			stage.OnAfterSPECIFICATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPECIFICATION_TYPE:
		if stage.OnAfterSPECIFICATION_TYPEDeleteCallback != nil {
			staged := any(staged).(*SPECIFICATION_TYPE)
			stage.OnAfterSPECIFICATION_TYPEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPEC_HIERARCHY:
		if stage.OnAfterSPEC_HIERARCHYDeleteCallback != nil {
			staged := any(staged).(*SPEC_HIERARCHY)
			stage.OnAfterSPEC_HIERARCHYDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPEC_OBJECT:
		if stage.OnAfterSPEC_OBJECTDeleteCallback != nil {
			staged := any(staged).(*SPEC_OBJECT)
			stage.OnAfterSPEC_OBJECTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPEC_OBJECT_TYPE:
		if stage.OnAfterSPEC_OBJECT_TYPEDeleteCallback != nil {
			staged := any(staged).(*SPEC_OBJECT_TYPE)
			stage.OnAfterSPEC_OBJECT_TYPEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPEC_RELATION:
		if stage.OnAfterSPEC_RELATIONDeleteCallback != nil {
			staged := any(staged).(*SPEC_RELATION)
			stage.OnAfterSPEC_RELATIONDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *SPEC_RELATION_TYPE:
		if stage.OnAfterSPEC_RELATION_TYPEDeleteCallback != nil {
			staged := any(staged).(*SPEC_RELATION_TYPE)
			stage.OnAfterSPEC_RELATION_TYPEDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	case *XHTML_CONTENT:
		if stage.OnAfterXHTML_CONTENTDeleteCallback != nil {
			staged := any(staged).(*XHTML_CONTENT)
			stage.OnAfterXHTML_CONTENTDeleteCallback.OnAfterDelete(stage, staged, front)
		}
	default:
		_ = front
	}
}

// AfterReadFromFront is called after a Read from front
func AfterReadFromFront[Type Gongstruct](stage *Stage, instance *Type) {

	switch target := any(instance).(type) {
	// insertion point
	case *ALTERNATIVE_ID:
		if stage.OnAfterALTERNATIVE_IDReadCallback != nil {
			stage.OnAfterALTERNATIVE_IDReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		if stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANReadCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_DATE:
		if stage.OnAfterATTRIBUTE_DEFINITION_DATEReadCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_DATEReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		if stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONReadCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_INTEGER:
		if stage.OnAfterATTRIBUTE_DEFINITION_INTEGERReadCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_INTEGERReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_REAL:
		if stage.OnAfterATTRIBUTE_DEFINITION_REALReadCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_REALReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_STRING:
		if stage.OnAfterATTRIBUTE_DEFINITION_STRINGReadCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_STRINGReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_DEFINITION_XHTML:
		if stage.OnAfterATTRIBUTE_DEFINITION_XHTMLReadCallback != nil {
			stage.OnAfterATTRIBUTE_DEFINITION_XHTMLReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_VALUE_BOOLEAN:
		if stage.OnAfterATTRIBUTE_VALUE_BOOLEANReadCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_BOOLEANReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_VALUE_DATE:
		if stage.OnAfterATTRIBUTE_VALUE_DATEReadCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_DATEReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_VALUE_ENUMERATION:
		if stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONReadCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_VALUE_INTEGER:
		if stage.OnAfterATTRIBUTE_VALUE_INTEGERReadCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_INTEGERReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_VALUE_REAL:
		if stage.OnAfterATTRIBUTE_VALUE_REALReadCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_REALReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_VALUE_STRING:
		if stage.OnAfterATTRIBUTE_VALUE_STRINGReadCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_STRINGReadCallback.OnAfterRead(stage, target)
		}
	case *ATTRIBUTE_VALUE_XHTML:
		if stage.OnAfterATTRIBUTE_VALUE_XHTMLReadCallback != nil {
			stage.OnAfterATTRIBUTE_VALUE_XHTMLReadCallback.OnAfterRead(stage, target)
		}
	case *A_ALTERNATIVE_ID:
		if stage.OnAfterA_ALTERNATIVE_IDReadCallback != nil {
			stage.OnAfterA_ALTERNATIVE_IDReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		if stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		if stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_DATE:
		if stage.OnAfterA_ATTRIBUTE_VALUE_DATEReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_DATEReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		if stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_INTEGER:
		if stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_REAL:
		if stage.OnAfterA_ATTRIBUTE_VALUE_REALReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_REALReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_STRING:
		if stage.OnAfterA_ATTRIBUTE_VALUE_STRINGReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_STRINGReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_XHTML:
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLReadCallback.OnAfterRead(stage, target)
		}
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		if stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1ReadCallback != nil {
			stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1ReadCallback.OnAfterRead(stage, target)
		}
	case *A_CHILDREN:
		if stage.OnAfterA_CHILDRENReadCallback != nil {
			stage.OnAfterA_CHILDRENReadCallback.OnAfterRead(stage, target)
		}
	case *A_CORE_CONTENT:
		if stage.OnAfterA_CORE_CONTENTReadCallback != nil {
			stage.OnAfterA_CORE_CONTENTReadCallback.OnAfterRead(stage, target)
		}
	case *A_DATATYPES:
		if stage.OnAfterA_DATATYPESReadCallback != nil {
			stage.OnAfterA_DATATYPESReadCallback.OnAfterRead(stage, target)
		}
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFReadCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_DATATYPE_DEFINITION_DATE_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFReadCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFReadCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFReadCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_DATATYPE_DEFINITION_REAL_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFReadCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_DATATYPE_DEFINITION_STRING_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFReadCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		if stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFReadCallback != nil {
			stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_EDITABLE_ATTS:
		if stage.OnAfterA_EDITABLE_ATTSReadCallback != nil {
			stage.OnAfterA_EDITABLE_ATTSReadCallback.OnAfterRead(stage, target)
		}
	case *A_ENUM_VALUE_REF:
		if stage.OnAfterA_ENUM_VALUE_REFReadCallback != nil {
			stage.OnAfterA_ENUM_VALUE_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_OBJECT:
		if stage.OnAfterA_OBJECTReadCallback != nil {
			stage.OnAfterA_OBJECTReadCallback.OnAfterRead(stage, target)
		}
	case *A_PROPERTIES:
		if stage.OnAfterA_PROPERTIESReadCallback != nil {
			stage.OnAfterA_PROPERTIESReadCallback.OnAfterRead(stage, target)
		}
	case *A_RELATION_GROUP_TYPE_REF:
		if stage.OnAfterA_RELATION_GROUP_TYPE_REFReadCallback != nil {
			stage.OnAfterA_RELATION_GROUP_TYPE_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_SOURCE_1:
		if stage.OnAfterA_SOURCE_1ReadCallback != nil {
			stage.OnAfterA_SOURCE_1ReadCallback.OnAfterRead(stage, target)
		}
	case *A_SOURCE_SPECIFICATION_1:
		if stage.OnAfterA_SOURCE_SPECIFICATION_1ReadCallback != nil {
			stage.OnAfterA_SOURCE_SPECIFICATION_1ReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPECIFICATIONS:
		if stage.OnAfterA_SPECIFICATIONSReadCallback != nil {
			stage.OnAfterA_SPECIFICATIONSReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPECIFICATION_TYPE_REF:
		if stage.OnAfterA_SPECIFICATION_TYPE_REFReadCallback != nil {
			stage.OnAfterA_SPECIFICATION_TYPE_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPECIFIED_VALUES:
		if stage.OnAfterA_SPECIFIED_VALUESReadCallback != nil {
			stage.OnAfterA_SPECIFIED_VALUESReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPEC_ATTRIBUTES:
		if stage.OnAfterA_SPEC_ATTRIBUTESReadCallback != nil {
			stage.OnAfterA_SPEC_ATTRIBUTESReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPEC_OBJECTS:
		if stage.OnAfterA_SPEC_OBJECTSReadCallback != nil {
			stage.OnAfterA_SPEC_OBJECTSReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPEC_OBJECT_TYPE_REF:
		if stage.OnAfterA_SPEC_OBJECT_TYPE_REFReadCallback != nil {
			stage.OnAfterA_SPEC_OBJECT_TYPE_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPEC_RELATIONS:
		if stage.OnAfterA_SPEC_RELATIONSReadCallback != nil {
			stage.OnAfterA_SPEC_RELATIONSReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPEC_RELATION_GROUPS:
		if stage.OnAfterA_SPEC_RELATION_GROUPSReadCallback != nil {
			stage.OnAfterA_SPEC_RELATION_GROUPSReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPEC_RELATION_REF:
		if stage.OnAfterA_SPEC_RELATION_REFReadCallback != nil {
			stage.OnAfterA_SPEC_RELATION_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPEC_RELATION_TYPE_REF:
		if stage.OnAfterA_SPEC_RELATION_TYPE_REFReadCallback != nil {
			stage.OnAfterA_SPEC_RELATION_TYPE_REFReadCallback.OnAfterRead(stage, target)
		}
	case *A_SPEC_TYPES:
		if stage.OnAfterA_SPEC_TYPESReadCallback != nil {
			stage.OnAfterA_SPEC_TYPESReadCallback.OnAfterRead(stage, target)
		}
	case *A_THE_HEADER:
		if stage.OnAfterA_THE_HEADERReadCallback != nil {
			stage.OnAfterA_THE_HEADERReadCallback.OnAfterRead(stage, target)
		}
	case *A_TOOL_EXTENSIONS:
		if stage.OnAfterA_TOOL_EXTENSIONSReadCallback != nil {
			stage.OnAfterA_TOOL_EXTENSIONSReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPE_DEFINITION_BOOLEAN:
		if stage.OnAfterDATATYPE_DEFINITION_BOOLEANReadCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_BOOLEANReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPE_DEFINITION_DATE:
		if stage.OnAfterDATATYPE_DEFINITION_DATEReadCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_DATEReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPE_DEFINITION_ENUMERATION:
		if stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONReadCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPE_DEFINITION_INTEGER:
		if stage.OnAfterDATATYPE_DEFINITION_INTEGERReadCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_INTEGERReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPE_DEFINITION_REAL:
		if stage.OnAfterDATATYPE_DEFINITION_REALReadCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_REALReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPE_DEFINITION_STRING:
		if stage.OnAfterDATATYPE_DEFINITION_STRINGReadCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_STRINGReadCallback.OnAfterRead(stage, target)
		}
	case *DATATYPE_DEFINITION_XHTML:
		if stage.OnAfterDATATYPE_DEFINITION_XHTMLReadCallback != nil {
			stage.OnAfterDATATYPE_DEFINITION_XHTMLReadCallback.OnAfterRead(stage, target)
		}
	case *EMBEDDED_VALUE:
		if stage.OnAfterEMBEDDED_VALUEReadCallback != nil {
			stage.OnAfterEMBEDDED_VALUEReadCallback.OnAfterRead(stage, target)
		}
	case *ENUM_VALUE:
		if stage.OnAfterENUM_VALUEReadCallback != nil {
			stage.OnAfterENUM_VALUEReadCallback.OnAfterRead(stage, target)
		}
	case *RELATION_GROUP:
		if stage.OnAfterRELATION_GROUPReadCallback != nil {
			stage.OnAfterRELATION_GROUPReadCallback.OnAfterRead(stage, target)
		}
	case *RELATION_GROUP_TYPE:
		if stage.OnAfterRELATION_GROUP_TYPEReadCallback != nil {
			stage.OnAfterRELATION_GROUP_TYPEReadCallback.OnAfterRead(stage, target)
		}
	case *REQ_IF:
		if stage.OnAfterREQ_IFReadCallback != nil {
			stage.OnAfterREQ_IFReadCallback.OnAfterRead(stage, target)
		}
	case *REQ_IF_CONTENT:
		if stage.OnAfterREQ_IF_CONTENTReadCallback != nil {
			stage.OnAfterREQ_IF_CONTENTReadCallback.OnAfterRead(stage, target)
		}
	case *REQ_IF_HEADER:
		if stage.OnAfterREQ_IF_HEADERReadCallback != nil {
			stage.OnAfterREQ_IF_HEADERReadCallback.OnAfterRead(stage, target)
		}
	case *REQ_IF_TOOL_EXTENSION:
		if stage.OnAfterREQ_IF_TOOL_EXTENSIONReadCallback != nil {
			stage.OnAfterREQ_IF_TOOL_EXTENSIONReadCallback.OnAfterRead(stage, target)
		}
	case *SPECIFICATION:
		if stage.OnAfterSPECIFICATIONReadCallback != nil {
			stage.OnAfterSPECIFICATIONReadCallback.OnAfterRead(stage, target)
		}
	case *SPECIFICATION_TYPE:
		if stage.OnAfterSPECIFICATION_TYPEReadCallback != nil {
			stage.OnAfterSPECIFICATION_TYPEReadCallback.OnAfterRead(stage, target)
		}
	case *SPEC_HIERARCHY:
		if stage.OnAfterSPEC_HIERARCHYReadCallback != nil {
			stage.OnAfterSPEC_HIERARCHYReadCallback.OnAfterRead(stage, target)
		}
	case *SPEC_OBJECT:
		if stage.OnAfterSPEC_OBJECTReadCallback != nil {
			stage.OnAfterSPEC_OBJECTReadCallback.OnAfterRead(stage, target)
		}
	case *SPEC_OBJECT_TYPE:
		if stage.OnAfterSPEC_OBJECT_TYPEReadCallback != nil {
			stage.OnAfterSPEC_OBJECT_TYPEReadCallback.OnAfterRead(stage, target)
		}
	case *SPEC_RELATION:
		if stage.OnAfterSPEC_RELATIONReadCallback != nil {
			stage.OnAfterSPEC_RELATIONReadCallback.OnAfterRead(stage, target)
		}
	case *SPEC_RELATION_TYPE:
		if stage.OnAfterSPEC_RELATION_TYPEReadCallback != nil {
			stage.OnAfterSPEC_RELATION_TYPEReadCallback.OnAfterRead(stage, target)
		}
	case *XHTML_CONTENT:
		if stage.OnAfterXHTML_CONTENTReadCallback != nil {
			stage.OnAfterXHTML_CONTENTReadCallback.OnAfterRead(stage, target)
		}
	default:
		_ = target
	}
}

// SetCallbackAfterUpdateFromFront is a function to set up callback that is robust to refactoring
func SetCallbackAfterUpdateFromFront[Type Gongstruct](stage *Stage, callback OnAfterUpdateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
	// insertion point
	case *ALTERNATIVE_ID:
		stage.OnAfterALTERNATIVE_IDUpdateCallback = any(callback).(OnAfterUpdateInterface[ALTERNATIVE_ID])
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_BOOLEAN])
	case *ATTRIBUTE_DEFINITION_DATE:
		stage.OnAfterATTRIBUTE_DEFINITION_DATEUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_DATE])
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_ENUMERATION])
	case *ATTRIBUTE_DEFINITION_INTEGER:
		stage.OnAfterATTRIBUTE_DEFINITION_INTEGERUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_INTEGER])
	case *ATTRIBUTE_DEFINITION_REAL:
		stage.OnAfterATTRIBUTE_DEFINITION_REALUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_REAL])
	case *ATTRIBUTE_DEFINITION_STRING:
		stage.OnAfterATTRIBUTE_DEFINITION_STRINGUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_STRING])
	case *ATTRIBUTE_DEFINITION_XHTML:
		stage.OnAfterATTRIBUTE_DEFINITION_XHTMLUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_DEFINITION_XHTML])
	case *ATTRIBUTE_VALUE_BOOLEAN:
		stage.OnAfterATTRIBUTE_VALUE_BOOLEANUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_VALUE_BOOLEAN])
	case *ATTRIBUTE_VALUE_DATE:
		stage.OnAfterATTRIBUTE_VALUE_DATEUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_VALUE_DATE])
	case *ATTRIBUTE_VALUE_ENUMERATION:
		stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_VALUE_ENUMERATION])
	case *ATTRIBUTE_VALUE_INTEGER:
		stage.OnAfterATTRIBUTE_VALUE_INTEGERUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_VALUE_INTEGER])
	case *ATTRIBUTE_VALUE_REAL:
		stage.OnAfterATTRIBUTE_VALUE_REALUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_VALUE_REAL])
	case *ATTRIBUTE_VALUE_STRING:
		stage.OnAfterATTRIBUTE_VALUE_STRINGUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_VALUE_STRING])
	case *ATTRIBUTE_VALUE_XHTML:
		stage.OnAfterATTRIBUTE_VALUE_XHTMLUpdateCallback = any(callback).(OnAfterUpdateInterface[ATTRIBUTE_VALUE_XHTML])
	case *A_ALTERNATIVE_ID:
		stage.OnAfterA_ALTERNATIVE_IDUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ALTERNATIVE_ID])
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_BOOLEAN_REF])
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_DATE_REF])
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_ENUMERATION_REF])
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_INTEGER_REF])
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_REAL_REF])
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_STRING_REF])
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_DEFINITION_XHTML_REF])
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_BOOLEAN])
	case *A_ATTRIBUTE_VALUE_DATE:
		stage.OnAfterA_ATTRIBUTE_VALUE_DATEUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_DATE])
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_ENUMERATION])
	case *A_ATTRIBUTE_VALUE_INTEGER:
		stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_INTEGER])
	case *A_ATTRIBUTE_VALUE_REAL:
		stage.OnAfterA_ATTRIBUTE_VALUE_REALUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_REAL])
	case *A_ATTRIBUTE_VALUE_STRING:
		stage.OnAfterA_ATTRIBUTE_VALUE_STRINGUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_STRING])
	case *A_ATTRIBUTE_VALUE_XHTML:
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_XHTML])
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1UpdateCallback = any(callback).(OnAfterUpdateInterface[A_ATTRIBUTE_VALUE_XHTML_1])
	case *A_CHILDREN:
		stage.OnAfterA_CHILDRENUpdateCallback = any(callback).(OnAfterUpdateInterface[A_CHILDREN])
	case *A_CORE_CONTENT:
		stage.OnAfterA_CORE_CONTENTUpdateCallback = any(callback).(OnAfterUpdateInterface[A_CORE_CONTENT])
	case *A_DATATYPES:
		stage.OnAfterA_DATATYPESUpdateCallback = any(callback).(OnAfterUpdateInterface[A_DATATYPES])
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_DATATYPE_DEFINITION_BOOLEAN_REF])
	case *A_DATATYPE_DEFINITION_DATE_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_DATATYPE_DEFINITION_DATE_REF])
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_DATATYPE_DEFINITION_ENUMERATION_REF])
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_DATATYPE_DEFINITION_INTEGER_REF])
	case *A_DATATYPE_DEFINITION_REAL_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_DATATYPE_DEFINITION_REAL_REF])
	case *A_DATATYPE_DEFINITION_STRING_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_DATATYPE_DEFINITION_STRING_REF])
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_DATATYPE_DEFINITION_XHTML_REF])
	case *A_EDITABLE_ATTS:
		stage.OnAfterA_EDITABLE_ATTSUpdateCallback = any(callback).(OnAfterUpdateInterface[A_EDITABLE_ATTS])
	case *A_ENUM_VALUE_REF:
		stage.OnAfterA_ENUM_VALUE_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_ENUM_VALUE_REF])
	case *A_OBJECT:
		stage.OnAfterA_OBJECTUpdateCallback = any(callback).(OnAfterUpdateInterface[A_OBJECT])
	case *A_PROPERTIES:
		stage.OnAfterA_PROPERTIESUpdateCallback = any(callback).(OnAfterUpdateInterface[A_PROPERTIES])
	case *A_RELATION_GROUP_TYPE_REF:
		stage.OnAfterA_RELATION_GROUP_TYPE_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_RELATION_GROUP_TYPE_REF])
	case *A_SOURCE_1:
		stage.OnAfterA_SOURCE_1UpdateCallback = any(callback).(OnAfterUpdateInterface[A_SOURCE_1])
	case *A_SOURCE_SPECIFICATION_1:
		stage.OnAfterA_SOURCE_SPECIFICATION_1UpdateCallback = any(callback).(OnAfterUpdateInterface[A_SOURCE_SPECIFICATION_1])
	case *A_SPECIFICATIONS:
		stage.OnAfterA_SPECIFICATIONSUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPECIFICATIONS])
	case *A_SPECIFICATION_TYPE_REF:
		stage.OnAfterA_SPECIFICATION_TYPE_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPECIFICATION_TYPE_REF])
	case *A_SPECIFIED_VALUES:
		stage.OnAfterA_SPECIFIED_VALUESUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPECIFIED_VALUES])
	case *A_SPEC_ATTRIBUTES:
		stage.OnAfterA_SPEC_ATTRIBUTESUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPEC_ATTRIBUTES])
	case *A_SPEC_OBJECTS:
		stage.OnAfterA_SPEC_OBJECTSUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPEC_OBJECTS])
	case *A_SPEC_OBJECT_TYPE_REF:
		stage.OnAfterA_SPEC_OBJECT_TYPE_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPEC_OBJECT_TYPE_REF])
	case *A_SPEC_RELATIONS:
		stage.OnAfterA_SPEC_RELATIONSUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPEC_RELATIONS])
	case *A_SPEC_RELATION_GROUPS:
		stage.OnAfterA_SPEC_RELATION_GROUPSUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPEC_RELATION_GROUPS])
	case *A_SPEC_RELATION_REF:
		stage.OnAfterA_SPEC_RELATION_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPEC_RELATION_REF])
	case *A_SPEC_RELATION_TYPE_REF:
		stage.OnAfterA_SPEC_RELATION_TYPE_REFUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPEC_RELATION_TYPE_REF])
	case *A_SPEC_TYPES:
		stage.OnAfterA_SPEC_TYPESUpdateCallback = any(callback).(OnAfterUpdateInterface[A_SPEC_TYPES])
	case *A_THE_HEADER:
		stage.OnAfterA_THE_HEADERUpdateCallback = any(callback).(OnAfterUpdateInterface[A_THE_HEADER])
	case *A_TOOL_EXTENSIONS:
		stage.OnAfterA_TOOL_EXTENSIONSUpdateCallback = any(callback).(OnAfterUpdateInterface[A_TOOL_EXTENSIONS])
	case *DATATYPE_DEFINITION_BOOLEAN:
		stage.OnAfterDATATYPE_DEFINITION_BOOLEANUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPE_DEFINITION_BOOLEAN])
	case *DATATYPE_DEFINITION_DATE:
		stage.OnAfterDATATYPE_DEFINITION_DATEUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPE_DEFINITION_DATE])
	case *DATATYPE_DEFINITION_ENUMERATION:
		stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPE_DEFINITION_ENUMERATION])
	case *DATATYPE_DEFINITION_INTEGER:
		stage.OnAfterDATATYPE_DEFINITION_INTEGERUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPE_DEFINITION_INTEGER])
	case *DATATYPE_DEFINITION_REAL:
		stage.OnAfterDATATYPE_DEFINITION_REALUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPE_DEFINITION_REAL])
	case *DATATYPE_DEFINITION_STRING:
		stage.OnAfterDATATYPE_DEFINITION_STRINGUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPE_DEFINITION_STRING])
	case *DATATYPE_DEFINITION_XHTML:
		stage.OnAfterDATATYPE_DEFINITION_XHTMLUpdateCallback = any(callback).(OnAfterUpdateInterface[DATATYPE_DEFINITION_XHTML])
	case *EMBEDDED_VALUE:
		stage.OnAfterEMBEDDED_VALUEUpdateCallback = any(callback).(OnAfterUpdateInterface[EMBEDDED_VALUE])
	case *ENUM_VALUE:
		stage.OnAfterENUM_VALUEUpdateCallback = any(callback).(OnAfterUpdateInterface[ENUM_VALUE])
	case *RELATION_GROUP:
		stage.OnAfterRELATION_GROUPUpdateCallback = any(callback).(OnAfterUpdateInterface[RELATION_GROUP])
	case *RELATION_GROUP_TYPE:
		stage.OnAfterRELATION_GROUP_TYPEUpdateCallback = any(callback).(OnAfterUpdateInterface[RELATION_GROUP_TYPE])
	case *REQ_IF:
		stage.OnAfterREQ_IFUpdateCallback = any(callback).(OnAfterUpdateInterface[REQ_IF])
	case *REQ_IF_CONTENT:
		stage.OnAfterREQ_IF_CONTENTUpdateCallback = any(callback).(OnAfterUpdateInterface[REQ_IF_CONTENT])
	case *REQ_IF_HEADER:
		stage.OnAfterREQ_IF_HEADERUpdateCallback = any(callback).(OnAfterUpdateInterface[REQ_IF_HEADER])
	case *REQ_IF_TOOL_EXTENSION:
		stage.OnAfterREQ_IF_TOOL_EXTENSIONUpdateCallback = any(callback).(OnAfterUpdateInterface[REQ_IF_TOOL_EXTENSION])
	case *SPECIFICATION:
		stage.OnAfterSPECIFICATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECIFICATION])
	case *SPECIFICATION_TYPE:
		stage.OnAfterSPECIFICATION_TYPEUpdateCallback = any(callback).(OnAfterUpdateInterface[SPECIFICATION_TYPE])
	case *SPEC_HIERARCHY:
		stage.OnAfterSPEC_HIERARCHYUpdateCallback = any(callback).(OnAfterUpdateInterface[SPEC_HIERARCHY])
	case *SPEC_OBJECT:
		stage.OnAfterSPEC_OBJECTUpdateCallback = any(callback).(OnAfterUpdateInterface[SPEC_OBJECT])
	case *SPEC_OBJECT_TYPE:
		stage.OnAfterSPEC_OBJECT_TYPEUpdateCallback = any(callback).(OnAfterUpdateInterface[SPEC_OBJECT_TYPE])
	case *SPEC_RELATION:
		stage.OnAfterSPEC_RELATIONUpdateCallback = any(callback).(OnAfterUpdateInterface[SPEC_RELATION])
	case *SPEC_RELATION_TYPE:
		stage.OnAfterSPEC_RELATION_TYPEUpdateCallback = any(callback).(OnAfterUpdateInterface[SPEC_RELATION_TYPE])
	case *XHTML_CONTENT:
		stage.OnAfterXHTML_CONTENTUpdateCallback = any(callback).(OnAfterUpdateInterface[XHTML_CONTENT])
	}
}
func SetCallbackAfterCreateFromFront[Type Gongstruct](stage *Stage, callback OnAfterCreateInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ALTERNATIVE_ID:
		stage.OnAfterALTERNATIVE_IDCreateCallback = any(callback).(OnAfterCreateInterface[ALTERNATIVE_ID])
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_DEFINITION_BOOLEAN])
	case *ATTRIBUTE_DEFINITION_DATE:
		stage.OnAfterATTRIBUTE_DEFINITION_DATECreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_DEFINITION_DATE])
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_DEFINITION_ENUMERATION])
	case *ATTRIBUTE_DEFINITION_INTEGER:
		stage.OnAfterATTRIBUTE_DEFINITION_INTEGERCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_DEFINITION_INTEGER])
	case *ATTRIBUTE_DEFINITION_REAL:
		stage.OnAfterATTRIBUTE_DEFINITION_REALCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_DEFINITION_REAL])
	case *ATTRIBUTE_DEFINITION_STRING:
		stage.OnAfterATTRIBUTE_DEFINITION_STRINGCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_DEFINITION_STRING])
	case *ATTRIBUTE_DEFINITION_XHTML:
		stage.OnAfterATTRIBUTE_DEFINITION_XHTMLCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_DEFINITION_XHTML])
	case *ATTRIBUTE_VALUE_BOOLEAN:
		stage.OnAfterATTRIBUTE_VALUE_BOOLEANCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_VALUE_BOOLEAN])
	case *ATTRIBUTE_VALUE_DATE:
		stage.OnAfterATTRIBUTE_VALUE_DATECreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_VALUE_DATE])
	case *ATTRIBUTE_VALUE_ENUMERATION:
		stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_VALUE_ENUMERATION])
	case *ATTRIBUTE_VALUE_INTEGER:
		stage.OnAfterATTRIBUTE_VALUE_INTEGERCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_VALUE_INTEGER])
	case *ATTRIBUTE_VALUE_REAL:
		stage.OnAfterATTRIBUTE_VALUE_REALCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_VALUE_REAL])
	case *ATTRIBUTE_VALUE_STRING:
		stage.OnAfterATTRIBUTE_VALUE_STRINGCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_VALUE_STRING])
	case *ATTRIBUTE_VALUE_XHTML:
		stage.OnAfterATTRIBUTE_VALUE_XHTMLCreateCallback = any(callback).(OnAfterCreateInterface[ATTRIBUTE_VALUE_XHTML])
	case *A_ALTERNATIVE_ID:
		stage.OnAfterA_ALTERNATIVE_IDCreateCallback = any(callback).(OnAfterCreateInterface[A_ALTERNATIVE_ID])
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_BOOLEAN_REF])
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_DATE_REF])
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_ENUMERATION_REF])
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_INTEGER_REF])
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_REAL_REF])
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_STRING_REF])
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_DEFINITION_XHTML_REF])
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_VALUE_BOOLEAN])
	case *A_ATTRIBUTE_VALUE_DATE:
		stage.OnAfterA_ATTRIBUTE_VALUE_DATECreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_VALUE_DATE])
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_VALUE_ENUMERATION])
	case *A_ATTRIBUTE_VALUE_INTEGER:
		stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_VALUE_INTEGER])
	case *A_ATTRIBUTE_VALUE_REAL:
		stage.OnAfterA_ATTRIBUTE_VALUE_REALCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_VALUE_REAL])
	case *A_ATTRIBUTE_VALUE_STRING:
		stage.OnAfterA_ATTRIBUTE_VALUE_STRINGCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_VALUE_STRING])
	case *A_ATTRIBUTE_VALUE_XHTML:
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLCreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_VALUE_XHTML])
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1CreateCallback = any(callback).(OnAfterCreateInterface[A_ATTRIBUTE_VALUE_XHTML_1])
	case *A_CHILDREN:
		stage.OnAfterA_CHILDRENCreateCallback = any(callback).(OnAfterCreateInterface[A_CHILDREN])
	case *A_CORE_CONTENT:
		stage.OnAfterA_CORE_CONTENTCreateCallback = any(callback).(OnAfterCreateInterface[A_CORE_CONTENT])
	case *A_DATATYPES:
		stage.OnAfterA_DATATYPESCreateCallback = any(callback).(OnAfterCreateInterface[A_DATATYPES])
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_DATATYPE_DEFINITION_BOOLEAN_REF])
	case *A_DATATYPE_DEFINITION_DATE_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_DATATYPE_DEFINITION_DATE_REF])
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_DATATYPE_DEFINITION_ENUMERATION_REF])
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_DATATYPE_DEFINITION_INTEGER_REF])
	case *A_DATATYPE_DEFINITION_REAL_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_DATATYPE_DEFINITION_REAL_REF])
	case *A_DATATYPE_DEFINITION_STRING_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_DATATYPE_DEFINITION_STRING_REF])
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_DATATYPE_DEFINITION_XHTML_REF])
	case *A_EDITABLE_ATTS:
		stage.OnAfterA_EDITABLE_ATTSCreateCallback = any(callback).(OnAfterCreateInterface[A_EDITABLE_ATTS])
	case *A_ENUM_VALUE_REF:
		stage.OnAfterA_ENUM_VALUE_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_ENUM_VALUE_REF])
	case *A_OBJECT:
		stage.OnAfterA_OBJECTCreateCallback = any(callback).(OnAfterCreateInterface[A_OBJECT])
	case *A_PROPERTIES:
		stage.OnAfterA_PROPERTIESCreateCallback = any(callback).(OnAfterCreateInterface[A_PROPERTIES])
	case *A_RELATION_GROUP_TYPE_REF:
		stage.OnAfterA_RELATION_GROUP_TYPE_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_RELATION_GROUP_TYPE_REF])
	case *A_SOURCE_1:
		stage.OnAfterA_SOURCE_1CreateCallback = any(callback).(OnAfterCreateInterface[A_SOURCE_1])
	case *A_SOURCE_SPECIFICATION_1:
		stage.OnAfterA_SOURCE_SPECIFICATION_1CreateCallback = any(callback).(OnAfterCreateInterface[A_SOURCE_SPECIFICATION_1])
	case *A_SPECIFICATIONS:
		stage.OnAfterA_SPECIFICATIONSCreateCallback = any(callback).(OnAfterCreateInterface[A_SPECIFICATIONS])
	case *A_SPECIFICATION_TYPE_REF:
		stage.OnAfterA_SPECIFICATION_TYPE_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_SPECIFICATION_TYPE_REF])
	case *A_SPECIFIED_VALUES:
		stage.OnAfterA_SPECIFIED_VALUESCreateCallback = any(callback).(OnAfterCreateInterface[A_SPECIFIED_VALUES])
	case *A_SPEC_ATTRIBUTES:
		stage.OnAfterA_SPEC_ATTRIBUTESCreateCallback = any(callback).(OnAfterCreateInterface[A_SPEC_ATTRIBUTES])
	case *A_SPEC_OBJECTS:
		stage.OnAfterA_SPEC_OBJECTSCreateCallback = any(callback).(OnAfterCreateInterface[A_SPEC_OBJECTS])
	case *A_SPEC_OBJECT_TYPE_REF:
		stage.OnAfterA_SPEC_OBJECT_TYPE_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_SPEC_OBJECT_TYPE_REF])
	case *A_SPEC_RELATIONS:
		stage.OnAfterA_SPEC_RELATIONSCreateCallback = any(callback).(OnAfterCreateInterface[A_SPEC_RELATIONS])
	case *A_SPEC_RELATION_GROUPS:
		stage.OnAfterA_SPEC_RELATION_GROUPSCreateCallback = any(callback).(OnAfterCreateInterface[A_SPEC_RELATION_GROUPS])
	case *A_SPEC_RELATION_REF:
		stage.OnAfterA_SPEC_RELATION_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_SPEC_RELATION_REF])
	case *A_SPEC_RELATION_TYPE_REF:
		stage.OnAfterA_SPEC_RELATION_TYPE_REFCreateCallback = any(callback).(OnAfterCreateInterface[A_SPEC_RELATION_TYPE_REF])
	case *A_SPEC_TYPES:
		stage.OnAfterA_SPEC_TYPESCreateCallback = any(callback).(OnAfterCreateInterface[A_SPEC_TYPES])
	case *A_THE_HEADER:
		stage.OnAfterA_THE_HEADERCreateCallback = any(callback).(OnAfterCreateInterface[A_THE_HEADER])
	case *A_TOOL_EXTENSIONS:
		stage.OnAfterA_TOOL_EXTENSIONSCreateCallback = any(callback).(OnAfterCreateInterface[A_TOOL_EXTENSIONS])
	case *DATATYPE_DEFINITION_BOOLEAN:
		stage.OnAfterDATATYPE_DEFINITION_BOOLEANCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPE_DEFINITION_BOOLEAN])
	case *DATATYPE_DEFINITION_DATE:
		stage.OnAfterDATATYPE_DEFINITION_DATECreateCallback = any(callback).(OnAfterCreateInterface[DATATYPE_DEFINITION_DATE])
	case *DATATYPE_DEFINITION_ENUMERATION:
		stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPE_DEFINITION_ENUMERATION])
	case *DATATYPE_DEFINITION_INTEGER:
		stage.OnAfterDATATYPE_DEFINITION_INTEGERCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPE_DEFINITION_INTEGER])
	case *DATATYPE_DEFINITION_REAL:
		stage.OnAfterDATATYPE_DEFINITION_REALCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPE_DEFINITION_REAL])
	case *DATATYPE_DEFINITION_STRING:
		stage.OnAfterDATATYPE_DEFINITION_STRINGCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPE_DEFINITION_STRING])
	case *DATATYPE_DEFINITION_XHTML:
		stage.OnAfterDATATYPE_DEFINITION_XHTMLCreateCallback = any(callback).(OnAfterCreateInterface[DATATYPE_DEFINITION_XHTML])
	case *EMBEDDED_VALUE:
		stage.OnAfterEMBEDDED_VALUECreateCallback = any(callback).(OnAfterCreateInterface[EMBEDDED_VALUE])
	case *ENUM_VALUE:
		stage.OnAfterENUM_VALUECreateCallback = any(callback).(OnAfterCreateInterface[ENUM_VALUE])
	case *RELATION_GROUP:
		stage.OnAfterRELATION_GROUPCreateCallback = any(callback).(OnAfterCreateInterface[RELATION_GROUP])
	case *RELATION_GROUP_TYPE:
		stage.OnAfterRELATION_GROUP_TYPECreateCallback = any(callback).(OnAfterCreateInterface[RELATION_GROUP_TYPE])
	case *REQ_IF:
		stage.OnAfterREQ_IFCreateCallback = any(callback).(OnAfterCreateInterface[REQ_IF])
	case *REQ_IF_CONTENT:
		stage.OnAfterREQ_IF_CONTENTCreateCallback = any(callback).(OnAfterCreateInterface[REQ_IF_CONTENT])
	case *REQ_IF_HEADER:
		stage.OnAfterREQ_IF_HEADERCreateCallback = any(callback).(OnAfterCreateInterface[REQ_IF_HEADER])
	case *REQ_IF_TOOL_EXTENSION:
		stage.OnAfterREQ_IF_TOOL_EXTENSIONCreateCallback = any(callback).(OnAfterCreateInterface[REQ_IF_TOOL_EXTENSION])
	case *SPECIFICATION:
		stage.OnAfterSPECIFICATIONCreateCallback = any(callback).(OnAfterCreateInterface[SPECIFICATION])
	case *SPECIFICATION_TYPE:
		stage.OnAfterSPECIFICATION_TYPECreateCallback = any(callback).(OnAfterCreateInterface[SPECIFICATION_TYPE])
	case *SPEC_HIERARCHY:
		stage.OnAfterSPEC_HIERARCHYCreateCallback = any(callback).(OnAfterCreateInterface[SPEC_HIERARCHY])
	case *SPEC_OBJECT:
		stage.OnAfterSPEC_OBJECTCreateCallback = any(callback).(OnAfterCreateInterface[SPEC_OBJECT])
	case *SPEC_OBJECT_TYPE:
		stage.OnAfterSPEC_OBJECT_TYPECreateCallback = any(callback).(OnAfterCreateInterface[SPEC_OBJECT_TYPE])
	case *SPEC_RELATION:
		stage.OnAfterSPEC_RELATIONCreateCallback = any(callback).(OnAfterCreateInterface[SPEC_RELATION])
	case *SPEC_RELATION_TYPE:
		stage.OnAfterSPEC_RELATION_TYPECreateCallback = any(callback).(OnAfterCreateInterface[SPEC_RELATION_TYPE])
	case *XHTML_CONTENT:
		stage.OnAfterXHTML_CONTENTCreateCallback = any(callback).(OnAfterCreateInterface[XHTML_CONTENT])
	}
}
func SetCallbackAfterDeleteFromFront[Type Gongstruct](stage *Stage, callback OnAfterDeleteInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ALTERNATIVE_ID:
		stage.OnAfterALTERNATIVE_IDDeleteCallback = any(callback).(OnAfterDeleteInterface[ALTERNATIVE_ID])
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_BOOLEAN])
	case *ATTRIBUTE_DEFINITION_DATE:
		stage.OnAfterATTRIBUTE_DEFINITION_DATEDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_DATE])
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_ENUMERATION])
	case *ATTRIBUTE_DEFINITION_INTEGER:
		stage.OnAfterATTRIBUTE_DEFINITION_INTEGERDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_INTEGER])
	case *ATTRIBUTE_DEFINITION_REAL:
		stage.OnAfterATTRIBUTE_DEFINITION_REALDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_REAL])
	case *ATTRIBUTE_DEFINITION_STRING:
		stage.OnAfterATTRIBUTE_DEFINITION_STRINGDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_STRING])
	case *ATTRIBUTE_DEFINITION_XHTML:
		stage.OnAfterATTRIBUTE_DEFINITION_XHTMLDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_DEFINITION_XHTML])
	case *ATTRIBUTE_VALUE_BOOLEAN:
		stage.OnAfterATTRIBUTE_VALUE_BOOLEANDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_VALUE_BOOLEAN])
	case *ATTRIBUTE_VALUE_DATE:
		stage.OnAfterATTRIBUTE_VALUE_DATEDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_VALUE_DATE])
	case *ATTRIBUTE_VALUE_ENUMERATION:
		stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_VALUE_ENUMERATION])
	case *ATTRIBUTE_VALUE_INTEGER:
		stage.OnAfterATTRIBUTE_VALUE_INTEGERDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_VALUE_INTEGER])
	case *ATTRIBUTE_VALUE_REAL:
		stage.OnAfterATTRIBUTE_VALUE_REALDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_VALUE_REAL])
	case *ATTRIBUTE_VALUE_STRING:
		stage.OnAfterATTRIBUTE_VALUE_STRINGDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_VALUE_STRING])
	case *ATTRIBUTE_VALUE_XHTML:
		stage.OnAfterATTRIBUTE_VALUE_XHTMLDeleteCallback = any(callback).(OnAfterDeleteInterface[ATTRIBUTE_VALUE_XHTML])
	case *A_ALTERNATIVE_ID:
		stage.OnAfterA_ALTERNATIVE_IDDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ALTERNATIVE_ID])
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_BOOLEAN_REF])
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_DATE_REF])
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_ENUMERATION_REF])
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_INTEGER_REF])
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_REAL_REF])
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_STRING_REF])
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_DEFINITION_XHTML_REF])
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_BOOLEAN])
	case *A_ATTRIBUTE_VALUE_DATE:
		stage.OnAfterA_ATTRIBUTE_VALUE_DATEDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_DATE])
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_ENUMERATION])
	case *A_ATTRIBUTE_VALUE_INTEGER:
		stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_INTEGER])
	case *A_ATTRIBUTE_VALUE_REAL:
		stage.OnAfterA_ATTRIBUTE_VALUE_REALDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_REAL])
	case *A_ATTRIBUTE_VALUE_STRING:
		stage.OnAfterA_ATTRIBUTE_VALUE_STRINGDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_STRING])
	case *A_ATTRIBUTE_VALUE_XHTML:
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_XHTML])
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1DeleteCallback = any(callback).(OnAfterDeleteInterface[A_ATTRIBUTE_VALUE_XHTML_1])
	case *A_CHILDREN:
		stage.OnAfterA_CHILDRENDeleteCallback = any(callback).(OnAfterDeleteInterface[A_CHILDREN])
	case *A_CORE_CONTENT:
		stage.OnAfterA_CORE_CONTENTDeleteCallback = any(callback).(OnAfterDeleteInterface[A_CORE_CONTENT])
	case *A_DATATYPES:
		stage.OnAfterA_DATATYPESDeleteCallback = any(callback).(OnAfterDeleteInterface[A_DATATYPES])
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_DATATYPE_DEFINITION_BOOLEAN_REF])
	case *A_DATATYPE_DEFINITION_DATE_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_DATATYPE_DEFINITION_DATE_REF])
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_DATATYPE_DEFINITION_ENUMERATION_REF])
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_DATATYPE_DEFINITION_INTEGER_REF])
	case *A_DATATYPE_DEFINITION_REAL_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_DATATYPE_DEFINITION_REAL_REF])
	case *A_DATATYPE_DEFINITION_STRING_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_DATATYPE_DEFINITION_STRING_REF])
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_DATATYPE_DEFINITION_XHTML_REF])
	case *A_EDITABLE_ATTS:
		stage.OnAfterA_EDITABLE_ATTSDeleteCallback = any(callback).(OnAfterDeleteInterface[A_EDITABLE_ATTS])
	case *A_ENUM_VALUE_REF:
		stage.OnAfterA_ENUM_VALUE_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_ENUM_VALUE_REF])
	case *A_OBJECT:
		stage.OnAfterA_OBJECTDeleteCallback = any(callback).(OnAfterDeleteInterface[A_OBJECT])
	case *A_PROPERTIES:
		stage.OnAfterA_PROPERTIESDeleteCallback = any(callback).(OnAfterDeleteInterface[A_PROPERTIES])
	case *A_RELATION_GROUP_TYPE_REF:
		stage.OnAfterA_RELATION_GROUP_TYPE_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_RELATION_GROUP_TYPE_REF])
	case *A_SOURCE_1:
		stage.OnAfterA_SOURCE_1DeleteCallback = any(callback).(OnAfterDeleteInterface[A_SOURCE_1])
	case *A_SOURCE_SPECIFICATION_1:
		stage.OnAfterA_SOURCE_SPECIFICATION_1DeleteCallback = any(callback).(OnAfterDeleteInterface[A_SOURCE_SPECIFICATION_1])
	case *A_SPECIFICATIONS:
		stage.OnAfterA_SPECIFICATIONSDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPECIFICATIONS])
	case *A_SPECIFICATION_TYPE_REF:
		stage.OnAfterA_SPECIFICATION_TYPE_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPECIFICATION_TYPE_REF])
	case *A_SPECIFIED_VALUES:
		stage.OnAfterA_SPECIFIED_VALUESDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPECIFIED_VALUES])
	case *A_SPEC_ATTRIBUTES:
		stage.OnAfterA_SPEC_ATTRIBUTESDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPEC_ATTRIBUTES])
	case *A_SPEC_OBJECTS:
		stage.OnAfterA_SPEC_OBJECTSDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPEC_OBJECTS])
	case *A_SPEC_OBJECT_TYPE_REF:
		stage.OnAfterA_SPEC_OBJECT_TYPE_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPEC_OBJECT_TYPE_REF])
	case *A_SPEC_RELATIONS:
		stage.OnAfterA_SPEC_RELATIONSDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPEC_RELATIONS])
	case *A_SPEC_RELATION_GROUPS:
		stage.OnAfterA_SPEC_RELATION_GROUPSDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPEC_RELATION_GROUPS])
	case *A_SPEC_RELATION_REF:
		stage.OnAfterA_SPEC_RELATION_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPEC_RELATION_REF])
	case *A_SPEC_RELATION_TYPE_REF:
		stage.OnAfterA_SPEC_RELATION_TYPE_REFDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPEC_RELATION_TYPE_REF])
	case *A_SPEC_TYPES:
		stage.OnAfterA_SPEC_TYPESDeleteCallback = any(callback).(OnAfterDeleteInterface[A_SPEC_TYPES])
	case *A_THE_HEADER:
		stage.OnAfterA_THE_HEADERDeleteCallback = any(callback).(OnAfterDeleteInterface[A_THE_HEADER])
	case *A_TOOL_EXTENSIONS:
		stage.OnAfterA_TOOL_EXTENSIONSDeleteCallback = any(callback).(OnAfterDeleteInterface[A_TOOL_EXTENSIONS])
	case *DATATYPE_DEFINITION_BOOLEAN:
		stage.OnAfterDATATYPE_DEFINITION_BOOLEANDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPE_DEFINITION_BOOLEAN])
	case *DATATYPE_DEFINITION_DATE:
		stage.OnAfterDATATYPE_DEFINITION_DATEDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPE_DEFINITION_DATE])
	case *DATATYPE_DEFINITION_ENUMERATION:
		stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPE_DEFINITION_ENUMERATION])
	case *DATATYPE_DEFINITION_INTEGER:
		stage.OnAfterDATATYPE_DEFINITION_INTEGERDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPE_DEFINITION_INTEGER])
	case *DATATYPE_DEFINITION_REAL:
		stage.OnAfterDATATYPE_DEFINITION_REALDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPE_DEFINITION_REAL])
	case *DATATYPE_DEFINITION_STRING:
		stage.OnAfterDATATYPE_DEFINITION_STRINGDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPE_DEFINITION_STRING])
	case *DATATYPE_DEFINITION_XHTML:
		stage.OnAfterDATATYPE_DEFINITION_XHTMLDeleteCallback = any(callback).(OnAfterDeleteInterface[DATATYPE_DEFINITION_XHTML])
	case *EMBEDDED_VALUE:
		stage.OnAfterEMBEDDED_VALUEDeleteCallback = any(callback).(OnAfterDeleteInterface[EMBEDDED_VALUE])
	case *ENUM_VALUE:
		stage.OnAfterENUM_VALUEDeleteCallback = any(callback).(OnAfterDeleteInterface[ENUM_VALUE])
	case *RELATION_GROUP:
		stage.OnAfterRELATION_GROUPDeleteCallback = any(callback).(OnAfterDeleteInterface[RELATION_GROUP])
	case *RELATION_GROUP_TYPE:
		stage.OnAfterRELATION_GROUP_TYPEDeleteCallback = any(callback).(OnAfterDeleteInterface[RELATION_GROUP_TYPE])
	case *REQ_IF:
		stage.OnAfterREQ_IFDeleteCallback = any(callback).(OnAfterDeleteInterface[REQ_IF])
	case *REQ_IF_CONTENT:
		stage.OnAfterREQ_IF_CONTENTDeleteCallback = any(callback).(OnAfterDeleteInterface[REQ_IF_CONTENT])
	case *REQ_IF_HEADER:
		stage.OnAfterREQ_IF_HEADERDeleteCallback = any(callback).(OnAfterDeleteInterface[REQ_IF_HEADER])
	case *REQ_IF_TOOL_EXTENSION:
		stage.OnAfterREQ_IF_TOOL_EXTENSIONDeleteCallback = any(callback).(OnAfterDeleteInterface[REQ_IF_TOOL_EXTENSION])
	case *SPECIFICATION:
		stage.OnAfterSPECIFICATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECIFICATION])
	case *SPECIFICATION_TYPE:
		stage.OnAfterSPECIFICATION_TYPEDeleteCallback = any(callback).(OnAfterDeleteInterface[SPECIFICATION_TYPE])
	case *SPEC_HIERARCHY:
		stage.OnAfterSPEC_HIERARCHYDeleteCallback = any(callback).(OnAfterDeleteInterface[SPEC_HIERARCHY])
	case *SPEC_OBJECT:
		stage.OnAfterSPEC_OBJECTDeleteCallback = any(callback).(OnAfterDeleteInterface[SPEC_OBJECT])
	case *SPEC_OBJECT_TYPE:
		stage.OnAfterSPEC_OBJECT_TYPEDeleteCallback = any(callback).(OnAfterDeleteInterface[SPEC_OBJECT_TYPE])
	case *SPEC_RELATION:
		stage.OnAfterSPEC_RELATIONDeleteCallback = any(callback).(OnAfterDeleteInterface[SPEC_RELATION])
	case *SPEC_RELATION_TYPE:
		stage.OnAfterSPEC_RELATION_TYPEDeleteCallback = any(callback).(OnAfterDeleteInterface[SPEC_RELATION_TYPE])
	case *XHTML_CONTENT:
		stage.OnAfterXHTML_CONTENTDeleteCallback = any(callback).(OnAfterDeleteInterface[XHTML_CONTENT])
	}
}
func SetCallbackAfterReadFromFront[Type Gongstruct](stage *Stage, callback OnAfterReadInterface[Type]) {

	var instance Type
	switch any(instance).(type) {
		// insertion point
	case *ALTERNATIVE_ID:
		stage.OnAfterALTERNATIVE_IDReadCallback = any(callback).(OnAfterReadInterface[ALTERNATIVE_ID])
	case *ATTRIBUTE_DEFINITION_BOOLEAN:
		stage.OnAfterATTRIBUTE_DEFINITION_BOOLEANReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_DEFINITION_BOOLEAN])
	case *ATTRIBUTE_DEFINITION_DATE:
		stage.OnAfterATTRIBUTE_DEFINITION_DATEReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_DEFINITION_DATE])
	case *ATTRIBUTE_DEFINITION_ENUMERATION:
		stage.OnAfterATTRIBUTE_DEFINITION_ENUMERATIONReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_DEFINITION_ENUMERATION])
	case *ATTRIBUTE_DEFINITION_INTEGER:
		stage.OnAfterATTRIBUTE_DEFINITION_INTEGERReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_DEFINITION_INTEGER])
	case *ATTRIBUTE_DEFINITION_REAL:
		stage.OnAfterATTRIBUTE_DEFINITION_REALReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_DEFINITION_REAL])
	case *ATTRIBUTE_DEFINITION_STRING:
		stage.OnAfterATTRIBUTE_DEFINITION_STRINGReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_DEFINITION_STRING])
	case *ATTRIBUTE_DEFINITION_XHTML:
		stage.OnAfterATTRIBUTE_DEFINITION_XHTMLReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_DEFINITION_XHTML])
	case *ATTRIBUTE_VALUE_BOOLEAN:
		stage.OnAfterATTRIBUTE_VALUE_BOOLEANReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_VALUE_BOOLEAN])
	case *ATTRIBUTE_VALUE_DATE:
		stage.OnAfterATTRIBUTE_VALUE_DATEReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_VALUE_DATE])
	case *ATTRIBUTE_VALUE_ENUMERATION:
		stage.OnAfterATTRIBUTE_VALUE_ENUMERATIONReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_VALUE_ENUMERATION])
	case *ATTRIBUTE_VALUE_INTEGER:
		stage.OnAfterATTRIBUTE_VALUE_INTEGERReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_VALUE_INTEGER])
	case *ATTRIBUTE_VALUE_REAL:
		stage.OnAfterATTRIBUTE_VALUE_REALReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_VALUE_REAL])
	case *ATTRIBUTE_VALUE_STRING:
		stage.OnAfterATTRIBUTE_VALUE_STRINGReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_VALUE_STRING])
	case *ATTRIBUTE_VALUE_XHTML:
		stage.OnAfterATTRIBUTE_VALUE_XHTMLReadCallback = any(callback).(OnAfterReadInterface[ATTRIBUTE_VALUE_XHTML])
	case *A_ALTERNATIVE_ID:
		stage.OnAfterA_ALTERNATIVE_IDReadCallback = any(callback).(OnAfterReadInterface[A_ALTERNATIVE_ID])
	case *A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_BOOLEAN_REFReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_BOOLEAN_REF])
	case *A_ATTRIBUTE_DEFINITION_DATE_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_DATE_REFReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_DATE_REF])
	case *A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_ENUMERATION_REFReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_ENUMERATION_REF])
	case *A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_INTEGER_REFReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_INTEGER_REF])
	case *A_ATTRIBUTE_DEFINITION_REAL_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_REAL_REFReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_REAL_REF])
	case *A_ATTRIBUTE_DEFINITION_STRING_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_STRING_REFReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_STRING_REF])
	case *A_ATTRIBUTE_DEFINITION_XHTML_REF:
		stage.OnAfterA_ATTRIBUTE_DEFINITION_XHTML_REFReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_DEFINITION_XHTML_REF])
	case *A_ATTRIBUTE_VALUE_BOOLEAN:
		stage.OnAfterA_ATTRIBUTE_VALUE_BOOLEANReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_VALUE_BOOLEAN])
	case *A_ATTRIBUTE_VALUE_DATE:
		stage.OnAfterA_ATTRIBUTE_VALUE_DATEReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_VALUE_DATE])
	case *A_ATTRIBUTE_VALUE_ENUMERATION:
		stage.OnAfterA_ATTRIBUTE_VALUE_ENUMERATIONReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_VALUE_ENUMERATION])
	case *A_ATTRIBUTE_VALUE_INTEGER:
		stage.OnAfterA_ATTRIBUTE_VALUE_INTEGERReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_VALUE_INTEGER])
	case *A_ATTRIBUTE_VALUE_REAL:
		stage.OnAfterA_ATTRIBUTE_VALUE_REALReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_VALUE_REAL])
	case *A_ATTRIBUTE_VALUE_STRING:
		stage.OnAfterA_ATTRIBUTE_VALUE_STRINGReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_VALUE_STRING])
	case *A_ATTRIBUTE_VALUE_XHTML:
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTMLReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_VALUE_XHTML])
	case *A_ATTRIBUTE_VALUE_XHTML_1:
		stage.OnAfterA_ATTRIBUTE_VALUE_XHTML_1ReadCallback = any(callback).(OnAfterReadInterface[A_ATTRIBUTE_VALUE_XHTML_1])
	case *A_CHILDREN:
		stage.OnAfterA_CHILDRENReadCallback = any(callback).(OnAfterReadInterface[A_CHILDREN])
	case *A_CORE_CONTENT:
		stage.OnAfterA_CORE_CONTENTReadCallback = any(callback).(OnAfterReadInterface[A_CORE_CONTENT])
	case *A_DATATYPES:
		stage.OnAfterA_DATATYPESReadCallback = any(callback).(OnAfterReadInterface[A_DATATYPES])
	case *A_DATATYPE_DEFINITION_BOOLEAN_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_BOOLEAN_REFReadCallback = any(callback).(OnAfterReadInterface[A_DATATYPE_DEFINITION_BOOLEAN_REF])
	case *A_DATATYPE_DEFINITION_DATE_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_DATE_REFReadCallback = any(callback).(OnAfterReadInterface[A_DATATYPE_DEFINITION_DATE_REF])
	case *A_DATATYPE_DEFINITION_ENUMERATION_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_ENUMERATION_REFReadCallback = any(callback).(OnAfterReadInterface[A_DATATYPE_DEFINITION_ENUMERATION_REF])
	case *A_DATATYPE_DEFINITION_INTEGER_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_INTEGER_REFReadCallback = any(callback).(OnAfterReadInterface[A_DATATYPE_DEFINITION_INTEGER_REF])
	case *A_DATATYPE_DEFINITION_REAL_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_REAL_REFReadCallback = any(callback).(OnAfterReadInterface[A_DATATYPE_DEFINITION_REAL_REF])
	case *A_DATATYPE_DEFINITION_STRING_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_STRING_REFReadCallback = any(callback).(OnAfterReadInterface[A_DATATYPE_DEFINITION_STRING_REF])
	case *A_DATATYPE_DEFINITION_XHTML_REF:
		stage.OnAfterA_DATATYPE_DEFINITION_XHTML_REFReadCallback = any(callback).(OnAfterReadInterface[A_DATATYPE_DEFINITION_XHTML_REF])
	case *A_EDITABLE_ATTS:
		stage.OnAfterA_EDITABLE_ATTSReadCallback = any(callback).(OnAfterReadInterface[A_EDITABLE_ATTS])
	case *A_ENUM_VALUE_REF:
		stage.OnAfterA_ENUM_VALUE_REFReadCallback = any(callback).(OnAfterReadInterface[A_ENUM_VALUE_REF])
	case *A_OBJECT:
		stage.OnAfterA_OBJECTReadCallback = any(callback).(OnAfterReadInterface[A_OBJECT])
	case *A_PROPERTIES:
		stage.OnAfterA_PROPERTIESReadCallback = any(callback).(OnAfterReadInterface[A_PROPERTIES])
	case *A_RELATION_GROUP_TYPE_REF:
		stage.OnAfterA_RELATION_GROUP_TYPE_REFReadCallback = any(callback).(OnAfterReadInterface[A_RELATION_GROUP_TYPE_REF])
	case *A_SOURCE_1:
		stage.OnAfterA_SOURCE_1ReadCallback = any(callback).(OnAfterReadInterface[A_SOURCE_1])
	case *A_SOURCE_SPECIFICATION_1:
		stage.OnAfterA_SOURCE_SPECIFICATION_1ReadCallback = any(callback).(OnAfterReadInterface[A_SOURCE_SPECIFICATION_1])
	case *A_SPECIFICATIONS:
		stage.OnAfterA_SPECIFICATIONSReadCallback = any(callback).(OnAfterReadInterface[A_SPECIFICATIONS])
	case *A_SPECIFICATION_TYPE_REF:
		stage.OnAfterA_SPECIFICATION_TYPE_REFReadCallback = any(callback).(OnAfterReadInterface[A_SPECIFICATION_TYPE_REF])
	case *A_SPECIFIED_VALUES:
		stage.OnAfterA_SPECIFIED_VALUESReadCallback = any(callback).(OnAfterReadInterface[A_SPECIFIED_VALUES])
	case *A_SPEC_ATTRIBUTES:
		stage.OnAfterA_SPEC_ATTRIBUTESReadCallback = any(callback).(OnAfterReadInterface[A_SPEC_ATTRIBUTES])
	case *A_SPEC_OBJECTS:
		stage.OnAfterA_SPEC_OBJECTSReadCallback = any(callback).(OnAfterReadInterface[A_SPEC_OBJECTS])
	case *A_SPEC_OBJECT_TYPE_REF:
		stage.OnAfterA_SPEC_OBJECT_TYPE_REFReadCallback = any(callback).(OnAfterReadInterface[A_SPEC_OBJECT_TYPE_REF])
	case *A_SPEC_RELATIONS:
		stage.OnAfterA_SPEC_RELATIONSReadCallback = any(callback).(OnAfterReadInterface[A_SPEC_RELATIONS])
	case *A_SPEC_RELATION_GROUPS:
		stage.OnAfterA_SPEC_RELATION_GROUPSReadCallback = any(callback).(OnAfterReadInterface[A_SPEC_RELATION_GROUPS])
	case *A_SPEC_RELATION_REF:
		stage.OnAfterA_SPEC_RELATION_REFReadCallback = any(callback).(OnAfterReadInterface[A_SPEC_RELATION_REF])
	case *A_SPEC_RELATION_TYPE_REF:
		stage.OnAfterA_SPEC_RELATION_TYPE_REFReadCallback = any(callback).(OnAfterReadInterface[A_SPEC_RELATION_TYPE_REF])
	case *A_SPEC_TYPES:
		stage.OnAfterA_SPEC_TYPESReadCallback = any(callback).(OnAfterReadInterface[A_SPEC_TYPES])
	case *A_THE_HEADER:
		stage.OnAfterA_THE_HEADERReadCallback = any(callback).(OnAfterReadInterface[A_THE_HEADER])
	case *A_TOOL_EXTENSIONS:
		stage.OnAfterA_TOOL_EXTENSIONSReadCallback = any(callback).(OnAfterReadInterface[A_TOOL_EXTENSIONS])
	case *DATATYPE_DEFINITION_BOOLEAN:
		stage.OnAfterDATATYPE_DEFINITION_BOOLEANReadCallback = any(callback).(OnAfterReadInterface[DATATYPE_DEFINITION_BOOLEAN])
	case *DATATYPE_DEFINITION_DATE:
		stage.OnAfterDATATYPE_DEFINITION_DATEReadCallback = any(callback).(OnAfterReadInterface[DATATYPE_DEFINITION_DATE])
	case *DATATYPE_DEFINITION_ENUMERATION:
		stage.OnAfterDATATYPE_DEFINITION_ENUMERATIONReadCallback = any(callback).(OnAfterReadInterface[DATATYPE_DEFINITION_ENUMERATION])
	case *DATATYPE_DEFINITION_INTEGER:
		stage.OnAfterDATATYPE_DEFINITION_INTEGERReadCallback = any(callback).(OnAfterReadInterface[DATATYPE_DEFINITION_INTEGER])
	case *DATATYPE_DEFINITION_REAL:
		stage.OnAfterDATATYPE_DEFINITION_REALReadCallback = any(callback).(OnAfterReadInterface[DATATYPE_DEFINITION_REAL])
	case *DATATYPE_DEFINITION_STRING:
		stage.OnAfterDATATYPE_DEFINITION_STRINGReadCallback = any(callback).(OnAfterReadInterface[DATATYPE_DEFINITION_STRING])
	case *DATATYPE_DEFINITION_XHTML:
		stage.OnAfterDATATYPE_DEFINITION_XHTMLReadCallback = any(callback).(OnAfterReadInterface[DATATYPE_DEFINITION_XHTML])
	case *EMBEDDED_VALUE:
		stage.OnAfterEMBEDDED_VALUEReadCallback = any(callback).(OnAfterReadInterface[EMBEDDED_VALUE])
	case *ENUM_VALUE:
		stage.OnAfterENUM_VALUEReadCallback = any(callback).(OnAfterReadInterface[ENUM_VALUE])
	case *RELATION_GROUP:
		stage.OnAfterRELATION_GROUPReadCallback = any(callback).(OnAfterReadInterface[RELATION_GROUP])
	case *RELATION_GROUP_TYPE:
		stage.OnAfterRELATION_GROUP_TYPEReadCallback = any(callback).(OnAfterReadInterface[RELATION_GROUP_TYPE])
	case *REQ_IF:
		stage.OnAfterREQ_IFReadCallback = any(callback).(OnAfterReadInterface[REQ_IF])
	case *REQ_IF_CONTENT:
		stage.OnAfterREQ_IF_CONTENTReadCallback = any(callback).(OnAfterReadInterface[REQ_IF_CONTENT])
	case *REQ_IF_HEADER:
		stage.OnAfterREQ_IF_HEADERReadCallback = any(callback).(OnAfterReadInterface[REQ_IF_HEADER])
	case *REQ_IF_TOOL_EXTENSION:
		stage.OnAfterREQ_IF_TOOL_EXTENSIONReadCallback = any(callback).(OnAfterReadInterface[REQ_IF_TOOL_EXTENSION])
	case *SPECIFICATION:
		stage.OnAfterSPECIFICATIONReadCallback = any(callback).(OnAfterReadInterface[SPECIFICATION])
	case *SPECIFICATION_TYPE:
		stage.OnAfterSPECIFICATION_TYPEReadCallback = any(callback).(OnAfterReadInterface[SPECIFICATION_TYPE])
	case *SPEC_HIERARCHY:
		stage.OnAfterSPEC_HIERARCHYReadCallback = any(callback).(OnAfterReadInterface[SPEC_HIERARCHY])
	case *SPEC_OBJECT:
		stage.OnAfterSPEC_OBJECTReadCallback = any(callback).(OnAfterReadInterface[SPEC_OBJECT])
	case *SPEC_OBJECT_TYPE:
		stage.OnAfterSPEC_OBJECT_TYPEReadCallback = any(callback).(OnAfterReadInterface[SPEC_OBJECT_TYPE])
	case *SPEC_RELATION:
		stage.OnAfterSPEC_RELATIONReadCallback = any(callback).(OnAfterReadInterface[SPEC_RELATION])
	case *SPEC_RELATION_TYPE:
		stage.OnAfterSPEC_RELATION_TYPEReadCallback = any(callback).(OnAfterReadInterface[SPEC_RELATION_TYPE])
	case *XHTML_CONTENT:
		stage.OnAfterXHTML_CONTENTReadCallback = any(callback).(OnAfterReadInterface[XHTML_CONTENT])
	}
}
