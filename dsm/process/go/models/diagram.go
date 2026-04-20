package models

type Diagram struct {
	Name string

	LibraryAbstractFields
	AbstractTypeFields

	IsChecked   bool
	IsEditable_ bool

	IsShowPrefix bool // display shapes with their prefix

	DefaultBoxWidth  float64
	DefaultBoxHeigth float64

	Width  float64
	Height float64

	elementWhoseDiagramListIsDisplayed AbstractType
}

func (d *Diagram) IsEditable() bool {
	return d.IsEditable_
}

func (d *Diagram) SetEditable(v bool) {
	d.IsEditable_ = v
}

func (d *Diagram) GetDefaultBoxHeigth() float64 {
	return d.DefaultBoxHeigth
}

func (d *Diagram) GetDefaultBoxWidth() float64 {
	return d.DefaultBoxWidth
}

func (d *Diagram) GetElementWhoseDiagramListIsDisplayed() AbstractType {
	return d.elementWhoseDiagramListIsDisplayed
}

func (d *Diagram) SetElementWhoseDiagramListIsDisplayed(v AbstractType) {
	d.elementWhoseDiagramListIsDisplayed = v
}

func (d *Diagram) GetIsChecked() bool {
	return d.IsChecked
}

func (d *Diagram) SetIsChecked(v bool) {
	d.IsChecked = v
}

func (d *Diagram) GetIsShowPrefix() bool {
	return d.IsShowPrefix
}

func (d *Diagram) SetIsShowPrefix(v bool) {
	d.IsShowPrefix = v
}
