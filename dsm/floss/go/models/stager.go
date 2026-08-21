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

	load_fullstack "github.com/fullstack-lang/gong/lib/load/go/fullstack"
	load "github.com/fullstack-lang/gong/lib/load/go/models"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"

	slider "github.com/fullstack-lang/gong/lib/slider/go/models"
	slider_stack "github.com/fullstack-lang/gong/lib/slider/go/stack"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage
	probeForm  ProbeIF

	treeStage             *tree.Stage
	systemDiagramSvgStage *svg.Stage
	flossDiagramSvgStage  *svg.Stage
	sliderStage           *slider.Stage
	ssgStage              *ssg.Stage
	loadStage             *load.Stage
	fileName              string // fileName is used to store the name of the file to load or save
	buttonStage           *button.Stage

	svgObjectDiagramFloss *svg.SVG
	diagramFloss          *DiagramFloss // diagram is the current diagram being displayed

	// present in all "dsm" applications
	// map to navigate from abstract elements to all diagrams where they are displayed
	map_Element_Diagrams map[AbstractType][]*DiagramFloss
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

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage
	stager.treeStage = tree_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.ssgStage = ssg_stack.NewLevel1Stack("", "", "", true, true).Stage
	stager.systemDiagramSvgStage = svg_stack.NewStack(r, "system diagram svg", "", "", "", true, true).Stage
	stager.flossDiagramSvgStage = svg_stack.NewStack(r, "floss diagram svg", "", "", "", true, true).Stage
	stager.sliderStage = slider_stack.NewStack(r, "floss sliders", "", "", "", true, true).Stage
	stager.loadStage, _ = load_fullstack.NewStackInstance(r, "")
	stager.buttonStage = button_stack.NewStack(r, "", "", "", "", true, true).Stage

	stager.createViews()

	// Setup your before commit sequence

	beforeCommit := func(stage *Stage) {
		stager.enforceSemantic()
	}
	afterCommit := func(stage *Stage) {
		stager.ux_tree()
		stager.svg()
		stager.ux_slider()
		stager.createViews()
		stager.button()
		stager.load()
	}


	stager.stage.RegisterBeforeCommit(beforeCommit)
	stager.stage.RegisterAfterCommit(afterCommit)
	beforeCommit(stager.stage)
	afterCommit(stager.stage)

	return stager
}

var _ slider.Target = (*Stager)(nil)

func (stager *Stager) GetSliderStage() *slider.Stage {
	return stager.sliderStage
}

func (stager *Stager) OnAfterUpdateSliderElement() {
	stager.enforceSemantic()
	stager.svg()
	stager.stage.CommitWithSuspendedCallbacks()
}


func (stager *Stager) GetSvgObject() *svg.SVG {
	return stager.svgObjectDiagramFloss
}

