// do not modify, generated file
package level1stack

import (
	"fmt"
	"log"
	"strings"

	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models"
	"github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go/models/probe"

	embeddedgo "github.com/fullstack-lang/gong/app/xsd/tests/musicxml/go"

	"net/http"

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
	R     *http.ServeMux
}

func (stack *Level1Stack) Run(addr string) error {
	return split_static.RunServer(stack.R, addr)
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
		// "go/models/diagrams/diagrams.go", the path is "../../models/diagrams/diagrams.go"
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
		err := stage.ParseAstFile(unmarshallFromCode, true)

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
	stage.SetOrchestratorOnAfterUpdate[models.A_directive]()
	stage.SetOrchestratorOnAfterUpdate[models.A_measure]()
	stage.SetOrchestratorOnAfterUpdate[models.A_measure_1]()
	stage.SetOrchestratorOnAfterUpdate[models.A_part]()
	stage.SetOrchestratorOnAfterUpdate[models.A_part_1]()
	stage.SetOrchestratorOnAfterUpdate[models.Accidental]()
	stage.SetOrchestratorOnAfterUpdate[models.Accidental_mark]()
	stage.SetOrchestratorOnAfterUpdate[models.Accidental_text]()
	stage.SetOrchestratorOnAfterUpdate[models.Accord]()
	stage.SetOrchestratorOnAfterUpdate[models.Accordion_registration]()
	stage.SetOrchestratorOnAfterUpdate[models.Appearance]()
	stage.SetOrchestratorOnAfterUpdate[models.Arpeggiate]()
	stage.SetOrchestratorOnAfterUpdate[models.Arrow]()
	stage.SetOrchestratorOnAfterUpdate[models.Articulations]()
	stage.SetOrchestratorOnAfterUpdate[models.Assess]()
	stage.SetOrchestratorOnAfterUpdate[models.Attributes]()
	stage.SetOrchestratorOnAfterUpdate[models.Backup]()
	stage.SetOrchestratorOnAfterUpdate[models.Bar_style_color]()
	stage.SetOrchestratorOnAfterUpdate[models.Barline]()
	stage.SetOrchestratorOnAfterUpdate[models.Barre]()
	stage.SetOrchestratorOnAfterUpdate[models.Bass]()
	stage.SetOrchestratorOnAfterUpdate[models.Bass_step]()
	stage.SetOrchestratorOnAfterUpdate[models.Beam]()
	stage.SetOrchestratorOnAfterUpdate[models.Beat_repeat]()
	stage.SetOrchestratorOnAfterUpdate[models.Beat_unit_tied]()
	stage.SetOrchestratorOnAfterUpdate[models.Beater]()
	stage.SetOrchestratorOnAfterUpdate[models.Bend]()
	stage.SetOrchestratorOnAfterUpdate[models.Bookmark]()
	stage.SetOrchestratorOnAfterUpdate[models.Bracket]()
	stage.SetOrchestratorOnAfterUpdate[models.Breath_mark]()
	stage.SetOrchestratorOnAfterUpdate[models.Caesura]()
	stage.SetOrchestratorOnAfterUpdate[models.Cancel]()
	stage.SetOrchestratorOnAfterUpdate[models.Clef]()
	stage.SetOrchestratorOnAfterUpdate[models.Coda]()
	stage.SetOrchestratorOnAfterUpdate[models.Credit]()
	stage.SetOrchestratorOnAfterUpdate[models.Dashes]()
	stage.SetOrchestratorOnAfterUpdate[models.Defaults]()
	stage.SetOrchestratorOnAfterUpdate[models.Degree]()
	stage.SetOrchestratorOnAfterUpdate[models.Degree_alter]()
	stage.SetOrchestratorOnAfterUpdate[models.Degree_type]()
	stage.SetOrchestratorOnAfterUpdate[models.Degree_value]()
	stage.SetOrchestratorOnAfterUpdate[models.Direction]()
	stage.SetOrchestratorOnAfterUpdate[models.Direction_type]()
	stage.SetOrchestratorOnAfterUpdate[models.Distance]()
	stage.SetOrchestratorOnAfterUpdate[models.Double]()
	stage.SetOrchestratorOnAfterUpdate[models.Dynamics]()
	stage.SetOrchestratorOnAfterUpdate[models.Effect]()
	stage.SetOrchestratorOnAfterUpdate[models.Elision]()
	stage.SetOrchestratorOnAfterUpdate[models.Empty]()
	stage.SetOrchestratorOnAfterUpdate[models.Empty_font]()
	stage.SetOrchestratorOnAfterUpdate[models.Empty_line]()
	stage.SetOrchestratorOnAfterUpdate[models.Empty_placement]()
	stage.SetOrchestratorOnAfterUpdate[models.Empty_placement_smufl]()
	stage.SetOrchestratorOnAfterUpdate[models.Empty_print_object_style_align]()
	stage.SetOrchestratorOnAfterUpdate[models.Empty_print_style]()
	stage.SetOrchestratorOnAfterUpdate[models.Empty_print_style_align]()
	stage.SetOrchestratorOnAfterUpdate[models.Empty_print_style_align_id]()
	stage.SetOrchestratorOnAfterUpdate[models.Empty_trill_sound]()
	stage.SetOrchestratorOnAfterUpdate[models.Encoding]()
	stage.SetOrchestratorOnAfterUpdate[models.Ending]()
	stage.SetOrchestratorOnAfterUpdate[models.Extend]()
	stage.SetOrchestratorOnAfterUpdate[models.Feature]()
	stage.SetOrchestratorOnAfterUpdate[models.Fermata]()
	stage.SetOrchestratorOnAfterUpdate[models.Figure]()
	stage.SetOrchestratorOnAfterUpdate[models.Figured_bass]()
	stage.SetOrchestratorOnAfterUpdate[models.Fingering]()
	stage.SetOrchestratorOnAfterUpdate[models.First_fret]()
	stage.SetOrchestratorOnAfterUpdate[models.For_part]()
	stage.SetOrchestratorOnAfterUpdate[models.Formatted_symbol]()
	stage.SetOrchestratorOnAfterUpdate[models.Formatted_symbol_id]()
	stage.SetOrchestratorOnAfterUpdate[models.Formatted_text]()
	stage.SetOrchestratorOnAfterUpdate[models.Formatted_text_id]()
	stage.SetOrchestratorOnAfterUpdate[models.Forward]()
	stage.SetOrchestratorOnAfterUpdate[models.Frame]()
	stage.SetOrchestratorOnAfterUpdate[models.Frame_note]()
	stage.SetOrchestratorOnAfterUpdate[models.Fret]()
	stage.SetOrchestratorOnAfterUpdate[models.Glass]()
	stage.SetOrchestratorOnAfterUpdate[models.Glissando]()
	stage.SetOrchestratorOnAfterUpdate[models.Glyph]()
	stage.SetOrchestratorOnAfterUpdate[models.Grace]()
	stage.SetOrchestratorOnAfterUpdate[models.Group_barline]()
	stage.SetOrchestratorOnAfterUpdate[models.Group_name]()
	stage.SetOrchestratorOnAfterUpdate[models.Group_symbol]()
	stage.SetOrchestratorOnAfterUpdate[models.Grouping]()
	stage.SetOrchestratorOnAfterUpdate[models.Hammer_on_pull_off]()
	stage.SetOrchestratorOnAfterUpdate[models.Handbell]()
	stage.SetOrchestratorOnAfterUpdate[models.Harmon_closed]()
	stage.SetOrchestratorOnAfterUpdate[models.Harmon_mute]()
	stage.SetOrchestratorOnAfterUpdate[models.Harmonic]()
	stage.SetOrchestratorOnAfterUpdate[models.Harmony]()
	stage.SetOrchestratorOnAfterUpdate[models.Harmony_alter]()
	stage.SetOrchestratorOnAfterUpdate[models.Harp_pedals]()
	stage.SetOrchestratorOnAfterUpdate[models.Heel_toe]()
	stage.SetOrchestratorOnAfterUpdate[models.Hole]()
	stage.SetOrchestratorOnAfterUpdate[models.Hole_closed]()
	stage.SetOrchestratorOnAfterUpdate[models.Horizontal_turn]()
	stage.SetOrchestratorOnAfterUpdate[models.Identification]()
	stage.SetOrchestratorOnAfterUpdate[models.Image]()
	stage.SetOrchestratorOnAfterUpdate[models.Instrument]()
	stage.SetOrchestratorOnAfterUpdate[models.Instrument_change]()
	stage.SetOrchestratorOnAfterUpdate[models.Instrument_link]()
	stage.SetOrchestratorOnAfterUpdate[models.Interchangeable]()
	stage.SetOrchestratorOnAfterUpdate[models.Inversion]()
	stage.SetOrchestratorOnAfterUpdate[models.Key]()
	stage.SetOrchestratorOnAfterUpdate[models.Key_accidental]()
	stage.SetOrchestratorOnAfterUpdate[models.Key_octave]()
	stage.SetOrchestratorOnAfterUpdate[models.Kind]()
	stage.SetOrchestratorOnAfterUpdate[models.Level]()
	stage.SetOrchestratorOnAfterUpdate[models.Line_detail]()
	stage.SetOrchestratorOnAfterUpdate[models.Line_width]()
	stage.SetOrchestratorOnAfterUpdate[models.Link]()
	stage.SetOrchestratorOnAfterUpdate[models.Listen]()
	stage.SetOrchestratorOnAfterUpdate[models.Listening]()
	stage.SetOrchestratorOnAfterUpdate[models.Lyric]()
	stage.SetOrchestratorOnAfterUpdate[models.Lyric_font]()
	stage.SetOrchestratorOnAfterUpdate[models.Lyric_language]()
	stage.SetOrchestratorOnAfterUpdate[models.Measure_layout]()
	stage.SetOrchestratorOnAfterUpdate[models.Measure_numbering]()
	stage.SetOrchestratorOnAfterUpdate[models.Measure_repeat]()
	stage.SetOrchestratorOnAfterUpdate[models.Measure_style]()
	stage.SetOrchestratorOnAfterUpdate[models.Membrane]()
	stage.SetOrchestratorOnAfterUpdate[models.Metal]()
	stage.SetOrchestratorOnAfterUpdate[models.Metronome]()
	stage.SetOrchestratorOnAfterUpdate[models.Metronome_beam]()
	stage.SetOrchestratorOnAfterUpdate[models.Metronome_note]()
	stage.SetOrchestratorOnAfterUpdate[models.Metronome_tied]()
	stage.SetOrchestratorOnAfterUpdate[models.Metronome_tuplet]()
	stage.SetOrchestratorOnAfterUpdate[models.Midi_device]()
	stage.SetOrchestratorOnAfterUpdate[models.Midi_instrument]()
	stage.SetOrchestratorOnAfterUpdate[models.Miscellaneous]()
	stage.SetOrchestratorOnAfterUpdate[models.Miscellaneous_field]()
	stage.SetOrchestratorOnAfterUpdate[models.Mordent]()
	stage.SetOrchestratorOnAfterUpdate[models.Multiple_rest]()
	stage.SetOrchestratorOnAfterUpdate[models.Name_display]()
	stage.SetOrchestratorOnAfterUpdate[models.Non_arpeggiate]()
	stage.SetOrchestratorOnAfterUpdate[models.Notations]()
	stage.SetOrchestratorOnAfterUpdate[models.Note]()
	stage.SetOrchestratorOnAfterUpdate[models.Note_size]()
	stage.SetOrchestratorOnAfterUpdate[models.Note_type]()
	stage.SetOrchestratorOnAfterUpdate[models.Notehead]()
	stage.SetOrchestratorOnAfterUpdate[models.Notehead_text]()
	stage.SetOrchestratorOnAfterUpdate[models.Numeral]()
	stage.SetOrchestratorOnAfterUpdate[models.Numeral_key]()
	stage.SetOrchestratorOnAfterUpdate[models.Numeral_root]()
	stage.SetOrchestratorOnAfterUpdate[models.Octave_shift]()
	stage.SetOrchestratorOnAfterUpdate[models.Offset]()
	stage.SetOrchestratorOnAfterUpdate[models.Opus]()
	stage.SetOrchestratorOnAfterUpdate[models.Ornaments]()
	stage.SetOrchestratorOnAfterUpdate[models.Other_appearance]()
	stage.SetOrchestratorOnAfterUpdate[models.Other_direction]()
	stage.SetOrchestratorOnAfterUpdate[models.Other_listening]()
	stage.SetOrchestratorOnAfterUpdate[models.Other_notation]()
	stage.SetOrchestratorOnAfterUpdate[models.Other_placement_text]()
	stage.SetOrchestratorOnAfterUpdate[models.Other_play]()
	stage.SetOrchestratorOnAfterUpdate[models.Other_text]()
	stage.SetOrchestratorOnAfterUpdate[models.Page_layout]()
	stage.SetOrchestratorOnAfterUpdate[models.Page_margins]()
	stage.SetOrchestratorOnAfterUpdate[models.Part_clef]()
	stage.SetOrchestratorOnAfterUpdate[models.Part_group]()
	stage.SetOrchestratorOnAfterUpdate[models.Part_link]()
	stage.SetOrchestratorOnAfterUpdate[models.Part_list]()
	stage.SetOrchestratorOnAfterUpdate[models.Part_name]()
	stage.SetOrchestratorOnAfterUpdate[models.Part_symbol]()
	stage.SetOrchestratorOnAfterUpdate[models.Part_transpose]()
	stage.SetOrchestratorOnAfterUpdate[models.Pedal]()
	stage.SetOrchestratorOnAfterUpdate[models.Pedal_tuning]()
	stage.SetOrchestratorOnAfterUpdate[models.Per_minute]()
	stage.SetOrchestratorOnAfterUpdate[models.Percussion]()
	stage.SetOrchestratorOnAfterUpdate[models.Pitch]()
	stage.SetOrchestratorOnAfterUpdate[models.Pitched]()
	stage.SetOrchestratorOnAfterUpdate[models.Placement_text]()
	stage.SetOrchestratorOnAfterUpdate[models.Play]()
	stage.SetOrchestratorOnAfterUpdate[models.Player]()
	stage.SetOrchestratorOnAfterUpdate[models.Principal_voice]()
	stage.SetOrchestratorOnAfterUpdate[models.Print]()
	stage.SetOrchestratorOnAfterUpdate[models.Release]()
	stage.SetOrchestratorOnAfterUpdate[models.Repeat]()
	stage.SetOrchestratorOnAfterUpdate[models.Rest]()
	stage.SetOrchestratorOnAfterUpdate[models.Root]()
	stage.SetOrchestratorOnAfterUpdate[models.Root_step]()
	stage.SetOrchestratorOnAfterUpdate[models.Scaling]()
	stage.SetOrchestratorOnAfterUpdate[models.Scordatura]()
	stage.SetOrchestratorOnAfterUpdate[models.Score_instrument]()
	stage.SetOrchestratorOnAfterUpdate[models.Score_part]()
	stage.SetOrchestratorOnAfterUpdate[models.Score_partwise]()
	stage.SetOrchestratorOnAfterUpdate[models.Score_timewise]()
	stage.SetOrchestratorOnAfterUpdate[models.Segno]()
	stage.SetOrchestratorOnAfterUpdate[models.Slash]()
	stage.SetOrchestratorOnAfterUpdate[models.Slide]()
	stage.SetOrchestratorOnAfterUpdate[models.Slur]()
	stage.SetOrchestratorOnAfterUpdate[models.Sound]()
	stage.SetOrchestratorOnAfterUpdate[models.Staff_details]()
	stage.SetOrchestratorOnAfterUpdate[models.Staff_divide]()
	stage.SetOrchestratorOnAfterUpdate[models.Staff_layout]()
	stage.SetOrchestratorOnAfterUpdate[models.Staff_size]()
	stage.SetOrchestratorOnAfterUpdate[models.Staff_tuning]()
	stage.SetOrchestratorOnAfterUpdate[models.Stem]()
	stage.SetOrchestratorOnAfterUpdate[models.Stick]()
	stage.SetOrchestratorOnAfterUpdate[models.String_mute]()
	stage.SetOrchestratorOnAfterUpdate[models.String_type]()
	stage.SetOrchestratorOnAfterUpdate[models.Strong_accent]()
	stage.SetOrchestratorOnAfterUpdate[models.Style_text]()
	stage.SetOrchestratorOnAfterUpdate[models.Supports]()
	stage.SetOrchestratorOnAfterUpdate[models.Swing]()
	stage.SetOrchestratorOnAfterUpdate[models.Sync]()
	stage.SetOrchestratorOnAfterUpdate[models.System_dividers]()
	stage.SetOrchestratorOnAfterUpdate[models.System_layout]()
	stage.SetOrchestratorOnAfterUpdate[models.System_margins]()
	stage.SetOrchestratorOnAfterUpdate[models.Tap]()
	stage.SetOrchestratorOnAfterUpdate[models.Technical]()
	stage.SetOrchestratorOnAfterUpdate[models.Text_element_data]()
	stage.SetOrchestratorOnAfterUpdate[models.Tie]()
	stage.SetOrchestratorOnAfterUpdate[models.Tied]()
	stage.SetOrchestratorOnAfterUpdate[models.Time]()
	stage.SetOrchestratorOnAfterUpdate[models.Time_modification]()
	stage.SetOrchestratorOnAfterUpdate[models.Timpani]()
	stage.SetOrchestratorOnAfterUpdate[models.Transpose]()
	stage.SetOrchestratorOnAfterUpdate[models.Tremolo]()
	stage.SetOrchestratorOnAfterUpdate[models.Tuplet]()
	stage.SetOrchestratorOnAfterUpdate[models.Tuplet_dot]()
	stage.SetOrchestratorOnAfterUpdate[models.Tuplet_number]()
	stage.SetOrchestratorOnAfterUpdate[models.Tuplet_portion]()
	stage.SetOrchestratorOnAfterUpdate[models.Tuplet_type]()
	stage.SetOrchestratorOnAfterUpdate[models.Typed_text]()
	stage.SetOrchestratorOnAfterUpdate[models.Unpitched]()
	stage.SetOrchestratorOnAfterUpdate[models.Virtual_instrument]()
	stage.SetOrchestratorOnAfterUpdate[models.Wait]()
	stage.SetOrchestratorOnAfterUpdate[models.Wavy_line]()
	stage.SetOrchestratorOnAfterUpdate[models.Wedge]()
	stage.SetOrchestratorOnAfterUpdate[models.Wood]()
	stage.SetOrchestratorOnAfterUpdate[models.Work]()

	return
}
