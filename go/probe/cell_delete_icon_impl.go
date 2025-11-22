// generated code - do not edit
package probe

import (
	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/go/models"
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
	case *models.GongBasicField:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.GongEnum:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.GongEnumValue:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.GongLink:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.GongNote:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.GongStruct:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.GongTimeField:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.MetaReference:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.ModelPkg:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.PointerToGongStructField:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	case *models.SliceOfPointerToGongStructField:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.probe.stageOfInterest.Commit()

	updateAndCommitTable[T](cellDeleteIconImpl.probe)
	updateAndCommitTree(cellDeleteIconImpl.probe)
	cellDeleteIconImpl.probe.tableStage.Commit()
}
