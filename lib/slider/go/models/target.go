package models

type Target interface {
	OnAfterUpdateSliderElement()
	GetSliderStage() *Stage
}
