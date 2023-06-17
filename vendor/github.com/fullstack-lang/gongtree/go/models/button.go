package models

type Button struct {
	Name string // used for indentfication of click
	Icon string

	Impl ButtonImplInterface
}

type ButtonImplInterface interface {

	// ButtonUpdated function is called each time a Button is modified
	ButtonUpdated(tage *StageStruct, button, updatedButton *Button)
}

func (button *Button) OnAfterUpdate(stage *StageStruct, _, frontButton *Button) {

	if button.Impl != nil {
		button.Impl.ButtonUpdated(stage, button, frontButton)
	}
}
