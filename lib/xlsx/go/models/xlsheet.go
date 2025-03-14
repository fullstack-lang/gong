package models

import (
	"github.com/tealeg/xlsx/v3"
)

type XLSheet struct {
	Name   string
	MaxRow int
	MaxCol int

	sheet  *xlsx.Sheet
	NbRows int

	Rows       []*XLRow
	SheetCells []*XLCell

	arrayOfString2D *[][]string
}

func (xLSheet *XLSheet) GetArrayOfString2S() (arrayOfString2D *[][]string) {
	return xLSheet.arrayOfString2D
}

func (xLSheet *XLSheet) FillUpArrayOfString2S() {

	tmp := make([][]string, xLSheet.NbRows)
	xLSheet.arrayOfString2D = &tmp

	for rowId := 0; rowId < xLSheet.NbRows; rowId = rowId + 1 {
		(*xLSheet.arrayOfString2D)[rowId] = make([]string, xLSheet.MaxCol)

		xlRow := xLSheet.Rows[rowId]
		for columnId := 0; columnId < xLSheet.MaxCol; columnId = columnId + 1 {
			xLCell := (*xlRow).Cells[columnId]
			(*xLSheet.arrayOfString2D)[rowId][columnId] = xLCell.Name
		}
	}

}
