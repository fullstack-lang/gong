package models

type DiagramIF interface {
	GetName() string
	GetIsChecked() bool
	GetIsShowPrefix() bool
	GetElementWhoseDiagramListIsDisplayed() AbstractType
	SetElementWhoseDiagramListIsDisplayed(AbstractType)
	GetDefaultBoxWidth() float64
	GetDefaultBoxHeigth() float64
	IsEditable() bool
}
