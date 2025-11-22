// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/sim/go/models"
)

func NewCellDeleteIconImplPointerToGongstruct[T models.PointerToGongstruct](
	Instance T,
	probe *Probe,
) (cellDeleteIconImpl *CellDeleteIconImplPointerToGongstruct[T]) {
	cellDeleteIconImpl = new(CellDeleteIconImplPointerToGongstruct[T])
	cellDeleteIconImpl.Instance = Instance
	cellDeleteIconImpl.probe = probe
	return
}

type CellDeleteIconImplPointerToGongstruct[T models.PointerToGongstruct] struct {
	Instance T
	probe    *Probe
}

func (cellDeleteIconImpl *CellDeleteIconImplPointerToGongstruct[T]) CellIconUpdated(stage *gongtable.Stage,
	row, updatedCellIcon *gongtable.CellIcon) {
	// log.Println("CellIconUpdate: CellIconUpdated", updatedCellIcon.Name)

	switch instancesTyped := any(cellDeleteIconImpl.Instance).(type) {
	// insertion point
	case *models.Command:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.DummyAgent:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Engine:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Event:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Status:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.UpdateState:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.probe.stageOfInterest.Commit()

	updateAndCommitTable[T](cellDeleteIconImpl.probe)
	updateAndCommitTree(cellDeleteIconImpl.probe)
	cellDeleteIconImpl.probe.tableStage.Commit()
}
