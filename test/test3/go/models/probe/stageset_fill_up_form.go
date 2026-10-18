// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test3/go/models"
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
	case *models.A:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Date", inst.Date, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Duration", inst.Duration, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("FloatValue", inst.FloatValue, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IntValue", inst.IntValue, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("EnumString", inst.EnumString, formGroup, probe.formStage)
		StageSetEnumIntFieldToForm("EnumInt", inst.EnumInt, formGroup, probe.formStage)
		StageSetAssociationFieldToForm("B", inst.B, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.B](), probe.formStage)

		{
			// Slice of pointers: Bs
			div := (&form.FormDiv{Name: "Bs"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Bs {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Bs",
				Label: "Bs",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetAssociationFieldToForm("C", inst.C, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.C](), probe.formStage)

		{
			// Slice of pointers: Cs
			div := (&form.FormDiv{Name: "Cs"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Cs {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Cs",
				Label: "Cs",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("UUID", inst.UUID, probe.formStage, formGroup)
	case *models.B:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.As {
				if src.B == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A", "B", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.As {
				for _, target := range src.Bs {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A", "Bs", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
