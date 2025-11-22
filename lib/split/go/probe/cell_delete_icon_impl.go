// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/lib/split/go/models"
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
	case *models.AsSplit:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.AsSplitArea:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Button:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Cursor:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.FavIcon:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Form:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Load:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.LogoOnTheLeft:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.LogoOnTheRight:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Markdown:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Slider:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Split:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Svg:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Table:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Title:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Tone:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Tree:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.View:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.Xlsx:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.probe.stageOfInterest.Commit()

	updateAndCommitTable[T](cellDeleteIconImpl.probe)
	updateAndCommitTree(cellDeleteIconImpl.probe)
	cellDeleteIconImpl.probe.tableStage.Commit()
}
