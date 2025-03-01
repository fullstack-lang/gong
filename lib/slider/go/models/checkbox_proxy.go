package models

// Generic checkbox creation function
func NewCheckbox(
	target Target,
	name string,
	labelForTrue string,
	labelForFalse string,
	valueRef *bool,
) *Checkbox {
	checkbox := new(Checkbox).Stage(target.GetSliderStage())
	checkbox.Name = name
	checkbox.LabelForTrue = labelForTrue
	checkbox.LabelForFalse = labelForFalse
	checkbox.ValueBool = *valueRef

	proxy := NewCheckboxProxy(
		checkbox,
		valueRef,
		target,
	)

	checkbox.Proxy = proxy

	return checkbox
}

// NewCheckboxProxy creates a new proxy for a checkbox
func NewCheckboxProxy(
	checkbox *Checkbox,
	value *bool,
	target Target,
) *CheckboxProxy {
	proxy := new(CheckboxProxy)
	proxy.checkbox = checkbox
	proxy.Value = value
	proxy.target = target
	return proxy
}

// CheckboxProxy is a generic proxy for both int and float64
type CheckboxProxy struct {
	checkbox *Checkbox
	Value    *bool
	target   Target
}

// Updated handles updating values when the checkbox changes
func (proxy *CheckboxProxy) Updated() {
	*proxy.Value = proxy.checkbox.ValueBool

	proxy.target.OnAfterUpdateSliderElement()
}
