// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
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
	case *models.AsSplit:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Direction", inst.Direction, formGroup, probe.formStage)

		{
			// Slice of pointers: AsSplitAreas
			div := (&form.FormDiv{Name: "AsSplitAreas"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.AsSplitAreas {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "AsSplitAreas",
				Label: "AsSplitAreas",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsSizeInPixel", inst.IsSizeInPixel, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsWithCustomGutterSize", inst.IsWithCustomGutterSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("GutterSize", inst.GutterSize, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.AsSplit == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AsSplitArea", "AsSplit", refNames, formGroup, probe.formStage)
		}
	case *models.AsSplitArea:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ShowNameInHeader", inst.ShowNameInHeader, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Size", inst.Size, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsAny", inst.IsAny, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("AsSplit", inst.AsSplit, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.AsSplit](), probe.formStage)
		StageSetAssociationFieldToForm("Button", inst.Button, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Button](), probe.formStage)
		StageSetAssociationFieldToForm("Cursor", inst.Cursor, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Cursor](), probe.formStage)
		StageSetAssociationFieldToForm("Form", inst.Form, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Form](), probe.formStage)
		StageSetAssociationFieldToForm("Load", inst.Load, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Load](), probe.formStage)
		StageSetAssociationFieldToForm("Markdown", inst.Markdown, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Markdown](), probe.formStage)
		StageSetAssociationFieldToForm("Slider", inst.Slider, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Slider](), probe.formStage)
		StageSetAssociationFieldToForm("Split", inst.Split, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Split](), probe.formStage)
		StageSetAssociationFieldToForm("Svg", inst.Svg, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Svg](), probe.formStage)
		StageSetAssociationFieldToForm("Table", inst.Table, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Table](), probe.formStage)
		StageSetAssociationFieldToForm("Tone", inst.Tone, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Tone](), probe.formStage)
		StageSetAssociationFieldToForm("Tree", inst.Tree, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Tree](), probe.formStage)
		StageSetAssociationFieldToForm("Threejs", inst.Threejs, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Threejs](), probe.formStage)
		StageSetAssociationFieldToForm("Xlsx", inst.Xlsx, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.Xlsx](), probe.formStage)
		StageSetBasicFieldtoForm("HasDiv", inst.HasDiv, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("DivStyle", inst.DivStyle, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplits {
				for _, target := range src.AsSplitAreas {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AsSplit", "AsSplitAreas", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Views {
				for _, target := range src.RootAsSplitAreas {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.View", "RootAsSplitAreas", refNames, formGroup, probe.formStage)
		}
	case *models.Button:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StackName", inst.StackName, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Button == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AsSplitArea", "Button", refNames, formGroup, probe.formStage)
		}
	case *models.Cursor:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StackName", inst.StackName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Style", inst.Style, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Cursor == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AsSplitArea", "Cursor", refNames, formGroup, probe.formStage)
		}
	case *models.FavIcon:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SVG", inst.SVG, probe.formStage, formGroup)
	case *models.Form:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StackName", inst.StackName, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Form == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AsSplitArea", "Form", refNames, formGroup, probe.formStage)
		}
	case *models.Load:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StackName", inst.StackName, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Load == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AsSplitArea", "Load", refNames, formGroup, probe.formStage)
		}
	case *models.LogoOnTheLeft:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SVG", inst.SVG, probe.formStage, formGroup)
	case *models.LogoOnTheRight:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Width", inst.Width, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Height", inst.Height, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SVG", inst.SVG, probe.formStage, formGroup)
	case *models.Markdown:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StackName", inst.StackName, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Markdown == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AsSplitArea", "Markdown", refNames, formGroup, probe.formStage)
		}
	case *models.Slider:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StackName", inst.StackName, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Slider == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AsSplitArea", "Slider", refNames, formGroup, probe.formStage)
		}
	case *models.Split:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StackName", inst.StackName, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Split == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AsSplitArea", "Split", refNames, formGroup, probe.formStage)
		}
	case *models.Svg:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StackName", inst.StackName, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Style", inst.Style, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Svg == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AsSplitArea", "Svg", refNames, formGroup, probe.formStage)
		}
	case *models.Table:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StackName", inst.StackName, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Table == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AsSplitArea", "Table", refNames, formGroup, probe.formStage)
		}
	case *models.Threejs:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StackName", inst.StackName, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Threejs == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AsSplitArea", "Threejs", refNames, formGroup, probe.formStage)
		}
	case *models.Title:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
	case *models.Tone:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StackName", inst.StackName, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Tone == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AsSplitArea", "Tone", refNames, formGroup, probe.formStage)
		}
	case *models.Tree:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StackName", inst.StackName, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Tree == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AsSplitArea", "Tree", refNames, formGroup, probe.formStage)
		}
	case *models.View:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ShowViewName", inst.ShowViewName, probe.formStage, formGroup)

		{
			// Slice of pointers: RootAsSplitAreas
			div := (&form.FormDiv{Name: "RootAsSplitAreas"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RootAsSplitAreas {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RootAsSplitAreas",
				Label: "RootAsSplitAreas",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsSelectedView", inst.IsSelectedView, probe.formStage, formGroup)
		StageSetEnumStringFieldToForm("Direction", inst.Direction, formGroup, probe.formStage)
		StageSetBasicFieldtoForm("IsSecondaryView", inst.IsSecondaryView, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsSizeInPixel", inst.IsSizeInPixel, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("IsWithCustomGutterSize", inst.IsWithCustomGutterSize, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("GutterSize", inst.GutterSize, probe.formStage, formGroup)
	case *models.Xlsx:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("StackName", inst.StackName, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.AsSplitAreas {
				if src.Xlsx == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.AsSplitArea", "Xlsx", refNames, formGroup, probe.formStage)
		}
	default:
		_ = inst
	}
}
