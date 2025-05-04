// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/table/go/models"
)

func NewCellDeleteIconImpl[T models.Gongstruct](
	Instance *T,
	probe *Probe,
) (cellDeleteIconImpl *CellDeleteIconImpl[T]) {
	cellDeleteIconImpl = new(CellDeleteIconImpl[T])
	cellDeleteIconImpl.Instance = Instance
	cellDeleteIconImpl.probe = probe
	return
}

type CellDeleteIconImpl[T models.Gongstruct] struct {
	Instance *T
	probe    *Probe
}

func (cellDeleteIconImpl *CellDeleteIconImpl[T]) CellIconUpdated(stage *gongtable.Stage,
	row, updatedCellIcon *gongtable.CellIcon) {
	log.Println("CellIconUpdate: CellIconUpdated", updatedCellIcon.Name)

	switch instancesTyped := any(cellDeleteIconImpl.Instance).(type) {
	// insertion point
	case *models.Cell:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.CellBoolean:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.CellFloat64:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.CellIcon:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.CellInt:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.CellString:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.CheckBox:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DisplayedColumn:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.FormDiv:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.FormEditAssocButton:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.FormField:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.FormFieldDate:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.FormFieldDateTime:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.FormFieldFloat64:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.FormFieldInt:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.FormFieldSelect:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.FormFieldString:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.FormFieldTime:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.FormGroup:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.FormSortAssocButton:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Option:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Row:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Table:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.probe.stageOfInterest.Commit()

	updateAndCommitTable[T](cellDeleteIconImpl.probe)
	updateAndCommitTree(cellDeleteIconImpl.probe)
	cellDeleteIconImpl.probe.tableStage.Commit()
}
