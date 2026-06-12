// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/form/go/models"

	"github.com/fullstack-lang/gong/lib/table/go/models"
)

const FormName = "Form"

func FillUpForm(
	instance any,
	formGroup *form.FormGroup,
	probe *Probe,
) {

	switch instanceWithInferedType := any(instance).(type) {
	// insertion point
	case *models.Button:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Icon", instanceWithInferedType.Icon, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("SVGIcon", instanceWithInferedType.SVGIcon, formGroup, probe)
		BasicFieldtoForm("IsDisabled", instanceWithInferedType.IsDisabled, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasToolTip", instanceWithInferedType.HasToolTip, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ToolTipText", instanceWithInferedType.ToolTipText, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		EnumTypeStringToForm("ToolTipPosition", instanceWithInferedType.ToolTipPosition, instanceWithInferedType, probe.formStage, formGroup)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Table, *models.Button](
				"Table",
				"Buttons",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Table) []*models.Button {
					return owner.Buttons
				})
		}

	case *models.Cell:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationFieldToForm("CellString", instanceWithInferedType.CellString, formGroup, probe)
		AssociationFieldToForm("CellFloat64", instanceWithInferedType.CellFloat64, formGroup, probe)
		AssociationFieldToForm("CellInt", instanceWithInferedType.CellInt, formGroup, probe)
		AssociationFieldToForm("CellBool", instanceWithInferedType.CellBool, formGroup, probe)
		AssociationFieldToForm("CellIcon", instanceWithInferedType.CellIcon, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Row, *models.Cell](
				"Row",
				"Cells",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Row) []*models.Cell {
					return owner.Cells
				})
		}

	case *models.CellBoolean:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.CellFloat64:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.CellIcon:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Icon", instanceWithInferedType.Icon, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NeedsConfirmation", instanceWithInferedType.NeedsConfirmation, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("ConfirmationMessage", instanceWithInferedType.ConfirmationMessage, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.CellInt:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.CellString:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("Value", instanceWithInferedType.Value, instanceWithInferedType, probe.formStage, formGroup,
			true, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.DisplayedColumn:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Table, *models.DisplayedColumn](
				"Table",
				"DisplayedColumns",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Table) []*models.DisplayedColumn {
					return owner.DisplayedColumns
				})
		}

	case *models.Row:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Cells", instanceWithInferedType, &instanceWithInferedType.Cells, formGroup, probe)
		BasicFieldtoForm("IsChecked", instanceWithInferedType.IsChecked, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)
		{
			AssociationReverseSliceToForm[*models.Table, *models.Row](
				"Table",
				"Rows",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Table) []*models.Row {
					return owner.Rows
				})
		}
		{
			AssociationReverseSliceToForm[*models.Table, *models.Row](
				"Table",
				"RowsSelectedForBulkDelete",
				instanceWithInferedType,
				formGroup,
				probe,
				func(owner *models.Table) []*models.Row {
					return owner.RowsSelectedForBulkDelete
				})
		}

	case *models.SVGIcon:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, false, 0)
		BasicFieldtoForm("SVG", instanceWithInferedType.SVG, instanceWithInferedType, probe.formStage, formGroup,
			false, true, 600, true, 300)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	case *models.Table:
		// insertion point
		BasicFieldtoForm("Name", instanceWithInferedType.Name, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("DisplayedColumns", instanceWithInferedType, &instanceWithInferedType.DisplayedColumns, formGroup, probe)
		AssociationSliceToForm("Rows", instanceWithInferedType, &instanceWithInferedType.Rows, formGroup, probe)
		BasicFieldtoForm("HasFiltering", instanceWithInferedType.HasFiltering, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasColumnSorting", instanceWithInferedType.HasColumnSorting, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasPaginator", instanceWithInferedType.HasPaginator, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasCheckableRows", instanceWithInferedType.HasCheckableRows, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasSaveButton", instanceWithInferedType.HasSaveButton, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SaveButtonLabel", instanceWithInferedType.SaveButtonLabel, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasBulkDeleteButton", instanceWithInferedType.HasBulkDeleteButton, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("BulkDeleteButtonTooltip", instanceWithInferedType.BulkDeleteButtonTooltip, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("RowsSelectedForBulkDelete", instanceWithInferedType, &instanceWithInferedType.RowsSelectedForBulkDelete, formGroup, probe)
		BasicFieldtoForm("CanDragDropRows", instanceWithInferedType.CanDragDropRows, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("HasCloseButton", instanceWithInferedType.HasCloseButton, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("SavingInProgress", instanceWithInferedType.SavingInProgress, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		BasicFieldtoForm("NbOfStickyColumns", instanceWithInferedType.NbOfStickyColumns, instanceWithInferedType, probe.formStage, formGroup,
			false, false, 0, false, 0)
		AssociationSliceToForm("Buttons", instanceWithInferedType, &instanceWithInferedType.Buttons, formGroup, probe)
		formDivDivider := (&form.FormDiv{
			Name:       "",
			IsADivider: true,
		}).Stage(probe.formStage)
		formGroup.FormDivs = append(formGroup.FormDivs, formDivDivider)

	default:
		_ = instanceWithInferedType
	}
}
