package data

import (
	"embed"
	"fmt"
	"sort"

	"github.com/gin-gonic/gin"

	gongrouter_fullstack "github.com/fullstack-lang/gongrouter/go/fullstack"
	gongrouter_models "github.com/fullstack-lang/gongrouter/go/models"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_fullstack "github.com/fullstack-lang/gongtree/go/fullstack"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	gongtable_fullstack "github.com/fullstack-lang/gongtable/go/fullstack"
	gongtable_models "github.com/fullstack-lang/gongtable/go/models"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong_models "github.com/fullstack-lang/gong/go/models"

	"github.com/fullstack-lang/gong/test/go/models"
)

func Load(
	r *gin.Engine,
	goModelsDir embed.FS,
	stackPath string,
	stageOfInterest *models.StageStruct) {

	gongStage := gong_fullstack.NewStackInstance(r, stackPath)

	gong_models.LoadEmbedded(gongStage, goModelsDir)

	gongtreeStageLegacy := gongtree_fullstack.NewStackInstance(r, stackPath)

	// treeForSelectingDate that is on the sidebar
	stageForSidebarTree := gongtree_fullstack.NewStackInstance(r, stackPath+"-sidebar")
	stageForMainTable := gongtable_fullstack.NewStackInstance(r, stackPath)
	fillUpSelectTableWithDummyStuff(stageForMainTable, "Table")
	stageForMainTable.Commit()

	// (legacy) configure routing of table and editor router
	gongrouterStage := gongrouter_fullstack.NewStackInstance(r, stackPath)
	tableRouter := new(gongrouter_models.Outlet).Stage(gongrouterStage)
	tableRouter.Name = "github_com_fullstack_lang_gong_test_go" + "_table" + "_" + stackPath

	editorRouter := new(gongrouter_models.Outlet).Stage(gongrouterStage)
	editorRouter.Name = "github_com_fullstack_lang_gong_test_go" + "_editor" + "_" + stackPath
	// end of legacy

	// create tree
	// set up the gongTree to display elements
	treeOfGongObjectsLegacy := (&gongtree_models.Tree{Name: "gong"}).Stage(gongtreeStageLegacy)

	treeOfGongStructs := (&gongtree_models.Tree{Name: "gong"}).Stage(stageForSidebarTree)

	// collect all gong struct to construe the true
	setOfGongStructs := *gong_models.GetGongstructInstancesSet[gong_models.GongStruct](gongStage)

	sliceOfGongStructsSorted := make([]*gong_models.GongStruct, len(setOfGongStructs))
	i := 0
	for k := range setOfGongStructs {
		sliceOfGongStructsSorted[i] = k
		i++
	}
	sort.Slice(sliceOfGongStructsSorted, func(i, j int) bool {
		return sliceOfGongStructsSorted[i].Name < sliceOfGongStructsSorted[j].Name
	})

	for _, gongStructLegacy := range sliceOfGongStructsSorted {

		nodeGongstruct := (&gongtree_models.Node{Name: gongStructLegacy.Name}).Stage(gongtreeStageLegacy)
		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewNodeImplGongstructLegacy(gongStructLegacy, gongrouterStage, tableRouter)

		// add add button
		addButton := (&gongtree_models.Button{
			Name: gongStructLegacy.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(gongtreeStageLegacy)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstructLegacy(
			gongStructLegacy,
			gongtree_buttons.BUTTON_add,
			gongrouterStage,
			editorRouter,
		)

		treeOfGongObjectsLegacy.RootNodes = append(treeOfGongObjectsLegacy.RootNodes, nodeGongstruct)
	}
	gongtreeStageLegacy.Commit()

	for _, gongStruct := range sliceOfGongStructsSorted {

		nodeGongstruct := (&gongtree_models.Node{Name: gongStruct.Name}).Stage(stageForSidebarTree)
		nodeGongstruct.IsNodeClickable = true
		nodeGongstruct.Impl = NewNodeImplGongstruct(gongStruct, stageForMainTable, stageOfInterest)

		// add add button
		addButton := (&gongtree_models.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(stageForSidebarTree)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongStruct,
			gongtree_buttons.BUTTON_add,
		)

		treeOfGongStructs.RootNodes = append(treeOfGongStructs.RootNodes, nodeGongstruct)
	}
	stageForSidebarTree.Commit()

	gongrouterStage.Commit()

}

func fillUpSelectTableWithDummyStuff(stage *gongtable_models.StageStruct, tableName string) {
	nbRows := 0
	nbColumns := 1
	table := new(gongtable_models.Table).Stage(stage)
	table.Name = tableName
	table.HasColumnSorting = false
	table.HasFiltering = false
	table.HasPaginator = false
	table.HasCheckableRows = false
	table.HasSaveButton = false

	for j := 0; j < nbColumns; j++ {
		column := new(gongtable_models.DisplayedColumn).Stage(stage)
		column.Name = "Select a Struct on the left tab to view instances"
		table.DisplayedColumns = append(table.DisplayedColumns, column)
	}

	for i := 0; i < nbRows; i++ {
		row := new(gongtable_models.Row).Stage(stage)
		row.Name = fmt.Sprintf("Row %d", i)
		table.Rows = append(table.Rows, row)

		if i%2 == 0 {
			row.IsChecked = true
		}

		for j := 0; j < nbColumns; j++ {
			cell := new(gongtable_models.Cell).Stage(stage)
			cell.Name = fmt.Sprintf("Row %d - Column %d", i, j)

			cellString := new(gongtable_models.CellString).Stage(stage)
			cellString.Name = cell.Name
			cellString.Value = cell.Name
			cell.CellString = cellString

			row.Cells = append(row.Cells, cell)
		}
	}
}
