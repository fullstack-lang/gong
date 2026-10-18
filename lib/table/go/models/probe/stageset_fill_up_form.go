// generated code - do not edit
package probe

import (
	"sort"
	"strings"

	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/table/go/models"
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

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tables {
				for _, target := range src.Buttons {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Table", "Buttons", refNames, formGroup, probe.formStage)
		}
	case *models.Cell:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetAssociationFieldToForm("CellString", inst.CellString, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.CellString](), probe.formStage)
		StageSetAssociationFieldToForm("CellFloat64", inst.CellFloat64, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.CellFloat64](), probe.formStage)
		StageSetAssociationFieldToForm("CellInt", inst.CellInt, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.CellInt](), probe.formStage)
		StageSetAssociationFieldToForm("CellBool", inst.CellBool, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.CellBoolean](), probe.formStage)
		StageSetAssociationFieldToForm("CellIcon", inst.CellIcon, formGroup, probe.stageSet.Stage.GetInstancesSet[*models.CellIcon](), probe.formStage)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Rows {
				for _, target := range src.Cells {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Row", "Cells", refNames, formGroup, probe.formStage)
		}
	case *models.CellBoolean:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Cells {
				if src.CellBool == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Cell", "CellBool", refNames, formGroup, probe.formStage)
		}
	case *models.CellFloat64:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Cells {
				if src.CellFloat64 == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Cell", "CellFloat64", refNames, formGroup, probe.formStage)
		}
	case *models.CellIcon:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Icon", inst.Icon, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NeedsConfirmation", inst.NeedsConfirmation, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("ConfirmationMessage", inst.ConfirmationMessage, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Cells {
				if src.CellIcon == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Cell", "CellIcon", refNames, formGroup, probe.formStage)
		}
	case *models.CellInt:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Cells {
				if src.CellInt == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Cell", "CellInt", refNames, formGroup, probe.formStage)
		}
	case *models.CellString:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("Value", inst.Value, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Cells {
				if src.CellString == inst {
					refNames = append(refNames, src.GetName())
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Cell", "CellString", refNames, formGroup, probe.formStage)
		}
	case *models.DisplayedColumn:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tables {
				for _, target := range src.DisplayedColumns {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Table", "DisplayedColumns", refNames, formGroup, probe.formStage)
		}
	case *models.Row:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: Cells
			div := (&form.FormDiv{Name: "Cells"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Cells {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Cells",
				Label: "Cells",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("IsChecked", inst.IsChecked, probe.formStage, formGroup)

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tables {
				for _, target := range src.Rows {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Table", "Rows", refNames, formGroup, probe.formStage)
		}

		{
			var refNames []string
			for src := range probe.stageSet.Stage.Tables {
				for _, target := range src.RowsSelectedForBulkDelete {
					if target == inst {
						refNames = append(refNames, src.GetName())
						break
					}
				}
			}
			sort.Strings(refNames)
			StageSetAssociationReverseFieldToForm("models.Table", "RowsSelectedForBulkDelete", refNames, formGroup, probe.formStage)
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
	case *models.Table:
		StageSetBasicFieldtoForm("Name", inst.Name, probe.formStage, formGroup)

		{
			// Slice of pointers: DisplayedColumns
			div := (&form.FormDiv{Name: "DisplayedColumns"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.DisplayedColumns {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "DisplayedColumns",
				Label: "DisplayedColumns",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}

		{
			// Slice of pointers: Rows
			div := (&form.FormDiv{Name: "Rows"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.Rows {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "Rows",
				Label: "Rows",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("HasFiltering", inst.HasFiltering, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasColumnSorting", inst.HasColumnSorting, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasPaginator", inst.HasPaginator, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasCheckableRows", inst.HasCheckableRows, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasSaveButton", inst.HasSaveButton, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SaveButtonLabel", inst.SaveButtonLabel, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasBulkDeleteButton", inst.HasBulkDeleteButton, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("BulkDeleteButtonTooltip", inst.BulkDeleteButtonTooltip, probe.formStage, formGroup)

		{
			// Slice of pointers: RowsSelectedForBulkDelete
			div := (&form.FormDiv{Name: "RowsSelectedForBulkDelete"}).Stage(probe.formStage)
			formGroup.FormDivs = append(formGroup.FormDivs, div)
			var names []string
			for _, elem := range inst.RowsSelectedForBulkDelete {
				if elem != nil {
					names = append(names, elem.GetName())
				}
			}
			fld := (&form.FormField{
				Name:  "RowsSelectedForBulkDelete",
				Label: "RowsSelectedForBulkDelete",
				FormFieldString: &form.FormFieldString{
					Value: strings.Join(names, ", "),
				},
			}).Stage(probe.formStage)
			div.FormFields = append(div.FormFields, fld)
		}
		StageSetBasicFieldtoForm("CanDragDropRows", inst.CanDragDropRows, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("HasCloseButton", inst.HasCloseButton, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("SavingInProgress", inst.SavingInProgress, probe.formStage, formGroup)
		StageSetBasicFieldtoForm("NbOfStickyColumns", inst.NbOfStickyColumns, probe.formStage, formGroup)

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
	default:
		_ = inst
	}
}
