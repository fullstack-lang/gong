// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"github.com/gin-gonic/gin"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"

	ssg_stack "github.com/fullstack-lang/gong/lib/ssg/go/level1stack"
	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"

	load "github.com/fullstack-lang/gong/lib/load/go/models"
	load_stack "github.com/fullstack-lang/gong/lib/load/go/stack"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"

	markdown "github.com/fullstack-lang/gong/lib/markdown/go/models"
	markdown_stack "github.com/fullstack-lang/gong/lib/markdown/go/stack"
)

type Stager struct {
	stage       *Stage
	splitStage  *split.Stage
	asSplitArea *split.AsSplitArea

	svgStage  *svg.Stage
	svgObject *svg.SVG

	ssgStage  *ssg.Stage
	treeStage *tree.Stage

	loadStage     *load.Stage
	fileName      string // fileName is used to store the name of the file to load or save
	buttonStage   *button.Stage
	markdownStage *markdown.Stage

	probeForm ProbeIF

	desk *Desk

	// necessary for "--dsm"
	map_Element_Diagrams map[AbstractType][]*Diagram
}

func NewStager(
	r *gin.Engine,
	stage *Stage,
	probeForm ProbeIF,
) (stager *Stager) {
	stager = new(Stager)
	stager.probeForm = probeForm

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage

	stager.stage = stage
	stage.SetDeltaMode(true)

	stager.svgStage = svg_stack.NewStack(r, stage.GetName(), "", "", "", true, true).Stage
	stager.ssgStage = ssg_stack.NewLevel1Stack("", "", "", true, true).Stage
	stager.treeStage = tree_stack.NewStack(r, stage.GetName(), "", "", "", true, true).Stage
	stager.loadStage = load_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.buttonStage = button_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.markdownStage = markdown_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.createViews()

	beforeCommit := func(stage *Stage) {
		stager.enforceSemantic()
	}
	afterCommit := func(stage *Stage) {
		stager.ux_tree()
		stager.ux_svg()
		stager.ux_button()
		stager.load()
		stager.ux_markdown()
	}

	stager.stage.RegisterBeforeCommit(beforeCommit)
	stager.stage.RegisterAfterCommit(afterCommit)
	beforeCommit(stager.stage)
	afterCommit(stager.stage)

	return
}

func (stager *Stager) GetAsSplitArea() (asSplitArea *split.AsSplitArea) {
	asSplitArea = stager.asSplitArea
	return
}

func (stager *Stager) GetGongtreeStage() *tree.Stage {
	return stager.treeStage
}
