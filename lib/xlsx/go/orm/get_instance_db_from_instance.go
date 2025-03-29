// generated code - do not edit
package orm

import (
	"github.com/fullstack-lang/gong/lib/xlsx/go/models"
)

type GongstructDB interface {
}

func GetInstanceDBFromInstance[T models.Gongstruct, T2 GongstructDB](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (ret *T2) {

	switch concreteInstance := any(instance).(type) {
	// insertion point for per struct backup
	case *models.DisplaySelection:
		displayselectionInstance := any(concreteInstance).(*models.DisplaySelection)
		ret2 := backRepo.BackRepoDisplaySelection.GetDisplaySelectionDBFromDisplaySelectionPtr(displayselectionInstance)
		ret = any(ret2).(*T2)
	case *models.XLCell:
		xlcellInstance := any(concreteInstance).(*models.XLCell)
		ret2 := backRepo.BackRepoXLCell.GetXLCellDBFromXLCellPtr(xlcellInstance)
		ret = any(ret2).(*T2)
	case *models.XLFile:
		xlfileInstance := any(concreteInstance).(*models.XLFile)
		ret2 := backRepo.BackRepoXLFile.GetXLFileDBFromXLFilePtr(xlfileInstance)
		ret = any(ret2).(*T2)
	case *models.XLRow:
		xlrowInstance := any(concreteInstance).(*models.XLRow)
		ret2 := backRepo.BackRepoXLRow.GetXLRowDBFromXLRowPtr(xlrowInstance)
		ret = any(ret2).(*T2)
	case *models.XLSheet:
		xlsheetInstance := any(concreteInstance).(*models.XLSheet)
		ret2 := backRepo.BackRepoXLSheet.GetXLSheetDBFromXLSheetPtr(xlsheetInstance)
		ret = any(ret2).(*T2)
	default:
		_ = concreteInstance
	}
	return
}

func GetID[T models.Gongstruct](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance *T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.DisplaySelection:
		tmp := GetInstanceDBFromInstance[models.DisplaySelection, DisplaySelectionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.XLCell:
		tmp := GetInstanceDBFromInstance[models.XLCell, XLCellDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.XLFile:
		tmp := GetInstanceDBFromInstance[models.XLFile, XLFileDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.XLRow:
		tmp := GetInstanceDBFromInstance[models.XLRow, XLRowDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.XLSheet:
		tmp := GetInstanceDBFromInstance[models.XLSheet, XLSheetDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}

func GetIDPointer[T models.PointerToGongstruct](
	stage *models.Stage,
	backRepo *BackRepoStruct,
	instance T) (id int) {

	switch inst := any(instance).(type) {
	// insertion point for per struct backup
	case *models.DisplaySelection:
		tmp := GetInstanceDBFromInstance[models.DisplaySelection, DisplaySelectionDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.XLCell:
		tmp := GetInstanceDBFromInstance[models.XLCell, XLCellDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.XLFile:
		tmp := GetInstanceDBFromInstance[models.XLFile, XLFileDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.XLRow:
		tmp := GetInstanceDBFromInstance[models.XLRow, XLRowDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	case *models.XLSheet:
		tmp := GetInstanceDBFromInstance[models.XLSheet, XLSheetDB](
			stage, backRepo, inst,
		)
		id = int(tmp.ID)
	default:
		_ = inst
	}
	return
}
