package models

// Generic button creation function
func NewButton(
	target Target,
	name string,
	icon string,

) *Button {
	button := new(Button).Stage(target.GetButtonsStage())
	button.Name = name
	button.Icon = icon

	proxy := NewButtonProxy(
		button,
		target,
	)

	button.Impl = proxy

	return button
}

// NewButtonProxy creates a new proxy for a button
func NewButtonProxy(
	button *Button,
	target Target,
) *ButtonProxy {
	proxy := new(ButtonProxy)
	proxy.button = button
	proxy.target = target
	return proxy
}

// ButtonProxy is a generic proxy for both int and float64
type ButtonProxy struct {
	button *Button
	Value  *bool
	target Target
}

// Updated handles updating values when the button changes
func (proxy *ButtonProxy) ButtonUpdated(*Stage, *Button, *Button) {

	proxy.target.OnAfterUpdateButton()
}

type Target interface {
	OnAfterUpdateButton()
	GetButtonsStage() *Stage
}
