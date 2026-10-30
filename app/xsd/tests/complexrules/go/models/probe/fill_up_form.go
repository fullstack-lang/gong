// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/app/xsd/tests/complexrules/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/complexrules/go/models/refs"
	"github.com/fullstack-lang/gong/app/xsd/tests/complexrules/go/models/common"
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
	case *models.Catalog:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Target
			div := (&form.FormDiv{Name: "Target"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Target {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Target",
				Label: "Target",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		AssociationSliceToForm("Item", inst, &inst.Item, formGroup, probe)
	case *models.DerivedItemType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Code", inst.Code, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Score", inst.Score, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("TargetRef", inst.TargetRef, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OptionA", inst.OptionA, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("OptionB", inst.OptionB, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Catalogs {
				for _, target := range src.Item {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Catalog", "Item", refNames, formGroup, probe.formStage)
		}
	case *refs.RefTargetType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Id", inst.Id, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Status", inst.Status, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("CreatedBy", inst.CreatedBy, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Label", inst.Label, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageOfInterest.Catalogs {
				for _, target := range src.Target {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Catalog", "Target", refNames, formGroup, probe.formStage)
		}
	case *common.BaseItemType:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Code", inst.Code, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Score", inst.Score, probe.formStage, formGroup)
	default:
		_ = inst
	}
}
