// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/form/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.CheckBox:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.FormDiv, *models.CheckBox](
				"FormDiv",
				"CheckBoxs",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.FormDiv) []*models.CheckBox {
					return owner.CheckBoxs
				})
		}

	case *models.FormDiv:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("FormFields", instanceWithInferedType, &instanceWithInferedType.FormFields, formGroup, probe)
		AssociationSliceToForm("CheckBoxs", instanceWithInferedType, &instanceWithInferedType.CheckBoxs, formGroup, probe)
		AssociationFieldToForm("FormEditAssocButton", instanceWithInferedType.FormEditAssocButton, formGroup, probe)
		AssociationFieldToForm("FormSortAssocButton", instanceWithInferedType.FormSortAssocButton, formGroup, probe)
		BasicFieldtoForm("IsADivider", instanceWithInferedType.IsADivider, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsAStartAccordionGroup", instanceWithInferedType.IsAStartAccordionGroup, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("AccordionGroupName", instanceWithInferedType.AccordionGroupName, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsAEndAccordionGroup", instanceWithInferedType.IsAEndAccordionGroup, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.FormGroup, *models.FormDiv](
				"FormGroup",
				"FormDivs",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.FormGroup) []*models.FormDiv {
					return owner.FormDivs
				})
		}

	case *models.FormEditAssocButton:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("AssociationStorage", instanceWithInferedType.AssociationStorage, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasChanged", instanceWithInferedType.HasChanged, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsForSavePurpose", instanceWithInferedType.IsForSavePurpose, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasToolTip", instanceWithInferedType.HasToolTip, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ToolTipText", instanceWithInferedType.ToolTipText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MatTooltipShowDelay", instanceWithInferedType.MatTooltipShowDelay, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.FormField:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("InputTypeEnum", instanceWithInferedType.InputTypeEnum, instanceWithInferedType, probe.formStage, formGroup)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Placeholder", instanceWithInferedType.Placeholder, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("FormFieldString", instanceWithInferedType.FormFieldString, formGroup, probe)
		AssociationFieldToForm("FormFieldFloat64", instanceWithInferedType.FormFieldFloat64, formGroup, probe)
		AssociationFieldToForm("FormFieldInt", instanceWithInferedType.FormFieldInt, formGroup, probe)
		AssociationFieldToForm("FormFieldDate", instanceWithInferedType.FormFieldDate, formGroup, probe)
		AssociationFieldToForm("FormFieldTime", instanceWithInferedType.FormFieldTime, formGroup, probe)
		AssociationFieldToForm("FormFieldDateTime", instanceWithInferedType.FormFieldDateTime, formGroup, probe)
		AssociationFieldToForm("FormFieldSelect", instanceWithInferedType.FormFieldSelect, formGroup, probe)
		BasicFieldtoForm("HasBespokeWidth", instanceWithInferedType.HasBespokeWidth, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BespokeWidthPx", instanceWithInferedType.BespokeWidthPx, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasBespokeHeight", instanceWithInferedType.HasBespokeHeight, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BespokeHeightPx", instanceWithInferedType.BespokeHeightPx, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.FormDiv, *models.FormField](
				"FormDiv",
				"FormFields",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.FormDiv) []*models.FormField {
					return owner.FormFields
				})
		}

	case *models.FormFieldDate:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.FormFieldDateTime:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.FormFieldFloat64:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasMinValidator", instanceWithInferedType.HasMinValidator, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MinValue", instanceWithInferedType.MinValue, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasMaxValidator", instanceWithInferedType.HasMaxValidator, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MaxValue", instanceWithInferedType.MaxValue, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.FormFieldInt:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasMinValidator", instanceWithInferedType.HasMinValidator, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MinValue", instanceWithInferedType.MinValue, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasMaxValidator", instanceWithInferedType.HasMaxValidator, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MaxValue", instanceWithInferedType.MaxValue, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.FormFieldSelect:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("Value", instanceWithInferedType.Value, formGroup, probe)
		AssociationSliceToForm("Options", instanceWithInferedType, &instanceWithInferedType.Options, formGroup, probe)
		BasicFieldtoForm("CanBeEmpty", instanceWithInferedType.CanBeEmpty, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("PreserveInitialOrder", instanceWithInferedType.PreserveInitialOrder, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.FormFieldString:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsTextArea", instanceWithInferedType.IsTextArea, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.FormFieldTime:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Step", instanceWithInferedType.Step, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.FormGroup:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("TypeLabel", instanceWithInferedType.TypeLabel, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("FormDivs", instanceWithInferedType, &instanceWithInferedType.FormDivs, formGroup, probe)
		BasicFieldtoForm("HasSuppressButton", instanceWithInferedType.HasSuppressButton, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasSuppressButtonBeenPressed", instanceWithInferedType.HasSuppressButtonBeenPressed, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.FormSortAssocButton:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Label", instanceWithInferedType.Label, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasToolTip", instanceWithInferedType.HasToolTip, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ToolTipText", instanceWithInferedType.ToolTipText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MatTooltipShowDelay", instanceWithInferedType.MatTooltipShowDelay, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("FormEditAssocButton", instanceWithInferedType.FormEditAssocButton, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Option:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.FormFieldSelect, *models.Option](
				"FormFieldSelect",
				"Options",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.FormFieldSelect) []*models.Option {
					return owner.Options
				})
		}

	default:
		_ = instanceWithInferedType
	}
}
