// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test3/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.A:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			true, true, 600, true, 300, false)
		formGroup.FormDivs = append(formGroup.FormDivs, (&form.FormDiv{
			Name:       "Time fields",
			IsAStartAccordionGroup: true,
			AccordionGroupName: "Time fields",
		}).Stage(probe.formStage))
		BasicFieldtoForm("Date", instanceWithInferedType.Date, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, true)
		BasicFieldtoForm("Duration", instanceWithInferedType.Duration, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formGroup.FormDivs = append(formGroup.FormDivs, (&form.FormDiv{
			Name:       "",
			IsAEndAccordionGroup:   true,
		}).Stage(probe.formStage))
		BasicFieldtoForm("FloatValue", instanceWithInferedType.FloatValue, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		BasicFieldtoForm("IntValue", instanceWithInferedType.IntValue, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		EnumTypeStringToForm("EnumString", instanceWithInferedType.EnumString, instanceWithInferedType, probe.formStage, formGroup)
		EnumTypeIntToForm("EnumInt", instanceWithInferedType.EnumInt, instanceWithInferedType, probe.formStage, formGroup)
		AssociationFieldToForm("B", instanceWithInferedType.B, formGroup, probe)
		AssociationSliceToForm("Bs", instanceWithInferedType, &instanceWithInferedType.Bs, formGroup, probe)
		AssociationFieldToForm("C", instanceWithInferedType.C, formGroup, probe)
		AssociationSliceToForm("Cs", instanceWithInferedType, &instanceWithInferedType.Cs, formGroup, probe)
		BasicFieldtoForm("UUID", instanceWithInferedType.UUID, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.B:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A, *models.B](
				"A",
				"Bs",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A) []*models.B {
					return owner.Bs
				})
		}

	case *models.C:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0, false)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.A, *models.C](
				"A",
				"Cs",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.A) []*models.C {
					return owner.Cs
				})
		}

	default:
		_ = instanceWithInferedType
	}
}
