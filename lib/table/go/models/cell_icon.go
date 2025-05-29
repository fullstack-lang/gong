package models

type CellIcon struct {
	Name string

	// reference of the material icon (ex "home", "delete", "edit")
	Icon string

	// needs confirmation ? (for instance if the icon is a delete icon)
	NeedsConfirmation   bool
	ConfirmationMessage string

	// swagger:ignore
	Impl CellIconImplInterface
}

type CellIconImplInterface interface {

	// CellIconUpdated function is called each time a CellIcon is modified
	CellIconUpdated(stage *Stage, cellIcon, updatedCellIcon *CellIcon)
}

func (cellIcon *CellIcon) OnAfterUpdate(stage *Stage, stagedInstance, frontCellIcon *CellIcon) {

	if cellIcon.Impl != nil {
		cellIcon.Impl.CellIconUpdated(stage, cellIcon, frontCellIcon)
	}
}
