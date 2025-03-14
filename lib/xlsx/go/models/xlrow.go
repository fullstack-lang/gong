package models

import (
	"github.com/tealeg/xlsx/v3"
)

type XLRow struct {
	Name     string
	RowIndex int

	row *xlsx.Row

	Cells []*XLCell
}
