// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/table/go/models"
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
	cellDeleteIconImpl.Instance.UnstageVoid(cellDeleteIconImpl.probe.stageOfInterest)

	cellDeleteIconImpl.probe.stageOfInterest.Commit()

	updateProbeTable[T](cellDeleteIconImpl.probe)
	updateAndCommitTree(cellDeleteIconImpl.probe)
	cellDeleteIconImpl.probe.tableStage.Commit()
}
