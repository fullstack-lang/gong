// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"

	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"

	ssg_stack "github.com/fullstack-lang/gong/lib/ssg/go/level1stack"
	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	load_fullstack "github.com/fullstack-lang/gong/lib/load/go/fullstack"

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
	buttonStage *button.Stage

	svgObject *svg.SVG
	diagram   *Diagram
	fileName  string

	// necessary for "--dsm"
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
	stager.map_Element_Diagrams = make(map[AbstractType][]*Diagram)

	// enable delta mode on the stage
	// so that changes are tracked and undo/redo are possible
	stage.SetDeltaMode(true)

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage
	stager.treeStage = tree_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.ssgStage = ssg_stack.NewLevel1Stack("", "", "", true, true).Stage
	stager.svgStage = svg_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.loadStage, _ = load_fullstack.NewStackInstance(r, "")
	stager.buttonStage = button_stack.NewStack(r, "", "", "", "", true, true).Stage

	stager.createViews()

	beforeCommit := func(stage *Stage) {
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

type BeforeCommitImplementation struct {
	stager *Stager
}

func (c *BeforeCommitImplementation) BeforeCommit(stage *Stage) {

}
