// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/tree/go/models"
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
		StageSetBasicFieldtoForm("Icon", inst.Icon, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("SVGIcon", inst.SVGIcon, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.SVGIcon](), probe.formStage)
		StageSetBasicFieldtoForm("IsDisabled", inst.IsDisabled, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasToolTip", inst.HasToolTip, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ToolTipText", inst.ToolTipText, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("ToolTipPosition", inst.ToolTipPosition, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("ClientOnX", inst.ClientOnX, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ClientOnY", inst.ClientOnY, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Menus {
				for _, target := range src.Buttons {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Menu", "Buttons", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Nodes {
				for _, target := range src.Buttons {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Node", "Buttons", refNames, formGroup, probe.formStage)
		}
	case *models.Menu:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

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

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Nodes {
				if src.Menu == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Node", "Menu", refNames, formGroup, probe.formStage)
		}
	case *models.Node:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsWithPrefix", inst.IsWithPrefix, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Prefix", inst.Prefix, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("FontStyle", inst.FontStyle, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("BackgroundColor", inst.BackgroundColor, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsExpanded", inst.IsExpanded, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasCheckboxButton", inst.HasCheckboxButton, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsCheckboxDisabled", inst.IsCheckboxDisabled, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CheckboxHasToolTip", inst.CheckboxHasToolTip, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("CheckboxToolTipText", inst.CheckboxToolTipText, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("CheckboxToolTipPosition", inst.CheckboxToolTipPosition, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("HasSecondCheckboxButton", inst.HasSecondCheckboxButton, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsSecondCheckboxChecked", inst.IsSecondCheckboxChecked, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsSecondCheckboxDisabled", inst.IsSecondCheckboxDisabled, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SecondCheckboxHasToolTip", inst.SecondCheckboxHasToolTip, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SecondCheckboxToolTipText", inst.SecondCheckboxToolTipText, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("SecondCheckboxToolTipPosition", inst.SecondCheckboxToolTipPosition, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("TextAfterSecondCheckbox", inst.TextAfterSecondCheckbox, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasToolTip", inst.HasToolTip, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ToolTipText", inst.ToolTipText, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("ToolTipPosition", inst.ToolTipPosition, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("ClientOnY", inst.ClientOnY, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsInEditMode", inst.IsInEditMode, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsNodeClickable", inst.IsNodeClickable, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsWithPreceedingIcon", inst.IsWithPreceedingIcon, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("PreceedingIcon", inst.PreceedingIcon, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("PreceedingSVGIcon", inst.PreceedingSVGIcon, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.SVGIcon](), probe.formStage)

		{
			// Slice of pointers: Children
			div := (&form.FormDiv{Name: "Children"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Children {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Children",
				Label: "Children",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

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
		StageSetAssociationFieldToForm("Menu", inst.Menu, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Menu](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Nodes {
				for _, target := range src.Children {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Node", "Children", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Trees {
				for _, target := range src.RootNodes {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Tree", "RootNodes", refNames, formGroup, probe.formStage)
		}
	case *models.SVGIcon:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SVG", inst.SVG, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Buttons {
				if src.SVGIcon == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Button", "SVGIcon", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Nodes {
				if src.PreceedingSVGIcon == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Node", "PreceedingSVGIcon", refNames, formGroup, probe.formStage)
		}
	case *models.Tree:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: RootNodes
			div := (&form.FormDiv{Name: "RootNodes"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RootNodes {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RootNodes",
				Label: "RootNodes",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("HaveSearch", inst.HaveSearch, probe.formStage, formGroup)
	default:
		_ = inst
	}
}
