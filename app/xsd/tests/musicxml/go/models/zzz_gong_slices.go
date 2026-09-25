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

	// Compute reverse map for named struct Dynamics
	// insertion point per field
	stage.Dynamics_Other_dynamics_reverseMap = make(map[*Other_text]*Dynamics)
	for dynamics := range stage.Dynamicss {
		_ = dynamics
		for _, _other_text := range dynamics.Other_dynamics {
			stage.Dynamics_Other_dynamics_reverseMap[_other_text] = dynamics
		}
	}

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

	// Compute reverse map for named struct Figured_bass
	// insertion point per field
	stage.Figured_bass_Figure_reverseMap = make(map[*Figure]*Figured_bass)
	for figured_bass := range stage.Figured_basss {
		_ = figured_bass
		for _, _figure := range figured_bass.Figure {
			stage.Figured_bass_Figure_reverseMap[_figure] = figured_bass
		}
	}

	// Compute reverse map for named struct Frame
	// insertion point per field
	stage.Frame_Frame_note_reverseMap = make(map[*Frame_note]*Frame)
	for frame := range stage.Frames {
		_ = frame
		for _, _frame_note := range frame.Frame_note {
			stage.Frame_Frame_note_reverseMap[_frame_note] = frame
		}
	}

	// Compute reverse map for named struct Grouping
	// insertion point per field
	stage.Grouping_Feature_reverseMap = make(map[*Feature]*Grouping)
	for grouping := range stage.Groupings {
		_ = grouping
		for _, _feature := range grouping.Feature {
			stage.Grouping_Feature_reverseMap[_feature] = grouping
		}
	}

	// Compute reverse map for named struct Harmony
	// insertion point per field
	stage.Harmony_Degree_reverseMap = make(map[*Degree]*Harmony)
	for harmony := range stage.Harmonys {
		_ = harmony
		for _, _degree := range harmony.Degree {
			stage.Harmony_Degree_reverseMap[_degree] = harmony
		}
	}

	// Compute reverse map for named struct Harp_pedals
	// insertion point per field
	stage.Harp_pedals_Pedal_tuning_reverseMap = make(map[*Pedal_tuning]*Harp_pedals)
	for harp_pedals := range stage.Harp_pedalss {
		_ = harp_pedals
		for _, _pedal_tuning := range harp_pedals.Pedal_tuning {
			stage.Harp_pedals_Pedal_tuning_reverseMap[_pedal_tuning] = harp_pedals
		}
	}

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

	// Compute reverse map for named struct Key
	// insertion point per field
	stage.Key_Key_octave_reverseMap = make(map[*Key_octave]*Key)
	for key := range stage.Keys {
		_ = key
		for _, _key_octave := range key.Key_octave {
			stage.Key_Key_octave_reverseMap[_key_octave] = key
		}
	}

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

	// Compute reverse map for named struct Metronome_note
	// insertion point per field
	stage.Metronome_note_Metronome_beam_reverseMap = make(map[*Metronome_beam]*Metronome_note)
	for metronome_note := range stage.Metronome_notes {
		_ = metronome_note
		for _, _metronome_beam := range metronome_note.Metronome_beam {
			stage.Metronome_note_Metronome_beam_reverseMap[_metronome_beam] = metronome_note
		}
	}

	// Compute reverse map for named struct Miscellaneous
	// insertion point per field
	stage.Miscellaneous_Miscellaneous_field_reverseMap = make(map[*Miscellaneous_field]*Miscellaneous)
	for miscellaneous := range stage.Miscellaneouss {
		_ = miscellaneous
		for _, _miscellaneous_field := range miscellaneous.Miscellaneous_field {
			stage.Miscellaneous_Miscellaneous_field_reverseMap[_miscellaneous_field] = miscellaneous
		}
	}

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

	// Compute reverse map for named struct Part_link
	// insertion point per field
	stage.Part_link_Instrument_link_reverseMap = make(map[*Instrument_link]*Part_link)
	for part_link := range stage.Part_links {
		_ = part_link
		for _, _instrument_link := range part_link.Instrument_link {
			stage.Part_link_Instrument_link_reverseMap[_instrument_link] = part_link
		}
	}

	// Compute reverse map for named struct Play
	// insertion point per field
	stage.Play_Other_play_reverseMap = make(map[*Other_play]*Play)
	for play := range stage.Plays {
		_ = play
		for _, _other_play := range play.Other_play {
			stage.Play_Other_play_reverseMap[_other_play] = play
		}
	}

	// Compute reverse map for named struct Print
	// insertion point per field
	stage.Print_Staff_layout_reverseMap = make(map[*Staff_layout]*Print)
	for print := range stage.Prints {
		_ = print
		for _, _staff_layout := range print.Staff_layout {
			stage.Print_Staff_layout_reverseMap[_staff_layout] = print
		}
	}

	// Compute reverse map for named struct Scordatura
	// insertion point per field
	stage.Scordatura_Accord_reverseMap = make(map[*Accord]*Scordatura)
	for scordatura := range stage.Scordaturas {
		_ = scordatura
		for _, _accord := range scordatura.Accord {
			stage.Scordatura_Accord_reverseMap[_accord] = scordatura
		}
	}

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

	// Compute reverse map for named struct Tuplet_portion
	// insertion point per field
	stage.Tuplet_portion_Tuplet_dot_reverseMap = make(map[*Tuplet_dot]*Tuplet_portion)
	for tuplet_portion := range stage.Tuplet_portions {
		_ = tuplet_portion
		for _, _tuplet_dot := range tuplet_portion.Tuplet_dot {
			stage.Tuplet_portion_Tuplet_dot_reverseMap[_tuplet_dot] = tuplet_portion
		}
	}

	// end of insertion point per named struct
}

func (stage *Stage) GetInstances() (res []GongstructIF) {
	// insertion point per named struct
	res = __gong__appendInstances(res, stage.A_directives)

	res = __gong__appendInstances(res, stage.A_measures)

	res = __gong__appendInstances(res, stage.A_measure_1s)

	res = __gong__appendInstances(res, stage.A_parts)

	res = __gong__appendInstances(res, stage.A_part_1s)

	res = __gong__appendInstances(res, stage.Accidentals)

	res = __gong__appendInstances(res, stage.Accidental_marks)

	res = __gong__appendInstances(res, stage.Accidental_texts)

	res = __gong__appendInstances(res, stage.Accords)

	res = __gong__appendInstances(res, stage.Accordion_registrations)

	res = __gong__appendInstances(res, stage.Appearances)

	res = __gong__appendInstances(res, stage.Arpeggiates)

	res = __gong__appendInstances(res, stage.Arrows)

	res = __gong__appendInstances(res, stage.Articulationss)

	res = __gong__appendInstances(res, stage.Assesss)

	res = __gong__appendInstances(res, stage.Attributess)

	res = __gong__appendInstances(res, stage.Backups)

	res = __gong__appendInstances(res, stage.Bar_style_colors)

	res = __gong__appendInstances(res, stage.Barlines)

	res = __gong__appendInstances(res, stage.Barres)

	res = __gong__appendInstances(res, stage.Basss)

	res = __gong__appendInstances(res, stage.Bass_steps)

	res = __gong__appendInstances(res, stage.Beams)

	res = __gong__appendInstances(res, stage.Beat_repeats)

	res = __gong__appendInstances(res, stage.Beat_unit_tieds)

	res = __gong__appendInstances(res, stage.Beaters)

	res = __gong__appendInstances(res, stage.Bends)

	res = __gong__appendInstances(res, stage.Bookmarks)

	res = __gong__appendInstances(res, stage.Brackets)

	res = __gong__appendInstances(res, stage.Breath_marks)

	res = __gong__appendInstances(res, stage.Caesuras)

	res = __gong__appendInstances(res, stage.Cancels)

	res = __gong__appendInstances(res, stage.Clefs)

	res = __gong__appendInstances(res, stage.Codas)

	res = __gong__appendInstances(res, stage.Credits)

	res = __gong__appendInstances(res, stage.Dashess)

	res = __gong__appendInstances(res, stage.Defaultss)

	res = __gong__appendInstances(res, stage.Degrees)

	res = __gong__appendInstances(res, stage.Degree_alters)

	res = __gong__appendInstances(res, stage.Degree_types)

	res = __gong__appendInstances(res, stage.Degree_values)

	res = __gong__appendInstances(res, stage.Directions)

	res = __gong__appendInstances(res, stage.Direction_types)

	res = __gong__appendInstances(res, stage.Distances)

	res = __gong__appendInstances(res, stage.Doubles)

	res = __gong__appendInstances(res, stage.Dynamicss)

	res = __gong__appendInstances(res, stage.Effects)

	res = __gong__appendInstances(res, stage.Elisions)

	res = __gong__appendInstances(res, stage.Emptys)

	res = __gong__appendInstances(res, stage.Empty_fonts)

	res = __gong__appendInstances(res, stage.Empty_lines)

	res = __gong__appendInstances(res, stage.Empty_placements)

	res = __gong__appendInstances(res, stage.Empty_placement_smufls)

	res = __gong__appendInstances(res, stage.Empty_print_object_style_aligns)

	res = __gong__appendInstances(res, stage.Empty_print_styles)

	res = __gong__appendInstances(res, stage.Empty_print_style_aligns)

	res = __gong__appendInstances(res, stage.Empty_print_style_align_ids)

	res = __gong__appendInstances(res, stage.Empty_trill_sounds)

	res = __gong__appendInstances(res, stage.Encodings)

	res = __gong__appendInstances(res, stage.Endings)

	res = __gong__appendInstances(res, stage.Extends)

	res = __gong__appendInstances(res, stage.Features)

	res = __gong__appendInstances(res, stage.Fermatas)

	res = __gong__appendInstances(res, stage.Figures)

	res = __gong__appendInstances(res, stage.Figured_basss)

	res = __gong__appendInstances(res, stage.Fingerings)

	res = __gong__appendInstances(res, stage.First_frets)

	res = __gong__appendInstances(res, stage.For_parts)

	res = __gong__appendInstances(res, stage.Formatted_symbols)

	res = __gong__appendInstances(res, stage.Formatted_symbol_ids)

	res = __gong__appendInstances(res, stage.Formatted_texts)

	res = __gong__appendInstances(res, stage.Formatted_text_ids)

	res = __gong__appendInstances(res, stage.Forwards)

	res = __gong__appendInstances(res, stage.Frames)

	res = __gong__appendInstances(res, stage.Frame_notes)

	res = __gong__appendInstances(res, stage.Frets)

	res = __gong__appendInstances(res, stage.Glasss)

	res = __gong__appendInstances(res, stage.Glissandos)

	res = __gong__appendInstances(res, stage.Glyphs)

	res = __gong__appendInstances(res, stage.Graces)

	res = __gong__appendInstances(res, stage.Group_barlines)

	res = __gong__appendInstances(res, stage.Group_names)

	res = __gong__appendInstances(res, stage.Group_symbols)

	res = __gong__appendInstances(res, stage.Groupings)

	res = __gong__appendInstances(res, stage.Hammer_on_pull_offs)

	res = __gong__appendInstances(res, stage.Handbells)

	res = __gong__appendInstances(res, stage.Harmon_closeds)

	res = __gong__appendInstances(res, stage.Harmon_mutes)

	res = __gong__appendInstances(res, stage.Harmonics)

	res = __gong__appendInstances(res, stage.Harmonys)

	res = __gong__appendInstances(res, stage.Harmony_alters)

	res = __gong__appendInstances(res, stage.Harp_pedalss)

	res = __gong__appendInstances(res, stage.Heel_toes)

	res = __gong__appendInstances(res, stage.Holes)

	res = __gong__appendInstances(res, stage.Hole_closeds)

	res = __gong__appendInstances(res, stage.Horizontal_turns)

	res = __gong__appendInstances(res, stage.Identifications)

	res = __gong__appendInstances(res, stage.Images)

	res = __gong__appendInstances(res, stage.Instruments)

	res = __gong__appendInstances(res, stage.Instrument_changes)

	res = __gong__appendInstances(res, stage.Instrument_links)

	res = __gong__appendInstances(res, stage.Interchangeables)

	res = __gong__appendInstances(res, stage.Inversions)

	res = __gong__appendInstances(res, stage.Keys)

	res = __gong__appendInstances(res, stage.Key_accidentals)

	res = __gong__appendInstances(res, stage.Key_octaves)

	res = __gong__appendInstances(res, stage.Kinds)

	res = __gong__appendInstances(res, stage.Levels)

	res = __gong__appendInstances(res, stage.Line_details)

	res = __gong__appendInstances(res, stage.Line_widths)

	res = __gong__appendInstances(res, stage.Links)

	res = __gong__appendInstances(res, stage.Listens)

	res = __gong__appendInstances(res, stage.Listenings)

	res = __gong__appendInstances(res, stage.Lyrics)

	res = __gong__appendInstances(res, stage.Lyric_fonts)

	res = __gong__appendInstances(res, stage.Lyric_languages)

	res = __gong__appendInstances(res, stage.Measure_layouts)

	res = __gong__appendInstances(res, stage.Measure_numberings)

	res = __gong__appendInstances(res, stage.Measure_repeats)

	res = __gong__appendInstances(res, stage.Measure_styles)

	res = __gong__appendInstances(res, stage.Membranes)

	res = __gong__appendInstances(res, stage.Metals)

	res = __gong__appendInstances(res, stage.Metronomes)

	res = __gong__appendInstances(res, stage.Metronome_beams)

	res = __gong__appendInstances(res, stage.Metronome_notes)

	res = __gong__appendInstances(res, stage.Metronome_tieds)

	res = __gong__appendInstances(res, stage.Metronome_tuplets)

	res = __gong__appendInstances(res, stage.Midi_devices)

	res = __gong__appendInstances(res, stage.Midi_instruments)

	res = __gong__appendInstances(res, stage.Miscellaneouss)

	res = __gong__appendInstances(res, stage.Miscellaneous_fields)

	res = __gong__appendInstances(res, stage.Mordents)

	res = __gong__appendInstances(res, stage.Multiple_rests)

	res = __gong__appendInstances(res, stage.Name_displays)

	res = __gong__appendInstances(res, stage.Non_arpeggiates)

	res = __gong__appendInstances(res, stage.Notationss)

	res = __gong__appendInstances(res, stage.Notes)

	res = __gong__appendInstances(res, stage.Note_sizes)

	res = __gong__appendInstances(res, stage.Note_types)

	res = __gong__appendInstances(res, stage.Noteheads)

	res = __gong__appendInstances(res, stage.Notehead_texts)

	res = __gong__appendInstances(res, stage.Numerals)

	res = __gong__appendInstances(res, stage.Numeral_keys)

	res = __gong__appendInstances(res, stage.Numeral_roots)

	res = __gong__appendInstances(res, stage.Octave_shifts)

	res = __gong__appendInstances(res, stage.Offsets)

	res = __gong__appendInstances(res, stage.Opuss)

	res = __gong__appendInstances(res, stage.Ornamentss)

	res = __gong__appendInstances(res, stage.Other_appearances)

	res = __gong__appendInstances(res, stage.Other_directions)

	res = __gong__appendInstances(res, stage.Other_listenings)

	res = __gong__appendInstances(res, stage.Other_notations)

	res = __gong__appendInstances(res, stage.Other_placement_texts)

	res = __gong__appendInstances(res, stage.Other_plays)

	res = __gong__appendInstances(res, stage.Other_texts)

	res = __gong__appendInstances(res, stage.Page_layouts)

	res = __gong__appendInstances(res, stage.Page_marginss)

	res = __gong__appendInstances(res, stage.Part_clefs)

	res = __gong__appendInstances(res, stage.Part_groups)

	res = __gong__appendInstances(res, stage.Part_links)

	res = __gong__appendInstances(res, stage.Part_lists)

	res = __gong__appendInstances(res, stage.Part_names)

	res = __gong__appendInstances(res, stage.Part_symbols)

	res = __gong__appendInstances(res, stage.Part_transposes)

	res = __gong__appendInstances(res, stage.Pedals)

	res = __gong__appendInstances(res, stage.Pedal_tunings)

	res = __gong__appendInstances(res, stage.Per_minutes)

	res = __gong__appendInstances(res, stage.Percussions)

	res = __gong__appendInstances(res, stage.Pitchs)

	res = __gong__appendInstances(res, stage.Pitcheds)

	res = __gong__appendInstances(res, stage.Placement_texts)

	res = __gong__appendInstances(res, stage.Plays)

	res = __gong__appendInstances(res, stage.Players)

	res = __gong__appendInstances(res, stage.Principal_voices)

	res = __gong__appendInstances(res, stage.Prints)

	res = __gong__appendInstances(res, stage.Releases)

	res = __gong__appendInstances(res, stage.Repeats)

	res = __gong__appendInstances(res, stage.Rests)

	res = __gong__appendInstances(res, stage.Roots)

	res = __gong__appendInstances(res, stage.Root_steps)

	res = __gong__appendInstances(res, stage.Scalings)

	res = __gong__appendInstances(res, stage.Scordaturas)

	res = __gong__appendInstances(res, stage.Score_instruments)

	res = __gong__appendInstances(res, stage.Score_parts)

	res = __gong__appendInstances(res, stage.Score_partwises)

	res = __gong__appendInstances(res, stage.Score_timewises)

	res = __gong__appendInstances(res, stage.Segnos)

	res = __gong__appendInstances(res, stage.Slashs)

	res = __gong__appendInstances(res, stage.Slides)

	res = __gong__appendInstances(res, stage.Slurs)

	res = __gong__appendInstances(res, stage.Sounds)

	res = __gong__appendInstances(res, stage.Staff_detailss)

	res = __gong__appendInstances(res, stage.Staff_divides)

	res = __gong__appendInstances(res, stage.Staff_layouts)

	res = __gong__appendInstances(res, stage.Staff_sizes)

	res = __gong__appendInstances(res, stage.Staff_tunings)

	res = __gong__appendInstances(res, stage.Stems)

	res = __gong__appendInstances(res, stage.Sticks)

	res = __gong__appendInstances(res, stage.String_mutes)

	res = __gong__appendInstances(res, stage.String_types)

	res = __gong__appendInstances(res, stage.Strong_accents)

	res = __gong__appendInstances(res, stage.Style_texts)

	res = __gong__appendInstances(res, stage.Supportss)

	res = __gong__appendInstances(res, stage.Swings)

	res = __gong__appendInstances(res, stage.Syncs)

	res = __gong__appendInstances(res, stage.System_dividerss)

	res = __gong__appendInstances(res, stage.System_layouts)

	res = __gong__appendInstances(res, stage.System_marginss)

	res = __gong__appendInstances(res, stage.Taps)

	res = __gong__appendInstances(res, stage.Technicals)

	res = __gong__appendInstances(res, stage.Text_element_datas)

	res = __gong__appendInstances(res, stage.Ties)

	res = __gong__appendInstances(res, stage.Tieds)

	res = __gong__appendInstances(res, stage.Times)

	res = __gong__appendInstances(res, stage.Time_modifications)

	res = __gong__appendInstances(res, stage.Timpanis)

	res = __gong__appendInstances(res, stage.Transposes)

	res = __gong__appendInstances(res, stage.Tremolos)

	res = __gong__appendInstances(res, stage.Tuplets)

	res = __gong__appendInstances(res, stage.Tuplet_dots)

	res = __gong__appendInstances(res, stage.Tuplet_numbers)

	res = __gong__appendInstances(res, stage.Tuplet_portions)

	res = __gong__appendInstances(res, stage.Tuplet_types)

	res = __gong__appendInstances(res, stage.Typed_texts)

	res = __gong__appendInstances(res, stage.Unpitcheds)

	res = __gong__appendInstances(res, stage.Virtual_instruments)

	res = __gong__appendInstances(res, stage.Waits)

	res = __gong__appendInstances(res, stage.Wavy_lines)

	res = __gong__appendInstances(res, stage.Wedges)

	res = __gong__appendInstances(res, stage.Woods)

	res = __gong__appendInstances(res, stage.Works)

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
func (a_directive *A_directive) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_directive)
}

func (a_measure *A_measure) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_measure)
}

func (a_measure_1 *A_measure_1) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_measure_1)
}

func (a_part *A_part) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_part)
}

func (a_part_1 *A_part_1) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, a_part_1)
}

func (accidental *Accidental) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, accidental)
}

func (accidental_mark *Accidental_mark) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, accidental_mark)
}

func (accidental_text *Accidental_text) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, accidental_text)
}

func (accord *Accord) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, accord)
}

func (accordion_registration *Accordion_registration) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, accordion_registration)
}

func (appearance *Appearance) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, appearance)
}

func (arpeggiate *Arpeggiate) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, arpeggiate)
}

func (arrow *Arrow) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, arrow)
}

func (articulations *Articulations) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, articulations)
}

func (assess *Assess) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, assess)
}

func (attributes *Attributes) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, attributes)
}

func (backup *Backup) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, backup)
}

func (bar_style_color *Bar_style_color) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, bar_style_color)
}

func (barline *Barline) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, barline)
}

func (barre *Barre) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, barre)
}

func (bass *Bass) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, bass)
}

func (bass_step *Bass_step) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, bass_step)
}

func (beam *Beam) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, beam)
}

func (beat_repeat *Beat_repeat) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, beat_repeat)
}

func (beat_unit_tied *Beat_unit_tied) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, beat_unit_tied)
}

func (beater *Beater) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, beater)
}

func (bend *Bend) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, bend)
}

func (bookmark *Bookmark) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, bookmark)
}

func (bracket *Bracket) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, bracket)
}

func (breath_mark *Breath_mark) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, breath_mark)
}

func (caesura *Caesura) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, caesura)
}

func (cancel *Cancel) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, cancel)
}

func (clef *Clef) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, clef)
}

func (coda *Coda) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, coda)
}

func (credit *Credit) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, credit)
}

func (dashes *Dashes) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, dashes)
}

func (defaults *Defaults) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, defaults)
}

func (degree *Degree) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, degree)
}

func (degree_alter *Degree_alter) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, degree_alter)
}

func (degree_type *Degree_type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, degree_type)
}

func (degree_value *Degree_value) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, degree_value)
}

func (direction *Direction) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, direction)
}

func (direction_type *Direction_type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, direction_type)
}

func (distance *Distance) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, distance)
}

func (double *Double) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, double)
}

func (dynamics *Dynamics) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, dynamics)
}

func (effect *Effect) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, effect)
}

func (elision *Elision) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, elision)
}

func (empty *Empty) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, empty)
}

func (empty_font *Empty_font) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, empty_font)
}

func (empty_line *Empty_line) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, empty_line)
}

func (empty_placement *Empty_placement) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, empty_placement)
}

func (empty_placement_smufl *Empty_placement_smufl) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, empty_placement_smufl)
}

func (empty_print_object_style_align *Empty_print_object_style_align) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, empty_print_object_style_align)
}

func (empty_print_style *Empty_print_style) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, empty_print_style)
}

func (empty_print_style_align *Empty_print_style_align) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, empty_print_style_align)
}

func (empty_print_style_align_id *Empty_print_style_align_id) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, empty_print_style_align_id)
}

func (empty_trill_sound *Empty_trill_sound) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, empty_trill_sound)
}

func (encoding *Encoding) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, encoding)
}

func (ending *Ending) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ending)
}

func (extend *Extend) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, extend)
}

func (feature *Feature) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, feature)
}

func (fermata *Fermata) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, fermata)
}

func (figure *Figure) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, figure)
}

func (figured_bass *Figured_bass) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, figured_bass)
}

func (fingering *Fingering) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, fingering)
}

func (first_fret *First_fret) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, first_fret)
}

func (for_part *For_part) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, for_part)
}

func (formatted_symbol *Formatted_symbol) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formatted_symbol)
}

func (formatted_symbol_id *Formatted_symbol_id) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formatted_symbol_id)
}

func (formatted_text *Formatted_text) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formatted_text)
}

func (formatted_text_id *Formatted_text_id) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, formatted_text_id)
}

func (forward *Forward) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, forward)
}

func (frame *Frame) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, frame)
}

func (frame_note *Frame_note) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, frame_note)
}

func (fret *Fret) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, fret)
}

func (glass *Glass) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, glass)
}

func (glissando *Glissando) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, glissando)
}

func (glyph *Glyph) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, glyph)
}

func (grace *Grace) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, grace)
}

func (group_barline *Group_barline) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, group_barline)
}

func (group_name *Group_name) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, group_name)
}

func (group_symbol *Group_symbol) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, group_symbol)
}

func (grouping *Grouping) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, grouping)
}

func (hammer_on_pull_off *Hammer_on_pull_off) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, hammer_on_pull_off)
}

func (handbell *Handbell) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, handbell)
}

func (harmon_closed *Harmon_closed) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, harmon_closed)
}

func (harmon_mute *Harmon_mute) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, harmon_mute)
}

func (harmonic *Harmonic) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, harmonic)
}

func (harmony *Harmony) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, harmony)
}

func (harmony_alter *Harmony_alter) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, harmony_alter)
}

func (harp_pedals *Harp_pedals) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, harp_pedals)
}

func (heel_toe *Heel_toe) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, heel_toe)
}

func (hole *Hole) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, hole)
}

func (hole_closed *Hole_closed) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, hole_closed)
}

func (horizontal_turn *Horizontal_turn) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, horizontal_turn)
}

func (identification *Identification) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, identification)
}

func (image *Image) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, image)
}

func (instrument *Instrument) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, instrument)
}

func (instrument_change *Instrument_change) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, instrument_change)
}

func (instrument_link *Instrument_link) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, instrument_link)
}

func (interchangeable *Interchangeable) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, interchangeable)
}

func (inversion *Inversion) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, inversion)
}

func (key *Key) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, key)
}

func (key_accidental *Key_accidental) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, key_accidental)
}

func (key_octave *Key_octave) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, key_octave)
}

func (kind *Kind) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, kind)
}

func (level *Level) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, level)
}

func (line_detail *Line_detail) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, line_detail)
}

func (line_width *Line_width) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, line_width)
}

func (link *Link) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, link)
}

func (listen *Listen) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, listen)
}

func (listening *Listening) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, listening)
}

func (lyric *Lyric) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, lyric)
}

func (lyric_font *Lyric_font) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, lyric_font)
}

func (lyric_language *Lyric_language) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, lyric_language)
}

func (measure_layout *Measure_layout) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, measure_layout)
}

func (measure_numbering *Measure_numbering) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, measure_numbering)
}

func (measure_repeat *Measure_repeat) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, measure_repeat)
}

func (measure_style *Measure_style) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, measure_style)
}

func (membrane *Membrane) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, membrane)
}

func (metal *Metal) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, metal)
}

func (metronome *Metronome) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, metronome)
}

func (metronome_beam *Metronome_beam) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, metronome_beam)
}

func (metronome_note *Metronome_note) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, metronome_note)
}

func (metronome_tied *Metronome_tied) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, metronome_tied)
}

func (metronome_tuplet *Metronome_tuplet) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, metronome_tuplet)
}

func (midi_device *Midi_device) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, midi_device)
}

func (midi_instrument *Midi_instrument) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, midi_instrument)
}

func (miscellaneous *Miscellaneous) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, miscellaneous)
}

func (miscellaneous_field *Miscellaneous_field) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, miscellaneous_field)
}

func (mordent *Mordent) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, mordent)
}

func (multiple_rest *Multiple_rest) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, multiple_rest)
}

func (name_display *Name_display) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, name_display)
}

func (non_arpeggiate *Non_arpeggiate) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, non_arpeggiate)
}

func (notations *Notations) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notations)
}

func (note *Note) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, note)
}

func (note_size *Note_size) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, note_size)
}

func (note_type *Note_type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, note_type)
}

func (notehead *Notehead) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notehead)
}

func (notehead_text *Notehead_text) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, notehead_text)
}

func (numeral *Numeral) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, numeral)
}

func (numeral_key *Numeral_key) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, numeral_key)
}

func (numeral_root *Numeral_root) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, numeral_root)
}

func (octave_shift *Octave_shift) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, octave_shift)
}

func (offset *Offset) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, offset)
}

func (opus *Opus) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, opus)
}

func (ornaments *Ornaments) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, ornaments)
}

func (other_appearance *Other_appearance) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, other_appearance)
}

func (other_direction *Other_direction) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, other_direction)
}

func (other_listening *Other_listening) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, other_listening)
}

func (other_notation *Other_notation) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, other_notation)
}

func (other_placement_text *Other_placement_text) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, other_placement_text)
}

func (other_play *Other_play) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, other_play)
}

func (other_text *Other_text) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, other_text)
}

func (page_layout *Page_layout) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, page_layout)
}

func (page_margins *Page_margins) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, page_margins)
}

func (part_clef *Part_clef) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, part_clef)
}

func (part_group *Part_group) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, part_group)
}

func (part_link *Part_link) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, part_link)
}

func (part_list *Part_list) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, part_list)
}

func (part_name *Part_name) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, part_name)
}

func (part_symbol *Part_symbol) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, part_symbol)
}

func (part_transpose *Part_transpose) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, part_transpose)
}

func (pedal *Pedal) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pedal)
}

func (pedal_tuning *Pedal_tuning) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pedal_tuning)
}

func (per_minute *Per_minute) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, per_minute)
}

func (percussion *Percussion) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, percussion)
}

func (pitch *Pitch) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pitch)
}

func (pitched *Pitched) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, pitched)
}

func (placement_text *Placement_text) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, placement_text)
}

func (play *Play) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, play)
}

func (player *Player) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, player)
}

func (principal_voice *Principal_voice) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, principal_voice)
}

func (print *Print) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, print)
}

func (release *Release) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, release)
}

func (repeat *Repeat) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, repeat)
}

func (rest *Rest) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, rest)
}

func (root *Root) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, root)
}

func (root_step *Root_step) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, root_step)
}

func (scaling *Scaling) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, scaling)
}

func (scordatura *Scordatura) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, scordatura)
}

func (score_instrument *Score_instrument) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, score_instrument)
}

func (score_part *Score_part) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, score_part)
}

func (score_partwise *Score_partwise) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, score_partwise)
}

func (score_timewise *Score_timewise) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, score_timewise)
}

func (segno *Segno) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, segno)
}

func (slash *Slash) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, slash)
}

func (slide *Slide) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, slide)
}

func (slur *Slur) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, slur)
}

func (sound *Sound) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, sound)
}

func (staff_details *Staff_details) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, staff_details)
}

func (staff_divide *Staff_divide) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, staff_divide)
}

func (staff_layout *Staff_layout) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, staff_layout)
}

func (staff_size *Staff_size) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, staff_size)
}

func (staff_tuning *Staff_tuning) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, staff_tuning)
}

func (stem *Stem) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stem)
}

func (stick *Stick) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, stick)
}

func (string_mute *String_mute) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, string_mute)
}

func (string_type *String_type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, string_type)
}

func (strong_accent *Strong_accent) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, strong_accent)
}

func (style_text *Style_text) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, style_text)
}

func (supports *Supports) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, supports)
}

func (swing *Swing) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, swing)
}

func (sync *Sync) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, sync)
}

func (system_dividers *System_dividers) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, system_dividers)
}

func (system_layout *System_layout) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, system_layout)
}

func (system_margins *System_margins) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, system_margins)
}

func (tap *Tap) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tap)
}

func (technical *Technical) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, technical)
}

func (text_element_data *Text_element_data) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, text_element_data)
}

func (tie *Tie) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tie)
}

func (tied *Tied) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tied)
}

func (time *Time) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, time)
}

func (time_modification *Time_modification) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, time_modification)
}

func (timpani *Timpani) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, timpani)
}

func (transpose *Transpose) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, transpose)
}

func (tremolo *Tremolo) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tremolo)
}

func (tuplet *Tuplet) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tuplet)
}

func (tuplet_dot *Tuplet_dot) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tuplet_dot)
}

func (tuplet_number *Tuplet_number) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tuplet_number)
}

func (tuplet_portion *Tuplet_portion) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tuplet_portion)
}

func (tuplet_type *Tuplet_type) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, tuplet_type)
}

func (typed_text *Typed_text) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, typed_text)
}

func (unpitched *Unpitched) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, unpitched)
}

func (virtual_instrument *Virtual_instrument) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, virtual_instrument)
}

func (wait *Wait) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, wait)
}

func (wavy_line *Wavy_line) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, wavy_line)
}

func (wedge *Wedge) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, wedge)
}

func (wood *Wood) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, wood)
}

func (work *Work) GongGetUUID(stage *Stage) string {
	return __gong__getUUID(stage, work)
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
	__gong__computeReferencePass1(stage, stage.A_directives, &stage.A_directives_reference, &stage.A_directives_referenceOrder, &stage.A_directives_instance)

	__gong__computeReferencePass1(stage, stage.A_measures, &stage.A_measures_reference, &stage.A_measures_referenceOrder, &stage.A_measures_instance)

	__gong__computeReferencePass1(stage, stage.A_measure_1s, &stage.A_measure_1s_reference, &stage.A_measure_1s_referenceOrder, &stage.A_measure_1s_instance)

	__gong__computeReferencePass1(stage, stage.A_parts, &stage.A_parts_reference, &stage.A_parts_referenceOrder, &stage.A_parts_instance)

	__gong__computeReferencePass1(stage, stage.A_part_1s, &stage.A_part_1s_reference, &stage.A_part_1s_referenceOrder, &stage.A_part_1s_instance)

	__gong__computeReferencePass1(stage, stage.Accidentals, &stage.Accidentals_reference, &stage.Accidentals_referenceOrder, &stage.Accidentals_instance)

	__gong__computeReferencePass1(stage, stage.Accidental_marks, &stage.Accidental_marks_reference, &stage.Accidental_marks_referenceOrder, &stage.Accidental_marks_instance)

	__gong__computeReferencePass1(stage, stage.Accidental_texts, &stage.Accidental_texts_reference, &stage.Accidental_texts_referenceOrder, &stage.Accidental_texts_instance)

	__gong__computeReferencePass1(stage, stage.Accords, &stage.Accords_reference, &stage.Accords_referenceOrder, &stage.Accords_instance)

	__gong__computeReferencePass1(stage, stage.Accordion_registrations, &stage.Accordion_registrations_reference, &stage.Accordion_registrations_referenceOrder, &stage.Accordion_registrations_instance)

	__gong__computeReferencePass1(stage, stage.Appearances, &stage.Appearances_reference, &stage.Appearances_referenceOrder, &stage.Appearances_instance)

	__gong__computeReferencePass1(stage, stage.Arpeggiates, &stage.Arpeggiates_reference, &stage.Arpeggiates_referenceOrder, &stage.Arpeggiates_instance)

	__gong__computeReferencePass1(stage, stage.Arrows, &stage.Arrows_reference, &stage.Arrows_referenceOrder, &stage.Arrows_instance)

	__gong__computeReferencePass1(stage, stage.Articulationss, &stage.Articulationss_reference, &stage.Articulationss_referenceOrder, &stage.Articulationss_instance)

	__gong__computeReferencePass1(stage, stage.Assesss, &stage.Assesss_reference, &stage.Assesss_referenceOrder, &stage.Assesss_instance)

	__gong__computeReferencePass1(stage, stage.Attributess, &stage.Attributess_reference, &stage.Attributess_referenceOrder, &stage.Attributess_instance)

	__gong__computeReferencePass1(stage, stage.Backups, &stage.Backups_reference, &stage.Backups_referenceOrder, &stage.Backups_instance)

	__gong__computeReferencePass1(stage, stage.Bar_style_colors, &stage.Bar_style_colors_reference, &stage.Bar_style_colors_referenceOrder, &stage.Bar_style_colors_instance)

	__gong__computeReferencePass1(stage, stage.Barlines, &stage.Barlines_reference, &stage.Barlines_referenceOrder, &stage.Barlines_instance)

	__gong__computeReferencePass1(stage, stage.Barres, &stage.Barres_reference, &stage.Barres_referenceOrder, &stage.Barres_instance)

	__gong__computeReferencePass1(stage, stage.Basss, &stage.Basss_reference, &stage.Basss_referenceOrder, &stage.Basss_instance)

	__gong__computeReferencePass1(stage, stage.Bass_steps, &stage.Bass_steps_reference, &stage.Bass_steps_referenceOrder, &stage.Bass_steps_instance)

	__gong__computeReferencePass1(stage, stage.Beams, &stage.Beams_reference, &stage.Beams_referenceOrder, &stage.Beams_instance)

	__gong__computeReferencePass1(stage, stage.Beat_repeats, &stage.Beat_repeats_reference, &stage.Beat_repeats_referenceOrder, &stage.Beat_repeats_instance)

	__gong__computeReferencePass1(stage, stage.Beat_unit_tieds, &stage.Beat_unit_tieds_reference, &stage.Beat_unit_tieds_referenceOrder, &stage.Beat_unit_tieds_instance)

	__gong__computeReferencePass1(stage, stage.Beaters, &stage.Beaters_reference, &stage.Beaters_referenceOrder, &stage.Beaters_instance)

	__gong__computeReferencePass1(stage, stage.Bends, &stage.Bends_reference, &stage.Bends_referenceOrder, &stage.Bends_instance)

	__gong__computeReferencePass1(stage, stage.Bookmarks, &stage.Bookmarks_reference, &stage.Bookmarks_referenceOrder, &stage.Bookmarks_instance)

	__gong__computeReferencePass1(stage, stage.Brackets, &stage.Brackets_reference, &stage.Brackets_referenceOrder, &stage.Brackets_instance)

	__gong__computeReferencePass1(stage, stage.Breath_marks, &stage.Breath_marks_reference, &stage.Breath_marks_referenceOrder, &stage.Breath_marks_instance)

	__gong__computeReferencePass1(stage, stage.Caesuras, &stage.Caesuras_reference, &stage.Caesuras_referenceOrder, &stage.Caesuras_instance)

	__gong__computeReferencePass1(stage, stage.Cancels, &stage.Cancels_reference, &stage.Cancels_referenceOrder, &stage.Cancels_instance)

	__gong__computeReferencePass1(stage, stage.Clefs, &stage.Clefs_reference, &stage.Clefs_referenceOrder, &stage.Clefs_instance)

	__gong__computeReferencePass1(stage, stage.Codas, &stage.Codas_reference, &stage.Codas_referenceOrder, &stage.Codas_instance)

	__gong__computeReferencePass1(stage, stage.Credits, &stage.Credits_reference, &stage.Credits_referenceOrder, &stage.Credits_instance)

	__gong__computeReferencePass1(stage, stage.Dashess, &stage.Dashess_reference, &stage.Dashess_referenceOrder, &stage.Dashess_instance)

	__gong__computeReferencePass1(stage, stage.Defaultss, &stage.Defaultss_reference, &stage.Defaultss_referenceOrder, &stage.Defaultss_instance)

	__gong__computeReferencePass1(stage, stage.Degrees, &stage.Degrees_reference, &stage.Degrees_referenceOrder, &stage.Degrees_instance)

	__gong__computeReferencePass1(stage, stage.Degree_alters, &stage.Degree_alters_reference, &stage.Degree_alters_referenceOrder, &stage.Degree_alters_instance)

	__gong__computeReferencePass1(stage, stage.Degree_types, &stage.Degree_types_reference, &stage.Degree_types_referenceOrder, &stage.Degree_types_instance)

	__gong__computeReferencePass1(stage, stage.Degree_values, &stage.Degree_values_reference, &stage.Degree_values_referenceOrder, &stage.Degree_values_instance)

	__gong__computeReferencePass1(stage, stage.Directions, &stage.Directions_reference, &stage.Directions_referenceOrder, &stage.Directions_instance)

	__gong__computeReferencePass1(stage, stage.Direction_types, &stage.Direction_types_reference, &stage.Direction_types_referenceOrder, &stage.Direction_types_instance)

	__gong__computeReferencePass1(stage, stage.Distances, &stage.Distances_reference, &stage.Distances_referenceOrder, &stage.Distances_instance)

	__gong__computeReferencePass1(stage, stage.Doubles, &stage.Doubles_reference, &stage.Doubles_referenceOrder, &stage.Doubles_instance)

	__gong__computeReferencePass1(stage, stage.Dynamicss, &stage.Dynamicss_reference, &stage.Dynamicss_referenceOrder, &stage.Dynamicss_instance)

	__gong__computeReferencePass1(stage, stage.Effects, &stage.Effects_reference, &stage.Effects_referenceOrder, &stage.Effects_instance)

	__gong__computeReferencePass1(stage, stage.Elisions, &stage.Elisions_reference, &stage.Elisions_referenceOrder, &stage.Elisions_instance)

	__gong__computeReferencePass1(stage, stage.Emptys, &stage.Emptys_reference, &stage.Emptys_referenceOrder, &stage.Emptys_instance)

	__gong__computeReferencePass1(stage, stage.Empty_fonts, &stage.Empty_fonts_reference, &stage.Empty_fonts_referenceOrder, &stage.Empty_fonts_instance)

	__gong__computeReferencePass1(stage, stage.Empty_lines, &stage.Empty_lines_reference, &stage.Empty_lines_referenceOrder, &stage.Empty_lines_instance)

	__gong__computeReferencePass1(stage, stage.Empty_placements, &stage.Empty_placements_reference, &stage.Empty_placements_referenceOrder, &stage.Empty_placements_instance)

	__gong__computeReferencePass1(stage, stage.Empty_placement_smufls, &stage.Empty_placement_smufls_reference, &stage.Empty_placement_smufls_referenceOrder, &stage.Empty_placement_smufls_instance)

	__gong__computeReferencePass1(stage, stage.Empty_print_object_style_aligns, &stage.Empty_print_object_style_aligns_reference, &stage.Empty_print_object_style_aligns_referenceOrder, &stage.Empty_print_object_style_aligns_instance)

	__gong__computeReferencePass1(stage, stage.Empty_print_styles, &stage.Empty_print_styles_reference, &stage.Empty_print_styles_referenceOrder, &stage.Empty_print_styles_instance)

	__gong__computeReferencePass1(stage, stage.Empty_print_style_aligns, &stage.Empty_print_style_aligns_reference, &stage.Empty_print_style_aligns_referenceOrder, &stage.Empty_print_style_aligns_instance)

	__gong__computeReferencePass1(stage, stage.Empty_print_style_align_ids, &stage.Empty_print_style_align_ids_reference, &stage.Empty_print_style_align_ids_referenceOrder, &stage.Empty_print_style_align_ids_instance)

	__gong__computeReferencePass1(stage, stage.Empty_trill_sounds, &stage.Empty_trill_sounds_reference, &stage.Empty_trill_sounds_referenceOrder, &stage.Empty_trill_sounds_instance)

	__gong__computeReferencePass1(stage, stage.Encodings, &stage.Encodings_reference, &stage.Encodings_referenceOrder, &stage.Encodings_instance)

	__gong__computeReferencePass1(stage, stage.Endings, &stage.Endings_reference, &stage.Endings_referenceOrder, &stage.Endings_instance)

	__gong__computeReferencePass1(stage, stage.Extends, &stage.Extends_reference, &stage.Extends_referenceOrder, &stage.Extends_instance)

	__gong__computeReferencePass1(stage, stage.Features, &stage.Features_reference, &stage.Features_referenceOrder, &stage.Features_instance)

	__gong__computeReferencePass1(stage, stage.Fermatas, &stage.Fermatas_reference, &stage.Fermatas_referenceOrder, &stage.Fermatas_instance)

	__gong__computeReferencePass1(stage, stage.Figures, &stage.Figures_reference, &stage.Figures_referenceOrder, &stage.Figures_instance)

	__gong__computeReferencePass1(stage, stage.Figured_basss, &stage.Figured_basss_reference, &stage.Figured_basss_referenceOrder, &stage.Figured_basss_instance)

	__gong__computeReferencePass1(stage, stage.Fingerings, &stage.Fingerings_reference, &stage.Fingerings_referenceOrder, &stage.Fingerings_instance)

	__gong__computeReferencePass1(stage, stage.First_frets, &stage.First_frets_reference, &stage.First_frets_referenceOrder, &stage.First_frets_instance)

	__gong__computeReferencePass1(stage, stage.For_parts, &stage.For_parts_reference, &stage.For_parts_referenceOrder, &stage.For_parts_instance)

	__gong__computeReferencePass1(stage, stage.Formatted_symbols, &stage.Formatted_symbols_reference, &stage.Formatted_symbols_referenceOrder, &stage.Formatted_symbols_instance)

	__gong__computeReferencePass1(stage, stage.Formatted_symbol_ids, &stage.Formatted_symbol_ids_reference, &stage.Formatted_symbol_ids_referenceOrder, &stage.Formatted_symbol_ids_instance)

	__gong__computeReferencePass1(stage, stage.Formatted_texts, &stage.Formatted_texts_reference, &stage.Formatted_texts_referenceOrder, &stage.Formatted_texts_instance)

	__gong__computeReferencePass1(stage, stage.Formatted_text_ids, &stage.Formatted_text_ids_reference, &stage.Formatted_text_ids_referenceOrder, &stage.Formatted_text_ids_instance)

	__gong__computeReferencePass1(stage, stage.Forwards, &stage.Forwards_reference, &stage.Forwards_referenceOrder, &stage.Forwards_instance)

	__gong__computeReferencePass1(stage, stage.Frames, &stage.Frames_reference, &stage.Frames_referenceOrder, &stage.Frames_instance)

	__gong__computeReferencePass1(stage, stage.Frame_notes, &stage.Frame_notes_reference, &stage.Frame_notes_referenceOrder, &stage.Frame_notes_instance)

	__gong__computeReferencePass1(stage, stage.Frets, &stage.Frets_reference, &stage.Frets_referenceOrder, &stage.Frets_instance)

	__gong__computeReferencePass1(stage, stage.Glasss, &stage.Glasss_reference, &stage.Glasss_referenceOrder, &stage.Glasss_instance)

	__gong__computeReferencePass1(stage, stage.Glissandos, &stage.Glissandos_reference, &stage.Glissandos_referenceOrder, &stage.Glissandos_instance)

	__gong__computeReferencePass1(stage, stage.Glyphs, &stage.Glyphs_reference, &stage.Glyphs_referenceOrder, &stage.Glyphs_instance)

	__gong__computeReferencePass1(stage, stage.Graces, &stage.Graces_reference, &stage.Graces_referenceOrder, &stage.Graces_instance)

	__gong__computeReferencePass1(stage, stage.Group_barlines, &stage.Group_barlines_reference, &stage.Group_barlines_referenceOrder, &stage.Group_barlines_instance)

	__gong__computeReferencePass1(stage, stage.Group_names, &stage.Group_names_reference, &stage.Group_names_referenceOrder, &stage.Group_names_instance)

	__gong__computeReferencePass1(stage, stage.Group_symbols, &stage.Group_symbols_reference, &stage.Group_symbols_referenceOrder, &stage.Group_symbols_instance)

	__gong__computeReferencePass1(stage, stage.Groupings, &stage.Groupings_reference, &stage.Groupings_referenceOrder, &stage.Groupings_instance)

	__gong__computeReferencePass1(stage, stage.Hammer_on_pull_offs, &stage.Hammer_on_pull_offs_reference, &stage.Hammer_on_pull_offs_referenceOrder, &stage.Hammer_on_pull_offs_instance)

	__gong__computeReferencePass1(stage, stage.Handbells, &stage.Handbells_reference, &stage.Handbells_referenceOrder, &stage.Handbells_instance)

	__gong__computeReferencePass1(stage, stage.Harmon_closeds, &stage.Harmon_closeds_reference, &stage.Harmon_closeds_referenceOrder, &stage.Harmon_closeds_instance)

	__gong__computeReferencePass1(stage, stage.Harmon_mutes, &stage.Harmon_mutes_reference, &stage.Harmon_mutes_referenceOrder, &stage.Harmon_mutes_instance)

	__gong__computeReferencePass1(stage, stage.Harmonics, &stage.Harmonics_reference, &stage.Harmonics_referenceOrder, &stage.Harmonics_instance)

	__gong__computeReferencePass1(stage, stage.Harmonys, &stage.Harmonys_reference, &stage.Harmonys_referenceOrder, &stage.Harmonys_instance)

	__gong__computeReferencePass1(stage, stage.Harmony_alters, &stage.Harmony_alters_reference, &stage.Harmony_alters_referenceOrder, &stage.Harmony_alters_instance)

	__gong__computeReferencePass1(stage, stage.Harp_pedalss, &stage.Harp_pedalss_reference, &stage.Harp_pedalss_referenceOrder, &stage.Harp_pedalss_instance)

	__gong__computeReferencePass1(stage, stage.Heel_toes, &stage.Heel_toes_reference, &stage.Heel_toes_referenceOrder, &stage.Heel_toes_instance)

	__gong__computeReferencePass1(stage, stage.Holes, &stage.Holes_reference, &stage.Holes_referenceOrder, &stage.Holes_instance)

	__gong__computeReferencePass1(stage, stage.Hole_closeds, &stage.Hole_closeds_reference, &stage.Hole_closeds_referenceOrder, &stage.Hole_closeds_instance)

	__gong__computeReferencePass1(stage, stage.Horizontal_turns, &stage.Horizontal_turns_reference, &stage.Horizontal_turns_referenceOrder, &stage.Horizontal_turns_instance)

	__gong__computeReferencePass1(stage, stage.Identifications, &stage.Identifications_reference, &stage.Identifications_referenceOrder, &stage.Identifications_instance)

	__gong__computeReferencePass1(stage, stage.Images, &stage.Images_reference, &stage.Images_referenceOrder, &stage.Images_instance)

	__gong__computeReferencePass1(stage, stage.Instruments, &stage.Instruments_reference, &stage.Instruments_referenceOrder, &stage.Instruments_instance)

	__gong__computeReferencePass1(stage, stage.Instrument_changes, &stage.Instrument_changes_reference, &stage.Instrument_changes_referenceOrder, &stage.Instrument_changes_instance)

	__gong__computeReferencePass1(stage, stage.Instrument_links, &stage.Instrument_links_reference, &stage.Instrument_links_referenceOrder, &stage.Instrument_links_instance)

	__gong__computeReferencePass1(stage, stage.Interchangeables, &stage.Interchangeables_reference, &stage.Interchangeables_referenceOrder, &stage.Interchangeables_instance)

	__gong__computeReferencePass1(stage, stage.Inversions, &stage.Inversions_reference, &stage.Inversions_referenceOrder, &stage.Inversions_instance)

	__gong__computeReferencePass1(stage, stage.Keys, &stage.Keys_reference, &stage.Keys_referenceOrder, &stage.Keys_instance)

	__gong__computeReferencePass1(stage, stage.Key_accidentals, &stage.Key_accidentals_reference, &stage.Key_accidentals_referenceOrder, &stage.Key_accidentals_instance)

	__gong__computeReferencePass1(stage, stage.Key_octaves, &stage.Key_octaves_reference, &stage.Key_octaves_referenceOrder, &stage.Key_octaves_instance)

	__gong__computeReferencePass1(stage, stage.Kinds, &stage.Kinds_reference, &stage.Kinds_referenceOrder, &stage.Kinds_instance)

	__gong__computeReferencePass1(stage, stage.Levels, &stage.Levels_reference, &stage.Levels_referenceOrder, &stage.Levels_instance)

	__gong__computeReferencePass1(stage, stage.Line_details, &stage.Line_details_reference, &stage.Line_details_referenceOrder, &stage.Line_details_instance)

	__gong__computeReferencePass1(stage, stage.Line_widths, &stage.Line_widths_reference, &stage.Line_widths_referenceOrder, &stage.Line_widths_instance)

	__gong__computeReferencePass1(stage, stage.Links, &stage.Links_reference, &stage.Links_referenceOrder, &stage.Links_instance)

	__gong__computeReferencePass1(stage, stage.Listens, &stage.Listens_reference, &stage.Listens_referenceOrder, &stage.Listens_instance)

	__gong__computeReferencePass1(stage, stage.Listenings, &stage.Listenings_reference, &stage.Listenings_referenceOrder, &stage.Listenings_instance)

	__gong__computeReferencePass1(stage, stage.Lyrics, &stage.Lyrics_reference, &stage.Lyrics_referenceOrder, &stage.Lyrics_instance)

	__gong__computeReferencePass1(stage, stage.Lyric_fonts, &stage.Lyric_fonts_reference, &stage.Lyric_fonts_referenceOrder, &stage.Lyric_fonts_instance)

	__gong__computeReferencePass1(stage, stage.Lyric_languages, &stage.Lyric_languages_reference, &stage.Lyric_languages_referenceOrder, &stage.Lyric_languages_instance)

	__gong__computeReferencePass1(stage, stage.Measure_layouts, &stage.Measure_layouts_reference, &stage.Measure_layouts_referenceOrder, &stage.Measure_layouts_instance)

	__gong__computeReferencePass1(stage, stage.Measure_numberings, &stage.Measure_numberings_reference, &stage.Measure_numberings_referenceOrder, &stage.Measure_numberings_instance)

	__gong__computeReferencePass1(stage, stage.Measure_repeats, &stage.Measure_repeats_reference, &stage.Measure_repeats_referenceOrder, &stage.Measure_repeats_instance)

	__gong__computeReferencePass1(stage, stage.Measure_styles, &stage.Measure_styles_reference, &stage.Measure_styles_referenceOrder, &stage.Measure_styles_instance)

	__gong__computeReferencePass1(stage, stage.Membranes, &stage.Membranes_reference, &stage.Membranes_referenceOrder, &stage.Membranes_instance)

	__gong__computeReferencePass1(stage, stage.Metals, &stage.Metals_reference, &stage.Metals_referenceOrder, &stage.Metals_instance)

	__gong__computeReferencePass1(stage, stage.Metronomes, &stage.Metronomes_reference, &stage.Metronomes_referenceOrder, &stage.Metronomes_instance)

	__gong__computeReferencePass1(stage, stage.Metronome_beams, &stage.Metronome_beams_reference, &stage.Metronome_beams_referenceOrder, &stage.Metronome_beams_instance)

	__gong__computeReferencePass1(stage, stage.Metronome_notes, &stage.Metronome_notes_reference, &stage.Metronome_notes_referenceOrder, &stage.Metronome_notes_instance)

	__gong__computeReferencePass1(stage, stage.Metronome_tieds, &stage.Metronome_tieds_reference, &stage.Metronome_tieds_referenceOrder, &stage.Metronome_tieds_instance)

	__gong__computeReferencePass1(stage, stage.Metronome_tuplets, &stage.Metronome_tuplets_reference, &stage.Metronome_tuplets_referenceOrder, &stage.Metronome_tuplets_instance)

	__gong__computeReferencePass1(stage, stage.Midi_devices, &stage.Midi_devices_reference, &stage.Midi_devices_referenceOrder, &stage.Midi_devices_instance)

	__gong__computeReferencePass1(stage, stage.Midi_instruments, &stage.Midi_instruments_reference, &stage.Midi_instruments_referenceOrder, &stage.Midi_instruments_instance)

	__gong__computeReferencePass1(stage, stage.Miscellaneouss, &stage.Miscellaneouss_reference, &stage.Miscellaneouss_referenceOrder, &stage.Miscellaneouss_instance)

	__gong__computeReferencePass1(stage, stage.Miscellaneous_fields, &stage.Miscellaneous_fields_reference, &stage.Miscellaneous_fields_referenceOrder, &stage.Miscellaneous_fields_instance)

	__gong__computeReferencePass1(stage, stage.Mordents, &stage.Mordents_reference, &stage.Mordents_referenceOrder, &stage.Mordents_instance)

	__gong__computeReferencePass1(stage, stage.Multiple_rests, &stage.Multiple_rests_reference, &stage.Multiple_rests_referenceOrder, &stage.Multiple_rests_instance)

	__gong__computeReferencePass1(stage, stage.Name_displays, &stage.Name_displays_reference, &stage.Name_displays_referenceOrder, &stage.Name_displays_instance)

	__gong__computeReferencePass1(stage, stage.Non_arpeggiates, &stage.Non_arpeggiates_reference, &stage.Non_arpeggiates_referenceOrder, &stage.Non_arpeggiates_instance)

	__gong__computeReferencePass1(stage, stage.Notationss, &stage.Notationss_reference, &stage.Notationss_referenceOrder, &stage.Notationss_instance)

	__gong__computeReferencePass1(stage, stage.Notes, &stage.Notes_reference, &stage.Notes_referenceOrder, &stage.Notes_instance)

	__gong__computeReferencePass1(stage, stage.Note_sizes, &stage.Note_sizes_reference, &stage.Note_sizes_referenceOrder, &stage.Note_sizes_instance)

	__gong__computeReferencePass1(stage, stage.Note_types, &stage.Note_types_reference, &stage.Note_types_referenceOrder, &stage.Note_types_instance)

	__gong__computeReferencePass1(stage, stage.Noteheads, &stage.Noteheads_reference, &stage.Noteheads_referenceOrder, &stage.Noteheads_instance)

	__gong__computeReferencePass1(stage, stage.Notehead_texts, &stage.Notehead_texts_reference, &stage.Notehead_texts_referenceOrder, &stage.Notehead_texts_instance)

	__gong__computeReferencePass1(stage, stage.Numerals, &stage.Numerals_reference, &stage.Numerals_referenceOrder, &stage.Numerals_instance)

	__gong__computeReferencePass1(stage, stage.Numeral_keys, &stage.Numeral_keys_reference, &stage.Numeral_keys_referenceOrder, &stage.Numeral_keys_instance)

	__gong__computeReferencePass1(stage, stage.Numeral_roots, &stage.Numeral_roots_reference, &stage.Numeral_roots_referenceOrder, &stage.Numeral_roots_instance)

	__gong__computeReferencePass1(stage, stage.Octave_shifts, &stage.Octave_shifts_reference, &stage.Octave_shifts_referenceOrder, &stage.Octave_shifts_instance)

	__gong__computeReferencePass1(stage, stage.Offsets, &stage.Offsets_reference, &stage.Offsets_referenceOrder, &stage.Offsets_instance)

	__gong__computeReferencePass1(stage, stage.Opuss, &stage.Opuss_reference, &stage.Opuss_referenceOrder, &stage.Opuss_instance)

	__gong__computeReferencePass1(stage, stage.Ornamentss, &stage.Ornamentss_reference, &stage.Ornamentss_referenceOrder, &stage.Ornamentss_instance)

	__gong__computeReferencePass1(stage, stage.Other_appearances, &stage.Other_appearances_reference, &stage.Other_appearances_referenceOrder, &stage.Other_appearances_instance)

	__gong__computeReferencePass1(stage, stage.Other_directions, &stage.Other_directions_reference, &stage.Other_directions_referenceOrder, &stage.Other_directions_instance)

	__gong__computeReferencePass1(stage, stage.Other_listenings, &stage.Other_listenings_reference, &stage.Other_listenings_referenceOrder, &stage.Other_listenings_instance)

	__gong__computeReferencePass1(stage, stage.Other_notations, &stage.Other_notations_reference, &stage.Other_notations_referenceOrder, &stage.Other_notations_instance)

	__gong__computeReferencePass1(stage, stage.Other_placement_texts, &stage.Other_placement_texts_reference, &stage.Other_placement_texts_referenceOrder, &stage.Other_placement_texts_instance)

	__gong__computeReferencePass1(stage, stage.Other_plays, &stage.Other_plays_reference, &stage.Other_plays_referenceOrder, &stage.Other_plays_instance)

	__gong__computeReferencePass1(stage, stage.Other_texts, &stage.Other_texts_reference, &stage.Other_texts_referenceOrder, &stage.Other_texts_instance)

	__gong__computeReferencePass1(stage, stage.Page_layouts, &stage.Page_layouts_reference, &stage.Page_layouts_referenceOrder, &stage.Page_layouts_instance)

	__gong__computeReferencePass1(stage, stage.Page_marginss, &stage.Page_marginss_reference, &stage.Page_marginss_referenceOrder, &stage.Page_marginss_instance)

	__gong__computeReferencePass1(stage, stage.Part_clefs, &stage.Part_clefs_reference, &stage.Part_clefs_referenceOrder, &stage.Part_clefs_instance)

	__gong__computeReferencePass1(stage, stage.Part_groups, &stage.Part_groups_reference, &stage.Part_groups_referenceOrder, &stage.Part_groups_instance)

	__gong__computeReferencePass1(stage, stage.Part_links, &stage.Part_links_reference, &stage.Part_links_referenceOrder, &stage.Part_links_instance)

	__gong__computeReferencePass1(stage, stage.Part_lists, &stage.Part_lists_reference, &stage.Part_lists_referenceOrder, &stage.Part_lists_instance)

	__gong__computeReferencePass1(stage, stage.Part_names, &stage.Part_names_reference, &stage.Part_names_referenceOrder, &stage.Part_names_instance)

	__gong__computeReferencePass1(stage, stage.Part_symbols, &stage.Part_symbols_reference, &stage.Part_symbols_referenceOrder, &stage.Part_symbols_instance)

	__gong__computeReferencePass1(stage, stage.Part_transposes, &stage.Part_transposes_reference, &stage.Part_transposes_referenceOrder, &stage.Part_transposes_instance)

	__gong__computeReferencePass1(stage, stage.Pedals, &stage.Pedals_reference, &stage.Pedals_referenceOrder, &stage.Pedals_instance)

	__gong__computeReferencePass1(stage, stage.Pedal_tunings, &stage.Pedal_tunings_reference, &stage.Pedal_tunings_referenceOrder, &stage.Pedal_tunings_instance)

	__gong__computeReferencePass1(stage, stage.Per_minutes, &stage.Per_minutes_reference, &stage.Per_minutes_referenceOrder, &stage.Per_minutes_instance)

	__gong__computeReferencePass1(stage, stage.Percussions, &stage.Percussions_reference, &stage.Percussions_referenceOrder, &stage.Percussions_instance)

	__gong__computeReferencePass1(stage, stage.Pitchs, &stage.Pitchs_reference, &stage.Pitchs_referenceOrder, &stage.Pitchs_instance)

	__gong__computeReferencePass1(stage, stage.Pitcheds, &stage.Pitcheds_reference, &stage.Pitcheds_referenceOrder, &stage.Pitcheds_instance)

	__gong__computeReferencePass1(stage, stage.Placement_texts, &stage.Placement_texts_reference, &stage.Placement_texts_referenceOrder, &stage.Placement_texts_instance)

	__gong__computeReferencePass1(stage, stage.Plays, &stage.Plays_reference, &stage.Plays_referenceOrder, &stage.Plays_instance)

	__gong__computeReferencePass1(stage, stage.Players, &stage.Players_reference, &stage.Players_referenceOrder, &stage.Players_instance)

	__gong__computeReferencePass1(stage, stage.Principal_voices, &stage.Principal_voices_reference, &stage.Principal_voices_referenceOrder, &stage.Principal_voices_instance)

	__gong__computeReferencePass1(stage, stage.Prints, &stage.Prints_reference, &stage.Prints_referenceOrder, &stage.Prints_instance)

	__gong__computeReferencePass1(stage, stage.Releases, &stage.Releases_reference, &stage.Releases_referenceOrder, &stage.Releases_instance)

	__gong__computeReferencePass1(stage, stage.Repeats, &stage.Repeats_reference, &stage.Repeats_referenceOrder, &stage.Repeats_instance)

	__gong__computeReferencePass1(stage, stage.Rests, &stage.Rests_reference, &stage.Rests_referenceOrder, &stage.Rests_instance)

	__gong__computeReferencePass1(stage, stage.Roots, &stage.Roots_reference, &stage.Roots_referenceOrder, &stage.Roots_instance)

	__gong__computeReferencePass1(stage, stage.Root_steps, &stage.Root_steps_reference, &stage.Root_steps_referenceOrder, &stage.Root_steps_instance)

	__gong__computeReferencePass1(stage, stage.Scalings, &stage.Scalings_reference, &stage.Scalings_referenceOrder, &stage.Scalings_instance)

	__gong__computeReferencePass1(stage, stage.Scordaturas, &stage.Scordaturas_reference, &stage.Scordaturas_referenceOrder, &stage.Scordaturas_instance)

	__gong__computeReferencePass1(stage, stage.Score_instruments, &stage.Score_instruments_reference, &stage.Score_instruments_referenceOrder, &stage.Score_instruments_instance)

	__gong__computeReferencePass1(stage, stage.Score_parts, &stage.Score_parts_reference, &stage.Score_parts_referenceOrder, &stage.Score_parts_instance)

	__gong__computeReferencePass1(stage, stage.Score_partwises, &stage.Score_partwises_reference, &stage.Score_partwises_referenceOrder, &stage.Score_partwises_instance)

	__gong__computeReferencePass1(stage, stage.Score_timewises, &stage.Score_timewises_reference, &stage.Score_timewises_referenceOrder, &stage.Score_timewises_instance)

	__gong__computeReferencePass1(stage, stage.Segnos, &stage.Segnos_reference, &stage.Segnos_referenceOrder, &stage.Segnos_instance)

	__gong__computeReferencePass1(stage, stage.Slashs, &stage.Slashs_reference, &stage.Slashs_referenceOrder, &stage.Slashs_instance)

	__gong__computeReferencePass1(stage, stage.Slides, &stage.Slides_reference, &stage.Slides_referenceOrder, &stage.Slides_instance)

	__gong__computeReferencePass1(stage, stage.Slurs, &stage.Slurs_reference, &stage.Slurs_referenceOrder, &stage.Slurs_instance)

	__gong__computeReferencePass1(stage, stage.Sounds, &stage.Sounds_reference, &stage.Sounds_referenceOrder, &stage.Sounds_instance)

	__gong__computeReferencePass1(stage, stage.Staff_detailss, &stage.Staff_detailss_reference, &stage.Staff_detailss_referenceOrder, &stage.Staff_detailss_instance)

	__gong__computeReferencePass1(stage, stage.Staff_divides, &stage.Staff_divides_reference, &stage.Staff_divides_referenceOrder, &stage.Staff_divides_instance)

	__gong__computeReferencePass1(stage, stage.Staff_layouts, &stage.Staff_layouts_reference, &stage.Staff_layouts_referenceOrder, &stage.Staff_layouts_instance)

	__gong__computeReferencePass1(stage, stage.Staff_sizes, &stage.Staff_sizes_reference, &stage.Staff_sizes_referenceOrder, &stage.Staff_sizes_instance)

	__gong__computeReferencePass1(stage, stage.Staff_tunings, &stage.Staff_tunings_reference, &stage.Staff_tunings_referenceOrder, &stage.Staff_tunings_instance)

	__gong__computeReferencePass1(stage, stage.Stems, &stage.Stems_reference, &stage.Stems_referenceOrder, &stage.Stems_instance)

	__gong__computeReferencePass1(stage, stage.Sticks, &stage.Sticks_reference, &stage.Sticks_referenceOrder, &stage.Sticks_instance)

	__gong__computeReferencePass1(stage, stage.String_mutes, &stage.String_mutes_reference, &stage.String_mutes_referenceOrder, &stage.String_mutes_instance)

	__gong__computeReferencePass1(stage, stage.String_types, &stage.String_types_reference, &stage.String_types_referenceOrder, &stage.String_types_instance)

	__gong__computeReferencePass1(stage, stage.Strong_accents, &stage.Strong_accents_reference, &stage.Strong_accents_referenceOrder, &stage.Strong_accents_instance)

	__gong__computeReferencePass1(stage, stage.Style_texts, &stage.Style_texts_reference, &stage.Style_texts_referenceOrder, &stage.Style_texts_instance)

	__gong__computeReferencePass1(stage, stage.Supportss, &stage.Supportss_reference, &stage.Supportss_referenceOrder, &stage.Supportss_instance)

	__gong__computeReferencePass1(stage, stage.Swings, &stage.Swings_reference, &stage.Swings_referenceOrder, &stage.Swings_instance)

	__gong__computeReferencePass1(stage, stage.Syncs, &stage.Syncs_reference, &stage.Syncs_referenceOrder, &stage.Syncs_instance)

	__gong__computeReferencePass1(stage, stage.System_dividerss, &stage.System_dividerss_reference, &stage.System_dividerss_referenceOrder, &stage.System_dividerss_instance)

	__gong__computeReferencePass1(stage, stage.System_layouts, &stage.System_layouts_reference, &stage.System_layouts_referenceOrder, &stage.System_layouts_instance)

	__gong__computeReferencePass1(stage, stage.System_marginss, &stage.System_marginss_reference, &stage.System_marginss_referenceOrder, &stage.System_marginss_instance)

	__gong__computeReferencePass1(stage, stage.Taps, &stage.Taps_reference, &stage.Taps_referenceOrder, &stage.Taps_instance)

	__gong__computeReferencePass1(stage, stage.Technicals, &stage.Technicals_reference, &stage.Technicals_referenceOrder, &stage.Technicals_instance)

	__gong__computeReferencePass1(stage, stage.Text_element_datas, &stage.Text_element_datas_reference, &stage.Text_element_datas_referenceOrder, &stage.Text_element_datas_instance)

	__gong__computeReferencePass1(stage, stage.Ties, &stage.Ties_reference, &stage.Ties_referenceOrder, &stage.Ties_instance)

	__gong__computeReferencePass1(stage, stage.Tieds, &stage.Tieds_reference, &stage.Tieds_referenceOrder, &stage.Tieds_instance)

	__gong__computeReferencePass1(stage, stage.Times, &stage.Times_reference, &stage.Times_referenceOrder, &stage.Times_instance)

	__gong__computeReferencePass1(stage, stage.Time_modifications, &stage.Time_modifications_reference, &stage.Time_modifications_referenceOrder, &stage.Time_modifications_instance)

	__gong__computeReferencePass1(stage, stage.Timpanis, &stage.Timpanis_reference, &stage.Timpanis_referenceOrder, &stage.Timpanis_instance)

	__gong__computeReferencePass1(stage, stage.Transposes, &stage.Transposes_reference, &stage.Transposes_referenceOrder, &stage.Transposes_instance)

	__gong__computeReferencePass1(stage, stage.Tremolos, &stage.Tremolos_reference, &stage.Tremolos_referenceOrder, &stage.Tremolos_instance)

	__gong__computeReferencePass1(stage, stage.Tuplets, &stage.Tuplets_reference, &stage.Tuplets_referenceOrder, &stage.Tuplets_instance)

	__gong__computeReferencePass1(stage, stage.Tuplet_dots, &stage.Tuplet_dots_reference, &stage.Tuplet_dots_referenceOrder, &stage.Tuplet_dots_instance)

	__gong__computeReferencePass1(stage, stage.Tuplet_numbers, &stage.Tuplet_numbers_reference, &stage.Tuplet_numbers_referenceOrder, &stage.Tuplet_numbers_instance)

	__gong__computeReferencePass1(stage, stage.Tuplet_portions, &stage.Tuplet_portions_reference, &stage.Tuplet_portions_referenceOrder, &stage.Tuplet_portions_instance)

	__gong__computeReferencePass1(stage, stage.Tuplet_types, &stage.Tuplet_types_reference, &stage.Tuplet_types_referenceOrder, &stage.Tuplet_types_instance)

	__gong__computeReferencePass1(stage, stage.Typed_texts, &stage.Typed_texts_reference, &stage.Typed_texts_referenceOrder, &stage.Typed_texts_instance)

	__gong__computeReferencePass1(stage, stage.Unpitcheds, &stage.Unpitcheds_reference, &stage.Unpitcheds_referenceOrder, &stage.Unpitcheds_instance)

	__gong__computeReferencePass1(stage, stage.Virtual_instruments, &stage.Virtual_instruments_reference, &stage.Virtual_instruments_referenceOrder, &stage.Virtual_instruments_instance)

	__gong__computeReferencePass1(stage, stage.Waits, &stage.Waits_reference, &stage.Waits_referenceOrder, &stage.Waits_instance)

	__gong__computeReferencePass1(stage, stage.Wavy_lines, &stage.Wavy_lines_reference, &stage.Wavy_lines_referenceOrder, &stage.Wavy_lines_instance)

	__gong__computeReferencePass1(stage, stage.Wedges, &stage.Wedges_reference, &stage.Wedges_referenceOrder, &stage.Wedges_instance)

	__gong__computeReferencePass1(stage, stage.Woods, &stage.Woods_reference, &stage.Woods_referenceOrder, &stage.Woods_instance)

	__gong__computeReferencePass1(stage, stage.Works, &stage.Works_reference, &stage.Works_referenceOrder, &stage.Works_instance)

	// insertion point per named struct
	__gong__computeReferencePass2(stage.A_directives, stage.A_directives_reference, stage)

	__gong__computeReferencePass2(stage.A_measures, stage.A_measures_reference, stage)

	__gong__computeReferencePass2(stage.A_measure_1s, stage.A_measure_1s_reference, stage)

	__gong__computeReferencePass2(stage.A_parts, stage.A_parts_reference, stage)

	__gong__computeReferencePass2(stage.A_part_1s, stage.A_part_1s_reference, stage)

	__gong__computeReferencePass2(stage.Accidentals, stage.Accidentals_reference, stage)

	__gong__computeReferencePass2(stage.Accidental_marks, stage.Accidental_marks_reference, stage)

	__gong__computeReferencePass2(stage.Accidental_texts, stage.Accidental_texts_reference, stage)

	__gong__computeReferencePass2(stage.Accords, stage.Accords_reference, stage)

	__gong__computeReferencePass2(stage.Accordion_registrations, stage.Accordion_registrations_reference, stage)

	__gong__computeReferencePass2(stage.Appearances, stage.Appearances_reference, stage)

	__gong__computeReferencePass2(stage.Arpeggiates, stage.Arpeggiates_reference, stage)

	__gong__computeReferencePass2(stage.Arrows, stage.Arrows_reference, stage)

	__gong__computeReferencePass2(stage.Articulationss, stage.Articulationss_reference, stage)

	__gong__computeReferencePass2(stage.Assesss, stage.Assesss_reference, stage)

	__gong__computeReferencePass2(stage.Attributess, stage.Attributess_reference, stage)

	__gong__computeReferencePass2(stage.Backups, stage.Backups_reference, stage)

	__gong__computeReferencePass2(stage.Bar_style_colors, stage.Bar_style_colors_reference, stage)

	__gong__computeReferencePass2(stage.Barlines, stage.Barlines_reference, stage)

	__gong__computeReferencePass2(stage.Barres, stage.Barres_reference, stage)

	__gong__computeReferencePass2(stage.Basss, stage.Basss_reference, stage)

	__gong__computeReferencePass2(stage.Bass_steps, stage.Bass_steps_reference, stage)

	__gong__computeReferencePass2(stage.Beams, stage.Beams_reference, stage)

	__gong__computeReferencePass2(stage.Beat_repeats, stage.Beat_repeats_reference, stage)

	__gong__computeReferencePass2(stage.Beat_unit_tieds, stage.Beat_unit_tieds_reference, stage)

	__gong__computeReferencePass2(stage.Beaters, stage.Beaters_reference, stage)

	__gong__computeReferencePass2(stage.Bends, stage.Bends_reference, stage)

	__gong__computeReferencePass2(stage.Bookmarks, stage.Bookmarks_reference, stage)

	__gong__computeReferencePass2(stage.Brackets, stage.Brackets_reference, stage)

	__gong__computeReferencePass2(stage.Breath_marks, stage.Breath_marks_reference, stage)

	__gong__computeReferencePass2(stage.Caesuras, stage.Caesuras_reference, stage)

	__gong__computeReferencePass2(stage.Cancels, stage.Cancels_reference, stage)

	__gong__computeReferencePass2(stage.Clefs, stage.Clefs_reference, stage)

	__gong__computeReferencePass2(stage.Codas, stage.Codas_reference, stage)

	__gong__computeReferencePass2(stage.Credits, stage.Credits_reference, stage)

	__gong__computeReferencePass2(stage.Dashess, stage.Dashess_reference, stage)

	__gong__computeReferencePass2(stage.Defaultss, stage.Defaultss_reference, stage)

	__gong__computeReferencePass2(stage.Degrees, stage.Degrees_reference, stage)

	__gong__computeReferencePass2(stage.Degree_alters, stage.Degree_alters_reference, stage)

	__gong__computeReferencePass2(stage.Degree_types, stage.Degree_types_reference, stage)

	__gong__computeReferencePass2(stage.Degree_values, stage.Degree_values_reference, stage)

	__gong__computeReferencePass2(stage.Directions, stage.Directions_reference, stage)

	__gong__computeReferencePass2(stage.Direction_types, stage.Direction_types_reference, stage)

	__gong__computeReferencePass2(stage.Distances, stage.Distances_reference, stage)

	__gong__computeReferencePass2(stage.Doubles, stage.Doubles_reference, stage)

	__gong__computeReferencePass2(stage.Dynamicss, stage.Dynamicss_reference, stage)

	__gong__computeReferencePass2(stage.Effects, stage.Effects_reference, stage)

	__gong__computeReferencePass2(stage.Elisions, stage.Elisions_reference, stage)

	__gong__computeReferencePass2(stage.Emptys, stage.Emptys_reference, stage)

	__gong__computeReferencePass2(stage.Empty_fonts, stage.Empty_fonts_reference, stage)

	__gong__computeReferencePass2(stage.Empty_lines, stage.Empty_lines_reference, stage)

	__gong__computeReferencePass2(stage.Empty_placements, stage.Empty_placements_reference, stage)

	__gong__computeReferencePass2(stage.Empty_placement_smufls, stage.Empty_placement_smufls_reference, stage)

	__gong__computeReferencePass2(stage.Empty_print_object_style_aligns, stage.Empty_print_object_style_aligns_reference, stage)

	__gong__computeReferencePass2(stage.Empty_print_styles, stage.Empty_print_styles_reference, stage)

	__gong__computeReferencePass2(stage.Empty_print_style_aligns, stage.Empty_print_style_aligns_reference, stage)

	__gong__computeReferencePass2(stage.Empty_print_style_align_ids, stage.Empty_print_style_align_ids_reference, stage)

	__gong__computeReferencePass2(stage.Empty_trill_sounds, stage.Empty_trill_sounds_reference, stage)

	__gong__computeReferencePass2(stage.Encodings, stage.Encodings_reference, stage)

	__gong__computeReferencePass2(stage.Endings, stage.Endings_reference, stage)

	__gong__computeReferencePass2(stage.Extends, stage.Extends_reference, stage)

	__gong__computeReferencePass2(stage.Features, stage.Features_reference, stage)

	__gong__computeReferencePass2(stage.Fermatas, stage.Fermatas_reference, stage)

	__gong__computeReferencePass2(stage.Figures, stage.Figures_reference, stage)

	__gong__computeReferencePass2(stage.Figured_basss, stage.Figured_basss_reference, stage)

	__gong__computeReferencePass2(stage.Fingerings, stage.Fingerings_reference, stage)

	__gong__computeReferencePass2(stage.First_frets, stage.First_frets_reference, stage)

	__gong__computeReferencePass2(stage.For_parts, stage.For_parts_reference, stage)

	__gong__computeReferencePass2(stage.Formatted_symbols, stage.Formatted_symbols_reference, stage)

	__gong__computeReferencePass2(stage.Formatted_symbol_ids, stage.Formatted_symbol_ids_reference, stage)

	__gong__computeReferencePass2(stage.Formatted_texts, stage.Formatted_texts_reference, stage)

	__gong__computeReferencePass2(stage.Formatted_text_ids, stage.Formatted_text_ids_reference, stage)

	__gong__computeReferencePass2(stage.Forwards, stage.Forwards_reference, stage)

	__gong__computeReferencePass2(stage.Frames, stage.Frames_reference, stage)

	__gong__computeReferencePass2(stage.Frame_notes, stage.Frame_notes_reference, stage)

	__gong__computeReferencePass2(stage.Frets, stage.Frets_reference, stage)

	__gong__computeReferencePass2(stage.Glasss, stage.Glasss_reference, stage)

	__gong__computeReferencePass2(stage.Glissandos, stage.Glissandos_reference, stage)

	__gong__computeReferencePass2(stage.Glyphs, stage.Glyphs_reference, stage)

	__gong__computeReferencePass2(stage.Graces, stage.Graces_reference, stage)

	__gong__computeReferencePass2(stage.Group_barlines, stage.Group_barlines_reference, stage)

	__gong__computeReferencePass2(stage.Group_names, stage.Group_names_reference, stage)

	__gong__computeReferencePass2(stage.Group_symbols, stage.Group_symbols_reference, stage)

	__gong__computeReferencePass2(stage.Groupings, stage.Groupings_reference, stage)

	__gong__computeReferencePass2(stage.Hammer_on_pull_offs, stage.Hammer_on_pull_offs_reference, stage)

	__gong__computeReferencePass2(stage.Handbells, stage.Handbells_reference, stage)

	__gong__computeReferencePass2(stage.Harmon_closeds, stage.Harmon_closeds_reference, stage)

	__gong__computeReferencePass2(stage.Harmon_mutes, stage.Harmon_mutes_reference, stage)

	__gong__computeReferencePass2(stage.Harmonics, stage.Harmonics_reference, stage)

	__gong__computeReferencePass2(stage.Harmonys, stage.Harmonys_reference, stage)

	__gong__computeReferencePass2(stage.Harmony_alters, stage.Harmony_alters_reference, stage)

	__gong__computeReferencePass2(stage.Harp_pedalss, stage.Harp_pedalss_reference, stage)

	__gong__computeReferencePass2(stage.Heel_toes, stage.Heel_toes_reference, stage)

	__gong__computeReferencePass2(stage.Holes, stage.Holes_reference, stage)

	__gong__computeReferencePass2(stage.Hole_closeds, stage.Hole_closeds_reference, stage)

	__gong__computeReferencePass2(stage.Horizontal_turns, stage.Horizontal_turns_reference, stage)

	__gong__computeReferencePass2(stage.Identifications, stage.Identifications_reference, stage)

	__gong__computeReferencePass2(stage.Images, stage.Images_reference, stage)

	__gong__computeReferencePass2(stage.Instruments, stage.Instruments_reference, stage)

	__gong__computeReferencePass2(stage.Instrument_changes, stage.Instrument_changes_reference, stage)

	__gong__computeReferencePass2(stage.Instrument_links, stage.Instrument_links_reference, stage)

	__gong__computeReferencePass2(stage.Interchangeables, stage.Interchangeables_reference, stage)

	__gong__computeReferencePass2(stage.Inversions, stage.Inversions_reference, stage)

	__gong__computeReferencePass2(stage.Keys, stage.Keys_reference, stage)

	__gong__computeReferencePass2(stage.Key_accidentals, stage.Key_accidentals_reference, stage)

	__gong__computeReferencePass2(stage.Key_octaves, stage.Key_octaves_reference, stage)

	__gong__computeReferencePass2(stage.Kinds, stage.Kinds_reference, stage)

	__gong__computeReferencePass2(stage.Levels, stage.Levels_reference, stage)

	__gong__computeReferencePass2(stage.Line_details, stage.Line_details_reference, stage)

	__gong__computeReferencePass2(stage.Line_widths, stage.Line_widths_reference, stage)

	__gong__computeReferencePass2(stage.Links, stage.Links_reference, stage)

	__gong__computeReferencePass2(stage.Listens, stage.Listens_reference, stage)

	__gong__computeReferencePass2(stage.Listenings, stage.Listenings_reference, stage)

	__gong__computeReferencePass2(stage.Lyrics, stage.Lyrics_reference, stage)

	__gong__computeReferencePass2(stage.Lyric_fonts, stage.Lyric_fonts_reference, stage)

	__gong__computeReferencePass2(stage.Lyric_languages, stage.Lyric_languages_reference, stage)

	__gong__computeReferencePass2(stage.Measure_layouts, stage.Measure_layouts_reference, stage)

	__gong__computeReferencePass2(stage.Measure_numberings, stage.Measure_numberings_reference, stage)

	__gong__computeReferencePass2(stage.Measure_repeats, stage.Measure_repeats_reference, stage)

	__gong__computeReferencePass2(stage.Measure_styles, stage.Measure_styles_reference, stage)

	__gong__computeReferencePass2(stage.Membranes, stage.Membranes_reference, stage)

	__gong__computeReferencePass2(stage.Metals, stage.Metals_reference, stage)

	__gong__computeReferencePass2(stage.Metronomes, stage.Metronomes_reference, stage)

	__gong__computeReferencePass2(stage.Metronome_beams, stage.Metronome_beams_reference, stage)

	__gong__computeReferencePass2(stage.Metronome_notes, stage.Metronome_notes_reference, stage)

	__gong__computeReferencePass2(stage.Metronome_tieds, stage.Metronome_tieds_reference, stage)

	__gong__computeReferencePass2(stage.Metronome_tuplets, stage.Metronome_tuplets_reference, stage)

	__gong__computeReferencePass2(stage.Midi_devices, stage.Midi_devices_reference, stage)

	__gong__computeReferencePass2(stage.Midi_instruments, stage.Midi_instruments_reference, stage)

	__gong__computeReferencePass2(stage.Miscellaneouss, stage.Miscellaneouss_reference, stage)

	__gong__computeReferencePass2(stage.Miscellaneous_fields, stage.Miscellaneous_fields_reference, stage)

	__gong__computeReferencePass2(stage.Mordents, stage.Mordents_reference, stage)

	__gong__computeReferencePass2(stage.Multiple_rests, stage.Multiple_rests_reference, stage)

	__gong__computeReferencePass2(stage.Name_displays, stage.Name_displays_reference, stage)

	__gong__computeReferencePass2(stage.Non_arpeggiates, stage.Non_arpeggiates_reference, stage)

	__gong__computeReferencePass2(stage.Notationss, stage.Notationss_reference, stage)

	__gong__computeReferencePass2(stage.Notes, stage.Notes_reference, stage)

	__gong__computeReferencePass2(stage.Note_sizes, stage.Note_sizes_reference, stage)

	__gong__computeReferencePass2(stage.Note_types, stage.Note_types_reference, stage)

	__gong__computeReferencePass2(stage.Noteheads, stage.Noteheads_reference, stage)

	__gong__computeReferencePass2(stage.Notehead_texts, stage.Notehead_texts_reference, stage)

	__gong__computeReferencePass2(stage.Numerals, stage.Numerals_reference, stage)

	__gong__computeReferencePass2(stage.Numeral_keys, stage.Numeral_keys_reference, stage)

	__gong__computeReferencePass2(stage.Numeral_roots, stage.Numeral_roots_reference, stage)

	__gong__computeReferencePass2(stage.Octave_shifts, stage.Octave_shifts_reference, stage)

	__gong__computeReferencePass2(stage.Offsets, stage.Offsets_reference, stage)

	__gong__computeReferencePass2(stage.Opuss, stage.Opuss_reference, stage)

	__gong__computeReferencePass2(stage.Ornamentss, stage.Ornamentss_reference, stage)

	__gong__computeReferencePass2(stage.Other_appearances, stage.Other_appearances_reference, stage)

	__gong__computeReferencePass2(stage.Other_directions, stage.Other_directions_reference, stage)

	__gong__computeReferencePass2(stage.Other_listenings, stage.Other_listenings_reference, stage)

	__gong__computeReferencePass2(stage.Other_notations, stage.Other_notations_reference, stage)

	__gong__computeReferencePass2(stage.Other_placement_texts, stage.Other_placement_texts_reference, stage)

	__gong__computeReferencePass2(stage.Other_plays, stage.Other_plays_reference, stage)

	__gong__computeReferencePass2(stage.Other_texts, stage.Other_texts_reference, stage)

	__gong__computeReferencePass2(stage.Page_layouts, stage.Page_layouts_reference, stage)

	__gong__computeReferencePass2(stage.Page_marginss, stage.Page_marginss_reference, stage)

	__gong__computeReferencePass2(stage.Part_clefs, stage.Part_clefs_reference, stage)

	__gong__computeReferencePass2(stage.Part_groups, stage.Part_groups_reference, stage)

	__gong__computeReferencePass2(stage.Part_links, stage.Part_links_reference, stage)

	__gong__computeReferencePass2(stage.Part_lists, stage.Part_lists_reference, stage)

	__gong__computeReferencePass2(stage.Part_names, stage.Part_names_reference, stage)

	__gong__computeReferencePass2(stage.Part_symbols, stage.Part_symbols_reference, stage)

	__gong__computeReferencePass2(stage.Part_transposes, stage.Part_transposes_reference, stage)

	__gong__computeReferencePass2(stage.Pedals, stage.Pedals_reference, stage)

	__gong__computeReferencePass2(stage.Pedal_tunings, stage.Pedal_tunings_reference, stage)

	__gong__computeReferencePass2(stage.Per_minutes, stage.Per_minutes_reference, stage)

	__gong__computeReferencePass2(stage.Percussions, stage.Percussions_reference, stage)

	__gong__computeReferencePass2(stage.Pitchs, stage.Pitchs_reference, stage)

	__gong__computeReferencePass2(stage.Pitcheds, stage.Pitcheds_reference, stage)

	__gong__computeReferencePass2(stage.Placement_texts, stage.Placement_texts_reference, stage)

	__gong__computeReferencePass2(stage.Plays, stage.Plays_reference, stage)

	__gong__computeReferencePass2(stage.Players, stage.Players_reference, stage)

	__gong__computeReferencePass2(stage.Principal_voices, stage.Principal_voices_reference, stage)

	__gong__computeReferencePass2(stage.Prints, stage.Prints_reference, stage)

	__gong__computeReferencePass2(stage.Releases, stage.Releases_reference, stage)

	__gong__computeReferencePass2(stage.Repeats, stage.Repeats_reference, stage)

	__gong__computeReferencePass2(stage.Rests, stage.Rests_reference, stage)

	__gong__computeReferencePass2(stage.Roots, stage.Roots_reference, stage)

	__gong__computeReferencePass2(stage.Root_steps, stage.Root_steps_reference, stage)

	__gong__computeReferencePass2(stage.Scalings, stage.Scalings_reference, stage)

	__gong__computeReferencePass2(stage.Scordaturas, stage.Scordaturas_reference, stage)

	__gong__computeReferencePass2(stage.Score_instruments, stage.Score_instruments_reference, stage)

	__gong__computeReferencePass2(stage.Score_parts, stage.Score_parts_reference, stage)

	__gong__computeReferencePass2(stage.Score_partwises, stage.Score_partwises_reference, stage)

	__gong__computeReferencePass2(stage.Score_timewises, stage.Score_timewises_reference, stage)

	__gong__computeReferencePass2(stage.Segnos, stage.Segnos_reference, stage)

	__gong__computeReferencePass2(stage.Slashs, stage.Slashs_reference, stage)

	__gong__computeReferencePass2(stage.Slides, stage.Slides_reference, stage)

	__gong__computeReferencePass2(stage.Slurs, stage.Slurs_reference, stage)

	__gong__computeReferencePass2(stage.Sounds, stage.Sounds_reference, stage)

	__gong__computeReferencePass2(stage.Staff_detailss, stage.Staff_detailss_reference, stage)

	__gong__computeReferencePass2(stage.Staff_divides, stage.Staff_divides_reference, stage)

	__gong__computeReferencePass2(stage.Staff_layouts, stage.Staff_layouts_reference, stage)

	__gong__computeReferencePass2(stage.Staff_sizes, stage.Staff_sizes_reference, stage)

	__gong__computeReferencePass2(stage.Staff_tunings, stage.Staff_tunings_reference, stage)

	__gong__computeReferencePass2(stage.Stems, stage.Stems_reference, stage)

	__gong__computeReferencePass2(stage.Sticks, stage.Sticks_reference, stage)

	__gong__computeReferencePass2(stage.String_mutes, stage.String_mutes_reference, stage)

	__gong__computeReferencePass2(stage.String_types, stage.String_types_reference, stage)

	__gong__computeReferencePass2(stage.Strong_accents, stage.Strong_accents_reference, stage)

	__gong__computeReferencePass2(stage.Style_texts, stage.Style_texts_reference, stage)

	__gong__computeReferencePass2(stage.Supportss, stage.Supportss_reference, stage)

	__gong__computeReferencePass2(stage.Swings, stage.Swings_reference, stage)

	__gong__computeReferencePass2(stage.Syncs, stage.Syncs_reference, stage)

	__gong__computeReferencePass2(stage.System_dividerss, stage.System_dividerss_reference, stage)

	__gong__computeReferencePass2(stage.System_layouts, stage.System_layouts_reference, stage)

	__gong__computeReferencePass2(stage.System_marginss, stage.System_marginss_reference, stage)

	__gong__computeReferencePass2(stage.Taps, stage.Taps_reference, stage)

	__gong__computeReferencePass2(stage.Technicals, stage.Technicals_reference, stage)

	__gong__computeReferencePass2(stage.Text_element_datas, stage.Text_element_datas_reference, stage)

	__gong__computeReferencePass2(stage.Ties, stage.Ties_reference, stage)

	__gong__computeReferencePass2(stage.Tieds, stage.Tieds_reference, stage)

	__gong__computeReferencePass2(stage.Times, stage.Times_reference, stage)

	__gong__computeReferencePass2(stage.Time_modifications, stage.Time_modifications_reference, stage)

	__gong__computeReferencePass2(stage.Timpanis, stage.Timpanis_reference, stage)

	__gong__computeReferencePass2(stage.Transposes, stage.Transposes_reference, stage)

	__gong__computeReferencePass2(stage.Tremolos, stage.Tremolos_reference, stage)

	__gong__computeReferencePass2(stage.Tuplets, stage.Tuplets_reference, stage)

	__gong__computeReferencePass2(stage.Tuplet_dots, stage.Tuplet_dots_reference, stage)

	__gong__computeReferencePass2(stage.Tuplet_numbers, stage.Tuplet_numbers_reference, stage)

	__gong__computeReferencePass2(stage.Tuplet_portions, stage.Tuplet_portions_reference, stage)

	__gong__computeReferencePass2(stage.Tuplet_types, stage.Tuplet_types_reference, stage)

	__gong__computeReferencePass2(stage.Typed_texts, stage.Typed_texts_reference, stage)

	__gong__computeReferencePass2(stage.Unpitcheds, stage.Unpitcheds_reference, stage)

	__gong__computeReferencePass2(stage.Virtual_instruments, stage.Virtual_instruments_reference, stage)

	__gong__computeReferencePass2(stage.Waits, stage.Waits_reference, stage)

	__gong__computeReferencePass2(stage.Wavy_lines, stage.Wavy_lines_reference, stage)

	__gong__computeReferencePass2(stage.Wedges, stage.Wedges_reference, stage)

	__gong__computeReferencePass2(stage.Woods, stage.Woods_reference, stage)

	__gong__computeReferencePass2(stage.Works, stage.Works_reference, stage)

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
	return __gong__getOrder(stage.A_directive_stagedOrder, stage.A_directives_referenceOrder, a_directive, "A_directive")
}

func (a_measure *A_measure) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_measure_stagedOrder, stage.A_measures_referenceOrder, a_measure, "A_measure")
}

func (a_measure_1 *A_measure_1) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_measure_1_stagedOrder, stage.A_measure_1s_referenceOrder, a_measure_1, "A_measure_1")
}

func (a_part *A_part) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_part_stagedOrder, stage.A_parts_referenceOrder, a_part, "A_part")
}

func (a_part_1 *A_part_1) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.A_part_1_stagedOrder, stage.A_part_1s_referenceOrder, a_part_1, "A_part_1")
}

func (accidental *Accidental) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Accidental_stagedOrder, stage.Accidentals_referenceOrder, accidental, "Accidental")
}

func (accidental_mark *Accidental_mark) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Accidental_mark_stagedOrder, stage.Accidental_marks_referenceOrder, accidental_mark, "Accidental_mark")
}

func (accidental_text *Accidental_text) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Accidental_text_stagedOrder, stage.Accidental_texts_referenceOrder, accidental_text, "Accidental_text")
}

func (accord *Accord) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Accord_stagedOrder, stage.Accords_referenceOrder, accord, "Accord")
}

func (accordion_registration *Accordion_registration) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Accordion_registration_stagedOrder, stage.Accordion_registrations_referenceOrder, accordion_registration, "Accordion_registration")
}

func (appearance *Appearance) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Appearance_stagedOrder, stage.Appearances_referenceOrder, appearance, "Appearance")
}

func (arpeggiate *Arpeggiate) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Arpeggiate_stagedOrder, stage.Arpeggiates_referenceOrder, arpeggiate, "Arpeggiate")
}

func (arrow *Arrow) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Arrow_stagedOrder, stage.Arrows_referenceOrder, arrow, "Arrow")
}

func (articulations *Articulations) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Articulations_stagedOrder, stage.Articulationss_referenceOrder, articulations, "Articulations")
}

func (assess *Assess) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Assess_stagedOrder, stage.Assesss_referenceOrder, assess, "Assess")
}

func (attributes *Attributes) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Attributes_stagedOrder, stage.Attributess_referenceOrder, attributes, "Attributes")
}

func (backup *Backup) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Backup_stagedOrder, stage.Backups_referenceOrder, backup, "Backup")
}

func (bar_style_color *Bar_style_color) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Bar_style_color_stagedOrder, stage.Bar_style_colors_referenceOrder, bar_style_color, "Bar_style_color")
}

func (barline *Barline) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Barline_stagedOrder, stage.Barlines_referenceOrder, barline, "Barline")
}

func (barre *Barre) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Barre_stagedOrder, stage.Barres_referenceOrder, barre, "Barre")
}

func (bass *Bass) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Bass_stagedOrder, stage.Basss_referenceOrder, bass, "Bass")
}

func (bass_step *Bass_step) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Bass_step_stagedOrder, stage.Bass_steps_referenceOrder, bass_step, "Bass_step")
}

func (beam *Beam) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Beam_stagedOrder, stage.Beams_referenceOrder, beam, "Beam")
}

func (beat_repeat *Beat_repeat) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Beat_repeat_stagedOrder, stage.Beat_repeats_referenceOrder, beat_repeat, "Beat_repeat")
}

func (beat_unit_tied *Beat_unit_tied) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Beat_unit_tied_stagedOrder, stage.Beat_unit_tieds_referenceOrder, beat_unit_tied, "Beat_unit_tied")
}

func (beater *Beater) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Beater_stagedOrder, stage.Beaters_referenceOrder, beater, "Beater")
}

func (bend *Bend) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Bend_stagedOrder, stage.Bends_referenceOrder, bend, "Bend")
}

func (bookmark *Bookmark) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Bookmark_stagedOrder, stage.Bookmarks_referenceOrder, bookmark, "Bookmark")
}

func (bracket *Bracket) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Bracket_stagedOrder, stage.Brackets_referenceOrder, bracket, "Bracket")
}

func (breath_mark *Breath_mark) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Breath_mark_stagedOrder, stage.Breath_marks_referenceOrder, breath_mark, "Breath_mark")
}

func (caesura *Caesura) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Caesura_stagedOrder, stage.Caesuras_referenceOrder, caesura, "Caesura")
}

func (cancel *Cancel) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Cancel_stagedOrder, stage.Cancels_referenceOrder, cancel, "Cancel")
}

func (clef *Clef) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Clef_stagedOrder, stage.Clefs_referenceOrder, clef, "Clef")
}

func (coda *Coda) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Coda_stagedOrder, stage.Codas_referenceOrder, coda, "Coda")
}

func (credit *Credit) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Credit_stagedOrder, stage.Credits_referenceOrder, credit, "Credit")
}

func (dashes *Dashes) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Dashes_stagedOrder, stage.Dashess_referenceOrder, dashes, "Dashes")
}

func (defaults *Defaults) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Defaults_stagedOrder, stage.Defaultss_referenceOrder, defaults, "Defaults")
}

func (degree *Degree) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Degree_stagedOrder, stage.Degrees_referenceOrder, degree, "Degree")
}

func (degree_alter *Degree_alter) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Degree_alter_stagedOrder, stage.Degree_alters_referenceOrder, degree_alter, "Degree_alter")
}

func (degree_type *Degree_type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Degree_type_stagedOrder, stage.Degree_types_referenceOrder, degree_type, "Degree_type")
}

func (degree_value *Degree_value) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Degree_value_stagedOrder, stage.Degree_values_referenceOrder, degree_value, "Degree_value")
}

func (direction *Direction) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Direction_stagedOrder, stage.Directions_referenceOrder, direction, "Direction")
}

func (direction_type *Direction_type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Direction_type_stagedOrder, stage.Direction_types_referenceOrder, direction_type, "Direction_type")
}

func (distance *Distance) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Distance_stagedOrder, stage.Distances_referenceOrder, distance, "Distance")
}

func (double *Double) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Double_stagedOrder, stage.Doubles_referenceOrder, double, "Double")
}

func (dynamics *Dynamics) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Dynamics_stagedOrder, stage.Dynamicss_referenceOrder, dynamics, "Dynamics")
}

func (effect *Effect) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Effect_stagedOrder, stage.Effects_referenceOrder, effect, "Effect")
}

func (elision *Elision) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Elision_stagedOrder, stage.Elisions_referenceOrder, elision, "Elision")
}

func (empty *Empty) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Empty_stagedOrder, stage.Emptys_referenceOrder, empty, "Empty")
}

func (empty_font *Empty_font) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Empty_font_stagedOrder, stage.Empty_fonts_referenceOrder, empty_font, "Empty_font")
}

func (empty_line *Empty_line) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Empty_line_stagedOrder, stage.Empty_lines_referenceOrder, empty_line, "Empty_line")
}

func (empty_placement *Empty_placement) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Empty_placement_stagedOrder, stage.Empty_placements_referenceOrder, empty_placement, "Empty_placement")
}

func (empty_placement_smufl *Empty_placement_smufl) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Empty_placement_smufl_stagedOrder, stage.Empty_placement_smufls_referenceOrder, empty_placement_smufl, "Empty_placement_smufl")
}

func (empty_print_object_style_align *Empty_print_object_style_align) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Empty_print_object_style_align_stagedOrder, stage.Empty_print_object_style_aligns_referenceOrder, empty_print_object_style_align, "Empty_print_object_style_align")
}

func (empty_print_style *Empty_print_style) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Empty_print_style_stagedOrder, stage.Empty_print_styles_referenceOrder, empty_print_style, "Empty_print_style")
}

func (empty_print_style_align *Empty_print_style_align) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Empty_print_style_align_stagedOrder, stage.Empty_print_style_aligns_referenceOrder, empty_print_style_align, "Empty_print_style_align")
}

func (empty_print_style_align_id *Empty_print_style_align_id) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Empty_print_style_align_id_stagedOrder, stage.Empty_print_style_align_ids_referenceOrder, empty_print_style_align_id, "Empty_print_style_align_id")
}

func (empty_trill_sound *Empty_trill_sound) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Empty_trill_sound_stagedOrder, stage.Empty_trill_sounds_referenceOrder, empty_trill_sound, "Empty_trill_sound")
}

func (encoding *Encoding) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Encoding_stagedOrder, stage.Encodings_referenceOrder, encoding, "Encoding")
}

func (ending *Ending) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Ending_stagedOrder, stage.Endings_referenceOrder, ending, "Ending")
}

func (extend *Extend) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Extend_stagedOrder, stage.Extends_referenceOrder, extend, "Extend")
}

func (feature *Feature) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Feature_stagedOrder, stage.Features_referenceOrder, feature, "Feature")
}

func (fermata *Fermata) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Fermata_stagedOrder, stage.Fermatas_referenceOrder, fermata, "Fermata")
}

func (figure *Figure) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Figure_stagedOrder, stage.Figures_referenceOrder, figure, "Figure")
}

func (figured_bass *Figured_bass) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Figured_bass_stagedOrder, stage.Figured_basss_referenceOrder, figured_bass, "Figured_bass")
}

func (fingering *Fingering) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Fingering_stagedOrder, stage.Fingerings_referenceOrder, fingering, "Fingering")
}

func (first_fret *First_fret) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.First_fret_stagedOrder, stage.First_frets_referenceOrder, first_fret, "First_fret")
}

func (for_part *For_part) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.For_part_stagedOrder, stage.For_parts_referenceOrder, for_part, "For_part")
}

func (formatted_symbol *Formatted_symbol) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Formatted_symbol_stagedOrder, stage.Formatted_symbols_referenceOrder, formatted_symbol, "Formatted_symbol")
}

func (formatted_symbol_id *Formatted_symbol_id) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Formatted_symbol_id_stagedOrder, stage.Formatted_symbol_ids_referenceOrder, formatted_symbol_id, "Formatted_symbol_id")
}

func (formatted_text *Formatted_text) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Formatted_text_stagedOrder, stage.Formatted_texts_referenceOrder, formatted_text, "Formatted_text")
}

func (formatted_text_id *Formatted_text_id) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Formatted_text_id_stagedOrder, stage.Formatted_text_ids_referenceOrder, formatted_text_id, "Formatted_text_id")
}

func (forward *Forward) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Forward_stagedOrder, stage.Forwards_referenceOrder, forward, "Forward")
}

func (frame *Frame) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Frame_stagedOrder, stage.Frames_referenceOrder, frame, "Frame")
}

func (frame_note *Frame_note) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Frame_note_stagedOrder, stage.Frame_notes_referenceOrder, frame_note, "Frame_note")
}

func (fret *Fret) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Fret_stagedOrder, stage.Frets_referenceOrder, fret, "Fret")
}

func (glass *Glass) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Glass_stagedOrder, stage.Glasss_referenceOrder, glass, "Glass")
}

func (glissando *Glissando) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Glissando_stagedOrder, stage.Glissandos_referenceOrder, glissando, "Glissando")
}

func (glyph *Glyph) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Glyph_stagedOrder, stage.Glyphs_referenceOrder, glyph, "Glyph")
}

func (grace *Grace) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Grace_stagedOrder, stage.Graces_referenceOrder, grace, "Grace")
}

func (group_barline *Group_barline) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Group_barline_stagedOrder, stage.Group_barlines_referenceOrder, group_barline, "Group_barline")
}

func (group_name *Group_name) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Group_name_stagedOrder, stage.Group_names_referenceOrder, group_name, "Group_name")
}

func (group_symbol *Group_symbol) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Group_symbol_stagedOrder, stage.Group_symbols_referenceOrder, group_symbol, "Group_symbol")
}

func (grouping *Grouping) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Grouping_stagedOrder, stage.Groupings_referenceOrder, grouping, "Grouping")
}

func (hammer_on_pull_off *Hammer_on_pull_off) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Hammer_on_pull_off_stagedOrder, stage.Hammer_on_pull_offs_referenceOrder, hammer_on_pull_off, "Hammer_on_pull_off")
}

func (handbell *Handbell) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Handbell_stagedOrder, stage.Handbells_referenceOrder, handbell, "Handbell")
}

func (harmon_closed *Harmon_closed) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Harmon_closed_stagedOrder, stage.Harmon_closeds_referenceOrder, harmon_closed, "Harmon_closed")
}

func (harmon_mute *Harmon_mute) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Harmon_mute_stagedOrder, stage.Harmon_mutes_referenceOrder, harmon_mute, "Harmon_mute")
}

func (harmonic *Harmonic) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Harmonic_stagedOrder, stage.Harmonics_referenceOrder, harmonic, "Harmonic")
}

func (harmony *Harmony) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Harmony_stagedOrder, stage.Harmonys_referenceOrder, harmony, "Harmony")
}

func (harmony_alter *Harmony_alter) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Harmony_alter_stagedOrder, stage.Harmony_alters_referenceOrder, harmony_alter, "Harmony_alter")
}

func (harp_pedals *Harp_pedals) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Harp_pedals_stagedOrder, stage.Harp_pedalss_referenceOrder, harp_pedals, "Harp_pedals")
}

func (heel_toe *Heel_toe) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Heel_toe_stagedOrder, stage.Heel_toes_referenceOrder, heel_toe, "Heel_toe")
}

func (hole *Hole) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Hole_stagedOrder, stage.Holes_referenceOrder, hole, "Hole")
}

func (hole_closed *Hole_closed) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Hole_closed_stagedOrder, stage.Hole_closeds_referenceOrder, hole_closed, "Hole_closed")
}

func (horizontal_turn *Horizontal_turn) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Horizontal_turn_stagedOrder, stage.Horizontal_turns_referenceOrder, horizontal_turn, "Horizontal_turn")
}

func (identification *Identification) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Identification_stagedOrder, stage.Identifications_referenceOrder, identification, "Identification")
}

func (image *Image) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Image_stagedOrder, stage.Images_referenceOrder, image, "Image")
}

func (instrument *Instrument) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Instrument_stagedOrder, stage.Instruments_referenceOrder, instrument, "Instrument")
}

func (instrument_change *Instrument_change) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Instrument_change_stagedOrder, stage.Instrument_changes_referenceOrder, instrument_change, "Instrument_change")
}

func (instrument_link *Instrument_link) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Instrument_link_stagedOrder, stage.Instrument_links_referenceOrder, instrument_link, "Instrument_link")
}

func (interchangeable *Interchangeable) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Interchangeable_stagedOrder, stage.Interchangeables_referenceOrder, interchangeable, "Interchangeable")
}

func (inversion *Inversion) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Inversion_stagedOrder, stage.Inversions_referenceOrder, inversion, "Inversion")
}

func (key *Key) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Key_stagedOrder, stage.Keys_referenceOrder, key, "Key")
}

func (key_accidental *Key_accidental) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Key_accidental_stagedOrder, stage.Key_accidentals_referenceOrder, key_accidental, "Key_accidental")
}

func (key_octave *Key_octave) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Key_octave_stagedOrder, stage.Key_octaves_referenceOrder, key_octave, "Key_octave")
}

func (kind *Kind) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Kind_stagedOrder, stage.Kinds_referenceOrder, kind, "Kind")
}

func (level *Level) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Level_stagedOrder, stage.Levels_referenceOrder, level, "Level")
}

func (line_detail *Line_detail) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Line_detail_stagedOrder, stage.Line_details_referenceOrder, line_detail, "Line_detail")
}

func (line_width *Line_width) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Line_width_stagedOrder, stage.Line_widths_referenceOrder, line_width, "Line_width")
}

func (link *Link) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Link_stagedOrder, stage.Links_referenceOrder, link, "Link")
}

func (listen *Listen) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Listen_stagedOrder, stage.Listens_referenceOrder, listen, "Listen")
}

func (listening *Listening) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Listening_stagedOrder, stage.Listenings_referenceOrder, listening, "Listening")
}

func (lyric *Lyric) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Lyric_stagedOrder, stage.Lyrics_referenceOrder, lyric, "Lyric")
}

func (lyric_font *Lyric_font) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Lyric_font_stagedOrder, stage.Lyric_fonts_referenceOrder, lyric_font, "Lyric_font")
}

func (lyric_language *Lyric_language) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Lyric_language_stagedOrder, stage.Lyric_languages_referenceOrder, lyric_language, "Lyric_language")
}

func (measure_layout *Measure_layout) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Measure_layout_stagedOrder, stage.Measure_layouts_referenceOrder, measure_layout, "Measure_layout")
}

func (measure_numbering *Measure_numbering) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Measure_numbering_stagedOrder, stage.Measure_numberings_referenceOrder, measure_numbering, "Measure_numbering")
}

func (measure_repeat *Measure_repeat) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Measure_repeat_stagedOrder, stage.Measure_repeats_referenceOrder, measure_repeat, "Measure_repeat")
}

func (measure_style *Measure_style) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Measure_style_stagedOrder, stage.Measure_styles_referenceOrder, measure_style, "Measure_style")
}

func (membrane *Membrane) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Membrane_stagedOrder, stage.Membranes_referenceOrder, membrane, "Membrane")
}

func (metal *Metal) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Metal_stagedOrder, stage.Metals_referenceOrder, metal, "Metal")
}

func (metronome *Metronome) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Metronome_stagedOrder, stage.Metronomes_referenceOrder, metronome, "Metronome")
}

func (metronome_beam *Metronome_beam) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Metronome_beam_stagedOrder, stage.Metronome_beams_referenceOrder, metronome_beam, "Metronome_beam")
}

func (metronome_note *Metronome_note) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Metronome_note_stagedOrder, stage.Metronome_notes_referenceOrder, metronome_note, "Metronome_note")
}

func (metronome_tied *Metronome_tied) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Metronome_tied_stagedOrder, stage.Metronome_tieds_referenceOrder, metronome_tied, "Metronome_tied")
}

func (metronome_tuplet *Metronome_tuplet) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Metronome_tuplet_stagedOrder, stage.Metronome_tuplets_referenceOrder, metronome_tuplet, "Metronome_tuplet")
}

func (midi_device *Midi_device) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Midi_device_stagedOrder, stage.Midi_devices_referenceOrder, midi_device, "Midi_device")
}

func (midi_instrument *Midi_instrument) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Midi_instrument_stagedOrder, stage.Midi_instruments_referenceOrder, midi_instrument, "Midi_instrument")
}

func (miscellaneous *Miscellaneous) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Miscellaneous_stagedOrder, stage.Miscellaneouss_referenceOrder, miscellaneous, "Miscellaneous")
}

func (miscellaneous_field *Miscellaneous_field) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Miscellaneous_field_stagedOrder, stage.Miscellaneous_fields_referenceOrder, miscellaneous_field, "Miscellaneous_field")
}

func (mordent *Mordent) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Mordent_stagedOrder, stage.Mordents_referenceOrder, mordent, "Mordent")
}

func (multiple_rest *Multiple_rest) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Multiple_rest_stagedOrder, stage.Multiple_rests_referenceOrder, multiple_rest, "Multiple_rest")
}

func (name_display *Name_display) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Name_display_stagedOrder, stage.Name_displays_referenceOrder, name_display, "Name_display")
}

func (non_arpeggiate *Non_arpeggiate) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Non_arpeggiate_stagedOrder, stage.Non_arpeggiates_referenceOrder, non_arpeggiate, "Non_arpeggiate")
}

func (notations *Notations) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Notations_stagedOrder, stage.Notationss_referenceOrder, notations, "Notations")
}

func (note *Note) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Note_stagedOrder, stage.Notes_referenceOrder, note, "Note")
}

func (note_size *Note_size) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Note_size_stagedOrder, stage.Note_sizes_referenceOrder, note_size, "Note_size")
}

func (note_type *Note_type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Note_type_stagedOrder, stage.Note_types_referenceOrder, note_type, "Note_type")
}

func (notehead *Notehead) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Notehead_stagedOrder, stage.Noteheads_referenceOrder, notehead, "Notehead")
}

func (notehead_text *Notehead_text) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Notehead_text_stagedOrder, stage.Notehead_texts_referenceOrder, notehead_text, "Notehead_text")
}

func (numeral *Numeral) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Numeral_stagedOrder, stage.Numerals_referenceOrder, numeral, "Numeral")
}

func (numeral_key *Numeral_key) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Numeral_key_stagedOrder, stage.Numeral_keys_referenceOrder, numeral_key, "Numeral_key")
}

func (numeral_root *Numeral_root) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Numeral_root_stagedOrder, stage.Numeral_roots_referenceOrder, numeral_root, "Numeral_root")
}

func (octave_shift *Octave_shift) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Octave_shift_stagedOrder, stage.Octave_shifts_referenceOrder, octave_shift, "Octave_shift")
}

func (offset *Offset) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Offset_stagedOrder, stage.Offsets_referenceOrder, offset, "Offset")
}

func (opus *Opus) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Opus_stagedOrder, stage.Opuss_referenceOrder, opus, "Opus")
}

func (ornaments *Ornaments) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Ornaments_stagedOrder, stage.Ornamentss_referenceOrder, ornaments, "Ornaments")
}

func (other_appearance *Other_appearance) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Other_appearance_stagedOrder, stage.Other_appearances_referenceOrder, other_appearance, "Other_appearance")
}

func (other_direction *Other_direction) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Other_direction_stagedOrder, stage.Other_directions_referenceOrder, other_direction, "Other_direction")
}

func (other_listening *Other_listening) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Other_listening_stagedOrder, stage.Other_listenings_referenceOrder, other_listening, "Other_listening")
}

func (other_notation *Other_notation) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Other_notation_stagedOrder, stage.Other_notations_referenceOrder, other_notation, "Other_notation")
}

func (other_placement_text *Other_placement_text) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Other_placement_text_stagedOrder, stage.Other_placement_texts_referenceOrder, other_placement_text, "Other_placement_text")
}

func (other_play *Other_play) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Other_play_stagedOrder, stage.Other_plays_referenceOrder, other_play, "Other_play")
}

func (other_text *Other_text) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Other_text_stagedOrder, stage.Other_texts_referenceOrder, other_text, "Other_text")
}

func (page_layout *Page_layout) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Page_layout_stagedOrder, stage.Page_layouts_referenceOrder, page_layout, "Page_layout")
}

func (page_margins *Page_margins) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Page_margins_stagedOrder, stage.Page_marginss_referenceOrder, page_margins, "Page_margins")
}

func (part_clef *Part_clef) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Part_clef_stagedOrder, stage.Part_clefs_referenceOrder, part_clef, "Part_clef")
}

func (part_group *Part_group) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Part_group_stagedOrder, stage.Part_groups_referenceOrder, part_group, "Part_group")
}

func (part_link *Part_link) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Part_link_stagedOrder, stage.Part_links_referenceOrder, part_link, "Part_link")
}

func (part_list *Part_list) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Part_list_stagedOrder, stage.Part_lists_referenceOrder, part_list, "Part_list")
}

func (part_name *Part_name) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Part_name_stagedOrder, stage.Part_names_referenceOrder, part_name, "Part_name")
}

func (part_symbol *Part_symbol) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Part_symbol_stagedOrder, stage.Part_symbols_referenceOrder, part_symbol, "Part_symbol")
}

func (part_transpose *Part_transpose) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Part_transpose_stagedOrder, stage.Part_transposes_referenceOrder, part_transpose, "Part_transpose")
}

func (pedal *Pedal) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Pedal_stagedOrder, stage.Pedals_referenceOrder, pedal, "Pedal")
}

func (pedal_tuning *Pedal_tuning) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Pedal_tuning_stagedOrder, stage.Pedal_tunings_referenceOrder, pedal_tuning, "Pedal_tuning")
}

func (per_minute *Per_minute) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Per_minute_stagedOrder, stage.Per_minutes_referenceOrder, per_minute, "Per_minute")
}

func (percussion *Percussion) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Percussion_stagedOrder, stage.Percussions_referenceOrder, percussion, "Percussion")
}

func (pitch *Pitch) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Pitch_stagedOrder, stage.Pitchs_referenceOrder, pitch, "Pitch")
}

func (pitched *Pitched) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Pitched_stagedOrder, stage.Pitcheds_referenceOrder, pitched, "Pitched")
}

func (placement_text *Placement_text) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Placement_text_stagedOrder, stage.Placement_texts_referenceOrder, placement_text, "Placement_text")
}

func (play *Play) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Play_stagedOrder, stage.Plays_referenceOrder, play, "Play")
}

func (player *Player) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Player_stagedOrder, stage.Players_referenceOrder, player, "Player")
}

func (principal_voice *Principal_voice) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Principal_voice_stagedOrder, stage.Principal_voices_referenceOrder, principal_voice, "Principal_voice")
}

func (print *Print) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Print_stagedOrder, stage.Prints_referenceOrder, print, "Print")
}

func (release *Release) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Release_stagedOrder, stage.Releases_referenceOrder, release, "Release")
}

func (repeat *Repeat) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Repeat_stagedOrder, stage.Repeats_referenceOrder, repeat, "Repeat")
}

func (rest *Rest) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Rest_stagedOrder, stage.Rests_referenceOrder, rest, "Rest")
}

func (root *Root) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Root_stagedOrder, stage.Roots_referenceOrder, root, "Root")
}

func (root_step *Root_step) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Root_step_stagedOrder, stage.Root_steps_referenceOrder, root_step, "Root_step")
}

func (scaling *Scaling) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Scaling_stagedOrder, stage.Scalings_referenceOrder, scaling, "Scaling")
}

func (scordatura *Scordatura) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Scordatura_stagedOrder, stage.Scordaturas_referenceOrder, scordatura, "Scordatura")
}

func (score_instrument *Score_instrument) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Score_instrument_stagedOrder, stage.Score_instruments_referenceOrder, score_instrument, "Score_instrument")
}

func (score_part *Score_part) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Score_part_stagedOrder, stage.Score_parts_referenceOrder, score_part, "Score_part")
}

func (score_partwise *Score_partwise) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Score_partwise_stagedOrder, stage.Score_partwises_referenceOrder, score_partwise, "Score_partwise")
}

func (score_timewise *Score_timewise) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Score_timewise_stagedOrder, stage.Score_timewises_referenceOrder, score_timewise, "Score_timewise")
}

func (segno *Segno) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Segno_stagedOrder, stage.Segnos_referenceOrder, segno, "Segno")
}

func (slash *Slash) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Slash_stagedOrder, stage.Slashs_referenceOrder, slash, "Slash")
}

func (slide *Slide) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Slide_stagedOrder, stage.Slides_referenceOrder, slide, "Slide")
}

func (slur *Slur) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Slur_stagedOrder, stage.Slurs_referenceOrder, slur, "Slur")
}

func (sound *Sound) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Sound_stagedOrder, stage.Sounds_referenceOrder, sound, "Sound")
}

func (staff_details *Staff_details) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Staff_details_stagedOrder, stage.Staff_detailss_referenceOrder, staff_details, "Staff_details")
}

func (staff_divide *Staff_divide) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Staff_divide_stagedOrder, stage.Staff_divides_referenceOrder, staff_divide, "Staff_divide")
}

func (staff_layout *Staff_layout) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Staff_layout_stagedOrder, stage.Staff_layouts_referenceOrder, staff_layout, "Staff_layout")
}

func (staff_size *Staff_size) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Staff_size_stagedOrder, stage.Staff_sizes_referenceOrder, staff_size, "Staff_size")
}

func (staff_tuning *Staff_tuning) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Staff_tuning_stagedOrder, stage.Staff_tunings_referenceOrder, staff_tuning, "Staff_tuning")
}

func (stem *Stem) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Stem_stagedOrder, stage.Stems_referenceOrder, stem, "Stem")
}

func (stick *Stick) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Stick_stagedOrder, stage.Sticks_referenceOrder, stick, "Stick")
}

func (string_mute *String_mute) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.String_mute_stagedOrder, stage.String_mutes_referenceOrder, string_mute, "String_mute")
}

func (string_type *String_type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.String_type_stagedOrder, stage.String_types_referenceOrder, string_type, "String_type")
}

func (strong_accent *Strong_accent) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Strong_accent_stagedOrder, stage.Strong_accents_referenceOrder, strong_accent, "Strong_accent")
}

func (style_text *Style_text) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Style_text_stagedOrder, stage.Style_texts_referenceOrder, style_text, "Style_text")
}

func (supports *Supports) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Supports_stagedOrder, stage.Supportss_referenceOrder, supports, "Supports")
}

func (swing *Swing) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Swing_stagedOrder, stage.Swings_referenceOrder, swing, "Swing")
}

func (sync *Sync) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Sync_stagedOrder, stage.Syncs_referenceOrder, sync, "Sync")
}

func (system_dividers *System_dividers) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.System_dividers_stagedOrder, stage.System_dividerss_referenceOrder, system_dividers, "System_dividers")
}

func (system_layout *System_layout) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.System_layout_stagedOrder, stage.System_layouts_referenceOrder, system_layout, "System_layout")
}

func (system_margins *System_margins) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.System_margins_stagedOrder, stage.System_marginss_referenceOrder, system_margins, "System_margins")
}

func (tap *Tap) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tap_stagedOrder, stage.Taps_referenceOrder, tap, "Tap")
}

func (technical *Technical) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Technical_stagedOrder, stage.Technicals_referenceOrder, technical, "Technical")
}

func (text_element_data *Text_element_data) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Text_element_data_stagedOrder, stage.Text_element_datas_referenceOrder, text_element_data, "Text_element_data")
}

func (tie *Tie) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tie_stagedOrder, stage.Ties_referenceOrder, tie, "Tie")
}

func (tied *Tied) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tied_stagedOrder, stage.Tieds_referenceOrder, tied, "Tied")
}

func (time *Time) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Time_stagedOrder, stage.Times_referenceOrder, time, "Time")
}

func (time_modification *Time_modification) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Time_modification_stagedOrder, stage.Time_modifications_referenceOrder, time_modification, "Time_modification")
}

func (timpani *Timpani) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Timpani_stagedOrder, stage.Timpanis_referenceOrder, timpani, "Timpani")
}

func (transpose *Transpose) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Transpose_stagedOrder, stage.Transposes_referenceOrder, transpose, "Transpose")
}

func (tremolo *Tremolo) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tremolo_stagedOrder, stage.Tremolos_referenceOrder, tremolo, "Tremolo")
}

func (tuplet *Tuplet) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tuplet_stagedOrder, stage.Tuplets_referenceOrder, tuplet, "Tuplet")
}

func (tuplet_dot *Tuplet_dot) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tuplet_dot_stagedOrder, stage.Tuplet_dots_referenceOrder, tuplet_dot, "Tuplet_dot")
}

func (tuplet_number *Tuplet_number) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tuplet_number_stagedOrder, stage.Tuplet_numbers_referenceOrder, tuplet_number, "Tuplet_number")
}

func (tuplet_portion *Tuplet_portion) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tuplet_portion_stagedOrder, stage.Tuplet_portions_referenceOrder, tuplet_portion, "Tuplet_portion")
}

func (tuplet_type *Tuplet_type) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Tuplet_type_stagedOrder, stage.Tuplet_types_referenceOrder, tuplet_type, "Tuplet_type")
}

func (typed_text *Typed_text) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Typed_text_stagedOrder, stage.Typed_texts_referenceOrder, typed_text, "Typed_text")
}

func (unpitched *Unpitched) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Unpitched_stagedOrder, stage.Unpitcheds_referenceOrder, unpitched, "Unpitched")
}

func (virtual_instrument *Virtual_instrument) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Virtual_instrument_stagedOrder, stage.Virtual_instruments_referenceOrder, virtual_instrument, "Virtual_instrument")
}

func (wait *Wait) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Wait_stagedOrder, stage.Waits_referenceOrder, wait, "Wait")
}

func (wavy_line *Wavy_line) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Wavy_line_stagedOrder, stage.Wavy_lines_referenceOrder, wavy_line, "Wavy_line")
}

func (wedge *Wedge) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Wedge_stagedOrder, stage.Wedges_referenceOrder, wedge, "Wedge")
}

func (wood *Wood) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Wood_stagedOrder, stage.Woods_referenceOrder, wood, "Wood")
}

func (work *Work) GongGetOrder(stage *Stage) uint {
	return __gong__getOrder(stage.Work_stagedOrder, stage.Works_referenceOrder, work, "Work")
}

// GongGetIdentifier returns a unique identifier of the instance in the staging area
// This identifier is composed of the Gongstruct name and the order of the instance
// in the staging area
// It is used to identify instances across sessions
// insertion point per named struct
func (a_directive *A_directive) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_directive, a_directive.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_directive *A_directive) GongGetReferenceIdentifier(stage *Stage) string {
	return a_directive.GongGetIdentifier(stage)
}

func (a_measure *A_measure) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_measure, a_measure.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_measure *A_measure) GongGetReferenceIdentifier(stage *Stage) string {
	return a_measure.GongGetIdentifier(stage)
}

func (a_measure_1 *A_measure_1) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_measure_1, a_measure_1.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_measure_1 *A_measure_1) GongGetReferenceIdentifier(stage *Stage) string {
	return a_measure_1.GongGetIdentifier(stage)
}

func (a_part *A_part) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_part, a_part.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_part *A_part) GongGetReferenceIdentifier(stage *Stage) string {
	return a_part.GongGetIdentifier(stage)
}

func (a_part_1 *A_part_1) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(a_part_1, a_part_1.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (a_part_1 *A_part_1) GongGetReferenceIdentifier(stage *Stage) string {
	return a_part_1.GongGetIdentifier(stage)
}

func (accidental *Accidental) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(accidental, accidental.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (accidental *Accidental) GongGetReferenceIdentifier(stage *Stage) string {
	return accidental.GongGetIdentifier(stage)
}

func (accidental_mark *Accidental_mark) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(accidental_mark, accidental_mark.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (accidental_mark *Accidental_mark) GongGetReferenceIdentifier(stage *Stage) string {
	return accidental_mark.GongGetIdentifier(stage)
}

func (accidental_text *Accidental_text) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(accidental_text, accidental_text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (accidental_text *Accidental_text) GongGetReferenceIdentifier(stage *Stage) string {
	return accidental_text.GongGetIdentifier(stage)
}

func (accord *Accord) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(accord, accord.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (accord *Accord) GongGetReferenceIdentifier(stage *Stage) string {
	return accord.GongGetIdentifier(stage)
}

func (accordion_registration *Accordion_registration) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(accordion_registration, accordion_registration.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (accordion_registration *Accordion_registration) GongGetReferenceIdentifier(stage *Stage) string {
	return accordion_registration.GongGetIdentifier(stage)
}

func (appearance *Appearance) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(appearance, appearance.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (appearance *Appearance) GongGetReferenceIdentifier(stage *Stage) string {
	return appearance.GongGetIdentifier(stage)
}

func (arpeggiate *Arpeggiate) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(arpeggiate, arpeggiate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arpeggiate *Arpeggiate) GongGetReferenceIdentifier(stage *Stage) string {
	return arpeggiate.GongGetIdentifier(stage)
}

func (arrow *Arrow) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(arrow, arrow.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (arrow *Arrow) GongGetReferenceIdentifier(stage *Stage) string {
	return arrow.GongGetIdentifier(stage)
}

func (articulations *Articulations) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(articulations, articulations.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (articulations *Articulations) GongGetReferenceIdentifier(stage *Stage) string {
	return articulations.GongGetIdentifier(stage)
}

func (assess *Assess) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(assess, assess.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (assess *Assess) GongGetReferenceIdentifier(stage *Stage) string {
	return assess.GongGetIdentifier(stage)
}

func (attributes *Attributes) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(attributes, attributes.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (attributes *Attributes) GongGetReferenceIdentifier(stage *Stage) string {
	return attributes.GongGetIdentifier(stage)
}

func (backup *Backup) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(backup, backup.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (backup *Backup) GongGetReferenceIdentifier(stage *Stage) string {
	return backup.GongGetIdentifier(stage)
}

func (bar_style_color *Bar_style_color) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(bar_style_color, bar_style_color.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bar_style_color *Bar_style_color) GongGetReferenceIdentifier(stage *Stage) string {
	return bar_style_color.GongGetIdentifier(stage)
}

func (barline *Barline) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(barline, barline.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (barline *Barline) GongGetReferenceIdentifier(stage *Stage) string {
	return barline.GongGetIdentifier(stage)
}

func (barre *Barre) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(barre, barre.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (barre *Barre) GongGetReferenceIdentifier(stage *Stage) string {
	return barre.GongGetIdentifier(stage)
}

func (bass *Bass) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(bass, bass.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bass *Bass) GongGetReferenceIdentifier(stage *Stage) string {
	return bass.GongGetIdentifier(stage)
}

func (bass_step *Bass_step) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(bass_step, bass_step.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bass_step *Bass_step) GongGetReferenceIdentifier(stage *Stage) string {
	return bass_step.GongGetIdentifier(stage)
}

func (beam *Beam) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(beam, beam.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (beam *Beam) GongGetReferenceIdentifier(stage *Stage) string {
	return beam.GongGetIdentifier(stage)
}

func (beat_repeat *Beat_repeat) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(beat_repeat, beat_repeat.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (beat_repeat *Beat_repeat) GongGetReferenceIdentifier(stage *Stage) string {
	return beat_repeat.GongGetIdentifier(stage)
}

func (beat_unit_tied *Beat_unit_tied) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(beat_unit_tied, beat_unit_tied.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (beat_unit_tied *Beat_unit_tied) GongGetReferenceIdentifier(stage *Stage) string {
	return beat_unit_tied.GongGetIdentifier(stage)
}

func (beater *Beater) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(beater, beater.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (beater *Beater) GongGetReferenceIdentifier(stage *Stage) string {
	return beater.GongGetIdentifier(stage)
}

func (bend *Bend) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(bend, bend.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bend *Bend) GongGetReferenceIdentifier(stage *Stage) string {
	return bend.GongGetIdentifier(stage)
}

func (bookmark *Bookmark) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(bookmark, bookmark.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bookmark *Bookmark) GongGetReferenceIdentifier(stage *Stage) string {
	return bookmark.GongGetIdentifier(stage)
}

func (bracket *Bracket) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(bracket, bracket.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (bracket *Bracket) GongGetReferenceIdentifier(stage *Stage) string {
	return bracket.GongGetIdentifier(stage)
}

func (breath_mark *Breath_mark) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(breath_mark, breath_mark.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (breath_mark *Breath_mark) GongGetReferenceIdentifier(stage *Stage) string {
	return breath_mark.GongGetIdentifier(stage)
}

func (caesura *Caesura) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(caesura, caesura.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (caesura *Caesura) GongGetReferenceIdentifier(stage *Stage) string {
	return caesura.GongGetIdentifier(stage)
}

func (cancel *Cancel) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(cancel, cancel.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (cancel *Cancel) GongGetReferenceIdentifier(stage *Stage) string {
	return cancel.GongGetIdentifier(stage)
}

func (clef *Clef) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(clef, clef.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (clef *Clef) GongGetReferenceIdentifier(stage *Stage) string {
	return clef.GongGetIdentifier(stage)
}

func (coda *Coda) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(coda, coda.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (coda *Coda) GongGetReferenceIdentifier(stage *Stage) string {
	return coda.GongGetIdentifier(stage)
}

func (credit *Credit) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(credit, credit.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (credit *Credit) GongGetReferenceIdentifier(stage *Stage) string {
	return credit.GongGetIdentifier(stage)
}

func (dashes *Dashes) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(dashes, dashes.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dashes *Dashes) GongGetReferenceIdentifier(stage *Stage) string {
	return dashes.GongGetIdentifier(stage)
}

func (defaults *Defaults) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(defaults, defaults.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (defaults *Defaults) GongGetReferenceIdentifier(stage *Stage) string {
	return defaults.GongGetIdentifier(stage)
}

func (degree *Degree) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(degree, degree.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (degree *Degree) GongGetReferenceIdentifier(stage *Stage) string {
	return degree.GongGetIdentifier(stage)
}

func (degree_alter *Degree_alter) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(degree_alter, degree_alter.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (degree_alter *Degree_alter) GongGetReferenceIdentifier(stage *Stage) string {
	return degree_alter.GongGetIdentifier(stage)
}

func (degree_type *Degree_type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(degree_type, degree_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (degree_type *Degree_type) GongGetReferenceIdentifier(stage *Stage) string {
	return degree_type.GongGetIdentifier(stage)
}

func (degree_value *Degree_value) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(degree_value, degree_value.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (degree_value *Degree_value) GongGetReferenceIdentifier(stage *Stage) string {
	return degree_value.GongGetIdentifier(stage)
}

func (direction *Direction) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(direction, direction.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (direction *Direction) GongGetReferenceIdentifier(stage *Stage) string {
	return direction.GongGetIdentifier(stage)
}

func (direction_type *Direction_type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(direction_type, direction_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (direction_type *Direction_type) GongGetReferenceIdentifier(stage *Stage) string {
	return direction_type.GongGetIdentifier(stage)
}

func (distance *Distance) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(distance, distance.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (distance *Distance) GongGetReferenceIdentifier(stage *Stage) string {
	return distance.GongGetIdentifier(stage)
}

func (double *Double) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(double, double.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (double *Double) GongGetReferenceIdentifier(stage *Stage) string {
	return double.GongGetIdentifier(stage)
}

func (dynamics *Dynamics) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(dynamics, dynamics.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (dynamics *Dynamics) GongGetReferenceIdentifier(stage *Stage) string {
	return dynamics.GongGetIdentifier(stage)
}

func (effect *Effect) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(effect, effect.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (effect *Effect) GongGetReferenceIdentifier(stage *Stage) string {
	return effect.GongGetIdentifier(stage)
}

func (elision *Elision) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(elision, elision.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (elision *Elision) GongGetReferenceIdentifier(stage *Stage) string {
	return elision.GongGetIdentifier(stage)
}

func (empty *Empty) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(empty, empty.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty *Empty) GongGetReferenceIdentifier(stage *Stage) string {
	return empty.GongGetIdentifier(stage)
}

func (empty_font *Empty_font) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(empty_font, empty_font.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_font *Empty_font) GongGetReferenceIdentifier(stage *Stage) string {
	return empty_font.GongGetIdentifier(stage)
}

func (empty_line *Empty_line) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(empty_line, empty_line.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_line *Empty_line) GongGetReferenceIdentifier(stage *Stage) string {
	return empty_line.GongGetIdentifier(stage)
}

func (empty_placement *Empty_placement) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(empty_placement, empty_placement.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_placement *Empty_placement) GongGetReferenceIdentifier(stage *Stage) string {
	return empty_placement.GongGetIdentifier(stage)
}

func (empty_placement_smufl *Empty_placement_smufl) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(empty_placement_smufl, empty_placement_smufl.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_placement_smufl *Empty_placement_smufl) GongGetReferenceIdentifier(stage *Stage) string {
	return empty_placement_smufl.GongGetIdentifier(stage)
}

func (empty_print_object_style_align *Empty_print_object_style_align) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(empty_print_object_style_align, empty_print_object_style_align.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_print_object_style_align *Empty_print_object_style_align) GongGetReferenceIdentifier(stage *Stage) string {
	return empty_print_object_style_align.GongGetIdentifier(stage)
}

func (empty_print_style *Empty_print_style) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(empty_print_style, empty_print_style.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_print_style *Empty_print_style) GongGetReferenceIdentifier(stage *Stage) string {
	return empty_print_style.GongGetIdentifier(stage)
}

func (empty_print_style_align *Empty_print_style_align) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(empty_print_style_align, empty_print_style_align.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_print_style_align *Empty_print_style_align) GongGetReferenceIdentifier(stage *Stage) string {
	return empty_print_style_align.GongGetIdentifier(stage)
}

func (empty_print_style_align_id *Empty_print_style_align_id) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(empty_print_style_align_id, empty_print_style_align_id.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_print_style_align_id *Empty_print_style_align_id) GongGetReferenceIdentifier(stage *Stage) string {
	return empty_print_style_align_id.GongGetIdentifier(stage)
}

func (empty_trill_sound *Empty_trill_sound) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(empty_trill_sound, empty_trill_sound.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (empty_trill_sound *Empty_trill_sound) GongGetReferenceIdentifier(stage *Stage) string {
	return empty_trill_sound.GongGetIdentifier(stage)
}

func (encoding *Encoding) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(encoding, encoding.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (encoding *Encoding) GongGetReferenceIdentifier(stage *Stage) string {
	return encoding.GongGetIdentifier(stage)
}

func (ending *Ending) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ending, ending.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ending *Ending) GongGetReferenceIdentifier(stage *Stage) string {
	return ending.GongGetIdentifier(stage)
}

func (extend *Extend) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(extend, extend.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (extend *Extend) GongGetReferenceIdentifier(stage *Stage) string {
	return extend.GongGetIdentifier(stage)
}

func (feature *Feature) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(feature, feature.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (feature *Feature) GongGetReferenceIdentifier(stage *Stage) string {
	return feature.GongGetIdentifier(stage)
}

func (fermata *Fermata) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(fermata, fermata.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (fermata *Fermata) GongGetReferenceIdentifier(stage *Stage) string {
	return fermata.GongGetIdentifier(stage)
}

func (figure *Figure) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(figure, figure.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (figure *Figure) GongGetReferenceIdentifier(stage *Stage) string {
	return figure.GongGetIdentifier(stage)
}

func (figured_bass *Figured_bass) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(figured_bass, figured_bass.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (figured_bass *Figured_bass) GongGetReferenceIdentifier(stage *Stage) string {
	return figured_bass.GongGetIdentifier(stage)
}

func (fingering *Fingering) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(fingering, fingering.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (fingering *Fingering) GongGetReferenceIdentifier(stage *Stage) string {
	return fingering.GongGetIdentifier(stage)
}

func (first_fret *First_fret) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(first_fret, first_fret.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (first_fret *First_fret) GongGetReferenceIdentifier(stage *Stage) string {
	return first_fret.GongGetIdentifier(stage)
}

func (for_part *For_part) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(for_part, for_part.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (for_part *For_part) GongGetReferenceIdentifier(stage *Stage) string {
	return for_part.GongGetIdentifier(stage)
}

func (formatted_symbol *Formatted_symbol) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formatted_symbol, formatted_symbol.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formatted_symbol *Formatted_symbol) GongGetReferenceIdentifier(stage *Stage) string {
	return formatted_symbol.GongGetIdentifier(stage)
}

func (formatted_symbol_id *Formatted_symbol_id) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formatted_symbol_id, formatted_symbol_id.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formatted_symbol_id *Formatted_symbol_id) GongGetReferenceIdentifier(stage *Stage) string {
	return formatted_symbol_id.GongGetIdentifier(stage)
}

func (formatted_text *Formatted_text) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formatted_text, formatted_text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formatted_text *Formatted_text) GongGetReferenceIdentifier(stage *Stage) string {
	return formatted_text.GongGetIdentifier(stage)
}

func (formatted_text_id *Formatted_text_id) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(formatted_text_id, formatted_text_id.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (formatted_text_id *Formatted_text_id) GongGetReferenceIdentifier(stage *Stage) string {
	return formatted_text_id.GongGetIdentifier(stage)
}

func (forward *Forward) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(forward, forward.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (forward *Forward) GongGetReferenceIdentifier(stage *Stage) string {
	return forward.GongGetIdentifier(stage)
}

func (frame *Frame) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(frame, frame.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (frame *Frame) GongGetReferenceIdentifier(stage *Stage) string {
	return frame.GongGetIdentifier(stage)
}

func (frame_note *Frame_note) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(frame_note, frame_note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (frame_note *Frame_note) GongGetReferenceIdentifier(stage *Stage) string {
	return frame_note.GongGetIdentifier(stage)
}

func (fret *Fret) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(fret, fret.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (fret *Fret) GongGetReferenceIdentifier(stage *Stage) string {
	return fret.GongGetIdentifier(stage)
}

func (glass *Glass) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(glass, glass.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (glass *Glass) GongGetReferenceIdentifier(stage *Stage) string {
	return glass.GongGetIdentifier(stage)
}

func (glissando *Glissando) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(glissando, glissando.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (glissando *Glissando) GongGetReferenceIdentifier(stage *Stage) string {
	return glissando.GongGetIdentifier(stage)
}

func (glyph *Glyph) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(glyph, glyph.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (glyph *Glyph) GongGetReferenceIdentifier(stage *Stage) string {
	return glyph.GongGetIdentifier(stage)
}

func (grace *Grace) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(grace, grace.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (grace *Grace) GongGetReferenceIdentifier(stage *Stage) string {
	return grace.GongGetIdentifier(stage)
}

func (group_barline *Group_barline) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(group_barline, group_barline.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group_barline *Group_barline) GongGetReferenceIdentifier(stage *Stage) string {
	return group_barline.GongGetIdentifier(stage)
}

func (group_name *Group_name) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(group_name, group_name.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group_name *Group_name) GongGetReferenceIdentifier(stage *Stage) string {
	return group_name.GongGetIdentifier(stage)
}

func (group_symbol *Group_symbol) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(group_symbol, group_symbol.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (group_symbol *Group_symbol) GongGetReferenceIdentifier(stage *Stage) string {
	return group_symbol.GongGetIdentifier(stage)
}

func (grouping *Grouping) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(grouping, grouping.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (grouping *Grouping) GongGetReferenceIdentifier(stage *Stage) string {
	return grouping.GongGetIdentifier(stage)
}

func (hammer_on_pull_off *Hammer_on_pull_off) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(hammer_on_pull_off, hammer_on_pull_off.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (hammer_on_pull_off *Hammer_on_pull_off) GongGetReferenceIdentifier(stage *Stage) string {
	return hammer_on_pull_off.GongGetIdentifier(stage)
}

func (handbell *Handbell) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(handbell, handbell.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (handbell *Handbell) GongGetReferenceIdentifier(stage *Stage) string {
	return handbell.GongGetIdentifier(stage)
}

func (harmon_closed *Harmon_closed) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(harmon_closed, harmon_closed.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harmon_closed *Harmon_closed) GongGetReferenceIdentifier(stage *Stage) string {
	return harmon_closed.GongGetIdentifier(stage)
}

func (harmon_mute *Harmon_mute) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(harmon_mute, harmon_mute.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harmon_mute *Harmon_mute) GongGetReferenceIdentifier(stage *Stage) string {
	return harmon_mute.GongGetIdentifier(stage)
}

func (harmonic *Harmonic) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(harmonic, harmonic.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harmonic *Harmonic) GongGetReferenceIdentifier(stage *Stage) string {
	return harmonic.GongGetIdentifier(stage)
}

func (harmony *Harmony) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(harmony, harmony.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harmony *Harmony) GongGetReferenceIdentifier(stage *Stage) string {
	return harmony.GongGetIdentifier(stage)
}

func (harmony_alter *Harmony_alter) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(harmony_alter, harmony_alter.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harmony_alter *Harmony_alter) GongGetReferenceIdentifier(stage *Stage) string {
	return harmony_alter.GongGetIdentifier(stage)
}

func (harp_pedals *Harp_pedals) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(harp_pedals, harp_pedals.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (harp_pedals *Harp_pedals) GongGetReferenceIdentifier(stage *Stage) string {
	return harp_pedals.GongGetIdentifier(stage)
}

func (heel_toe *Heel_toe) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(heel_toe, heel_toe.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (heel_toe *Heel_toe) GongGetReferenceIdentifier(stage *Stage) string {
	return heel_toe.GongGetIdentifier(stage)
}

func (hole *Hole) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(hole, hole.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (hole *Hole) GongGetReferenceIdentifier(stage *Stage) string {
	return hole.GongGetIdentifier(stage)
}

func (hole_closed *Hole_closed) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(hole_closed, hole_closed.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (hole_closed *Hole_closed) GongGetReferenceIdentifier(stage *Stage) string {
	return hole_closed.GongGetIdentifier(stage)
}

func (horizontal_turn *Horizontal_turn) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(horizontal_turn, horizontal_turn.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (horizontal_turn *Horizontal_turn) GongGetReferenceIdentifier(stage *Stage) string {
	return horizontal_turn.GongGetIdentifier(stage)
}

func (identification *Identification) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(identification, identification.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (identification *Identification) GongGetReferenceIdentifier(stage *Stage) string {
	return identification.GongGetIdentifier(stage)
}

func (image *Image) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(image, image.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (image *Image) GongGetReferenceIdentifier(stage *Stage) string {
	return image.GongGetIdentifier(stage)
}

func (instrument *Instrument) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(instrument, instrument.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (instrument *Instrument) GongGetReferenceIdentifier(stage *Stage) string {
	return instrument.GongGetIdentifier(stage)
}

func (instrument_change *Instrument_change) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(instrument_change, instrument_change.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (instrument_change *Instrument_change) GongGetReferenceIdentifier(stage *Stage) string {
	return instrument_change.GongGetIdentifier(stage)
}

func (instrument_link *Instrument_link) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(instrument_link, instrument_link.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (instrument_link *Instrument_link) GongGetReferenceIdentifier(stage *Stage) string {
	return instrument_link.GongGetIdentifier(stage)
}

func (interchangeable *Interchangeable) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(interchangeable, interchangeable.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (interchangeable *Interchangeable) GongGetReferenceIdentifier(stage *Stage) string {
	return interchangeable.GongGetIdentifier(stage)
}

func (inversion *Inversion) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(inversion, inversion.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (inversion *Inversion) GongGetReferenceIdentifier(stage *Stage) string {
	return inversion.GongGetIdentifier(stage)
}

func (key *Key) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(key, key.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (key *Key) GongGetReferenceIdentifier(stage *Stage) string {
	return key.GongGetIdentifier(stage)
}

func (key_accidental *Key_accidental) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(key_accidental, key_accidental.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (key_accidental *Key_accidental) GongGetReferenceIdentifier(stage *Stage) string {
	return key_accidental.GongGetIdentifier(stage)
}

func (key_octave *Key_octave) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(key_octave, key_octave.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (key_octave *Key_octave) GongGetReferenceIdentifier(stage *Stage) string {
	return key_octave.GongGetIdentifier(stage)
}

func (kind *Kind) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(kind, kind.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (kind *Kind) GongGetReferenceIdentifier(stage *Stage) string {
	return kind.GongGetIdentifier(stage)
}

func (level *Level) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(level, level.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (level *Level) GongGetReferenceIdentifier(stage *Stage) string {
	return level.GongGetIdentifier(stage)
}

func (line_detail *Line_detail) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(line_detail, line_detail.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (line_detail *Line_detail) GongGetReferenceIdentifier(stage *Stage) string {
	return line_detail.GongGetIdentifier(stage)
}

func (line_width *Line_width) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(line_width, line_width.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (line_width *Line_width) GongGetReferenceIdentifier(stage *Stage) string {
	return line_width.GongGetIdentifier(stage)
}

func (link *Link) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(link, link.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (link *Link) GongGetReferenceIdentifier(stage *Stage) string {
	return link.GongGetIdentifier(stage)
}

func (listen *Listen) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(listen, listen.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (listen *Listen) GongGetReferenceIdentifier(stage *Stage) string {
	return listen.GongGetIdentifier(stage)
}

func (listening *Listening) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(listening, listening.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (listening *Listening) GongGetReferenceIdentifier(stage *Stage) string {
	return listening.GongGetIdentifier(stage)
}

func (lyric *Lyric) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(lyric, lyric.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lyric *Lyric) GongGetReferenceIdentifier(stage *Stage) string {
	return lyric.GongGetIdentifier(stage)
}

func (lyric_font *Lyric_font) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(lyric_font, lyric_font.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lyric_font *Lyric_font) GongGetReferenceIdentifier(stage *Stage) string {
	return lyric_font.GongGetIdentifier(stage)
}

func (lyric_language *Lyric_language) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(lyric_language, lyric_language.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (lyric_language *Lyric_language) GongGetReferenceIdentifier(stage *Stage) string {
	return lyric_language.GongGetIdentifier(stage)
}

func (measure_layout *Measure_layout) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(measure_layout, measure_layout.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measure_layout *Measure_layout) GongGetReferenceIdentifier(stage *Stage) string {
	return measure_layout.GongGetIdentifier(stage)
}

func (measure_numbering *Measure_numbering) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(measure_numbering, measure_numbering.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measure_numbering *Measure_numbering) GongGetReferenceIdentifier(stage *Stage) string {
	return measure_numbering.GongGetIdentifier(stage)
}

func (measure_repeat *Measure_repeat) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(measure_repeat, measure_repeat.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measure_repeat *Measure_repeat) GongGetReferenceIdentifier(stage *Stage) string {
	return measure_repeat.GongGetIdentifier(stage)
}

func (measure_style *Measure_style) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(measure_style, measure_style.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (measure_style *Measure_style) GongGetReferenceIdentifier(stage *Stage) string {
	return measure_style.GongGetIdentifier(stage)
}

func (membrane *Membrane) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(membrane, membrane.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (membrane *Membrane) GongGetReferenceIdentifier(stage *Stage) string {
	return membrane.GongGetIdentifier(stage)
}

func (metal *Metal) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(metal, metal.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metal *Metal) GongGetReferenceIdentifier(stage *Stage) string {
	return metal.GongGetIdentifier(stage)
}

func (metronome *Metronome) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(metronome, metronome.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metronome *Metronome) GongGetReferenceIdentifier(stage *Stage) string {
	return metronome.GongGetIdentifier(stage)
}

func (metronome_beam *Metronome_beam) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(metronome_beam, metronome_beam.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metronome_beam *Metronome_beam) GongGetReferenceIdentifier(stage *Stage) string {
	return metronome_beam.GongGetIdentifier(stage)
}

func (metronome_note *Metronome_note) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(metronome_note, metronome_note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metronome_note *Metronome_note) GongGetReferenceIdentifier(stage *Stage) string {
	return metronome_note.GongGetIdentifier(stage)
}

func (metronome_tied *Metronome_tied) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(metronome_tied, metronome_tied.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metronome_tied *Metronome_tied) GongGetReferenceIdentifier(stage *Stage) string {
	return metronome_tied.GongGetIdentifier(stage)
}

func (metronome_tuplet *Metronome_tuplet) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(metronome_tuplet, metronome_tuplet.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (metronome_tuplet *Metronome_tuplet) GongGetReferenceIdentifier(stage *Stage) string {
	return metronome_tuplet.GongGetIdentifier(stage)
}

func (midi_device *Midi_device) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(midi_device, midi_device.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (midi_device *Midi_device) GongGetReferenceIdentifier(stage *Stage) string {
	return midi_device.GongGetIdentifier(stage)
}

func (midi_instrument *Midi_instrument) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(midi_instrument, midi_instrument.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (midi_instrument *Midi_instrument) GongGetReferenceIdentifier(stage *Stage) string {
	return midi_instrument.GongGetIdentifier(stage)
}

func (miscellaneous *Miscellaneous) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(miscellaneous, miscellaneous.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (miscellaneous *Miscellaneous) GongGetReferenceIdentifier(stage *Stage) string {
	return miscellaneous.GongGetIdentifier(stage)
}

func (miscellaneous_field *Miscellaneous_field) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(miscellaneous_field, miscellaneous_field.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (miscellaneous_field *Miscellaneous_field) GongGetReferenceIdentifier(stage *Stage) string {
	return miscellaneous_field.GongGetIdentifier(stage)
}

func (mordent *Mordent) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(mordent, mordent.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (mordent *Mordent) GongGetReferenceIdentifier(stage *Stage) string {
	return mordent.GongGetIdentifier(stage)
}

func (multiple_rest *Multiple_rest) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(multiple_rest, multiple_rest.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (multiple_rest *Multiple_rest) GongGetReferenceIdentifier(stage *Stage) string {
	return multiple_rest.GongGetIdentifier(stage)
}

func (name_display *Name_display) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(name_display, name_display.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (name_display *Name_display) GongGetReferenceIdentifier(stage *Stage) string {
	return name_display.GongGetIdentifier(stage)
}

func (non_arpeggiate *Non_arpeggiate) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(non_arpeggiate, non_arpeggiate.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (non_arpeggiate *Non_arpeggiate) GongGetReferenceIdentifier(stage *Stage) string {
	return non_arpeggiate.GongGetIdentifier(stage)
}

func (notations *Notations) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notations, notations.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notations *Notations) GongGetReferenceIdentifier(stage *Stage) string {
	return notations.GongGetIdentifier(stage)
}

func (note *Note) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(note, note.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note *Note) GongGetReferenceIdentifier(stage *Stage) string {
	return note.GongGetIdentifier(stage)
}

func (note_size *Note_size) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(note_size, note_size.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note_size *Note_size) GongGetReferenceIdentifier(stage *Stage) string {
	return note_size.GongGetIdentifier(stage)
}

func (note_type *Note_type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(note_type, note_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (note_type *Note_type) GongGetReferenceIdentifier(stage *Stage) string {
	return note_type.GongGetIdentifier(stage)
}

func (notehead *Notehead) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notehead, notehead.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notehead *Notehead) GongGetReferenceIdentifier(stage *Stage) string {
	return notehead.GongGetIdentifier(stage)
}

func (notehead_text *Notehead_text) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(notehead_text, notehead_text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (notehead_text *Notehead_text) GongGetReferenceIdentifier(stage *Stage) string {
	return notehead_text.GongGetIdentifier(stage)
}

func (numeral *Numeral) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(numeral, numeral.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (numeral *Numeral) GongGetReferenceIdentifier(stage *Stage) string {
	return numeral.GongGetIdentifier(stage)
}

func (numeral_key *Numeral_key) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(numeral_key, numeral_key.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (numeral_key *Numeral_key) GongGetReferenceIdentifier(stage *Stage) string {
	return numeral_key.GongGetIdentifier(stage)
}

func (numeral_root *Numeral_root) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(numeral_root, numeral_root.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (numeral_root *Numeral_root) GongGetReferenceIdentifier(stage *Stage) string {
	return numeral_root.GongGetIdentifier(stage)
}

func (octave_shift *Octave_shift) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(octave_shift, octave_shift.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (octave_shift *Octave_shift) GongGetReferenceIdentifier(stage *Stage) string {
	return octave_shift.GongGetIdentifier(stage)
}

func (offset *Offset) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(offset, offset.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (offset *Offset) GongGetReferenceIdentifier(stage *Stage) string {
	return offset.GongGetIdentifier(stage)
}

func (opus *Opus) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(opus, opus.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (opus *Opus) GongGetReferenceIdentifier(stage *Stage) string {
	return opus.GongGetIdentifier(stage)
}

func (ornaments *Ornaments) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(ornaments, ornaments.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (ornaments *Ornaments) GongGetReferenceIdentifier(stage *Stage) string {
	return ornaments.GongGetIdentifier(stage)
}

func (other_appearance *Other_appearance) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(other_appearance, other_appearance.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (other_appearance *Other_appearance) GongGetReferenceIdentifier(stage *Stage) string {
	return other_appearance.GongGetIdentifier(stage)
}

func (other_direction *Other_direction) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(other_direction, other_direction.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (other_direction *Other_direction) GongGetReferenceIdentifier(stage *Stage) string {
	return other_direction.GongGetIdentifier(stage)
}

func (other_listening *Other_listening) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(other_listening, other_listening.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (other_listening *Other_listening) GongGetReferenceIdentifier(stage *Stage) string {
	return other_listening.GongGetIdentifier(stage)
}

func (other_notation *Other_notation) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(other_notation, other_notation.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (other_notation *Other_notation) GongGetReferenceIdentifier(stage *Stage) string {
	return other_notation.GongGetIdentifier(stage)
}

func (other_placement_text *Other_placement_text) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(other_placement_text, other_placement_text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (other_placement_text *Other_placement_text) GongGetReferenceIdentifier(stage *Stage) string {
	return other_placement_text.GongGetIdentifier(stage)
}

func (other_play *Other_play) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(other_play, other_play.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (other_play *Other_play) GongGetReferenceIdentifier(stage *Stage) string {
	return other_play.GongGetIdentifier(stage)
}

func (other_text *Other_text) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(other_text, other_text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (other_text *Other_text) GongGetReferenceIdentifier(stage *Stage) string {
	return other_text.GongGetIdentifier(stage)
}

func (page_layout *Page_layout) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(page_layout, page_layout.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (page_layout *Page_layout) GongGetReferenceIdentifier(stage *Stage) string {
	return page_layout.GongGetIdentifier(stage)
}

func (page_margins *Page_margins) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(page_margins, page_margins.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (page_margins *Page_margins) GongGetReferenceIdentifier(stage *Stage) string {
	return page_margins.GongGetIdentifier(stage)
}

func (part_clef *Part_clef) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(part_clef, part_clef.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part_clef *Part_clef) GongGetReferenceIdentifier(stage *Stage) string {
	return part_clef.GongGetIdentifier(stage)
}

func (part_group *Part_group) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(part_group, part_group.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part_group *Part_group) GongGetReferenceIdentifier(stage *Stage) string {
	return part_group.GongGetIdentifier(stage)
}

func (part_link *Part_link) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(part_link, part_link.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part_link *Part_link) GongGetReferenceIdentifier(stage *Stage) string {
	return part_link.GongGetIdentifier(stage)
}

func (part_list *Part_list) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(part_list, part_list.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part_list *Part_list) GongGetReferenceIdentifier(stage *Stage) string {
	return part_list.GongGetIdentifier(stage)
}

func (part_name *Part_name) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(part_name, part_name.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part_name *Part_name) GongGetReferenceIdentifier(stage *Stage) string {
	return part_name.GongGetIdentifier(stage)
}

func (part_symbol *Part_symbol) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(part_symbol, part_symbol.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part_symbol *Part_symbol) GongGetReferenceIdentifier(stage *Stage) string {
	return part_symbol.GongGetIdentifier(stage)
}

func (part_transpose *Part_transpose) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(part_transpose, part_transpose.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (part_transpose *Part_transpose) GongGetReferenceIdentifier(stage *Stage) string {
	return part_transpose.GongGetIdentifier(stage)
}

func (pedal *Pedal) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pedal, pedal.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pedal *Pedal) GongGetReferenceIdentifier(stage *Stage) string {
	return pedal.GongGetIdentifier(stage)
}

func (pedal_tuning *Pedal_tuning) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pedal_tuning, pedal_tuning.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pedal_tuning *Pedal_tuning) GongGetReferenceIdentifier(stage *Stage) string {
	return pedal_tuning.GongGetIdentifier(stage)
}

func (per_minute *Per_minute) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(per_minute, per_minute.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (per_minute *Per_minute) GongGetReferenceIdentifier(stage *Stage) string {
	return per_minute.GongGetIdentifier(stage)
}

func (percussion *Percussion) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(percussion, percussion.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (percussion *Percussion) GongGetReferenceIdentifier(stage *Stage) string {
	return percussion.GongGetIdentifier(stage)
}

func (pitch *Pitch) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pitch, pitch.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pitch *Pitch) GongGetReferenceIdentifier(stage *Stage) string {
	return pitch.GongGetIdentifier(stage)
}

func (pitched *Pitched) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(pitched, pitched.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (pitched *Pitched) GongGetReferenceIdentifier(stage *Stage) string {
	return pitched.GongGetIdentifier(stage)
}

func (placement_text *Placement_text) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(placement_text, placement_text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (placement_text *Placement_text) GongGetReferenceIdentifier(stage *Stage) string {
	return placement_text.GongGetIdentifier(stage)
}

func (play *Play) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(play, play.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (play *Play) GongGetReferenceIdentifier(stage *Stage) string {
	return play.GongGetIdentifier(stage)
}

func (player *Player) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(player, player.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (player *Player) GongGetReferenceIdentifier(stage *Stage) string {
	return player.GongGetIdentifier(stage)
}

func (principal_voice *Principal_voice) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(principal_voice, principal_voice.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (principal_voice *Principal_voice) GongGetReferenceIdentifier(stage *Stage) string {
	return principal_voice.GongGetIdentifier(stage)
}

func (print *Print) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(print, print.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (print *Print) GongGetReferenceIdentifier(stage *Stage) string {
	return print.GongGetIdentifier(stage)
}

func (release *Release) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(release, release.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (release *Release) GongGetReferenceIdentifier(stage *Stage) string {
	return release.GongGetIdentifier(stage)
}

func (repeat *Repeat) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(repeat, repeat.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (repeat *Repeat) GongGetReferenceIdentifier(stage *Stage) string {
	return repeat.GongGetIdentifier(stage)
}

func (rest *Rest) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(rest, rest.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (rest *Rest) GongGetReferenceIdentifier(stage *Stage) string {
	return rest.GongGetIdentifier(stage)
}

func (root *Root) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(root, root.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (root *Root) GongGetReferenceIdentifier(stage *Stage) string {
	return root.GongGetIdentifier(stage)
}

func (root_step *Root_step) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(root_step, root_step.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (root_step *Root_step) GongGetReferenceIdentifier(stage *Stage) string {
	return root_step.GongGetIdentifier(stage)
}

func (scaling *Scaling) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(scaling, scaling.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (scaling *Scaling) GongGetReferenceIdentifier(stage *Stage) string {
	return scaling.GongGetIdentifier(stage)
}

func (scordatura *Scordatura) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(scordatura, scordatura.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (scordatura *Scordatura) GongGetReferenceIdentifier(stage *Stage) string {
	return scordatura.GongGetIdentifier(stage)
}

func (score_instrument *Score_instrument) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(score_instrument, score_instrument.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (score_instrument *Score_instrument) GongGetReferenceIdentifier(stage *Stage) string {
	return score_instrument.GongGetIdentifier(stage)
}

func (score_part *Score_part) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(score_part, score_part.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (score_part *Score_part) GongGetReferenceIdentifier(stage *Stage) string {
	return score_part.GongGetIdentifier(stage)
}

func (score_partwise *Score_partwise) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(score_partwise, score_partwise.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (score_partwise *Score_partwise) GongGetReferenceIdentifier(stage *Stage) string {
	return score_partwise.GongGetIdentifier(stage)
}

func (score_timewise *Score_timewise) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(score_timewise, score_timewise.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (score_timewise *Score_timewise) GongGetReferenceIdentifier(stage *Stage) string {
	return score_timewise.GongGetIdentifier(stage)
}

func (segno *Segno) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(segno, segno.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (segno *Segno) GongGetReferenceIdentifier(stage *Stage) string {
	return segno.GongGetIdentifier(stage)
}

func (slash *Slash) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(slash, slash.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (slash *Slash) GongGetReferenceIdentifier(stage *Stage) string {
	return slash.GongGetIdentifier(stage)
}

func (slide *Slide) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(slide, slide.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (slide *Slide) GongGetReferenceIdentifier(stage *Stage) string {
	return slide.GongGetIdentifier(stage)
}

func (slur *Slur) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(slur, slur.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (slur *Slur) GongGetReferenceIdentifier(stage *Stage) string {
	return slur.GongGetIdentifier(stage)
}

func (sound *Sound) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(sound, sound.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sound *Sound) GongGetReferenceIdentifier(stage *Stage) string {
	return sound.GongGetIdentifier(stage)
}

func (staff_details *Staff_details) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(staff_details, staff_details.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staff_details *Staff_details) GongGetReferenceIdentifier(stage *Stage) string {
	return staff_details.GongGetIdentifier(stage)
}

func (staff_divide *Staff_divide) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(staff_divide, staff_divide.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staff_divide *Staff_divide) GongGetReferenceIdentifier(stage *Stage) string {
	return staff_divide.GongGetIdentifier(stage)
}

func (staff_layout *Staff_layout) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(staff_layout, staff_layout.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staff_layout *Staff_layout) GongGetReferenceIdentifier(stage *Stage) string {
	return staff_layout.GongGetIdentifier(stage)
}

func (staff_size *Staff_size) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(staff_size, staff_size.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staff_size *Staff_size) GongGetReferenceIdentifier(stage *Stage) string {
	return staff_size.GongGetIdentifier(stage)
}

func (staff_tuning *Staff_tuning) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(staff_tuning, staff_tuning.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (staff_tuning *Staff_tuning) GongGetReferenceIdentifier(stage *Stage) string {
	return staff_tuning.GongGetIdentifier(stage)
}

func (stem *Stem) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stem, stem.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stem *Stem) GongGetReferenceIdentifier(stage *Stage) string {
	return stem.GongGetIdentifier(stage)
}

func (stick *Stick) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(stick, stick.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (stick *Stick) GongGetReferenceIdentifier(stage *Stage) string {
	return stick.GongGetIdentifier(stage)
}

func (string_mute *String_mute) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(string_mute, string_mute.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (string_mute *String_mute) GongGetReferenceIdentifier(stage *Stage) string {
	return string_mute.GongGetIdentifier(stage)
}

func (string_type *String_type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(string_type, string_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (string_type *String_type) GongGetReferenceIdentifier(stage *Stage) string {
	return string_type.GongGetIdentifier(stage)
}

func (strong_accent *Strong_accent) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(strong_accent, strong_accent.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (strong_accent *Strong_accent) GongGetReferenceIdentifier(stage *Stage) string {
	return strong_accent.GongGetIdentifier(stage)
}

func (style_text *Style_text) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(style_text, style_text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (style_text *Style_text) GongGetReferenceIdentifier(stage *Stage) string {
	return style_text.GongGetIdentifier(stage)
}

func (supports *Supports) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(supports, supports.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (supports *Supports) GongGetReferenceIdentifier(stage *Stage) string {
	return supports.GongGetIdentifier(stage)
}

func (swing *Swing) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(swing, swing.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (swing *Swing) GongGetReferenceIdentifier(stage *Stage) string {
	return swing.GongGetIdentifier(stage)
}

func (sync *Sync) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(sync, sync.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (sync *Sync) GongGetReferenceIdentifier(stage *Stage) string {
	return sync.GongGetIdentifier(stage)
}

func (system_dividers *System_dividers) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(system_dividers, system_dividers.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (system_dividers *System_dividers) GongGetReferenceIdentifier(stage *Stage) string {
	return system_dividers.GongGetIdentifier(stage)
}

func (system_layout *System_layout) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(system_layout, system_layout.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (system_layout *System_layout) GongGetReferenceIdentifier(stage *Stage) string {
	return system_layout.GongGetIdentifier(stage)
}

func (system_margins *System_margins) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(system_margins, system_margins.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (system_margins *System_margins) GongGetReferenceIdentifier(stage *Stage) string {
	return system_margins.GongGetIdentifier(stage)
}

func (tap *Tap) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tap, tap.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tap *Tap) GongGetReferenceIdentifier(stage *Stage) string {
	return tap.GongGetIdentifier(stage)
}

func (technical *Technical) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(technical, technical.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (technical *Technical) GongGetReferenceIdentifier(stage *Stage) string {
	return technical.GongGetIdentifier(stage)
}

func (text_element_data *Text_element_data) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(text_element_data, text_element_data.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (text_element_data *Text_element_data) GongGetReferenceIdentifier(stage *Stage) string {
	return text_element_data.GongGetIdentifier(stage)
}

func (tie *Tie) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tie, tie.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tie *Tie) GongGetReferenceIdentifier(stage *Stage) string {
	return tie.GongGetIdentifier(stage)
}

func (tied *Tied) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tied, tied.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tied *Tied) GongGetReferenceIdentifier(stage *Stage) string {
	return tied.GongGetIdentifier(stage)
}

func (time *Time) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(time, time.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (time *Time) GongGetReferenceIdentifier(stage *Stage) string {
	return time.GongGetIdentifier(stage)
}

func (time_modification *Time_modification) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(time_modification, time_modification.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (time_modification *Time_modification) GongGetReferenceIdentifier(stage *Stage) string {
	return time_modification.GongGetIdentifier(stage)
}

func (timpani *Timpani) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(timpani, timpani.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (timpani *Timpani) GongGetReferenceIdentifier(stage *Stage) string {
	return timpani.GongGetIdentifier(stage)
}

func (transpose *Transpose) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(transpose, transpose.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (transpose *Transpose) GongGetReferenceIdentifier(stage *Stage) string {
	return transpose.GongGetIdentifier(stage)
}

func (tremolo *Tremolo) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tremolo, tremolo.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tremolo *Tremolo) GongGetReferenceIdentifier(stage *Stage) string {
	return tremolo.GongGetIdentifier(stage)
}

func (tuplet *Tuplet) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tuplet, tuplet.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tuplet *Tuplet) GongGetReferenceIdentifier(stage *Stage) string {
	return tuplet.GongGetIdentifier(stage)
}

func (tuplet_dot *Tuplet_dot) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tuplet_dot, tuplet_dot.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tuplet_dot *Tuplet_dot) GongGetReferenceIdentifier(stage *Stage) string {
	return tuplet_dot.GongGetIdentifier(stage)
}

func (tuplet_number *Tuplet_number) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tuplet_number, tuplet_number.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tuplet_number *Tuplet_number) GongGetReferenceIdentifier(stage *Stage) string {
	return tuplet_number.GongGetIdentifier(stage)
}

func (tuplet_portion *Tuplet_portion) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tuplet_portion, tuplet_portion.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tuplet_portion *Tuplet_portion) GongGetReferenceIdentifier(stage *Stage) string {
	return tuplet_portion.GongGetIdentifier(stage)
}

func (tuplet_type *Tuplet_type) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(tuplet_type, tuplet_type.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (tuplet_type *Tuplet_type) GongGetReferenceIdentifier(stage *Stage) string {
	return tuplet_type.GongGetIdentifier(stage)
}

func (typed_text *Typed_text) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(typed_text, typed_text.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (typed_text *Typed_text) GongGetReferenceIdentifier(stage *Stage) string {
	return typed_text.GongGetIdentifier(stage)
}

func (unpitched *Unpitched) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(unpitched, unpitched.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (unpitched *Unpitched) GongGetReferenceIdentifier(stage *Stage) string {
	return unpitched.GongGetIdentifier(stage)
}

func (virtual_instrument *Virtual_instrument) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(virtual_instrument, virtual_instrument.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (virtual_instrument *Virtual_instrument) GongGetReferenceIdentifier(stage *Stage) string {
	return virtual_instrument.GongGetIdentifier(stage)
}

func (wait *Wait) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(wait, wait.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (wait *Wait) GongGetReferenceIdentifier(stage *Stage) string {
	return wait.GongGetIdentifier(stage)
}

func (wavy_line *Wavy_line) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(wavy_line, wavy_line.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (wavy_line *Wavy_line) GongGetReferenceIdentifier(stage *Stage) string {
	return wavy_line.GongGetIdentifier(stage)
}

func (wedge *Wedge) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(wedge, wedge.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (wedge *Wedge) GongGetReferenceIdentifier(stage *Stage) string {
	return wedge.GongGetIdentifier(stage)
}

func (wood *Wood) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(wood, wood.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (wood *Wood) GongGetReferenceIdentifier(stage *Stage) string {
	return wood.GongGetIdentifier(stage)
}

func (work *Work) GongGetIdentifier(stage *Stage) string {
	return __gong__formatIdentifier(work, work.GongGetOrder(stage))
}

// GongGetReferenceIdentifier returns an identifier when it was staged (it may have been unstaged since)
func (work *Work) GongGetReferenceIdentifier(stage *Stage) string {
	return work.GongGetIdentifier(stage)
}

// MarshallIdentifier returns the code to instantiate the instance
// in a marshalling file
// insertion point per named struct
func (a_directive *A_directive) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_directive.GongGetIdentifier(stage), "A_directive", a_directive.Name)
}

func (a_measure *A_measure) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_measure.GongGetIdentifier(stage), "A_measure", a_measure.Name)
}

func (a_measure_1 *A_measure_1) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_measure_1.GongGetIdentifier(stage), "A_measure_1", a_measure_1.Name)
}

func (a_part *A_part) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_part.GongGetIdentifier(stage), "A_part", a_part.Name)
}

func (a_part_1 *A_part_1) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(a_part_1.GongGetIdentifier(stage), "A_part_1", a_part_1.Name)
}

func (accidental *Accidental) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(accidental.GongGetIdentifier(stage), "Accidental", accidental.Name)
}

func (accidental_mark *Accidental_mark) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(accidental_mark.GongGetIdentifier(stage), "Accidental_mark", accidental_mark.Name)
}

func (accidental_text *Accidental_text) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(accidental_text.GongGetIdentifier(stage), "Accidental_text", accidental_text.Name)
}

func (accord *Accord) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(accord.GongGetIdentifier(stage), "Accord", accord.Name)
}

func (accordion_registration *Accordion_registration) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(accordion_registration.GongGetIdentifier(stage), "Accordion_registration", accordion_registration.Name)
}

func (appearance *Appearance) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(appearance.GongGetIdentifier(stage), "Appearance", appearance.Name)
}

func (arpeggiate *Arpeggiate) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(arpeggiate.GongGetIdentifier(stage), "Arpeggiate", arpeggiate.Name)
}

func (arrow *Arrow) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(arrow.GongGetIdentifier(stage), "Arrow", arrow.Name)
}

func (articulations *Articulations) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(articulations.GongGetIdentifier(stage), "Articulations", articulations.Name)
}

func (assess *Assess) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(assess.GongGetIdentifier(stage), "Assess", assess.Name)
}

func (attributes *Attributes) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(attributes.GongGetIdentifier(stage), "Attributes", attributes.Name)
}

func (backup *Backup) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(backup.GongGetIdentifier(stage), "Backup", backup.Name)
}

func (bar_style_color *Bar_style_color) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(bar_style_color.GongGetIdentifier(stage), "Bar_style_color", bar_style_color.Name)
}

func (barline *Barline) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(barline.GongGetIdentifier(stage), "Barline", barline.Name)
}

func (barre *Barre) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(barre.GongGetIdentifier(stage), "Barre", barre.Name)
}

func (bass *Bass) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(bass.GongGetIdentifier(stage), "Bass", bass.Name)
}

func (bass_step *Bass_step) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(bass_step.GongGetIdentifier(stage), "Bass_step", bass_step.Name)
}

func (beam *Beam) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(beam.GongGetIdentifier(stage), "Beam", beam.Name)
}

func (beat_repeat *Beat_repeat) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(beat_repeat.GongGetIdentifier(stage), "Beat_repeat", beat_repeat.Name)
}

func (beat_unit_tied *Beat_unit_tied) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(beat_unit_tied.GongGetIdentifier(stage), "Beat_unit_tied", beat_unit_tied.Name)
}

func (beater *Beater) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(beater.GongGetIdentifier(stage), "Beater", beater.Name)
}

func (bend *Bend) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(bend.GongGetIdentifier(stage), "Bend", bend.Name)
}

func (bookmark *Bookmark) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(bookmark.GongGetIdentifier(stage), "Bookmark", bookmark.Name)
}

func (bracket *Bracket) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(bracket.GongGetIdentifier(stage), "Bracket", bracket.Name)
}

func (breath_mark *Breath_mark) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(breath_mark.GongGetIdentifier(stage), "Breath_mark", breath_mark.Name)
}

func (caesura *Caesura) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(caesura.GongGetIdentifier(stage), "Caesura", caesura.Name)
}

func (cancel *Cancel) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(cancel.GongGetIdentifier(stage), "Cancel", cancel.Name)
}

func (clef *Clef) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(clef.GongGetIdentifier(stage), "Clef", clef.Name)
}

func (coda *Coda) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(coda.GongGetIdentifier(stage), "Coda", coda.Name)
}

func (credit *Credit) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(credit.GongGetIdentifier(stage), "Credit", credit.Name)
}

func (dashes *Dashes) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(dashes.GongGetIdentifier(stage), "Dashes", dashes.Name)
}

func (defaults *Defaults) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(defaults.GongGetIdentifier(stage), "Defaults", defaults.Name)
}

func (degree *Degree) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(degree.GongGetIdentifier(stage), "Degree", degree.Name)
}

func (degree_alter *Degree_alter) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(degree_alter.GongGetIdentifier(stage), "Degree_alter", degree_alter.Name)
}

func (degree_type *Degree_type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(degree_type.GongGetIdentifier(stage), "Degree_type", degree_type.Name)
}

func (degree_value *Degree_value) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(degree_value.GongGetIdentifier(stage), "Degree_value", degree_value.Name)
}

func (direction *Direction) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(direction.GongGetIdentifier(stage), "Direction", direction.Name)
}

func (direction_type *Direction_type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(direction_type.GongGetIdentifier(stage), "Direction_type", direction_type.Name)
}

func (distance *Distance) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(distance.GongGetIdentifier(stage), "Distance", distance.Name)
}

func (double *Double) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(double.GongGetIdentifier(stage), "Double", double.Name)
}

func (dynamics *Dynamics) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(dynamics.GongGetIdentifier(stage), "Dynamics", dynamics.Name)
}

func (effect *Effect) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(effect.GongGetIdentifier(stage), "Effect", effect.Name)
}

func (elision *Elision) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(elision.GongGetIdentifier(stage), "Elision", elision.Name)
}

func (empty *Empty) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(empty.GongGetIdentifier(stage), "Empty", empty.Name)
}

func (empty_font *Empty_font) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(empty_font.GongGetIdentifier(stage), "Empty_font", empty_font.Name)
}

func (empty_line *Empty_line) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(empty_line.GongGetIdentifier(stage), "Empty_line", empty_line.Name)
}

func (empty_placement *Empty_placement) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(empty_placement.GongGetIdentifier(stage), "Empty_placement", empty_placement.Name)
}

func (empty_placement_smufl *Empty_placement_smufl) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(empty_placement_smufl.GongGetIdentifier(stage), "Empty_placement_smufl", empty_placement_smufl.Name)
}

func (empty_print_object_style_align *Empty_print_object_style_align) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(empty_print_object_style_align.GongGetIdentifier(stage), "Empty_print_object_style_align", empty_print_object_style_align.Name)
}

func (empty_print_style *Empty_print_style) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(empty_print_style.GongGetIdentifier(stage), "Empty_print_style", empty_print_style.Name)
}

func (empty_print_style_align *Empty_print_style_align) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(empty_print_style_align.GongGetIdentifier(stage), "Empty_print_style_align", empty_print_style_align.Name)
}

func (empty_print_style_align_id *Empty_print_style_align_id) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(empty_print_style_align_id.GongGetIdentifier(stage), "Empty_print_style_align_id", empty_print_style_align_id.Name)
}

func (empty_trill_sound *Empty_trill_sound) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(empty_trill_sound.GongGetIdentifier(stage), "Empty_trill_sound", empty_trill_sound.Name)
}

func (encoding *Encoding) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(encoding.GongGetIdentifier(stage), "Encoding", encoding.Name)
}

func (ending *Ending) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ending.GongGetIdentifier(stage), "Ending", ending.Name)
}

func (extend *Extend) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(extend.GongGetIdentifier(stage), "Extend", extend.Name)
}

func (feature *Feature) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(feature.GongGetIdentifier(stage), "Feature", feature.Name)
}

func (fermata *Fermata) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(fermata.GongGetIdentifier(stage), "Fermata", fermata.Name)
}

func (figure *Figure) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(figure.GongGetIdentifier(stage), "Figure", figure.Name)
}

func (figured_bass *Figured_bass) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(figured_bass.GongGetIdentifier(stage), "Figured_bass", figured_bass.Name)
}

func (fingering *Fingering) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(fingering.GongGetIdentifier(stage), "Fingering", fingering.Name)
}

func (first_fret *First_fret) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(first_fret.GongGetIdentifier(stage), "First_fret", first_fret.Name)
}

func (for_part *For_part) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(for_part.GongGetIdentifier(stage), "For_part", for_part.Name)
}

func (formatted_symbol *Formatted_symbol) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formatted_symbol.GongGetIdentifier(stage), "Formatted_symbol", formatted_symbol.Name)
}

func (formatted_symbol_id *Formatted_symbol_id) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formatted_symbol_id.GongGetIdentifier(stage), "Formatted_symbol_id", formatted_symbol_id.Name)
}

func (formatted_text *Formatted_text) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formatted_text.GongGetIdentifier(stage), "Formatted_text", formatted_text.Name)
}

func (formatted_text_id *Formatted_text_id) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(formatted_text_id.GongGetIdentifier(stage), "Formatted_text_id", formatted_text_id.Name)
}

func (forward *Forward) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(forward.GongGetIdentifier(stage), "Forward", forward.Name)
}

func (frame *Frame) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(frame.GongGetIdentifier(stage), "Frame", frame.Name)
}

func (frame_note *Frame_note) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(frame_note.GongGetIdentifier(stage), "Frame_note", frame_note.Name)
}

func (fret *Fret) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(fret.GongGetIdentifier(stage), "Fret", fret.Name)
}

func (glass *Glass) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(glass.GongGetIdentifier(stage), "Glass", glass.Name)
}

func (glissando *Glissando) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(glissando.GongGetIdentifier(stage), "Glissando", glissando.Name)
}

func (glyph *Glyph) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(glyph.GongGetIdentifier(stage), "Glyph", glyph.Name)
}

func (grace *Grace) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(grace.GongGetIdentifier(stage), "Grace", grace.Name)
}

func (group_barline *Group_barline) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(group_barline.GongGetIdentifier(stage), "Group_barline", group_barline.Name)
}

func (group_name *Group_name) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(group_name.GongGetIdentifier(stage), "Group_name", group_name.Name)
}

func (group_symbol *Group_symbol) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(group_symbol.GongGetIdentifier(stage), "Group_symbol", group_symbol.Name)
}

func (grouping *Grouping) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(grouping.GongGetIdentifier(stage), "Grouping", grouping.Name)
}

func (hammer_on_pull_off *Hammer_on_pull_off) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(hammer_on_pull_off.GongGetIdentifier(stage), "Hammer_on_pull_off", hammer_on_pull_off.Name)
}

func (handbell *Handbell) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(handbell.GongGetIdentifier(stage), "Handbell", handbell.Name)
}

func (harmon_closed *Harmon_closed) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(harmon_closed.GongGetIdentifier(stage), "Harmon_closed", harmon_closed.Name)
}

func (harmon_mute *Harmon_mute) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(harmon_mute.GongGetIdentifier(stage), "Harmon_mute", harmon_mute.Name)
}

func (harmonic *Harmonic) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(harmonic.GongGetIdentifier(stage), "Harmonic", harmonic.Name)
}

func (harmony *Harmony) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(harmony.GongGetIdentifier(stage), "Harmony", harmony.Name)
}

func (harmony_alter *Harmony_alter) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(harmony_alter.GongGetIdentifier(stage), "Harmony_alter", harmony_alter.Name)
}

func (harp_pedals *Harp_pedals) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(harp_pedals.GongGetIdentifier(stage), "Harp_pedals", harp_pedals.Name)
}

func (heel_toe *Heel_toe) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(heel_toe.GongGetIdentifier(stage), "Heel_toe", heel_toe.Name)
}

func (hole *Hole) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(hole.GongGetIdentifier(stage), "Hole", hole.Name)
}

func (hole_closed *Hole_closed) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(hole_closed.GongGetIdentifier(stage), "Hole_closed", hole_closed.Name)
}

func (horizontal_turn *Horizontal_turn) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(horizontal_turn.GongGetIdentifier(stage), "Horizontal_turn", horizontal_turn.Name)
}

func (identification *Identification) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(identification.GongGetIdentifier(stage), "Identification", identification.Name)
}

func (image *Image) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(image.GongGetIdentifier(stage), "Image", image.Name)
}

func (instrument *Instrument) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(instrument.GongGetIdentifier(stage), "Instrument", instrument.Name)
}

func (instrument_change *Instrument_change) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(instrument_change.GongGetIdentifier(stage), "Instrument_change", instrument_change.Name)
}

func (instrument_link *Instrument_link) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(instrument_link.GongGetIdentifier(stage), "Instrument_link", instrument_link.Name)
}

func (interchangeable *Interchangeable) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(interchangeable.GongGetIdentifier(stage), "Interchangeable", interchangeable.Name)
}

func (inversion *Inversion) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(inversion.GongGetIdentifier(stage), "Inversion", inversion.Name)
}

func (key *Key) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(key.GongGetIdentifier(stage), "Key", key.Name)
}

func (key_accidental *Key_accidental) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(key_accidental.GongGetIdentifier(stage), "Key_accidental", key_accidental.Name)
}

func (key_octave *Key_octave) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(key_octave.GongGetIdentifier(stage), "Key_octave", key_octave.Name)
}

func (kind *Kind) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(kind.GongGetIdentifier(stage), "Kind", kind.Name)
}

func (level *Level) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(level.GongGetIdentifier(stage), "Level", level.Name)
}

func (line_detail *Line_detail) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(line_detail.GongGetIdentifier(stage), "Line_detail", line_detail.Name)
}

func (line_width *Line_width) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(line_width.GongGetIdentifier(stage), "Line_width", line_width.Name)
}

func (link *Link) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(link.GongGetIdentifier(stage), "Link", link.Name)
}

func (listen *Listen) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(listen.GongGetIdentifier(stage), "Listen", listen.Name)
}

func (listening *Listening) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(listening.GongGetIdentifier(stage), "Listening", listening.Name)
}

func (lyric *Lyric) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(lyric.GongGetIdentifier(stage), "Lyric", lyric.Name)
}

func (lyric_font *Lyric_font) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(lyric_font.GongGetIdentifier(stage), "Lyric_font", lyric_font.Name)
}

func (lyric_language *Lyric_language) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(lyric_language.GongGetIdentifier(stage), "Lyric_language", lyric_language.Name)
}

func (measure_layout *Measure_layout) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(measure_layout.GongGetIdentifier(stage), "Measure_layout", measure_layout.Name)
}

func (measure_numbering *Measure_numbering) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(measure_numbering.GongGetIdentifier(stage), "Measure_numbering", measure_numbering.Name)
}

func (measure_repeat *Measure_repeat) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(measure_repeat.GongGetIdentifier(stage), "Measure_repeat", measure_repeat.Name)
}

func (measure_style *Measure_style) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(measure_style.GongGetIdentifier(stage), "Measure_style", measure_style.Name)
}

func (membrane *Membrane) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(membrane.GongGetIdentifier(stage), "Membrane", membrane.Name)
}

func (metal *Metal) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(metal.GongGetIdentifier(stage), "Metal", metal.Name)
}

func (metronome *Metronome) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(metronome.GongGetIdentifier(stage), "Metronome", metronome.Name)
}

func (metronome_beam *Metronome_beam) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(metronome_beam.GongGetIdentifier(stage), "Metronome_beam", metronome_beam.Name)
}

func (metronome_note *Metronome_note) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(metronome_note.GongGetIdentifier(stage), "Metronome_note", metronome_note.Name)
}

func (metronome_tied *Metronome_tied) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(metronome_tied.GongGetIdentifier(stage), "Metronome_tied", metronome_tied.Name)
}

func (metronome_tuplet *Metronome_tuplet) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(metronome_tuplet.GongGetIdentifier(stage), "Metronome_tuplet", metronome_tuplet.Name)
}

func (midi_device *Midi_device) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(midi_device.GongGetIdentifier(stage), "Midi_device", midi_device.Name)
}

func (midi_instrument *Midi_instrument) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(midi_instrument.GongGetIdentifier(stage), "Midi_instrument", midi_instrument.Name)
}

func (miscellaneous *Miscellaneous) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(miscellaneous.GongGetIdentifier(stage), "Miscellaneous", miscellaneous.Name)
}

func (miscellaneous_field *Miscellaneous_field) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(miscellaneous_field.GongGetIdentifier(stage), "Miscellaneous_field", miscellaneous_field.Name)
}

func (mordent *Mordent) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(mordent.GongGetIdentifier(stage), "Mordent", mordent.Name)
}

func (multiple_rest *Multiple_rest) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(multiple_rest.GongGetIdentifier(stage), "Multiple_rest", multiple_rest.Name)
}

func (name_display *Name_display) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(name_display.GongGetIdentifier(stage), "Name_display", name_display.Name)
}

func (non_arpeggiate *Non_arpeggiate) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(non_arpeggiate.GongGetIdentifier(stage), "Non_arpeggiate", non_arpeggiate.Name)
}

func (notations *Notations) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notations.GongGetIdentifier(stage), "Notations", notations.Name)
}

func (note *Note) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(note.GongGetIdentifier(stage), "Note", note.Name)
}

func (note_size *Note_size) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(note_size.GongGetIdentifier(stage), "Note_size", note_size.Name)
}

func (note_type *Note_type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(note_type.GongGetIdentifier(stage), "Note_type", note_type.Name)
}

func (notehead *Notehead) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notehead.GongGetIdentifier(stage), "Notehead", notehead.Name)
}

func (notehead_text *Notehead_text) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(notehead_text.GongGetIdentifier(stage), "Notehead_text", notehead_text.Name)
}

func (numeral *Numeral) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(numeral.GongGetIdentifier(stage), "Numeral", numeral.Name)
}

func (numeral_key *Numeral_key) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(numeral_key.GongGetIdentifier(stage), "Numeral_key", numeral_key.Name)
}

func (numeral_root *Numeral_root) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(numeral_root.GongGetIdentifier(stage), "Numeral_root", numeral_root.Name)
}

func (octave_shift *Octave_shift) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(octave_shift.GongGetIdentifier(stage), "Octave_shift", octave_shift.Name)
}

func (offset *Offset) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(offset.GongGetIdentifier(stage), "Offset", offset.Name)
}

func (opus *Opus) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(opus.GongGetIdentifier(stage), "Opus", opus.Name)
}

func (ornaments *Ornaments) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(ornaments.GongGetIdentifier(stage), "Ornaments", ornaments.Name)
}

func (other_appearance *Other_appearance) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(other_appearance.GongGetIdentifier(stage), "Other_appearance", other_appearance.Name)
}

func (other_direction *Other_direction) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(other_direction.GongGetIdentifier(stage), "Other_direction", other_direction.Name)
}

func (other_listening *Other_listening) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(other_listening.GongGetIdentifier(stage), "Other_listening", other_listening.Name)
}

func (other_notation *Other_notation) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(other_notation.GongGetIdentifier(stage), "Other_notation", other_notation.Name)
}

func (other_placement_text *Other_placement_text) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(other_placement_text.GongGetIdentifier(stage), "Other_placement_text", other_placement_text.Name)
}

func (other_play *Other_play) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(other_play.GongGetIdentifier(stage), "Other_play", other_play.Name)
}

func (other_text *Other_text) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(other_text.GongGetIdentifier(stage), "Other_text", other_text.Name)
}

func (page_layout *Page_layout) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(page_layout.GongGetIdentifier(stage), "Page_layout", page_layout.Name)
}

func (page_margins *Page_margins) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(page_margins.GongGetIdentifier(stage), "Page_margins", page_margins.Name)
}

func (part_clef *Part_clef) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(part_clef.GongGetIdentifier(stage), "Part_clef", part_clef.Name)
}

func (part_group *Part_group) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(part_group.GongGetIdentifier(stage), "Part_group", part_group.Name)
}

func (part_link *Part_link) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(part_link.GongGetIdentifier(stage), "Part_link", part_link.Name)
}

func (part_list *Part_list) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(part_list.GongGetIdentifier(stage), "Part_list", part_list.Name)
}

func (part_name *Part_name) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(part_name.GongGetIdentifier(stage), "Part_name", part_name.Name)
}

func (part_symbol *Part_symbol) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(part_symbol.GongGetIdentifier(stage), "Part_symbol", part_symbol.Name)
}

func (part_transpose *Part_transpose) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(part_transpose.GongGetIdentifier(stage), "Part_transpose", part_transpose.Name)
}

func (pedal *Pedal) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pedal.GongGetIdentifier(stage), "Pedal", pedal.Name)
}

func (pedal_tuning *Pedal_tuning) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pedal_tuning.GongGetIdentifier(stage), "Pedal_tuning", pedal_tuning.Name)
}

func (per_minute *Per_minute) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(per_minute.GongGetIdentifier(stage), "Per_minute", per_minute.Name)
}

func (percussion *Percussion) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(percussion.GongGetIdentifier(stage), "Percussion", percussion.Name)
}

func (pitch *Pitch) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pitch.GongGetIdentifier(stage), "Pitch", pitch.Name)
}

func (pitched *Pitched) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(pitched.GongGetIdentifier(stage), "Pitched", pitched.Name)
}

func (placement_text *Placement_text) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(placement_text.GongGetIdentifier(stage), "Placement_text", placement_text.Name)
}

func (play *Play) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(play.GongGetIdentifier(stage), "Play", play.Name)
}

func (player *Player) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(player.GongGetIdentifier(stage), "Player", player.Name)
}

func (principal_voice *Principal_voice) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(principal_voice.GongGetIdentifier(stage), "Principal_voice", principal_voice.Name)
}

func (print *Print) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(print.GongGetIdentifier(stage), "Print", print.Name)
}

func (release *Release) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(release.GongGetIdentifier(stage), "Release", release.Name)
}

func (repeat *Repeat) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(repeat.GongGetIdentifier(stage), "Repeat", repeat.Name)
}

func (rest *Rest) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(rest.GongGetIdentifier(stage), "Rest", rest.Name)
}

func (root *Root) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(root.GongGetIdentifier(stage), "Root", root.Name)
}

func (root_step *Root_step) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(root_step.GongGetIdentifier(stage), "Root_step", root_step.Name)
}

func (scaling *Scaling) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(scaling.GongGetIdentifier(stage), "Scaling", scaling.Name)
}

func (scordatura *Scordatura) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(scordatura.GongGetIdentifier(stage), "Scordatura", scordatura.Name)
}

func (score_instrument *Score_instrument) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(score_instrument.GongGetIdentifier(stage), "Score_instrument", score_instrument.Name)
}

func (score_part *Score_part) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(score_part.GongGetIdentifier(stage), "Score_part", score_part.Name)
}

func (score_partwise *Score_partwise) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(score_partwise.GongGetIdentifier(stage), "Score_partwise", score_partwise.Name)
}

func (score_timewise *Score_timewise) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(score_timewise.GongGetIdentifier(stage), "Score_timewise", score_timewise.Name)
}

func (segno *Segno) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(segno.GongGetIdentifier(stage), "Segno", segno.Name)
}

func (slash *Slash) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(slash.GongGetIdentifier(stage), "Slash", slash.Name)
}

func (slide *Slide) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(slide.GongGetIdentifier(stage), "Slide", slide.Name)
}

func (slur *Slur) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(slur.GongGetIdentifier(stage), "Slur", slur.Name)
}

func (sound *Sound) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(sound.GongGetIdentifier(stage), "Sound", sound.Name)
}

func (staff_details *Staff_details) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(staff_details.GongGetIdentifier(stage), "Staff_details", staff_details.Name)
}

func (staff_divide *Staff_divide) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(staff_divide.GongGetIdentifier(stage), "Staff_divide", staff_divide.Name)
}

func (staff_layout *Staff_layout) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(staff_layout.GongGetIdentifier(stage), "Staff_layout", staff_layout.Name)
}

func (staff_size *Staff_size) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(staff_size.GongGetIdentifier(stage), "Staff_size", staff_size.Name)
}

func (staff_tuning *Staff_tuning) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(staff_tuning.GongGetIdentifier(stage), "Staff_tuning", staff_tuning.Name)
}

func (stem *Stem) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stem.GongGetIdentifier(stage), "Stem", stem.Name)
}

func (stick *Stick) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(stick.GongGetIdentifier(stage), "Stick", stick.Name)
}

func (string_mute *String_mute) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(string_mute.GongGetIdentifier(stage), "String_mute", string_mute.Name)
}

func (string_type *String_type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(string_type.GongGetIdentifier(stage), "String_type", string_type.Name)
}

func (strong_accent *Strong_accent) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(strong_accent.GongGetIdentifier(stage), "Strong_accent", strong_accent.Name)
}

func (style_text *Style_text) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(style_text.GongGetIdentifier(stage), "Style_text", style_text.Name)
}

func (supports *Supports) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(supports.GongGetIdentifier(stage), "Supports", supports.Name)
}

func (swing *Swing) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(swing.GongGetIdentifier(stage), "Swing", swing.Name)
}

func (sync *Sync) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(sync.GongGetIdentifier(stage), "Sync", sync.Name)
}

func (system_dividers *System_dividers) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(system_dividers.GongGetIdentifier(stage), "System_dividers", system_dividers.Name)
}

func (system_layout *System_layout) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(system_layout.GongGetIdentifier(stage), "System_layout", system_layout.Name)
}

func (system_margins *System_margins) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(system_margins.GongGetIdentifier(stage), "System_margins", system_margins.Name)
}

func (tap *Tap) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tap.GongGetIdentifier(stage), "Tap", tap.Name)
}

func (technical *Technical) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(technical.GongGetIdentifier(stage), "Technical", technical.Name)
}

func (text_element_data *Text_element_data) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(text_element_data.GongGetIdentifier(stage), "Text_element_data", text_element_data.Name)
}

func (tie *Tie) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tie.GongGetIdentifier(stage), "Tie", tie.Name)
}

func (tied *Tied) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tied.GongGetIdentifier(stage), "Tied", tied.Name)
}

func (time *Time) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(time.GongGetIdentifier(stage), "Time", time.Name)
}

func (time_modification *Time_modification) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(time_modification.GongGetIdentifier(stage), "Time_modification", time_modification.Name)
}

func (timpani *Timpani) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(timpani.GongGetIdentifier(stage), "Timpani", timpani.Name)
}

func (transpose *Transpose) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(transpose.GongGetIdentifier(stage), "Transpose", transpose.Name)
}

func (tremolo *Tremolo) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tremolo.GongGetIdentifier(stage), "Tremolo", tremolo.Name)
}

func (tuplet *Tuplet) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tuplet.GongGetIdentifier(stage), "Tuplet", tuplet.Name)
}

func (tuplet_dot *Tuplet_dot) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tuplet_dot.GongGetIdentifier(stage), "Tuplet_dot", tuplet_dot.Name)
}

func (tuplet_number *Tuplet_number) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tuplet_number.GongGetIdentifier(stage), "Tuplet_number", tuplet_number.Name)
}

func (tuplet_portion *Tuplet_portion) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tuplet_portion.GongGetIdentifier(stage), "Tuplet_portion", tuplet_portion.Name)
}

func (tuplet_type *Tuplet_type) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(tuplet_type.GongGetIdentifier(stage), "Tuplet_type", tuplet_type.Name)
}

func (typed_text *Typed_text) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(typed_text.GongGetIdentifier(stage), "Typed_text", typed_text.Name)
}

func (unpitched *Unpitched) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(unpitched.GongGetIdentifier(stage), "Unpitched", unpitched.Name)
}

func (virtual_instrument *Virtual_instrument) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(virtual_instrument.GongGetIdentifier(stage), "Virtual_instrument", virtual_instrument.Name)
}

func (wait *Wait) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(wait.GongGetIdentifier(stage), "Wait", wait.Name)
}

func (wavy_line *Wavy_line) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(wavy_line.GongGetIdentifier(stage), "Wavy_line", wavy_line.Name)
}

func (wedge *Wedge) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(wedge.GongGetIdentifier(stage), "Wedge", wedge.Name)
}

func (wood *Wood) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(wood.GongGetIdentifier(stage), "Wood", wood.Name)
}

func (work *Work) GongMarshallIdentifier(stage *Stage) string {
	return __gong__marshallIdentifier(work.GongGetIdentifier(stage), "Work", work.Name)
}

// insertion point for unstaging
func (a_directive *A_directive) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_directive.GongGetReferenceIdentifier(stage))
}

func (a_measure *A_measure) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_measure.GongGetReferenceIdentifier(stage))
}

func (a_measure_1 *A_measure_1) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_measure_1.GongGetReferenceIdentifier(stage))
}

func (a_part *A_part) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_part.GongGetReferenceIdentifier(stage))
}

func (a_part_1 *A_part_1) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(a_part_1.GongGetReferenceIdentifier(stage))
}

func (accidental *Accidental) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(accidental.GongGetReferenceIdentifier(stage))
}

func (accidental_mark *Accidental_mark) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(accidental_mark.GongGetReferenceIdentifier(stage))
}

func (accidental_text *Accidental_text) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(accidental_text.GongGetReferenceIdentifier(stage))
}

func (accord *Accord) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(accord.GongGetReferenceIdentifier(stage))
}

func (accordion_registration *Accordion_registration) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(accordion_registration.GongGetReferenceIdentifier(stage))
}

func (appearance *Appearance) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(appearance.GongGetReferenceIdentifier(stage))
}

func (arpeggiate *Arpeggiate) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(arpeggiate.GongGetReferenceIdentifier(stage))
}

func (arrow *Arrow) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(arrow.GongGetReferenceIdentifier(stage))
}

func (articulations *Articulations) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(articulations.GongGetReferenceIdentifier(stage))
}

func (assess *Assess) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(assess.GongGetReferenceIdentifier(stage))
}

func (attributes *Attributes) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(attributes.GongGetReferenceIdentifier(stage))
}

func (backup *Backup) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(backup.GongGetReferenceIdentifier(stage))
}

func (bar_style_color *Bar_style_color) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(bar_style_color.GongGetReferenceIdentifier(stage))
}

func (barline *Barline) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(barline.GongGetReferenceIdentifier(stage))
}

func (barre *Barre) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(barre.GongGetReferenceIdentifier(stage))
}

func (bass *Bass) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(bass.GongGetReferenceIdentifier(stage))
}

func (bass_step *Bass_step) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(bass_step.GongGetReferenceIdentifier(stage))
}

func (beam *Beam) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(beam.GongGetReferenceIdentifier(stage))
}

func (beat_repeat *Beat_repeat) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(beat_repeat.GongGetReferenceIdentifier(stage))
}

func (beat_unit_tied *Beat_unit_tied) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(beat_unit_tied.GongGetReferenceIdentifier(stage))
}

func (beater *Beater) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(beater.GongGetReferenceIdentifier(stage))
}

func (bend *Bend) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(bend.GongGetReferenceIdentifier(stage))
}

func (bookmark *Bookmark) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(bookmark.GongGetReferenceIdentifier(stage))
}

func (bracket *Bracket) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(bracket.GongGetReferenceIdentifier(stage))
}

func (breath_mark *Breath_mark) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(breath_mark.GongGetReferenceIdentifier(stage))
}

func (caesura *Caesura) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(caesura.GongGetReferenceIdentifier(stage))
}

func (cancel *Cancel) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(cancel.GongGetReferenceIdentifier(stage))
}

func (clef *Clef) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(clef.GongGetReferenceIdentifier(stage))
}

func (coda *Coda) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(coda.GongGetReferenceIdentifier(stage))
}

func (credit *Credit) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(credit.GongGetReferenceIdentifier(stage))
}

func (dashes *Dashes) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(dashes.GongGetReferenceIdentifier(stage))
}

func (defaults *Defaults) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(defaults.GongGetReferenceIdentifier(stage))
}

func (degree *Degree) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(degree.GongGetReferenceIdentifier(stage))
}

func (degree_alter *Degree_alter) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(degree_alter.GongGetReferenceIdentifier(stage))
}

func (degree_type *Degree_type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(degree_type.GongGetReferenceIdentifier(stage))
}

func (degree_value *Degree_value) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(degree_value.GongGetReferenceIdentifier(stage))
}

func (direction *Direction) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(direction.GongGetReferenceIdentifier(stage))
}

func (direction_type *Direction_type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(direction_type.GongGetReferenceIdentifier(stage))
}

func (distance *Distance) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(distance.GongGetReferenceIdentifier(stage))
}

func (double *Double) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(double.GongGetReferenceIdentifier(stage))
}

func (dynamics *Dynamics) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(dynamics.GongGetReferenceIdentifier(stage))
}

func (effect *Effect) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(effect.GongGetReferenceIdentifier(stage))
}

func (elision *Elision) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(elision.GongGetReferenceIdentifier(stage))
}

func (empty *Empty) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(empty.GongGetReferenceIdentifier(stage))
}

func (empty_font *Empty_font) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(empty_font.GongGetReferenceIdentifier(stage))
}

func (empty_line *Empty_line) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(empty_line.GongGetReferenceIdentifier(stage))
}

func (empty_placement *Empty_placement) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(empty_placement.GongGetReferenceIdentifier(stage))
}

func (empty_placement_smufl *Empty_placement_smufl) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(empty_placement_smufl.GongGetReferenceIdentifier(stage))
}

func (empty_print_object_style_align *Empty_print_object_style_align) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(empty_print_object_style_align.GongGetReferenceIdentifier(stage))
}

func (empty_print_style *Empty_print_style) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(empty_print_style.GongGetReferenceIdentifier(stage))
}

func (empty_print_style_align *Empty_print_style_align) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(empty_print_style_align.GongGetReferenceIdentifier(stage))
}

func (empty_print_style_align_id *Empty_print_style_align_id) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(empty_print_style_align_id.GongGetReferenceIdentifier(stage))
}

func (empty_trill_sound *Empty_trill_sound) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(empty_trill_sound.GongGetReferenceIdentifier(stage))
}

func (encoding *Encoding) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(encoding.GongGetReferenceIdentifier(stage))
}

func (ending *Ending) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ending.GongGetReferenceIdentifier(stage))
}

func (extend *Extend) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(extend.GongGetReferenceIdentifier(stage))
}

func (feature *Feature) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(feature.GongGetReferenceIdentifier(stage))
}

func (fermata *Fermata) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(fermata.GongGetReferenceIdentifier(stage))
}

func (figure *Figure) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(figure.GongGetReferenceIdentifier(stage))
}

func (figured_bass *Figured_bass) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(figured_bass.GongGetReferenceIdentifier(stage))
}

func (fingering *Fingering) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(fingering.GongGetReferenceIdentifier(stage))
}

func (first_fret *First_fret) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(first_fret.GongGetReferenceIdentifier(stage))
}

func (for_part *For_part) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(for_part.GongGetReferenceIdentifier(stage))
}

func (formatted_symbol *Formatted_symbol) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formatted_symbol.GongGetReferenceIdentifier(stage))
}

func (formatted_symbol_id *Formatted_symbol_id) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formatted_symbol_id.GongGetReferenceIdentifier(stage))
}

func (formatted_text *Formatted_text) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formatted_text.GongGetReferenceIdentifier(stage))
}

func (formatted_text_id *Formatted_text_id) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(formatted_text_id.GongGetReferenceIdentifier(stage))
}

func (forward *Forward) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(forward.GongGetReferenceIdentifier(stage))
}

func (frame *Frame) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(frame.GongGetReferenceIdentifier(stage))
}

func (frame_note *Frame_note) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(frame_note.GongGetReferenceIdentifier(stage))
}

func (fret *Fret) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(fret.GongGetReferenceIdentifier(stage))
}

func (glass *Glass) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(glass.GongGetReferenceIdentifier(stage))
}

func (glissando *Glissando) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(glissando.GongGetReferenceIdentifier(stage))
}

func (glyph *Glyph) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(glyph.GongGetReferenceIdentifier(stage))
}

func (grace *Grace) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(grace.GongGetReferenceIdentifier(stage))
}

func (group_barline *Group_barline) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(group_barline.GongGetReferenceIdentifier(stage))
}

func (group_name *Group_name) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(group_name.GongGetReferenceIdentifier(stage))
}

func (group_symbol *Group_symbol) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(group_symbol.GongGetReferenceIdentifier(stage))
}

func (grouping *Grouping) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(grouping.GongGetReferenceIdentifier(stage))
}

func (hammer_on_pull_off *Hammer_on_pull_off) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(hammer_on_pull_off.GongGetReferenceIdentifier(stage))
}

func (handbell *Handbell) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(handbell.GongGetReferenceIdentifier(stage))
}

func (harmon_closed *Harmon_closed) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(harmon_closed.GongGetReferenceIdentifier(stage))
}

func (harmon_mute *Harmon_mute) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(harmon_mute.GongGetReferenceIdentifier(stage))
}

func (harmonic *Harmonic) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(harmonic.GongGetReferenceIdentifier(stage))
}

func (harmony *Harmony) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(harmony.GongGetReferenceIdentifier(stage))
}

func (harmony_alter *Harmony_alter) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(harmony_alter.GongGetReferenceIdentifier(stage))
}

func (harp_pedals *Harp_pedals) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(harp_pedals.GongGetReferenceIdentifier(stage))
}

func (heel_toe *Heel_toe) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(heel_toe.GongGetReferenceIdentifier(stage))
}

func (hole *Hole) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(hole.GongGetReferenceIdentifier(stage))
}

func (hole_closed *Hole_closed) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(hole_closed.GongGetReferenceIdentifier(stage))
}

func (horizontal_turn *Horizontal_turn) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(horizontal_turn.GongGetReferenceIdentifier(stage))
}

func (identification *Identification) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(identification.GongGetReferenceIdentifier(stage))
}

func (image *Image) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(image.GongGetReferenceIdentifier(stage))
}

func (instrument *Instrument) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(instrument.GongGetReferenceIdentifier(stage))
}

func (instrument_change *Instrument_change) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(instrument_change.GongGetReferenceIdentifier(stage))
}

func (instrument_link *Instrument_link) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(instrument_link.GongGetReferenceIdentifier(stage))
}

func (interchangeable *Interchangeable) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(interchangeable.GongGetReferenceIdentifier(stage))
}

func (inversion *Inversion) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(inversion.GongGetReferenceIdentifier(stage))
}

func (key *Key) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(key.GongGetReferenceIdentifier(stage))
}

func (key_accidental *Key_accidental) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(key_accidental.GongGetReferenceIdentifier(stage))
}

func (key_octave *Key_octave) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(key_octave.GongGetReferenceIdentifier(stage))
}

func (kind *Kind) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(kind.GongGetReferenceIdentifier(stage))
}

func (level *Level) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(level.GongGetReferenceIdentifier(stage))
}

func (line_detail *Line_detail) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(line_detail.GongGetReferenceIdentifier(stage))
}

func (line_width *Line_width) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(line_width.GongGetReferenceIdentifier(stage))
}

func (link *Link) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(link.GongGetReferenceIdentifier(stage))
}

func (listen *Listen) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(listen.GongGetReferenceIdentifier(stage))
}

func (listening *Listening) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(listening.GongGetReferenceIdentifier(stage))
}

func (lyric *Lyric) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(lyric.GongGetReferenceIdentifier(stage))
}

func (lyric_font *Lyric_font) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(lyric_font.GongGetReferenceIdentifier(stage))
}

func (lyric_language *Lyric_language) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(lyric_language.GongGetReferenceIdentifier(stage))
}

func (measure_layout *Measure_layout) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(measure_layout.GongGetReferenceIdentifier(stage))
}

func (measure_numbering *Measure_numbering) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(measure_numbering.GongGetReferenceIdentifier(stage))
}

func (measure_repeat *Measure_repeat) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(measure_repeat.GongGetReferenceIdentifier(stage))
}

func (measure_style *Measure_style) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(measure_style.GongGetReferenceIdentifier(stage))
}

func (membrane *Membrane) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(membrane.GongGetReferenceIdentifier(stage))
}

func (metal *Metal) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(metal.GongGetReferenceIdentifier(stage))
}

func (metronome *Metronome) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(metronome.GongGetReferenceIdentifier(stage))
}

func (metronome_beam *Metronome_beam) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(metronome_beam.GongGetReferenceIdentifier(stage))
}

func (metronome_note *Metronome_note) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(metronome_note.GongGetReferenceIdentifier(stage))
}

func (metronome_tied *Metronome_tied) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(metronome_tied.GongGetReferenceIdentifier(stage))
}

func (metronome_tuplet *Metronome_tuplet) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(metronome_tuplet.GongGetReferenceIdentifier(stage))
}

func (midi_device *Midi_device) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(midi_device.GongGetReferenceIdentifier(stage))
}

func (midi_instrument *Midi_instrument) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(midi_instrument.GongGetReferenceIdentifier(stage))
}

func (miscellaneous *Miscellaneous) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(miscellaneous.GongGetReferenceIdentifier(stage))
}

func (miscellaneous_field *Miscellaneous_field) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(miscellaneous_field.GongGetReferenceIdentifier(stage))
}

func (mordent *Mordent) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(mordent.GongGetReferenceIdentifier(stage))
}

func (multiple_rest *Multiple_rest) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(multiple_rest.GongGetReferenceIdentifier(stage))
}

func (name_display *Name_display) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(name_display.GongGetReferenceIdentifier(stage))
}

func (non_arpeggiate *Non_arpeggiate) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(non_arpeggiate.GongGetReferenceIdentifier(stage))
}

func (notations *Notations) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notations.GongGetReferenceIdentifier(stage))
}

func (note *Note) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(note.GongGetReferenceIdentifier(stage))
}

func (note_size *Note_size) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(note_size.GongGetReferenceIdentifier(stage))
}

func (note_type *Note_type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(note_type.GongGetReferenceIdentifier(stage))
}

func (notehead *Notehead) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notehead.GongGetReferenceIdentifier(stage))
}

func (notehead_text *Notehead_text) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(notehead_text.GongGetReferenceIdentifier(stage))
}

func (numeral *Numeral) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(numeral.GongGetReferenceIdentifier(stage))
}

func (numeral_key *Numeral_key) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(numeral_key.GongGetReferenceIdentifier(stage))
}

func (numeral_root *Numeral_root) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(numeral_root.GongGetReferenceIdentifier(stage))
}

func (octave_shift *Octave_shift) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(octave_shift.GongGetReferenceIdentifier(stage))
}

func (offset *Offset) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(offset.GongGetReferenceIdentifier(stage))
}

func (opus *Opus) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(opus.GongGetReferenceIdentifier(stage))
}

func (ornaments *Ornaments) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(ornaments.GongGetReferenceIdentifier(stage))
}

func (other_appearance *Other_appearance) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(other_appearance.GongGetReferenceIdentifier(stage))
}

func (other_direction *Other_direction) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(other_direction.GongGetReferenceIdentifier(stage))
}

func (other_listening *Other_listening) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(other_listening.GongGetReferenceIdentifier(stage))
}

func (other_notation *Other_notation) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(other_notation.GongGetReferenceIdentifier(stage))
}

func (other_placement_text *Other_placement_text) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(other_placement_text.GongGetReferenceIdentifier(stage))
}

func (other_play *Other_play) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(other_play.GongGetReferenceIdentifier(stage))
}

func (other_text *Other_text) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(other_text.GongGetReferenceIdentifier(stage))
}

func (page_layout *Page_layout) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(page_layout.GongGetReferenceIdentifier(stage))
}

func (page_margins *Page_margins) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(page_margins.GongGetReferenceIdentifier(stage))
}

func (part_clef *Part_clef) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(part_clef.GongGetReferenceIdentifier(stage))
}

func (part_group *Part_group) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(part_group.GongGetReferenceIdentifier(stage))
}

func (part_link *Part_link) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(part_link.GongGetReferenceIdentifier(stage))
}

func (part_list *Part_list) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(part_list.GongGetReferenceIdentifier(stage))
}

func (part_name *Part_name) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(part_name.GongGetReferenceIdentifier(stage))
}

func (part_symbol *Part_symbol) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(part_symbol.GongGetReferenceIdentifier(stage))
}

func (part_transpose *Part_transpose) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(part_transpose.GongGetReferenceIdentifier(stage))
}

func (pedal *Pedal) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pedal.GongGetReferenceIdentifier(stage))
}

func (pedal_tuning *Pedal_tuning) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pedal_tuning.GongGetReferenceIdentifier(stage))
}

func (per_minute *Per_minute) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(per_minute.GongGetReferenceIdentifier(stage))
}

func (percussion *Percussion) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(percussion.GongGetReferenceIdentifier(stage))
}

func (pitch *Pitch) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pitch.GongGetReferenceIdentifier(stage))
}

func (pitched *Pitched) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(pitched.GongGetReferenceIdentifier(stage))
}

func (placement_text *Placement_text) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(placement_text.GongGetReferenceIdentifier(stage))
}

func (play *Play) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(play.GongGetReferenceIdentifier(stage))
}

func (player *Player) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(player.GongGetReferenceIdentifier(stage))
}

func (principal_voice *Principal_voice) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(principal_voice.GongGetReferenceIdentifier(stage))
}

func (print *Print) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(print.GongGetReferenceIdentifier(stage))
}

func (release *Release) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(release.GongGetReferenceIdentifier(stage))
}

func (repeat *Repeat) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(repeat.GongGetReferenceIdentifier(stage))
}

func (rest *Rest) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(rest.GongGetReferenceIdentifier(stage))
}

func (root *Root) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(root.GongGetReferenceIdentifier(stage))
}

func (root_step *Root_step) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(root_step.GongGetReferenceIdentifier(stage))
}

func (scaling *Scaling) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(scaling.GongGetReferenceIdentifier(stage))
}

func (scordatura *Scordatura) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(scordatura.GongGetReferenceIdentifier(stage))
}

func (score_instrument *Score_instrument) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(score_instrument.GongGetReferenceIdentifier(stage))
}

func (score_part *Score_part) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(score_part.GongGetReferenceIdentifier(stage))
}

func (score_partwise *Score_partwise) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(score_partwise.GongGetReferenceIdentifier(stage))
}

func (score_timewise *Score_timewise) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(score_timewise.GongGetReferenceIdentifier(stage))
}

func (segno *Segno) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(segno.GongGetReferenceIdentifier(stage))
}

func (slash *Slash) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(slash.GongGetReferenceIdentifier(stage))
}

func (slide *Slide) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(slide.GongGetReferenceIdentifier(stage))
}

func (slur *Slur) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(slur.GongGetReferenceIdentifier(stage))
}

func (sound *Sound) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(sound.GongGetReferenceIdentifier(stage))
}

func (staff_details *Staff_details) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(staff_details.GongGetReferenceIdentifier(stage))
}

func (staff_divide *Staff_divide) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(staff_divide.GongGetReferenceIdentifier(stage))
}

func (staff_layout *Staff_layout) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(staff_layout.GongGetReferenceIdentifier(stage))
}

func (staff_size *Staff_size) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(staff_size.GongGetReferenceIdentifier(stage))
}

func (staff_tuning *Staff_tuning) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(staff_tuning.GongGetReferenceIdentifier(stage))
}

func (stem *Stem) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stem.GongGetReferenceIdentifier(stage))
}

func (stick *Stick) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(stick.GongGetReferenceIdentifier(stage))
}

func (string_mute *String_mute) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(string_mute.GongGetReferenceIdentifier(stage))
}

func (string_type *String_type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(string_type.GongGetReferenceIdentifier(stage))
}

func (strong_accent *Strong_accent) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(strong_accent.GongGetReferenceIdentifier(stage))
}

func (style_text *Style_text) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(style_text.GongGetReferenceIdentifier(stage))
}

func (supports *Supports) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(supports.GongGetReferenceIdentifier(stage))
}

func (swing *Swing) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(swing.GongGetReferenceIdentifier(stage))
}

func (sync *Sync) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(sync.GongGetReferenceIdentifier(stage))
}

func (system_dividers *System_dividers) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(system_dividers.GongGetReferenceIdentifier(stage))
}

func (system_layout *System_layout) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(system_layout.GongGetReferenceIdentifier(stage))
}

func (system_margins *System_margins) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(system_margins.GongGetReferenceIdentifier(stage))
}

func (tap *Tap) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tap.GongGetReferenceIdentifier(stage))
}

func (technical *Technical) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(technical.GongGetReferenceIdentifier(stage))
}

func (text_element_data *Text_element_data) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(text_element_data.GongGetReferenceIdentifier(stage))
}

func (tie *Tie) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tie.GongGetReferenceIdentifier(stage))
}

func (tied *Tied) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tied.GongGetReferenceIdentifier(stage))
}

func (time *Time) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(time.GongGetReferenceIdentifier(stage))
}

func (time_modification *Time_modification) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(time_modification.GongGetReferenceIdentifier(stage))
}

func (timpani *Timpani) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(timpani.GongGetReferenceIdentifier(stage))
}

func (transpose *Transpose) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(transpose.GongGetReferenceIdentifier(stage))
}

func (tremolo *Tremolo) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tremolo.GongGetReferenceIdentifier(stage))
}

func (tuplet *Tuplet) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tuplet.GongGetReferenceIdentifier(stage))
}

func (tuplet_dot *Tuplet_dot) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tuplet_dot.GongGetReferenceIdentifier(stage))
}

func (tuplet_number *Tuplet_number) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tuplet_number.GongGetReferenceIdentifier(stage))
}

func (tuplet_portion *Tuplet_portion) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tuplet_portion.GongGetReferenceIdentifier(stage))
}

func (tuplet_type *Tuplet_type) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(tuplet_type.GongGetReferenceIdentifier(stage))
}

func (typed_text *Typed_text) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(typed_text.GongGetReferenceIdentifier(stage))
}

func (unpitched *Unpitched) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(unpitched.GongGetReferenceIdentifier(stage))
}

func (virtual_instrument *Virtual_instrument) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(virtual_instrument.GongGetReferenceIdentifier(stage))
}

func (wait *Wait) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(wait.GongGetReferenceIdentifier(stage))
}

func (wavy_line *Wavy_line) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(wavy_line.GongGetReferenceIdentifier(stage))
}

func (wedge *Wedge) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(wedge.GongGetReferenceIdentifier(stage))
}

func (wood *Wood) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(wood.GongGetReferenceIdentifier(stage))
}

func (work *Work) GongMarshallUnstaging(stage *Stage) string {
	return __gong__marshallUnstaging(work.GongGetReferenceIdentifier(stage))
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

func __gong__appendInstances[T interface {
	comparable
	GongstructIF
}](res []GongstructIF, m map[T]struct{}) []GongstructIF {
	for instance := range m {
		res = append(res, instance)
	}
	return res
}

func __gong__getUUID(stage *Stage, instance GongstructIF) string {
	if __gong__, ok := any(instance).(interface{ GongGetUUIDCustom(stage *Stage) string }); ok {
		return __gong__.GongGetUUIDCustom(stage)
	}
	return GongGenerateReproducibleUUIDv4(GongGetGongstructNameFromPointer(instance), uint64(stage.GetOrder(instance)))
}

func __gong__computeReferencePass1[T interface {
	comparable
	GongstructIF
}](
	stage *Stage,
	staged map[T]struct{},
	ref *map[T]T,
	refOrder *map[T]uint,
	inst *map[T]T,
) {
	*ref = make(map[T]T, len(staged))
	*refOrder = make(map[T]uint, len(staged))
	*inst = make(map[T]T, len(staged))
	for instance := range staged {
		_copy := instance.GongCopy().(T)
		(*ref)[instance] = _copy
		(*inst)[_copy] = instance
		(*refOrder)[_copy] = instance.GongGetOrder(stage)
	}
}

func __gong__computeReferencePass2[T interface {
	comparable
	GongstructIF
	GongReconstructPointersFromReferences(*Stage, T)
}](staged map[T]struct{}, reference map[T]T, stage *Stage) {
	for instance := range staged {
		reference[instance].GongReconstructPointersFromReferences(stage, instance)
	}
}

func __gong__getOrder[T comparable](stagedOrder, refOrder map[T]uint, instance T, typeName string) uint {
	if order, ok := stagedOrder[instance]; ok {
		return order
	}
	if order, ok := refOrder[instance]; ok {
		return order
	}
	log.Printf("instance %p of type %s was not staged and does not have a reference order", any(instance), typeName)
	return 0
}

func __gong__formatIdentifier(s GongstructIF, order uint) string {
	return fmt.Sprintf("__%s__%08d_", s.GongGetGongstructName(), order)
}

func __gong__marshallIdentifier(identifier, structName, name string) string {
	decl := strings.ReplaceAll(GongIdentifiersDecls, "{{Identifier}}", identifier)
	decl = strings.ReplaceAll(decl, "{{GeneratedStructName}}", structName)
	return strings.ReplaceAll(decl, "{{GeneratedFieldNameValue}}", __gong__toRawStringLiteral(name))
}

func __gong__marshallUnstaging(identifier string) string {
	return strings.ReplaceAll(GongUnstageStmt, "{{Identifier}}", identifier)
}

// end of template
