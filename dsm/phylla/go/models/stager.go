// generated boilerplate code
// edit the file for adding other stages
package models

import (
	"github.com/gin-gonic/gin"

	button "github.com/fullstack-lang/gong/lib/button/go/models"
	button_stack "github.com/fullstack-lang/gong/lib/button/go/stack"

	load_fullstack "github.com/fullstack-lang/gong/lib/load/go/fullstack"
	load "github.com/fullstack-lang/gong/lib/load/go/models"

	slider "github.com/fullstack-lang/gong/lib/slider/go/models"
	slider_stack "github.com/fullstack-lang/gong/lib/slider/go/stack"

	split "github.com/fullstack-lang/gong/lib/split/go/models"
	split_stack "github.com/fullstack-lang/gong/lib/split/go/stack"

	ssg_stack "github.com/fullstack-lang/gong/lib/ssg/go/level1stack"
	ssg "github.com/fullstack-lang/gong/lib/ssg/go/models"

	threejs "github.com/fullstack-lang/gong/lib/threejs/go/models"
	threejs_stack "github.com/fullstack-lang/gong/lib/threejs/go/stack"

	tree "github.com/fullstack-lang/gong/lib/tree/go/models"
	tree_stack "github.com/fullstack-lang/gong/lib/tree/go/stack"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"
)

type Stager struct {
	stage      *Stage
	splitStage *split.Stage
	probeForm  ProbeIF

	buttonStage  *button.Stage  // "buttonStage" is the DSM mandatory name (to be changed)
	loadStage    *load.Stage    // mandatory
	threejsStage *threejs.Stage // "treeStage" is the DSM mandatory name (to be changed)

	treeStage2D *tree.Stage
	treeStage3D *tree.Stage
	sliderStage *slider.Stage
	ssgStage    *ssg.Stage // mandatory
	svgStage    *svg.Stage

	svgObject *svg.SVG

	// DSM mandatory
	// map to navigate from abstract elements to all diagrams where they are displayed
	map_Element_Diagrams map[AbstractType][]DiagramIF

	// DSM mandatory
	fileName        string // fileName is used to store the name of the file to load or save
	persistanceFile string

	// DSM specific
	// the plant that is currently selected for the form
	selectedPlant *Plant

	// maps
	m_Plant_Library map[*Plant]*Library
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
	stager.sliderStage = slider_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage
	stager.ssgStage = ssg_stack.NewLevel1Stack("", "", "", true, true).Stage
	stager.svgStage = svg_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.threejsStage = threejs_stack.NewStack(r, "", "", "", "", true, true).Stage

	stager.treeStage2D = tree_stack.NewStack(r, "treeStage2D", "", "", "", true, true).Stage
	stager.treeStage3D = tree_stack.NewStack(r, "treeStage3D", "", "", "", true, true).Stage

	stager.createViews()

	beforeCommit := func(stage *Stage) {
		stager.enforceSemantic()
	}
	afterCommit := func(stage *Stage) {
		stager.ux_tree() // DSM mandatory name, to be changed
		stager.button()
		stager.load()
		stager.ux_slider()
		stager.ux_svg_plant_diagram()
		stager.ux_3d_plant_diagram()
	}

	stager.stage.RegisterBeforeCommit(beforeCommit)
	stager.stage.RegisterAfterCommit(afterCommit)

	beforeCommit(stager.stage)
	afterCommit(stager.stage)

	for plant := range *GetGongstructInstancesSetFromPointerType[*Plant](stage) {
		if plant.IsSelected {
			stager.probeForm.FillUpFormFromGongstruct(plant, GetPointerToGongstructName[*Plant]())
			break
		}
	}
	return
}

func (stager *Stager) GetSvgStage() *svg.Stage {
	return stager.svgStage
}

type BeforeCommitImplementation struct {
	stager *Stager
}

func (c *BeforeCommitImplementation) BeforeCommit(stage *Stage) {

}

func (stager *Stager) GetSvgObject() *svg.SVG {
	return stager.svgObject
}

func (stager *Stager) GetCurrentPlant() *Plant {
	return stager.selectedPlant
}

func (stager *Stager) GetSliderStage() *slider.Stage {
	return stager.sliderStage
}

func (stager *Stager) GetThreejsStage() *threejs.Stage {
	return stager.threejsStage
}
