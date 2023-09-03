// generated code - do not edit
package data

import (
	"fmt"

	"github.com/gin-gonic/gin"

	gongtable_fullstack "github.com/fullstack-lang/gongtable/go/fullstack"
	form "github.com/fullstack-lang/gongtable/go/models"
	gongtable_models "github.com/fullstack-lang/gongtable/go/models"
)

func NewOnAssocEditon(
	r *gin.Engine,
	formStage *form.StageStruct,
) (onAssocEdition *OnAssocEditon) {

	onAssocEdition = new(OnAssocEditon)

	onAssocEdition.r = r
	onAssocEdition.formStage = formStage

	return
}

type OnAssocEditon struct {
	r         *gin.Engine
	formStage *form.StageStruct
}

func (onAssocEditon *OnAssocEditon) OnButtonPressed() {

	tableStackName := onAssocEditon.formStage.GetPath() + string(form.StackNamePostFixForTableForAssociation)

	// tableStackName supposed to be "test-form-table"
	tableStageForSelection, _ := gongtable_fullstack.NewStackInstance(onAssocEditon.r, tableStackName)

	nbRows := 10
	nbColumns := 1
	table := new(gongtable_models.Table).Stage(tableStageForSelection)
	table.Name = string(form.TableSelectExtraName)
	table.HasColumnSorting = true
	table.HasFiltering = true
	table.HasPaginator = false
	table.HasCheckableRows = true
	table.HasSaveButton = true

	for j := 0; j < nbColumns; j++ {
		column := new(gongtable_models.DisplayedColumn).Stage(tableStageForSelection)
		column.Name = fmt.Sprintf("Column %d", j)
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	for i := 0; i < nbRows; i++ {
		row := new(gongtable_models.Row).Stage(tableStageForSelection)
		row.Name = fmt.Sprintf("Row %d", i)
		table.Rows = append(table.Rows, row)

		if i%2 == 0 {
			row.IsChecked = true
		}

		for j := 0; j < nbColumns; j++ {
			cell := new(gongtable_models.Cell).Stage(tableStageForSelection)
			cell.Name = fmt.Sprintf("Row %d - Column %d", i, j)

			cellString := new(gongtable_models.CellString).Stage(tableStageForSelection)
			cellString.Name = cell.Name
			cellString.Value = cell.Name
			cell.CellString = cellString

			row.Cells = append(row.Cells, cell)
		}
	}
	tableStageForSelection.Commit()
}
