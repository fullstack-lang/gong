package models

type Target interface {
	OnAfterUpdateButton()
	GetButtonsStage() *Stage
}
