// generated code (do not edit)
package models

type DiagramIF interface {
	GetName() string
	GetIsChecked() bool
	SetIsChecked(bool)
	SetIsShowPrefix(bool)
	GetIsShowPrefix() bool
	GetElementWhoseDiagramListIsDisplayed() AbstractType
	SetElementWhoseDiagramListIsDisplayed(AbstractType)
	GetDefaultBoxWidth() float64
	GetDefaultBoxHeigth() float64
	IsEditable() bool
}
