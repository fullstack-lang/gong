package models

type Button struct {
	Name string // used for indentfication of click
	Icon string // material angular icon

	// SVG Icon, overides the angular materal icon
	SVGIcon *SVGIcon

	// Deprecated, uses OnClick instead
	Impl ButtonImplInterface
	// Deprecated, uses OnClick instead
	OnUpdate func(stage *Stage, updatedButton *Button)

	OnClick func() // is called when the button is clicked

	IsDisabled bool

	HasToolTip      bool
	ToolTipText     string
	ToolTipPosition ToolTipPositionEnum
}

type ToolTipPositionEnum string

const (
	Below ToolTipPositionEnum = "below"
	Above ToolTipPositionEnum = "above"
	Left  ToolTipPositionEnum = "left"
	Right ToolTipPositionEnum = "right"
)

type ButtonImplInterface interface {
	// ButtonUpdated function is called each time a Button is modified
	ButtonUpdated(tage *Stage, button, updatedButton *Button)
}

func (button *Button) OnAfterUpdate(stage *Stage, _, frontButton *Button) {
	if button.Impl != nil {
		button.Impl.ButtonUpdated(stage, button, frontButton)
	}
	if button.OnUpdate != nil {
		button.OnUpdate(stage, frontButton)
	}
}
