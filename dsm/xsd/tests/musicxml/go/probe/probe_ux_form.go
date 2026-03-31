// generated code - do not edit
package probe

import (
	form "github.com/fullstack-lang/gong/lib/table/go/models"

	"github.com/fullstack-lang/gong/dsm/xsd/tests/musicxml/go/models"
)

// ux_form updates the current form if there is one
func (probe *Probe) ux_form() {
	var formGroup *form.FormGroup
	for fg := range probe.formStage.FormGroups {
		formGroup = fg
	}
	if formGroup != nil {
		switch onSave := formGroup.OnSave.(type) { // insertion point
		case *A_directiveFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "A_directive", true)
			} else {
				FillUpFormFromGongstruct(onSave.a_directive, probe)
			}
		case *A_measureFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "A_measure", true)
			} else {
				FillUpFormFromGongstruct(onSave.a_measure, probe)
			}
		case *A_measure_1FormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "A_measure_1", true)
			} else {
				FillUpFormFromGongstruct(onSave.a_measure_1, probe)
			}
		case *A_partFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "A_part", true)
			} else {
				FillUpFormFromGongstruct(onSave.a_part, probe)
			}
		case *A_part_1FormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "A_part_1", true)
			} else {
				FillUpFormFromGongstruct(onSave.a_part_1, probe)
			}
		case *AccidentalFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Accidental", true)
			} else {
				FillUpFormFromGongstruct(onSave.accidental, probe)
			}
		case *Accidental_markFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Accidental_mark", true)
			} else {
				FillUpFormFromGongstruct(onSave.accidental_mark, probe)
			}
		case *Accidental_textFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Accidental_text", true)
			} else {
				FillUpFormFromGongstruct(onSave.accidental_text, probe)
			}
		case *AccordFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Accord", true)
			} else {
				FillUpFormFromGongstruct(onSave.accord, probe)
			}
		case *Accordion_registrationFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Accordion_registration", true)
			} else {
				FillUpFormFromGongstruct(onSave.accordion_registration, probe)
			}
		case *AppearanceFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Appearance", true)
			} else {
				FillUpFormFromGongstruct(onSave.appearance, probe)
			}
		case *ArpeggiateFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Arpeggiate", true)
			} else {
				FillUpFormFromGongstruct(onSave.arpeggiate, probe)
			}
		case *ArrowFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Arrow", true)
			} else {
				FillUpFormFromGongstruct(onSave.arrow, probe)
			}
		case *ArticulationsFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Articulations", true)
			} else {
				FillUpFormFromGongstruct(onSave.articulations, probe)
			}
		case *AssessFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Assess", true)
			} else {
				FillUpFormFromGongstruct(onSave.assess, probe)
			}
		case *AttributesFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Attributes", true)
			} else {
				FillUpFormFromGongstruct(onSave.attributes, probe)
			}
		case *BackupFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Backup", true)
			} else {
				FillUpFormFromGongstruct(onSave.backup, probe)
			}
		case *Bar_style_colorFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Bar_style_color", true)
			} else {
				FillUpFormFromGongstruct(onSave.bar_style_color, probe)
			}
		case *BarlineFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Barline", true)
			} else {
				FillUpFormFromGongstruct(onSave.barline, probe)
			}
		case *BarreFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Barre", true)
			} else {
				FillUpFormFromGongstruct(onSave.barre, probe)
			}
		case *BassFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Bass", true)
			} else {
				FillUpFormFromGongstruct(onSave.bass, probe)
			}
		case *Bass_stepFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Bass_step", true)
			} else {
				FillUpFormFromGongstruct(onSave.bass_step, probe)
			}
		case *BeamFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Beam", true)
			} else {
				FillUpFormFromGongstruct(onSave.beam, probe)
			}
		case *Beat_repeatFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Beat_repeat", true)
			} else {
				FillUpFormFromGongstruct(onSave.beat_repeat, probe)
			}
		case *Beat_unit_tiedFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Beat_unit_tied", true)
			} else {
				FillUpFormFromGongstruct(onSave.beat_unit_tied, probe)
			}
		case *BeaterFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Beater", true)
			} else {
				FillUpFormFromGongstruct(onSave.beater, probe)
			}
		case *BendFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Bend", true)
			} else {
				FillUpFormFromGongstruct(onSave.bend, probe)
			}
		case *BookmarkFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Bookmark", true)
			} else {
				FillUpFormFromGongstruct(onSave.bookmark, probe)
			}
		case *BracketFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Bracket", true)
			} else {
				FillUpFormFromGongstruct(onSave.bracket, probe)
			}
		case *Breath_markFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Breath_mark", true)
			} else {
				FillUpFormFromGongstruct(onSave.breath_mark, probe)
			}
		case *CaesuraFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Caesura", true)
			} else {
				FillUpFormFromGongstruct(onSave.caesura, probe)
			}
		case *CancelFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Cancel", true)
			} else {
				FillUpFormFromGongstruct(onSave.cancel, probe)
			}
		case *ClefFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Clef", true)
			} else {
				FillUpFormFromGongstruct(onSave.clef, probe)
			}
		case *CodaFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Coda", true)
			} else {
				FillUpFormFromGongstruct(onSave.coda, probe)
			}
		case *CreditFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Credit", true)
			} else {
				FillUpFormFromGongstruct(onSave.credit, probe)
			}
		case *DashesFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Dashes", true)
			} else {
				FillUpFormFromGongstruct(onSave.dashes, probe)
			}
		case *DefaultsFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Defaults", true)
			} else {
				FillUpFormFromGongstruct(onSave.defaults, probe)
			}
		case *DegreeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Degree", true)
			} else {
				FillUpFormFromGongstruct(onSave.degree, probe)
			}
		case *Degree_alterFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Degree_alter", true)
			} else {
				FillUpFormFromGongstruct(onSave.degree_alter, probe)
			}
		case *Degree_typeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Degree_type", true)
			} else {
				FillUpFormFromGongstruct(onSave.degree_type, probe)
			}
		case *Degree_valueFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Degree_value", true)
			} else {
				FillUpFormFromGongstruct(onSave.degree_value, probe)
			}
		case *DirectionFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Direction", true)
			} else {
				FillUpFormFromGongstruct(onSave.direction, probe)
			}
		case *Direction_typeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Direction_type", true)
			} else {
				FillUpFormFromGongstruct(onSave.direction_type, probe)
			}
		case *DistanceFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Distance", true)
			} else {
				FillUpFormFromGongstruct(onSave.distance, probe)
			}
		case *DoubleFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Double", true)
			} else {
				FillUpFormFromGongstruct(onSave.double, probe)
			}
		case *DynamicsFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Dynamics", true)
			} else {
				FillUpFormFromGongstruct(onSave.dynamics, probe)
			}
		case *EffectFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Effect", true)
			} else {
				FillUpFormFromGongstruct(onSave.effect, probe)
			}
		case *ElisionFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Elision", true)
			} else {
				FillUpFormFromGongstruct(onSave.elision, probe)
			}
		case *EmptyFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Empty", true)
			} else {
				FillUpFormFromGongstruct(onSave.empty, probe)
			}
		case *Empty_fontFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Empty_font", true)
			} else {
				FillUpFormFromGongstruct(onSave.empty_font, probe)
			}
		case *Empty_lineFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Empty_line", true)
			} else {
				FillUpFormFromGongstruct(onSave.empty_line, probe)
			}
		case *Empty_placementFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Empty_placement", true)
			} else {
				FillUpFormFromGongstruct(onSave.empty_placement, probe)
			}
		case *Empty_placement_smuflFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Empty_placement_smufl", true)
			} else {
				FillUpFormFromGongstruct(onSave.empty_placement_smufl, probe)
			}
		case *Empty_print_object_style_alignFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Empty_print_object_style_align", true)
			} else {
				FillUpFormFromGongstruct(onSave.empty_print_object_style_align, probe)
			}
		case *Empty_print_styleFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Empty_print_style", true)
			} else {
				FillUpFormFromGongstruct(onSave.empty_print_style, probe)
			}
		case *Empty_print_style_alignFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Empty_print_style_align", true)
			} else {
				FillUpFormFromGongstruct(onSave.empty_print_style_align, probe)
			}
		case *Empty_print_style_align_idFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Empty_print_style_align_id", true)
			} else {
				FillUpFormFromGongstruct(onSave.empty_print_style_align_id, probe)
			}
		case *Empty_trill_soundFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Empty_trill_sound", true)
			} else {
				FillUpFormFromGongstruct(onSave.empty_trill_sound, probe)
			}
		case *EncodingFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Encoding", true)
			} else {
				FillUpFormFromGongstruct(onSave.encoding, probe)
			}
		case *EndingFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Ending", true)
			} else {
				FillUpFormFromGongstruct(onSave.ending, probe)
			}
		case *ExtendFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Extend", true)
			} else {
				FillUpFormFromGongstruct(onSave.extend, probe)
			}
		case *FeatureFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Feature", true)
			} else {
				FillUpFormFromGongstruct(onSave.feature, probe)
			}
		case *FermataFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Fermata", true)
			} else {
				FillUpFormFromGongstruct(onSave.fermata, probe)
			}
		case *FigureFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Figure", true)
			} else {
				FillUpFormFromGongstruct(onSave.figure, probe)
			}
		case *Figured_bassFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Figured_bass", true)
			} else {
				FillUpFormFromGongstruct(onSave.figured_bass, probe)
			}
		case *FingeringFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Fingering", true)
			} else {
				FillUpFormFromGongstruct(onSave.fingering, probe)
			}
		case *First_fretFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "First_fret", true)
			} else {
				FillUpFormFromGongstruct(onSave.first_fret, probe)
			}
		case *For_partFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "For_part", true)
			} else {
				FillUpFormFromGongstruct(onSave.for_part, probe)
			}
		case *Formatted_symbolFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Formatted_symbol", true)
			} else {
				FillUpFormFromGongstruct(onSave.formatted_symbol, probe)
			}
		case *Formatted_symbol_idFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Formatted_symbol_id", true)
			} else {
				FillUpFormFromGongstruct(onSave.formatted_symbol_id, probe)
			}
		case *Formatted_textFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Formatted_text", true)
			} else {
				FillUpFormFromGongstruct(onSave.formatted_text, probe)
			}
		case *Formatted_text_idFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Formatted_text_id", true)
			} else {
				FillUpFormFromGongstruct(onSave.formatted_text_id, probe)
			}
		case *ForwardFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Forward", true)
			} else {
				FillUpFormFromGongstruct(onSave.forward, probe)
			}
		case *FrameFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Frame", true)
			} else {
				FillUpFormFromGongstruct(onSave.frame, probe)
			}
		case *Frame_noteFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Frame_note", true)
			} else {
				FillUpFormFromGongstruct(onSave.frame_note, probe)
			}
		case *FretFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Fret", true)
			} else {
				FillUpFormFromGongstruct(onSave.fret, probe)
			}
		case *GlassFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Glass", true)
			} else {
				FillUpFormFromGongstruct(onSave.glass, probe)
			}
		case *GlissandoFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Glissando", true)
			} else {
				FillUpFormFromGongstruct(onSave.glissando, probe)
			}
		case *GlyphFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Glyph", true)
			} else {
				FillUpFormFromGongstruct(onSave.glyph, probe)
			}
		case *GraceFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Grace", true)
			} else {
				FillUpFormFromGongstruct(onSave.grace, probe)
			}
		case *Group_barlineFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Group_barline", true)
			} else {
				FillUpFormFromGongstruct(onSave.group_barline, probe)
			}
		case *Group_nameFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Group_name", true)
			} else {
				FillUpFormFromGongstruct(onSave.group_name, probe)
			}
		case *Group_symbolFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Group_symbol", true)
			} else {
				FillUpFormFromGongstruct(onSave.group_symbol, probe)
			}
		case *GroupingFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Grouping", true)
			} else {
				FillUpFormFromGongstruct(onSave.grouping, probe)
			}
		case *Hammer_on_pull_offFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Hammer_on_pull_off", true)
			} else {
				FillUpFormFromGongstruct(onSave.hammer_on_pull_off, probe)
			}
		case *HandbellFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Handbell", true)
			} else {
				FillUpFormFromGongstruct(onSave.handbell, probe)
			}
		case *Harmon_closedFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Harmon_closed", true)
			} else {
				FillUpFormFromGongstruct(onSave.harmon_closed, probe)
			}
		case *Harmon_muteFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Harmon_mute", true)
			} else {
				FillUpFormFromGongstruct(onSave.harmon_mute, probe)
			}
		case *HarmonicFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Harmonic", true)
			} else {
				FillUpFormFromGongstruct(onSave.harmonic, probe)
			}
		case *HarmonyFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Harmony", true)
			} else {
				FillUpFormFromGongstruct(onSave.harmony, probe)
			}
		case *Harmony_alterFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Harmony_alter", true)
			} else {
				FillUpFormFromGongstruct(onSave.harmony_alter, probe)
			}
		case *Harp_pedalsFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Harp_pedals", true)
			} else {
				FillUpFormFromGongstruct(onSave.harp_pedals, probe)
			}
		case *Heel_toeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Heel_toe", true)
			} else {
				FillUpFormFromGongstruct(onSave.heel_toe, probe)
			}
		case *HoleFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Hole", true)
			} else {
				FillUpFormFromGongstruct(onSave.hole, probe)
			}
		case *Hole_closedFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Hole_closed", true)
			} else {
				FillUpFormFromGongstruct(onSave.hole_closed, probe)
			}
		case *Horizontal_turnFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Horizontal_turn", true)
			} else {
				FillUpFormFromGongstruct(onSave.horizontal_turn, probe)
			}
		case *IdentificationFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Identification", true)
			} else {
				FillUpFormFromGongstruct(onSave.identification, probe)
			}
		case *ImageFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Image", true)
			} else {
				FillUpFormFromGongstruct(onSave.image, probe)
			}
		case *InstrumentFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Instrument", true)
			} else {
				FillUpFormFromGongstruct(onSave.instrument, probe)
			}
		case *Instrument_changeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Instrument_change", true)
			} else {
				FillUpFormFromGongstruct(onSave.instrument_change, probe)
			}
		case *Instrument_linkFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Instrument_link", true)
			} else {
				FillUpFormFromGongstruct(onSave.instrument_link, probe)
			}
		case *InterchangeableFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Interchangeable", true)
			} else {
				FillUpFormFromGongstruct(onSave.interchangeable, probe)
			}
		case *InversionFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Inversion", true)
			} else {
				FillUpFormFromGongstruct(onSave.inversion, probe)
			}
		case *KeyFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Key", true)
			} else {
				FillUpFormFromGongstruct(onSave.key, probe)
			}
		case *Key_accidentalFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Key_accidental", true)
			} else {
				FillUpFormFromGongstruct(onSave.key_accidental, probe)
			}
		case *Key_octaveFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Key_octave", true)
			} else {
				FillUpFormFromGongstruct(onSave.key_octave, probe)
			}
		case *KindFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Kind", true)
			} else {
				FillUpFormFromGongstruct(onSave.kind, probe)
			}
		case *LevelFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Level", true)
			} else {
				FillUpFormFromGongstruct(onSave.level, probe)
			}
		case *Line_detailFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Line_detail", true)
			} else {
				FillUpFormFromGongstruct(onSave.line_detail, probe)
			}
		case *Line_widthFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Line_width", true)
			} else {
				FillUpFormFromGongstruct(onSave.line_width, probe)
			}
		case *LinkFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Link", true)
			} else {
				FillUpFormFromGongstruct(onSave.link, probe)
			}
		case *ListenFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Listen", true)
			} else {
				FillUpFormFromGongstruct(onSave.listen, probe)
			}
		case *ListeningFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Listening", true)
			} else {
				FillUpFormFromGongstruct(onSave.listening, probe)
			}
		case *LyricFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Lyric", true)
			} else {
				FillUpFormFromGongstruct(onSave.lyric, probe)
			}
		case *Lyric_fontFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Lyric_font", true)
			} else {
				FillUpFormFromGongstruct(onSave.lyric_font, probe)
			}
		case *Lyric_languageFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Lyric_language", true)
			} else {
				FillUpFormFromGongstruct(onSave.lyric_language, probe)
			}
		case *Measure_layoutFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Measure_layout", true)
			} else {
				FillUpFormFromGongstruct(onSave.measure_layout, probe)
			}
		case *Measure_numberingFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Measure_numbering", true)
			} else {
				FillUpFormFromGongstruct(onSave.measure_numbering, probe)
			}
		case *Measure_repeatFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Measure_repeat", true)
			} else {
				FillUpFormFromGongstruct(onSave.measure_repeat, probe)
			}
		case *Measure_styleFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Measure_style", true)
			} else {
				FillUpFormFromGongstruct(onSave.measure_style, probe)
			}
		case *MembraneFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Membrane", true)
			} else {
				FillUpFormFromGongstruct(onSave.membrane, probe)
			}
		case *MetalFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Metal", true)
			} else {
				FillUpFormFromGongstruct(onSave.metal, probe)
			}
		case *MetronomeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Metronome", true)
			} else {
				FillUpFormFromGongstruct(onSave.metronome, probe)
			}
		case *Metronome_beamFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Metronome_beam", true)
			} else {
				FillUpFormFromGongstruct(onSave.metronome_beam, probe)
			}
		case *Metronome_noteFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Metronome_note", true)
			} else {
				FillUpFormFromGongstruct(onSave.metronome_note, probe)
			}
		case *Metronome_tiedFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Metronome_tied", true)
			} else {
				FillUpFormFromGongstruct(onSave.metronome_tied, probe)
			}
		case *Metronome_tupletFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Metronome_tuplet", true)
			} else {
				FillUpFormFromGongstruct(onSave.metronome_tuplet, probe)
			}
		case *Midi_deviceFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Midi_device", true)
			} else {
				FillUpFormFromGongstruct(onSave.midi_device, probe)
			}
		case *Midi_instrumentFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Midi_instrument", true)
			} else {
				FillUpFormFromGongstruct(onSave.midi_instrument, probe)
			}
		case *MiscellaneousFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Miscellaneous", true)
			} else {
				FillUpFormFromGongstruct(onSave.miscellaneous, probe)
			}
		case *Miscellaneous_fieldFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Miscellaneous_field", true)
			} else {
				FillUpFormFromGongstruct(onSave.miscellaneous_field, probe)
			}
		case *MordentFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Mordent", true)
			} else {
				FillUpFormFromGongstruct(onSave.mordent, probe)
			}
		case *Multiple_restFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Multiple_rest", true)
			} else {
				FillUpFormFromGongstruct(onSave.multiple_rest, probe)
			}
		case *Name_displayFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Name_display", true)
			} else {
				FillUpFormFromGongstruct(onSave.name_display, probe)
			}
		case *Non_arpeggiateFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Non_arpeggiate", true)
			} else {
				FillUpFormFromGongstruct(onSave.non_arpeggiate, probe)
			}
		case *NotationsFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Notations", true)
			} else {
				FillUpFormFromGongstruct(onSave.notations, probe)
			}
		case *NoteFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Note", true)
			} else {
				FillUpFormFromGongstruct(onSave.note, probe)
			}
		case *Note_sizeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Note_size", true)
			} else {
				FillUpFormFromGongstruct(onSave.note_size, probe)
			}
		case *Note_typeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Note_type", true)
			} else {
				FillUpFormFromGongstruct(onSave.note_type, probe)
			}
		case *NoteheadFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Notehead", true)
			} else {
				FillUpFormFromGongstruct(onSave.notehead, probe)
			}
		case *Notehead_textFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Notehead_text", true)
			} else {
				FillUpFormFromGongstruct(onSave.notehead_text, probe)
			}
		case *NumeralFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Numeral", true)
			} else {
				FillUpFormFromGongstruct(onSave.numeral, probe)
			}
		case *Numeral_keyFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Numeral_key", true)
			} else {
				FillUpFormFromGongstruct(onSave.numeral_key, probe)
			}
		case *Numeral_rootFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Numeral_root", true)
			} else {
				FillUpFormFromGongstruct(onSave.numeral_root, probe)
			}
		case *Octave_shiftFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Octave_shift", true)
			} else {
				FillUpFormFromGongstruct(onSave.octave_shift, probe)
			}
		case *OffsetFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Offset", true)
			} else {
				FillUpFormFromGongstruct(onSave.offset, probe)
			}
		case *OpusFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Opus", true)
			} else {
				FillUpFormFromGongstruct(onSave.opus, probe)
			}
		case *OrnamentsFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Ornaments", true)
			} else {
				FillUpFormFromGongstruct(onSave.ornaments, probe)
			}
		case *Other_appearanceFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Other_appearance", true)
			} else {
				FillUpFormFromGongstruct(onSave.other_appearance, probe)
			}
		case *Other_directionFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Other_direction", true)
			} else {
				FillUpFormFromGongstruct(onSave.other_direction, probe)
			}
		case *Other_listeningFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Other_listening", true)
			} else {
				FillUpFormFromGongstruct(onSave.other_listening, probe)
			}
		case *Other_notationFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Other_notation", true)
			} else {
				FillUpFormFromGongstruct(onSave.other_notation, probe)
			}
		case *Other_placement_textFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Other_placement_text", true)
			} else {
				FillUpFormFromGongstruct(onSave.other_placement_text, probe)
			}
		case *Other_playFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Other_play", true)
			} else {
				FillUpFormFromGongstruct(onSave.other_play, probe)
			}
		case *Other_textFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Other_text", true)
			} else {
				FillUpFormFromGongstruct(onSave.other_text, probe)
			}
		case *Page_layoutFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Page_layout", true)
			} else {
				FillUpFormFromGongstruct(onSave.page_layout, probe)
			}
		case *Page_marginsFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Page_margins", true)
			} else {
				FillUpFormFromGongstruct(onSave.page_margins, probe)
			}
		case *Part_clefFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Part_clef", true)
			} else {
				FillUpFormFromGongstruct(onSave.part_clef, probe)
			}
		case *Part_groupFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Part_group", true)
			} else {
				FillUpFormFromGongstruct(onSave.part_group, probe)
			}
		case *Part_linkFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Part_link", true)
			} else {
				FillUpFormFromGongstruct(onSave.part_link, probe)
			}
		case *Part_listFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Part_list", true)
			} else {
				FillUpFormFromGongstruct(onSave.part_list, probe)
			}
		case *Part_nameFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Part_name", true)
			} else {
				FillUpFormFromGongstruct(onSave.part_name, probe)
			}
		case *Part_symbolFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Part_symbol", true)
			} else {
				FillUpFormFromGongstruct(onSave.part_symbol, probe)
			}
		case *Part_transposeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Part_transpose", true)
			} else {
				FillUpFormFromGongstruct(onSave.part_transpose, probe)
			}
		case *PedalFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Pedal", true)
			} else {
				FillUpFormFromGongstruct(onSave.pedal, probe)
			}
		case *Pedal_tuningFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Pedal_tuning", true)
			} else {
				FillUpFormFromGongstruct(onSave.pedal_tuning, probe)
			}
		case *Per_minuteFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Per_minute", true)
			} else {
				FillUpFormFromGongstruct(onSave.per_minute, probe)
			}
		case *PercussionFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Percussion", true)
			} else {
				FillUpFormFromGongstruct(onSave.percussion, probe)
			}
		case *PitchFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Pitch", true)
			} else {
				FillUpFormFromGongstruct(onSave.pitch, probe)
			}
		case *PitchedFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Pitched", true)
			} else {
				FillUpFormFromGongstruct(onSave.pitched, probe)
			}
		case *Placement_textFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Placement_text", true)
			} else {
				FillUpFormFromGongstruct(onSave.placement_text, probe)
			}
		case *PlayFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Play", true)
			} else {
				FillUpFormFromGongstruct(onSave.play, probe)
			}
		case *PlayerFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Player", true)
			} else {
				FillUpFormFromGongstruct(onSave.player, probe)
			}
		case *Principal_voiceFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Principal_voice", true)
			} else {
				FillUpFormFromGongstruct(onSave.principal_voice, probe)
			}
		case *PrintFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Print", true)
			} else {
				FillUpFormFromGongstruct(onSave.print, probe)
			}
		case *ReleaseFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Release", true)
			} else {
				FillUpFormFromGongstruct(onSave.release, probe)
			}
		case *RepeatFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Repeat", true)
			} else {
				FillUpFormFromGongstruct(onSave.repeat, probe)
			}
		case *RestFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Rest", true)
			} else {
				FillUpFormFromGongstruct(onSave.rest, probe)
			}
		case *RootFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Root", true)
			} else {
				FillUpFormFromGongstruct(onSave.root, probe)
			}
		case *Root_stepFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Root_step", true)
			} else {
				FillUpFormFromGongstruct(onSave.root_step, probe)
			}
		case *ScalingFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Scaling", true)
			} else {
				FillUpFormFromGongstruct(onSave.scaling, probe)
			}
		case *ScordaturaFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Scordatura", true)
			} else {
				FillUpFormFromGongstruct(onSave.scordatura, probe)
			}
		case *Score_instrumentFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Score_instrument", true)
			} else {
				FillUpFormFromGongstruct(onSave.score_instrument, probe)
			}
		case *Score_partFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Score_part", true)
			} else {
				FillUpFormFromGongstruct(onSave.score_part, probe)
			}
		case *Score_partwiseFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Score_partwise", true)
			} else {
				FillUpFormFromGongstruct(onSave.score_partwise, probe)
			}
		case *Score_timewiseFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Score_timewise", true)
			} else {
				FillUpFormFromGongstruct(onSave.score_timewise, probe)
			}
		case *SegnoFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Segno", true)
			} else {
				FillUpFormFromGongstruct(onSave.segno, probe)
			}
		case *SlashFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Slash", true)
			} else {
				FillUpFormFromGongstruct(onSave.slash, probe)
			}
		case *SlideFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Slide", true)
			} else {
				FillUpFormFromGongstruct(onSave.slide, probe)
			}
		case *SlurFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Slur", true)
			} else {
				FillUpFormFromGongstruct(onSave.slur, probe)
			}
		case *SoundFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Sound", true)
			} else {
				FillUpFormFromGongstruct(onSave.sound, probe)
			}
		case *Staff_detailsFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Staff_details", true)
			} else {
				FillUpFormFromGongstruct(onSave.staff_details, probe)
			}
		case *Staff_divideFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Staff_divide", true)
			} else {
				FillUpFormFromGongstruct(onSave.staff_divide, probe)
			}
		case *Staff_layoutFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Staff_layout", true)
			} else {
				FillUpFormFromGongstruct(onSave.staff_layout, probe)
			}
		case *Staff_sizeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Staff_size", true)
			} else {
				FillUpFormFromGongstruct(onSave.staff_size, probe)
			}
		case *Staff_tuningFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Staff_tuning", true)
			} else {
				FillUpFormFromGongstruct(onSave.staff_tuning, probe)
			}
		case *StemFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Stem", true)
			} else {
				FillUpFormFromGongstruct(onSave.stem, probe)
			}
		case *StickFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Stick", true)
			} else {
				FillUpFormFromGongstruct(onSave.stick, probe)
			}
		case *String_muteFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "String_mute", true)
			} else {
				FillUpFormFromGongstruct(onSave.string_mute, probe)
			}
		case *String_typeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "String_type", true)
			} else {
				FillUpFormFromGongstruct(onSave.string_type, probe)
			}
		case *Strong_accentFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Strong_accent", true)
			} else {
				FillUpFormFromGongstruct(onSave.strong_accent, probe)
			}
		case *Style_textFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Style_text", true)
			} else {
				FillUpFormFromGongstruct(onSave.style_text, probe)
			}
		case *SupportsFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Supports", true)
			} else {
				FillUpFormFromGongstruct(onSave.supports, probe)
			}
		case *SwingFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Swing", true)
			} else {
				FillUpFormFromGongstruct(onSave.swing, probe)
			}
		case *SyncFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Sync", true)
			} else {
				FillUpFormFromGongstruct(onSave.sync, probe)
			}
		case *System_dividersFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "System_dividers", true)
			} else {
				FillUpFormFromGongstruct(onSave.system_dividers, probe)
			}
		case *System_layoutFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "System_layout", true)
			} else {
				FillUpFormFromGongstruct(onSave.system_layout, probe)
			}
		case *System_marginsFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "System_margins", true)
			} else {
				FillUpFormFromGongstruct(onSave.system_margins, probe)
			}
		case *TapFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Tap", true)
			} else {
				FillUpFormFromGongstruct(onSave.tap, probe)
			}
		case *TechnicalFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Technical", true)
			} else {
				FillUpFormFromGongstruct(onSave.technical, probe)
			}
		case *Text_element_dataFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Text_element_data", true)
			} else {
				FillUpFormFromGongstruct(onSave.text_element_data, probe)
			}
		case *TieFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Tie", true)
			} else {
				FillUpFormFromGongstruct(onSave.tie, probe)
			}
		case *TiedFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Tied", true)
			} else {
				FillUpFormFromGongstruct(onSave.tied, probe)
			}
		case *TimeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Time", true)
			} else {
				FillUpFormFromGongstruct(onSave.time, probe)
			}
		case *Time_modificationFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Time_modification", true)
			} else {
				FillUpFormFromGongstruct(onSave.time_modification, probe)
			}
		case *TimpaniFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Timpani", true)
			} else {
				FillUpFormFromGongstruct(onSave.timpani, probe)
			}
		case *TransposeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Transpose", true)
			} else {
				FillUpFormFromGongstruct(onSave.transpose, probe)
			}
		case *TremoloFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Tremolo", true)
			} else {
				FillUpFormFromGongstruct(onSave.tremolo, probe)
			}
		case *TupletFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Tuplet", true)
			} else {
				FillUpFormFromGongstruct(onSave.tuplet, probe)
			}
		case *Tuplet_dotFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Tuplet_dot", true)
			} else {
				FillUpFormFromGongstruct(onSave.tuplet_dot, probe)
			}
		case *Tuplet_numberFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Tuplet_number", true)
			} else {
				FillUpFormFromGongstruct(onSave.tuplet_number, probe)
			}
		case *Tuplet_portionFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Tuplet_portion", true)
			} else {
				FillUpFormFromGongstruct(onSave.tuplet_portion, probe)
			}
		case *Tuplet_typeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Tuplet_type", true)
			} else {
				FillUpFormFromGongstruct(onSave.tuplet_type, probe)
			}
		case *Typed_textFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Typed_text", true)
			} else {
				FillUpFormFromGongstruct(onSave.typed_text, probe)
			}
		case *UnpitchedFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Unpitched", true)
			} else {
				FillUpFormFromGongstruct(onSave.unpitched, probe)
			}
		case *Virtual_instrumentFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Virtual_instrument", true)
			} else {
				FillUpFormFromGongstruct(onSave.virtual_instrument, probe)
			}
		case *WaitFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Wait", true)
			} else {
				FillUpFormFromGongstruct(onSave.wait, probe)
			}
		case *Wavy_lineFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Wavy_line", true)
			} else {
				FillUpFormFromGongstruct(onSave.wavy_line, probe)
			}
		case *WedgeFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Wedge", true)
			} else {
				FillUpFormFromGongstruct(onSave.wedge, probe)
			}
		case *WoodFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Wood", true)
			} else {
				FillUpFormFromGongstruct(onSave.wood, probe)
			}
		case *WorkFormCallback:
			if onSave.CreationMode {
				FillUpFormFromGongstructName(probe, "Work", true)
			} else {
				FillUpFormFromGongstruct(onSave.work, probe)
			}
		}
	}
}

func FillUpFormFromGongstructName(
	probe *Probe,
	gongstructName string,
	isNewInstance bool,
) {
	formStage := probe.formStage
	formStage.Reset()

	var prefix string

	if isNewInstance {
		prefix = ""
	} else {
		prefix = ""
	}

	switch gongstructName {
	// insertion point
	case "A_directive":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "A_directive Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_directiveFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_directive := new(models.A_directive)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_directive, formGroup, probe)
	case "A_measure":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "A_measure Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_measureFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_measure := new(models.A_measure)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_measure, formGroup, probe)
	case "A_measure_1":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "A_measure_1 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_measure_1FormCallback(
			nil,
			probe,
			formGroup,
		)
		a_measure_1 := new(models.A_measure_1)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_measure_1, formGroup, probe)
	case "A_part":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "A_part Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_partFormCallback(
			nil,
			probe,
			formGroup,
		)
		a_part := new(models.A_part)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_part, formGroup, probe)
	case "A_part_1":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "A_part_1 Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__A_part_1FormCallback(
			nil,
			probe,
			formGroup,
		)
		a_part_1 := new(models.A_part_1)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(a_part_1, formGroup, probe)
	case "Accidental":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Accidental Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AccidentalFormCallback(
			nil,
			probe,
			formGroup,
		)
		accidental := new(models.Accidental)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(accidental, formGroup, probe)
	case "Accidental_mark":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Accidental_mark Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Accidental_markFormCallback(
			nil,
			probe,
			formGroup,
		)
		accidental_mark := new(models.Accidental_mark)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(accidental_mark, formGroup, probe)
	case "Accidental_text":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Accidental_text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Accidental_textFormCallback(
			nil,
			probe,
			formGroup,
		)
		accidental_text := new(models.Accidental_text)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(accidental_text, formGroup, probe)
	case "Accord":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Accord Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AccordFormCallback(
			nil,
			probe,
			formGroup,
		)
		accord := new(models.Accord)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(accord, formGroup, probe)
	case "Accordion_registration":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Accordion_registration Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Accordion_registrationFormCallback(
			nil,
			probe,
			formGroup,
		)
		accordion_registration := new(models.Accordion_registration)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(accordion_registration, formGroup, probe)
	case "Appearance":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Appearance Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AppearanceFormCallback(
			nil,
			probe,
			formGroup,
		)
		appearance := new(models.Appearance)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(appearance, formGroup, probe)
	case "Arpeggiate":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Arpeggiate Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArpeggiateFormCallback(
			nil,
			probe,
			formGroup,
		)
		arpeggiate := new(models.Arpeggiate)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(arpeggiate, formGroup, probe)
	case "Arrow":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Arrow Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArrowFormCallback(
			nil,
			probe,
			formGroup,
		)
		arrow := new(models.Arrow)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(arrow, formGroup, probe)
	case "Articulations":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Articulations Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ArticulationsFormCallback(
			nil,
			probe,
			formGroup,
		)
		articulations := new(models.Articulations)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(articulations, formGroup, probe)
	case "Assess":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Assess Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AssessFormCallback(
			nil,
			probe,
			formGroup,
		)
		assess := new(models.Assess)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(assess, formGroup, probe)
	case "Attributes":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Attributes Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__AttributesFormCallback(
			nil,
			probe,
			formGroup,
		)
		attributes := new(models.Attributes)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(attributes, formGroup, probe)
	case "Backup":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Backup Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BackupFormCallback(
			nil,
			probe,
			formGroup,
		)
		backup := new(models.Backup)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(backup, formGroup, probe)
	case "Bar_style_color":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Bar_style_color Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Bar_style_colorFormCallback(
			nil,
			probe,
			formGroup,
		)
		bar_style_color := new(models.Bar_style_color)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bar_style_color, formGroup, probe)
	case "Barline":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Barline Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BarlineFormCallback(
			nil,
			probe,
			formGroup,
		)
		barline := new(models.Barline)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(barline, formGroup, probe)
	case "Barre":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Barre Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BarreFormCallback(
			nil,
			probe,
			formGroup,
		)
		barre := new(models.Barre)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(barre, formGroup, probe)
	case "Bass":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Bass Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BassFormCallback(
			nil,
			probe,
			formGroup,
		)
		bass := new(models.Bass)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bass, formGroup, probe)
	case "Bass_step":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Bass_step Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Bass_stepFormCallback(
			nil,
			probe,
			formGroup,
		)
		bass_step := new(models.Bass_step)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bass_step, formGroup, probe)
	case "Beam":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Beam Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BeamFormCallback(
			nil,
			probe,
			formGroup,
		)
		beam := new(models.Beam)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(beam, formGroup, probe)
	case "Beat_repeat":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Beat_repeat Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Beat_repeatFormCallback(
			nil,
			probe,
			formGroup,
		)
		beat_repeat := new(models.Beat_repeat)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(beat_repeat, formGroup, probe)
	case "Beat_unit_tied":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Beat_unit_tied Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Beat_unit_tiedFormCallback(
			nil,
			probe,
			formGroup,
		)
		beat_unit_tied := new(models.Beat_unit_tied)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(beat_unit_tied, formGroup, probe)
	case "Beater":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Beater Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BeaterFormCallback(
			nil,
			probe,
			formGroup,
		)
		beater := new(models.Beater)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(beater, formGroup, probe)
	case "Bend":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Bend Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BendFormCallback(
			nil,
			probe,
			formGroup,
		)
		bend := new(models.Bend)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bend, formGroup, probe)
	case "Bookmark":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Bookmark Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BookmarkFormCallback(
			nil,
			probe,
			formGroup,
		)
		bookmark := new(models.Bookmark)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bookmark, formGroup, probe)
	case "Bracket":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Bracket Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__BracketFormCallback(
			nil,
			probe,
			formGroup,
		)
		bracket := new(models.Bracket)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(bracket, formGroup, probe)
	case "Breath_mark":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Breath_mark Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Breath_markFormCallback(
			nil,
			probe,
			formGroup,
		)
		breath_mark := new(models.Breath_mark)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(breath_mark, formGroup, probe)
	case "Caesura":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Caesura Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CaesuraFormCallback(
			nil,
			probe,
			formGroup,
		)
		caesura := new(models.Caesura)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(caesura, formGroup, probe)
	case "Cancel":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Cancel Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CancelFormCallback(
			nil,
			probe,
			formGroup,
		)
		cancel := new(models.Cancel)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(cancel, formGroup, probe)
	case "Clef":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Clef Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ClefFormCallback(
			nil,
			probe,
			formGroup,
		)
		clef := new(models.Clef)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(clef, formGroup, probe)
	case "Coda":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Coda Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CodaFormCallback(
			nil,
			probe,
			formGroup,
		)
		coda := new(models.Coda)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(coda, formGroup, probe)
	case "Credit":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Credit Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__CreditFormCallback(
			nil,
			probe,
			formGroup,
		)
		credit := new(models.Credit)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(credit, formGroup, probe)
	case "Dashes":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Dashes Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DashesFormCallback(
			nil,
			probe,
			formGroup,
		)
		dashes := new(models.Dashes)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(dashes, formGroup, probe)
	case "Defaults":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Defaults Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DefaultsFormCallback(
			nil,
			probe,
			formGroup,
		)
		defaults := new(models.Defaults)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(defaults, formGroup, probe)
	case "Degree":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Degree Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DegreeFormCallback(
			nil,
			probe,
			formGroup,
		)
		degree := new(models.Degree)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(degree, formGroup, probe)
	case "Degree_alter":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Degree_alter Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Degree_alterFormCallback(
			nil,
			probe,
			formGroup,
		)
		degree_alter := new(models.Degree_alter)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(degree_alter, formGroup, probe)
	case "Degree_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Degree_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Degree_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		degree_type := new(models.Degree_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(degree_type, formGroup, probe)
	case "Degree_value":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Degree_value Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Degree_valueFormCallback(
			nil,
			probe,
			formGroup,
		)
		degree_value := new(models.Degree_value)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(degree_value, formGroup, probe)
	case "Direction":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Direction Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DirectionFormCallback(
			nil,
			probe,
			formGroup,
		)
		direction := new(models.Direction)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(direction, formGroup, probe)
	case "Direction_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Direction_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Direction_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		direction_type := new(models.Direction_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(direction_type, formGroup, probe)
	case "Distance":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Distance Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DistanceFormCallback(
			nil,
			probe,
			formGroup,
		)
		distance := new(models.Distance)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(distance, formGroup, probe)
	case "Double":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Double Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DoubleFormCallback(
			nil,
			probe,
			formGroup,
		)
		double := new(models.Double)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(double, formGroup, probe)
	case "Dynamics":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Dynamics Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__DynamicsFormCallback(
			nil,
			probe,
			formGroup,
		)
		dynamics := new(models.Dynamics)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(dynamics, formGroup, probe)
	case "Effect":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Effect Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EffectFormCallback(
			nil,
			probe,
			formGroup,
		)
		effect := new(models.Effect)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(effect, formGroup, probe)
	case "Elision":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Elision Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ElisionFormCallback(
			nil,
			probe,
			formGroup,
		)
		elision := new(models.Elision)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(elision, formGroup, probe)
	case "Empty":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Empty Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EmptyFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty := new(models.Empty)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty, formGroup, probe)
	case "Empty_font":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Empty_font Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_fontFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_font := new(models.Empty_font)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_font, formGroup, probe)
	case "Empty_line":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Empty_line Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_lineFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_line := new(models.Empty_line)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_line, formGroup, probe)
	case "Empty_placement":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Empty_placement Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_placementFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_placement := new(models.Empty_placement)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_placement, formGroup, probe)
	case "Empty_placement_smufl":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Empty_placement_smufl Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_placement_smuflFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_placement_smufl := new(models.Empty_placement_smufl)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_placement_smufl, formGroup, probe)
	case "Empty_print_object_style_align":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Empty_print_object_style_align Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_print_object_style_alignFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_print_object_style_align := new(models.Empty_print_object_style_align)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_print_object_style_align, formGroup, probe)
	case "Empty_print_style":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Empty_print_style Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_print_styleFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_print_style := new(models.Empty_print_style)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_print_style, formGroup, probe)
	case "Empty_print_style_align":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Empty_print_style_align Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_print_style_alignFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_print_style_align := new(models.Empty_print_style_align)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_print_style_align, formGroup, probe)
	case "Empty_print_style_align_id":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Empty_print_style_align_id Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_print_style_align_idFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_print_style_align_id := new(models.Empty_print_style_align_id)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_print_style_align_id, formGroup, probe)
	case "Empty_trill_sound":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Empty_trill_sound Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Empty_trill_soundFormCallback(
			nil,
			probe,
			formGroup,
		)
		empty_trill_sound := new(models.Empty_trill_sound)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(empty_trill_sound, formGroup, probe)
	case "Encoding":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Encoding Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EncodingFormCallback(
			nil,
			probe,
			formGroup,
		)
		encoding := new(models.Encoding)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(encoding, formGroup, probe)
	case "Ending":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Ending Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__EndingFormCallback(
			nil,
			probe,
			formGroup,
		)
		ending := new(models.Ending)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(ending, formGroup, probe)
	case "Extend":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Extend Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ExtendFormCallback(
			nil,
			probe,
			formGroup,
		)
		extend := new(models.Extend)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(extend, formGroup, probe)
	case "Feature":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Feature Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FeatureFormCallback(
			nil,
			probe,
			formGroup,
		)
		feature := new(models.Feature)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(feature, formGroup, probe)
	case "Fermata":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Fermata Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FermataFormCallback(
			nil,
			probe,
			formGroup,
		)
		fermata := new(models.Fermata)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(fermata, formGroup, probe)
	case "Figure":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Figure Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FigureFormCallback(
			nil,
			probe,
			formGroup,
		)
		figure := new(models.Figure)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(figure, formGroup, probe)
	case "Figured_bass":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Figured_bass Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Figured_bassFormCallback(
			nil,
			probe,
			formGroup,
		)
		figured_bass := new(models.Figured_bass)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(figured_bass, formGroup, probe)
	case "Fingering":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Fingering Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FingeringFormCallback(
			nil,
			probe,
			formGroup,
		)
		fingering := new(models.Fingering)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(fingering, formGroup, probe)
	case "First_fret":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "First_fret Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__First_fretFormCallback(
			nil,
			probe,
			formGroup,
		)
		first_fret := new(models.First_fret)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(first_fret, formGroup, probe)
	case "For_part":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "For_part Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__For_partFormCallback(
			nil,
			probe,
			formGroup,
		)
		for_part := new(models.For_part)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(for_part, formGroup, probe)
	case "Formatted_symbol":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Formatted_symbol Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Formatted_symbolFormCallback(
			nil,
			probe,
			formGroup,
		)
		formatted_symbol := new(models.Formatted_symbol)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formatted_symbol, formGroup, probe)
	case "Formatted_symbol_id":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Formatted_symbol_id Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Formatted_symbol_idFormCallback(
			nil,
			probe,
			formGroup,
		)
		formatted_symbol_id := new(models.Formatted_symbol_id)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formatted_symbol_id, formGroup, probe)
	case "Formatted_text":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Formatted_text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Formatted_textFormCallback(
			nil,
			probe,
			formGroup,
		)
		formatted_text := new(models.Formatted_text)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formatted_text, formGroup, probe)
	case "Formatted_text_id":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Formatted_text_id Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Formatted_text_idFormCallback(
			nil,
			probe,
			formGroup,
		)
		formatted_text_id := new(models.Formatted_text_id)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(formatted_text_id, formGroup, probe)
	case "Forward":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Forward Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ForwardFormCallback(
			nil,
			probe,
			formGroup,
		)
		forward := new(models.Forward)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(forward, formGroup, probe)
	case "Frame":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Frame Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FrameFormCallback(
			nil,
			probe,
			formGroup,
		)
		frame := new(models.Frame)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(frame, formGroup, probe)
	case "Frame_note":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Frame_note Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Frame_noteFormCallback(
			nil,
			probe,
			formGroup,
		)
		frame_note := new(models.Frame_note)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(frame_note, formGroup, probe)
	case "Fret":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Fret Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__FretFormCallback(
			nil,
			probe,
			formGroup,
		)
		fret := new(models.Fret)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(fret, formGroup, probe)
	case "Glass":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Glass Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GlassFormCallback(
			nil,
			probe,
			formGroup,
		)
		glass := new(models.Glass)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(glass, formGroup, probe)
	case "Glissando":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Glissando Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GlissandoFormCallback(
			nil,
			probe,
			formGroup,
		)
		glissando := new(models.Glissando)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(glissando, formGroup, probe)
	case "Glyph":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Glyph Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GlyphFormCallback(
			nil,
			probe,
			formGroup,
		)
		glyph := new(models.Glyph)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(glyph, formGroup, probe)
	case "Grace":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Grace Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GraceFormCallback(
			nil,
			probe,
			formGroup,
		)
		grace := new(models.Grace)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(grace, formGroup, probe)
	case "Group_barline":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Group_barline Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Group_barlineFormCallback(
			nil,
			probe,
			formGroup,
		)
		group_barline := new(models.Group_barline)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(group_barline, formGroup, probe)
	case "Group_name":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Group_name Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Group_nameFormCallback(
			nil,
			probe,
			formGroup,
		)
		group_name := new(models.Group_name)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(group_name, formGroup, probe)
	case "Group_symbol":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Group_symbol Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Group_symbolFormCallback(
			nil,
			probe,
			formGroup,
		)
		group_symbol := new(models.Group_symbol)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(group_symbol, formGroup, probe)
	case "Grouping":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Grouping Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__GroupingFormCallback(
			nil,
			probe,
			formGroup,
		)
		grouping := new(models.Grouping)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(grouping, formGroup, probe)
	case "Hammer_on_pull_off":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Hammer_on_pull_off Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Hammer_on_pull_offFormCallback(
			nil,
			probe,
			formGroup,
		)
		hammer_on_pull_off := new(models.Hammer_on_pull_off)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(hammer_on_pull_off, formGroup, probe)
	case "Handbell":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Handbell Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HandbellFormCallback(
			nil,
			probe,
			formGroup,
		)
		handbell := new(models.Handbell)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(handbell, formGroup, probe)
	case "Harmon_closed":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Harmon_closed Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Harmon_closedFormCallback(
			nil,
			probe,
			formGroup,
		)
		harmon_closed := new(models.Harmon_closed)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(harmon_closed, formGroup, probe)
	case "Harmon_mute":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Harmon_mute Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Harmon_muteFormCallback(
			nil,
			probe,
			formGroup,
		)
		harmon_mute := new(models.Harmon_mute)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(harmon_mute, formGroup, probe)
	case "Harmonic":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Harmonic Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HarmonicFormCallback(
			nil,
			probe,
			formGroup,
		)
		harmonic := new(models.Harmonic)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(harmonic, formGroup, probe)
	case "Harmony":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Harmony Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HarmonyFormCallback(
			nil,
			probe,
			formGroup,
		)
		harmony := new(models.Harmony)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(harmony, formGroup, probe)
	case "Harmony_alter":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Harmony_alter Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Harmony_alterFormCallback(
			nil,
			probe,
			formGroup,
		)
		harmony_alter := new(models.Harmony_alter)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(harmony_alter, formGroup, probe)
	case "Harp_pedals":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Harp_pedals Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Harp_pedalsFormCallback(
			nil,
			probe,
			formGroup,
		)
		harp_pedals := new(models.Harp_pedals)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(harp_pedals, formGroup, probe)
	case "Heel_toe":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Heel_toe Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Heel_toeFormCallback(
			nil,
			probe,
			formGroup,
		)
		heel_toe := new(models.Heel_toe)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(heel_toe, formGroup, probe)
	case "Hole":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Hole Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__HoleFormCallback(
			nil,
			probe,
			formGroup,
		)
		hole := new(models.Hole)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(hole, formGroup, probe)
	case "Hole_closed":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Hole_closed Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Hole_closedFormCallback(
			nil,
			probe,
			formGroup,
		)
		hole_closed := new(models.Hole_closed)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(hole_closed, formGroup, probe)
	case "Horizontal_turn":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Horizontal_turn Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Horizontal_turnFormCallback(
			nil,
			probe,
			formGroup,
		)
		horizontal_turn := new(models.Horizontal_turn)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(horizontal_turn, formGroup, probe)
	case "Identification":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Identification Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__IdentificationFormCallback(
			nil,
			probe,
			formGroup,
		)
		identification := new(models.Identification)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(identification, formGroup, probe)
	case "Image":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Image Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ImageFormCallback(
			nil,
			probe,
			formGroup,
		)
		image := new(models.Image)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(image, formGroup, probe)
	case "Instrument":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Instrument Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InstrumentFormCallback(
			nil,
			probe,
			formGroup,
		)
		instrument := new(models.Instrument)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(instrument, formGroup, probe)
	case "Instrument_change":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Instrument_change Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Instrument_changeFormCallback(
			nil,
			probe,
			formGroup,
		)
		instrument_change := new(models.Instrument_change)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(instrument_change, formGroup, probe)
	case "Instrument_link":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Instrument_link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Instrument_linkFormCallback(
			nil,
			probe,
			formGroup,
		)
		instrument_link := new(models.Instrument_link)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(instrument_link, formGroup, probe)
	case "Interchangeable":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Interchangeable Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InterchangeableFormCallback(
			nil,
			probe,
			formGroup,
		)
		interchangeable := new(models.Interchangeable)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(interchangeable, formGroup, probe)
	case "Inversion":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Inversion Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__InversionFormCallback(
			nil,
			probe,
			formGroup,
		)
		inversion := new(models.Inversion)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(inversion, formGroup, probe)
	case "Key":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Key Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KeyFormCallback(
			nil,
			probe,
			formGroup,
		)
		key := new(models.Key)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(key, formGroup, probe)
	case "Key_accidental":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Key_accidental Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Key_accidentalFormCallback(
			nil,
			probe,
			formGroup,
		)
		key_accidental := new(models.Key_accidental)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(key_accidental, formGroup, probe)
	case "Key_octave":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Key_octave Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Key_octaveFormCallback(
			nil,
			probe,
			formGroup,
		)
		key_octave := new(models.Key_octave)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(key_octave, formGroup, probe)
	case "Kind":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Kind Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__KindFormCallback(
			nil,
			probe,
			formGroup,
		)
		kind := new(models.Kind)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(kind, formGroup, probe)
	case "Level":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Level Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LevelFormCallback(
			nil,
			probe,
			formGroup,
		)
		level := new(models.Level)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(level, formGroup, probe)
	case "Line_detail":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Line_detail Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Line_detailFormCallback(
			nil,
			probe,
			formGroup,
		)
		line_detail := new(models.Line_detail)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(line_detail, formGroup, probe)
	case "Line_width":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Line_width Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Line_widthFormCallback(
			nil,
			probe,
			formGroup,
		)
		line_width := new(models.Line_width)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(line_width, formGroup, probe)
	case "Link":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LinkFormCallback(
			nil,
			probe,
			formGroup,
		)
		link := new(models.Link)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(link, formGroup, probe)
	case "Listen":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Listen Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ListenFormCallback(
			nil,
			probe,
			formGroup,
		)
		listen := new(models.Listen)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(listen, formGroup, probe)
	case "Listening":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Listening Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ListeningFormCallback(
			nil,
			probe,
			formGroup,
		)
		listening := new(models.Listening)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(listening, formGroup, probe)
	case "Lyric":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Lyric Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__LyricFormCallback(
			nil,
			probe,
			formGroup,
		)
		lyric := new(models.Lyric)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(lyric, formGroup, probe)
	case "Lyric_font":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Lyric_font Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Lyric_fontFormCallback(
			nil,
			probe,
			formGroup,
		)
		lyric_font := new(models.Lyric_font)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(lyric_font, formGroup, probe)
	case "Lyric_language":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Lyric_language Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Lyric_languageFormCallback(
			nil,
			probe,
			formGroup,
		)
		lyric_language := new(models.Lyric_language)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(lyric_language, formGroup, probe)
	case "Measure_layout":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Measure_layout Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Measure_layoutFormCallback(
			nil,
			probe,
			formGroup,
		)
		measure_layout := new(models.Measure_layout)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(measure_layout, formGroup, probe)
	case "Measure_numbering":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Measure_numbering Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Measure_numberingFormCallback(
			nil,
			probe,
			formGroup,
		)
		measure_numbering := new(models.Measure_numbering)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(measure_numbering, formGroup, probe)
	case "Measure_repeat":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Measure_repeat Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Measure_repeatFormCallback(
			nil,
			probe,
			formGroup,
		)
		measure_repeat := new(models.Measure_repeat)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(measure_repeat, formGroup, probe)
	case "Measure_style":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Measure_style Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Measure_styleFormCallback(
			nil,
			probe,
			formGroup,
		)
		measure_style := new(models.Measure_style)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(measure_style, formGroup, probe)
	case "Membrane":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Membrane Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MembraneFormCallback(
			nil,
			probe,
			formGroup,
		)
		membrane := new(models.Membrane)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(membrane, formGroup, probe)
	case "Metal":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Metal Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MetalFormCallback(
			nil,
			probe,
			formGroup,
		)
		metal := new(models.Metal)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(metal, formGroup, probe)
	case "Metronome":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Metronome Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MetronomeFormCallback(
			nil,
			probe,
			formGroup,
		)
		metronome := new(models.Metronome)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(metronome, formGroup, probe)
	case "Metronome_beam":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Metronome_beam Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Metronome_beamFormCallback(
			nil,
			probe,
			formGroup,
		)
		metronome_beam := new(models.Metronome_beam)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(metronome_beam, formGroup, probe)
	case "Metronome_note":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Metronome_note Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Metronome_noteFormCallback(
			nil,
			probe,
			formGroup,
		)
		metronome_note := new(models.Metronome_note)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(metronome_note, formGroup, probe)
	case "Metronome_tied":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Metronome_tied Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Metronome_tiedFormCallback(
			nil,
			probe,
			formGroup,
		)
		metronome_tied := new(models.Metronome_tied)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(metronome_tied, formGroup, probe)
	case "Metronome_tuplet":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Metronome_tuplet Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Metronome_tupletFormCallback(
			nil,
			probe,
			formGroup,
		)
		metronome_tuplet := new(models.Metronome_tuplet)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(metronome_tuplet, formGroup, probe)
	case "Midi_device":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Midi_device Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Midi_deviceFormCallback(
			nil,
			probe,
			formGroup,
		)
		midi_device := new(models.Midi_device)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(midi_device, formGroup, probe)
	case "Midi_instrument":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Midi_instrument Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Midi_instrumentFormCallback(
			nil,
			probe,
			formGroup,
		)
		midi_instrument := new(models.Midi_instrument)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(midi_instrument, formGroup, probe)
	case "Miscellaneous":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Miscellaneous Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MiscellaneousFormCallback(
			nil,
			probe,
			formGroup,
		)
		miscellaneous := new(models.Miscellaneous)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(miscellaneous, formGroup, probe)
	case "Miscellaneous_field":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Miscellaneous_field Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Miscellaneous_fieldFormCallback(
			nil,
			probe,
			formGroup,
		)
		miscellaneous_field := new(models.Miscellaneous_field)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(miscellaneous_field, formGroup, probe)
	case "Mordent":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Mordent Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__MordentFormCallback(
			nil,
			probe,
			formGroup,
		)
		mordent := new(models.Mordent)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(mordent, formGroup, probe)
	case "Multiple_rest":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Multiple_rest Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Multiple_restFormCallback(
			nil,
			probe,
			formGroup,
		)
		multiple_rest := new(models.Multiple_rest)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(multiple_rest, formGroup, probe)
	case "Name_display":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Name_display Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Name_displayFormCallback(
			nil,
			probe,
			formGroup,
		)
		name_display := new(models.Name_display)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(name_display, formGroup, probe)
	case "Non_arpeggiate":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Non_arpeggiate Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Non_arpeggiateFormCallback(
			nil,
			probe,
			formGroup,
		)
		non_arpeggiate := new(models.Non_arpeggiate)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(non_arpeggiate, formGroup, probe)
	case "Notations":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Notations Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NotationsFormCallback(
			nil,
			probe,
			formGroup,
		)
		notations := new(models.Notations)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(notations, formGroup, probe)
	case "Note":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Note Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteFormCallback(
			nil,
			probe,
			formGroup,
		)
		note := new(models.Note)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(note, formGroup, probe)
	case "Note_size":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Note_size Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Note_sizeFormCallback(
			nil,
			probe,
			formGroup,
		)
		note_size := new(models.Note_size)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(note_size, formGroup, probe)
	case "Note_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Note_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Note_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		note_type := new(models.Note_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(note_type, formGroup, probe)
	case "Notehead":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Notehead Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NoteheadFormCallback(
			nil,
			probe,
			formGroup,
		)
		notehead := new(models.Notehead)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(notehead, formGroup, probe)
	case "Notehead_text":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Notehead_text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Notehead_textFormCallback(
			nil,
			probe,
			formGroup,
		)
		notehead_text := new(models.Notehead_text)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(notehead_text, formGroup, probe)
	case "Numeral":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Numeral Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__NumeralFormCallback(
			nil,
			probe,
			formGroup,
		)
		numeral := new(models.Numeral)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(numeral, formGroup, probe)
	case "Numeral_key":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Numeral_key Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Numeral_keyFormCallback(
			nil,
			probe,
			formGroup,
		)
		numeral_key := new(models.Numeral_key)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(numeral_key, formGroup, probe)
	case "Numeral_root":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Numeral_root Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Numeral_rootFormCallback(
			nil,
			probe,
			formGroup,
		)
		numeral_root := new(models.Numeral_root)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(numeral_root, formGroup, probe)
	case "Octave_shift":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Octave_shift Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Octave_shiftFormCallback(
			nil,
			probe,
			formGroup,
		)
		octave_shift := new(models.Octave_shift)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(octave_shift, formGroup, probe)
	case "Offset":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Offset Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OffsetFormCallback(
			nil,
			probe,
			formGroup,
		)
		offset := new(models.Offset)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(offset, formGroup, probe)
	case "Opus":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Opus Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OpusFormCallback(
			nil,
			probe,
			formGroup,
		)
		opus := new(models.Opus)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(opus, formGroup, probe)
	case "Ornaments":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Ornaments Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__OrnamentsFormCallback(
			nil,
			probe,
			formGroup,
		)
		ornaments := new(models.Ornaments)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(ornaments, formGroup, probe)
	case "Other_appearance":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Other_appearance Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_appearanceFormCallback(
			nil,
			probe,
			formGroup,
		)
		other_appearance := new(models.Other_appearance)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(other_appearance, formGroup, probe)
	case "Other_direction":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Other_direction Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_directionFormCallback(
			nil,
			probe,
			formGroup,
		)
		other_direction := new(models.Other_direction)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(other_direction, formGroup, probe)
	case "Other_listening":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Other_listening Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_listeningFormCallback(
			nil,
			probe,
			formGroup,
		)
		other_listening := new(models.Other_listening)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(other_listening, formGroup, probe)
	case "Other_notation":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Other_notation Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_notationFormCallback(
			nil,
			probe,
			formGroup,
		)
		other_notation := new(models.Other_notation)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(other_notation, formGroup, probe)
	case "Other_placement_text":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Other_placement_text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_placement_textFormCallback(
			nil,
			probe,
			formGroup,
		)
		other_placement_text := new(models.Other_placement_text)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(other_placement_text, formGroup, probe)
	case "Other_play":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Other_play Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_playFormCallback(
			nil,
			probe,
			formGroup,
		)
		other_play := new(models.Other_play)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(other_play, formGroup, probe)
	case "Other_text":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Other_text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Other_textFormCallback(
			nil,
			probe,
			formGroup,
		)
		other_text := new(models.Other_text)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(other_text, formGroup, probe)
	case "Page_layout":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Page_layout Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Page_layoutFormCallback(
			nil,
			probe,
			formGroup,
		)
		page_layout := new(models.Page_layout)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(page_layout, formGroup, probe)
	case "Page_margins":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Page_margins Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Page_marginsFormCallback(
			nil,
			probe,
			formGroup,
		)
		page_margins := new(models.Page_margins)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(page_margins, formGroup, probe)
	case "Part_clef":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Part_clef Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_clefFormCallback(
			nil,
			probe,
			formGroup,
		)
		part_clef := new(models.Part_clef)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(part_clef, formGroup, probe)
	case "Part_group":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Part_group Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_groupFormCallback(
			nil,
			probe,
			formGroup,
		)
		part_group := new(models.Part_group)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(part_group, formGroup, probe)
	case "Part_link":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Part_link Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_linkFormCallback(
			nil,
			probe,
			formGroup,
		)
		part_link := new(models.Part_link)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(part_link, formGroup, probe)
	case "Part_list":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Part_list Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_listFormCallback(
			nil,
			probe,
			formGroup,
		)
		part_list := new(models.Part_list)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(part_list, formGroup, probe)
	case "Part_name":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Part_name Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_nameFormCallback(
			nil,
			probe,
			formGroup,
		)
		part_name := new(models.Part_name)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(part_name, formGroup, probe)
	case "Part_symbol":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Part_symbol Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_symbolFormCallback(
			nil,
			probe,
			formGroup,
		)
		part_symbol := new(models.Part_symbol)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(part_symbol, formGroup, probe)
	case "Part_transpose":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Part_transpose Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Part_transposeFormCallback(
			nil,
			probe,
			formGroup,
		)
		part_transpose := new(models.Part_transpose)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(part_transpose, formGroup, probe)
	case "Pedal":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Pedal Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PedalFormCallback(
			nil,
			probe,
			formGroup,
		)
		pedal := new(models.Pedal)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(pedal, formGroup, probe)
	case "Pedal_tuning":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Pedal_tuning Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Pedal_tuningFormCallback(
			nil,
			probe,
			formGroup,
		)
		pedal_tuning := new(models.Pedal_tuning)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(pedal_tuning, formGroup, probe)
	case "Per_minute":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Per_minute Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Per_minuteFormCallback(
			nil,
			probe,
			formGroup,
		)
		per_minute := new(models.Per_minute)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(per_minute, formGroup, probe)
	case "Percussion":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Percussion Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PercussionFormCallback(
			nil,
			probe,
			formGroup,
		)
		percussion := new(models.Percussion)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(percussion, formGroup, probe)
	case "Pitch":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Pitch Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PitchFormCallback(
			nil,
			probe,
			formGroup,
		)
		pitch := new(models.Pitch)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(pitch, formGroup, probe)
	case "Pitched":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Pitched Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PitchedFormCallback(
			nil,
			probe,
			formGroup,
		)
		pitched := new(models.Pitched)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(pitched, formGroup, probe)
	case "Placement_text":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Placement_text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Placement_textFormCallback(
			nil,
			probe,
			formGroup,
		)
		placement_text := new(models.Placement_text)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(placement_text, formGroup, probe)
	case "Play":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Play Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlayFormCallback(
			nil,
			probe,
			formGroup,
		)
		play := new(models.Play)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(play, formGroup, probe)
	case "Player":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Player Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PlayerFormCallback(
			nil,
			probe,
			formGroup,
		)
		player := new(models.Player)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(player, formGroup, probe)
	case "Principal_voice":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Principal_voice Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Principal_voiceFormCallback(
			nil,
			probe,
			formGroup,
		)
		principal_voice := new(models.Principal_voice)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(principal_voice, formGroup, probe)
	case "Print":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Print Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__PrintFormCallback(
			nil,
			probe,
			formGroup,
		)
		print := new(models.Print)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(print, formGroup, probe)
	case "Release":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Release Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ReleaseFormCallback(
			nil,
			probe,
			formGroup,
		)
		release := new(models.Release)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(release, formGroup, probe)
	case "Repeat":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Repeat Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RepeatFormCallback(
			nil,
			probe,
			formGroup,
		)
		repeat := new(models.Repeat)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(repeat, formGroup, probe)
	case "Rest":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Rest Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RestFormCallback(
			nil,
			probe,
			formGroup,
		)
		rest := new(models.Rest)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(rest, formGroup, probe)
	case "Root":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Root Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__RootFormCallback(
			nil,
			probe,
			formGroup,
		)
		root := new(models.Root)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(root, formGroup, probe)
	case "Root_step":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Root_step Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Root_stepFormCallback(
			nil,
			probe,
			formGroup,
		)
		root_step := new(models.Root_step)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(root_step, formGroup, probe)
	case "Scaling":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Scaling Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ScalingFormCallback(
			nil,
			probe,
			formGroup,
		)
		scaling := new(models.Scaling)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(scaling, formGroup, probe)
	case "Scordatura":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Scordatura Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__ScordaturaFormCallback(
			nil,
			probe,
			formGroup,
		)
		scordatura := new(models.Scordatura)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(scordatura, formGroup, probe)
	case "Score_instrument":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Score_instrument Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Score_instrumentFormCallback(
			nil,
			probe,
			formGroup,
		)
		score_instrument := new(models.Score_instrument)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(score_instrument, formGroup, probe)
	case "Score_part":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Score_part Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Score_partFormCallback(
			nil,
			probe,
			formGroup,
		)
		score_part := new(models.Score_part)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(score_part, formGroup, probe)
	case "Score_partwise":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Score_partwise Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Score_partwiseFormCallback(
			nil,
			probe,
			formGroup,
		)
		score_partwise := new(models.Score_partwise)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(score_partwise, formGroup, probe)
	case "Score_timewise":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Score_timewise Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Score_timewiseFormCallback(
			nil,
			probe,
			formGroup,
		)
		score_timewise := new(models.Score_timewise)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(score_timewise, formGroup, probe)
	case "Segno":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Segno Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SegnoFormCallback(
			nil,
			probe,
			formGroup,
		)
		segno := new(models.Segno)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(segno, formGroup, probe)
	case "Slash":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Slash Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SlashFormCallback(
			nil,
			probe,
			formGroup,
		)
		slash := new(models.Slash)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(slash, formGroup, probe)
	case "Slide":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Slide Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SlideFormCallback(
			nil,
			probe,
			formGroup,
		)
		slide := new(models.Slide)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(slide, formGroup, probe)
	case "Slur":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Slur Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SlurFormCallback(
			nil,
			probe,
			formGroup,
		)
		slur := new(models.Slur)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(slur, formGroup, probe)
	case "Sound":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Sound Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SoundFormCallback(
			nil,
			probe,
			formGroup,
		)
		sound := new(models.Sound)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(sound, formGroup, probe)
	case "Staff_details":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Staff_details Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_detailsFormCallback(
			nil,
			probe,
			formGroup,
		)
		staff_details := new(models.Staff_details)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(staff_details, formGroup, probe)
	case "Staff_divide":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Staff_divide Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_divideFormCallback(
			nil,
			probe,
			formGroup,
		)
		staff_divide := new(models.Staff_divide)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(staff_divide, formGroup, probe)
	case "Staff_layout":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Staff_layout Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_layoutFormCallback(
			nil,
			probe,
			formGroup,
		)
		staff_layout := new(models.Staff_layout)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(staff_layout, formGroup, probe)
	case "Staff_size":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Staff_size Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_sizeFormCallback(
			nil,
			probe,
			formGroup,
		)
		staff_size := new(models.Staff_size)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(staff_size, formGroup, probe)
	case "Staff_tuning":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Staff_tuning Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Staff_tuningFormCallback(
			nil,
			probe,
			formGroup,
		)
		staff_tuning := new(models.Staff_tuning)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(staff_tuning, formGroup, probe)
	case "Stem":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Stem Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StemFormCallback(
			nil,
			probe,
			formGroup,
		)
		stem := new(models.Stem)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stem, formGroup, probe)
	case "Stick":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Stick Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__StickFormCallback(
			nil,
			probe,
			formGroup,
		)
		stick := new(models.Stick)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(stick, formGroup, probe)
	case "String_mute":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "String_mute Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__String_muteFormCallback(
			nil,
			probe,
			formGroup,
		)
		string_mute := new(models.String_mute)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(string_mute, formGroup, probe)
	case "String_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "String_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__String_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		string_type := new(models.String_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(string_type, formGroup, probe)
	case "Strong_accent":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Strong_accent Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Strong_accentFormCallback(
			nil,
			probe,
			formGroup,
		)
		strong_accent := new(models.Strong_accent)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(strong_accent, formGroup, probe)
	case "Style_text":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Style_text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Style_textFormCallback(
			nil,
			probe,
			formGroup,
		)
		style_text := new(models.Style_text)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(style_text, formGroup, probe)
	case "Supports":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Supports Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SupportsFormCallback(
			nil,
			probe,
			formGroup,
		)
		supports := new(models.Supports)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(supports, formGroup, probe)
	case "Swing":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Swing Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SwingFormCallback(
			nil,
			probe,
			formGroup,
		)
		swing := new(models.Swing)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(swing, formGroup, probe)
	case "Sync":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Sync Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__SyncFormCallback(
			nil,
			probe,
			formGroup,
		)
		sync := new(models.Sync)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(sync, formGroup, probe)
	case "System_dividers":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "System_dividers Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__System_dividersFormCallback(
			nil,
			probe,
			formGroup,
		)
		system_dividers := new(models.System_dividers)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(system_dividers, formGroup, probe)
	case "System_layout":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "System_layout Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__System_layoutFormCallback(
			nil,
			probe,
			formGroup,
		)
		system_layout := new(models.System_layout)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(system_layout, formGroup, probe)
	case "System_margins":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "System_margins Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__System_marginsFormCallback(
			nil,
			probe,
			formGroup,
		)
		system_margins := new(models.System_margins)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(system_margins, formGroup, probe)
	case "Tap":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Tap Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TapFormCallback(
			nil,
			probe,
			formGroup,
		)
		tap := new(models.Tap)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tap, formGroup, probe)
	case "Technical":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Technical Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TechnicalFormCallback(
			nil,
			probe,
			formGroup,
		)
		technical := new(models.Technical)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(technical, formGroup, probe)
	case "Text_element_data":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Text_element_data Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Text_element_dataFormCallback(
			nil,
			probe,
			formGroup,
		)
		text_element_data := new(models.Text_element_data)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(text_element_data, formGroup, probe)
	case "Tie":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Tie Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TieFormCallback(
			nil,
			probe,
			formGroup,
		)
		tie := new(models.Tie)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tie, formGroup, probe)
	case "Tied":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Tied Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TiedFormCallback(
			nil,
			probe,
			formGroup,
		)
		tied := new(models.Tied)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tied, formGroup, probe)
	case "Time":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Time Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TimeFormCallback(
			nil,
			probe,
			formGroup,
		)
		time := new(models.Time)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(time, formGroup, probe)
	case "Time_modification":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Time_modification Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Time_modificationFormCallback(
			nil,
			probe,
			formGroup,
		)
		time_modification := new(models.Time_modification)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(time_modification, formGroup, probe)
	case "Timpani":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Timpani Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TimpaniFormCallback(
			nil,
			probe,
			formGroup,
		)
		timpani := new(models.Timpani)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(timpani, formGroup, probe)
	case "Transpose":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Transpose Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TransposeFormCallback(
			nil,
			probe,
			formGroup,
		)
		transpose := new(models.Transpose)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(transpose, formGroup, probe)
	case "Tremolo":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Tremolo Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TremoloFormCallback(
			nil,
			probe,
			formGroup,
		)
		tremolo := new(models.Tremolo)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tremolo, formGroup, probe)
	case "Tuplet":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Tuplet Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__TupletFormCallback(
			nil,
			probe,
			formGroup,
		)
		tuplet := new(models.Tuplet)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tuplet, formGroup, probe)
	case "Tuplet_dot":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Tuplet_dot Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Tuplet_dotFormCallback(
			nil,
			probe,
			formGroup,
		)
		tuplet_dot := new(models.Tuplet_dot)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tuplet_dot, formGroup, probe)
	case "Tuplet_number":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Tuplet_number Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Tuplet_numberFormCallback(
			nil,
			probe,
			formGroup,
		)
		tuplet_number := new(models.Tuplet_number)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tuplet_number, formGroup, probe)
	case "Tuplet_portion":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Tuplet_portion Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Tuplet_portionFormCallback(
			nil,
			probe,
			formGroup,
		)
		tuplet_portion := new(models.Tuplet_portion)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tuplet_portion, formGroup, probe)
	case "Tuplet_type":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Tuplet_type Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Tuplet_typeFormCallback(
			nil,
			probe,
			formGroup,
		)
		tuplet_type := new(models.Tuplet_type)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(tuplet_type, formGroup, probe)
	case "Typed_text":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Typed_text Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Typed_textFormCallback(
			nil,
			probe,
			formGroup,
		)
		typed_text := new(models.Typed_text)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(typed_text, formGroup, probe)
	case "Unpitched":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Unpitched Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__UnpitchedFormCallback(
			nil,
			probe,
			formGroup,
		)
		unpitched := new(models.Unpitched)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(unpitched, formGroup, probe)
	case "Virtual_instrument":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Virtual_instrument Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Virtual_instrumentFormCallback(
			nil,
			probe,
			formGroup,
		)
		virtual_instrument := new(models.Virtual_instrument)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(virtual_instrument, formGroup, probe)
	case "Wait":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Wait Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WaitFormCallback(
			nil,
			probe,
			formGroup,
		)
		wait := new(models.Wait)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(wait, formGroup, probe)
	case "Wavy_line":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Wavy_line Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__Wavy_lineFormCallback(
			nil,
			probe,
			formGroup,
		)
		wavy_line := new(models.Wavy_line)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(wavy_line, formGroup, probe)
	case "Wedge":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Wedge Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WedgeFormCallback(
			nil,
			probe,
			formGroup,
		)
		wedge := new(models.Wedge)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(wedge, formGroup, probe)
	case "Wood":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Wood Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WoodFormCallback(
			nil,
			probe,
			formGroup,
		)
		wood := new(models.Wood)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(wood, formGroup, probe)
	case "Work":
		formGroup := (&form.FormGroup{
			Name:  FormName,
			Label: prefix + "Work Form",
		}).Stage(formStage)
		formGroup.OnSave = __gong__New__WorkFormCallback(
			nil,
			probe,
			formGroup,
		)
		work := new(models.Work)
		formGroup.HasSuppressButton = !isNewInstance
		FillUpForm(work, formGroup, probe)
	}
	formStage.Commit()
}
