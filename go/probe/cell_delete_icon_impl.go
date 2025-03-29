// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/go/models"
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

func (cellDeleteIconImpl *CellDeleteIconImpl[T]) CellIconUpdated(stage *gongtable.Stage,
	row, updatedCellIcon *gongtable.CellIcon) {
	log.Println("CellIconUpdate: CellIconUpdated", updatedCellIcon.Name)

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
	case *models.Meta:
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

	fillUpTable[T](cellDeleteIconImpl.probe)
	fillUpTree(cellDeleteIconImpl.probe)
	cellDeleteIconImpl.probe.tableStage.Commit()
}

