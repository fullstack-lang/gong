package main

import (
	"slices"
	"time"

	"github.com/fullstack-lang/gong/lib/table/go/models"
	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time
var _ = slices.Index[[]int, int]

// function will stage objects
func _(stage *models.Stage) {

	// insertion point for declaration of instances to stage

	__Button__00000000_ := (&models.Button{Name: `Delete`}).Stage(stage)

	__Cell__00000000_ := (&models.Cell{Name: `Row 1 - Cell 1`}).Stage(stage)
	__Cell__00000001_ := (&models.Cell{Name: `Row 1 - Cell 2`}).Stage(stage)
	__Cell__00000002_ := (&models.Cell{Name: `Row 1 - Cell 3`}).Stage(stage)
	__Cell__00000003_ := (&models.Cell{Name: `Row 1 - Cell 4`}).Stage(stage)
	__Cell__00000004_ := (&models.Cell{Name: `Row 1 - Cell 5`}).Stage(stage)
	__Cell__00000005_ := (&models.Cell{Name: `Row 2 - Cell 1`}).Stage(stage)
	__Cell__00000006_ := (&models.Cell{Name: `Row 2 - Cell 2`}).Stage(stage)
	__Cell__00000007_ := (&models.Cell{Name: `Row 2 - Cell 3`}).Stage(stage)
	__Cell__00000008_ := (&models.Cell{Name: `Row 2 - Cell 4`}).Stage(stage)
	__Cell__00000009_ := (&models.Cell{Name: `Row 2 - Cell 5`}).Stage(stage)
	__Cell__00000010_ := (&models.Cell{Name: `T2 - R1 - A`}).Stage(stage)
	__Cell__00000011_ := (&models.Cell{Name: `T2 - R1 - B`}).Stage(stage)
	__Cell__00000012_ := (&models.Cell{Name: `T2 - R2 - A`}).Stage(stage)
	__Cell__00000013_ := (&models.Cell{Name: `T2 - R2 - B`}).Stage(stage)
	__Cell__00000014_ := (&models.Cell{Name: `T2 - R3 - A`}).Stage(stage)
	__Cell__00000015_ := (&models.Cell{Name: `T2 - R3 - B`}).Stage(stage)

	__CellBoolean__00000000_ := (&models.CellBoolean{Name: `Row 1 -Cell 1 - Cell False`}).Stage(stage)
	__CellBoolean__00000001_ := (&models.CellBoolean{Name: `Row 2 - Cell 1 - Cell true`}).Stage(stage)

	__CellFloat64__00000000_ := (&models.CellFloat64{Name: `Row 1 - Cell2 - Float`}).Stage(stage)
	__CellFloat64__00000001_ := (&models.CellFloat64{Name: `Row 2 - Cell 2 - Float`}).Stage(stage)

	__CellIcon__00000000_ := (&models.CellIcon{Name: `Row 1 - Cell 3 - Delete`}).Stage(stage)
	__CellIcon__00000001_ := (&models.CellIcon{Name: `Row 2 - Cell 3 - home`}).Stage(stage)

	__CellInt__00000000_ := (&models.CellInt{Name: `Row 1 - Cell 4 - Int`}).Stage(stage)
	__CellInt__00000001_ := (&models.CellInt{Name: `Row 2 - Cell 4 - Int`}).Stage(stage)

	__CellString__00000000_ := (&models.CellString{Name: `Row 1 - Cell 5`}).Stage(stage)
	__CellString__00000001_ := (&models.CellString{Name: `Row 2 - Cell 5`}).Stage(stage)
	__CellString__00000002_ := (&models.CellString{Name: `T2 - R1 - A`}).Stage(stage)
	__CellString__00000003_ := (&models.CellString{Name: `T2 - R1 - B`}).Stage(stage)
	__CellString__00000004_ := (&models.CellString{Name: `T2 - R2 - A`}).Stage(stage)
	__CellString__00000005_ := (&models.CellString{Name: `T2 - R2 - B`}).Stage(stage)
	__CellString__00000006_ := (&models.CellString{Name: `T2 - R3 - A`}).Stage(stage)
	__CellString__00000007_ := (&models.CellString{Name: `T2 - R3 - B`}).Stage(stage)

	__DisplayedColumn__00000000_ := (&models.DisplayedColumn{Name: `A`}).Stage(stage)
	__DisplayedColumn__00000001_ := (&models.DisplayedColumn{Name: `B`}).Stage(stage)
	__DisplayedColumn__00000002_ := (&models.DisplayedColumn{Name: `Column 1 - Boolean`}).Stage(stage)
	__DisplayedColumn__00000003_ := (&models.DisplayedColumn{Name: `Column 2 - Float64`}).Stage(stage)
	__DisplayedColumn__00000004_ := (&models.DisplayedColumn{Name: `Column 3 - Icon`}).Stage(stage)
	__DisplayedColumn__00000005_ := (&models.DisplayedColumn{Name: `Column 4 - Int`}).Stage(stage)
	__DisplayedColumn__00000006_ := (&models.DisplayedColumn{Name: `Column 5 - String`}).Stage(stage)

	__Row__00000000_ := (&models.Row{Name: `Row 1`}).Stage(stage)
	__Row__00000001_ := (&models.Row{Name: `Row 2`}).Stage(stage)
	__Row__00000002_ := (&models.Row{Name: `Table 2 - Row 1`}).Stage(stage)
	__Row__00000003_ := (&models.Row{Name: `Table 2 - Row 2`}).Stage(stage)
	__Row__00000004_ := (&models.Row{Name: `Table 2 - Row 3`}).Stage(stage)

	__Table__00000000_ := (&models.Table{Name: `EmptyTable`}).Stage(stage)
	__Table__00000001_ := (&models.Table{Name: `Table 2`}).Stage(stage)
	__Table__00000002_ := (&models.Table{Name: `Table with 5 types`}).Stage(stage)

	// insertion point for initialization of values

	__Button__00000000_.Name = `Delete`
	__Button__00000000_.Icon = `delete`
	__Button__00000000_.IsDisabled = false
	__Button__00000000_.HasToolTip = false
	__Button__00000000_.ToolTipText = `test to delete`
	__Button__00000000_.ToolTipPosition = models.Above

	__Cell__00000000_.Name = `Row 1 - Cell 1`

	__Cell__00000001_.Name = `Row 1 - Cell 2`

	__Cell__00000002_.Name = `Row 1 - Cell 3`

	__Cell__00000003_.Name = `Row 1 - Cell 4`

	__Cell__00000004_.Name = `Row 1 - Cell 5`

	__Cell__00000005_.Name = `Row 2 - Cell 1`

	__Cell__00000006_.Name = `Row 2 - Cell 2`

	__Cell__00000007_.Name = `Row 2 - Cell 3`

	__Cell__00000008_.Name = `Row 2 - Cell 4`

	__Cell__00000009_.Name = `Row 2 - Cell 5`

	__Cell__00000010_.Name = `T2 - R1 - A`

	__Cell__00000011_.Name = `T2 - R1 - B`

	__Cell__00000012_.Name = `T2 - R2 - A`

	__Cell__00000013_.Name = `T2 - R2 - B`

	__Cell__00000014_.Name = `T2 - R3 - A`

	__Cell__00000015_.Name = `T2 - R3 - B`

	__CellBoolean__00000000_.Name = `Row 1 -Cell 1 - Cell False`
	__CellBoolean__00000000_.Value = false

	__CellBoolean__00000001_.Name = `Row 2 - Cell 1 - Cell true`
	__CellBoolean__00000001_.Value = true

	__CellFloat64__00000000_.Name = `Row 1 - Cell2 - Float`
	__CellFloat64__00000000_.Value = 20.433333

	__CellFloat64__00000001_.Name = `Row 2 - Cell 2 - Float`
	__CellFloat64__00000001_.Value = 18.550000

	__CellIcon__00000000_.Name = `Row 1 - Cell 3 - Delete`
	__CellIcon__00000000_.Icon = `delete`
	__CellIcon__00000000_.NeedsConfirmation = true
	__CellIcon__00000000_.ConfirmationMessage = `Do you confirm deletion`

	__CellIcon__00000001_.Name = `Row 2 - Cell 3 - home`
	__CellIcon__00000001_.Icon = `edit`
	__CellIcon__00000001_.NeedsConfirmation = false
	__CellIcon__00000001_.ConfirmationMessage = ``

	__CellInt__00000000_.Name = `Row 1 - Cell 4 - Int`
	__CellInt__00000000_.Value = 10

	__CellInt__00000001_.Name = `Row 2 - Cell 4 - Int`
	__CellInt__00000001_.Value = 288

	__CellString__00000000_.Name = `Row 1 - Cell 5`
	__CellString__00000000_.Value = `Je ferais le métier  des idoles antiques
Et comme elles je veux me faire redorer`

	__CellString__00000001_.Name = `Row 2 - Cell 5`
	__CellString__00000001_.Value = `des idoles antiques`

	__CellString__00000002_.Name = `T2 - R1 - A`
	__CellString__00000002_.Value = `T2 - R1 - A`

	__CellString__00000003_.Name = `T2 - R1 - B`
	__CellString__00000003_.Value = `T2 - R1 - B`

	__CellString__00000004_.Name = `T2 - R2 - A`
	__CellString__00000004_.Value = `T2 - R2 - A`

	__CellString__00000005_.Name = `T2 - R2 - B`
	__CellString__00000005_.Value = `T2 - R2 - B`

	__CellString__00000006_.Name = `T2 - R3 - A`
	__CellString__00000006_.Value = `T2 - R3 - A`

	__CellString__00000007_.Name = `T2 - R3 - B`
	__CellString__00000007_.Value = `T2 - R3 - B`

	__DisplayedColumn__00000000_.Name = `A`

	__DisplayedColumn__00000001_.Name = `B`

	__DisplayedColumn__00000002_.Name = `Column 1 - Boolean`

	__DisplayedColumn__00000003_.Name = `Column 2 - Float64`

	__DisplayedColumn__00000004_.Name = `Column 3 - Icon`

	__DisplayedColumn__00000005_.Name = `Column 4 - Int`

	__DisplayedColumn__00000006_.Name = `Column 5 - String`

	__Row__00000000_.Name = `Row 1`
	__Row__00000000_.IsChecked = true

	__Row__00000001_.Name = `Row 2`
	__Row__00000001_.IsChecked = false

	__Row__00000002_.Name = `Table 2 - Row 1`
	__Row__00000002_.IsChecked = false

	__Row__00000003_.Name = `Table 2 - Row 2`
	__Row__00000003_.IsChecked = false

	__Row__00000004_.Name = `Table 2 - Row 3`
	__Row__00000004_.IsChecked = false

	__Table__00000000_.Name = `EmptyTable`
	__Table__00000000_.HasFiltering = false
	__Table__00000000_.HasColumnSorting = false
	__Table__00000000_.HasPaginator = false
	__Table__00000000_.HasCheckableRows = false
	__Table__00000000_.HasSaveButton = false
	__Table__00000000_.SaveButtonLabel = ``
	__Table__00000000_.CanDragDropRows = false
	__Table__00000000_.HasCloseButton = false
	__Table__00000000_.SavingInProgress = false
	__Table__00000000_.NbOfStickyColumns = 0

	__Table__00000001_.Name = `Table 2`
	__Table__00000001_.HasFiltering = false
	__Table__00000001_.HasColumnSorting = false
	__Table__00000001_.HasPaginator = false
	__Table__00000001_.HasCheckableRows = false
	__Table__00000001_.HasSaveButton = true
	__Table__00000001_.SaveButtonLabel = ``
	__Table__00000001_.CanDragDropRows = true
	__Table__00000001_.HasCloseButton = true
	__Table__00000001_.SavingInProgress = false
	__Table__00000001_.NbOfStickyColumns = 0

	__Table__00000002_.Name = `Table with 5 types`
	__Table__00000002_.HasFiltering = true
	__Table__00000002_.HasColumnSorting = true
	__Table__00000002_.HasPaginator = true
	__Table__00000002_.HasCheckableRows = true
	__Table__00000002_.HasSaveButton = true
	__Table__00000002_.SaveButtonLabel = ``
	__Table__00000002_.CanDragDropRows = false
	__Table__00000002_.HasCloseButton = false
	__Table__00000002_.SavingInProgress = false
	__Table__00000002_.NbOfStickyColumns = 3

	// insertion point for setup of pointers
	__Button__00000000_.SVGIcon = nil
	__Cell__00000000_.CellString = nil
	__Cell__00000000_.CellFloat64 = nil
	__Cell__00000000_.CellInt = nil
	__Cell__00000000_.CellBool = __CellBoolean__00000000_
	__Cell__00000000_.CellIcon = nil
	__Cell__00000001_.CellString = nil
	__Cell__00000001_.CellFloat64 = __CellFloat64__00000000_
	__Cell__00000001_.CellInt = nil
	__Cell__00000001_.CellBool = nil
	__Cell__00000001_.CellIcon = nil
	__Cell__00000002_.CellString = nil
	__Cell__00000002_.CellFloat64 = nil
	__Cell__00000002_.CellInt = nil
	__Cell__00000002_.CellBool = nil
	__Cell__00000002_.CellIcon = __CellIcon__00000000_
	__Cell__00000003_.CellString = nil
	__Cell__00000003_.CellFloat64 = nil
	__Cell__00000003_.CellInt = __CellInt__00000000_
	__Cell__00000003_.CellBool = nil
	__Cell__00000003_.CellIcon = nil
	__Cell__00000004_.CellString = __CellString__00000000_
	__Cell__00000004_.CellFloat64 = nil
	__Cell__00000004_.CellInt = nil
	__Cell__00000004_.CellBool = nil
	__Cell__00000004_.CellIcon = nil
	__Cell__00000005_.CellString = nil
	__Cell__00000005_.CellFloat64 = nil
	__Cell__00000005_.CellInt = nil
	__Cell__00000005_.CellBool = __CellBoolean__00000001_
	__Cell__00000005_.CellIcon = nil
	__Cell__00000006_.CellString = nil
	__Cell__00000006_.CellFloat64 = __CellFloat64__00000001_
	__Cell__00000006_.CellInt = nil
	__Cell__00000006_.CellBool = nil
	__Cell__00000006_.CellIcon = nil
	__Cell__00000007_.CellString = nil
	__Cell__00000007_.CellFloat64 = nil
	__Cell__00000007_.CellInt = nil
	__Cell__00000007_.CellBool = nil
	__Cell__00000007_.CellIcon = __CellIcon__00000001_
	__Cell__00000008_.CellString = nil
	__Cell__00000008_.CellFloat64 = nil
	__Cell__00000008_.CellInt = __CellInt__00000001_
	__Cell__00000008_.CellBool = nil
	__Cell__00000008_.CellIcon = nil
	__Cell__00000009_.CellString = __CellString__00000001_
	__Cell__00000009_.CellFloat64 = nil
	__Cell__00000009_.CellInt = nil
	__Cell__00000009_.CellBool = nil
	__Cell__00000009_.CellIcon = nil
	__Cell__00000010_.CellString = __CellString__00000002_
	__Cell__00000010_.CellFloat64 = nil
	__Cell__00000010_.CellInt = nil
	__Cell__00000010_.CellBool = nil
	__Cell__00000010_.CellIcon = nil
	__Cell__00000011_.CellString = __CellString__00000003_
	__Cell__00000011_.CellFloat64 = nil
	__Cell__00000011_.CellInt = nil
	__Cell__00000011_.CellBool = nil
	__Cell__00000011_.CellIcon = nil
	__Cell__00000012_.CellString = __CellString__00000004_
	__Cell__00000012_.CellFloat64 = nil
	__Cell__00000012_.CellInt = nil
	__Cell__00000012_.CellBool = nil
	__Cell__00000012_.CellIcon = nil
	__Cell__00000013_.CellString = __CellString__00000005_
	__Cell__00000013_.CellFloat64 = nil
	__Cell__00000013_.CellInt = nil
	__Cell__00000013_.CellBool = nil
	__Cell__00000013_.CellIcon = nil
	__Cell__00000014_.CellString = __CellString__00000006_
	__Cell__00000014_.CellFloat64 = nil
	__Cell__00000014_.CellInt = nil
	__Cell__00000014_.CellBool = nil
	__Cell__00000014_.CellIcon = nil
	__Cell__00000015_.CellString = __CellString__00000007_
	__Cell__00000015_.CellFloat64 = nil
	__Cell__00000015_.CellInt = nil
	__Cell__00000015_.CellBool = nil
	__Cell__00000015_.CellIcon = nil
	__Row__00000000_.Cells = append(__Row__00000000_.Cells, __Cell__00000000_)
	__Row__00000000_.Cells = append(__Row__00000000_.Cells, __Cell__00000001_)
	__Row__00000000_.Cells = append(__Row__00000000_.Cells, __Cell__00000002_)
	__Row__00000000_.Cells = append(__Row__00000000_.Cells, __Cell__00000003_)
	__Row__00000000_.Cells = append(__Row__00000000_.Cells, __Cell__00000004_)
	__Row__00000001_.Cells = append(__Row__00000001_.Cells, __Cell__00000005_)
	__Row__00000001_.Cells = append(__Row__00000001_.Cells, __Cell__00000006_)
	__Row__00000001_.Cells = append(__Row__00000001_.Cells, __Cell__00000007_)
	__Row__00000001_.Cells = append(__Row__00000001_.Cells, __Cell__00000008_)
	__Row__00000001_.Cells = append(__Row__00000001_.Cells, __Cell__00000009_)
	__Row__00000002_.Cells = append(__Row__00000002_.Cells, __Cell__00000011_)
	__Row__00000002_.Cells = append(__Row__00000002_.Cells, __Cell__00000010_)
	__Row__00000003_.Cells = append(__Row__00000003_.Cells, __Cell__00000013_)
	__Row__00000003_.Cells = append(__Row__00000003_.Cells, __Cell__00000012_)
	__Row__00000004_.Cells = append(__Row__00000004_.Cells, __Cell__00000014_)
	__Row__00000004_.Cells = append(__Row__00000004_.Cells, __Cell__00000015_)
	__Table__00000001_.DisplayedColumns = append(__Table__00000001_.DisplayedColumns, __DisplayedColumn__00000000_)
	__Table__00000001_.Rows = append(__Table__00000001_.Rows, __Row__00000003_)
	__Table__00000001_.Rows = append(__Table__00000001_.Rows, __Row__00000002_)
	__Table__00000001_.Rows = append(__Table__00000001_.Rows, __Row__00000004_)
	__Table__00000002_.DisplayedColumns = append(__Table__00000002_.DisplayedColumns, __DisplayedColumn__00000002_)
	__Table__00000002_.DisplayedColumns = append(__Table__00000002_.DisplayedColumns, __DisplayedColumn__00000003_)
	__Table__00000002_.DisplayedColumns = append(__Table__00000002_.DisplayedColumns, __DisplayedColumn__00000004_)
	__Table__00000002_.DisplayedColumns = append(__Table__00000002_.DisplayedColumns, __DisplayedColumn__00000005_)
	__Table__00000002_.DisplayedColumns = append(__Table__00000002_.DisplayedColumns, __DisplayedColumn__00000006_)
	__Table__00000002_.Rows = append(__Table__00000002_.Rows, __Row__00000000_)
	__Table__00000002_.Rows = append(__Table__00000002_.Rows, __Row__00000001_)
	__Table__00000002_.Buttons = append(__Table__00000002_.Buttons, __Button__00000000_)
}
