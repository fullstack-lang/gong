// generated code - do not edit
package models

import (
	"cmp"
	"fmt"
	"slices"

	"github.com/xuri/excelize/v2"
)

func (stage *Stage) SerializeStage(filename string) {
	stage.SerializeStage2(filename, false)
}

func (stage *Stage) SerializeStage2(filename string, addIDs bool) {
	f := stage.__gong__buildExcelizeFile(addIDs)
	if err := f.SaveAs(filename); err != nil {
		fmt.Println("cannot write xl file : ", err)
	}
}

func (stage *Stage) __gong__buildExcelizeFile(addIDs bool) *excelize.File {
	f := excelize.NewFile()
	{
		// insertion point
		{
			var instances []GongstructIF
			for instance := range stage.A_directives {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_directive", instances, (*A_directive)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_measures {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_measure", instances, (*A_measure)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_measure_1s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_measure_1", instances, (*A_measure_1)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_parts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_part", instances, (*A_part)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.A_part_1s {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "A_part_1", instances, (*A_part_1)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Accidentals {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Accidental", instances, (*Accidental)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Accidental_marks {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Accidental_mark", instances, (*Accidental_mark)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Accidental_texts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Accidental_text", instances, (*Accidental_text)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Accords {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Accord", instances, (*Accord)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Accordion_registrations {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Accordion_registration", instances, (*Accordion_registration)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Appearances {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Appearance", instances, (*Appearance)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Arpeggiates {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Arpeggiate", instances, (*Arpeggiate)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Arrows {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Arrow", instances, (*Arrow)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Articulationss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Articulations", instances, (*Articulations)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Assesss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Assess", instances, (*Assess)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Attributess {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Attributes", instances, (*Attributes)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Backups {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Backup", instances, (*Backup)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Bar_style_colors {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Bar_style_color", instances, (*Bar_style_color)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Barlines {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Barline", instances, (*Barline)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Barres {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Barre", instances, (*Barre)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Basss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Bass", instances, (*Bass)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Bass_steps {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Bass_step", instances, (*Bass_step)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Beams {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Beam", instances, (*Beam)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Beat_repeats {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Beat_repeat", instances, (*Beat_repeat)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Beat_unit_tieds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Beat_unit_tied", instances, (*Beat_unit_tied)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Beaters {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Beater", instances, (*Beater)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Bends {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Bend", instances, (*Bend)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Bookmarks {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Bookmark", instances, (*Bookmark)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Brackets {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Bracket", instances, (*Bracket)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Breath_marks {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Breath_mark", instances, (*Breath_mark)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Caesuras {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Caesura", instances, (*Caesura)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Cancels {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Cancel", instances, (*Cancel)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Clefs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Clef", instances, (*Clef)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Codas {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Coda", instances, (*Coda)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Credits {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Credit", instances, (*Credit)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Dashess {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Dashes", instances, (*Dashes)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Defaultss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Defaults", instances, (*Defaults)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Degrees {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Degree", instances, (*Degree)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Degree_alters {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Degree_alter", instances, (*Degree_alter)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Degree_types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Degree_type", instances, (*Degree_type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Degree_values {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Degree_value", instances, (*Degree_value)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Directions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Direction", instances, (*Direction)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Direction_types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Direction_type", instances, (*Direction_type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Distances {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Distance", instances, (*Distance)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Doubles {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Double", instances, (*Double)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Dynamicss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Dynamics", instances, (*Dynamics)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Effects {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Effect", instances, (*Effect)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Elisions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Elision", instances, (*Elision)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Emptys {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Empty", instances, (*Empty)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Empty_fonts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Empty_font", instances, (*Empty_font)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Empty_lines {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Empty_line", instances, (*Empty_line)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Empty_placements {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Empty_placement", instances, (*Empty_placement)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Empty_placement_smufls {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Empty_placement_smufl", instances, (*Empty_placement_smufl)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Empty_print_object_style_aligns {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Empty_print_object_style_align", instances, (*Empty_print_object_style_align)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Empty_print_styles {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Empty_print_style", instances, (*Empty_print_style)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Empty_print_style_aligns {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Empty_print_style_align", instances, (*Empty_print_style_align)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Empty_print_style_align_ids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Empty_print_style_align_id", instances, (*Empty_print_style_align_id)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Empty_trill_sounds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Empty_trill_sound", instances, (*Empty_trill_sound)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Encodings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Encoding", instances, (*Encoding)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Endings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Ending", instances, (*Ending)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Extends {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Extend", instances, (*Extend)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Features {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Feature", instances, (*Feature)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Fermatas {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Fermata", instances, (*Fermata)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Figures {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Figure", instances, (*Figure)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Figured_basss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Figured_bass", instances, (*Figured_bass)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Fingerings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Fingering", instances, (*Fingering)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.First_frets {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "First_fret", instances, (*First_fret)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.For_parts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "For_part", instances, (*For_part)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Formatted_symbols {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Formatted_symbol", instances, (*Formatted_symbol)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Formatted_symbol_ids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Formatted_symbol_id", instances, (*Formatted_symbol_id)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Formatted_texts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Formatted_text", instances, (*Formatted_text)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Formatted_text_ids {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Formatted_text_id", instances, (*Formatted_text_id)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Forwards {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Forward", instances, (*Forward)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Frames {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Frame", instances, (*Frame)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Frame_notes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Frame_note", instances, (*Frame_note)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Frets {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Fret", instances, (*Fret)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Glasss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Glass", instances, (*Glass)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Glissandos {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Glissando", instances, (*Glissando)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Glyphs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Glyph", instances, (*Glyph)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Graces {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Grace", instances, (*Grace)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Group_barlines {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Group_barline", instances, (*Group_barline)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Group_names {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Group_name", instances, (*Group_name)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Group_symbols {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Group_symbol", instances, (*Group_symbol)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Groupings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Grouping", instances, (*Grouping)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Hammer_on_pull_offs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Hammer_on_pull_off", instances, (*Hammer_on_pull_off)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Handbells {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Handbell", instances, (*Handbell)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Harmon_closeds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Harmon_closed", instances, (*Harmon_closed)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Harmon_mutes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Harmon_mute", instances, (*Harmon_mute)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Harmonics {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Harmonic", instances, (*Harmonic)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Harmonys {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Harmony", instances, (*Harmony)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Harmony_alters {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Harmony_alter", instances, (*Harmony_alter)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Harp_pedalss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Harp_pedals", instances, (*Harp_pedals)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Heel_toes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Heel_toe", instances, (*Heel_toe)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Holes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Hole", instances, (*Hole)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Hole_closeds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Hole_closed", instances, (*Hole_closed)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Horizontal_turns {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Horizontal_turn", instances, (*Horizontal_turn)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Identifications {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Identification", instances, (*Identification)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Images {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Image", instances, (*Image)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Instruments {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Instrument", instances, (*Instrument)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Instrument_changes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Instrument_change", instances, (*Instrument_change)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Instrument_links {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Instrument_link", instances, (*Instrument_link)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Interchangeables {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Interchangeable", instances, (*Interchangeable)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Inversions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Inversion", instances, (*Inversion)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Keys {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Key", instances, (*Key)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Key_accidentals {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Key_accidental", instances, (*Key_accidental)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Key_octaves {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Key_octave", instances, (*Key_octave)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Kinds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Kind", instances, (*Kind)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Levels {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Level", instances, (*Level)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Line_details {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Line_detail", instances, (*Line_detail)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Line_widths {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Line_width", instances, (*Line_width)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Links {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Link", instances, (*Link)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Listens {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Listen", instances, (*Listen)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Listenings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Listening", instances, (*Listening)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Lyrics {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Lyric", instances, (*Lyric)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Lyric_fonts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Lyric_font", instances, (*Lyric_font)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Lyric_languages {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Lyric_language", instances, (*Lyric_language)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Measure_layouts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Measure_layout", instances, (*Measure_layout)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Measure_numberings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Measure_numbering", instances, (*Measure_numbering)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Measure_repeats {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Measure_repeat", instances, (*Measure_repeat)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Measure_styles {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Measure_style", instances, (*Measure_style)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Membranes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Membrane", instances, (*Membrane)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Metals {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Metal", instances, (*Metal)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Metronomes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Metronome", instances, (*Metronome)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Metronome_beams {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Metronome_beam", instances, (*Metronome_beam)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Metronome_notes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Metronome_note", instances, (*Metronome_note)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Metronome_tieds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Metronome_tied", instances, (*Metronome_tied)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Metronome_tuplets {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Metronome_tuplet", instances, (*Metronome_tuplet)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Midi_devices {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Midi_device", instances, (*Midi_device)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Midi_instruments {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Midi_instrument", instances, (*Midi_instrument)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Miscellaneouss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Miscellaneous", instances, (*Miscellaneous)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Miscellaneous_fields {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Miscellaneous_field", instances, (*Miscellaneous_field)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Mordents {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Mordent", instances, (*Mordent)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Multiple_rests {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Multiple_rest", instances, (*Multiple_rest)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Name_displays {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Name_display", instances, (*Name_display)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Non_arpeggiates {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Non_arpeggiate", instances, (*Non_arpeggiate)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Notationss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Notations", instances, (*Notations)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Notes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Note", instances, (*Note)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Note_sizes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Note_size", instances, (*Note_size)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Note_types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Note_type", instances, (*Note_type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Noteheads {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Notehead", instances, (*Notehead)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Notehead_texts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Notehead_text", instances, (*Notehead_text)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Numerals {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Numeral", instances, (*Numeral)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Numeral_keys {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Numeral_key", instances, (*Numeral_key)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Numeral_roots {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Numeral_root", instances, (*Numeral_root)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Octave_shifts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Octave_shift", instances, (*Octave_shift)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Offsets {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Offset", instances, (*Offset)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Opuss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Opus", instances, (*Opus)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Ornamentss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Ornaments", instances, (*Ornaments)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Other_appearances {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Other_appearance", instances, (*Other_appearance)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Other_directions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Other_direction", instances, (*Other_direction)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Other_listenings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Other_listening", instances, (*Other_listening)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Other_notations {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Other_notation", instances, (*Other_notation)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Other_placement_texts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Other_placement_text", instances, (*Other_placement_text)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Other_plays {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Other_play", instances, (*Other_play)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Other_texts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Other_text", instances, (*Other_text)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Page_layouts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Page_layout", instances, (*Page_layout)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Page_marginss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Page_margins", instances, (*Page_margins)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Part_clefs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Part_clef", instances, (*Part_clef)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Part_groups {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Part_group", instances, (*Part_group)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Part_links {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Part_link", instances, (*Part_link)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Part_lists {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Part_list", instances, (*Part_list)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Part_names {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Part_name", instances, (*Part_name)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Part_symbols {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Part_symbol", instances, (*Part_symbol)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Part_transposes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Part_transpose", instances, (*Part_transpose)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Pedals {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Pedal", instances, (*Pedal)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Pedal_tunings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Pedal_tuning", instances, (*Pedal_tuning)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Per_minutes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Per_minute", instances, (*Per_minute)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Percussions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Percussion", instances, (*Percussion)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Pitchs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Pitch", instances, (*Pitch)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Pitcheds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Pitched", instances, (*Pitched)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Placement_texts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Placement_text", instances, (*Placement_text)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Plays {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Play", instances, (*Play)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Players {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Player", instances, (*Player)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Principal_voices {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Principal_voice", instances, (*Principal_voice)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Prints {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Print", instances, (*Print)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Releases {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Release", instances, (*Release)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Repeats {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Repeat", instances, (*Repeat)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Rests {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Rest", instances, (*Rest)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Roots {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Root", instances, (*Root)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Root_steps {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Root_step", instances, (*Root_step)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Scalings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Scaling", instances, (*Scaling)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Scordaturas {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Scordatura", instances, (*Scordatura)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Score_instruments {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Score_instrument", instances, (*Score_instrument)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Score_parts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Score_part", instances, (*Score_part)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Score_partwises {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Score_partwise", instances, (*Score_partwise)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Score_timewises {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Score_timewise", instances, (*Score_timewise)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Segnos {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Segno", instances, (*Segno)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Slashs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Slash", instances, (*Slash)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Slides {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Slide", instances, (*Slide)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Slurs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Slur", instances, (*Slur)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Sounds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Sound", instances, (*Sound)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Staff_detailss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Staff_details", instances, (*Staff_details)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Staff_divides {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Staff_divide", instances, (*Staff_divide)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Staff_layouts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Staff_layout", instances, (*Staff_layout)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Staff_sizes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Staff_size", instances, (*Staff_size)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Staff_tunings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Staff_tuning", instances, (*Staff_tuning)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Stems {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Stem", instances, (*Stem)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Sticks {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Stick", instances, (*Stick)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.String_mutes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "String_mute", instances, (*String_mute)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.String_types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "String_type", instances, (*String_type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Strong_accents {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Strong_accent", instances, (*Strong_accent)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Style_texts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Style_text", instances, (*Style_text)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Supportss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Supports", instances, (*Supports)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Swings {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Swing", instances, (*Swing)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Syncs {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Sync", instances, (*Sync)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.System_dividerss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "System_dividers", instances, (*System_dividers)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.System_layouts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "System_layout", instances, (*System_layout)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.System_marginss {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "System_margins", instances, (*System_margins)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Taps {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Tap", instances, (*Tap)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Technicals {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Technical", instances, (*Technical)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Text_element_datas {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Text_element_data", instances, (*Text_element_data)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Ties {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Tie", instances, (*Tie)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Tieds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Tied", instances, (*Tied)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Times {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Time", instances, (*Time)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Time_modifications {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Time_modification", instances, (*Time_modification)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Timpanis {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Timpani", instances, (*Timpani)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Transposes {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Transpose", instances, (*Transpose)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Tremolos {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Tremolo", instances, (*Tremolo)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Tuplets {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Tuplet", instances, (*Tuplet)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Tuplet_dots {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Tuplet_dot", instances, (*Tuplet_dot)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Tuplet_numbers {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Tuplet_number", instances, (*Tuplet_number)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Tuplet_portions {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Tuplet_portion", instances, (*Tuplet_portion)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Tuplet_types {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Tuplet_type", instances, (*Tuplet_type)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Typed_texts {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Typed_text", instances, (*Typed_text)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Unpitcheds {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Unpitched", instances, (*Unpitched)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Virtual_instruments {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Virtual_instrument", instances, (*Virtual_instrument)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Waits {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Wait", instances, (*Wait)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Wavy_lines {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Wavy_line", instances, (*Wavy_line)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Wedges {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Wedge", instances, (*Wedge)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Woods {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Wood", instances, (*Wood)(nil).GongGetFieldHeaders(), addIDs)
		}
		{
			var instances []GongstructIF
			for instance := range stage.Works {
				instances = append(instances, instance)
			}
			stage.SerializeExcelize(f, "Work", instances, (*Work)(nil).GongGetFieldHeaders(), addIDs)
		}
	}

	// Create a style with wrap text enabled
	wrapStyle, err := f.NewStyle(&excelize.Style{
		Alignment: &excelize.Alignment{
			WrapText: true,
		},
	})
	_ = wrapStyle
	if err != nil {
		fmt.Println("failed to create style:", err)
		return f
	}

	// Create a style with bold text
	boldStyle, err := f.NewStyle(&excelize.Style{
		Font: &excelize.Font{
			Bold: true,
		},
	})
	_ = boldStyle
	if err != nil {
		fmt.Println("failed to create bold style:", err)
		return f
	}

	// Get all sheet names
	sheetList := f.GetSheetList()

	for _, sheet := range sheetList {
		// Use a lazy iterator instead of loading all rows into memory
		rows, err := f.Rows(sheet)
		if err != nil {
			fmt.Printf("failed to get rows iterator for sheet %q: %v\n", sheet, err)
			continue
		}

		// Check if there is at least one row, and move the iterator to it
		if !rows.Next() {
			rows.Close() // Always close iterators
			continue
		}

		// Read ONLY the first row
		firstRow, err := rows.Columns()

		// Close the iterator immediately since we don't need the rest of the sheet
		rows.Close()

		if err != nil {
			fmt.Printf("failed to get columns for sheet %q: %v\n", sheet, err)
			continue
		}

		// If the first row is completely empty, skip
		if len(firstRow) == 0 {
			continue
		}

		// Track the first and last “used” column in the first row,
		// so we can later apply an AutoFilter from the first to last used col
		var firstUsedColIdx, lastUsedColIdx int

		for colIdx, cellValue := range firstRow {
			if cellValue == "" {
				// Skip columns with empty first-row cells
				continue
			}

			// Convert zero-based colIdx to 1-based for Excelize,
			// then get the column name (A, B, C, etc.)
			colName, err := excelize.ColumnNumberToName(colIdx + 1)
			if err != nil {
				fmt.Printf("failed to convert column number: %v\n", err)
				continue
			}

			// Apply wrap-text style to this entire column
			colRange := colName + ":" + colName
			if err := f.SetColStyle(sheet, colRange, wrapStyle); err != nil {
				fmt.Printf("failed to set col style on %s: %v\n", colRange, err)
				continue
			}

			// Make the first row (cell in row 1) bold in this column
			cellRef := fmt.Sprintf("%s1", colName)
			if err := f.SetCellStyle(sheet, cellRef, cellRef, boldStyle); err != nil {
				fmt.Printf("failed to set cell style on %s: %v\n", cellRef, err)
				continue
			}

			// Update our “first used” and “last used” column indices
			if firstUsedColIdx == 0 {
				firstUsedColIdx = colIdx + 1
			}
			if colIdx+1 > lastUsedColIdx {
				lastUsedColIdx = colIdx + 1
			}
		}

		// If we found at least one non-empty column in row 1, enable AutoFilter
		if firstUsedColIdx != 0 && lastUsedColIdx >= firstUsedColIdx {
			startCol, _ := excelize.ColumnNumberToName(firstUsedColIdx)
			endCol, _ := excelize.ColumnNumberToName(lastUsedColIdx)
			styleRange := fmt.Sprintf("%s:%s", startCol, endCol)
			autoFilterRange := fmt.Sprintf("%s1:%s1", startCol, endCol)
			startCellString := fmt.Sprintf("%s1", startCol)
			endCellString := fmt.Sprintf("%s1", endCol)

			if err := f.SetColStyle(sheet, styleRange, wrapStyle); err != nil {
				fmt.Println("failed to set column style:", err)
				return f
			}

			// Apply the bold style to the first row (A1:XFD1)
			if err := f.SetCellStyle(sheet, startCellString, endCellString, boldStyle); err != nil {
				fmt.Println("failed to set bold style:", err)
				return f
			}

			var opts []excelize.AutoFilterOptions
			if err := f.AutoFilter(sheet, autoFilterRange, opts); err != nil {
				fmt.Printf("failed to enable auto filter on range %s: %v\n", autoFilterRange, err)
			}
		}
	}

	var tab ExcelizeTabulator
	tab.SetExcelizeFile(f)
	{
		f.DeleteSheet("Sheet1")
	}
	return f
}

// SerializeStageAsBytes serializes the stage to a pure in-memory Excel file and returns the bytes.
func (stage *Stage) SerializeStageAsBytes(addIDs bool) ([]byte, error) {
	f := stage.__gong__buildExcelizeFile(addIDs)
	buf, err := f.WriteToBuffer()
	if err != nil {
		return nil, err
	}
	return buf.Bytes(), nil
}

func __gong__shortenString(s string) string {
	if len(s) > 31 {
		return s[:31]
	}
	return s
}

// Tabulator is an interface for writing to a table strings
type Tabulator interface {
	AddSheet(sheetName string)
	AddRow(sheetName string) int
	AddCell(sheetName string, rowId, columnIndex int, value string)
}

type ExcelizeTabulator struct {
	f *excelize.File
}

func (tab *ExcelizeTabulator) SetExcelizeFile(f *excelize.File) {
	tab.f = f
}

func (tab *ExcelizeTabulator) AddSheet(sheetName string) {

}

func (tab *ExcelizeTabulator) AddRow(sheetName string) (rowId int) {
	return
}

func (tab *ExcelizeTabulator) AddCell(sheetName string, rowId, columnIndex int, value string) {

}

// SerializeExcelize is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelize(f *excelize.File, name string, instances []GongstructIF, fields []GongFieldHeader, addIDs bool) {
	sheetName := __gong__shortenString(name)

	// Create a new sheet.
	f.NewSheet(sheetName)

	sortedSlice := make([]GongstructIF, len(instances))
	copy(sortedSlice, instances)
	slices.SortFunc(sortedSlice, func(a, b GongstructIF) int {
		return cmp.Compare(a.GetName(), b.GetName())
	})

	line := 1

	for index, fieldHeader := range fields {
		if !addIDs {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(index+1)), line), fieldHeader.Name)
		} else {
			f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+1)), line), fieldHeader.Name)
			switch fieldHeader.GongFieldValueType {
			case GongFieldValueTypePointer:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":ID")
			case GongFieldValueTypeSliceOfPointers:
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line),
					fieldHeader.Name+":"+fieldHeader.TargetGongstructName+":IDs")
			default:
				// if index is 0, this is the ID of the instance
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), "ID")
				} else {
					// one have to put the type of the cell
					header := fieldHeader.Name
					switch fieldHeader.GongFieldValueType {
					case GongFieldValueTypeInt:
						header += ":int"
					case GongFieldValueTypeIntDuration:
						header += ":duration"
					case GongFieldValueTypeFloat:
						header += ":float"
					case GongFieldValueTypeBool:
						header += ":bool"
					case GongFieldValueTypeString:
						header += ":string"
					case GongFieldValueTypeDate:
						header += ":date"
					default:
						header += ":basicType"
					}
					header += ":noID"
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), header)
				}
			}
		}
	}

	// AutoFilter starting from A1
	f.AutoFilter(sheetName,
		fmt.Sprintf("%s%d", GongIntToLetters(int32(1)), line),
		[]excelize.AutoFilterOptions{})

	for _, instance := range sortedSlice {
		line = line + 1

		// 3. Add the ID value in column B

		for index, fieldName := range fields {
			fieldStringValue := stage.GetFieldStringValueFromPointer(instance, fieldName.Name)
			if !addIDs {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(index+1)), line), fieldStringValue.GetValueString())
			} else {
				f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+1)), line), fieldStringValue.GetValueString())
				if index == 0 {
					f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), instance.GongGetUUID(stage))
				} else {
					switch fieldStringValue.GongFieldValueType {
					case GongFieldValueTypePointer, GongFieldValueTypeSliceOfPointers:
						f.SetCellStr(sheetName, fmt.Sprintf("%s%d", GongIntToLetters(int32(2*index+2)), line), fieldStringValue.ids)
					}
				}

			}
		}
	}
}

// SerializeExcelizePointer is the Stage method for Excel serialization.
func (stage *Stage) SerializeExcelizePointer[Type PointerToGongstruct](f *excelize.File) {
	stage.SerializeExcelizePointer2[Type](f, false)
}

// SerializeExcelizePointer2 is the Stage method for Excel serialization with optional IDs.
func (stage *Stage) SerializeExcelizePointer2[Type PointerToGongstruct](f *excelize.File, addIDs bool) {
	var ret Type
	set := *stage.GetInstancesSet[Type]()
	var instances []GongstructIF
	for key := range set {
		instances = append(instances, key)
	}
	stage.SerializeExcelize(f, ret.GongGetGongstructName(), instances, ret.GongGetFieldHeaders(), addIDs)
}
