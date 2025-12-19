package models

type TableRow struct {
	Name         string
	Content      string
	Node         *Node
	TableColumns []*TableColumn
}
