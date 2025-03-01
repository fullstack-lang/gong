package models

type Layout struct {
	Name string

	Groups []*Group
}

type Group struct {
	Name       string
	Percentage float64
	Buttons    []*Button
}

// for instance is minor / is major
type Button struct {
	Name string

	Label string

	Icon string

	Proxy ButtonProxyInterface
}

type ButtonProxyInterface interface {
	Updated()
}

func (button *Button) OnAfterUpdate(
	stage *StageStruct,
	stageButton, frontButton *Button) {

	if button.Proxy != nil {
		button.Proxy.Updated()
	}
}
