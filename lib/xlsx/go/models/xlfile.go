package models

import (
	"github.com/tealeg/xlsx/v3"
)

type XLFile struct {
	Name string

	file *xlsx.File

	NbSheets int
	Sheets   []*XLSheet
}

func (xlfile *XLFile) Open(stage *Stage, path string) {

	xlfile.Name = path
	var err error
	xlfile.file, err = xlsx.OpenFile(path)

	if err != nil {
		panic(err)
	}
	// wb now contains a reference to the workbook
	// show all the sheets in the workbook
	// fmt.Println("Sheets in this file:")
	for _, sh := range xlfile.file.Sheets {
		// fmt.Println(i, sh.Name)

		xlsheet := new(XLSheet).Stage(stage)
		xlsheet.Name = sh.Name
		xlsheet.sheet = sh
		xlsheet.MaxCol = sh.MaxCol
		xlsheet.MaxRow = sh.MaxRow

		xlfile.Sheets = append(xlfile.Sheets, xlsheet)

		emptyRow := false
		rowIndex := 0
		for !emptyRow {

			cell, err := sh.Cell(rowIndex, 0)
			if err != nil {
				continue
			}
			if cell.Value == "" {
				emptyRow = true
				continue
			}
			row, _ := sh.Row(rowIndex)
			xlrow := new(XLRow).Stage(stage)
			xlrow.row = row
			//			xlrow.Name = xlsheet.Name + "-" + fmt.Sprintf("%4d", rowIndex)
			xlrow.Name = cell.Value
			xlsheet.Rows = append(xlsheet.Rows, xlrow)

			xlsheet.NbRows = rowIndex + 1

			// get cells
			for colIndex := 0; colIndex < xlsheet.MaxCol; colIndex = colIndex + 1 {
				xlcell := new(XLCell).Stage(stage)

				xlcell.cell, err = sh.Cell(rowIndex, colIndex)
				if err != nil {
					continue
				}

				xlcell.Name = xlcell.cell.Value
				xlcell.X = colIndex
				xlcell.Y = rowIndex
				xlrow.Cells = append(xlrow.Cells, xlcell)
				xlsheet.SheetCells = append(xlsheet.SheetCells, xlcell)
			}
			rowIndex = rowIndex + 1
		}
		// fmt.Println("Sheet ", xlsheet.Name, "Nb Rows", xlsheet.NbRows)
	}
	// fmt.Println("----")

	xlfile.NbSheets = len(xlfile.file.Sheets)
}

var nbRows int

func rowVisitor(r *xlsx.Row) error {
	nbRows = nbRows + 1
	return nil
}
