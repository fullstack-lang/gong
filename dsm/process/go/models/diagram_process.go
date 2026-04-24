package models

type DiagramProcess struct {
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

	// DSM specific fields
	Process_Shapes              []*ProcessShape
	map_Process_ProcessShape    map[*Process]*ProcessShape
	ProcesssWhoseNodeIsExpanded []*Process
	IsProcesssNodeExpanded      bool

	ProcessComposition_Shapes           []*ProcessCompositionShape
	map_Process_ProcessCompositionShape map[*Process]*ProcessCompositionShape
}

func (d *DiagramProcess) IsEditable() bool {
	return d.IsEditable_
}

func (d *DiagramProcess) SetEditable(v bool) {
	d.IsEditable_ = v
}

func (d *DiagramProcess) GetDefaultBoxHeigth() float64 {
	return d.DefaultBoxHeigth
}

func (d *DiagramProcess) GetDefaultBoxWidth() float64 {
	return d.DefaultBoxWidth
}

func (d *DiagramProcess) GetElementWhoseDiagramListIsDisplayed() AbstractType {
	return d.elementWhoseDiagramListIsDisplayed
}

func (d *DiagramProcess) SetElementWhoseDiagramListIsDisplayed(v AbstractType) {
	d.elementWhoseDiagramListIsDisplayed = v
}

func (d *DiagramProcess) GetIsChecked() bool {
	return d.IsChecked
}

func (d *DiagramProcess) SetIsChecked(v bool) {
	d.IsChecked = v
}

func (d *DiagramProcess) GetIsShowPrefix() bool {
	return d.IsShowPrefix
}

func (d *DiagramProcess) SetIsShowPrefix(v bool) {
	d.IsShowPrefix = v
}
