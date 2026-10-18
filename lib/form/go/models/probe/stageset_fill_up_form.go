// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/form/go/models"
)

var (
	_ = sort.Strings
	_ = strings.Join
)

func StageSetFillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *StageSetProbe,
) {
	switch inst := instance.(type) {
	case *models.CheckBox:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.FormDivs {
				for _, target := range src.CheckBoxs {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.FormDiv", "CheckBoxs", refNames, formGroup, probe.formStage)
		}
	case *models.FormDiv:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: FormFields
			div := (&form.FormDiv{Name: "FormFields"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.FormFields {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "FormFields",
				Label: "FormFields",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: CheckBoxs
			div := (&form.FormDiv{Name: "CheckBoxs"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.CheckBoxs {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "CheckBoxs",
				Label: "CheckBoxs",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("FormEditAssocButton", inst.FormEditAssocButton, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.FormEditAssocButton](), probe.formStage)
		StageSetAssociationFieldToForm("FormSortAssocButton", inst.FormSortAssocButton, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.FormSortAssocButton](), probe.formStage)
		StageSetBasicFieldtoForm("IsADivider", inst.IsADivider, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsAStartAccordionGroup", inst.IsAStartAccordionGroup, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AccordionGroupName", inst.AccordionGroupName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsAEndAccordionGroup", inst.IsAEndAccordionGroup, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.FormGroups {
				for _, target := range src.FormDivs {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.FormGroup", "FormDivs", refNames, formGroup, probe.formStage)
		}
	case *models.FormEditAssocButton:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Label", inst.Label, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("AssociationStorage", inst.AssociationStorage, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasChanged", inst.HasChanged, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsForSavePurpose", inst.IsForSavePurpose, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasToolTip", inst.HasToolTip, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ToolTipText", inst.ToolTipText, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MatTooltipShowDelay", inst.MatTooltipShowDelay, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.FormDivs {
				if src.FormEditAssocButton == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.FormDiv", "FormEditAssocButton", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.FormSortAssocButtons {
				if src.FormEditAssocButton == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.FormSortAssocButton", "FormEditAssocButton", refNames, formGroup, probe.formStage)
		}
	case *models.FormField:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("InputTypeEnum", inst.InputTypeEnum, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("Label", inst.Label, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Placeholder", inst.Placeholder, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("FormFieldString", inst.FormFieldString, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.FormFieldString](), probe.formStage)
		StageSetAssociationFieldToForm("FormFieldFloat64", inst.FormFieldFloat64, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.FormFieldFloat64](), probe.formStage)
		StageSetAssociationFieldToForm("FormFieldInt", inst.FormFieldInt, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.FormFieldInt](), probe.formStage)
		StageSetAssociationFieldToForm("FormFieldDate", inst.FormFieldDate, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.FormFieldDate](), probe.formStage)
		StageSetAssociationFieldToForm("FormFieldTime", inst.FormFieldTime, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.FormFieldTime](), probe.formStage)
		StageSetAssociationFieldToForm("FormFieldDateTime", inst.FormFieldDateTime, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.FormFieldDateTime](), probe.formStage)
		StageSetAssociationFieldToForm("FormFieldSelect", inst.FormFieldSelect, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.FormFieldSelect](), probe.formStage)
		StageSetBasicFieldtoForm("HasBespokeWidth", inst.HasBespokeWidth, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BespokeWidthPx", inst.BespokeWidthPx, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasBespokeHeight", inst.HasBespokeHeight, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BespokeHeightPx", inst.BespokeHeightPx, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.FormDivs {
				for _, target := range src.FormFields {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.FormDiv", "FormFields", refNames, formGroup, probe.formStage)
		}
	case *models.FormFieldDate:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.FormFields {
				if src.FormFieldDate == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.FormField", "FormFieldDate", refNames, formGroup, probe.formStage)
		}
	case *models.FormFieldDateTime:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.FormFields {
				if src.FormFieldDateTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.FormField", "FormFieldDateTime", refNames, formGroup, probe.formStage)
		}
	case *models.FormFieldFloat64:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasMinValidator", inst.HasMinValidator, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MinValue", inst.MinValue, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasMaxValidator", inst.HasMaxValidator, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MaxValue", inst.MaxValue, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.FormFields {
				if src.FormFieldFloat64 == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.FormField", "FormFieldFloat64", refNames, formGroup, probe.formStage)
		}
	case *models.FormFieldInt:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasMinValidator", inst.HasMinValidator, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MinValue", inst.MinValue, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasMaxValidator", inst.HasMaxValidator, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MaxValue", inst.MaxValue, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.FormFields {
				if src.FormFieldInt == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.FormField", "FormFieldInt", refNames, formGroup, probe.formStage)
		}
	case *models.FormFieldSelect:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Value", inst.Value, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Option](), probe.formStage)

		{
			// Slice of pointers: Options
			div := (&form.FormDiv{Name: "Options"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Options {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Options",
				Label: "Options",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("CanBeEmpty", inst.CanBeEmpty, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("PreserveInitialOrder", inst.PreserveInitialOrder, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.FormFields {
				if src.FormFieldSelect == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.FormField", "FormFieldSelect", refNames, formGroup, probe.formStage)
		}
	case *models.FormFieldString:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsTextArea", inst.IsTextArea, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.FormFields {
				if src.FormFieldString == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.FormField", "FormFieldString", refNames, formGroup, probe.formStage)
		}
	case *models.FormFieldTime:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Step", inst.Step, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.FormFields {
				if src.FormFieldTime == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.FormField", "FormFieldTime", refNames, formGroup, probe.formStage)
		}
	case *models.FormGroup:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Label", inst.Label, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TypeLabel", inst.TypeLabel, probe.formStage, formGroup)

		{
			// Slice of pointers: FormDivs
			div := (&form.FormDiv{Name: "FormDivs"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.FormDivs {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "FormDivs",
				Label: "FormDivs",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("HasSuppressButton", inst.HasSuppressButton, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasSuppressButtonBeenPressed", inst.HasSuppressButtonBeenPressed, probe.formStage, formGroup)
	case *models.FormSortAssocButton:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Label", inst.Label, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasToolTip", inst.HasToolTip, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ToolTipText", inst.ToolTipText, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MatTooltipShowDelay", inst.MatTooltipShowDelay, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("FormEditAssocButton", inst.FormEditAssocButton, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.FormEditAssocButton](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.FormDivs {
				if src.FormSortAssocButton == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.FormDiv", "FormSortAssocButton", refNames, formGroup, probe.formStage)
		}
	case *models.Option:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.FormFieldSelects {
				if src.Value == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.FormFieldSelect", "Value", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.FormFieldSelects {
				for _, target := range src.Options {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.FormFieldSelect", "Options", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
