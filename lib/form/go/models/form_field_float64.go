package models

type FormFieldFloat64 struct {
	Name  string
	Value float64

	HasMinValidator bool
	MinValue        float64

	HasMaxValidator bool
	MaxValue        float64

	IsSlider     bool
	Step         float64
	IsLiveUpdate bool

	// swagger:ignore
	Proxy FormElementProxyInterface
}

func (formFieldFloat64 *FormFieldFloat64) OnAfterUpdate(
	stage *Stage,
	stagedInstance, frontInstance *FormFieldFloat64) {

	stagedInstance.Value = frontInstance.Value
	if stagedInstance.Proxy != nil {
		stagedInstance.Proxy.Updated()
	}
}
