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

	form "github.com/fullstack-lang/gong/lib/form/go/models"
	form_stack "github.com/fullstack-lang/gong/lib/form/go/stack"

	svg "github.com/fullstack-lang/gong/lib/svg/go/models"
	svg_stack "github.com/fullstack-lang/gong/lib/svg/go/stack"

	markdown "github.com/fullstack-lang/gong/lib/markdown/go/models"
	markdown_stack "github.com/fullstack-lang/gong/lib/markdown/go/stack"

	cursor "github.com/fullstack-lang/gong/lib/cursor/go/models"
	cursor_stack "github.com/fullstack-lang/gong/lib/cursor/go/stack"

	tone "github.com/fullstack-lang/gong/lib/tone/go/models"
	tone_stack "github.com/fullstack-lang/gong/lib/tone/go/stack"
)

type ThreeJSStageUpdaterInterface interface {
	UpdateThreeJSStage(stager *Stager)
	StartMovieRecordingVase3D(stager *Stager, plant *PlantAbstract, vase3DDiagram *Vase3DDiagram)
	StopMovieRecording(stager *Stager)
	IsMovieRecording() bool
	GetMovieRecordingFrameCount() int
}

type Stool3DStageUpdaterInterface interface {
	UpdateStool3DStage(stager *Stager)
}

type Clock3DStageUpdaterInterface interface {
	UpdateClock3DStage(stager *Stager)
}

type Plant3DStageUpdaterInterface interface {
	UpdatePlant3DStage(stager *Stager)
}

type Stager struct {
	stage      *Stage
	splitStage *split.Stage
	probeForm  ProbeIF

	buttonStage  *button.Stage  // "buttonStage" is the DSM mandatory name (to be changed)
	loadStage    *load.Stage    // mandatory
	threejsStage     *threejs.Stage // "treeStage" is the DSM mandatory name (to be changed)
	stool3dStage     *threejs.Stage
	clock3dStage     *threejs.Stage
	plant3dStage     *threejs.Stage

	treeStage2D      *tree.Stage
	treeStage3D      *tree.Stage
	sliderStage      *slider.Stage
	sliderStoolStage *slider.Stage
	sliderClockStage *slider.Stage
	sliderMusicStage *slider.Stage
	plantFormStage   *form.Stage
	ssgStage         *ssg.Stage // mandatory
	svgPlantStage    *svg.Stage
	svgVaseStage     *svg.Stage
	svgMusicStage    *svg.Stage
	svgStage         *svg.Stage
	buttonMusicStage *button.Stage
	toneStage        *tone.Stage
	cursorStage      *cursor.Stage
	cursor           *cursor.Cursor
	markdownStage    *markdown.Stage

	svgObject *svg.SVG

	musicNotes          []*MusicNoteData
	circumferenceLength float64

	// DSM mandatory
	// map to navigate from abstract elements to all diagrams where they are displayed
	map_Element_Diagrams map[AbstractType][]DiagramIF

	// DSM mandatory
	fileName        string // fileName is used to store the name of the file to load or save
	persistanceFile string

	// DSM specific
	// the plant that is currently selected for the form
	selectedPlant *PlantAbstract

	threeJSUpdater ThreeJSStageUpdaterInterface
	stool3DUpdater Stool3DStageUpdaterInterface
	clock3DUpdater Clock3DStageUpdaterInterface
	plant3DUpdater Plant3DStageUpdaterInterface

	// maps
	m_Plant_Library map[*PlantAbstract]*Library
}

func NewStager(
	r *gin.Engine,
	stage *Stage,
	probeForm ProbeIF,
	persistanceFile string,
	threeJSUpdater ThreeJSStageUpdaterInterface,
	stool3DUpdater Stool3DStageUpdaterInterface,
	clock3DUpdater Clock3DStageUpdaterInterface,
	plant3DUpdater Plant3DStageUpdaterInterface,
) (stager *Stager) {

	stager = new(Stager)

	stager.stage = stage
	stager.probeForm = probeForm
	stager.persistanceFile = persistanceFile
	stager.threeJSUpdater = threeJSUpdater
	stager.stool3DUpdater = stool3DUpdater
	stager.clock3DUpdater = clock3DUpdater
	stager.plant3DUpdater = plant3DUpdater

	// the root split name is "" by convention. Is is the same for all gong applications
	// that do not develop their specific angular component
	stager.buttonStage = button_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.loadStage, _ = load_fullstack.NewStackInstance(r, "")
	stager.sliderStage = slider_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.sliderStoolStage = slider_stack.NewStack(r, "sliderStoolStage", "", "", "", true, true).Stage
	stager.sliderClockStage = slider_stack.NewStack(r, "sliderClockStage", "", "", "", true, true).Stage
	stager.sliderMusicStage = slider_stack.NewStack(r, "sliderMusicStage", "", "", "", true, true).Stage
	stager.splitStage = split_stack.NewStack(r, "", "", "", "", false, false).Stage
	stager.ssgStage = ssg_stack.NewLevel1Stack("", "", "", true, true).Stage
	stager.svgPlantStage = svg_stack.NewStack(r, "svgPlantStage", "", "", "", true, true).Stage
	stager.svgVaseStage = svg_stack.NewStack(r, "svgVaseStage", "", "", "", true, true).Stage
	stager.svgMusicStage = svg_stack.NewStack(r, "svgMusicStage", "", "", "", true, true).Stage
	stager.svgStage = stager.svgPlantStage
	stager.buttonMusicStage = button_stack.NewStack(r, "buttonMusicStage", "", "", "", true, true).Stage
	stager.toneStage = tone_stack.NewStack(r, "toneStage", "", "", "", true, true).Stage
	stager.cursorStage = cursor_stack.NewStack(r, "cursorStage", "", "", "", true, true).Stage
	stager.threejsStage = threejs_stack.NewStack(r, "", "", "", "", true, true).Stage
	stager.stool3dStage = threejs_stack.NewStack(r, "stool3d", "", "", "", true, true).Stage
	stager.clock3dStage = threejs_stack.NewStack(r, "clock3d", "", "", "", true, true).Stage
	stager.plant3dStage = threejs_stack.NewStack(r, "plant3d", "", "", "", true, true).Stage
	stager.markdownStage = markdown_stack.NewStack(r, "", "", "", "", true, true).Stage

	stager.cursor = new(cursor.Cursor).Stage(stager.cursorStage)
	stager.cursorStage.Commit()

	stager.treeStage2D = tree_stack.NewStack(r, "treeStage2D", "", "", "", true, true).Stage
	stager.treeStage3D = tree_stack.NewStack(r, "treeStage3D", "", "", "", true, true).Stage
	stager.plantFormStage = form_stack.NewStack(r, "plantFormStage", "", "", "", true, true).Stage
	form.SetOrchestratorOnAfterUpdate[form.FormGroup](stager.plantFormStage)

	stager.createViews()

	beforeCommit := func(stage *Stage) {
		stager.enforceSemantic()
	}
	afterCommit := func(stage *Stage) {
		stager.createViews()
		stager.ux_tree() // DSM mandatory name, to be changed
		stager.button()
		stager.load()
		stager.ux_markdown()
		stager.updateSelectedViewFromPlant(stager.GetCurrentPlant())
		stager.ux_slider()
		stager.ux_slider_stool()
		stager.ux_slider_clock()
		stager.ux_slider_music()
		stager.ux_button_music()
		stager.ux_plant_form()
		stager.ux_svg_plant_diagram()
		stager.ux_svg_music()
		stager.ux_tone_music()
		stager.ux_cursor_music()
		stager.UpdateThreeJSStage()
		stager.UpdateStool3DStage()
		stager.UpdateClock3DStage()
		stager.UpdatePlant3DStage()
	}

	stager.stage.RegisterBeforeCommit(beforeCommit)
	stager.stage.RegisterAfterCommit(afterCommit)

	beforeCommit(stager.stage)
	afterCommit(stager.stage)

	for plant := range *GetGongstructInstancesSetFromPointerType[*PlantAbstract](stage) {
		if plant.IsSelected {
			stager.probeForm.FillUpFormFromGongstruct(plant, GetPointerToGongstructName[*PlantAbstract]())
			break
		}
	}
	return
}

func (stager *Stager) GetStage() *Stage {
	return stager.stage
}

func (stager *Stager) EnforceSemantic() {
	stager.enforceSemantic()
}

func (stager *Stager) UxSlider() {
	stager.ux_slider()
}

func (stager *Stager) SetThreeJSUpdater(updater ThreeJSStageUpdaterInterface) {
	stager.threeJSUpdater = updater
}

func (stager *Stager) UpdateThreeJSStage() {
	if stager.threeJSUpdater != nil {
		stager.threeJSUpdater.UpdateThreeJSStage(stager)
	}
}

func (stager *Stager) SetStool3DUpdater(updater Stool3DStageUpdaterInterface) {
	stager.stool3DUpdater = updater
}

func (stager *Stager) UpdateStool3DStage() {
	if stager.stool3DUpdater != nil {
		stager.stool3DUpdater.UpdateStool3DStage(stager)
	}
}

func (stager *Stager) StartMovieRecordingVase3D(plant *PlantAbstract, vase3DDiagram *Vase3DDiagram) {
	if stager.threeJSUpdater != nil {
		stager.threeJSUpdater.StartMovieRecordingVase3D(stager, plant, vase3DDiagram)
	}
}

func (stager *Stager) StopMovieRecording() {
	if stager.threeJSUpdater != nil {
		stager.threeJSUpdater.StopMovieRecording(stager)
	}
}

func (stager *Stager) IsMovieRecording() bool {
	if stager.threeJSUpdater != nil {
		return stager.threeJSUpdater.IsMovieRecording()
	}
	return false
}

func (stager *Stager) GetMovieRecordingFrameCount() int {
	if stager.threeJSUpdater != nil {
		return stager.threeJSUpdater.GetMovieRecordingFrameCount()
	}
	return 0
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

func (stager *Stager) GetCurrentPlant() *PlantAbstract {
	return stager.selectedPlant
}

func (stager *Stager) GetSliderStage() *slider.Stage {
	return stager.sliderStage
}

func (stager *Stager) GetThreejsStage() *threejs.Stage {
	return stager.threejsStage
}

func (stager *Stager) GetStool3dStage() *threejs.Stage {
	return stager.stool3dStage
}

func (stager *Stager) GetSliderStoolStage() *slider.Stage {
	return stager.sliderStoolStage
}

func (stager *Stager) SetClock3DUpdater(updater Clock3DStageUpdaterInterface) {
	stager.clock3DUpdater = updater
}

func (stager *Stager) UpdateClock3DStage() {
	if stager.clock3DUpdater != nil {
		stager.clock3DUpdater.UpdateClock3DStage(stager)
	}
}

func (stager *Stager) GetClock3dStage() *threejs.Stage {
	return stager.clock3dStage
}

func (stager *Stager) GetSliderClockStage() *slider.Stage {
	return stager.sliderClockStage
}

func (stager *Stager) GetMarkdownStage() *markdown.Stage {
	return stager.markdownStage
}

func (stager *Stager) SetPlant3DUpdater(updater Plant3DStageUpdaterInterface) {
	stager.plant3DUpdater = updater
}

func (stager *Stager) UpdatePlant3DStage() {
	if stager.plant3DUpdater != nil {
		stager.plant3DUpdater.UpdatePlant3DStage(stager)
	}
}

func (stager *Stager) GetPlant3dStage() *threejs.Stage {
	return stager.plant3dStage
}

