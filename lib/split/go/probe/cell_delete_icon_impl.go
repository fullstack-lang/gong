// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
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
	Instance   *T
	probe *Probe
}

func (cellDeleteIconImpl *CellDeleteIconImpl[T]) CellIconUpdated(stage *gongtable.StageStruct,
	row, updatedCellIcon *gongtable.CellIcon) {
	log.Println("CellIconUpdate: CellIconUpdated", updatedCellIcon.Name)

	switch instancesTyped := any(cellDeleteIconImpl.Instance).(type) {
	// insertion point
	case *models.AsSplit:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.AsSplitArea:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Doc:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Form:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Split:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Svg:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Table:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Tree:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.View:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.probe.stageOfInterest.Commit()

	fillUpTable[T](cellDeleteIconImpl.probe)
	fillUpTree(cellDeleteIconImpl.probe)
	cellDeleteIconImpl.probe.tableStage.Commit()
}

