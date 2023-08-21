package models

type Row struct {
	Name  string
	Cells []*Cell

	IsChecked bool
}
