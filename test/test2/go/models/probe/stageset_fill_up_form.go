// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/fullstack-lang/gong/test/test2/go/models/x"
	"github.com/fullstack-lang/gong/test/test2/go/models/y"
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
		StageSetBasicFieldtoForm("NumberField", inst.NumberField, probe.formStage, formGroup)
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
		StageSetAssociationFieldToForm("X", inst.X, formGroup, probe.stageSet.XStage.GetInstancesSet[*x.X](), probe.formStage)
		StageSetBasicFieldtoForm("Foo", inst.Foo, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bar", inst.Bar, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Zorgh", inst.Zorgh, probe.formStage, formGroup)
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
	case *x.X:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("Y", inst.Y, formGroup, probe.stageSet.YStage.GetInstancesSet[*y.Y](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.As {
				if src.X == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A", "X", refNames, formGroup, probe.formStage)
		}
	case *y.Y:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.XStage.Xs {
				if src.Y == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("x.X", "Y", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
