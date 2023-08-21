package data

import (
	"log"

	gong_models "github.com/fullstack-lang/gong/go/models"
	gongtable_models "github.com/fullstack-lang/gongtable/go/models"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
)

type NodeImplGongstruct struct {
	gongStruct      *gong_models.GongStruct
	gongtableStage  *gongtable_models.StageStruct
	stageOfInterest *models.StageStruct
}

func NewNodeImplGongstruct(
	gongStruct *gong_models.GongStruct,
	gongtableStage *gongtable_models.StageStruct,
	stageOfInterest *models.StageStruct,
) (nodeImplGongstruct *NodeImplGongstruct) {

	nodeImplGongstruct = new(NodeImplGongstruct)
	nodeImplGongstruct.gongStruct = gongStruct
	nodeImplGongstruct.gongtableStage = gongtableStage
	nodeImplGongstruct.stageOfInterest = stageOfInterest
	return
}

func (nodeImplGongstruct *NodeImplGongstruct) OnAfterUpdate(
	gongtreeStage *gongtree_models.StageStruct,
	stagedNode, frontNode *gongtree_models.Node) {

	// setting the value of the staged node	to the new value
	// otherwise, the expansion memory is lost
	if stagedNode.IsExpanded != frontNode.IsExpanded {
		stagedNode.IsExpanded = frontNode.IsExpanded
		return
	}

	// if node is unchecked
	if stagedNode.IsChecked && !frontNode.IsChecked {

	}

	// if node is checked, add gongstructshape
	if !stagedNode.IsChecked && frontNode.IsChecked {

	}

	// the node was selected. Therefore, one request the
	// table to route to the table
	log.Println("NodeImplGongstruct:OnAfterUpdate with: ", nodeImplGongstruct.gongStruct.GetName())

	nodeImplGongstruct.gongtableStage.Reset()

	tableStage := nodeImplGongstruct.gongtableStage

	if nodeImplGongstruct.gongStruct.GetName() == "Astruct" {

		setOfAstructs := (*models.GetGongstructInstancesSet[models.Astruct](nodeImplGongstruct.stageOfInterest))

		nbColumns := 1
		table := new(gongtable_models.Table).Stage(tableStage)
		table.Name = "Table"
		table.HasColumnSorting = false
		table.HasFiltering = false
		table.HasPaginator = false
		table.HasCheckableRows = false
		table.HasSaveButton = false

		for j := 0; j < nbColumns; j++ {
			column := new(gongtable_models.DisplayedColumn).Stage(tableStage)
			column.Name = "Name"
			table.DisplayedColumns = append(table.DisplayedColumns, column)
		}

		for astruct := range setOfAstructs {
			row := new(gongtable_models.Row).Stage(tableStage)
			row.Name = astruct.Name
			table.Rows = append(table.Rows, row)

			for j := 0; j < nbColumns; j++ {
				cell := new(gongtable_models.Cell).Stage(tableStage)
				cell.Name = astruct.Name

				cellString := new(gongtable_models.CellString).Stage(tableStage)
				cellString.Name = cell.Name
				cellString.Value = cell.Name
				cell.CellString = cellString

				row.Cells = append(row.Cells, cell)
			}
		}
	}

	tableStage.Commit()
}
