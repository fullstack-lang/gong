package data

import (
	"fmt"
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

	tableStage := nodeImplGongstruct.gongtableStage
	tableStage.Reset()
	tableStage.Commit()

	table := new(gongtable_models.Table).Stage(tableStage)
	table.Name = "Table"
	table.HasColumnSorting = false
	table.HasFiltering = false
	table.HasPaginator = false
	table.HasCheckableRows = false
	table.HasSaveButton = false

	if nodeImplGongstruct.gongStruct.GetName() == "Astruct" {

		fields := models.GetFields[models.Astruct]()

		setOfAstructs := (*models.GetGongstructInstancesSet[models.Astruct](nodeImplGongstruct.stageOfInterest))

		for _, fieldName := range fields {
			column := new(gongtable_models.DisplayedColumn).Stage(tableStage)
			column.Name = fieldName
			table.DisplayedColumns = append(table.DisplayedColumns, column)
		}

		for astruct := range setOfAstructs {
			row := new(gongtable_models.Row).Stage(tableStage)
			row.Name = astruct.Name
			table.Rows = append(table.Rows, row)

			for _, fieldName := range fields {
				value := models.GetFieldStringValue[models.Astruct](*astruct, fieldName)

				log.Println(fieldName, value)
				cell := (&gongtable_models.Cell{
					Name: value,
				}).Stage(tableStage)

				cellString := (&gongtable_models.CellString{
					Name:  value,
					Value: value,
				}).Stage(tableStage)
				cell.CellString = cellString

				row.Cells = append(row.Cells, cell)
			}
		}
	}

	tableStage.Commit()
}

func fillUpTable[T models.Gongstruct](
	nodeImplGongstruct *NodeImplGongstruct,
	tableStage *gongtable_models.StageStruct,
	table *gongtable_models.Table) {

	fields := models.GetFields[T]()

	setOfStructs := (*models.GetGongstructInstancesSet[T](nodeImplGongstruct.stageOfInterest))

	for _, fieldName := range fields {
		column := new(gongtable_models.DisplayedColumn).Stage(tableStage)
		column.Name = fieldName
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	fieldIndex := 0
	for structInstance := range setOfStructs {
		row := new(gongtable_models.Row).Stage(tableStage)
		row.Name = models.GetFieldStringValue[T](*structInstance, "Name")
		table.Rows = append(table.Rows, row)

		for _, fieldName := range fields {
			value := models.GetFieldStringValue[T](*structInstance, fieldName)
			value = fmt.Sprintf("%d", fieldIndex) + " " + value
			fieldIndex++
			log.Println(fieldName, value)
			cell := (&gongtable_models.Cell{
				Name: value,
			}).Stage(tableStage)

			cellString := (&gongtable_models.CellString{
				Name:  value,
				Value: value,
			}).Stage(tableStage)
			cell.CellString = cellString

			row.Cells = append(row.Cells, cell)
		}
	}
}
