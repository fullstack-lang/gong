// generated code - do not edit
package models

import (
	"crypto/sha256"
	"encoding/binary"
	"fmt"
	"log"
	"sort"
	"strings"
	"time"
)

var (
	__GongSliceTemplate_time__dummyDeclaration time.Duration
	_                                          = __GongSliceTemplate_time__dummyDeclaration
)

// ComputeReverseMaps computes the reverse map, for all intances, for all slice to pointers field
// Its complexity is in O(n)O(p) where p is the number of pointers
func (stage *Stage) ComputeReverseMaps() {
	// insertion point per named struct
	// Compute reverse map for named struct A_directive
	// insertion point per field

	// Compute reverse map for named struct A_measure
	// insertion point per field
	stage.A_measure_Note_reverseMap = make(map[*Note]*A_measure)
	for a_measure := range stage.A_measures {
		_ = a_measure
		for _, _note := range a_measure.Note {
			stage.A_measure_Note_reverseMap[_note] = a_measure
		}
	}
	stage.A_measure_Backup_reverseMap = make(map[*Backup]*A_measure)
	for a_measure := range stage.A_measures {
		_ = a_measure
		for _, _backup := range a_measure.Backup {
			stage.A_measure_Backup_reverseMap[_backup] = a_measure
		}
	}
	stage.A_measure_Forward_reverseMap = make(map[*Forward]*A_measure)
	for a_measure := range stage.A_measures {
		_ = a_measure
		for _, _forward := range a_measure.Forward {
			stage.A_measure_Forward_reverseMap[_forward] = a_measure
		}
	}
	stage.A_measure_Direction_reverseMap = make(map[*Direction]*A_measure)
	for a_measure := range stage.A_measures {
		_ = a_measure
		for _, _direction := range a_measure.Direction {
			stage.A_measure_Direction_reverseMap[_direction] = a_measure
		}
	}
	stage.A_measure_Attributes_reverseMap = make(map[*Attributes]*A_measure)
	for a_measure := range stage.A_measures {
		_ = a_measure
		for _, _attributes := range a_measure.Attributes {
			stage.A_measure_Attributes_reverseMap[_attributes] = a_measure
		}
	}
	stage.A_measure_Harmony_reverseMap = make(map[*Harmony]*A_measure)
	for a_measure := range stage.A_measures {
		_ = a_measure
		for _, _harmony := range a_measure.Harmony {
			stage.A_measure_Harmony_reverseMap[_harmony] = a_measure
		}
	}
	stage.A_measure_Figured_bass_reverseMap = make(map[*Figured_bass]*A_measure)
	for a_measure := range stage.A_measures {
		_ = a_measure
		for _, _figured_bass := range a_measure.Figured_bass {
			stage.A_measure_Figured_bass_reverseMap[_figured_bass] = a_measure
		}
	}
	stage.A_measure_Print_reverseMap = make(map[*Print]*A_measure)
	for a_measure := range stage.A_measures {
		_ = a_measure
		for _, _print := range a_measure.Print {
			stage.A_measure_Print_reverseMap[_print] = a_measure
		}
	}
	stage.A_measure_Sound_reverseMap = make(map[*Sound]*A_measure)
	for a_measure := range stage.A_measures {
		_ = a_measure
		for _, _sound := range a_measure.Sound {
			stage.A_measure_Sound_reverseMap[_sound] = a_measure
		}
	}
	stage.A_measure_Listening_reverseMap = make(map[*Listening]*A_measure)
	for a_measure := range stage.A_measures {
		_ = a_measure
		for _, _listening := range a_measure.Listening {
			stage.A_measure_Listening_reverseMap[_listening] = a_measure
		}
	}
	stage.A_measure_Barline_reverseMap = make(map[*Barline]*A_measure)
	for a_measure := range stage.A_measures {
		_ = a_measure
		for _, _barline := range a_measure.Barline {
			stage.A_measure_Barline_reverseMap[_barline] = a_measure
		}
	}
	stage.A_measure_Grouping_reverseMap = make(map[*Grouping]*A_measure)
	for a_measure := range stage.A_measures {
		_ = a_measure
		for _, _grouping := range a_measure.Grouping {
			stage.A_measure_Grouping_reverseMap[_grouping] = a_measure
		}
	}
	stage.A_measure_Link_reverseMap = make(map[*Link]*A_measure)
	for a_measure := range stage.A_measures {
		_ = a_measure
		for _, _link := range a_measure.Link {
			stage.A_measure_Link_reverseMap[_link] = a_measure
		}
	}
	stage.A_measure_Bookmark_reverseMap = make(map[*Bookmark]*A_measure)
	for a_measure := range stage.A_measures {
		_ = a_measure
		for _, _bookmark := range a_measure.Bookmark {
			stage.A_measure_Bookmark_reverseMap[_bookmark] = a_measure
		}
	}

	// Compute reverse map for named struct A_measure_1
	// insertion point per field
	stage.A_measure_1_Part_reverseMap = make(map[*A_part_1]*A_measure_1)
	for a_measure_1 := range stage.A_measure_1s {
		_ = a_measure_1
		for _, _a_part_1 := range a_measure_1.Part {
			stage.A_measure_1_Part_reverseMap[_a_part_1] = a_measure_1
		}
	}

	// Compute reverse map for named struct A_part
	// insertion point per field
	stage.A_part_Measure_reverseMap = make(map[*A_measure]*A_part)
	for a_part := range stage.A_parts {
		_ = a_part
		for _, _a_measure := range a_part.Measure {
			stage.A_part_Measure_reverseMap[_a_measure] = a_part
		}
	}

	// Compute reverse map for named struct A_part_1
	// insertion point per field
	stage.A_part_1_Note_reverseMap = make(map[*Note]*A_part_1)
	for a_part_1 := range stage.A_part_1s {
		_ = a_part_1
		for _, _note := range a_part_1.Note {
			stage.A_part_1_Note_reverseMap[_note] = a_part_1
		}
	}
	stage.A_part_1_Backup_reverseMap = make(map[*Backup]*A_part_1)
	for a_part_1 := range stage.A_part_1s {
		_ = a_part_1
		for _, _backup := range a_part_1.Backup {
			stage.A_part_1_Backup_reverseMap[_backup] = a_part_1
		}
	}
	stage.A_part_1_Forward_reverseMap = make(map[*Forward]*A_part_1)
	for a_part_1 := range stage.A_part_1s {
		_ = a_part_1
		for _, _forward := range a_part_1.Forward {
			stage.A_part_1_Forward_reverseMap[_forward] = a_part_1
		}
	}
	stage.A_part_1_Direction_reverseMap = make(map[*Direction]*A_part_1)
	for a_part_1 := range stage.A_part_1s {
		_ = a_part_1
		for _, _direction := range a_part_1.Direction {
			stage.A_part_1_Direction_reverseMap[_direction] = a_part_1
		}
	}
	stage.A_part_1_Attributes_reverseMap = make(map[*Attributes]*A_part_1)
	for a_part_1 := range stage.A_part_1s {
		_ = a_part_1
		for _, _attributes := range a_part_1.Attributes {
			stage.A_part_1_Attributes_reverseMap[_attributes] = a_part_1
		}
	}
	stage.A_part_1_Harmony_reverseMap = make(map[*Harmony]*A_part_1)
	for a_part_1 := range stage.A_part_1s {
		_ = a_part_1
		for _, _harmony := range a_part_1.Harmony {
			stage.A_part_1_Harmony_reverseMap[_harmony] = a_part_1
		}
	}
	stage.A_part_1_Figured_bass_reverseMap = make(map[*Figured_bass]*A_part_1)
	for a_part_1 := range stage.A_part_1s {
		_ = a_part_1
		for _, _figured_bass := range a_part_1.Figured_bass {
			stage.A_part_1_Figured_bass_reverseMap[_figured_bass] = a_part_1
		}
	}
	stage.A_part_1_Print_reverseMap = make(map[*Print]*A_part_1)
	for a_part_1 := range stage.A_part_1s {
		_ = a_part_1
		for _, _print := range a_part_1.Print {
			stage.A_part_1_Print_reverseMap[_print] = a_part_1
		}
	}
	stage.A_part_1_Sound_reverseMap = make(map[*Sound]*A_part_1)
	for a_part_1 := range stage.A_part_1s {
		_ = a_part_1
		for _, _sound := range a_part_1.Sound {
			stage.A_part_1_Sound_reverseMap[_sound] = a_part_1
		}
	}
	stage.A_part_1_Listening_reverseMap = make(map[*Listening]*A_part_1)
	for a_part_1 := range stage.A_part_1s {
		_ = a_part_1
		for _, _listening := range a_part_1.Listening {
			stage.A_part_1_Listening_reverseMap[_listening] = a_part_1
		}
	}
	stage.A_part_1_Barline_reverseMap = make(map[*Barline]*A_part_1)
	for a_part_1 := range stage.A_part_1s {
		_ = a_part_1
		for _, _barline := range a_part_1.Barline {
			stage.A_part_1_Barline_reverseMap[_barline] = a_part_1
		}
	}
	stage.A_part_1_Grouping_reverseMap = make(map[*Grouping]*A_part_1)
	for a_part_1 := range stage.A_part_1s {
		_ = a_part_1
		for _, _grouping := range a_part_1.Grouping {
			stage.A_part_1_Grouping_reverseMap[_grouping] = a_part_1
		}
	}
	stage.A_part_1_Link_reverseMap = make(map[*Link]*A_part_1)
	for a_part_1 := range stage.A_part_1s {
		_ = a_part_1
		for _, _link := range a_part_1.Link {
			stage.A_part_1_Link_reverseMap[_link] = a_part_1
		}
	}
	stage.A_part_1_Bookmark_reverseMap = make(map[*Bookmark]*A_part_1)
	for a_part_1 := range stage.A_part_1s {
		_ = a_part_1
		for _, _bookmark := range a_part_1.Bookmark {
			stage.A_part_1_Bookmark_reverseMap[_bookmark] = a_part_1
		}
	}

	// Compute reverse map for named struct Accidental
	// insertion point per field

	// Compute reverse map for named struct Accidental_mark
	// insertion point per field

	// Compute reverse map for named struct Accidental_text
	// insertion point per field

	// Compute reverse map for named struct Accord
	// insertion point per field

	// Compute reverse map for named struct Accordion_registration
	// insertion point per field

	// Compute reverse map for named struct Appearance
	// insertion point per field
	stage.Appearance_Line_width_reverseMap = make(map[*Line_width]*Appearance)
	for appearance := range stage.Appearances {
		_ = appearance
		for _, _line_width := range appearance.Line_width {
			stage.Appearance_Line_width_reverseMap[_line_width] = appearance
		}
	}
	stage.Appearance_Note_size_reverseMap = make(map[*Note_size]*Appearance)
	for appearance := range stage.Appearances {
		_ = appearance
		for _, _note_size := range appearance.Note_size {
			stage.Appearance_Note_size_reverseMap[_note_size] = appearance
		}
	}
	stage.Appearance_Distance_reverseMap = make(map[*Distance]*Appearance)
	for appearance := range stage.Appearances {
		_ = appearance
		for _, _distance := range appearance.Distance {
			stage.Appearance_Distance_reverseMap[_distance] = appearance
		}
	}
	stage.Appearance_Glyph_reverseMap = make(map[*Glyph]*Appearance)
	for appearance := range stage.Appearances {
		_ = appearance
		for _, _glyph := range appearance.Glyph {
			stage.Appearance_Glyph_reverseMap[_glyph] = appearance
		}
	}
	stage.Appearance_Other_appearance_reverseMap = make(map[*Other_appearance]*Appearance)
	for appearance := range stage.Appearances {
		_ = appearance
		for _, _other_appearance := range appearance.Other_appearance {
			stage.Appearance_Other_appearance_reverseMap[_other_appearance] = appearance
		}
	}

	// Compute reverse map for named struct Arpeggiate
	// insertion point per field

	// Compute reverse map for named struct Arrow
	// insertion point per field

	// Compute reverse map for named struct Articulations
	// insertion point per field
	stage.Articulations_Accent_reverseMap = make(map[*Empty_placement]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _empty_placement := range articulations.Accent {
			stage.Articulations_Accent_reverseMap[_empty_placement] = articulations
		}
	}
	stage.Articulations_Strong_accent_reverseMap = make(map[*Strong_accent]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _strong_accent := range articulations.Strong_accent {
			stage.Articulations_Strong_accent_reverseMap[_strong_accent] = articulations
		}
	}
	stage.Articulations_Staccato_reverseMap = make(map[*Empty_placement]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _empty_placement := range articulations.Staccato {
			stage.Articulations_Staccato_reverseMap[_empty_placement] = articulations
		}
	}
	stage.Articulations_Tenuto_reverseMap = make(map[*Empty_placement]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _empty_placement := range articulations.Tenuto {
			stage.Articulations_Tenuto_reverseMap[_empty_placement] = articulations
		}
	}
	stage.Articulations_Detached_legato_reverseMap = make(map[*Empty_placement]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _empty_placement := range articulations.Detached_legato {
			stage.Articulations_Detached_legato_reverseMap[_empty_placement] = articulations
		}
	}
	stage.Articulations_Staccatissimo_reverseMap = make(map[*Empty_placement]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _empty_placement := range articulations.Staccatissimo {
			stage.Articulations_Staccatissimo_reverseMap[_empty_placement] = articulations
		}
	}
	stage.Articulations_Spiccato_reverseMap = make(map[*Empty_placement]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _empty_placement := range articulations.Spiccato {
			stage.Articulations_Spiccato_reverseMap[_empty_placement] = articulations
		}
	}
	stage.Articulations_Scoop_reverseMap = make(map[*Empty_line]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _empty_line := range articulations.Scoop {
			stage.Articulations_Scoop_reverseMap[_empty_line] = articulations
		}
	}
	stage.Articulations_Plop_reverseMap = make(map[*Empty_line]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _empty_line := range articulations.Plop {
			stage.Articulations_Plop_reverseMap[_empty_line] = articulations
		}
	}
	stage.Articulations_Doit_reverseMap = make(map[*Empty_line]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _empty_line := range articulations.Doit {
			stage.Articulations_Doit_reverseMap[_empty_line] = articulations
		}
	}
	stage.Articulations_Falloff_reverseMap = make(map[*Empty_line]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _empty_line := range articulations.Falloff {
			stage.Articulations_Falloff_reverseMap[_empty_line] = articulations
		}
	}
	stage.Articulations_Breath_mark_reverseMap = make(map[*Breath_mark]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _breath_mark := range articulations.Breath_mark {
			stage.Articulations_Breath_mark_reverseMap[_breath_mark] = articulations
		}
	}
	stage.Articulations_Caesura_reverseMap = make(map[*Caesura]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _caesura := range articulations.Caesura {
			stage.Articulations_Caesura_reverseMap[_caesura] = articulations
		}
	}
	stage.Articulations_Stress_reverseMap = make(map[*Empty_placement]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _empty_placement := range articulations.Stress {
			stage.Articulations_Stress_reverseMap[_empty_placement] = articulations
		}
	}
	stage.Articulations_Unstress_reverseMap = make(map[*Empty_placement]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _empty_placement := range articulations.Unstress {
			stage.Articulations_Unstress_reverseMap[_empty_placement] = articulations
		}
	}
	stage.Articulations_Soft_accent_reverseMap = make(map[*Empty_placement]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _empty_placement := range articulations.Soft_accent {
			stage.Articulations_Soft_accent_reverseMap[_empty_placement] = articulations
		}
	}
	stage.Articulations_Other_articulation_reverseMap = make(map[*Other_placement_text]*Articulations)
	for articulations := range stage.Articulationss {
		_ = articulations
		for _, _other_placement_text := range articulations.Other_articulation {
			stage.Articulations_Other_articulation_reverseMap[_other_placement_text] = articulations
		}
	}

	// Compute reverse map for named struct Assess
	// insertion point per field

	// Compute reverse map for named struct Attributes
	// insertion point per field
	stage.Attributes_Key_reverseMap = make(map[*Key]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _key := range attributes.Key {
			stage.Attributes_Key_reverseMap[_key] = attributes
		}
	}
	stage.Attributes_Time_reverseMap = make(map[*Time]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _time := range attributes.Time {
			stage.Attributes_Time_reverseMap[_time] = attributes
		}
	}
	stage.Attributes_Clef_reverseMap = make(map[*Clef]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _clef := range attributes.Clef {
			stage.Attributes_Clef_reverseMap[_clef] = attributes
		}
	}
	stage.Attributes_Staff_details_reverseMap = make(map[*Staff_details]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _staff_details := range attributes.Staff_details {
			stage.Attributes_Staff_details_reverseMap[_staff_details] = attributes
		}
	}
	stage.Attributes_Transpose_reverseMap = make(map[*Transpose]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _transpose := range attributes.Transpose {
			stage.Attributes_Transpose_reverseMap[_transpose] = attributes
		}
	}
	stage.Attributes_For_part_reverseMap = make(map[*For_part]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _for_part := range attributes.For_part {
			stage.Attributes_For_part_reverseMap[_for_part] = attributes
		}
	}
	stage.Attributes_Directive_reverseMap = make(map[*A_directive]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _a_directive := range attributes.Directive {
			stage.Attributes_Directive_reverseMap[_a_directive] = attributes
		}
	}
	stage.Attributes_Measure_style_reverseMap = make(map[*Measure_style]*Attributes)
	for attributes := range stage.Attributess {
		_ = attributes
		for _, _measure_style := range attributes.Measure_style {
			stage.Attributes_Measure_style_reverseMap[_measure_style] = attributes
		}
	}

	// Compute reverse map for named struct Backup
	// insertion point per field

	// Compute reverse map for named struct Bar_style_color
	// insertion point per field

	// Compute reverse map for named struct Barline
	// insertion point per field

	// Compute reverse map for named struct Barre
	// insertion point per field

	// Compute reverse map for named struct Bass
	// insertion point per field

	// Compute reverse map for named struct Bass_step
	// insertion point per field

	// Compute reverse map for named struct Beam
	// insertion point per field

	// Compute reverse map for named struct Beat_repeat
	// insertion point per field

	// Compute reverse map for named struct Beat_unit_tied
	// insertion point per field

	// Compute reverse map for named struct Beater
	// insertion point per field

	// Compute reverse map for named struct Bend
	// insertion point per field

	// Compute reverse map for named struct Bookmark
	// insertion point per field

	// Compute reverse map for named struct Bracket
	// insertion point per field

	// Compute reverse map for named struct Breath_mark
	// insertion point per field

	// Compute reverse map for named struct Caesura
	// insertion point per field

	// Compute reverse map for named struct Cancel
	// insertion point per field

	// Compute reverse map for named struct Clef
	// insertion point per field

	// Compute reverse map for named struct Coda
	// insertion point per field

	// Compute reverse map for named struct Credit
	// insertion point per field
	stage.Credit_Link_reverseMap = make(map[*Link]*Credit)
	for credit := range stage.Credits {
		_ = credit
		for _, _link := range credit.Link {
			stage.Credit_Link_reverseMap[_link] = credit
		}
	}
	stage.Credit_Bookmark_reverseMap = make(map[*Bookmark]*Credit)
	for credit := range stage.Credits {
		_ = credit
		for _, _bookmark := range credit.Bookmark {
			stage.Credit_Bookmark_reverseMap[_bookmark] = credit
		}
	}
	stage.Credit_Credit_words_reverseMap = make(map[*Formatted_text_id]*Credit)
	for credit := range stage.Credits {
		_ = credit
		for _, _formatted_text_id := range credit.Credit_words {
			stage.Credit_Credit_words_reverseMap[_formatted_text_id] = credit
		}
	}
	stage.Credit_Credit_symbol_reverseMap = make(map[*Formatted_symbol_id]*Credit)
	for credit := range stage.Credits {
		_ = credit
		for _, _formatted_symbol_id := range credit.Credit_symbol {
			stage.Credit_Credit_symbol_reverseMap[_formatted_symbol_id] = credit
		}
	}

	// Compute reverse map for named struct Dashes
	// insertion point per field

	// Compute reverse map for named struct Defaults
	// insertion point per field
	stage.Defaults_Staff_layout_reverseMap = make(map[*Staff_layout]*Defaults)
	for defaults := range stage.Defaultss {
		_ = defaults
		for _, _staff_layout := range defaults.Staff_layout {
			stage.Defaults_Staff_layout_reverseMap[_staff_layout] = defaults
		}
	}
	stage.Defaults_Lyric_font_reverseMap = make(map[*Lyric_font]*Defaults)
	for defaults := range stage.Defaultss {
		_ = defaults
		for _, _lyric_font := range defaults.Lyric_font {
			stage.Defaults_Lyric_font_reverseMap[_lyric_font] = defaults
		}
	}
	stage.Defaults_Lyric_language_reverseMap = make(map[*Lyric_language]*Defaults)
	for defaults := range stage.Defaultss {
		_ = defaults
		for _, _lyric_language := range defaults.Lyric_language {
			stage.Defaults_Lyric_language_reverseMap[_lyric_language] = defaults
		}
	}

	// Compute reverse map for named struct Degree
	// insertion point per field

	// Compute reverse map for named struct Degree_alter
	// insertion point per field

	// Compute reverse map for named struct Degree_type
	// insertion point per field

	// Compute reverse map for named struct Degree_value
	// insertion point per field

	// Compute reverse map for named struct Direction
	// insertion point per field
	stage.Direction_Direction_type_reverseMap = make(map[*Direction_type]*Direction)
	for direction := range stage.Directions {
		_ = direction
		for _, _direction_type := range direction.Direction_type {
			stage.Direction_Direction_type_reverseMap[_direction_type] = direction
		}
	}

	// Compute reverse map for named struct Direction_type
	// insertion point per field
	stage.Direction_type_Rehearsal_reverseMap = make(map[*Formatted_text_id]*Direction_type)
	for direction_type := range stage.Direction_types {
		_ = direction_type
		for _, _formatted_text_id := range direction_type.Rehearsal {
			stage.Direction_type_Rehearsal_reverseMap[_formatted_text_id] = direction_type
		}
	}
	stage.Direction_type_Segno_reverseMap = make(map[*Segno]*Direction_type)
	for direction_type := range stage.Direction_types {
		_ = direction_type
		for _, _segno := range direction_type.Segno {
			stage.Direction_type_Segno_reverseMap[_segno] = direction_type
		}
	}
	stage.Direction_type_Coda_reverseMap = make(map[*Coda]*Direction_type)
	for direction_type := range stage.Direction_types {
		_ = direction_type
		for _, _coda := range direction_type.Coda {
			stage.Direction_type_Coda_reverseMap[_coda] = direction_type
		}
	}
	stage.Direction_type_Words_reverseMap = make(map[*Formatted_text_id]*Direction_type)
	for direction_type := range stage.Direction_types {
		_ = direction_type
		for _, _formatted_text_id := range direction_type.Words {
			stage.Direction_type_Words_reverseMap[_formatted_text_id] = direction_type
		}
	}
	stage.Direction_type_Symbol_reverseMap = make(map[*Formatted_symbol_id]*Direction_type)
	for direction_type := range stage.Direction_types {
		_ = direction_type
		for _, _formatted_symbol_id := range direction_type.Symbol {
			stage.Direction_type_Symbol_reverseMap[_formatted_symbol_id] = direction_type
		}
	}
	stage.Direction_type_Dynamics_reverseMap = make(map[*Dynamics]*Direction_type)
	for direction_type := range stage.Direction_types {
		_ = direction_type
		for _, _dynamics := range direction_type.Dynamics {
			stage.Direction_type_Dynamics_reverseMap[_dynamics] = direction_type
		}
	}
	stage.Direction_type_Percussion_reverseMap = make(map[*Percussion]*Direction_type)
	for direction_type := range stage.Direction_types {
		_ = direction_type
		for _, _percussion := range direction_type.Percussion {
			stage.Direction_type_Percussion_reverseMap[_percussion] = direction_type
		}
	}

	// Compute reverse map for named struct Distance
	// insertion point per field

	// Compute reverse map for named struct Double
	// insertion point per field

	// Compute reverse map for named struct Dynamics
	// insertion point per field
	stage.Dynamics_Other_dynamics_reverseMap = make(map[*Other_text]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _other_text := range dynamics.Other_dynamics {
			stage.Dynamics_Other_dynamics_reverseMap[_other_text] = dynamics
		}
	}

	// Compute reverse map for named struct Effect
	// insertion point per field

	// Compute reverse map for named struct Elision
	// insertion point per field

	// Compute reverse map for named struct Empty
	// insertion point per field

	// Compute reverse map for named struct Empty_font
	// insertion point per field

	// Compute reverse map for named struct Empty_line
	// insertion point per field

	// Compute reverse map for named struct Empty_placement
	// insertion point per field

	// Compute reverse map for named struct Empty_placement_smufl
	// insertion point per field

	// Compute reverse map for named struct Empty_print_object_style_align
	// insertion point per field

	// Compute reverse map for named struct Empty_print_style
	// insertion point per field

	// Compute reverse map for named struct Empty_print_style_align
	// insertion point per field

	// Compute reverse map for named struct Empty_print_style_align_id
	// insertion point per field

	// Compute reverse map for named struct Empty_trill_sound
	// insertion point per field

	// Compute reverse map for named struct Encoding
	// insertion point per field
	stage.Encoding_Encoder_reverseMap = make(map[*Typed_text]*Encoding)
	for encoding := range stage.Encodings {
		_ = encoding
		for _, _typed_text := range encoding.Encoder {
			stage.Encoding_Encoder_reverseMap[_typed_text] = encoding
		}
	}
	stage.Encoding_Supports_reverseMap = make(map[*Supports]*Encoding)
	for encoding := range stage.Encodings {
		_ = encoding
		for _, _supports := range encoding.Supports {
			stage.Encoding_Supports_reverseMap[_supports] = encoding
		}
	}

	// Compute reverse map for named struct Ending
	// insertion point per field

	// Compute reverse map for named struct Extend
	// insertion point per field

	// Compute reverse map for named struct Feature
	// insertion point per field

	// Compute reverse map for named struct Fermata
	// insertion point per field

	// Compute reverse map for named struct Figure
	// insertion point per field

	// Compute reverse map for named struct Figured_bass
	// insertion point per field
	stage.Figured_bass_Figure_reverseMap = make(map[*Figure]*Figured_bass)
	for figured_bass := range stage.Figured_basss {
		_ = figured_bass
		for _, _figure := range figured_bass.Figure {
			stage.Figured_bass_Figure_reverseMap[_figure] = figured_bass
		}
	}

	// Compute reverse map for named struct Fingering
	// insertion point per field

	// Compute reverse map for named struct First_fret
	// insertion point per field

	// Compute reverse map for named struct For_part
	// insertion point per field

	// Compute reverse map for named struct Formatted_symbol
	// insertion point per field

	// Compute reverse map for named struct Formatted_symbol_id
	// insertion point per field

	// Compute reverse map for named struct Formatted_text
	// insertion point per field

	// Compute reverse map for named struct Formatted_text_id
	// insertion point per field

	// Compute reverse map for named struct Forward
	// insertion point per field

	// Compute reverse map for named struct Frame
	// insertion point per field
	stage.Frame_Frame_note_reverseMap = make(map[*Frame_note]*Frame)
	for frame := range stage.Frames {
		_ = frame
		for _, _frame_note := range frame.Frame_note {
			stage.Frame_Frame_note_reverseMap[_frame_note] = frame
		}
	}

	// Compute reverse map for named struct Frame_note
	// insertion point per field

	// Compute reverse map for named struct Fret
	// insertion point per field

	// Compute reverse map for named struct Glass
	// insertion point per field

	// Compute reverse map for named struct Glissando
	// insertion point per field

	// Compute reverse map for named struct Glyph
	// insertion point per field

	// Compute reverse map for named struct Grace
	// insertion point per field

	// Compute reverse map for named struct Group_barline
	// insertion point per field

	// Compute reverse map for named struct Group_name
	// insertion point per field

	// Compute reverse map for named struct Group_symbol
	// insertion point per field

	// Compute reverse map for named struct Grouping
	// insertion point per field
	stage.Grouping_Feature_reverseMap = make(map[*Feature]*Grouping)
	for grouping := range stage.Groupings {
		_ = grouping
		for _, _feature := range grouping.Feature {
			stage.Grouping_Feature_reverseMap[_feature] = grouping
		}
	}

	// Compute reverse map for named struct Hammer_on_pull_off
	// insertion point per field

	// Compute reverse map for named struct Handbell
	// insertion point per field

	// Compute reverse map for named struct Harmon_closed
	// insertion point per field

	// Compute reverse map for named struct Harmon_mute
	// insertion point per field

	// Compute reverse map for named struct Harmonic
	// insertion point per field

	// Compute reverse map for named struct Harmony
	// insertion point per field
	stage.Harmony_Degree_reverseMap = make(map[*Degree]*Harmony)
	for harmony := range stage.Harmonys {
		_ = harmony
		for _, _degree := range harmony.Degree {
			stage.Harmony_Degree_reverseMap[_degree] = harmony
		}
	}

	// Compute reverse map for named struct Harmony_alter
	// insertion point per field

	// Compute reverse map for named struct Harp_pedals
	// insertion point per field
	stage.Harp_pedals_Pedal_tuning_reverseMap = make(map[*Pedal_tuning]*Harp_pedals)
	for harp_pedals := range stage.Harp_pedalss {
		_ = harp_pedals
		for _, _pedal_tuning := range harp_pedals.Pedal_tuning {
			stage.Harp_pedals_Pedal_tuning_reverseMap[_pedal_tuning] = harp_pedals
		}
	}

	// Compute reverse map for named struct Heel_toe
	// insertion point per field

	// Compute reverse map for named struct Hole
	// insertion point per field

	// Compute reverse map for named struct Hole_closed
	// insertion point per field

	// Compute reverse map for named struct Horizontal_turn
	// insertion point per field

	// Compute reverse map for named struct Identification
	// insertion point per field
	stage.Identification_Creator_reverseMap = make(map[*Typed_text]*Identification)
	for identification := range stage.Identifications {
		_ = identification
		for _, _typed_text := range identification.Creator {
			stage.Identification_Creator_reverseMap[_typed_text] = identification
		}
	}
	stage.Identification_Rights_reverseMap = make(map[*Typed_text]*Identification)
	for identification := range stage.Identifications {
		_ = identification
		for _, _typed_text := range identification.Rights {
			stage.Identification_Rights_reverseMap[_typed_text] = identification
		}
	}
	stage.Identification_Relation_reverseMap = make(map[*Typed_text]*Identification)
	for identification := range stage.Identifications {
		_ = identification
		for _, _typed_text := range identification.Relation {
			stage.Identification_Relation_reverseMap[_typed_text] = identification
		}
	}

	// Compute reverse map for named struct Image
	// insertion point per field

	// Compute reverse map for named struct Instrument
	// insertion point per field

	// Compute reverse map for named struct Instrument_change
	// insertion point per field

	// Compute reverse map for named struct Instrument_link
	// insertion point per field

	// Compute reverse map for named struct Interchangeable
	// insertion point per field

	// Compute reverse map for named struct Inversion
	// insertion point per field

	// Compute reverse map for named struct Key
	// insertion point per field
	stage.Key_Key_octave_reverseMap = make(map[*Key_octave]*Key)
	for key := range stage.Keys {
		_ = key
		for _, _key_octave := range key.Key_octave {
			stage.Key_Key_octave_reverseMap[_key_octave] = key
		}
	}

	// Compute reverse map for named struct Key_accidental
	// insertion point per field

	// Compute reverse map for named struct Key_octave
	// insertion point per field

	// Compute reverse map for named struct Kind
	// insertion point per field

	// Compute reverse map for named struct Level
	// insertion point per field

	// Compute reverse map for named struct Line_detail
	// insertion point per field

	// Compute reverse map for named struct Line_width
	// insertion point per field

	// Compute reverse map for named struct Link
	// insertion point per field

	// Compute reverse map for named struct Listen
	// insertion point per field
	stage.Listen_Assess_reverseMap = make(map[*Assess]*Listen)
	for listen := range stage.Listens {
		_ = listen
		for _, _assess := range listen.Assess {
			stage.Listen_Assess_reverseMap[_assess] = listen
		}
	}
	stage.Listen_Wait_reverseMap = make(map[*Wait]*Listen)
	for listen := range stage.Listens {
		_ = listen
		for _, _wait := range listen.Wait {
			stage.Listen_Wait_reverseMap[_wait] = listen
		}
	}
	stage.Listen_Other_listen_reverseMap = make(map[*Other_listening]*Listen)
	for listen := range stage.Listens {
		_ = listen
		for _, _other_listening := range listen.Other_listen {
			stage.Listen_Other_listen_reverseMap[_other_listening] = listen
		}
	}

	// Compute reverse map for named struct Listening
	// insertion point per field
	stage.Listening_Sync_reverseMap = make(map[*Sync]*Listening)
	for listening := range stage.Listenings {
		_ = listening
		for _, _sync := range listening.Sync {
			stage.Listening_Sync_reverseMap[_sync] = listening
		}
	}
	stage.Listening_Other_listening_reverseMap = make(map[*Other_listening]*Listening)
	for listening := range stage.Listenings {
		_ = listening
		for _, _other_listening := range listening.Other_listening {
			stage.Listening_Other_listening_reverseMap[_other_listening] = listening
		}
	}

	// Compute reverse map for named struct Lyric
	// insertion point per field
	stage.Lyric_Elision_reverseMap = make(map[*Elision]*Lyric)
	for lyric := range stage.Lyrics {
		_ = lyric
		for _, _elision := range lyric.Elision {
			stage.Lyric_Elision_reverseMap[_elision] = lyric
		}
	}
	stage.Lyric_Text_reverseMap = make(map[*Text_element_data]*Lyric)
	for lyric := range stage.Lyrics {
		_ = lyric
		for _, _text_element_data := range lyric.Text {
			stage.Lyric_Text_reverseMap[_text_element_data] = lyric
		}
	}

	// Compute reverse map for named struct Lyric_font
	// insertion point per field

	// Compute reverse map for named struct Lyric_language
	// insertion point per field

	// Compute reverse map for named struct Measure_layout
	// insertion point per field

	// Compute reverse map for named struct Measure_numbering
	// insertion point per field

	// Compute reverse map for named struct Measure_repeat
	// insertion point per field

	// Compute reverse map for named struct Measure_style
	// insertion point per field

	// Compute reverse map for named struct Membrane
	// insertion point per field

	// Compute reverse map for named struct Metal
	// insertion point per field

	// Compute reverse map for named struct Metronome
	// insertion point per field
	stage.Metronome_Beat_unit_tied_reverseMap = make(map[*Beat_unit_tied]*Metronome)
	for metronome := range stage.Metronomes {
		_ = metronome
		for _, _beat_unit_tied := range metronome.Beat_unit_tied {
			stage.Metronome_Beat_unit_tied_reverseMap[_beat_unit_tied] = metronome
		}
	}
	stage.Metronome_Metronome_note_reverseMap = make(map[*Metronome_note]*Metronome)
	for metronome := range stage.Metronomes {
		_ = metronome
		for _, _metronome_note := range metronome.Metronome_note {
			stage.Metronome_Metronome_note_reverseMap[_metronome_note] = metronome
		}
	}

	// Compute reverse map for named struct Metronome_beam
	// insertion point per field

	// Compute reverse map for named struct Metronome_note
	// insertion point per field
	stage.Metronome_note_Metronome_beam_reverseMap = make(map[*Metronome_beam]*Metronome_note)
	for metronome_note := range stage.Metronome_notes {
		_ = metronome_note
		for _, _metronome_beam := range metronome_note.Metronome_beam {
			stage.Metronome_note_Metronome_beam_reverseMap[_metronome_beam] = metronome_note
		}
	}

	// Compute reverse map for named struct Metronome_tied
	// insertion point per field

	// Compute reverse map for named struct Metronome_tuplet
	// insertion point per field

	// Compute reverse map for named struct Midi_device
	// insertion point per field

	// Compute reverse map for named struct Midi_instrument
	// insertion point per field

	// Compute reverse map for named struct Miscellaneous
	// insertion point per field
	stage.Miscellaneous_Miscellaneous_field_reverseMap = make(map[*Miscellaneous_field]*Miscellaneous)
	for miscellaneous := range stage.Miscellaneouss {
		_ = miscellaneous
		for _, _miscellaneous_field := range miscellaneous.Miscellaneous_field {
			stage.Miscellaneous_Miscellaneous_field_reverseMap[_miscellaneous_field] = miscellaneous
		}
	}

	// Compute reverse map for named struct Miscellaneous_field
	// insertion point per field

	// Compute reverse map for named struct Mordent
	// insertion point per field

	// Compute reverse map for named struct Multiple_rest
	// insertion point per field

	// Compute reverse map for named struct Name_display
	// insertion point per field
	stage.Name_display_Display_text_reverseMap = make(map[*Formatted_text]*Name_display)
	for name_display := range stage.Name_displays {
		_ = name_display
		for _, _formatted_text := range name_display.Display_text {
			stage.Name_display_Display_text_reverseMap[_formatted_text] = name_display
		}
	}
	stage.Name_display_Accidental_text_reverseMap = make(map[*Accidental_text]*Name_display)
	for name_display := range stage.Name_displays {
		_ = name_display
		for _, _accidental_text := range name_display.Accidental_text {
			stage.Name_display_Accidental_text_reverseMap[_accidental_text] = name_display
		}
	}

	// Compute reverse map for named struct Non_arpeggiate
	// insertion point per field

	// Compute reverse map for named struct Notations
	// insertion point per field
	stage.Notations_Tied_reverseMap = make(map[*Tied]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _tied := range notations.Tied {
			stage.Notations_Tied_reverseMap[_tied] = notations
		}
	}
	stage.Notations_Slur_reverseMap = make(map[*Slur]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _slur := range notations.Slur {
			stage.Notations_Slur_reverseMap[_slur] = notations
		}
	}
	stage.Notations_Tuplet_reverseMap = make(map[*Tuplet]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _tuplet := range notations.Tuplet {
			stage.Notations_Tuplet_reverseMap[_tuplet] = notations
		}
	}
	stage.Notations_Glissando_reverseMap = make(map[*Glissando]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _glissando := range notations.Glissando {
			stage.Notations_Glissando_reverseMap[_glissando] = notations
		}
	}
	stage.Notations_Slide_reverseMap = make(map[*Slide]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _slide := range notations.Slide {
			stage.Notations_Slide_reverseMap[_slide] = notations
		}
	}
	stage.Notations_Ornaments_reverseMap = make(map[*Ornaments]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _ornaments := range notations.Ornaments {
			stage.Notations_Ornaments_reverseMap[_ornaments] = notations
		}
	}
	stage.Notations_Technical_reverseMap = make(map[*Technical]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _technical := range notations.Technical {
			stage.Notations_Technical_reverseMap[_technical] = notations
		}
	}
	stage.Notations_Articulations_reverseMap = make(map[*Articulations]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _articulations := range notations.Articulations {
			stage.Notations_Articulations_reverseMap[_articulations] = notations
		}
	}
	stage.Notations_Dynamics_reverseMap = make(map[*Dynamics]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _dynamics := range notations.Dynamics {
			stage.Notations_Dynamics_reverseMap[_dynamics] = notations
		}
	}
	stage.Notations_Fermata_reverseMap = make(map[*Fermata]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _fermata := range notations.Fermata {
			stage.Notations_Fermata_reverseMap[_fermata] = notations
		}
	}
	stage.Notations_Arpeggiate_reverseMap = make(map[*Arpeggiate]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _arpeggiate := range notations.Arpeggiate {
			stage.Notations_Arpeggiate_reverseMap[_arpeggiate] = notations
		}
	}
	stage.Notations_Non_arpeggiate_reverseMap = make(map[*Non_arpeggiate]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _non_arpeggiate := range notations.Non_arpeggiate {
			stage.Notations_Non_arpeggiate_reverseMap[_non_arpeggiate] = notations
		}
	}
	stage.Notations_Accidental_mark_reverseMap = make(map[*Accidental_mark]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _accidental_mark := range notations.Accidental_mark {
			stage.Notations_Accidental_mark_reverseMap[_accidental_mark] = notations
		}
	}
	stage.Notations_Other_notation_reverseMap = make(map[*Other_notation]*Notations)
	for notations := range stage.Notationss {
		_ = notations
		for _, _other_notation := range notations.Other_notation {
			stage.Notations_Other_notation_reverseMap[_other_notation] = notations
		}
	}

	// Compute reverse map for named struct Note
	// insertion point per field
	stage.Note_Instrument_reverseMap = make(map[*Instrument]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _instrument := range note.Instrument {
			stage.Note_Instrument_reverseMap[_instrument] = note
		}
	}
	stage.Note_Dot_reverseMap = make(map[*Empty_placement]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _empty_placement := range note.Dot {
			stage.Note_Dot_reverseMap[_empty_placement] = note
		}
	}
	stage.Note_Notations_reverseMap = make(map[*Notations]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _notations := range note.Notations {
			stage.Note_Notations_reverseMap[_notations] = note
		}
	}
	stage.Note_Lyric_reverseMap = make(map[*Lyric]*Note)
	for note := range stage.Notes {
		_ = note
		for _, _lyric := range note.Lyric {
			stage.Note_Lyric_reverseMap[_lyric] = note
		}
	}

	// Compute reverse map for named struct Note_size
	// insertion point per field

	// Compute reverse map for named struct Note_type
	// insertion point per field

	// Compute reverse map for named struct Notehead
	// insertion point per field

	// Compute reverse map for named struct Notehead_text
	// insertion point per field
	stage.Notehead_text_Display_text_reverseMap = make(map[*Formatted_text]*Notehead_text)
	for notehead_text := range stage.Notehead_texts {
		_ = notehead_text
		for _, _formatted_text := range notehead_text.Display_text {
			stage.Notehead_text_Display_text_reverseMap[_formatted_text] = notehead_text
		}
	}
	stage.Notehead_text_Accidental_text_reverseMap = make(map[*Accidental_text]*Notehead_text)
	for notehead_text := range stage.Notehead_texts {
		_ = notehead_text
		for _, _accidental_text := range notehead_text.Accidental_text {
			stage.Notehead_text_Accidental_text_reverseMap[_accidental_text] = notehead_text
		}
	}

	// Compute reverse map for named struct Numeral
	// insertion point per field

	// Compute reverse map for named struct Numeral_key
	// insertion point per field

	// Compute reverse map for named struct Numeral_root
	// insertion point per field

	// Compute reverse map for named struct Octave_shift
	// insertion point per field

	// Compute reverse map for named struct Offset
	// insertion point per field

	// Compute reverse map for named struct Opus
	// insertion point per field

	// Compute reverse map for named struct Ornaments
	// insertion point per field
	stage.Ornaments_Trill_mark_reverseMap = make(map[*Empty_trill_sound]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _empty_trill_sound := range ornaments.Trill_mark {
			stage.Ornaments_Trill_mark_reverseMap[_empty_trill_sound] = ornaments
		}
	}
	stage.Ornaments_Turn_reverseMap = make(map[*Horizontal_turn]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _horizontal_turn := range ornaments.Turn {
			stage.Ornaments_Turn_reverseMap[_horizontal_turn] = ornaments
		}
	}
	stage.Ornaments_Delayed_turn_reverseMap = make(map[*Horizontal_turn]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _horizontal_turn := range ornaments.Delayed_turn {
			stage.Ornaments_Delayed_turn_reverseMap[_horizontal_turn] = ornaments
		}
	}
	stage.Ornaments_Inverted_turn_reverseMap = make(map[*Horizontal_turn]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _horizontal_turn := range ornaments.Inverted_turn {
			stage.Ornaments_Inverted_turn_reverseMap[_horizontal_turn] = ornaments
		}
	}
	stage.Ornaments_Delayed_inverted_turn_reverseMap = make(map[*Horizontal_turn]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _horizontal_turn := range ornaments.Delayed_inverted_turn {
			stage.Ornaments_Delayed_inverted_turn_reverseMap[_horizontal_turn] = ornaments
		}
	}
	stage.Ornaments_Vertical_turn_reverseMap = make(map[*Empty_trill_sound]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _empty_trill_sound := range ornaments.Vertical_turn {
			stage.Ornaments_Vertical_turn_reverseMap[_empty_trill_sound] = ornaments
		}
	}
	stage.Ornaments_Inverted_vertical_turn_reverseMap = make(map[*Empty_trill_sound]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _empty_trill_sound := range ornaments.Inverted_vertical_turn {
			stage.Ornaments_Inverted_vertical_turn_reverseMap[_empty_trill_sound] = ornaments
		}
	}
	stage.Ornaments_Shake_reverseMap = make(map[*Empty_trill_sound]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _empty_trill_sound := range ornaments.Shake {
			stage.Ornaments_Shake_reverseMap[_empty_trill_sound] = ornaments
		}
	}
	stage.Ornaments_Wavy_line_reverseMap = make(map[*Wavy_line]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _wavy_line := range ornaments.Wavy_line {
			stage.Ornaments_Wavy_line_reverseMap[_wavy_line] = ornaments
		}
	}
	stage.Ornaments_Mordent_reverseMap = make(map[*Mordent]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _mordent := range ornaments.Mordent {
			stage.Ornaments_Mordent_reverseMap[_mordent] = ornaments
		}
	}
	stage.Ornaments_Inverted_mordent_reverseMap = make(map[*Mordent]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _mordent := range ornaments.Inverted_mordent {
			stage.Ornaments_Inverted_mordent_reverseMap[_mordent] = ornaments
		}
	}
	stage.Ornaments_Schleifer_reverseMap = make(map[*Empty_placement]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _empty_placement := range ornaments.Schleifer {
			stage.Ornaments_Schleifer_reverseMap[_empty_placement] = ornaments
		}
	}
	stage.Ornaments_Tremolo_reverseMap = make(map[*Tremolo]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _tremolo := range ornaments.Tremolo {
			stage.Ornaments_Tremolo_reverseMap[_tremolo] = ornaments
		}
	}
	stage.Ornaments_Haydn_reverseMap = make(map[*Empty_trill_sound]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _empty_trill_sound := range ornaments.Haydn {
			stage.Ornaments_Haydn_reverseMap[_empty_trill_sound] = ornaments
		}
	}
	stage.Ornaments_Other_ornament_reverseMap = make(map[*Other_placement_text]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _other_placement_text := range ornaments.Other_ornament {
			stage.Ornaments_Other_ornament_reverseMap[_other_placement_text] = ornaments
		}
	}
	stage.Ornaments_Accidental_mark_reverseMap = make(map[*Accidental_mark]*Ornaments)
	for ornaments := range stage.Ornamentss {
		_ = ornaments
		for _, _accidental_mark := range ornaments.Accidental_mark {
			stage.Ornaments_Accidental_mark_reverseMap[_accidental_mark] = ornaments
		}
	}

	// Compute reverse map for named struct Other_appearance
	// insertion point per field

	// Compute reverse map for named struct Other_direction
	// insertion point per field

	// Compute reverse map for named struct Other_listening
	// insertion point per field

	// Compute reverse map for named struct Other_notation
	// insertion point per field

	// Compute reverse map for named struct Other_placement_text
	// insertion point per field

	// Compute reverse map for named struct Other_play
	// insertion point per field

	// Compute reverse map for named struct Other_text
	// insertion point per field

	// Compute reverse map for named struct Page_layout
	// insertion point per field

	// Compute reverse map for named struct Page_margins
	// insertion point per field

	// Compute reverse map for named struct Part_clef
	// insertion point per field

	// Compute reverse map for named struct Part_group
	// insertion point per field

	// Compute reverse map for named struct Part_link
	// insertion point per field
	stage.Part_link_Instrument_link_reverseMap = make(map[*Instrument_link]*Part_link)
	for part_link := range stage.Part_links {
		_ = part_link
		for _, _instrument_link := range part_link.Instrument_link {
			stage.Part_link_Instrument_link_reverseMap[_instrument_link] = part_link
		}
	}

	// Compute reverse map for named struct Part_list
	// insertion point per field

	// Compute reverse map for named struct Part_name
	// insertion point per field

	// Compute reverse map for named struct Part_symbol
	// insertion point per field

	// Compute reverse map for named struct Part_transpose
	// insertion point per field

	// Compute reverse map for named struct Pedal
	// insertion point per field

	// Compute reverse map for named struct Pedal_tuning
	// insertion point per field

	// Compute reverse map for named struct Per_minute
	// insertion point per field

	// Compute reverse map for named struct Percussion
	// insertion point per field

	// Compute reverse map for named struct Pitch
	// insertion point per field

	// Compute reverse map for named struct Pitched
	// insertion point per field

	// Compute reverse map for named struct Placement_text
	// insertion point per field

	// Compute reverse map for named struct Play
	// insertion point per field
	stage.Play_Other_play_reverseMap = make(map[*Other_play]*Play)
	for play := range stage.Plays {
		_ = play
		for _, _other_play := range play.Other_play {
			stage.Play_Other_play_reverseMap[_other_play] = play
		}
	}

	// Compute reverse map for named struct Player
	// insertion point per field

	// Compute reverse map for named struct Principal_voice
	// insertion point per field

	// Compute reverse map for named struct Print
	// insertion point per field
	stage.Print_Staff_layout_reverseMap = make(map[*Staff_layout]*Print)
	for print := range stage.Prints {
		_ = print
		for _, _staff_layout := range print.Staff_layout {
			stage.Print_Staff_layout_reverseMap[_staff_layout] = print
		}
	}

	// Compute reverse map for named struct Release
	// insertion point per field

	// Compute reverse map for named struct Repeat
	// insertion point per field

	// Compute reverse map for named struct Rest
	// insertion point per field

	// Compute reverse map for named struct Root
	// insertion point per field

	// Compute reverse map for named struct Root_step
	// insertion point per field

	// Compute reverse map for named struct Scaling
	// insertion point per field

	// Compute reverse map for named struct Scordatura
	// insertion point per field
	stage.Scordatura_Accord_reverseMap = make(map[*Accord]*Scordatura)
	for scordatura := range stage.Scordaturas {
		_ = scordatura
		for _, _accord := range scordatura.Accord {
			stage.Scordatura_Accord_reverseMap[_accord] = scordatura
		}
	}

	// Compute reverse map for named struct Score_instrument
	// insertion point per field

	// Compute reverse map for named struct Score_part
	// insertion point per field
	stage.Score_part_Part_link_reverseMap = make(map[*Part_link]*Score_part)
	for score_part := range stage.Score_parts {
		_ = score_part
		for _, _part_link := range score_part.Part_link {
			stage.Score_part_Part_link_reverseMap[_part_link] = score_part
		}
	}
	stage.Score_part_Score_instrument_reverseMap = make(map[*Score_instrument]*Score_part)
	for score_part := range stage.Score_parts {
		_ = score_part
		for _, _score_instrument := range score_part.Score_instrument {
			stage.Score_part_Score_instrument_reverseMap[_score_instrument] = score_part
		}
	}
	stage.Score_part_Player_reverseMap = make(map[*Player]*Score_part)
	for score_part := range stage.Score_parts {
		_ = score_part
		for _, _player := range score_part.Player {
			stage.Score_part_Player_reverseMap[_player] = score_part
		}
	}
	stage.Score_part_Midi_device_reverseMap = make(map[*Midi_device]*Score_part)
	for score_part := range stage.Score_parts {
		_ = score_part
		for _, _midi_device := range score_part.Midi_device {
			stage.Score_part_Midi_device_reverseMap[_midi_device] = score_part
		}
	}
	stage.Score_part_Midi_instrument_reverseMap = make(map[*Midi_instrument]*Score_part)
	for score_part := range stage.Score_parts {
		_ = score_part
		for _, _midi_instrument := range score_part.Midi_instrument {
			stage.Score_part_Midi_instrument_reverseMap[_midi_instrument] = score_part
		}
	}

	// Compute reverse map for named struct Score_partwise
	// insertion point per field
	stage.Score_partwise_Credit_reverseMap = make(map[*Credit]*Score_partwise)
	for score_partwise := range stage.Score_partwises {
		_ = score_partwise
		for _, _credit := range score_partwise.Credit {
			stage.Score_partwise_Credit_reverseMap[_credit] = score_partwise
		}
	}
	stage.Score_partwise_Part_reverseMap = make(map[*A_part]*Score_partwise)
	for score_partwise := range stage.Score_partwises {
		_ = score_partwise
		for _, _a_part := range score_partwise.Part {
			stage.Score_partwise_Part_reverseMap[_a_part] = score_partwise
		}
	}

	// Compute reverse map for named struct Score_timewise
	// insertion point per field
	stage.Score_timewise_Credit_reverseMap = make(map[*Credit]*Score_timewise)
	for score_timewise := range stage.Score_timewises {
		_ = score_timewise
		for _, _credit := range score_timewise.Credit {
			stage.Score_timewise_Credit_reverseMap[_credit] = score_timewise
		}
	}
	stage.Score_timewise_Measure_reverseMap = make(map[*A_measure_1]*Score_timewise)
	for score_timewise := range stage.Score_timewises {
		_ = score_timewise
		for _, _a_measure_1 := range score_timewise.Measure {
			stage.Score_timewise_Measure_reverseMap[_a_measure_1] = score_timewise
		}
	}

	// Compute reverse map for named struct Segno
	// insertion point per field

	// Compute reverse map for named struct Slash
	// insertion point per field

	// Compute reverse map for named struct Slide
	// insertion point per field

	// Compute reverse map for named struct Slur
	// insertion point per field

	// Compute reverse map for named struct Sound
	// insertion point per field
	stage.Sound_Instrument_change_reverseMap = make(map[*Instrument_change]*Sound)
	for sound := range stage.Sounds {
		_ = sound
		for _, _instrument_change := range sound.Instrument_change {
			stage.Sound_Instrument_change_reverseMap[_instrument_change] = sound
		}
	}
	stage.Sound_Midi_device_reverseMap = make(map[*Midi_device]*Sound)
	for sound := range stage.Sounds {
		_ = sound
		for _, _midi_device := range sound.Midi_device {
			stage.Sound_Midi_device_reverseMap[_midi_device] = sound
		}
	}
	stage.Sound_Midi_instrument_reverseMap = make(map[*Midi_instrument]*Sound)
	for sound := range stage.Sounds {
		_ = sound
		for _, _midi_instrument := range sound.Midi_instrument {
			stage.Sound_Midi_instrument_reverseMap[_midi_instrument] = sound
		}
	}
	stage.Sound_Play_reverseMap = make(map[*Play]*Sound)
	for sound := range stage.Sounds {
		_ = sound
		for _, _play := range sound.Play {
			stage.Sound_Play_reverseMap[_play] = sound
		}
	}

	// Compute reverse map for named struct Staff_details
	// insertion point per field
	stage.Staff_details_Line_detail_reverseMap = make(map[*Line_detail]*Staff_details)
	for staff_details := range stage.Staff_detailss {
		_ = staff_details
		for _, _line_detail := range staff_details.Line_detail {
			stage.Staff_details_Line_detail_reverseMap[_line_detail] = staff_details
		}
	}
	stage.Staff_details_Staff_tuning_reverseMap = make(map[*Staff_tuning]*Staff_details)
	for staff_details := range stage.Staff_detailss {
		_ = staff_details
		for _, _staff_tuning := range staff_details.Staff_tuning {
			stage.Staff_details_Staff_tuning_reverseMap[_staff_tuning] = staff_details
		}
	}

	// Compute reverse map for named struct Staff_divide
	// insertion point per field

	// Compute reverse map for named struct Staff_layout
	// insertion point per field

	// Compute reverse map for named struct Staff_size
	// insertion point per field

	// Compute reverse map for named struct Staff_tuning
	// insertion point per field

	// Compute reverse map for named struct Stem
	// insertion point per field

	// Compute reverse map for named struct Stick
	// insertion point per field

	// Compute reverse map for named struct String_mute
	// insertion point per field

	// Compute reverse map for named struct String_type
	// insertion point per field

	// Compute reverse map for named struct Strong_accent
	// insertion point per field

	// Compute reverse map for named struct Style_text
	// insertion point per field

	// Compute reverse map for named struct Supports
	// insertion point per field

	// Compute reverse map for named struct Swing
	// insertion point per field

	// Compute reverse map for named struct Sync
	// insertion point per field

	// Compute reverse map for named struct System_dividers
	// insertion point per field

	// Compute reverse map for named struct System_layout
	// insertion point per field

	// Compute reverse map for named struct System_margins
	// insertion point per field

	// Compute reverse map for named struct Tap
	// insertion point per field

	// Compute reverse map for named struct Technical
	// insertion point per field
	stage.Technical_Up_bow_reverseMap = make(map[*Empty_placement]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _empty_placement := range technical.Up_bow {
			stage.Technical_Up_bow_reverseMap[_empty_placement] = technical
		}
	}
	stage.Technical_Down_bow_reverseMap = make(map[*Empty_placement]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _empty_placement := range technical.Down_bow {
			stage.Technical_Down_bow_reverseMap[_empty_placement] = technical
		}
	}
	stage.Technical_Harmonic_reverseMap = make(map[*Harmonic]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _harmonic := range technical.Harmonic {
			stage.Technical_Harmonic_reverseMap[_harmonic] = technical
		}
	}
	stage.Technical_Open_string_reverseMap = make(map[*Empty_placement]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _empty_placement := range technical.Open_string {
			stage.Technical_Open_string_reverseMap[_empty_placement] = technical
		}
	}
	stage.Technical_Thumb_position_reverseMap = make(map[*Empty_placement]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _empty_placement := range technical.Thumb_position {
			stage.Technical_Thumb_position_reverseMap[_empty_placement] = technical
		}
	}
	stage.Technical_Fingering_reverseMap = make(map[*Fingering]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _fingering := range technical.Fingering {
			stage.Technical_Fingering_reverseMap[_fingering] = technical
		}
	}
	stage.Technical_Pluck_reverseMap = make(map[*Placement_text]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _placement_text := range technical.Pluck {
			stage.Technical_Pluck_reverseMap[_placement_text] = technical
		}
	}
	stage.Technical_Double_tongue_reverseMap = make(map[*Empty_placement]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _empty_placement := range technical.Double_tongue {
			stage.Technical_Double_tongue_reverseMap[_empty_placement] = technical
		}
	}
	stage.Technical_Triple_tongue_reverseMap = make(map[*Empty_placement]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _empty_placement := range technical.Triple_tongue {
			stage.Technical_Triple_tongue_reverseMap[_empty_placement] = technical
		}
	}
	stage.Technical_Stopped_reverseMap = make(map[*Empty_placement_smufl]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _empty_placement_smufl := range technical.Stopped {
			stage.Technical_Stopped_reverseMap[_empty_placement_smufl] = technical
		}
	}
	stage.Technical_Snap_pizzicato_reverseMap = make(map[*Empty_placement]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _empty_placement := range technical.Snap_pizzicato {
			stage.Technical_Snap_pizzicato_reverseMap[_empty_placement] = technical
		}
	}
	stage.Technical_Fret_reverseMap = make(map[*Fret]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _fret := range technical.Fret {
			stage.Technical_Fret_reverseMap[_fret] = technical
		}
	}
	stage.Technical_String_reverseMap = make(map[*String_type]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _string_type := range technical.String {
			stage.Technical_String_reverseMap[_string_type] = technical
		}
	}
	stage.Technical_Hammer_on_reverseMap = make(map[*Hammer_on_pull_off]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _hammer_on_pull_off := range technical.Hammer_on {
			stage.Technical_Hammer_on_reverseMap[_hammer_on_pull_off] = technical
		}
	}
	stage.Technical_Pull_off_reverseMap = make(map[*Hammer_on_pull_off]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _hammer_on_pull_off := range technical.Pull_off {
			stage.Technical_Pull_off_reverseMap[_hammer_on_pull_off] = technical
		}
	}
	stage.Technical_Bend_reverseMap = make(map[*Bend]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _bend := range technical.Bend {
			stage.Technical_Bend_reverseMap[_bend] = technical
		}
	}
	stage.Technical_Tap_reverseMap = make(map[*Tap]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _tap := range technical.Tap {
			stage.Technical_Tap_reverseMap[_tap] = technical
		}
	}
	stage.Technical_Heel_reverseMap = make(map[*Heel_toe]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _heel_toe := range technical.Heel {
			stage.Technical_Heel_reverseMap[_heel_toe] = technical
		}
	}
	stage.Technical_Toe_reverseMap = make(map[*Heel_toe]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _heel_toe := range technical.Toe {
			stage.Technical_Toe_reverseMap[_heel_toe] = technical
		}
	}
	stage.Technical_Fingernails_reverseMap = make(map[*Empty_placement]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _empty_placement := range technical.Fingernails {
			stage.Technical_Fingernails_reverseMap[_empty_placement] = technical
		}
	}
	stage.Technical_Hole_reverseMap = make(map[*Hole]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _hole := range technical.Hole {
			stage.Technical_Hole_reverseMap[_hole] = technical
		}
	}
	stage.Technical_Arrow_reverseMap = make(map[*Arrow]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _arrow := range technical.Arrow {
			stage.Technical_Arrow_reverseMap[_arrow] = technical
		}
	}
	stage.Technical_Handbell_reverseMap = make(map[*Handbell]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _handbell := range technical.Handbell {
			stage.Technical_Handbell_reverseMap[_handbell] = technical
		}
	}
	stage.Technical_Brass_bend_reverseMap = make(map[*Empty_placement]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _empty_placement := range technical.Brass_bend {
			stage.Technical_Brass_bend_reverseMap[_empty_placement] = technical
		}
	}
	stage.Technical_Flip_reverseMap = make(map[*Empty_placement]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _empty_placement := range technical.Flip {
			stage.Technical_Flip_reverseMap[_empty_placement] = technical
		}
	}
	stage.Technical_Smear_reverseMap = make(map[*Empty_placement]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _empty_placement := range technical.Smear {
			stage.Technical_Smear_reverseMap[_empty_placement] = technical
		}
	}
	stage.Technical_Open_reverseMap = make(map[*Empty_placement_smufl]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _empty_placement_smufl := range technical.Open {
			stage.Technical_Open_reverseMap[_empty_placement_smufl] = technical
		}
	}
	stage.Technical_Half_muted_reverseMap = make(map[*Empty_placement_smufl]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _empty_placement_smufl := range technical.Half_muted {
			stage.Technical_Half_muted_reverseMap[_empty_placement_smufl] = technical
		}
	}
	stage.Technical_Harmon_mute_reverseMap = make(map[*Harmon_mute]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _harmon_mute := range technical.Harmon_mute {
			stage.Technical_Harmon_mute_reverseMap[_harmon_mute] = technical
		}
	}
	stage.Technical_Golpe_reverseMap = make(map[*Empty_placement]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _empty_placement := range technical.Golpe {
			stage.Technical_Golpe_reverseMap[_empty_placement] = technical
		}
	}
	stage.Technical_Other_technical_reverseMap = make(map[*Other_placement_text]*Technical)
	for technical := range stage.Technicals {
		_ = technical
		for _, _other_placement_text := range technical.Other_technical {
			stage.Technical_Other_technical_reverseMap[_other_placement_text] = technical
		}
	}

	// Compute reverse map for named struct Text_element_data
	// insertion point per field

	// Compute reverse map for named struct Tie
	// insertion point per field

	// Compute reverse map for named struct Tied
	// insertion point per field

	// Compute reverse map for named struct Time
	// insertion point per field

	// Compute reverse map for named struct Time_modification
	// insertion point per field

	// Compute reverse map for named struct Timpani
	// insertion point per field

	// Compute reverse map for named struct Transpose
	// insertion point per field

	// Compute reverse map for named struct Tremolo
	// insertion point per field

	// Compute reverse map for named struct Tuplet
	// insertion point per field

	// Compute reverse map for named struct Tuplet_dot
	// insertion point per field

	// Compute reverse map for named struct Tuplet_number
	// insertion point per field

	// Compute reverse map for named struct Tuplet_portion
	// insertion point per field
	stage.Tuplet_portion_Tuplet_dot_reverseMap = make(map[*Tuplet_dot]*Tuplet_portion)
	for tuplet_portion := range stage.Tuplet_portions {
		_ = tuplet_portion
		for _, _tuplet_dot := range tuplet_portion.Tuplet_dot {
			stage.Tuplet_portion_Tuplet_dot_reverseMap[_tuplet_dot] = tuplet_portion
		}
	}

	// Compute reverse map for named struct Tuplet_type
	// insertion point per field

	// Compute reverse map for named struct Typed_text
	// insertion point per field

	// Compute reverse map for named struct Unpitched
	// insertion point per field

	// Compute reverse map for named struct Virtual_instrument
	// insertion point per field

	// Compute reverse map for named struct Wait
	// insertion point per field

	// Compute reverse map for named struct Wavy_line
	// insertion point per field

	// Compute reverse map for named struct Wedge
	// insertion point per field

	// Compute reverse map for named struct Wood
	// insertion point per field

	// Compute reverse map for named struct Work
	// insertion point per field

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	for instance := range stage.A_directives {
		res = append(res, instance)
	}

	for instance := range stage.A_measures {
		res = append(res, instance)
	}

	for instance := range stage.A_measure_1s {
		res = append(res, instance)
	}

	for instance := range stage.A_parts {
		res = append(res, instance)
	}

	for instance := range stage.A_part_1s {
		res = append(res, instance)
	}

	for instance := range stage.Accidentals {
		res = append(res, instance)
	}

	for instance := range stage.Accidental_marks {
		res = append(res, instance)
	}

	for instance := range stage.Accidental_texts {
		res = append(res, instance)
	}

	for instance := range stage.Accords {
		res = append(res, instance)
	}

	for instance := range stage.Accordion_registrations {
		res = append(res, instance)
	}

	for instance := range stage.Appearances {
		res = append(res, instance)
	}

	for instance := range stage.Arpeggiates {
		res = append(res, instance)
	}

	for instance := range stage.Arrows {
		res = append(res, instance)
	}

	for instance := range stage.Articulationss {
		res = append(res, instance)
	}

	for instance := range stage.Assesss {
		res = append(res, instance)
	}

	for instance := range stage.Attributess {
		res = append(res, instance)
	}

	for instance := range stage.Backups {
		res = append(res, instance)
	}

	for instance := range stage.Bar_style_colors {
		res = append(res, instance)
	}

	for instance := range stage.Barlines {
		res = append(res, instance)
	}

	for instance := range stage.Barres {
		res = append(res, instance)
	}

	for instance := range stage.Basss {
		res = append(res, instance)
	}

	for instance := range stage.Bass_steps {
		res = append(res, instance)
	}

	for instance := range stage.Beams {
		res = append(res, instance)
	}

	for instance := range stage.Beat_repeats {
		res = append(res, instance)
	}

	for instance := range stage.Beat_unit_tieds {
		res = append(res, instance)
	}

	for instance := range stage.Beaters {
		res = append(res, instance)
	}

	for instance := range stage.Bends {
		res = append(res, instance)
	}

	for instance := range stage.Bookmarks {
		res = append(res, instance)
	}

	for instance := range stage.Brackets {
		res = append(res, instance)
	}

	for instance := range stage.Breath_marks {
		res = append(res, instance)
	}

	for instance := range stage.Caesuras {
		res = append(res, instance)
	}

	for instance := range stage.Cancels {
		res = append(res, instance)
	}

	for instance := range stage.Clefs {
		res = append(res, instance)
	}

	for instance := range stage.Codas {
		res = append(res, instance)
	}

	for instance := range stage.Credits {
		res = append(res, instance)
	}

	for instance := range stage.Dashess {
		res = append(res, instance)
	}

	for instance := range stage.Defaultss {
		res = append(res, instance)
	}

	for instance := range stage.Degrees {
		res = append(res, instance)
	}

	for instance := range stage.Degree_alters {
		res = append(res, instance)
	}

	for instance := range stage.Degree_types {
		res = append(res, instance)
	}

	for instance := range stage.Degree_values {
		res = append(res, instance)
	}

	for instance := range stage.Directions {
		res = append(res, instance)
	}

	for instance := range stage.Direction_types {
		res = append(res, instance)
	}

	for instance := range stage.Distances {
		res = append(res, instance)
	}

	for instance := range stage.Doubles {
		res = append(res, instance)
	}

	for instance := range stage.Dynamicss {
		res = append(res, instance)
	}

	for instance := range stage.Effects {
		res = append(res, instance)
	}

	for instance := range stage.Elisions {
		res = append(res, instance)
	}

	for instance := range stage.Emptys {
		res = append(res, instance)
	}

	for instance := range stage.Empty_fonts {
		res = append(res, instance)
	}

	for instance := range stage.Empty_lines {
		res = append(res, instance)
	}

	for instance := range stage.Empty_placements {
		res = append(res, instance)
	}

	for instance := range stage.Empty_placement_smufls {
		res = append(res, instance)
	}

	for instance := range stage.Empty_print_object_style_aligns {
		res = append(res, instance)
	}

	for instance := range stage.Empty_print_styles {
		res = append(res, instance)
	}

	for instance := range stage.Empty_print_style_aligns {
		res = append(res, instance)
	}

	for instance := range stage.Empty_print_style_align_ids {
		res = append(res, instance)
	}

	for instance := range stage.Empty_trill_sounds {
		res = append(res, instance)
	}

	for instance := range stage.Encodings {
		res = append(res, instance)
	}

	for instance := range stage.Endings {
		res = append(res, instance)
	}

	for instance := range stage.Extends {
		res = append(res, instance)
	}

	for instance := range stage.Features {
		res = append(res, instance)
	}

	for instance := range stage.Fermatas {
		res = append(res, instance)
	}

	for instance := range stage.Figures {
		res = append(res, instance)
	}

	for instance := range stage.Figured_basss {
		res = append(res, instance)
	}

	for instance := range stage.Fingerings {
		res = append(res, instance)
	}

	for instance := range stage.First_frets {
		res = append(res, instance)
	}

	for instance := range stage.For_parts {
		res = append(res, instance)
	}

	for instance := range stage.Formatted_symbols {
		res = append(res, instance)
	}

	for instance := range stage.Formatted_symbol_ids {
		res = append(res, instance)
	}

	for instance := range stage.Formatted_texts {
		res = append(res, instance)
	}

	for instance := range stage.Formatted_text_ids {
		res = append(res, instance)
	}

	for instance := range stage.Forwards {
		res = append(res, instance)
	}

	for instance := range stage.Frames {
		res = append(res, instance)
	}

	for instance := range stage.Frame_notes {
		res = append(res, instance)
	}

	for instance := range stage.Frets {
		res = append(res, instance)
	}

	for instance := range stage.Glasss {
		res = append(res, instance)
	}

	for instance := range stage.Glissandos {
		res = append(res, instance)
	}

	for instance := range stage.Glyphs {
		res = append(res, instance)
	}

	for instance := range stage.Graces {
		res = append(res, instance)
	}

	for instance := range stage.Group_barlines {
		res = append(res, instance)
	}

	for instance := range stage.Group_names {
		res = append(res, instance)
	}

	for instance := range stage.Group_symbols {
		res = append(res, instance)
	}

	for instance := range stage.Groupings {
		res = append(res, instance)
	}

	for instance := range stage.Hammer_on_pull_offs {
		res = append(res, instance)
	}

	for instance := range stage.Handbells {
		res = append(res, instance)
	}

	for instance := range stage.Harmon_closeds {
		res = append(res, instance)
	}

	for instance := range stage.Harmon_mutes {
		res = append(res, instance)
	}

	for instance := range stage.Harmonics {
		res = append(res, instance)
	}

	for instance := range stage.Harmonys {
		res = append(res, instance)
	}

	for instance := range stage.Harmony_alters {
		res = append(res, instance)
	}

	for instance := range stage.Harp_pedalss {
		res = append(res, instance)
	}

	for instance := range stage.Heel_toes {
		res = append(res, instance)
	}

	for instance := range stage.Holes {
		res = append(res, instance)
	}

	for instance := range stage.Hole_closeds {
		res = append(res, instance)
	}

	for instance := range stage.Horizontal_turns {
		res = append(res, instance)
	}

	for instance := range stage.Identifications {
		res = append(res, instance)
	}

	for instance := range stage.Images {
		res = append(res, instance)
	}

	for instance := range stage.Instruments {
		res = append(res, instance)
	}

	for instance := range stage.Instrument_changes {
		res = append(res, instance)
	}

	for instance := range stage.Instrument_links {
		res = append(res, instance)
	}

	for instance := range stage.Interchangeables {
		res = append(res, instance)
	}

	for instance := range stage.Inversions {
		res = append(res, instance)
	}

	for instance := range stage.Keys {
		res = append(res, instance)
	}

	for instance := range stage.Key_accidentals {
		res = append(res, instance)
	}

	for instance := range stage.Key_octaves {
		res = append(res, instance)
	}

	for instance := range stage.Kinds {
		res = append(res, instance)
	}

	for instance := range stage.Levels {
		res = append(res, instance)
	}

	for instance := range stage.Line_details {
		res = append(res, instance)
	}

	for instance := range stage.Line_widths {
		res = append(res, instance)
	}

	for instance := range stage.Links {
		res = append(res, instance)
	}

	for instance := range stage.Listens {
		res = append(res, instance)
	}

	for instance := range stage.Listenings {
		res = append(res, instance)
	}

	for instance := range stage.Lyrics {
		res = append(res, instance)
	}

	for instance := range stage.Lyric_fonts {
		res = append(res, instance)
	}

	for instance := range stage.Lyric_languages {
		res = append(res, instance)
	}

	for instance := range stage.Measure_layouts {
		res = append(res, instance)
	}

	for instance := range stage.Measure_numberings {
		res = append(res, instance)
	}

	for instance := range stage.Measure_repeats {
		res = append(res, instance)
	}

	for instance := range stage.Measure_styles {
		res = append(res, instance)
	}

	for instance := range stage.Membranes {
		res = append(res, instance)
	}

	for instance := range stage.Metals {
		res = append(res, instance)
	}

	for instance := range stage.Metronomes {
		res = append(res, instance)
	}

	for instance := range stage.Metronome_beams {
		res = append(res, instance)
	}

	for instance := range stage.Metronome_notes {
		res = append(res, instance)
	}

	for instance := range stage.Metronome_tieds {
		res = append(res, instance)
	}

	for instance := range stage.Metronome_tuplets {
		res = append(res, instance)
	}

	for instance := range stage.Midi_devices {
		res = append(res, instance)
	}

	for instance := range stage.Midi_instruments {
		res = append(res, instance)
	}

	for instance := range stage.Miscellaneouss {
		res = append(res, instance)
	}

	for instance := range stage.Miscellaneous_fields {
		res = append(res, instance)
	}

	for instance := range stage.Mordents {
		res = append(res, instance)
	}

	for instance := range stage.Multiple_rests {
		res = append(res, instance)
	}

	for instance := range stage.Name_displays {
		res = append(res, instance)
	}

	for instance := range stage.Non_arpeggiates {
		res = append(res, instance)
	}

	for instance := range stage.Notationss {
		res = append(res, instance)
	}

	for instance := range stage.Notes {
		res = append(res, instance)
	}

	for instance := range stage.Note_sizes {
		res = append(res, instance)
	}

	for instance := range stage.Note_types {
		res = append(res, instance)
	}

	for instance := range stage.Noteheads {
		res = append(res, instance)
	}

	for instance := range stage.Notehead_texts {
		res = append(res, instance)
	}

	for instance := range stage.Numerals {
		res = append(res, instance)
	}

	for instance := range stage.Numeral_keys {
		res = append(res, instance)
	}

	for instance := range stage.Numeral_roots {
		res = append(res, instance)
	}

	for instance := range stage.Octave_shifts {
		res = append(res, instance)
	}

	for instance := range stage.Offsets {
		res = append(res, instance)
	}

	for instance := range stage.Opuss {
		res = append(res, instance)
	}

	for instance := range stage.Ornamentss {
		res = append(res, instance)
	}

	for instance := range stage.Other_appearances {
		res = append(res, instance)
	}

	for instance := range stage.Other_directions {
		res = append(res, instance)
	}

	for instance := range stage.Other_listenings {
		res = append(res, instance)
	}

	for instance := range stage.Other_notations {
		res = append(res, instance)
	}

	for instance := range stage.Other_placement_texts {
		res = append(res, instance)
	}

	for instance := range stage.Other_plays {
		res = append(res, instance)
	}

	for instance := range stage.Other_texts {
		res = append(res, instance)
	}

	for instance := range stage.Page_layouts {
		res = append(res, instance)
	}

	for instance := range stage.Page_marginss {
		res = append(res, instance)
	}

	for instance := range stage.Part_clefs {
		res = append(res, instance)
	}

	for instance := range stage.Part_groups {
		res = append(res, instance)
	}

	for instance := range stage.Part_links {
		res = append(res, instance)
	}

	for instance := range stage.Part_lists {
		res = append(res, instance)
	}

	for instance := range stage.Part_names {
		res = append(res, instance)
	}

	for instance := range stage.Part_symbols {
		res = append(res, instance)
	}

	for instance := range stage.Part_transposes {
		res = append(res, instance)
	}

	for instance := range stage.Pedals {
		res = append(res, instance)
	}

	for instance := range stage.Pedal_tunings {
		res = append(res, instance)
	}

	for instance := range stage.Per_minutes {
		res = append(res, instance)
	}

	for instance := range stage.Percussions {
		res = append(res, instance)
	}

	for instance := range stage.Pitchs {
		res = append(res, instance)
	}

	for instance := range stage.Pitcheds {
		res = append(res, instance)
	}

	for instance := range stage.Placement_texts {
		res = append(res, instance)
	}

	for instance := range stage.Plays {
		res = append(res, instance)
	}

	for instance := range stage.Players {
		res = append(res, instance)
	}

	for instance := range stage.Principal_voices {
		res = append(res, instance)
	}

	for instance := range stage.Prints {
		res = append(res, instance)
	}

	for instance := range stage.Releases {
		res = append(res, instance)
	}

	for instance := range stage.Repeats {
		res = append(res, instance)
	}

	for instance := range stage.Rests {
		res = append(res, instance)
	}

	for instance := range stage.Roots {
		res = append(res, instance)
	}

	for instance := range stage.Root_steps {
		res = append(res, instance)
	}

	for instance := range stage.Scalings {
		res = append(res, instance)
	}

	for instance := range stage.Scordaturas {
		res = append(res, instance)
	}

	for instance := range stage.Score_instruments {
		res = append(res, instance)
	}

	for instance := range stage.Score_parts {
		res = append(res, instance)
	}

	for instance := range stage.Score_partwises {
		res = append(res, instance)
	}

	for instance := range stage.Score_timewises {
		res = append(res, instance)
	}

	for instance := range stage.Segnos {
		res = append(res, instance)
	}

	for instance := range stage.Slashs {
		res = append(res, instance)
	}

	for instance := range stage.Slides {
		res = append(res, instance)
	}

	for instance := range stage.Slurs {
		res = append(res, instance)
	}

	for instance := range stage.Sounds {
		res = append(res, instance)
	}

	for instance := range stage.Staff_detailss {
		res = append(res, instance)
	}

	for instance := range stage.Staff_divides {
		res = append(res, instance)
	}

	for instance := range stage.Staff_layouts {
		res = append(res, instance)
	}

	for instance := range stage.Staff_sizes {
		res = append(res, instance)
	}

	for instance := range stage.Staff_tunings {
		res = append(res, instance)
	}

	for instance := range stage.Stems {
		res = append(res, instance)
	}

	for instance := range stage.Sticks {
		res = append(res, instance)
	}

	for instance := range stage.String_mutes {
		res = append(res, instance)
	}

	for instance := range stage.String_types {
		res = append(res, instance)
	}

	for instance := range stage.Strong_accents {
		res = append(res, instance)
	}

	for instance := range stage.Style_texts {
		res = append(res, instance)
	}

	for instance := range stage.Supportss {
		res = append(res, instance)
	}

	for instance := range stage.Swings {
		res = append(res, instance)
	}

	for instance := range stage.Syncs {
		res = append(res, instance)
	}

	for instance := range stage.System_dividerss {
		res = append(res, instance)
	}

	for instance := range stage.System_layouts {
		res = append(res, instance)
	}

	for instance := range stage.System_marginss {
		res = append(res, instance)
	}

	for instance := range stage.Taps {
		res = append(res, instance)
	}

	for instance := range stage.Technicals {
		res = append(res, instance)
	}

	for instance := range stage.Text_element_datas {
		res = append(res, instance)
	}

	for instance := range stage.Ties {
		res = append(res, instance)
	}

	for instance := range stage.Tieds {
		res = append(res, instance)
	}

	for instance := range stage.Times {
		res = append(res, instance)
	}

	for instance := range stage.Time_modifications {
		res = append(res, instance)
	}

	for instance := range stage.Timpanis {
		res = append(res, instance)
	}

	for instance := range stage.Transposes {
		res = append(res, instance)
	}

	for instance := range stage.Tremolos {
		res = append(res, instance)
	}

	for instance := range stage.Tuplets {
		res = append(res, instance)
	}

	for instance := range stage.Tuplet_dots {
		res = append(res, instance)
	}

	for instance := range stage.Tuplet_numbers {
		res = append(res, instance)
	}

	for instance := range stage.Tuplet_portions {
		res = append(res, instance)
	}

	for instance := range stage.Tuplet_types {
		res = append(res, instance)
	}

	for instance := range stage.Typed_texts {
		res = append(res, instance)
	}

	for instance := range stage.Unpitcheds {
		res = append(res, instance)
	}

	for instance := range stage.Virtual_instruments {
		res = append(res, instance)
	}

	for instance := range stage.Waits {
		res = append(res, instance)
	}

	for instance := range stage.Wavy_lines {
		res = append(res, instance)
	}

	for instance := range stage.Wedges {
		res = append(res, instance)
	}

	for instance := range stage.Woods {
		res = append(res, instance)
	}

	for instance := range stage.Works {
		res = append(res, instance)
	}

	return
}

// insertion point per named struct
func (a_directive *A_directive) GongCopy() GongstructIF {
	newInstance := new(A_directive)
	a_directive.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_measure *A_measure) GongCopy() GongstructIF {
	newInstance := new(A_measure)
	a_measure.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_measure_1 *A_measure_1) GongCopy() GongstructIF {
	newInstance := new(A_measure_1)
	a_measure_1.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_part *A_part) GongCopy() GongstructIF {
	newInstance := new(A_part)
	a_part.GongCopyBasicFields(newInstance)
	return newInstance
}

func (a_part_1 *A_part_1) GongCopy() GongstructIF {
	newInstance := new(A_part_1)
	a_part_1.GongCopyBasicFields(newInstance)
	return newInstance
}

func (accidental *Accidental) GongCopy() GongstructIF {
	newInstance := new(Accidental)
	accidental.GongCopyBasicFields(newInstance)
	return newInstance
}

func (accidental_mark *Accidental_mark) GongCopy() GongstructIF {
	newInstance := new(Accidental_mark)
	accidental_mark.GongCopyBasicFields(newInstance)
	return newInstance
}

func (accidental_text *Accidental_text) GongCopy() GongstructIF {
	newInstance := new(Accidental_text)
	accidental_text.GongCopyBasicFields(newInstance)
	return newInstance
}

func (accord *Accord) GongCopy() GongstructIF {
	newInstance := new(Accord)
	accord.GongCopyBasicFields(newInstance)
	return newInstance
}

func (accordion_registration *Accordion_registration) GongCopy() GongstructIF {
	newInstance := new(Accordion_registration)
	accordion_registration.GongCopyBasicFields(newInstance)
	return newInstance
}

func (appearance *Appearance) GongCopy() GongstructIF {
	newInstance := new(Appearance)
	appearance.GongCopyBasicFields(newInstance)
	return newInstance
}

func (arpeggiate *Arpeggiate) GongCopy() GongstructIF {
	newInstance := new(Arpeggiate)
	arpeggiate.GongCopyBasicFields(newInstance)
	return newInstance
}

func (arrow *Arrow) GongCopy() GongstructIF {
	newInstance := new(Arrow)
	arrow.GongCopyBasicFields(newInstance)
	return newInstance
}

func (articulations *Articulations) GongCopy() GongstructIF {
	newInstance := new(Articulations)
	articulations.GongCopyBasicFields(newInstance)
	return newInstance
}

func (assess *Assess) GongCopy() GongstructIF {
	newInstance := new(Assess)
	assess.GongCopyBasicFields(newInstance)
	return newInstance
}

func (attributes *Attributes) GongCopy() GongstructIF {
	newInstance := new(Attributes)
	attributes.GongCopyBasicFields(newInstance)
	return newInstance
}

func (backup *Backup) GongCopy() GongstructIF {
	newInstance := new(Backup)
	backup.GongCopyBasicFields(newInstance)
	return newInstance
}

func (bar_style_color *Bar_style_color) GongCopy() GongstructIF {
	newInstance := new(Bar_style_color)
	bar_style_color.GongCopyBasicFields(newInstance)
	return newInstance
}

func (barline *Barline) GongCopy() GongstructIF {
	newInstance := new(Barline)
	barline.GongCopyBasicFields(newInstance)
	return newInstance
}

func (barre *Barre) GongCopy() GongstructIF {
	newInstance := new(Barre)
	barre.GongCopyBasicFields(newInstance)
	return newInstance
}

func (bass *Bass) GongCopy() GongstructIF {
	newInstance := new(Bass)
	bass.GongCopyBasicFields(newInstance)
	return newInstance
}

func (bass_step *Bass_step) GongCopy() GongstructIF {
	newInstance := new(Bass_step)
	bass_step.GongCopyBasicFields(newInstance)
	return newInstance
}

func (beam *Beam) GongCopy() GongstructIF {
	newInstance := new(Beam)
	beam.GongCopyBasicFields(newInstance)
	return newInstance
}

func (beat_repeat *Beat_repeat) GongCopy() GongstructIF {
	newInstance := new(Beat_repeat)
	beat_repeat.GongCopyBasicFields(newInstance)
	return newInstance
}

func (beat_unit_tied *Beat_unit_tied) GongCopy() GongstructIF {
	newInstance := new(Beat_unit_tied)
	beat_unit_tied.GongCopyBasicFields(newInstance)
	return newInstance
}

func (beater *Beater) GongCopy() GongstructIF {
	newInstance := new(Beater)
	beater.GongCopyBasicFields(newInstance)
	return newInstance
}

func (bend *Bend) GongCopy() GongstructIF {
	newInstance := new(Bend)
	bend.GongCopyBasicFields(newInstance)
	return newInstance
}

func (bookmark *Bookmark) GongCopy() GongstructIF {
	newInstance := new(Bookmark)
	bookmark.GongCopyBasicFields(newInstance)
	return newInstance
}

func (bracket *Bracket) GongCopy() GongstructIF {
	newInstance := new(Bracket)
	bracket.GongCopyBasicFields(newInstance)
	return newInstance
}

func (breath_mark *Breath_mark) GongCopy() GongstructIF {
	newInstance := new(Breath_mark)
	breath_mark.GongCopyBasicFields(newInstance)
	return newInstance
}

func (caesura *Caesura) GongCopy() GongstructIF {
	newInstance := new(Caesura)
	caesura.GongCopyBasicFields(newInstance)
	return newInstance
}

func (cancel *Cancel) GongCopy() GongstructIF {
	newInstance := new(Cancel)
	cancel.GongCopyBasicFields(newInstance)
	return newInstance
}

func (clef *Clef) GongCopy() GongstructIF {
	newInstance := new(Clef)
	clef.GongCopyBasicFields(newInstance)
	return newInstance
}

func (coda *Coda) GongCopy() GongstructIF {
	newInstance := new(Coda)
	coda.GongCopyBasicFields(newInstance)
	return newInstance
}

func (credit *Credit) GongCopy() GongstructIF {
	newInstance := new(Credit)
	credit.GongCopyBasicFields(newInstance)
	return newInstance
}

func (dashes *Dashes) GongCopy() GongstructIF {
	newInstance := new(Dashes)
	dashes.GongCopyBasicFields(newInstance)
	return newInstance
}

func (defaults *Defaults) GongCopy() GongstructIF {
	newInstance := new(Defaults)
	defaults.GongCopyBasicFields(newInstance)
	return newInstance
}

func (degree *Degree) GongCopy() GongstructIF {
	newInstance := new(Degree)
	degree.GongCopyBasicFields(newInstance)
	return newInstance
}

func (degree_alter *Degree_alter) GongCopy() GongstructIF {
	newInstance := new(Degree_alter)
	degree_alter.GongCopyBasicFields(newInstance)
	return newInstance
}

func (degree_type *Degree_type) GongCopy() GongstructIF {
	newInstance := new(Degree_type)
	degree_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (degree_value *Degree_value) GongCopy() GongstructIF {
	newInstance := new(Degree_value)
	degree_value.GongCopyBasicFields(newInstance)
	return newInstance
}

func (direction *Direction) GongCopy() GongstructIF {
	newInstance := new(Direction)
	direction.GongCopyBasicFields(newInstance)
	return newInstance
}

func (direction_type *Direction_type) GongCopy() GongstructIF {
	newInstance := new(Direction_type)
	direction_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (distance *Distance) GongCopy() GongstructIF {
	newInstance := new(Distance)
	distance.GongCopyBasicFields(newInstance)
	return newInstance
}

func (double *Double) GongCopy() GongstructIF {
	newInstance := new(Double)
	double.GongCopyBasicFields(newInstance)
	return newInstance
}

func (dynamics *Dynamics) GongCopy() GongstructIF {
	newInstance := new(Dynamics)
	dynamics.GongCopyBasicFields(newInstance)
	return newInstance
}

func (effect *Effect) GongCopy() GongstructIF {
	newInstance := new(Effect)
	effect.GongCopyBasicFields(newInstance)
	return newInstance
}

func (elision *Elision) GongCopy() GongstructIF {
	newInstance := new(Elision)
	elision.GongCopyBasicFields(newInstance)
	return newInstance
}

func (empty *Empty) GongCopy() GongstructIF {
	newInstance := new(Empty)
	empty.GongCopyBasicFields(newInstance)
	return newInstance
}

func (empty_font *Empty_font) GongCopy() GongstructIF {
	newInstance := new(Empty_font)
	empty_font.GongCopyBasicFields(newInstance)
	return newInstance
}

func (empty_line *Empty_line) GongCopy() GongstructIF {
	newInstance := new(Empty_line)
	empty_line.GongCopyBasicFields(newInstance)
	return newInstance
}

func (empty_placement *Empty_placement) GongCopy() GongstructIF {
	newInstance := new(Empty_placement)
	empty_placement.GongCopyBasicFields(newInstance)
	return newInstance
}

func (empty_placement_smufl *Empty_placement_smufl) GongCopy() GongstructIF {
	newInstance := new(Empty_placement_smufl)
	empty_placement_smufl.GongCopyBasicFields(newInstance)
	return newInstance
}

func (empty_print_object_style_align *Empty_print_object_style_align) GongCopy() GongstructIF {
	newInstance := new(Empty_print_object_style_align)
	empty_print_object_style_align.GongCopyBasicFields(newInstance)
	return newInstance
}

func (empty_print_style *Empty_print_style) GongCopy() GongstructIF {
	newInstance := new(Empty_print_style)
	empty_print_style.GongCopyBasicFields(newInstance)
	return newInstance
}

func (empty_print_style_align *Empty_print_style_align) GongCopy() GongstructIF {
	newInstance := new(Empty_print_style_align)
	empty_print_style_align.GongCopyBasicFields(newInstance)
	return newInstance
}

func (empty_print_style_align_id *Empty_print_style_align_id) GongCopy() GongstructIF {
	newInstance := new(Empty_print_style_align_id)
	empty_print_style_align_id.GongCopyBasicFields(newInstance)
	return newInstance
}

func (empty_trill_sound *Empty_trill_sound) GongCopy() GongstructIF {
	newInstance := new(Empty_trill_sound)
	empty_trill_sound.GongCopyBasicFields(newInstance)
	return newInstance
}

func (encoding *Encoding) GongCopy() GongstructIF {
	newInstance := new(Encoding)
	encoding.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ending *Ending) GongCopy() GongstructIF {
	newInstance := new(Ending)
	ending.GongCopyBasicFields(newInstance)
	return newInstance
}

func (extend *Extend) GongCopy() GongstructIF {
	newInstance := new(Extend)
	extend.GongCopyBasicFields(newInstance)
	return newInstance
}

func (feature *Feature) GongCopy() GongstructIF {
	newInstance := new(Feature)
	feature.GongCopyBasicFields(newInstance)
	return newInstance
}

func (fermata *Fermata) GongCopy() GongstructIF {
	newInstance := new(Fermata)
	fermata.GongCopyBasicFields(newInstance)
	return newInstance
}

func (figure *Figure) GongCopy() GongstructIF {
	newInstance := new(Figure)
	figure.GongCopyBasicFields(newInstance)
	return newInstance
}

func (figured_bass *Figured_bass) GongCopy() GongstructIF {
	newInstance := new(Figured_bass)
	figured_bass.GongCopyBasicFields(newInstance)
	return newInstance
}

func (fingering *Fingering) GongCopy() GongstructIF {
	newInstance := new(Fingering)
	fingering.GongCopyBasicFields(newInstance)
	return newInstance
}

func (first_fret *First_fret) GongCopy() GongstructIF {
	newInstance := new(First_fret)
	first_fret.GongCopyBasicFields(newInstance)
	return newInstance
}

func (for_part *For_part) GongCopy() GongstructIF {
	newInstance := new(For_part)
	for_part.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formatted_symbol *Formatted_symbol) GongCopy() GongstructIF {
	newInstance := new(Formatted_symbol)
	formatted_symbol.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formatted_symbol_id *Formatted_symbol_id) GongCopy() GongstructIF {
	newInstance := new(Formatted_symbol_id)
	formatted_symbol_id.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formatted_text *Formatted_text) GongCopy() GongstructIF {
	newInstance := new(Formatted_text)
	formatted_text.GongCopyBasicFields(newInstance)
	return newInstance
}

func (formatted_text_id *Formatted_text_id) GongCopy() GongstructIF {
	newInstance := new(Formatted_text_id)
	formatted_text_id.GongCopyBasicFields(newInstance)
	return newInstance
}

func (forward *Forward) GongCopy() GongstructIF {
	newInstance := new(Forward)
	forward.GongCopyBasicFields(newInstance)
	return newInstance
}

func (frame *Frame) GongCopy() GongstructIF {
	newInstance := new(Frame)
	frame.GongCopyBasicFields(newInstance)
	return newInstance
}

func (frame_note *Frame_note) GongCopy() GongstructIF {
	newInstance := new(Frame_note)
	frame_note.GongCopyBasicFields(newInstance)
	return newInstance
}

func (fret *Fret) GongCopy() GongstructIF {
	newInstance := new(Fret)
	fret.GongCopyBasicFields(newInstance)
	return newInstance
}

func (glass *Glass) GongCopy() GongstructIF {
	newInstance := new(Glass)
	glass.GongCopyBasicFields(newInstance)
	return newInstance
}

func (glissando *Glissando) GongCopy() GongstructIF {
	newInstance := new(Glissando)
	glissando.GongCopyBasicFields(newInstance)
	return newInstance
}

func (glyph *Glyph) GongCopy() GongstructIF {
	newInstance := new(Glyph)
	glyph.GongCopyBasicFields(newInstance)
	return newInstance
}

func (grace *Grace) GongCopy() GongstructIF {
	newInstance := new(Grace)
	grace.GongCopyBasicFields(newInstance)
	return newInstance
}

func (group_barline *Group_barline) GongCopy() GongstructIF {
	newInstance := new(Group_barline)
	group_barline.GongCopyBasicFields(newInstance)
	return newInstance
}

func (group_name *Group_name) GongCopy() GongstructIF {
	newInstance := new(Group_name)
	group_name.GongCopyBasicFields(newInstance)
	return newInstance
}

func (group_symbol *Group_symbol) GongCopy() GongstructIF {
	newInstance := new(Group_symbol)
	group_symbol.GongCopyBasicFields(newInstance)
	return newInstance
}

func (grouping *Grouping) GongCopy() GongstructIF {
	newInstance := new(Grouping)
	grouping.GongCopyBasicFields(newInstance)
	return newInstance
}

func (hammer_on_pull_off *Hammer_on_pull_off) GongCopy() GongstructIF {
	newInstance := new(Hammer_on_pull_off)
	hammer_on_pull_off.GongCopyBasicFields(newInstance)
	return newInstance
}

func (handbell *Handbell) GongCopy() GongstructIF {
	newInstance := new(Handbell)
	handbell.GongCopyBasicFields(newInstance)
	return newInstance
}

func (harmon_closed *Harmon_closed) GongCopy() GongstructIF {
	newInstance := new(Harmon_closed)
	harmon_closed.GongCopyBasicFields(newInstance)
	return newInstance
}

func (harmon_mute *Harmon_mute) GongCopy() GongstructIF {
	newInstance := new(Harmon_mute)
	harmon_mute.GongCopyBasicFields(newInstance)
	return newInstance
}

func (harmonic *Harmonic) GongCopy() GongstructIF {
	newInstance := new(Harmonic)
	harmonic.GongCopyBasicFields(newInstance)
	return newInstance
}

func (harmony *Harmony) GongCopy() GongstructIF {
	newInstance := new(Harmony)
	harmony.GongCopyBasicFields(newInstance)
	return newInstance
}

func (harmony_alter *Harmony_alter) GongCopy() GongstructIF {
	newInstance := new(Harmony_alter)
	harmony_alter.GongCopyBasicFields(newInstance)
	return newInstance
}

func (harp_pedals *Harp_pedals) GongCopy() GongstructIF {
	newInstance := new(Harp_pedals)
	harp_pedals.GongCopyBasicFields(newInstance)
	return newInstance
}

func (heel_toe *Heel_toe) GongCopy() GongstructIF {
	newInstance := new(Heel_toe)
	heel_toe.GongCopyBasicFields(newInstance)
	return newInstance
}

func (hole *Hole) GongCopy() GongstructIF {
	newInstance := new(Hole)
	hole.GongCopyBasicFields(newInstance)
	return newInstance
}

func (hole_closed *Hole_closed) GongCopy() GongstructIF {
	newInstance := new(Hole_closed)
	hole_closed.GongCopyBasicFields(newInstance)
	return newInstance
}

func (horizontal_turn *Horizontal_turn) GongCopy() GongstructIF {
	newInstance := new(Horizontal_turn)
	horizontal_turn.GongCopyBasicFields(newInstance)
	return newInstance
}

func (identification *Identification) GongCopy() GongstructIF {
	newInstance := new(Identification)
	identification.GongCopyBasicFields(newInstance)
	return newInstance
}

func (image *Image) GongCopy() GongstructIF {
	newInstance := new(Image)
	image.GongCopyBasicFields(newInstance)
	return newInstance
}

func (instrument *Instrument) GongCopy() GongstructIF {
	newInstance := new(Instrument)
	instrument.GongCopyBasicFields(newInstance)
	return newInstance
}

func (instrument_change *Instrument_change) GongCopy() GongstructIF {
	newInstance := new(Instrument_change)
	instrument_change.GongCopyBasicFields(newInstance)
	return newInstance
}

func (instrument_link *Instrument_link) GongCopy() GongstructIF {
	newInstance := new(Instrument_link)
	instrument_link.GongCopyBasicFields(newInstance)
	return newInstance
}

func (interchangeable *Interchangeable) GongCopy() GongstructIF {
	newInstance := new(Interchangeable)
	interchangeable.GongCopyBasicFields(newInstance)
	return newInstance
}

func (inversion *Inversion) GongCopy() GongstructIF {
	newInstance := new(Inversion)
	inversion.GongCopyBasicFields(newInstance)
	return newInstance
}

func (key *Key) GongCopy() GongstructIF {
	newInstance := new(Key)
	key.GongCopyBasicFields(newInstance)
	return newInstance
}

func (key_accidental *Key_accidental) GongCopy() GongstructIF {
	newInstance := new(Key_accidental)
	key_accidental.GongCopyBasicFields(newInstance)
	return newInstance
}

func (key_octave *Key_octave) GongCopy() GongstructIF {
	newInstance := new(Key_octave)
	key_octave.GongCopyBasicFields(newInstance)
	return newInstance
}

func (kind *Kind) GongCopy() GongstructIF {
	newInstance := new(Kind)
	kind.GongCopyBasicFields(newInstance)
	return newInstance
}

func (level *Level) GongCopy() GongstructIF {
	newInstance := new(Level)
	level.GongCopyBasicFields(newInstance)
	return newInstance
}

func (line_detail *Line_detail) GongCopy() GongstructIF {
	newInstance := new(Line_detail)
	line_detail.GongCopyBasicFields(newInstance)
	return newInstance
}

func (line_width *Line_width) GongCopy() GongstructIF {
	newInstance := new(Line_width)
	line_width.GongCopyBasicFields(newInstance)
	return newInstance
}

func (link *Link) GongCopy() GongstructIF {
	newInstance := new(Link)
	link.GongCopyBasicFields(newInstance)
	return newInstance
}

func (listen *Listen) GongCopy() GongstructIF {
	newInstance := new(Listen)
	listen.GongCopyBasicFields(newInstance)
	return newInstance
}

func (listening *Listening) GongCopy() GongstructIF {
	newInstance := new(Listening)
	listening.GongCopyBasicFields(newInstance)
	return newInstance
}

func (lyric *Lyric) GongCopy() GongstructIF {
	newInstance := new(Lyric)
	lyric.GongCopyBasicFields(newInstance)
	return newInstance
}

func (lyric_font *Lyric_font) GongCopy() GongstructIF {
	newInstance := new(Lyric_font)
	lyric_font.GongCopyBasicFields(newInstance)
	return newInstance
}

func (lyric_language *Lyric_language) GongCopy() GongstructIF {
	newInstance := new(Lyric_language)
	lyric_language.GongCopyBasicFields(newInstance)
	return newInstance
}

func (measure_layout *Measure_layout) GongCopy() GongstructIF {
	newInstance := new(Measure_layout)
	measure_layout.GongCopyBasicFields(newInstance)
	return newInstance
}

func (measure_numbering *Measure_numbering) GongCopy() GongstructIF {
	newInstance := new(Measure_numbering)
	measure_numbering.GongCopyBasicFields(newInstance)
	return newInstance
}

func (measure_repeat *Measure_repeat) GongCopy() GongstructIF {
	newInstance := new(Measure_repeat)
	measure_repeat.GongCopyBasicFields(newInstance)
	return newInstance
}

func (measure_style *Measure_style) GongCopy() GongstructIF {
	newInstance := new(Measure_style)
	measure_style.GongCopyBasicFields(newInstance)
	return newInstance
}

func (membrane *Membrane) GongCopy() GongstructIF {
	newInstance := new(Membrane)
	membrane.GongCopyBasicFields(newInstance)
	return newInstance
}

func (metal *Metal) GongCopy() GongstructIF {
	newInstance := new(Metal)
	metal.GongCopyBasicFields(newInstance)
	return newInstance
}

func (metronome *Metronome) GongCopy() GongstructIF {
	newInstance := new(Metronome)
	metronome.GongCopyBasicFields(newInstance)
	return newInstance
}

func (metronome_beam *Metronome_beam) GongCopy() GongstructIF {
	newInstance := new(Metronome_beam)
	metronome_beam.GongCopyBasicFields(newInstance)
	return newInstance
}

func (metronome_note *Metronome_note) GongCopy() GongstructIF {
	newInstance := new(Metronome_note)
	metronome_note.GongCopyBasicFields(newInstance)
	return newInstance
}

func (metronome_tied *Metronome_tied) GongCopy() GongstructIF {
	newInstance := new(Metronome_tied)
	metronome_tied.GongCopyBasicFields(newInstance)
	return newInstance
}

func (metronome_tuplet *Metronome_tuplet) GongCopy() GongstructIF {
	newInstance := new(Metronome_tuplet)
	metronome_tuplet.GongCopyBasicFields(newInstance)
	return newInstance
}

func (midi_device *Midi_device) GongCopy() GongstructIF {
	newInstance := new(Midi_device)
	midi_device.GongCopyBasicFields(newInstance)
	return newInstance
}

func (midi_instrument *Midi_instrument) GongCopy() GongstructIF {
	newInstance := new(Midi_instrument)
	midi_instrument.GongCopyBasicFields(newInstance)
	return newInstance
}

func (miscellaneous *Miscellaneous) GongCopy() GongstructIF {
	newInstance := new(Miscellaneous)
	miscellaneous.GongCopyBasicFields(newInstance)
	return newInstance
}

func (miscellaneous_field *Miscellaneous_field) GongCopy() GongstructIF {
	newInstance := new(Miscellaneous_field)
	miscellaneous_field.GongCopyBasicFields(newInstance)
	return newInstance
}

func (mordent *Mordent) GongCopy() GongstructIF {
	newInstance := new(Mordent)
	mordent.GongCopyBasicFields(newInstance)
	return newInstance
}

func (multiple_rest *Multiple_rest) GongCopy() GongstructIF {
	newInstance := new(Multiple_rest)
	multiple_rest.GongCopyBasicFields(newInstance)
	return newInstance
}

func (name_display *Name_display) GongCopy() GongstructIF {
	newInstance := new(Name_display)
	name_display.GongCopyBasicFields(newInstance)
	return newInstance
}

func (non_arpeggiate *Non_arpeggiate) GongCopy() GongstructIF {
	newInstance := new(Non_arpeggiate)
	non_arpeggiate.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notations *Notations) GongCopy() GongstructIF {
	newInstance := new(Notations)
	notations.GongCopyBasicFields(newInstance)
	return newInstance
}

func (note *Note) GongCopy() GongstructIF {
	newInstance := new(Note)
	note.GongCopyBasicFields(newInstance)
	return newInstance
}

func (note_size *Note_size) GongCopy() GongstructIF {
	newInstance := new(Note_size)
	note_size.GongCopyBasicFields(newInstance)
	return newInstance
}

func (note_type *Note_type) GongCopy() GongstructIF {
	newInstance := new(Note_type)
	note_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notehead *Notehead) GongCopy() GongstructIF {
	newInstance := new(Notehead)
	notehead.GongCopyBasicFields(newInstance)
	return newInstance
}

func (notehead_text *Notehead_text) GongCopy() GongstructIF {
	newInstance := new(Notehead_text)
	notehead_text.GongCopyBasicFields(newInstance)
	return newInstance
}

func (numeral *Numeral) GongCopy() GongstructIF {
	newInstance := new(Numeral)
	numeral.GongCopyBasicFields(newInstance)
	return newInstance
}

func (numeral_key *Numeral_key) GongCopy() GongstructIF {
	newInstance := new(Numeral_key)
	numeral_key.GongCopyBasicFields(newInstance)
	return newInstance
}

func (numeral_root *Numeral_root) GongCopy() GongstructIF {
	newInstance := new(Numeral_root)
	numeral_root.GongCopyBasicFields(newInstance)
	return newInstance
}

func (octave_shift *Octave_shift) GongCopy() GongstructIF {
	newInstance := new(Octave_shift)
	octave_shift.GongCopyBasicFields(newInstance)
	return newInstance
}

func (offset *Offset) GongCopy() GongstructIF {
	newInstance := new(Offset)
	offset.GongCopyBasicFields(newInstance)
	return newInstance
}

func (opus *Opus) GongCopy() GongstructIF {
	newInstance := new(Opus)
	opus.GongCopyBasicFields(newInstance)
	return newInstance
}

func (ornaments *Ornaments) GongCopy() GongstructIF {
	newInstance := new(Ornaments)
	ornaments.GongCopyBasicFields(newInstance)
	return newInstance
}

func (other_appearance *Other_appearance) GongCopy() GongstructIF {
	newInstance := new(Other_appearance)
	other_appearance.GongCopyBasicFields(newInstance)
	return newInstance
}

func (other_direction *Other_direction) GongCopy() GongstructIF {
	newInstance := new(Other_direction)
	other_direction.GongCopyBasicFields(newInstance)
	return newInstance
}

func (other_listening *Other_listening) GongCopy() GongstructIF {
	newInstance := new(Other_listening)
	other_listening.GongCopyBasicFields(newInstance)
	return newInstance
}

func (other_notation *Other_notation) GongCopy() GongstructIF {
	newInstance := new(Other_notation)
	other_notation.GongCopyBasicFields(newInstance)
	return newInstance
}

func (other_placement_text *Other_placement_text) GongCopy() GongstructIF {
	newInstance := new(Other_placement_text)
	other_placement_text.GongCopyBasicFields(newInstance)
	return newInstance
}

func (other_play *Other_play) GongCopy() GongstructIF {
	newInstance := new(Other_play)
	other_play.GongCopyBasicFields(newInstance)
	return newInstance
}

func (other_text *Other_text) GongCopy() GongstructIF {
	newInstance := new(Other_text)
	other_text.GongCopyBasicFields(newInstance)
	return newInstance
}

func (page_layout *Page_layout) GongCopy() GongstructIF {
	newInstance := new(Page_layout)
	page_layout.GongCopyBasicFields(newInstance)
	return newInstance
}

func (page_margins *Page_margins) GongCopy() GongstructIF {
	newInstance := new(Page_margins)
	page_margins.GongCopyBasicFields(newInstance)
	return newInstance
}

func (part_clef *Part_clef) GongCopy() GongstructIF {
	newInstance := new(Part_clef)
	part_clef.GongCopyBasicFields(newInstance)
	return newInstance
}

func (part_group *Part_group) GongCopy() GongstructIF {
	newInstance := new(Part_group)
	part_group.GongCopyBasicFields(newInstance)
	return newInstance
}

func (part_link *Part_link) GongCopy() GongstructIF {
	newInstance := new(Part_link)
	part_link.GongCopyBasicFields(newInstance)
	return newInstance
}

func (part_list *Part_list) GongCopy() GongstructIF {
	newInstance := new(Part_list)
	part_list.GongCopyBasicFields(newInstance)
	return newInstance
}

func (part_name *Part_name) GongCopy() GongstructIF {
	newInstance := new(Part_name)
	part_name.GongCopyBasicFields(newInstance)
	return newInstance
}

func (part_symbol *Part_symbol) GongCopy() GongstructIF {
	newInstance := new(Part_symbol)
	part_symbol.GongCopyBasicFields(newInstance)
	return newInstance
}

func (part_transpose *Part_transpose) GongCopy() GongstructIF {
	newInstance := new(Part_transpose)
	part_transpose.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pedal *Pedal) GongCopy() GongstructIF {
	newInstance := new(Pedal)
	pedal.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pedal_tuning *Pedal_tuning) GongCopy() GongstructIF {
	newInstance := new(Pedal_tuning)
	pedal_tuning.GongCopyBasicFields(newInstance)
	return newInstance
}

func (per_minute *Per_minute) GongCopy() GongstructIF {
	newInstance := new(Per_minute)
	per_minute.GongCopyBasicFields(newInstance)
	return newInstance
}

func (percussion *Percussion) GongCopy() GongstructIF {
	newInstance := new(Percussion)
	percussion.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pitch *Pitch) GongCopy() GongstructIF {
	newInstance := new(Pitch)
	pitch.GongCopyBasicFields(newInstance)
	return newInstance
}

func (pitched *Pitched) GongCopy() GongstructIF {
	newInstance := new(Pitched)
	pitched.GongCopyBasicFields(newInstance)
	return newInstance
}

func (placement_text *Placement_text) GongCopy() GongstructIF {
	newInstance := new(Placement_text)
	placement_text.GongCopyBasicFields(newInstance)
	return newInstance
}

func (play *Play) GongCopy() GongstructIF {
	newInstance := new(Play)
	play.GongCopyBasicFields(newInstance)
	return newInstance
}

func (player *Player) GongCopy() GongstructIF {
	newInstance := new(Player)
	player.GongCopyBasicFields(newInstance)
	return newInstance
}

func (principal_voice *Principal_voice) GongCopy() GongstructIF {
	newInstance := new(Principal_voice)
	principal_voice.GongCopyBasicFields(newInstance)
	return newInstance
}

func (print *Print) GongCopy() GongstructIF {
	newInstance := new(Print)
	print.GongCopyBasicFields(newInstance)
	return newInstance
}

func (release *Release) GongCopy() GongstructIF {
	newInstance := new(Release)
	release.GongCopyBasicFields(newInstance)
	return newInstance
}

func (repeat *Repeat) GongCopy() GongstructIF {
	newInstance := new(Repeat)
	repeat.GongCopyBasicFields(newInstance)
	return newInstance
}

func (rest *Rest) GongCopy() GongstructIF {
	newInstance := new(Rest)
	rest.GongCopyBasicFields(newInstance)
	return newInstance
}

func (root *Root) GongCopy() GongstructIF {
	newInstance := new(Root)
	root.GongCopyBasicFields(newInstance)
	return newInstance
}

func (root_step *Root_step) GongCopy() GongstructIF {
	newInstance := new(Root_step)
	root_step.GongCopyBasicFields(newInstance)
	return newInstance
}

func (scaling *Scaling) GongCopy() GongstructIF {
	newInstance := new(Scaling)
	scaling.GongCopyBasicFields(newInstance)
	return newInstance
}

func (scordatura *Scordatura) GongCopy() GongstructIF {
	newInstance := new(Scordatura)
	scordatura.GongCopyBasicFields(newInstance)
	return newInstance
}

func (score_instrument *Score_instrument) GongCopy() GongstructIF {
	newInstance := new(Score_instrument)
	score_instrument.GongCopyBasicFields(newInstance)
	return newInstance
}

func (score_part *Score_part) GongCopy() GongstructIF {
	newInstance := new(Score_part)
	score_part.GongCopyBasicFields(newInstance)
	return newInstance
}

func (score_partwise *Score_partwise) GongCopy() GongstructIF {
	newInstance := new(Score_partwise)
	score_partwise.GongCopyBasicFields(newInstance)
	return newInstance
}

func (score_timewise *Score_timewise) GongCopy() GongstructIF {
	newInstance := new(Score_timewise)
	score_timewise.GongCopyBasicFields(newInstance)
	return newInstance
}

func (segno *Segno) GongCopy() GongstructIF {
	newInstance := new(Segno)
	segno.GongCopyBasicFields(newInstance)
	return newInstance
}

func (slash *Slash) GongCopy() GongstructIF {
	newInstance := new(Slash)
	slash.GongCopyBasicFields(newInstance)
	return newInstance
}

func (slide *Slide) GongCopy() GongstructIF {
	newInstance := new(Slide)
	slide.GongCopyBasicFields(newInstance)
	return newInstance
}

func (slur *Slur) GongCopy() GongstructIF {
	newInstance := new(Slur)
	slur.GongCopyBasicFields(newInstance)
	return newInstance
}

func (sound *Sound) GongCopy() GongstructIF {
	newInstance := new(Sound)
	sound.GongCopyBasicFields(newInstance)
	return newInstance
}

func (staff_details *Staff_details) GongCopy() GongstructIF {
	newInstance := new(Staff_details)
	staff_details.GongCopyBasicFields(newInstance)
	return newInstance
}

func (staff_divide *Staff_divide) GongCopy() GongstructIF {
	newInstance := new(Staff_divide)
	staff_divide.GongCopyBasicFields(newInstance)
	return newInstance
}

func (staff_layout *Staff_layout) GongCopy() GongstructIF {
	newInstance := new(Staff_layout)
	staff_layout.GongCopyBasicFields(newInstance)
	return newInstance
}

func (staff_size *Staff_size) GongCopy() GongstructIF {
	newInstance := new(Staff_size)
	staff_size.GongCopyBasicFields(newInstance)
	return newInstance
}

func (staff_tuning *Staff_tuning) GongCopy() GongstructIF {
	newInstance := new(Staff_tuning)
	staff_tuning.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stem *Stem) GongCopy() GongstructIF {
	newInstance := new(Stem)
	stem.GongCopyBasicFields(newInstance)
	return newInstance
}

func (stick *Stick) GongCopy() GongstructIF {
	newInstance := new(Stick)
	stick.GongCopyBasicFields(newInstance)
	return newInstance
}

func (string_mute *String_mute) GongCopy() GongstructIF {
	newInstance := new(String_mute)
	string_mute.GongCopyBasicFields(newInstance)
	return newInstance
}

func (string_type *String_type) GongCopy() GongstructIF {
	newInstance := new(String_type)
	string_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (strong_accent *Strong_accent) GongCopy() GongstructIF {
	newInstance := new(Strong_accent)
	strong_accent.GongCopyBasicFields(newInstance)
	return newInstance
}

func (style_text *Style_text) GongCopy() GongstructIF {
	newInstance := new(Style_text)
	style_text.GongCopyBasicFields(newInstance)
	return newInstance
}

func (supports *Supports) GongCopy() GongstructIF {
	newInstance := new(Supports)
	supports.GongCopyBasicFields(newInstance)
	return newInstance
}

func (swing *Swing) GongCopy() GongstructIF {
	newInstance := new(Swing)
	swing.GongCopyBasicFields(newInstance)
	return newInstance
}

func (sync *Sync) GongCopy() GongstructIF {
	newInstance := new(Sync)
	sync.GongCopyBasicFields(newInstance)
	return newInstance
}

func (system_dividers *System_dividers) GongCopy() GongstructIF {
	newInstance := new(System_dividers)
	system_dividers.GongCopyBasicFields(newInstance)
	return newInstance
}

func (system_layout *System_layout) GongCopy() GongstructIF {
	newInstance := new(System_layout)
	system_layout.GongCopyBasicFields(newInstance)
	return newInstance
}

func (system_margins *System_margins) GongCopy() GongstructIF {
	newInstance := new(System_margins)
	system_margins.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tap *Tap) GongCopy() GongstructIF {
	newInstance := new(Tap)
	tap.GongCopyBasicFields(newInstance)
	return newInstance
}

func (technical *Technical) GongCopy() GongstructIF {
	newInstance := new(Technical)
	technical.GongCopyBasicFields(newInstance)
	return newInstance
}

func (text_element_data *Text_element_data) GongCopy() GongstructIF {
	newInstance := new(Text_element_data)
	text_element_data.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tie *Tie) GongCopy() GongstructIF {
	newInstance := new(Tie)
	tie.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tied *Tied) GongCopy() GongstructIF {
	newInstance := new(Tied)
	tied.GongCopyBasicFields(newInstance)
	return newInstance
}

func (time *Time) GongCopy() GongstructIF {
	newInstance := new(Time)
	time.GongCopyBasicFields(newInstance)
	return newInstance
}

func (time_modification *Time_modification) GongCopy() GongstructIF {
	newInstance := new(Time_modification)
	time_modification.GongCopyBasicFields(newInstance)
	return newInstance
}

func (timpani *Timpani) GongCopy() GongstructIF {
	newInstance := new(Timpani)
	timpani.GongCopyBasicFields(newInstance)
	return newInstance
}

func (transpose *Transpose) GongCopy() GongstructIF {
	newInstance := new(Transpose)
	transpose.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tremolo *Tremolo) GongCopy() GongstructIF {
	newInstance := new(Tremolo)
	tremolo.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tuplet *Tuplet) GongCopy() GongstructIF {
	newInstance := new(Tuplet)
	tuplet.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tuplet_dot *Tuplet_dot) GongCopy() GongstructIF {
	newInstance := new(Tuplet_dot)
	tuplet_dot.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tuplet_number *Tuplet_number) GongCopy() GongstructIF {
	newInstance := new(Tuplet_number)
	tuplet_number.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tuplet_portion *Tuplet_portion) GongCopy() GongstructIF {
	newInstance := new(Tuplet_portion)
	tuplet_portion.GongCopyBasicFields(newInstance)
	return newInstance
}

func (tuplet_type *Tuplet_type) GongCopy() GongstructIF {
	newInstance := new(Tuplet_type)
	tuplet_type.GongCopyBasicFields(newInstance)
	return newInstance
}

func (typed_text *Typed_text) GongCopy() GongstructIF {
	newInstance := new(Typed_text)
	typed_text.GongCopyBasicFields(newInstance)
	return newInstance
}

func (unpitched *Unpitched) GongCopy() GongstructIF {
	newInstance := new(Unpitched)
	unpitched.GongCopyBasicFields(newInstance)
	return newInstance
}

func (virtual_instrument *Virtual_instrument) GongCopy() GongstructIF {
	newInstance := new(Virtual_instrument)
	virtual_instrument.GongCopyBasicFields(newInstance)
	return newInstance
}

func (wait *Wait) GongCopy() GongstructIF {
	newInstance := new(Wait)
	wait.GongCopyBasicFields(newInstance)
	return newInstance
}

func (wavy_line *Wavy_line) GongCopy() GongstructIF {
	newInstance := new(Wavy_line)
	wavy_line.GongCopyBasicFields(newInstance)
	return newInstance
}

func (wedge *Wedge) GongCopy() GongstructIF {
	newInstance := new(Wedge)
	wedge.GongCopyBasicFields(newInstance)
	return newInstance
}

func (wood *Wood) GongCopy() GongstructIF {
	newInstance := new(Wood)
	wood.GongCopyBasicFields(newInstance)
	return newInstance
}

func (work *Work) GongCopy() GongstructIF {
	newInstance := new(Work)
	work.GongCopyBasicFields(newInstance)
	return newInstance
}

// insertion point per named struct
func (a_directive *A_directive) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_directive).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_directive), uint64(stage.GetOrder(a_directive)))
	return
}

func (a_measure *A_measure) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_measure).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_measure), uint64(stage.GetOrder(a_measure)))
	return
}

func (a_measure_1 *A_measure_1) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_measure_1).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_measure_1), uint64(stage.GetOrder(a_measure_1)))
	return
}

func (a_part *A_part) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_part).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_part), uint64(stage.GetOrder(a_part)))
	return
}

func (a_part_1 *A_part_1) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(a_part_1).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(a_part_1), uint64(stage.GetOrder(a_part_1)))
	return
}

func (accidental *Accidental) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(accidental).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(accidental), uint64(stage.GetOrder(accidental)))
	return
}

func (accidental_mark *Accidental_mark) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(accidental_mark).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(accidental_mark), uint64(stage.GetOrder(accidental_mark)))
	return
}

func (accidental_text *Accidental_text) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(accidental_text).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(accidental_text), uint64(stage.GetOrder(accidental_text)))
	return
}

func (accord *Accord) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(accord).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(accord), uint64(stage.GetOrder(accord)))
	return
}

func (accordion_registration *Accordion_registration) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(accordion_registration).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(accordion_registration), uint64(stage.GetOrder(accordion_registration)))
	return
}

func (appearance *Appearance) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(appearance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(appearance), uint64(stage.GetOrder(appearance)))
	return
}

func (arpeggiate *Arpeggiate) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(arpeggiate).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(arpeggiate), uint64(stage.GetOrder(arpeggiate)))
	return
}

func (arrow *Arrow) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(arrow).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(arrow), uint64(stage.GetOrder(arrow)))
	return
}

func (articulations *Articulations) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(articulations).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(articulations), uint64(stage.GetOrder(articulations)))
	return
}

func (assess *Assess) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(assess).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(assess), uint64(stage.GetOrder(assess)))
	return
}

func (attributes *Attributes) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(attributes).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(attributes), uint64(stage.GetOrder(attributes)))
	return
}

func (backup *Backup) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(backup).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(backup), uint64(stage.GetOrder(backup)))
	return
}

func (bar_style_color *Bar_style_color) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bar_style_color).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(bar_style_color), uint64(stage.GetOrder(bar_style_color)))
	return
}

func (barline *Barline) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(barline).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(barline), uint64(stage.GetOrder(barline)))
	return
}

func (barre *Barre) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(barre).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(barre), uint64(stage.GetOrder(barre)))
	return
}

func (bass *Bass) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bass).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(bass), uint64(stage.GetOrder(bass)))
	return
}

func (bass_step *Bass_step) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bass_step).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(bass_step), uint64(stage.GetOrder(bass_step)))
	return
}

func (beam *Beam) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(beam).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(beam), uint64(stage.GetOrder(beam)))
	return
}

func (beat_repeat *Beat_repeat) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(beat_repeat).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(beat_repeat), uint64(stage.GetOrder(beat_repeat)))
	return
}

func (beat_unit_tied *Beat_unit_tied) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(beat_unit_tied).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(beat_unit_tied), uint64(stage.GetOrder(beat_unit_tied)))
	return
}

func (beater *Beater) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(beater).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(beater), uint64(stage.GetOrder(beater)))
	return
}

func (bend *Bend) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bend).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(bend), uint64(stage.GetOrder(bend)))
	return
}

func (bookmark *Bookmark) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bookmark).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(bookmark), uint64(stage.GetOrder(bookmark)))
	return
}

func (bracket *Bracket) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(bracket).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(bracket), uint64(stage.GetOrder(bracket)))
	return
}

func (breath_mark *Breath_mark) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(breath_mark).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(breath_mark), uint64(stage.GetOrder(breath_mark)))
	return
}

func (caesura *Caesura) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(caesura).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(caesura), uint64(stage.GetOrder(caesura)))
	return
}

func (cancel *Cancel) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(cancel).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(cancel), uint64(stage.GetOrder(cancel)))
	return
}

func (clef *Clef) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(clef).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(clef), uint64(stage.GetOrder(clef)))
	return
}

func (coda *Coda) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(coda).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(coda), uint64(stage.GetOrder(coda)))
	return
}

func (credit *Credit) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(credit).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(credit), uint64(stage.GetOrder(credit)))
	return
}

func (dashes *Dashes) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(dashes).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(dashes), uint64(stage.GetOrder(dashes)))
	return
}

func (defaults *Defaults) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(defaults).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(defaults), uint64(stage.GetOrder(defaults)))
	return
}

func (degree *Degree) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(degree).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(degree), uint64(stage.GetOrder(degree)))
	return
}

func (degree_alter *Degree_alter) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(degree_alter).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(degree_alter), uint64(stage.GetOrder(degree_alter)))
	return
}

func (degree_type *Degree_type) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(degree_type).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(degree_type), uint64(stage.GetOrder(degree_type)))
	return
}

func (degree_value *Degree_value) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(degree_value).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(degree_value), uint64(stage.GetOrder(degree_value)))
	return
}

func (direction *Direction) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(direction).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(direction), uint64(stage.GetOrder(direction)))
	return
}

func (direction_type *Direction_type) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(direction_type).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(direction_type), uint64(stage.GetOrder(direction_type)))
	return
}

func (distance *Distance) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(distance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(distance), uint64(stage.GetOrder(distance)))
	return
}

func (double *Double) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(double).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(double), uint64(stage.GetOrder(double)))
	return
}

func (dynamics *Dynamics) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(dynamics).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(dynamics), uint64(stage.GetOrder(dynamics)))
	return
}

func (effect *Effect) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(effect).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(effect), uint64(stage.GetOrder(effect)))
	return
}

func (elision *Elision) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(elision).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(elision), uint64(stage.GetOrder(elision)))
	return
}

func (empty *Empty) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(empty).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(empty), uint64(stage.GetOrder(empty)))
	return
}

func (empty_font *Empty_font) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(empty_font).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(empty_font), uint64(stage.GetOrder(empty_font)))
	return
}

func (empty_line *Empty_line) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(empty_line).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(empty_line), uint64(stage.GetOrder(empty_line)))
	return
}

func (empty_placement *Empty_placement) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(empty_placement).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(empty_placement), uint64(stage.GetOrder(empty_placement)))
	return
}

func (empty_placement_smufl *Empty_placement_smufl) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(empty_placement_smufl).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(empty_placement_smufl), uint64(stage.GetOrder(empty_placement_smufl)))
	return
}

func (empty_print_object_style_align *Empty_print_object_style_align) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(empty_print_object_style_align).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(empty_print_object_style_align), uint64(stage.GetOrder(empty_print_object_style_align)))
	return
}

func (empty_print_style *Empty_print_style) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(empty_print_style).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(empty_print_style), uint64(stage.GetOrder(empty_print_style)))
	return
}

func (empty_print_style_align *Empty_print_style_align) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(empty_print_style_align).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(empty_print_style_align), uint64(stage.GetOrder(empty_print_style_align)))
	return
}

func (empty_print_style_align_id *Empty_print_style_align_id) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(empty_print_style_align_id).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(empty_print_style_align_id), uint64(stage.GetOrder(empty_print_style_align_id)))
	return
}

func (empty_trill_sound *Empty_trill_sound) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(empty_trill_sound).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(empty_trill_sound), uint64(stage.GetOrder(empty_trill_sound)))
	return
}

func (encoding *Encoding) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(encoding).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(encoding), uint64(stage.GetOrder(encoding)))
	return
}

func (ending *Ending) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(ending).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(ending), uint64(stage.GetOrder(ending)))
	return
}

func (extend *Extend) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(extend).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(extend), uint64(stage.GetOrder(extend)))
	return
}

func (feature *Feature) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(feature).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(feature), uint64(stage.GetOrder(feature)))
	return
}

func (fermata *Fermata) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(fermata).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(fermata), uint64(stage.GetOrder(fermata)))
	return
}

func (figure *Figure) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(figure).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(figure), uint64(stage.GetOrder(figure)))
	return
}

func (figured_bass *Figured_bass) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(figured_bass).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(figured_bass), uint64(stage.GetOrder(figured_bass)))
	return
}

func (fingering *Fingering) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(fingering).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(fingering), uint64(stage.GetOrder(fingering)))
	return
}

func (first_fret *First_fret) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(first_fret).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(first_fret), uint64(stage.GetOrder(first_fret)))
	return
}

func (for_part *For_part) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(for_part).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(for_part), uint64(stage.GetOrder(for_part)))
	return
}

func (formatted_symbol *Formatted_symbol) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formatted_symbol).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(formatted_symbol), uint64(stage.GetOrder(formatted_symbol)))
	return
}

func (formatted_symbol_id *Formatted_symbol_id) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formatted_symbol_id).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(formatted_symbol_id), uint64(stage.GetOrder(formatted_symbol_id)))
	return
}

func (formatted_text *Formatted_text) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formatted_text).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(formatted_text), uint64(stage.GetOrder(formatted_text)))
	return
}

func (formatted_text_id *Formatted_text_id) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(formatted_text_id).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(formatted_text_id), uint64(stage.GetOrder(formatted_text_id)))
	return
}

func (forward *Forward) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(forward).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(forward), uint64(stage.GetOrder(forward)))
	return
}

func (frame *Frame) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(frame).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(frame), uint64(stage.GetOrder(frame)))
	return
}

func (frame_note *Frame_note) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(frame_note).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(frame_note), uint64(stage.GetOrder(frame_note)))
	return
}

func (fret *Fret) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(fret).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(fret), uint64(stage.GetOrder(fret)))
	return
}

func (glass *Glass) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(glass).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(glass), uint64(stage.GetOrder(glass)))
	return
}

func (glissando *Glissando) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(glissando).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(glissando), uint64(stage.GetOrder(glissando)))
	return
}

func (glyph *Glyph) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(glyph).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(glyph), uint64(stage.GetOrder(glyph)))
	return
}

func (grace *Grace) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(grace).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(grace), uint64(stage.GetOrder(grace)))
	return
}

func (group_barline *Group_barline) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(group_barline).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(group_barline), uint64(stage.GetOrder(group_barline)))
	return
}

func (group_name *Group_name) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(group_name).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(group_name), uint64(stage.GetOrder(group_name)))
	return
}

func (group_symbol *Group_symbol) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(group_symbol).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(group_symbol), uint64(stage.GetOrder(group_symbol)))
	return
}

func (grouping *Grouping) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(grouping).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(grouping), uint64(stage.GetOrder(grouping)))
	return
}

func (hammer_on_pull_off *Hammer_on_pull_off) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(hammer_on_pull_off).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(hammer_on_pull_off), uint64(stage.GetOrder(hammer_on_pull_off)))
	return
}

func (handbell *Handbell) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(handbell).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(handbell), uint64(stage.GetOrder(handbell)))
	return
}

func (harmon_closed *Harmon_closed) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(harmon_closed).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(harmon_closed), uint64(stage.GetOrder(harmon_closed)))
	return
}

func (harmon_mute *Harmon_mute) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(harmon_mute).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(harmon_mute), uint64(stage.GetOrder(harmon_mute)))
	return
}

func (harmonic *Harmonic) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(harmonic).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(harmonic), uint64(stage.GetOrder(harmonic)))
	return
}

func (harmony *Harmony) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(harmony).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(harmony), uint64(stage.GetOrder(harmony)))
	return
}

func (harmony_alter *Harmony_alter) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(harmony_alter).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(harmony_alter), uint64(stage.GetOrder(harmony_alter)))
	return
}

func (harp_pedals *Harp_pedals) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(harp_pedals).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(harp_pedals), uint64(stage.GetOrder(harp_pedals)))
	return
}

func (heel_toe *Heel_toe) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(heel_toe).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(heel_toe), uint64(stage.GetOrder(heel_toe)))
	return
}

func (hole *Hole) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(hole).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(hole), uint64(stage.GetOrder(hole)))
	return
}

func (hole_closed *Hole_closed) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(hole_closed).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(hole_closed), uint64(stage.GetOrder(hole_closed)))
	return
}

func (horizontal_turn *Horizontal_turn) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(horizontal_turn).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(horizontal_turn), uint64(stage.GetOrder(horizontal_turn)))
	return
}

func (identification *Identification) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(identification).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(identification), uint64(stage.GetOrder(identification)))
	return
}

func (image *Image) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(image).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(image), uint64(stage.GetOrder(image)))
	return
}

func (instrument *Instrument) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(instrument).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instrument), uint64(stage.GetOrder(instrument)))
	return
}

func (instrument_change *Instrument_change) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(instrument_change).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instrument_change), uint64(stage.GetOrder(instrument_change)))
	return
}

func (instrument_link *Instrument_link) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(instrument_link).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instrument_link), uint64(stage.GetOrder(instrument_link)))
	return
}

func (interchangeable *Interchangeable) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(interchangeable).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(interchangeable), uint64(stage.GetOrder(interchangeable)))
	return
}

func (inversion *Inversion) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(inversion).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(inversion), uint64(stage.GetOrder(inversion)))
	return
}

func (key *Key) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(key).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(key), uint64(stage.GetOrder(key)))
	return
}

func (key_accidental *Key_accidental) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(key_accidental).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(key_accidental), uint64(stage.GetOrder(key_accidental)))
	return
}

func (key_octave *Key_octave) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(key_octave).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(key_octave), uint64(stage.GetOrder(key_octave)))
	return
}

func (kind *Kind) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(kind).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(kind), uint64(stage.GetOrder(kind)))
	return
}

func (level *Level) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(level).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(level), uint64(stage.GetOrder(level)))
	return
}

func (line_detail *Line_detail) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(line_detail).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(line_detail), uint64(stage.GetOrder(line_detail)))
	return
}

func (line_width *Line_width) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(line_width).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(line_width), uint64(stage.GetOrder(line_width)))
	return
}

func (link *Link) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(link).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(link), uint64(stage.GetOrder(link)))
	return
}

func (listen *Listen) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(listen).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(listen), uint64(stage.GetOrder(listen)))
	return
}

func (listening *Listening) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(listening).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(listening), uint64(stage.GetOrder(listening)))
	return
}

func (lyric *Lyric) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(lyric).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(lyric), uint64(stage.GetOrder(lyric)))
	return
}

func (lyric_font *Lyric_font) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(lyric_font).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(lyric_font), uint64(stage.GetOrder(lyric_font)))
	return
}

func (lyric_language *Lyric_language) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(lyric_language).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(lyric_language), uint64(stage.GetOrder(lyric_language)))
	return
}

func (measure_layout *Measure_layout) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(measure_layout).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(measure_layout), uint64(stage.GetOrder(measure_layout)))
	return
}

func (measure_numbering *Measure_numbering) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(measure_numbering).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(measure_numbering), uint64(stage.GetOrder(measure_numbering)))
	return
}

func (measure_repeat *Measure_repeat) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(measure_repeat).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(measure_repeat), uint64(stage.GetOrder(measure_repeat)))
	return
}

func (measure_style *Measure_style) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(measure_style).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(measure_style), uint64(stage.GetOrder(measure_style)))
	return
}

func (membrane *Membrane) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(membrane).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(membrane), uint64(stage.GetOrder(membrane)))
	return
}

func (metal *Metal) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(metal).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(metal), uint64(stage.GetOrder(metal)))
	return
}

func (metronome *Metronome) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(metronome).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(metronome), uint64(stage.GetOrder(metronome)))
	return
}

func (metronome_beam *Metronome_beam) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(metronome_beam).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(metronome_beam), uint64(stage.GetOrder(metronome_beam)))
	return
}

func (metronome_note *Metronome_note) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(metronome_note).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(metronome_note), uint64(stage.GetOrder(metronome_note)))
	return
}

func (metronome_tied *Metronome_tied) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(metronome_tied).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(metronome_tied), uint64(stage.GetOrder(metronome_tied)))
	return
}

func (metronome_tuplet *Metronome_tuplet) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(metronome_tuplet).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(metronome_tuplet), uint64(stage.GetOrder(metronome_tuplet)))
	return
}

func (midi_device *Midi_device) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(midi_device).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(midi_device), uint64(stage.GetOrder(midi_device)))
	return
}

func (midi_instrument *Midi_instrument) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(midi_instrument).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(midi_instrument), uint64(stage.GetOrder(midi_instrument)))
	return
}

func (miscellaneous *Miscellaneous) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(miscellaneous).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(miscellaneous), uint64(stage.GetOrder(miscellaneous)))
	return
}

func (miscellaneous_field *Miscellaneous_field) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(miscellaneous_field).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(miscellaneous_field), uint64(stage.GetOrder(miscellaneous_field)))
	return
}

func (mordent *Mordent) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(mordent).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(mordent), uint64(stage.GetOrder(mordent)))
	return
}

func (multiple_rest *Multiple_rest) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(multiple_rest).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(multiple_rest), uint64(stage.GetOrder(multiple_rest)))
	return
}

func (name_display *Name_display) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(name_display).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(name_display), uint64(stage.GetOrder(name_display)))
	return
}

func (non_arpeggiate *Non_arpeggiate) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(non_arpeggiate).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(non_arpeggiate), uint64(stage.GetOrder(non_arpeggiate)))
	return
}

func (notations *Notations) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(notations).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(notations), uint64(stage.GetOrder(notations)))
	return
}

func (note *Note) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(note).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(note), uint64(stage.GetOrder(note)))
	return
}

func (note_size *Note_size) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(note_size).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(note_size), uint64(stage.GetOrder(note_size)))
	return
}

func (note_type *Note_type) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(note_type).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(note_type), uint64(stage.GetOrder(note_type)))
	return
}

func (notehead *Notehead) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(notehead).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(notehead), uint64(stage.GetOrder(notehead)))
	return
}

func (notehead_text *Notehead_text) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(notehead_text).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(notehead_text), uint64(stage.GetOrder(notehead_text)))
	return
}

func (numeral *Numeral) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(numeral).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(numeral), uint64(stage.GetOrder(numeral)))
	return
}

func (numeral_key *Numeral_key) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(numeral_key).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(numeral_key), uint64(stage.GetOrder(numeral_key)))
	return
}

func (numeral_root *Numeral_root) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(numeral_root).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(numeral_root), uint64(stage.GetOrder(numeral_root)))
	return
}

func (octave_shift *Octave_shift) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(octave_shift).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(octave_shift), uint64(stage.GetOrder(octave_shift)))
	return
}

func (offset *Offset) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(offset).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(offset), uint64(stage.GetOrder(offset)))
	return
}

func (opus *Opus) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(opus).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(opus), uint64(stage.GetOrder(opus)))
	return
}

func (ornaments *Ornaments) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(ornaments).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(ornaments), uint64(stage.GetOrder(ornaments)))
	return
}

func (other_appearance *Other_appearance) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(other_appearance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(other_appearance), uint64(stage.GetOrder(other_appearance)))
	return
}

func (other_direction *Other_direction) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(other_direction).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(other_direction), uint64(stage.GetOrder(other_direction)))
	return
}

func (other_listening *Other_listening) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(other_listening).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(other_listening), uint64(stage.GetOrder(other_listening)))
	return
}

func (other_notation *Other_notation) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(other_notation).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(other_notation), uint64(stage.GetOrder(other_notation)))
	return
}

func (other_placement_text *Other_placement_text) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(other_placement_text).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(other_placement_text), uint64(stage.GetOrder(other_placement_text)))
	return
}

func (other_play *Other_play) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(other_play).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(other_play), uint64(stage.GetOrder(other_play)))
	return
}

func (other_text *Other_text) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(other_text).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(other_text), uint64(stage.GetOrder(other_text)))
	return
}

func (page_layout *Page_layout) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(page_layout).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(page_layout), uint64(stage.GetOrder(page_layout)))
	return
}

func (page_margins *Page_margins) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(page_margins).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(page_margins), uint64(stage.GetOrder(page_margins)))
	return
}

func (part_clef *Part_clef) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(part_clef).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(part_clef), uint64(stage.GetOrder(part_clef)))
	return
}

func (part_group *Part_group) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(part_group).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(part_group), uint64(stage.GetOrder(part_group)))
	return
}

func (part_link *Part_link) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(part_link).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(part_link), uint64(stage.GetOrder(part_link)))
	return
}

func (part_list *Part_list) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(part_list).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(part_list), uint64(stage.GetOrder(part_list)))
	return
}

func (part_name *Part_name) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(part_name).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(part_name), uint64(stage.GetOrder(part_name)))
	return
}

func (part_symbol *Part_symbol) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(part_symbol).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(part_symbol), uint64(stage.GetOrder(part_symbol)))
	return
}

func (part_transpose *Part_transpose) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(part_transpose).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(part_transpose), uint64(stage.GetOrder(part_transpose)))
	return
}

func (pedal *Pedal) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(pedal).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(pedal), uint64(stage.GetOrder(pedal)))
	return
}

func (pedal_tuning *Pedal_tuning) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(pedal_tuning).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(pedal_tuning), uint64(stage.GetOrder(pedal_tuning)))
	return
}

func (per_minute *Per_minute) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(per_minute).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(per_minute), uint64(stage.GetOrder(per_minute)))
	return
}

func (percussion *Percussion) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(percussion).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(percussion), uint64(stage.GetOrder(percussion)))
	return
}

func (pitch *Pitch) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(pitch).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(pitch), uint64(stage.GetOrder(pitch)))
	return
}

func (pitched *Pitched) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(pitched).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(pitched), uint64(stage.GetOrder(pitched)))
	return
}

func (placement_text *Placement_text) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(placement_text).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(placement_text), uint64(stage.GetOrder(placement_text)))
	return
}

func (play *Play) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(play).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(play), uint64(stage.GetOrder(play)))
	return
}

func (player *Player) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(player).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(player), uint64(stage.GetOrder(player)))
	return
}

func (principal_voice *Principal_voice) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(principal_voice).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(principal_voice), uint64(stage.GetOrder(principal_voice)))
	return
}

func (print *Print) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(print).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(print), uint64(stage.GetOrder(print)))
	return
}

func (release *Release) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(release).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(release), uint64(stage.GetOrder(release)))
	return
}

func (repeat *Repeat) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(repeat).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(repeat), uint64(stage.GetOrder(repeat)))
	return
}

func (rest *Rest) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(rest).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(rest), uint64(stage.GetOrder(rest)))
	return
}

func (root *Root) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(root).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(root), uint64(stage.GetOrder(root)))
	return
}

func (root_step *Root_step) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(root_step).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(root_step), uint64(stage.GetOrder(root_step)))
	return
}

func (scaling *Scaling) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(scaling).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(scaling), uint64(stage.GetOrder(scaling)))
	return
}

func (scordatura *Scordatura) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(scordatura).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(scordatura), uint64(stage.GetOrder(scordatura)))
	return
}

func (score_instrument *Score_instrument) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(score_instrument).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(score_instrument), uint64(stage.GetOrder(score_instrument)))
	return
}

func (score_part *Score_part) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(score_part).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(score_part), uint64(stage.GetOrder(score_part)))
	return
}

func (score_partwise *Score_partwise) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(score_partwise).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(score_partwise), uint64(stage.GetOrder(score_partwise)))
	return
}

func (score_timewise *Score_timewise) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(score_timewise).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(score_timewise), uint64(stage.GetOrder(score_timewise)))
	return
}

func (segno *Segno) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(segno).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(segno), uint64(stage.GetOrder(segno)))
	return
}

func (slash *Slash) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(slash).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(slash), uint64(stage.GetOrder(slash)))
	return
}

func (slide *Slide) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(slide).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(slide), uint64(stage.GetOrder(slide)))
	return
}

func (slur *Slur) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(slur).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(slur), uint64(stage.GetOrder(slur)))
	return
}

func (sound *Sound) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(sound).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(sound), uint64(stage.GetOrder(sound)))
	return
}

func (staff_details *Staff_details) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(staff_details).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(staff_details), uint64(stage.GetOrder(staff_details)))
	return
}

func (staff_divide *Staff_divide) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(staff_divide).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(staff_divide), uint64(stage.GetOrder(staff_divide)))
	return
}

func (staff_layout *Staff_layout) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(staff_layout).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(staff_layout), uint64(stage.GetOrder(staff_layout)))
	return
}

func (staff_size *Staff_size) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(staff_size).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(staff_size), uint64(stage.GetOrder(staff_size)))
	return
}

func (staff_tuning *Staff_tuning) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(staff_tuning).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(staff_tuning), uint64(stage.GetOrder(staff_tuning)))
	return
}

func (stem *Stem) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stem).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stem), uint64(stage.GetOrder(stem)))
	return
}

func (stick *Stick) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(stick).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(stick), uint64(stage.GetOrder(stick)))
	return
}

func (string_mute *String_mute) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(string_mute).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(string_mute), uint64(stage.GetOrder(string_mute)))
	return
}

func (string_type *String_type) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(string_type).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(string_type), uint64(stage.GetOrder(string_type)))
	return
}

func (strong_accent *Strong_accent) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(strong_accent).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(strong_accent), uint64(stage.GetOrder(strong_accent)))
	return
}

func (style_text *Style_text) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(style_text).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(style_text), uint64(stage.GetOrder(style_text)))
	return
}

func (supports *Supports) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(supports).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(supports), uint64(stage.GetOrder(supports)))
	return
}

func (swing *Swing) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(swing).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(swing), uint64(stage.GetOrder(swing)))
	return
}

func (sync *Sync) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(sync).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(sync), uint64(stage.GetOrder(sync)))
	return
}

func (system_dividers *System_dividers) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(system_dividers).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(system_dividers), uint64(stage.GetOrder(system_dividers)))
	return
}

func (system_layout *System_layout) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(system_layout).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(system_layout), uint64(stage.GetOrder(system_layout)))
	return
}

func (system_margins *System_margins) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(system_margins).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(system_margins), uint64(stage.GetOrder(system_margins)))
	return
}

func (tap *Tap) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tap).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tap), uint64(stage.GetOrder(tap)))
	return
}

func (technical *Technical) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(technical).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(technical), uint64(stage.GetOrder(technical)))
	return
}

func (text_element_data *Text_element_data) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(text_element_data).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(text_element_data), uint64(stage.GetOrder(text_element_data)))
	return
}

func (tie *Tie) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tie).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tie), uint64(stage.GetOrder(tie)))
	return
}

func (tied *Tied) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tied).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tied), uint64(stage.GetOrder(tied)))
	return
}

func (time *Time) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(time).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(time), uint64(stage.GetOrder(time)))
	return
}

func (time_modification *Time_modification) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(time_modification).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(time_modification), uint64(stage.GetOrder(time_modification)))
	return
}

func (timpani *Timpani) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(timpani).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(timpani), uint64(stage.GetOrder(timpani)))
	return
}

func (transpose *Transpose) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(transpose).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(transpose), uint64(stage.GetOrder(transpose)))
	return
}

func (tremolo *Tremolo) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tremolo).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tremolo), uint64(stage.GetOrder(tremolo)))
	return
}

func (tuplet *Tuplet) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tuplet).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tuplet), uint64(stage.GetOrder(tuplet)))
	return
}

func (tuplet_dot *Tuplet_dot) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tuplet_dot).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tuplet_dot), uint64(stage.GetOrder(tuplet_dot)))
	return
}

func (tuplet_number *Tuplet_number) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tuplet_number).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tuplet_number), uint64(stage.GetOrder(tuplet_number)))
	return
}

func (tuplet_portion *Tuplet_portion) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tuplet_portion).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tuplet_portion), uint64(stage.GetOrder(tuplet_portion)))
	return
}

func (tuplet_type *Tuplet_type) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(tuplet_type).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(tuplet_type), uint64(stage.GetOrder(tuplet_type)))
	return
}

func (typed_text *Typed_text) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(typed_text).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(typed_text), uint64(stage.GetOrder(typed_text)))
	return
}

func (unpitched *Unpitched) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(unpitched).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(unpitched), uint64(stage.GetOrder(unpitched)))
	return
}

func (virtual_instrument *Virtual_instrument) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(virtual_instrument).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(virtual_instrument), uint64(stage.GetOrder(virtual_instrument)))
	return
}

func (wait *Wait) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(wait).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(wait), uint64(stage.GetOrder(wait)))
	return
}

func (wavy_line *Wavy_line) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(wavy_line).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(wavy_line), uint64(stage.GetOrder(wavy_line)))
	return
}

func (wedge *Wedge) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(wedge).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(wedge), uint64(stage.GetOrder(wedge)))
	return
}

func (wood *Wood) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(wood).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(wood), uint64(stage.GetOrder(wood)))
	return
}

func (work *Work) GongGetUUID(stage *Stage) (uuid string) {

	if __gong__, ok := any(work).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}

	uuid = GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(work), uint64(stage.GetOrder(work)))
	return
}


type GongstructDiffable[T any] interface {
	GongstructPtr
	GongMarshallIdentifier(stage *Stage) string
	GongMarshallUnstaging(stage *Stage) string
	GongMarshallAllFields(stage *Stage) (string, string)
	GongReconstructPointersFromInstances(stage *Stage)
	GongDiff(stage *Stage, other T) []string
}

func computeCommitsForType[T GongstructDiffable[T]](
	stage *Stage,
	stagedInstances map[T]struct{},
	stagedOrder map[T]uint,
	referenceInstances map[T]T,
	referenceOrder *map[T]uint,
	instancesMap map[T]T,
	newInstancesSlice *[]string,
	fieldsEditSlice *[]string,
	deletedInstancesSlice *[]string,
	newInstancesReverseSlice *[]string,
	fieldsEditReverseSlice *[]string,
	deletedInstancesReverseSlice *[]string,
	lenNewInstances *int,
	lenDeletedInstances *int,
	lenModifiedInstances *int,
) {
	var newInstances []T
	var deletedInstances []T

	// parse all staged instances and check if they have a reference
	for instance := range stagedInstances {
		if ref, ok := referenceInstances[instance]; !ok {
			newInstances = append(newInstances, instance)
			*newInstancesSlice = append(*newInstancesSlice, instance.GongMarshallIdentifier(stage))
			if *referenceOrder == nil {
				*referenceOrder = make(map[T]uint)
			}
			(*referenceOrder)[instance] = stagedOrder[instance]
			*newInstancesReverseSlice = append(*newInstancesReverseSlice, instance.GongMarshallUnstaging(stage))
			fieldInitializers, pointersInitializations := instance.GongMarshallAllFields(stage)
			*fieldsEditSlice = append(*fieldsEditSlice, fieldInitializers+pointersInitializations)
		} else {
			stagedOrder[ref] = stagedOrder[instance]
			ref.GongReconstructPointersFromInstances(stage) // reconstruct ref with pointers from the stage
			diffs := instance.GongDiff(stage, ref)
			reverseDiffs := ref.GongDiff(stage, instance)
			if len(diffs) > 0 {
				var fieldsEdit string
				if instance.GetName() != "" {
					fieldsEdit += fmt.Sprintf("\n\t// %s", instance.GetName())
				} else {
					fieldsEdit += "\n\t//"
				}
				for _, diff := range diffs {
					fieldsEdit += diff
				}
				*fieldsEditSlice = append(*fieldsEditSlice, fieldsEdit)
				for _, reverseDiff := range reverseDiffs {
					*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, reverseDiff)
				}
				*lenModifiedInstances++
			}
		}
	}

	// parse all reference instances and check if they are still staged
	for _, ref := range referenceInstances {
		instance := instancesMap[ref] // get the instance corresponding to the reference
		if _, ok := stagedInstances[instance]; !ok { // if the instance is not staged anymore, it means it has been unstaged
			deletedInstances = append(deletedInstances, ref)
			*deletedInstancesSlice = append(*deletedInstancesSlice, ref.GongMarshallUnstaging(stage))
			*deletedInstancesReverseSlice = append(*deletedInstancesReverseSlice, ref.GongMarshallIdentifier(stage))
			fieldInitializers, pointersInitializations := ref.GongMarshallAllFields(stage)
			*fieldsEditReverseSlice = append(*fieldsEditReverseSlice, fieldInitializers+pointersInitializations)
		}
	}

	*lenNewInstances += len(newInstances)
	*lenDeletedInstances += len(deletedInstances)
}

func (stage *Stage) ComputeForwardAndBackwardCommits() {
	var lenNewInstances int
	var lenModifiedInstances int
	var lenDeletedInstances int

	var newInstancesSlice []string
	var fieldsEditSlice []string
	var deletedInstancesSlice []string

	var newInstancesReverseSlice []string
	var fieldsEditReverseSlice []string
	var deletedInstancesReverseSlice []string

	// first clean the staging area to remove non staged instances
	// from pointers fields and slices of pointers fields
	stage.Clean()

	// insertion point per named struct
	computeCommitsForType(
		stage,
		stage.A_directives,
		stage.A_directive_stagedOrder,
		stage.A_directives_reference,
		&stage.A_directives_referenceOrder,
		stage.A_directives_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_measures,
		stage.A_measure_stagedOrder,
		stage.A_measures_reference,
		&stage.A_measures_referenceOrder,
		stage.A_measures_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_measure_1s,
		stage.A_measure_1_stagedOrder,
		stage.A_measure_1s_reference,
		&stage.A_measure_1s_referenceOrder,
		stage.A_measure_1s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_parts,
		stage.A_part_stagedOrder,
		stage.A_parts_reference,
		&stage.A_parts_referenceOrder,
		stage.A_parts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.A_part_1s,
		stage.A_part_1_stagedOrder,
		stage.A_part_1s_reference,
		&stage.A_part_1s_referenceOrder,
		stage.A_part_1s_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Accidentals,
		stage.Accidental_stagedOrder,
		stage.Accidentals_reference,
		&stage.Accidentals_referenceOrder,
		stage.Accidentals_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Accidental_marks,
		stage.Accidental_mark_stagedOrder,
		stage.Accidental_marks_reference,
		&stage.Accidental_marks_referenceOrder,
		stage.Accidental_marks_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Accidental_texts,
		stage.Accidental_text_stagedOrder,
		stage.Accidental_texts_reference,
		&stage.Accidental_texts_referenceOrder,
		stage.Accidental_texts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Accords,
		stage.Accord_stagedOrder,
		stage.Accords_reference,
		&stage.Accords_referenceOrder,
		stage.Accords_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Accordion_registrations,
		stage.Accordion_registration_stagedOrder,
		stage.Accordion_registrations_reference,
		&stage.Accordion_registrations_referenceOrder,
		stage.Accordion_registrations_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Appearances,
		stage.Appearance_stagedOrder,
		stage.Appearances_reference,
		&stage.Appearances_referenceOrder,
		stage.Appearances_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Arpeggiates,
		stage.Arpeggiate_stagedOrder,
		stage.Arpeggiates_reference,
		&stage.Arpeggiates_referenceOrder,
		stage.Arpeggiates_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Arrows,
		stage.Arrow_stagedOrder,
		stage.Arrows_reference,
		&stage.Arrows_referenceOrder,
		stage.Arrows_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Articulationss,
		stage.Articulations_stagedOrder,
		stage.Articulationss_reference,
		&stage.Articulationss_referenceOrder,
		stage.Articulationss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Assesss,
		stage.Assess_stagedOrder,
		stage.Assesss_reference,
		&stage.Assesss_referenceOrder,
		stage.Assesss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Attributess,
		stage.Attributes_stagedOrder,
		stage.Attributess_reference,
		&stage.Attributess_referenceOrder,
		stage.Attributess_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Backups,
		stage.Backup_stagedOrder,
		stage.Backups_reference,
		&stage.Backups_referenceOrder,
		stage.Backups_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Bar_style_colors,
		stage.Bar_style_color_stagedOrder,
		stage.Bar_style_colors_reference,
		&stage.Bar_style_colors_referenceOrder,
		stage.Bar_style_colors_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Barlines,
		stage.Barline_stagedOrder,
		stage.Barlines_reference,
		&stage.Barlines_referenceOrder,
		stage.Barlines_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Barres,
		stage.Barre_stagedOrder,
		stage.Barres_reference,
		&stage.Barres_referenceOrder,
		stage.Barres_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Basss,
		stage.Bass_stagedOrder,
		stage.Basss_reference,
		&stage.Basss_referenceOrder,
		stage.Basss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Bass_steps,
		stage.Bass_step_stagedOrder,
		stage.Bass_steps_reference,
		&stage.Bass_steps_referenceOrder,
		stage.Bass_steps_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Beams,
		stage.Beam_stagedOrder,
		stage.Beams_reference,
		&stage.Beams_referenceOrder,
		stage.Beams_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Beat_repeats,
		stage.Beat_repeat_stagedOrder,
		stage.Beat_repeats_reference,
		&stage.Beat_repeats_referenceOrder,
		stage.Beat_repeats_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Beat_unit_tieds,
		stage.Beat_unit_tied_stagedOrder,
		stage.Beat_unit_tieds_reference,
		&stage.Beat_unit_tieds_referenceOrder,
		stage.Beat_unit_tieds_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Beaters,
		stage.Beater_stagedOrder,
		stage.Beaters_reference,
		&stage.Beaters_referenceOrder,
		stage.Beaters_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Bends,
		stage.Bend_stagedOrder,
		stage.Bends_reference,
		&stage.Bends_referenceOrder,
		stage.Bends_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Bookmarks,
		stage.Bookmark_stagedOrder,
		stage.Bookmarks_reference,
		&stage.Bookmarks_referenceOrder,
		stage.Bookmarks_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Brackets,
		stage.Bracket_stagedOrder,
		stage.Brackets_reference,
		&stage.Brackets_referenceOrder,
		stage.Brackets_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Breath_marks,
		stage.Breath_mark_stagedOrder,
		stage.Breath_marks_reference,
		&stage.Breath_marks_referenceOrder,
		stage.Breath_marks_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Caesuras,
		stage.Caesura_stagedOrder,
		stage.Caesuras_reference,
		&stage.Caesuras_referenceOrder,
		stage.Caesuras_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Cancels,
		stage.Cancel_stagedOrder,
		stage.Cancels_reference,
		&stage.Cancels_referenceOrder,
		stage.Cancels_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Clefs,
		stage.Clef_stagedOrder,
		stage.Clefs_reference,
		&stage.Clefs_referenceOrder,
		stage.Clefs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Codas,
		stage.Coda_stagedOrder,
		stage.Codas_reference,
		&stage.Codas_referenceOrder,
		stage.Codas_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Credits,
		stage.Credit_stagedOrder,
		stage.Credits_reference,
		&stage.Credits_referenceOrder,
		stage.Credits_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Dashess,
		stage.Dashes_stagedOrder,
		stage.Dashess_reference,
		&stage.Dashess_referenceOrder,
		stage.Dashess_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Defaultss,
		stage.Defaults_stagedOrder,
		stage.Defaultss_reference,
		&stage.Defaultss_referenceOrder,
		stage.Defaultss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Degrees,
		stage.Degree_stagedOrder,
		stage.Degrees_reference,
		&stage.Degrees_referenceOrder,
		stage.Degrees_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Degree_alters,
		stage.Degree_alter_stagedOrder,
		stage.Degree_alters_reference,
		&stage.Degree_alters_referenceOrder,
		stage.Degree_alters_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Degree_types,
		stage.Degree_type_stagedOrder,
		stage.Degree_types_reference,
		&stage.Degree_types_referenceOrder,
		stage.Degree_types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Degree_values,
		stage.Degree_value_stagedOrder,
		stage.Degree_values_reference,
		&stage.Degree_values_referenceOrder,
		stage.Degree_values_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Directions,
		stage.Direction_stagedOrder,
		stage.Directions_reference,
		&stage.Directions_referenceOrder,
		stage.Directions_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Direction_types,
		stage.Direction_type_stagedOrder,
		stage.Direction_types_reference,
		&stage.Direction_types_referenceOrder,
		stage.Direction_types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Distances,
		stage.Distance_stagedOrder,
		stage.Distances_reference,
		&stage.Distances_referenceOrder,
		stage.Distances_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Doubles,
		stage.Double_stagedOrder,
		stage.Doubles_reference,
		&stage.Doubles_referenceOrder,
		stage.Doubles_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Dynamicss,
		stage.Dynamics_stagedOrder,
		stage.Dynamicss_reference,
		&stage.Dynamicss_referenceOrder,
		stage.Dynamicss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Effects,
		stage.Effect_stagedOrder,
		stage.Effects_reference,
		&stage.Effects_referenceOrder,
		stage.Effects_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Elisions,
		stage.Elision_stagedOrder,
		stage.Elisions_reference,
		&stage.Elisions_referenceOrder,
		stage.Elisions_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Emptys,
		stage.Empty_stagedOrder,
		stage.Emptys_reference,
		&stage.Emptys_referenceOrder,
		stage.Emptys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Empty_fonts,
		stage.Empty_font_stagedOrder,
		stage.Empty_fonts_reference,
		&stage.Empty_fonts_referenceOrder,
		stage.Empty_fonts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Empty_lines,
		stage.Empty_line_stagedOrder,
		stage.Empty_lines_reference,
		&stage.Empty_lines_referenceOrder,
		stage.Empty_lines_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Empty_placements,
		stage.Empty_placement_stagedOrder,
		stage.Empty_placements_reference,
		&stage.Empty_placements_referenceOrder,
		stage.Empty_placements_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Empty_placement_smufls,
		stage.Empty_placement_smufl_stagedOrder,
		stage.Empty_placement_smufls_reference,
		&stage.Empty_placement_smufls_referenceOrder,
		stage.Empty_placement_smufls_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Empty_print_object_style_aligns,
		stage.Empty_print_object_style_align_stagedOrder,
		stage.Empty_print_object_style_aligns_reference,
		&stage.Empty_print_object_style_aligns_referenceOrder,
		stage.Empty_print_object_style_aligns_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Empty_print_styles,
		stage.Empty_print_style_stagedOrder,
		stage.Empty_print_styles_reference,
		&stage.Empty_print_styles_referenceOrder,
		stage.Empty_print_styles_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Empty_print_style_aligns,
		stage.Empty_print_style_align_stagedOrder,
		stage.Empty_print_style_aligns_reference,
		&stage.Empty_print_style_aligns_referenceOrder,
		stage.Empty_print_style_aligns_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Empty_print_style_align_ids,
		stage.Empty_print_style_align_id_stagedOrder,
		stage.Empty_print_style_align_ids_reference,
		&stage.Empty_print_style_align_ids_referenceOrder,
		stage.Empty_print_style_align_ids_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Empty_trill_sounds,
		stage.Empty_trill_sound_stagedOrder,
		stage.Empty_trill_sounds_reference,
		&stage.Empty_trill_sounds_referenceOrder,
		stage.Empty_trill_sounds_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Encodings,
		stage.Encoding_stagedOrder,
		stage.Encodings_reference,
		&stage.Encodings_referenceOrder,
		stage.Encodings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Endings,
		stage.Ending_stagedOrder,
		stage.Endings_reference,
		&stage.Endings_referenceOrder,
		stage.Endings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Extends,
		stage.Extend_stagedOrder,
		stage.Extends_reference,
		&stage.Extends_referenceOrder,
		stage.Extends_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Features,
		stage.Feature_stagedOrder,
		stage.Features_reference,
		&stage.Features_referenceOrder,
		stage.Features_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Fermatas,
		stage.Fermata_stagedOrder,
		stage.Fermatas_reference,
		&stage.Fermatas_referenceOrder,
		stage.Fermatas_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Figures,
		stage.Figure_stagedOrder,
		stage.Figures_reference,
		&stage.Figures_referenceOrder,
		stage.Figures_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Figured_basss,
		stage.Figured_bass_stagedOrder,
		stage.Figured_basss_reference,
		&stage.Figured_basss_referenceOrder,
		stage.Figured_basss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Fingerings,
		stage.Fingering_stagedOrder,
		stage.Fingerings_reference,
		&stage.Fingerings_referenceOrder,
		stage.Fingerings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.First_frets,
		stage.First_fret_stagedOrder,
		stage.First_frets_reference,
		&stage.First_frets_referenceOrder,
		stage.First_frets_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.For_parts,
		stage.For_part_stagedOrder,
		stage.For_parts_reference,
		&stage.For_parts_referenceOrder,
		stage.For_parts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Formatted_symbols,
		stage.Formatted_symbol_stagedOrder,
		stage.Formatted_symbols_reference,
		&stage.Formatted_symbols_referenceOrder,
		stage.Formatted_symbols_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Formatted_symbol_ids,
		stage.Formatted_symbol_id_stagedOrder,
		stage.Formatted_symbol_ids_reference,
		&stage.Formatted_symbol_ids_referenceOrder,
		stage.Formatted_symbol_ids_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Formatted_texts,
		stage.Formatted_text_stagedOrder,
		stage.Formatted_texts_reference,
		&stage.Formatted_texts_referenceOrder,
		stage.Formatted_texts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Formatted_text_ids,
		stage.Formatted_text_id_stagedOrder,
		stage.Formatted_text_ids_reference,
		&stage.Formatted_text_ids_referenceOrder,
		stage.Formatted_text_ids_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Forwards,
		stage.Forward_stagedOrder,
		stage.Forwards_reference,
		&stage.Forwards_referenceOrder,
		stage.Forwards_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Frames,
		stage.Frame_stagedOrder,
		stage.Frames_reference,
		&stage.Frames_referenceOrder,
		stage.Frames_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Frame_notes,
		stage.Frame_note_stagedOrder,
		stage.Frame_notes_reference,
		&stage.Frame_notes_referenceOrder,
		stage.Frame_notes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Frets,
		stage.Fret_stagedOrder,
		stage.Frets_reference,
		&stage.Frets_referenceOrder,
		stage.Frets_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Glasss,
		stage.Glass_stagedOrder,
		stage.Glasss_reference,
		&stage.Glasss_referenceOrder,
		stage.Glasss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Glissandos,
		stage.Glissando_stagedOrder,
		stage.Glissandos_reference,
		&stage.Glissandos_referenceOrder,
		stage.Glissandos_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Glyphs,
		stage.Glyph_stagedOrder,
		stage.Glyphs_reference,
		&stage.Glyphs_referenceOrder,
		stage.Glyphs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Graces,
		stage.Grace_stagedOrder,
		stage.Graces_reference,
		&stage.Graces_referenceOrder,
		stage.Graces_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Group_barlines,
		stage.Group_barline_stagedOrder,
		stage.Group_barlines_reference,
		&stage.Group_barlines_referenceOrder,
		stage.Group_barlines_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Group_names,
		stage.Group_name_stagedOrder,
		stage.Group_names_reference,
		&stage.Group_names_referenceOrder,
		stage.Group_names_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Group_symbols,
		stage.Group_symbol_stagedOrder,
		stage.Group_symbols_reference,
		&stage.Group_symbols_referenceOrder,
		stage.Group_symbols_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Groupings,
		stage.Grouping_stagedOrder,
		stage.Groupings_reference,
		&stage.Groupings_referenceOrder,
		stage.Groupings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Hammer_on_pull_offs,
		stage.Hammer_on_pull_off_stagedOrder,
		stage.Hammer_on_pull_offs_reference,
		&stage.Hammer_on_pull_offs_referenceOrder,
		stage.Hammer_on_pull_offs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Handbells,
		stage.Handbell_stagedOrder,
		stage.Handbells_reference,
		&stage.Handbells_referenceOrder,
		stage.Handbells_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Harmon_closeds,
		stage.Harmon_closed_stagedOrder,
		stage.Harmon_closeds_reference,
		&stage.Harmon_closeds_referenceOrder,
		stage.Harmon_closeds_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Harmon_mutes,
		stage.Harmon_mute_stagedOrder,
		stage.Harmon_mutes_reference,
		&stage.Harmon_mutes_referenceOrder,
		stage.Harmon_mutes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Harmonics,
		stage.Harmonic_stagedOrder,
		stage.Harmonics_reference,
		&stage.Harmonics_referenceOrder,
		stage.Harmonics_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Harmonys,
		stage.Harmony_stagedOrder,
		stage.Harmonys_reference,
		&stage.Harmonys_referenceOrder,
		stage.Harmonys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Harmony_alters,
		stage.Harmony_alter_stagedOrder,
		stage.Harmony_alters_reference,
		&stage.Harmony_alters_referenceOrder,
		stage.Harmony_alters_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Harp_pedalss,
		stage.Harp_pedals_stagedOrder,
		stage.Harp_pedalss_reference,
		&stage.Harp_pedalss_referenceOrder,
		stage.Harp_pedalss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Heel_toes,
		stage.Heel_toe_stagedOrder,
		stage.Heel_toes_reference,
		&stage.Heel_toes_referenceOrder,
		stage.Heel_toes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Holes,
		stage.Hole_stagedOrder,
		stage.Holes_reference,
		&stage.Holes_referenceOrder,
		stage.Holes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Hole_closeds,
		stage.Hole_closed_stagedOrder,
		stage.Hole_closeds_reference,
		&stage.Hole_closeds_referenceOrder,
		stage.Hole_closeds_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Horizontal_turns,
		stage.Horizontal_turn_stagedOrder,
		stage.Horizontal_turns_reference,
		&stage.Horizontal_turns_referenceOrder,
		stage.Horizontal_turns_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Identifications,
		stage.Identification_stagedOrder,
		stage.Identifications_reference,
		&stage.Identifications_referenceOrder,
		stage.Identifications_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Images,
		stage.Image_stagedOrder,
		stage.Images_reference,
		&stage.Images_referenceOrder,
		stage.Images_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Instruments,
		stage.Instrument_stagedOrder,
		stage.Instruments_reference,
		&stage.Instruments_referenceOrder,
		stage.Instruments_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Instrument_changes,
		stage.Instrument_change_stagedOrder,
		stage.Instrument_changes_reference,
		&stage.Instrument_changes_referenceOrder,
		stage.Instrument_changes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Instrument_links,
		stage.Instrument_link_stagedOrder,
		stage.Instrument_links_reference,
		&stage.Instrument_links_referenceOrder,
		stage.Instrument_links_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Interchangeables,
		stage.Interchangeable_stagedOrder,
		stage.Interchangeables_reference,
		&stage.Interchangeables_referenceOrder,
		stage.Interchangeables_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Inversions,
		stage.Inversion_stagedOrder,
		stage.Inversions_reference,
		&stage.Inversions_referenceOrder,
		stage.Inversions_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Keys,
		stage.Key_stagedOrder,
		stage.Keys_reference,
		&stage.Keys_referenceOrder,
		stage.Keys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Key_accidentals,
		stage.Key_accidental_stagedOrder,
		stage.Key_accidentals_reference,
		&stage.Key_accidentals_referenceOrder,
		stage.Key_accidentals_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Key_octaves,
		stage.Key_octave_stagedOrder,
		stage.Key_octaves_reference,
		&stage.Key_octaves_referenceOrder,
		stage.Key_octaves_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Kinds,
		stage.Kind_stagedOrder,
		stage.Kinds_reference,
		&stage.Kinds_referenceOrder,
		stage.Kinds_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Levels,
		stage.Level_stagedOrder,
		stage.Levels_reference,
		&stage.Levels_referenceOrder,
		stage.Levels_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Line_details,
		stage.Line_detail_stagedOrder,
		stage.Line_details_reference,
		&stage.Line_details_referenceOrder,
		stage.Line_details_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Line_widths,
		stage.Line_width_stagedOrder,
		stage.Line_widths_reference,
		&stage.Line_widths_referenceOrder,
		stage.Line_widths_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Links,
		stage.Link_stagedOrder,
		stage.Links_reference,
		&stage.Links_referenceOrder,
		stage.Links_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Listens,
		stage.Listen_stagedOrder,
		stage.Listens_reference,
		&stage.Listens_referenceOrder,
		stage.Listens_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Listenings,
		stage.Listening_stagedOrder,
		stage.Listenings_reference,
		&stage.Listenings_referenceOrder,
		stage.Listenings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Lyrics,
		stage.Lyric_stagedOrder,
		stage.Lyrics_reference,
		&stage.Lyrics_referenceOrder,
		stage.Lyrics_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Lyric_fonts,
		stage.Lyric_font_stagedOrder,
		stage.Lyric_fonts_reference,
		&stage.Lyric_fonts_referenceOrder,
		stage.Lyric_fonts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Lyric_languages,
		stage.Lyric_language_stagedOrder,
		stage.Lyric_languages_reference,
		&stage.Lyric_languages_referenceOrder,
		stage.Lyric_languages_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Measure_layouts,
		stage.Measure_layout_stagedOrder,
		stage.Measure_layouts_reference,
		&stage.Measure_layouts_referenceOrder,
		stage.Measure_layouts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Measure_numberings,
		stage.Measure_numbering_stagedOrder,
		stage.Measure_numberings_reference,
		&stage.Measure_numberings_referenceOrder,
		stage.Measure_numberings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Measure_repeats,
		stage.Measure_repeat_stagedOrder,
		stage.Measure_repeats_reference,
		&stage.Measure_repeats_referenceOrder,
		stage.Measure_repeats_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Measure_styles,
		stage.Measure_style_stagedOrder,
		stage.Measure_styles_reference,
		&stage.Measure_styles_referenceOrder,
		stage.Measure_styles_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Membranes,
		stage.Membrane_stagedOrder,
		stage.Membranes_reference,
		&stage.Membranes_referenceOrder,
		stage.Membranes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Metals,
		stage.Metal_stagedOrder,
		stage.Metals_reference,
		&stage.Metals_referenceOrder,
		stage.Metals_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Metronomes,
		stage.Metronome_stagedOrder,
		stage.Metronomes_reference,
		&stage.Metronomes_referenceOrder,
		stage.Metronomes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Metronome_beams,
		stage.Metronome_beam_stagedOrder,
		stage.Metronome_beams_reference,
		&stage.Metronome_beams_referenceOrder,
		stage.Metronome_beams_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Metronome_notes,
		stage.Metronome_note_stagedOrder,
		stage.Metronome_notes_reference,
		&stage.Metronome_notes_referenceOrder,
		stage.Metronome_notes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Metronome_tieds,
		stage.Metronome_tied_stagedOrder,
		stage.Metronome_tieds_reference,
		&stage.Metronome_tieds_referenceOrder,
		stage.Metronome_tieds_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Metronome_tuplets,
		stage.Metronome_tuplet_stagedOrder,
		stage.Metronome_tuplets_reference,
		&stage.Metronome_tuplets_referenceOrder,
		stage.Metronome_tuplets_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Midi_devices,
		stage.Midi_device_stagedOrder,
		stage.Midi_devices_reference,
		&stage.Midi_devices_referenceOrder,
		stage.Midi_devices_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Midi_instruments,
		stage.Midi_instrument_stagedOrder,
		stage.Midi_instruments_reference,
		&stage.Midi_instruments_referenceOrder,
		stage.Midi_instruments_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Miscellaneouss,
		stage.Miscellaneous_stagedOrder,
		stage.Miscellaneouss_reference,
		&stage.Miscellaneouss_referenceOrder,
		stage.Miscellaneouss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Miscellaneous_fields,
		stage.Miscellaneous_field_stagedOrder,
		stage.Miscellaneous_fields_reference,
		&stage.Miscellaneous_fields_referenceOrder,
		stage.Miscellaneous_fields_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Mordents,
		stage.Mordent_stagedOrder,
		stage.Mordents_reference,
		&stage.Mordents_referenceOrder,
		stage.Mordents_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Multiple_rests,
		stage.Multiple_rest_stagedOrder,
		stage.Multiple_rests_reference,
		&stage.Multiple_rests_referenceOrder,
		stage.Multiple_rests_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Name_displays,
		stage.Name_display_stagedOrder,
		stage.Name_displays_reference,
		&stage.Name_displays_referenceOrder,
		stage.Name_displays_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Non_arpeggiates,
		stage.Non_arpeggiate_stagedOrder,
		stage.Non_arpeggiates_reference,
		&stage.Non_arpeggiates_referenceOrder,
		stage.Non_arpeggiates_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Notationss,
		stage.Notations_stagedOrder,
		stage.Notationss_reference,
		&stage.Notationss_referenceOrder,
		stage.Notationss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Notes,
		stage.Note_stagedOrder,
		stage.Notes_reference,
		&stage.Notes_referenceOrder,
		stage.Notes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Note_sizes,
		stage.Note_size_stagedOrder,
		stage.Note_sizes_reference,
		&stage.Note_sizes_referenceOrder,
		stage.Note_sizes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Note_types,
		stage.Note_type_stagedOrder,
		stage.Note_types_reference,
		&stage.Note_types_referenceOrder,
		stage.Note_types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Noteheads,
		stage.Notehead_stagedOrder,
		stage.Noteheads_reference,
		&stage.Noteheads_referenceOrder,
		stage.Noteheads_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Notehead_texts,
		stage.Notehead_text_stagedOrder,
		stage.Notehead_texts_reference,
		&stage.Notehead_texts_referenceOrder,
		stage.Notehead_texts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Numerals,
		stage.Numeral_stagedOrder,
		stage.Numerals_reference,
		&stage.Numerals_referenceOrder,
		stage.Numerals_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Numeral_keys,
		stage.Numeral_key_stagedOrder,
		stage.Numeral_keys_reference,
		&stage.Numeral_keys_referenceOrder,
		stage.Numeral_keys_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Numeral_roots,
		stage.Numeral_root_stagedOrder,
		stage.Numeral_roots_reference,
		&stage.Numeral_roots_referenceOrder,
		stage.Numeral_roots_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Octave_shifts,
		stage.Octave_shift_stagedOrder,
		stage.Octave_shifts_reference,
		&stage.Octave_shifts_referenceOrder,
		stage.Octave_shifts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Offsets,
		stage.Offset_stagedOrder,
		stage.Offsets_reference,
		&stage.Offsets_referenceOrder,
		stage.Offsets_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Opuss,
		stage.Opus_stagedOrder,
		stage.Opuss_reference,
		&stage.Opuss_referenceOrder,
		stage.Opuss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Ornamentss,
		stage.Ornaments_stagedOrder,
		stage.Ornamentss_reference,
		&stage.Ornamentss_referenceOrder,
		stage.Ornamentss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Other_appearances,
		stage.Other_appearance_stagedOrder,
		stage.Other_appearances_reference,
		&stage.Other_appearances_referenceOrder,
		stage.Other_appearances_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Other_directions,
		stage.Other_direction_stagedOrder,
		stage.Other_directions_reference,
		&stage.Other_directions_referenceOrder,
		stage.Other_directions_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Other_listenings,
		stage.Other_listening_stagedOrder,
		stage.Other_listenings_reference,
		&stage.Other_listenings_referenceOrder,
		stage.Other_listenings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Other_notations,
		stage.Other_notation_stagedOrder,
		stage.Other_notations_reference,
		&stage.Other_notations_referenceOrder,
		stage.Other_notations_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Other_placement_texts,
		stage.Other_placement_text_stagedOrder,
		stage.Other_placement_texts_reference,
		&stage.Other_placement_texts_referenceOrder,
		stage.Other_placement_texts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Other_plays,
		stage.Other_play_stagedOrder,
		stage.Other_plays_reference,
		&stage.Other_plays_referenceOrder,
		stage.Other_plays_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Other_texts,
		stage.Other_text_stagedOrder,
		stage.Other_texts_reference,
		&stage.Other_texts_referenceOrder,
		stage.Other_texts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Page_layouts,
		stage.Page_layout_stagedOrder,
		stage.Page_layouts_reference,
		&stage.Page_layouts_referenceOrder,
		stage.Page_layouts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Page_marginss,
		stage.Page_margins_stagedOrder,
		stage.Page_marginss_reference,
		&stage.Page_marginss_referenceOrder,
		stage.Page_marginss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Part_clefs,
		stage.Part_clef_stagedOrder,
		stage.Part_clefs_reference,
		&stage.Part_clefs_referenceOrder,
		stage.Part_clefs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Part_groups,
		stage.Part_group_stagedOrder,
		stage.Part_groups_reference,
		&stage.Part_groups_referenceOrder,
		stage.Part_groups_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Part_links,
		stage.Part_link_stagedOrder,
		stage.Part_links_reference,
		&stage.Part_links_referenceOrder,
		stage.Part_links_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Part_lists,
		stage.Part_list_stagedOrder,
		stage.Part_lists_reference,
		&stage.Part_lists_referenceOrder,
		stage.Part_lists_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Part_names,
		stage.Part_name_stagedOrder,
		stage.Part_names_reference,
		&stage.Part_names_referenceOrder,
		stage.Part_names_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Part_symbols,
		stage.Part_symbol_stagedOrder,
		stage.Part_symbols_reference,
		&stage.Part_symbols_referenceOrder,
		stage.Part_symbols_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Part_transposes,
		stage.Part_transpose_stagedOrder,
		stage.Part_transposes_reference,
		&stage.Part_transposes_referenceOrder,
		stage.Part_transposes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Pedals,
		stage.Pedal_stagedOrder,
		stage.Pedals_reference,
		&stage.Pedals_referenceOrder,
		stage.Pedals_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Pedal_tunings,
		stage.Pedal_tuning_stagedOrder,
		stage.Pedal_tunings_reference,
		&stage.Pedal_tunings_referenceOrder,
		stage.Pedal_tunings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Per_minutes,
		stage.Per_minute_stagedOrder,
		stage.Per_minutes_reference,
		&stage.Per_minutes_referenceOrder,
		stage.Per_minutes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Percussions,
		stage.Percussion_stagedOrder,
		stage.Percussions_reference,
		&stage.Percussions_referenceOrder,
		stage.Percussions_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Pitchs,
		stage.Pitch_stagedOrder,
		stage.Pitchs_reference,
		&stage.Pitchs_referenceOrder,
		stage.Pitchs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Pitcheds,
		stage.Pitched_stagedOrder,
		stage.Pitcheds_reference,
		&stage.Pitcheds_referenceOrder,
		stage.Pitcheds_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Placement_texts,
		stage.Placement_text_stagedOrder,
		stage.Placement_texts_reference,
		&stage.Placement_texts_referenceOrder,
		stage.Placement_texts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Plays,
		stage.Play_stagedOrder,
		stage.Plays_reference,
		&stage.Plays_referenceOrder,
		stage.Plays_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Players,
		stage.Player_stagedOrder,
		stage.Players_reference,
		&stage.Players_referenceOrder,
		stage.Players_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Principal_voices,
		stage.Principal_voice_stagedOrder,
		stage.Principal_voices_reference,
		&stage.Principal_voices_referenceOrder,
		stage.Principal_voices_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Prints,
		stage.Print_stagedOrder,
		stage.Prints_reference,
		&stage.Prints_referenceOrder,
		stage.Prints_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Releases,
		stage.Release_stagedOrder,
		stage.Releases_reference,
		&stage.Releases_referenceOrder,
		stage.Releases_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Repeats,
		stage.Repeat_stagedOrder,
		stage.Repeats_reference,
		&stage.Repeats_referenceOrder,
		stage.Repeats_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Rests,
		stage.Rest_stagedOrder,
		stage.Rests_reference,
		&stage.Rests_referenceOrder,
		stage.Rests_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Roots,
		stage.Root_stagedOrder,
		stage.Roots_reference,
		&stage.Roots_referenceOrder,
		stage.Roots_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Root_steps,
		stage.Root_step_stagedOrder,
		stage.Root_steps_reference,
		&stage.Root_steps_referenceOrder,
		stage.Root_steps_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Scalings,
		stage.Scaling_stagedOrder,
		stage.Scalings_reference,
		&stage.Scalings_referenceOrder,
		stage.Scalings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Scordaturas,
		stage.Scordatura_stagedOrder,
		stage.Scordaturas_reference,
		&stage.Scordaturas_referenceOrder,
		stage.Scordaturas_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Score_instruments,
		stage.Score_instrument_stagedOrder,
		stage.Score_instruments_reference,
		&stage.Score_instruments_referenceOrder,
		stage.Score_instruments_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Score_parts,
		stage.Score_part_stagedOrder,
		stage.Score_parts_reference,
		&stage.Score_parts_referenceOrder,
		stage.Score_parts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Score_partwises,
		stage.Score_partwise_stagedOrder,
		stage.Score_partwises_reference,
		&stage.Score_partwises_referenceOrder,
		stage.Score_partwises_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Score_timewises,
		stage.Score_timewise_stagedOrder,
		stage.Score_timewises_reference,
		&stage.Score_timewises_referenceOrder,
		stage.Score_timewises_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Segnos,
		stage.Segno_stagedOrder,
		stage.Segnos_reference,
		&stage.Segnos_referenceOrder,
		stage.Segnos_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Slashs,
		stage.Slash_stagedOrder,
		stage.Slashs_reference,
		&stage.Slashs_referenceOrder,
		stage.Slashs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Slides,
		stage.Slide_stagedOrder,
		stage.Slides_reference,
		&stage.Slides_referenceOrder,
		stage.Slides_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Slurs,
		stage.Slur_stagedOrder,
		stage.Slurs_reference,
		&stage.Slurs_referenceOrder,
		stage.Slurs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Sounds,
		stage.Sound_stagedOrder,
		stage.Sounds_reference,
		&stage.Sounds_referenceOrder,
		stage.Sounds_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Staff_detailss,
		stage.Staff_details_stagedOrder,
		stage.Staff_detailss_reference,
		&stage.Staff_detailss_referenceOrder,
		stage.Staff_detailss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Staff_divides,
		stage.Staff_divide_stagedOrder,
		stage.Staff_divides_reference,
		&stage.Staff_divides_referenceOrder,
		stage.Staff_divides_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Staff_layouts,
		stage.Staff_layout_stagedOrder,
		stage.Staff_layouts_reference,
		&stage.Staff_layouts_referenceOrder,
		stage.Staff_layouts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Staff_sizes,
		stage.Staff_size_stagedOrder,
		stage.Staff_sizes_reference,
		&stage.Staff_sizes_referenceOrder,
		stage.Staff_sizes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Staff_tunings,
		stage.Staff_tuning_stagedOrder,
		stage.Staff_tunings_reference,
		&stage.Staff_tunings_referenceOrder,
		stage.Staff_tunings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Stems,
		stage.Stem_stagedOrder,
		stage.Stems_reference,
		&stage.Stems_referenceOrder,
		stage.Stems_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Sticks,
		stage.Stick_stagedOrder,
		stage.Sticks_reference,
		&stage.Sticks_referenceOrder,
		stage.Sticks_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.String_mutes,
		stage.String_mute_stagedOrder,
		stage.String_mutes_reference,
		&stage.String_mutes_referenceOrder,
		stage.String_mutes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.String_types,
		stage.String_type_stagedOrder,
		stage.String_types_reference,
		&stage.String_types_referenceOrder,
		stage.String_types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Strong_accents,
		stage.Strong_accent_stagedOrder,
		stage.Strong_accents_reference,
		&stage.Strong_accents_referenceOrder,
		stage.Strong_accents_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Style_texts,
		stage.Style_text_stagedOrder,
		stage.Style_texts_reference,
		&stage.Style_texts_referenceOrder,
		stage.Style_texts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Supportss,
		stage.Supports_stagedOrder,
		stage.Supportss_reference,
		&stage.Supportss_referenceOrder,
		stage.Supportss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Swings,
		stage.Swing_stagedOrder,
		stage.Swings_reference,
		&stage.Swings_referenceOrder,
		stage.Swings_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Syncs,
		stage.Sync_stagedOrder,
		stage.Syncs_reference,
		&stage.Syncs_referenceOrder,
		stage.Syncs_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.System_dividerss,
		stage.System_dividers_stagedOrder,
		stage.System_dividerss_reference,
		&stage.System_dividerss_referenceOrder,
		stage.System_dividerss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.System_layouts,
		stage.System_layout_stagedOrder,
		stage.System_layouts_reference,
		&stage.System_layouts_referenceOrder,
		stage.System_layouts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.System_marginss,
		stage.System_margins_stagedOrder,
		stage.System_marginss_reference,
		&stage.System_marginss_referenceOrder,
		stage.System_marginss_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Taps,
		stage.Tap_stagedOrder,
		stage.Taps_reference,
		&stage.Taps_referenceOrder,
		stage.Taps_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Technicals,
		stage.Technical_stagedOrder,
		stage.Technicals_reference,
		&stage.Technicals_referenceOrder,
		stage.Technicals_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Text_element_datas,
		stage.Text_element_data_stagedOrder,
		stage.Text_element_datas_reference,
		&stage.Text_element_datas_referenceOrder,
		stage.Text_element_datas_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Ties,
		stage.Tie_stagedOrder,
		stage.Ties_reference,
		&stage.Ties_referenceOrder,
		stage.Ties_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Tieds,
		stage.Tied_stagedOrder,
		stage.Tieds_reference,
		&stage.Tieds_referenceOrder,
		stage.Tieds_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Times,
		stage.Time_stagedOrder,
		stage.Times_reference,
		&stage.Times_referenceOrder,
		stage.Times_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Time_modifications,
		stage.Time_modification_stagedOrder,
		stage.Time_modifications_reference,
		&stage.Time_modifications_referenceOrder,
		stage.Time_modifications_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Timpanis,
		stage.Timpani_stagedOrder,
		stage.Timpanis_reference,
		&stage.Timpanis_referenceOrder,
		stage.Timpanis_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Transposes,
		stage.Transpose_stagedOrder,
		stage.Transposes_reference,
		&stage.Transposes_referenceOrder,
		stage.Transposes_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Tremolos,
		stage.Tremolo_stagedOrder,
		stage.Tremolos_reference,
		&stage.Tremolos_referenceOrder,
		stage.Tremolos_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Tuplets,
		stage.Tuplet_stagedOrder,
		stage.Tuplets_reference,
		&stage.Tuplets_referenceOrder,
		stage.Tuplets_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Tuplet_dots,
		stage.Tuplet_dot_stagedOrder,
		stage.Tuplet_dots_reference,
		&stage.Tuplet_dots_referenceOrder,
		stage.Tuplet_dots_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Tuplet_numbers,
		stage.Tuplet_number_stagedOrder,
		stage.Tuplet_numbers_reference,
		&stage.Tuplet_numbers_referenceOrder,
		stage.Tuplet_numbers_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Tuplet_portions,
		stage.Tuplet_portion_stagedOrder,
		stage.Tuplet_portions_reference,
		&stage.Tuplet_portions_referenceOrder,
		stage.Tuplet_portions_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Tuplet_types,
		stage.Tuplet_type_stagedOrder,
		stage.Tuplet_types_reference,
		&stage.Tuplet_types_referenceOrder,
		stage.Tuplet_types_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Typed_texts,
		stage.Typed_text_stagedOrder,
		stage.Typed_texts_reference,
		&stage.Typed_texts_referenceOrder,
		stage.Typed_texts_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Unpitcheds,
		stage.Unpitched_stagedOrder,
		stage.Unpitcheds_reference,
		&stage.Unpitcheds_referenceOrder,
		stage.Unpitcheds_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Virtual_instruments,
		stage.Virtual_instrument_stagedOrder,
		stage.Virtual_instruments_reference,
		&stage.Virtual_instruments_referenceOrder,
		stage.Virtual_instruments_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Waits,
		stage.Wait_stagedOrder,
		stage.Waits_reference,
		&stage.Waits_referenceOrder,
		stage.Waits_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Wavy_lines,
		stage.Wavy_line_stagedOrder,
		stage.Wavy_lines_reference,
		&stage.Wavy_lines_referenceOrder,
		stage.Wavy_lines_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Wedges,
		stage.Wedge_stagedOrder,
		stage.Wedges_reference,
		&stage.Wedges_referenceOrder,
		stage.Wedges_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Woods,
		stage.Wood_stagedOrder,
		stage.Woods_reference,
		&stage.Woods_referenceOrder,
		stage.Woods_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)
	computeCommitsForType(
		stage,
		stage.Works,
		stage.Work_stagedOrder,
		stage.Works_reference,
		&stage.Works_referenceOrder,
		stage.Works_instance,
		&newInstancesSlice,
		&fieldsEditSlice,
		&deletedInstancesSlice,
		&newInstancesReverseSlice,
		&fieldsEditReverseSlice,
		&deletedInstancesReverseSlice,
		&lenNewInstances,
		&lenDeletedInstances,
		&lenModifiedInstances,
	)

	if lenNewInstances > 0 || lenDeletedInstances > 0 || lenModifiedInstances > 0 {

		// sort the stmt to have reproductible forward/backward commit
		sort.Strings(newInstancesSlice)
		newInstancesStmt := strings.Join(newInstancesSlice, "")
		sort.Strings(fieldsEditSlice)
		fieldsEditStmt := strings.Join(fieldsEditSlice, "")
		sort.Strings(deletedInstancesSlice)
		deletedInstancesStmt := strings.Join(deletedInstancesSlice, "")

		sort.Strings(newInstancesReverseSlice)
		newInstancesReverseStmt := strings.Join(newInstancesReverseSlice, "")
		sort.Strings(fieldsEditReverseSlice)
		fieldsEditReverseStmt := strings.Join(fieldsEditReverseSlice, "")
		sort.Strings(deletedInstancesReverseSlice)
		deletedInstancesReverseStmt := strings.Join(deletedInstancesReverseSlice, "")

		forwardCommit := newInstancesStmt + fieldsEditStmt + deletedInstancesStmt
		forwardCommit += "\n\tstage.Commit()"
		stage.forwardCommits = append(stage.forwardCommits, forwardCommit)

		backwardCommit := deletedInstancesReverseStmt + fieldsEditReverseStmt + newInstancesReverseStmt
		backwardCommit += "\n\tstage.Commit()"
		// append to the end of the backward commits slice
		stage.backwardCommits = append(stage.backwardCommits, backwardCommit)
		stage.modified = true
	} else {
		stage.modified = false
	}
}

// ComputeReferenceAndOrders will creates a deep copy of each of the staged elements
func (stage *Stage) ComputeReferenceAndOrders() {
	// insertion point per named struct
	stage.A_directives_reference = make(map[*A_directive]*A_directive)
	stage.A_directives_referenceOrder = make(map[*A_directive]uint) // diff Unstage needs the reference order
	stage.A_directives_instance = make(map[*A_directive]*A_directive)
	for instance := range stage.A_directives {
		_copy := instance.GongCopy().(*A_directive)
		stage.A_directives_reference[instance] = _copy
		stage.A_directives_instance[_copy] = instance
		stage.A_directives_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_measures_reference = make(map[*A_measure]*A_measure)
	stage.A_measures_referenceOrder = make(map[*A_measure]uint) // diff Unstage needs the reference order
	stage.A_measures_instance = make(map[*A_measure]*A_measure)
	for instance := range stage.A_measures {
		_copy := instance.GongCopy().(*A_measure)
		stage.A_measures_reference[instance] = _copy
		stage.A_measures_instance[_copy] = instance
		stage.A_measures_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_measure_1s_reference = make(map[*A_measure_1]*A_measure_1)
	stage.A_measure_1s_referenceOrder = make(map[*A_measure_1]uint) // diff Unstage needs the reference order
	stage.A_measure_1s_instance = make(map[*A_measure_1]*A_measure_1)
	for instance := range stage.A_measure_1s {
		_copy := instance.GongCopy().(*A_measure_1)
		stage.A_measure_1s_reference[instance] = _copy
		stage.A_measure_1s_instance[_copy] = instance
		stage.A_measure_1s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_parts_reference = make(map[*A_part]*A_part)
	stage.A_parts_referenceOrder = make(map[*A_part]uint) // diff Unstage needs the reference order
	stage.A_parts_instance = make(map[*A_part]*A_part)
	for instance := range stage.A_parts {
		_copy := instance.GongCopy().(*A_part)
		stage.A_parts_reference[instance] = _copy
		stage.A_parts_instance[_copy] = instance
		stage.A_parts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.A_part_1s_reference = make(map[*A_part_1]*A_part_1)
	stage.A_part_1s_referenceOrder = make(map[*A_part_1]uint) // diff Unstage needs the reference order
	stage.A_part_1s_instance = make(map[*A_part_1]*A_part_1)
	for instance := range stage.A_part_1s {
		_copy := instance.GongCopy().(*A_part_1)
		stage.A_part_1s_reference[instance] = _copy
		stage.A_part_1s_instance[_copy] = instance
		stage.A_part_1s_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Accidentals_reference = make(map[*Accidental]*Accidental)
	stage.Accidentals_referenceOrder = make(map[*Accidental]uint) // diff Unstage needs the reference order
	stage.Accidentals_instance = make(map[*Accidental]*Accidental)
	for instance := range stage.Accidentals {
		_copy := instance.GongCopy().(*Accidental)
		stage.Accidentals_reference[instance] = _copy
		stage.Accidentals_instance[_copy] = instance
		stage.Accidentals_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Accidental_marks_reference = make(map[*Accidental_mark]*Accidental_mark)
	stage.Accidental_marks_referenceOrder = make(map[*Accidental_mark]uint) // diff Unstage needs the reference order
	stage.Accidental_marks_instance = make(map[*Accidental_mark]*Accidental_mark)
	for instance := range stage.Accidental_marks {
		_copy := instance.GongCopy().(*Accidental_mark)
		stage.Accidental_marks_reference[instance] = _copy
		stage.Accidental_marks_instance[_copy] = instance
		stage.Accidental_marks_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Accidental_texts_reference = make(map[*Accidental_text]*Accidental_text)
	stage.Accidental_texts_referenceOrder = make(map[*Accidental_text]uint) // diff Unstage needs the reference order
	stage.Accidental_texts_instance = make(map[*Accidental_text]*Accidental_text)
	for instance := range stage.Accidental_texts {
		_copy := instance.GongCopy().(*Accidental_text)
		stage.Accidental_texts_reference[instance] = _copy
		stage.Accidental_texts_instance[_copy] = instance
		stage.Accidental_texts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Accords_reference = make(map[*Accord]*Accord)
	stage.Accords_referenceOrder = make(map[*Accord]uint) // diff Unstage needs the reference order
	stage.Accords_instance = make(map[*Accord]*Accord)
	for instance := range stage.Accords {
		_copy := instance.GongCopy().(*Accord)
		stage.Accords_reference[instance] = _copy
		stage.Accords_instance[_copy] = instance
		stage.Accords_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Accordion_registrations_reference = make(map[*Accordion_registration]*Accordion_registration)
	stage.Accordion_registrations_referenceOrder = make(map[*Accordion_registration]uint) // diff Unstage needs the reference order
	stage.Accordion_registrations_instance = make(map[*Accordion_registration]*Accordion_registration)
	for instance := range stage.Accordion_registrations {
		_copy := instance.GongCopy().(*Accordion_registration)
		stage.Accordion_registrations_reference[instance] = _copy
		stage.Accordion_registrations_instance[_copy] = instance
		stage.Accordion_registrations_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Appearances_reference = make(map[*Appearance]*Appearance)
	stage.Appearances_referenceOrder = make(map[*Appearance]uint) // diff Unstage needs the reference order
	stage.Appearances_instance = make(map[*Appearance]*Appearance)
	for instance := range stage.Appearances {
		_copy := instance.GongCopy().(*Appearance)
		stage.Appearances_reference[instance] = _copy
		stage.Appearances_instance[_copy] = instance
		stage.Appearances_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Arpeggiates_reference = make(map[*Arpeggiate]*Arpeggiate)
	stage.Arpeggiates_referenceOrder = make(map[*Arpeggiate]uint) // diff Unstage needs the reference order
	stage.Arpeggiates_instance = make(map[*Arpeggiate]*Arpeggiate)
	for instance := range stage.Arpeggiates {
		_copy := instance.GongCopy().(*Arpeggiate)
		stage.Arpeggiates_reference[instance] = _copy
		stage.Arpeggiates_instance[_copy] = instance
		stage.Arpeggiates_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Arrows_reference = make(map[*Arrow]*Arrow)
	stage.Arrows_referenceOrder = make(map[*Arrow]uint) // diff Unstage needs the reference order
	stage.Arrows_instance = make(map[*Arrow]*Arrow)
	for instance := range stage.Arrows {
		_copy := instance.GongCopy().(*Arrow)
		stage.Arrows_reference[instance] = _copy
		stage.Arrows_instance[_copy] = instance
		stage.Arrows_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Articulationss_reference = make(map[*Articulations]*Articulations)
	stage.Articulationss_referenceOrder = make(map[*Articulations]uint) // diff Unstage needs the reference order
	stage.Articulationss_instance = make(map[*Articulations]*Articulations)
	for instance := range stage.Articulationss {
		_copy := instance.GongCopy().(*Articulations)
		stage.Articulationss_reference[instance] = _copy
		stage.Articulationss_instance[_copy] = instance
		stage.Articulationss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Assesss_reference = make(map[*Assess]*Assess)
	stage.Assesss_referenceOrder = make(map[*Assess]uint) // diff Unstage needs the reference order
	stage.Assesss_instance = make(map[*Assess]*Assess)
	for instance := range stage.Assesss {
		_copy := instance.GongCopy().(*Assess)
		stage.Assesss_reference[instance] = _copy
		stage.Assesss_instance[_copy] = instance
		stage.Assesss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Attributess_reference = make(map[*Attributes]*Attributes)
	stage.Attributess_referenceOrder = make(map[*Attributes]uint) // diff Unstage needs the reference order
	stage.Attributess_instance = make(map[*Attributes]*Attributes)
	for instance := range stage.Attributess {
		_copy := instance.GongCopy().(*Attributes)
		stage.Attributess_reference[instance] = _copy
		stage.Attributess_instance[_copy] = instance
		stage.Attributess_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Backups_reference = make(map[*Backup]*Backup)
	stage.Backups_referenceOrder = make(map[*Backup]uint) // diff Unstage needs the reference order
	stage.Backups_instance = make(map[*Backup]*Backup)
	for instance := range stage.Backups {
		_copy := instance.GongCopy().(*Backup)
		stage.Backups_reference[instance] = _copy
		stage.Backups_instance[_copy] = instance
		stage.Backups_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Bar_style_colors_reference = make(map[*Bar_style_color]*Bar_style_color)
	stage.Bar_style_colors_referenceOrder = make(map[*Bar_style_color]uint) // diff Unstage needs the reference order
	stage.Bar_style_colors_instance = make(map[*Bar_style_color]*Bar_style_color)
	for instance := range stage.Bar_style_colors {
		_copy := instance.GongCopy().(*Bar_style_color)
		stage.Bar_style_colors_reference[instance] = _copy
		stage.Bar_style_colors_instance[_copy] = instance
		stage.Bar_style_colors_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Barlines_reference = make(map[*Barline]*Barline)
	stage.Barlines_referenceOrder = make(map[*Barline]uint) // diff Unstage needs the reference order
	stage.Barlines_instance = make(map[*Barline]*Barline)
	for instance := range stage.Barlines {
		_copy := instance.GongCopy().(*Barline)
		stage.Barlines_reference[instance] = _copy
		stage.Barlines_instance[_copy] = instance
		stage.Barlines_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Barres_reference = make(map[*Barre]*Barre)
	stage.Barres_referenceOrder = make(map[*Barre]uint) // diff Unstage needs the reference order
	stage.Barres_instance = make(map[*Barre]*Barre)
	for instance := range stage.Barres {
		_copy := instance.GongCopy().(*Barre)
		stage.Barres_reference[instance] = _copy
		stage.Barres_instance[_copy] = instance
		stage.Barres_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Basss_reference = make(map[*Bass]*Bass)
	stage.Basss_referenceOrder = make(map[*Bass]uint) // diff Unstage needs the reference order
	stage.Basss_instance = make(map[*Bass]*Bass)
	for instance := range stage.Basss {
		_copy := instance.GongCopy().(*Bass)
		stage.Basss_reference[instance] = _copy
		stage.Basss_instance[_copy] = instance
		stage.Basss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Bass_steps_reference = make(map[*Bass_step]*Bass_step)
	stage.Bass_steps_referenceOrder = make(map[*Bass_step]uint) // diff Unstage needs the reference order
	stage.Bass_steps_instance = make(map[*Bass_step]*Bass_step)
	for instance := range stage.Bass_steps {
		_copy := instance.GongCopy().(*Bass_step)
		stage.Bass_steps_reference[instance] = _copy
		stage.Bass_steps_instance[_copy] = instance
		stage.Bass_steps_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Beams_reference = make(map[*Beam]*Beam)
	stage.Beams_referenceOrder = make(map[*Beam]uint) // diff Unstage needs the reference order
	stage.Beams_instance = make(map[*Beam]*Beam)
	for instance := range stage.Beams {
		_copy := instance.GongCopy().(*Beam)
		stage.Beams_reference[instance] = _copy
		stage.Beams_instance[_copy] = instance
		stage.Beams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Beat_repeats_reference = make(map[*Beat_repeat]*Beat_repeat)
	stage.Beat_repeats_referenceOrder = make(map[*Beat_repeat]uint) // diff Unstage needs the reference order
	stage.Beat_repeats_instance = make(map[*Beat_repeat]*Beat_repeat)
	for instance := range stage.Beat_repeats {
		_copy := instance.GongCopy().(*Beat_repeat)
		stage.Beat_repeats_reference[instance] = _copy
		stage.Beat_repeats_instance[_copy] = instance
		stage.Beat_repeats_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Beat_unit_tieds_reference = make(map[*Beat_unit_tied]*Beat_unit_tied)
	stage.Beat_unit_tieds_referenceOrder = make(map[*Beat_unit_tied]uint) // diff Unstage needs the reference order
	stage.Beat_unit_tieds_instance = make(map[*Beat_unit_tied]*Beat_unit_tied)
	for instance := range stage.Beat_unit_tieds {
		_copy := instance.GongCopy().(*Beat_unit_tied)
		stage.Beat_unit_tieds_reference[instance] = _copy
		stage.Beat_unit_tieds_instance[_copy] = instance
		stage.Beat_unit_tieds_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Beaters_reference = make(map[*Beater]*Beater)
	stage.Beaters_referenceOrder = make(map[*Beater]uint) // diff Unstage needs the reference order
	stage.Beaters_instance = make(map[*Beater]*Beater)
	for instance := range stage.Beaters {
		_copy := instance.GongCopy().(*Beater)
		stage.Beaters_reference[instance] = _copy
		stage.Beaters_instance[_copy] = instance
		stage.Beaters_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Bends_reference = make(map[*Bend]*Bend)
	stage.Bends_referenceOrder = make(map[*Bend]uint) // diff Unstage needs the reference order
	stage.Bends_instance = make(map[*Bend]*Bend)
	for instance := range stage.Bends {
		_copy := instance.GongCopy().(*Bend)
		stage.Bends_reference[instance] = _copy
		stage.Bends_instance[_copy] = instance
		stage.Bends_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Bookmarks_reference = make(map[*Bookmark]*Bookmark)
	stage.Bookmarks_referenceOrder = make(map[*Bookmark]uint) // diff Unstage needs the reference order
	stage.Bookmarks_instance = make(map[*Bookmark]*Bookmark)
	for instance := range stage.Bookmarks {
		_copy := instance.GongCopy().(*Bookmark)
		stage.Bookmarks_reference[instance] = _copy
		stage.Bookmarks_instance[_copy] = instance
		stage.Bookmarks_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Brackets_reference = make(map[*Bracket]*Bracket)
	stage.Brackets_referenceOrder = make(map[*Bracket]uint) // diff Unstage needs the reference order
	stage.Brackets_instance = make(map[*Bracket]*Bracket)
	for instance := range stage.Brackets {
		_copy := instance.GongCopy().(*Bracket)
		stage.Brackets_reference[instance] = _copy
		stage.Brackets_instance[_copy] = instance
		stage.Brackets_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Breath_marks_reference = make(map[*Breath_mark]*Breath_mark)
	stage.Breath_marks_referenceOrder = make(map[*Breath_mark]uint) // diff Unstage needs the reference order
	stage.Breath_marks_instance = make(map[*Breath_mark]*Breath_mark)
	for instance := range stage.Breath_marks {
		_copy := instance.GongCopy().(*Breath_mark)
		stage.Breath_marks_reference[instance] = _copy
		stage.Breath_marks_instance[_copy] = instance
		stage.Breath_marks_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Caesuras_reference = make(map[*Caesura]*Caesura)
	stage.Caesuras_referenceOrder = make(map[*Caesura]uint) // diff Unstage needs the reference order
	stage.Caesuras_instance = make(map[*Caesura]*Caesura)
	for instance := range stage.Caesuras {
		_copy := instance.GongCopy().(*Caesura)
		stage.Caesuras_reference[instance] = _copy
		stage.Caesuras_instance[_copy] = instance
		stage.Caesuras_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Cancels_reference = make(map[*Cancel]*Cancel)
	stage.Cancels_referenceOrder = make(map[*Cancel]uint) // diff Unstage needs the reference order
	stage.Cancels_instance = make(map[*Cancel]*Cancel)
	for instance := range stage.Cancels {
		_copy := instance.GongCopy().(*Cancel)
		stage.Cancels_reference[instance] = _copy
		stage.Cancels_instance[_copy] = instance
		stage.Cancels_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Clefs_reference = make(map[*Clef]*Clef)
	stage.Clefs_referenceOrder = make(map[*Clef]uint) // diff Unstage needs the reference order
	stage.Clefs_instance = make(map[*Clef]*Clef)
	for instance := range stage.Clefs {
		_copy := instance.GongCopy().(*Clef)
		stage.Clefs_reference[instance] = _copy
		stage.Clefs_instance[_copy] = instance
		stage.Clefs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Codas_reference = make(map[*Coda]*Coda)
	stage.Codas_referenceOrder = make(map[*Coda]uint) // diff Unstage needs the reference order
	stage.Codas_instance = make(map[*Coda]*Coda)
	for instance := range stage.Codas {
		_copy := instance.GongCopy().(*Coda)
		stage.Codas_reference[instance] = _copy
		stage.Codas_instance[_copy] = instance
		stage.Codas_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Credits_reference = make(map[*Credit]*Credit)
	stage.Credits_referenceOrder = make(map[*Credit]uint) // diff Unstage needs the reference order
	stage.Credits_instance = make(map[*Credit]*Credit)
	for instance := range stage.Credits {
		_copy := instance.GongCopy().(*Credit)
		stage.Credits_reference[instance] = _copy
		stage.Credits_instance[_copy] = instance
		stage.Credits_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Dashess_reference = make(map[*Dashes]*Dashes)
	stage.Dashess_referenceOrder = make(map[*Dashes]uint) // diff Unstage needs the reference order
	stage.Dashess_instance = make(map[*Dashes]*Dashes)
	for instance := range stage.Dashess {
		_copy := instance.GongCopy().(*Dashes)
		stage.Dashess_reference[instance] = _copy
		stage.Dashess_instance[_copy] = instance
		stage.Dashess_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Defaultss_reference = make(map[*Defaults]*Defaults)
	stage.Defaultss_referenceOrder = make(map[*Defaults]uint) // diff Unstage needs the reference order
	stage.Defaultss_instance = make(map[*Defaults]*Defaults)
	for instance := range stage.Defaultss {
		_copy := instance.GongCopy().(*Defaults)
		stage.Defaultss_reference[instance] = _copy
		stage.Defaultss_instance[_copy] = instance
		stage.Defaultss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Degrees_reference = make(map[*Degree]*Degree)
	stage.Degrees_referenceOrder = make(map[*Degree]uint) // diff Unstage needs the reference order
	stage.Degrees_instance = make(map[*Degree]*Degree)
	for instance := range stage.Degrees {
		_copy := instance.GongCopy().(*Degree)
		stage.Degrees_reference[instance] = _copy
		stage.Degrees_instance[_copy] = instance
		stage.Degrees_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Degree_alters_reference = make(map[*Degree_alter]*Degree_alter)
	stage.Degree_alters_referenceOrder = make(map[*Degree_alter]uint) // diff Unstage needs the reference order
	stage.Degree_alters_instance = make(map[*Degree_alter]*Degree_alter)
	for instance := range stage.Degree_alters {
		_copy := instance.GongCopy().(*Degree_alter)
		stage.Degree_alters_reference[instance] = _copy
		stage.Degree_alters_instance[_copy] = instance
		stage.Degree_alters_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Degree_types_reference = make(map[*Degree_type]*Degree_type)
	stage.Degree_types_referenceOrder = make(map[*Degree_type]uint) // diff Unstage needs the reference order
	stage.Degree_types_instance = make(map[*Degree_type]*Degree_type)
	for instance := range stage.Degree_types {
		_copy := instance.GongCopy().(*Degree_type)
		stage.Degree_types_reference[instance] = _copy
		stage.Degree_types_instance[_copy] = instance
		stage.Degree_types_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Degree_values_reference = make(map[*Degree_value]*Degree_value)
	stage.Degree_values_referenceOrder = make(map[*Degree_value]uint) // diff Unstage needs the reference order
	stage.Degree_values_instance = make(map[*Degree_value]*Degree_value)
	for instance := range stage.Degree_values {
		_copy := instance.GongCopy().(*Degree_value)
		stage.Degree_values_reference[instance] = _copy
		stage.Degree_values_instance[_copy] = instance
		stage.Degree_values_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Directions_reference = make(map[*Direction]*Direction)
	stage.Directions_referenceOrder = make(map[*Direction]uint) // diff Unstage needs the reference order
	stage.Directions_instance = make(map[*Direction]*Direction)
	for instance := range stage.Directions {
		_copy := instance.GongCopy().(*Direction)
		stage.Directions_reference[instance] = _copy
		stage.Directions_instance[_copy] = instance
		stage.Directions_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Direction_types_reference = make(map[*Direction_type]*Direction_type)
	stage.Direction_types_referenceOrder = make(map[*Direction_type]uint) // diff Unstage needs the reference order
	stage.Direction_types_instance = make(map[*Direction_type]*Direction_type)
	for instance := range stage.Direction_types {
		_copy := instance.GongCopy().(*Direction_type)
		stage.Direction_types_reference[instance] = _copy
		stage.Direction_types_instance[_copy] = instance
		stage.Direction_types_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Distances_reference = make(map[*Distance]*Distance)
	stage.Distances_referenceOrder = make(map[*Distance]uint) // diff Unstage needs the reference order
	stage.Distances_instance = make(map[*Distance]*Distance)
	for instance := range stage.Distances {
		_copy := instance.GongCopy().(*Distance)
		stage.Distances_reference[instance] = _copy
		stage.Distances_instance[_copy] = instance
		stage.Distances_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Doubles_reference = make(map[*Double]*Double)
	stage.Doubles_referenceOrder = make(map[*Double]uint) // diff Unstage needs the reference order
	stage.Doubles_instance = make(map[*Double]*Double)
	for instance := range stage.Doubles {
		_copy := instance.GongCopy().(*Double)
		stage.Doubles_reference[instance] = _copy
		stage.Doubles_instance[_copy] = instance
		stage.Doubles_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Dynamicss_reference = make(map[*Dynamics]*Dynamics)
	stage.Dynamicss_referenceOrder = make(map[*Dynamics]uint) // diff Unstage needs the reference order
	stage.Dynamicss_instance = make(map[*Dynamics]*Dynamics)
	for instance := range stage.Dynamicss {
		_copy := instance.GongCopy().(*Dynamics)
		stage.Dynamicss_reference[instance] = _copy
		stage.Dynamicss_instance[_copy] = instance
		stage.Dynamicss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Effects_reference = make(map[*Effect]*Effect)
	stage.Effects_referenceOrder = make(map[*Effect]uint) // diff Unstage needs the reference order
	stage.Effects_instance = make(map[*Effect]*Effect)
	for instance := range stage.Effects {
		_copy := instance.GongCopy().(*Effect)
		stage.Effects_reference[instance] = _copy
		stage.Effects_instance[_copy] = instance
		stage.Effects_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Elisions_reference = make(map[*Elision]*Elision)
	stage.Elisions_referenceOrder = make(map[*Elision]uint) // diff Unstage needs the reference order
	stage.Elisions_instance = make(map[*Elision]*Elision)
	for instance := range stage.Elisions {
		_copy := instance.GongCopy().(*Elision)
		stage.Elisions_reference[instance] = _copy
		stage.Elisions_instance[_copy] = instance
		stage.Elisions_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Emptys_reference = make(map[*Empty]*Empty)
	stage.Emptys_referenceOrder = make(map[*Empty]uint) // diff Unstage needs the reference order
	stage.Emptys_instance = make(map[*Empty]*Empty)
	for instance := range stage.Emptys {
		_copy := instance.GongCopy().(*Empty)
		stage.Emptys_reference[instance] = _copy
		stage.Emptys_instance[_copy] = instance
		stage.Emptys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Empty_fonts_reference = make(map[*Empty_font]*Empty_font)
	stage.Empty_fonts_referenceOrder = make(map[*Empty_font]uint) // diff Unstage needs the reference order
	stage.Empty_fonts_instance = make(map[*Empty_font]*Empty_font)
	for instance := range stage.Empty_fonts {
		_copy := instance.GongCopy().(*Empty_font)
		stage.Empty_fonts_reference[instance] = _copy
		stage.Empty_fonts_instance[_copy] = instance
		stage.Empty_fonts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Empty_lines_reference = make(map[*Empty_line]*Empty_line)
	stage.Empty_lines_referenceOrder = make(map[*Empty_line]uint) // diff Unstage needs the reference order
	stage.Empty_lines_instance = make(map[*Empty_line]*Empty_line)
	for instance := range stage.Empty_lines {
		_copy := instance.GongCopy().(*Empty_line)
		stage.Empty_lines_reference[instance] = _copy
		stage.Empty_lines_instance[_copy] = instance
		stage.Empty_lines_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Empty_placements_reference = make(map[*Empty_placement]*Empty_placement)
	stage.Empty_placements_referenceOrder = make(map[*Empty_placement]uint) // diff Unstage needs the reference order
	stage.Empty_placements_instance = make(map[*Empty_placement]*Empty_placement)
	for instance := range stage.Empty_placements {
		_copy := instance.GongCopy().(*Empty_placement)
		stage.Empty_placements_reference[instance] = _copy
		stage.Empty_placements_instance[_copy] = instance
		stage.Empty_placements_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Empty_placement_smufls_reference = make(map[*Empty_placement_smufl]*Empty_placement_smufl)
	stage.Empty_placement_smufls_referenceOrder = make(map[*Empty_placement_smufl]uint) // diff Unstage needs the reference order
	stage.Empty_placement_smufls_instance = make(map[*Empty_placement_smufl]*Empty_placement_smufl)
	for instance := range stage.Empty_placement_smufls {
		_copy := instance.GongCopy().(*Empty_placement_smufl)
		stage.Empty_placement_smufls_reference[instance] = _copy
		stage.Empty_placement_smufls_instance[_copy] = instance
		stage.Empty_placement_smufls_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Empty_print_object_style_aligns_reference = make(map[*Empty_print_object_style_align]*Empty_print_object_style_align)
	stage.Empty_print_object_style_aligns_referenceOrder = make(map[*Empty_print_object_style_align]uint) // diff Unstage needs the reference order
	stage.Empty_print_object_style_aligns_instance = make(map[*Empty_print_object_style_align]*Empty_print_object_style_align)
	for instance := range stage.Empty_print_object_style_aligns {
		_copy := instance.GongCopy().(*Empty_print_object_style_align)
		stage.Empty_print_object_style_aligns_reference[instance] = _copy
		stage.Empty_print_object_style_aligns_instance[_copy] = instance
		stage.Empty_print_object_style_aligns_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Empty_print_styles_reference = make(map[*Empty_print_style]*Empty_print_style)
	stage.Empty_print_styles_referenceOrder = make(map[*Empty_print_style]uint) // diff Unstage needs the reference order
	stage.Empty_print_styles_instance = make(map[*Empty_print_style]*Empty_print_style)
	for instance := range stage.Empty_print_styles {
		_copy := instance.GongCopy().(*Empty_print_style)
		stage.Empty_print_styles_reference[instance] = _copy
		stage.Empty_print_styles_instance[_copy] = instance
		stage.Empty_print_styles_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Empty_print_style_aligns_reference = make(map[*Empty_print_style_align]*Empty_print_style_align)
	stage.Empty_print_style_aligns_referenceOrder = make(map[*Empty_print_style_align]uint) // diff Unstage needs the reference order
	stage.Empty_print_style_aligns_instance = make(map[*Empty_print_style_align]*Empty_print_style_align)
	for instance := range stage.Empty_print_style_aligns {
		_copy := instance.GongCopy().(*Empty_print_style_align)
		stage.Empty_print_style_aligns_reference[instance] = _copy
		stage.Empty_print_style_aligns_instance[_copy] = instance
		stage.Empty_print_style_aligns_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Empty_print_style_align_ids_reference = make(map[*Empty_print_style_align_id]*Empty_print_style_align_id)
	stage.Empty_print_style_align_ids_referenceOrder = make(map[*Empty_print_style_align_id]uint) // diff Unstage needs the reference order
	stage.Empty_print_style_align_ids_instance = make(map[*Empty_print_style_align_id]*Empty_print_style_align_id)
	for instance := range stage.Empty_print_style_align_ids {
		_copy := instance.GongCopy().(*Empty_print_style_align_id)
		stage.Empty_print_style_align_ids_reference[instance] = _copy
		stage.Empty_print_style_align_ids_instance[_copy] = instance
		stage.Empty_print_style_align_ids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Empty_trill_sounds_reference = make(map[*Empty_trill_sound]*Empty_trill_sound)
	stage.Empty_trill_sounds_referenceOrder = make(map[*Empty_trill_sound]uint) // diff Unstage needs the reference order
	stage.Empty_trill_sounds_instance = make(map[*Empty_trill_sound]*Empty_trill_sound)
	for instance := range stage.Empty_trill_sounds {
		_copy := instance.GongCopy().(*Empty_trill_sound)
		stage.Empty_trill_sounds_reference[instance] = _copy
		stage.Empty_trill_sounds_instance[_copy] = instance
		stage.Empty_trill_sounds_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Encodings_reference = make(map[*Encoding]*Encoding)
	stage.Encodings_referenceOrder = make(map[*Encoding]uint) // diff Unstage needs the reference order
	stage.Encodings_instance = make(map[*Encoding]*Encoding)
	for instance := range stage.Encodings {
		_copy := instance.GongCopy().(*Encoding)
		stage.Encodings_reference[instance] = _copy
		stage.Encodings_instance[_copy] = instance
		stage.Encodings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Endings_reference = make(map[*Ending]*Ending)
	stage.Endings_referenceOrder = make(map[*Ending]uint) // diff Unstage needs the reference order
	stage.Endings_instance = make(map[*Ending]*Ending)
	for instance := range stage.Endings {
		_copy := instance.GongCopy().(*Ending)
		stage.Endings_reference[instance] = _copy
		stage.Endings_instance[_copy] = instance
		stage.Endings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Extends_reference = make(map[*Extend]*Extend)
	stage.Extends_referenceOrder = make(map[*Extend]uint) // diff Unstage needs the reference order
	stage.Extends_instance = make(map[*Extend]*Extend)
	for instance := range stage.Extends {
		_copy := instance.GongCopy().(*Extend)
		stage.Extends_reference[instance] = _copy
		stage.Extends_instance[_copy] = instance
		stage.Extends_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Features_reference = make(map[*Feature]*Feature)
	stage.Features_referenceOrder = make(map[*Feature]uint) // diff Unstage needs the reference order
	stage.Features_instance = make(map[*Feature]*Feature)
	for instance := range stage.Features {
		_copy := instance.GongCopy().(*Feature)
		stage.Features_reference[instance] = _copy
		stage.Features_instance[_copy] = instance
		stage.Features_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Fermatas_reference = make(map[*Fermata]*Fermata)
	stage.Fermatas_referenceOrder = make(map[*Fermata]uint) // diff Unstage needs the reference order
	stage.Fermatas_instance = make(map[*Fermata]*Fermata)
	for instance := range stage.Fermatas {
		_copy := instance.GongCopy().(*Fermata)
		stage.Fermatas_reference[instance] = _copy
		stage.Fermatas_instance[_copy] = instance
		stage.Fermatas_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Figures_reference = make(map[*Figure]*Figure)
	stage.Figures_referenceOrder = make(map[*Figure]uint) // diff Unstage needs the reference order
	stage.Figures_instance = make(map[*Figure]*Figure)
	for instance := range stage.Figures {
		_copy := instance.GongCopy().(*Figure)
		stage.Figures_reference[instance] = _copy
		stage.Figures_instance[_copy] = instance
		stage.Figures_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Figured_basss_reference = make(map[*Figured_bass]*Figured_bass)
	stage.Figured_basss_referenceOrder = make(map[*Figured_bass]uint) // diff Unstage needs the reference order
	stage.Figured_basss_instance = make(map[*Figured_bass]*Figured_bass)
	for instance := range stage.Figured_basss {
		_copy := instance.GongCopy().(*Figured_bass)
		stage.Figured_basss_reference[instance] = _copy
		stage.Figured_basss_instance[_copy] = instance
		stage.Figured_basss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Fingerings_reference = make(map[*Fingering]*Fingering)
	stage.Fingerings_referenceOrder = make(map[*Fingering]uint) // diff Unstage needs the reference order
	stage.Fingerings_instance = make(map[*Fingering]*Fingering)
	for instance := range stage.Fingerings {
		_copy := instance.GongCopy().(*Fingering)
		stage.Fingerings_reference[instance] = _copy
		stage.Fingerings_instance[_copy] = instance
		stage.Fingerings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.First_frets_reference = make(map[*First_fret]*First_fret)
	stage.First_frets_referenceOrder = make(map[*First_fret]uint) // diff Unstage needs the reference order
	stage.First_frets_instance = make(map[*First_fret]*First_fret)
	for instance := range stage.First_frets {
		_copy := instance.GongCopy().(*First_fret)
		stage.First_frets_reference[instance] = _copy
		stage.First_frets_instance[_copy] = instance
		stage.First_frets_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.For_parts_reference = make(map[*For_part]*For_part)
	stage.For_parts_referenceOrder = make(map[*For_part]uint) // diff Unstage needs the reference order
	stage.For_parts_instance = make(map[*For_part]*For_part)
	for instance := range stage.For_parts {
		_copy := instance.GongCopy().(*For_part)
		stage.For_parts_reference[instance] = _copy
		stage.For_parts_instance[_copy] = instance
		stage.For_parts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Formatted_symbols_reference = make(map[*Formatted_symbol]*Formatted_symbol)
	stage.Formatted_symbols_referenceOrder = make(map[*Formatted_symbol]uint) // diff Unstage needs the reference order
	stage.Formatted_symbols_instance = make(map[*Formatted_symbol]*Formatted_symbol)
	for instance := range stage.Formatted_symbols {
		_copy := instance.GongCopy().(*Formatted_symbol)
		stage.Formatted_symbols_reference[instance] = _copy
		stage.Formatted_symbols_instance[_copy] = instance
		stage.Formatted_symbols_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Formatted_symbol_ids_reference = make(map[*Formatted_symbol_id]*Formatted_symbol_id)
	stage.Formatted_symbol_ids_referenceOrder = make(map[*Formatted_symbol_id]uint) // diff Unstage needs the reference order
	stage.Formatted_symbol_ids_instance = make(map[*Formatted_symbol_id]*Formatted_symbol_id)
	for instance := range stage.Formatted_symbol_ids {
		_copy := instance.GongCopy().(*Formatted_symbol_id)
		stage.Formatted_symbol_ids_reference[instance] = _copy
		stage.Formatted_symbol_ids_instance[_copy] = instance
		stage.Formatted_symbol_ids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Formatted_texts_reference = make(map[*Formatted_text]*Formatted_text)
	stage.Formatted_texts_referenceOrder = make(map[*Formatted_text]uint) // diff Unstage needs the reference order
	stage.Formatted_texts_instance = make(map[*Formatted_text]*Formatted_text)
	for instance := range stage.Formatted_texts {
		_copy := instance.GongCopy().(*Formatted_text)
		stage.Formatted_texts_reference[instance] = _copy
		stage.Formatted_texts_instance[_copy] = instance
		stage.Formatted_texts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Formatted_text_ids_reference = make(map[*Formatted_text_id]*Formatted_text_id)
	stage.Formatted_text_ids_referenceOrder = make(map[*Formatted_text_id]uint) // diff Unstage needs the reference order
	stage.Formatted_text_ids_instance = make(map[*Formatted_text_id]*Formatted_text_id)
	for instance := range stage.Formatted_text_ids {
		_copy := instance.GongCopy().(*Formatted_text_id)
		stage.Formatted_text_ids_reference[instance] = _copy
		stage.Formatted_text_ids_instance[_copy] = instance
		stage.Formatted_text_ids_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Forwards_reference = make(map[*Forward]*Forward)
	stage.Forwards_referenceOrder = make(map[*Forward]uint) // diff Unstage needs the reference order
	stage.Forwards_instance = make(map[*Forward]*Forward)
	for instance := range stage.Forwards {
		_copy := instance.GongCopy().(*Forward)
		stage.Forwards_reference[instance] = _copy
		stage.Forwards_instance[_copy] = instance
		stage.Forwards_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Frames_reference = make(map[*Frame]*Frame)
	stage.Frames_referenceOrder = make(map[*Frame]uint) // diff Unstage needs the reference order
	stage.Frames_instance = make(map[*Frame]*Frame)
	for instance := range stage.Frames {
		_copy := instance.GongCopy().(*Frame)
		stage.Frames_reference[instance] = _copy
		stage.Frames_instance[_copy] = instance
		stage.Frames_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Frame_notes_reference = make(map[*Frame_note]*Frame_note)
	stage.Frame_notes_referenceOrder = make(map[*Frame_note]uint) // diff Unstage needs the reference order
	stage.Frame_notes_instance = make(map[*Frame_note]*Frame_note)
	for instance := range stage.Frame_notes {
		_copy := instance.GongCopy().(*Frame_note)
		stage.Frame_notes_reference[instance] = _copy
		stage.Frame_notes_instance[_copy] = instance
		stage.Frame_notes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Frets_reference = make(map[*Fret]*Fret)
	stage.Frets_referenceOrder = make(map[*Fret]uint) // diff Unstage needs the reference order
	stage.Frets_instance = make(map[*Fret]*Fret)
	for instance := range stage.Frets {
		_copy := instance.GongCopy().(*Fret)
		stage.Frets_reference[instance] = _copy
		stage.Frets_instance[_copy] = instance
		stage.Frets_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Glasss_reference = make(map[*Glass]*Glass)
	stage.Glasss_referenceOrder = make(map[*Glass]uint) // diff Unstage needs the reference order
	stage.Glasss_instance = make(map[*Glass]*Glass)
	for instance := range stage.Glasss {
		_copy := instance.GongCopy().(*Glass)
		stage.Glasss_reference[instance] = _copy
		stage.Glasss_instance[_copy] = instance
		stage.Glasss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Glissandos_reference = make(map[*Glissando]*Glissando)
	stage.Glissandos_referenceOrder = make(map[*Glissando]uint) // diff Unstage needs the reference order
	stage.Glissandos_instance = make(map[*Glissando]*Glissando)
	for instance := range stage.Glissandos {
		_copy := instance.GongCopy().(*Glissando)
		stage.Glissandos_reference[instance] = _copy
		stage.Glissandos_instance[_copy] = instance
		stage.Glissandos_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Glyphs_reference = make(map[*Glyph]*Glyph)
	stage.Glyphs_referenceOrder = make(map[*Glyph]uint) // diff Unstage needs the reference order
	stage.Glyphs_instance = make(map[*Glyph]*Glyph)
	for instance := range stage.Glyphs {
		_copy := instance.GongCopy().(*Glyph)
		stage.Glyphs_reference[instance] = _copy
		stage.Glyphs_instance[_copy] = instance
		stage.Glyphs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Graces_reference = make(map[*Grace]*Grace)
	stage.Graces_referenceOrder = make(map[*Grace]uint) // diff Unstage needs the reference order
	stage.Graces_instance = make(map[*Grace]*Grace)
	for instance := range stage.Graces {
		_copy := instance.GongCopy().(*Grace)
		stage.Graces_reference[instance] = _copy
		stage.Graces_instance[_copy] = instance
		stage.Graces_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Group_barlines_reference = make(map[*Group_barline]*Group_barline)
	stage.Group_barlines_referenceOrder = make(map[*Group_barline]uint) // diff Unstage needs the reference order
	stage.Group_barlines_instance = make(map[*Group_barline]*Group_barline)
	for instance := range stage.Group_barlines {
		_copy := instance.GongCopy().(*Group_barline)
		stage.Group_barlines_reference[instance] = _copy
		stage.Group_barlines_instance[_copy] = instance
		stage.Group_barlines_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Group_names_reference = make(map[*Group_name]*Group_name)
	stage.Group_names_referenceOrder = make(map[*Group_name]uint) // diff Unstage needs the reference order
	stage.Group_names_instance = make(map[*Group_name]*Group_name)
	for instance := range stage.Group_names {
		_copy := instance.GongCopy().(*Group_name)
		stage.Group_names_reference[instance] = _copy
		stage.Group_names_instance[_copy] = instance
		stage.Group_names_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Group_symbols_reference = make(map[*Group_symbol]*Group_symbol)
	stage.Group_symbols_referenceOrder = make(map[*Group_symbol]uint) // diff Unstage needs the reference order
	stage.Group_symbols_instance = make(map[*Group_symbol]*Group_symbol)
	for instance := range stage.Group_symbols {
		_copy := instance.GongCopy().(*Group_symbol)
		stage.Group_symbols_reference[instance] = _copy
		stage.Group_symbols_instance[_copy] = instance
		stage.Group_symbols_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Groupings_reference = make(map[*Grouping]*Grouping)
	stage.Groupings_referenceOrder = make(map[*Grouping]uint) // diff Unstage needs the reference order
	stage.Groupings_instance = make(map[*Grouping]*Grouping)
	for instance := range stage.Groupings {
		_copy := instance.GongCopy().(*Grouping)
		stage.Groupings_reference[instance] = _copy
		stage.Groupings_instance[_copy] = instance
		stage.Groupings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Hammer_on_pull_offs_reference = make(map[*Hammer_on_pull_off]*Hammer_on_pull_off)
	stage.Hammer_on_pull_offs_referenceOrder = make(map[*Hammer_on_pull_off]uint) // diff Unstage needs the reference order
	stage.Hammer_on_pull_offs_instance = make(map[*Hammer_on_pull_off]*Hammer_on_pull_off)
	for instance := range stage.Hammer_on_pull_offs {
		_copy := instance.GongCopy().(*Hammer_on_pull_off)
		stage.Hammer_on_pull_offs_reference[instance] = _copy
		stage.Hammer_on_pull_offs_instance[_copy] = instance
		stage.Hammer_on_pull_offs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Handbells_reference = make(map[*Handbell]*Handbell)
	stage.Handbells_referenceOrder = make(map[*Handbell]uint) // diff Unstage needs the reference order
	stage.Handbells_instance = make(map[*Handbell]*Handbell)
	for instance := range stage.Handbells {
		_copy := instance.GongCopy().(*Handbell)
		stage.Handbells_reference[instance] = _copy
		stage.Handbells_instance[_copy] = instance
		stage.Handbells_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Harmon_closeds_reference = make(map[*Harmon_closed]*Harmon_closed)
	stage.Harmon_closeds_referenceOrder = make(map[*Harmon_closed]uint) // diff Unstage needs the reference order
	stage.Harmon_closeds_instance = make(map[*Harmon_closed]*Harmon_closed)
	for instance := range stage.Harmon_closeds {
		_copy := instance.GongCopy().(*Harmon_closed)
		stage.Harmon_closeds_reference[instance] = _copy
		stage.Harmon_closeds_instance[_copy] = instance
		stage.Harmon_closeds_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Harmon_mutes_reference = make(map[*Harmon_mute]*Harmon_mute)
	stage.Harmon_mutes_referenceOrder = make(map[*Harmon_mute]uint) // diff Unstage needs the reference order
	stage.Harmon_mutes_instance = make(map[*Harmon_mute]*Harmon_mute)
	for instance := range stage.Harmon_mutes {
		_copy := instance.GongCopy().(*Harmon_mute)
		stage.Harmon_mutes_reference[instance] = _copy
		stage.Harmon_mutes_instance[_copy] = instance
		stage.Harmon_mutes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Harmonics_reference = make(map[*Harmonic]*Harmonic)
	stage.Harmonics_referenceOrder = make(map[*Harmonic]uint) // diff Unstage needs the reference order
	stage.Harmonics_instance = make(map[*Harmonic]*Harmonic)
	for instance := range stage.Harmonics {
		_copy := instance.GongCopy().(*Harmonic)
		stage.Harmonics_reference[instance] = _copy
		stage.Harmonics_instance[_copy] = instance
		stage.Harmonics_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Harmonys_reference = make(map[*Harmony]*Harmony)
	stage.Harmonys_referenceOrder = make(map[*Harmony]uint) // diff Unstage needs the reference order
	stage.Harmonys_instance = make(map[*Harmony]*Harmony)
	for instance := range stage.Harmonys {
		_copy := instance.GongCopy().(*Harmony)
		stage.Harmonys_reference[instance] = _copy
		stage.Harmonys_instance[_copy] = instance
		stage.Harmonys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Harmony_alters_reference = make(map[*Harmony_alter]*Harmony_alter)
	stage.Harmony_alters_referenceOrder = make(map[*Harmony_alter]uint) // diff Unstage needs the reference order
	stage.Harmony_alters_instance = make(map[*Harmony_alter]*Harmony_alter)
	for instance := range stage.Harmony_alters {
		_copy := instance.GongCopy().(*Harmony_alter)
		stage.Harmony_alters_reference[instance] = _copy
		stage.Harmony_alters_instance[_copy] = instance
		stage.Harmony_alters_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Harp_pedalss_reference = make(map[*Harp_pedals]*Harp_pedals)
	stage.Harp_pedalss_referenceOrder = make(map[*Harp_pedals]uint) // diff Unstage needs the reference order
	stage.Harp_pedalss_instance = make(map[*Harp_pedals]*Harp_pedals)
	for instance := range stage.Harp_pedalss {
		_copy := instance.GongCopy().(*Harp_pedals)
		stage.Harp_pedalss_reference[instance] = _copy
		stage.Harp_pedalss_instance[_copy] = instance
		stage.Harp_pedalss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Heel_toes_reference = make(map[*Heel_toe]*Heel_toe)
	stage.Heel_toes_referenceOrder = make(map[*Heel_toe]uint) // diff Unstage needs the reference order
	stage.Heel_toes_instance = make(map[*Heel_toe]*Heel_toe)
	for instance := range stage.Heel_toes {
		_copy := instance.GongCopy().(*Heel_toe)
		stage.Heel_toes_reference[instance] = _copy
		stage.Heel_toes_instance[_copy] = instance
		stage.Heel_toes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Holes_reference = make(map[*Hole]*Hole)
	stage.Holes_referenceOrder = make(map[*Hole]uint) // diff Unstage needs the reference order
	stage.Holes_instance = make(map[*Hole]*Hole)
	for instance := range stage.Holes {
		_copy := instance.GongCopy().(*Hole)
		stage.Holes_reference[instance] = _copy
		stage.Holes_instance[_copy] = instance
		stage.Holes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Hole_closeds_reference = make(map[*Hole_closed]*Hole_closed)
	stage.Hole_closeds_referenceOrder = make(map[*Hole_closed]uint) // diff Unstage needs the reference order
	stage.Hole_closeds_instance = make(map[*Hole_closed]*Hole_closed)
	for instance := range stage.Hole_closeds {
		_copy := instance.GongCopy().(*Hole_closed)
		stage.Hole_closeds_reference[instance] = _copy
		stage.Hole_closeds_instance[_copy] = instance
		stage.Hole_closeds_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Horizontal_turns_reference = make(map[*Horizontal_turn]*Horizontal_turn)
	stage.Horizontal_turns_referenceOrder = make(map[*Horizontal_turn]uint) // diff Unstage needs the reference order
	stage.Horizontal_turns_instance = make(map[*Horizontal_turn]*Horizontal_turn)
	for instance := range stage.Horizontal_turns {
		_copy := instance.GongCopy().(*Horizontal_turn)
		stage.Horizontal_turns_reference[instance] = _copy
		stage.Horizontal_turns_instance[_copy] = instance
		stage.Horizontal_turns_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Identifications_reference = make(map[*Identification]*Identification)
	stage.Identifications_referenceOrder = make(map[*Identification]uint) // diff Unstage needs the reference order
	stage.Identifications_instance = make(map[*Identification]*Identification)
	for instance := range stage.Identifications {
		_copy := instance.GongCopy().(*Identification)
		stage.Identifications_reference[instance] = _copy
		stage.Identifications_instance[_copy] = instance
		stage.Identifications_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Images_reference = make(map[*Image]*Image)
	stage.Images_referenceOrder = make(map[*Image]uint) // diff Unstage needs the reference order
	stage.Images_instance = make(map[*Image]*Image)
	for instance := range stage.Images {
		_copy := instance.GongCopy().(*Image)
		stage.Images_reference[instance] = _copy
		stage.Images_instance[_copy] = instance
		stage.Images_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Instruments_reference = make(map[*Instrument]*Instrument)
	stage.Instruments_referenceOrder = make(map[*Instrument]uint) // diff Unstage needs the reference order
	stage.Instruments_instance = make(map[*Instrument]*Instrument)
	for instance := range stage.Instruments {
		_copy := instance.GongCopy().(*Instrument)
		stage.Instruments_reference[instance] = _copy
		stage.Instruments_instance[_copy] = instance
		stage.Instruments_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Instrument_changes_reference = make(map[*Instrument_change]*Instrument_change)
	stage.Instrument_changes_referenceOrder = make(map[*Instrument_change]uint) // diff Unstage needs the reference order
	stage.Instrument_changes_instance = make(map[*Instrument_change]*Instrument_change)
	for instance := range stage.Instrument_changes {
		_copy := instance.GongCopy().(*Instrument_change)
		stage.Instrument_changes_reference[instance] = _copy
		stage.Instrument_changes_instance[_copy] = instance
		stage.Instrument_changes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Instrument_links_reference = make(map[*Instrument_link]*Instrument_link)
	stage.Instrument_links_referenceOrder = make(map[*Instrument_link]uint) // diff Unstage needs the reference order
	stage.Instrument_links_instance = make(map[*Instrument_link]*Instrument_link)
	for instance := range stage.Instrument_links {
		_copy := instance.GongCopy().(*Instrument_link)
		stage.Instrument_links_reference[instance] = _copy
		stage.Instrument_links_instance[_copy] = instance
		stage.Instrument_links_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Interchangeables_reference = make(map[*Interchangeable]*Interchangeable)
	stage.Interchangeables_referenceOrder = make(map[*Interchangeable]uint) // diff Unstage needs the reference order
	stage.Interchangeables_instance = make(map[*Interchangeable]*Interchangeable)
	for instance := range stage.Interchangeables {
		_copy := instance.GongCopy().(*Interchangeable)
		stage.Interchangeables_reference[instance] = _copy
		stage.Interchangeables_instance[_copy] = instance
		stage.Interchangeables_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Inversions_reference = make(map[*Inversion]*Inversion)
	stage.Inversions_referenceOrder = make(map[*Inversion]uint) // diff Unstage needs the reference order
	stage.Inversions_instance = make(map[*Inversion]*Inversion)
	for instance := range stage.Inversions {
		_copy := instance.GongCopy().(*Inversion)
		stage.Inversions_reference[instance] = _copy
		stage.Inversions_instance[_copy] = instance
		stage.Inversions_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Keys_reference = make(map[*Key]*Key)
	stage.Keys_referenceOrder = make(map[*Key]uint) // diff Unstage needs the reference order
	stage.Keys_instance = make(map[*Key]*Key)
	for instance := range stage.Keys {
		_copy := instance.GongCopy().(*Key)
		stage.Keys_reference[instance] = _copy
		stage.Keys_instance[_copy] = instance
		stage.Keys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Key_accidentals_reference = make(map[*Key_accidental]*Key_accidental)
	stage.Key_accidentals_referenceOrder = make(map[*Key_accidental]uint) // diff Unstage needs the reference order
	stage.Key_accidentals_instance = make(map[*Key_accidental]*Key_accidental)
	for instance := range stage.Key_accidentals {
		_copy := instance.GongCopy().(*Key_accidental)
		stage.Key_accidentals_reference[instance] = _copy
		stage.Key_accidentals_instance[_copy] = instance
		stage.Key_accidentals_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Key_octaves_reference = make(map[*Key_octave]*Key_octave)
	stage.Key_octaves_referenceOrder = make(map[*Key_octave]uint) // diff Unstage needs the reference order
	stage.Key_octaves_instance = make(map[*Key_octave]*Key_octave)
	for instance := range stage.Key_octaves {
		_copy := instance.GongCopy().(*Key_octave)
		stage.Key_octaves_reference[instance] = _copy
		stage.Key_octaves_instance[_copy] = instance
		stage.Key_octaves_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Kinds_reference = make(map[*Kind]*Kind)
	stage.Kinds_referenceOrder = make(map[*Kind]uint) // diff Unstage needs the reference order
	stage.Kinds_instance = make(map[*Kind]*Kind)
	for instance := range stage.Kinds {
		_copy := instance.GongCopy().(*Kind)
		stage.Kinds_reference[instance] = _copy
		stage.Kinds_instance[_copy] = instance
		stage.Kinds_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Levels_reference = make(map[*Level]*Level)
	stage.Levels_referenceOrder = make(map[*Level]uint) // diff Unstage needs the reference order
	stage.Levels_instance = make(map[*Level]*Level)
	for instance := range stage.Levels {
		_copy := instance.GongCopy().(*Level)
		stage.Levels_reference[instance] = _copy
		stage.Levels_instance[_copy] = instance
		stage.Levels_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Line_details_reference = make(map[*Line_detail]*Line_detail)
	stage.Line_details_referenceOrder = make(map[*Line_detail]uint) // diff Unstage needs the reference order
	stage.Line_details_instance = make(map[*Line_detail]*Line_detail)
	for instance := range stage.Line_details {
		_copy := instance.GongCopy().(*Line_detail)
		stage.Line_details_reference[instance] = _copy
		stage.Line_details_instance[_copy] = instance
		stage.Line_details_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Line_widths_reference = make(map[*Line_width]*Line_width)
	stage.Line_widths_referenceOrder = make(map[*Line_width]uint) // diff Unstage needs the reference order
	stage.Line_widths_instance = make(map[*Line_width]*Line_width)
	for instance := range stage.Line_widths {
		_copy := instance.GongCopy().(*Line_width)
		stage.Line_widths_reference[instance] = _copy
		stage.Line_widths_instance[_copy] = instance
		stage.Line_widths_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Links_reference = make(map[*Link]*Link)
	stage.Links_referenceOrder = make(map[*Link]uint) // diff Unstage needs the reference order
	stage.Links_instance = make(map[*Link]*Link)
	for instance := range stage.Links {
		_copy := instance.GongCopy().(*Link)
		stage.Links_reference[instance] = _copy
		stage.Links_instance[_copy] = instance
		stage.Links_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Listens_reference = make(map[*Listen]*Listen)
	stage.Listens_referenceOrder = make(map[*Listen]uint) // diff Unstage needs the reference order
	stage.Listens_instance = make(map[*Listen]*Listen)
	for instance := range stage.Listens {
		_copy := instance.GongCopy().(*Listen)
		stage.Listens_reference[instance] = _copy
		stage.Listens_instance[_copy] = instance
		stage.Listens_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Listenings_reference = make(map[*Listening]*Listening)
	stage.Listenings_referenceOrder = make(map[*Listening]uint) // diff Unstage needs the reference order
	stage.Listenings_instance = make(map[*Listening]*Listening)
	for instance := range stage.Listenings {
		_copy := instance.GongCopy().(*Listening)
		stage.Listenings_reference[instance] = _copy
		stage.Listenings_instance[_copy] = instance
		stage.Listenings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Lyrics_reference = make(map[*Lyric]*Lyric)
	stage.Lyrics_referenceOrder = make(map[*Lyric]uint) // diff Unstage needs the reference order
	stage.Lyrics_instance = make(map[*Lyric]*Lyric)
	for instance := range stage.Lyrics {
		_copy := instance.GongCopy().(*Lyric)
		stage.Lyrics_reference[instance] = _copy
		stage.Lyrics_instance[_copy] = instance
		stage.Lyrics_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Lyric_fonts_reference = make(map[*Lyric_font]*Lyric_font)
	stage.Lyric_fonts_referenceOrder = make(map[*Lyric_font]uint) // diff Unstage needs the reference order
	stage.Lyric_fonts_instance = make(map[*Lyric_font]*Lyric_font)
	for instance := range stage.Lyric_fonts {
		_copy := instance.GongCopy().(*Lyric_font)
		stage.Lyric_fonts_reference[instance] = _copy
		stage.Lyric_fonts_instance[_copy] = instance
		stage.Lyric_fonts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Lyric_languages_reference = make(map[*Lyric_language]*Lyric_language)
	stage.Lyric_languages_referenceOrder = make(map[*Lyric_language]uint) // diff Unstage needs the reference order
	stage.Lyric_languages_instance = make(map[*Lyric_language]*Lyric_language)
	for instance := range stage.Lyric_languages {
		_copy := instance.GongCopy().(*Lyric_language)
		stage.Lyric_languages_reference[instance] = _copy
		stage.Lyric_languages_instance[_copy] = instance
		stage.Lyric_languages_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Measure_layouts_reference = make(map[*Measure_layout]*Measure_layout)
	stage.Measure_layouts_referenceOrder = make(map[*Measure_layout]uint) // diff Unstage needs the reference order
	stage.Measure_layouts_instance = make(map[*Measure_layout]*Measure_layout)
	for instance := range stage.Measure_layouts {
		_copy := instance.GongCopy().(*Measure_layout)
		stage.Measure_layouts_reference[instance] = _copy
		stage.Measure_layouts_instance[_copy] = instance
		stage.Measure_layouts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Measure_numberings_reference = make(map[*Measure_numbering]*Measure_numbering)
	stage.Measure_numberings_referenceOrder = make(map[*Measure_numbering]uint) // diff Unstage needs the reference order
	stage.Measure_numberings_instance = make(map[*Measure_numbering]*Measure_numbering)
	for instance := range stage.Measure_numberings {
		_copy := instance.GongCopy().(*Measure_numbering)
		stage.Measure_numberings_reference[instance] = _copy
		stage.Measure_numberings_instance[_copy] = instance
		stage.Measure_numberings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Measure_repeats_reference = make(map[*Measure_repeat]*Measure_repeat)
	stage.Measure_repeats_referenceOrder = make(map[*Measure_repeat]uint) // diff Unstage needs the reference order
	stage.Measure_repeats_instance = make(map[*Measure_repeat]*Measure_repeat)
	for instance := range stage.Measure_repeats {
		_copy := instance.GongCopy().(*Measure_repeat)
		stage.Measure_repeats_reference[instance] = _copy
		stage.Measure_repeats_instance[_copy] = instance
		stage.Measure_repeats_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Measure_styles_reference = make(map[*Measure_style]*Measure_style)
	stage.Measure_styles_referenceOrder = make(map[*Measure_style]uint) // diff Unstage needs the reference order
	stage.Measure_styles_instance = make(map[*Measure_style]*Measure_style)
	for instance := range stage.Measure_styles {
		_copy := instance.GongCopy().(*Measure_style)
		stage.Measure_styles_reference[instance] = _copy
		stage.Measure_styles_instance[_copy] = instance
		stage.Measure_styles_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Membranes_reference = make(map[*Membrane]*Membrane)
	stage.Membranes_referenceOrder = make(map[*Membrane]uint) // diff Unstage needs the reference order
	stage.Membranes_instance = make(map[*Membrane]*Membrane)
	for instance := range stage.Membranes {
		_copy := instance.GongCopy().(*Membrane)
		stage.Membranes_reference[instance] = _copy
		stage.Membranes_instance[_copy] = instance
		stage.Membranes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Metals_reference = make(map[*Metal]*Metal)
	stage.Metals_referenceOrder = make(map[*Metal]uint) // diff Unstage needs the reference order
	stage.Metals_instance = make(map[*Metal]*Metal)
	for instance := range stage.Metals {
		_copy := instance.GongCopy().(*Metal)
		stage.Metals_reference[instance] = _copy
		stage.Metals_instance[_copy] = instance
		stage.Metals_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Metronomes_reference = make(map[*Metronome]*Metronome)
	stage.Metronomes_referenceOrder = make(map[*Metronome]uint) // diff Unstage needs the reference order
	stage.Metronomes_instance = make(map[*Metronome]*Metronome)
	for instance := range stage.Metronomes {
		_copy := instance.GongCopy().(*Metronome)
		stage.Metronomes_reference[instance] = _copy
		stage.Metronomes_instance[_copy] = instance
		stage.Metronomes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Metronome_beams_reference = make(map[*Metronome_beam]*Metronome_beam)
	stage.Metronome_beams_referenceOrder = make(map[*Metronome_beam]uint) // diff Unstage needs the reference order
	stage.Metronome_beams_instance = make(map[*Metronome_beam]*Metronome_beam)
	for instance := range stage.Metronome_beams {
		_copy := instance.GongCopy().(*Metronome_beam)
		stage.Metronome_beams_reference[instance] = _copy
		stage.Metronome_beams_instance[_copy] = instance
		stage.Metronome_beams_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Metronome_notes_reference = make(map[*Metronome_note]*Metronome_note)
	stage.Metronome_notes_referenceOrder = make(map[*Metronome_note]uint) // diff Unstage needs the reference order
	stage.Metronome_notes_instance = make(map[*Metronome_note]*Metronome_note)
	for instance := range stage.Metronome_notes {
		_copy := instance.GongCopy().(*Metronome_note)
		stage.Metronome_notes_reference[instance] = _copy
		stage.Metronome_notes_instance[_copy] = instance
		stage.Metronome_notes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Metronome_tieds_reference = make(map[*Metronome_tied]*Metronome_tied)
	stage.Metronome_tieds_referenceOrder = make(map[*Metronome_tied]uint) // diff Unstage needs the reference order
	stage.Metronome_tieds_instance = make(map[*Metronome_tied]*Metronome_tied)
	for instance := range stage.Metronome_tieds {
		_copy := instance.GongCopy().(*Metronome_tied)
		stage.Metronome_tieds_reference[instance] = _copy
		stage.Metronome_tieds_instance[_copy] = instance
		stage.Metronome_tieds_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Metronome_tuplets_reference = make(map[*Metronome_tuplet]*Metronome_tuplet)
	stage.Metronome_tuplets_referenceOrder = make(map[*Metronome_tuplet]uint) // diff Unstage needs the reference order
	stage.Metronome_tuplets_instance = make(map[*Metronome_tuplet]*Metronome_tuplet)
	for instance := range stage.Metronome_tuplets {
		_copy := instance.GongCopy().(*Metronome_tuplet)
		stage.Metronome_tuplets_reference[instance] = _copy
		stage.Metronome_tuplets_instance[_copy] = instance
		stage.Metronome_tuplets_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Midi_devices_reference = make(map[*Midi_device]*Midi_device)
	stage.Midi_devices_referenceOrder = make(map[*Midi_device]uint) // diff Unstage needs the reference order
	stage.Midi_devices_instance = make(map[*Midi_device]*Midi_device)
	for instance := range stage.Midi_devices {
		_copy := instance.GongCopy().(*Midi_device)
		stage.Midi_devices_reference[instance] = _copy
		stage.Midi_devices_instance[_copy] = instance
		stage.Midi_devices_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Midi_instruments_reference = make(map[*Midi_instrument]*Midi_instrument)
	stage.Midi_instruments_referenceOrder = make(map[*Midi_instrument]uint) // diff Unstage needs the reference order
	stage.Midi_instruments_instance = make(map[*Midi_instrument]*Midi_instrument)
	for instance := range stage.Midi_instruments {
		_copy := instance.GongCopy().(*Midi_instrument)
		stage.Midi_instruments_reference[instance] = _copy
		stage.Midi_instruments_instance[_copy] = instance
		stage.Midi_instruments_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Miscellaneouss_reference = make(map[*Miscellaneous]*Miscellaneous)
	stage.Miscellaneouss_referenceOrder = make(map[*Miscellaneous]uint) // diff Unstage needs the reference order
	stage.Miscellaneouss_instance = make(map[*Miscellaneous]*Miscellaneous)
	for instance := range stage.Miscellaneouss {
		_copy := instance.GongCopy().(*Miscellaneous)
		stage.Miscellaneouss_reference[instance] = _copy
		stage.Miscellaneouss_instance[_copy] = instance
		stage.Miscellaneouss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Miscellaneous_fields_reference = make(map[*Miscellaneous_field]*Miscellaneous_field)
	stage.Miscellaneous_fields_referenceOrder = make(map[*Miscellaneous_field]uint) // diff Unstage needs the reference order
	stage.Miscellaneous_fields_instance = make(map[*Miscellaneous_field]*Miscellaneous_field)
	for instance := range stage.Miscellaneous_fields {
		_copy := instance.GongCopy().(*Miscellaneous_field)
		stage.Miscellaneous_fields_reference[instance] = _copy
		stage.Miscellaneous_fields_instance[_copy] = instance
		stage.Miscellaneous_fields_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Mordents_reference = make(map[*Mordent]*Mordent)
	stage.Mordents_referenceOrder = make(map[*Mordent]uint) // diff Unstage needs the reference order
	stage.Mordents_instance = make(map[*Mordent]*Mordent)
	for instance := range stage.Mordents {
		_copy := instance.GongCopy().(*Mordent)
		stage.Mordents_reference[instance] = _copy
		stage.Mordents_instance[_copy] = instance
		stage.Mordents_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Multiple_rests_reference = make(map[*Multiple_rest]*Multiple_rest)
	stage.Multiple_rests_referenceOrder = make(map[*Multiple_rest]uint) // diff Unstage needs the reference order
	stage.Multiple_rests_instance = make(map[*Multiple_rest]*Multiple_rest)
	for instance := range stage.Multiple_rests {
		_copy := instance.GongCopy().(*Multiple_rest)
		stage.Multiple_rests_reference[instance] = _copy
		stage.Multiple_rests_instance[_copy] = instance
		stage.Multiple_rests_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Name_displays_reference = make(map[*Name_display]*Name_display)
	stage.Name_displays_referenceOrder = make(map[*Name_display]uint) // diff Unstage needs the reference order
	stage.Name_displays_instance = make(map[*Name_display]*Name_display)
	for instance := range stage.Name_displays {
		_copy := instance.GongCopy().(*Name_display)
		stage.Name_displays_reference[instance] = _copy
		stage.Name_displays_instance[_copy] = instance
		stage.Name_displays_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Non_arpeggiates_reference = make(map[*Non_arpeggiate]*Non_arpeggiate)
	stage.Non_arpeggiates_referenceOrder = make(map[*Non_arpeggiate]uint) // diff Unstage needs the reference order
	stage.Non_arpeggiates_instance = make(map[*Non_arpeggiate]*Non_arpeggiate)
	for instance := range stage.Non_arpeggiates {
		_copy := instance.GongCopy().(*Non_arpeggiate)
		stage.Non_arpeggiates_reference[instance] = _copy
		stage.Non_arpeggiates_instance[_copy] = instance
		stage.Non_arpeggiates_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Notationss_reference = make(map[*Notations]*Notations)
	stage.Notationss_referenceOrder = make(map[*Notations]uint) // diff Unstage needs the reference order
	stage.Notationss_instance = make(map[*Notations]*Notations)
	for instance := range stage.Notationss {
		_copy := instance.GongCopy().(*Notations)
		stage.Notationss_reference[instance] = _copy
		stage.Notationss_instance[_copy] = instance
		stage.Notationss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Notes_reference = make(map[*Note]*Note)
	stage.Notes_referenceOrder = make(map[*Note]uint) // diff Unstage needs the reference order
	stage.Notes_instance = make(map[*Note]*Note)
	for instance := range stage.Notes {
		_copy := instance.GongCopy().(*Note)
		stage.Notes_reference[instance] = _copy
		stage.Notes_instance[_copy] = instance
		stage.Notes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Note_sizes_reference = make(map[*Note_size]*Note_size)
	stage.Note_sizes_referenceOrder = make(map[*Note_size]uint) // diff Unstage needs the reference order
	stage.Note_sizes_instance = make(map[*Note_size]*Note_size)
	for instance := range stage.Note_sizes {
		_copy := instance.GongCopy().(*Note_size)
		stage.Note_sizes_reference[instance] = _copy
		stage.Note_sizes_instance[_copy] = instance
		stage.Note_sizes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Note_types_reference = make(map[*Note_type]*Note_type)
	stage.Note_types_referenceOrder = make(map[*Note_type]uint) // diff Unstage needs the reference order
	stage.Note_types_instance = make(map[*Note_type]*Note_type)
	for instance := range stage.Note_types {
		_copy := instance.GongCopy().(*Note_type)
		stage.Note_types_reference[instance] = _copy
		stage.Note_types_instance[_copy] = instance
		stage.Note_types_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Noteheads_reference = make(map[*Notehead]*Notehead)
	stage.Noteheads_referenceOrder = make(map[*Notehead]uint) // diff Unstage needs the reference order
	stage.Noteheads_instance = make(map[*Notehead]*Notehead)
	for instance := range stage.Noteheads {
		_copy := instance.GongCopy().(*Notehead)
		stage.Noteheads_reference[instance] = _copy
		stage.Noteheads_instance[_copy] = instance
		stage.Noteheads_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Notehead_texts_reference = make(map[*Notehead_text]*Notehead_text)
	stage.Notehead_texts_referenceOrder = make(map[*Notehead_text]uint) // diff Unstage needs the reference order
	stage.Notehead_texts_instance = make(map[*Notehead_text]*Notehead_text)
	for instance := range stage.Notehead_texts {
		_copy := instance.GongCopy().(*Notehead_text)
		stage.Notehead_texts_reference[instance] = _copy
		stage.Notehead_texts_instance[_copy] = instance
		stage.Notehead_texts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Numerals_reference = make(map[*Numeral]*Numeral)
	stage.Numerals_referenceOrder = make(map[*Numeral]uint) // diff Unstage needs the reference order
	stage.Numerals_instance = make(map[*Numeral]*Numeral)
	for instance := range stage.Numerals {
		_copy := instance.GongCopy().(*Numeral)
		stage.Numerals_reference[instance] = _copy
		stage.Numerals_instance[_copy] = instance
		stage.Numerals_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Numeral_keys_reference = make(map[*Numeral_key]*Numeral_key)
	stage.Numeral_keys_referenceOrder = make(map[*Numeral_key]uint) // diff Unstage needs the reference order
	stage.Numeral_keys_instance = make(map[*Numeral_key]*Numeral_key)
	for instance := range stage.Numeral_keys {
		_copy := instance.GongCopy().(*Numeral_key)
		stage.Numeral_keys_reference[instance] = _copy
		stage.Numeral_keys_instance[_copy] = instance
		stage.Numeral_keys_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Numeral_roots_reference = make(map[*Numeral_root]*Numeral_root)
	stage.Numeral_roots_referenceOrder = make(map[*Numeral_root]uint) // diff Unstage needs the reference order
	stage.Numeral_roots_instance = make(map[*Numeral_root]*Numeral_root)
	for instance := range stage.Numeral_roots {
		_copy := instance.GongCopy().(*Numeral_root)
		stage.Numeral_roots_reference[instance] = _copy
		stage.Numeral_roots_instance[_copy] = instance
		stage.Numeral_roots_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Octave_shifts_reference = make(map[*Octave_shift]*Octave_shift)
	stage.Octave_shifts_referenceOrder = make(map[*Octave_shift]uint) // diff Unstage needs the reference order
	stage.Octave_shifts_instance = make(map[*Octave_shift]*Octave_shift)
	for instance := range stage.Octave_shifts {
		_copy := instance.GongCopy().(*Octave_shift)
		stage.Octave_shifts_reference[instance] = _copy
		stage.Octave_shifts_instance[_copy] = instance
		stage.Octave_shifts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Offsets_reference = make(map[*Offset]*Offset)
	stage.Offsets_referenceOrder = make(map[*Offset]uint) // diff Unstage needs the reference order
	stage.Offsets_instance = make(map[*Offset]*Offset)
	for instance := range stage.Offsets {
		_copy := instance.GongCopy().(*Offset)
		stage.Offsets_reference[instance] = _copy
		stage.Offsets_instance[_copy] = instance
		stage.Offsets_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Opuss_reference = make(map[*Opus]*Opus)
	stage.Opuss_referenceOrder = make(map[*Opus]uint) // diff Unstage needs the reference order
	stage.Opuss_instance = make(map[*Opus]*Opus)
	for instance := range stage.Opuss {
		_copy := instance.GongCopy().(*Opus)
		stage.Opuss_reference[instance] = _copy
		stage.Opuss_instance[_copy] = instance
		stage.Opuss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Ornamentss_reference = make(map[*Ornaments]*Ornaments)
	stage.Ornamentss_referenceOrder = make(map[*Ornaments]uint) // diff Unstage needs the reference order
	stage.Ornamentss_instance = make(map[*Ornaments]*Ornaments)
	for instance := range stage.Ornamentss {
		_copy := instance.GongCopy().(*Ornaments)
		stage.Ornamentss_reference[instance] = _copy
		stage.Ornamentss_instance[_copy] = instance
		stage.Ornamentss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Other_appearances_reference = make(map[*Other_appearance]*Other_appearance)
	stage.Other_appearances_referenceOrder = make(map[*Other_appearance]uint) // diff Unstage needs the reference order
	stage.Other_appearances_instance = make(map[*Other_appearance]*Other_appearance)
	for instance := range stage.Other_appearances {
		_copy := instance.GongCopy().(*Other_appearance)
		stage.Other_appearances_reference[instance] = _copy
		stage.Other_appearances_instance[_copy] = instance
		stage.Other_appearances_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Other_directions_reference = make(map[*Other_direction]*Other_direction)
	stage.Other_directions_referenceOrder = make(map[*Other_direction]uint) // diff Unstage needs the reference order
	stage.Other_directions_instance = make(map[*Other_direction]*Other_direction)
	for instance := range stage.Other_directions {
		_copy := instance.GongCopy().(*Other_direction)
		stage.Other_directions_reference[instance] = _copy
		stage.Other_directions_instance[_copy] = instance
		stage.Other_directions_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Other_listenings_reference = make(map[*Other_listening]*Other_listening)
	stage.Other_listenings_referenceOrder = make(map[*Other_listening]uint) // diff Unstage needs the reference order
	stage.Other_listenings_instance = make(map[*Other_listening]*Other_listening)
	for instance := range stage.Other_listenings {
		_copy := instance.GongCopy().(*Other_listening)
		stage.Other_listenings_reference[instance] = _copy
		stage.Other_listenings_instance[_copy] = instance
		stage.Other_listenings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Other_notations_reference = make(map[*Other_notation]*Other_notation)
	stage.Other_notations_referenceOrder = make(map[*Other_notation]uint) // diff Unstage needs the reference order
	stage.Other_notations_instance = make(map[*Other_notation]*Other_notation)
	for instance := range stage.Other_notations {
		_copy := instance.GongCopy().(*Other_notation)
		stage.Other_notations_reference[instance] = _copy
		stage.Other_notations_instance[_copy] = instance
		stage.Other_notations_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Other_placement_texts_reference = make(map[*Other_placement_text]*Other_placement_text)
	stage.Other_placement_texts_referenceOrder = make(map[*Other_placement_text]uint) // diff Unstage needs the reference order
	stage.Other_placement_texts_instance = make(map[*Other_placement_text]*Other_placement_text)
	for instance := range stage.Other_placement_texts {
		_copy := instance.GongCopy().(*Other_placement_text)
		stage.Other_placement_texts_reference[instance] = _copy
		stage.Other_placement_texts_instance[_copy] = instance
		stage.Other_placement_texts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Other_plays_reference = make(map[*Other_play]*Other_play)
	stage.Other_plays_referenceOrder = make(map[*Other_play]uint) // diff Unstage needs the reference order
	stage.Other_plays_instance = make(map[*Other_play]*Other_play)
	for instance := range stage.Other_plays {
		_copy := instance.GongCopy().(*Other_play)
		stage.Other_plays_reference[instance] = _copy
		stage.Other_plays_instance[_copy] = instance
		stage.Other_plays_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Other_texts_reference = make(map[*Other_text]*Other_text)
	stage.Other_texts_referenceOrder = make(map[*Other_text]uint) // diff Unstage needs the reference order
	stage.Other_texts_instance = make(map[*Other_text]*Other_text)
	for instance := range stage.Other_texts {
		_copy := instance.GongCopy().(*Other_text)
		stage.Other_texts_reference[instance] = _copy
		stage.Other_texts_instance[_copy] = instance
		stage.Other_texts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Page_layouts_reference = make(map[*Page_layout]*Page_layout)
	stage.Page_layouts_referenceOrder = make(map[*Page_layout]uint) // diff Unstage needs the reference order
	stage.Page_layouts_instance = make(map[*Page_layout]*Page_layout)
	for instance := range stage.Page_layouts {
		_copy := instance.GongCopy().(*Page_layout)
		stage.Page_layouts_reference[instance] = _copy
		stage.Page_layouts_instance[_copy] = instance
		stage.Page_layouts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Page_marginss_reference = make(map[*Page_margins]*Page_margins)
	stage.Page_marginss_referenceOrder = make(map[*Page_margins]uint) // diff Unstage needs the reference order
	stage.Page_marginss_instance = make(map[*Page_margins]*Page_margins)
	for instance := range stage.Page_marginss {
		_copy := instance.GongCopy().(*Page_margins)
		stage.Page_marginss_reference[instance] = _copy
		stage.Page_marginss_instance[_copy] = instance
		stage.Page_marginss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Part_clefs_reference = make(map[*Part_clef]*Part_clef)
	stage.Part_clefs_referenceOrder = make(map[*Part_clef]uint) // diff Unstage needs the reference order
	stage.Part_clefs_instance = make(map[*Part_clef]*Part_clef)
	for instance := range stage.Part_clefs {
		_copy := instance.GongCopy().(*Part_clef)
		stage.Part_clefs_reference[instance] = _copy
		stage.Part_clefs_instance[_copy] = instance
		stage.Part_clefs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Part_groups_reference = make(map[*Part_group]*Part_group)
	stage.Part_groups_referenceOrder = make(map[*Part_group]uint) // diff Unstage needs the reference order
	stage.Part_groups_instance = make(map[*Part_group]*Part_group)
	for instance := range stage.Part_groups {
		_copy := instance.GongCopy().(*Part_group)
		stage.Part_groups_reference[instance] = _copy
		stage.Part_groups_instance[_copy] = instance
		stage.Part_groups_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Part_links_reference = make(map[*Part_link]*Part_link)
	stage.Part_links_referenceOrder = make(map[*Part_link]uint) // diff Unstage needs the reference order
	stage.Part_links_instance = make(map[*Part_link]*Part_link)
	for instance := range stage.Part_links {
		_copy := instance.GongCopy().(*Part_link)
		stage.Part_links_reference[instance] = _copy
		stage.Part_links_instance[_copy] = instance
		stage.Part_links_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Part_lists_reference = make(map[*Part_list]*Part_list)
	stage.Part_lists_referenceOrder = make(map[*Part_list]uint) // diff Unstage needs the reference order
	stage.Part_lists_instance = make(map[*Part_list]*Part_list)
	for instance := range stage.Part_lists {
		_copy := instance.GongCopy().(*Part_list)
		stage.Part_lists_reference[instance] = _copy
		stage.Part_lists_instance[_copy] = instance
		stage.Part_lists_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Part_names_reference = make(map[*Part_name]*Part_name)
	stage.Part_names_referenceOrder = make(map[*Part_name]uint) // diff Unstage needs the reference order
	stage.Part_names_instance = make(map[*Part_name]*Part_name)
	for instance := range stage.Part_names {
		_copy := instance.GongCopy().(*Part_name)
		stage.Part_names_reference[instance] = _copy
		stage.Part_names_instance[_copy] = instance
		stage.Part_names_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Part_symbols_reference = make(map[*Part_symbol]*Part_symbol)
	stage.Part_symbols_referenceOrder = make(map[*Part_symbol]uint) // diff Unstage needs the reference order
	stage.Part_symbols_instance = make(map[*Part_symbol]*Part_symbol)
	for instance := range stage.Part_symbols {
		_copy := instance.GongCopy().(*Part_symbol)
		stage.Part_symbols_reference[instance] = _copy
		stage.Part_symbols_instance[_copy] = instance
		stage.Part_symbols_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Part_transposes_reference = make(map[*Part_transpose]*Part_transpose)
	stage.Part_transposes_referenceOrder = make(map[*Part_transpose]uint) // diff Unstage needs the reference order
	stage.Part_transposes_instance = make(map[*Part_transpose]*Part_transpose)
	for instance := range stage.Part_transposes {
		_copy := instance.GongCopy().(*Part_transpose)
		stage.Part_transposes_reference[instance] = _copy
		stage.Part_transposes_instance[_copy] = instance
		stage.Part_transposes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Pedals_reference = make(map[*Pedal]*Pedal)
	stage.Pedals_referenceOrder = make(map[*Pedal]uint) // diff Unstage needs the reference order
	stage.Pedals_instance = make(map[*Pedal]*Pedal)
	for instance := range stage.Pedals {
		_copy := instance.GongCopy().(*Pedal)
		stage.Pedals_reference[instance] = _copy
		stage.Pedals_instance[_copy] = instance
		stage.Pedals_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Pedal_tunings_reference = make(map[*Pedal_tuning]*Pedal_tuning)
	stage.Pedal_tunings_referenceOrder = make(map[*Pedal_tuning]uint) // diff Unstage needs the reference order
	stage.Pedal_tunings_instance = make(map[*Pedal_tuning]*Pedal_tuning)
	for instance := range stage.Pedal_tunings {
		_copy := instance.GongCopy().(*Pedal_tuning)
		stage.Pedal_tunings_reference[instance] = _copy
		stage.Pedal_tunings_instance[_copy] = instance
		stage.Pedal_tunings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Per_minutes_reference = make(map[*Per_minute]*Per_minute)
	stage.Per_minutes_referenceOrder = make(map[*Per_minute]uint) // diff Unstage needs the reference order
	stage.Per_minutes_instance = make(map[*Per_minute]*Per_minute)
	for instance := range stage.Per_minutes {
		_copy := instance.GongCopy().(*Per_minute)
		stage.Per_minutes_reference[instance] = _copy
		stage.Per_minutes_instance[_copy] = instance
		stage.Per_minutes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Percussions_reference = make(map[*Percussion]*Percussion)
	stage.Percussions_referenceOrder = make(map[*Percussion]uint) // diff Unstage needs the reference order
	stage.Percussions_instance = make(map[*Percussion]*Percussion)
	for instance := range stage.Percussions {
		_copy := instance.GongCopy().(*Percussion)
		stage.Percussions_reference[instance] = _copy
		stage.Percussions_instance[_copy] = instance
		stage.Percussions_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Pitchs_reference = make(map[*Pitch]*Pitch)
	stage.Pitchs_referenceOrder = make(map[*Pitch]uint) // diff Unstage needs the reference order
	stage.Pitchs_instance = make(map[*Pitch]*Pitch)
	for instance := range stage.Pitchs {
		_copy := instance.GongCopy().(*Pitch)
		stage.Pitchs_reference[instance] = _copy
		stage.Pitchs_instance[_copy] = instance
		stage.Pitchs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Pitcheds_reference = make(map[*Pitched]*Pitched)
	stage.Pitcheds_referenceOrder = make(map[*Pitched]uint) // diff Unstage needs the reference order
	stage.Pitcheds_instance = make(map[*Pitched]*Pitched)
	for instance := range stage.Pitcheds {
		_copy := instance.GongCopy().(*Pitched)
		stage.Pitcheds_reference[instance] = _copy
		stage.Pitcheds_instance[_copy] = instance
		stage.Pitcheds_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Placement_texts_reference = make(map[*Placement_text]*Placement_text)
	stage.Placement_texts_referenceOrder = make(map[*Placement_text]uint) // diff Unstage needs the reference order
	stage.Placement_texts_instance = make(map[*Placement_text]*Placement_text)
	for instance := range stage.Placement_texts {
		_copy := instance.GongCopy().(*Placement_text)
		stage.Placement_texts_reference[instance] = _copy
		stage.Placement_texts_instance[_copy] = instance
		stage.Placement_texts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Plays_reference = make(map[*Play]*Play)
	stage.Plays_referenceOrder = make(map[*Play]uint) // diff Unstage needs the reference order
	stage.Plays_instance = make(map[*Play]*Play)
	for instance := range stage.Plays {
		_copy := instance.GongCopy().(*Play)
		stage.Plays_reference[instance] = _copy
		stage.Plays_instance[_copy] = instance
		stage.Plays_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Players_reference = make(map[*Player]*Player)
	stage.Players_referenceOrder = make(map[*Player]uint) // diff Unstage needs the reference order
	stage.Players_instance = make(map[*Player]*Player)
	for instance := range stage.Players {
		_copy := instance.GongCopy().(*Player)
		stage.Players_reference[instance] = _copy
		stage.Players_instance[_copy] = instance
		stage.Players_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Principal_voices_reference = make(map[*Principal_voice]*Principal_voice)
	stage.Principal_voices_referenceOrder = make(map[*Principal_voice]uint) // diff Unstage needs the reference order
	stage.Principal_voices_instance = make(map[*Principal_voice]*Principal_voice)
	for instance := range stage.Principal_voices {
		_copy := instance.GongCopy().(*Principal_voice)
		stage.Principal_voices_reference[instance] = _copy
		stage.Principal_voices_instance[_copy] = instance
		stage.Principal_voices_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Prints_reference = make(map[*Print]*Print)
	stage.Prints_referenceOrder = make(map[*Print]uint) // diff Unstage needs the reference order
	stage.Prints_instance = make(map[*Print]*Print)
	for instance := range stage.Prints {
		_copy := instance.GongCopy().(*Print)
		stage.Prints_reference[instance] = _copy
		stage.Prints_instance[_copy] = instance
		stage.Prints_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Releases_reference = make(map[*Release]*Release)
	stage.Releases_referenceOrder = make(map[*Release]uint) // diff Unstage needs the reference order
	stage.Releases_instance = make(map[*Release]*Release)
	for instance := range stage.Releases {
		_copy := instance.GongCopy().(*Release)
		stage.Releases_reference[instance] = _copy
		stage.Releases_instance[_copy] = instance
		stage.Releases_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Repeats_reference = make(map[*Repeat]*Repeat)
	stage.Repeats_referenceOrder = make(map[*Repeat]uint) // diff Unstage needs the reference order
	stage.Repeats_instance = make(map[*Repeat]*Repeat)
	for instance := range stage.Repeats {
		_copy := instance.GongCopy().(*Repeat)
		stage.Repeats_reference[instance] = _copy
		stage.Repeats_instance[_copy] = instance
		stage.Repeats_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Rests_reference = make(map[*Rest]*Rest)
	stage.Rests_referenceOrder = make(map[*Rest]uint) // diff Unstage needs the reference order
	stage.Rests_instance = make(map[*Rest]*Rest)
	for instance := range stage.Rests {
		_copy := instance.GongCopy().(*Rest)
		stage.Rests_reference[instance] = _copy
		stage.Rests_instance[_copy] = instance
		stage.Rests_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Roots_reference = make(map[*Root]*Root)
	stage.Roots_referenceOrder = make(map[*Root]uint) // diff Unstage needs the reference order
	stage.Roots_instance = make(map[*Root]*Root)
	for instance := range stage.Roots {
		_copy := instance.GongCopy().(*Root)
		stage.Roots_reference[instance] = _copy
		stage.Roots_instance[_copy] = instance
		stage.Roots_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Root_steps_reference = make(map[*Root_step]*Root_step)
	stage.Root_steps_referenceOrder = make(map[*Root_step]uint) // diff Unstage needs the reference order
	stage.Root_steps_instance = make(map[*Root_step]*Root_step)
	for instance := range stage.Root_steps {
		_copy := instance.GongCopy().(*Root_step)
		stage.Root_steps_reference[instance] = _copy
		stage.Root_steps_instance[_copy] = instance
		stage.Root_steps_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Scalings_reference = make(map[*Scaling]*Scaling)
	stage.Scalings_referenceOrder = make(map[*Scaling]uint) // diff Unstage needs the reference order
	stage.Scalings_instance = make(map[*Scaling]*Scaling)
	for instance := range stage.Scalings {
		_copy := instance.GongCopy().(*Scaling)
		stage.Scalings_reference[instance] = _copy
		stage.Scalings_instance[_copy] = instance
		stage.Scalings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Scordaturas_reference = make(map[*Scordatura]*Scordatura)
	stage.Scordaturas_referenceOrder = make(map[*Scordatura]uint) // diff Unstage needs the reference order
	stage.Scordaturas_instance = make(map[*Scordatura]*Scordatura)
	for instance := range stage.Scordaturas {
		_copy := instance.GongCopy().(*Scordatura)
		stage.Scordaturas_reference[instance] = _copy
		stage.Scordaturas_instance[_copy] = instance
		stage.Scordaturas_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Score_instruments_reference = make(map[*Score_instrument]*Score_instrument)
	stage.Score_instruments_referenceOrder = make(map[*Score_instrument]uint) // diff Unstage needs the reference order
	stage.Score_instruments_instance = make(map[*Score_instrument]*Score_instrument)
	for instance := range stage.Score_instruments {
		_copy := instance.GongCopy().(*Score_instrument)
		stage.Score_instruments_reference[instance] = _copy
		stage.Score_instruments_instance[_copy] = instance
		stage.Score_instruments_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Score_parts_reference = make(map[*Score_part]*Score_part)
	stage.Score_parts_referenceOrder = make(map[*Score_part]uint) // diff Unstage needs the reference order
	stage.Score_parts_instance = make(map[*Score_part]*Score_part)
	for instance := range stage.Score_parts {
		_copy := instance.GongCopy().(*Score_part)
		stage.Score_parts_reference[instance] = _copy
		stage.Score_parts_instance[_copy] = instance
		stage.Score_parts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Score_partwises_reference = make(map[*Score_partwise]*Score_partwise)
	stage.Score_partwises_referenceOrder = make(map[*Score_partwise]uint) // diff Unstage needs the reference order
	stage.Score_partwises_instance = make(map[*Score_partwise]*Score_partwise)
	for instance := range stage.Score_partwises {
		_copy := instance.GongCopy().(*Score_partwise)
		stage.Score_partwises_reference[instance] = _copy
		stage.Score_partwises_instance[_copy] = instance
		stage.Score_partwises_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Score_timewises_reference = make(map[*Score_timewise]*Score_timewise)
	stage.Score_timewises_referenceOrder = make(map[*Score_timewise]uint) // diff Unstage needs the reference order
	stage.Score_timewises_instance = make(map[*Score_timewise]*Score_timewise)
	for instance := range stage.Score_timewises {
		_copy := instance.GongCopy().(*Score_timewise)
		stage.Score_timewises_reference[instance] = _copy
		stage.Score_timewises_instance[_copy] = instance
		stage.Score_timewises_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Segnos_reference = make(map[*Segno]*Segno)
	stage.Segnos_referenceOrder = make(map[*Segno]uint) // diff Unstage needs the reference order
	stage.Segnos_instance = make(map[*Segno]*Segno)
	for instance := range stage.Segnos {
		_copy := instance.GongCopy().(*Segno)
		stage.Segnos_reference[instance] = _copy
		stage.Segnos_instance[_copy] = instance
		stage.Segnos_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Slashs_reference = make(map[*Slash]*Slash)
	stage.Slashs_referenceOrder = make(map[*Slash]uint) // diff Unstage needs the reference order
	stage.Slashs_instance = make(map[*Slash]*Slash)
	for instance := range stage.Slashs {
		_copy := instance.GongCopy().(*Slash)
		stage.Slashs_reference[instance] = _copy
		stage.Slashs_instance[_copy] = instance
		stage.Slashs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Slides_reference = make(map[*Slide]*Slide)
	stage.Slides_referenceOrder = make(map[*Slide]uint) // diff Unstage needs the reference order
	stage.Slides_instance = make(map[*Slide]*Slide)
	for instance := range stage.Slides {
		_copy := instance.GongCopy().(*Slide)
		stage.Slides_reference[instance] = _copy
		stage.Slides_instance[_copy] = instance
		stage.Slides_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Slurs_reference = make(map[*Slur]*Slur)
	stage.Slurs_referenceOrder = make(map[*Slur]uint) // diff Unstage needs the reference order
	stage.Slurs_instance = make(map[*Slur]*Slur)
	for instance := range stage.Slurs {
		_copy := instance.GongCopy().(*Slur)
		stage.Slurs_reference[instance] = _copy
		stage.Slurs_instance[_copy] = instance
		stage.Slurs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Sounds_reference = make(map[*Sound]*Sound)
	stage.Sounds_referenceOrder = make(map[*Sound]uint) // diff Unstage needs the reference order
	stage.Sounds_instance = make(map[*Sound]*Sound)
	for instance := range stage.Sounds {
		_copy := instance.GongCopy().(*Sound)
		stage.Sounds_reference[instance] = _copy
		stage.Sounds_instance[_copy] = instance
		stage.Sounds_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Staff_detailss_reference = make(map[*Staff_details]*Staff_details)
	stage.Staff_detailss_referenceOrder = make(map[*Staff_details]uint) // diff Unstage needs the reference order
	stage.Staff_detailss_instance = make(map[*Staff_details]*Staff_details)
	for instance := range stage.Staff_detailss {
		_copy := instance.GongCopy().(*Staff_details)
		stage.Staff_detailss_reference[instance] = _copy
		stage.Staff_detailss_instance[_copy] = instance
		stage.Staff_detailss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Staff_divides_reference = make(map[*Staff_divide]*Staff_divide)
	stage.Staff_divides_referenceOrder = make(map[*Staff_divide]uint) // diff Unstage needs the reference order
	stage.Staff_divides_instance = make(map[*Staff_divide]*Staff_divide)
	for instance := range stage.Staff_divides {
		_copy := instance.GongCopy().(*Staff_divide)
		stage.Staff_divides_reference[instance] = _copy
		stage.Staff_divides_instance[_copy] = instance
		stage.Staff_divides_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Staff_layouts_reference = make(map[*Staff_layout]*Staff_layout)
	stage.Staff_layouts_referenceOrder = make(map[*Staff_layout]uint) // diff Unstage needs the reference order
	stage.Staff_layouts_instance = make(map[*Staff_layout]*Staff_layout)
	for instance := range stage.Staff_layouts {
		_copy := instance.GongCopy().(*Staff_layout)
		stage.Staff_layouts_reference[instance] = _copy
		stage.Staff_layouts_instance[_copy] = instance
		stage.Staff_layouts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Staff_sizes_reference = make(map[*Staff_size]*Staff_size)
	stage.Staff_sizes_referenceOrder = make(map[*Staff_size]uint) // diff Unstage needs the reference order
	stage.Staff_sizes_instance = make(map[*Staff_size]*Staff_size)
	for instance := range stage.Staff_sizes {
		_copy := instance.GongCopy().(*Staff_size)
		stage.Staff_sizes_reference[instance] = _copy
		stage.Staff_sizes_instance[_copy] = instance
		stage.Staff_sizes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Staff_tunings_reference = make(map[*Staff_tuning]*Staff_tuning)
	stage.Staff_tunings_referenceOrder = make(map[*Staff_tuning]uint) // diff Unstage needs the reference order
	stage.Staff_tunings_instance = make(map[*Staff_tuning]*Staff_tuning)
	for instance := range stage.Staff_tunings {
		_copy := instance.GongCopy().(*Staff_tuning)
		stage.Staff_tunings_reference[instance] = _copy
		stage.Staff_tunings_instance[_copy] = instance
		stage.Staff_tunings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Stems_reference = make(map[*Stem]*Stem)
	stage.Stems_referenceOrder = make(map[*Stem]uint) // diff Unstage needs the reference order
	stage.Stems_instance = make(map[*Stem]*Stem)
	for instance := range stage.Stems {
		_copy := instance.GongCopy().(*Stem)
		stage.Stems_reference[instance] = _copy
		stage.Stems_instance[_copy] = instance
		stage.Stems_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Sticks_reference = make(map[*Stick]*Stick)
	stage.Sticks_referenceOrder = make(map[*Stick]uint) // diff Unstage needs the reference order
	stage.Sticks_instance = make(map[*Stick]*Stick)
	for instance := range stage.Sticks {
		_copy := instance.GongCopy().(*Stick)
		stage.Sticks_reference[instance] = _copy
		stage.Sticks_instance[_copy] = instance
		stage.Sticks_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.String_mutes_reference = make(map[*String_mute]*String_mute)
	stage.String_mutes_referenceOrder = make(map[*String_mute]uint) // diff Unstage needs the reference order
	stage.String_mutes_instance = make(map[*String_mute]*String_mute)
	for instance := range stage.String_mutes {
		_copy := instance.GongCopy().(*String_mute)
		stage.String_mutes_reference[instance] = _copy
		stage.String_mutes_instance[_copy] = instance
		stage.String_mutes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.String_types_reference = make(map[*String_type]*String_type)
	stage.String_types_referenceOrder = make(map[*String_type]uint) // diff Unstage needs the reference order
	stage.String_types_instance = make(map[*String_type]*String_type)
	for instance := range stage.String_types {
		_copy := instance.GongCopy().(*String_type)
		stage.String_types_reference[instance] = _copy
		stage.String_types_instance[_copy] = instance
		stage.String_types_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Strong_accents_reference = make(map[*Strong_accent]*Strong_accent)
	stage.Strong_accents_referenceOrder = make(map[*Strong_accent]uint) // diff Unstage needs the reference order
	stage.Strong_accents_instance = make(map[*Strong_accent]*Strong_accent)
	for instance := range stage.Strong_accents {
		_copy := instance.GongCopy().(*Strong_accent)
		stage.Strong_accents_reference[instance] = _copy
		stage.Strong_accents_instance[_copy] = instance
		stage.Strong_accents_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Style_texts_reference = make(map[*Style_text]*Style_text)
	stage.Style_texts_referenceOrder = make(map[*Style_text]uint) // diff Unstage needs the reference order
	stage.Style_texts_instance = make(map[*Style_text]*Style_text)
	for instance := range stage.Style_texts {
		_copy := instance.GongCopy().(*Style_text)
		stage.Style_texts_reference[instance] = _copy
		stage.Style_texts_instance[_copy] = instance
		stage.Style_texts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Supportss_reference = make(map[*Supports]*Supports)
	stage.Supportss_referenceOrder = make(map[*Supports]uint) // diff Unstage needs the reference order
	stage.Supportss_instance = make(map[*Supports]*Supports)
	for instance := range stage.Supportss {
		_copy := instance.GongCopy().(*Supports)
		stage.Supportss_reference[instance] = _copy
		stage.Supportss_instance[_copy] = instance
		stage.Supportss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Swings_reference = make(map[*Swing]*Swing)
	stage.Swings_referenceOrder = make(map[*Swing]uint) // diff Unstage needs the reference order
	stage.Swings_instance = make(map[*Swing]*Swing)
	for instance := range stage.Swings {
		_copy := instance.GongCopy().(*Swing)
		stage.Swings_reference[instance] = _copy
		stage.Swings_instance[_copy] = instance
		stage.Swings_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Syncs_reference = make(map[*Sync]*Sync)
	stage.Syncs_referenceOrder = make(map[*Sync]uint) // diff Unstage needs the reference order
	stage.Syncs_instance = make(map[*Sync]*Sync)
	for instance := range stage.Syncs {
		_copy := instance.GongCopy().(*Sync)
		stage.Syncs_reference[instance] = _copy
		stage.Syncs_instance[_copy] = instance
		stage.Syncs_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.System_dividerss_reference = make(map[*System_dividers]*System_dividers)
	stage.System_dividerss_referenceOrder = make(map[*System_dividers]uint) // diff Unstage needs the reference order
	stage.System_dividerss_instance = make(map[*System_dividers]*System_dividers)
	for instance := range stage.System_dividerss {
		_copy := instance.GongCopy().(*System_dividers)
		stage.System_dividerss_reference[instance] = _copy
		stage.System_dividerss_instance[_copy] = instance
		stage.System_dividerss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.System_layouts_reference = make(map[*System_layout]*System_layout)
	stage.System_layouts_referenceOrder = make(map[*System_layout]uint) // diff Unstage needs the reference order
	stage.System_layouts_instance = make(map[*System_layout]*System_layout)
	for instance := range stage.System_layouts {
		_copy := instance.GongCopy().(*System_layout)
		stage.System_layouts_reference[instance] = _copy
		stage.System_layouts_instance[_copy] = instance
		stage.System_layouts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.System_marginss_reference = make(map[*System_margins]*System_margins)
	stage.System_marginss_referenceOrder = make(map[*System_margins]uint) // diff Unstage needs the reference order
	stage.System_marginss_instance = make(map[*System_margins]*System_margins)
	for instance := range stage.System_marginss {
		_copy := instance.GongCopy().(*System_margins)
		stage.System_marginss_reference[instance] = _copy
		stage.System_marginss_instance[_copy] = instance
		stage.System_marginss_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Taps_reference = make(map[*Tap]*Tap)
	stage.Taps_referenceOrder = make(map[*Tap]uint) // diff Unstage needs the reference order
	stage.Taps_instance = make(map[*Tap]*Tap)
	for instance := range stage.Taps {
		_copy := instance.GongCopy().(*Tap)
		stage.Taps_reference[instance] = _copy
		stage.Taps_instance[_copy] = instance
		stage.Taps_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Technicals_reference = make(map[*Technical]*Technical)
	stage.Technicals_referenceOrder = make(map[*Technical]uint) // diff Unstage needs the reference order
	stage.Technicals_instance = make(map[*Technical]*Technical)
	for instance := range stage.Technicals {
		_copy := instance.GongCopy().(*Technical)
		stage.Technicals_reference[instance] = _copy
		stage.Technicals_instance[_copy] = instance
		stage.Technicals_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Text_element_datas_reference = make(map[*Text_element_data]*Text_element_data)
	stage.Text_element_datas_referenceOrder = make(map[*Text_element_data]uint) // diff Unstage needs the reference order
	stage.Text_element_datas_instance = make(map[*Text_element_data]*Text_element_data)
	for instance := range stage.Text_element_datas {
		_copy := instance.GongCopy().(*Text_element_data)
		stage.Text_element_datas_reference[instance] = _copy
		stage.Text_element_datas_instance[_copy] = instance
		stage.Text_element_datas_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Ties_reference = make(map[*Tie]*Tie)
	stage.Ties_referenceOrder = make(map[*Tie]uint) // diff Unstage needs the reference order
	stage.Ties_instance = make(map[*Tie]*Tie)
	for instance := range stage.Ties {
		_copy := instance.GongCopy().(*Tie)
		stage.Ties_reference[instance] = _copy
		stage.Ties_instance[_copy] = instance
		stage.Ties_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Tieds_reference = make(map[*Tied]*Tied)
	stage.Tieds_referenceOrder = make(map[*Tied]uint) // diff Unstage needs the reference order
	stage.Tieds_instance = make(map[*Tied]*Tied)
	for instance := range stage.Tieds {
		_copy := instance.GongCopy().(*Tied)
		stage.Tieds_reference[instance] = _copy
		stage.Tieds_instance[_copy] = instance
		stage.Tieds_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Times_reference = make(map[*Time]*Time)
	stage.Times_referenceOrder = make(map[*Time]uint) // diff Unstage needs the reference order
	stage.Times_instance = make(map[*Time]*Time)
	for instance := range stage.Times {
		_copy := instance.GongCopy().(*Time)
		stage.Times_reference[instance] = _copy
		stage.Times_instance[_copy] = instance
		stage.Times_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Time_modifications_reference = make(map[*Time_modification]*Time_modification)
	stage.Time_modifications_referenceOrder = make(map[*Time_modification]uint) // diff Unstage needs the reference order
	stage.Time_modifications_instance = make(map[*Time_modification]*Time_modification)
	for instance := range stage.Time_modifications {
		_copy := instance.GongCopy().(*Time_modification)
		stage.Time_modifications_reference[instance] = _copy
		stage.Time_modifications_instance[_copy] = instance
		stage.Time_modifications_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Timpanis_reference = make(map[*Timpani]*Timpani)
	stage.Timpanis_referenceOrder = make(map[*Timpani]uint) // diff Unstage needs the reference order
	stage.Timpanis_instance = make(map[*Timpani]*Timpani)
	for instance := range stage.Timpanis {
		_copy := instance.GongCopy().(*Timpani)
		stage.Timpanis_reference[instance] = _copy
		stage.Timpanis_instance[_copy] = instance
		stage.Timpanis_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Transposes_reference = make(map[*Transpose]*Transpose)
	stage.Transposes_referenceOrder = make(map[*Transpose]uint) // diff Unstage needs the reference order
	stage.Transposes_instance = make(map[*Transpose]*Transpose)
	for instance := range stage.Transposes {
		_copy := instance.GongCopy().(*Transpose)
		stage.Transposes_reference[instance] = _copy
		stage.Transposes_instance[_copy] = instance
		stage.Transposes_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Tremolos_reference = make(map[*Tremolo]*Tremolo)
	stage.Tremolos_referenceOrder = make(map[*Tremolo]uint) // diff Unstage needs the reference order
	stage.Tremolos_instance = make(map[*Tremolo]*Tremolo)
	for instance := range stage.Tremolos {
		_copy := instance.GongCopy().(*Tremolo)
		stage.Tremolos_reference[instance] = _copy
		stage.Tremolos_instance[_copy] = instance
		stage.Tremolos_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Tuplets_reference = make(map[*Tuplet]*Tuplet)
	stage.Tuplets_referenceOrder = make(map[*Tuplet]uint) // diff Unstage needs the reference order
	stage.Tuplets_instance = make(map[*Tuplet]*Tuplet)
	for instance := range stage.Tuplets {
		_copy := instance.GongCopy().(*Tuplet)
		stage.Tuplets_reference[instance] = _copy
		stage.Tuplets_instance[_copy] = instance
		stage.Tuplets_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Tuplet_dots_reference = make(map[*Tuplet_dot]*Tuplet_dot)
	stage.Tuplet_dots_referenceOrder = make(map[*Tuplet_dot]uint) // diff Unstage needs the reference order
	stage.Tuplet_dots_instance = make(map[*Tuplet_dot]*Tuplet_dot)
	for instance := range stage.Tuplet_dots {
		_copy := instance.GongCopy().(*Tuplet_dot)
		stage.Tuplet_dots_reference[instance] = _copy
		stage.Tuplet_dots_instance[_copy] = instance
		stage.Tuplet_dots_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Tuplet_numbers_reference = make(map[*Tuplet_number]*Tuplet_number)
	stage.Tuplet_numbers_referenceOrder = make(map[*Tuplet_number]uint) // diff Unstage needs the reference order
	stage.Tuplet_numbers_instance = make(map[*Tuplet_number]*Tuplet_number)
	for instance := range stage.Tuplet_numbers {
		_copy := instance.GongCopy().(*Tuplet_number)
		stage.Tuplet_numbers_reference[instance] = _copy
		stage.Tuplet_numbers_instance[_copy] = instance
		stage.Tuplet_numbers_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Tuplet_portions_reference = make(map[*Tuplet_portion]*Tuplet_portion)
	stage.Tuplet_portions_referenceOrder = make(map[*Tuplet_portion]uint) // diff Unstage needs the reference order
	stage.Tuplet_portions_instance = make(map[*Tuplet_portion]*Tuplet_portion)
	for instance := range stage.Tuplet_portions {
		_copy := instance.GongCopy().(*Tuplet_portion)
		stage.Tuplet_portions_reference[instance] = _copy
		stage.Tuplet_portions_instance[_copy] = instance
		stage.Tuplet_portions_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Tuplet_types_reference = make(map[*Tuplet_type]*Tuplet_type)
	stage.Tuplet_types_referenceOrder = make(map[*Tuplet_type]uint) // diff Unstage needs the reference order
	stage.Tuplet_types_instance = make(map[*Tuplet_type]*Tuplet_type)
	for instance := range stage.Tuplet_types {
		_copy := instance.GongCopy().(*Tuplet_type)
		stage.Tuplet_types_reference[instance] = _copy
		stage.Tuplet_types_instance[_copy] = instance
		stage.Tuplet_types_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Typed_texts_reference = make(map[*Typed_text]*Typed_text)
	stage.Typed_texts_referenceOrder = make(map[*Typed_text]uint) // diff Unstage needs the reference order
	stage.Typed_texts_instance = make(map[*Typed_text]*Typed_text)
	for instance := range stage.Typed_texts {
		_copy := instance.GongCopy().(*Typed_text)
		stage.Typed_texts_reference[instance] = _copy
		stage.Typed_texts_instance[_copy] = instance
		stage.Typed_texts_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Unpitcheds_reference = make(map[*Unpitched]*Unpitched)
	stage.Unpitcheds_referenceOrder = make(map[*Unpitched]uint) // diff Unstage needs the reference order
	stage.Unpitcheds_instance = make(map[*Unpitched]*Unpitched)
	for instance := range stage.Unpitcheds {
		_copy := instance.GongCopy().(*Unpitched)
		stage.Unpitcheds_reference[instance] = _copy
		stage.Unpitcheds_instance[_copy] = instance
		stage.Unpitcheds_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Virtual_instruments_reference = make(map[*Virtual_instrument]*Virtual_instrument)
	stage.Virtual_instruments_referenceOrder = make(map[*Virtual_instrument]uint) // diff Unstage needs the reference order
	stage.Virtual_instruments_instance = make(map[*Virtual_instrument]*Virtual_instrument)
	for instance := range stage.Virtual_instruments {
		_copy := instance.GongCopy().(*Virtual_instrument)
		stage.Virtual_instruments_reference[instance] = _copy
		stage.Virtual_instruments_instance[_copy] = instance
		stage.Virtual_instruments_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Waits_reference = make(map[*Wait]*Wait)
	stage.Waits_referenceOrder = make(map[*Wait]uint) // diff Unstage needs the reference order
	stage.Waits_instance = make(map[*Wait]*Wait)
	for instance := range stage.Waits {
		_copy := instance.GongCopy().(*Wait)
		stage.Waits_reference[instance] = _copy
		stage.Waits_instance[_copy] = instance
		stage.Waits_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Wavy_lines_reference = make(map[*Wavy_line]*Wavy_line)
	stage.Wavy_lines_referenceOrder = make(map[*Wavy_line]uint) // diff Unstage needs the reference order
	stage.Wavy_lines_instance = make(map[*Wavy_line]*Wavy_line)
	for instance := range stage.Wavy_lines {
		_copy := instance.GongCopy().(*Wavy_line)
		stage.Wavy_lines_reference[instance] = _copy
		stage.Wavy_lines_instance[_copy] = instance
		stage.Wavy_lines_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Wedges_reference = make(map[*Wedge]*Wedge)
	stage.Wedges_referenceOrder = make(map[*Wedge]uint) // diff Unstage needs the reference order
	stage.Wedges_instance = make(map[*Wedge]*Wedge)
	for instance := range stage.Wedges {
		_copy := instance.GongCopy().(*Wedge)
		stage.Wedges_reference[instance] = _copy
		stage.Wedges_instance[_copy] = instance
		stage.Wedges_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Woods_reference = make(map[*Wood]*Wood)
	stage.Woods_referenceOrder = make(map[*Wood]uint) // diff Unstage needs the reference order
	stage.Woods_instance = make(map[*Wood]*Wood)
	for instance := range stage.Woods {
		_copy := instance.GongCopy().(*Wood)
		stage.Woods_reference[instance] = _copy
		stage.Woods_instance[_copy] = instance
		stage.Woods_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	stage.Works_reference = make(map[*Work]*Work)
	stage.Works_referenceOrder = make(map[*Work]uint) // diff Unstage needs the reference order
	stage.Works_instance = make(map[*Work]*Work)
	for instance := range stage.Works {
		_copy := instance.GongCopy().(*Work)
		stage.Works_reference[instance] = _copy
		stage.Works_instance[_copy] = instance
		stage.Works_referenceOrder[_copy] = instance.GongGetOrder(stage)
	}

	// insertion point per named struct
	for instance := range stage.A_directives {
		reference := stage.A_directives_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_measures {
		reference := stage.A_measures_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_measure_1s {
		reference := stage.A_measure_1s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_parts {
		reference := stage.A_parts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.A_part_1s {
		reference := stage.A_part_1s_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Accidentals {
		reference := stage.Accidentals_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Accidental_marks {
		reference := stage.Accidental_marks_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Accidental_texts {
		reference := stage.Accidental_texts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Accords {
		reference := stage.Accords_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Accordion_registrations {
		reference := stage.Accordion_registrations_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Appearances {
		reference := stage.Appearances_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Arpeggiates {
		reference := stage.Arpeggiates_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Arrows {
		reference := stage.Arrows_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Articulationss {
		reference := stage.Articulationss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Assesss {
		reference := stage.Assesss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Attributess {
		reference := stage.Attributess_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Backups {
		reference := stage.Backups_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Bar_style_colors {
		reference := stage.Bar_style_colors_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Barlines {
		reference := stage.Barlines_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Barres {
		reference := stage.Barres_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Basss {
		reference := stage.Basss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Bass_steps {
		reference := stage.Bass_steps_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Beams {
		reference := stage.Beams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Beat_repeats {
		reference := stage.Beat_repeats_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Beat_unit_tieds {
		reference := stage.Beat_unit_tieds_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Beaters {
		reference := stage.Beaters_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Bends {
		reference := stage.Bends_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Bookmarks {
		reference := stage.Bookmarks_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Brackets {
		reference := stage.Brackets_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Breath_marks {
		reference := stage.Breath_marks_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Caesuras {
		reference := stage.Caesuras_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Cancels {
		reference := stage.Cancels_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Clefs {
		reference := stage.Clefs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Codas {
		reference := stage.Codas_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Credits {
		reference := stage.Credits_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Dashess {
		reference := stage.Dashess_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Defaultss {
		reference := stage.Defaultss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Degrees {
		reference := stage.Degrees_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Degree_alters {
		reference := stage.Degree_alters_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Degree_types {
		reference := stage.Degree_types_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Degree_values {
		reference := stage.Degree_values_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Directions {
		reference := stage.Directions_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Direction_types {
		reference := stage.Direction_types_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Distances {
		reference := stage.Distances_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Doubles {
		reference := stage.Doubles_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Dynamicss {
		reference := stage.Dynamicss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Effects {
		reference := stage.Effects_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Elisions {
		reference := stage.Elisions_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Emptys {
		reference := stage.Emptys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Empty_fonts {
		reference := stage.Empty_fonts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Empty_lines {
		reference := stage.Empty_lines_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Empty_placements {
		reference := stage.Empty_placements_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Empty_placement_smufls {
		reference := stage.Empty_placement_smufls_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Empty_print_object_style_aligns {
		reference := stage.Empty_print_object_style_aligns_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Empty_print_styles {
		reference := stage.Empty_print_styles_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Empty_print_style_aligns {
		reference := stage.Empty_print_style_aligns_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Empty_print_style_align_ids {
		reference := stage.Empty_print_style_align_ids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Empty_trill_sounds {
		reference := stage.Empty_trill_sounds_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Encodings {
		reference := stage.Encodings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Endings {
		reference := stage.Endings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Extends {
		reference := stage.Extends_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Features {
		reference := stage.Features_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Fermatas {
		reference := stage.Fermatas_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Figures {
		reference := stage.Figures_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Figured_basss {
		reference := stage.Figured_basss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Fingerings {
		reference := stage.Fingerings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.First_frets {
		reference := stage.First_frets_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.For_parts {
		reference := stage.For_parts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Formatted_symbols {
		reference := stage.Formatted_symbols_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Formatted_symbol_ids {
		reference := stage.Formatted_symbol_ids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Formatted_texts {
		reference := stage.Formatted_texts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Formatted_text_ids {
		reference := stage.Formatted_text_ids_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Forwards {
		reference := stage.Forwards_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Frames {
		reference := stage.Frames_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Frame_notes {
		reference := stage.Frame_notes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Frets {
		reference := stage.Frets_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Glasss {
		reference := stage.Glasss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Glissandos {
		reference := stage.Glissandos_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Glyphs {
		reference := stage.Glyphs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Graces {
		reference := stage.Graces_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Group_barlines {
		reference := stage.Group_barlines_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Group_names {
		reference := stage.Group_names_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Group_symbols {
		reference := stage.Group_symbols_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Groupings {
		reference := stage.Groupings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Hammer_on_pull_offs {
		reference := stage.Hammer_on_pull_offs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Handbells {
		reference := stage.Handbells_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Harmon_closeds {
		reference := stage.Harmon_closeds_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Harmon_mutes {
		reference := stage.Harmon_mutes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Harmonics {
		reference := stage.Harmonics_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Harmonys {
		reference := stage.Harmonys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Harmony_alters {
		reference := stage.Harmony_alters_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Harp_pedalss {
		reference := stage.Harp_pedalss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Heel_toes {
		reference := stage.Heel_toes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Holes {
		reference := stage.Holes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Hole_closeds {
		reference := stage.Hole_closeds_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Horizontal_turns {
		reference := stage.Horizontal_turns_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Identifications {
		reference := stage.Identifications_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Images {
		reference := stage.Images_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Instruments {
		reference := stage.Instruments_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Instrument_changes {
		reference := stage.Instrument_changes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Instrument_links {
		reference := stage.Instrument_links_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Interchangeables {
		reference := stage.Interchangeables_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Inversions {
		reference := stage.Inversions_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Keys {
		reference := stage.Keys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Key_accidentals {
		reference := stage.Key_accidentals_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Key_octaves {
		reference := stage.Key_octaves_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Kinds {
		reference := stage.Kinds_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Levels {
		reference := stage.Levels_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Line_details {
		reference := stage.Line_details_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Line_widths {
		reference := stage.Line_widths_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Links {
		reference := stage.Links_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Listens {
		reference := stage.Listens_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Listenings {
		reference := stage.Listenings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Lyrics {
		reference := stage.Lyrics_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Lyric_fonts {
		reference := stage.Lyric_fonts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Lyric_languages {
		reference := stage.Lyric_languages_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Measure_layouts {
		reference := stage.Measure_layouts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Measure_numberings {
		reference := stage.Measure_numberings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Measure_repeats {
		reference := stage.Measure_repeats_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Measure_styles {
		reference := stage.Measure_styles_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Membranes {
		reference := stage.Membranes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Metals {
		reference := stage.Metals_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Metronomes {
		reference := stage.Metronomes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Metronome_beams {
		reference := stage.Metronome_beams_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Metronome_notes {
		reference := stage.Metronome_notes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Metronome_tieds {
		reference := stage.Metronome_tieds_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Metronome_tuplets {
		reference := stage.Metronome_tuplets_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Midi_devices {
		reference := stage.Midi_devices_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Midi_instruments {
		reference := stage.Midi_instruments_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Miscellaneouss {
		reference := stage.Miscellaneouss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Miscellaneous_fields {
		reference := stage.Miscellaneous_fields_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Mordents {
		reference := stage.Mordents_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Multiple_rests {
		reference := stage.Multiple_rests_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Name_displays {
		reference := stage.Name_displays_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Non_arpeggiates {
		reference := stage.Non_arpeggiates_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Notationss {
		reference := stage.Notationss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Notes {
		reference := stage.Notes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Note_sizes {
		reference := stage.Note_sizes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Note_types {
		reference := stage.Note_types_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Noteheads {
		reference := stage.Noteheads_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Notehead_texts {
		reference := stage.Notehead_texts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Numerals {
		reference := stage.Numerals_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Numeral_keys {
		reference := stage.Numeral_keys_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Numeral_roots {
		reference := stage.Numeral_roots_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Octave_shifts {
		reference := stage.Octave_shifts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Offsets {
		reference := stage.Offsets_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Opuss {
		reference := stage.Opuss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Ornamentss {
		reference := stage.Ornamentss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Other_appearances {
		reference := stage.Other_appearances_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Other_directions {
		reference := stage.Other_directions_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Other_listenings {
		reference := stage.Other_listenings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Other_notations {
		reference := stage.Other_notations_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Other_placement_texts {
		reference := stage.Other_placement_texts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Other_plays {
		reference := stage.Other_plays_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Other_texts {
		reference := stage.Other_texts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Page_layouts {
		reference := stage.Page_layouts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Page_marginss {
		reference := stage.Page_marginss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Part_clefs {
		reference := stage.Part_clefs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Part_groups {
		reference := stage.Part_groups_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Part_links {
		reference := stage.Part_links_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Part_lists {
		reference := stage.Part_lists_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Part_names {
		reference := stage.Part_names_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Part_symbols {
		reference := stage.Part_symbols_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Part_transposes {
		reference := stage.Part_transposes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Pedals {
		reference := stage.Pedals_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Pedal_tunings {
		reference := stage.Pedal_tunings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Per_minutes {
		reference := stage.Per_minutes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Percussions {
		reference := stage.Percussions_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Pitchs {
		reference := stage.Pitchs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Pitcheds {
		reference := stage.Pitcheds_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Placement_texts {
		reference := stage.Placement_texts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Plays {
		reference := stage.Plays_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Players {
		reference := stage.Players_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Principal_voices {
		reference := stage.Principal_voices_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Prints {
		reference := stage.Prints_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Releases {
		reference := stage.Releases_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Repeats {
		reference := stage.Repeats_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Rests {
		reference := stage.Rests_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Roots {
		reference := stage.Roots_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Root_steps {
		reference := stage.Root_steps_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Scalings {
		reference := stage.Scalings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Scordaturas {
		reference := stage.Scordaturas_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Score_instruments {
		reference := stage.Score_instruments_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Score_parts {
		reference := stage.Score_parts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Score_partwises {
		reference := stage.Score_partwises_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Score_timewises {
		reference := stage.Score_timewises_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Segnos {
		reference := stage.Segnos_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Slashs {
		reference := stage.Slashs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Slides {
		reference := stage.Slides_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Slurs {
		reference := stage.Slurs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Sounds {
		reference := stage.Sounds_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Staff_detailss {
		reference := stage.Staff_detailss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Staff_divides {
		reference := stage.Staff_divides_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Staff_layouts {
		reference := stage.Staff_layouts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Staff_sizes {
		reference := stage.Staff_sizes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Staff_tunings {
		reference := stage.Staff_tunings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Stems {
		reference := stage.Stems_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Sticks {
		reference := stage.Sticks_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.String_mutes {
		reference := stage.String_mutes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.String_types {
		reference := stage.String_types_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Strong_accents {
		reference := stage.Strong_accents_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Style_texts {
		reference := stage.Style_texts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Supportss {
		reference := stage.Supportss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Swings {
		reference := stage.Swings_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Syncs {
		reference := stage.Syncs_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.System_dividerss {
		reference := stage.System_dividerss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.System_layouts {
		reference := stage.System_layouts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.System_marginss {
		reference := stage.System_marginss_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Taps {
		reference := stage.Taps_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Technicals {
		reference := stage.Technicals_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Text_element_datas {
		reference := stage.Text_element_datas_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Ties {
		reference := stage.Ties_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Tieds {
		reference := stage.Tieds_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Times {
		reference := stage.Times_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Time_modifications {
		reference := stage.Time_modifications_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Timpanis {
		reference := stage.Timpanis_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Transposes {
		reference := stage.Transposes_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Tremolos {
		reference := stage.Tremolos_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Tuplets {
		reference := stage.Tuplets_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Tuplet_dots {
		reference := stage.Tuplet_dots_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Tuplet_numbers {
		reference := stage.Tuplet_numbers_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Tuplet_portions {
		reference := stage.Tuplet_portions_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Tuplet_types {
		reference := stage.Tuplet_types_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Typed_texts {
		reference := stage.Typed_texts_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Unpitcheds {
		reference := stage.Unpitcheds_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Virtual_instruments {
		reference := stage.Virtual_instruments_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Waits {
		reference := stage.Waits_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Wavy_lines {
		reference := stage.Wavy_lines_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Wedges {
		reference := stage.Wedges_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Woods {
		reference := stage.Woods_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	for instance := range stage.Works {
		reference := stage.Works_reference[instance]
		reference.GongReconstructPointersFromReferences(stage, instance)
	}

	stage.recomputeOrders()
}

// GongGetOrder returns the order of the instance in the staging area
// This order is set at staging time, and reflects the order of creation of the instances
// in the staging area
// It is used when rendering slices of GongstructIF to keep a deterministic order
// which is important for frontends such as web frontends
// to avoid unnecessary re-renderings
// insertion point per named struct
func (a_directive *A_directive) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_directive_stagedOrder[a_directive]; ok {
		return order
	}
	if order, ok := stage.A_directives_referenceOrder[a_directive]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_directive was not staged and does not have a reference order", a_directive)
		return 0
	}
}

func (a_measure *A_measure) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_measure_stagedOrder[a_measure]; ok {
		return order
	}
	if order, ok := stage.A_measures_referenceOrder[a_measure]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_measure was not staged and does not have a reference order", a_measure)
		return 0
	}
}

func (a_measure_1 *A_measure_1) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_measure_1_stagedOrder[a_measure_1]; ok {
		return order
	}
	if order, ok := stage.A_measure_1s_referenceOrder[a_measure_1]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_measure_1 was not staged and does not have a reference order", a_measure_1)
		return 0
	}
}

func (a_part *A_part) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_part_stagedOrder[a_part]; ok {
		return order
	}
	if order, ok := stage.A_parts_referenceOrder[a_part]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_part was not staged and does not have a reference order", a_part)
		return 0
	}
}

func (a_part_1 *A_part_1) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.A_part_1_stagedOrder[a_part_1]; ok {
		return order
	}
	if order, ok := stage.A_part_1s_referenceOrder[a_part_1]; ok {
		return order
	} else {
		log.Printf("instance %p of type A_part_1 was not staged and does not have a reference order", a_part_1)
		return 0
	}
}

func (accidental *Accidental) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Accidental_stagedOrder[accidental]; ok {
		return order
	}
	if order, ok := stage.Accidentals_referenceOrder[accidental]; ok {
		return order
	} else {
		log.Printf("instance %p of type Accidental was not staged and does not have a reference order", accidental)
		return 0
	}
}

func (accidental_mark *Accidental_mark) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Accidental_mark_stagedOrder[accidental_mark]; ok {
		return order
	}
	if order, ok := stage.Accidental_marks_referenceOrder[accidental_mark]; ok {
		return order
	} else {
		log.Printf("instance %p of type Accidental_mark was not staged and does not have a reference order", accidental_mark)
		return 0
	}
}

func (accidental_text *Accidental_text) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Accidental_text_stagedOrder[accidental_text]; ok {
		return order
	}
	if order, ok := stage.Accidental_texts_referenceOrder[accidental_text]; ok {
		return order
	} else {
		log.Printf("instance %p of type Accidental_text was not staged and does not have a reference order", accidental_text)
		return 0
	}
}

func (accord *Accord) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Accord_stagedOrder[accord]; ok {
		return order
	}
	if order, ok := stage.Accords_referenceOrder[accord]; ok {
		return order
	} else {
		log.Printf("instance %p of type Accord was not staged and does not have a reference order", accord)
		return 0
	}
}

func (accordion_registration *Accordion_registration) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Accordion_registration_stagedOrder[accordion_registration]; ok {
		return order
	}
	if order, ok := stage.Accordion_registrations_referenceOrder[accordion_registration]; ok {
		return order
	} else {
		log.Printf("instance %p of type Accordion_registration was not staged and does not have a reference order", accordion_registration)
		return 0
	}
}

func (appearance *Appearance) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Appearance_stagedOrder[appearance]; ok {
		return order
	}
	if order, ok := stage.Appearances_referenceOrder[appearance]; ok {
		return order
	} else {
		log.Printf("instance %p of type Appearance was not staged and does not have a reference order", appearance)
		return 0
	}
}

func (arpeggiate *Arpeggiate) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Arpeggiate_stagedOrder[arpeggiate]; ok {
		return order
	}
	if order, ok := stage.Arpeggiates_referenceOrder[arpeggiate]; ok {
		return order
	} else {
		log.Printf("instance %p of type Arpeggiate was not staged and does not have a reference order", arpeggiate)
		return 0
	}
}

func (arrow *Arrow) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Arrow_stagedOrder[arrow]; ok {
		return order
	}
	if order, ok := stage.Arrows_referenceOrder[arrow]; ok {
		return order
	} else {
		log.Printf("instance %p of type Arrow was not staged and does not have a reference order", arrow)
		return 0
	}
}

func (articulations *Articulations) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Articulations_stagedOrder[articulations]; ok {
		return order
	}
	if order, ok := stage.Articulationss_referenceOrder[articulations]; ok {
		return order
	} else {
		log.Printf("instance %p of type Articulations was not staged and does not have a reference order", articulations)
		return 0
	}
}

func (assess *Assess) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Assess_stagedOrder[assess]; ok {
		return order
	}
	if order, ok := stage.Assesss_referenceOrder[assess]; ok {
		return order
	} else {
		log.Printf("instance %p of type Assess was not staged and does not have a reference order", assess)
		return 0
	}
}

func (attributes *Attributes) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Attributes_stagedOrder[attributes]; ok {
		return order
	}
	if order, ok := stage.Attributess_referenceOrder[attributes]; ok {
		return order
	} else {
		log.Printf("instance %p of type Attributes was not staged and does not have a reference order", attributes)
		return 0
	}
}

func (backup *Backup) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Backup_stagedOrder[backup]; ok {
		return order
	}
	if order, ok := stage.Backups_referenceOrder[backup]; ok {
		return order
	} else {
		log.Printf("instance %p of type Backup was not staged and does not have a reference order", backup)
		return 0
	}
}

func (bar_style_color *Bar_style_color) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Bar_style_color_stagedOrder[bar_style_color]; ok {
		return order
	}
	if order, ok := stage.Bar_style_colors_referenceOrder[bar_style_color]; ok {
		return order
	} else {
		log.Printf("instance %p of type Bar_style_color was not staged and does not have a reference order", bar_style_color)
		return 0
	}
}

func (barline *Barline) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Barline_stagedOrder[barline]; ok {
		return order
	}
	if order, ok := stage.Barlines_referenceOrder[barline]; ok {
		return order
	} else {
		log.Printf("instance %p of type Barline was not staged and does not have a reference order", barline)
		return 0
	}
}

func (barre *Barre) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Barre_stagedOrder[barre]; ok {
		return order
	}
	if order, ok := stage.Barres_referenceOrder[barre]; ok {
		return order
	} else {
		log.Printf("instance %p of type Barre was not staged and does not have a reference order", barre)
		return 0
	}
}

func (bass *Bass) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Bass_stagedOrder[bass]; ok {
		return order
	}
	if order, ok := stage.Basss_referenceOrder[bass]; ok {
		return order
	} else {
		log.Printf("instance %p of type Bass was not staged and does not have a reference order", bass)
		return 0
	}
}

func (bass_step *Bass_step) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Bass_step_stagedOrder[bass_step]; ok {
		return order
	}
	if order, ok := stage.Bass_steps_referenceOrder[bass_step]; ok {
		return order
	} else {
		log.Printf("instance %p of type Bass_step was not staged and does not have a reference order", bass_step)
		return 0
	}
}

func (beam *Beam) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Beam_stagedOrder[beam]; ok {
		return order
	}
	if order, ok := stage.Beams_referenceOrder[beam]; ok {
		return order
	} else {
		log.Printf("instance %p of type Beam was not staged and does not have a reference order", beam)
		return 0
	}
}

func (beat_repeat *Beat_repeat) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Beat_repeat_stagedOrder[beat_repeat]; ok {
		return order
	}
	if order, ok := stage.Beat_repeats_referenceOrder[beat_repeat]; ok {
		return order
	} else {
		log.Printf("instance %p of type Beat_repeat was not staged and does not have a reference order", beat_repeat)
		return 0
	}
}

func (beat_unit_tied *Beat_unit_tied) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Beat_unit_tied_stagedOrder[beat_unit_tied]; ok {
		return order
	}
	if order, ok := stage.Beat_unit_tieds_referenceOrder[beat_unit_tied]; ok {
		return order
	} else {
		log.Printf("instance %p of type Beat_unit_tied was not staged and does not have a reference order", beat_unit_tied)
		return 0
	}
}

func (beater *Beater) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Beater_stagedOrder[beater]; ok {
		return order
	}
	if order, ok := stage.Beaters_referenceOrder[beater]; ok {
		return order
	} else {
		log.Printf("instance %p of type Beater was not staged and does not have a reference order", beater)
		return 0
	}
}

func (bend *Bend) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Bend_stagedOrder[bend]; ok {
		return order
	}
	if order, ok := stage.Bends_referenceOrder[bend]; ok {
		return order
	} else {
		log.Printf("instance %p of type Bend was not staged and does not have a reference order", bend)
		return 0
	}
}

func (bookmark *Bookmark) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Bookmark_stagedOrder[bookmark]; ok {
		return order
	}
	if order, ok := stage.Bookmarks_referenceOrder[bookmark]; ok {
		return order
	} else {
		log.Printf("instance %p of type Bookmark was not staged and does not have a reference order", bookmark)
		return 0
	}
}

func (bracket *Bracket) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Bracket_stagedOrder[bracket]; ok {
		return order
	}
	if order, ok := stage.Brackets_referenceOrder[bracket]; ok {
		return order
	} else {
		log.Printf("instance %p of type Bracket was not staged and does not have a reference order", bracket)
		return 0
	}
}

func (breath_mark *Breath_mark) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Breath_mark_stagedOrder[breath_mark]; ok {
		return order
	}
	if order, ok := stage.Breath_marks_referenceOrder[breath_mark]; ok {
		return order
	} else {
		log.Printf("instance %p of type Breath_mark was not staged and does not have a reference order", breath_mark)
		return 0
	}
}

func (caesura *Caesura) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Caesura_stagedOrder[caesura]; ok {
		return order
	}
	if order, ok := stage.Caesuras_referenceOrder[caesura]; ok {
		return order
	} else {
		log.Printf("instance %p of type Caesura was not staged and does not have a reference order", caesura)
		return 0
	}
}

func (cancel *Cancel) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Cancel_stagedOrder[cancel]; ok {
		return order
	}
	if order, ok := stage.Cancels_referenceOrder[cancel]; ok {
		return order
	} else {
		log.Printf("instance %p of type Cancel was not staged and does not have a reference order", cancel)
		return 0
	}
}

func (clef *Clef) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Clef_stagedOrder[clef]; ok {
		return order
	}
	if order, ok := stage.Clefs_referenceOrder[clef]; ok {
		return order
	} else {
		log.Printf("instance %p of type Clef was not staged and does not have a reference order", clef)
		return 0
	}
}

func (coda *Coda) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Coda_stagedOrder[coda]; ok {
		return order
	}
	if order, ok := stage.Codas_referenceOrder[coda]; ok {
		return order
	} else {
		log.Printf("instance %p of type Coda was not staged and does not have a reference order", coda)
		return 0
	}
}

func (credit *Credit) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Credit_stagedOrder[credit]; ok {
		return order
	}
	if order, ok := stage.Credits_referenceOrder[credit]; ok {
		return order
	} else {
		log.Printf("instance %p of type Credit was not staged and does not have a reference order", credit)
		return 0
	}
}

func (dashes *Dashes) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Dashes_stagedOrder[dashes]; ok {
		return order
	}
	if order, ok := stage.Dashess_referenceOrder[dashes]; ok {
		return order
	} else {
		log.Printf("instance %p of type Dashes was not staged and does not have a reference order", dashes)
		return 0
	}
}

func (defaults *Defaults) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Defaults_stagedOrder[defaults]; ok {
		return order
	}
	if order, ok := stage.Defaultss_referenceOrder[defaults]; ok {
		return order
	} else {
		log.Printf("instance %p of type Defaults was not staged and does not have a reference order", defaults)
		return 0
	}
}

func (degree *Degree) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Degree_stagedOrder[degree]; ok {
		return order
	}
	if order, ok := stage.Degrees_referenceOrder[degree]; ok {
		return order
	} else {
		log.Printf("instance %p of type Degree was not staged and does not have a reference order", degree)
		return 0
	}
}

func (degree_alter *Degree_alter) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Degree_alter_stagedOrder[degree_alter]; ok {
		return order
	}
	if order, ok := stage.Degree_alters_referenceOrder[degree_alter]; ok {
		return order
	} else {
		log.Printf("instance %p of type Degree_alter was not staged and does not have a reference order", degree_alter)
		return 0
	}
}

func (degree_type *Degree_type) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Degree_type_stagedOrder[degree_type]; ok {
		return order
	}
	if order, ok := stage.Degree_types_referenceOrder[degree_type]; ok {
		return order
	} else {
		log.Printf("instance %p of type Degree_type was not staged and does not have a reference order", degree_type)
		return 0
	}
}

func (degree_value *Degree_value) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Degree_value_stagedOrder[degree_value]; ok {
		return order
	}
	if order, ok := stage.Degree_values_referenceOrder[degree_value]; ok {
		return order
	} else {
		log.Printf("instance %p of type Degree_value was not staged and does not have a reference order", degree_value)
		return 0
	}
}

func (direction *Direction) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Direction_stagedOrder[direction]; ok {
		return order
	}
	if order, ok := stage.Directions_referenceOrder[direction]; ok {
		return order
	} else {
		log.Printf("instance %p of type Direction was not staged and does not have a reference order", direction)
		return 0
	}
}

func (direction_type *Direction_type) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Direction_type_stagedOrder[direction_type]; ok {
		return order
	}
	if order, ok := stage.Direction_types_referenceOrder[direction_type]; ok {
		return order
	} else {
		log.Printf("instance %p of type Direction_type was not staged and does not have a reference order", direction_type)
		return 0
	}
}

func (distance *Distance) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Distance_stagedOrder[distance]; ok {
		return order
	}
	if order, ok := stage.Distances_referenceOrder[distance]; ok {
		return order
	} else {
		log.Printf("instance %p of type Distance was not staged and does not have a reference order", distance)
		return 0
	}
}

func (double *Double) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Double_stagedOrder[double]; ok {
		return order
	}
	if order, ok := stage.Doubles_referenceOrder[double]; ok {
		return order
	} else {
		log.Printf("instance %p of type Double was not staged and does not have a reference order", double)
		return 0
	}
}

func (dynamics *Dynamics) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Dynamics_stagedOrder[dynamics]; ok {
		return order
	}
	if order, ok := stage.Dynamicss_referenceOrder[dynamics]; ok {
		return order
	} else {
		log.Printf("instance %p of type Dynamics was not staged and does not have a reference order", dynamics)
		return 0
	}
}

func (effect *Effect) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Effect_stagedOrder[effect]; ok {
		return order
	}
	if order, ok := stage.Effects_referenceOrder[effect]; ok {
		return order
	} else {
		log.Printf("instance %p of type Effect was not staged and does not have a reference order", effect)
		return 0
	}
}

func (elision *Elision) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Elision_stagedOrder[elision]; ok {
		return order
	}
	if order, ok := stage.Elisions_referenceOrder[elision]; ok {
		return order
	} else {
		log.Printf("instance %p of type Elision was not staged and does not have a reference order", elision)
		return 0
	}
}

func (empty *Empty) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Empty_stagedOrder[empty]; ok {
		return order
	}
	if order, ok := stage.Emptys_referenceOrder[empty]; ok {
		return order
	} else {
		log.Printf("instance %p of type Empty was not staged and does not have a reference order", empty)
		return 0
	}
}

func (empty_font *Empty_font) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Empty_font_stagedOrder[empty_font]; ok {
		return order
	}
	if order, ok := stage.Empty_fonts_referenceOrder[empty_font]; ok {
		return order
	} else {
		log.Printf("instance %p of type Empty_font was not staged and does not have a reference order", empty_font)
		return 0
	}
}

func (empty_line *Empty_line) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Empty_line_stagedOrder[empty_line]; ok {
		return order
	}
	if order, ok := stage.Empty_lines_referenceOrder[empty_line]; ok {
		return order
	} else {
		log.Printf("instance %p of type Empty_line was not staged and does not have a reference order", empty_line)
		return 0
	}
}

func (empty_placement *Empty_placement) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Empty_placement_stagedOrder[empty_placement]; ok {
		return order
	}
	if order, ok := stage.Empty_placements_referenceOrder[empty_placement]; ok {
		return order
	} else {
		log.Printf("instance %p of type Empty_placement was not staged and does not have a reference order", empty_placement)
		return 0
	}
}

func (empty_placement_smufl *Empty_placement_smufl) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Empty_placement_smufl_stagedOrder[empty_placement_smufl]; ok {
		return order
	}
	if order, ok := stage.Empty_placement_smufls_referenceOrder[empty_placement_smufl]; ok {
		return order
	} else {
		log.Printf("instance %p of type Empty_placement_smufl was not staged and does not have a reference order", empty_placement_smufl)
		return 0
	}
}

func (empty_print_object_style_align *Empty_print_object_style_align) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Empty_print_object_style_align_stagedOrder[empty_print_object_style_align]; ok {
		return order
	}
	if order, ok := stage.Empty_print_object_style_aligns_referenceOrder[empty_print_object_style_align]; ok {
		return order
	} else {
		log.Printf("instance %p of type Empty_print_object_style_align was not staged and does not have a reference order", empty_print_object_style_align)
		return 0
	}
}

func (empty_print_style *Empty_print_style) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Empty_print_style_stagedOrder[empty_print_style]; ok {
		return order
	}
	if order, ok := stage.Empty_print_styles_referenceOrder[empty_print_style]; ok {
		return order
	} else {
		log.Printf("instance %p of type Empty_print_style was not staged and does not have a reference order", empty_print_style)
		return 0
	}
}

func (empty_print_style_align *Empty_print_style_align) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Empty_print_style_align_stagedOrder[empty_print_style_align]; ok {
		return order
	}
	if order, ok := stage.Empty_print_style_aligns_referenceOrder[empty_print_style_align]; ok {
		return order
	} else {
		log.Printf("instance %p of type Empty_print_style_align was not staged and does not have a reference order", empty_print_style_align)
		return 0
	}
}

func (empty_print_style_align_id *Empty_print_style_align_id) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Empty_print_style_align_id_stagedOrder[empty_print_style_align_id]; ok {
		return order
	}
	if order, ok := stage.Empty_print_style_align_ids_referenceOrder[empty_print_style_align_id]; ok {
		return order
	} else {
		log.Printf("instance %p of type Empty_print_style_align_id was not staged and does not have a reference order", empty_print_style_align_id)
		return 0
	}
}

func (empty_trill_sound *Empty_trill_sound) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Empty_trill_sound_stagedOrder[empty_trill_sound]; ok {
		return order
	}
	if order, ok := stage.Empty_trill_sounds_referenceOrder[empty_trill_sound]; ok {
		return order
	} else {
		log.Printf("instance %p of type Empty_trill_sound was not staged and does not have a reference order", empty_trill_sound)
		return 0
	}
}

func (encoding *Encoding) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Encoding_stagedOrder[encoding]; ok {
		return order
	}
	if order, ok := stage.Encodings_referenceOrder[encoding]; ok {
		return order
	} else {
		log.Printf("instance %p of type Encoding was not staged and does not have a reference order", encoding)
		return 0
	}
}

func (ending *Ending) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Ending_stagedOrder[ending]; ok {
		return order
	}
	if order, ok := stage.Endings_referenceOrder[ending]; ok {
		return order
	} else {
		log.Printf("instance %p of type Ending was not staged and does not have a reference order", ending)
		return 0
	}
}

func (extend *Extend) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Extend_stagedOrder[extend]; ok {
		return order
	}
	if order, ok := stage.Extends_referenceOrder[extend]; ok {
		return order
	} else {
		log.Printf("instance %p of type Extend was not staged and does not have a reference order", extend)
		return 0
	}
}

func (feature *Feature) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Feature_stagedOrder[feature]; ok {
		return order
	}
	if order, ok := stage.Features_referenceOrder[feature]; ok {
		return order
	} else {
		log.Printf("instance %p of type Feature was not staged and does not have a reference order", feature)
		return 0
	}
}

func (fermata *Fermata) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Fermata_stagedOrder[fermata]; ok {
		return order
	}
	if order, ok := stage.Fermatas_referenceOrder[fermata]; ok {
		return order
	} else {
		log.Printf("instance %p of type Fermata was not staged and does not have a reference order", fermata)
		return 0
	}
}

func (figure *Figure) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Figure_stagedOrder[figure]; ok {
		return order
	}
	if order, ok := stage.Figures_referenceOrder[figure]; ok {
		return order
	} else {
		log.Printf("instance %p of type Figure was not staged and does not have a reference order", figure)
		return 0
	}
}

func (figured_bass *Figured_bass) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Figured_bass_stagedOrder[figured_bass]; ok {
		return order
	}
	if order, ok := stage.Figured_basss_referenceOrder[figured_bass]; ok {
		return order
	} else {
		log.Printf("instance %p of type Figured_bass was not staged and does not have a reference order", figured_bass)
		return 0
	}
}

func (fingering *Fingering) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Fingering_stagedOrder[fingering]; ok {
		return order
	}
	if order, ok := stage.Fingerings_referenceOrder[fingering]; ok {
		return order
	} else {
		log.Printf("instance %p of type Fingering was not staged and does not have a reference order", fingering)
		return 0
	}
}

func (first_fret *First_fret) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.First_fret_stagedOrder[first_fret]; ok {
		return order
	}
	if order, ok := stage.First_frets_referenceOrder[first_fret]; ok {
		return order
	} else {
		log.Printf("instance %p of type First_fret was not staged and does not have a reference order", first_fret)
		return 0
	}
}

func (for_part *For_part) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.For_part_stagedOrder[for_part]; ok {
		return order
	}
	if order, ok := stage.For_parts_referenceOrder[for_part]; ok {
		return order
	} else {
		log.Printf("instance %p of type For_part was not staged and does not have a reference order", for_part)
		return 0
	}
}

func (formatted_symbol *Formatted_symbol) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Formatted_symbol_stagedOrder[formatted_symbol]; ok {
		return order
	}
	if order, ok := stage.Formatted_symbols_referenceOrder[formatted_symbol]; ok {
		return order
	} else {
		log.Printf("instance %p of type Formatted_symbol was not staged and does not have a reference order", formatted_symbol)
		return 0
	}
}

func (formatted_symbol_id *Formatted_symbol_id) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Formatted_symbol_id_stagedOrder[formatted_symbol_id]; ok {
		return order
	}
	if order, ok := stage.Formatted_symbol_ids_referenceOrder[formatted_symbol_id]; ok {
		return order
	} else {
		log.Printf("instance %p of type Formatted_symbol_id was not staged and does not have a reference order", formatted_symbol_id)
		return 0
	}
}

func (formatted_text *Formatted_text) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Formatted_text_stagedOrder[formatted_text]; ok {
		return order
	}
	if order, ok := stage.Formatted_texts_referenceOrder[formatted_text]; ok {
		return order
	} else {
		log.Printf("instance %p of type Formatted_text was not staged and does not have a reference order", formatted_text)
		return 0
	}
}

func (formatted_text_id *Formatted_text_id) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Formatted_text_id_stagedOrder[formatted_text_id]; ok {
		return order
	}
	if order, ok := stage.Formatted_text_ids_referenceOrder[formatted_text_id]; ok {
		return order
	} else {
		log.Printf("instance %p of type Formatted_text_id was not staged and does not have a reference order", formatted_text_id)
		return 0
	}
}

func (forward *Forward) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Forward_stagedOrder[forward]; ok {
		return order
	}
	if order, ok := stage.Forwards_referenceOrder[forward]; ok {
		return order
	} else {
		log.Printf("instance %p of type Forward was not staged and does not have a reference order", forward)
		return 0
	}
}

func (frame *Frame) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Frame_stagedOrder[frame]; ok {
		return order
	}
	if order, ok := stage.Frames_referenceOrder[frame]; ok {
		return order
	} else {
		log.Printf("instance %p of type Frame was not staged and does not have a reference order", frame)
		return 0
	}
}

func (frame_note *Frame_note) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Frame_note_stagedOrder[frame_note]; ok {
		return order
	}
	if order, ok := stage.Frame_notes_referenceOrder[frame_note]; ok {
		return order
	} else {
		log.Printf("instance %p of type Frame_note was not staged and does not have a reference order", frame_note)
		return 0
	}
}

func (fret *Fret) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Fret_stagedOrder[fret]; ok {
		return order
	}
	if order, ok := stage.Frets_referenceOrder[fret]; ok {
		return order
	} else {
		log.Printf("instance %p of type Fret was not staged and does not have a reference order", fret)
		return 0
	}
}

func (glass *Glass) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Glass_stagedOrder[glass]; ok {
		return order
	}
	if order, ok := stage.Glasss_referenceOrder[glass]; ok {
		return order
	} else {
		log.Printf("instance %p of type Glass was not staged and does not have a reference order", glass)
		return 0
	}
}

func (glissando *Glissando) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Glissando_stagedOrder[glissando]; ok {
		return order
	}
	if order, ok := stage.Glissandos_referenceOrder[glissando]; ok {
		return order
	} else {
		log.Printf("instance %p of type Glissando was not staged and does not have a reference order", glissando)
		return 0
	}
}

func (glyph *Glyph) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Glyph_stagedOrder[glyph]; ok {
		return order
	}
	if order, ok := stage.Glyphs_referenceOrder[glyph]; ok {
		return order
	} else {
		log.Printf("instance %p of type Glyph was not staged and does not have a reference order", glyph)
		return 0
	}
}

func (grace *Grace) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Grace_stagedOrder[grace]; ok {
		return order
	}
	if order, ok := stage.Graces_referenceOrder[grace]; ok {
		return order
	} else {
		log.Printf("instance %p of type Grace was not staged and does not have a reference order", grace)
		return 0
	}
}

func (group_barline *Group_barline) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Group_barline_stagedOrder[group_barline]; ok {
		return order
	}
	if order, ok := stage.Group_barlines_referenceOrder[group_barline]; ok {
		return order
	} else {
		log.Printf("instance %p of type Group_barline was not staged and does not have a reference order", group_barline)
		return 0
	}
}

func (group_name *Group_name) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Group_name_stagedOrder[group_name]; ok {
		return order
	}
	if order, ok := stage.Group_names_referenceOrder[group_name]; ok {
		return order
	} else {
		log.Printf("instance %p of type Group_name was not staged and does not have a reference order", group_name)
		return 0
	}
}

func (group_symbol *Group_symbol) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Group_symbol_stagedOrder[group_symbol]; ok {
		return order
	}
	if order, ok := stage.Group_symbols_referenceOrder[group_symbol]; ok {
		return order
	} else {
		log.Printf("instance %p of type Group_symbol was not staged and does not have a reference order", group_symbol)
		return 0
	}
}

func (grouping *Grouping) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Grouping_stagedOrder[grouping]; ok {
		return order
	}
	if order, ok := stage.Groupings_referenceOrder[grouping]; ok {
		return order
	} else {
		log.Printf("instance %p of type Grouping was not staged and does not have a reference order", grouping)
		return 0
	}
}

func (hammer_on_pull_off *Hammer_on_pull_off) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Hammer_on_pull_off_stagedOrder[hammer_on_pull_off]; ok {
		return order
	}
	if order, ok := stage.Hammer_on_pull_offs_referenceOrder[hammer_on_pull_off]; ok {
		return order
	} else {
		log.Printf("instance %p of type Hammer_on_pull_off was not staged and does not have a reference order", hammer_on_pull_off)
		return 0
	}
}

func (handbell *Handbell) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Handbell_stagedOrder[handbell]; ok {
		return order
	}
	if order, ok := stage.Handbells_referenceOrder[handbell]; ok {
		return order
	} else {
		log.Printf("instance %p of type Handbell was not staged and does not have a reference order", handbell)
		return 0
	}
}

func (harmon_closed *Harmon_closed) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Harmon_closed_stagedOrder[harmon_closed]; ok {
		return order
	}
	if order, ok := stage.Harmon_closeds_referenceOrder[harmon_closed]; ok {
		return order
	} else {
		log.Printf("instance %p of type Harmon_closed was not staged and does not have a reference order", harmon_closed)
		return 0
	}
}

func (harmon_mute *Harmon_mute) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Harmon_mute_stagedOrder[harmon_mute]; ok {
		return order
	}
	if order, ok := stage.Harmon_mutes_referenceOrder[harmon_mute]; ok {
		return order
	} else {
		log.Printf("instance %p of type Harmon_mute was not staged and does not have a reference order", harmon_mute)
		return 0
	}
}

func (harmonic *Harmonic) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Harmonic_stagedOrder[harmonic]; ok {
		return order
	}
	if order, ok := stage.Harmonics_referenceOrder[harmonic]; ok {
		return order
	} else {
		log.Printf("instance %p of type Harmonic was not staged and does not have a reference order", harmonic)
		return 0
	}
}

func (harmony *Harmony) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Harmony_stagedOrder[harmony]; ok {
		return order
	}
	if order, ok := stage.Harmonys_referenceOrder[harmony]; ok {
		return order
	} else {
		log.Printf("instance %p of type Harmony was not staged and does not have a reference order", harmony)
		return 0
	}
}

func (harmony_alter *Harmony_alter) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Harmony_alter_stagedOrder[harmony_alter]; ok {
		return order
	}
	if order, ok := stage.Harmony_alters_referenceOrder[harmony_alter]; ok {
		return order
	} else {
		log.Printf("instance %p of type Harmony_alter was not staged and does not have a reference order", harmony_alter)
		return 0
	}
}

func (harp_pedals *Harp_pedals) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Harp_pedals_stagedOrder[harp_pedals]; ok {
		return order
	}
	if order, ok := stage.Harp_pedalss_referenceOrder[harp_pedals]; ok {
		return order
	} else {
		log.Printf("instance %p of type Harp_pedals was not staged and does not have a reference order", harp_pedals)
		return 0
	}
}

func (heel_toe *Heel_toe) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Heel_toe_stagedOrder[heel_toe]; ok {
		return order
	}
	if order, ok := stage.Heel_toes_referenceOrder[heel_toe]; ok {
		return order
	} else {
		log.Printf("instance %p of type Heel_toe was not staged and does not have a reference order", heel_toe)
		return 0
	}
}

func (hole *Hole) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Hole_stagedOrder[hole]; ok {
		return order
	}
	if order, ok := stage.Holes_referenceOrder[hole]; ok {
		return order
	} else {
		log.Printf("instance %p of type Hole was not staged and does not have a reference order", hole)
		return 0
	}
}

func (hole_closed *Hole_closed) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Hole_closed_stagedOrder[hole_closed]; ok {
		return order
	}
	if order, ok := stage.Hole_closeds_referenceOrder[hole_closed]; ok {
		return order
	} else {
		log.Printf("instance %p of type Hole_closed was not staged and does not have a reference order", hole_closed)
		return 0
	}
}

func (horizontal_turn *Horizontal_turn) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Horizontal_turn_stagedOrder[horizontal_turn]; ok {
		return order
	}
	if order, ok := stage.Horizontal_turns_referenceOrder[horizontal_turn]; ok {
		return order
	} else {
		log.Printf("instance %p of type Horizontal_turn was not staged and does not have a reference order", horizontal_turn)
		return 0
	}
}

func (identification *Identification) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Identification_stagedOrder[identification]; ok {
		return order
	}
	if order, ok := stage.Identifications_referenceOrder[identification]; ok {
		return order
	} else {
		log.Printf("instance %p of type Identification was not staged and does not have a reference order", identification)
		return 0
	}
}

func (image *Image) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Image_stagedOrder[image]; ok {
		return order
	}
	if order, ok := stage.Images_referenceOrder[image]; ok {
		return order
	} else {
		log.Printf("instance %p of type Image was not staged and does not have a reference order", image)
		return 0
	}
}

func (instrument *Instrument) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Instrument_stagedOrder[instrument]; ok {
		return order
	}
	if order, ok := stage.Instruments_referenceOrder[instrument]; ok {
		return order
	} else {
		log.Printf("instance %p of type Instrument was not staged and does not have a reference order", instrument)
		return 0
	}
}

func (instrument_change *Instrument_change) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Instrument_change_stagedOrder[instrument_change]; ok {
		return order
	}
	if order, ok := stage.Instrument_changes_referenceOrder[instrument_change]; ok {
		return order
	} else {
		log.Printf("instance %p of type Instrument_change was not staged and does not have a reference order", instrument_change)
		return 0
	}
}

func (instrument_link *Instrument_link) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Instrument_link_stagedOrder[instrument_link]; ok {
		return order
	}
	if order, ok := stage.Instrument_links_referenceOrder[instrument_link]; ok {
		return order
	} else {
		log.Printf("instance %p of type Instrument_link was not staged and does not have a reference order", instrument_link)
		return 0
	}
}

func (interchangeable *Interchangeable) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Interchangeable_stagedOrder[interchangeable]; ok {
		return order
	}
	if order, ok := stage.Interchangeables_referenceOrder[interchangeable]; ok {
		return order
	} else {
		log.Printf("instance %p of type Interchangeable was not staged and does not have a reference order", interchangeable)
		return 0
	}
}

func (inversion *Inversion) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Inversion_stagedOrder[inversion]; ok {
		return order
	}
	if order, ok := stage.Inversions_referenceOrder[inversion]; ok {
		return order
	} else {
		log.Printf("instance %p of type Inversion was not staged and does not have a reference order", inversion)
		return 0
	}
}

func (key *Key) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Key_stagedOrder[key]; ok {
		return order
	}
	if order, ok := stage.Keys_referenceOrder[key]; ok {
		return order
	} else {
		log.Printf("instance %p of type Key was not staged and does not have a reference order", key)
		return 0
	}
}

func (key_accidental *Key_accidental) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Key_accidental_stagedOrder[key_accidental]; ok {
		return order
	}
	if order, ok := stage.Key_accidentals_referenceOrder[key_accidental]; ok {
		return order
	} else {
		log.Printf("instance %p of type Key_accidental was not staged and does not have a reference order", key_accidental)
		return 0
	}
}

func (key_octave *Key_octave) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Key_octave_stagedOrder[key_octave]; ok {
		return order
	}
	if order, ok := stage.Key_octaves_referenceOrder[key_octave]; ok {
		return order
	} else {
		log.Printf("instance %p of type Key_octave was not staged and does not have a reference order", key_octave)
		return 0
	}
}

func (kind *Kind) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Kind_stagedOrder[kind]; ok {
		return order
	}
	if order, ok := stage.Kinds_referenceOrder[kind]; ok {
		return order
	} else {
		log.Printf("instance %p of type Kind was not staged and does not have a reference order", kind)
		return 0
	}
}

func (level *Level) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Level_stagedOrder[level]; ok {
		return order
	}
	if order, ok := stage.Levels_referenceOrder[level]; ok {
		return order
	} else {
		log.Printf("instance %p of type Level was not staged and does not have a reference order", level)
		return 0
	}
}

func (line_detail *Line_detail) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Line_detail_stagedOrder[line_detail]; ok {
		return order
	}
	if order, ok := stage.Line_details_referenceOrder[line_detail]; ok {
		return order
	} else {
		log.Printf("instance %p of type Line_detail was not staged and does not have a reference order", line_detail)
		return 0
	}
}

func (line_width *Line_width) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Line_width_stagedOrder[line_width]; ok {
		return order
	}
	if order, ok := stage.Line_widths_referenceOrder[line_width]; ok {
		return order
	} else {
		log.Printf("instance %p of type Line_width was not staged and does not have a reference order", line_width)
		return 0
	}
}

func (link *Link) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Link_stagedOrder[link]; ok {
		return order
	}
	if order, ok := stage.Links_referenceOrder[link]; ok {
		return order
	} else {
		log.Printf("instance %p of type Link was not staged and does not have a reference order", link)
		return 0
	}
}

func (listen *Listen) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Listen_stagedOrder[listen]; ok {
		return order
	}
	if order, ok := stage.Listens_referenceOrder[listen]; ok {
		return order
	} else {
		log.Printf("instance %p of type Listen was not staged and does not have a reference order", listen)
		return 0
	}
}

func (listening *Listening) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Listening_stagedOrder[listening]; ok {
		return order
	}
	if order, ok := stage.Listenings_referenceOrder[listening]; ok {
		return order
	} else {
		log.Printf("instance %p of type Listening was not staged and does not have a reference order", listening)
		return 0
	}
}

func (lyric *Lyric) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Lyric_stagedOrder[lyric]; ok {
		return order
	}
	if order, ok := stage.Lyrics_referenceOrder[lyric]; ok {
		return order
	} else {
		log.Printf("instance %p of type Lyric was not staged and does not have a reference order", lyric)
		return 0
	}
}

func (lyric_font *Lyric_font) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Lyric_font_stagedOrder[lyric_font]; ok {
		return order
	}
	if order, ok := stage.Lyric_fonts_referenceOrder[lyric_font]; ok {
		return order
	} else {
		log.Printf("instance %p of type Lyric_font was not staged and does not have a reference order", lyric_font)
		return 0
	}
}

func (lyric_language *Lyric_language) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Lyric_language_stagedOrder[lyric_language]; ok {
		return order
	}
	if order, ok := stage.Lyric_languages_referenceOrder[lyric_language]; ok {
		return order
	} else {
		log.Printf("instance %p of type Lyric_language was not staged and does not have a reference order", lyric_language)
		return 0
	}
}

func (measure_layout *Measure_layout) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Measure_layout_stagedOrder[measure_layout]; ok {
		return order
	}
	if order, ok := stage.Measure_layouts_referenceOrder[measure_layout]; ok {
		return order
	} else {
		log.Printf("instance %p of type Measure_layout was not staged and does not have a reference order", measure_layout)
		return 0
	}
}

func (measure_numbering *Measure_numbering) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Measure_numbering_stagedOrder[measure_numbering]; ok {
		return order
	}
	if order, ok := stage.Measure_numberings_referenceOrder[measure_numbering]; ok {
		return order
	} else {
		log.Printf("instance %p of type Measure_numbering was not staged and does not have a reference order", measure_numbering)
		return 0
	}
}

func (measure_repeat *Measure_repeat) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Measure_repeat_stagedOrder[measure_repeat]; ok {
		return order
	}
	if order, ok := stage.Measure_repeats_referenceOrder[measure_repeat]; ok {
		return order
	} else {
		log.Printf("instance %p of type Measure_repeat was not staged and does not have a reference order", measure_repeat)
		return 0
	}
}

func (measure_style *Measure_style) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Measure_style_stagedOrder[measure_style]; ok {
		return order
	}
	if order, ok := stage.Measure_styles_referenceOrder[measure_style]; ok {
		return order
	} else {
		log.Printf("instance %p of type Measure_style was not staged and does not have a reference order", measure_style)
		return 0
	}
}

func (membrane *Membrane) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Membrane_stagedOrder[membrane]; ok {
		return order
	}
	if order, ok := stage.Membranes_referenceOrder[membrane]; ok {
		return order
	} else {
		log.Printf("instance %p of type Membrane was not staged and does not have a reference order", membrane)
		return 0
	}
}

func (metal *Metal) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Metal_stagedOrder[metal]; ok {
		return order
	}
	if order, ok := stage.Metals_referenceOrder[metal]; ok {
		return order
	} else {
		log.Printf("instance %p of type Metal was not staged and does not have a reference order", metal)
		return 0
	}
}

func (metronome *Metronome) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Metronome_stagedOrder[metronome]; ok {
		return order
	}
	if order, ok := stage.Metronomes_referenceOrder[metronome]; ok {
		return order
	} else {
		log.Printf("instance %p of type Metronome was not staged and does not have a reference order", metronome)
		return 0
	}
}

func (metronome_beam *Metronome_beam) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Metronome_beam_stagedOrder[metronome_beam]; ok {
		return order
	}
	if order, ok := stage.Metronome_beams_referenceOrder[metronome_beam]; ok {
		return order
	} else {
		log.Printf("instance %p of type Metronome_beam was not staged and does not have a reference order", metronome_beam)
		return 0
	}
}

func (metronome_note *Metronome_note) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Metronome_note_stagedOrder[metronome_note]; ok {
		return order
	}
	if order, ok := stage.Metronome_notes_referenceOrder[metronome_note]; ok {
		return order
	} else {
		log.Printf("instance %p of type Metronome_note was not staged and does not have a reference order", metronome_note)
		return 0
	}
}

func (metronome_tied *Metronome_tied) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Metronome_tied_stagedOrder[metronome_tied]; ok {
		return order
	}
	if order, ok := stage.Metronome_tieds_referenceOrder[metronome_tied]; ok {
		return order
	} else {
		log.Printf("instance %p of type Metronome_tied was not staged and does not have a reference order", metronome_tied)
		return 0
	}
}

func (metronome_tuplet *Metronome_tuplet) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Metronome_tuplet_stagedOrder[metronome_tuplet]; ok {
		return order
	}
	if order, ok := stage.Metronome_tuplets_referenceOrder[metronome_tuplet]; ok {
		return order
	} else {
		log.Printf("instance %p of type Metronome_tuplet was not staged and does not have a reference order", metronome_tuplet)
		return 0
	}
}

func (midi_device *Midi_device) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Midi_device_stagedOrder[midi_device]; ok {
		return order
	}
	if order, ok := stage.Midi_devices_referenceOrder[midi_device]; ok {
		return order
	} else {
		log.Printf("instance %p of type Midi_device was not staged and does not have a reference order", midi_device)
		return 0
	}
}

func (midi_instrument *Midi_instrument) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Midi_instrument_stagedOrder[midi_instrument]; ok {
		return order
	}
	if order, ok := stage.Midi_instruments_referenceOrder[midi_instrument]; ok {
		return order
	} else {
		log.Printf("instance %p of type Midi_instrument was not staged and does not have a reference order", midi_instrument)
		return 0
	}
}

func (miscellaneous *Miscellaneous) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Miscellaneous_stagedOrder[miscellaneous]; ok {
		return order
	}
	if order, ok := stage.Miscellaneouss_referenceOrder[miscellaneous]; ok {
		return order
	} else {
		log.Printf("instance %p of type Miscellaneous was not staged and does not have a reference order", miscellaneous)
		return 0
	}
}

func (miscellaneous_field *Miscellaneous_field) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Miscellaneous_field_stagedOrder[miscellaneous_field]; ok {
		return order
	}
	if order, ok := stage.Miscellaneous_fields_referenceOrder[miscellaneous_field]; ok {
		return order
	} else {
		log.Printf("instance %p of type Miscellaneous_field was not staged and does not have a reference order", miscellaneous_field)
		return 0
	}
}

func (mordent *Mordent) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Mordent_stagedOrder[mordent]; ok {
		return order
	}
	if order, ok := stage.Mordents_referenceOrder[mordent]; ok {
		return order
	} else {
		log.Printf("instance %p of type Mordent was not staged and does not have a reference order", mordent)
		return 0
	}
}

func (multiple_rest *Multiple_rest) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Multiple_rest_stagedOrder[multiple_rest]; ok {
		return order
	}
	if order, ok := stage.Multiple_rests_referenceOrder[multiple_rest]; ok {
		return order
	} else {
		log.Printf("instance %p of type Multiple_rest was not staged and does not have a reference order", multiple_rest)
		return 0
	}
}

func (name_display *Name_display) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Name_display_stagedOrder[name_display]; ok {
		return order
	}
	if order, ok := stage.Name_displays_referenceOrder[name_display]; ok {
		return order
	} else {
		log.Printf("instance %p of type Name_display was not staged and does not have a reference order", name_display)
		return 0
	}
}

func (non_arpeggiate *Non_arpeggiate) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Non_arpeggiate_stagedOrder[non_arpeggiate]; ok {
		return order
	}
	if order, ok := stage.Non_arpeggiates_referenceOrder[non_arpeggiate]; ok {
		return order
	} else {
		log.Printf("instance %p of type Non_arpeggiate was not staged and does not have a reference order", non_arpeggiate)
		return 0
	}
}

func (notations *Notations) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Notations_stagedOrder[notations]; ok {
		return order
	}
	if order, ok := stage.Notationss_referenceOrder[notations]; ok {
		return order
	} else {
		log.Printf("instance %p of type Notations was not staged and does not have a reference order", notations)
		return 0
	}
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Note_stagedOrder[note]; ok {
		return order
	}
	if order, ok := stage.Notes_referenceOrder[note]; ok {
		return order
	} else {
		log.Printf("instance %p of type Note was not staged and does not have a reference order", note)
		return 0
	}
}

func (note_size *Note_size) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Note_size_stagedOrder[note_size]; ok {
		return order
	}
	if order, ok := stage.Note_sizes_referenceOrder[note_size]; ok {
		return order
	} else {
		log.Printf("instance %p of type Note_size was not staged and does not have a reference order", note_size)
		return 0
	}
}

func (note_type *Note_type) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Note_type_stagedOrder[note_type]; ok {
		return order
	}
	if order, ok := stage.Note_types_referenceOrder[note_type]; ok {
		return order
	} else {
		log.Printf("instance %p of type Note_type was not staged and does not have a reference order", note_type)
		return 0
	}
}

func (notehead *Notehead) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Notehead_stagedOrder[notehead]; ok {
		return order
	}
	if order, ok := stage.Noteheads_referenceOrder[notehead]; ok {
		return order
	} else {
		log.Printf("instance %p of type Notehead was not staged and does not have a reference order", notehead)
		return 0
	}
}

func (notehead_text *Notehead_text) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Notehead_text_stagedOrder[notehead_text]; ok {
		return order
	}
	if order, ok := stage.Notehead_texts_referenceOrder[notehead_text]; ok {
		return order
	} else {
		log.Printf("instance %p of type Notehead_text was not staged and does not have a reference order", notehead_text)
		return 0
	}
}

func (numeral *Numeral) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Numeral_stagedOrder[numeral]; ok {
		return order
	}
	if order, ok := stage.Numerals_referenceOrder[numeral]; ok {
		return order
	} else {
		log.Printf("instance %p of type Numeral was not staged and does not have a reference order", numeral)
		return 0
	}
}

func (numeral_key *Numeral_key) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Numeral_key_stagedOrder[numeral_key]; ok {
		return order
	}
	if order, ok := stage.Numeral_keys_referenceOrder[numeral_key]; ok {
		return order
	} else {
		log.Printf("instance %p of type Numeral_key was not staged and does not have a reference order", numeral_key)
		return 0
	}
}

func (numeral_root *Numeral_root) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Numeral_root_stagedOrder[numeral_root]; ok {
		return order
	}
	if order, ok := stage.Numeral_roots_referenceOrder[numeral_root]; ok {
		return order
	} else {
		log.Printf("instance %p of type Numeral_root was not staged and does not have a reference order", numeral_root)
		return 0
	}
}

func (octave_shift *Octave_shift) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Octave_shift_stagedOrder[octave_shift]; ok {
		return order
	}
	if order, ok := stage.Octave_shifts_referenceOrder[octave_shift]; ok {
		return order
	} else {
		log.Printf("instance %p of type Octave_shift was not staged and does not have a reference order", octave_shift)
		return 0
	}
}

func (offset *Offset) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Offset_stagedOrder[offset]; ok {
		return order
	}
	if order, ok := stage.Offsets_referenceOrder[offset]; ok {
		return order
	} else {
		log.Printf("instance %p of type Offset was not staged and does not have a reference order", offset)
		return 0
	}
}

func (opus *Opus) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Opus_stagedOrder[opus]; ok {
		return order
	}
	if order, ok := stage.Opuss_referenceOrder[opus]; ok {
		return order
	} else {
		log.Printf("instance %p of type Opus was not staged and does not have a reference order", opus)
		return 0
	}
}

func (ornaments *Ornaments) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Ornaments_stagedOrder[ornaments]; ok {
		return order
	}
	if order, ok := stage.Ornamentss_referenceOrder[ornaments]; ok {
		return order
	} else {
		log.Printf("instance %p of type Ornaments was not staged and does not have a reference order", ornaments)
		return 0
	}
}

func (other_appearance *Other_appearance) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Other_appearance_stagedOrder[other_appearance]; ok {
		return order
	}
	if order, ok := stage.Other_appearances_referenceOrder[other_appearance]; ok {
		return order
	} else {
		log.Printf("instance %p of type Other_appearance was not staged and does not have a reference order", other_appearance)
		return 0
	}
}

func (other_direction *Other_direction) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Other_direction_stagedOrder[other_direction]; ok {
		return order
	}
	if order, ok := stage.Other_directions_referenceOrder[other_direction]; ok {
		return order
	} else {
		log.Printf("instance %p of type Other_direction was not staged and does not have a reference order", other_direction)
		return 0
	}
}

func (other_listening *Other_listening) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Other_listening_stagedOrder[other_listening]; ok {
		return order
	}
	if order, ok := stage.Other_listenings_referenceOrder[other_listening]; ok {
		return order
	} else {
		log.Printf("instance %p of type Other_listening was not staged and does not have a reference order", other_listening)
		return 0
	}
}

func (other_notation *Other_notation) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Other_notation_stagedOrder[other_notation]; ok {
		return order
	}
	if order, ok := stage.Other_notations_referenceOrder[other_notation]; ok {
		return order
	} else {
		log.Printf("instance %p of type Other_notation was not staged and does not have a reference order", other_notation)
		return 0
	}
}

func (other_placement_text *Other_placement_text) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Other_placement_text_stagedOrder[other_placement_text]; ok {
		return order
	}
	if order, ok := stage.Other_placement_texts_referenceOrder[other_placement_text]; ok {
		return order
	} else {
		log.Printf("instance %p of type Other_placement_text was not staged and does not have a reference order", other_placement_text)
		return 0
	}
}

func (other_play *Other_play) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Other_play_stagedOrder[other_play]; ok {
		return order
	}
	if order, ok := stage.Other_plays_referenceOrder[other_play]; ok {
		return order
	} else {
		log.Printf("instance %p of type Other_play was not staged and does not have a reference order", other_play)
		return 0
	}
}

func (other_text *Other_text) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Other_text_stagedOrder[other_text]; ok {
		return order
	}
	if order, ok := stage.Other_texts_referenceOrder[other_text]; ok {
		return order
	} else {
		log.Printf("instance %p of type Other_text was not staged and does not have a reference order", other_text)
		return 0
	}
}

func (page_layout *Page_layout) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Page_layout_stagedOrder[page_layout]; ok {
		return order
	}
	if order, ok := stage.Page_layouts_referenceOrder[page_layout]; ok {
		return order
	} else {
		log.Printf("instance %p of type Page_layout was not staged and does not have a reference order", page_layout)
		return 0
	}
}

func (page_margins *Page_margins) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Page_margins_stagedOrder[page_margins]; ok {
		return order
	}
	if order, ok := stage.Page_marginss_referenceOrder[page_margins]; ok {
		return order
	} else {
		log.Printf("instance %p of type Page_margins was not staged and does not have a reference order", page_margins)
		return 0
	}
}

func (part_clef *Part_clef) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Part_clef_stagedOrder[part_clef]; ok {
		return order
	}
	if order, ok := stage.Part_clefs_referenceOrder[part_clef]; ok {
		return order
	} else {
		log.Printf("instance %p of type Part_clef was not staged and does not have a reference order", part_clef)
		return 0
	}
}

func (part_group *Part_group) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Part_group_stagedOrder[part_group]; ok {
		return order
	}
	if order, ok := stage.Part_groups_referenceOrder[part_group]; ok {
		return order
	} else {
		log.Printf("instance %p of type Part_group was not staged and does not have a reference order", part_group)
		return 0
	}
}

func (part_link *Part_link) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Part_link_stagedOrder[part_link]; ok {
		return order
	}
	if order, ok := stage.Part_links_referenceOrder[part_link]; ok {
		return order
	} else {
		log.Printf("instance %p of type Part_link was not staged and does not have a reference order", part_link)
		return 0
	}
}

func (part_list *Part_list) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Part_list_stagedOrder[part_list]; ok {
		return order
	}
	if order, ok := stage.Part_lists_referenceOrder[part_list]; ok {
		return order
	} else {
		log.Printf("instance %p of type Part_list was not staged and does not have a reference order", part_list)
		return 0
	}
}

func (part_name *Part_name) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Part_name_stagedOrder[part_name]; ok {
		return order
	}
	if order, ok := stage.Part_names_referenceOrder[part_name]; ok {
		return order
	} else {
		log.Printf("instance %p of type Part_name was not staged and does not have a reference order", part_name)
		return 0
	}
}

func (part_symbol *Part_symbol) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Part_symbol_stagedOrder[part_symbol]; ok {
		return order
	}
	if order, ok := stage.Part_symbols_referenceOrder[part_symbol]; ok {
		return order
	} else {
		log.Printf("instance %p of type Part_symbol was not staged and does not have a reference order", part_symbol)
		return 0
	}
}

func (part_transpose *Part_transpose) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Part_transpose_stagedOrder[part_transpose]; ok {
		return order
	}
	if order, ok := stage.Part_transposes_referenceOrder[part_transpose]; ok {
		return order
	} else {
		log.Printf("instance %p of type Part_transpose was not staged and does not have a reference order", part_transpose)
		return 0
	}
}

func (pedal *Pedal) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Pedal_stagedOrder[pedal]; ok {
		return order
	}
	if order, ok := stage.Pedals_referenceOrder[pedal]; ok {
		return order
	} else {
		log.Printf("instance %p of type Pedal was not staged and does not have a reference order", pedal)
		return 0
	}
}

func (pedal_tuning *Pedal_tuning) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Pedal_tuning_stagedOrder[pedal_tuning]; ok {
		return order
	}
	if order, ok := stage.Pedal_tunings_referenceOrder[pedal_tuning]; ok {
		return order
	} else {
		log.Printf("instance %p of type Pedal_tuning was not staged and does not have a reference order", pedal_tuning)
		return 0
	}
}

func (per_minute *Per_minute) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Per_minute_stagedOrder[per_minute]; ok {
		return order
	}
	if order, ok := stage.Per_minutes_referenceOrder[per_minute]; ok {
		return order
	} else {
		log.Printf("instance %p of type Per_minute was not staged and does not have a reference order", per_minute)
		return 0
	}
}

func (percussion *Percussion) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Percussion_stagedOrder[percussion]; ok {
		return order
	}
	if order, ok := stage.Percussions_referenceOrder[percussion]; ok {
		return order
	} else {
		log.Printf("instance %p of type Percussion was not staged and does not have a reference order", percussion)
		return 0
	}
}

func (pitch *Pitch) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Pitch_stagedOrder[pitch]; ok {
		return order
	}
	if order, ok := stage.Pitchs_referenceOrder[pitch]; ok {
		return order
	} else {
		log.Printf("instance %p of type Pitch was not staged and does not have a reference order", pitch)
		return 0
	}
}

func (pitched *Pitched) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Pitched_stagedOrder[pitched]; ok {
		return order
	}
	if order, ok := stage.Pitcheds_referenceOrder[pitched]; ok {
		return order
	} else {
		log.Printf("instance %p of type Pitched was not staged and does not have a reference order", pitched)
		return 0
	}
}

func (placement_text *Placement_text) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Placement_text_stagedOrder[placement_text]; ok {
		return order
	}
	if order, ok := stage.Placement_texts_referenceOrder[placement_text]; ok {
		return order
	} else {
		log.Printf("instance %p of type Placement_text was not staged and does not have a reference order", placement_text)
		return 0
	}
}

func (play *Play) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Play_stagedOrder[play]; ok {
		return order
	}
	if order, ok := stage.Plays_referenceOrder[play]; ok {
		return order
	} else {
		log.Printf("instance %p of type Play was not staged and does not have a reference order", play)
		return 0
	}
}

func (player *Player) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Player_stagedOrder[player]; ok {
		return order
	}
	if order, ok := stage.Players_referenceOrder[player]; ok {
		return order
	} else {
		log.Printf("instance %p of type Player was not staged and does not have a reference order", player)
		return 0
	}
}

func (principal_voice *Principal_voice) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Principal_voice_stagedOrder[principal_voice]; ok {
		return order
	}
	if order, ok := stage.Principal_voices_referenceOrder[principal_voice]; ok {
		return order
	} else {
		log.Printf("instance %p of type Principal_voice was not staged and does not have a reference order", principal_voice)
		return 0
	}
}

func (print *Print) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Print_stagedOrder[print]; ok {
		return order
	}
	if order, ok := stage.Prints_referenceOrder[print]; ok {
		return order
	} else {
		log.Printf("instance %p of type Print was not staged and does not have a reference order", print)
		return 0
	}
}

func (release *Release) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Release_stagedOrder[release]; ok {
		return order
	}
	if order, ok := stage.Releases_referenceOrder[release]; ok {
		return order
	} else {
		log.Printf("instance %p of type Release was not staged and does not have a reference order", release)
		return 0
	}
}

func (repeat *Repeat) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Repeat_stagedOrder[repeat]; ok {
		return order
	}
	if order, ok := stage.Repeats_referenceOrder[repeat]; ok {
		return order
	} else {
		log.Printf("instance %p of type Repeat was not staged and does not have a reference order", repeat)
		return 0
	}
}

func (rest *Rest) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Rest_stagedOrder[rest]; ok {
		return order
	}
	if order, ok := stage.Rests_referenceOrder[rest]; ok {
		return order
	} else {
		log.Printf("instance %p of type Rest was not staged and does not have a reference order", rest)
		return 0
	}
}

func (root *Root) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Root_stagedOrder[root]; ok {
		return order
	}
	if order, ok := stage.Roots_referenceOrder[root]; ok {
		return order
	} else {
		log.Printf("instance %p of type Root was not staged and does not have a reference order", root)
		return 0
	}
}

func (root_step *Root_step) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Root_step_stagedOrder[root_step]; ok {
		return order
	}
	if order, ok := stage.Root_steps_referenceOrder[root_step]; ok {
		return order
	} else {
		log.Printf("instance %p of type Root_step was not staged and does not have a reference order", root_step)
		return 0
	}
}

func (scaling *Scaling) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Scaling_stagedOrder[scaling]; ok {
		return order
	}
	if order, ok := stage.Scalings_referenceOrder[scaling]; ok {
		return order
	} else {
		log.Printf("instance %p of type Scaling was not staged and does not have a reference order", scaling)
		return 0
	}
}

func (scordatura *Scordatura) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Scordatura_stagedOrder[scordatura]; ok {
		return order
	}
	if order, ok := stage.Scordaturas_referenceOrder[scordatura]; ok {
		return order
	} else {
		log.Printf("instance %p of type Scordatura was not staged and does not have a reference order", scordatura)
		return 0
	}
}

func (score_instrument *Score_instrument) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Score_instrument_stagedOrder[score_instrument]; ok {
		return order
	}
	if order, ok := stage.Score_instruments_referenceOrder[score_instrument]; ok {
		return order
	} else {
		log.Printf("instance %p of type Score_instrument was not staged and does not have a reference order", score_instrument)
		return 0
	}
}

func (score_part *Score_part) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Score_part_stagedOrder[score_part]; ok {
		return order
	}
	if order, ok := stage.Score_parts_referenceOrder[score_part]; ok {
		return order
	} else {
		log.Printf("instance %p of type Score_part was not staged and does not have a reference order", score_part)
		return 0
	}
}

func (score_partwise *Score_partwise) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Score_partwise_stagedOrder[score_partwise]; ok {
		return order
	}
	if order, ok := stage.Score_partwises_referenceOrder[score_partwise]; ok {
		return order
	} else {
		log.Printf("instance %p of type Score_partwise was not staged and does not have a reference order", score_partwise)
		return 0
	}
}

func (score_timewise *Score_timewise) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Score_timewise_stagedOrder[score_timewise]; ok {
		return order
	}
	if order, ok := stage.Score_timewises_referenceOrder[score_timewise]; ok {
		return order
	} else {
		log.Printf("instance %p of type Score_timewise was not staged and does not have a reference order", score_timewise)
		return 0
	}
}

func (segno *Segno) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Segno_stagedOrder[segno]; ok {
		return order
	}
	if order, ok := stage.Segnos_referenceOrder[segno]; ok {
		return order
	} else {
		log.Printf("instance %p of type Segno was not staged and does not have a reference order", segno)
		return 0
	}
}

func (slash *Slash) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Slash_stagedOrder[slash]; ok {
		return order
	}
	if order, ok := stage.Slashs_referenceOrder[slash]; ok {
		return order
	} else {
		log.Printf("instance %p of type Slash was not staged and does not have a reference order", slash)
		return 0
	}
}

func (slide *Slide) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Slide_stagedOrder[slide]; ok {
		return order
	}
	if order, ok := stage.Slides_referenceOrder[slide]; ok {
		return order
	} else {
		log.Printf("instance %p of type Slide was not staged and does not have a reference order", slide)
		return 0
	}
}

func (slur *Slur) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Slur_stagedOrder[slur]; ok {
		return order
	}
	if order, ok := stage.Slurs_referenceOrder[slur]; ok {
		return order
	} else {
		log.Printf("instance %p of type Slur was not staged and does not have a reference order", slur)
		return 0
	}
}

func (sound *Sound) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Sound_stagedOrder[sound]; ok {
		return order
	}
	if order, ok := stage.Sounds_referenceOrder[sound]; ok {
		return order
	} else {
		log.Printf("instance %p of type Sound was not staged and does not have a reference order", sound)
		return 0
	}
}

func (staff_details *Staff_details) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Staff_details_stagedOrder[staff_details]; ok {
		return order
	}
	if order, ok := stage.Staff_detailss_referenceOrder[staff_details]; ok {
		return order
	} else {
		log.Printf("instance %p of type Staff_details was not staged and does not have a reference order", staff_details)
		return 0
	}
}

func (staff_divide *Staff_divide) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Staff_divide_stagedOrder[staff_divide]; ok {
		return order
	}
	if order, ok := stage.Staff_divides_referenceOrder[staff_divide]; ok {
		return order
	} else {
		log.Printf("instance %p of type Staff_divide was not staged and does not have a reference order", staff_divide)
		return 0
	}
}

func (staff_layout *Staff_layout) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Staff_layout_stagedOrder[staff_layout]; ok {
		return order
	}
	if order, ok := stage.Staff_layouts_referenceOrder[staff_layout]; ok {
		return order
	} else {
		log.Printf("instance %p of type Staff_layout was not staged and does not have a reference order", staff_layout)
		return 0
	}
}

func (staff_size *Staff_size) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Staff_size_stagedOrder[staff_size]; ok {
		return order
	}
	if order, ok := stage.Staff_sizes_referenceOrder[staff_size]; ok {
		return order
	} else {
		log.Printf("instance %p of type Staff_size was not staged and does not have a reference order", staff_size)
		return 0
	}
}

func (staff_tuning *Staff_tuning) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Staff_tuning_stagedOrder[staff_tuning]; ok {
		return order
	}
	if order, ok := stage.Staff_tunings_referenceOrder[staff_tuning]; ok {
		return order
	} else {
		log.Printf("instance %p of type Staff_tuning was not staged and does not have a reference order", staff_tuning)
		return 0
	}
}

func (stem *Stem) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Stem_stagedOrder[stem]; ok {
		return order
	}
	if order, ok := stage.Stems_referenceOrder[stem]; ok {
		return order
	} else {
		log.Printf("instance %p of type Stem was not staged and does not have a reference order", stem)
		return 0
	}
}

func (stick *Stick) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Stick_stagedOrder[stick]; ok {
		return order
	}
	if order, ok := stage.Sticks_referenceOrder[stick]; ok {
		return order
	} else {
		log.Printf("instance %p of type Stick was not staged and does not have a reference order", stick)
		return 0
	}
}

func (string_mute *String_mute) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.String_mute_stagedOrder[string_mute]; ok {
		return order
	}
	if order, ok := stage.String_mutes_referenceOrder[string_mute]; ok {
		return order
	} else {
		log.Printf("instance %p of type String_mute was not staged and does not have a reference order", string_mute)
		return 0
	}
}

func (string_type *String_type) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.String_type_stagedOrder[string_type]; ok {
		return order
	}
	if order, ok := stage.String_types_referenceOrder[string_type]; ok {
		return order
	} else {
		log.Printf("instance %p of type String_type was not staged and does not have a reference order", string_type)
		return 0
	}
}

func (strong_accent *Strong_accent) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Strong_accent_stagedOrder[strong_accent]; ok {
		return order
	}
	if order, ok := stage.Strong_accents_referenceOrder[strong_accent]; ok {
		return order
	} else {
		log.Printf("instance %p of type Strong_accent was not staged and does not have a reference order", strong_accent)
		return 0
	}
}

func (style_text *Style_text) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Style_text_stagedOrder[style_text]; ok {
		return order
	}
	if order, ok := stage.Style_texts_referenceOrder[style_text]; ok {
		return order
	} else {
		log.Printf("instance %p of type Style_text was not staged and does not have a reference order", style_text)
		return 0
	}
}

func (supports *Supports) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Supports_stagedOrder[supports]; ok {
		return order
	}
	if order, ok := stage.Supportss_referenceOrder[supports]; ok {
		return order
	} else {
		log.Printf("instance %p of type Supports was not staged and does not have a reference order", supports)
		return 0
	}
}

func (swing *Swing) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Swing_stagedOrder[swing]; ok {
		return order
	}
	if order, ok := stage.Swings_referenceOrder[swing]; ok {
		return order
	} else {
		log.Printf("instance %p of type Swing was not staged and does not have a reference order", swing)
		return 0
	}
}

func (sync *Sync) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Sync_stagedOrder[sync]; ok {
		return order
	}
	if order, ok := stage.Syncs_referenceOrder[sync]; ok {
		return order
	} else {
		log.Printf("instance %p of type Sync was not staged and does not have a reference order", sync)
		return 0
	}
}

func (system_dividers *System_dividers) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.System_dividers_stagedOrder[system_dividers]; ok {
		return order
	}
	if order, ok := stage.System_dividerss_referenceOrder[system_dividers]; ok {
		return order
	} else {
		log.Printf("instance %p of type System_dividers was not staged and does not have a reference order", system_dividers)
		return 0
	}
}

func (system_layout *System_layout) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.System_layout_stagedOrder[system_layout]; ok {
		return order
	}
	if order, ok := stage.System_layouts_referenceOrder[system_layout]; ok {
		return order
	} else {
		log.Printf("instance %p of type System_layout was not staged and does not have a reference order", system_layout)
		return 0
	}
}

func (system_margins *System_margins) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.System_margins_stagedOrder[system_margins]; ok {
		return order
	}
	if order, ok := stage.System_marginss_referenceOrder[system_margins]; ok {
		return order
	} else {
		log.Printf("instance %p of type System_margins was not staged and does not have a reference order", system_margins)
		return 0
	}
}

func (tap *Tap) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Tap_stagedOrder[tap]; ok {
		return order
	}
	if order, ok := stage.Taps_referenceOrder[tap]; ok {
		return order
	} else {
		log.Printf("instance %p of type Tap was not staged and does not have a reference order", tap)
		return 0
	}
}

func (technical *Technical) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Technical_stagedOrder[technical]; ok {
		return order
	}
	if order, ok := stage.Technicals_referenceOrder[technical]; ok {
		return order
	} else {
		log.Printf("instance %p of type Technical was not staged and does not have a reference order", technical)
		return 0
	}
}

func (text_element_data *Text_element_data) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Text_element_data_stagedOrder[text_element_data]; ok {
		return order
	}
	if order, ok := stage.Text_element_datas_referenceOrder[text_element_data]; ok {
		return order
	} else {
		log.Printf("instance %p of type Text_element_data was not staged and does not have a reference order", text_element_data)
		return 0
	}
}

func (tie *Tie) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Tie_stagedOrder[tie]; ok {
		return order
	}
	if order, ok := stage.Ties_referenceOrder[tie]; ok {
		return order
	} else {
		log.Printf("instance %p of type Tie was not staged and does not have a reference order", tie)
		return 0
	}
}

func (tied *Tied) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Tied_stagedOrder[tied]; ok {
		return order
	}
	if order, ok := stage.Tieds_referenceOrder[tied]; ok {
		return order
	} else {
		log.Printf("instance %p of type Tied was not staged and does not have a reference order", tied)
		return 0
	}
}

func (time *Time) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Time_stagedOrder[time]; ok {
		return order
	}
	if order, ok := stage.Times_referenceOrder[time]; ok {
		return order
	} else {
		log.Printf("instance %p of type Time was not staged and does not have a reference order", time)
		return 0
	}
}

func (time_modification *Time_modification) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Time_modification_stagedOrder[time_modification]; ok {
		return order
	}
	if order, ok := stage.Time_modifications_referenceOrder[time_modification]; ok {
		return order
	} else {
		log.Printf("instance %p of type Time_modification was not staged and does not have a reference order", time_modification)
		return 0
	}
}

func (timpani *Timpani) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Timpani_stagedOrder[timpani]; ok {
		return order
	}
	if order, ok := stage.Timpanis_referenceOrder[timpani]; ok {
		return order
	} else {
		log.Printf("instance %p of type Timpani was not staged and does not have a reference order", timpani)
		return 0
	}
}

func (transpose *Transpose) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Transpose_stagedOrder[transpose]; ok {
		return order
	}
	if order, ok := stage.Transposes_referenceOrder[transpose]; ok {
		return order
	} else {
		log.Printf("instance %p of type Transpose was not staged and does not have a reference order", transpose)
		return 0
	}
}

func (tremolo *Tremolo) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Tremolo_stagedOrder[tremolo]; ok {
		return order
	}
	if order, ok := stage.Tremolos_referenceOrder[tremolo]; ok {
		return order
	} else {
		log.Printf("instance %p of type Tremolo was not staged and does not have a reference order", tremolo)
		return 0
	}
}

func (tuplet *Tuplet) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Tuplet_stagedOrder[tuplet]; ok {
		return order
	}
	if order, ok := stage.Tuplets_referenceOrder[tuplet]; ok {
		return order
	} else {
		log.Printf("instance %p of type Tuplet was not staged and does not have a reference order", tuplet)
		return 0
	}
}

func (tuplet_dot *Tuplet_dot) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Tuplet_dot_stagedOrder[tuplet_dot]; ok {
		return order
	}
	if order, ok := stage.Tuplet_dots_referenceOrder[tuplet_dot]; ok {
		return order
	} else {
		log.Printf("instance %p of type Tuplet_dot was not staged and does not have a reference order", tuplet_dot)
		return 0
	}
}

func (tuplet_number *Tuplet_number) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Tuplet_number_stagedOrder[tuplet_number]; ok {
		return order
	}
	if order, ok := stage.Tuplet_numbers_referenceOrder[tuplet_number]; ok {
		return order
	} else {
		log.Printf("instance %p of type Tuplet_number was not staged and does not have a reference order", tuplet_number)
		return 0
	}
}

func (tuplet_portion *Tuplet_portion) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Tuplet_portion_stagedOrder[tuplet_portion]; ok {
		return order
	}
	if order, ok := stage.Tuplet_portions_referenceOrder[tuplet_portion]; ok {
		return order
	} else {
		log.Printf("instance %p of type Tuplet_portion was not staged and does not have a reference order", tuplet_portion)
		return 0
	}
}

func (tuplet_type *Tuplet_type) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Tuplet_type_stagedOrder[tuplet_type]; ok {
		return order
	}
	if order, ok := stage.Tuplet_types_referenceOrder[tuplet_type]; ok {
		return order
	} else {
		log.Printf("instance %p of type Tuplet_type was not staged and does not have a reference order", tuplet_type)
		return 0
	}
}

func (typed_text *Typed_text) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Typed_text_stagedOrder[typed_text]; ok {
		return order
	}
	if order, ok := stage.Typed_texts_referenceOrder[typed_text]; ok {
		return order
	} else {
		log.Printf("instance %p of type Typed_text was not staged and does not have a reference order", typed_text)
		return 0
	}
}

func (unpitched *Unpitched) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Unpitched_stagedOrder[unpitched]; ok {
		return order
	}
	if order, ok := stage.Unpitcheds_referenceOrder[unpitched]; ok {
		return order
	} else {
		log.Printf("instance %p of type Unpitched was not staged and does not have a reference order", unpitched)
		return 0
	}
}

func (virtual_instrument *Virtual_instrument) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Virtual_instrument_stagedOrder[virtual_instrument]; ok {
		return order
	}
	if order, ok := stage.Virtual_instruments_referenceOrder[virtual_instrument]; ok {
		return order
	} else {
		log.Printf("instance %p of type Virtual_instrument was not staged and does not have a reference order", virtual_instrument)
		return 0
	}
}

func (wait *Wait) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Wait_stagedOrder[wait]; ok {
		return order
	}
	if order, ok := stage.Waits_referenceOrder[wait]; ok {
		return order
	} else {
		log.Printf("instance %p of type Wait was not staged and does not have a reference order", wait)
		return 0
	}
}

func (wavy_line *Wavy_line) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Wavy_line_stagedOrder[wavy_line]; ok {
		return order
	}
	if order, ok := stage.Wavy_lines_referenceOrder[wavy_line]; ok {
		return order
	} else {
		log.Printf("instance %p of type Wavy_line was not staged and does not have a reference order", wavy_line)
		return 0
	}
}

func (wedge *Wedge) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Wedge_stagedOrder[wedge]; ok {
		return order
	}
	if order, ok := stage.Wedges_referenceOrder[wedge]; ok {
		return order
	} else {
		log.Printf("instance %p of type Wedge was not staged and does not have a reference order", wedge)
		return 0
	}
}

func (wood *Wood) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Wood_stagedOrder[wood]; ok {
		return order
	}
	if order, ok := stage.Woods_referenceOrder[wood]; ok {
		return order
	} else {
		log.Printf("instance %p of type Wood was not staged and does not have a reference order", wood)
		return 0
	}
}

func (work *Work) GongGetOrder(stage *Stage) uint {
	if order, ok := stage.Work_stagedOrder[work]; ok {
		return order
	}
	if order, ok := stage.Works_referenceOrder[work]; ok {
		return order
	} else {
		log.Printf("instance %p of type Work was not staged and does not have a reference order", work)
		return 0
	}
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (a_directive *A_directive) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_directive.GongGetGongstructName(), a_directive.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_directive *A_directive) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_directive.GongGetGongstructName(), a_directive.GongGetOrder(stage))
}

func (a_measure *A_measure) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_measure.GongGetGongstructName(), a_measure.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_measure *A_measure) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_measure.GongGetGongstructName(), a_measure.GongGetOrder(stage))
}

func (a_measure_1 *A_measure_1) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_measure_1.GongGetGongstructName(), a_measure_1.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_measure_1 *A_measure_1) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_measure_1.GongGetGongstructName(), a_measure_1.GongGetOrder(stage))
}

func (a_part *A_part) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_part.GongGetGongstructName(), a_part.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_part *A_part) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_part.GongGetGongstructName(), a_part.GongGetOrder(stage))
}

func (a_part_1 *A_part_1) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_part_1.GongGetGongstructName(), a_part_1.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_part_1 *A_part_1) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", a_part_1.GongGetGongstructName(), a_part_1.GongGetOrder(stage))
}

func (accidental *Accidental) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", accidental.GongGetGongstructName(), accidental.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (accidental *Accidental) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", accidental.GongGetGongstructName(), accidental.GongGetOrder(stage))
}

func (accidental_mark *Accidental_mark) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", accidental_mark.GongGetGongstructName(), accidental_mark.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (accidental_mark *Accidental_mark) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", accidental_mark.GongGetGongstructName(), accidental_mark.GongGetOrder(stage))
}

func (accidental_text *Accidental_text) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", accidental_text.GongGetGongstructName(), accidental_text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (accidental_text *Accidental_text) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", accidental_text.GongGetGongstructName(), accidental_text.GongGetOrder(stage))
}

func (accord *Accord) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", accord.GongGetGongstructName(), accord.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (accord *Accord) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", accord.GongGetGongstructName(), accord.GongGetOrder(stage))
}

func (accordion_registration *Accordion_registration) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", accordion_registration.GongGetGongstructName(), accordion_registration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (accordion_registration *Accordion_registration) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", accordion_registration.GongGetGongstructName(), accordion_registration.GongGetOrder(stage))
}

func (appearance *Appearance) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", appearance.GongGetGongstructName(), appearance.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (appearance *Appearance) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", appearance.GongGetGongstructName(), appearance.GongGetOrder(stage))
}

func (arpeggiate *Arpeggiate) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", arpeggiate.GongGetGongstructName(), arpeggiate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arpeggiate *Arpeggiate) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", arpeggiate.GongGetGongstructName(), arpeggiate.GongGetOrder(stage))
}

func (arrow *Arrow) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", arrow.GongGetGongstructName(), arrow.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arrow *Arrow) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", arrow.GongGetGongstructName(), arrow.GongGetOrder(stage))
}

func (articulations *Articulations) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", articulations.GongGetGongstructName(), articulations.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (articulations *Articulations) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", articulations.GongGetGongstructName(), articulations.GongGetOrder(stage))
}

func (assess *Assess) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", assess.GongGetGongstructName(), assess.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (assess *Assess) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", assess.GongGetGongstructName(), assess.GongGetOrder(stage))
}

func (attributes *Attributes) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attributes.GongGetGongstructName(), attributes.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attributes *Attributes) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", attributes.GongGetGongstructName(), attributes.GongGetOrder(stage))
}

func (backup *Backup) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", backup.GongGetGongstructName(), backup.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (backup *Backup) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", backup.GongGetGongstructName(), backup.GongGetOrder(stage))
}

func (bar_style_color *Bar_style_color) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bar_style_color.GongGetGongstructName(), bar_style_color.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bar_style_color *Bar_style_color) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bar_style_color.GongGetGongstructName(), bar_style_color.GongGetOrder(stage))
}

func (barline *Barline) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", barline.GongGetGongstructName(), barline.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (barline *Barline) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", barline.GongGetGongstructName(), barline.GongGetOrder(stage))
}

func (barre *Barre) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", barre.GongGetGongstructName(), barre.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (barre *Barre) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", barre.GongGetGongstructName(), barre.GongGetOrder(stage))
}

func (bass *Bass) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bass.GongGetGongstructName(), bass.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bass *Bass) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bass.GongGetGongstructName(), bass.GongGetOrder(stage))
}

func (bass_step *Bass_step) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bass_step.GongGetGongstructName(), bass_step.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bass_step *Bass_step) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bass_step.GongGetGongstructName(), bass_step.GongGetOrder(stage))
}

func (beam *Beam) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", beam.GongGetGongstructName(), beam.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (beam *Beam) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", beam.GongGetGongstructName(), beam.GongGetOrder(stage))
}

func (beat_repeat *Beat_repeat) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", beat_repeat.GongGetGongstructName(), beat_repeat.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (beat_repeat *Beat_repeat) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", beat_repeat.GongGetGongstructName(), beat_repeat.GongGetOrder(stage))
}

func (beat_unit_tied *Beat_unit_tied) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", beat_unit_tied.GongGetGongstructName(), beat_unit_tied.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (beat_unit_tied *Beat_unit_tied) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", beat_unit_tied.GongGetGongstructName(), beat_unit_tied.GongGetOrder(stage))
}

func (beater *Beater) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", beater.GongGetGongstructName(), beater.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (beater *Beater) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", beater.GongGetGongstructName(), beater.GongGetOrder(stage))
}

func (bend *Bend) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bend.GongGetGongstructName(), bend.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bend *Bend) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bend.GongGetGongstructName(), bend.GongGetOrder(stage))
}

func (bookmark *Bookmark) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bookmark.GongGetGongstructName(), bookmark.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bookmark *Bookmark) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bookmark.GongGetGongstructName(), bookmark.GongGetOrder(stage))
}

func (bracket *Bracket) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bracket.GongGetGongstructName(), bracket.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bracket *Bracket) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", bracket.GongGetGongstructName(), bracket.GongGetOrder(stage))
}

func (breath_mark *Breath_mark) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", breath_mark.GongGetGongstructName(), breath_mark.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (breath_mark *Breath_mark) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", breath_mark.GongGetGongstructName(), breath_mark.GongGetOrder(stage))
}

func (caesura *Caesura) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", caesura.GongGetGongstructName(), caesura.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (caesura *Caesura) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", caesura.GongGetGongstructName(), caesura.GongGetOrder(stage))
}

func (cancel *Cancel) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cancel.GongGetGongstructName(), cancel.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cancel *Cancel) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", cancel.GongGetGongstructName(), cancel.GongGetOrder(stage))
}

func (clef *Clef) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", clef.GongGetGongstructName(), clef.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (clef *Clef) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", clef.GongGetGongstructName(), clef.GongGetOrder(stage))
}

func (coda *Coda) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", coda.GongGetGongstructName(), coda.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coda *Coda) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", coda.GongGetGongstructName(), coda.GongGetOrder(stage))
}

func (credit *Credit) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", credit.GongGetGongstructName(), credit.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (credit *Credit) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", credit.GongGetGongstructName(), credit.GongGetOrder(stage))
}

func (dashes *Dashes) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dashes.GongGetGongstructName(), dashes.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dashes *Dashes) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dashes.GongGetGongstructName(), dashes.GongGetOrder(stage))
}

func (defaults *Defaults) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", defaults.GongGetGongstructName(), defaults.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (defaults *Defaults) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", defaults.GongGetGongstructName(), defaults.GongGetOrder(stage))
}

func (degree *Degree) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", degree.GongGetGongstructName(), degree.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (degree *Degree) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", degree.GongGetGongstructName(), degree.GongGetOrder(stage))
}

func (degree_alter *Degree_alter) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", degree_alter.GongGetGongstructName(), degree_alter.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (degree_alter *Degree_alter) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", degree_alter.GongGetGongstructName(), degree_alter.GongGetOrder(stage))
}

func (degree_type *Degree_type) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", degree_type.GongGetGongstructName(), degree_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (degree_type *Degree_type) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", degree_type.GongGetGongstructName(), degree_type.GongGetOrder(stage))
}

func (degree_value *Degree_value) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", degree_value.GongGetGongstructName(), degree_value.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (degree_value *Degree_value) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", degree_value.GongGetGongstructName(), degree_value.GongGetOrder(stage))
}

func (direction *Direction) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", direction.GongGetGongstructName(), direction.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (direction *Direction) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", direction.GongGetGongstructName(), direction.GongGetOrder(stage))
}

func (direction_type *Direction_type) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", direction_type.GongGetGongstructName(), direction_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (direction_type *Direction_type) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", direction_type.GongGetGongstructName(), direction_type.GongGetOrder(stage))
}

func (distance *Distance) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", distance.GongGetGongstructName(), distance.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (distance *Distance) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", distance.GongGetGongstructName(), distance.GongGetOrder(stage))
}

func (double *Double) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", double.GongGetGongstructName(), double.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (double *Double) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", double.GongGetGongstructName(), double.GongGetOrder(stage))
}

func (dynamics *Dynamics) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dynamics.GongGetGongstructName(), dynamics.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dynamics *Dynamics) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", dynamics.GongGetGongstructName(), dynamics.GongGetOrder(stage))
}

func (effect *Effect) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", effect.GongGetGongstructName(), effect.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (effect *Effect) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", effect.GongGetGongstructName(), effect.GongGetOrder(stage))
}

func (elision *Elision) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", elision.GongGetGongstructName(), elision.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elision *Elision) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", elision.GongGetGongstructName(), elision.GongGetOrder(stage))
}

func (empty *Empty) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty.GongGetGongstructName(), empty.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty *Empty) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty.GongGetGongstructName(), empty.GongGetOrder(stage))
}

func (empty_font *Empty_font) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_font.GongGetGongstructName(), empty_font.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_font *Empty_font) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_font.GongGetGongstructName(), empty_font.GongGetOrder(stage))
}

func (empty_line *Empty_line) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_line.GongGetGongstructName(), empty_line.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_line *Empty_line) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_line.GongGetGongstructName(), empty_line.GongGetOrder(stage))
}

func (empty_placement *Empty_placement) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_placement.GongGetGongstructName(), empty_placement.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_placement *Empty_placement) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_placement.GongGetGongstructName(), empty_placement.GongGetOrder(stage))
}

func (empty_placement_smufl *Empty_placement_smufl) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_placement_smufl.GongGetGongstructName(), empty_placement_smufl.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_placement_smufl *Empty_placement_smufl) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_placement_smufl.GongGetGongstructName(), empty_placement_smufl.GongGetOrder(stage))
}

func (empty_print_object_style_align *Empty_print_object_style_align) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_print_object_style_align.GongGetGongstructName(), empty_print_object_style_align.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_print_object_style_align *Empty_print_object_style_align) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_print_object_style_align.GongGetGongstructName(), empty_print_object_style_align.GongGetOrder(stage))
}

func (empty_print_style *Empty_print_style) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_print_style.GongGetGongstructName(), empty_print_style.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_print_style *Empty_print_style) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_print_style.GongGetGongstructName(), empty_print_style.GongGetOrder(stage))
}

func (empty_print_style_align *Empty_print_style_align) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_print_style_align.GongGetGongstructName(), empty_print_style_align.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_print_style_align *Empty_print_style_align) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_print_style_align.GongGetGongstructName(), empty_print_style_align.GongGetOrder(stage))
}

func (empty_print_style_align_id *Empty_print_style_align_id) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_print_style_align_id.GongGetGongstructName(), empty_print_style_align_id.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_print_style_align_id *Empty_print_style_align_id) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_print_style_align_id.GongGetGongstructName(), empty_print_style_align_id.GongGetOrder(stage))
}

func (empty_trill_sound *Empty_trill_sound) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_trill_sound.GongGetGongstructName(), empty_trill_sound.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_trill_sound *Empty_trill_sound) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", empty_trill_sound.GongGetGongstructName(), empty_trill_sound.GongGetOrder(stage))
}

func (encoding *Encoding) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", encoding.GongGetGongstructName(), encoding.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (encoding *Encoding) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", encoding.GongGetGongstructName(), encoding.GongGetOrder(stage))
}

func (ending *Ending) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", ending.GongGetGongstructName(), ending.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ending *Ending) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", ending.GongGetGongstructName(), ending.GongGetOrder(stage))
}

func (extend *Extend) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", extend.GongGetGongstructName(), extend.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (extend *Extend) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", extend.GongGetGongstructName(), extend.GongGetOrder(stage))
}

func (feature *Feature) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", feature.GongGetGongstructName(), feature.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (feature *Feature) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", feature.GongGetGongstructName(), feature.GongGetOrder(stage))
}

func (fermata *Fermata) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", fermata.GongGetGongstructName(), fermata.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (fermata *Fermata) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", fermata.GongGetGongstructName(), fermata.GongGetOrder(stage))
}

func (figure *Figure) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", figure.GongGetGongstructName(), figure.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (figure *Figure) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", figure.GongGetGongstructName(), figure.GongGetOrder(stage))
}

func (figured_bass *Figured_bass) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", figured_bass.GongGetGongstructName(), figured_bass.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (figured_bass *Figured_bass) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", figured_bass.GongGetGongstructName(), figured_bass.GongGetOrder(stage))
}

func (fingering *Fingering) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", fingering.GongGetGongstructName(), fingering.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (fingering *Fingering) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", fingering.GongGetGongstructName(), fingering.GongGetOrder(stage))
}

func (first_fret *First_fret) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", first_fret.GongGetGongstructName(), first_fret.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (first_fret *First_fret) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", first_fret.GongGetGongstructName(), first_fret.GongGetOrder(stage))
}

func (for_part *For_part) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", for_part.GongGetGongstructName(), for_part.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (for_part *For_part) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", for_part.GongGetGongstructName(), for_part.GongGetOrder(stage))
}

func (formatted_symbol *Formatted_symbol) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formatted_symbol.GongGetGongstructName(), formatted_symbol.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formatted_symbol *Formatted_symbol) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formatted_symbol.GongGetGongstructName(), formatted_symbol.GongGetOrder(stage))
}

func (formatted_symbol_id *Formatted_symbol_id) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formatted_symbol_id.GongGetGongstructName(), formatted_symbol_id.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formatted_symbol_id *Formatted_symbol_id) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formatted_symbol_id.GongGetGongstructName(), formatted_symbol_id.GongGetOrder(stage))
}

func (formatted_text *Formatted_text) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formatted_text.GongGetGongstructName(), formatted_text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formatted_text *Formatted_text) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formatted_text.GongGetGongstructName(), formatted_text.GongGetOrder(stage))
}

func (formatted_text_id *Formatted_text_id) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formatted_text_id.GongGetGongstructName(), formatted_text_id.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formatted_text_id *Formatted_text_id) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", formatted_text_id.GongGetGongstructName(), formatted_text_id.GongGetOrder(stage))
}

func (forward *Forward) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", forward.GongGetGongstructName(), forward.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (forward *Forward) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", forward.GongGetGongstructName(), forward.GongGetOrder(stage))
}

func (frame *Frame) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", frame.GongGetGongstructName(), frame.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (frame *Frame) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", frame.GongGetGongstructName(), frame.GongGetOrder(stage))
}

func (frame_note *Frame_note) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", frame_note.GongGetGongstructName(), frame_note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (frame_note *Frame_note) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", frame_note.GongGetGongstructName(), frame_note.GongGetOrder(stage))
}

func (fret *Fret) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", fret.GongGetGongstructName(), fret.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (fret *Fret) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", fret.GongGetGongstructName(), fret.GongGetOrder(stage))
}

func (glass *Glass) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", glass.GongGetGongstructName(), glass.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (glass *Glass) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", glass.GongGetGongstructName(), glass.GongGetOrder(stage))
}

func (glissando *Glissando) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", glissando.GongGetGongstructName(), glissando.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (glissando *Glissando) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", glissando.GongGetGongstructName(), glissando.GongGetOrder(stage))
}

func (glyph *Glyph) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", glyph.GongGetGongstructName(), glyph.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (glyph *Glyph) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", glyph.GongGetGongstructName(), glyph.GongGetOrder(stage))
}

func (grace *Grace) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", grace.GongGetGongstructName(), grace.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (grace *Grace) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", grace.GongGetGongstructName(), grace.GongGetOrder(stage))
}

func (group_barline *Group_barline) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group_barline.GongGetGongstructName(), group_barline.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group_barline *Group_barline) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group_barline.GongGetGongstructName(), group_barline.GongGetOrder(stage))
}

func (group_name *Group_name) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group_name.GongGetGongstructName(), group_name.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group_name *Group_name) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group_name.GongGetGongstructName(), group_name.GongGetOrder(stage))
}

func (group_symbol *Group_symbol) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group_symbol.GongGetGongstructName(), group_symbol.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group_symbol *Group_symbol) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", group_symbol.GongGetGongstructName(), group_symbol.GongGetOrder(stage))
}

func (grouping *Grouping) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", grouping.GongGetGongstructName(), grouping.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (grouping *Grouping) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", grouping.GongGetGongstructName(), grouping.GongGetOrder(stage))
}

func (hammer_on_pull_off *Hammer_on_pull_off) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", hammer_on_pull_off.GongGetGongstructName(), hammer_on_pull_off.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (hammer_on_pull_off *Hammer_on_pull_off) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", hammer_on_pull_off.GongGetGongstructName(), hammer_on_pull_off.GongGetOrder(stage))
}

func (handbell *Handbell) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", handbell.GongGetGongstructName(), handbell.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (handbell *Handbell) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", handbell.GongGetGongstructName(), handbell.GongGetOrder(stage))
}

func (harmon_closed *Harmon_closed) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", harmon_closed.GongGetGongstructName(), harmon_closed.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harmon_closed *Harmon_closed) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", harmon_closed.GongGetGongstructName(), harmon_closed.GongGetOrder(stage))
}

func (harmon_mute *Harmon_mute) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", harmon_mute.GongGetGongstructName(), harmon_mute.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harmon_mute *Harmon_mute) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", harmon_mute.GongGetGongstructName(), harmon_mute.GongGetOrder(stage))
}

func (harmonic *Harmonic) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", harmonic.GongGetGongstructName(), harmonic.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harmonic *Harmonic) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", harmonic.GongGetGongstructName(), harmonic.GongGetOrder(stage))
}

func (harmony *Harmony) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", harmony.GongGetGongstructName(), harmony.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harmony *Harmony) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", harmony.GongGetGongstructName(), harmony.GongGetOrder(stage))
}

func (harmony_alter *Harmony_alter) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", harmony_alter.GongGetGongstructName(), harmony_alter.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harmony_alter *Harmony_alter) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", harmony_alter.GongGetGongstructName(), harmony_alter.GongGetOrder(stage))
}

func (harp_pedals *Harp_pedals) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", harp_pedals.GongGetGongstructName(), harp_pedals.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harp_pedals *Harp_pedals) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", harp_pedals.GongGetGongstructName(), harp_pedals.GongGetOrder(stage))
}

func (heel_toe *Heel_toe) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", heel_toe.GongGetGongstructName(), heel_toe.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (heel_toe *Heel_toe) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", heel_toe.GongGetGongstructName(), heel_toe.GongGetOrder(stage))
}

func (hole *Hole) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", hole.GongGetGongstructName(), hole.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (hole *Hole) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", hole.GongGetGongstructName(), hole.GongGetOrder(stage))
}

func (hole_closed *Hole_closed) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", hole_closed.GongGetGongstructName(), hole_closed.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (hole_closed *Hole_closed) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", hole_closed.GongGetGongstructName(), hole_closed.GongGetOrder(stage))
}

func (horizontal_turn *Horizontal_turn) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", horizontal_turn.GongGetGongstructName(), horizontal_turn.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (horizontal_turn *Horizontal_turn) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", horizontal_turn.GongGetGongstructName(), horizontal_turn.GongGetOrder(stage))
}

func (identification *Identification) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", identification.GongGetGongstructName(), identification.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (identification *Identification) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", identification.GongGetGongstructName(), identification.GongGetOrder(stage))
}

func (image *Image) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", image.GongGetGongstructName(), image.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (image *Image) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", image.GongGetGongstructName(), image.GongGetOrder(stage))
}

func (instrument *Instrument) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", instrument.GongGetGongstructName(), instrument.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (instrument *Instrument) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", instrument.GongGetGongstructName(), instrument.GongGetOrder(stage))
}

func (instrument_change *Instrument_change) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", instrument_change.GongGetGongstructName(), instrument_change.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (instrument_change *Instrument_change) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", instrument_change.GongGetGongstructName(), instrument_change.GongGetOrder(stage))
}

func (instrument_link *Instrument_link) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", instrument_link.GongGetGongstructName(), instrument_link.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (instrument_link *Instrument_link) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", instrument_link.GongGetGongstructName(), instrument_link.GongGetOrder(stage))
}

func (interchangeable *Interchangeable) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", interchangeable.GongGetGongstructName(), interchangeable.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (interchangeable *Interchangeable) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", interchangeable.GongGetGongstructName(), interchangeable.GongGetOrder(stage))
}

func (inversion *Inversion) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", inversion.GongGetGongstructName(), inversion.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (inversion *Inversion) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", inversion.GongGetGongstructName(), inversion.GongGetOrder(stage))
}

func (key *Key) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", key.GongGetGongstructName(), key.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (key *Key) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", key.GongGetGongstructName(), key.GongGetOrder(stage))
}

func (key_accidental *Key_accidental) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", key_accidental.GongGetGongstructName(), key_accidental.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (key_accidental *Key_accidental) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", key_accidental.GongGetGongstructName(), key_accidental.GongGetOrder(stage))
}

func (key_octave *Key_octave) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", key_octave.GongGetGongstructName(), key_octave.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (key_octave *Key_octave) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", key_octave.GongGetGongstructName(), key_octave.GongGetOrder(stage))
}

func (kind *Kind) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", kind.GongGetGongstructName(), kind.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (kind *Kind) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", kind.GongGetGongstructName(), kind.GongGetOrder(stage))
}

func (level *Level) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", level.GongGetGongstructName(), level.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (level *Level) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", level.GongGetGongstructName(), level.GongGetOrder(stage))
}

func (line_detail *Line_detail) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", line_detail.GongGetGongstructName(), line_detail.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (line_detail *Line_detail) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", line_detail.GongGetGongstructName(), line_detail.GongGetOrder(stage))
}

func (line_width *Line_width) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", line_width.GongGetGongstructName(), line_width.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (line_width *Line_width) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", line_width.GongGetGongstructName(), line_width.GongGetOrder(stage))
}

func (link *Link) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", link.GongGetGongstructName(), link.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (link *Link) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", link.GongGetGongstructName(), link.GongGetOrder(stage))
}

func (listen *Listen) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", listen.GongGetGongstructName(), listen.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (listen *Listen) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", listen.GongGetGongstructName(), listen.GongGetOrder(stage))
}

func (listening *Listening) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", listening.GongGetGongstructName(), listening.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (listening *Listening) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", listening.GongGetGongstructName(), listening.GongGetOrder(stage))
}

func (lyric *Lyric) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", lyric.GongGetGongstructName(), lyric.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lyric *Lyric) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", lyric.GongGetGongstructName(), lyric.GongGetOrder(stage))
}

func (lyric_font *Lyric_font) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", lyric_font.GongGetGongstructName(), lyric_font.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lyric_font *Lyric_font) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", lyric_font.GongGetGongstructName(), lyric_font.GongGetOrder(stage))
}

func (lyric_language *Lyric_language) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", lyric_language.GongGetGongstructName(), lyric_language.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lyric_language *Lyric_language) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", lyric_language.GongGetGongstructName(), lyric_language.GongGetOrder(stage))
}

func (measure_layout *Measure_layout) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", measure_layout.GongGetGongstructName(), measure_layout.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measure_layout *Measure_layout) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", measure_layout.GongGetGongstructName(), measure_layout.GongGetOrder(stage))
}

func (measure_numbering *Measure_numbering) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", measure_numbering.GongGetGongstructName(), measure_numbering.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measure_numbering *Measure_numbering) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", measure_numbering.GongGetGongstructName(), measure_numbering.GongGetOrder(stage))
}

func (measure_repeat *Measure_repeat) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", measure_repeat.GongGetGongstructName(), measure_repeat.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measure_repeat *Measure_repeat) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", measure_repeat.GongGetGongstructName(), measure_repeat.GongGetOrder(stage))
}

func (measure_style *Measure_style) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", measure_style.GongGetGongstructName(), measure_style.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measure_style *Measure_style) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", measure_style.GongGetGongstructName(), measure_style.GongGetOrder(stage))
}

func (membrane *Membrane) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", membrane.GongGetGongstructName(), membrane.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (membrane *Membrane) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", membrane.GongGetGongstructName(), membrane.GongGetOrder(stage))
}

func (metal *Metal) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", metal.GongGetGongstructName(), metal.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metal *Metal) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", metal.GongGetGongstructName(), metal.GongGetOrder(stage))
}

func (metronome *Metronome) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", metronome.GongGetGongstructName(), metronome.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metronome *Metronome) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", metronome.GongGetGongstructName(), metronome.GongGetOrder(stage))
}

func (metronome_beam *Metronome_beam) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", metronome_beam.GongGetGongstructName(), metronome_beam.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metronome_beam *Metronome_beam) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", metronome_beam.GongGetGongstructName(), metronome_beam.GongGetOrder(stage))
}

func (metronome_note *Metronome_note) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", metronome_note.GongGetGongstructName(), metronome_note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metronome_note *Metronome_note) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", metronome_note.GongGetGongstructName(), metronome_note.GongGetOrder(stage))
}

func (metronome_tied *Metronome_tied) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", metronome_tied.GongGetGongstructName(), metronome_tied.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metronome_tied *Metronome_tied) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", metronome_tied.GongGetGongstructName(), metronome_tied.GongGetOrder(stage))
}

func (metronome_tuplet *Metronome_tuplet) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", metronome_tuplet.GongGetGongstructName(), metronome_tuplet.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metronome_tuplet *Metronome_tuplet) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", metronome_tuplet.GongGetGongstructName(), metronome_tuplet.GongGetOrder(stage))
}

func (midi_device *Midi_device) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", midi_device.GongGetGongstructName(), midi_device.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (midi_device *Midi_device) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", midi_device.GongGetGongstructName(), midi_device.GongGetOrder(stage))
}

func (midi_instrument *Midi_instrument) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", midi_instrument.GongGetGongstructName(), midi_instrument.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (midi_instrument *Midi_instrument) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", midi_instrument.GongGetGongstructName(), midi_instrument.GongGetOrder(stage))
}

func (miscellaneous *Miscellaneous) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", miscellaneous.GongGetGongstructName(), miscellaneous.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (miscellaneous *Miscellaneous) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", miscellaneous.GongGetGongstructName(), miscellaneous.GongGetOrder(stage))
}

func (miscellaneous_field *Miscellaneous_field) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", miscellaneous_field.GongGetGongstructName(), miscellaneous_field.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (miscellaneous_field *Miscellaneous_field) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", miscellaneous_field.GongGetGongstructName(), miscellaneous_field.GongGetOrder(stage))
}

func (mordent *Mordent) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", mordent.GongGetGongstructName(), mordent.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (mordent *Mordent) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", mordent.GongGetGongstructName(), mordent.GongGetOrder(stage))
}

func (multiple_rest *Multiple_rest) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", multiple_rest.GongGetGongstructName(), multiple_rest.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (multiple_rest *Multiple_rest) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", multiple_rest.GongGetGongstructName(), multiple_rest.GongGetOrder(stage))
}

func (name_display *Name_display) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", name_display.GongGetGongstructName(), name_display.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (name_display *Name_display) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", name_display.GongGetGongstructName(), name_display.GongGetOrder(stage))
}

func (non_arpeggiate *Non_arpeggiate) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", non_arpeggiate.GongGetGongstructName(), non_arpeggiate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (non_arpeggiate *Non_arpeggiate) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", non_arpeggiate.GongGetGongstructName(), non_arpeggiate.GongGetOrder(stage))
}

func (notations *Notations) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notations.GongGetGongstructName(), notations.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notations *Notations) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notations.GongGetGongstructName(), notations.GongGetOrder(stage))
}

func (note *Note) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note.GongGetGongstructName(), note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note *Note) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note.GongGetGongstructName(), note.GongGetOrder(stage))
}

func (note_size *Note_size) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note_size.GongGetGongstructName(), note_size.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note_size *Note_size) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note_size.GongGetGongstructName(), note_size.GongGetOrder(stage))
}

func (note_type *Note_type) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note_type.GongGetGongstructName(), note_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note_type *Note_type) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", note_type.GongGetGongstructName(), note_type.GongGetOrder(stage))
}

func (notehead *Notehead) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notehead.GongGetGongstructName(), notehead.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notehead *Notehead) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notehead.GongGetGongstructName(), notehead.GongGetOrder(stage))
}

func (notehead_text *Notehead_text) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notehead_text.GongGetGongstructName(), notehead_text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notehead_text *Notehead_text) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", notehead_text.GongGetGongstructName(), notehead_text.GongGetOrder(stage))
}

func (numeral *Numeral) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", numeral.GongGetGongstructName(), numeral.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (numeral *Numeral) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", numeral.GongGetGongstructName(), numeral.GongGetOrder(stage))
}

func (numeral_key *Numeral_key) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", numeral_key.GongGetGongstructName(), numeral_key.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (numeral_key *Numeral_key) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", numeral_key.GongGetGongstructName(), numeral_key.GongGetOrder(stage))
}

func (numeral_root *Numeral_root) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", numeral_root.GongGetGongstructName(), numeral_root.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (numeral_root *Numeral_root) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", numeral_root.GongGetGongstructName(), numeral_root.GongGetOrder(stage))
}

func (octave_shift *Octave_shift) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", octave_shift.GongGetGongstructName(), octave_shift.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (octave_shift *Octave_shift) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", octave_shift.GongGetGongstructName(), octave_shift.GongGetOrder(stage))
}

func (offset *Offset) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", offset.GongGetGongstructName(), offset.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (offset *Offset) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", offset.GongGetGongstructName(), offset.GongGetOrder(stage))
}

func (opus *Opus) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", opus.GongGetGongstructName(), opus.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (opus *Opus) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", opus.GongGetGongstructName(), opus.GongGetOrder(stage))
}

func (ornaments *Ornaments) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", ornaments.GongGetGongstructName(), ornaments.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ornaments *Ornaments) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", ornaments.GongGetGongstructName(), ornaments.GongGetOrder(stage))
}

func (other_appearance *Other_appearance) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", other_appearance.GongGetGongstructName(), other_appearance.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (other_appearance *Other_appearance) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", other_appearance.GongGetGongstructName(), other_appearance.GongGetOrder(stage))
}

func (other_direction *Other_direction) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", other_direction.GongGetGongstructName(), other_direction.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (other_direction *Other_direction) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", other_direction.GongGetGongstructName(), other_direction.GongGetOrder(stage))
}

func (other_listening *Other_listening) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", other_listening.GongGetGongstructName(), other_listening.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (other_listening *Other_listening) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", other_listening.GongGetGongstructName(), other_listening.GongGetOrder(stage))
}

func (other_notation *Other_notation) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", other_notation.GongGetGongstructName(), other_notation.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (other_notation *Other_notation) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", other_notation.GongGetGongstructName(), other_notation.GongGetOrder(stage))
}

func (other_placement_text *Other_placement_text) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", other_placement_text.GongGetGongstructName(), other_placement_text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (other_placement_text *Other_placement_text) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", other_placement_text.GongGetGongstructName(), other_placement_text.GongGetOrder(stage))
}

func (other_play *Other_play) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", other_play.GongGetGongstructName(), other_play.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (other_play *Other_play) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", other_play.GongGetGongstructName(), other_play.GongGetOrder(stage))
}

func (other_text *Other_text) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", other_text.GongGetGongstructName(), other_text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (other_text *Other_text) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", other_text.GongGetGongstructName(), other_text.GongGetOrder(stage))
}

func (page_layout *Page_layout) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", page_layout.GongGetGongstructName(), page_layout.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (page_layout *Page_layout) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", page_layout.GongGetGongstructName(), page_layout.GongGetOrder(stage))
}

func (page_margins *Page_margins) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", page_margins.GongGetGongstructName(), page_margins.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (page_margins *Page_margins) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", page_margins.GongGetGongstructName(), page_margins.GongGetOrder(stage))
}

func (part_clef *Part_clef) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part_clef.GongGetGongstructName(), part_clef.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part_clef *Part_clef) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part_clef.GongGetGongstructName(), part_clef.GongGetOrder(stage))
}

func (part_group *Part_group) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part_group.GongGetGongstructName(), part_group.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part_group *Part_group) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part_group.GongGetGongstructName(), part_group.GongGetOrder(stage))
}

func (part_link *Part_link) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part_link.GongGetGongstructName(), part_link.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part_link *Part_link) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part_link.GongGetGongstructName(), part_link.GongGetOrder(stage))
}

func (part_list *Part_list) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part_list.GongGetGongstructName(), part_list.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part_list *Part_list) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part_list.GongGetGongstructName(), part_list.GongGetOrder(stage))
}

func (part_name *Part_name) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part_name.GongGetGongstructName(), part_name.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part_name *Part_name) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part_name.GongGetGongstructName(), part_name.GongGetOrder(stage))
}

func (part_symbol *Part_symbol) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part_symbol.GongGetGongstructName(), part_symbol.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part_symbol *Part_symbol) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part_symbol.GongGetGongstructName(), part_symbol.GongGetOrder(stage))
}

func (part_transpose *Part_transpose) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part_transpose.GongGetGongstructName(), part_transpose.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part_transpose *Part_transpose) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", part_transpose.GongGetGongstructName(), part_transpose.GongGetOrder(stage))
}

func (pedal *Pedal) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pedal.GongGetGongstructName(), pedal.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pedal *Pedal) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pedal.GongGetGongstructName(), pedal.GongGetOrder(stage))
}

func (pedal_tuning *Pedal_tuning) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pedal_tuning.GongGetGongstructName(), pedal_tuning.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pedal_tuning *Pedal_tuning) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pedal_tuning.GongGetGongstructName(), pedal_tuning.GongGetOrder(stage))
}

func (per_minute *Per_minute) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", per_minute.GongGetGongstructName(), per_minute.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (per_minute *Per_minute) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", per_minute.GongGetGongstructName(), per_minute.GongGetOrder(stage))
}

func (percussion *Percussion) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", percussion.GongGetGongstructName(), percussion.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (percussion *Percussion) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", percussion.GongGetGongstructName(), percussion.GongGetOrder(stage))
}

func (pitch *Pitch) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pitch.GongGetGongstructName(), pitch.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pitch *Pitch) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pitch.GongGetGongstructName(), pitch.GongGetOrder(stage))
}

func (pitched *Pitched) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pitched.GongGetGongstructName(), pitched.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pitched *Pitched) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", pitched.GongGetGongstructName(), pitched.GongGetOrder(stage))
}

func (placement_text *Placement_text) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", placement_text.GongGetGongstructName(), placement_text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (placement_text *Placement_text) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", placement_text.GongGetGongstructName(), placement_text.GongGetOrder(stage))
}

func (play *Play) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", play.GongGetGongstructName(), play.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (play *Play) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", play.GongGetGongstructName(), play.GongGetOrder(stage))
}

func (player *Player) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", player.GongGetGongstructName(), player.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (player *Player) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", player.GongGetGongstructName(), player.GongGetOrder(stage))
}

func (principal_voice *Principal_voice) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", principal_voice.GongGetGongstructName(), principal_voice.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (principal_voice *Principal_voice) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", principal_voice.GongGetGongstructName(), principal_voice.GongGetOrder(stage))
}

func (print *Print) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", print.GongGetGongstructName(), print.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (print *Print) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", print.GongGetGongstructName(), print.GongGetOrder(stage))
}

func (release *Release) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", release.GongGetGongstructName(), release.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (release *Release) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", release.GongGetGongstructName(), release.GongGetOrder(stage))
}

func (repeat *Repeat) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", repeat.GongGetGongstructName(), repeat.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (repeat *Repeat) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", repeat.GongGetGongstructName(), repeat.GongGetOrder(stage))
}

func (rest *Rest) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rest.GongGetGongstructName(), rest.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rest *Rest) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", rest.GongGetGongstructName(), rest.GongGetOrder(stage))
}

func (root *Root) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", root.GongGetGongstructName(), root.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (root *Root) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", root.GongGetGongstructName(), root.GongGetOrder(stage))
}

func (root_step *Root_step) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", root_step.GongGetGongstructName(), root_step.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (root_step *Root_step) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", root_step.GongGetGongstructName(), root_step.GongGetOrder(stage))
}

func (scaling *Scaling) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", scaling.GongGetGongstructName(), scaling.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (scaling *Scaling) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", scaling.GongGetGongstructName(), scaling.GongGetOrder(stage))
}

func (scordatura *Scordatura) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", scordatura.GongGetGongstructName(), scordatura.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (scordatura *Scordatura) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", scordatura.GongGetGongstructName(), scordatura.GongGetOrder(stage))
}

func (score_instrument *Score_instrument) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", score_instrument.GongGetGongstructName(), score_instrument.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (score_instrument *Score_instrument) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", score_instrument.GongGetGongstructName(), score_instrument.GongGetOrder(stage))
}

func (score_part *Score_part) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", score_part.GongGetGongstructName(), score_part.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (score_part *Score_part) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", score_part.GongGetGongstructName(), score_part.GongGetOrder(stage))
}

func (score_partwise *Score_partwise) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", score_partwise.GongGetGongstructName(), score_partwise.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (score_partwise *Score_partwise) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", score_partwise.GongGetGongstructName(), score_partwise.GongGetOrder(stage))
}

func (score_timewise *Score_timewise) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", score_timewise.GongGetGongstructName(), score_timewise.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (score_timewise *Score_timewise) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", score_timewise.GongGetGongstructName(), score_timewise.GongGetOrder(stage))
}

func (segno *Segno) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", segno.GongGetGongstructName(), segno.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (segno *Segno) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", segno.GongGetGongstructName(), segno.GongGetOrder(stage))
}

func (slash *Slash) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", slash.GongGetGongstructName(), slash.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (slash *Slash) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", slash.GongGetGongstructName(), slash.GongGetOrder(stage))
}

func (slide *Slide) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", slide.GongGetGongstructName(), slide.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (slide *Slide) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", slide.GongGetGongstructName(), slide.GongGetOrder(stage))
}

func (slur *Slur) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", slur.GongGetGongstructName(), slur.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (slur *Slur) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", slur.GongGetGongstructName(), slur.GongGetOrder(stage))
}

func (sound *Sound) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", sound.GongGetGongstructName(), sound.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sound *Sound) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", sound.GongGetGongstructName(), sound.GongGetOrder(stage))
}

func (staff_details *Staff_details) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staff_details.GongGetGongstructName(), staff_details.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staff_details *Staff_details) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staff_details.GongGetGongstructName(), staff_details.GongGetOrder(stage))
}

func (staff_divide *Staff_divide) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staff_divide.GongGetGongstructName(), staff_divide.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staff_divide *Staff_divide) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staff_divide.GongGetGongstructName(), staff_divide.GongGetOrder(stage))
}

func (staff_layout *Staff_layout) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staff_layout.GongGetGongstructName(), staff_layout.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staff_layout *Staff_layout) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staff_layout.GongGetGongstructName(), staff_layout.GongGetOrder(stage))
}

func (staff_size *Staff_size) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staff_size.GongGetGongstructName(), staff_size.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staff_size *Staff_size) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staff_size.GongGetGongstructName(), staff_size.GongGetOrder(stage))
}

func (staff_tuning *Staff_tuning) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staff_tuning.GongGetGongstructName(), staff_tuning.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staff_tuning *Staff_tuning) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", staff_tuning.GongGetGongstructName(), staff_tuning.GongGetOrder(stage))
}

func (stem *Stem) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stem.GongGetGongstructName(), stem.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stem *Stem) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stem.GongGetGongstructName(), stem.GongGetOrder(stage))
}

func (stick *Stick) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stick.GongGetGongstructName(), stick.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stick *Stick) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", stick.GongGetGongstructName(), stick.GongGetOrder(stage))
}

func (string_mute *String_mute) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", string_mute.GongGetGongstructName(), string_mute.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (string_mute *String_mute) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", string_mute.GongGetGongstructName(), string_mute.GongGetOrder(stage))
}

func (string_type *String_type) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", string_type.GongGetGongstructName(), string_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (string_type *String_type) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", string_type.GongGetGongstructName(), string_type.GongGetOrder(stage))
}

func (strong_accent *Strong_accent) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", strong_accent.GongGetGongstructName(), strong_accent.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (strong_accent *Strong_accent) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", strong_accent.GongGetGongstructName(), strong_accent.GongGetOrder(stage))
}

func (style_text *Style_text) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", style_text.GongGetGongstructName(), style_text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (style_text *Style_text) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", style_text.GongGetGongstructName(), style_text.GongGetOrder(stage))
}

func (supports *Supports) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", supports.GongGetGongstructName(), supports.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (supports *Supports) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", supports.GongGetGongstructName(), supports.GongGetOrder(stage))
}

func (swing *Swing) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", swing.GongGetGongstructName(), swing.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (swing *Swing) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", swing.GongGetGongstructName(), swing.GongGetOrder(stage))
}

func (sync *Sync) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", sync.GongGetGongstructName(), sync.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sync *Sync) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", sync.GongGetGongstructName(), sync.GongGetOrder(stage))
}

func (system_dividers *System_dividers) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", system_dividers.GongGetGongstructName(), system_dividers.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (system_dividers *System_dividers) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", system_dividers.GongGetGongstructName(), system_dividers.GongGetOrder(stage))
}

func (system_layout *System_layout) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", system_layout.GongGetGongstructName(), system_layout.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (system_layout *System_layout) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", system_layout.GongGetGongstructName(), system_layout.GongGetOrder(stage))
}

func (system_margins *System_margins) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", system_margins.GongGetGongstructName(), system_margins.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (system_margins *System_margins) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", system_margins.GongGetGongstructName(), system_margins.GongGetOrder(stage))
}

func (tap *Tap) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tap.GongGetGongstructName(), tap.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tap *Tap) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tap.GongGetGongstructName(), tap.GongGetOrder(stage))
}

func (technical *Technical) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", technical.GongGetGongstructName(), technical.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (technical *Technical) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", technical.GongGetGongstructName(), technical.GongGetOrder(stage))
}

func (text_element_data *Text_element_data) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", text_element_data.GongGetGongstructName(), text_element_data.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (text_element_data *Text_element_data) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", text_element_data.GongGetGongstructName(), text_element_data.GongGetOrder(stage))
}

func (tie *Tie) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tie.GongGetGongstructName(), tie.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tie *Tie) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tie.GongGetGongstructName(), tie.GongGetOrder(stage))
}

func (tied *Tied) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tied.GongGetGongstructName(), tied.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tied *Tied) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tied.GongGetGongstructName(), tied.GongGetOrder(stage))
}

func (time *Time) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", time.GongGetGongstructName(), time.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (time *Time) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", time.GongGetGongstructName(), time.GongGetOrder(stage))
}

func (time_modification *Time_modification) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", time_modification.GongGetGongstructName(), time_modification.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (time_modification *Time_modification) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", time_modification.GongGetGongstructName(), time_modification.GongGetOrder(stage))
}

func (timpani *Timpani) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", timpani.GongGetGongstructName(), timpani.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (timpani *Timpani) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", timpani.GongGetGongstructName(), timpani.GongGetOrder(stage))
}

func (transpose *Transpose) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", transpose.GongGetGongstructName(), transpose.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (transpose *Transpose) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", transpose.GongGetGongstructName(), transpose.GongGetOrder(stage))
}

func (tremolo *Tremolo) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tremolo.GongGetGongstructName(), tremolo.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tremolo *Tremolo) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tremolo.GongGetGongstructName(), tremolo.GongGetOrder(stage))
}

func (tuplet *Tuplet) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tuplet.GongGetGongstructName(), tuplet.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tuplet *Tuplet) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tuplet.GongGetGongstructName(), tuplet.GongGetOrder(stage))
}

func (tuplet_dot *Tuplet_dot) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tuplet_dot.GongGetGongstructName(), tuplet_dot.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tuplet_dot *Tuplet_dot) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tuplet_dot.GongGetGongstructName(), tuplet_dot.GongGetOrder(stage))
}

func (tuplet_number *Tuplet_number) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tuplet_number.GongGetGongstructName(), tuplet_number.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tuplet_number *Tuplet_number) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tuplet_number.GongGetGongstructName(), tuplet_number.GongGetOrder(stage))
}

func (tuplet_portion *Tuplet_portion) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tuplet_portion.GongGetGongstructName(), tuplet_portion.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tuplet_portion *Tuplet_portion) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tuplet_portion.GongGetGongstructName(), tuplet_portion.GongGetOrder(stage))
}

func (tuplet_type *Tuplet_type) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tuplet_type.GongGetGongstructName(), tuplet_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tuplet_type *Tuplet_type) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", tuplet_type.GongGetGongstructName(), tuplet_type.GongGetOrder(stage))
}

func (typed_text *Typed_text) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", typed_text.GongGetGongstructName(), typed_text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (typed_text *Typed_text) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", typed_text.GongGetGongstructName(), typed_text.GongGetOrder(stage))
}

func (unpitched *Unpitched) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", unpitched.GongGetGongstructName(), unpitched.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unpitched *Unpitched) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", unpitched.GongGetGongstructName(), unpitched.GongGetOrder(stage))
}

func (virtual_instrument *Virtual_instrument) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", virtual_instrument.GongGetGongstructName(), virtual_instrument.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (virtual_instrument *Virtual_instrument) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", virtual_instrument.GongGetGongstructName(), virtual_instrument.GongGetOrder(stage))
}

func (wait *Wait) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", wait.GongGetGongstructName(), wait.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (wait *Wait) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", wait.GongGetGongstructName(), wait.GongGetOrder(stage))
}

func (wavy_line *Wavy_line) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", wavy_line.GongGetGongstructName(), wavy_line.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (wavy_line *Wavy_line) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", wavy_line.GongGetGongstructName(), wavy_line.GongGetOrder(stage))
}

func (wedge *Wedge) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", wedge.GongGetGongstructName(), wedge.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (wedge *Wedge) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", wedge.GongGetGongstructName(), wedge.GongGetOrder(stage))
}

func (wood *Wood) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", wood.GongGetGongstructName(), wood.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (wood *Wood) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", wood.GongGetGongstructName(), wood.GongGetOrder(stage))
}

func (work *Work) GongGetIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", work.GongGetGongstructName(), work.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (work *Work) GongGetReferenceIdentifier(stage *Stage) string {
	return fmt.Sprintf("__%s__%08d_", work.GongGetGongstructName(), work.GongGetOrder(stage))
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (a_directive *A_directive) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_directive.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_directive")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_directive.Name))
	return
}

func (a_measure *A_measure) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_measure.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_measure")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_measure.Name))
	return
}

func (a_measure_1 *A_measure_1) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_measure_1.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_measure_1")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_measure_1.Name))
	return
}

func (a_part *A_part) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_part.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_part")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_part.Name))
	return
}

func (a_part_1 *A_part_1) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_part_1.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "A_part_1")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(a_part_1.Name))
	return
}

func (accidental *Accidental) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", accidental.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Accidental")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(accidental.Name))
	return
}

func (accidental_mark *Accidental_mark) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", accidental_mark.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Accidental_mark")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(accidental_mark.Name))
	return
}

func (accidental_text *Accidental_text) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", accidental_text.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Accidental_text")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(accidental_text.Name))
	return
}

func (accord *Accord) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", accord.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Accord")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(accord.Name))
	return
}

func (accordion_registration *Accordion_registration) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", accordion_registration.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Accordion_registration")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(accordion_registration.Name))
	return
}

func (appearance *Appearance) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", appearance.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Appearance")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(appearance.Name))
	return
}

func (arpeggiate *Arpeggiate) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", arpeggiate.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Arpeggiate")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(arpeggiate.Name))
	return
}

func (arrow *Arrow) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", arrow.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Arrow")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(arrow.Name))
	return
}

func (articulations *Articulations) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", articulations.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Articulations")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(articulations.Name))
	return
}

func (assess *Assess) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", assess.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Assess")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(assess.Name))
	return
}

func (attributes *Attributes) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attributes.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Attributes")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(attributes.Name))
	return
}

func (backup *Backup) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", backup.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Backup")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(backup.Name))
	return
}

func (bar_style_color *Bar_style_color) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bar_style_color.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bar_style_color")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(bar_style_color.Name))
	return
}

func (barline *Barline) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", barline.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Barline")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(barline.Name))
	return
}

func (barre *Barre) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", barre.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Barre")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(barre.Name))
	return
}

func (bass *Bass) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bass.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bass")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(bass.Name))
	return
}

func (bass_step *Bass_step) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bass_step.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bass_step")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(bass_step.Name))
	return
}

func (beam *Beam) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", beam.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Beam")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(beam.Name))
	return
}

func (beat_repeat *Beat_repeat) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", beat_repeat.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Beat_repeat")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(beat_repeat.Name))
	return
}

func (beat_unit_tied *Beat_unit_tied) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", beat_unit_tied.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Beat_unit_tied")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(beat_unit_tied.Name))
	return
}

func (beater *Beater) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", beater.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Beater")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(beater.Name))
	return
}

func (bend *Bend) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bend.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bend")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(bend.Name))
	return
}

func (bookmark *Bookmark) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bookmark.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bookmark")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(bookmark.Name))
	return
}

func (bracket *Bracket) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bracket.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Bracket")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(bracket.Name))
	return
}

func (breath_mark *Breath_mark) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", breath_mark.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Breath_mark")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(breath_mark.Name))
	return
}

func (caesura *Caesura) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", caesura.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Caesura")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(caesura.Name))
	return
}

func (cancel *Cancel) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cancel.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Cancel")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(cancel.Name))
	return
}

func (clef *Clef) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", clef.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Clef")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(clef.Name))
	return
}

func (coda *Coda) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", coda.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Coda")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(coda.Name))
	return
}

func (credit *Credit) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", credit.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Credit")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(credit.Name))
	return
}

func (dashes *Dashes) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dashes.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Dashes")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(dashes.Name))
	return
}

func (defaults *Defaults) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", defaults.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Defaults")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(defaults.Name))
	return
}

func (degree *Degree) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", degree.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Degree")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(degree.Name))
	return
}

func (degree_alter *Degree_alter) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", degree_alter.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Degree_alter")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(degree_alter.Name))
	return
}

func (degree_type *Degree_type) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", degree_type.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Degree_type")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(degree_type.Name))
	return
}

func (degree_value *Degree_value) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", degree_value.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Degree_value")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(degree_value.Name))
	return
}

func (direction *Direction) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", direction.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Direction")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(direction.Name))
	return
}

func (direction_type *Direction_type) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", direction_type.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Direction_type")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(direction_type.Name))
	return
}

func (distance *Distance) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", distance.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Distance")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(distance.Name))
	return
}

func (double *Double) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", double.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Double")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(double.Name))
	return
}

func (dynamics *Dynamics) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dynamics.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Dynamics")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(dynamics.Name))
	return
}

func (effect *Effect) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", effect.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Effect")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(effect.Name))
	return
}

func (elision *Elision) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", elision.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Elision")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(elision.Name))
	return
}

func (empty *Empty) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(empty.Name))
	return
}

func (empty_font *Empty_font) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_font.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_font")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(empty_font.Name))
	return
}

func (empty_line *Empty_line) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_line.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_line")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(empty_line.Name))
	return
}

func (empty_placement *Empty_placement) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_placement.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_placement")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(empty_placement.Name))
	return
}

func (empty_placement_smufl *Empty_placement_smufl) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_placement_smufl.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_placement_smufl")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(empty_placement_smufl.Name))
	return
}

func (empty_print_object_style_align *Empty_print_object_style_align) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_print_object_style_align.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_print_object_style_align")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(empty_print_object_style_align.Name))
	return
}

func (empty_print_style *Empty_print_style) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_print_style.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_print_style")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(empty_print_style.Name))
	return
}

func (empty_print_style_align *Empty_print_style_align) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_print_style_align.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_print_style_align")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(empty_print_style_align.Name))
	return
}

func (empty_print_style_align_id *Empty_print_style_align_id) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_print_style_align_id.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_print_style_align_id")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(empty_print_style_align_id.Name))
	return
}

func (empty_trill_sound *Empty_trill_sound) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_trill_sound.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Empty_trill_sound")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(empty_trill_sound.Name))
	return
}

func (encoding *Encoding) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", encoding.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Encoding")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(encoding.Name))
	return
}

func (ending *Ending) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", ending.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Ending")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(ending.Name))
	return
}

func (extend *Extend) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", extend.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Extend")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(extend.Name))
	return
}

func (feature *Feature) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", feature.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Feature")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(feature.Name))
	return
}

func (fermata *Fermata) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", fermata.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Fermata")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(fermata.Name))
	return
}

func (figure *Figure) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", figure.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Figure")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(figure.Name))
	return
}

func (figured_bass *Figured_bass) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", figured_bass.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Figured_bass")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(figured_bass.Name))
	return
}

func (fingering *Fingering) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", fingering.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Fingering")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(fingering.Name))
	return
}

func (first_fret *First_fret) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", first_fret.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "First_fret")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(first_fret.Name))
	return
}

func (for_part *For_part) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", for_part.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "For_part")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(for_part.Name))
	return
}

func (formatted_symbol *Formatted_symbol) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formatted_symbol.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Formatted_symbol")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(formatted_symbol.Name))
	return
}

func (formatted_symbol_id *Formatted_symbol_id) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formatted_symbol_id.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Formatted_symbol_id")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(formatted_symbol_id.Name))
	return
}

func (formatted_text *Formatted_text) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formatted_text.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Formatted_text")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(formatted_text.Name))
	return
}

func (formatted_text_id *Formatted_text_id) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formatted_text_id.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Formatted_text_id")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(formatted_text_id.Name))
	return
}

func (forward *Forward) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", forward.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Forward")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(forward.Name))
	return
}

func (frame *Frame) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", frame.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Frame")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(frame.Name))
	return
}

func (frame_note *Frame_note) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", frame_note.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Frame_note")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(frame_note.Name))
	return
}

func (fret *Fret) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", fret.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Fret")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(fret.Name))
	return
}

func (glass *Glass) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", glass.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Glass")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(glass.Name))
	return
}

func (glissando *Glissando) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", glissando.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Glissando")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(glissando.Name))
	return
}

func (glyph *Glyph) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", glyph.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Glyph")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(glyph.Name))
	return
}

func (grace *Grace) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", grace.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Grace")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(grace.Name))
	return
}

func (group_barline *Group_barline) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group_barline.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group_barline")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(group_barline.Name))
	return
}

func (group_name *Group_name) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group_name.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group_name")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(group_name.Name))
	return
}

func (group_symbol *Group_symbol) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group_symbol.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Group_symbol")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(group_symbol.Name))
	return
}

func (grouping *Grouping) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", grouping.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Grouping")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(grouping.Name))
	return
}

func (hammer_on_pull_off *Hammer_on_pull_off) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", hammer_on_pull_off.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Hammer_on_pull_off")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(hammer_on_pull_off.Name))
	return
}

func (handbell *Handbell) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", handbell.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Handbell")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(handbell.Name))
	return
}

func (harmon_closed *Harmon_closed) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", harmon_closed.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Harmon_closed")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(harmon_closed.Name))
	return
}

func (harmon_mute *Harmon_mute) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", harmon_mute.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Harmon_mute")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(harmon_mute.Name))
	return
}

func (harmonic *Harmonic) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", harmonic.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Harmonic")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(harmonic.Name))
	return
}

func (harmony *Harmony) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", harmony.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Harmony")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(harmony.Name))
	return
}

func (harmony_alter *Harmony_alter) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", harmony_alter.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Harmony_alter")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(harmony_alter.Name))
	return
}

func (harp_pedals *Harp_pedals) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", harp_pedals.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Harp_pedals")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(harp_pedals.Name))
	return
}

func (heel_toe *Heel_toe) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", heel_toe.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Heel_toe")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(heel_toe.Name))
	return
}

func (hole *Hole) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", hole.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Hole")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(hole.Name))
	return
}

func (hole_closed *Hole_closed) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", hole_closed.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Hole_closed")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(hole_closed.Name))
	return
}

func (horizontal_turn *Horizontal_turn) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", horizontal_turn.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Horizontal_turn")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(horizontal_turn.Name))
	return
}

func (identification *Identification) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", identification.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Identification")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(identification.Name))
	return
}

func (image *Image) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", image.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Image")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(image.Name))
	return
}

func (instrument *Instrument) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", instrument.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Instrument")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(instrument.Name))
	return
}

func (instrument_change *Instrument_change) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", instrument_change.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Instrument_change")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(instrument_change.Name))
	return
}

func (instrument_link *Instrument_link) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", instrument_link.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Instrument_link")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(instrument_link.Name))
	return
}

func (interchangeable *Interchangeable) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", interchangeable.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Interchangeable")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(interchangeable.Name))
	return
}

func (inversion *Inversion) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", inversion.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Inversion")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(inversion.Name))
	return
}

func (key *Key) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", key.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Key")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(key.Name))
	return
}

func (key_accidental *Key_accidental) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", key_accidental.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Key_accidental")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(key_accidental.Name))
	return
}

func (key_octave *Key_octave) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", key_octave.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Key_octave")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(key_octave.Name))
	return
}

func (kind *Kind) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", kind.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Kind")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(kind.Name))
	return
}

func (level *Level) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", level.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Level")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(level.Name))
	return
}

func (line_detail *Line_detail) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", line_detail.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Line_detail")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(line_detail.Name))
	return
}

func (line_width *Line_width) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", line_width.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Line_width")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(line_width.Name))
	return
}

func (link *Link) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", link.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Link")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(link.Name))
	return
}

func (listen *Listen) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", listen.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Listen")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(listen.Name))
	return
}

func (listening *Listening) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", listening.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Listening")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(listening.Name))
	return
}

func (lyric *Lyric) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", lyric.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Lyric")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(lyric.Name))
	return
}

func (lyric_font *Lyric_font) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", lyric_font.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Lyric_font")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(lyric_font.Name))
	return
}

func (lyric_language *Lyric_language) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", lyric_language.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Lyric_language")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(lyric_language.Name))
	return
}

func (measure_layout *Measure_layout) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", measure_layout.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Measure_layout")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(measure_layout.Name))
	return
}

func (measure_numbering *Measure_numbering) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", measure_numbering.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Measure_numbering")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(measure_numbering.Name))
	return
}

func (measure_repeat *Measure_repeat) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", measure_repeat.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Measure_repeat")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(measure_repeat.Name))
	return
}

func (measure_style *Measure_style) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", measure_style.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Measure_style")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(measure_style.Name))
	return
}

func (membrane *Membrane) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", membrane.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Membrane")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(membrane.Name))
	return
}

func (metal *Metal) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", metal.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Metal")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(metal.Name))
	return
}

func (metronome *Metronome) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", metronome.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Metronome")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(metronome.Name))
	return
}

func (metronome_beam *Metronome_beam) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", metronome_beam.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Metronome_beam")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(metronome_beam.Name))
	return
}

func (metronome_note *Metronome_note) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", metronome_note.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Metronome_note")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(metronome_note.Name))
	return
}

func (metronome_tied *Metronome_tied) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", metronome_tied.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Metronome_tied")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(metronome_tied.Name))
	return
}

func (metronome_tuplet *Metronome_tuplet) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", metronome_tuplet.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Metronome_tuplet")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(metronome_tuplet.Name))
	return
}

func (midi_device *Midi_device) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", midi_device.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Midi_device")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(midi_device.Name))
	return
}

func (midi_instrument *Midi_instrument) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", midi_instrument.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Midi_instrument")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(midi_instrument.Name))
	return
}

func (miscellaneous *Miscellaneous) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", miscellaneous.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Miscellaneous")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(miscellaneous.Name))
	return
}

func (miscellaneous_field *Miscellaneous_field) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", miscellaneous_field.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Miscellaneous_field")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(miscellaneous_field.Name))
	return
}

func (mordent *Mordent) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", mordent.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Mordent")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(mordent.Name))
	return
}

func (multiple_rest *Multiple_rest) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", multiple_rest.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Multiple_rest")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(multiple_rest.Name))
	return
}

func (name_display *Name_display) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", name_display.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Name_display")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name_display.Name))
	return
}

func (non_arpeggiate *Non_arpeggiate) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", non_arpeggiate.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Non_arpeggiate")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(non_arpeggiate.Name))
	return
}

func (notations *Notations) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notations.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Notations")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(notations.Name))
	return
}

func (note *Note) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(note.Name))
	return
}

func (note_size *Note_size) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note_size.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note_size")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(note_size.Name))
	return
}

func (note_type *Note_type) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note_type.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Note_type")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(note_type.Name))
	return
}

func (notehead *Notehead) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notehead.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Notehead")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(notehead.Name))
	return
}

func (notehead_text *Notehead_text) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notehead_text.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Notehead_text")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(notehead_text.Name))
	return
}

func (numeral *Numeral) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", numeral.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Numeral")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(numeral.Name))
	return
}

func (numeral_key *Numeral_key) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", numeral_key.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Numeral_key")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(numeral_key.Name))
	return
}

func (numeral_root *Numeral_root) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", numeral_root.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Numeral_root")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(numeral_root.Name))
	return
}

func (octave_shift *Octave_shift) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", octave_shift.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Octave_shift")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(octave_shift.Name))
	return
}

func (offset *Offset) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", offset.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Offset")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(offset.Name))
	return
}

func (opus *Opus) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", opus.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Opus")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(opus.Name))
	return
}

func (ornaments *Ornaments) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", ornaments.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Ornaments")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(ornaments.Name))
	return
}

func (other_appearance *Other_appearance) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", other_appearance.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Other_appearance")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(other_appearance.Name))
	return
}

func (other_direction *Other_direction) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", other_direction.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Other_direction")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(other_direction.Name))
	return
}

func (other_listening *Other_listening) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", other_listening.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Other_listening")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(other_listening.Name))
	return
}

func (other_notation *Other_notation) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", other_notation.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Other_notation")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(other_notation.Name))
	return
}

func (other_placement_text *Other_placement_text) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", other_placement_text.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Other_placement_text")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(other_placement_text.Name))
	return
}

func (other_play *Other_play) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", other_play.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Other_play")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(other_play.Name))
	return
}

func (other_text *Other_text) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", other_text.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Other_text")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(other_text.Name))
	return
}

func (page_layout *Page_layout) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", page_layout.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Page_layout")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(page_layout.Name))
	return
}

func (page_margins *Page_margins) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", page_margins.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Page_margins")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(page_margins.Name))
	return
}

func (part_clef *Part_clef) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part_clef.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Part_clef")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(part_clef.Name))
	return
}

func (part_group *Part_group) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part_group.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Part_group")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(part_group.Name))
	return
}

func (part_link *Part_link) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part_link.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Part_link")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(part_link.Name))
	return
}

func (part_list *Part_list) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part_list.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Part_list")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(part_list.Name))
	return
}

func (part_name *Part_name) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part_name.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Part_name")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(part_name.Name))
	return
}

func (part_symbol *Part_symbol) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part_symbol.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Part_symbol")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(part_symbol.Name))
	return
}

func (part_transpose *Part_transpose) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part_transpose.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Part_transpose")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(part_transpose.Name))
	return
}

func (pedal *Pedal) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pedal.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Pedal")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(pedal.Name))
	return
}

func (pedal_tuning *Pedal_tuning) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pedal_tuning.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Pedal_tuning")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(pedal_tuning.Name))
	return
}

func (per_minute *Per_minute) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", per_minute.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Per_minute")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(per_minute.Name))
	return
}

func (percussion *Percussion) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", percussion.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Percussion")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(percussion.Name))
	return
}

func (pitch *Pitch) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pitch.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Pitch")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(pitch.Name))
	return
}

func (pitched *Pitched) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pitched.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Pitched")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(pitched.Name))
	return
}

func (placement_text *Placement_text) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", placement_text.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Placement_text")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(placement_text.Name))
	return
}

func (play *Play) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", play.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Play")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(play.Name))
	return
}

func (player *Player) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", player.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Player")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(player.Name))
	return
}

func (principal_voice *Principal_voice) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", principal_voice.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Principal_voice")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(principal_voice.Name))
	return
}

func (print *Print) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", print.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Print")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(print.Name))
	return
}

func (release *Release) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", release.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Release")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(release.Name))
	return
}

func (repeat *Repeat) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", repeat.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Repeat")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(repeat.Name))
	return
}

func (rest *Rest) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rest.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Rest")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(rest.Name))
	return
}

func (root *Root) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", root.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Root")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(root.Name))
	return
}

func (root_step *Root_step) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", root_step.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Root_step")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(root_step.Name))
	return
}

func (scaling *Scaling) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", scaling.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Scaling")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(scaling.Name))
	return
}

func (scordatura *Scordatura) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", scordatura.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Scordatura")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(scordatura.Name))
	return
}

func (score_instrument *Score_instrument) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", score_instrument.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Score_instrument")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(score_instrument.Name))
	return
}

func (score_part *Score_part) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", score_part.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Score_part")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(score_part.Name))
	return
}

func (score_partwise *Score_partwise) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", score_partwise.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Score_partwise")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(score_partwise.Name))
	return
}

func (score_timewise *Score_timewise) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", score_timewise.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Score_timewise")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(score_timewise.Name))
	return
}

func (segno *Segno) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", segno.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Segno")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(segno.Name))
	return
}

func (slash *Slash) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", slash.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Slash")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(slash.Name))
	return
}

func (slide *Slide) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", slide.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Slide")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(slide.Name))
	return
}

func (slur *Slur) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", slur.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Slur")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(slur.Name))
	return
}

func (sound *Sound) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", sound.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Sound")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(sound.Name))
	return
}

func (staff_details *Staff_details) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staff_details.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Staff_details")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(staff_details.Name))
	return
}

func (staff_divide *Staff_divide) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staff_divide.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Staff_divide")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(staff_divide.Name))
	return
}

func (staff_layout *Staff_layout) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staff_layout.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Staff_layout")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(staff_layout.Name))
	return
}

func (staff_size *Staff_size) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staff_size.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Staff_size")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(staff_size.Name))
	return
}

func (staff_tuning *Staff_tuning) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staff_tuning.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Staff_tuning")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(staff_tuning.Name))
	return
}

func (stem *Stem) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stem.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Stem")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stem.Name))
	return
}

func (stick *Stick) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stick.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Stick")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(stick.Name))
	return
}

func (string_mute *String_mute) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", string_mute.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "String_mute")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(string_mute.Name))
	return
}

func (string_type *String_type) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", string_type.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "String_type")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(string_type.Name))
	return
}

func (strong_accent *Strong_accent) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", strong_accent.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Strong_accent")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(strong_accent.Name))
	return
}

func (style_text *Style_text) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", style_text.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Style_text")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(style_text.Name))
	return
}

func (supports *Supports) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", supports.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Supports")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(supports.Name))
	return
}

func (swing *Swing) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", swing.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Swing")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(swing.Name))
	return
}

func (sync *Sync) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", sync.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Sync")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(sync.Name))
	return
}

func (system_dividers *System_dividers) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", system_dividers.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "System_dividers")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(system_dividers.Name))
	return
}

func (system_layout *System_layout) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", system_layout.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "System_layout")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(system_layout.Name))
	return
}

func (system_margins *System_margins) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", system_margins.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "System_margins")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(system_margins.Name))
	return
}

func (tap *Tap) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tap.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tap")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tap.Name))
	return
}

func (technical *Technical) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", technical.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Technical")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(technical.Name))
	return
}

func (text_element_data *Text_element_data) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", text_element_data.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Text_element_data")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(text_element_data.Name))
	return
}

func (tie *Tie) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tie.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tie")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tie.Name))
	return
}

func (tied *Tied) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tied.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tied")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tied.Name))
	return
}

func (time *Time) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", time.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Time")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(time.Name))
	return
}

func (time_modification *Time_modification) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", time_modification.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Time_modification")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(time_modification.Name))
	return
}

func (timpani *Timpani) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", timpani.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Timpani")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(timpani.Name))
	return
}

func (transpose *Transpose) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", transpose.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Transpose")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(transpose.Name))
	return
}

func (tremolo *Tremolo) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tremolo.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tremolo")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tremolo.Name))
	return
}

func (tuplet *Tuplet) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tuplet.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tuplet")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tuplet.Name))
	return
}

func (tuplet_dot *Tuplet_dot) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tuplet_dot.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tuplet_dot")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tuplet_dot.Name))
	return
}

func (tuplet_number *Tuplet_number) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tuplet_number.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tuplet_number")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tuplet_number.Name))
	return
}

func (tuplet_portion *Tuplet_portion) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tuplet_portion.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tuplet_portion")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tuplet_portion.Name))
	return
}

func (tuplet_type *Tuplet_type) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tuplet_type.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Tuplet_type")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(tuplet_type.Name))
	return
}

func (typed_text *Typed_text) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", typed_text.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Typed_text")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(typed_text.Name))
	return
}

func (unpitched *Unpitched) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", unpitched.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Unpitched")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(unpitched.Name))
	return
}

func (virtual_instrument *Virtual_instrument) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", virtual_instrument.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Virtual_instrument")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(virtual_instrument.Name))
	return
}

func (wait *Wait) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", wait.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Wait")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(wait.Name))
	return
}

func (wavy_line *Wavy_line) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", wavy_line.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Wavy_line")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(wavy_line.Name))
	return
}

func (wedge *Wedge) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", wedge.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Wedge")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(wedge.Name))
	return
}

func (wood *Wood) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", wood.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Wood")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(wood.Name))
	return
}

func (work *Work) GongMarshallIdentifier(stage *Stage) (decl string) {
	decl = GongIdentifiersDecls
	decl = strings.ReplaceAll(decl, "{{Identifier}}", work.GongGetIdentifier(stage))
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", "Work")
	decl = strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(work.Name))
	return
}

// insertion point for unstaging
func (a_directive *A_directive) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_directive.GongGetReferenceIdentifier(stage))
	return
}

func (a_measure *A_measure) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_measure.GongGetReferenceIdentifier(stage))
	return
}

func (a_measure_1 *A_measure_1) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_measure_1.GongGetReferenceIdentifier(stage))
	return
}

func (a_part *A_part) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_part.GongGetReferenceIdentifier(stage))
	return
}

func (a_part_1 *A_part_1) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", a_part_1.GongGetReferenceIdentifier(stage))
	return
}

func (accidental *Accidental) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", accidental.GongGetReferenceIdentifier(stage))
	return
}

func (accidental_mark *Accidental_mark) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", accidental_mark.GongGetReferenceIdentifier(stage))
	return
}

func (accidental_text *Accidental_text) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", accidental_text.GongGetReferenceIdentifier(stage))
	return
}

func (accord *Accord) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", accord.GongGetReferenceIdentifier(stage))
	return
}

func (accordion_registration *Accordion_registration) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", accordion_registration.GongGetReferenceIdentifier(stage))
	return
}

func (appearance *Appearance) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", appearance.GongGetReferenceIdentifier(stage))
	return
}

func (arpeggiate *Arpeggiate) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", arpeggiate.GongGetReferenceIdentifier(stage))
	return
}

func (arrow *Arrow) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", arrow.GongGetReferenceIdentifier(stage))
	return
}

func (articulations *Articulations) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", articulations.GongGetReferenceIdentifier(stage))
	return
}

func (assess *Assess) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", assess.GongGetReferenceIdentifier(stage))
	return
}

func (attributes *Attributes) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", attributes.GongGetReferenceIdentifier(stage))
	return
}

func (backup *Backup) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", backup.GongGetReferenceIdentifier(stage))
	return
}

func (bar_style_color *Bar_style_color) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bar_style_color.GongGetReferenceIdentifier(stage))
	return
}

func (barline *Barline) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", barline.GongGetReferenceIdentifier(stage))
	return
}

func (barre *Barre) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", barre.GongGetReferenceIdentifier(stage))
	return
}

func (bass *Bass) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bass.GongGetReferenceIdentifier(stage))
	return
}

func (bass_step *Bass_step) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bass_step.GongGetReferenceIdentifier(stage))
	return
}

func (beam *Beam) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", beam.GongGetReferenceIdentifier(stage))
	return
}

func (beat_repeat *Beat_repeat) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", beat_repeat.GongGetReferenceIdentifier(stage))
	return
}

func (beat_unit_tied *Beat_unit_tied) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", beat_unit_tied.GongGetReferenceIdentifier(stage))
	return
}

func (beater *Beater) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", beater.GongGetReferenceIdentifier(stage))
	return
}

func (bend *Bend) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bend.GongGetReferenceIdentifier(stage))
	return
}

func (bookmark *Bookmark) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bookmark.GongGetReferenceIdentifier(stage))
	return
}

func (bracket *Bracket) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", bracket.GongGetReferenceIdentifier(stage))
	return
}

func (breath_mark *Breath_mark) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", breath_mark.GongGetReferenceIdentifier(stage))
	return
}

func (caesura *Caesura) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", caesura.GongGetReferenceIdentifier(stage))
	return
}

func (cancel *Cancel) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", cancel.GongGetReferenceIdentifier(stage))
	return
}

func (clef *Clef) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", clef.GongGetReferenceIdentifier(stage))
	return
}

func (coda *Coda) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", coda.GongGetReferenceIdentifier(stage))
	return
}

func (credit *Credit) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", credit.GongGetReferenceIdentifier(stage))
	return
}

func (dashes *Dashes) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dashes.GongGetReferenceIdentifier(stage))
	return
}

func (defaults *Defaults) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", defaults.GongGetReferenceIdentifier(stage))
	return
}

func (degree *Degree) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", degree.GongGetReferenceIdentifier(stage))
	return
}

func (degree_alter *Degree_alter) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", degree_alter.GongGetReferenceIdentifier(stage))
	return
}

func (degree_type *Degree_type) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", degree_type.GongGetReferenceIdentifier(stage))
	return
}

func (degree_value *Degree_value) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", degree_value.GongGetReferenceIdentifier(stage))
	return
}

func (direction *Direction) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", direction.GongGetReferenceIdentifier(stage))
	return
}

func (direction_type *Direction_type) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", direction_type.GongGetReferenceIdentifier(stage))
	return
}

func (distance *Distance) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", distance.GongGetReferenceIdentifier(stage))
	return
}

func (double *Double) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", double.GongGetReferenceIdentifier(stage))
	return
}

func (dynamics *Dynamics) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", dynamics.GongGetReferenceIdentifier(stage))
	return
}

func (effect *Effect) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", effect.GongGetReferenceIdentifier(stage))
	return
}

func (elision *Elision) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", elision.GongGetReferenceIdentifier(stage))
	return
}

func (empty *Empty) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty.GongGetReferenceIdentifier(stage))
	return
}

func (empty_font *Empty_font) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_font.GongGetReferenceIdentifier(stage))
	return
}

func (empty_line *Empty_line) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_line.GongGetReferenceIdentifier(stage))
	return
}

func (empty_placement *Empty_placement) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_placement.GongGetReferenceIdentifier(stage))
	return
}

func (empty_placement_smufl *Empty_placement_smufl) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_placement_smufl.GongGetReferenceIdentifier(stage))
	return
}

func (empty_print_object_style_align *Empty_print_object_style_align) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_print_object_style_align.GongGetReferenceIdentifier(stage))
	return
}

func (empty_print_style *Empty_print_style) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_print_style.GongGetReferenceIdentifier(stage))
	return
}

func (empty_print_style_align *Empty_print_style_align) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_print_style_align.GongGetReferenceIdentifier(stage))
	return
}

func (empty_print_style_align_id *Empty_print_style_align_id) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_print_style_align_id.GongGetReferenceIdentifier(stage))
	return
}

func (empty_trill_sound *Empty_trill_sound) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", empty_trill_sound.GongGetReferenceIdentifier(stage))
	return
}

func (encoding *Encoding) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", encoding.GongGetReferenceIdentifier(stage))
	return
}

func (ending *Ending) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", ending.GongGetReferenceIdentifier(stage))
	return
}

func (extend *Extend) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", extend.GongGetReferenceIdentifier(stage))
	return
}

func (feature *Feature) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", feature.GongGetReferenceIdentifier(stage))
	return
}

func (fermata *Fermata) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", fermata.GongGetReferenceIdentifier(stage))
	return
}

func (figure *Figure) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", figure.GongGetReferenceIdentifier(stage))
	return
}

func (figured_bass *Figured_bass) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", figured_bass.GongGetReferenceIdentifier(stage))
	return
}

func (fingering *Fingering) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", fingering.GongGetReferenceIdentifier(stage))
	return
}

func (first_fret *First_fret) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", first_fret.GongGetReferenceIdentifier(stage))
	return
}

func (for_part *For_part) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", for_part.GongGetReferenceIdentifier(stage))
	return
}

func (formatted_symbol *Formatted_symbol) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formatted_symbol.GongGetReferenceIdentifier(stage))
	return
}

func (formatted_symbol_id *Formatted_symbol_id) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formatted_symbol_id.GongGetReferenceIdentifier(stage))
	return
}

func (formatted_text *Formatted_text) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formatted_text.GongGetReferenceIdentifier(stage))
	return
}

func (formatted_text_id *Formatted_text_id) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", formatted_text_id.GongGetReferenceIdentifier(stage))
	return
}

func (forward *Forward) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", forward.GongGetReferenceIdentifier(stage))
	return
}

func (frame *Frame) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", frame.GongGetReferenceIdentifier(stage))
	return
}

func (frame_note *Frame_note) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", frame_note.GongGetReferenceIdentifier(stage))
	return
}

func (fret *Fret) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", fret.GongGetReferenceIdentifier(stage))
	return
}

func (glass *Glass) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", glass.GongGetReferenceIdentifier(stage))
	return
}

func (glissando *Glissando) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", glissando.GongGetReferenceIdentifier(stage))
	return
}

func (glyph *Glyph) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", glyph.GongGetReferenceIdentifier(stage))
	return
}

func (grace *Grace) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", grace.GongGetReferenceIdentifier(stage))
	return
}

func (group_barline *Group_barline) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group_barline.GongGetReferenceIdentifier(stage))
	return
}

func (group_name *Group_name) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group_name.GongGetReferenceIdentifier(stage))
	return
}

func (group_symbol *Group_symbol) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", group_symbol.GongGetReferenceIdentifier(stage))
	return
}

func (grouping *Grouping) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", grouping.GongGetReferenceIdentifier(stage))
	return
}

func (hammer_on_pull_off *Hammer_on_pull_off) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", hammer_on_pull_off.GongGetReferenceIdentifier(stage))
	return
}

func (handbell *Handbell) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", handbell.GongGetReferenceIdentifier(stage))
	return
}

func (harmon_closed *Harmon_closed) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", harmon_closed.GongGetReferenceIdentifier(stage))
	return
}

func (harmon_mute *Harmon_mute) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", harmon_mute.GongGetReferenceIdentifier(stage))
	return
}

func (harmonic *Harmonic) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", harmonic.GongGetReferenceIdentifier(stage))
	return
}

func (harmony *Harmony) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", harmony.GongGetReferenceIdentifier(stage))
	return
}

func (harmony_alter *Harmony_alter) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", harmony_alter.GongGetReferenceIdentifier(stage))
	return
}

func (harp_pedals *Harp_pedals) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", harp_pedals.GongGetReferenceIdentifier(stage))
	return
}

func (heel_toe *Heel_toe) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", heel_toe.GongGetReferenceIdentifier(stage))
	return
}

func (hole *Hole) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", hole.GongGetReferenceIdentifier(stage))
	return
}

func (hole_closed *Hole_closed) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", hole_closed.GongGetReferenceIdentifier(stage))
	return
}

func (horizontal_turn *Horizontal_turn) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", horizontal_turn.GongGetReferenceIdentifier(stage))
	return
}

func (identification *Identification) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", identification.GongGetReferenceIdentifier(stage))
	return
}

func (image *Image) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", image.GongGetReferenceIdentifier(stage))
	return
}

func (instrument *Instrument) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", instrument.GongGetReferenceIdentifier(stage))
	return
}

func (instrument_change *Instrument_change) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", instrument_change.GongGetReferenceIdentifier(stage))
	return
}

func (instrument_link *Instrument_link) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", instrument_link.GongGetReferenceIdentifier(stage))
	return
}

func (interchangeable *Interchangeable) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", interchangeable.GongGetReferenceIdentifier(stage))
	return
}

func (inversion *Inversion) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", inversion.GongGetReferenceIdentifier(stage))
	return
}

func (key *Key) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", key.GongGetReferenceIdentifier(stage))
	return
}

func (key_accidental *Key_accidental) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", key_accidental.GongGetReferenceIdentifier(stage))
	return
}

func (key_octave *Key_octave) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", key_octave.GongGetReferenceIdentifier(stage))
	return
}

func (kind *Kind) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", kind.GongGetReferenceIdentifier(stage))
	return
}

func (level *Level) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", level.GongGetReferenceIdentifier(stage))
	return
}

func (line_detail *Line_detail) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", line_detail.GongGetReferenceIdentifier(stage))
	return
}

func (line_width *Line_width) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", line_width.GongGetReferenceIdentifier(stage))
	return
}

func (link *Link) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", link.GongGetReferenceIdentifier(stage))
	return
}

func (listen *Listen) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", listen.GongGetReferenceIdentifier(stage))
	return
}

func (listening *Listening) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", listening.GongGetReferenceIdentifier(stage))
	return
}

func (lyric *Lyric) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", lyric.GongGetReferenceIdentifier(stage))
	return
}

func (lyric_font *Lyric_font) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", lyric_font.GongGetReferenceIdentifier(stage))
	return
}

func (lyric_language *Lyric_language) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", lyric_language.GongGetReferenceIdentifier(stage))
	return
}

func (measure_layout *Measure_layout) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", measure_layout.GongGetReferenceIdentifier(stage))
	return
}

func (measure_numbering *Measure_numbering) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", measure_numbering.GongGetReferenceIdentifier(stage))
	return
}

func (measure_repeat *Measure_repeat) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", measure_repeat.GongGetReferenceIdentifier(stage))
	return
}

func (measure_style *Measure_style) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", measure_style.GongGetReferenceIdentifier(stage))
	return
}

func (membrane *Membrane) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", membrane.GongGetReferenceIdentifier(stage))
	return
}

func (metal *Metal) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", metal.GongGetReferenceIdentifier(stage))
	return
}

func (metronome *Metronome) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", metronome.GongGetReferenceIdentifier(stage))
	return
}

func (metronome_beam *Metronome_beam) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", metronome_beam.GongGetReferenceIdentifier(stage))
	return
}

func (metronome_note *Metronome_note) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", metronome_note.GongGetReferenceIdentifier(stage))
	return
}

func (metronome_tied *Metronome_tied) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", metronome_tied.GongGetReferenceIdentifier(stage))
	return
}

func (metronome_tuplet *Metronome_tuplet) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", metronome_tuplet.GongGetReferenceIdentifier(stage))
	return
}

func (midi_device *Midi_device) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", midi_device.GongGetReferenceIdentifier(stage))
	return
}

func (midi_instrument *Midi_instrument) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", midi_instrument.GongGetReferenceIdentifier(stage))
	return
}

func (miscellaneous *Miscellaneous) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", miscellaneous.GongGetReferenceIdentifier(stage))
	return
}

func (miscellaneous_field *Miscellaneous_field) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", miscellaneous_field.GongGetReferenceIdentifier(stage))
	return
}

func (mordent *Mordent) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", mordent.GongGetReferenceIdentifier(stage))
	return
}

func (multiple_rest *Multiple_rest) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", multiple_rest.GongGetReferenceIdentifier(stage))
	return
}

func (name_display *Name_display) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", name_display.GongGetReferenceIdentifier(stage))
	return
}

func (non_arpeggiate *Non_arpeggiate) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", non_arpeggiate.GongGetReferenceIdentifier(stage))
	return
}

func (notations *Notations) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notations.GongGetReferenceIdentifier(stage))
	return
}

func (note *Note) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note.GongGetReferenceIdentifier(stage))
	return
}

func (note_size *Note_size) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note_size.GongGetReferenceIdentifier(stage))
	return
}

func (note_type *Note_type) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", note_type.GongGetReferenceIdentifier(stage))
	return
}

func (notehead *Notehead) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notehead.GongGetReferenceIdentifier(stage))
	return
}

func (notehead_text *Notehead_text) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", notehead_text.GongGetReferenceIdentifier(stage))
	return
}

func (numeral *Numeral) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", numeral.GongGetReferenceIdentifier(stage))
	return
}

func (numeral_key *Numeral_key) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", numeral_key.GongGetReferenceIdentifier(stage))
	return
}

func (numeral_root *Numeral_root) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", numeral_root.GongGetReferenceIdentifier(stage))
	return
}

func (octave_shift *Octave_shift) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", octave_shift.GongGetReferenceIdentifier(stage))
	return
}

func (offset *Offset) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", offset.GongGetReferenceIdentifier(stage))
	return
}

func (opus *Opus) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", opus.GongGetReferenceIdentifier(stage))
	return
}

func (ornaments *Ornaments) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", ornaments.GongGetReferenceIdentifier(stage))
	return
}

func (other_appearance *Other_appearance) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", other_appearance.GongGetReferenceIdentifier(stage))
	return
}

func (other_direction *Other_direction) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", other_direction.GongGetReferenceIdentifier(stage))
	return
}

func (other_listening *Other_listening) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", other_listening.GongGetReferenceIdentifier(stage))
	return
}

func (other_notation *Other_notation) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", other_notation.GongGetReferenceIdentifier(stage))
	return
}

func (other_placement_text *Other_placement_text) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", other_placement_text.GongGetReferenceIdentifier(stage))
	return
}

func (other_play *Other_play) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", other_play.GongGetReferenceIdentifier(stage))
	return
}

func (other_text *Other_text) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", other_text.GongGetReferenceIdentifier(stage))
	return
}

func (page_layout *Page_layout) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", page_layout.GongGetReferenceIdentifier(stage))
	return
}

func (page_margins *Page_margins) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", page_margins.GongGetReferenceIdentifier(stage))
	return
}

func (part_clef *Part_clef) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part_clef.GongGetReferenceIdentifier(stage))
	return
}

func (part_group *Part_group) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part_group.GongGetReferenceIdentifier(stage))
	return
}

func (part_link *Part_link) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part_link.GongGetReferenceIdentifier(stage))
	return
}

func (part_list *Part_list) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part_list.GongGetReferenceIdentifier(stage))
	return
}

func (part_name *Part_name) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part_name.GongGetReferenceIdentifier(stage))
	return
}

func (part_symbol *Part_symbol) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part_symbol.GongGetReferenceIdentifier(stage))
	return
}

func (part_transpose *Part_transpose) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", part_transpose.GongGetReferenceIdentifier(stage))
	return
}

func (pedal *Pedal) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pedal.GongGetReferenceIdentifier(stage))
	return
}

func (pedal_tuning *Pedal_tuning) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pedal_tuning.GongGetReferenceIdentifier(stage))
	return
}

func (per_minute *Per_minute) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", per_minute.GongGetReferenceIdentifier(stage))
	return
}

func (percussion *Percussion) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", percussion.GongGetReferenceIdentifier(stage))
	return
}

func (pitch *Pitch) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pitch.GongGetReferenceIdentifier(stage))
	return
}

func (pitched *Pitched) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", pitched.GongGetReferenceIdentifier(stage))
	return
}

func (placement_text *Placement_text) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", placement_text.GongGetReferenceIdentifier(stage))
	return
}

func (play *Play) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", play.GongGetReferenceIdentifier(stage))
	return
}

func (player *Player) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", player.GongGetReferenceIdentifier(stage))
	return
}

func (principal_voice *Principal_voice) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", principal_voice.GongGetReferenceIdentifier(stage))
	return
}

func (print *Print) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", print.GongGetReferenceIdentifier(stage))
	return
}

func (release *Release) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", release.GongGetReferenceIdentifier(stage))
	return
}

func (repeat *Repeat) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", repeat.GongGetReferenceIdentifier(stage))
	return
}

func (rest *Rest) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", rest.GongGetReferenceIdentifier(stage))
	return
}

func (root *Root) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", root.GongGetReferenceIdentifier(stage))
	return
}

func (root_step *Root_step) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", root_step.GongGetReferenceIdentifier(stage))
	return
}

func (scaling *Scaling) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", scaling.GongGetReferenceIdentifier(stage))
	return
}

func (scordatura *Scordatura) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", scordatura.GongGetReferenceIdentifier(stage))
	return
}

func (score_instrument *Score_instrument) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", score_instrument.GongGetReferenceIdentifier(stage))
	return
}

func (score_part *Score_part) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", score_part.GongGetReferenceIdentifier(stage))
	return
}

func (score_partwise *Score_partwise) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", score_partwise.GongGetReferenceIdentifier(stage))
	return
}

func (score_timewise *Score_timewise) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", score_timewise.GongGetReferenceIdentifier(stage))
	return
}

func (segno *Segno) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", segno.GongGetReferenceIdentifier(stage))
	return
}

func (slash *Slash) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", slash.GongGetReferenceIdentifier(stage))
	return
}

func (slide *Slide) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", slide.GongGetReferenceIdentifier(stage))
	return
}

func (slur *Slur) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", slur.GongGetReferenceIdentifier(stage))
	return
}

func (sound *Sound) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", sound.GongGetReferenceIdentifier(stage))
	return
}

func (staff_details *Staff_details) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staff_details.GongGetReferenceIdentifier(stage))
	return
}

func (staff_divide *Staff_divide) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staff_divide.GongGetReferenceIdentifier(stage))
	return
}

func (staff_layout *Staff_layout) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staff_layout.GongGetReferenceIdentifier(stage))
	return
}

func (staff_size *Staff_size) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staff_size.GongGetReferenceIdentifier(stage))
	return
}

func (staff_tuning *Staff_tuning) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", staff_tuning.GongGetReferenceIdentifier(stage))
	return
}

func (stem *Stem) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stem.GongGetReferenceIdentifier(stage))
	return
}

func (stick *Stick) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", stick.GongGetReferenceIdentifier(stage))
	return
}

func (string_mute *String_mute) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", string_mute.GongGetReferenceIdentifier(stage))
	return
}

func (string_type *String_type) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", string_type.GongGetReferenceIdentifier(stage))
	return
}

func (strong_accent *Strong_accent) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", strong_accent.GongGetReferenceIdentifier(stage))
	return
}

func (style_text *Style_text) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", style_text.GongGetReferenceIdentifier(stage))
	return
}

func (supports *Supports) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", supports.GongGetReferenceIdentifier(stage))
	return
}

func (swing *Swing) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", swing.GongGetReferenceIdentifier(stage))
	return
}

func (sync *Sync) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", sync.GongGetReferenceIdentifier(stage))
	return
}

func (system_dividers *System_dividers) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", system_dividers.GongGetReferenceIdentifier(stage))
	return
}

func (system_layout *System_layout) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", system_layout.GongGetReferenceIdentifier(stage))
	return
}

func (system_margins *System_margins) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", system_margins.GongGetReferenceIdentifier(stage))
	return
}

func (tap *Tap) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tap.GongGetReferenceIdentifier(stage))
	return
}

func (technical *Technical) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", technical.GongGetReferenceIdentifier(stage))
	return
}

func (text_element_data *Text_element_data) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", text_element_data.GongGetReferenceIdentifier(stage))
	return
}

func (tie *Tie) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tie.GongGetReferenceIdentifier(stage))
	return
}

func (tied *Tied) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tied.GongGetReferenceIdentifier(stage))
	return
}

func (time *Time) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", time.GongGetReferenceIdentifier(stage))
	return
}

func (time_modification *Time_modification) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", time_modification.GongGetReferenceIdentifier(stage))
	return
}

func (timpani *Timpani) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", timpani.GongGetReferenceIdentifier(stage))
	return
}

func (transpose *Transpose) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", transpose.GongGetReferenceIdentifier(stage))
	return
}

func (tremolo *Tremolo) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tremolo.GongGetReferenceIdentifier(stage))
	return
}

func (tuplet *Tuplet) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tuplet.GongGetReferenceIdentifier(stage))
	return
}

func (tuplet_dot *Tuplet_dot) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tuplet_dot.GongGetReferenceIdentifier(stage))
	return
}

func (tuplet_number *Tuplet_number) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tuplet_number.GongGetReferenceIdentifier(stage))
	return
}

func (tuplet_portion *Tuplet_portion) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tuplet_portion.GongGetReferenceIdentifier(stage))
	return
}

func (tuplet_type *Tuplet_type) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", tuplet_type.GongGetReferenceIdentifier(stage))
	return
}

func (typed_text *Typed_text) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", typed_text.GongGetReferenceIdentifier(stage))
	return
}

func (unpitched *Unpitched) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", unpitched.GongGetReferenceIdentifier(stage))
	return
}

func (virtual_instrument *Virtual_instrument) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", virtual_instrument.GongGetReferenceIdentifier(stage))
	return
}

func (wait *Wait) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", wait.GongGetReferenceIdentifier(stage))
	return
}

func (wavy_line *Wavy_line) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", wavy_line.GongGetReferenceIdentifier(stage))
	return
}

func (wedge *Wedge) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", wedge.GongGetReferenceIdentifier(stage))
	return
}

func (wood *Wood) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", wood.GongGetReferenceIdentifier(stage))
	return
}

func (work *Work) GongMarshallUnstaging(stage *Stage) (decl string) {
	decl = GongUnstageStmt
	decl = strings.ReplaceAll(decl, "{{Identifier}}", work.GongGetReferenceIdentifier(stage))
	return
}

func GongIntToLetters(number int32) (letters string) {
	number--
	if firstLetter := number / 26; firstLetter > 0 {
		letters += GongIntToLetters(firstLetter)
		letters += string('A' + number%26)
	} else {
		letters += string('A' + number)
	}

	return
}

// GongGenerateReproducibleUUIDv4 creates a deterministic UUIDv4 based on a string and a positive integer.
func GongGenerateReproducibleUUIDv4(seedStr string, seedInt uint64) string {
	// 1. Create a deterministic hash from the inputs using SHA-256
	h := sha256.New()

	// Write the string to the hash
	h.Write([]byte(seedStr))

	// Write the integer to the hash (using BigEndian to ensure consistency across architectures)
	intBytes := make([]byte, 8)
	binary.BigEndian.PutUint64(intBytes, seedInt)
	h.Write(intBytes)

	// 2. Extract the first 16 bytes from our resulting hash
	hashBytes := h.Sum(nil)
	uuid := make([]byte, 16)
	copy(uuid, hashBytes[:16])

	// 3. Set the Version to 4 (0100 in binary)
	// We take the 7th byte, clear the top 4 bits with & 0x0f, and set the top bits to 0100 with | 0x40
	uuid[6] = (uuid[6] & 0x0f) | 0x40

	// 4. Set the Variant to RFC4122 (10 in binary)
	// We take the 9th byte, clear the top 2 bits with & 0x3f, and set the top bits to 10 with | 0x80
	uuid[8] = (uuid[8] & 0x3f) | 0x80

	// 5. Format and return the byte array as a standard UUID string
	return fmt.Sprintf("%08x-%04x-%04x-%04x-%012x",
		uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:16])
}

// end of template
