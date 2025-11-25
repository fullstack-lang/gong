package models

type Diagram struct {
	Name           string
	IsChecked      bool
	IsExpanded     bool
	IsEditable_    bool
	IsInRenameMode bool

	State_Shapes      []*StateShape
	Transition_Shapes []*Transition_Shape
}

func (d *Diagram) IsEditable() bool {
	return d.IsEditable_
}
