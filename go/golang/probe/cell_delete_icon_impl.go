package probe

const CellDeleteIconImplTemplate = `// generated code - do not edit
package probe

import (
	"log"

	gongtable "github.com/fullstack-lang/gongtable/go/models"

	"{{PkgPathRoot}}/models"
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
	// insertion point{{` + string(rune(CellDeleteIconImplSwitchCase)) + `}}
	default:
		_ = instancesTyped
	}
	cellDeleteIconImpl.playground.stageOfInterest.Commit()

	fillUpTable[T](cellDeleteIconImpl.playground)
	fillUpTree(cellDeleteIconImpl.playground)
	cellDeleteIconImpl.playground.tableStage.Commit()
}

`

type CellDeleteIconImplInsertionId int

const (
	CellDeleteIconImplSwitchCase CellDeleteIconImplInsertionId = iota
	CellDeleteIconImplInsertionNb
)

var CellDeleteIconImplSubTemplateCode map[string]string = // new line
map[string]string{
	string(rune(ButtonImplPerGongstructCallToForm)): `
	case *models.{{Structname}}:
		instancesTyped.Unstage(cellDeleteIconImpl.playground.stageOfInterest)`,
}
