package models

type FormFieldInt struct {
	Name  string
	Value int

	HasMinValidator bool
	MinValue        int

	HasMaxValidator bool
	MaxValue        int

	IsSlider     bool
	Step         int
	IsLiveUpdate bool

	// swagger:ignore
	Proxy FormElementProxyInterface
}

func (formFieldInt *FormFieldInt) OnAfterUpdate(
	stage *Stage,
	stagedInstance, frontInstance *FormFieldInt) {

	stagedInstance.Value = frontInstance.Value
	if stagedInstance.Proxy != nil {
		stagedInstance.Proxy.Updated()
	}
}
