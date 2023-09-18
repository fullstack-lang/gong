// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"github.com/fullstack-lang/gong/go/models"
)

func NewCellDeleteIconImpl[T models.Gongstruct](
	Instance *T,
	playground *Playground,
) (cellDeleteIconImpl *CellDeleteIconImpl[T]) {
	cellDeleteIconImpl = new(CellDeleteIconImpl[T])
	cellDeleteIconImpl.Instance = Instance
	cellDeleteIconImpl.playground = playground
	return
}

type CellDeleteIconImpl[T models.Gongstruct] struct {
	Instance   *T
	playground *Playground
}

func (cellDeleteIconImpl *CellDeleteIconImpl[T]) CellIconUpdated(stage *gongtable.StageStruct,
	row, updatedCellIcon *gongtable.CellIcon) {
	log.Println("CellIconUpdate: CellIconUpdated", updatedCellIcon.Name)

	switch instancesTyped := any(cellDeleteIconImpl.Instance).(type) {
	// insertion point
	case *models.GongBasicField:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.GongEnum:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.GongEnumValue:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.GongLink:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.GongNote:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.GongStruct:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.GongTimeField:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.Meta:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.MetaReference:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.ModelPkg:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.PointerToGongStructField:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	case *models.SliceOfPointerToGongStructField:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.playground.stageOfInterest.Commit()

	fillUpTable[T](cellDeleteIconImpl.playground)
	fillUpTree(cellDeleteIconImpl.playground)
	cellDeleteIconImpl.playground.tableStage.Commit()
}

