package models

// Cell is the sum type for cell
type Cell struct {
	Name string

	CellString  *CellString
	CellFloat64 *CellFloat64
	CellInt     *CellInt
	CellBool    *CellBoolean
	CellIcon    *CellIcon
}
