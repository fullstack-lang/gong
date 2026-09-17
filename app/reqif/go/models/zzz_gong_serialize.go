// generated code - do not edit
package models

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/xuri/excelize/v2"
)

func (stage *Stage) SerializeStage(filename string) {
	stage.SerializeStage2(filename, false)
}

func (stage *Stage) SerializeStage2(filename string, addIDs bool) {
	f := stage.__gong__buildExcelizeFile(addIDs)
	if err := f.SaveAs(filename); err != nil {
		fmt.Println("cannot write xl file : ", err)
	}
}

func (stage *Stage) __gong__buildExcelizeFile(addIDs bool) *excelize.File {
	f := excelize.NewFile()
	{
		// insertion point
		{
			var instances []GongstructIF
			for instance := range stage.ALTERNATIVE_IDs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ALTERNATIVE_ID", instances, (*ALTERNATIVE_ID)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_DEFINITION_BOOLEANs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_DEFINITION_BOOLEAN", instances, (*ATTRIBUTE_DEFINITION_BOOLEAN)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_DEFINITION_BOOLEAN_Renderings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_DEFINITION_BOOLEAN_Rendering", instances, (*ATTRIBUTE_DEFINITION_BOOLEAN_Rendering)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_DEFINITION_DATEs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_DEFINITION_DATE", instances, (*ATTRIBUTE_DEFINITION_DATE)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_DEFINITION_DATE_Renderings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_DEFINITION_DATE_Rendering", instances, (*ATTRIBUTE_DEFINITION_DATE_Rendering)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_DEFINITION_ENUMERATIONs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_DEFINITION_ENUMERATION", instances, (*ATTRIBUTE_DEFINITION_ENUMERATION)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_DEFINITION_ENUMERATION_Renderings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_DEFINITION_ENUMERATION_Rendering", instances, (*ATTRIBUTE_DEFINITION_ENUMERATION_Rendering)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_DEFINITION_INTEGERs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_DEFINITION_INTEGER", instances, (*ATTRIBUTE_DEFINITION_INTEGER)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_DEFINITION_INTEGER_Renderings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_DEFINITION_INTEGER_Rendering", instances, (*ATTRIBUTE_DEFINITION_INTEGER_Rendering)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_DEFINITION_REALs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_DEFINITION_REAL", instances, (*ATTRIBUTE_DEFINITION_REAL)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_DEFINITION_REAL_Renderings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_DEFINITION_REAL_Rendering", instances, (*ATTRIBUTE_DEFINITION_REAL_Rendering)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_DEFINITION_Renderings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_DEFINITION_Rendering", instances, (*ATTRIBUTE_DEFINITION_Rendering)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_DEFINITION_STRINGs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_DEFINITION_STRING", instances, (*ATTRIBUTE_DEFINITION_STRING)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_DEFINITION_STRING_Renderings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_DEFINITION_STRING_Rendering", instances, (*ATTRIBUTE_DEFINITION_STRING_Rendering)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_DEFINITION_XHTMLs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_DEFINITION_XHTML", instances, (*ATTRIBUTE_DEFINITION_XHTML)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_DEFINITION_XHTML_Renderings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_DEFINITION_XHTML_Rendering", instances, (*ATTRIBUTE_DEFINITION_XHTML_Rendering)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_VALUE_BOOLEANs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_VALUE_BOOLEAN", instances, (*ATTRIBUTE_VALUE_BOOLEAN)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_VALUE_DATEs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_VALUE_DATE", instances, (*ATTRIBUTE_VALUE_DATE)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_VALUE_ENUMERATIONs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_VALUE_ENUMERATION", instances, (*ATTRIBUTE_VALUE_ENUMERATION)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_VALUE_INTEGERs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_VALUE_INTEGER", instances, (*ATTRIBUTE_VALUE_INTEGER)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_VALUE_REALs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_VALUE_REAL", instances, (*ATTRIBUTE_VALUE_REAL)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_VALUE_STRINGs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_VALUE_STRING", instances, (*ATTRIBUTE_VALUE_STRING)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ATTRIBUTE_VALUE_XHTMLs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ATTRIBUTE_VALUE_XHTML", instances, (*ATTRIBUTE_VALUE_XHTML)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ALTERNATIVE_IDs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ALTERNATIVE_ID", instances, (*A_ALTERNATIVE_ID)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ATTRIBUTE_DEFINITION_BOOLEAN_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ATTRIBUTE_DEFINITION_BOOLEAN_REF", instances, (*A_ATTRIBUTE_DEFINITION_BOOLEAN_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ATTRIBUTE_DEFINITION_DATE_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ATTRIBUTE_DEFINITION_DATE_REF", instances, (*A_ATTRIBUTE_DEFINITION_DATE_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ATTRIBUTE_DEFINITION_ENUMERATION_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ATTRIBUTE_DEFINITION_ENUMERATION_REF", instances, (*A_ATTRIBUTE_DEFINITION_ENUMERATION_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ATTRIBUTE_DEFINITION_INTEGER_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ATTRIBUTE_DEFINITION_INTEGER_REF", instances, (*A_ATTRIBUTE_DEFINITION_INTEGER_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ATTRIBUTE_DEFINITION_REAL_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ATTRIBUTE_DEFINITION_REAL_REF", instances, (*A_ATTRIBUTE_DEFINITION_REAL_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ATTRIBUTE_DEFINITION_STRING_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ATTRIBUTE_DEFINITION_STRING_REF", instances, (*A_ATTRIBUTE_DEFINITION_STRING_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ATTRIBUTE_DEFINITION_XHTML_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ATTRIBUTE_DEFINITION_XHTML_REF", instances, (*A_ATTRIBUTE_DEFINITION_XHTML_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ATTRIBUTE_VALUE_BOOLEANs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ATTRIBUTE_VALUE_BOOLEAN", instances, (*A_ATTRIBUTE_VALUE_BOOLEAN)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ATTRIBUTE_VALUE_DATEs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ATTRIBUTE_VALUE_DATE", instances, (*A_ATTRIBUTE_VALUE_DATE)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ATTRIBUTE_VALUE_ENUMERATIONs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ATTRIBUTE_VALUE_ENUMERATION", instances, (*A_ATTRIBUTE_VALUE_ENUMERATION)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ATTRIBUTE_VALUE_INTEGERs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ATTRIBUTE_VALUE_INTEGER", instances, (*A_ATTRIBUTE_VALUE_INTEGER)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ATTRIBUTE_VALUE_REALs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ATTRIBUTE_VALUE_REAL", instances, (*A_ATTRIBUTE_VALUE_REAL)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ATTRIBUTE_VALUE_STRINGs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ATTRIBUTE_VALUE_STRING", instances, (*A_ATTRIBUTE_VALUE_STRING)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ATTRIBUTE_VALUE_XHTMLs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ATTRIBUTE_VALUE_XHTML", instances, (*A_ATTRIBUTE_VALUE_XHTML)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ATTRIBUTE_VALUE_XHTML_1s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ATTRIBUTE_VALUE_XHTML_1", instances, (*A_ATTRIBUTE_VALUE_XHTML_1)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_CHILDRENs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_CHILDREN", instances, (*A_CHILDREN)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_CORE_CONTENTs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_CORE_CONTENT", instances, (*A_CORE_CONTENT)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_DATATYPESs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_DATATYPES", instances, (*A_DATATYPES)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_DATATYPE_DEFINITION_BOOLEAN_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_DATATYPE_DEFINITION_BOOLEAN_REF", instances, (*A_DATATYPE_DEFINITION_BOOLEAN_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_DATATYPE_DEFINITION_DATE_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_DATATYPE_DEFINITION_DATE_REF", instances, (*A_DATATYPE_DEFINITION_DATE_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_DATATYPE_DEFINITION_ENUMERATION_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_DATATYPE_DEFINITION_ENUMERATION_REF", instances, (*A_DATATYPE_DEFINITION_ENUMERATION_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_DATATYPE_DEFINITION_INTEGER_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_DATATYPE_DEFINITION_INTEGER_REF", instances, (*A_DATATYPE_DEFINITION_INTEGER_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_DATATYPE_DEFINITION_REAL_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_DATATYPE_DEFINITION_REAL_REF", instances, (*A_DATATYPE_DEFINITION_REAL_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_DATATYPE_DEFINITION_STRING_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_DATATYPE_DEFINITION_STRING_REF", instances, (*A_DATATYPE_DEFINITION_STRING_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_DATATYPE_DEFINITION_XHTML_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_DATATYPE_DEFINITION_XHTML_REF", instances, (*A_DATATYPE_DEFINITION_XHTML_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_EDITABLE_ATTSs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_EDITABLE_ATTS", instances, (*A_EDITABLE_ATTS)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_ENUM_VALUE_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_ENUM_VALUE_REF", instances, (*A_ENUM_VALUE_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_OBJECTs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_OBJECT", instances, (*A_OBJECT)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_PROPERTIESs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_PROPERTIES", instances, (*A_PROPERTIES)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_RELATION_GROUP_TYPE_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_RELATION_GROUP_TYPE_REF", instances, (*A_RELATION_GROUP_TYPE_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_SOURCE_1s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_SOURCE_1", instances, (*A_SOURCE_1)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_SOURCE_SPECIFICATION_1s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_SOURCE_SPECIFICATION_1", instances, (*A_SOURCE_SPECIFICATION_1)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_SPECIFICATIONSs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_SPECIFICATIONS", instances, (*A_SPECIFICATIONS)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_SPECIFICATION_TYPE_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_SPECIFICATION_TYPE_REF", instances, (*A_SPECIFICATION_TYPE_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_SPECIFIED_VALUESs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_SPECIFIED_VALUES", instances, (*A_SPECIFIED_VALUES)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_SPEC_ATTRIBUTESs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_SPEC_ATTRIBUTES", instances, (*A_SPEC_ATTRIBUTES)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_SPEC_OBJECTSs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_SPEC_OBJECTS", instances, (*A_SPEC_OBJECTS)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_SPEC_OBJECT_TYPE_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_SPEC_OBJECT_TYPE_REF", instances, (*A_SPEC_OBJECT_TYPE_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_SPEC_RELATIONSs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_SPEC_RELATIONS", instances, (*A_SPEC_RELATIONS)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_SPEC_RELATION_GROUPSs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_SPEC_RELATION_GROUPS", instances, (*A_SPEC_RELATION_GROUPS)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_SPEC_RELATION_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_SPEC_RELATION_REF", instances, (*A_SPEC_RELATION_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_SPEC_RELATION_TYPE_REFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_SPEC_RELATION_TYPE_REF", instances, (*A_SPEC_RELATION_TYPE_REF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_SPEC_TYPESs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_SPEC_TYPES", instances, (*A_SPEC_TYPES)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_THE_HEADERs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_THE_HEADER", instances, (*A_THE_HEADER)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_TOOL_EXTENSIONSs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_TOOL_EXTENSIONS", instances, (*A_TOOL_EXTENSIONS)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DATATYPE_DEFINITION_BOOLEANs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DATATYPE_DEFINITION_BOOLEAN", instances, (*DATATYPE_DEFINITION_BOOLEAN)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DATATYPE_DEFINITION_DATEs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DATATYPE_DEFINITION_DATE", instances, (*DATATYPE_DEFINITION_DATE)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DATATYPE_DEFINITION_ENUMERATIONs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DATATYPE_DEFINITION_ENUMERATION", instances, (*DATATYPE_DEFINITION_ENUMERATION)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DATATYPE_DEFINITION_INTEGERs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DATATYPE_DEFINITION_INTEGER", instances, (*DATATYPE_DEFINITION_INTEGER)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DATATYPE_DEFINITION_REALs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DATATYPE_DEFINITION_REAL", instances, (*DATATYPE_DEFINITION_REAL)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DATATYPE_DEFINITION_STRINGs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DATATYPE_DEFINITION_STRING", instances, (*DATATYPE_DEFINITION_STRING)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.DATATYPE_DEFINITION_XHTMLs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "DATATYPE_DEFINITION_XHTML", instances, (*DATATYPE_DEFINITION_XHTML)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EMBEDDED_VALUEs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EMBEDDED_VALUE", instances, (*EMBEDDED_VALUE)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.ENUM_VALUEs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "ENUM_VALUE", instances, (*ENUM_VALUE)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EmbeddedJpgImages {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EmbeddedJpgImage", instances, (*EmbeddedJpgImage)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EmbeddedPngImages {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EmbeddedPngImage", instances, (*EmbeddedPngImage)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.EmbeddedSvgImages {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "EmbeddedSvgImage", instances, (*EmbeddedSvgImage)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Kills {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Kill", instances, (*Kill)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Map_identifier_bools {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Map_identifier_bool", instances, (*Map_identifier_bool)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RELATION_GROUPs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RELATION_GROUP", instances, (*RELATION_GROUP)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.RELATION_GROUP_TYPEs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "RELATION_GROUP_TYPE", instances, (*RELATION_GROUP_TYPE)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.REQ_IFs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "REQ_IF", instances, (*REQ_IF)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.REQ_IF_CONTENTs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "REQ_IF_CONTENT", instances, (*REQ_IF_CONTENT)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.REQ_IF_HEADERs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "REQ_IF_HEADER", instances, (*REQ_IF_HEADER)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.REQ_IF_TOOL_EXTENSIONs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "REQ_IF_TOOL_EXTENSION", instances, (*REQ_IF_TOOL_EXTENSION)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SPECIFICATIONs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SPECIFICATION", instances, (*SPECIFICATION)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SPECIFICATION_Renderings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SPECIFICATION_Rendering", instances, (*SPECIFICATION_Rendering)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SPECIFICATION_TYPEs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SPECIFICATION_TYPE", instances, (*SPECIFICATION_TYPE)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SPEC_HIERARCHYs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SPEC_HIERARCHY", instances, (*SPEC_HIERARCHY)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SPEC_OBJECTs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SPEC_OBJECT", instances, (*SPEC_OBJECT)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SPEC_OBJECT_TYPEs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SPEC_OBJECT_TYPE", instances, (*SPEC_OBJECT_TYPE)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SPEC_OBJECT_TYPE_Renderings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SPEC_OBJECT_TYPE_Rendering", instances, (*SPEC_OBJECT_TYPE_Rendering)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SPEC_RELATIONs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SPEC_RELATION", instances, (*SPEC_RELATION)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.SPEC_RELATION_TYPEs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "SPEC_RELATION_TYPE", instances, (*SPEC_RELATION_TYPE)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StaticWebSites {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StaticWebSite", instances, (*StaticWebSite)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StaticWebSiteChapters {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StaticWebSiteChapter", instances, (*StaticWebSiteChapter)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StaticWebSiteGeneratedImages {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StaticWebSiteGeneratedImage", instances, (*StaticWebSiteGeneratedImage)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StaticWebSiteImages {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StaticWebSiteImage", instances, (*StaticWebSiteImage)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.StaticWebSiteParagraphs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "StaticWebSiteParagraph", instances, (*StaticWebSiteParagraph)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.XHTML_CONTENTs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "XHTML_CONTENT", instances, (*XHTML_CONTENT)(nil).GongGetFieldHeaders(), addIDs)
		}
	}

	// Create a style with wrap text enabled
	wrapStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	})
	_ = wrapStyle
	if err != nil {
		fmt.Println("failed to create style:", err)
		return f
	}

	// Create a style with bold text
	boldStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})
	_ = boldStyle
	if err != nil {
		fmt.Println("failed to create bold style:", err)
		return f
	}

	// Get all sheet names
	sheetList := f.GetSheetList()

	for _, sheet := range sheetList {
		// Use a lazy iterator instead of loading all rows into memory
		rows, err := f.Rows(sheet)
		if err != nil {
			fmt.Printf("failed to get rows iterator for sheet %q: %v\n", sheet, err)
			continue
		}

		// Check if there is at least one row, and move the iterator to it
		if !rows.Next() {
			rows.Close() // Always close iterators
			continue
		}

		// Read ONLY the first row
		firstRow, err := rows.Columns()

		// Close the iterator immediately since we don't need the rest of the sheet
		rows.Close()

		if err != nil {
			fmt.Printf("failed to get columns for sheet %q: %v\n", sheet, err)
			continue
		}

		// If the first row is completely empty, skip
		if len(firstRow) == 0 {
			continue
		}

		// Track the first and last “used” column in the first row,
		// so we can later apply an AutoFilter from the first to last used col
		var firstUsedColIdx, lastUsedColIdx int

		for colIdx, cellValue := range firstRow {
			if cellValue == "" {
				// Skip columns with empty first-row cells
				continue
			}

			// Convert zero-based colIdx to 1-based for Excelize,
			// then get the column name (A, B, C, etc.)
			colName, err := excelize.ColumnNumberToName(colIdx + 1)
			if err != nil {
				fmt.Printf("failed to convert column number: %v\n", err)
				continue
			}

			// Apply wrap-text style to this entire column
			colRange := colName + ":" + colName
			if err := f.SetColStyle(sheet, colRange, wrapStyle); err != nil {
				fmt.Printf("failed to set col style on %s: %v\n", colRange, err)
				continue
			}

			// Make the first row (cell in row 1) bold in this column
			cellRef := fmt.Sprintf("%s1", colName)
			if err := f.SetCellStyle(sheet, cellRef, cellRef, boldStyle); err != nil {
				fmt.Printf("failed to set cell style on %s: %v\n", cellRef, err)
				continue
			}

			// Update our “first used” and “last used” column indices
			if firstUsedColIdx == 0 {
				firstUsedColIdx = colIdx + 1
			}
			if colIdx+1 > lastUsedColIdx {
				lastUsedColIdx = colIdx + 1
			}
		}

		// If we found at least one non-empty column in row 1, enable AutoFilter
		if firstUsedColIdx != 0 && lastUsedColIdx >= firstUsedColIdx {
			startCol, _ := excelize.ColumnNumberToName(firstUsedColIdx)
			endCol, _ := excelize.ColumnNumberToName(lastUsedColIdx)
			styleRange := fmt.Sprintf("%s:%s", startCol, endCol)
			autoFilterRange := fmt.Sprintf("%s1:%s1", startCol, endCol)
			startCellString := fmt.Sprintf("%s1", startCol)
			endCellString := fmt.Sprintf("%s1", endCol)

			if err := f.SetColStyle(sheet, styleRange, wrapStyle); err != nil {
				fmt.Println("failed to set column style:", err)
				return f
			}

			// Apply the bold style to the first row (A1:XFD1)
			if err := f.SetCellStyle(sheet, startCellString, endCellString, boldStyle); err != nil {
				fmt.Println("failed to set bold style:", err)
				return f
			}

			var opts []excelize.AutoFilterOptions
			if err := f.AutoFilter(sheet, autoFilterRange, opts); err != nil {
				fmt.Printf("failed to enable auto filter on range %s: %v\n", autoFilterRange, err)
			}
		}
	}

	var tab ExcelizeTabulator
	tab.SetExcelizeFile(f)
	{
		f.DeleteSheet("Sheet1")
	}
	return f
}

// SerializeStageAsBytes serializes the stage to a pure in-memory Excel file and returns the bytes.
func (stage *Stage) SerializeStageAsBytes(addIDs bool) ([]byte, error) {
	f := stage.__gong__buildExcelizeFile(addIDs)
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func __gong__shortenString(s string) string {
	if len(s) > 31 {
		return s[:31]
	}
	return s
}

// Tabulator is an interface for writing to a table strings
type Tabulator interface {
	AddSheet(sheetName string)
	AddRow(sheetName string) int
	AddCell(sheetName string, rowId, columnIndex int, value string)
}

type ExcelizeTabulator struct {
	f *excelize.File
}

func (tab *ExcelizeTabulator) SetExcelizeFile(f *excelize.File) {
	tab.f = f
}

func (tab *ExcelizeTabulator) AddSheet(sheetName string) {

}

func (tab *ExcelizeTabulator) AddRow(sheetName string) (rowId int) {
	return
}

func (tab *ExcelizeTabulator) AddCell(sheetName string, rowId, columnIndex int, value string) {

}

// SerializeExcelize is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelize(f *excelize.File, name string, instances []GongstructIF, fields []GongFieldHeader, addIDs bool) {
	sheetName := __gong__shortenString(name)

	// Create a new sheet.
	f.NewSheet(sheetName)

	sortedSlice := make([]GongstructIF, len(instances))
	copy(sortedSlice, instances)
	slices.SortFunc(sortedSlice, func(a, b GongstructIF) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	line := 1

	for index, fieldHeader := range fields {
		if !addIDs {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(index+1)), line), fieldHeader.Name)
		} else {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+1)), line), fieldHeader.Name)
			switch fieldHeader.GongFieldValueType {
			case GongFieldValueTypePointer:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":ID")
			case GongFieldValueTypeSliceOfPointers:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":IDs")
			default:
				// if index is 0, this is the ID of the instance
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), "ID")
				} else {
					// one have to put the type of the cell
					header := fieldHeader.Name
					switch fieldHeader.GongFieldValueType {
					case GongFieldValueTypeInt:
						header += ":int"
					case GongFieldValueTypeIntDuration:
						header += ":duration"
					case GongFieldValueTypeFloat:
						header += ":float"
					case GongFieldValueTypeBool:
						header += ":bool"
					case GongFieldValueTypeString:
						header += ":string"
					case GongFieldValueTypeDate:
						header += ":date"
					default:
						header += ":basicType"
					}
					header += ":noID"
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), header)
				}
			}
		}
	}

	// AutoFilter starting from A1
	f.AutoFilter(sheetName,
		fmt.Sprintf("%s%d", GongIntToLetters(int32(1)), line),
		[]excelize.AutoFilterOptions{})

	for _, instance := range sortedSlice {
		line = line + 1

		// 3. Add the ID value in column B

		for index, fieldName := range fields {
			fieldStringValue := stage.GetFieldStringValueFromPointer(instance, fieldName.Name)
			if !addIDs {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(index+1)), line), fieldStringValue.GetValueString())
			} else {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+1)), line), fieldStringValue.GetValueString())
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), instance.GongGetUUID(stage))
				} else {
					switch fieldStringValue.GongFieldValueType {
					case GongFieldValueTypePointer, GongFieldValueTypeSliceOfPointers:
						f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), fieldStringValue.ids)
					}
				}

			}
		}
	}
}

// SerializeExcelizePointer is the Stage method for Excel serialization.
func (stage *Stage) SerializeExcelizePointer[Type PointerToGongstruct](f *excelize.File) {
	stage.SerializeExcelizePointer2[Type](f, false)
}

// SerializeExcelizePointer2 is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelizePointer2[Type PointerToGongstruct](f *excelize.File, addIDs bool) {
	var ret Type
	set := *stage.GetInstancesSet[Type]()
	var instances []GongstructIF
	for key := range set {
		instances = append(instances, key)
	}
	stage.SerializeExcelize(f, ret.GongGetGongstructName(), instances, ret.GongGetFieldHeaders(), addIDs)
}
