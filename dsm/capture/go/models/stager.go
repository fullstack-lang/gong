// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"github.com/gin-gonic/gin"

	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"

	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	ssg_stack "github.com/fullstack-lang/gong/lib/ssg/go/level1stack"
	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	load_stack "github.com/fullstack-lang/gong/lib/load/go/stack"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage
	probeForm  ProbeIF



	treeStage   *tree.Stage
	svgStage    *svg.Stage
	ssgStage    *ssg.Stage
	loadStage   *load.Stage
	fileName    string // fileName is used to store the name of the file to load or save
	buttonStage *button.Stage

	productToLibrary map[*Deliverable]*Library
	taskToLibrary    map[*Concern]*Library
	// reverse map
	map_Concern_Stakeholder map[*Concern][]*Stakeholder

	svgObject *svg.SVG
	diagram   *Diagram

	// map to navigate from abstract elements to all diagrams where they are displayed
	map_Element_Diagrams map[AbstractType][]*Diagram
}

func NewStager(
	r *gin.Engine,
	stage *Stage,
	probeForm ProbeIF,
) (stager *Stager) {
	stager = new(Stager)

	stager.stage = stage
	stager.probeForm = probeForm

	// enable delta mode on the stage
	// so that changes are tracked and undo/redo are possible
	stage.SetDeltaMode(true)

	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage
	stager.treeStage = tree_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.ssgStage = ssg_stack.NewLevel1Stack("", "", "", true, true).Stage
	stager.svgStage = svg_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.loadStage = load_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.buttonStage = button_stack.NewStack(r, "", "", "", "", true, true).Stage

	stager.createViews()

	// Setup your before commit sequence

	beforeCommit := func(stage *Stage) {
		stager.enforceThereIsARootLibrary()
		stager.enforceSemantic()
	}
	afterCommit := func(stage *Stage) {
		stager.ux_tree()
		stager.svg()
		stager.button()
		stager.load()
	}

	stager.stage.RegisterBeforeCommit(beforeCommit)
	stager.stage.RegisterAfterCommit(afterCommit)
	beforeCommit(stager.stage)
	afterCommit(stager.stage)

	return
}
