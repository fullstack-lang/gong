// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/slider/go/models"
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
	case *models.Checkbox:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ValueBool", inst.ValueBool, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LabelForTrue", inst.LabelForTrue, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("LabelForFalse", inst.LabelForFalse, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Groups {
				for _, target := range src.Checkboxes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Group", "Checkboxes", refNames, formGroup, probe.formStage)
		}
	case *models.Group:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Percentage", inst.Percentage, probe.formStage, formGroup)

		{
			// Slice of pointers: Sliders
			div := (&form.FormDiv{Name: "Sliders"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Sliders {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Sliders",
				Label: "Sliders",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Checkboxes
			div := (&form.FormDiv{Name: "Checkboxes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Checkboxes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Checkboxes",
				Label: "Checkboxes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Layouts {
				for _, target := range src.Groups {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Layout", "Groups", refNames, formGroup, probe.formStage)
		}
	case *models.Layout:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Groups
			div := (&form.FormDiv{Name: "Groups"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Groups {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Groups",
				Label: "Groups",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsWithCustomGutterSize", inst.IsWithCustomGutterSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("GutterSize", inst.GutterSize, probe.formStage, formGroup)
	case *models.Slider:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsFloat64", inst.IsFloat64, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsInt", inst.IsInt, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MinInt", inst.MinInt, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MaxInt", inst.MaxInt, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StepInt", inst.StepInt, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ValueInt", inst.ValueInt, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MinFloat64", inst.MinFloat64, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("MaxFloat64", inst.MaxFloat64, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StepFloat64", inst.StepFloat64, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ValueFloat64", inst.ValueFloat64, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsDisabled", inst.IsDisabled, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Groups {
				for _, target := range src.Sliders {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Group", "Sliders", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
