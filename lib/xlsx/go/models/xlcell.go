package models

import (
	"github.com/tealeg/xlsx/v3"
)

type XLCell struct {
	Name string

	X, Y int

	cell *xlsx.Cell
}
