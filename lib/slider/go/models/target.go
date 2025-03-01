package models

type Target interface {
	OnAfterUpdateSliderElement()
	GetSliderStage() *StageStruct
}
