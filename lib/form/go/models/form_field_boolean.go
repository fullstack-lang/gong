package models

type CheckBox struct {
	Name  string
	Value bool

	LabelForTrue  string
	LabelForFalse string

	IsDisabled   bool
	IsLiveUpdate bool

	// swagger:ignore
	Proxy FormElementProxyInterface
}

func (checkBox *CheckBox) OnAfterUpdate(
	stage *Stage,
	stagedInstance, frontInstance *CheckBox) {

	stagedInstance.Value = frontInstance.Value
	if stagedInstance.Proxy != nil {
		stagedInstance.Proxy.Updated()
	}
}
