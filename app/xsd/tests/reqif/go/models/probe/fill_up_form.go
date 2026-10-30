// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/reqif/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/reqif/go/models/xmlxsd"
)

const FormName = "Form"

var (
	_ = sort.Strings
	_ = strings.Join
)

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {
	switch inst := instance.(type) {
	case *models.ALTERNATIVE_ID:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_ALTERNATIVE_IDs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_ALTERNATIVE_ID", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}
	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IS_EDITABLE", inst.IS_EDITABLE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("DEFAULT_VALUE", inst.DEFAULT_VALUE, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_BOOLEAN](), probe.formStage)
		StageSetAssociationFieldToForm("TYPE_XSD", inst.TYPE_XSD, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_DATATYPE_DEFINITION_BOOLEAN_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_SPEC_ATTRIBUTESs {
				for _, target := range src.ATTRIBUTE_DEFINITION_BOOLEAN {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_SPEC_ATTRIBUTES", "ATTRIBUTE_DEFINITION_BOOLEAN", refNames, formGroup, probe.formStage)
		}
	case *models.ATTRIBUTE_DEFINITION_DATE:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IS_EDITABLE", inst.IS_EDITABLE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("DEFAULT_VALUE", inst.DEFAULT_VALUE, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_DATE](), probe.formStage)
		StageSetAssociationFieldToForm("TYPE_XSD", inst.TYPE_XSD, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_DATATYPE_DEFINITION_DATE_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_SPEC_ATTRIBUTESs {
				for _, target := range src.ATTRIBUTE_DEFINITION_DATE {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_SPEC_ATTRIBUTES", "ATTRIBUTE_DEFINITION_DATE", refNames, formGroup, probe.formStage)
		}
	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IS_EDITABLE", inst.IS_EDITABLE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MULTI_VALUED", inst.MULTI_VALUED, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("DEFAULT_VALUE", inst.DEFAULT_VALUE, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_ENUMERATION](), probe.formStage)
		StageSetAssociationFieldToForm("TYPE_XSD", inst.TYPE_XSD, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_DATATYPE_DEFINITION_ENUMERATION_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_SPEC_ATTRIBUTESs {
				for _, target := range src.ATTRIBUTE_DEFINITION_ENUMERATION {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_SPEC_ATTRIBUTES", "ATTRIBUTE_DEFINITION_ENUMERATION", refNames, formGroup, probe.formStage)
		}
	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IS_EDITABLE", inst.IS_EDITABLE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("DEFAULT_VALUE", inst.DEFAULT_VALUE, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_INTEGER](), probe.formStage)
		StageSetAssociationFieldToForm("TYPE_XSD", inst.TYPE_XSD, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_DATATYPE_DEFINITION_INTEGER_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_SPEC_ATTRIBUTESs {
				for _, target := range src.ATTRIBUTE_DEFINITION_INTEGER {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_SPEC_ATTRIBUTES", "ATTRIBUTE_DEFINITION_INTEGER", refNames, formGroup, probe.formStage)
		}
	case *models.ATTRIBUTE_DEFINITION_REAL:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IS_EDITABLE", inst.IS_EDITABLE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("DEFAULT_VALUE", inst.DEFAULT_VALUE, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_REAL](), probe.formStage)
		StageSetAssociationFieldToForm("TYPE_XSD", inst.TYPE_XSD, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_DATATYPE_DEFINITION_REAL_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_SPEC_ATTRIBUTESs {
				for _, target := range src.ATTRIBUTE_DEFINITION_REAL {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_SPEC_ATTRIBUTES", "ATTRIBUTE_DEFINITION_REAL", refNames, formGroup, probe.formStage)
		}
	case *models.ATTRIBUTE_DEFINITION_STRING:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IS_EDITABLE", inst.IS_EDITABLE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("DEFAULT_VALUE", inst.DEFAULT_VALUE, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_STRING](), probe.formStage)
		StageSetAssociationFieldToForm("TYPE_XSD", inst.TYPE_XSD, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_DATATYPE_DEFINITION_STRING_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_SPEC_ATTRIBUTESs {
				for _, target := range src.ATTRIBUTE_DEFINITION_STRING {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_SPEC_ATTRIBUTES", "ATTRIBUTE_DEFINITION_STRING", refNames, formGroup, probe.formStage)
		}
	case *models.ATTRIBUTE_DEFINITION_XHTML:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IS_EDITABLE", inst.IS_EDITABLE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("DEFAULT_VALUE", inst.DEFAULT_VALUE, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_XHTML](), probe.formStage)
		StageSetAssociationFieldToForm("TYPE_XSD", inst.TYPE_XSD, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_DATATYPE_DEFINITION_XHTML_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_SPEC_ATTRIBUTESs {
				for _, target := range src.ATTRIBUTE_DEFINITION_XHTML {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_SPEC_ATTRIBUTES", "ATTRIBUTE_DEFINITION_XHTML", refNames, formGroup, probe.formStage)
		}
	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("THE_VALUE", inst.THE_VALUE, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("DEFINITION", inst.DEFINITION, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_ATTRIBUTE_VALUE_BOOLEANs {
				for _, target := range src.ATTRIBUTE_VALUE_BOOLEAN {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_ATTRIBUTE_VALUE_BOOLEAN", "ATTRIBUTE_VALUE_BOOLEAN", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_ATTRIBUTE_VALUE_XHTML_1s {
				for _, target := range src.ATTRIBUTE_VALUE_BOOLEAN {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_ATTRIBUTE_VALUE_XHTML_1", "ATTRIBUTE_VALUE_BOOLEAN", refNames, formGroup, probe.formStage)
		}
	case *models.ATTRIBUTE_VALUE_DATE:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("THE_VALUE", inst.THE_VALUE, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("DEFINITION", inst.DEFINITION, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_DATE_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_ATTRIBUTE_VALUE_DATEs {
				for _, target := range src.ATTRIBUTE_VALUE_DATE {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_ATTRIBUTE_VALUE_DATE", "ATTRIBUTE_VALUE_DATE", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_ATTRIBUTE_VALUE_XHTML_1s {
				for _, target := range src.ATTRIBUTE_VALUE_DATE {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_ATTRIBUTE_VALUE_XHTML_1", "ATTRIBUTE_VALUE_DATE", refNames, formGroup, probe.formStage)
		}
	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("DEFINITION", inst.DEFINITION, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF](), probe.formStage)
		StageSetAssociationFieldToForm("VALUES", inst.VALUES, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ENUM_VALUE_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_ATTRIBUTE_VALUE_ENUMERATIONs {
				for _, target := range src.ATTRIBUTE_VALUE_ENUMERATION {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_ATTRIBUTE_VALUE_ENUMERATION", "ATTRIBUTE_VALUE_ENUMERATION", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_ATTRIBUTE_VALUE_XHTML_1s {
				for _, target := range src.ATTRIBUTE_VALUE_ENUMERATION {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_ATTRIBUTE_VALUE_XHTML_1", "ATTRIBUTE_VALUE_ENUMERATION", refNames, formGroup, probe.formStage)
		}
	case *models.ATTRIBUTE_VALUE_INTEGER:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("THE_VALUE", inst.THE_VALUE, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("DEFINITION", inst.DEFINITION, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_INTEGER_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_ATTRIBUTE_VALUE_INTEGERs {
				for _, target := range src.ATTRIBUTE_VALUE_INTEGER {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_ATTRIBUTE_VALUE_INTEGER", "ATTRIBUTE_VALUE_INTEGER", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_ATTRIBUTE_VALUE_XHTML_1s {
				for _, target := range src.ATTRIBUTE_VALUE_INTEGER {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_ATTRIBUTE_VALUE_XHTML_1", "ATTRIBUTE_VALUE_INTEGER", refNames, formGroup, probe.formStage)
		}
	case *models.ATTRIBUTE_VALUE_REAL:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("THE_VALUE", inst.THE_VALUE, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("DEFINITION", inst.DEFINITION, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_REAL_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_ATTRIBUTE_VALUE_REALs {
				for _, target := range src.ATTRIBUTE_VALUE_REAL {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_ATTRIBUTE_VALUE_REAL", "ATTRIBUTE_VALUE_REAL", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_ATTRIBUTE_VALUE_XHTML_1s {
				for _, target := range src.ATTRIBUTE_VALUE_REAL {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_ATTRIBUTE_VALUE_XHTML_1", "ATTRIBUTE_VALUE_REAL", refNames, formGroup, probe.formStage)
		}
	case *models.ATTRIBUTE_VALUE_STRING:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("THE_VALUE", inst.THE_VALUE, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("DEFINITION", inst.DEFINITION, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_STRING_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_ATTRIBUTE_VALUE_STRINGs {
				for _, target := range src.ATTRIBUTE_VALUE_STRING {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_ATTRIBUTE_VALUE_STRING", "ATTRIBUTE_VALUE_STRING", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_ATTRIBUTE_VALUE_XHTML_1s {
				for _, target := range src.ATTRIBUTE_VALUE_STRING {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_ATTRIBUTE_VALUE_XHTML_1", "ATTRIBUTE_VALUE_STRING", refNames, formGroup, probe.formStage)
		}
	case *models.ATTRIBUTE_VALUE_XHTML:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IS_SIMPLIFIED", inst.IS_SIMPLIFIED, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("THE_VALUE", inst.THE_VALUE, formGroup, probe.stageOfInterest.GetInstancesSet[*models.XHTML_CONTENT](), probe.formStage)
		StageSetAssociationFieldToForm("THE_ORIGINAL_VALUE", inst.THE_ORIGINAL_VALUE, formGroup, probe.stageOfInterest.GetInstancesSet[*models.XHTML_CONTENT](), probe.formStage)
		StageSetAssociationFieldToForm("DEFINITION", inst.DEFINITION, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_DEFINITION_XHTML_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_ATTRIBUTE_VALUE_XHTMLs {
				for _, target := range src.ATTRIBUTE_VALUE_XHTML {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_ATTRIBUTE_VALUE_XHTML", "ATTRIBUTE_VALUE_XHTML", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_ATTRIBUTE_VALUE_XHTML_1s {
				for _, target := range src.ATTRIBUTE_VALUE_XHTML {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_ATTRIBUTE_VALUE_XHTML_1", "ATTRIBUTE_VALUE_XHTML", refNames, formGroup, probe.formStage)
		}
	case *models.A_ALTERNATIVE_ID:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.ALTERNATIVE_ID](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_BOOLEANs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_BOOLEAN", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_DATEs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_DATE", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_ENUMERATIONs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_ENUMERATION", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_INTEGERs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_INTEGER", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_REALs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_REAL", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_STRINGs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_STRING", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_XHTMLs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_XHTML", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.DATATYPE_DEFINITION_BOOLEANs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DATATYPE_DEFINITION_BOOLEAN", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.DATATYPE_DEFINITION_DATEs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DATATYPE_DEFINITION_DATE", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.DATATYPE_DEFINITION_ENUMERATIONs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DATATYPE_DEFINITION_ENUMERATION", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.DATATYPE_DEFINITION_INTEGERs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DATATYPE_DEFINITION_INTEGER", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.DATATYPE_DEFINITION_REALs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DATATYPE_DEFINITION_REAL", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.DATATYPE_DEFINITION_STRINGs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DATATYPE_DEFINITION_STRING", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.DATATYPE_DEFINITION_XHTMLs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DATATYPE_DEFINITION_XHTML", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.ENUM_VALUEs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ENUM_VALUE", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.RELATION_GROUPs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.RELATION_GROUP", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.RELATION_GROUP_TYPEs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.RELATION_GROUP_TYPE", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPECIFICATIONs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPECIFICATION", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPECIFICATION_TYPEs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPECIFICATION_TYPE", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPEC_HIERARCHYs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPEC_HIERARCHY", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPEC_OBJECTs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPEC_OBJECT", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPEC_OBJECT_TYPEs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPEC_OBJECT_TYPE", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPEC_RELATIONs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPEC_RELATION", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPEC_RELATION_TYPEs {
				if src.ALTERNATIVE_ID == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPEC_RELATION_TYPE", "ALTERNATIVE_ID", refNames, formGroup, probe.formStage)
		}
	case *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ATTRIBUTE_DEFINITION_BOOLEAN_REF", inst.ATTRIBUTE_DEFINITION_BOOLEAN_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_VALUE_BOOLEANs {
				if src.DEFINITION == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_VALUE_BOOLEAN", "DEFINITION", refNames, formGroup, probe.formStage)
		}
	case *models.A_ATTRIBUTE_DEFINITION_DATE_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ATTRIBUTE_DEFINITION_DATE_REF", inst.ATTRIBUTE_DEFINITION_DATE_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_VALUE_DATEs {
				if src.DEFINITION == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_VALUE_DATE", "DEFINITION", refNames, formGroup, probe.formStage)
		}
	case *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ATTRIBUTE_DEFINITION_ENUMERATION_REF", inst.ATTRIBUTE_DEFINITION_ENUMERATION_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_VALUE_ENUMERATIONs {
				if src.DEFINITION == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_VALUE_ENUMERATION", "DEFINITION", refNames, formGroup, probe.formStage)
		}
	case *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ATTRIBUTE_DEFINITION_INTEGER_REF", inst.ATTRIBUTE_DEFINITION_INTEGER_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_VALUE_INTEGERs {
				if src.DEFINITION == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_VALUE_INTEGER", "DEFINITION", refNames, formGroup, probe.formStage)
		}
	case *models.A_ATTRIBUTE_DEFINITION_REAL_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ATTRIBUTE_DEFINITION_REAL_REF", inst.ATTRIBUTE_DEFINITION_REAL_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_VALUE_REALs {
				if src.DEFINITION == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_VALUE_REAL", "DEFINITION", refNames, formGroup, probe.formStage)
		}
	case *models.A_ATTRIBUTE_DEFINITION_STRING_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ATTRIBUTE_DEFINITION_STRING_REF", inst.ATTRIBUTE_DEFINITION_STRING_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_VALUE_STRINGs {
				if src.DEFINITION == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_VALUE_STRING", "DEFINITION", refNames, formGroup, probe.formStage)
		}
	case *models.A_ATTRIBUTE_DEFINITION_XHTML_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ATTRIBUTE_DEFINITION_XHTML_REF", inst.ATTRIBUTE_DEFINITION_XHTML_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_VALUE_XHTMLs {
				if src.DEFINITION == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_VALUE_XHTML", "DEFINITION", refNames, formGroup, probe.formStage)
		}
	case *models.A_ATTRIBUTE_VALUE_BOOLEAN:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("ATTRIBUTE_VALUE_BOOLEAN", inst, &inst.ATTRIBUTE_VALUE_BOOLEAN, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_BOOLEANs {
				if src.DEFAULT_VALUE == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_BOOLEAN", "DEFAULT_VALUE", refNames, formGroup, probe.formStage)
		}
	case *models.A_ATTRIBUTE_VALUE_DATE:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("ATTRIBUTE_VALUE_DATE", inst, &inst.ATTRIBUTE_VALUE_DATE, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_DATEs {
				if src.DEFAULT_VALUE == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_DATE", "DEFAULT_VALUE", refNames, formGroup, probe.formStage)
		}
	case *models.A_ATTRIBUTE_VALUE_ENUMERATION:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("ATTRIBUTE_VALUE_ENUMERATION", inst, &inst.ATTRIBUTE_VALUE_ENUMERATION, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_ENUMERATIONs {
				if src.DEFAULT_VALUE == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_ENUMERATION", "DEFAULT_VALUE", refNames, formGroup, probe.formStage)
		}
	case *models.A_ATTRIBUTE_VALUE_INTEGER:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("ATTRIBUTE_VALUE_INTEGER", inst, &inst.ATTRIBUTE_VALUE_INTEGER, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_INTEGERs {
				if src.DEFAULT_VALUE == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_INTEGER", "DEFAULT_VALUE", refNames, formGroup, probe.formStage)
		}
	case *models.A_ATTRIBUTE_VALUE_REAL:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("ATTRIBUTE_VALUE_REAL", inst, &inst.ATTRIBUTE_VALUE_REAL, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_REALs {
				if src.DEFAULT_VALUE == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_REAL", "DEFAULT_VALUE", refNames, formGroup, probe.formStage)
		}
	case *models.A_ATTRIBUTE_VALUE_STRING:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("ATTRIBUTE_VALUE_STRING", inst, &inst.ATTRIBUTE_VALUE_STRING, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_STRINGs {
				if src.DEFAULT_VALUE == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_STRING", "DEFAULT_VALUE", refNames, formGroup, probe.formStage)
		}
	case *models.A_ATTRIBUTE_VALUE_XHTML:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("ATTRIBUTE_VALUE_XHTML", inst, &inst.ATTRIBUTE_VALUE_XHTML, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_XHTMLs {
				if src.DEFAULT_VALUE == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_XHTML", "DEFAULT_VALUE", refNames, formGroup, probe.formStage)
		}
	case *models.A_ATTRIBUTE_VALUE_XHTML_1:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("ATTRIBUTE_VALUE_BOOLEAN", inst, &inst.ATTRIBUTE_VALUE_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_DATE", inst, &inst.ATTRIBUTE_VALUE_DATE, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_ENUMERATION", inst, &inst.ATTRIBUTE_VALUE_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_INTEGER", inst, &inst.ATTRIBUTE_VALUE_INTEGER, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_REAL", inst, &inst.ATTRIBUTE_VALUE_REAL, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_STRING", inst, &inst.ATTRIBUTE_VALUE_STRING, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_XHTML", inst, &inst.ATTRIBUTE_VALUE_XHTML, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPECIFICATIONs {
				if src.VALUES == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPECIFICATION", "VALUES", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPEC_OBJECTs {
				if src.VALUES == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPEC_OBJECT", "VALUES", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPEC_RELATIONs {
				if src.VALUES == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPEC_RELATION", "VALUES", refNames, formGroup, probe.formStage)
		}
	case *models.A_CHILDREN:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("SPEC_HIERARCHY", inst, &inst.SPEC_HIERARCHY, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPECIFICATIONs {
				if src.CHILDREN == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPECIFICATION", "CHILDREN", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPEC_HIERARCHYs {
				if src.CHILDREN == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPEC_HIERARCHY", "CHILDREN", refNames, formGroup, probe.formStage)
		}
	case *models.A_CORE_CONTENT:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("REQ_IF_CONTENT", inst.REQ_IF_CONTENT, formGroup, probe.stageOfInterest.GetInstancesSet[*models.REQ_IF_CONTENT](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.REQ_IFs {
				if src.CORE_CONTENT == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.REQ_IF", "CORE_CONTENT", refNames, formGroup, probe.formStage)
		}
	case *models.A_DATATYPES:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("DATATYPE_DEFINITION_BOOLEAN", inst, &inst.DATATYPE_DEFINITION_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_DATE", inst, &inst.DATATYPE_DEFINITION_DATE, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_ENUMERATION", inst, &inst.DATATYPE_DEFINITION_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_INTEGER", inst, &inst.DATATYPE_DEFINITION_INTEGER, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_REAL", inst, &inst.DATATYPE_DEFINITION_REAL, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_STRING", inst, &inst.DATATYPE_DEFINITION_STRING, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_XHTML", inst, &inst.DATATYPE_DEFINITION_XHTML, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.REQ_IF_CONTENTs {
				if src.DATATYPES == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.REQ_IF_CONTENT", "DATATYPES", refNames, formGroup, probe.formStage)
		}
	case *models.A_DATATYPE_DEFINITION_BOOLEAN_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DATATYPE_DEFINITION_BOOLEAN_REF", inst.DATATYPE_DEFINITION_BOOLEAN_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_BOOLEANs {
				if src.TYPE_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_BOOLEAN", "TYPE_XSD", refNames, formGroup, probe.formStage)
		}
	case *models.A_DATATYPE_DEFINITION_DATE_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DATATYPE_DEFINITION_DATE_REF", inst.DATATYPE_DEFINITION_DATE_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_DATEs {
				if src.TYPE_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_DATE", "TYPE_XSD", refNames, formGroup, probe.formStage)
		}
	case *models.A_DATATYPE_DEFINITION_ENUMERATION_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DATATYPE_DEFINITION_ENUMERATION_REF", inst.DATATYPE_DEFINITION_ENUMERATION_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_ENUMERATIONs {
				if src.TYPE_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_ENUMERATION", "TYPE_XSD", refNames, formGroup, probe.formStage)
		}
	case *models.A_DATATYPE_DEFINITION_INTEGER_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DATATYPE_DEFINITION_INTEGER_REF", inst.DATATYPE_DEFINITION_INTEGER_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_INTEGERs {
				if src.TYPE_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_INTEGER", "TYPE_XSD", refNames, formGroup, probe.formStage)
		}
	case *models.A_DATATYPE_DEFINITION_REAL_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DATATYPE_DEFINITION_REAL_REF", inst.DATATYPE_DEFINITION_REAL_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_REALs {
				if src.TYPE_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_REAL", "TYPE_XSD", refNames, formGroup, probe.formStage)
		}
	case *models.A_DATATYPE_DEFINITION_STRING_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DATATYPE_DEFINITION_STRING_REF", inst.DATATYPE_DEFINITION_STRING_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_STRINGs {
				if src.TYPE_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_STRING", "TYPE_XSD", refNames, formGroup, probe.formStage)
		}
	case *models.A_DATATYPE_DEFINITION_XHTML_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DATATYPE_DEFINITION_XHTML_REF", inst.DATATYPE_DEFINITION_XHTML_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_DEFINITION_XHTMLs {
				if src.TYPE_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_DEFINITION_XHTML", "TYPE_XSD", refNames, formGroup, probe.formStage)
		}
	case *models.A_EDITABLE_ATTS:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ATTRIBUTE_DEFINITION_BOOLEAN_REF", inst.ATTRIBUTE_DEFINITION_BOOLEAN_REF, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ATTRIBUTE_DEFINITION_DATE_REF", inst.ATTRIBUTE_DEFINITION_DATE_REF, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ATTRIBUTE_DEFINITION_ENUMERATION_REF", inst.ATTRIBUTE_DEFINITION_ENUMERATION_REF, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ATTRIBUTE_DEFINITION_INTEGER_REF", inst.ATTRIBUTE_DEFINITION_INTEGER_REF, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ATTRIBUTE_DEFINITION_REAL_REF", inst.ATTRIBUTE_DEFINITION_REAL_REF, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ATTRIBUTE_DEFINITION_STRING_REF", inst.ATTRIBUTE_DEFINITION_STRING_REF, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ATTRIBUTE_DEFINITION_XHTML_REF", inst.ATTRIBUTE_DEFINITION_XHTML_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPEC_HIERARCHYs {
				if src.EDITABLE_ATTS == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPEC_HIERARCHY", "EDITABLE_ATTS", refNames, formGroup, probe.formStage)
		}
	case *models.A_ENUM_VALUE_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ENUM_VALUE_REF", inst.ENUM_VALUE_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_VALUE_ENUMERATIONs {
				if src.VALUES == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_VALUE_ENUMERATION", "VALUES", refNames, formGroup, probe.formStage)
		}
	case *models.A_OBJECT:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SPEC_OBJECT_REF", inst.SPEC_OBJECT_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPEC_HIERARCHYs {
				if src.OBJECT == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPEC_HIERARCHY", "OBJECT", refNames, formGroup, probe.formStage)
		}
	case *models.A_PROPERTIES:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("EMBEDDED_VALUE", inst.EMBEDDED_VALUE, formGroup, probe.stageOfInterest.GetInstancesSet[*models.EMBEDDED_VALUE](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ENUM_VALUEs {
				if src.PROPERTIES == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ENUM_VALUE", "PROPERTIES", refNames, formGroup, probe.formStage)
		}
	case *models.A_RELATION_GROUP_TYPE_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("RELATION_GROUP_TYPE_REF", inst.RELATION_GROUP_TYPE_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.RELATION_GROUPs {
				if src.TYPE_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.RELATION_GROUP", "TYPE_XSD", refNames, formGroup, probe.formStage)
		}
	case *models.A_SOURCE:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SPEC_OBJECT_REF", inst.SPEC_OBJECT_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPEC_RELATIONs {
				if src.SOURCE == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPEC_RELATION", "SOURCE", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPEC_RELATIONs {
				if src.TARGET == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPEC_RELATION", "TARGET", refNames, formGroup, probe.formStage)
		}
	case *models.A_SOURCE_SPECIFICATION:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SPECIFICATION_REF", inst.SPECIFICATION_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.RELATION_GROUPs {
				if src.SOURCE_SPECIFICATION == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.RELATION_GROUP", "SOURCE_SPECIFICATION", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.RELATION_GROUPs {
				if src.TARGET_SPECIFICATION == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.RELATION_GROUP", "TARGET_SPECIFICATION", refNames, formGroup, probe.formStage)
		}
	case *models.A_SPECIFICATIONS:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("SPECIFICATION", inst, &inst.SPECIFICATION, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.REQ_IF_CONTENTs {
				if src.SPECIFICATIONS == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.REQ_IF_CONTENT", "SPECIFICATIONS", refNames, formGroup, probe.formStage)
		}
	case *models.A_SPECIFICATION_TYPE_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SPECIFICATION_TYPE_REF", inst.SPECIFICATION_TYPE_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPECIFICATIONs {
				if src.TYPE_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPECIFICATION", "TYPE_XSD", refNames, formGroup, probe.formStage)
		}
	case *models.A_SPECIFIED_VALUES:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("ENUM_VALUE", inst, &inst.ENUM_VALUE, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.DATATYPE_DEFINITION_ENUMERATIONs {
				if src.SPECIFIED_VALUES == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.DATATYPE_DEFINITION_ENUMERATION", "SPECIFIED_VALUES", refNames, formGroup, probe.formStage)
		}
	case *models.A_SPEC_ATTRIBUTES:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_BOOLEAN", inst, &inst.ATTRIBUTE_DEFINITION_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_DATE", inst, &inst.ATTRIBUTE_DEFINITION_DATE, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_ENUMERATION", inst, &inst.ATTRIBUTE_DEFINITION_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_INTEGER", inst, &inst.ATTRIBUTE_DEFINITION_INTEGER, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_REAL", inst, &inst.ATTRIBUTE_DEFINITION_REAL, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_STRING", inst, &inst.ATTRIBUTE_DEFINITION_STRING, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_XHTML", inst, &inst.ATTRIBUTE_DEFINITION_XHTML, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.RELATION_GROUP_TYPEs {
				if src.SPEC_ATTRIBUTES == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.RELATION_GROUP_TYPE", "SPEC_ATTRIBUTES", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPECIFICATION_TYPEs {
				if src.SPEC_ATTRIBUTES == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPECIFICATION_TYPE", "SPEC_ATTRIBUTES", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPEC_OBJECT_TYPEs {
				if src.SPEC_ATTRIBUTES == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPEC_OBJECT_TYPE", "SPEC_ATTRIBUTES", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPEC_RELATION_TYPEs {
				if src.SPEC_ATTRIBUTES == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPEC_RELATION_TYPE", "SPEC_ATTRIBUTES", refNames, formGroup, probe.formStage)
		}
	case *models.A_SPEC_OBJECTS:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("SPEC_OBJECT", inst, &inst.SPEC_OBJECT, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.REQ_IF_CONTENTs {
				if src.SPEC_OBJECTS == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.REQ_IF_CONTENT", "SPEC_OBJECTS", refNames, formGroup, probe.formStage)
		}
	case *models.A_SPEC_OBJECT_TYPE_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SPEC_OBJECT_TYPE_REF", inst.SPEC_OBJECT_TYPE_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPEC_OBJECTs {
				if src.TYPE_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPEC_OBJECT", "TYPE_XSD", refNames, formGroup, probe.formStage)
		}
	case *models.A_SPEC_RELATIONS:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("SPEC_RELATION", inst, &inst.SPEC_RELATION, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.REQ_IF_CONTENTs {
				if src.SPEC_RELATIONS == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.REQ_IF_CONTENT", "SPEC_RELATIONS", refNames, formGroup, probe.formStage)
		}
	case *models.A_SPEC_RELATION_GROUPS:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("RELATION_GROUP", inst, &inst.RELATION_GROUP, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.REQ_IF_CONTENTs {
				if src.SPEC_RELATION_GROUPS == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.REQ_IF_CONTENT", "SPEC_RELATION_GROUPS", refNames, formGroup, probe.formStage)
		}
	case *models.A_SPEC_RELATION_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SPEC_RELATION_REF", inst.SPEC_RELATION_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.RELATION_GROUPs {
				if src.SPEC_RELATIONS == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.RELATION_GROUP", "SPEC_RELATIONS", refNames, formGroup, probe.formStage)
		}
	case *models.A_SPEC_RELATION_TYPE_REF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SPEC_RELATION_TYPE_REF", inst.SPEC_RELATION_TYPE_REF, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.SPEC_RELATIONs {
				if src.TYPE_XSD == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.SPEC_RELATION", "TYPE_XSD", refNames, formGroup, probe.formStage)
		}
	case *models.A_SPEC_TYPES:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("RELATION_GROUP_TYPE", inst, &inst.RELATION_GROUP_TYPE, formGroup, probe)
		AssociationSliceToForm("SPEC_OBJECT_TYPE", inst, &inst.SPEC_OBJECT_TYPE, formGroup, probe)
		AssociationSliceToForm("SPEC_RELATION_TYPE", inst, &inst.SPEC_RELATION_TYPE, formGroup, probe)
		AssociationSliceToForm("SPECIFICATION_TYPE", inst, &inst.SPECIFICATION_TYPE, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.REQ_IF_CONTENTs {
				if src.SPEC_TYPES == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.REQ_IF_CONTENT", "SPEC_TYPES", refNames, formGroup, probe.formStage)
		}
	case *models.A_THE_HEADER:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("REQ_IF_HEADER", inst.REQ_IF_HEADER, formGroup, probe.stageOfInterest.GetInstancesSet[*models.REQ_IF_HEADER](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.REQ_IFs {
				if src.THE_HEADER == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.REQ_IF", "THE_HEADER", refNames, formGroup, probe.formStage)
		}
	case *models.A_TOOL_EXTENSIONS:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		AssociationSliceToForm("REQ_IF_TOOL_EXTENSION", inst, &inst.REQ_IF_TOOL_EXTENSION, formGroup, probe)

		{
			var refNames []string
			for src := range probe.stageOfInterest.REQ_IFs {
				if src.TOOL_EXTENSIONS == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.REQ_IF", "TOOL_EXTENSIONS", refNames, formGroup, probe.formStage)
		}
	case *models.DATATYPE_DEFINITION_BOOLEAN:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_DATATYPESs {
				for _, target := range src.DATATYPE_DEFINITION_BOOLEAN {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_DATATYPES", "DATATYPE_DEFINITION_BOOLEAN", refNames, formGroup, probe.formStage)
		}
	case *models.DATATYPE_DEFINITION_DATE:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_DATATYPESs {
				for _, target := range src.DATATYPE_DEFINITION_DATE {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_DATATYPES", "DATATYPE_DEFINITION_DATE", refNames, formGroup, probe.formStage)
		}
	case *models.DATATYPE_DEFINITION_ENUMERATION:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("SPECIFIED_VALUES", inst.SPECIFIED_VALUES, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SPECIFIED_VALUES](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_DATATYPESs {
				for _, target := range src.DATATYPE_DEFINITION_ENUMERATION {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_DATATYPES", "DATATYPE_DEFINITION_ENUMERATION", refNames, formGroup, probe.formStage)
		}
	case *models.DATATYPE_DEFINITION_INTEGER:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MAX_XSD", inst.MAX_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MIN_XSD", inst.MIN_XSD, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_DATATYPESs {
				for _, target := range src.DATATYPE_DEFINITION_INTEGER {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_DATATYPES", "DATATYPE_DEFINITION_INTEGER", refNames, formGroup, probe.formStage)
		}
	case *models.DATATYPE_DEFINITION_REAL:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ACCURACY", inst.ACCURACY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MAX_XSD", inst.MAX_XSD, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MIN_XSD", inst.MIN_XSD, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_DATATYPESs {
				for _, target := range src.DATATYPE_DEFINITION_REAL {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_DATATYPES", "DATATYPE_DEFINITION_REAL", refNames, formGroup, probe.formStage)
		}
	case *models.DATATYPE_DEFINITION_STRING:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MAX_LENGTH", inst.MAX_LENGTH, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_DATATYPESs {
				for _, target := range src.DATATYPE_DEFINITION_STRING {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_DATATYPES", "DATATYPE_DEFINITION_STRING", refNames, formGroup, probe.formStage)
		}
	case *models.DATATYPE_DEFINITION_XHTML:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_DATATYPESs {
				for _, target := range src.DATATYPE_DEFINITION_XHTML {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_DATATYPES", "DATATYPE_DEFINITION_XHTML", refNames, formGroup, probe.formStage)
		}
	case *models.EMBEDDED_VALUE:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("KEY", inst.KEY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OTHER_CONTENT", inst.OTHER_CONTENT, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_PROPERTIESs {
				if src.EMBEDDED_VALUE == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_PROPERTIES", "EMBEDDED_VALUE", refNames, formGroup, probe.formStage)
		}
	case *models.ENUM_VALUE:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("PROPERTIES", inst.PROPERTIES, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_PROPERTIES](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_SPECIFIED_VALUESs {
				for _, target := range src.ENUM_VALUE {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_SPECIFIED_VALUES", "ENUM_VALUE", refNames, formGroup, probe.formStage)
		}
	case *models.RELATION_GROUP:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("SOURCE_SPECIFICATION", inst.SOURCE_SPECIFICATION, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SOURCE_SPECIFICATION](), probe.formStage)
		StageSetAssociationFieldToForm("SPEC_RELATIONS", inst.SPEC_RELATIONS, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_RELATION_REF](), probe.formStage)
		StageSetAssociationFieldToForm("TARGET_SPECIFICATION", inst.TARGET_SPECIFICATION, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SOURCE_SPECIFICATION](), probe.formStage)
		StageSetAssociationFieldToForm("TYPE_XSD", inst.TYPE_XSD, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_RELATION_GROUP_TYPE_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_SPEC_RELATION_GROUPSs {
				for _, target := range src.RELATION_GROUP {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_SPEC_RELATION_GROUPS", "RELATION_GROUP", refNames, formGroup, probe.formStage)
		}
	case *models.RELATION_GROUP_TYPE:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("SPEC_ATTRIBUTES", inst.SPEC_ATTRIBUTES, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_ATTRIBUTES](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_SPEC_TYPESs {
				for _, target := range src.RELATION_GROUP_TYPE {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_SPEC_TYPES", "RELATION_GROUP_TYPE", refNames, formGroup, probe.formStage)
		}
	case *models.REQ_IF:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Lang", inst.Lang, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("THE_HEADER", inst.THE_HEADER, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_THE_HEADER](), probe.formStage)
		StageSetAssociationFieldToForm("CORE_CONTENT", inst.CORE_CONTENT, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_CORE_CONTENT](), probe.formStage)
		StageSetAssociationFieldToForm("TOOL_EXTENSIONS", inst.TOOL_EXTENSIONS, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_TOOL_EXTENSIONS](), probe.formStage)
	case *models.REQ_IF_CONTENT:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("DATATYPES", inst.DATATYPES, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_DATATYPES](), probe.formStage)
		StageSetAssociationFieldToForm("SPEC_TYPES", inst.SPEC_TYPES, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_TYPES](), probe.formStage)
		StageSetAssociationFieldToForm("SPEC_OBJECTS", inst.SPEC_OBJECTS, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_OBJECTS](), probe.formStage)
		StageSetAssociationFieldToForm("SPEC_RELATIONS", inst.SPEC_RELATIONS, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_RELATIONS](), probe.formStage)
		StageSetAssociationFieldToForm("SPECIFICATIONS", inst.SPECIFICATIONS, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SPECIFICATIONS](), probe.formStage)
		StageSetAssociationFieldToForm("SPEC_RELATION_GROUPS", inst.SPEC_RELATION_GROUPS, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_RELATION_GROUPS](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_CORE_CONTENTs {
				if src.REQ_IF_CONTENT == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_CORE_CONTENT", "REQ_IF_CONTENT", refNames, formGroup, probe.formStage)
		}
	case *models.REQ_IF_HEADER:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("COMMENT", inst.COMMENT, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CREATION_TIME", inst.CREATION_TIME, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("REPOSITORY_ID", inst.REPOSITORY_ID, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("REQ_IF_TOOL_ID", inst.REQ_IF_TOOL_ID, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("REQ_IF_VERSION", inst.REQ_IF_VERSION, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SOURCE_TOOL_ID", inst.SOURCE_TOOL_ID, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TITLE", inst.TITLE, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_THE_HEADERs {
				if src.REQ_IF_HEADER == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_THE_HEADER", "REQ_IF_HEADER", refNames, formGroup, probe.formStage)
		}
	case *models.REQ_IF_TOOL_EXTENSION:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_TOOL_EXTENSIONSs {
				for _, target := range src.REQ_IF_TOOL_EXTENSION {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_TOOL_EXTENSIONS", "REQ_IF_TOOL_EXTENSION", refNames, formGroup, probe.formStage)
		}
	case *models.SPECIFICATION:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("CHILDREN", inst.CHILDREN, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_CHILDREN](), probe.formStage)
		StageSetAssociationFieldToForm("VALUES", inst.VALUES, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_XHTML_1](), probe.formStage)
		StageSetAssociationFieldToForm("TYPE_XSD", inst.TYPE_XSD, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SPECIFICATION_TYPE_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_SPECIFICATIONSs {
				for _, target := range src.SPECIFICATION {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_SPECIFICATIONS", "SPECIFICATION", refNames, formGroup, probe.formStage)
		}
	case *models.SPECIFICATION_TYPE:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("SPEC_ATTRIBUTES", inst.SPEC_ATTRIBUTES, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_ATTRIBUTES](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_SPEC_TYPESs {
				for _, target := range src.SPECIFICATION_TYPE {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_SPEC_TYPES", "SPECIFICATION_TYPE", refNames, formGroup, probe.formStage)
		}
	case *models.SPEC_HIERARCHY:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IS_EDITABLE", inst.IS_EDITABLE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IS_TABLE_INTERNAL", inst.IS_TABLE_INTERNAL, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("CHILDREN", inst.CHILDREN, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_CHILDREN](), probe.formStage)
		StageSetAssociationFieldToForm("EDITABLE_ATTS", inst.EDITABLE_ATTS, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_EDITABLE_ATTS](), probe.formStage)
		StageSetAssociationFieldToForm("OBJECT", inst.OBJECT, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_OBJECT](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_CHILDRENs {
				for _, target := range src.SPEC_HIERARCHY {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_CHILDREN", "SPEC_HIERARCHY", refNames, formGroup, probe.formStage)
		}
	case *models.SPEC_OBJECT:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("VALUES", inst.VALUES, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_XHTML_1](), probe.formStage)
		StageSetAssociationFieldToForm("TYPE_XSD", inst.TYPE_XSD, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_OBJECT_TYPE_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_SPEC_OBJECTSs {
				for _, target := range src.SPEC_OBJECT {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_SPEC_OBJECTS", "SPEC_OBJECT", refNames, formGroup, probe.formStage)
		}
	case *models.SPEC_OBJECT_TYPE:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("SPEC_ATTRIBUTES", inst.SPEC_ATTRIBUTES, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_ATTRIBUTES](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_SPEC_TYPESs {
				for _, target := range src.SPEC_OBJECT_TYPE {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_SPEC_TYPES", "SPEC_OBJECT_TYPE", refNames, formGroup, probe.formStage)
		}
	case *models.SPEC_RELATION:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("VALUES", inst.VALUES, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ATTRIBUTE_VALUE_XHTML_1](), probe.formStage)
		StageSetAssociationFieldToForm("SOURCE", inst.SOURCE, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SOURCE](), probe.formStage)
		StageSetAssociationFieldToForm("TARGET", inst.TARGET, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SOURCE](), probe.formStage)
		StageSetAssociationFieldToForm("TYPE_XSD", inst.TYPE_XSD, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_RELATION_TYPE_REF](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_SPEC_RELATIONSs {
				for _, target := range src.SPEC_RELATION {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_SPEC_RELATIONS", "SPEC_RELATION", refNames, formGroup, probe.formStage)
		}
	case *models.SPEC_RELATION_TYPE:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DESC", inst.DESC, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IDENTIFIER", inst.IDENTIFIER, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LAST_CHANGE", inst.LAST_CHANGE, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LONG_NAME", inst.LONG_NAME, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("ALTERNATIVE_ID", inst.ALTERNATIVE_ID, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_ALTERNATIVE_ID](), probe.formStage)
		StageSetAssociationFieldToForm("SPEC_ATTRIBUTES", inst.SPEC_ATTRIBUTES, formGroup, probe.stageOfInterest.GetInstancesSet[*models.A_SPEC_ATTRIBUTES](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.A_SPEC_TYPESs {
				for _, target := range src.SPEC_RELATION_TYPE {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A_SPEC_TYPES", "SPEC_RELATION_TYPE", refNames, formGroup, probe.formStage)
		}
	case *models.XHTML_CONTENT:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("EnclosedText", inst.EnclosedText, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_VALUE_XHTMLs {
				if src.THE_VALUE == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_VALUE_XHTML", "THE_VALUE", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.ATTRIBUTE_VALUE_XHTMLs {
				if src.THE_ORIGINAL_VALUE == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.ATTRIBUTE_VALUE_XHTML", "THE_ORIGINAL_VALUE", refNames, formGroup, probe.formStage)
		}
	case *xmlxsd.Xmlxsd_SchemaElement:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
	default:
		_ = inst
	}
}
