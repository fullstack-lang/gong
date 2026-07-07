// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"github.com/gin-gonic/gin"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"

	load_fullstack "github.com/fullstack-lang/gong/lib/load/go/fullstack"
	load "github.com/fullstack-lang/gong/lib/load/go/models"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"

	ssg_stack "github.com/fullstack-lang/gong/lib/ssg/go/level1stack"
	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage
	probeForm  ProbeIF

	buttonStage *button.Stage // "buttonStage" is the DSM mandatory name (to be changed)
	loadStage   *load.Stage   // mandatory

	treeStage *tree.Stage // "treeStage" is the DSM mandatory name (to be changed)
	ssgStage  *ssg.Stage  // mandatory

	svgStage *svg.Stage

	svgObject *svg.SVG

	// DSM mandatory
	// map to navigate from abstract elements to all diagrams where they are displayed
	map_Element_Diagrams map[AbstractType][]DiagramIF

	// DSM mandatory
	fileName        string // fileName is used to store the name of the file to load or save
	persistanceFile string
}

func NewStager(
	r *gin.Engine,
	stage *Stage,
	probeForm ProbeIF,
) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.probeForm = probeForm

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.buttonStage = button_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.loadStage, _ = load_fullstack.NewStackInstance(r, "")

	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage
	stager.ssgStage = ssg_stack.NewLevel1Stack("", "", "", true, true).Stage
	stager.treeStage = tree_stack.NewStack(r, "", "", "", "", true, true).Stage

	split.StageBranch(stager.splitStage, &split.View{
		Name: "Data Probe & Data Model",
		RootAsSplitAreas: []*split.AsSplitArea{
			{
				Split: &split.Split{
					StackName: stage.GetProbeSplitStageName(),
				},
			},
		},
	})

	stager.splitStage.Commit()

	callbacks := &BeforeCommitImplementation{
		stager: stager,
	}
	stager.stage.OnInitCommitFromBackCallback = callbacks
	callbacks.BeforeCommit(stage)

	return
}

type BeforeCommitImplementation struct {
	stager *Stager
}

func (c *BeforeCommitImplementation) BeforeCommit(stage *Stage) {

}

func (stager *Stager) GetSvgObject() *svg.SVG {
	return stager.svgObject
}
