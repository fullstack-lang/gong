// do not modify, generated file
package level1stack

import (
	"fmt"
	"log"
	"strings"

	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/probe"

	embeddedgo "github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go"

	"github.com/gin-gonic/gin"

	split_static "github.com/fullstack-lang/gong/lib/split/go/static"
)

// hook marhalling to stage
type BeforeCommitImplementation struct {
	marshallOnCommit string

	packageName string
}

func (impl *BeforeCommitImplementation) BeforeCommit(stage *models.Stage) {

	if stage.GetGongMarshallingMode() == models.GongMarshallingAppendCommit {
		stage.ComputeForwardAndBackwardCommits()
		stage.ComputeReferenceAndOrders()
	}

	// the ".go" is not provided
	filename := impl.marshallOnCommit
	if !strings.HasSuffix(filename, ".go") {
		filename = filename + ".go"
	}

	packageName := impl.packageName
	if packageName == "" {
		packageName = "main"
	}

	stage.MarshallFile(fmt.Sprintf("./%s", filename), "github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models", packageName)
}

type Level1Stack struct {
	Stage *models.Stage
	Probe *probe.Probe
	R     *gin.Engine
}

func NewLevel1Stack(
	stackPath string,
	unmarshallFromCode string,
	marshallOnCommit string,
	withProbe bool,
	embeddedDiagrams bool,
) (level1Stack *Level1Stack) {
	return NewLevel1StackDelta(stackPath, unmarshallFromCode, marshallOnCommit, withProbe, embeddedDiagrams, false)
}

func NewLevel1StackDelta(
	stackPath string,
	unmarshallFromCode string,
	marshallOnCommit string,
	withProbe bool,
	embeddedDiagrams bool,
	deltaMode bool,
) (level1Stack *Level1Stack) {

	level1Stack = new(Level1Stack)
	stage := models.NewStage(stackPath)

	if deltaMode {
		stage.SetDeltaMode(true)
	}

	level1Stack.Stage = stage

	level1Stack.R = split_static.ServeStaticFiles(false)
	if withProbe {
		// if the application edits the diagrams via the probe, it is surmised
		// that the application is launched from "go/cmd/<appl>/". Therefore, to reach
		// "go/diagrams/diagrams.go", the path is "../../diagrams/diagrams.go"
		level1Stack.Probe = probe.NewProbe(
			level1Stack.R,
			embeddedgo.GoModelsDir,
			embeddedgo.GoDiagramsDir,
			embeddedDiagrams,
			stage,
		)

		stage.SetProbeIF(level1Stack.Probe)
	}

	if unmarshallFromCode != "" {
		err := models.ParseAstFile(stage, unmarshallFromCode, true)

		// if the application is run with -unmarshallFromCode=xxx.go -marshallOnCommit
		// xxx.go might be absent the first time. However, this shall not be a show stopper.
		if err != nil {
			log.Println("no file to read " + err.Error())
		}

		stage.ComputeReverseMaps()
		stage.ComputeInstancesNb()
		stage.ComputeReferenceAndOrders()
	} else {
		// in case the database is used, checkout the content to the stage
		stage.Checkout()
	}

	// hook automatic marshall to go code at every commit
	if marshallOnCommit != "" {
		hook := new(BeforeCommitImplementation)
		hook.marshallOnCommit = marshallOnCommit
		stage.OnInitCommitCallback = hook
	}

	// add orchestration
	// insertion point
	models.SetOrchestratorOnAfterUpdate[models.A_directive](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_measure](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_measure_1](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_part](stage)
	models.SetOrchestratorOnAfterUpdate[models.A_part_1](stage)
	models.SetOrchestratorOnAfterUpdate[models.Accidental](stage)
	models.SetOrchestratorOnAfterUpdate[models.Accidental_mark](stage)
	models.SetOrchestratorOnAfterUpdate[models.Accidental_text](stage)
	models.SetOrchestratorOnAfterUpdate[models.Accord](stage)
	models.SetOrchestratorOnAfterUpdate[models.Accordion_registration](stage)
	models.SetOrchestratorOnAfterUpdate[models.Appearance](stage)
	models.SetOrchestratorOnAfterUpdate[models.Arpeggiate](stage)
	models.SetOrchestratorOnAfterUpdate[models.Arrow](stage)
	models.SetOrchestratorOnAfterUpdate[models.Articulations](stage)
	models.SetOrchestratorOnAfterUpdate[models.Assess](stage)
	models.SetOrchestratorOnAfterUpdate[models.Attributes](stage)
	models.SetOrchestratorOnAfterUpdate[models.Backup](stage)
	models.SetOrchestratorOnAfterUpdate[models.Bar_style_color](stage)
	models.SetOrchestratorOnAfterUpdate[models.Barline](stage)
	models.SetOrchestratorOnAfterUpdate[models.Barre](stage)
	models.SetOrchestratorOnAfterUpdate[models.Bass](stage)
	models.SetOrchestratorOnAfterUpdate[models.Bass_step](stage)
	models.SetOrchestratorOnAfterUpdate[models.Beam](stage)
	models.SetOrchestratorOnAfterUpdate[models.Beat_repeat](stage)
	models.SetOrchestratorOnAfterUpdate[models.Beat_unit_tied](stage)
	models.SetOrchestratorOnAfterUpdate[models.Beater](stage)
	models.SetOrchestratorOnAfterUpdate[models.Bend](stage)
	models.SetOrchestratorOnAfterUpdate[models.Bookmark](stage)
	models.SetOrchestratorOnAfterUpdate[models.Bracket](stage)
	models.SetOrchestratorOnAfterUpdate[models.Breath_mark](stage)
	models.SetOrchestratorOnAfterUpdate[models.Caesura](stage)
	models.SetOrchestratorOnAfterUpdate[models.Cancel](stage)
	models.SetOrchestratorOnAfterUpdate[models.Clef](stage)
	models.SetOrchestratorOnAfterUpdate[models.Coda](stage)
	models.SetOrchestratorOnAfterUpdate[models.Credit](stage)
	models.SetOrchestratorOnAfterUpdate[models.Dashes](stage)
	models.SetOrchestratorOnAfterUpdate[models.Defaults](stage)
	models.SetOrchestratorOnAfterUpdate[models.Degree](stage)
	models.SetOrchestratorOnAfterUpdate[models.Degree_alter](stage)
	models.SetOrchestratorOnAfterUpdate[models.Degree_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Degree_value](stage)
	models.SetOrchestratorOnAfterUpdate[models.Direction](stage)
	models.SetOrchestratorOnAfterUpdate[models.Direction_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Distance](stage)
	models.SetOrchestratorOnAfterUpdate[models.Double](stage)
	models.SetOrchestratorOnAfterUpdate[models.Dynamics](stage)
	models.SetOrchestratorOnAfterUpdate[models.Effect](stage)
	models.SetOrchestratorOnAfterUpdate[models.Elision](stage)
	models.SetOrchestratorOnAfterUpdate[models.Empty](stage)
	models.SetOrchestratorOnAfterUpdate[models.Empty_font](stage)
	models.SetOrchestratorOnAfterUpdate[models.Empty_line](stage)
	models.SetOrchestratorOnAfterUpdate[models.Empty_placement](stage)
	models.SetOrchestratorOnAfterUpdate[models.Empty_placement_smufl](stage)
	models.SetOrchestratorOnAfterUpdate[models.Empty_print_object_style_align](stage)
	models.SetOrchestratorOnAfterUpdate[models.Empty_print_style](stage)
	models.SetOrchestratorOnAfterUpdate[models.Empty_print_style_align](stage)
	models.SetOrchestratorOnAfterUpdate[models.Empty_print_style_align_id](stage)
	models.SetOrchestratorOnAfterUpdate[models.Empty_trill_sound](stage)
	models.SetOrchestratorOnAfterUpdate[models.Encoding](stage)
	models.SetOrchestratorOnAfterUpdate[models.Ending](stage)
	models.SetOrchestratorOnAfterUpdate[models.Extend](stage)
	models.SetOrchestratorOnAfterUpdate[models.Feature](stage)
	models.SetOrchestratorOnAfterUpdate[models.Fermata](stage)
	models.SetOrchestratorOnAfterUpdate[models.Figure](stage)
	models.SetOrchestratorOnAfterUpdate[models.Figured_bass](stage)
	models.SetOrchestratorOnAfterUpdate[models.Fingering](stage)
	models.SetOrchestratorOnAfterUpdate[models.First_fret](stage)
	models.SetOrchestratorOnAfterUpdate[models.For_part](stage)
	models.SetOrchestratorOnAfterUpdate[models.Formatted_symbol](stage)
	models.SetOrchestratorOnAfterUpdate[models.Formatted_symbol_id](stage)
	models.SetOrchestratorOnAfterUpdate[models.Formatted_text](stage)
	models.SetOrchestratorOnAfterUpdate[models.Formatted_text_id](stage)
	models.SetOrchestratorOnAfterUpdate[models.Forward](stage)
	models.SetOrchestratorOnAfterUpdate[models.Frame](stage)
	models.SetOrchestratorOnAfterUpdate[models.Frame_note](stage)
	models.SetOrchestratorOnAfterUpdate[models.Fret](stage)
	models.SetOrchestratorOnAfterUpdate[models.Glass](stage)
	models.SetOrchestratorOnAfterUpdate[models.Glissando](stage)
	models.SetOrchestratorOnAfterUpdate[models.Glyph](stage)
	models.SetOrchestratorOnAfterUpdate[models.Grace](stage)
	models.SetOrchestratorOnAfterUpdate[models.Group_barline](stage)
	models.SetOrchestratorOnAfterUpdate[models.Group_name](stage)
	models.SetOrchestratorOnAfterUpdate[models.Group_symbol](stage)
	models.SetOrchestratorOnAfterUpdate[models.Grouping](stage)
	models.SetOrchestratorOnAfterUpdate[models.Hammer_on_pull_off](stage)
	models.SetOrchestratorOnAfterUpdate[models.Handbell](stage)
	models.SetOrchestratorOnAfterUpdate[models.Harmon_closed](stage)
	models.SetOrchestratorOnAfterUpdate[models.Harmon_mute](stage)
	models.SetOrchestratorOnAfterUpdate[models.Harmonic](stage)
	models.SetOrchestratorOnAfterUpdate[models.Harmony](stage)
	models.SetOrchestratorOnAfterUpdate[models.Harmony_alter](stage)
	models.SetOrchestratorOnAfterUpdate[models.Harp_pedals](stage)
	models.SetOrchestratorOnAfterUpdate[models.Heel_toe](stage)
	models.SetOrchestratorOnAfterUpdate[models.Hole](stage)
	models.SetOrchestratorOnAfterUpdate[models.Hole_closed](stage)
	models.SetOrchestratorOnAfterUpdate[models.Horizontal_turn](stage)
	models.SetOrchestratorOnAfterUpdate[models.Identification](stage)
	models.SetOrchestratorOnAfterUpdate[models.Image](stage)
	models.SetOrchestratorOnAfterUpdate[models.Instrument](stage)
	models.SetOrchestratorOnAfterUpdate[models.Instrument_change](stage)
	models.SetOrchestratorOnAfterUpdate[models.Instrument_link](stage)
	models.SetOrchestratorOnAfterUpdate[models.Interchangeable](stage)
	models.SetOrchestratorOnAfterUpdate[models.Inversion](stage)
	models.SetOrchestratorOnAfterUpdate[models.Key](stage)
	models.SetOrchestratorOnAfterUpdate[models.Key_accidental](stage)
	models.SetOrchestratorOnAfterUpdate[models.Key_octave](stage)
	models.SetOrchestratorOnAfterUpdate[models.Kind](stage)
	models.SetOrchestratorOnAfterUpdate[models.Level](stage)
	models.SetOrchestratorOnAfterUpdate[models.Line_detail](stage)
	models.SetOrchestratorOnAfterUpdate[models.Line_width](stage)
	models.SetOrchestratorOnAfterUpdate[models.Link](stage)
	models.SetOrchestratorOnAfterUpdate[models.Listen](stage)
	models.SetOrchestratorOnAfterUpdate[models.Listening](stage)
	models.SetOrchestratorOnAfterUpdate[models.Lyric](stage)
	models.SetOrchestratorOnAfterUpdate[models.Lyric_font](stage)
	models.SetOrchestratorOnAfterUpdate[models.Lyric_language](stage)
	models.SetOrchestratorOnAfterUpdate[models.Measure_layout](stage)
	models.SetOrchestratorOnAfterUpdate[models.Measure_numbering](stage)
	models.SetOrchestratorOnAfterUpdate[models.Measure_repeat](stage)
	models.SetOrchestratorOnAfterUpdate[models.Measure_style](stage)
	models.SetOrchestratorOnAfterUpdate[models.Membrane](stage)
	models.SetOrchestratorOnAfterUpdate[models.Metal](stage)
	models.SetOrchestratorOnAfterUpdate[models.Metronome](stage)
	models.SetOrchestratorOnAfterUpdate[models.Metronome_beam](stage)
	models.SetOrchestratorOnAfterUpdate[models.Metronome_note](stage)
	models.SetOrchestratorOnAfterUpdate[models.Metronome_tied](stage)
	models.SetOrchestratorOnAfterUpdate[models.Metronome_tuplet](stage)
	models.SetOrchestratorOnAfterUpdate[models.Midi_device](stage)
	models.SetOrchestratorOnAfterUpdate[models.Midi_instrument](stage)
	models.SetOrchestratorOnAfterUpdate[models.Miscellaneous](stage)
	models.SetOrchestratorOnAfterUpdate[models.Miscellaneous_field](stage)
	models.SetOrchestratorOnAfterUpdate[models.Mordent](stage)
	models.SetOrchestratorOnAfterUpdate[models.Multiple_rest](stage)
	models.SetOrchestratorOnAfterUpdate[models.Name_display](stage)
	models.SetOrchestratorOnAfterUpdate[models.Non_arpeggiate](stage)
	models.SetOrchestratorOnAfterUpdate[models.Notations](stage)
	models.SetOrchestratorOnAfterUpdate[models.Note](stage)
	models.SetOrchestratorOnAfterUpdate[models.Note_size](stage)
	models.SetOrchestratorOnAfterUpdate[models.Note_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Notehead](stage)
	models.SetOrchestratorOnAfterUpdate[models.Notehead_text](stage)
	models.SetOrchestratorOnAfterUpdate[models.Numeral](stage)
	models.SetOrchestratorOnAfterUpdate[models.Numeral_key](stage)
	models.SetOrchestratorOnAfterUpdate[models.Numeral_root](stage)
	models.SetOrchestratorOnAfterUpdate[models.Octave_shift](stage)
	models.SetOrchestratorOnAfterUpdate[models.Offset](stage)
	models.SetOrchestratorOnAfterUpdate[models.Opus](stage)
	models.SetOrchestratorOnAfterUpdate[models.Ornaments](stage)
	models.SetOrchestratorOnAfterUpdate[models.Other_appearance](stage)
	models.SetOrchestratorOnAfterUpdate[models.Other_direction](stage)
	models.SetOrchestratorOnAfterUpdate[models.Other_listening](stage)
	models.SetOrchestratorOnAfterUpdate[models.Other_notation](stage)
	models.SetOrchestratorOnAfterUpdate[models.Other_placement_text](stage)
	models.SetOrchestratorOnAfterUpdate[models.Other_play](stage)
	models.SetOrchestratorOnAfterUpdate[models.Other_text](stage)
	models.SetOrchestratorOnAfterUpdate[models.Page_layout](stage)
	models.SetOrchestratorOnAfterUpdate[models.Page_margins](stage)
	models.SetOrchestratorOnAfterUpdate[models.Part_clef](stage)
	models.SetOrchestratorOnAfterUpdate[models.Part_group](stage)
	models.SetOrchestratorOnAfterUpdate[models.Part_link](stage)
	models.SetOrchestratorOnAfterUpdate[models.Part_list](stage)
	models.SetOrchestratorOnAfterUpdate[models.Part_name](stage)
	models.SetOrchestratorOnAfterUpdate[models.Part_symbol](stage)
	models.SetOrchestratorOnAfterUpdate[models.Part_transpose](stage)
	models.SetOrchestratorOnAfterUpdate[models.Pedal](stage)
	models.SetOrchestratorOnAfterUpdate[models.Pedal_tuning](stage)
	models.SetOrchestratorOnAfterUpdate[models.Per_minute](stage)
	models.SetOrchestratorOnAfterUpdate[models.Percussion](stage)
	models.SetOrchestratorOnAfterUpdate[models.Pitch](stage)
	models.SetOrchestratorOnAfterUpdate[models.Pitched](stage)
	models.SetOrchestratorOnAfterUpdate[models.Placement_text](stage)
	models.SetOrchestratorOnAfterUpdate[models.Play](stage)
	models.SetOrchestratorOnAfterUpdate[models.Player](stage)
	models.SetOrchestratorOnAfterUpdate[models.Principal_voice](stage)
	models.SetOrchestratorOnAfterUpdate[models.Print](stage)
	models.SetOrchestratorOnAfterUpdate[models.Release](stage)
	models.SetOrchestratorOnAfterUpdate[models.Repeat](stage)
	models.SetOrchestratorOnAfterUpdate[models.Rest](stage)
	models.SetOrchestratorOnAfterUpdate[models.Root](stage)
	models.SetOrchestratorOnAfterUpdate[models.Root_step](stage)
	models.SetOrchestratorOnAfterUpdate[models.Scaling](stage)
	models.SetOrchestratorOnAfterUpdate[models.Scordatura](stage)
	models.SetOrchestratorOnAfterUpdate[models.Score_instrument](stage)
	models.SetOrchestratorOnAfterUpdate[models.Score_part](stage)
	models.SetOrchestratorOnAfterUpdate[models.Score_partwise](stage)
	models.SetOrchestratorOnAfterUpdate[models.Score_timewise](stage)
	models.SetOrchestratorOnAfterUpdate[models.Segno](stage)
	models.SetOrchestratorOnAfterUpdate[models.Slash](stage)
	models.SetOrchestratorOnAfterUpdate[models.Slide](stage)
	models.SetOrchestratorOnAfterUpdate[models.Slur](stage)
	models.SetOrchestratorOnAfterUpdate[models.Sound](stage)
	models.SetOrchestratorOnAfterUpdate[models.Staff_details](stage)
	models.SetOrchestratorOnAfterUpdate[models.Staff_divide](stage)
	models.SetOrchestratorOnAfterUpdate[models.Staff_layout](stage)
	models.SetOrchestratorOnAfterUpdate[models.Staff_size](stage)
	models.SetOrchestratorOnAfterUpdate[models.Staff_tuning](stage)
	models.SetOrchestratorOnAfterUpdate[models.Stem](stage)
	models.SetOrchestratorOnAfterUpdate[models.Stick](stage)
	models.SetOrchestratorOnAfterUpdate[models.String_mute](stage)
	models.SetOrchestratorOnAfterUpdate[models.String_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Strong_accent](stage)
	models.SetOrchestratorOnAfterUpdate[models.Style_text](stage)
	models.SetOrchestratorOnAfterUpdate[models.Supports](stage)
	models.SetOrchestratorOnAfterUpdate[models.Swing](stage)
	models.SetOrchestratorOnAfterUpdate[models.Sync](stage)
	models.SetOrchestratorOnAfterUpdate[models.System_dividers](stage)
	models.SetOrchestratorOnAfterUpdate[models.System_layout](stage)
	models.SetOrchestratorOnAfterUpdate[models.System_margins](stage)
	models.SetOrchestratorOnAfterUpdate[models.Tap](stage)
	models.SetOrchestratorOnAfterUpdate[models.Technical](stage)
	models.SetOrchestratorOnAfterUpdate[models.Text_element_data](stage)
	models.SetOrchestratorOnAfterUpdate[models.Tie](stage)
	models.SetOrchestratorOnAfterUpdate[models.Tied](stage)
	models.SetOrchestratorOnAfterUpdate[models.Time](stage)
	models.SetOrchestratorOnAfterUpdate[models.Time_modification](stage)
	models.SetOrchestratorOnAfterUpdate[models.Timpani](stage)
	models.SetOrchestratorOnAfterUpdate[models.Transpose](stage)
	models.SetOrchestratorOnAfterUpdate[models.Tremolo](stage)
	models.SetOrchestratorOnAfterUpdate[models.Tuplet](stage)
	models.SetOrchestratorOnAfterUpdate[models.Tuplet_dot](stage)
	models.SetOrchestratorOnAfterUpdate[models.Tuplet_number](stage)
	models.SetOrchestratorOnAfterUpdate[models.Tuplet_portion](stage)
	models.SetOrchestratorOnAfterUpdate[models.Tuplet_type](stage)
	models.SetOrchestratorOnAfterUpdate[models.Typed_text](stage)
	models.SetOrchestratorOnAfterUpdate[models.Unpitched](stage)
	models.SetOrchestratorOnAfterUpdate[models.Virtual_instrument](stage)
	models.SetOrchestratorOnAfterUpdate[models.Wait](stage)
	models.SetOrchestratorOnAfterUpdate[models.Wavy_line](stage)
	models.SetOrchestratorOnAfterUpdate[models.Wedge](stage)
	models.SetOrchestratorOnAfterUpdate[models.Wood](stage)
	models.SetOrchestratorOnAfterUpdate[models.Work](stage)

	return
}
