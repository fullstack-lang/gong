// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/slider/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Checkbox:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ValueBool", instanceWithInferedType.ValueBool, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LabelForTrue", instanceWithInferedType.LabelForTrue, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("LabelForFalse", instanceWithInferedType.LabelForFalse, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Group, *models.Checkbox](
				"Group",
				"Checkboxes",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Group) []*models.Checkbox {
					return owner.Checkboxes
				})
		}

	case *models.Group:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Percentage", instanceWithInferedType.Percentage, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Sliders", instanceWithInferedType, &instanceWithInferedType.Sliders, formGroup, probe)
		AssociationSliceToForm("Checkboxes", instanceWithInferedType, &instanceWithInferedType.Checkboxes, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Layout, *models.Group](
				"Layout",
				"Groups",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Layout) []*models.Group {
					return owner.Groups
				})
		}

	case *models.Layout:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Groups", instanceWithInferedType, &instanceWithInferedType.Groups, formGroup, probe)
		BasicFieldtoForm("IsWithCustomGutterSize", instanceWithInferedType.IsWithCustomGutterSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("GutterSize", instanceWithInferedType.GutterSize, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Slider:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsFloat64", instanceWithInferedType.IsFloat64, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("IsInt", instanceWithInferedType.IsInt, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MinInt", instanceWithInferedType.MinInt, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MaxInt", instanceWithInferedType.MaxInt, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StepInt", instanceWithInferedType.StepInt, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ValueInt", instanceWithInferedType.ValueInt, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MinFloat64", instanceWithInferedType.MinFloat64, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("MaxFloat64", instanceWithInferedType.MaxFloat64, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("StepFloat64", instanceWithInferedType.StepFloat64, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ValueFloat64", instanceWithInferedType.ValueFloat64, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Group, *models.Slider](
				"Group",
				"Sliders",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Group) []*models.Slider {
					return owner.Sliders
				})
		}

	default:
		_ = instanceWithInferedType
	}
}
