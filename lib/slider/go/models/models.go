package models

type Layout struct {
	Name string

	Groups []*Group
}

type Group struct {
	Name       string
	Percentage float64
	Sliders    []*Slider
	Checkboxes []*Checkbox
}

type Slider struct {
	Name string

	IsFloat64 bool
	IsInt     bool

	MinInt   int
	MaxInt   int
	StepInt  int
	ValueInt int

	MinFloat64   float64
	MaxFloat64   float64
	StepFloat64  float64
	ValueFloat64 float64

	Proxy SliderProxyInterface
}

type SliderProxyInterface interface {
	Updated()
}

func (slider *Slider) OnAfterUpdate(
	stage *Stage,
	stageSlider, frontSlider *Slider) {

	slider.ValueFloat64 = frontSlider.ValueFloat64
	slider.ValueInt = frontSlider.ValueInt

	if slider.Proxy != nil {
		slider.Proxy.Updated()
	}
}

// for instance is minor / is major
type Checkbox struct {
	Name string

	ValueBool bool

	LabelForTrue  string
	LabelForFalse string

	Proxy CheckboxProxyInterface
}

type CheckboxProxyInterface interface {
	Updated()
}

func (checkbox *Checkbox) OnAfterUpdate(
	stage *Stage,
	stageCheckbox, frontCheckbox *Checkbox) {

	checkbox.ValueBool = frontCheckbox.ValueBool

	if checkbox.Proxy != nil {
		checkbox.Proxy.Updated()
	}
}
