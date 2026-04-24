// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"os"

	"github.com/gin-gonic/gin"

	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"
	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	split "github.com/fullstack-lang/gong/lib/split/go/models"
	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	load_stack "github.com/fullstack-lang/gong/lib/load/go/stack"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage

	treeObjectsSimulationStage *tree.Stage
	treeDiagramStage           *tree.Stage
	svgStage                   *svg.Stage
	loadStage                  *load.Stage
	buttonTransitionsStage     *button.Stage
	buttonExportXLStage        *button.Stage

	// maps

	// singloton architecture
	architecture *Architecture

	set_StartStates map[*State]struct{}

	map_state_nextStates map[*State][]*State

	map_stateMachine_objects map[*StateMachine][]*Object
	map_state_stateMachine   map[*State]*StateMachine
	map_diagram_stateMachine map[*Diagram]*StateMachine

	probeForm ProbeIF

	diagram                *Diagram
	svgObject              *svg.SVG
	map_SvgRect_StateShape map[*svg.Rect]*StateShape

	filename string
}

func (stager *Stager) GetStage() *Stage {
	return stager.stage
}

// OnAfterUpdateButton implements models.Target.
func (stager *Stager) OnAfterUpdateButton() {

}

func (stager *Stager) GetButtonsStage() *button.Stage {
	return stager.buttonTransitionsStage
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
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage

	stackName := "statemachines"
	stager.treeObjectsSimulationStage = tree_stack.NewStack(r, stackName+"-objects", "", "", "", true, true).Stage
	stager.treeDiagramStage = tree_stack.NewStack(r, stackName+"-diagrams", "", "", "", true, true).Stage
	stager.svgStage = svg_stack.NewStack(r, stackName, "", "", "", true, true).Stage
	stager.loadStage = load_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.buttonTransitionsStage = button_stack.NewStack(r, stackName+"-transitions", "", "", "", false, false).Stage
	stager.buttonExportXLStage = button_stack.NewStack(r, stackName+"-exportXL", "", "", "", true, true).Stage

	stager.create_views()

	beforeCommit := func(stage *Stage) {
		stager.enforce_semantic()
	}
	afterCommit := func(stage *Stage) {
		stager.treeSimulation()
		stager.buttonSimulation()
		stager.buttonsExportXL()
		stager.svg()
		stager.load()
		stager.treeDiagrams()
	}

	stager.stage.RegisterBeforeCommit(beforeCommit)
	stager.stage.RegisterAfterCommit(afterCommit)
	beforeCommit(stager.stage)
	afterCommit(stager.stage)

	// hook the stage on a kill command
	// 	curl --request POST \
	//   --url http://localhost:8000/api/gongstatesmachines/go/v1/kills?Name="gongstatesmachines" \
	//   --header 'content-type: application/json' \
	//   --data '{"Name": "Die !"}'
	stage.OnAfterKillCreateCallback = &OnAfterKillCreateCallback{}

	return
}

type Kill struct {
	Name string
}

type OnAfterKillCreateCallback struct {
}

// OnAfterCreate implements OnAfterCreateInterface.
func (o *OnAfterKillCreateCallback) OnAfterCreate(stage *Stage, instance *Kill) {
	os.Exit(0)
}

type OnInitCommitFromBackCallback struct {
	stager *Stager
}
