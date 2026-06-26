// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/dsm/xsd/tests/reqif/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.ALTERNATIVE_ID:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.ATTRIBUTE_DEFINITION_BOOLEAN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("DEFAULT_VALUE", instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_SPEC_ATTRIBUTES, *models.ATTRIBUTE_DEFINITION_BOOLEAN](
				"A_SPEC_ATTRIBUTES",
				"ATTRIBUTE_DEFINITION_BOOLEAN",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_SPEC_ATTRIBUTES) []*models.ATTRIBUTE_DEFINITION_BOOLEAN {
					return owner.ATTRIBUTE_DEFINITION_BOOLEAN
				})
		}

	case *models.ATTRIBUTE_DEFINITION_DATE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("DEFAULT_VALUE", instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_SPEC_ATTRIBUTES, *models.ATTRIBUTE_DEFINITION_DATE](
				"A_SPEC_ATTRIBUTES",
				"ATTRIBUTE_DEFINITION_DATE",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_SPEC_ATTRIBUTES) []*models.ATTRIBUTE_DEFINITION_DATE {
					return owner.ATTRIBUTE_DEFINITION_DATE
				})
		}

	case *models.ATTRIBUTE_DEFINITION_ENUMERATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MULTI_VALUED", instanceWithInferedType.MULTI_VALUED, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("DEFAULT_VALUE", instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_SPEC_ATTRIBUTES, *models.ATTRIBUTE_DEFINITION_ENUMERATION](
				"A_SPEC_ATTRIBUTES",
				"ATTRIBUTE_DEFINITION_ENUMERATION",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_SPEC_ATTRIBUTES) []*models.ATTRIBUTE_DEFINITION_ENUMERATION {
					return owner.ATTRIBUTE_DEFINITION_ENUMERATION
				})
		}

	case *models.ATTRIBUTE_DEFINITION_INTEGER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("DEFAULT_VALUE", instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_SPEC_ATTRIBUTES, *models.ATTRIBUTE_DEFINITION_INTEGER](
				"A_SPEC_ATTRIBUTES",
				"ATTRIBUTE_DEFINITION_INTEGER",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_SPEC_ATTRIBUTES) []*models.ATTRIBUTE_DEFINITION_INTEGER {
					return owner.ATTRIBUTE_DEFINITION_INTEGER
				})
		}

	case *models.ATTRIBUTE_DEFINITION_REAL:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("DEFAULT_VALUE", instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_SPEC_ATTRIBUTES, *models.ATTRIBUTE_DEFINITION_REAL](
				"A_SPEC_ATTRIBUTES",
				"ATTRIBUTE_DEFINITION_REAL",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_SPEC_ATTRIBUTES) []*models.ATTRIBUTE_DEFINITION_REAL {
					return owner.ATTRIBUTE_DEFINITION_REAL
				})
		}

	case *models.ATTRIBUTE_DEFINITION_STRING:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("DEFAULT_VALUE", instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_SPEC_ATTRIBUTES, *models.ATTRIBUTE_DEFINITION_STRING](
				"A_SPEC_ATTRIBUTES",
				"ATTRIBUTE_DEFINITION_STRING",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_SPEC_ATTRIBUTES) []*models.ATTRIBUTE_DEFINITION_STRING {
					return owner.ATTRIBUTE_DEFINITION_STRING
				})
		}

	case *models.ATTRIBUTE_DEFINITION_XHTML:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("DEFAULT_VALUE", instanceWithInferedType.DEFAULT_VALUE, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_SPEC_ATTRIBUTES, *models.ATTRIBUTE_DEFINITION_XHTML](
				"A_SPEC_ATTRIBUTES",
				"ATTRIBUTE_DEFINITION_XHTML",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_SPEC_ATTRIBUTES) []*models.ATTRIBUTE_DEFINITION_XHTML {
					return owner.ATTRIBUTE_DEFINITION_XHTML
				})
		}

	case *models.ATTRIBUTE_VALUE_BOOLEAN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_ATTRIBUTE_VALUE_BOOLEAN, *models.ATTRIBUTE_VALUE_BOOLEAN](
				"A_ATTRIBUTE_VALUE_BOOLEAN",
				"ATTRIBUTE_VALUE_BOOLEAN",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_ATTRIBUTE_VALUE_BOOLEAN) []*models.ATTRIBUTE_VALUE_BOOLEAN {
					return owner.ATTRIBUTE_VALUE_BOOLEAN
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_ATTRIBUTE_VALUE_XHTML_1, *models.ATTRIBUTE_VALUE_BOOLEAN](
				"A_ATTRIBUTE_VALUE_XHTML_1",
				"ATTRIBUTE_VALUE_BOOLEAN",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_ATTRIBUTE_VALUE_XHTML_1) []*models.ATTRIBUTE_VALUE_BOOLEAN {
					return owner.ATTRIBUTE_VALUE_BOOLEAN
				})
		}

	case *models.ATTRIBUTE_VALUE_DATE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_ATTRIBUTE_VALUE_DATE, *models.ATTRIBUTE_VALUE_DATE](
				"A_ATTRIBUTE_VALUE_DATE",
				"ATTRIBUTE_VALUE_DATE",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_ATTRIBUTE_VALUE_DATE) []*models.ATTRIBUTE_VALUE_DATE {
					return owner.ATTRIBUTE_VALUE_DATE
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_ATTRIBUTE_VALUE_XHTML_1, *models.ATTRIBUTE_VALUE_DATE](
				"A_ATTRIBUTE_VALUE_XHTML_1",
				"ATTRIBUTE_VALUE_DATE",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_ATTRIBUTE_VALUE_XHTML_1) []*models.ATTRIBUTE_VALUE_DATE {
					return owner.ATTRIBUTE_VALUE_DATE
				})
		}

	case *models.ATTRIBUTE_VALUE_ENUMERATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)
		AssociationFieldToForm("VALUES", instanceWithInferedType.VALUES, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_ATTRIBUTE_VALUE_ENUMERATION, *models.ATTRIBUTE_VALUE_ENUMERATION](
				"A_ATTRIBUTE_VALUE_ENUMERATION",
				"ATTRIBUTE_VALUE_ENUMERATION",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_ATTRIBUTE_VALUE_ENUMERATION) []*models.ATTRIBUTE_VALUE_ENUMERATION {
					return owner.ATTRIBUTE_VALUE_ENUMERATION
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_ATTRIBUTE_VALUE_XHTML_1, *models.ATTRIBUTE_VALUE_ENUMERATION](
				"A_ATTRIBUTE_VALUE_XHTML_1",
				"ATTRIBUTE_VALUE_ENUMERATION",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_ATTRIBUTE_VALUE_XHTML_1) []*models.ATTRIBUTE_VALUE_ENUMERATION {
					return owner.ATTRIBUTE_VALUE_ENUMERATION
				})
		}

	case *models.ATTRIBUTE_VALUE_INTEGER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_ATTRIBUTE_VALUE_INTEGER, *models.ATTRIBUTE_VALUE_INTEGER](
				"A_ATTRIBUTE_VALUE_INTEGER",
				"ATTRIBUTE_VALUE_INTEGER",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_ATTRIBUTE_VALUE_INTEGER) []*models.ATTRIBUTE_VALUE_INTEGER {
					return owner.ATTRIBUTE_VALUE_INTEGER
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_ATTRIBUTE_VALUE_XHTML_1, *models.ATTRIBUTE_VALUE_INTEGER](
				"A_ATTRIBUTE_VALUE_XHTML_1",
				"ATTRIBUTE_VALUE_INTEGER",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_ATTRIBUTE_VALUE_XHTML_1) []*models.ATTRIBUTE_VALUE_INTEGER {
					return owner.ATTRIBUTE_VALUE_INTEGER
				})
		}

	case *models.ATTRIBUTE_VALUE_REAL:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_ATTRIBUTE_VALUE_REAL, *models.ATTRIBUTE_VALUE_REAL](
				"A_ATTRIBUTE_VALUE_REAL",
				"ATTRIBUTE_VALUE_REAL",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_ATTRIBUTE_VALUE_REAL) []*models.ATTRIBUTE_VALUE_REAL {
					return owner.ATTRIBUTE_VALUE_REAL
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_ATTRIBUTE_VALUE_XHTML_1, *models.ATTRIBUTE_VALUE_REAL](
				"A_ATTRIBUTE_VALUE_XHTML_1",
				"ATTRIBUTE_VALUE_REAL",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_ATTRIBUTE_VALUE_XHTML_1) []*models.ATTRIBUTE_VALUE_REAL {
					return owner.ATTRIBUTE_VALUE_REAL
				})
		}

	case *models.ATTRIBUTE_VALUE_STRING:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("THE_VALUE", instanceWithInferedType.THE_VALUE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_ATTRIBUTE_VALUE_STRING, *models.ATTRIBUTE_VALUE_STRING](
				"A_ATTRIBUTE_VALUE_STRING",
				"ATTRIBUTE_VALUE_STRING",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_ATTRIBUTE_VALUE_STRING) []*models.ATTRIBUTE_VALUE_STRING {
					return owner.ATTRIBUTE_VALUE_STRING
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_ATTRIBUTE_VALUE_XHTML_1, *models.ATTRIBUTE_VALUE_STRING](
				"A_ATTRIBUTE_VALUE_XHTML_1",
				"ATTRIBUTE_VALUE_STRING",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_ATTRIBUTE_VALUE_XHTML_1) []*models.ATTRIBUTE_VALUE_STRING {
					return owner.ATTRIBUTE_VALUE_STRING
				})
		}

	case *models.ATTRIBUTE_VALUE_XHTML:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IS_SIMPLIFIED", instanceWithInferedType.IS_SIMPLIFIED, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("THE_VALUE", instanceWithInferedType.THE_VALUE, formGroup, probe)
		AssociationFieldToForm("THE_ORIGINAL_VALUE", instanceWithInferedType.THE_ORIGINAL_VALUE, formGroup, probe)
		AssociationFieldToForm("DEFINITION", instanceWithInferedType.DEFINITION, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_ATTRIBUTE_VALUE_XHTML, *models.ATTRIBUTE_VALUE_XHTML](
				"A_ATTRIBUTE_VALUE_XHTML",
				"ATTRIBUTE_VALUE_XHTML",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_ATTRIBUTE_VALUE_XHTML) []*models.ATTRIBUTE_VALUE_XHTML {
					return owner.ATTRIBUTE_VALUE_XHTML
				})
		}
		{
			AssociationReverseSliceToForm[*models.A_ATTRIBUTE_VALUE_XHTML_1, *models.ATTRIBUTE_VALUE_XHTML](
				"A_ATTRIBUTE_VALUE_XHTML_1",
				"ATTRIBUTE_VALUE_XHTML",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_ATTRIBUTE_VALUE_XHTML_1) []*models.ATTRIBUTE_VALUE_XHTML {
					return owner.ATTRIBUTE_VALUE_XHTML
				})
		}

	case *models.A_ALTERNATIVE_ID:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_ATTRIBUTE_DEFINITION_BOOLEAN_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_BOOLEAN_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_BOOLEAN_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_ATTRIBUTE_DEFINITION_DATE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_DATE_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_DATE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_ATTRIBUTE_DEFINITION_ENUMERATION_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_ENUMERATION_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_ENUMERATION_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_ATTRIBUTE_DEFINITION_INTEGER_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_INTEGER_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_INTEGER_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_ATTRIBUTE_DEFINITION_REAL_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_REAL_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_REAL_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_ATTRIBUTE_DEFINITION_STRING_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_STRING_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_STRING_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_ATTRIBUTE_DEFINITION_XHTML_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_XHTML_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_XHTML_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_ATTRIBUTE_VALUE_BOOLEAN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ATTRIBUTE_VALUE_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_BOOLEAN, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_ATTRIBUTE_VALUE_DATE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ATTRIBUTE_VALUE_DATE", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_DATE, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_ATTRIBUTE_VALUE_ENUMERATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ATTRIBUTE_VALUE_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_ENUMERATION, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_ATTRIBUTE_VALUE_INTEGER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ATTRIBUTE_VALUE_INTEGER", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_INTEGER, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_ATTRIBUTE_VALUE_REAL:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ATTRIBUTE_VALUE_REAL", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_REAL, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_ATTRIBUTE_VALUE_STRING:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ATTRIBUTE_VALUE_STRING", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_STRING, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_ATTRIBUTE_VALUE_XHTML:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ATTRIBUTE_VALUE_XHTML", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_XHTML, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_ATTRIBUTE_VALUE_XHTML_1:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ATTRIBUTE_VALUE_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_DATE", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_DATE, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_INTEGER", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_INTEGER, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_REAL", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_REAL, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_STRING", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_STRING, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_VALUE_XHTML", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_VALUE_XHTML, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_CHILDREN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SPEC_HIERARCHY", instanceWithInferedType, &instanceWithInferedType.SPEC_HIERARCHY, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_CORE_CONTENT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("REQ_IF_CONTENT", instanceWithInferedType.REQ_IF_CONTENT, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_DATATYPES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("DATATYPE_DEFINITION_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.DATATYPE_DEFINITION_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_DATE", instanceWithInferedType, &instanceWithInferedType.DATATYPE_DEFINITION_DATE, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.DATATYPE_DEFINITION_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_INTEGER", instanceWithInferedType, &instanceWithInferedType.DATATYPE_DEFINITION_INTEGER, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_REAL", instanceWithInferedType, &instanceWithInferedType.DATATYPE_DEFINITION_REAL, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_STRING", instanceWithInferedType, &instanceWithInferedType.DATATYPE_DEFINITION_STRING, formGroup, probe)
		AssociationSliceToForm("DATATYPE_DEFINITION_XHTML", instanceWithInferedType, &instanceWithInferedType.DATATYPE_DEFINITION_XHTML, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_DATATYPE_DEFINITION_BOOLEAN_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DATATYPE_DEFINITION_BOOLEAN_REF", instanceWithInferedType.DATATYPE_DEFINITION_BOOLEAN_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_DATATYPE_DEFINITION_DATE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DATATYPE_DEFINITION_DATE_REF", instanceWithInferedType.DATATYPE_DEFINITION_DATE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_DATATYPE_DEFINITION_ENUMERATION_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DATATYPE_DEFINITION_ENUMERATION_REF", instanceWithInferedType.DATATYPE_DEFINITION_ENUMERATION_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_DATATYPE_DEFINITION_INTEGER_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DATATYPE_DEFINITION_INTEGER_REF", instanceWithInferedType.DATATYPE_DEFINITION_INTEGER_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_DATATYPE_DEFINITION_REAL_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DATATYPE_DEFINITION_REAL_REF", instanceWithInferedType.DATATYPE_DEFINITION_REAL_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_DATATYPE_DEFINITION_STRING_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DATATYPE_DEFINITION_STRING_REF", instanceWithInferedType.DATATYPE_DEFINITION_STRING_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_DATATYPE_DEFINITION_XHTML_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DATATYPE_DEFINITION_XHTML_REF", instanceWithInferedType.DATATYPE_DEFINITION_XHTML_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_EDITABLE_ATTS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_BOOLEAN_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_BOOLEAN_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_DATE_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_DATE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_ENUMERATION_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_ENUMERATION_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_INTEGER_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_INTEGER_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_REAL_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_REAL_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_STRING_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_STRING_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ATTRIBUTE_DEFINITION_XHTML_REF", instanceWithInferedType.ATTRIBUTE_DEFINITION_XHTML_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_ENUM_VALUE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ENUM_VALUE_REF", instanceWithInferedType.ENUM_VALUE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_OBJECT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SPEC_OBJECT_REF", instanceWithInferedType.SPEC_OBJECT_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_PROPERTIES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("EMBEDDED_VALUE", instanceWithInferedType.EMBEDDED_VALUE, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_RELATION_GROUP_TYPE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("RELATION_GROUP_TYPE_REF", instanceWithInferedType.RELATION_GROUP_TYPE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_SOURCE_1:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SPEC_OBJECT_REF", instanceWithInferedType.SPEC_OBJECT_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_SOURCE_SPECIFICATION_1:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("SPECIFICATION_REF", instanceWithInferedType.SPECIFICATION_REF, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_SPECIFICATIONS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SPECIFICATION", instanceWithInferedType, &instanceWithInferedType.SPECIFICATION, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_SPECIFICATION_TYPE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SPECIFICATION_TYPE_REF", instanceWithInferedType.SPECIFICATION_TYPE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_SPECIFIED_VALUES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ENUM_VALUE", instanceWithInferedType, &instanceWithInferedType.ENUM_VALUE, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_SPEC_ATTRIBUTES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_BOOLEAN", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_DEFINITION_BOOLEAN, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_DATE", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_DEFINITION_DATE, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_ENUMERATION", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_DEFINITION_ENUMERATION, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_INTEGER", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_DEFINITION_INTEGER, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_REAL", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_DEFINITION_REAL, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_STRING", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_DEFINITION_STRING, formGroup, probe)
		AssociationSliceToForm("ATTRIBUTE_DEFINITION_XHTML", instanceWithInferedType, &instanceWithInferedType.ATTRIBUTE_DEFINITION_XHTML, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_SPEC_OBJECTS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SPEC_OBJECT", instanceWithInferedType, &instanceWithInferedType.SPEC_OBJECT, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_SPEC_OBJECT_TYPE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SPEC_OBJECT_TYPE_REF", instanceWithInferedType.SPEC_OBJECT_TYPE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_SPEC_RELATIONS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("SPEC_RELATION", instanceWithInferedType, &instanceWithInferedType.SPEC_RELATION, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_SPEC_RELATION_GROUPS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("RELATION_GROUP", instanceWithInferedType, &instanceWithInferedType.RELATION_GROUP, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_SPEC_RELATION_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SPEC_RELATION_REF", instanceWithInferedType.SPEC_RELATION_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_SPEC_RELATION_TYPE_REF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SPEC_RELATION_TYPE_REF", instanceWithInferedType.SPEC_RELATION_TYPE_REF, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_SPEC_TYPES:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("RELATION_GROUP_TYPE", instanceWithInferedType, &instanceWithInferedType.RELATION_GROUP_TYPE, formGroup, probe)
		AssociationSliceToForm("SPEC_OBJECT_TYPE", instanceWithInferedType, &instanceWithInferedType.SPEC_OBJECT_TYPE, formGroup, probe)
		AssociationSliceToForm("SPEC_RELATION_TYPE", instanceWithInferedType, &instanceWithInferedType.SPEC_RELATION_TYPE, formGroup, probe)
		AssociationSliceToForm("SPECIFICATION_TYPE", instanceWithInferedType, &instanceWithInferedType.SPECIFICATION_TYPE, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_THE_HEADER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("REQ_IF_HEADER", instanceWithInferedType.REQ_IF_HEADER, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.A_TOOL_EXTENSIONS:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationSliceToForm("REQ_IF_TOOL_EXTENSION", instanceWithInferedType, &instanceWithInferedType.REQ_IF_TOOL_EXTENSION, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.DATATYPE_DEFINITION_BOOLEAN:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_DATATYPES, *models.DATATYPE_DEFINITION_BOOLEAN](
				"A_DATATYPES",
				"DATATYPE_DEFINITION_BOOLEAN",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_DATATYPES) []*models.DATATYPE_DEFINITION_BOOLEAN {
					return owner.DATATYPE_DEFINITION_BOOLEAN
				})
		}

	case *models.DATATYPE_DEFINITION_DATE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_DATATYPES, *models.DATATYPE_DEFINITION_DATE](
				"A_DATATYPES",
				"DATATYPE_DEFINITION_DATE",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_DATATYPES) []*models.DATATYPE_DEFINITION_DATE {
					return owner.DATATYPE_DEFINITION_DATE
				})
		}

	case *models.DATATYPE_DEFINITION_ENUMERATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("SPECIFIED_VALUES", instanceWithInferedType.SPECIFIED_VALUES, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_DATATYPES, *models.DATATYPE_DEFINITION_ENUMERATION](
				"A_DATATYPES",
				"DATATYPE_DEFINITION_ENUMERATION",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_DATATYPES) []*models.DATATYPE_DEFINITION_ENUMERATION {
					return owner.DATATYPE_DEFINITION_ENUMERATION
				})
		}

	case *models.DATATYPE_DEFINITION_INTEGER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MAX", instanceWithInferedType.MAX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MIN", instanceWithInferedType.MIN, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_DATATYPES, *models.DATATYPE_DEFINITION_INTEGER](
				"A_DATATYPES",
				"DATATYPE_DEFINITION_INTEGER",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_DATATYPES) []*models.DATATYPE_DEFINITION_INTEGER {
					return owner.DATATYPE_DEFINITION_INTEGER
				})
		}

	case *models.DATATYPE_DEFINITION_REAL:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("ACCURACY", instanceWithInferedType.ACCURACY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MAX", instanceWithInferedType.MAX, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MIN", instanceWithInferedType.MIN, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_DATATYPES, *models.DATATYPE_DEFINITION_REAL](
				"A_DATATYPES",
				"DATATYPE_DEFINITION_REAL",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_DATATYPES) []*models.DATATYPE_DEFINITION_REAL {
					return owner.DATATYPE_DEFINITION_REAL
				})
		}

	case *models.DATATYPE_DEFINITION_STRING:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("MAX_LENGTH", instanceWithInferedType.MAX_LENGTH, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_DATATYPES, *models.DATATYPE_DEFINITION_STRING](
				"A_DATATYPES",
				"DATATYPE_DEFINITION_STRING",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_DATATYPES) []*models.DATATYPE_DEFINITION_STRING {
					return owner.DATATYPE_DEFINITION_STRING
				})
		}

	case *models.DATATYPE_DEFINITION_XHTML:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_DATATYPES, *models.DATATYPE_DEFINITION_XHTML](
				"A_DATATYPES",
				"DATATYPE_DEFINITION_XHTML",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_DATATYPES) []*models.DATATYPE_DEFINITION_XHTML {
					return owner.DATATYPE_DEFINITION_XHTML
				})
		}

	case *models.EMBEDDED_VALUE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("KEY", instanceWithInferedType.KEY, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("OTHER_CONTENT", instanceWithInferedType.OTHER_CONTENT, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.ENUM_VALUE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("PROPERTIES", instanceWithInferedType.PROPERTIES, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_SPECIFIED_VALUES, *models.ENUM_VALUE](
				"A_SPECIFIED_VALUES",
				"ENUM_VALUE",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_SPECIFIED_VALUES) []*models.ENUM_VALUE {
					return owner.ENUM_VALUE
				})
		}

	case *models.RELATION_GROUP:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("SOURCE_SPECIFICATION", instanceWithInferedType.SOURCE_SPECIFICATION, formGroup, probe)
		AssociationFieldToForm("SPEC_RELATIONS", instanceWithInferedType.SPEC_RELATIONS, formGroup, probe)
		AssociationFieldToForm("TARGET_SPECIFICATION", instanceWithInferedType.TARGET_SPECIFICATION, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_SPEC_RELATION_GROUPS, *models.RELATION_GROUP](
				"A_SPEC_RELATION_GROUPS",
				"RELATION_GROUP",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_SPEC_RELATION_GROUPS) []*models.RELATION_GROUP {
					return owner.RELATION_GROUP
				})
		}

	case *models.RELATION_GROUP_TYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("SPEC_ATTRIBUTES", instanceWithInferedType.SPEC_ATTRIBUTES, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_SPEC_TYPES, *models.RELATION_GROUP_TYPE](
				"A_SPEC_TYPES",
				"RELATION_GROUP_TYPE",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_SPEC_TYPES) []*models.RELATION_GROUP_TYPE {
					return owner.RELATION_GROUP_TYPE
				})
		}

	case *models.REQ_IF:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("Lang", instanceWithInferedType.Lang, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("THE_HEADER", instanceWithInferedType.THE_HEADER, formGroup, probe)
		AssociationFieldToForm("CORE_CONTENT", instanceWithInferedType.CORE_CONTENT, formGroup, probe)
		AssociationFieldToForm("TOOL_EXTENSIONS", instanceWithInferedType.TOOL_EXTENSIONS, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.REQ_IF_CONTENT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("DATATYPES", instanceWithInferedType.DATATYPES, formGroup, probe)
		AssociationFieldToForm("SPEC_TYPES", instanceWithInferedType.SPEC_TYPES, formGroup, probe)
		AssociationFieldToForm("SPEC_OBJECTS", instanceWithInferedType.SPEC_OBJECTS, formGroup, probe)
		AssociationFieldToForm("SPEC_RELATIONS", instanceWithInferedType.SPEC_RELATIONS, formGroup, probe)
		AssociationFieldToForm("SPECIFICATIONS", instanceWithInferedType.SPECIFICATIONS, formGroup, probe)
		AssociationFieldToForm("SPEC_RELATION_GROUPS", instanceWithInferedType.SPEC_RELATION_GROUPS, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.REQ_IF_HEADER:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("COMMENT", instanceWithInferedType.COMMENT, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("CREATION_TIME", instanceWithInferedType.CREATION_TIME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("REPOSITORY_ID", instanceWithInferedType.REPOSITORY_ID, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("REQ_IF_TOOL_ID", instanceWithInferedType.REQ_IF_TOOL_ID, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("REQ_IF_VERSION", instanceWithInferedType.REQ_IF_VERSION, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("SOURCE_TOOL_ID", instanceWithInferedType.SOURCE_TOOL_ID, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("TITLE", instanceWithInferedType.TITLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.REQ_IF_TOOL_EXTENSION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_TOOL_EXTENSIONS, *models.REQ_IF_TOOL_EXTENSION](
				"A_TOOL_EXTENSIONS",
				"REQ_IF_TOOL_EXTENSION",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_TOOL_EXTENSIONS) []*models.REQ_IF_TOOL_EXTENSION {
					return owner.REQ_IF_TOOL_EXTENSION
				})
		}

	case *models.SPECIFICATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("CHILDREN", instanceWithInferedType.CHILDREN, formGroup, probe)
		AssociationFieldToForm("VALUES", instanceWithInferedType.VALUES, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_SPECIFICATIONS, *models.SPECIFICATION](
				"A_SPECIFICATIONS",
				"SPECIFICATION",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_SPECIFICATIONS) []*models.SPECIFICATION {
					return owner.SPECIFICATION
				})
		}

	case *models.SPECIFICATION_TYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("SPEC_ATTRIBUTES", instanceWithInferedType.SPEC_ATTRIBUTES, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_SPEC_TYPES, *models.SPECIFICATION_TYPE](
				"A_SPEC_TYPES",
				"SPECIFICATION_TYPE",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_SPEC_TYPES) []*models.SPECIFICATION_TYPE {
					return owner.SPECIFICATION_TYPE
				})
		}

	case *models.SPEC_HIERARCHY:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IS_EDITABLE", instanceWithInferedType.IS_EDITABLE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IS_TABLE_INTERNAL", instanceWithInferedType.IS_TABLE_INTERNAL, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("CHILDREN", instanceWithInferedType.CHILDREN, formGroup, probe)
		AssociationFieldToForm("EDITABLE_ATTS", instanceWithInferedType.EDITABLE_ATTS, formGroup, probe)
		AssociationFieldToForm("OBJECT", instanceWithInferedType.OBJECT, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_CHILDREN, *models.SPEC_HIERARCHY](
				"A_CHILDREN",
				"SPEC_HIERARCHY",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_CHILDREN) []*models.SPEC_HIERARCHY {
					return owner.SPEC_HIERARCHY
				})
		}

	case *models.SPEC_OBJECT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("VALUES", instanceWithInferedType.VALUES, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_SPEC_OBJECTS, *models.SPEC_OBJECT](
				"A_SPEC_OBJECTS",
				"SPEC_OBJECT",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_SPEC_OBJECTS) []*models.SPEC_OBJECT {
					return owner.SPEC_OBJECT
				})
		}

	case *models.SPEC_OBJECT_TYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("SPEC_ATTRIBUTES", instanceWithInferedType.SPEC_ATTRIBUTES, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_SPEC_TYPES, *models.SPEC_OBJECT_TYPE](
				"A_SPEC_TYPES",
				"SPEC_OBJECT_TYPE",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_SPEC_TYPES) []*models.SPEC_OBJECT_TYPE {
					return owner.SPEC_OBJECT_TYPE
				})
		}

	case *models.SPEC_RELATION:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("VALUES", instanceWithInferedType.VALUES, formGroup, probe)
		AssociationFieldToForm("SOURCE", instanceWithInferedType.SOURCE, formGroup, probe)
		AssociationFieldToForm("TARGET", instanceWithInferedType.TARGET, formGroup, probe)
		AssociationFieldToForm("TYPE", instanceWithInferedType.TYPE, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_SPEC_RELATIONS, *models.SPEC_RELATION](
				"A_SPEC_RELATIONS",
				"SPEC_RELATION",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_SPEC_RELATIONS) []*models.SPEC_RELATION {
					return owner.SPEC_RELATION
				})
		}

	case *models.SPEC_RELATION_TYPE:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("DESC", instanceWithInferedType.DESC, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IDENTIFIER", instanceWithInferedType.IDENTIFIER, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LAST_CHANGE", instanceWithInferedType.LAST_CHANGE, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("LONG_NAME", instanceWithInferedType.LONG_NAME, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		AssociationFieldToForm("ALTERNATIVE_ID", instanceWithInferedType.ALTERNATIVE_ID, formGroup, probe)
		AssociationFieldToForm("SPEC_ATTRIBUTES", instanceWithInferedType.SPEC_ATTRIBUTES, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A_SPEC_TYPES, *models.SPEC_RELATION_TYPE](
				"A_SPEC_TYPES",
				"SPEC_RELATION_TYPE",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A_SPEC_TYPES) []*models.SPEC_RELATION_TYPE {
					return owner.SPEC_RELATION_TYPE
				})
		}

	case *models.XHTML_CONTENT:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("EnclosedText", instanceWithInferedType.EnclosedText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	default:
		_ = instanceWithInferedType
	}
}
