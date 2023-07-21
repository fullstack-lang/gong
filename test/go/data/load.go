package data

import (
	"embed"
	"sort"

	"github.com/gin-gonic/gin"

	gongrouter_fullstack "github.com/fullstack-lang/gongrouter/go/fullstack"
	gongrouter_models "github.com/fullstack-lang/gongrouter/go/models"

	gongtree_buttons "github.com/fullstack-lang/gongtree/go/buttons"
	gongtree_fullstack "github.com/fullstack-lang/gongtree/go/fullstack"
	gongtree_models "github.com/fullstack-lang/gongtree/go/models"

	gong_fullstack "github.com/fullstack-lang/gong/go/fullstack"
	gong_models "github.com/fullstack-lang/gong/go/models"
)

func Load(
	r *gin.Engine,
	goModelsDir embed.FS,
	stackPath string) {

	gongStage := gong_fullstack.NewStackInstance(r, stackPath)

	gong_models.LoadEmbedded(gongStage, goModelsDir)

	gongtreeStage := gongtree_fullstack.NewStackInstance(r, stackPath)

	// configure routing of table and editor router
	gongrouterStage := gongrouter_fullstack.NewStackInstance(r, stackPath)
	tableRouter := new(gongrouter_models.Outlet).Stage(gongrouterStage)
	tableRouter.Name = "github_com_fullstack_lang_gong_test_go" + "_table" + "_" + stackPath

	editorRouter := new(gongrouter_models.Outlet).Stage(gongrouterStage)
	editorRouter.Name = "github_com_fullstack_lang_gong_test_go" + "_editor" + "_" + stackPath

	// create tree
	// set up the gongTree to display elements
	treeOfGongObjects := (&gongtree_models.Tree{Name: "gong"}).Stage(gongtreeStage)
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

	for _, gongStruct := range sliceOfGongStructsSorted {

		nodeGongstruct := (&gongtree_models.Node{Name: gongStruct.Name}).Stage(gongtreeStage)
		nodeGongstruct.IsNodeClickable = true

		nodeGongstruct.Impl = NewNodeImplGongstruct(gongStruct, gongrouterStage, tableRouter)

		// add add button
		addButton := (&gongtree_models.Button{
			Name: gongStruct.Name + " " + string(gongtree_buttons.BUTTON_add),
			Icon: string(gongtree_buttons.BUTTON_add)}).Stage(gongtreeStage)
		nodeGongstruct.Buttons = append(nodeGongstruct.Buttons, addButton)
		addButton.Impl = NewButtonImplGongstruct(
			gongtree_buttons.BUTTON_add,
		)

		treeOfGongObjects.RootNodes = append(treeOfGongObjects.RootNodes, nodeGongstruct)
	}
	gongtreeStage.Commit()

	gongrouterStage.Commit()

}
