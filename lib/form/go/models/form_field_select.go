package models

type FormFieldSelect struct {
	Name  string
	Value *Option

	Options []*Option

	CanBeEmpty bool

	// PreserveInitialOrder overides default alphabetical
	// ordering of Options in the pulldown menu
	PreserveInitialOrder bool

	IsLiveUpdate bool

	// swagger:ignore
	Proxy FormElementProxyInterface
}

func (formFieldSelect *FormFieldSelect) OnAfterUpdate(
	stage *Stage,
	stagedInstance, frontInstance *FormFieldSelect) {

	stagedInstance.Value = frontInstance.Value
	if stagedInstance.Proxy != nil {
		stagedInstance.Proxy.Updated()
	}
}
