// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/button/go/models"
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
	case *models.Button:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Label", inst.Label, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Icon", inst.Icon, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsDisabled", inst.IsDisabled, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Color", inst.Color, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("MatButtonType", inst.MatButtonType, formGroup, probe.formStage)
		StageSetEnumStringFieldToForm("MatButtonAppearance", inst.MatButtonAppearance, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("HasToolTip", inst.HasToolTip, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ToolTipText", inst.ToolTipText, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("ToolTipPosition", inst.ToolTipPosition, formGroup, probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Groups {
				for _, target := range src.Buttons {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Group", "Buttons", refNames, formGroup, probe.formStage)
		}
	case *models.ButtonToggle:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Label", inst.Label, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Icon", inst.Icon, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsDisabled", inst.IsDisabled, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.GroupToogles {
				for _, target := range src.ButtonToggles {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.GroupToogle", "ButtonToggles", refNames, formGroup, probe.formStage)
		}
	case *models.Group:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Percentage", inst.Percentage, probe.formStage, formGroup)

		{
			// Slice of pointers: Buttons
			div := (&form.FormDiv{Name: "Buttons"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Buttons {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Buttons",
				Label: "Buttons",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("NbColumns", inst.NbColumns, probe.formStage, formGroup)

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
	case *models.GroupToogle:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Percentage", inst.Percentage, probe.formStage, formGroup)

		{
			// Slice of pointers: ButtonToggles
			div := (&form.FormDiv{Name: "ButtonToggles"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.ButtonToggles {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "ButtonToggles",
				Label: "ButtonToggles",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsSingleSelector", inst.IsSingleSelector, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Layouts {
				for _, target := range src.GroupToogles {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Layout", "GroupToogles", refNames, formGroup, probe.formStage)
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

		{
			// Slice of pointers: GroupToogles
			div := (&form.FormDiv{Name: "GroupToogles"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.GroupToogles {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "GroupToogles",
				Label: "GroupToogles",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
	default:
		_ = inst
	}
}
