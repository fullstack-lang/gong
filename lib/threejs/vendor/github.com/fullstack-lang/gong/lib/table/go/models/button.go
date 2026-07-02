package models

type Button struct {
	Name string // used for indentfication of click
	Icon string // material angular icon

	// SVG Icon, overides the angular materal icon
	SVGIcon *SVGIcon

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

func (button *Button) OnAfterUpdate(_ *Stage, _, _ *Button) {
	if button.OnClick != nil {
		button.OnClick()
	}
}
