package models

type Diagram struct {
	Name           string
	IsChecked      bool
	IsExpanded     bool
	IsEditable_    bool
	IsInRenameMode bool

	Product_Shapes []*ProductShape
}

func (d *Diagram) IsEditable() bool {
	return d.IsEditable_
}
