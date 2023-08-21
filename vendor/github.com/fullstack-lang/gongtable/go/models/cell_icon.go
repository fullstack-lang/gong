package models

type CellIcon struct {
	Name string

	// reference of the material icon (ex "home", "delete", "edit")
	Icon string
}

func (cellIcon *CellIcon) OnAfterUpdate(stage *StageStruct, stagedInstance, frontInstance *CellIcon) {

	//
}
