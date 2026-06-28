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
	load_fullstack "github.com/fullstack-lang/gong/lib/load/go/fullstack"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage
	probeForm  ProbeIF

	treeStage *tree.Stage
	// the tree stage can be very deep. The zoomTreeStage display on the tree starting from the current diagram.
	zoomTreeStage            *tree.Stage
	processDiagramSvgStage   *svg.Stage
	structureDiagramSvgStage *svg.Stage
	ssgStage                 *ssg.Stage
	loadStage                *load.Stage
	fileName                 string // fileName is used to store the name of the file to load or save
	buttonStage              *button.Stage

	svgObjectDiagramProcess *svg.SVG
	diagramProcess          *DiagramProcess // diagram is the current diagram being displayed

	// present in all "dsm" applications
	// map to navigate from abstract elements to all diagrams where they are displayed
	map_Element_Diagrams map[AbstractType][]*DiagramProcess

	// reverse map to navigate
	rm_Data_DataFlows        map[*Data][]*DataFlow
	rm_Resource_Participants map[*Resource][]*Participant
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
	stager.treeStage = tree_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.zoomTreeStage = tree_stack.NewStack(r, "zoom tree", "", "", "", true, true).Stage
	stager.ssgStage = ssg_stack.NewLevel1Stack("", "", "", true, true).Stage
	stager.processDiagramSvgStage = svg_stack.NewStack(r, "process diagram svg", "", "", "", true, true).Stage
	stager.structureDiagramSvgStage = svg_stack.NewStack(r, "structure diagram svg", "", "", "", true, true).Stage
	stager.loadStage, _ = load_fullstack.NewStackInstance(r, "")
	stager.buttonStage = button_stack.NewStack(r, "", "", "", "", true, true).Stage

	stager.createViews()

	// Setup your before commit sequence

	beforeCommit := func(stage *Stage) {
		stager.enforceSemantic()
	}
	afterCommit := func(stage *Stage) {
		stager.ux_tree()
		stager.treeZoom()
		stager.svg()
		stager.button()
		stager.load()
	}

	stager.stage.RegisterBeforeCommit(beforeCommit)
	stager.stage.RegisterAfterCommit(afterCommit)
	beforeCommit(stager.stage)
	afterCommit(stager.stage)

	return stager
}

type BeforeCommitImplementation struct {
	stager *Stager
}

func (c *BeforeCommitImplementation) BeforeCommit(stage *Stage) {
}

func (stager *Stager) GetSvgObject() *svg.SVG {
	return stager.svgObjectDiagramProcess
}
