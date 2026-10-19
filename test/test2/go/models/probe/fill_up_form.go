// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/test/test2/go/models"
	"github.com/fullstack-lang/gong/test/test2/go/models/x"
	"github.com/fullstack-lang/gong/test/test2/go/models/y"
	model "github.com/fullstack-lang/gong/test/test2/go/models/x/models"
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
	case *models.A:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NumberField", inst.NumberField, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("B", inst.B, formGroup, probe.stageOfInterest.GetInstancesSet[*models.B](), probe.formStage)
		AssociationSliceToForm("Bs", inst, &inst.Bs, formGroup, probe)
		StageSetAssociationFieldToForm("X", inst.X, formGroup, probe.stageOfInterest.XStage.GetInstancesSet[*x.X](), probe.formStage)
		StageSetBasicFieldtoForm("Foo", inst.Foo, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Bar", inst.Bar, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Zorgh", inst.Zorgh, probe.formStage, formGroup)
	case *models.B:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.As {
				if src.B == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A", "B", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageOfInterest.As {
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
		StageSetAssociationFieldToForm("Y", inst.Y, formGroup, probe.stageOfInterest.YStage.GetInstancesSet[*y.Y](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.As {
				if src.X == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.A", "X", refNames, formGroup, probe.formStage)
		}
	case *y.Y:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SubModel", inst.SubModel, formGroup, probe.stageOfInterest.ModelStage.GetInstancesSet[*model.SubModel](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageOfInterest.XStage.Xs {
				if src.Y == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("x.X", "Y", refNames, formGroup, probe.formStage)
		}
	case *model.SubModel:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.YStage.Ys {
				if src.SubModel == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("y.Y", "SubModel", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
