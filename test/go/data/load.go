package data

import (
	"embed"
	"sort"

	"github.com/gin-gonic/gin"

	gongrouter_fullstack "github.com/fullstack-lang/gongrouter/go/fullstack"

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
	gongrouterStage := gongrouter_fullstack.NewStackInstance(r, stackPath)
	_ = gongrouterStage

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
		treeOfGongObjects.RootNodes = append(treeOfGongObjects.RootNodes, nodeGongstruct)
	}
	gongtreeStage.Commit()
}
