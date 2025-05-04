package probe

const CellDeleteIconImplTemplate = `// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gong/lib/table/go/models"

	"{{PkgPathRoot}}/models"
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
	// insertion point{{` + string(rune(CellDeleteIconImplSwitchCase)) + `}}
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.probe.stageOfInterest.Commit()

	updateAndCommitTable[T](cellDeleteIconImpl.probe)
	updateAndCommitTree(cellDeleteIconImpl.probe)
	cellDeleteIconImpl.probe.tableStage.Commit()
}
`

type CellDeleteIconImplInsertionId int

const (
	CellDeleteIconImplSwitchCase CellDeleteIconImplInsertionId = iota
	CellDeleteIconImplInsertionNb
)

var CellDeleteIconImplSubTemplateCode map[string]string = // new line
map[string]string{
	string(rune(CellDeleteIconImplSwitchCase)): `
	case *models.{{Structname}}:
		instancesTyped.Unstage(cellDeleteIconImpl.probe.stageOfInterest)`,
}
