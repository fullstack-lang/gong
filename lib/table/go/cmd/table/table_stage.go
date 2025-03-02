package main

import (
	"time"

	"github.com/fullstack-lang/gongtable/go/models"

	// injection point for ident package import declaration{{ImportPackageDeclaration}}
)

// generated in order to avoid error in the package import
// if there are no elements in the stage to marshall
var _ time.Time

// _ point for meta package dummy declaration{{ImportPackageDummyDeclaration}}

// When parsed, those maps will help with the renaming process
var _ map[string]any = map[string]any{
	// injection point for docLink to identifiers{{EntriesDocLinkStringDocLinkIdentifier}}
}

// function will stage objects
func _(stage *models.StageStruct) {

	// Declaration of instances to stage

	__Cell__000000_Row_1_Cell_1 := (&models.Cell{}).Stage(stage)
	__Cell__000001_Row_1_Cell_2 := (&models.Cell{}).Stage(stage)
	__Cell__000002_Row_1_Cell_3 := (&models.Cell{}).Stage(stage)
	__Cell__000003_Row_1_Cell_4 := (&models.Cell{}).Stage(stage)
	__Cell__000004_Row_1_Cell_5 := (&models.Cell{}).Stage(stage)
	__Cell__000005_Row_2_Cell_1 := (&models.Cell{}).Stage(stage)
	__Cell__000006_Row_2_Cell_2 := (&models.Cell{}).Stage(stage)
	__Cell__000007_Row_2_Cell_3 := (&models.Cell{}).Stage(stage)
	__Cell__000008_Row_2_Cell_4 := (&models.Cell{}).Stage(stage)
	__Cell__000009_Row_2_Cell_5 := (&models.Cell{}).Stage(stage)
	__Cell__000010_T2_R1_A := (&models.Cell{}).Stage(stage)
	__Cell__000011_T2_R1_B := (&models.Cell{}).Stage(stage)
	__Cell__000012_T2_R2_A := (&models.Cell{}).Stage(stage)
	__Cell__000013_T2_R2_B := (&models.Cell{}).Stage(stage)
	__Cell__000014_T2_R3_A := (&models.Cell{}).Stage(stage)
	__Cell__000015_T2_R3_B := (&models.Cell{}).Stage(stage)

	__CellBoolean__000000_Row_1_Cell_1_Cell_False := (&models.CellBoolean{}).Stage(stage)
	__CellBoolean__000001_Row_2_Cell_1_Cell_true := (&models.CellBoolean{}).Stage(stage)

	__CellFloat64__000000_Row_1_Cell2_Float := (&models.CellFloat64{}).Stage(stage)
	__CellFloat64__000001_Row_2_Cell_2_Float := (&models.CellFloat64{}).Stage(stage)

	__CellIcon__000000_Row_1_Cell_3_Delete := (&models.CellIcon{}).Stage(stage)
	__CellIcon__000001_Row_2_Cell_3_home := (&models.CellIcon{}).Stage(stage)

	__CellInt__000000_Row_1_Cell_4_Int := (&models.CellInt{}).Stage(stage)
	__CellInt__000001_Row_2_Cell_4_Int := (&models.CellInt{}).Stage(stage)

	__CellString__000000_Row_1_Cell_5 := (&models.CellString{}).Stage(stage)
	__CellString__000001_Row_2_Cell_5 := (&models.CellString{}).Stage(stage)
	__CellString__000002_T2_R1_A := (&models.CellString{}).Stage(stage)
	__CellString__000003_T2_R1_B := (&models.CellString{}).Stage(stage)
	__CellString__000004_T2_R2_A := (&models.CellString{}).Stage(stage)
	__CellString__000005_T2_R2_B := (&models.CellString{}).Stage(stage)
	__CellString__000006_T2_R3_A := (&models.CellString{}).Stage(stage)
	__CellString__000007_T2_R3_B := (&models.CellString{}).Stage(stage)

	__DisplayedColumn__000000_A := (&models.DisplayedColumn{}).Stage(stage)
	__DisplayedColumn__000001_B := (&models.DisplayedColumn{}).Stage(stage)
	__DisplayedColumn__000002_Column_1_Boolean := (&models.DisplayedColumn{}).Stage(stage)
	__DisplayedColumn__000003_Column_2_Float64 := (&models.DisplayedColumn{}).Stage(stage)
	__DisplayedColumn__000004_Column_3_Icon := (&models.DisplayedColumn{}).Stage(stage)
	__DisplayedColumn__000005_Column_4_Int := (&models.DisplayedColumn{}).Stage(stage)
	__DisplayedColumn__000006_Column_5_String := (&models.DisplayedColumn{}).Stage(stage)

	__Row__000000_Row_1 := (&models.Row{}).Stage(stage)
	__Row__000001_Row_2 := (&models.Row{}).Stage(stage)
	__Row__000002_Table_2_Row_1 := (&models.Row{}).Stage(stage)
	__Row__000003_Table_2_Row_2 := (&models.Row{}).Stage(stage)
	__Row__000004_Table_2_Row_3 := (&models.Row{}).Stage(stage)

	__Table__000000_EmptyTable := (&models.Table{}).Stage(stage)
	__Table__000001_Table_2 := (&models.Table{}).Stage(stage)
	__Table__000002_Table_with_5_types := (&models.Table{}).Stage(stage)

	// Setup of values

	__Cell__000000_Row_1_Cell_1.Name = `Row 1 - Cell 1`

	__Cell__000001_Row_1_Cell_2.Name = `Row 1 - Cell 2`

	__Cell__000002_Row_1_Cell_3.Name = `Row 1 - Cell 3`

	__Cell__000003_Row_1_Cell_4.Name = `Row 1 - Cell 4`

	__Cell__000004_Row_1_Cell_5.Name = `Row 1 - Cell 5`

	__Cell__000005_Row_2_Cell_1.Name = `Row 2 - Cell 1`

	__Cell__000006_Row_2_Cell_2.Name = `Row 2 - Cell 2`

	__Cell__000007_Row_2_Cell_3.Name = `Row 2 - Cell 3`

	__Cell__000008_Row_2_Cell_4.Name = `Row 2 - Cell 4`

	__Cell__000009_Row_2_Cell_5.Name = `Row 2 - Cell 5`

	__Cell__000010_T2_R1_A.Name = `T2 - R1 - A`

	__Cell__000011_T2_R1_B.Name = `T2 - R1 - B`

	__Cell__000012_T2_R2_A.Name = `T2 - R2 - A`

	__Cell__000013_T2_R2_B.Name = `T2 - R2 - B`

	__Cell__000014_T2_R3_A.Name = `T2 - R3 - A`

	__Cell__000015_T2_R3_B.Name = `T2 - R3 - B`

	__CellBoolean__000000_Row_1_Cell_1_Cell_False.Name = `Row 1 -Cell 1 - Cell False`
	__CellBoolean__000000_Row_1_Cell_1_Cell_False.Value = false

	__CellBoolean__000001_Row_2_Cell_1_Cell_true.Name = `Row 2 - Cell 1 - Cell true`
	__CellBoolean__000001_Row_2_Cell_1_Cell_true.Value = true

	__CellFloat64__000000_Row_1_Cell2_Float.Name = `Row 1 - Cell2 - Float`
	__CellFloat64__000000_Row_1_Cell2_Float.Value = 20.433333

	__CellFloat64__000001_Row_2_Cell_2_Float.Name = `Row 2 - Cell 2 - Float`
	__CellFloat64__000001_Row_2_Cell_2_Float.Value = 18.550000

	__CellIcon__000000_Row_1_Cell_3_Delete.Name = `Row 1 - Cell 3 - Delete`
	__CellIcon__000000_Row_1_Cell_3_Delete.Icon = `delete`

	__CellIcon__000001_Row_2_Cell_3_home.Name = `Row 2 - Cell 3 - home`
	__CellIcon__000001_Row_2_Cell_3_home.Icon = `edit`

	__CellInt__000000_Row_1_Cell_4_Int.Name = `Row 1 - Cell 4 - Int`
	__CellInt__000000_Row_1_Cell_4_Int.Value = 10

	__CellInt__000001_Row_2_Cell_4_Int.Name = `Row 2 - Cell 4 - Int`
	__CellInt__000001_Row_2_Cell_4_Int.Value = 288

	__CellString__000000_Row_1_Cell_5.Name = `Row 1 - Cell 5`
	__CellString__000000_Row_1_Cell_5.Value = `Je ferais le métier`

	__CellString__000001_Row_2_Cell_5.Name = `Row 2 - Cell 5`
	__CellString__000001_Row_2_Cell_5.Value = `des idoles antiques`

	__CellString__000002_T2_R1_A.Name = `T2 - R1 - A`
	__CellString__000002_T2_R1_A.Value = `T2 - R1 - A`

	__CellString__000003_T2_R1_B.Name = `T2 - R1 - B`
	__CellString__000003_T2_R1_B.Value = `T2 - R1 - B`

	__CellString__000004_T2_R2_A.Name = `T2 - R2 - A`
	__CellString__000004_T2_R2_A.Value = `T2 - R2 - A`

	__CellString__000005_T2_R2_B.Name = `T2 - R2 - B`
	__CellString__000005_T2_R2_B.Value = `T2 - R2 - B`

	__CellString__000006_T2_R3_A.Name = `T2 - R3 - A`
	__CellString__000006_T2_R3_A.Value = `T2 - R3 - A`

	__CellString__000007_T2_R3_B.Name = `T2 - R3 - B`
	__CellString__000007_T2_R3_B.Value = `T2 - R3 - B`

	__DisplayedColumn__000000_A.Name = `A`

	__DisplayedColumn__000001_B.Name = `B`

	__DisplayedColumn__000002_Column_1_Boolean.Name = `Column 1 - Boolean`

	__DisplayedColumn__000003_Column_2_Float64.Name = `Column 2 - Float64`

	__DisplayedColumn__000004_Column_3_Icon.Name = `Column 3 - Icon`

	__DisplayedColumn__000005_Column_4_Int.Name = `Column 4 - Int`

	__DisplayedColumn__000006_Column_5_String.Name = `Column 5 - String`

	__Row__000000_Row_1.Name = `Row 1`
	__Row__000000_Row_1.IsChecked = true

	__Row__000001_Row_2.Name = `Row 2`
	__Row__000001_Row_2.IsChecked = false

	__Row__000002_Table_2_Row_1.Name = `Table 2 - Row 1`
	__Row__000002_Table_2_Row_1.IsChecked = false

	__Row__000003_Table_2_Row_2.Name = `Table 2 - Row 2`
	__Row__000003_Table_2_Row_2.IsChecked = false

	__Row__000004_Table_2_Row_3.Name = `Table 2 - Row 3`
	__Row__000004_Table_2_Row_3.IsChecked = false

	__Table__000000_EmptyTable.Name = `EmptyTable`
	__Table__000000_EmptyTable.HasFiltering = false
	__Table__000000_EmptyTable.HasColumnSorting = false
	__Table__000000_EmptyTable.HasPaginator = false
	__Table__000000_EmptyTable.HasCheckableRows = false
	__Table__000000_EmptyTable.HasSaveButton = false
	__Table__000000_EmptyTable.CanDragDropRows = false
	__Table__000000_EmptyTable.HasCloseButton = false
	__Table__000000_EmptyTable.SavingInProgress = false
	__Table__000000_EmptyTable.NbOfStickyColumns = 0

	__Table__000001_Table_2.Name = `Table 2`
	__Table__000001_Table_2.HasFiltering = false
	__Table__000001_Table_2.HasColumnSorting = false
	__Table__000001_Table_2.HasPaginator = false
	__Table__000001_Table_2.HasCheckableRows = false
	__Table__000001_Table_2.HasSaveButton = true
	__Table__000001_Table_2.CanDragDropRows = true
	__Table__000001_Table_2.HasCloseButton = true
	__Table__000001_Table_2.SavingInProgress = false
	__Table__000001_Table_2.NbOfStickyColumns = 0

	__Table__000002_Table_with_5_types.Name = `Table with 5 types`
	__Table__000002_Table_with_5_types.HasFiltering = true
	__Table__000002_Table_with_5_types.HasColumnSorting = true
	__Table__000002_Table_with_5_types.HasPaginator = true
	__Table__000002_Table_with_5_types.HasCheckableRows = true
	__Table__000002_Table_with_5_types.HasSaveButton = true
	__Table__000002_Table_with_5_types.CanDragDropRows = false
	__Table__000002_Table_with_5_types.HasCloseButton = false
	__Table__000002_Table_with_5_types.SavingInProgress = false
	__Table__000002_Table_with_5_types.NbOfStickyColumns = 3

	// Setup of pointers
	__Cell__000000_Row_1_Cell_1.CellBool = __CellBoolean__000000_Row_1_Cell_1_Cell_False
	__Cell__000001_Row_1_Cell_2.CellFloat64 = __CellFloat64__000000_Row_1_Cell2_Float
	__Cell__000002_Row_1_Cell_3.CellIcon = __CellIcon__000000_Row_1_Cell_3_Delete
	__Cell__000003_Row_1_Cell_4.CellInt = __CellInt__000000_Row_1_Cell_4_Int
	__Cell__000004_Row_1_Cell_5.CellString = __CellString__000000_Row_1_Cell_5
	__Cell__000005_Row_2_Cell_1.CellBool = __CellBoolean__000001_Row_2_Cell_1_Cell_true
	__Cell__000006_Row_2_Cell_2.CellFloat64 = __CellFloat64__000001_Row_2_Cell_2_Float
	__Cell__000007_Row_2_Cell_3.CellIcon = __CellIcon__000001_Row_2_Cell_3_home
	__Cell__000008_Row_2_Cell_4.CellInt = __CellInt__000001_Row_2_Cell_4_Int
	__Cell__000009_Row_2_Cell_5.CellString = __CellString__000001_Row_2_Cell_5
	__Cell__000010_T2_R1_A.CellString = __CellString__000002_T2_R1_A
	__Cell__000011_T2_R1_B.CellString = __CellString__000003_T2_R1_B
	__Cell__000012_T2_R2_A.CellString = __CellString__000004_T2_R2_A
	__Cell__000013_T2_R2_B.CellString = __CellString__000005_T2_R2_B
	__Cell__000014_T2_R3_A.CellString = __CellString__000006_T2_R3_A
	__Cell__000015_T2_R3_B.CellString = __CellString__000007_T2_R3_B
	__Row__000000_Row_1.Cells = append(__Row__000000_Row_1.Cells, __Cell__000000_Row_1_Cell_1)
	__Row__000000_Row_1.Cells = append(__Row__000000_Row_1.Cells, __Cell__000001_Row_1_Cell_2)
	__Row__000000_Row_1.Cells = append(__Row__000000_Row_1.Cells, __Cell__000002_Row_1_Cell_3)
	__Row__000000_Row_1.Cells = append(__Row__000000_Row_1.Cells, __Cell__000003_Row_1_Cell_4)
	__Row__000000_Row_1.Cells = append(__Row__000000_Row_1.Cells, __Cell__000004_Row_1_Cell_5)
	__Row__000001_Row_2.Cells = append(__Row__000001_Row_2.Cells, __Cell__000005_Row_2_Cell_1)
	__Row__000001_Row_2.Cells = append(__Row__000001_Row_2.Cells, __Cell__000006_Row_2_Cell_2)
	__Row__000001_Row_2.Cells = append(__Row__000001_Row_2.Cells, __Cell__000007_Row_2_Cell_3)
	__Row__000001_Row_2.Cells = append(__Row__000001_Row_2.Cells, __Cell__000008_Row_2_Cell_4)
	__Row__000001_Row_2.Cells = append(__Row__000001_Row_2.Cells, __Cell__000009_Row_2_Cell_5)
	__Row__000002_Table_2_Row_1.Cells = append(__Row__000002_Table_2_Row_1.Cells, __Cell__000011_T2_R1_B)
	__Row__000002_Table_2_Row_1.Cells = append(__Row__000002_Table_2_Row_1.Cells, __Cell__000010_T2_R1_A)
	__Row__000003_Table_2_Row_2.Cells = append(__Row__000003_Table_2_Row_2.Cells, __Cell__000013_T2_R2_B)
	__Row__000003_Table_2_Row_2.Cells = append(__Row__000003_Table_2_Row_2.Cells, __Cell__000012_T2_R2_A)
	__Row__000004_Table_2_Row_3.Cells = append(__Row__000004_Table_2_Row_3.Cells, __Cell__000014_T2_R3_A)
	__Row__000004_Table_2_Row_3.Cells = append(__Row__000004_Table_2_Row_3.Cells, __Cell__000015_T2_R3_B)
	__Table__000001_Table_2.DisplayedColumns = append(__Table__000001_Table_2.DisplayedColumns, __DisplayedColumn__000000_A)
	__Table__000001_Table_2.Rows = append(__Table__000001_Table_2.Rows, __Row__000003_Table_2_Row_2)
	__Table__000001_Table_2.Rows = append(__Table__000001_Table_2.Rows, __Row__000002_Table_2_Row_1)
	__Table__000001_Table_2.Rows = append(__Table__000001_Table_2.Rows, __Row__000004_Table_2_Row_3)
	__Table__000002_Table_with_5_types.DisplayedColumns = append(__Table__000002_Table_with_5_types.DisplayedColumns, __DisplayedColumn__000002_Column_1_Boolean)
	__Table__000002_Table_with_5_types.DisplayedColumns = append(__Table__000002_Table_with_5_types.DisplayedColumns, __DisplayedColumn__000003_Column_2_Float64)
	__Table__000002_Table_with_5_types.DisplayedColumns = append(__Table__000002_Table_with_5_types.DisplayedColumns, __DisplayedColumn__000004_Column_3_Icon)
	__Table__000002_Table_with_5_types.DisplayedColumns = append(__Table__000002_Table_with_5_types.DisplayedColumns, __DisplayedColumn__000005_Column_4_Int)
	__Table__000002_Table_with_5_types.DisplayedColumns = append(__Table__000002_Table_with_5_types.DisplayedColumns, __DisplayedColumn__000006_Column_5_String)
	__Table__000002_Table_with_5_types.Rows = append(__Table__000002_Table_with_5_types.Rows, __Row__000000_Row_1)
	__Table__000002_Table_with_5_types.Rows = append(__Table__000002_Table_with_5_types.Rows, __Row__000001_Row_2)
}
